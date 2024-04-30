package itypelib

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

// ITypeLib server interface.
type TypeLibServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GetTypeInfoCount method provides the number of automation type descriptions in
	// the type information table.
	//
	// The GetTypeInfoCount method specifies whether the automation server provides Type
	// description information.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	GetTypeInfoCount(context.Context, *GetTypeInfoCountRequest) (*GetTypeInfoCountResponse, error)

	// The GetTypeInfo method provides access to the Type description information exposed
	// by the automation server.
	//
	// The GetTypeInfo method retrieves the automation type description that has the specified
	// ordinal position within the type information table.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1 and the entire HRESULT ( ../ms-dtyp/a9046ed2-bfb2-4d56-a719-2824afce59ac
	// ) DWORD does not match the value in the following table, a fatal failure occurred.
	//
	// * If the severity bit is set to 1 and the entire HRESULT DWORD matches the value
	// in the following table, a failure occurred.
	//
	// Return value/code
	//
	// # Description
	//
	// 0x8002000B
	//
	// DISP_E_BADINDEX
	//
	// SHOULD be returned when the value of the passed in iTInfo argument was not 0. See
	// [MS-ERREF].
	GetTypeInfo(context.Context, *GetTypeInfoRequest) (*GetTypeInfoResponse, error)

	// The GetTypeInfoType method retrieves the TYPEKIND value associated with an automation
	// type description.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1 and the entire *HRESULT* DWORD does not match a
	// value in the following table, a fatal failure occurred.
	//
	// * If the severity bit is set to 1 and the entire *HRESULT* DWORD matches a value
	// in the following table, a failure occurred.
	//
	// Return value/code
	//
	// # Description
	//
	// 0x8002802B
	//
	// TYPE_E_ELEMENTNOTFOUND
	//
	// The value of index did not specify the ordinal position of an element in the type
	// information table. See [MS-ERREF].
	GetTypeInfoType(context.Context, *GetTypeInfoTypeRequest) (*GetTypeInfoTypeResponse, error)

	// The GetTypeInfoOfGuid method retrieves the automation type description with the specified
	// GUID from the server's type information table.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1 and the entire *HRESULT* DWORD does not match a
	// value in the following table, a fatal failure occurred.
	//
	// * If the severity bit is set to 1 and the entire *HRESULT* DWORD matches a value
	// in the following table, a failure occurred.
	//
	// Return value/code
	//
	// # Description
	//
	// 0x8002802B
	//
	// TYPE_E_ELEMENTNOTFOUND
	//
	// The value of guid did not correspond to any entry in the Type information table,
	// or the value of guid was IID_NULL. See [MS-ERREF].
	GetTypeInfoOfGUID(context.Context, *GetTypeInfoOfGUIDRequest) (*GetTypeInfoOfGUIDResponse, error)

	// The GetLibAttr method retrieves a structure that contains the attributes declared
	// with the Type library.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	GetLibAttribute(context.Context, *GetLibAttributeRequest) (*GetLibAttributeResponse, error)

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

	// The IsName method indicates whether the specified string matches the name of a type
	// or type member that is contained in the automation type library or its binding context.
	//
	// Return Values: The method MUST return the information in an HRESULT data structure,
	// which is defined in [MS-ERREF] section 2.1. The severity bit in the structure identifies
	// the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	IsName(context.Context, *IsNameRequest) (*IsNameResponse, error)

	// The FindName method retrieves references to types, or type members, contained in
	// the automation type library whose names match a specified string.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	FindName(context.Context, *FindNameRequest) (*FindNameResponse, error)

	// Opnum12NotUsedOnWire operation.
	// Opnum12NotUsedOnWire
}

func RegisterTypeLibServer(conn dcerpc.Conn, o TypeLibServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTypeLibServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TypeLibSyntaxV0_0))...)
}

func NewTypeLibServerHandle(o TypeLibServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TypeLibServerHandle(ctx, o, opNum, r)
	}
}

func TypeLibServerHandle(ctx context.Context, o TypeLibServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetTypeInfoCount
		in := &GetTypeInfoCountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTypeInfoCount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // GetTypeInfo
		in := &GetTypeInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTypeInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // GetTypeInfoType
		in := &GetTypeInfoTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTypeInfoType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // GetTypeInfoOfGuid
		in := &GetTypeInfoOfGUIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTypeInfoOfGUID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // GetLibAttr
		in := &GetLibAttributeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLibAttribute(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // GetTypeComp
		in := &GetTypeCompRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTypeComp(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // GetDocumentation
		in := &GetDocumentationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDocumentation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // IsName
		in := &IsNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.IsName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // FindName
		in := &FindNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FindName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // Opnum12NotUsedOnWire
		// Opnum12NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}
