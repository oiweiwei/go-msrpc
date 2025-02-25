package rfc4757

import (
	"crypto/rc4"
	"errors"

	"github.com/oiweiwei/go-msrpc/ssp/crypto"
)

func DeriveEncryptionKey(key []byte, seqNum []byte) ([]byte, error) {

	if len(seqNum) < 4 {
		return nil, errors.New("sequence number must be at least 4 bytes")
	}

	// for (i = 0; i < 16; i++) Klocal[i] = Kss[i] ^ 0xF0;
	klocal := make([]byte, 16)
	for i := range klocal {
		klocal[i] = key[i] ^ 0xF0
	}
	// Kcrypt = HMAC(Klocal, (int32)0);
	kcrypt, err := crypto.HMACMD5(klocal, int32(0))
	if err != nil {
		return nil, err
	}
	// Kcrypt = HMAC(Kcrypt, (int32)seq);
	kcrypt, err = crypto.HMACMD5(kcrypt, seqNum[:4])
	if err != nil {
		return nil, err
	}

	return kcrypt, nil
}

func DeriveSigningKey(key []byte) ([]byte, error) {
	// Ksign = HMAC(Kss, "signaturekey");
	ksign, err := crypto.HMACMD5(key, ([]byte)("signaturekey\x00"))
	if err != nil {
		return nil, err
	}
	return ksign, nil
}

func NewCipher(key []byte, seqNum []byte) (*rc4.Cipher, error) {

	kcrypt, err := DeriveEncryptionKey(key, seqNum)
	if err != nil {
		return nil, err
	}

	return rc4.NewCipher(kcrypt)
}

func XORSequenceNumber(key []byte, cksum []byte, seqNum []byte) error {
	// Kseq = HMAC(Kss, (int32)0);
	kseq, err := crypto.HMACMD5(key, int32(0))
	if err != nil {
		return err
	}
	// Kseq = HMAC(Kseq, Token.SGN_CKSUM);
	kseq, err = crypto.HMACMD5(kseq, cksum)
	if err != nil {
		return err
	}
	// RC4(Kseq, Token.SND_SEQ);
	if err := crypto.RC4K(kseq, seqNum); err != nil {
		return err
	}
	return nil
}
