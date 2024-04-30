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
		in := &GetOCSPPropertyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetOCSPProperty(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // SetOCSPProperty
		in := &SetOCSPPropertyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetOCSPProperty(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // GetCAConfigInformation
		in := &GetCAConfigInformationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCAConfigInformation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // SetCAConfigInformation
		in := &SetCAConfigInformationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetCAConfigInformation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // GetSecurity
		in := &GetSecurityRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSecurity(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // SetSecurity
		in := &SetSecurityRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSecurity(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // GetSigningCertificates
		in := &GetSigningCertificatesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSigningCertificates(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // GetHashAlgorithms
		in := &GetHashAlgorithmsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetHashAlgorithms(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // GetMyRoles
		in := &GetMyRolesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMyRoles(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // Ping
		in := &PingRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Ping(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
