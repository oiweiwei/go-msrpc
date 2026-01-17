package imsmqtransactiondispenser

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	mqac "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac"
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
	_ = dcom.GoPackage
	_ = idispatch.GoPackage
	_ = mqac.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

var (
	// IMSMQTransactionDispenser interface identifier d7d6e083-dccd-11d0-aa4b-0060970debae
	TransactionDispenserIID = &dcom.IID{Data1: 0xd7d6e083, Data2: 0xdccd, Data3: 0x11d0, Data4: []byte{0xaa, 0x4b, 0x00, 0x60, 0x97, 0x0d, 0xeb, 0xae}}
	// Syntax UUID
	TransactionDispenserSyntaxUUID = &uuid.UUID{TimeLow: 0xd7d6e083, TimeMid: 0xdccd, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x4b, Node: [6]uint8{0x0, 0x60, 0x97, 0xd, 0xeb, 0xae}}
	// Syntax ID
	TransactionDispenserSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: TransactionDispenserSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQTransactionDispenser interface.
type TransactionDispenserClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// BeginTransaction operation.
	BeginTransaction(context.Context, *BeginTransactionRequest, ...dcerpc.CallOption) (*BeginTransactionResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) TransactionDispenserClient
}

type xxx_DefaultTransactionDispenserClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultTransactionDispenserClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultTransactionDispenserClient) BeginTransaction(ctx context.Context, in *BeginTransactionRequest, opts ...dcerpc.CallOption) (*BeginTransactionResponse, error) {
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
	out := &BeginTransactionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransactionDispenserClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultTransactionDispenserClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultTransactionDispenserClient) IPID(ctx context.Context, ipid *dcom.IPID) TransactionDispenserClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultTransactionDispenserClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewTransactionDispenserClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (TransactionDispenserClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(TransactionDispenserSyntaxV0_0))...)
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
	return &xxx_DefaultTransactionDispenserClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_BeginTransactionOperation structure represents the BeginTransaction operation
type xxx_BeginTransactionOperation struct {
	This        *dcom.ORPCThis    `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat    `idl:"name:That" json:"that"`
	Transaction *mqac.Transaction `idl:"name:ptransaction" json:"transaction"`
	Return      int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_BeginTransactionOperation) OpNum() int { return 7 }

func (o *xxx_BeginTransactionOperation) OpName() string {
	return "/IMSMQTransactionDispenser/v0/BeginTransaction"
}

func (o *xxx_BeginTransactionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BeginTransactionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BeginTransactionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_BeginTransactionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BeginTransactionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ptransaction {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQTransaction}(interface))
	{
		if o.Transaction != nil {
			_ptr_ptransaction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Transaction != nil {
					if err := o.Transaction.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Transaction{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Transaction, _ptr_ptransaction); err != nil {
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

func (o *xxx_BeginTransactionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ptransaction {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQTransaction}(interface))
	{
		_ptr_ptransaction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Transaction == nil {
				o.Transaction = &mqac.Transaction{}
			}
			if err := o.Transaction.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ptransaction := func(ptr interface{}) { o.Transaction = *ptr.(**mqac.Transaction) }
		if err := w.ReadPointer(&o.Transaction, _s_ptransaction, _ptr_ptransaction); err != nil {
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

// BeginTransactionRequest structure represents the BeginTransaction operation request
type BeginTransactionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *BeginTransactionRequest) xxx_ToOp(ctx context.Context, op *xxx_BeginTransactionOperation) *xxx_BeginTransactionOperation {
	if op == nil {
		op = &xxx_BeginTransactionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *BeginTransactionRequest) xxx_FromOp(ctx context.Context, op *xxx_BeginTransactionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *BeginTransactionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BeginTransactionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BeginTransactionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BeginTransactionResponse structure represents the BeginTransaction operation response
type BeginTransactionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat    `idl:"name:That" json:"that"`
	Transaction *mqac.Transaction `idl:"name:ptransaction" json:"transaction"`
	// Return: The BeginTransaction return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BeginTransactionResponse) xxx_ToOp(ctx context.Context, op *xxx_BeginTransactionOperation) *xxx_BeginTransactionOperation {
	if op == nil {
		op = &xxx_BeginTransactionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Transaction = o.Transaction
	op.Return = o.Return
	return op
}

func (o *BeginTransactionResponse) xxx_FromOp(ctx context.Context, op *xxx_BeginTransactionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Transaction = op.Transaction
	o.Return = op.Return
}
func (o *BeginTransactionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BeginTransactionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BeginTransactionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
