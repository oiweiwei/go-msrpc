package ivdsopenvdisk

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
	// IVdsOpenVDisk interface identifier 75c8f324-f715-4fe3-a28e-f9011b61a4a1
	OpenVDiskIID = &dcom.IID{Data1: 0x75c8f324, Data2: 0xf715, Data3: 0x4fe3, Data4: []byte{0xa2, 0x8e, 0xf9, 0x01, 0x1b, 0x61, 0xa4, 0xa1}}
	// Syntax UUID
	OpenVDiskSyntaxUUID = &uuid.UUID{TimeLow: 0x75c8f324, TimeMid: 0xf715, TimeHiAndVersion: 0x4fe3, ClockSeqHiAndReserved: 0xa2, ClockSeqLow: 0x8e, Node: [6]uint8{0xf9, 0x1, 0x1b, 0x61, 0xa4, 0xa1}}
	// Syntax ID
	OpenVDiskSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: OpenVDiskSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsOpenVDisk interface.
type OpenVDiskClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The Attach method creates an operating system disk device for a virtual disk.
	//
	// Return Values: This method MUST return zero to indicate success, or return an implementation-specific
	// nonzero error code to indicate failure.
	Attach(context.Context, *AttachRequest, ...dcerpc.CallOption) (*AttachResponse, error)

	// The Detach method removes the operating system disk device that represents a virtual
	// disk.
	//
	// Return Values: The method MUST return zero to indicate success, or return an implementation-specific
	// nonzero error code to indicate failure.
	Detach(context.Context, *DetachRequest, ...dcerpc.CallOption) (*DetachResponse, error)

	// The DetachAndDelete method removes the operating system disk device that represents
	// a virtual disk and deletes any backing store file.
	//
	// Return Values: The method MUST return zero to indicate success, or return an implementation-specific
	// nonzero error code to indicate failure.
	DetachAndDelete(context.Context, *DetachAndDeleteRequest, ...dcerpc.CallOption) (*DetachAndDeleteResponse, error)

	// The Compact method reduces the size of the virtual disk file (the backing store).
	// This requires that the virtual disk be detached. Compact is applicable only to differencing
	// type virtual disks and virtual disks created using CREATE_VIRTUAL_DISK_FLAG_NONE.
	// The Compact method does not change the size of the virtual disk.
	//
	// Return Values: The method MUST return zero to indicate success, or return an implementation-specific
	// nonzero error code to indicate failure.
	Compact(context.Context, *CompactRequest, ...dcerpc.CallOption) (*CompactResponse, error)

	//	This method is applicable only to differencing type virtual disks. The Merge method moves all data blocks from a differencing virtual disk into its parent virtual disk. Merging a virtual disk requires that the virtual disk be detached during the operation. Both the virtual disk and its parent are opened READ|WRITE using the IVdsVDisk::Open method called against the virtual disk with an appropriate value for the ReadWriteDepth, as described later in this section.<145>
	//
	// For example, to merge a differencing disk that is a child of a single parent disk
	// into that parent disk, call the IVdsVDisk::Open method on the child disk with the
	// ReadWriteDepth parameter set to the value 2. This value opens both disks with the
	// READ and WRITE flags set, which is necessary for disks to be merged with subsequent
	// call to the IVdsOpenVDisk::Merge method.
	//
	// Return Values: The method MUST return zero to indicate success, or return an implementation-specific
	// nonzero error code to indicate failure.
	Merge(context.Context, *MergeRequest, ...dcerpc.CallOption) (*MergeResponse, error)

	//	The Expand method increases the size of a virtual disk. Expanding a virtual disk requires that the virtual disk be detached during the operation. The virtual disk file is opened with READ|WRITE privileges using the IVdsVDisk::Open method.<147>
	//
	// Return Values: The method MUST return zero to indicate success, or return an implementation-specific
	// nonzero error code to indicate failure.
	Expand(context.Context, *ExpandRequest, ...dcerpc.CallOption) (*ExpandResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) OpenVDiskClient
}

