package ifolderaction

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = idispatch.GoPackage
)

// IFolderAction server interface.
type FolderActionServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Age operation.
	GetAge(context.Context, *GetAgeRequest) (*GetAgeResponse, error)

	// Age operation.
	SetAge(context.Context, *SetAgeRequest) (*SetAgeResponse, error)

	// Size operation.
	GetSize(context.Context, *GetSizeRequest) (*GetSizeResponse, error)

	// Size operation.
	SetSize(context.Context, *SetSizeRequest) (*SetSizeResponse, error)

	// Actions operation.
	GetActions(context.Context, *GetActionsRequest) (*GetActionsResponse, error)

	// Actions operation.
	SetActions(context.Context, *SetActionsRequest) (*SetActionsResponse, error)

	// SendCabTo operation.
	GetSendCabTo(context.Context, *GetSendCabToRequest) (*GetSendCabToResponse, error)

	// SendCabTo operation.
	SetSendCabTo(context.Context, *SetSendCabToRequest) (*SetSendCabToResponse, error)
}

func RegisterFolderActionServer(conn dcerpc.Conn, o FolderActionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewFolderActionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(FolderActionSyntaxV0_0))...)
}

func NewFolderActionServerHandle(o FolderActionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return FolderActionServerHandle(ctx, o, opNum, r)
	}
}

func FolderActionServerHandle(ctx context.Context, o FolderActionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Age
		op := &xxx_GetAgeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAgeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAge(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Age
		op := &xxx_SetAgeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAgeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAge(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // Size
		op := &xxx_GetSizeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSizeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Size
		op := &xxx_SetSizeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSizeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // Actions
		op := &xxx_GetActionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetActionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetActions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // Actions
		op := &xxx_SetActionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetActionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetActions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // SendCabTo
		op := &xxx_GetSendCabToOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSendCabToRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSendCabTo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // SendCabTo
		op := &xxx_SetSendCabToOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSendCabToRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSendCabTo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IFolderAction
type UnimplementedFolderActionServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedFolderActionServer) GetAge(context.Context, *GetAgeRequest) (*GetAgeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFolderActionServer) SetAge(context.Context, *SetAgeRequest) (*SetAgeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFolderActionServer) GetSize(context.Context, *GetSizeRequest) (*GetSizeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFolderActionServer) SetSize(context.Context, *SetSizeRequest) (*SetSizeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFolderActionServer) GetActions(context.Context, *GetActionsRequest) (*GetActionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFolderActionServer) SetActions(context.Context, *SetActionsRequest) (*SetActionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFolderActionServer) GetSendCabTo(context.Context, *GetSendCabToRequest) (*GetSendCabToResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFolderActionServer) SetSendCabTo(context.Context, *SetSendCabToRequest) (*SetSendCabToResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ FolderActionServer = (*UnimplementedFolderActionServer)(nil)
