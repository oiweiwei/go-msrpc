//go:build exclude

// epm.go script dumps all interfaces registered on the remote server and also
// maps particular endpoints using Map method.
package main

import (
	"context"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strings"

	"github.com/oiweiwei/go-msrpc/dcerpc"
	"github.com/oiweiwei/go-msrpc/midl/uuid"
	"github.com/oiweiwei/go-msrpc/msrpc/dcetypes"
	"github.com/oiweiwei/go-msrpc/msrpc/dnsp/dnsserver/v5"
	"github.com/oiweiwei/go-msrpc/msrpc/epm/epm/v3"
	"github.com/oiweiwei/go-msrpc/msrpc/even/eventlog/v0"
	"github.com/oiweiwei/go-msrpc/msrpc/even6/ieventservice/v1"
	"github.com/oiweiwei/go-msrpc/msrpc/nrpc/logon/v1"
	"github.com/oiweiwei/go-msrpc/msrpc/well_known"
	"github.com/rs/zerolog"

	"github.com/oiweiwei/go-msrpc/ssp"
	"github.com/oiweiwei/go-msrpc/ssp/credential"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"

	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/win32"
)

var short bool
var cpu string

func init() {
	flag.BoolVar(&short, "s", false, "short form")
	flag.StringVar(&cpu, "cpupprof", "/tmp/cpprof.pprof", "cpu profile")
	flag.Parse()
}

func init() {
	// add credentials.
	gssapi.AddCredential(credential.NewFromNTHash(os.Getenv("USERNAME"), os.Getenv("PASSWORD_MD4")))
	// add mechanism.
	gssapi.AddMechanism(ssp.SPNEGO)
	gssapi.AddMechanism(ssp.NTLM)
}

func main() {

	if cpu != "" {
		f, err := os.Create(cpu)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	ctx := gssapi.NewSecurityContext(context.Background())

	cc, err := dcerpc.Dial(ctx, os.Getenv("SERVER"), well_known.EndpointMapper())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer cc.Close(ctx)

	cli, err := epm.NewEpmClient(ctx, cc, dcerpc.WithSeal(), dcerpc.WithTargetName(os.Getenv("TARGET")), dcerpc.WithVerifyBitMask(true), dcerpc.WithVerifyPresenetation(true), dcerpc.WithVerifyHeader2(true), dcerpc.WithLogger(zerolog.New(os.Stdout)))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	resp, err := cli.Lookup(ctx, &epm.LookupRequest{MaxEntries: 500})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	uConv := func(u *uuid.UUID) string {
		if n := strings.ToLower((*well_known.UUID)(u).Name()); n != "" {
			return n
		}
		return u.String()
	}

	_ = uConv

	if short {

		for i, entries := range resp.Entries {
			fmt.Printf("[%d] %s (%s)\n", i, entries.Tower.Binding().URL(uConv), entries.Tower.Binding().StringBinding)
		}

	} else {

		for i, entries := range resp.Entries {
			fmt.Println(i, " --- ", entries.Annotation, entries.Object)
			for i, floor := range entries.Tower.Floors() {
				if i == 0 {
					fmt.Println(i, floor, "//", (*well_known.UUID)(floor.UUID).Describe())
				} else {
					fmt.Println(i, floor)
				}
			}
		}
	}

	fmt.Println("logon")

	MapSyntax(ctx, cli, logon.LogonSyntaxV1_0)

	fmt.Println("dnsserver")

	MapSyntax(ctx, cli, dnsserver.DNSServerSyntaxV5_0)

	fmt.Println("epm")

	MapSyntax(ctx, cli, epm.EpmSyntaxV3_0)

	fmt.Println("ievent6")

	MapSyntax(ctx, cli, ieventservice.EventServiceSyntaxV1_0)

	fmt.Println("ievent")

	MapSyntax(ctx, cli, eventlog.EventlogSyntaxV0_0)
}

func MapSyntax(ctx context.Context, cli epm.EpmClient, syntax *dcerpc.SyntaxID) {

	resp, err := cli.Map(ctx, &epm.MapRequest{
		MapTower: dcetypes.FloorsToTower([]*dcetypes.Floor{
			{
				Protocol:     uint8(dcetypes.ProtocolUUID),
				UUID:         syntax.IfUUID,
				VersionMajor: syntax.IfVersionMajor,
				Data:         []byte{0, 0},
			},
			{
				Protocol:     uint8(dcetypes.ProtocolUUID),
				UUID:         dcerpc.TransferNDR,
				VersionMajor: syntax.IfVersionMajor,
				Data:         binary.LittleEndian.AppendUint16(nil, syntax.IfVersionMinor),
			},
			{
				Protocol: uint8(dcetypes.ProtocolRPC_CO),
				Data:     []byte{0, 0},
			},
			{
				Protocol: uint8(dcetypes.ProtocolTCP),
				Data:     []byte{0, 0},
			},
			{
				Protocol: uint8(dcetypes.ProtocolIP),
				Data:     []byte{0, 0, 0, 0},
			},
		}),
		MaxTowers: 100,
	})

	if err != nil {
		fmt.Println("error", err)
		return
	}

	for i, tower := range resp.Towers {
		fmt.Println(i, " --- ", "*")
		for i, floor := range tower.Floors() {
			fmt.Println(i, floor)
		}
	}
}
