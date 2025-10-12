package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	config "github.com/oiweiwei/go-msrpc/config"
	config_flag "github.com/oiweiwei/go-msrpc/config/flag"
	"github.com/oiweiwei/go-msrpc/dcerpc"
	"github.com/oiweiwei/go-msrpc/msrpc/mgmt/mgmt/v1"
	"github.com/oiweiwei/go-msrpc/msrpc/well_known"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
)

var (
	cfg = config.New().DisableEPM()
)

var Statistics = [4]string{
	"RPC_C_STATS_CALLS_IN",
	"RPC_C_STATS_CALLS_OUT",
	"RPC_C_STATS_PKTS_IN",
	"RPC_C_STATS_PKTS_OUT",
}

func init() {

	config_flag.BindFlags(cfg, flag.CommandLine)
}

func main() {

	if err := config_flag.ParseAndValidate(cfg, flag.CommandLine); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	ctx := gssapi.NewSecurityContext(context.Background())

	cc, err := dcerpc.Dial(ctx, cfg.ServerAddr(), cfg.DialOptions(ctx)...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer cc.Close(ctx)

	cli2, err := mgmt.NewManagementClient(ctx, cc, cfg.ClientOptions(ctx)...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	stats, err := cli2.InquireStats(ctx, &mgmt.InquireStatsRequest{Count: 4})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for i, v := range stats.Statistics {
		if i >= len(Statistics) {
			fmt.Printf("unknown statistic %d: %d\n", i, v)
		}
		fmt.Printf("%s: %d\n", Statistics[i], v)
	}

	ifs, err := cli2.InquireInterfaceIDs(ctx, &mgmt.InquireInterfaceIDsRequest{})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for i, iff := range ifs.InterfaceIDVector.InterfaceID {
		fmt.Printf("[%d]: %s %s %d.%d\n", i, iff.UUID.UUID(), (*well_known.UUID)(iff.UUID.UUID()).Describe(), iff.VersMajor, iff.VersMinor)
	}

	resp2, err := cli2.InquirePrincName(ctx, &mgmt.InquirePrincNameRequest{
		AuthnProto:    uint32(dcerpc.AuthTypeWinNT),
		PrincNameSize: 1024},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("Server principal name:", resp2.PrincName)
}
