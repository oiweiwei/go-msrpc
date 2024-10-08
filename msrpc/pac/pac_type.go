package pac

import (
	"bytes"
	"fmt"

	ndr "github.com/oiweiwei/go-msrpc/ndr"

	claims "github.com/oiweiwei/go-msrpc/msrpc/adts/claims/claims/v1"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
)

type PAC struct {
	Version                          int                       `json:"version"`
	Buffers                          []*PACInfoBuffer          `json:"pac_info_buffer"`
	LogonInformation                 *KerberosValidationInfo   `json:"logon_information,omitempty"`
	CredentialInformation            *PACCredentialInfo        `json:"credential_information,omitempty"`
	ServerChecksum                   *PACSignatureData         `json:"server_checksum,omitempty"`
	KDCChecksum                      *PACSignatureData         `json:"kdc_checksum,omitempty"`
	ClientNameAndTicketInformation   *PACClientInfo            `json:"client_name_and_ticket_information,omitempty"`
	ConstrainedDelegationInformation *S4UDelegationInfo        `json:"constrained_delegation_information,omitempty"`
	UPNAndDNSInformation             *UPNDNSInfo               `json:"upn_and_dns_information,omitempty"`
	ClientClaimsInformation          *claims.ClaimsSetMetadata `json:"client_claims_information,omitempty"`
	DeviceInformation                *PACDeviceInfo            `json:"device_information,omitempty"`
	DeviceClaimsInformation          *claims.ClaimsSetMetadata `json:"device_claims_information,omitempty"`
	TicketChecksum                   *PACSignatureData         `json:"ticket_checksum,omitempty"`
	Attributes                       *PACAttributesInfo        `json:"attributes,omitempty"`
	RequestorSID                     *dtyp.SID                 `json:"requestor_sid,omitempty"`
	ExtendedKDCChecksum              *PACSignatureData         `json:"extended_kdc_checksum,omitempty"`
	RequestorGUID                    *dtyp.GUID                `json:"requestor_guid,omitempty"`
}

func (p *PAC) Unmarshal(b []byte) error {

	var pac PACType

	if err := ndr.Unmarshal(b, &pac, ndr.Opaque); err != nil {
		return fmt.Errorf("unmarshal_pac: headers: %w", err)
	}

	p.Version = int(pac.Version)
	p.Buffers = pac.Buffers

	for _, buffer := range pac.Buffers {

		b := b[buffer.Offset : buffer.Offset+uint64(buffer.BufferLength)]

		if len(b) == 0 {
			continue
		}

		switch buffer.Type {
		case PACInfoBufferTypeLogonInfo:
			var v KerberosValidationInfo
			if err := ndr.UnmarshalWithTypeSerializationV1(b, ndr.UnmarshalerPointer(&v)); err != nil {
				return fmt.Errorf("unmarshal_pac: kerberos_validation_info: %w", err)
			}
			p.LogonInformation = &v
		case PACInfoBufferTypeCredentialsInfo:
			var v PACCredentialInfo
			if err := ndr.Unmarshal(b, &v, ndr.Opaque); err != nil {
				return fmt.Errorf("unmarshal_pac: credential_info: %w", err)
			}
			p.CredentialInformation = &v
		case PACInfoBufferTypeServerChecksum:
			var v PACSignatureData
			if err := ndr.Unmarshal(b, &v, ndr.Opaque); err != nil {
				return fmt.Errorf("unmarhsal_pac: server_checksum: %w", err)
			}
			p.ServerChecksum = &v
		case PACInfoBufferTypeKDCChecksum:
			var v PACSignatureData
			if err := ndr.Unmarshal(b, &v, ndr.Opaque); err != nil {
				return fmt.Errorf("unmarshal_pac: kdc_checksum: %w", err)
			}
			p.KDCChecksum = &v
		case PACInfoBufferTypeClientNameAndTicketInfo:
			var v PACClientInfo
			if err := ndr.Unmarshal(b, &v, ndr.Opaque); err != nil {
				return fmt.Errorf("unmarshal_pac: client_name_and_ticket_information: %w", err)
			}
			p.ClientNameAndTicketInformation = &v
		case PACInfoBufferTypeConstrainedDelegationInfo:
			var v S4UDelegationInfo
			if err := ndr.UnmarshalWithTypeSerializationV1(b, ndr.UnmarshalerPointer(&v)); err != nil {
				return fmt.Errorf("unmarshal_pac: constrained_delegation_information: %w", err)
			}
			p.ConstrainedDelegationInformation = &v
		case PACInfoBufferTypeUPNAndDNSInfo:
			var v UPNDNSInfo
			if err := v.Unmarshal(b); err != nil {
				return fmt.Errorf("unmarshal_pac: upn_and_dns_information: %w", err)
			}
			p.UPNAndDNSInformation = &v
		case PACInfoBufferTypeClientClaimsInfo:
			var v claims.ClaimsSetMetadata
			if err := ndr.UnmarshalWithTypeSerializationV1(b, ndr.UnmarshalerPointer(&v)); err != nil {
				return fmt.Errorf("unmarshal_pac: client_claims_information: %w", err)
			}
			p.ClientClaimsInformation = &v
		case PACInfoBufferTypeDeviceInfo:
			var v PACDeviceInfo
			if err := ndr.UnmarshalWithTypeSerializationV1(b, ndr.UnmarshalerPointer(&v)); err != nil {
				return fmt.Errorf("unmarshal_pac: device_information: %w", err)
			}
			p.DeviceInformation = &v
		case PACInfoBufferTypeDeviceClaimsInfo:
			var v claims.ClaimsSetMetadata
			if err := ndr.UnmarshalWithTypeSerializationV1(b, ndr.UnmarshalerPointer(&v)); err != nil {
				return fmt.Errorf("unmarshal_pac: client_claims_information: %w", err)
			}
			p.DeviceClaimsInformation = &v
		case PACInfoBufferTypeTicketChecksum:
			var v PACSignatureData
			if err := ndr.Unmarshal(b, &v, ndr.Opaque); err != nil {
				return fmt.Errorf("unmarshal_pac: ticket_checksum: %w", err)
			}
			p.TicketChecksum = &v
		case PACInfoBufferTypeAttributes:
			// XXX: the format of the PAC_ATTRIBUTE_INFO is defined as uint32 flag-size in bits
			// and array of flags of rounded (flag-size / 32)
			var v PACAttributesInfo
			if err := ndr.Unmarshal(b, &v, ndr.Opaque); err != nil {
				return fmt.Errorf("unmarshal_pac: attributes: %w", err)
			}
			p.Attributes = &v
		case PACInfoBufferTypeRequestorSID:
			var v dtyp.SID
			if err := ndr.Unmarshal(b, &v, ndr.Opaque); err != nil {
				return fmt.Errorf("unmarshal_pac: requestor_sid: %w", err)
			}
			p.RequestorSID = &v
		case PACInfoBufferTypeExtendedKDCChecksum:
			var v PACSignatureData
			if err := ndr.Unmarshal(b, &v, ndr.Opaque); err != nil {
				return fmt.Errorf("unmarshal_pac: extended_kdc_checksum: %w", err)
			}
			p.ExtendedKDCChecksum = &v
		case PACInfoBufferTypeRequestorGUID:
			var v dtyp.GUID
			if err := ndr.Unmarshal(b, &v, ndr.Opaque); err != nil {
				return fmt.Errorf("unmarshal_pac: requestor_guid: %w", err)
			}
			p.RequestorGUID = &v
		}
	}

	return nil
}

