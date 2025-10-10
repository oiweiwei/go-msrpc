package iwrmmachinegroup

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
	wsrm "github.com/oiweiwei/go-msrpc/msrpc/dcom/wsrm"
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
	_ = wsrm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wsrm"
)

var (
	// IWRMMachineGroup interface identifier 943991a5-b3fe-41fa-9696-7f7b656ee34b
	IwrmMachineGroupIID = &dcom.IID{Data1: 0x943991a5, Data2: 0xb3fe, Data3: 0x41fa, Data4: []byte{0x96, 0x96, 0x7f, 0x7b, 0x65, 0x6e, 0xe3, 0x4b}}
	// Syntax UUID
	IwrmMachineGroupSyntaxUUID = &uuid.UUID{TimeLow: 0x943991a5, TimeMid: 0xb3fe, TimeHiAndVersion: 0x41fa, ClockSeqHiAndReserved: 0x96, ClockSeqLow: 0x96, Node: [6]uint8{0x7f, 0x7b, 0x65, 0x6e, 0xe3, 0x4b}}
	// Syntax ID
	IwrmMachineGroupSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: IwrmMachineGroupSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWRMMachineGroup interface.
type IwrmMachineGroupClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	CreateMachineGroup(context.Context, *CreateMachineGroupRequest, ...dcerpc.CallOption) (*CreateMachineGroupResponse, error)

	GetMachineGroupInfo(context.Context, *GetMachineGroupInfoRequest, ...dcerpc.CallOption) (*GetMachineGroupInfoResponse, error)

	ModifyMachineGroup(context.Context, *ModifyMachineGroupRequest, ...dcerpc.CallOption) (*ModifyMachineGroupResponse, error)

	DeleteMachineGroup(context.Context, *DeleteMachineGroupRequest, ...dcerpc.CallOption) (*DeleteMachineGroupResponse, error)

	RenameMachineGroup(context.Context, *RenameMachineGroupRequest, ...dcerpc.CallOption) (*RenameMachineGroupResponse, error)

	AddMachine(context.Context, *AddMachineRequest, ...dcerpc.CallOption) (*AddMachineResponse, error)

	GetMachineInfo(context.Context, *GetMachineInfoRequest, ...dcerpc.CallOption) (*GetMachineInfoResponse, error)

	ModifyMachineInfo(context.Context, *ModifyMachineInfoRequest, ...dcerpc.CallOption) (*ModifyMachineInfoResponse, error)

	DeleteMachine(context.Context, *DeleteMachineRequest, ...dcerpc.CallOption) (*DeleteMachineResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) IwrmMachineGroupClient
}

type xxx_DefaultIwrmMachineGroupClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultIwrmMachineGroupClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultIwrmMachineGroupClient) CreateMachineGroup(ctx context.Context, in *CreateMachineGroupRequest, opts ...dcerpc.CallOption) (*CreateMachineGroupResponse, error) {
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
	out := &CreateMachineGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIwrmMachineGroupClient) GetMachineGroupInfo(ctx context.Context, in *GetMachineGroupInfoRequest, opts ...dcerpc.CallOption) (*GetMachineGroupInfoResponse, error) {
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
	out := &GetMachineGroupInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIwrmMachineGroupClient) ModifyMachineGroup(ctx context.Context, in *ModifyMachineGroupRequest, opts ...dcerpc.CallOption) (*ModifyMachineGroupResponse, error) {
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
	out := &ModifyMachineGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIwrmMachineGroupClient) DeleteMachineGroup(ctx context.Context, in *DeleteMachineGroupRequest, opts ...dcerpc.CallOption) (*DeleteMachineGroupResponse, error) {
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
	out := &DeleteMachineGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIwrmMachineGroupClient) RenameMachineGroup(ctx context.Context, in *RenameMachineGroupRequest, opts ...dcerpc.CallOption) (*RenameMachineGroupResponse, error) {
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
	out := &RenameMachineGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIwrmMachineGroupClient) AddMachine(ctx context.Context, in *AddMachineRequest, opts ...dcerpc.CallOption) (*AddMachineResponse, error) {
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
	out := &AddMachineResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIwrmMachineGroupClient) GetMachineInfo(ctx context.Context, in *GetMachineInfoRequest, opts ...dcerpc.CallOption) (*GetMachineInfoResponse, error) {
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
	out := &GetMachineInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIwrmMachineGroupClient) ModifyMachineInfo(ctx context.Context, in *ModifyMachineInfoRequest, opts ...dcerpc.CallOption) (*ModifyMachineInfoResponse, error) {
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
	out := &ModifyMachineInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIwrmMachineGroupClient) DeleteMachine(ctx context.Context, in *DeleteMachineRequest, opts ...dcerpc.CallOption) (*DeleteMachineResponse, error) {
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
	out := &DeleteMachineResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultIwrmMachineGroupClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultIwrmMachineGroupClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultIwrmMachineGroupClient) IPID(ctx context.Context, ipid *dcom.IPID) IwrmMachineGroupClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultIwrmMachineGroupClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewIwrmMachineGroupClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (IwrmMachineGroupClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(IwrmMachineGroupSyntaxV0_0))...)
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
	return &xxx_DefaultIwrmMachineGroupClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_CreateMachineGroupOperation structure represents the CreateMachineGroup operation
