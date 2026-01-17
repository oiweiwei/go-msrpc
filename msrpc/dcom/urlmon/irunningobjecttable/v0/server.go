package irunningobjecttable

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = iunknown.GoPackage
)

// IRunningObjectTable server interface.
type RunningObjectTableServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Register operation.
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)

	// Revoke operation.
	Revoke(context.Context, *RevokeRequest) (*RevokeResponse, error)

	// IsRunning operation.
	IsRunning(context.Context, *IsRunningRequest) (*IsRunningResponse, error)

	// GetObject operation.
	GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error)

	// NoteChangeTime operation.
	NoteChangeTime(context.Context, *NoteChangeTimeRequest) (*NoteChangeTimeResponse, error)

	// GetTimeOfLastChange operation.
	GetTimeOfLastChange(context.Context, *GetTimeOfLastChangeRequest) (*GetTimeOfLastChangeResponse, error)

	// EnumRunning operation.
	EnumRunning(context.Context, *EnumRunningRequest) (*EnumRunningResponse, error)
}

func RegisterRunningObjectTableServer(conn dcerpc.Conn, o RunningObjectTableServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewRunningObjectTableServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(RunningObjectTableSyntaxV0_0))...)
}

func NewRunningObjectTableServerHandle(o RunningObjectTableServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return RunningObjectTableServerHandle(ctx, o, opNum, r)
	}
}

func RunningObjectTableServerHandle(ctx context.Context, o RunningObjectTableServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Register
		op := &xxx_RegisterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RegisterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Register(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Revoke
		op := &xxx_RevokeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RevokeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Revoke(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // IsRunning
		op := &xxx_IsRunningOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsRunningRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsRunning(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // GetObject
		op := &xxx_GetObjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetObjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetObject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // NoteChangeTime
		op := &xxx_NoteChangeTimeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NoteChangeTimeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NoteChangeTime(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // GetTimeOfLastChange
		op := &xxx_GetTimeOfLastChangeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTimeOfLastChangeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTimeOfLastChange(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // EnumRunning
		op := &xxx_EnumRunningOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumRunningRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumRunning(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IRunningObjectTable
type UnimplementedRunningObjectTableServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedRunningObjectTableServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRunningObjectTableServer) Revoke(context.Context, *RevokeRequest) (*RevokeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRunningObjectTableServer) IsRunning(context.Context, *IsRunningRequest) (*IsRunningResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRunningObjectTableServer) GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRunningObjectTableServer) NoteChangeTime(context.Context, *NoteChangeTimeRequest) (*NoteChangeTimeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRunningObjectTableServer) GetTimeOfLastChange(context.Context, *GetTimeOfLastChangeRequest) (*GetTimeOfLastChangeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedRunningObjectTableServer) EnumRunning(context.Context, *EnumRunningRequest) (*EnumRunningResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ RunningObjectTableServer = (*UnimplementedRunningObjectTableServer)(nil)
