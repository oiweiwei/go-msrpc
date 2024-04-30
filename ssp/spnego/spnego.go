// spnego package implements the SPNEGO security service client.
//
// This package also contains client-side GSSAPI bindings (InitSecurityContext, Wrap, Unwrap and so on).
package spnego

import (
	"context"
	"encoding/asn1"
	"fmt"

	cb "golang.org/x/crypto/cryptobyte"
	cb_asn1 "golang.org/x/crypto/cryptobyte/asn1"
)

const Application = 1 << 6

// The negotiation state.
type State int

var (
	// No further negotiation message from the peer is expected,
	// and the security context is established for the sender.
	AcceptCompleted State = 0
	// At least one additional negotiation message from the peer is
	// needed to establish the security context.
	AcceptIncomplete State = 1
	// The sender terminates the negotiation.
	Reject State = 2
	// The sender indicates that the exchange of MIC tokens will be
	// REQUIRED if per-message integrity services are available on the
	// mechanism context to be established.
	//
	// This value SHALL only be present in the first reply from the target.
	RequestMIC State = 3
)

// The negotiate token initialization message.
type NegTokenInit struct {
	// This field contains one or more security mechanisms
	// available for the initiator, in decreasing preference order
	// (favorite choice first).
	MechTypes []asn1.ObjectIdentifier
	// This field, if present, contains the service options that are
	// requested to establish the context.
	ReqFlags asn1.BitString
	// This field, if present, contains the optimistic mechanism token.
	MechToken []byte
	// This field, if present, contains an MIC token for the mechanism
	// list in the initial negotiation message.
	MechTokenMIC []byte
	// Negotiate hints.
	HintName string
	// Hint address.
	HintAddress []byte
}

const (
	HintName = "not_defined_in_RFC4178@please_ignore"
)

// Marshal function marshals the negotiate token initialization.
func (tok *NegTokenInit) Marshal(ctx context.Context) ([]byte, error) {

	b := cb.NewBuilder(nil)

	// SPNEGO
	b.AddASN1(cb_asn1.Tag(Application).Constructed(), func(b *cb.Builder) {
		// SPNEGO OID
		b.AddASN1ObjectIdentifier(MechanismTypeSPNEGO)
		// [0] NegTokenInit
		b.AddASN1(cb_asn1.Tag(0).ContextSpecific().Constructed(), func(b *cb.Builder) {
			// NegTokenInit ::= SEQUENCE
			b.AddASN1(cb_asn1.SEQUENCE, func(b *cb.Builder) {
				// [0] MechTypeList
				if len(tok.MechTypes) > 0 {
					b.AddASN1(cb_asn1.Tag(0).ContextSpecific().Constructed(), func(b *cb.Builder) {
						// SEQUENCE OF
						b.AddASN1(cb_asn1.SEQUENCE, func(b *cb.Builder) {
							for _, value := range tok.MechTypes {
								// OID
								b.AddASN1ObjectIdentifier(value)
							}
						})
					})
				}
				// [1] ContextFlags OPTIONAL
				if len(tok.ReqFlags.Bytes) > 0 {
					b.AddASN1(cb_asn1.Tag(1).ContextSpecific().Constructed(), func(b *cb.Builder) {
						b.AddASN1BitString(tok.ReqFlags.Bytes)
					})
				}

				// [2] MechToken OCTET STRING OPTIONAL
				if len(tok.MechToken) > 0 {
					b.AddASN1(cb_asn1.Tag(2).ContextSpecific().Constructed(), func(b *cb.Builder) {
						b.AddASN1OctetString(tok.MechToken)
					})
				}

				if tok.HintName != "" || tok.HintAddress != nil {
					// [3] NegHints OPTIONAL
					b.AddASN1(cb_asn1.Tag(3).ContextSpecific().Constructed(), func(b *cb.Builder) {
						b.AddASN1(cb_asn1.SEQUENCE, func(b *cb.Builder) {
							// hintName[0] GeneralString OPTIONAL
							if tok.HintName != "" {
								b.AddASN1(cb_asn1.Tag(0).ContextSpecific().Constructed(), func(b *cb.Builder) {
									b.AddASN1(cb_asn1.GeneralString, func(b *cb.Builder) {
										b.AddBytes([]byte(tok.HintName))
									})
								})
							}
							if tok.HintAddress != nil {
								// hintAddress[1] OCTET STRING OPTIONAL
								b.AddASN1(cb_asn1.Tag(1).ContextSpecific().Constructed(), func(b *cb.Builder) {
									b.AddASN1OctetString(tok.HintAddress)
								})
							}
						})
					})
					// [4] MechTokenMIC OCTET STRING OPTIONAL
					if len(tok.MechTokenMIC) > 0 {
						b.AddASN1(cb_asn1.Tag(4).ContextSpecific().Constructed(), func(b *cb.Builder) {
							b.AddASN1OctetString(tok.MechTokenMIC)
						})
					}
				} else {
					// [3] MechTokenMIC OCTET STRING OPTIONAL
					if len(tok.MechTokenMIC) > 0 {
						b.AddASN1(cb_asn1.Tag(3).ContextSpecific().Constructed(), func(b *cb.Builder) {
							b.AddASN1OctetString(tok.MechTokenMIC)
						})
					}
				}
			})
		})
	})

	bytes, err := b.Bytes()
	if err != nil {
		return nil, fmt.Errorf("neg_token_init: marshal: %w", err)
	}

	return bytes, nil

}

