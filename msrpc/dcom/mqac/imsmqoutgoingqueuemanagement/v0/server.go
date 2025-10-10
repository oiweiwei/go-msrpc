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
type OutgoingQueueManagementServer interface {

	// IMSMQManagement base class.
	imsmqmanagement.ManagementServer

	// State operation.
	GetState(context.Context, *GetStateRequest) (*GetStateResponse, error)

	// NextHops operation.
	GetNextHops(context.Context, *GetNextHopsRequest) (*GetNextHopsResponse, error)

	// EodGetSendInfo operation.
	EODGetSendInfo(context.Context, *EODGetSendInfoRequest) (*EODGetSendInfoResponse, error)

	// Resume operation.
	Resume(context.Context, *ResumeRequest) (*ResumeResponse, error)

	// Pause operation.
	Pause(context.Context, *PauseRequest) (*PauseResponse, error)

	// EodResend operation.
	EODResend(context.Context, *EODResendRequest) (*EODResendResponse, error)
}

func RegisterOutgoingQueueManagementServer(conn dcerpc.Conn, o OutgoingQueueManagementServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewOutgoingQueueManagementServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(OutgoingQueueManagementSyntaxV0_0))...)
}

func NewOutgoingQueueManagementServerHandle(o OutgoingQueueManagementServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return OutgoingQueueManagementServerHandle(ctx, o, opNum, r)
	}
}

func OutgoingQueueManagementServerHandle(ctx context.Context, o OutgoingQueueManagementServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 16 {
		// IMSMQManagement base method.
		return imsmqmanagement.ManagementServerHandle(ctx, o, opNum, r)
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
		op := &xxx_EODGetSendInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EODGetSendInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EODGetSendInfo(ctx, req)
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
		op := &xxx_EODResendOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EODResendRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EODResend(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQOutgoingQueueManagement
type UnimplementedOutgoingQueueManagementServer struct {
	imsmqmanagement.UnimplementedManagementServer
}

func (UnimplementedOutgoingQueueManagementServer) GetState(context.Context, *GetStateRequest) (*GetStateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOutgoingQueueManagementServer) GetNextHops(context.Context, *GetNextHopsRequest) (*GetNextHopsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOutgoingQueueManagementServer) EODGetSendInfo(context.Context, *EODGetSendInfoRequest) (*EODGetSendInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOutgoingQueueManagementServer) Resume(context.Context, *ResumeRequest) (*ResumeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOutgoingQueueManagementServer) Pause(context.Context, *PauseRequest) (*PauseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOutgoingQueueManagementServer) EODResend(context.Context, *EODResendRequest) (*EODResendResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ OutgoingQueueManagementServer = (*UnimplementedOutgoingQueueManagementServer)(nil)
