package ivdsvolumemf

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
	// IVdsVolumeMF interface identifier ee2d5ded-6236-4169-931d-b9778ce03dc6
	VolumeMFIID = &dcom.IID{Data1: 0xee2d5ded, Data2: 0x6236, Data3: 0x4169, Data4: []byte{0x93, 0x1d, 0xb9, 0x77, 0x8c, 0xe0, 0x3d, 0xc6}}
	// Syntax UUID
	VolumeMFSyntaxUUID = &uuid.UUID{TimeLow: 0xee2d5ded, TimeMid: 0x6236, TimeHiAndVersion: 0x4169, ClockSeqHiAndReserved: 0x93, ClockSeqLow: 0x1d, Node: [6]uint8{0xb9, 0x77, 0x8c, 0xe0, 0x3d, 0xc6}}
	// Syntax ID
	VolumeMFSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: VolumeMFSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsVolumeMF interface.
type VolumeMFClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The GetFileSystemProperties method returns property details about the file system
	// on the current volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetFileSystemProperties(context.Context, *GetFileSystemPropertiesRequest, ...dcerpc.CallOption) (*GetFileSystemPropertiesResponse, error)

	// The Format method formats a file system on the current volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Format(context.Context, *FormatRequest, ...dcerpc.CallOption) (*FormatResponse, error)

	// The AddAccessPath method adds an access path to the current volume.<129>
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	AddAccessPath(context.Context, *AddAccessPathRequest, ...dcerpc.CallOption) (*AddAccessPathResponse, error)

	// The QueryAccessPaths method returns a list of access paths and a drive letter as
	// a single case-insensitive Unicode character, if one exists, for the current volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryAccessPaths(context.Context, *QueryAccessPathsRequest, ...dcerpc.CallOption) (*QueryAccessPathsResponse, error)

	// The QueryReparsePoints method returns all reparse points for the current volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryReparsePoints(context.Context, *QueryReparsePointsRequest, ...dcerpc.CallOption) (*QueryReparsePointsResponse, error)

	// The DeleteAccessPath method removes the access path from the current volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	DeleteAccessPath(context.Context, *DeleteAccessPathRequest, ...dcerpc.CallOption) (*DeleteAccessPathResponse, error)

	// The Mount method mounts a volume.
	//
	// This method has no parameters.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Mount(context.Context, *MountRequest, ...dcerpc.CallOption) (*MountResponse, error)

	// The Dismount method dismounts a mounted volume.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Dismount(context.Context, *DismountRequest, ...dcerpc.CallOption) (*DismountResponse, error)

	// The SetFileSystemFlags method sets the file system flags.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	SetFileSystemFlags(context.Context, *SetFileSystemFlagsRequest, ...dcerpc.CallOption) (*SetFileSystemFlagsResponse, error)

	// The ClearFileSystemFlags method clears the file system flags.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	ClearFileSystemFlags(context.Context, *ClearFileSystemFlagsRequest, ...dcerpc.CallOption) (*ClearFileSystemFlagsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) VolumeMFClient
}

type xxx_DefaultVolumeMFClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultVolumeMFClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultVolumeMFClient) GetFileSystemProperties(ctx context.Context, in *GetFileSystemPropertiesRequest, opts ...dcerpc.CallOption) (*GetFileSystemPropertiesResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &GetFileSystemPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeMFClient) Format(ctx context.Context, in *FormatRequest, opts ...dcerpc.CallOption) (*FormatResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &FormatResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeMFClient) AddAccessPath(ctx context.Context, in *AddAccessPathRequest, opts ...dcerpc.CallOption) (*AddAccessPathResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &AddAccessPathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeMFClient) QueryAccessPaths(ctx context.Context, in *QueryAccessPathsRequest, opts ...dcerpc.CallOption) (*QueryAccessPathsResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &QueryAccessPathsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeMFClient) QueryReparsePoints(ctx context.Context, in *QueryReparsePointsRequest, opts ...dcerpc.CallOption) (*QueryReparsePointsResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &QueryReparsePointsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeMFClient) DeleteAccessPath(ctx context.Context, in *DeleteAccessPathRequest, opts ...dcerpc.CallOption) (*DeleteAccessPathResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &DeleteAccessPathResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeMFClient) Mount(ctx context.Context, in *MountRequest, opts ...dcerpc.CallOption) (*MountResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &MountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeMFClient) Dismount(ctx context.Context, in *DismountRequest, opts ...dcerpc.CallOption) (*DismountResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &DismountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeMFClient) SetFileSystemFlags(ctx context.Context, in *SetFileSystemFlagsRequest, opts ...dcerpc.CallOption) (*SetFileSystemFlagsResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &SetFileSystemFlagsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeMFClient) ClearFileSystemFlags(ctx context.Context, in *ClearFileSystemFlagsRequest, opts ...dcerpc.CallOption) (*ClearFileSystemFlagsResponse, error) {
	op := in.xxx_ToOp(ctx)
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
	out := &ClearFileSystemFlagsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVolumeMFClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultVolumeMFClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultVolumeMFClient) IPID(ctx context.Context, ipid *dcom.IPID) VolumeMFClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultVolumeMFClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewVolumeMFClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (VolumeMFClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(VolumeMFSyntaxV0_0))...)
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
	return &xxx_DefaultVolumeMFClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetFileSystemPropertiesOperation structure represents the GetFileSystemProperties operation
