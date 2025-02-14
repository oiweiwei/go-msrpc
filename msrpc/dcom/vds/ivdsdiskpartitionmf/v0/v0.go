package ivdsdiskpartitionmf

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
	// IVdsDiskPartitionMF interface identifier 538684e0-ba3d-4bc0-aca9-164aff85c2a9
	DiskPartitionMFIID = &dcom.IID{Data1: 0x538684e0, Data2: 0xba3d, Data3: 0x4bc0, Data4: []byte{0xac, 0xa9, 0x16, 0x4a, 0xff, 0x85, 0xc2, 0xa9}}
	// Syntax UUID
	DiskPartitionMFSyntaxUUID = &uuid.UUID{TimeLow: 0x538684e0, TimeMid: 0xba3d, TimeHiAndVersion: 0x4bc0, ClockSeqHiAndReserved: 0xac, ClockSeqLow: 0xa9, Node: [6]uint8{0x16, 0x4a, 0xff, 0x85, 0xc2, 0xa9}}
	// Syntax ID
	DiskPartitionMFSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: DiskPartitionMFSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsDiskPartitionMF interface.
type DiskPartitionMFClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The GetPartitionFileSystemProperties method returns property details about the file
	// system on a disk partition at a specified byte offset. This method is only supported
	// on OEM, ESP, recovery, and unknown partitions.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// ERROR_SUCCESS (0x00000000)
	GetPartitionFileSystemProperties(context.Context, *GetPartitionFileSystemPropertiesRequest, ...dcerpc.CallOption) (*GetPartitionFileSystemPropertiesResponse, error)

	// The GetPartitionFileSystemTypeName method retrieves the name of the file system on
	// a disk partition at a specified byte offset. This method is only supported on OEM,
	// ESP, recovery, and unknown partitions.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetPartitionFileSystemTypeName(context.Context, *GetPartitionFileSystemTypeNameRequest, ...dcerpc.CallOption) (*GetPartitionFileSystemTypeNameResponse, error)

	// The QueryPartitionFileSystemFormatSupport method retrieves the properties of the
	// file systems that support formatting a disk partition at a specified byte offset.
	// This method is only supported on OEM, ESP, recovery, and unknown partitions.
	//
	// Note  This method is not valid on CD/DVD or super floppy devices. These devices do
	// not support partition tables.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.<107>
	//
	// ERROR_SUCCESS (0x00000000)
	QueryPartitionFileSystemFormatSupport(context.Context, *QueryPartitionFileSystemFormatSupportRequest, ...dcerpc.CallOption) (*QueryPartitionFileSystemFormatSupportResponse, error)

	// The FormatPartitionEx method formats an existing OEM, ESP, or unknown partition.
	// This method is only supported on OEM, ESP, recovery, and unknown partitions.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// ERROR_SUCCESS (0x00000000)
	FormatPartitionEx(context.Context, *FormatPartitionExRequest, ...dcerpc.CallOption) (*FormatPartitionExResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) DiskPartitionMFClient
}

type xxx_DefaultDiskPartitionMFClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultDiskPartitionMFClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultDiskPartitionMFClient) GetPartitionFileSystemProperties(ctx context.Context, in *GetPartitionFileSystemPropertiesRequest, opts ...dcerpc.CallOption) (*GetPartitionFileSystemPropertiesResponse, error) {
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
	out := &GetPartitionFileSystemPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDiskPartitionMFClient) GetPartitionFileSystemTypeName(ctx context.Context, in *GetPartitionFileSystemTypeNameRequest, opts ...dcerpc.CallOption) (*GetPartitionFileSystemTypeNameResponse, error) {
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
	out := &GetPartitionFileSystemTypeNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDiskPartitionMFClient) QueryPartitionFileSystemFormatSupport(ctx context.Context, in *QueryPartitionFileSystemFormatSupportRequest, opts ...dcerpc.CallOption) (*QueryPartitionFileSystemFormatSupportResponse, error) {
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
	out := &QueryPartitionFileSystemFormatSupportResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDiskPartitionMFClient) FormatPartitionEx(ctx context.Context, in *FormatPartitionExRequest, opts ...dcerpc.CallOption) (*FormatPartitionExResponse, error) {
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
	out := &FormatPartitionExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDiskPartitionMFClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultDiskPartitionMFClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultDiskPartitionMFClient) IPID(ctx context.Context, ipid *dcom.IPID) DiskPartitionMFClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultDiskPartitionMFClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewDiskPartitionMFClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (DiskPartitionMFClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(DiskPartitionMFSyntaxV0_0))...)
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
	return &xxx_DefaultDiskPartitionMFClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetPartitionFileSystemPropertiesOperation structure represents the GetPartitionFileSystemProperties operation
