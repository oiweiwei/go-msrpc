package ivsssnapshotmgmt

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	scmp "github.com/oiweiwei/go-msrpc/msrpc/dcom/scmp"
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
	_ = iunknown.GoPackage
	_ = scmp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/scmp"
)

var (
	// IVssSnapshotMgmt interface identifier fa7df749-66e7-4986-a27f-e2f04ae53772
	SnapshotManagementIID = &dcom.IID{Data1: 0xfa7df749, Data2: 0x66e7, Data3: 0x4986, Data4: []byte{0xa2, 0x7f, 0xe2, 0xf0, 0x4a, 0xe5, 0x37, 0x72}}
	// Syntax UUID
	SnapshotManagementSyntaxUUID = &uuid.UUID{TimeLow: 0xfa7df749, TimeMid: 0x66e7, TimeHiAndVersion: 0x4986, ClockSeqHiAndReserved: 0xa2, ClockSeqLow: 0x7f, Node: [6]uint8{0xe2, 0xf0, 0x4a, 0xe5, 0x37, 0x72}}
	// Syntax ID
	SnapshotManagementSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: SnapshotManagementSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVssSnapshotMgmt interface.
type SnapshotManagementClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	GetProviderManagementInterface(context.Context, *GetProviderManagementInterfaceRequest, ...dcerpc.CallOption) (*GetProviderManagementInterfaceResponse, error)

	QueryVolumesSupportedForSnapshots(context.Context, *QueryVolumesSupportedForSnapshotsRequest, ...dcerpc.CallOption) (*QueryVolumesSupportedForSnapshotsResponse, error)

	QuerySnapshotsByVolume(context.Context, *QuerySnapshotsByVolumeRequest, ...dcerpc.CallOption) (*QuerySnapshotsByVolumeResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) SnapshotManagementClient
}

type xxx_DefaultSnapshotManagementClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultSnapshotManagementClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultSnapshotManagementClient) GetProviderManagementInterface(ctx context.Context, in *GetProviderManagementInterfaceRequest, opts ...dcerpc.CallOption) (*GetProviderManagementInterfaceResponse, error) {
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
	out := &GetProviderManagementInterfaceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSnapshotManagementClient) QueryVolumesSupportedForSnapshots(ctx context.Context, in *QueryVolumesSupportedForSnapshotsRequest, opts ...dcerpc.CallOption) (*QueryVolumesSupportedForSnapshotsResponse, error) {
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
	out := &QueryVolumesSupportedForSnapshotsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSnapshotManagementClient) QuerySnapshotsByVolume(ctx context.Context, in *QuerySnapshotsByVolumeRequest, opts ...dcerpc.CallOption) (*QuerySnapshotsByVolumeResponse, error) {
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
	out := &QuerySnapshotsByVolumeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultSnapshotManagementClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultSnapshotManagementClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultSnapshotManagementClient) IPID(ctx context.Context, ipid *dcom.IPID) SnapshotManagementClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultSnapshotManagementClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewSnapshotManagementClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (SnapshotManagementClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(SnapshotManagementSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iunknown.NewUnknownClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultSnapshotManagementClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetProviderManagementInterfaceOperation structure represents the GetProviderMgmtInterface operation
type xxx_GetProviderManagementInterfaceOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ProviderID  *scmp.ID       `idl:"name:ProviderId" json:"provider_id"`
	InterfaceID *dcom.IID      `idl:"name:InterfaceId" json:"interface_id"`
	Interface   *dcom.Unknown  `idl:"name:ppItf" json:"interface"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetProviderManagementInterfaceOperation) OpNum() int { return 3 }

func (o *xxx_GetProviderManagementInterfaceOperation) OpName() string {
	return "/IVssSnapshotMgmt/v0/GetProviderMgmtInterface"
}

func (o *xxx_GetProviderManagementInterfaceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProviderManagementInterfaceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ProviderId {in} (1:{alias=VSS_ID, names=GUID}(struct))
	{
		if o.ProviderID != nil {
			if err := o.ProviderID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&scmp.ID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// InterfaceId {in} (1:{alias=REFIID}*(1))(2:{alias=IID, names=GUID}(struct))
	{
		if o.InterfaceID != nil {
			if err := o.InterfaceID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.IID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetProviderManagementInterfaceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ProviderId {in} (1:{alias=VSS_ID, names=GUID}(struct))
	{
		if o.ProviderID == nil {
			o.ProviderID = &scmp.ID{}
		}
		if err := o.ProviderID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// InterfaceId {in} (1:{alias=REFIID,pointer=ref}*(1))(2:{alias=IID, names=GUID}(struct))
	{
		if o.InterfaceID == nil {
			o.InterfaceID = &dcom.IID{}
		}
		if err := o.InterfaceID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProviderManagementInterfaceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProviderManagementInterfaceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppItf {out, iid_is={InterfaceId}} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		if o.Interface != nil {
			_ptr_ppItf := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Interface != nil {
					if err := o.Interface.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dcom.Unknown{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Interface, _ptr_ppItf); err != nil {
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

func (o *xxx_GetProviderManagementInterfaceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppItf {out, iid_is={InterfaceId}} (1:{pointer=ref}*(2)*(1))(2:{alias=IUnknown}(interface))
	{
		_ptr_ppItf := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Interface == nil {
				o.Interface = &dcom.Unknown{}
			}
			if err := o.Interface.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppItf := func(ptr interface{}) { o.Interface = *ptr.(**dcom.Unknown) }
		if err := w.ReadPointer(&o.Interface, _s_ppItf, _ptr_ppItf); err != nil {
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

// GetProviderManagementInterfaceRequest structure represents the GetProviderMgmtInterface operation request
type GetProviderManagementInterfaceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	ProviderID  *scmp.ID       `idl:"name:ProviderId" json:"provider_id"`
	InterfaceID *dcom.IID      `idl:"name:InterfaceId" json:"interface_id"`
}

func (o *GetProviderManagementInterfaceRequest) xxx_ToOp(ctx context.Context, op *xxx_GetProviderManagementInterfaceOperation) *xxx_GetProviderManagementInterfaceOperation {
	if op == nil {
		op = &xxx_GetProviderManagementInterfaceOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ProviderID = o.ProviderID
	op.InterfaceID = o.InterfaceID
	return op
}

func (o *GetProviderManagementInterfaceRequest) xxx_FromOp(ctx context.Context, op *xxx_GetProviderManagementInterfaceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ProviderID = op.ProviderID
	o.InterfaceID = op.InterfaceID
}
func (o *GetProviderManagementInterfaceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetProviderManagementInterfaceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetProviderManagementInterfaceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetProviderManagementInterfaceResponse structure represents the GetProviderMgmtInterface operation response
type GetProviderManagementInterfaceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Interface *dcom.Unknown  `idl:"name:ppItf" json:"interface"`
	// Return: The GetProviderMgmtInterface return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetProviderManagementInterfaceResponse) xxx_ToOp(ctx context.Context, op *xxx_GetProviderManagementInterfaceOperation) *xxx_GetProviderManagementInterfaceOperation {
	if op == nil {
		op = &xxx_GetProviderManagementInterfaceOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Interface = o.Interface
	op.Return = o.Return
	return op
}

func (o *GetProviderManagementInterfaceResponse) xxx_FromOp(ctx context.Context, op *xxx_GetProviderManagementInterfaceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Interface = op.Interface
	o.Return = op.Return
}
func (o *GetProviderManagementInterfaceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetProviderManagementInterfaceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetProviderManagementInterfaceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryVolumesSupportedForSnapshotsOperation structure represents the QueryVolumesSupportedForSnapshots operation
type xxx_QueryVolumesSupportedForSnapshotsOperation struct {
	This       *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat             `idl:"name:That" json:"that"`
	ProviderID *scmp.ID                   `idl:"name:ProviderId" json:"provider_id"`
	Context    int32                      `idl:"name:lContext" json:"context"`
	Enum       *scmp.EnumManagementObject `idl:"name:ppEnum" json:"enum"`
	Return     int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryVolumesSupportedForSnapshotsOperation) OpNum() int { return 4 }

func (o *xxx_QueryVolumesSupportedForSnapshotsOperation) OpName() string {
	return "/IVssSnapshotMgmt/v0/QueryVolumesSupportedForSnapshots"
}

func (o *xxx_QueryVolumesSupportedForSnapshotsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryVolumesSupportedForSnapshotsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ProviderId {in} (1:{alias=VSS_ID, names=GUID}(struct))
	{
		if o.ProviderID != nil {
			if err := o.ProviderID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&scmp.ID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lContext {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Context); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryVolumesSupportedForSnapshotsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ProviderId {in} (1:{alias=VSS_ID, names=GUID}(struct))
	{
		if o.ProviderID == nil {
			o.ProviderID = &scmp.ID{}
		}
		if err := o.ProviderID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lContext {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Context); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryVolumesSupportedForSnapshotsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryVolumesSupportedForSnapshotsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVssEnumMgmtObject}(interface))
	{
		if o.Enum != nil {
			_ptr_ppEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&scmp.EnumManagementObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_ppEnum); err != nil {
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

func (o *xxx_QueryVolumesSupportedForSnapshotsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVssEnumMgmtObject}(interface))
	{
		_ptr_ppEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &scmp.EnumManagementObject{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppEnum := func(ptr interface{}) { o.Enum = *ptr.(**scmp.EnumManagementObject) }
		if err := w.ReadPointer(&o.Enum, _s_ppEnum, _ptr_ppEnum); err != nil {
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

// QueryVolumesSupportedForSnapshotsRequest structure represents the QueryVolumesSupportedForSnapshots operation request
type QueryVolumesSupportedForSnapshotsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	ProviderID *scmp.ID       `idl:"name:ProviderId" json:"provider_id"`
	Context    int32          `idl:"name:lContext" json:"context"`
}

func (o *QueryVolumesSupportedForSnapshotsRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryVolumesSupportedForSnapshotsOperation) *xxx_QueryVolumesSupportedForSnapshotsOperation {
	if op == nil {
		op = &xxx_QueryVolumesSupportedForSnapshotsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.ProviderID = o.ProviderID
	op.Context = o.Context
	return op
}

func (o *QueryVolumesSupportedForSnapshotsRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryVolumesSupportedForSnapshotsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ProviderID = op.ProviderID
	o.Context = op.Context
}
func (o *QueryVolumesSupportedForSnapshotsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryVolumesSupportedForSnapshotsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryVolumesSupportedForSnapshotsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryVolumesSupportedForSnapshotsResponse structure represents the QueryVolumesSupportedForSnapshots operation response
type QueryVolumesSupportedForSnapshotsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat             `idl:"name:That" json:"that"`
	Enum *scmp.EnumManagementObject `idl:"name:ppEnum" json:"enum"`
	// Return: The QueryVolumesSupportedForSnapshots return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryVolumesSupportedForSnapshotsResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryVolumesSupportedForSnapshotsOperation) *xxx_QueryVolumesSupportedForSnapshotsOperation {
	if op == nil {
		op = &xxx_QueryVolumesSupportedForSnapshotsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Enum = o.Enum
	op.Return = o.Return
	return op
}

func (o *QueryVolumesSupportedForSnapshotsResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryVolumesSupportedForSnapshotsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *QueryVolumesSupportedForSnapshotsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryVolumesSupportedForSnapshotsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryVolumesSupportedForSnapshotsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QuerySnapshotsByVolumeOperation structure represents the QuerySnapshotsByVolume operation
type xxx_QuerySnapshotsByVolumeOperation struct {
	This       *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat   `idl:"name:That" json:"that"`
	VolumeName string           `idl:"name:pwszVolumeName" json:"volume_name"`
	ProviderID *scmp.ID         `idl:"name:ProviderId" json:"provider_id"`
	Enum       *scmp.EnumObject `idl:"name:ppEnum" json:"enum"`
	Return     int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_QuerySnapshotsByVolumeOperation) OpNum() int { return 5 }

func (o *xxx_QuerySnapshotsByVolumeOperation) OpName() string {
	return "/IVssSnapshotMgmt/v0/QuerySnapshotsByVolume"
}

func (o *xxx_QuerySnapshotsByVolumeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySnapshotsByVolumeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszVolumeName {in} (1:{string, pointer=unique, alias=VSS_PWSZ}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.VolumeName != "" {
			_ptr_pwszVolumeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.VolumeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.VolumeName, _ptr_pwszVolumeName); err != nil {
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
	// ProviderId {in} (1:{alias=VSS_ID, names=GUID}(struct))
	{
		if o.ProviderID != nil {
			if err := o.ProviderID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&scmp.ID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_QuerySnapshotsByVolumeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszVolumeName {in} (1:{string, pointer=unique, alias=VSS_PWSZ}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_pwszVolumeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.VolumeName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszVolumeName := func(ptr interface{}) { o.VolumeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.VolumeName, _s_pwszVolumeName, _ptr_pwszVolumeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ProviderId {in} (1:{alias=VSS_ID, names=GUID}(struct))
	{
		if o.ProviderID == nil {
			o.ProviderID = &scmp.ID{}
		}
		if err := o.ProviderID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySnapshotsByVolumeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySnapshotsByVolumeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVssEnumObject}(interface))
	{
		if o.Enum != nil {
			_ptr_ppEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&scmp.EnumObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_ppEnum); err != nil {
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

func (o *xxx_QuerySnapshotsByVolumeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVssEnumObject}(interface))
	{
		_ptr_ppEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &scmp.EnumObject{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppEnum := func(ptr interface{}) { o.Enum = *ptr.(**scmp.EnumObject) }
		if err := w.ReadPointer(&o.Enum, _s_ppEnum, _ptr_ppEnum); err != nil {
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

// QuerySnapshotsByVolumeRequest structure represents the QuerySnapshotsByVolume operation request
type QuerySnapshotsByVolumeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	VolumeName string         `idl:"name:pwszVolumeName" json:"volume_name"`
	ProviderID *scmp.ID       `idl:"name:ProviderId" json:"provider_id"`
}

func (o *QuerySnapshotsByVolumeRequest) xxx_ToOp(ctx context.Context, op *xxx_QuerySnapshotsByVolumeOperation) *xxx_QuerySnapshotsByVolumeOperation {
	if op == nil {
		op = &xxx_QuerySnapshotsByVolumeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.VolumeName = o.VolumeName
	op.ProviderID = o.ProviderID
	return op
}

func (o *QuerySnapshotsByVolumeRequest) xxx_FromOp(ctx context.Context, op *xxx_QuerySnapshotsByVolumeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeName = op.VolumeName
	o.ProviderID = op.ProviderID
}
func (o *QuerySnapshotsByVolumeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QuerySnapshotsByVolumeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QuerySnapshotsByVolumeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QuerySnapshotsByVolumeResponse structure represents the QuerySnapshotsByVolume operation response
type QuerySnapshotsByVolumeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Enum *scmp.EnumObject `idl:"name:ppEnum" json:"enum"`
	// Return: The QuerySnapshotsByVolume return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QuerySnapshotsByVolumeResponse) xxx_ToOp(ctx context.Context, op *xxx_QuerySnapshotsByVolumeOperation) *xxx_QuerySnapshotsByVolumeOperation {
	if op == nil {
		op = &xxx_QuerySnapshotsByVolumeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Enum = o.Enum
	op.Return = o.Return
	return op
}

func (o *QuerySnapshotsByVolumeResponse) xxx_FromOp(ctx context.Context, op *xxx_QuerySnapshotsByVolumeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *QuerySnapshotsByVolumeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QuerySnapshotsByVolumeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QuerySnapshotsByVolumeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
