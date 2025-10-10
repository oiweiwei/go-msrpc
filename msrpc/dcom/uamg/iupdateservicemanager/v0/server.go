package iupdateservicemanager

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

// IUpdateServiceManager server interface.
type UpdateServiceManagerServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	GetServices(context.Context, *GetServicesRequest) (*GetServicesResponse, error)

	// Opnum9NotUsedOnWire operation.
	// Opnum9NotUsedOnWire

	RegisterServiceWithAu(context.Context, *RegisterServiceWithAuRequest) (*RegisterServiceWithAuResponse, error)

	RemoveService(context.Context, *RemoveServiceRequest) (*RemoveServiceResponse, error)

	// Opnum12NotUsedOnWire operation.
	// Opnum12NotUsedOnWire

	AddScanPackageService(context.Context, *AddScanPackageServiceRequest) (*AddScanPackageServiceResponse, error)

	SetOption(context.Context, *SetOptionRequest) (*SetOptionResponse, error)
}

func RegisterUpdateServiceManagerServer(conn dcerpc.Conn, o UpdateServiceManagerServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateServiceManagerServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateServiceManagerSyntaxV0_0))...)
}

func NewUpdateServiceManagerServerHandle(o UpdateServiceManagerServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateServiceManagerServerHandle(ctx, o, opNum, r)
	}
}

func UpdateServiceManagerServerHandle(ctx context.Context, o UpdateServiceManagerServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Services
		op := &xxx_GetServicesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServicesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServices(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // Opnum9NotUsedOnWire
		// Opnum9NotUsedOnWire
		return nil, nil
	case 9: // RegisterServiceWithAU
		op := &xxx_RegisterServiceWithAuOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RegisterServiceWithAuRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RegisterServiceWithAu(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // RemoveService
		op := &xxx_RemoveServiceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveServiceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveService(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // Opnum12NotUsedOnWire
		// Opnum12NotUsedOnWire
		return nil, nil
	case 12: // AddScanPackageService
		op := &xxx_AddScanPackageServiceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddScanPackageServiceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddScanPackageService(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // SetOption
		op := &xxx_SetOptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetOptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetOption(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdateServiceManager
type UnimplementedUpdateServiceManagerServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedUpdateServiceManagerServer) GetServices(context.Context, *GetServicesRequest) (*GetServicesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceManagerServer) RegisterServiceWithAu(context.Context, *RegisterServiceWithAuRequest) (*RegisterServiceWithAuResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceManagerServer) RemoveService(context.Context, *RemoveServiceRequest) (*RemoveServiceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceManagerServer) AddScanPackageService(context.Context, *AddScanPackageServiceRequest) (*AddScanPackageServiceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceManagerServer) SetOption(context.Context, *SetOptionRequest) (*SetOptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateServiceManagerServer = (*UnimplementedUpdateServiceManagerServer)(nil)
