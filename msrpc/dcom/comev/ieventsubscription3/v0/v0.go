package ieventsubscription3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	ieventsubscription2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ieventsubscription2/v0"
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
	_ = ieventsubscription2.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/comev"
)

var (
	// IEventSubscription3 interface identifier fbc1d17d-c498-43a0-81af-423ddd530af6
	EventSubscription3IID = &dcom.IID{Data1: 0xfbc1d17d, Data2: 0xc498, Data3: 0x43a0, Data4: []byte{0x81, 0xaf, 0x42, 0x3d, 0xdd, 0x53, 0x0a, 0xf6}}
	// Syntax UUID
	EventSubscription3SyntaxUUID = &uuid.UUID{TimeLow: 0xfbc1d17d, TimeMid: 0xc498, TimeHiAndVersion: 0x43a0, ClockSeqHiAndReserved: 0x81, ClockSeqLow: 0xaf, Node: [6]uint8{0x42, 0x3d, 0xdd, 0x53, 0xa, 0xf6}}
	// Syntax ID
	EventSubscription3SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: EventSubscription3SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IEventSubscription3 interface.
type EventSubscription3Client interface {

	// IEventSubscription2 retrieval method.
	EventSubscription2() ieventsubscription2.EventSubscription2Client

	// EventClassPartitionID operation.
	GetEventClassPartitionID(context.Context, *GetEventClassPartitionIDRequest, ...dcerpc.CallOption) (*GetEventClassPartitionIDResponse, error)

	// EventClassPartitionID operation.
	SetEventClassPartitionID(context.Context, *SetEventClassPartitionIDRequest, ...dcerpc.CallOption) (*SetEventClassPartitionIDResponse, error)

	// EventClassApplicationID operation.
	GetEventClassApplicationID(context.Context, *GetEventClassApplicationIDRequest, ...dcerpc.CallOption) (*GetEventClassApplicationIDResponse, error)

	// EventClassApplicationID operation.
	SetEventClassApplicationID(context.Context, *SetEventClassApplicationIDRequest, ...dcerpc.CallOption) (*SetEventClassApplicationIDResponse, error)

	// SubscriberPartitionID operation.
	GetSubscriberPartitionID(context.Context, *GetSubscriberPartitionIDRequest, ...dcerpc.CallOption) (*GetSubscriberPartitionIDResponse, error)

	// SubscriberPartitionID operation.
	SetSubscriberPartitionID(context.Context, *SetSubscriberPartitionIDRequest, ...dcerpc.CallOption) (*SetSubscriberPartitionIDResponse, error)

	// SubscriberApplicationID operation.
	GetSubscriberApplicationID(context.Context, *GetSubscriberApplicationIDRequest, ...dcerpc.CallOption) (*GetSubscriberApplicationIDResponse, error)

	// SubscriberApplicationID operation.
	SetSubscriberApplicationID(context.Context, *SetSubscriberApplicationIDRequest, ...dcerpc.CallOption) (*SetSubscriberApplicationIDResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) EventSubscription3Client
}

type xxx_DefaultEventSubscription3Client struct {
	ieventsubscription2.EventSubscription2Client
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultEventSubscription3Client) EventSubscription2() ieventsubscription2.EventSubscription2Client {
	return o.EventSubscription2Client
}

func (o *xxx_DefaultEventSubscription3Client) GetEventClassPartitionID(ctx context.Context, in *GetEventClassPartitionIDRequest, opts ...dcerpc.CallOption) (*GetEventClassPartitionIDResponse, error) {
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
	out := &GetEventClassPartitionIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscription3Client) SetEventClassPartitionID(ctx context.Context, in *SetEventClassPartitionIDRequest, opts ...dcerpc.CallOption) (*SetEventClassPartitionIDResponse, error) {
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
	out := &SetEventClassPartitionIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscription3Client) GetEventClassApplicationID(ctx context.Context, in *GetEventClassApplicationIDRequest, opts ...dcerpc.CallOption) (*GetEventClassApplicationIDResponse, error) {
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
	out := &GetEventClassApplicationIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscription3Client) SetEventClassApplicationID(ctx context.Context, in *SetEventClassApplicationIDRequest, opts ...dcerpc.CallOption) (*SetEventClassApplicationIDResponse, error) {
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
	out := &SetEventClassApplicationIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscription3Client) GetSubscriberPartitionID(ctx context.Context, in *GetSubscriberPartitionIDRequest, opts ...dcerpc.CallOption) (*GetSubscriberPartitionIDResponse, error) {
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
	out := &GetSubscriberPartitionIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscription3Client) SetSubscriberPartitionID(ctx context.Context, in *SetSubscriberPartitionIDRequest, opts ...dcerpc.CallOption) (*SetSubscriberPartitionIDResponse, error) {
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
	out := &SetSubscriberPartitionIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscription3Client) GetSubscriberApplicationID(ctx context.Context, in *GetSubscriberApplicationIDRequest, opts ...dcerpc.CallOption) (*GetSubscriberApplicationIDResponse, error) {
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
	out := &GetSubscriberApplicationIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscription3Client) SetSubscriberApplicationID(ctx context.Context, in *SetSubscriberApplicationIDRequest, opts ...dcerpc.CallOption) (*SetSubscriberApplicationIDResponse, error) {
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
	out := &SetSubscriberApplicationIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscription3Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultEventSubscription3Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultEventSubscription3Client) IPID(ctx context.Context, ipid *dcom.IPID) EventSubscription3Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultEventSubscription3Client{
		EventSubscription2Client: o.EventSubscription2Client.IPID(ctx, ipid),
		cc:                       o.cc,
		ipid:                     ipid,
	}
}

