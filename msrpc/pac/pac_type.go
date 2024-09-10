package pac

import (
	"bytes"
	"fmt"

	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	ndr "github.com/oiweiwei/go-msrpc/ndr"
)

type PAC struct {
	Version                          int                     `json:"version"`
	ZeroSignatureRaw                 []byte                  `json:"-"`
	LogonInformation                 *KerberosValidationInfo `json:"logon_information,omitempty"`
	ServerChecksum                   *PACSignatureData       `json:"server_checksum,omitempty"`
	KDCChecksum                      *PACSignatureData       `json:"kdc_checksum,omitempty"`
	ClientNameAndTicketInformation   *PACClientInfo          `json:"client_name_and_ticket_information,omitempty"`
	ConstrainedDelegationInformation *S4UDelegationInfo      `json:"constrained_delegation_information,omitempty"`
	UPNAndDNSInformation             *UPNDNSInfo             `json:"upn_and_dns_information,omitempty"`
	TicketChecksum                   *PACSignatureData       `json:"ticket_checksum,omitempty"`
	Attributes                       *PACAttributesInfo      `json:"attributes,omitempty"`
	RequestorSID                     *dtyp.SID               `json:"requestor_sid,omitempty"`
	ExtendedKDCChecksum              *PACSignatureData       `json:"extended_kdc_checksum,omitempty"`
	RequestorGUID                    *dtyp.GUID              `json:"requestor_guid,omitempty"`
}

func (p *PAC) Unmarshal(b []byte) error {

	var pac PACType

	if err := ndr.Unmarshal(b, &pac, ndr.Opaque); err != nil {
		return fmt.Errorf("unmarshal_pac: headers: %w", err)
	}

	p.Version = int(pac.Version)

	// allocate data for signature verification.
	p.ZeroSignatureRaw = make([]byte, len(b))
	copy(p.ZeroSignatureRaw, b)

	for _, buffer := range pac.Buffers {

		b := b[buffer.Offset : buffer.Offset+uint64(buffer.BufferLength)]

		if buffer.Type == 0x00000006 || buffer.Type == 0x00000007 || buffer.Type == 0x00000010 || buffer.Type == 0x00000013 {
			// clear the signature data
			clear(p.ZeroSignatureRaw[buffer.Offset : buffer.Offset+uint64(buffer.BufferLength)])
		}

		switch buffer.Type {
		case 0x00000001:
			var v KerberosValidationInfo
			if err := ndr.UnmarshalWithTypeSerializationV1(b, ndr.UnmarshalerPointer(&v)); err != nil {
				return fmt.Errorf("unmarshal_pac: kerberos_validation_info: %w", err)
			}
			p.LogonInformation = &v
		case 0x00000002:
			// TODO: Credentials information
		case 0x00000006:
			var v PACSignatureData
			if err := ndr.Unmarshal(b, &v, ndr.Opaque); err != nil {
				return fmt.Errorf("unmarhsal_pac: server_checksum: %w", err)
			}
			p.ServerChecksum = &v
		case 0x00000007:
			var v PACSignatureData
			if err := ndr.Unmarshal(b, &v, ndr.Opaque); err != nil {
				return fmt.Errorf("unmarshal_pac: kdc_checksum: %w", err)
			}
			p.KDCChecksum = &v
		case 0x0000000A:
			var v PACClientInfo
			if err := ndr.Unmarshal(b, &v, ndr.Opaque); err != nil {
				return fmt.Errorf("unmarshal_pac: client_name_and_ticket_information: %w", err)
			}
			p.ClientNameAndTicketInformation = &v
		case 0x0000000B:
			var v S4UDelegationInfo
			if err := ndr.UnmarshalWithTypeSerializationV1(b, ndr.UnmarshalerPointer(&v)); err != nil {
				return fmt.Errorf("unmarshal_pac: constrained_delegation_information: %w", err)
			}
			p.ConstrainedDelegationInformation = &v
		case 0x0000000C:
			var v UPNDNSInfo
			if err := v.Unmarshal(b); err != nil {
				return fmt.Errorf("unmarshal_pac: upn_and_dns_information: %w", err)
			}
			p.UPNAndDNSInformation = &v
		case 0x0000000E:
			// TODO: Device information
		case 0x0000000F:
			// TODO: Device claims information
		case 0x00000010:
			var v PACSignatureData
			if err := ndr.Unmarshal(b, &v, ndr.Opaque); err != nil {
				return fmt.Errorf("unmarshal_pac: ticket_checksum: %w", err)
			}
			p.TicketChecksum = &v
		case 0x00000011:
			// XXX: the format of the PAC_ATTRIBUTE_INFO is defined as uint32 flag-size in bits
			// and array of flags of rounded (flag-size / 32)
			var v PACAttributesInfo
			if err := ndr.Unmarshal(b, &v, ndr.Opaque); err != nil {
				return fmt.Errorf("unmarshal_pac: attributes: %w", err)
			}
			p.Attributes = &v
		case 0x00000012:
			var v dtyp.SID
			if err := ndr.Unmarshal(b, &v, ndr.Opaque); err != nil {
				return fmt.Errorf("unmarshal_pac: requestor_sid: %w", err)
			}
			p.RequestorSID = &v
		case 0x00000013:
			var v PACSignatureData
			if err := ndr.Unmarshal(b, &v, ndr.Opaque); err != nil {
				return fmt.Errorf("unmarshal_pac: extended_kdc_checksum: %w", err)
			}
			p.ExtendedKDCChecksum = &v
		case 0x00000014:
			var v dtyp.GUID
			if err := ndr.Unmarshal(b, &v, ndr.Opaque); err != nil {
				return fmt.Errorf("unmarshal_pac: requestor_guid: %w", err)
			}
			p.RequestorGUID = &v
		}
	}

	return nil
}

