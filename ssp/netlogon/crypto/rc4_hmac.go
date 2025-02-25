package crypto

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rc4"

	"github.com/oiweiwei/go-msrpc/ssp/crypto"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
)

type RC4MD5 struct {
	key      []byte
	isClient bool
}

var (
	rc4Sign = [...]byte{0x77, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00}
	rc4Seal = [...]byte{0x77, 0x00, 0x7A, 0x00, 0xFF, 0xFF, 0x00, 0x00}
)

func (c *RC4MD5) MICHeader(ctx context.Context, seqNum uint64) []byte {
	return append(rc4Sign[:], MakeSeqNum(ctx, seqNum, c.isClient)...)
}

func (c *RC4MD5) WrapHeader(ctx context.Context, seqNum uint64) []byte {
	return append(rc4Seal[:], MakeSeqNum(ctx, seqNum, c.isClient)...)
}

func (c *RC4MD5) Size(ctx context.Context, conf bool) int {
	sz := (8 /* header */ + 8 /* seq_num */ + 8 /* cksum */)
	if conf {
		sz += 8 /* confounder */
	}
	return sz
}

func (c *RC4MD5) MakeSignature(ctx context.Context, seqNum uint64, forSign [][]byte) ([]byte, error) {

	sgn := c.MICHeader(ctx, seqNum)

	tmp := md5.New()
	// write zero pad (4).
	tmp.Write(make([]byte, 4))
	// write header.
	tmp.Write(sgn[:8])

	for _, b := range forSign {
		// write buffer.
		tmp.Write(b)
	}

	// hmac digest `tmp`.
	ckH := hmac.New(md5.New, c.key)
	ckH.Write(tmp.Sum(nil))

	// add checksum to signature.
	sgn = append(sgn, ckH.Sum(nil)[:8]...)

	// derive first key for sequence number encryption.
	k1, err := crypto.HMACMD5(c.key, make([]byte, 4))
	if err != nil {
		return nil, err
	}

	// derivce sequence number encryption key.
	snK, err := crypto.HMACMD5(k1, sgn[16:24])
	if err != nil {
		return nil, err
	}

	// rc4 cipher.
	snC, err := rc4.NewCipher(snK)
	if err != nil {
		return nil, err
	}

	// encrypt sequence number.
	snC.XORKeyStream(sgn[8:16], sgn[8:16])

	return sgn, nil
}

func (c *RC4MD5) EncryptionKey(ctx context.Context, seqNum uint64) ([]byte, error) {

	eK0 := make([]byte, len(c.key))
	for i := range eK0 {
		eK0[i] = c.key[i] ^ 0xF0
	}

	eK1, err := crypto.HMACMD5(eK0, make([]byte, 4))
	if err != nil {
		return nil, err
	}

	eK, err := crypto.HMACMD5(eK1, MakeSeqNum(ctx, seqNum, c.isClient))
	if err != nil {
		return nil, err
	}

	return eK, nil
}

func (c *RC4MD5) Wrap(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte) ([]byte, error) {

	sgn := c.WrapHeader(ctx, seqNum)

	confounder := make([]byte, 8)

	if _, err := rand.Read(confounder); err != nil {
		return nil, err
	}

	tmp := md5.New()
	// write zero pad (4).
	tmp.Write(make([]byte, 4))
	// write header.
	tmp.Write(sgn[:8])
	// write confounder.
	tmp.Write(confounder)

	for _, b := range forSign {
		// write buffer.
		tmp.Write(b)
	}

	// hmac digest `tmp`.
	ckH := hmac.New(md5.New, c.key)
	ckH.Write(tmp.Sum(nil))

	// add checksum to signature.
	sgn = append(sgn, ckH.Sum(nil)[:8]...)

	eK, err := c.EncryptionKey(ctx, seqNum)
	if err != nil {
		return nil, err
	}

	eC, err := rc4.NewCipher(eK)
	if err != nil {
		return nil, err
	}

	// FIXME-MSDN: [MS-NRPC 3.3.4.2.1]:
	// The server MUST use this key to initialize RC4 and encrypt the Confounder
	// field and then the data. The server MUST initialize RC4 only once, before
	// encrypting the Confounder field.
	//
	// However, impacket implementation works this way (and Windows Server too).

	// encrypt confounder.
	eC.XORKeyStream(confounder, confounder)
	for _, b := range forSeal {
		eC, _ := rc4.NewCipher(eK)
		// encrypt buffer.
		eC.XORKeyStream(b, b)
	}

	// derive first key for sequence number encryption.
	snK1, err := crypto.HMACMD5(c.key, make([]byte, 4))
	if err != nil {
		return nil, err
	}

	// derive sequence number encryption key.
	snK, err := crypto.HMACMD5(snK1, sgn[16:24])
	if err != nil {
		return nil, err
	}

	// rc4 cipher.
	snC, err := rc4.NewCipher(snK)
	if err != nil {
		return nil, err
	}

	// encrypt sequence number.
	snC.XORKeyStream(sgn[8:16], sgn[8:16])

	sgn = append(sgn, confounder...)

	return sgn, nil

}

func (c *RC4MD5) Unwrap(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte, sgn []byte) ([]byte, error) {

	if len(sgn) < 32 {
		return nil, gssapi.ErrDefectiveToken
	}

	confounder := sgn[24:32]

	eK, err := c.EncryptionKey(ctx, seqNum)
	if err != nil {
		return nil, err
	}

	eC, err := rc4.NewCipher(eK)
	if err != nil {
		return nil, err
	}

	eC.XORKeyStream(confounder, confounder)
	for _, b := range forSeal {
		eC, _ := rc4.NewCipher(eK)
		eC.XORKeyStream(b, b)
	}

	tmp := md5.New()
	// write zero pad (4).
	tmp.Write(make([]byte, 4))
	// write header.
	tmp.Write(sgn[:8])
	// write confounder.
	tmp.Write(confounder)

	for _, b := range forSign {
		// write buffer.
		tmp.Write(b)
	}

	// hmac digest `tmp`.
	ckH := hmac.New(md5.New, c.key)
	ckH.Write(tmp.Sum(nil))

	// derive first key for sequence number decryption.
	snK1, err := crypto.HMACMD5(c.key, make([]byte, 4))
	if err != nil {
		return nil, err
	}

	// derive sequence number decryption key.
	snK, err := crypto.HMACMD5(snK1, sgn[16:24])
	if err != nil {
		return nil, err
	}

	// rc4 cipher.
	snC, err := rc4.NewCipher(snK)
	if err != nil {
		return nil, err
	}

	// decrypt sequence number.
	snC.XORKeyStream(sgn[8:16], sgn[8:16])

	// add checksum to signature.
	expSgn := c.WrapHeader(ctx, seqNum)
	expSgn = append(expSgn, ckH.Sum(nil)[:8]...)
	expSgn = append(expSgn, confounder...)

	return expSgn, nil
}
