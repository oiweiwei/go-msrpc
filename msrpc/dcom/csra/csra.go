// The csra package implements the CSRA client protocol.
//
// # Introduction
//
// The Certificate Services Remote Administration Protocol consists of a set of Distributed
// Component Object Model (DCOM) interfaces, as specified in [MS-DCOM], that allow administrative
// tools to configure the state and policy of a certification authority (CA) on a server.
//
// For a complete understanding of this protocol, familiarity with public key infrastructure
// (PKI) concepts such as asymmetric and symmetric cryptography, asymmetric and symmetric
// encryption techniques, digital certificate concepts, and cryptographic key establishment
// is required. A comprehensive understanding of the X.509 standard, as specified in
// [X509], is also required.
//
// The Handbook of Applied Cryptography provides an excellent introduction to cryptography
// and PKI concepts. For more information, see [CRYPTO]. The X.509 standard, as specified
// in [X509], provides an excellent introduction to PKI and certificate concepts. Certificate
// revocation and status checking provides an excellent introduction to certificate
// revocation lists (CRLs) and revocation concepts. For more information, see [MSFT-CRL].
//
// # Overview
//
// The Certificate Services Remote Administration Protocol consists of a set of DCOM
// interfaces, as specified in [MS-DCOM], that allow administrative tools to configure
// the state and policy of a CA on a server. The administrative tools can perform such
// functions as getting or setting properties on a CA, retrieving data, revoking certificates,
// or retrieving escrowed private keys from a CA.
//
// The following figure reflects only CA administration, not the normal operation of
// the CA. The protocol for the normal operation of the Microsoft CA is specified in
// [MS-WCCE].
package csra

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
	GoPackage = "dcom/csra"
)

// CertTransportBlob structure represents CERTTRANSBLOB RPC structure.
//
// The CERTTRANSBLOB structure defines a byte buffer that is used to store and request
// certificates, transmit responses, manipulate Unicode strings, and marshal property
// values.
type CertTransportBlob struct {
	// cb:  An unsigned integer value that MUST contain the length, in bytes, of the buffer
	// that is pointed to by pb.
	Length uint32 `idl:"name:cb" json:"length"`
	// pb:  The BYTE buffer that contains the binary contents being transported in this
	// CERTTRANSBLOB. That content consists of any of the following entities:
	//
	// * A certificate.
	//
	// * A certificate request.
	//
	// * CA ( c6451297-197d-4b4b-b786-3f3187b67b8f#gt_c925d5d7-a442-4ba4-9586-5f94ccec847a
	// ) properties.
	//
	// * Any common structure that is defined in section 2.2.1 ( d9e0f247-2b38-466d-934b-83093c6a11a5
	// ) other than VARIANT ( 8d5e0fb0-f357-48b2-808c-bb125fd0609e ) or CERTVIEWRESTRICTION
	// ( 5503c7fa-c78e-4fda-adc9-21030751bce7 ).
	//
	// * Any common structure that is defined in [MS-WCCE] ( ../ms-wcce/446a0fca-7f27-4436-965d-191635518466
	// ) section 2.2.2 ( ../ms-wcce/a2d33e71-31d9-4934-a369-07ed8c502ae5 ).
	//
	// The CERTTRANSBLOB structure is empty when cb is set to 0 and pb is set to NULL.
	//
	// The marshaling of other structures that can be passed in the pb byte buffer of CERTTRANSBLOB
	// is defined in [MS-WCCE] section 2.2.2.
	//
	// All instances of CERTTRANSBLOB that are used by this protocol MUST use the marshaling
	// rules that are described in the following sections or in [MS-WCCE] section 2.2.2.
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
// The CATRANSPROP structure encapsulates information about a CA property. The CATRANSPROP
// structure and the marshaling of one or more CATRANSPROP structures into a CERTTRANSBLOB
// structure is specified in [MS-WCCE] section 2.2.2.3.
type CATransportProperty struct {
	PropertyID        int32  `idl:"name:lPropID" json:"property_id"`
	PropertyType      uint8  `idl:"name:propType" json:"property_type"`
	_                 uint8  `idl:"name:Reserved"`
	PropertyFlags     uint16 `idl:"name:propFlags" json:"property_flags"`
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

// CertTransDBAttribute structure represents CERTTRANSDBATTRIBUTE RPC structure.
//
// The CERTTRANSDBATTRIBUTE structure is encoded within a CERTTRANSBLOB structure. The
// CERTTRANSDBATTRIBUTE structure is used by the server to return attribute information
// that is associated with a request to the client (upon the client's query via invocation
// of the EnumAttributesOrExtensions method of the ICertAdminD interface).
type CertTransDBAttribute struct {
	// obwszName:  An integer that contains the offset from the beginning of the byte array
	// buffer that is pointed to by the pb member in the containing CERTTRANSBLOB structure
	// to where the string that contains the name of this attribute can be found. The format
	// is a null-terminated UNICODE string. The offset MUST be divisible by 4.
	NameOffset uint32 `idl:"name:obwszName" json:"name_offset"`
	// obwszValue:  An integer that contains the offset from the beginning of the byte array
	// buffer that is pointed to by the pb member in the containing CERTTRANSBLOB structure
	// to where the string that contains the value of this attribute can be found. The format
	// is a null-terminated UNICODE string. The offset MUST be divisible by 4.
	ValueOffset uint32 `idl:"name:obwszValue" json:"value_offset"`
}

func (o *CertTransDBAttribute) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CertTransDBAttribute) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.NameOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.ValueOffset); err != nil {
		return err
	}
	return nil
}
func (o *CertTransDBAttribute) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.NameOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValueOffset); err != nil {
		return err
	}
	return nil
}

