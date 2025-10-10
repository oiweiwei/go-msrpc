package iwindowsdriverupdate

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iupdate "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdate/v0"
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
	_ = iupdate.GoPackage
)

// IWindowsDriverUpdate server interface.
type WindowsDriverUpdateServer interface {

	// IUpdate base class.
	iupdate.UpdateServer

	GetDriverClass(context.Context, *GetDriverClassRequest) (*GetDriverClassResponse, error)

	GetDriverHardwareID(context.Context, *GetDriverHardwareIDRequest) (*GetDriverHardwareIDResponse, error)

	GetDriverManufacturer(context.Context, *GetDriverManufacturerRequest) (*GetDriverManufacturerResponse, error)

	GetDriverModel(context.Context, *GetDriverModelRequest) (*GetDriverModelResponse, error)

	GetDriverProvider(context.Context, *GetDriverProviderRequest) (*GetDriverProviderResponse, error)

	GetDriverVerDate(context.Context, *GetDriverVerDateRequest) (*GetDriverVerDateResponse, error)

	GetDeviceProblemNumber(context.Context, *GetDeviceProblemNumberRequest) (*GetDeviceProblemNumberResponse, error)

	GetDeviceStatus(context.Context, *GetDeviceStatusRequest) (*GetDeviceStatusResponse, error)
}

func RegisterWindowsDriverUpdateServer(conn dcerpc.Conn, o WindowsDriverUpdateServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWindowsDriverUpdateServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WindowsDriverUpdateSyntaxV0_0))...)
}

func NewWindowsDriverUpdateServerHandle(o WindowsDriverUpdateServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WindowsDriverUpdateServerHandle(ctx, o, opNum, r)
	}
}

func WindowsDriverUpdateServerHandle(ctx context.Context, o WindowsDriverUpdateServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 52 {
		// IUpdate base method.
		return iupdate.UpdateServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 52: // DriverClass
		op := &xxx_GetDriverClassOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDriverClassRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDriverClass(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 53: // DriverHardwareID
		op := &xxx_GetDriverHardwareIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDriverHardwareIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDriverHardwareID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 54: // DriverManufacturer
		op := &xxx_GetDriverManufacturerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDriverManufacturerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDriverManufacturer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 55: // DriverModel
		op := &xxx_GetDriverModelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDriverModelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDriverModel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 56: // DriverProvider
		op := &xxx_GetDriverProviderOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDriverProviderRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDriverProvider(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 57: // DriverVerDate
		op := &xxx_GetDriverVerDateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDriverVerDateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDriverVerDate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 58: // DeviceProblemNumber
		op := &xxx_GetDeviceProblemNumberOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDeviceProblemNumberRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDeviceProblemNumber(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 59: // DeviceStatus
		op := &xxx_GetDeviceStatusOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDeviceStatusRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDeviceStatus(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWindowsDriverUpdate
type UnimplementedWindowsDriverUpdateServer struct {
	iupdate.UnimplementedUpdateServer
}

func (UnimplementedWindowsDriverUpdateServer) GetDriverClass(context.Context, *GetDriverClassRequest) (*GetDriverClassResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdateServer) GetDriverHardwareID(context.Context, *GetDriverHardwareIDRequest) (*GetDriverHardwareIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdateServer) GetDriverManufacturer(context.Context, *GetDriverManufacturerRequest) (*GetDriverManufacturerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdateServer) GetDriverModel(context.Context, *GetDriverModelRequest) (*GetDriverModelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdateServer) GetDriverProvider(context.Context, *GetDriverProviderRequest) (*GetDriverProviderResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdateServer) GetDriverVerDate(context.Context, *GetDriverVerDateRequest) (*GetDriverVerDateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdateServer) GetDeviceProblemNumber(context.Context, *GetDeviceProblemNumberRequest) (*GetDeviceProblemNumberResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdateServer) GetDeviceStatus(context.Context, *GetDeviceStatusRequest) (*GetDeviceStatusResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ WindowsDriverUpdateServer = (*UnimplementedWindowsDriverUpdateServer)(nil)
