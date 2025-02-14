package ifsrmquotatemplate

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	fsrm "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm"
	ifsrmquotabase "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmquotabase/v0"
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
	_ = ifsrmquotabase.GoPackage
	_ = oaut.GoPackage
	_ = fsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmQuotaTemplate interface identifier a2efab31-295e-46bb-b976-e86d58b52e8b
	QuotaTemplateIID = &dcom.IID{Data1: 0xa2efab31, Data2: 0x295e, Data3: 0x46bb, Data4: []byte{0xb9, 0x76, 0xe8, 0x6d, 0x58, 0xb5, 0x2e, 0x8b}}
	// Syntax UUID
	QuotaTemplateSyntaxUUID = &uuid.UUID{TimeLow: 0xa2efab31, TimeMid: 0x295e, TimeHiAndVersion: 0x46bb, ClockSeqHiAndReserved: 0xb9, ClockSeqLow: 0x76, Node: [6]uint8{0xe8, 0x6d, 0x58, 0xb5, 0x2e, 0x8b}}
	// Syntax ID
	QuotaTemplateSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: QuotaTemplateSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmQuotaTemplate interface.
type QuotaTemplateClient interface {

	// IFsrmQuotaBase retrieval method.
	QuotaBase() ifsrmquotabase.QuotaBaseClient

	// Name operation.
	GetName(context.Context, *GetNameRequest, ...dcerpc.CallOption) (*GetNameResponse, error)

	// Name operation.
	SetName(context.Context, *SetNameRequest, ...dcerpc.CallOption) (*SetNameResponse, error)

	// CopyTemplate operation.
	CopyTemplate(context.Context, *CopyTemplateRequest, ...dcerpc.CallOption) (*CopyTemplateResponse, error)

	// CommitAndUpdateDerived operation.
	CommitAndUpdateDerived(context.Context, *CommitAndUpdateDerivedRequest, ...dcerpc.CallOption) (*CommitAndUpdateDerivedResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) QuotaTemplateClient
}

type xxx_DefaultQuotaTemplateClient struct {
	ifsrmquotabase.QuotaBaseClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultQuotaTemplateClient) QuotaBase() ifsrmquotabase.QuotaBaseClient {
	return o.QuotaBaseClient
}

func (o *xxx_DefaultQuotaTemplateClient) GetName(ctx context.Context, in *GetNameRequest, opts ...dcerpc.CallOption) (*GetNameResponse, error) {
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
	out := &GetNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaTemplateClient) SetName(ctx context.Context, in *SetNameRequest, opts ...dcerpc.CallOption) (*SetNameResponse, error) {
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
	out := &SetNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaTemplateClient) CopyTemplate(ctx context.Context, in *CopyTemplateRequest, opts ...dcerpc.CallOption) (*CopyTemplateResponse, error) {
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
	out := &CopyTemplateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaTemplateClient) CommitAndUpdateDerived(ctx context.Context, in *CommitAndUpdateDerivedRequest, opts ...dcerpc.CallOption) (*CommitAndUpdateDerivedResponse, error) {
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
	out := &CommitAndUpdateDerivedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaTemplateClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultQuotaTemplateClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultQuotaTemplateClient) IPID(ctx context.Context, ipid *dcom.IPID) QuotaTemplateClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultQuotaTemplateClient{
		QuotaBaseClient: o.QuotaBaseClient.IPID(ctx, ipid),
		cc:              o.cc,
		ipid:            ipid,
	}
}

func NewQuotaTemplateClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (QuotaTemplateClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(QuotaTemplateSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ifsrmquotabase.NewQuotaBaseClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultQuotaTemplateClient{
		QuotaBaseClient: base,
		cc:              cc,
		ipid:            ipid,
	}, nil
}

// xxx_GetNameOperation structure represents the Name operation
type xxx_GetNameOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name   *oaut.String   `idl:"name:name" json:"name"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNameOperation) OpNum() int { return 22 }

func (o *xxx_GetNameOperation) OpName() string { return "/IFsrmQuotaTemplate/v0/Name" }

func (o *xxx_GetNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// name {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Name != nil {
			_ptr_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Name != nil {
					if err := o.Name.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_name); err != nil {
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

func (o *xxx_GetNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// name {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Name == nil {
				o.Name = &oaut.String{}
			}
			if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_name := func(ptr interface{}) { o.Name = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Name, _s_name, _ptr_name); err != nil {
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

// GetNameRequest structure represents the Name operation request
type GetNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNameOperation) *xxx_GetNameOperation {
	if op == nil {
		op = &xxx_GetNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNameResponse structure represents the Name operation response
type GetNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name *oaut.String   `idl:"name:name" json:"name"`
	// Return: The Name return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNameOperation) *xxx_GetNameOperation {
	if op == nil {
		op = &xxx_GetNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Name = o.Name
	op.Return = o.Return
	return op
}

func (o *GetNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Name = op.Name
	o.Return = op.Return
}
func (o *GetNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetNameOperation structure represents the Name operation
type xxx_SetNameOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Name   *oaut.String   `idl:"name:name" json:"name"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNameOperation) OpNum() int { return 23 }

