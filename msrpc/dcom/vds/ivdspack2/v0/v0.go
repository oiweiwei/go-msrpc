package ivdspack2

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
	// IVdsPack2 interface identifier 13b50bff-290a-47dd-8558-b7c58db1a71a
	Pack2IID = &dcom.IID{Data1: 0x13b50bff, Data2: 0x290a, Data3: 0x47dd, Data4: []byte{0x85, 0x58, 0xb7, 0xc5, 0x8d, 0xb1, 0xa7, 0x1a}}
	// Syntax UUID
	Pack2SyntaxUUID = &uuid.UUID{TimeLow: 0x13b50bff, TimeMid: 0x290a, TimeHiAndVersion: 0x47dd, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0x58, Node: [6]uint8{0xb7, 0xc5, 0x8d, 0xb1, 0xa7, 0x1a}}
	// Syntax ID
	Pack2SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: Pack2SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsPack2 interface.
type Pack2Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The CreateVolume2 method creates a volume in a disk pack with an optional alignment
	// parameter.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	//
	// IVdsPack2::CreateVolume2 has the same sequencing rules as IVdsPack::CreateVolume
	// (Opnum 7), as specified in section 3.4.5.2.19.5.
	CreateVolume2(context.Context, *CreateVolume2Request, ...dcerpc.CallOption) (*CreateVolume2Response, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Pack2Client
}

type xxx_DefaultPack2Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultPack2Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultPack2Client) CreateVolume2(ctx context.Context, in *CreateVolume2Request, opts ...dcerpc.CallOption) (*CreateVolume2Response, error) {
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
	out := &CreateVolume2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultPack2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultPack2Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultPack2Client) IPID(ctx context.Context, ipid *dcom.IPID) Pack2Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultPack2Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewPack2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Pack2Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Pack2SyntaxV0_0))...)
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
	return &xxx_DefaultPack2Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_CreateVolume2Operation structure represents the CreateVolume2 operation
type xxx_CreateVolume2Operation struct {
	This           *dcom.ORPCThis   `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat   `idl:"name:That" json:"that"`
	Type           vds.VolumeType   `idl:"name:type" json:"type"`
	InputDiskArray []*vds.InputDisk `idl:"name:pInputDiskArray;size_is:(lNumberOfDisks)" json:"input_disk_array"`
	NumberOfDisks  int32            `idl:"name:lNumberOfDisks" json:"number_of_disks"`
	StripeSize     uint32           `idl:"name:ulStripeSize" json:"stripe_size"`
	Align          uint32           `idl:"name:ulAlign" json:"align"`
	Async          *vds.Async       `idl:"name:ppAsync" json:"async"`
	Return         int32            `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateVolume2Operation) OpNum() int { return 3 }

func (o *xxx_CreateVolume2Operation) OpName() string { return "/IVdsPack2/v0/CreateVolume2" }

func (o *xxx_CreateVolume2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
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

func (o *xxx_CreateVolume2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// type {in} (1:{alias=VDS_VOLUME_TYPE}(enum))
	{
		if err := w.WriteEnum(uint16(o.Type)); err != nil {
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
	// ulStripeSize {in} (1:(uint32))
	{
		if err := w.WriteData(o.StripeSize); err != nil {
			return err
		}
	}
	// ulAlign {in} (1:(uint32))
	{
		if err := w.WriteData(o.Align); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVolume2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// type {in} (1:{alias=VDS_VOLUME_TYPE}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Type)); err != nil {
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
	// ulStripeSize {in} (1:(uint32))
	{
		if err := w.ReadData(&o.StripeSize); err != nil {
			return err
		}
	}
	// ulAlign {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Align); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVolume2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateVolume2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CreateVolume2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CreateVolume2Request structure represents the CreateVolume2 operation request
type CreateVolume2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// type: A value from the VDS_VOLUME_TYPE enumeration that indicates the type of volume
	// to create.
	Type vds.VolumeType `idl:"name:type" json:"type"`
	// pInputDiskArray: An array of VDS_INPUT_DISK structures that indicate the disks on
	// which to create the volume.<86>
	InputDiskArray []*vds.InputDisk `idl:"name:pInputDiskArray;size_is:(lNumberOfDisks)" json:"input_disk_array"`
	// lNumberOfDisks: The number of elements in pInputDiskArray.
	NumberOfDisks int32 `idl:"name:lNumberOfDisks" json:"number_of_disks"`
	// ulStripeSize: The stripe size, in bytes, of the new volume.<87>
	StripeSize uint32 `idl:"name:ulStripeSize" json:"stripe_size"`
	// ulAlign: The number of bytes for the volume alignment. If zero is specified, the
	// server determines the alignment value based on the size of the disk on which the
	// volume is created.<88>
	Align uint32 `idl:"name:ulAlign" json:"align"`
}

func (o *CreateVolume2Request) xxx_ToOp(ctx context.Context, op *xxx_CreateVolume2Operation) *xxx_CreateVolume2Operation {
	if op == nil {
		op = &xxx_CreateVolume2Operation{}
	}
	if o == nil {
		return op
	}
	o.This = op.This
	o.Type = op.Type
	o.InputDiskArray = op.InputDiskArray
	o.NumberOfDisks = op.NumberOfDisks
	o.StripeSize = op.StripeSize
	o.Align = op.Align
	return op
}

func (o *CreateVolume2Request) xxx_FromOp(ctx context.Context, op *xxx_CreateVolume2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Type = op.Type
	o.InputDiskArray = op.InputDiskArray
	o.NumberOfDisks = op.NumberOfDisks
	o.StripeSize = op.StripeSize
	o.Align = op.Align
}
func (o *CreateVolume2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CreateVolume2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateVolume2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateVolume2Response structure represents the CreateVolume2 operation response
type CreateVolume2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppAsync: A pointer to an IVdsAsync interface that, if the operation is successfully
	// completed, receives the IVdsAsync interface to monitor and control this operation.
	// Callers MUST release the interface when they are done with it. If the IVdsAsync::Wait
	// (Opnum 4) method is called on the interface, the interfaces returned in the VDS_ASYNC_OUTPUT
	// structure MUST be released as well. For more information on handling asynchronous
	// tasks, see section 3.4.5.1.9.
	Async *vds.Async `idl:"name:ppAsync" json:"async"`
	// Return: The CreateVolume2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateVolume2Response) xxx_ToOp(ctx context.Context, op *xxx_CreateVolume2Operation) *xxx_CreateVolume2Operation {
	if op == nil {
		op = &xxx_CreateVolume2Operation{}
	}
	if o == nil {
		return op
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
	return op
}

func (o *CreateVolume2Response) xxx_FromOp(ctx context.Context, op *xxx_CreateVolume2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Async = op.Async
	o.Return = op.Return
}
func (o *CreateVolume2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CreateVolume2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateVolume2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
