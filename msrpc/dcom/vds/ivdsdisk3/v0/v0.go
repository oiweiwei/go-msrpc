package ivdsdisk3

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
	// IVdsDisk3 interface identifier 8f4b2f5d-ec15-4357-992f-473ef10975b9
	Disk3IID = &dcom.IID{Data1: 0x8f4b2f5d, Data2: 0xec15, Data3: 0x4357, Data4: []byte{0x99, 0x2f, 0x47, 0x3e, 0xf1, 0x09, 0x75, 0xb9}}
	// Syntax UUID
	Disk3SyntaxUUID = &uuid.UUID{TimeLow: 0x8f4b2f5d, TimeMid: 0xec15, TimeHiAndVersion: 0x4357, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0x2f, Node: [6]uint8{0x47, 0x3e, 0xf1, 0x9, 0x75, 0xb9}}
	// Syntax ID
	Disk3SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: Disk3SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsDisk3 interface.
type Disk3Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// GetProperties2 operation.
	GetProperties2(context.Context, *GetProperties2Request, ...dcerpc.CallOption) (*GetProperties2Response, error)

	// The QueryFreeExtents method enumerates a disk's free extents.Returns all free extents
	// on the disk and aligns them to the alignment value supplied in the ulAlign parameter.
	// If there is no alignment value supplied, QueryFreeExtents aligns the free extents
	// based on the default alignment values.
	//
	// Return Values: QueryFreeExtents MUST return zero to indicate success, or an implementation-specific,
	// nonzero error code to indicate failure.
	//
	// Free extents are not returned for CD/DVD, or super floppy devices.
	//
	// If the disk has no partition format (it is not formatted as either MBR or GPT), then
	// this method MUST return VDS_E_DISK_NOT_INITIALIZED.
	QueryFreeExtents(context.Context, *QueryFreeExtentsRequest, ...dcerpc.CallOption) (*QueryFreeExtentsResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) Disk3Client
}

type xxx_DefaultDisk3Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultDisk3Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultDisk3Client) GetProperties2(ctx context.Context, in *GetProperties2Request, opts ...dcerpc.CallOption) (*GetProperties2Response, error) {
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
	out := &GetProperties2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDisk3Client) QueryFreeExtents(ctx context.Context, in *QueryFreeExtentsRequest, opts ...dcerpc.CallOption) (*QueryFreeExtentsResponse, error) {
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
	out := &QueryFreeExtentsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDisk3Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultDisk3Client) IPID(ctx context.Context, ipid *dcom.IPID) Disk3Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultDisk3Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewDisk3Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Disk3Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Disk3SyntaxV0_0))...)
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
	return &xxx_DefaultDisk3Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetProperties2Operation structure represents the GetProperties2 operation
type xxx_GetProperties2Operation struct {
	This           *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat     `idl:"name:That" json:"that"`
	DiskProperties *vds.DiskProperty2 `idl:"name:pDiskProperties" json:"disk_properties"`
	Return         int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_GetProperties2Operation) OpNum() int { return 3 }

func (o *xxx_GetProperties2Operation) OpName() string { return "/IVdsDisk3/v0/GetProperties2" }

