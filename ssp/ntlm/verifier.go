package ntlm

import (
	"context"
	"fmt"

	"github.com/oiweiwei/go-msrpc/ssp/credential"
)

type Verifier interface {
	VerifyChallenge(context.Context, *SecurityParameters, Credential, []byte, *AuthenticateMessage) (*ChallengeResponse, error)
}

type LocalVerifier struct {
	Config   *Config
	Database credential.Database
}

func (v *LocalVerifier) Lookup(ctx context.Context, cred Credential) (Credential, bool) {
	if credential.IsAnonymous(cred) && v.Config.AllowAnonymous {
		return credential.Anonymous(), true
	}
	return v.Database.Lookup(ctx, cred)
}

func (v *LocalVerifier) VerifyChallenge(ctx context.Context, session *SecurityParameters, cred Credential, nonce []byte, am *AuthenticateMessage) (*ChallengeResponse, error) {

	var (
		ntlm NTLMVersion
	)

	ntlm = &V1{Config: v.Config, SecurityParameters: session}

	if len(am.NTChallengeResponse) > 24 {
		ntlm = &V2{Config: v.Config, SecurityParameters: session}
	}

	cred, ok := v.Lookup(ctx, cred)
	if !ok {
		return nil, fmt.Errorf("ntlm: local verifyer: credential not found for %s", cred.DomainName()+"\\"+cred.UserName())
	}

	resp, err := ntlm.AuthenticateResponse(ctx, cred, am, nonce)
	if err != nil {
		return nil, fmt.Errorf("ntlm: local provider: authenticate verification failed: %w", err)
	}

	resp.KeyExchangeKey, err = ntlm.KeyExchangeKey(ctx, &ChallengeMessage{ServerChallenge: nonce}, resp)
	if err != nil {
		return nil, fmt.Errorf("ntlm: local provider: key exchange key failed: %w", err)
	}

	return resp, nil
}