// CertTransDBExtension structure represents CERTTRANSDBEXTENSION RPC structure.
//
// The CERTTRANSDBEXTENSION structure is encoded within a CERTTRANSBLOB structure. The
// CERTTRANSDBEXTENSION structure is used by the server to return certificate extension
// information, as specified in [RFC3280] section 4, that is associated with a request.
// This associated request to the client occurs when the client performs a query by
// invoking the EnumAttributesOrExtensions method of the ICertAdminD interface.
type CertTransDBExtension struct {
	// obwszName:  An unsigned integer that contains the offset from the beginning of the
	// byte array buffer that is pointed to by the pb member in the containing CERTTRANSBLOB
	// structure to the string representation of an OID (1) of this extension (as specified
	// in [X680]). The string format is a null-terminated UNICODE string. The offset MUST
	// be divisible by 4.
	NameOffset uint32 `idl:"name:obwszName" json:"name_offset"`
	// ExtFlags:  An integer value that specifies the flags that are associated with the
	// extension. The following diagram shows its contents.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| Nigiro                                                        |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	// C (1 bit): C is the ExtensionCriticalFlag, as defined in section 3.1.1.3, the Extension_Flags
	// ADM element. A value of 0 means the extension is not critical. A value of 1 means
	// the extension is critical.
	//
	// D (1 bit): D is the ExtensionDisabledFlag, as defined in section 3.1.1.3, the Extension_Flags
	// ADM element. A value of 0 means the extension is not disabled. A value of 1 means
	// the extension is disabled.
	//
	// Nigiro (2 bytes): The Nigiro field is defined as follows:
	//
	//	+------------------------+----------------------------------------------------------------------------------+
	//	|   MIRRORED (NIGIRO)    |                                                                                  |
	//	|          BYTE          |                                     MEANING                                      |
	//	|                        |                                                                                  |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| 0x8000                 | The extension comes from the request.                                            |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| 0x4000                 | The extension was added by the CA. The CA assigns a value of 2 if the extension  |
	//	|                        | was added by the policy module of the CA.                                        |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| 0xC000                 | The extension was added by the CA. The CA assigns a value of 3 if the extension  |
	//	|                        | was added interactively by a human administrator of the CA.                      |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| 0x2000                 | The extension was added by the CA. The CA assigns a value of 4 if the extension  |
	//	|                        | was added by the certificate server engine and not the policy module component   |
	//	|                        | of the CA.                                                                       |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| 0xA000                 | The extension was in the preceding certificate, which might occur, for example,  |
	//	|                        | when a certificate is renewed.                                                   |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| 0x6000                 | The extension comes from an imported certificate (a certificate that was         |
	//	|                        | imported into the CA database).                                                  |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| 0xE000                 | The extension comes from a PKCS7 request.                                        |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| 0x1000                 | The extension comes from a CMC request.                                          |
	//	+------------------------+----------------------------------------------------------------------------------+
	//	| 0x9000                 | The extension comes from the current CA signing certificate.                     |
	//	+------------------------+----------------------------------------------------------------------------------+
	ExtFlags int32 `idl:"name:ExtFlags" json:"ext_flags"`
	// cbValue:  An unsigned integer value that contains the length, in bytes, of data that
	// is referenced by the obValue parameter.
	ValueLength uint32 `idl:"name:cbValue" json:"value_length"`
	// obValue:  An unsigned integer that contains the offset from the beginning of the
	// byte array buffer that is pointed to by the pb member in the containing CERTTRANSBLOB
	// structure to where the value for this extension can be found. The length of the value
	// is specified in the cbValue field. The value is in ASN.1 Distinguished Encoding Rules
	// (DER) format for the extension, as specified in [X660]. The offset MUST be divisible
	// by 4.
	ValueOffset uint32 `idl:"name:obValue" json:"value_offset"`
}