type xxx_GetFileSystemPropertiesOperation struct {
	This               *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat          `idl:"name:That" json:"that"`
	FileSystemProperty *vds.FileSystemProperty `idl:"name:pFileSystemProp" json:"file_system_property"`
	Return             int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetFileSystemPropertiesOperation) OpNum() int { return 3 }

func (o *xxx_GetFileSystemPropertiesOperation) OpName() string {
	return "/IVdsVolumeMF/v0/GetFileSystemProperties"
}

func (o *xxx_GetFileSystemPropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileSystemPropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetFileSystemPropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetFileSystemPropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetFileSystemPropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pFileSystemProp {out} (1:{pointer=ref}*(1))(2:{alias=VDS_FILE_SYSTEM_PROP}(struct))
	{
		if o.FileSystemProperty != nil {
			if err := o.FileSystemProperty.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.FileSystemProperty{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetFileSystemPropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pFileSystemProp {out} (1:{pointer=ref}*(1))(2:{alias=VDS_FILE_SYSTEM_PROP}(struct))
	{
		if o.FileSystemProperty == nil {
			o.FileSystemProperty = &vds.FileSystemProperty{}
		}
		if err := o.FileSystemProperty.UnmarshalNDR(ctx, w); err != nil {
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

// GetFileSystemPropertiesRequest structure represents the GetFileSystemProperties operation request
type GetFileSystemPropertiesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetFileSystemPropertiesRequest) xxx_ToOp(ctx context.Context) *xxx_GetFileSystemPropertiesOperation {
	if o == nil {
		return &xxx_GetFileSystemPropertiesOperation{}
	}
	return &xxx_GetFileSystemPropertiesOperation{
		This: o.This,
	}
}

func (o *GetFileSystemPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetFileSystemPropertiesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetFileSystemPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetFileSystemPropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileSystemPropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetFileSystemPropertiesResponse structure represents the GetFileSystemProperties operation response
type GetFileSystemPropertiesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pFileSystemProp: A pointer to a VDS_FILE_SYSTEM_PROP structure that, if the operation
	// is successfully completed, receives the properties of the file system on the volume.
	FileSystemProperty *vds.FileSystemProperty `idl:"name:pFileSystemProp" json:"file_system_property"`
	// Return: The GetFileSystemProperties return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetFileSystemPropertiesResponse) xxx_ToOp(ctx context.Context) *xxx_GetFileSystemPropertiesOperation {
	if o == nil {
		return &xxx_GetFileSystemPropertiesOperation{}
	}
	return &xxx_GetFileSystemPropertiesOperation{
		That:               o.That,
		FileSystemProperty: o.FileSystemProperty,
		Return:             o.Return,
	}
}

func (o *GetFileSystemPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetFileSystemPropertiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileSystemProperty = op.FileSystemProperty
	o.Return = op.Return
}
func (o *GetFileSystemPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetFileSystemPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetFileSystemPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FormatOperation structure represents the Format operation
type xxx_FormatOperation struct {
	This               *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat     `idl:"name:That" json:"that"`
	Type               vds.FileSystemType `idl:"name:type" json:"type"`
	Label              string             `idl:"name:pwszLabel;string" json:"label"`
	UnitAllocationSize uint32             `idl:"name:dwUnitAllocationSize" json:"unit_allocation_size"`
	Force              int32              `idl:"name:bForce" json:"force"`
	QuickFormat        int32              `idl:"name:bQuickFormat" json:"quick_format"`
	EnableCompression  int32              `idl:"name:bEnableCompression" json:"enable_compression"`
	Async              *vds.Async         `idl:"name:ppAsync" json:"async"`
	Return             int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_FormatOperation) OpNum() int { return 4 }

func (o *xxx_FormatOperation) OpName() string { return "/IVdsVolumeMF/v0/Format" }

func (o *xxx_FormatOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// type {in} (1:{alias=VDS_FILE_SYSTEM_TYPE}(enum))
	{
		if err := w.WriteEnum(uint16(o.Type)); err != nil {
			return err
		}
	}
	// pwszLabel {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Label); err != nil {
			return err
		}
	}
	// dwUnitAllocationSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.UnitAllocationSize); err != nil {
			return err
		}
	}
	// bForce {in} (1:(int32))
	{
		if err := w.WriteData(o.Force); err != nil {
			return err
		}
	}
	// bQuickFormat {in} (1:(int32))
	{
		if err := w.WriteData(o.QuickFormat); err != nil {
			return err
		}
	}
	// bEnableCompression {in} (1:(int32))
	{
		if err := w.WriteData(o.EnableCompression); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// type {in} (1:{alias=VDS_FILE_SYSTEM_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Type)); err != nil {
			return err
		}
	}
	// pwszLabel {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Label); err != nil {
			return err
		}
	}
	// dwUnitAllocationSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.UnitAllocationSize); err != nil {
			return err
		}
	}
	// bForce {in} (1:(int32))
	{
		if err := w.ReadData(&o.Force); err != nil {
			return err
		}
	}
	// bQuickFormat {in} (1:(int32))
	{
		if err := w.ReadData(&o.QuickFormat); err != nil {
			return err
		}
	}
	// bEnableCompression {in} (1:(int32))
	{
		if err := w.ReadData(&o.EnableCompression); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_FormatOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// FormatRequest structure represents the Format operation request
type FormatRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// type: A file system type that is enumerated by VDS_FILE_SYSTEM_TYPE. Clients that
	// format by using file systems that are not enumerated by VDS_FILE_SYSTEM_TYPE can
	// use the IVdsVolumeMF2::FormatEx method.
	Type vds.FileSystemType `idl:"name:type" json:"type"`
	// pwszLabel: A null-terminated Unicode label to assign to the new file system. The
	// maximum label size is file system-dependent.
	Label string `idl:"name:pwszLabel;string" json:"label"`
	// dwUnitAllocationSize: The size, in bytes, of the allocation unit for the file system.
	// The value MUST be a power of 2. The allocation unit range is file system-dependent.
	UnitAllocationSize uint32 `idl:"name:dwUnitAllocationSize" json:"unit_allocation_size"`
	// bForce: A Boolean that determines whether the format is forced, even if the volume
	// is in use.
	Force int32 `idl:"name:bForce" json:"force"`
	// bQuickFormat: A Boolean that determines whether a file system is quick format. A
	// quick format does not verify each sector on the volume.
	QuickFormat int32 `idl:"name:bQuickFormat" json:"quick_format"`
	// bEnableCompression: A Boolean that determines whether a file system is created with
	// compression enabled.<127>
	EnableCompression int32 `idl:"name:bEnableCompression" json:"enable_compression"`
}