type xxx_CreateMachineGroupOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	ParentMachineGroupID *oaut.String   `idl:"name:bstrParentMachineGroupId" json:"parent_machine_group_id"`
	MachineGroupInfo     *oaut.String   `idl:"name:bstrMachineGroupInfo" json:"machine_group_info"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateMachineGroupOperation) OpNum() int { return 7 }

func (o *xxx_CreateMachineGroupOperation) OpName() string {
	return "/IWRMMachineGroup/v0/CreateMachineGroup"
}

func (o *xxx_CreateMachineGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateMachineGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrParentMachineGroupId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ParentMachineGroupID != nil {
			_ptr_bstrParentMachineGroupId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ParentMachineGroupID != nil {
					if err := o.ParentMachineGroupID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ParentMachineGroupID, _ptr_bstrParentMachineGroupId); err != nil {
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
	// bstrMachineGroupInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineGroupInfo != nil {
			_ptr_bstrMachineGroupInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MachineGroupInfo != nil {
					if err := o.MachineGroupInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineGroupInfo, _ptr_bstrMachineGroupInfo); err != nil {
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

func (o *xxx_CreateMachineGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrParentMachineGroupId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrParentMachineGroupId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ParentMachineGroupID == nil {
				o.ParentMachineGroupID = &oaut.String{}
			}
			if err := o.ParentMachineGroupID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrParentMachineGroupId := func(ptr interface{}) { o.ParentMachineGroupID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ParentMachineGroupID, _s_bstrParentMachineGroupId, _ptr_bstrParentMachineGroupId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrMachineGroupInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMachineGroupInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineGroupInfo == nil {
				o.MachineGroupInfo = &oaut.String{}
			}
			if err := o.MachineGroupInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMachineGroupInfo := func(ptr interface{}) { o.MachineGroupInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineGroupInfo, _s_bstrMachineGroupInfo, _ptr_bstrMachineGroupInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateMachineGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateMachineGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateMachineGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateMachineGroupRequest structure represents the CreateMachineGroup operation request
type CreateMachineGroupRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	ParentMachineGroupID *oaut.String   `idl:"name:bstrParentMachineGroupId" json:"parent_machine_group_id"`
	MachineGroupInfo     *oaut.String   `idl:"name:bstrMachineGroupInfo" json:"machine_group_info"`
}

func (o *CreateMachineGroupRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateMachineGroupOperation) *xxx_CreateMachineGroupOperation {
	if op == nil {
		op = &xxx_CreateMachineGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ParentMachineGroupID = o.ParentMachineGroupID
	op.MachineGroupInfo = o.MachineGroupInfo
	return op
}

func (o *CreateMachineGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateMachineGroupOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ParentMachineGroupID = op.ParentMachineGroupID
	o.MachineGroupInfo = op.MachineGroupInfo
}
func (o *CreateMachineGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateMachineGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateMachineGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateMachineGroupResponse structure represents the CreateMachineGroup operation response
type CreateMachineGroupResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CreateMachineGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateMachineGroupResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateMachineGroupOperation) *xxx_CreateMachineGroupOperation {
	if op == nil {
		op = &xxx_CreateMachineGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CreateMachineGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateMachineGroupOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CreateMachineGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateMachineGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateMachineGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMachineGroupInfoOperation structure represents the GetMachineGroupInfo operation
type xxx_GetMachineGroupInfoOperation struct {
	This             *dcom.ORPCThis `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	MachineGroupID   *oaut.String   `idl:"name:bstrMachineGroupId" json:"machine_group_id"`
	MachineGroupInfo *oaut.String   `idl:"name:pbstrMachineGroupInfo" json:"machine_group_info"`
	Return           int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMachineGroupInfoOperation) OpNum() int { return 8 }

