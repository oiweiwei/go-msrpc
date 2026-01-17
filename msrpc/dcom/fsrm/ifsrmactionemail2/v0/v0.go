package ifsrmactionemail2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	ifsrmactionemail "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmactionemail/v0"
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
	_ = ifsrmactionemail.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmActionEmail2 interface identifier 8276702f-2532-4839-89bf-4872609a2ea4
	ActionEmail2IID = &dcom.IID{Data1: 0x8276702f, Data2: 0x2532, Data3: 0x4839, Data4: []byte{0x89, 0xbf, 0x48, 0x72, 0x60, 0x9a, 0x2e, 0xa4}}
	// Syntax UUID
	ActionEmail2SyntaxUUID = &uuid.UUID{TimeLow: 0x8276702f, TimeMid: 0x2532, TimeHiAndVersion: 0x4839, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0xbf, Node: [6]uint8{0x48, 0x72, 0x60, 0x9a, 0x2e, 0xa4}}
	// Syntax ID
	ActionEmail2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ActionEmail2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmActionEmail2 interface.
type ActionEmail2Client interface {

	// IFsrmActionEmail retrieval method.
	ActionEmail() ifsrmactionemail.ActionEmailClient

	// AttachmentFileListSize operation.
	GetAttachmentFileListSize(context.Context, *GetAttachmentFileListSizeRequest, ...dcerpc.CallOption) (*GetAttachmentFileListSizeResponse, error)

	// AttachmentFileListSize operation.
	SetAttachmentFileListSize(context.Context, *SetAttachmentFileListSizeRequest, ...dcerpc.CallOption) (*SetAttachmentFileListSizeResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ActionEmail2Client
}

type xxx_DefaultActionEmail2Client struct {
	ifsrmactionemail.ActionEmailClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultActionEmail2Client) ActionEmail() ifsrmactionemail.ActionEmailClient {
	return o.ActionEmailClient
}

func (o *xxx_DefaultActionEmail2Client) GetAttachmentFileListSize(ctx context.Context, in *GetAttachmentFileListSizeRequest, opts ...dcerpc.CallOption) (*GetAttachmentFileListSizeResponse, error) {
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
	out := &GetAttachmentFileListSizeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionEmail2Client) SetAttachmentFileListSize(ctx context.Context, in *SetAttachmentFileListSizeRequest, opts ...dcerpc.CallOption) (*SetAttachmentFileListSizeResponse, error) {
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
	out := &SetAttachmentFileListSizeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionEmail2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultActionEmail2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultActionEmail2Client) IPID(ctx context.Context, ipid *dcom.IPID) ActionEmail2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultActionEmail2Client{
		ActionEmailClient: o.ActionEmailClient.IPID(ctx, ipid),
		cc:                o.cc,
		ipid:              ipid,
	}
}

func NewActionEmail2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ActionEmail2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ActionEmail2SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ifsrmactionemail.NewActionEmailClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultActionEmail2Client{
		ActionEmailClient: base,
		cc:                cc,
		ipid:              ipid,
	}, nil
}

// xxx_GetAttachmentFileListSizeOperation structure represents the AttachmentFileListSize operation
type xxx_GetAttachmentFileListSizeOperation struct {
	This                   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                   *dcom.ORPCThat `idl:"name:That" json:"that"`
	AttachmentFileListSize int32          `idl:"name:attachmentFileListSize" json:"attachment_file_list_size"`
	Return                 int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAttachmentFileListSizeOperation) OpNum() int { return 26 }

func (o *xxx_GetAttachmentFileListSizeOperation) OpName() string {
	return "/IFsrmActionEmail2/v0/AttachmentFileListSize"
}

func (o *xxx_GetAttachmentFileListSizeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAttachmentFileListSizeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAttachmentFileListSizeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAttachmentFileListSizeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAttachmentFileListSizeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// attachmentFileListSize {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.AttachmentFileListSize); err != nil {
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

func (o *xxx_GetAttachmentFileListSizeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// attachmentFileListSize {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.AttachmentFileListSize); err != nil {
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

// GetAttachmentFileListSizeRequest structure represents the AttachmentFileListSize operation request
type GetAttachmentFileListSizeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAttachmentFileListSizeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAttachmentFileListSizeOperation) *xxx_GetAttachmentFileListSizeOperation {
	if op == nil {
		op = &xxx_GetAttachmentFileListSizeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetAttachmentFileListSizeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAttachmentFileListSizeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAttachmentFileListSizeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAttachmentFileListSizeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAttachmentFileListSizeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAttachmentFileListSizeResponse structure represents the AttachmentFileListSize operation response
type GetAttachmentFileListSizeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                   *dcom.ORPCThat `idl:"name:That" json:"that"`
	AttachmentFileListSize int32          `idl:"name:attachmentFileListSize" json:"attachment_file_list_size"`
	// Return: The AttachmentFileListSize return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAttachmentFileListSizeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAttachmentFileListSizeOperation) *xxx_GetAttachmentFileListSizeOperation {
	if op == nil {
		op = &xxx_GetAttachmentFileListSizeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.AttachmentFileListSize = o.AttachmentFileListSize
	op.Return = o.Return
	return op
}

func (o *GetAttachmentFileListSizeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAttachmentFileListSizeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.AttachmentFileListSize = op.AttachmentFileListSize
	o.Return = op.Return
}
func (o *GetAttachmentFileListSizeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAttachmentFileListSizeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAttachmentFileListSizeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetAttachmentFileListSizeOperation structure represents the AttachmentFileListSize operation
type xxx_SetAttachmentFileListSizeOperation struct {
	This                   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                   *dcom.ORPCThat `idl:"name:That" json:"that"`
	AttachmentFileListSize int32          `idl:"name:attachmentFileListSize" json:"attachment_file_list_size"`
	Return                 int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetAttachmentFileListSizeOperation) OpNum() int { return 27 }

func (o *xxx_SetAttachmentFileListSizeOperation) OpName() string {
	return "/IFsrmActionEmail2/v0/AttachmentFileListSize"
}

func (o *xxx_SetAttachmentFileListSizeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAttachmentFileListSizeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// attachmentFileListSize {in} (1:(int32))
	{
		if err := w.WriteData(o.AttachmentFileListSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAttachmentFileListSizeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// attachmentFileListSize {in} (1:(int32))
	{
		if err := w.ReadData(&o.AttachmentFileListSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAttachmentFileListSizeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAttachmentFileListSizeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetAttachmentFileListSizeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetAttachmentFileListSizeRequest structure represents the AttachmentFileListSize operation request
type SetAttachmentFileListSizeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                   *dcom.ORPCThis `idl:"name:This" json:"this"`
	AttachmentFileListSize int32          `idl:"name:attachmentFileListSize" json:"attachment_file_list_size"`
}

func (o *SetAttachmentFileListSizeRequest) xxx_ToOp(ctx context.Context, op *xxx_SetAttachmentFileListSizeOperation) *xxx_SetAttachmentFileListSizeOperation {
	if op == nil {
		op = &xxx_SetAttachmentFileListSizeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.AttachmentFileListSize = o.AttachmentFileListSize
	return op
}

func (o *SetAttachmentFileListSizeRequest) xxx_FromOp(ctx context.Context, op *xxx_SetAttachmentFileListSizeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.AttachmentFileListSize = op.AttachmentFileListSize
}
func (o *SetAttachmentFileListSizeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetAttachmentFileListSizeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAttachmentFileListSizeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetAttachmentFileListSizeResponse structure represents the AttachmentFileListSize operation response
type SetAttachmentFileListSizeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AttachmentFileListSize return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetAttachmentFileListSizeResponse) xxx_ToOp(ctx context.Context, op *xxx_SetAttachmentFileListSizeOperation) *xxx_SetAttachmentFileListSizeOperation {
	if op == nil {
		op = &xxx_SetAttachmentFileListSizeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetAttachmentFileListSizeResponse) xxx_FromOp(ctx context.Context, op *xxx_SetAttachmentFileListSizeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetAttachmentFileListSizeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetAttachmentFileListSizeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAttachmentFileListSizeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
