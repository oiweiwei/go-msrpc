package iupdateserviceregistration

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

// IUpdateServiceRegistration server interface.
type UpdateServiceRegistrationServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The IUpdateServiceRegistration::RegistrationState (opnum 8) method retrieves an enumeration
	// value that describes the state of the service registration.
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
	// This method SHOULD return the value of the RegistrationState ADM element.
	GetRegistrationState(context.Context, *GetRegistrationStateRequest) (*GetRegistrationStateResponse, error)

	// The IUpdateService::ServiceID (opnum 16) method retrieves the unique identifier for
	// the update service.
	//
	// The IUpdateHistoryEntry::ServiceID (opnum 18) method retrieves the unique identifier
	// of the update service that provided the update for which the operation was performed.
	//
	// The IUpdateSearcher::ServiceID (opnum 24) method retrieves the unique identifier
	// of the update server used to search against.
	//
	// The IUpdateSearcher::ServiceID (opnum 25) method sets the unique identifier of the
	// update server used to search against.
	//
	// The IUpdateServiceRegistration::ServiceID (opnum 9) method retrieves the service
	// identifier.
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
	// This method SHOULD return the value of the ServiceID ADM element.
	GetServiceID(context.Context, *GetServiceIDRequest) (*GetServiceIDResponse, error)

	// The IUpdateServiceRegistration::IsPendingRegistrationWithAU (opnum 10) method retrieves
	// whether the service is pending registration with the automatic update agent.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// * If no service registration record is found for the given service ID, the server
	// MUST return an error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the IsPendingRegistrationWithAU ADM element.
	GetIsPendingRegistrationWithAU(context.Context, *GetIsPendingRegistrationWithAURequest) (*GetIsPendingRegistrationWithAUResponse, error)

	// The IUpdateServiceRegistration::Service (opnum 11) method retrieves information about
	// the service.
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
	// This method SHOULD return the value of the Service ADM element.
	GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error)
}

func RegisterUpdateServiceRegistrationServer(conn dcerpc.Conn, o UpdateServiceRegistrationServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateServiceRegistrationServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateServiceRegistrationSyntaxV0_0))...)
}

func NewUpdateServiceRegistrationServerHandle(o UpdateServiceRegistrationServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateServiceRegistrationServerHandle(ctx, o, opNum, r)
	}
}

func UpdateServiceRegistrationServerHandle(ctx context.Context, o UpdateServiceRegistrationServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // RegistrationState
		op := &xxx_GetRegistrationStateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRegistrationStateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRegistrationState(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // ServiceID
		op := &xxx_GetServiceIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServiceIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServiceID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // IsPendingRegistrationWithAU
		op := &xxx_GetIsPendingRegistrationWithAUOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsPendingRegistrationWithAURequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsPendingRegistrationWithAU(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // Service
		op := &xxx_GetServiceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServiceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetService(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdateServiceRegistration
type UnimplementedUpdateServiceRegistrationServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedUpdateServiceRegistrationServer) GetRegistrationState(context.Context, *GetRegistrationStateRequest) (*GetRegistrationStateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceRegistrationServer) GetServiceID(context.Context, *GetServiceIDRequest) (*GetServiceIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceRegistrationServer) GetIsPendingRegistrationWithAU(context.Context, *GetIsPendingRegistrationWithAURequest) (*GetIsPendingRegistrationWithAUResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceRegistrationServer) GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateServiceRegistrationServer = (*UnimplementedUpdateServiceRegistrationServer)(nil)
