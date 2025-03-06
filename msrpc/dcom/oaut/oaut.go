// The oaut package implements the OAUT client protocol.
//
// # Introduction
//
// The OLE Automation Protocol uses Distributed Component Object Model (DCOM) as its
// transport layer. It provides support for an additional set of types, a late-bound
// calling mechanism, and type description and discovery. The late-bound calling mechanism
// is based on dispatch identifiers and a dispatching table that maps the identifiers
// to specific operations. The dispatch identifiers and the dispatching table are specified
// by using IDL extensions specified in this document. Type description and discovery
// are based on a set of IDL extensions and a set of interfaces that are implemented
// by type library and type description servers.
//
// # Overview
//
// The OLE Automation Protocol extends the use of DCOM by providing support for marshaling
// an additional set of types (known as automation types) and by providing support for
// exposing COM components to automation clients (such as scripting engines) through
// a late-bound calling alternative to the traditional COM calls. Additionally, the
// OLE Automation Protocol specifies how a type browser can discover and interpret type
// information provided by a type description server.
//
// The automation client and server can be present on the same machine, or on different
// machines connected by a network. Automation takes advantage of functionality provided
// by the Microsoft Component Object Model (for more information, see [MSDN-COM]) and
// the Microsoft Distributed Component Object Model (as specified in [MS-DCOM]) for
// creating, activating, and managing the lifetime of the objects exposed by an automation
// server.
//
// To support late-bound calling, the OLE Automation Protocol specifies the following:
//
// * How a server defines a set of automation methods that can be dispatched, based
// on a dispatch ID (DISPID) ( 5583e1b8-454c-4147-9f56-f72416a15bee#gt_3792c5cc-783c-4b2a-a71e-c1ec3f432474
// ).
//
// * How the server provides a way to map a method name to the DISPID.
//
// * How the server performs the late-bound call, based on the DISPID.
//
// The automation methods are defined by using extensions to the IDL language specified
// in [C706] sections 6, 7, 8, 9, 10, 11, 12, 13, and 14. These extensions provide the
// definition of automation interfaces containing automation methods and properties.
// Each IDL definition of an automation method and property can have a unique (per interface)
// integer value associated with it. This value is defined as a DISPID and is statically
// discoverable (from the IDL specification of the method), and dynamically discoverable
// (through a call to IDispatch::GetIDsOfNames (section 3.1.4.3)). This value is then
// used by automation clients to invoke the automation method, or to set or retrieve
// an automation property (through a call to IDispatch::Invoke).
//
// To support this late-bound calling mechanism, Automation defines a set of types,
// VARIANT (section 2.2.29) being the most important of them. A VARIANT can be thought
// of as a discriminated union of all automation-supported types. The OLE Automation
// Protocol imposes the following restriction on the automation interfaces: All types
// of method arguments and properties can be stored as VARIANT structures.
//
// The following illustration shows a generic automation call sequence:
package oaut

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
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
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/oaut"
)

// VariantTrue represents the VARIANT_TRUE RPC constant
var VariantTrue = 65535

// VariantFalse represents the VARIANT_FALSE RPC constant
var VariantFalse = 0

// VarEnum type represents VARENUM RPC enumeration.
//
// The VARENUM enumeration constants are used in the discriminant field, vt, of the
// VARIANT type specified in section 2.2.29.2. When present, the VT_BYREF flag MUST
// be in an OR relation with another value to specify the byref argument for the VARIANT.
// The VT_EMPTY and VT_NULL values MUST NOT be specified with the VT_BYREF bit flag.
//
// The following values are also used in the discriminant field, vt, of the TYPEDESC
// structure specified in section 2.2.37.
//
// The meaning of each VARIANT type constant is specified as follows. The Context column
// specifies the context in which each constant is used. A constant MUST NOT be used
// in a VARIANT unless it is specified with a "V". A constant MUST NOT be used in a
// SAFEARRAY unless it is specified with an "S". A constant MUST NOT be used in a TYPEDESC
// unless it is specified with a "T".
type VarEnum uint16

var (
	// VT_EMPTY:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V       | The type of the contained field is undefined. When this flag is specified, the   |
	//	|         | VARIANT MUST NOT contain a data field. The VARIANT definition is specified in    |
	//	|         | section 2.2.29.2.                                                                |
	//	+---------+----------------------------------------------------------------------------------+
	VarEmpty VarEnum = 0
	// VT_NULL:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V       | The type of the contained field is NULL. When this flag is specified, the        |
	//	|         | VARIANT MUST NOT contain a data field. The VARIANT definition is specified in    |
	//	|         | section 2.2.29.2.                                                                |
	//	+---------+----------------------------------------------------------------------------------+
	VarNull VarEnum = 1
	// VT_I2:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | a 2-byte signed integer.                                                         |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumI2 VarEnum = 2
	// VT_I4:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | a 4-byte signed integer.                                                         |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumI4 VarEnum = 3
	// VT_R4:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | a 4-byte IEEE floating-point number.                                             |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumR4 VarEnum = 4
	// VT_R8:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | an 8-byte IEEE floating-point number.                                            |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumR8 VarEnum = 5
	// VT_CY:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | CURRENCY (see section 2.2.24).                                                   |
	//	+---------+----------------------------------------------------------------------------------+
	VarCurrency VarEnum = 6
	// VT_DATE:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | DATE (see section 2.2.25).                                                       |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumDate VarEnum = 7
	// VT_BSTR:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | BSTR (see section 2.2.23).                                                       |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumString VarEnum = 8
	// VT_DISPATCH:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | a pointer to IDispatch (see section 3.1.4).                                      |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumDispatch VarEnum = 9
	// VT_ERROR:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | HRESULT.                                                                         |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumError VarEnum = 10
	// VT_BOOL:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | VARIANT_BOOL (see section 2.2.27).                                               |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumBool VarEnum = 11
	// VT_VARIANT:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | VARIANT (see section 2.2.29). It MUST appear with the bit flag VT_BYREF.         |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumVariant VarEnum = 12
	// VT_UNKNOWN:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | a pointer to IUnknown.                                                           |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumUnknown VarEnum = 13
	// VT_DECIMAL:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | DECIMAL (see section 2.2.26).                                                    |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumDecimal VarEnum = 14
	// VT_I1:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | a 1-byte integer.                                                                |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumI1 VarEnum = 16
	// VT_UI1:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | a 1-byte unsigned integer.                                                       |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumUI1 VarEnum = 17
	// VT_UI2:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | a 2-byte unsigned integer.                                                       |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumUI2 VarEnum = 18
	// VT_UI4:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | a 4-byte unsigned integer.                                                       |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumUI4 VarEnum = 19
	// VT_I8:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | an 8-byte signed integer.                                                        |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumI8 VarEnum = 20
	// VT_UI8:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | an 8-byte unsigned integer.                                                      |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumUI8 VarEnum = 21
	// VT_INT:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | a 4-byte signed integer.                                                         |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumInt VarEnum = 22
	// VT_UINT:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S, T | Either the specified type, or the type of the element or contained field MUST be |
	//	|         | a 4-byte unsigned integer.                                                       |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumUint VarEnum = 23
	// VT_VOID:
	//
	//	+---------+----------------------------------+
	//	|         |                                  |
	//	| CONTEXT |           DESCRIPTION            |
	//	|         |                                  |
	//	+---------+----------------------------------+
	//	+---------+----------------------------------+
	//	| T       | The specified type MUST be void. |
	//	+---------+----------------------------------+
	VarEnumVoid VarEnum = 24
	// VT_HRESULT:
	//
	//	+---------+-------------------------------------+
	//	|         |                                     |
	//	| CONTEXT |             DESCRIPTION             |
	//	|         |                                     |
	//	+---------+-------------------------------------+
	//	+---------+-------------------------------------+
	//	| T       | The specified type MUST be HRESULT. |
	//	+---------+-------------------------------------+
	VarEnumHResult VarEnum = 25
	// VT_PTR:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| T       | The specified type MUST be a unique pointer, as specified in [C706] section      |
	//	|         | 4.2.20.2.                                                                        |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumPointer VarEnum = 26
	// VT_SAFEARRAY:
	//
	//	+---------+--------------------------------------------------------+
	//	|         |                                                        |
	//	| CONTEXT |                      DESCRIPTION                       |
	//	|         |                                                        |
	//	+---------+--------------------------------------------------------+
	//	+---------+--------------------------------------------------------+
	//	| T       | The specified type MUST be SAFEARRAY (section 2.2.30). |
	//	+---------+--------------------------------------------------------+
	VarEnumSafeArray VarEnum = 27
	// VT_CARRAY:
	//
	//	+---------+------------------------------------------------+
	//	|         |                                                |
	//	| CONTEXT |                  DESCRIPTION                   |
	//	|         |                                                |
	//	+---------+------------------------------------------------+
	//	+---------+------------------------------------------------+
	//	| T       | The specified type MUST be a fixed-size array. |
	//	+---------+------------------------------------------------+
	VarEnumCarray VarEnum = 28
	// VT_USERDEFINED:
	//
	//	+---------+------------------------------------------+
	//	|         |                                          |
	//	| CONTEXT |               DESCRIPTION                |
	//	|         |                                          |
	//	+---------+------------------------------------------+
	//	+---------+------------------------------------------+
	//	| T       | The specified type MUST be user defined. |
	//	+---------+------------------------------------------+
	VarEnumUserdefined VarEnum = 29
	// VT_LPSTR:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| T       | The specified type MUST be a NULL-terminated string, as specified in [C706]      |
	//	|         | section 14.3.4.                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	VarCharString VarEnum = 30
	// VT_LPWSTR:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| T       | The specified type MUST be a zero-terminated string of UNICODE characters, as    |
	//	|         | specified in [C706], section 14.3.4.                                             |
	//	+---------+----------------------------------------------------------------------------------+
	VarUnicodeString VarEnum = 31
	// VT_RECORD:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S    | The type of the element or contained field MUST be a BRECORD (see section        |
	//	|         | 2.2.28.2).                                                                       |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumRecord VarEnum = 36
	// VT_INT_PTR:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| T       | The specified type MUST be either a 4-byte or an 8-byte signed integer. The      |
	//	|         | size of the integer is platform specific and determines the system pointer size  |
	//	|         | value, as specified in section 2.2.21.                                           |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumIntPointer VarEnum = 37
	// VT_UINT_PTR:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| T       | The specified type MUST be either a 4 byte or an 8 byte unsigned integer. The    |
	//	|         | size of the integer is platform specific and determines the system pointer size  |
	//	|         | value, as specified in section 2.2.21.                                           |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumUintPointer VarEnum = 38
	// VT_ARRAY:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S    | The type of the element or contained field MUST be a SAFEARRAY (see section      |
	//	|         | 2.2.30.10).                                                                      |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumArray VarEnum = 8192
	// VT_BYREF:
	//
	//	+---------+----------------------------------------------------------------------------------+
	//	|         |                                                                                  |
	//	| CONTEXT |                                   DESCRIPTION                                    |
	//	|         |                                                                                  |
	//	+---------+----------------------------------------------------------------------------------+
	//	+---------+----------------------------------------------------------------------------------+
	//	| V, S    | The type of the element or contained field MUST be a pointer to one of the types |
	//	|         | listed in the previous rows of this table. If present, this bit flag MUST appear |
	//	|         | in a VARIANT discriminant (see section 2.2.28) with one of the previous flags.   |
	//	+---------+----------------------------------------------------------------------------------+
	VarEnumByref VarEnum = 16384
)

func (o VarEnum) String() string {
	switch o {
	case VarEmpty:
		return "VarEmpty"
	case VarNull:
		return "VarNull"
	case VarEnumI2:
		return "VarEnumI2"
	case VarEnumI4:
		return "VarEnumI4"
	case VarEnumR4:
		return "VarEnumR4"
	case VarEnumR8:
		return "VarEnumR8"
	case VarCurrency:
		return "VarCurrency"
	case VarEnumDate:
		return "VarEnumDate"
	case VarEnumString:
		return "VarEnumString"
	case VarEnumDispatch:
		return "VarEnumDispatch"
	case VarEnumError:
		return "VarEnumError"
	case VarEnumBool:
		return "VarEnumBool"
	case VarEnumVariant:
		return "VarEnumVariant"
	case VarEnumUnknown:
		return "VarEnumUnknown"
	case VarEnumDecimal:
		return "VarEnumDecimal"
	case VarEnumI1:
		return "VarEnumI1"
	case VarEnumUI1:
		return "VarEnumUI1"
	case VarEnumUI2:
		return "VarEnumUI2"
	case VarEnumUI4:
		return "VarEnumUI4"
	case VarEnumI8:
		return "VarEnumI8"
	case VarEnumUI8:
		return "VarEnumUI8"
	case VarEnumInt:
		return "VarEnumInt"
	case VarEnumUint:
		return "VarEnumUint"
	case VarEnumVoid:
		return "VarEnumVoid"
	case VarEnumHResult:
		return "VarEnumHResult"
	case VarEnumPointer:
		return "VarEnumPointer"
	case VarEnumSafeArray:
		return "VarEnumSafeArray"
	case VarEnumCarray:
		return "VarEnumCarray"
	case VarEnumUserdefined:
		return "VarEnumUserdefined"
	case VarCharString:
		return "VarCharString"
	case VarUnicodeString:
		return "VarUnicodeString"
	case VarEnumRecord:
		return "VarEnumRecord"
	case VarEnumIntPointer:
		return "VarEnumIntPointer"
	case VarEnumUintPointer:
		return "VarEnumUintPointer"
	case VarEnumArray:
		return "VarEnumArray"
	case VarEnumByref:
		return "VarEnumByref"
	}
	return "Invalid"
}

// AdvFeatureFlags type represents ADVFEATUREFLAGS RPC enumeration.
//
// The following values are used in the field fFeatures of a SAFEARRAY (section 2.2.30.10)
// data type.
type AdvFeatureFlags uint16

var (
	// FADF_AUTO:  MUST be set if the SAFEARRAY is allocated on the stack. This flag MUST
	// be ignored on receipt.
	AdvFeatureFlagAuto AdvFeatureFlags = 1
	// FADF_STATIC:  MUST be set if the SAFEARRAY is statically allocated. This flag MUST
	// be ignored on receipt.
	AdvFeatureFlagStatic AdvFeatureFlags = 2
	// FADF_EMBEDDED:  MUST be set if the SAFEARRAY is embedded in a structure. This flag
	// MUST be ignored on receipt.
	AdvFeatureFlagEmbedded AdvFeatureFlags = 4
	// FADF_FIXEDSIZE:  MUST be set if the SAFEARRAY cannot be resized or reallocated.
	// This flag MUST be ignored on receipt.
	AdvFeatureFlagFixedSize AdvFeatureFlags = 16
	// FADF_RECORD:  The SAFEARRAY MUST contain elements of a UDT (see section 2.2.28.1)
	AdvFeatureFlagRecord AdvFeatureFlags = 32
	// FADF_HAVEIID:  The SAFEARRAY MUST contain MInterfacePointers elements.
	AdvFeatureFlagHaveIID AdvFeatureFlags = 64
	// FADF_HAVEVARTYPE:  If this bit flag is set, the high word of the cLocks field of
	// the SAFEARRAY MUST contain a VARIANT type constant that describes the type of the
	// array's elements (see sections 2.2.7 and 2.2.30.10).
	AdvFeatureFlagHaveVarType AdvFeatureFlags = 128
	// FADF_BSTR:  The SAFEARRAY MUST contain an array of BSTR elements (see section 2.2.23).
	AdvFeatureFlagString AdvFeatureFlags = 256
	// FADF_UNKNOWN:  The SAFEARRAY MUST contain an array of pointers to IUnknown.
	AdvFeatureFlagUnknown AdvFeatureFlags = 512
	// FADF_DISPATCH:  The SAFEARRAY MUST contain an array of pointers to IDispatch (see
	// section 3.1.4).
	AdvFeatureFlagDispatch AdvFeatureFlags = 1024
	// FADF_VARIANT:  The SAFEARRAY MUST contain an array of VARIANT instances.
	AdvFeatureFlagVariant AdvFeatureFlags = 2048
)

func (o AdvFeatureFlags) String() string {
	switch o {
	case AdvFeatureFlagAuto:
		return "AdvFeatureFlagAuto"
	case AdvFeatureFlagStatic:
		return "AdvFeatureFlagStatic"
	case AdvFeatureFlagEmbedded:
		return "AdvFeatureFlagEmbedded"
	case AdvFeatureFlagFixedSize:
		return "AdvFeatureFlagFixedSize"
	case AdvFeatureFlagRecord:
		return "AdvFeatureFlagRecord"
	case AdvFeatureFlagHaveIID:
		return "AdvFeatureFlagHaveIID"
	case AdvFeatureFlagHaveVarType:
		return "AdvFeatureFlagHaveVarType"
	case AdvFeatureFlagString:
		return "AdvFeatureFlagString"
	case AdvFeatureFlagUnknown:
		return "AdvFeatureFlagUnknown"
	case AdvFeatureFlagDispatch:
		return "AdvFeatureFlagDispatch"
	case AdvFeatureFlagVariant:
		return "AdvFeatureFlagVariant"
	}
	return "Invalid"
}

// SafeArrayType type represents SF_TYPE RPC enumeration.
//
// The SF_TYPE enumeration values are used in the discriminant field, sfType, of a SAFEARRAYUNION
// structure.
//
// The SAFEARRAY feature constants are defined in the SF_TYPE enumeration.
type SafeArrayType uint32

var (
	// SF_ERROR:  This value means that the SAFEARRAY was incorrectly marshaled. The receiver
	// MUST reject any call that has a SAFEARRAY argument with this flag specified, by raising
	// an RPC_X_BAD_STUB_DATA RPC exception.
	//
	// Hex value is 0x0000000A.
	SafeArrayTypeError SafeArrayType = 10
	// SF_I1:  The type of the elements contained in the SAFEARRAY MUST be a 1-byte integer.
	//
	// Hex value is 0x00000010.
	SafeArrayTypeI1 SafeArrayType = 16
	// SF_I2:  The type of the elements contained in the SAFEARRAY MUST be a 2-byte integer.
	//
	// Hex value is 0x00000002.
	SafeArrayTypeI2 SafeArrayType = 2
	// SF_I4:  The type of the elements contained in the SAFEARRAY MUST be a 4-byte integer.
	//
	// Hex value is 0x00000003.
	SafeArrayTypeI4 SafeArrayType = 3
	// SF_I8:  The type of the elements contained in the SAFEARRAY MUST be an 8-byte integer.
	//
	// Hex value is 0x00000014.
	SafeArrayTypeI8 SafeArrayType = 20
	// SF_BSTR:  The type of the elements contained in the SAFEARRAY MUST be a BSTR.
	//
	// Hex value is 0x00000008.
	SafeArrayTypeString SafeArrayType = 8
	// SF_UNKNOWN:  The type of the elements contained in the SAFEARRAY MUST be a pointer
	// to IUnknown.
	//
	// Hex value is 0x0000000D.
	SafeArrayTypeUnknown SafeArrayType = 13
	// SF_DISPATCH:  The type of the elements contained in the SAFEARRAY MUST be a pointer
	// to IDispatch (see section 3.1.4).
	//
	// Hex value is 0x00000009.
	SafeArrayTypeDispatch SafeArrayType = 9
	// SF_VARIANT:  The type of the elements contained in the SAFEARRAY MUST be VARIANT.
	//
	// Hex value is 0x0000000C.
	SafeArrayTypeVariant SafeArrayType = 12
	// SF_RECORD:  The type of the elements contained in the SAFEARRAY is a user-defined
	// type (UDT) (as defined in section 2.2.28.1.
	//
	// Hex value is 0x00000024.
	SafeArrayTypeRecord SafeArrayType = 36
	// SF_HAVEIID:  The type of the elements contained in the SAFEARRAY MUST be an MInterfacePointer.
	//
	// Hex value is 0x0000800D.
	SafeArrayTypeHaveIID SafeArrayType = 32781
)

func (o SafeArrayType) String() string {
	switch o {
	case SafeArrayTypeError:
		return "SafeArrayTypeError"
	case SafeArrayTypeI1:
		return "SafeArrayTypeI1"
	case SafeArrayTypeI2:
		return "SafeArrayTypeI2"
	case SafeArrayTypeI4:
		return "SafeArrayTypeI4"
	case SafeArrayTypeI8:
		return "SafeArrayTypeI8"
	case SafeArrayTypeString:
		return "SafeArrayTypeString"
	case SafeArrayTypeUnknown:
		return "SafeArrayTypeUnknown"
	case SafeArrayTypeDispatch:
		return "SafeArrayTypeDispatch"
	case SafeArrayTypeVariant:
		return "SafeArrayTypeVariant"
	case SafeArrayTypeRecord:
		return "SafeArrayTypeRecord"
	case SafeArrayTypeHaveIID:
		return "SafeArrayTypeHaveIID"
	}
	return "Invalid"
}

// CallConv type represents CALLCONV RPC enumeration.
//
// The CALLCONV enumeration values are used in the callconv field of a FUNCDESC to identify
// the calling convention of a local method defined in the automation type library module,
// as specified in sections 2.2.42 and 2.2.49.9 .
//
// The following calling convention constants are defined in the CALLCONV enumeration:
type CallConv uint32

var (
	// CC_CDECL:  MUST be set if the method was declared with the cdecl keyword.
	CallConvCdecl CallConv = 1
	// CC_PASCAL:  MUST be set if the method was declared with the pascal keyword.
	CallConvPascal CallConv = 2
	// CC_STDCALL:  MUST be set if the method was declared with the stdcall keyword.
	CallConvStdcall CallConv = 4
)

func (o CallConv) String() string {
	switch o {
	case CallConvCdecl:
		return "CallConvCdecl"
	case CallConvPascal:
		return "CallConvPascal"
	case CallConvStdcall:
		return "CallConvStdcall"
	}
	return "Invalid"
}

// FuncFlags type represents FUNCFLAGS RPC enumeration.
//
// The FUNCFLAGS enumeration values are used in the wFuncFlags field of a FUNCDESC to
// identify features of a function, as specified in section 2.2.42.
//
// The function feature constants are defined in the FUNCFLAGS enumeration.
type FuncFlags uint16

var (
	// FUNCFLAG_FRESTRICTED:  MUST be set if the method or property was declared with the
	// [restricted] attribute (as specified in section 2.2.49.5.1).
	FuncFlagsRestricted FuncFlags = 1
	// FUNCFLAG_FSOURCE:  MUST be set if the method or property is a member of an interface
	// declared with the [source] attribute (as specified in section 2.2.49.8).
	FuncFlagsSource FuncFlags = 2
	// FUNCFLAG_FBINDABLE:  MUST be set if the property was declared with the [bindable]
	// attribute (as specified in section 2.2.49.5.2).
	FuncFlagsBindable FuncFlags = 4
	// FUNCFLAG_FREQUESTEDIT:  MUST be set if the property was declared with the [requestedit]
	// attribute (as specified in section 2.2.49.5.2).
	FuncFlagsRequestEdit FuncFlags = 8
	// FUNCFLAG_FDISPLAYBIND:  MUST be set if the property was declared with the [displaybind]
	// attribute (as specified in section 2.2.49.5.2).
	FuncFlagsDisplayBind FuncFlags = 16
	// FUNCFLAG_FDEFAULTBIND:  MUST be set if the property was declared with the [defaultbind]
	// attribute (as specified in section 2.2.49.5.2).
	FuncFlagsDefaultBind FuncFlags = 32
	// FUNCFLAG_FHIDDEN:  MUST be set if the method or property was declared with the [hidden]
	// attribute (as specified in section 2.2.49.5.1).
	FuncFlagsHidden FuncFlags = 64
	// FUNCFLAG_FUSESGETLASTERROR:  MUST be set if the method or property was declared
	// with the [usesgetlasterror] attribute (as specified in section 2.2.49.9) and MUST
	// be ignored on receipt.
	FuncFlagsFusesgetlasterror FuncFlags = 128
	// FUNCFLAG_FDEFAULTCOLLELEM:  MUST be set if the method or property was declared with
	// the [defaultcollelem] attribute (as specified in section 2.2.49.5.1).
	FuncFlagsDefaultCollElem FuncFlags = 256
	// FUNCFLAG_FUIDEFAULT:  MUST be set if the method or property was declared with the
	// [uidefault] attribute (as specified in section 2.2.49.5.1).
	FuncFlagsUIDefault FuncFlags = 512
	// FUNCFLAG_FNONBROWSABLE:  MUST be set if the property was declared with the [nonbrowsable]
	// attribute (as specified in section 2.2.49.5.1).
	FuncFlagsNonBrowsable FuncFlags = 1024
	// FUNCFLAG_FREPLACEABLE:  MUST be set if the property was declared with the [replaceable]
	// attribute (as specified in section 2.2.49.5.1). MUST be ignored on receipt.
	FuncFlagsReplaceable FuncFlags = 2048
	// FUNCFLAG_FIMMEDIATEBIND:  MUST be set if the property was declared with the [immediatebind]
	// attribute (as specified in section 2.2.49.5.2).
	FuncFlagsImmediateBind FuncFlags = 4096
)

