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

	// The IUpdateServiceManager::Services (opnum 8) method retrieves a collection of update
	// services registered with the update agent.
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
	// This method SHOULD return the value of the Services ADM element.
	GetServices(context.Context, *GetServicesRequest) (*GetServicesResponse, error)

	// Opnum9NotUsedOnWire operation.
	// Opnum9NotUsedOnWire

	// The IUpdateServiceManager::RegisterServiceWithAU (opnum 10) method registers an update
	// service with the automatic update agent.
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
	// This method SHOULD register the given service with the automatic update agent through
	// an implementation-defined interface. This method SHOULD set the ServiceRegisteredWithAU
	// ADM element to the given update service identifier.
	RegisterServiceWithAU(context.Context, *RegisterServiceWithAURequest) (*RegisterServiceWithAUResponse, error)

	// The IUpdateServiceManager::RemoveService (opnum 11) method removes an update service
	// from the update agent.
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
	// This method SHOULD trigger the update agent to remove the given service through an
	// implementation-defined interface. The method SHOULD remove the given service from
	// the Services ADM element.
	RemoveService(context.Context, *RemoveServiceRequest) (*RemoveServiceResponse, error)

	// Opnum12NotUsedOnWire operation.
	// Opnum12NotUsedOnWire

	// The IUpdateServiceManager::AddScanPackageService (opnum 13) method registers an update
	// service based on a scan package.
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
	// This method SHOULD trigger the update agent to add a service based on this scan package
	// through an implementation-dependent<47> interface. This method SHOULD add the given
	// service to the Services ADM element and return an IUpdateService instance representing
	// the added service.
	AddScanPackageService(context.Context, *AddScanPackageServiceRequest) (*AddScanPackageServiceResponse, error)

	// The IUpdateServiceManager::SetOption (opnum 14) method sets options on this interface.
	// The "AllowedServiceID" option restricts the IUpdateServiceManager::RegisterServiceWithAU
	// (opnum 10) (section 3.44.4.2) method to work only with the given service ID. The
	// "AllowWarningUI" option controls whether a warning UI is displayed when changing
	// the service registered with the automatic update agent.
	//
	// Return Values: The method MUST return information in an HRESULT  data structure.
	// The severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// If optionName is "AllowedServiceID", this method SHOULD set the value of the AllowedServiceID
	// ADM element to the string stored in optionValue. If optionName is "AllowWarningUI",
	// this method SHOULD set the value of the AllowWarningUI ADM element to the VARIANT_BOOL
	// stored in optionValue.
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
		op := &xxx_RegisterServiceWithAUOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RegisterServiceWithAURequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RegisterServiceWithAU(ctx, req)
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
func (UnimplementedUpdateServiceManagerServer) RegisterServiceWithAU(context.Context, *RegisterServiceWithAURequest) (*RegisterServiceWithAUResponse, error) {
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
