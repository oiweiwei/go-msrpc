//go:build exclude

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/oiweiwei/go-msrpc/dcerpc"

	"github.com/oiweiwei/go-msrpc/ssp/gssapi"

	config "github.com/oiweiwei/go-msrpc/config"
	config_flag "github.com/oiweiwei/go-msrpc/config/flag"

	"github.com/oiweiwei/go-msrpc/msrpc/srvs/srvsvc/v3"

	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/ntstatus"
	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/win32"

	. "github.com/oiweiwei/go-msrpc/examples/common"
)

var (
	cfg = config.New().DisableEPM()
)

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

	cli, err := srvsvc.NewSrvsvcClient(ctx, cc, cfg.ClientOptions(ctx)...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	enums, err := cli.ShareEnum(ctx, &srvsvc.ShareEnumRequest{
		ServerName: "",
		Info: &srvsvc.ShareEnum{
			Level: 503,
			ShareInfo: &srvsvc.ShareEnumUnion{
				Value: &srvsvc.ShareEnumUnion_Level503{Level503: &srvsvc.ShareInfo503Container{}},
			},
		},
		PreferredMaximumLength: 0xffffffff,
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println(J(enums))
}
