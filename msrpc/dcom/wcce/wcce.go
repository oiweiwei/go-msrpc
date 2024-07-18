// The wcce package implements the WCCE client protocol.
//
// # Introduction
//
// The Windows Client Certificate Enrollment Protocol consists of a set of DCOM interfaces
// (as specified in [MS-DCOM]) that allow clients to request various services from a
// certification authority (CA). These services enable X.509 (as specified in [X509])
// digital certificate enrollment, issuance, revocation, and property retrieval.
//
// Active Directory can be used to store domain policies for certificate enrollment.
// An implementation of the protocol that is specified in this document might retrieve
// Active Directory objects (1) and attributes that define these enrollment policies.
// Because Active Directory is an independent component with its own protocols, the
// exact process for Active Directory discovery and objects retrieval is covered in
// [MS-ADTS].
//
// Familiarity with public key infrastructure (PKI) concepts such as asymmetric and
// symmetric cryptography, digital certificates, and cryptographic key exchange is required
// for a complete understanding of this specification. In addition, a comprehensive
// understanding of the [X509] standard is required for a complete understanding of
// the protocol and its usage. For a comprehensive introduction to cryptography and
// PKI concepts, see [SCHNEIER]. PKI basics and certificate concepts are as specified
// in [X509]. For an introduction to certificate revocation lists (CRLs) and revocation
// concepts, see [MSFT-CRL].
//
// # Overview
//
// The Windows Client Certificate Enrollment Protocol is built from two DCOM interfaces:
// ICertRequestD and ICertRequestD2, successive versions. The two DCOM interfaces allow
// a client to interact with a CA to request a certificate and to obtain certain information
// about the CA. This document specifies the protocol, the Windows Client Certificate
// Enrollment Protocol, but also specifies certain elements of the behavior of the client
// and the CA (the server), because those behaviors are reflected in or influence protocol
// behavior.
//
// The Windows Client Certificate Enrollment Protocol occurs between one client and
// one server. However, the client and the server are subject to variation, so the enrollment
// process can appear very complex. Other machines and services can also interact with
// the client and/or the server during enrollment, but those interactions depend on
// the particular variations in use.
//
// Two elements of a server are subject to variation. These elements are independent
// of each other and independent of the implementation of the Windows Client Certificate
// Enrollment Protocol stack. This protocol specification refers to these elements as
// follows:
//
// * CA policy algorithm ( 719b890d-62e6-4322-b9b1-1f34d11535b4#gt_cd0e4dab-0331-4123-a538-df8e4e626a71
// )
//
// This algorithm determines 1) whether to issue the certificate requested, and 2) how
// to populate the fields of a certificate that is issued.
//
// * CA exit algorithm ( 719b890d-62e6-4322-b9b1-1f34d11535b4#gt_240a9746-c99e-4765-a9ee-6d60f1a9ffd1
// )
//
// The optional algorithm that is invoked when a certificate is issued. This algorithm
// might store a copy of that certificate in one or more repositories, or the algorithm
// might make a log entry or notify some person of the issuance ( 719b890d-62e6-4322-b9b1-1f34d11535b4#gt_71f399e7-7026-46bb-b7c2-8fd4872b900f
// ) of the certificate.
//
// The variants of interest in the CA policy algorithm are as follows:
//
// * Hard-coded
//
// A policy algorithm that performs the same operation on certificate requests regardless
// of the information specified in the request is called a hard-coded policy algorithm.
// A simple, hard-coded policy algorithm might issue any certificate that is requested.
//
// * Manual
//
// A policy algorithm that requires human intervention in order to determine whether
// or not to issue a certificate is called a manual policy algorithm. A simple manual
// policy algorithm accepts the requester's choice of certificate fields, presents the
// requested certificate to an administrator, and asks the administrator whether or
// not to issue the certificate.
//
// * Policy-driven via certificate templates ( 719b890d-62e6-4322-b9b1-1f34d11535b4#gt_9c47ca85-9343-4e01-96d4-53d96d3df60e
// )
//
// A policy algorithm that determines whether or not to issue certificates based on
// enrollment policies specified in a certificate template [MS-CRTD] ( ../ms-crtd/4c6950e4-1dc2-4ae3-98c3-b8919bb73822
// ). Each certificate template in a collection of certificate templates describes a
// kind of certificate with its fields. The security descriptor on the certificate template
// provides an access control list (ACL) ( 719b890d-62e6-4322-b9b1-1f34d11535b4#gt_9f92aa05-dd0a-45f2-88d6-89f1fb654395
// ) that can include the Enroll permission for an individual or, more typically, a
// group of individuals. A policy algorithm that strictly implements a policy stored
// as certificate templates is described in section 3.2.2.6.2.1.4 ( e8e51249-b699-4004-97de-cb8cbe2c4c9c
// ).
//
// *Note* The capability to base certificate policy on user types is not available for
// a standalone CA ( 719b890d-62e6-4322-b9b1-1f34d11535b4#gt_6593a312-1130-482c-aa85-6840f7b1859f
// ) since standalone CAs do not support the use of certificate templates.
//
// One aspect of a client subject to variation is whether certificate templates are
// used to form certificate requests.
package wcce

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	ndr "github.com/oiweiwei/go-msrpc/ndr"
)