func (o FuncFlags) String() string {
	switch o {
	case FuncFlagsRestricted:
		return "FuncFlagsRestricted"
	case FuncFlagsSource:
		return "FuncFlagsSource"
	case FuncFlagsBindable:
		return "FuncFlagsBindable"
	case FuncFlagsRequestEdit:
		return "FuncFlagsRequestEdit"
	case FuncFlagsDisplayBind:
		return "FuncFlagsDisplayBind"
	case FuncFlagsDefaultBind:
		return "FuncFlagsDefaultBind"
	case FuncFlagsHidden:
		return "FuncFlagsHidden"
	case FuncFlagsFusesgetlasterror:
		return "FuncFlagsFusesgetlasterror"
	case FuncFlagsDefaultCollElem:
		return "FuncFlagsDefaultCollElem"
	case FuncFlagsUIDefault:
		return "FuncFlagsUIDefault"
	case FuncFlagsNonBrowsable:
		return "FuncFlagsNonBrowsable"
	case FuncFlagsReplaceable:
		return "FuncFlagsReplaceable"
	case FuncFlagsImmediateBind:
		return "FuncFlagsImmediateBind"
	}
	return "Invalid"
}

// FuncKind type represents FUNCKIND RPC enumeration.
//
// The FUNCKIND enumeration values are used in the funckind field of a FUNCDESC to specify
// the way that a method is accessed, as specified in section 2.2.42.
//
// The following function access constants are defined in the FUNCKIND enumeration.
type FuncKind uint32

var (
	// FUNC_PUREVIRTUAL:  MUST be set if the method described by the FUNCDESC structure
	// is a member of an interface whose associated TYPEKIND value is TKIND_INTERFACE (as
	// specified in section 2.2.17).
	FuncKindPureVirtual FuncKind = 1
	// FUNC_STATIC:  MUST be set if the method described by the FUNCDESC structure is a
	// method member of the module defined with the automation scope (as specified in section
	// 2.2.49.9).
	FuncKindStatic FuncKind = 3
	// FUNC_DISPATCH:  MUST be set if the method described by the FUNCDESC structure is
	// a member of an interface whose associated TYPEKIND value is TKIND_DISPATCH (as specified
	// in section 2.2.17). MUST NOT be set if the FUNC_PUREVIRTUAL flag is set.
	FuncKindDispatch FuncKind = 4
)

func (o FuncKind) String() string {
	switch o {
	case FuncKindPureVirtual:
		return "FuncKindPureVirtual"
	case FuncKindStatic:
		return "FuncKindStatic"
	case FuncKindDispatch:
		return "FuncKindDispatch"
	}
	return "Invalid"
}

// ImplTypeFlags type represents IMPLTYPEFLAGS RPC enumeration.
//
// The IMPLTYPEFLAGS enumeration values are stored in the pImplTypeFlags parameter of
// the ITypeInfo::GetImplTypeFlags method to specify the implementation features of
// a COM server, as specified in section 3.7.4.7.
//
// The following implementation type feature constants are defined in the IMPLTYPEFLAGS
// enumeration.
type ImplTypeFlags uint16

var (
	// IMPLTYPEFLAG_FDEFAULT:  MUST be set if the interface was declared with the [default]
	// attribute (as specified in section 2.2.49.8).
	ImplTypeFlagsDefault ImplTypeFlags = 1
	// IMPLTYPEFLAG_FSOURCE:  MUST be set if the interface was declared with the [source]
	// or [defaultvtable] attributes (as specified in section 2.2.49.8).
	ImplTypeFlagsSource ImplTypeFlags = 2
	// IMPLTYPEFLAG_FRESTRICTED:  MUST be set if the interface was declared with the [restricted]
	// attribute (as specified in section 2.2.49.8).
	ImplTypeFlagsRestricted ImplTypeFlags = 4
	// IMPLTYPEFLAG_FDEFAULTVTABLE:  MUST be set if the interface was declared with the
	// [defaultvtable] attribute (as specified in section 2.2.49.8).
	ImplTypeFlagsDefaultVTable ImplTypeFlags = 8
)

func (o ImplTypeFlags) String() string {
	switch o {
	case ImplTypeFlagsDefault:
		return "ImplTypeFlagsDefault"
	case ImplTypeFlagsSource:
		return "ImplTypeFlagsSource"
	case ImplTypeFlagsRestricted:
		return "ImplTypeFlagsRestricted"
	case ImplTypeFlagsDefaultVTable:
		return "ImplTypeFlagsDefaultVTable"
	}
	return "Invalid"
}

// InvokeKind type represents INVOKEKIND RPC enumeration.
//
// The INVOKEKIND enumeration values are used in the invkind field of a FUNCDESC (section
// 2.2.42) to specify the way that a method is invoked using IDispatch::Invoke (section
// 3.1.4.4). They are also used in the ITypeInfo2::GetFuncIndexOfMemId, ITypeInfo::GetDllEntry
// and ITypeComp::Bind methods to distinguish between properties and property accessor
// methods that have the same MEMBERID (section 2.2.35) but are invoked differently.
//
// Fields and parameters that contain function invocation constants MUST contain a single
// INVOKEKIND value, and MUST NOT contain bitwise combinations of multiple INVOKEKIND
// values.
//
// The function invocation constants are defined in the INVOKEKIND enumeration.
type InvokeKind uint32

var (
	// INVOKE_FUNC:  MUST be set if the type member is a method declared without the [propget],
	// [propput], or [propputref] attributes, or to specify that a client method request
	// MUST NOT return a property.
	InvokeKindFunc InvokeKind = 1
	// INVOKE_PROPERTYGET:  MUST be set if the type member is a property declared with
	// the [propget] attribute (as specified in section 2.2.49.5.1), or to specify that
	// a client method request MUST NOT return anything but an ODL dispinterface property
	// (as specified in section 2.2.49.5.3) or a property declared with the [propget] attribute.
	InvokeKindPropertyGet InvokeKind = 2
	// INVOKE_PROPERTYPUT:  MUST be set if the type member is a property declared with
	// the [propput] attribute (as specified in section 2.2.49.5.1), or to specify that
	// a client method request MUST NOT return anything but a property declared with the
	// [propput] attribute.
	InvokeKindPropertyPut InvokeKind = 4
	// INVOKE_PROPERTYPUTREF:  MUST be set if the type member is a property declared with
	// the [propputref] attribute (as specified in section 2.2.49.5.1), or to specify that
	// a client method request MUST NOT return anything but a property declared with the
	// [propputref] attribute.
	InvokeKindPropertyPutReference InvokeKind = 8
)

func (o InvokeKind) String() string {
	switch o {
	case InvokeKindFunc:
		return "InvokeKindFunc"
	case InvokeKindPropertyGet:
		return "InvokeKindPropertyGet"
	case InvokeKindPropertyPut:
		return "InvokeKindPropertyPut"
	case InvokeKindPropertyPutReference:
		return "InvokeKindPropertyPutReference"
	}
	return "Invalid"
}

// ParamFlags type represents PARAMFLAGS RPC enumeration.
//
// The PARAMFLAGS enumeration values are used in the wParamFlags field of a PARAMFLAGS
// to identify the features of a method parameter, as specified in section 2.2.40.
//
// The following parameter feature constants are defined in the PARAMFLAGS enumeration.
type ParamFlags uint16

var (
	// PARAMFLAG_NONE:  The behavior of the parameter is not specified.
	ParamFlagsNone ParamFlags = 0
	// PARAMFLAG_FIN:  MUST be set if the parameter was declared by using the [in] attribute
	// (for more information, see section 2.2.49.6).
	ParamFlagsIn ParamFlags = 1
	// PARAMFLAG_FOUT:  MUST be set if the parameter was declared by using the [out] attribute
	// (for more information, see section 2.2.49.5).
	ParamFlagsOut ParamFlags = 2
	// PARAMFLAG_FLCID:  MUST be set if the parameter was declared by using the [lcid]
	// attribute (for more information, see section 2.2.49.6).
	ParamFlagsLocaleID ParamFlags = 4
	// PARAMFLAG_FRETVAL:  MUST be set if the parameter was declared by using the [retval]
	// attribute (for more information, see section 2.2.49.6).
	ParamFlagsReturnValue ParamFlags = 8
	// PARAMFLAG_FOPT:  MUST be set if the parameter was declared by using the [optional]
	// attribute (for more information, see section 2.2.49.6). MUST be set if the PARAMFLAG_FHASDEFAULT
	// flag is set.
	ParamFlagsOptional ParamFlags = 16
	// PARAMFLAG_FHASDEFAULT:  MUST be set if the parameter was declared by using the [defaultvalue]
	// attribute (for more information, see section 2.2.49.6).
	ParamFlagsHasDefault ParamFlags = 32
	// PARAMFLAG_FHASCUSTDATA:  MAY<2> be set if the parameter was declared by using the
	// [custom] attribute (for more information, see section 2.2.49.2).
	ParamFlagsHasCustomData ParamFlags = 64
)

func (o ParamFlags) String() string {
	switch o {
	case ParamFlagsNone:
		return "ParamFlagsNone"
	case ParamFlagsIn:
		return "ParamFlagsIn"
	case ParamFlagsOut:
		return "ParamFlagsOut"
	case ParamFlagsLocaleID:
		return "ParamFlagsLocaleID"
	case ParamFlagsReturnValue:
		return "ParamFlagsReturnValue"
	case ParamFlagsOptional:
		return "ParamFlagsOptional"
	case ParamFlagsHasDefault:
		return "ParamFlagsHasDefault"
	case ParamFlagsHasCustomData:
		return "ParamFlagsHasCustomData"
	}
	return "Invalid"
}

// TypeFlags type represents TYPEFLAGS RPC enumeration.
//
// The TYPEFLAGS enumeration values are used in the wTypeFlags field of a TYPEATTR to
// specify the features of a type, as specified in section 2.2.44. They also are used
// in the pTypeFlags parameter of the ITypeInfo2::GetTypeFlags method.
//
// The function invocation constants are defined in the TYPEFLAGS enumeration.
type TypeFlags uint16

var (
	// TYPEFLAG_FAPPOBJECT:  MUST be set if the type was declared with the [appobject]
	// attribute (see section 2.2.49.8).
	TypeFlagsAppObject TypeFlags = 1
	// TYPEFLAG_FCANCREATE:  MUST NOT be set if the type was declared with the [noncreatable]
	// attribute (see section 2.2.49.8). Otherwise, MUST be set.
	TypeFlagsCanCreate TypeFlags = 2
	// TYPEFLAG_FLICENSED:  MUST be set if the type was declared with the [licensed] attribute
	// (see section 2.2.49.8).
	TypeFlagsLicensed TypeFlags = 4
	// TYPEFLAG_FPREDECLID:  MUST be set if the type was declared with the [predeclid]
	// or [appobject] attributes (see section 2.2.49.8).
	TypeFlagsPredeclID TypeFlags = 8
	// TYPEFLAG_FHIDDEN:  MUST be set if the type was declared with the [hidden] attribute
	// (see section 2.2.49.8).
	TypeFlagsHidden TypeFlags = 16
	// TYPEFLAG_FCONTROL:  MUST be set if the type was declared with the [control] attribute
	// (see section 2.2.49.8).
	TypeFlagsControl TypeFlags = 32
	// TYPEFLAG_FDUAL:  MUST be set if the type was declared with the [dual] attribute
	// (see section 2.2.49.4.2).
	TypeFlagsDual TypeFlags = 64
	// TYPEFLAG_FNONEXTENSIBLE:  MUST be set if the type was declared with the [nonextensible]
	// attribute (see section 2.2.49.4).
	TypeFlagsNonExtensible TypeFlags = 128
	// TYPEFLAG_FOLEAUTOMATION:  MUST be set if the type is a DCOM interface that was declared
	// with the [oleautomation] or [dual] attributes (see section 2.2.49.4). MUST NOT be
	// set if the type is a dispinterface.
	TypeFlagsOLEAutomation TypeFlags = 256
	// TYPEFLAG_FRESTRICTED:  MUST be set if the type was declared with the [restricted]
	// attribute (see section 2.2.49.5.1).
	TypeFlagsRestricted TypeFlags = 512
	// TYPEFLAG_FAGGREGATABLE:  MUST be set if the type was declared with the [aggregatable]
	// attribute (see section 2.2.49.8).
	TypeFlagsAggregatable TypeFlags = 1024
	// TYPEFLAG_FREPLACEABLE:  MUST be set if the type contains a member that was declared
	// with the [replaceable] attribute (see section 2.2.49.5.1). MUST be ignored on receipt.
	TypeFlagsReplaceable TypeFlags = 2048
	// TYPEFLAG_FDISPATCHABLE:  MUST be set if the type derives from IDispatch, either
	// directly or indirectly. MUST be set if the type is a dispinterface or dual interface
	// <3> (see section 2.2.49.4.2).
	TypeFlagsDispatchable TypeFlags = 4096
	// TYPEFLAG_FPROXY:  MUST be set if the type was declared with the [proxy] attribute
	// (see section 2.2.49.4). MUST be ignored on receipt.
	TypeFlagsProxy TypeFlags = 16384
)

func (o TypeFlags) String() string {
	switch o {
	case TypeFlagsAppObject:
		return "TypeFlagsAppObject"
	case TypeFlagsCanCreate:
		return "TypeFlagsCanCreate"
	case TypeFlagsLicensed:
		return "TypeFlagsLicensed"
	case TypeFlagsPredeclID:
		return "TypeFlagsPredeclID"
	case TypeFlagsHidden:
		return "TypeFlagsHidden"
	case TypeFlagsControl:
		return "TypeFlagsControl"
	case TypeFlagsDual:
		return "TypeFlagsDual"
	case TypeFlagsNonExtensible:
		return "TypeFlagsNonExtensible"
	case TypeFlagsOLEAutomation:
		return "TypeFlagsOLEAutomation"
	case TypeFlagsRestricted:
		return "TypeFlagsRestricted"
	case TypeFlagsAggregatable:
		return "TypeFlagsAggregatable"
	case TypeFlagsReplaceable:
		return "TypeFlagsReplaceable"
	case TypeFlagsDispatchable:
		return "TypeFlagsDispatchable"
	case TypeFlagsProxy:
		return "TypeFlagsProxy"
	}
	return "Invalid"
}

// TypeKind type represents TYPEKIND RPC enumeration.
//
// The TYPEKIND enumeration values are used in the typekind field of a TYPEATTR to specify
// the features of a type, as specified in section 2.2.44. They are also used in the
// pTypeKind parameter of the ITypeInfo2::GetTypeKind method, as specified in section
// 3.9.4.1.
//
// The type kind constants are defined in the TYPEKIND enumeration.
type TypeKind uint32

var (
	// TKIND_ENUM:  MUST be used if the type is an enumeration that was defined with the
	// typedef and enum keywords.
	TypeKindEnum TypeKind = 0
	// TKIND_RECORD:  MUST be used if the type is a structure that was defined with the
	// typedef and struct keywords.
	TypeKindRecord TypeKind = 1
	// TKIND_MODULE:  MUST be used if the type is a module that was defined with the module
	// keyword.
	TypeKindModule TypeKind = 2
	// TKIND_INTERFACE:  MUST be used if the type is a DCOM interface that was defined
	// with the interface keyword.
	TypeKindInterface TypeKind = 3
	// TKIND_DISPATCH:  MUST be used if the type is a dispinterface that was defined with
	// either the dispinterface keyword or the interface keyword with the [dual] attribute.
	TypeKindDispatch TypeKind = 4
	// TKIND_COCLASS:  MUST be used if the type is a COM server that was defined with the
	// coclass keyword.
	TypeKindCoclass TypeKind = 5
	// TKIND_ALIAS:  MUST be used if the type is an alias for a predefined type that was
	// defined with the typedef keyword and added to the automation scope by using the [public]
	// attribute as specified in section 2.2.49.3.
	TypeKindAlias TypeKind = 6
	// TKIND_UNION:  MUST be used if the type is a union that was defined with the typedef
	// and union keywords.
	TypeKindUnion TypeKind = 7
)

func (o TypeKind) String() string {
	switch o {
	case TypeKindEnum:
		return "TypeKindEnum"
	case TypeKindRecord:
		return "TypeKindRecord"
	case TypeKindModule:
		return "TypeKindModule"
	case TypeKindInterface:
		return "TypeKindInterface"
	case TypeKindDispatch:
		return "TypeKindDispatch"
	case TypeKindCoclass:
		return "TypeKindCoclass"
	case TypeKindAlias:
		return "TypeKindAlias"
	case TypeKindUnion:
		return "TypeKindUnion"
	}
	return "Invalid"
}

// VarFlags type represents VARFLAGS RPC enumeration.
//
// The VARFLAGS enumeration values are used in the wVarFlags field of a VARDESC to specify
// the features of a field, constant, or ODL dispinterface property, as specified in
// section 2.2.43.
//
// The variable feature constants are defined in the VARFLAGS enumeration.
type VarFlags uint16

var (
	// VARFLAG_FREADONLY:  MUST be set if the variable is an ODL dispinterface property
	// that was declared with the [readonly] attribute (see section 2.2.49.5.3).
	VarFlagsReadOnly VarFlags = 1
	// VARFLAG_FSOURCE:  MUST be set if the variable is a property member of an ODL interface
	// that was declared with the [source] attribute (see section 2.2.49.8).
	VarFlagsSource VarFlags = 2
	// VARFLAG_FBINDABLE:  MUST be set if the variable is an ODL dispinterface property
	// that was declared with the [bindable] attribute (see section 2.2.49.5.2).
	VarFlagsBindable VarFlags = 4
	// VARFLAG_FREQUESTEDIT:  MUST be set if the variable is an ODL dispinterface property
	// that was declared with the [requestedit] attribute (see section 2.2.49.5.2).
	VarFlagsRequestEdit VarFlags = 8
	// VARFLAG_FDISPLAYBIND:  MUST be set if the variable is an ODL dispinterface property
	// that was declared with the [displaybind] attribute (see section 2.2.49.5.2).
	VarFlagsDisplayBind VarFlags = 16
	// VARFLAG_FDEFAULTBIND:  MUST be set if the variable is an ODL dispinterface property
	// that was declared with the [defaultbind] attribute (see section 2.2.49.5.2).
	VarFlagsDefaultBind VarFlags = 32
	// VARFLAG_FHIDDEN:  MUST be set if the variable is a member of a type that was declared
	// with the [hidden] attribute (see section 2.2.49.5.1).
	VarFlagsHidden VarFlags = 64
	// VARFLAG_FRESTRICTED:  MUST be set if the variable is a member of a type that was
	// declared with the [restricted] attribute (see section 2.2.49.5.1).
	VarFlagsRestricted VarFlags = 128
	// VARFLAG_FDEFAULTCOLLELEM:  MUST be set if the variable is an ODL dispinterface property
	// that was declared with the [defaultcollelem] attribute (see section 2.2.49.5.1).
	VarFlagsDefaultCollElem VarFlags = 256
	// VARFLAG_FUIDEFAULT:  MUST be set if the variable is an ODL dispinterface property
	// that was declared with the [uidefault] attribute (see section 2.2.49.5.1).
	VarFlagsUIDefault VarFlags = 512
	// VARFLAG_FNONBROWSABLE:  MUST be set if the variable is an ODL dispinterface property
	// that was declared with the [nonbrowsable] attribute (see section 2.2.49.5.1).
	VarFlagsNonBrowsable VarFlags = 1024
	// VARFLAG_FREPLACEABLE:  MUST be set if the variable is an ODL dispinterface property
	// that was declared with the [replaceable] attribute (see section 2.2.49.5.1). MUST
	// be ignored on receipt.
	VarFlagsReplaceable VarFlags = 2048
	// VARFLAG_FIMMEDIATEBIND:  MUST be set if the variable is an ODL dispinterface property
	// that was declared with the [immediatebind] attribute (see section 2.2.49.5.2).
	VarFlagsImmediateBind VarFlags = 4096
)

func (o VarFlags) String() string {
	switch o {
	case VarFlagsReadOnly:
		return "VarFlagsReadOnly"
	case VarFlagsSource:
		return "VarFlagsSource"
	case VarFlagsBindable:
		return "VarFlagsBindable"
	case VarFlagsRequestEdit:
		return "VarFlagsRequestEdit"
	case VarFlagsDisplayBind:
		return "VarFlagsDisplayBind"
	case VarFlagsDefaultBind:
		return "VarFlagsDefaultBind"
	case VarFlagsHidden:
		return "VarFlagsHidden"
	case VarFlagsRestricted:
		return "VarFlagsRestricted"
	case VarFlagsDefaultCollElem:
		return "VarFlagsDefaultCollElem"
	case VarFlagsUIDefault:
		return "VarFlagsUIDefault"
	case VarFlagsNonBrowsable:
		return "VarFlagsNonBrowsable"
	case VarFlagsReplaceable:
		return "VarFlagsReplaceable"
	case VarFlagsImmediateBind:
		return "VarFlagsImmediateBind"
	}
	return "Invalid"
}

// VarKind type represents VARKIND RPC enumeration.
//
// The VARKIND enumeration values are used in the varkind field of a VARDESC to specify
// the kind of element that is described by the VARDESC, as specified in section 2.2.43.
//
// The variable kind constants are defined in the VARKIND enumeration:
type VarKind uint32

var (
	// VAR_PERINSTANCE:  MUST be used if the VARDESC describes a member of a structure
	// or union.
	VarKindPerinstance VarKind = 0
	// VAR_STATIC:  MUST be used if the VARDESC describes an appobject coclass (see section
	// 2.2.49.8).
	VarKindStatic VarKind = 1
	// VAR_CONST:  MUST be used if the VARDESC describes a member of a module or enumeration.
	VarKindConst VarKind = 2
	// VAR_DISPATCH:  MUST be used if the VARDESC describes an ODL dispinterface property
	// (see section 2.2.49.5.3).
	VarKindDispatch VarKind = 3
)

func (o VarKind) String() string {
	switch o {
	case VarKindPerinstance:
		return "VarKindPerinstance"
	case VarKindStatic:
		return "VarKindStatic"
	case VarKindConst:
		return "VarKindConst"
	case VarKindDispatch:
		return "VarKindDispatch"
	}
	return "Invalid"
}

// LibFlags type represents LIBFLAGS RPC enumeration.
//
// The LIBFLAGS enumeration values are used in the wLibFlags field of a TLIBATTR to
// specify the features of the automation scope of an ITypeLib server, as specified
// in section 2.2.45.
//
// The Type library feature constants are defined in the LIBFLAGS enumeration.
type LibFlags uint32

var (
	// LIBFLAG_FRESTRICTED:  MUST be set if the automation scope was declared with the
	// [restricted] attribute (as specified in section 2.2.49.2).
	LibFlagsRestricted LibFlags = 1
	// LIBFLAG_FCONTROL:  MUST be set if the automation scope was declared with the [control]
	// attribute (as specified in section 2.2.49.2).
	LibFlagsControl LibFlags = 2
	// LIBFLAG_FHIDDEN:  MUST be set if the automation scope was declared with the [hidden]
	// attribute (as specified in section 2.2.49.2).
	LibFlagsHidden LibFlags = 4
	// LIBFLAG_FHASDISKIMAGE:  MAY be set <4>and MUST be ignored on receipt.
	LibFlagsHasDiskImage LibFlags = 8
)

func (o LibFlags) String() string {
	switch o {
	case LibFlagsRestricted:
		return "LibFlagsRestricted"
	case LibFlagsControl:
		return "LibFlagsControl"
	case LibFlagsHidden:
		return "LibFlagsHidden"
	case LibFlagsHasDiskImage:
		return "LibFlagsHasDiskImage"
	}
	return "Invalid"
}

