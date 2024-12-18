package ifsrmactionreport

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
	// IFsrmActionReport interface identifier 2dbe63c4-b340-48a0-a5b0-158e07fc567e
	ActionReportIID = &dcom.IID{Data1: 0x2dbe63c4, Data2: 0xb340, Data3: 0x48a0, Data4: []byte{0xa5, 0xb0, 0x15, 0x8e, 0x07, 0xfc, 0x56, 0x7e}}
	// Syntax UUID
	ActionReportSyntaxUUID = &uuid.UUID{TimeLow: 0x2dbe63c4, TimeMid: 0xb340, TimeHiAndVersion: 0x48a0, ClockSeqHiAndReserved: 0xa5, ClockSeqLow: 0xb0, Node: [6]uint8{0x15, 0x8e, 0x7, 0xfc, 0x56, 0x7e}}
	// Syntax ID
	ActionReportSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ActionReportSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmActionReport interface.
type ActionReportClient interface {

	// IFsrmAction retrieval method.
	Action() ifsrmaction.ActionClient

	// ReportTypes operation.
	GetReportTypes(context.Context, *GetReportTypesRequest, ...dcerpc.CallOption) (*GetReportTypesResponse, error)

	// ReportTypes operation.
	SetReportTypes(context.Context, *SetReportTypesRequest, ...dcerpc.CallOption) (*SetReportTypesResponse, error)

	// MailTo operation.
	GetMailTo(context.Context, *GetMailToRequest, ...dcerpc.CallOption) (*GetMailToResponse, error)

	// MailTo operation.
	SetMailTo(context.Context, *SetMailToRequest, ...dcerpc.CallOption) (*SetMailToResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ActionReportClient
}

type xxx_DefaultActionReportClient struct {
	ifsrmaction.ActionClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultActionReportClient) Action() ifsrmaction.ActionClient {
	return o.ActionClient
}

func (o *xxx_DefaultActionReportClient) GetReportTypes(ctx context.Context, in *GetReportTypesRequest, opts ...dcerpc.CallOption) (*GetReportTypesResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetReportTypesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionReportClient) SetReportTypes(ctx context.Context, in *SetReportTypesRequest, opts ...dcerpc.CallOption) (*SetReportTypesResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &SetReportTypesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultActionReportClient) GetMailTo(ctx context.Context, in *GetMailToRequest, opts ...dcerpc.CallOption) (*GetMailToResponse, error) {
	op := in.xxx_ToOp(ctx)
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

func (o *xxx_DefaultActionReportClient) SetMailTo(ctx context.Context, in *SetMailToRequest, opts ...dcerpc.CallOption) (*SetMailToResponse, error) {
	op := in.xxx_ToOp(ctx)
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

func (o *xxx_DefaultActionReportClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultActionReportClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultActionReportClient) IPID(ctx context.Context, ipid *dcom.IPID) ActionReportClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultActionReportClient{
		ActionClient: o.ActionClient.IPID(ctx, ipid),
		cc:           o.cc,
		ipid:         ipid,
	}
}

func NewActionReportClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ActionReportClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ActionReportSyntaxV0_0))...)
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
	return &xxx_DefaultActionReportClient{
		ActionClient: base,
		cc:           cc,
		ipid:         ipid,
	}, nil
}

// xxx_GetReportTypesOperation structure represents the ReportTypes operation
type xxx_GetReportTypesOperation struct {
	This        *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ReportTypes *oaut.SafeArray `idl:"name:reportTypes" json:"report_types"`
	Return      int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetReportTypesOperation) OpNum() int { return 12 }

func (o *xxx_GetReportTypesOperation) OpName() string { return "/IFsrmActionReport/v0/ReportTypes" }

