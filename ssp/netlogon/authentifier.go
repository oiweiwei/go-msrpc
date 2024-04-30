package netlogon

import (
	"context"
	"fmt"
	"strings"

	"github.com/oiweiwei/go-msrpc/ssp/netlogon/crypto"
)

type Authentifier struct {
	Config *Config
	state  *SecurityService
}

type SecurityService struct {
	Key                    []byte
	InboundSequenceNumber  uint64
	OutboundSequenceNumber uint64
	OutboundCipher         crypto.Cipher
	InboundCipher          crypto.Cipher
}

func (a *Authentifier) SessionKey() []byte {
	if a.state != nil {
		return a.state.Key
	}
	return nil
}

func (a *Authentifier) AuthMessageInit(ctx context.Context) ([]byte, error) {

	m := &AuthMessage{
		MessageType: MessageTypeRequest,
	}

	cred := a.Config.Credential
	if cred == nil {
		return nil, fmt.Errorf("netlogon: init: auth_message_init: credentials required")
	}

	if dn := cred.DomainName(); dn != "" {
		if strings.Contains(dn, ".") {
			m.DNSDomainName = dn
		} else {
			m.NetBIOSDomainName = dn
		}
	}

	if cn := cred.Workstation(); cn != "" {
		if strings.Contains(cn, ".") {
			m.DNSHostName = cn
		} else {
			m.NetBIOSComputerName = cn
		}
	}

	b, err := m.Marshal()
	if err != nil {
		return nil, fmt.Errorf("netlogon: init: auth_message_init: %w", err)
	}

	return b, nil
}

func (a *Authentifier) AuthMessageReply(ctx context.Context, b []byte) ([]byte, error) {

	if len(b) != 12 {
		return nil, fmt.Errorf("netlogon: init: auth_message_reply: token length mismatch")
	}

	m := &AuthMessage{}

	if err := m.Unmarshal(b); err != nil {
		return nil, fmt.Errorf("netlogon: init: auth_message_reply: %w", err)
	}

	if m.MessageType != MessageTypeResponse || m.Flags != NegFlag(0) {
		return nil, fmt.Errorf("netlogon: init: auth_message_reply: invalid token type or flag")
	}

	if err := a.makeSecurityService(ctx); err != nil {
		return nil, fmt.Errorf("netlogon: init: auth_message_reply: make_security_service: %w", err)
	}

	return nil, nil
}

func (a *Authentifier) makeSecurityService(ctx context.Context) error {

	a.state = &SecurityService{}

	key, err := ComputeSessionKey(ctx, a.Config.Capabilities, a.Config.Credential, a.Config.ClientChallenge, a.Config.ServerChallenge)
	if err != nil {
		return err
	}

	a.state.Key = key

	etype := crypto.EtypeAESSHA2
	if !a.Config.Capabilities.IsSet(CapAES_SHA2) {
		etype = crypto.EtypeRC4MD5
	}

	clientCipher, serverCipher := crypto.NewCipher(ctx, etype, a.state.Key, true), crypto.NewCipher(ctx, etype, a.state.Key, false)

	if a.Config.IsServer {
		a.state.InboundCipher, a.state.OutboundCipher = clientCipher, serverCipher
		a.state.OutboundSequenceNumber = 1
	} else {
		a.state.InboundCipher, a.state.OutboundCipher = serverCipher, clientCipher
		a.state.InboundSequenceNumber = 1
	}

	return nil
}

func (a *Authentifier) MakeOutboundSignature(ctx context.Context, forSign [][]byte) ([]byte, error) {
	sgn, err := a.state.OutboundCipher.MakeSignature(ctx, a.state.OutboundSequenceNumber, forSign)
	if err != nil {
		return nil, fmt.Errorf("netlogon: make outbound signature: %w", err)
	}
	a.state.OutboundSequenceNumber += 2
	return sgn, nil
}

func (a *Authentifier) MakeInboundSignature(ctx context.Context, forSign [][]byte) ([]byte, error) {
	sgn, err := a.state.InboundCipher.MakeSignature(ctx, a.state.InboundSequenceNumber, forSign)
	if err != nil {
		return nil, fmt.Errorf("netlogon: make inbound signature: %w", err)
	}
	a.state.InboundSequenceNumber += 2
	return sgn, nil
}

func (a *Authentifier) WrapOutboundPayload(ctx context.Context, forSign, forSeal [][]byte) ([]byte, error) {
	sgn, err := a.state.OutboundCipher.Wrap(ctx, a.state.OutboundSequenceNumber, forSign, forSeal)
	if err != nil {
		return nil, fmt.Errorf("netlogon: wrap outbound payload: %w", err)
	}

	a.state.OutboundSequenceNumber += 2
	return sgn, nil
}

func (a *Authentifier) UnwrapInboundPayload(ctx context.Context, forSign, forSeal [][]byte, sgn []byte) ([]byte, error) {
	expSgn, err := a.state.InboundCipher.Unwrap(ctx, a.state.InboundSequenceNumber, forSign, forSeal, sgn)
	if err != nil {
		return nil, fmt.Errorf("netlogon: unwrap inbound payload: %w", err)
	}
	a.state.InboundSequenceNumber += 2
	return expSgn, nil
}

func (a *Authentifier) OutboundSignatureSize(ctx context.Context, conf bool) int {
	// use outbound cipher as we wrap the outbound data.
	return a.state.OutboundCipher.Size(ctx, conf)
}
