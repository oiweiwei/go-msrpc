package itypeinfo2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	itypeinfo "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/itypeinfo/v0"
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
	_ = itypeinfo.GoPackage
)

// ITypeInfo2 server interface.
type TypeInfo2Server interface {

	// ITypeInfo base class.
	itypeinfo.TypeInfoServer

	// The GetTypeKind method returns the TYPEKIND value associated with the automation
	// type description.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	GetTypeKind(context.Context, *GetTypeKindRequest) (*GetTypeKindResponse, error)

	// The GetTypeFlags method returns the TYPEFLAGS value associated with the automation
	// type description.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	GetTypeFlags(context.Context, *GetTypeFlagsRequest) (*GetTypeFlagsResponse, error)

	// The GetFuncIndexOfMemId method retrieves the location of an element in the data member
	// table, given its associated MEMBERID.
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
	// The values of memid and invKind did not specify a member of the type. See [MS-ERREF].
	GetFuncIndexOfMemberIDs(context.Context, *GetFuncIndexOfMemberIDsRequest) (*GetFuncIndexOfMemberIDsResponse, error)

	// The GetVarIndexOfMemId method retrieves the location of an element in the data member
	// table by using the associated MEMBERID of the element.
	//
	// The method is received by the server in an RPC_REQUEST packet.
	//
	// Return Values: The method MUST return information in an HRESULT data structure that
	// is defined in [MS-ERREF] section 2.1. The severity bit in the structure identifies
	// the following conditions:
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
	// The value of memid did not specify a member of the type. See [MS-ERREF].
	GetVarIndexOfMemberIDs(context.Context, *GetVarIndexOfMemberIDsRequest) (*GetVarIndexOfMemberIDsResponse, error)

	// The GetCustData method retrieves the value of a custom data item associated with
	// the automation type library.
	//
	// The GetCustData method retrieves the value of a custom data item associated with
	// the type.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	GetCustomData(context.Context, *GetCustomDataRequest) (*GetCustomDataResponse, error)

	// The GetFuncCustData method retrieves the value of a custom data item associated with
	// the specified method.
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
	// table. See [MS-ERREF].
	GetFuncCustomData(context.Context, *GetFuncCustomDataRequest) (*GetFuncCustomDataResponse, error)

	// The GetParamCustData method retrieves the value of a custom data item associated
	// with the specified method parameter.
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
	// The value of indexFunc did not specify the ordinal position of an element in the
	// interface table, or the value of indexParam did not specify the ordinal position
	// of an element in the method's parameter table. See [MS-ERREF].
	GetParamCustomData(context.Context, *GetParamCustomDataRequest) (*GetParamCustomDataResponse, error)

	// The GetVarCustData method retrieves the value of a custom data item associated with
	// the specified data member.
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
	// The value of index did not specify the ordinal position of an element in the data
	// member table. See [MS-ERREF].
	GetVarCustomData(context.Context, *GetVarCustomDataRequest) (*GetVarCustomDataResponse, error)

	// The GetImplTypeCustData method retrieves the value of a custom data item associated
	// with the specified interface member of a coclass.
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
	// table, or the type is not a coclass. See [MS-ERREF].
	GetImplTypeCustomData(context.Context, *GetImplTypeCustomDataRequest) (*GetImplTypeCustomDataResponse, error)

	// The GetDocumentation2 method retrieves values associated with the automation type
	// library.
	//
	// The GetDocumentation2 method retrieves values associated with a type member.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	GetDocumentation2(context.Context, *GetDocumentation2Request) (*GetDocumentation2Response, error)

	// The GetAllCustData method retrieves the values of all custom data items associated
	// with the automation type library.
	//
	// The GetAllCustData method retrieves all the custom data items associated with the
	// automation type description.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	GetAllCustomData(context.Context, *GetAllCustomDataRequest) (*GetAllCustomDataResponse, error)

	// The GetAllFuncCustData method retrieves all of the custom data items associated with
	// the specified method.
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
	// The value of index did not specify the ordinal position of an element in the data
	// member table. See [MS-ERREF].
	GetAllFuncCustomData(context.Context, *GetAllFuncCustomDataRequest) (*GetAllFuncCustomDataResponse, error)

	// The GetAllParamCustData method retrieves all of the custom data items associated
	// with the specified parameter.
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
	// The value of indexFunc did not specify the ordinal position of an element in the
	// method table, or the value of indexParam did not specify the ordinal position of
	// an element in the parameter table.
	GetAllParamCustomData(context.Context, *GetAllParamCustomDataRequest) (*GetAllParamCustomDataResponse, error)

	// The GetAllVarCustData method retrieves all of the custom data items associated with
	// the specified data member.
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
	// The value of index did not specify the ordinal position of an element in the data
	// member table. See [MS-ERREF].
	GetAllVarCustomData(context.Context, *GetAllVarCustomDataRequest) (*GetAllVarCustomDataResponse, error)

	// The GetAllImplTypeCustData method retrieves all of the custom data items associated
	// with the specified data member.
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
	// table, or the type is not a coclass ( 5583e1b8-454c-4147-9f56-f72416a15bee#gt_670b0ee2-d101-41b0-ac77-6ac7dbeee7dc
	// ). See [MS-ERREF].
	GetAllImplTypeCustomData(context.Context, *GetAllImplTypeCustomDataRequest) (*GetAllImplTypeCustomDataResponse, error)
}

