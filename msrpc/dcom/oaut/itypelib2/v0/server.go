package itypelib2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	itypelib "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/itypelib/v0"
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
	_ = itypelib.GoPackage
)

// ITypeLib2 server interface.
type TypeLib2Server interface {

	// ITypeLib base class.
	itypelib.TypeLibServer

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

	// The GetLibStatistics method returns statistics about the unique names in the automation
	// type library.
	//
	// Return Values: The method MUST return information in an HRESULT data structure, defined
	// in [MS-ERREF] section 2.1. The severity bit in the structure identifies the following
	// conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	GetLibStatistics(context.Context, *GetLibStatisticsRequest) (*GetLibStatisticsResponse, error)

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
}

func RegisterTypeLib2Server(conn dcerpc.Conn, o TypeLib2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTypeLib2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TypeLib2SyntaxV0_0))...)
}

func NewTypeLib2ServerHandle(o TypeLib2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TypeLib2ServerHandle(ctx, o, opNum, r)
	}
}

func TypeLib2ServerHandle(ctx context.Context, o TypeLib2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 13 {
		// ITypeLib base method.
		return itypelib.TypeLibServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 13: // GetCustData
		in := &GetCustomDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCustomData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // GetLibStatistics
		in := &GetLibStatisticsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLibStatistics(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // GetDocumentation2
		in := &GetDocumentation2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDocumentation2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // GetAllCustData
		in := &GetAllCustomDataRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAllCustomData(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
