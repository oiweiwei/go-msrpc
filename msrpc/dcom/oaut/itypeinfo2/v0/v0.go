package itypeinfo2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	itypeinfo "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/itypeinfo/v0"
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
	_ = itypeinfo.GoPackage
	_ = oaut.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/oaut"
)

var (
	// ITypeInfo2 interface identifier 00020412-0000-0000-c000-000000000046
	TypeInfo2IID = &dcom.IID{Data1: 0x00020412, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	// Syntax UUID
	TypeInfo2SyntaxUUID = &uuid.UUID{TimeLow: 0x20412, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	// Syntax ID
	TypeInfo2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: TypeInfo2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ITypeInfo2 interface.
type TypeInfo2Client interface {

	// ITypeInfo retrieval method.
	TypeInfo() itypeinfo.TypeInfoClient

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
	GetTypeKind(context.Context, *GetTypeKindRequest, ...dcerpc.CallOption) (*GetTypeKindResponse, error)

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
	GetTypeFlags(context.Context, *GetTypeFlagsRequest, ...dcerpc.CallOption) (*GetTypeFlagsResponse, error)

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
	GetFuncIndexOfMemberIDs(context.Context, *GetFuncIndexOfMemberIDsRequest, ...dcerpc.CallOption) (*GetFuncIndexOfMemberIDsResponse, error)

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
	GetVarIndexOfMemberIDs(context.Context, *GetVarIndexOfMemberIDsRequest, ...dcerpc.CallOption) (*GetVarIndexOfMemberIDsResponse, error)

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
	GetCustomData(context.Context, *GetCustomDataRequest, ...dcerpc.CallOption) (*GetCustomDataResponse, error)

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
	GetFuncCustomData(context.Context, *GetFuncCustomDataRequest, ...dcerpc.CallOption) (*GetFuncCustomDataResponse, error)

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
	GetParamCustomData(context.Context, *GetParamCustomDataRequest, ...dcerpc.CallOption) (*GetParamCustomDataResponse, error)

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
	GetVarCustomData(context.Context, *GetVarCustomDataRequest, ...dcerpc.CallOption) (*GetVarCustomDataResponse, error)

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
	GetImplTypeCustomData(context.Context, *GetImplTypeCustomDataRequest, ...dcerpc.CallOption) (*GetImplTypeCustomDataResponse, error)

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
	GetDocumentation2(context.Context, *GetDocumentation2Request, ...dcerpc.CallOption) (*GetDocumentation2Response, error)

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
	GetAllCustomData(context.Context, *GetAllCustomDataRequest, ...dcerpc.CallOption) (*GetAllCustomDataResponse, error)

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
	GetAllFuncCustomData(context.Context, *GetAllFuncCustomDataRequest, ...dcerpc.CallOption) (*GetAllFuncCustomDataResponse, error)

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
	GetAllParamCustomData(context.Context, *GetAllParamCustomDataRequest, ...dcerpc.CallOption) (*GetAllParamCustomDataResponse, error)

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
	GetAllVarCustomData(context.Context, *GetAllVarCustomDataRequest, ...dcerpc.CallOption) (*GetAllVarCustomDataResponse, error)

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
	GetAllImplTypeCustomData(context.Context, *GetAllImplTypeCustomDataRequest, ...dcerpc.CallOption) (*GetAllImplTypeCustomDataResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) TypeInfo2Client
}

type xxx_DefaultTypeInfo2Client struct {
	itypeinfo.TypeInfoClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultTypeInfo2Client) TypeInfo() itypeinfo.TypeInfoClient {
	return o.TypeInfoClient
}

func (o *xxx_DefaultTypeInfo2Client) GetTypeKind(ctx context.Context, in *GetTypeKindRequest, opts ...dcerpc.CallOption) (*GetTypeKindResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetTypeKindResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfo2Client) GetTypeFlags(ctx context.Context, in *GetTypeFlagsRequest, opts ...dcerpc.CallOption) (*GetTypeFlagsResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetTypeFlagsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfo2Client) GetFuncIndexOfMemberIDs(ctx context.Context, in *GetFuncIndexOfMemberIDsRequest, opts ...dcerpc.CallOption) (*GetFuncIndexOfMemberIDsResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetFuncIndexOfMemberIDsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfo2Client) GetVarIndexOfMemberIDs(ctx context.Context, in *GetVarIndexOfMemberIDsRequest, opts ...dcerpc.CallOption) (*GetVarIndexOfMemberIDsResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetVarIndexOfMemberIDsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfo2Client) GetCustomData(ctx context.Context, in *GetCustomDataRequest, opts ...dcerpc.CallOption) (*GetCustomDataResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetCustomDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfo2Client) GetFuncCustomData(ctx context.Context, in *GetFuncCustomDataRequest, opts ...dcerpc.CallOption) (*GetFuncCustomDataResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetFuncCustomDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfo2Client) GetParamCustomData(ctx context.Context, in *GetParamCustomDataRequest, opts ...dcerpc.CallOption) (*GetParamCustomDataResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetParamCustomDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfo2Client) GetVarCustomData(ctx context.Context, in *GetVarCustomDataRequest, opts ...dcerpc.CallOption) (*GetVarCustomDataResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetVarCustomDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfo2Client) GetImplTypeCustomData(ctx context.Context, in *GetImplTypeCustomDataRequest, opts ...dcerpc.CallOption) (*GetImplTypeCustomDataResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetImplTypeCustomDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfo2Client) GetDocumentation2(ctx context.Context, in *GetDocumentation2Request, opts ...dcerpc.CallOption) (*GetDocumentation2Response, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetDocumentation2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfo2Client) GetAllCustomData(ctx context.Context, in *GetAllCustomDataRequest, opts ...dcerpc.CallOption) (*GetAllCustomDataResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetAllCustomDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfo2Client) GetAllFuncCustomData(ctx context.Context, in *GetAllFuncCustomDataRequest, opts ...dcerpc.CallOption) (*GetAllFuncCustomDataResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetAllFuncCustomDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfo2Client) GetAllParamCustomData(ctx context.Context, in *GetAllParamCustomDataRequest, opts ...dcerpc.CallOption) (*GetAllParamCustomDataResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetAllParamCustomDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfo2Client) GetAllVarCustomData(ctx context.Context, in *GetAllVarCustomDataRequest, opts ...dcerpc.CallOption) (*GetAllVarCustomDataResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetAllVarCustomDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfo2Client) GetAllImplTypeCustomData(ctx context.Context, in *GetAllImplTypeCustomDataRequest, opts ...dcerpc.CallOption) (*GetAllImplTypeCustomDataResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetAllImplTypeCustomDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTypeInfo2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultTypeInfo2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultTypeInfo2Client) IPID(ctx context.Context, ipid *dcom.IPID) TypeInfo2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultTypeInfo2Client{
		TypeInfoClient: o.TypeInfoClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewTypeInfo2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (TypeInfo2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(TypeInfo2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := itypeinfo.NewTypeInfoClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultTypeInfo2Client{
		TypeInfoClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetTypeKindOperation structure represents the GetTypeKind operation
type xxx_GetTypeKindOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	TypeKind oaut.TypeKind  `idl:"name:pTypeKind" json:"type_kind"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTypeKindOperation) OpNum() int { return 22 }

func (o *xxx_GetTypeKindOperation) OpName() string { return "/ITypeInfo2/v0/GetTypeKind" }

func (o *xxx_GetTypeKindOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeKindOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetTypeKindOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetTypeKindOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeKindOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pTypeKind {out} (1:{pointer=ref}*(1))(2:{v1_enum, alias=TYPEKIND}(enum))
	{
		if err := w.WriteData(uint32(o.TypeKind)); err != nil {
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

func (o *xxx_GetTypeKindOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pTypeKind {out} (1:{pointer=ref}*(1))(2:{v1_enum, alias=TYPEKIND}(enum))
	{
		if err := w.ReadData((*uint32)(&o.TypeKind)); err != nil {
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

// GetTypeKindRequest structure represents the GetTypeKind operation request
type GetTypeKindRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetTypeKindRequest) xxx_ToOp(ctx context.Context) *xxx_GetTypeKindOperation {
	if o == nil {
		return &xxx_GetTypeKindOperation{}
	}
	return &xxx_GetTypeKindOperation{
		This: o.This,
	}
}

func (o *GetTypeKindRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTypeKindOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTypeKindRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetTypeKindRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeKindOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTypeKindResponse structure represents the GetTypeKind operation response
type GetTypeKindResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pTypeKind: MUST be set to the TYPEKIND value associated with the automation type
	// description, as specified in section 2.2.17.
	TypeKind oaut.TypeKind `idl:"name:pTypeKind" json:"type_kind"`
	// Return: The GetTypeKind return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTypeKindResponse) xxx_ToOp(ctx context.Context) *xxx_GetTypeKindOperation {
	if o == nil {
		return &xxx_GetTypeKindOperation{}
	}
	return &xxx_GetTypeKindOperation{
		That:     o.That,
		TypeKind: o.TypeKind,
		Return:   o.Return,
	}
}

func (o *GetTypeKindResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTypeKindOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TypeKind = op.TypeKind
	o.Return = op.Return
}
func (o *GetTypeKindResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetTypeKindResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeKindOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTypeFlagsOperation structure represents the GetTypeFlags operation
type xxx_GetTypeFlagsOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	TypeFlags uint32         `idl:"name:pTypeFlags" json:"type_flags"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTypeFlagsOperation) OpNum() int { return 23 }

func (o *xxx_GetTypeFlagsOperation) OpName() string { return "/ITypeInfo2/v0/GetTypeFlags" }

func (o *xxx_GetTypeFlagsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeFlagsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetTypeFlagsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetTypeFlagsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeFlagsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pTypeFlags {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.TypeFlags); err != nil {
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

func (o *xxx_GetTypeFlagsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pTypeFlags {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.TypeFlags); err != nil {
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

// GetTypeFlagsRequest structure represents the GetTypeFlags operation request
type GetTypeFlagsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetTypeFlagsRequest) xxx_ToOp(ctx context.Context) *xxx_GetTypeFlagsOperation {
	if o == nil {
		return &xxx_GetTypeFlagsOperation{}
	}
	return &xxx_GetTypeFlagsOperation{
		This: o.This,
	}
}

func (o *GetTypeFlagsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTypeFlagsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTypeFlagsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetTypeFlagsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeFlagsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTypeFlagsResponse structure represents the GetTypeFlags operation response
type GetTypeFlagsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pTypeFlags: MUST be set either to a combination of the TYPEFLAGS type feature constants
	// specified in section 2.2.16, or to 0.
	TypeFlags uint32 `idl:"name:pTypeFlags" json:"type_flags"`
	// Return: The GetTypeFlags return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTypeFlagsResponse) xxx_ToOp(ctx context.Context) *xxx_GetTypeFlagsOperation {
	if o == nil {
		return &xxx_GetTypeFlagsOperation{}
	}
	return &xxx_GetTypeFlagsOperation{
		That:      o.That,
		TypeFlags: o.TypeFlags,
		Return:    o.Return,
	}
}

func (o *GetTypeFlagsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTypeFlagsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TypeFlags = op.TypeFlags
	o.Return = op.Return
}
func (o *GetTypeFlagsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetTypeFlagsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeFlagsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFuncIndexOfMemberIDsOperation structure represents the GetFuncIndexOfMemId operation
type xxx_GetFuncIndexOfMemberIDsOperation struct {
	This      *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat  `idl:"name:That" json:"that"`
	MemberID  int32           `idl:"name:memid" json:"member_id"`
	InvKind   oaut.InvokeKind `idl:"name:invKind" json:"inv_kind"`
	FuncIndex uint32          `idl:"name:pFuncIndex" json:"func_index"`
	Return    int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFuncIndexOfMemberIDsOperation) OpNum() int { return 24 }

func (o *xxx_GetFuncIndexOfMemberIDsOperation) OpName() string {
	return "/ITypeInfo2/v0/GetFuncIndexOfMemId"
}

func (o *xxx_GetFuncIndexOfMemberIDsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFuncIndexOfMemberIDsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
		if err := w.WriteData(uint32(o.InvKind)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFuncIndexOfMemberIDsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
		if err := w.ReadData((*uint32)(&o.InvKind)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFuncIndexOfMemberIDsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFuncIndexOfMemberIDsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pFuncIndex {out} (1:{pointer=ref}*(1))(2:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.FuncIndex); err != nil {
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

func (o *xxx_GetFuncIndexOfMemberIDsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pFuncIndex {out} (1:{pointer=ref}*(1))(2:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.FuncIndex); err != nil {
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

// GetFuncIndexOfMemberIDsRequest structure represents the GetFuncIndexOfMemId operation request
type GetFuncIndexOfMemberIDsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// memid: MUST be a MEMBERID, as specified in section 2.2.35.
	MemberID int32 `idl:"name:memid" json:"member_id"`
	// invKind: MUST be set to one of the values of the INVOKEKIND enumeration (as specified
	// in section 2.2.14) or to 0.
	InvKind oaut.InvokeKind `idl:"name:invKind" json:"inv_kind"`
}

func (o *GetFuncIndexOfMemberIDsRequest) xxx_ToOp(ctx context.Context) *xxx_GetFuncIndexOfMemberIDsOperation {
	if o == nil {
		return &xxx_GetFuncIndexOfMemberIDsOperation{}
	}
	return &xxx_GetFuncIndexOfMemberIDsOperation{
		This:     o.This,
		MemberID: o.MemberID,
		InvKind:  o.InvKind,
	}
}

func (o *GetFuncIndexOfMemberIDsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFuncIndexOfMemberIDsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MemberID = op.MemberID
	o.InvKind = op.InvKind
}
func (o *GetFuncIndexOfMemberIDsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetFuncIndexOfMemberIDsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFuncIndexOfMemberIDsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFuncIndexOfMemberIDsResponse structure represents the GetFuncIndexOfMemId operation response
type GetFuncIndexOfMemberIDsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pFuncIndex: MUST be set to the ordinal position in the method table of the element
	// specified by the values of memid and invKind as described below, or to –1 if no
	// such element exists.
	//
	// If invKind is not 0, the specified element is the one whose MEMBERID matches the
	// value of memid, and whose associated INVOKEKIND constant (see FUNCDESC) matches the
	// value of invKind.
	FuncIndex uint32 `idl:"name:pFuncIndex" json:"func_index"`
	// Return: The GetFuncIndexOfMemId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFuncIndexOfMemberIDsResponse) xxx_ToOp(ctx context.Context) *xxx_GetFuncIndexOfMemberIDsOperation {
	if o == nil {
		return &xxx_GetFuncIndexOfMemberIDsOperation{}
	}
	return &xxx_GetFuncIndexOfMemberIDsOperation{
		That:      o.That,
		FuncIndex: o.FuncIndex,
		Return:    o.Return,
	}
}

func (o *GetFuncIndexOfMemberIDsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFuncIndexOfMemberIDsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FuncIndex = op.FuncIndex
	o.Return = op.Return
}
func (o *GetFuncIndexOfMemberIDsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetFuncIndexOfMemberIDsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFuncIndexOfMemberIDsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetVarIndexOfMemberIDsOperation structure represents the GetVarIndexOfMemId operation
type xxx_GetVarIndexOfMemberIDsOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	MemberID int32          `idl:"name:memid" json:"member_id"`
	VarIndex uint32         `idl:"name:pVarIndex" json:"var_index"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetVarIndexOfMemberIDsOperation) OpNum() int { return 25 }

func (o *xxx_GetVarIndexOfMemberIDsOperation) OpName() string {
	return "/ITypeInfo2/v0/GetVarIndexOfMemId"
}

func (o *xxx_GetVarIndexOfMemberIDsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVarIndexOfMemberIDsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetVarIndexOfMemberIDsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetVarIndexOfMemberIDsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVarIndexOfMemberIDsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pVarIndex {out} (1:{pointer=ref}*(1))(2:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.VarIndex); err != nil {
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

func (o *xxx_GetVarIndexOfMemberIDsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pVarIndex {out} (1:{pointer=ref}*(1))(2:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.VarIndex); err != nil {
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

// GetVarIndexOfMemberIDsRequest structure represents the GetVarIndexOfMemId operation request
type GetVarIndexOfMemberIDsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// memid: MUST be a MEMBERID, as specified in section 2.2.35. MUST NOT be MEMBERID_NIL.
	MemberID int32 `idl:"name:memid" json:"member_id"`
}

func (o *GetVarIndexOfMemberIDsRequest) xxx_ToOp(ctx context.Context) *xxx_GetVarIndexOfMemberIDsOperation {
	if o == nil {
		return &xxx_GetVarIndexOfMemberIDsOperation{}
	}
	return &xxx_GetVarIndexOfMemberIDsOperation{
		This:     o.This,
		MemberID: o.MemberID,
	}
}

func (o *GetVarIndexOfMemberIDsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetVarIndexOfMemberIDsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MemberID = op.MemberID
}
func (o *GetVarIndexOfMemberIDsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetVarIndexOfMemberIDsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVarIndexOfMemberIDsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetVarIndexOfMemberIDsResponse structure represents the GetVarIndexOfMemId operation response
type GetVarIndexOfMemberIDsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pVarIndex: MUST be set to the ordinal position in the data member table of the element
	// whose MEMBERID is specified by memid, if such an element exists. If the method returns
	// a failure code, the value MUST be ignored on receipt.
	VarIndex uint32 `idl:"name:pVarIndex" json:"var_index"`
	// Return: The GetVarIndexOfMemId return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetVarIndexOfMemberIDsResponse) xxx_ToOp(ctx context.Context) *xxx_GetVarIndexOfMemberIDsOperation {
	if o == nil {
		return &xxx_GetVarIndexOfMemberIDsOperation{}
	}
	return &xxx_GetVarIndexOfMemberIDsOperation{
		That:     o.That,
		VarIndex: o.VarIndex,
		Return:   o.Return,
	}
}

func (o *GetVarIndexOfMemberIDsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetVarIndexOfMemberIDsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.VarIndex = op.VarIndex
	o.Return = op.Return
}
func (o *GetVarIndexOfMemberIDsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetVarIndexOfMemberIDsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVarIndexOfMemberIDsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCustomDataOperation structure represents the GetCustData operation
type xxx_GetCustomDataOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	GUID     *dtyp.GUID     `idl:"name:guid" json:"guid"`
	VarValue *oaut.Variant  `idl:"name:pVarVal" json:"var_value"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCustomDataOperation) OpNum() int { return 26 }

func (o *xxx_GetCustomDataOperation) OpName() string { return "/ITypeInfo2/v0/GetCustData" }

func (o *xxx_GetCustomDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCustomDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// guid {in} (1:{alias=REFGUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID != nil {
			if err := o.GUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetCustomDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// guid {in} (1:{alias=REFGUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID == nil {
			o.GUID = &dtyp.GUID{}
		}
		if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCustomDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCustomDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pVarVal {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.VarValue != nil {
			_ptr_pVarVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.VarValue != nil {
					if err := o.VarValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.VarValue, _ptr_pVarVal); err != nil {
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

func (o *xxx_GetCustomDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pVarVal {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pVarVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.VarValue == nil {
				o.VarValue = &oaut.Variant{}
			}
			if err := o.VarValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pVarVal := func(ptr interface{}) { o.VarValue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.VarValue, _s_pVarVal, _ptr_pVarVal); err != nil {
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

// GetCustomDataRequest structure represents the GetCustData operation request
type GetCustomDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// guid: MUST be the GUID associated with the custom data item using the [custom] attribute,
	// as specified in section 2.2.49.2.
	//
	// guid: MUST be a GUID associated with the custom data item.
	GUID *dtyp.GUID `idl:"name:guid" json:"guid"`
}

func (o *GetCustomDataRequest) xxx_ToOp(ctx context.Context) *xxx_GetCustomDataOperation {
	if o == nil {
		return &xxx_GetCustomDataOperation{}
	}
	return &xxx_GetCustomDataOperation{
		This: o.This,
		GUID: o.GUID,
	}
}

func (o *GetCustomDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCustomDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.GUID = op.GUID
}
func (o *GetCustomDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetCustomDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCustomDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCustomDataResponse structure represents the GetCustData operation response
type GetCustomDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pVarVal: MUST be set to the value of the custom data item, or VT_EMPTY if there is
	// no custom data item associated with the specified GUID.
	//
	// pVarVal: MUST be set to the value associated with the GUID using the [custom] attribute
	// (as specified in section 2.2.49.3), or VT_EMPTY if the type does not have a value
	// associated with the GUID.
	VarValue *oaut.Variant `idl:"name:pVarVal" json:"var_value"`
	// Return: The GetCustData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCustomDataResponse) xxx_ToOp(ctx context.Context) *xxx_GetCustomDataOperation {
	if o == nil {
		return &xxx_GetCustomDataOperation{}
	}
	return &xxx_GetCustomDataOperation{
		That:     o.That,
		VarValue: o.VarValue,
		Return:   o.Return,
	}
}

func (o *GetCustomDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCustomDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.VarValue = op.VarValue
	o.Return = op.Return
}
func (o *GetCustomDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetCustomDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCustomDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFuncCustomDataOperation structure represents the GetFuncCustData operation
type xxx_GetFuncCustomDataOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Index    uint32         `idl:"name:index" json:"index"`
	GUID     *dtyp.GUID     `idl:"name:guid" json:"guid"`
	VarValue *oaut.Variant  `idl:"name:pVarVal" json:"var_value"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFuncCustomDataOperation) OpNum() int { return 27 }

func (o *xxx_GetFuncCustomDataOperation) OpName() string { return "/ITypeInfo2/v0/GetFuncCustData" }

func (o *xxx_GetFuncCustomDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFuncCustomDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// guid {in} (1:{alias=REFGUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID != nil {
			if err := o.GUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetFuncCustomDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// guid {in} (1:{alias=REFGUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID == nil {
			o.GUID = &dtyp.GUID{}
		}
		if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFuncCustomDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFuncCustomDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pVarVal {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.VarValue != nil {
			_ptr_pVarVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.VarValue != nil {
					if err := o.VarValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.VarValue, _ptr_pVarVal); err != nil {
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

func (o *xxx_GetFuncCustomDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pVarVal {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pVarVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.VarValue == nil {
				o.VarValue = &oaut.Variant{}
			}
			if err := o.VarValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pVarVal := func(ptr interface{}) { o.VarValue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.VarValue, _s_pVarVal, _ptr_pVarVal); err != nil {
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

// GetFuncCustomDataRequest structure represents the GetFuncCustData operation request
type GetFuncCustomDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// index: MUST specify an ordinal position in the method table and MUST be less than
	// the value of the cFuncs field in the TYPEATTR structure of the associated type, as
	// specified in sections 2.2.44 and 3.7.4.1.
	Index uint32 `idl:"name:index" json:"index"`
	// guid: MUST be the GUID associated with the custom data item using the [custom] attribute,
	// as specified in section 2.2.49.5.1.
	GUID *dtyp.GUID `idl:"name:guid" json:"guid"`
}

func (o *GetFuncCustomDataRequest) xxx_ToOp(ctx context.Context) *xxx_GetFuncCustomDataOperation {
	if o == nil {
		return &xxx_GetFuncCustomDataOperation{}
	}
	return &xxx_GetFuncCustomDataOperation{
		This:  o.This,
		Index: o.Index,
		GUID:  o.GUID,
	}
}

func (o *GetFuncCustomDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFuncCustomDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
	o.GUID = op.GUID
}
func (o *GetFuncCustomDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetFuncCustomDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFuncCustomDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFuncCustomDataResponse structure represents the GetFuncCustData operation response
type GetFuncCustomDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pVarVal: MUST be set to the value of the custom data item, or VT_EMPTY if index and
	// guid do not specify a custom data item.
	VarValue *oaut.Variant `idl:"name:pVarVal" json:"var_value"`
	// Return: The GetFuncCustData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFuncCustomDataResponse) xxx_ToOp(ctx context.Context) *xxx_GetFuncCustomDataOperation {
	if o == nil {
		return &xxx_GetFuncCustomDataOperation{}
	}
	return &xxx_GetFuncCustomDataOperation{
		That:     o.That,
		VarValue: o.VarValue,
		Return:   o.Return,
	}
}

func (o *GetFuncCustomDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFuncCustomDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.VarValue = op.VarValue
	o.Return = op.Return
}
func (o *GetFuncCustomDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetFuncCustomDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFuncCustomDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetParamCustomDataOperation structure represents the GetParamCustData operation
type xxx_GetParamCustomDataOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	IndexFunc  uint32         `idl:"name:indexFunc" json:"index_func"`
	IndexParam uint32         `idl:"name:indexParam" json:"index_param"`
	GUID       *dtyp.GUID     `idl:"name:guid" json:"guid"`
	VarValue   *oaut.Variant  `idl:"name:pVarVal" json:"var_value"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetParamCustomDataOperation) OpNum() int { return 28 }

func (o *xxx_GetParamCustomDataOperation) OpName() string { return "/ITypeInfo2/v0/GetParamCustData" }

func (o *xxx_GetParamCustomDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetParamCustomDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// indexFunc {in} (1:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.IndexFunc); err != nil {
			return err
		}
	}
	// indexParam {in} (1:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.IndexParam); err != nil {
			return err
		}
	}
	// guid {in} (1:{alias=REFGUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID != nil {
			if err := o.GUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetParamCustomDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// indexFunc {in} (1:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.IndexFunc); err != nil {
			return err
		}
	}
	// indexParam {in} (1:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.IndexParam); err != nil {
			return err
		}
	}
	// guid {in} (1:{alias=REFGUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID == nil {
			o.GUID = &dtyp.GUID{}
		}
		if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetParamCustomDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetParamCustomDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pVarVal {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.VarValue != nil {
			_ptr_pVarVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.VarValue != nil {
					if err := o.VarValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.VarValue, _ptr_pVarVal); err != nil {
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

func (o *xxx_GetParamCustomDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pVarVal {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pVarVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.VarValue == nil {
				o.VarValue = &oaut.Variant{}
			}
			if err := o.VarValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pVarVal := func(ptr interface{}) { o.VarValue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.VarValue, _s_pVarVal, _ptr_pVarVal); err != nil {
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

// GetParamCustomDataRequest structure represents the GetParamCustData operation request
type GetParamCustomDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// indexFunc: MUST specify an ordinal position in the method table and MUST be less
	// than the value of the cFuncs field in the TYPEATTR structure of the associated type,
	// as specified in sections 2.2.44 and 3.7.4.1.
	IndexFunc uint32 `idl:"name:indexFunc" json:"index_func"`
	// indexParam: MUST specify an ordinal position in the parameter table of the method
	// specified by indexFunc. The value of indexParam MUST be less than the value of the
	// cParams field in the FUNCDESC structure of the associated method, as specified in
	// sections 2.2.42 and 3.7.4.3.
	IndexParam uint32 `idl:"name:indexParam" json:"index_param"`
	// guid: MUST be the GUID associated with the custom data item using the [custom] attribute,
	// as specified in section 2.2.49.6.
	GUID *dtyp.GUID `idl:"name:guid" json:"guid"`
}

func (o *GetParamCustomDataRequest) xxx_ToOp(ctx context.Context) *xxx_GetParamCustomDataOperation {
	if o == nil {
		return &xxx_GetParamCustomDataOperation{}
	}
	return &xxx_GetParamCustomDataOperation{
		This:       o.This,
		IndexFunc:  o.IndexFunc,
		IndexParam: o.IndexParam,
		GUID:       o.GUID,
	}
}

func (o *GetParamCustomDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetParamCustomDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.IndexFunc = op.IndexFunc
	o.IndexParam = op.IndexParam
	o.GUID = op.GUID
}
func (o *GetParamCustomDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetParamCustomDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetParamCustomDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetParamCustomDataResponse structure represents the GetParamCustData operation response
type GetParamCustomDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pVarVal: MUST be set to the value of the custom data item, or to VT_EMPTY if the
	// parameter does not have an associated custom data item.
	VarValue *oaut.Variant `idl:"name:pVarVal" json:"var_value"`
	// Return: The GetParamCustData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetParamCustomDataResponse) xxx_ToOp(ctx context.Context) *xxx_GetParamCustomDataOperation {
	if o == nil {
		return &xxx_GetParamCustomDataOperation{}
	}
	return &xxx_GetParamCustomDataOperation{
		That:     o.That,
		VarValue: o.VarValue,
		Return:   o.Return,
	}
}

func (o *GetParamCustomDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetParamCustomDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.VarValue = op.VarValue
	o.Return = op.Return
}
func (o *GetParamCustomDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetParamCustomDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetParamCustomDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetVarCustomDataOperation structure represents the GetVarCustData operation
type xxx_GetVarCustomDataOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Index    uint32         `idl:"name:index" json:"index"`
	GUID     *dtyp.GUID     `idl:"name:guid" json:"guid"`
	VarValue *oaut.Variant  `idl:"name:pVarVal" json:"var_value"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetVarCustomDataOperation) OpNum() int { return 29 }

func (o *xxx_GetVarCustomDataOperation) OpName() string { return "/ITypeInfo2/v0/GetVarCustData" }

func (o *xxx_GetVarCustomDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVarCustomDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// guid {in} (1:{alias=REFGUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID != nil {
			if err := o.GUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetVarCustomDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// guid {in} (1:{alias=REFGUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID == nil {
			o.GUID = &dtyp.GUID{}
		}
		if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVarCustomDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVarCustomDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pVarVal {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.VarValue != nil {
			_ptr_pVarVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.VarValue != nil {
					if err := o.VarValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.VarValue, _ptr_pVarVal); err != nil {
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

func (o *xxx_GetVarCustomDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pVarVal {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pVarVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.VarValue == nil {
				o.VarValue = &oaut.Variant{}
			}
			if err := o.VarValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pVarVal := func(ptr interface{}) { o.VarValue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.VarValue, _s_pVarVal, _ptr_pVarVal); err != nil {
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

// GetVarCustomDataRequest structure represents the GetVarCustData operation request
type GetVarCustomDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// index: MUST specify an ordinal position in the data member table and MUST be less
	// than the value of the cVars field in the TYPEATTR structure of the associated type,
	// as specified in sections 2.2.44 and 3.7.4.1.
	Index uint32 `idl:"name:index" json:"index"`
	// guid: MUST be the GUID associated with the custom data item using the [custom] attribute,
	// as specified in section 2.2.49.5.
	GUID *dtyp.GUID `idl:"name:guid" json:"guid"`
}

func (o *GetVarCustomDataRequest) xxx_ToOp(ctx context.Context) *xxx_GetVarCustomDataOperation {
	if o == nil {
		return &xxx_GetVarCustomDataOperation{}
	}
	return &xxx_GetVarCustomDataOperation{
		This:  o.This,
		Index: o.Index,
		GUID:  o.GUID,
	}
}

func (o *GetVarCustomDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetVarCustomDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
	o.GUID = op.GUID
}
func (o *GetVarCustomDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetVarCustomDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVarCustomDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetVarCustomDataResponse structure represents the GetVarCustData operation response
type GetVarCustomDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pVarVal: MUST be set to the value of the custom data item, or to VT_EMPTY if the
	// type does not have a value associated with the GUID.
	VarValue *oaut.Variant `idl:"name:pVarVal" json:"var_value"`
	// Return: The GetVarCustData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetVarCustomDataResponse) xxx_ToOp(ctx context.Context) *xxx_GetVarCustomDataOperation {
	if o == nil {
		return &xxx_GetVarCustomDataOperation{}
	}
	return &xxx_GetVarCustomDataOperation{
		That:     o.That,
		VarValue: o.VarValue,
		Return:   o.Return,
	}
}

func (o *GetVarCustomDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetVarCustomDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.VarValue = op.VarValue
	o.Return = op.Return
}
func (o *GetVarCustomDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetVarCustomDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVarCustomDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetImplTypeCustomDataOperation structure represents the GetImplTypeCustData operation
type xxx_GetImplTypeCustomDataOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Index    uint32         `idl:"name:index" json:"index"`
	GUID     *dtyp.GUID     `idl:"name:guid" json:"guid"`
	VarValue *oaut.Variant  `idl:"name:pVarVal" json:"var_value"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetImplTypeCustomDataOperation) OpNum() int { return 30 }

func (o *xxx_GetImplTypeCustomDataOperation) OpName() string {
	return "/ITypeInfo2/v0/GetImplTypeCustData"
}

func (o *xxx_GetImplTypeCustomDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetImplTypeCustomDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// guid {in} (1:{alias=REFGUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID != nil {
			if err := o.GUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetImplTypeCustomDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// guid {in} (1:{alias=REFGUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID == nil {
			o.GUID = &dtyp.GUID{}
		}
		if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetImplTypeCustomDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetImplTypeCustomDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pVarVal {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.VarValue != nil {
			_ptr_pVarVal := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.VarValue != nil {
					if err := o.VarValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.VarValue, _ptr_pVarVal); err != nil {
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

func (o *xxx_GetImplTypeCustomDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pVarVal {out} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pVarVal := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.VarValue == nil {
				o.VarValue = &oaut.Variant{}
			}
			if err := o.VarValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pVarVal := func(ptr interface{}) { o.VarValue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.VarValue, _s_pVarVal, _ptr_pVarVal); err != nil {
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

// GetImplTypeCustomDataRequest structure represents the GetImplTypeCustData operation request
type GetImplTypeCustomDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// index: MUST specify an ordinal position in the interface table and MUST be less than
	// the value of the cImplTypes field in the TYPEATTR structure of the associated type,
	// as specified in sections 2.2.44 and 3.7.4.1.
	Index uint32 `idl:"name:index" json:"index"`
	// guid: MUST be the GUID associated with the custom data item using the [custom] attribute,
	// as specified in section 2.2.49.8.
	GUID *dtyp.GUID `idl:"name:guid" json:"guid"`
}

func (o *GetImplTypeCustomDataRequest) xxx_ToOp(ctx context.Context) *xxx_GetImplTypeCustomDataOperation {
	if o == nil {
		return &xxx_GetImplTypeCustomDataOperation{}
	}
	return &xxx_GetImplTypeCustomDataOperation{
		This:  o.This,
		Index: o.Index,
		GUID:  o.GUID,
	}
}

func (o *GetImplTypeCustomDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetImplTypeCustomDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
	o.GUID = op.GUID
}
func (o *GetImplTypeCustomDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetImplTypeCustomDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetImplTypeCustomDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetImplTypeCustomDataResponse structure represents the GetImplTypeCustData operation response
type GetImplTypeCustomDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pVarVal: MUST be set to the value of the custom data item, or to VT_EMPTY if the
	// type does not have a value associated with the GUID.
	VarValue *oaut.Variant `idl:"name:pVarVal" json:"var_value"`
	// Return: The GetImplTypeCustData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetImplTypeCustomDataResponse) xxx_ToOp(ctx context.Context) *xxx_GetImplTypeCustomDataOperation {
	if o == nil {
		return &xxx_GetImplTypeCustomDataOperation{}
	}
	return &xxx_GetImplTypeCustomDataOperation{
		That:     o.That,
		VarValue: o.VarValue,
		Return:   o.Return,
	}
}

func (o *GetImplTypeCustomDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetImplTypeCustomDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.VarValue = op.VarValue
	o.Return = op.Return
}
func (o *GetImplTypeCustomDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetImplTypeCustomDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetImplTypeCustomDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDocumentation2Operation structure represents the GetDocumentation2 operation
type xxx_GetDocumentation2Operation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	MemberID          int32          `idl:"name:memid" json:"member_id"`
	LocaleID          uint32         `idl:"name:lcid" json:"locale_id"`
	PointerFlags      uint32         `idl:"name:refPtrFlags" json:"pointer_flags"`
	HelpString        *oaut.String   `idl:"name:pbstrHelpString" json:"help_string"`
	HelpStringContext uint32         `idl:"name:pdwHelpStringContext" json:"help_string_context"`
	HelpStringDLL     *oaut.String   `idl:"name:pbstrHelpStringDll" json:"help_string_dll"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDocumentation2Operation) OpNum() int { return 31 }

func (o *xxx_GetDocumentation2Operation) OpName() string { return "/ITypeInfo2/v0/GetDocumentation2" }

func (o *xxx_GetDocumentation2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDocumentation2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lcid {in} (1:{alias=LCID, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.LocaleID); err != nil {
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

func (o *xxx_GetDocumentation2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lcid {in} (1:{alias=LCID, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LocaleID); err != nil {
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

func (o *xxx_GetDocumentation2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDocumentation2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrHelpString {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.HelpString != nil {
			_ptr_pbstrHelpString := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.HelpString != nil {
					if err := o.HelpString.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.HelpString, _ptr_pbstrHelpString); err != nil {
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
	// pdwHelpStringContext {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.HelpStringContext); err != nil {
			return err
		}
	}
	// pbstrHelpStringDll {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.HelpStringDLL != nil {
			_ptr_pbstrHelpStringDll := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.HelpStringDLL != nil {
					if err := o.HelpStringDLL.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.HelpStringDLL, _ptr_pbstrHelpStringDll); err != nil {
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

func (o *xxx_GetDocumentation2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrHelpString {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrHelpString := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.HelpString == nil {
				o.HelpString = &oaut.String{}
			}
			if err := o.HelpString.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrHelpString := func(ptr interface{}) { o.HelpString = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.HelpString, _s_pbstrHelpString, _ptr_pbstrHelpString); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwHelpStringContext {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.HelpStringContext); err != nil {
			return err
		}
	}
	// pbstrHelpStringDll {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrHelpStringDll := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.HelpStringDLL == nil {
				o.HelpStringDLL = &oaut.String{}
			}
			if err := o.HelpStringDLL.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrHelpStringDll := func(ptr interface{}) { o.HelpStringDLL = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.HelpStringDLL, _s_pbstrHelpStringDll, _ptr_pbstrHelpStringDll); err != nil {
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

// GetDocumentation2Request structure represents the GetDocumentation2 operation request
type GetDocumentation2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// memid: MUST be the MEMBERID of a member of the type (as specified in section 2.2.35),
	// or MEMBERID_NIL.
	MemberID int32 `idl:"name:memid" json:"member_id"`
	// lcid: MUST be the locale ID of the specified type or type library.
	//
	// lcid: MUST be the Locale ID associated with the specified type member.
	LocaleID uint32 `idl:"name:lcid" json:"locale_id"`
	// refPtrFlags: MUST be 0, or a combination of the bit flags specified in the following
	// table.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|                                   |                                                                                  |
	//	|               VALUE               |                                     MEANING                                      |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| TYPELIB_HelpStringArg 0x00000001  | MUST specify that the client is interested in the actual pBstrHelpString [out]   |
	//	|                                   | argument.                                                                        |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| TYPELIB_HelpContextArg 0x00000002 | MUST specify that the client is interested in the actual pdwHelpStringContext    |
	//	|                                   | [out] argument.                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| TYPELIB_HelpFileArg 0x00000004    | MUST specify that the client is interested in the actual pBstrHelpStringDll      |
	//	|                                   | [out] argument.                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	PointerFlags uint32 `idl:"name:refPtrFlags" json:"pointer_flags"`
}

func (o *GetDocumentation2Request) xxx_ToOp(ctx context.Context) *xxx_GetDocumentation2Operation {
	if o == nil {
		return &xxx_GetDocumentation2Operation{}
	}
	return &xxx_GetDocumentation2Operation{
		This:         o.This,
		MemberID:     o.MemberID,
		LocaleID:     o.LocaleID,
		PointerFlags: o.PointerFlags,
	}
}

func (o *GetDocumentation2Request) xxx_FromOp(ctx context.Context, op *xxx_GetDocumentation2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MemberID = op.MemberID
	o.LocaleID = op.LocaleID
	o.PointerFlags = op.PointerFlags
}
func (o *GetDocumentation2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDocumentation2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDocumentation2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDocumentation2Response structure represents the GetDocumentation2 operation response
type GetDocumentation2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbstrHelpString: MUST be set to an implementation-specific BSTR type<63> if the TYPELIB_HelpStringArg
	// bit flag is set in refPtrFlags. MUST be set to a NULL BSTR otherwise.
	//
	// pbstrHelpString: If the TYPEINFO_HelpStringContextArg and TYPEINFO_HelpStringDllArg
	// bit flags are set in refPtrFlags, pbstrHelpString MUST be set to an implementation-specific
	// BSTR<60> . Otherwise, MUST be set to a NULL BSTR.
	HelpString *oaut.String `idl:"name:pbstrHelpString" json:"help_string"`
	// pdwHelpStringContext: MUST be set to the value that was associated with the specified
	// type or type library using the [helpstringcontext] attribute (see section 2.2.49.2)
	// if the TYPELIB_HelpContextArg bit flag is set in refPtrFlags. MUST be set to 0 otherwise.
	//
	// pdwHelpStringContext: MUST be set to the value that was associated with the specified
	// type or type member using the [helpstringcontext] attribute (see IDL Automation Scope)
	// if the TYPEINFO_HelpStringContextArg bit flag is set in refPtrFlags. MUST be set
	// to 0 otherwise.
	HelpStringContext uint32 `idl:"name:pdwHelpStringContext" json:"help_string_context"`
	// pbstrHelpStringDll: MUST be set to the documentation string that was associated with
	// the specified type or type library using the [helpstringdll] attribute (see section
	// 2.2.49.2) if the TYPELIB_HelpFileArg bit flag is set in refPtrFlags. MUST be set
	// to a NULL BSTR otherwise.
	//
	// pbstrHelpStringDll: MUST be set to the documentation string that was associated with
	// the specified type or type member using the [helpstringdll] attribute (see IDL Automation
	// Scope) if the TYPEINFO_HelpStringDllArg bit flag is set in refPtrFlags. MUST be set
	// to a NULL BSTR otherwise.
	HelpStringDLL *oaut.String `idl:"name:pbstrHelpStringDll" json:"help_string_dll"`
	// Return: The GetDocumentation2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDocumentation2Response) xxx_ToOp(ctx context.Context) *xxx_GetDocumentation2Operation {
	if o == nil {
		return &xxx_GetDocumentation2Operation{}
	}
	return &xxx_GetDocumentation2Operation{
		That:              o.That,
		HelpString:        o.HelpString,
		HelpStringContext: o.HelpStringContext,
		HelpStringDLL:     o.HelpStringDLL,
		Return:            o.Return,
	}
}

func (o *GetDocumentation2Response) xxx_FromOp(ctx context.Context, op *xxx_GetDocumentation2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.HelpString = op.HelpString
	o.HelpStringContext = op.HelpStringContext
	o.HelpStringDLL = op.HelpStringDLL
	o.Return = op.Return
}
func (o *GetDocumentation2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDocumentation2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDocumentation2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAllCustomDataOperation structure represents the GetAllCustData operation
type xxx_GetAllCustomDataOperation struct {
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat   `idl:"name:That" json:"that"`
	CustomData *oaut.CustomData `idl:"name:pCustData" json:"custom_data"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAllCustomDataOperation) OpNum() int { return 32 }

func (o *xxx_GetAllCustomDataOperation) OpName() string { return "/ITypeInfo2/v0/GetAllCustData" }

func (o *xxx_GetAllCustomDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllCustomDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAllCustomDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAllCustomDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllCustomDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pCustData {out} (1:{pointer=ref}*(1))(2:{alias=CUSTDATA}(struct))
	{
		if o.CustomData != nil {
			if err := o.CustomData.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.CustomData{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetAllCustomDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pCustData {out} (1:{pointer=ref}*(1))(2:{alias=CUSTDATA}(struct))
	{
		if o.CustomData == nil {
			o.CustomData = &oaut.CustomData{}
		}
		if err := o.CustomData.UnmarshalNDR(ctx, w); err != nil {
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

// GetAllCustomDataRequest structure represents the GetAllCustData operation request
type GetAllCustomDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAllCustomDataRequest) xxx_ToOp(ctx context.Context) *xxx_GetAllCustomDataOperation {
	if o == nil {
		return &xxx_GetAllCustomDataOperation{}
	}
	return &xxx_GetAllCustomDataOperation{
		This: o.This,
	}
}

func (o *GetAllCustomDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAllCustomDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAllCustomDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetAllCustomDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllCustomDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAllCustomDataResponse structure represents the GetAllCustData operation response
type GetAllCustomDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pCustData: MUST be set to a CUSTDATA structure that contains an array of custom data
	// items, as specified in section 2.2.47. The structure's cCustData field MUST be set
	// to 0 and its prgCustData field set to NULL if there are no custom data items associated
	// with the automation type library.
	//
	// pCustData: MUST be set to a CUSTDATA structure that contains an array of custom data
	// items, as specified in section 2.2.47. The structure's cCustData field MUST be set
	// to 0 and its prgCustData field set to NULL, if there are no custom data items associated
	// with the automation type description.
	CustomData *oaut.CustomData `idl:"name:pCustData" json:"custom_data"`
	// Return: The GetAllCustData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAllCustomDataResponse) xxx_ToOp(ctx context.Context) *xxx_GetAllCustomDataOperation {
	if o == nil {
		return &xxx_GetAllCustomDataOperation{}
	}
	return &xxx_GetAllCustomDataOperation{
		That:       o.That,
		CustomData: o.CustomData,
		Return:     o.Return,
	}
}

func (o *GetAllCustomDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAllCustomDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CustomData = op.CustomData
	o.Return = op.Return
}
func (o *GetAllCustomDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetAllCustomDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllCustomDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAllFuncCustomDataOperation structure represents the GetAllFuncCustData operation
type xxx_GetAllFuncCustomDataOperation struct {
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Index      uint32           `idl:"name:index" json:"index"`
	CustomData *oaut.CustomData `idl:"name:pCustData" json:"custom_data"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAllFuncCustomDataOperation) OpNum() int { return 33 }

func (o *xxx_GetAllFuncCustomDataOperation) OpName() string {
	return "/ITypeInfo2/v0/GetAllFuncCustData"
}

func (o *xxx_GetAllFuncCustomDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllFuncCustomDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAllFuncCustomDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAllFuncCustomDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllFuncCustomDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pCustData {out} (1:{pointer=ref}*(1))(2:{alias=CUSTDATA}(struct))
	{
		if o.CustomData != nil {
			if err := o.CustomData.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.CustomData{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetAllFuncCustomDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pCustData {out} (1:{pointer=ref}*(1))(2:{alias=CUSTDATA}(struct))
	{
		if o.CustomData == nil {
			o.CustomData = &oaut.CustomData{}
		}
		if err := o.CustomData.UnmarshalNDR(ctx, w); err != nil {
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

// GetAllFuncCustomDataRequest structure represents the GetAllFuncCustData operation request
type GetAllFuncCustomDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// index: MUST specify an ordinal position in the method table and MUST be less than
	// the value of the cFuncs field in the TYPEATTR structure of the associated type, as
	// specified in sections 2.2.44 and 3.7.4.1.
	Index uint32 `idl:"name:index" json:"index"`
}

func (o *GetAllFuncCustomDataRequest) xxx_ToOp(ctx context.Context) *xxx_GetAllFuncCustomDataOperation {
	if o == nil {
		return &xxx_GetAllFuncCustomDataOperation{}
	}
	return &xxx_GetAllFuncCustomDataOperation{
		This:  o.This,
		Index: o.Index,
	}
}

func (o *GetAllFuncCustomDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAllFuncCustomDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
}
func (o *GetAllFuncCustomDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetAllFuncCustomDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllFuncCustomDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAllFuncCustomDataResponse structure represents the GetAllFuncCustData operation response
type GetAllFuncCustomDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pCustData: MUST be set to a CUSTDATA structure that contains an array of custom data
	// items, as specified in section 2.2.49.5.1. The structure's cCustData field MUST be
	// set to 0 and its prgCustData field set to NULL, if there are no custom data items
	// associated with the method.
	CustomData *oaut.CustomData `idl:"name:pCustData" json:"custom_data"`
	// Return: The GetAllFuncCustData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAllFuncCustomDataResponse) xxx_ToOp(ctx context.Context) *xxx_GetAllFuncCustomDataOperation {
	if o == nil {
		return &xxx_GetAllFuncCustomDataOperation{}
	}
	return &xxx_GetAllFuncCustomDataOperation{
		That:       o.That,
		CustomData: o.CustomData,
		Return:     o.Return,
	}
}

func (o *GetAllFuncCustomDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAllFuncCustomDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CustomData = op.CustomData
	o.Return = op.Return
}
func (o *GetAllFuncCustomDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetAllFuncCustomDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllFuncCustomDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAllParamCustomDataOperation structure represents the GetAllParamCustData operation
type xxx_GetAllParamCustomDataOperation struct {
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat   `idl:"name:That" json:"that"`
	IndexFunc  uint32           `idl:"name:indexFunc" json:"index_func"`
	IndexParam uint32           `idl:"name:indexParam" json:"index_param"`
	CustomData *oaut.CustomData `idl:"name:pCustData" json:"custom_data"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAllParamCustomDataOperation) OpNum() int { return 34 }

func (o *xxx_GetAllParamCustomDataOperation) OpName() string {
	return "/ITypeInfo2/v0/GetAllParamCustData"
}

func (o *xxx_GetAllParamCustomDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllParamCustomDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// indexFunc {in} (1:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.IndexFunc); err != nil {
			return err
		}
	}
	// indexParam {in} (1:{alias=UINT}(uint32))
	{
		if err := w.WriteData(o.IndexParam); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllParamCustomDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// indexFunc {in} (1:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.IndexFunc); err != nil {
			return err
		}
	}
	// indexParam {in} (1:{alias=UINT}(uint32))
	{
		if err := w.ReadData(&o.IndexParam); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllParamCustomDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllParamCustomDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pCustData {out} (1:{pointer=ref}*(1))(2:{alias=CUSTDATA}(struct))
	{
		if o.CustomData != nil {
			if err := o.CustomData.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.CustomData{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetAllParamCustomDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pCustData {out} (1:{pointer=ref}*(1))(2:{alias=CUSTDATA}(struct))
	{
		if o.CustomData == nil {
			o.CustomData = &oaut.CustomData{}
		}
		if err := o.CustomData.UnmarshalNDR(ctx, w); err != nil {
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

// GetAllParamCustomDataRequest structure represents the GetAllParamCustData operation request
type GetAllParamCustomDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// indexFunc: MUST specify an ordinal position in the method table and MUST be less
	// than the value of the cFuncs field in the TYPEATTR structure of the associated type,
	// as specified in sections 2.2.44 and 3.7.4.1.
	IndexFunc uint32 `idl:"name:indexFunc" json:"index_func"`
	// indexParam: MUST specify an ordinal position in the parameter table of the method
	// specified by indexFunc. The value of indexParam MUST be less than the value of the
	// cParams field in the FUNCDESC structure of the associated method, as specified in
	// sections 2.2.42 and 3.7.4.3.
	IndexParam uint32 `idl:"name:indexParam" json:"index_param"`
}

func (o *GetAllParamCustomDataRequest) xxx_ToOp(ctx context.Context) *xxx_GetAllParamCustomDataOperation {
	if o == nil {
		return &xxx_GetAllParamCustomDataOperation{}
	}
	return &xxx_GetAllParamCustomDataOperation{
		This:       o.This,
		IndexFunc:  o.IndexFunc,
		IndexParam: o.IndexParam,
	}
}

func (o *GetAllParamCustomDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAllParamCustomDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.IndexFunc = op.IndexFunc
	o.IndexParam = op.IndexParam
}
func (o *GetAllParamCustomDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetAllParamCustomDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllParamCustomDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAllParamCustomDataResponse structure represents the GetAllParamCustData operation response
type GetAllParamCustomDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pCustData: MUST be set to a CUSTDATA structure that contains an array of custom data
	// items, as specified in section 2.2.49.6. The structure's cCustData field MUST be
	// set to 0 and its prgCustData field set to NULL, if there are no custom data items
	// associated with the parameter.
	CustomData *oaut.CustomData `idl:"name:pCustData" json:"custom_data"`
	// Return: The GetAllParamCustData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAllParamCustomDataResponse) xxx_ToOp(ctx context.Context) *xxx_GetAllParamCustomDataOperation {
	if o == nil {
		return &xxx_GetAllParamCustomDataOperation{}
	}
	return &xxx_GetAllParamCustomDataOperation{
		That:       o.That,
		CustomData: o.CustomData,
		Return:     o.Return,
	}
}

func (o *GetAllParamCustomDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAllParamCustomDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CustomData = op.CustomData
	o.Return = op.Return
}
func (o *GetAllParamCustomDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetAllParamCustomDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllParamCustomDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAllVarCustomDataOperation structure represents the GetAllVarCustData operation
type xxx_GetAllVarCustomDataOperation struct {
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Index      uint32           `idl:"name:index" json:"index"`
	CustomData *oaut.CustomData `idl:"name:pCustData" json:"custom_data"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAllVarCustomDataOperation) OpNum() int { return 35 }

func (o *xxx_GetAllVarCustomDataOperation) OpName() string { return "/ITypeInfo2/v0/GetAllVarCustData" }

func (o *xxx_GetAllVarCustomDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllVarCustomDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAllVarCustomDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAllVarCustomDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllVarCustomDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pCustData {out} (1:{pointer=ref}*(1))(2:{alias=CUSTDATA}(struct))
	{
		if o.CustomData != nil {
			if err := o.CustomData.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.CustomData{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetAllVarCustomDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pCustData {out} (1:{pointer=ref}*(1))(2:{alias=CUSTDATA}(struct))
	{
		if o.CustomData == nil {
			o.CustomData = &oaut.CustomData{}
		}
		if err := o.CustomData.UnmarshalNDR(ctx, w); err != nil {
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

// GetAllVarCustomDataRequest structure represents the GetAllVarCustData operation request
type GetAllVarCustomDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// index: MUST specify an ordinal position in the data member table and MUST be less
	// than the value of the cVars field in the TYPEATTR structure of the associated type,
	// as specified in sections 2.2.44 and 3.7.4.1.
	Index uint32 `idl:"name:index" json:"index"`
}

func (o *GetAllVarCustomDataRequest) xxx_ToOp(ctx context.Context) *xxx_GetAllVarCustomDataOperation {
	if o == nil {
		return &xxx_GetAllVarCustomDataOperation{}
	}
	return &xxx_GetAllVarCustomDataOperation{
		This:  o.This,
		Index: o.Index,
	}
}

func (o *GetAllVarCustomDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAllVarCustomDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
}
func (o *GetAllVarCustomDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetAllVarCustomDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllVarCustomDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAllVarCustomDataResponse structure represents the GetAllVarCustData operation response
type GetAllVarCustomDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pCustData: MUST be set to a CUSTDATA structure that contains an array of custom data
	// items, as specified in section 2.2.49.5. The structure's cCustData field MUST be
	// set to 0 and its prgCustData field set to NULL, if there are no custom data items
	// associated with the data member.
	CustomData *oaut.CustomData `idl:"name:pCustData" json:"custom_data"`
	// Return: The GetAllVarCustData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAllVarCustomDataResponse) xxx_ToOp(ctx context.Context) *xxx_GetAllVarCustomDataOperation {
	if o == nil {
		return &xxx_GetAllVarCustomDataOperation{}
	}
	return &xxx_GetAllVarCustomDataOperation{
		That:       o.That,
		CustomData: o.CustomData,
		Return:     o.Return,
	}
}

func (o *GetAllVarCustomDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAllVarCustomDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CustomData = op.CustomData
	o.Return = op.Return
}
func (o *GetAllVarCustomDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetAllVarCustomDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllVarCustomDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAllImplTypeCustomDataOperation structure represents the GetAllImplTypeCustData operation
type xxx_GetAllImplTypeCustomDataOperation struct {
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Index      uint32           `idl:"name:index" json:"index"`
	CustomData *oaut.CustomData `idl:"name:pCustData" json:"custom_data"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAllImplTypeCustomDataOperation) OpNum() int { return 36 }

func (o *xxx_GetAllImplTypeCustomDataOperation) OpName() string {
	return "/ITypeInfo2/v0/GetAllImplTypeCustData"
}

func (o *xxx_GetAllImplTypeCustomDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllImplTypeCustomDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAllImplTypeCustomDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAllImplTypeCustomDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllImplTypeCustomDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pCustData {out} (1:{pointer=ref}*(1))(2:{alias=CUSTDATA}(struct))
	{
		if o.CustomData != nil {
			if err := o.CustomData.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oaut.CustomData{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetAllImplTypeCustomDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pCustData {out} (1:{pointer=ref}*(1))(2:{alias=CUSTDATA}(struct))
	{
		if o.CustomData == nil {
			o.CustomData = &oaut.CustomData{}
		}
		if err := o.CustomData.UnmarshalNDR(ctx, w); err != nil {
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

// GetAllImplTypeCustomDataRequest structure represents the GetAllImplTypeCustData operation request
type GetAllImplTypeCustomDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// index: MUST specify an ordinal position in the interface table and MUST be less than
	// the value of the cImplTypes field in the TYPEATTR structure of the associated type,
	// as specified in sections 2.2.44 and 3.7.4.1.
	Index uint32 `idl:"name:index" json:"index"`
}

func (o *GetAllImplTypeCustomDataRequest) xxx_ToOp(ctx context.Context) *xxx_GetAllImplTypeCustomDataOperation {
	if o == nil {
		return &xxx_GetAllImplTypeCustomDataOperation{}
	}
	return &xxx_GetAllImplTypeCustomDataOperation{
		This:  o.This,
		Index: o.Index,
	}
}

func (o *GetAllImplTypeCustomDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAllImplTypeCustomDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Index = op.Index
}
func (o *GetAllImplTypeCustomDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetAllImplTypeCustomDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllImplTypeCustomDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAllImplTypeCustomDataResponse structure represents the GetAllImplTypeCustData operation response
type GetAllImplTypeCustomDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pCustData: MUST be set to a CUSTDATA structure that contains an array of custom data
	// items, as specified in section 2.2.49.8. The structure's cCustData field MUST be
	// set to 0 and its prgCustData field set to NULL if there are no custom data items
	// associated with the interface.
	CustomData *oaut.CustomData `idl:"name:pCustData" json:"custom_data"`
	// Return: The GetAllImplTypeCustData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAllImplTypeCustomDataResponse) xxx_ToOp(ctx context.Context) *xxx_GetAllImplTypeCustomDataOperation {
	if o == nil {
		return &xxx_GetAllImplTypeCustomDataOperation{}
	}
	return &xxx_GetAllImplTypeCustomDataOperation{
		That:       o.That,
		CustomData: o.CustomData,
		Return:     o.Return,
	}
}

func (o *GetAllImplTypeCustomDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAllImplTypeCustomDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CustomData = op.CustomData
	o.Return = op.Return
}
func (o *GetAllImplTypeCustomDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetAllImplTypeCustomDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllImplTypeCustomDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
