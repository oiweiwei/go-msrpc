package ifsrmactionemail

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	ifsrmaction "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmaction/v0"
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
	_ = ifsrmaction.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmActionEmail interface identifier d646567d-26ae-4caa-9f84-4e0aad207fca
	ActionEmailIID = &dcom.IID{Data1: 0xd646567d, Data2: 0x26ae, Data3: 0x4caa, Data4: []byte{0x9f, 0x84, 0x4e, 0x0a, 0xad, 0x20, 0x7f, 0xca}}
	// Syntax UUID
	ActionEmailSyntaxUUID = &uuid.UUID{TimeLow: 0xd646567d, TimeMid: 0x26ae, TimeHiAndVersion: 0x4caa, ClockSeqHiAndReserved: 0x9f, ClockSeqLow: 0x84, Node: [6]uint8{0x4e, 0xa, 0xad, 0x20, 0x7f, 0xca}}
	// Syntax ID
	ActionEmailSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ActionEmailSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmActionEmail interface.
type ActionEmailClient interface {

	// IFsrmAction retrieval method.
	Action() ifsrmaction.ActionClient

	// MailFrom operation.
	GetMailFrom(context.Context, *GetMailFromRequest, ...dcerpc.CallOption) (*GetMailFromResponse, error)

	// MailFrom operation.
	SetMailFrom(context.Context, *SetMailFromRequest, ...dcerpc.CallOption) (*SetMailFromResponse, error)

	// MailReplyTo operation.
	GetMailReplyTo(context.Context, *GetMailReplyToRequest, ...dcerpc.CallOption) (*GetMailReplyToResponse, error)

	// MailReplyTo operation.
	SetMailReplyTo(context.Context, *SetMailReplyToRequest, ...dcerpc.CallOption) (*SetMailReplyToResponse, error)

	// MailTo operation.
	GetMailTo(context.Context, *GetMailToRequest, ...dcerpc.CallOption) (*GetMailToResponse, error)

	// MailTo operation.
	SetMailTo(context.Context, *SetMailToRequest, ...dcerpc.CallOption) (*SetMailToResponse, error)

	// MailCc operation.
	GetMailCC(context.Context, *GetMailCCRequest, ...dcerpc.CallOption) (*GetMailCCResponse, error)

	// MailCc operation.
	SetMailCC(context.Context, *SetMailCCRequest, ...dcerpc.CallOption) (*SetMailCCResponse, error)

	// MailBcc operation.
	GetMailBCC(context.Context, *GetMailBCCRequest, ...dcerpc.CallOption) (*GetMailBCCResponse, error)

	// MailBcc operation.
	SetMailBCC(context.Context, *SetMailBCCRequest, ...dcerpc.CallOption) (*SetMailBCCResponse, error)

	// MailSubject operation.
	GetMailSubject(context.Context, *GetMailSubjectRequest, ...dcerpc.CallOption) (*GetMailSubjectResponse, error)

	// MailSubject operation.
	SetMailSubject(context.Context, *SetMailSubjectRequest, ...dcerpc.CallOption) (*SetMailSubjectResponse, error)

	// MessageText operation.
	GetMessageText(context.Context, *GetMessageTextRequest, ...dcerpc.CallOption) (*GetMessageTextResponse, error)

	// MessageText operation.
	SetMessageText(context.Context, *SetMessageTextRequest, ...dcerpc.CallOption) (*SetMessageTextResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ActionEmailClient
}

type xxx_DefaultActionEmailClient struct {
	ifsrmaction.ActionClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultActionEmailClient) Action() ifsrmaction.ActionClient {
	return o.ActionClient
}

func (o *xxx_DefaultActionEmailClient) GetMailFrom(ctx context.Context, in *GetMailFromRequest, opts ...dcerpc.CallOption) (*GetMailFromResponse, error) {
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
	out := &GetMailFromResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionEmailClient) SetMailFrom(ctx context.Context, in *SetMailFromRequest, opts ...dcerpc.CallOption) (*SetMailFromResponse, error) {
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
	out := &SetMailFromResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionEmailClient) GetMailReplyTo(ctx context.Context, in *GetMailReplyToRequest, opts ...dcerpc.CallOption) (*GetMailReplyToResponse, error) {
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
	out := &GetMailReplyToResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionEmailClient) SetMailReplyTo(ctx context.Context, in *SetMailReplyToRequest, opts ...dcerpc.CallOption) (*SetMailReplyToResponse, error) {
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
	out := &SetMailReplyToResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionEmailClient) GetMailTo(ctx context.Context, in *GetMailToRequest, opts ...dcerpc.CallOption) (*GetMailToResponse, error) {
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
	out := &GetMailToResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionEmailClient) SetMailTo(ctx context.Context, in *SetMailToRequest, opts ...dcerpc.CallOption) (*SetMailToResponse, error) {
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
	out := &SetMailToResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionEmailClient) GetMailCC(ctx context.Context, in *GetMailCCRequest, opts ...dcerpc.CallOption) (*GetMailCCResponse, error) {
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
	out := &GetMailCCResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionEmailClient) SetMailCC(ctx context.Context, in *SetMailCCRequest, opts ...dcerpc.CallOption) (*SetMailCCResponse, error) {
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
	out := &SetMailCCResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionEmailClient) GetMailBCC(ctx context.Context, in *GetMailBCCRequest, opts ...dcerpc.CallOption) (*GetMailBCCResponse, error) {
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
	out := &GetMailBCCResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionEmailClient) SetMailBCC(ctx context.Context, in *SetMailBCCRequest, opts ...dcerpc.CallOption) (*SetMailBCCResponse, error) {
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
	out := &SetMailBCCResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionEmailClient) GetMailSubject(ctx context.Context, in *GetMailSubjectRequest, opts ...dcerpc.CallOption) (*GetMailSubjectResponse, error) {
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
	out := &GetMailSubjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionEmailClient) SetMailSubject(ctx context.Context, in *SetMailSubjectRequest, opts ...dcerpc.CallOption) (*SetMailSubjectResponse, error) {
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
	out := &SetMailSubjectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionEmailClient) GetMessageText(ctx context.Context, in *GetMessageTextRequest, opts ...dcerpc.CallOption) (*GetMessageTextResponse, error) {
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
	out := &GetMessageTextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionEmailClient) SetMessageText(ctx context.Context, in *SetMessageTextRequest, opts ...dcerpc.CallOption) (*SetMessageTextResponse, error) {
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
	out := &SetMessageTextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionEmailClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultActionEmailClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultActionEmailClient) IPID(ctx context.Context, ipid *dcom.IPID) ActionEmailClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultActionEmailClient{
		ActionClient: o.ActionClient.IPID(ctx, ipid),
		cc:           o.cc,
		ipid:         ipid,
	}
}

func NewActionEmailClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ActionEmailClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ActionEmailSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ifsrmaction.NewActionClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultActionEmailClient{
		ActionClient: base,
		cc:           cc,
		ipid:         ipid,
	}, nil
}

// xxx_GetMailFromOperation structure represents the MailFrom operation
type xxx_GetMailFromOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailFrom *oaut.String   `idl:"name:mailFrom" json:"mail_from"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMailFromOperation) OpNum() int { return 12 }

func (o *xxx_GetMailFromOperation) OpName() string { return "/IFsrmActionEmail/v0/MailFrom" }

func (o *xxx_GetMailFromOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMailFromOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMailFromOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMailFromOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMailFromOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// mailFrom {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MailFrom != nil {
			_ptr_mailFrom := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MailFrom != nil {
					if err := o.MailFrom.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MailFrom, _ptr_mailFrom); err != nil {
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

func (o *xxx_GetMailFromOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// mailFrom {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_mailFrom := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MailFrom == nil {
				o.MailFrom = &oaut.String{}
			}
			if err := o.MailFrom.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_mailFrom := func(ptr interface{}) { o.MailFrom = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MailFrom, _s_mailFrom, _ptr_mailFrom); err != nil {
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

// GetMailFromRequest structure represents the MailFrom operation request
type GetMailFromRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMailFromRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMailFromOperation) *xxx_GetMailFromOperation {
	if op == nil {
		op = &xxx_GetMailFromOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMailFromRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMailFromOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMailFromRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMailFromRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMailFromOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMailFromResponse structure represents the MailFrom operation response
type GetMailFromResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailFrom *oaut.String   `idl:"name:mailFrom" json:"mail_from"`
	// Return: The MailFrom return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMailFromResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMailFromOperation) *xxx_GetMailFromOperation {
	if op == nil {
		op = &xxx_GetMailFromOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MailFrom = o.MailFrom
	op.Return = o.Return
	return op
}

func (o *GetMailFromResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMailFromOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MailFrom = op.MailFrom
	o.Return = op.Return
}
func (o *GetMailFromResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMailFromResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMailFromOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMailFromOperation structure represents the MailFrom operation
type xxx_SetMailFromOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailFrom *oaut.String   `idl:"name:mailFrom" json:"mail_from"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMailFromOperation) OpNum() int { return 13 }

