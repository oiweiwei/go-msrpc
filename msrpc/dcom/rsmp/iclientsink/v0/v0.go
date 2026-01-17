package iclientsink

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rsmp"
)

var (
	// IClientSink interface identifier 879c8bbe-41b0-11d1-be11-00c04fb6bf70
	ClientSinkIID = &dcom.IID{Data1: 0x879c8bbe, Data2: 0x41b0, Data3: 0x11d1, Data4: []byte{0xbe, 0x11, 0x00, 0xc0, 0x4f, 0xb6, 0xbf, 0x70}}
	// Syntax UUID
	ClientSinkSyntaxUUID = &uuid.UUID{TimeLow: 0x879c8bbe, TimeMid: 0x41b0, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xbe, ClockSeqLow: 0x11, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xb6, 0xbf, 0x70}}
	// Syntax ID
	ClientSinkSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ClientSinkSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IClientSink interface.
type ClientSinkClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// OnNotify operation.
	OnNotify(context.Context, *OnNotifyRequest, ...dcerpc.CallOption) (*OnNotifyResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ClientSinkClient
}

type xxx_DefaultClientSinkClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultClientSinkClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultClientSinkClient) OnNotify(ctx context.Context, in *OnNotifyRequest, opts ...dcerpc.CallOption) (*OnNotifyResponse, error) {
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
	out := &OnNotifyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClientSinkClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultClientSinkClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultClientSinkClient) IPID(ctx context.Context, ipid *dcom.IPID) ClientSinkClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultClientSinkClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewClientSinkClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ClientSinkClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ClientSinkSyntaxV0_0))...)
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
	return &xxx_DefaultClientSinkClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_OnNotifyOperation structure represents the OnNotify operation
type xxx_OnNotifyOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Type      uint32         `idl:"name:dwType" json:"type"`
	Operation uint32         `idl:"name:dwOperation" json:"operation"`
	ID        *dtyp.GUID     `idl:"name:lpIdentifier" json:"id"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_OnNotifyOperation) OpNum() int { return 0 }

func (o *xxx_OnNotifyOperation) OpName() string { return "/IClientSink/v0/OnNotify" }

func (o *xxx_OnNotifyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnNotifyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// dwOperation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Operation); err != nil {
			return err
		}
	}
	// lpIdentifier {in} (1:{alias=LPNTMS_GUID}*(1))(2:{alias=GUID}(struct))
	{
		if o.ID != nil {
			if err := o.ID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_OnNotifyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// dwOperation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Operation); err != nil {
			return err
		}
	}
	// lpIdentifier {in} (1:{alias=LPNTMS_GUID,pointer=ref}*(1))(2:{alias=GUID}(struct))
	{
		if o.ID == nil {
			o.ID = &dtyp.GUID{}
		}
		if err := o.ID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnNotifyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnNotifyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_OnNotifyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// OnNotifyRequest structure represents the OnNotify operation request
type OnNotifyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	Type      uint32         `idl:"name:dwType" json:"type"`
	Operation uint32         `idl:"name:dwOperation" json:"operation"`
	ID        *dtyp.GUID     `idl:"name:lpIdentifier" json:"id"`
}

func (o *OnNotifyRequest) xxx_ToOp(ctx context.Context, op *xxx_OnNotifyOperation) *xxx_OnNotifyOperation {
	if op == nil {
		op = &xxx_OnNotifyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Type = o.Type
	op.Operation = o.Operation
	op.ID = o.ID
	return op
}

func (o *OnNotifyRequest) xxx_FromOp(ctx context.Context, op *xxx_OnNotifyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Type = op.Type
	o.Operation = op.Operation
	o.ID = op.ID
}
func (o *OnNotifyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OnNotifyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OnNotifyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OnNotifyResponse structure represents the OnNotify operation response
type OnNotifyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The OnNotify return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OnNotifyResponse) xxx_ToOp(ctx context.Context, op *xxx_OnNotifyOperation) *xxx_OnNotifyOperation {
	if op == nil {
		op = &xxx_OnNotifyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *OnNotifyResponse) xxx_FromOp(ctx context.Context, op *xxx_OnNotifyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *OnNotifyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OnNotifyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OnNotifyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
