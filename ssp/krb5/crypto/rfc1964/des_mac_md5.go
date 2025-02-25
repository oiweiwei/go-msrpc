package rfc1964

import (
	"crypto/md5"
	"hash"

	ssp_crypto "github.com/oiweiwei/go-msrpc/ssp/crypto"

	"github.com/oiweiwei/gokrb5.fork/v9/crypto"
	"github.com/oiweiwei/gokrb5.fork/v9/crypto/etype"
	"github.com/oiweiwei/gokrb5.fork/v9/types"
)

type DesMacMd5 struct {
	key   []byte
	etype etype.EType
	hash  hash.Hash
}

func NewDesMacMd5(key []byte, etype etype.EType) *DesMacMd5 {
	return &DesMacMd5{key: key, etype: etype, hash: md5.New()}
}

func (h *DesMacMd5) Write(p []byte) (n int, err error) {
	return h.hash.Write(p)
}

func (h *DesMacMd5) Sum(in []byte) ([]byte, error) {

	sum := h.hash.Sum(nil)

	cksum, _, err := h.etype.EncryptData(h.key, sum)
	if err != nil {
		return nil, err
	}

	return append(in, cksum...), nil
}

func DESMACMD5(key types.EncryptionKey, data ...any) ([]byte, error) {

	etype, err := crypto.GetEtype(key.KeyType)
	if err != nil {
		return nil, err
	}

	h := NewDesMacMd5(key.KeyValue, etype)

	if err := ssp_crypto.WriteHash(h, data...); err != nil {
		return nil, err
	}

	return h.Sum(nil)
}
