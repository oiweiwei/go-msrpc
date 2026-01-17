package ibindctx

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
)

// IBindCtx server interface.
type BindContextServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// RegisterObjectBound operation.
	RegisterObjectBound(context.Context, *RegisterObjectBoundRequest) (*RegisterObjectBoundResponse, error)

	// RevokeObjectBound operation.
	RevokeObjectBound(context.Context, *RevokeObjectBoundRequest) (*RevokeObjectBoundResponse, error)

	// ReleaseBoundObjects operation.
	ReleaseBoundObjects(context.Context, *ReleaseBoundObjectsRequest) (*ReleaseBoundObjectsResponse, error)

	// SetBindOptions operation.
	SetBindOptions(context.Context, *SetBindOptionsRequest) (*SetBindOptionsResponse, error)

	// GetBindOptions operation.
	GetBindOptions(context.Context, *GetBindOptionsRequest) (*GetBindOptionsResponse, error)

	// GetRunningObjectTable operation.
	GetRunningObjectTable(context.Context, *GetRunningObjectTableRequest) (*GetRunningObjectTableResponse, error)

	// RegisterObjectParam operation.
	RegisterObjectParam(context.Context, *RegisterObjectParamRequest) (*RegisterObjectParamResponse, error)

	// GetObjectParam operation.
	GetObjectParam(context.Context, *GetObjectParamRequest) (*GetObjectParamResponse, error)

	// EnumObjectParam operation.
	EnumObjectParam(context.Context, *EnumObjectParamRequest) (*EnumObjectParamResponse, error)

	// RevokeObjectParam operation.
	RevokeObjectParam(context.Context, *RevokeObjectParamRequest) (*RevokeObjectParamResponse, error)
}

func RegisterBindContextServer(conn dcerpc.Conn, o BindContextServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewBindContextServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(BindContextSyntaxV0_0))...)
}

func NewBindContextServerHandle(o BindContextServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return BindContextServerHandle(ctx, o, opNum, r)
	}
}

func BindContextServerHandle(ctx context.Context, o BindContextServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // RegisterObjectBound
		op := &xxx_RegisterObjectBoundOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RegisterObjectBoundRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RegisterObjectBound(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // RevokeObjectBound
		op := &xxx_RevokeObjectBoundOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RevokeObjectBoundRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RevokeObjectBound(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // ReleaseBoundObjects
		op := &xxx_ReleaseBoundObjectsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReleaseBoundObjectsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReleaseBoundObjects(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // SetBindOptions
		op := &xxx_SetBindOptionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetBindOptionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetBindOptions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // GetBindOptions
		op := &xxx_GetBindOptionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetBindOptionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetBindOptions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // GetRunningObjectTable
		op := &xxx_GetRunningObjectTableOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRunningObjectTableRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRunningObjectTable(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // RegisterObjectParam
		op := &xxx_RegisterObjectParamOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RegisterObjectParamRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RegisterObjectParam(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // GetObjectParam
		op := &xxx_GetObjectParamOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetObjectParamRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetObjectParam(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // EnumObjectParam
		op := &xxx_EnumObjectParamOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumObjectParamRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumObjectParam(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // RevokeObjectParam
		op := &xxx_RevokeObjectParamOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RevokeObjectParamRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RevokeObjectParam(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IBindCtx
type UnimplementedBindContextServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedBindContextServer) RegisterObjectBound(context.Context, *RegisterObjectBoundRequest) (*RegisterObjectBoundResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedBindContextServer) RevokeObjectBound(context.Context, *RevokeObjectBoundRequest) (*RevokeObjectBoundResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedBindContextServer) ReleaseBoundObjects(context.Context, *ReleaseBoundObjectsRequest) (*ReleaseBoundObjectsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedBindContextServer) SetBindOptions(context.Context, *SetBindOptionsRequest) (*SetBindOptionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedBindContextServer) GetBindOptions(context.Context, *GetBindOptionsRequest) (*GetBindOptionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedBindContextServer) GetRunningObjectTable(context.Context, *GetRunningObjectTableRequest) (*GetRunningObjectTableResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedBindContextServer) RegisterObjectParam(context.Context, *RegisterObjectParamRequest) (*RegisterObjectParamResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedBindContextServer) GetObjectParam(context.Context, *GetObjectParamRequest) (*GetObjectParamResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedBindContextServer) EnumObjectParam(context.Context, *EnumObjectParamRequest) (*EnumObjectParamResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedBindContextServer) RevokeObjectParam(context.Context, *RevokeObjectParamRequest) (*RevokeObjectParamResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ BindContextServer = (*UnimplementedBindContextServer)(nil)
