package ienumcertviewextension

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

// IEnumCERTVIEWEXTENSION server interface.
type EnumCertViewExtensionServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Next operation.
	Next(context.Context, *NextRequest) (*NextResponse, error)

	// GetName operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// GetFlags operation.
	GetFlags(context.Context, *GetFlagsRequest) (*GetFlagsResponse, error)

	// GetValue operation.
	GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error)

	// Skip operation.
	Skip(context.Context, *SkipRequest) (*SkipResponse, error)

	// Reset operation.
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)

	// Clone operation.
	Clone(context.Context, *CloneRequest) (*CloneResponse, error)
}

func RegisterEnumCertViewExtensionServer(conn dcerpc.Conn, o EnumCertViewExtensionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEnumCertViewExtensionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EnumCertViewExtensionSyntaxV0_0))...)
}

func NewEnumCertViewExtensionServerHandle(o EnumCertViewExtensionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EnumCertViewExtensionServerHandle(ctx, o, opNum, r)
	}
}

func EnumCertViewExtensionServerHandle(ctx context.Context, o EnumCertViewExtensionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 9: // GetFlags
		op := &xxx_GetFlagsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFlagsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFlags(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // GetValue
		op := &xxx_GetValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // Skip
		op := &xxx_SkipOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SkipRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Skip(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // Reset
		op := &xxx_ResetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Reset(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Clone
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

// Unimplemented IEnumCERTVIEWEXTENSION
type UnimplementedEnumCertViewExtensionServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedEnumCertViewExtensionServer) Next(context.Context, *NextRequest) (*NextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewExtensionServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewExtensionServer) GetFlags(context.Context, *GetFlagsRequest) (*GetFlagsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewExtensionServer) GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewExtensionServer) Skip(context.Context, *SkipRequest) (*SkipResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewExtensionServer) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewExtensionServer) Clone(context.Context, *CloneRequest) (*CloneResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ EnumCertViewExtensionServer = (*UnimplementedEnumCertViewExtensionServer)(nil)
