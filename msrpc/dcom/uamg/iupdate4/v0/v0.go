package iupdate4

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iupdate3 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iupdate3/v0"
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
	_ = iupdate3.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/uamg"
)

var (
	// IUpdate4 interface identifier 27e94b0d-5139-49a2-9a61-93522dc54652
	Update4IID = &dcom.IID{Data1: 0x27e94b0d, Data2: 0x5139, Data3: 0x49a2, Data4: []byte{0x9a, 0x61, 0x93, 0x52, 0x2d, 0xc5, 0x46, 0x52}}
	// Syntax UUID
	Update4SyntaxUUID = &uuid.UUID{TimeLow: 0x27e94b0d, TimeMid: 0x5139, TimeHiAndVersion: 0x49a2, ClockSeqHiAndReserved: 0x9a, ClockSeqLow: 0x61, Node: [6]uint8{0x93, 0x52, 0x2d, 0xc5, 0x46, 0x52}}
	// Syntax ID
	Update4SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: Update4SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IUpdate4 interface.
type Update4Client interface {

	// IUpdate3 retrieval method.
	Update3() iupdate3.Update3Client

	GetPerUser(context.Context, *GetPerUserRequest, ...dcerpc.CallOption) (*GetPerUserResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Update4Client
}

type xxx_DefaultUpdate4Client struct {
	iupdate3.Update3Client
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultUpdate4Client) Update3() iupdate3.Update3Client {
	return o.Update3Client
}

func (o *xxx_DefaultUpdate4Client) GetPerUser(ctx context.Context, in *GetPerUserRequest, opts ...dcerpc.CallOption) (*GetPerUserResponse, error) {
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
	out := &GetPerUserResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdate4Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultUpdate4Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultUpdate4Client) IPID(ctx context.Context, ipid *dcom.IPID) Update4Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultUpdate4Client{
		Update3Client: o.Update3Client.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewUpdate4Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Update4Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Update4SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iupdate3.NewUpdate3Client(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultUpdate4Client{
		Update3Client: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetPerUserOperation structure represents the PerUser operation
type xxx_GetPerUserOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue int16          `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPerUserOperation) OpNum() int { return 57 }

func (o *xxx_GetPerUserOperation) OpName() string { return "/IUpdate4/v0/PerUser" }

func (o *xxx_GetPerUserOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPerUserOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPerUserOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPerUserOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPerUserOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPerUserOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// GetPerUserRequest structure represents the PerUser operation request
type GetPerUserRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPerUserRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPerUserOperation) *xxx_GetPerUserOperation {
	if op == nil {
		op = &xxx_GetPerUserOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPerUserRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPerUserOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPerUserRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPerUserRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPerUserOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPerUserResponse structure represents the PerUser operation response
type GetPerUserResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue int16          `idl:"name:retval" json:"return_value"`
	// Return: The PerUser return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPerUserResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPerUserOperation) *xxx_GetPerUserOperation {
	if op == nil {
		op = &xxx_GetPerUserOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetPerUserResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPerUserOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetPerUserResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPerUserResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPerUserOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
