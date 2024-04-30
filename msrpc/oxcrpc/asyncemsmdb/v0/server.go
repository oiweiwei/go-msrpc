package asyncemsmdb

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

// asyncemsmdb server interface.
type AsyncemsmdbServer interface {

	// The EcDoAsyncWaitEx method is an asynchronous call that the server does not complete
	// until events are pending on the Session Context, up to a 5-minute duration of no
	// client activity. If no events are available within 5 minutes of the time that the
	// client last accessed the server<36> through a call to the EcDoRpcExt2 method, as
	// specified in section 3.1.4.2, the server returns the call and does not set the NotificationPending
	// flag in the pulFlagsOut parameter. If an event is pending, the server completes the
	// call immediately and returns the NotificationPending flag in the pulFlagsOut parameter.
	// This call requires an active asynchronous context handle to be returned from the
	// EcDoAsyncConnectEx method on the EMSMDB interface, as specified in section 3.1.4.1.
	// The asynchronous context handle is associated with the Session Context.
	//
	// This method is part of notification handling. For more information about notifications,
	// see [MS-OXCNOTIF].
	//
	// Return Values: If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code or one of the protocol-defined
	// error codes listed in the following table.
	//
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	|   ERROR CODE    |            |                                                                                  |
	//	|      NAME       |   VALUE    |                                     MEANING                                      |
	//	|                 |            |                                                                                  |
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	| Rejected        | 0x000007EE | An EcDoAsyncWaitEx method call is already outstanding on this asynchronous       |
	//	|                 |            | context handle.<37>                                                              |
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	| Exiting         | 0x000003ED | The server is shutting down.                                                     |
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol as specified in [MS-RPCE].
	DoAsyncWaitEx(context.Context, *DoAsyncWaitExRequest) (*DoAsyncWaitExResponse, error)
}

func RegisterAsyncemsmdbServer(conn dcerpc.Conn, o AsyncemsmdbServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAsyncemsmdbServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AsyncemsmdbSyntaxV0_1))...)
}

func NewAsyncemsmdbServerHandle(o AsyncemsmdbServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AsyncemsmdbServerHandle(ctx, o, opNum, r)
	}
}

func AsyncemsmdbServerHandle(ctx context.Context, o AsyncemsmdbServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // EcDoAsyncWaitEx
		in := &DoAsyncWaitExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DoAsyncWaitEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
