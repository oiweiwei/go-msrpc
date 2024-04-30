package iclusterlog

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

// IClusterLog server interface.
type ClusterLogServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GenerateClusterLog method writes a file that contains diagnostic information
	// about failover clusters for the server on which it executes. The content and format
	// of the file are implementation-specific, but SHOULD contain diagnostic information.
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
	GenerateClusterLog(context.Context, *GenerateClusterLogRequest) (*GenerateClusterLogResponse, error)

	// The GenerateTimeSpanLog method writes a file that contains diagnostic information
	// about failover clusters for the server on which it executes. The log entries in the
	// file date back only for the specified number of minutes. The content and format of
	// the file is implementation-specific, but SHOULD contain diagnostic information.
	//
	// Return Values: Return values are the same as the return values for the GenerateClusterLog
	// method specified in section 3.12.4.1.
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
	GenerateTimeSpanLog(context.Context, *GenerateTimeSpanLogRequest) (*GenerateTimeSpanLogResponse, error)

	// The GenerateClusterLogInLocalTime method<41> writes a file that contains diagnostic
	// information about failover clusters for the server on which it executes. The file
	// uses local time instead of GMT. The content and format of the file are implementation-specific
	// but SHOULD contain diagnostic information.
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
	GenerateClusterLogInLocalTime(context.Context, *GenerateClusterLogInLocalTimeRequest) (*GenerateClusterLogInLocalTimeResponse, error)

	// The GenerateTimeSpanLogInLocalTime method<42> writes a file that contains diagnostic
	// information about failover clusters for the server on which it executes. The log
	// entries in the file date back only for the specified number of minutes. The file
	// uses local time instead of GMT. The content and format of the file is implementation-specific
	// but SHOULD contain diagnostic information.
	//
	// Return Values: Return values are the same as the return values for the GenerateClusterLog
	// method specified in section 3.12.4.1.
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
	GenerateTimeSpanLogInLocalTime(context.Context, *GenerateTimeSpanLogInLocalTimeRequest) (*GenerateTimeSpanLogInLocalTimeResponse, error)
}

func RegisterClusterLogServer(conn dcerpc.Conn, o ClusterLogServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewClusterLogServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ClusterLogSyntaxV0_0))...)
}

func NewClusterLogServerHandle(o ClusterLogServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ClusterLogServerHandle(ctx, o, opNum, r)
	}
}

func ClusterLogServerHandle(ctx context.Context, o ClusterLogServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GenerateClusterLog
		in := &GenerateClusterLogRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GenerateClusterLog(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // GenerateTimeSpanLog
		in := &GenerateTimeSpanLogRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GenerateTimeSpanLog(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // GenerateClusterLogInLocalTime
		in := &GenerateClusterLogInLocalTimeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GenerateClusterLogInLocalTime(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // GenerateTimeSpanLogInLocalTime
		in := &GenerateTimeSpanLogInLocalTimeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GenerateTimeSpanLogInLocalTime(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