var (
	_ = context.Background
	_ = fmt.Errorf
	_ = utf16.Encode
	_ = strings.TrimPrefix
	_ = ndr.ZeroString
	_ = (*uuid.UUID)(nil)
	_ = (*dcerpc.SyntaxID)(nil)
	_ = (*errors.Error)(nil)
	_ = dcom.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wcce"
)

// CertTransportBlob structure represents CERTTRANSBLOB RPC structure.
//
// The CERTTRANSBLOB structure defines a byte buffer that is used to store certificates,
// request certificates, transmit responses, manipulate [UNICODE] strings, and marshal
// property values.
type CertTransportBlob struct {
	// cb:  Unsigned integer value that MUST contain the length of the buffer pointed to
	// by pb in bytes.
	Length uint32 `idl:"name:cb" json:"length"`
	// pb:  Byte buffer that MUST contain the binary contents being transported in this
	// CERTTRANSBLOB.
	//
	// CERTTRANSBLOB is empty when both cb and pb are set to 0.
	//
	// The following sections specify marshaling of all supported structures that can be
	// passed in the pb Byte buffer of CERTTRANSBLOB.
	//
	// All instances of CERTTRANSBLOB used by this protocol MUST use one of the marshaling
	// rules described in the following sections.
	Buffer []byte `idl:"name:pb;size_is:(cb);pointer:unique" json:"buffer"`
}

func (o *CertTransportBlob) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.Length == 0 {
		o.Length = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CertTransportBlob) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if o.Buffer != nil || o.Length > 0 {
		_ptr_pb := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Length)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Buffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Buffer[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Buffer, _ptr_pb); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CertTransportBlob) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	_ptr_pb := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Length > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Length)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]byte, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if err := w.ReadData(&o.Buffer[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pb := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Buffer, _s_pb, _ptr_pb); err != nil {
		return err
	}
	return nil
}