type xxx_DefaultOpenVDiskClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultOpenVDiskClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultOpenVDiskClient) Attach(ctx context.Context, in *AttachRequest, opts ...dcerpc.CallOption) (*AttachResponse, error) {
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
	out := &AttachResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOpenVDiskClient) Detach(ctx context.Context, in *DetachRequest, opts ...dcerpc.CallOption) (*DetachResponse, error) {
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
	out := &DetachResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOpenVDiskClient) DetachAndDelete(ctx context.Context, in *DetachAndDeleteRequest, opts ...dcerpc.CallOption) (*DetachAndDeleteResponse, error) {
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
	out := &DetachAndDeleteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOpenVDiskClient) Compact(ctx context.Context, in *CompactRequest, opts ...dcerpc.CallOption) (*CompactResponse, error) {
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
	out := &CompactResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOpenVDiskClient) Merge(ctx context.Context, in *MergeRequest, opts ...dcerpc.CallOption) (*MergeResponse, error) {
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
	out := &MergeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOpenVDiskClient) Expand(ctx context.Context, in *ExpandRequest, opts ...dcerpc.CallOption) (*ExpandResponse, error) {
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
	out := &ExpandResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultOpenVDiskClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultOpenVDiskClient) IPID(ctx context.Context, ipid *dcom.IPID) OpenVDiskClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultOpenVDiskClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewOpenVDiskClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (OpenVDiskClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(OpenVDiskSyntaxV0_0))...)
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
	return &xxx_DefaultOpenVDiskClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_AttachOperation structure represents the Attach operation
type xxx_AttachOperation struct {
	This                     *dcom.ORPCThis            `idl:"name:This" json:"this"`
	That                     *dcom.ORPCThat            `idl:"name:That" json:"that"`
	StringSecurityDescriptor string                    `idl:"name:pStringSecurityDescriptor;pointer:unique" json:"string_security_descriptor"`
	Flags                    vds.AttachVirtualDiskFlag `idl:"name:Flags" json:"flags"`
	ProviderSpecificFlags    uint32                    `idl:"name:ProviderSpecificFlags" json:"provider_specific_flags"`
	TimeoutInMs              uint32                    `idl:"name:TimeoutInMs" json:"timeout_in_ms"`
	Async                    *vds.Async                `idl:"name:ppAsync" json:"async"`
	Return                   int32                     `idl:"name:Return" json:"return"`
}

func (o *xxx_AttachOperation) OpNum() int { return 3 }

func (o *xxx_AttachOperation) OpName() string { return "/IVdsOpenVDisk/v0/Attach" }

func (o *xxx_AttachOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AttachOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pStringSecurityDescriptor {in} (1:{pointer=unique, alias=LPWSTR}*(1)[dim:0,string](wchar))
	{
		if o.StringSecurityDescriptor != "" {
			_ptr_pStringSecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.StringSecurityDescriptor); err != nil {
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
	// Flags {in} (1:{alias=ATTACH_VIRTUAL_DISK_FLAG}(enum))
	{
		if err := w.WriteData(uint16(o.Flags)); err != nil {
			return err
		}
	}
	// ProviderSpecificFlags {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.ProviderSpecificFlags); err != nil {
			return err
		}
	}
	// TimeoutInMs {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.TimeoutInMs); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AttachOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pStringSecurityDescriptor {in} (1:{pointer=unique, alias=LPWSTR}*(1)[dim:0,string](wchar))
	{
		_ptr_pStringSecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.StringSecurityDescriptor); err != nil {
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
	// Flags {in} (1:{alias=ATTACH_VIRTUAL_DISK_FLAG}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Flags)); err != nil {
			return err
		}
	}
	// ProviderSpecificFlags {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.ProviderSpecificFlags); err != nil {
			return err
		}
	}
	// TimeoutInMs {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.TimeoutInMs); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AttachOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AttachOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AttachOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AttachRequest structure represents the Attach operation request
type AttachRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pStringSecurityDescriptor: A NULL-terminated wide-character string containing the
	// security descriptor to be applied to the virtual disk.<140> If this parameter is
	// NULL, the security descriptor in the caller's access token MUST be used.
	StringSecurityDescriptor string `idl:"name:pStringSecurityDescriptor;pointer:unique" json:"string_security_descriptor"`
	// Flags: A bitmask of ATTACH_VIRTUAL_DISK_FLAG (section 2.2.2.20.1.1) enumeration values
	// specifying virtual disk attaching options.
	Flags vds.AttachVirtualDiskFlag `idl:"name:Flags" json:"flags"`
	// ProviderSpecificFlags: A bitmask of flags that are specific to the type of virtual
	// disk that is being surfaced. These flags are provider-specific.<141>
	ProviderSpecificFlags uint32 `idl:"name:ProviderSpecificFlags" json:"provider_specific_flags"`
	// TimeoutInMs: The length of time, in milliseconds, before this method MAY<142> return
	// after waiting for the virtual disk to be surfaced completely. If this parameter is
	// zero, the method returns immediately without waiting for the disk to be surfaced.
	// If this parameter is INFINITE, the method does not return until the surfacing operation
	// is complete. If this parameter is set to a value other than zero or INFINITE and
	// the time-out value is reached, the method guarantees that the disk is not surfaced
	// after the operation is complete.
	TimeoutInMs uint32 `idl:"name:TimeoutInMs" json:"timeout_in_ms"`
}

func (o *AttachRequest) xxx_ToOp(ctx context.Context) *xxx_AttachOperation {
	if o == nil {
		return &xxx_AttachOperation{}
	}
	return &xxx_AttachOperation{
		This:                     o.This,
		StringSecurityDescriptor: o.StringSecurityDescriptor,
		Flags:                    o.Flags,
		ProviderSpecificFlags:    o.ProviderSpecificFlags,
		TimeoutInMs:              o.TimeoutInMs,
	}
}

func (o *AttachRequest) xxx_FromOp(ctx context.Context, op *xxx_AttachOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.StringSecurityDescriptor = op.StringSecurityDescriptor
	o.Flags = op.Flags
	o.ProviderSpecificFlags = op.ProviderSpecificFlags
	o.TimeoutInMs = op.TimeoutInMs
}
func (o *AttachRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AttachRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AttachOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AttachResponse structure represents the Attach operation response
type AttachResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync (section 3.1.3.1) interface that, if the operation
	// is successfully completed, receives the IVdsAsync interface to monitor and control
	// this operation. Callers MUST release the interface when they are done with it.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The Attach return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AttachResponse) xxx_ToOp(ctx context.Context) *xxx_AttachOperation {
	if o == nil {
		return &xxx_AttachOperation{}
	}
	return &xxx_AttachOperation{
		That:   o.That,
		Async:  o.Async,
		Return: o.Return,
	}
}

func (o *AttachResponse) xxx_FromOp(ctx context.Context, op *xxx_AttachOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *AttachResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AttachResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AttachOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DetachOperation structure represents the Detach operation
type xxx_DetachOperation struct {
	This                  *dcom.ORPCThis            `idl:"name:This" json:"this"`
	That                  *dcom.ORPCThat            `idl:"name:That" json:"that"`
	Flags                 vds.DetachVirtualDiskFlag `idl:"name:Flags" json:"flags"`
	ProviderSpecificFlags uint32                    `idl:"name:ProviderSpecificFlags" json:"provider_specific_flags"`
	Return                int32                     `idl:"name:Return" json:"return"`
}

func (o *xxx_DetachOperation) OpNum() int { return 4 }

func (o *xxx_DetachOperation) OpName() string { return "/IVdsOpenVDisk/v0/Detach" }

func (o *xxx_DetachOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DetachOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Flags {in} (1:{alias=DETACH_VIRTUAL_DISK_FLAG}(enum))
	{
		if err := w.WriteData(uint16(o.Flags)); err != nil {
			return err
		}
	}
	// ProviderSpecificFlags {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.ProviderSpecificFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DetachOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Flags {in} (1:{alias=DETACH_VIRTUAL_DISK_FLAG}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Flags)); err != nil {
			return err
		}
	}
	// ProviderSpecificFlags {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.ProviderSpecificFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DetachOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DetachOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DetachOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DetachRequest structure represents the Detach operation request
type DetachRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// Flags: A DETACH_VIRTUAL_DISK_FLAG (section 2.2.2.20.1.2) enumeration value that specifies
	// how the virtual disk is to be detached.
	Flags vds.DetachVirtualDiskFlag `idl:"name:Flags" json:"flags"`
	// ProviderSpecificFlags: Flags specific to the type of virtual disk being detached.<143>
	ProviderSpecificFlags uint32 `idl:"name:ProviderSpecificFlags" json:"provider_specific_flags"`
}

