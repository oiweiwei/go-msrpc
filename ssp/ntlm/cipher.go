package ntlm

import (
	"bytes"
	"context"
	"encoding/binary"

	"hash"
	"hash/crc32"

	"crypto/hmac"
	"crypto/md5"
	"crypto/rc4"
)

// The HMAC-MD5 checksum.
func HMACMD5(key []byte) func(uint32) hash.Hash {
	return func(seqNum uint32) hash.Hash {
		h := hmac.New(md5.New, key)
		binary.Write(h, binary.LittleEndian, seqNum)
		return h
	}
}

// The CRC32 checksum.
func CRC32(key []byte) func(uint32) hash.Hash {
	return func(_ uint32) hash.Hash { return crc32.NewIEEE() }
}

// RC4K performs XOR key stream on bytes `src` once.
func RC4K(ctx context.Context, key []byte, src []byte) ([]byte, error) {
	c, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	dst := make([]byte, len(src))
	c.XORKeyStream(dst, src)
	return dst, nil
}

func NewCipher(ctx context.Context, key []byte, hashFunc func(uint32) hash.Hash) (*Cipher, error) {
	c, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return &Cipher{c, hashFunc}, nil
}

type Cipher struct {
	cipher   *rc4.Cipher
	hashFunc func(uint32) hash.Hash
}

// XORKeyStream function XOr-s the data with a random cipher.
func (c *Cipher) XORKeyStream(data any) error {

	// xor byte stream.
	if b, ok := data.([]byte); ok {
		c.cipher.XORKeyStream(b, b)
		return nil
	}

	// xor generic data.
	buf := bytes.NewBuffer(nil)
	if err := binary.Write(buf, binary.LittleEndian, data); err != nil {
		return err
	}

	b := buf.Bytes()
	c.cipher.XORKeyStream(b, b)

	if err := binary.Read(bytes.NewBuffer(b), binary.LittleEndian, data); err != nil {
		return err
	}

	return nil
}

// HMAC_MD5 function computes hash over the data provided and returns
// the first 8 bytes.
func (c *Cipher) Checksum(seqNum uint32, ms ...[]byte) ([]byte, error) {

	h := c.hashFunc(seqNum)

	for _, m := range ms {
		h.Write(m)
	}

	var pad []byte

	if h.Size() < 8 {
		pad = make([]byte, 8-h.Size())
	}

	return h.Sum(pad)[:8], nil
}
