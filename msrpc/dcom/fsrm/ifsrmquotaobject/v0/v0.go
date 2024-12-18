package ifsrmquotaobject

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmQuotaObject interface identifier 42dc3511-61d5-48ae-b6dc-59fc00c0a8d6
	QuotaObjectIID = &dcom.IID{Data1: 0x42dc3511, Data2: 0x61d5, Data3: 0x48ae, Data4: []byte{0xb6, 0xdc, 0x59, 0xfc, 0x00, 0xc0, 0xa8, 0xd6}}
	// Syntax UUID
	QuotaObjectSyntaxUUID = &uuid.UUID{TimeLow: 0x42dc3511, TimeMid: 0x61d5, TimeHiAndVersion: 0x48ae, ClockSeqHiAndReserved: 0xb6, ClockSeqLow: 0xdc, Node: [6]uint8{0x59, 0xfc, 0x0, 0xc0, 0xa8, 0xd6}}
	// Syntax ID
	QuotaObjectSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: QuotaObjectSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmQuotaObject interface.
type QuotaObjectClient interface {

	// IFsrmQuotaBase retrieval method.
	QuotaBase() ifsrmquotabase.QuotaBaseClient

	// Path operation.
	GetPath(context.Context, *GetPathRequest, ...dcerpc.CallOption) (*GetPathResponse, error)

	// UserSid operation.
	GetUserSID(context.Context, *GetUserSIDRequest, ...dcerpc.CallOption) (*GetUserSIDResponse, error)

	// UserAccount operation.
	GetUserAccount(context.Context, *GetUserAccountRequest, ...dcerpc.CallOption) (*GetUserAccountResponse, error)

	// SourceTemplateName operation.
	GetSourceTemplateName(context.Context, *GetSourceTemplateNameRequest, ...dcerpc.CallOption) (*GetSourceTemplateNameResponse, error)

	// MatchesSourceTemplate operation.
	GetMatchesSourceTemplate(context.Context, *GetMatchesSourceTemplateRequest, ...dcerpc.CallOption) (*GetMatchesSourceTemplateResponse, error)

	// ApplyTemplate operation.
	ApplyTemplate(context.Context, *ApplyTemplateRequest, ...dcerpc.CallOption) (*ApplyTemplateResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) QuotaObjectClient
}

type xxx_DefaultQuotaObjectClient struct {
	ifsrmquotabase.QuotaBaseClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultQuotaObjectClient) QuotaBase() ifsrmquotabase.QuotaBaseClient {
	return o.QuotaBaseClient
}

func (o *xxx_DefaultQuotaObjectClient) GetPath(ctx context.Context, in *GetPathRequest, opts ...dcerpc.CallOption) (*GetPathResponse, error) {
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
	out := &GetPathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaObjectClient) GetUserSID(ctx context.Context, in *GetUserSIDRequest, opts ...dcerpc.CallOption) (*GetUserSIDResponse, error) {
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
	out := &GetUserSIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaObjectClient) GetUserAccount(ctx context.Context, in *GetUserAccountRequest, opts ...dcerpc.CallOption) (*GetUserAccountResponse, error) {
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
	out := &GetUserAccountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaObjectClient) GetSourceTemplateName(ctx context.Context, in *GetSourceTemplateNameRequest, opts ...dcerpc.CallOption) (*GetSourceTemplateNameResponse, error) {
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
	out := &GetSourceTemplateNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaObjectClient) GetMatchesSourceTemplate(ctx context.Context, in *GetMatchesSourceTemplateRequest, opts ...dcerpc.CallOption) (*GetMatchesSourceTemplateResponse, error) {
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
	out := &GetMatchesSourceTemplateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaObjectClient) ApplyTemplate(ctx context.Context, in *ApplyTemplateRequest, opts ...dcerpc.CallOption) (*ApplyTemplateResponse, error) {
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
	out := &ApplyTemplateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQuotaObjectClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultQuotaObjectClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultQuotaObjectClient) IPID(ctx context.Context, ipid *dcom.IPID) QuotaObjectClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultQuotaObjectClient{
		QuotaBaseClient: o.QuotaBaseClient.IPID(ctx, ipid),
		cc:              o.cc,
		ipid:            ipid,
	}
}

func NewQuotaObjectClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (QuotaObjectClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(QuotaObjectSyntaxV0_0))...)
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
	return &xxx_DefaultQuotaObjectClient{
		QuotaBaseClient: base,
		cc:              cc,
		ipid:            ipid,
	}, nil
}

// xxx_GetPathOperation structure represents the Path operation
type xxx_GetPathOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path   *oaut.String   `idl:"name:path" json:"path"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPathOperation) OpNum() int { return 22 }

func (o *xxx_GetPathOperation) OpName() string { return "/IFsrmQuotaObject/v0/Path" }

func (o *xxx_GetPathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// path {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Path != nil {
			_ptr_path := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Path != nil {
					if err := o.Path.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Path, _ptr_path); err != nil {
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

func (o *xxx_GetPathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// path {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_path := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Path == nil {
				o.Path = &oaut.String{}
			}
			if err := o.Path.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_path := func(ptr interface{}) { o.Path = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Path, _s_path, _ptr_path); err != nil {
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

// GetPathRequest structure represents the Path operation request
type GetPathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPathRequest) xxx_ToOp(ctx context.Context) *xxx_GetPathOperation {
	if o == nil {
		return &xxx_GetPathOperation{}
	}
	return &xxx_GetPathOperation{
		This: o.This,
	}
}

func (o *GetPathRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetPathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPathResponse structure represents the Path operation response
type GetPathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path *oaut.String   `idl:"name:path" json:"path"`
	// Return: The Path return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPathResponse) xxx_ToOp(ctx context.Context) *xxx_GetPathOperation {
	if o == nil {
		return &xxx_GetPathOperation{}
	}
	return &xxx_GetPathOperation{
		That:   o.That,
		Path:   o.Path,
		Return: o.Return,
	}
}

func (o *GetPathResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Path = op.Path
	o.Return = op.Return
}
func (o *GetPathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetPathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetUserSIDOperation structure represents the UserSid operation
type xxx_GetUserSIDOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	UserSID *oaut.String   `idl:"name:userSid" json:"user_sid"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetUserSIDOperation) OpNum() int { return 23 }

func (o *xxx_GetUserSIDOperation) OpName() string { return "/IFsrmQuotaObject/v0/UserSid" }