// Unmarshal function unmarshals the negotiate token initialization.
func (tok *NegTokenInit) Unmarshal(ctx context.Context, b []byte) error {

	s := cb.String(b)

	// SPNEGO
	if !s.ReadASN1(&s, cb_asn1.Tag(Application).Constructed()) {
		return fmt.Errorf("neg_token_init: unmarshal: application read error")
	}
	var oid asn1.ObjectIdentifier
	// SPNEGO OID
	if !s.ReadASN1ObjectIdentifier((*asn1.ObjectIdentifier)(&oid)) || !oid.Equal(MechanismTypeSPNEGO) {
		return fmt.Errorf("neg_token_init: unmarshal: header read error")
	}
	// [0] NegTokenInit
	if !s.ReadASN1(&s, cb_asn1.Tag(0).ContextSpecific().Constructed()) {
		return fmt.Errorf("neg_token_init: unmarshal: message error")
	}
	// NegTokenInit ::= SEQUENCE
	if !s.ReadASN1(&s, cb_asn1.SEQUENCE) {
		return fmt.Errorf("neg_token_resp: unmarshal: sequence error")
	}

	tag, present := cb.String{}, false

	// [0] MechTypeList
	if s.ReadOptionalASN1(&tag, &present, cb_asn1.Tag(0).ContextSpecific().Constructed()) {
		if present {
			// SEQUENCE OF
			if !tag.ReadASN1(&tag, cb_asn1.SEQUENCE) {
				return fmt.Errorf("neg_token_init: unmarshal: mech_types read error")
			}
			// OID
			for len(tag) > 0 {
				var oid asn1.ObjectIdentifier
				if tag.ReadASN1ObjectIdentifier((*asn1.ObjectIdentifier)(&oid)) {
					tok.MechTypes = append(tok.MechTypes, oid)
				} else {
					return fmt.Errorf("neg_token_init: unmarshal: mech_types read error")
				}
			}
		}
	} else {
		return fmt.Errorf("neg_token_init: unmarshal: mech_types read error")
	}
	// [1] ContextFlags OPTIONAL
	if s.ReadOptionalASN1(&tag, &present, cb_asn1.Tag(1).ContextSpecific().Constructed()) {
		if present {
			if !tag.ReadASN1BitString(&tok.ReqFlags) {
				return fmt.Errorf("neg_token_init: unmarshal: req_flags read error")
			}
		}
	} else {
		return fmt.Errorf("neg_token_init: unmarshal: req_flags read error")
	}
	// [2] MechToken OCTET STRING OPTIONAL
	if !s.ReadOptionalASN1OctetString(&tok.MechToken, nil, cb_asn1.Tag(2).ContextSpecific().Constructed()) {
		return fmt.Errorf("neg_token_init: unmarshal: mech_token read error")
	}

	if s.ReadOptionalASN1(&tag, &present, cb_asn1.Tag(3).ContextSpecific().Constructed()) {
		out, hint := cb_asn1.Tag(0), cb.String{}
		if present {
			if tag.ReadAnyASN1(&tag, &out) {
				switch out {
				case cb_asn1.SEQUENCE:
					// [3] negHints NegHints OPTIONAL
					{
						// hintName[0] GeneralString OPTIONAL
						if tag.ReadOptionalASN1(&hint, &present, cb_asn1.Tag(0).ContextSpecific().Constructed()) {
							if present {
								if hint.ReadASN1(&hint, cb_asn1.GeneralString) {
									tok.HintName = string(hint)
								} else {
									return fmt.Errorf("neg_token_init2: unmarshal: hint_name read error")
								}
							}
						} else {
							return fmt.Errorf("neg_token_init2: unmarshal: hint_name read error")
						}
						//  hintAddress[1] OCTET STRING OPTIONAL
						if !tag.ReadOptionalASN1OctetString(&tok.HintAddress, nil, cb_asn1.Tag(1).ContextSpecific().Constructed()) {
							return fmt.Errorf("neg_token_init2: unmarshal: hint_address read error")
						}
					}

					// [4] MechTokenMIC OCTET STRING OPTIONAL
					if !s.ReadOptionalASN1OctetString(&tok.MechTokenMIC, nil, cb_asn1.Tag(4).ContextSpecific().Constructed()) {
						return fmt.Errorf("neg_token_init: unmarshal: mech_token_mic read error")
					}

				case cb_asn1.OCTET_STRING:
					// [3] MechTokenMIC OCTET STRING OPTIONAL
					tok.MechTokenMIC = []byte(tag)
				default:
					return fmt.Errorf("neg_token_init: unmarshal: mech_token_mic unknown tag %d", out)
				}
			} else {
				return fmt.Errorf("neg_token_init: unmarshal: mech_token_mic read error")
			}
		}
	} else {
		return fmt.Errorf("neg_token_init: unmarshal: mech_token_mic read error")
	}

	return nil
}