// SystemKind type represents SYSKIND RPC enumeration.
//
// SYSKIND is used in the syskind field of a TLIBATTR to specify the system pointer
// size value, as specified in section 2.2.45.
//
// The system pointer size constants are defined in the SYSKIND enumeration.
type SystemKind uint32

var (
	// SYS_WIN32:  MUST be set if the automation type library uses 32 bits for pointer-sized
	// values.
	SystemKindWin32 SystemKind = 1
	// SYS_WIN64:  MUST be set if the automation type library uses 64 bits for pointer-sized
	// values.
	SystemKindWin64 SystemKind = 3
)

func (o SystemKind) String() string {
	switch o {
	case SystemKindWin32:
		return "SystemKindWin32"
	case SystemKindWin64:
		return "SystemKindWin64"
	}
	return "Invalid"
}

// DescKind type represents DESCKIND RPC enumeration.
//
// The DESCKIND Name Description Constants enumeration values are used by the ITypeComp::Bind
// method to indicate the kind of element to which a name has been bound, as specified
// in section 3.5.4.1.
//
// The name description constants are defined in the DESCKIND enumeration.
type DescKind uint32

var (
	// DESCKIND_NONE:  MUST be set if there is no element bound to the name.
	DescKindNone DescKind = 0
	// DESCKIND_FUNCDESC:  MUST be set if the name is bound to a method or property accessor
	// method. MUST NOT be set if the name is bound to an ODL dispinterface property.
	DescKindFuncDesc DescKind = 1
	// DESCKIND_VARDESC:  MUST be set if the name is bound to a data element or ODL dispinterface
	// property.
	DescKindVarDesc DescKind = 2
	// DESCKIND_TYPECOMP:  MUST be set if the name is bound to an enumeration or module.
	DescKindTypeComp DescKind = 3
	// DESCKIND_IMPLICITAPPOBJ:  MUST be set if the name is bound to an appobject coclass
	// (see section 2.2.49.8) or a member of its default nonsource interface (also see 2.2.49.8).
	DescKindImplicitAppObject DescKind = 4
)

func (o DescKind) String() string {
	switch o {
	case DescKindNone:
		return "DescKindNone"
	case DescKindFuncDesc:
		return "DescKindFuncDesc"
	case DescKindVarDesc:
		return "DescKindVarDesc"
	case DescKindTypeComp:
		return "DescKindTypeComp"
	case DescKindImplicitAppObject:
		return "DescKindImplicitAppObject"
	}
	return "Invalid"
}

// FlaggedWordBlob structure represents FLAGGED_WORD_BLOB RPC structure.
//
// The FLAGGED_WORD_BLOB structure defines a type for transferring length-prefixed data.
type FlaggedWordBlob struct {
	// cBytes:  MUST be the size, in bytes, of the asData array.
	//
	// Note  A value of 0xFFFFFFFF MUST be considered as representing a null BSTR.
	BytesCount uint32 `idl:"name:cBytes" json:"bytes_count"`
	// clSize:  MUST be the total number of unsigned shorts in the asData array. This value
	// MUST be half the value of cBytes, rounded up, unless this is a null  BSTR. In the
	// latter case, a value of 0 MUST be used.
	Size uint32 `idl:"name:clSize" json:"size"`
	// asData:  An array of unsigned shorts. If clSize is 0, asData MUST not contain any
	// elements.
	//
	// Data of this type MUST be marshaled as specified in [C706], section 14, with the
	// exception that it MUST be marshaled by using a little-endian data representation
	// regardless of the data representation format label. For more information, see [C706]
	// section 14.2.5.
	Data string `idl:"name:asData;size_is:(clSize)" json:"data"`
}

func (o *FlaggedWordBlob) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != "" && o.Size == 0 {
		o.Size = uint32(ndr.UTF16Len(o.Data))
	}
	if o.BytesCount == uint32(0) {
		_exprclSize := uint32(0)
		if o.Size != 0 {
			_exprclSize = uint32((o.Size * 2))
		} else {
			_exprclSize = uint32(4294967295)
		}
		o.BytesCount = uint32(_exprclSize)
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *FlaggedWordBlob) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.Size)
	return []uint64{
		dimSize1,
	}
}
func (o *FlaggedWordBlob) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.BytesCount); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	_Data_buf := utf16.Encode([]rune(o.Data))
	if uint64(len(_Data_buf)) > sizeInfo[0] {
		_Data_buf = _Data_buf[:sizeInfo[0]]
	}
	for i1 := range _Data_buf {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(_Data_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_Data_buf); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *FlaggedWordBlob) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.BytesCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.Size > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.Size)
	}
	var _Data_buf []uint16
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array _Data_buf", sizeInfo[0])
	}
	_Data_buf = make([]uint16, sizeInfo[0])
	for i1 := range _Data_buf {
		i1 := i1
		if err := w.ReadData(&_Data_buf[i1]); err != nil {
			return err
		}
	}
	o.Data = strings.TrimRight(string(utf16.Decode(_Data_buf)), ndr.ZeroString)
	return nil
}

// String structure represents _BSTR RPC structure.
type String FlaggedWordBlob

func (o *String) FlaggedWordBlob() *FlaggedWordBlob { return (*FlaggedWordBlob)(o) }

func (o *String) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != "" && o.Size == 0 {
		o.Size = uint32(ndr.UTF16Len(o.Data))
	}
	if o.BytesCount == uint32(0) {
		_exprclSize := uint32(0)
		if o.Size != 0 {
			_exprclSize = uint32((o.Size * 2))
		} else {
			_exprclSize = uint32(4294967295)
		}
		o.BytesCount = uint32(_exprclSize)
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *String) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.Size)
	return []uint64{
		dimSize1,
	}
}
func (o *String) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteData(o.BytesCount); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	_Data_buf := utf16.Encode([]rune(o.Data))
	if uint64(len(_Data_buf)) > sizeInfo[0] {
		_Data_buf = _Data_buf[:sizeInfo[0]]
	}
	for i1 := range _Data_buf {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if err := w.WriteData(_Data_buf[i1]); err != nil {
			return err
		}
	}
	for i1 := len(_Data_buf); uint64(i1) < sizeInfo[0]; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *String) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadData(&o.BytesCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.Size > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.Size)
	}
	var _Data_buf []uint16
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array _Data_buf", sizeInfo[0])
	}
	_Data_buf = make([]uint16, sizeInfo[0])
	for i1 := range _Data_buf {
		i1 := i1
		if err := w.ReadData(&_Data_buf[i1]); err != nil {
			return err
		}
	}
	o.Data = strings.TrimRight(string(utf16.Decode(_Data_buf)), ndr.ZeroString)
	return nil
}

// Currency structure represents CURRENCY RPC structure.
//
// The CURRENCY type specifies currency information. It is represented as an 8-byte
// integer, scaled by 10,000, to give a fixed-point number with 15 digits to the left
// of the decimal point, and four digits to the right. This representation provides
// a range of 922337203685477.5807 to –922337203685477.5808. For example, $5.25 is
// stored as the value 52500.
type Currency struct {
	Int64 int64 `idl:"name:int64" json:"int64"`
}

func (o *Currency) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Currency) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Int64); err != nil {
		return err
	}
	return nil
}
func (o *Currency) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Int64); err != nil {
		return err
	}
	return nil
}

// Decimal structure represents DECIMAL RPC structure.
//
// The DECIMAL structure specifies a sign and scale for a number. Decimal variables
// are represented as 96-bit unsigned integers that are scaled by a variable power of
// 10.
type Decimal struct {
	// wReserved:  MUST be set to 0 and MUST be ignored by the recipient.
	_ uint16 `idl:"name:wReserved"`
	// scale:  MUST be the power of 10 by which to divide the 96-bit integer represented
	// by Hi32 * 2^64 + Lo64. The value MUST be in the range of 0 to 28, inclusive.
	//
	//	+--------+-------------------------------------------+
	//	|        |                                           |
	//	| VALUE  |                  MEANING                  |
	//	|        |                                           |
	//	+--------+-------------------------------------------+
	//	+--------+-------------------------------------------+
	//	| 0 — 28 | Order of magnitude of the decimal number. |
	//	+--------+-------------------------------------------+
	Scale uint8 `idl:"name:scale" json:"scale"`
	// sign:  MUST equal one of the following values.
	//
	//	+-------+----------------------------------------+
	//	|       |                                        |
	//	| VALUE |                MEANING                 |
	//	|       |                                        |
	//	+-------+----------------------------------------+
	//	+-------+----------------------------------------+
	//	|     0 | The decimal contains a positive value. |
	//	+-------+----------------------------------------+
	//	| 0x80  | The decimal contains a negative value. |
	//	+-------+----------------------------------------+
	Sign uint8 `idl:"name:sign" json:"sign"`
	// Hi32:  MUST be the high 32 bits of the 96-bit integer that is scaled and signed to
	// represent the final DECIMAL value.
	Hi32 uint32 `idl:"name:Hi32" json:"hi32"`
	// Lo64:  MUST be the low 64 bits of the 96-bit integer that is scaled and signed to
	// represent the final DECIMAL value.
	Lo64 uint64 `idl:"name:Lo64" json:"lo64"`
}

func (o *Decimal) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Decimal) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	// reserved wReserved
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.Scale); err != nil {
		return err
	}
	if err := w.WriteData(o.Sign); err != nil {
		return err
	}
	if err := w.WriteData(o.Hi32); err != nil {
		return err
	}
	if err := w.WriteData(o.Lo64); err != nil {
		return err
	}
	return nil
}
func (o *Decimal) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	// reserved wReserved
	var _wReserved uint16
	if err := w.ReadData(&_wReserved); err != nil {
		return err
	}
	if err := w.ReadData(&o.Scale); err != nil {
		return err
	}
	if err := w.ReadData(&o.Sign); err != nil {
		return err
	}
	if err := w.ReadData(&o.Hi32); err != nil {
		return err
	}
	if err := w.ReadData(&o.Lo64); err != nil {
		return err
	}
	return nil
}

// Record structure represents _BRECORD RPC structure.
type Record struct {
	Flags      uint32                 `idl:"name:fFlags" json:"flags"`
	Size       uint32                 `idl:"name:clSize" json:"size"`
	RecordInfo *dcom.InterfacePointer `idl:"name:pRecInfo" json:"record_info"`
	Record     []byte                 `idl:"name:pRecord;size_is:(clSize)" json:"record"`
}

func (o *Record) xxx_PreparePayload(ctx context.Context) error {
	if o.Record != nil && o.Size == 0 {
		o.Size = uint32(len(o.Record))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if o.RecordInfo != nil {
		_ptr_pRecInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.RecordInfo != nil {
				if err := o.RecordInfo.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dcom.InterfacePointer{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.RecordInfo, _ptr_pRecInfo); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Record != nil || o.Size > 0 {
		_ptr_pRecord := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Size)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Record {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Record[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Record); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Record, _ptr_pRecord); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Record) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	_ptr_pRecInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.RecordInfo == nil {
			o.RecordInfo = &dcom.InterfacePointer{}
		}
		if err := o.RecordInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pRecInfo := func(ptr interface{}) { o.RecordInfo = *ptr.(**dcom.InterfacePointer) }
	if err := w.ReadPointer(&o.RecordInfo, _s_pRecInfo, _ptr_pRecInfo); err != nil {
		return err
	}
	_ptr_pRecord := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Size > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Size)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Record", sizeInfo[0])
		}
		o.Record = make([]byte, sizeInfo[0])
		for i1 := range o.Record {
			i1 := i1
			if err := w.ReadData(&o.Record[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pRecord := func(ptr interface{}) { o.Record = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Record, _s_pRecord, _ptr_pRecord); err != nil {
		return err
	}
	return nil
}

// Variant structure represents _VARIANT RPC structure.
type Variant struct {
	Size     uint32            `idl:"name:clSize" json:"size"`
	_        uint32            `idl:"name:rpcReserved"`
	VT       uint16            `idl:"name:vt" json:"vt"`
	_        uint16            `idl:"name:wReserved1"`
	_        uint16            `idl:"name:wReserved2"`
	_        uint16            `idl:"name:wReserved3"`
	VarUnion *Variant_VarUnion `idl:"name:_varUnion;switch_is:vt" json:"var_union"`
}

func (o *Variant) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	// reserved rpcReserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	if err := w.WriteData(o.VT); err != nil {
		return err
	}
	// reserved wReserved1
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	// reserved wReserved2
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	// reserved wReserved3
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	_swVarUnion := uint32(o.VT)
	if o.VarUnion != nil {
		if err := o.VarUnion.MarshalUnionNDR(ctx, w, _swVarUnion); err != nil {
			return err
		}
	} else {
		if err := (&Variant_VarUnion{}).MarshalUnionNDR(ctx, w, _swVarUnion); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	// reserved rpcReserved
	var _rpcReserved uint32
	if err := w.ReadData(&_rpcReserved); err != nil {
		return err
	}
	if err := w.ReadData(&o.VT); err != nil {
		return err
	}
	// reserved wReserved1
	var _wReserved1 uint16
	if err := w.ReadData(&_wReserved1); err != nil {
		return err
	}
	// reserved wReserved2
	var _wReserved2 uint16
	if err := w.ReadData(&_wReserved2); err != nil {
		return err
	}
	// reserved wReserved3
	var _wReserved3 uint16
	if err := w.ReadData(&_wReserved3); err != nil {
		return err
	}
	if o.VarUnion == nil {
		o.VarUnion = &Variant_VarUnion{}
	}
	_swVarUnion := uint32(o.VT)
	if err := o.VarUnion.UnmarshalUnionNDR(ctx, w, _swVarUnion); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion structure represents _VARIANT union anonymous member.
type Variant_VarUnion struct {
	// Types that are assignable to Value
	//
	// *Variant_VarUnion_LongLongValue
	// *Variant_VarUnion_Long
	// *Variant_VarUnion_Byte
	// *Variant_VarUnion_Short
	// *Variant_VarUnion_Float
	// *Variant_VarUnion_Double
	// *Variant_VarUnion_Bool
	// *Variant_VarUnion_HResult
	// *Variant_VarUnion_Currency
	// *Variant_VarUnion_Date
	// *Variant_VarUnion_BSTR
	// *Variant_VarUnion_IUnknown
	// *Variant_VarUnion_IDispatch
	// *Variant_VarUnion_SafeArray
	// *Variant_VarUnion_Brecord
	// *Variant_VarUnion_BytePtr
	// *Variant_VarUnion_ShortPtr
	// *Variant_VarUnion_LongPtr
	// *Variant_VarUnion_LongLongPtr
	// *Variant_VarUnion_FloatPtr
	// *Variant_VarUnion_DoublePtr
	// *Variant_VarUnion_BoolPtr
	// *Variant_VarUnion_HResultPtr
	// *Variant_VarUnion_CurrencyPtr
	// *Variant_VarUnion_DatePtr
	// *Variant_VarUnion_BSTRPtr
	// *Variant_VarUnion_IUnknownPtr
	// *Variant_VarUnion_IDispatchPtr
	// *Variant_VarUnion_SafeArrayPtr
	// *Variant_VarUnion_VariantPtr
	// *Variant_VarUnion_Char
	// *Variant_VarUnion_Ushort
	// *Variant_VarUnion_Ulong
	// *Variant_VarUnion_UlongLong
	// *Variant_VarUnion_Int
	// *Variant_VarUnion_Uint
	// *Variant_VarUnion_Decimal
	// *Variant_VarUnion_CharPtr
	// *Variant_VarUnion_UshortPtr
	// *Variant_VarUnion_UlongPtr
	// *Variant_VarUnion_UlongLongPtr
	// *Variant_VarUnion_IntPtr
	// *Variant_VarUnion_UintPtr
	// *Variant_VarUnion_DecimalPtr
	// *Variant_VarUnion_0
	// *Variant_VarUnion_1
	Value is_Variant_VarUnion `json:"value"`
}

func (o *Variant_VarUnion) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *Variant_VarUnion_LongLongValue:
		if value != nil {
			return value.LongLongValue
		}
	case *Variant_VarUnion_Long:
		if value != nil {
			return value.Long
		}
	case *Variant_VarUnion_Byte:
		if value != nil {
			return value.Byte
		}
	case *Variant_VarUnion_Short:
		if value != nil {
			return value.Short
		}
	case *Variant_VarUnion_Float:
		if value != nil {
			return value.Float
		}
	case *Variant_VarUnion_Double:
		if value != nil {
			return value.Double
		}
	case *Variant_VarUnion_Bool:
		if value != nil {
			return value.Bool
		}
	case *Variant_VarUnion_HResult:
		if value != nil {
			return value.HResult
		}
	case *Variant_VarUnion_Currency:
		if value != nil {
			return value.Currency
		}
	case *Variant_VarUnion_Date:
		if value != nil {
			return value.Date
		}
	case *Variant_VarUnion_BSTR:
		if value != nil {
			return value.BSTR
		}
	case *Variant_VarUnion_IUnknown:
		if value != nil {
			return value.IUnknown
		}
	case *Variant_VarUnion_IDispatch:
		if value != nil {
			return value.IDispatch
		}
	case *Variant_VarUnion_SafeArray:
		if value != nil {
			return value.SafeArray
		}
	case *Variant_VarUnion_Brecord:
		if value != nil {
			return value.Brecord
		}
	case *Variant_VarUnion_BytePtr:
		if value != nil {
			return value.BytePtr
		}
	case *Variant_VarUnion_ShortPtr:
		if value != nil {
			return value.ShortPtr
		}
	case *Variant_VarUnion_LongPtr:
		if value != nil {
			return value.LongPtr
		}
	case *Variant_VarUnion_LongLongPtr:
		if value != nil {
			return value.LongLongPtr
		}
	case *Variant_VarUnion_FloatPtr:
		if value != nil {
			return value.FloatPtr
		}
	case *Variant_VarUnion_DoublePtr:
		if value != nil {
			return value.DoublePtr
		}
	case *Variant_VarUnion_BoolPtr:
		if value != nil {
			return value.BoolPtr
		}
	case *Variant_VarUnion_HResultPtr:
		if value != nil {
			return value.HResultPtr
		}
	case *Variant_VarUnion_CurrencyPtr:
		if value != nil {
			return value.CurrencyPtr
		}
	case *Variant_VarUnion_DatePtr:
		if value != nil {
			return value.DatePtr
		}
	case *Variant_VarUnion_BSTRPtr:
		if value != nil {
			return value.BSTRPtr
		}
	case *Variant_VarUnion_IUnknownPtr:
		if value != nil {
			return value.IUnknownPtr
		}
	case *Variant_VarUnion_IDispatchPtr:
		if value != nil {
			return value.IDispatchPtr
		}
	case *Variant_VarUnion_SafeArrayPtr:
		if value != nil {
			return value.SafeArrayPtr
		}
	case *Variant_VarUnion_VariantPtr:
		if value != nil {
			return value.VariantPtr
		}
	case *Variant_VarUnion_Char:
		if value != nil {
			return value.Char
		}
	case *Variant_VarUnion_Ushort:
		if value != nil {
			return value.Ushort
		}
	case *Variant_VarUnion_Ulong:
		if value != nil {
			return value.Ulong
		}
	case *Variant_VarUnion_UlongLong:
		if value != nil {
			return value.UlongLong
		}
	case *Variant_VarUnion_Int:
		if value != nil {
			return value.Int
		}
	case *Variant_VarUnion_Uint:
		if value != nil {
			return value.Uint
		}
	case *Variant_VarUnion_Decimal:
		if value != nil {
			return value.Decimal
		}
	case *Variant_VarUnion_CharPtr:
		if value != nil {
			return value.CharPtr
		}
	case *Variant_VarUnion_UshortPtr:
		if value != nil {
			return value.UshortPtr
		}
	case *Variant_VarUnion_UlongPtr:
		if value != nil {
			return value.UlongPtr
		}
	case *Variant_VarUnion_UlongLongPtr:
		if value != nil {
			return value.UlongLongPtr
		}
	case *Variant_VarUnion_IntPtr:
		if value != nil {
			return value.IntPtr
		}
	case *Variant_VarUnion_UintPtr:
		if value != nil {
			return value.UintPtr
		}
	case *Variant_VarUnion_DecimalPtr:
		if value != nil {
			return value.DecimalPtr
		}
	}
	return nil
}

type is_Variant_VarUnion interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_Variant_VarUnion()
}

func (o *Variant_VarUnion) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *Variant_VarUnion_LongLongValue:
		return uint32(20)
	case *Variant_VarUnion_Long:
		return uint32(3)
	case *Variant_VarUnion_Byte:
		return uint32(17)
	case *Variant_VarUnion_Short:
		return uint32(2)
	case *Variant_VarUnion_Float:
		return uint32(4)
	case *Variant_VarUnion_Double:
		return uint32(5)
	case *Variant_VarUnion_Bool:
		return uint32(11)
	case *Variant_VarUnion_HResult:
		return uint32(10)
	case *Variant_VarUnion_Currency:
		return uint32(6)
	case *Variant_VarUnion_Date:
		return uint32(7)
	case *Variant_VarUnion_BSTR:
		return uint32(8)
	case *Variant_VarUnion_IUnknown:
		return uint32(13)
	case *Variant_VarUnion_IDispatch:
		return uint32(9)
	case *Variant_VarUnion_SafeArray:
		return uint32(8192)
	case *Variant_VarUnion_Brecord:
		switch sw {
		case uint32(36),
			uint32(16420):
			return sw
		}
		return uint32(36)
	case *Variant_VarUnion_BytePtr:
		return uint32(16401)
	case *Variant_VarUnion_ShortPtr:
		return uint32(16386)
	case *Variant_VarUnion_LongPtr:
		return uint32(16387)
	case *Variant_VarUnion_LongLongPtr:
		return uint32(16404)
	case *Variant_VarUnion_FloatPtr:
		return uint32(16388)
	case *Variant_VarUnion_DoublePtr:
		return uint32(16389)
	case *Variant_VarUnion_BoolPtr:
		return uint32(16395)
	case *Variant_VarUnion_HResultPtr:
		return uint32(16394)
	case *Variant_VarUnion_CurrencyPtr:
		return uint32(16390)
	case *Variant_VarUnion_DatePtr:
		return uint32(16391)
	case *Variant_VarUnion_BSTRPtr:
		return uint32(16392)
	case *Variant_VarUnion_IUnknownPtr:
		return uint32(16397)
	case *Variant_VarUnion_IDispatchPtr:
		return uint32(16393)
	case *Variant_VarUnion_SafeArrayPtr:
		return uint32(24576)
	case *Variant_VarUnion_VariantPtr:
		return uint32(16396)
	case *Variant_VarUnion_Char:
		return uint32(16)
	case *Variant_VarUnion_Ushort:
		return uint32(18)
	case *Variant_VarUnion_Ulong:
		return uint32(19)
	case *Variant_VarUnion_UlongLong:
		return uint32(21)
	case *Variant_VarUnion_Int:
		return uint32(22)
	case *Variant_VarUnion_Uint:
		return uint32(23)
	case *Variant_VarUnion_Decimal:
		return uint32(14)
	case *Variant_VarUnion_CharPtr:
		return uint32(16400)
	case *Variant_VarUnion_UshortPtr:
		return uint32(16402)
	case *Variant_VarUnion_UlongPtr:
		return uint32(16403)
	case *Variant_VarUnion_UlongLongPtr:
		return uint32(16405)
	case *Variant_VarUnion_IntPtr:
		return uint32(16406)
	case *Variant_VarUnion_UintPtr:
		return uint32(16407)
	case *Variant_VarUnion_DecimalPtr:
		return uint32(16398)
	case *Variant_VarUnion_0:
		return uint32(0)
	case *Variant_VarUnion_1:
		return uint32(1)
	}
	return uint32(0)
}

