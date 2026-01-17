package itypelib2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
		op := &xxx_GetCustomDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCustomDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCustomData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // GetLibStatistics
		op := &xxx_GetLibStatisticsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLibStatisticsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLibStatistics(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // GetDocumentation2
		op := &xxx_GetDocumentation2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDocumentation2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDocumentation2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // GetAllCustData
		op := &xxx_GetAllCustomDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAllCustomDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAllCustomData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ITypeLib2
type UnimplementedTypeLib2Server struct {
	itypelib.UnimplementedTypeLibServer
}

func (UnimplementedTypeLib2Server) GetCustomData(context.Context, *GetCustomDataRequest) (*GetCustomDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeLib2Server) GetLibStatistics(context.Context, *GetLibStatisticsRequest) (*GetLibStatisticsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeLib2Server) GetDocumentation2(context.Context, *GetDocumentation2Request) (*GetDocumentation2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTypeLib2Server) GetAllCustomData(context.Context, *GetAllCustomDataRequest) (*GetAllCustomDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ TypeLib2Server = (*UnimplementedTypeLib2Server)(nil)
