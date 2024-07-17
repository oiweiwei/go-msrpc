package netlogon

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/oiweiwei/go-msrpc/ssp/crypto"
)

type SecureCredential struct {
	caps Cap
	key  []byte
	cred []byte
}

func NewSecureCredential(ctx context.Context, cfg *Config) (*SecureCredential, error) {

	key, err := ComputeSessionKey(ctx, cfg.Capabilities, cfg.Credential, cfg.ClientChallenge, cfg.ServerChallenge)
	if err != nil {
		return nil, fmt.Errorf("secure_credentials: %v", err)
	}

	sc := &SecureCredential{caps: cfg.Capabilities, key: key}

	if sc.cred, err = sc.Encrypt(ctx, cfg.ClientChallenge[:8]); err != nil {
		return nil, fmt.Errorf("secure_credentials: %v", err)
	}

	return sc, nil
}

func (a *SecureCredential) Next(ctx context.Context, inc uint32) ([]byte, error) {
	// extract credentials uint32.
	sC := binary.LittleEndian.Uint32(a.cred)
	// add credentials to the increment.
	binary.LittleEndian.PutUint32(a.cred, sC+inc)
	// encrypt the credentials.
	return a.Encrypt(ctx, a.cred)
}

func (a *SecureCredential) Verify(ctx context.Context, inc uint32, cred []byte) error {

	expCred, err := a.Next(ctx, inc)
	if err != nil {
		return fmt.Errorf("verify_credentials: %v", err)
	}

	if !bytes.Equal(cred, expCred) {
		return fmt.Errorf("verify_credentials: credential mismatch: expected %s got %s",
			hex.EncodeToString(expCred), hex.EncodeToString(cred))
	}

	return nil
}

func (a *SecureCredential) Encrypt(ctx context.Context, cred []byte) ([]byte, error) {

	if len(a.key) < 14 {
		return nil, fmt.Errorf("compute_credential: invalid session key")
	}

	if a.caps.IsSet(CapAES_SHA2) {
		return crypto.AES_CFB(a.key, make([]byte, 16), cred, false), nil
	}

	return crypto.DES_ECB(a.key[7:14], crypto.DES_ECB(a.key[:7], cred, true), true), nil
}
