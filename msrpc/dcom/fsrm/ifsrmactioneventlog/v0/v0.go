package ifsrmactioneventlog

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	fsrm "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm"
	ifsrmaction "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmaction/v0"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
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
	_ = ifsrmaction.GoPackage
	_ = fsrm.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmActionEventLog interface identifier 4c8f96c3-5d94-4f37-a4f4-f56ab463546f
	ActionEventLogIID = &dcom.IID{Data1: 0x4c8f96c3, Data2: 0x5d94, Data3: 0x4f37, Data4: []byte{0xa4, 0xf4, 0xf5, 0x6a, 0xb4, 0x63, 0x54, 0x6f}}
	// Syntax UUID
	ActionEventLogSyntaxUUID = &uuid.UUID{TimeLow: 0x4c8f96c3, TimeMid: 0x5d94, TimeHiAndVersion: 0x4f37, ClockSeqHiAndReserved: 0xa4, ClockSeqLow: 0xf4, Node: [6]uint8{0xf5, 0x6a, 0xb4, 0x63, 0x54, 0x6f}}
	// Syntax ID
	ActionEventLogSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ActionEventLogSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmActionEventLog interface.
type ActionEventLogClient interface {

	// IFsrmAction retrieval method.
	Action() ifsrmaction.ActionClient

	// EventType operation.
	GetEventType(context.Context, *GetEventTypeRequest, ...dcerpc.CallOption) (*GetEventTypeResponse, error)

	// EventType operation.
	SetEventType(context.Context, *SetEventTypeRequest, ...dcerpc.CallOption) (*SetEventTypeResponse, error)

	// MessageText operation.
	GetMessageText(context.Context, *GetMessageTextRequest, ...dcerpc.CallOption) (*GetMessageTextResponse, error)

	// MessageText operation.
	SetMessageText(context.Context, *SetMessageTextRequest, ...dcerpc.CallOption) (*SetMessageTextResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ActionEventLogClient
}

type xxx_DefaultActionEventLogClient struct {
	ifsrmaction.ActionClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultActionEventLogClient) Action() ifsrmaction.ActionClient {
	return o.ActionClient
}

func (o *xxx_DefaultActionEventLogClient) GetEventType(ctx context.Context, in *GetEventTypeRequest, opts ...dcerpc.CallOption) (*GetEventTypeResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetEventTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionEventLogClient) SetEventType(ctx context.Context, in *SetEventTypeRequest, opts ...dcerpc.CallOption) (*SetEventTypeResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &SetEventTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionEventLogClient) GetMessageText(ctx context.Context, in *GetMessageTextRequest, opts ...dcerpc.CallOption) (*GetMessageTextResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetMessageTextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionEventLogClient) SetMessageText(ctx context.Context, in *SetMessageTextRequest, opts ...dcerpc.CallOption) (*SetMessageTextResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &SetMessageTextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionEventLogClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultActionEventLogClient) IPID(ctx context.Context, ipid *dcom.IPID) ActionEventLogClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultActionEventLogClient{
		ActionClient: o.ActionClient.IPID(ctx, ipid),
		cc:           o.cc,
		ipid:         ipid,
	}
}
func NewActionEventLogClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ActionEventLogClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ActionEventLogSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ifsrmaction.NewActionClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultActionEventLogClient{
		ActionClient: base,
		cc:           cc,
		ipid:         ipid,
	}, nil
}

// xxx_GetEventTypeOperation structure represents the EventType operation
type xxx_GetEventTypeOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventType fsrm.EventType `idl:"name:eventType" json:"event_type"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEventTypeOperation) OpNum() int { return 12 }

func (o *xxx_GetEventTypeOperation) OpName() string { return "/IFsrmActionEventLog/v0/EventType" }

func (o *xxx_GetEventTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_GetEventTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_GetEventTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// eventType {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmEventType}(enum))
	{
		if err := w.WriteData(uint16(o.EventType)); err != nil {
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

func (o *xxx_GetEventTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// eventType {out, retval} (1:{pointer=ref}*(1))(2:{alias=FsrmEventType}(enum))
	{
		if err := w.ReadData((*uint16)(&o.EventType)); err != nil {
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

// GetEventTypeRequest structure represents the EventType operation request
type GetEventTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetEventTypeRequest) xxx_ToOp(ctx context.Context) *xxx_GetEventTypeOperation {
	if o == nil {
		return &xxx_GetEventTypeOperation{}
	}
	return &xxx_GetEventTypeOperation{
		This: o.This,
	}
}

func (o *GetEventTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEventTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetEventTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetEventTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEventTypeResponse structure represents the EventType operation response
type GetEventTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventType fsrm.EventType `idl:"name:eventType" json:"event_type"`
	// Return: The EventType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetEventTypeResponse) xxx_ToOp(ctx context.Context) *xxx_GetEventTypeOperation {
	if o == nil {
		return &xxx_GetEventTypeOperation{}
	}
	return &xxx_GetEventTypeOperation{
		That:      o.That,
		EventType: o.EventType,
		Return:    o.Return,
	}
}

func (o *GetEventTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEventTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.EventType = op.EventType
	o.Return = op.Return
}
func (o *GetEventTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetEventTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetEventTypeOperation structure represents the EventType operation
type xxx_SetEventTypeOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventType fsrm.EventType `idl:"name:eventType" json:"event_type"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetEventTypeOperation) OpNum() int { return 13 }

