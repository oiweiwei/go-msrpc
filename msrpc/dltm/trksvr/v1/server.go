package trksvr

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
)

// trksvr server interface.
type TrksvrServer interface {

	// The LnkSvrMessage method provides a way to send and receive messages to the DLT Central
	// Manager server to query or update information.
	//
	// Return Values: See the following table and the explanation after it for more information
	// on return values.
	//
	// Exceptions Thrown: None.
	//
	// The following table contains failure and success return values that have special
	// behavior in this protocol. All failure values not listed in this table MUST be treated
	// identically. Similarly, all success values not listed in this table MUST be treated
	// identically. Except where otherwise stated, a return value MUST NOT be a value from
	// this table. Except where otherwise specified, the server MUST return a success value.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                              |                                                                                  |
	//	|                CONSTANT/VALUE                |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| TRK_E_NOT_FOUND 0x8DEAD01B                   | A requested object was not found.                                                |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| TRK_E_VOLUME_QUOTA_EXCEEDED 0x8DEAD01C       | The server received a CREATE_VOLUME subrequest of a SYNC_VOLUMES request, but    |
	//	|                                              | the ServerVolumeTable size limit for the RequestMachine value has already been   |
	//	|                                              | reached.                                                                         |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| TRK_E_SERVER_TOO_BUSY 0x8DEAD01E             | The server is busy, and the client is to retry the request at a later time.      |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| TRK_S_OUT_OF_SYNC 0x0DEAD100                 | The VolumeSequenceNumber of a MOVE_NOTIFICATION request is incorrect. See        |
	//	|                                              | section 3.1.4.2.                                                                 |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| TRK_S_VOLUME_NOT_FOUND 0x0DEAD102            | The VolumeID in a request was not found in the server's ServerVolumeTable.       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| TRK_S_VOLUME_NOT_OWNED 0x0DEAD103            | A notification was sent to the LnkSvrMessage method, but the RequestMachine for  |
	//	|                                              | the request was not the VolumeOwner for a VolumeID in the request.               |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| TRK_S_NOTIFICATION_QUOTA_EXCEEDED 0x0DEAD107 | The server received a MOVE_NOTIFICATION request, but the FileTable size limit    |
	//	|                                              | has already been reached.                                                        |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The LnkSvrMessage method has only a single parameter, a union of type TRKSVR_MESSAGE_UNION
	// (see section 2.2.12). But that union is defined to hold one of several types of requests,
	// referred to in this protocol specification as messages. The message type for a given
	// request is specified in the MessageType field of the TRKSVR_MESSAGE_UNION. The possible
	// message types are defined in section 2.2.11. The formats of the different messages
	// are defined in the sub-sections of section 2.2.12. The responses by the server to
	// those different messages are specified in the remaining subsections of section 3.1.4,
	// according to the MessageType field of the union.
	//
	// Except where otherwise noted, the server that receives a request MUST ignore and
	// leave unmodified all fields in the TRKSVR_MESSAGE_UNION structure of the pMsg parameter,
	// as well as the structures referenced by the MessageUnion field of the TRKSVR_MESSAGE_UNION.
	//
	// For security purposes, the server SHOULD restrict access to clients that are members
	// of the authenticated users built-in security group. The client's identity is determined
	// as described in [MS-RPCE], section 3.3.3.4.3.<4>
	//
	// The TRKSVR_MESSAGE_UNION structure of the pMsg parameter contains a Priority field.
	// The server MAY use this value to decide to reject some requests with a TRK_E_SERVER_TOO_BUSY
	// return value, but it MUST NOT use this value to change the ordering of processing
	// of messages from a caller.<5>
	//
	// In this processing of this method call, the MachineID of the client that makes the
	// request MUST be used as the RequestMachine value.<6>
	//
	// Note  During the processing of a LnkSvrMessage call, the server can call back to
	// the client by using the LnkSvrMessageCallback method. See sections 3.1.4.4 and 3.2.4.1
	// for more information.
	ServerMessage(context.Context, *ServerMessageRequest) (*ServerMessageResponse, error)

	// The LnkSvrMessageCallback method is an RPC callback method that provides a means
	// for the DLT Central Manager server to call back to the client during a LnkSvrMessage
	// call. As defined in section 3.1.4, this callback only occurs for SYNC_VOLUMES messages
	// (for an example of this message, see section 3.2.5.3).
	//
	// For more details on when this callback is used by the server, see section 3.1.4.4.
	//
	// Return Values: The return value is typed as an HRESULT, but for this method, a value
	// of zero indicates success, and all other values indicate failure. Any nonzero value
	// MUST be treated identically as a failure value.
	//
	// The client MUST respond to this request by executing the steps in section 3.2.4.4
	// on each of the TRKSVR_SYNC_VOLUME structures in the TRKSVR_CALL_SYNC_VOLUMES structure
	// within the pMsg parameter. In this way, the client is responding as though it received
	// the updated structure in the completion of the LnkSvrMessage request.
	//
	// If any subrequest indicates a failure—that is, if the hr field of any TRKSVR_SYNC_VOLUME
	// structure is not zero—the client MUST return to the server with a return value
	// that indicates failure.
	//
	// For example, in a typical case where this callback method is used, processing proceeds
	// in the following order:
	//
	// * The client sends a SYNC_VOLUMES message to the server by calling LnkSvrMessage,
	// as described, for example, in section 3.2.5.3.
	//
	// * The server processes the request, updates the TRKSVR_CALL_SYNC_VOLUMES array in
	// the request, and calls LnkSvrMessageCallback on the client.
	//
	// * The client processes the subrequests in the updated TRKSVR_CALL_SYNC_VOLUMES array,
	// as defined in section 3.2.4.4.
	//
	// * The client returns from the LnkSvrMessageCallback method to the server.
	//
	// * The server sets the *cProcessed* field of the TRKSVR_CALL_SYNC_VOLUMES structure
	// to zero, and returns from the LnkSvrMessage method to the client.
	//
	// * The client again performs the processing defined in section 3.2.4.4. But because
	// the *cProcessed* field has been set to zero, the client takes no additional action,
	// as defined in that section.
	ServerMessageCallback(context.Context, *ServerMessageCallbackRequest) (*ServerMessageCallbackResponse, error)
}

func RegisterTrksvrServer(conn dcerpc.Conn, o TrksvrServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewTrksvrServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(TrksvrSyntaxV1_0))...)
}

func NewTrksvrServerHandle(o TrksvrServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return TrksvrServerHandle(ctx, o, opNum, r)
	}
}

func TrksvrServerHandle(ctx context.Context, o TrksvrServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // LnkSvrMessage
		op := &xxx_ServerMessageOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ServerMessageRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ServerMessage(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // LnkSvrMessageCallback
		op := &xxx_ServerMessageCallbackOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ServerMessageCallbackRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ServerMessageCallback(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented trksvr
type UnimplementedTrksvrServer struct {
}

func (UnimplementedTrksvrServer) ServerMessage(context.Context, *ServerMessageRequest) (*ServerMessageResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedTrksvrServer) ServerMessageCallback(context.Context, *ServerMessageCallbackRequest) (*ServerMessageCallbackResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ TrksvrServer = (*UnimplementedTrksvrServer)(nil)
