package iwbemremoterefresher

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
	wmi "github.com/oiweiwei/go-msrpc/msrpc/dcom/wmi"
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
	_ = wmi.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/wmi"
)

var (
	// IWbemRemoteRefresher interface identifier f1e9c5b2-f59b-11d2-b362-00105a1f8177
	RemoteRefresherIID = &dcom.IID{Data1: 0xf1e9c5b2, Data2: 0xf59b, Data3: 0x11d2, Data4: []byte{0xb3, 0x62, 0x00, 0x10, 0x5a, 0x1f, 0x81, 0x77}}
	// Syntax UUID
	RemoteRefresherSyntaxUUID = &uuid.UUID{TimeLow: 0xf1e9c5b2, TimeMid: 0xf59b, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0x62, Node: [6]uint8{0x0, 0x10, 0x5a, 0x1f, 0x81, 0x77}}
	// Syntax ID
	RemoteRefresherSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: RemoteRefresherSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IWbemRemoteRefresher interface.
type RemoteRefresherClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The IWbemRemoteRefresher::RemoteRefresh method MUST return the updated collection
	// of CIM instances and enumerations previously configured by the IWbemRefreshingServices
	// interface pointer.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call.
	//
	// The server MUST return WBEM_S_NO_ERROR (specified in section 2.2.11) to indicate
	// the successful completion of the method.
	//
	// The IWbemRemoteRefresher::RemoteRefresh method MUST be called on the IWbemRemoteRefresher
	// interface pointer returned as a member of the _WBEM_REFRESH_INFO structure from IWbemRefreshingServices
	// methods or on the interface returned by IWbemRefreshingServices::GetRemoteRefresher
	// method invocation.
	RemoteRefresh(context.Context, *RemoteRefreshRequest, ...dcerpc.CallOption) (*RemoteRefreshResponse, error)

	// The IWbemRemoteRefresher::StopRefreshing method MUST remove a set of CIM instances
	// or enumerations from the collection previously configured by the IWbemRefreshingServices
	// interface pointer.
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. In case of success, the server MUST return WBEM_S_NO_ERROR (as
	// specified in section 2.2.11) to indicate the successful completion of the method.
	//
	// The IWbemRemoteRefresher::StopRefreshing method MUST be called on the IWbemRemoteRefresher
	// interface pointer that is returned as a member of the _WBEM_REFRESH_INFO structure
	// from the methods of the IWbemRefreshingServices interface or on the interface that
	// is returned by the IWbemRefreshingServices::GetRemoteRefresher method invocation.
	StopRefreshing(context.Context, *StopRefreshingRequest, ...dcerpc.CallOption) (*StopRefreshingResponse, error)

	// The IWbemRemoteRefresher::Opnum5NotUsedOnWire method MUST return a random GUID that
	// identifies the server object that receives the call.
	//
	//
	// Return Values: This method MUST return an HRESULT value that MUST indicate the status
	// of the method call. The server MUST return WBEM_S_NO_ERROR (specified in section
	// 2.2.11) to indicate the successful completion of the method.
	//
	// Opnum5NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) RemoteRefresherClient
}

type xxx_DefaultRemoteRefresherClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultRemoteRefresherClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultRemoteRefresherClient) RemoteRefresh(ctx context.Context, in *RemoteRefreshRequest, opts ...dcerpc.CallOption) (*RemoteRefreshResponse, error) {
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
	out := &RemoteRefreshResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteRefresherClient) StopRefreshing(ctx context.Context, in *StopRefreshingRequest, opts ...dcerpc.CallOption) (*StopRefreshingResponse, error) {
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
	out := &StopRefreshingResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemoteRefresherClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRemoteRefresherClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultRemoteRefresherClient) IPID(ctx context.Context, ipid *dcom.IPID) RemoteRefresherClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultRemoteRefresherClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewRemoteRefresherClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RemoteRefresherClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RemoteRefresherSyntaxV0_0))...)
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
	return &xxx_DefaultRemoteRefresherClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_RemoteRefreshOperation structure represents the RemoteRefresh operation
