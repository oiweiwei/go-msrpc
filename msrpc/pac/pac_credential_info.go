package pac

import (
	"errors"
	"fmt"

	"github.com/jcmturner/gokrb5/v8/crypto"
	"github.com/jcmturner/gokrb5/v8/iana/keyusage"
	"github.com/jcmturner/gokrb5/v8/types"
	ndr "github.com/oiweiwei/go-msrpc/ndr"
)

func (o *PACCredentialInfo) DecryptCredentialData(key types.EncryptionKey) (*PACCredentialData, error) {

	if len(o.SerializedData) == 0 {
		return nil, nil
	}

	if int32(o.EncryptionType) != key.KeyType {
		return nil, errors.New("pac: invalid key type")
	}

	b, err := crypto.DecryptMessage(o.SerializedData, key, keyusage.KERB_NON_KERB_SALT)
	if err != nil {
		return nil, fmt.Errorf("pac: decrypt credential data: %w", err)
	}

	var pkg PACCredentialData

	if err := ndr.UnmarshalWithTypeSerializationV1(b, ndr.UnmarshalerPointer(&pkg)); err != nil {
		return nil, fmt.Errorf("pac: unmarshal credential data: %w", err)
	}

	for _, cred := range pkg.Credentials {
		if cred.PackageName != nil && cred.PackageName.Buffer == "NTLM" {

			if cred.NTLMSupplementalCredential == nil {
				cred.NTLMSupplementalCredential = new(NTLMSupplementalCredential)
			}

			if err = ndr.Unmarshal([]byte(cred.Credentials), cred.NTLMSupplementalCredential, ndr.Opaque); err != nil {
				return nil, fmt.Errorf("pac: unmarshal ntlm supplemental credentials: %w", err)
			}

			if cred.NTLMSupplementalCredential.Flags&0x01 == 0 {
				cred.NTLMSupplementalCredential.LMPassword = nil
			}

			if cred.NTLMSupplementalCredential.Flags&0x02 == 0 {
				cred.NTLMSupplementalCredential.NTPassword = nil
			}
		}
	}

	return &pkg, nil
}
