package imsmqcoordinatedtransactiondispenser3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	mqac "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac"
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
	_ = mqac.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

var (
	// IMSMQCoordinatedTransactionDispenser3 interface identifier eba96b14-2168-11d3-898c-00e02c074f6b
	ImsmqCoordinatedTransactionDispenser3IID = &dcom.IID{Data1: 0xeba96b14, Data2: 0x2168, Data3: 0x11d3, Data4: []byte{0x89, 0x8c, 0x00, 0xe0, 0x2c, 0x07, 0x4f, 0x6b}}
	// Syntax UUID
	ImsmqCoordinatedTransactionDispenser3SyntaxUUID = &uuid.UUID{TimeLow: 0xeba96b14, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	// Syntax ID
	ImsmqCoordinatedTransactionDispenser3SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ImsmqCoordinatedTransactionDispenser3SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQCoordinatedTransactionDispenser3 interface.
type ImsmqCoordinatedTransactionDispenser3Client interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// BeginTransaction operation.
	BeginTransaction(context.Context, *BeginTransactionRequest, ...dcerpc.CallOption) (*BeginTransactionResponse, error)

	// Properties operation.
	GetProperties(context.Context, *GetPropertiesRequest, ...dcerpc.CallOption) (*GetPropertiesResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ImsmqCoordinatedTransactionDispenser3Client
}

type xxx_DefaultImsmqCoordinatedTransactionDispenser3Client struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultImsmqCoordinatedTransactionDispenser3Client) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultImsmqCoordinatedTransactionDispenser3Client) BeginTransaction(ctx context.Context, in *BeginTransactionRequest, opts ...dcerpc.CallOption) (*BeginTransactionResponse, error) {
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
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqCoordinatedTransactionDispenser3Client) GetProperties(ctx context.Context, in *GetPropertiesRequest, opts ...dcerpc.CallOption) (*GetPropertiesResponse, error) {
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
	out := &GetPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultImsmqCoordinatedTransactionDispenser3Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultImsmqCoordinatedTransactionDispenser3Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultImsmqCoordinatedTransactionDispenser3Client) IPID(ctx context.Context, ipid *dcom.IPID) ImsmqCoordinatedTransactionDispenser3Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultImsmqCoordinatedTransactionDispenser3Client{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewImsmqCoordinatedTransactionDispenser3Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ImsmqCoordinatedTransactionDispenser3Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ImsmqCoordinatedTransactionDispenser3SyntaxV0_0))...)
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
	return &xxx_DefaultImsmqCoordinatedTransactionDispenser3Client{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_BeginTransactionOperation structure represents the BeginTransaction operation
type xxx_BeginTransactionOperation struct {
	This         *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Ptransaction *mqac.ImsmqTransaction3 `idl:"name:ptransaction" json:"ptransaction"`
	Return       int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_BeginTransactionOperation) OpNum() int { return 7 }

func (o *xxx_BeginTransactionOperation) OpName() string {
	return "/IMSMQCoordinatedTransactionDispenser3/v0/BeginTransaction"
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
	// ptransaction {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQTransaction3}(interface))
	{
		if o.Ptransaction != nil {
			_ptr_ptransaction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Ptransaction != nil {
					if err := o.Ptransaction.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.ImsmqTransaction3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Ptransaction, _ptr_ptransaction); err != nil {
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
	// ptransaction {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQTransaction3}(interface))
	{
		_ptr_ptransaction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Ptransaction == nil {
				o.Ptransaction = &mqac.ImsmqTransaction3{}
			}
			if err := o.Ptransaction.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ptransaction := func(ptr interface{}) { o.Ptransaction = *ptr.(**mqac.ImsmqTransaction3) }
		if err := w.ReadPointer(&o.Ptransaction, _s_ptransaction, _ptr_ptransaction); err != nil {
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
	That         *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Ptransaction *mqac.ImsmqTransaction3 `idl:"name:ptransaction" json:"ptransaction"`
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
	op.Ptransaction = o.Ptransaction
	op.Return = o.Return
	return op
}

func (o *BeginTransactionResponse) xxx_FromOp(ctx context.Context, op *xxx_BeginTransactionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Ptransaction = op.Ptransaction
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

// xxx_GetPropertiesOperation structure represents the Properties operation
type xxx_GetPropertiesOperation struct {
	This            *dcom.ORPCThis `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	PpcolProperties *oaut.Dispatch `idl:"name:ppcolProperties" json:"ppcol_properties"`
	Return          int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertiesOperation) OpNum() int { return 8 }

func (o *xxx_GetPropertiesOperation) OpName() string {
	return "/IMSMQCoordinatedTransactionDispenser3/v0/Properties"
}

func (o *xxx_GetPropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppcolProperties {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		if o.PpcolProperties != nil {
			_ptr_ppcolProperties := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PpcolProperties != nil {
					if err := o.PpcolProperties.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Dispatch{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PpcolProperties, _ptr_ppcolProperties); err != nil {
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

func (o *xxx_GetPropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppcolProperties {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IDispatch}(interface))
	{
		_ptr_ppcolProperties := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.PpcolProperties == nil {
				o.PpcolProperties = &oaut.Dispatch{}
			}
			if err := o.PpcolProperties.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppcolProperties := func(ptr interface{}) { o.PpcolProperties = *ptr.(**oaut.Dispatch) }
		if err := w.ReadPointer(&o.PpcolProperties, _s_ppcolProperties, _ptr_ppcolProperties); err != nil {
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

// GetPropertiesRequest structure represents the Properties operation request
type GetPropertiesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPropertiesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesOperation) *xxx_GetPropertiesOperation {
	if op == nil {
		op = &xxx_GetPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertiesResponse structure represents the Properties operation response
type GetPropertiesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat `idl:"name:That" json:"that"`
	PpcolProperties *oaut.Dispatch `idl:"name:ppcolProperties" json:"ppcol_properties"`
	// Return: The Properties return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertiesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesOperation) *xxx_GetPropertiesOperation {
	if op == nil {
		op = &xxx_GetPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PpcolProperties = o.PpcolProperties
	op.Return = o.Return
	return op
}

func (o *GetPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PpcolProperties = op.PpcolProperties
	o.Return = op.Return
}
func (o *GetPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
