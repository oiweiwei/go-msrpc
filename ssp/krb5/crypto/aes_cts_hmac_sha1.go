package crypto

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"hash"

	"github.com/oiweiwei/gokrb5.fork/v9/crypto/common"
	"github.com/oiweiwei/gokrb5.fork/v9/crypto/etype"
	"github.com/oiweiwei/gokrb5.fork/v9/iana/keyusage"
)

var (
	CFXFlagSendByAcceptor = uint8(0x01)
	CFXFlagSealed         = uint8(0x02)
	CFXFlagAcceptorSubKey = uint8(0x04)
)

// AESCTSHMACSHA1 represents the integrity/confidentiality routines
// for the aes128-cts-hmac-sha1 and aes256-cts-hmac-sha1.
type AESCTSHMACSHA1 struct {
	setting CipherSetting
	// usages.
	ckU, ecU int
}

func NewAESCipher(ctx context.Context, setting CipherSetting) (Cipher, error) {

	ckU, ecU := keyusage.GSSAPI_INITIATOR_SIGN, keyusage.GSSAPI_INITIATOR_SEAL
	if !setting.IsLocal {
		ckU, ecU = keyusage.GSSAPI_ACCEPTOR_SIGN, keyusage.GSSAPI_ACCEPTOR_SEAL
	}

	return &AESCTSHMACSHA1{setting: setting, ckU: ckU, ecU: ecU}, nil
}

func (c *AESCTSHMACSHA1) ChecksumHash() (hash.Hash, error) {
	return c.newHash(common.GetUsageKc(uint32(c.ckU)))
}

func (c *AESCTSHMACSHA1) IntegrityHash() (hash.Hash, error) {
	return c.newHash(common.GetUsageKi(uint32(c.ecU)))
}

func (c *AESCTSHMACSHA1) newHash(usage []byte) (hash.Hash, error) {
	k, err := c.setting.Type.DeriveKey(c.setting.Key.KeyValue, usage)
	if err != nil {
		return nil, fmt.Errorf("unable to derive key for checksum: %w", err)
	}
	return &SizedHash{Hash: hmac.New(c.setting.Type.GetHashFunc(), k), etype: c.setting.Type}, nil
}

func (c *AESCTSHMACSHA1) flags() uint8 {
	flags := uint8(0)
	if !c.setting.IsLocal {
		flags |= CFXFlagSendByAcceptor
	}
	if c.setting.IsSubKey {
		flags |= CFXFlagAcceptorSubKey
	}
	return flags
}

const (
	EC  = 16
	RRC = 28
)

func (c *AESCTSHMACSHA1) Wrap(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte) ([]byte, error) {
	b, err := c.wrap(ctx, seqNum, forSign, forSeal)
	if err != nil {
		return nil, fmt.Errorf("aes-cts-hmac-sha1: wrap: %w", err)
	}
	return b, nil
}

func (c *AESCTSHMACSHA1) wrap(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte) ([]byte, error) {

	eB, hdr, cc := bytes.NewBuffer(nil), c.WrapHeader(ctx, seqNum), c.setting.Type.GetConfounderByteSize()

	// gen confounder.
	confounder := make([]byte, cc)
	if _, err := rand.Read(confounder); err != nil {
		return nil, fmt.Errorf("read confounder: %w", err)
	}

	// gen ec.
	ec := bytes.Repeat([]byte{0xFF}, EC)
	// set ec value (16). (pad = 1, block_size = 16).
	binary.BigEndian.PutUint16(hdr[4:6], EC)

	// write confounder.
	eB.Write(confounder)

	for i := range forSeal {
		// write buffer.
		eB.Write(forSeal[i])
	}

	// write ec.
	eB.Write(ec)
	// write header.
	eB.Write(hdr)

	iH, err := c.IntegrityHash()
	if err != nil {
		return nil, fmt.Errorf("make integrity hash: %w", err)
	}

	// write confounder.
	iH.Write(confounder)

	for i := range forSign {
		// write buffer.
		iH.Write(forSign[i])
	}

	// write ec.
	iH.Write(ec)
	// write header.
	iH.Write(hdr)

	key, err := c.setting.Type.DeriveKey(c.setting.Key.KeyValue, common.GetUsageKe(uint32(c.ecU)))
	if err != nil {
		return nil, fmt.Errorf("derive key: %w", err)
	}

	_, b, err := c.setting.Type.EncryptData(key, eB.Bytes())
	if err != nil {
		return nil, fmt.Errorf("encrypt data: %w", err)
	}

	b = Rotate(iH.Sum(b), EC+RRC)

	sgn := make([]byte, EC+RRC+cc)
	b = b[copy(sgn, b):]

	for i := range forSeal {
		b = b[copy(forSeal[i], b):]
	}

	// set rrc value.
	binary.BigEndian.PutUint16(hdr[6:8], RRC)

	return append(hdr, sgn...), nil
}

