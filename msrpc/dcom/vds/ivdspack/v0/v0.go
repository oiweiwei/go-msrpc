package ivdspack

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
	vds "github.com/oiweiwei/go-msrpc/msrpc/dcom/vds"
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
	_ = vds.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/vds"
)

var (
	// IVdsPack interface identifier 3b69d7f5-9d94-4648-91ca-79939ba263bf
	PackIID = &dcom.IID{Data1: 0x3b69d7f5, Data2: 0x9d94, Data3: 0x4648, Data4: []byte{0x91, 0xca, 0x79, 0x93, 0x9b, 0xa2, 0x63, 0xbf}}
	// Syntax UUID
	PackSyntaxUUID = &uuid.UUID{TimeLow: 0x3b69d7f5, TimeMid: 0x9d94, TimeHiAndVersion: 0x4648, ClockSeqHiAndReserved: 0x91, ClockSeqLow: 0xca, Node: [6]uint8{0x79, 0x93, 0x9b, 0xa2, 0x63, 0xbf}}
	// Syntax ID
	PackSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: PackSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsPack interface.
type PackClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// GetProperties operation.
	GetProperties(context.Context, *GetPropertiesRequest, ...dcerpc.CallOption) (*GetPropertiesResponse, error)

	// The GetProvider method retrieves the provider that the disk pack belongs to.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetProvider(context.Context, *GetProviderRequest, ...dcerpc.CallOption) (*GetProviderResponse, error)

	// The QueryVolumes method retrieves the volumes of a disk pack.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryVolumes(context.Context, *QueryVolumesRequest, ...dcerpc.CallOption) (*QueryVolumesResponse, error)

	// The QueryDisks method retrieves the disks of a disk pack.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryDisks(context.Context, *QueryDisksRequest, ...dcerpc.CallOption) (*QueryDisksResponse, error)

	// The CreateVolume method creates a volume in a disk pack.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	CreateVolume(context.Context, *CreateVolumeRequest, ...dcerpc.CallOption) (*CreateVolumeResponse, error)

	// This method initializes a disk that has no partitioning format defined, and then
	// adds the disk to the disk pack. AddDisk cannot redefine the partitioning format on
	// a disk.<82>
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	AddDisk(context.Context, *AddDiskRequest, ...dcerpc.CallOption) (*AddDiskResponse, error)

	// The MigrateDisks method migrates a set of disks from one pack to another pack.<83>
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// ERROR_SUCCESS (0x00000000)
	MigrateDisks(context.Context, *MigrateDisksRequest, ...dcerpc.CallOption) (*MigrateDisksResponse, error)

	// Opnum10NotUsedOnWire operation.
	// Opnum10NotUsedOnWire

	// The IVdsPack::RemoveMissingDisk method removes the specified missing disk from a
	// disk pack. This method only applies to dynamic disks. At least one dynamic disk needs
	// to be present to enumerate missing disks.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	RemoveMissingDisk(context.Context, *RemoveMissingDiskRequest, ...dcerpc.CallOption) (*RemoveMissingDiskResponse, error)

	// The Recover method restores a disk pack to a healthy state. This method is not supported
	// on basic disk packs or the INVALID dynamic disk pack (the value of VDS_PACK_PROP::pwszName
	// is INVALID for this pack). The INVALID dynamic disk pack contains dynamic disks that
	// have failed to be joined to the owning pack because there are errors or data corruption
	// has occurred.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Recover(context.Context, *RecoverRequest, ...dcerpc.CallOption) (*RecoverResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) PackClient
}

