package trksvr

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dltm "github.com/oiweiwei/go-msrpc/msrpc/dltm"
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
	_ = dltm.GoPackage
)

var (
	// import guard
	GoPackage = "dltm"
)

var (
	// Syntax UUID
	TrksvrSyntaxUUID = &uuid.UUID{TimeLow: 0x4da1c422, TimeMid: 0x943d, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xac, ClockSeqLow: 0xae, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xc2, 0xaa, 0x3f}}
	// Syntax ID
	TrksvrSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: TrksvrSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// trksvr interface.
type TrksvrClient interface {

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
	ServerMessage(context.Context, *ServerMessageRequest, ...dcerpc.CallOption) (*ServerMessageResponse, error)

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
	ServerMessageCallback(context.Context, *ServerMessageCallbackRequest, ...dcerpc.CallOption) (*ServerMessageCallbackResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultTrksvrClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultTrksvrClient) ServerMessage(ctx context.Context, in *ServerMessageRequest, opts ...dcerpc.CallOption) (*ServerMessageResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ServerMessageResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTrksvrClient) ServerMessageCallback(ctx context.Context, in *ServerMessageCallbackRequest, opts ...dcerpc.CallOption) (*ServerMessageCallbackResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ServerMessageCallbackResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTrksvrClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultTrksvrClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewTrksvrClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (TrksvrClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(TrksvrSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultTrksvrClient{cc: cc}, nil
}

// xxx_ServerMessageOperation structure represents the LnkSvrMessage operation
type xxx_ServerMessageOperation struct {
	Message *dltm.MessageUnion `idl:"name:pMsg" json:"message"`
	Return  int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_ServerMessageOperation) OpNum() int { return 0 }

func (o *xxx_ServerMessageOperation) OpName() string { return "/trksvr/v1/LnkSvrMessage" }

func (o *xxx_ServerMessageOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerMessageOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pMsg {in, out} (1:{pointer=ref}*(1))(2:{alias=TRKSVR_MESSAGE_UNION}(struct))
	{
		if o.Message != nil {
			if err := o.Message.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dltm.MessageUnion{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerMessageOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pMsg {in, out} (1:{pointer=ref}*(1))(2:{alias=TRKSVR_MESSAGE_UNION}(struct))
	{
		if o.Message == nil {
			o.Message = &dltm.MessageUnion{}
		}
		if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerMessageOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerMessageOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pMsg {in, out} (1:{pointer=ref}*(1))(2:{alias=TRKSVR_MESSAGE_UNION}(struct))
	{
		if o.Message != nil {
			if err := o.Message.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dltm.MessageUnion{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerMessageOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pMsg {in, out} (1:{pointer=ref}*(1))(2:{alias=TRKSVR_MESSAGE_UNION}(struct))
	{
		if o.Message == nil {
			o.Message = &dltm.MessageUnion{}
		}
		if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ServerMessageRequest structure represents the LnkSvrMessage operation request
type ServerMessageRequest struct {
	// pMsg: Pointer to a message, in the format of a TRKSVR_MESSAGE_UNION structure. If
	// this method fails, as indicated by a failure return value, the client MUST ignore
	// any changes made by the server to this structure.
	Message *dltm.MessageUnion `idl:"name:pMsg" json:"message"`
}

func (o *ServerMessageRequest) xxx_ToOp(ctx context.Context, op *xxx_ServerMessageOperation) *xxx_ServerMessageOperation {
	if op == nil {
		op = &xxx_ServerMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.Message = o.Message
	return op
}

func (o *ServerMessageRequest) xxx_FromOp(ctx context.Context, op *xxx_ServerMessageOperation) {
	if o == nil {
		return
	}
	o.Message = op.Message
}
func (o *ServerMessageRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ServerMessageRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerMessageOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ServerMessageResponse structure represents the LnkSvrMessage operation response
type ServerMessageResponse struct {
	// pMsg: Pointer to a message, in the format of a TRKSVR_MESSAGE_UNION structure. If
	// this method fails, as indicated by a failure return value, the client MUST ignore
	// any changes made by the server to this structure.
	Message *dltm.MessageUnion `idl:"name:pMsg" json:"message"`
	// Return: The LnkSvrMessage return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ServerMessageResponse) xxx_ToOp(ctx context.Context, op *xxx_ServerMessageOperation) *xxx_ServerMessageOperation {
	if op == nil {
		op = &xxx_ServerMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *ServerMessageResponse) xxx_FromOp(ctx context.Context, op *xxx_ServerMessageOperation) {
	if o == nil {
		return
	}
	o.Message = op.Message
	o.Return = op.Return
}
func (o *ServerMessageResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ServerMessageResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerMessageOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ServerMessageCallbackOperation structure represents the LnkSvrMessageCallback operation
type xxx_ServerMessageCallbackOperation struct {
	Message *dltm.MessageUnion `idl:"name:pMsg" json:"message"`
	Return  int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_ServerMessageCallbackOperation) OpNum() int { return 1 }

func (o *xxx_ServerMessageCallbackOperation) OpName() string {
	return "/trksvr/v1/LnkSvrMessageCallback"
}

func (o *xxx_ServerMessageCallbackOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerMessageCallbackOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pMsg {in, out} (1:{pointer=ref}*(1))(2:{alias=TRKSVR_MESSAGE_UNION}(struct))
	{
		if o.Message != nil {
			if err := o.Message.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dltm.MessageUnion{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerMessageCallbackOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pMsg {in, out} (1:{pointer=ref}*(1))(2:{alias=TRKSVR_MESSAGE_UNION}(struct))
	{
		if o.Message == nil {
			o.Message = &dltm.MessageUnion{}
		}
		if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerMessageCallbackOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerMessageCallbackOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pMsg {in, out} (1:{pointer=ref}*(1))(2:{alias=TRKSVR_MESSAGE_UNION}(struct))
	{
		if o.Message != nil {
			if err := o.Message.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dltm.MessageUnion{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerMessageCallbackOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pMsg {in, out} (1:{pointer=ref}*(1))(2:{alias=TRKSVR_MESSAGE_UNION}(struct))
	{
		if o.Message == nil {
			o.Message = &dltm.MessageUnion{}
		}
		if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ServerMessageCallbackRequest structure represents the LnkSvrMessageCallback operation request
type ServerMessageCallbackRequest struct {
	// pMsg: Pointer to a message, in the format of a TRKSVR_MESSAGE_UNION structure. The
	// MessageType field in this structure MUST be set to SYNC_VOLUMES.
	Message *dltm.MessageUnion `idl:"name:pMsg" json:"message"`
}

func (o *ServerMessageCallbackRequest) xxx_ToOp(ctx context.Context, op *xxx_ServerMessageCallbackOperation) *xxx_ServerMessageCallbackOperation {
	if op == nil {
		op = &xxx_ServerMessageCallbackOperation{}
	}
	if o == nil {
		return op
	}
	op.Message = o.Message
	return op
}

func (o *ServerMessageCallbackRequest) xxx_FromOp(ctx context.Context, op *xxx_ServerMessageCallbackOperation) {
	if o == nil {
		return
	}
	o.Message = op.Message
}
func (o *ServerMessageCallbackRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ServerMessageCallbackRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerMessageCallbackOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ServerMessageCallbackResponse structure represents the LnkSvrMessageCallback operation response
type ServerMessageCallbackResponse struct {
	// pMsg: Pointer to a message, in the format of a TRKSVR_MESSAGE_UNION structure. The
	// MessageType field in this structure MUST be set to SYNC_VOLUMES.
	Message *dltm.MessageUnion `idl:"name:pMsg" json:"message"`
	// Return: The LnkSvrMessageCallback return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ServerMessageCallbackResponse) xxx_ToOp(ctx context.Context, op *xxx_ServerMessageCallbackOperation) *xxx_ServerMessageCallbackOperation {
	if op == nil {
		op = &xxx_ServerMessageCallbackOperation{}
	}
	if o == nil {
		return op
	}
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *ServerMessageCallbackResponse) xxx_FromOp(ctx context.Context, op *xxx_ServerMessageCallbackOperation) {
	if o == nil {
		return
	}
	o.Message = op.Message
	o.Return = op.Return
}
func (o *ServerMessageCallbackResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ServerMessageCallbackResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerMessageCallbackOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