func (o *Variant_VarUnion) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint32(20):
		_o, _ := o.Value.(*Variant_VarUnion_LongLongValue)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_LongLongValue{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(3):
		_o, _ := o.Value.(*Variant_VarUnion_Long)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_Long{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(17):
		_o, _ := o.Value.(*Variant_VarUnion_Byte)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_Byte{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*Variant_VarUnion_Short)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_Short{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(4):
		_o, _ := o.Value.(*Variant_VarUnion_Float)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_Float{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(5):
		_o, _ := o.Value.(*Variant_VarUnion_Double)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_Double{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(11):
		_o, _ := o.Value.(*Variant_VarUnion_Bool)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_Bool{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(10):
		_o, _ := o.Value.(*Variant_VarUnion_HResult)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_HResult{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(6):
		_o, _ := o.Value.(*Variant_VarUnion_Currency)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_Currency{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(7):
		_o, _ := o.Value.(*Variant_VarUnion_Date)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_Date{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(8):
		_o, _ := o.Value.(*Variant_VarUnion_BSTR)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_BSTR{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(13):
		_o, _ := o.Value.(*Variant_VarUnion_IUnknown)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_IUnknown{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(9):
		_o, _ := o.Value.(*Variant_VarUnion_IDispatch)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_IDispatch{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(8192):
		_o, _ := o.Value.(*Variant_VarUnion_SafeArray)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_SafeArray{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(36),
		uint32(16420):
		_o, _ := o.Value.(*Variant_VarUnion_Brecord)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_Brecord{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(16401):
		_o, _ := o.Value.(*Variant_VarUnion_BytePtr)
		if _o != nil {
			_ptr_o := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := _o.MarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&_o, _ptr_o); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	case uint32(16386):
		_o, _ := o.Value.(*Variant_VarUnion_ShortPtr)
		if _o != nil {
			_ptr_o := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := _o.MarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&_o, _ptr_o); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	case uint32(16387):
		_o, _ := o.Value.(*Variant_VarUnion_LongPtr)
		if _o != nil {
			_ptr_o := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := _o.MarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&_o, _ptr_o); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	case uint32(16404):
		_o, _ := o.Value.(*Variant_VarUnion_LongLongPtr)
		if _o != nil {
			_ptr_o := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := _o.MarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&_o, _ptr_o); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	case uint32(16388):
		_o, _ := o.Value.(*Variant_VarUnion_FloatPtr)
		if _o != nil {
			_ptr_o := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := _o.MarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&_o, _ptr_o); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	case uint32(16389):
		_o, _ := o.Value.(*Variant_VarUnion_DoublePtr)
		if _o != nil {
			_ptr_o := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := _o.MarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&_o, _ptr_o); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	case uint32(16395):
		_o, _ := o.Value.(*Variant_VarUnion_BoolPtr)
		if _o != nil {
			_ptr_o := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := _o.MarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&_o, _ptr_o); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	case uint32(16394):
		_o, _ := o.Value.(*Variant_VarUnion_HResultPtr)
		if _o != nil {
			_ptr_o := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := _o.MarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&_o, _ptr_o); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	case uint32(16390):
		_o, _ := o.Value.(*Variant_VarUnion_CurrencyPtr)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_CurrencyPtr{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(16391):
		_o, _ := o.Value.(*Variant_VarUnion_DatePtr)
		if _o != nil {
			_ptr_o := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := _o.MarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&_o, _ptr_o); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	case uint32(16392):
		_o, _ := o.Value.(*Variant_VarUnion_BSTRPtr)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_BSTRPtr{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(16397):
		_o, _ := o.Value.(*Variant_VarUnion_IUnknownPtr)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_IUnknownPtr{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(16393):
		_o, _ := o.Value.(*Variant_VarUnion_IDispatchPtr)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_IDispatchPtr{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(24576):
		_o, _ := o.Value.(*Variant_VarUnion_SafeArrayPtr)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_SafeArrayPtr{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(16396):
		_o, _ := o.Value.(*Variant_VarUnion_VariantPtr)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_VariantPtr{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(16):
		_o, _ := o.Value.(*Variant_VarUnion_Char)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_Char{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(18):
		_o, _ := o.Value.(*Variant_VarUnion_Ushort)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_Ushort{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(19):
		_o, _ := o.Value.(*Variant_VarUnion_Ulong)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_Ulong{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(21):
		_o, _ := o.Value.(*Variant_VarUnion_UlongLong)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_UlongLong{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(22):
		_o, _ := o.Value.(*Variant_VarUnion_Int)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_Int{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(23):
		_o, _ := o.Value.(*Variant_VarUnion_Uint)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_Uint{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(14):
		_o, _ := o.Value.(*Variant_VarUnion_Decimal)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_Decimal{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(16400):
		_o, _ := o.Value.(*Variant_VarUnion_CharPtr)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_CharPtr{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(16402):
		_o, _ := o.Value.(*Variant_VarUnion_UshortPtr)
		if _o != nil {
			_ptr_o := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := _o.MarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&_o, _ptr_o); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	case uint32(16403):
		_o, _ := o.Value.(*Variant_VarUnion_UlongPtr)
		if _o != nil {
			_ptr_o := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := _o.MarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&_o, _ptr_o); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	case uint32(16405):
		_o, _ := o.Value.(*Variant_VarUnion_UlongLongPtr)
		if _o != nil {
			_ptr_o := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := _o.MarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&_o, _ptr_o); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	case uint32(16406):
		_o, _ := o.Value.(*Variant_VarUnion_IntPtr)
		if _o != nil {
			_ptr_o := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := _o.MarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&_o, _ptr_o); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	case uint32(16407):
		_o, _ := o.Value.(*Variant_VarUnion_UintPtr)
		if _o != nil {
			_ptr_o := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := _o.MarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&_o, _ptr_o); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
	case uint32(16398):
		_o, _ := o.Value.(*Variant_VarUnion_DecimalPtr)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Variant_VarUnion_DecimalPtr{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(0):
	case uint32(1):
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *Variant_VarUnion) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(8); err != nil {
		return err
	}
	switch sw {
	case uint32(20):
		o.Value = &Variant_VarUnion_LongLongValue{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(3):
		o.Value = &Variant_VarUnion_Long{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(17):
		o.Value = &Variant_VarUnion_Byte{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &Variant_VarUnion_Short{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(4):
		o.Value = &Variant_VarUnion_Float{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(5):
		o.Value = &Variant_VarUnion_Double{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(11):
		o.Value = &Variant_VarUnion_Bool{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(10):
		o.Value = &Variant_VarUnion_HResult{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(6):
		o.Value = &Variant_VarUnion_Currency{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(7):
		o.Value = &Variant_VarUnion_Date{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(8):
		o.Value = &Variant_VarUnion_BSTR{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(13):
		o.Value = &Variant_VarUnion_IUnknown{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(9):
		o.Value = &Variant_VarUnion_IDispatch{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(8192):
		o.Value = &Variant_VarUnion_SafeArray{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(36),
		uint32(16420):
		o.Value = &Variant_VarUnion_Brecord{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(16401):
		_ptr_o := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			o.Value = &Variant_VarUnion_BytePtr{}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_o := func(ptr interface{}) { o.Value = *ptr.(**Variant_VarUnion_BytePtr) }
		if err := w.ReadPointer(&o.Value, _s_o, _ptr_o); err != nil {
			return err
		}
	case uint32(16386):
		_ptr_o := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			o.Value = &Variant_VarUnion_ShortPtr{}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_o := func(ptr interface{}) { o.Value = *ptr.(**Variant_VarUnion_ShortPtr) }
		if err := w.ReadPointer(&o.Value, _s_o, _ptr_o); err != nil {
			return err
		}
	case uint32(16387):
		_ptr_o := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			o.Value = &Variant_VarUnion_LongPtr{}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_o := func(ptr interface{}) { o.Value = *ptr.(**Variant_VarUnion_LongPtr) }
		if err := w.ReadPointer(&o.Value, _s_o, _ptr_o); err != nil {
			return err
		}
	case uint32(16404):
		_ptr_o := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			o.Value = &Variant_VarUnion_LongLongPtr{}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_o := func(ptr interface{}) { o.Value = *ptr.(**Variant_VarUnion_LongLongPtr) }
		if err := w.ReadPointer(&o.Value, _s_o, _ptr_o); err != nil {
			return err
		}
	case uint32(16388):
		_ptr_o := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			o.Value = &Variant_VarUnion_FloatPtr{}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_o := func(ptr interface{}) { o.Value = *ptr.(**Variant_VarUnion_FloatPtr) }
		if err := w.ReadPointer(&o.Value, _s_o, _ptr_o); err != nil {
			return err
		}
	case uint32(16389):
		_ptr_o := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			o.Value = &Variant_VarUnion_DoublePtr{}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_o := func(ptr interface{}) { o.Value = *ptr.(**Variant_VarUnion_DoublePtr) }
		if err := w.ReadPointer(&o.Value, _s_o, _ptr_o); err != nil {
			return err
		}
	case uint32(16395):
		_ptr_o := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			o.Value = &Variant_VarUnion_BoolPtr{}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_o := func(ptr interface{}) { o.Value = *ptr.(**Variant_VarUnion_BoolPtr) }
		if err := w.ReadPointer(&o.Value, _s_o, _ptr_o); err != nil {
			return err
		}
	case uint32(16394):
		_ptr_o := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			o.Value = &Variant_VarUnion_HResultPtr{}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_o := func(ptr interface{}) { o.Value = *ptr.(**Variant_VarUnion_HResultPtr) }
		if err := w.ReadPointer(&o.Value, _s_o, _ptr_o); err != nil {
			return err
		}
	case uint32(16390):
		o.Value = &Variant_VarUnion_CurrencyPtr{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(16391):
		_ptr_o := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			o.Value = &Variant_VarUnion_DatePtr{}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_o := func(ptr interface{}) { o.Value = *ptr.(**Variant_VarUnion_DatePtr) }
		if err := w.ReadPointer(&o.Value, _s_o, _ptr_o); err != nil {
			return err
		}
	case uint32(16392):
		o.Value = &Variant_VarUnion_BSTRPtr{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(16397):
		o.Value = &Variant_VarUnion_IUnknownPtr{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(16393):
		o.Value = &Variant_VarUnion_IDispatchPtr{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(24576):
		o.Value = &Variant_VarUnion_SafeArrayPtr{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(16396):
		o.Value = &Variant_VarUnion_VariantPtr{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(16):
		o.Value = &Variant_VarUnion_Char{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(18):
		o.Value = &Variant_VarUnion_Ushort{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(19):
		o.Value = &Variant_VarUnion_Ulong{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(21):
		o.Value = &Variant_VarUnion_UlongLong{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(22):
		o.Value = &Variant_VarUnion_Int{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(23):
		o.Value = &Variant_VarUnion_Uint{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(14):
		o.Value = &Variant_VarUnion_Decimal{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(16400):
		o.Value = &Variant_VarUnion_CharPtr{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(16402):
		_ptr_o := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			o.Value = &Variant_VarUnion_UshortPtr{}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_o := func(ptr interface{}) { o.Value = *ptr.(**Variant_VarUnion_UshortPtr) }
		if err := w.ReadPointer(&o.Value, _s_o, _ptr_o); err != nil {
			return err
		}
	case uint32(16403):
		_ptr_o := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			o.Value = &Variant_VarUnion_UlongPtr{}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_o := func(ptr interface{}) { o.Value = *ptr.(**Variant_VarUnion_UlongPtr) }
		if err := w.ReadPointer(&o.Value, _s_o, _ptr_o); err != nil {
			return err
		}
	case uint32(16405):
		_ptr_o := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			o.Value = &Variant_VarUnion_UlongLongPtr{}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_o := func(ptr interface{}) { o.Value = *ptr.(**Variant_VarUnion_UlongLongPtr) }
		if err := w.ReadPointer(&o.Value, _s_o, _ptr_o); err != nil {
			return err
		}
	case uint32(16406):
		_ptr_o := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			o.Value = &Variant_VarUnion_IntPtr{}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_o := func(ptr interface{}) { o.Value = *ptr.(**Variant_VarUnion_IntPtr) }
		if err := w.ReadPointer(&o.Value, _s_o, _ptr_o); err != nil {
			return err
		}
	case uint32(16407):
		_ptr_o := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			o.Value = &Variant_VarUnion_UintPtr{}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_o := func(ptr interface{}) { o.Value = *ptr.(**Variant_VarUnion_UintPtr) }
		if err := w.ReadPointer(&o.Value, _s_o, _ptr_o); err != nil {
			return err
		}
	case uint32(16398):
		o.Value = &Variant_VarUnion_DecimalPtr{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(0):
		o.Value = &Variant_VarUnion_0{}
	case uint32(1):
		o.Value = &Variant_VarUnion_1{}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// Variant_VarUnion_LongLongValue structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 20
type Variant_VarUnion_LongLongValue struct {
	LongLongValue int64 `idl:"name:llVal" json:"long_long_value"`
}

func (*Variant_VarUnion_LongLongValue) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_LongLongValue) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.LongLongValue); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_LongLongValue) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.LongLongValue); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_Long structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 3
type Variant_VarUnion_Long struct {
	Long int32 `idl:"name:lVal" json:"long"`
}

func (*Variant_VarUnion_Long) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_Long) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Long); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_Long) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Long); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_Byte structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 17
type Variant_VarUnion_Byte struct {
	Byte uint8 `idl:"name:bVal" json:"byte"`
}

func (*Variant_VarUnion_Byte) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_Byte) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Byte); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_Byte) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Byte); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_Short structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 2
type Variant_VarUnion_Short struct {
	Short int16 `idl:"name:iVal" json:"short"`
}

func (*Variant_VarUnion_Short) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_Short) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Short); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_Short) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Short); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_Float structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 4
type Variant_VarUnion_Float struct {
	Float float32 `idl:"name:fltVal" json:"float"`
}

func (*Variant_VarUnion_Float) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_Float) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Float); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_Float) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Float); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_Double structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 5
type Variant_VarUnion_Double struct {
	Double float64 `idl:"name:dblVal" json:"double"`
}

func (*Variant_VarUnion_Double) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_Double) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Double); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_Double) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Double); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_Bool structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 11
type Variant_VarUnion_Bool struct {
	Bool int16 `idl:"name:boolVal" json:"bool"`
}

func (*Variant_VarUnion_Bool) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_Bool) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Bool); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_Bool) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Bool); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_HResult structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 10
type Variant_VarUnion_HResult struct {
	HResult int32 `idl:"name:scode" json:"h_result"`
}

func (*Variant_VarUnion_HResult) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_HResult) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.HResult); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_HResult) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.HResult); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_Currency structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 6
type Variant_VarUnion_Currency struct {
	Currency *Currency `idl:"name:cyVal" json:"currency"`
}

func (*Variant_VarUnion_Currency) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_Currency) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Currency != nil {
		if err := o.Currency.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Currency{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_VarUnion_Currency) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Currency == nil {
		o.Currency = &Currency{}
	}
	if err := o.Currency.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_Date structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 7
type Variant_VarUnion_Date struct {
	Date float64 `idl:"name:date" json:"date"`
}

func (*Variant_VarUnion_Date) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_Date) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Date); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_Date) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Date); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_BSTR structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 8
type Variant_VarUnion_BSTR struct {
	BSTR *String `idl:"name:bstrVal" json:"bstr"`
}

func (*Variant_VarUnion_BSTR) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_BSTR) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.BSTR != nil {
		_ptr_bstrVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.BSTR != nil {
				if err := o.BSTR.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&String{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.BSTR, _ptr_bstrVal); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_VarUnion_BSTR) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_bstrVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.BSTR == nil {
			o.BSTR = &String{}
		}
		if err := o.BSTR.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_bstrVal := func(ptr interface{}) { o.BSTR = *ptr.(**String) }
	if err := w.ReadPointer(&o.BSTR, _s_bstrVal, _ptr_bstrVal); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_IUnknown structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 13
type Variant_VarUnion_IUnknown struct {
	IUnknown *dcom.Unknown `idl:"name:punkVal" json:"iunknown"`
}

func (*Variant_VarUnion_IUnknown) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_IUnknown) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IUnknown != nil {
		_ptr_punkVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.IUnknown != nil {
				if err := o.IUnknown.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IUnknown, _ptr_punkVal); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_VarUnion_IUnknown) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_punkVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.IUnknown == nil {
			o.IUnknown = &dcom.Unknown{}
		}
		if err := o.IUnknown.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_punkVal := func(ptr interface{}) { o.IUnknown = *ptr.(**dcom.Unknown) }
	if err := w.ReadPointer(&o.IUnknown, _s_punkVal, _ptr_punkVal); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_IDispatch structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 9
type Variant_VarUnion_IDispatch struct {
	IDispatch *Dispatch `idl:"name:pdispVal" json:"idispatch"`
}

func (*Variant_VarUnion_IDispatch) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_IDispatch) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IDispatch != nil {
		_ptr_pdispVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.IDispatch != nil {
				if err := o.IDispatch.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Dispatch{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IDispatch, _ptr_pdispVal); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_VarUnion_IDispatch) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pdispVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.IDispatch == nil {
			o.IDispatch = &Dispatch{}
		}
		if err := o.IDispatch.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pdispVal := func(ptr interface{}) { o.IDispatch = *ptr.(**Dispatch) }
	if err := w.ReadPointer(&o.IDispatch, _s_pdispVal, _ptr_pdispVal); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_SafeArray structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 8192
type Variant_VarUnion_SafeArray struct {
	SafeArray *SafeArray `idl:"name:parray" json:"safe_array"`
}

func (*Variant_VarUnion_SafeArray) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_SafeArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SafeArray != nil {
		_ptr_parray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SafeArray != nil {
				_ptr_parray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.SafeArray != nil {
						if err := o.SafeArray.MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&SafeArray{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.SafeArray, _ptr_parray); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SafeArray, _ptr_parray); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_VarUnion_SafeArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_parray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_parray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SafeArray == nil {
				o.SafeArray = &SafeArray{}
			}
			if err := o.SafeArray.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_parray := func(ptr interface{}) { o.SafeArray = *ptr.(**SafeArray) }
		if err := w.ReadPointer(&o.SafeArray, _s_parray, _ptr_parray); err != nil {
			return err
		}
		return nil
	})
	_s_parray := func(ptr interface{}) { o.SafeArray = *ptr.(**SafeArray) }
	if err := w.ReadPointer(&o.SafeArray, _s_parray, _ptr_parray); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_Brecord structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 36, 16420
type Variant_VarUnion_Brecord struct {
	Brecord *Record `idl:"name:brecVal" json:"brecord"`
}

func (*Variant_VarUnion_Brecord) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_Brecord) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Brecord != nil {
		_ptr_brecVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Brecord != nil {
				if err := o.Brecord.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Record{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Brecord, _ptr_brecVal); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_VarUnion_Brecord) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_brecVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Brecord == nil {
			o.Brecord = &Record{}
		}
		if err := o.Brecord.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_brecVal := func(ptr interface{}) { o.Brecord = *ptr.(**Record) }
	if err := w.ReadPointer(&o.Brecord, _s_brecVal, _ptr_brecVal); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_BytePtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16401
type Variant_VarUnion_BytePtr struct {
	BytePtr uint8 `idl:"name:pbVal" json:"byte_ptr"`
}

func (*Variant_VarUnion_BytePtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_BytePtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.BytePtr); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_BytePtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.BytePtr); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_ShortPtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16386
type Variant_VarUnion_ShortPtr struct {
	ShortPtr int16 `idl:"name:piVal" json:"short_ptr"`
}

func (*Variant_VarUnion_ShortPtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_ShortPtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.ShortPtr); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_ShortPtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.ShortPtr); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_LongPtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16387
type Variant_VarUnion_LongPtr struct {
	LongPtr int32 `idl:"name:plVal" json:"long_ptr"`
}

func (*Variant_VarUnion_LongPtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_LongPtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.LongPtr); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_LongPtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.LongPtr); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_LongLongPtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16404
type Variant_VarUnion_LongLongPtr struct {
	LongLongPtr int64 `idl:"name:pllVal" json:"long_long_ptr"`
}

func (*Variant_VarUnion_LongLongPtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_LongLongPtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.LongLongPtr); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_LongLongPtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.LongLongPtr); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_FloatPtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16388
type Variant_VarUnion_FloatPtr struct {
	FloatPtr float32 `idl:"name:pfltVal" json:"float_ptr"`
}

func (*Variant_VarUnion_FloatPtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_FloatPtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.FloatPtr); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_FloatPtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.FloatPtr); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_DoublePtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16389
type Variant_VarUnion_DoublePtr struct {
	DoublePtr float64 `idl:"name:pdblVal" json:"double_ptr"`
}

func (*Variant_VarUnion_DoublePtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_DoublePtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.DoublePtr); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_DoublePtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.DoublePtr); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_BoolPtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16395
type Variant_VarUnion_BoolPtr struct {
	BoolPtr int16 `idl:"name:pboolVal" json:"bool_ptr"`
}

func (*Variant_VarUnion_BoolPtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_BoolPtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.BoolPtr); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_BoolPtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.BoolPtr); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_HResultPtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16394
type Variant_VarUnion_HResultPtr struct {
	HResultPtr int32 `idl:"name:pscode" json:"hresult_ptr"`
}

func (*Variant_VarUnion_HResultPtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_HResultPtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.HResultPtr); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_HResultPtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.HResultPtr); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_CurrencyPtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16390
type Variant_VarUnion_CurrencyPtr struct {
	CurrencyPtr *Currency `idl:"name:pcyVal" json:"currency_ptr"`
}

func (*Variant_VarUnion_CurrencyPtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_CurrencyPtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.CurrencyPtr != nil {
		_ptr_pcyVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.CurrencyPtr != nil {
				if err := o.CurrencyPtr.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Currency{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CurrencyPtr, _ptr_pcyVal); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_VarUnion_CurrencyPtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pcyVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.CurrencyPtr == nil {
			o.CurrencyPtr = &Currency{}
		}
		if err := o.CurrencyPtr.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pcyVal := func(ptr interface{}) { o.CurrencyPtr = *ptr.(**Currency) }
	if err := w.ReadPointer(&o.CurrencyPtr, _s_pcyVal, _ptr_pcyVal); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_DatePtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16391
type Variant_VarUnion_DatePtr struct {
	DatePtr float64 `idl:"name:pdate" json:"date_ptr"`
}

func (*Variant_VarUnion_DatePtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_DatePtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.DatePtr); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_DatePtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.DatePtr); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_BSTRPtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16392
type Variant_VarUnion_BSTRPtr struct {
	BSTRPtr *String `idl:"name:pbstrVal" json:"bstr_ptr"`
}

func (*Variant_VarUnion_BSTRPtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_BSTRPtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.BSTRPtr != nil {
		_ptr_pbstrVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.BSTRPtr != nil {
				_ptr_pbstrVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.BSTRPtr != nil {
						if err := o.BSTRPtr.MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&String{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.BSTRPtr, _ptr_pbstrVal); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.BSTRPtr, _ptr_pbstrVal); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_VarUnion_BSTRPtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pbstrVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_pbstrVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.BSTRPtr == nil {
				o.BSTRPtr = &String{}
			}
			if err := o.BSTRPtr.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrVal := func(ptr interface{}) { o.BSTRPtr = *ptr.(**String) }
		if err := w.ReadPointer(&o.BSTRPtr, _s_pbstrVal, _ptr_pbstrVal); err != nil {
			return err
		}
		return nil
	})
	_s_pbstrVal := func(ptr interface{}) { o.BSTRPtr = *ptr.(**String) }
	if err := w.ReadPointer(&o.BSTRPtr, _s_pbstrVal, _ptr_pbstrVal); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_IUnknownPtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16397
type Variant_VarUnion_IUnknownPtr struct {
	IUnknownPtr *dcom.Unknown `idl:"name:ppunkVal" json:"iunknown_ptr"`
}

func (*Variant_VarUnion_IUnknownPtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_IUnknownPtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IUnknownPtr != nil {
		_ptr_ppunkVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.IUnknownPtr != nil {
				_ptr_ppunkVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.IUnknownPtr != nil {
						if err := o.IUnknownPtr.MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.IUnknownPtr, _ptr_ppunkVal); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IUnknownPtr, _ptr_ppunkVal); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_VarUnion_IUnknownPtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ppunkVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_ppunkVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.IUnknownPtr == nil {
				o.IUnknownPtr = &dcom.Unknown{}
			}
			if err := o.IUnknownPtr.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppunkVal := func(ptr interface{}) { o.IUnknownPtr = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.IUnknownPtr, _s_ppunkVal, _ptr_ppunkVal); err != nil {
			return err
		}
		return nil
	})
	_s_ppunkVal := func(ptr interface{}) { o.IUnknownPtr = *ptr.(**dcom.Unknown) }
	if err := w.ReadPointer(&o.IUnknownPtr, _s_ppunkVal, _ptr_ppunkVal); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_IDispatchPtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16393
type Variant_VarUnion_IDispatchPtr struct {
	IDispatchPtr *Dispatch `idl:"name:ppdispVal" json:"idispatch_ptr"`
}

func (*Variant_VarUnion_IDispatchPtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_IDispatchPtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.IDispatchPtr != nil {
		_ptr_ppdispVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.IDispatchPtr != nil {
				_ptr_ppdispVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.IDispatchPtr != nil {
						if err := o.IDispatchPtr.MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&Dispatch{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.IDispatchPtr, _ptr_ppdispVal); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.IDispatchPtr, _ptr_ppdispVal); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_VarUnion_IDispatchPtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_ppdispVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_ppdispVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.IDispatchPtr == nil {
				o.IDispatchPtr = &Dispatch{}
			}
			if err := o.IDispatchPtr.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppdispVal := func(ptr interface{}) { o.IDispatchPtr = *ptr.(**Dispatch) }
		if err := w.ReadPointer(&o.IDispatchPtr, _s_ppdispVal, _ptr_ppdispVal); err != nil {
			return err
		}
		return nil
	})
	_s_ppdispVal := func(ptr interface{}) { o.IDispatchPtr = *ptr.(**Dispatch) }
	if err := w.ReadPointer(&o.IDispatchPtr, _s_ppdispVal, _ptr_ppdispVal); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_SafeArrayPtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 24576
type Variant_VarUnion_SafeArrayPtr struct {
	SafeArrayPtr *SafeArray `idl:"name:pparray" json:"safe_array_ptr"`
}

func (*Variant_VarUnion_SafeArrayPtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_SafeArrayPtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.SafeArrayPtr != nil {
		_ptr_pparray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.SafeArrayPtr != nil {
				_ptr_pparray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.SafeArrayPtr != nil {
						_ptr_pparray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if o.SafeArrayPtr != nil {
								if err := o.SafeArrayPtr.MarshalNDR(ctx, w); err != nil {
									return err
								}
							} else {
								if err := (&SafeArray{}).MarshalNDR(ctx, w); err != nil {
									return err
								}
							}
							return nil
						})
						if err := w.WritePointer(&o.SafeArrayPtr, _ptr_pparray); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.SafeArrayPtr, _ptr_pparray); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SafeArrayPtr, _ptr_pparray); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_VarUnion_SafeArrayPtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pparray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_pparray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_pparray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.SafeArrayPtr == nil {
					o.SafeArrayPtr = &SafeArray{}
				}
				if err := o.SafeArrayPtr.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_pparray := func(ptr interface{}) { o.SafeArrayPtr = *ptr.(**SafeArray) }
			if err := w.ReadPointer(&o.SafeArrayPtr, _s_pparray, _ptr_pparray); err != nil {
				return err
			}
			return nil
		})
		_s_pparray := func(ptr interface{}) { o.SafeArrayPtr = *ptr.(**SafeArray) }
		if err := w.ReadPointer(&o.SafeArrayPtr, _s_pparray, _ptr_pparray); err != nil {
			return err
		}
		return nil
	})
	_s_pparray := func(ptr interface{}) { o.SafeArrayPtr = *ptr.(**SafeArray) }
	if err := w.ReadPointer(&o.SafeArrayPtr, _s_pparray, _ptr_pparray); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_VariantPtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16396
type Variant_VarUnion_VariantPtr struct {
	VariantPtr *Variant `idl:"name:pvarVal" json:"variant_ptr"`
}

func (*Variant_VarUnion_VariantPtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_VariantPtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.VariantPtr != nil {
		_ptr_pvarVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.VariantPtr != nil {
				_ptr_pvarVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.VariantPtr != nil {
						if err := o.VariantPtr.MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&Variant{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.VariantPtr, _ptr_pvarVal); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.VariantPtr, _ptr_pvarVal); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_VarUnion_VariantPtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pvarVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_pvarVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.VariantPtr == nil {
				o.VariantPtr = &Variant{}
			}
			if err := o.VariantPtr.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvarVal := func(ptr interface{}) { o.VariantPtr = *ptr.(**Variant) }
		if err := w.ReadPointer(&o.VariantPtr, _s_pvarVal, _ptr_pvarVal); err != nil {
			return err
		}
		return nil
	})
	_s_pvarVal := func(ptr interface{}) { o.VariantPtr = *ptr.(**Variant) }
	if err := w.ReadPointer(&o.VariantPtr, _s_pvarVal, _ptr_pvarVal); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_Char structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16
type Variant_VarUnion_Char struct {
	Char uint8 `idl:"name:cVal" json:"char"`
}

func (*Variant_VarUnion_Char) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_Char) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Char); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_Char) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Char); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_Ushort structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 18
type Variant_VarUnion_Ushort struct {
	Ushort uint16 `idl:"name:uiVal" json:"ushort"`
}

func (*Variant_VarUnion_Ushort) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_Ushort) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Ushort); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_Ushort) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Ushort); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_Ulong structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 19
type Variant_VarUnion_Ulong struct {
	Ulong uint32 `idl:"name:ulVal" json:"ulong"`
}

func (*Variant_VarUnion_Ulong) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_Ulong) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Ulong); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_Ulong) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Ulong); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_UlongLong structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 21
type Variant_VarUnion_UlongLong struct {
	UlongLong uint64 `idl:"name:ullVal" json:"ulong_long"`
}

func (*Variant_VarUnion_UlongLong) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_UlongLong) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.UlongLong); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_UlongLong) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.UlongLong); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_Int structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 22
type Variant_VarUnion_Int struct {
	Int int32 `idl:"name:intVal" json:"int"`
}

