package ieventsystem2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	ieventsystem "github.com/oiweiwei/go-msrpc/msrpc/dcom/comev/ieventsystem/v0"
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
	_ = ieventsystem.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/comev"
)

var (
	// IEventSystem2 interface identifier 99cc098f-a48a-4e9c-8e58-965c0afc19d5
	EventSystem2IID = &dcom.IID{Data1: 0x99cc098f, Data2: 0xa48a, Data3: 0x4e9c, Data4: []byte{0x8e, 0x58, 0x96, 0x5c, 0x0a, 0xfc, 0x19, 0xd5}}
	// Syntax UUID
	EventSystem2SyntaxUUID = &uuid.UUID{TimeLow: 0x99cc098f, TimeMid: 0xa48a, TimeHiAndVersion: 0x4e9c, ClockSeqHiAndReserved: 0x8e, ClockSeqLow: 0x58, Node: [6]uint8{0x96, 0x5c, 0xa, 0xfc, 0x19, 0xd5}}
	// Syntax ID
	EventSystem2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: EventSystem2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IEventSystem2 interface.
type EventSystem2Client interface {

	// IEventSystem retrieval method.
	EventSystem() ieventsystem.EventSystemClient

	// The GetVersion method retrieves the version of the server implementation of the protocol.
	//
	// Return Values: An HRESULT specifying success or failure. All success codes MUST be
	// treated the same, and all failure codes MUST be treated the same.
	GetVersion(context.Context, *GetVersionRequest, ...dcerpc.CallOption) (*GetVersionResponse, error)

	// The VerifyTransientSubscribers method verifies the state of the transient subscribers.
	//
	// This method has no parameters.
	//
	// Return Values: This function MUST return S_OK.
	VerifyTransientSubscribers(context.Context, *VerifyTransientSubscribersRequest, ...dcerpc.CallOption) (*VerifyTransientSubscribersResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) EventSystem2Client
}

type xxx_DefaultEventSystem2Client struct {
	ieventsystem.EventSystemClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultEventSystem2Client) EventSystem() ieventsystem.EventSystemClient {
	return o.EventSystemClient
}

func (o *xxx_DefaultEventSystem2Client) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...dcerpc.CallOption) (*GetVersionResponse, error) {
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
	out := &GetVersionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSystem2Client) VerifyTransientSubscribers(ctx context.Context, in *VerifyTransientSubscribersRequest, opts ...dcerpc.CallOption) (*VerifyTransientSubscribersResponse, error) {
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
	out := &VerifyTransientSubscribersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultEventSystem2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultEventSystem2Client) IPID(ctx context.Context, ipid *dcom.IPID) EventSystem2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultEventSystem2Client{
		EventSystemClient: o.EventSystemClient.IPID(ctx, ipid),
		cc:                o.cc,
		ipid:              ipid,
	}
}
func NewEventSystem2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (EventSystem2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(EventSystem2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ieventsystem.NewEventSystemClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultEventSystem2Client{
		EventSystemClient: base,
		cc:                cc,
		ipid:              ipid,
	}, nil
}

// xxx_GetVersionOperation structure represents the GetVersion operation
type xxx_GetVersionOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Version int32          `idl:"name:pnVersion" json:"version"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetVersionOperation) OpNum() int { return 13 }

func (o *xxx_GetVersionOperation) OpName() string { return "/IEventSystem2/v0/GetVersion" }

func (o *xxx_GetVersionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetVersionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetVersionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pnVersion {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Version); err != nil {
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

func (o *xxx_GetVersionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pnVersion {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Version); err != nil {
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

// GetVersionRequest structure represents the GetVersion operation request
type GetVersionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetVersionRequest) xxx_ToOp(ctx context.Context) *xxx_GetVersionOperation {
	if o == nil {
		return &xxx_GetVersionOperation{}
	}
	return &xxx_GetVersionOperation{
		This: o.This,
	}
}

func (o *GetVersionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetVersionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetVersionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetVersionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVersionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetVersionResponse structure represents the GetVersion operation response
type GetVersionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pnVersion: If the function returns a success HRESULT, it MUST contain one of the
	// following:
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 | The server does not support the IEventSubscription3 and IEventClass3 interfaces, |
	//	|            | nor does it support the PartitionID and ApplicationID properties on the          |
	//	|            | subscription (section 3.1.1.2) and event class (section 3.1.1.1) objects.        |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 | The server supports the IEventSubscription3 and IEventClass3 interfaces. It      |
	//	|            | also supports the PartitionID and ApplicationID properties on the subscription   |
	//	|            | (section 3.1.1.2) and event class (section 3.1.1.1) objects.                     |
	//	+------------+----------------------------------------------------------------------------------+
	Version int32 `idl:"name:pnVersion" json:"version"`
	// Return: The GetVersion return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetVersionResponse) xxx_ToOp(ctx context.Context) *xxx_GetVersionOperation {
	if o == nil {
		return &xxx_GetVersionOperation{}
	}
	return &xxx_GetVersionOperation{
		That:    o.That,
		Version: o.Version,
		Return:  o.Return,
	}
}

func (o *GetVersionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetVersionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Version = op.Version
	o.Return = op.Return
}
func (o *GetVersionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetVersionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVersionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_VerifyTransientSubscribersOperation structure represents the VerifyTransientSubscribers operation
type xxx_VerifyTransientSubscribersOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_VerifyTransientSubscribersOperation) OpNum() int { return 14 }

func (o *xxx_VerifyTransientSubscribersOperation) OpName() string {
	return "/IEventSystem2/v0/VerifyTransientSubscribers"
}

func (o *xxx_VerifyTransientSubscribersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_VerifyTransientSubscribersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_VerifyTransientSubscribersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_VerifyTransientSubscribersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_VerifyTransientSubscribersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_VerifyTransientSubscribersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// VerifyTransientSubscribersRequest structure represents the VerifyTransientSubscribers operation request
type VerifyTransientSubscribersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *VerifyTransientSubscribersRequest) xxx_ToOp(ctx context.Context) *xxx_VerifyTransientSubscribersOperation {
	if o == nil {
		return &xxx_VerifyTransientSubscribersOperation{}
	}
	return &xxx_VerifyTransientSubscribersOperation{
		This: o.This,
	}
}

func (o *VerifyTransientSubscribersRequest) xxx_FromOp(ctx context.Context, op *xxx_VerifyTransientSubscribersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *VerifyTransientSubscribersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *VerifyTransientSubscribersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_VerifyTransientSubscribersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// VerifyTransientSubscribersResponse structure represents the VerifyTransientSubscribers operation response
type VerifyTransientSubscribersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The VerifyTransientSubscribers return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *VerifyTransientSubscribersResponse) xxx_ToOp(ctx context.Context) *xxx_VerifyTransientSubscribersOperation {
	if o == nil {
		return &xxx_VerifyTransientSubscribersOperation{}
	}
	return &xxx_VerifyTransientSubscribersOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *VerifyTransientSubscribersResponse) xxx_FromOp(ctx context.Context, op *xxx_VerifyTransientSubscribersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *VerifyTransientSubscribersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *VerifyTransientSubscribersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_VerifyTransientSubscribersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
