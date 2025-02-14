package itypeinfo

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
)

// ITypeInfo server interface.
type TypeInfoServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

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
	GetTypeAttribute(context.Context, *GetTypeAttributeRequest) (*GetTypeAttributeResponse, error)

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
	GetTypeComp(context.Context, *GetTypeCompRequest) (*GetTypeCompResponse, error)

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
	GetFuncDesc(context.Context, *GetFuncDescRequest) (*GetFuncDescResponse, error)

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
	GetVarDesc(context.Context, *GetVarDescRequest) (*GetVarDescResponse, error)

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
	GetNames(context.Context, *GetNamesRequest) (*GetNamesResponse, error)

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
	GetReferenceTypeOfImplType(context.Context, *GetReferenceTypeOfImplTypeRequest) (*GetReferenceTypeOfImplTypeResponse, error)

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
	GetImplTypeFlags(context.Context, *GetImplTypeFlagsRequest) (*GetImplTypeFlagsResponse, error)

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
	GetDocumentation(context.Context, *GetDocumentationRequest) (*GetDocumentationResponse, error)

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
	GetDLLEntry(context.Context, *GetDLLEntryRequest) (*GetDLLEntryResponse, error)

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
	GetReferenceTypeInfo(context.Context, *GetReferenceTypeInfoRequest) (*GetReferenceTypeInfoResponse, error)

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
	CreateInstance(context.Context, *CreateInstanceRequest) (*CreateInstanceResponse, error)

	// The GetMops method has no effect when called across the wire.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	GetMops(context.Context, *GetMopsRequest) (*GetMopsResponse, error)

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
	GetContainingTypeLib(context.Context, *GetContainingTypeLibRequest) (*GetContainingTypeLibResponse, error)

	// Opnum19NotUsedOnWire operation.
	// Opnum19NotUsedOnWire

	// Opnum20NotUsedOnWire operation.
	// Opnum20NotUsedOnWire

	// Opnum21NotUsedOnWire operation.
	// Opnum21NotUsedOnWire
}

func RegisterTypeInfoServer(conn dcerpc.Conn, o TypeInfoServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTypeInfoServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TypeInfoSyntaxV0_0))...)
}

func NewTypeInfoServerHandle(o TypeInfoServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TypeInfoServerHandle(ctx, o, opNum, r)
	}
}

func TypeInfoServerHandle(ctx context.Context, o TypeInfoServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetTypeAttr
		op := &xxx_GetTypeAttributeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTypeAttributeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTypeAttribute(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetTypeComp
		op := &xxx_GetTypeCompOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTypeCompRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTypeComp(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // GetFuncDesc
		op := &xxx_GetFuncDescOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFuncDescRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFuncDesc(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // GetVarDesc
		op := &xxx_GetVarDescOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetVarDescRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetVarDesc(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // GetNames
		op := &xxx_GetNamesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNamesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNames(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // GetRefTypeOfImplType
		op := &xxx_GetReferenceTypeOfImplTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetReferenceTypeOfImplTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetReferenceTypeOfImplType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // GetImplTypeFlags
		op := &xxx_GetImplTypeFlagsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetImplTypeFlagsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetImplTypeFlags(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Opnum10NotUsedOnWire
		// Opnum10NotUsedOnWire
		return nil, nil
	case 11: // Opnum11NotUsedOnWire
		// Opnum11NotUsedOnWire
		return nil, nil
	case 12: // GetDocumentation
		op := &xxx_GetDocumentationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDocumentationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDocumentation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // GetDllEntry
		op := &xxx_GetDLLEntryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDLLEntryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDLLEntry(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // GetRefTypeInfo
		op := &xxx_GetReferenceTypeInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetReferenceTypeInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetReferenceTypeInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // Opnum15NotUsedOnWire
		// Opnum15NotUsedOnWire
		return nil, nil
	case 16: // CreateInstance
		op := &xxx_CreateInstanceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateInstanceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateInstance(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // GetMops
		op := &xxx_GetMopsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMopsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMops(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // GetContainingTypeLib
		op := &xxx_GetContainingTypeLibOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetContainingTypeLibRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetContainingTypeLib(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // Opnum19NotUsedOnWire
		// Opnum19NotUsedOnWire
		return nil, nil
	case 20: // Opnum20NotUsedOnWire
		// Opnum20NotUsedOnWire
		return nil, nil
	case 21: // Opnum21NotUsedOnWire
		// Opnum21NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented ITypeInfo
type UnimplementedTypeInfoServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedTypeInfoServer) GetTypeAttribute(context.Context, *GetTypeAttributeRequest) (*GetTypeAttributeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfoServer) GetTypeComp(context.Context, *GetTypeCompRequest) (*GetTypeCompResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfoServer) GetFuncDesc(context.Context, *GetFuncDescRequest) (*GetFuncDescResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfoServer) GetVarDesc(context.Context, *GetVarDescRequest) (*GetVarDescResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfoServer) GetNames(context.Context, *GetNamesRequest) (*GetNamesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfoServer) GetReferenceTypeOfImplType(context.Context, *GetReferenceTypeOfImplTypeRequest) (*GetReferenceTypeOfImplTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfoServer) GetImplTypeFlags(context.Context, *GetImplTypeFlagsRequest) (*GetImplTypeFlagsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfoServer) GetDocumentation(context.Context, *GetDocumentationRequest) (*GetDocumentationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfoServer) GetDLLEntry(context.Context, *GetDLLEntryRequest) (*GetDLLEntryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfoServer) GetReferenceTypeInfo(context.Context, *GetReferenceTypeInfoRequest) (*GetReferenceTypeInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfoServer) CreateInstance(context.Context, *CreateInstanceRequest) (*CreateInstanceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfoServer) GetMops(context.Context, *GetMopsRequest) (*GetMopsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfoServer) GetContainingTypeLib(context.Context, *GetContainingTypeLibRequest) (*GetContainingTypeLibResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ TypeInfoServer = (*UnimplementedTypeInfoServer)(nil)