// The negotiate token response.
type NegTokenResp struct {
	// This field, if present, contains the state of the negotiation.
	State State
	// This field SHALL only be present in the first reply from the
	// target. It MUST be one of the mechanism(s) offered by the initiator.
	SupportedMech asn1.ObjectIdentifier
	// This field, if present, contains tokens specific to the mechanism
	// selected.
	ResponseToken []byte
	// This field, if present, contains an MIC token for the mechanism
	// list in the initial negotiation message.
	MechListMIC []byte
}

// Marshal function marshals the negotiate token response.
func (tok *NegTokenResp) Marshal(ctx context.Context) ([]byte, error) {

	b := cb.NewBuilder(nil)

	// [1] NegTokenResp
	b.AddASN1(cb_asn1.Tag(1).ContextSpecific().Constructed(), func(b *cb.Builder) {
		// NegTokenResp ::= SEQUENCE
		b.AddASN1(cb_asn1.SEQUENCE, func(b *cb.Builder) {
			// [0] ENUMERATED [...] OPTIONAL
			if tok.State != 0 {
				b.AddASN1(cb_asn1.Tag(0).ContextSpecific().Constructed(), func(b *cb.Builder) {
					b.AddASN1Enum(int64(tok.State))
				})
			}
			// [1] MechType OPTIONAL
			if tok.SupportedMech != nil {
				b.AddASN1(cb_asn1.Tag(1).ContextSpecific().Constructed(), func(b *cb.Builder) {
					b.AddASN1(cb_asn1.SEQUENCE, func(b *cb.Builder) {
						b.AddASN1ObjectIdentifier(tok.SupportedMech)
					})
				})
			}
			// [2] OCTET STRING OPTIONAL
			if len(tok.ResponseToken) > 0 {
				b.AddASN1(cb_asn1.Tag(2).ContextSpecific().Constructed(), func(b *cb.Builder) {
					b.AddASN1OctetString(tok.ResponseToken)
				})
			}
			// [3] OCTET STRING OPTIONAL
			if len(tok.MechListMIC) > 0 {
				b.AddASN1(cb_asn1.Tag(3).ContextSpecific().Constructed(), func(b *cb.Builder) {
					b.AddASN1OctetString(tok.MechListMIC)
				})
			}
		})
	})

	return b.Bytes()
}

// Unmarshal function unmarshals the negotiate token response.
func (tok *NegTokenResp) Unmarshal(ctx context.Context, b []byte) error {

	s := cb.String(b)

	// [1] NegTokenResp
	if !s.ReadASN1(&s, cb_asn1.Tag(1).ContextSpecific().Constructed()) {
		return fmt.Errorf("neg_token_resp: unmarshal: message error")
	}

	// NegTokenResp ::= SEQUENCE
	if !s.ReadASN1(&s, cb_asn1.SEQUENCE) {
		return fmt.Errorf("neg_token_resp: unmarshal: sequence error")
	}

	tag, present := cb.String{}, false
	// [0] ENUMERATED [...] OPTIONAL
	if s.ReadOptionalASN1(&tag, &present, cb_asn1.Tag(0).ContextSpecific().Constructed()) {
		if present {
			tag.ReadASN1Enum((*int)(&tok.State))
		}
	} else {
		return fmt.Errorf("neg_token_resp: unmarshal: state read error")
	}
	// [1] MechType OPTIONAL
	if s.ReadOptionalASN1(&tag, &present, cb_asn1.Tag(1).ContextSpecific().Constructed()) {
		if present {
			if tag.ReadOptionalASN1(&tag, &present, cb_asn1.SEQUENCE) {
				if present {
					tag.ReadASN1ObjectIdentifier((*asn1.ObjectIdentifier)(&tok.SupportedMech))
				}
			} else {
				return fmt.Errorf("neg_token_resp: unmarshal: mech_type sequence read error")
			}
		}
	} else {
		return fmt.Errorf("neg_token_resp: unmarshal: mech_type read error")
	}
	// [2] OCTET STRING OPTIONAL
	if !s.ReadOptionalASN1OctetString(&tok.ResponseToken, nil, cb_asn1.Tag(2).ContextSpecific().Constructed()) {
		return fmt.Errorf("neg_token_resp: unmarshal: response_token read error")
	}
	// [3] OCTET STRING OPTIONAL
	if !s.ReadOptionalASN1OctetString(&tok.MechListMIC, nil, cb_asn1.Tag(3).ContextSpecific().Constructed()) {
		return fmt.Errorf("neg_token_resp: unmarshal: mech_list_mic read error")
	}

	return nil
}