func (p *PAC) Marshal() ([]byte, error) {

	buf := new(bytes.Buffer)

	var pac PACType

	pac.Version = uint32(p.Version)

	// LogonInformation
	if p.LogonInformation != nil {
		b, err := ndr.MarshalWithTypeSerializationV1(ndr.MarshalerPointer(p.LogonInformation))
		if err != nil {
			return nil, err
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x00000001, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// ServerChecksum
	if p.ServerChecksum != nil {
		b, err := ndr.Marshal(p.ServerChecksum, ndr.Opaque)
		if err != nil {
			return nil, err
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x00000006, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// KDCChecksum
	if p.KDCChecksum != nil {
		b, err := ndr.Marshal(p.KDCChecksum, ndr.Opaque)
		if err != nil {
			return nil, err
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x00000007, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// ClientNameAndTicketInformation
	if p.ClientNameAndTicketInformation != nil {
		b, err := ndr.Marshal(p.ClientNameAndTicketInformation, ndr.Opaque)
		if err != nil {
			return nil, err
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x0000000A, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// ConstrainedDelegationInformation
	if p.ConstrainedDelegationInformation != nil {
		b, err := ndr.MarshalWithTypeSerializationV1(p.ConstrainedDelegationInformation)
		if err != nil {
			return nil, err
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x0000000B, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// UPNAndDNSInformation
	if p.UPNAndDNSInformation != nil {
		b, err := p.UPNAndDNSInformation.Marshal()
		if err != nil {
			return nil, err
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x0000000C, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// TicketChecksum
	if p.TicketChecksum != nil {
		b, err := ndr.Marshal(p.TicketChecksum, ndr.Opaque)
		if err != nil {
			return nil, err
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x00000010, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// Attributes.
	if p.Attributes != nil {
		b, err := ndr.Marshal(p.Attributes, ndr.Opaque)
		if err != nil {
			return nil, err
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x00000011, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// RequestorSID
	if p.RequestorSID != nil {
		b, err := ndr.Marshal(p.RequestorSID, ndr.Opaque)
		if err != nil {
			return nil, err
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x00000012, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// ExtendedKDCChecksum
	if p.ExtendedKDCChecksum != nil {
		b, err := ndr.Marshal(p.ExtendedKDCChecksum, ndr.Opaque)
		if err != nil {
			return nil, err
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x00000013, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// RequestorGUID
	if p.RequestorGUID != nil {
		b, err := ndr.Marshal(p.RequestorGUID, ndr.Opaque)
		if err != nil {
			return nil, err
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x00000014, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	for _, buffer := range pac.Buffers {
		buffer.Offset += uint64(len(pac.Buffers)*16 + 8)
	}

	b, err := ndr.Marshal(&pac, ndr.Opaque)
	if err != nil {
		return nil, err
	}

	return append(b, buf.Bytes()...), nil
}

func writeWithPad(buf *bytes.Buffer, b []byte) {
	buf.Write(b)
	if pad := len(b) % 8; pad != 0 {
		buf.Write(make([]byte, 8-pad))
	}
}
