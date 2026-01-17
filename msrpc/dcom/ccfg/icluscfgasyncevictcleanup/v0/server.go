package icluscfgasyncevictcleanup

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = idispatch.GoPackage
)

// IClusCfgAsyncEvictCleanup server interface.
type AsyncEvictCleanupServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The CleanupNode method removes all persistent artifacts that exist on the node after
	// it is evicted from a cluster.
	//
	// The Failover Cluster: Management API (ClusAPI) Protocol server provides a method
	// to evict a node from a cluster. Evicting a node from a cluster is specified in [MS-CMRP],
	// section 3.1.1.5. Once evicted, the node can be restored to its precluster installation
	// state.
	//
	// This method is idempotent. After it is invoked, the following actions MUST occur.
	//
	// * The target node MUST no longer be a server for the Failover Cluster: Management
	// API (ClusAPI) Protocol (as specified in [MS-CMRP]) until the node is reconfigured
	// as a member of a cluster by using implementation-specific methods between servers.
	//
	// * In any subsequent query of the ClusterInstallationState Registry Value, by means
	// of the Windows Remote Registry Protocol ( [MS-RRP] ( ../ms-rrp/0fa3191d-bb79-490a-81bd-54c2601b7a78
	// ) ), as specified in [MS-CMRP] section 3.1.3.1, the server MUST return that the Value
	// is set to 0x00000001 (eClusterInstallStateFilesCopied).
	//
	// * In any subsequent calls of the Service Control Manager Remote Protocol ( [MS-SCMR]
	// ( ../ms-scmr/705b624a-13de-43cc-b8a2-99573da3635f ) ) OpenService method for service
	// name "ClusSvc", the server MUST complete with error 1060 (ERROR_SERVICE_DOES_NOT_EXIST).
	// This behavior is in contrast to the behavior specified in the first bullet of [MS-CMRP]
	// section 3.1.3.2.
	//
	// * Reset any other implementation-specific values to their precluster installation
	// state.
	//
	// This method MUST NOT be invoked while the node is a configured member of a cluster.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, the method has failed. If the 12-bit facility code (bits
	// 16-27) is set to 0x007, the value contains a Win32 error code (defined in [MS-ERREF])
	// in the lower 16 bits. A zero value or positive values indicate success, with the
	// lower 16 bits in positive nonzero values containing warnings or flags that are defined
	// in the method implementation.
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
	// Exceptions thrown:
	//
	// No exceptions are thrown except those that are thrown by the underlying RPC protocol,
	// as specified in [MS-RPCE].
	CleanupNode(context.Context, *CleanupNodeRequest) (*CleanupNodeResponse, error)
}

func RegisterAsyncEvictCleanupServer(conn dcerpc.Conn, o AsyncEvictCleanupServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAsyncEvictCleanupServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AsyncEvictCleanupSyntaxV0_0))...)
}

func NewAsyncEvictCleanupServerHandle(o AsyncEvictCleanupServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AsyncEvictCleanupServerHandle(ctx, o, opNum, r)
	}
}

func AsyncEvictCleanupServerHandle(ctx context.Context, o AsyncEvictCleanupServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // CleanupNode
		op := &xxx_CleanupNodeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CleanupNodeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CleanupNode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IClusCfgAsyncEvictCleanup
type UnimplementedAsyncEvictCleanupServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedAsyncEvictCleanupServer) CleanupNode(context.Context, *CleanupNodeRequest) (*CleanupNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AsyncEvictCleanupServer = (*UnimplementedAsyncEvictCleanupServer)(nil)
