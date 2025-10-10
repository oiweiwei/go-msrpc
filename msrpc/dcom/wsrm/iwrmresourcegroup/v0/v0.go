package iwrmresourcegroup

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
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
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wsrm"
)

var (
	// IWRMResourceGroup interface identifier bc681469-9dd9-4bf4-9b3d-709f69efe431
	IwrmResourceGroupIID = &dcom.IID{Data1: 0xbc681469, Data2: 0x9dd9, Data3: 0x4bf4, Data4: []byte{0x9b, 0x3d, 0x70, 0x9f, 0x69, 0xef, 0xe4, 0x31}}
	// Syntax UUID
	IwrmResourceGroupSyntaxUUID = &uuid.UUID{TimeLow: 0xbc681469, TimeMid: 0x9dd9, TimeHiAndVersion: 0x4bf4, ClockSeqHiAndReserved: 0x9b, ClockSeqLow: 0x3d, Node: [6]uint8{0x70, 0x9f, 0x69, 0xef, 0xe4, 0x31}}
	// Syntax ID
	IwrmResourceGroupSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: IwrmResourceGroupSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWRMResourceGroup interface.
type IwrmResourceGroupClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	GetResourceGroupInfo(context.Context, *GetResourceGroupInfoRequest, ...dcerpc.CallOption) (*GetResourceGroupInfoResponse, error)

	ModifyResourceGroup(context.Context, *ModifyResourceGroupRequest, ...dcerpc.CallOption) (*ModifyResourceGroupResponse, error)

	CreateResourceGroup(context.Context, *CreateResourceGroupRequest, ...dcerpc.CallOption) (*CreateResourceGroupResponse, error)

	DeleteResourceGroup(context.Context, *DeleteResourceGroupRequest, ...dcerpc.CallOption) (*DeleteResourceGroupResponse, error)

	RenameResourceGroup(context.Context, *RenameResourceGroupRequest, ...dcerpc.CallOption) (*RenameResourceGroupResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) IwrmResourceGroupClient
}

type xxx_DefaultIwrmResourceGroupClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultIwrmResourceGroupClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultIwrmResourceGroupClient) GetResourceGroupInfo(ctx context.Context, in *GetResourceGroupInfoRequest, opts ...dcerpc.CallOption) (*GetResourceGroupInfoResponse, error) {
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
	out := &GetResourceGroupInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIwrmResourceGroupClient) ModifyResourceGroup(ctx context.Context, in *ModifyResourceGroupRequest, opts ...dcerpc.CallOption) (*ModifyResourceGroupResponse, error) {
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
	out := &ModifyResourceGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIwrmResourceGroupClient) CreateResourceGroup(ctx context.Context, in *CreateResourceGroupRequest, opts ...dcerpc.CallOption) (*CreateResourceGroupResponse, error) {
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
	out := &CreateResourceGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIwrmResourceGroupClient) DeleteResourceGroup(ctx context.Context, in *DeleteResourceGroupRequest, opts ...dcerpc.CallOption) (*DeleteResourceGroupResponse, error) {
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
	out := &DeleteResourceGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIwrmResourceGroupClient) RenameResourceGroup(ctx context.Context, in *RenameResourceGroupRequest, opts ...dcerpc.CallOption) (*RenameResourceGroupResponse, error) {
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
	out := &RenameResourceGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIwrmResourceGroupClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultIwrmResourceGroupClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultIwrmResourceGroupClient) IPID(ctx context.Context, ipid *dcom.IPID) IwrmResourceGroupClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultIwrmResourceGroupClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewIwrmResourceGroupClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (IwrmResourceGroupClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(IwrmResourceGroupSyntaxV0_0))...)
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
	return &xxx_DefaultIwrmResourceGroupClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetResourceGroupInfoOperation structure represents the GetResourceGroupInfo operation
