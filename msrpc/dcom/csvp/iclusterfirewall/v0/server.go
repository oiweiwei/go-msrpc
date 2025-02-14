package iclusterfirewall

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

// IClusterFirewall server interface.
type ClusterFirewallServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The InitializeAdapterConfiguration method initializes the server Firewall State to
	// process subsequent calls of GetNextAdapterFirewallConfiguration.
	//
	// This method is called at least once before GetNextAdapterFirewallConfiguration.
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
	InitializeAdapterConfiguration(context.Context, *InitializeAdapterConfigurationRequest) (*InitializeAdapterConfigurationResponse, error)

	// The GetNextAdapterFirewallConfiguration method returns information about a specific
	// network adapter attached to the system.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                                  |
	//	|       VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK         | The call was successful.                                                         |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | The value the client specified in idx is greater than or equal                   |
	//	|                         | to the cRetAdapters value returned by the previous call to                       |
	//	|                         | InitializeAdapterConfiguration.                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x8000FFFF E_UNEXPECTED | InitializeAdapterConfiguration has not yet been called.                          |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 4.
	//
	// The server returns the following information to the client:
	//
	// * The output parameters set to the values specified previously.
	GetNextAdapterFirewallConfiguration(context.Context, *GetNextAdapterFirewallConfigurationRequest) (*GetNextAdapterFirewallConfigurationResponse, error)
}

func RegisterClusterFirewallServer(conn dcerpc.Conn, o ClusterFirewallServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewClusterFirewallServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ClusterFirewallSyntaxV0_0))...)
}

func NewClusterFirewallServerHandle(o ClusterFirewallServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ClusterFirewallServerHandle(ctx, o, opNum, r)
	}
}

func ClusterFirewallServerHandle(ctx context.Context, o ClusterFirewallServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // InitializeAdapterConfiguration
		op := &xxx_InitializeAdapterConfigurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InitializeAdapterConfigurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InitializeAdapterConfiguration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // GetNextAdapterFirewallConfiguration
		op := &xxx_GetNextAdapterFirewallConfigurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNextAdapterFirewallConfigurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNextAdapterFirewallConfiguration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IClusterFirewall
type UnimplementedClusterFirewallServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedClusterFirewallServer) InitializeAdapterConfiguration(context.Context, *InitializeAdapterConfigurationRequest) (*InitializeAdapterConfigurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterFirewallServer) GetNextAdapterFirewallConfiguration(context.Context, *GetNextAdapterFirewallConfigurationRequest) (*GetNextAdapterFirewallConfigurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ClusterFirewallServer = (*UnimplementedClusterFirewallServer)(nil)
