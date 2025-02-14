package ivdsvolumeplex

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
	// IVdsVolumePlex interface identifier 4daa0135-e1d1-40f1-aaa5-3cc1e53221c3
	VolumePlexIID = &dcom.IID{Data1: 0x4daa0135, Data2: 0xe1d1, Data3: 0x40f1, Data4: []byte{0xaa, 0xa5, 0x3c, 0xc1, 0xe5, 0x32, 0x21, 0xc3}}
	// Syntax UUID
	VolumePlexSyntaxUUID = &uuid.UUID{TimeLow: 0x4daa0135, TimeMid: 0xe1d1, TimeHiAndVersion: 0x40f1, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0xa5, Node: [6]uint8{0x3c, 0xc1, 0xe5, 0x32, 0x21, 0xc3}}
	// Syntax ID
	VolumePlexSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: VolumePlexSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsVolumePlex interface.
type VolumePlexClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// GetProperties operation.
	GetProperties(context.Context, *GetPropertiesRequest, ...dcerpc.CallOption) (*GetPropertiesResponse, error)

	// The GetVolume method retrieves the volume that the volume plex belongs to.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetVolume(context.Context, *GetVolumeRequest, ...dcerpc.CallOption) (*GetVolumeResponse, error)

	// QueryExtents operation.
	QueryExtents(context.Context, *QueryExtentsRequest, ...dcerpc.CallOption) (*QueryExtentsResponse, error)

	// The Repair method repairs a fault-tolerant volume plex by moving defective members
	// to good disks. Only plexes that are RAID-5, striped with parity, can be repaired
	// with this method.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// ERROR_SUCCESS (0x00000000)
	Repair(context.Context, *RepairRequest, ...dcerpc.CallOption) (*RepairResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) VolumePlexClient
}

type xxx_DefaultVolumePlexClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultVolumePlexClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultVolumePlexClient) GetProperties(ctx context.Context, in *GetPropertiesRequest, opts ...dcerpc.CallOption) (*GetPropertiesResponse, error) {
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

func (o *xxx_DefaultVolumePlexClient) GetVolume(ctx context.Context, in *GetVolumeRequest, opts ...dcerpc.CallOption) (*GetVolumeResponse, error) {
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
	out := &GetVolumeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumePlexClient) QueryExtents(ctx context.Context, in *QueryExtentsRequest, opts ...dcerpc.CallOption) (*QueryExtentsResponse, error) {
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
	out := &QueryExtentsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumePlexClient) Repair(ctx context.Context, in *RepairRequest, opts ...dcerpc.CallOption) (*RepairResponse, error) {
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
	out := &RepairResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumePlexClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultVolumePlexClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultVolumePlexClient) IPID(ctx context.Context, ipid *dcom.IPID) VolumePlexClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultVolumePlexClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewVolumePlexClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (VolumePlexClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(VolumePlexSyntaxV0_0))...)
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
	return &xxx_DefaultVolumePlexClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetPropertiesOperation structure represents the GetProperties operation
type xxx_GetPropertiesOperation struct {
	This           *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat          `idl:"name:That" json:"that"`
	PlexProperties *vds.VolumePlexProperty `idl:"name:pPlexProperties" json:"plex_properties"`
	Return         int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertiesOperation) OpNum() int { return 3 }

