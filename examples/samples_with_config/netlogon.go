//go:build exclude

package main

// netlogon.go script gets the domain info from the remote server.
import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/oiweiwei/go-msrpc/dcerpc"

	"github.com/oiweiwei/go-msrpc/ssp/gssapi"

	config "github.com/oiweiwei/go-msrpc/config"
	config_flag "github.com/oiweiwei/go-msrpc/config/flag"

	"github.com/oiweiwei/go-msrpc/msrpc/nrpc/logon/v1"

	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/win32"

	. "github.com/oiweiwei/go-msrpc/examples/common"
)

var (
	cfg = config.New()
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
		fmt.Fprintln(os.Stderr, "dial", err)
		return
	}

	cli, err := logon.NewLogonClient(ctx, cc, cfg.ClientOptions(ctx)...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "new_logon_client", err)
		return
	}

	resp, err := cli.GetDCName(ctx, &logon.GetDCNameRequest{
		ComputerName: cfg.Workstation,
		Flags:        logon.DSReturnDNSName | logon.DSIPRequired,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println(J(resp.DomainControllerInfo))

	// do valid call.
	trusts, err := cli.EnumerateDomainTrusts(ctx, &logon.EnumerateDomainTrustsRequest{
		ServerName: resp.DomainControllerInfo.DomainControllerName,
		Flags:      logon.TrustTypeForestMember,
	})
	if err != nil {
		fmt.Println(J(trusts))
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for _, dom := range trusts.Domains.Domains {
		fmt.Println(J(dom))
	}
}
