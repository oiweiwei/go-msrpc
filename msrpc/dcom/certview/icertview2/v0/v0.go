package icertview2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	icertview "github.com/oiweiwei/go-msrpc/msrpc/dcom/certview/icertview/v0"
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
	_ = icertview.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/certview"
)

var (
	// ICertView2 interface identifier d594b282-8851-4b61-9c66-3edadf848863
	CertView2IID = &dcom.IID{Data1: 0xd594b282, Data2: 0x8851, Data3: 0x4b61, Data4: []byte{0x9c, 0x66, 0x3e, 0xda, 0xdf, 0x84, 0x88, 0x63}}
	// Syntax UUID
	CertView2SyntaxUUID = &uuid.UUID{TimeLow: 0xd594b282, TimeMid: 0x8851, TimeHiAndVersion: 0x4b61, ClockSeqHiAndReserved: 0x9c, ClockSeqLow: 0x66, Node: [6]uint8{0x3e, 0xda, 0xdf, 0x84, 0x88, 0x63}}
	// Syntax ID
	CertView2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: CertView2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// ICertView2 interface.
type CertView2Client interface {

	// ICertView retrieval method.
	CertView() icertview.CertViewClient

	// SetTable operation.
	SetTable(context.Context, *SetTableRequest, ...dcerpc.CallOption) (*SetTableResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) CertView2Client
}

type xxx_DefaultCertView2Client struct {
	icertview.CertViewClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultCertView2Client) CertView() icertview.CertViewClient {
	return o.CertViewClient
}

func (o *xxx_DefaultCertView2Client) SetTable(ctx context.Context, in *SetTableRequest, opts ...dcerpc.CallOption) (*SetTableResponse, error) {
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
	out := &SetTableResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultCertView2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultCertView2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultCertView2Client) IPID(ctx context.Context, ipid *dcom.IPID) CertView2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultCertView2Client{
		CertViewClient: o.CertViewClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewCertView2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (CertView2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(CertView2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := icertview.NewCertViewClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultCertView2Client{
		CertViewClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_SetTableOperation structure represents the SetTable operation
type xxx_SetTableOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Table  int32          `idl:"name:Table" json:"table"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetTableOperation) OpNum() int { return 15 }

func (o *xxx_SetTableOperation) OpName() string { return "/ICertView2/v0/SetTable" }

func (o *xxx_SetTableOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTableOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Table {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Table); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTableOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Table {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Table); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTableOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTableOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetTableOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetTableRequest structure represents the SetTable operation request
type SetTableRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Table int32          `idl:"name:Table" json:"table"`
}

func (o *SetTableRequest) xxx_ToOp(ctx context.Context, op *xxx_SetTableOperation) *xxx_SetTableOperation {
	if op == nil {
		op = &xxx_SetTableOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Table = o.Table
	return op
}

func (o *SetTableRequest) xxx_FromOp(ctx context.Context, op *xxx_SetTableOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Table = op.Table
}
func (o *SetTableRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetTableRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTableOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetTableResponse structure represents the SetTable operation response
type SetTableResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetTable return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetTableResponse) xxx_ToOp(ctx context.Context, op *xxx_SetTableOperation) *xxx_SetTableOperation {
	if op == nil {
		op = &xxx_SetTableOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetTableResponse) xxx_FromOp(ctx context.Context, op *xxx_SetTableOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetTableResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetTableResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTableOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
