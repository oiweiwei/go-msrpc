package imsmqqueuemanagement

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

// IMSMQQueueManagement server interface.
type QueueManagementServer interface {

	// IMSMQManagement base class.
	imsmqmanagement.ManagementServer

	// JournalMessageCount operation.
	GetJournalMessageCount(context.Context, *GetJournalMessageCountRequest) (*GetJournalMessageCountResponse, error)

	// BytesInJournal operation.
	GetBytesInJournal(context.Context, *GetBytesInJournalRequest) (*GetBytesInJournalResponse, error)

	// EodGetReceiveInfo operation.
	EODGetReceiveInfo(context.Context, *EODGetReceiveInfoRequest) (*EODGetReceiveInfoResponse, error)
}

func RegisterQueueManagementServer(conn dcerpc.Conn, o QueueManagementServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewQueueManagementServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(QueueManagementSyntaxV0_0))...)
}

func NewQueueManagementServerHandle(o QueueManagementServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return QueueManagementServerHandle(ctx, o, opNum, r)
	}
}

func QueueManagementServerHandle(ctx context.Context, o QueueManagementServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 16 {
		// IMSMQManagement base method.
		return imsmqmanagement.ManagementServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 16: // JournalMessageCount
		op := &xxx_GetJournalMessageCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetJournalMessageCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetJournalMessageCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // BytesInJournal
		op := &xxx_GetBytesInJournalOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetBytesInJournalRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetBytesInJournal(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // EodGetReceiveInfo
		op := &xxx_EODGetReceiveInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EODGetReceiveInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EODGetReceiveInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQQueueManagement
type UnimplementedQueueManagementServer struct {
	imsmqmanagement.UnimplementedManagementServer
}

func (UnimplementedQueueManagementServer) GetJournalMessageCount(context.Context, *GetJournalMessageCountRequest) (*GetJournalMessageCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueManagementServer) GetBytesInJournal(context.Context, *GetBytesInJournalRequest) (*GetBytesInJournalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedQueueManagementServer) EODGetReceiveInfo(context.Context, *EODGetReceiveInfoRequest) (*EODGetReceiveInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ QueueManagementServer = (*UnimplementedQueueManagementServer)(nil)