type xxx_RemoteRefreshOperation struct {
	This          *dcom.ORPCThis         `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat         `idl:"name:That" json:"that"`
	Flags         int32                  `idl:"name:lFlags" json:"flags"`
	ObjectsLength int32                  `idl:"name:plNumObjects" json:"objects_length"`
	Objects       []*wmi.RefreshedObject `idl:"name:paObjects;size_is:(, plNumObjects)" json:"objects"`
	Return        int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoteRefreshOperation) OpNum() int { return 3 }

func (o *xxx_RemoteRefreshOperation) OpName() string { return "/IWbemRemoteRefresher/v0/RemoteRefresh" }

func (o *xxx_RemoteRefreshOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteRefreshOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteRefreshOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteRefreshOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Objects != nil && o.ObjectsLength == 0 {
		o.ObjectsLength = int32(len(o.Objects))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteRefreshOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plNumObjects {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.ObjectsLength); err != nil {
			return err
		}
	}
	// paObjects {out} (1:{pointer=ref}*(2)*(1))(2:{alias=WBEM_REFRESHED_OBJECT}[dim:0,size_is=plNumObjects](struct))
	{
		if o.Objects != nil || o.ObjectsLength > 0 {
			_ptr_paObjects := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ObjectsLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Objects {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.Objects[i1] != nil {
						if err := o.Objects[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&wmi.RefreshedObject{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.Objects); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&wmi.RefreshedObject{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Objects, _ptr_paObjects); err != nil {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoteRefreshOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plNumObjects {out} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.ObjectsLength); err != nil {
			return err
		}
	}
	// paObjects {out} (1:{pointer=ref}*(2)*(1))(2:{alias=WBEM_REFRESHED_OBJECT}[dim:0,size_is=plNumObjects](struct))
	{
		_ptr_paObjects := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Objects", sizeInfo[0])
			}
			o.Objects = make([]*wmi.RefreshedObject, sizeInfo[0])
			for i1 := range o.Objects {
				i1 := i1
				if o.Objects[i1] == nil {
					o.Objects[i1] = &wmi.RefreshedObject{}
				}
				if err := o.Objects[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_paObjects := func(ptr interface{}) { o.Objects = *ptr.(*[]*wmi.RefreshedObject) }
		if err := w.ReadPointer(&o.Objects, _s_paObjects, _ptr_paObjects); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoteRefreshRequest structure represents the RemoteRefresh operation request
type RemoteRefreshRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lFlags: This parameter is not used, and its value MUST be 0x0.
	Flags int32 `idl:"name:lFlags" json:"flags"`
}

func (o *RemoteRefreshRequest) xxx_ToOp(ctx context.Context, op *xxx_RemoteRefreshOperation) *xxx_RemoteRefreshOperation {
	if op == nil {
		op = &xxx_RemoteRefreshOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Flags = o.Flags
	return op
}

func (o *RemoteRefreshRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoteRefreshOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Flags = op.Flags
}
func (o *RemoteRefreshRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemoteRefreshRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteRefreshOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoteRefreshResponse structure represents the RemoteRefresh operation response
type RemoteRefreshResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// plNumObjects: If successful, plNumObjects MUST be a pointer to the number of CIM
	// instances and enumerations that the method returns. It MUST NOT be NULL.
	ObjectsLength int32 `idl:"name:plNumObjects" json:"objects_length"`
	// paObjects: If successful, paObjects MUST be a pointer to an array of WBEM_REFRESHED_OBJECT
	// objects specified in section 2.2.15. The array MUST contain CIM instances and enumerations.
	// It MUST NOT be NULL.
	Objects []*wmi.RefreshedObject `idl:"name:paObjects;size_is:(, plNumObjects)" json:"objects"`
	// Return: The RemoteRefresh return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RemoteRefreshResponse) xxx_ToOp(ctx context.Context, op *xxx_RemoteRefreshOperation) *xxx_RemoteRefreshOperation {
	if op == nil {
		op = &xxx_RemoteRefreshOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ObjectsLength = o.ObjectsLength
	op.Objects = o.Objects
	op.Return = o.Return
	return op
}

func (o *RemoteRefreshResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoteRefreshOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ObjectsLength = op.ObjectsLength
	o.Objects = op.Objects
	o.Return = op.Return
}
func (o *RemoteRefreshResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemoteRefreshResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoteRefreshOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StopRefreshingOperation structure represents the StopRefreshing operation
type xxx_StopRefreshingOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	IDsLength int32          `idl:"name:lNumIds" json:"ids_length"`
	IDs       []int32        `idl:"name:aplIds;size_is:(lNumIds)" json:"ids"`
	Flags     int32          `idl:"name:lFlags" json:"flags"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_StopRefreshingOperation) OpNum() int { return 4 }