type xxx_DefaultPackClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultPackClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultPackClient) GetProperties(ctx context.Context, in *GetPropertiesRequest, opts ...dcerpc.CallOption) (*GetPropertiesResponse, error) {
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
	out := &GetPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPackClient) GetProvider(ctx context.Context, in *GetProviderRequest, opts ...dcerpc.CallOption) (*GetProviderResponse, error) {
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
	out := &GetProviderResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPackClient) QueryVolumes(ctx context.Context, in *QueryVolumesRequest, opts ...dcerpc.CallOption) (*QueryVolumesResponse, error) {
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
	out := &QueryVolumesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPackClient) QueryDisks(ctx context.Context, in *QueryDisksRequest, opts ...dcerpc.CallOption) (*QueryDisksResponse, error) {
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
	out := &QueryDisksResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPackClient) CreateVolume(ctx context.Context, in *CreateVolumeRequest, opts ...dcerpc.CallOption) (*CreateVolumeResponse, error) {
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
	out := &CreateVolumeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPackClient) AddDisk(ctx context.Context, in *AddDiskRequest, opts ...dcerpc.CallOption) (*AddDiskResponse, error) {
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
	out := &AddDiskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPackClient) MigrateDisks(ctx context.Context, in *MigrateDisksRequest, opts ...dcerpc.CallOption) (*MigrateDisksResponse, error) {
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
	out := &MigrateDisksResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPackClient) RemoveMissingDisk(ctx context.Context, in *RemoveMissingDiskRequest, opts ...dcerpc.CallOption) (*RemoveMissingDiskResponse, error) {
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
	out := &RemoveMissingDiskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPackClient) Recover(ctx context.Context, in *RecoverRequest, opts ...dcerpc.CallOption) (*RecoverResponse, error) {
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
	out := &RecoverResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPackClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultPackClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultPackClient) IPID(ctx context.Context, ipid *dcom.IPID) PackClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultPackClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewPackClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (PackClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(PackSyntaxV0_0))...)
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
	return &xxx_DefaultPackClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetPropertiesOperation structure represents the GetProperties operation
type xxx_GetPropertiesOperation struct {
	This         *dcom.ORPCThis    `idl:"name:This" json:"this"`
	That         *dcom.ORPCThat    `idl:"name:That" json:"that"`
	PackProperty *vds.PackProperty `idl:"name:pPackProp" json:"pack_property"`
	Return       int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertiesOperation) OpNum() int { return 3 }

func (o *xxx_GetPropertiesOperation) OpName() string { return "/IVdsPack/v0/GetProperties" }