func (o *xxx_GetReportTypesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReportTypesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetReportTypesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetReportTypesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetReportTypesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// reportTypes {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.ReportTypes != nil {
			_ptr_reportTypes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReportTypes != nil {
					if err := o.ReportTypes.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReportTypes, _ptr_reportTypes); err != nil {
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

func (o *xxx_GetReportTypesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// reportTypes {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_reportTypes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReportTypes == nil {
				o.ReportTypes = &oaut.SafeArray{}
			}
			if err := o.ReportTypes.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_reportTypes := func(ptr interface{}) { o.ReportTypes = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.ReportTypes, _s_reportTypes, _ptr_reportTypes); err != nil {
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

// GetReportTypesRequest structure represents the ReportTypes operation request
type GetReportTypesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetReportTypesRequest) xxx_ToOp(ctx context.Context) *xxx_GetReportTypesOperation {
	if o == nil {
		return &xxx_GetReportTypesOperation{}
	}
	return &xxx_GetReportTypesOperation{
		This: o.This,
	}
}

func (o *GetReportTypesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetReportTypesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetReportTypesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetReportTypesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReportTypesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetReportTypesResponse structure represents the ReportTypes operation response
type GetReportTypesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ReportTypes *oaut.SafeArray `idl:"name:reportTypes" json:"report_types"`
	// Return: The ReportTypes return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetReportTypesResponse) xxx_ToOp(ctx context.Context) *xxx_GetReportTypesOperation {
	if o == nil {
		return &xxx_GetReportTypesOperation{}
	}
	return &xxx_GetReportTypesOperation{
		That:        o.That,
		ReportTypes: o.ReportTypes,
		Return:      o.Return,
	}
}

func (o *GetReportTypesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetReportTypesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReportTypes = op.ReportTypes
	o.Return = op.Return
}
func (o *GetReportTypesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetReportTypesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetReportTypesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetReportTypesOperation structure represents the ReportTypes operation
type xxx_SetReportTypesOperation struct {
	This        *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat  `idl:"name:That" json:"that"`
	ReportTypes *oaut.SafeArray `idl:"name:reportTypes" json:"report_types"`
	Return      int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetReportTypesOperation) OpNum() int { return 13 }

func (o *xxx_SetReportTypesOperation) OpName() string { return "/IFsrmActionReport/v0/ReportTypes" }

func (o *xxx_SetReportTypesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetReportTypesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// reportTypes {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.ReportTypes != nil {
			_ptr_reportTypes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReportTypes != nil {
					if err := o.ReportTypes.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReportTypes, _ptr_reportTypes); err != nil {
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

func (o *xxx_SetReportTypesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// reportTypes {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_reportTypes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReportTypes == nil {
				o.ReportTypes = &oaut.SafeArray{}
			}
			if err := o.ReportTypes.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_reportTypes := func(ptr interface{}) { o.ReportTypes = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.ReportTypes, _s_reportTypes, _ptr_reportTypes); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetReportTypesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetReportTypesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetReportTypesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetReportTypesRequest structure represents the ReportTypes operation request
type SetReportTypesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis  `idl:"name:This" json:"this"`
	ReportTypes *oaut.SafeArray `idl:"name:reportTypes" json:"report_types"`
}

func (o *SetReportTypesRequest) xxx_ToOp(ctx context.Context) *xxx_SetReportTypesOperation {
	if o == nil {
		return &xxx_SetReportTypesOperation{}
	}
	return &xxx_SetReportTypesOperation{
		This:        o.This,
		ReportTypes: o.ReportTypes,
	}
}

func (o *SetReportTypesRequest) xxx_FromOp(ctx context.Context, op *xxx_SetReportTypesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ReportTypes = op.ReportTypes
}
func (o *SetReportTypesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetReportTypesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetReportTypesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetReportTypesResponse structure represents the ReportTypes operation response
type SetReportTypesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ReportTypes return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetReportTypesResponse) xxx_ToOp(ctx context.Context) *xxx_SetReportTypesOperation {
	if o == nil {
		return &xxx_SetReportTypesOperation{}
	}
	return &xxx_SetReportTypesOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetReportTypesResponse) xxx_FromOp(ctx context.Context, op *xxx_SetReportTypesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetReportTypesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetReportTypesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetReportTypesOperation{}
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

func (o *xxx_GetMailToOperation) OpNum() int { return 14 }

func (o *xxx_GetMailToOperation) OpName() string { return "/IFsrmActionReport/v0/MailTo" }

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

func (o *GetMailToRequest) xxx_ToOp(ctx context.Context) *xxx_GetMailToOperation {
	if o == nil {
		return &xxx_GetMailToOperation{}
	}
	return &xxx_GetMailToOperation{
		This: o.This,
	}
}

func (o *GetMailToRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMailToOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMailToRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
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

func (o *GetMailToResponse) xxx_ToOp(ctx context.Context) *xxx_GetMailToOperation {
	if o == nil {
		return &xxx_GetMailToOperation{}
	}
	return &xxx_GetMailToOperation{
		That:   o.That,
		MailTo: o.MailTo,
		Return: o.Return,
	}
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
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
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

func (o *xxx_SetMailToOperation) OpNum() int { return 15 }

func (o *xxx_SetMailToOperation) OpName() string { return "/IFsrmActionReport/v0/MailTo" }

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

func (o *SetMailToRequest) xxx_ToOp(ctx context.Context) *xxx_SetMailToOperation {
	if o == nil {
		return &xxx_SetMailToOperation{}
	}
	return &xxx_SetMailToOperation{
		This:   o.This,
		MailTo: o.MailTo,
	}
}

func (o *SetMailToRequest) xxx_FromOp(ctx context.Context, op *xxx_SetMailToOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MailTo = op.MailTo
}
func (o *SetMailToRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
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

func (o *SetMailToResponse) xxx_ToOp(ctx context.Context) *xxx_SetMailToOperation {
	if o == nil {
		return &xxx_SetMailToOperation{}
	}
	return &xxx_SetMailToOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetMailToResponse) xxx_FromOp(ctx context.Context, op *xxx_SetMailToOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetMailToResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetMailToResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetMailToOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
