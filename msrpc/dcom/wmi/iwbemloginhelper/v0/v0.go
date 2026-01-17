package iwbemloginhelper

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wmi"
)

var (
	// IWbemLoginHelper interface identifier 541679ab-2e5f-11d3-b34e-00104bcc4b4a
	LoginHelperIID = &dcom.IID{Data1: 0x541679ab, Data2: 0x2e5f, Data3: 0x11d3, Data4: []byte{0xb3, 0x4e, 0x00, 0x10, 0x4b, 0xcc, 0x4b, 0x4a}}
	// Syntax UUID
	LoginHelperSyntaxUUID = &uuid.UUID{TimeLow: 0x541679ab, TimeMid: 0x2e5f, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0x4e, Node: [6]uint8{0x0, 0x10, 0x4b, 0xcc, 0x4b, 0x4a}}
	// Syntax ID
	LoginHelperSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: LoginHelperSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWbemLoginHelper interface.
type LoginHelperClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The IWbemLoginHelper::SetEvent MUST return WBEM_S_NO_ERROR. The SetEvent method SHOULD
	// NOT perform any action.<57>
	//
	// The opnum of the SetEvent method equals 3.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	//
	// If the method fails, the server MUST return an HRESULT whose S (severity) bit is
	// set as specified in [MS-ERREF] section 2.1. The actual HRESULT value is implementation
	// dependent.
	SetEvent(context.Context, *SetEventRequest, ...dcerpc.CallOption) (*SetEventResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) LoginHelperClient
}

type xxx_DefaultLoginHelperClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultLoginHelperClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultLoginHelperClient) SetEvent(ctx context.Context, in *SetEventRequest, opts ...dcerpc.CallOption) (*SetEventResponse, error) {
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
	out := &SetEventResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultLoginHelperClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultLoginHelperClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultLoginHelperClient) IPID(ctx context.Context, ipid *dcom.IPID) LoginHelperClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultLoginHelperClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewLoginHelperClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (LoginHelperClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(LoginHelperSyntaxV0_0))...)
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
	return &xxx_DefaultLoginHelperClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_SetEventOperation structure represents the SetEvent operation
type xxx_SetEventOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	EventToSet string         `idl:"name:sEventToSet" json:"event_to_set"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetEventOperation) OpNum() int { return 3 }

func (o *xxx_SetEventOperation) OpName() string { return "/IWbemLoginHelper/v0/SetEvent" }

func (o *xxx_SetEventOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// sEventToSet {in} (1:{alias=LPCSTR}*(1)[dim:0,string](char))
	{
		if err := ndr.WriteCharString(ctx, w, o.EventToSet); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// sEventToSet {in} (1:{alias=LPCSTR,pointer=ref}*(1)[dim:0,string](char))
	{
		if err := ndr.ReadCharString(ctx, w, &o.EventToSet); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetEventOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetEventRequest structure represents the SetEvent operation request
type SetEventRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// sEventToSet: MUST contain the name of the event to be signaled. This parameter MUST
	// NOT be NULL.
	EventToSet string `idl:"name:sEventToSet" json:"event_to_set"`
}

func (o *SetEventRequest) xxx_ToOp(ctx context.Context, op *xxx_SetEventOperation) *xxx_SetEventOperation {
	if op == nil {
		op = &xxx_SetEventOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.EventToSet = o.EventToSet
	return op
}

func (o *SetEventRequest) xxx_FromOp(ctx context.Context, op *xxx_SetEventOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.EventToSet = op.EventToSet
}
func (o *SetEventRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetEventRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEventOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetEventResponse structure represents the SetEvent operation response
type SetEventResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetEvent return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetEventResponse) xxx_ToOp(ctx context.Context, op *xxx_SetEventOperation) *xxx_SetEventOperation {
	if op == nil {
		op = &xxx_SetEventOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetEventResponse) xxx_FromOp(ctx context.Context, op *xxx_SetEventOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetEventResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetEventResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetEventOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
