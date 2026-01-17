package iupdateservicemanager2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iupdateservicemanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdateservicemanager/v0"
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
	_ = iupdateservicemanager.GoPackage
)

// IUpdateServiceManager2 server interface.
type UpdateServiceManager2Server interface {

	// IUpdateServiceManager base class.
	iupdateservicemanager.UpdateServiceManagerServer

	// The IUpdateSearcher::ClientApplicationID (opnum 10) method retrieves the string used
	// to identify the current client application.
	//
	// The IUpdateSession::ClientApplicationID (opnum 9) method sets the identifier of the
	// calling application.
	//
	// The IUpdateHistoryEntry::ClientApplicationID (opnum 16) method retrieves the ID of
	// the application that initiated the operation.
	//
	// The IUpdateServiceManager2::ClientApplicationID (opnum 16) method sets a string that
	// identifies the client application that is using this interface.
	//
	// The IUpdateSession::ClientApplicationID (opnum 8) method retrieves the identifier
	// of the calling application.
	//
	// The IUpdateSearcher::ClientApplicationID (opnum 11) method sets the string used to
	// identify the current client application.
	//
	// The IUpdateServiceManager2::ClientApplicationID (opnum 15) method retrieves a string
	// that identifies the client application that is using this interface.
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
	// This method SHOULD return the value of the ClientApplicationID ADM element.
	GetClientApplicationID(context.Context, *GetClientApplicationIDRequest) (*GetClientApplicationIDResponse, error)

	// The IUpdateSearcher::ClientApplicationID (opnum 10) method retrieves the string used
	// to identify the current client application.
	//
	// The IUpdateSession::ClientApplicationID (opnum 9) method sets the identifier of the
	// calling application.
	//
	// The IUpdateHistoryEntry::ClientApplicationID (opnum 16) method retrieves the ID of
	// the application that initiated the operation.
	//
	// The IUpdateServiceManager2::ClientApplicationID (opnum 16) method sets a string that
	// identifies the client application that is using this interface.
	//
	// The IUpdateSession::ClientApplicationID (opnum 8) method retrieves the identifier
	// of the calling application.
	//
	// The IUpdateSearcher::ClientApplicationID (opnum 11) method sets the string used to
	// identify the current client application.
	//
	// The IUpdateServiceManager2::ClientApplicationID (opnum 15) method retrieves a string
	// that identifies the client application that is using this interface.
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
	// This method SHOULD return the value of the ClientApplicationID ADM element.
	SetClientApplicationID(context.Context, *SetClientApplicationIDRequest) (*SetClientApplicationIDResponse, error)

	// The IUpdateServiceManager2::QueryServiceRegistration (opnum 17) method retrieves
	// the update service registration record for a service.
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
	// This method SHOULD return the service registration record for the identified service
	// from the ServiceRegistrations ADM element.
	QueryServiceRegistration(context.Context, *QueryServiceRegistrationRequest) (*QueryServiceRegistrationResponse, error)

	// The IUpdateServiceManager2::AddService2 (opnum 18) method registers an update service
	// with the update agent.
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
	// This method SHOULD trigger the update agent to add this update service through an
	// implementation-dependent<48> interface. This method SHOULD add the given service
	// to the Services ADM element maintained by the IUpdateServiceManager (section 3.44)
	// interface and return an IUpdateServiceRegistration instance that represents the service
	// registration.
	AddService2(context.Context, *AddService2Request) (*AddService2Response, error)
}

func RegisterUpdateServiceManager2Server(conn dcerpc.Conn, o UpdateServiceManager2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateServiceManager2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateServiceManager2SyntaxV0_0))...)
}

func NewUpdateServiceManager2ServerHandle(o UpdateServiceManager2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateServiceManager2ServerHandle(ctx, o, opNum, r)
	}
}

func UpdateServiceManager2ServerHandle(ctx context.Context, o UpdateServiceManager2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 14 {
		// IUpdateServiceManager base method.
		return iupdateservicemanager.UpdateServiceManagerServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 14: // ClientApplicationID
		op := &xxx_GetClientApplicationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClientApplicationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClientApplicationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // ClientApplicationID
		op := &xxx_SetClientApplicationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetClientApplicationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetClientApplicationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // QueryServiceRegistration
		op := &xxx_QueryServiceRegistrationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryServiceRegistrationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryServiceRegistration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // AddService2
		op := &xxx_AddService2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddService2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddService2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdateServiceManager2
type UnimplementedUpdateServiceManager2Server struct {
	iupdateservicemanager.UnimplementedUpdateServiceManagerServer
}

func (UnimplementedUpdateServiceManager2Server) GetClientApplicationID(context.Context, *GetClientApplicationIDRequest) (*GetClientApplicationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceManager2Server) SetClientApplicationID(context.Context, *SetClientApplicationIDRequest) (*SetClientApplicationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceManager2Server) QueryServiceRegistration(context.Context, *QueryServiceRegistrationRequest) (*QueryServiceRegistrationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceManager2Server) AddService2(context.Context, *AddService2Request) (*AddService2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateServiceManager2Server = (*UnimplementedUpdateServiceManager2Server)(nil)
