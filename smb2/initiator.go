package smb2

import (
	"context"
	"encoding/asn1"

	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
)

func SecurtiyContextInitiator(opts ...gssapi.ContextOption) *Initiator {

	i := Initiator{ctx: gssapi.NewSecurityContext(context.Background(), opts...)}

	for _, opt := range opts {
		if opt, ok := (any)(opt).(gssapi.Option); ok {
			i.opts = append(i.opts, opt)
		}
	}

	// allways include DCEStyle.
	i.opts = append(i.opts, gssapi.WithRequest(gssapi.DCEStyle))

	return &i
}

type Initiator struct {
	ctx  context.Context
	opts []gssapi.Option
}

func (i *Initiator) InitSecContext() ([]byte, error) {

	i.ctx = gssapi.ResetSecurityContext(i.ctx)

	tok, err := gssapi.InitSecurityContext(i.ctx, &gssapi.Token{}, i.opts...)
	if err != nil {
		return nil, err
	}

	return tok.Payload, nil
}

func (i *Initiator) AcceptSecContext(b []byte) ([]byte, error) {

	tok, err := gssapi.InitSecurityContext(i.ctx, &gssapi.Token{Payload: b}, i.opts...)
	if err != nil {
		return nil, err
	}

	return tok.Payload, nil
}

func (i *Initiator) OID() asn1.ObjectIdentifier {

	for _, mechanism := range gssapi.ListMechanisms(i.ctx) {
		return (asn1.ObjectIdentifier)(mechanism.Type())
	}

	return nil
}

func (i *Initiator) Sum(b []byte) []byte {

	tok, err := gssapi.MakeSignature(i.ctx, &gssapi.MessageToken{Payload: b}, i.opts...)
	if err != nil {
		return nil
	}

	return tok.Signature
}

func (i *Initiator) SessionKey() []byte {

	key, ok := gssapi.GetAttribute(i.ctx, gssapi.AttributeSessionKey, i.opts...)
	if !ok {
		return nil
	}

	return key.([]byte)
}
