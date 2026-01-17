package iclustersetup

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
		op := &xxx_ConfigServiceSecretOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConfigServiceSecretRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConfigServiceSecret(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // RetrieveSvcSecret
		op := &xxx_RetrieveServiceSecretOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RetrieveServiceSecretRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RetrieveServiceSecret(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // RetrieveHostLabel
		op := &xxx_RetrieveHostLabelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RetrieveHostLabelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RetrieveHostLabel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // GetFunctionalLevel
		op := &xxx_GetFunctionalLevelOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetFunctionalLevelRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetFunctionalLevel(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // Opnum7Reserved
		// Opnum7Reserved
		return nil, nil
	case 8: // Opnum8Reserved
		// Opnum8Reserved
		return nil, nil
	case 9: // ConfigClusterCert
		op := &xxx_ConfigClusterCertOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConfigClusterCertRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConfigClusterCert(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // RetrieveClusterCert
		op := &xxx_RetrieveClusterCertOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RetrieveClusterCertRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RetrieveClusterCert(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // GenerateClusterCert
		op := &xxx_GenerateClusterCertOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GenerateClusterCertRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GenerateClusterCert(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // GetUpgradeVersion
		op := &xxx_GetUpgradeVersionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUpgradeVersionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUpgradeVersion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // Opnum13Reserved
		// Opnum13Reserved
		return nil, nil
	case 14: // ConfigClusterCerV2
		op := &xxx_ConfigClusterCerV2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConfigClusterCerV2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConfigClusterCerV2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // RetrieveClusterCertV2
		op := &xxx_RetrieveClusterCertV2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RetrieveClusterCertV2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RetrieveClusterCertV2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // GenerateClusterCertV2
		op := &xxx_GenerateClusterCertV2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GenerateClusterCertV2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GenerateClusterCertV2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IClusterSetup
type UnimplementedClusterSetupServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedClusterSetupServer) ConfigServiceSecret(context.Context, *ConfigServiceSecretRequest) (*ConfigServiceSecretResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterSetupServer) RetrieveServiceSecret(context.Context, *RetrieveServiceSecretRequest) (*RetrieveServiceSecretResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterSetupServer) RetrieveHostLabel(context.Context, *RetrieveHostLabelRequest) (*RetrieveHostLabelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterSetupServer) GetFunctionalLevel(context.Context, *GetFunctionalLevelRequest) (*GetFunctionalLevelResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterSetupServer) ConfigClusterCert(context.Context, *ConfigClusterCertRequest) (*ConfigClusterCertResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterSetupServer) RetrieveClusterCert(context.Context, *RetrieveClusterCertRequest) (*RetrieveClusterCertResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterSetupServer) GenerateClusterCert(context.Context, *GenerateClusterCertRequest) (*GenerateClusterCertResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterSetupServer) GetUpgradeVersion(context.Context, *GetUpgradeVersionRequest) (*GetUpgradeVersionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterSetupServer) ConfigClusterCerV2(context.Context, *ConfigClusterCerV2Request) (*ConfigClusterCerV2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterSetupServer) RetrieveClusterCertV2(context.Context, *RetrieveClusterCertV2Request) (*RetrieveClusterCertV2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterSetupServer) GenerateClusterCertV2(context.Context, *GenerateClusterCertV2Request) (*GenerateClusterCertV2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ClusterSetupServer = (*UnimplementedClusterSetupServer)(nil)
