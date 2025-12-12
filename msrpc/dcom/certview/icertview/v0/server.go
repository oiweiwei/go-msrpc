package icertview

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

// ICertView server interface.
type CertViewServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// OpenConnection operation.
	OpenConnection(context.Context, *OpenConnectionRequest) (*OpenConnectionResponse, error)

	// EnumCertViewColumn operation.
	EnumCertViewColumn(context.Context, *EnumCertViewColumnRequest) (*EnumCertViewColumnResponse, error)

	// GetColumnCount operation.
	GetColumnCount(context.Context, *GetColumnCountRequest) (*GetColumnCountResponse, error)

	// GetColumnIndex operation.
	GetColumnIndex(context.Context, *GetColumnIndexRequest) (*GetColumnIndexResponse, error)

	// SetResultColumnCount operation.
	SetResultColumnCount(context.Context, *SetResultColumnCountRequest) (*SetResultColumnCountResponse, error)

	// SetResultColumn operation.
	SetResultColumn(context.Context, *SetResultColumnRequest) (*SetResultColumnResponse, error)

	// SetRestriction operation.
	SetRestriction(context.Context, *SetRestrictionRequest) (*SetRestrictionResponse, error)

	// OpenView operation.
	OpenView(context.Context, *OpenViewRequest) (*OpenViewResponse, error)
}

func RegisterCertViewServer(conn dcerpc.Conn, o CertViewServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCertViewServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CertViewSyntaxV0_0))...)
}

func NewCertViewServerHandle(o CertViewServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CertViewServerHandle(ctx, o, opNum, r)
	}
}

func CertViewServerHandle(ctx context.Context, o CertViewServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // OpenConnection
		op := &xxx_OpenConnectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenConnectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenConnection(ctx, req)
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
	case 9: // GetColumnCount
		op := &xxx_GetColumnCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetColumnCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetColumnCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // GetColumnIndex
		op := &xxx_GetColumnIndexOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetColumnIndexRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetColumnIndex(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // SetResultColumnCount
		op := &xxx_SetResultColumnCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetResultColumnCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetResultColumnCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // SetResultColumn
		op := &xxx_SetResultColumnOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetResultColumnRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetResultColumn(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // SetRestriction
		op := &xxx_SetRestrictionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetRestrictionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetRestriction(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // OpenView
		op := &xxx_OpenViewOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenViewRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenView(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented ICertView
type UnimplementedCertViewServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedCertViewServer) OpenConnection(context.Context, *OpenConnectionRequest) (*OpenConnectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertViewServer) EnumCertViewColumn(context.Context, *EnumCertViewColumnRequest) (*EnumCertViewColumnResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertViewServer) GetColumnCount(context.Context, *GetColumnCountRequest) (*GetColumnCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertViewServer) GetColumnIndex(context.Context, *GetColumnIndexRequest) (*GetColumnIndexResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertViewServer) SetResultColumnCount(context.Context, *SetResultColumnCountRequest) (*SetResultColumnCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertViewServer) SetResultColumn(context.Context, *SetResultColumnRequest) (*SetResultColumnResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertViewServer) SetRestriction(context.Context, *SetRestrictionRequest) (*SetRestrictionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedCertViewServer) OpenView(context.Context, *OpenViewRequest) (*OpenViewResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ CertViewServer = (*UnimplementedCertViewServer)(nil)
