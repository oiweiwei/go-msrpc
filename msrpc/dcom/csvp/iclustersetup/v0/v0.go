package iclustersetup

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	csvp "github.com/oiweiwei/go-msrpc/msrpc/dcom/csvp"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = oaut.GoPackage
	_ = csvp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/csvp"
)

var (
	// IClusterSetup interface identifier 491260b5-05c9-40d9-b7f2-1f7bdae0927f
	ClusterSetupIID = &dcom.IID{Data1: 0x491260b5, Data2: 0x05c9, Data3: 0x40d9, Data4: []byte{0xb7, 0xf2, 0x1f, 0x7b, 0xda, 0xe0, 0x92, 0x7f}}
	// Syntax UUID
	ClusterSetupSyntaxUUID = &uuid.UUID{TimeLow: 0x491260b5, TimeMid: 0x5c9, TimeHiAndVersion: 0x40d9, ClockSeqHiAndReserved: 0xb7, ClockSeqLow: 0xf2, Node: [6]uint8{0x1f, 0x7b, 0xda, 0xe0, 0x92, 0x7f}}
	// Syntax ID
	ClusterSetupSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ClusterSetupSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IClusterSetup interface.
type ClusterSetupClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

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
	ConfigServiceSecret(context.Context, *ConfigServiceSecretRequest, ...dcerpc.CallOption) (*ConfigServiceSecretResponse, error)

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
	RetrieveServiceSecret(context.Context, *RetrieveServiceSecretRequest, ...dcerpc.CallOption) (*RetrieveServiceSecretResponse, error)

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
	RetrieveHostLabel(context.Context, *RetrieveHostLabelRequest, ...dcerpc.CallOption) (*RetrieveHostLabelResponse, error)

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
	GetFunctionalLevel(context.Context, *GetFunctionalLevelRequest, ...dcerpc.CallOption) (*GetFunctionalLevelResponse, error)

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
	ConfigClusterCert(context.Context, *ConfigClusterCertRequest, ...dcerpc.CallOption) (*ConfigClusterCertResponse, error)

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
	RetrieveClusterCert(context.Context, *RetrieveClusterCertRequest, ...dcerpc.CallOption) (*RetrieveClusterCertResponse, error)

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
	GenerateClusterCert(context.Context, *GenerateClusterCertRequest, ...dcerpc.CallOption) (*GenerateClusterCertResponse, error)

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
	GetUpgradeVersion(context.Context, *GetUpgradeVersionRequest, ...dcerpc.CallOption) (*GetUpgradeVersionResponse, error)

	// Opnum13Reserved operation.
	// Opnum13Reserved

	// ConfigClusterCerV2 operation.
	ConfigClusterCerV2(context.Context, *ConfigClusterCerV2Request, ...dcerpc.CallOption) (*ConfigClusterCerV2Response, error)

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
	RetrieveClusterCertV2(context.Context, *RetrieveClusterCertV2Request, ...dcerpc.CallOption) (*RetrieveClusterCertV2Response, error)

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
	GenerateClusterCertV2(context.Context, *GenerateClusterCertV2Request, ...dcerpc.CallOption) (*GenerateClusterCertV2Response, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ClusterSetupClient
}

type xxx_DefaultClusterSetupClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultClusterSetupClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultClusterSetupClient) ConfigServiceSecret(ctx context.Context, in *ConfigServiceSecretRequest, opts ...dcerpc.CallOption) (*ConfigServiceSecretResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConfigServiceSecretResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterSetupClient) RetrieveServiceSecret(ctx context.Context, in *RetrieveServiceSecretRequest, opts ...dcerpc.CallOption) (*RetrieveServiceSecretResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RetrieveServiceSecretResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterSetupClient) RetrieveHostLabel(ctx context.Context, in *RetrieveHostLabelRequest, opts ...dcerpc.CallOption) (*RetrieveHostLabelResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RetrieveHostLabelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterSetupClient) GetFunctionalLevel(ctx context.Context, in *GetFunctionalLevelRequest, opts ...dcerpc.CallOption) (*GetFunctionalLevelResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetFunctionalLevelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterSetupClient) ConfigClusterCert(ctx context.Context, in *ConfigClusterCertRequest, opts ...dcerpc.CallOption) (*ConfigClusterCertResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConfigClusterCertResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterSetupClient) RetrieveClusterCert(ctx context.Context, in *RetrieveClusterCertRequest, opts ...dcerpc.CallOption) (*RetrieveClusterCertResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RetrieveClusterCertResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterSetupClient) GenerateClusterCert(ctx context.Context, in *GenerateClusterCertRequest, opts ...dcerpc.CallOption) (*GenerateClusterCertResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GenerateClusterCertResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterSetupClient) GetUpgradeVersion(ctx context.Context, in *GetUpgradeVersionRequest, opts ...dcerpc.CallOption) (*GetUpgradeVersionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetUpgradeVersionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterSetupClient) ConfigClusterCerV2(ctx context.Context, in *ConfigClusterCerV2Request, opts ...dcerpc.CallOption) (*ConfigClusterCerV2Response, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConfigClusterCerV2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterSetupClient) RetrieveClusterCertV2(ctx context.Context, in *RetrieveClusterCertV2Request, opts ...dcerpc.CallOption) (*RetrieveClusterCertV2Response, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RetrieveClusterCertV2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterSetupClient) GenerateClusterCertV2(ctx context.Context, in *GenerateClusterCertV2Request, opts ...dcerpc.CallOption) (*GenerateClusterCertV2Response, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GenerateClusterCertV2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterSetupClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultClusterSetupClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultClusterSetupClient) IPID(ctx context.Context, ipid *dcom.IPID) ClusterSetupClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultClusterSetupClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewClusterSetupClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ClusterSetupClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ClusterSetupSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iunknown.NewUnknownClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultClusterSetupClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_ConfigServiceSecretOperation structure represents the ConfigSvcSecret operation
type xxx_ConfigServiceSecretOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	SecretBlob *oaut.String   `idl:"name:SecretBLOB" json:"secret_blob"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ConfigServiceSecretOperation) OpNum() int { return 3 }

func (o *xxx_ConfigServiceSecretOperation) OpName() string {
	return "/IClusterSetup/v0/ConfigSvcSecret"
}

func (o *xxx_ConfigServiceSecretOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConfigServiceSecretOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// SecretBLOB {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SecretBlob != nil {
			_ptr_SecretBLOB := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SecretBlob != nil {
					if err := o.SecretBlob.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SecretBlob, _ptr_SecretBLOB); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConfigServiceSecretOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SecretBLOB {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_SecretBLOB := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SecretBlob == nil {
				o.SecretBlob = &oaut.String{}
			}
			if err := o.SecretBlob.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_SecretBLOB := func(ptr interface{}) { o.SecretBlob = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SecretBlob, _s_SecretBLOB, _ptr_SecretBLOB); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConfigServiceSecretOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConfigServiceSecretOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConfigServiceSecretOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ConfigServiceSecretRequest structure represents the ConfigSvcSecret operation request
type ConfigServiceSecretRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// SecretBLOB: The cluster secret for the cluster in which this server is or will be
	// a node.
	SecretBlob *oaut.String `idl:"name:SecretBLOB" json:"secret_blob"`
}

func (o *ConfigServiceSecretRequest) xxx_ToOp(ctx context.Context) *xxx_ConfigServiceSecretOperation {
	if o == nil {
		return &xxx_ConfigServiceSecretOperation{}
	}
	return &xxx_ConfigServiceSecretOperation{
		This:       o.This,
		SecretBlob: o.SecretBlob,
	}
}

func (o *ConfigServiceSecretRequest) xxx_FromOp(ctx context.Context, op *xxx_ConfigServiceSecretOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SecretBlob = op.SecretBlob
}
func (o *ConfigServiceSecretRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ConfigServiceSecretRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConfigServiceSecretOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConfigServiceSecretResponse structure represents the ConfigSvcSecret operation response
type ConfigServiceSecretResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ConfigSvcSecret return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ConfigServiceSecretResponse) xxx_ToOp(ctx context.Context) *xxx_ConfigServiceSecretOperation {
	if o == nil {
		return &xxx_ConfigServiceSecretOperation{}
	}
	return &xxx_ConfigServiceSecretOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ConfigServiceSecretResponse) xxx_FromOp(ctx context.Context, op *xxx_ConfigServiceSecretOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ConfigServiceSecretResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ConfigServiceSecretResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConfigServiceSecretOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RetrieveServiceSecretOperation structure represents the RetrieveSvcSecret operation
type xxx_RetrieveServiceSecretOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	SecretBlob *oaut.String   `idl:"name:SecretBLOB" json:"secret_blob"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RetrieveServiceSecretOperation) OpNum() int { return 4 }

