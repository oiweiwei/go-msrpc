package ieventclass2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	ieventclass "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ieventclass/v0"
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
	_ = ieventclass.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/comev"
)

var (
	// IEventClass2 interface identifier fb2b72a1-7a68-11d1-88f9-0080c7d771bf
	EventClass2IID = &dcom.IID{Data1: 0xfb2b72a1, Data2: 0x7a68, Data3: 0x11d1, Data4: []byte{0x88, 0xf9, 0x00, 0x80, 0xc7, 0xd7, 0x71, 0xbf}}
	// Syntax UUID
	EventClass2SyntaxUUID = &uuid.UUID{TimeLow: 0xfb2b72a1, TimeMid: 0x7a68, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x88, ClockSeqLow: 0xf9, Node: [6]uint8{0x0, 0x80, 0xc7, 0xd7, 0x71, 0xbf}}
	// Syntax ID
	EventClass2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: EventClass2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IEventClass2 interface.
type EventClass2Client interface {

	// IEventClass retrieval method.
	EventClass() ieventclass.EventClassClient

	// PublisherID operation.
	GetPublisherID(context.Context, *GetPublisherIDRequest, ...dcerpc.CallOption) (*GetPublisherIDResponse, error)

	// PublisherID operation.
	SetPublisherID(context.Context, *SetPublisherIDRequest, ...dcerpc.CallOption) (*SetPublisherIDResponse, error)

	// MultiInterfacePublisherFilterCLSID operation.
	GetMultiInterfacePublisherFilterClassID(context.Context, *GetMultiInterfacePublisherFilterClassIDRequest, ...dcerpc.CallOption) (*GetMultiInterfacePublisherFilterClassIDResponse, error)

	// MultiInterfacePublisherFilterCLSID operation.
	SetMultiInterfacePublisherFilterClassID(context.Context, *SetMultiInterfacePublisherFilterClassIDRequest, ...dcerpc.CallOption) (*SetMultiInterfacePublisherFilterClassIDResponse, error)

	// AllowInprocActivation operation.
	GetAllowInProcessActivation(context.Context, *GetAllowInProcessActivationRequest, ...dcerpc.CallOption) (*GetAllowInProcessActivationResponse, error)

	// AllowInprocActivation operation.
	SetAllowInProcessActivation(context.Context, *SetAllowInProcessActivationRequest, ...dcerpc.CallOption) (*SetAllowInProcessActivationResponse, error)

	// FireInParallel operation.
	GetFireInParallel(context.Context, *GetFireInParallelRequest, ...dcerpc.CallOption) (*GetFireInParallelResponse, error)

	// FireInParallel operation.
	SetFireInParallel(context.Context, *SetFireInParallelRequest, ...dcerpc.CallOption) (*SetFireInParallelResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) EventClass2Client
}

type xxx_DefaultEventClass2Client struct {
	ieventclass.EventClassClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultEventClass2Client) EventClass() ieventclass.EventClassClient {
	return o.EventClassClient
}

func (o *xxx_DefaultEventClass2Client) GetPublisherID(ctx context.Context, in *GetPublisherIDRequest, opts ...dcerpc.CallOption) (*GetPublisherIDResponse, error) {
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
	out := &GetPublisherIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClass2Client) SetPublisherID(ctx context.Context, in *SetPublisherIDRequest, opts ...dcerpc.CallOption) (*SetPublisherIDResponse, error) {
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
	out := &SetPublisherIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClass2Client) GetMultiInterfacePublisherFilterClassID(ctx context.Context, in *GetMultiInterfacePublisherFilterClassIDRequest, opts ...dcerpc.CallOption) (*GetMultiInterfacePublisherFilterClassIDResponse, error) {
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
	out := &GetMultiInterfacePublisherFilterClassIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClass2Client) SetMultiInterfacePublisherFilterClassID(ctx context.Context, in *SetMultiInterfacePublisherFilterClassIDRequest, opts ...dcerpc.CallOption) (*SetMultiInterfacePublisherFilterClassIDResponse, error) {
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
	out := &SetMultiInterfacePublisherFilterClassIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClass2Client) GetAllowInProcessActivation(ctx context.Context, in *GetAllowInProcessActivationRequest, opts ...dcerpc.CallOption) (*GetAllowInProcessActivationResponse, error) {
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
	out := &GetAllowInProcessActivationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClass2Client) SetAllowInProcessActivation(ctx context.Context, in *SetAllowInProcessActivationRequest, opts ...dcerpc.CallOption) (*SetAllowInProcessActivationResponse, error) {
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
	out := &SetAllowInProcessActivationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClass2Client) GetFireInParallel(ctx context.Context, in *GetFireInParallelRequest, opts ...dcerpc.CallOption) (*GetFireInParallelResponse, error) {
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
	out := &GetFireInParallelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClass2Client) SetFireInParallel(ctx context.Context, in *SetFireInParallelRequest, opts ...dcerpc.CallOption) (*SetFireInParallelResponse, error) {
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
	out := &SetFireInParallelResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventClass2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultEventClass2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultEventClass2Client) IPID(ctx context.Context, ipid *dcom.IPID) EventClass2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultEventClass2Client{
		EventClassClient: o.EventClassClient.IPID(ctx, ipid),
		cc:               o.cc,
		ipid:             ipid,
	}
}

