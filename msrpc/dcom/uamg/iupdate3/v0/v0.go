package iupdate3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iupdate2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdate2/v0"
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
	_ = iupdate2.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/uamg"
)

var (
	// IUpdate3 interface identifier 112eda6b-95b3-476f-9d90-aee82c6b8181
	Update3IID = &dcom.IID{Data1: 0x112eda6b, Data2: 0x95b3, Data3: 0x476f, Data4: []byte{0x9d, 0x90, 0xae, 0xe8, 0x2c, 0x6b, 0x81, 0x81}}
	// Syntax UUID
	Update3SyntaxUUID = &uuid.UUID{TimeLow: 0x112eda6b, TimeMid: 0x95b3, TimeHiAndVersion: 0x476f, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0x90, Node: [6]uint8{0xae, 0xe8, 0x2c, 0x6b, 0x81, 0x81}}
	// Syntax ID
	Update3SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: Update3SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IUpdate3 interface.
type Update3Client interface {

	// IUpdate2 retrieval method.
	Update2() iupdate2.Update2Client

	GetBrowseOnly(context.Context, *GetBrowseOnlyRequest, ...dcerpc.CallOption) (*GetBrowseOnlyResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Update3Client
}

type xxx_DefaultUpdate3Client struct {
	iupdate2.Update2Client
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultUpdate3Client) Update2() iupdate2.Update2Client {
	return o.Update2Client
}

func (o *xxx_DefaultUpdate3Client) GetBrowseOnly(ctx context.Context, in *GetBrowseOnlyRequest, opts ...dcerpc.CallOption) (*GetBrowseOnlyResponse, error) {
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
	out := &GetBrowseOnlyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdate3Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultUpdate3Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultUpdate3Client) IPID(ctx context.Context, ipid *dcom.IPID) Update3Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultUpdate3Client{
		Update2Client: o.Update2Client.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewUpdate3Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Update3Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Update3SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iupdate2.NewUpdate2Client(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultUpdate3Client{
		Update2Client: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetBrowseOnlyOperation structure represents the BrowseOnly operation
type xxx_GetBrowseOnlyOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue int16          `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetBrowseOnlyOperation) OpNum() int { return 56 }

func (o *xxx_GetBrowseOnlyOperation) OpName() string { return "/IUpdate3/v0/BrowseOnly" }

func (o *xxx_GetBrowseOnlyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBrowseOnlyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetBrowseOnlyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetBrowseOnlyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBrowseOnlyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetBrowseOnlyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// GetBrowseOnlyRequest structure represents the BrowseOnly operation request
type GetBrowseOnlyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetBrowseOnlyRequest) xxx_ToOp(ctx context.Context, op *xxx_GetBrowseOnlyOperation) *xxx_GetBrowseOnlyOperation {
	if op == nil {
		op = &xxx_GetBrowseOnlyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetBrowseOnlyRequest) xxx_FromOp(ctx context.Context, op *xxx_GetBrowseOnlyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetBrowseOnlyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetBrowseOnlyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBrowseOnlyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetBrowseOnlyResponse structure represents the BrowseOnly operation response
type GetBrowseOnlyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue int16          `idl:"name:retval" json:"return_value"`
	// Return: The BrowseOnly return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetBrowseOnlyResponse) xxx_ToOp(ctx context.Context, op *xxx_GetBrowseOnlyOperation) *xxx_GetBrowseOnlyOperation {
	if op == nil {
		op = &xxx_GetBrowseOnlyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetBrowseOnlyResponse) xxx_FromOp(ctx context.Context, op *xxx_GetBrowseOnlyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetBrowseOnlyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetBrowseOnlyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBrowseOnlyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
