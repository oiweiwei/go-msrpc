package crypto

import (
	"context"
	"crypto/rand"
	"fmt"

	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
	"github.com/oiweiwei/go-msrpc/ssp/krb5/crypto/rfc1964"
)

type DES struct {
	setting CipherSetting
}

func NewDESCipher(ctx context.Context, setting CipherSetting) (Cipher, error) {
	return &DES{setting: setting}, nil
}

func (c *DES) Wrap(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte) ([]byte, error) {
	b, err := c.wrap(ctx, seqNum, forSign, forSeal)
	if err != nil {
		return nil, fmt.Errorf("des: wrap: %w", err)
	}
	return b, nil
}

func (c *DES) wrap(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte) ([]byte, error) {

	key := c.setting.Key

	tok := rfc1964.NewWrapToken()
	tok.SetSequenceNumber(uint32(seqNum), c.setting.IsLocal)
	tok.SetEncryption(len(forSeal) > 0)

	_, err := rand.Read(tok.Confounder)
	if err != nil {
		return nil, err
	}

	cksum, err := rfc1964.DESMACMD5(key, tok.Header(), tok.Confounder, forSign)
	if err != nil {
		return nil, err
	}
	copy(tok.Checksum, cksum)

	if err := rfc1964.EncryptSequenceNumber(key, cksum, tok.SequenceNumber); err != nil {
		return nil, err
	}

	if tok.UseEncryption() {
		cipher := rfc1964.NewCipher(key)
		cipher.AddBuffer(tok.Confounder)
		for _, b := range forSeal {
			cipher.AddBuffer(b)
		}
		if err := cipher.Encrypt(); err != nil {
			return nil, err
		}
	}

	return EncodeASN1Value(tok.Marshal(), KRB5OID, true /* always use dce-style */, forSign)
}

func (c *DES) Unwrap(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte, sgn []byte) (bool, error) {
	ok, err := c.unwrap(ctx, seqNum, forSign, forSeal, sgn)
	if err != nil {
		return ok, fmt.Errorf("des: unwrap: %w", err)
	}
	return ok, nil
}

func (c *DES) unwrap(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte, sgn []byte) (bool, error) {

	key := c.setting.Key

	expTok := rfc1964.NewWrapToken()
	expTok.SetSequenceNumber(uint32(seqNum), c.setting.IsLocal)

	// trim asn1 header.
	sgn, _, err := ParseASN1Value(sgn, forSign)
	if err != nil {
		return false, gssapi.ErrDefectiveToken
	}

	tok := rfc1964.NewWrapToken()
	if err := tok.Unmarshal(sgn); err != nil {
		return false, gssapi.ErrDefectiveToken
	}

	expTok.SetEncryption(tok.UseEncryption())

	if tok.UseEncryption() {
		cipher := rfc1964.NewCipher(key)
		cipher.AddBuffer(tok.Confounder)
		for _, b := range forSeal {
			cipher.AddBuffer(b)
		}
		if err := cipher.Decrypt(); err != nil {
			return false, err
		}
	}

	// fix confounder.
	copy(expTok.Confounder, tok.Confounder)

	// compute checksum.
	cksum, err := rfc1964.DESMACMD5(key, tok.Header(), tok.Confounder, forSign)
	if err != nil {
		return false, err
	}
	copy(expTok.Checksum, cksum)

	if err := rfc1964.DecryptSequenceNumber(key, cksum, tok.SequenceNumber); err != nil {
		return false, err
	}

	return expTok.Equals(tok), nil
}

func (c *DES) MakeSignature(ctx context.Context, seqNum uint64, forSign [][]byte) ([]byte, error) {

	key := c.setting.Key

	tok := rfc1964.NewMICToken()
	tok.SetSequenceNumber(uint32(seqNum), c.setting.IsLocal)

	cksum, err := rfc1964.DESMACMD5(key, tok.Header(), forSign)
	if err != nil {
		return nil, err
	}
	copy(tok.Checksum, cksum)

	if err := rfc1964.EncryptSequenceNumber(key, cksum, tok.SequenceNumber); err != nil {
		return nil, err
	}

	return EncodeASN1Value(tok.Marshal(), KRB5OID, true /* always use dce-style */, forSign)
}

func (c *DES) Size(ctx context.Context, conf bool) int {
	sz := (16 /* header[16] */ + 8 /* cksum[8] */ + 13 /* asn1 header */)
	if conf {
		sz += 8 /* confounder[8] */
	}
	return sz
}