func (o *FormatRequest) xxx_ToOp(ctx context.Context) *xxx_FormatOperation {
	if o == nil {
		return &xxx_FormatOperation{}
	}
	return &xxx_FormatOperation{
		This:               o.This,
		Type:               o.Type,
		Label:              o.Label,
		UnitAllocationSize: o.UnitAllocationSize,
		Force:              o.Force,
		QuickFormat:        o.QuickFormat,
		EnableCompression:  o.EnableCompression,
	}
}

func (o *FormatRequest) xxx_FromOp(ctx context.Context, op *xxx_FormatOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Type = op.Type
	o.Label = op.Label
	o.UnitAllocationSize = op.UnitAllocationSize
	o.Force = op.Force
	o.QuickFormat = op.QuickFormat
	o.EnableCompression = op.EnableCompression
}
func (o *FormatRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *FormatRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FormatOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FormatResponse structure represents the Format operation response
type FormatResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync interface that, if the operation is successfully
	// completed, receives the IVdsAsync interface to monitor and control this operation.
	// Callers MUST release the interface when they are done with it.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The Format return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FormatResponse) xxx_ToOp(ctx context.Context) *xxx_FormatOperation {
	if o == nil {
		return &xxx_FormatOperation{}
	}
	return &xxx_FormatOperation{
		That:   o.That,
		Async:  o.Async,
		Return: o.Return,
	}
}

