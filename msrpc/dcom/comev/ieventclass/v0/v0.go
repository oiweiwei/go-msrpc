package ieventclass

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/comev"
)

var (
	// IEventClass interface identifier fb2b72a0-7a68-11d1-88f9-0080c7d771bf
	EventClassIID = &dcom.IID{Data1: 0xfb2b72a0, Data2: 0x7a68, Data3: 0x11d1, Data4: []byte{0x88, 0xf9, 0x00, 0x80, 0xc7, 0xd7, 0x71, 0xbf}}
	// Syntax UUID
	EventClassSyntaxUUID = &uuid.UUID{TimeLow: 0xfb2b72a0, TimeMid: 0x7a68, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x88, ClockSeqLow: 0xf9, Node: [6]uint8{0x0, 0x80, 0xc7, 0xd7, 0x71, 0xbf}}
	// Syntax ID
	EventClassSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: EventClassSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IEventClass interface.
type EventClassClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// EventClassID operation.
	GetEventClassID(context.Context, *GetEventClassIDRequest, ...dcerpc.CallOption) (*GetEventClassIDResponse, error)

	// EventClassID operation.
	SetEventClassID(context.Context, *SetEventClassIDRequest, ...dcerpc.CallOption) (*SetEventClassIDResponse, error)

	// EventClassName operation.
	GetEventClassName(context.Context, *GetEventClassNameRequest, ...dcerpc.CallOption) (*GetEventClassNameResponse, error)

	// EventClassName operation.
	SetEventClassName(context.Context, *SetEventClassNameRequest, ...dcerpc.CallOption) (*SetEventClassNameResponse, error)

	// OwnerSID operation.
	GetOwnerSID(context.Context, *GetOwnerSIDRequest, ...dcerpc.CallOption) (*GetOwnerSIDResponse, error)

	// OwnerSID operation.
	SetOwnerSID(context.Context, *SetOwnerSIDRequest, ...dcerpc.CallOption) (*SetOwnerSIDResponse, error)

	// FiringInterfaceID operation.
	GetFiringInterfaceID(context.Context, *GetFiringInterfaceIDRequest, ...dcerpc.CallOption) (*GetFiringInterfaceIDResponse, error)

	// FiringInterfaceID operation.
	SetFiringInterfaceID(context.Context, *SetFiringInterfaceIDRequest, ...dcerpc.CallOption) (*SetFiringInterfaceIDResponse, error)

	// Description operation.
	GetDescription(context.Context, *GetDescriptionRequest, ...dcerpc.CallOption) (*GetDescriptionResponse, error)

	// Description operation.
	SetDescription(context.Context, *SetDescriptionRequest, ...dcerpc.CallOption) (*SetDescriptionResponse, error)

	// Opnum17NotUsedOnWire operation.
	// Opnum17NotUsedOnWire

	// Opnum18NotUsedOnWire operation.
	// Opnum18NotUsedOnWire

	// TypeLib operation.
	GetTypeLib(context.Context, *GetTypeLibRequest, ...dcerpc.CallOption) (*GetTypeLibResponse, error)

	// TypeLib operation.
	SetTypeLib(context.Context, *SetTypeLibRequest, ...dcerpc.CallOption) (*SetTypeLibResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) EventClassClient
}

