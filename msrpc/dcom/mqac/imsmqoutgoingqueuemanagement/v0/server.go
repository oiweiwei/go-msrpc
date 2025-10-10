package imsmqoutgoingqueuemanagement

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	imsmqmanagement "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqmanagement/v0"
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
	_ = imsmqmanagement.GoPackage
)

// IMSMQOutgoingQueueManagement server interface.
type ImsmqOutgoingQueueManagementServer interface {

	// IMSMQManagement base class.
	imsmqmanagement.ImsmqManagementServer

	// State operation.
	GetState(context.Context, *GetStateRequest) (*GetStateResponse, error)

	// NextHops operation.
	GetNextHops(context.Context, *GetNextHopsRequest) (*GetNextHopsResponse, error)

	// EodGetSendInfo operation.
	EodGetSendInfo(context.Context, *EodGetSendInfoRequest) (*EodGetSendInfoResponse, error)

	// Resume operation.
	Resume(context.Context, *ResumeRequest) (*ResumeResponse, error)

	// Pause operation.
	Pause(context.Context, *PauseRequest) (*PauseResponse, error)

	// EodResend operation.
	EodResend(context.Context, *EodResendRequest) (*EodResendResponse, error)
}

func RegisterImsmqOutgoingQueueManagementServer(conn dcerpc.Conn, o ImsmqOutgoingQueueManagementServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImsmqOutgoingQueueManagementServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImsmqOutgoingQueueManagementSyntaxV0_0))...)
}

func NewImsmqOutgoingQueueManagementServerHandle(o ImsmqOutgoingQueueManagementServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImsmqOutgoingQueueManagementServerHandle(ctx, o, opNum, r)
	}
}

func ImsmqOutgoingQueueManagementServerHandle(ctx context.Context, o ImsmqOutgoingQueueManagementServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 16 {
		// IMSMQManagement base method.
		return imsmqmanagement.ImsmqManagementServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 16: // State
		op := &xxx_GetStateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetStateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetState(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // NextHops
		op := &xxx_GetNextHopsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNextHopsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNextHops(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // EodGetSendInfo
		op := &xxx_EodGetSendInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EodGetSendInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EodGetSendInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // Resume
		op := &xxx_ResumeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResumeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Resume(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // Pause
		op := &xxx_PauseOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PauseRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Pause(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // EodResend
		op := &xxx_EodResendOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EodResendRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EodResend(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQOutgoingQueueManagement
type UnimplementedImsmqOutgoingQueueManagementServer struct {
	imsmqmanagement.UnimplementedImsmqManagementServer
}

func (UnimplementedImsmqOutgoingQueueManagementServer) GetState(context.Context, *GetStateRequest) (*GetStateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqOutgoingQueueManagementServer) GetNextHops(context.Context, *GetNextHopsRequest) (*GetNextHopsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqOutgoingQueueManagementServer) EodGetSendInfo(context.Context, *EodGetSendInfoRequest) (*EodGetSendInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqOutgoingQueueManagementServer) Resume(context.Context, *ResumeRequest) (*ResumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqOutgoingQueueManagementServer) Pause(context.Context, *PauseRequest) (*PauseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqOutgoingQueueManagementServer) EodResend(context.Context, *EodResendRequest) (*EodResendResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ImsmqOutgoingQueueManagementServer = (*UnimplementedImsmqOutgoingQueueManagementServer)(nil)
