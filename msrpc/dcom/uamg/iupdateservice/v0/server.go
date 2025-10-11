package iupdateservice

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

// IUpdateService server interface.
type UpdateServiceServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The ICategory::Name (opnum 8) method retrieves the name of the update category.
	//
	// The IUpdateService::Name (opnum 8) method retrieves the name of the update service.
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
	// This method SHOULD return the value of the Name ADM element.
	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	// The IUpdateService::ContentValidationCert (opnum 9) method retrieves a cryptographic
	// hash of the certificate used to sign content delivered by this update service.
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
	// This method SHOULD return the value of the ContentValidationCert ADM element.
	GetContentValidationCert(context.Context, *GetContentValidationCertRequest) (*GetContentValidationCertResponse, error)

	// The IUpdateService::ExpirationDate (opnum 10) method retrieves the date on which
	// the authorization cabinet file for the update service expires.
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
	// This method SHOULD return the value of the ExpirationDate ADM element.
	GetExpirationDate(context.Context, *GetExpirationDateRequest) (*GetExpirationDateResponse, error)

	// The IUpdateService::IsManaged (opnum 11) method retrieves whether the update service
	// is managed by an administrator.
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
	// This method SHOULD return the value of the IsManaged ADM element.
	GetIsManaged(context.Context, *GetIsManagedRequest) (*GetIsManagedResponse, error)

	// The IUpdateService::IsRegisteredWithAU (opnum 12) method retrieves whether the update
	// service is registered with the automatic update agent.
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
	// This method SHOULD return the value of the IsRegisteredWithAU ADM element.
	GetIsRegisteredWithAU(context.Context, *GetIsRegisteredWithAURequest) (*GetIsRegisteredWithAUResponse, error)

	// The IUpdateService::IssueDate (opnum 13) method retrieves the date on which the authorization
	// cabinet file for the update service was issued.
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
	// This method SHOULD return the value of the IssueDate ADM element.
	GetIssueDate(context.Context, *GetIssueDateRequest) (*GetIssueDateResponse, error)

	// The IUpdateService::OffersWindowsUpdates (opnum 14) method retrieves whether the
	// update service offers updates for Windows.
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
	// This method SHOULD return the value of the OffersWindowsUpdates ADM element.
	GetOffersWindowsUpdates(context.Context, *GetOffersWindowsUpdatesRequest) (*GetOffersWindowsUpdatesResponse, error)

	// The IUpdateService::RedirectUrls (opnum 15) method retrieves a collection of URLs
	// that provide the locations of the redirector cabinet files for the update service.
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
	// This method SHOULD return the value of the RedirectUrls ADM element.
	GetRedirectUrls(context.Context, *GetRedirectUrlsRequest) (*GetRedirectUrlsResponse, error)

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

	// The IUpdateService::IsScanPackageService (opnum 17) method retrieves whether the
	// update service is based on a scan package.
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
	// This method SHOULD return the value of the IsScanPackageService ADM element.
	GetIsScanPackageService(context.Context, *GetIsScanPackageServiceRequest) (*GetIsScanPackageServiceResponse, error)

	// The IUpdateService::CanRegisterWithAU (opnum 18) method retrieves whether the update
	// service can be registered with the automatic update agent.
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
	// This method SHOULD return the value of the CanRegisterWithAU ADM element.
	GetCanRegisterWithAU(context.Context, *GetCanRegisterWithAURequest) (*GetCanRegisterWithAUResponse, error)

	// The IUpdateService::ServiceUrl (opnum 19) method retrieves the URL for the update
	// service.
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
	// This method SHOULD return the value of the ServiceUrl ADM element.
	GetServiceURL(context.Context, *GetServiceURLRequest) (*GetServiceURLResponse, error)

	// The IUpdateService::SetupPrefix (opnum 20) method retrieves the prefix for the setup
	// files for this update service.
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
	// This method SHOULD return the value of the SetupPrefix ADM element.
	GetSetupPrefix(context.Context, *GetSetupPrefixRequest) (*GetSetupPrefixResponse, error)
}

func RegisterUpdateServiceServer(conn dcerpc.Conn, o UpdateServiceServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateServiceServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateServiceSyntaxV0_0))...)
}

func NewUpdateServiceServerHandle(o UpdateServiceServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateServiceServerHandle(ctx, o, opNum, r)
	}
}

func UpdateServiceServerHandle(ctx context.Context, o UpdateServiceServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // Name
		op := &xxx_GetNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // ContentValidationCert
		op := &xxx_GetContentValidationCertOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetContentValidationCertRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetContentValidationCert(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ExpirationDate
		op := &xxx_GetExpirationDateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetExpirationDateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetExpirationDate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // IsManaged
		op := &xxx_GetIsManagedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsManagedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsManaged(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // IsRegisteredWithAU
		op := &xxx_GetIsRegisteredWithAUOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsRegisteredWithAURequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsRegisteredWithAU(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // IssueDate
		op := &xxx_GetIssueDateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIssueDateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIssueDate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // OffersWindowsUpdates
		op := &xxx_GetOffersWindowsUpdatesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOffersWindowsUpdatesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOffersWindowsUpdates(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // RedirectUrls
		op := &xxx_GetRedirectUrlsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRedirectUrlsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRedirectUrls(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // ServiceID
		op := &xxx_GetServiceIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServiceIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServiceID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // IsScanPackageService
		op := &xxx_GetIsScanPackageServiceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsScanPackageServiceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsScanPackageService(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // CanRegisterWithAU
		op := &xxx_GetCanRegisterWithAUOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCanRegisterWithAURequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCanRegisterWithAU(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // ServiceUrl
		op := &xxx_GetServiceURLOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServiceURLRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServiceURL(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // SetupPrefix
		op := &xxx_GetSetupPrefixOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSetupPrefixRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSetupPrefix(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdateService
type UnimplementedUpdateServiceServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedUpdateServiceServer) GetName(context.Context, *GetNameRequest) (*GetNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceServer) GetContentValidationCert(context.Context, *GetContentValidationCertRequest) (*GetContentValidationCertResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceServer) GetExpirationDate(context.Context, *GetExpirationDateRequest) (*GetExpirationDateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceServer) GetIsManaged(context.Context, *GetIsManagedRequest) (*GetIsManagedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceServer) GetIsRegisteredWithAU(context.Context, *GetIsRegisteredWithAURequest) (*GetIsRegisteredWithAUResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceServer) GetIssueDate(context.Context, *GetIssueDateRequest) (*GetIssueDateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceServer) GetOffersWindowsUpdates(context.Context, *GetOffersWindowsUpdatesRequest) (*GetOffersWindowsUpdatesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceServer) GetRedirectUrls(context.Context, *GetRedirectUrlsRequest) (*GetRedirectUrlsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceServer) GetServiceID(context.Context, *GetServiceIDRequest) (*GetServiceIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceServer) GetIsScanPackageService(context.Context, *GetIsScanPackageServiceRequest) (*GetIsScanPackageServiceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceServer) GetCanRegisterWithAU(context.Context, *GetCanRegisterWithAURequest) (*GetCanRegisterWithAUResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceServer) GetServiceURL(context.Context, *GetServiceURLRequest) (*GetServiceURLResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceServer) GetSetupPrefix(context.Context, *GetSetupPrefixRequest) (*GetSetupPrefixResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateServiceServer = (*UnimplementedUpdateServiceServer)(nil)