type xxx_GetPartitionFileSystemPropertiesOperation struct {
	This               *dcom.ORPCThis          `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat          `idl:"name:That" json:"that"`
	Offset             uint64                  `idl:"name:ullOffset" json:"offset"`
	FileSystemProperty *vds.FileSystemProperty `idl:"name:pFileSystemProp" json:"file_system_property"`
	Return             int32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPartitionFileSystemPropertiesOperation) OpNum() int { return 3 }

func (o *xxx_GetPartitionFileSystemPropertiesOperation) OpName() string {
	return "/IVdsDiskPartitionMF/v0/GetPartitionFileSystemProperties"
}

func (o *xxx_GetPartitionFileSystemPropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPartitionFileSystemPropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ullOffset {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.Offset); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPartitionFileSystemPropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ullOffset {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.Offset); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPartitionFileSystemPropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPartitionFileSystemPropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPartitionFileSystemPropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// GetPartitionFileSystemPropertiesRequest structure represents the GetPartitionFileSystemProperties operation request
type GetPartitionFileSystemPropertiesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ullOffset: The byte offset of the partition from the beginning of the disk. This
	// MUST be the offset at the start of a partition.
	Offset uint64 `idl:"name:ullOffset" json:"offset"`
}

func (o *GetPartitionFileSystemPropertiesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPartitionFileSystemPropertiesOperation) *xxx_GetPartitionFileSystemPropertiesOperation {
	if op == nil {
		op = &xxx_GetPartitionFileSystemPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Offset = op.Offset
	return op
}

func (o *GetPartitionFileSystemPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPartitionFileSystemPropertiesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Offset = op.Offset
}
func (o *GetPartitionFileSystemPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPartitionFileSystemPropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPartitionFileSystemPropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPartitionFileSystemPropertiesResponse structure represents the GetPartitionFileSystemProperties operation response
type GetPartitionFileSystemPropertiesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pFileSystemProp: A pointer to a VDS_FILE_SYSTEM_PROP structure that, if the operation
	// is successfully completed, receives the properties of the file system on the partition.
	FileSystemProperty *vds.FileSystemProperty `idl:"name:pFileSystemProp" json:"file_system_property"`
	// Return: The GetPartitionFileSystemProperties return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPartitionFileSystemPropertiesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPartitionFileSystemPropertiesOperation) *xxx_GetPartitionFileSystemPropertiesOperation {
	if op == nil {
		op = &xxx_GetPartitionFileSystemPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.FileSystemProperty = op.FileSystemProperty
	o.Return = op.Return
	return op
}

func (o *GetPartitionFileSystemPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPartitionFileSystemPropertiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileSystemProperty = op.FileSystemProperty
	o.Return = op.Return
}
func (o *GetPartitionFileSystemPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPartitionFileSystemPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPartitionFileSystemPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPartitionFileSystemTypeNameOperation structure represents the GetPartitionFileSystemTypeName operation
type xxx_GetPartitionFileSystemTypeNameOperation struct {
	This               *dcom.ORPCThis `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat `idl:"name:That" json:"that"`
	Offset             uint64         `idl:"name:ullOffset" json:"offset"`
	FileSystemTypeName string         `idl:"name:ppwszFileSystemTypeName;string" json:"file_system_type_name"`
	Return             int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPartitionFileSystemTypeNameOperation) OpNum() int { return 4 }

func (o *xxx_GetPartitionFileSystemTypeNameOperation) OpName() string {
	return "/IVdsDiskPartitionMF/v0/GetPartitionFileSystemTypeName"
}

