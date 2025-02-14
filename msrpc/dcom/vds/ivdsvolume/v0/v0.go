package ivdsvolume

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
	// IVdsVolume interface identifier 88306bb2-e71f-478c-86a2-79da200a0f11
	VolumeIID = &dcom.IID{Data1: 0x88306bb2, Data2: 0xe71f, Data3: 0x478c, Data4: []byte{0x86, 0xa2, 0x79, 0xda, 0x20, 0x0a, 0x0f, 0x11}}
	// Syntax UUID
	VolumeSyntaxUUID = &uuid.UUID{TimeLow: 0x88306bb2, TimeMid: 0xe71f, TimeHiAndVersion: 0x478c, ClockSeqHiAndReserved: 0x86, ClockSeqLow: 0xa2, Node: [6]uint8{0x79, 0xda, 0x20, 0xa, 0xf, 0x11}}
	// Syntax ID
	VolumeSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: VolumeSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsVolume interface.
type VolumeClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// GetProperties operation.
	GetProperties(context.Context, *GetPropertiesRequest, ...dcerpc.CallOption) (*GetPropertiesResponse, error)

	// GetPack operation.
	GetPack(context.Context, *GetPackRequest, ...dcerpc.CallOption) (*GetPackResponse, error)

	// The QueryPlexes method enumerates the plexes of a volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryPlexes(context.Context, *QueryPlexesRequest, ...dcerpc.CallOption) (*QueryPlexesResponse, error)

	// The Extend method expands the size of the current volume by adding disk extents to
	// each member of each plex.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Extend(context.Context, *ExtendRequest, ...dcerpc.CallOption) (*ExtendResponse, error)

	// Shrink operation.
	Shrink(context.Context, *ShrinkRequest, ...dcerpc.CallOption) (*ShrinkResponse, error)

	// The AddPlex method adds a volume as a plex to the current volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	AddPlex(context.Context, *AddPlexRequest, ...dcerpc.CallOption) (*AddPlexResponse, error)

	// The BreakPlex method removes a specified plex from the current volume. The interface
	// pointer for the new volume object can be retrieved by calling IVdsAsync::Wait through
	// the ppAsync parameter. The VDS_ASYNC_OUTPUT structure that is returned contains the
	// volume object interface pointer in the bvp.pVolumeUnk member.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	BreakPlex(context.Context, *BreakPlexRequest, ...dcerpc.CallOption) (*BreakPlexResponse, error)

	// The RemovePlex method removes a specified plex from a volume. The last plex of a
	// volume cannot be removed.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	RemovePlex(context.Context, *RemovePlexRequest, ...dcerpc.CallOption) (*RemovePlexResponse, error)

	// The Delete method deletes all plexes in a volume.<118>
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// ERROR_SUCCESS (0x00000000)
	Delete(context.Context, *DeleteRequest, ...dcerpc.CallOption) (*DeleteResponse, error)

	// SetFlags operation.
	SetFlags(context.Context, *SetFlagsRequest, ...dcerpc.CallOption) (*SetFlagsResponse, error)

	// ClearFlags operation.
	ClearFlags(context.Context, *ClearFlagsRequest, ...dcerpc.CallOption) (*ClearFlagsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) VolumeClient
}