func (o *xxx_GetMachineGroupInfoOperation) OpName() string {
	return "/IWRMMachineGroup/v0/GetMachineGroupInfo"
}

func (o *xxx_GetMachineGroupInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMachineGroupInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrMachineGroupId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineGroupID != nil {
			_ptr_bstrMachineGroupId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MachineGroupID != nil {
					if err := o.MachineGroupID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineGroupID, _ptr_bstrMachineGroupId); err != nil {
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

func (o *xxx_GetMachineGroupInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrMachineGroupId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMachineGroupId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineGroupID == nil {
				o.MachineGroupID = &oaut.String{}
			}
			if err := o.MachineGroupID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMachineGroupId := func(ptr interface{}) { o.MachineGroupID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineGroupID, _s_bstrMachineGroupId, _ptr_bstrMachineGroupId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMachineGroupInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMachineGroupInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrMachineGroupInfo {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineGroupInfo != nil {
			_ptr_pbstrMachineGroupInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MachineGroupInfo != nil {
					if err := o.MachineGroupInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineGroupInfo, _ptr_pbstrMachineGroupInfo); err != nil {
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

func (o *xxx_GetMachineGroupInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrMachineGroupInfo {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrMachineGroupInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineGroupInfo == nil {
				o.MachineGroupInfo = &oaut.String{}
			}
			if err := o.MachineGroupInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrMachineGroupInfo := func(ptr interface{}) { o.MachineGroupInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineGroupInfo, _s_pbstrMachineGroupInfo, _ptr_pbstrMachineGroupInfo); err != nil {
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

// GetMachineGroupInfoRequest structure represents the GetMachineGroupInfo operation request
type GetMachineGroupInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	MachineGroupID *oaut.String   `idl:"name:bstrMachineGroupId" json:"machine_group_id"`
}

func (o *GetMachineGroupInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMachineGroupInfoOperation) *xxx_GetMachineGroupInfoOperation {
	if op == nil {
		op = &xxx_GetMachineGroupInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MachineGroupID = o.MachineGroupID
	return op
}

func (o *GetMachineGroupInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMachineGroupInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MachineGroupID = op.MachineGroupID
}
func (o *GetMachineGroupInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMachineGroupInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMachineGroupInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMachineGroupInfoResponse structure represents the GetMachineGroupInfo operation response
type GetMachineGroupInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That             *dcom.ORPCThat `idl:"name:That" json:"that"`
	MachineGroupInfo *oaut.String   `idl:"name:pbstrMachineGroupInfo" json:"machine_group_info"`
	// Return: The GetMachineGroupInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMachineGroupInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMachineGroupInfoOperation) *xxx_GetMachineGroupInfoOperation {
	if op == nil {
		op = &xxx_GetMachineGroupInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MachineGroupInfo = o.MachineGroupInfo
	op.Return = o.Return
	return op
}

func (o *GetMachineGroupInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMachineGroupInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MachineGroupInfo = op.MachineGroupInfo
	o.Return = op.Return
}
func (o *GetMachineGroupInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMachineGroupInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMachineGroupInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ModifyMachineGroupOperation structure represents the ModifyMachineGroup operation
type xxx_ModifyMachineGroupOperation struct {
	This               *dcom.ORPCThis                `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat                `idl:"name:That" json:"that"`
	MachineGroupID     *oaut.String                  `idl:"name:bstrMachineGroupId" json:"machine_group_id"`
	MachineGroupInfo   *oaut.String                  `idl:"name:bstrMachineGroupInfo" json:"machine_group_info"`
	EnumMGMergeOptions wsrm.MachineGroupMergeOptions `idl:"name:enumMGMergeOptions" json:"enum_mg_merge_options"`
	Return             int32                         `idl:"name:Return" json:"return"`
}

func (o *xxx_ModifyMachineGroupOperation) OpNum() int { return 9 }

func (o *xxx_ModifyMachineGroupOperation) OpName() string {
	return "/IWRMMachineGroup/v0/ModifyMachineGroup"
}

func (o *xxx_ModifyMachineGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyMachineGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrMachineGroupId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineGroupID != nil {
			_ptr_bstrMachineGroupId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MachineGroupID != nil {
					if err := o.MachineGroupID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineGroupID, _ptr_bstrMachineGroupId); err != nil {
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
	// bstrMachineGroupInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineGroupInfo != nil {
			_ptr_bstrMachineGroupInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MachineGroupInfo != nil {
					if err := o.MachineGroupInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineGroupInfo, _ptr_bstrMachineGroupInfo); err != nil {
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
	// enumMGMergeOptions {in} (1:{v1_enum, alias=MACHINE_GROUP_MERGE_OPTIONS}(enum))
	{
		if err := w.WriteEnum(uint32(o.EnumMGMergeOptions)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyMachineGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrMachineGroupId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMachineGroupId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineGroupID == nil {
				o.MachineGroupID = &oaut.String{}
			}
			if err := o.MachineGroupID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMachineGroupId := func(ptr interface{}) { o.MachineGroupID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineGroupID, _s_bstrMachineGroupId, _ptr_bstrMachineGroupId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrMachineGroupInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMachineGroupInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineGroupInfo == nil {
				o.MachineGroupInfo = &oaut.String{}
			}
			if err := o.MachineGroupInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMachineGroupInfo := func(ptr interface{}) { o.MachineGroupInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineGroupInfo, _s_bstrMachineGroupInfo, _ptr_bstrMachineGroupInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// enumMGMergeOptions {in} (1:{v1_enum, alias=MACHINE_GROUP_MERGE_OPTIONS}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.EnumMGMergeOptions)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyMachineGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyMachineGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ModifyMachineGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ModifyMachineGroupRequest structure represents the ModifyMachineGroup operation request
type ModifyMachineGroupRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This               *dcom.ORPCThis                `idl:"name:This" json:"this"`
	MachineGroupID     *oaut.String                  `idl:"name:bstrMachineGroupId" json:"machine_group_id"`
	MachineGroupInfo   *oaut.String                  `idl:"name:bstrMachineGroupInfo" json:"machine_group_info"`
	EnumMGMergeOptions wsrm.MachineGroupMergeOptions `idl:"name:enumMGMergeOptions" json:"enum_mg_merge_options"`
}

func (o *ModifyMachineGroupRequest) xxx_ToOp(ctx context.Context, op *xxx_ModifyMachineGroupOperation) *xxx_ModifyMachineGroupOperation {
	if op == nil {
		op = &xxx_ModifyMachineGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MachineGroupID = o.MachineGroupID
	op.MachineGroupInfo = o.MachineGroupInfo
	op.EnumMGMergeOptions = o.EnumMGMergeOptions
	return op
}

func (o *ModifyMachineGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_ModifyMachineGroupOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MachineGroupID = op.MachineGroupID
	o.MachineGroupInfo = op.MachineGroupInfo
	o.EnumMGMergeOptions = op.EnumMGMergeOptions
}
func (o *ModifyMachineGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ModifyMachineGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyMachineGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ModifyMachineGroupResponse structure represents the ModifyMachineGroup operation response
type ModifyMachineGroupResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ModifyMachineGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ModifyMachineGroupResponse) xxx_ToOp(ctx context.Context, op *xxx_ModifyMachineGroupOperation) *xxx_ModifyMachineGroupOperation {
	if op == nil {
		op = &xxx_ModifyMachineGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ModifyMachineGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_ModifyMachineGroupOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ModifyMachineGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ModifyMachineGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyMachineGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteMachineGroupOperation structure represents the DeleteMachineGroup operation
type xxx_DeleteMachineGroupOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	MachineGroupID *oaut.String   `idl:"name:bstrMachineGroupId" json:"machine_group_id"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteMachineGroupOperation) OpNum() int { return 10 }

func (o *xxx_DeleteMachineGroupOperation) OpName() string {
	return "/IWRMMachineGroup/v0/DeleteMachineGroup"
}

func (o *xxx_DeleteMachineGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteMachineGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrMachineGroupId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineGroupID != nil {
			_ptr_bstrMachineGroupId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MachineGroupID != nil {
					if err := o.MachineGroupID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineGroupID, _ptr_bstrMachineGroupId); err != nil {
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

func (o *xxx_DeleteMachineGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrMachineGroupId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMachineGroupId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineGroupID == nil {
				o.MachineGroupID = &oaut.String{}
			}
			if err := o.MachineGroupID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMachineGroupId := func(ptr interface{}) { o.MachineGroupID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineGroupID, _s_bstrMachineGroupId, _ptr_bstrMachineGroupId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteMachineGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteMachineGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteMachineGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteMachineGroupRequest structure represents the DeleteMachineGroup operation request
type DeleteMachineGroupRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	MachineGroupID *oaut.String   `idl:"name:bstrMachineGroupId" json:"machine_group_id"`
}

func (o *DeleteMachineGroupRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteMachineGroupOperation) *xxx_DeleteMachineGroupOperation {
	if op == nil {
		op = &xxx_DeleteMachineGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MachineGroupID = o.MachineGroupID
	return op
}

func (o *DeleteMachineGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteMachineGroupOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MachineGroupID = op.MachineGroupID
}
func (o *DeleteMachineGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteMachineGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteMachineGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteMachineGroupResponse structure represents the DeleteMachineGroup operation response
type DeleteMachineGroupResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteMachineGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteMachineGroupResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteMachineGroupOperation) *xxx_DeleteMachineGroupOperation {
	if op == nil {
		op = &xxx_DeleteMachineGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteMachineGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteMachineGroupOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteMachineGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteMachineGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteMachineGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RenameMachineGroupOperation structure represents the RenameMachineGroup operation
type xxx_RenameMachineGroupOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	OldMachineGroupName *oaut.String   `idl:"name:bstrOldMachineGroupName" json:"old_machine_group_name"`
	NewMachineGroupName *oaut.String   `idl:"name:bstrNewMachineGroupName" json:"new_machine_group_name"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RenameMachineGroupOperation) OpNum() int { return 11 }

func (o *xxx_RenameMachineGroupOperation) OpName() string {
	return "/IWRMMachineGroup/v0/RenameMachineGroup"
}

func (o *xxx_RenameMachineGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameMachineGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrOldMachineGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.OldMachineGroupName != nil {
			_ptr_bstrOldMachineGroupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OldMachineGroupName != nil {
					if err := o.OldMachineGroupName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OldMachineGroupName, _ptr_bstrOldMachineGroupName); err != nil {
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
	// bstrNewMachineGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.NewMachineGroupName != nil {
			_ptr_bstrNewMachineGroupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.NewMachineGroupName != nil {
					if err := o.NewMachineGroupName.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NewMachineGroupName, _ptr_bstrNewMachineGroupName); err != nil {
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

func (o *xxx_RenameMachineGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrOldMachineGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrOldMachineGroupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OldMachineGroupName == nil {
				o.OldMachineGroupName = &oaut.String{}
			}
			if err := o.OldMachineGroupName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrOldMachineGroupName := func(ptr interface{}) { o.OldMachineGroupName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.OldMachineGroupName, _s_bstrOldMachineGroupName, _ptr_bstrOldMachineGroupName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrNewMachineGroupName {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrNewMachineGroupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.NewMachineGroupName == nil {
				o.NewMachineGroupName = &oaut.String{}
			}
			if err := o.NewMachineGroupName.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrNewMachineGroupName := func(ptr interface{}) { o.NewMachineGroupName = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.NewMachineGroupName, _s_bstrNewMachineGroupName, _ptr_bstrNewMachineGroupName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameMachineGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RenameMachineGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RenameMachineGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RenameMachineGroupRequest structure represents the RenameMachineGroup operation request
type RenameMachineGroupRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	OldMachineGroupName *oaut.String   `idl:"name:bstrOldMachineGroupName" json:"old_machine_group_name"`
	NewMachineGroupName *oaut.String   `idl:"name:bstrNewMachineGroupName" json:"new_machine_group_name"`
}

func (o *RenameMachineGroupRequest) xxx_ToOp(ctx context.Context, op *xxx_RenameMachineGroupOperation) *xxx_RenameMachineGroupOperation {
	if op == nil {
		op = &xxx_RenameMachineGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.OldMachineGroupName = o.OldMachineGroupName
	op.NewMachineGroupName = o.NewMachineGroupName
	return op
}

func (o *RenameMachineGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_RenameMachineGroupOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.OldMachineGroupName = op.OldMachineGroupName
	o.NewMachineGroupName = op.NewMachineGroupName
}
func (o *RenameMachineGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RenameMachineGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RenameMachineGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RenameMachineGroupResponse structure represents the RenameMachineGroup operation response
type RenameMachineGroupResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RenameMachineGroup return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RenameMachineGroupResponse) xxx_ToOp(ctx context.Context, op *xxx_RenameMachineGroupOperation) *xxx_RenameMachineGroupOperation {
	if op == nil {
		op = &xxx_RenameMachineGroupOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RenameMachineGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_RenameMachineGroupOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RenameMachineGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RenameMachineGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RenameMachineGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddMachineOperation structure represents the AddMachine operation
type xxx_AddMachineOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	ParentMachineGroupID *oaut.String   `idl:"name:bstrParentMachineGroupId" json:"parent_machine_group_id"`
	MachineInfo          *oaut.String   `idl:"name:bstrMachineInfo" json:"machine_info"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AddMachineOperation) OpNum() int { return 12 }

func (o *xxx_AddMachineOperation) OpName() string { return "/IWRMMachineGroup/v0/AddMachine" }

func (o *xxx_AddMachineOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMachineOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrParentMachineGroupId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ParentMachineGroupID != nil {
			_ptr_bstrParentMachineGroupId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ParentMachineGroupID != nil {
					if err := o.ParentMachineGroupID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ParentMachineGroupID, _ptr_bstrParentMachineGroupId); err != nil {
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
	// bstrMachineInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineInfo != nil {
			_ptr_bstrMachineInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MachineInfo != nil {
					if err := o.MachineInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineInfo, _ptr_bstrMachineInfo); err != nil {
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

func (o *xxx_AddMachineOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrParentMachineGroupId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrParentMachineGroupId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ParentMachineGroupID == nil {
				o.ParentMachineGroupID = &oaut.String{}
			}
			if err := o.ParentMachineGroupID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrParentMachineGroupId := func(ptr interface{}) { o.ParentMachineGroupID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ParentMachineGroupID, _s_bstrParentMachineGroupId, _ptr_bstrParentMachineGroupId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrMachineInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMachineInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineInfo == nil {
				o.MachineInfo = &oaut.String{}
			}
			if err := o.MachineInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMachineInfo := func(ptr interface{}) { o.MachineInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineInfo, _s_bstrMachineInfo, _ptr_bstrMachineInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMachineOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddMachineOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AddMachineOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AddMachineRequest structure represents the AddMachine operation request
type AddMachineRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	ParentMachineGroupID *oaut.String   `idl:"name:bstrParentMachineGroupId" json:"parent_machine_group_id"`
	MachineInfo          *oaut.String   `idl:"name:bstrMachineInfo" json:"machine_info"`
}

func (o *AddMachineRequest) xxx_ToOp(ctx context.Context, op *xxx_AddMachineOperation) *xxx_AddMachineOperation {
	if op == nil {
		op = &xxx_AddMachineOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ParentMachineGroupID = o.ParentMachineGroupID
	op.MachineInfo = o.MachineInfo
	return op
}

func (o *AddMachineRequest) xxx_FromOp(ctx context.Context, op *xxx_AddMachineOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ParentMachineGroupID = op.ParentMachineGroupID
	o.MachineInfo = op.MachineInfo
}
func (o *AddMachineRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AddMachineRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddMachineOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddMachineResponse structure represents the AddMachine operation response
type AddMachineResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AddMachine return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddMachineResponse) xxx_ToOp(ctx context.Context, op *xxx_AddMachineOperation) *xxx_AddMachineOperation {
	if op == nil {
		op = &xxx_AddMachineOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *AddMachineResponse) xxx_FromOp(ctx context.Context, op *xxx_AddMachineOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AddMachineResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AddMachineResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddMachineOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMachineInfoOperation structure represents the GetMachineInfo operation
type xxx_GetMachineInfoOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MachineID   *oaut.String   `idl:"name:bstrMachineId" json:"machine_id"`
	MachineInfo *oaut.String   `idl:"name:pbstrMachineInfo" json:"machine_info"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMachineInfoOperation) OpNum() int { return 13 }

func (o *xxx_GetMachineInfoOperation) OpName() string { return "/IWRMMachineGroup/v0/GetMachineInfo" }

func (o *xxx_GetMachineInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMachineInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrMachineId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineID != nil {
			_ptr_bstrMachineId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MachineID != nil {
					if err := o.MachineID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineID, _ptr_bstrMachineId); err != nil {
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

func (o *xxx_GetMachineInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrMachineId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMachineId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineID == nil {
				o.MachineID = &oaut.String{}
			}
			if err := o.MachineID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMachineId := func(ptr interface{}) { o.MachineID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineID, _s_bstrMachineId, _ptr_bstrMachineId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMachineInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMachineInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbstrMachineInfo {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineInfo != nil {
			_ptr_pbstrMachineInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MachineInfo != nil {
					if err := o.MachineInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineInfo, _ptr_pbstrMachineInfo); err != nil {
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

func (o *xxx_GetMachineInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbstrMachineInfo {out} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_pbstrMachineInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineInfo == nil {
				o.MachineInfo = &oaut.String{}
			}
			if err := o.MachineInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pbstrMachineInfo := func(ptr interface{}) { o.MachineInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineInfo, _s_pbstrMachineInfo, _ptr_pbstrMachineInfo); err != nil {
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

// GetMachineInfoRequest structure represents the GetMachineInfo operation request
type GetMachineInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	MachineID *oaut.String   `idl:"name:bstrMachineId" json:"machine_id"`
}

func (o *GetMachineInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMachineInfoOperation) *xxx_GetMachineInfoOperation {
	if op == nil {
		op = &xxx_GetMachineInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.MachineID = o.MachineID
	return op
}

func (o *GetMachineInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMachineInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.MachineID = op.MachineID
}
func (o *GetMachineInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMachineInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMachineInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMachineInfoResponse structure represents the GetMachineInfo operation response
type GetMachineInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	MachineInfo *oaut.String   `idl:"name:pbstrMachineInfo" json:"machine_info"`
	// Return: The GetMachineInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetMachineInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMachineInfoOperation) *xxx_GetMachineInfoOperation {
	if op == nil {
		op = &xxx_GetMachineInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.MachineInfo = o.MachineInfo
	op.Return = o.Return
	return op
}

func (o *GetMachineInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMachineInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.MachineInfo = op.MachineInfo
	o.Return = op.Return
}
func (o *GetMachineInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMachineInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMachineInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ModifyMachineInfoOperation structure represents the ModifyMachineInfo operation
type xxx_ModifyMachineInfoOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	ParentMachineGroupID *oaut.String   `idl:"name:bstrParentMachineGroupId" json:"parent_machine_group_id"`
	MachineID            *oaut.String   `idl:"name:bstrMachineId" json:"machine_id"`
	MachineInfo          *oaut.String   `idl:"name:bstrMachineInfo" json:"machine_info"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ModifyMachineInfoOperation) OpNum() int { return 14 }

func (o *xxx_ModifyMachineInfoOperation) OpName() string {
	return "/IWRMMachineGroup/v0/ModifyMachineInfo"
}

func (o *xxx_ModifyMachineInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyMachineInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrParentMachineGroupId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ParentMachineGroupID != nil {
			_ptr_bstrParentMachineGroupId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ParentMachineGroupID != nil {
					if err := o.ParentMachineGroupID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ParentMachineGroupID, _ptr_bstrParentMachineGroupId); err != nil {
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
	// bstrMachineId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineID != nil {
			_ptr_bstrMachineId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MachineID != nil {
					if err := o.MachineID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineID, _ptr_bstrMachineId); err != nil {
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
	// bstrMachineInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineInfo != nil {
			_ptr_bstrMachineInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MachineInfo != nil {
					if err := o.MachineInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineInfo, _ptr_bstrMachineInfo); err != nil {
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

func (o *xxx_ModifyMachineInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrParentMachineGroupId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrParentMachineGroupId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ParentMachineGroupID == nil {
				o.ParentMachineGroupID = &oaut.String{}
			}
			if err := o.ParentMachineGroupID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrParentMachineGroupId := func(ptr interface{}) { o.ParentMachineGroupID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ParentMachineGroupID, _s_bstrParentMachineGroupId, _ptr_bstrParentMachineGroupId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrMachineId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMachineId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineID == nil {
				o.MachineID = &oaut.String{}
			}
			if err := o.MachineID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMachineId := func(ptr interface{}) { o.MachineID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineID, _s_bstrMachineId, _ptr_bstrMachineId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrMachineInfo {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMachineInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineInfo == nil {
				o.MachineInfo = &oaut.String{}
			}
			if err := o.MachineInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMachineInfo := func(ptr interface{}) { o.MachineInfo = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineInfo, _s_bstrMachineInfo, _ptr_bstrMachineInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyMachineInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ModifyMachineInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ModifyMachineInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ModifyMachineInfoRequest structure represents the ModifyMachineInfo operation request
type ModifyMachineInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	ParentMachineGroupID *oaut.String   `idl:"name:bstrParentMachineGroupId" json:"parent_machine_group_id"`
	MachineID            *oaut.String   `idl:"name:bstrMachineId" json:"machine_id"`
	MachineInfo          *oaut.String   `idl:"name:bstrMachineInfo" json:"machine_info"`
}

func (o *ModifyMachineInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_ModifyMachineInfoOperation) *xxx_ModifyMachineInfoOperation {
	if op == nil {
		op = &xxx_ModifyMachineInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ParentMachineGroupID = o.ParentMachineGroupID
	op.MachineID = o.MachineID
	op.MachineInfo = o.MachineInfo
	return op
}

func (o *ModifyMachineInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_ModifyMachineInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ParentMachineGroupID = op.ParentMachineGroupID
	o.MachineID = op.MachineID
	o.MachineInfo = op.MachineInfo
}
func (o *ModifyMachineInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ModifyMachineInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyMachineInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ModifyMachineInfoResponse structure represents the ModifyMachineInfo operation response
type ModifyMachineInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ModifyMachineInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ModifyMachineInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_ModifyMachineInfoOperation) *xxx_ModifyMachineInfoOperation {
	if op == nil {
		op = &xxx_ModifyMachineInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ModifyMachineInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_ModifyMachineInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ModifyMachineInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ModifyMachineInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ModifyMachineInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteMachineOperation structure represents the DeleteMachine operation
type xxx_DeleteMachineOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	ParentMachineGroupID *oaut.String   `idl:"name:bstrParentMachineGroupId" json:"parent_machine_group_id"`
	MachineID            *oaut.String   `idl:"name:bstrMachineId" json:"machine_id"`
	Recursive            bool           `idl:"name:bRecursive" json:"recursive"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteMachineOperation) OpNum() int { return 15 }

func (o *xxx_DeleteMachineOperation) OpName() string { return "/IWRMMachineGroup/v0/DeleteMachine" }

func (o *xxx_DeleteMachineOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteMachineOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bstrParentMachineGroupId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ParentMachineGroupID != nil {
			_ptr_bstrParentMachineGroupId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ParentMachineGroupID != nil {
					if err := o.ParentMachineGroupID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ParentMachineGroupID, _ptr_bstrParentMachineGroupId); err != nil {
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
	// bstrMachineId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.MachineID != nil {
			_ptr_bstrMachineId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MachineID != nil {
					if err := o.MachineID.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineID, _ptr_bstrMachineId); err != nil {
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
	// bRecursive {in} (1:{alias=BOOL}(int32))
	{
		if !o.Recursive {
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

func (o *xxx_DeleteMachineOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bstrParentMachineGroupId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrParentMachineGroupId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ParentMachineGroupID == nil {
				o.ParentMachineGroupID = &oaut.String{}
			}
			if err := o.ParentMachineGroupID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrParentMachineGroupId := func(ptr interface{}) { o.ParentMachineGroupID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ParentMachineGroupID, _s_bstrParentMachineGroupId, _ptr_bstrParentMachineGroupId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bstrMachineId {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_bstrMachineId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.MachineID == nil {
				o.MachineID = &oaut.String{}
			}
			if err := o.MachineID.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_bstrMachineId := func(ptr interface{}) { o.MachineID = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.MachineID, _s_bstrMachineId, _ptr_bstrMachineId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// bRecursive {in} (1:{alias=BOOL}(int32))
	{
		var _bRecursive int32
		if err := w.ReadData(&_bRecursive); err != nil {
			return err
		}
		o.Recursive = _bRecursive != 0
	}
	return nil
}

func (o *xxx_DeleteMachineOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteMachineOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteMachineOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteMachineRequest structure represents the DeleteMachine operation request
type DeleteMachineRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	ParentMachineGroupID *oaut.String   `idl:"name:bstrParentMachineGroupId" json:"parent_machine_group_id"`
	MachineID            *oaut.String   `idl:"name:bstrMachineId" json:"machine_id"`
	Recursive            bool           `idl:"name:bRecursive" json:"recursive"`
}

func (o *DeleteMachineRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteMachineOperation) *xxx_DeleteMachineOperation {
	if op == nil {
		op = &xxx_DeleteMachineOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ParentMachineGroupID = o.ParentMachineGroupID
	op.MachineID = o.MachineID
	op.Recursive = o.Recursive
	return op
}

func (o *DeleteMachineRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteMachineOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ParentMachineGroupID = op.ParentMachineGroupID
	o.MachineID = op.MachineID
	o.Recursive = op.Recursive
}
func (o *DeleteMachineRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteMachineRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteMachineOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteMachineResponse structure represents the DeleteMachine operation response
type DeleteMachineResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteMachine return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteMachineResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteMachineOperation) *xxx_DeleteMachineOperation {
	if op == nil {
		op = &xxx_DeleteMachineOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteMachineResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteMachineOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteMachineResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteMachineResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteMachineOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
