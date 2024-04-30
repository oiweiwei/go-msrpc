package iapphostconfigexception

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

// IAppHostConfigException server interface.
type AppHostConfigExceptionServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// LineNumber operation.
	GetLineNumber(context.Context, *GetLineNumberRequest) (*GetLineNumberResponse, error)

	// FileName operation.
	GetFileName(context.Context, *GetFileNameRequest) (*GetFileNameResponse, error)

	// ConfigPath operation.
	GetConfigPath(context.Context, *GetConfigPathRequest) (*GetConfigPathResponse, error)

	// ErrorLine operation.
	GetErrorLine(context.Context, *GetErrorLineRequest) (*GetErrorLineResponse, error)

	// PreErrorLine operation.
	GetPreErrorLine(context.Context, *GetPreErrorLineRequest) (*GetPreErrorLineResponse, error)

	// PostErrorLine operation.
	GetPostErrorLine(context.Context, *GetPostErrorLineRequest) (*GetPostErrorLineResponse, error)

	// ErrorString operation.
	GetErrorString(context.Context, *GetErrorStringRequest) (*GetErrorStringResponse, error)
}

func RegisterAppHostConfigExceptionServer(conn dcerpc.Conn, o AppHostConfigExceptionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostConfigExceptionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostConfigExceptionSyntaxV0_0))...)
}

func NewAppHostConfigExceptionServerHandle(o AppHostConfigExceptionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostConfigExceptionServerHandle(ctx, o, opNum, r)
	}
}

func AppHostConfigExceptionServerHandle(ctx context.Context, o AppHostConfigExceptionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // LineNumber
		in := &GetLineNumberRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLineNumber(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // FileName
		in := &GetFileNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFileName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // ConfigPath
		in := &GetConfigPathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetConfigPath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // ErrorLine
		in := &GetErrorLineRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetErrorLine(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // PreErrorLine
		in := &GetPreErrorLineRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPreErrorLine(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // PostErrorLine
		in := &GetPostErrorLineRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPostErrorLine(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // ErrorString
		in := &GetErrorStringRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetErrorString(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
