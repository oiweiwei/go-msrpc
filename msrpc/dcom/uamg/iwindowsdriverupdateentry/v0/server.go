package iwindowsdriverupdateentry

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

// IWindowsDriverUpdateEntry server interface.
type WindowsDriverUpdateEntryServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	GetDriverClass(context.Context, *GetDriverClassRequest) (*GetDriverClassResponse, error)

	GetDriverHardwareID(context.Context, *GetDriverHardwareIDRequest) (*GetDriverHardwareIDResponse, error)

	GetDriverManufacturer(context.Context, *GetDriverManufacturerRequest) (*GetDriverManufacturerResponse, error)

	GetDriverModel(context.Context, *GetDriverModelRequest) (*GetDriverModelResponse, error)

	GetDriverProvider(context.Context, *GetDriverProviderRequest) (*GetDriverProviderResponse, error)

	GetDriverVerDate(context.Context, *GetDriverVerDateRequest) (*GetDriverVerDateResponse, error)

	GetDeviceProblemNumber(context.Context, *GetDeviceProblemNumberRequest) (*GetDeviceProblemNumberResponse, error)

	GetDeviceStatus(context.Context, *GetDeviceStatusRequest) (*GetDeviceStatusResponse, error)
}

func RegisterWindowsDriverUpdateEntryServer(conn dcerpc.Conn, o WindowsDriverUpdateEntryServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWindowsDriverUpdateEntryServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WindowsDriverUpdateEntrySyntaxV0_0))...)
}

func NewWindowsDriverUpdateEntryServerHandle(o WindowsDriverUpdateEntryServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WindowsDriverUpdateEntryServerHandle(ctx, o, opNum, r)
	}
}

func WindowsDriverUpdateEntryServerHandle(ctx context.Context, o WindowsDriverUpdateEntryServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // DriverClass
		op := &xxx_GetDriverClassOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDriverClassRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDriverClass(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // DriverHardwareID
		op := &xxx_GetDriverHardwareIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDriverHardwareIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDriverHardwareID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // DriverManufacturer
		op := &xxx_GetDriverManufacturerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDriverManufacturerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDriverManufacturer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // DriverModel
		op := &xxx_GetDriverModelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDriverModelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDriverModel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // DriverProvider
		op := &xxx_GetDriverProviderOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDriverProviderRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDriverProvider(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // DriverVerDate
		op := &xxx_GetDriverVerDateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDriverVerDateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDriverVerDate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // DeviceProblemNumber
		op := &xxx_GetDeviceProblemNumberOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDeviceProblemNumberRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDeviceProblemNumber(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // DeviceStatus
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

// Unimplemented IWindowsDriverUpdateEntry
type UnimplementedWindowsDriverUpdateEntryServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedWindowsDriverUpdateEntryServer) GetDriverClass(context.Context, *GetDriverClassRequest) (*GetDriverClassResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdateEntryServer) GetDriverHardwareID(context.Context, *GetDriverHardwareIDRequest) (*GetDriverHardwareIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdateEntryServer) GetDriverManufacturer(context.Context, *GetDriverManufacturerRequest) (*GetDriverManufacturerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdateEntryServer) GetDriverModel(context.Context, *GetDriverModelRequest) (*GetDriverModelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdateEntryServer) GetDriverProvider(context.Context, *GetDriverProviderRequest) (*GetDriverProviderResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdateEntryServer) GetDriverVerDate(context.Context, *GetDriverVerDateRequest) (*GetDriverVerDateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdateEntryServer) GetDeviceProblemNumber(context.Context, *GetDeviceProblemNumberRequest) (*GetDeviceProblemNumberResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdateEntryServer) GetDeviceStatus(context.Context, *GetDeviceStatusRequest) (*GetDeviceStatusResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ WindowsDriverUpdateEntryServer = (*UnimplementedWindowsDriverUpdateEntryServer)(nil)
