package rfc1964

import (
	"bytes"
	"fmt"

	"github.com/oiweiwei/gokrb5.fork/v9/crypto"
	"github.com/oiweiwei/gokrb5.fork/v9/crypto/rfc3961"
	"github.com/oiweiwei/gokrb5.fork/v9/types"
)

type DESCipher struct {
	key     types.EncryptionKey
	buf     bytes.Buffer
	buffers [][]byte
}

func NewCipher(key types.EncryptionKey) *DESCipher {
	return &DESCipher{key: key}
}

func Encrypt(key types.EncryptionKey, data ...[]byte) error {
	h := NewCipher(key)
	for _, d := range data {
		h.AddBuffer(d)
	}
	return h.Encrypt()
}

func Decrypt(key types.EncryptionKey, data ...[]byte) error {
	h := NewCipher(key)
	for _, d := range data {
		h.AddBuffer(d)
	}
	return h.Decrypt()
}

func (e *DESCipher) AddBuffer(b []byte) {
	e.buf.Write(b)
	e.buffers = append(e.buffers, b)
}

func (e *DESCipher) DeriveKey() []byte {
	key := make([]byte, len(e.key.KeyValue))
	for i := range key {
		key[i] = e.key.KeyValue[i] ^ 0xF0
	}
	return key
}

func (e *DESCipher) Encrypt() error {
	return e.applyCipher(true)
}

func (e *DESCipher) Decrypt() error {
	return e.applyCipher(false)
}

func (e *DESCipher) applyCipher(encrypt bool) error {

	etype, err := crypto.GetEtype(e.key.KeyType)
	if err != nil {
		return fmt.Errorf("get etype: %w", err)
	}

	var b []byte

	if encrypt {
		if _, b, err = etype.EncryptData(e.DeriveKey(), e.buf.Bytes()); err != nil {
			return fmt.Errorf("encrypt buffer: %w", err)
		}
	} else {
		if b, err = etype.DecryptData(e.DeriveKey(), e.buf.Bytes()); err != nil {
			return fmt.Errorf("decrypt buffer: %w", err)
		}
	}

	for _, data := range e.buffers {
		b = b[copy(data, b):]
	}

	if len(b) > 0 {
		return fmt.Errorf("unexpected data left after encryption")
	}

	return nil

}

func EncryptSequenceNumber(key types.EncryptionKey, cksum []byte, seqNum []byte) error {
	etype, err := crypto.GetEtype(key.KeyType)
	if err != nil {
		return fmt.Errorf("encrypt sequence number: get etype: %w", err)
	}
	_, seq, err := rfc3961.DESEncryptData(key.KeyValue, seqNum, cksum, etype)
	if err != nil {
		return fmt.Errorf("encrypt sequence number: %w", err)
	}
	copy(seqNum, seq)
	return nil
}

func DecryptSequenceNumber(key types.EncryptionKey, cksum []byte, seqNum []byte) error {
	etype, err := crypto.GetEtype(key.KeyType)
	if err != nil {
		return fmt.Errorf("decrypt sequence number: get etype: %w", err)
	}
	seq, err := rfc3961.DESDecryptData(key.KeyValue, seqNum, cksum, etype)
	if err != nil {
		return fmt.Errorf("decrypt sequence number: %X: %w", seqNum, err)
	}
	copy(seqNum, seq)
	return nil
}
