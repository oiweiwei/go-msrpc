package ieventsysteminitialize

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
)

var (
	// import guard
	GoPackage = "dcom/comev"
)

var (
	// IEventSystemInitialize interface identifier a0e8f27a-888c-11d1-b763-00c04fb926af
	EventSystemInitializeIID = &dcom.IID{Data1: 0xa0e8f27a, Data2: 0x888c, Data3: 0x11d1, Data4: []byte{0xb7, 0x63, 0x00, 0xc0, 0x4f, 0xb9, 0x26, 0xaf}}
	// Syntax UUID
	EventSystemInitializeSyntaxUUID = &uuid.UUID{TimeLow: 0xa0e8f27a, TimeMid: 0x888c, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xb7, ClockSeqLow: 0x63, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xb9, 0x26, 0xaf}}
	// Syntax ID
	EventSystemInitializeSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: EventSystemInitializeSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IEventSystemInitialize interface.
type EventSystemInitializeClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The SetCOMCatalogBehaviour method controls the event system CatalogMode and RetainSubKeys
	// state variables.
	//
	// Return Values: The server MUST return S_OK.
	//
	// After this method is called, the event system CatalogMode state variable will be
	// true (server in catalog mode) and the RetainSubKeys variable will be set based on
	// the bRetainSubKeys parameter. If the client does not call this method, the server
	// will be in non-catalog mode. The Store, Remove, and RemoveS methods of IEventSystem
	// will have different behavior when the server is in catalog mode. See section 3.1.1.3
	// for more details.
	SetCOMCatalogBehaviour(context.Context, *SetCOMCatalogBehaviourRequest, ...dcerpc.CallOption) (*SetCOMCatalogBehaviourResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) EventSystemInitializeClient
}

type xxx_DefaultEventSystemInitializeClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultEventSystemInitializeClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultEventSystemInitializeClient) SetCOMCatalogBehaviour(ctx context.Context, in *SetCOMCatalogBehaviourRequest, opts ...dcerpc.CallOption) (*SetCOMCatalogBehaviourResponse, error) {
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
	out := &SetCOMCatalogBehaviourResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSystemInitializeClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultEventSystemInitializeClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultEventSystemInitializeClient) IPID(ctx context.Context, ipid *dcom.IPID) EventSystemInitializeClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultEventSystemInitializeClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewEventSystemInitializeClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (EventSystemInitializeClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(EventSystemInitializeSyntaxV0_0))...)
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
	return &xxx_DefaultEventSystemInitializeClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_SetCOMCatalogBehaviourOperation structure represents the SetCOMCatalogBehaviour operation
type xxx_SetCOMCatalogBehaviourOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	RetainSubKeys bool           `idl:"name:bRetainSubKeys" json:"retain_sub_keys"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetCOMCatalogBehaviourOperation) OpNum() int { return 3 }

func (o *xxx_SetCOMCatalogBehaviourOperation) OpName() string {
	return "/IEventSystemInitialize/v0/SetCOMCatalogBehaviour"
}

func (o *xxx_SetCOMCatalogBehaviourOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCOMCatalogBehaviourOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bRetainSubKeys {in} (1:{alias=BOOL}(int32))
	{
		if !o.RetainSubKeys {
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

func (o *xxx_SetCOMCatalogBehaviourOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bRetainSubKeys {in} (1:{alias=BOOL}(int32))
	{
		var _bRetainSubKeys int32
		if err := w.ReadData(&_bRetainSubKeys); err != nil {
			return err
		}
		o.RetainSubKeys = _bRetainSubKeys != 0
	}
	return nil
}

func (o *xxx_SetCOMCatalogBehaviourOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCOMCatalogBehaviourOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetCOMCatalogBehaviourOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetCOMCatalogBehaviourRequest structure represents the SetCOMCatalogBehaviour operation request
type SetCOMCatalogBehaviourRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bRetainSubKeys: This value is copied to the RetainSubKeys state variable (see section
	// 3.1.1.3) of the event system.
	RetainSubKeys bool `idl:"name:bRetainSubKeys" json:"retain_sub_keys"`
}

func (o *SetCOMCatalogBehaviourRequest) xxx_ToOp(ctx context.Context) *xxx_SetCOMCatalogBehaviourOperation {
	if o == nil {
		return &xxx_SetCOMCatalogBehaviourOperation{}
	}
	return &xxx_SetCOMCatalogBehaviourOperation{
		This:          o.This,
		RetainSubKeys: o.RetainSubKeys,
	}
}

func (o *SetCOMCatalogBehaviourRequest) xxx_FromOp(ctx context.Context, op *xxx_SetCOMCatalogBehaviourOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.RetainSubKeys = op.RetainSubKeys
}
func (o *SetCOMCatalogBehaviourRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetCOMCatalogBehaviourRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCOMCatalogBehaviourOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetCOMCatalogBehaviourResponse structure represents the SetCOMCatalogBehaviour operation response
type SetCOMCatalogBehaviourResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetCOMCatalogBehaviour return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetCOMCatalogBehaviourResponse) xxx_ToOp(ctx context.Context) *xxx_SetCOMCatalogBehaviourOperation {
	if o == nil {
		return &xxx_SetCOMCatalogBehaviourOperation{}
	}
	return &xxx_SetCOMCatalogBehaviourOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetCOMCatalogBehaviourResponse) xxx_FromOp(ctx context.Context, op *xxx_SetCOMCatalogBehaviourOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetCOMCatalogBehaviourResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetCOMCatalogBehaviourResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCOMCatalogBehaviourOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
