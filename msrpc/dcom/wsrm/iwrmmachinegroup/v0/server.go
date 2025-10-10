package iwrmmachinegroup

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

// IWRMMachineGroup server interface.
type MachineGroupServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	CreateMachineGroup(context.Context, *CreateMachineGroupRequest) (*CreateMachineGroupResponse, error)

	GetMachineGroupInfo(context.Context, *GetMachineGroupInfoRequest) (*GetMachineGroupInfoResponse, error)

	ModifyMachineGroup(context.Context, *ModifyMachineGroupRequest) (*ModifyMachineGroupResponse, error)

	DeleteMachineGroup(context.Context, *DeleteMachineGroupRequest) (*DeleteMachineGroupResponse, error)

	RenameMachineGroup(context.Context, *RenameMachineGroupRequest) (*RenameMachineGroupResponse, error)

	AddMachine(context.Context, *AddMachineRequest) (*AddMachineResponse, error)

	GetMachineInfo(context.Context, *GetMachineInfoRequest) (*GetMachineInfoResponse, error)

	ModifyMachineInfo(context.Context, *ModifyMachineInfoRequest) (*ModifyMachineInfoResponse, error)

	DeleteMachine(context.Context, *DeleteMachineRequest) (*DeleteMachineResponse, error)
}

func RegisterMachineGroupServer(conn dcerpc.Conn, o MachineGroupServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewMachineGroupServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(MachineGroupSyntaxV0_0))...)
}

func NewMachineGroupServerHandle(o MachineGroupServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return MachineGroupServerHandle(ctx, o, opNum, r)
	}
}

func MachineGroupServerHandle(ctx context.Context, o MachineGroupServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // CreateMachineGroup
		op := &xxx_CreateMachineGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateMachineGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateMachineGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // GetMachineGroupInfo
		op := &xxx_GetMachineGroupInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMachineGroupInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMachineGroupInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ModifyMachineGroup
		op := &xxx_ModifyMachineGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ModifyMachineGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ModifyMachineGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // DeleteMachineGroup
		op := &xxx_DeleteMachineGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteMachineGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteMachineGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // RenameMachineGroup
		op := &xxx_RenameMachineGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RenameMachineGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RenameMachineGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // AddMachine
		op := &xxx_AddMachineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddMachineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddMachine(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // GetMachineInfo
		op := &xxx_GetMachineInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMachineInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMachineInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // ModifyMachineInfo
		op := &xxx_ModifyMachineInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ModifyMachineInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ModifyMachineInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // DeleteMachine
		op := &xxx_DeleteMachineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteMachineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteMachine(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWRMMachineGroup
type UnimplementedMachineGroupServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedMachineGroupServer) CreateMachineGroup(context.Context, *CreateMachineGroupRequest) (*CreateMachineGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMachineGroupServer) GetMachineGroupInfo(context.Context, *GetMachineGroupInfoRequest) (*GetMachineGroupInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMachineGroupServer) ModifyMachineGroup(context.Context, *ModifyMachineGroupRequest) (*ModifyMachineGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMachineGroupServer) DeleteMachineGroup(context.Context, *DeleteMachineGroupRequest) (*DeleteMachineGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMachineGroupServer) RenameMachineGroup(context.Context, *RenameMachineGroupRequest) (*RenameMachineGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMachineGroupServer) AddMachine(context.Context, *AddMachineRequest) (*AddMachineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMachineGroupServer) GetMachineInfo(context.Context, *GetMachineInfoRequest) (*GetMachineInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMachineGroupServer) ModifyMachineInfo(context.Context, *ModifyMachineInfoRequest) (*ModifyMachineInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedMachineGroupServer) DeleteMachine(context.Context, *DeleteMachineRequest) (*DeleteMachineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ MachineGroupServer = (*UnimplementedMachineGroupServer)(nil)
