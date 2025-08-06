//go:build exclude

// netlogon_sec_channel.go script establishes the secure channel and validates credentials
// provided via creds parameter. (specify creds as --creds 'DOMAIN\username%password').
//
// NOTE: MACHINE_ACCOUNT_USERNAME, and MACHINE_ACCOUNT_PASSWORD are machine account password acquired while joining
// domain. To join domain you must do following:
//
//	# https://manpages.ubuntu.com/manpages/trusty/man8/adcli.8.html
//	$ sudo adcli join CONTOSO.NET --domain-controller=${SERVER} --show-password --show-details
//
// Password will be located under computer-password key, and MACHINE_ACCOUNT_USERNAME must be set to computer-name + '$'.
// (if computer-name is MYPC, then MACHINE_ACCOUNT_USERNAME will be MYPC$)
package main

import (
	"context"
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/oiweiwei/go-msrpc/dcerpc"

	config "github.com/oiweiwei/go-msrpc/config"
	config_flag "github.com/oiweiwei/go-msrpc/config/flag"

	"github.com/oiweiwei/go-msrpc/ssp/credential"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
	"github.com/oiweiwei/go-msrpc/ssp/ntlm"

	"github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	"github.com/oiweiwei/go-msrpc/msrpc/nrpc/logon/v1"

	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/hresult"
	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/ntstatus"
	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/win32"

	. "github.com/oiweiwei/go-msrpc/examples/common"
)

var (
	cfg          = config.New().UseMachineAccount().UseNetlogonSSP()
	uCred        string
	wkstaDNSName string
)

func init() {

	config_flag.BindFlags(cfg, flag.CommandLine)

	flag.StringVar(&uCred, "user-credential", `Administrator%P@ssw0rd`, "user credentials to verify (Domain\\Username%Password")
	flag.StringVar(&wkstaDNSName, "workstation-dns-name", "COMPUTER.MSAD.LOCAL", "workstation dns name")
}

func main() {

	if err := config_flag.ParseAndValidate(cfg, flag.CommandLine); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if len(cfg.MachineAccountCredentials()) == 0 {
		fmt.Fprintln(os.Stderr, "machine account credential is required, falling back to provided")
		cfg.Credential.MachineAccountPassword = cfg.Credential.Password
		cfg.Credential.MachineAccountNTHash = cfg.Credential.NTHash
	}

	// should be only one credential.
	maCred := cfg.MachineAccountCredentials()[0]

	ctx := gssapi.NewSecurityContext(context.Background())

	cc, err := dcerpc.Dial(ctx, cfg.ServerAddr(), cfg.DialOptions(ctx)...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "dial", err)
		return
	}
	defer cc.Close(ctx)

	// WithSign() is no longer supported after CVE-2020-1472 fix.
	cli, err := logon.NewSecureChannelClient(ctx, cc, cfg.ClientOptions(ctx)...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "new_secure_channel_client", err)
		return
	}

	uCred := credential.NewFromString(uCred)

	ntlmV2 := &ntlm.V2{Config: ntlm.NewConfig()}

	nonce := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}

	resp, err := ntlmV2.ChallengeResponse(ctx, uCred, &ntlm.ChallengeMessage{
		ServerChallenge: nonce,
		TargetInfo: ntlm.AttrValues{
			ntlm.AttrNetBIOSDomainName: &ntlm.Value{NetBIOSDomainName: maCred.DomainName()},
			ntlm.AttrTargetName:        &ntlm.Value{TargetName: maCred.Workstation()},
		},
	}, nonce)

	samLogon, err := cli.SAMLogon(ctx, &logon.SAMLogonRequest{
		LogonServer:  cli.DomainControllerInfo().DomainControllerName,
		ComputerName: maCred.Workstation(),
		LogonLevel:   logon.LogonInfoClassNetworkTransitiveInformation,
		LogonInformation: &logon.Level{
			Value: &logon.Level_LogonNetworkTransitive{LogonNetworkTransitive: &logon.NetworkInfo{
				Identity: &logon.LogonIdentityInfo{
					ParameterControl: logon.IdentityAllowWorkstationTrustAccount,
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

	fmt.Println(J(samLogon))

	cbKey := samLogon.ValidationInformation.GetValue().(*logon.ValidationSAMInfo4).UserSessionKey

	key := append(cbKey.Data[0].Data, cbKey.Data[1].Data...)
	fmt.Println(" ----------------------- ")
	fmt.Println("SAM_SESSION_KEY:\t", hex.EncodeToString(key))
	fmt.Println(" ----------------------- ")
	fmt.Println("NTLM_SESSION_BASE_KEY:\t", hex.EncodeToString(resp.SessionBaseKey))
	fmt.Println(" ----------------------- ")

	if hex.EncodeToString(key) != hex.EncodeToString(resp.SessionBaseKey) {
		fmt.Fprintln(os.Stderr, "SAM_SESSION_KEY does not match NTLM_SESSION_BASE_KEY")
		return
	}

	domain, err := cli.GetDomainInfo(ctx, &logon.GetDomainInfoRequest{
		ServerName:   cli.DomainControllerInfo().DomainControllerName,
		ComputerName: maCred.Workstation(),
		Level:        1,
		WorkstationBuffer: &logon.WorkstationInformation{Value: &logon.WorkstationInformation_WorkstationInfo{
			WorkstationInfo: &logon.WorkstationInfo{
				DNSHostName: wkstaDNSName,
			},
		}},
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println(J(domain))
}
