package ivdsadvanceddisk

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = vds.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/vds"
)

var (
	// IVdsAdvancedDisk interface identifier 6e6f6b40-977c-4069-bddd-ac710059f8c0
	AdvancedDiskIID = &dcom.IID{Data1: 0x6e6f6b40, Data2: 0x977c, Data3: 0x4069, Data4: []byte{0xbd, 0xdd, 0xac, 0x71, 0x00, 0x59, 0xf8, 0xc0}}
	// Syntax UUID
	AdvancedDiskSyntaxUUID = &uuid.UUID{TimeLow: 0x6e6f6b40, TimeMid: 0x977c, TimeHiAndVersion: 0x4069, ClockSeqHiAndReserved: 0xbd, ClockSeqLow: 0xdd, Node: [6]uint8{0xac, 0x71, 0x0, 0x59, 0xf8, 0xc0}}
	// Syntax ID
	AdvancedDiskSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AdvancedDiskSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsAdvancedDisk interface.
type AdvancedDiskClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The GetPartitionProperties method retrieves the properties of a partition on the
	// disk at a specified byte offset.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetPartitionProperties(context.Context, *GetPartitionPropertiesRequest, ...dcerpc.CallOption) (*GetPartitionPropertiesResponse, error)

	// The QueryPartitions method enumerates a disk's partitions.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryPartitions(context.Context, *QueryPartitionsRequest, ...dcerpc.CallOption) (*QueryPartitionsResponse, error)

	// The CreatePartition method creates a partition on a disk at a specified byte offset.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.<94>
	//
	// Note  Creating or deleting partitions on dynamic disks is not supported.
	CreatePartition(context.Context, *CreatePartitionRequest, ...dcerpc.CallOption) (*CreatePartitionResponse, error)

	// The DeletePartition method deletes a partition from the disk at a specified byte
	// offset.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// Note  Creating or deleting partitions on dynamic disks is not supported.
	//
	// ERROR_SUCCESS (0x00000000)
	DeletePartition(context.Context, *DeletePartitionRequest, ...dcerpc.CallOption) (*DeletePartitionResponse, error)

	// The ChangeAttributes method changes the attributes of the partition at byte offset
	// ullOffset on the disk.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	ChangeAttributes(context.Context, *ChangeAttributesRequest, ...dcerpc.CallOption) (*ChangeAttributesResponse, error)

	// The AssignDriveLetter method assigns a drive letter to an existing OEM, ESP, or unknown
	// partition.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	AssignDriveLetter(context.Context, *AssignDriveLetterRequest, ...dcerpc.CallOption) (*AssignDriveLetterResponse, error)

	// The DeleteDriveLetter method deletes a drive letter that is assigned to an OEM, ESP,
	// or unknown partition.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	DeleteDriveLetter(context.Context, *DeleteDriveLetterRequest, ...dcerpc.CallOption) (*DeleteDriveLetterResponse, error)

	// The GetDriveLetter method retrieves the drive letter of an OEM, ESP, or unknown partition
	// on the disk at a specified byte offset.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// ERROR_SUCCESS (0x00000000)
	GetDriveLetter(context.Context, *GetDriveLetterRequest, ...dcerpc.CallOption) (*GetDriveLetterResponse, error)

	// The FormatPartition method formats an existing OEM, ESP, or unknown partition.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// ERROR_SUCCESS (0x00000000)
	FormatPartition(context.Context, *FormatPartitionRequest, ...dcerpc.CallOption) (*FormatPartitionResponse, error)

	// The Clean method cleans a disk.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	Clean(context.Context, *CleanRequest, ...dcerpc.CallOption) (*CleanResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AdvancedDiskClient
}

type xxx_DefaultAdvancedDiskClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAdvancedDiskClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultAdvancedDiskClient) GetPartitionProperties(ctx context.Context, in *GetPartitionPropertiesRequest, opts ...dcerpc.CallOption) (*GetPartitionPropertiesResponse, error) {
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
	out := &GetPartitionPropertiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAdvancedDiskClient) QueryPartitions(ctx context.Context, in *QueryPartitionsRequest, opts ...dcerpc.CallOption) (*QueryPartitionsResponse, error) {
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
	out := &QueryPartitionsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAdvancedDiskClient) CreatePartition(ctx context.Context, in *CreatePartitionRequest, opts ...dcerpc.CallOption) (*CreatePartitionResponse, error) {
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
	out := &CreatePartitionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAdvancedDiskClient) DeletePartition(ctx context.Context, in *DeletePartitionRequest, opts ...dcerpc.CallOption) (*DeletePartitionResponse, error) {
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
	out := &DeletePartitionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAdvancedDiskClient) ChangeAttributes(ctx context.Context, in *ChangeAttributesRequest, opts ...dcerpc.CallOption) (*ChangeAttributesResponse, error) {
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
	out := &ChangeAttributesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAdvancedDiskClient) AssignDriveLetter(ctx context.Context, in *AssignDriveLetterRequest, opts ...dcerpc.CallOption) (*AssignDriveLetterResponse, error) {
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
	out := &AssignDriveLetterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAdvancedDiskClient) DeleteDriveLetter(ctx context.Context, in *DeleteDriveLetterRequest, opts ...dcerpc.CallOption) (*DeleteDriveLetterResponse, error) {
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
	out := &DeleteDriveLetterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAdvancedDiskClient) GetDriveLetter(ctx context.Context, in *GetDriveLetterRequest, opts ...dcerpc.CallOption) (*GetDriveLetterResponse, error) {
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
	out := &GetDriveLetterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAdvancedDiskClient) FormatPartition(ctx context.Context, in *FormatPartitionRequest, opts ...dcerpc.CallOption) (*FormatPartitionResponse, error) {
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
	out := &FormatPartitionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAdvancedDiskClient) Clean(ctx context.Context, in *CleanRequest, opts ...dcerpc.CallOption) (*CleanResponse, error) {
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
	out := &CleanResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAdvancedDiskClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAdvancedDiskClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAdvancedDiskClient) IPID(ctx context.Context, ipid *dcom.IPID) AdvancedDiskClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAdvancedDiskClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewAdvancedDiskClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AdvancedDiskClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AdvancedDiskSyntaxV0_0))...)
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
	return &xxx_DefaultAdvancedDiskClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetPartitionPropertiesOperation structure represents the GetPartitionProperties operation
type xxx_GetPartitionPropertiesOperation struct {
	This              *dcom.ORPCThis         `idl:"name:This" json:"this"`
	That              *dcom.ORPCThat         `idl:"name:That" json:"that"`
	Offset            uint64                 `idl:"name:ullOffset" json:"offset"`
	PartitionProperty *vds.PartitionProperty `idl:"name:pPartitionProp" json:"partition_property"`
	Return            int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPartitionPropertiesOperation) OpNum() int { return 3 }

func (o *xxx_GetPartitionPropertiesOperation) OpName() string {
	return "/IVdsAdvancedDisk/v0/GetPartitionProperties"
}