func IsSignatureBuffer(buffer *PACInfoBuffer) bool {
	if buffer != nil {
		return buffer.Type == 0x00000006 || buffer.Type == 0x00000007 || buffer.Type == 0x00000010 || buffer.Type == 0x00000013
	}

	return false
}

// FillInSignatureData function fills the signature data in the buffer.
func FillInSignatureData(b []byte, sign *PACInfoBuffer, data []byte) ([]byte, error) {

	if !IsSignatureBuffer(sign) {
		return b, nil
	}

	if len(b) < int(sign.Offset+uint64(sign.BufferLength)) {
		return nil, fmt.Errorf("fill_signature_data: buffer too small")
	}

	if sign.BufferLength < 4 {
		return nil, fmt.Errorf("fill_signature_data: buffer too small")
	}

	if n := copy(b[sign.Offset+4:sign.Offset+uint64(sign.BufferLength)], data); n != len(data) {
		return nil, fmt.Errorf("fill_signature_data: short write")
	}

	return b, nil
}

// ZeroOutSignatureData function clears the signature data in the buffer.
func ZeroOutSignatureData(b []byte, sign *PACInfoBuffer) ([]byte, error) {

	if !IsSignatureBuffer(sign) {
		return b, nil
	}

	if len(b) < int(sign.Offset+uint64(sign.BufferLength)) {
		return nil, fmt.Errorf("zero_out_signature_data: buffer too small")
	}

	if sign.BufferLength < 4 {
		return nil, fmt.Errorf("zero_out_signature_data: buffer too small")
	}

	clear(b[sign.Offset+4 : sign.Offset+uint64(sign.BufferLength)])
	return b, nil
}

