package iapphostwritableadminmanager

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iapphostadminmanager "github.com/oiweiwei/go-msrpc/msrpc/dcom/iisa/iapphostadminmanager/v0"
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
	_ = iapphostadminmanager.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/iisa"
)

var (
	// IAppHostWritableAdminManager interface identifier fa7660f6-7b3f-4237-a8bf-ed0ad0dcbbd9
	AppHostWritableAdminManagerIID = &dcom.IID{Data1: 0xfa7660f6, Data2: 0x7b3f, Data3: 0x4237, Data4: []byte{0xa8, 0xbf, 0xed, 0x0a, 0xd0, 0xdc, 0xbb, 0xd9}}
	// Syntax UUID
	AppHostWritableAdminManagerSyntaxUUID = &uuid.UUID{TimeLow: 0xfa7660f6, TimeMid: 0x7b3f, TimeHiAndVersion: 0x4237, ClockSeqHiAndReserved: 0xa8, ClockSeqLow: 0xbf, Node: [6]uint8{0xed, 0xa, 0xd0, 0xdc, 0xbb, 0xd9}}
	// Syntax ID
	AppHostWritableAdminManagerSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AppHostWritableAdminManagerSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IAppHostWritableAdminManager interface.
type AppHostWritableAdminManagerClient interface {

	// IAppHostAdminManager retrieval method.
	AppHostAdminManager() iapphostadminmanager.AppHostAdminManagerClient

	// CommitChanges operation.
	CommitChanges(context.Context, *CommitChangesRequest, ...dcerpc.CallOption) (*CommitChangesResponse, error)

	// CommitPath operation.
	GetCommitPath(context.Context, *GetCommitPathRequest, ...dcerpc.CallOption) (*GetCommitPathResponse, error)

	// CommitPath operation.
	SetCommitPath(context.Context, *SetCommitPathRequest, ...dcerpc.CallOption) (*SetCommitPathResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AppHostWritableAdminManagerClient
}

type xxx_DefaultAppHostWritableAdminManagerClient struct {
	iapphostadminmanager.AppHostAdminManagerClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAppHostWritableAdminManagerClient) AppHostAdminManager() iapphostadminmanager.AppHostAdminManagerClient {
	return o.AppHostAdminManagerClient
}

func (o *xxx_DefaultAppHostWritableAdminManagerClient) CommitChanges(ctx context.Context, in *CommitChangesRequest, opts ...dcerpc.CallOption) (*CommitChangesResponse, error) {
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
	out := &CommitChangesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostWritableAdminManagerClient) GetCommitPath(ctx context.Context, in *GetCommitPathRequest, opts ...dcerpc.CallOption) (*GetCommitPathResponse, error) {
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
	out := &GetCommitPathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostWritableAdminManagerClient) SetCommitPath(ctx context.Context, in *SetCommitPathRequest, opts ...dcerpc.CallOption) (*SetCommitPathResponse, error) {
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
	out := &SetCommitPathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAppHostWritableAdminManagerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAppHostWritableAdminManagerClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAppHostWritableAdminManagerClient) IPID(ctx context.Context, ipid *dcom.IPID) AppHostWritableAdminManagerClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAppHostWritableAdminManagerClient{
		AppHostAdminManagerClient: o.AppHostAdminManagerClient.IPID(ctx, ipid),
		cc:                        o.cc,
		ipid:                      ipid,
	}
}

func NewAppHostWritableAdminManagerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AppHostWritableAdminManagerClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AppHostWritableAdminManagerSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iapphostadminmanager.NewAppHostAdminManagerClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultAppHostWritableAdminManagerClient{
		AppHostAdminManagerClient: base,
		cc:                        cc,
		ipid:                      ipid,
	}, nil
}

// xxx_CommitChangesOperation structure represents the CommitChanges operation
type xxx_CommitChangesOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CommitChangesOperation) OpNum() int { return 7 }

func (o *xxx_CommitChangesOperation) OpName() string {
	return "/IAppHostWritableAdminManager/v0/CommitChanges"
}

