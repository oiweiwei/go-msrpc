//go:generate go run ./gen/main.go -o v1_secure_channel_wrappers.zz.go
package logon

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"time"

	"github.com/oiweiwei/go-msrpc/dcerpc"

	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
	"github.com/oiweiwei/go-msrpc/ssp/netlogon"
)

type LogonSecureChannelClient interface {
	LogonClient
	Encrypt(context.Context, []byte) ([]byte, error)
}

type xxx_SecureChannelClient struct {
	LogonClient
	sCred *netlogon.SecureCredential
}

var SecureChannel_T = &xxx_SecureChannelClient{}

func NewSecureChannelClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (LogonSecureChannelClient, error) {

	cli, err := NewLogonClient(ctx, cc, opts...)
	if err != nil {
		return nil, err
	}

	sspCred := gssapi.GetCredential(cli.Conn().Context(), "", nil, gssapi.InitiateAndAccept)
	if sspCred == nil {
		return nil, fmt.Errorf("secure_channel: credentials missing")
	}

	creds, ok := sspCred.Value().(netlogon.Credential)
	if !ok || creds == nil {
		return nil, fmt.Errorf("secure_channel: credentials missing")
	}

	cfg := &netlogon.Config{
		Capabilities:    netlogon.CapAES_SHA2 | netlogon.CapStrongKey | netlogon.CapSecureRPC | netlogon.CapRC4,
		Credential:      creds,
		ClientChallenge: make([]byte, 8),
	}

	if _, err := rand.Read(cfg.ClientChallenge); err != nil {
		return nil, fmt.Errorf("secure_channel: %v", err)
	}

	if creds.Workstation() == "" {
		return nil, fmt.Errorf("secure_channel: workstation missing")
	}

	dc, err := cli.GetDCName(ctx, &GetDCNameRequest{
		ComputerName: creds.Workstation(),
		Flags:        1<<30 /* locate dns names */ | 1<<9, /* locate ips */
	})
	if err != nil {
		return nil, fmt.Errorf("secure_channel: dc_name: %v", err)
	}

	chal, err := cli.RequestChallenge(ctx, &RequestChallengeRequest{
		PrimaryName:     dc.DomainControllerInfo.DomainControllerName,
		ComputerName:    creds.Workstation(),
		ClientChallenge: &Credential{Data: cfg.ClientChallenge},
	})
	if err != nil {
		return nil, fmt.Errorf("secure_channel: request_challenge: %v", err)
	}

	cfg.ServerChallenge = chal.ServerChallenge.Data

	sCred, err := netlogon.NewSecureCredential(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("secure_channel: %v", err)
	}

	clientCred, err := sCred.Encrypt(ctx, cfg.ClientChallenge)
	if err != nil {
		return nil, fmt.Errorf("secure_channel: client_credentials: %v", err)
	}

	auth3, err := cli.Authenticate3(ctx, &Authenticate3Request{
		PrimaryName:       dc.DomainControllerInfo.DomainControllerName,
		AccountName:       creds.UserName(),
		SecureChannelType: SecureChannelTypeWorkstationSecureChannel,
		ComputerName:      creds.Workstation(),
		ClientCredential:  &Credential{Data: clientCred},
		NegotiateFlags:    uint32(cfg.Capabilities),
	})
	if err != nil {
		return nil, fmt.Errorf("secure_channel: auth3: %v", err)
	}

	expServerCred, err := sCred.Encrypt(ctx, cfg.ServerChallenge)
	if err != nil {
		return nil, fmt.Errorf("secure_channel: auth3: server_credentials: %v", err)
	}

	if !bytes.Equal(expServerCred, auth3.ServerCredential.Data) {
		return nil, fmt.Errorf("secure_channel: auth3: invalid server credentials")
	}

	// upgrade to secure channel.
	if err := cli.AlterContext(ctx, append(opts, dcerpc.WithSecurityConfig(cfg))...); err != nil {
		return nil, fmt.Errorf("secure_channel: %v", err)
	}

	return &xxx_SecureChannelClient{
		LogonClient: cli,
		sCred:       sCred,
	}, nil
}

func (o *xxx_SecureChannelClient) Encrypt(ctx context.Context, b []byte) ([]byte, error) {
	return o.sCred.Encrypt(ctx, b)
}

func (o *xxx_SecureChannelClient) VerifyAuthenticator(ctx context.Context, ra *Authenticator) error {
	return o.sCred.Verify(ctx, 1, ra.Credential.Data)
}

func (o *xxx_SecureChannelClient) SetAuthenticators(ctx context.Context, a, ra **Authenticator) error {

	var err error

	*a = &Authenticator{
		Timestamp:  uint32(time.Now().Unix()),
		Credential: &Credential{},
	}

	if (*a).Credential.Data, err = o.sCred.Next(ctx, (*a).Timestamp); err != nil {
		return err
	}

	*ra = &Authenticator{}

	return nil
}