type xxx_GetResourceGroupInfoOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	ResourceGroupName *oaut.String   `idl:"name:bstrResourceGroupName" json:"resource_group_name"`
	ResourceGroupInfo *oaut.String   `idl:"name:pbstrResourceGroupInfo" json:"resource_group_info"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetResourceGroupInfoOperation) OpNum() int { return 7 }

func (o *xxx_GetResourceGroupInfoOperation) OpName() string {
	return "/IWRMResourceGroup/v0/GetResourceGroupInfo"
}

func (o *xxx_GetResourceGroupInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResourceGroupInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ResourceGroupName != nil {
			_ptr_bstrResourceGroupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResourceGroupName != nil {
					if err := o.ResourceGroupName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResourceGroupName, _ptr_bstrResourceGroupName); err != nil {
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

func (o *xxx_GetResourceGroupInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrResourceGroupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResourceGroupName == nil {
				o.ResourceGroupName = &oaut.String{}
			}
			if err := o.ResourceGroupName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrResourceGroupName := func(ptr interface{}) { o.ResourceGroupName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ResourceGroupName, _s_bstrResourceGroupName, _ptr_bstrResourceGroupName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResourceGroupInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResourceGroupInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrResourceGroupInfo {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ResourceGroupInfo != nil {
			_ptr_pbstrResourceGroupInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResourceGroupInfo != nil {
					if err := o.ResourceGroupInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResourceGroupInfo, _ptr_pbstrResourceGroupInfo); err != nil {
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

func (o *xxx_GetResourceGroupInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrResourceGroupInfo {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrResourceGroupInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResourceGroupInfo == nil {
				o.ResourceGroupInfo = &oaut.String{}
			}
			if err := o.ResourceGroupInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrResourceGroupInfo := func(ptr interface{}) { o.ResourceGroupInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ResourceGroupInfo, _s_pbstrResourceGroupInfo, _ptr_pbstrResourceGroupInfo); err != nil {
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

// GetResourceGroupInfoRequest structure represents the GetResourceGroupInfo operation request
type GetResourceGroupInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	ResourceGroupName *oaut.String   `idl:"name:bstrResourceGroupName" json:"resource_group_name"`
}

func (o *GetResourceGroupInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetResourceGroupInfoOperation) *xxx_GetResourceGroupInfoOperation {
	if op == nil {
		op = &xxx_GetResourceGroupInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ResourceGroupName = o.ResourceGroupName
	return op
}

func (o *GetResourceGroupInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetResourceGroupInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ResourceGroupName = op.ResourceGroupName
}
func (o *GetResourceGroupInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetResourceGroupInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResourceGroupInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetResourceGroupInfoResponse structure represents the GetResourceGroupInfo operation response
type GetResourceGroupInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	ResourceGroupInfo *oaut.String   `idl:"name:pbstrResourceGroupInfo" json:"resource_group_info"`
	// Return: The GetResourceGroupInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetResourceGroupInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetResourceGroupInfoOperation) *xxx_GetResourceGroupInfoOperation {
	if op == nil {
		op = &xxx_GetResourceGroupInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ResourceGroupInfo = o.ResourceGroupInfo
	op.Return = o.Return
	return op
}

func (o *GetResourceGroupInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetResourceGroupInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ResourceGroupInfo = op.ResourceGroupInfo
	o.Return = op.Return
}
func (o *GetResourceGroupInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetResourceGroupInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResourceGroupInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ModifyResourceGroupOperation structure represents the ModifyResourceGroup operation
type xxx_ModifyResourceGroupOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	ResourceGroupInfo *oaut.String   `idl:"name:bstrResourceGroupInfo" json:"resource_group_info"`
	Overwrite         bool           `idl:"name:bOverwrite" json:"overwrite"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ModifyResourceGroupOperation) OpNum() int { return 8 }

func (o *xxx_ModifyResourceGroupOperation) OpName() string {
	return "/IWRMResourceGroup/v0/ModifyResourceGroup"
}

