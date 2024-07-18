package drsuapi

import (
	"context"
	"crypto/md5"
	"crypto/rc4"
	"fmt"

	"github.com/oiweiwei/go-msrpc/ndr"
	"github.com/oiweiwei/go-msrpc/ssp/crypto"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
)

// DecryptData decrypts the data using the session key from the context.
func DecryptData(ctx context.Context, b []byte) ([]byte, error) {
	key, ok := gssapi.GetAttribute(ctx, gssapi.AttributeSessionKey)
	if !ok {
		return nil, fmt.Errorf("unable to get session key")
	}
	bkey, _ := key.([]byte)
	return DecryptDataWithKey(ctx, bkey, b)
}

// DecryptDataWithKey decrypts the data using the provided key.
func DecryptDataWithKey(ctx context.Context, key []byte, b []byte) ([]byte, error) {

	if len(key) == 0 {
		return nil, fmt.Errorf("unable to get session key")
	}

	payload := EncryptedPayload{}

	if err := ndr.Unmarshal(b, &payload, ndr.Opaque); err != nil {
		return nil, fmt.Errorf("unmarshal encrypted payload: %w", err)
	}

	h := md5.New()
	h.Write(key)
	h.Write(payload.Salt)
	key = h.Sum(nil)

	data := make([]byte, len(payload.EncryptedData)+4)
	copy(data[4:], payload.EncryptedData)

	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("init cipher: %w", err)
	}
	cipher.XORKeyStream(data, data)

	return data[4:], nil
}

// DecryptHash decrypts the hash using the session key and rid.
func DecryptHash(ctx context.Context, rid uint32, b []byte) ([]byte, error) {
	b, err := DecryptData(ctx, b)
	if err != nil {
		return nil, err
	}
	return crypto.DES_ECB_LM(rid, b)
}
