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
type ImsmqQueueManagementServer interface {

	// IMSMQManagement base class.
	imsmqmanagement.ImsmqManagementServer

	// JournalMessageCount operation.
	GetJournalMessageCount(context.Context, *GetJournalMessageCountRequest) (*GetJournalMessageCountResponse, error)

	// BytesInJournal operation.
	GetBytesInJournal(context.Context, *GetBytesInJournalRequest) (*GetBytesInJournalResponse, error)

	// EodGetReceiveInfo operation.
	EodGetReceiveInfo(context.Context, *EodGetReceiveInfoRequest) (*EodGetReceiveInfoResponse, error)
}

func RegisterImsmqQueueManagementServer(conn dcerpc.Conn, o ImsmqQueueManagementServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewImsmqQueueManagementServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ImsmqQueueManagementSyntaxV0_0))...)
}

func NewImsmqQueueManagementServerHandle(o ImsmqQueueManagementServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ImsmqQueueManagementServerHandle(ctx, o, opNum, r)
	}
}

func ImsmqQueueManagementServerHandle(ctx context.Context, o ImsmqQueueManagementServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 16 {
		// IMSMQManagement base method.
		return imsmqmanagement.ImsmqManagementServerHandle(ctx, o, opNum, r)
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
		op := &xxx_EodGetReceiveInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EodGetReceiveInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EodGetReceiveInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQQueueManagement
type UnimplementedImsmqQueueManagementServer struct {
	imsmqmanagement.UnimplementedImsmqManagementServer
}

func (UnimplementedImsmqQueueManagementServer) GetJournalMessageCount(context.Context, *GetJournalMessageCountRequest) (*GetJournalMessageCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueManagementServer) GetBytesInJournal(context.Context, *GetBytesInJournalRequest) (*GetBytesInJournalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedImsmqQueueManagementServer) EodGetReceiveInfo(context.Context, *EodGetReceiveInfoRequest) (*EodGetReceiveInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ImsmqQueueManagementServer = (*UnimplementedImsmqQueueManagementServer)(nil)