func (o *DetachRequest) xxx_ToOp(ctx context.Context) *xxx_DetachOperation {
	if o == nil {
		return &xxx_DetachOperation{}
	}
	return &xxx_DetachOperation{
		This:                  o.This,
		Flags:                 o.Flags,
		ProviderSpecificFlags: o.ProviderSpecificFlags,
	}
}

func (o *DetachRequest) xxx_FromOp(ctx context.Context, op *xxx_DetachOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flags = op.Flags
	o.ProviderSpecificFlags = op.ProviderSpecificFlags
}
func (o *DetachRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DetachRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DetachOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DetachResponse structure represents the Detach operation response
type DetachResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Detach return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DetachResponse) xxx_ToOp(ctx context.Context) *xxx_DetachOperation {
	if o == nil {
		return &xxx_DetachOperation{}
	}
	return &xxx_DetachOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *DetachResponse) xxx_FromOp(ctx context.Context, op *xxx_DetachOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DetachResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DetachResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DetachOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DetachAndDeleteOperation structure represents the DetachAndDelete operation
type xxx_DetachAndDeleteOperation struct {
	This                  *dcom.ORPCThis            `idl:"name:This" json:"this"`
	That                  *dcom.ORPCThat            `idl:"name:That" json:"that"`
	Flags                 vds.DetachVirtualDiskFlag `idl:"name:Flags" json:"flags"`
	ProviderSpecificFlags uint32                    `idl:"name:ProviderSpecificFlags" json:"provider_specific_flags"`
	Return                int32                     `idl:"name:Return" json:"return"`
}

func (o *xxx_DetachAndDeleteOperation) OpNum() int { return 5 }

func (o *xxx_DetachAndDeleteOperation) OpName() string { return "/IVdsOpenVDisk/v0/DetachAndDelete" }

func (o *xxx_DetachAndDeleteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DetachAndDeleteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Flags {in} (1:{alias=DETACH_VIRTUAL_DISK_FLAG}(enum))
	{
		if err := w.WriteData(uint16(o.Flags)); err != nil {
			return err
		}
	}
	// ProviderSpecificFlags {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.ProviderSpecificFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DetachAndDeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Flags {in} (1:{alias=DETACH_VIRTUAL_DISK_FLAG}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Flags)); err != nil {
			return err
		}
	}
	// ProviderSpecificFlags {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.ProviderSpecificFlags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DetachAndDeleteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DetachAndDeleteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DetachAndDeleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DetachAndDeleteRequest structure represents the DetachAndDelete operation request
type DetachAndDeleteRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// Flags: A DETACH_VIRTUAL_DISK_FLAG (section 2.2.2.20.1.2) enumeration value that specifies
	// how the virtual disk is to be detached.
	Flags vds.DetachVirtualDiskFlag `idl:"name:Flags" json:"flags"`
	// ProviderSpecificFlags: Flags specific to the type of virtual disk being detached.<144>
	ProviderSpecificFlags uint32 `idl:"name:ProviderSpecificFlags" json:"provider_specific_flags"`
}

func (o *DetachAndDeleteRequest) xxx_ToOp(ctx context.Context) *xxx_DetachAndDeleteOperation {
	if o == nil {
		return &xxx_DetachAndDeleteOperation{}
	}
	return &xxx_DetachAndDeleteOperation{
		This:                  o.This,
		Flags:                 o.Flags,
		ProviderSpecificFlags: o.ProviderSpecificFlags,
	}
}

