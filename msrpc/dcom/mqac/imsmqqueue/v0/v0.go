package imsmqqueue

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	mqac "github.com/oiweiwei/go-msrpc/msrpc/dcom/mqac"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
	_ = mqac.GoPackage
	_ = oaut.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/mqac"
)

var (
	// IMSMQQueue interface identifier d7d6e076-dccd-11d0-aa4b-0060970debae
	QueueIID = &dcom.IID{Data1: 0xd7d6e076, Data2: 0xdccd, Data3: 0x11d0, Data4: []byte{0xaa, 0x4b, 0x00, 0x60, 0x97, 0x0d, 0xeb, 0xae}}
	// Syntax UUID
	QueueSyntaxUUID = &uuid.UUID{TimeLow: 0xd7d6e076, TimeMid: 0xdccd, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x4b, Node: [6]uint8{0x0, 0x60, 0x97, 0xd, 0xeb, 0xae}}
	// Syntax ID
	QueueSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: QueueSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IMSMQQueue interface.
type QueueClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

	// Access operation.
	GetAccess(context.Context, *GetAccessRequest, ...dcerpc.CallOption) (*GetAccessResponse, error)

	// ShareMode operation.
	GetShareMode(context.Context, *GetShareModeRequest, ...dcerpc.CallOption) (*GetShareModeResponse, error)

	// QueueInfo operation.
	GetQueueInfo(context.Context, *GetQueueInfoRequest, ...dcerpc.CallOption) (*GetQueueInfoResponse, error)

	// Handle operation.
	GetHandle(context.Context, *GetHandleRequest, ...dcerpc.CallOption) (*GetHandleResponse, error)

	// IsOpen operation.
	GetIsOpen(context.Context, *GetIsOpenRequest, ...dcerpc.CallOption) (*GetIsOpenResponse, error)

	// Close operation.
	Close(context.Context, *CloseRequest, ...dcerpc.CallOption) (*CloseResponse, error)

	// Receive operation.
	Receive(context.Context, *ReceiveRequest, ...dcerpc.CallOption) (*ReceiveResponse, error)

	// Peek operation.
	Peek(context.Context, *PeekRequest, ...dcerpc.CallOption) (*PeekResponse, error)

	// EnableNotification operation.
	EnableNotification(context.Context, *EnableNotificationRequest, ...dcerpc.CallOption) (*EnableNotificationResponse, error)

	// Reset operation.
	Reset(context.Context, *ResetRequest, ...dcerpc.CallOption) (*ResetResponse, error)

	// ReceiveCurrent operation.
	ReceiveCurrent(context.Context, *ReceiveCurrentRequest, ...dcerpc.CallOption) (*ReceiveCurrentResponse, error)

	// PeekNext operation.
	PeekNext(context.Context, *PeekNextRequest, ...dcerpc.CallOption) (*PeekNextResponse, error)

	// PeekCurrent operation.
	PeekCurrent(context.Context, *PeekCurrentRequest, ...dcerpc.CallOption) (*PeekCurrentResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) QueueClient
}

