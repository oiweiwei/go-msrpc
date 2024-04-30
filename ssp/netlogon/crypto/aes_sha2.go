package crypto

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"

	"github.com/oiweiwei/go-msrpc/ssp/crypto"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
)

type AESSHA2 struct {
	key      []byte
	isClient bool
}

var (
	aesMICHeader  = [...]byte{0x13, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00}
	aesWrapHeader = [...]byte{0x13, 0x00, 0x1A, 0x00, 0xFF, 0xFF, 0x00, 0x00}
)

func (c *AESSHA2) MICHeader(ctx context.Context, seqNum uint64) []byte {
	return append(aesMICHeader[:], MakeSeqNum(ctx, seqNum, c.isClient)...)
}

func (c *AESSHA2) WrapHeader(ctx context.Context, seqNum uint64) []byte {
	return append(aesWrapHeader[:], MakeSeqNum(ctx, seqNum, c.isClient)...)
}

func (c *AESSHA2) Size(ctx context.Context, conf bool) int {
	sz := (8 /* header */ + 8 /* seq_number */ + 8 /* cksum */ + 24 /* pad */)
	sz += 8 /* confounder */
	return sz
}

func (c *AESSHA2) DeriveKey(ctx context.Context, key []byte) []byte {
	local := make([]byte, len(key))
	for i := range local {
		local[i] = key[i] ^ 0xF0
	}
	return local
}

func (c *AESSHA2) MakeSignature(ctx context.Context, seqNum uint64, forSign [][]byte) ([]byte, error) {

	sgn := c.MICHeader(ctx, seqNum)

	ckH := hmac.New(sha256.New, c.key)

	ckH.Write(sgn[:8])
	for _, b := range forSign {
		ckH.Write(b)
	}

	ckSum := ckH.Sum(nil)[:8]

	snC := crypto.NewAESCFB8(c.key, bytes.Repeat(ckSum, 2), false) /* always false */
	snC.XORKeyStream(sgn[8:16], sgn[8:16])

	sgn = append(sgn, ckSum...)
	sgn = append(sgn, make([]byte, 24)...)

	return sgn, nil
}

func (c *AESSHA2) Wrap(ctx context.Context, seqNum uint64, forSign [][]byte, forSeal [][]byte) ([]byte, error) {

	sgn := c.WrapHeader(ctx, seqNum)

	confounder := make([]byte, 8)

	if _, err := rand.Read(confounder); err != nil {
		return nil, err
	}

	ckH := hmac.New(sha256.New, c.key)

	// write header.
	ckH.Write(sgn[:8])
	// write confounder.
	ckH.Write(confounder)
	for _, b := range forSign {
		// write buffers.
		ckH.Write(b)
	}

	ckSum := ckH.Sum(nil)[:8]

	eC := crypto.NewAESCFB8(c.DeriveKey(ctx, c.key), bytes.Repeat(sgn[8:16], 2), false)
	// encrypt confounder.
	eC.XORKeyStream(confounder, confounder)
	for _, b := range forSeal {
		// encrypt buffer.
		eC.XORKeyStream(b, b)
	}

	snC := crypto.NewAESCFB8(c.key, bytes.Repeat(ckSum, 2), false)
	snC.XORKeyStream(sgn[8:16], sgn[8:16])

	sgn = append(sgn, ckSum...)
	sgn = append(sgn, confounder...)
	sgn = append(sgn, make([]byte, 24)...)

	return sgn, nil
}

func (c *AESSHA2) Unwrap(ctx context.Context, seqNum uint64, forSign, forSeal [][]byte, sgn []byte) ([]byte, error) {

	if len(sgn) < 32 {
		return nil, gssapi.ErrDefectiveToken
	}

	ckSum, confounder := sgn[16:24], sgn[24:32]

	snC := crypto.NewAESCFB8(c.key, bytes.Repeat(ckSum, 2), true)
	snC.XORKeyStream(sgn[8:16], sgn[8:16])

	eC := crypto.NewAESCFB8(c.DeriveKey(ctx, c.key), bytes.Repeat(sgn[8:16], 2), true)
	// decrypt confounder.
	eC.XORKeyStream(confounder, confounder)
	for _, b := range forSeal {
		// decrypt buffer.
		eC.XORKeyStream(b, b)
	}

	ckH := hmac.New(sha256.New, c.key)

	// write header.
	ckH.Write(sgn[:8])
	// write confounder.
	ckH.Write(confounder)
	for _, b := range forSign {
		// write buffers.
		ckH.Write(b)
	}

	expSgn := c.WrapHeader(ctx, seqNum)
	expSgn = append(expSgn, ckH.Sum(nil)[:8]...)
	expSgn = append(expSgn, confounder...)
	expSgn = append(expSgn, make([]byte, 24)...)

	return expSgn, nil
}