func (*Variant_VarUnion_Int) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_Int) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Int); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_Int) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Int); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_Uint structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 23
type Variant_VarUnion_Uint struct {
	Uint uint32 `idl:"name:uintVal" json:"uint"`
}

func (*Variant_VarUnion_Uint) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_Uint) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Uint); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_Uint) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Uint); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_Decimal structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 14
type Variant_VarUnion_Decimal struct {
	Decimal *Decimal `idl:"name:decVal" json:"decimal"`
}

func (*Variant_VarUnion_Decimal) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_Decimal) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Decimal != nil {
		if err := o.Decimal.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&Decimal{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_VarUnion_Decimal) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Decimal == nil {
		o.Decimal = &Decimal{}
	}
	if err := o.Decimal.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_CharPtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16400
type Variant_VarUnion_CharPtr struct {
	CharPtr string `idl:"name:pcVal" json:"char_ptr"`
}

func (*Variant_VarUnion_CharPtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_CharPtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.CharPtr != "" {
		_ptr_pcVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteCharString(ctx, w, o.CharPtr); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.CharPtr, _ptr_pcVal); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_VarUnion_CharPtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pcVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadCharString(ctx, w, &o.CharPtr); err != nil {
			return err
		}
		return nil
	})
	_s_pcVal := func(ptr interface{}) { o.CharPtr = *ptr.(*string) }
	if err := w.ReadPointer(&o.CharPtr, _s_pcVal, _ptr_pcVal); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_UshortPtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16402
type Variant_VarUnion_UshortPtr struct {
	UshortPtr uint16 `idl:"name:puiVal" json:"ushort_ptr"`
}

func (*Variant_VarUnion_UshortPtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_UshortPtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.UshortPtr); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_UshortPtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.UshortPtr); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_UlongPtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16403
type Variant_VarUnion_UlongPtr struct {
	UlongPtr uint32 `idl:"name:pulVal" json:"ulong_ptr"`
}

func (*Variant_VarUnion_UlongPtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_UlongPtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.UlongPtr); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_UlongPtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.UlongPtr); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_UlongLongPtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16405
type Variant_VarUnion_UlongLongPtr struct {
	UlongLongPtr uint64 `idl:"name:pullVal" json:"ulong_long_ptr"`
}

func (*Variant_VarUnion_UlongLongPtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_UlongLongPtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.UlongLongPtr); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_UlongLongPtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.UlongLongPtr); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_IntPtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16406
type Variant_VarUnion_IntPtr struct {
	IntPtr int32 `idl:"name:pintVal" json:"int_ptr"`
}

func (*Variant_VarUnion_IntPtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_IntPtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.IntPtr); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_IntPtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.IntPtr); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_UintPtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16407
type Variant_VarUnion_UintPtr struct {
	UintPtr uint32 `idl:"name:puintVal" json:"uint_ptr"`
}

func (*Variant_VarUnion_UintPtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_UintPtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.UintPtr); err != nil {
		return err
	}
	return nil
}
func (o *Variant_VarUnion_UintPtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.UintPtr); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_DecimalPtr structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 16398
type Variant_VarUnion_DecimalPtr struct {
	DecimalPtr *Decimal `idl:"name:pdecVal" json:"decimal_ptr"`
}

func (*Variant_VarUnion_DecimalPtr) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_DecimalPtr) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.DecimalPtr != nil {
		_ptr_pdecVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.DecimalPtr != nil {
				if err := o.DecimalPtr.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Decimal{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DecimalPtr, _ptr_pdecVal); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Variant_VarUnion_DecimalPtr) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_pdecVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.DecimalPtr == nil {
			o.DecimalPtr = &Decimal{}
		}
		if err := o.DecimalPtr.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pdecVal := func(ptr interface{}) { o.DecimalPtr = *ptr.(**Decimal) }
	if err := w.ReadPointer(&o.DecimalPtr, _s_pdecVal, _ptr_pdecVal); err != nil {
		return err
	}
	return nil
}

// Variant_VarUnion_0 structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 0
type Variant_VarUnion_0 struct {
}

func (*Variant_VarUnion_0) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *Variant_VarUnion_0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// Variant_VarUnion_1 structure represents Variant_VarUnion RPC union arm.
//
// It has following labels: 1
type Variant_VarUnion_1 struct {
}

func (*Variant_VarUnion_1) is_Variant_VarUnion() {}

func (o *Variant_VarUnion_1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return nil
}
func (o *Variant_VarUnion_1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	return nil
}

// SafeArrayBound structure represents SAFEARRAYBOUND RPC structure.
//
// The SAFEARRAYBOUND structure specifies the bounds of one dimension of a SAFEARRAY.
type SafeArrayBound struct {
	// cElements:  MUST be set to the number of elements in the current dimension. MUST
	// be nonzero.
	ElementsCount uint32 `idl:"name:cElements" json:"elements_count"`
	// lLbound:  MUST be set to the lower bound of the current dimension.
	//
	// Data of this type MUST be marshaled as specified in [C706] section 14, except that
	// it MUST be marshaled by using a little-endian data representation, regardless of
	// the data representation format label. For more information, see [C706] section 14.2.5.
	LowerBound int32 `idl:"name:lLbound" json:"lower_bound"`
}

func (o *SafeArrayBound) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayBound) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsCount); err != nil {
		return err
	}
	if err := w.WriteData(o.LowerBound); err != nil {
		return err
	}
	return nil
}
func (o *SafeArrayBound) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.LowerBound); err != nil {
		return err
	}
	return nil
}

// SafeArrayString structure represents SAFEARR_BSTR RPC structure.
//
// The SAFEARR_BSTR structure specifies an array of BSTRs (see section 2.2.23).
type SafeArrayString struct {
	// Size:  MUST be set to the total number of elements in the array.
	Size uint32 `idl:"name:Size" json:"size"`
	// aBstr:  MUST be an array of BSTRs (see section 2.2.23).
	//
	// Data of this type MUST be marshaled as specified in [C706] section 14, with the exception
	// that it MUST be marshaled by using a little-endian data representation, regardless
	// of the data representation format label. For more information, see [C706] section
	// 14.2.5.
	String []*String `idl:"name:aBstr;size_is:(Size);pointer:ref" json:"string"`
}

func (o *SafeArrayString) xxx_PreparePayload(ctx context.Context) error {
	if o.String != nil && o.Size == 0 {
		o.Size = uint32(len(o.String))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayString) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if o.String != nil || o.Size > 0 {
		_ptr_aBstr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Size)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.String {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.String[i1] != nil {
					_ptr_aBstr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.String[i1] != nil {
							if err := o.String[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&String{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.String[i1], _ptr_aBstr); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.String); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.String, _ptr_aBstr); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayString) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	_ptr_aBstr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Size > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Size)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.String", sizeInfo[0])
		}
		o.String = make([]*String, sizeInfo[0])
		for i1 := range o.String {
			i1 := i1
			_ptr_aBstr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.String[i1] == nil {
					o.String[i1] = &String{}
				}
				if err := o.String[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_aBstr := func(ptr interface{}) { o.String[i1] = *ptr.(**String) }
			if err := w.ReadPointer(&o.String[i1], _s_aBstr, _ptr_aBstr); err != nil {
				return err
			}
		}
		return nil
	})
	_s_aBstr := func(ptr interface{}) { o.String = *ptr.(*[]*String) }
	if err := w.ReadPointer(&o.String, _s_aBstr, _ptr_aBstr); err != nil {
		return err
	}
	return nil
}

// SafeArrayUnknown structure represents SAFEARR_UNKNOWN RPC structure.
//
// The SAFEARR_UNKNOWN structure specifies an array of MInterfacePointers elements (see
// [MS-DCOM] section 2.2.14) whose IPID is IID_IUnknown (see section 1.9).
type SafeArrayUnknown struct {
	// Size:  MUST be set to the total number of elements in the array.
	Size uint32 `idl:"name:Size" json:"size"`
	// apUnknown:  MUST be an array of MInterfacePointer (see [MS-DCOM], section 2.2.1.10).
	// The iid field in the OBJREF MUST be IID_IUnknown (see section 1.9).
	//
	// The Size field of this type MUST be marshaled as specified in [C706] section 14,
	// with the exception that it MUST be marshaled by using a little-endian data representation,
	// regardless of the data representation format label. For more information, see [C706]
	// section 14.2.5.
	Unknown []*dcom.Unknown `idl:"name:apUnknown;size_is:(Size);pointer:ref" json:"unknown"`
}

func (o *SafeArrayUnknown) xxx_PreparePayload(ctx context.Context) error {
	if o.Unknown != nil && o.Size == 0 {
		o.Size = uint32(len(o.Unknown))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayUnknown) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if o.Unknown != nil || o.Size > 0 {
		_ptr_apUnknown := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Size)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Unknown {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Unknown[i1] != nil {
					_ptr_apUnknown := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Unknown[i1] != nil {
							if err := o.Unknown[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Unknown[i1], _ptr_apUnknown); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Unknown); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Unknown, _ptr_apUnknown); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayUnknown) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	_ptr_apUnknown := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Size > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Size)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Unknown", sizeInfo[0])
		}
		o.Unknown = make([]*dcom.Unknown, sizeInfo[0])
		for i1 := range o.Unknown {
			i1 := i1
			_ptr_apUnknown := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Unknown[i1] == nil {
					o.Unknown[i1] = &dcom.Unknown{}
				}
				if err := o.Unknown[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_apUnknown := func(ptr interface{}) { o.Unknown[i1] = *ptr.(**dcom.Unknown) }
			if err := w.ReadPointer(&o.Unknown[i1], _s_apUnknown, _ptr_apUnknown); err != nil {
				return err
			}
		}
		return nil
	})
	_s_apUnknown := func(ptr interface{}) { o.Unknown = *ptr.(*[]*dcom.Unknown) }
	if err := w.ReadPointer(&o.Unknown, _s_apUnknown, _ptr_apUnknown); err != nil {
		return err
	}
	return nil
}

// SafeArrayDispatch structure represents SAFEARR_DISPATCH RPC structure.
//
// The SAFEARR_DISPATCH structure specifies an array of MInterfacePointer elements (see
// [MS-DCOM] section 2.2.14) whose IPID is IID_IDispatch (see section 1.9).
type SafeArrayDispatch struct {
	// Size:  MUST be set to the total number of elements in the array.
	Size uint32 `idl:"name:Size" json:"size"`
	// apDispatch:  MUST be an array of MInterfacePointer elements (see [MS-DCOM] section
	// 2.2.14). The iid field in the OBJREF MUST be IID_IDispatch (see section 1.9).
	//
	// The Size field of this type MUST be marshaled as specified in [C706] section 14,
	// with the exception that it MUST be marshaled by using a little-endian data representation,
	// regardless of the data representation format label. For more information, see [C706]
	// section 14.2.5.
	Dispatch []*Dispatch `idl:"name:apDispatch;size_is:(Size);pointer:ref" json:"dispatch"`
}

func (o *SafeArrayDispatch) xxx_PreparePayload(ctx context.Context) error {
	if o.Dispatch != nil && o.Size == 0 {
		o.Size = uint32(len(o.Dispatch))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayDispatch) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if o.Dispatch != nil || o.Size > 0 {
		_ptr_apDispatch := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Size)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Dispatch {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Dispatch[i1] != nil {
					_ptr_apDispatch := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Dispatch[i1] != nil {
							if err := o.Dispatch[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&Dispatch{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Dispatch[i1], _ptr_apDispatch); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Dispatch); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Dispatch, _ptr_apDispatch); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayDispatch) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	_ptr_apDispatch := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Size > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Size)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Dispatch", sizeInfo[0])
		}
		o.Dispatch = make([]*Dispatch, sizeInfo[0])
		for i1 := range o.Dispatch {
			i1 := i1
			_ptr_apDispatch := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Dispatch[i1] == nil {
					o.Dispatch[i1] = &Dispatch{}
				}
				if err := o.Dispatch[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_apDispatch := func(ptr interface{}) { o.Dispatch[i1] = *ptr.(**Dispatch) }
			if err := w.ReadPointer(&o.Dispatch[i1], _s_apDispatch, _ptr_apDispatch); err != nil {
				return err
			}
		}
		return nil
	})
	_s_apDispatch := func(ptr interface{}) { o.Dispatch = *ptr.(*[]*Dispatch) }
	if err := w.ReadPointer(&o.Dispatch, _s_apDispatch, _ptr_apDispatch); err != nil {
		return err
	}
	return nil
}

// SafeArrayVariant structure represents SAFEARR_VARIANT RPC structure.
//
// The SAFEARR_VARIANT structure specifies an array of VARIANT types.
type SafeArrayVariant struct {
	// Size:  MUST be set to the total number of elements in the array. MUST be nonzero.
	Size uint32 `idl:"name:Size" json:"size"`
	// aVariant:  MUST be an array of VARIANT types. For more information, see section 2.2.29.2.
	//
	// The Size field of this type MUST be marshaled as specified in [C706] section 14,
	// with the exception that it MUST be marshaled by using a little-endian data representation,
	// regardless of the data representation format label. For more information, see [C706]
	// section 14.2.5.
	Variant []*Variant `idl:"name:aVariant;size_is:(Size);pointer:ref" json:"variant"`
}

func (o *SafeArrayVariant) xxx_PreparePayload(ctx context.Context) error {
	if o.Variant != nil && o.Size == 0 {
		o.Size = uint32(len(o.Variant))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayVariant) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if o.Variant != nil || o.Size > 0 {
		_ptr_aVariant := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Size)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Variant {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Variant[i1] != nil {
					_ptr_aVariant := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Variant[i1] != nil {
							if err := o.Variant[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&Variant{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Variant[i1], _ptr_aVariant); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Variant); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Variant, _ptr_aVariant); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayVariant) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	_ptr_aVariant := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Size > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Size)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Variant", sizeInfo[0])
		}
		o.Variant = make([]*Variant, sizeInfo[0])
		for i1 := range o.Variant {
			i1 := i1
			_ptr_aVariant := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Variant[i1] == nil {
					o.Variant[i1] = &Variant{}
				}
				if err := o.Variant[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_aVariant := func(ptr interface{}) { o.Variant[i1] = *ptr.(**Variant) }
			if err := w.ReadPointer(&o.Variant[i1], _s_aVariant, _ptr_aVariant); err != nil {
				return err
			}
		}
		return nil
	})
	_s_aVariant := func(ptr interface{}) { o.Variant = *ptr.(*[]*Variant) }
	if err := w.ReadPointer(&o.Variant, _s_aVariant, _ptr_aVariant); err != nil {
		return err
	}
	return nil
}

// SafeArrayRecord structure represents SAFEARR_BRECORD RPC structure.
//
// The SAFEARR_BRECORD structure specifies an array of UDTs.
type SafeArrayRecord struct {
	// Size:  The number of BRECORD elements in the aRecord array. This MUST be set to 1.
	Size uint32 `idl:"name:Size" json:"size"`
	// aRecord:  MUST be the collection of UDTs as specified in BRECORD (see section 2.2.28.2).
	//
	// The Size field of this type MUST be marshaled as specified in [C706] section 14,
	// with the exception that it MUST be marshaled by using a little-endian data representation,
	// regardless of the data representation format label. For more information, see [C706]
	// section 14.2.5.
	Record []*Record `idl:"name:aRecord;size_is:(Size);pointer:ref" json:"record"`
}

func (o *SafeArrayRecord) xxx_PreparePayload(ctx context.Context) error {
	if o.Record != nil && o.Size == 0 {
		o.Size = uint32(len(o.Record))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayRecord) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if o.Record != nil || o.Size > 0 {
		_ptr_aRecord := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Size)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Record {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Record[i1] != nil {
					_ptr_aRecord := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Record[i1] != nil {
							if err := o.Record[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&Record{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Record[i1], _ptr_aRecord); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Record); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Record, _ptr_aRecord); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayRecord) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	_ptr_aRecord := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Size > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Size)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Record", sizeInfo[0])
		}
		o.Record = make([]*Record, sizeInfo[0])
		for i1 := range o.Record {
			i1 := i1
			_ptr_aRecord := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Record[i1] == nil {
					o.Record[i1] = &Record{}
				}
				if err := o.Record[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_aRecord := func(ptr interface{}) { o.Record[i1] = *ptr.(**Record) }
			if err := w.ReadPointer(&o.Record[i1], _s_aRecord, _ptr_aRecord); err != nil {
				return err
			}
		}
		return nil
	})
	_s_aRecord := func(ptr interface{}) { o.Record = *ptr.(*[]*Record) }
	if err := w.ReadPointer(&o.Record, _s_aRecord, _ptr_aRecord); err != nil {
		return err
	}
	return nil
}

// SafeArrayHaveIID structure represents SAFEARR_HAVEIID RPC structure.
//
// The SAFEARR_HAVEIID structure defines an array of MInterfacePointers (see [MS-DCOM]
// section 2.2.14
type SafeArrayHaveIID struct {
	// Size:  MUST be set to the total number of elements in the array. This MUST be nonzero.
	Size uint32 `idl:"name:Size" json:"size"`
	// apUnknown:  MUST be an array of MInterfacePointer elements. The OBJREF iid field
	// MUST be the same as the value of the iid.
	Unknown []*dcom.Unknown `idl:"name:apUnknown;size_is:(Size);pointer:ref" json:"unknown"`
	// iid:  MUST specify the IID of each of the elements in the SAFEARRAY.
	//
	// The Size and iid fields of this type MUST be marshaled as specified in [C706] section
	// 14, with the exception that it MUST be marshaled by using a little-endian data representation,
	// regardless of the data representation format label. For more information, see [C706]
	// section 14.2.5.
	IID *dcom.IID `idl:"name:iid" json:"iid"`
}

