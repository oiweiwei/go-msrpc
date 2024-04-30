package krb5

import (
	"fmt"
	"time"

	"github.com/jcmturner/gokrb5/v8/asn1tools"
	"github.com/jcmturner/gokrb5/v8/crypto"
	"github.com/jcmturner/gokrb5/v8/iana/asnAppTag"
	"github.com/jcmturner/gokrb5/v8/iana/keyusage"
	"github.com/jcmturner/gokrb5/v8/messages"
	"github.com/jcmturner/gokrb5/v8/types"

	"github.com/jcmturner/gofork/encoding/asn1"
)

type APReq messages.APReq

func (m *APReq) SequenceNumber() int64 {
	if m != nil {
		return m.Authenticator.SeqNumber
	}
	return 0
}

func (m *APReq) Marshal() ([]byte, error) {
	return (*messages.APReq)(m).Marshal()
}

func (m *APReq) Unmarshal(b []byte) error {
	return (*messages.APReq)(m).Unmarshal(b)
}

func (m *APReq) DecryptAuthenticator(key types.EncryptionKey) error {
	return (*messages.APReq)(m).DecryptAuthenticator(key)
}

type EncAPRepPart messages.EncAPRepPart

func (m *EncAPRepPart) Marshal() ([]byte, error) {
	b, err := asn1.Marshal(*m)
	if err != nil {
		return nil, err
	}
	return asn1tools.AddASNAppTag(b, asnAppTag.EncAPRepPart), nil
}

func (m *EncAPRepPart) Unmarshal(b []byte) error {
	_, err := asn1.UnmarshalWithParams(b, m, fmt.Sprintf("application,explicit,tag:%v", asnAppTag.EncAPRepPart))
	return err
}

type APRep struct {
	messages.APRep
	EncAPRepPart
}

func (m *APRep) SequenceNumber() int64 {
	if m != nil {
		return m.EncAPRepPart.SequenceNumber
	}
	return 0
}

func (m *APRep) Unmarshal(b []byte) error {
	return m.APRep.Unmarshal(b)
}

func (m *APRep) DecryptEncPart(key types.EncryptionKey) error {
	b, err := crypto.DecryptEncPart(m.EncPart, key, keyusage.AP_REP_ENCPART)
	if err != nil {
		return err
	}
	return m.EncAPRepPart.Unmarshal(b)
}

func (m *APRep) EncryptEncPart(e EncAPRepPart, key types.EncryptionKey) error {

	b, err := e.Marshal()
	if err != nil {
		return err
	}

	if m.APRep.EncPart, err = crypto.GetEncryptedData(b, key, keyusage.AP_REP_ENCPART, 0); err != nil {
		return err
	}

	return nil
}

func (m *APRep) MarshalWithSeqNumber(seqNum int64, key types.EncryptionKey) ([]byte, error) {

	t := time.Now().UTC()

	return m.MarshalWithEncPart(EncAPRepPart{
		Cusec:          int((t.UnixNano() / int64(time.Microsecond)) - (t.Unix() * 1e6)),
		CTime:          t,
		SequenceNumber: seqNum,
	}, key)

}

func (m *APRep) MarshalWithEncPart(e EncAPRepPart, key types.EncryptionKey) ([]byte, error) {

	m = &APRep{
		APRep: messages.APRep{
			PVNO:    m.PVNO,
			MsgType: m.MsgType,
		},
		EncAPRepPart: e,
	}

	if err := m.EncryptEncPart(e, key); err != nil {
		return nil, err
	}

	return m.Marshal()
}

func (m *APRep) Marshal() ([]byte, error) {
	b, err := asn1.Marshal(m.APRep)
	if err != nil {
		return nil, err
	}
	return asn1tools.AddASNAppTag(b, asnAppTag.APREP), nil
}
