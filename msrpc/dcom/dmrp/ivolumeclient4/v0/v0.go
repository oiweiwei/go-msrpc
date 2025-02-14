package ivolumeclient4

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
)

var (
	// import guard
	GoPackage = "dcom/dmrp"
)

var (
	// IVolumeClient4 interface identifier deb01010-3a37-4d26-99df-e2bb6ae3ac61
	VolumeClient4IID = &dcom.IID{Data1: 0xdeb01010, Data2: 0x3a37, Data3: 0x4d26, Data4: []byte{0x99, 0xdf, 0xe2, 0xbb, 0x6a, 0xe3, 0xac, 0x61}}
	// Syntax UUID
	VolumeClient4SyntaxUUID = &uuid.UUID{TimeLow: 0xdeb01010, TimeMid: 0x3a37, TimeHiAndVersion: 0x4d26, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0xdf, Node: [6]uint8{0xe2, 0xbb, 0x6a, 0xe3, 0xac, 0x61}}
	// Syntax ID
	VolumeClient4SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: VolumeClient4SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVolumeClient4 interface.
type VolumeClient4Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The RefreshEx method refreshes the server's cache of storage objects, including regions,
	// removable media and CD-ROM drive media, file systems, and drive letters.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF]; see also section 2.2.1
	// for HRESULT values predefined by the Disk Management Remote Protocol).
	//
	// The handling of this message is identical to the handling of IVolumeClient::Refresh
	// and IVolumeClient3::Refresh except that the server MUST perform an extra low-level
	// refresh of the list of storage objects by looking for missing dynamic disks or dynamic
	// disks that were missing and are now present. This verification updates the status
	// for missing disks, volumes that reside on missing disks, or disk regions that reside
	// on missing disks.<228>
	//
	// In addition to the preceding actions, the server MUST check whether the lengths of
	// the disks have changed and make appropriate changes to the disk objects in the list
	// of storage objects.
	RefreshEx(context.Context, *RefreshExRequest, ...dcerpc.CallOption) (*RefreshExResponse, error)

	// The GetVolumeDeviceName method retrieves the Windows NT operating system device name
	// of a dynamic volume on the server.
	//
	// Return Values: The method MUST return 0 or a nonerror HRESULT on success, or an implementation-specific
	// nonzero error code on failure (as specified in [MS-ERREF]; see also section 2.2.1
	// for HRESULT values predefined by the Disk Management Remote Protocol).
	GetVolumeDeviceName(context.Context, *GetVolumeDeviceNameRequest, ...dcerpc.CallOption) (*GetVolumeDeviceNameResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) VolumeClient4Client
}

type xxx_DefaultVolumeClient4Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultVolumeClient4Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultVolumeClient4Client) RefreshEx(ctx context.Context, in *RefreshExRequest, opts ...dcerpc.CallOption) (*RefreshExResponse, error) {
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
	out := &RefreshExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClient4Client) GetVolumeDeviceName(ctx context.Context, in *GetVolumeDeviceNameRequest, opts ...dcerpc.CallOption) (*GetVolumeDeviceNameResponse, error) {
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
	out := &GetVolumeDeviceNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClient4Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultVolumeClient4Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultVolumeClient4Client) IPID(ctx context.Context, ipid *dcom.IPID) VolumeClient4Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultVolumeClient4Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewVolumeClient4Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (VolumeClient4Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(VolumeClient4SyntaxV0_0))...)
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
	return &xxx_DefaultVolumeClient4Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_RefreshExOperation structure represents the RefreshEx operation
type xxx_RefreshExOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RefreshExOperation) OpNum() int { return 3 }

func (o *xxx_RefreshExOperation) OpName() string { return "/IVolumeClient4/v0/RefreshEx" }

