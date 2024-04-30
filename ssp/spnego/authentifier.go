package spnego

import (
	"context"
	"encoding/asn1"
	"fmt"

	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
)

type Config struct {
	// The services available.
	Capabilities gssapi.Cap
	// The list of negotiated mechanisms.
	MechanismsList []gssapi.MechanismFactory
}

type Authentifier struct {
	// The authentifier configuration.
	*Config
	// The selected mechanism.
	Mechanism gssapi.Mechanism
	// The retrieved Mechanism List.
	RetrievedMechanismList []asn1.ObjectIdentifier
	// Require mechanism list MIC.
	RequireMechanismListMIC bool
}

func (a *Authentifier) Negotiate(ctx context.Context) ([]byte, error) {

	var err error

	a.Mechanism, err = a.MechanismsList[0].New(ctx)
	if err != nil {
		return nil, fmt.Errorf("spnego: init: mechanism new: %w", err)
	}

	// initiate payload.

	mechTok, err := a.Mechanism.Init(ctx, &gssapi.Token{})
	if err != nil {
		return nil, fmt.Errorf("spnego: init: mechanism: %w", err)
	}

	neg := &NegTokenInit{
		MechTypes: a.MakeMechanismList(ctx),
		MechToken: mechTok.Payload,
	}

	b, err := neg.Marshal(ctx)
	if err != nil {
		return nil, fmt.Errorf("spnego: init: marshal neg token init: %w", err)
	}

	return b, nil
}

func (a *Authentifier) Respond(ctx context.Context, b []byte) ([]byte, error) {

	var err error

	resp := &NegTokenResp{}

	if a.IsNegTokenInit(ctx, b) {

		init := &NegTokenInit{}

		if err := init.Unmarshal(ctx, b); err != nil {
			return nil, fmt.Errorf("spnego: init: unmarshal neg token init: %w", err)
		}

		a.RetrievedMechanismList = init.MechTypes

	loop:
		// select the mechanism from the retrieved list.
		for i := range a.MechanismsList {
			for _, rm := range a.RetrievedMechanismList {
				if a.MechanismsList[i].Type().Equal((gssapi.OID)(rm)) {
					if a.Mechanism, err = a.MechanismsList[i].New(ctx); err != nil {
						continue
					}
					break loop
				}
			}
		}

		if a.Mechanism == nil {
			return nil, gssapi.ErrUnavailable
		}

	} else {

		if err := resp.Unmarshal(ctx, b); err != nil {
			return nil, fmt.Errorf("spnego: init: unmarshal response: %w", err)
		}

		if resp.State == Reject {
			return nil, fmt.Errorf("spnego: init: %w", ErrReject)
		}

		if resp.State == AcceptCompleted {

			// the spnego negotiation completed successfully, verify
			// the mechanism list mic.

			b, err := asn1.Marshal(a.MakeMechanismList(ctx))
			if err != nil {
				return nil, fmt.Errorf("spnego: init: marshal mech list: %w", err)
			}

			if len(resp.MechListMIC) > 0 || a.RequireMechanismListMIC {
				err = a.Mechanism.VerifySignature(ctx, &gssapi.MessageToken{Payload: b, Signature: resp.MechListMIC})
				if err != nil {
					return nil, fmt.Errorf("spnego: init: verify mech list mic: %w", err)
				}
			}

			// reset the security service.
			// (note that sequence number must not be resetted, only cipher sequence).

			rst, ok := (any)(a.Mechanism).(interface{ ResetSecurityService(context.Context) error })
			if ok {
				if err := rst.ResetSecurityService(ctx); err != nil {
					return nil, fmt.Errorf("spnego: init: reset security service: %w", err)
				}
			}

			return nil, nil
		}
	}

	mechTok, err := a.Mechanism.Init(ctx, &gssapi.Token{
		Payload: resp.ResponseToken,
	})
	if err != nil {
		return nil, err
	}

	resp = &NegTokenResp{
		ResponseToken: mechTok.Payload,
	}

	if gssapi.IsComplete(ctx) {

		// create mechanism list mic.

		b, err := asn1.Marshal(a.MakeMechanismList(ctx))
		if err != nil {
			return nil, fmt.Errorf("spnego: init: marshal mech list: %w", err)
		}

		sgn, err := a.Mechanism.MakeSignature(ctx, &gssapi.MessageToken{Payload: b})
		if err != nil {
			return nil, fmt.Errorf("spnego: init: make mech list signature: %w", err)
		}

		resp.MechListMIC = sgn.Signature

		rst, ok := interface{}(a.Mechanism).(interface{ ResetSecurityService(context.Context) error })
		if ok {
			if err := rst.ResetSecurityService(ctx); err != nil {
				return nil, fmt.Errorf("spnego: init: reset security service: %w", err)
			}
		}
	}

	b, err = resp.Marshal(ctx)
	if err != nil {
		return nil, fmt.Errorf("spnego: init: marshal neg token resp: %w", err)
	}

	return b, nil
}

func (a *Authentifier) IsNegTokenInit(ctx context.Context, b []byte) bool {
	return len(b) > 0 && b[0] == Application|0
}

func (a *Authentifier) MakeMechanismList(ctx context.Context) []asn1.ObjectIdentifier {

	if len(a.RetrievedMechanismList) != 0 {
		return a.RetrievedMechanismList
	}

	mechTypes := make([]asn1.ObjectIdentifier, len(a.Config.MechanismsList))

	for i, mech := range a.Config.MechanismsList {
		mechTypes[i] = (asn1.ObjectIdentifier)(mech.Type())
	}

	return mechTypes
}

func (a *Authentifier) SelectMechanism(ctx context.Context, oid gssapi.OID) gssapi.Mechanism {

	// select mechanism based on oid or first entry if default...
	return nil

}