type xxx_DefaultVolumeClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultVolumeClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultVolumeClient) GetProperties(ctx context.Context, in *GetPropertiesRequest, opts ...dcerpc.CallOption) (*GetPropertiesResponse, error) {
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

func (o *xxx_DefaultVolumeClient) GetPack(ctx context.Context, in *GetPackRequest, opts ...dcerpc.CallOption) (*GetPackResponse, error) {
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
	out := &GetPackResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClient) QueryPlexes(ctx context.Context, in *QueryPlexesRequest, opts ...dcerpc.CallOption) (*QueryPlexesResponse, error) {
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
	out := &QueryPlexesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClient) Extend(ctx context.Context, in *ExtendRequest, opts ...dcerpc.CallOption) (*ExtendResponse, error) {
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
	out := &ExtendResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClient) Shrink(ctx context.Context, in *ShrinkRequest, opts ...dcerpc.CallOption) (*ShrinkResponse, error) {
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
	out := &ShrinkResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClient) AddPlex(ctx context.Context, in *AddPlexRequest, opts ...dcerpc.CallOption) (*AddPlexResponse, error) {
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
	out := &AddPlexResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClient) BreakPlex(ctx context.Context, in *BreakPlexRequest, opts ...dcerpc.CallOption) (*BreakPlexResponse, error) {
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
	out := &BreakPlexResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClient) RemovePlex(ctx context.Context, in *RemovePlexRequest, opts ...dcerpc.CallOption) (*RemovePlexResponse, error) {
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
	out := &RemovePlexResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClient) Delete(ctx context.Context, in *DeleteRequest, opts ...dcerpc.CallOption) (*DeleteResponse, error) {
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
	out := &DeleteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClient) SetFlags(ctx context.Context, in *SetFlagsRequest, opts ...dcerpc.CallOption) (*SetFlagsResponse, error) {
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
	out := &SetFlagsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClient) ClearFlags(ctx context.Context, in *ClearFlagsRequest, opts ...dcerpc.CallOption) (*ClearFlagsResponse, error) {
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
	out := &ClearFlagsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultVolumeClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultVolumeClient) IPID(ctx context.Context, ipid *dcom.IPID) VolumeClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultVolumeClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewVolumeClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (VolumeClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(VolumeSyntaxV0_0))...)
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
	return &xxx_DefaultVolumeClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetPropertiesOperation structure represents the GetProperties operation
type xxx_GetPropertiesOperation struct {
	This             *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That             *dcom.ORPCThat      `idl:"name:That" json:"that"`
	VolumeProperties *vds.VolumeProperty `idl:"name:pVolumeProperties" json:"volume_properties"`
	Return           int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPropertiesOperation) OpNum() int { return 3 }

