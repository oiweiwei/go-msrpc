package ieventsubscription2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	ieventsubscription "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ieventsubscription/v0"
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
	_ = ieventsubscription.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/comev"
)

var (
	// IEventSubscription2 interface identifier 4a6b0e16-2e38-11d1-9965-00c04fbbb345
	EventSubscription2IID = &dcom.IID{Data1: 0x4a6b0e16, Data2: 0x2e38, Data3: 0x11d1, Data4: []byte{0x99, 0x65, 0x00, 0xc0, 0x4f, 0xbb, 0xb3, 0x45}}
	// Syntax UUID
	EventSubscription2SyntaxUUID = &uuid.UUID{TimeLow: 0x4a6b0e16, TimeMid: 0x2e38, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0x65, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xbb, 0xb3, 0x45}}
	// Syntax ID
	EventSubscription2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: EventSubscription2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IEventSubscription2 interface.
type EventSubscription2Client interface {

	// IEventSubscription retrieval method.
	EventSubscription() ieventsubscription.EventSubscriptionClient

	// FilterCriteria operation.
	GetFilterCriteria(context.Context, *GetFilterCriteriaRequest, ...dcerpc.CallOption) (*GetFilterCriteriaResponse, error)

	// FilterCriteria operation.
	SetFilterCriteria(context.Context, *SetFilterCriteriaRequest, ...dcerpc.CallOption) (*SetFilterCriteriaResponse, error)

	// SubscriberMoniker operation.
	GetSubscriberMoniker(context.Context, *GetSubscriberMonikerRequest, ...dcerpc.CallOption) (*GetSubscriberMonikerResponse, error)

	// SubscriberMoniker operation.
	SetSubscriberMoniker(context.Context, *SetSubscriberMonikerRequest, ...dcerpc.CallOption) (*SetSubscriberMonikerResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) EventSubscription2Client
}

type xxx_DefaultEventSubscription2Client struct {
	ieventsubscription.EventSubscriptionClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultEventSubscription2Client) EventSubscription() ieventsubscription.EventSubscriptionClient {
	return o.EventSubscriptionClient
}

func (o *xxx_DefaultEventSubscription2Client) GetFilterCriteria(ctx context.Context, in *GetFilterCriteriaRequest, opts ...dcerpc.CallOption) (*GetFilterCriteriaResponse, error) {
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
	out := &GetFilterCriteriaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscription2Client) SetFilterCriteria(ctx context.Context, in *SetFilterCriteriaRequest, opts ...dcerpc.CallOption) (*SetFilterCriteriaResponse, error) {
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
	out := &SetFilterCriteriaResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscription2Client) GetSubscriberMoniker(ctx context.Context, in *GetSubscriberMonikerRequest, opts ...dcerpc.CallOption) (*GetSubscriberMonikerResponse, error) {
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
	out := &GetSubscriberMonikerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscription2Client) SetSubscriberMoniker(ctx context.Context, in *SetSubscriberMonikerRequest, opts ...dcerpc.CallOption) (*SetSubscriberMonikerResponse, error) {
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
	out := &SetSubscriberMonikerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSubscription2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultEventSubscription2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultEventSubscription2Client) IPID(ctx context.Context, ipid *dcom.IPID) EventSubscription2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultEventSubscription2Client{
		EventSubscriptionClient: o.EventSubscriptionClient.IPID(ctx, ipid),
		cc:                      o.cc,
		ipid:                    ipid,
	}
}