func (o *FormatResponse) xxx_FromOp(ctx context.Context, op *xxx_FormatOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *FormatResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *FormatResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FormatOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddAccessPathOperation structure represents the AddAccessPath operation
type xxx_AddAccessPathOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path   string         `idl:"name:pwszPath;max_is:(259);string" json:"path"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AddAccessPathOperation) OpNum() int { return 5 }

func (o *xxx_AddAccessPathOperation) OpName() string { return "/IVdsVolumeMF/v0/AddAccessPath" }

func (o *xxx_AddAccessPathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	// cannot evaluate expression 260
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddAccessPathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,max_is=259,string,null](wchar))
	{
		dimSize1 := uint64(260)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := ndr.UTF16NLen(o.Path)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		_Path_buf := utf16.Encode([]rune(o.Path))
		if uint64(len(_Path_buf)) > sizeInfo[0]-1 {
			_Path_buf = _Path_buf[:sizeInfo[0]-1]
		}
		if o.Path != ndr.ZeroString {
			_Path_buf = append(_Path_buf, uint16(0))
		}
		for i1 := range _Path_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_Path_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_Path_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_AddAccessPathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,max_is=259,string,null](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _Path_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Path_buf", sizeInfo[0])
		}
		_Path_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Path_buf {
			i1 := i1
			if err := w.ReadData(&_Path_buf[i1]); err != nil {
				return err
			}
		}
		o.Path = strings.TrimRight(string(utf16.Decode(_Path_buf)), ndr.ZeroString)
	}
	return nil
}

func (o *xxx_AddAccessPathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddAccessPathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AddAccessPathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AddAccessPathRequest structure represents the AddAccessPath operation request
type AddAccessPathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszPath: A null-terminated Unicode string that indicates the access path. A trailing
	// backslash MUST be included if the access path is a drive letter (for example, "F:\").
	Path string `idl:"name:pwszPath;max_is:(259);string" json:"path"`
}

func (o *AddAccessPathRequest) xxx_ToOp(ctx context.Context) *xxx_AddAccessPathOperation {
	if o == nil {
		return &xxx_AddAccessPathOperation{}
	}
	return &xxx_AddAccessPathOperation{
		This: o.This,
		Path: o.Path,
	}
}

func (o *AddAccessPathRequest) xxx_FromOp(ctx context.Context, op *xxx_AddAccessPathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
}
func (o *AddAccessPathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddAccessPathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddAccessPathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddAccessPathResponse structure represents the AddAccessPath operation response
type AddAccessPathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AddAccessPath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddAccessPathResponse) xxx_ToOp(ctx context.Context) *xxx_AddAccessPathOperation {
	if o == nil {
		return &xxx_AddAccessPathOperation{}
	}
	return &xxx_AddAccessPathOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *AddAccessPathResponse) xxx_FromOp(ctx context.Context, op *xxx_AddAccessPathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AddAccessPathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddAccessPathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddAccessPathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryAccessPathsOperation structure represents the QueryAccessPaths operation
type xxx_QueryAccessPathsOperation struct {
	This                *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat `idl:"name:That" json:"that"`
	PathArray           []string       `idl:"name:pwszPathArray;size_is:(, plNumberOfAccessPaths);string" json:"path_array"`
	NumberOfAccessPaths int32          `idl:"name:plNumberOfAccessPaths" json:"number_of_access_paths"`
	Return              int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryAccessPathsOperation) OpNum() int { return 6 }

func (o *xxx_QueryAccessPathsOperation) OpName() string { return "/IVdsVolumeMF/v0/QueryAccessPaths" }

func (o *xxx_QueryAccessPathsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryAccessPathsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryAccessPathsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_QueryAccessPathsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.PathArray != nil && o.NumberOfAccessPaths == 0 {
		o.NumberOfAccessPaths = int32(len(o.PathArray))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryAccessPathsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pwszPathArray {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=plNumberOfAccessPaths]*(1)[dim:0,string,null](wchar))
	{
		if o.PathArray != nil || o.NumberOfAccessPaths > 0 {
			_ptr_pwszPathArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NumberOfAccessPaths)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.PathArray {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.PathArray[i1] != "" {
						_ptr_pwszPathArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.PathArray[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.PathArray[i1], _ptr_pwszPathArray); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.PathArray); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PathArray, _ptr_pwszPathArray); err != nil {
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
	// plNumberOfAccessPaths {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.NumberOfAccessPaths); err != nil {
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

func (o *xxx_QueryAccessPathsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pwszPathArray {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=plNumberOfAccessPaths]*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszPathArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.PathArray", sizeInfo[0])
			}
			o.PathArray = make([]string, sizeInfo[0])
			for i1 := range o.PathArray {
				i1 := i1
				_ptr_pwszPathArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.PathArray[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_pwszPathArray := func(ptr interface{}) { o.PathArray[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.PathArray[i1], _s_pwszPathArray, _ptr_pwszPathArray); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pwszPathArray := func(ptr interface{}) { o.PathArray = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.PathArray, _s_pwszPathArray, _ptr_pwszPathArray); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// plNumberOfAccessPaths {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.NumberOfAccessPaths); err != nil {
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

// QueryAccessPathsRequest structure represents the QueryAccessPaths operation request
type QueryAccessPathsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *QueryAccessPathsRequest) xxx_ToOp(ctx context.Context) *xxx_QueryAccessPathsOperation {
	if o == nil {
		return &xxx_QueryAccessPathsOperation{}
	}
	return &xxx_QueryAccessPathsOperation{
		This: o.This,
	}
}

func (o *QueryAccessPathsRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryAccessPathsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *QueryAccessPathsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryAccessPathsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryAccessPathsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryAccessPathsResponse structure represents the QueryAccessPaths operation response
type QueryAccessPathsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pwszPathArray: A pointer to an array of strings that, if the operation is successfully
	// completed, receives the array of access paths.
	PathArray []string `idl:"name:pwszPathArray;size_is:(, plNumberOfAccessPaths);string" json:"path_array"`
	// plNumberOfAccessPaths: A pointer to a variable that, if the operation is successfully
	// completed, receives the total number of elements returned in pwszPathArray.
	NumberOfAccessPaths int32 `idl:"name:plNumberOfAccessPaths" json:"number_of_access_paths"`
	// Return: The QueryAccessPaths return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryAccessPathsResponse) xxx_ToOp(ctx context.Context) *xxx_QueryAccessPathsOperation {
	if o == nil {
		return &xxx_QueryAccessPathsOperation{}
	}
	return &xxx_QueryAccessPathsOperation{
		That:                o.That,
		PathArray:           o.PathArray,
		NumberOfAccessPaths: o.NumberOfAccessPaths,
		Return:              o.Return,
	}
}

func (o *QueryAccessPathsResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryAccessPathsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PathArray = op.PathArray
	o.NumberOfAccessPaths = op.NumberOfAccessPaths
	o.Return = op.Return
}
func (o *QueryAccessPathsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryAccessPathsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryAccessPathsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryReparsePointsOperation structure represents the QueryReparsePoints operation
type xxx_QueryReparsePointsOperation struct {
	This                           *dcom.ORPCThis              `idl:"name:This" json:"this"`
	That                           *dcom.ORPCThat              `idl:"name:That" json:"that"`
	ReparsePointProperties         []*vds.ReparsePointProperty `idl:"name:ppReparsePointProps;size_is:(, plNumberOfReparsePointProps)" json:"reparse_point_properties"`
	NumberOfReparsePointProperties int32                       `idl:"name:plNumberOfReparsePointProps" json:"number_of_reparse_point_properties"`
	Return                         int32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryReparsePointsOperation) OpNum() int { return 7 }

func (o *xxx_QueryReparsePointsOperation) OpName() string {
	return "/IVdsVolumeMF/v0/QueryReparsePoints"
}

func (o *xxx_QueryReparsePointsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryReparsePointsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryReparsePointsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_QueryReparsePointsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ReparsePointProperties != nil && o.NumberOfReparsePointProperties == 0 {
		o.NumberOfReparsePointProperties = int32(len(o.ReparsePointProperties))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryReparsePointsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppReparsePointProps {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VDS_REPARSE_POINT_PROP}[dim:0,size_is=plNumberOfReparsePointProps](struct))
	{
		if o.ReparsePointProperties != nil || o.NumberOfReparsePointProperties > 0 {
			_ptr_ppReparsePointProps := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NumberOfReparsePointProperties)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ReparsePointProperties {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.ReparsePointProperties[i1] != nil {
						if err := o.ReparsePointProperties[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&vds.ReparsePointProperty{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.ReparsePointProperties); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&vds.ReparsePointProperty{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReparsePointProperties, _ptr_ppReparsePointProps); err != nil {
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
	// plNumberOfReparsePointProps {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.NumberOfReparsePointProperties); err != nil {
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

func (o *xxx_QueryReparsePointsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppReparsePointProps {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VDS_REPARSE_POINT_PROP}[dim:0,size_is=plNumberOfReparsePointProps](struct))
	{
		_ptr_ppReparsePointProps := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ReparsePointProperties", sizeInfo[0])
			}
			o.ReparsePointProperties = make([]*vds.ReparsePointProperty, sizeInfo[0])
			for i1 := range o.ReparsePointProperties {
				i1 := i1
				if o.ReparsePointProperties[i1] == nil {
					o.ReparsePointProperties[i1] = &vds.ReparsePointProperty{}
				}
				if err := o.ReparsePointProperties[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppReparsePointProps := func(ptr interface{}) { o.ReparsePointProperties = *ptr.(*[]*vds.ReparsePointProperty) }
		if err := w.ReadPointer(&o.ReparsePointProperties, _s_ppReparsePointProps, _ptr_ppReparsePointProps); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// plNumberOfReparsePointProps {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.NumberOfReparsePointProperties); err != nil {
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

// QueryReparsePointsRequest structure represents the QueryReparsePoints operation request
type QueryReparsePointsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *QueryReparsePointsRequest) xxx_ToOp(ctx context.Context) *xxx_QueryReparsePointsOperation {
	if o == nil {
		return &xxx_QueryReparsePointsOperation{}
	}
	return &xxx_QueryReparsePointsOperation{
		This: o.This,
	}
}

func (o *QueryReparsePointsRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryReparsePointsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *QueryReparsePointsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryReparsePointsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryReparsePointsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryReparsePointsResponse structure represents the QueryReparsePoints operation response
type QueryReparsePointsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppReparsePointProps: A pointer to an array of VDS_REPARSE_POINT_PROP structures that,
	// if the operation is successfully completed, receives the array of reparse point properties.
	ReparsePointProperties []*vds.ReparsePointProperty `idl:"name:ppReparsePointProps;size_is:(, plNumberOfReparsePointProps)" json:"reparse_point_properties"`
	// plNumberOfReparsePointProps: A pointer to a variable that, if the operation is successfully
	// completed, receives the total number of elements returned in ppReparsePointPorps.
	NumberOfReparsePointProperties int32 `idl:"name:plNumberOfReparsePointProps" json:"number_of_reparse_point_properties"`
	// Return: The QueryReparsePoints return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryReparsePointsResponse) xxx_ToOp(ctx context.Context) *xxx_QueryReparsePointsOperation {
	if o == nil {
		return &xxx_QueryReparsePointsOperation{}
	}
	return &xxx_QueryReparsePointsOperation{
		That:                           o.That,
		ReparsePointProperties:         o.ReparsePointProperties,
		NumberOfReparsePointProperties: o.NumberOfReparsePointProperties,
		Return:                         o.Return,
	}
}

func (o *QueryReparsePointsResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryReparsePointsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReparsePointProperties = op.ReparsePointProperties
	o.NumberOfReparsePointProperties = op.NumberOfReparsePointProperties
	o.Return = op.Return
}
func (o *QueryReparsePointsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryReparsePointsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryReparsePointsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteAccessPathOperation structure represents the DeleteAccessPath operation
type xxx_DeleteAccessPathOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Path   string         `idl:"name:pwszPath;max_is:(259);string" json:"path"`
	Force  int32          `idl:"name:bForce" json:"force"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteAccessPathOperation) OpNum() int { return 8 }

