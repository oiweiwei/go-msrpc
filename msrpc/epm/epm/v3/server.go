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
		in := &InsertRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Insert(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // ept_delete
		in := &DeleteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Delete(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // ept_lookup
		in := &LookupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Lookup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // ept_map
		in := &MapRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Map(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // ept_lookup_handle_free
		in := &LookupHandleFreeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.LookupHandleFree(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // ept_inq_object
		in := &InquireObjectRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.InquireObject(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // ept_mgmt_delete
		in := &ManagementDeleteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ManagementDelete(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