// CATransportProperty structure represents CATRANSPROP RPC structure.
//
// The CATRANSPROP structure encapsulates information about a CA property. For a list
// of CA properties, see section 3.2.1.4.3.2. An array of these structures is carried
// in a CERTTRANSBLOB (section 2.2.2.2) structure, and is returned by GetCAPropertyInfo,
// as specified in section 3.2.1.4.3.3. Note that this structure does not contain property
// values themselves; rather, CATRANSPROP contains information about properties.
//
// A CERTTRANSBLOB (section 2.2.2.2) structure MUST be used to return an array of CATRANSPROP
// (section 2.2.2.3) structures, where the count of array elements is returned in a
// separate output parameter of the remote procedure call. It MUST also contain a null-terminated
// Unicode string for each CATRANSPROP (section 2.2.2.3) structure that represents the
// display name of the CA property.
//
// The following tables show the sequence of fields in the byte array referenced by
// the pb field of the CERTTRANSBLOB (section 2.2.2.2) structure when used to transfer
// an array of CATRANSPROP (section 2.2.2.3) structures and their corresponding data.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| CATRANSPROP Structures (variable)                                                                                             |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Byte Array (variable)                                                                                                         |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type CATransportProperty struct {
	// lPropID:  Integer value that MUST contain the property identifier. For the list of
	// supported properties, see section 3.2.1.4.3.2.
	PropertyID int32 `idl:"name:lPropID" json:"property_id"`
	// propType:  Byte value that MUST contain the data type for the property. Must be one
	// of the following values.
	//
	//	+---------------------+-----------------------------------------+
	//	|                     |                                         |
	//	|        VALUE        |                 MEANING                 |
	//	|                     |                                         |
	//	+---------------------+-----------------------------------------+
	//	+---------------------+-----------------------------------------+
	//	| PROPTYPE_LONG 0x1   | Property type is a signed long integer. |
	//	+---------------------+-----------------------------------------+
	//	| PROPTYPE_DATE 0x2   | Property type is a date-time value.     |
	//	+---------------------+-----------------------------------------+
	//	| PROPTYPE_BINARY 0x3 | Property type is binary data.           |
	//	+---------------------+-----------------------------------------+
	//	| PROPTYPE_STRING 0x4 | Property type is a string.              |
	//	+---------------------+-----------------------------------------+
	PropertyType uint8 `idl:"name:propType" json:"property_type"`
	// Reserved:  MUST be set to 0 and ignored upon receipt.
	_ uint8 `idl:"name:Reserved"`
	// propFlags:  16-bit flag field.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	|   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |
	//	|   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | I |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	// Where the bits are defined as:
	//
	//	+-------+----------------------------------------------------------------------------------+
	//	|       |                                                                                  |
	//	| VALUE |                                   DESCRIPTION                                    |
	//	|       |                                                                                  |
	//	+-------+----------------------------------------------------------------------------------+
	//	+-------+----------------------------------------------------------------------------------+
	//	| I     | This bit provides indication that the property is indexed and has multiple       |
	//	|       | values. If this bit is set to 1, then a property is indexed. If the bit is set   |
	//	|       | to 0, then the property is not indexed.                                          |
	//	+-------+----------------------------------------------------------------------------------+
	PropertyFlags uint16 `idl:"name:propFlags" json:"property_flags"`
	// obwszDisplayName:  Integer that MUST contain the offset to the string that contains
	// the display name of this property, where the offset begins at the beginning of the
	// byte array referenced by the pb field of the containing CERTTRANSBLOB (section 2.2.2.2)
	// structure. The string format MUST be null-terminated [UNICODE]. The offset MUST be
	// DWORD-aligned. For marshaling information about this property, see Marshaling CATRANSPROP
	// in a CERTTRANSBLOB (section 2.2.2.3.1).
	DisplayNameOffset uint32 `idl:"name:obwszDisplayName" json:"display_name_offset"`
}

func (o *CATransportProperty) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CATransportProperty) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.PropertyID); err != nil {
		return err
	}
	if err := w.WriteData(o.PropertyType); err != nil {
		return err
	}
	// reserved Reserved
	if err := w.WriteData(uint8(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.PropertyFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.DisplayNameOffset); err != nil {
		return err
	}
	return nil
}
func (o *CATransportProperty) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyID); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyType); err != nil {
		return err
	}
	// reserved Reserved
	var _Reserved uint8
	if err := w.ReadData(&_Reserved); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.DisplayNameOffset); err != nil {
		return err
	}
	return nil
}