// Marshal function marshals the PAC structure into a byte array also setting the
// PAC.Buffers field to the Buffers slice. This buffer slice will contain the
// information about the binary layout.
func (p *PAC) Marshal() ([]byte, error) {

	buf := new(bytes.Buffer)

	var pac PACType

	pac.Version = uint32(p.Version)

	// LogonInformation
	if p.LogonInformation != nil {
		b, err := ndr.MarshalWithTypeSerializationV1(ndr.MarshalerPointer(p.LogonInformation))
		if err != nil {
			return nil, fmt.Errorf("marshal_pac: logon_information: %w", err)
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x00000001, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// CredentialInformation
	if p.CredentialInformation != nil {
		b, err := ndr.Marshal(p.CredentialInformation, ndr.Opaque)
		if err != nil {
			return nil, fmt.Errorf("marshal_pac: credential_info: %w", err)
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x00000002, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// ServerChecksum
	if p.ServerChecksum != nil {
		b, err := ndr.Marshal(p.ServerChecksum, ndr.Opaque)
		if err != nil {
			return nil, fmt.Errorf("marshal_pac: server_checksum: %w", err)
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x00000006, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// KDCChecksum
	if p.KDCChecksum != nil {
		b, err := ndr.Marshal(p.KDCChecksum, ndr.Opaque)
		if err != nil {
			return nil, fmt.Errorf("marshal_pac: kdc_checksum: %w", err)
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x00000007, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// ClientNameAndTicketInformation
	if p.ClientNameAndTicketInformation != nil {
		b, err := ndr.Marshal(p.ClientNameAndTicketInformation, ndr.Opaque)
		if err != nil {
			return nil, fmt.Errorf("marshal_pac: client_name_and_ticket_information: %w", err)
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x0000000A, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// ConstrainedDelegationInformation
	if p.ConstrainedDelegationInformation != nil {
		b, err := ndr.MarshalWithTypeSerializationV1(ndr.MarshalerPointer(p.ConstrainedDelegationInformation))
		if err != nil {
			return nil, fmt.Errorf("marshal_pac: constrained_delegation_information: %w", err)
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x0000000B, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// UPNAndDNSInformation
	if p.UPNAndDNSInformation != nil {
		b, err := p.UPNAndDNSInformation.Marshal()
		if err != nil {
			return nil, fmt.Errorf("marshal_pac: upn_and_dns_information: %w", err)
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x0000000C, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// ClientClaimsInformation
	if p.ClientClaimsInformation != nil {
		b, err := ndr.MarshalWithTypeSerializationV1(ndr.MarshalerPointer(p.ClientClaimsInformation))
		if err != nil {
			return nil, fmt.Errorf("marshal_pac: client_claims_information: %w", err)
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x0000000D, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// DeviceInformation
	if p.DeviceInformation != nil {
		b, err := ndr.MarshalWithTypeSerializationV1(ndr.MarshalerPointer(p.DeviceInformation))
		if err != nil {
			return nil, fmt.Errorf("marshal_pac: device_information: %w", err)
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x0000000E, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// DeviceClaimsInformation
	if p.DeviceClaimsInformation != nil {
		b, err := ndr.MarshalWithTypeSerializationV1(ndr.MarshalerPointer(p.DeviceClaimsInformation))
		if err != nil {
			return nil, fmt.Errorf("marshal_pac: device_claims_information: %w", err)
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x0000000F, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// TicketChecksum
	if p.TicketChecksum != nil {
		b, err := ndr.Marshal(p.TicketChecksum, ndr.Opaque)
		if err != nil {
			return nil, fmt.Errorf("marshal_pac: ticket_checksum: %w", err)
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x00000010, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// Attributes.
	if p.Attributes != nil {
		b, err := ndr.Marshal(p.Attributes, ndr.Opaque)
		if err != nil {
			return nil, fmt.Errorf("marshal_pac: attributes: %w", err)
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x00000011, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// RequestorSID
	if p.RequestorSID != nil {
		b, err := ndr.Marshal(p.RequestorSID, ndr.Opaque)
		if err != nil {
			return nil, fmt.Errorf("marshal_pac: requestor_sid: %w", err)
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x00000012, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// ExtendedKDCChecksum
	if p.ExtendedKDCChecksum != nil {
		b, err := ndr.Marshal(p.ExtendedKDCChecksum, ndr.Opaque)
		if err != nil {
			return nil, fmt.Errorf("marshal_pac: extended_kdc_checksum: %w", err)
		}
		pac.Buffers = append(pac.Buffers, &PACInfoBuffer{Type: 0x00000013, Offset: uint64(buf.Len()), BufferLength: uint32(len(b))})
		writeWithPad(buf, b)
	}

	// RequestorGUID
	if p.RequestorGUID != nil {
		b, err := ndr.Marshal(p.RequestorGUID, ndr.Opaque)
		if err != nil {
			return nil, fmt.Errorf("marshal_pac: requestor_guid: %w", err)
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

	p.Buffers = pac.Buffers

	return append(b, buf.Bytes()...), nil
}

func writeWithPad(buf *bytes.Buffer, b []byte) {
	buf.Write(b)
	if pad := len(b) % 8; pad != 0 {
		buf.Write(make([]byte, 8-pad))
	}
}
