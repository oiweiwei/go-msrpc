package mgmt

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// mgmt server interface.
type ManagementServer interface {

	// rpc__mgmt_inq_if_ids operation.
	InquireInterfaceIDs(context.Context, *InquireInterfaceIDsRequest) (*InquireInterfaceIDsResponse, error)

	// rpc__mgmt_inq_stats operation.
	InquireStats(context.Context, *InquireStatsRequest) (*InquireStatsResponse, error)

	// rpc__mgmt_is_server_listening operation.
	IsServerListening(context.Context, *IsServerListeningRequest) (*IsServerListeningResponse, error)

	// rpc__mgmt_stop_server_listening operation.
	StopServerListening(context.Context, *StopServerListeningRequest) (*StopServerListeningResponse, error)

	// rpc__mgmt_inq_princ_name operation.
	InquirePrincName(context.Context, *InquirePrincNameRequest) (*InquirePrincNameResponse, error)
}

func RegisterManagementServer(conn dcerpc.Conn, o ManagementServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewManagementServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ManagementSyntaxV1_0))...)
}

func NewManagementServerHandle(o ManagementServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ManagementServerHandle(ctx, o, opNum, r)
	}
}

func ManagementServerHandle(ctx context.Context, o ManagementServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // rpc__mgmt_inq_if_ids
		op := &xxx_InquireInterfaceIDsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InquireInterfaceIDsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InquireInterfaceIDs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // rpc__mgmt_inq_stats
		op := &xxx_InquireStatsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InquireStatsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InquireStats(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // rpc__mgmt_is_server_listening
		op := &xxx_IsServerListeningOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsServerListeningRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsServerListening(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // rpc__mgmt_stop_server_listening
		op := &xxx_StopServerListeningOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StopServerListeningRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StopServerListening(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // rpc__mgmt_inq_princ_name
		op := &xxx_InquirePrincNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InquirePrincNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InquirePrincName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented mgmt
type UnimplementedManagementServer struct {
}

func (UnimplementedManagementServer) InquireInterfaceIDs(context.Context, *InquireInterfaceIDsRequest) (*InquireInterfaceIDsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManagementServer) InquireStats(context.Context, *InquireStatsRequest) (*InquireStatsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManagementServer) IsServerListening(context.Context, *IsServerListeningRequest) (*IsServerListeningResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManagementServer) StopServerListening(context.Context, *StopServerListeningRequest) (*StopServerListeningResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedManagementServer) InquirePrincName(context.Context, *InquirePrincNameRequest) (*InquirePrincNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ManagementServer = (*UnimplementedManagementServer)(nil)