func NewEventClass2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (EventClass2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(EventClass2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ieventclass.NewEventClassClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultEventClass2Client{
		EventClassClient: base,
		cc:               cc,
		ipid:             ipid,
	}, nil
}

// xxx_GetPublisherIDOperation structure represents the PublisherID operation
type xxx_GetPublisherIDOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	PublisherID *oaut.String   `idl:"name:pbstrPublisherID" json:"publisher_id"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPublisherIDOperation) OpNum() int { return 21 }

func (o *xxx_GetPublisherIDOperation) OpName() string { return "/IEventClass2/v0/PublisherID" }

func (o *xxx_GetPublisherIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPublisherIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPublisherIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPublisherIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrPublisherID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PublisherID != nil {
			_ptr_pbstrPublisherID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PublisherID != nil {
					if err := o.PublisherID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PublisherID, _ptr_pbstrPublisherID); err != nil {
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

func (o *xxx_GetPublisherIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrPublisherID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrPublisherID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PublisherID == nil {
				o.PublisherID = &oaut.String{}
			}
			if err := o.PublisherID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrPublisherID := func(ptr interface{}) { o.PublisherID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PublisherID, _s_pbstrPublisherID, _ptr_pbstrPublisherID); err != nil {
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

// GetPublisherIDRequest structure represents the PublisherID operation request
type GetPublisherIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPublisherIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPublisherIDOperation) *xxx_GetPublisherIDOperation {
	if op == nil {
		op = &xxx_GetPublisherIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPublisherIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPublisherIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPublisherIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPublisherIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPublisherIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPublisherIDResponse structure represents the PublisherID operation response
type GetPublisherIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	PublisherID *oaut.String   `idl:"name:pbstrPublisherID" json:"publisher_id"`
	// Return: The PublisherID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPublisherIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPublisherIDOperation) *xxx_GetPublisherIDOperation {
	if op == nil {
		op = &xxx_GetPublisherIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PublisherID = o.PublisherID
	op.Return = o.Return
	return op
}

func (o *GetPublisherIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPublisherIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PublisherID = op.PublisherID
	o.Return = op.Return
}
func (o *GetPublisherIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPublisherIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPublisherIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetPublisherIDOperation structure represents the PublisherID operation
type xxx_SetPublisherIDOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	PublisherID *oaut.String   `idl:"name:bstrPublisherID" json:"publisher_id"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetPublisherIDOperation) OpNum() int { return 22 }

func (o *xxx_SetPublisherIDOperation) OpName() string { return "/IEventClass2/v0/PublisherID" }

