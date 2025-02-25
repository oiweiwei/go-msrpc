package rfc1964

import (
	"bytes"
	"encoding/binary"
	"errors"
)

var ErrInvalidToken = errors.New("invalid token")

const (
	// MICTokenID is the token ID for MIC.
	MICTokenID = 0x0101
	// WrapTokenID is the token ID for wrapping.
	WrapTokenID = 0x0201
)

const (
	// Signature algorithm DES MAC MD5.
	SignAlgorithmDESMACMD5 = 0x0000
	// Signature algorithm MD2.5
	SignAlgorithmMD2_5 = 0x0100
	// Signature algorithm DES MAC.
	SignAlgorithmDESMAC = 0x0200
	// Seal algorithm none.
	SealAlgorithmNone = 0xffff
	// Seal algorithm DES.
	SealAlgorithmDES = 0x0000
)

// MICToken is a token used for MIC.
type MICToken struct {
	TokenID            uint16
	SignatureAlgorithm uint16
	Filler             [4]byte
	SequenceNumber     []byte
	Checksum           []byte
}

// Equals returns true if the two tokens are equal.
func (tok *MICToken) Equals(other *MICToken) bool {
	return tok.TokenID == other.TokenID &&
		tok.SignatureAlgorithm == other.SignatureAlgorithm &&
		tok.Filler == other.Filler &&
		bytes.Equal(tok.SequenceNumber, other.SequenceNumber) &&
		bytes.Equal(tok.Checksum, other.Checksum)
}

// SetSequenceNumber sets the sequence number of the token.
func (tok *MICToken) SetSequenceNumber(seqNum uint32, isLocal bool) {
	binary.LittleEndian.PutUint32(tok.SequenceNumber, seqNum)
	if !isLocal {
		copy(tok.SequenceNumber[4:], []byte{0xff, 0xff, 0xff, 0xff})
	}
}

// NewMICToken creates a new MICToken.
func NewMICToken() *MICToken {
	tok := &MICToken{
		TokenID:            MICTokenID,
		SignatureAlgorithm: SignAlgorithmDESMACMD5,
		Filler:             [4]byte{0xff, 0xff, 0xff, 0xff},
		SequenceNumber:     make([]byte, 8),
		Checksum:           make([]byte, 8),
	}
	return tok
}

// Size returns the size of the token.
func (tok *MICToken) Size() int {
	return 24
}

// Header returns the header of the token.
func (tok *MICToken) Header() []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint16(b[0:2], tok.TokenID)
	binary.BigEndian.PutUint16(b[2:4], tok.SignatureAlgorithm)
	copy(b[4:8], tok.Filler[:])
	return b
}

// Marshal returns the byte representation of the token.
func (tok *MICToken) Marshal() []byte {
	b := make([]byte, tok.Size())
	binary.BigEndian.PutUint16(b[0:2], tok.TokenID)
	binary.BigEndian.PutUint16(b[2:4], tok.SignatureAlgorithm)
	copy(b[4:8], tok.Filler[:])
	copy(b[8:16], tok.SequenceNumber)
	copy(b[16:24], tok.Checksum)
	return b
}

// Unmarshal parses the byte representation of the token.
func (tok *MICToken) Unmarshal(b []byte) error {
	if len(b) < 24 {
		return ErrInvalidToken
	}
	tok.TokenID = binary.BigEndian.Uint16(b[0:2])
	tok.SignatureAlgorithm = binary.BigEndian.Uint16(b[2:4])
	copy(tok.Filler[:], b[4:8])
	copy(tok.SequenceNumber, b[8:16])
	copy(tok.Checksum, b[16:24])
	return nil
}

// WrapToken is a token used for wrapping.
type WrapToken struct {
	TokenID            uint16
	SignatureAlgorithm uint16
	SealAlgorithm      uint16
	Filler             [2]byte
	SequenceNumber     []byte
	Checksum           []byte
	Confounder         []byte
}

// NewWrapToken creates a new WrapToken.
func NewWrapToken() *WrapToken {
	tok := &WrapToken{
		TokenID:            WrapTokenID,
		SignatureAlgorithm: SignAlgorithmDESMACMD5,
		SealAlgorithm:      SealAlgorithmNone,
		Filler:             [2]byte{0xff, 0xff},
		SequenceNumber:     make([]byte, 8),
		Checksum:           make([]byte, 8),
		Confounder:         make([]byte, 8),
	}
	return tok
}

func (tok *WrapToken) SetEncryption(encrypt bool) {
	if encrypt {
		tok.SealAlgorithm = SealAlgorithmDES
	} else {
		tok.SealAlgorithm = SealAlgorithmNone
	}
}

func (tok *WrapToken) UseEncryption() bool {
	return tok.SealAlgorithm != SealAlgorithmNone
}

// SetSequenceNumber sets the sequence number of the token.
func (tok *WrapToken) SetSequenceNumber(seqNum uint32, isLocal bool) {
	binary.LittleEndian.PutUint32(tok.SequenceNumber, seqNum)
	if !isLocal {
		copy(tok.SequenceNumber[4:], []byte{0xff, 0xff, 0xff, 0xff})
	}
}

// Equals returns true if the two tokens are equal.
func (tok *WrapToken) Equals(other *WrapToken) bool {
	return tok.TokenID == other.TokenID &&
		tok.SignatureAlgorithm == other.SignatureAlgorithm &&
		tok.SealAlgorithm == other.SealAlgorithm &&
		tok.Filler == other.Filler &&
		bytes.Equal(tok.SequenceNumber, other.SequenceNumber) &&
		bytes.Equal(tok.Checksum, other.Checksum) &&
		bytes.Equal(tok.Confounder, other.Confounder)
}

// Size returns the size of the token.
func (tok *WrapToken) Size() int {
	return 32
}

// Header returns the header of the token.
func (tok *WrapToken) Header() []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint16(b[0:2], tok.TokenID)
	binary.BigEndian.PutUint16(b[2:4], tok.SignatureAlgorithm)
	binary.BigEndian.PutUint16(b[4:6], tok.SealAlgorithm)
	copy(b[6:8], tok.Filler[:])
	return b
}

// Marshal returns the byte representation of the token.
func (tok *WrapToken) Marshal() []byte {
	b := make([]byte, tok.Size())
	binary.BigEndian.PutUint16(b[0:2], tok.TokenID)
	binary.BigEndian.PutUint16(b[2:4], tok.SignatureAlgorithm)
	binary.BigEndian.PutUint16(b[4:6], tok.SealAlgorithm)
	copy(b[6:8], tok.Filler[:])
	copy(b[8:16], tok.SequenceNumber)
	copy(b[16:24], tok.Checksum)
	copy(b[24:32], tok.Confounder)
	return b
}

// Unmarshal parses the byte representation of the token.
func (tok *WrapToken) Unmarshal(b []byte) error {
	if len(b) < 32 {
		return ErrInvalidToken
	}
	tok.TokenID = binary.BigEndian.Uint16(b[0:2])
	tok.SignatureAlgorithm = binary.BigEndian.Uint16(b[2:4])
	tok.SealAlgorithm = binary.BigEndian.Uint16(b[4:6])
	copy(tok.Filler[:], b[6:8])
	copy(tok.SequenceNumber, b[8:16])
	copy(tok.Checksum, b[16:24])
	copy(tok.Confounder, b[24:32])
	return nil
}