func (o *xxx_SetNameOperation) OpName() string { return "/IFsrmQuotaTemplate/v0/Name" }

func (o *xxx_SetNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// name {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Name != nil {
			_ptr_name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Name != nil {
					if err := o.Name.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_name); err != nil {
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

func (o *xxx_SetNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// name {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Name == nil {
				o.Name = &oaut.String{}
			}
			if err := o.Name.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_name := func(ptr interface{}) { o.Name = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Name, _s_name, _ptr_name); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetNameRequest structure represents the Name operation request
type SetNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	Name *oaut.String   `idl:"name:name" json:"name"`
}

func (o *SetNameRequest) xxx_ToOp(ctx context.Context, op *xxx_SetNameOperation) *xxx_SetNameOperation {
	if op == nil {
		op = &xxx_SetNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Name = o.Name
	return op
}

func (o *SetNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Name = op.Name
}
func (o *SetNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNameResponse structure represents the Name operation response
type SetNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Name return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetNameResponse) xxx_ToOp(ctx context.Context, op *xxx_SetNameOperation) *xxx_SetNameOperation {
	if op == nil {
		op = &xxx_SetNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CopyTemplateOperation structure represents the CopyTemplate operation
type xxx_CopyTemplateOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	QuotaTemplateName *oaut.String   `idl:"name:quotaTemplateName" json:"quota_template_name"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CopyTemplateOperation) OpNum() int { return 24 }

func (o *xxx_CopyTemplateOperation) OpName() string { return "/IFsrmQuotaTemplate/v0/CopyTemplate" }

func (o *xxx_CopyTemplateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CopyTemplateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// quotaTemplateName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.QuotaTemplateName != nil {
			_ptr_quotaTemplateName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QuotaTemplateName != nil {
					if err := o.QuotaTemplateName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QuotaTemplateName, _ptr_quotaTemplateName); err != nil {
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

func (o *xxx_CopyTemplateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// quotaTemplateName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_quotaTemplateName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QuotaTemplateName == nil {
				o.QuotaTemplateName = &oaut.String{}
			}
			if err := o.QuotaTemplateName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_quotaTemplateName := func(ptr interface{}) { o.QuotaTemplateName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.QuotaTemplateName, _s_quotaTemplateName, _ptr_quotaTemplateName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CopyTemplateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CopyTemplateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CopyTemplateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CopyTemplateRequest structure represents the CopyTemplate operation request
type CopyTemplateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	QuotaTemplateName *oaut.String   `idl:"name:quotaTemplateName" json:"quota_template_name"`
}

func (o *CopyTemplateRequest) xxx_ToOp(ctx context.Context, op *xxx_CopyTemplateOperation) *xxx_CopyTemplateOperation {
	if op == nil {
		op = &xxx_CopyTemplateOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.QuotaTemplateName = o.QuotaTemplateName
	return op
}

func (o *CopyTemplateRequest) xxx_FromOp(ctx context.Context, op *xxx_CopyTemplateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.QuotaTemplateName = op.QuotaTemplateName
}
func (o *CopyTemplateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CopyTemplateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CopyTemplateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CopyTemplateResponse structure represents the CopyTemplate operation response
type CopyTemplateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CopyTemplate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CopyTemplateResponse) xxx_ToOp(ctx context.Context, op *xxx_CopyTemplateOperation) *xxx_CopyTemplateOperation {
	if op == nil {
		op = &xxx_CopyTemplateOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CopyTemplateResponse) xxx_FromOp(ctx context.Context, op *xxx_CopyTemplateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CopyTemplateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CopyTemplateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CopyTemplateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CommitAndUpdateDerivedOperation structure represents the CommitAndUpdateDerived operation
type xxx_CommitAndUpdateDerivedOperation struct {
	This                 *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat             `idl:"name:That" json:"that"`
	CommitOptions        fsrm.CommitOptions         `idl:"name:commitOptions" json:"commit_options"`
	ApplyOptions         fsrm.TemplateApplyOptions  `idl:"name:applyOptions" json:"apply_options"`
	DerivedObjectsResult *fsrm.DerivedObjectsResult `idl:"name:derivedObjectsResult" json:"derived_objects_result"`
	Return               int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_CommitAndUpdateDerivedOperation) OpNum() int { return 25 }

func (o *xxx_CommitAndUpdateDerivedOperation) OpName() string {
	return "/IFsrmQuotaTemplate/v0/CommitAndUpdateDerived"
}

func (o *xxx_CommitAndUpdateDerivedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitAndUpdateDerivedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// commitOptions {in} (1:{alias=FsrmCommitOptions}(enum))
	{
		if err := w.WriteEnum(uint16(o.CommitOptions)); err != nil {
			return err
		}
	}
	// applyOptions {in} (1:{alias=FsrmTemplateApplyOptions}(enum))
	{
		if err := w.WriteEnum(uint16(o.ApplyOptions)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitAndUpdateDerivedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// commitOptions {in} (1:{alias=FsrmCommitOptions}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.CommitOptions)); err != nil {
			return err
		}
	}
	// applyOptions {in} (1:{alias=FsrmTemplateApplyOptions}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.ApplyOptions)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitAndUpdateDerivedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitAndUpdateDerivedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// derivedObjectsResult {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmDerivedObjectsResult}(interface))
	{
		if o.DerivedObjectsResult != nil {
			_ptr_derivedObjectsResult := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.DerivedObjectsResult != nil {
					if err := o.DerivedObjectsResult.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.DerivedObjectsResult{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DerivedObjectsResult, _ptr_derivedObjectsResult); err != nil {
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

func (o *xxx_CommitAndUpdateDerivedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// derivedObjectsResult {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmDerivedObjectsResult}(interface))
	{
		_ptr_derivedObjectsResult := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.DerivedObjectsResult == nil {
				o.DerivedObjectsResult = &fsrm.DerivedObjectsResult{}
			}
			if err := o.DerivedObjectsResult.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_derivedObjectsResult := func(ptr interface{}) { o.DerivedObjectsResult = *ptr.(**fsrm.DerivedObjectsResult) }
		if err := w.ReadPointer(&o.DerivedObjectsResult, _s_derivedObjectsResult, _ptr_derivedObjectsResult); err != nil {
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

// CommitAndUpdateDerivedRequest structure represents the CommitAndUpdateDerived operation request
type CommitAndUpdateDerivedRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This          *dcom.ORPCThis            `idl:"name:This" json:"this"`
	CommitOptions fsrm.CommitOptions        `idl:"name:commitOptions" json:"commit_options"`
	ApplyOptions  fsrm.TemplateApplyOptions `idl:"name:applyOptions" json:"apply_options"`
}

func (o *CommitAndUpdateDerivedRequest) xxx_ToOp(ctx context.Context, op *xxx_CommitAndUpdateDerivedOperation) *xxx_CommitAndUpdateDerivedOperation {
	if op == nil {
		op = &xxx_CommitAndUpdateDerivedOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.CommitOptions = o.CommitOptions
	op.ApplyOptions = o.ApplyOptions
	return op
}

func (o *CommitAndUpdateDerivedRequest) xxx_FromOp(ctx context.Context, op *xxx_CommitAndUpdateDerivedOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.CommitOptions = op.CommitOptions
	o.ApplyOptions = op.ApplyOptions
}
func (o *CommitAndUpdateDerivedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CommitAndUpdateDerivedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CommitAndUpdateDerivedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CommitAndUpdateDerivedResponse structure represents the CommitAndUpdateDerived operation response
type CommitAndUpdateDerivedResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                 *dcom.ORPCThat             `idl:"name:That" json:"that"`
	DerivedObjectsResult *fsrm.DerivedObjectsResult `idl:"name:derivedObjectsResult" json:"derived_objects_result"`
	// Return: The CommitAndUpdateDerived return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CommitAndUpdateDerivedResponse) xxx_ToOp(ctx context.Context, op *xxx_CommitAndUpdateDerivedOperation) *xxx_CommitAndUpdateDerivedOperation {
	if op == nil {
		op = &xxx_CommitAndUpdateDerivedOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.DerivedObjectsResult = o.DerivedObjectsResult
	op.Return = o.Return
	return op
}

func (o *CommitAndUpdateDerivedResponse) xxx_FromOp(ctx context.Context, op *xxx_CommitAndUpdateDerivedOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DerivedObjectsResult = op.DerivedObjectsResult
	o.Return = op.Return
}
func (o *CommitAndUpdateDerivedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CommitAndUpdateDerivedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CommitAndUpdateDerivedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
