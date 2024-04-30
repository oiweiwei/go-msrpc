//go:build exclude

// netlogon_sec_channel.go script establishes the secure channel and validates credentials
// provided via creds parameter. (specify creds as --creds 'DOMAIN\username%password').
//
// NOTE: SAM_USERNAME, and SAM_PASSWORD are machine account password acquired while joining
// domain. To join domain you must do following:
//
//	# https://manpages.ubuntu.com/manpages/trusty/man8/adcli.8.html
//	$ sudo adcli join CONTOSO.NET --domain-controller=${SERVER} --show-password --show-details
//
// Password will be located under computer-password key, and SAM_USERNAME must be set to computer-name + '$'.
// (if computer-name is MYPC, then SAM_USERNAME will be MYPC$)
package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/oiweiwei/go-msrpc/dcerpc"
	"github.com/oiweiwei/go-msrpc/ssp"
	"github.com/oiweiwei/go-msrpc/ssp/credential"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
	"github.com/oiweiwei/go-msrpc/ssp/ntlm"
	"github.com/rs/zerolog"

	"github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	"github.com/oiweiwei/go-msrpc/msrpc/epm/epm/v3"
	"github.com/oiweiwei/go-msrpc/msrpc/nrpc/logon/v1"

	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/ntstatus"
	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/win32"
)

func init() {
	// add credentials.
	gssapi.AddCredential(credential.NewFromPassword(os.Getenv("SAM_USERNAME"), os.Getenv("SAM_PASSWORD"), os.Getenv("SAM_WORKSTATION")))
	// add mechanism.
	gssapi.AddMechanism(ssp.SPNEGO)
	gssapi.AddMechanism(ssp.NTLM)
	gssapi.AddMechanism(ssp.Netlogon)
}

var uCred string

func init() {
	flag.StringVar(&uCred, "creds", "", "user credentials to verify")
	flag.Parse()
}

func j(v any) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}

func main() {

	sAMCred := credential.NewFromPassword(os.Getenv("SAM_USERNAME"), os.Getenv("SAM_PASSWORD"), os.Getenv("SAM_WORKSTATION"))

	ctx := gssapi.NewSecurityContext(context.Background())

	cc, err := dcerpc.Dial(ctx, os.Getenv("SERVER"), dcerpc.WithLogger(zerolog.New(os.Stdout)), epm.EndpointMapper(ctx, os.Getenv("SERVER"), dcerpc.WithLogger(zerolog.New(os.Stdout))))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer cc.Close(ctx)

	// WithSign() is no longer supported after CVE-2020-1472 fix.
	cli, err := logon.NewSecureChannelClient(ctx, cc, dcerpc.WithSeal(), dcerpc.WithEndpoint("ncacn_ip_tcp:"), dcerpc.WithLogger(zerolog.New(os.Stdout)))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	dcs, err := cli.GetDCName(ctx, &logon.GetDCNameRequest{
		ComputerName: sAMCred.Workstation(),
		Flags:        1<<30 /* locate dns names */ | 1<<9, /* locate ips */
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println(j(dcs))

	uCred := credential.NewFromString(uCred)

	ntlmV2 := &ntlm.V2{Config: ntlm.NewConfig()}

	nonce := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}

	resp, err := ntlmV2.ChallengeResponse(ctx, uCred, &ntlm.ChallengeMessage{
		ServerChallenge: nonce,
		TargetInfo: ntlm.AttrValues{
			ntlm.AttrNetBIOSDomainName: &ntlm.Value{NetBIOSDomainName: sAMCred.DomainName()},
			ntlm.AttrTargetName:        &ntlm.Value{TargetName: sAMCred.Workstation()},
		},
	}, nonce)

	samLogon, err := cli.SAMLogon(ctx, &logon.SAMLogonRequest{
		LogonServer:  dcs.DomainControllerInfo.DomainControllerName,
		ComputerName: sAMCred.Workstation(),
		LogonLevel:   logon.LogonInfoClassNetworkTransitiveInformation,
		LogonInformation: &logon.Level{
			Value: &logon.Level_LogonNetworkTransitive{LogonNetworkTransitive: &logon.NetworkInfo{
				Identity: &logon.LogonIdentityInfo{
					ParameterControl: 1 << 11,
					LogonDomainName:  &dtyp.UnicodeString{Buffer: uCred.DomainName()},
					UserName:         &dtyp.UnicodeString{Buffer: uCred.UserName()},
					Workstation:      &dtyp.UnicodeString{Buffer: uCred.Workstation()},
				},
				LMChallenge:         &logon.LMChallenge{Data: nonce},
				LMChallengeResponse: &logon.String{Buffer: resp.LM},
				NTChallengeResponse: &logon.String{Buffer: resp.NT},
			}},
		},
		ValidationLevel: logon.ValidationInfoClassSAMInfo4,
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println(j(samLogon))

	cbKey := samLogon.ValidationInformation.GetValue().(*logon.ValidationSAMInfo4).UserSessionKey

	key := append(cbKey.Data[0].Data, cbKey.Data[1].Data...)
	fmt.Println(" ----------------------- ")
	fmt.Println("SAM_SESSION_KEY:", hex.EncodeToString(key))
	fmt.Println(" ----------------------- ")
	fmt.Println("NTLM_SESSION_BASE_KEY:", hex.EncodeToString(resp.SessionBaseKey))
	fmt.Println(" ----------------------- ")

	domain, err := cli.GetDomainInfo(ctx, &logon.GetDomainInfoRequest{
		ServerName:   dcs.DomainControllerInfo.DomainControllerName,
		ComputerName: sAMCred.Workstation(),
		Level:        1,
		WorkstationBuffer: &logon.WorkstationInformation{Value: &logon.WorkstationInformation_WorkstationInfo{
			WorkstationInfo: &logon.WorkstationInfo{
				DNSHostName: "COMPUTER1.MSAD2.com",
			},
		}},
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println(j(domain))
}
