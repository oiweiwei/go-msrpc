package iclusterupdate

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

// IClusterUpdate server interface.
type ClusterUpdateServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GetUpdates method queries the local server for all of the updates that are installed
	// on the local server.
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
	GetUpdates(context.Context, *GetUpdatesRequest) (*GetUpdatesResponse, error)

	// The Count method returns the number of updates that are installed on the local server.
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
	Count(context.Context, *CountRequest) (*CountResponse, error)
}

func RegisterClusterUpdateServer(conn dcerpc.Conn, o ClusterUpdateServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewClusterUpdateServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ClusterUpdateSyntaxV0_0))...)
}

func NewClusterUpdateServerHandle(o ClusterUpdateServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ClusterUpdateServerHandle(ctx, o, opNum, r)
	}
}

func ClusterUpdateServerHandle(ctx context.Context, o ClusterUpdateServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetUpdates
		op := &xxx_GetUpdatesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetUpdatesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetUpdates(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // Count
		op := &xxx_CountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Count(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IClusterUpdate
type UnimplementedClusterUpdateServer struct {
	iunknown.UnimplementedUnknownServer
}

func (UnimplementedClusterUpdateServer) GetUpdates(context.Context, *GetUpdatesRequest) (*GetUpdatesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusterUpdateServer) Count(context.Context, *CountRequest) (*CountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ ClusterUpdateServer = (*UnimplementedClusterUpdateServer)(nil)