func NewEventSubscription2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (EventSubscription2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(EventSubscription2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ieventsubscription.NewEventSubscriptionClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultEventSubscription2Client{
		EventSubscriptionClient: base,
		cc:                      cc,
		ipid:                    ipid,
	}, nil
}

// xxx_GetFilterCriteriaOperation structure represents the FilterCriteria operation
type xxx_GetFilterCriteriaOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	FilterCriteria *oaut.String   `idl:"name:pbstrFilterCriteria" json:"filter_criteria"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFilterCriteriaOperation) OpNum() int { return 41 }

func (o *xxx_GetFilterCriteriaOperation) OpName() string {
	return "/IEventSubscription2/v0/FilterCriteria"
}

func (o *xxx_GetFilterCriteriaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFilterCriteriaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetFilterCriteriaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetFilterCriteriaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFilterCriteriaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrFilterCriteria {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FilterCriteria != nil {
			_ptr_pbstrFilterCriteria := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FilterCriteria != nil {
					if err := o.FilterCriteria.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FilterCriteria, _ptr_pbstrFilterCriteria); err != nil {
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

func (o *xxx_GetFilterCriteriaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrFilterCriteria {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrFilterCriteria := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FilterCriteria == nil {
				o.FilterCriteria = &oaut.String{}
			}
			if err := o.FilterCriteria.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrFilterCriteria := func(ptr interface{}) { o.FilterCriteria = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FilterCriteria, _s_pbstrFilterCriteria, _ptr_pbstrFilterCriteria); err != nil {
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

// GetFilterCriteriaRequest structure represents the FilterCriteria operation request
type GetFilterCriteriaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetFilterCriteriaRequest) xxx_ToOp(ctx context.Context) *xxx_GetFilterCriteriaOperation {
	if o == nil {
		return &xxx_GetFilterCriteriaOperation{}
	}
	return &xxx_GetFilterCriteriaOperation{
		This: o.This,
	}
}

func (o *GetFilterCriteriaRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFilterCriteriaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetFilterCriteriaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetFilterCriteriaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFilterCriteriaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFilterCriteriaResponse structure represents the FilterCriteria operation response
type GetFilterCriteriaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	FilterCriteria *oaut.String   `idl:"name:pbstrFilterCriteria" json:"filter_criteria"`
	// Return: The FilterCriteria return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFilterCriteriaResponse) xxx_ToOp(ctx context.Context) *xxx_GetFilterCriteriaOperation {
	if o == nil {
		return &xxx_GetFilterCriteriaOperation{}
	}
	return &xxx_GetFilterCriteriaOperation{
		That:           o.That,
		FilterCriteria: o.FilterCriteria,
		Return:         o.Return,
	}
}

func (o *GetFilterCriteriaResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFilterCriteriaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FilterCriteria = op.FilterCriteria
	o.Return = op.Return
}
func (o *GetFilterCriteriaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetFilterCriteriaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFilterCriteriaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetFilterCriteriaOperation structure represents the FilterCriteria operation
type xxx_SetFilterCriteriaOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	FilterCriteria *oaut.String   `idl:"name:bstrFilterCriteria" json:"filter_criteria"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetFilterCriteriaOperation) OpNum() int { return 42 }

func (o *xxx_SetFilterCriteriaOperation) OpName() string {
	return "/IEventSubscription2/v0/FilterCriteria"
}