func (o *DetachAndDeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_DetachAndDeleteOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flags = op.Flags
	o.ProviderSpecificFlags = op.ProviderSpecificFlags
}
func (o *DetachAndDeleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DetachAndDeleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DetachAndDeleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DetachAndDeleteResponse structure represents the DetachAndDelete operation response
type DetachAndDeleteResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DetachAndDelete return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DetachAndDeleteResponse) xxx_ToOp(ctx context.Context) *xxx_DetachAndDeleteOperation {
	if o == nil {
		return &xxx_DetachAndDeleteOperation{}
	}
	return &xxx_DetachAndDeleteOperation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *DetachAndDeleteResponse) xxx_FromOp(ctx context.Context, op *xxx_DetachAndDeleteOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DetachAndDeleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DetachAndDeleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DetachAndDeleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CompactOperation structure represents the Compact operation
type xxx_CompactOperation struct {
	This   *dcom.ORPCThis             `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat             `idl:"name:That" json:"that"`
	Flags  vds.CompactVirtualDiskFlag `idl:"name:Flags" json:"flags"`
	_      uint32                     `idl:"name:Reserved"`
	Async  *vds.Async                 `idl:"name:ppAsync" json:"async"`
	Return int32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_CompactOperation) OpNum() int { return 6 }

func (o *xxx_CompactOperation) OpName() string { return "/IVdsOpenVDisk/v0/Compact" }

func (o *xxx_CompactOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CompactOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Flags {in} (1:{alias=COMPACT_VIRTUAL_DISK_FLAG}(enum))
	{
		if err := w.WriteData(uint16(o.Flags)); err != nil {
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
	return nil
}

func (o *xxx_CompactOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Flags {in} (1:{alias=COMPACT_VIRTUAL_DISK_FLAG}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Flags)); err != nil {
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
	return nil
}

func (o *xxx_CompactOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CompactOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CompactOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CompactRequest structure represents the Compact operation request
type CompactRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// Flags: A COMPACT_VIRTUAL_DISK_FLAG (section 2.2.2.20.1.3) enumeration value that
	// specifies how the virtual disk is to be compacted.
	Flags vds.CompactVirtualDiskFlag `idl:"name:Flags" json:"flags"`
}

func (o *CompactRequest) xxx_ToOp(ctx context.Context) *xxx_CompactOperation {
	if o == nil {
		return &xxx_CompactOperation{}
	}
	return &xxx_CompactOperation{
		This:  o.This,
		Flags: o.Flags,
	}
}

func (o *CompactRequest) xxx_FromOp(ctx context.Context, op *xxx_CompactOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flags = op.Flags
}
func (o *CompactRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CompactRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CompactOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CompactResponse structure represents the Compact operation response
type CompactResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync (section 3.1.3.1) interface that if the operation
	// is successfully completed receives the IVdsAsync interface to monitor and control
	// this operation. Callers MUST release the interface when they are done with it.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The Compact return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CompactResponse) xxx_ToOp(ctx context.Context) *xxx_CompactOperation {
	if o == nil {
		return &xxx_CompactOperation{}
	}
	return &xxx_CompactOperation{
		That:   o.That,
		Async:  o.Async,
		Return: o.Return,
	}
}

func (o *CompactResponse) xxx_FromOp(ctx context.Context, op *xxx_CompactOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *CompactResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CompactResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CompactOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MergeOperation structure represents the Merge operation
type xxx_MergeOperation struct {
	This       *dcom.ORPCThis           `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat           `idl:"name:That" json:"that"`
	Flags      vds.MergeVirtualDiskFlag `idl:"name:Flags" json:"flags"`
	MergeDepth uint32                   `idl:"name:MergeDepth" json:"merge_depth"`
	Async      *vds.Async               `idl:"name:ppAsync" json:"async"`
	Return     int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_MergeOperation) OpNum() int { return 7 }

func (o *xxx_MergeOperation) OpName() string { return "/IVdsOpenVDisk/v0/Merge" }

func (o *xxx_MergeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MergeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Flags {in} (1:{alias=MERGE_VIRTUAL_DISK_FLAG}(enum))
	{
		if err := w.WriteData(uint16(o.Flags)); err != nil {
			return err
		}
	}
	// MergeDepth {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.MergeDepth); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MergeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Flags {in} (1:{alias=MERGE_VIRTUAL_DISK_FLAG}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Flags)); err != nil {
			return err
		}
	}
	// MergeDepth {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.MergeDepth); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MergeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MergeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_MergeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// MergeRequest structure represents the Merge operation request
type MergeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// Flags: A MERGE_VIRTUAL_DISK_FLAG (section 2.2.2.20.1.4) enumeration value that specifies
	// how the virtual disk is to be merged.
	Flags vds.MergeVirtualDiskFlag `idl:"name:Flags" json:"flags"`
	// MergeDepth: Number of parent backing store files in the differencing chain to be
	// updated. For example, if MergeDepth has a value of 1, the data blocks from the given
	// differencing disk are moved into its parent. If the given differencing disk's parent
	// is also a differencing disk, (in other words the given disk is diskA, its parent
	// is diskB, and diskB's parent is diskC), and the MergeDepth parameter value is 2,
	// the data blocks from the given differencing disk (diskA) are moved into its parent
	// (diskB), and then its parent's (diskB's) data blocks are moved into its parent (diskC).<146>
	MergeDepth uint32 `idl:"name:MergeDepth" json:"merge_depth"`
}

