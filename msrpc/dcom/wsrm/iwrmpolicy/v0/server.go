package iwrmpolicy

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

// IWRMPolicy server interface.
type PolicyServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	GetPolicyInfo(context.Context, *GetPolicyInfoRequest) (*GetPolicyInfoResponse, error)

	CreatePolicy(context.Context, *CreatePolicyRequest) (*CreatePolicyResponse, error)

	ModifyPolicy(context.Context, *ModifyPolicyRequest) (*ModifyPolicyResponse, error)

	DeletePolicy(context.Context, *DeletePolicyRequest) (*DeletePolicyResponse, error)

	RenameAllocationPolicy(context.Context, *RenameAllocationPolicyRequest) (*RenameAllocationPolicyResponse, error)

	MoveBefore(context.Context, *MoveBeforeRequest) (*MoveBeforeResponse, error)

	MoveAfter(context.Context, *MoveAfterRequest) (*MoveAfterResponse, error)

	SetCALDefaultPolicyName(context.Context, *SetCALDefaultPolicyNameRequest) (*SetCALDefaultPolicyNameResponse, error)

	GetCALDefaultPolicyName(context.Context, *GetCALDefaultPolicyNameRequest) (*GetCALDefaultPolicyNameResponse, error)

	GetProcessList(context.Context, *GetProcessListRequest) (*GetProcessListResponse, error)

	GetCurrentPolicy(context.Context, *GetCurrentPolicyRequest) (*GetCurrentPolicyResponse, error)

	SetCurrentPolicy(context.Context, *SetCurrentPolicyRequest) (*SetCurrentPolicyResponse, error)

	GetCurrentStateAndActivePolicyName(context.Context, *GetCurrentStateAndActivePolicyNameRequest) (*GetCurrentStateAndActivePolicyNameResponse, error)

	GetConditionalPolicy(context.Context, *GetConditionalPolicyRequest) (*GetConditionalPolicyResponse, error)

	SetConditionalPolicy(context.Context, *SetConditionalPolicyRequest) (*SetConditionalPolicyResponse, error)
}

func RegisterPolicyServer(conn dcerpc.Conn, o PolicyServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewPolicyServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(PolicySyntaxV0_0))...)
}

func NewPolicyServerHandle(o PolicyServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return PolicyServerHandle(ctx, o, opNum, r)
	}
}

func PolicyServerHandle(ctx context.Context, o PolicyServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // GetPolicyInfo
		op := &xxx_GetPolicyInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPolicyInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPolicyInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // CreatePolicy
		op := &xxx_CreatePolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreatePolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreatePolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ModifyPolicy
		op := &xxx_ModifyPolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ModifyPolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ModifyPolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // DeletePolicy
		op := &xxx_DeletePolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeletePolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeletePolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // RenameAllocationPolicy
		op := &xxx_RenameAllocationPolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RenameAllocationPolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RenameAllocationPolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // MoveBefore
		op := &xxx_MoveBeforeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MoveBeforeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MoveBefore(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // MoveAfter
		op := &xxx_MoveAfterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MoveAfterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MoveAfter(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // SetCalDefaultPolicyName
		op := &xxx_SetCALDefaultPolicyNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetCALDefaultPolicyNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetCALDefaultPolicyName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // GetCalDefaultPolicyName
		op := &xxx_GetCALDefaultPolicyNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCALDefaultPolicyNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCALDefaultPolicyName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // GetProcessList
		op := &xxx_GetProcessListOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetProcessListRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetProcessList(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // GetCurrentPolicy
		op := &xxx_GetCurrentPolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCurrentPolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCurrentPolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // SetCurrentPolicy
		op := &xxx_SetCurrentPolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetCurrentPolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetCurrentPolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // GetCurrentStateAndActivePolicyName
		op := &xxx_GetCurrentStateAndActivePolicyNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCurrentStateAndActivePolicyNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCurrentStateAndActivePolicyName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // GetConditionalPolicy
		op := &xxx_GetConditionalPolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetConditionalPolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetConditionalPolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // SetConditionalPolicy
		op := &xxx_SetConditionalPolicyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetConditionalPolicyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetConditionalPolicy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWRMPolicy
type UnimplementedPolicyServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedPolicyServer) GetPolicyInfo(context.Context, *GetPolicyInfoRequest) (*GetPolicyInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) CreatePolicy(context.Context, *CreatePolicyRequest) (*CreatePolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) ModifyPolicy(context.Context, *ModifyPolicyRequest) (*ModifyPolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) DeletePolicy(context.Context, *DeletePolicyRequest) (*DeletePolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) RenameAllocationPolicy(context.Context, *RenameAllocationPolicyRequest) (*RenameAllocationPolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) MoveBefore(context.Context, *MoveBeforeRequest) (*MoveBeforeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) MoveAfter(context.Context, *MoveAfterRequest) (*MoveAfterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) SetCALDefaultPolicyName(context.Context, *SetCALDefaultPolicyNameRequest) (*SetCALDefaultPolicyNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) GetCALDefaultPolicyName(context.Context, *GetCALDefaultPolicyNameRequest) (*GetCALDefaultPolicyNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) GetProcessList(context.Context, *GetProcessListRequest) (*GetProcessListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) GetCurrentPolicy(context.Context, *GetCurrentPolicyRequest) (*GetCurrentPolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) SetCurrentPolicy(context.Context, *SetCurrentPolicyRequest) (*SetCurrentPolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) GetCurrentStateAndActivePolicyName(context.Context, *GetCurrentStateAndActivePolicyNameRequest) (*GetCurrentStateAndActivePolicyNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) GetConditionalPolicy(context.Context, *GetConditionalPolicyRequest) (*GetConditionalPolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedPolicyServer) SetConditionalPolicy(context.Context, *SetConditionalPolicyRequest) (*SetConditionalPolicyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ PolicyServer = (*UnimplementedPolicyServer)(nil)
