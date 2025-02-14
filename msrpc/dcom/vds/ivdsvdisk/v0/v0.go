package ivdsvdisk

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
	// IVdsVDisk interface identifier 1e062b84-e5e6-4b4b-8a25-67b81e8f13e8
	VDiskIID = &dcom.IID{Data1: 0x1e062b84, Data2: 0xe5e6, Data3: 0x4b4b, Data4: []byte{0x8a, 0x25, 0x67, 0xb8, 0x1e, 0x8f, 0x13, 0xe8}}
	// Syntax UUID
	VDiskSyntaxUUID = &uuid.UUID{TimeLow: 0x1e062b84, TimeMid: 0xe5e6, TimeHiAndVersion: 0x4b4b, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x25, Node: [6]uint8{0x67, 0xb8, 0x1e, 0x8f, 0x13, 0xe8}}
	// Syntax ID
	VDiskSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: VDiskSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsVDisk interface.
type VDiskClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The Open method opens a handle to the specified virtual disk file and returns an
	// IVdsOpenVDisk (section 3.1.15.2) interface pointer to an object representing the
	// open virtual disk (an OpenVirtualDisk object). Release the IVdsOpenVDisk interface
	// to close the handle to the virtual disk.
	//
	// Return Values: The method MUST return zero to indicate success, or return an implementation-specific
	// nonzero error code to indicate failure.
	Open(context.Context, *OpenRequest, ...dcerpc.CallOption) (*OpenResponse, error)

	// GetProperties operation.
	GetProperties(context.Context, *GetPropertiesRequest, ...dcerpc.CallOption) (*GetPropertiesResponse, error)

	// The GetHostVolume method returns an interface pointer to the volume object for the
	// volume on which the virtual disk backing store file resides.<139>
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetHostVolume(context.Context, *GetHostVolumeRequest, ...dcerpc.CallOption) (*GetHostVolumeResponse, error)

	// The GetDeviceName method returns the device name of the disk.
	//
	// Return Values: The method MUST return zero to indicate success, or return an implementation-specific
	// nonzero error code to indicate failure.
	GetDeviceName(context.Context, *GetDeviceNameRequest, ...dcerpc.CallOption) (*GetDeviceNameResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) VDiskClient
}

type xxx_DefaultVDiskClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultVDiskClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultVDiskClient) Open(ctx context.Context, in *OpenRequest, opts ...dcerpc.CallOption) (*OpenResponse, error) {
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
	out := &OpenResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVDiskClient) GetProperties(ctx context.Context, in *GetPropertiesRequest, opts ...dcerpc.CallOption) (*GetPropertiesResponse, error) {
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

func (o *xxx_DefaultVDiskClient) GetHostVolume(ctx context.Context, in *GetHostVolumeRequest, opts ...dcerpc.CallOption) (*GetHostVolumeResponse, error) {
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
	out := &GetHostVolumeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVDiskClient) GetDeviceName(ctx context.Context, in *GetDeviceNameRequest, opts ...dcerpc.CallOption) (*GetDeviceNameResponse, error) {
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
	out := &GetDeviceNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVDiskClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultVDiskClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultVDiskClient) IPID(ctx context.Context, ipid *dcom.IPID) VDiskClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultVDiskClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewVDiskClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (VDiskClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(VDiskSyntaxV0_0))...)
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
	return &xxx_DefaultVDiskClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_OpenOperation structure represents the Open operation
type xxx_OpenOperation struct {
	This           *dcom.ORPCThis            `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat            `idl:"name:That" json:"that"`
	AccessMask     vds.VirtualDiskAccessMask `idl:"name:AccessMask" json:"access_mask"`
	Flags          vds.OpenVirtualDiskFlag   `idl:"name:Flags" json:"flags"`
	ReadWriteDepth uint32                    `idl:"name:ReadWriteDepth" json:"read_write_depth"`
	OpenVDisk      *vds.OpenVDisk            `idl:"name:ppOpenVDisk" json:"open_v_disk"`
	Return         int32                     `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenOperation) OpNum() int { return 3 }

func (o *xxx_OpenOperation) OpName() string { return "/IVdsVDisk/v0/Open" }

