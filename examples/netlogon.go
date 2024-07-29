//go:build exclude

package main

// netlogon.go script gets the domain info from the remote server.
import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/oiweiwei/go-msrpc/dcerpc"
	"github.com/oiweiwei/go-msrpc/msrpc/epm/epm/v3"
	"github.com/oiweiwei/go-msrpc/msrpc/nrpc/logon/v1"
	"github.com/rs/zerolog"

	"github.com/oiweiwei/go-msrpc/ssp"
	"github.com/oiweiwei/go-msrpc/ssp/credential"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"

	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/win32"
)

func init() {
	// add credentials.
	gssapi.AddCredential(credential.NewFromPassword(os.Getenv("USERNAME"), os.Getenv("PASSWORD")))
	// add mechanism.
	gssapi.AddMechanism(ssp.SPNEGO)
	gssapi.AddMechanism(ssp.NTLM)
}

func j(v any) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}

func main() {

	ctx := gssapi.NewSecurityContext(context.Background())

	cc, err := dcerpc.Dial(ctx, os.Getenv("SERVER"), dcerpc.WithLogger(zerolog.New(os.Stdout)), epm.EndpointMapper(ctx, os.Getenv("SERVER"), dcerpc.WithLogger(zerolog.New(os.Stdout))))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	cli, err := logon.NewLogonClient(ctx, cc, dcerpc.WithSign())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	resp, err := cli.GetDCName(ctx, &logon.GetDCNameRequest{
		ComputerName: "PC$",
		Flags:        logon.DSReturnDNSName | logon.DSIPRequired,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Printf("%s\n", j(resp.DomainControllerInfo))

	// do valid call.
	trusts, err := cli.EnumerateDomainTrusts(ctx, &logon.EnumerateDomainTrustsRequest{
		ServerName: resp.DomainControllerInfo.DomainControllerName,
		Flags:      logon.TrustTypeForestMember,
	})
	if err != nil {

		fmt.Println(j(trusts))
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for _, dom := range trusts.Domains.Domains {
		fmt.Println(j(dom))
	}

}
