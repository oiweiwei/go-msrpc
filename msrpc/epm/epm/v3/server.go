package epm

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
)

// epm server interface.
type EpmServer interface {

	// ept_insert operation.
	Insert(context.Context, *InsertRequest) (*InsertResponse, error)

	// ept_delete operation.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)

	// ept_lookup operation.
	Lookup(context.Context, *LookupRequest) (*LookupResponse, error)

	// ept_map operation.
	Map(context.Context, *MapRequest) (*MapResponse, error)

	// ept_lookup_handle_free operation.
	LookupHandleFree(context.Context, *LookupHandleFreeRequest) (*LookupHandleFreeResponse, error)

	// ept_inq_object operation.
	InquireObject(context.Context, *InquireObjectRequest) (*InquireObjectResponse, error)

	// ept_mgmt_delete operation.
	ManagementDelete(context.Context, *ManagementDeleteRequest) (*ManagementDeleteResponse, error)
}

func RegisterEpmServer(conn dcerpc.Conn, o EpmServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEpmServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EpmSyntaxV3_0))...)
}

func NewEpmServerHandle(o EpmServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EpmServerHandle(ctx, o, opNum, r)
	}
}

func EpmServerHandle(ctx context.Context, o EpmServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // ept_insert
		op := &xxx_InsertOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InsertRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Insert(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // ept_delete
		op := &xxx_DeleteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Delete(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // ept_lookup
		op := &xxx_LookupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &LookupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Lookup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // ept_map
		op := &xxx_MapOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MapRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Map(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // ept_lookup_handle_free
		op := &xxx_LookupHandleFreeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &LookupHandleFreeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.LookupHandleFree(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // ept_inq_object
		op := &xxx_InquireObjectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InquireObjectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InquireObject(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // ept_mgmt_delete
		op := &xxx_ManagementDeleteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ManagementDeleteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ManagementDelete(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented epm
type UnimplementedEpmServer struct {
}

func (UnimplementedEpmServer) Insert(context.Context, *InsertRequest) (*InsertResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEpmServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEpmServer) Lookup(context.Context, *LookupRequest) (*LookupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEpmServer) Map(context.Context, *MapRequest) (*MapResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEpmServer) LookupHandleFree(context.Context, *LookupHandleFreeRequest) (*LookupHandleFreeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEpmServer) InquireObject(context.Context, *InquireObjectRequest) (*InquireObjectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEpmServer) ManagementDelete(context.Context, *ManagementDeleteRequest) (*ManagementDeleteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ EpmServer = (*UnimplementedEpmServer)(nil)