func (o *xxx_GetUserSIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserSIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetUserSIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetUserSIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserSIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// userSid {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.UserSID != nil {
			_ptr_userSid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.UserSID != nil {
					if err := o.UserSID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.UserSID, _ptr_userSid); err != nil {
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

func (o *xxx_GetUserSIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// userSid {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_userSid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.UserSID == nil {
				o.UserSID = &oaut.String{}
			}
			if err := o.UserSID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_userSid := func(ptr interface{}) { o.UserSID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.UserSID, _s_userSid, _ptr_userSid); err != nil {
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

// GetUserSIDRequest structure represents the UserSid operation request
type GetUserSIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetUserSIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetUserSIDOperation {
	if o == nil {
		return &xxx_GetUserSIDOperation{}
	}
	return &xxx_GetUserSIDOperation{
		This: o.This,
	}
}

func (o *GetUserSIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetUserSIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetUserSIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetUserSIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUserSIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetUserSIDResponse structure represents the UserSid operation response
type GetUserSIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	UserSID *oaut.String   `idl:"name:userSid" json:"user_sid"`
	// Return: The UserSid return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetUserSIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetUserSIDOperation {
	if o == nil {
		return &xxx_GetUserSIDOperation{}
	}
	return &xxx_GetUserSIDOperation{
		That:    o.That,
		UserSID: o.UserSID,
		Return:  o.Return,
	}
}

func (o *GetUserSIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetUserSIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.UserSID = op.UserSID
	o.Return = op.Return
}
func (o *GetUserSIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetUserSIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUserSIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetUserAccountOperation structure represents the UserAccount operation
type xxx_GetUserAccountOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	UserAccount *oaut.String   `idl:"name:userAccount" json:"user_account"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetUserAccountOperation) OpNum() int { return 24 }

func (o *xxx_GetUserAccountOperation) OpName() string { return "/IFsrmQuotaObject/v0/UserAccount" }

func (o *xxx_GetUserAccountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserAccountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetUserAccountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetUserAccountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUserAccountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// userAccount {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.UserAccount != nil {
			_ptr_userAccount := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.UserAccount != nil {
					if err := o.UserAccount.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.UserAccount, _ptr_userAccount); err != nil {
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

func (o *xxx_GetUserAccountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// userAccount {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_userAccount := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.UserAccount == nil {
				o.UserAccount = &oaut.String{}
			}
			if err := o.UserAccount.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_userAccount := func(ptr interface{}) { o.UserAccount = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.UserAccount, _s_userAccount, _ptr_userAccount); err != nil {
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

// GetUserAccountRequest structure represents the UserAccount operation request
type GetUserAccountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetUserAccountRequest) xxx_ToOp(ctx context.Context) *xxx_GetUserAccountOperation {
	if o == nil {
		return &xxx_GetUserAccountOperation{}
	}
	return &xxx_GetUserAccountOperation{
		This: o.This,
	}
}

func (o *GetUserAccountRequest) xxx_FromOp(ctx context.Context, op *xxx_GetUserAccountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetUserAccountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetUserAccountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUserAccountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetUserAccountResponse structure represents the UserAccount operation response
type GetUserAccountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	UserAccount *oaut.String   `idl:"name:userAccount" json:"user_account"`
	// Return: The UserAccount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetUserAccountResponse) xxx_ToOp(ctx context.Context) *xxx_GetUserAccountOperation {
	if o == nil {
		return &xxx_GetUserAccountOperation{}
	}
	return &xxx_GetUserAccountOperation{
		That:        o.That,
		UserAccount: o.UserAccount,
		Return:      o.Return,
	}
}

func (o *GetUserAccountResponse) xxx_FromOp(ctx context.Context, op *xxx_GetUserAccountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.UserAccount = op.UserAccount
	o.Return = op.Return
}
func (o *GetUserAccountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetUserAccountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUserAccountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSourceTemplateNameOperation structure represents the SourceTemplateName operation
type xxx_GetSourceTemplateNameOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	QuotaTemplateName *oaut.String   `idl:"name:quotaTemplateName" json:"quota_template_name"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSourceTemplateNameOperation) OpNum() int { return 25 }

func (o *xxx_GetSourceTemplateNameOperation) OpName() string {
	return "/IFsrmQuotaObject/v0/SourceTemplateName"
}

func (o *xxx_GetSourceTemplateNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSourceTemplateNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetSourceTemplateNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetSourceTemplateNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSourceTemplateNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// quotaTemplateName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSourceTemplateNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// quotaTemplateName {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetSourceTemplateNameRequest structure represents the SourceTemplateName operation request
type GetSourceTemplateNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetSourceTemplateNameRequest) xxx_ToOp(ctx context.Context) *xxx_GetSourceTemplateNameOperation {
	if o == nil {
		return &xxx_GetSourceTemplateNameOperation{}
	}
	return &xxx_GetSourceTemplateNameOperation{
		This: o.This,
	}
}

func (o *GetSourceTemplateNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSourceTemplateNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetSourceTemplateNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetSourceTemplateNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSourceTemplateNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSourceTemplateNameResponse structure represents the SourceTemplateName operation response
type GetSourceTemplateNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	QuotaTemplateName *oaut.String   `idl:"name:quotaTemplateName" json:"quota_template_name"`
	// Return: The SourceTemplateName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetSourceTemplateNameResponse) xxx_ToOp(ctx context.Context) *xxx_GetSourceTemplateNameOperation {
	if o == nil {
		return &xxx_GetSourceTemplateNameOperation{}
	}
	return &xxx_GetSourceTemplateNameOperation{
		That:              o.That,
		QuotaTemplateName: o.QuotaTemplateName,
		Return:            o.Return,
	}
}

func (o *GetSourceTemplateNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSourceTemplateNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.QuotaTemplateName = op.QuotaTemplateName
	o.Return = op.Return
}
func (o *GetSourceTemplateNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetSourceTemplateNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSourceTemplateNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMatchesSourceTemplateOperation structure represents the MatchesSourceTemplate operation
type xxx_GetMatchesSourceTemplateOperation struct {
	This    *dcom.ORPCThis `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Matches int16          `idl:"name:matches" json:"matches"`
	Return  int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMatchesSourceTemplateOperation) OpNum() int { return 26 }

func (o *xxx_GetMatchesSourceTemplateOperation) OpName() string {
	return "/IFsrmQuotaObject/v0/MatchesSourceTemplate"
}

func (o *xxx_GetMatchesSourceTemplateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMatchesSourceTemplateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetMatchesSourceTemplateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetMatchesSourceTemplateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMatchesSourceTemplateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// matches {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Matches); err != nil {
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

func (o *xxx_GetMatchesSourceTemplateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// matches {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Matches); err != nil {
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

// GetMatchesSourceTemplateRequest structure represents the MatchesSourceTemplate operation request
type GetMatchesSourceTemplateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetMatchesSourceTemplateRequest) xxx_ToOp(ctx context.Context) *xxx_GetMatchesSourceTemplateOperation {
	if o == nil {
		return &xxx_GetMatchesSourceTemplateOperation{}
	}
	return &xxx_GetMatchesSourceTemplateOperation{
		This: o.This,
	}
}

func (o *GetMatchesSourceTemplateRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMatchesSourceTemplateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetMatchesSourceTemplateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetMatchesSourceTemplateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMatchesSourceTemplateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMatchesSourceTemplateResponse structure represents the MatchesSourceTemplate operation response
type GetMatchesSourceTemplateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Matches int16          `idl:"name:matches" json:"matches"`
	// Return: The MatchesSourceTemplate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMatchesSourceTemplateResponse) xxx_ToOp(ctx context.Context) *xxx_GetMatchesSourceTemplateOperation {
	if o == nil {
		return &xxx_GetMatchesSourceTemplateOperation{}
	}
	return &xxx_GetMatchesSourceTemplateOperation{
		That:    o.That,
		Matches: o.Matches,
		Return:  o.Return,
	}
}

func (o *GetMatchesSourceTemplateResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMatchesSourceTemplateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Matches = op.Matches
	o.Return = op.Return
}
func (o *GetMatchesSourceTemplateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetMatchesSourceTemplateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMatchesSourceTemplateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ApplyTemplateOperation structure represents the ApplyTemplate operation
type xxx_ApplyTemplateOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	QuotaTemplateName *oaut.String   `idl:"name:quotaTemplateName" json:"quota_template_name"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ApplyTemplateOperation) OpNum() int { return 27 }

func (o *xxx_ApplyTemplateOperation) OpName() string { return "/IFsrmQuotaObject/v0/ApplyTemplate" }

func (o *xxx_ApplyTemplateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ApplyTemplateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ApplyTemplateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_ApplyTemplateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ApplyTemplateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ApplyTemplateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ApplyTemplateRequest structure represents the ApplyTemplate operation request
type ApplyTemplateRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	QuotaTemplateName *oaut.String   `idl:"name:quotaTemplateName" json:"quota_template_name"`
}

func (o *ApplyTemplateRequest) xxx_ToOp(ctx context.Context) *xxx_ApplyTemplateOperation {
	if o == nil {
		return &xxx_ApplyTemplateOperation{}
	}
	return &xxx_ApplyTemplateOperation{
		This:              o.This,
		QuotaTemplateName: o.QuotaTemplateName,
	}
}

func (o *ApplyTemplateRequest) xxx_FromOp(ctx context.Context, op *xxx_ApplyTemplateOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.QuotaTemplateName = op.QuotaTemplateName
}
func (o *ApplyTemplateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ApplyTemplateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ApplyTemplateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ApplyTemplateResponse structure represents the ApplyTemplate operation response
type ApplyTemplateResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ApplyTemplate return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ApplyTemplateResponse) xxx_ToOp(ctx context.Context) *xxx_ApplyTemplateOperation {
	if o == nil {
		return &xxx_ApplyTemplateOperation{}
	}
	return &xxx_ApplyTemplateOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ApplyTemplateResponse) xxx_FromOp(ctx context.Context, op *xxx_ApplyTemplateOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ApplyTemplateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ApplyTemplateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ApplyTemplateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
