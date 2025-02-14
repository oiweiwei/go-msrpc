package ivdsserviceuninstalldisk

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
	// IVdsServiceUninstallDisk interface identifier b6b22da8-f903-4be7-b492-c09d875ac9da
	ServiceUninstallDiskIID = &dcom.IID{Data1: 0xb6b22da8, Data2: 0xf903, Data3: 0x4be7, Data4: []byte{0xb4, 0x92, 0xc0, 0x9d, 0x87, 0x5a, 0xc9, 0xda}}
	// Syntax UUID
	ServiceUninstallDiskSyntaxUUID = &uuid.UUID{TimeLow: 0xb6b22da8, TimeMid: 0xf903, TimeHiAndVersion: 0x4be7, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0x92, Node: [6]uint8{0xc0, 0x9d, 0x87, 0x5a, 0xc9, 0xda}}
	// Syntax ID
	ServiceUninstallDiskSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ServiceUninstallDiskSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsServiceUninstallDisk interface.
type ServiceUninstallDiskClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The GetDiskIdFromLunInfo method retrieves the VDS object ID of a disk that corresponds
	// to a specified LUN information structure.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetDiskIDFromLUNInfo(context.Context, *GetDiskIDFromLUNInfoRequest, ...dcerpc.CallOption) (*GetDiskIDFromLUNInfoResponse, error)

	// The UninstallDisks method uninstalls a specific set of disks when it is given a list
	// of the VDS object IDs for the disks. All volumes that are contained wholly or partially
	// on the disks are also uninstalled, and the obsolete mount points are removed.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	UninstallDisks(context.Context, *UninstallDisksRequest, ...dcerpc.CallOption) (*UninstallDisksResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ServiceUninstallDiskClient
}

type xxx_DefaultServiceUninstallDiskClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultServiceUninstallDiskClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultServiceUninstallDiskClient) GetDiskIDFromLUNInfo(ctx context.Context, in *GetDiskIDFromLUNInfoRequest, opts ...dcerpc.CallOption) (*GetDiskIDFromLUNInfoResponse, error) {
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
	out := &GetDiskIDFromLUNInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceUninstallDiskClient) UninstallDisks(ctx context.Context, in *UninstallDisksRequest, opts ...dcerpc.CallOption) (*UninstallDisksResponse, error) {
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
	out := &UninstallDisksResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultServiceUninstallDiskClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultServiceUninstallDiskClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultServiceUninstallDiskClient) IPID(ctx context.Context, ipid *dcom.IPID) ServiceUninstallDiskClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultServiceUninstallDiskClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewServiceUninstallDiskClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ServiceUninstallDiskClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ServiceUninstallDiskSyntaxV0_0))...)
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
	return &xxx_DefaultServiceUninstallDiskClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetDiskIDFromLUNInfoOperation structure represents the GetDiskIdFromLunInfo operation
type xxx_GetDiskIDFromLUNInfoOperation struct {
	This    *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That    *dcom.ORPCThat      `idl:"name:That" json:"that"`
	LUNInfo *vds.LUNInformation `idl:"name:pLunInfo" json:"lun_info"`
	DiskID  *vds.ObjectID       `idl:"name:pDiskId" json:"disk_id"`
	Return  int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDiskIDFromLUNInfoOperation) OpNum() int { return 3 }

func (o *xxx_GetDiskIDFromLUNInfoOperation) OpName() string {
	return "/IVdsServiceUninstallDisk/v0/GetDiskIdFromLunInfo"
}

func (o *xxx_GetDiskIDFromLUNInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDiskIDFromLUNInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pLunInfo {in} (1:{pointer=ref}*(1))(2:{alias=VDS_LUN_INFORMATION}(struct))
	{
		if o.LUNInfo != nil {
			if err := o.LUNInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.LUNInformation{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDiskIDFromLUNInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pLunInfo {in} (1:{pointer=ref}*(1))(2:{alias=VDS_LUN_INFORMATION}(struct))
	{
		if o.LUNInfo == nil {
			o.LUNInfo = &vds.LUNInformation{}
		}
		if err := o.LUNInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDiskIDFromLUNInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDiskIDFromLUNInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pDiskId {out} (1:{pointer=ref}*(1))(2:{alias=VDS_OBJECT_ID, names=GUID}(struct))
	{
		if o.DiskID != nil {
			if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&vds.ObjectID{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_GetDiskIDFromLUNInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pDiskId {out} (1:{pointer=ref}*(1))(2:{alias=VDS_OBJECT_ID, names=GUID}(struct))
	{
		if o.DiskID == nil {
			o.DiskID = &vds.ObjectID{}
		}
		if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
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

// GetDiskIDFromLUNInfoRequest structure represents the GetDiskIdFromLunInfo operation request
type GetDiskIDFromLUNInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pLunInfo: A pointer to a VDS_LUN_INFORMATION structure that stores the disk's LUN
	// information.
	LUNInfo *vds.LUNInformation `idl:"name:pLunInfo" json:"lun_info"`
}

func (o *GetDiskIDFromLUNInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDiskIDFromLUNInfoOperation) *xxx_GetDiskIDFromLUNInfoOperation {
	if op == nil {
		op = &xxx_GetDiskIDFromLUNInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.LUNInfo = o.LUNInfo
	return op
}

func (o *GetDiskIDFromLUNInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDiskIDFromLUNInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.LUNInfo = op.LUNInfo
}
func (o *GetDiskIDFromLUNInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDiskIDFromLUNInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDiskIDFromLUNInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDiskIDFromLUNInfoResponse structure represents the GetDiskIdFromLunInfo operation response
type GetDiskIDFromLUNInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pDiskId: A pointer to a VDS_OBJECT_ID structure that, if the operation is successfully
	// completed, receives the VDS object ID of the disk object that corresponds to the
	// LUN information that pLunInfo specifies.
	DiskID *vds.ObjectID `idl:"name:pDiskId" json:"disk_id"`
	// Return: The GetDiskIdFromLunInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetDiskIDFromLUNInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDiskIDFromLUNInfoOperation) *xxx_GetDiskIDFromLUNInfoOperation {
	if op == nil {
		op = &xxx_GetDiskIDFromLUNInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.DiskID = o.DiskID
	op.Return = o.Return
	return op
}

func (o *GetDiskIDFromLUNInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDiskIDFromLUNInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DiskID = op.DiskID
	o.Return = op.Return
}
func (o *GetDiskIDFromLUNInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDiskIDFromLUNInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDiskIDFromLUNInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UninstallDisksOperation structure represents the UninstallDisks operation
type xxx_UninstallDisksOperation struct {
	This        *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat  `idl:"name:That" json:"that"`
	DiskIDArray []*vds.ObjectID `idl:"name:pDiskIdArray;size_is:(ulCount)" json:"disk_id_array"`
	Count       uint32          `idl:"name:ulCount" json:"count"`
	Force       bool            `idl:"name:bForce" json:"force"`
	Reboot      bool            `idl:"name:pbReboot" json:"reboot"`
	Results     []int32         `idl:"name:pResults;size_is:(ulCount)" json:"results"`
	Return      int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_UninstallDisksOperation) OpNum() int { return 4 }

func (o *xxx_UninstallDisksOperation) OpName() string {
	return "/IVdsServiceUninstallDisk/v0/UninstallDisks"
}

func (o *xxx_UninstallDisksOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.DiskIDArray != nil && o.Count == 0 {
		o.Count = uint32(len(o.DiskIDArray))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UninstallDisksOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// pDiskIdArray {in} (1:{pointer=ref}*(1))(2:{alias=VDS_OBJECT_ID, names=GUID}[dim:0,size_is=ulCount](struct))
	{
		dimSize1 := uint64(o.Count)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.DiskIDArray {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.DiskIDArray[i1] != nil {
				if err := o.DiskIDArray[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&vds.ObjectID{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.DiskIDArray); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&vds.ObjectID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ulCount {in} (1:(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	// bForce {in} (1:(bool))
	{
		if err := w.WriteData(o.Force); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UninstallDisksOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// pDiskIdArray {in} (1:{pointer=ref}*(1))(2:{alias=VDS_OBJECT_ID, names=GUID}[dim:0,size_is=ulCount](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.DiskIDArray", sizeInfo[0])
		}
		o.DiskIDArray = make([]*vds.ObjectID, sizeInfo[0])
		for i1 := range o.DiskIDArray {
			i1 := i1
			if o.DiskIDArray[i1] == nil {
				o.DiskIDArray[i1] = &vds.ObjectID{}
			}
			if err := o.DiskIDArray[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ulCount {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	// bForce {in} (1:(bool))
	{
		if err := w.ReadData(&o.Force); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UninstallDisksOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UninstallDisksOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pbReboot {out} (1:{pointer=ref}*(1)(bool))
	{
		if err := w.WriteData(o.Reboot); err != nil {
			return err
		}
	}
	// pResults {out} (1:{pointer=ref}*(1))(2:{alias=HRESULT, names=LONG}[dim:0,size_is=ulCount](int32))
	{
		dimSize1 := uint64(o.Count)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Results {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Results[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Results); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(int32(0)); err != nil {
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

func (o *xxx_UninstallDisksOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pbReboot {out} (1:{pointer=ref}*(1)(bool))
	{
		if err := w.ReadData(&o.Reboot); err != nil {
			return err
		}
	}
	// pResults {out} (1:{pointer=ref}*(1))(2:{alias=HRESULT, names=LONG}[dim:0,size_is=ulCount](int32))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Results", sizeInfo[0])
		}
		o.Results = make([]int32, sizeInfo[0])
		for i1 := range o.Results {
			i1 := i1
			if err := w.ReadData(&o.Results[i1]); err != nil {
				return err
			}
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

// UninstallDisksRequest structure represents the UninstallDisks operation request
type UninstallDisksRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// pDiskIdArray: A pointer to an array of VDS_OBJECT_ID structures that store the VDS
	// object IDs of the disks to be uninstalled.
	DiskIDArray []*vds.ObjectID `idl:"name:pDiskIdArray;size_is:(ulCount)" json:"disk_id_array"`
	// ulCount: The number of disks that are specified in pDiskIdArray.
	Count uint32 `idl:"name:ulCount" json:"count"`
	// bForce: A Boolean that determines whether the volume dismount is forced.
	Force bool `idl:"name:bForce" json:"force"`
}

func (o *UninstallDisksRequest) xxx_ToOp(ctx context.Context, op *xxx_UninstallDisksOperation) *xxx_UninstallDisksOperation {
	if op == nil {
		op = &xxx_UninstallDisksOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.DiskIDArray = o.DiskIDArray
	op.Count = o.Count
	op.Force = o.Force
	return op
}

func (o *UninstallDisksRequest) xxx_FromOp(ctx context.Context, op *xxx_UninstallDisksOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskIDArray = op.DiskIDArray
	o.Count = op.Count
	o.Force = op.Force
}
func (o *UninstallDisksRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *UninstallDisksRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UninstallDisksOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UninstallDisksResponse structure represents the UninstallDisks operation response
type UninstallDisksResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbReboot: A pointer to a Boolean that, if the operation is successfully completed,
	// indicates whether the user is required to reboot the remote machine to complete the
	// uninstall process.
	Reboot bool `idl:"name:pbReboot" json:"reboot"`
	// pResults: A pointer to an array of HRESULT values that, if the operation is successfully
	// completed, receives an HRESULT for each disk uninstall request. There MUST be one
	// HRESULT value in the array for each disk in pDiskIdArray. If any disk fails to uninstall
	// properly, the error code for that failure is received in the corresponding entry
	// in pResults.
	Results []int32 `idl:"name:pResults;size_is:(ulCount)" json:"results"`
	// Return: The UninstallDisks return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *UninstallDisksResponse) xxx_ToOp(ctx context.Context, op *xxx_UninstallDisksOperation) *xxx_UninstallDisksOperation {
	if op == nil {
		op = &xxx_UninstallDisksOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Reboot = o.Reboot
	op.Results = o.Results
	op.Return = o.Return
	return op
}

func (o *UninstallDisksResponse) xxx_FromOp(ctx context.Context, op *xxx_UninstallDisksOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Reboot = op.Reboot
	o.Results = op.Results
	o.Return = op.Return
}
func (o *UninstallDisksResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *UninstallDisksResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UninstallDisksOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
