package ivdsvdprovider

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
	// IVdsVdProvider interface identifier b481498c-8354-45f9-84a0-0bdd2832a91f
	VDiskProviderIID = &dcom.IID{Data1: 0xb481498c, Data2: 0x8354, Data3: 0x45f9, Data4: []byte{0x84, 0xa0, 0x0b, 0xdd, 0x28, 0x32, 0xa9, 0x1f}}
	// Syntax UUID
	VDiskProviderSyntaxUUID = &uuid.UUID{TimeLow: 0xb481498c, TimeMid: 0x8354, TimeHiAndVersion: 0x45f9, ClockSeqHiAndReserved: 0x84, ClockSeqLow: 0xa0, Node: [6]uint8{0xb, 0xdd, 0x28, 0x32, 0xa9, 0x1f}}
	// Syntax ID
	VDiskProviderSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: VDiskProviderSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsVdProvider interface.
type VDiskProviderClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The QueryVDisks method returns a list of virtual disks that are managed by the provider.
	//
	// Return Values: The method MUST return zero or a nonerror HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryVDisks(context.Context, *QueryVDisksRequest, ...dcerpc.CallOption) (*QueryVDisksResponse, error)

	// The CreateVDisk method defines a new virtual disk. This method creates a virtual
	// disk file to be used as the backing store for the virtual disk.
	//
	// Return Values: The method MUST return zero to indicate success, or return an implementation-specific
	// nonzero error code to indicate failure. For the HRESULT values predefined by the
	// Virtual Disk Service Remote Protocol, see section 2.2.3.
	CreateVDisk(context.Context, *CreateVDiskRequest, ...dcerpc.CallOption) (*CreateVDiskResponse, error)

	// The AddVDisk method creates a virtual disk object representing the specified virtual
	// disk and adds it to the list of virtual disks managed by the provider. This method
	// returns an IVdsVDisk (section 3.1.15.1) interface pointer to the specified virtual
	// disk object.
	//
	// Return Values: The method MUST return zero or a nonerror HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	AddVDisk(context.Context, *AddVDiskRequest, ...dcerpc.CallOption) (*AddVDiskResponse, error)

	// The GetDiskFromVDisk method returns an IVdsDisk (section 3.1.12.1) interface pointer
	// for a virtual disk given an IVdsVDisk (section 3.1.15.1) interface pointer.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetDiskFromVDisk(context.Context, *GetDiskFromVDiskRequest, ...dcerpc.CallOption) (*GetDiskFromVDiskResponse, error)

	// The GetVDiskFromDisk method returns an IVdsVDisk (section 3.1.15.1) interface pointer
	// for the virtual disk given an IVdsDisk (section 3.1.12.1) interface pointer.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetVDiskFromDisk(context.Context, *GetVDiskFromDiskRequest, ...dcerpc.CallOption) (*GetVDiskFromDiskResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) VDiskProviderClient
}

type xxx_DefaultVDiskProviderClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultVDiskProviderClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultVDiskProviderClient) QueryVDisks(ctx context.Context, in *QueryVDisksRequest, opts ...dcerpc.CallOption) (*QueryVDisksResponse, error) {
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
	out := &QueryVDisksResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVDiskProviderClient) CreateVDisk(ctx context.Context, in *CreateVDiskRequest, opts ...dcerpc.CallOption) (*CreateVDiskResponse, error) {
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
	out := &CreateVDiskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVDiskProviderClient) AddVDisk(ctx context.Context, in *AddVDiskRequest, opts ...dcerpc.CallOption) (*AddVDiskResponse, error) {
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
	out := &AddVDiskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVDiskProviderClient) GetDiskFromVDisk(ctx context.Context, in *GetDiskFromVDiskRequest, opts ...dcerpc.CallOption) (*GetDiskFromVDiskResponse, error) {
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
	out := &GetDiskFromVDiskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVDiskProviderClient) GetVDiskFromDisk(ctx context.Context, in *GetVDiskFromDiskRequest, opts ...dcerpc.CallOption) (*GetVDiskFromDiskResponse, error) {
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
	out := &GetVDiskFromDiskResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultVDiskProviderClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultVDiskProviderClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultVDiskProviderClient) IPID(ctx context.Context, ipid *dcom.IPID) VDiskProviderClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultVDiskProviderClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewVDiskProviderClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (VDiskProviderClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(VDiskProviderSyntaxV0_0))...)
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
	return &xxx_DefaultVDiskProviderClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_QueryVDisksOperation structure represents the QueryVDisks operation