func (o *xxx_GetProperties2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProperties2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetProperties2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetProperties2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProperties2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pDiskProperties {out} (1:{pointer=ref}*(1))(2:{alias=VDS_DISK_PROP2}(struct))
	{
		if o.DiskProperties != nil {
			if err := o.DiskProperties.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.DiskProperty2{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetProperties2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pDiskProperties {out} (1:{pointer=ref}*(1))(2:{alias=VDS_DISK_PROP2}(struct))
	{
		if o.DiskProperties == nil {
			o.DiskProperties = &vds.DiskProperty2{}
		}
		if err := o.DiskProperties.UnmarshalNDR(ctx, w); err != nil {
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

// GetProperties2Request structure represents the GetProperties2 operation request
type GetProperties2Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetProperties2Request) xxx_ToOp(ctx context.Context) *xxx_GetProperties2Operation {
	if o == nil {
		return &xxx_GetProperties2Operation{}
	}
	return &xxx_GetProperties2Operation{
		This: o.This,
	}
}

func (o *GetProperties2Request) xxx_FromOp(ctx context.Context, op *xxx_GetProperties2Operation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetProperties2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetProperties2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetProperties2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetProperties2Response structure represents the GetProperties2 operation response
type GetProperties2Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That           *dcom.ORPCThat     `idl:"name:That" json:"that"`
	DiskProperties *vds.DiskProperty2 `idl:"name:pDiskProperties" json:"disk_properties"`
	// Return: The GetProperties2 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetProperties2Response) xxx_ToOp(ctx context.Context) *xxx_GetProperties2Operation {
	if o == nil {
		return &xxx_GetProperties2Operation{}
	}
	return &xxx_GetProperties2Operation{
		That:           o.That,
		DiskProperties: o.DiskProperties,
		Return:         o.Return,
	}
}

func (o *GetProperties2Response) xxx_FromOp(ctx context.Context, op *xxx_GetProperties2Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DiskProperties = op.DiskProperties
	o.Return = op.Return
}
func (o *GetProperties2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetProperties2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetProperties2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryFreeExtentsOperation structure represents the QueryFreeExtents operation
type xxx_QueryFreeExtentsOperation struct {
	This                *dcom.ORPCThis        `idl:"name:This" json:"this"`
	That                *dcom.ORPCThat        `idl:"name:That" json:"that"`
	Align               uint32                `idl:"name:ulAlign" json:"align"`
	FreeExtentArray     []*vds.DiskFreeExtent `idl:"name:ppFreeExtentArray;size_is:(, plNumberOfFreeExtents)" json:"free_extent_array"`
	NumberOfFreeExtents int32                 `idl:"name:plNumberOfFreeExtents" json:"number_of_free_extents"`
	Return              int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryFreeExtentsOperation) OpNum() int { return 4 }

func (o *xxx_QueryFreeExtentsOperation) OpName() string { return "/IVdsDisk3/v0/QueryFreeExtents" }

func (o *xxx_QueryFreeExtentsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryFreeExtentsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// ulAlign {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Align); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryFreeExtentsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// ulAlign {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Align); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryFreeExtentsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.FreeExtentArray != nil && o.NumberOfFreeExtents == 0 {
		o.NumberOfFreeExtents = int32(len(o.FreeExtentArray))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryFreeExtentsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppFreeExtentArray {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VDS_DISK_FREE_EXTENT}[dim:0,size_is=plNumberOfFreeExtents](struct))
	{
		if o.FreeExtentArray != nil || o.NumberOfFreeExtents > 0 {
			_ptr_ppFreeExtentArray := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NumberOfFreeExtents)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.FreeExtentArray {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.FreeExtentArray[i1] != nil {
						if err := o.FreeExtentArray[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&vds.DiskFreeExtent{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.FreeExtentArray); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&vds.DiskFreeExtent{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FreeExtentArray, _ptr_ppFreeExtentArray); err != nil {
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
	// plNumberOfFreeExtents {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.NumberOfFreeExtents); err != nil {
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

func (o *xxx_QueryFreeExtentsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppFreeExtentArray {out} (1:{pointer=ref}*(2)*(1))(2:{alias=VDS_DISK_FREE_EXTENT}[dim:0,size_is=plNumberOfFreeExtents](struct))
	{
		_ptr_ppFreeExtentArray := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.FreeExtentArray", sizeInfo[0])
			}
			o.FreeExtentArray = make([]*vds.DiskFreeExtent, sizeInfo[0])
			for i1 := range o.FreeExtentArray {
				i1 := i1
				if o.FreeExtentArray[i1] == nil {
					o.FreeExtentArray[i1] = &vds.DiskFreeExtent{}
				}
				if err := o.FreeExtentArray[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppFreeExtentArray := func(ptr interface{}) { o.FreeExtentArray = *ptr.(*[]*vds.DiskFreeExtent) }
		if err := w.ReadPointer(&o.FreeExtentArray, _s_ppFreeExtentArray, _ptr_ppFreeExtentArray); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// plNumberOfFreeExtents {out} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.NumberOfFreeExtents); err != nil {
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

// QueryFreeExtentsRequest structure represents the QueryFreeExtents operation request
type QueryFreeExtentsRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ulAlign: The alignment value. If ulAlign is 0, the default alignment value is used.
	Align uint32 `idl:"name:ulAlign" json:"align"`
}

func (o *QueryFreeExtentsRequest) xxx_ToOp(ctx context.Context) *xxx_QueryFreeExtentsOperation {
	if o == nil {
		return &xxx_QueryFreeExtentsOperation{}
	}
	return &xxx_QueryFreeExtentsOperation{
		This:  o.This,
		Align: o.Align,
	}
}

func (o *QueryFreeExtentsRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryFreeExtentsOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Align = op.Align
}
func (o *QueryFreeExtentsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryFreeExtentsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryFreeExtentsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryFreeExtentsResponse structure represents the QueryFreeExtents operation response
type QueryFreeExtentsResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppFreeExtentArray: Pointer to an array of VDS_DISK_FREE_EXTENT structures that, if
	// the operation is successful, receives the array of disk extent structures.
	FreeExtentArray []*vds.DiskFreeExtent `idl:"name:ppFreeExtentArray;size_is:(, plNumberOfFreeExtents)" json:"free_extent_array"`
	// plNumberOfFreeExtents: Pointer to a variable that, if the operation is successfully
	// completed, receives the total number of elements in ppFreeExtentArray.
	NumberOfFreeExtents int32 `idl:"name:plNumberOfFreeExtents" json:"number_of_free_extents"`
	// Return: The QueryFreeExtents return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryFreeExtentsResponse) xxx_ToOp(ctx context.Context) *xxx_QueryFreeExtentsOperation {
	if o == nil {
		return &xxx_QueryFreeExtentsOperation{}
	}
	return &xxx_QueryFreeExtentsOperation{
		That:                o.That,
		FreeExtentArray:     o.FreeExtentArray,
		NumberOfFreeExtents: o.NumberOfFreeExtents,
		Return:              o.Return,
	}
}

func (o *QueryFreeExtentsResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryFreeExtentsOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.FreeExtentArray = op.FreeExtentArray
	o.NumberOfFreeExtents = op.NumberOfFreeExtents
	o.Return = op.Return
}
func (o *QueryFreeExtentsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryFreeExtentsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryFreeExtentsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
