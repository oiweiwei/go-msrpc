package itypeinfo2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
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
		in := &GetTypeKindRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTypeKind(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // GetTypeFlags
		in := &GetTypeFlagsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTypeFlags(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // GetFuncIndexOfMemId
		in := &GetFuncIndexOfMemberIDsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFuncIndexOfMemberIDs(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // GetVarIndexOfMemId
		in := &GetVarIndexOfMemberIDsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetVarIndexOfMemberIDs(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // GetCustData
		in := &GetCustomDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCustomData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // GetFuncCustData
		in := &GetFuncCustomDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFuncCustomData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // GetParamCustData
		in := &GetParamCustomDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetParamCustomData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // GetVarCustData
		in := &GetVarCustomDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetVarCustomData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // GetImplTypeCustData
		in := &GetImplTypeCustomDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetImplTypeCustomData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // GetDocumentation2
		in := &GetDocumentation2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDocumentation2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // GetAllCustData
		in := &GetAllCustomDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAllCustomData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // GetAllFuncCustData
		in := &GetAllFuncCustomDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAllFuncCustomData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // GetAllParamCustData
		in := &GetAllParamCustomDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAllParamCustomData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // GetAllVarCustData
		in := &GetAllVarCustomDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAllVarCustomData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // GetAllImplTypeCustData
		in := &GetAllImplTypeCustomDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAllImplTypeCustomData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
