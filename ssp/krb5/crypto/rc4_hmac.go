package crypto

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rc4"
	"encoding/asn1"
	"encoding/binary"
	"errors"
	"fmt"
	"hash"

	"github.com/jcmturner/gokrb5/v8/asn1tools"
	"github.com/jcmturner/gokrb5/v8/types"

	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
)

var (
	KRB5OID = asn1.ObjectIdentifier{1, 2, 840, 113554, 1, 2, 2}
)

type RC4HMAC struct {
	// indicates whether this is an initiator.
	init bool
	// encryption key.
	key types.EncryptionKey
}

func NewRC4Cipher(ctx context.Context, key types.EncryptionKey, isServer bool) (Cipher, error) {
	return &RC4HMAC{init: !isServer, key: key}, nil
}

func (c *RC4HMAC) Wrap(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte) ([]byte, error) {
	b, err := c.wrap(ctx, seqNum, forSign, forSeal)
	if err != nil {
		return nil, fmt.Errorf("rc4-hmac: wrap: %w", err)
	}
	return b, nil
}

func (c *RC4HMAC) wrap(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte) ([]byte, error) {

	hdr := c.WrapHeader(ctx, seqNum)

	_, err := rand.Read(hdr[24:32])
	if err != nil {
		return nil, err
	}

	// Ksign = HMAC(Kss, "signaturekey");
	// Sgn_Cksum = MD5((int32)13, Token.Header, data);
	// Sgn_Cksum = HMAC(Ksign, Sgn_Cksum);
	cksum := c.hmac(c.hmac(c.key.KeyValue, []byte("signaturekey\x00")),
		c.md5(int32(13), hdr[:8], hdr[24:32], forSign))
	// memcpy(Token.SGN_CKSUM, Sgn_Cksum, 8);
	copy(hdr[16:24], cksum[:8])

	// for (i = 0; i < 16; i++) Klocal[i] = Kss[i] ^ 0xF0;
	klocal := make([]byte, 16)
	for i := range klocal {
		klocal[i] = c.key.KeyValue[i] ^ 0xF0
	}

	// Kcrypt = HMAC(Klocal, (int32)0);
	// Kcrypt = HMAC(Kcrypt, (int32)seq);
	// RC4(Kcrypt, Token.Confounder);
	cipher, _ := rc4.NewCipher(c.hmac(c.hmac(klocal, int32(0)), hdr[8:12]))
	cipher.XORKeyStream(hdr[24:32], hdr[24:32])
	for i := range forSeal {
		cipher.XORKeyStream(forSeal[i], forSeal[i])
	}

	// Kseq = HMAC(Kss, (int32)0);
	// Kseq = HMAC(Kseq, Token.SGN_CKSUM);
	// RC4(Kseq, Token.SND_SEQ);
	c.rc4(c.hmac(c.hmac(c.key.KeyValue, int32(0)), cksum[:8]), hdr[8:16])

	return c.withASN1Header(ctx, hdr)
}

func (c *RC4HMAC) withASN1Header(ctx context.Context, b []byte) ([]byte, error) {

	oid, err := asn1.Marshal(KRB5OID)
	if err != nil {
		return nil, fmt.Errorf("marshal krb5oid: %w", err)
	}

	return asn1tools.AddASNAppTag(append(oid, b...), 0), nil
}

func (c *RC4HMAC) Unwrap(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte, sgn []byte) (bool, error) {
	ok, err := c.unwrap(ctx, seqNum, forSign, forSeal, sgn)
	if err != nil {
		return ok, fmt.Errorf("rc4-hmac: unwrap: %w", err)
	}
	return ok, nil
}

var errDataTruncated = asn1.SyntaxError{Msg: "data truncated"}

// fixASN1Header function works with non-DCE (non-mutual) authentication RC4-HMAC tokens.
// These tokens are constructed over complete wrapped payload, hence the ASN.1 header length
// is seems to be truncated.
func (c *RC4HMAC) fixASN1Header(ctx context.Context, b []byte, forSign [][]byte) ([]byte, error) {

	var val asn1.RawValue

	// try marshal token, if content is valid, return as is.
	_, err := asn1.Unmarshal(b, &val)
	if err == nil || !errors.Is(err, errDataTruncated) {
		return b, err
	}

	// data truncated error returned, we need to include payload
	// to match the ASN.1 header length.

	sz := 0
	for _, forSign := range forSign {
		sz += len(forSign)
	}

	// unmarshal again.
	_, err = asn1.Unmarshal(append(b, make([]byte, sz)...), &val)
	if err != nil {
		return b, err
	}

	oid, err := asn1.Marshal(KRB5OID)
	if err != nil {
		return nil, fmt.Errorf("marshal krb5oid: %w", err)
	}

	// trim oid.
	sgn := val.Bytes[len(oid):]
	// trim trailer and add oid again.
	return c.withASN1Header(ctx, sgn[:32])

}