func (o *xxx_SetFilterCriteriaOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFilterCriteriaOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrFilterCriteria {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FilterCriteria != nil {
			_ptr_bstrFilterCriteria := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FilterCriteria != nil {
					if err := o.FilterCriteria.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FilterCriteria, _ptr_bstrFilterCriteria); err != nil {
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

func (o *xxx_SetFilterCriteriaOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrFilterCriteria {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrFilterCriteria := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FilterCriteria == nil {
				o.FilterCriteria = &oaut.String{}
			}
			if err := o.FilterCriteria.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrFilterCriteria := func(ptr interface{}) { o.FilterCriteria = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FilterCriteria, _s_bstrFilterCriteria, _ptr_bstrFilterCriteria); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFilterCriteriaOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFilterCriteriaOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetFilterCriteriaOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetFilterCriteriaRequest structure represents the FilterCriteria operation request
type SetFilterCriteriaRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	FilterCriteria *oaut.String   `idl:"name:bstrFilterCriteria" json:"filter_criteria"`
}

func (o *SetFilterCriteriaRequest) xxx_ToOp(ctx context.Context) *xxx_SetFilterCriteriaOperation {
	if o == nil {
		return &xxx_SetFilterCriteriaOperation{}
	}
	return &xxx_SetFilterCriteriaOperation{
		This:           o.This,
		FilterCriteria: o.FilterCriteria,
	}
}

func (o *SetFilterCriteriaRequest) xxx_FromOp(ctx context.Context, op *xxx_SetFilterCriteriaOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FilterCriteria = op.FilterCriteria
}
func (o *SetFilterCriteriaRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetFilterCriteriaRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFilterCriteriaOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetFilterCriteriaResponse structure represents the FilterCriteria operation response
type SetFilterCriteriaResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The FilterCriteria return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetFilterCriteriaResponse) xxx_ToOp(ctx context.Context) *xxx_SetFilterCriteriaOperation {
	if o == nil {
		return &xxx_SetFilterCriteriaOperation{}
	}
	return &xxx_SetFilterCriteriaOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetFilterCriteriaResponse) xxx_FromOp(ctx context.Context, op *xxx_SetFilterCriteriaOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetFilterCriteriaResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetFilterCriteriaResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFilterCriteriaOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSubscriberMonikerOperation structure represents the SubscriberMoniker operation
type xxx_GetSubscriberMonikerOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Moniker *oaut.String   `idl:"name:pbstrMoniker" json:"moniker"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSubscriberMonikerOperation) OpNum() int { return 43 }

func (o *xxx_GetSubscriberMonikerOperation) OpName() string {
	return "/IEventSubscription2/v0/SubscriberMoniker"
}

func (o *xxx_GetSubscriberMonikerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberMonikerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSubscriberMonikerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSubscriberMonikerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSubscriberMonikerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrMoniker {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Moniker != nil {
			_ptr_pbstrMoniker := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Moniker != nil {
					if err := o.Moniker.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Moniker, _ptr_pbstrMoniker); err != nil {
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

func (o *xxx_GetSubscriberMonikerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrMoniker {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrMoniker := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Moniker == nil {
				o.Moniker = &oaut.String{}
			}
			if err := o.Moniker.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrMoniker := func(ptr interface{}) { o.Moniker = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Moniker, _s_pbstrMoniker, _ptr_pbstrMoniker); err != nil {
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

// GetSubscriberMonikerRequest structure represents the SubscriberMoniker operation request
type GetSubscriberMonikerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSubscriberMonikerRequest) xxx_ToOp(ctx context.Context) *xxx_GetSubscriberMonikerOperation {
	if o == nil {
		return &xxx_GetSubscriberMonikerOperation{}
	}
	return &xxx_GetSubscriberMonikerOperation{
		This: o.This,
	}
}

func (o *GetSubscriberMonikerRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSubscriberMonikerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSubscriberMonikerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSubscriberMonikerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubscriberMonikerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSubscriberMonikerResponse structure represents the SubscriberMoniker operation response
type GetSubscriberMonikerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Moniker *oaut.String   `idl:"name:pbstrMoniker" json:"moniker"`
	// Return: The SubscriberMoniker return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSubscriberMonikerResponse) xxx_ToOp(ctx context.Context) *xxx_GetSubscriberMonikerOperation {
	if o == nil {
		return &xxx_GetSubscriberMonikerOperation{}
	}
	return &xxx_GetSubscriberMonikerOperation{
		That:    o.That,
		Moniker: o.Moniker,
		Return:  o.Return,
	}
}

func (o *GetSubscriberMonikerResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSubscriberMonikerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Moniker = op.Moniker
	o.Return = op.Return
}
func (o *GetSubscriberMonikerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSubscriberMonikerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSubscriberMonikerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSubscriberMonikerOperation structure represents the SubscriberMoniker operation
type xxx_SetSubscriberMonikerOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Moniker *oaut.String   `idl:"name:bstrMoniker" json:"moniker"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSubscriberMonikerOperation) OpNum() int { return 44 }

func (o *xxx_SetSubscriberMonikerOperation) OpName() string {
	return "/IEventSubscription2/v0/SubscriberMoniker"
}

func (o *xxx_SetSubscriberMonikerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriberMonikerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrMoniker {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Moniker != nil {
			_ptr_bstrMoniker := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Moniker != nil {
					if err := o.Moniker.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Moniker, _ptr_bstrMoniker); err != nil {
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

func (o *xxx_SetSubscriberMonikerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrMoniker {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMoniker := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Moniker == nil {
				o.Moniker = &oaut.String{}
			}
			if err := o.Moniker.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMoniker := func(ptr interface{}) { o.Moniker = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Moniker, _s_bstrMoniker, _ptr_bstrMoniker); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriberMonikerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSubscriberMonikerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetSubscriberMonikerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetSubscriberMonikerRequest structure represents the SubscriberMoniker operation request
type SetSubscriberMonikerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	Moniker *oaut.String   `idl:"name:bstrMoniker" json:"moniker"`
}

func (o *SetSubscriberMonikerRequest) xxx_ToOp(ctx context.Context) *xxx_SetSubscriberMonikerOperation {
	if o == nil {
		return &xxx_SetSubscriberMonikerOperation{}
	}
	return &xxx_SetSubscriberMonikerOperation{
		This:    o.This,
		Moniker: o.Moniker,
	}
}

func (o *SetSubscriberMonikerRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSubscriberMonikerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Moniker = op.Moniker
}
func (o *SetSubscriberMonikerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetSubscriberMonikerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubscriberMonikerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSubscriberMonikerResponse structure represents the SubscriberMoniker operation response
type SetSubscriberMonikerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SubscriberMoniker return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetSubscriberMonikerResponse) xxx_ToOp(ctx context.Context) *xxx_SetSubscriberMonikerOperation {
	if o == nil {
		return &xxx_SetSubscriberMonikerOperation{}
	}
	return &xxx_SetSubscriberMonikerOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetSubscriberMonikerResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSubscriberMonikerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetSubscriberMonikerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetSubscriberMonikerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSubscriberMonikerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
