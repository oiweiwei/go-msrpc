package itypecomp

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

// ITypeComp server interface.
type TypeCompServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The Bind method retrieves automation type description server instances and related
	// structures that provide access to a method, property or data member whose name corresponds
	// to a specified string.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, which
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
	// 0x8002802C
	//
	// TYPE_E_AMBIGUOUSNAME
	//
	// The values of szName and wFlags match more than one element in the binding context.
	// See [MS-ERREF].
	//
	// 0x80028CA0
	//
	// TYPE_E_TYPEMISMATCH
	//
	// The value of szName matched a single element in the binding context, but the *INVOKEKIND*
	// value that is associated with that element did not match the value of wFlags. See
	// [MS-ERREF].
	Bind(context.Context, *BindRequest) (*BindResponse, error)

	// The BindType method retrieves a reference to an automation type description whose
	// name corresponds to a specified string.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	BindType(context.Context, *BindTypeRequest) (*BindTypeResponse, error)
}

func RegisterTypeCompServer(conn dcerpc.Conn, o TypeCompServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTypeCompServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TypeCompSyntaxV0_0))...)
}

func NewTypeCompServerHandle(o TypeCompServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TypeCompServerHandle(ctx, o, opNum, r)
	}
}

func TypeCompServerHandle(ctx context.Context, o TypeCompServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Bind
		in := &BindRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Bind(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // BindType
		in := &BindTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.BindType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