type xxx_DefaultQueueClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultQueueClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultQueueClient) GetAccess(ctx context.Context, in *GetAccessRequest, opts ...dcerpc.CallOption) (*GetAccessResponse, error) {
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
	out := &GetAccessResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueClient) GetShareMode(ctx context.Context, in *GetShareModeRequest, opts ...dcerpc.CallOption) (*GetShareModeResponse, error) {
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
	out := &GetShareModeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueClient) GetQueueInfo(ctx context.Context, in *GetQueueInfoRequest, opts ...dcerpc.CallOption) (*GetQueueInfoResponse, error) {
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
	out := &GetQueueInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueClient) GetHandle(ctx context.Context, in *GetHandleRequest, opts ...dcerpc.CallOption) (*GetHandleResponse, error) {
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
	out := &GetHandleResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueClient) GetIsOpen(ctx context.Context, in *GetIsOpenRequest, opts ...dcerpc.CallOption) (*GetIsOpenResponse, error) {
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
	out := &GetIsOpenResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueClient) Close(ctx context.Context, in *CloseRequest, opts ...dcerpc.CallOption) (*CloseResponse, error) {
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
	out := &CloseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueClient) Receive(ctx context.Context, in *ReceiveRequest, opts ...dcerpc.CallOption) (*ReceiveResponse, error) {
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
	out := &ReceiveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueClient) Peek(ctx context.Context, in *PeekRequest, opts ...dcerpc.CallOption) (*PeekResponse, error) {
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
	out := &PeekResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueClient) EnableNotification(ctx context.Context, in *EnableNotificationRequest, opts ...dcerpc.CallOption) (*EnableNotificationResponse, error) {
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
	out := &EnableNotificationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueClient) Reset(ctx context.Context, in *ResetRequest, opts ...dcerpc.CallOption) (*ResetResponse, error) {
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
	out := &ResetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueClient) ReceiveCurrent(ctx context.Context, in *ReceiveCurrentRequest, opts ...dcerpc.CallOption) (*ReceiveCurrentResponse, error) {
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
	out := &ReceiveCurrentResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueClient) PeekNext(ctx context.Context, in *PeekNextRequest, opts ...dcerpc.CallOption) (*PeekNextResponse, error) {
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
	out := &PeekNextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueClient) PeekCurrent(ctx context.Context, in *PeekCurrentRequest, opts ...dcerpc.CallOption) (*PeekCurrentResponse, error) {
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
	out := &PeekCurrentResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultQueueClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultQueueClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultQueueClient) IPID(ctx context.Context, ipid *dcom.IPID) QueueClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultQueueClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewQueueClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (QueueClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(QueueSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultQueueClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetAccessOperation structure represents the Access operation
type xxx_GetAccessOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Access int32          `idl:"name:plAccess" json:"access"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAccessOperation) OpNum() int { return 7 }

func (o *xxx_GetAccessOperation) OpName() string { return "/IMSMQQueue/v0/Access" }

func (o *xxx_GetAccessOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAccessOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetAccessOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetAccessOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAccessOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plAccess {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Access); err != nil {
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

func (o *xxx_GetAccessOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plAccess {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Access); err != nil {
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

// GetAccessRequest structure represents the Access operation request
type GetAccessRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetAccessRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAccessOperation) *xxx_GetAccessOperation {
	if op == nil {
		op = &xxx_GetAccessOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetAccessRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAccessOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetAccessRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAccessRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAccessOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAccessResponse structure represents the Access operation response
type GetAccessResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Access int32          `idl:"name:plAccess" json:"access"`
	// Return: The Access return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetAccessResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAccessOperation) *xxx_GetAccessOperation {
	if op == nil {
		op = &xxx_GetAccessOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Access = o.Access
	op.Return = o.Return
	return op
}

func (o *GetAccessResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAccessOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Access = op.Access
	o.Return = op.Return
}
func (o *GetAccessResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAccessResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAccessOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetShareModeOperation structure represents the ShareMode operation
type xxx_GetShareModeOperation struct {
	This      *dcom.ORPCThis `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	ShareMode int32          `idl:"name:plShareMode" json:"share_mode"`
	Return    int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetShareModeOperation) OpNum() int { return 8 }

func (o *xxx_GetShareModeOperation) OpName() string { return "/IMSMQQueue/v0/ShareMode" }

func (o *xxx_GetShareModeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetShareModeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetShareModeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetShareModeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetShareModeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plShareMode {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.ShareMode); err != nil {
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

func (o *xxx_GetShareModeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plShareMode {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.ShareMode); err != nil {
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

// GetShareModeRequest structure represents the ShareMode operation request
type GetShareModeRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetShareModeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetShareModeOperation) *xxx_GetShareModeOperation {
	if op == nil {
		op = &xxx_GetShareModeOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetShareModeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetShareModeOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetShareModeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetShareModeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetShareModeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetShareModeResponse structure represents the ShareMode operation response
type GetShareModeResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat `idl:"name:That" json:"that"`
	ShareMode int32          `idl:"name:plShareMode" json:"share_mode"`
	// Return: The ShareMode return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetShareModeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetShareModeOperation) *xxx_GetShareModeOperation {
	if op == nil {
		op = &xxx_GetShareModeOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ShareMode = o.ShareMode
	op.Return = o.Return
	return op
}

func (o *GetShareModeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetShareModeOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ShareMode = op.ShareMode
	o.Return = op.Return
}
func (o *GetShareModeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetShareModeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetShareModeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetQueueInfoOperation structure represents the QueueInfo operation
type xxx_GetQueueInfoOperation struct {
	This      *dcom.ORPCThis  `idl:"name:This" json:"this"`
	That      *dcom.ORPCThat  `idl:"name:That" json:"that"`
	QueueInfo *mqac.QueueInfo `idl:"name:ppqinfo" json:"queue_info"`
	Return    int32           `idl:"name:Return" json:"return"`
}

func (o *xxx_GetQueueInfoOperation) OpNum() int { return 9 }

func (o *xxx_GetQueueInfoOperation) OpName() string { return "/IMSMQQueue/v0/QueueInfo" }

func (o *xxx_GetQueueInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQueueInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetQueueInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetQueueInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQueueInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppqinfo {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo}(interface))
	{
		if o.QueueInfo != nil {
			_ptr_ppqinfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.QueueInfo != nil {
					if err := o.QueueInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.QueueInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.QueueInfo, _ptr_ppqinfo); err != nil {
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

func (o *xxx_GetQueueInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppqinfo {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQQueueInfo}(interface))
	{
		_ptr_ppqinfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.QueueInfo == nil {
				o.QueueInfo = &mqac.QueueInfo{}
			}
			if err := o.QueueInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppqinfo := func(ptr interface{}) { o.QueueInfo = *ptr.(**mqac.QueueInfo) }
		if err := w.ReadPointer(&o.QueueInfo, _s_ppqinfo, _ptr_ppqinfo); err != nil {
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

// GetQueueInfoRequest structure represents the QueueInfo operation request
type GetQueueInfoRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetQueueInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetQueueInfoOperation) *xxx_GetQueueInfoOperation {
	if op == nil {
		op = &xxx_GetQueueInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetQueueInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetQueueInfoOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetQueueInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetQueueInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQueueInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetQueueInfoResponse structure represents the QueueInfo operation response
type GetQueueInfoResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That      *dcom.ORPCThat  `idl:"name:That" json:"that"`
	QueueInfo *mqac.QueueInfo `idl:"name:ppqinfo" json:"queue_info"`
	// Return: The QueueInfo return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetQueueInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetQueueInfoOperation) *xxx_GetQueueInfoOperation {
	if op == nil {
		op = &xxx_GetQueueInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.QueueInfo = o.QueueInfo
	op.Return = o.Return
	return op
}

func (o *GetQueueInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetQueueInfoOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.QueueInfo = op.QueueInfo
	o.Return = op.Return
}
func (o *GetQueueInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetQueueInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQueueInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetHandleOperation structure represents the Handle operation
type xxx_GetHandleOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle int32          `idl:"name:plHandle" json:"handle"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetHandleOperation) OpNum() int { return 10 }

func (o *xxx_GetHandleOperation) OpName() string { return "/IMSMQQueue/v0/Handle" }

func (o *xxx_GetHandleOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHandleOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetHandleOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetHandleOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetHandleOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// plHandle {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.WriteData(o.Handle); err != nil {
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

func (o *xxx_GetHandleOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// plHandle {out, retval} (1:{pointer=ref}*(1)(int32))
	{
		if err := w.ReadData(&o.Handle); err != nil {
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

// GetHandleRequest structure represents the Handle operation request
type GetHandleRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetHandleRequest) xxx_ToOp(ctx context.Context, op *xxx_GetHandleOperation) *xxx_GetHandleOperation {
	if op == nil {
		op = &xxx_GetHandleOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetHandleRequest) xxx_FromOp(ctx context.Context, op *xxx_GetHandleOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetHandleRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetHandleRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetHandleOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetHandleResponse structure represents the Handle operation response
type GetHandleResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Handle int32          `idl:"name:plHandle" json:"handle"`
	// Return: The Handle return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetHandleResponse) xxx_ToOp(ctx context.Context, op *xxx_GetHandleOperation) *xxx_GetHandleOperation {
	if op == nil {
		op = &xxx_GetHandleOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Handle = o.Handle
	op.Return = o.Return
	return op
}

func (o *GetHandleResponse) xxx_FromOp(ctx context.Context, op *xxx_GetHandleOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Handle = op.Handle
	o.Return = op.Return
}
func (o *GetHandleResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetHandleResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetHandleOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIsOpenOperation structure represents the IsOpen operation
type xxx_GetIsOpenOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsOpen int16          `idl:"name:pisOpen" json:"is_open"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIsOpenOperation) OpNum() int { return 11 }

func (o *xxx_GetIsOpenOperation) OpName() string { return "/IMSMQQueue/v0/IsOpen" }

func (o *xxx_GetIsOpenOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsOpenOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetIsOpenOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_GetIsOpenOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIsOpenOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// pisOpen {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.WriteData(o.IsOpen); err != nil {
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

func (o *xxx_GetIsOpenOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// pisOpen {out, retval} (1:{pointer=ref}*(1)(int16))
	{
		if err := w.ReadData(&o.IsOpen); err != nil {
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

// GetIsOpenRequest structure represents the IsOpen operation request
type GetIsOpenRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIsOpenRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIsOpenOperation) *xxx_GetIsOpenOperation {
	if op == nil {
		op = &xxx_GetIsOpenOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIsOpenRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIsOpenOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIsOpenRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIsOpenRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsOpenOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIsOpenResponse structure represents the IsOpen operation response
type GetIsOpenResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsOpen int16          `idl:"name:pisOpen" json:"is_open"`
	// Return: The IsOpen return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIsOpenResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIsOpenOperation) *xxx_GetIsOpenOperation {
	if op == nil {
		op = &xxx_GetIsOpenOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.IsOpen = o.IsOpen
	op.Return = o.Return
	return op
}

func (o *GetIsOpenResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIsOpenOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IsOpen = op.IsOpen
	o.Return = op.Return
}
func (o *GetIsOpenResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIsOpenResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIsOpenOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseOperation structure represents the Close operation
type xxx_CloseOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseOperation) OpNum() int { return 12 }

func (o *xxx_CloseOperation) OpName() string { return "/IMSMQQueue/v0/Close" }

func (o *xxx_CloseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CloseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_CloseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_CloseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// CloseRequest structure represents the Close operation request
type CloseRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CloseRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseOperation) *xxx_CloseOperation {
	if op == nil {
		op = &xxx_CloseOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *CloseRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CloseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseResponse structure represents the Close operation response
type CloseResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Close return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloseResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseOperation) *xxx_CloseOperation {
	if op == nil {
		op = &xxx_CloseOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *CloseResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CloseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReceiveOperation structure represents the Receive operation
type xxx_ReceiveOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	Message              *mqac.Message  `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReceiveOperation) OpNum() int { return 13 }

func (o *xxx_ReceiveOperation) OpName() string { return "/IMSMQQueue/v0/Receive" }

func (o *xxx_ReceiveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Transaction != nil {
			_ptr_Transaction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Transaction != nil {
					if err := o.Transaction.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Transaction, _ptr_Transaction); err != nil {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ReceiveTimeout != nil {
			_ptr_ReceiveTimeout := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReceiveTimeout != nil {
					if err := o.ReceiveTimeout.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
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
	return nil
}

func (o *xxx_ReceiveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_Transaction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Transaction == nil {
				o.Transaction = &oaut.Variant{}
			}
			if err := o.Transaction.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Transaction := func(ptr interface{}) { o.Transaction = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Transaction, _s_Transaction, _ptr_Transaction); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ReceiveTimeout := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReceiveTimeout == nil {
				o.ReceiveTimeout = &oaut.Variant{}
			}
			if err := o.ReceiveTimeout.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReceiveTimeout := func(ptr interface{}) { o.ReceiveTimeout = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ReceiveTimeout, _s_ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_ReceiveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// ReceiveRequest structure represents the Receive operation request
type ReceiveRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
}

func (o *ReceiveRequest) xxx_ToOp(ctx context.Context, op *xxx_ReceiveOperation) *xxx_ReceiveOperation {
	if op == nil {
		op = &xxx_ReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Transaction = o.Transaction
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.ReceiveTimeout = o.ReceiveTimeout
	return op
}

func (o *ReceiveRequest) xxx_FromOp(ctx context.Context, op *xxx_ReceiveOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Transaction = op.Transaction
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.ReceiveTimeout = op.ReceiveTimeout
}
func (o *ReceiveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReceiveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReceiveResponse structure represents the Receive operation response
type ReceiveResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message  `idl:"name:ppmsg" json:"message"`
	// Return: The Receive return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReceiveResponse) xxx_ToOp(ctx context.Context, op *xxx_ReceiveOperation) *xxx_ReceiveOperation {
	if op == nil {
		op = &xxx_ReceiveOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *ReceiveResponse) xxx_FromOp(ctx context.Context, op *xxx_ReceiveOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *ReceiveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReceiveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PeekOperation structure represents the Peek operation
type xxx_PeekOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	Message              *mqac.Message  `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PeekOperation) OpNum() int { return 14 }

func (o *xxx_PeekOperation) OpName() string { return "/IMSMQQueue/v0/Peek" }

func (o *xxx_PeekOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ReceiveTimeout != nil {
			_ptr_ReceiveTimeout := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReceiveTimeout != nil {
					if err := o.ReceiveTimeout.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
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
	return nil
}

func (o *xxx_PeekOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ReceiveTimeout := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReceiveTimeout == nil {
				o.ReceiveTimeout = &oaut.Variant{}
			}
			if err := o.ReceiveTimeout.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReceiveTimeout := func(ptr interface{}) { o.ReceiveTimeout = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ReceiveTimeout, _s_ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_PeekOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// PeekRequest structure represents the Peek operation request
type PeekRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
}

func (o *PeekRequest) xxx_ToOp(ctx context.Context, op *xxx_PeekOperation) *xxx_PeekOperation {
	if op == nil {
		op = &xxx_PeekOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.ReceiveTimeout = o.ReceiveTimeout
	return op
}

func (o *PeekRequest) xxx_FromOp(ctx context.Context, op *xxx_PeekOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.ReceiveTimeout = op.ReceiveTimeout
}
func (o *PeekRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PeekRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PeekResponse structure represents the Peek operation response
type PeekResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message  `idl:"name:ppmsg" json:"message"`
	// Return: The Peek return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PeekResponse) xxx_ToOp(ctx context.Context, op *xxx_PeekOperation) *xxx_PeekOperation {
	if op == nil {
		op = &xxx_PeekOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *PeekResponse) xxx_FromOp(ctx context.Context, op *xxx_PeekOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *PeekResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PeekResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnableNotificationOperation structure represents the EnableNotification operation
type xxx_EnableNotificationOperation struct {
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat `idl:"name:That" json:"that"`
	Event          *mqac.Event    `idl:"name:Event" json:"event"`
	Cursor         *oaut.Variant  `idl:"name:Cursor" json:"cursor"`
	ReceiveTimeout *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	Return         int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_EnableNotificationOperation) OpNum() int { return 15 }

func (o *xxx_EnableNotificationOperation) OpName() string { return "/IMSMQQueue/v0/EnableNotification" }

func (o *xxx_EnableNotificationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableNotificationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Event {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQEvent}(interface))
	{
		if o.Event != nil {
			_ptr_Event := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Event != nil {
					if err := o.Event.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Event{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Event, _ptr_Event); err != nil {
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
	// Cursor {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Cursor != nil {
			_ptr_Cursor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Cursor != nil {
					if err := o.Cursor.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Cursor, _ptr_Cursor); err != nil {
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
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ReceiveTimeout != nil {
			_ptr_ReceiveTimeout := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReceiveTimeout != nil {
					if err := o.ReceiveTimeout.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
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
	return nil
}

func (o *xxx_EnableNotificationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Event {in} (1:{pointer=ref}*(1))(2:{alias=IMSMQEvent}(interface))
	{
		_ptr_Event := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Event == nil {
				o.Event = &mqac.Event{}
			}
			if err := o.Event.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Event := func(ptr interface{}) { o.Event = *ptr.(**mqac.Event) }
		if err := w.ReadPointer(&o.Event, _s_Event, _ptr_Event); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Cursor {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_Cursor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Cursor == nil {
				o.Cursor = &oaut.Variant{}
			}
			if err := o.Cursor.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Cursor := func(ptr interface{}) { o.Cursor = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Cursor, _s_Cursor, _ptr_Cursor); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ReceiveTimeout := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReceiveTimeout == nil {
				o.ReceiveTimeout = &oaut.Variant{}
			}
			if err := o.ReceiveTimeout.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReceiveTimeout := func(ptr interface{}) { o.ReceiveTimeout = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ReceiveTimeout, _s_ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableNotificationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableNotificationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_EnableNotificationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// EnableNotificationRequest structure represents the EnableNotification operation request
type EnableNotificationRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This           *dcom.ORPCThis `idl:"name:This" json:"this"`
	Event          *mqac.Event    `idl:"name:Event" json:"event"`
	Cursor         *oaut.Variant  `idl:"name:Cursor" json:"cursor"`
	ReceiveTimeout *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
}

func (o *EnableNotificationRequest) xxx_ToOp(ctx context.Context, op *xxx_EnableNotificationOperation) *xxx_EnableNotificationOperation {
	if op == nil {
		op = &xxx_EnableNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Event = o.Event
	op.Cursor = o.Cursor
	op.ReceiveTimeout = o.ReceiveTimeout
	return op
}

func (o *EnableNotificationRequest) xxx_FromOp(ctx context.Context, op *xxx_EnableNotificationOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Event = op.Event
	o.Cursor = op.Cursor
	o.ReceiveTimeout = op.ReceiveTimeout
}
func (o *EnableNotificationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnableNotificationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnableNotificationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnableNotificationResponse structure represents the EnableNotification operation response
type EnableNotificationResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The EnableNotification return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EnableNotificationResponse) xxx_ToOp(ctx context.Context, op *xxx_EnableNotificationOperation) *xxx_EnableNotificationOperation {
	if op == nil {
		op = &xxx_EnableNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *EnableNotificationResponse) xxx_FromOp(ctx context.Context, op *xxx_EnableNotificationOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *EnableNotificationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnableNotificationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnableNotificationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ResetOperation structure represents the Reset operation
type xxx_ResetOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ResetOperation) OpNum() int { return 16 }

func (o *xxx_ResetOperation) OpName() string { return "/IMSMQQueue/v0/Reset" }

func (o *xxx_ResetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ResetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_ResetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ResetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// ResetRequest structure represents the Reset operation request
type ResetRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *ResetRequest) xxx_ToOp(ctx context.Context, op *xxx_ResetOperation) *xxx_ResetOperation {
	if op == nil {
		op = &xxx_ResetOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *ResetRequest) xxx_FromOp(ctx context.Context, op *xxx_ResetOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *ResetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ResetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ResetResponse structure represents the Reset operation response
type ResetResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Reset return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ResetResponse) xxx_ToOp(ctx context.Context, op *xxx_ResetOperation) *xxx_ResetOperation {
	if op == nil {
		op = &xxx_ResetOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *ResetResponse) xxx_FromOp(ctx context.Context, op *xxx_ResetOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ResetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ResetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReceiveCurrentOperation structure represents the ReceiveCurrent operation
type xxx_ReceiveCurrentOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	Message              *mqac.Message  `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ReceiveCurrentOperation) OpNum() int { return 17 }

func (o *xxx_ReceiveCurrentOperation) OpName() string { return "/IMSMQQueue/v0/ReceiveCurrent" }

func (o *xxx_ReceiveCurrentOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveCurrentOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.Transaction != nil {
			_ptr_Transaction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Transaction != nil {
					if err := o.Transaction.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Transaction, _ptr_Transaction); err != nil {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ReceiveTimeout != nil {
			_ptr_ReceiveTimeout := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReceiveTimeout != nil {
					if err := o.ReceiveTimeout.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
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
	return nil
}

func (o *xxx_ReceiveCurrentOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// Transaction {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_Transaction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Transaction == nil {
				o.Transaction = &oaut.Variant{}
			}
			if err := o.Transaction.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_Transaction := func(ptr interface{}) { o.Transaction = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.Transaction, _s_Transaction, _ptr_Transaction); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ReceiveTimeout := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReceiveTimeout == nil {
				o.ReceiveTimeout = &oaut.Variant{}
			}
			if err := o.ReceiveTimeout.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReceiveTimeout := func(ptr interface{}) { o.ReceiveTimeout = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ReceiveTimeout, _s_ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveCurrentOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveCurrentOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_ReceiveCurrentOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// ReceiveCurrentRequest structure represents the ReceiveCurrent operation request
type ReceiveCurrentRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	Transaction          *oaut.Variant  `idl:"name:Transaction" json:"transaction"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
}

func (o *ReceiveCurrentRequest) xxx_ToOp(ctx context.Context, op *xxx_ReceiveCurrentOperation) *xxx_ReceiveCurrentOperation {
	if op == nil {
		op = &xxx_ReceiveCurrentOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Transaction = o.Transaction
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.ReceiveTimeout = o.ReceiveTimeout
	return op
}

func (o *ReceiveCurrentRequest) xxx_FromOp(ctx context.Context, op *xxx_ReceiveCurrentOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Transaction = op.Transaction
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.ReceiveTimeout = op.ReceiveTimeout
}
func (o *ReceiveCurrentRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReceiveCurrentRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveCurrentOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReceiveCurrentResponse structure represents the ReceiveCurrent operation response
type ReceiveCurrentResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message  `idl:"name:ppmsg" json:"message"`
	// Return: The ReceiveCurrent return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ReceiveCurrentResponse) xxx_ToOp(ctx context.Context, op *xxx_ReceiveCurrentOperation) *xxx_ReceiveCurrentOperation {
	if op == nil {
		op = &xxx_ReceiveCurrentOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *ReceiveCurrentResponse) xxx_FromOp(ctx context.Context, op *xxx_ReceiveCurrentOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *ReceiveCurrentResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReceiveCurrentResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveCurrentOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PeekNextOperation structure represents the PeekNext operation
type xxx_PeekNextOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	Message              *mqac.Message  `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PeekNextOperation) OpNum() int { return 18 }

func (o *xxx_PeekNextOperation) OpName() string { return "/IMSMQQueue/v0/PeekNext" }

func (o *xxx_PeekNextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekNextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ReceiveTimeout != nil {
			_ptr_ReceiveTimeout := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReceiveTimeout != nil {
					if err := o.ReceiveTimeout.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
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
	return nil
}

func (o *xxx_PeekNextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ReceiveTimeout := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReceiveTimeout == nil {
				o.ReceiveTimeout = &oaut.Variant{}
			}
			if err := o.ReceiveTimeout.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReceiveTimeout := func(ptr interface{}) { o.ReceiveTimeout = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ReceiveTimeout, _s_ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekNextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekNextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_PeekNextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// PeekNextRequest structure represents the PeekNext operation request
type PeekNextRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
}

func (o *PeekNextRequest) xxx_ToOp(ctx context.Context, op *xxx_PeekNextOperation) *xxx_PeekNextOperation {
	if op == nil {
		op = &xxx_PeekNextOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.ReceiveTimeout = o.ReceiveTimeout
	return op
}

func (o *PeekNextRequest) xxx_FromOp(ctx context.Context, op *xxx_PeekNextOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.ReceiveTimeout = op.ReceiveTimeout
}
func (o *PeekNextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PeekNextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekNextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PeekNextResponse structure represents the PeekNext operation response
type PeekNextResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message  `idl:"name:ppmsg" json:"message"`
	// Return: The PeekNext return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PeekNextResponse) xxx_ToOp(ctx context.Context, op *xxx_PeekNextOperation) *xxx_PeekNextOperation {
	if op == nil {
		op = &xxx_PeekNextOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *PeekNextResponse) xxx_FromOp(ctx context.Context, op *xxx_PeekNextOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *PeekNextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PeekNextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekNextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PeekCurrentOperation structure represents the PeekCurrent operation
type xxx_PeekCurrentOperation struct {
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	That                 *dcom.ORPCThat `idl:"name:That" json:"that"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
	Message              *mqac.Message  `idl:"name:ppmsg" json:"message"`
	Return               int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_PeekCurrentOperation) OpNum() int { return 19 }

func (o *xxx_PeekCurrentOperation) OpName() string { return "/IMSMQQueue/v0/PeekCurrent" }

func (o *xxx_PeekCurrentOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekCurrentOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantDestinationQueue != nil {
			_ptr_WantDestinationQueue := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantDestinationQueue != nil {
					if err := o.WantDestinationQueue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
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
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.WantBody != nil {
			_ptr_WantBody := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WantBody != nil {
					if err := o.WantBody.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WantBody, _ptr_WantBody); err != nil {
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
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT}*(1))(3:{alias=_VARIANT}(struct))
	{
		if o.ReceiveTimeout != nil {
			_ptr_ReceiveTimeout := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReceiveTimeout != nil {
					if err := o.ReceiveTimeout.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.Variant{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
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
	return nil
}

func (o *xxx_PeekCurrentOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// WantDestinationQueue {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantDestinationQueue := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantDestinationQueue == nil {
				o.WantDestinationQueue = &oaut.Variant{}
			}
			if err := o.WantDestinationQueue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantDestinationQueue := func(ptr interface{}) { o.WantDestinationQueue = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantDestinationQueue, _s_WantDestinationQueue, _ptr_WantDestinationQueue); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// WantBody {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_WantBody := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WantBody == nil {
				o.WantBody = &oaut.Variant{}
			}
			if err := o.WantBody.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_WantBody := func(ptr interface{}) { o.WantBody = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.WantBody, _s_WantBody, _ptr_WantBody); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ReceiveTimeout {in, optional} (1:{pointer=ref}*(2))(2:{alias=VARIANT,pointer=ref}*(1))(3:{alias=_VARIANT}(struct))
	{
		_ptr_ReceiveTimeout := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReceiveTimeout == nil {
				o.ReceiveTimeout = &oaut.Variant{}
			}
			if err := o.ReceiveTimeout.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReceiveTimeout := func(ptr interface{}) { o.ReceiveTimeout = *ptr.(**oaut.Variant) }
		if err := w.ReadPointer(&o.ReceiveTimeout, _s_ReceiveTimeout, _ptr_ReceiveTimeout); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekCurrentOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PeekCurrentOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage}(interface))
	{
		if o.Message != nil {
			_ptr_ppmsg := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&mqac.Message{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_ppmsg); err != nil {
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

func (o *xxx_PeekCurrentOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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
	// ppmsg {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IMSMQMessage}(interface))
	{
		_ptr_ppmsg := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &mqac.Message{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppmsg := func(ptr interface{}) { o.Message = *ptr.(**mqac.Message) }
		if err := w.ReadPointer(&o.Message, _s_ppmsg, _ptr_ppmsg); err != nil {
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

// PeekCurrentRequest structure represents the PeekCurrent operation request
type PeekCurrentRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This                 *dcom.ORPCThis `idl:"name:This" json:"this"`
	WantDestinationQueue *oaut.Variant  `idl:"name:WantDestinationQueue" json:"want_destination_queue"`
	WantBody             *oaut.Variant  `idl:"name:WantBody" json:"want_body"`
	ReceiveTimeout       *oaut.Variant  `idl:"name:ReceiveTimeout" json:"receive_timeout"`
}

func (o *PeekCurrentRequest) xxx_ToOp(ctx context.Context, op *xxx_PeekCurrentOperation) *xxx_PeekCurrentOperation {
	if op == nil {
		op = &xxx_PeekCurrentOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.WantDestinationQueue = o.WantDestinationQueue
	op.WantBody = o.WantBody
	op.ReceiveTimeout = o.ReceiveTimeout
	return op
}

func (o *PeekCurrentRequest) xxx_FromOp(ctx context.Context, op *xxx_PeekCurrentOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.WantDestinationQueue = op.WantDestinationQueue
	o.WantBody = op.WantBody
	o.ReceiveTimeout = op.ReceiveTimeout
}
func (o *PeekCurrentRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PeekCurrentRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekCurrentOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PeekCurrentResponse structure represents the PeekCurrent operation response
type PeekCurrentResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That    *dcom.ORPCThat `idl:"name:That" json:"that"`
	Message *mqac.Message  `idl:"name:ppmsg" json:"message"`
	// Return: The PeekCurrent return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PeekCurrentResponse) xxx_ToOp(ctx context.Context, op *xxx_PeekCurrentOperation) *xxx_PeekCurrentOperation {
	if op == nil {
		op = &xxx_PeekCurrentOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Message = o.Message
	op.Return = o.Return
	return op
}

func (o *PeekCurrentResponse) xxx_FromOp(ctx context.Context, op *xxx_PeekCurrentOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Message = op.Message
	o.Return = op.Return
}
func (o *PeekCurrentResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PeekCurrentResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PeekCurrentOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