type xxx_QueryVDisksOperation struct {
	This   *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat  `idl:"name:That" json:"that"`
	Enum   *vds.EnumObject `idl:"name:ppEnum" json:"enum"`
	Return int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryVDisksOperation) OpNum() int { return 3 }

func (o *xxx_QueryVDisksOperation) OpName() string { return "/IVdsVdProvider/v0/QueryVDisks" }

func (o *xxx_QueryVDisksOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryVDisksOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryVDisksOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_QueryVDisksOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryVDisksOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryVDisksOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// QueryVDisksRequest structure represents the QueryVDisks operation request
type QueryVDisksRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *QueryVDisksRequest) xxx_ToOp(ctx context.Context) *xxx_QueryVDisksOperation {
	if o == nil {
		return &xxx_QueryVDisksOperation{}
	}
	return &xxx_QueryVDisksOperation{
		This: o.This,
	}
}

func (o *QueryVDisksRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryVDisksOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *QueryVDisksRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryVDisksRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryVDisksOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryVDisksResponse structure represents the QueryVDisks operation response
type QueryVDisksResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppEnum: A pointer to an IEnumVdsObject (section 3.1.1.1) interface. If the operation
	// is successfully completed, the pointer receives the IEnumVdsObject interface of the
	// object, which contains an enumeration of virtual disk objects in the provider. Callers
	// MUST release the interface when they are finished with it.
	Enum *vds.EnumObject `idl:"name:ppEnum" json:"enum"`
	// Return: The QueryVDisks return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryVDisksResponse) xxx_ToOp(ctx context.Context) *xxx_QueryVDisksOperation {
	if o == nil {
		return &xxx_QueryVDisksOperation{}
	}
	return &xxx_QueryVDisksOperation{
		That:   o.That,
		Enum:   o.Enum,
		Return: o.Return,
	}
}