func (o *xxx_GetPartitionPropertiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPartitionPropertiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetPartitionPropertiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetPartitionPropertiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPartitionPropertiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pPartitionProp {out} (1:{pointer=ref}*(1))(2:{alias=VDS_PARTITION_PROP}(struct))
	{
		if o.PartitionProperty != nil {
			if err := o.PartitionProperty.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.PartitionProperty{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetPartitionPropertiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pPartitionProp {out} (1:{pointer=ref}*(1))(2:{alias=VDS_PARTITION_PROP}(struct))
	{
		if o.PartitionProperty == nil {
			o.PartitionProperty = &vds.PartitionProperty{}
		}
		if err := o.PartitionProperty.UnmarshalNDR(ctx, w); err != nil {
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

// GetPartitionPropertiesRequest structure represents the GetPartitionProperties operation request
type GetPartitionPropertiesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ullOffset: The byte offset of the partition, from the beginning of the disk. This
	// offset MUST be the offset of a start of a partition.
	Offset uint64 `idl:"name:ullOffset" json:"offset"`
}

func (o *GetPartitionPropertiesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPartitionPropertiesOperation) *xxx_GetPartitionPropertiesOperation {
	if op == nil {
		op = &xxx_GetPartitionPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Offset = o.Offset
	return op
}

func (o *GetPartitionPropertiesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPartitionPropertiesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Offset = op.Offset
}
func (o *GetPartitionPropertiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPartitionPropertiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPartitionPropertiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPartitionPropertiesResponse structure represents the GetPartitionProperties operation response
type GetPartitionPropertiesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pPartitionProp: A pointer to a VDS_PARTITION_PROP structure that, if the operation
	// is successfully completed, receives the properties of the partition.
	PartitionProperty *vds.PartitionProperty `idl:"name:pPartitionProp" json:"partition_property"`
	// Return: The GetPartitionProperties return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetPartitionPropertiesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPartitionPropertiesOperation) *xxx_GetPartitionPropertiesOperation {
	if op == nil {
		op = &xxx_GetPartitionPropertiesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PartitionProperty = o.PartitionProperty
	op.Return = o.Return
	return op
}

func (o *GetPartitionPropertiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPartitionPropertiesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PartitionProperty = op.PartitionProperty
	o.Return = op.Return
}
func (o *GetPartitionPropertiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPartitionPropertiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPartitionPropertiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryPartitionsOperation structure represents the QueryPartitions operation
type xxx_QueryPartitionsOperation struct {
	This                   *dcom.ORPCThis           `idl:"name:This" json:"this"`
	That                   *dcom.ORPCThat           `idl:"name:That" json:"that"`
	PartitionPropertyArray []*vds.PartitionProperty `idl:"name:ppPartitionPropArray;size_is:(, plNumberOfPartitions)" json:"partition_property_array"`
	NumberOfPartitions     int32                    `idl:"name:plNumberOfPartitions" json:"number_of_partitions"`
	Return                 int32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryPartitionsOperation) OpNum() int { return 4 }

func (o *xxx_QueryPartitionsOperation) OpName() string { return "/IVdsAdvancedDisk/v0/QueryPartitions" }

func (o *xxx_QueryPartitionsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryPartitionsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_QueryPartitionsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_QueryPartitionsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.PartitionPropertyArray != nil && o.NumberOfPartitions == 0 {
		o.NumberOfPartitions = int32(len(o.PartitionPropertyArray))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryPartitionsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppPartitionPropArray {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VDS_PARTITION_PROP}[dim:0,size_is=plNumberOfPartitions](struct))
	{
		if o.PartitionPropertyArray != nil || o.NumberOfPartitions > 0 {
			_ptr_ppPartitionPropArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NumberOfPartitions)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.PartitionPropertyArray {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.PartitionPropertyArray[i1] != nil {
						if err := o.PartitionPropertyArray[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&vds.PartitionProperty{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.PartitionPropertyArray); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&vds.PartitionProperty{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PartitionPropertyArray, _ptr_ppPartitionPropArray); err != nil {
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
	// plNumberOfPartitions {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.NumberOfPartitions); err != nil {
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

func (o *xxx_QueryPartitionsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppPartitionPropArray {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VDS_PARTITION_PROP}[dim:0,size_is=plNumberOfPartitions](struct))
	{
		_ptr_ppPartitionPropArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.PartitionPropertyArray", sizeInfo[0])
			}
			o.PartitionPropertyArray = make([]*vds.PartitionProperty, sizeInfo[0])
			for i1 := range o.PartitionPropertyArray {
				i1 := i1
				if o.PartitionPropertyArray[i1] == nil {
					o.PartitionPropertyArray[i1] = &vds.PartitionProperty{}
				}
				if err := o.PartitionPropertyArray[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppPartitionPropArray := func(ptr interface{}) { o.PartitionPropertyArray = *ptr.(*[]*vds.PartitionProperty) }
		if err := w.ReadPointer(&o.PartitionPropertyArray, _s_ppPartitionPropArray, _ptr_ppPartitionPropArray); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// plNumberOfPartitions {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.NumberOfPartitions); err != nil {
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

// QueryPartitionsRequest structure represents the QueryPartitions operation request
type QueryPartitionsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *QueryPartitionsRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryPartitionsOperation) *xxx_QueryPartitionsOperation {
	if op == nil {
		op = &xxx_QueryPartitionsOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *QueryPartitionsRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryPartitionsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *QueryPartitionsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryPartitionsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryPartitionsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryPartitionsResponse structure represents the QueryPartitions operation response
type QueryPartitionsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppPartitionPropArray: A pointer to an array of VDS_PARTITION_PROP structures that,
	// if the operation is successfully completed, receives the array of partition properties.
	PartitionPropertyArray []*vds.PartitionProperty `idl:"name:ppPartitionPropArray;size_is:(, plNumberOfPartitions)" json:"partition_property_array"`
	// plNumberOfPartitions: A pointer to a variable that, if the operation is successfully
	// completed, receives the total number of elements in ppPartitionPropArray.
	NumberOfPartitions int32 `idl:"name:plNumberOfPartitions" json:"number_of_partitions"`
	// Return: The QueryPartitions return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryPartitionsResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryPartitionsOperation) *xxx_QueryPartitionsOperation {
	if op == nil {
		op = &xxx_QueryPartitionsOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.PartitionPropertyArray = o.PartitionPropertyArray
	op.NumberOfPartitions = o.NumberOfPartitions
	op.Return = o.Return
	return op
}

func (o *QueryPartitionsResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryPartitionsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.PartitionPropertyArray = op.PartitionPropertyArray
	o.NumberOfPartitions = op.NumberOfPartitions
	o.Return = op.Return
}
func (o *QueryPartitionsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryPartitionsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryPartitionsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreatePartitionOperation structure represents the CreatePartition operation
type xxx_CreatePartitionOperation struct {
	This       *dcom.ORPCThis                 `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat                 `idl:"name:That" json:"that"`
	Offset     uint64                         `idl:"name:ullOffset" json:"offset"`
	Size       uint64                         `idl:"name:ullSize" json:"size"`
	Parameters *vds.CreatePartitionParameters `idl:"name:para" json:"parameters"`
	Async      *vds.Async                     `idl:"name:ppAsync" json:"async"`
	Return     int32                          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreatePartitionOperation) OpNum() int { return 5 }

func (o *xxx_CreatePartitionOperation) OpName() string { return "/IVdsAdvancedDisk/v0/CreatePartition" }

func (o *xxx_CreatePartitionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePartitionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ullSize {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.Size); err != nil {
			return err
		}
	}
	// para {in} (1:{pointer=ref}*(1))(2:{alias=CREATE_PARTITION_PARAMETERS}(struct))
	{
		if o.Parameters != nil {
			if err := o.Parameters.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.CreatePartitionParameters{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CreatePartitionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ullSize {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.Size); err != nil {
			return err
		}
	}
	// para {in} (1:{pointer=ref}*(1))(2:{alias=CREATE_PARTITION_PARAMETERS}(struct))
	{
		if o.Parameters == nil {
			o.Parameters = &vds.CreatePartitionParameters{}
		}
		if err := o.Parameters.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePartitionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreatePartitionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreatePartitionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreatePartitionRequest structure represents the CreatePartition operation request
type CreatePartitionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ullOffset: MUST be the byte offset from the beginning of the disk at which to create
	// the new partition.
	Offset uint64 `idl:"name:ullOffset" json:"offset"`
	// ullSize: MUST be the size of the new partition, in bytes.
	Size uint64 `idl:"name:ullSize" json:"size"`
	// para: MUST be a pointer to a CREATE_PARTITION_PARAMETERS structure that describes
	// the new partition to create.
	Parameters *vds.CreatePartitionParameters `idl:"name:para" json:"parameters"`
}

func (o *CreatePartitionRequest) xxx_ToOp(ctx context.Context, op *xxx_CreatePartitionOperation) *xxx_CreatePartitionOperation {
	if op == nil {
		op = &xxx_CreatePartitionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Offset = o.Offset
	op.Size = o.Size
	op.Parameters = o.Parameters
	return op
}

func (o *CreatePartitionRequest) xxx_FromOp(ctx context.Context, op *xxx_CreatePartitionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Offset = op.Offset
	o.Size = op.Size
	o.Parameters = op.Parameters
}
func (o *CreatePartitionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreatePartitionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreatePartitionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreatePartitionResponse structure represents the CreatePartition operation response
type CreatePartitionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: MUST be a pointer to an IVdsAsync interface that, upon successful completion,
	// receives the IVdsAsync interface to monitor and control this operation. Callers MUST
	// release the interface received when they are done with it. If the IVdsAsync::Wait
	// method is called on the interface, the interfaces returned in the VDS_ASYNC_OUTPUT
	// structure MUST be released as well.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The CreatePartition return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreatePartitionResponse) xxx_ToOp(ctx context.Context, op *xxx_CreatePartitionOperation) *xxx_CreatePartitionOperation {
	if op == nil {
		op = &xxx_CreatePartitionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Async = o.Async
	op.Return = o.Return
	return op
}

func (o *CreatePartitionResponse) xxx_FromOp(ctx context.Context, op *xxx_CreatePartitionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *CreatePartitionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreatePartitionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreatePartitionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeletePartitionOperation structure represents the DeletePartition operation
type xxx_DeletePartitionOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	Offset         uint64         `idl:"name:ullOffset" json:"offset"`
	Force          int32          `idl:"name:bForce" json:"force"`
	ForceProtected int32          `idl:"name:bForceProtected" json:"force_protected"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeletePartitionOperation) OpNum() int { return 6 }

func (o *xxx_DeletePartitionOperation) OpName() string { return "/IVdsAdvancedDisk/v0/DeletePartition" }

func (o *xxx_DeletePartitionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeletePartitionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bForce {in} (1:(int32))
	{
		if err := w.WriteData(o.Force); err != nil {
			return err
		}
	}
	// bForceProtected {in} (1:(int32))
	{
		if err := w.WriteData(o.ForceProtected); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeletePartitionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bForce {in} (1:(int32))
	{
		if err := w.ReadData(&o.Force); err != nil {
			return err
		}
	}
	// bForceProtected {in} (1:(int32))
	{
		if err := w.ReadData(&o.ForceProtected); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeletePartitionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeletePartitionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeletePartitionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeletePartitionRequest structure represents the DeletePartition operation request
type DeletePartitionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ullOffset: The byte offset of the partition, from the beginning of the disk. This
	// offset MUST be the offset at the start of a partition.
	Offset uint64 `idl:"name:ullOffset" json:"offset"`
	// bForce: A Boolean that determines whether the partition deletion is forced. If this
	// parameter is set to a nonzero value, calls to lock and dismount any file system on
	// the partition that fail are ignored. If this parameter is set to zero and any file
	// systems on the partition failed to be locked or dismounted, the server returns VDS_E_DEVICE_IN_USE.
	Force int32 `idl:"name:bForce" json:"force"`
	// bForceProtected: A Boolean value that determines whether deletion of a protected
	// partition  will be forced.<95>
	ForceProtected int32 `idl:"name:bForceProtected" json:"force_protected"`
}

func (o *DeletePartitionRequest) xxx_ToOp(ctx context.Context, op *xxx_DeletePartitionOperation) *xxx_DeletePartitionOperation {
	if op == nil {
		op = &xxx_DeletePartitionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Offset = o.Offset
	op.Force = o.Force
	op.ForceProtected = o.ForceProtected
	return op
}

func (o *DeletePartitionRequest) xxx_FromOp(ctx context.Context, op *xxx_DeletePartitionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Offset = op.Offset
	o.Force = op.Force
	o.ForceProtected = op.ForceProtected
}
func (o *DeletePartitionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeletePartitionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeletePartitionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeletePartitionResponse structure represents the DeletePartition operation response
type DeletePartitionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeletePartition return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeletePartitionResponse) xxx_ToOp(ctx context.Context, op *xxx_DeletePartitionOperation) *xxx_DeletePartitionOperation {
	if op == nil {
		op = &xxx_DeletePartitionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeletePartitionResponse) xxx_FromOp(ctx context.Context, op *xxx_DeletePartitionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeletePartitionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeletePartitionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeletePartitionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ChangeAttributesOperation structure represents the ChangeAttributes operation
type xxx_ChangeAttributesOperation struct {
	This       *dcom.ORPCThis                  `idl:"name:This" json:"this"`
	That       *dcom.ORPCThat                  `idl:"name:That" json:"that"`
	Offset     uint64                          `idl:"name:ullOffset" json:"offset"`
	Parameters *vds.ChangeAttributesParameters `idl:"name:para" json:"parameters"`
	Return     int32                           `idl:"name:Return" json:"return"`
}

func (o *xxx_ChangeAttributesOperation) OpNum() int { return 7 }

func (o *xxx_ChangeAttributesOperation) OpName() string {
	return "/IVdsAdvancedDisk/v0/ChangeAttributes"
}

func (o *xxx_ChangeAttributesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangeAttributesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// para {in} (1:{pointer=ref}*(1))(2:{alias=CHANGE_ATTRIBUTES_PARAMETERS}(struct))
	{
		if o.Parameters != nil {
			if err := o.Parameters.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.ChangeAttributesParameters{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ChangeAttributesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// para {in} (1:{pointer=ref}*(1))(2:{alias=CHANGE_ATTRIBUTES_PARAMETERS}(struct))
	{
		if o.Parameters == nil {
			o.Parameters = &vds.ChangeAttributesParameters{}
		}
		if err := o.Parameters.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangeAttributesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangeAttributesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ChangeAttributesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ChangeAttributesRequest structure represents the ChangeAttributes operation request
type ChangeAttributesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ullOffset: The byte offset of the partition, from the beginning of the disk. This
	// offset MUST be the offset of the start of a partition.
	Offset uint64 `idl:"name:ullOffset" json:"offset"`
	// para: A pointer to a CHANGE_ATTRIBUTES_PARAMETERS structure that describes the attributes
	// to change.
	Parameters *vds.ChangeAttributesParameters `idl:"name:para" json:"parameters"`
}

func (o *ChangeAttributesRequest) xxx_ToOp(ctx context.Context, op *xxx_ChangeAttributesOperation) *xxx_ChangeAttributesOperation {
	if op == nil {
		op = &xxx_ChangeAttributesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Offset = o.Offset
	op.Parameters = o.Parameters
	return op
}

func (o *ChangeAttributesRequest) xxx_FromOp(ctx context.Context, op *xxx_ChangeAttributesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Offset = op.Offset
	o.Parameters = op.Parameters
}
func (o *ChangeAttributesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ChangeAttributesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ChangeAttributesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ChangeAttributesResponse structure represents the ChangeAttributes operation response
type ChangeAttributesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ChangeAttributes return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ChangeAttributesResponse) xxx_ToOp(ctx context.Context, op *xxx_ChangeAttributesOperation) *xxx_ChangeAttributesOperation {
	if op == nil {
		op = &xxx_ChangeAttributesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ChangeAttributesResponse) xxx_FromOp(ctx context.Context, op *xxx_ChangeAttributesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ChangeAttributesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ChangeAttributesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ChangeAttributesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AssignDriveLetterOperation structure represents the AssignDriveLetter operation
type xxx_AssignDriveLetterOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Offset uint64         `idl:"name:ullOffset" json:"offset"`
	Letter uint16         `idl:"name:wcLetter" json:"letter"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_AssignDriveLetterOperation) OpNum() int { return 8 }

func (o *xxx_AssignDriveLetterOperation) OpName() string {
	return "/IVdsAdvancedDisk/v0/AssignDriveLetter"
}

func (o *xxx_AssignDriveLetterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AssignDriveLetterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// wcLetter {in} (1:{alias=WCHAR}(wchar))
	{
		if err := w.WriteData(o.Letter); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AssignDriveLetterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// wcLetter {in} (1:{alias=WCHAR}(wchar))
	{
		if err := w.ReadData(&o.Letter); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AssignDriveLetterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AssignDriveLetterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AssignDriveLetterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// AssignDriveLetterRequest structure represents the AssignDriveLetter operation request
type AssignDriveLetterRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ullOffset: The byte offset of the partition, from the beginning of the disk. This
	// offset MUST be the offset of a start of a partition.
	Offset uint64 `idl:"name:ullOffset" json:"offset"`
	// wcLetter: The drive letter to assign, as a single uppercase or lowercase alphabetical
	// (A-Z) Unicode character.
	Letter uint16 `idl:"name:wcLetter" json:"letter"`
}

func (o *AssignDriveLetterRequest) xxx_ToOp(ctx context.Context, op *xxx_AssignDriveLetterOperation) *xxx_AssignDriveLetterOperation {
	if op == nil {
		op = &xxx_AssignDriveLetterOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Offset = o.Offset
	op.Letter = o.Letter
	return op
}

func (o *AssignDriveLetterRequest) xxx_FromOp(ctx context.Context, op *xxx_AssignDriveLetterOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Offset = op.Offset
	o.Letter = op.Letter
}
func (o *AssignDriveLetterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AssignDriveLetterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AssignDriveLetterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AssignDriveLetterResponse structure represents the AssignDriveLetter operation response
type AssignDriveLetterResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The AssignDriveLetter return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AssignDriveLetterResponse) xxx_ToOp(ctx context.Context, op *xxx_AssignDriveLetterOperation) *xxx_AssignDriveLetterOperation {
	if op == nil {
		op = &xxx_AssignDriveLetterOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *AssignDriveLetterResponse) xxx_FromOp(ctx context.Context, op *xxx_AssignDriveLetterOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *AssignDriveLetterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AssignDriveLetterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AssignDriveLetterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteDriveLetterOperation structure represents the DeleteDriveLetter operation
type xxx_DeleteDriveLetterOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Offset uint64         `idl:"name:ullOffset" json:"offset"`
	Letter uint16         `idl:"name:wcLetter" json:"letter"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteDriveLetterOperation) OpNum() int { return 9 }

func (o *xxx_DeleteDriveLetterOperation) OpName() string {
	return "/IVdsAdvancedDisk/v0/DeleteDriveLetter"
}

func (o *xxx_DeleteDriveLetterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteDriveLetterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// wcLetter {in} (1:{alias=WCHAR}(wchar))
	{
		if err := w.WriteData(o.Letter); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteDriveLetterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// wcLetter {in} (1:{alias=WCHAR}(wchar))
	{
		if err := w.ReadData(&o.Letter); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteDriveLetterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteDriveLetterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_DeleteDriveLetterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// DeleteDriveLetterRequest structure represents the DeleteDriveLetter operation request
type DeleteDriveLetterRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ullOffset: The byte offset of the partition from the beginning of the disk. This
	// offset MUST be the offset of a start of a partition.
	Offset uint64 `idl:"name:ullOffset" json:"offset"`
	// wcLetter: The drive letter to delete as a single uppercase or lowercase alphabetical
	// (A-Z) Unicode character.
	Letter uint16 `idl:"name:wcLetter" json:"letter"`
}

func (o *DeleteDriveLetterRequest) xxx_ToOp(ctx context.Context, op *xxx_DeleteDriveLetterOperation) *xxx_DeleteDriveLetterOperation {
	if op == nil {
		op = &xxx_DeleteDriveLetterOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Offset = o.Offset
	op.Letter = o.Letter
	return op
}

func (o *DeleteDriveLetterRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteDriveLetterOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Offset = op.Offset
	o.Letter = op.Letter
}
func (o *DeleteDriveLetterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *DeleteDriveLetterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteDriveLetterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteDriveLetterResponse structure represents the DeleteDriveLetter operation response
type DeleteDriveLetterResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The DeleteDriveLetter return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DeleteDriveLetterResponse) xxx_ToOp(ctx context.Context, op *xxx_DeleteDriveLetterOperation) *xxx_DeleteDriveLetterOperation {
	if op == nil {
		op = &xxx_DeleteDriveLetterOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *DeleteDriveLetterResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteDriveLetterOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *DeleteDriveLetterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *DeleteDriveLetterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteDriveLetterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDriveLetterOperation structure represents the GetDriveLetter operation
type xxx_GetDriveLetterOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Offset uint64         `idl:"name:ullOffset" json:"offset"`
	Letter rune           `idl:"name:pwcLetter" json:"letter"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDriveLetterOperation) OpNum() int { return 10 }

func (o *xxx_GetDriveLetterOperation) OpName() string { return "/IVdsAdvancedDisk/v0/GetDriveLetter" }

func (o *xxx_GetDriveLetterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDriveLetterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetDriveLetterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetDriveLetterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDriveLetterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pwcLetter {out} (1:{rune, pointer=ref}*(1))(2:{alias=WCHAR}(wchar))
	{
		if err := w.WriteData(uint16(o.Letter)); err != nil {
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

func (o *xxx_GetDriveLetterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pwcLetter {out} (1:{rune, pointer=ref}*(1))(2:{alias=WCHAR}(wchar))
	{
		var _Letter uint16
		if err := w.ReadData(&_Letter); err != nil {
			return err
		}
		o.Letter = rune(_Letter)
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetDriveLetterRequest structure represents the GetDriveLetter operation request
type GetDriveLetterRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ullOffset: The byte offset of the partition, from the beginning of the disk. This
	// offset MUST be the offset of a start of a partition.
	Offset uint64 `idl:"name:ullOffset" json:"offset"`
}

func (o *GetDriveLetterRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDriveLetterOperation) *xxx_GetDriveLetterOperation {
	if op == nil {
		op = &xxx_GetDriveLetterOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Offset = o.Offset
	return op
}

func (o *GetDriveLetterRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDriveLetterOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Offset = op.Offset
}
func (o *GetDriveLetterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDriveLetterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDriveLetterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDriveLetterResponse structure represents the GetDriveLetter operation response
type GetDriveLetterResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pwcLetter: A pointer to a Unicode character that will receive an uppercase or lowercase
	// alphabetical (A-Z) drive letter for the partition at byte offset ullOffset.
	Letter rune `idl:"name:pwcLetter" json:"letter"`
	// Return: The GetDriveLetter return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDriveLetterResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDriveLetterOperation) *xxx_GetDriveLetterOperation {
	if op == nil {
		op = &xxx_GetDriveLetterOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Letter = o.Letter
	op.Return = o.Return
	return op
}

func (o *GetDriveLetterResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDriveLetterOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Letter = op.Letter
	o.Return = op.Return
}
func (o *GetDriveLetterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDriveLetterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDriveLetterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FormatPartitionOperation structure represents the FormatPartition operation
type xxx_FormatPartitionOperation struct {
	This               *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That               *dcom.ORPCThat     `idl:"name:That" json:"that"`
	Offset             uint64             `idl:"name:ullOffset" json:"offset"`
	Type               vds.FileSystemType `idl:"name:type" json:"type"`
	Label              string             `idl:"name:pwszLabel;string" json:"label"`
	UnitAllocationSize uint32             `idl:"name:dwUnitAllocationSize" json:"unit_allocation_size"`
	Force              int32              `idl:"name:bForce" json:"force"`
	QuickFormat        int32              `idl:"name:bQuickFormat" json:"quick_format"`
	EnableCompression  int32              `idl:"name:bEnableCompression" json:"enable_compression"`
	Async              *vds.Async         `idl:"name:ppAsync" json:"async"`
	Return             int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_FormatPartitionOperation) OpNum() int { return 11 }

func (o *xxx_FormatPartitionOperation) OpName() string { return "/IVdsAdvancedDisk/v0/FormatPartition" }

func (o *xxx_FormatPartitionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatPartitionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_FormatPartitionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_FormatPartitionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatPartitionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_FormatPartitionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// FormatPartitionRequest structure represents the FormatPartition operation request
type FormatPartitionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ullOffset: The byte offset of the partition, from the beginning of the disk. This
	// offset MUST be the offset of a start of a partition.
	Offset uint64 `idl:"name:ullOffset" json:"offset"`
	// type: A file system type that is enumerated by VDS_FILE_SYSTEM_TYPE. Clients that
	// want to format by using file systems that are not enumerated by VDS_FILE_SYSTEM_TYPE
	// (section 2.2.1.2.9) can use the IVdsDiskPartitionMF::FormatPartionEx (section 3.4.5.2.29.4)
	// or IVdsDiskPartitionMF2::FormatPartitionEx2 methods.
	Type vds.FileSystemType `idl:"name:type" json:"type"`
	// pwszLabel: A null-terminated Unicode string representing the partition label. The
	// maximum label size is file system-dependent.
	Label string `idl:"name:pwszLabel;string" json:"label"`
	// dwUnitAllocationSize: The size, in bytes, of the allocation unit for the file system.
	// The value MUST be a power of 2. Allocation unit range is file system-dependent.
	UnitAllocationSize uint32 `idl:"name:dwUnitAllocationSize" json:"unit_allocation_size"`
	// bForce: A Boolean that determines whether the format is forced, regardless of whether
	// the partition is in use.
	Force int32 `idl:"name:bForce" json:"force"`
	// bQuickFormat: A Boolean that determines whether a file system is quick formatted.
	// A quick format does not verify each sector on the volume.
	QuickFormat int32 `idl:"name:bQuickFormat" json:"quick_format"`
	// bEnableCompression: A Boolean that determines whether a file system is created with
	// compression enabled.<97>
	EnableCompression int32 `idl:"name:bEnableCompression" json:"enable_compression"`
}

func (o *FormatPartitionRequest) xxx_ToOp(ctx context.Context, op *xxx_FormatPartitionOperation) *xxx_FormatPartitionOperation {
	if op == nil {
		op = &xxx_FormatPartitionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Offset = o.Offset
	op.Type = o.Type
	op.Label = o.Label
	op.UnitAllocationSize = o.UnitAllocationSize
	op.Force = o.Force
	op.QuickFormat = o.QuickFormat
	op.EnableCompression = o.EnableCompression
	return op
}

func (o *FormatPartitionRequest) xxx_FromOp(ctx context.Context, op *xxx_FormatPartitionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Offset = op.Offset
	o.Type = op.Type
	o.Label = op.Label
	o.UnitAllocationSize = op.UnitAllocationSize
	o.Force = op.Force
	o.QuickFormat = op.QuickFormat
	o.EnableCompression = op.EnableCompression
}
func (o *FormatPartitionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FormatPartitionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FormatPartitionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FormatPartitionResponse structure represents the FormatPartition operation response
type FormatPartitionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync interface that, if the operation is successfully
	// completed, receives the IVdsAsync interface to monitor and control this operation.
	// Callers MUST release the interface when they are done with it.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The FormatPartition return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FormatPartitionResponse) xxx_ToOp(ctx context.Context, op *xxx_FormatPartitionOperation) *xxx_FormatPartitionOperation {
	if op == nil {
		op = &xxx_FormatPartitionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Async = o.Async
	op.Return = o.Return
	return op
}

func (o *FormatPartitionResponse) xxx_FromOp(ctx context.Context, op *xxx_FormatPartitionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *FormatPartitionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FormatPartitionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FormatPartitionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CleanOperation structure represents the Clean operation
type xxx_CleanOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Force     int32          `idl:"name:bForce" json:"force"`
	ForceOEM  int32          `idl:"name:bForceOEM" json:"force_oem"`
	FullClean int32          `idl:"name:bFullClean" json:"full_clean"`
	Async     *vds.Async     `idl:"name:ppAsync" json:"async"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CleanOperation) OpNum() int { return 12 }

func (o *xxx_CleanOperation) OpName() string { return "/IVdsAdvancedDisk/v0/Clean" }

func (o *xxx_CleanOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// bForceOEM {in} (1:(int32))
	{
		if err := w.WriteData(o.ForceOEM); err != nil {
			return err
		}
	}
	// bFullClean {in} (1:(int32))
	{
		if err := w.WriteData(o.FullClean); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// bForceOEM {in} (1:(int32))
	{
		if err := w.ReadData(&o.ForceOEM); err != nil {
			return err
		}
	}
	// bFullClean {in} (1:(int32))
	{
		if err := w.ReadData(&o.FullClean); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CleanOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CleanOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CleanRequest structure represents the Clean operation request
type CleanRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// bForce: A Boolean value that indicates whether the cleaning operation will be forced.
	// If set, the method attempts to clean the disk, even if data volumes or ESP partitions
	// are present.
	Force int32 `idl:"name:bForce" json:"force"`
	// bForceOEM: A Boolean value that indicates whether the cleaning operation of an OEM
	// partition will be forced. If the disk contains an OEM partition, but bForceOEM is
	// not set, the operation SHOULD fail. If the value is set, the method attempts to clean
	// the disk, even if OEM partitions are present.<99>
	ForceOEM int32 `idl:"name:bForceOEM" json:"force_oem"`
	// bFullClean: A Boolean value specifying whether the cleaning operation removes all
	// the data from the disk.
	FullClean int32 `idl:"name:bFullClean" json:"full_clean"`
}

func (o *CleanRequest) xxx_ToOp(ctx context.Context, op *xxx_CleanOperation) *xxx_CleanOperation {
	if op == nil {
		op = &xxx_CleanOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Force = o.Force
	op.ForceOEM = o.ForceOEM
	op.FullClean = o.FullClean
	return op
}

func (o *CleanRequest) xxx_FromOp(ctx context.Context, op *xxx_CleanOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Force = op.Force
	o.ForceOEM = op.ForceOEM
	o.FullClean = op.FullClean
}
func (o *CleanRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CleanRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CleanOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CleanResponse structure represents the Clean operation response
type CleanResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync interface that, if the operation is successfully
	// completed, receives the IVdsAsync interface to monitor and control this operation.
	// Callers MUST release the interface when they are done with it.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The Clean return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CleanResponse) xxx_ToOp(ctx context.Context, op *xxx_CleanOperation) *xxx_CleanOperation {
	if op == nil {
		op = &xxx_CleanOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Async = o.Async
	op.Return = o.Return
	return op
}

func (o *CleanResponse) xxx_FromOp(ctx context.Context, op *xxx_CleanOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *CleanResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CleanResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CleanOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
