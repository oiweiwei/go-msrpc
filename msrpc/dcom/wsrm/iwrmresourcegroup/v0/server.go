package iwrmresourcegroup

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

// IWRMResourceGroup server interface.
type IwrmResourceGroupServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	GetResourceGroupInfo(context.Context, *GetResourceGroupInfoRequest) (*GetResourceGroupInfoResponse, error)

	ModifyResourceGroup(context.Context, *ModifyResourceGroupRequest) (*ModifyResourceGroupResponse, error)

	CreateResourceGroup(context.Context, *CreateResourceGroupRequest) (*CreateResourceGroupResponse, error)

	DeleteResourceGroup(context.Context, *DeleteResourceGroupRequest) (*DeleteResourceGroupResponse, error)

	RenameResourceGroup(context.Context, *RenameResourceGroupRequest) (*RenameResourceGroupResponse, error)
}

func RegisterIwrmResourceGroupServer(conn dcerpc.Conn, o IwrmResourceGroupServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewIwrmResourceGroupServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(IwrmResourceGroupSyntaxV0_0))...)
}

func NewIwrmResourceGroupServerHandle(o IwrmResourceGroupServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return IwrmResourceGroupServerHandle(ctx, o, opNum, r)
	}
}

func IwrmResourceGroupServerHandle(ctx context.Context, o IwrmResourceGroupServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // GetResourceGroupInfo
		op := &xxx_GetResourceGroupInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetResourceGroupInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetResourceGroupInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // ModifyResourceGroup
		op := &xxx_ModifyResourceGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ModifyResourceGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ModifyResourceGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // CreateResourceGroup
		op := &xxx_CreateResourceGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateResourceGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateResourceGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // DeleteResourceGroup
		op := &xxx_DeleteResourceGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteResourceGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteResourceGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // RenameResourceGroup
		op := &xxx_RenameResourceGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RenameResourceGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RenameResourceGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWRMResourceGroup
type UnimplementedIwrmResourceGroupServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedIwrmResourceGroupServer) GetResourceGroupInfo(context.Context, *GetResourceGroupInfoRequest) (*GetResourceGroupInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIwrmResourceGroupServer) ModifyResourceGroup(context.Context, *ModifyResourceGroupRequest) (*ModifyResourceGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIwrmResourceGroupServer) CreateResourceGroup(context.Context, *CreateResourceGroupRequest) (*CreateResourceGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIwrmResourceGroupServer) DeleteResourceGroup(context.Context, *DeleteResourceGroupRequest) (*DeleteResourceGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedIwrmResourceGroupServer) RenameResourceGroup(context.Context, *RenameResourceGroupRequest) (*RenameResourceGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ IwrmResourceGroupServer = (*UnimplementedIwrmResourceGroupServer)(nil)
