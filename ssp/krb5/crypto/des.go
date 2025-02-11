package crypto

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"fmt"

	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
	"github.com/oiweiwei/gokrb5.fork/v9/crypto/etype"
	"github.com/oiweiwei/gokrb5.fork/v9/crypto/rfc3961"
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

	hdr := c.WrapHeader(ctx, seqNum)

	_, err := rand.Read(hdr[24:32])
	if err != nil {
		return nil, err
	}

	cksum, err := DESMACMD5(ctx, c.setting.Key.KeyValue, c.setting.Type, hdr[:8], hdr[24:32], forSign)
	if err != nil {
		return nil, err
	}
	copy(hdr[16:24], cksum)

	_, seq, err := rfc3961.DESEncryptData(c.setting.Key.KeyValue, hdr[8:16], cksum, c.setting.Type)
	if err != nil {
		return nil, err
	}
	copy(hdr[8:16], seq)

	eB := bytes.NewBuffer(nil)
	eB.Write(hdr[24:32])
	for _, b := range forSeal {
		eB.Write(b)
	}

	key := make([]byte, len(c.setting.Key.KeyValue))
	for i := range key {
		key[i] = c.setting.Key.KeyValue[i] ^ 0xF0
	}

	_, b, err := c.setting.Type.EncryptData(key, eB.Bytes())
	if err != nil {
		return nil, err
	}

	b = b[copy(hdr[24:32], b):]

	for i := range forSeal {
		b = b[copy(forSeal[i], b):]
	}

	if len(b) > 0 {
		return nil, fmt.Errorf("unexpected data left after encryption")
	}

	return EncodeASN1Value(hdr, KRB5OID, true /* always use dce-style */, forSign)
}

func (c *DES) Unwrap(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte, sgn []byte) (bool, error) {
	ok, err := c.unwrap(ctx, seqNum, forSign, forSeal, sgn)
	if err != nil {
		return ok, fmt.Errorf("des: unwrap: %w", err)
	}
	return ok, nil
}

func (c *DES) unwrap(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte, sgn []byte) (bool, error) {

	// trim asn1 header.
	sgn, _, err := ParseASN1Value(sgn, forSign)
	if err != nil {
		return false, gssapi.ErrDefectiveToken
	}

	hdr := c.WrapHeader(ctx, seqNum)

	eB := bytes.NewBuffer(nil)
	eB.Write(sgn[24:32])
	for _, b := range forSeal {
		eB.Write(b)
	}

	key := make([]byte, len(c.setting.Key.KeyValue))
	for i := range key {
		key[i] = c.setting.Key.KeyValue[i] ^ 0xF0
	}

	b, err := c.setting.Type.DecryptData(key, eB.Bytes())
	if err != nil {
		return false, err
	}

	b = b[copy(hdr[24:32], b):]

	for i := range forSeal {
		if len(b) < len(forSeal[i]) {
			return false, fmt.Errorf("unexpected data left after decryption")
		}
		b = b[copy(forSeal[i], b):]
	}

	if len(b) > 0 {
		return false, fmt.Errorf("unexpected data left after decryption")
	}

	// fix confounder.
	copy(sgn[24:32], hdr[24:32])

	// compute checksum.
	cksum, err := DESMACMD5(ctx, c.setting.Key.KeyValue, c.setting.Type, hdr[:8], hdr[24:32], forSign)
	if err != nil {
		return false, err
	}
	copy(hdr[16:24], cksum)

	// decrypt sequence number.
	seq, err := rfc3961.DESDecryptData(c.setting.Key.KeyValue, sgn[8:16], cksum, c.setting.Type)
	if err != nil {
		return false, err
	}
	copy(sgn[8:16], seq)

	return bytes.Equal(sgn, hdr), nil
}

func DESMACMD5(ctx context.Context, key []byte, eType etype.EType, hdr, confounder []byte, data [][]byte) ([]byte, error) {

	iH := md5.New()

	iH.Write(hdr)
	iH.Write(confounder)
	for _, b := range data {
		iH.Write(b)
	}

	cksum, _, err := eType.EncryptData(key, iH.Sum(nil))
	if err != nil {
		return nil, err
	}

	return cksum, nil
}

func (c *DES) MakeSignature(ctx context.Context, seqNum uint64, forSign [][]byte) ([]byte, error) {

	hdr := c.MICHeader(ctx, seqNum)

	cksum, err := DESMACMD5(ctx, c.setting.Key.KeyValue, c.setting.Type, hdr[:8], nil, forSign)
	if err != nil {
		return nil, err
	}
	copy(hdr[16:24], cksum)

	_, seq, err := rfc3961.DESEncryptData(c.setting.Key.KeyValue, hdr[8:16], cksum, c.setting.Type)
	if err != nil {
		return nil, err
	}
	copy(hdr[8:16], seq)

	return EncodeASN1Value(hdr, KRB5OID, true /* always use dce-style */, forSign)
}

func (c *DES) Size(ctx context.Context, conf bool) int {
	sz := (16 /* header[16] */ + 8 /* cksum[8] */ + 13 /* asn1 header */)
	if conf {
		sz += 8 /* confounder[8] */
	}
	return sz
}

func (c *DES) WrapHeader(ctx context.Context, seqNum uint64) []byte {

	hdr := make([]byte, 32)

	// token_id.
	hdr[0] = 0x02
	hdr[1] = 0x01
	// sgn_alg.
	hdr[2] = 0x00 // des mac md5
	hdr[3] = 0x00
	// seal_alg.
	hdr[4] = 0x00 // des
	hdr[5] = 0x00
	// filler.
	hdr[6] = 0xff
	hdr[7] = 0xff
	// seq_number.
	binary.LittleEndian.PutUint32(hdr[8:12], uint32(seqNum))
	// filler.
	if !c.setting.IsLocal {
		hdr[12] = 0xff
		hdr[13] = 0xff
		hdr[14] = 0xff
		hdr[15] = 0xff
	}

	return hdr
}

func (c *DES) MICHeader(ctx context.Context, seqNum uint64) []byte {

	hdr := make([]byte, 24)

	// token_id.
	hdr[0] = 0x01
	hdr[1] = 0x01
	// sgn_alg.
	hdr[2] = 0x00 // des mac md5
	hdr[3] = 0x00
	// filler.
	hdr[4] = 0xff
	hdr[5] = 0xff
	hdr[6] = 0xff
	hdr[7] = 0xff
	// seq_number.
	binary.LittleEndian.PutUint32(hdr[8:12], uint32(seqNum))
	// filler.
	if !c.setting.IsLocal {
		hdr[12] = 0xff
		hdr[13] = 0xff
		hdr[14] = 0xff
		hdr[15] = 0xff
	}

	return hdr
}