func RegisterTypeInfo2Server(conn dcerpc.Conn, o TypeInfo2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTypeInfo2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TypeInfo2SyntaxV0_0))...)
}

func NewTypeInfo2ServerHandle(o TypeInfo2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TypeInfo2ServerHandle(ctx, o, opNum, r)
	}
}

func TypeInfo2ServerHandle(ctx context.Context, o TypeInfo2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 22 {
		// ITypeInfo base method.
		return itypeinfo.TypeInfoServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 22: // GetTypeKind
		op := &xxx_GetTypeKindOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTypeKindRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTypeKind(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // GetTypeFlags
		op := &xxx_GetTypeFlagsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTypeFlagsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTypeFlags(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // GetFuncIndexOfMemId
		op := &xxx_GetFuncIndexOfMemberIDsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFuncIndexOfMemberIDsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFuncIndexOfMemberIDs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // GetVarIndexOfMemId
		op := &xxx_GetVarIndexOfMemberIDsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetVarIndexOfMemberIDsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetVarIndexOfMemberIDs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // GetCustData
		op := &xxx_GetCustomDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCustomDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCustomData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // GetFuncCustData
		op := &xxx_GetFuncCustomDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFuncCustomDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFuncCustomData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // GetParamCustData
		op := &xxx_GetParamCustomDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetParamCustomDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetParamCustomData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // GetVarCustData
		op := &xxx_GetVarCustomDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetVarCustomDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetVarCustomData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // GetImplTypeCustData
		op := &xxx_GetImplTypeCustomDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetImplTypeCustomDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetImplTypeCustomData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // GetDocumentation2
		op := &xxx_GetDocumentation2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDocumentation2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDocumentation2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // GetAllCustData
		op := &xxx_GetAllCustomDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAllCustomDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAllCustomData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // GetAllFuncCustData
		op := &xxx_GetAllFuncCustomDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAllFuncCustomDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAllFuncCustomData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // GetAllParamCustData
		op := &xxx_GetAllParamCustomDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAllParamCustomDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAllParamCustomData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // GetAllVarCustData
		op := &xxx_GetAllVarCustomDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAllVarCustomDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAllVarCustomData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // GetAllImplTypeCustData
		op := &xxx_GetAllImplTypeCustomDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAllImplTypeCustomDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAllImplTypeCustomData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ITypeInfo2
type UnimplementedTypeInfo2Server struct {
	itypeinfo.UnimplementedTypeInfoServer
}

func (UnimplementedTypeInfo2Server) GetTypeKind(context.Context, *GetTypeKindRequest) (*GetTypeKindResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfo2Server) GetTypeFlags(context.Context, *GetTypeFlagsRequest) (*GetTypeFlagsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfo2Server) GetFuncIndexOfMemberIDs(context.Context, *GetFuncIndexOfMemberIDsRequest) (*GetFuncIndexOfMemberIDsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfo2Server) GetVarIndexOfMemberIDs(context.Context, *GetVarIndexOfMemberIDsRequest) (*GetVarIndexOfMemberIDsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfo2Server) GetCustomData(context.Context, *GetCustomDataRequest) (*GetCustomDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfo2Server) GetFuncCustomData(context.Context, *GetFuncCustomDataRequest) (*GetFuncCustomDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfo2Server) GetParamCustomData(context.Context, *GetParamCustomDataRequest) (*GetParamCustomDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfo2Server) GetVarCustomData(context.Context, *GetVarCustomDataRequest) (*GetVarCustomDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfo2Server) GetImplTypeCustomData(context.Context, *GetImplTypeCustomDataRequest) (*GetImplTypeCustomDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfo2Server) GetDocumentation2(context.Context, *GetDocumentation2Request) (*GetDocumentation2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfo2Server) GetAllCustomData(context.Context, *GetAllCustomDataRequest) (*GetAllCustomDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfo2Server) GetAllFuncCustomData(context.Context, *GetAllFuncCustomDataRequest) (*GetAllFuncCustomDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfo2Server) GetAllParamCustomData(context.Context, *GetAllParamCustomDataRequest) (*GetAllParamCustomDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfo2Server) GetAllVarCustomData(context.Context, *GetAllVarCustomDataRequest) (*GetAllVarCustomDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeInfo2Server) GetAllImplTypeCustomData(context.Context, *GetAllImplTypeCustomDataRequest) (*GetAllImplTypeCustomDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ TypeInfo2Server = (*UnimplementedTypeInfo2Server)(nil)
