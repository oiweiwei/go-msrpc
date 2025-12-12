package ienumcertviewattribute

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

// IEnumCERTVIEWATTRIBUTE server interface.
type EnumCertViewAttributeServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Next operation.
	Next(context.Context, *NextRequest) (*NextResponse, error)

	// GetName operation.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// GetValue operation.
	GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error)

	// Skip operation.
	Skip(context.Context, *SkipRequest) (*SkipResponse, error)

	// Reset operation.
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)

	// Clone operation.
	Clone(context.Context, *CloneRequest) (*CloneResponse, error)
}

func RegisterEnumCertViewAttributeServer(conn dcerpc.Conn, o EnumCertViewAttributeServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEnumCertViewAttributeServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EnumCertViewAttributeSyntaxV0_0))...)
}

func NewEnumCertViewAttributeServerHandle(o EnumCertViewAttributeServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EnumCertViewAttributeServerHandle(ctx, o, opNum, r)
	}
}

func EnumCertViewAttributeServerHandle(ctx context.Context, o EnumCertViewAttributeServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 9: // GetValue
		op := &xxx_GetValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Skip
		op := &xxx_SkipOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SkipRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Skip(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // Reset
		op := &xxx_ResetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Reset(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // Clone
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

// Unimplemented IEnumCERTVIEWATTRIBUTE
type UnimplementedEnumCertViewAttributeServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedEnumCertViewAttributeServer) Next(context.Context, *NextRequest) (*NextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewAttributeServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewAttributeServer) GetValue(context.Context, *GetValueRequest) (*GetValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewAttributeServer) Skip(context.Context, *SkipRequest) (*SkipResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewAttributeServer) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewAttributeServer) Clone(context.Context, *CloneRequest) (*CloneResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ EnumCertViewAttributeServer = (*UnimplementedEnumCertViewAttributeServer)(nil)
