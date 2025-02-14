package iapphostpropertyexception

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iapphostconfigexception "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostconfigexception/v0"
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
	_ = iapphostconfigexception.GoPackage
)

// IAppHostPropertyException server interface.
type AppHostPropertyExceptionServer interface {

	// IAppHostConfigException base class.
	iapphostconfigexception.AppHostConfigExceptionServer

	// InvalidValue operation.
	GetInvalidValue(context.Context, *GetInvalidValueRequest) (*GetInvalidValueResponse, error)

	// ValidationFailureReason operation.
	GetValidationFailureReason(context.Context, *GetValidationFailureReasonRequest) (*GetValidationFailureReasonResponse, error)

	// ValidationFailureParameters operation.
	GetValidationFailureParameters(context.Context, *GetValidationFailureParametersRequest) (*GetValidationFailureParametersResponse, error)
}

func RegisterAppHostPropertyExceptionServer(conn dcerpc.Conn, o AppHostPropertyExceptionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostPropertyExceptionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostPropertyExceptionSyntaxV0_0))...)
}

func NewAppHostPropertyExceptionServerHandle(o AppHostPropertyExceptionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostPropertyExceptionServerHandle(ctx, o, opNum, r)
	}
}

func AppHostPropertyExceptionServerHandle(ctx context.Context, o AppHostPropertyExceptionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 10 {
		// IAppHostConfigException base method.
		return iapphostconfigexception.AppHostConfigExceptionServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 10: // InvalidValue
		op := &xxx_GetInvalidValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetInvalidValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetInvalidValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // ValidationFailureReason
		op := &xxx_GetValidationFailureReasonOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetValidationFailureReasonRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetValidationFailureReason(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // ValidationFailureParameters
		op := &xxx_GetValidationFailureParametersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetValidationFailureParametersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetValidationFailureParameters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostPropertyException
type UnimplementedAppHostPropertyExceptionServer struct {
	iapphostconfigexception.UnimplementedAppHostConfigExceptionServer
}

func (UnimplementedAppHostPropertyExceptionServer) GetInvalidValue(context.Context, *GetInvalidValueRequest) (*GetInvalidValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertyExceptionServer) GetValidationFailureReason(context.Context, *GetValidationFailureReasonRequest) (*GetValidationFailureReasonResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostPropertyExceptionServer) GetValidationFailureParameters(context.Context, *GetValidationFailureParametersRequest) (*GetValidationFailureParametersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostPropertyExceptionServer = (*UnimplementedAppHostPropertyExceptionServer)(nil)