func (o *xxx_GetPropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pPackProp {out} (1:{pointer=ref}*(1))(2:{alias=VDS_PACK_PROP}(struct))
	{
		if o.PackProperty != nil {
			if err := o.PackProperty.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.PackProperty{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetPropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pPackProp {out} (1:{pointer=ref}*(1))(2:{alias=VDS_PACK_PROP}(struct))
	{
		if o.PackProperty == nil {
			o.PackProperty = &vds.PackProperty{}
		}
		if err := o.PackProperty.UnmarshalNDR(ctx, w); err != nil {
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

// GetPropertiesRequest structure represents the GetProperties operation request
type GetPropertiesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPropertiesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesOperation) *xxx_GetPropertiesOperation {
	if op == nil {
		op = &xxx_GetPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPropertiesResponse structure represents the GetProperties operation response
type GetPropertiesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That         *dcom.ORPCThat    `idl:"name:That" json:"that"`
	PackProperty *vds.PackProperty `idl:"name:pPackProp" json:"pack_property"`
	// Return: The GetProperties return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPropertiesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPropertiesOperation) *xxx_GetPropertiesOperation {
	if op == nil {
		op = &xxx_GetPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PackProperty = o.PackProperty
	op.Return = o.Return
	return op
}

func (o *GetPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PackProperty = op.PackProperty
	o.Return = op.Return
}
func (o *GetPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetProviderOperation structure represents the GetProvider operation
type xxx_GetProviderOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	Provider *vds.Provider  `idl:"name:ppProvider" json:"provider"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetProviderOperation) OpNum() int { return 4 }

func (o *xxx_GetProviderOperation) OpName() string { return "/IVdsPack/v0/GetProvider" }

func (o *xxx_GetProviderOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProviderOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetProviderOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetProviderOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProviderOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppProvider {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsProvider}(interface))
	{
		if o.Provider != nil {
			_ptr_ppProvider := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Provider != nil {
					if err := o.Provider.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.Provider{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Provider, _ptr_ppProvider); err != nil {
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

func (o *xxx_GetProviderOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppProvider {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsProvider}(interface))
	{
		_ptr_ppProvider := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Provider == nil {
				o.Provider = &vds.Provider{}
			}
			if err := o.Provider.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppProvider := func(ptr interface{}) { o.Provider = *ptr.(**vds.Provider) }
		if err := w.ReadPointer(&o.Provider, _s_ppProvider, _ptr_ppProvider); err != nil {
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

// GetProviderRequest structure represents the GetProvider operation request
type GetProviderRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetProviderRequest) xxx_ToOp(ctx context.Context, op *xxx_GetProviderOperation) *xxx_GetProviderOperation {
	if op == nil {
		op = &xxx_GetProviderOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetProviderRequest) xxx_FromOp(ctx context.Context, op *xxx_GetProviderOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetProviderRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetProviderRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetProviderOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetProviderResponse structure represents the GetProvider operation response
type GetProviderResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppProvider: A pointer to an IVdsProvider interface that, if the operation is successfully
	// completed, receives the IVdsProvider interface of the provider object that the pack
	// belongs to. Callers MUST release the interface when they are done with it.
	Provider *vds.Provider `idl:"name:ppProvider" json:"provider"`
	// Return: The GetProvider return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetProviderResponse) xxx_ToOp(ctx context.Context, op *xxx_GetProviderOperation) *xxx_GetProviderOperation {
	if op == nil {
		op = &xxx_GetProviderOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Provider = o.Provider
	op.Return = o.Return
	return op
}

func (o *GetProviderResponse) xxx_FromOp(ctx context.Context, op *xxx_GetProviderOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Provider = op.Provider
	o.Return = op.Return
}
func (o *GetProviderResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetProviderResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetProviderOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryVolumesOperation structure represents the QueryVolumes operation
type xxx_QueryVolumesOperation struct {
	This   *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Enum   *vds.EnumObject `idl:"name:ppEnum" json:"enum"`
	Return int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryVolumesOperation) OpNum() int { return 5 }

func (o *xxx_QueryVolumesOperation) OpName() string { return "/IVdsPack/v0/QueryVolumes" }

func (o *xxx_QueryVolumesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryVolumesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryVolumesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_QueryVolumesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryVolumesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumVdsObject}(interface))
	{
		if o.Enum != nil {
			_ptr_ppEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.EnumObject{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_QueryVolumesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumVdsObject}(interface))
	{
		_ptr_ppEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &vds.EnumObject{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppEnum := func(ptr interface{}) { o.Enum = *ptr.(**vds.EnumObject) }
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

// QueryVolumesRequest structure represents the QueryVolumes operation request
type QueryVolumesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *QueryVolumesRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryVolumesOperation) *xxx_QueryVolumesOperation {
	if op == nil {
		op = &xxx_QueryVolumesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *QueryVolumesRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryVolumesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *QueryVolumesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryVolumesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryVolumesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryVolumesResponse structure represents the QueryVolumes operation response
type QueryVolumesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppEnum: A pointer to an IEnumVdsObject interface that, if the operation is successfully
	// completed, receives the IEnumVdsObject interface of the object that contains an enumeration
	// of volume objects in the pack. Callers MUST release the interface when they are done
	// with it.
	Enum *vds.EnumObject `idl:"name:ppEnum" json:"enum"`
	// Return: The QueryVolumes return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryVolumesResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryVolumesOperation) *xxx_QueryVolumesOperation {
	if op == nil {
		op = &xxx_QueryVolumesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Enum = o.Enum
	op.Return = o.Return
	return op
}

func (o *QueryVolumesResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryVolumesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *QueryVolumesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryVolumesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryVolumesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryDisksOperation structure represents the QueryDisks operation
type xxx_QueryDisksOperation struct {
	This   *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Enum   *vds.EnumObject `idl:"name:ppEnum" json:"enum"`
	Return int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryDisksOperation) OpNum() int { return 6 }

func (o *xxx_QueryDisksOperation) OpName() string { return "/IVdsPack/v0/QueryDisks" }

func (o *xxx_QueryDisksOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDisksOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryDisksOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_QueryDisksOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryDisksOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumVdsObject}(interface))
	{
		if o.Enum != nil {
			_ptr_ppEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.EnumObject{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_QueryDisksOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppEnum {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IEnumVdsObject}(interface))
	{
		_ptr_ppEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &vds.EnumObject{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppEnum := func(ptr interface{}) { o.Enum = *ptr.(**vds.EnumObject) }
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

// QueryDisksRequest structure represents the QueryDisks operation request
type QueryDisksRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *QueryDisksRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryDisksOperation) *xxx_QueryDisksOperation {
	if op == nil {
		op = &xxx_QueryDisksOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *QueryDisksRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryDisksOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *QueryDisksRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryDisksRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryDisksOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryDisksResponse structure represents the QueryDisks operation response
type QueryDisksResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppEnum: A pointer to an IEnumVdsObject interface that, if the operation is successfully
	// completed, receives the IEnumVdsObject interface of the object containing an enumeration
	// of disk objects in the pack. Callers MUST release the interface when they are done
	// with it.
	Enum *vds.EnumObject `idl:"name:ppEnum" json:"enum"`
	// Return: The QueryDisks return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryDisksResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryDisksOperation) *xxx_QueryDisksOperation {
	if op == nil {
		op = &xxx_QueryDisksOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Enum = o.Enum
	op.Return = o.Return
	return op
}

func (o *QueryDisksResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryDisksOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *QueryDisksResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryDisksResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryDisksOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateVolumeOperation structure represents the CreateVolume operation
type xxx_CreateVolumeOperation struct {
	This           *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Type           vds.VolumeType   `idl:"name:type" json:"type"`
	InputDiskArray []*vds.InputDisk `idl:"name:pInputDiskArray;size_is:(lNumberOfDisks)" json:"input_disk_array"`
	NumberOfDisks  int32            `idl:"name:lNumberOfDisks" json:"number_of_disks"`
	StripeSize     uint32           `idl:"name:ulStripeSize" json:"stripe_size"`
	Async          *vds.Async       `idl:"name:ppAsync" json:"async"`
	Return         int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateVolumeOperation) OpNum() int { return 7 }

func (o *xxx_CreateVolumeOperation) OpName() string { return "/IVdsPack/v0/CreateVolume" }

func (o *xxx_CreateVolumeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InputDiskArray != nil && o.NumberOfDisks == 0 {
		o.NumberOfDisks = int32(len(o.InputDiskArray))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVolumeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// type {in} (1:{alias=VDS_VOLUME_TYPE}(enum))
	{
		if err := w.WriteEnum(uint16(o.Type)); err != nil {
			return err
		}
	}
	// pInputDiskArray {in} (1:{pointer=ref}*(1))(2:{alias=VDS_INPUT_DISK}[dim:0,size_is=lNumberOfDisks](struct))
	{
		dimSize1 := uint64(o.NumberOfDisks)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.InputDiskArray {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.InputDiskArray[i1] != nil {
				if err := o.InputDiskArray[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&vds.InputDisk{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.InputDiskArray); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&vds.InputDisk{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lNumberOfDisks {in} (1:(int32))
	{
		if err := w.WriteData(o.NumberOfDisks); err != nil {
			return err
		}
	}
	// ulStripeSize {in} (1:(uint32))
	{
		if err := w.WriteData(o.StripeSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVolumeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// type {in} (1:{alias=VDS_VOLUME_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Type)); err != nil {
			return err
		}
	}
	// pInputDiskArray {in} (1:{pointer=ref}*(1))(2:{alias=VDS_INPUT_DISK}[dim:0,size_is=lNumberOfDisks](struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.InputDiskArray", sizeInfo[0])
		}
		o.InputDiskArray = make([]*vds.InputDisk, sizeInfo[0])
		for i1 := range o.InputDiskArray {
			i1 := i1
			if o.InputDiskArray[i1] == nil {
				o.InputDiskArray[i1] = &vds.InputDisk{}
			}
			if err := o.InputDiskArray[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lNumberOfDisks {in} (1:(int32))
	{
		if err := w.ReadData(&o.NumberOfDisks); err != nil {
			return err
		}
	}
	// ulStripeSize {in} (1:(uint32))
	{
		if err := w.ReadData(&o.StripeSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVolumeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVolumeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppAsync {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsAsync}(interface))
	{
		if o.Async != nil {
			_ptr_ppAsync := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Async != nil {
					if err := o.Async.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.Async{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Async, _ptr_ppAsync); err != nil {
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

func (o *xxx_CreateVolumeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppAsync {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsAsync}(interface))
	{
		_ptr_ppAsync := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Async == nil {
				o.Async = &vds.Async{}
			}
			if err := o.Async.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppAsync := func(ptr interface{}) { o.Async = *ptr.(**vds.Async) }
		if err := w.ReadPointer(&o.Async, _s_ppAsync, _ptr_ppAsync); err != nil {
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

// CreateVolumeRequest structure represents the CreateVolume operation request
type CreateVolumeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// type: A value from the VDS_VOLUME_TYPE enumeration that indicates the type of volume
	// to create.
	Type vds.VolumeType `idl:"name:type" json:"type"`
	// pInputDiskArray: An array of VDS_INPUT_DISK structures that indicate the disks on
	// which to create the volume.<80>
	InputDiskArray []*vds.InputDisk `idl:"name:pInputDiskArray;size_is:(lNumberOfDisks)" json:"input_disk_array"`
	// lNumberOfDisks: The number of elements in pInputDiskArray.
	NumberOfDisks int32 `idl:"name:lNumberOfDisks" json:"number_of_disks"`
	// ulStripeSize: The stripe size of the new volume.<81>
	StripeSize uint32 `idl:"name:ulStripeSize" json:"stripe_size"`
}

func (o *CreateVolumeRequest) xxx_ToOp(ctx context.Context, op *xxx_CreateVolumeOperation) *xxx_CreateVolumeOperation {
	if op == nil {
		op = &xxx_CreateVolumeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Type = o.Type
	op.InputDiskArray = o.InputDiskArray
	op.NumberOfDisks = o.NumberOfDisks
	op.StripeSize = o.StripeSize
	return op
}

func (o *CreateVolumeRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateVolumeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Type = op.Type
	o.InputDiskArray = op.InputDiskArray
	o.NumberOfDisks = op.NumberOfDisks
	o.StripeSize = op.StripeSize
}
func (o *CreateVolumeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateVolumeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateVolumeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateVolumeResponse structure represents the CreateVolume operation response
type CreateVolumeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync interface that, if the operation is successfully
	// completed, receives the IVdsAsync interface to monitor and control this operation.
	// Callers MUST release the interface when they are done with it. If the IVdsAsync::Wait
	// (Opnum 4) method is called on the interface, the interfaces returned in the VDS_ASYNC_OUTPUT
	// structure MUST be released as well. For information on handling asynchronous tasks,
	// see section 3.4.5.1.9.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The CreateVolume return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateVolumeResponse) xxx_ToOp(ctx context.Context, op *xxx_CreateVolumeOperation) *xxx_CreateVolumeOperation {
	if op == nil {
		op = &xxx_CreateVolumeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Async = o.Async
	op.Return = o.Return
	return op
}

func (o *CreateVolumeResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateVolumeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *CreateVolumeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateVolumeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateVolumeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddDiskOperation structure represents the AddDisk operation
type xxx_AddDiskOperation struct {
	This           *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat     `idl:"name:That" json:"that"`
	DiskID         *vds.ObjectID      `idl:"name:DiskId" json:"disk_id"`
	PartitionStyle vds.PartitionStyle `idl:"name:PartitionStyle" json:"partition_style"`
	AsHotSpare     int32              `idl:"name:bAsHotSpare" json:"as_hot_spare"`
	Return         int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_AddDiskOperation) OpNum() int { return 8 }

func (o *xxx_AddDiskOperation) OpName() string { return "/IVdsPack/v0/AddDisk" }

func (o *xxx_AddDiskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddDiskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// DiskId {in} (1:{alias=VDS_OBJECT_ID, names=GUID}(struct))
	{
		if o.DiskID != nil {
			if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.ObjectID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// PartitionStyle {in} (1:{alias=VDS_PARTITION_STYLE}(enum))
	{
		if err := w.WriteEnum(uint16(o.PartitionStyle)); err != nil {
			return err
		}
	}
	// bAsHotSpare {in} (1:(int32))
	{
		if err := w.WriteData(o.AsHotSpare); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddDiskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// DiskId {in} (1:{alias=VDS_OBJECT_ID, names=GUID}(struct))
	{
		if o.DiskID == nil {
			o.DiskID = &vds.ObjectID{}
		}
		if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// PartitionStyle {in} (1:{alias=VDS_PARTITION_STYLE}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.PartitionStyle)); err != nil {
			return err
		}
	}
	// bAsHotSpare {in} (1:(int32))
	{
		if err := w.ReadData(&o.AsHotSpare); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddDiskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddDiskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AddDiskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AddDiskRequest structure represents the AddDisk operation request
type AddDiskRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The VDS object ID of the disk object.
	DiskID *vds.ObjectID `idl:"name:DiskId" json:"disk_id"`
	// PartitionStyle: A value from the VDS_PARTITION_STYLE enumeration that indicates the
	// partition format.
	PartitionStyle vds.PartitionStyle `idl:"name:PartitionStyle" json:"partition_style"`
	// bAsHotSpare: The Virtual Disk Service Remote Protocol does not support this parameter;
	// callers MUST set it to FALSE.
	AsHotSpare int32 `idl:"name:bAsHotSpare" json:"as_hot_spare"`
}

func (o *AddDiskRequest) xxx_ToOp(ctx context.Context, op *xxx_AddDiskOperation) *xxx_AddDiskOperation {
	if op == nil {
		op = &xxx_AddDiskOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DiskID = o.DiskID
	op.PartitionStyle = o.PartitionStyle
	op.AsHotSpare = o.AsHotSpare
	return op
}

func (o *AddDiskRequest) xxx_FromOp(ctx context.Context, op *xxx_AddDiskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.PartitionStyle = op.PartitionStyle
	o.AsHotSpare = op.AsHotSpare
}
func (o *AddDiskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AddDiskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddDiskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddDiskResponse structure represents the AddDisk operation response
type AddDiskResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AddDisk return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddDiskResponse) xxx_ToOp(ctx context.Context, op *xxx_AddDiskOperation) *xxx_AddDiskOperation {
	if op == nil {
		op = &xxx_AddDiskOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *AddDiskResponse) xxx_FromOp(ctx context.Context, op *xxx_AddDiskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AddDiskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AddDiskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddDiskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MigrateDisksOperation structure represents the MigrateDisks operation
type xxx_MigrateDisksOperation struct {
	This          *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat  `idl:"name:That" json:"that"`
	DiskArray     []*vds.ObjectID `idl:"name:pDiskArray;size_is:(lNumberOfDisks)" json:"disk_array"`
	NumberOfDisks int32           `idl:"name:lNumberOfDisks" json:"number_of_disks"`
	TargetPack    *vds.ObjectID   `idl:"name:TargetPack" json:"target_pack"`
	Force         int32           `idl:"name:bForce" json:"force"`
	QueryOnly     int32           `idl:"name:bQueryOnly" json:"query_only"`
	Results       []int32         `idl:"name:pResults;size_is:(lNumberOfDisks)" json:"results"`
	RebootNeeded  int32           `idl:"name:pbRebootNeeded" json:"reboot_needed"`
	Return        int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_MigrateDisksOperation) OpNum() int { return 9 }

func (o *xxx_MigrateDisksOperation) OpName() string { return "/IVdsPack/v0/MigrateDisks" }

func (o *xxx_MigrateDisksOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.DiskArray != nil && o.NumberOfDisks == 0 {
		o.NumberOfDisks = int32(len(o.DiskArray))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MigrateDisksOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pDiskArray {in} (1:{pointer=ref}*(1))(2:{alias=VDS_OBJECT_ID, names=GUID}[dim:0,size_is=lNumberOfDisks](struct))
	{
		dimSize1 := uint64(o.NumberOfDisks)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.DiskArray {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.DiskArray[i1] != nil {
				if err := o.DiskArray[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&vds.ObjectID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.DiskArray); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&vds.ObjectID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lNumberOfDisks {in} (1:(int32))
	{
		if err := w.WriteData(o.NumberOfDisks); err != nil {
			return err
		}
	}
	// TargetPack {in} (1:{alias=VDS_OBJECT_ID, names=GUID}(struct))
	{
		if o.TargetPack != nil {
			if err := o.TargetPack.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.ObjectID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// bForce {in} (1:(int32))
	{
		if err := w.WriteData(o.Force); err != nil {
			return err
		}
	}
	// bQueryOnly {in} (1:(int32))
	{
		if err := w.WriteData(o.QueryOnly); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MigrateDisksOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pDiskArray {in} (1:{pointer=ref}*(1))(2:{alias=VDS_OBJECT_ID, names=GUID}[dim:0,size_is=lNumberOfDisks](struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.DiskArray", sizeInfo[0])
		}
		o.DiskArray = make([]*vds.ObjectID, sizeInfo[0])
		for i1 := range o.DiskArray {
			i1 := i1
			if o.DiskArray[i1] == nil {
				o.DiskArray[i1] = &vds.ObjectID{}
			}
			if err := o.DiskArray[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lNumberOfDisks {in} (1:(int32))
	{
		if err := w.ReadData(&o.NumberOfDisks); err != nil {
			return err
		}
	}
	// TargetPack {in} (1:{alias=VDS_OBJECT_ID, names=GUID}(struct))
	{
		if o.TargetPack == nil {
			o.TargetPack = &vds.ObjectID{}
		}
		if err := o.TargetPack.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// bForce {in} (1:(int32))
	{
		if err := w.ReadData(&o.Force); err != nil {
			return err
		}
	}
	// bQueryOnly {in} (1:(int32))
	{
		if err := w.ReadData(&o.QueryOnly); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MigrateDisksOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MigrateDisksOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pResults {out} (1:{pointer=ref}*(1))(2:{alias=HRESULT, names=LONG}[dim:0,size_is=lNumberOfDisks](int32))
	{
		dimSize1 := uint64(o.NumberOfDisks)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Results {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Results[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Results); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		}
	}
	// pbRebootNeeded {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.RebootNeeded); err != nil {
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

func (o *xxx_MigrateDisksOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pResults {out} (1:{pointer=ref}*(1))(2:{alias=HRESULT, names=LONG}[dim:0,size_is=lNumberOfDisks](int32))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Results", sizeInfo[0])
		}
		o.Results = make([]int32, sizeInfo[0])
		for i1 := range o.Results {
			i1 := i1
			if err := w.ReadData(&o.Results[i1]); err != nil {
				return err
			}
		}
	}
	// pbRebootNeeded {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.RebootNeeded); err != nil {
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

// MigrateDisksRequest structure represents the MigrateDisks operation request
type MigrateDisksRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pDiskArray: A pointer to an array of VDS object IDs--one for each disk object that
	// corresponds to the disks to migrate.
	DiskArray []*vds.ObjectID `idl:"name:pDiskArray;size_is:(lNumberOfDisks)" json:"disk_array"`
	// lNumberOfDisks: The number of disks specified in pDiskArray.
	NumberOfDisks int32 `idl:"name:lNumberOfDisks" json:"number_of_disks"`
	// TargetPack: The VDS object ID of the pack object.
	TargetPack *vds.ObjectID `idl:"name:TargetPack" json:"target_pack"`
	// bForce: A Boolean that determines whether disk migration is forced. When the client
	// makes the call to migrate disks, the provider(s) that owns the disks is notified
	// by the server that the disks are about to be migrated. The provider(s) can respond
	// to this notification with an error.
	Force int32 `idl:"name:bForce" json:"force"`
	// bQueryOnly: A Boolean that determines whether the disk migration will actually happen.
	QueryOnly int32 `idl:"name:bQueryOnly" json:"query_only"`
}

func (o *MigrateDisksRequest) xxx_ToOp(ctx context.Context, op *xxx_MigrateDisksOperation) *xxx_MigrateDisksOperation {
	if op == nil {
		op = &xxx_MigrateDisksOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DiskArray = o.DiskArray
	op.NumberOfDisks = o.NumberOfDisks
	op.TargetPack = o.TargetPack
	op.Force = o.Force
	op.QueryOnly = o.QueryOnly
	return op
}

func (o *MigrateDisksRequest) xxx_FromOp(ctx context.Context, op *xxx_MigrateDisksOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskArray = op.DiskArray
	o.NumberOfDisks = op.NumberOfDisks
	o.TargetPack = op.TargetPack
	o.Force = op.Force
	o.QueryOnly = op.QueryOnly
}
func (o *MigrateDisksRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MigrateDisksRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MigrateDisksOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MigrateDisksResponse structure represents the MigrateDisks operation response
type MigrateDisksResponse struct {
	// XXX: lNumberOfDisks is an implicit input depedency for output parameters
	NumberOfDisks int32 `idl:"name:lNumberOfDisks" json:"number_of_disks"`
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pResults: A pointer to an array of HRESULT values that, if the operation is successfully
	// completed, receives the HRESULTs returned by each disk migration request. There MUST
	// be one HRESULT value in the array for each disk in pDiskArray. If any of the disks
	// fail to migrate properly, the specific error code for that failure is received in
	// the corresponding entry in pResults.
	Results []int32 `idl:"name:pResults;size_is:(lNumberOfDisks)" json:"results"`
	// pbRebootNeeded: A pointer to a Boolean that, if the operation is successfully completed,
	// receives an indication of whether the user needs to reboot the remote machine in
	// order to complete the migration process.
	RebootNeeded int32 `idl:"name:pbRebootNeeded" json:"reboot_needed"`
	// Return: The MigrateDisks return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *MigrateDisksResponse) xxx_ToOp(ctx context.Context, op *xxx_MigrateDisksOperation) *xxx_MigrateDisksOperation {
	if op == nil {
		op = &xxx_MigrateDisksOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.NumberOfDisks == int32(0) {
		op.NumberOfDisks = o.NumberOfDisks
	}

	op.That = o.That
	op.Results = o.Results
	op.RebootNeeded = o.RebootNeeded
	op.Return = o.Return
	return op
}

func (o *MigrateDisksResponse) xxx_FromOp(ctx context.Context, op *xxx_MigrateDisksOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.NumberOfDisks = op.NumberOfDisks

	o.That = op.That
	o.Results = op.Results
	o.RebootNeeded = op.RebootNeeded
	o.Return = op.Return
}
func (o *MigrateDisksResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MigrateDisksResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MigrateDisksOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveMissingDiskOperation structure represents the RemoveMissingDisk operation
type xxx_RemoveMissingDiskOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	DiskID *vds.ObjectID  `idl:"name:DiskId" json:"disk_id"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveMissingDiskOperation) OpNum() int { return 11 }

func (o *xxx_RemoveMissingDiskOperation) OpName() string { return "/IVdsPack/v0/RemoveMissingDisk" }

func (o *xxx_RemoveMissingDiskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMissingDiskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// DiskId {in} (1:{alias=VDS_OBJECT_ID, names=GUID}(struct))
	{
		if o.DiskID != nil {
			if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.ObjectID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_RemoveMissingDiskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// DiskId {in} (1:{alias=VDS_OBJECT_ID, names=GUID}(struct))
	{
		if o.DiskID == nil {
			o.DiskID = &vds.ObjectID{}
		}
		if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMissingDiskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveMissingDiskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RemoveMissingDiskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RemoveMissingDiskRequest structure represents the RemoveMissingDisk operation request
type RemoveMissingDiskRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The VDS object ID of the disk object to remove.
	DiskID *vds.ObjectID `idl:"name:DiskId" json:"disk_id"`
}

func (o *RemoveMissingDiskRequest) xxx_ToOp(ctx context.Context, op *xxx_RemoveMissingDiskOperation) *xxx_RemoveMissingDiskOperation {
	if op == nil {
		op = &xxx_RemoveMissingDiskOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DiskID = o.DiskID
	return op
}

func (o *RemoveMissingDiskRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveMissingDiskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *RemoveMissingDiskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemoveMissingDiskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveMissingDiskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveMissingDiskResponse structure represents the RemoveMissingDisk operation response
type RemoveMissingDiskResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RemoveMissingDisk return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoveMissingDiskResponse) xxx_ToOp(ctx context.Context, op *xxx_RemoveMissingDiskOperation) *xxx_RemoveMissingDiskOperation {
	if op == nil {
		op = &xxx_RemoveMissingDiskOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *RemoveMissingDiskResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveMissingDiskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RemoveMissingDiskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemoveMissingDiskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveMissingDiskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RecoverOperation structure represents the Recover operation
type xxx_RecoverOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Async  *vds.Async     `idl:"name:ppAsync" json:"async"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RecoverOperation) OpNum() int { return 12 }

func (o *xxx_RecoverOperation) OpName() string { return "/IVdsPack/v0/Recover" }

func (o *xxx_RecoverOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RecoverOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RecoverOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_RecoverOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RecoverOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppAsync {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsAsync}(interface))
	{
		if o.Async != nil {
			_ptr_ppAsync := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Async != nil {
					if err := o.Async.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.Async{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Async, _ptr_ppAsync); err != nil {
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

func (o *xxx_RecoverOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppAsync {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsAsync}(interface))
	{
		_ptr_ppAsync := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Async == nil {
				o.Async = &vds.Async{}
			}
			if err := o.Async.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppAsync := func(ptr interface{}) { o.Async = *ptr.(**vds.Async) }
		if err := w.ReadPointer(&o.Async, _s_ppAsync, _ptr_ppAsync); err != nil {
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

// RecoverRequest structure represents the Recover operation request
type RecoverRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *RecoverRequest) xxx_ToOp(ctx context.Context, op *xxx_RecoverOperation) *xxx_RecoverOperation {
	if op == nil {
		op = &xxx_RecoverOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *RecoverRequest) xxx_FromOp(ctx context.Context, op *xxx_RecoverOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *RecoverRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RecoverRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RecoverOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RecoverResponse structure represents the Recover operation response
type RecoverResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync interface that, if the operation is successfully
	// completed, receives the IVdsAsync interface to monitor and control this operation.
	// Callers MUST release the interface when they are done with it.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The Recover return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RecoverResponse) xxx_ToOp(ctx context.Context, op *xxx_RecoverOperation) *xxx_RecoverOperation {
	if op == nil {
		op = &xxx_RecoverOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Async = o.Async
	op.Return = o.Return
	return op
}

func (o *RecoverResponse) xxx_FromOp(ctx context.Context, op *xxx_RecoverOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *RecoverResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RecoverResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RecoverOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
