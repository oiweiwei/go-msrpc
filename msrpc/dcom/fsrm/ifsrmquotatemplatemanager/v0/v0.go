package ifsrmquotatemplatemanager

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
	_ = fsrm.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmQuotaTemplateManager interface identifier 4173ac41-172d-4d52-963c-fdc7e415f717
	QuotaTemplateManagerIID = &dcom.IID{Data1: 0x4173ac41, Data2: 0x172d, Data3: 0x4d52, Data4: []byte{0x96, 0x3c, 0xfd, 0xc7, 0xe4, 0x15, 0xf7, 0x17}}
	// Syntax UUID
	QuotaTemplateManagerSyntaxUUID = &uuid.UUID{TimeLow: 0x4173ac41, TimeMid: 0x172d, TimeHiAndVersion: 0x4d52, ClockSeqHiAndReserved: 0x96, ClockSeqLow: 0x3c, Node: [6]uint8{0xfd, 0xc7, 0xe4, 0x15, 0xf7, 0x17}}
	// Syntax ID
	QuotaTemplateManagerSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: QuotaTemplateManagerSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmQuotaTemplateManager interface.
type QuotaTemplateManagerClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// CreateTemplate operation.
	CreateTemplate(context.Context, *CreateTemplateRequest, ...dcerpc.CallOption) (*CreateTemplateResponse, error)

	// GetTemplate operation.
	GetTemplate(context.Context, *GetTemplateRequest, ...dcerpc.CallOption) (*GetTemplateResponse, error)

	// EnumTemplates operation.
	EnumTemplates(context.Context, *EnumTemplatesRequest, ...dcerpc.CallOption) (*EnumTemplatesResponse, error)

	// ExportTemplates operation.
	ExportTemplates(context.Context, *ExportTemplatesRequest, ...dcerpc.CallOption) (*ExportTemplatesResponse, error)

	// ImportTemplates operation.
	ImportTemplates(context.Context, *ImportTemplatesRequest, ...dcerpc.CallOption) (*ImportTemplatesResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) QuotaTemplateManagerClient
}

type xxx_DefaultQuotaTemplateManagerClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultQuotaTemplateManagerClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultQuotaTemplateManagerClient) CreateTemplate(ctx context.Context, in *CreateTemplateRequest, opts ...dcerpc.CallOption) (*CreateTemplateResponse, error) {
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
	out := &CreateTemplateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaTemplateManagerClient) GetTemplate(ctx context.Context, in *GetTemplateRequest, opts ...dcerpc.CallOption) (*GetTemplateResponse, error) {
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
	out := &GetTemplateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaTemplateManagerClient) EnumTemplates(ctx context.Context, in *EnumTemplatesRequest, opts ...dcerpc.CallOption) (*EnumTemplatesResponse, error) {
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
	out := &EnumTemplatesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaTemplateManagerClient) ExportTemplates(ctx context.Context, in *ExportTemplatesRequest, opts ...dcerpc.CallOption) (*ExportTemplatesResponse, error) {
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
	out := &ExportTemplatesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaTemplateManagerClient) ImportTemplates(ctx context.Context, in *ImportTemplatesRequest, opts ...dcerpc.CallOption) (*ImportTemplatesResponse, error) {
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
	out := &ImportTemplatesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaTemplateManagerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultQuotaTemplateManagerClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultQuotaTemplateManagerClient) IPID(ctx context.Context, ipid *dcom.IPID) QuotaTemplateManagerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultQuotaTemplateManagerClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewQuotaTemplateManagerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (QuotaTemplateManagerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(QuotaTemplateManagerSyntaxV0_0))...)
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
	return &xxx_DefaultQuotaTemplateManagerClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_CreateTemplateOperation structure represents the CreateTemplate operation
type xxx_CreateTemplateOperation struct {
	This          *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat      `idl:"name:That" json:"that"`
	QuotaTemplate *fsrm.QuotaTemplate `idl:"name:quotaTemplate" json:"quota_template"`
	Return        int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateTemplateOperation) OpNum() int { return 7 }

func (o *xxx_CreateTemplateOperation) OpName() string {
	return "/IFsrmQuotaTemplateManager/v0/CreateTemplate"
}

