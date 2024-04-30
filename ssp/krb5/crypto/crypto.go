package crypto

import (
	"context"
	"fmt"

	"github.com/jcmturner/gokrb5/v8/crypto"
	"github.com/jcmturner/gokrb5/v8/iana/etypeID"
	"github.com/jcmturner/gokrb5/v8/types"
)

type Cipher interface {
	Wrap(context.Context, uint64, [][]byte, [][]byte) ([]byte, error)
	Unwrap(context.Context, uint64, [][]byte, [][]byte, []byte) (bool, error)
	MakeSignature(context.Context, uint64, [][]byte) ([]byte, error)
	Size(context.Context, bool) int
}

func NewCipher(ctx context.Context, key types.EncryptionKey, isSubKey bool, isServer bool) (Cipher, error) {
	switch etype, _ := crypto.GetEtype(key.KeyType); etype.GetETypeID() {
	case etypeID.AES128_CTS_HMAC_SHA1_96, etypeID.AES256_CTS_HMAC_SHA1_96:
		return NewAESCipher(ctx, key, isServer, isSubKey)
	case etypeID.RC4_HMAC:
		return NewRC4Cipher(ctx, key, isServer)
	default:
		return nil, fmt.Errorf("unsupported encryption type %d", etype.GetETypeID())
	}
}