func (o *xxx_DeleteAccessPathOperation) OpName() string { return "/IVdsVolumeMF/v0/DeleteAccessPath" }

func (o *xxx_DeleteAccessPathOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	// cannot evaluate expression 260
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteAccessPathOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,max_is=259,string,null](wchar))
	{
		dimSize1 := uint64(260)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := ndr.UTF16NLen(o.Path)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		_Path_buf := utf16.Encode([]rune(o.Path))
		if uint64(len(_Path_buf)) > sizeInfo[0]-1 {
			_Path_buf = _Path_buf[:sizeInfo[0]-1]
		}
		if o.Path != ndr.ZeroString {
			_Path_buf = append(_Path_buf, uint16(0))
		}
		for i1 := range _Path_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_Path_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_Path_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint16(0)); err != nil {
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
	return nil
}

func (o *xxx_DeleteAccessPathOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,max_is=259,string,null](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _Path_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Path_buf", sizeInfo[0])
		}
		_Path_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Path_buf {
			i1 := i1
			if err := w.ReadData(&_Path_buf[i1]); err != nil {
				return err
			}
		}
		o.Path = strings.TrimRight(string(utf16.Decode(_Path_buf)), ndr.ZeroString)
	}
	// bForce {in} (1:(int32))
	{
		if err := w.ReadData(&o.Force); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteAccessPathOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteAccessPathOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteAccessPathOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteAccessPathRequest structure represents the DeleteAccessPath operation request
type DeleteAccessPathRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pwszPath: A Unicode string indicating the access path, for example, "C:\myfolder\mydocuments".
	Path string `idl:"name:pwszPath;max_is:(259);string" json:"path"`
	// bForce: A Boolean that determines whether an access is deleted unconditionally, even
	// if the volume is in use. This parameter is meaningful only when the access path is
	// a drive letter.
	Force int32 `idl:"name:bForce" json:"force"`
}

