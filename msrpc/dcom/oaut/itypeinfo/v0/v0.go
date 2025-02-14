package itypeinfo

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = iunknown.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/oaut"
)

var (
	// ITypeInfo interface identifier 00020401-0000-0000-c000-000000000046
	TypeInfoIID = &dcom.IID{Data1: 0x00020401, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	// Syntax UUID
	TypeInfoSyntaxUUID = &uuid.UUID{TimeLow: 0x20401, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	// Syntax ID
	TypeInfoSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: TypeInfoSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ITypeInfo interface.
type TypeInfoClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The GetTypeAttr method retrieves a TYPEATTR structure that contains information about
	// the type described by the ITypeInfo server.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	GetTypeAttribute(context.Context, *GetTypeAttributeRequest, ...dcerpc.CallOption) (*GetTypeAttributeResponse, error)

	// The GetTypeComp method retrieves a reference to the ITypeComp server instance associated
	// with the ITypeInfo server.
	//
	// The GetTypeComp method retrieves a reference to the ITypeComp server instance associated
	// with the ITypeLib server.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the sevierty bit is set to 1, the method failed and encountered a fatal error.
	GetTypeComp(context.Context, *GetTypeCompRequest, ...dcerpc.CallOption) (*GetTypeCompResponse, error)

	// The GetFuncDesc method retrieves a FUNCDESC structure that contains information about
	// a member of the ITypeInfo server's method or dispatch method table.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD does not match a value
	// in the following table, a fatal failure occurred.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD matches a value in
	// the following table, a failure occurred.
	//
	// Return value/code
	//
	// # Description
	//
	// 0x8002802B
	//
	// TYPE_E_ELEMENTNOTFOUND
	//
	// The value of index did not specify the ordinal position of an element in the method
	// table.
	GetFuncDesc(context.Context, *GetFuncDescRequest, ...dcerpc.CallOption) (*GetFuncDescResponse, error)

	// The GetVarDesc method retrieves a VARDESC structure that contains information about
	// a member of the ITypeInfo server's data member table.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD does not match a value
	// in the following table, a fatal failure occurred.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD matches a value in
	// the following table, a failure occurred.
	//
	// Return value/code
	//
	// # Description
	//
	// 0x8002802B
	//
	// TYPE_E_ELEMENTNOTFOUND
	//
	// The value of index and invKind did not specify the ordinal position of an element
	// in the method table. See [MS-ERREF].
	GetVarDesc(context.Context, *GetVarDescRequest, ...dcerpc.CallOption) (*GetVarDescResponse, error)

	// The GetNames method retrieves the data member name or the method and parameter names
	// associated with a specified MEMBERID.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD does not match a value
	// in the following table, a fatal failure occurred.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD matches a value in
	// the following table, a failure occurred.
	//
	// Return value/code
	//
	// # Description
	//
	// 0x8002802B
	//
	// TYPE_E_ELEMENTNOTFOUND
	//
	// The value of memid did not specify the MEMBERID of a member of the type. See [MS-ERREF].
	//
	// If the memid field corresponds to multiple property accessor methods, the contents
	// of rgBstrNames MUST correspond to the [propget] property accessor.
	//
	// If the ITypeInfo server represents an appobject coclass (see section 2.2.49.8) and
	// memid is MEMBERID_DEFAULTINST, the first element of rgBstrNames MUST be set to the
	// name of the coclass.
	//
	// In all other cases, the first element of rgBstrNames MUST be set to the name of the
	// method or data element in the binding context of the ITypeInfo server that corresponds
	// to the value of memid.
	//
	// If memid specifies a method or property accessor method, the remaining elements of
	// rgBstrNames MUST be set to the names of entries in its parameter table, in the order
	// in which they are stored in the parameter table.
	//
	// If memid specifies a put or putref property, the rgBstrNames array MUST NOT include
	// the name of the [retval] parameter. If memid specifies a member of a dispinterface,
	// the rgBstrNames array MUST NOT include the names of [lcid] or [retval] parameters
	// (see section 3.1.4.4). In all other cases, the rgBstrNames array MUST include all
	// members of the parameter table.
	GetNames(context.Context, *GetNamesRequest, ...dcerpc.CallOption) (*GetNamesResponse, error)

	// The GetRefTypeOfImplType method retrieves the HREFTYPE corresponding to the automation
	// type description of an interface that is inherited, implemented, or referenced by
	// the ITypeInfo server.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD does not match a value
	// in the following table, a fatal failure occurred.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD matches a value in
	// the following table, a failure occurred.
	//
	// Return value/code
	//
	// # Description
	//
	// 0x8002802B
	//
	// TYPE_E_ELEMENTNOTFOUND
	//
	// The value of index did not specify the ordinal position of an element in the interface
	// table, or the value of index was -1 and the type was not a dual interface. See [MS-ERREF].
	GetReferenceTypeOfImplType(context.Context, *GetReferenceTypeOfImplTypeRequest, ...dcerpc.CallOption) (*GetReferenceTypeOfImplTypeResponse, error)

	// The GetImplTypeFlags method retrieves the IMPLTYPEFLAGS values associated with an
	// interface member of a coclass.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD does not match a value
	// in the following table, a fatal failure occurred.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD matches a value in
	// the following table, a failure occurred.
	//
	// Return value/code
	//
	// # Description
	//
	// 0x8002802B
	//
	// TYPE_E_ELEMENTNOTFOUND
	//
	// The value of index did not specify the ordinal position of an element in the interface
	// table. See [MS-ERREF].
	GetImplTypeFlags(context.Context, *GetImplTypeFlagsRequest, ...dcerpc.CallOption) (*GetImplTypeFlagsResponse, error)

	// Opnum10NotUsedOnWire operation.
	// Opnum10NotUsedOnWire

	// Opnum11NotUsedOnWire operation.
	// Opnum11NotUsedOnWire

	// The GetDocumentation method retrieves the documentation resources associated with
	// a type member.
	//
	// The GetDocumentation method retrieves the documentation resources associated with
	// the automation type library.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// If memid is MEMBERID_NIL, the values of pBstrName, pBstrDocString, pdwHelpContext,
	// and pBstrHelpFile MUST correspond to the attributes declared with the type, as specified
	// in section 2.2.49.3. Otherwise, they MUST correspond to the attributes declared with
	// the specified member of the type.
	GetDocumentation(context.Context, *GetDocumentationRequest, ...dcerpc.CallOption) (*GetDocumentationResponse, error)

	// The GetDllEntry method retrieves values associated with a local-only method defined
	// in a module.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD does not match a value
	// in the following table, a fatal failure occurred.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD matches a value in
	// the following table, a failure occurred.
	//
	// Return value/code
	//
	// # Description
	//
	// 0x800288BD
	//
	// TYPE_E_BADMODULEKIND
	//
	// The type is not a module. See [MS-ERREF].
	//
	// 0x8002802C
	//
	// TYPE_E_AMBIGUOUSNAME
	//
	// The values of memid and invKind match more than one element in the binding context.
	// See [MS-ERREF].
	//
	// 0x8002802B
	//
	// TYPE_E_ELEMENTNOTFOUND
	//
	// The value of memid and invKind did not specify the ordinal position of an element
	// in the interface table, or the type is not a coclass ( 5583e1b8-454c-4147-9f56-f72416a15bee#gt_670b0ee2-d101-41b0-ac77-6ac7dbeee7dc
	// ). See [MS-ERREF].
	GetDLLEntry(context.Context, *GetDLLEntryRequest, ...dcerpc.CallOption) (*GetDLLEntryResponse, error)

	// The GetRefTypeInfo method retrieves an automation type description that describes
	// a type that is inherited, implemented, or referenced by the ITypeInfo server.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	GetReferenceTypeInfo(context.Context, *GetReferenceTypeInfoRequest, ...dcerpc.CallOption) (*GetReferenceTypeInfoResponse, error)

	// Opnum15NotUsedOnWire operation.
	// Opnum15NotUsedOnWire

	// The CreateInstance method creates a new instance of a type that describes a COM server
	// (coclass).
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD does not match a value
	// in the following table, a fatal failure occurred.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD matches a value in
	// the following table, a failure occurred.
	//
	// Return value/code
	//
	// # Description
	//
	// 0x80000004
	//
	// TYPE_E_NOINTERFACE
	//
	// The value of riid did not specify a known type. See [MS-ERREF].
	//
	// 0x800288BD
	//
	// TYPE_E_BADMODULEKIND
	//
	// The ITypeInfo server specified a non-coclass type. See [MS-ERREF].
	CreateInstance(context.Context, *CreateInstanceRequest, ...dcerpc.CallOption) (*CreateInstanceResponse, error)

	// The GetMops method has no effect when called across the wire.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	GetMops(context.Context, *GetMopsRequest, ...dcerpc.CallOption) (*GetMopsResponse, error)

	// The GetContainingTypeLib method retrieves the ITypeLib server instance whose type
	// information table contains the ITypeInfo instance, and the index of the ITypeInfo
	// instance within the type information table.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	GetContainingTypeLib(context.Context, *GetContainingTypeLibRequest, ...dcerpc.CallOption) (*GetContainingTypeLibResponse, error)

	// Opnum19NotUsedOnWire operation.
	// Opnum19NotUsedOnWire

	// Opnum20NotUsedOnWire operation.
	// Opnum20NotUsedOnWire

	// Opnum21NotUsedOnWire operation.
	// Opnum21NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) TypeInfoClient
}

type xxx_DefaultTypeInfoClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultTypeInfoClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultTypeInfoClient) GetTypeAttribute(ctx context.Context, in *GetTypeAttributeRequest, opts ...dcerpc.CallOption) (*GetTypeAttributeResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetTypeAttributeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfoClient) GetTypeComp(ctx context.Context, in *GetTypeCompRequest, opts ...dcerpc.CallOption) (*GetTypeCompResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetTypeCompResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfoClient) GetFuncDesc(ctx context.Context, in *GetFuncDescRequest, opts ...dcerpc.CallOption) (*GetFuncDescResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetFuncDescResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfoClient) GetVarDesc(ctx context.Context, in *GetVarDescRequest, opts ...dcerpc.CallOption) (*GetVarDescResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetVarDescResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfoClient) GetNames(ctx context.Context, in *GetNamesRequest, opts ...dcerpc.CallOption) (*GetNamesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNamesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfoClient) GetReferenceTypeOfImplType(ctx context.Context, in *GetReferenceTypeOfImplTypeRequest, opts ...dcerpc.CallOption) (*GetReferenceTypeOfImplTypeResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetReferenceTypeOfImplTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfoClient) GetImplTypeFlags(ctx context.Context, in *GetImplTypeFlagsRequest, opts ...dcerpc.CallOption) (*GetImplTypeFlagsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetImplTypeFlagsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfoClient) GetDocumentation(ctx context.Context, in *GetDocumentationRequest, opts ...dcerpc.CallOption) (*GetDocumentationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetDocumentationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfoClient) GetDLLEntry(ctx context.Context, in *GetDLLEntryRequest, opts ...dcerpc.CallOption) (*GetDLLEntryResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetDLLEntryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfoClient) GetReferenceTypeInfo(ctx context.Context, in *GetReferenceTypeInfoRequest, opts ...dcerpc.CallOption) (*GetReferenceTypeInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetReferenceTypeInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfoClient) CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...dcerpc.CallOption) (*CreateInstanceResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateInstanceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfoClient) GetMops(ctx context.Context, in *GetMopsRequest, opts ...dcerpc.CallOption) (*GetMopsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetMopsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfoClient) GetContainingTypeLib(ctx context.Context, in *GetContainingTypeLibRequest, opts ...dcerpc.CallOption) (*GetContainingTypeLibResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetContainingTypeLibResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfoClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultTypeInfoClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultTypeInfoClient) IPID(ctx context.Context, ipid *dcom.IPID) TypeInfoClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultTypeInfoClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewTypeInfoClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (TypeInfoClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(TypeInfoSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iunknown.NewUnknownClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultTypeInfoClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetTypeAttributeOperation structure represents the GetTypeAttr operation
type xxx_GetTypeAttributeOperation struct {
	This          *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat      `idl:"name:That" json:"that"`
	TypeAttribute *oaut.TypeAttribute `idl:"name:ppTypeAttr" json:"type_attribute"`
	_             uint32              `idl:"name:pReserved"`
	Return        int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTypeAttributeOperation) OpNum() int { return 3 }

func (o *xxx_GetTypeAttributeOperation) OpName() string { return "/ITypeInfo/v0/GetTypeAttr" }

func (o *xxx_GetTypeAttributeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeAttributeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeAttributeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeAttributeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeAttributeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppTypeAttr {out} (1:{pointer=ref}*(2))(2:{alias=LPTYPEATTR}*(1))(3:{alias=TYPEATTR}(struct))
	{
		if o.TypeAttribute != nil {
			_ptr_ppTypeAttr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.TypeAttribute != nil {
					if err := o.TypeAttribute.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.TypeAttribute{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TypeAttribute, _ptr_ppTypeAttr); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pReserved {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		// reserved pReserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeAttributeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppTypeAttr {out} (1:{pointer=ref}*(2))(2:{alias=LPTYPEATTR,pointer=ref}*(1))(3:{alias=TYPEATTR}(struct))
	{
		_ptr_ppTypeAttr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.TypeAttribute == nil {
				o.TypeAttribute = &oaut.TypeAttribute{}
			}
			if err := o.TypeAttribute.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppTypeAttr := func(ptr interface{}) { o.TypeAttribute = *ptr.(**oaut.TypeAttribute) }
		if err := w.ReadPointer(&o.TypeAttribute, _s_ppTypeAttr, _ptr_ppTypeAttr); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pReserved {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		// reserved pReserved
		var _pReserved uint32
		if err := w.ReadData(&_pReserved); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetTypeAttributeRequest structure represents the GetTypeAttr operation request
type GetTypeAttributeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetTypeAttributeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTypeAttributeOperation) *xxx_GetTypeAttributeOperation {
	if op == nil {
		op = &xxx_GetTypeAttributeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetTypeAttributeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTypeAttributeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTypeAttributeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetTypeAttributeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeAttributeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTypeAttributeResponse structure represents the GetTypeAttr operation response
type GetTypeAttributeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppTypeAttr: MUST be set to a TYPEATTR structure whose data values describe the type
	// associated with the ITypeInfo server, as specified in section 2.2.44.
	TypeAttribute *oaut.TypeAttribute `idl:"name:ppTypeAttr" json:"type_attribute"`
	// Return: The GetTypeAttr return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTypeAttributeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTypeAttributeOperation) *xxx_GetTypeAttributeOperation {
	if op == nil {
		op = &xxx_GetTypeAttributeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.TypeAttribute = o.TypeAttribute
	op.Return = o.Return
	return op
}

func (o *GetTypeAttributeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTypeAttributeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TypeAttribute = op.TypeAttribute
	o.Return = op.Return
}
func (o *GetTypeAttributeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetTypeAttributeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeAttributeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTypeCompOperation structure represents the GetTypeComp operation
type xxx_GetTypeCompOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Comp   *oaut.TypeComp `idl:"name:ppTComp" json:"comp"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTypeCompOperation) OpNum() int { return 4 }

func (o *xxx_GetTypeCompOperation) OpName() string { return "/ITypeInfo/v0/GetTypeComp" }

func (o *xxx_GetTypeCompOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeCompOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeCompOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeCompOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeCompOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppTComp {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ITypeComp}(interface))
	{
		if o.Comp != nil {
			_ptr_ppTComp := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Comp != nil {
					if err := o.Comp.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.TypeComp{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Comp, _ptr_ppTComp); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeCompOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppTComp {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ITypeComp}(interface))
	{
		_ptr_ppTComp := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Comp == nil {
				o.Comp = &oaut.TypeComp{}
			}
			if err := o.Comp.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppTComp := func(ptr interface{}) { o.Comp = *ptr.(**oaut.TypeComp) }
		if err := w.ReadPointer(&o.Comp, _s_ppTComp, _ptr_ppTComp); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetTypeCompRequest structure represents the GetTypeComp operation request
type GetTypeCompRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetTypeCompRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTypeCompOperation) *xxx_GetTypeCompOperation {
	if op == nil {
		op = &xxx_GetTypeCompOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetTypeCompRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTypeCompOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTypeCompRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetTypeCompRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeCompOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTypeCompResponse structure represents the GetTypeComp operation response
type GetTypeCompResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppTComp: MUST be set to a reference to the ITypeComp server instance associated with
	// the ITypeInfo server (see section 3.5).
	//
	// ppTComp: MUST be set to a reference to the ITypeComp server instance associated with
	// the automation type library, or to NULL if the automation type library does not have
	// an associated ITypeComp server instance.
	Comp *oaut.TypeComp `idl:"name:ppTComp" json:"comp"`
	// Return: The GetTypeComp return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTypeCompResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTypeCompOperation) *xxx_GetTypeCompOperation {
	if op == nil {
		op = &xxx_GetTypeCompOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Comp = o.Comp
	op.Return = o.Return
	return op
}

func (o *GetTypeCompResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTypeCompOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Comp = op.Comp
	o.Return = op.Return
}
func (o *GetTypeCompResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetTypeCompResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeCompOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFuncDescOperation structure represents the GetFuncDesc operation
type xxx_GetFuncDescOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Index    uint32         `idl:"name:index" json:"index"`
	FuncDesc *oaut.FuncDesc `idl:"name:ppFuncDesc" json:"func_desc"`
	_        uint32         `idl:"name:pReserved"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFuncDescOperation) OpNum() int { return 5 }

func (o *xxx_GetFuncDescOperation) OpName() string { return "/ITypeInfo/v0/GetFuncDesc" }

func (o *xxx_GetFuncDescOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFuncDescOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// index {in} (1:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFuncDescOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// index {in} (1:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFuncDescOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFuncDescOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppFuncDesc {out} (1:{pointer=ref}*(2))(2:{alias=LPFUNCDESC}*(1))(3:{alias=FUNCDESC}(struct))
	{
		if o.FuncDesc != nil {
			_ptr_ppFuncDesc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FuncDesc != nil {
					if err := o.FuncDesc.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.FuncDesc{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FuncDesc, _ptr_ppFuncDesc); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pReserved {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		// reserved pReserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFuncDescOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppFuncDesc {out} (1:{pointer=ref}*(2))(2:{alias=LPFUNCDESC,pointer=ref}*(1))(3:{alias=FUNCDESC}(struct))
	{
		_ptr_ppFuncDesc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FuncDesc == nil {
				o.FuncDesc = &oaut.FuncDesc{}
			}
			if err := o.FuncDesc.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppFuncDesc := func(ptr interface{}) { o.FuncDesc = *ptr.(**oaut.FuncDesc) }
		if err := w.ReadPointer(&o.FuncDesc, _s_ppFuncDesc, _ptr_ppFuncDesc); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pReserved {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		// reserved pReserved
		var _pReserved uint32
		if err := w.ReadData(&_pReserved); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetFuncDescRequest structure represents the GetFuncDesc operation request
type GetFuncDescRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// index: MUST equal the ordinal position in the method table (if the type describes
	// a DCOM interface or module) or the dispatch method table (if the type describes a
	// dispinterface) of the method whose description is to be returned. The value of index
	// MUST be less than the value of the cFuncs field in the TYPEATTR (section 2.2.44)
	// structure returned by the GetTypeAttr (section 3.7.4.1) method.
	Index uint32 `idl:"name:index" json:"index"`
}

func (o *GetFuncDescRequest) xxx_ToOp(ctx context.Context, op *xxx_GetFuncDescOperation) *xxx_GetFuncDescOperation {
	if op == nil {
		op = &xxx_GetFuncDescOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Index = o.Index
	return op
}

func (o *GetFuncDescRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFuncDescOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
}
func (o *GetFuncDescRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetFuncDescRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFuncDescOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFuncDescResponse structure represents the GetFuncDesc operation response
type GetFuncDescResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppFuncDesc: MUST be set to a FUNCDESC structure (see section 2.2.42) that contains
	// the data values associated with the specified member of the method or dispatch method
	// table, or NULL if no such member exists.
	FuncDesc *oaut.FuncDesc `idl:"name:ppFuncDesc" json:"func_desc"`
	// Return: The GetFuncDesc return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFuncDescResponse) xxx_ToOp(ctx context.Context, op *xxx_GetFuncDescOperation) *xxx_GetFuncDescOperation {
	if op == nil {
		op = &xxx_GetFuncDescOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.FuncDesc = o.FuncDesc
	op.Return = o.Return
	return op
}

func (o *GetFuncDescResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFuncDescOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FuncDesc = op.FuncDesc
	o.Return = op.Return
}
func (o *GetFuncDescResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetFuncDescResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFuncDescOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetVarDescOperation structure represents the GetVarDesc operation
type xxx_GetVarDescOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Index   uint32         `idl:"name:index" json:"index"`
	VarDesc *oaut.VarDesc  `idl:"name:ppVarDesc" json:"var_desc"`
	_       uint32         `idl:"name:pReserved"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetVarDescOperation) OpNum() int { return 6 }

func (o *xxx_GetVarDescOperation) OpName() string { return "/ITypeInfo/v0/GetVarDesc" }

func (o *xxx_GetVarDescOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVarDescOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// index {in} (1:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVarDescOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// index {in} (1:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVarDescOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVarDescOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppVarDesc {out} (1:{pointer=ref}*(2))(2:{alias=LPVARDESC}*(1))(3:{alias=VARDESC}(struct))
	{
		if o.VarDesc != nil {
			_ptr_ppVarDesc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.VarDesc != nil {
					if err := o.VarDesc.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.VarDesc{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.VarDesc, _ptr_ppVarDesc); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pReserved {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		// reserved pReserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVarDescOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppVarDesc {out} (1:{pointer=ref}*(2))(2:{alias=LPVARDESC,pointer=ref}*(1))(3:{alias=VARDESC}(struct))
	{
		_ptr_ppVarDesc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.VarDesc == nil {
				o.VarDesc = &oaut.VarDesc{}
			}
			if err := o.VarDesc.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppVarDesc := func(ptr interface{}) { o.VarDesc = *ptr.(**oaut.VarDesc) }
		if err := w.ReadPointer(&o.VarDesc, _s_ppVarDesc, _ptr_ppVarDesc); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pReserved {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		// reserved pReserved
		var _pReserved uint32
		if err := w.ReadData(&_pReserved); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetVarDescRequest structure represents the GetVarDesc operation request
type GetVarDescRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// index: MUST equal the data member table index value of the data member whose description
	// is to be returned. The value of index MUST be less than the value of the cVars field
	// in the TYPEATTR structure returned by the GetTypeAttr method, as specified in 3.7.4.1
	// and 2.2.44.
	Index uint32 `idl:"name:index" json:"index"`
}

func (o *GetVarDescRequest) xxx_ToOp(ctx context.Context, op *xxx_GetVarDescOperation) *xxx_GetVarDescOperation {
	if op == nil {
		op = &xxx_GetVarDescOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Index = o.Index
	return op
}

func (o *GetVarDescRequest) xxx_FromOp(ctx context.Context, op *xxx_GetVarDescOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
}
func (o *GetVarDescRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetVarDescRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVarDescOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetVarDescResponse structure represents the GetVarDesc operation response
type GetVarDescResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppVarDesc: MUST be set to a VARDESC structure (see section 2.2.43) that contains
	// the data values associated with the specified member of the data member table, or
	// NULL if no such member exists.
	VarDesc *oaut.VarDesc `idl:"name:ppVarDesc" json:"var_desc"`
	// Return: The GetVarDesc return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetVarDescResponse) xxx_ToOp(ctx context.Context, op *xxx_GetVarDescOperation) *xxx_GetVarDescOperation {
	if op == nil {
		op = &xxx_GetVarDescOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.VarDesc = o.VarDesc
	op.Return = o.Return
	return op
}

func (o *GetVarDescResponse) xxx_FromOp(ctx context.Context, op *xxx_GetVarDescOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.VarDesc = op.VarDesc
	o.Return = op.Return
}
func (o *GetVarDescResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetVarDescResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVarDescOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNamesOperation structure represents the GetNames operation
type xxx_GetNamesOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	MemberID      int32          `idl:"name:memid" json:"member_id"`
	Names         []*oaut.String `idl:"name:rgBstrNames;size_is:(cMaxNames);length_is:(pcNames)" json:"names"`
	MaxNamesCount uint32         `idl:"name:cMaxNames" json:"max_names_count"`
	NamesCount    uint32         `idl:"name:pcNames" json:"names_count"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNamesOperation) OpNum() int { return 7 }

func (o *xxx_GetNamesOperation) OpName() string { return "/ITypeInfo/v0/GetNames" }

func (o *xxx_GetNamesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNamesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// memid {in} (1:{alias=MEMBERID, names=LONG}(int32))
	{
		if err := w.WriteData(o.MemberID); err != nil {
			return err
		}
	}
	// cMaxNames {in} (1:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.MaxNamesCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNamesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// memid {in} (1:{alias=MEMBERID, names=LONG}(int32))
	{
		if err := w.ReadData(&o.MemberID); err != nil {
			return err
		}
	}
	// cMaxNames {in} (1:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.MaxNamesCount); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNamesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Names != nil && o.NamesCount == 0 {
		o.NamesCount = uint32(len(o.Names))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNamesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// rgBstrNames {out} (1:{pointer=ref}*(1))(2:{pointer=unique, alias=BSTR}[dim:0,size_is=cMaxNames,length_is=pcNames]*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		dimSize1 := uint64(o.MaxNamesCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.NamesCount)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.Names {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.Names[i1] != nil {
				_ptr_rgBstrNames := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if o.Names[i1] != nil {
						if err := o.Names[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
					return nil
				})
				if err := w.WritePointer(&o.Names[i1], _ptr_rgBstrNames); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.Names); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pcNames {out} (1:{pointer=ref}*(1))(2:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.NamesCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNamesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// rgBstrNames {out} (1:{pointer=ref}*(1))(2:{pointer=unique, alias=BSTR}[dim:0,size_is=cMaxNames,length_is=pcNames]*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Names", sizeInfo[0])
		}
		o.Names = make([]*oaut.String, sizeInfo[0])
		for i1 := range o.Names {
			i1 := i1
			_ptr_rgBstrNames := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.Names[i1] == nil {
					o.Names[i1] = &oaut.String{}
				}
				if err := o.Names[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_rgBstrNames := func(ptr interface{}) { o.Names[i1] = *ptr.(**oaut.String) }
			if err := w.ReadPointer(&o.Names[i1], _s_rgBstrNames, _ptr_rgBstrNames); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcNames {out} (1:{pointer=ref}*(1))(2:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.NamesCount); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNamesRequest structure represents the GetNames operation request
type GetNamesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// memid: MUST be a MEMBERID (section 2.2.35).
	MemberID int32 `idl:"name:memid" json:"member_id"`
	// cMaxNames: MUST specify the maximum length of the rgBstrNames array that the client
	// can accept.
	MaxNamesCount uint32 `idl:"name:cMaxNames" json:"max_names_count"`
}

func (o *GetNamesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNamesOperation) *xxx_GetNamesOperation {
	if op == nil {
		op = &xxx_GetNamesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MemberID = o.MemberID
	op.MaxNamesCount = o.MaxNamesCount
	return op
}

func (o *GetNamesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNamesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MemberID = op.MemberID
	o.MaxNamesCount = op.MaxNamesCount
}
func (o *GetNamesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNamesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNamesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNamesResponse structure represents the GetNames operation response
type GetNamesResponse struct {
	// XXX: cMaxNames is an implicit input depedency for output parameters
	MaxNamesCount uint32 `idl:"name:cMaxNames" json:"max_names_count"`

	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// rgBstrNames: MUST be set to an array of BSTR. If pcNames is 0, rgBstrNames MUST be
	// NULL.
	Names []*oaut.String `idl:"name:rgBstrNames;size_is:(cMaxNames);length_is:(pcNames)" json:"names"`
	// pcNames: MUST be set to the length of the rgBstrNames array.
	NamesCount uint32 `idl:"name:pcNames" json:"names_count"`
	// Return: The GetNames return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNamesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNamesOperation) *xxx_GetNamesOperation {
	if op == nil {
		op = &xxx_GetNamesOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.MaxNamesCount == uint32(0) {
		op.MaxNamesCount = o.MaxNamesCount
	}

	op.That = o.That
	op.Names = o.Names
	op.NamesCount = o.NamesCount
	op.Return = o.Return
	return op
}

func (o *GetNamesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNamesOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.MaxNamesCount = op.MaxNamesCount

	o.That = op.That
	o.Names = op.Names
	o.NamesCount = op.NamesCount
	o.Return = op.Return
}
func (o *GetNamesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNamesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNamesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetReferenceTypeOfImplTypeOperation structure represents the GetRefTypeOfImplType operation
type xxx_GetReferenceTypeOfImplTypeOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	Index         uint32         `idl:"name:index" json:"index"`
	ReferenceType uint32         `idl:"name:pRefType" json:"reference_type"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetReferenceTypeOfImplTypeOperation) OpNum() int { return 8 }

func (o *xxx_GetReferenceTypeOfImplTypeOperation) OpName() string {
	return "/ITypeInfo/v0/GetRefTypeOfImplType"
}

func (o *xxx_GetReferenceTypeOfImplTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReferenceTypeOfImplTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// index {in} (1:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReferenceTypeOfImplTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// index {in} (1:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReferenceTypeOfImplTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReferenceTypeOfImplTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pRefType {out} (1:{pointer=ref}*(1))(2:{alias=HREFTYPE, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.ReferenceType); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReferenceTypeOfImplTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pRefType {out} (1:{pointer=ref}*(1))(2:{alias=HREFTYPE, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ReferenceType); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetReferenceTypeOfImplTypeRequest structure represents the GetRefTypeOfImplType operation request
type GetReferenceTypeOfImplTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// index: MUST be a nonnegative integer, or -1.
	//
	// If the ITypeInfo server describes a dual interface (see sections 2.2.49.4.2 and 3.7.1),
	// the value of index MUST be 0 or -1.
	//
	// If the ITypeInfo server describes a coclass, the value of index MUST be nonnegative
	// and less than the value of the cImplTypes field in the TYPEATTR (section 2.2.44)
	// structure returned by the GetTypeAttr (section 3.7.4.1) method.
	Index uint32 `idl:"name:index" json:"index"`
}

func (o *GetReferenceTypeOfImplTypeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetReferenceTypeOfImplTypeOperation) *xxx_GetReferenceTypeOfImplTypeOperation {
	if op == nil {
		op = &xxx_GetReferenceTypeOfImplTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Index = o.Index
	return op
}

func (o *GetReferenceTypeOfImplTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetReferenceTypeOfImplTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
}
func (o *GetReferenceTypeOfImplTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetReferenceTypeOfImplTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReferenceTypeOfImplTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetReferenceTypeOfImplTypeResponse structure represents the GetRefTypeOfImplType operation response
type GetReferenceTypeOfImplTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pRefType: MUST be set to one of the following values, if index is -1 or specifies
	// an interface table entry.
	//
	// If the ITypeInfo server describes a dual interface and index is -1, pRefType is specified
	// by the TYPEKIND value associated with the ITypeInfo server (see section 2.2.44):
	//
	// TKIND_DISPATCH: MUST be set to the HREFTYPE of the partner interface.
	//
	// TKIND_INTERFACE: MUST be set to the HREFTYPE of the partner dispinterface.
	//
	// In all other cases, pRefType is specified by the interface table member whose ordinal
	// position is specified by index:
	//
	// If the interface table member is a dual interface and the ITypeInfo server describes
	// a DCOM interface or partner interface, pRefType MUST be the HREFTYPE of the partner
	// interface of the interface table member.
	//
	// Note  This is the only case where an OLE Automation Protocol interface method returns
	// a partner interface by default.
	//
	// If the interface table member is a dual interface and the ITypeInfo server describes
	// a coclass, pRefType MUST be the HREFTYPE of the partner dispinterface of the interface
	// table member.
	ReferenceType uint32 `idl:"name:pRefType" json:"reference_type"`
	// Return: The GetRefTypeOfImplType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetReferenceTypeOfImplTypeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetReferenceTypeOfImplTypeOperation) *xxx_GetReferenceTypeOfImplTypeOperation {
	if op == nil {
		op = &xxx_GetReferenceTypeOfImplTypeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReferenceType = o.ReferenceType
	op.Return = o.Return
	return op
}

func (o *GetReferenceTypeOfImplTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetReferenceTypeOfImplTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReferenceType = op.ReferenceType
	o.Return = op.Return
}
func (o *GetReferenceTypeOfImplTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetReferenceTypeOfImplTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReferenceTypeOfImplTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetImplTypeFlagsOperation structure represents the GetImplTypeFlags operation
type xxx_GetImplTypeFlagsOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	Index         uint32         `idl:"name:index" json:"index"`
	ImplTypeFlags int32          `idl:"name:pImplTypeFlags" json:"impl_type_flags"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetImplTypeFlagsOperation) OpNum() int { return 9 }

func (o *xxx_GetImplTypeFlagsOperation) OpName() string { return "/ITypeInfo/v0/GetImplTypeFlags" }

func (o *xxx_GetImplTypeFlagsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetImplTypeFlagsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// index {in} (1:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetImplTypeFlagsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// index {in} (1:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetImplTypeFlagsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetImplTypeFlagsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pImplTypeFlags {out} (1:{pointer=ref}*(1))(2:{alias=INT}(int32))
	{
		if err := w.WriteData(o.ImplTypeFlags); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetImplTypeFlagsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pImplTypeFlags {out} (1:{pointer=ref}*(1))(2:{alias=INT}(int32))
	{
		if err := w.ReadData(&o.ImplTypeFlags); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetImplTypeFlagsRequest structure represents the GetImplTypeFlags operation request
type GetImplTypeFlagsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// index: MUST be the ordinal position in the coclass interface table of the interface
	// whose associated IMPLTYPEFLAGS values are to be returned.
	Index uint32 `idl:"name:index" json:"index"`
}

func (o *GetImplTypeFlagsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetImplTypeFlagsOperation) *xxx_GetImplTypeFlagsOperation {
	if op == nil {
		op = &xxx_GetImplTypeFlagsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Index = o.Index
	return op
}

func (o *GetImplTypeFlagsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetImplTypeFlagsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
}
func (o *GetImplTypeFlagsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetImplTypeFlagsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetImplTypeFlagsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetImplTypeFlagsResponse structure represents the GetImplTypeFlags operation response
type GetImplTypeFlagsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pImplTypeFlags: MUST be set to either a combination of the IMPLTYPEFLAGS feature
	// constants specified in section 2.2.13, or 0.
	ImplTypeFlags int32 `idl:"name:pImplTypeFlags" json:"impl_type_flags"`
	// Return: The GetImplTypeFlags return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetImplTypeFlagsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetImplTypeFlagsOperation) *xxx_GetImplTypeFlagsOperation {
	if op == nil {
		op = &xxx_GetImplTypeFlagsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ImplTypeFlags = o.ImplTypeFlags
	op.Return = o.Return
	return op
}

func (o *GetImplTypeFlagsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetImplTypeFlagsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ImplTypeFlags = op.ImplTypeFlags
	o.Return = op.Return
}
func (o *GetImplTypeFlagsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetImplTypeFlagsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetImplTypeFlagsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDocumentationOperation structure represents the GetDocumentation operation
type xxx_GetDocumentationOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	MemberID     int32          `idl:"name:memid" json:"member_id"`
	PointerFlags uint32         `idl:"name:refPtrFlags" json:"pointer_flags"`
	Name         *oaut.String   `idl:"name:pBstrName" json:"name"`
	DocString    *oaut.String   `idl:"name:pBstrDocString" json:"doc_string"`
	HelpContext  uint32         `idl:"name:pdwHelpContext" json:"help_context"`
	HelpFile     *oaut.String   `idl:"name:pBstrHelpFile" json:"help_file"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDocumentationOperation) OpNum() int { return 12 }

func (o *xxx_GetDocumentationOperation) OpName() string { return "/ITypeInfo/v0/GetDocumentation" }

func (o *xxx_GetDocumentationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDocumentationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// memid {in} (1:{alias=MEMBERID, names=LONG}(int32))
	{
		if err := w.WriteData(o.MemberID); err != nil {
			return err
		}
	}
	// refPtrFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PointerFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDocumentationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// memid {in} (1:{alias=MEMBERID, names=LONG}(int32))
	{
		if err := w.ReadData(&o.MemberID); err != nil {
			return err
		}
	}
	// refPtrFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PointerFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDocumentationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDocumentationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pBstrName {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Name != nil {
			_ptr_pBstrName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Name != nil {
					if err := o.Name.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_pBstrName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pBstrDocString {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DocString != nil {
			_ptr_pBstrDocString := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DocString != nil {
					if err := o.DocString.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DocString, _ptr_pBstrDocString); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pdwHelpContext {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.HelpContext); err != nil {
			return err
		}
	}
	// pBstrHelpFile {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.HelpFile != nil {
			_ptr_pBstrHelpFile := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.HelpFile != nil {
					if err := o.HelpFile.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.HelpFile, _ptr_pBstrHelpFile); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDocumentationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pBstrName {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pBstrName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Name == nil {
				o.Name = &oaut.String{}
			}
			if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pBstrName := func(ptr interface{}) { o.Name = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Name, _s_pBstrName, _ptr_pBstrName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pBstrDocString {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pBstrDocString := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DocString == nil {
				o.DocString = &oaut.String{}
			}
			if err := o.DocString.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pBstrDocString := func(ptr interface{}) { o.DocString = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DocString, _s_pBstrDocString, _ptr_pBstrDocString); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwHelpContext {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.HelpContext); err != nil {
			return err
		}
	}
	// pBstrHelpFile {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pBstrHelpFile := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.HelpFile == nil {
				o.HelpFile = &oaut.String{}
			}
			if err := o.HelpFile.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pBstrHelpFile := func(ptr interface{}) { o.HelpFile = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.HelpFile, _s_pBstrHelpFile, _ptr_pBstrHelpFile); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetDocumentationRequest structure represents the GetDocumentation operation request
type GetDocumentationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// memid: MUST be either the MEMBERID of a method or data member in the binding context
	// of the ITypeInfo server (see section 3.5.4.1.1), or MEMBERID_NIL (see section 2.2.35.1).
	MemberID int32 `idl:"name:memid" json:"member_id"`
	// refPtrFlags: MUST be a combination of the bit flags specified in the following table,
	// or 0.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|                                    |                                                                                  |
	//	|               VALUE                |                                     MEANING                                      |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| TYPEINFO_NameArg 0x00000001        | MUST specify that the client is interested in the actual pBstrName [out]         |
	//	|                                    | argument.                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| TYPEINFO_DocStringArg 0x00000002   | MUST specify that the client is interested in the actual pBstrDocString [out]    |
	//	|                                    | argument.                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| TYPEINFO_HelpContextArg 0x00000004 | MUST specify that the client is interested in the actual pdwHelpContext [out]    |
	//	|                                    | argument.                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| TYPEINFO_HelpFileArg 0x00000008    | MUST specify that the client is interested in the actual pBstrHelpFile [out]     |
	//	|                                    | argument.                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// refPtrFlags: MUST be a combination of 0, or the bit flags specified in the following
	// table.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|                                   |                                                                                  |
	//	|               VALUE               |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| TYPELIB_NameArg 0x00000001        | MUST specify that the client is interested in the actual pBstrName [out]         |
	//	|                                   | argument.                                                                        |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| TYPELIB_DocStringArg 0x00000002   | MUST specify that the client is interested in the actual pBstrDocString [out]    |
	//	|                                   | argument.                                                                        |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| TYPELIB_HelpContextArg 0x00000004 | MUST specify that the client is interested in the actual pdwHelpContext [out]    |
	//	|                                   | argument.                                                                        |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| TYPELIB_HelpFileArg 0x00000008    | MUST specify that the client is interested in the actual pBstrHelpFile [out]     |
	//	|                                   | argument.                                                                        |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	PointerFlags uint32 `idl:"name:refPtrFlags" json:"pointer_flags"`
}

func (o *GetDocumentationRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDocumentationOperation) *xxx_GetDocumentationOperation {
	if op == nil {
		op = &xxx_GetDocumentationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MemberID = o.MemberID
	op.PointerFlags = o.PointerFlags
	return op
}

func (o *GetDocumentationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDocumentationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MemberID = op.MemberID
	o.PointerFlags = op.PointerFlags
}
func (o *GetDocumentationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDocumentationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDocumentationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDocumentationResponse structure represents the GetDocumentation operation response
type GetDocumentationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pBstrName: If the TYPEINFO_NameArg bit flag is set in refPtrFlags,pBstrName MUST
	// be set to a BSTR that contains the name of the type or specified type member. Othernwise,
	// pBstrName MUST be set to a NULL BSTR.
	//
	// pBstrName: MUST be set to a BSTR that contains the name of the specified type or
	// Type library if the TYPELIB_NameArg bit flag is set in refPtrFlags. MUST be set to
	// a NULL BSTR otherwise.
	Name *oaut.String `idl:"name:pBstrName" json:"name"`
	// pBstrDocString: MUST be set to the documentation string that was associated with
	// the type or specified type member using the [helpstring] attribute (see section 2.2.49.2),
	// if the TYPEINFO_DocStringArg bit flag is set in refPtrFlags. MAY be set to an implementation-specific
	// string<59> if no [helpstring] attribute is specified. MUST be set to a NULL BSTR
	// otherwise.
	//
	// pBstrDocString: MUST be set to the documentation string that was associated with
	// the specified type or Type library using the [helpstring] attribute (see section
	// 2.2.49.2), if the TYPELIB_DocStringArg bit flag is set in refPtrFlags. MAY be set
	// to an implementation-specific string<61> if no [helpstring] attribute is specified.
	// MUST be set to a NULL BSTR otherwise.
	DocString *oaut.String `idl:"name:pBstrDocString" json:"doc_string"`
	// pdwHelpContext: MUST be set to the value that was associated with the type or specified
	// type member using the [helpcontext] attribute (see section 2.2.49.2), if the TYPEINFO_HelpContextArg
	// bit flag is set in refPtrFlags. MUST be set to 0 otherwise.
	//
	// pdwHelpContext: MUST be set to the value that was associated with the specified type
	// or Type library using the [helpcontext] attribute (see section 2.2.49.2), if the
	// TYPELIB_HelpContextArg bit flag is set in refPtrFlags. MUST be set to 0 otherwise.
	HelpContext uint32 `idl:"name:pdwHelpContext" json:"help_context"`
	// pBstrHelpFile: MUST be set to the documentation string that was associated with the
	// type or specified type member using the [helpfile] attribute (see section 2.2.49.2),
	// if the TYPEINFO_HelpFileArg bit flag is set in refPtrFlags. MUST be set to a NULL
	// BSTR otherwise.
	//
	// pBstrHelpFile: MUST be set to the documentation string that was associated with the
	// specified type or Type library using the [helpfile] attribute (see section 2.2.49.2),
	// if the TYPELIB_HelpFileArg bit flag is set in refPtrFlags. MUST be set to a NULL
	// BSTR otherwise.
	HelpFile *oaut.String `idl:"name:pBstrHelpFile" json:"help_file"`
	// Return: The GetDocumentation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDocumentationResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDocumentationOperation) *xxx_GetDocumentationOperation {
	if op == nil {
		op = &xxx_GetDocumentationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Name = o.Name
	op.DocString = o.DocString
	op.HelpContext = o.HelpContext
	op.HelpFile = o.HelpFile
	op.Return = o.Return
	return op
}

func (o *GetDocumentationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDocumentationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Name = op.Name
	o.DocString = op.DocString
	o.HelpContext = op.HelpContext
	o.HelpFile = op.HelpFile
	o.Return = op.Return
}
func (o *GetDocumentationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDocumentationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDocumentationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDLLEntryOperation structure represents the GetDllEntry operation
type xxx_GetDLLEntryOperation struct {
	This         *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat  `idl:"name:That" json:"that"`
	MemberID     int32           `idl:"name:memid" json:"member_id"`
	InvKind      oaut.InvokeKind `idl:"name:invKind" json:"inv_kind"`
	PointerFlags uint32          `idl:"name:refPtrFlags" json:"pointer_flags"`
	DLLName      *oaut.String    `idl:"name:pBstrDllName" json:"dll_name"`
	Name         *oaut.String    `idl:"name:pBstrName" json:"name"`
	Ordinal      uint16          `idl:"name:pwOrdinal" json:"ordinal"`
	Return       int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDLLEntryOperation) OpNum() int { return 13 }

func (o *xxx_GetDLLEntryOperation) OpName() string { return "/ITypeInfo/v0/GetDllEntry" }

func (o *xxx_GetDLLEntryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDLLEntryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// memid {in} (1:{alias=MEMBERID, names=LONG}(int32))
	{
		if err := w.WriteData(o.MemberID); err != nil {
			return err
		}
	}
	// invKind {in} (1:{v1_enum, alias=INVOKEKIND}(enum))
	{
		if err := w.WriteEnum(uint32(o.InvKind)); err != nil {
			return err
		}
	}
	// refPtrFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PointerFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDLLEntryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// memid {in} (1:{alias=MEMBERID, names=LONG}(int32))
	{
		if err := w.ReadData(&o.MemberID); err != nil {
			return err
		}
	}
	// invKind {in} (1:{v1_enum, alias=INVOKEKIND}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.InvKind)); err != nil {
			return err
		}
	}
	// refPtrFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PointerFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDLLEntryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDLLEntryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pBstrDllName {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.DLLName != nil {
			_ptr_pBstrDllName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DLLName != nil {
					if err := o.DLLName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DLLName, _ptr_pBstrDllName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pBstrName {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Name != nil {
			_ptr_pBstrName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Name != nil {
					if err := o.Name.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_pBstrName); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pwOrdinal {out} (1:{pointer=ref}*(1))(2:{alias=WORD}(uint16))
	{
		if err := w.WriteData(o.Ordinal); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDLLEntryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pBstrDllName {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pBstrDllName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DLLName == nil {
				o.DLLName = &oaut.String{}
			}
			if err := o.DLLName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pBstrDllName := func(ptr interface{}) { o.DLLName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.DLLName, _s_pBstrDllName, _ptr_pBstrDllName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pBstrName {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pBstrName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Name == nil {
				o.Name = &oaut.String{}
			}
			if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pBstrName := func(ptr interface{}) { o.Name = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Name, _s_pBstrName, _ptr_pBstrName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pwOrdinal {out} (1:{pointer=ref}*(1))(2:{alias=WORD}(uint16))
	{
		if err := w.ReadData(&o.Ordinal); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetDLLEntryRequest structure represents the GetDllEntry operation request
type GetDLLEntryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// memid: MUST be the MEMBERID of a method member of the module defined in the automation
	// type library.
	MemberID int32 `idl:"name:memid" json:"member_id"`
	// invKind: MUST be a value of the INVOKEKIND (section 2.2.14) enumeration that specifies
	// a single property accessor method, if memid corresponds to a property with multiple
	// accessors.
	InvKind oaut.InvokeKind `idl:"name:invKind" json:"inv_kind"`
	// refPtrFlags: MUST be a combination of the bit flags specified in the following table,
	// or 0.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|                                |                                                                                  |
	//	|             VALUE              |                                     MEANING                                      |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| TYPEINFO_DLLNameArg 0x00000001 | MUST specify that the client is interested in the actual pBstrDllName [out]      |
	//	|                                | argument.                                                                        |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| TYPEINFO_NameArg 0x00000002    | MUST specify that the client is interested in the actual pBstrName [out]         |
	//	|                                | argument.                                                                        |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| TYPEINFO_OrdinalArg 0x00000004 | MUST specify that the client is interested in the actual pwOrdinal [out]         |
	//	|                                | argument.                                                                        |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	PointerFlags uint32 `idl:"name:refPtrFlags" json:"pointer_flags"`
}

func (o *GetDLLEntryRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDLLEntryOperation) *xxx_GetDLLEntryOperation {
	if op == nil {
		op = &xxx_GetDLLEntryOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MemberID = o.MemberID
	op.InvKind = o.InvKind
	op.PointerFlags = o.PointerFlags
	return op
}

func (o *GetDLLEntryRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDLLEntryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MemberID = op.MemberID
	o.InvKind = op.InvKind
	o.PointerFlags = op.PointerFlags
}
func (o *GetDLLEntryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDLLEntryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDLLEntryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDLLEntryResponse structure represents the GetDllEntry operation response
type GetDLLEntryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pBstrDllName: MUST be set to the value associated with the method using the [dllname]
	// attribute (see section 2.2.49.9) if the TYPEINFO_DllNameArg bit flag is set in refPtrFlags.
	// MUST be set to a NULL BSTR otherwise.
	DLLName *oaut.String `idl:"name:pBstrDllName" json:"dll_name"`
	// pBstrName: MUST be set to the value associated with the method using the [entry]
	// attribute (see section 2.2.49.9), if the associated value is a string and the TYPEINFO_NameArg
	// bit flag is set in refPtrFlags. MUST be set to a NULL BSTR otherwise.
	Name *oaut.String `idl:"name:pBstrName" json:"name"`
	// pwOrdinal: MUST be set to the value associated with the method using the [entry]
	// attribute (see section 2.2.49.9), if the associated value is an integer and the TYPEINFO_OrdinalArg
	// bit flag is set in refPtrFlags. MUST be set to 0 otherwise.
	Ordinal uint16 `idl:"name:pwOrdinal" json:"ordinal"`
	// Return: The GetDllEntry return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDLLEntryResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDLLEntryOperation) *xxx_GetDLLEntryOperation {
	if op == nil {
		op = &xxx_GetDLLEntryOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.DLLName = o.DLLName
	op.Name = o.Name
	op.Ordinal = o.Ordinal
	op.Return = o.Return
	return op
}

func (o *GetDLLEntryResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDLLEntryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DLLName = op.DLLName
	o.Name = op.Name
	o.Ordinal = op.Ordinal
	o.Return = op.Return
}
func (o *GetDLLEntryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDLLEntryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDLLEntryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetReferenceTypeInfoOperation structure represents the GetRefTypeInfo operation
type xxx_GetReferenceTypeInfoOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	TypeHandle uint32         `idl:"name:hRefType" json:"type_handle"`
	TypeInfo   *oaut.TypeInfo `idl:"name:ppTInfo" json:"type_info"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetReferenceTypeInfoOperation) OpNum() int { return 14 }

func (o *xxx_GetReferenceTypeInfoOperation) OpName() string { return "/ITypeInfo/v0/GetRefTypeInfo" }

func (o *xxx_GetReferenceTypeInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReferenceTypeInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// hRefType {in} (1:{alias=HREFTYPE, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.TypeHandle); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReferenceTypeInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// hRefType {in} (1:{alias=HREFTYPE, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TypeHandle); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReferenceTypeInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReferenceTypeInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppTInfo {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ITypeInfo}(interface))
	{
		if o.TypeInfo != nil {
			_ptr_ppTInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.TypeInfo != nil {
					if err := o.TypeInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.TypeInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TypeInfo, _ptr_ppTInfo); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReferenceTypeInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppTInfo {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ITypeInfo}(interface))
	{
		_ptr_ppTInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.TypeInfo == nil {
				o.TypeInfo = &oaut.TypeInfo{}
			}
			if err := o.TypeInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppTInfo := func(ptr interface{}) { o.TypeInfo = *ptr.(**oaut.TypeInfo) }
		if err := w.ReadPointer(&o.TypeInfo, _s_ppTInfo, _ptr_ppTInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetReferenceTypeInfoRequest structure represents the GetRefTypeInfo operation request
type GetReferenceTypeInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// hRefType: MUST be an HREFTYPE (section 2.2.36) that has been provided by the ITypeInfo
	// server instance whose GetRefTypeInfo method is being called.
	TypeHandle uint32 `idl:"name:hRefType" json:"type_handle"`
}

func (o *GetReferenceTypeInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetReferenceTypeInfoOperation) *xxx_GetReferenceTypeInfoOperation {
	if op == nil {
		op = &xxx_GetReferenceTypeInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.TypeHandle = o.TypeHandle
	return op
}

func (o *GetReferenceTypeInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetReferenceTypeInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.TypeHandle = op.TypeHandle
}
func (o *GetReferenceTypeInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetReferenceTypeInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReferenceTypeInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetReferenceTypeInfoResponse structure represents the GetRefTypeInfo operation response
type GetReferenceTypeInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppTInfo: MUST be set to a reference to an ITypeInfo server instance that provides
	// an automation type description of the inherited or referenced interface, or NULL
	// if hRefType does not specify an available interface.
	TypeInfo *oaut.TypeInfo `idl:"name:ppTInfo" json:"type_info"`
	// Return: The GetRefTypeInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetReferenceTypeInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetReferenceTypeInfoOperation) *xxx_GetReferenceTypeInfoOperation {
	if op == nil {
		op = &xxx_GetReferenceTypeInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.TypeInfo = o.TypeInfo
	op.Return = o.Return
	return op
}

func (o *GetReferenceTypeInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetReferenceTypeInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TypeInfo = op.TypeInfo
	o.Return = op.Return
}
func (o *GetReferenceTypeInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetReferenceTypeInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReferenceTypeInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateInstanceOperation structure represents the CreateInstance operation
type xxx_CreateInstanceOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	IID    *dcom.IID      `idl:"name:riid" json:"iid"`
	Object *dcom.Unknown  `idl:"name:ppvObj" json:"object"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateInstanceOperation) OpNum() int { return 16 }

func (o *xxx_CreateInstanceOperation) OpName() string { return "/ITypeInfo/v0/CreateInstance" }

func (o *xxx_CreateInstanceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateInstanceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// riid {in} (1:{alias=REFIID}*(1))(2:{alias=IID, names=GUID}(struct))
	{
		if o.IID != nil {
			if err := o.IID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.IID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CreateInstanceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// riid {in} (1:{alias=REFIID,pointer=ref}*(1))(2:{alias=IID, names=GUID}(struct))
	{
		if o.IID == nil {
			o.IID = &dcom.IID{}
		}
		if err := o.IID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateInstanceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateInstanceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppvObj {out, iid_is={riid}} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		if o.Object != nil {
			_ptr_ppvObj := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Object != nil {
					if err := o.Object.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Object, _ptr_ppvObj); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateInstanceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppvObj {out, iid_is={riid}} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_ppvObj := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Object == nil {
				o.Object = &dcom.Unknown{}
			}
			if err := o.Object.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppvObj := func(ptr interface{}) { o.Object = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.Object, _s_ppvObj, _ptr_ppvObj); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateInstanceRequest structure represents the CreateInstance operation request
type CreateInstanceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// riid: MUST be an IID.
	IID *dcom.IID `idl:"name:riid" json:"iid"`
}

func (o *CreateInstanceRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateInstanceOperation) *xxx_CreateInstanceOperation {
	if op == nil {
		op = &xxx_CreateInstanceOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.IID = o.IID
	return op
}

func (o *CreateInstanceRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateInstanceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.IID = op.IID
}
func (o *CreateInstanceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateInstanceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateInstanceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateInstanceResponse structure represents the CreateInstance operation response
type CreateInstanceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppvObj:  MUST be set to reference an existing instance of the coclass described
	// by the ITypeInfo server, if the coclass was declared with the [predeclid] or [appobject]
	// attributes and an instance of the coclass exists. MUST be set to NULL if the coclass
	// was declared with the [noncreatable] attribute. Otherwise, MUST be set to a new instance
	// of the coclass described by the ITypeInfo server or NULL if the class cannot be instantiated.
	Object *dcom.Unknown `idl:"name:ppvObj" json:"object"`
	// Return: The CreateInstance return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateInstanceResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateInstanceOperation) *xxx_CreateInstanceOperation {
	if op == nil {
		op = &xxx_CreateInstanceOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Object = o.Object
	op.Return = o.Return
	return op
}

func (o *CreateInstanceResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateInstanceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Object = op.Object
	o.Return = op.Return
}
func (o *CreateInstanceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateInstanceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateInstanceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMopsOperation structure represents the GetMops operation
type xxx_GetMopsOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	MemberID int32          `idl:"name:memid" json:"member_id"`
	Mops     *oaut.String   `idl:"name:pBstrMops" json:"mops"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMopsOperation) OpNum() int { return 17 }

func (o *xxx_GetMopsOperation) OpName() string { return "/ITypeInfo/v0/GetMops" }

func (o *xxx_GetMopsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMopsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// memid {in} (1:{alias=MEMBERID, names=LONG}(int32))
	{
		if err := w.WriteData(o.MemberID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMopsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// memid {in} (1:{alias=MEMBERID, names=LONG}(int32))
	{
		if err := w.ReadData(&o.MemberID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMopsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMopsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pBstrMops {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Mops != nil {
			_ptr_pBstrMops := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Mops != nil {
					if err := o.Mops.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Mops, _ptr_pBstrMops); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMopsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pBstrMops {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pBstrMops := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Mops == nil {
				o.Mops = &oaut.String{}
			}
			if err := o.Mops.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pBstrMops := func(ptr interface{}) { o.Mops = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Mops, _s_pBstrMops, _ptr_pBstrMops); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetMopsRequest structure represents the GetMops operation request
type GetMopsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// memid: MUST be a nonzero MEMBERID.
	MemberID int32 `idl:"name:memid" json:"member_id"`
}

func (o *GetMopsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMopsOperation) *xxx_GetMopsOperation {
	if op == nil {
		op = &xxx_GetMopsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MemberID = o.MemberID
	return op
}

func (o *GetMopsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMopsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MemberID = op.MemberID
}
func (o *GetMopsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMopsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMopsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMopsResponse structure represents the GetMops operation response
type GetMopsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pBstrMops: MUST be set to a NULL BSTR.
	Mops *oaut.String `idl:"name:pBstrMops" json:"mops"`
	// Return: The GetMops return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMopsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMopsOperation) *xxx_GetMopsOperation {
	if op == nil {
		op = &xxx_GetMopsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Mops = o.Mops
	op.Return = o.Return
	return op
}

func (o *GetMopsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMopsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Mops = op.Mops
	o.Return = op.Return
}
func (o *GetMopsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMopsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMopsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetContainingTypeLibOperation structure represents the GetContainingTypeLib operation
type xxx_GetContainingTypeLibOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Lib    *oaut.TypeLib  `idl:"name:ppTLib" json:"lib"`
	Index  uint32         `idl:"name:pIndex" json:"index"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetContainingTypeLibOperation) OpNum() int { return 18 }

func (o *xxx_GetContainingTypeLibOperation) OpName() string {
	return "/ITypeInfo/v0/GetContainingTypeLib"
}

func (o *xxx_GetContainingTypeLibOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetContainingTypeLibOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetContainingTypeLibOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetContainingTypeLibOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetContainingTypeLibOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppTLib {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ITypeLib}(interface))
	{
		if o.Lib != nil {
			_ptr_ppTLib := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Lib != nil {
					if err := o.Lib.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.TypeLib{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Lib, _ptr_ppTLib); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pIndex {out} (1:{pointer=ref}*(1))(2:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetContainingTypeLibOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppTLib {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ITypeLib}(interface))
	{
		_ptr_ppTLib := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Lib == nil {
				o.Lib = &oaut.TypeLib{}
			}
			if err := o.Lib.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppTLib := func(ptr interface{}) { o.Lib = *ptr.(**oaut.TypeLib) }
		if err := w.ReadPointer(&o.Lib, _s_ppTLib, _ptr_ppTLib); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pIndex {out} (1:{pointer=ref}*(1))(2:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetContainingTypeLibRequest structure represents the GetContainingTypeLib operation request
type GetContainingTypeLibRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetContainingTypeLibRequest) xxx_ToOp(ctx context.Context, op *xxx_GetContainingTypeLibOperation) *xxx_GetContainingTypeLibOperation {
	if op == nil {
		op = &xxx_GetContainingTypeLibOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetContainingTypeLibRequest) xxx_FromOp(ctx context.Context, op *xxx_GetContainingTypeLibOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetContainingTypeLibRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetContainingTypeLibRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetContainingTypeLibOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetContainingTypeLibResponse structure represents the GetContainingTypeLib operation response
type GetContainingTypeLibResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppTLib: MUST be set to a reference to an ITypeLib server instance (see section 3.11).
	Lib *oaut.TypeLib `idl:"name:ppTLib" json:"lib"`
	// pIndex: MUST be set to the index value of the current automation type description
	// within the type information table (see section 3.11.1).
	Index uint32 `idl:"name:pIndex" json:"index"`
	// Return: The GetContainingTypeLib return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetContainingTypeLibResponse) xxx_ToOp(ctx context.Context, op *xxx_GetContainingTypeLibOperation) *xxx_GetContainingTypeLibOperation {
	if op == nil {
		op = &xxx_GetContainingTypeLibOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Lib = o.Lib
	op.Index = o.Index
	op.Return = o.Return
	return op
}

func (o *GetContainingTypeLibResponse) xxx_FromOp(ctx context.Context, op *xxx_GetContainingTypeLibOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Lib = op.Lib
	o.Index = op.Index
	o.Return = op.Return
}
func (o *GetContainingTypeLibResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetContainingTypeLibResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetContainingTypeLibOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