func (o *CertTransDBExtension) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CertTransDBExtension) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.NameOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.ExtFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.ValueLength); err != nil {
		return err
	}
	if err := w.WriteData(o.ValueOffset); err != nil {
		return err
	}
	return nil
}
func (o *CertTransDBExtension) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.NameOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.ExtFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValueLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValueOffset); err != nil {
		return err
	}
	return nil
}

// CertTransDBColumn structure represents CERTTRANSDBCOLUMN RPC structure.
//
// The CERTTRANSDBCOLUMN structure is encoded within a CERTTRANSBLOB structure. The
// CERTTRANSDBCOLUMN structure contains schema information about a particular database
// column that is associated with a specific table to the client. This associated table
// is invoked when the client queries the EnumViewColumn or EnumViewColumnTable method
// of the ICertAdminD and ICertAdminD2 interfaces, respectively.
//
// The CERTTRANSDBCOLUMN structure (section 2.2.1.7) is encoded within the byte array
// that is referenced by the pb member of a CERTTRANSBLOB structure (section 2.2.1.4).
//
// The packet that contains an array of some number, "N", of CERTTRANSDBCOLUMN structures
// is specified in the following packet diagrams. The actual value of "N" is a separate
// return parameter for the EnumViewColumn (section 3.1.4.1.9) and EnumViewColumnTable
// (section 3.2.4.2.5) methods.
//
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| CERTTRANSDBCOLUMN Structures (variable)                                                                                       |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| Column_Schema_Data (variable)                                                                                                 |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
//	| ...                                                                                                                           |
//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
type CertTransDBColumn struct {
	// Type:  This field describes the column. It consists of two WORDs, a high WORD and
	// a low WORD, which are used separately.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| Column Flags                                                  | Column Type                                                   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	Type uint32 `idl:"name:Type" json:"type"`
	// Index:  An unsigned integer value that specifies the identifier for the column in
	// the server database.
	Index uint32 `idl:"name:Index" json:"index"`
	// cbMax:  An unsigned integer value that specifies the maximum data size, in bytes,
	// that this column can contain.
	MaxLength uint32 `idl:"name:cbMax" json:"max_length"`
	// obwszName:  An integer that contains the offset from the beginning of the byte array
	// buffer that is pointed to by the pb member in the containing CERTTRANSBLOB structure,
	// to where the string that contains the name of this column can be found. The string
	// format is a null-terminated UNICODE string. The offset MUST be divisible by 4.
	NameOffset uint32 `idl:"name:obwszName" json:"name_offset"`
	// obwszDisplayName:  An integer that contains the offset from the beginning of the
	// byte array buffer that is pointed to by the pb member in the containing CERTTRANSBLOB
	// structure, to where the string that contains the display name of this column can
	// be found. The string format is a null-terminated UNICODE string. The offset MUST
	// be divisible by 4.
	DisplayNameOffset uint32 `idl:"name:obwszDisplayName" json:"display_name_offset"`
}

func (o *CertTransDBColumn) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CertTransDBColumn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteData(o.Index); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxLength); err != nil {
		return err
	}
	if err := w.WriteData(o.NameOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.DisplayNameOffset); err != nil {
		return err
	}
	return nil
}
func (o *CertTransDBColumn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadData(&o.Index); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.NameOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.DisplayNameOffset); err != nil {
		return err
	}
	return nil
}