func (o *xxx_RetrieveServiceSecretOperation) OpName() string {
	return "/IClusterSetup/v0/RetrieveSvcSecret"
}

func (o *xxx_RetrieveServiceSecretOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveServiceSecretOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveServiceSecretOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveServiceSecretOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveServiceSecretOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// SecretBLOB {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SecretBlob != nil {
			_ptr_SecretBLOB := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SecretBlob != nil {
					if err := o.SecretBlob.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SecretBlob, _ptr_SecretBLOB); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveServiceSecretOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// SecretBLOB {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_SecretBLOB := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SecretBlob == nil {
				o.SecretBlob = &oaut.String{}
			}
			if err := o.SecretBlob.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_SecretBLOB := func(ptr interface{}) { o.SecretBlob = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SecretBlob, _s_SecretBLOB, _ptr_SecretBLOB); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RetrieveServiceSecretRequest structure represents the RetrieveSvcSecret operation request
type RetrieveServiceSecretRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *RetrieveServiceSecretRequest) xxx_ToOp(ctx context.Context) *xxx_RetrieveServiceSecretOperation {
	if o == nil {
		return &xxx_RetrieveServiceSecretOperation{}
	}
	return &xxx_RetrieveServiceSecretOperation{
		This: o.This,
	}
}