func (o *xxx_SetEventTypeOperation) OpName() string { return "/IFsrmActionEventLog/v0/EventType" }

func (o *xxx_SetEventTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// eventType {in} (1:{alias=FsrmEventType}(enum))
	{
		if err := w.WriteData(uint16(o.EventType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// eventType {in} (1:{alias=FsrmEventType}(enum))
	{
		if err := w.ReadData((*uint16)(&o.EventType)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetEventTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetEventTypeRequest structure represents the EventType operation request
type SetEventTypeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	EventType fsrm.EventType `idl:"name:eventType" json:"event_type"`
}

func (o *SetEventTypeRequest) xxx_ToOp(ctx context.Context) *xxx_SetEventTypeOperation {
	if o == nil {
		return &xxx_SetEventTypeOperation{}
	}
	return &xxx_SetEventTypeOperation{
		This:      o.This,
		EventType: o.EventType,
	}
}

func (o *SetEventTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_SetEventTypeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.EventType = op.EventType
}
func (o *SetEventTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetEventTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEventTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetEventTypeResponse structure represents the EventType operation response
type SetEventTypeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The EventType return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetEventTypeResponse) xxx_ToOp(ctx context.Context) *xxx_SetEventTypeOperation {
	if o == nil {
		return &xxx_SetEventTypeOperation{}
	}
	return &xxx_SetEventTypeOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetEventTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_SetEventTypeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetEventTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetEventTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEventTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMessageTextOperation structure represents the MessageText operation
type xxx_GetMessageTextOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MessageText *oaut.String   `idl:"name:messageText" json:"message_text"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMessageTextOperation) OpNum() int { return 14 }

func (o *xxx_GetMessageTextOperation) OpName() string { return "/IFsrmActionEventLog/v0/MessageText" }

func (o *xxx_GetMessageTextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMessageTextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_GetMessageTextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_GetMessageTextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMessageTextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// messageText {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MessageText != nil {
			_ptr_messageText := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MessageText != nil {
					if err := o.MessageText.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MessageText, _ptr_messageText); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMessageTextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// messageText {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_messageText := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MessageText == nil {
				o.MessageText = &oaut.String{}
			}
			if err := o.MessageText.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_messageText := func(ptr interface{}) { o.MessageText = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MessageText, _s_messageText, _ptr_messageText); err != nil {
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

// GetMessageTextRequest structure represents the MessageText operation request
type GetMessageTextRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMessageTextRequest) xxx_ToOp(ctx context.Context) *xxx_GetMessageTextOperation {
	if o == nil {
		return &xxx_GetMessageTextOperation{}
	}
	return &xxx_GetMessageTextOperation{
		This: o.This,
	}
}

func (o *GetMessageTextRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMessageTextOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMessageTextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetMessageTextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMessageTextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMessageTextResponse structure represents the MessageText operation response
type GetMessageTextResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MessageText *oaut.String   `idl:"name:messageText" json:"message_text"`
	// Return: The MessageText return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMessageTextResponse) xxx_ToOp(ctx context.Context) *xxx_GetMessageTextOperation {
	if o == nil {
		return &xxx_GetMessageTextOperation{}
	}
	return &xxx_GetMessageTextOperation{
		That:        o.That,
		MessageText: o.MessageText,
		Return:      o.Return,
	}
}

func (o *GetMessageTextResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMessageTextOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MessageText = op.MessageText
	o.Return = op.Return
}
func (o *GetMessageTextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetMessageTextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMessageTextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMessageTextOperation structure represents the MessageText operation
type xxx_SetMessageTextOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MessageText *oaut.String   `idl:"name:messageText" json:"message_text"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMessageTextOperation) OpNum() int { return 15 }

func (o *xxx_SetMessageTextOperation) OpName() string { return "/IFsrmActionEventLog/v0/MessageText" }

func (o *xxx_SetMessageTextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMessageTextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// messageText {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MessageText != nil {
			_ptr_messageText := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MessageText != nil {
					if err := o.MessageText.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MessageText, _ptr_messageText); err != nil {
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

func (o *xxx_SetMessageTextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// messageText {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_messageText := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MessageText == nil {
				o.MessageText = &oaut.String{}
			}
			if err := o.MessageText.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_messageText := func(ptr interface{}) { o.MessageText = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MessageText, _s_messageText, _ptr_messageText); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMessageTextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMessageTextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMessageTextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMessageTextRequest structure represents the MessageText operation request
type SetMessageTextRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	MessageText *oaut.String   `idl:"name:messageText" json:"message_text"`
}

func (o *SetMessageTextRequest) xxx_ToOp(ctx context.Context) *xxx_SetMessageTextOperation {
	if o == nil {
		return &xxx_SetMessageTextOperation{}
	}
	return &xxx_SetMessageTextOperation{
		This:        o.This,
		MessageText: o.MessageText,
	}
}

func (o *SetMessageTextRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMessageTextOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MessageText = op.MessageText
}
func (o *SetMessageTextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetMessageTextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMessageTextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMessageTextResponse structure represents the MessageText operation response
type SetMessageTextResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MessageText return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMessageTextResponse) xxx_ToOp(ctx context.Context) *xxx_SetMessageTextOperation {
	if o == nil {
		return &xxx_SetMessageTextOperation{}
	}
	return &xxx_SetMessageTextOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetMessageTextResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMessageTextOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMessageTextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetMessageTextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMessageTextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