// CertTransDBResultColumn structure represents CERTTRANSDBRESULTCOLUMN RPC structure.
//
// The CERTTRANSDBRESULTCOLUMN structure is encoded within a CERTTRANSBLOB structure.
// The CERTTRANSDBRESULTCOLUMN structure is used by the server to return the result
// of a CA database query that is done by the client (upon the client's query via invocation
// of the OpenView or EnumView method of the ICertAdminD interface).
//
// The OpenView and EnumView methods return data in the form of a CERTTRANSBLOB structure
// whose pb member points to an array of one or more CERTTRANSDBRESULTROW structures.
// Each CERTTRANSDBRESULTROW structure contains one or more CERTTRANSDBRESULTCOLUMN
// structures.
//
// The CERTTRANSDBRESULTCOLUMN structure contains data for a specific column in a specific
// row.
type CertTransDBResultColumn struct {
	// Type:  This field describes the column. It consists of two WORDs, a high WORD and
	// a low WORD, which are used separately.
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| Column Flags                                                  | Column Type                                                   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	Type uint32 `idl:"name:Type" json:"type"`
	// Index:  An unsigned integer value that specifies the identifier for the column in
	// the relevant table.
	Index uint32 `idl:"name:Index" json:"index"`
	// obValue:  An unsigned integer that contains the offset from the beginning of the
	// corresponding CERTTRANSDBRESULTROW structure to where the value for this column can
	// be found. The length of the value is specified in the cbValue field. The offset MUST
	// be DWORD aligned.
	ValueOffset uint32 `idl:"name:obValue" json:"value_offset"`
	// cbValue:  An unsigned integer value that specifies the length, in bytes, of the value
	// for the specific column.
	ValueLength uint32 `idl:"name:cbValue" json:"value_length"`
}

func (o *CertTransDBResultColumn) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CertTransDBResultColumn) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if err := w.WriteData(o.Index); err != nil {
		return err
	}
	if err := w.WriteData(o.ValueOffset); err != nil {
		return err
	}
	if err := w.WriteData(o.ValueLength); err != nil {
		return err
	}
	return nil
}
func (o *CertTransDBResultColumn) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	if err := w.ReadData(&o.Index); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValueOffset); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValueLength); err != nil {
		return err
	}
	return nil
}

// CertTransDBResultRow structure represents CERTTRANSDBRESULTROW RPC structure.
//
// The CERTTRANSDBRESULTROW structure is encoded within a CERTTRANSBLOB structure. The
// CERTTRANSDBRESULTROW structure is used by the server to return the result of the
// database query done by the client (upon the client's query via invocation of OpenView
// or EnumView methods of the ICertAdminD interface). This structure contains data for
// a specific row.
type CertTransDBResultRow struct {
	// rowid:  Unsigned integer value that specifies the identifier for the row.
	RowID uint32 `idl:"name:rowid" json:"row_id"`
	// ccol:  Unsigned integer value that specifies the count of CERTTRANSDBRESULTCOLUMN
	// structures. Each structure contains the value of a specific column in the row identified
	// by rowid.
	ColumnsCount uint32 `idl:"name:ccol" json:"columns_count"`
	// cbrow:  Unsigned integer value that specifies the total size of row data (in bytes).
	// This is the sum of the size of CERTTRANSDBRESULTROW structure, size of each CERTTRANSDBRESULTCOLUMN
	// structure for the row (the count of which is specified by ccol), and the DWORD-rounded-up
	// size of each column value.
	RowLength uint32 `idl:"name:cbrow" json:"row_length"`
}

func (o *CertTransDBResultRow) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CertTransDBResultRow) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.RowID); err != nil {
		return err
	}
	if err := w.WriteData(o.ColumnsCount); err != nil {
		return err
	}
	if err := w.WriteData(o.RowLength); err != nil {
		return err
	}
	return nil
}
func (o *CertTransDBResultRow) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.RowID); err != nil {
		return err
	}
	if err := w.ReadData(&o.ColumnsCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.RowLength); err != nil {
		return err
	}
	return nil
}

