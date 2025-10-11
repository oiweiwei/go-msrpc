package imsmqtransaction3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	imsmqtransaction2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac/imsmqtransaction2/v0"
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
	_ = imsmqtransaction2.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

var (
	// IMSMQTransaction3 interface identifier eba96b13-2168-11d3-898c-00e02c074f6b
	Transaction3IID = &dcom.IID{Data1: 0xeba96b13, Data2: 0x2168, Data3: 0x11d3, Data4: []byte{0x89, 0x8c, 0x00, 0xe0, 0x2c, 0x07, 0x4f, 0x6b}}
	// Syntax UUID
	Transaction3SyntaxUUID = &uuid.UUID{TimeLow: 0xeba96b13, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	// Syntax ID
	Transaction3SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: Transaction3SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQTransaction3 interface.
type Transaction3Client interface {

	// IMSMQTransaction2 retrieval method.
	Transaction2() imsmqtransaction2.Transaction2Client

	// The ITransaction method is received by the server in an RPC_REQUEST packet. In response,
	// the server returns the ITransaction interface on the underlying transaction object.
	//
	// Return Values: The method MUST return S_OK (0x00000000) to indicate success or an
	// implementation-specific error HRESULT on failure.
	//
	// The pvarITransaction output parameter MUST be set to the ITransaction interface pointer
	// of the Transaction instance variable.
	GetITransaction(context.Context, *GetITransactionRequest, ...dcerpc.CallOption) (*GetITransactionResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Transaction3Client
}

type xxx_DefaultTransaction3Client struct {
	imsmqtransaction2.Transaction2Client
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultTransaction3Client) Transaction2() imsmqtransaction2.Transaction2Client {
	return o.Transaction2Client
}

func (o *xxx_DefaultTransaction3Client) GetITransaction(ctx context.Context, in *GetITransactionRequest, opts ...dcerpc.CallOption) (*GetITransactionResponse, error) {
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
	out := &GetITransactionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTransaction3Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultTransaction3Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultTransaction3Client) IPID(ctx context.Context, ipid *dcom.IPID) Transaction3Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultTransaction3Client{
		Transaction2Client: o.Transaction2Client.IPID(ctx, ipid),
		cc:                 o.cc,
		ipid:               ipid,
	}
}

func NewTransaction3Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Transaction3Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Transaction3SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := imsmqtransaction2.NewTransaction2Client(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultTransaction3Client{
		Transaction2Client: base,
		cc:                 cc,
		ipid:               ipid,
	}, nil
}

// xxx_GetITransactionOperation structure represents the ITransaction operation
type xxx_GetITransactionOperation struct {
	This         *dcom.ORPCThis `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat `idl:"name:That" json:"that"`
	ITransaction *oaut.Variant  `idl:"name:pvarITransaction" json:"i_transaction"`
	Return       int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetITransactionOperation) OpNum() int { return 12 }

func (o *xxx_GetITransactionOperation) OpName() string { return "/IMSMQTransaction3/v0/ITransaction" }

func (o *xxx_GetITransactionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetITransactionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetITransactionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetITransactionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetITransactionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pvarITransaction {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ITransaction != nil {
			_ptr_pvarITransaction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ITransaction != nil {
					if err := o.ITransaction.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ITransaction, _ptr_pvarITransaction); err != nil {
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

func (o *xxx_GetITransactionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pvarITransaction {out, retval} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_pvarITransaction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ITransaction == nil {
				o.ITransaction = &oaut.Variant{}
			}
			if err := o.ITransaction.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pvarITransaction := func(ptr interface{}) { o.ITransaction = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ITransaction, _s_pvarITransaction, _ptr_pvarITransaction); err != nil {
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

// GetITransactionRequest structure represents the ITransaction operation request
type GetITransactionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetITransactionRequest) xxx_ToOp(ctx context.Context, op *xxx_GetITransactionOperation) *xxx_GetITransactionOperation {
	if op == nil {
		op = &xxx_GetITransactionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetITransactionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetITransactionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetITransactionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetITransactionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetITransactionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetITransactionResponse structure represents the ITransaction operation response
type GetITransactionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pvarITransaction: A pointer to a VARIANT (VT_UNKNOWN or VT_EMPTY) that, when successfully
	// completed, contains the underlying transaction object.
	ITransaction *oaut.Variant `idl:"name:pvarITransaction" json:"i_transaction"`
	// Return: The ITransaction return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetITransactionResponse) xxx_ToOp(ctx context.Context, op *xxx_GetITransactionOperation) *xxx_GetITransactionOperation {
	if op == nil {
		op = &xxx_GetITransactionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ITransaction = o.ITransaction
	op.Return = o.Return
	return op
}

func (o *GetITransactionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetITransactionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ITransaction = op.ITransaction
	o.Return = op.Return
}
func (o *GetITransactionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetITransactionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetITransactionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
