package iautomaticupdates

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
)

var (
	// import guard
	GoPackage = "dcom/uamg"
)

var (
	// IAutomaticUpdates interface identifier 673425bf-c082-4c7c-bdfd-569464b8e0ce
	AutomaticUpdatesIID = &dcom.IID{Data1: 0x673425bf, Data2: 0xc082, Data3: 0x4c7c, Data4: []byte{0xbd, 0xfd, 0x56, 0x94, 0x64, 0xb8, 0xe0, 0xce}}
	// Syntax UUID
	AutomaticUpdatesSyntaxUUID = &uuid.UUID{TimeLow: 0x673425bf, TimeMid: 0xc082, TimeHiAndVersion: 0x4c7c, ClockSeqHiAndReserved: 0xbd, ClockSeqLow: 0xfd, Node: [6]uint8{0x56, 0x94, 0x64, 0xb8, 0xe0, 0xce}}
	// Syntax ID
	AutomaticUpdatesSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AutomaticUpdatesSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IAutomaticUpdates interface.
type AutomaticUpdatesClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The server SHOULD trigger the automatic update agent to perform an update search.
	DetectNow(context.Context, *DetectNowRequest, ...dcerpc.CallOption) (*DetectNowResponse, error)

	// Opnum9NotUsedOnWire operation.
	// Opnum9NotUsedOnWire

	// Opnum10NotUsedOnWire operation.
	// Opnum10NotUsedOnWire

	// Opnum11NotUsedOnWire operation.
	// Opnum11NotUsedOnWire

	// Opnum12NotUsedOnWire operation.
	// Opnum12NotUsedOnWire

	// Opnum13NotUsedOnWire operation.
	// Opnum13NotUsedOnWire

	// Opnum14NotUsedOnWire operation.
	// Opnum14NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AutomaticUpdatesClient
}

type xxx_DefaultAutomaticUpdatesClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAutomaticUpdatesClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultAutomaticUpdatesClient) DetectNow(ctx context.Context, in *DetectNowRequest, opts ...dcerpc.CallOption) (*DetectNowResponse, error) {
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
	out := &DetectNowResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAutomaticUpdatesClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAutomaticUpdatesClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAutomaticUpdatesClient) IPID(ctx context.Context, ipid *dcom.IPID) AutomaticUpdatesClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAutomaticUpdatesClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewAutomaticUpdatesClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AutomaticUpdatesClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AutomaticUpdatesSyntaxV0_0))...)
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
	return &xxx_DefaultAutomaticUpdatesClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_DetectNowOperation structure represents the DetectNow operation
type xxx_DetectNowOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DetectNowOperation) OpNum() int { return 7 }

func (o *xxx_DetectNowOperation) OpName() string { return "/IAutomaticUpdates/v0/DetectNow" }

func (o *xxx_DetectNowOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DetectNowOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DetectNowOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_DetectNowOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DetectNowOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DetectNowOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DetectNowRequest structure represents the DetectNow operation request
type DetectNowRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *DetectNowRequest) xxx_ToOp(ctx context.Context, op *xxx_DetectNowOperation) *xxx_DetectNowOperation {
	if op == nil {
		op = &xxx_DetectNowOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *DetectNowRequest) xxx_FromOp(ctx context.Context, op *xxx_DetectNowOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *DetectNowRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DetectNowRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DetectNowOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DetectNowResponse structure represents the DetectNow operation response
type DetectNowResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DetectNow return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DetectNowResponse) xxx_ToOp(ctx context.Context, op *xxx_DetectNowOperation) *xxx_DetectNowOperation {
	if op == nil {
		op = &xxx_DetectNowOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DetectNowResponse) xxx_FromOp(ctx context.Context, op *xxx_DetectNowOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DetectNowResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DetectNowResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DetectNowOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
