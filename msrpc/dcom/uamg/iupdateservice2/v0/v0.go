package iupdateservice2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iupdateservice "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdateservice/v0"
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
	_ = iupdateservice.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/uamg"
)

var (
	// IUpdateService2 interface identifier 1518b460-6518-4172-940f-c75883b24ceb
	UpdateService2IID = &dcom.IID{Data1: 0x1518b460, Data2: 0x6518, Data3: 0x4172, Data4: []byte{0x94, 0x0f, 0xc7, 0x58, 0x83, 0xb2, 0x4c, 0xeb}}
	// Syntax UUID
	UpdateService2SyntaxUUID = &uuid.UUID{TimeLow: 0x1518b460, TimeMid: 0x6518, TimeHiAndVersion: 0x4172, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0xf, Node: [6]uint8{0xc7, 0x58, 0x83, 0xb2, 0x4c, 0xeb}}
	// Syntax ID
	UpdateService2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: UpdateService2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IUpdateService2 interface.
type UpdateService2Client interface {

	// IUpdateService retrieval method.
	UpdateService() iupdateservice.UpdateServiceClient

	GetIsDefaultAUService(context.Context, *GetIsDefaultAUServiceRequest, ...dcerpc.CallOption) (*GetIsDefaultAUServiceResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) UpdateService2Client
}

type xxx_DefaultUpdateService2Client struct {
	iupdateservice.UpdateServiceClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultUpdateService2Client) UpdateService() iupdateservice.UpdateServiceClient {
	return o.UpdateServiceClient
}

func (o *xxx_DefaultUpdateService2Client) GetIsDefaultAUService(ctx context.Context, in *GetIsDefaultAUServiceRequest, opts ...dcerpc.CallOption) (*GetIsDefaultAUServiceResponse, error) {
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
	out := &GetIsDefaultAUServiceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateService2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultUpdateService2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultUpdateService2Client) IPID(ctx context.Context, ipid *dcom.IPID) UpdateService2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultUpdateService2Client{
		UpdateServiceClient: o.UpdateServiceClient.IPID(ctx, ipid),
		cc:                  o.cc,
		ipid:                ipid,
	}
}

func NewUpdateService2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (UpdateService2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(UpdateService2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iupdateservice.NewUpdateServiceClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultUpdateService2Client{
		UpdateServiceClient: base,
		cc:                  cc,
		ipid:                ipid,
	}, nil
}

// xxx_GetIsDefaultAUServiceOperation structure represents the IsDefaultAUService operation
type xxx_GetIsDefaultAUServiceOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue int16          `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIsDefaultAUServiceOperation) OpNum() int { return 20 }

func (o *xxx_GetIsDefaultAUServiceOperation) OpName() string {
	return "/IUpdateService2/v0/IsDefaultAUService"
}

func (o *xxx_GetIsDefaultAUServiceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsDefaultAUServiceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIsDefaultAUServiceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIsDefaultAUServiceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsDefaultAUServiceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.ReturnValue); err != nil {
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

func (o *xxx_GetIsDefaultAUServiceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.ReturnValue); err != nil {
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

// GetIsDefaultAUServiceRequest structure represents the IsDefaultAUService operation request
type GetIsDefaultAUServiceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIsDefaultAUServiceRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIsDefaultAUServiceOperation) *xxx_GetIsDefaultAUServiceOperation {
	if op == nil {
		op = &xxx_GetIsDefaultAUServiceOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIsDefaultAUServiceRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIsDefaultAUServiceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIsDefaultAUServiceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIsDefaultAUServiceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsDefaultAUServiceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIsDefaultAUServiceResponse structure represents the IsDefaultAUService operation response
type GetIsDefaultAUServiceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue int16          `idl:"name:retval" json:"return_value"`
	// Return: The IsDefaultAUService return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIsDefaultAUServiceResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIsDefaultAUServiceOperation) *xxx_GetIsDefaultAUServiceOperation {
	if op == nil {
		op = &xxx_GetIsDefaultAUServiceOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetIsDefaultAUServiceResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIsDefaultAUServiceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetIsDefaultAUServiceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIsDefaultAUServiceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsDefaultAUServiceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