func (o *xxx_SetPublisherIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPublisherIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrPublisherID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PublisherID != nil {
			_ptr_bstrPublisherID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PublisherID != nil {
					if err := o.PublisherID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PublisherID, _ptr_bstrPublisherID); err != nil {
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

func (o *xxx_SetPublisherIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrPublisherID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPublisherID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PublisherID == nil {
				o.PublisherID = &oaut.String{}
			}
			if err := o.PublisherID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPublisherID := func(ptr interface{}) { o.PublisherID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PublisherID, _s_bstrPublisherID, _ptr_bstrPublisherID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPublisherIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPublisherIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetPublisherIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetPublisherIDRequest structure represents the PublisherID operation request
type SetPublisherIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	PublisherID *oaut.String   `idl:"name:bstrPublisherID" json:"publisher_id"`
}

func (o *SetPublisherIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetPublisherIDOperation) *xxx_SetPublisherIDOperation {
	if op == nil {
		op = &xxx_SetPublisherIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PublisherID = o.PublisherID
	return op
}

func (o *SetPublisherIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetPublisherIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PublisherID = op.PublisherID
}
func (o *SetPublisherIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetPublisherIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPublisherIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetPublisherIDResponse structure represents the PublisherID operation response
type SetPublisherIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The PublisherID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetPublisherIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetPublisherIDOperation) *xxx_SetPublisherIDOperation {
	if op == nil {
		op = &xxx_SetPublisherIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetPublisherIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetPublisherIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetPublisherIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetPublisherIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPublisherIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMultiInterfacePublisherFilterClassIDOperation structure represents the MultiInterfacePublisherFilterCLSID operation
type xxx_GetMultiInterfacePublisherFilterClassIDOperation struct {
	This                   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                   *dcom.ORPCThat `idl:"name:That" json:"that"`
	PublisherFilterClassID *oaut.String   `idl:"name:pbstrPubFilCLSID" json:"publisher_filter_class_id"`
	Return                 int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMultiInterfacePublisherFilterClassIDOperation) OpNum() int { return 23 }

func (o *xxx_GetMultiInterfacePublisherFilterClassIDOperation) OpName() string {
	return "/IEventClass2/v0/MultiInterfacePublisherFilterCLSID"
}

func (o *xxx_GetMultiInterfacePublisherFilterClassIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMultiInterfacePublisherFilterClassIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMultiInterfacePublisherFilterClassIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMultiInterfacePublisherFilterClassIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMultiInterfacePublisherFilterClassIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrPubFilCLSID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PublisherFilterClassID != nil {
			_ptr_pbstrPubFilCLSID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PublisherFilterClassID != nil {
					if err := o.PublisherFilterClassID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PublisherFilterClassID, _ptr_pbstrPubFilCLSID); err != nil {
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

func (o *xxx_GetMultiInterfacePublisherFilterClassIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrPubFilCLSID {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrPubFilCLSID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PublisherFilterClassID == nil {
				o.PublisherFilterClassID = &oaut.String{}
			}
			if err := o.PublisherFilterClassID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrPubFilCLSID := func(ptr interface{}) { o.PublisherFilterClassID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PublisherFilterClassID, _s_pbstrPubFilCLSID, _ptr_pbstrPubFilCLSID); err != nil {
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

// GetMultiInterfacePublisherFilterClassIDRequest structure represents the MultiInterfacePublisherFilterCLSID operation request
type GetMultiInterfacePublisherFilterClassIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMultiInterfacePublisherFilterClassIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMultiInterfacePublisherFilterClassIDOperation) *xxx_GetMultiInterfacePublisherFilterClassIDOperation {
	if op == nil {
		op = &xxx_GetMultiInterfacePublisherFilterClassIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMultiInterfacePublisherFilterClassIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMultiInterfacePublisherFilterClassIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMultiInterfacePublisherFilterClassIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMultiInterfacePublisherFilterClassIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMultiInterfacePublisherFilterClassIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMultiInterfacePublisherFilterClassIDResponse structure represents the MultiInterfacePublisherFilterCLSID operation response
type GetMultiInterfacePublisherFilterClassIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                   *dcom.ORPCThat `idl:"name:That" json:"that"`
	PublisherFilterClassID *oaut.String   `idl:"name:pbstrPubFilCLSID" json:"publisher_filter_class_id"`
	// Return: The MultiInterfacePublisherFilterCLSID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMultiInterfacePublisherFilterClassIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMultiInterfacePublisherFilterClassIDOperation) *xxx_GetMultiInterfacePublisherFilterClassIDOperation {
	if op == nil {
		op = &xxx_GetMultiInterfacePublisherFilterClassIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PublisherFilterClassID = o.PublisherFilterClassID
	op.Return = o.Return
	return op
}

func (o *GetMultiInterfacePublisherFilterClassIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMultiInterfacePublisherFilterClassIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PublisherFilterClassID = op.PublisherFilterClassID
	o.Return = op.Return
}
func (o *GetMultiInterfacePublisherFilterClassIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMultiInterfacePublisherFilterClassIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMultiInterfacePublisherFilterClassIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMultiInterfacePublisherFilterClassIDOperation structure represents the MultiInterfacePublisherFilterCLSID operation
type xxx_SetMultiInterfacePublisherFilterClassIDOperation struct {
	This                   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                   *dcom.ORPCThat `idl:"name:That" json:"that"`
	PublisherFilterClassID *oaut.String   `idl:"name:bstrPubFilCLSID" json:"publisher_filter_class_id"`
	Return                 int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMultiInterfacePublisherFilterClassIDOperation) OpNum() int { return 24 }

func (o *xxx_SetMultiInterfacePublisherFilterClassIDOperation) OpName() string {
	return "/IEventClass2/v0/MultiInterfacePublisherFilterCLSID"
}

func (o *xxx_SetMultiInterfacePublisherFilterClassIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMultiInterfacePublisherFilterClassIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrPubFilCLSID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.PublisherFilterClassID != nil {
			_ptr_bstrPubFilCLSID := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PublisherFilterClassID != nil {
					if err := o.PublisherFilterClassID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PublisherFilterClassID, _ptr_bstrPubFilCLSID); err != nil {
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

func (o *xxx_SetMultiInterfacePublisherFilterClassIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrPubFilCLSID {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrPubFilCLSID := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PublisherFilterClassID == nil {
				o.PublisherFilterClassID = &oaut.String{}
			}
			if err := o.PublisherFilterClassID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrPubFilCLSID := func(ptr interface{}) { o.PublisherFilterClassID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.PublisherFilterClassID, _s_bstrPubFilCLSID, _ptr_bstrPubFilCLSID); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMultiInterfacePublisherFilterClassIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMultiInterfacePublisherFilterClassIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMultiInterfacePublisherFilterClassIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMultiInterfacePublisherFilterClassIDRequest structure represents the MultiInterfacePublisherFilterCLSID operation request
type SetMultiInterfacePublisherFilterClassIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                   *dcom.ORPCThis `idl:"name:This" json:"this"`
	PublisherFilterClassID *oaut.String   `idl:"name:bstrPubFilCLSID" json:"publisher_filter_class_id"`
}

func (o *SetMultiInterfacePublisherFilterClassIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetMultiInterfacePublisherFilterClassIDOperation) *xxx_SetMultiInterfacePublisherFilterClassIDOperation {
	if op == nil {
		op = &xxx_SetMultiInterfacePublisherFilterClassIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.PublisherFilterClassID = o.PublisherFilterClassID
	return op
}

func (o *SetMultiInterfacePublisherFilterClassIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMultiInterfacePublisherFilterClassIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PublisherFilterClassID = op.PublisherFilterClassID
}
func (o *SetMultiInterfacePublisherFilterClassIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetMultiInterfacePublisherFilterClassIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMultiInterfacePublisherFilterClassIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMultiInterfacePublisherFilterClassIDResponse structure represents the MultiInterfacePublisherFilterCLSID operation response
type SetMultiInterfacePublisherFilterClassIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MultiInterfacePublisherFilterCLSID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMultiInterfacePublisherFilterClassIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetMultiInterfacePublisherFilterClassIDOperation) *xxx_SetMultiInterfacePublisherFilterClassIDOperation {
	if op == nil {
		op = &xxx_SetMultiInterfacePublisherFilterClassIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetMultiInterfacePublisherFilterClassIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMultiInterfacePublisherFilterClassIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMultiInterfacePublisherFilterClassIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetMultiInterfacePublisherFilterClassIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMultiInterfacePublisherFilterClassIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAllowInProcessActivationOperation structure represents the AllowInprocActivation operation
type xxx_GetAllowInProcessActivationOperation struct {
	This                     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                     *dcom.ORPCThat `idl:"name:That" json:"that"`
	AllowInProcessActivation bool           `idl:"name:pfAllowInprocActivation" json:"allow_in_process_activation"`
	Return                   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAllowInProcessActivationOperation) OpNum() int { return 25 }

func (o *xxx_GetAllowInProcessActivationOperation) OpName() string {
	return "/IEventClass2/v0/AllowInprocActivation"
}

func (o *xxx_GetAllowInProcessActivationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllowInProcessActivationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAllowInProcessActivationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAllowInProcessActivationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAllowInProcessActivationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pfAllowInprocActivation {out, retval} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		if !o.AllowInProcessActivation {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
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

func (o *xxx_GetAllowInProcessActivationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pfAllowInprocActivation {out, retval} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		var _bAllowInProcessActivation int32
		if err := w.ReadData(&_bAllowInProcessActivation); err != nil {
			return err
		}
		o.AllowInProcessActivation = _bAllowInProcessActivation != 0
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetAllowInProcessActivationRequest structure represents the AllowInprocActivation operation request
type GetAllowInProcessActivationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAllowInProcessActivationRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAllowInProcessActivationOperation) *xxx_GetAllowInProcessActivationOperation {
	if op == nil {
		op = &xxx_GetAllowInProcessActivationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetAllowInProcessActivationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAllowInProcessActivationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAllowInProcessActivationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAllowInProcessActivationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllowInProcessActivationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAllowInProcessActivationResponse structure represents the AllowInprocActivation operation response
type GetAllowInProcessActivationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                     *dcom.ORPCThat `idl:"name:That" json:"that"`
	AllowInProcessActivation bool           `idl:"name:pfAllowInprocActivation" json:"allow_in_process_activation"`
	// Return: The AllowInprocActivation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAllowInProcessActivationResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAllowInProcessActivationOperation) *xxx_GetAllowInProcessActivationOperation {
	if op == nil {
		op = &xxx_GetAllowInProcessActivationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.AllowInProcessActivation = o.AllowInProcessActivation
	op.Return = o.Return
	return op
}

func (o *GetAllowInProcessActivationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAllowInProcessActivationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.AllowInProcessActivation = op.AllowInProcessActivation
	o.Return = op.Return
}
func (o *GetAllowInProcessActivationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAllowInProcessActivationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAllowInProcessActivationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetAllowInProcessActivationOperation structure represents the AllowInprocActivation operation
type xxx_SetAllowInProcessActivationOperation struct {
	This                     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                     *dcom.ORPCThat `idl:"name:That" json:"that"`
	AllowInProcessActivation bool           `idl:"name:fAllowInprocActivation" json:"allow_in_process_activation"`
	Return                   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetAllowInProcessActivationOperation) OpNum() int { return 26 }

func (o *xxx_SetAllowInProcessActivationOperation) OpName() string {
	return "/IEventClass2/v0/AllowInprocActivation"
}

func (o *xxx_SetAllowInProcessActivationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAllowInProcessActivationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// fAllowInprocActivation {in} (1:{alias=BOOL}(int32))
	{
		if !o.AllowInProcessActivation {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SetAllowInProcessActivationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// fAllowInprocActivation {in} (1:{alias=BOOL}(int32))
	{
		var _bAllowInProcessActivation int32
		if err := w.ReadData(&_bAllowInProcessActivation); err != nil {
			return err
		}
		o.AllowInProcessActivation = _bAllowInProcessActivation != 0
	}
	return nil
}

func (o *xxx_SetAllowInProcessActivationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAllowInProcessActivationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetAllowInProcessActivationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetAllowInProcessActivationRequest structure represents the AllowInprocActivation operation request
type SetAllowInProcessActivationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                     *dcom.ORPCThis `idl:"name:This" json:"this"`
	AllowInProcessActivation bool           `idl:"name:fAllowInprocActivation" json:"allow_in_process_activation"`
}

func (o *SetAllowInProcessActivationRequest) xxx_ToOp(ctx context.Context, op *xxx_SetAllowInProcessActivationOperation) *xxx_SetAllowInProcessActivationOperation {
	if op == nil {
		op = &xxx_SetAllowInProcessActivationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.AllowInProcessActivation = o.AllowInProcessActivation
	return op
}

func (o *SetAllowInProcessActivationRequest) xxx_FromOp(ctx context.Context, op *xxx_SetAllowInProcessActivationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.AllowInProcessActivation = op.AllowInProcessActivation
}
func (o *SetAllowInProcessActivationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetAllowInProcessActivationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAllowInProcessActivationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetAllowInProcessActivationResponse structure represents the AllowInprocActivation operation response
type SetAllowInProcessActivationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AllowInprocActivation return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetAllowInProcessActivationResponse) xxx_ToOp(ctx context.Context, op *xxx_SetAllowInProcessActivationOperation) *xxx_SetAllowInProcessActivationOperation {
	if op == nil {
		op = &xxx_SetAllowInProcessActivationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetAllowInProcessActivationResponse) xxx_FromOp(ctx context.Context, op *xxx_SetAllowInProcessActivationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetAllowInProcessActivationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetAllowInProcessActivationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAllowInProcessActivationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetFireInParallelOperation structure represents the FireInParallel operation
type xxx_GetFireInParallelOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	FireInParallel bool           `idl:"name:pfFireInParallel" json:"fire_in_parallel"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFireInParallelOperation) OpNum() int { return 27 }

func (o *xxx_GetFireInParallelOperation) OpName() string { return "/IEventClass2/v0/FireInParallel" }

func (o *xxx_GetFireInParallelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFireInParallelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetFireInParallelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetFireInParallelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFireInParallelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pfFireInParallel {out, retval} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		if !o.FireInParallel {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
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

func (o *xxx_GetFireInParallelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pfFireInParallel {out, retval} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		var _bFireInParallel int32
		if err := w.ReadData(&_bFireInParallel); err != nil {
			return err
		}
		o.FireInParallel = _bFireInParallel != 0
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetFireInParallelRequest structure represents the FireInParallel operation request
type GetFireInParallelRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetFireInParallelRequest) xxx_ToOp(ctx context.Context, op *xxx_GetFireInParallelOperation) *xxx_GetFireInParallelOperation {
	if op == nil {
		op = &xxx_GetFireInParallelOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetFireInParallelRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFireInParallelOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetFireInParallelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetFireInParallelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFireInParallelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFireInParallelResponse structure represents the FireInParallel operation response
type GetFireInParallelResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	FireInParallel bool           `idl:"name:pfFireInParallel" json:"fire_in_parallel"`
	// Return: The FireInParallel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFireInParallelResponse) xxx_ToOp(ctx context.Context, op *xxx_GetFireInParallelOperation) *xxx_GetFireInParallelOperation {
	if op == nil {
		op = &xxx_GetFireInParallelOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.FireInParallel = o.FireInParallel
	op.Return = o.Return
	return op
}

func (o *GetFireInParallelResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFireInParallelOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FireInParallel = op.FireInParallel
	o.Return = op.Return
}
func (o *GetFireInParallelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetFireInParallelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFireInParallelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetFireInParallelOperation structure represents the FireInParallel operation
type xxx_SetFireInParallelOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	FireInParallel bool           `idl:"name:fFireInParallel" json:"fire_in_parallel"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetFireInParallelOperation) OpNum() int { return 28 }

func (o *xxx_SetFireInParallelOperation) OpName() string { return "/IEventClass2/v0/FireInParallel" }

func (o *xxx_SetFireInParallelOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFireInParallelOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// fFireInParallel {in} (1:{alias=BOOL}(int32))
	{
		if !o.FireInParallel {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SetFireInParallelOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// fFireInParallel {in} (1:{alias=BOOL}(int32))
	{
		var _bFireInParallel int32
		if err := w.ReadData(&_bFireInParallel); err != nil {
			return err
		}
		o.FireInParallel = _bFireInParallel != 0
	}
	return nil
}

func (o *xxx_SetFireInParallelOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFireInParallelOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetFireInParallelOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetFireInParallelRequest structure represents the FireInParallel operation request
type SetFireInParallelRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	FireInParallel bool           `idl:"name:fFireInParallel" json:"fire_in_parallel"`
}

func (o *SetFireInParallelRequest) xxx_ToOp(ctx context.Context, op *xxx_SetFireInParallelOperation) *xxx_SetFireInParallelOperation {
	if op == nil {
		op = &xxx_SetFireInParallelOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.FireInParallel = o.FireInParallel
	return op
}

func (o *SetFireInParallelRequest) xxx_FromOp(ctx context.Context, op *xxx_SetFireInParallelOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.FireInParallel = op.FireInParallel
}
func (o *SetFireInParallelRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetFireInParallelRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFireInParallelOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetFireInParallelResponse structure represents the FireInParallel operation response
type SetFireInParallelResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The FireInParallel return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetFireInParallelResponse) xxx_ToOp(ctx context.Context, op *xxx_SetFireInParallelOperation) *xxx_SetFireInParallelOperation {
	if op == nil {
		op = &xxx_SetFireInParallelOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetFireInParallelResponse) xxx_FromOp(ctx context.Context, op *xxx_SetFireInParallelOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetFireInParallelResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetFireInParallelResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFireInParallelOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