func (c *RC4HMAC) unwrap(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte, sgn []byte) (bool, error) {

	if len(sgn) < 8 {
		return false, gssapi.ErrDefectiveToken
	}

	sgn, err := c.fixASN1Header(ctx, sgn, forSign)
	if err != nil {
		return false, gssapi.ErrDefectiveToken
	}

	hdr := c.WrapHeader(ctx, seqNum)
	copy(hdr[24:32], sgn[len(sgn)-8:])

	// for (i = 0; i < 16; i++) Klocal[i] = Kss[i] ^ 0xF0;
	klocal := make([]byte, 16)
	for i := range klocal {
		klocal[i] = c.key.KeyValue[i] ^ 0xF0
	}

	// Kcrypt = HMAC(Klocal, (int32)0);
	// Kcrypt = HMAC(Kcrypt, (int32)seq);
	// RC4(Kcrypt, Token.Confounder);
	cipher, _ := rc4.NewCipher(c.hmac(c.hmac(klocal, int32(0)), hdr[8:12]))
	cipher.XORKeyStream(hdr[24:32], hdr[24:32])
	for i := range forSeal {
		cipher.XORKeyStream(forSeal[i], forSeal[i])
	}

	// Ksign = HMAC(Kss, "signaturekey");
	// Sgn_Cksum = MD5((int32)13, Token.Header, data);
	// Sgn_Cksum = HMAC(Ksign, Sgn_Cksum);
	cksum := c.hmac(c.hmac(c.key.KeyValue, []byte("signaturekey\x00")),
		c.md5(int32(13), hdr[:8], hdr[24:32], forSign))
	// memcpy(Token.SGN_CKSUM, Sgn_Cksum, 8);
	copy(hdr[16:24], cksum[:8])

	// Kseq = HMAC(Kss, (int32)0);
	// Kseq = HMAC(Kseq, Token.SGN_CKSUM);
	// RC4(Kseq, Token.SND_SEQ);
	c.rc4(c.hmac(c.hmac(c.key.KeyValue, int32(0)), cksum[:8]), hdr[8:16])

	expSgn, err := c.withASN1Header(ctx, hdr)
	if err != nil {
		return false, err
	}

	// set decrypted confounder back.
	copy(sgn[len(sgn)-8:], hdr[24:32])

	return bytes.Equal(expSgn, sgn), nil
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

	hdr := c.MICHeader(ctx, seqNum)

	// Ksign = HMAC(Kss, "signaturekey");
	// Sgn_Cksum = MD5((int32)15, Token.Header, data);
	// Sgn_Cksum = HMAC(Ksign, Sgn_Cksum);
	cksum := c.hmac(c.hmac(c.key.KeyValue, []byte("signaturekey\x00")),
		c.md5(int32(15), hdr[:8], forSign))
	// memcpy(Token.SGN_CKSUM, Sgn_Cksum, 8);
	copy(hdr[16:24], cksum[:8])
	// Kseq = HMAC(Kss, (int32)0);
	// Kseq = HMAC(Kseq, Token.SGN_CKSUM);
	// RC4(Kseq, Token.SND_SEQ);
	c.rc4(c.hmac(c.hmac(c.key.KeyValue, int32(0)), cksum[:8]), hdr[8:16])

	return c.withASN1Header(ctx, hdr)
}

func (c *RC4HMAC) hmac(k []byte, data ...interface{}) []byte {
	return c.digest(hmac.New(md5.New, k), data...)
}

func (c *RC4HMAC) md5(data ...interface{}) []byte {
	return c.digest(md5.New(), data...)
}

func (c *RC4HMAC) rc4(key []byte, data []byte) {
	cipher, _ := rc4.NewCipher(key)
	cipher.XORKeyStream(data, data)
}

func (c *RC4HMAC) digest(h hash.Hash, data ...interface{}) []byte {
	for _, d := range data {
		switch d := d.(type) {
		case []byte:
			h.Write(d)
		case [][]byte:
			for i := range d {
				h.Write(d[i])
			}
		default:
			binary.Write(h, binary.LittleEndian, d)
		}
	}
	return h.Sum(nil)
}

func (c *RC4HMAC) MICHeader(ctx context.Context, seqNum uint64) []byte {

	hdr := make([]byte, 24)

	// token_type 0x01 0x01
	hdr[0] = 0x01
	hdr[1] = 0x01
	// HMAC
	hdr[2] = 0x11
	hdr[3] = 0x00
	// filler.
	hdr[4] = 0xff
	hdr[5] = 0xff
	hdr[6] = 0xff
	hdr[7] = 0xff
	// seq_number.
	binary.BigEndian.PutUint32(hdr[8:12], uint32(seqNum))
	// filler.
	if !c.init {
		hdr[12] = 0xff
		hdr[13] = 0xff
		hdr[14] = 0xff
		hdr[15] = 0xff
	}
	return hdr

}

func (c *RC4HMAC) WrapHeader(ctx context.Context, seqNum uint64) []byte {

	hdr := make([]byte, 32)

	// token_id.
	hdr[0] = 0x02
	hdr[1] = 0x01
	// sgn_alg.
	hdr[2] = 0x11 // hmac
	hdr[3] = 0x00
	// seal_alg.
	hdr[4] = 0x10
	hdr[5] = 0x00
	// filler.
	hdr[6] = 0xff
	hdr[7] = 0xff
	// seq_number.
	binary.BigEndian.PutUint32(hdr[8:12], uint32(seqNum))
	// filler.
	if !c.init {
		hdr[12] = 0xff
		hdr[13] = 0xff
		hdr[14] = 0xff
		hdr[15] = 0xff
	}

	return hdr
}
