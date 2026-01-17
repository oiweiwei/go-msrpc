package iwindowsdriverupdate

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = iupdate.GoPackage
)

// IWindowsDriverUpdate server interface.
type WindowsDriverUpdateServer interface {

	// IUpdate base class.
	iupdate.UpdateServer

	// The IWindowsDriverUpdate::DriverClass (opnum 53) method retrieves the class of this
	// driver.
	//
	// The IWindowsDriverUpdateEntry::DriverClass (opnum 8) method retrieves the class of
	// this driver.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the string returned by a call to the DriverClass method
	// on the IWindowsDriverUpdateEntry instance referred to by the DefaultWindowsDriverUpdateEntry
	// ADM element.
	GetDriverClass(context.Context, *GetDriverClassRequest) (*GetDriverClassResponse, error)

	// The IWindowsDriverUpdateEntry::DriverHardwareID (opnum 9) method retrieves the hardware
	// ID or compatible ID that this driver matches to be installable.
	//
	// The IWindowsDriverUpdate::DriverHardwareID (opnum 54) method retrieves the hardware
	// ID or compatible ID that this driver matches to be installable.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the DriverHardwareID ADM element.
	GetDriverHardwareID(context.Context, *GetDriverHardwareIDRequest) (*GetDriverHardwareIDResponse, error)

	// The IWindowsDriverUpdate::DriverManufacturer (opnum 55) method retrieves the language-invariant
	// name of the driver manufacturer.
	//
	// The IWindowsDriverUpdateEntry::DriverManufacturer (opnum 10) method retrieves the
	// language-invariant name of the driver manufacturer.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the string returned by a call to the DriverManufacturer
	// method on the IWindowsDriverUpdateEntry instance referred to by the DefaultWindowsDriverUpdateEntry
	// ADM element.
	GetDriverManufacturer(context.Context, *GetDriverManufacturerRequest) (*GetDriverManufacturerResponse, error)

	// The IWindowsDriverUpdate::DriverModel (opnum 56) method retrieves the language-invariant
	// model name of the device for which this driver is intended.
	//
	// The IWindowsDriverUpdateEntry::DriverModel (opnum 11) method retrieves the language-invariant
	// model name of the device for which this driver is intended.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the string returned by a call to the DriverModel method
	// on the IWindowsDriverUpdateEntry instance referred to by the DefaultWindowsDriverUpdateEntry
	// ADM element.
	GetDriverModel(context.Context, *GetDriverModelRequest) (*GetDriverModelResponse, error)

	// The IWindowsDriverUpdate::DriverProvider (opnum 57) method retrieves the language-invariant
	// name of the provider of this driver.
	//
	// The IWindowsDriverUpdateEntry::DriverProvider (opnum 12) method retrieves the language-invariant
	// name of the provider of this driver.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the string returned by a call to the DriverProvider method
	// on the IWindowsDriverUpdateEntry instance referred to by the DefaultWindowsDriverUpdateEntry
	// ADM element.
	GetDriverProvider(context.Context, *GetDriverProviderRequest) (*GetDriverProviderResponse, error)

	// The IWindowsDriverUpdate::DriverVerDate (opnum 58) method retrieves the version date
	// of the driver.
	//
	// The IWindowsDriverUpdateEntry::DriverVerDate (opnum 13) method retrieves the version
	// date of the driver.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value returned by a call to the DriverVerDate method
	// on the IWindowsDriverUpdateEntry instance referred to by the DefaultWindowsDriverUpdateEntry
	// ADM element.
	GetDriverVerDate(context.Context, *GetDriverVerDateRequest) (*GetDriverVerDateResponse, error)

	// The IWindowsDriverUpdateEntry::DeviceProblemNumber (opnum 14) method retrieves the
	// problem number of the device this driver matches.
	//
	// The IWindowsDriverUpdate::DeviceProblemNumber (opnum 59) method retrieves the problem
	// number of the device this driver matches.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the DeviceProblemNumber ADM element.
	GetDeviceProblemNumber(context.Context, *GetDeviceProblemNumberRequest) (*GetDeviceProblemNumberResponse, error)

	// The IWindowsDriverUpdateEntry::DeviceStatus (opnum 15) method retrieves the status
	// of the device this driver matches.
	//
	// The IWindowsDriverUpdate::DeviceStatus (opnum 60) method retrieves the status of
	// the device this driver matches.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the DeviceStatus ADM element.
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
