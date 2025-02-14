package ivdsadvisesink

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
	// IVdsAdviseSink interface identifier 8326cd1d-cf59-4936-b786-5efc08798e25
	AdviseSinkIID = &dcom.IID{Data1: 0x8326cd1d, Data2: 0xcf59, Data3: 0x4936, Data4: []byte{0xb7, 0x86, 0x5e, 0xfc, 0x08, 0x79, 0x8e, 0x25}}
	// Syntax UUID
	AdviseSinkSyntaxUUID = &uuid.UUID{TimeLow: 0x8326cd1d, TimeMid: 0xcf59, TimeHiAndVersion: 0x4936, ClockSeqHiAndReserved: 0xb7, ClockSeqLow: 0x86, Node: [6]uint8{0x5e, 0xfc, 0x8, 0x79, 0x8e, 0x25}}
	// Syntax ID
	AdviseSinkSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: AdviseSinkSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IVdsAdviseSink interface.
type AdviseSinkClient interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The OnNotify method passes notifications from VDS to applications.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	OnNotify(context.Context, *OnNotifyRequest, ...dcerpc.CallOption) (*OnNotifyResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) AdviseSinkClient
}

type xxx_DefaultAdviseSinkClient struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultAdviseSinkClient) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultAdviseSinkClient) OnNotify(ctx context.Context, in *OnNotifyRequest, opts ...dcerpc.CallOption) (*OnNotifyResponse, error) {
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
	out := &OnNotifyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAdviseSinkClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAdviseSinkClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultAdviseSinkClient) IPID(ctx context.Context, ipid *dcom.IPID) AdviseSinkClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultAdviseSinkClient{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewAdviseSinkClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AdviseSinkClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AdviseSinkSyntaxV0_0))...)
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
	return &xxx_DefaultAdviseSinkClient{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_OnNotifyOperation structure represents the OnNotify operation
type xxx_OnNotifyOperation struct {
	This                  *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That                  *dcom.ORPCThat      `idl:"name:That" json:"that"`
	NumberOfNotifications int32               `idl:"name:lNumberOfNotifications" json:"number_of_notifications"`
	NotificationArray     []*vds.Notification `idl:"name:pNotificationArray;size_is:(lNumberOfNotifications)" json:"notification_array"`
	Return                int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_OnNotifyOperation) OpNum() int { return 3 }

func (o *xxx_OnNotifyOperation) OpName() string { return "/IVdsAdviseSink/v0/OnNotify" }

func (o *xxx_OnNotifyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.NotificationArray != nil && o.NumberOfNotifications == 0 {
		o.NumberOfNotifications = int32(len(o.NotificationArray))
	}
	if o.NumberOfNotifications < int32(1) || o.NumberOfNotifications > int32(100) {
		return fmt.Errorf("NumberOfNotifications is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnNotifyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// lNumberOfNotifications {in} (1:{range=(1,100)}(int32))
	{
		if err := w.WriteData(o.NumberOfNotifications); err != nil {
			return err
		}
	}
	// pNotificationArray {in} (1:{pointer=ref}*(1))(2:{alias=VDS_NOTIFICATION}[dim:0,size_is=lNumberOfNotifications](struct))
	{
		dimSize1 := uint64(o.NumberOfNotifications)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.NotificationArray {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.NotificationArray[i1] != nil {
				if err := o.NotificationArray[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&vds.Notification{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.NotificationArray); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&vds.Notification{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_OnNotifyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// lNumberOfNotifications {in} (1:{range=(1,100)}(int32))
	{
		if err := w.ReadData(&o.NumberOfNotifications); err != nil {
			return err
		}
	}
	// pNotificationArray {in} (1:{pointer=ref}*(1))(2:{alias=VDS_NOTIFICATION}[dim:0,size_is=lNumberOfNotifications](struct))
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
			return fmt.Errorf("buffer overflow for size %d of array o.NotificationArray", sizeInfo[0])
		}
		o.NotificationArray = make([]*vds.Notification, sizeInfo[0])
		for i1 := range o.NotificationArray {
			i1 := i1
			if o.NotificationArray[i1] == nil {
				o.NotificationArray[i1] = &vds.Notification{}
			}
			if err := o.NotificationArray[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_OnNotifyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnNotifyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_OnNotifyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// OnNotifyRequest structure represents the OnNotify operation request
type OnNotifyRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// lNumberOfNotifications: The number of notifications that are specified in pNotificationArray.
	// This parameter MUST be a value from 1 through 100.
	NumberOfNotifications int32 `idl:"name:lNumberOfNotifications" json:"number_of_notifications"`
	// pNotificationArray: An array of VDS_NOTIFICATION structures.
	NotificationArray []*vds.Notification `idl:"name:pNotificationArray;size_is:(lNumberOfNotifications)" json:"notification_array"`
}

func (o *OnNotifyRequest) xxx_ToOp(ctx context.Context, op *xxx_OnNotifyOperation) *xxx_OnNotifyOperation {
	if op == nil {
		op = &xxx_OnNotifyOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.NumberOfNotifications = o.NumberOfNotifications
	op.NotificationArray = o.NotificationArray
	return op
}

func (o *OnNotifyRequest) xxx_FromOp(ctx context.Context, op *xxx_OnNotifyOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.NumberOfNotifications = op.NumberOfNotifications
	o.NotificationArray = op.NotificationArray
}
func (o *OnNotifyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OnNotifyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OnNotifyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OnNotifyResponse structure represents the OnNotify operation response
type OnNotifyResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The OnNotify return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OnNotifyResponse) xxx_ToOp(ctx context.Context, op *xxx_OnNotifyOperation) *xxx_OnNotifyOperation {
	if op == nil {
		op = &xxx_OnNotifyOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *OnNotifyResponse) xxx_FromOp(ctx context.Context, op *xxx_OnNotifyOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *OnNotifyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OnNotifyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OnNotifyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
