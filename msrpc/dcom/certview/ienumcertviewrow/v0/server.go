package ienumcertviewrow

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

// IEnumCERTVIEWROW server interface.
type EnumCertViewRowServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// Next operation.
	Next(context.Context, *NextRequest) (*NextResponse, error)

	// EnumCertViewColumn operation.
	EnumCertViewColumn(context.Context, *EnumCertViewColumnRequest) (*EnumCertViewColumnResponse, error)

	// EnumCertViewAttribute operation.
	EnumCertViewAttribute(context.Context, *EnumCertViewAttributeRequest) (*EnumCertViewAttributeResponse, error)

	// EnumCertViewExtension operation.
	EnumCertViewExtension(context.Context, *EnumCertViewExtensionRequest) (*EnumCertViewExtensionResponse, error)

	// Skip operation.
	Skip(context.Context, *SkipRequest) (*SkipResponse, error)

	// Reset operation.
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)

	// Clone operation.
	Clone(context.Context, *CloneRequest) (*CloneResponse, error)

	// GetMaxIndex operation.
	GetMaxIndex(context.Context, *GetMaxIndexRequest) (*GetMaxIndexResponse, error)
}

func RegisterEnumCertViewRowServer(conn dcerpc.Conn, o EnumCertViewRowServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewEnumCertViewRowServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(EnumCertViewRowSyntaxV0_0))...)
}

func NewEnumCertViewRowServerHandle(o EnumCertViewRowServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return EnumCertViewRowServerHandle(ctx, o, opNum, r)
	}
}

func EnumCertViewRowServerHandle(ctx context.Context, o EnumCertViewRowServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 8: // EnumCertViewColumn
		op := &xxx_EnumCertViewColumnOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumCertViewColumnRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumCertViewColumn(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // EnumCertViewAttribute
		op := &xxx_EnumCertViewAttributeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumCertViewAttributeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumCertViewAttribute(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // EnumCertViewExtension
		op := &xxx_EnumCertViewExtensionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumCertViewExtensionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumCertViewExtension(ctx, req)
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
	case 14: // GetMaxIndex
		op := &xxx_GetMaxIndexOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMaxIndexRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMaxIndex(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IEnumCERTVIEWROW
type UnimplementedEnumCertViewRowServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedEnumCertViewRowServer) Next(context.Context, *NextRequest) (*NextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewRowServer) EnumCertViewColumn(context.Context, *EnumCertViewColumnRequest) (*EnumCertViewColumnResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewRowServer) EnumCertViewAttribute(context.Context, *EnumCertViewAttributeRequest) (*EnumCertViewAttributeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewRowServer) EnumCertViewExtension(context.Context, *EnumCertViewExtensionRequest) (*EnumCertViewExtensionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewRowServer) Skip(context.Context, *SkipRequest) (*SkipResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewRowServer) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewRowServer) Clone(context.Context, *CloneRequest) (*CloneResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedEnumCertViewRowServer) GetMaxIndex(context.Context, *GetMaxIndexRequest) (*GetMaxIndexResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ EnumCertViewRowServer = (*UnimplementedEnumCertViewRowServer)(nil)
