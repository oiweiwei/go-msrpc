package iclusternetwork2

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

// IClusterNetwork2 server interface.
type ClusterNetwork2Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The SendRTMessage method determines whether roundtrip communication works between
	// two network addresses.
	//
	// The server SHOULD fail this method if the server Initialization State is FALSE.
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
	SendRTMessage(context.Context, *SendRTMessageRequest) (*SendRTMessageResponse, error)

	// The InitializeNode method prepares the server in an implementation-specific way to
	// execute the other methods in the interface. It also informs the client about what
	// port will be used and version information.
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
	// The opnum field value for this method is 4.
	InitializeNode(context.Context, *InitializeNodeRequest) (*InitializeNodeResponse, error)

	// The GetIpConfigSerialized method queries the network adapter configuration and returns
	// select information about the adapters.
	//
	// The server SHOULD support this method even if the server Initialization State is
	// FALSE.
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
	GetIPConfigSerialized(context.Context, *GetIPConfigSerializedRequest) (*GetIPConfigSerializedResponse, error)

	// The CleanupNode method cleans up any state initialized by InitializeNode.
	//
	// The server SHOULD fail this method if the server Initialization State is FALSE.
	//
	// This method has no parameters.
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
	// The opnum field value for this method is 6.
	CleanupNode(context.Context, *CleanupNodeRequest) (*CleanupNodeResponse, error)

	// The QueryFirewallConfiguration method determines whether the firewall state of the
	// server is compatible with use in a failover cluster. The firewall settings that constitute
	// failover cluster compatibility are implementation-specific. When the server firewall
	// enforces policies specified in [MS-FASP], the server SHOULD determine the firewall
	// state according to how the group of rules is enabled, as specified later in this
	// section.
	//
	// The server SHOULD support this method even if the server Initialization State is
	// FALSE.
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
	// The opnum field value for this method is 7.
	QueryFirewallConfiguration(context.Context, *QueryFirewallConfigurationRequest) (*QueryFirewallConfigurationResponse, error)

	// The ProcessAddRoutes method<30> adds Route elements to a Route Collection and initiates
	// monitoring of these routes for packet loss and status data.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.1 and 2.2.
	//
	//	+-------------------+---------------------------------+
	//	|      RETURN       |                                 |
	//	|    VALUE/CODE     |           DESCRIPTION           |
	//	|                   |                                 |
	//	+-------------------+---------------------------------+
	//	+-------------------+---------------------------------+
	//	| 0x00000000 S_OK   | The call was successful.        |
	//	+-------------------+---------------------------------+
	//	| 0x80004005 E_FAIL | Route Monitoring State is TRUE. |
	//	+-------------------+---------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	//
	// The opnum field value for this method is 8.
	ProcessAddRoutes(context.Context, *ProcessAddRoutesRequest) (*ProcessAddRoutesResponse, error)

	// The GetAddRoutesStatus method<31> retrieves packet loss information and status for
	// the Route elements in the Route Collection previously added with the ProcessAddRoutes
	// method.
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
	// RPC protocol specified in [MS-RPCE].
	//
	// The opnum field value for this method is 9.
	GetAddRoutesStatus(context.Context, *GetAddRoutesStatusRequest) (*GetAddRoutesStatusResponse, error)

	// Opnum10Reserved operation.
	// Opnum10Reserved

	// The CancelAddRoutesRequest method<32> stops packet loss and status monitoring for
	// Route elements previously added in a ProcessAddRoutes (section 3.6.4.6) invocation
	// and removes these routes from Route Collection.
	//
	// This method has no parameters.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.1 and 2.2.
	//
	//	+-------------------+----------------------------------+
	//	|      RETURN       |                                  |
	//	|    VALUE/CODE     |           DESCRIPTION            |
	//	|                   |                                  |
	//	+-------------------+----------------------------------+
	//	+-------------------+----------------------------------+
	//	| 0x00000000 S_OK   | The call was successful.         |
	//	+-------------------+----------------------------------+
	//	| 0x80004005 E_FAIL | Route Monitoring State is FALSE. |
	//	+-------------------+----------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	//
	// The opnum field value for this method is 11.
	CancelAddRoutesRequest(context.Context, *CancelAddRoutesRequestRequest) (*CancelAddRoutesRequestResponse, error)
}

func RegisterClusterNetwork2Server(conn dcerpc.Conn, o ClusterNetwork2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewClusterNetwork2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ClusterNetwork2SyntaxV0_0))...)
}

func NewClusterNetwork2ServerHandle(o ClusterNetwork2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ClusterNetwork2ServerHandle(ctx, o, opNum, r)
	}
}

func ClusterNetwork2ServerHandle(ctx context.Context, o ClusterNetwork2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // SendRTMessage
		in := &SendRTMessageRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SendRTMessage(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // InitializeNode
		in := &InitializeNodeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.InitializeNode(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // GetIpConfigSerialized
		in := &GetIPConfigSerializedRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetIPConfigSerialized(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // CleanupNode
		in := &CleanupNodeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CleanupNode(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // QueryFirewallConfiguration
		in := &QueryFirewallConfigurationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryFirewallConfiguration(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // ProcessAddRoutes
		in := &ProcessAddRoutesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ProcessAddRoutes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // GetAddRoutesStatus
		in := &GetAddRoutesStatusRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAddRoutesStatus(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // Opnum10Reserved
		// Opnum10Reserved
		return nil, nil
	case 11: // CancelAddRoutesRequest
		in := &CancelAddRoutesRequestRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CancelAddRoutesRequest(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
