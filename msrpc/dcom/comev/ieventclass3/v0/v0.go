package ieventclass3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	ieventclass2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ieventclass2/v0"
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
	_ = dcom.GoPackage
	_ = ieventclass2.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/comev"
)

var (
	// IEventClass3 interface identifier 7fb7ea43-2d76-4ea8-8cd9-3decc270295e
	EventClass3IID = &dcom.IID{Data1: 0x7fb7ea43, Data2: 0x2d76, Data3: 0x4ea8, Data4: []byte{0x8c, 0xd9, 0x3d, 0xec, 0xc2, 0x70, 0x29, 0x5e}}
	// Syntax UUID
	EventClass3SyntaxUUID = &uuid.UUID{TimeLow: 0x7fb7ea43, TimeMid: 0x2d76, TimeHiAndVersion: 0x4ea8, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0xd9, Node: [6]uint8{0x3d, 0xec, 0xc2, 0x70, 0x29, 0x5e}}
	// Syntax ID
	EventClass3SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: EventClass3SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IEventClass3 interface.
type EventClass3Client interface {

	// IEventClass2 retrieval method.
	EventClass2() ieventclass2.EventClass2Client

	// EventClassPartitionID operation.
	GetEventClassPartitionID(context.Context, *GetEventClassPartitionIDRequest, ...dcerpc.CallOption) (*GetEventClassPartitionIDResponse, error)

	// EventClassPartitionID operation.
	SetEventClassPartitionID(context.Context, *SetEventClassPartitionIDRequest, ...dcerpc.CallOption) (*SetEventClassPartitionIDResponse, error)

	// EventClassApplicationID operation.
	GetEventClassApplicationID(context.Context, *GetEventClassApplicationIDRequest, ...dcerpc.CallOption) (*GetEventClassApplicationIDResponse, error)

	// EventClassApplicationID operation.
	SetEventClassApplicationID(context.Context, *SetEventClassApplicationIDRequest, ...dcerpc.CallOption) (*SetEventClassApplicationIDResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) EventClass3Client
}

type xxx_DefaultEventClass3Client struct {
	ieventclass2.EventClass2Client
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultEventClass3Client) EventClass2() ieventclass2.EventClass2Client {
	return o.EventClass2Client
}

func (o *xxx_DefaultEventClass3Client) GetEventClassPartitionID(ctx context.Context, in *GetEventClassPartitionIDRequest, opts ...dcerpc.CallOption) (*GetEventClassPartitionIDResponse, error) {
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
	out := &GetEventClassPartitionIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClass3Client) SetEventClassPartitionID(ctx context.Context, in *SetEventClassPartitionIDRequest, opts ...dcerpc.CallOption) (*SetEventClassPartitionIDResponse, error) {
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
	out := &SetEventClassPartitionIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClass3Client) GetEventClassApplicationID(ctx context.Context, in *GetEventClassApplicationIDRequest, opts ...dcerpc.CallOption) (*GetEventClassApplicationIDResponse, error) {
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
	out := &GetEventClassApplicationIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClass3Client) SetEventClassApplicationID(ctx context.Context, in *SetEventClassApplicationIDRequest, opts ...dcerpc.CallOption) (*SetEventClassApplicationIDResponse, error) {
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
	out := &SetEventClassApplicationIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClass3Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultEventClass3Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultEventClass3Client) IPID(ctx context.Context, ipid *dcom.IPID) EventClass3Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultEventClass3Client{
		EventClass2Client: o.EventClass2Client.IPID(ctx, ipid),
		cc:                o.cc,
		ipid:              ipid,
	}
}