func (o *DeleteAccessPathRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteAccessPathOperation {
	if o == nil {
		return &xxx_DeleteAccessPathOperation{}
	}
	return &xxx_DeleteAccessPathOperation{
		This:  o.This,
		Path:  o.Path,
		Force: o.Force,
	}
}

func (o *DeleteAccessPathRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteAccessPathOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Path = op.Path
	o.Force = op.Force
}
func (o *DeleteAccessPathRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteAccessPathRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteAccessPathOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteAccessPathResponse structure represents the DeleteAccessPath operation response
type DeleteAccessPathResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteAccessPath return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteAccessPathResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteAccessPathOperation {
	if o == nil {
		return &xxx_DeleteAccessPathOperation{}
	}
	return &xxx_DeleteAccessPathOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *DeleteAccessPathResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteAccessPathOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteAccessPathResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteAccessPathResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteAccessPathOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MountOperation structure represents the Mount operation
type xxx_MountOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_MountOperation) OpNum() int { return 9 }

func (o *xxx_MountOperation) OpName() string { return "/IVdsVolumeMF/v0/Mount" }

func (o *xxx_MountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_MountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_MountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_MountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// MountRequest structure represents the Mount operation request
type MountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *MountRequest) xxx_ToOp(ctx context.Context) *xxx_MountOperation {
	if o == nil {
		return &xxx_MountOperation{}
	}
	return &xxx_MountOperation{
		This: o.This,
	}
}

func (o *MountRequest) xxx_FromOp(ctx context.Context, op *xxx_MountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *MountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *MountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MountResponse structure represents the Mount operation response
type MountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Mount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *MountResponse) xxx_ToOp(ctx context.Context) *xxx_MountOperation {
	if o == nil {
		return &xxx_MountOperation{}
	}
	return &xxx_MountOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *MountResponse) xxx_FromOp(ctx context.Context, op *xxx_MountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *MountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *MountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DismountOperation structure represents the Dismount operation
type xxx_DismountOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Force     int32          `idl:"name:bForce" json:"force"`
	Permanent int32          `idl:"name:bPermanent" json:"permanent"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DismountOperation) OpNum() int { return 10 }

func (o *xxx_DismountOperation) OpName() string { return "/IVdsVolumeMF/v0/Dismount" }

func (o *xxx_DismountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DismountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bPermanent {in} (1:(int32))
	{
		if err := w.WriteData(o.Permanent); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DismountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bPermanent {in} (1:(int32))
	{
		if err := w.ReadData(&o.Permanent); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DismountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DismountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DismountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DismountRequest structure represents the Dismount operation request
type DismountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bForce: A Boolean that determines whether the current volume is dismounted unconditionally,
	// even if the volume is in use.
	Force int32 `idl:"name:bForce" json:"force"`
	// bPermanent: A Boolean that determines whether a volume MUST be dismounted permanently
	// by taking the volume offline after dismounting it.
	Permanent int32 `idl:"name:bPermanent" json:"permanent"`
}

func (o *DismountRequest) xxx_ToOp(ctx context.Context) *xxx_DismountOperation {
	if o == nil {
		return &xxx_DismountOperation{}
	}
	return &xxx_DismountOperation{
		This:      o.This,
		Force:     o.Force,
		Permanent: o.Permanent,
	}
}

func (o *DismountRequest) xxx_FromOp(ctx context.Context, op *xxx_DismountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Force = op.Force
	o.Permanent = op.Permanent
}
func (o *DismountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DismountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DismountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DismountResponse structure represents the Dismount operation response
type DismountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Dismount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DismountResponse) xxx_ToOp(ctx context.Context) *xxx_DismountOperation {
	if o == nil {
		return &xxx_DismountOperation{}
	}
	return &xxx_DismountOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *DismountResponse) xxx_FromOp(ctx context.Context, op *xxx_DismountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DismountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DismountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DismountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetFileSystemFlagsOperation structure represents the SetFileSystemFlags operation
type xxx_SetFileSystemFlagsOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Flags  uint32         `idl:"name:ulFlags" json:"flags"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetFileSystemFlagsOperation) OpNum() int { return 11 }

func (o *xxx_SetFileSystemFlagsOperation) OpName() string {
	return "/IVdsVolumeMF/v0/SetFileSystemFlags"
}

func (o *xxx_SetFileSystemFlagsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileSystemFlagsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetFileSystemFlagsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_SetFileSystemFlagsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFileSystemFlagsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetFileSystemFlagsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// SetFileSystemFlagsRequest structure represents the SetFileSystemFlags operation request
type SetFileSystemFlagsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ulFlags: Callers MUST set the VDS_FPF_COMPRESSED flag.
	Flags uint32 `idl:"name:ulFlags" json:"flags"`
}

func (o *SetFileSystemFlagsRequest) xxx_ToOp(ctx context.Context) *xxx_SetFileSystemFlagsOperation {
	if o == nil {
		return &xxx_SetFileSystemFlagsOperation{}
	}
	return &xxx_SetFileSystemFlagsOperation{
		This:  o.This,
		Flags: o.Flags,
	}
}

func (o *SetFileSystemFlagsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetFileSystemFlagsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flags = op.Flags
}
func (o *SetFileSystemFlagsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetFileSystemFlagsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFileSystemFlagsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetFileSystemFlagsResponse structure represents the SetFileSystemFlags operation response
type SetFileSystemFlagsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The SetFileSystemFlags return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetFileSystemFlagsResponse) xxx_ToOp(ctx context.Context) *xxx_SetFileSystemFlagsOperation {
	if o == nil {
		return &xxx_SetFileSystemFlagsOperation{}
	}
	return &xxx_SetFileSystemFlagsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *SetFileSystemFlagsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetFileSystemFlagsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetFileSystemFlagsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetFileSystemFlagsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFileSystemFlagsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ClearFileSystemFlagsOperation structure represents the ClearFileSystemFlags operation
type xxx_ClearFileSystemFlagsOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Flags  uint32         `idl:"name:ulFlags" json:"flags"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ClearFileSystemFlagsOperation) OpNum() int { return 12 }

func (o *xxx_ClearFileSystemFlagsOperation) OpName() string {
	return "/IVdsVolumeMF/v0/ClearFileSystemFlags"
}

func (o *xxx_ClearFileSystemFlagsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearFileSystemFlagsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ClearFileSystemFlagsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_ClearFileSystemFlagsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearFileSystemFlagsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ClearFileSystemFlagsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ClearFileSystemFlagsRequest structure represents the ClearFileSystemFlags operation request
type ClearFileSystemFlagsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ulFlags: Callers MUST clear the VDS_FPF_COMPRESSED flag.
	Flags uint32 `idl:"name:ulFlags" json:"flags"`
}

func (o *ClearFileSystemFlagsRequest) xxx_ToOp(ctx context.Context) *xxx_ClearFileSystemFlagsOperation {
	if o == nil {
		return &xxx_ClearFileSystemFlagsOperation{}
	}
	return &xxx_ClearFileSystemFlagsOperation{
		This:  o.This,
		Flags: o.Flags,
	}
}

func (o *ClearFileSystemFlagsRequest) xxx_FromOp(ctx context.Context, op *xxx_ClearFileSystemFlagsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flags = op.Flags
}
func (o *ClearFileSystemFlagsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ClearFileSystemFlagsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearFileSystemFlagsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ClearFileSystemFlagsResponse structure represents the ClearFileSystemFlags operation response
type ClearFileSystemFlagsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ClearFileSystemFlags return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ClearFileSystemFlagsResponse) xxx_ToOp(ctx context.Context) *xxx_ClearFileSystemFlagsOperation {
	if o == nil {
		return &xxx_ClearFileSystemFlagsOperation{}
	}
	return &xxx_ClearFileSystemFlagsOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ClearFileSystemFlagsResponse) xxx_FromOp(ctx context.Context, op *xxx_ClearFileSystemFlagsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ClearFileSystemFlagsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ClearFileSystemFlagsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearFileSystemFlagsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