func (o *xxx_CreateTemplateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTemplateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateTemplateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CreateTemplateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateTemplateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// quotaTemplate {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmQuotaTemplate}(interface))
	{
		if o.QuotaTemplate != nil {
			_ptr_quotaTemplate := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QuotaTemplate != nil {
					if err := o.QuotaTemplate.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.QuotaTemplate{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QuotaTemplate, _ptr_quotaTemplate); err != nil {
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

func (o *xxx_CreateTemplateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// quotaTemplate {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmQuotaTemplate}(interface))
	{
		_ptr_quotaTemplate := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QuotaTemplate == nil {
				o.QuotaTemplate = &fsrm.QuotaTemplate{}
			}
			if err := o.QuotaTemplate.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_quotaTemplate := func(ptr interface{}) { o.QuotaTemplate = *ptr.(**fsrm.QuotaTemplate) }
		if err := w.ReadPointer(&o.QuotaTemplate, _s_quotaTemplate, _ptr_quotaTemplate); err != nil {
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

// CreateTemplateRequest structure represents the CreateTemplate operation request
type CreateTemplateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CreateTemplateRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateTemplateOperation) *xxx_CreateTemplateOperation {
	if op == nil {
		op = &xxx_CreateTemplateOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *CreateTemplateRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateTemplateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CreateTemplateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateTemplateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateTemplateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateTemplateResponse structure represents the CreateTemplate operation response
type CreateTemplateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That          *dcom.ORPCThat      `idl:"name:That" json:"that"`
	QuotaTemplate *fsrm.QuotaTemplate `idl:"name:quotaTemplate" json:"quota_template"`
	// Return: The CreateTemplate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateTemplateResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateTemplateOperation) *xxx_CreateTemplateOperation {
	if op == nil {
		op = &xxx_CreateTemplateOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.QuotaTemplate = o.QuotaTemplate
	op.Return = o.Return
	return op
}

func (o *CreateTemplateResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateTemplateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.QuotaTemplate = op.QuotaTemplate
	o.Return = op.Return
}
func (o *CreateTemplateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateTemplateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateTemplateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTemplateOperation structure represents the GetTemplate operation
type xxx_GetTemplateOperation struct {
	This          *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat      `idl:"name:That" json:"that"`
	Name          *oaut.String        `idl:"name:name" json:"name"`
	QuotaTemplate *fsrm.QuotaTemplate `idl:"name:quotaTemplate" json:"quota_template"`
	Return        int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTemplateOperation) OpNum() int { return 8 }

func (o *xxx_GetTemplateOperation) OpName() string {
	return "/IFsrmQuotaTemplateManager/v0/GetTemplate"
}

func (o *xxx_GetTemplateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTemplateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// name {in, default_value={""}} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
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

func (o *xxx_GetTemplateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// name {in, default_value={""}} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
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

func (o *xxx_GetTemplateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTemplateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// quotaTemplate {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmQuotaTemplate}(interface))
	{
		if o.QuotaTemplate != nil {
			_ptr_quotaTemplate := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QuotaTemplate != nil {
					if err := o.QuotaTemplate.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.QuotaTemplate{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QuotaTemplate, _ptr_quotaTemplate); err != nil {
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

func (o *xxx_GetTemplateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// quotaTemplate {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmQuotaTemplate}(interface))
	{
		_ptr_quotaTemplate := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QuotaTemplate == nil {
				o.QuotaTemplate = &fsrm.QuotaTemplate{}
			}
			if err := o.QuotaTemplate.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_quotaTemplate := func(ptr interface{}) { o.QuotaTemplate = *ptr.(**fsrm.QuotaTemplate) }
		if err := w.ReadPointer(&o.QuotaTemplate, _s_quotaTemplate, _ptr_quotaTemplate); err != nil {
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

// GetTemplateRequest structure represents the GetTemplate operation request
type GetTemplateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	Name *oaut.String   `idl:"name:name" json:"name"`
}

func (o *GetTemplateRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTemplateOperation) *xxx_GetTemplateOperation {
	if op == nil {
		op = &xxx_GetTemplateOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Name = o.Name
	return op
}

func (o *GetTemplateRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTemplateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Name = op.Name
}
func (o *GetTemplateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetTemplateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTemplateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTemplateResponse structure represents the GetTemplate operation response
type GetTemplateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That          *dcom.ORPCThat      `idl:"name:That" json:"that"`
	QuotaTemplate *fsrm.QuotaTemplate `idl:"name:quotaTemplate" json:"quota_template"`
	// Return: The GetTemplate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTemplateResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTemplateOperation) *xxx_GetTemplateOperation {
	if op == nil {
		op = &xxx_GetTemplateOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.QuotaTemplate = o.QuotaTemplate
	op.Return = o.Return
	return op
}

func (o *GetTemplateResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTemplateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.QuotaTemplate = op.QuotaTemplate
	o.Return = op.Return
}
func (o *GetTemplateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetTemplateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTemplateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumTemplatesOperation structure represents the EnumTemplates operation
type xxx_EnumTemplatesOperation struct {
	This           *dcom.ORPCThis              `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat              `idl:"name:That" json:"that"`
	Options        fsrm.EnumOptions            `idl:"name:options" json:"options"`
	QuotaTemplates *fsrm.CommittableCollection `idl:"name:quotaTemplates" json:"quota_templates"`
	Return         int32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumTemplatesOperation) OpNum() int { return 9 }

func (o *xxx_EnumTemplatesOperation) OpName() string {
	return "/IFsrmQuotaTemplateManager/v0/EnumTemplates"
}

func (o *xxx_EnumTemplatesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumTemplatesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.WriteEnum(uint16(o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumTemplatesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// options {in, default_value={0}} (1:{alias=FsrmEnumOptions}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Options)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumTemplatesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumTemplatesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// quotaTemplates {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		if o.QuotaTemplates != nil {
			_ptr_quotaTemplates := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QuotaTemplates != nil {
					if err := o.QuotaTemplates.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.CommittableCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QuotaTemplates, _ptr_quotaTemplates); err != nil {
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

func (o *xxx_EnumTemplatesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// quotaTemplates {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		_ptr_quotaTemplates := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QuotaTemplates == nil {
				o.QuotaTemplates = &fsrm.CommittableCollection{}
			}
			if err := o.QuotaTemplates.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_quotaTemplates := func(ptr interface{}) { o.QuotaTemplates = *ptr.(**fsrm.CommittableCollection) }
		if err := w.ReadPointer(&o.QuotaTemplates, _s_quotaTemplates, _ptr_quotaTemplates); err != nil {
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

// EnumTemplatesRequest structure represents the EnumTemplates operation request
type EnumTemplatesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis   `idl:"name:This" json:"this"`
	Options fsrm.EnumOptions `idl:"name:options" json:"options"`
}

func (o *EnumTemplatesRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumTemplatesOperation) *xxx_EnumTemplatesOperation {
	if op == nil {
		op = &xxx_EnumTemplatesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Options = o.Options
	return op
}

func (o *EnumTemplatesRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumTemplatesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Options = op.Options
}
func (o *EnumTemplatesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumTemplatesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumTemplatesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumTemplatesResponse structure represents the EnumTemplates operation response
type EnumTemplatesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat              `idl:"name:That" json:"that"`
	QuotaTemplates *fsrm.CommittableCollection `idl:"name:quotaTemplates" json:"quota_templates"`
	// Return: The EnumTemplates return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnumTemplatesResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumTemplatesOperation) *xxx_EnumTemplatesOperation {
	if op == nil {
		op = &xxx_EnumTemplatesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.QuotaTemplates = o.QuotaTemplates
	op.Return = o.Return
	return op
}

func (o *EnumTemplatesResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumTemplatesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.QuotaTemplates = op.QuotaTemplates
	o.Return = op.Return
}
func (o *EnumTemplatesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumTemplatesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumTemplatesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ExportTemplatesOperation structure represents the ExportTemplates operation
type xxx_ExportTemplatesOperation struct {
	This                     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                     *dcom.ORPCThat `idl:"name:That" json:"that"`
	QuotaTemplateNamesArray  *oaut.Variant  `idl:"name:quotaTemplateNamesArray" json:"quota_template_names_array"`
	SerializedQuotaTemplates *oaut.String   `idl:"name:serializedQuotaTemplates" json:"serialized_quota_templates"`
	Return                   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ExportTemplatesOperation) OpNum() int { return 10 }

func (o *xxx_ExportTemplatesOperation) OpName() string {
	return "/IFsrmQuotaTemplateManager/v0/ExportTemplates"
}

func (o *xxx_ExportTemplatesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportTemplatesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// quotaTemplateNamesArray {in, default_value={<nil>}} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.QuotaTemplateNamesArray != nil {
			_ptr_quotaTemplateNamesArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QuotaTemplateNamesArray != nil {
					if err := o.QuotaTemplateNamesArray.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QuotaTemplateNamesArray, _ptr_quotaTemplateNamesArray); err != nil {
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

func (o *xxx_ExportTemplatesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// quotaTemplateNamesArray {in, default_value={<nil>}} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_quotaTemplateNamesArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QuotaTemplateNamesArray == nil {
				o.QuotaTemplateNamesArray = &oaut.Variant{}
			}
			if err := o.QuotaTemplateNamesArray.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_quotaTemplateNamesArray := func(ptr interface{}) { o.QuotaTemplateNamesArray = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.QuotaTemplateNamesArray, _s_quotaTemplateNamesArray, _ptr_quotaTemplateNamesArray); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportTemplatesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExportTemplatesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// serializedQuotaTemplates {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SerializedQuotaTemplates != nil {
			_ptr_serializedQuotaTemplates := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SerializedQuotaTemplates != nil {
					if err := o.SerializedQuotaTemplates.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SerializedQuotaTemplates, _ptr_serializedQuotaTemplates); err != nil {
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

func (o *xxx_ExportTemplatesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// serializedQuotaTemplates {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_serializedQuotaTemplates := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SerializedQuotaTemplates == nil {
				o.SerializedQuotaTemplates = &oaut.String{}
			}
			if err := o.SerializedQuotaTemplates.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_serializedQuotaTemplates := func(ptr interface{}) { o.SerializedQuotaTemplates = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SerializedQuotaTemplates, _s_serializedQuotaTemplates, _ptr_serializedQuotaTemplates); err != nil {
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

// ExportTemplatesRequest structure represents the ExportTemplates operation request
type ExportTemplatesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                    *dcom.ORPCThis `idl:"name:This" json:"this"`
	QuotaTemplateNamesArray *oaut.Variant  `idl:"name:quotaTemplateNamesArray" json:"quota_template_names_array"`
}

func (o *ExportTemplatesRequest) xxx_ToOp(ctx context.Context, op *xxx_ExportTemplatesOperation) *xxx_ExportTemplatesOperation {
	if op == nil {
		op = &xxx_ExportTemplatesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.QuotaTemplateNamesArray = o.QuotaTemplateNamesArray
	return op
}

func (o *ExportTemplatesRequest) xxx_FromOp(ctx context.Context, op *xxx_ExportTemplatesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.QuotaTemplateNamesArray = op.QuotaTemplateNamesArray
}
func (o *ExportTemplatesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ExportTemplatesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportTemplatesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExportTemplatesResponse structure represents the ExportTemplates operation response
type ExportTemplatesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That                     *dcom.ORPCThat `idl:"name:That" json:"that"`
	SerializedQuotaTemplates *oaut.String   `idl:"name:serializedQuotaTemplates" json:"serialized_quota_templates"`
	// Return: The ExportTemplates return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExportTemplatesResponse) xxx_ToOp(ctx context.Context, op *xxx_ExportTemplatesOperation) *xxx_ExportTemplatesOperation {
	if op == nil {
		op = &xxx_ExportTemplatesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.SerializedQuotaTemplates = o.SerializedQuotaTemplates
	op.Return = o.Return
	return op
}

func (o *ExportTemplatesResponse) xxx_FromOp(ctx context.Context, op *xxx_ExportTemplatesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SerializedQuotaTemplates = op.SerializedQuotaTemplates
	o.Return = op.Return
}
func (o *ExportTemplatesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ExportTemplatesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExportTemplatesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ImportTemplatesOperation structure represents the ImportTemplates operation
type xxx_ImportTemplatesOperation struct {
	This                     *dcom.ORPCThis              `idl:"name:This" json:"this"`
	That                     *dcom.ORPCThat              `idl:"name:That" json:"that"`
	SerializedQuotaTemplates *oaut.String                `idl:"name:serializedQuotaTemplates" json:"serialized_quota_templates"`
	QuotaTemplateNamesArray  *oaut.Variant               `idl:"name:quotaTemplateNamesArray" json:"quota_template_names_array"`
	QuotaTemplates           *fsrm.CommittableCollection `idl:"name:quotaTemplates" json:"quota_templates"`
	Return                   int32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_ImportTemplatesOperation) OpNum() int { return 11 }

func (o *xxx_ImportTemplatesOperation) OpName() string {
	return "/IFsrmQuotaTemplateManager/v0/ImportTemplates"
}

func (o *xxx_ImportTemplatesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportTemplatesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// serializedQuotaTemplates {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.SerializedQuotaTemplates != nil {
			_ptr_serializedQuotaTemplates := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SerializedQuotaTemplates != nil {
					if err := o.SerializedQuotaTemplates.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SerializedQuotaTemplates, _ptr_serializedQuotaTemplates); err != nil {
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
	// quotaTemplateNamesArray {in, default_value={<nil>}} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.QuotaTemplateNamesArray != nil {
			_ptr_quotaTemplateNamesArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QuotaTemplateNamesArray != nil {
					if err := o.QuotaTemplateNamesArray.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QuotaTemplateNamesArray, _ptr_quotaTemplateNamesArray); err != nil {
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

func (o *xxx_ImportTemplatesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// serializedQuotaTemplates {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_serializedQuotaTemplates := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SerializedQuotaTemplates == nil {
				o.SerializedQuotaTemplates = &oaut.String{}
			}
			if err := o.SerializedQuotaTemplates.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_serializedQuotaTemplates := func(ptr interface{}) { o.SerializedQuotaTemplates = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.SerializedQuotaTemplates, _s_serializedQuotaTemplates, _ptr_serializedQuotaTemplates); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// quotaTemplateNamesArray {in, default_value={<nil>}} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_quotaTemplateNamesArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QuotaTemplateNamesArray == nil {
				o.QuotaTemplateNamesArray = &oaut.Variant{}
			}
			if err := o.QuotaTemplateNamesArray.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_quotaTemplateNamesArray := func(ptr interface{}) { o.QuotaTemplateNamesArray = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.QuotaTemplateNamesArray, _s_quotaTemplateNamesArray, _ptr_quotaTemplateNamesArray); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportTemplatesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ImportTemplatesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// quotaTemplates {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		if o.QuotaTemplates != nil {
			_ptr_quotaTemplates := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QuotaTemplates != nil {
					if err := o.QuotaTemplates.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&fsrm.CommittableCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QuotaTemplates, _ptr_quotaTemplates); err != nil {
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

func (o *xxx_ImportTemplatesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// quotaTemplates {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IFsrmCommittableCollection}(interface))
	{
		_ptr_quotaTemplates := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QuotaTemplates == nil {
				o.QuotaTemplates = &fsrm.CommittableCollection{}
			}
			if err := o.QuotaTemplates.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_quotaTemplates := func(ptr interface{}) { o.QuotaTemplates = *ptr.(**fsrm.CommittableCollection) }
		if err := w.ReadPointer(&o.QuotaTemplates, _s_quotaTemplates, _ptr_quotaTemplates); err != nil {
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

// ImportTemplatesRequest structure represents the ImportTemplates operation request
type ImportTemplatesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                     *dcom.ORPCThis `idl:"name:This" json:"this"`
	SerializedQuotaTemplates *oaut.String   `idl:"name:serializedQuotaTemplates" json:"serialized_quota_templates"`
	QuotaTemplateNamesArray  *oaut.Variant  `idl:"name:quotaTemplateNamesArray" json:"quota_template_names_array"`
}

func (o *ImportTemplatesRequest) xxx_ToOp(ctx context.Context, op *xxx_ImportTemplatesOperation) *xxx_ImportTemplatesOperation {
	if op == nil {
		op = &xxx_ImportTemplatesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.SerializedQuotaTemplates = o.SerializedQuotaTemplates
	op.QuotaTemplateNamesArray = o.QuotaTemplateNamesArray
	return op
}

func (o *ImportTemplatesRequest) xxx_FromOp(ctx context.Context, op *xxx_ImportTemplatesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SerializedQuotaTemplates = op.SerializedQuotaTemplates
	o.QuotaTemplateNamesArray = op.QuotaTemplateNamesArray
}
func (o *ImportTemplatesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ImportTemplatesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportTemplatesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ImportTemplatesResponse structure represents the ImportTemplates operation response
type ImportTemplatesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat              `idl:"name:That" json:"that"`
	QuotaTemplates *fsrm.CommittableCollection `idl:"name:quotaTemplates" json:"quota_templates"`
	// Return: The ImportTemplates return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ImportTemplatesResponse) xxx_ToOp(ctx context.Context, op *xxx_ImportTemplatesOperation) *xxx_ImportTemplatesOperation {
	if op == nil {
		op = &xxx_ImportTemplatesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.QuotaTemplates = o.QuotaTemplates
	op.Return = o.Return
	return op
}

func (o *ImportTemplatesResponse) xxx_FromOp(ctx context.Context, op *xxx_ImportTemplatesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.QuotaTemplates = op.QuotaTemplates
	o.Return = op.Return
}
func (o *ImportTemplatesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ImportTemplatesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ImportTemplatesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
