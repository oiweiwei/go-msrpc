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
		in := &GetExecutablePathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetExecutablePath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // ExecutablePath
		in := &SetExecutablePathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetExecutablePath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // Arguments
		in := &GetArgumentsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetArguments(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // Arguments
		in := &SetArgumentsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetArguments(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // Account
		in := &GetAccountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAccount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // Account
		in := &SetAccountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetAccount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // WorkingDirectory
		in := &GetWorkingDirectoryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetWorkingDirectory(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // WorkingDirectory
		in := &SetWorkingDirectoryRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetWorkingDirectory(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // MonitorCommand
		in := &GetMonitorCommandRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMonitorCommand(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // MonitorCommand
		in := &SetMonitorCommandRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMonitorCommand(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // KillTimeOut
		in := &GetKillTimeoutRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetKillTimeout(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // KillTimeOut
		in := &SetKillTimeoutRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetKillTimeout(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // LogResult
		in := &GetLogResultRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetLogResult(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // LogResult
		in := &SetLogResultRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetLogResult(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
