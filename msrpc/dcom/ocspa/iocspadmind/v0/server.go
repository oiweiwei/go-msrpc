package iocspadmind

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = iunknown.GoPackage
)

// IOCSPAdminD server interface.
type OCSPAdminDServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// This method retrieves the value of a responder property from the Online Responder
	// Service.
	GetOCSPProperty(context.Context, *GetOCSPPropertyRequest) (*GetOCSPPropertyResponse, error)

	// This method configures the value of a responder property on the server.
	SetOCSPProperty(context.Context, *SetOCSPPropertyRequest) (*SetOCSPPropertyResponse, error)

	// The GetCAConfigInformation method retrieves all the properties associated with a
	// particular revocation configuration.
	GetCAConfigInformation(context.Context, *GetCAConfigInformationRequest) (*GetCAConfigInformationResponse, error)

	// This method sets all the properties for a particular revocation configuration.
	SetCAConfigInformation(context.Context, *SetCAConfigInformationRequest) (*SetCAConfigInformationResponse, error)

	// The GetSecurity method is used to retrieve the security descriptor associated with
	// the responder.
	GetSecurity(context.Context, *GetSecurityRequest) (*GetSecurityResponse, error)

	// The SetSecurity method is used to set the Online Responder Security, as defined in
	// the Abstract Data Model.
	SetSecurity(context.Context, *SetSecurityRequest) (*SetSecurityResponse, error)

	// The GetSigningCertficates method retrieves a list of certificates available at the
	// responder machine that can be used to sign responses to OCSP requests regarding certificates
	// issued by the CA certificate specified.
	GetSigningCertificates(context.Context, *GetSigningCertificatesRequest) (*GetSigningCertificatesResponse, error)

	// The GetHashAlgorithms method retrieves the list of hash algorithms available at the
	// responder that could be used along with the signing certificate associated with a
	// revocation configuration to sign OCSP responses.
	GetHashAlgorithms(context.Context, *GetHashAlgorithmsRequest) (*GetHashAlgorithmsResponse, error)

	// The GetMyRoles method retrieves the Online Responder Roles [CIMC-PP] assigned to
	// the user that calls the method.
	GetMyRoles(context.Context, *GetMyRolesRequest) (*GetMyRolesResponse, error)

	// This method queries the Online Responder Service to find out whether it is running.
	//
	// This method has no parameters.
	//
	// If the Online Responder Service is running, the server MUST return success (0) when
	// this method is invoked.
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

func RegisterOCSPAdminDServer(conn dcerpc.Conn, o OCSPAdminDServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewOCSPAdminDServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(OCSPAdminDSyntaxV0_0))...)
}

func NewOCSPAdminDServerHandle(o OCSPAdminDServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return OCSPAdminDServerHandle(ctx, o, opNum, r)
	}
}

func OCSPAdminDServerHandle(ctx context.Context, o OCSPAdminDServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetOCSPProperty
		op := &xxx_GetOCSPPropertyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOCSPPropertyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOCSPProperty(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // SetOCSPProperty
		op := &xxx_SetOCSPPropertyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetOCSPPropertyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetOCSPProperty(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // GetCAConfigInformation
		op := &xxx_GetCAConfigInformationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCAConfigInformationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCAConfigInformation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // SetCAConfigInformation
		op := &xxx_SetCAConfigInformationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetCAConfigInformationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetCAConfigInformation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // GetSecurity
		op := &xxx_GetSecurityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSecurityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSecurity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // SetSecurity
		op := &xxx_SetSecurityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSecurityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSecurity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // GetSigningCertificates
		op := &xxx_GetSigningCertificatesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSigningCertificatesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSigningCertificates(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // GetHashAlgorithms
		op := &xxx_GetHashAlgorithmsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetHashAlgorithmsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetHashAlgorithms(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // GetMyRoles
		op := &xxx_GetMyRolesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMyRolesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMyRoles(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // Ping
		op := &xxx_PingOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PingRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Ping(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IOCSPAdminD
type UnimplementedOCSPAdminDServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedOCSPAdminDServer) GetOCSPProperty(context.Context, *GetOCSPPropertyRequest) (*GetOCSPPropertyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOCSPAdminDServer) SetOCSPProperty(context.Context, *SetOCSPPropertyRequest) (*SetOCSPPropertyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOCSPAdminDServer) GetCAConfigInformation(context.Context, *GetCAConfigInformationRequest) (*GetCAConfigInformationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOCSPAdminDServer) SetCAConfigInformation(context.Context, *SetCAConfigInformationRequest) (*SetCAConfigInformationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOCSPAdminDServer) GetSecurity(context.Context, *GetSecurityRequest) (*GetSecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOCSPAdminDServer) SetSecurity(context.Context, *SetSecurityRequest) (*SetSecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOCSPAdminDServer) GetSigningCertificates(context.Context, *GetSigningCertificatesRequest) (*GetSigningCertificatesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOCSPAdminDServer) GetHashAlgorithms(context.Context, *GetHashAlgorithmsRequest) (*GetHashAlgorithmsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOCSPAdminDServer) GetMyRoles(context.Context, *GetMyRolesRequest) (*GetMyRolesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedOCSPAdminDServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ OCSPAdminDServer = (*UnimplementedOCSPAdminDServer)(nil)
