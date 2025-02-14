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
		op := &xxx_GetLineNumberOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLineNumberRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLineNumber(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // FileName
		op := &xxx_GetFileNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFileNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFileName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // ConfigPath
		op := &xxx_GetConfigPathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetConfigPathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetConfigPath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // ErrorLine
		op := &xxx_GetErrorLineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetErrorLineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetErrorLine(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // PreErrorLine
		op := &xxx_GetPreErrorLineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPreErrorLineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPreErrorLine(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // PostErrorLine
		op := &xxx_GetPostErrorLineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPostErrorLineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPostErrorLine(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ErrorString
		op := &xxx_GetErrorStringOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetErrorStringRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetErrorString(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IAppHostConfigException
type UnimplementedAppHostConfigExceptionServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedAppHostConfigExceptionServer) GetLineNumber(context.Context, *GetLineNumberRequest) (*GetLineNumberResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigExceptionServer) GetFileName(context.Context, *GetFileNameRequest) (*GetFileNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigExceptionServer) GetConfigPath(context.Context, *GetConfigPathRequest) (*GetConfigPathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigExceptionServer) GetErrorLine(context.Context, *GetErrorLineRequest) (*GetErrorLineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigExceptionServer) GetPreErrorLine(context.Context, *GetPreErrorLineRequest) (*GetPreErrorLineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigExceptionServer) GetPostErrorLine(context.Context, *GetPostErrorLineRequest) (*GetPostErrorLineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAppHostConfigExceptionServer) GetErrorString(context.Context, *GetErrorStringRequest) (*GetErrorStringResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AppHostConfigExceptionServer = (*UnimplementedAppHostConfigExceptionServer)(nil)
