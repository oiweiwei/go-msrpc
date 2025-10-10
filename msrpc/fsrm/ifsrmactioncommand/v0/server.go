package ifsrmactioncommand

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	ifsrmaction "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmaction/v0"
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
	_ = ifsrmaction.GoPackage
)

// IFsrmActionCommand server interface.
type ActionCommandServer interface {

	// IFsrmAction base class.
	ifsrmaction.ActionServer

	// ExecutablePath operation.
	GetExecutablePath(context.Context, *GetExecutablePathRequest) (*GetExecutablePathResponse, error)

	// ExecutablePath operation.
	SetExecutablePath(context.Context, *SetExecutablePathRequest) (*SetExecutablePathResponse, error)

	// Arguments operation.
	GetArguments(context.Context, *GetArgumentsRequest) (*GetArgumentsResponse, error)

	// Arguments operation.
	SetArguments(context.Context, *SetArgumentsRequest) (*SetArgumentsResponse, error)

	// Account operation.
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error)

	// Account operation.
	SetAccount(context.Context, *SetAccountRequest) (*SetAccountResponse, error)

	// WorkingDirectory operation.
	GetWorkingDirectory(context.Context, *GetWorkingDirectoryRequest) (*GetWorkingDirectoryResponse, error)

	// WorkingDirectory operation.
	SetWorkingDirectory(context.Context, *SetWorkingDirectoryRequest) (*SetWorkingDirectoryResponse, error)

	// MonitorCommand operation.
	GetMonitorCommand(context.Context, *GetMonitorCommandRequest) (*GetMonitorCommandResponse, error)

	// MonitorCommand operation.
	SetMonitorCommand(context.Context, *SetMonitorCommandRequest) (*SetMonitorCommandResponse, error)

	// KillTimeOut operation.
	GetKillTimeout(context.Context, *GetKillTimeoutRequest) (*GetKillTimeoutResponse, error)

	// KillTimeOut operation.
	SetKillTimeout(context.Context, *SetKillTimeoutRequest) (*SetKillTimeoutResponse, error)

	// LogResult operation.
	GetLogResult(context.Context, *GetLogResultRequest) (*GetLogResultResponse, error)

	// LogResult operation.
	SetLogResult(context.Context, *SetLogResultRequest) (*SetLogResultResponse, error)
}

func RegisterActionCommandServer(conn dcerpc.Conn, o ActionCommandServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewActionCommandServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ActionCommandSyntaxV0_0))...)
}

func NewActionCommandServerHandle(o ActionCommandServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ActionCommandServerHandle(ctx, o, opNum, r)
	}
}

func ActionCommandServerHandle(ctx context.Context, o ActionCommandServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 12 {
		// IFsrmAction base method.
		return ifsrmaction.ActionServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 12: // ExecutablePath
		op := &xxx_GetExecutablePathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetExecutablePathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetExecutablePath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // ExecutablePath
		op := &xxx_SetExecutablePathOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetExecutablePathRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetExecutablePath(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // Arguments
		op := &xxx_GetArgumentsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetArgumentsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetArguments(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // Arguments
		op := &xxx_SetArgumentsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetArgumentsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetArguments(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // Account
		op := &xxx_GetAccountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAccountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAccount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // Account
		op := &xxx_SetAccountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAccountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAccount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // WorkingDirectory
		op := &xxx_GetWorkingDirectoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetWorkingDirectoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetWorkingDirectory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // WorkingDirectory
		op := &xxx_SetWorkingDirectoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetWorkingDirectoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetWorkingDirectory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // MonitorCommand
		op := &xxx_GetMonitorCommandOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMonitorCommandRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMonitorCommand(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // MonitorCommand
		op := &xxx_SetMonitorCommandOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMonitorCommandRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMonitorCommand(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // KillTimeOut
		op := &xxx_GetKillTimeoutOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetKillTimeoutRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetKillTimeout(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // KillTimeOut
		op := &xxx_SetKillTimeoutOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetKillTimeoutRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetKillTimeout(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // LogResult
		op := &xxx_GetLogResultOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLogResultRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLogResult(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // LogResult
		op := &xxx_SetLogResultOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetLogResultRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetLogResult(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFsrmActionCommand
type UnimplementedActionCommandServer struct {
	ifsrmaction.UnimplementedActionServer
}

func (UnimplementedActionCommandServer) GetExecutablePath(context.Context, *GetExecutablePathRequest) (*GetExecutablePathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionCommandServer) SetExecutablePath(context.Context, *SetExecutablePathRequest) (*SetExecutablePathResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionCommandServer) GetArguments(context.Context, *GetArgumentsRequest) (*GetArgumentsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionCommandServer) SetArguments(context.Context, *SetArgumentsRequest) (*SetArgumentsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionCommandServer) GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionCommandServer) SetAccount(context.Context, *SetAccountRequest) (*SetAccountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionCommandServer) GetWorkingDirectory(context.Context, *GetWorkingDirectoryRequest) (*GetWorkingDirectoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionCommandServer) SetWorkingDirectory(context.Context, *SetWorkingDirectoryRequest) (*SetWorkingDirectoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionCommandServer) GetMonitorCommand(context.Context, *GetMonitorCommandRequest) (*GetMonitorCommandResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionCommandServer) SetMonitorCommand(context.Context, *SetMonitorCommandRequest) (*SetMonitorCommandResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionCommandServer) GetKillTimeout(context.Context, *GetKillTimeoutRequest) (*GetKillTimeoutResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionCommandServer) SetKillTimeout(context.Context, *SetKillTimeoutRequest) (*SetKillTimeoutResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionCommandServer) GetLogResult(context.Context, *GetLogResultRequest) (*GetLogResultResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedActionCommandServer) SetLogResult(context.Context, *SetLogResultRequest) (*SetLogResultResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ActionCommandServer = (*UnimplementedActionCommandServer)(nil)
