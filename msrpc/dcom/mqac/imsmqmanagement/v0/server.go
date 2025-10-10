package imsmqmanagement

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
)

// IMSMQManagement server interface.
type ManagementServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Init operation.
	Init(context.Context, *InitRequest) (*InitResponse, error)

	// FormatName operation.
	GetFormatName(context.Context, *GetFormatNameRequest) (*GetFormatNameResponse, error)

	// Machine operation.
	GetMachine(context.Context, *GetMachineRequest) (*GetMachineResponse, error)

	// MessageCount operation.
	GetMessageCount(context.Context, *GetMessageCountRequest) (*GetMessageCountResponse, error)

	// ForeignStatus operation.
	GetForeignStatus(context.Context, *GetForeignStatusRequest) (*GetForeignStatusResponse, error)

	// QueueType operation.
	GetQueueType(context.Context, *GetQueueTypeRequest) (*GetQueueTypeResponse, error)

	// IsLocal operation.
	GetIsLocal(context.Context, *GetIsLocalRequest) (*GetIsLocalResponse, error)

	// TransactionalStatus operation.
	GetTransactionalStatus(context.Context, *GetTransactionalStatusRequest) (*GetTransactionalStatusResponse, error)

	// BytesInQueue operation.
	GetBytesInQueue(context.Context, *GetBytesInQueueRequest) (*GetBytesInQueueResponse, error)
}

func RegisterManagementServer(conn dcerpc.Conn, o ManagementServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewManagementServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ManagementSyntaxV0_0))...)
}

func NewManagementServerHandle(o ManagementServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ManagementServerHandle(ctx, o, opNum, r)
	}
}

func ManagementServerHandle(ctx context.Context, o ManagementServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Init
		op := &xxx_InitOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InitRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Init(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // FormatName
		op := &xxx_GetFormatNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFormatNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFormatName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Machine
		op := &xxx_GetMachineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMachineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMachine(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // MessageCount
		op := &xxx_GetMessageCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMessageCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMessageCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // ForeignStatus
		op := &xxx_GetForeignStatusOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetForeignStatusRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetForeignStatus(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // QueueType
		op := &xxx_GetQueueTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetQueueTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetQueueType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // IsLocal
		op := &xxx_GetIsLocalOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsLocalRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsLocal(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // TransactionalStatus
		op := &xxx_GetTransactionalStatusOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTransactionalStatusRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTransactionalStatus(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // BytesInQueue
		op := &xxx_GetBytesInQueueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetBytesInQueueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetBytesInQueue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IMSMQManagement
type UnimplementedManagementServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedManagementServer) Init(context.Context, *InitRequest) (*InitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManagementServer) GetFormatName(context.Context, *GetFormatNameRequest) (*GetFormatNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManagementServer) GetMachine(context.Context, *GetMachineRequest) (*GetMachineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManagementServer) GetMessageCount(context.Context, *GetMessageCountRequest) (*GetMessageCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManagementServer) GetForeignStatus(context.Context, *GetForeignStatusRequest) (*GetForeignStatusResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManagementServer) GetQueueType(context.Context, *GetQueueTypeRequest) (*GetQueueTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManagementServer) GetIsLocal(context.Context, *GetIsLocalRequest) (*GetIsLocalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManagementServer) GetTransactionalStatus(context.Context, *GetTransactionalStatusRequest) (*GetTransactionalStatusResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManagementServer) GetBytesInQueue(context.Context, *GetBytesInQueueRequest) (*GetBytesInQueueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ManagementServer = (*UnimplementedManagementServer)(nil)