func (o *xxx_CommitChangesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitChangesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CommitChangesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CommitChangesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CommitChangesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CommitChangesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CommitChangesRequest structure represents the CommitChanges operation request
type CommitChangesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CommitChangesRequest) xxx_ToOp(ctx context.Context, op *xxx_CommitChangesOperation) *xxx_CommitChangesOperation {
	if op == nil {
		op = &xxx_CommitChangesOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *CommitChangesRequest) xxx_FromOp(ctx context.Context, op *xxx_CommitChangesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CommitChangesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CommitChangesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CommitChangesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CommitChangesResponse structure represents the CommitChanges operation response
type CommitChangesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CommitChanges return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CommitChangesResponse) xxx_ToOp(ctx context.Context, op *xxx_CommitChangesOperation) *xxx_CommitChangesOperation {
	if op == nil {
		op = &xxx_CommitChangesOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *CommitChangesResponse) xxx_FromOp(ctx context.Context, op *xxx_CommitChangesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CommitChangesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CommitChangesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CommitChangesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetCommitPathOperation structure represents the CommitPath operation
type xxx_GetCommitPathOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	CommitPath *oaut.String   `idl:"name:pbstrCommitPath" json:"commit_path"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCommitPathOperation) OpNum() int { return 8 }

func (o *xxx_GetCommitPathOperation) OpName() string {
	return "/IAppHostWritableAdminManager/v0/CommitPath"
}

func (o *xxx_GetCommitPathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCommitPathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetCommitPathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetCommitPathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCommitPathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrCommitPath {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CommitPath != nil {
			_ptr_pbstrCommitPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CommitPath != nil {
					if err := o.CommitPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CommitPath, _ptr_pbstrCommitPath); err != nil {
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

func (o *xxx_GetCommitPathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrCommitPath {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrCommitPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CommitPath == nil {
				o.CommitPath = &oaut.String{}
			}
			if err := o.CommitPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrCommitPath := func(ptr interface{}) { o.CommitPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.CommitPath, _s_pbstrCommitPath, _ptr_pbstrCommitPath); err != nil {
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

// GetCommitPathRequest structure represents the CommitPath operation request
type GetCommitPathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetCommitPathRequest) xxx_ToOp(ctx context.Context, op *xxx_GetCommitPathOperation) *xxx_GetCommitPathOperation {
	if op == nil {
		op = &xxx_GetCommitPathOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetCommitPathRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCommitPathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCommitPathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetCommitPathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCommitPathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCommitPathResponse structure represents the CommitPath operation response
type GetCommitPathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	CommitPath *oaut.String   `idl:"name:pbstrCommitPath" json:"commit_path"`
	// Return: The CommitPath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCommitPathResponse) xxx_ToOp(ctx context.Context, op *xxx_GetCommitPathOperation) *xxx_GetCommitPathOperation {
	if op == nil {
		op = &xxx_GetCommitPathOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.CommitPath = op.CommitPath
	o.Return = op.Return
	return op
}

func (o *GetCommitPathResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCommitPathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.CommitPath = op.CommitPath
	o.Return = op.Return
}
func (o *GetCommitPathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCommitPathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCommitPathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetCommitPathOperation structure represents the CommitPath operation
type xxx_SetCommitPathOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	CommitPath *oaut.String   `idl:"name:bstrCommitPath" json:"commit_path"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetCommitPathOperation) OpNum() int { return 9 }

func (o *xxx_SetCommitPathOperation) OpName() string {
	return "/IAppHostWritableAdminManager/v0/CommitPath"
}

func (o *xxx_SetCommitPathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCommitPathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrCommitPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.CommitPath != nil {
			_ptr_bstrCommitPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.CommitPath != nil {
					if err := o.CommitPath.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.CommitPath, _ptr_bstrCommitPath); err != nil {
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

func (o *xxx_SetCommitPathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrCommitPath {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrCommitPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.CommitPath == nil {
				o.CommitPath = &oaut.String{}
			}
			if err := o.CommitPath.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrCommitPath := func(ptr interface{}) { o.CommitPath = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.CommitPath, _s_bstrCommitPath, _ptr_bstrCommitPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCommitPathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCommitPathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetCommitPathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetCommitPathRequest structure represents the CommitPath operation request
type SetCommitPathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	CommitPath *oaut.String   `idl:"name:bstrCommitPath" json:"commit_path"`
}

func (o *SetCommitPathRequest) xxx_ToOp(ctx context.Context, op *xxx_SetCommitPathOperation) *xxx_SetCommitPathOperation {
	if op == nil {
		op = &xxx_SetCommitPathOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.CommitPath = op.CommitPath
	return op
}

func (o *SetCommitPathRequest) xxx_FromOp(ctx context.Context, op *xxx_SetCommitPathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.CommitPath = op.CommitPath
}
func (o *SetCommitPathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetCommitPathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCommitPathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetCommitPathResponse structure represents the CommitPath operation response
type SetCommitPathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CommitPath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetCommitPathResponse) xxx_ToOp(ctx context.Context, op *xxx_SetCommitPathOperation) *xxx_SetCommitPathOperation {
	if op == nil {
		op = &xxx_SetCommitPathOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *SetCommitPathResponse) xxx_FromOp(ctx context.Context, op *xxx_SetCommitPathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetCommitPathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetCommitPathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCommitPathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
