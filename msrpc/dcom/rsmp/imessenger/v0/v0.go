package imessenger

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	rsmp "github.com/oiweiwei/go-msrpc/msrpc/dcom/rsmp"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = rsmp.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rsmp"
)

var (
	// IMessenger interface identifier 081e7188-c080-4ff3-9238-29f66d6cabfd
	MessengerIID = &dcom.IID{Data1: 0x081e7188, Data2: 0xc080, Data3: 0x4ff3, Data4: []byte{0x92, 0x38, 0x29, 0xf6, 0x6d, 0x6c, 0xab, 0xfd}}
	// Syntax UUID
	MessengerSyntaxUUID = &uuid.UUID{TimeLow: 0x81e7188, TimeMid: 0xc080, TimeHiAndVersion: 0x4ff3, ClockSeqHiAndReserved: 0x92, ClockSeqLow: 0x38, Node: [6]uint8{0x29, 0xf6, 0x6d, 0x6c, 0xab, 0xfd}}
	// Syntax ID
	MessengerSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: MessengerSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMessenger interface.
type MessengerClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The SendMessage method adds a message to the send queue.
	SendMessage(context.Context, *SendMessageRequest, ...dcerpc.CallOption) (*SendMessageResponse, error)

	// The RecallMessage method retrieves a message from the send queue.
	RecallMessage(context.Context, *RecallMessageRequest, ...dcerpc.CallOption) (*RecallMessageResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) MessengerClient
}

