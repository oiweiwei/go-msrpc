//go:build exclude

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/rs/zerolog"

	"github.com/oiweiwei/go-msrpc/dcerpc"
	"github.com/oiweiwei/go-msrpc/smb2"
	"github.com/oiweiwei/go-msrpc/ssp"
	"github.com/oiweiwei/go-msrpc/ssp/credential"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"

	"github.com/oiweiwei/go-msrpc/msrpc/epm/epm/v3"
	"github.com/oiweiwei/go-msrpc/msrpc/srvs/srvsvc/v3"

	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/ntstatus"
	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/win32"
)

func init() {
	// add credentials.
	gssapi.AddCredential(credential.NewFromPassword(os.Getenv("USERNAME"), os.Getenv("PASSWORD")))
	// add mechanism.
	gssapi.AddMechanism(ssp.SPNEGO)
	gssapi.AddMechanism(ssp.KRB5)
	gssapi.AddMechanism(ssp.NTLM)
}

func j(v any) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}

func main() {

	ctx := gssapi.NewSecurityContext(context.Background())

	log := zerolog.New(os.Stderr)

	dialer := smb2.NewDialer(
		smb2.WithDialect(smb2.SMB311),
		// smb2.WithSeal(),
		smb2.WithSecurity(gssapi.WithTargetName(os.Getenv("TARGET"))))

	opts := []dcerpc.Option{
		dcerpc.WithLogger(log),
		epm.EndpointMapper(ctx, os.Getenv("SERVER"), dcerpc.WithLogger(log)),
		dcerpc.WithSMBDialer(dialer),
	}

	cc, err := dcerpc.Dial(ctx, os.Getenv("SERVER"), opts...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	cli, err := srvsvc.NewSrvsvcClient(ctx, cc, dcerpc.WithInsecure())
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

	fmt.Println(j(enums))
}
