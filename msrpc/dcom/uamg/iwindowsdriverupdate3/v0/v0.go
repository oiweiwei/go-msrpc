package iwindowsdriverupdate3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iwindowsdriverupdate2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iwindowsdriverupdate2/v0"
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
	_ = iwindowsdriverupdate2.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/uamg"
)

var (
	// IWindowsDriverUpdate3 interface identifier 49ebd502-4a96-41bd-9e3e-4c5057f4250c
	WindowsDriverUpdate3IID = &dcom.IID{Data1: 0x49ebd502, Data2: 0x4a96, Data3: 0x41bd, Data4: []byte{0x9e, 0x3e, 0x4c, 0x50, 0x57, 0xf4, 0x25, 0x0c}}
	// Syntax UUID
	WindowsDriverUpdate3SyntaxUUID = &uuid.UUID{TimeLow: 0x49ebd502, TimeMid: 0x4a96, TimeHiAndVersion: 0x41bd, ClockSeqHiAndReserved: 0x9e, ClockSeqLow: 0x3e, Node: [6]uint8{0x4c, 0x50, 0x57, 0xf4, 0x25, 0xc}}
	// Syntax ID
	WindowsDriverUpdate3SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: WindowsDriverUpdate3SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWindowsDriverUpdate3 interface.
type WindowsDriverUpdate3Client interface {

	// IWindowsDriverUpdate2 retrieval method.
	WindowsDriverUpdate2() iwindowsdriverupdate2.WindowsDriverUpdate2Client

	// The IUpdate3::BrowseOnly (opnum 57) method retrieves whether the update is browse-only.
	//
	// The IWindowsDriverUpdate3::BrowseOnly (opnum 65) method retrieves whether the update
	// is browse-only. If an update is browse-only, then it is not recommended for the update
	// to be installed automatically.
	//
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
	// This method SHOULD return the value of the BrowseOnly ADM element.
	GetBrowseOnly(context.Context, *GetBrowseOnlyRequest, ...dcerpc.CallOption) (*GetBrowseOnlyResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) WindowsDriverUpdate3Client
}

type xxx_DefaultWindowsDriverUpdate3Client struct {
	iwindowsdriverupdate2.WindowsDriverUpdate2Client
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultWindowsDriverUpdate3Client) WindowsDriverUpdate2() iwindowsdriverupdate2.WindowsDriverUpdate2Client {
	return o.WindowsDriverUpdate2Client
}

func (o *xxx_DefaultWindowsDriverUpdate3Client) GetBrowseOnly(ctx context.Context, in *GetBrowseOnlyRequest, opts ...dcerpc.CallOption) (*GetBrowseOnlyResponse, error) {
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

func (o *xxx_DefaultWindowsDriverUpdate3Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultWindowsDriverUpdate3Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultWindowsDriverUpdate3Client) IPID(ctx context.Context, ipid *dcom.IPID) WindowsDriverUpdate3Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultWindowsDriverUpdate3Client{
		WindowsDriverUpdate2Client: o.WindowsDriverUpdate2Client.IPID(ctx, ipid),
		cc:                         o.cc,
		ipid:                       ipid,
	}
}

func NewWindowsDriverUpdate3Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (WindowsDriverUpdate3Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(WindowsDriverUpdate3SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iwindowsdriverupdate2.NewWindowsDriverUpdate2Client(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultWindowsDriverUpdate3Client{
		WindowsDriverUpdate2Client: base,
		cc:                         cc,
		ipid:                       ipid,
	}, nil
}

// xxx_GetBrowseOnlyOperation structure represents the BrowseOnly operation
type xxx_GetBrowseOnlyOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue int16          `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetBrowseOnlyOperation) OpNum() int { return 64 }

func (o *xxx_GetBrowseOnlyOperation) OpName() string { return "/IWindowsDriverUpdate3/v0/BrowseOnly" }

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
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: MUST be set either to VARIANT_TRUE if this update is not recommended to be
	// installed automatically or to VARIANT_FALSE if not.
	//
	// retval: MUST be set either to VARIANT_TRUE if this update is browse-only, or to VARIANT_FALSE
	// otherwise.
	ReturnValue int16 `idl:"name:retval" json:"return_value"`
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
