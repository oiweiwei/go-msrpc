package rfc4757

import "github.com/oiweiwei/go-msrpc/ssp/crypto"

const (
	KeySaltWrap = int32(13)
	KeySaltMIC  = int32(15)
)

func ComputeWrapChecksum(key []byte, data ...any) ([]byte, error) {
	return ComputeChecksum(key, KeySaltWrap, data...)
}

func ComputeMICChecksum(key []byte, data ...any) ([]byte, error) {
	return ComputeChecksum(key, KeySaltMIC, data...)
}

func ComputeChecksum(key []byte, salt int32, data ...any) ([]byte, error) {
	ksign, err := DeriveSigningKey(key)
	if err != nil {
		return nil, err
	}
	// Sgn_Cksum = MD5((int32)salt, Token.Header, data);
	sgnCksum, err := crypto.MD5(append([]any{salt}, data...)...)
	if err != nil {
		return nil, err
	}
	// Sgn_Cksum = HMAC(Ksign, Sgn_Cksum);
	sgnCksum, err = crypto.HMACMD5(ksign, sgnCksum)
	if err != nil {
		return nil, err
	}
	return sgnCksum, nil
}