func (o *xxx_OpenOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// AccessMask {in} (1:{alias=VIRTUAL_DISK_ACCESS_MASK}(enum))
	{
		if err := w.WriteEnum(uint16(o.AccessMask)); err != nil {
			return err
		}
	}
	// Flags {in} (1:{alias=OPEN_VIRTUAL_DISK_FLAG}(enum))
	{
		if err := w.WriteEnum(uint16(o.Flags)); err != nil {
			return err
		}
	}
	// ReadWriteDepth {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.ReadWriteDepth); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// AccessMask {in} (1:{alias=VIRTUAL_DISK_ACCESS_MASK}(enum))
	{
		_eAccessMask := uint16(o.AccessMask)
		if err := w.ReadEnum(&_eAccessMask); err != nil {
			return err
		}
		o.AccessMask = vds.VirtualDiskAccessMask(_eAccessMask)
	}
	// Flags {in} (1:{alias=OPEN_VIRTUAL_DISK_FLAG}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Flags)); err != nil {
			return err
		}
	}
	// ReadWriteDepth {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.ReadWriteDepth); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppOpenVDisk {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsOpenVDisk}(interface))
	{
		if o.OpenVDisk != nil {
			_ptr_ppOpenVDisk := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.OpenVDisk != nil {
					if err := o.OpenVDisk.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.OpenVDisk{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.OpenVDisk, _ptr_ppOpenVDisk); err != nil {
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

func (o *xxx_OpenOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppOpenVDisk {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsOpenVDisk}(interface))
	{
		_ptr_ppOpenVDisk := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.OpenVDisk == nil {
				o.OpenVDisk = &vds.OpenVDisk{}
			}
			if err := o.OpenVDisk.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppOpenVDisk := func(ptr interface{}) { o.OpenVDisk = *ptr.(**vds.OpenVDisk) }
		if err := w.ReadPointer(&o.OpenVDisk, _s_ppOpenVDisk, _ptr_ppOpenVDisk); err != nil {
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

// OpenRequest structure represents the Open operation request
type OpenRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// AccessMask: A VIRTUAL_DISK_ACCESS_MASK (section 2.2.2.19.1.4) structure that contains
	// the set of access rights to be applied to the opened virtual disk.
	AccessMask vds.VirtualDiskAccessMask `idl:"name:AccessMask" json:"access_mask"`
	// Flags: A bitmask of OPEN_VIRTUAL_DISK_FLAG (section 2.2.2.19.1.2) flags specifying
	// how the virtual disk is to be opened.
	Flags vds.OpenVirtualDiskFlag `idl:"name:Flags" json:"flags"`
	// ReadWriteDepth: This is applicable only to differencing type virtual disks. The number
	// of backing stores (files) to open read/write. This count includes the child. The
	// remaining stores in the differencing chain MUST be opened as read-only. For example,
	// given a differencing disk with two parents (diskA is the differencing disk whose
	// parent is diskB, and since diskB is a differencing disk, it has a parent which is
	// diskC), entering '2' for this parameter will open the differencing disk (diskA) and
	// the parent used to create this differencing disk (diskB) as read-write. In this case,
	// diskB is also a differencing disk and its parent (diskC) is opened as read-only.
	ReadWriteDepth uint32 `idl:"name:ReadWriteDepth" json:"read_write_depth"`
}

func (o *OpenRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenOperation) *xxx_OpenOperation {
	if op == nil {
		op = &xxx_OpenOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.AccessMask = o.AccessMask
	op.Flags = o.Flags
	op.ReadWriteDepth = o.ReadWriteDepth
	return op
}

func (o *OpenRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.AccessMask = op.AccessMask
	o.Flags = op.Flags
	o.ReadWriteDepth = op.ReadWriteDepth
}
func (o *OpenRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenResponse structure represents the Open operation response
type OpenResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	OpenVDisk *vds.OpenVDisk `idl:"name:ppOpenVDisk" json:"open_v_disk"`
	// Return: The Open return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenOperation) *xxx_OpenOperation {
	if op == nil {
		op = &xxx_OpenOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.OpenVDisk = o.OpenVDisk
	op.Return = o.Return
	return op
}

func (o *OpenResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.OpenVDisk = op.OpenVDisk
	o.Return = op.Return
}
func (o *OpenResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPropertiesOperation structure represents the GetProperties operation
type xxx_GetPropertiesOperation struct {
	This           *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat       `idl:"name:That" json:"that"`
	DiskProperties *vds.VDiskProperties `idl:"name:pDiskProperties" json:"disk_properties"`
	Return         int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertiesOperation) OpNum() int { return 4 }

func (o *xxx_GetPropertiesOperation) OpName() string { return "/IVdsVDisk/v0/GetProperties" }

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
	// pDiskProperties {out} (1:{alias=PVDS_VDISK_PROPERTIES}*(1))(2:{alias=VDS_VDISK_PROPERTIES}(struct))
	{
		if o.DiskProperties != nil {
			if err := o.DiskProperties.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.VDiskProperties{}).MarshalNDR(ctx, w); err != nil {
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
	// pDiskProperties {out} (1:{alias=PVDS_VDISK_PROPERTIES,pointer=ref}*(1))(2:{alias=VDS_VDISK_PROPERTIES}(struct))
	{
		if o.DiskProperties == nil {
			o.DiskProperties = &vds.VDiskProperties{}
		}
		if err := o.DiskProperties.UnmarshalNDR(ctx, w); err != nil {
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
	That           *dcom.ORPCThat       `idl:"name:That" json:"that"`
	DiskProperties *vds.VDiskProperties `idl:"name:pDiskProperties" json:"disk_properties"`
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
	op.DiskProperties = o.DiskProperties
	op.Return = o.Return
	return op
}

func (o *GetPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DiskProperties = op.DiskProperties
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

// xxx_GetHostVolumeOperation structure represents the GetHostVolume operation
type xxx_GetHostVolumeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Volume *vds.Volume    `idl:"name:ppVolume" json:"volume"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetHostVolumeOperation) OpNum() int { return 5 }

func (o *xxx_GetHostVolumeOperation) OpName() string { return "/IVdsVDisk/v0/GetHostVolume" }

func (o *xxx_GetHostVolumeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHostVolumeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetHostVolumeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetHostVolumeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHostVolumeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppVolume {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsVolume}(interface))
	{
		if o.Volume != nil {
			_ptr_ppVolume := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Volume != nil {
					if err := o.Volume.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.Volume{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Volume, _ptr_ppVolume); err != nil {
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

func (o *xxx_GetHostVolumeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppVolume {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsVolume}(interface))
	{
		_ptr_ppVolume := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Volume == nil {
				o.Volume = &vds.Volume{}
			}
			if err := o.Volume.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppVolume := func(ptr interface{}) { o.Volume = *ptr.(**vds.Volume) }
		if err := w.ReadPointer(&o.Volume, _s_ppVolume, _ptr_ppVolume); err != nil {
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

// GetHostVolumeRequest structure represents the GetHostVolume operation request
type GetHostVolumeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetHostVolumeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetHostVolumeOperation) *xxx_GetHostVolumeOperation {
	if op == nil {
		op = &xxx_GetHostVolumeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetHostVolumeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetHostVolumeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetHostVolumeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetHostVolumeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetHostVolumeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetHostVolumeResponse structure represents the GetHostVolume operation response
type GetHostVolumeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Volume *vds.Volume    `idl:"name:ppVolume" json:"volume"`
	// Return: The GetHostVolume return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetHostVolumeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetHostVolumeOperation) *xxx_GetHostVolumeOperation {
	if op == nil {
		op = &xxx_GetHostVolumeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Volume = o.Volume
	op.Return = o.Return
	return op
}

func (o *GetHostVolumeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetHostVolumeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Volume = op.Volume
	o.Return = op.Return
}
func (o *GetHostVolumeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetHostVolumeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetHostVolumeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDeviceNameOperation structure represents the GetDeviceName operation
type xxx_GetDeviceNameOperation struct {
	This       *dcom.ORPCThis `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat `idl:"name:That" json:"that"`
	DeviceName string         `idl:"name:ppDeviceName;string" json:"device_name"`
	Return     int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDeviceNameOperation) OpNum() int { return 6 }

func (o *xxx_GetDeviceNameOperation) OpName() string { return "/IVdsVDisk/v0/GetDeviceName" }

func (o *xxx_GetDeviceNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDeviceNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDeviceNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDeviceNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDeviceNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppDeviceName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.DeviceName != "" {
			_ptr_ppDeviceName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.DeviceName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DeviceName, _ptr_ppDeviceName); err != nil {
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

func (o *xxx_GetDeviceNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppDeviceName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ppDeviceName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.DeviceName); err != nil {
				return err
			}
			return nil
		})
		_s_ppDeviceName := func(ptr interface{}) { o.DeviceName = *ptr.(*string) }
		if err := w.ReadPointer(&o.DeviceName, _s_ppDeviceName, _ptr_ppDeviceName); err != nil {
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

// GetDeviceNameRequest structure represents the GetDeviceName operation request
type GetDeviceNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetDeviceNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDeviceNameOperation) *xxx_GetDeviceNameOperation {
	if op == nil {
		op = &xxx_GetDeviceNameOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetDeviceNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDeviceNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetDeviceNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDeviceNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDeviceNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDeviceNameResponse structure represents the GetDeviceName operation response
type GetDeviceNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppDeviceName: A pointer to a variable that receives the device name of the disk.
	DeviceName string `idl:"name:ppDeviceName;string" json:"device_name"`
	// Return: The GetDeviceName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDeviceNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDeviceNameOperation) *xxx_GetDeviceNameOperation {
	if op == nil {
		op = &xxx_GetDeviceNameOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.DeviceName = o.DeviceName
	op.Return = o.Return
	return op
}

func (o *GetDeviceNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDeviceNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DeviceName = op.DeviceName
	o.Return = op.Return
}
func (o *GetDeviceNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDeviceNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDeviceNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