func (o *xxx_GetPropertiesOperation) OpName() string { return "/IVdsVolumePlex/v0/GetProperties" }

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
	// pPlexProperties {out} (1:{pointer=ref}*(1))(2:{alias=VDS_VOLUME_PLEX_PROP}(struct))
	{
		if o.PlexProperties != nil {
			if err := o.PlexProperties.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.VolumePlexProperty{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
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
	// pPlexProperties {out} (1:{pointer=ref}*(1))(2:{alias=VDS_VOLUME_PLEX_PROP}(struct))
	{
		if o.PlexProperties == nil {
			o.PlexProperties = &vds.VolumePlexProperty{}
		}
		if err := o.PlexProperties.UnmarshalNDR(ctx, w); err != nil {
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
	o.This = op.This
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
	That           *dcom.ORPCThat          `idl:"name:That" json:"that"`
	PlexProperties *vds.VolumePlexProperty `idl:"name:pPlexProperties" json:"plex_properties"`
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
	o.That = op.That
	o.PlexProperties = op.PlexProperties
	o.Return = op.Return
	return op
}

func (o *GetPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PlexProperties = op.PlexProperties
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

// xxx_GetVolumeOperation structure represents the GetVolume operation
type xxx_GetVolumeOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Volume *vds.Volume    `idl:"name:ppVolume" json:"volume"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetVolumeOperation) OpNum() int { return 4 }

func (o *xxx_GetVolumeOperation) OpName() string { return "/IVdsVolumePlex/v0/GetVolume" }

func (o *xxx_GetVolumeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVolumeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetVolumeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetVolumeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVolumeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetVolumeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// GetVolumeRequest structure represents the GetVolume operation request
type GetVolumeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetVolumeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetVolumeOperation) *xxx_GetVolumeOperation {
	if op == nil {
		op = &xxx_GetVolumeOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetVolumeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetVolumeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetVolumeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetVolumeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVolumeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetVolumeResponse structure represents the GetVolume operation response
type GetVolumeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppVolume: A pointer to an IVdsVolume interface that, if the operation is successfully
	// completed, receives the IVdsVolume interface of the volume object that the volume
	// plex belongs to. Callers MUST release the interface when they are done with it.
	Volume *vds.Volume `idl:"name:ppVolume" json:"volume"`
	// Return: The GetVolume return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetVolumeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetVolumeOperation) *xxx_GetVolumeOperation {
	if op == nil {
		op = &xxx_GetVolumeOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Volume = op.Volume
	o.Return = op.Return
	return op
}

func (o *GetVolumeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetVolumeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Volume = op.Volume
	o.Return = op.Return
}
func (o *GetVolumeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetVolumeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVolumeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryExtentsOperation structure represents the QueryExtents operation
type xxx_QueryExtentsOperation struct {
	This            *dcom.ORPCThis    `idl:"name:This" json:"this"`
	That            *dcom.ORPCThat    `idl:"name:That" json:"that"`
	ExtentArray     []*vds.DiskExtent `idl:"name:ppExtentArray;size_is:(, plNumberOfExtents)" json:"extent_array"`
	NumberOfExtents int32             `idl:"name:plNumberOfExtents" json:"number_of_extents"`
	Return          int32             `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryExtentsOperation) OpNum() int { return 5 }

func (o *xxx_QueryExtentsOperation) OpName() string { return "/IVdsVolumePlex/v0/QueryExtents" }

func (o *xxx_QueryExtentsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryExtentsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryExtentsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_QueryExtentsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ExtentArray != nil && o.NumberOfExtents == 0 {
		o.NumberOfExtents = int32(len(o.ExtentArray))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryExtentsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppExtentArray {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VDS_DISK_EXTENT}[dim:0,size_is=plNumberOfExtents](struct))
	{
		if o.ExtentArray != nil || o.NumberOfExtents > 0 {
			_ptr_ppExtentArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NumberOfExtents)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ExtentArray {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.ExtentArray[i1] != nil {
						if err := o.ExtentArray[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&vds.DiskExtent{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.ExtentArray); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&vds.DiskExtent{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ExtentArray, _ptr_ppExtentArray); err != nil {
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
	// plNumberOfExtents {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.NumberOfExtents); err != nil {
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

func (o *xxx_QueryExtentsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppExtentArray {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VDS_DISK_EXTENT}[dim:0,size_is=plNumberOfExtents](struct))
	{
		_ptr_ppExtentArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ExtentArray", sizeInfo[0])
			}
			o.ExtentArray = make([]*vds.DiskExtent, sizeInfo[0])
			for i1 := range o.ExtentArray {
				i1 := i1
				if o.ExtentArray[i1] == nil {
					o.ExtentArray[i1] = &vds.DiskExtent{}
				}
				if err := o.ExtentArray[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppExtentArray := func(ptr interface{}) { o.ExtentArray = *ptr.(*[]*vds.DiskExtent) }
		if err := w.ReadPointer(&o.ExtentArray, _s_ppExtentArray, _ptr_ppExtentArray); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// plNumberOfExtents {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.NumberOfExtents); err != nil {
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

// QueryExtentsRequest structure represents the QueryExtents operation request
type QueryExtentsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *QueryExtentsRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryExtentsOperation) *xxx_QueryExtentsOperation {
	if op == nil {
		op = &xxx_QueryExtentsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *QueryExtentsRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryExtentsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *QueryExtentsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryExtentsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryExtentsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryExtentsResponse structure represents the QueryExtents operation response
type QueryExtentsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That            *dcom.ORPCThat    `idl:"name:That" json:"that"`
	ExtentArray     []*vds.DiskExtent `idl:"name:ppExtentArray;size_is:(, plNumberOfExtents)" json:"extent_array"`
	NumberOfExtents int32             `idl:"name:plNumberOfExtents" json:"number_of_extents"`
	// Return: The QueryExtents return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryExtentsResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryExtentsOperation) *xxx_QueryExtentsOperation {
	if op == nil {
		op = &xxx_QueryExtentsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.ExtentArray = op.ExtentArray
	o.NumberOfExtents = op.NumberOfExtents
	o.Return = op.Return
	return op
}

func (o *QueryExtentsResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryExtentsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ExtentArray = op.ExtentArray
	o.NumberOfExtents = op.NumberOfExtents
	o.Return = op.Return
}
func (o *QueryExtentsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryExtentsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryExtentsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RepairOperation structure represents the Repair operation
type xxx_RepairOperation struct {
	This           *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat   `idl:"name:That" json:"that"`
	InputDiskArray []*vds.InputDisk `idl:"name:pInputDiskArray;size_is:(lNumberOfDisks)" json:"input_disk_array"`
	NumberOfDisks  int32            `idl:"name:lNumberOfDisks" json:"number_of_disks"`
	Async          *vds.Async       `idl:"name:ppAsync" json:"async"`
	Return         int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_RepairOperation) OpNum() int { return 6 }

func (o *xxx_RepairOperation) OpName() string { return "/IVdsVolumePlex/v0/Repair" }

func (o *xxx_RepairOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
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

func (o *xxx_RepairOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_RepairOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_RepairOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RepairOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RepairOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RepairRequest structure represents the Repair operation request
type RepairRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pInputDiskArray: An array of VDS_INPUT_DISK structures that describe the replacement
	// disks. Only diskId and ullSize SHOULD be specified in each VDS_INPUT_DISK element.
	// Only one new disk can be passed to this method at a time.
	InputDiskArray []*vds.InputDisk `idl:"name:pInputDiskArray;size_is:(lNumberOfDisks)" json:"input_disk_array"`
	// lNumberOfDisks: The number of elements in pInputDiskArray. Only one new disk can
	// be passed to this method at a time.
	NumberOfDisks int32 `idl:"name:lNumberOfDisks" json:"number_of_disks"`
}

func (o *RepairRequest) xxx_ToOp(ctx context.Context, op *xxx_RepairOperation) *xxx_RepairOperation {
	if op == nil {
		op = &xxx_RepairOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.InputDiskArray = op.InputDiskArray
	o.NumberOfDisks = op.NumberOfDisks
	return op
}

func (o *RepairRequest) xxx_FromOp(ctx context.Context, op *xxx_RepairOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.InputDiskArray = op.InputDiskArray
	o.NumberOfDisks = op.NumberOfDisks
}
func (o *RepairRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RepairRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RepairOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RepairResponse structure represents the Repair operation response
type RepairResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync interface that, if the operation is successfully
	// completed, receives the IVdsAsync interface to monitor and control this operation.
	// Callers MUST release the interface when they are done with it. If the Wait method
	// is called on the interface, the interface returned in the VDS_ASYNC_OUTPUT structure
	// MUST be released as well. For information on asynchronous tasks, see section 3.4.5.1.9.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The Repair return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RepairResponse) xxx_ToOp(ctx context.Context, op *xxx_RepairOperation) *xxx_RepairOperation {
	if op == nil {
		op = &xxx_RepairOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
	return op
}

func (o *RepairResponse) xxx_FromOp(ctx context.Context, op *xxx_RepairOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *RepairResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RepairResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RepairOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
