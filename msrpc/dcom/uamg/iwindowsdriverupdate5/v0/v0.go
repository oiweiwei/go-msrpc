package iwindowsdriverupdate5

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	uamg "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg"
	iwindowsdriverupdate4 "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg/iwindowsdriverupdate4/v0"
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
	_ = iwindowsdriverupdate4.GoPackage
	_ = uamg.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/uamg"
)

var (
	// IWindowsDriverUpdate5 interface identifier 70cf5c82-8642-42bb-9dbc-0cfd263c6c4f
	WindowsDriverUpdate5IID = &dcom.IID{Data1: 0x70cf5c82, Data2: 0x8642, Data3: 0x42bb, Data4: []byte{0x9d, 0xbc, 0x0c, 0xfd, 0x26, 0x3c, 0x6c, 0x4f}}
	// Syntax UUID
	WindowsDriverUpdate5SyntaxUUID = &uuid.UUID{TimeLow: 0x70cf5c82, TimeMid: 0x8642, TimeHiAndVersion: 0x42bb, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0xbc, Node: [6]uint8{0xc, 0xfd, 0x26, 0x3c, 0x6c, 0x4f}}
	// Syntax ID
	WindowsDriverUpdate5SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: WindowsDriverUpdate5SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWindowsDriverUpdate5 interface.
type WindowsDriverUpdate5Client interface {

	// IWindowsDriverUpdate4 retrieval method.
	WindowsDriverUpdate4() iwindowsdriverupdate4.WindowsDriverUpdate4Client

	GetAutoSelection(context.Context, *GetAutoSelectionRequest, ...dcerpc.CallOption) (*GetAutoSelectionResponse, error)

	GetAutoDownload(context.Context, *GetAutoDownloadRequest, ...dcerpc.CallOption) (*GetAutoDownloadResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) WindowsDriverUpdate5Client
}

type xxx_DefaultWindowsDriverUpdate5Client struct {
	iwindowsdriverupdate4.WindowsDriverUpdate4Client
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultWindowsDriverUpdate5Client) WindowsDriverUpdate4() iwindowsdriverupdate4.WindowsDriverUpdate4Client {
	return o.WindowsDriverUpdate4Client
}

func (o *xxx_DefaultWindowsDriverUpdate5Client) GetAutoSelection(ctx context.Context, in *GetAutoSelectionRequest, opts ...dcerpc.CallOption) (*GetAutoSelectionResponse, error) {
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
	out := &GetAutoSelectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWindowsDriverUpdate5Client) GetAutoDownload(ctx context.Context, in *GetAutoDownloadRequest, opts ...dcerpc.CallOption) (*GetAutoDownloadResponse, error) {
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
	out := &GetAutoDownloadResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWindowsDriverUpdate5Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultWindowsDriverUpdate5Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultWindowsDriverUpdate5Client) IPID(ctx context.Context, ipid *dcom.IPID) WindowsDriverUpdate5Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultWindowsDriverUpdate5Client{
		WindowsDriverUpdate4Client: o.WindowsDriverUpdate4Client.IPID(ctx, ipid),
		cc:                         o.cc,
		ipid:                       ipid,
	}
}

func NewWindowsDriverUpdate5Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (WindowsDriverUpdate5Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(WindowsDriverUpdate5SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iwindowsdriverupdate4.NewWindowsDriverUpdate4Client(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultWindowsDriverUpdate5Client{
		WindowsDriverUpdate4Client: base,
		cc:                         cc,
		ipid:                       ipid,
	}, nil
}

// xxx_GetAutoSelectionOperation structure represents the AutoSelection operation
type xxx_GetAutoSelectionOperation struct {
	This        *dcom.ORPCThis         `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat         `idl:"name:That" json:"that"`
	ReturnValue uamg.AutoSelectionMode `idl:"name:retval" json:"return_value"`
	Return      int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAutoSelectionOperation) OpNum() int { return 67 }

func (o *xxx_GetAutoSelectionOperation) OpName() string {
	return "/IWindowsDriverUpdate5/v0/AutoSelection"
}

func (o *xxx_GetAutoSelectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAutoSelectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAutoSelectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAutoSelectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAutoSelectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{v1_enum, alias=AutoSelectionMode}(enum))
	{
		if err := w.WriteEnum(uint32(o.ReturnValue)); err != nil {
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

func (o *xxx_GetAutoSelectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{v1_enum, alias=AutoSelectionMode}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.ReturnValue)); err != nil {
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

// GetAutoSelectionRequest structure represents the AutoSelection operation request
type GetAutoSelectionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAutoSelectionRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAutoSelectionOperation) *xxx_GetAutoSelectionOperation {
	if op == nil {
		op = &xxx_GetAutoSelectionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetAutoSelectionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAutoSelectionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAutoSelectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAutoSelectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAutoSelectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAutoSelectionResponse structure represents the AutoSelection operation response
type GetAutoSelectionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat         `idl:"name:That" json:"that"`
	ReturnValue uamg.AutoSelectionMode `idl:"name:retval" json:"return_value"`
	// Return: The AutoSelection return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAutoSelectionResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAutoSelectionOperation) *xxx_GetAutoSelectionOperation {
	if op == nil {
		op = &xxx_GetAutoSelectionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetAutoSelectionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAutoSelectionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetAutoSelectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAutoSelectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAutoSelectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetAutoDownloadOperation structure represents the AutoDownload operation
type xxx_GetAutoDownloadOperation struct {
	This        *dcom.ORPCThis        `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat        `idl:"name:That" json:"that"`
	ReturnValue uamg.AutoDownloadMode `idl:"name:retval" json:"return_value"`
	Return      int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAutoDownloadOperation) OpNum() int { return 68 }

func (o *xxx_GetAutoDownloadOperation) OpName() string {
	return "/IWindowsDriverUpdate5/v0/AutoDownload"
}

func (o *xxx_GetAutoDownloadOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAutoDownloadOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAutoDownloadOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAutoDownloadOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAutoDownloadOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{v1_enum, alias=AutoDownloadMode}(enum))
	{
		if err := w.WriteEnum(uint32(o.ReturnValue)); err != nil {
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

func (o *xxx_GetAutoDownloadOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{v1_enum, alias=AutoDownloadMode}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.ReturnValue)); err != nil {
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

// GetAutoDownloadRequest structure represents the AutoDownload operation request
type GetAutoDownloadRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAutoDownloadRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAutoDownloadOperation) *xxx_GetAutoDownloadOperation {
	if op == nil {
		op = &xxx_GetAutoDownloadOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetAutoDownloadRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAutoDownloadOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAutoDownloadRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAutoDownloadRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAutoDownloadOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAutoDownloadResponse structure represents the AutoDownload operation response
type GetAutoDownloadResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat        `idl:"name:That" json:"that"`
	ReturnValue uamg.AutoDownloadMode `idl:"name:retval" json:"return_value"`
	// Return: The AutoDownload return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAutoDownloadResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAutoDownloadOperation) *xxx_GetAutoDownloadOperation {
	if op == nil {
		op = &xxx_GetAutoDownloadOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetAutoDownloadResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAutoDownloadOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetAutoDownloadResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAutoDownloadResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAutoDownloadOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