func (o *QueryVDisksResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryVDisksOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Enum = op.Enum
	o.Return = op.Return
}
func (o *QueryVDisksResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryVDisksResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryVDisksOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateVDiskOperation structure represents the CreateVDisk operation
type xxx_CreateVDiskOperation struct {
	This                     *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That                     *dcom.ORPCThat             `idl:"name:That" json:"that"`
	VirtualDeviceType        *vds.VirtualStorageType    `idl:"name:VirtualDeviceType" json:"virtual_device_type"`
	Path                     string                     `idl:"name:pPath;string" json:"path"`
	StringSecurityDescriptor string                     `idl:"name:pStringSecurityDescriptor;string;pointer:unique" json:"string_security_descriptor"`
	Flags                    vds.CreateVirtualDiskFlag  `idl:"name:Flags" json:"flags"`
	ProviderSpecificFlags    uint32                     `idl:"name:ProviderSpecificFlags" json:"provider_specific_flags"`
	_                        uint32                     `idl:"name:Reserved"`
	CreateDiskParameters     *vds.CreateVDiskParameters `idl:"name:pCreateDiskParameters" json:"create_disk_parameters"`
	Async                    *vds.Async                 `idl:"name:ppAsync;pointer:unique" json:"async"`
	Return                   int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateVDiskOperation) OpNum() int { return 4 }

func (o *xxx_CreateVDiskOperation) OpName() string { return "/IVdsVdProvider/v0/CreateVDisk" }

func (o *xxx_CreateVDiskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVDiskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// VirtualDeviceType {in} (1:{alias=PVIRTUAL_STORAGE_TYPE}*(1))(2:{alias=VIRTUAL_STORAGE_TYPE}(struct))
	{
		if o.VirtualDeviceType != nil {
			if err := o.VirtualDeviceType.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.VirtualStorageType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pPath {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
			return err
		}
	}
	// pStringSecurityDescriptor {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.StringSecurityDescriptor != "" {
			_ptr_pStringSecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.StringSecurityDescriptor); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.StringSecurityDescriptor, _ptr_pStringSecurityDescriptor); err != nil {
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
	// Flags {in} (1:{alias=CREATE_VIRTUAL_DISK_FLAG}(enum))
	{
		if err := w.WriteEnum(uint16(o.Flags)); err != nil {
			return err
		}
	}
	// ProviderSpecificFlags {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.ProviderSpecificFlags); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{alias=ULONG}(uint32))
	{
		// reserved Reserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	// pCreateDiskParameters {in} (1:{alias=PVDS_CREATE_VDISK_PARAMETERS}*(1))(2:{alias=VDS_CREATE_VDISK_PARAMETERS}(struct))
	{
		if o.CreateDiskParameters != nil {
			if err := o.CreateDiskParameters.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.CreateVDiskParameters{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppAsync {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IVdsAsync}(interface))
	{
		if o.Async != nil {
			_ptr_ppAsync := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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
	return nil
}

func (o *xxx_CreateVDiskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// VirtualDeviceType {in} (1:{alias=PVIRTUAL_STORAGE_TYPE,pointer=ref}*(1))(2:{alias=VIRTUAL_STORAGE_TYPE}(struct))
	{
		if o.VirtualDeviceType == nil {
			o.VirtualDeviceType = &vds.VirtualStorageType{}
		}
		if err := o.VirtualDeviceType.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pPath {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	// pStringSecurityDescriptor {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pStringSecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.StringSecurityDescriptor); err != nil {
				return err
			}
			return nil
		})
		_s_pStringSecurityDescriptor := func(ptr interface{}) { o.StringSecurityDescriptor = *ptr.(*string) }
		if err := w.ReadPointer(&o.StringSecurityDescriptor, _s_pStringSecurityDescriptor, _ptr_pStringSecurityDescriptor); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Flags {in} (1:{alias=CREATE_VIRTUAL_DISK_FLAG}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Flags)); err != nil {
			return err
		}
	}
	// ProviderSpecificFlags {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.ProviderSpecificFlags); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{alias=ULONG}(uint32))
	{
		// reserved Reserved
		var _Reserved uint32
		if err := w.ReadData(&_Reserved); err != nil {
			return err
		}
	}
	// pCreateDiskParameters {in} (1:{alias=PVDS_CREATE_VDISK_PARAMETERS,pointer=ref}*(1))(2:{alias=VDS_CREATE_VDISK_PARAMETERS}(struct))
	{
		if o.CreateDiskParameters == nil {
			o.CreateDiskParameters = &vds.CreateVDiskParameters{}
		}
		if err := o.CreateDiskParameters.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppAsync {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IVdsAsync}(interface))
	{
		_ptr_ppAsync := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
	return nil
}

func (o *xxx_CreateVDiskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVDiskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppAsync {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IVdsAsync}(interface))
	{
		if o.Async != nil {
			_ptr_ppAsync := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateVDiskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppAsync {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=IVdsAsync}(interface))
	{
		_ptr_ppAsync := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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

// CreateVDiskRequest structure represents the CreateVDisk operation request
type CreateVDiskRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// VirtualDeviceType: Pointer to a VIRTUAL_STORAGE_TYPE (section 2.2.1.3.23) structure
	// that specifies the type of virtual hard disk to be created.
	VirtualDeviceType *vds.VirtualStorageType `idl:"name:VirtualDeviceType" json:"virtual_device_type"`
	// pPath: A NULL-terminated wide-character string containing the name and directory
	// path for the backing file to be created for the virtual hard disk.
	Path string `idl:"name:pPath;string" json:"path"`
	// pStringSecurityDescriptor: A NULL-terminated wide-character string containing the
	// security descriptor to be applied to the virtual disk. Security descriptors MUST
	// be in the Security Descriptor Definition Language (see [MSDN-SDDLforDevObj]).<77>
	// If this parameter is NULL, the security descriptor in the caller's access token (see
	// [MSFT-WSM/WEDWNK]) MUST be used.
	StringSecurityDescriptor string `idl:"name:pStringSecurityDescriptor;string;pointer:unique" json:"string_security_descriptor"`
	// Flags: Bitmask of flags specifying how the virtual disk is to be created.
	Flags vds.CreateVirtualDiskFlag `idl:"name:Flags" json:"flags"`
	// ProviderSpecificFlags: A bitmask of flags that are specific to the type of virtual
	// hard disk that is being surfaced. These flags are provider-specific.<78>
	ProviderSpecificFlags uint32 `idl:"name:ProviderSpecificFlags" json:"provider_specific_flags"`
	// pCreateDiskParameters: Pointer to a VDS_CREATE_VDISK_PARAMETERS (section 2.2.2.18.2.1)
	// structure that contains the virtual hard disk creation parameters.
	CreateDiskParameters *vds.CreateVDiskParameters `idl:"name:pCreateDiskParameters" json:"create_disk_parameters"`
	// ppAsync: A pointer to an IVdsAsync (section 3.1.3.1) interface that, if the operation
	// is successfully completed, receives the IVdsAsync interface to monitor and control
	// this operation. Callers MUST release the interface when they are done with it.
	Async *vds.Async `idl:"name:ppAsync;pointer:unique" json:"async"`
}

func (o *CreateVDiskRequest) xxx_ToOp(ctx context.Context) *xxx_CreateVDiskOperation {
	if o == nil {
		return &xxx_CreateVDiskOperation{}
	}
	return &xxx_CreateVDiskOperation{
		This:                     o.This,
		VirtualDeviceType:        o.VirtualDeviceType,
		Path:                     o.Path,
		StringSecurityDescriptor: o.StringSecurityDescriptor,
		Flags:                    o.Flags,
		ProviderSpecificFlags:    o.ProviderSpecificFlags,
		CreateDiskParameters:     o.CreateDiskParameters,
		Async:                    o.Async,
	}
}

func (o *CreateVDiskRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateVDiskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VirtualDeviceType = op.VirtualDeviceType
	o.Path = op.Path
	o.StringSecurityDescriptor = op.StringSecurityDescriptor
	o.Flags = op.Flags
	o.ProviderSpecificFlags = op.ProviderSpecificFlags
	o.CreateDiskParameters = op.CreateDiskParameters
	o.Async = op.Async
}
func (o *CreateVDiskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateVDiskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateVDiskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateVDiskResponse structure represents the CreateVDisk operation response
type CreateVDiskResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync (section 3.1.3.1) interface that, if the operation
	// is successfully completed, receives the IVdsAsync interface to monitor and control
	// this operation. Callers MUST release the interface when they are done with it.
	Async *vds.Async `idl:"name:ppAsync;pointer:unique" json:"async"`
	// Return: The CreateVDisk return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateVDiskResponse) xxx_ToOp(ctx context.Context) *xxx_CreateVDiskOperation {
	if o == nil {
		return &xxx_CreateVDiskOperation{}
	}
	return &xxx_CreateVDiskOperation{
		That:   o.That,
		Async:  o.Async,
		Return: o.Return,
	}
}

func (o *CreateVDiskResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateVDiskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *CreateVDiskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateVDiskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateVDiskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddVDiskOperation structure represents the AddVDisk operation
type xxx_AddVDiskOperation struct {
	This              *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat          `idl:"name:That" json:"that"`
	VirtualDeviceType *vds.VirtualStorageType `idl:"name:VirtualDeviceType" json:"virtual_device_type"`
	Path              string                  `idl:"name:pPath;string" json:"path"`
	VDisk             *vds.VDisk              `idl:"name:ppVDisk" json:"v_disk"`
	Return            int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_AddVDiskOperation) OpNum() int { return 5 }

func (o *xxx_AddVDiskOperation) OpName() string { return "/IVdsVdProvider/v0/AddVDisk" }

func (o *xxx_AddVDiskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddVDiskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// VirtualDeviceType {in} (1:{alias=PVIRTUAL_STORAGE_TYPE}*(1))(2:{alias=VIRTUAL_STORAGE_TYPE}(struct))
	{
		if o.VirtualDeviceType != nil {
			if err := o.VirtualDeviceType.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.VirtualStorageType{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pPath {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Path); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddVDiskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// VirtualDeviceType {in} (1:{alias=PVIRTUAL_STORAGE_TYPE,pointer=ref}*(1))(2:{alias=VIRTUAL_STORAGE_TYPE}(struct))
	{
		if o.VirtualDeviceType == nil {
			o.VirtualDeviceType = &vds.VirtualStorageType{}
		}
		if err := o.VirtualDeviceType.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pPath {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Path); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddVDiskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddVDiskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppVDisk {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsVDisk}(interface))
	{
		if o.VDisk != nil {
			_ptr_ppVDisk := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.VDisk != nil {
					if err := o.VDisk.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.VDisk{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.VDisk, _ptr_ppVDisk); err != nil {
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

func (o *xxx_AddVDiskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppVDisk {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsVDisk}(interface))
	{
		_ptr_ppVDisk := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.VDisk == nil {
				o.VDisk = &vds.VDisk{}
			}
			if err := o.VDisk.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppVDisk := func(ptr interface{}) { o.VDisk = *ptr.(**vds.VDisk) }
		if err := w.ReadPointer(&o.VDisk, _s_ppVDisk, _ptr_ppVDisk); err != nil {
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

// AddVDiskRequest structure represents the AddVDisk operation request
type AddVDiskRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// VirtualDeviceType: A pointer to a VIRTUAL_STORAGE_TYPE (section 2.2.1.3.23) structure
	// that specifies the type of virtual hard disk to open.
	VirtualDeviceType *vds.VirtualStorageType `idl:"name:VirtualDeviceType" json:"virtual_device_type"`
	// pPath: A NULL-terminated wide-character string containing the fully qualified pathname
	// for the virtual disk's backing file.
	Path string `idl:"name:pPath;string" json:"path"`
}

func (o *AddVDiskRequest) xxx_ToOp(ctx context.Context) *xxx_AddVDiskOperation {
	if o == nil {
		return &xxx_AddVDiskOperation{}
	}
	return &xxx_AddVDiskOperation{
		This:              o.This,
		VirtualDeviceType: o.VirtualDeviceType,
		Path:              o.Path,
	}
}

func (o *AddVDiskRequest) xxx_FromOp(ctx context.Context, op *xxx_AddVDiskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VirtualDeviceType = op.VirtualDeviceType
	o.Path = op.Path
}
func (o *AddVDiskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddVDiskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddVDiskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddVDiskResponse structure represents the AddVDisk operation response
type AddVDiskResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	VDisk *vds.VDisk     `idl:"name:ppVDisk" json:"v_disk"`
	// Return: The AddVDisk return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AddVDiskResponse) xxx_ToOp(ctx context.Context) *xxx_AddVDiskOperation {
	if o == nil {
		return &xxx_AddVDiskOperation{}
	}
	return &xxx_AddVDiskOperation{
		That:   o.That,
		VDisk:  o.VDisk,
		Return: o.Return,
	}
}

func (o *AddVDiskResponse) xxx_FromOp(ctx context.Context, op *xxx_AddVDiskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.VDisk = op.VDisk
	o.Return = op.Return
}
func (o *AddVDiskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddVDiskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddVDiskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDiskFromVDiskOperation structure represents the GetDiskFromVDisk operation
type xxx_GetDiskFromVDiskOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	VDisk  *vds.VDisk     `idl:"name:pVDisk" json:"v_disk"`
	Disk   *vds.Disk      `idl:"name:ppDisk" json:"disk"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDiskFromVDiskOperation) OpNum() int { return 6 }

func (o *xxx_GetDiskFromVDiskOperation) OpName() string { return "/IVdsVdProvider/v0/GetDiskFromVDisk" }

func (o *xxx_GetDiskFromVDiskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDiskFromVDiskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pVDisk {in} (1:{pointer=ref}*(1))(2:{alias=IVdsVDisk}(interface))
	{
		if o.VDisk != nil {
			_ptr_pVDisk := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.VDisk != nil {
					if err := o.VDisk.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.VDisk{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.VDisk, _ptr_pVDisk); err != nil {
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

func (o *xxx_GetDiskFromVDiskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pVDisk {in} (1:{pointer=ref}*(1))(2:{alias=IVdsVDisk}(interface))
	{
		_ptr_pVDisk := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.VDisk == nil {
				o.VDisk = &vds.VDisk{}
			}
			if err := o.VDisk.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pVDisk := func(ptr interface{}) { o.VDisk = *ptr.(**vds.VDisk) }
		if err := w.ReadPointer(&o.VDisk, _s_pVDisk, _ptr_pVDisk); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDiskFromVDiskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDiskFromVDiskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppDisk {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsDisk}(interface))
	{
		if o.Disk != nil {
			_ptr_ppDisk := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Disk != nil {
					if err := o.Disk.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.Disk{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Disk, _ptr_ppDisk); err != nil {
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

func (o *xxx_GetDiskFromVDiskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppDisk {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsDisk}(interface))
	{
		_ptr_ppDisk := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Disk == nil {
				o.Disk = &vds.Disk{}
			}
			if err := o.Disk.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppDisk := func(ptr interface{}) { o.Disk = *ptr.(**vds.Disk) }
		if err := w.ReadPointer(&o.Disk, _s_ppDisk, _ptr_ppDisk); err != nil {
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

// GetDiskFromVDiskRequest structure represents the GetDiskFromVDisk operation request
type GetDiskFromVDiskRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pVDisk: The IVdsVDisk interface pointer for the virtual disk.
	VDisk *vds.VDisk `idl:"name:pVDisk" json:"v_disk"`
}

func (o *GetDiskFromVDiskRequest) xxx_ToOp(ctx context.Context) *xxx_GetDiskFromVDiskOperation {
	if o == nil {
		return &xxx_GetDiskFromVDiskOperation{}
	}
	return &xxx_GetDiskFromVDiskOperation{
		This:  o.This,
		VDisk: o.VDisk,
	}
}

func (o *GetDiskFromVDiskRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDiskFromVDiskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.VDisk = op.VDisk
}
func (o *GetDiskFromVDiskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDiskFromVDiskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDiskFromVDiskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDiskFromVDiskResponse structure represents the GetDiskFromVDisk operation response
type GetDiskFromVDiskResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	Disk *vds.Disk      `idl:"name:ppDisk" json:"disk"`
	// Return: The GetDiskFromVDisk return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDiskFromVDiskResponse) xxx_ToOp(ctx context.Context) *xxx_GetDiskFromVDiskOperation {
	if o == nil {
		return &xxx_GetDiskFromVDiskOperation{}
	}
	return &xxx_GetDiskFromVDiskOperation{
		That:   o.That,
		Disk:   o.Disk,
		Return: o.Return,
	}
}

func (o *GetDiskFromVDiskResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDiskFromVDiskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Disk = op.Disk
	o.Return = op.Return
}
func (o *GetDiskFromVDiskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDiskFromVDiskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDiskFromVDiskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetVDiskFromDiskOperation structure represents the GetVDiskFromDisk operation
type xxx_GetVDiskFromDiskOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Disk   *vds.Disk      `idl:"name:pDisk" json:"disk"`
	VDisk  *vds.VDisk     `idl:"name:ppVDisk" json:"v_disk"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetVDiskFromDiskOperation) OpNum() int { return 7 }

func (o *xxx_GetVDiskFromDiskOperation) OpName() string { return "/IVdsVdProvider/v0/GetVDiskFromDisk" }

func (o *xxx_GetVDiskFromDiskOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVDiskFromDiskOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pDisk {in} (1:{pointer=ref}*(1))(2:{alias=IVdsDisk}(interface))
	{
		if o.Disk != nil {
			_ptr_pDisk := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Disk != nil {
					if err := o.Disk.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.Disk{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Disk, _ptr_pDisk); err != nil {
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

func (o *xxx_GetVDiskFromDiskOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pDisk {in} (1:{pointer=ref}*(1))(2:{alias=IVdsDisk}(interface))
	{
		_ptr_pDisk := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Disk == nil {
				o.Disk = &vds.Disk{}
			}
			if err := o.Disk.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pDisk := func(ptr interface{}) { o.Disk = *ptr.(**vds.Disk) }
		if err := w.ReadPointer(&o.Disk, _s_pDisk, _ptr_pDisk); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVDiskFromDiskOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVDiskFromDiskOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppVDisk {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsVDisk}(interface))
	{
		if o.VDisk != nil {
			_ptr_ppVDisk := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.VDisk != nil {
					if err := o.VDisk.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&vds.VDisk{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.VDisk, _ptr_ppVDisk); err != nil {
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

func (o *xxx_GetVDiskFromDiskOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppVDisk {out} (1:{pointer=ref}*(2)*(1))(2:{alias=IVdsVDisk}(interface))
	{
		_ptr_ppVDisk := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.VDisk == nil {
				o.VDisk = &vds.VDisk{}
			}
			if err := o.VDisk.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppVDisk := func(ptr interface{}) { o.VDisk = *ptr.(**vds.VDisk) }
		if err := w.ReadPointer(&o.VDisk, _s_ppVDisk, _ptr_ppVDisk); err != nil {
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

// GetVDiskFromDiskRequest structure represents the GetVDiskFromDisk operation request
type GetVDiskFromDiskRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pDisk: The IVdsDisk interface pointer to a disk.
	Disk *vds.Disk `idl:"name:pDisk" json:"disk"`
}

func (o *GetVDiskFromDiskRequest) xxx_ToOp(ctx context.Context) *xxx_GetVDiskFromDiskOperation {
	if o == nil {
		return &xxx_GetVDiskFromDiskOperation{}
	}
	return &xxx_GetVDiskFromDiskOperation{
		This: o.This,
		Disk: o.Disk,
	}
}

func (o *GetVDiskFromDiskRequest) xxx_FromOp(ctx context.Context, op *xxx_GetVDiskFromDiskOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Disk = op.Disk
}
func (o *GetVDiskFromDiskRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetVDiskFromDiskRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVDiskFromDiskOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetVDiskFromDiskResponse structure represents the GetVDiskFromDisk operation response
type GetVDiskFromDiskResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That  *dcom.ORPCThat `idl:"name:That" json:"that"`
	VDisk *vds.VDisk     `idl:"name:ppVDisk" json:"v_disk"`
	// Return: The GetVDiskFromDisk return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetVDiskFromDiskResponse) xxx_ToOp(ctx context.Context) *xxx_GetVDiskFromDiskOperation {
	if o == nil {
		return &xxx_GetVDiskFromDiskOperation{}
	}
	return &xxx_GetVDiskFromDiskOperation{
		That:   o.That,
		VDisk:  o.VDisk,
		Return: o.Return,
	}
}

func (o *GetVDiskFromDiskResponse) xxx_FromOp(ctx context.Context, op *xxx_GetVDiskFromDiskOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.VDisk = op.VDisk
	o.Return = op.Return
}
func (o *GetVDiskFromDiskResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetVDiskFromDiskResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVDiskFromDiskOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
