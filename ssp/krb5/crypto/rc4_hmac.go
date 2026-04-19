package crypto

import (
	"context"
	"crypto/rand"
	"fmt"

	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
	"github.com/oiweiwei/go-msrpc/ssp/krb5/crypto/rfc4757"
)

type RC4HMAC struct {
	setting CipherSetting
}

func NewRC4Cipher(ctx context.Context, setting CipherSetting) (Cipher, error) {
	return &RC4HMAC{setting: setting}, nil
}

func (c *RC4HMAC) ParseSignature(ctx context.Context, payload []byte) ([]byte, []byte, error) {

	_, oft, _, err := ParseASN1Value(payload, [][]byte{payload})
	if err != nil {
		return nil, nil, fmt.Errorf("des: parse header: %w", err)
	}

	if len(payload) < oft+4 {
		return nil, nil, fmt.Errorf("des: parse header: invalid payload size: %d < %d", len(payload), oft)
	}

	switch payload[oft+4] {
	case 0x02 /* wrap token */ :
		oft += 32
	case 0x01 /* mic token */ :
		oft += 24
	default:
		return nil, nil, fmt.Errorf("des: parse header: invalid token type: %d", payload[oft+4])
	}

	return payload[:oft], payload[oft:], nil
}

func (c *RC4HMAC) Wrap(ctx context.Context, seqNum uint64, payload []byte, conf bool) ([]byte, error) {
	var forSeal [][]byte
	if conf {
		forSeal = [][]byte{payload}
	}
	return c.WrapEx(ctx, seqNum, [][]byte{payload}, forSeal)
}

func (c *RC4HMAC) WrapEx(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte) ([]byte, error) {
	b, err := c.wrap(ctx, seqNum, forSign, forSeal)
	if err != nil {
		return nil, fmt.Errorf("rc4-hmac: wrap: %w", err)
	}
	return b, nil
}

func (c *RC4HMAC) wrap(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte) ([]byte, error) {

	key := c.setting.Key.KeyValue

	tok := rfc4757.NewWrapToken()
	tok.SetSequenceNumber(uint32(seqNum), c.setting.IsLocal)
	tok.SetEncryption(len(forSeal) > 0)

	_, err := rand.Read(tok.Confounder)
	if err != nil {
		return nil, err
	}

	sgnCksum, err := rfc4757.ComputeWrapChecksum(key, tok.Header(), tok.Confounder, forSign)
	if err != nil {
		return nil, err
	}
	// memcpy(Token.SGN_CKSUM, Sgn_Cksum, 8);
	copy(tok.Checksum, sgnCksum[:8])

	if tok.UseEncryption() {
		cipher, err := rfc4757.NewCipher(key, tok.SequenceNumber[:4])
		if err != nil {
			return nil, err
		}
		cipher.XORKeyStream(tok.Confounder, tok.Confounder)
		for i := range forSeal {
			cipher.XORKeyStream(forSeal[i], forSeal[i])
		}
	}

	if err := rfc4757.XORSequenceNumber(key, tok.Checksum, tok.SequenceNumber); err != nil {
		return nil, err
	}

	return EncodeASN1Value(tok.Marshal(), KRB5OID, true /* always use dce-style */, forSign)
}

func (c *RC4HMAC) Unwrap(ctx context.Context, seqNum uint64, payload []byte, sgn []byte) ([]byte, bool, error) {

	var err error

	if len(sgn) == 0 {
		// if sgn is empty, it's encoded as part of the payload.
		if sgn, payload, err = c.ParseSignature(ctx, payload); err != nil {
			return nil, false, fmt.Errorf("rc4-hmac: unwrap: parse token: %w", err)
		}
	}

	ok, err := c.unwrapEx(ctx, seqNum, [][]byte{payload}, [][]byte{payload}, sgn)
	if err != nil {
		return nil, false, fmt.Errorf("rc4-hmac: unwrap: %w", err)
	}

	return sgn, ok, nil
}

func (c *RC4HMAC) UnwrapEx(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte, sgn []byte) (bool, error) {
	ok, err := c.unwrapEx(ctx, seqNum, forSign, forSeal, sgn)
	if err != nil {
		return ok, fmt.Errorf("rc4-hmac: unwrap: %w", err)
	}
	return ok, nil
}

func (c *RC4HMAC) unwrapEx(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte, sgn []byte) (bool, error) {

	key := c.setting.Key.KeyValue

	expTok := rfc4757.NewWrapToken()
	expTok.SetSequenceNumber(uint32(seqNum), c.setting.IsLocal)

	// trim asn1 header.
	sgn, _, _, err := ParseASN1Value(sgn, forSign)
	if err != nil {
		return false, gssapi.ErrDefectiveToken
	}

	tok := rfc4757.NewWrapToken()
	if err := tok.Unmarshal(sgn); err != nil {
		return false, gssapi.ErrDefectiveToken
	}

	// copy confounder from received signature.
	copy(expTok.Confounder, tok.Confounder)

	if expTok.SetEncryption(tok.UseEncryption()); tok.UseEncryption() {
		cipher, err := rfc4757.NewCipher(key, expTok.SequenceNumber[:4])
		if err != nil {
			return false, err
		}
		cipher.XORKeyStream(expTok.Confounder, expTok.Confounder)
		for i := range forSeal {
			cipher.XORKeyStream(forSeal[i], forSeal[i])
		}
	}

	sgnCksum, err := rfc4757.ComputeWrapChecksum(key, expTok.Header(), expTok.Confounder, forSign)
	if err != nil {
		return false, err
	}
	// memcpy(Token.SGN_CKSUM, Sgn_Cksum, 8);
	copy(expTok.Checksum, sgnCksum[:8])

	if err := rfc4757.XORSequenceNumber(key, expTok.Checksum, expTok.SequenceNumber); err != nil {
		return false, err
	}

	// set decrypted confounder back.
	copy(tok.Confounder, expTok.Confounder)

	return tok.Equals(expTok), nil
}

func (c *RC4HMAC) Size(ctx context.Context, conf bool) int {
	sz := (16 /* header[16] */ + 8 /* cksum[8] */ + 13 /* asn1 header */)
	if conf {
		sz += 8 /* confounder[8] */
	}
	return sz
}

func (c *RC4HMAC) MakeSignature(ctx context.Context, seqNum uint64, forSign [][]byte) ([]byte, error) {
	b, err := c.makeSignature(ctx, seqNum, forSign)
	if err != nil {
		return nil, fmt.Errorf("rc4-hmac: make signature: %w", err)
	}
	return b, nil
}

func (c *RC4HMAC) makeSignature(ctx context.Context, seqNum uint64, forSign [][]byte) ([]byte, error) {

	key := c.setting.Key.KeyValue

	tok := rfc4757.NewMICToken()
	tok.SetSequenceNumber(uint32(seqNum), c.setting.IsLocal)

	sgnCksum, err := rfc4757.ComputeMICChecksum(key, tok.Header(), forSign)
	if err != nil {
		return nil, err
	}
	// memcpy(Token.SGN_CKSUM, Sgn_Cksum, 8);
	copy(tok.Checksum, sgnCksum[:8])

	if err := rfc4757.XORSequenceNumber(key, tok.Checksum, tok.SequenceNumber); err != nil {
		return nil, err
	}

	return EncodeASN1Value(tok.Marshal(), KRB5OID, true /* always use dce-style */, forSign)
}