func (o *xxx_SetMailFromOperation) OpName() string { return "/IFsrmActionEmail/v0/MailFrom" }

func (o *xxx_SetMailFromOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailFromOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// mailFrom {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MailFrom != nil {
			_ptr_mailFrom := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MailFrom != nil {
					if err := o.MailFrom.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MailFrom, _ptr_mailFrom); err != nil {
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
	return nil
}

func (o *xxx_SetMailFromOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// mailFrom {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_mailFrom := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MailFrom == nil {
				o.MailFrom = &oaut.String{}
			}
			if err := o.MailFrom.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_mailFrom := func(ptr interface{}) { o.MailFrom = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MailFrom, _s_mailFrom, _ptr_mailFrom); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailFromOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailFromOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMailFromOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMailFromRequest structure represents the MailFrom operation request
type SetMailFromRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	MailFrom *oaut.String   `idl:"name:mailFrom" json:"mail_from"`
}

func (o *SetMailFromRequest) xxx_ToOp(ctx context.Context, op *xxx_SetMailFromOperation) *xxx_SetMailFromOperation {
	if op == nil {
		op = &xxx_SetMailFromOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MailFrom = o.MailFrom
	return op
}

func (o *SetMailFromRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMailFromOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MailFrom = op.MailFrom
}
func (o *SetMailFromRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetMailFromRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMailFromOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMailFromResponse structure represents the MailFrom operation response
type SetMailFromResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MailFrom return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMailFromResponse) xxx_ToOp(ctx context.Context, op *xxx_SetMailFromOperation) *xxx_SetMailFromOperation {
	if op == nil {
		op = &xxx_SetMailFromOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetMailFromResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMailFromOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMailFromResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetMailFromResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMailFromOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMailReplyToOperation structure represents the MailReplyTo operation
type xxx_GetMailReplyToOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailReplyTo *oaut.String   `idl:"name:mailReplyTo" json:"mail_reply_to"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMailReplyToOperation) OpNum() int { return 14 }

func (o *xxx_GetMailReplyToOperation) OpName() string { return "/IFsrmActionEmail/v0/MailReplyTo" }

func (o *xxx_GetMailReplyToOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMailReplyToOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMailReplyToOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMailReplyToOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMailReplyToOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// mailReplyTo {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MailReplyTo != nil {
			_ptr_mailReplyTo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MailReplyTo != nil {
					if err := o.MailReplyTo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MailReplyTo, _ptr_mailReplyTo); err != nil {
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

func (o *xxx_GetMailReplyToOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// mailReplyTo {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_mailReplyTo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MailReplyTo == nil {
				o.MailReplyTo = &oaut.String{}
			}
			if err := o.MailReplyTo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_mailReplyTo := func(ptr interface{}) { o.MailReplyTo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MailReplyTo, _s_mailReplyTo, _ptr_mailReplyTo); err != nil {
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

// GetMailReplyToRequest structure represents the MailReplyTo operation request
type GetMailReplyToRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMailReplyToRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMailReplyToOperation) *xxx_GetMailReplyToOperation {
	if op == nil {
		op = &xxx_GetMailReplyToOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMailReplyToRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMailReplyToOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMailReplyToRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMailReplyToRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMailReplyToOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMailReplyToResponse structure represents the MailReplyTo operation response
type GetMailReplyToResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailReplyTo *oaut.String   `idl:"name:mailReplyTo" json:"mail_reply_to"`
	// Return: The MailReplyTo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMailReplyToResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMailReplyToOperation) *xxx_GetMailReplyToOperation {
	if op == nil {
		op = &xxx_GetMailReplyToOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MailReplyTo = o.MailReplyTo
	op.Return = o.Return
	return op
}

func (o *GetMailReplyToResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMailReplyToOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MailReplyTo = op.MailReplyTo
	o.Return = op.Return
}
func (o *GetMailReplyToResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMailReplyToResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMailReplyToOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMailReplyToOperation structure represents the MailReplyTo operation
type xxx_SetMailReplyToOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailReplyTo *oaut.String   `idl:"name:mailReplyTo" json:"mail_reply_to"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMailReplyToOperation) OpNum() int { return 15 }

func (o *xxx_SetMailReplyToOperation) OpName() string { return "/IFsrmActionEmail/v0/MailReplyTo" }

func (o *xxx_SetMailReplyToOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailReplyToOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// mailReplyTo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MailReplyTo != nil {
			_ptr_mailReplyTo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MailReplyTo != nil {
					if err := o.MailReplyTo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MailReplyTo, _ptr_mailReplyTo); err != nil {
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
	return nil
}

func (o *xxx_SetMailReplyToOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// mailReplyTo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_mailReplyTo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MailReplyTo == nil {
				o.MailReplyTo = &oaut.String{}
			}
			if err := o.MailReplyTo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_mailReplyTo := func(ptr interface{}) { o.MailReplyTo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MailReplyTo, _s_mailReplyTo, _ptr_mailReplyTo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailReplyToOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailReplyToOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMailReplyToOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMailReplyToRequest structure represents the MailReplyTo operation request
type SetMailReplyToRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	MailReplyTo *oaut.String   `idl:"name:mailReplyTo" json:"mail_reply_to"`
}

func (o *SetMailReplyToRequest) xxx_ToOp(ctx context.Context, op *xxx_SetMailReplyToOperation) *xxx_SetMailReplyToOperation {
	if op == nil {
		op = &xxx_SetMailReplyToOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MailReplyTo = o.MailReplyTo
	return op
}

func (o *SetMailReplyToRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMailReplyToOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MailReplyTo = op.MailReplyTo
}
func (o *SetMailReplyToRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetMailReplyToRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMailReplyToOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMailReplyToResponse structure represents the MailReplyTo operation response
type SetMailReplyToResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MailReplyTo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMailReplyToResponse) xxx_ToOp(ctx context.Context, op *xxx_SetMailReplyToOperation) *xxx_SetMailReplyToOperation {
	if op == nil {
		op = &xxx_SetMailReplyToOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetMailReplyToResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMailReplyToOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMailReplyToResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetMailReplyToResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMailReplyToOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMailToOperation structure represents the MailTo operation
type xxx_GetMailToOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailTo *oaut.String   `idl:"name:mailTo" json:"mail_to"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMailToOperation) OpNum() int { return 16 }

func (o *xxx_GetMailToOperation) OpName() string { return "/IFsrmActionEmail/v0/MailTo" }

func (o *xxx_GetMailToOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMailToOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMailToOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMailToOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMailToOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// mailTo {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MailTo != nil {
			_ptr_mailTo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MailTo != nil {
					if err := o.MailTo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MailTo, _ptr_mailTo); err != nil {
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

func (o *xxx_GetMailToOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// mailTo {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_mailTo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MailTo == nil {
				o.MailTo = &oaut.String{}
			}
			if err := o.MailTo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_mailTo := func(ptr interface{}) { o.MailTo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MailTo, _s_mailTo, _ptr_mailTo); err != nil {
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

// GetMailToRequest structure represents the MailTo operation request
type GetMailToRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMailToRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMailToOperation) *xxx_GetMailToOperation {
	if op == nil {
		op = &xxx_GetMailToOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMailToRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMailToOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMailToRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMailToRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMailToOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMailToResponse structure represents the MailTo operation response
type GetMailToResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailTo *oaut.String   `idl:"name:mailTo" json:"mail_to"`
	// Return: The MailTo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMailToResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMailToOperation) *xxx_GetMailToOperation {
	if op == nil {
		op = &xxx_GetMailToOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MailTo = o.MailTo
	op.Return = o.Return
	return op
}

func (o *GetMailToResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMailToOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MailTo = op.MailTo
	o.Return = op.Return
}
func (o *GetMailToResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMailToResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMailToOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMailToOperation structure represents the MailTo operation
type xxx_SetMailToOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailTo *oaut.String   `idl:"name:mailTo" json:"mail_to"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMailToOperation) OpNum() int { return 17 }

func (o *xxx_SetMailToOperation) OpName() string { return "/IFsrmActionEmail/v0/MailTo" }

func (o *xxx_SetMailToOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailToOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// mailTo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MailTo != nil {
			_ptr_mailTo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MailTo != nil {
					if err := o.MailTo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MailTo, _ptr_mailTo); err != nil {
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
	return nil
}

func (o *xxx_SetMailToOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// mailTo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_mailTo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MailTo == nil {
				o.MailTo = &oaut.String{}
			}
			if err := o.MailTo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_mailTo := func(ptr interface{}) { o.MailTo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MailTo, _s_mailTo, _ptr_mailTo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailToOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailToOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMailToOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMailToRequest structure represents the MailTo operation request
type SetMailToRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	MailTo *oaut.String   `idl:"name:mailTo" json:"mail_to"`
}

func (o *SetMailToRequest) xxx_ToOp(ctx context.Context, op *xxx_SetMailToOperation) *xxx_SetMailToOperation {
	if op == nil {
		op = &xxx_SetMailToOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MailTo = o.MailTo
	return op
}

func (o *SetMailToRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMailToOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MailTo = op.MailTo
}
func (o *SetMailToRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetMailToRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMailToOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMailToResponse structure represents the MailTo operation response
type SetMailToResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MailTo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMailToResponse) xxx_ToOp(ctx context.Context, op *xxx_SetMailToOperation) *xxx_SetMailToOperation {
	if op == nil {
		op = &xxx_SetMailToOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetMailToResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMailToOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMailToResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetMailToResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMailToOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMailCCOperation structure represents the MailCc operation
type xxx_GetMailCCOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailCC *oaut.String   `idl:"name:mailCc" json:"mail_cc"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMailCCOperation) OpNum() int { return 18 }

func (o *xxx_GetMailCCOperation) OpName() string { return "/IFsrmActionEmail/v0/MailCc" }

func (o *xxx_GetMailCCOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMailCCOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMailCCOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMailCCOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMailCCOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// mailCc {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MailCC != nil {
			_ptr_mailCc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MailCC != nil {
					if err := o.MailCC.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MailCC, _ptr_mailCc); err != nil {
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

func (o *xxx_GetMailCCOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// mailCc {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_mailCc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MailCC == nil {
				o.MailCC = &oaut.String{}
			}
			if err := o.MailCC.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_mailCc := func(ptr interface{}) { o.MailCC = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MailCC, _s_mailCc, _ptr_mailCc); err != nil {
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

// GetMailCCRequest structure represents the MailCc operation request
type GetMailCCRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMailCCRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMailCCOperation) *xxx_GetMailCCOperation {
	if op == nil {
		op = &xxx_GetMailCCOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMailCCRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMailCCOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMailCCRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMailCCRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMailCCOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMailCCResponse structure represents the MailCc operation response
type GetMailCCResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailCC *oaut.String   `idl:"name:mailCc" json:"mail_cc"`
	// Return: The MailCc return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMailCCResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMailCCOperation) *xxx_GetMailCCOperation {
	if op == nil {
		op = &xxx_GetMailCCOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MailCC = o.MailCC
	op.Return = o.Return
	return op
}

func (o *GetMailCCResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMailCCOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MailCC = op.MailCC
	o.Return = op.Return
}
func (o *GetMailCCResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMailCCResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMailCCOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMailCCOperation structure represents the MailCc operation
type xxx_SetMailCCOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailCC *oaut.String   `idl:"name:mailCc" json:"mail_cc"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMailCCOperation) OpNum() int { return 19 }

func (o *xxx_SetMailCCOperation) OpName() string { return "/IFsrmActionEmail/v0/MailCc" }

func (o *xxx_SetMailCCOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailCCOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// mailCc {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MailCC != nil {
			_ptr_mailCc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MailCC != nil {
					if err := o.MailCC.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MailCC, _ptr_mailCc); err != nil {
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
	return nil
}

func (o *xxx_SetMailCCOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// mailCc {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_mailCc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MailCC == nil {
				o.MailCC = &oaut.String{}
			}
			if err := o.MailCC.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_mailCc := func(ptr interface{}) { o.MailCC = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MailCC, _s_mailCc, _ptr_mailCc); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailCCOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailCCOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMailCCOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMailCCRequest structure represents the MailCc operation request
type SetMailCCRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	MailCC *oaut.String   `idl:"name:mailCc" json:"mail_cc"`
}

func (o *SetMailCCRequest) xxx_ToOp(ctx context.Context, op *xxx_SetMailCCOperation) *xxx_SetMailCCOperation {
	if op == nil {
		op = &xxx_SetMailCCOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MailCC = o.MailCC
	return op
}

func (o *SetMailCCRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMailCCOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MailCC = op.MailCC
}
func (o *SetMailCCRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetMailCCRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMailCCOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMailCCResponse structure represents the MailCc operation response
type SetMailCCResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MailCc return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMailCCResponse) xxx_ToOp(ctx context.Context, op *xxx_SetMailCCOperation) *xxx_SetMailCCOperation {
	if op == nil {
		op = &xxx_SetMailCCOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetMailCCResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMailCCOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMailCCResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetMailCCResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMailCCOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMailBCCOperation structure represents the MailBcc operation
type xxx_GetMailBCCOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailBCC *oaut.String   `idl:"name:mailBcc" json:"mail_bcc"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMailBCCOperation) OpNum() int { return 20 }

func (o *xxx_GetMailBCCOperation) OpName() string { return "/IFsrmActionEmail/v0/MailBcc" }

func (o *xxx_GetMailBCCOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMailBCCOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMailBCCOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMailBCCOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMailBCCOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// mailBcc {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MailBCC != nil {
			_ptr_mailBcc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MailBCC != nil {
					if err := o.MailBCC.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MailBCC, _ptr_mailBcc); err != nil {
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

func (o *xxx_GetMailBCCOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// mailBcc {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_mailBcc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MailBCC == nil {
				o.MailBCC = &oaut.String{}
			}
			if err := o.MailBCC.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_mailBcc := func(ptr interface{}) { o.MailBCC = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MailBCC, _s_mailBcc, _ptr_mailBcc); err != nil {
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

// GetMailBCCRequest structure represents the MailBcc operation request
type GetMailBCCRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMailBCCRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMailBCCOperation) *xxx_GetMailBCCOperation {
	if op == nil {
		op = &xxx_GetMailBCCOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMailBCCRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMailBCCOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMailBCCRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMailBCCRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMailBCCOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMailBCCResponse structure represents the MailBcc operation response
type GetMailBCCResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailBCC *oaut.String   `idl:"name:mailBcc" json:"mail_bcc"`
	// Return: The MailBcc return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMailBCCResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMailBCCOperation) *xxx_GetMailBCCOperation {
	if op == nil {
		op = &xxx_GetMailBCCOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MailBCC = o.MailBCC
	op.Return = o.Return
	return op
}

func (o *GetMailBCCResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMailBCCOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MailBCC = op.MailBCC
	o.Return = op.Return
}
func (o *GetMailBCCResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMailBCCResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMailBCCOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMailBCCOperation structure represents the MailBcc operation
type xxx_SetMailBCCOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailBCC *oaut.String   `idl:"name:mailBcc" json:"mail_bcc"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMailBCCOperation) OpNum() int { return 21 }

func (o *xxx_SetMailBCCOperation) OpName() string { return "/IFsrmActionEmail/v0/MailBcc" }

func (o *xxx_SetMailBCCOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailBCCOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// mailBcc {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MailBCC != nil {
			_ptr_mailBcc := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MailBCC != nil {
					if err := o.MailBCC.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MailBCC, _ptr_mailBcc); err != nil {
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
	return nil
}

func (o *xxx_SetMailBCCOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// mailBcc {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_mailBcc := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MailBCC == nil {
				o.MailBCC = &oaut.String{}
			}
			if err := o.MailBCC.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_mailBcc := func(ptr interface{}) { o.MailBCC = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MailBCC, _s_mailBcc, _ptr_mailBcc); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailBCCOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailBCCOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMailBCCOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMailBCCRequest structure represents the MailBcc operation request
type SetMailBCCRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	MailBCC *oaut.String   `idl:"name:mailBcc" json:"mail_bcc"`
}

func (o *SetMailBCCRequest) xxx_ToOp(ctx context.Context, op *xxx_SetMailBCCOperation) *xxx_SetMailBCCOperation {
	if op == nil {
		op = &xxx_SetMailBCCOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MailBCC = o.MailBCC
	return op
}

func (o *SetMailBCCRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMailBCCOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MailBCC = op.MailBCC
}
func (o *SetMailBCCRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetMailBCCRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMailBCCOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMailBCCResponse structure represents the MailBcc operation response
type SetMailBCCResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MailBcc return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMailBCCResponse) xxx_ToOp(ctx context.Context, op *xxx_SetMailBCCOperation) *xxx_SetMailBCCOperation {
	if op == nil {
		op = &xxx_SetMailBCCOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetMailBCCResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMailBCCOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMailBCCResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetMailBCCResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMailBCCOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMailSubjectOperation structure represents the MailSubject operation
type xxx_GetMailSubjectOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailSubject *oaut.String   `idl:"name:mailSubject" json:"mail_subject"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMailSubjectOperation) OpNum() int { return 22 }

func (o *xxx_GetMailSubjectOperation) OpName() string { return "/IFsrmActionEmail/v0/MailSubject" }

func (o *xxx_GetMailSubjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMailSubjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMailSubjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMailSubjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMailSubjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// mailSubject {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MailSubject != nil {
			_ptr_mailSubject := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MailSubject != nil {
					if err := o.MailSubject.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MailSubject, _ptr_mailSubject); err != nil {
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

func (o *xxx_GetMailSubjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// mailSubject {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_mailSubject := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MailSubject == nil {
				o.MailSubject = &oaut.String{}
			}
			if err := o.MailSubject.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_mailSubject := func(ptr interface{}) { o.MailSubject = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MailSubject, _s_mailSubject, _ptr_mailSubject); err != nil {
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

// GetMailSubjectRequest structure represents the MailSubject operation request
type GetMailSubjectRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMailSubjectRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMailSubjectOperation) *xxx_GetMailSubjectOperation {
	if op == nil {
		op = &xxx_GetMailSubjectOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMailSubjectRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMailSubjectOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMailSubjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMailSubjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMailSubjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMailSubjectResponse structure represents the MailSubject operation response
type GetMailSubjectResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailSubject *oaut.String   `idl:"name:mailSubject" json:"mail_subject"`
	// Return: The MailSubject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMailSubjectResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMailSubjectOperation) *xxx_GetMailSubjectOperation {
	if op == nil {
		op = &xxx_GetMailSubjectOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MailSubject = o.MailSubject
	op.Return = o.Return
	return op
}

func (o *GetMailSubjectResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMailSubjectOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MailSubject = op.MailSubject
	o.Return = op.Return
}
func (o *GetMailSubjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMailSubjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMailSubjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMailSubjectOperation structure represents the MailSubject operation
type xxx_SetMailSubjectOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MailSubject *oaut.String   `idl:"name:mailSubject" json:"mail_subject"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMailSubjectOperation) OpNum() int { return 23 }

func (o *xxx_SetMailSubjectOperation) OpName() string { return "/IFsrmActionEmail/v0/MailSubject" }

func (o *xxx_SetMailSubjectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailSubjectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// mailSubject {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MailSubject != nil {
			_ptr_mailSubject := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MailSubject != nil {
					if err := o.MailSubject.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MailSubject, _ptr_mailSubject); err != nil {
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
	return nil
}

func (o *xxx_SetMailSubjectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// mailSubject {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_mailSubject := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MailSubject == nil {
				o.MailSubject = &oaut.String{}
			}
			if err := o.MailSubject.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_mailSubject := func(ptr interface{}) { o.MailSubject = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MailSubject, _s_mailSubject, _ptr_mailSubject); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailSubjectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMailSubjectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMailSubjectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMailSubjectRequest structure represents the MailSubject operation request
type SetMailSubjectRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	MailSubject *oaut.String   `idl:"name:mailSubject" json:"mail_subject"`
}

func (o *SetMailSubjectRequest) xxx_ToOp(ctx context.Context, op *xxx_SetMailSubjectOperation) *xxx_SetMailSubjectOperation {
	if op == nil {
		op = &xxx_SetMailSubjectOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MailSubject = o.MailSubject
	return op
}

func (o *SetMailSubjectRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMailSubjectOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MailSubject = op.MailSubject
}
func (o *SetMailSubjectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetMailSubjectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMailSubjectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMailSubjectResponse structure represents the MailSubject operation response
type SetMailSubjectResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MailSubject return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMailSubjectResponse) xxx_ToOp(ctx context.Context, op *xxx_SetMailSubjectOperation) *xxx_SetMailSubjectOperation {
	if op == nil {
		op = &xxx_SetMailSubjectOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetMailSubjectResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMailSubjectOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMailSubjectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetMailSubjectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMailSubjectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMessageTextOperation structure represents the MessageText operation
type xxx_GetMessageTextOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MessageText *oaut.String   `idl:"name:messageText" json:"message_text"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMessageTextOperation) OpNum() int { return 24 }

func (o *xxx_GetMessageTextOperation) OpName() string { return "/IFsrmActionEmail/v0/MessageText" }

func (o *xxx_GetMessageTextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMessageTextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMessageTextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMessageTextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMessageTextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// messageText {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MessageText != nil {
			_ptr_messageText := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MessageText != nil {
					if err := o.MessageText.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MessageText, _ptr_messageText); err != nil {
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

func (o *xxx_GetMessageTextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// messageText {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_messageText := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MessageText == nil {
				o.MessageText = &oaut.String{}
			}
			if err := o.MessageText.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_messageText := func(ptr interface{}) { o.MessageText = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MessageText, _s_messageText, _ptr_messageText); err != nil {
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

// GetMessageTextRequest structure represents the MessageText operation request
type GetMessageTextRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMessageTextRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMessageTextOperation) *xxx_GetMessageTextOperation {
	if op == nil {
		op = &xxx_GetMessageTextOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetMessageTextRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMessageTextOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMessageTextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMessageTextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMessageTextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMessageTextResponse structure represents the MessageText operation response
type GetMessageTextResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MessageText *oaut.String   `idl:"name:messageText" json:"message_text"`
	// Return: The MessageText return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMessageTextResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMessageTextOperation) *xxx_GetMessageTextOperation {
	if op == nil {
		op = &xxx_GetMessageTextOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MessageText = o.MessageText
	op.Return = o.Return
	return op
}

func (o *GetMessageTextResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMessageTextOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MessageText = op.MessageText
	o.Return = op.Return
}
func (o *GetMessageTextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMessageTextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMessageTextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetMessageTextOperation structure represents the MessageText operation
type xxx_SetMessageTextOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MessageText *oaut.String   `idl:"name:messageText" json:"message_text"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetMessageTextOperation) OpNum() int { return 25 }

func (o *xxx_SetMessageTextOperation) OpName() string { return "/IFsrmActionEmail/v0/MessageText" }

func (o *xxx_SetMessageTextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMessageTextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// messageText {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MessageText != nil {
			_ptr_messageText := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MessageText != nil {
					if err := o.MessageText.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MessageText, _ptr_messageText); err != nil {
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
	return nil
}

func (o *xxx_SetMessageTextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// messageText {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_messageText := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MessageText == nil {
				o.MessageText = &oaut.String{}
			}
			if err := o.MessageText.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_messageText := func(ptr interface{}) { o.MessageText = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MessageText, _s_messageText, _ptr_messageText); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMessageTextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetMessageTextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetMessageTextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetMessageTextRequest structure represents the MessageText operation request
type SetMessageTextRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	MessageText *oaut.String   `idl:"name:messageText" json:"message_text"`
}

func (o *SetMessageTextRequest) xxx_ToOp(ctx context.Context, op *xxx_SetMessageTextOperation) *xxx_SetMessageTextOperation {
	if op == nil {
		op = &xxx_SetMessageTextOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MessageText = o.MessageText
	return op
}

func (o *SetMessageTextRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMessageTextOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MessageText = op.MessageText
}
func (o *SetMessageTextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetMessageTextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMessageTextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetMessageTextResponse structure represents the MessageText operation response
type SetMessageTextResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The MessageText return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetMessageTextResponse) xxx_ToOp(ctx context.Context, op *xxx_SetMessageTextOperation) *xxx_SetMessageTextOperation {
	if op == nil {
		op = &xxx_SetMessageTextOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetMessageTextResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMessageTextOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMessageTextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetMessageTextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMessageTextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