func (o *RetrieveServiceSecretRequest) xxx_FromOp(ctx context.Context, op *xxx_RetrieveServiceSecretOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *RetrieveServiceSecretRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RetrieveServiceSecretRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RetrieveServiceSecretOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RetrieveServiceSecretResponse structure represents the RetrieveSvcSecret operation response
type RetrieveServiceSecretResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// SecretBLOB: The value of the cluster secret as stored on this server.
	SecretBlob *oaut.String `idl:"name:SecretBLOB" json:"secret_blob"`
	// Return: The RetrieveSvcSecret return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RetrieveServiceSecretResponse) xxx_ToOp(ctx context.Context) *xxx_RetrieveServiceSecretOperation {
	if o == nil {
		return &xxx_RetrieveServiceSecretOperation{}
	}
	return &xxx_RetrieveServiceSecretOperation{
		That:       o.That,
		SecretBlob: o.SecretBlob,
		Return:     o.Return,
	}
}

func (o *RetrieveServiceSecretResponse) xxx_FromOp(ctx context.Context, op *xxx_RetrieveServiceSecretOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SecretBlob = op.SecretBlob
	o.Return = op.Return
}
func (o *RetrieveServiceSecretResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RetrieveServiceSecretResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RetrieveServiceSecretOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RetrieveHostLabelOperation structure represents the RetrieveHostLabel operation
type xxx_RetrieveHostLabelOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	HostLabel *oaut.String   `idl:"name:HostLabel" json:"host_label"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RetrieveHostLabelOperation) OpNum() int { return 5 }

func (o *xxx_RetrieveHostLabelOperation) OpName() string {
	return "/IClusterSetup/v0/RetrieveHostLabel"
}

func (o *xxx_RetrieveHostLabelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveHostLabelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveHostLabelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveHostLabelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveHostLabelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// HostLabel {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.HostLabel != nil {
			_ptr_HostLabel := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.HostLabel != nil {
					if err := o.HostLabel.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.HostLabel, _ptr_HostLabel); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveHostLabelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// HostLabel {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_HostLabel := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.HostLabel == nil {
				o.HostLabel = &oaut.String{}
			}
			if err := o.HostLabel.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_HostLabel := func(ptr interface{}) { o.HostLabel = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.HostLabel, _s_HostLabel, _ptr_HostLabel); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RetrieveHostLabelRequest structure represents the RetrieveHostLabel operation request
type RetrieveHostLabelRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *RetrieveHostLabelRequest) xxx_ToOp(ctx context.Context) *xxx_RetrieveHostLabelOperation {
	if o == nil {
		return &xxx_RetrieveHostLabelOperation{}
	}
	return &xxx_RetrieveHostLabelOperation{
		This: o.This,
	}
}

func (o *RetrieveHostLabelRequest) xxx_FromOp(ctx context.Context, op *xxx_RetrieveHostLabelOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *RetrieveHostLabelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RetrieveHostLabelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RetrieveHostLabelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RetrieveHostLabelResponse structure represents the RetrieveHostLabel operation response
type RetrieveHostLabelResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// HostLabel: The host name of the server. This is the first part of the FQDN.
	HostLabel *oaut.String `idl:"name:HostLabel" json:"host_label"`
	// Return: The RetrieveHostLabel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RetrieveHostLabelResponse) xxx_ToOp(ctx context.Context) *xxx_RetrieveHostLabelOperation {
	if o == nil {
		return &xxx_RetrieveHostLabelOperation{}
	}
	return &xxx_RetrieveHostLabelOperation{
		That:      o.That,
		HostLabel: o.HostLabel,
		Return:    o.Return,
	}
}

func (o *RetrieveHostLabelResponse) xxx_FromOp(ctx context.Context, op *xxx_RetrieveHostLabelOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.HostLabel = op.HostLabel
	o.Return = op.Return
}
func (o *RetrieveHostLabelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RetrieveHostLabelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RetrieveHostLabelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFunctionalLevelOperation structure represents the GetFunctionalLevel operation
type xxx_GetFunctionalLevelOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	FunctionalLevel uint16         `idl:"name:FunctionalLevel" json:"functional_level"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFunctionalLevelOperation) OpNum() int { return 6 }

func (o *xxx_GetFunctionalLevelOperation) OpName() string {
	return "/IClusterSetup/v0/GetFunctionalLevel"
}

