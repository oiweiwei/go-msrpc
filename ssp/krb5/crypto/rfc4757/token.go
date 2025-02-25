package rfc4757

import (
	"encoding/binary"

	"github.com/oiweiwei/go-msrpc/ssp/krb5/crypto/rfc1964"
)

const (
	// HMAC signature algorithm.
	SignAlgorithmHMAC = 0x1100
	// RC4 seal algorithm.
	SealAlgorithmRC4 = 0x1000
)

// MICToken is a token used for MIC.
type MICToken rfc1964.MICToken

// Equals returns true if the two tokens are equal.
func (tok *MICToken) Equals(other *MICToken) bool {
	return (*rfc1964.MICToken)(tok).Equals((*rfc1964.MICToken)(other))
}

// SetSequenceNumber sets the sequence number of the token.
func (tok *MICToken) SetSequenceNumber(seqNum uint32, isLocal bool) *MICToken {
	binary.BigEndian.PutUint32(tok.SequenceNumber, seqNum)
	if !isLocal {
		copy(tok.SequenceNumber[4:], []byte{0xff, 0xff, 0xff, 0xff})
	}
	return tok
}

// NewMICToken creates a new MICToken.
func NewMICToken() *MICToken {
	tok := (*MICToken)(rfc1964.NewMICToken())
	tok.SignatureAlgorithm = SignAlgorithmHMAC
	return tok
}

// Size returns the size of the token.
func (tok *MICToken) Size() int { return (*rfc1964.MICToken)(tok).Size() }

// Header returns the header of the token.
func (tok *MICToken) Header() []byte { return (*rfc1964.MICToken)(tok).Header() }

// Marshal returns the byte representation of the token.
func (tok *MICToken) Marshal() []byte { return (*rfc1964.MICToken)(tok).Marshal() }

// Unmarshal parses the byte representation of the token.
func (tok *MICToken) Unmarshal(b []byte) error { return (*rfc1964.MICToken)(tok).Unmarshal(b) }

// WrapToken is a token used for wrapping.
type WrapToken rfc1964.WrapToken

// Equals returns true if the two tokens are equal.
func (tok *WrapToken) Equals(other *WrapToken) bool {
	return (*rfc1964.WrapToken)(tok).Equals((*rfc1964.WrapToken)(other))
}

// SetSequenceNumber sets the sequence number of the token.
func (tok *WrapToken) SetSequenceNumber(seqNum uint32, isLocal bool) *WrapToken {
	binary.BigEndian.PutUint32(tok.SequenceNumber, seqNum)
	if !isLocal {
		copy(tok.SequenceNumber[4:], []byte{0xff, 0xff, 0xff, 0xff})
	}
	return tok
}

func (tok *WrapToken) SetEncryption(encrypt bool) *WrapToken {
	if encrypt {
		tok.SealAlgorithm = SealAlgorithmRC4
	} else {
		tok.SealAlgorithm = rfc1964.SealAlgorithmNone
	}
	return tok
}

func (tok *WrapToken) UseEncryption() bool {
	return tok.SealAlgorithm != rfc1964.SealAlgorithmNone
}

// NewWrapToken creates a new WrapToken.
func NewWrapToken() *WrapToken {
	tok := (*WrapToken)(rfc1964.NewWrapToken())
	tok.SignatureAlgorithm = SignAlgorithmHMAC
	return tok
}

// Size returns the size of the token.
func (tok *WrapToken) Size() int { return (*rfc1964.WrapToken)(tok).Size() }

// Header returns the header of the token.
func (tok *WrapToken) Header() []byte { return (*rfc1964.WrapToken)(tok).Header() }

// Marshal returns the byte representation of the token.
func (tok *WrapToken) Marshal() []byte { return (*rfc1964.WrapToken)(tok).Marshal() }

// Unmarshal parses the byte representation of the token.
func (tok *WrapToken) Unmarshal(b []byte) error { return (*rfc1964.WrapToken)(tok).Unmarshal(b) }
