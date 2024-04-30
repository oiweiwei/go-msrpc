package iclustersetup

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

// IClusterSetup server interface.
type ClusterSetupServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The ConfigSvcSecret method stores the cluster secret in an implementation-specific
	// manner on the server.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 3.
	ConfigServiceSecret(context.Context, *ConfigServiceSecretRequest) (*ConfigServiceSecretResponse, error)

	// The RetrieveSvcSecret method returns the cluster secret stored on this server.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND | The cluster secret has not yet been configured by a previous call to             |
	//	|                                 | ConfigSvcSecret.                                                                 |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 4.
	RetrieveServiceSecret(context.Context, *RetrieveServiceSecretRequest) (*RetrieveServiceSecretResponse, error)

	// The RetrieveHostLabel method returns the fully qualified domain name (FQDN) of the
	// server.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 5.
	RetrieveHostLabel(context.Context, *RetrieveHostLabelRequest) (*RetrieveHostLabelResponse, error)

	// The GetFunctionalLevel method SHOULD<33> return the maximum functional level of the
	// cluster supported by this server.
	//
	// Return Values: A signed 32-bit value that indicates the return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] section 2.1 and section 2.2.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 6.
	GetFunctionalLevel(context.Context, *GetFunctionalLevelRequest) (*GetFunctionalLevelResponse, error)

	// Opnum7Reserved operation.
	// Opnum7Reserved

	// Opnum8Reserved operation.
	// Opnum8Reserved

	// The ConfigClusterCert method SHOULD<34> store the certificate and cluster secret
	// in an implementation-specific manner on the server.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] section 2.1 and section 2.2.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 9.
	ConfigClusterCert(context.Context, *ConfigClusterCertRequest) (*ConfigClusterCertResponse, error)

	// The RetrieveClusterCert method SHOULD<35> return the certificate and cluster secret
	// stored on the server.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] section 2.1 and section 2.2.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND | The certificate or cluster secret has not yet been configured by a previous call |
	//	|                                 | to RetrieveClusterCert.                                                          |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 10.
	RetrieveClusterCert(context.Context, *RetrieveClusterCertRequest) (*RetrieveClusterCertResponse, error)

	// The GenerateClusterCert method SHOULD<36> generate and return a new certificate.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] section 2.1 and section 2.2.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 11.
	GenerateClusterCert(context.Context, *GenerateClusterCertRequest) (*GenerateClusterCertResponse, error)

	// The GetUpgradeVersion method SHOULD<37> return the maximum upgrade version of the
	// cluster supported by this server.
	//
	// Return Values: A signed 32-bit value that indicates the return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] section 2.1 and section 2.2.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 12.
	GetUpgradeVersion(context.Context, *GetUpgradeVersionRequest) (*GetUpgradeVersionResponse, error)

	// Opnum13Reserved operation.
	// Opnum13Reserved

	// ConfigClusterCerV2 operation.
	ConfigClusterCerV2(context.Context, *ConfigClusterCerV2Request) (*ConfigClusterCerV2Response, error)

	// The RetrieveClusterCertV2 method<39> SHOULD return the certificate and cluster secret
	// stored on the server.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.1 and 2.2.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND | The certificate or cluster secret has not yet been configured by a previous call |
	//	|                                 | to ConfigClusterCertV2.                                                          |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 15.
	RetrieveClusterCertV2(context.Context, *RetrieveClusterCertV2Request) (*RetrieveClusterCertV2Response, error)

	// The GenerateClusterCertV2 method<40> SHOULD generate and return a new certificate.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.1 and 2.2.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 16.
	GenerateClusterCertV2(context.Context, *GenerateClusterCertV2Request) (*GenerateClusterCertV2Response, error)
}

func RegisterClusterSetupServer(conn dcerpc.Conn, o ClusterSetupServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewClusterSetupServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ClusterSetupSyntaxV0_0))...)
}

func NewClusterSetupServerHandle(o ClusterSetupServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ClusterSetupServerHandle(ctx, o, opNum, r)
	}
}

func ClusterSetupServerHandle(ctx context.Context, o ClusterSetupServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // ConfigSvcSecret
		in := &ConfigServiceSecretRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ConfigServiceSecret(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // RetrieveSvcSecret
		in := &RetrieveServiceSecretRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RetrieveServiceSecret(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // RetrieveHostLabel
		in := &RetrieveHostLabelRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RetrieveHostLabel(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // GetFunctionalLevel
		in := &GetFunctionalLevelRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFunctionalLevel(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // Opnum7Reserved
		// Opnum7Reserved
		return nil, nil
	case 8: // Opnum8Reserved
		// Opnum8Reserved
		return nil, nil
	case 9: // ConfigClusterCert
		in := &ConfigClusterCertRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ConfigClusterCert(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // RetrieveClusterCert
		in := &RetrieveClusterCertRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RetrieveClusterCert(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // GenerateClusterCert
		in := &GenerateClusterCertRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GenerateClusterCert(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // GetUpgradeVersion
		in := &GetUpgradeVersionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetUpgradeVersion(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // Opnum13Reserved
		// Opnum13Reserved
		return nil, nil
	case 14: // ConfigClusterCerV2
		in := &ConfigClusterCerV2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ConfigClusterCerV2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // RetrieveClusterCertV2
		in := &RetrieveClusterCertV2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RetrieveClusterCertV2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // GenerateClusterCertV2
		in := &GenerateClusterCertV2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GenerateClusterCertV2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
