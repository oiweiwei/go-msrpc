package iclustercleanup

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

// IClusterCleanup server interface.
type ClusterCleanupServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The CleanUpEvictedNode method removes all persistent artifacts that exist on the
	// server after it is evicted from a cluster.
	//
	// This method is idempotent. After it is invoked, the target server can no longer be
	// a server for the Failover Cluster: Cluster Management Remote Protocol (ClusAPI) ([MS-CMRP])
	// until the server is reconfigured as a member of a cluster by using implementation-specific
	// methods between servers.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+-------------------------+---------------------------------------------------------------------------+
	//	|         RETURN          |                                                                           |
	//	|       VALUE/CODE        |                                DESCRIPTION                                |
	//	|                         |                                                                           |
	//	+-------------------------+---------------------------------------------------------------------------+
	//	+-------------------------+---------------------------------------------------------------------------+
	//	| 0x00000000 S_OK         | The call was successful.                                                  |
	//	+-------------------------+---------------------------------------------------------------------------+
	//	| 0x80070102 WAIT_TIMEOUT | The Cleanup Timer (section 3.8.2.2) expired before cleanup was completed. |
	//	+-------------------------+---------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 3.
	CleanupEvictedNode(context.Context, *CleanupEvictedNodeRequest) (*CleanupEvictedNodeResponse, error)

	// The ClearPR method performs a SCSI PERSISTENT RESERVE OUT command (see [SPC-3] section
	// 6.12) with a REGISTER AND IGNORE EXISTING KEY action, followed by a CLEAR action.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------+--------------------------+
	//	|             RETURN              |                          |
	//	|           VALUE/CODE            |       DESCRIPTION        |
	//	|                                 |                          |
	//	+---------------------------------+--------------------------+
	//	+---------------------------------+--------------------------+
	//	| 0x00000000 S_OK                 | The call was successful. |
	//	+---------------------------------+--------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND | The disk was not found.  |
	//	+---------------------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 4.
	ClearPR(context.Context, *ClearPRRequest) (*ClearPRResponse, error)
}

func RegisterClusterCleanupServer(conn dcerpc.Conn, o ClusterCleanupServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewClusterCleanupServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ClusterCleanupSyntaxV0_0))...)
}

func NewClusterCleanupServerHandle(o ClusterCleanupServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ClusterCleanupServerHandle(ctx, o, opNum, r)
	}
}

func ClusterCleanupServerHandle(ctx context.Context, o ClusterCleanupServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // CleanUpEvictedNode
		op := &xxx_CleanupEvictedNodeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CleanupEvictedNodeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CleanupEvictedNode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // ClearPR
		op := &xxx_ClearPROperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ClearPRRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ClearPR(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IClusterCleanup
type UnimplementedClusterCleanupServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedClusterCleanupServer) CleanupEvictedNode(context.Context, *CleanupEvictedNodeRequest) (*CleanupEvictedNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterCleanupServer) ClearPR(context.Context, *ClearPRRequest) (*ClearPRResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ClusterCleanupServer = (*UnimplementedClusterCleanupServer)(nil)