type xxx_DefaultMessengerClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultMessengerClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultMessengerClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...dcerpc.CallOption) (*SendMessageResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SendMessageResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessengerClient) RecallMessage(ctx context.Context, in *RecallMessageRequest, opts ...dcerpc.CallOption) (*RecallMessageResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RecallMessageResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultMessengerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultMessengerClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultMessengerClient) IPID(ctx context.Context, ipid *dcom.IPID) MessengerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultMessengerClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewMessengerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (MessengerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(MessengerSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iunknown.NewUnknownClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultMessengerClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_SendMessageOperation structure represents the SendMessage operation
type xxx_SendMessageOperation struct {
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat   `idl:"name:That" json:"that"`
	RSMMessage *rsmp.RSMMessage `idl:"name:lpRsmMessage;pointer:unique" json:"rsm_message"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_SendMessageOperation) OpNum() int { return 3 }

func (o *xxx_SendMessageOperation) OpName() string { return "/IMessenger/v0/SendMessage" }

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
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpRsmMessage {in} (1:{pointer=unique, alias=LPRSM_MESSAGE}*(1))(2:{alias=RSM_MESSAGE}(struct))
	{
		if o.RSMMessage != nil {
			_ptr_lpRsmMessage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RSMMessage != nil {
					if err := o.RSMMessage.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&rsmp.RSMMessage{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RSMMessage, _ptr_lpRsmMessage); err != nil {
				return err
			}
		} else {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendMessageOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpRsmMessage {in} (1:{pointer=unique, alias=LPRSM_MESSAGE}*(1))(2:{alias=RSM_MESSAGE}(struct))
	{
		_ptr_lpRsmMessage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RSMMessage == nil {
				o.RSMMessage = &rsmp.RSMMessage{}
			}
			if err := o.RSMMessage.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpRsmMessage := func(ptr interface{}) { o.RSMMessage = *ptr.(**rsmp.RSMMessage) }
		if err := w.ReadPointer(&o.RSMMessage, _s_lpRsmMessage, _ptr_lpRsmMessage); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
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
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_SendMessageOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
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

// SendMessageRequest structure represents the SendMessage operation request
type SendMessageRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpRsmMessage: A pointer to an RSM_MESSAGE (section 2.2.6.1) structure describing
	// the message that is to be sent.
	//
	//	+------------------------------------+--------------------------+
	//	|               RETURN               |                          |
	//	|             VALUE/CODE             |       DESCRIPTION        |
	//	|                                    |                          |
	//	+------------------------------------+--------------------------+
	//	+------------------------------------+--------------------------+
	//	| 0x00000000 S_OK                    | The call was successful. |
	//	+------------------------------------+--------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | A parameter is missing.  |
	//	+------------------------------------+--------------------------+
	RSMMessage *rsmp.RSMMessage `idl:"name:lpRsmMessage;pointer:unique" json:"rsm_message"`
}

func (o *SendMessageRequest) xxx_ToOp(ctx context.Context, op *xxx_SendMessageOperation) *xxx_SendMessageOperation {
	if op == nil {
		op = &xxx_SendMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.RSMMessage = o.RSMMessage
	return op
}

func (o *SendMessageRequest) xxx_FromOp(ctx context.Context, op *xxx_SendMessageOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RSMMessage = op.RSMMessage
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

// SendMessageResponse structure represents the SendMessage operation response
type SendMessageResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SendMessage return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SendMessageResponse) xxx_ToOp(ctx context.Context, op *xxx_SendMessageOperation) *xxx_SendMessageOperation {
	if op == nil {
		op = &xxx_SendMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SendMessageResponse) xxx_FromOp(ctx context.Context, op *xxx_SendMessageOperation) {
	if o == nil {
		return
	}
	o.That = op.That
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

// xxx_RecallMessageOperation structure represents the RecallMessage operation
type xxx_RecallMessageOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	GUID   *dtyp.GUID     `idl:"name:lpGuid" json:"guid"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RecallMessageOperation) OpNum() int { return 4 }

func (o *xxx_RecallMessageOperation) OpName() string { return "/IMessenger/v0/RecallMessage" }

func (o *xxx_RecallMessageOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RecallMessageOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpGuid {in} (1:{alias=LPGUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID != nil {
			if err := o.GUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_RecallMessageOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpGuid {in} (1:{alias=LPGUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.GUID == nil {
			o.GUID = &dtyp.GUID{}
		}
		if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RecallMessageOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RecallMessageOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RecallMessageOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
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

// RecallMessageRequest structure represents the RecallMessage operation request
type RecallMessageRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lpGuid: A pointer to the identifier of the message to retrieve.
	//
	//	+------------------------------------+--------------------------+
	//	|               RETURN               |                          |
	//	|             VALUE/CODE             |       DESCRIPTION        |
	//	|                                    |                          |
	//	+------------------------------------+--------------------------+
	//	+------------------------------------+--------------------------+
	//	| 0x00000000 S_OK                    | The call was successful. |
	//	+------------------------------------+--------------------------+
	//	| 0x80070057 ERROR_INVALID_PARAMETER | A parameter is missing.  |
	//	+------------------------------------+--------------------------+
	//
	// After the server receives this message, it MUST verify that lpGuid is not equal to
	// NULL. If parameter validation fails, the server MUST immediately fail the operation
	// and return ERROR_INVALID_PARAMETER (0x80070057).
	//
	// If parameter validation succeeds, the server MUST perform the following actions:
	//
	// * Create an event for synchronization.
	//
	// * Get the Destination List from the Sent Message List, and verify that it is not
	// NULL.
	//
	// * Find the message in the Sent Message List.
	//
	// * Remove the message from the list.
	//
	// If there are no more active requests, the server MUST hide any existing notifications,
	// destroy the handler object and its corresponding dialog, and free the message object.
	GUID *dtyp.GUID `idl:"name:lpGuid" json:"guid"`
}

func (o *RecallMessageRequest) xxx_ToOp(ctx context.Context, op *xxx_RecallMessageOperation) *xxx_RecallMessageOperation {
	if op == nil {
		op = &xxx_RecallMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.GUID = o.GUID
	return op
}

func (o *RecallMessageRequest) xxx_FromOp(ctx context.Context, op *xxx_RecallMessageOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.GUID = op.GUID
}
func (o *RecallMessageRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RecallMessageRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RecallMessageOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RecallMessageResponse structure represents the RecallMessage operation response
type RecallMessageResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RecallMessage return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RecallMessageResponse) xxx_ToOp(ctx context.Context, op *xxx_RecallMessageOperation) *xxx_RecallMessageOperation {
	if op == nil {
		op = &xxx_RecallMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RecallMessageResponse) xxx_FromOp(ctx context.Context, op *xxx_RecallMessageOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RecallMessageResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RecallMessageResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RecallMessageOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