func (o *SafeArrayHaveIID) xxx_PreparePayload(ctx context.Context) error {
	if o.Unknown != nil && o.Size == 0 {
		o.Size = uint32(len(o.Unknown))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayHaveIID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if o.Unknown != nil || o.Size > 0 {
		_ptr_apUnknown := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Size)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Unknown {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Unknown[i1] != nil {
					_ptr_apUnknown := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Unknown[i1] != nil {
							if err := o.Unknown[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Unknown[i1], _ptr_apUnknown); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Unknown); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Unknown, _ptr_apUnknown); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.IID != nil {
		if err := o.IID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dcom.IID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *SafeArrayHaveIID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	_ptr_apUnknown := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Size > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Size)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Unknown", sizeInfo[0])
		}
		o.Unknown = make([]*dcom.Unknown, sizeInfo[0])
		for i1 := range o.Unknown {
			i1 := i1
			_ptr_apUnknown := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Unknown[i1] == nil {
					o.Unknown[i1] = &dcom.Unknown{}
				}
				if err := o.Unknown[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_apUnknown := func(ptr interface{}) { o.Unknown[i1] = *ptr.(**dcom.Unknown) }
			if err := w.ReadPointer(&o.Unknown[i1], _s_apUnknown, _ptr_apUnknown); err != nil {
				return err
			}
		}
		return nil
	})
	_s_apUnknown := func(ptr interface{}) { o.Unknown = *ptr.(*[]*dcom.Unknown) }
	if err := w.ReadPointer(&o.Unknown, _s_apUnknown, _ptr_apUnknown); err != nil {
		return err
	}
	if o.IID == nil {
		o.IID = &dcom.IID{}
	}
	if err := o.IID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ByteSizedArray structure represents BYTE_SIZEDARR RPC structure.
//
// The BYTE_SIZEDARR structure specifies a BYTE array.
type ByteSizedArray struct {
	// clSize:  MUST be set to the total number of elements in the array. This MUST be nonzero.
	Size uint32 `idl:"name:clSize" json:"size"`
	// pData:  MUST be an array of BYTEs.
	//
	// Data of this type MUST be marshaled as specified in [C706] section 14, with the exception
	// that it MUST be marshaled by using a little-endian data representation, regardless
	// of the data representation format label. For more information, see [C706] section
	// 14.2.5.
	Data []byte `idl:"name:pData;size_is:(clSize)" json:"data"`
}

func (o *ByteSizedArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.Size == 0 {
		o.Size = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ByteSizedArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if o.Data != nil || o.Size > 0 {
		_ptr_pData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Size)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
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
		})
		if err := w.WritePointer(&o.Data, _ptr_pData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ByteSizedArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	_ptr_pData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Size > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Size)
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
	})
	_s_pData := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Data, _s_pData, _ptr_pData); err != nil {
		return err
	}
	return nil
}

// WordSizedArray structure represents WORD_SIZEDARR RPC structure.
//
// The WORD_SIZEDARR structure defines an array of unsigned 16-bit integers.
type WordSizedArray struct {
	// clSize:  MUST be set to the total number of elements in the array. This MUST be nonzero.
	Size uint32 `idl:"name:clSize" json:"size"`
	// pData:  MUST be an array of WORD elements.
	//
	// Data of this type MUST be marshaled as specified in [C706] section 14, with the exception
	// that it MUST be marshaled by using a little-endian data representation, regardless
	// of the data representation format label. For more information, see [C706] section
	// 14.2.5.
	Data []uint16 `idl:"name:pData;size_is:(clSize)" json:"data"`
}

func (o *WordSizedArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.Size == 0 {
		o.Size = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *WordSizedArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if o.Data != nil || o.Size > 0 {
		_ptr_pData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Size)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
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
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Data, _ptr_pData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *WordSizedArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	_ptr_pData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Size > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Size)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]uint16, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pData := func(ptr interface{}) { o.Data = *ptr.(*[]uint16) }
	if err := w.ReadPointer(&o.Data, _s_pData, _ptr_pData); err != nil {
		return err
	}
	return nil
}

// DwordSizedArray structure represents DWORD_SIZEDARR RPC structure.
//
// The DWORD_SIZEDARR structure defines an array of unsigned 32-bit integers.
type DwordSizedArray struct {
	// clSize:  MUST be set to the number of elements within the array. This MUST be nonzero.
	Size uint32 `idl:"name:clSize" json:"size"`
	// pData:  MUST be an array of DWORD elements.
	//
	// Data of this type MUST be marshaled as specified in [C706] section 14, with the exception
	// that it MUST be marshaled by using a little-endian data representation, regardless
	// of the data representation format label. For more information, see [C706] section
	// 14.2.5.
	Data []uint32 `idl:"name:pData;size_is:(clSize)" json:"data"`
}

func (o *DwordSizedArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.Size == 0 {
		o.Size = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DwordSizedArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if o.Data != nil || o.Size > 0 {
		_ptr_pData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Size)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
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
				if err := w.WriteData(uint32(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Data, _ptr_pData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *DwordSizedArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	_ptr_pData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Size > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Size)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]uint32, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pData := func(ptr interface{}) { o.Data = *ptr.(*[]uint32) }
	if err := w.ReadPointer(&o.Data, _s_pData, _ptr_pData); err != nil {
		return err
	}
	return nil
}

// HyperSizedArray structure represents HYPER_SIZEDARR RPC structure.
//
// The HYPER_SIZEDARR structure defines an array of 64-bit integers.
type HyperSizedArray struct {
	// clSize:  MUST be set to the total number of elements in the array. This MUST be nonzero.
	Size uint32 `idl:"name:clSize" json:"size"`
	// pData:  MUST be an array of hyper elements.
	//
	// Data of this type MUST be marshaled as specified in [C706] section 14, with the exception
	// that it MUST be marshaled by using a little-endian data representation, regardless
	// of the data representation format label. For more information, see [C706] section
	// 14.2.5.
	Data []int64 `idl:"name:pData;size_is:(clSize)" json:"data"`
}

func (o *HyperSizedArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Data != nil && o.Size == 0 {
		o.Size = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *HyperSizedArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if o.Data != nil || o.Size > 0 {
		_ptr_pData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Size)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
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
				if err := w.WriteData(int64(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Data, _ptr_pData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *HyperSizedArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	_ptr_pData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Size > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Size)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]int64, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_pData := func(ptr interface{}) { o.Data = *ptr.(*[]int64) }
	if err := w.ReadPointer(&o.Data, _s_pData, _ptr_pData); err != nil {
		return err
	}
	return nil
}

// SafeArrayUnion structure represents SAFEARRAYUNION RPC union.
//
// The SAFEARRAYUNION union defines the mapping between the discriminant value and the
// contained array data.
//
// _wireSAFEARRAY_UNION: MUST contain an instance of the type, according to the value
// of the union discriminant.
//
// The sfType field MUST be marshaled as specified in [C706] section 14, with the exception
// that it MUST be marshaled by using a little-endian data representation, regardless
// of the data representation format label. For more information, see [C706] section
// 14.2.5.
type SafeArrayUnion struct {
	SafeArrayType uint32
	// Types that are assignable to Value
	//
	// *SafeArrayUnion_String
	// *SafeArrayUnion_Unknown
	// *SafeArrayUnion_Dispatch
	// *SafeArrayUnion_Variant
	// *SafeArrayUnion_Record
	// *SafeArrayUnion_HaveIID
	// *SafeArrayUnion_Byte
	// *SafeArrayUnion_Word
	// *SafeArrayUnion_Long
	// *SafeArrayUnion_Hyper
	Value is_SafeArrayUnion `json:"value"`
}

func (o *SafeArrayUnion) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *SafeArrayUnion_String:
		if value != nil {
			return value.String
		}
	case *SafeArrayUnion_Unknown:
		if value != nil {
			return value.Unknown
		}
	case *SafeArrayUnion_Dispatch:
		if value != nil {
			return value.Dispatch
		}
	case *SafeArrayUnion_Variant:
		if value != nil {
			return value.Variant
		}
	case *SafeArrayUnion_Record:
		if value != nil {
			return value.Record
		}
	case *SafeArrayUnion_HaveIID:
		if value != nil {
			return value.HaveIID
		}
	case *SafeArrayUnion_Byte:
		if value != nil {
			return value.Byte
		}
	case *SafeArrayUnion_Word:
		if value != nil {
			return value.Word
		}
	case *SafeArrayUnion_Long:
		if value != nil {
			return value.Long
		}
	case *SafeArrayUnion_Hyper:
		if value != nil {
			return value.Hyper
		}
	}
	return nil
}

type is_SafeArrayUnion interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_SafeArrayUnion()
}