func NewEventSubscription3Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (EventSubscription3Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(EventSubscription3SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ieventsubscription2.NewEventSubscription2Client(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultEventSubscription3Client{
		EventSubscription2Client: base,
		cc:                       cc,
		ipid:                     ipid,
	}, nil
}

// xxx_GetEventClassPartitionIDOperation structure represents the EventClassPartitionID operation
type xxx_GetEventClassPartitionIDOperation struct {
	This                  *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                  *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventClassPartitionID *oaut.String   `idl:"name:pbstrEventClassPartitionID" json:"event_class_partition_id"`
	Return                int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEventClassPartitionIDOperation) OpNum() int { return 45 }

func (o *xxx_GetEventClassPartitionIDOperation) OpName() string {
	return "/IEventSubscription3/v0/EventClassPartitionID"
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

func (o *GetEventClassPartitionIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetEventClassPartitionIDOperation {
	if o == nil {
		return &xxx_GetEventClassPartitionIDOperation{}
	}
	return &xxx_GetEventClassPartitionIDOperation{
		This: o.This,
	}
}

func (o *GetEventClassPartitionIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEventClassPartitionIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetEventClassPartitionIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
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

func (o *GetEventClassPartitionIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetEventClassPartitionIDOperation {
	if o == nil {
		return &xxx_GetEventClassPartitionIDOperation{}
	}
	return &xxx_GetEventClassPartitionIDOperation{
		That:                  o.That,
		EventClassPartitionID: o.EventClassPartitionID,
		Return:                o.Return,
	}
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
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
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

func (o *xxx_SetEventClassPartitionIDOperation) OpNum() int { return 46 }

func (o *xxx_SetEventClassPartitionIDOperation) OpName() string {
	return "/IEventSubscription3/v0/EventClassPartitionID"
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

func (o *SetEventClassPartitionIDRequest) xxx_ToOp(ctx context.Context) *xxx_SetEventClassPartitionIDOperation {
	if o == nil {
		return &xxx_SetEventClassPartitionIDOperation{}
	}
	return &xxx_SetEventClassPartitionIDOperation{
		This:                  o.This,
		EventClassPartitionID: o.EventClassPartitionID,
	}
}

func (o *SetEventClassPartitionIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetEventClassPartitionIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.EventClassPartitionID = op.EventClassPartitionID
}
func (o *SetEventClassPartitionIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
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

func (o *SetEventClassPartitionIDResponse) xxx_ToOp(ctx context.Context) *xxx_SetEventClassPartitionIDOperation {
	if o == nil {
		return &xxx_SetEventClassPartitionIDOperation{}
	}
	return &xxx_SetEventClassPartitionIDOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetEventClassPartitionIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetEventClassPartitionIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetEventClassPartitionIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
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

func (o *xxx_GetEventClassApplicationIDOperation) OpNum() int { return 47 }

func (o *xxx_GetEventClassApplicationIDOperation) OpName() string {
	return "/IEventSubscription3/v0/EventClassApplicationID"
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

func (o *GetEventClassApplicationIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetEventClassApplicationIDOperation {
	if o == nil {
		return &xxx_GetEventClassApplicationIDOperation{}
	}
	return &xxx_GetEventClassApplicationIDOperation{
		This: o.This,
	}
}

func (o *GetEventClassApplicationIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEventClassApplicationIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetEventClassApplicationIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
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

func (o *GetEventClassApplicationIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetEventClassApplicationIDOperation {
	if o == nil {
		return &xxx_GetEventClassApplicationIDOperation{}
	}
	return &xxx_GetEventClassApplicationIDOperation{
		That:                    o.That,
		EventClassApplicationID: o.EventClassApplicationID,
		Return:                  o.Return,
	}
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
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
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

func (o *xxx_SetEventClassApplicationIDOperation) OpNum() int { return 48 }

func (o *xxx_SetEventClassApplicationIDOperation) OpName() string {
	return "/IEventSubscription3/v0/EventClassApplicationID"
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

func (o *SetEventClassApplicationIDRequest) xxx_ToOp(ctx context.Context) *xxx_SetEventClassApplicationIDOperation {
	if o == nil {
		return &xxx_SetEventClassApplicationIDOperation{}
	}
	return &xxx_SetEventClassApplicationIDOperation{
		This:                    o.This,
		EventClassApplicationID: o.EventClassApplicationID,
	}
}

func (o *SetEventClassApplicationIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetEventClassApplicationIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.EventClassApplicationID = op.EventClassApplicationID
}
func (o *SetEventClassApplicationIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
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

func (o *SetEventClassApplicationIDResponse) xxx_ToOp(ctx context.Context) *xxx_SetEventClassApplicationIDOperation {
	if o == nil {
		return &xxx_SetEventClassApplicationIDOperation{}
	}
	return &xxx_SetEventClassApplicationIDOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetEventClassApplicationIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetEventClassApplicationIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetEventClassApplicationIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetEventClassApplicationIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEventClassApplicationIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSubscriberPartitionIDOperation structure represents the SubscriberPartitionID operation
type xxx_GetSubscriberPartitionIDOperation struct {
	This                  *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                  *dcom.ORPCThat `idl:"name:That" json:"that"`
	SubscriberPartitionID *oaut.String   `idl:"name:pbstrSubscriberPartitionID" json:"subscriber_partition_id"`
	Return                int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSubscriberPartitionIDOperation) OpNum() int { return 49 }

func (o *xxx_GetSubscriberPartitionIDOperation) OpName() string {
	return "/IEventSubscription3/v0/SubscriberPartitionID"
}

func (o *xxx_GetSubscriberPartitionIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberPartitionIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSubscriberPartitionIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSubscriberPartitionIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberPartitionIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrSubscriberPartitionID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SubscriberPartitionID != nil {
			_ptr_pbstrSubscriberPartitionID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SubscriberPartitionID != nil {
					if err := o.SubscriberPartitionID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SubscriberPartitionID, _ptr_pbstrSubscriberPartitionID); err != nil {
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

func (o *xxx_GetSubscriberPartitionIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrSubscriberPartitionID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrSubscriberPartitionID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SubscriberPartitionID == nil {
				o.SubscriberPartitionID = &oaut.String{}
			}
			if err := o.SubscriberPartitionID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrSubscriberPartitionID := func(ptr interface{}) { o.SubscriberPartitionID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SubscriberPartitionID, _s_pbstrSubscriberPartitionID, _ptr_pbstrSubscriberPartitionID); err != nil {
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

// GetSubscriberPartitionIDRequest structure represents the SubscriberPartitionID operation request
type GetSubscriberPartitionIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSubscriberPartitionIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetSubscriberPartitionIDOperation {
	if o == nil {
		return &xxx_GetSubscriberPartitionIDOperation{}
	}
	return &xxx_GetSubscriberPartitionIDOperation{
		This: o.This,
	}
}

func (o *GetSubscriberPartitionIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSubscriberPartitionIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSubscriberPartitionIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSubscriberPartitionIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubscriberPartitionIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSubscriberPartitionIDResponse structure represents the SubscriberPartitionID operation response
type GetSubscriberPartitionIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                  *dcom.ORPCThat `idl:"name:That" json:"that"`
	SubscriberPartitionID *oaut.String   `idl:"name:pbstrSubscriberPartitionID" json:"subscriber_partition_id"`
	// Return: The SubscriberPartitionID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSubscriberPartitionIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetSubscriberPartitionIDOperation {
	if o == nil {
		return &xxx_GetSubscriberPartitionIDOperation{}
	}
	return &xxx_GetSubscriberPartitionIDOperation{
		That:                  o.That,
		SubscriberPartitionID: o.SubscriberPartitionID,
		Return:                o.Return,
	}
}

func (o *GetSubscriberPartitionIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSubscriberPartitionIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SubscriberPartitionID = op.SubscriberPartitionID
	o.Return = op.Return
}
func (o *GetSubscriberPartitionIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSubscriberPartitionIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubscriberPartitionIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSubscriberPartitionIDOperation structure represents the SubscriberPartitionID operation
type xxx_SetSubscriberPartitionIDOperation struct {
	This                  *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                  *dcom.ORPCThat `idl:"name:That" json:"that"`
	SubscriberPartitionID *oaut.String   `idl:"name:bstrSubscriberPartitionID" json:"subscriber_partition_id"`
	Return                int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSubscriberPartitionIDOperation) OpNum() int { return 50 }

func (o *xxx_SetSubscriberPartitionIDOperation) OpName() string {
	return "/IEventSubscription3/v0/SubscriberPartitionID"
}

func (o *xxx_SetSubscriberPartitionIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriberPartitionIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrSubscriberPartitionID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SubscriberPartitionID != nil {
			_ptr_bstrSubscriberPartitionID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SubscriberPartitionID != nil {
					if err := o.SubscriberPartitionID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SubscriberPartitionID, _ptr_bstrSubscriberPartitionID); err != nil {
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

func (o *xxx_SetSubscriberPartitionIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrSubscriberPartitionID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrSubscriberPartitionID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SubscriberPartitionID == nil {
				o.SubscriberPartitionID = &oaut.String{}
			}
			if err := o.SubscriberPartitionID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrSubscriberPartitionID := func(ptr interface{}) { o.SubscriberPartitionID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SubscriberPartitionID, _s_bstrSubscriberPartitionID, _ptr_bstrSubscriberPartitionID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriberPartitionIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriberPartitionIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSubscriberPartitionIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSubscriberPartitionIDRequest structure represents the SubscriberPartitionID operation request
type SetSubscriberPartitionIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                  *dcom.ORPCThis `idl:"name:This" json:"this"`
	SubscriberPartitionID *oaut.String   `idl:"name:bstrSubscriberPartitionID" json:"subscriber_partition_id"`
}

func (o *SetSubscriberPartitionIDRequest) xxx_ToOp(ctx context.Context) *xxx_SetSubscriberPartitionIDOperation {
	if o == nil {
		return &xxx_SetSubscriberPartitionIDOperation{}
	}
	return &xxx_SetSubscriberPartitionIDOperation{
		This:                  o.This,
		SubscriberPartitionID: o.SubscriberPartitionID,
	}
}

func (o *SetSubscriberPartitionIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSubscriberPartitionIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SubscriberPartitionID = op.SubscriberPartitionID
}
func (o *SetSubscriberPartitionIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSubscriberPartitionIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubscriberPartitionIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSubscriberPartitionIDResponse structure represents the SubscriberPartitionID operation response
type SetSubscriberPartitionIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SubscriberPartitionID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSubscriberPartitionIDResponse) xxx_ToOp(ctx context.Context) *xxx_SetSubscriberPartitionIDOperation {
	if o == nil {
		return &xxx_SetSubscriberPartitionIDOperation{}
	}
	return &xxx_SetSubscriberPartitionIDOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetSubscriberPartitionIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSubscriberPartitionIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSubscriberPartitionIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSubscriberPartitionIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubscriberPartitionIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSubscriberApplicationIDOperation structure represents the SubscriberApplicationID operation
type xxx_GetSubscriberApplicationIDOperation struct {
	This                    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat `idl:"name:That" json:"that"`
	SubscriberApplicationID *oaut.String   `idl:"name:pbstrSubscriberApplicationID" json:"subscriber_application_id"`
	Return                  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSubscriberApplicationIDOperation) OpNum() int { return 51 }

func (o *xxx_GetSubscriberApplicationIDOperation) OpName() string {
	return "/IEventSubscription3/v0/SubscriberApplicationID"
}

func (o *xxx_GetSubscriberApplicationIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberApplicationIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSubscriberApplicationIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSubscriberApplicationIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberApplicationIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrSubscriberApplicationID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SubscriberApplicationID != nil {
			_ptr_pbstrSubscriberApplicationID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SubscriberApplicationID != nil {
					if err := o.SubscriberApplicationID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SubscriberApplicationID, _ptr_pbstrSubscriberApplicationID); err != nil {
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

func (o *xxx_GetSubscriberApplicationIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrSubscriberApplicationID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrSubscriberApplicationID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SubscriberApplicationID == nil {
				o.SubscriberApplicationID = &oaut.String{}
			}
			if err := o.SubscriberApplicationID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrSubscriberApplicationID := func(ptr interface{}) { o.SubscriberApplicationID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SubscriberApplicationID, _s_pbstrSubscriberApplicationID, _ptr_pbstrSubscriberApplicationID); err != nil {
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

// GetSubscriberApplicationIDRequest structure represents the SubscriberApplicationID operation request
type GetSubscriberApplicationIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSubscriberApplicationIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetSubscriberApplicationIDOperation {
	if o == nil {
		return &xxx_GetSubscriberApplicationIDOperation{}
	}
	return &xxx_GetSubscriberApplicationIDOperation{
		This: o.This,
	}
}

func (o *GetSubscriberApplicationIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSubscriberApplicationIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSubscriberApplicationIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSubscriberApplicationIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubscriberApplicationIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSubscriberApplicationIDResponse structure represents the SubscriberApplicationID operation response
type GetSubscriberApplicationIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                    *dcom.ORPCThat `idl:"name:That" json:"that"`
	SubscriberApplicationID *oaut.String   `idl:"name:pbstrSubscriberApplicationID" json:"subscriber_application_id"`
	// Return: The SubscriberApplicationID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSubscriberApplicationIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetSubscriberApplicationIDOperation {
	if o == nil {
		return &xxx_GetSubscriberApplicationIDOperation{}
	}
	return &xxx_GetSubscriberApplicationIDOperation{
		That:                    o.That,
		SubscriberApplicationID: o.SubscriberApplicationID,
		Return:                  o.Return,
	}
}

func (o *GetSubscriberApplicationIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSubscriberApplicationIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SubscriberApplicationID = op.SubscriberApplicationID
	o.Return = op.Return
}
func (o *GetSubscriberApplicationIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSubscriberApplicationIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubscriberApplicationIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSubscriberApplicationIDOperation structure represents the SubscriberApplicationID operation
type xxx_SetSubscriberApplicationIDOperation struct {
	This                    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                    *dcom.ORPCThat `idl:"name:That" json:"that"`
	SubscriberApplicationID *oaut.String   `idl:"name:bstrSubscriberApplicationID" json:"subscriber_application_id"`
	Return                  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSubscriberApplicationIDOperation) OpNum() int { return 52 }

func (o *xxx_SetSubscriberApplicationIDOperation) OpName() string {
	return "/IEventSubscription3/v0/SubscriberApplicationID"
}

func (o *xxx_SetSubscriberApplicationIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriberApplicationIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrSubscriberApplicationID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SubscriberApplicationID != nil {
			_ptr_bstrSubscriberApplicationID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SubscriberApplicationID != nil {
					if err := o.SubscriberApplicationID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SubscriberApplicationID, _ptr_bstrSubscriberApplicationID); err != nil {
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

func (o *xxx_SetSubscriberApplicationIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrSubscriberApplicationID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrSubscriberApplicationID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SubscriberApplicationID == nil {
				o.SubscriberApplicationID = &oaut.String{}
			}
			if err := o.SubscriberApplicationID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrSubscriberApplicationID := func(ptr interface{}) { o.SubscriberApplicationID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SubscriberApplicationID, _s_bstrSubscriberApplicationID, _ptr_bstrSubscriberApplicationID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriberApplicationIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriberApplicationIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSubscriberApplicationIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSubscriberApplicationIDRequest structure represents the SubscriberApplicationID operation request
type SetSubscriberApplicationIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                    *dcom.ORPCThis `idl:"name:This" json:"this"`
	SubscriberApplicationID *oaut.String   `idl:"name:bstrSubscriberApplicationID" json:"subscriber_application_id"`
}

func (o *SetSubscriberApplicationIDRequest) xxx_ToOp(ctx context.Context) *xxx_SetSubscriberApplicationIDOperation {
	if o == nil {
		return &xxx_SetSubscriberApplicationIDOperation{}
	}
	return &xxx_SetSubscriberApplicationIDOperation{
		This:                    o.This,
		SubscriberApplicationID: o.SubscriberApplicationID,
	}
}

func (o *SetSubscriberApplicationIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSubscriberApplicationIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SubscriberApplicationID = op.SubscriberApplicationID
}
func (o *SetSubscriberApplicationIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSubscriberApplicationIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubscriberApplicationIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSubscriberApplicationIDResponse structure represents the SubscriberApplicationID operation response
type SetSubscriberApplicationIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SubscriberApplicationID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSubscriberApplicationIDResponse) xxx_ToOp(ctx context.Context) *xxx_SetSubscriberApplicationIDOperation {
	if o == nil {
		return &xxx_SetSubscriberApplicationIDOperation{}
	}
	return &xxx_SetSubscriberApplicationIDOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetSubscriberApplicationIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSubscriberApplicationIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSubscriberApplicationIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSubscriberApplicationIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubscriberApplicationIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
