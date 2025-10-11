package iwindowsdriverupdate2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iwindowsdriverupdate "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iwindowsdriverupdate/v0"
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
	_ = iwindowsdriverupdate.GoPackage
)

// IWindowsDriverUpdate2 server interface.
type WindowsDriverUpdate2Server interface {

	// IWindowsDriverUpdate base class.
	iwindowsdriverupdate.WindowsDriverUpdateServer

	// The IWindowsDriverUpdate2::RebootRequired (opnum 61) method retrieves whether a reboot
	// is needed to complete installation or uninstallation of the update.
	//
	// The IUpdate2::RebootRequired (opnum 53) method retrieves whether a reboot is needed
	// to complete installation or uninstallation of the update.
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
	// This method SHOULD return the value of the RebootRequired ADM element.
	GetRebootRequired(context.Context, *GetRebootRequiredRequest) (*GetRebootRequiredResponse, error)

	// The IUpdate2::IsPresent (opnum 54) method retrieves whether the update is installed
	// for one or more products.
	//
	// The IWindowsDriverUpdate2::IsPresent (opnum 62) method retrieves whether the update
	// is installed for one or more products.
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
	// This method SHOULD return the value of the IsPresent ADM element.
	GetIsPresent(context.Context, *GetIsPresentRequest) (*GetIsPresentResponse, error)

	// The IUpdate2::CveIDs (opnum 55) method retrieves the CVE IDs associated with the
	// update.
	//
	// The IWindowsDriverUpdate2::CveIDs (opnum 63) method retrieves the CVE IDs associated
	// with the update.
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
	// This method SHOULD return the value of the CveIDs ADM element.
	GetCveIDs(context.Context, *GetCveIDsRequest) (*GetCveIDsResponse, error)

	// Opnum64NotUsedOnWire operation.
	// Opnum64NotUsedOnWire
}

func RegisterWindowsDriverUpdate2Server(conn dcerpc.Conn, o WindowsDriverUpdate2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWindowsDriverUpdate2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WindowsDriverUpdate2SyntaxV0_0))...)
}

func NewWindowsDriverUpdate2ServerHandle(o WindowsDriverUpdate2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WindowsDriverUpdate2ServerHandle(ctx, o, opNum, r)
	}
}

func WindowsDriverUpdate2ServerHandle(ctx context.Context, o WindowsDriverUpdate2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 60 {
		// IWindowsDriverUpdate base method.
		return iwindowsdriverupdate.WindowsDriverUpdateServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 60: // RebootRequired
		op := &xxx_GetRebootRequiredOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRebootRequiredRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRebootRequired(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 61: // IsPresent
		op := &xxx_GetIsPresentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsPresentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsPresent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 62: // CveIDs
		op := &xxx_GetCveIDsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCveIDsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCveIDs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 63: // Opnum64NotUsedOnWire
		// Opnum64NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}

// Unimplemented IWindowsDriverUpdate2
type UnimplementedWindowsDriverUpdate2Server struct {
	iwindowsdriverupdate.UnimplementedWindowsDriverUpdateServer
}

func (UnimplementedWindowsDriverUpdate2Server) GetRebootRequired(context.Context, *GetRebootRequiredRequest) (*GetRebootRequiredResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdate2Server) GetIsPresent(context.Context, *GetIsPresentRequest) (*GetIsPresentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWindowsDriverUpdate2Server) GetCveIDs(context.Context, *GetCveIDsRequest) (*GetCveIDsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ WindowsDriverUpdate2Server = (*UnimplementedWindowsDriverUpdate2Server)(nil)