func (o *SafeArrayUnion) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(uint32(o.SafeArrayType)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch o.SafeArrayType {
	case uint32(8):
		_o, _ := o.Value.(*SafeArrayUnion_String)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SafeArrayUnion_String{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(13):
		_o, _ := o.Value.(*SafeArrayUnion_Unknown)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SafeArrayUnion_Unknown{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(9):
		_o, _ := o.Value.(*SafeArrayUnion_Dispatch)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SafeArrayUnion_Dispatch{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(12):
		_o, _ := o.Value.(*SafeArrayUnion_Variant)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SafeArrayUnion_Variant{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(36):
		_o, _ := o.Value.(*SafeArrayUnion_Record)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SafeArrayUnion_Record{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(32781):
		_o, _ := o.Value.(*SafeArrayUnion_HaveIID)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SafeArrayUnion_HaveIID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(16):
		_o, _ := o.Value.(*SafeArrayUnion_Byte)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SafeArrayUnion_Byte{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*SafeArrayUnion_Word)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SafeArrayUnion_Word{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(3):
		_o, _ := o.Value.(*SafeArrayUnion_Long)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SafeArrayUnion_Long{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(20):
		_o, _ := o.Value.(*SafeArrayUnion_Hyper)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SafeArrayUnion_Hyper{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.SafeArrayType)
	}
	return nil
}

func (o *SafeArrayUnion) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadData((*uint32)(&o.SafeArrayType)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch o.SafeArrayType {
	case uint32(8):
		o.Value = &SafeArrayUnion_String{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(13):
		o.Value = &SafeArrayUnion_Unknown{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(9):
		o.Value = &SafeArrayUnion_Dispatch{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(12):
		o.Value = &SafeArrayUnion_Variant{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(36):
		o.Value = &SafeArrayUnion_Record{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(32781):
		o.Value = &SafeArrayUnion_HaveIID{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(16):
		o.Value = &SafeArrayUnion_Byte{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &SafeArrayUnion_Word{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(3):
		o.Value = &SafeArrayUnion_Long{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(20):
		o.Value = &SafeArrayUnion_Hyper{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", o.SafeArrayType)
	}
	return nil
}

// SafeArrayUnion_String structure represents SAFEARRAYUNION RPC union arm.
//
// It has following labels: 8
type SafeArrayUnion_String struct {
	String *SafeArrayString `idl:"name:BstrStr" json:"string"`
}

func (*SafeArrayUnion_String) is_SafeArrayUnion() {}

func (o *SafeArrayUnion_String) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.String != nil {
		if err := o.String.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SafeArrayString{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayUnion_String) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.String == nil {
		o.String = &SafeArrayString{}
	}
	if err := o.String.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SafeArrayUnion_Unknown structure represents SAFEARRAYUNION RPC union arm.
//
// It has following labels: 13
type SafeArrayUnion_Unknown struct {
	Unknown *SafeArrayUnknown `idl:"name:UnknownStr" json:"unknown"`
}

func (*SafeArrayUnion_Unknown) is_SafeArrayUnion() {}

func (o *SafeArrayUnion_Unknown) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Unknown != nil {
		if err := o.Unknown.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SafeArrayUnknown{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayUnion_Unknown) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Unknown == nil {
		o.Unknown = &SafeArrayUnknown{}
	}
	if err := o.Unknown.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SafeArrayUnion_Dispatch structure represents SAFEARRAYUNION RPC union arm.
//
// It has following labels: 9
type SafeArrayUnion_Dispatch struct {
	Dispatch *SafeArrayDispatch `idl:"name:DispatchStr" json:"dispatch"`
}

func (*SafeArrayUnion_Dispatch) is_SafeArrayUnion() {}

func (o *SafeArrayUnion_Dispatch) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Dispatch != nil {
		if err := o.Dispatch.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SafeArrayDispatch{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayUnion_Dispatch) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Dispatch == nil {
		o.Dispatch = &SafeArrayDispatch{}
	}
	if err := o.Dispatch.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SafeArrayUnion_Variant structure represents SAFEARRAYUNION RPC union arm.
//
// It has following labels: 12
type SafeArrayUnion_Variant struct {
	Variant *SafeArrayVariant `idl:"name:VariantStr" json:"variant"`
}

func (*SafeArrayUnion_Variant) is_SafeArrayUnion() {}

func (o *SafeArrayUnion_Variant) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Variant != nil {
		if err := o.Variant.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SafeArrayVariant{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayUnion_Variant) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Variant == nil {
		o.Variant = &SafeArrayVariant{}
	}
	if err := o.Variant.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SafeArrayUnion_Record structure represents SAFEARRAYUNION RPC union arm.
//
// It has following labels: 36
type SafeArrayUnion_Record struct {
	Record *SafeArrayRecord `idl:"name:RecordStr" json:"record"`
}

func (*SafeArrayUnion_Record) is_SafeArrayUnion() {}

func (o *SafeArrayUnion_Record) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Record != nil {
		if err := o.Record.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SafeArrayRecord{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayUnion_Record) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Record == nil {
		o.Record = &SafeArrayRecord{}
	}
	if err := o.Record.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SafeArrayUnion_HaveIID structure represents SAFEARRAYUNION RPC union arm.
//
// It has following labels: 32781
type SafeArrayUnion_HaveIID struct {
	HaveIID *SafeArrayHaveIID `idl:"name:HaveIidStr" json:"have_iid"`
}

func (*SafeArrayUnion_HaveIID) is_SafeArrayUnion() {}

func (o *SafeArrayUnion_HaveIID) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.HaveIID != nil {
		if err := o.HaveIID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SafeArrayHaveIID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayUnion_HaveIID) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.HaveIID == nil {
		o.HaveIID = &SafeArrayHaveIID{}
	}
	if err := o.HaveIID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SafeArrayUnion_Byte structure represents SAFEARRAYUNION RPC union arm.
//
// It has following labels: 16
type SafeArrayUnion_Byte struct {
	Byte *ByteSizedArray `idl:"name:ByteStr" json:"byte"`
}

func (*SafeArrayUnion_Byte) is_SafeArrayUnion() {}

func (o *SafeArrayUnion_Byte) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Byte != nil {
		if err := o.Byte.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ByteSizedArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayUnion_Byte) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Byte == nil {
		o.Byte = &ByteSizedArray{}
	}
	if err := o.Byte.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SafeArrayUnion_Word structure represents SAFEARRAYUNION RPC union arm.
//
// It has following labels: 2
type SafeArrayUnion_Word struct {
	Word *WordSizedArray `idl:"name:WordStr" json:"word"`
}

func (*SafeArrayUnion_Word) is_SafeArrayUnion() {}

func (o *SafeArrayUnion_Word) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Word != nil {
		if err := o.Word.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&WordSizedArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayUnion_Word) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Word == nil {
		o.Word = &WordSizedArray{}
	}
	if err := o.Word.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SafeArrayUnion_Long structure represents SAFEARRAYUNION RPC union arm.
//
// It has following labels: 3
type SafeArrayUnion_Long struct {
	Long *DwordSizedArray `idl:"name:LongStr" json:"long"`
}

func (*SafeArrayUnion_Long) is_SafeArrayUnion() {}

func (o *SafeArrayUnion_Long) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Long != nil {
		if err := o.Long.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&DwordSizedArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayUnion_Long) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Long == nil {
		o.Long = &DwordSizedArray{}
	}
	if err := o.Long.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SafeArrayUnion_Hyper structure represents SAFEARRAYUNION RPC union arm.
//
// It has following labels: 20
type SafeArrayUnion_Hyper struct {
	Hyper *HyperSizedArray `idl:"name:HyperStr" json:"hyper"`
}

func (*SafeArrayUnion_Hyper) is_SafeArrayUnion() {}

func (o *SafeArrayUnion_Hyper) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Hyper != nil {
		if err := o.Hyper.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&HyperSizedArray{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArrayUnion_Hyper) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if o.Hyper == nil {
		o.Hyper = &HyperSizedArray{}
	}
	if err := o.Hyper.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SafeArray structure represents _SAFEARRAY RPC structure.
type SafeArray struct {
	DimsCount      uint16            `idl:"name:cDims" json:"dims_count"`
	Features       uint16            `idl:"name:fFeatures" json:"features"`
	ElementsLength uint32            `idl:"name:cbElements" json:"elements_length"`
	LocksCount     uint32            `idl:"name:cLocks" json:"locks_count"`
	ArrayStructs   *SafeArrayUnion   `idl:"name:uArrayStructs" json:"array_structs"`
	Bound          []*SafeArrayBound `idl:"name:rgsabound;size_is:(cDims)" json:"bound"`
}

func (o *SafeArray) xxx_PreparePayload(ctx context.Context) error {
	if o.Bound != nil && o.DimsCount == 0 {
		o.DimsCount = uint16(len(o.Bound))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *SafeArray) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DimsCount)
	return []uint64{
		dimSize1,
	}
}
func (o *SafeArray) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.DimsCount); err != nil {
		return err
	}
	if err := w.WriteData(o.Features); err != nil {
		return err
	}
	if err := w.WriteData(o.ElementsLength); err != nil {
		return err
	}
	if err := w.WriteData(o.LocksCount); err != nil {
		return err
	}
	if o.ArrayStructs != nil {
		if err := o.ArrayStructs.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SafeArrayUnion{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	for i1 := range o.Bound {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.Bound[i1] != nil {
			if err := o.Bound[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SafeArrayBound{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Bound); uint64(i1) < sizeInfo[0]; i1++ {
		if err := (&SafeArrayBound{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SafeArray) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.DimsCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.Features); err != nil {
		return err
	}
	if err := w.ReadData(&o.ElementsLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocksCount); err != nil {
		return err
	}
	if o.ArrayStructs == nil {
		o.ArrayStructs = &SafeArrayUnion{}
	}
	if err := o.ArrayStructs.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DimsCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DimsCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Bound", sizeInfo[0])
	}
	o.Bound = make([]*SafeArrayBound, sizeInfo[0])
	for i1 := range o.Bound {
		i1 := i1
		if o.Bound[i1] == nil {
			o.Bound[i1] = &SafeArrayBound{}
		}
		if err := o.Bound[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// RecordInfo structure represents RecordInfo RPC structure.
//
// The RecordInfoData structure specifies information that is needed to identify the
// definition of a UDT, such as the GUID and version of the automation scope that defines
// the UDT, and the GUID of the type itself.
type RecordInfo struct {
	// libraryGuid:   MUST be set to a GUID that identifies the IDL automation scope of
	// the UDT (see section 2.2.49.2).
	LibraryGUID *dtyp.GUID `idl:"name:libraryGuid" json:"library_guid"`
	// verMajor:   MUST be set to the major version of the UDT automation scope (see section
	// 2.2.49.2).
	VerMajor uint32 `idl:"name:verMajor" json:"ver_major"`
	// recGuid:  MUST be set to the GUID of the UDT.
	RecordGUID *dtyp.GUID `idl:"name:recGuid" json:"record_guid"`
	// verMinor:   MUST be set to the minor version of the UDT's automation scope (see section
	// 2.2.49.2).
	VerMinor uint32 `idl:"name:verMinor" json:"ver_minor"`
	// Lcid:   MUST be set to the locale ID of the UDT's automation scope (see section 2.2.49.2).
	//
	// RecordInfoData structures allow a client and a server to fully specify the identity
	// of the UDT type being marshaled in the containing BRECORD (section 2.2.28.2). The
	// client and the server MUST be able to reference the same type definition, by sharing
	// through an unspecified mechanism a consistent view of the IDL automation scope of
	// the UDT.<13>
	//
	// Data of this type MUST be marshaled as specified in [C706] section 14, with the exception
	// that it MUST be marshaled by using a little-endian data representation, regardless
	// of the data representation format label. For more information, see [C706] section
	// 14.2.5.
	LocaleID uint32 `idl:"name:Lcid" json:"locale_id"`
}

func (o *RecordInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RecordInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.LibraryGUID != nil {
		if err := o.LibraryGUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.VerMajor); err != nil {
		return err
	}
	if o.RecordGUID != nil {
		if err := o.RecordGUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.VerMinor); err != nil {
		return err
	}
	if err := w.WriteData(o.LocaleID); err != nil {
		return err
	}
	return nil
}
func (o *RecordInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.LibraryGUID == nil {
		o.LibraryGUID = &dtyp.GUID{}
	}
	if err := o.LibraryGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.VerMajor); err != nil {
		return err
	}
	if o.RecordGUID == nil {
		o.RecordGUID = &dtyp.GUID{}
	}
	if err := o.RecordGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.VerMinor); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocaleID); err != nil {
		return err
	}
	return nil
}

// DispatchParams structure represents DISPPARAMS RPC structure.
//
// The DISPPARAMS structure is used by the Invoke method (see section 3.1.4.4) to contain
// the arguments passed to a method or property.
type DispatchParams struct {
	// rgvarg:  MUST be the array of arguments passed to the method or property call.
	Args []*Variant `idl:"name:rgvarg;size_is:(cArgs)" json:"args"`
	// rgdispidNamedArgs:  MUST be the array of DISPIDs corresponding to the named arguments
	// (see section 3.1.4.4).
	NamedArgs []int32 `idl:"name:rgdispidNamedArgs;size_is:(cNamedArgs)" json:"named_args"`
	// cArgs:  MUST equal the number of arguments passed to the method.
	ArgsCount uint32 `idl:"name:cArgs" json:"args_count"`
	// cNamedArgs:  MUST equal the number of named arguments passed to the method. This
	// value MUST be less than or equal to the value of cArgs.
	//
	// The arguments passed in DISPPARAMS MUST be stored as specified in section 3.1.4.4.2.
	NamedArgsCount uint32 `idl:"name:cNamedArgs" json:"named_args_count"`
}

func (o *DispatchParams) xxx_PreparePayload(ctx context.Context) error {
	if o.Args != nil && o.ArgsCount == 0 {
		o.ArgsCount = uint32(len(o.Args))
	}
	if o.NamedArgs != nil && o.NamedArgsCount == 0 {
		o.NamedArgsCount = uint32(len(o.NamedArgs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *DispatchParams) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Args != nil || o.ArgsCount > 0 {
		_ptr_rgvarg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ArgsCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Args {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Args[i1] != nil {
					_ptr_rgvarg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.Args[i1] != nil {
							if err := o.Args[i1].MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&Variant{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Args[i1], _ptr_rgvarg); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Args); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Args, _ptr_rgvarg); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.NamedArgs != nil || o.NamedArgsCount > 0 {
		_ptr_rgdispidNamedArgs := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.NamedArgsCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.NamedArgs {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.NamedArgs[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.NamedArgs); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(int32(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.NamedArgs, _ptr_rgdispidNamedArgs); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ArgsCount); err != nil {
		return err
	}
	if err := w.WriteData(o.NamedArgsCount); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *DispatchParams) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_rgvarg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ArgsCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ArgsCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Args", sizeInfo[0])
		}
		o.Args = make([]*Variant, sizeInfo[0])
		for i1 := range o.Args {
			i1 := i1
			_ptr_rgvarg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Args[i1] == nil {
					o.Args[i1] = &Variant{}
				}
				if err := o.Args[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_rgvarg := func(ptr interface{}) { o.Args[i1] = *ptr.(**Variant) }
			if err := w.ReadPointer(&o.Args[i1], _s_rgvarg, _ptr_rgvarg); err != nil {
				return err
			}
		}
		return nil
	})
	_s_rgvarg := func(ptr interface{}) { o.Args = *ptr.(*[]*Variant) }
	if err := w.ReadPointer(&o.Args, _s_rgvarg, _ptr_rgvarg); err != nil {
		return err
	}
	_ptr_rgdispidNamedArgs := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.NamedArgsCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.NamedArgsCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.NamedArgs", sizeInfo[0])
		}
		o.NamedArgs = make([]int32, sizeInfo[0])
		for i1 := range o.NamedArgs {
			i1 := i1
			if err := w.ReadData(&o.NamedArgs[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_rgdispidNamedArgs := func(ptr interface{}) { o.NamedArgs = *ptr.(*[]int32) }
	if err := w.ReadPointer(&o.NamedArgs, _s_rgdispidNamedArgs, _ptr_rgdispidNamedArgs); err != nil {
		return err
	}
	if err := w.ReadData(&o.ArgsCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.NamedArgsCount); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// ExceptionInfo structure represents EXCEPINFO RPC structure.
//
// The EXCEPINFO structure is filled in by an automation server to describe an exception
// that occurred during a call to Invoke (as specified in section 3.1.4.4). If no exception
// occurred, the server MUST set both wCode and scode to 0.
type ExceptionInfo struct {
	// wCode:   An implementation-specific<14> value that identifies an error.
	//
	//	+--------------+----------------------------------------------------------------------------------+
	//	|              |                                                                                  |
	//	|    VALUE     |                                     MEANING                                      |
	//	|              |                                                                                  |
	//	+--------------+----------------------------------------------------------------------------------+
	//	+--------------+----------------------------------------------------------------------------------+
	//	|            0 | The value MUST be zero for either of the following conditions: This field does   |
	//	|              | not contain an error code. The value in the scode field is nonzero.              |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 1000 < value | Implementation-specific error values MUST be greater than 1000.                  |
	//	+--------------+----------------------------------------------------------------------------------+
	Code uint16 `idl:"name:wCode" json:"code"`
	// wReserved:  MUST be set to 0, and MUST be ignored on receipt.
	_ uint16 `idl:"name:wReserved"`
	// bstrSource:  MUST<15> be set to an implementation-specific string that identifies
	// the source of the exception.
	Source *String `idl:"name:bstrSource" json:"source"`
	// bstrDescription:  MUST<16> be set to an implementation-specific string, or to a NULL
	// BSTR if no description is available.
	Description *String `idl:"name:bstrDescription" json:"description"`
	// bstrHelpFile:  MUST<17> be set to an implementation-specific string, or to a NULL
	// BSTR if no help is available.
	HelpFile *String `idl:"name:bstrHelpFile" json:"help_file"`
	// dwHelpContext:  MUST<18> be set to an implementation-specific integer. If bstrHelpFile
	// is NULL, this field MUST be set to 0, and MUST be ignored on receipt.
	HelpContext uint32 `idl:"name:dwHelpContext" json:"help_context"`
	// pvReserved:  MUST be set to NULL, and MUST be ignored on receipt.
	_ uint64 `idl:"name:pvReserved"`
	// pfnDeferredFillIn:  MAY be set to NULL, and MUST be ignored on receipt.<19>
	PfnDeferredFillIn uint64 `idl:"name:pfnDeferredFillIn" json:"pfn_deferred_fill_in"`
	// scode:  MUST be set to a failure HRESULT that describes the error, or to 0 to indicate
	// that it does not contain an error code. If wCode is nonzero, this field MUST be set
	// to 0.
	HResult int32 `idl:"name:scode" json:"h_result"`
}

func (o *ExceptionInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ExceptionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Code); err != nil {
		return err
	}
	// reserved wReserved
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	if o.Source != nil {
		_ptr_bstrSource := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Source != nil {
				if err := o.Source.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&String{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Source, _ptr_bstrSource); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Description != nil {
		_ptr_bstrDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Description != nil {
				if err := o.Description.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&String{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Description, _ptr_bstrDescription); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.HelpFile != nil {
		_ptr_bstrHelpFile := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.HelpFile != nil {
				if err := o.HelpFile.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&String{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.HelpFile, _ptr_bstrHelpFile); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.HelpContext); err != nil {
		return err
	}
	// reserved pvReserved
	if err := w.WriteData(uint64(0)); err != nil {
		return err
	}
	if err := w.WriteData(ndr.Uint3264(o.PfnDeferredFillIn)); err != nil {
		return err
	}
	if err := w.WriteData(o.HResult); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *ExceptionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Code); err != nil {
		return err
	}
	// reserved wReserved
	var _wReserved uint16
	if err := w.ReadData(&_wReserved); err != nil {
		return err
	}
	_ptr_bstrSource := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Source == nil {
			o.Source = &String{}
		}
		if err := o.Source.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_bstrSource := func(ptr interface{}) { o.Source = *ptr.(**String) }
	if err := w.ReadPointer(&o.Source, _s_bstrSource, _ptr_bstrSource); err != nil {
		return err
	}
	_ptr_bstrDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Description == nil {
			o.Description = &String{}
		}
		if err := o.Description.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_bstrDescription := func(ptr interface{}) { o.Description = *ptr.(**String) }
	if err := w.ReadPointer(&o.Description, _s_bstrDescription, _ptr_bstrDescription); err != nil {
		return err
	}
	_ptr_bstrHelpFile := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.HelpFile == nil {
			o.HelpFile = &String{}
		}
		if err := o.HelpFile.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_bstrHelpFile := func(ptr interface{}) { o.HelpFile = *ptr.(**String) }
	if err := w.ReadPointer(&o.HelpFile, _s_bstrHelpFile, _ptr_bstrHelpFile); err != nil {
		return err
	}
	if err := w.ReadData(&o.HelpContext); err != nil {
		return err
	}
	// reserved pvReserved
	var _pvReserved uint64
	if err := w.ReadData((*ndr.Uint3264)(&_pvReserved)); err != nil {
		return err
	}
	if err := w.ReadData((*ndr.Uint3264)(&o.PfnDeferredFillIn)); err != nil {
		return err
	}
	if err := w.ReadData(&o.HResult); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// TypeDesc structure represents TYPEDESC RPC structure.
//
// The TYPEDESC structure is used in the ARRAYDESC, ELEMDESC, and TYPEATTR structures
// to identify and describe the type of a data member, the return type of a method,
// or the type of a method parameter.
type TypeDesc struct {
	// _tdUnion:  MUST contain an instance of the type, according to the VARENUM value provided
	// in the vt field.
	Union *TypeDesc_Union `idl:"name:_tdUnion;switch_is:vt" json:"union"`
	// vt:  MUST be set to one of the values that are specified as available to a TYPEDESC
	// and identified with a "T" in the Context column of the table in 2.2.7. MUST be set
	// to VT_PTR if the ELEMDESC is contained in a VARDESC that describes an appobject coclass,
	// as specified in section 2.2.49.8.
	VT uint16 `idl:"name:vt" json:"vt"`
}

func (o *TypeDesc) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TypeDesc) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	_swUnion := uint16(o.VT)
	if o.Union != nil {
		if err := o.Union.MarshalUnionNDR(ctx, w, _swUnion); err != nil {
			return err
		}
	} else {
		if err := (&TypeDesc_Union{}).MarshalUnionNDR(ctx, w, _swUnion); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.VT); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *TypeDesc) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.Union == nil {
		o.Union = &TypeDesc_Union{}
	}
	_swUnion := uint16(o.VT)
	if err := o.Union.UnmarshalUnionNDR(ctx, w, _swUnion); err != nil {
		return err
	}
	if err := w.ReadData(&o.VT); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// TypeDesc_Union structure represents TYPEDESC union anonymous member.
//
// The TYPEDESC structure is used in the ARRAYDESC, ELEMDESC, and TYPEATTR structures
// to identify and describe the type of a data member, the return type of a method,
// or the type of a method parameter.
type TypeDesc_Union struct {
	// Types that are assignable to Value
	//
	// *TypeDesc_Union_TypeDesc
	// *TypeDesc_Union_ArrayDesc
	// *TypeDesc_Union_HrefType
	Value is_TypeDesc_Union `json:"value"`
}

func (o *TypeDesc_Union) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *TypeDesc_Union_TypeDesc:
		if value != nil {
			return value.TypeDesc
		}
	case *TypeDesc_Union_ArrayDesc:
		if value != nil {
			return value.ArrayDesc
		}
	case *TypeDesc_Union_HrefType:
		if value != nil {
			return value.HrefType
		}
	}
	return nil
}

type is_TypeDesc_Union interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_TypeDesc_Union()
}

func (o *TypeDesc_Union) NDRSwitchValue(sw uint16) uint16 {
	if o == nil {
		return uint16(0)
	}
	switch (interface{})(o.Value).(type) {
	case *TypeDesc_Union_TypeDesc:
		switch sw {
		case uint16(26),
			uint16(27):
			return sw
		}
		return uint16(26)
	case *TypeDesc_Union_ArrayDesc:
		return uint16(28)
	case *TypeDesc_Union_HrefType:
		return uint16(29)
	}
	return uint16(0)
}

func (o *TypeDesc_Union) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint16) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint16(sw)); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(26),
		uint16(27):
		_o, _ := o.Value.(*TypeDesc_Union_TypeDesc)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TypeDesc_Union_TypeDesc{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(28):
		_o, _ := o.Value.(*TypeDesc_Union_ArrayDesc)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TypeDesc_Union_ArrayDesc{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint16(29):
		_o, _ := o.Value.(*TypeDesc_Union_HrefType)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&TypeDesc_Union_HrefType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *TypeDesc_Union) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint16) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint16)(&sw)); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint16(26),
		uint16(27):
		o.Value = &TypeDesc_Union_TypeDesc{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(28):
		o.Value = &TypeDesc_Union_ArrayDesc{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint16(29):
		o.Value = &TypeDesc_Union_HrefType{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// TypeDesc_Union_TypeDesc structure represents TypeDesc_Union RPC union arm.
//
// It has following labels: 26, 27
type TypeDesc_Union_TypeDesc struct {
	// lptdesc:   MUST refer to a TYPEDESC that specifies the element type. If the ELEMDESC
	// is contained in a VARDESC that describes an appobject coclass, the TYPEDESC MUST
	// specify the type of the coclass.
	TypeDesc *TypeDesc `idl:"name:lptdesc" json:"type_desc"`
}

func (*TypeDesc_Union_TypeDesc) is_TypeDesc_Union() {}

func (o *TypeDesc_Union_TypeDesc) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.TypeDesc != nil {
		_ptr_lptdesc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.TypeDesc != nil {
				if err := o.TypeDesc.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&TypeDesc{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.TypeDesc, _ptr_lptdesc); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TypeDesc_Union_TypeDesc) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_lptdesc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.TypeDesc == nil {
			o.TypeDesc = &TypeDesc{}
		}
		if err := o.TypeDesc.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_lptdesc := func(ptr interface{}) { o.TypeDesc = *ptr.(**TypeDesc) }
	if err := w.ReadPointer(&o.TypeDesc, _s_lptdesc, _ptr_lptdesc); err != nil {
		return err
	}
	return nil
}

// TypeDesc_Union_ArrayDesc structure represents TypeDesc_Union RPC union arm.
//
// It has following labels: 28
type TypeDesc_Union_ArrayDesc struct {
	// lpadesc:   MUST refer to an ARRAYDESC that describes a fixed-length array.
	ArrayDesc *ArrayDesc `idl:"name:lpadesc" json:"array_desc"`
}

func (*TypeDesc_Union_ArrayDesc) is_TypeDesc_Union() {}

func (o *TypeDesc_Union_ArrayDesc) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.ArrayDesc != nil {
		_ptr_lpadesc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ArrayDesc != nil {
				if err := o.ArrayDesc.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ArrayDesc{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ArrayDesc, _ptr_lpadesc); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *TypeDesc_Union_ArrayDesc) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_lpadesc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ArrayDesc == nil {
			o.ArrayDesc = &ArrayDesc{}
		}
		if err := o.ArrayDesc.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_lpadesc := func(ptr interface{}) { o.ArrayDesc = *ptr.(**ArrayDesc) }
	if err := w.ReadPointer(&o.ArrayDesc, _s_lpadesc, _ptr_lpadesc); err != nil {
		return err
	}
	return nil
}

// TypeDesc_Union_HrefType structure represents TypeDesc_Union RPC union arm.
//
// It has following labels: 29
type TypeDesc_Union_HrefType struct {
	// hreftype:   MUST be set to an HREFTYPE that identifies the UDT (see section 2.2.28).
	HrefType uint32 `idl:"name:hreftype" json:"href_type"`
}

func (*TypeDesc_Union_HrefType) is_TypeDesc_Union() {}

func (o *TypeDesc_Union_HrefType) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.HrefType); err != nil {
		return err
	}
	return nil
}
func (o *TypeDesc_Union_HrefType) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.HrefType); err != nil {
		return err
	}
	return nil
}

// ArrayDesc structure represents ARRAYDESC RPC structure.
//
// The ARRAYDESC structure is used in a TYPEDESC structure to specify the dimensions
// of an array and the type of its elements.
type ArrayDesc struct {
	// tdescElem:  MUST contain a TYPEDESC that specifies the type of the elements in the
	// array as specified in section 2.2.37.
	TypeDescElem *TypeDesc `idl:"name:tdescElem" json:"type_desc_elem"`
	// cDims:  MUST be set to the number of dimensions in the array.
	DimsCount uint16 `idl:"name:cDims" json:"dims_count"`
	// rgbounds:  MUST refer to a SAFEARRAYBOUND that specifies the maximum index value
	// for each dimension of the array, as specified in section 2.2.30.1.
	Bounds []*SafeArrayBound `idl:"name:rgbounds;size_is:(cDims)" json:"bounds"`
}

func (o *ArrayDesc) xxx_PreparePayload(ctx context.Context) error {
	if o.Bounds != nil && o.DimsCount == 0 {
		o.DimsCount = uint16(len(o.Bounds))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *ArrayDesc) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DimsCount)
	return []uint64{
		dimSize1,
	}
}
func (o *ArrayDesc) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.TypeDescElem != nil {
		if err := o.TypeDescElem.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TypeDesc{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.DimsCount); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	for i1 := range o.Bounds {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.Bounds[i1] != nil {
			if err := o.Bounds[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SafeArrayBound{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Bounds); uint64(i1) < sizeInfo[0]; i1++ {
		if err := (&SafeArrayBound{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ArrayDesc) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.TypeDescElem == nil {
		o.TypeDescElem = &TypeDesc{}
	}
	if err := o.TypeDescElem.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.DimsCount); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.DimsCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.DimsCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Bounds", sizeInfo[0])
	}
	o.Bounds = make([]*SafeArrayBound, sizeInfo[0])
	for i1 := range o.Bounds {
		i1 := i1
		if o.Bounds[i1] == nil {
			o.Bounds[i1] = &SafeArrayBound{}
		}
		if err := o.Bounds[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// ParamDescEx structure represents PARAMDESCEX RPC structure.
//
// The PARAMDESCEX structure is used in a PARAMDESC (section 2.2.40) structure to specify
// information about the default value of a parameter.
type ParamDescEx struct {
	// cBytes:  MUST be set to an implementation-specific value.<20>
	BytesCount uint32 `idl:"name:cBytes" json:"bytes_count"`
	// varDefaultValue:  MUST contain a VARIANT that specifies the default value of the
	// parameter.
	VarDefaultValue *Variant `idl:"name:varDefaultValue" json:"var_default_value"`
}

func (o *ParamDescEx) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ParamDescEx) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.BytesCount); err != nil {
		return err
	}
	if o.VarDefaultValue != nil {
		_ptr_varDefaultValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.VarDefaultValue != nil {
				if err := o.VarDefaultValue.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Variant{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.VarDefaultValue, _ptr_varDefaultValue); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ParamDescEx) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.BytesCount); err != nil {
		return err
	}
	_ptr_varDefaultValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.VarDefaultValue == nil {
			o.VarDefaultValue = &Variant{}
		}
		if err := o.VarDefaultValue.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_varDefaultValue := func(ptr interface{}) { o.VarDefaultValue = *ptr.(**Variant) }
	if err := w.ReadPointer(&o.VarDefaultValue, _s_varDefaultValue, _ptr_varDefaultValue); err != nil {
		return err
	}
	return nil
}

// ParamDesc structure represents PARAMDESC RPC structure.
//
// The PARAMDESC structure is used in an ELEMDESC (section 2.2.41) structure to specify
// the features of a method parameter.
type ParamDesc struct {
	// pparamdescex:  MUST refer to a PARAMDESCEX structure that specifies the default value
	// of the parameter if the PARAMFLAG_FHASDEFAULT flag is set in the wParamFlags field.
	// MUST be set to NULL otherwise.
	ParamDescEx *ParamDescEx `idl:"name:pparamdescex" json:"param_desc_ex"`
	// wParamFlags:  MUST be set to a combination of the PARAMFLAG (section 2.2.15) bit
	// flags if the PARAMDESC belongs to an element of the lprgelemdescParam array in a
	// FUNCDESC (section 2.2.42) structure. MUST be set to 0 otherwise.
	ParamFlags uint16 `idl:"name:wParamFlags" json:"param_flags"`
}

func (o *ParamDesc) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ParamDesc) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if o.ParamDescEx != nil {
		_ptr_pparamdescex := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ParamDescEx != nil {
				if err := o.ParamDescEx.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ParamDescEx{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ParamDescEx, _ptr_pparamdescex); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ParamFlags); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(7); err != nil {
		return err
	}
	return nil
}
func (o *ParamDesc) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	_ptr_pparamdescex := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ParamDescEx == nil {
			o.ParamDescEx = &ParamDescEx{}
		}
		if err := o.ParamDescEx.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_pparamdescex := func(ptr interface{}) { o.ParamDescEx = *ptr.(**ParamDescEx) }
	if err := w.ReadPointer(&o.ParamDescEx, _s_pparamdescex, _ptr_pparamdescex); err != nil {
		return err
	}
	if err := w.ReadData(&o.ParamFlags); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(7); err != nil {
		return err
	}
	return nil
}

// ElemDesc structure represents ELEMDESC RPC structure.
//
// The ELEMDESC structure is used in the FUNCDESC (section 2.2.42) and VARDESC (section
// 2.2.43) structures to describe a member of a structure, a parameter, or the return
// value of a method.
type ElemDesc struct {
	// tdesc:  MUST contain a TYPEDESC (section 2.2.37) that describes the element, parameter,
	// or return value.
	TypeDesc *TypeDesc `idl:"name:tdesc" json:"type_desc"`
	// paramdesc:  MUST contain a PARAMDESC that has the values as specified in section
	// 2.2.40, if the ELEMDESC is a member of the lprgelemdescParam array in a FUNCDESC
	// (section 2.2.42) structure. Otherwise, it MUST contain a PARAMDESC that has the data
	// fields pparamdescex and wParamFlags set to NULL and 0 respectively.
	ParamDesc *ParamDesc `idl:"name:paramdesc" json:"param_desc"`
}

func (o *ElemDesc) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ElemDesc) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.TypeDesc != nil {
		if err := o.TypeDesc.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TypeDesc{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ParamDesc != nil {
		if err := o.ParamDesc.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ParamDesc{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *ElemDesc) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.TypeDesc == nil {
		o.TypeDesc = &TypeDesc{}
	}
	if err := o.TypeDesc.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ParamDesc == nil {
		o.ParamDesc = &ParamDesc{}
	}
	if err := o.ParamDesc.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// FuncDesc structure represents FUNCDESC RPC structure.
//
// The FUNCDESC structure is used by an ITypeComp server or ITypeInfo server to describe
// a method, as specified in sections 3.5.4.1 and 3.7.4.3.
type FuncDesc struct {
	// memid:  MUST be set to the MEMBERID (section 2.2.35) of the method.
	MemberID int32 `idl:"name:memid" json:"member_id"`
	// lReserved1:  MUST be set to 0 and ignored on receipt. An HRESULT value is closely
	// related, or identical to an SCODE.
	_ []int32 `idl:"name:lReserved1;size_is:(cReserved2)"`
	// lprgelemdescParam:  MUST refer to an array of ELEMDESC that contains one entry for
	// each element in the method's parameter table.
	LprgelemdescParam []*ElemDesc `idl:"name:lprgelemdescParam;size_is:(cParams)" json:"lprgelemdesc_param"`
	// funckind:  MUST be set to one of the values of the FUNCKIND (section 2.2.12) enumeration.
	FuncKind FuncKind `idl:"name:funckind" json:"func_kind"`
	// invkind:  MUST be set to one of the values of the INVOKEKIND (section 2.2.14) enumeration.
	InvokeKind InvokeKind `idl:"name:invkind" json:"invoke_kind"`
	// callconv:  MUST be set to one of the values of the CALLCONV (section 2.2.10) enumeration.
	CallConv CallConv `idl:"name:callconv" json:"call_conv"`
	// cParams:  MUST be set to the length of the lprgelemdescParam array.
	ParamsCount int16 `idl:"name:cParams" json:"params_count"`
	// cParamsOpt:  SHOULD be set to the number of optional VARIANT parameters<21>. MUST
	// be set to -1 if the method was declared with the [vararg] attribute. Otherwise, MUST
	// be set to 0.
	ParamsOptCount int16 `idl:"name:cParamsOpt" json:"params_opt_count"`
	// oVft:  MUST be set to either 0 or to the opnum of the interface method multiplied
	// by the system pointer size value (as specified in sections 2.2.44 and 3.11.1).
	VFT int16 `idl:"name:oVft" json:"vft"`
	// cReserved2:  MUST be set to 0, and ignored on receipt.
	_ int16 `idl:"name:cReserved2"`
	// elemdescFunc:  MUST contain an ELEMDESC that specifies the return type of the method,
	// as specified in section 2.2.41.
	ElemDescFunc *ElemDesc `idl:"name:elemdescFunc" json:"elem_desc_func"`
	// wFuncFlags:  MUST be set to a combination of the FUNCFLAGS bit flags (as specified
	// in section 2.2.11), or set to 0.
	FuncFlags uint16 `idl:"name:wFuncFlags" json:"func_flags"`
}

func (o *FuncDesc) xxx_PreparePayload(ctx context.Context) error {
	if o.LprgelemdescParam != nil && o.ParamsCount == 0 {
		o.ParamsCount = int16(len(o.LprgelemdescParam))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *FuncDesc) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.MemberID); err != nil {
		return err
	}
	// reserved lReserved1
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	if o.LprgelemdescParam != nil || o.ParamsCount > 0 {
		_ptr_lprgelemdescParam := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.ParamsCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.LprgelemdescParam {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.LprgelemdescParam[i1] != nil {
					if err := o.LprgelemdescParam[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ElemDesc{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.LprgelemdescParam); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&ElemDesc{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.LprgelemdescParam, _ptr_lprgelemdescParam); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteEnum(uint32(o.FuncKind)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint32(o.InvokeKind)); err != nil {
		return err
	}
	if err := w.WriteEnum(uint32(o.CallConv)); err != nil {
		return err
	}
	if err := w.WriteData(o.ParamsCount); err != nil {
		return err
	}
	if err := w.WriteData(o.ParamsOptCount); err != nil {
		return err
	}
	if err := w.WriteData(o.VFT); err != nil {
		return err
	}
	// reserved cReserved2
	if err := w.WriteData(int16(0)); err != nil {
		return err
	}
	if o.ElemDescFunc != nil {
		if err := o.ElemDescFunc.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ElemDesc{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.FuncFlags); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *FuncDesc) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.MemberID); err != nil {
		return err
	}
	// reserved lReserved1
	var _lReserved1 []int32
	_ptr_lReserved1 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _lReserved1", sizeInfo[0])
		}
		_lReserved1 = make([]int32, sizeInfo[0])
		for i1 := range _lReserved1 {
			i1 := i1
			if err := w.ReadData(&_lReserved1[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_lReserved1 := func(ptr interface{}) { _lReserved1 = *ptr.(*[]int32) }
	if err := w.ReadPointer(&_lReserved1, _s_lReserved1, _ptr_lReserved1); err != nil {
		return err
	}
	_ptr_lprgelemdescParam := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.ParamsCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.ParamsCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.LprgelemdescParam", sizeInfo[0])
		}
		o.LprgelemdescParam = make([]*ElemDesc, sizeInfo[0])
		for i1 := range o.LprgelemdescParam {
			i1 := i1
			if o.LprgelemdescParam[i1] == nil {
				o.LprgelemdescParam[i1] = &ElemDesc{}
			}
			if err := o.LprgelemdescParam[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_lprgelemdescParam := func(ptr interface{}) { o.LprgelemdescParam = *ptr.(*[]*ElemDesc) }
	if err := w.ReadPointer(&o.LprgelemdescParam, _s_lprgelemdescParam, _ptr_lprgelemdescParam); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint32)(&o.FuncKind)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint32)(&o.InvokeKind)); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint32)(&o.CallConv)); err != nil {
		return err
	}
	if err := w.ReadData(&o.ParamsCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.ParamsOptCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.VFT); err != nil {
		return err
	}
	// reserved cReserved2
	var _cReserved2 int16
	if err := w.ReadData(&_cReserved2); err != nil {
		return err
	}
	if o.ElemDescFunc == nil {
		o.ElemDescFunc = &ElemDesc{}
	}
	if err := o.ElemDescFunc.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.FuncFlags); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// VarDesc structure represents VARDESC RPC structure.
//
// The VARDESC structure is used by an ITypeInfo server or ITypeComp server to describe
// a data member, constant, or ODL dispinterface property, as specified in sections
// 3.5.4.1 and 3.7.4.4.
type VarDesc struct {
	// memid:  MUST be set to the MEMBERID (section 2.2.35) of the data member, the constant,
	// or the ODL dispinterface property. MUST be set to MEMBERID_DEFAULTINST if the VARDESC
	// describes an appobject coclass, as specified in section 2.2.49.8.
	MemberID int32 `idl:"name:memid" json:"member_id"`
	// lpstrReserved:  MUST be set to NULL, and MUST be ignored by the recipient.
	_ string `idl:"name:lpstrReserved"`
	// _vdUnion:  MUST be set to an instance of the type, according to the value in the
	// varkind field.
	Union *VarDesc_Union `idl:"name:_vdUnion;switch_is:varkind" json:"union"`
	// elemdescVar:  MUST contain an ELEMDESC that describes the data member, constant,
	// or ODL dispinterface property and its type, as specified in section 2.2.41.
	ElemDescVar *ElemDesc `idl:"name:elemdescVar" json:"elem_desc_var"`
	// wVarFlags:  MUST be set to a combination of the VARFLAGS bit flags (as specified
	// in 2.2.18), or set to 0. MUST be set to 0 if the VARDESC describes an appobject coclass,
	// as specified in section 2.2.49.8.
	VarFlags uint16 `idl:"name:wVarFlags" json:"var_flags"`
	// varkind:  MUST be set to a value of the VARKIND enumeration. MUST be set to VAR_STATIC
	// if the VARDESC describes an appobject coclass, as specified in section 2.2.49.8.
	VarKind VarKind `idl:"name:varkind" json:"var_kind"`
}

func (o *VarDesc) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *VarDesc) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.MemberID); err != nil {
		return err
	}
	// reserved lpstrReserved
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	_swUnion := uint32(o.VarKind)
	if o.Union != nil {
		if err := o.Union.MarshalUnionNDR(ctx, w, _swUnion); err != nil {
			return err
		}
	} else {
		if err := (&VarDesc_Union{}).MarshalUnionNDR(ctx, w, _swUnion); err != nil {
			return err
		}
	}
	if o.ElemDescVar != nil {
		if err := o.ElemDescVar.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ElemDesc{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.VarFlags); err != nil {
		return err
	}
	if err := w.WriteEnum(uint32(o.VarKind)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *VarDesc) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.MemberID); err != nil {
		return err
	}
	// reserved lpstrReserved
	var _lpstrReserved string
	_ptr_lpstrReserved := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &_lpstrReserved); err != nil {
			return err
		}
		return nil
	})
	_s_lpstrReserved := func(ptr interface{}) { _lpstrReserved = *ptr.(*string) }
	if err := w.ReadPointer(&_lpstrReserved, _s_lpstrReserved, _ptr_lpstrReserved); err != nil {
		return err
	}
	if o.Union == nil {
		o.Union = &VarDesc_Union{}
	}
	_swUnion := uint32(o.VarKind)
	if err := o.Union.UnmarshalUnionNDR(ctx, w, _swUnion); err != nil {
		return err
	}
	if o.ElemDescVar == nil {
		o.ElemDescVar = &ElemDesc{}
	}
	if err := o.ElemDescVar.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.VarFlags); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint32)(&o.VarKind)); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// VarDesc_Union structure represents VARDESC union anonymous member.
//
// The VARDESC structure is used by an ITypeInfo server or ITypeComp server to describe
// a data member, constant, or ODL dispinterface property, as specified in sections
// 3.5.4.1 and 3.7.4.4.
type VarDesc_Union struct {
	// Types that are assignable to Value
	//
	// *VarDesc_Union_Instance
	// *VarDesc_Union_Value
	Value is_VarDesc_Union `json:"value"`
}

func (o *VarDesc_Union) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *VarDesc_Union_Instance:
		if value != nil {
			return value.Instance
		}
	case *VarDesc_Union_Value:
		if value != nil {
			return value.Value
		}
	}
	return nil
}

type is_VarDesc_Union interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_VarDesc_Union()
}

func (o *VarDesc_Union) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *VarDesc_Union_Instance:
		switch sw {
		case uint32(0),
			uint32(3),
			uint32(1):
			return sw
		}
		return uint32(0)
	case *VarDesc_Union_Value:
		return uint32(2)
	}
	return uint32(0)
}

func (o *VarDesc_Union) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(ndr.Enum(uint32(sw))); err != nil {
		return err
	}
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(0),
		uint32(3),
		uint32(1):
		_o, _ := o.Value.(*VarDesc_Union_Instance)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&VarDesc_Union_Instance{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*VarDesc_Union_Value)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&VarDesc_Union_Value{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *VarDesc_Union) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch(ndr.Enum((*uint32)(&sw))); err != nil {
		return err
	}
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(0),
		uint32(3),
		uint32(1):
		o.Value = &VarDesc_Union_Instance{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &VarDesc_Union_Value{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// VarDesc_Union_Instance structure represents VarDesc_Union RPC union arm.
//
// It has following labels: 0, 3, 1
type VarDesc_Union_Instance struct {
	// oInst:
	//
	// * VAR_PERINSTANCE: MUST be set to an implementation-specific value <22> ( ae19ad63-7433-4568-88e9-f70e5593547d#Appendix_A_22
	// )
	//
	// * VAR_DISPATCH: MUST be set to 0.
	//
	// * VAR_STATIC: MUST be set to 0.
	Instance uint32 `idl:"name:oInst" json:"instance"`
}

func (*VarDesc_Union_Instance) is_VarDesc_Union() {}

func (o *VarDesc_Union_Instance) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := w.WriteData(o.Instance); err != nil {
		return err
	}
	return nil
}
func (o *VarDesc_Union_Instance) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadData(&o.Instance); err != nil {
		return err
	}
	return nil
}

// VarDesc_Union_Value structure represents VarDesc_Union RPC union arm.
//
// It has following labels: 2
type VarDesc_Union_Value struct {
	Value *Variant `idl:"name:lpvarValue" json:"value"`
}

func (*VarDesc_Union_Value) is_VarDesc_Union() {}

func (o *VarDesc_Union_Value) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Value != nil {
		_ptr_lpvarValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Value != nil {
				_ptr_lpvarValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.Value != nil {
						if err := o.Value.MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&Variant{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.Value, _ptr_lpvarValue); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Value, _ptr_lpvarValue); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *VarDesc_Union_Value) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_lpvarValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		_ptr_lpvarValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Value == nil {
				o.Value = &Variant{}
			}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpvarValue := func(ptr interface{}) { o.Value = *ptr.(**Variant) }
		if err := w.ReadPointer(&o.Value, _s_lpvarValue, _ptr_lpvarValue); err != nil {
			return err
		}
		return nil
	})
	_s_lpvarValue := func(ptr interface{}) { o.Value = *ptr.(**Variant) }
	if err := w.ReadPointer(&o.Value, _s_lpvarValue, _ptr_lpvarValue); err != nil {
		return err
	}
	return nil
}

// TypeAttribute structure represents TYPEATTR RPC structure.
//
// The TYPEATTR structure is used by an ITypeInfo server to describe a type, as specified
// in section 3.7.4.1.
type TypeAttribute struct {
	// guid:  MUST be set to the GUID that is associated with the type, or to IID_NULL,
	// if the type was not declared with the [uuid] attribute (see section 2.2.49.2).
	GUID *dtyp.GUID `idl:"name:guid" json:"guid"`
	// lcid:  MUST be set to the locale ID of the type's member names and documentation
	// strings (see section 2.2.49.2).
	LocaleID uint32 `idl:"name:lcid" json:"locale_id"`
	// dwReserved1:  MUST be set to 0, and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved1"`
	// dwReserved2:  MUST be set to -1, and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved2"`
	// dwReserved3:  MUST be set to -1, and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved3"`
	// lpstrReserved4:  MUST be set to NULL, and MUST be ignored on receipt.
	_ string `idl:"name:lpstrReserved4"`
	// cbSizeInstance:  MUST be set to a value that is specified by the value of typekind.
	//
	//	+-------------------+----------------------------------------------------------------------------------+
	//	|     VALUE OF      |                                     VALUE OF                                     |
	//	|     TYPEKIND      |                                  CBSIZEINSTANCE                                  |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_COCLASS     | MUST be set to the system pointer size (see section 3.7.1.2).                    |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_DISPATCH    | MUST be set to the system pointer size (see section 3.7.1.2).                    |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_INTERFACE   | MUST be set to the system pointer size (see section 3.7.1.2).                    |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_MODULE      | MUST be set to 2.                                                                |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_ENUM        | MUST be set to an implementation-specific value<23> that specifies the size of   |
	//	|                   | an integer.                                                                      |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_UNION       | MUST be set to an implementation-specific value<24> that specifies the size of   |
	//	|                   | its largest element.                                                             |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_RECORD      | MUST be set to an implementation-specific value<25> that specifies the size in   |
	//	|                   | bytes, of the structure.                                                         |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_ALIAS       | MUST be set to an implementation-specific value<26> that specifies the size, in  |
	//	|                   | bytes, of the predefined type for which this type is an alias.                   |
	//	+-------------------+----------------------------------------------------------------------------------+
	LengthInstance uint32 `idl:"name:cbSizeInstance" json:"length_instance"`
	// typekind:  MUST be set to a value of the TYPEKIND enumeration, as specified in section
	// 2.2.17.
	TypeKind TypeKind `idl:"name:typekind" json:"type_kind"`
	// cFuncs:  MUST be set to a value specified by the value of typekind.
	//
	//	+-------------------+----------------------------------------------------------------------------------+
	//	|     VALUE OF      |                                     VALUE OF                                     |
	//	|     TYPEKIND      |                                      CFUNCS                                      |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_COCLASS     | MUST be set to 0.                                                                |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_DISPATCH    | MUST be set to the number of elements in the dispatch method table, as specified |
	//	|                   | in section 3.7.1.2.                                                              |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_INTERFACE   | MUST be set to the number of elements in the method table, as specified in       |
	//	|                   | section 3.7.1.2.                                                                 |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_MODULE      | MUST be set to the number of elements in the method table, as specified in       |
	//	|                   | section 3.7.1.2.                                                                 |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_ENUM        | MUST be set to 0.                                                                |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_UNION       | MUST be set to 0.                                                                |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_RECORD      | MUST be set to 0.                                                                |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_ALIAS       | MUST be set to 0.                                                                |
	//	+-------------------+----------------------------------------------------------------------------------+
	FuncsCount uint16 `idl:"name:cFuncs" json:"funcs_count"`
	// cVars:  MUST be set to the number of elements in the data member table, as specified
	// in section 3.7.1.2.
	VarsCount uint16 `idl:"name:cVars" json:"vars_count"`
	// cImplTypes:  MUST be set to the number of elements in the interface table, as specified
	// in section 3.7.1.2.
	ImplTypesCount uint16 `idl:"name:cImplTypes" json:"impl_types_count"`
	// cbSizeVft:  MUST be set to a value specified by the value of typekind.
	//
	//	+-------------------+----------------------------------------------------------------------------------+
	//	|     VALUE OF      |                                     VALUE OF                                     |
	//	|     TYPEKIND      |                                    CBSIZEVFT                                     |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_COCLASS     | MUST be set to 0.                                                                |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_DISPATCH    | MUST be set to the system pointer size value (see section 2.2.45) multiplied by  |
	//	|                   | seven.                                                                           |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_INTERFACE   | MUST be set to the system pointer size value multiplied by the number of methods |
	//	|                   | that are defined by the interface and all its inherited interfaces.              |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_MODULE      | MUST be set to 0.                                                                |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_ENUM        | MUST be set to 0.                                                                |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_UNION       | MUST be set to 0.                                                                |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_RECORD      | MUST be set to 0.                                                                |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| TKIND_ALIAS       | MUST be set to 0.                                                                |
	//	+-------------------+----------------------------------------------------------------------------------+
	LengthVFT uint16 `idl:"name:cbSizeVft" json:"length_vft"`
	// cbAlignment:  MUST be set to 0 or to an implementation-specific positive value.<27>
	AlignmentLength uint16 `idl:"name:cbAlignment" json:"alignment_length"`
	// wTypeFlags:  MUST be either a combination of the TYPEFLAGS bit flags that are specified
	// in section 2.2.16, or 0.
	TypeFlags uint16 `idl:"name:wTypeFlags" json:"type_flags"`
	// wMajorVerNum:  MUST be set to the major version number of the automation scope that
	// is associated with the ITypeLib server, as specified in section 2.2.49.2.
	MajorVerNum uint16 `idl:"name:wMajorVerNum" json:"major_ver_num"`
	// wMinorVerNum:  MUST be set to the minor version number of the automation scope that
	// is associated with the ITypeLib server, as specified in section 2.2.49.2.
	MinorVerNum uint16 `idl:"name:wMinorVerNum" json:"minor_ver_num"`
	// tdescAlias:  MUST contain a TYPEDESC (section 2.2.37) that describes the predefined
	// type for which this type is an alias, if typekind is set to TKIND_ALIAS. Otherwise,
	// MUST contain a TYPEDESC with the vt field set to VT_EMPTY.
	TypeDescAlias *TypeDesc `idl:"name:tdescAlias" json:"type_desc_alias"`
	// dwReserved5:  MUST be set to 0, and MUST be ignored on receipt.
	_ uint32 `idl:"name:dwReserved5"`
	// wReserved6:  MUST be set to 0, and MUST be ignored on receipt.
	_ uint16 `idl:"name:wReserved6"`
}

func (o *TypeAttribute) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TypeAttribute) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.GUID != nil {
		if err := o.GUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LocaleID); err != nil {
		return err
	}
	// reserved dwReserved1
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved dwReserved2
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved dwReserved3
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved lpstrReserved4
	if err := w.WritePointer(nil); err != nil {
		return err
	}
	if err := w.WriteData(o.LengthInstance); err != nil {
		return err
	}
	if err := w.WriteEnum(uint32(o.TypeKind)); err != nil {
		return err
	}
	if err := w.WriteData(o.FuncsCount); err != nil {
		return err
	}
	if err := w.WriteData(o.VarsCount); err != nil {
		return err
	}
	if err := w.WriteData(o.ImplTypesCount); err != nil {
		return err
	}
	if err := w.WriteData(o.LengthVFT); err != nil {
		return err
	}
	if err := w.WriteData(o.AlignmentLength); err != nil {
		return err
	}
	if err := w.WriteData(o.TypeFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.MajorVerNum); err != nil {
		return err
	}
	if err := w.WriteData(o.MinorVerNum); err != nil {
		return err
	}
	if o.TypeDescAlias != nil {
		if err := o.TypeDescAlias.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TypeDesc{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// reserved dwReserved5
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	// reserved wReserved6
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(9); err != nil {
		return err
	}
	return nil
}
func (o *TypeAttribute) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.GUID == nil {
		o.GUID = &dtyp.GUID{}
	}
	if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocaleID); err != nil {
		return err
	}
	// reserved dwReserved1
	var _dwReserved1 uint32
	if err := w.ReadData(&_dwReserved1); err != nil {
		return err
	}
	// reserved dwReserved2
	var _dwReserved2 uint32
	if err := w.ReadData(&_dwReserved2); err != nil {
		return err
	}
	// reserved dwReserved3
	var _dwReserved3 uint32
	if err := w.ReadData(&_dwReserved3); err != nil {
		return err
	}
	// reserved lpstrReserved4
	var _lpstrReserved4 string
	_ptr_lpstrReserved4 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &_lpstrReserved4); err != nil {
			return err
		}
		return nil
	})
	_s_lpstrReserved4 := func(ptr interface{}) { _lpstrReserved4 = *ptr.(*string) }
	if err := w.ReadPointer(&_lpstrReserved4, _s_lpstrReserved4, _ptr_lpstrReserved4); err != nil {
		return err
	}
	if err := w.ReadData(&o.LengthInstance); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint32)(&o.TypeKind)); err != nil {
		return err
	}
	if err := w.ReadData(&o.FuncsCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.VarsCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.ImplTypesCount); err != nil {
		return err
	}
	if err := w.ReadData(&o.LengthVFT); err != nil {
		return err
	}
	if err := w.ReadData(&o.AlignmentLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.TypeFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.MajorVerNum); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinorVerNum); err != nil {
		return err
	}
	if o.TypeDescAlias == nil {
		o.TypeDescAlias = &TypeDesc{}
	}
	if err := o.TypeDescAlias.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	// reserved dwReserved5
	var _dwReserved5 uint32
	if err := w.ReadData(&_dwReserved5); err != nil {
		return err
	}
	// reserved wReserved6
	var _wReserved6 uint16
	if err := w.ReadData(&_wReserved6); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(9); err != nil {
		return err
	}
	return nil
}

// TypeLibAttribute structure represents TLIBATTR RPC structure.
//
// The TLIBATTR structure is used to specify the attributes of an ITypeLib server, as
// specified in section 3.11.4.
type TypeLibAttribute struct {
	// guid:  MUST be set to the GUID of the automation scope that is associated with the
	// ITypeLib server, as specified in section 2.2.49.1.
	GUID *dtyp.GUID `idl:"name:guid" json:"guid"`
	// lcid:  MUST be set to the LCID of the automation scope that is associated with the
	// ITypeLib server, as specified in section 2.2.49.1.
	LocaleID uint32 `idl:"name:lcid" json:"locale_id"`
	// syskind:  MUST be set to a value of the SYSKIND enumeration, as specified in section
	// 2.2.21.
	//
	// The value of syskind specifies the system pointer-size value. If syskind is SYS_WIN32,
	// the system pointer-size value is 4. If syskind is SYS_WIN64, the system pointer-size
	// value is 8.
	SystemKind SystemKind `idl:"name:syskind" json:"system_kind"`
	// wMajorVerNum:  MUST be set to the major version number of the automation scope that
	// is associated with the ITypeLib server, as specified in section 2.2.49.2.
	MajorVerNum uint16 `idl:"name:wMajorVerNum" json:"major_ver_num"`
	// wMinorVerNum:  MUST be set to the minor version number of the automation scope that
	// is associated with the ITypeLib server, as specified in section 2.2.49.2.
	MinorVerNum uint16 `idl:"name:wMinorVerNum" json:"minor_ver_num"`
	// wLibFlags:  MUST be either a combination of the LIBFLAGS bit flags (as specified
	// in section 2.2.20) or 0.
	LibFlags uint16 `idl:"name:wLibFlags" json:"lib_flags"`
}

func (o *TypeLibAttribute) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TypeLibAttribute) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.GUID != nil {
		if err := o.GUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LocaleID); err != nil {
		return err
	}
	if err := w.WriteEnum(uint32(o.SystemKind)); err != nil {
		return err
	}
	if err := w.WriteData(o.MajorVerNum); err != nil {
		return err
	}
	if err := w.WriteData(o.MinorVerNum); err != nil {
		return err
	}
	if err := w.WriteData(o.LibFlags); err != nil {
		return err
	}
	if err := w.WriteTrailingGap(4); err != nil {
		return err
	}
	return nil
}
func (o *TypeLibAttribute) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.GUID == nil {
		o.GUID = &dtyp.GUID{}
	}
	if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocaleID); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint32)(&o.SystemKind)); err != nil {
		return err
	}
	if err := w.ReadData(&o.MajorVerNum); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinorVerNum); err != nil {
		return err
	}
	if err := w.ReadData(&o.LibFlags); err != nil {
		return err
	}
	if err := w.ReadTrailingGap(4); err != nil {
		return err
	}
	return nil
}

