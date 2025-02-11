package crypto

import (
	"context"
	"fmt"

	"github.com/oiweiwei/gokrb5.fork/v9/crypto/etype"
	"github.com/oiweiwei/gokrb5.fork/v9/iana/etypeID"
	"github.com/oiweiwei/gokrb5.fork/v9/types"
)

type Cipher interface {
	Wrap(context.Context, uint64, [][]byte, [][]byte) ([]byte, error)
	Unwrap(context.Context, uint64, [][]byte, [][]byte, []byte) (bool, error)
	MakeSignature(context.Context, uint64, [][]byte) ([]byte, error)
	Size(context.Context, bool) int
}

type CipherSetting struct {
	// Key is an encryption key for the cipher.
	Key types.EncryptionKey
	// Type is an encryption type for the cipher.
	Type etype.EType
	// IsSubKey is true if the key is a subkey.
	IsSubKey bool
	// IsLocal is true if cipher is an initiator's cipher.
	IsLocal bool
	// DCEStyle is true if the cipher is DCE style.
	DCEStyle bool
}

func NewCipher(ctx context.Context, setting CipherSetting) (Cipher, error) {
	switch setting.Type.GetETypeID() {
	case etypeID.AES128_CTS_HMAC_SHA1_96, etypeID.AES256_CTS_HMAC_SHA1_96:
		return NewAESCipher(ctx, setting)
	case etypeID.RC4_HMAC:
		return NewRC4Cipher(ctx, setting)
	case etypeID.DES_CBC_MD5, etypeID.DES_CBC_CRC:
		return NewDESCipher(ctx, setting)
	default:
		return nil, fmt.Errorf("unsupported encryption type %d", setting.Type.GetETypeID())
	}
}
