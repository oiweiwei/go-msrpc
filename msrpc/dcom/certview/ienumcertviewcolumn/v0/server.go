package ienumcertviewcolumn

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

// IEnumCERTVIEWCOLUMN server interface.
type EnumCertViewColumnServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Next operation.
	Next(context.Context, *NextRequest) (*NextResponse, error)

	// GetName operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// GetDisplayName operation.
	GetDisplayName(context.Context, *GetDisplayNameRequest) (*GetDisplayNameResponse, error)

	// GetType operation.
	GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error)

	// IsIndexed operation.
	IsIndexed(context.Context, *IsIndexedRequest) (*IsIndexedResponse, error)

	// GetMaxLength operation.
	GetMaxLength(context.Context, *GetMaxLengthRequest) (*GetMaxLengthResponse, error)

	// GetValue operation.
	GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error)

	// Skip operation.
	Skip(context.Context, *SkipRequest) (*SkipResponse, error)

	// Reset operation.
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)

	// Clone operation.
	Clone(context.Context, *CloneRequest) (*CloneResponse, error)
}

func RegisterEnumCertViewColumnServer(conn dcerpc.Conn, o EnumCertViewColumnServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEnumCertViewColumnServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EnumCertViewColumnSyntaxV0_0))...)
}

func NewEnumCertViewColumnServerHandle(o EnumCertViewColumnServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EnumCertViewColumnServerHandle(ctx, o, opNum, r)
	}
}

func EnumCertViewColumnServerHandle(ctx context.Context, o EnumCertViewColumnServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Next
		op := &xxx_NextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Next(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // GetName
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // GetDisplayName
		op := &xxx_GetDisplayNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDisplayNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDisplayName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // GetType
		op := &xxx_GetTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // IsIndexed
		op := &xxx_IsIndexedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &IsIndexedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.IsIndexed(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // GetMaxLength
		op := &xxx_GetMaxLengthOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMaxLengthRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMaxLength(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // GetValue
		op := &xxx_GetValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // Skip
		op := &xxx_SkipOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SkipRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Skip(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // Reset
		op := &xxx_ResetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Reset(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // Clone
		op := &xxx_CloneOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloneRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Clone(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IEnumCERTVIEWCOLUMN
type UnimplementedEnumCertViewColumnServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedEnumCertViewColumnServer) Next(context.Context, *NextRequest) (*NextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewColumnServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewColumnServer) GetDisplayName(context.Context, *GetDisplayNameRequest) (*GetDisplayNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewColumnServer) GetType(context.Context, *GetTypeRequest) (*GetTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewColumnServer) IsIndexed(context.Context, *IsIndexedRequest) (*IsIndexedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewColumnServer) GetMaxLength(context.Context, *GetMaxLengthRequest) (*GetMaxLengthResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewColumnServer) GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewColumnServer) Skip(context.Context, *SkipRequest) (*SkipResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewColumnServer) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewColumnServer) Clone(context.Context, *CloneRequest) (*CloneResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ EnumCertViewColumnServer = (*UnimplementedEnumCertViewColumnServer)(nil)