func (o *xxx_ModifyResourceGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyResourceGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrResourceGroupInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ResourceGroupInfo != nil {
			_ptr_bstrResourceGroupInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResourceGroupInfo != nil {
					if err := o.ResourceGroupInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResourceGroupInfo, _ptr_bstrResourceGroupInfo); err != nil {
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
	// bOverwrite {in} (1:{alias=BOOL}(int32))
	{
		if !o.Overwrite {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ModifyResourceGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrResourceGroupInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrResourceGroupInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResourceGroupInfo == nil {
				o.ResourceGroupInfo = &oaut.String{}
			}
			if err := o.ResourceGroupInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrResourceGroupInfo := func(ptr interface{}) { o.ResourceGroupInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ResourceGroupInfo, _s_bstrResourceGroupInfo, _ptr_bstrResourceGroupInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bOverwrite {in} (1:{alias=BOOL}(int32))
	{
		var _bOverwrite int32
		if err := w.ReadData(&_bOverwrite); err != nil {
			return err
		}
		o.Overwrite = _bOverwrite != 0
	}
	return nil
}

func (o *xxx_ModifyResourceGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyResourceGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ModifyResourceGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ModifyResourceGroupRequest structure represents the ModifyResourceGroup operation request
type ModifyResourceGroupRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	ResourceGroupInfo *oaut.String   `idl:"name:bstrResourceGroupInfo" json:"resource_group_info"`
	Overwrite         bool           `idl:"name:bOverwrite" json:"overwrite"`
}

func (o *ModifyResourceGroupRequest) xxx_ToOp(ctx context.Context, op *xxx_ModifyResourceGroupOperation) *xxx_ModifyResourceGroupOperation {
	if op == nil {
		op = &xxx_ModifyResourceGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ResourceGroupInfo = o.ResourceGroupInfo
	op.Overwrite = o.Overwrite
	return op
}

func (o *ModifyResourceGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_ModifyResourceGroupOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ResourceGroupInfo = op.ResourceGroupInfo
	o.Overwrite = op.Overwrite
}
func (o *ModifyResourceGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ModifyResourceGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyResourceGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ModifyResourceGroupResponse structure represents the ModifyResourceGroup operation response
type ModifyResourceGroupResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ModifyResourceGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ModifyResourceGroupResponse) xxx_ToOp(ctx context.Context, op *xxx_ModifyResourceGroupOperation) *xxx_ModifyResourceGroupOperation {
	if op == nil {
		op = &xxx_ModifyResourceGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ModifyResourceGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_ModifyResourceGroupOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ModifyResourceGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ModifyResourceGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyResourceGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateResourceGroupOperation structure represents the CreateResourceGroup operation
type xxx_CreateResourceGroupOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	ResourceGroupInfo *oaut.String   `idl:"name:bstrResourceGroupInfo" json:"resource_group_info"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateResourceGroupOperation) OpNum() int { return 9 }

func (o *xxx_CreateResourceGroupOperation) OpName() string {
	return "/IWRMResourceGroup/v0/CreateResourceGroup"
}

func (o *xxx_CreateResourceGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateResourceGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrResourceGroupInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ResourceGroupInfo != nil {
			_ptr_bstrResourceGroupInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResourceGroupInfo != nil {
					if err := o.ResourceGroupInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResourceGroupInfo, _ptr_bstrResourceGroupInfo); err != nil {
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

func (o *xxx_CreateResourceGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrResourceGroupInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrResourceGroupInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResourceGroupInfo == nil {
				o.ResourceGroupInfo = &oaut.String{}
			}
			if err := o.ResourceGroupInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrResourceGroupInfo := func(ptr interface{}) { o.ResourceGroupInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ResourceGroupInfo, _s_bstrResourceGroupInfo, _ptr_bstrResourceGroupInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateResourceGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateResourceGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateResourceGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateResourceGroupRequest structure represents the CreateResourceGroup operation request
type CreateResourceGroupRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	ResourceGroupInfo *oaut.String   `idl:"name:bstrResourceGroupInfo" json:"resource_group_info"`
}

func (o *CreateResourceGroupRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateResourceGroupOperation) *xxx_CreateResourceGroupOperation {
	if op == nil {
		op = &xxx_CreateResourceGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ResourceGroupInfo = o.ResourceGroupInfo
	return op
}

func (o *CreateResourceGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateResourceGroupOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ResourceGroupInfo = op.ResourceGroupInfo
}
func (o *CreateResourceGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateResourceGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateResourceGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateResourceGroupResponse structure represents the CreateResourceGroup operation response
type CreateResourceGroupResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreateResourceGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateResourceGroupResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateResourceGroupOperation) *xxx_CreateResourceGroupOperation {
	if op == nil {
		op = &xxx_CreateResourceGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CreateResourceGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateResourceGroupOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateResourceGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateResourceGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateResourceGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteResourceGroupOperation structure represents the DeleteResourceGroup operation
type xxx_DeleteResourceGroupOperation struct {
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat `idl:"name:That" json:"that"`
	ResourceGroupName *oaut.String   `idl:"name:bstrResourceGroupName" json:"resource_group_name"`
	Return            int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteResourceGroupOperation) OpNum() int { return 10 }

func (o *xxx_DeleteResourceGroupOperation) OpName() string {
	return "/IWRMResourceGroup/v0/DeleteResourceGroup"
}

func (o *xxx_DeleteResourceGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteResourceGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ResourceGroupName != nil {
			_ptr_bstrResourceGroupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ResourceGroupName != nil {
					if err := o.ResourceGroupName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ResourceGroupName, _ptr_bstrResourceGroupName); err != nil {
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

func (o *xxx_DeleteResourceGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrResourceGroupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ResourceGroupName == nil {
				o.ResourceGroupName = &oaut.String{}
			}
			if err := o.ResourceGroupName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrResourceGroupName := func(ptr interface{}) { o.ResourceGroupName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ResourceGroupName, _s_bstrResourceGroupName, _ptr_bstrResourceGroupName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteResourceGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteResourceGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteResourceGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteResourceGroupRequest structure represents the DeleteResourceGroup operation request
type DeleteResourceGroupRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This              *dcom.ORPCThis `idl:"name:This" json:"this"`
	ResourceGroupName *oaut.String   `idl:"name:bstrResourceGroupName" json:"resource_group_name"`
}

func (o *DeleteResourceGroupRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteResourceGroupOperation) *xxx_DeleteResourceGroupOperation {
	if op == nil {
		op = &xxx_DeleteResourceGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ResourceGroupName = o.ResourceGroupName
	return op
}

func (o *DeleteResourceGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteResourceGroupOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ResourceGroupName = op.ResourceGroupName
}
func (o *DeleteResourceGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteResourceGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteResourceGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteResourceGroupResponse structure represents the DeleteResourceGroup operation response
type DeleteResourceGroupResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteResourceGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteResourceGroupResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteResourceGroupOperation) *xxx_DeleteResourceGroupOperation {
	if op == nil {
		op = &xxx_DeleteResourceGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteResourceGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteResourceGroupOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteResourceGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteResourceGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteResourceGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RenameResourceGroupOperation structure represents the RenameResourceGroup operation
type xxx_RenameResourceGroupOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	NewResourceGroupName *oaut.String   `idl:"name:bstrNewResourceGroupName" json:"new_resource_group_name"`
	OldResourceGroupName *oaut.String   `idl:"name:bstrOldResourceGroupName" json:"old_resource_group_name"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RenameResourceGroupOperation) OpNum() int { return 11 }

func (o *xxx_RenameResourceGroupOperation) OpName() string {
	return "/IWRMResourceGroup/v0/RenameResourceGroup"
}

func (o *xxx_RenameResourceGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameResourceGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrNewResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.NewResourceGroupName != nil {
			_ptr_bstrNewResourceGroupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NewResourceGroupName != nil {
					if err := o.NewResourceGroupName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NewResourceGroupName, _ptr_bstrNewResourceGroupName); err != nil {
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
	// bstrOldResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.OldResourceGroupName != nil {
			_ptr_bstrOldResourceGroupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OldResourceGroupName != nil {
					if err := o.OldResourceGroupName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OldResourceGroupName, _ptr_bstrOldResourceGroupName); err != nil {
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

func (o *xxx_RenameResourceGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrNewResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrNewResourceGroupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NewResourceGroupName == nil {
				o.NewResourceGroupName = &oaut.String{}
			}
			if err := o.NewResourceGroupName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrNewResourceGroupName := func(ptr interface{}) { o.NewResourceGroupName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.NewResourceGroupName, _s_bstrNewResourceGroupName, _ptr_bstrNewResourceGroupName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrOldResourceGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrOldResourceGroupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OldResourceGroupName == nil {
				o.OldResourceGroupName = &oaut.String{}
			}
			if err := o.OldResourceGroupName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrOldResourceGroupName := func(ptr interface{}) { o.OldResourceGroupName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.OldResourceGroupName, _s_bstrOldResourceGroupName, _ptr_bstrOldResourceGroupName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameResourceGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameResourceGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RenameResourceGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RenameResourceGroupRequest structure represents the RenameResourceGroup operation request
type RenameResourceGroupRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	NewResourceGroupName *oaut.String   `idl:"name:bstrNewResourceGroupName" json:"new_resource_group_name"`
	OldResourceGroupName *oaut.String   `idl:"name:bstrOldResourceGroupName" json:"old_resource_group_name"`
}

func (o *RenameResourceGroupRequest) xxx_ToOp(ctx context.Context, op *xxx_RenameResourceGroupOperation) *xxx_RenameResourceGroupOperation {
	if op == nil {
		op = &xxx_RenameResourceGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.NewResourceGroupName = o.NewResourceGroupName
	op.OldResourceGroupName = o.OldResourceGroupName
	return op
}

func (o *RenameResourceGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_RenameResourceGroupOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.NewResourceGroupName = op.NewResourceGroupName
	o.OldResourceGroupName = op.OldResourceGroupName
}
func (o *RenameResourceGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RenameResourceGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RenameResourceGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RenameResourceGroupResponse structure represents the RenameResourceGroup operation response
type RenameResourceGroupResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RenameResourceGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RenameResourceGroupResponse) xxx_ToOp(ctx context.Context, op *xxx_RenameResourceGroupOperation) *xxx_RenameResourceGroupOperation {
	if op == nil {
		op = &xxx_RenameResourceGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RenameResourceGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_RenameResourceGroupOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RenameResourceGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RenameResourceGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RenameResourceGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
