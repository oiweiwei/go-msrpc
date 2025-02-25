package crypto

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"io"
	"math/bits"

	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rc4"
	"crypto/sha256"
	"hash/crc32"

	"golang.org/x/crypto/md4"
)

var (
	_ = fmt.Sprintf("")
)

var (
	ECBSize = 7
)

func AES_CFB(k []byte, iv []byte, m []byte, decrypt bool) []byte {
	c, err := aes.NewCipher(k)
	if err != nil {
		panic(err)
	}

	var dst = make([]byte, len(m)+aes.BlockSize)

	newCFB8(c, iv, decrypt).XORKeyStream(dst, m)
	return dst[:len(m)]
}

func NewAESCFB8(k []byte, iv []byte, decrypt bool) cipher.Stream {
	return AESInit(k, iv, decrypt)
}

func AESInit(k []byte, iv []byte, decrypt bool) cipher.Stream {
	c, err := aes.NewCipher(k)
	if err != nil {
		panic(err)
	}

	return newCFB8(c, iv, decrypt)
}

func AES(handle cipher.Stream, m []byte) []byte {
	var dst = make([]byte, len(m))
	handle.XORKeyStream(dst, m)

	return dst[:len(m)]
}

func ECB(in []byte) []byte {

	var out = []byte{
		(in[0] >> 1),
		((in[0] & 0x01) << 6) | (in[1] >> 2),
		((in[1] & 0x03) << 5) | (in[2] >> 3),
		((in[2] & 0x07) << 4) | (in[3] >> 4),
		((in[3] & 0x0F) << 3) | (in[4] >> 5),
		((in[4] & 0x1F) << 2) | (in[5] >> 6),
		((in[5] & 0x3F) << 1) | (in[6] >> 7),
		((in[6] & 0x7F) >> 0),
	}

	for i := range out {
		out[i] = (out[i] << 1) & 0xFE
	}

	return out
}

func ECBP(in []byte) []byte {

	var out = ECB(in)

	for i := range out {
		out[i] |= uint8(bits.OnesCount8(out[i])) % 2
	}

	return out
}

func CBC(b []byte) []byte {
	var (
		i, pb uint8
		res   = make([]byte, 8)
	)

	b = PAD(b, 7)

	for ; i < uint8(56); i++ {
		if b[i/8]&(1<<(7-(i%8))) > 0 {
			pb++
			res[i/7] |= 1 << (7 - (i+i/7)%8)
		}

		if (i+1)%7 == 0 {
			res[i/7] |= 1 - (pb % 2)
			pb = 0
		}
	}

	return res
}

func PAD(b []byte, n int) []byte {
	var cb = make([]byte, n)
	for i, vb := range b {
		if i >= n {
			break
		}

		cb[i] = vb

	}

	return cb
}

func DES(k []byte, d []byte) []byte {

	var (
		res = make([]byte, 8)
	)

	k = PAD(k, 7)
	d = PAD(d, 8)

	block, err := des.NewCipher(CBC(k[:7]))
	if err != nil {
		panic(err)
	}

	block.Encrypt(res, d)

	return res
}

func DES_ECB(k []byte, d []byte, encrypt bool) []byte {

	var (
		res = make([]byte, 8)
	)

	k = PAD(k, 7)
	d = PAD(d, 8)

	block, err := des.NewCipher(ECB(k[:7]))
	if err != nil {
		panic(err)
	}

	if encrypt {
		block.Encrypt(res, d)
	} else {
		block.Decrypt(res, d)
	}

	return res
}

func DESL(k []byte, d []byte) []byte {

	k = PAD(PAD(k, 16), 21)
	ret := make([]byte, 24)
	copy(ret, DES(k[:7], d))
	copy(ret[8:], DES(k[7:14], d))
	copy(ret[16:], DES(k[14:21], d))
	return ret

}

func MD4(ms ...any) ([]byte, error) {
	return MakeHash(md4.New(), ms...)
}

func HMACSHA256(k []byte, ms ...any) ([]byte, error) {
	return HMAC(k, sha256.New, ms...)
}

func MD5(ms ...any) ([]byte, error) {
	return MakeHash(md5.New(), ms...)
}

func HMACMD5(k []byte, ms ...any) ([]byte, error) {
	return HMAC(k, md5.New, ms...)
}

func HMAC(k []byte, h func() hash.Hash, ms ...any) ([]byte, error) {
	return MakeHash(hmac.New(h, k), ms...)
}

func RC4K(k []byte, m []byte) error {
	cipher, err := rc4.NewCipher(k)
	if err != nil {
		return err
	}
	cipher.XORKeyStream(m, m)
	return nil
}

func WriteHash(h io.Writer, ms ...any) error {
	for _, m := range ms {
		switch m := m.(type) {
		case [][]byte:
			for _, b := range m {
				if _, err := h.Write(b); err != nil {
					return err
				}
			}
		case []byte:
			if _, err := h.Write(m); err != nil {
				return err
			}
		default:
			if err := binary.Write(h, binary.LittleEndian, m); err != nil {
				return err
			}
		}
	}
	return nil
}

func MakeHash(h hash.Hash, ms ...any) ([]byte, error) {
	if err := WriteHash(h, ms...); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func Nonce(n int) ([]byte, error) {
	var b = make([]byte, n)
	if r, err := rand.Read(b); err != nil {
		return nil, err
	} else if r != n {
		return nil, errors.New("nonce generation number read mismatch.")
	}

	return b, nil
}

func CRC32(m []byte) ([]byte, error) {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, crc32.ChecksumIEEE(m))
	return b, nil
}