func NewEventClass3Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (EventClass3Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(EventClass3SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ieventclass2.NewEventClass2Client(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultEventClass3Client{
		EventClass2Client: base,
		cc:                cc,
		ipid:              ipid,
	}, nil
}

// xxx_GetEventClassPartitionIDOperation structure represents the EventClassPartitionID operation
type xxx_GetEventClassPartitionIDOperation struct {
	This                  *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                  *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventClassPartitionID *oaut.String   `idl:"name:pbstrEventClassPartitionID" json:"event_class_partition_id"`
	Return                int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEventClassPartitionIDOperation) OpNum() int { return 29 }

func (o *xxx_GetEventClassPartitionIDOperation) OpName() string {
	return "/IEventClass3/v0/EventClassPartitionID"
}

func (o *xxx_GetEventClassPartitionIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventClassPartitionIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetEventClassPartitionIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetEventClassPartitionIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventClassPartitionIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrEventClassPartitionID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EventClassPartitionID != nil {
			_ptr_pbstrEventClassPartitionID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EventClassPartitionID != nil {
					if err := o.EventClassPartitionID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EventClassPartitionID, _ptr_pbstrEventClassPartitionID); err != nil {
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

func (o *xxx_GetEventClassPartitionIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrEventClassPartitionID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrEventClassPartitionID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EventClassPartitionID == nil {
				o.EventClassPartitionID = &oaut.String{}
			}
			if err := o.EventClassPartitionID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrEventClassPartitionID := func(ptr interface{}) { o.EventClassPartitionID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.EventClassPartitionID, _s_pbstrEventClassPartitionID, _ptr_pbstrEventClassPartitionID); err != nil {
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

// GetEventClassPartitionIDRequest structure represents the EventClassPartitionID operation request
type GetEventClassPartitionIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetEventClassPartitionIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetEventClassPartitionIDOperation) *xxx_GetEventClassPartitionIDOperation {
	if op == nil {
		op = &xxx_GetEventClassPartitionIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetEventClassPartitionIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEventClassPartitionIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetEventClassPartitionIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetEventClassPartitionIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventClassPartitionIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEventClassPartitionIDResponse structure represents the EventClassPartitionID operation response
type GetEventClassPartitionIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                  *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventClassPartitionID *oaut.String   `idl:"name:pbstrEventClassPartitionID" json:"event_class_partition_id"`
	// Return: The EventClassPartitionID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetEventClassPartitionIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetEventClassPartitionIDOperation) *xxx_GetEventClassPartitionIDOperation {
	if op == nil {
		op = &xxx_GetEventClassPartitionIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.EventClassPartitionID = o.EventClassPartitionID
	op.Return = o.Return
	return op
}

func (o *GetEventClassPartitionIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEventClassPartitionIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.EventClassPartitionID = op.EventClassPartitionID
	o.Return = op.Return
}
func (o *GetEventClassPartitionIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetEventClassPartitionIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventClassPartitionIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetEventClassPartitionIDOperation structure represents the EventClassPartitionID operation
type xxx_SetEventClassPartitionIDOperation struct {
	This                  *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                  *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventClassPartitionID *oaut.String   `idl:"name:bstrEventClassPartitionID" json:"event_class_partition_id"`
	Return                int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetEventClassPartitionIDOperation) OpNum() int { return 30 }

func (o *xxx_SetEventClassPartitionIDOperation) OpName() string {
	return "/IEventClass3/v0/EventClassPartitionID"
}

func (o *xxx_SetEventClassPartitionIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventClassPartitionIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrEventClassPartitionID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EventClassPartitionID != nil {
			_ptr_bstrEventClassPartitionID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EventClassPartitionID != nil {
					if err := o.EventClassPartitionID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EventClassPartitionID, _ptr_bstrEventClassPartitionID); err != nil {
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

func (o *xxx_SetEventClassPartitionIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrEventClassPartitionID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrEventClassPartitionID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EventClassPartitionID == nil {
				o.EventClassPartitionID = &oaut.String{}
			}
			if err := o.EventClassPartitionID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrEventClassPartitionID := func(ptr interface{}) { o.EventClassPartitionID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.EventClassPartitionID, _s_bstrEventClassPartitionID, _ptr_bstrEventClassPartitionID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventClassPartitionIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventClassPartitionIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetEventClassPartitionIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetEventClassPartitionIDRequest structure represents the EventClassPartitionID operation request
type SetEventClassPartitionIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                  *dcom.ORPCThis `idl:"name:This" json:"this"`
	EventClassPartitionID *oaut.String   `idl:"name:bstrEventClassPartitionID" json:"event_class_partition_id"`
}

func (o *SetEventClassPartitionIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetEventClassPartitionIDOperation) *xxx_SetEventClassPartitionIDOperation {
	if op == nil {
		op = &xxx_SetEventClassPartitionIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.EventClassPartitionID = o.EventClassPartitionID
	return op
}

func (o *SetEventClassPartitionIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetEventClassPartitionIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.EventClassPartitionID = op.EventClassPartitionID
}
func (o *SetEventClassPartitionIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetEventClassPartitionIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEventClassPartitionIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetEventClassPartitionIDResponse structure represents the EventClassPartitionID operation response
type SetEventClassPartitionIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The EventClassPartitionID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetEventClassPartitionIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetEventClassPartitionIDOperation) *xxx_SetEventClassPartitionIDOperation {
	if op == nil {
		op = &xxx_SetEventClassPartitionIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetEventClassPartitionIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetEventClassPartitionIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetEventClassPartitionIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetEventClassPartitionIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEventClassPartitionIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetEventClassApplicationIDOperation structure represents the EventClassApplicationID operation
type xxx_GetEventClassApplicationIDOperation struct {
	This                    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventClassApplicationID *oaut.String   `idl:"name:pbstrEventClassApplicationID" json:"event_class_application_id"`
	Return                  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEventClassApplicationIDOperation) OpNum() int { return 31 }

func (o *xxx_GetEventClassApplicationIDOperation) OpName() string {
	return "/IEventClass3/v0/EventClassApplicationID"
}

func (o *xxx_GetEventClassApplicationIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventClassApplicationIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetEventClassApplicationIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetEventClassApplicationIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventClassApplicationIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrEventClassApplicationID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EventClassApplicationID != nil {
			_ptr_pbstrEventClassApplicationID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EventClassApplicationID != nil {
					if err := o.EventClassApplicationID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EventClassApplicationID, _ptr_pbstrEventClassApplicationID); err != nil {
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

func (o *xxx_GetEventClassApplicationIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrEventClassApplicationID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrEventClassApplicationID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EventClassApplicationID == nil {
				o.EventClassApplicationID = &oaut.String{}
			}
			if err := o.EventClassApplicationID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrEventClassApplicationID := func(ptr interface{}) { o.EventClassApplicationID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.EventClassApplicationID, _s_pbstrEventClassApplicationID, _ptr_pbstrEventClassApplicationID); err != nil {
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

// GetEventClassApplicationIDRequest structure represents the EventClassApplicationID operation request
type GetEventClassApplicationIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetEventClassApplicationIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetEventClassApplicationIDOperation) *xxx_GetEventClassApplicationIDOperation {
	if op == nil {
		op = &xxx_GetEventClassApplicationIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetEventClassApplicationIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEventClassApplicationIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetEventClassApplicationIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetEventClassApplicationIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventClassApplicationIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEventClassApplicationIDResponse structure represents the EventClassApplicationID operation response
type GetEventClassApplicationIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                    *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventClassApplicationID *oaut.String   `idl:"name:pbstrEventClassApplicationID" json:"event_class_application_id"`
	// Return: The EventClassApplicationID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetEventClassApplicationIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetEventClassApplicationIDOperation) *xxx_GetEventClassApplicationIDOperation {
	if op == nil {
		op = &xxx_GetEventClassApplicationIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.EventClassApplicationID = o.EventClassApplicationID
	op.Return = o.Return
	return op
}

func (o *GetEventClassApplicationIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEventClassApplicationIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.EventClassApplicationID = op.EventClassApplicationID
	o.Return = op.Return
}
func (o *GetEventClassApplicationIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetEventClassApplicationIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventClassApplicationIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetEventClassApplicationIDOperation structure represents the EventClassApplicationID operation
type xxx_SetEventClassApplicationIDOperation struct {
	This                    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventClassApplicationID *oaut.String   `idl:"name:bstrEventClassApplicationID" json:"event_class_application_id"`
	Return                  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetEventClassApplicationIDOperation) OpNum() int { return 32 }

func (o *xxx_SetEventClassApplicationIDOperation) OpName() string {
	return "/IEventClass3/v0/EventClassApplicationID"
}

func (o *xxx_SetEventClassApplicationIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventClassApplicationIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrEventClassApplicationID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EventClassApplicationID != nil {
			_ptr_bstrEventClassApplicationID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EventClassApplicationID != nil {
					if err := o.EventClassApplicationID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EventClassApplicationID, _ptr_bstrEventClassApplicationID); err != nil {
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

func (o *xxx_SetEventClassApplicationIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrEventClassApplicationID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrEventClassApplicationID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EventClassApplicationID == nil {
				o.EventClassApplicationID = &oaut.String{}
			}
			if err := o.EventClassApplicationID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrEventClassApplicationID := func(ptr interface{}) { o.EventClassApplicationID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.EventClassApplicationID, _s_bstrEventClassApplicationID, _ptr_bstrEventClassApplicationID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventClassApplicationIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventClassApplicationIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetEventClassApplicationIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetEventClassApplicationIDRequest structure represents the EventClassApplicationID operation request
type SetEventClassApplicationIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                    *dcom.ORPCThis `idl:"name:This" json:"this"`
	EventClassApplicationID *oaut.String   `idl:"name:bstrEventClassApplicationID" json:"event_class_application_id"`
}

func (o *SetEventClassApplicationIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetEventClassApplicationIDOperation) *xxx_SetEventClassApplicationIDOperation {
	if op == nil {
		op = &xxx_SetEventClassApplicationIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.EventClassApplicationID = o.EventClassApplicationID
	return op
}

func (o *SetEventClassApplicationIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetEventClassApplicationIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.EventClassApplicationID = op.EventClassApplicationID
}
func (o *SetEventClassApplicationIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetEventClassApplicationIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEventClassApplicationIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetEventClassApplicationIDResponse structure represents the EventClassApplicationID operation response
type SetEventClassApplicationIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The EventClassApplicationID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetEventClassApplicationIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetEventClassApplicationIDOperation) *xxx_SetEventClassApplicationIDOperation {
	if op == nil {
		op = &xxx_SetEventClassApplicationIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetEventClassApplicationIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetEventClassApplicationIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetEventClassApplicationIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetEventClassApplicationIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEventClassApplicationIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
