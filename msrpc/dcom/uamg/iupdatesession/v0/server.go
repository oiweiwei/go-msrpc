package iupdatesession

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

// IUpdateSession server interface.
type UpdateSessionServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	GetClientApplicationID(context.Context, *GetClientApplicationIDRequest) (*GetClientApplicationIDResponse, error)

	SetClientApplicationID(context.Context, *SetClientApplicationIDRequest) (*SetClientApplicationIDResponse, error)

	GetReadOnly(context.Context, *GetReadOnlyRequest) (*GetReadOnlyResponse, error)

	// Opnum11NotUsedOnWire operation.
	// Opnum11NotUsedOnWire

	// Opnum12NotUsedOnWire operation.
	// Opnum12NotUsedOnWire

	CreateUpdateSearcher(context.Context, *CreateUpdateSearcherRequest) (*CreateUpdateSearcherResponse, error)

	// Opnum14NotUsedOnWire operation.
	// Opnum14NotUsedOnWire

	// Opnum15NotUsedOnWire operation.
	// Opnum15NotUsedOnWire
}

func RegisterUpdateSessionServer(conn dcerpc.Conn, o UpdateSessionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateSessionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateSessionSyntaxV0_0))...)
}

func NewUpdateSessionServerHandle(o UpdateSessionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateSessionServerHandle(ctx, o, opNum, r)
	}
}

func UpdateSessionServerHandle(ctx context.Context, o UpdateSessionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // ClientApplicationID
		op := &xxx_GetClientApplicationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClientApplicationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClientApplicationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // ClientApplicationID
		op := &xxx_SetClientApplicationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetClientApplicationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetClientApplicationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ReadOnly
		op := &xxx_GetReadOnlyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetReadOnlyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetReadOnly(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Opnum11NotUsedOnWire
		// Opnum11NotUsedOnWire
		return nil, nil
	case 11: // Opnum12NotUsedOnWire
		// Opnum12NotUsedOnWire
		return nil, nil
	case 12: // CreateUpdateSearcher
		op := &xxx_CreateUpdateSearcherOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateUpdateSearcherRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateUpdateSearcher(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Opnum14NotUsedOnWire
		// Opnum14NotUsedOnWire
		return nil, nil
	case 14: // Opnum15NotUsedOnWire
		// Opnum15NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented IUpdateSession
type UnimplementedUpdateSessionServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedUpdateSessionServer) GetClientApplicationID(context.Context, *GetClientApplicationIDRequest) (*GetClientApplicationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSessionServer) SetClientApplicationID(context.Context, *SetClientApplicationIDRequest) (*SetClientApplicationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSessionServer) GetReadOnly(context.Context, *GetReadOnlyRequest) (*GetReadOnlyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSessionServer) CreateUpdateSearcher(context.Context, *CreateUpdateSearcherRequest) (*CreateUpdateSearcherResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateSessionServer = (*UnimplementedUpdateSessionServer)(nil)