// CAInfo structure represents CAINFO RPC structure.
//
// The CAINFO structure defines a basic informational block that describes a CA.
type CAInfo struct {
	// cbSize:  Unsigned integer value that MUST contain the size of this structure in bytes.
	Length uint32 `idl:"name:cbSize" json:"length"`
	// CAType:  Integer value that SHOULD contain a constant describing the CA type. The
	// value SHOULD be one of the values in the following table.
	//
	// Note  The value 0x00000002 MUST NOT be used for this parameter.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|                                   |                                                                                  |
	//	|               VALUE               |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| ENUM_ENTERPRISE_ROOTCA 0x00000000 | The CA is an enterprise root (self-signed) CA. For more information, see         |
	//	|                                   | [MSFT-PKI].                                                                      |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| ENUM_ENTERPRISE_SUBCA 0x00000001  | The CA is an enterprise subordinate CA. For more information, see [MSFT-PKI].    |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| ENUM_STANDALONE_ROOTCA 0x00000003 | The CA is a stand-alone root (self-signed) CA. For more information, see         |
	//	|                                   | [MSFT-PKI].                                                                      |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| ENUM_STANDALONE_SUBCA 0x00000004  | The CA is a stand-alone subordinate CA. For more information, see [MSFT-PKI].    |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| ENUM_UNKNOWN_CA 0x00000005        | The CA type is unknown.                                                          |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	CAType int32 `idl:"name:CAType" json:"ca_type"`
	// cCASignatureCerts:  Unsigned integer value that SHOULD contain the count of CA signing
	// certificates in the CA. A CA signing certificate contains a public key that is in
	// turn associated with the private key used to sign certificates that are issued by
	// the CA. For more information on CA signing certificates, see [MSFT-PKI].
	CASignatureCertsCount uint32 `idl:"name:cCASignatureCerts" json:"ca_signature_certs_count"`
	// cCAExchangeCerts:  Unsigned integer value that SHOULD contain the count of CA exchange
	// certificates in the CA. CA exchange certificates contain public keys that are used
	// to encrypt requests sent to a CA. For more information, see [MSFT-ARCHIVE].
	CAExchangeCertsCount uint32 `idl:"name:cCAExchangeCerts" json:"ca_exchange_certs_count"`
	// cExitAlgorithms:  Unsigned integer value that SHOULD contain the number of exit algorithms
	// that are installed and active for the CA.
	ExitAlgorithmsCount uint32 `idl:"name:cExitAlgorithms" json:"exit_algorithms_count"`
	// lPropIDMax:  Integer that SHOULD contain the maximum supported value for the PropID
	// parameter in the ICertRequestD2::GetCAProperty method. For more information on CA
	// properties, see section 3.2.1.4.3.2.
	PropertyIDMax int32 `idl:"name:lPropIDMax" json:"property_id_max"`
	// lRoleSeparationEnabled:  Integer value that SHOULD indicate whether CA role separation
	// is enabled on the CA. A value of 0 indicates that CA role separation is disabled;
	// a value of 1 indicates that it is enabled.
	RoleSeparationEnabled int32 `idl:"name:lRoleSeparationEnabled" json:"role_separation_enabled"`
	// cKRACertUsedCount:  Unsigned integer value that SHOULD contain the number of key
	// recovery agent (KRA) keys used to encrypt each archived private key.
	KRACertUsedCount uint32 `idl:"name:cKRACertUsedCount" json:"kra_cert_used_count"`
	// cKRACertCount:  Unsigned integer value that SHOULD contain the number of KRA keys
	// available for the CA to encrypt archived private keys.
	KRACertCount uint32 `idl:"name:cKRACertCount" json:"kra_cert_count"`
	// fAdvancedServer:  Unsigned integer value that SHOULD be set to 0 for standard CA
	// and 1 for advanced CA. This value is a Boolean value. The CA SHOULD return 0 or 1.
	AdvancedServer uint32 `idl:"name:fAdvancedServer" json:"advanced_server"`
}

func (o *CAInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CAInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.CAType); err != nil {
		return err
	}
	if err := w.WriteData(o.CASignatureCertsCount); err != nil {
		return err
	}
	if err := w.WriteData(o.CAExchangeCertsCount); err != nil {
		return err
	}
	if err := w.WriteData(o.ExitAlgorithmsCount); err != nil {
		return err
	}
	if err := w.WriteData(o.PropertyIDMax); err != nil {
		return err
	}
	if err := w.WriteData(o.RoleSeparationEnabled); err != nil {
		return err
	}
	if err := w.WriteData(o.KRACertUsedCount); err != nil {
		return err
	}
	if err := w.WriteData(o.KRACertCount); err != nil {
		return err
	}
	if err := w.WriteData(o.AdvancedServer); err != nil {
		return err
	}
	return nil
}
func (o *CAInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.CAType); err != nil {
		return err
	}
	if err := w.ReadData(&o.CASignatureCertsCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.CAExchangeCertsCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.ExitAlgorithmsCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyIDMax); err != nil {
		return err
	}
	if err := w.ReadData(&o.RoleSeparationEnabled); err != nil {
		return err
	}
	if err := w.ReadData(&o.KRACertUsedCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.KRACertCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.AdvancedServer); err != nil {
		return err
	}
	return nil
}

// CertRequestD structure represents ICertRequestD RPC structure.
type CertRequestD dcom.InterfacePointer

func (o *CertRequestD) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *CertRequestD) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *CertRequestD) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *CertRequestD) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *CertRequestD) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}

// CertRequestD2 structure represents ICertRequestD2 RPC structure.
type CertRequestD2 dcom.InterfacePointer

func (o *CertRequestD2) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *CertRequestD2) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.DataCount == 0 {
		o.DataCount = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *CertRequestD2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *CertRequestD2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.DataCount); err != nil {
		return err
	}
	for i1 := range o.Data {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(o.Data[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *CertRequestD2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.DataCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DataCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DataCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
	}
	o.Data = make([]byte, sizeInfo[0])
	for i1 := range o.Data {
		i1 := i1
		if err := w.ReadData(&o.Data[i1]); err != nil {
			return err
		}
	}
	return nil
}