func (c *AESCTSHMACSHA1) Unwrap(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte, sgn []byte) (bool, error) {
	ok, err := c.unwrap(ctx, seqNum, forSign, forSeal, sgn)
	if err != nil {
		return ok, fmt.Errorf("aes-cts-hmac-sha1: unwrap: %w", err)
	}
	return ok, nil
}

func (c *AESCTSHMACSHA1) unwrap(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte, sgn []byte) (bool, error) {

	// buffer for decryption.
	eB, hdr, cc := bytes.NewBuffer(nil), sgn[:16], c.setting.Type.GetConfounderByteSize()

	// write { ec | E"header" | confounder }
	eB.Write(sgn[16:])

	for i := range forSeal {
		// write { E"data" }
		eB.Write(forSeal[i])
	}

	rrc, ec := int(binary.BigEndian.Uint16(hdr[6:])), int(binary.BigEndian.Uint16(hdr[4:]))
	binary.BigEndian.PutUint16(hdr[6:8], uint16(0))

	// rotate { ec | E"header" | confounder | E"data" | mic } ->
	//        { confounder | E"data" | ec | E"header" | mic }
	b := Rotate(eB.Bytes(), -(rrc + ec))

	cksumSize := c.setting.Type.GetHMACBitLength() / 8
	// trim mic.
	b, cksum := b[:len(b)-cksumSize], b[len(b)-cksumSize:]

	key, err := c.setting.Type.DeriveKey(c.setting.Key.KeyValue, common.GetUsageKe(uint32(c.ecU)))
	if err != nil {
		return false, fmt.Errorf("derive key: %w", err)
	}

	b, err = c.setting.Type.DecryptData(key, b)
	if err != nil {
		return false, fmt.Errorf("decrypt data: %w", err)
	}

	iH, err := c.IntegrityHash()
	if err != nil {
		return false, fmt.Errorf("make integrity hash: %w", err)
	}

	// write confounder.
	iH.Write(b[:cc])
	b = b[cc:]

	for i := range forSeal {
		// decrypt the encrypred data.
		b = b[copy(forSeal[i], b):]
	}

	for i := range forSign {
		// write data for signing.
		iH.Write(forSign[i])
	}

	// write remaining of the header.
	iH.Write(b)

	return bytes.Equal(iH.Sum(nil), cksum), nil
}

func (c *AESCTSHMACSHA1) MakeSignature(ctx context.Context, seqNum uint64, forSgn [][]byte) ([]byte, error) {
	b, err := c.makeSignature(ctx, seqNum, forSgn)
	if err != nil {
		return nil, fmt.Errorf("aes-cts-hmac-sha1: make signature: %w", err)
	}
	return b, nil
}

func (c *AESCTSHMACSHA1) makeSignature(ctx context.Context, seqNum uint64, forSgn [][]byte) ([]byte, error) {

	// checksum hash.
	ckH, err := c.ChecksumHash()
	if err != nil {
		return nil, err
	}

	hdr := c.MICHeader(ctx, seqNum)

	for _, b := range forSgn {
		ckH.Write(b)
	}

	ckH.Write(hdr)

	return append(hdr, ckH.Sum(nil)...), nil
}

func (c *AESCTSHMACSHA1) Size(ctx context.Context, conf bool) int {
	sz := (16 /* hdr */ + 12 /* cksum */)
	if conf {
		sz += (16 /* confounder */ + 16 /* E"header" */ + 16 /* ec */)
	}
	return sz
}

func (c *AESCTSHMACSHA1) WrapHeader(ctx context.Context, seqNum uint64) []byte {
	hdr := make([]byte, 16)
	// token_type 0x05 0x04
	hdr[0] = 0x05
	hdr[1] = 0x04
	// flags
	hdr[2] = c.flags() | CFXFlagSealed
	// filler.
	hdr[3] = 0xff
	// sequence_number.
	binary.BigEndian.PutUint64(hdr[8:], uint64(seqNum))
	return hdr
}

func (c *AESCTSHMACSHA1) MICHeader(ctx context.Context, seqNum uint64) []byte {
	hdr := make([]byte, 16)
	// token_type 0x04 0x04
	hdr[0] = 0x04
	hdr[1] = 0x04
	// flags
	hdr[2] = c.flags()
	// filler.
	hdr[3] = 0xff
	hdr[4] = 0xff
	hdr[5] = 0xff
	hdr[6] = 0xff
	hdr[7] = 0xff
	// sequence_number.
	binary.BigEndian.PutUint64(hdr[8:], uint64(seqNum))
	return hdr
}

type SizedHash struct {
	hash.Hash
	etype etype.EType
	usage uint32
}

func (h *SizedHash) Sum(b []byte) []byte {
	return h.Hash.Sum(b)[:len(b)+h.etype.GetHMACBitLength()/8]
}

// Rotate function rotates the buffer over rotation point `rc`, ie
// out = in[rc:] + in[:rc].
func Rotate(b []byte, rc int) []byte {
	out, left := make([]byte, len(b)), (len(b)-rc)%len(b)
	copy(out[copy(out, b[left:]):], b[:left])
	return out
}
