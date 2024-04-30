package netlogon

import (
	"context"
	"encoding/binary"
	"fmt"

	"github.com/oiweiwei/go-msrpc/ssp/crypto"
	"github.com/oiweiwei/go-msrpc/text/encoding/utf16le"
)

func DeriveKey(ctx context.Context, cred Credential) ([]byte, error) {

	b, err := utf16le.Encode(cred.Password())
	if err != nil {
		return nil, err
	}

	return crypto.MD4(b)
}

func ComputeSessionKey(ctx context.Context, caps Cap, cred Credential, client, server []byte) ([]byte, error) {

	pwd, err := DeriveKey(ctx, cred)
	if err != nil {
		return nil, fmt.Errorf("compute_session_key: derive_key: %v", err)
	}

	if caps.IsSet(CapAES_SHA2) {
		return crypto.HMAC_SHA256(pwd, client, server)[:16], nil
	}

	if caps.IsSet(CapStrongKey) {
		tmp, err := crypto.MD5(make([]byte, 4), client, server)
		if err != nil {
			return nil, fmt.Errorf("compute_session_key: strong_key: %v", err)
		}
		k, err := crypto.HMAC_MD5(pwd, tmp)
		if err != nil {
			return nil, fmt.Errorf("compute_session_key: strong_key: %v", err)
		}
		return k[:16], nil
	}

	uC := binary.LittleEndian.Uint64(client)
	uS := binary.LittleEndian.Uint64(server)
	uCS := binary.LittleEndian.AppendUint64(nil, uC+uS)

	k := make([]byte, 16)
	copy(k[8:], crypto.DES(crypto.DES(uCS, pwd[:7]), pwd[7:14]))

	return k, nil
}
