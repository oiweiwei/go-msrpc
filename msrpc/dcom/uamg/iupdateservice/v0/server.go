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

	GetName(context.Context, *GetNameRequest) (*GetNameResponse, error)

	GetContentValidationCert(context.Context, *GetContentValidationCertRequest) (*GetContentValidationCertResponse, error)

	GetExpirationDate(context.Context, *GetExpirationDateRequest) (*GetExpirationDateResponse, error)

	GetIsManaged(context.Context, *GetIsManagedRequest) (*GetIsManagedResponse, error)

	GetIsRegisteredWithAu(context.Context, *GetIsRegisteredWithAuRequest) (*GetIsRegisteredWithAuResponse, error)

	GetIssueDate(context.Context, *GetIssueDateRequest) (*GetIssueDateResponse, error)

	GetOffersWindowsUpdates(context.Context, *GetOffersWindowsUpdatesRequest) (*GetOffersWindowsUpdatesResponse, error)

	GetRedirectUrls(context.Context, *GetRedirectUrlsRequest) (*GetRedirectUrlsResponse, error)

	GetServiceID(context.Context, *GetServiceIDRequest) (*GetServiceIDResponse, error)

	GetIsScanPackageService(context.Context, *GetIsScanPackageServiceRequest) (*GetIsScanPackageServiceResponse, error)

	GetCanRegisterWithAu(context.Context, *GetCanRegisterWithAuRequest) (*GetCanRegisterWithAuResponse, error)

	GetServiceURL(context.Context, *GetServiceURLRequest) (*GetServiceURLResponse, error)

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
		op := &xxx_GetIsRegisteredWithAuOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIsRegisteredWithAuRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIsRegisteredWithAu(ctx, req)
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
		op := &xxx_GetCanRegisterWithAuOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCanRegisterWithAuRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCanRegisterWithAu(ctx, req)
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
func (UnimplementedUpdateServiceServer) GetIsRegisteredWithAu(context.Context, *GetIsRegisteredWithAuRequest) (*GetIsRegisteredWithAuResponse, error) {
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
func (UnimplementedUpdateServiceServer) GetCanRegisterWithAu(context.Context, *GetCanRegisterWithAuRequest) (*GetCanRegisterWithAuResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceServer) GetServiceURL(context.Context, *GetServiceURLRequest) (*GetServiceURLResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateServiceServer) GetSetupPrefix(context.Context, *GetSetupPrefixRequest) (*GetSetupPrefixResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateServiceServer = (*UnimplementedUpdateServiceServer)(nil)
