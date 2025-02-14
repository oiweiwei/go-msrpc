package icluscfgasyncevictcleanup

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
	GoPackage = "dcom/ccfg"
)

var (
	// IClusCfgAsyncEvictCleanup interface identifier 52c80b95-c1ad-4240-8d89-72e9fa84025e
	AsyncEvictCleanupIID = &dcom.IID{Data1: 0x52c80b95, Data2: 0xc1ad, Data3: 0x4240, Data4: []byte{0x8d, 0x89, 0x72, 0xe9, 0xfa, 0x84, 0x02, 0x5e}}
	// Syntax UUID
	AsyncEvictCleanupSyntaxUUID = &uuid.UUID{TimeLow: 0x52c80b95, TimeMid: 0xc1ad, TimeHiAndVersion: 0x4240, ClockSeqHiAndReserved: 0x8d, ClockSeqLow: 0x89, Node: [6]uint8{0x72, 0xe9, 0xfa, 0x84, 0x2, 0x5e}}
	// Syntax ID
	AsyncEvictCleanupSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AsyncEvictCleanupSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IClusCfgAsyncEvictCleanup interface.
type AsyncEvictCleanupClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// CleanupNode operation.
	CleanupNode(context.Context, *CleanupNodeRequest, ...dcerpc.CallOption) (*CleanupNodeResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AsyncEvictCleanupClient
}

type xxx_DefaultAsyncEvictCleanupClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAsyncEvictCleanupClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultAsyncEvictCleanupClient) CleanupNode(ctx context.Context, in *CleanupNodeRequest, opts ...dcerpc.CallOption) (*CleanupNodeResponse, error) {
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
	out := &CleanupNodeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAsyncEvictCleanupClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAsyncEvictCleanupClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAsyncEvictCleanupClient) IPID(ctx context.Context, ipid *dcom.IPID) AsyncEvictCleanupClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAsyncEvictCleanupClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewAsyncEvictCleanupClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AsyncEvictCleanupClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AsyncEvictCleanupSyntaxV0_0))...)
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
	return &xxx_DefaultAsyncEvictCleanupClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_CleanupNodeOperation structure represents the CleanupNode operation
type xxx_CleanupNodeOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	EvictedNodeNameIn *oaut.String   `idl:"name:bstrEvictedNodeNameIn" json:"evicted_node_name_in"`
	DelayIn           int32          `idl:"name:nDelayIn" json:"delay_in"`
	TimeoutIn         int32          `idl:"name:nTimeoutIn" json:"timeout_in"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CleanupNodeOperation) OpNum() int { return 7 }

func (o *xxx_CleanupNodeOperation) OpName() string {
	return "/IClusCfgAsyncEvictCleanup/v0/CleanupNode"
}

func (o *xxx_CleanupNodeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanupNodeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrEvictedNodeNameIn {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.EvictedNodeNameIn != nil {
			_ptr_bstrEvictedNodeNameIn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.EvictedNodeNameIn != nil {
					if err := o.EvictedNodeNameIn.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.EvictedNodeNameIn, _ptr_bstrEvictedNodeNameIn); err != nil {
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
	// nDelayIn {in} (1:(int32))
	{
		if err := w.WriteData(o.DelayIn); err != nil {
			return err
		}
	}
	// nTimeoutIn {in} (1:(int32))
	{
		if err := w.WriteData(o.TimeoutIn); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanupNodeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrEvictedNodeNameIn {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrEvictedNodeNameIn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.EvictedNodeNameIn == nil {
				o.EvictedNodeNameIn = &oaut.String{}
			}
			if err := o.EvictedNodeNameIn.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrEvictedNodeNameIn := func(ptr interface{}) { o.EvictedNodeNameIn = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.EvictedNodeNameIn, _s_bstrEvictedNodeNameIn, _ptr_bstrEvictedNodeNameIn); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// nDelayIn {in} (1:(int32))
	{
		if err := w.ReadData(&o.DelayIn); err != nil {
			return err
		}
	}
	// nTimeoutIn {in} (1:(int32))
	{
		if err := w.ReadData(&o.TimeoutIn); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanupNodeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanupNodeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CleanupNodeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CleanupNodeRequest structure represents the CleanupNode operation request
type CleanupNodeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	EvictedNodeNameIn *oaut.String   `idl:"name:bstrEvictedNodeNameIn" json:"evicted_node_name_in"`
	DelayIn           int32          `idl:"name:nDelayIn" json:"delay_in"`
	TimeoutIn         int32          `idl:"name:nTimeoutIn" json:"timeout_in"`
}

func (o *CleanupNodeRequest) xxx_ToOp(ctx context.Context, op *xxx_CleanupNodeOperation) *xxx_CleanupNodeOperation {
	if op == nil {
		op = &xxx_CleanupNodeOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.EvictedNodeNameIn = op.EvictedNodeNameIn
	o.DelayIn = op.DelayIn
	o.TimeoutIn = op.TimeoutIn
	return op
}

func (o *CleanupNodeRequest) xxx_FromOp(ctx context.Context, op *xxx_CleanupNodeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.EvictedNodeNameIn = op.EvictedNodeNameIn
	o.DelayIn = op.DelayIn
	o.TimeoutIn = op.TimeoutIn
}
func (o *CleanupNodeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CleanupNodeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CleanupNodeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CleanupNodeResponse structure represents the CleanupNode operation response
type CleanupNodeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CleanupNode return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CleanupNodeResponse) xxx_ToOp(ctx context.Context, op *xxx_CleanupNodeOperation) *xxx_CleanupNodeOperation {
	if op == nil {
		op = &xxx_CleanupNodeOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *CleanupNodeResponse) xxx_FromOp(ctx context.Context, op *xxx_CleanupNodeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CleanupNodeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CleanupNodeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CleanupNodeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