// CustomDataItem structure represents CUSTDATAITEM RPC structure.
//
// The CUSTDATAITEM structure is used in a CUSTDATA structure to store custom data items,
// as specified in section 2.2.47.
type CustomDataItem struct {
	// guid:  MUST be set to the GUID associated with the custom data item that uses the
	// [custom] attribute, as specified in section 2.2.49.5.1.
	GUID *dtyp.GUID `idl:"name:guid" json:"guid"`
	// varValue:  MUST be set to the value of the custom data item.
	Value *Variant `idl:"name:varValue" json:"value"`
}

func (o *CustomDataItem) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CustomDataItem) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.GUID != nil {
		if err := o.GUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.Value != nil {
		_ptr_varValue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Value != nil {
				if err := o.Value.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Variant{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Value, _ptr_varValue); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CustomDataItem) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if o.GUID == nil {
		o.GUID = &dtyp.GUID{}
	}
	if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	_ptr_varValue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Value == nil {
			o.Value = &Variant{}
		}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_varValue := func(ptr interface{}) { o.Value = *ptr.(**Variant) }
	if err := w.ReadPointer(&o.Value, _s_varValue, _ptr_varValue); err != nil {
		return err
	}
	return nil
}

// CustomData structure represents CUSTDATA RPC structure.
//
// The CUSTDATA structure is used by an ITypeInfo2 server or ITypeLib2 server to retrieve
// custom data items, as specified in sections 3.9.4 and 3.13.4.
type CustomData struct {
	// cCustData:  MUST be set to the number of custom data items in prgCustData.
	CustomDataCount uint32 `idl:"name:cCustData" json:"custom_data_count"`
	// prgCustData:  MUST refer to an array of CUSTDATAITEM structures that contain custom
	// data items, as specified in section 2.2.46.
	CustomData []*CustomDataItem `idl:"name:prgCustData;size_is:(cCustData)" json:"custom_data"`
}

func (o *CustomData) xxx_PreparePayload(ctx context.Context) error {
	if o.CustomData != nil && o.CustomDataCount == 0 {
		o.CustomDataCount = uint32(len(o.CustomData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CustomData) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.CustomDataCount); err != nil {
		return err
	}
	if o.CustomData != nil || o.CustomDataCount > 0 {
		_ptr_prgCustData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.CustomDataCount)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.CustomData {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.CustomData[i1] != nil {
					if err := o.CustomData[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&CustomDataItem{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.CustomData); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&CustomDataItem{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CustomData, _ptr_prgCustData); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *CustomData) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.CustomDataCount); err != nil {
		return err
	}
	_ptr_prgCustData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.CustomDataCount > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.CustomDataCount)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.CustomData", sizeInfo[0])
		}
		o.CustomData = make([]*CustomDataItem, sizeInfo[0])
		for i1 := range o.CustomData {
			i1 := i1
			if o.CustomData[i1] == nil {
				o.CustomData[i1] = &CustomDataItem{}
			}
			if err := o.CustomData[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_prgCustData := func(ptr interface{}) { o.CustomData = *ptr.(*[]*CustomDataItem) }
	if err := w.ReadPointer(&o.CustomData, _s_prgCustData, _ptr_prgCustData); err != nil {
		return err
	}
	return nil
}

// TypeInfo2 structure represents ITypeInfo2 RPC structure.
type TypeInfo2 dcom.InterfacePointer

func (o *TypeInfo2) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *TypeInfo2) xxx_PreparePayload(ctx context.Context) error {
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

func (o *TypeInfo2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *TypeInfo2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *TypeInfo2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// TypeLib2 structure represents ITypeLib2 RPC structure.
type TypeLib2 dcom.InterfacePointer

func (o *TypeLib2) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *TypeLib2) xxx_PreparePayload(ctx context.Context) error {
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

func (o *TypeLib2) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *TypeLib2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *TypeLib2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// Dispatch structure represents IDispatch RPC structure.
type Dispatch dcom.InterfacePointer

func (o *Dispatch) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *Dispatch) xxx_PreparePayload(ctx context.Context) error {
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

func (o *Dispatch) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *Dispatch) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Dispatch) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// TypeComp structure represents ITypeComp RPC structure.
type TypeComp dcom.InterfacePointer

func (o *TypeComp) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *TypeComp) xxx_PreparePayload(ctx context.Context) error {
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

func (o *TypeComp) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *TypeComp) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *TypeComp) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// TypeLib structure represents ITypeLib RPC structure.
type TypeLib dcom.InterfacePointer

func (o *TypeLib) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *TypeLib) xxx_PreparePayload(ctx context.Context) error {
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

func (o *TypeLib) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *TypeLib) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *TypeLib) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// EnumVariant structure represents IEnumVARIANT RPC structure.
type EnumVariant dcom.InterfacePointer

func (o *EnumVariant) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *EnumVariant) xxx_PreparePayload(ctx context.Context) error {
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

func (o *EnumVariant) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EnumVariant) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *EnumVariant) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// TypeInfo structure represents ITypeInfo RPC structure.
type TypeInfo dcom.InterfacePointer

func (o *TypeInfo) InterfacePointer() *dcom.InterfacePointer { return (*dcom.InterfacePointer)(o) }

func (o *TypeInfo) xxx_PreparePayload(ctx context.Context) error {
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

func (o *TypeInfo) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.DataCount)
	return []uint64{
		dimSize1,
	}
}
func (o *TypeInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *TypeInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