func (o *xxx_GetFunctionalLevelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFunctionalLevelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFunctionalLevelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFunctionalLevelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFunctionalLevelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// FunctionalLevel {out} (1:{pointer=ref}*(1))(2:{alias=WORD}(uint16))
	{
		if err := w.WriteData(o.FunctionalLevel); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFunctionalLevelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// FunctionalLevel {out} (1:{pointer=ref}*(1))(2:{alias=WORD}(uint16))
	{
		if err := w.ReadData(&o.FunctionalLevel); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetFunctionalLevelRequest structure represents the GetFunctionalLevel operation request
type GetFunctionalLevelRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetFunctionalLevelRequest) xxx_ToOp(ctx context.Context) *xxx_GetFunctionalLevelOperation {
	if o == nil {
		return &xxx_GetFunctionalLevelOperation{}
	}
	return &xxx_GetFunctionalLevelOperation{
		This: o.This,
	}
}

func (o *GetFunctionalLevelRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFunctionalLevelOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetFunctionalLevelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetFunctionalLevelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFunctionalLevelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFunctionalLevelResponse structure represents the GetFunctionalLevel operation response
type GetFunctionalLevelResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// FunctionalLevel: The cluster functional level.
	FunctionalLevel uint16 `idl:"name:FunctionalLevel" json:"functional_level"`
	// Return: The GetFunctionalLevel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFunctionalLevelResponse) xxx_ToOp(ctx context.Context) *xxx_GetFunctionalLevelOperation {
	if o == nil {
		return &xxx_GetFunctionalLevelOperation{}
	}
	return &xxx_GetFunctionalLevelOperation{
		That:            o.That,
		FunctionalLevel: o.FunctionalLevel,
		Return:          o.Return,
	}
}

func (o *GetFunctionalLevelResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFunctionalLevelOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FunctionalLevel = op.FunctionalLevel
	o.Return = op.Return
}
func (o *GetFunctionalLevelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetFunctionalLevelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFunctionalLevelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConfigClusterCertOperation structure represents the ConfigClusterCert operation
type xxx_ConfigClusterCertOperation struct {
	This        *dcom.ORPCThis    `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat    `idl:"name:That" json:"that"`
	ClusterCert *csvp.ClusterCert `idl:"name:ClusterCert" json:"cluster_cert"`
	Return      int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_ConfigClusterCertOperation) OpNum() int { return 9 }

func (o *xxx_ConfigClusterCertOperation) OpName() string {
	return "/IClusterSetup/v0/ConfigClusterCert"
}

func (o *xxx_ConfigClusterCertOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConfigClusterCertOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ClusterCert {in} (1:{pointer=ref}*(1))(2:{alias=CLUSTER_CERT}(struct))
	{
		if o.ClusterCert != nil {
			if err := o.ClusterCert.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.ClusterCert{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ConfigClusterCertOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClusterCert {in} (1:{pointer=ref}*(1))(2:{alias=CLUSTER_CERT}(struct))
	{
		if o.ClusterCert == nil {
			o.ClusterCert = &csvp.ClusterCert{}
		}
		if err := o.ClusterCert.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConfigClusterCertOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConfigClusterCertOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConfigClusterCertOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ConfigClusterCertRequest structure represents the ConfigClusterCert operation request
type ConfigClusterCertRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ClusterCert: The certificate and cluster secret for the cluster in which this server
	// is or will be a node. The CLUSTER_CERT structure is defined in section 2.2.23.
	ClusterCert *csvp.ClusterCert `idl:"name:ClusterCert" json:"cluster_cert"`
}

func (o *ConfigClusterCertRequest) xxx_ToOp(ctx context.Context) *xxx_ConfigClusterCertOperation {
	if o == nil {
		return &xxx_ConfigClusterCertOperation{}
	}
	return &xxx_ConfigClusterCertOperation{
		This:        o.This,
		ClusterCert: o.ClusterCert,
	}
}

func (o *ConfigClusterCertRequest) xxx_FromOp(ctx context.Context, op *xxx_ConfigClusterCertOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ClusterCert = op.ClusterCert
}
func (o *ConfigClusterCertRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ConfigClusterCertRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConfigClusterCertOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConfigClusterCertResponse structure represents the ConfigClusterCert operation response
type ConfigClusterCertResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ConfigClusterCert return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ConfigClusterCertResponse) xxx_ToOp(ctx context.Context) *xxx_ConfigClusterCertOperation {
	if o == nil {
		return &xxx_ConfigClusterCertOperation{}
	}
	return &xxx_ConfigClusterCertOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ConfigClusterCertResponse) xxx_FromOp(ctx context.Context, op *xxx_ConfigClusterCertOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ConfigClusterCertResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ConfigClusterCertResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConfigClusterCertOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RetrieveClusterCertOperation structure represents the RetrieveClusterCert operation
type xxx_RetrieveClusterCertOperation struct {
	This        *dcom.ORPCThis    `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat    `idl:"name:That" json:"that"`
	ClusterCert *csvp.ClusterCert `idl:"name:ClusterCert" json:"cluster_cert"`
	Return      int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_RetrieveClusterCertOperation) OpNum() int { return 10 }

func (o *xxx_RetrieveClusterCertOperation) OpName() string {
	return "/IClusterSetup/v0/RetrieveClusterCert"
}

func (o *xxx_RetrieveClusterCertOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveClusterCertOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveClusterCertOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveClusterCertOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveClusterCertOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ClusterCert {out} (1:{pointer=ref}*(1))(2:{alias=CLUSTER_CERT}(struct))
	{
		if o.ClusterCert != nil {
			if err := o.ClusterCert.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.ClusterCert{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveClusterCertOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClusterCert {out} (1:{pointer=ref}*(1))(2:{alias=CLUSTER_CERT}(struct))
	{
		if o.ClusterCert == nil {
			o.ClusterCert = &csvp.ClusterCert{}
		}
		if err := o.ClusterCert.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RetrieveClusterCertRequest structure represents the RetrieveClusterCert operation request
type RetrieveClusterCertRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *RetrieveClusterCertRequest) xxx_ToOp(ctx context.Context) *xxx_RetrieveClusterCertOperation {
	if o == nil {
		return &xxx_RetrieveClusterCertOperation{}
	}
	return &xxx_RetrieveClusterCertOperation{
		This: o.This,
	}
}

func (o *RetrieveClusterCertRequest) xxx_FromOp(ctx context.Context, op *xxx_RetrieveClusterCertOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *RetrieveClusterCertRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RetrieveClusterCertRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RetrieveClusterCertOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RetrieveClusterCertResponse structure represents the RetrieveClusterCert operation response
type RetrieveClusterCertResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ClusterCert: The certificate and cluster secret for the cluster that is stored in
	// the node. The CLUSTER_CERT structure is defined in section 2.2.23.
	ClusterCert *csvp.ClusterCert `idl:"name:ClusterCert" json:"cluster_cert"`
	// Return: The RetrieveClusterCert return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RetrieveClusterCertResponse) xxx_ToOp(ctx context.Context) *xxx_RetrieveClusterCertOperation {
	if o == nil {
		return &xxx_RetrieveClusterCertOperation{}
	}
	return &xxx_RetrieveClusterCertOperation{
		That:        o.That,
		ClusterCert: o.ClusterCert,
		Return:      o.Return,
	}
}

func (o *RetrieveClusterCertResponse) xxx_FromOp(ctx context.Context, op *xxx_RetrieveClusterCertOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ClusterCert = op.ClusterCert
	o.Return = op.Return
}
func (o *RetrieveClusterCertResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RetrieveClusterCertResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RetrieveClusterCertOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GenerateClusterCertOperation structure represents the GenerateClusterCert operation
type xxx_GenerateClusterCertOperation struct {
	This        *dcom.ORPCThis    `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat    `idl:"name:That" json:"that"`
	ClusterCert *csvp.ClusterCert `idl:"name:ClusterCert" json:"cluster_cert"`
	Return      int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_GenerateClusterCertOperation) OpNum() int { return 11 }

func (o *xxx_GenerateClusterCertOperation) OpName() string {
	return "/IClusterSetup/v0/GenerateClusterCert"
}

func (o *xxx_GenerateClusterCertOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateClusterCertOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ClusterCert {in, out} (1:{pointer=ref}*(1))(2:{alias=CLUSTER_CERT}(struct))
	{
		if o.ClusterCert != nil {
			if err := o.ClusterCert.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.ClusterCert{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GenerateClusterCertOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClusterCert {in, out} (1:{pointer=ref}*(1))(2:{alias=CLUSTER_CERT}(struct))
	{
		if o.ClusterCert == nil {
			o.ClusterCert = &csvp.ClusterCert{}
		}
		if err := o.ClusterCert.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateClusterCertOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateClusterCertOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ClusterCert {in, out} (1:{pointer=ref}*(1))(2:{alias=CLUSTER_CERT}(struct))
	{
		if o.ClusterCert != nil {
			if err := o.ClusterCert.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.ClusterCert{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateClusterCertOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClusterCert {in, out} (1:{pointer=ref}*(1))(2:{alias=CLUSTER_CERT}(struct))
	{
		if o.ClusterCert == nil {
			o.ClusterCert = &csvp.ClusterCert{}
		}
		if err := o.ClusterCert.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GenerateClusterCertRequest structure represents the GenerateClusterCert operation request
type GenerateClusterCertRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ClusterCert: The new certificate for the cluster. The CLUSTER_CERT structure is defined
	// in section 2.2.23.
	ClusterCert *csvp.ClusterCert `idl:"name:ClusterCert" json:"cluster_cert"`
}

func (o *GenerateClusterCertRequest) xxx_ToOp(ctx context.Context) *xxx_GenerateClusterCertOperation {
	if o == nil {
		return &xxx_GenerateClusterCertOperation{}
	}
	return &xxx_GenerateClusterCertOperation{
		This:        o.This,
		ClusterCert: o.ClusterCert,
	}
}

func (o *GenerateClusterCertRequest) xxx_FromOp(ctx context.Context, op *xxx_GenerateClusterCertOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ClusterCert = op.ClusterCert
}
func (o *GenerateClusterCertRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GenerateClusterCertRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GenerateClusterCertOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GenerateClusterCertResponse structure represents the GenerateClusterCert operation response
type GenerateClusterCertResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ClusterCert: The new certificate for the cluster. The CLUSTER_CERT structure is defined
	// in section 2.2.23.
	ClusterCert *csvp.ClusterCert `idl:"name:ClusterCert" json:"cluster_cert"`
	// Return: The GenerateClusterCert return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GenerateClusterCertResponse) xxx_ToOp(ctx context.Context) *xxx_GenerateClusterCertOperation {
	if o == nil {
		return &xxx_GenerateClusterCertOperation{}
	}
	return &xxx_GenerateClusterCertOperation{
		That:        o.That,
		ClusterCert: o.ClusterCert,
		Return:      o.Return,
	}
}

func (o *GenerateClusterCertResponse) xxx_FromOp(ctx context.Context, op *xxx_GenerateClusterCertOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ClusterCert = op.ClusterCert
	o.Return = op.Return
}
func (o *GenerateClusterCertResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GenerateClusterCertResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GenerateClusterCertOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetUpgradeVersionOperation structure represents the GetUpgradeVersion operation
type xxx_GetUpgradeVersionOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	UpgradeVersion uint16         `idl:"name:UpgradeVersion" json:"upgrade_version"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetUpgradeVersionOperation) OpNum() int { return 12 }

func (o *xxx_GetUpgradeVersionOperation) OpName() string {
	return "/IClusterSetup/v0/GetUpgradeVersion"
}

func (o *xxx_GetUpgradeVersionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUpgradeVersionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUpgradeVersionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUpgradeVersionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUpgradeVersionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// UpgradeVersion {out} (1:{pointer=ref}*(1))(2:{alias=WORD}(uint16))
	{
		if err := w.WriteData(o.UpgradeVersion); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUpgradeVersionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// UpgradeVersion {out} (1:{pointer=ref}*(1))(2:{alias=WORD}(uint16))
	{
		if err := w.ReadData(&o.UpgradeVersion); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetUpgradeVersionRequest structure represents the GetUpgradeVersion operation request
type GetUpgradeVersionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetUpgradeVersionRequest) xxx_ToOp(ctx context.Context) *xxx_GetUpgradeVersionOperation {
	if o == nil {
		return &xxx_GetUpgradeVersionOperation{}
	}
	return &xxx_GetUpgradeVersionOperation{
		This: o.This,
	}
}

func (o *GetUpgradeVersionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetUpgradeVersionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetUpgradeVersionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetUpgradeVersionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUpgradeVersionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetUpgradeVersionResponse structure represents the GetUpgradeVersion operation response
type GetUpgradeVersionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// UpgradeVersion: An integer value representing the upgrade version.
	UpgradeVersion uint16 `idl:"name:UpgradeVersion" json:"upgrade_version"`
	// Return: The GetUpgradeVersion return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetUpgradeVersionResponse) xxx_ToOp(ctx context.Context) *xxx_GetUpgradeVersionOperation {
	if o == nil {
		return &xxx_GetUpgradeVersionOperation{}
	}
	return &xxx_GetUpgradeVersionOperation{
		That:           o.That,
		UpgradeVersion: o.UpgradeVersion,
		Return:         o.Return,
	}
}

func (o *GetUpgradeVersionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetUpgradeVersionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.UpgradeVersion = op.UpgradeVersion
	o.Return = op.Return
}
func (o *GetUpgradeVersionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetUpgradeVersionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUpgradeVersionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConfigClusterCerV2Operation structure represents the ConfigClusterCerV2 operation
type xxx_ConfigClusterCerV2Operation struct {
	This        *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat       `idl:"name:That" json:"that"`
	ClusterCert *csvp.ClusterCert    `idl:"name:ClusterCert" json:"cluster_cert"`
	CertType    csvp.ClusterCertType `idl:"name:certType" json:"cert_type"`
	Return      int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_ConfigClusterCerV2Operation) OpNum() int { return 14 }

func (o *xxx_ConfigClusterCerV2Operation) OpName() string {
	return "/IClusterSetup/v0/ConfigClusterCerV2"
}

func (o *xxx_ConfigClusterCerV2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConfigClusterCerV2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ClusterCert {in} (1:{pointer=ref}*(1))(2:{alias=CLUSTER_CERT}(struct))
	{
		if o.ClusterCert != nil {
			if err := o.ClusterCert.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.ClusterCert{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// certType {in} (1:{alias=CLUSTER_CERTTYPE}(enum))
	{
		if err := w.WriteData(uint16(o.CertType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConfigClusterCerV2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClusterCert {in} (1:{pointer=ref}*(1))(2:{alias=CLUSTER_CERT}(struct))
	{
		if o.ClusterCert == nil {
			o.ClusterCert = &csvp.ClusterCert{}
		}
		if err := o.ClusterCert.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// certType {in} (1:{alias=CLUSTER_CERTTYPE}(enum))
	{
		if err := w.ReadData((*uint16)(&o.CertType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConfigClusterCerV2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConfigClusterCerV2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConfigClusterCerV2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ConfigClusterCerV2Request structure represents the ConfigClusterCerV2 operation request
type ConfigClusterCerV2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis       `idl:"name:This" json:"this"`
	ClusterCert *csvp.ClusterCert    `idl:"name:ClusterCert" json:"cluster_cert"`
	CertType    csvp.ClusterCertType `idl:"name:certType" json:"cert_type"`
}

func (o *ConfigClusterCerV2Request) xxx_ToOp(ctx context.Context) *xxx_ConfigClusterCerV2Operation {
	if o == nil {
		return &xxx_ConfigClusterCerV2Operation{}
	}
	return &xxx_ConfigClusterCerV2Operation{
		This:        o.This,
		ClusterCert: o.ClusterCert,
		CertType:    o.CertType,
	}
}

func (o *ConfigClusterCerV2Request) xxx_FromOp(ctx context.Context, op *xxx_ConfigClusterCerV2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ClusterCert = op.ClusterCert
	o.CertType = op.CertType
}
func (o *ConfigClusterCerV2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ConfigClusterCerV2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConfigClusterCerV2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConfigClusterCerV2Response structure represents the ConfigClusterCerV2 operation response
type ConfigClusterCerV2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ConfigClusterCerV2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ConfigClusterCerV2Response) xxx_ToOp(ctx context.Context) *xxx_ConfigClusterCerV2Operation {
	if o == nil {
		return &xxx_ConfigClusterCerV2Operation{}
	}
	return &xxx_ConfigClusterCerV2Operation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ConfigClusterCerV2Response) xxx_FromOp(ctx context.Context, op *xxx_ConfigClusterCerV2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ConfigClusterCerV2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ConfigClusterCerV2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConfigClusterCerV2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RetrieveClusterCertV2Operation structure represents the RetrieveClusterCertV2 operation
type xxx_RetrieveClusterCertV2Operation struct {
	This        *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat       `idl:"name:That" json:"that"`
	ClusterCert *csvp.ClusterCert    `idl:"name:ClusterCert" json:"cluster_cert"`
	CertType    csvp.ClusterCertType `idl:"name:certType" json:"cert_type"`
	Return      int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_RetrieveClusterCertV2Operation) OpNum() int { return 15 }

func (o *xxx_RetrieveClusterCertV2Operation) OpName() string {
	return "/IClusterSetup/v0/RetrieveClusterCertV2"
}

func (o *xxx_RetrieveClusterCertV2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveClusterCertV2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ClusterCert {in} (1:{pointer=ref}*(1))(2:{alias=CLUSTER_CERT}(struct))
	{
		if o.ClusterCert != nil {
			if err := o.ClusterCert.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.ClusterCert{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_RetrieveClusterCertV2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClusterCert {in} (1:{pointer=ref}*(1))(2:{alias=CLUSTER_CERT}(struct))
	{
		if o.ClusterCert == nil {
			o.ClusterCert = &csvp.ClusterCert{}
		}
		if err := o.ClusterCert.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveClusterCertV2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveClusterCertV2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// certType {out} (1:{pointer=ref}*(1))(2:{alias=CLUSTER_CERTTYPE}(enum))
	{
		if err := w.WriteData(uint16(o.CertType)); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RetrieveClusterCertV2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// certType {out} (1:{pointer=ref}*(1))(2:{alias=CLUSTER_CERTTYPE}(enum))
	{
		if err := w.ReadData((*uint16)(&o.CertType)); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RetrieveClusterCertV2Request structure represents the RetrieveClusterCertV2 operation request
type RetrieveClusterCertV2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ClusterCert: The certificate and cluster secret for the cluster that is stored in
	// the node. The CLUSTER_CERT structure is defined in section 2.2.23.
	ClusterCert *csvp.ClusterCert `idl:"name:ClusterCert" json:"cluster_cert"`
}

func (o *RetrieveClusterCertV2Request) xxx_ToOp(ctx context.Context) *xxx_RetrieveClusterCertV2Operation {
	if o == nil {
		return &xxx_RetrieveClusterCertV2Operation{}
	}
	return &xxx_RetrieveClusterCertV2Operation{
		This:        o.This,
		ClusterCert: o.ClusterCert,
	}
}

func (o *RetrieveClusterCertV2Request) xxx_FromOp(ctx context.Context, op *xxx_RetrieveClusterCertV2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ClusterCert = op.ClusterCert
}
func (o *RetrieveClusterCertV2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RetrieveClusterCertV2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RetrieveClusterCertV2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RetrieveClusterCertV2Response structure represents the RetrieveClusterCertV2 operation response
type RetrieveClusterCertV2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat       `idl:"name:That" json:"that"`
	CertType csvp.ClusterCertType `idl:"name:certType" json:"cert_type"`
	// Return: The RetrieveClusterCertV2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RetrieveClusterCertV2Response) xxx_ToOp(ctx context.Context) *xxx_RetrieveClusterCertV2Operation {
	if o == nil {
		return &xxx_RetrieveClusterCertV2Operation{}
	}
	return &xxx_RetrieveClusterCertV2Operation{
		That:     o.That,
		CertType: o.CertType,
		Return:   o.Return,
	}
}

func (o *RetrieveClusterCertV2Response) xxx_FromOp(ctx context.Context, op *xxx_RetrieveClusterCertV2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CertType = op.CertType
	o.Return = op.Return
}
func (o *RetrieveClusterCertV2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RetrieveClusterCertV2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RetrieveClusterCertV2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GenerateClusterCertV2Operation structure represents the GenerateClusterCertV2 operation
type xxx_GenerateClusterCertV2Operation struct {
	This        *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat       `idl:"name:That" json:"that"`
	ClusterCert *csvp.ClusterCert    `idl:"name:ClusterCert" json:"cluster_cert"`
	CertType    csvp.ClusterCertType `idl:"name:certType" json:"cert_type"`
	Return      int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_GenerateClusterCertV2Operation) OpNum() int { return 16 }

func (o *xxx_GenerateClusterCertV2Operation) OpName() string {
	return "/IClusterSetup/v0/GenerateClusterCertV2"
}

func (o *xxx_GenerateClusterCertV2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateClusterCertV2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ClusterCert {in, out} (1:{pointer=ref}*(1))(2:{alias=CLUSTER_CERT}(struct))
	{
		if o.ClusterCert != nil {
			if err := o.ClusterCert.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.ClusterCert{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// certType {in} (1:{alias=CLUSTER_CERTTYPE}(enum))
	{
		if err := w.WriteData(uint16(o.CertType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateClusterCertV2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClusterCert {in, out} (1:{pointer=ref}*(1))(2:{alias=CLUSTER_CERT}(struct))
	{
		if o.ClusterCert == nil {
			o.ClusterCert = &csvp.ClusterCert{}
		}
		if err := o.ClusterCert.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// certType {in} (1:{alias=CLUSTER_CERTTYPE}(enum))
	{
		if err := w.ReadData((*uint16)(&o.CertType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateClusterCertV2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateClusterCertV2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ClusterCert {in, out} (1:{pointer=ref}*(1))(2:{alias=CLUSTER_CERT}(struct))
	{
		if o.ClusterCert != nil {
			if err := o.ClusterCert.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.ClusterCert{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GenerateClusterCertV2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClusterCert {in, out} (1:{pointer=ref}*(1))(2:{alias=CLUSTER_CERT}(struct))
	{
		if o.ClusterCert == nil {
			o.ClusterCert = &csvp.ClusterCert{}
		}
		if err := o.ClusterCert.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GenerateClusterCertV2Request structure represents the GenerateClusterCertV2 operation request
type GenerateClusterCertV2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ClusterCert: The new certificate for the cluster. The CLUSTER_CERT structure is defined
	// in section 2.2.23.
	ClusterCert *csvp.ClusterCert    `idl:"name:ClusterCert" json:"cluster_cert"`
	CertType    csvp.ClusterCertType `idl:"name:certType" json:"cert_type"`
}

func (o *GenerateClusterCertV2Request) xxx_ToOp(ctx context.Context) *xxx_GenerateClusterCertV2Operation {
	if o == nil {
		return &xxx_GenerateClusterCertV2Operation{}
	}
	return &xxx_GenerateClusterCertV2Operation{
		This:        o.This,
		ClusterCert: o.ClusterCert,
		CertType:    o.CertType,
	}
}

func (o *GenerateClusterCertV2Request) xxx_FromOp(ctx context.Context, op *xxx_GenerateClusterCertV2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ClusterCert = op.ClusterCert
	o.CertType = op.CertType
}
func (o *GenerateClusterCertV2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GenerateClusterCertV2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GenerateClusterCertV2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GenerateClusterCertV2Response structure represents the GenerateClusterCertV2 operation response
type GenerateClusterCertV2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ClusterCert: The new certificate for the cluster. The CLUSTER_CERT structure is defined
	// in section 2.2.23.
	ClusterCert *csvp.ClusterCert `idl:"name:ClusterCert" json:"cluster_cert"`
	// Return: The GenerateClusterCertV2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GenerateClusterCertV2Response) xxx_ToOp(ctx context.Context) *xxx_GenerateClusterCertV2Operation {
	if o == nil {
		return &xxx_GenerateClusterCertV2Operation{}
	}
	return &xxx_GenerateClusterCertV2Operation{
		That:        o.That,
		ClusterCert: o.ClusterCert,
		Return:      o.Return,
	}
}

func (o *GenerateClusterCertV2Response) xxx_FromOp(ctx context.Context, op *xxx_GenerateClusterCertV2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ClusterCert = op.ClusterCert
	o.Return = op.Return
}
func (o *GenerateClusterCertV2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GenerateClusterCertV2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GenerateClusterCertV2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
