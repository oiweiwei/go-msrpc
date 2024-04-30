package igettrackingdata

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	comt "github.com/oiweiwei/go-msrpc/msrpc/dcom/comt"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
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
	_ = comt.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/comt"
)

var (
	// IGetTrackingData interface identifier b60040e0-bcf3-11d1-861d-0080c729264d
	GetTrackingDataIID = &dcom.IID{Data1: 0xb60040e0, Data2: 0xbcf3, Data3: 0x11d1, Data4: []byte{0x86, 0x1d, 0x00, 0x80, 0xc7, 0x29, 0x26, 0x4d}}
	// Syntax UUID
	GetTrackingDataSyntaxUUID = &uuid.UUID{TimeLow: 0xb60040e0, TimeMid: 0xbcf3, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x86, ClockSeqLow: 0x1d, Node: [6]uint8{0x0, 0x80, 0xc7, 0x29, 0x26, 0x4d}}
	// Syntax ID
	GetTrackingDataSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: GetTrackingDataSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IGetTrackingData interface.
type GetTrackingDataClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// Opnum3NotUsedOnWire operation.
	// Opnum3NotUsedOnWire

	// A client calls this method to obtain tracking information for instance containers
	// across all conglomerations.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success and a failure
	// result (as specified in [MS-ERREF] section 2.1) on failure.
	GetContainerData(context.Context, *GetContainerDataRequest, ...dcerpc.CallOption) (*GetContainerDataResponse, error)

	// A client calls this method to obtain tracking information for components that have
	// one or more component instances in a given instance container.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success and a failure
	// result (as specified in [MS-ERREF] section 2.1) on failure.
	GetComponentDataByContainer(context.Context, *GetComponentDataByContainerRequest, ...dcerpc.CallOption) (*GetComponentDataByContainerResponse, error)

	// A client calls this method to obtain tracking information for a single component
	// that has component instances in an instance container.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success and a failure
	// result (as specified in [MS-ERREF] section 2.1) on failure.
	GetComponentDataByContainerAndClassID(context.Context, *GetComponentDataByContainerAndClassIDRequest, ...dcerpc.CallOption) (*GetComponentDataByContainerAndClassIDResponse, error)

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) GetTrackingDataClient
}

type xxx_DefaultGetTrackingDataClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultGetTrackingDataClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultGetTrackingDataClient) GetContainerData(ctx context.Context, in *GetContainerDataRequest, opts ...dcerpc.CallOption) (*GetContainerDataResponse, error) {
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
	out := &GetContainerDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultGetTrackingDataClient) GetComponentDataByContainer(ctx context.Context, in *GetComponentDataByContainerRequest, opts ...dcerpc.CallOption) (*GetComponentDataByContainerResponse, error) {
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
	out := &GetComponentDataByContainerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultGetTrackingDataClient) GetComponentDataByContainerAndClassID(ctx context.Context, in *GetComponentDataByContainerAndClassIDRequest, opts ...dcerpc.CallOption) (*GetComponentDataByContainerAndClassIDResponse, error) {
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
	out := &GetComponentDataByContainerAndClassIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultGetTrackingDataClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultGetTrackingDataClient) IPID(ctx context.Context, ipid *dcom.IPID) GetTrackingDataClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultGetTrackingDataClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}
func NewGetTrackingDataClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (GetTrackingDataClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(GetTrackingDataSyntaxV0_0))...)
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
	return &xxx_DefaultGetTrackingDataClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetContainerDataOperation structure represents the GetContainerData operation
