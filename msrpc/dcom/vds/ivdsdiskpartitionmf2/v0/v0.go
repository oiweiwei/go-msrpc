package ivdsdiskpartitionmf2

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
	// IVdsDiskPartitionMF2 interface identifier 9cbe50ca-f2d2-4bf4-ace1-96896b729625
	DiskPartitionMF2IID = &dcom.IID{Data1: 0x9cbe50ca, Data2: 0xf2d2, Data3: 0x4bf4, Data4: []byte{0xac, 0xe1, 0x96, 0x89, 0x6b, 0x72, 0x96, 0x25}}
	// Syntax UUID
	DiskPartitionMF2SyntaxUUID = &uuid.UUID{TimeLow: 0x9cbe50ca, TimeMid: 0xf2d2, TimeHiAndVersion: 0x4bf4, ClockSeqHiAndReserved: 0xac, ClockSeqLow: 0xe1, Node: [6]uint8{0x96, 0x89, 0x6b, 0x72, 0x96, 0x25}}
	// Syntax ID
	DiskPartitionMF2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: DiskPartitionMF2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsDiskPartitionMF2 interface.
type DiskPartitionMF2Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The FormatPartitionEx2 method formats an existing OEM, ESP, or unknown partition.
	// This method is only supported on OEM, ESP, recovery, and unknown partitions.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	FormatPartitionEx2(context.Context, *FormatPartitionEx2Request, ...dcerpc.CallOption) (*FormatPartitionEx2Response, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) DiskPartitionMF2Client
}

type xxx_DefaultDiskPartitionMF2Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultDiskPartitionMF2Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultDiskPartitionMF2Client) FormatPartitionEx2(ctx context.Context, in *FormatPartitionEx2Request, opts ...dcerpc.CallOption) (*FormatPartitionEx2Response, error) {
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
	out := &FormatPartitionEx2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDiskPartitionMF2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultDiskPartitionMF2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultDiskPartitionMF2Client) IPID(ctx context.Context, ipid *dcom.IPID) DiskPartitionMF2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultDiskPartitionMF2Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewDiskPartitionMF2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (DiskPartitionMF2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(DiskPartitionMF2SyntaxV0_0))...)
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
	return &xxx_DefaultDiskPartitionMF2Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_FormatPartitionEx2Operation structure represents the FormatPartitionEx2 operation
type xxx_FormatPartitionEx2Operation struct {
	This                      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                      *dcom.ORPCThat `idl:"name:That" json:"that"`
	Offset                    uint64         `idl:"name:ullOffset" json:"offset"`
	FileSystemTypeName        string         `idl:"name:pwszFileSystemTypeName;string;pointer:unique" json:"file_system_type_name"`
	FileSystemRevision        uint16         `idl:"name:usFileSystemRevision" json:"file_system_revision"`
	DesiredUnitAllocationSize uint32         `idl:"name:ulDesiredUnitAllocationSize" json:"desired_unit_allocation_size"`
	Label                     string         `idl:"name:pwszLabel;string;pointer:unique" json:"label"`
	Options                   uint32         `idl:"name:Options" json:"options"`
	Async                     *vds.Async     `idl:"name:ppAsync" json:"async"`
	Return                    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_FormatPartitionEx2Operation) OpNum() int { return 3 }

func (o *xxx_FormatPartitionEx2Operation) OpName() string {
	return "/IVdsDiskPartitionMF2/v0/FormatPartitionEx2"
}

func (o *xxx_FormatPartitionEx2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatPartitionEx2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pwszFileSystemTypeName {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
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
	// pwszLabel {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
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
	// Options {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatPartitionEx2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pwszFileSystemTypeName {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
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
	// pwszLabel {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
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
	// Options {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatPartitionEx2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FormatPartitionEx2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_FormatPartitionEx2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// FormatPartitionEx2Request structure represents the FormatPartitionEx2 operation request
type FormatPartitionEx2Request struct {
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
	// the minor revision (for example, 0x0250 represents version 2.50).
	FileSystemRevision uint16 `idl:"name:usFileSystemRevision" json:"file_system_revision"`
	// ulDesiredUnitAllocationSize: The size, in bytes, of the allocation unit for the file
	// system. The value MUST be a power of 2. If the value is 0, a default allocation unit
	// determined by the file system type is used. The allocation unit range is file system-dependent.
	DesiredUnitAllocationSize uint32 `idl:"name:ulDesiredUnitAllocationSize" json:"desired_unit_allocation_size"`
	// pwszLabel: The null-terminated Unicode string to assign to the new file system. The
	// maximum label size is file system-dependent.
	Label string `idl:"name:pwszLabel;string;pointer:unique" json:"label"`
	// Options: The combination of any values, by using a bitwise OR operator, that are
	// defined in the VDS_FORMAT_OPTION_FLAGS enumeration.
	Options uint32 `idl:"name:Options" json:"options"`
}

func (o *FormatPartitionEx2Request) xxx_ToOp(ctx context.Context, op *xxx_FormatPartitionEx2Operation) *xxx_FormatPartitionEx2Operation {
	if op == nil {
		op = &xxx_FormatPartitionEx2Operation{}
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
	o.Options = op.Options
	return op
}

func (o *FormatPartitionEx2Request) xxx_FromOp(ctx context.Context, op *xxx_FormatPartitionEx2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Offset = op.Offset
	o.FileSystemTypeName = op.FileSystemTypeName
	o.FileSystemRevision = op.FileSystemRevision
	o.DesiredUnitAllocationSize = op.DesiredUnitAllocationSize
	o.Label = op.Label
	o.Options = op.Options
}
func (o *FormatPartitionEx2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FormatPartitionEx2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FormatPartitionEx2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FormatPartitionEx2Response structure represents the FormatPartitionEx2 operation response
type FormatPartitionEx2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync interface that, if the operation is successfully
	// completed, receives the IVdsAsync interface pointer to monitor and control this operation.
	// Callers MUST release the interface when they are done with it.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The FormatPartitionEx2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FormatPartitionEx2Response) xxx_ToOp(ctx context.Context, op *xxx_FormatPartitionEx2Operation) *xxx_FormatPartitionEx2Operation {
	if op == nil {
		op = &xxx_FormatPartitionEx2Operation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
	return op
}

func (o *FormatPartitionEx2Response) xxx_FromOp(ctx context.Context, op *xxx_FormatPartitionEx2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *FormatPartitionEx2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FormatPartitionEx2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FormatPartitionEx2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
