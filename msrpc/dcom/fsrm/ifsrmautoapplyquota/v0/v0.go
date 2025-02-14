package ifsrmautoapplyquota

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
	ifsrmquotaobject "github.com/oiweiwei/go-msrpc/msrpc/dcom/fsrm/ifsrmquotaobject/v0"
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
	_ = ifsrmquotaobject.GoPackage
	_ = oaut.GoPackage
	_ = fsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/fsrm"
)

var (
	// IFsrmAutoApplyQuota interface identifier f82e5729-6aba-4740-bfc7-c7f58f75fb7b
	AutoApplyQuotaIID = &dcom.IID{Data1: 0xf82e5729, Data2: 0x6aba, Data3: 0x4740, Data4: []byte{0xbf, 0xc7, 0xc7, 0xf5, 0x8f, 0x75, 0xfb, 0x7b}}
	// Syntax UUID
	AutoApplyQuotaSyntaxUUID = &uuid.UUID{TimeLow: 0xf82e5729, TimeMid: 0x6aba, TimeHiAndVersion: 0x4740, ClockSeqHiAndReserved: 0xbf, ClockSeqLow: 0xc7, Node: [6]uint8{0xc7, 0xf5, 0x8f, 0x75, 0xfb, 0x7b}}
	// Syntax ID
	AutoApplyQuotaSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AutoApplyQuotaSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IFsrmAutoApplyQuota interface.
type AutoApplyQuotaClient interface {

	// IFsrmQuotaObject retrieval method.
	QuotaObject() ifsrmquotaobject.QuotaObjectClient

	// ExcludeFolders operation.
	GetExcludeFolders(context.Context, *GetExcludeFoldersRequest, ...dcerpc.CallOption) (*GetExcludeFoldersResponse, error)

	// ExcludeFolders operation.
	SetExcludeFolders(context.Context, *SetExcludeFoldersRequest, ...dcerpc.CallOption) (*SetExcludeFoldersResponse, error)

	// CommitAndUpdateDerived operation.
	CommitAndUpdateDerived(context.Context, *CommitAndUpdateDerivedRequest, ...dcerpc.CallOption) (*CommitAndUpdateDerivedResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AutoApplyQuotaClient
}

type xxx_DefaultAutoApplyQuotaClient struct {
	ifsrmquotaobject.QuotaObjectClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAutoApplyQuotaClient) QuotaObject() ifsrmquotaobject.QuotaObjectClient {
	return o.QuotaObjectClient
}

func (o *xxx_DefaultAutoApplyQuotaClient) GetExcludeFolders(ctx context.Context, in *GetExcludeFoldersRequest, opts ...dcerpc.CallOption) (*GetExcludeFoldersResponse, error) {
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
	out := &GetExcludeFoldersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAutoApplyQuotaClient) SetExcludeFolders(ctx context.Context, in *SetExcludeFoldersRequest, opts ...dcerpc.CallOption) (*SetExcludeFoldersResponse, error) {
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
	out := &SetExcludeFoldersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAutoApplyQuotaClient) CommitAndUpdateDerived(ctx context.Context, in *CommitAndUpdateDerivedRequest, opts ...dcerpc.CallOption) (*CommitAndUpdateDerivedResponse, error) {
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

func (o *xxx_DefaultAutoApplyQuotaClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAutoApplyQuotaClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAutoApplyQuotaClient) IPID(ctx context.Context, ipid *dcom.IPID) AutoApplyQuotaClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAutoApplyQuotaClient{
		QuotaObjectClient: o.QuotaObjectClient.IPID(ctx, ipid),
		cc:                o.cc,
		ipid:              ipid,
	}
}

func NewAutoApplyQuotaClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AutoApplyQuotaClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AutoApplyQuotaSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := ifsrmquotaobject.NewQuotaObjectClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultAutoApplyQuotaClient{
		QuotaObjectClient: base,
		cc:                cc,
		ipid:              ipid,
	}, nil
}

// xxx_GetExcludeFoldersOperation structure represents the ExcludeFolders operation
type xxx_GetExcludeFoldersOperation struct {
	This    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Folders *oaut.SafeArray `idl:"name:folders" json:"folders"`
	Return  int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetExcludeFoldersOperation) OpNum() int { return 28 }

func (o *xxx_GetExcludeFoldersOperation) OpName() string {
	return "/IFsrmAutoApplyQuota/v0/ExcludeFolders"
}

func (o *xxx_GetExcludeFoldersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetExcludeFoldersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetExcludeFoldersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetExcludeFoldersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetExcludeFoldersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// folders {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Folders != nil {
			_ptr_folders := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Folders != nil {
					if err := o.Folders.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Folders, _ptr_folders); err != nil {
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

func (o *xxx_GetExcludeFoldersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// folders {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=SAFEARRAY}*(1))(3:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_folders := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Folders == nil {
				o.Folders = &oaut.SafeArray{}
			}
			if err := o.Folders.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_folders := func(ptr interface{}) { o.Folders = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Folders, _s_folders, _ptr_folders); err != nil {
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

// GetExcludeFoldersRequest structure represents the ExcludeFolders operation request
type GetExcludeFoldersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetExcludeFoldersRequest) xxx_ToOp(ctx context.Context, op *xxx_GetExcludeFoldersOperation) *xxx_GetExcludeFoldersOperation {
	if op == nil {
		op = &xxx_GetExcludeFoldersOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetExcludeFoldersRequest) xxx_FromOp(ctx context.Context, op *xxx_GetExcludeFoldersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetExcludeFoldersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetExcludeFoldersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetExcludeFoldersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetExcludeFoldersResponse structure represents the ExcludeFolders operation response
type GetExcludeFoldersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Folders *oaut.SafeArray `idl:"name:folders" json:"folders"`
	// Return: The ExcludeFolders return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetExcludeFoldersResponse) xxx_ToOp(ctx context.Context, op *xxx_GetExcludeFoldersOperation) *xxx_GetExcludeFoldersOperation {
	if op == nil {
		op = &xxx_GetExcludeFoldersOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Folders = op.Folders
	o.Return = op.Return
	return op
}

func (o *GetExcludeFoldersResponse) xxx_FromOp(ctx context.Context, op *xxx_GetExcludeFoldersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Folders = op.Folders
	o.Return = op.Return
}
func (o *GetExcludeFoldersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetExcludeFoldersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetExcludeFoldersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetExcludeFoldersOperation structure represents the ExcludeFolders operation
type xxx_SetExcludeFoldersOperation struct {
	This    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Folders *oaut.SafeArray `idl:"name:folders" json:"folders"`
	Return  int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_SetExcludeFoldersOperation) OpNum() int { return 29 }

func (o *xxx_SetExcludeFoldersOperation) OpName() string {
	return "/IFsrmAutoApplyQuota/v0/ExcludeFolders"
}

func (o *xxx_SetExcludeFoldersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExcludeFoldersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// folders {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		if o.Folders != nil {
			_ptr_folders := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Folders != nil {
					if err := o.Folders.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.SafeArray{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Folders, _ptr_folders); err != nil {
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

func (o *xxx_SetExcludeFoldersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// folders {in} (1:{pointer=unique, alias=SAFEARRAY}*(1))(2:{pointer=unique, alias=_SAFEARRAY}(struct))
	{
		_ptr_folders := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Folders == nil {
				o.Folders = &oaut.SafeArray{}
			}
			if err := o.Folders.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_folders := func(ptr interface{}) { o.Folders = *ptr.(**oaut.SafeArray) }
		if err := w.ReadPointer(&o.Folders, _s_folders, _ptr_folders); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExcludeFoldersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetExcludeFoldersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetExcludeFoldersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetExcludeFoldersRequest structure represents the ExcludeFolders operation request
type SetExcludeFoldersRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This    *dcom.ORPCThis  `idl:"name:This" json:"this"`
	Folders *oaut.SafeArray `idl:"name:folders" json:"folders"`
}

func (o *SetExcludeFoldersRequest) xxx_ToOp(ctx context.Context, op *xxx_SetExcludeFoldersOperation) *xxx_SetExcludeFoldersOperation {
	if op == nil {
		op = &xxx_SetExcludeFoldersOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Folders = op.Folders
	return op
}

func (o *SetExcludeFoldersRequest) xxx_FromOp(ctx context.Context, op *xxx_SetExcludeFoldersOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Folders = op.Folders
}
func (o *SetExcludeFoldersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetExcludeFoldersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetExcludeFoldersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetExcludeFoldersResponse structure represents the ExcludeFolders operation response
type SetExcludeFoldersResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ExcludeFolders return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetExcludeFoldersResponse) xxx_ToOp(ctx context.Context, op *xxx_SetExcludeFoldersOperation) *xxx_SetExcludeFoldersOperation {
	if op == nil {
		op = &xxx_SetExcludeFoldersOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *SetExcludeFoldersResponse) xxx_FromOp(ctx context.Context, op *xxx_SetExcludeFoldersOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetExcludeFoldersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetExcludeFoldersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetExcludeFoldersOperation{}
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

func (o *xxx_CommitAndUpdateDerivedOperation) OpNum() int { return 30 }

func (o *xxx_CommitAndUpdateDerivedOperation) OpName() string {
	return "/IFsrmAutoApplyQuota/v0/CommitAndUpdateDerived"
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
	o.This = op.This
	o.CommitOptions = op.CommitOptions
	o.ApplyOptions = op.ApplyOptions
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
	o.That = op.That
	o.DerivedObjectsResult = op.DerivedObjectsResult
	o.Return = op.Return
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