type xxx_GetContainerDataOperation struct {
	This          *dcom.ORPCThis        `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat        `idl:"name:That" json:"that"`
	Containers    uint32                `idl:"name:nContainers" json:"containers"`
	ContainerData []*comt.ContainerData `idl:"name:aContainerData;size_is:(, nContainers)" json:"container_data"`
	Return        int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetContainerDataOperation) OpNum() int { return 4 }

func (o *xxx_GetContainerDataOperation) OpName() string {
	return "/IGetTrackingData/v0/GetContainerData"
}

func (o *xxx_GetContainerDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetContainerDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetContainerDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetContainerDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ContainerData != nil && o.Containers == 0 {
		o.Containers = uint32(len(o.ContainerData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetContainerDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// nContainers {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Containers); err != nil {
			return err
		}
	}
	// aContainerData {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ContainerData}[dim:0,size_is=nContainers](struct))
	{
		if o.ContainerData != nil || o.Containers > 0 {
			_ptr_aContainerData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.Containers)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ContainerData {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.ContainerData[i1] != nil {
						if err := o.ContainerData[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&comt.ContainerData{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.ContainerData); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&comt.ContainerData{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ContainerData, _ptr_aContainerData); err != nil {
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

func (o *xxx_GetContainerDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// nContainers {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Containers); err != nil {
			return err
		}
	}
	// aContainerData {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ContainerData}[dim:0,size_is=nContainers](struct))
	{
		_ptr_aContainerData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ContainerData", sizeInfo[0])
			}
			o.ContainerData = make([]*comt.ContainerData, sizeInfo[0])
			for i1 := range o.ContainerData {
				i1 := i1
				if o.ContainerData[i1] == nil {
					o.ContainerData[i1] = &comt.ContainerData{}
				}
				if err := o.ContainerData[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_aContainerData := func(ptr interface{}) { o.ContainerData = *ptr.(*[]*comt.ContainerData) }
		if err := w.ReadPointer(&o.ContainerData, _s_aContainerData, _ptr_aContainerData); err != nil {
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

// GetContainerDataRequest structure represents the GetContainerData operation request
type GetContainerDataRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetContainerDataRequest) xxx_ToOp(ctx context.Context) *xxx_GetContainerDataOperation {
	if o == nil {
		return &xxx_GetContainerDataOperation{}
	}
	return &xxx_GetContainerDataOperation{
		This: o.This,
	}
}

func (o *GetContainerDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetContainerDataOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetContainerDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetContainerDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetContainerDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetContainerDataResponse structure represents the GetContainerData operation response
type GetContainerDataResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// nContainers: A pointer to a variable that, upon successful completion, MUST contain
	// the number of elements in aContainerData.
	Containers uint32 `idl:"name:nContainers" json:"containers"`
	// aContainerData: A pointer to a variable that, upon successful completion, MUST contain
	// an array of zero or more ContainerData (section 2.2.3) structures. An array with
	// zero elements MUST be represented by null.
	ContainerData []*comt.ContainerData `idl:"name:aContainerData;size_is:(, nContainers)" json:"container_data"`
	// Return: The GetContainerData return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetContainerDataResponse) xxx_ToOp(ctx context.Context) *xxx_GetContainerDataOperation {
	if o == nil {
		return &xxx_GetContainerDataOperation{}
	}
	return &xxx_GetContainerDataOperation{
		That:          o.That,
		Containers:    o.Containers,
		ContainerData: o.ContainerData,
		Return:        o.Return,
	}
}

func (o *GetContainerDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetContainerDataOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Containers = op.Containers
	o.ContainerData = op.ContainerData
	o.Return = op.Return
}
func (o *GetContainerDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetContainerDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetContainerDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetComponentDataByContainerOperation structure represents the GetComponentDataByContainer operation
type xxx_GetComponentDataByContainerOperation struct {
	This          *dcom.ORPCThis        `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat        `idl:"name:That" json:"that"`
	ContainerID   uint32                `idl:"name:idContainer" json:"container_id"`
	Components    uint32                `idl:"name:nComponents" json:"components"`
	ComponentData []*comt.ComponentData `idl:"name:aComponentData;size_is:(, nComponents)" json:"component_data"`
	Return        int32                 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetComponentDataByContainerOperation) OpNum() int { return 5 }

func (o *xxx_GetComponentDataByContainerOperation) OpName() string {
	return "/IGetTrackingData/v0/GetComponentDataByContainer"
}

