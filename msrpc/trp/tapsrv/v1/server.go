package tapsrv

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// tapsrv server interface.
type TapsrvServer interface {

	// The ClientAttach method is called by the client to establish a binding instance with
	// the server.
	//
	// Return Values: The method returns 0 on success; otherwise, it returns a nonzero error
	// code, as specified in [MS-ERREF]. The following table includes a common error code.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80000048 LINEERR_OPERATIONFAILED | Generic error on the server.                                                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|                                -19 | Requesting administrator access via lProcessId equals 0xFFFFFFFD (-3), but the   |
	//	|                                    | user credentials of the client do not have administrator access on the server.   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	//
	// The opnum field value for this method is 0.
	//
	// Either the client or the server can reject unencrypted packets based on the configuration.<6>
	ClientAttach(context.Context, *ClientAttachRequest) (*ClientAttachResponse, error)

	// The ClientRequest method is called by the client to submit requests to the server.
	//
	// Return Values: This method has no return values. However, the status of the request
	// is encapsulated within the pBuffer parameter and contained in the TAPI32_MSG.Ack_ReturnValue
	// field.
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol as specified
	// in [MS-RPCE].
	//
	// The opnum field value for this method is 1.
	ClientRequest(context.Context, *ClientRequestRequest) (*ClientRequestResponse, error)

	// The ClientDetach method is called by a client when it finishes using the telephony
	// resources on a server. In response, the server frees the referenced binding instance
	// and releases the allocated resources associated with the client. For connection-oriented
	// clients, the server then calls RemoteSPDetach on the client to release the allocated
	// resources created on the client side.
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	//
	// The opnum field value for this method is 2.
	ClientDetach(context.Context, *ClientDetachRequest) (*ClientDetachResponse, error)
}

func RegisterTapsrvServer(conn dcerpc.Conn, o TapsrvServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTapsrvServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TapsrvSyntaxV1_0))...)
}

func NewTapsrvServerHandle(o TapsrvServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TapsrvServerHandle(ctx, o, opNum, r)
	}
}

func TapsrvServerHandle(ctx context.Context, o TapsrvServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // ClientAttach
		in := &ClientAttachRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ClientAttach(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // ClientRequest
		in := &ClientRequestRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ClientRequest(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // ClientDetach
		in := &ClientDetachRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ClientDetach(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