func (o *xxx_GetPartitionFileSystemTypeNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPartitionFileSystemTypeNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ullOffset {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.Offset); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPartitionFileSystemTypeNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ullOffset {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.Offset); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPartitionFileSystemTypeNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPartitionFileSystemTypeNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppwszFileSystemTypeName {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.FileSystemTypeName != "" {
			_ptr_ppwszFileSystemTypeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.FileSystemTypeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.FileSystemTypeName, _ptr_ppwszFileSystemTypeName); err != nil {
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

func (o *xxx_GetPartitionFileSystemTypeNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppwszFileSystemTypeName {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ppwszFileSystemTypeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.FileSystemTypeName); err != nil {
				return err
			}
			return nil
		})
		_s_ppwszFileSystemTypeName := func(ptr interface{}) { o.FileSystemTypeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.FileSystemTypeName, _s_ppwszFileSystemTypeName, _ptr_ppwszFileSystemTypeName); err != nil {
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

// GetPartitionFileSystemTypeNameRequest structure represents the GetPartitionFileSystemTypeName operation request
type GetPartitionFileSystemTypeNameRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ullOffset: The byte offset of the partition from the beginning of the disk. This
	// MUST be the offset at the start of a partition.
	Offset uint64 `idl:"name:ullOffset" json:"offset"`
}

func (o *GetPartitionFileSystemTypeNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPartitionFileSystemTypeNameOperation) *xxx_GetPartitionFileSystemTypeNameOperation {
	if op == nil {
		op = &xxx_GetPartitionFileSystemTypeNameOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Offset = op.Offset
	return op
}

func (o *GetPartitionFileSystemTypeNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPartitionFileSystemTypeNameOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Offset = op.Offset
}
func (o *GetPartitionFileSystemTypeNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPartitionFileSystemTypeNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPartitionFileSystemTypeNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPartitionFileSystemTypeNameResponse structure represents the GetPartitionFileSystemTypeName operation response
type GetPartitionFileSystemTypeNameResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppwszFileSystemTypeName: A pointer that, if the operation is successfully completed,
	// receives a null-terminated Unicode string with the file system name.
	FileSystemTypeName string `idl:"name:ppwszFileSystemTypeName;string" json:"file_system_type_name"`
	// Return: The GetPartitionFileSystemTypeName return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPartitionFileSystemTypeNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPartitionFileSystemTypeNameOperation) *xxx_GetPartitionFileSystemTypeNameOperation {
	if op == nil {
		op = &xxx_GetPartitionFileSystemTypeNameOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.FileSystemTypeName = op.FileSystemTypeName
	o.Return = op.Return
	return op
}

func (o *GetPartitionFileSystemTypeNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPartitionFileSystemTypeNameOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileSystemTypeName = op.FileSystemTypeName
	o.Return = op.Return
}
func (o *GetPartitionFileSystemTypeNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPartitionFileSystemTypeNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPartitionFileSystemTypeNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryPartitionFileSystemFormatSupportOperation structure represents the QueryPartitionFileSystemFormatSupport operation
type xxx_QueryPartitionFileSystemFormatSupportOperation struct {
	This                        *dcom.ORPCThis                         `idl:"name:This" json:"this"`
	That                        *dcom.ORPCThat                         `idl:"name:That" json:"that"`
	Offset                      uint64                                 `idl:"name:ullOffset" json:"offset"`
	FileSystemSupportProperties []*vds.FileSystemFormatSupportProperty `idl:"name:ppFileSystemSupportProps;size_is:(, plNumberOfFileSystems)" json:"file_system_support_properties"`
	NumberOfFileSystems         int32                                  `idl:"name:plNumberOfFileSystems" json:"number_of_file_systems"`
	Return                      int32                                  `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryPartitionFileSystemFormatSupportOperation) OpNum() int { return 5 }

func (o *xxx_QueryPartitionFileSystemFormatSupportOperation) OpName() string {
	return "/IVdsDiskPartitionMF/v0/QueryPartitionFileSystemFormatSupport"
}

func (o *xxx_QueryPartitionFileSystemFormatSupportOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryPartitionFileSystemFormatSupportOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ullOffset {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.Offset); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryPartitionFileSystemFormatSupportOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ullOffset {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.Offset); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryPartitionFileSystemFormatSupportOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.FileSystemSupportProperties != nil && o.NumberOfFileSystems == 0 {
		o.NumberOfFileSystems = int32(len(o.FileSystemSupportProperties))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryPartitionFileSystemFormatSupportOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppFileSystemSupportProps {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VDS_FILE_SYSTEM_FORMAT_SUPPORT_PROP}[dim:0,size_is=plNumberOfFileSystems](struct))
	{
		if o.FileSystemSupportProperties != nil || o.NumberOfFileSystems > 0 {
			_ptr_ppFileSystemSupportProps := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NumberOfFileSystems)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.FileSystemSupportProperties {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.FileSystemSupportProperties[i1] != nil {
						if err := o.FileSystemSupportProperties[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&vds.FileSystemFormatSupportProperty{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.FileSystemSupportProperties); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&vds.FileSystemFormatSupportProperty{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileSystemSupportProperties, _ptr_ppFileSystemSupportProps); err != nil {
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
	// plNumberOfFileSystems {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.NumberOfFileSystems); err != nil {
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

func (o *xxx_QueryPartitionFileSystemFormatSupportOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppFileSystemSupportProps {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VDS_FILE_SYSTEM_FORMAT_SUPPORT_PROP}[dim:0,size_is=plNumberOfFileSystems](struct))
	{
		_ptr_ppFileSystemSupportProps := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.FileSystemSupportProperties", sizeInfo[0])
			}
			o.FileSystemSupportProperties = make([]*vds.FileSystemFormatSupportProperty, sizeInfo[0])
			for i1 := range o.FileSystemSupportProperties {
				i1 := i1
				if o.FileSystemSupportProperties[i1] == nil {
					o.FileSystemSupportProperties[i1] = &vds.FileSystemFormatSupportProperty{}
				}
				if err := o.FileSystemSupportProperties[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppFileSystemSupportProps := func(ptr interface{}) { o.FileSystemSupportProperties = *ptr.(*[]*vds.FileSystemFormatSupportProperty) }
		if err := w.ReadPointer(&o.FileSystemSupportProperties, _s_ppFileSystemSupportProps, _ptr_ppFileSystemSupportProps); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// plNumberOfFileSystems {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.NumberOfFileSystems); err != nil {
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

// QueryPartitionFileSystemFormatSupportRequest structure represents the QueryPartitionFileSystemFormatSupport operation request
type QueryPartitionFileSystemFormatSupportRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ullOffset: The byte offset of the partition from the beginning of the disk. This
	// MUST be the offset at the start of a partition.
	Offset uint64 `idl:"name:ullOffset" json:"offset"`
}

func (o *QueryPartitionFileSystemFormatSupportRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryPartitionFileSystemFormatSupportOperation) *xxx_QueryPartitionFileSystemFormatSupportOperation {
	if op == nil {
		op = &xxx_QueryPartitionFileSystemFormatSupportOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Offset = op.Offset
	return op
}

func (o *QueryPartitionFileSystemFormatSupportRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryPartitionFileSystemFormatSupportOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Offset = op.Offset
}
func (o *QueryPartitionFileSystemFormatSupportRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryPartitionFileSystemFormatSupportRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryPartitionFileSystemFormatSupportOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryPartitionFileSystemFormatSupportResponse structure represents the QueryPartitionFileSystemFormatSupport operation response
type QueryPartitionFileSystemFormatSupportResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppFileSystemSupportProps: A pointer to an array of VDS_FILE_SYSTEM_FORMAT_SUPPORT_PROP
	// structures which, if the operation completes successfully, receives an array of properties
	// of each supported file system.
	FileSystemSupportProperties []*vds.FileSystemFormatSupportProperty `idl:"name:ppFileSystemSupportProps;size_is:(, plNumberOfFileSystems)" json:"file_system_support_properties"`
	// plNumberOfFileSystems: A pointer to a variable which, if the operation completes
	// successfully, receives the total number of elements returned in ppFileSystemSupportProps.
	NumberOfFileSystems int32 `idl:"name:plNumberOfFileSystems" json:"number_of_file_systems"`
	// Return: The QueryPartitionFileSystemFormatSupport return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryPartitionFileSystemFormatSupportResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryPartitionFileSystemFormatSupportOperation) *xxx_QueryPartitionFileSystemFormatSupportOperation {
	if op == nil {
		op = &xxx_QueryPartitionFileSystemFormatSupportOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.FileSystemSupportProperties = op.FileSystemSupportProperties
	o.NumberOfFileSystems = op.NumberOfFileSystems
	o.Return = op.Return
	return op
}

func (o *QueryPartitionFileSystemFormatSupportResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryPartitionFileSystemFormatSupportOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FileSystemSupportProperties = op.FileSystemSupportProperties
	o.NumberOfFileSystems = op.NumberOfFileSystems
	o.Return = op.Return
}
func (o *QueryPartitionFileSystemFormatSupportResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryPartitionFileSystemFormatSupportResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryPartitionFileSystemFormatSupportOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FormatPartitionExOperation structure represents the FormatPartitionEx operation
type xxx_FormatPartitionExOperation struct {
	This                      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Offset                    uint64         `idl:"name:ullOffset" json:"offset"`
	FileSystemTypeName        string         `idl:"name:pwszFileSystemTypeName;string;pointer:unique" json:"file_system_type_name"`
	FileSystemRevision        uint16         `idl:"name:usFileSystemRevision" json:"file_system_revision"`
	DesiredUnitAllocationSize uint32         `idl:"name:ulDesiredUnitAllocationSize" json:"desired_unit_allocation_size"`
	Label                     string         `idl:"name:pwszLabel;string;pointer:unique" json:"label"`
	Force                     int32          `idl:"name:bForce" json:"force"`
	QuickFormat               int32          `idl:"name:bQuickFormat" json:"quick_format"`
	EnableCompression         int32          `idl:"name:bEnableCompression" json:"enable_compression"`
	Async                     *vds.Async     `idl:"name:ppAsync" json:"async"`
	Return                    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_FormatPartitionExOperation) OpNum() int { return 6 }

func (o *xxx_FormatPartitionExOperation) OpName() string {
	return "/IVdsDiskPartitionMF/v0/FormatPartitionEx"
}

func (o *xxx_FormatPartitionExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatPartitionExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ullOffset {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.Offset); err != nil {
			return err
		}
	}
	// pwszFileSystemTypeName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.FileSystemTypeName != "" {
			_ptr_pwszFileSystemTypeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.FileSystemTypeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.FileSystemTypeName, _ptr_pwszFileSystemTypeName); err != nil {
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
	// usFileSystemRevision {in} (1:(uint16))
	{
		if err := w.WriteData(o.FileSystemRevision); err != nil {
			return err
		}
	}
	// ulDesiredUnitAllocationSize {in} (1:(uint32))
	{
		if err := w.WriteData(o.DesiredUnitAllocationSize); err != nil {
			return err
		}
	}
	// pwszLabel {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.Label != "" {
			_ptr_pwszLabel := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Label); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Label, _ptr_pwszLabel); err != nil {
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

func (o *xxx_FormatPartitionExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ullOffset {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.Offset); err != nil {
			return err
		}
	}
	// pwszFileSystemTypeName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_pwszFileSystemTypeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.FileSystemTypeName); err != nil {
				return err
			}
			return nil
		})
		_s_pwszFileSystemTypeName := func(ptr interface{}) { o.FileSystemTypeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.FileSystemTypeName, _s_pwszFileSystemTypeName, _ptr_pwszFileSystemTypeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// usFileSystemRevision {in} (1:(uint16))
	{
		if err := w.ReadData(&o.FileSystemRevision); err != nil {
			return err
		}
	}
	// ulDesiredUnitAllocationSize {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DesiredUnitAllocationSize); err != nil {
			return err
		}
	}
	// pwszLabel {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_pwszLabel := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Label); err != nil {
				return err
			}
			return nil
		})
		_s_pwszLabel := func(ptr interface{}) { o.Label = *ptr.(*string) }
		if err := w.ReadPointer(&o.Label, _s_pwszLabel, _ptr_pwszLabel); err != nil {
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

func (o *xxx_FormatPartitionExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatPartitionExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_FormatPartitionExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// FormatPartitionExRequest structure represents the FormatPartitionEx operation request
type FormatPartitionExRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ullOffset: The byte offset of the partition from the beginning of the disk. This
	// MUST be the offset at the start of a partition.
	Offset uint64 `idl:"name:ullOffset" json:"offset"`
	// pwszFileSystemTypeName: A null-terminated Unicode string that contains the name of
	// the file system with which to format the partition.
	FileSystemTypeName string `idl:"name:pwszFileSystemTypeName;string;pointer:unique" json:"file_system_type_name"`
	// usFileSystemRevision: A 16-bit, binary-coded decimal number that indicates the revision
	// of the file system, if any. The first two (most significant) digits (8-bits) indicate
	// the major revision while the last two (least significant) digits (8-bits) indicate
	// the minor revision (for example, 0x0250 represents revision 2.50).
	FileSystemRevision uint16 `idl:"name:usFileSystemRevision" json:"file_system_revision"`
	// ulDesiredUnitAllocationSize: The size, in bytes, of the allocation unit for the file
	// system. The value MUST be a power of 2. If the value is 0, a default allocation unit
	// determined by the file system type is used. The allocation unit range is file system-dependent.
	DesiredUnitAllocationSize uint32 `idl:"name:ulDesiredUnitAllocationSize" json:"desired_unit_allocation_size"`
	// pwszLabel: The null-terminated Unicode string to assign to the new file system. The
	// maximum label size is file system-dependent.
	Label string `idl:"name:pwszLabel;string;pointer:unique" json:"label"`
	// bForce: A Boolean that determines whether a file system format is forced, even if
	// the partition is in use.
	Force int32 `idl:"name:bForce" json:"force"`
	// bQuickFormat: A Boolean that determines whether a file system is quick formatted.
	// A quick format does not verify each sector on the volume.
	QuickFormat int32 `idl:"name:bQuickFormat" json:"quick_format"`
	// bEnableCompression: A Boolean that determines whether a file system is created with
	// compression enabled.<108>
	EnableCompression int32 `idl:"name:bEnableCompression" json:"enable_compression"`
}

func (o *FormatPartitionExRequest) xxx_ToOp(ctx context.Context, op *xxx_FormatPartitionExOperation) *xxx_FormatPartitionExOperation {
	if op == nil {
		op = &xxx_FormatPartitionExOperation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Offset = op.Offset
	o.FileSystemTypeName = op.FileSystemTypeName
	o.FileSystemRevision = op.FileSystemRevision
	o.DesiredUnitAllocationSize = op.DesiredUnitAllocationSize
	o.Label = op.Label
	o.Force = op.Force
	o.QuickFormat = op.QuickFormat
	o.EnableCompression = op.EnableCompression
	return op
}

func (o *FormatPartitionExRequest) xxx_FromOp(ctx context.Context, op *xxx_FormatPartitionExOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Offset = op.Offset
	o.FileSystemTypeName = op.FileSystemTypeName
	o.FileSystemRevision = op.FileSystemRevision
	o.DesiredUnitAllocationSize = op.DesiredUnitAllocationSize
	o.Label = op.Label
	o.Force = op.Force
	o.QuickFormat = op.QuickFormat
	o.EnableCompression = op.EnableCompression
}
func (o *FormatPartitionExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FormatPartitionExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FormatPartitionExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FormatPartitionExResponse structure represents the FormatPartitionEx operation response
type FormatPartitionExResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync interface that, if the operation is successfully
	// completed, receives the IVdsAsync interface to monitor and control this operation.
	// Callers MUST release the interface when they are done with it.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The FormatPartitionEx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FormatPartitionExResponse) xxx_ToOp(ctx context.Context, op *xxx_FormatPartitionExOperation) *xxx_FormatPartitionExOperation {
	if op == nil {
		op = &xxx_FormatPartitionExOperation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
	return op
}

func (o *FormatPartitionExResponse) xxx_FromOp(ctx context.Context, op *xxx_FormatPartitionExOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *FormatPartitionExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FormatPartitionExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FormatPartitionExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