func (o *xxx_GetPropertiesOperation) OpName() string { return "/IVdsVolume/v0/GetProperties" }

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
	// pVolumeProperties {out} (1:{pointer=ref}*(1))(2:{alias=VDS_VOLUME_PROP}(struct))
	{
		if o.VolumeProperties != nil {
			if err := o.VolumeProperties.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.VolumeProperty{}).MarshalNDR(ctx, w); err != nil {
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
	// pVolumeProperties {out} (1:{pointer=ref}*(1))(2:{alias=VDS_VOLUME_PROP}(struct))
	{
		if o.VolumeProperties == nil {
			o.VolumeProperties = &vds.VolumeProperty{}
		}
		if err := o.VolumeProperties.UnmarshalNDR(ctx, w); err != nil {
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
	That             *dcom.ORPCThat      `idl:"name:That" json:"that"`
	VolumeProperties *vds.VolumeProperty `idl:"name:pVolumeProperties" json:"volume_properties"`
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
	o.VolumeProperties = op.VolumeProperties
	o.Return = op.Return
	return op
}

func (o *GetPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPropertiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.VolumeProperties = op.VolumeProperties
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

// xxx_GetPackOperation structure represents the GetPack operation
type xxx_GetPackOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Pack   *vds.Pack      `idl:"name:ppPack" json:"pack"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPackOperation) OpNum() int { return 4 }

func (o *xxx_GetPackOperation) OpName() string { return "/IVdsVolume/v0/GetPack" }

func (o *xxx_GetPackOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPackOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPackOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPackOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPackOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppPack {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsPack}(interface))
	{
		if o.Pack != nil {
			_ptr_ppPack := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Pack != nil {
					if err := o.Pack.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.Pack{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Pack, _ptr_ppPack); err != nil {
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

func (o *xxx_GetPackOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppPack {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsPack}(interface))
	{
		_ptr_ppPack := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Pack == nil {
				o.Pack = &vds.Pack{}
			}
			if err := o.Pack.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppPack := func(ptr interface{}) { o.Pack = *ptr.(**vds.Pack) }
		if err := w.ReadPointer(&o.Pack, _s_ppPack, _ptr_ppPack); err != nil {
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

// GetPackRequest structure represents the GetPack operation request
type GetPackRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetPackRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPackOperation) *xxx_GetPackOperation {
	if op == nil {
		op = &xxx_GetPackOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *GetPackRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPackOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetPackRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPackRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPackOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPackResponse structure represents the GetPack operation response
type GetPackResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Pack *vds.Pack      `idl:"name:ppPack" json:"pack"`
	// Return: The GetPack return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPackResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPackOperation) *xxx_GetPackOperation {
	if op == nil {
		op = &xxx_GetPackOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Pack = op.Pack
	o.Return = op.Return
	return op
}

func (o *GetPackResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPackOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Pack = op.Pack
	o.Return = op.Return
}
func (o *GetPackResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPackResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPackOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryPlexesOperation structure represents the QueryPlexes operation
type xxx_QueryPlexesOperation struct {
	This   *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Enum   *vds.EnumObject `idl:"name:ppEnum" json:"enum"`
	Return int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryPlexesOperation) OpNum() int { return 5 }

func (o *xxx_QueryPlexesOperation) OpName() string { return "/IVdsVolume/v0/QueryPlexes" }

func (o *xxx_QueryPlexesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryPlexesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryPlexesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_QueryPlexesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryPlexesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryPlexesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// QueryPlexesRequest structure represents the QueryPlexes operation request
type QueryPlexesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *QueryPlexesRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryPlexesOperation) *xxx_QueryPlexesOperation {
	if op == nil {
		op = &xxx_QueryPlexesOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	return op
}

func (o *QueryPlexesRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryPlexesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *QueryPlexesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryPlexesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryPlexesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryPlexesResponse structure represents the QueryPlexes operation response
type QueryPlexesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppEnum: A pointer to an IEnumVdsObject interface that, if the operation is successfully
	// completed, receives the IEnumVdsObject interface of the object that contains an enumeration
	// of volume plex objects in the volume. Callers MUST release the interface when they
	// are done with it.
	Enum *vds.EnumObject `idl:"name:ppEnum" json:"enum"`
	// Return: The QueryPlexes return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryPlexesResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryPlexesOperation) *xxx_QueryPlexesOperation {
	if op == nil {
		op = &xxx_QueryPlexesOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
	return op
}

func (o *QueryPlexesResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryPlexesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *QueryPlexesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryPlexesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryPlexesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ExtendOperation structure represents the Extend operation
type xxx_ExtendOperation struct {
	This           *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat   `idl:"name:That" json:"that"`
	InputDiskArray []*vds.InputDisk `idl:"name:pInputDiskArray;size_is:(lNumberOfDisks);pointer:unique" json:"input_disk_array"`
	NumberOfDisks  int32            `idl:"name:lNumberOfDisks" json:"number_of_disks"`
	Async          *vds.Async       `idl:"name:ppAsync" json:"async"`
	Return         int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_ExtendOperation) OpNum() int { return 6 }

func (o *xxx_ExtendOperation) OpName() string { return "/IVdsVolume/v0/Extend" }

func (o *xxx_ExtendOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
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

func (o *xxx_ExtendOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pInputDiskArray {in} (1:{pointer=unique}*(1))(2:{alias=VDS_INPUT_DISK}[dim:0,size_is=lNumberOfDisks](struct))
	{
		if o.InputDiskArray != nil || o.NumberOfDisks > 0 {
			_ptr_pInputDiskArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
				return nil
			})
			if err := w.WritePointer(&o.InputDiskArray, _ptr_pInputDiskArray); err != nil {
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
	// lNumberOfDisks {in} (1:(int32))
	{
		if err := w.WriteData(o.NumberOfDisks); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExtendOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pInputDiskArray {in} (1:{pointer=unique}*(1))(2:{alias=VDS_INPUT_DISK}[dim:0,size_is=lNumberOfDisks](struct))
	{
		_ptr_pInputDiskArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
			return nil
		})
		_s_pInputDiskArray := func(ptr interface{}) { o.InputDiskArray = *ptr.(*[]*vds.InputDisk) }
		if err := w.ReadPointer(&o.InputDiskArray, _s_pInputDiskArray, _ptr_pInputDiskArray); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
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

func (o *xxx_ExtendOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExtendOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ExtendOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ExtendRequest structure represents the Extend operation request
type ExtendRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pInputDiskArray: A pointer to an array of VDS_INPUT_DISK structures that describe
	// the disk extents to add to the volume--one structure for each disk. Callers SHOULD
	// specify the member index for all the disk extents together with the Extend method,
	// unless the volume has only one plex with only one member.
	InputDiskArray []*vds.InputDisk `idl:"name:pInputDiskArray;size_is:(lNumberOfDisks);pointer:unique" json:"input_disk_array"`
	// lNumberOfDisks: The number of elements in pInputDiskArrray.<111>
	NumberOfDisks int32 `idl:"name:lNumberOfDisks" json:"number_of_disks"`
}

func (o *ExtendRequest) xxx_ToOp(ctx context.Context, op *xxx_ExtendOperation) *xxx_ExtendOperation {
	if op == nil {
		op = &xxx_ExtendOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.InputDiskArray = op.InputDiskArray
	o.NumberOfDisks = op.NumberOfDisks
	return op
}

func (o *ExtendRequest) xxx_FromOp(ctx context.Context, op *xxx_ExtendOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.InputDiskArray = op.InputDiskArray
	o.NumberOfDisks = op.NumberOfDisks
}
func (o *ExtendRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ExtendRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExtendOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExtendResponse structure represents the Extend operation response
type ExtendResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync interface that, if the operation is successfully
	// completed, receives the IVdsAsync interface to monitor and control this operation.
	// Callers MUST release the interface when they are done with it.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The Extend return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExtendResponse) xxx_ToOp(ctx context.Context, op *xxx_ExtendOperation) *xxx_ExtendOperation {
	if op == nil {
		op = &xxx_ExtendOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
	return op
}

func (o *ExtendResponse) xxx_FromOp(ctx context.Context, op *xxx_ExtendOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *ExtendResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ExtendResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExtendOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ShrinkOperation structure represents the Shrink operation
type xxx_ShrinkOperation struct {
	This                  *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                  *dcom.ORPCThat `idl:"name:That" json:"that"`
	NumberOfBytesToRemove uint64         `idl:"name:ullNumberOfBytesToRemove" json:"number_of_bytes_to_remove"`
	Async                 *vds.Async     `idl:"name:ppAsync" json:"async"`
	Return                int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ShrinkOperation) OpNum() int { return 7 }

func (o *xxx_ShrinkOperation) OpName() string { return "/IVdsVolume/v0/Shrink" }

func (o *xxx_ShrinkOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ShrinkOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ullNumberOfBytesToRemove {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.NumberOfBytesToRemove); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ShrinkOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ullNumberOfBytesToRemove {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.NumberOfBytesToRemove); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ShrinkOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ShrinkOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ShrinkOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ShrinkRequest structure represents the Shrink operation request
type ShrinkRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                  *dcom.ORPCThis `idl:"name:This" json:"this"`
	NumberOfBytesToRemove uint64         `idl:"name:ullNumberOfBytesToRemove" json:"number_of_bytes_to_remove"`
}

func (o *ShrinkRequest) xxx_ToOp(ctx context.Context, op *xxx_ShrinkOperation) *xxx_ShrinkOperation {
	if op == nil {
		op = &xxx_ShrinkOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.NumberOfBytesToRemove = op.NumberOfBytesToRemove
	return op
}

func (o *ShrinkRequest) xxx_FromOp(ctx context.Context, op *xxx_ShrinkOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.NumberOfBytesToRemove = op.NumberOfBytesToRemove
}
func (o *ShrinkRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ShrinkRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ShrinkOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ShrinkResponse structure represents the Shrink operation response
type ShrinkResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	Async *vds.Async     `idl:"name:ppAsync" json:"async"`
	// Return: The Shrink return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ShrinkResponse) xxx_ToOp(ctx context.Context, op *xxx_ShrinkOperation) *xxx_ShrinkOperation {
	if op == nil {
		op = &xxx_ShrinkOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
	return op
}

func (o *ShrinkResponse) xxx_FromOp(ctx context.Context, op *xxx_ShrinkOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *ShrinkResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ShrinkResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ShrinkOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddPlexOperation structure represents the AddPlex operation
type xxx_AddPlexOperation struct {
	This     *dcom.ORPCThis `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat `idl:"name:That" json:"that"`
	VolumeID *vds.ObjectID  `idl:"name:VolumeId" json:"volume_id"`
	Async    *vds.Async     `idl:"name:ppAsync" json:"async"`
	Return   int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AddPlexOperation) OpNum() int { return 8 }

func (o *xxx_AddPlexOperation) OpName() string { return "/IVdsVolume/v0/AddPlex" }

func (o *xxx_AddPlexOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddPlexOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// VolumeId {in} (1:{alias=VDS_OBJECT_ID, names=GUID}(struct))
	{
		if o.VolumeID != nil {
			if err := o.VolumeID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_AddPlexOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// VolumeId {in} (1:{alias=VDS_OBJECT_ID, names=GUID}(struct))
	{
		if o.VolumeID == nil {
			o.VolumeID = &vds.ObjectID{}
		}
		if err := o.VolumeID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddPlexOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddPlexOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AddPlexOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AddPlexRequest structure represents the AddPlex operation request
type AddPlexRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// VolumeId: The VDS object ID of the volume object to add as a plex.
	VolumeID *vds.ObjectID `idl:"name:VolumeId" json:"volume_id"`
}

func (o *AddPlexRequest) xxx_ToOp(ctx context.Context, op *xxx_AddPlexOperation) *xxx_AddPlexOperation {
	if op == nil {
		op = &xxx_AddPlexOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
	return op
}

func (o *AddPlexRequest) xxx_FromOp(ctx context.Context, op *xxx_AddPlexOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VolumeID = op.VolumeID
}
func (o *AddPlexRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AddPlexRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddPlexOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddPlexResponse structure represents the AddPlex operation response
type AddPlexResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync interface that, if the operation is successfully
	// completed, receives the IVdsAsync interface to monitor and control this operation.
	// Callers MUST release the interface when they are done with it.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The AddPlex return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddPlexResponse) xxx_ToOp(ctx context.Context, op *xxx_AddPlexOperation) *xxx_AddPlexOperation {
	if op == nil {
		op = &xxx_AddPlexOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
	return op
}

func (o *AddPlexResponse) xxx_FromOp(ctx context.Context, op *xxx_AddPlexOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *AddPlexResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AddPlexResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddPlexOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BreakPlexOperation structure represents the BreakPlex operation
type xxx_BreakPlexOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	PlexID *vds.ObjectID  `idl:"name:plexId" json:"plex_id"`
	Async  *vds.Async     `idl:"name:ppAsync" json:"async"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_BreakPlexOperation) OpNum() int { return 9 }

func (o *xxx_BreakPlexOperation) OpName() string { return "/IVdsVolume/v0/BreakPlex" }

func (o *xxx_BreakPlexOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BreakPlexOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// plexId {in} (1:{alias=VDS_OBJECT_ID, names=GUID}(struct))
	{
		if o.PlexID != nil {
			if err := o.PlexID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_BreakPlexOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// plexId {in} (1:{alias=VDS_OBJECT_ID, names=GUID}(struct))
	{
		if o.PlexID == nil {
			o.PlexID = &vds.ObjectID{}
		}
		if err := o.PlexID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BreakPlexOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BreakPlexOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BreakPlexOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// BreakPlexRequest structure represents the BreakPlex operation request
type BreakPlexRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// plexId: The GUID of the plex to be broken.
	PlexID *vds.ObjectID `idl:"name:plexId" json:"plex_id"`
}

func (o *BreakPlexRequest) xxx_ToOp(ctx context.Context, op *xxx_BreakPlexOperation) *xxx_BreakPlexOperation {
	if op == nil {
		op = &xxx_BreakPlexOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.PlexID = op.PlexID
	return op
}

func (o *BreakPlexRequest) xxx_FromOp(ctx context.Context, op *xxx_BreakPlexOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PlexID = op.PlexID
}
func (o *BreakPlexRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BreakPlexRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BreakPlexOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BreakPlexResponse structure represents the BreakPlex operation response
type BreakPlexResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync interface that, if the operation is successfully
	// completed, receives the IVdsAsync interface to monitor and control this operation.
	// Callers MUST release the interface when they are done with it. If the IVdsAsync::Wait
	// method is called on the interface, the interfaces returned in the VDS_ASYNC_OUTPUT
	// structure MUST be released as well. For information on asynchronous tasks, see section
	// 3.4.5.1.9.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The BreakPlex return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *BreakPlexResponse) xxx_ToOp(ctx context.Context, op *xxx_BreakPlexOperation) *xxx_BreakPlexOperation {
	if op == nil {
		op = &xxx_BreakPlexOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
	return op
}

func (o *BreakPlexResponse) xxx_FromOp(ctx context.Context, op *xxx_BreakPlexOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *BreakPlexResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BreakPlexResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BreakPlexOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemovePlexOperation structure represents the RemovePlex operation
type xxx_RemovePlexOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	PlexID *vds.ObjectID  `idl:"name:plexId" json:"plex_id"`
	Async  *vds.Async     `idl:"name:ppAsync" json:"async"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_RemovePlexOperation) OpNum() int { return 10 }

func (o *xxx_RemovePlexOperation) OpName() string { return "/IVdsVolume/v0/RemovePlex" }

func (o *xxx_RemovePlexOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemovePlexOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// plexId {in} (1:{alias=VDS_OBJECT_ID, names=GUID}(struct))
	{
		if o.PlexID != nil {
			if err := o.PlexID.MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RemovePlexOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// plexId {in} (1:{alias=VDS_OBJECT_ID, names=GUID}(struct))
	{
		if o.PlexID == nil {
			o.PlexID = &vds.ObjectID{}
		}
		if err := o.PlexID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemovePlexOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemovePlexOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RemovePlexOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RemovePlexRequest structure represents the RemovePlex operation request
type RemovePlexRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// plexId: The VDS object ID of the volume plex object to remove.
	PlexID *vds.ObjectID `idl:"name:plexId" json:"plex_id"`
}

func (o *RemovePlexRequest) xxx_ToOp(ctx context.Context, op *xxx_RemovePlexOperation) *xxx_RemovePlexOperation {
	if op == nil {
		op = &xxx_RemovePlexOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.PlexID = op.PlexID
	return op
}

func (o *RemovePlexRequest) xxx_FromOp(ctx context.Context, op *xxx_RemovePlexOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.PlexID = op.PlexID
}
func (o *RemovePlexRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemovePlexRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemovePlexOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemovePlexResponse structure represents the RemovePlex operation response
type RemovePlexResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync interface that, if the operation is successfully
	// completed, receives the IVdsAsync interface to monitor and control this operation.
	// Callers MUST release the interface when they are done with it.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The RemovePlex return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemovePlexResponse) xxx_ToOp(ctx context.Context, op *xxx_RemovePlexOperation) *xxx_RemovePlexOperation {
	if op == nil {
		op = &xxx_RemovePlexOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
	return op
}

func (o *RemovePlexResponse) xxx_FromOp(ctx context.Context, op *xxx_RemovePlexOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *RemovePlexResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemovePlexResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemovePlexOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteOperation structure represents the Delete operation
type xxx_DeleteOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Force  int32          `idl:"name:bForce" json:"force"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteOperation) OpNum() int { return 11 }

func (o *xxx_DeleteOperation) OpName() string { return "/IVdsVolume/v0/Delete" }

func (o *xxx_DeleteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bForce {in} (1:(int32))
	{
		if err := w.WriteData(o.Force); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bForce {in} (1:(int32))
	{
		if err := w.ReadData(&o.Force); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteRequest structure represents the Delete operation request
type DeleteRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bForce: A Boolean that determines whether all plexes in a volume are deleted when
	// the volume is in use.
	Force int32 `idl:"name:bForce" json:"force"`
}

func (o *DeleteRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteOperation) *xxx_DeleteOperation {
	if op == nil {
		op = &xxx_DeleteOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Force = op.Force
	return op
}

func (o *DeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Force = op.Force
}
func (o *DeleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteResponse structure represents the Delete operation response
type DeleteResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Delete return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteOperation) *xxx_DeleteOperation {
	if op == nil {
		op = &xxx_DeleteOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *DeleteResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetFlagsOperation structure represents the SetFlags operation
type xxx_SetFlagsOperation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	Flags         uint32         `idl:"name:ulFlags" json:"flags"`
	RevertOnClose int32          `idl:"name:bRevertOnClose" json:"revert_on_close"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetFlagsOperation) OpNum() int { return 12 }

func (o *xxx_SetFlagsOperation) OpName() string { return "/IVdsVolume/v0/SetFlags" }

func (o *xxx_SetFlagsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlagsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ulFlags {in} (1:(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// bRevertOnClose {in} (1:(int32))
	{
		if err := w.WriteData(o.RevertOnClose); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlagsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ulFlags {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// bRevertOnClose {in} (1:(int32))
	{
		if err := w.ReadData(&o.RevertOnClose); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlagsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlagsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetFlagsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetFlagsRequest structure represents the SetFlags operation request
type SetFlagsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	Flags         uint32         `idl:"name:ulFlags" json:"flags"`
	RevertOnClose int32          `idl:"name:bRevertOnClose" json:"revert_on_close"`
}

func (o *SetFlagsRequest) xxx_ToOp(ctx context.Context, op *xxx_SetFlagsOperation) *xxx_SetFlagsOperation {
	if op == nil {
		op = &xxx_SetFlagsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Flags = op.Flags
	o.RevertOnClose = op.RevertOnClose
	return op
}

func (o *SetFlagsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetFlagsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flags = op.Flags
	o.RevertOnClose = op.RevertOnClose
}
func (o *SetFlagsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetFlagsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFlagsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetFlagsResponse structure represents the SetFlags operation response
type SetFlagsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetFlags return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetFlagsResponse) xxx_ToOp(ctx context.Context, op *xxx_SetFlagsOperation) *xxx_SetFlagsOperation {
	if op == nil {
		op = &xxx_SetFlagsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *SetFlagsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetFlagsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetFlagsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetFlagsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFlagsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ClearFlagsOperation structure represents the ClearFlags operation
type xxx_ClearFlagsOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Flags  uint32         `idl:"name:ulFlags" json:"flags"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ClearFlagsOperation) OpNum() int { return 13 }

func (o *xxx_ClearFlagsOperation) OpName() string { return "/IVdsVolume/v0/ClearFlags" }

func (o *xxx_ClearFlagsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearFlagsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ulFlags {in} (1:(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearFlagsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ulFlags {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearFlagsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearFlagsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ClearFlagsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ClearFlagsRequest structure represents the ClearFlags operation request
type ClearFlagsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This  *dcom.ORPCThis `idl:"name:This" json:"this"`
	Flags uint32         `idl:"name:ulFlags" json:"flags"`
}

func (o *ClearFlagsRequest) xxx_ToOp(ctx context.Context, op *xxx_ClearFlagsOperation) *xxx_ClearFlagsOperation {
	if op == nil {
		op = &xxx_ClearFlagsOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Flags = op.Flags
	return op
}

func (o *ClearFlagsRequest) xxx_FromOp(ctx context.Context, op *xxx_ClearFlagsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flags = op.Flags
}
func (o *ClearFlagsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ClearFlagsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearFlagsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ClearFlagsResponse structure represents the ClearFlags operation response
type ClearFlagsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ClearFlags return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ClearFlagsResponse) xxx_ToOp(ctx context.Context, op *xxx_ClearFlagsOperation) *xxx_ClearFlagsOperation {
	if op == nil {
		op = &xxx_ClearFlagsOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Return = op.Return
	return op
}

func (o *ClearFlagsResponse) xxx_FromOp(ctx context.Context, op *xxx_ClearFlagsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ClearFlagsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ClearFlagsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearFlagsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