// CertViewRestriction structure represents CERTVIEWRESTRICTION RPC structure.
//
// The CERTVIEWRESTRICTION structure is used to restrict the data set that is returned
// by the CA server during calls to the OpenView method for the ICertAdminD interface.
//
// This structure is passed by RPC technology, as specified in [MS-RPCE], and does not
// need special marshaling.
type CertViewRestriction struct {
	// ColumnIndex:  An unsigned integer value that specifies the identifier for the database
	// column that is receiving the restriction.
	ColumnIndex uint32 `idl:"name:ColumnIndex" json:"column_index"`
	// SeekOperator:  An integer value that specifies the logical operator of the data-query
	// qualifier for the column. This parameter MUST be set to one of the following values.
	//
	//	+------------+--------------------------+
	//	|            |                          |
	//	|   VALUE    |         MEANING          |
	//	|            |                          |
	//	+------------+--------------------------+
	//	+------------+--------------------------+
	//	| 0x00000001 | Equal to                 |
	//	+------------+--------------------------+
	//	| 0x00000002 | Less than                |
	//	+------------+--------------------------+
	//	| 0x00000004 | Less than or equal to    |
	//	+------------+--------------------------+
	//	| 0x00000008 | Greater than or equal to |
	//	+------------+--------------------------+
	//	| 0x00000010 | Greater than             |
	//	+------------+--------------------------+
	SeekOperator int32 `idl:"name:SeekOperator" json:"seek_operator"`
	// SortOrder:  An integer value that specifies the sort order for the column. This parameter
	// MUST be set to one of the following values.
	//
	//	+------------+---------------+
	//	|            |               |
	//	|   VALUE    |    MEANING    |
	//	|            |               |
	//	+------------+---------------+
	//	+------------+---------------+
	//	| 0x00000000 | No sort order |
	//	+------------+---------------+
	//	| 0x00000001 | Ascending     |
	//	+------------+---------------+
	//	| 0x00000002 | Descending    |
	//	+------------+---------------+
	SortOrder int32 `idl:"name:SortOrder" json:"sort_order"`
	// pbValue:  A pointer to a byte array that specifies the value against which the value
	// in the corresponding column (specified by ColumnIndex) is compared, using SeekOperator.
	Value []byte `idl:"name:pbValue;size_is:(cbValue);pointer:unique" json:"value"`
	// cbValue:  An unsigned integer value that specifies the length of the byte array that
	// is pointed to by the pbValue field.
	ValueLength uint32 `idl:"name:cbValue" json:"value_length"`
}

func (o *CertViewRestriction) xxx_PreparePayload(ctx context.Context) error {
	if o.Value != nil && o.ValueLength == 0 {
		o.ValueLength = uint32(len(o.Value))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CertViewRestriction) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.ColumnIndex); err != nil {
		return err
	}
	if err := w.WriteData(o.SeekOperator); err != nil {
		return err
	}
	if err := w.WriteData(o.SortOrder); err != nil {
		return err
	}
	if o.Value != nil || o.ValueLength > 0 {
		_ptr_pbValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ValueLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Value {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Value[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Value); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Value, _ptr_pbValue); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ValueLength); err != nil {
		return err
	}
	return nil
}
func (o *CertViewRestriction) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.ColumnIndex); err != nil {
		return err
	}
	if err := w.ReadData(&o.SeekOperator); err != nil {
		return err
	}
	if err := w.ReadData(&o.SortOrder); err != nil {
		return err
	}
	_ptr_pbValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ValueLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ValueLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Value", sizeInfo[0])
		}
		o.Value = make([]byte, sizeInfo[0])
		for i1 := range o.Value {
			i1 := i1
			if err := w.ReadData(&o.Value[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pbValue := func(ptr interface{}) { o.Value = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Value, _s_pbValue, _ptr_pbValue); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValueLength); err != nil {
		return err
	}
	return nil
}

// CertAdminD2 structure represents ICertAdminD2 RPC structure.
type CertAdminD2 dcom.InterfacePointer

func (o *CertAdminD2) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *CertAdminD2) xxx_PreparePayload(ctx context.Context) error {
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
func (o *CertAdminD2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *CertAdminD2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *CertAdminD2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// CertAdminD structure represents ICertAdminD RPC structure.
type CertAdminD dcom.InterfacePointer

func (o *CertAdminD) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *CertAdminD) xxx_PreparePayload(ctx context.Context) error {
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
func (o *CertAdminD) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *CertAdminD) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *CertAdminD) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