func (o *xxx_RefreshExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RefreshExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_RefreshExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RefreshExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RefreshExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RefreshExRequest structure represents the RefreshEx operation request
type RefreshExRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *RefreshExRequest) xxx_ToOp(ctx context.Context, op *xxx_RefreshExOperation) *xxx_RefreshExOperation {
	if op == nil {
		op = &xxx_RefreshExOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *RefreshExRequest) xxx_FromOp(ctx context.Context, op *xxx_RefreshExOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *RefreshExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RefreshExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RefreshExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RefreshExResponse structure represents the RefreshEx operation response
type RefreshExResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The RefreshEx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RefreshExResponse) xxx_ToOp(ctx context.Context, op *xxx_RefreshExOperation) *xxx_RefreshExOperation {
	if op == nil {
		op = &xxx_RefreshExOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *RefreshExResponse) xxx_FromOp(ctx context.Context, op *xxx_RefreshExOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *RefreshExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RefreshExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RefreshExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetVolumeDeviceNameOperation structure represents the GetVolumeDeviceName operation
type xxx_GetVolumeDeviceNameOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	VolumeID           int64          `idl:"name:_volumeId" json:"volume_id"`
	VolumeDeviceLength uint32         `idl:"name:cchVolumeDevice" json:"volume_device_length"`
	VolumeDevice       string         `idl:"name:pwszVolumeDevice;size_is:(, cchVolumeDevice)" json:"volume_device"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetVolumeDeviceNameOperation) OpNum() int { return 4 }

func (o *xxx_GetVolumeDeviceNameOperation) OpName() string {
	return "/IVolumeClient4/v0/GetVolumeDeviceName"
}

func (o *xxx_GetVolumeDeviceNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVolumeDeviceNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// _volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.WriteData(o.VolumeID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVolumeDeviceNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// _volumeId {in} (1:{alias=LdmObjectId, names=LONGLONG}(int64))
	{
		if err := w.ReadData(&o.VolumeID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVolumeDeviceNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.VolumeDevice != "" && o.VolumeDeviceLength == 0 {
		o.VolumeDeviceLength = uint32(len(o.VolumeDevice))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVolumeDeviceNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// cchVolumeDevice {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.VolumeDeviceLength); err != nil {
			return err
		}
	}
	// pwszVolumeDevice {out} (1:{pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,size_is=cchVolumeDevice,string](wchar))
	{
		if o.VolumeDevice != "" || o.VolumeDeviceLength > 0 {
			_ptr_pwszVolumeDevice := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.VolumeDeviceLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				_VolumeDevice_buf := utf16.Encode([]rune(o.VolumeDevice))
				if uint64(len(_VolumeDevice_buf)) > sizeInfo[0] {
					_VolumeDevice_buf = _VolumeDevice_buf[:sizeInfo[0]]
				}
				for i1 := range _VolumeDevice_buf {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(_VolumeDevice_buf[i1]); err != nil {
						return err
					}
				}
				for i1 := len(_VolumeDevice_buf); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint16(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.VolumeDevice, _ptr_pwszVolumeDevice); err != nil {
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

func (o *xxx_GetVolumeDeviceNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// cchVolumeDevice {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.VolumeDeviceLength); err != nil {
			return err
		}
	}
	// pwszVolumeDevice {out} (1:{pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,size_is=cchVolumeDevice,string](wchar))
	{
		_ptr_pwszVolumeDevice := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			var _VolumeDevice_buf []uint16
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _VolumeDevice_buf", sizeInfo[0])
			}
			_VolumeDevice_buf = make([]uint16, sizeInfo[0])
			for i1 := range _VolumeDevice_buf {
				i1 := i1
				if err := w.ReadData(&_VolumeDevice_buf[i1]); err != nil {
					return err
				}
			}
			o.VolumeDevice = strings.TrimRight(string(utf16.Decode(_VolumeDevice_buf)), ndr.ZeroString)
			return nil
		})
		_s_pwszVolumeDevice := func(ptr interface{}) { o.VolumeDevice = *ptr.(*string) }
		if err := w.ReadPointer(&o.VolumeDevice, _s_pwszVolumeDevice, _ptr_pwszVolumeDevice); err != nil {
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

// GetVolumeDeviceNameRequest structure represents the GetVolumeDeviceName operation request
type GetVolumeDeviceNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// _volumeId: Specifies the OID of the volume whose path name is being returned.
	VolumeID int64 `idl:"name:_volumeId" json:"volume_id"`
}

func (o *GetVolumeDeviceNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetVolumeDeviceNameOperation) *xxx_GetVolumeDeviceNameOperation {
	if op == nil {
		op = &xxx_GetVolumeDeviceNameOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
	return op
}

func (o *GetVolumeDeviceNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetVolumeDeviceNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
}
func (o *GetVolumeDeviceNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetVolumeDeviceNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVolumeDeviceNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetVolumeDeviceNameResponse structure represents the GetVolumeDeviceName operation response
type GetVolumeDeviceNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// cchVolumeDevice: Number of characters returned in pwszVolumeDevice, including the
	// terminating null character.
	VolumeDeviceLength uint32 `idl:"name:cchVolumeDevice" json:"volume_device_length"`
	// pwszVolumeDevice: Pointer to a null-terminated array of characters that stores the
	// Windows NT device name of the volume specified by volumeId. The device name is in
	// the format \Device\DeviceName. Memory for the array is allocated by the server and
	// freed by the client.
	VolumeDevice string `idl:"name:pwszVolumeDevice;size_is:(, cchVolumeDevice)" json:"volume_device"`
	// Return: The GetVolumeDeviceName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetVolumeDeviceNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetVolumeDeviceNameOperation) *xxx_GetVolumeDeviceNameOperation {
	if op == nil {
		op = &xxx_GetVolumeDeviceNameOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.VolumeDeviceLength = op.VolumeDeviceLength
	o.VolumeDevice = op.VolumeDevice
	o.Return = op.Return
	return op
}

func (o *GetVolumeDeviceNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetVolumeDeviceNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.VolumeDeviceLength = op.VolumeDeviceLength
	o.VolumeDevice = op.VolumeDevice
	o.Return = op.Return
}
func (o *GetVolumeDeviceNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetVolumeDeviceNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVolumeDeviceNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