type xxx_DefaultEventClassClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultEventClassClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultEventClassClient) GetEventClassID(ctx context.Context, in *GetEventClassIDRequest, opts ...dcerpc.CallOption) (*GetEventClassIDResponse, error) {
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
	out := &GetEventClassIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClassClient) SetEventClassID(ctx context.Context, in *SetEventClassIDRequest, opts ...dcerpc.CallOption) (*SetEventClassIDResponse, error) {
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
	out := &SetEventClassIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClassClient) GetEventClassName(ctx context.Context, in *GetEventClassNameRequest, opts ...dcerpc.CallOption) (*GetEventClassNameResponse, error) {
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
	out := &GetEventClassNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClassClient) SetEventClassName(ctx context.Context, in *SetEventClassNameRequest, opts ...dcerpc.CallOption) (*SetEventClassNameResponse, error) {
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
	out := &SetEventClassNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClassClient) GetOwnerSID(ctx context.Context, in *GetOwnerSIDRequest, opts ...dcerpc.CallOption) (*GetOwnerSIDResponse, error) {
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
	out := &GetOwnerSIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClassClient) SetOwnerSID(ctx context.Context, in *SetOwnerSIDRequest, opts ...dcerpc.CallOption) (*SetOwnerSIDResponse, error) {
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
	out := &SetOwnerSIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClassClient) GetFiringInterfaceID(ctx context.Context, in *GetFiringInterfaceIDRequest, opts ...dcerpc.CallOption) (*GetFiringInterfaceIDResponse, error) {
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
	out := &GetFiringInterfaceIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClassClient) SetFiringInterfaceID(ctx context.Context, in *SetFiringInterfaceIDRequest, opts ...dcerpc.CallOption) (*SetFiringInterfaceIDResponse, error) {
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
	out := &SetFiringInterfaceIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClassClient) GetDescription(ctx context.Context, in *GetDescriptionRequest, opts ...dcerpc.CallOption) (*GetDescriptionResponse, error) {
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
	out := &GetDescriptionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClassClient) SetDescription(ctx context.Context, in *SetDescriptionRequest, opts ...dcerpc.CallOption) (*SetDescriptionResponse, error) {
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
	out := &SetDescriptionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClassClient) GetTypeLib(ctx context.Context, in *GetTypeLibRequest, opts ...dcerpc.CallOption) (*GetTypeLibResponse, error) {
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
	out := &GetTypeLibResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClassClient) SetTypeLib(ctx context.Context, in *SetTypeLibRequest, opts ...dcerpc.CallOption) (*SetTypeLibResponse, error) {
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
	out := &SetTypeLibResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClassClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultEventClassClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultEventClassClient) IPID(ctx context.Context, ipid *dcom.IPID) EventClassClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultEventClassClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewEventClassClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (EventClassClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(EventClassSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultEventClassClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetEventClassIDOperation structure represents the EventClassID operation
type xxx_GetEventClassIDOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventClassID *oaut.String   `idl:"name:pbstrEventClassID" json:"event_class_id"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEventClassIDOperation) OpNum() int { return 7 }

func (o *xxx_GetEventClassIDOperation) OpName() string { return "/IEventClass/v0/EventClassID" }

func (o *xxx_GetEventClassIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventClassIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetEventClassIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetEventClassIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventClassIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrEventClassID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EventClassID != nil {
			_ptr_pbstrEventClassID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EventClassID != nil {
					if err := o.EventClassID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EventClassID, _ptr_pbstrEventClassID); err != nil {
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

func (o *xxx_GetEventClassIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrEventClassID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrEventClassID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EventClassID == nil {
				o.EventClassID = &oaut.String{}
			}
			if err := o.EventClassID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrEventClassID := func(ptr interface{}) { o.EventClassID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.EventClassID, _s_pbstrEventClassID, _ptr_pbstrEventClassID); err != nil {
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

// GetEventClassIDRequest structure represents the EventClassID operation request
type GetEventClassIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetEventClassIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetEventClassIDOperation {
	if o == nil {
		return &xxx_GetEventClassIDOperation{}
	}
	return &xxx_GetEventClassIDOperation{
		This: o.This,
	}
}

func (o *GetEventClassIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEventClassIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetEventClassIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetEventClassIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventClassIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEventClassIDResponse structure represents the EventClassID operation response
type GetEventClassIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventClassID *oaut.String   `idl:"name:pbstrEventClassID" json:"event_class_id"`
	// Return: The EventClassID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetEventClassIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetEventClassIDOperation {
	if o == nil {
		return &xxx_GetEventClassIDOperation{}
	}
	return &xxx_GetEventClassIDOperation{
		That:         o.That,
		EventClassID: o.EventClassID,
		Return:       o.Return,
	}
}

func (o *GetEventClassIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEventClassIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.EventClassID = op.EventClassID
	o.Return = op.Return
}
func (o *GetEventClassIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetEventClassIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventClassIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetEventClassIDOperation structure represents the EventClassID operation
type xxx_SetEventClassIDOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventClassID *oaut.String   `idl:"name:bstrEventClassID" json:"event_class_id"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetEventClassIDOperation) OpNum() int { return 8 }

func (o *xxx_SetEventClassIDOperation) OpName() string { return "/IEventClass/v0/EventClassID" }

func (o *xxx_SetEventClassIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventClassIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrEventClassID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EventClassID != nil {
			_ptr_bstrEventClassID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EventClassID != nil {
					if err := o.EventClassID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EventClassID, _ptr_bstrEventClassID); err != nil {
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

func (o *xxx_SetEventClassIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrEventClassID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrEventClassID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EventClassID == nil {
				o.EventClassID = &oaut.String{}
			}
			if err := o.EventClassID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrEventClassID := func(ptr interface{}) { o.EventClassID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.EventClassID, _s_bstrEventClassID, _ptr_bstrEventClassID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventClassIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventClassIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetEventClassIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetEventClassIDRequest structure represents the EventClassID operation request
type SetEventClassIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	EventClassID *oaut.String   `idl:"name:bstrEventClassID" json:"event_class_id"`
}

func (o *SetEventClassIDRequest) xxx_ToOp(ctx context.Context) *xxx_SetEventClassIDOperation {
	if o == nil {
		return &xxx_SetEventClassIDOperation{}
	}
	return &xxx_SetEventClassIDOperation{
		This:         o.This,
		EventClassID: o.EventClassID,
	}
}

func (o *SetEventClassIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetEventClassIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.EventClassID = op.EventClassID
}
func (o *SetEventClassIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetEventClassIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEventClassIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetEventClassIDResponse structure represents the EventClassID operation response
type SetEventClassIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The EventClassID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetEventClassIDResponse) xxx_ToOp(ctx context.Context) *xxx_SetEventClassIDOperation {
	if o == nil {
		return &xxx_SetEventClassIDOperation{}
	}
	return &xxx_SetEventClassIDOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetEventClassIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetEventClassIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetEventClassIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetEventClassIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEventClassIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetEventClassNameOperation structure represents the EventClassName operation
type xxx_GetEventClassNameOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventClassName *oaut.String   `idl:"name:pbstrEventClassName" json:"event_class_name"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetEventClassNameOperation) OpNum() int { return 9 }

func (o *xxx_GetEventClassNameOperation) OpName() string { return "/IEventClass/v0/EventClassName" }

func (o *xxx_GetEventClassNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventClassNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetEventClassNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetEventClassNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetEventClassNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrEventClassName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EventClassName != nil {
			_ptr_pbstrEventClassName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EventClassName != nil {
					if err := o.EventClassName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EventClassName, _ptr_pbstrEventClassName); err != nil {
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

func (o *xxx_GetEventClassNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrEventClassName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrEventClassName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EventClassName == nil {
				o.EventClassName = &oaut.String{}
			}
			if err := o.EventClassName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrEventClassName := func(ptr interface{}) { o.EventClassName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.EventClassName, _s_pbstrEventClassName, _ptr_pbstrEventClassName); err != nil {
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

// GetEventClassNameRequest structure represents the EventClassName operation request
type GetEventClassNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetEventClassNameRequest) xxx_ToOp(ctx context.Context) *xxx_GetEventClassNameOperation {
	if o == nil {
		return &xxx_GetEventClassNameOperation{}
	}
	return &xxx_GetEventClassNameOperation{
		This: o.This,
	}
}

func (o *GetEventClassNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetEventClassNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetEventClassNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetEventClassNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventClassNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetEventClassNameResponse structure represents the EventClassName operation response
type GetEventClassNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventClassName *oaut.String   `idl:"name:pbstrEventClassName" json:"event_class_name"`
	// Return: The EventClassName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetEventClassNameResponse) xxx_ToOp(ctx context.Context) *xxx_GetEventClassNameOperation {
	if o == nil {
		return &xxx_GetEventClassNameOperation{}
	}
	return &xxx_GetEventClassNameOperation{
		That:           o.That,
		EventClassName: o.EventClassName,
		Return:         o.Return,
	}
}

func (o *GetEventClassNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetEventClassNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.EventClassName = op.EventClassName
	o.Return = op.Return
}
func (o *GetEventClassNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetEventClassNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetEventClassNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetEventClassNameOperation structure represents the EventClassName operation
type xxx_SetEventClassNameOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventClassName *oaut.String   `idl:"name:bstrEventClassName" json:"event_class_name"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetEventClassNameOperation) OpNum() int { return 10 }

func (o *xxx_SetEventClassNameOperation) OpName() string { return "/IEventClass/v0/EventClassName" }

func (o *xxx_SetEventClassNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventClassNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrEventClassName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EventClassName != nil {
			_ptr_bstrEventClassName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EventClassName != nil {
					if err := o.EventClassName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EventClassName, _ptr_bstrEventClassName); err != nil {
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

func (o *xxx_SetEventClassNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrEventClassName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrEventClassName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EventClassName == nil {
				o.EventClassName = &oaut.String{}
			}
			if err := o.EventClassName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrEventClassName := func(ptr interface{}) { o.EventClassName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.EventClassName, _s_bstrEventClassName, _ptr_bstrEventClassName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventClassNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventClassNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetEventClassNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetEventClassNameRequest structure represents the EventClassName operation request
type SetEventClassNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	EventClassName *oaut.String   `idl:"name:bstrEventClassName" json:"event_class_name"`
}

func (o *SetEventClassNameRequest) xxx_ToOp(ctx context.Context) *xxx_SetEventClassNameOperation {
	if o == nil {
		return &xxx_SetEventClassNameOperation{}
	}
	return &xxx_SetEventClassNameOperation{
		This:           o.This,
		EventClassName: o.EventClassName,
	}
}

func (o *SetEventClassNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetEventClassNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.EventClassName = op.EventClassName
}
func (o *SetEventClassNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetEventClassNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEventClassNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetEventClassNameResponse structure represents the EventClassName operation response
type SetEventClassNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The EventClassName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetEventClassNameResponse) xxx_ToOp(ctx context.Context) *xxx_SetEventClassNameOperation {
	if o == nil {
		return &xxx_SetEventClassNameOperation{}
	}
	return &xxx_SetEventClassNameOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetEventClassNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetEventClassNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetEventClassNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetEventClassNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEventClassNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetOwnerSIDOperation structure represents the OwnerSID operation
type xxx_GetOwnerSIDOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	OwnerSID *oaut.String   `idl:"name:pbstrOwnerSID" json:"owner_sid"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetOwnerSIDOperation) OpNum() int { return 11 }

func (o *xxx_GetOwnerSIDOperation) OpName() string { return "/IEventClass/v0/OwnerSID" }

func (o *xxx_GetOwnerSIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOwnerSIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetOwnerSIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetOwnerSIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOwnerSIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrOwnerSID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.OwnerSID != nil {
			_ptr_pbstrOwnerSID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OwnerSID != nil {
					if err := o.OwnerSID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OwnerSID, _ptr_pbstrOwnerSID); err != nil {
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

func (o *xxx_GetOwnerSIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrOwnerSID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrOwnerSID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OwnerSID == nil {
				o.OwnerSID = &oaut.String{}
			}
			if err := o.OwnerSID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrOwnerSID := func(ptr interface{}) { o.OwnerSID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.OwnerSID, _s_pbstrOwnerSID, _ptr_pbstrOwnerSID); err != nil {
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

// GetOwnerSIDRequest structure represents the OwnerSID operation request
type GetOwnerSIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetOwnerSIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetOwnerSIDOperation {
	if o == nil {
		return &xxx_GetOwnerSIDOperation{}
	}
	return &xxx_GetOwnerSIDOperation{
		This: o.This,
	}
}

func (o *GetOwnerSIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetOwnerSIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetOwnerSIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetOwnerSIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOwnerSIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetOwnerSIDResponse structure represents the OwnerSID operation response
type GetOwnerSIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	OwnerSID *oaut.String   `idl:"name:pbstrOwnerSID" json:"owner_sid"`
	// Return: The OwnerSID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetOwnerSIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetOwnerSIDOperation {
	if o == nil {
		return &xxx_GetOwnerSIDOperation{}
	}
	return &xxx_GetOwnerSIDOperation{
		That:     o.That,
		OwnerSID: o.OwnerSID,
		Return:   o.Return,
	}
}

func (o *GetOwnerSIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetOwnerSIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.OwnerSID = op.OwnerSID
	o.Return = op.Return
}
func (o *GetOwnerSIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetOwnerSIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOwnerSIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetOwnerSIDOperation structure represents the OwnerSID operation
type xxx_SetOwnerSIDOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	OwnerSID *oaut.String   `idl:"name:bstrOwnerSID" json:"owner_sid"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetOwnerSIDOperation) OpNum() int { return 12 }

func (o *xxx_SetOwnerSIDOperation) OpName() string { return "/IEventClass/v0/OwnerSID" }

func (o *xxx_SetOwnerSIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOwnerSIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrOwnerSID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.OwnerSID != nil {
			_ptr_bstrOwnerSID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OwnerSID != nil {
					if err := o.OwnerSID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OwnerSID, _ptr_bstrOwnerSID); err != nil {
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

func (o *xxx_SetOwnerSIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrOwnerSID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrOwnerSID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OwnerSID == nil {
				o.OwnerSID = &oaut.String{}
			}
			if err := o.OwnerSID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrOwnerSID := func(ptr interface{}) { o.OwnerSID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.OwnerSID, _s_bstrOwnerSID, _ptr_bstrOwnerSID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOwnerSIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOwnerSIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetOwnerSIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetOwnerSIDRequest structure represents the OwnerSID operation request
type SetOwnerSIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	OwnerSID *oaut.String   `idl:"name:bstrOwnerSID" json:"owner_sid"`
}

func (o *SetOwnerSIDRequest) xxx_ToOp(ctx context.Context) *xxx_SetOwnerSIDOperation {
	if o == nil {
		return &xxx_SetOwnerSIDOperation{}
	}
	return &xxx_SetOwnerSIDOperation{
		This:     o.This,
		OwnerSID: o.OwnerSID,
	}
}

func (o *SetOwnerSIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetOwnerSIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.OwnerSID = op.OwnerSID
}
func (o *SetOwnerSIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetOwnerSIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOwnerSIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetOwnerSIDResponse structure represents the OwnerSID operation response
type SetOwnerSIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The OwnerSID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetOwnerSIDResponse) xxx_ToOp(ctx context.Context) *xxx_SetOwnerSIDOperation {
	if o == nil {
		return &xxx_SetOwnerSIDOperation{}
	}
	return &xxx_SetOwnerSIDOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetOwnerSIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetOwnerSIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetOwnerSIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetOwnerSIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOwnerSIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFiringInterfaceIDOperation structure represents the FiringInterfaceID operation
type xxx_GetFiringInterfaceIDOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	FiringInterfaceID *oaut.String   `idl:"name:pbstrFiringInterfaceID" json:"firing_interface_id"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFiringInterfaceIDOperation) OpNum() int { return 13 }

func (o *xxx_GetFiringInterfaceIDOperation) OpName() string {
	return "/IEventClass/v0/FiringInterfaceID"
}

func (o *xxx_GetFiringInterfaceIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFiringInterfaceIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetFiringInterfaceIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetFiringInterfaceIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFiringInterfaceIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrFiringInterfaceID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FiringInterfaceID != nil {
			_ptr_pbstrFiringInterfaceID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FiringInterfaceID != nil {
					if err := o.FiringInterfaceID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FiringInterfaceID, _ptr_pbstrFiringInterfaceID); err != nil {
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

func (o *xxx_GetFiringInterfaceIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrFiringInterfaceID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrFiringInterfaceID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FiringInterfaceID == nil {
				o.FiringInterfaceID = &oaut.String{}
			}
			if err := o.FiringInterfaceID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrFiringInterfaceID := func(ptr interface{}) { o.FiringInterfaceID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FiringInterfaceID, _s_pbstrFiringInterfaceID, _ptr_pbstrFiringInterfaceID); err != nil {
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

// GetFiringInterfaceIDRequest structure represents the FiringInterfaceID operation request
type GetFiringInterfaceIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetFiringInterfaceIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetFiringInterfaceIDOperation {
	if o == nil {
		return &xxx_GetFiringInterfaceIDOperation{}
	}
	return &xxx_GetFiringInterfaceIDOperation{
		This: o.This,
	}
}

func (o *GetFiringInterfaceIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFiringInterfaceIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetFiringInterfaceIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetFiringInterfaceIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFiringInterfaceIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFiringInterfaceIDResponse structure represents the FiringInterfaceID operation response
type GetFiringInterfaceIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	FiringInterfaceID *oaut.String   `idl:"name:pbstrFiringInterfaceID" json:"firing_interface_id"`
	// Return: The FiringInterfaceID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFiringInterfaceIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetFiringInterfaceIDOperation {
	if o == nil {
		return &xxx_GetFiringInterfaceIDOperation{}
	}
	return &xxx_GetFiringInterfaceIDOperation{
		That:              o.That,
		FiringInterfaceID: o.FiringInterfaceID,
		Return:            o.Return,
	}
}

func (o *GetFiringInterfaceIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFiringInterfaceIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FiringInterfaceID = op.FiringInterfaceID
	o.Return = op.Return
}
func (o *GetFiringInterfaceIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetFiringInterfaceIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFiringInterfaceIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetFiringInterfaceIDOperation structure represents the FiringInterfaceID operation
type xxx_SetFiringInterfaceIDOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	FiringInterfaceID *oaut.String   `idl:"name:bstrFiringInterfaceID" json:"firing_interface_id"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetFiringInterfaceIDOperation) OpNum() int { return 14 }

func (o *xxx_SetFiringInterfaceIDOperation) OpName() string {
	return "/IEventClass/v0/FiringInterfaceID"
}

func (o *xxx_SetFiringInterfaceIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFiringInterfaceIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrFiringInterfaceID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.FiringInterfaceID != nil {
			_ptr_bstrFiringInterfaceID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FiringInterfaceID != nil {
					if err := o.FiringInterfaceID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FiringInterfaceID, _ptr_bstrFiringInterfaceID); err != nil {
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

func (o *xxx_SetFiringInterfaceIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrFiringInterfaceID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrFiringInterfaceID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.FiringInterfaceID == nil {
				o.FiringInterfaceID = &oaut.String{}
			}
			if err := o.FiringInterfaceID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrFiringInterfaceID := func(ptr interface{}) { o.FiringInterfaceID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.FiringInterfaceID, _s_bstrFiringInterfaceID, _ptr_bstrFiringInterfaceID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFiringInterfaceIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFiringInterfaceIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetFiringInterfaceIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetFiringInterfaceIDRequest structure represents the FiringInterfaceID operation request
type SetFiringInterfaceIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	FiringInterfaceID *oaut.String   `idl:"name:bstrFiringInterfaceID" json:"firing_interface_id"`
}

func (o *SetFiringInterfaceIDRequest) xxx_ToOp(ctx context.Context) *xxx_SetFiringInterfaceIDOperation {
	if o == nil {
		return &xxx_SetFiringInterfaceIDOperation{}
	}
	return &xxx_SetFiringInterfaceIDOperation{
		This:              o.This,
		FiringInterfaceID: o.FiringInterfaceID,
	}
}

func (o *SetFiringInterfaceIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetFiringInterfaceIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FiringInterfaceID = op.FiringInterfaceID
}
func (o *SetFiringInterfaceIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetFiringInterfaceIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFiringInterfaceIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetFiringInterfaceIDResponse structure represents the FiringInterfaceID operation response
type SetFiringInterfaceIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The FiringInterfaceID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetFiringInterfaceIDResponse) xxx_ToOp(ctx context.Context) *xxx_SetFiringInterfaceIDOperation {
	if o == nil {
		return &xxx_SetFiringInterfaceIDOperation{}
	}
	return &xxx_SetFiringInterfaceIDOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetFiringInterfaceIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetFiringInterfaceIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetFiringInterfaceIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetFiringInterfaceIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFiringInterfaceIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDescriptionOperation structure represents the Description operation
type xxx_GetDescriptionOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Description *oaut.String   `idl:"name:pbstrDescription" json:"description"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDescriptionOperation) OpNum() int { return 15 }

func (o *xxx_GetDescriptionOperation) OpName() string { return "/IEventClass/v0/Description" }

func (o *xxx_GetDescriptionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDescriptionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDescriptionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDescriptionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDescriptionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrDescription {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Description != nil {
			_ptr_pbstrDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Description != nil {
					if err := o.Description.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Description, _ptr_pbstrDescription); err != nil {
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

func (o *xxx_GetDescriptionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrDescription {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Description == nil {
				o.Description = &oaut.String{}
			}
			if err := o.Description.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrDescription := func(ptr interface{}) { o.Description = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Description, _s_pbstrDescription, _ptr_pbstrDescription); err != nil {
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

// GetDescriptionRequest structure represents the Description operation request
type GetDescriptionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDescriptionRequest) xxx_ToOp(ctx context.Context) *xxx_GetDescriptionOperation {
	if o == nil {
		return &xxx_GetDescriptionOperation{}
	}
	return &xxx_GetDescriptionOperation{
		This: o.This,
	}
}

func (o *GetDescriptionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDescriptionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDescriptionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDescriptionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDescriptionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDescriptionResponse structure represents the Description operation response
type GetDescriptionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Description *oaut.String   `idl:"name:pbstrDescription" json:"description"`
	// Return: The Description return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDescriptionResponse) xxx_ToOp(ctx context.Context) *xxx_GetDescriptionOperation {
	if o == nil {
		return &xxx_GetDescriptionOperation{}
	}
	return &xxx_GetDescriptionOperation{
		That:        o.That,
		Description: o.Description,
		Return:      o.Return,
	}
}

func (o *GetDescriptionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDescriptionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Description = op.Description
	o.Return = op.Return
}
func (o *GetDescriptionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDescriptionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDescriptionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetDescriptionOperation structure represents the Description operation
type xxx_SetDescriptionOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Description *oaut.String   `idl:"name:bstrDescription" json:"description"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDescriptionOperation) OpNum() int { return 16 }

func (o *xxx_SetDescriptionOperation) OpName() string { return "/IEventClass/v0/Description" }

func (o *xxx_SetDescriptionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDescriptionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrDescription {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Description != nil {
			_ptr_bstrDescription := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Description != nil {
					if err := o.Description.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Description, _ptr_bstrDescription); err != nil {
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

func (o *xxx_SetDescriptionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrDescription {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrDescription := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Description == nil {
				o.Description = &oaut.String{}
			}
			if err := o.Description.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrDescription := func(ptr interface{}) { o.Description = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Description, _s_bstrDescription, _ptr_bstrDescription); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDescriptionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDescriptionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetDescriptionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetDescriptionRequest structure represents the Description operation request
type SetDescriptionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	Description *oaut.String   `idl:"name:bstrDescription" json:"description"`
}

func (o *SetDescriptionRequest) xxx_ToOp(ctx context.Context) *xxx_SetDescriptionOperation {
	if o == nil {
		return &xxx_SetDescriptionOperation{}
	}
	return &xxx_SetDescriptionOperation{
		This:        o.This,
		Description: o.Description,
	}
}

func (o *SetDescriptionRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDescriptionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Description = op.Description
}
func (o *SetDescriptionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetDescriptionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDescriptionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDescriptionResponse structure represents the Description operation response
type SetDescriptionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Description return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetDescriptionResponse) xxx_ToOp(ctx context.Context) *xxx_SetDescriptionOperation {
	if o == nil {
		return &xxx_SetDescriptionOperation{}
	}
	return &xxx_SetDescriptionOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetDescriptionResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDescriptionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetDescriptionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetDescriptionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDescriptionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTypeLibOperation structure represents the TypeLib operation
type xxx_GetTypeLibOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	TypeLib *oaut.String   `idl:"name:pbstrTypeLib" json:"type_lib"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTypeLibOperation) OpNum() int { return 19 }

func (o *xxx_GetTypeLibOperation) OpName() string { return "/IEventClass/v0/TypeLib" }

func (o *xxx_GetTypeLibOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeLibOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetTypeLibOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetTypeLibOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTypeLibOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrTypeLib {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.TypeLib != nil {
			_ptr_pbstrTypeLib := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.TypeLib != nil {
					if err := o.TypeLib.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TypeLib, _ptr_pbstrTypeLib); err != nil {
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

func (o *xxx_GetTypeLibOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrTypeLib {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrTypeLib := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.TypeLib == nil {
				o.TypeLib = &oaut.String{}
			}
			if err := o.TypeLib.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrTypeLib := func(ptr interface{}) { o.TypeLib = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.TypeLib, _s_pbstrTypeLib, _ptr_pbstrTypeLib); err != nil {
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

// GetTypeLibRequest structure represents the TypeLib operation request
type GetTypeLibRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetTypeLibRequest) xxx_ToOp(ctx context.Context) *xxx_GetTypeLibOperation {
	if o == nil {
		return &xxx_GetTypeLibOperation{}
	}
	return &xxx_GetTypeLibOperation{
		This: o.This,
	}
}

func (o *GetTypeLibRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTypeLibOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTypeLibRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetTypeLibRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeLibOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTypeLibResponse structure represents the TypeLib operation response
type GetTypeLibResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	TypeLib *oaut.String   `idl:"name:pbstrTypeLib" json:"type_lib"`
	// Return: The TypeLib return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTypeLibResponse) xxx_ToOp(ctx context.Context) *xxx_GetTypeLibOperation {
	if o == nil {
		return &xxx_GetTypeLibOperation{}
	}
	return &xxx_GetTypeLibOperation{
		That:    o.That,
		TypeLib: o.TypeLib,
		Return:  o.Return,
	}
}

func (o *GetTypeLibResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTypeLibOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.TypeLib = op.TypeLib
	o.Return = op.Return
}
func (o *GetTypeLibResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetTypeLibResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTypeLibOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetTypeLibOperation structure represents the TypeLib operation
type xxx_SetTypeLibOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	TypeLib *oaut.String   `idl:"name:bstrTypeLib" json:"type_lib"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetTypeLibOperation) OpNum() int { return 20 }

func (o *xxx_SetTypeLibOperation) OpName() string { return "/IEventClass/v0/TypeLib" }

func (o *xxx_SetTypeLibOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTypeLibOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrTypeLib {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.TypeLib != nil {
			_ptr_bstrTypeLib := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.TypeLib != nil {
					if err := o.TypeLib.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.TypeLib, _ptr_bstrTypeLib); err != nil {
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

func (o *xxx_SetTypeLibOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrTypeLib {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrTypeLib := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.TypeLib == nil {
				o.TypeLib = &oaut.String{}
			}
			if err := o.TypeLib.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrTypeLib := func(ptr interface{}) { o.TypeLib = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.TypeLib, _s_bstrTypeLib, _ptr_bstrTypeLib); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTypeLibOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTypeLibOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetTypeLibOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetTypeLibRequest structure represents the TypeLib operation request
type SetTypeLibRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	TypeLib *oaut.String   `idl:"name:bstrTypeLib" json:"type_lib"`
}

func (o *SetTypeLibRequest) xxx_ToOp(ctx context.Context) *xxx_SetTypeLibOperation {
	if o == nil {
		return &xxx_SetTypeLibOperation{}
	}
	return &xxx_SetTypeLibOperation{
		This:    o.This,
		TypeLib: o.TypeLib,
	}
}

func (o *SetTypeLibRequest) xxx_FromOp(ctx context.Context, op *xxx_SetTypeLibOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.TypeLib = op.TypeLib
}
func (o *SetTypeLibRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetTypeLibRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTypeLibOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetTypeLibResponse structure represents the TypeLib operation response
type SetTypeLibResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The TypeLib return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetTypeLibResponse) xxx_ToOp(ctx context.Context) *xxx_SetTypeLibOperation {
	if o == nil {
		return &xxx_SetTypeLibOperation{}
	}
	return &xxx_SetTypeLibOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetTypeLibResponse) xxx_FromOp(ctx context.Context, op *xxx_SetTypeLibOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetTypeLibResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetTypeLibResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTypeLibOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