func (o *MergeRequest) xxx_ToOp(ctx context.Context) *xxx_MergeOperation {
	if o == nil {
		return &xxx_MergeOperation{}
	}
	return &xxx_MergeOperation{
		This:       o.This,
		Flags:      o.Flags,
		MergeDepth: o.MergeDepth,
	}
}

func (o *MergeRequest) xxx_FromOp(ctx context.Context, op *xxx_MergeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flags = op.Flags
	o.MergeDepth = op.MergeDepth
}
func (o *MergeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *MergeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MergeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MergeResponse structure represents the Merge operation response
type MergeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync (section 3.1.3.1) interface that, if the operation
	// is successfully completed, receives the IVdsAsync interface to monitor and control
	// this operation. Callers MUST release the interface when they are done with it.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The Merge return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *MergeResponse) xxx_ToOp(ctx context.Context) *xxx_MergeOperation {
	if o == nil {
		return &xxx_MergeOperation{}
	}
	return &xxx_MergeOperation{
		That:   o.That,
		Async:  o.Async,
		Return: o.Return,
	}
}

func (o *MergeResponse) xxx_FromOp(ctx context.Context, op *xxx_MergeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *MergeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *MergeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MergeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ExpandOperation structure represents the Expand operation
type xxx_ExpandOperation struct {
	This    *dcom.ORPCThis            `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat            `idl:"name:That" json:"that"`
	Flags   vds.ExpandVirtualDiskFlag `idl:"name:Flags" json:"flags"`
	NewSize uint64                    `idl:"name:NewSize" json:"new_size"`
	Async   *vds.Async                `idl:"name:ppAsync" json:"async"`
	Return  int32                     `idl:"name:Return" json:"return"`
}

func (o *xxx_ExpandOperation) OpNum() int { return 8 }

func (o *xxx_ExpandOperation) OpName() string { return "/IVdsOpenVDisk/v0/Expand" }

func (o *xxx_ExpandOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExpandOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Flags {in} (1:{alias=EXPAND_VIRTUAL_DISK_FLAG}(enum))
	{
		if err := w.WriteData(uint16(o.Flags)); err != nil {
			return err
		}
	}
	// NewSize {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.NewSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExpandOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Flags {in} (1:{alias=EXPAND_VIRTUAL_DISK_FLAG}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Flags)); err != nil {
			return err
		}
	}
	// NewSize {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.NewSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExpandOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ExpandOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ExpandOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ExpandRequest structure represents the Expand operation request
type ExpandRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// Flags: An EXPAND_VIRTUAL_DISK_FLAG (section 2.2.2.20.1.5) enumeration value that
	// specifies how the virtual disk is to be compacted.
	Flags vds.ExpandVirtualDiskFlag `idl:"name:Flags" json:"flags"`
	// NewSize: The desired size, in bytes, of the expanded virtual disk.
	NewSize uint64 `idl:"name:NewSize" json:"new_size"`
}

func (o *ExpandRequest) xxx_ToOp(ctx context.Context) *xxx_ExpandOperation {
	if o == nil {
		return &xxx_ExpandOperation{}
	}
	return &xxx_ExpandOperation{
		This:    o.This,
		Flags:   o.Flags,
		NewSize: o.NewSize,
	}
}

func (o *ExpandRequest) xxx_FromOp(ctx context.Context, op *xxx_ExpandOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flags = op.Flags
	o.NewSize = op.NewSize
}
func (o *ExpandRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ExpandRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExpandOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ExpandResponse structure represents the Expand operation response
type ExpandResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync (section 3.1.3.1) interface that, if the operation
	// is successfully completed, receives the IVdsAsync interface to monitor and control
	// this operation. Callers MUST release the interface when they are done with it.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The Expand return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ExpandResponse) xxx_ToOp(ctx context.Context) *xxx_ExpandOperation {
	if o == nil {
		return &xxx_ExpandOperation{}
	}
	return &xxx_ExpandOperation{
		That:   o.That,
		Async:  o.Async,
		Return: o.Return,
	}
}

func (o *ExpandResponse) xxx_FromOp(ctx context.Context, op *xxx_ExpandOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *ExpandResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ExpandResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ExpandOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
