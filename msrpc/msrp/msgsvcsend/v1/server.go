package msgsvcsend

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

// msgsvcsend server interface.
type MsgsvcsendServer interface {

	// The NetrSendMessage (Opnum 0) method is used to send a text message to a message
	// server.
	//
	// Return Values: An error_status_t value that indicates return status. If the method
	// returns a negative value, the method has failed. If the 12-bit facility code (bits
	// 16–27) is set to 0x007, the value contains a Win32 error code (defined in [MS-ERREF])
	// in the lower 16 bits. Zero or positive values indicate success, with the lower 16
	// bits in positive nonzero values containing warnings or flags defined in the method
	// implementation.
	//
	//	+------------------------------------+---------------------------------------------------------------------------+
	//	|               RETURN               |                                                                           |
	//	|             VALUE/CODE             |                                DESCRIPTION                                |
	//	|                                    |                                                                           |
	//	+------------------------------------+---------------------------------------------------------------------------+
	//	+------------------------------------+---------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | The operation completed successfully.                                     |
	//	+------------------------------------+---------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                                         |
	//	+------------------------------------+---------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The request is not supported.<22>                                         |
	//	+------------------------------------+---------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The parameter is incorrect.                                               |
	//	+------------------------------------+---------------------------------------------------------------------------+
	//	| 0x00000858 NERR_NetworkError       | A general network error occurred.                                         |
	//	+------------------------------------+---------------------------------------------------------------------------+
	//	| 0x000008E1 NERR_NameNotFound       | The message alias could not be found on the network.                      |
	//	+------------------------------------+---------------------------------------------------------------------------+
	//	| 0x000008E8 NERR_GrpMsgProcessor    | An error occurred in the domain message processor.                        |
	//	+------------------------------------+---------------------------------------------------------------------------+
	//	| 0x000008E9 NERR_PausedRemote       | The message was sent, but the recipient has paused the Messenger service. |
	//	+------------------------------------+---------------------------------------------------------------------------+
	//	| 0x000008EA NERR_BadReceive         | The message was sent but not received.                                    |
	//	+------------------------------------+---------------------------------------------------------------------------+
	//	| 0x000008EB NERR_NameInUse          | The message alias is currently in use. Try again later.                   |
	//	+------------------------------------+---------------------------------------------------------------------------+
	//	| 0x000008ED NERR_NotLocalName       | The name is not on the local computer.                                    |
	//	+------------------------------------+---------------------------------------------------------------------------+
	//	| 0x000008F1 NERR_TruncatedBroadcast | The broadcast message was truncated.                                      |
	//	+------------------------------------+---------------------------------------------------------------------------+
	//	| 0x000008F9 NERR_DuplicateName      | A duplicate message alias exists on the network.                          |
	//	+------------------------------------+---------------------------------------------------------------------------+
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
}

func RegisterMsgsvcsendServer(conn dcerpc.Conn, o MsgsvcsendServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewMsgsvcsendServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(MsgsvcsendSyntaxV1_0))...)
}

func NewMsgsvcsendServerHandle(o MsgsvcsendServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return MsgsvcsendServerHandle(ctx, o, opNum, r)
	}
}

func MsgsvcsendServerHandle(ctx context.Context, o MsgsvcsendServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // NetrSendMessage
		in := &SendMessageRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SendMessage(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
