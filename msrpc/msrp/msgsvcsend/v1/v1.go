package msgsvcsend

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

var (
	// import guard
	GoPackage = "msrp"
)

var (
	// Syntax UUID
	MsgsvcsendSyntaxUUID = &uuid.UUID{TimeLow: 0x5a7b91f8, TimeMid: 0xff00, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xa9, ClockSeqLow: 0xb2, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xb6, 0xe6, 0xfc}}
	// Syntax ID
	MsgsvcsendSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: MsgsvcsendSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// msgsvcsend interface.
type MsgsvcsendClient interface {

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
	SendMessage(context.Context, *SendMessageRequest, ...dcerpc.CallOption) (*SendMessageResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultMsgsvcsendClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultMsgsvcsendClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...dcerpc.CallOption) (*SendMessageResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SendMessageResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMsgsvcsendClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultMsgsvcsendClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewMsgsvcsendClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (MsgsvcsendClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(MsgsvcsendSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultMsgsvcsendClient{cc: cc}, nil
}

// xxx_SendMessageOperation structure represents the NetrSendMessage operation
type xxx_SendMessageOperation struct {
	From   string `idl:"name:From;string" json:"from"`
	To     string `idl:"name:To;string" json:"to"`
	Text   string `idl:"name:Text;string" json:"text"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SendMessageOperation) OpNum() int { return 0 }

func (o *xxx_SendMessageOperation) OpName() string { return "/msgsvcsend/v1/NetrSendMessage" }

func (o *xxx_SendMessageOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendMessageOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// From {in} (1:{string, alias=LPSTR}*(1)[dim:0,string,null](char))
	{
		if err := ndr.WriteCharNString(ctx, w, o.From); err != nil {
			return err
		}
	}
	// To {in} (1:{string, alias=LPSTR}*(1)[dim:0,string,null](char))
	{
		if err := ndr.WriteCharNString(ctx, w, o.To); err != nil {
			return err
		}
	}
	// Text {in} (1:{string, alias=LPSTR}*(1)[dim:0,string,null](char))
	{
		if err := ndr.WriteCharNString(ctx, w, o.Text); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendMessageOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// From {in} (1:{string, alias=LPSTR,pointer=ref}*(1)[dim:0,string,null](char))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.From); err != nil {
			return err
		}
	}
	// To {in} (1:{string, alias=LPSTR,pointer=ref}*(1)[dim:0,string,null](char))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.To); err != nil {
			return err
		}
	}
	// Text {in} (1:{string, alias=LPSTR,pointer=ref}*(1)[dim:0,string,null](char))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.Text); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendMessageOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendMessageOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendMessageOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SendMessageRequest structure represents the NetrSendMessage operation request
type SendMessageRequest struct {
	// From: A null-terminated string that MUST denote the name of the sender of the message.
	// The name is not guaranteed to be unique or reachable by this method. The string MUST
	// be expressed in the original equipment manufacturer (OEM) character set, as specified
	// in [MS-UCODEREF] section 2.2.1, of the invoker of this method.
	From string `idl:"name:From;string" json:"from"`
	// To: A null-terminated string that MUST represent the name of the intended recipient
	// of the message. The name is not guaranteed to be unique or reachable by this method.
	// The string is expressed in the OEM character set, as specified in [MS-UCODEREF] section
	// 2.2.1, of the invoker of this method.
	To string `idl:"name:To;string" json:"to"`
	// Text: A null-terminated string that MUST contain the message that is being sent to
	// the recipient in the To parameter. The string is expressed in the OEM character set,
	// as specified in [MS-UCODEREF] section 2.2.1.
	Text string `idl:"name:Text;string" json:"text"`
}

func (o *SendMessageRequest) xxx_ToOp(ctx context.Context, op *xxx_SendMessageOperation) *xxx_SendMessageOperation {
	if op == nil {
		op = &xxx_SendMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.From = o.From
	op.To = o.To
	op.Text = o.Text
	return op
}

func (o *SendMessageRequest) xxx_FromOp(ctx context.Context, op *xxx_SendMessageOperation) {
	if o == nil {
		return
	}
	o.From = op.From
	o.To = op.To
	o.Text = op.Text
}
func (o *SendMessageRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SendMessageRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SendMessageOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SendMessageResponse structure represents the NetrSendMessage operation response
type SendMessageResponse struct {
	// Return: The NetrSendMessage return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SendMessageResponse) xxx_ToOp(ctx context.Context, op *xxx_SendMessageOperation) *xxx_SendMessageOperation {
	if op == nil {
		op = &xxx_SendMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SendMessageResponse) xxx_FromOp(ctx context.Context, op *xxx_SendMessageOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SendMessageResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SendMessageResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SendMessageOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
