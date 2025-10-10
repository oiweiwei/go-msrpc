package ifsrmquotatemplateimported

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	ifsrmquotatemplate "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmquotatemplate/v0"
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
	_ = ifsrmquotatemplate.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmQuotaTemplateImported interface identifier 9a2bf113-a329-44cc-809a-5c00fce8da40
	QuotaTemplateImportedIID = &dcom.IID{Data1: 0x9a2bf113, Data2: 0xa329, Data3: 0x44cc, Data4: []byte{0x80, 0x9a, 0x5c, 0x00, 0xfc, 0xe8, 0xda, 0x40}}
	// Syntax UUID
	QuotaTemplateImportedSyntaxUUID = &uuid.UUID{TimeLow: 0x9a2bf113, TimeMid: 0xa329, TimeHiAndVersion: 0x44cc, ClockSeqHiAndReserved: 0x80, ClockSeqLow: 0x9a, Node: [6]uint8{0x5c, 0x0, 0xfc, 0xe8, 0xda, 0x40}}
	// Syntax ID
	QuotaTemplateImportedSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: QuotaTemplateImportedSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmQuotaTemplateImported interface.
type QuotaTemplateImportedClient interface {

	// IFsrmQuotaTemplate retrieval method.
	QuotaTemplate() ifsrmquotatemplate.QuotaTemplateClient

	// OverwriteOnCommit operation.
	GetOverwriteOnCommit(context.Context, *GetOverwriteOnCommitRequest, ...dcerpc.CallOption) (*GetOverwriteOnCommitResponse, error)

	// OverwriteOnCommit operation.
	SetOverwriteOnCommit(context.Context, *SetOverwriteOnCommitRequest, ...dcerpc.CallOption) (*SetOverwriteOnCommitResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) QuotaTemplateImportedClient
}

type xxx_DefaultQuotaTemplateImportedClient struct {
	ifsrmquotatemplate.QuotaTemplateClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultQuotaTemplateImportedClient) QuotaTemplate() ifsrmquotatemplate.QuotaTemplateClient {
	return o.QuotaTemplateClient
}

func (o *xxx_DefaultQuotaTemplateImportedClient) GetOverwriteOnCommit(ctx context.Context, in *GetOverwriteOnCommitRequest, opts ...dcerpc.CallOption) (*GetOverwriteOnCommitResponse, error) {
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
	out := &GetOverwriteOnCommitResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaTemplateImportedClient) SetOverwriteOnCommit(ctx context.Context, in *SetOverwriteOnCommitRequest, opts ...dcerpc.CallOption) (*SetOverwriteOnCommitResponse, error) {
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
	out := &SetOverwriteOnCommitResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaTemplateImportedClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultQuotaTemplateImportedClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultQuotaTemplateImportedClient) IPID(ctx context.Context, ipid *dcom.IPID) QuotaTemplateImportedClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultQuotaTemplateImportedClient{
		QuotaTemplateClient: o.QuotaTemplateClient.IPID(ctx, ipid),
		cc:                  o.cc,
		ipid:                ipid,
	}
}

func NewQuotaTemplateImportedClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (QuotaTemplateImportedClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(QuotaTemplateImportedSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ifsrmquotatemplate.NewQuotaTemplateClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultQuotaTemplateImportedClient{
		QuotaTemplateClient: base,
		cc:                  cc,
		ipid:                ipid,
	}, nil
}

// xxx_GetOverwriteOnCommitOperation structure represents the OverwriteOnCommit operation
type xxx_GetOverwriteOnCommitOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Overwrite int16          `idl:"name:overwrite" json:"overwrite"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetOverwriteOnCommitOperation) OpNum() int { return 26 }

func (o *xxx_GetOverwriteOnCommitOperation) OpName() string {
	return "/IFsrmQuotaTemplateImported/v0/OverwriteOnCommit"
}

func (o *xxx_GetOverwriteOnCommitOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOverwriteOnCommitOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetOverwriteOnCommitOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetOverwriteOnCommitOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOverwriteOnCommitOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// overwrite {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Overwrite); err != nil {
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

func (o *xxx_GetOverwriteOnCommitOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// overwrite {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Overwrite); err != nil {
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

// GetOverwriteOnCommitRequest structure represents the OverwriteOnCommit operation request
type GetOverwriteOnCommitRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetOverwriteOnCommitRequest) xxx_ToOp(ctx context.Context, op *xxx_GetOverwriteOnCommitOperation) *xxx_GetOverwriteOnCommitOperation {
	if op == nil {
		op = &xxx_GetOverwriteOnCommitOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetOverwriteOnCommitRequest) xxx_FromOp(ctx context.Context, op *xxx_GetOverwriteOnCommitOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetOverwriteOnCommitRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetOverwriteOnCommitRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOverwriteOnCommitOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetOverwriteOnCommitResponse structure represents the OverwriteOnCommit operation response
type GetOverwriteOnCommitResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Overwrite int16          `idl:"name:overwrite" json:"overwrite"`
	// Return: The OverwriteOnCommit return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetOverwriteOnCommitResponse) xxx_ToOp(ctx context.Context, op *xxx_GetOverwriteOnCommitOperation) *xxx_GetOverwriteOnCommitOperation {
	if op == nil {
		op = &xxx_GetOverwriteOnCommitOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Overwrite = o.Overwrite
	op.Return = o.Return
	return op
}

func (o *GetOverwriteOnCommitResponse) xxx_FromOp(ctx context.Context, op *xxx_GetOverwriteOnCommitOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Overwrite = op.Overwrite
	o.Return = op.Return
}
func (o *GetOverwriteOnCommitResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetOverwriteOnCommitResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOverwriteOnCommitOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetOverwriteOnCommitOperation structure represents the OverwriteOnCommit operation
type xxx_SetOverwriteOnCommitOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Overwrite int16          `idl:"name:overwrite" json:"overwrite"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetOverwriteOnCommitOperation) OpNum() int { return 27 }

func (o *xxx_SetOverwriteOnCommitOperation) OpName() string {
	return "/IFsrmQuotaTemplateImported/v0/OverwriteOnCommit"
}

func (o *xxx_SetOverwriteOnCommitOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOverwriteOnCommitOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// overwrite {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Overwrite); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOverwriteOnCommitOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// overwrite {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Overwrite); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOverwriteOnCommitOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOverwriteOnCommitOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetOverwriteOnCommitOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetOverwriteOnCommitRequest structure represents the OverwriteOnCommit operation request
type SetOverwriteOnCommitRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	Overwrite int16          `idl:"name:overwrite" json:"overwrite"`
}

func (o *SetOverwriteOnCommitRequest) xxx_ToOp(ctx context.Context, op *xxx_SetOverwriteOnCommitOperation) *xxx_SetOverwriteOnCommitOperation {
	if op == nil {
		op = &xxx_SetOverwriteOnCommitOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Overwrite = o.Overwrite
	return op
}

func (o *SetOverwriteOnCommitRequest) xxx_FromOp(ctx context.Context, op *xxx_SetOverwriteOnCommitOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Overwrite = op.Overwrite
}
func (o *SetOverwriteOnCommitRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetOverwriteOnCommitRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOverwriteOnCommitOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetOverwriteOnCommitResponse structure represents the OverwriteOnCommit operation response
type SetOverwriteOnCommitResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The OverwriteOnCommit return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetOverwriteOnCommitResponse) xxx_ToOp(ctx context.Context, op *xxx_SetOverwriteOnCommitOperation) *xxx_SetOverwriteOnCommitOperation {
	if op == nil {
		op = &xxx_SetOverwriteOnCommitOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetOverwriteOnCommitResponse) xxx_FromOp(ctx context.Context, op *xxx_SetOverwriteOnCommitOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetOverwriteOnCommitResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetOverwriteOnCommitResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOverwriteOnCommitOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
