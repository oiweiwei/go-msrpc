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
		in := &GetInvalidValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetInvalidValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // ValidationFailureReason
		in := &GetValidationFailureReasonRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetValidationFailureReason(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // ValidationFailureParameters
		in := &GetValidationFailureParametersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetValidationFailureParameters(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