func (o *xxx_GetComponentDataByContainerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetComponentDataByContainerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// idContainer {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ContainerID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetComponentDataByContainerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// idContainer {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ContainerID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetComponentDataByContainerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ComponentData != nil && o.Components == 0 {
		o.Components = uint32(len(o.ComponentData))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetComponentDataByContainerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// nComponents {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Components); err != nil {
			return err
		}
	}
	// aComponentData {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ComponentData}[dim:0,size_is=nComponents](struct))
	{
		if o.ComponentData != nil || o.Components > 0 {
			_ptr_aComponentData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.Components)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.ComponentData {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.ComponentData[i1] != nil {
						if err := o.ComponentData[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&comt.ComponentData{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.ComponentData); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&comt.ComponentData{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ComponentData, _ptr_aComponentData); err != nil {
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

func (o *xxx_GetComponentDataByContainerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// nComponents {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Components); err != nil {
			return err
		}
	}
	// aComponentData {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ComponentData}[dim:0,size_is=nComponents](struct))
	{
		_ptr_aComponentData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.ComponentData", sizeInfo[0])
			}
			o.ComponentData = make([]*comt.ComponentData, sizeInfo[0])
			for i1 := range o.ComponentData {
				i1 := i1
				if o.ComponentData[i1] == nil {
					o.ComponentData[i1] = &comt.ComponentData{}
				}
				if err := o.ComponentData[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_aComponentData := func(ptr interface{}) { o.ComponentData = *ptr.(*[]*comt.ComponentData) }
		if err := w.ReadPointer(&o.ComponentData, _s_aComponentData, _ptr_aComponentData); err != nil {
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

// GetComponentDataByContainerRequest structure represents the GetComponentDataByContainer operation request
type GetComponentDataByContainerRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// idContainer: The container legacy identifier of an instance container.
	ContainerID uint32 `idl:"name:idContainer" json:"container_id"`
}

func (o *GetComponentDataByContainerRequest) xxx_ToOp(ctx context.Context) *xxx_GetComponentDataByContainerOperation {
	if o == nil {
		return &xxx_GetComponentDataByContainerOperation{}
	}
	return &xxx_GetComponentDataByContainerOperation{
		This:        o.This,
		ContainerID: o.ContainerID,
	}
}

func (o *GetComponentDataByContainerRequest) xxx_FromOp(ctx context.Context, op *xxx_GetComponentDataByContainerOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ContainerID = op.ContainerID
}
func (o *GetComponentDataByContainerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetComponentDataByContainerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetComponentDataByContainerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetComponentDataByContainerResponse structure represents the GetComponentDataByContainer operation response
type GetComponentDataByContainerResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// nComponents: A pointer to a variable that, upon successful completion, MUST contain
	// the number of elements in aComponentData.
	Components uint32 `idl:"name:nComponents" json:"components"`
	// aComponentData: A pointer to a variable that, upon successful completion, MUST contain
	// an array of zero or more ComponentData (section 2.2.4) structures. An array with
	// zero elements MUST be represented by null.
	ComponentData []*comt.ComponentData `idl:"name:aComponentData;size_is:(, nComponents)" json:"component_data"`
	// Return: The GetComponentDataByContainer return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetComponentDataByContainerResponse) xxx_ToOp(ctx context.Context) *xxx_GetComponentDataByContainerOperation {
	if o == nil {
		return &xxx_GetComponentDataByContainerOperation{}
	}
	return &xxx_GetComponentDataByContainerOperation{
		That:          o.That,
		Components:    o.Components,
		ComponentData: o.ComponentData,
		Return:        o.Return,
	}
}

func (o *GetComponentDataByContainerResponse) xxx_FromOp(ctx context.Context, op *xxx_GetComponentDataByContainerOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Components = op.Components
	o.ComponentData = op.ComponentData
	o.Return = op.Return
}
func (o *GetComponentDataByContainerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetComponentDataByContainerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetComponentDataByContainerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetComponentDataByContainerAndClassIDOperation structure represents the GetComponentDataByContainerAndCLSID operation
type xxx_GetComponentDataByContainerAndClassIDOperation struct {
	This          *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat      `idl:"name:That" json:"that"`
	ContainerID   uint32              `idl:"name:idContainer" json:"container_id"`
	ClassID       *dtyp.GUID          `idl:"name:clsid" json:"class_id"`
	ComponentData *comt.ComponentData `idl:"name:ppComponentData" json:"component_data"`
	Return        int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_GetComponentDataByContainerAndClassIDOperation) OpNum() int { return 6 }

func (o *xxx_GetComponentDataByContainerAndClassIDOperation) OpName() string {
	return "/IGetTrackingData/v0/GetComponentDataByContainerAndCLSID"
}

func (o *xxx_GetComponentDataByContainerAndClassIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetComponentDataByContainerAndClassIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// idContainer {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ContainerID); err != nil {
			return err
		}
	}
	// clsid {in} (1:{alias=GUID}(struct))
	{
		if o.ClassID != nil {
			if err := o.ClassID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetComponentDataByContainerAndClassIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// idContainer {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ContainerID); err != nil {
			return err
		}
	}
	// clsid {in} (1:{alias=GUID}(struct))
	{
		if o.ClassID == nil {
			o.ClassID = &dtyp.GUID{}
		}
		if err := o.ClassID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetComponentDataByContainerAndClassIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetComponentDataByContainerAndClassIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppComponentData {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ComponentData}(struct))
	{
		if o.ComponentData != nil {
			_ptr_ppComponentData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ComponentData != nil {
					if err := o.ComponentData.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&comt.ComponentData{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ComponentData, _ptr_ppComponentData); err != nil {
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

func (o *xxx_GetComponentDataByContainerAndClassIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppComponentData {out} (1:{pointer=ref}*(2)*(1))(2:{alias=ComponentData}(struct))
	{
		_ptr_ppComponentData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ComponentData == nil {
				o.ComponentData = &comt.ComponentData{}
			}
			if err := o.ComponentData.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppComponentData := func(ptr interface{}) { o.ComponentData = *ptr.(**comt.ComponentData) }
		if err := w.ReadPointer(&o.ComponentData, _s_ppComponentData, _ptr_ppComponentData); err != nil {
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

// GetComponentDataByContainerAndClassIDRequest structure represents the GetComponentDataByContainerAndCLSID operation request
type GetComponentDataByContainerAndClassIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// idContainer: The container legacy identifier of an instance container.
	ContainerID uint32 `idl:"name:idContainer" json:"container_id"`
	// clsid: A pointer to the CLSID of a component.
	ClassID *dtyp.GUID `idl:"name:clsid" json:"class_id"`
}

func (o *GetComponentDataByContainerAndClassIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetComponentDataByContainerAndClassIDOperation {
	if o == nil {
		return &xxx_GetComponentDataByContainerAndClassIDOperation{}
	}
	return &xxx_GetComponentDataByContainerAndClassIDOperation{
		This:        o.This,
		ContainerID: o.ContainerID,
		ClassID:     o.ClassID,
	}
}

func (o *GetComponentDataByContainerAndClassIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetComponentDataByContainerAndClassIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.ContainerID = op.ContainerID
	o.ClassID = op.ClassID
}
func (o *GetComponentDataByContainerAndClassIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetComponentDataByContainerAndClassIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetComponentDataByContainerAndClassIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetComponentDataByContainerAndClassIDResponse structure represents the GetComponentDataByContainerAndCLSID operation response
type GetComponentDataByContainerAndClassIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppComponentData: A pointer to a variable that, upon successful completion, MUST contain
	// a pointer to a single ComponentData (section 2.2.4) structure.
	ComponentData *comt.ComponentData `idl:"name:ppComponentData" json:"component_data"`
	// Return: The GetComponentDataByContainerAndCLSID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetComponentDataByContainerAndClassIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetComponentDataByContainerAndClassIDOperation {
	if o == nil {
		return &xxx_GetComponentDataByContainerAndClassIDOperation{}
	}
	return &xxx_GetComponentDataByContainerAndClassIDOperation{
		That:          o.That,
		ComponentData: o.ComponentData,
		Return:        o.Return,
	}
}

func (o *GetComponentDataByContainerAndClassIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetComponentDataByContainerAndClassIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ComponentData = op.ComponentData
	o.Return = op.Return
}
func (o *GetComponentDataByContainerAndClassIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetComponentDataByContainerAndClassIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetComponentDataByContainerAndClassIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