func (o *xxx_StopRefreshingOperation) OpName() string {
	return "/IWbemRemoteRefresher/v0/StopRefreshing"
}

func (o *xxx_StopRefreshingOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.IDs != nil && o.IDsLength == 0 {
		o.IDsLength = int32(len(o.IDs))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopRefreshingOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lNumIds {in} (1:(int32))
	{
		if err := w.WriteData(o.IDsLength); err != nil {
			return err
		}
	}
	// aplIds {in} (1:{pointer=ref}*(1)[dim:0,size_is=lNumIds](int32))
	{
		dimSize1 := uint64(o.IDsLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.IDs {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.IDs[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.IDs); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopRefreshingOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lNumIds {in} (1:(int32))
	{
		if err := w.ReadData(&o.IDsLength); err != nil {
			return err
		}
	}
	// aplIds {in} (1:{pointer=ref}*(1)[dim:0,size_is=lNumIds](int32))
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
			return fmt.Errorf("buffer overflow for size %d of array o.IDs", sizeInfo[0])
		}
		o.IDs = make([]int32, sizeInfo[0])
		for i1 := range o.IDs {
			i1 := i1
			if err := w.ReadData(&o.IDs[i1]); err != nil {
				return err
			}
		}
	}
	// lFlags {in} (1:(int32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopRefreshingOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopRefreshingOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopRefreshingOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// Return {out} (1:{alias=HRESULT}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// StopRefreshingRequest structure represents the StopRefreshing operation request
type StopRefreshingRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lNumIds: MUST be the number of identifiers in the array of object identifiers in
	// the aplIds parameter.
	IDsLength int32 `idl:"name:lNumIds" json:"ids_length"`
	// aplIds: MUST be an array of object identifiers that MUST identify the CIM instances
	// and enumerations to stop refreshing. The object identifier is the m_lCancelId member
	// from the _WBEM_REFRESH_INFO structure that is specified in section 2.2.20 and MUST
	// be obtained from a previous call to the IWbemRefreshingServices::AddObjectToRefresher,
	// IWbemRefreshingServices::AddObjectToRefresherByTemplate, or IWbemRefreshingServices::AddEnumToRefresher
	// method specified in section 3.1.4.12.
	IDs []int32 `idl:"name:aplIds;size_is:(lNumIds)" json:"ids"`
	// lFlags: This parameter is not used, and its value MUST be 0x0.
	Flags int32 `idl:"name:lFlags" json:"flags"`
}

func (o *StopRefreshingRequest) xxx_ToOp(ctx context.Context, op *xxx_StopRefreshingOperation) *xxx_StopRefreshingOperation {
	if op == nil {
		op = &xxx_StopRefreshingOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.IDsLength = o.IDsLength
	op.IDs = o.IDs
	op.Flags = o.Flags
	return op
}

func (o *StopRefreshingRequest) xxx_FromOp(ctx context.Context, op *xxx_StopRefreshingOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.IDsLength = op.IDsLength
	o.IDs = op.IDs
	o.Flags = op.Flags
}
func (o *StopRefreshingRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *StopRefreshingRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StopRefreshingOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StopRefreshingResponse structure represents the StopRefreshing operation response
type StopRefreshingResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The StopRefreshing return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StopRefreshingResponse) xxx_ToOp(ctx context.Context, op *xxx_StopRefreshingOperation) *xxx_StopRefreshingOperation {
	if op == nil {
		op = &xxx_StopRefreshingOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *StopRefreshingResponse) xxx_FromOp(ctx context.Context, op *xxx_StopRefreshingOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *StopRefreshingResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *StopRefreshingResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StopRefreshingOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
