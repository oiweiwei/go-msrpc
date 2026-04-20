package termsrvnotification

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcetypes "github.com/oiweiwei/go-msrpc/msrpc/dcetypes"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	tsts "github.com/oiweiwei/go-msrpc/msrpc/tsts"
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
	_ = dcetypes.GoPackage
	_ = dtyp.GoPackage
	_ = tsts.GoPackage
)

var (
	// import guard
	GoPackage = "tsts"
)

var (
	// Syntax UUID
	TerminateServerNotificationSyntaxUUID = &uuid.UUID{TimeLow: 0x11899a43, TimeMid: 0x2b68, TimeHiAndVersion: 0x4a76, ClockSeqHiAndReserved: 0x92, ClockSeqLow: 0xe3, Node: [6]uint8{0xa3, 0xd6, 0xad, 0x8c, 0x26, 0xce}}
	// Syntax ID
	TerminateServerNotificationSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: TerminateServerNotificationSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// TermSrvNotification interface.
type TerminateServerNotificationClient interface {

	// The RpcWaitForSessionState method blocks until the state of the specified session
	// running on a terminal server changes to the desired state. The caller MUST have WINSTATION_QUERY
	// permission for the session. The method checks whether the caller has WINSTATION_QUERY
	// permission (section 3.1.1) by setting it as the Access Request mask, and fails if
	// the caller does not have the permission.
	//
	// Return Values:  The method MUST return S_OK (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	WaitForSessionState(context.Context, *WaitForSessionStateRequest, ...dcerpc.CallOption) (*WaitForSessionStateResponse, error)

	// The RpcRegisterAsyncNotification method registers for an event or events happening
	// on a terminal server. The caller MUST call RpcWaitAsyncNotification after calling
	// RpcRegisterAsyncNotification to wait for the notification. No special permissions
	// are required to call this method.
	//
	// Return Values:  The method MUST return S_OK (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	RegisterAsyncNotification(context.Context, *RegisterAsyncNotificationRequest, ...dcerpc.CallOption) (*RegisterAsyncNotificationResponse, error)

	// The RpcWaitAsyncNotification method starts the wait for the specified terminal server
	// notification. The notification object specified in RpcRegisterAsyncNotification is
	// called by RPC when a notification occurs. This is asynchronous notification and RpcWaitAsyncNotification
	// starts the wait for notification and returns. Specify the notification object using
	// RpcRegisterAsyncNotification and then start the notification wait process using RpcWaitAsyncNotification.
	// No special permissions are required to call this method.
	//
	// Return Values:  The method MUST return S_OK (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	WaitAsyncNotification(context.Context, *WaitAsyncNotificationRequest, ...dcerpc.CallOption) (*WaitAsyncNotificationResponse, error)

	// The RpcUnRegisterAsyncNotification method cancels the asynchronous operation of waiting
	// for notification from the terminal server. This MUST be called after RpcRegisterAsyncNotification.
	// The call to this method MUST be serialized if there are multiple threads running
	// otherwise the behavior of this function is unknown. No special permissions are required
	// to call this method.
	//
	// Return Values: The method MUST return S_OK (0x00000000) on success; otherwise, it
	// MUST return an implementation-specific negative value.
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	UnregisterAsyncNotification(context.Context, *UnregisterAsyncNotificationRequest, ...dcerpc.CallOption) (*UnregisterAsyncNotificationResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// Notify structure represents NOTIFY_HANDLE RPC structure.
type Notify dcetypes.ContextHandle

func (o *Notify) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Notify) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Notify) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Notify) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultTerminateServerNotificationClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultTerminateServerNotificationClient) WaitForSessionState(ctx context.Context, in *WaitForSessionStateRequest, opts ...dcerpc.CallOption) (*WaitForSessionStateResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WaitForSessionStateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTerminateServerNotificationClient) RegisterAsyncNotification(ctx context.Context, in *RegisterAsyncNotificationRequest, opts ...dcerpc.CallOption) (*RegisterAsyncNotificationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RegisterAsyncNotificationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTerminateServerNotificationClient) WaitAsyncNotification(ctx context.Context, in *WaitAsyncNotificationRequest, opts ...dcerpc.CallOption) (*WaitAsyncNotificationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WaitAsyncNotificationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTerminateServerNotificationClient) UnregisterAsyncNotification(ctx context.Context, in *UnregisterAsyncNotificationRequest, opts ...dcerpc.CallOption) (*UnregisterAsyncNotificationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UnregisterAsyncNotificationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultTerminateServerNotificationClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultTerminateServerNotificationClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewTerminateServerNotificationClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (TerminateServerNotificationClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(TerminateServerNotificationSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultTerminateServerNotificationClient{cc: cc}, nil
}

// xxx_WaitForSessionStateOperation structure represents the RpcWaitForSessionState operation
type xxx_WaitForSessionStateOperation struct {
	SessionID int32  `idl:"name:SessionId" json:"session_id"`
	State     int32  `idl:"name:State" json:"state"`
	Timeout   uint32 `idl:"name:Timeout" json:"timeout"`
	Return    int32  `idl:"name:Return" json:"return"`
}

// OpNum returns the operation number of RpcWaitForSessionState operation.
func (o *xxx_WaitForSessionStateOperation) OpNum() int { return 0 }

// OpName returns the operation name of RpcWaitForSessionState operation.
func (o *xxx_WaitForSessionStateOperation) OpName() string {
	return "/TermSrvNotification/v1/RpcWaitForSessionState"
}

func (o *xxx_WaitForSessionStateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForSessionStateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SessionId {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.SessionID); err != nil {
			return err
		}
	}
	// State {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.State); err != nil {
			return err
		}
	}
	// Timeout {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForSessionStateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SessionId {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.SessionID); err != nil {
			return err
		}
	}
	// State {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.State); err != nil {
			return err
		}
	}
	// Timeout {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForSessionStateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForSessionStateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitForSessionStateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WaitForSessionStateRequest structure represents the RpcWaitForSessionState operation request
type WaitForSessionStateRequest struct {
	// SessionId:  The ID of the session for which to await state change. This MUST NOT
	// be the session ID of any of the listener sessions.
	SessionID int32 `idl:"name:SessionId" json:"session_id"`
	// State:  The desired state of the session, as specified in WINSTATIONSTATECLASS (section
	// 2.2.1.9) with the exception of State_Idle and State_Listen, for which to wait. The
	// call will return when the session changes to this state.
	State int32 `idl:"name:State" json:"state"`
	// Timeout:  Maximum time, in milliseconds, to wait for the call to return.
	Timeout uint32 `idl:"name:Timeout" json:"timeout"`
}

func (o *WaitForSessionStateRequest) xxx_ToOp(ctx context.Context, op *xxx_WaitForSessionStateOperation) *xxx_WaitForSessionStateOperation {
	if op == nil {
		op = &xxx_WaitForSessionStateOperation{}
	}
	if o == nil {
		return op
	}
	op.SessionID = o.SessionID
	op.State = o.State
	op.Timeout = o.Timeout
	return op
}

func (o *WaitForSessionStateRequest) xxx_FromOp(ctx context.Context, op *xxx_WaitForSessionStateOperation) {
	if o == nil {
		return
	}
	o.SessionID = op.SessionID
	o.State = op.State
	o.Timeout = op.Timeout
}
func (o *WaitForSessionStateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WaitForSessionStateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WaitForSessionStateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MakeWaitForSessionStateRequest build a response structure from the given request structure.
func (o *WaitForSessionStateRequest) MakeResponse() *WaitForSessionStateResponse {
	return &WaitForSessionStateResponse{}
}

// OpNum returns the operation number of RpcWaitForSessionState operation.
func (o *WaitForSessionStateRequest) OpNum() int { return 0 }

// OpName returns the operation name of RpcWaitForSessionState operation.
func (o *WaitForSessionStateRequest) OpName() string {
	return "/TermSrvNotification/v1/RpcWaitForSessionState"
}

// WaitForSessionStateResponse structure represents the RpcWaitForSessionState operation response
type WaitForSessionStateResponse struct {
	// Return: The RpcWaitForSessionState return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *WaitForSessionStateResponse) xxx_ToOp(ctx context.Context, op *xxx_WaitForSessionStateOperation) *xxx_WaitForSessionStateOperation {
	if op == nil {
		op = &xxx_WaitForSessionStateOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *WaitForSessionStateResponse) xxx_FromOp(ctx context.Context, op *xxx_WaitForSessionStateOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *WaitForSessionStateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WaitForSessionStateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WaitForSessionStateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RegisterAsyncNotificationOperation structure represents the RpcRegisterAsyncNotification operation
type xxx_RegisterAsyncNotificationOperation struct {
	SessionID int32   `idl:"name:SessionId" json:"session_id"`
	Mask      uint32  `idl:"name:Mask" json:"mask"`
	Notify    *Notify `idl:"name:phNotify" json:"notify"`
	Return    int32   `idl:"name:Return" json:"return"`
}

// OpNum returns the operation number of RpcRegisterAsyncNotification operation.
func (o *xxx_RegisterAsyncNotificationOperation) OpNum() int { return 1 }

// OpName returns the operation name of RpcRegisterAsyncNotification operation.
func (o *xxx_RegisterAsyncNotificationOperation) OpName() string {
	return "/TermSrvNotification/v1/RpcRegisterAsyncNotification"
}

func (o *xxx_RegisterAsyncNotificationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterAsyncNotificationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SessionId {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.SessionID); err != nil {
			return err
		}
	}
	// Mask {in} (1:{alias=TNotificationId, names=ULONG}(uint32))
	{
		if err := w.WriteData(o.Mask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterAsyncNotificationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SessionId {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.SessionID); err != nil {
			return err
		}
	}
	// Mask {in} (1:{alias=TNotificationId, names=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Mask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterAsyncNotificationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RegisterAsyncNotificationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phNotify {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NOTIFY_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Notify != nil {
			if err := o.Notify.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RegisterAsyncNotificationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phNotify {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NOTIFY_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Notify == nil {
			o.Notify = &Notify{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
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

// RegisterAsyncNotificationRequest structure represents the RpcRegisterAsyncNotification operation request
type RegisterAsyncNotificationRequest struct {
	// SessionId:  The ID of the session for which to wait for notification. If the value
	// is -1, the call will register for notification for all sessions running on the terminal
	// server. This MUST NOT be the session ID of any of the listener sessions.
	SessionID int32 `idl:"name:SessionId" json:"session_id"`
	// Mask:  Specifies the type of notification to wait for. This is of the type TNotificationId.
	Mask uint32 `idl:"name:Mask" json:"mask"`
}

func (o *RegisterAsyncNotificationRequest) xxx_ToOp(ctx context.Context, op *xxx_RegisterAsyncNotificationOperation) *xxx_RegisterAsyncNotificationOperation {
	if op == nil {
		op = &xxx_RegisterAsyncNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.SessionID = o.SessionID
	op.Mask = o.Mask
	return op
}

func (o *RegisterAsyncNotificationRequest) xxx_FromOp(ctx context.Context, op *xxx_RegisterAsyncNotificationOperation) {
	if o == nil {
		return
	}
	o.SessionID = op.SessionID
	o.Mask = op.Mask
}
func (o *RegisterAsyncNotificationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RegisterAsyncNotificationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterAsyncNotificationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MakeRegisterAsyncNotificationRequest build a response structure from the given request structure.
func (o *RegisterAsyncNotificationRequest) MakeResponse() *RegisterAsyncNotificationResponse {
	return &RegisterAsyncNotificationResponse{}
}

// OpNum returns the operation number of RpcRegisterAsyncNotification operation.
func (o *RegisterAsyncNotificationRequest) OpNum() int { return 1 }

// OpName returns the operation name of RpcRegisterAsyncNotification operation.
func (o *RegisterAsyncNotificationRequest) OpName() string {
	return "/TermSrvNotification/v1/RpcRegisterAsyncNotification"
}

// RegisterAsyncNotificationResponse structure represents the RpcRegisterAsyncNotification operation response
type RegisterAsyncNotificationResponse struct {
	// phNotify: Handle to the notification object. For more information, see NOTIFY_HANDLE.
	Notify *Notify `idl:"name:phNotify" json:"notify"`
	// Return: The RpcRegisterAsyncNotification return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *RegisterAsyncNotificationResponse) xxx_ToOp(ctx context.Context, op *xxx_RegisterAsyncNotificationOperation) *xxx_RegisterAsyncNotificationOperation {
	if op == nil {
		op = &xxx_RegisterAsyncNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.Notify = o.Notify
	op.Return = o.Return
	return op
}

func (o *RegisterAsyncNotificationResponse) xxx_FromOp(ctx context.Context, op *xxx_RegisterAsyncNotificationOperation) {
	if o == nil {
		return
	}
	o.Notify = op.Notify
	o.Return = op.Return
}
func (o *RegisterAsyncNotificationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RegisterAsyncNotificationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RegisterAsyncNotificationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WaitAsyncNotificationOperation structure represents the RpcWaitAsyncNotification operation
type xxx_WaitAsyncNotificationOperation struct {
	Notify        *Notify               `idl:"name:hNotify" json:"notify"`
	SessionChange []*tsts.SessionChange `idl:"name:SessionChange;size_is:(, pEntries)" json:"session_change"`
	Entries       uint32                `idl:"name:pEntries" json:"entries"`
	Return        int32                 `idl:"name:Return" json:"return"`
}

// OpNum returns the operation number of RpcWaitAsyncNotification operation.
func (o *xxx_WaitAsyncNotificationOperation) OpNum() int { return 2 }

// OpName returns the operation name of RpcWaitAsyncNotification operation.
func (o *xxx_WaitAsyncNotificationOperation) OpName() string {
	return "/TermSrvNotification/v1/RpcWaitAsyncNotification"
}

func (o *xxx_WaitAsyncNotificationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitAsyncNotificationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNotify {in} (1:{context_handle, alias=NOTIFY_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Notify != nil {
			if err := o.Notify.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_WaitAsyncNotificationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNotify {in} (1:{context_handle, alias=NOTIFY_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Notify == nil {
			o.Notify = &Notify{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitAsyncNotificationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.SessionChange != nil && o.Entries == 0 {
		o.Entries = uint32(len(o.SessionChange))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WaitAsyncNotificationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// SessionChange {out} (1:{pointer=ref}*(2))(2:{alias=PSESSION_CHANGE}*(1))(3:{alias=SESSION_CHANGE}[dim:0,size_is=pEntries](struct))
	{
		if o.SessionChange != nil || o.Entries > 0 {
			_ptr_SessionChange := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.Entries)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.SessionChange {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.SessionChange[i1] != nil {
						if err := o.SessionChange[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&tsts.SessionChange{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.SessionChange); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&tsts.SessionChange{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SessionChange, _ptr_SessionChange); err != nil {
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
	// pEntries {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Entries); err != nil {
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

func (o *xxx_WaitAsyncNotificationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// SessionChange {out} (1:{pointer=ref}*(2))(2:{alias=PSESSION_CHANGE,pointer=ref}*(1))(3:{alias=SESSION_CHANGE}[dim:0,size_is=pEntries](struct))
	{
		_ptr_SessionChange := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.SessionChange", sizeInfo[0])
			}
			o.SessionChange = make([]*tsts.SessionChange, sizeInfo[0])
			for i1 := range o.SessionChange {
				i1 := i1
				if o.SessionChange[i1] == nil {
					o.SessionChange[i1] = &tsts.SessionChange{}
				}
				if err := o.SessionChange[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_SessionChange := func(ptr interface{}) { o.SessionChange = *ptr.(*[]*tsts.SessionChange) }
		if err := w.ReadPointer(&o.SessionChange, _s_SessionChange, _ptr_SessionChange); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pEntries {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Entries); err != nil {
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

// WaitAsyncNotificationRequest structure represents the RpcWaitAsyncNotification operation request
type WaitAsyncNotificationRequest struct {
	// hNotify:  The notification handle returned by RpcRegisterAsyncNotification.
	Notify *Notify `idl:"name:hNotify" json:"notify"`
}

func (o *WaitAsyncNotificationRequest) xxx_ToOp(ctx context.Context, op *xxx_WaitAsyncNotificationOperation) *xxx_WaitAsyncNotificationOperation {
	if op == nil {
		op = &xxx_WaitAsyncNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.Notify = o.Notify
	return op
}

func (o *WaitAsyncNotificationRequest) xxx_FromOp(ctx context.Context, op *xxx_WaitAsyncNotificationOperation) {
	if o == nil {
		return
	}
	o.Notify = op.Notify
}
func (o *WaitAsyncNotificationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *WaitAsyncNotificationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WaitAsyncNotificationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MakeWaitAsyncNotificationRequest build a response structure from the given request structure.
func (o *WaitAsyncNotificationRequest) MakeResponse() *WaitAsyncNotificationResponse {
	return &WaitAsyncNotificationResponse{}
}

// OpNum returns the operation number of RpcWaitAsyncNotification operation.
func (o *WaitAsyncNotificationRequest) OpNum() int { return 2 }

// OpName returns the operation name of RpcWaitAsyncNotification operation.
func (o *WaitAsyncNotificationRequest) OpName() string {
	return "/TermSrvNotification/v1/RpcWaitAsyncNotification"
}

// WaitAsyncNotificationResponse structure represents the RpcWaitAsyncNotification operation response
type WaitAsyncNotificationResponse struct {
	// SessionChange:  An array of type SESSION_CHANGE containing information about all
	// the sessions and the notifications received for that session.
	SessionChange []*tsts.SessionChange `idl:"name:SessionChange;size_is:(, pEntries)" json:"session_change"`
	// pEntries:  The number of entries returned in the SessionChange array.
	Entries uint32 `idl:"name:pEntries" json:"entries"`
	// Return: The RpcWaitAsyncNotification return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *WaitAsyncNotificationResponse) xxx_ToOp(ctx context.Context, op *xxx_WaitAsyncNotificationOperation) *xxx_WaitAsyncNotificationOperation {
	if op == nil {
		op = &xxx_WaitAsyncNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.SessionChange = o.SessionChange
	op.Entries = o.Entries
	op.Return = o.Return
	return op
}

func (o *WaitAsyncNotificationResponse) xxx_FromOp(ctx context.Context, op *xxx_WaitAsyncNotificationOperation) {
	if o == nil {
		return
	}
	o.SessionChange = op.SessionChange
	o.Entries = op.Entries
	o.Return = op.Return
}
func (o *WaitAsyncNotificationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *WaitAsyncNotificationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WaitAsyncNotificationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UnregisterAsyncNotificationOperation structure represents the RpcUnRegisterAsyncNotification operation
type xxx_UnregisterAsyncNotificationOperation struct {
	Notify *Notify `idl:"name:phNotify" json:"notify"`
	Return int32   `idl:"name:Return" json:"return"`
}

// OpNum returns the operation number of RpcUnRegisterAsyncNotification operation.
func (o *xxx_UnregisterAsyncNotificationOperation) OpNum() int { return 3 }

// OpName returns the operation name of RpcUnRegisterAsyncNotification operation.
func (o *xxx_UnregisterAsyncNotificationOperation) OpName() string {
	return "/TermSrvNotification/v1/RpcUnRegisterAsyncNotification"
}

func (o *xxx_UnregisterAsyncNotificationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnregisterAsyncNotificationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phNotify {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NOTIFY_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Notify != nil {
			if err := o.Notify.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_UnregisterAsyncNotificationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phNotify {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NOTIFY_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Notify == nil {
			o.Notify = &Notify{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnregisterAsyncNotificationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnregisterAsyncNotificationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phNotify {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NOTIFY_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Notify != nil {
			if err := o.Notify.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_UnregisterAsyncNotificationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phNotify {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=NOTIFY_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Notify == nil {
			o.Notify = &Notify{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
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

// UnregisterAsyncNotificationRequest structure represents the RpcUnRegisterAsyncNotification operation request
type UnregisterAsyncNotificationRequest struct {
	// phNotify: The notification handle returned by RpcRegisterAsyncNotification. This
	// is of type NOTIFY_HANDLE.
	Notify *Notify `idl:"name:phNotify" json:"notify"`
}

func (o *UnregisterAsyncNotificationRequest) xxx_ToOp(ctx context.Context, op *xxx_UnregisterAsyncNotificationOperation) *xxx_UnregisterAsyncNotificationOperation {
	if op == nil {
		op = &xxx_UnregisterAsyncNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.Notify = o.Notify
	return op
}

func (o *UnregisterAsyncNotificationRequest) xxx_FromOp(ctx context.Context, op *xxx_UnregisterAsyncNotificationOperation) {
	if o == nil {
		return
	}
	o.Notify = op.Notify
}
func (o *UnregisterAsyncNotificationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *UnregisterAsyncNotificationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnregisterAsyncNotificationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MakeUnregisterAsyncNotificationRequest build a response structure from the given request structure.
func (o *UnregisterAsyncNotificationRequest) MakeResponse() *UnregisterAsyncNotificationResponse {
	return &UnregisterAsyncNotificationResponse{}
}

// OpNum returns the operation number of RpcUnRegisterAsyncNotification operation.
func (o *UnregisterAsyncNotificationRequest) OpNum() int { return 3 }

// OpName returns the operation name of RpcUnRegisterAsyncNotification operation.
func (o *UnregisterAsyncNotificationRequest) OpName() string {
	return "/TermSrvNotification/v1/RpcUnRegisterAsyncNotification"
}

// UnregisterAsyncNotificationResponse structure represents the RpcUnRegisterAsyncNotification operation response
type UnregisterAsyncNotificationResponse struct {
	// phNotify: The notification handle returned by RpcRegisterAsyncNotification. This
	// is of type NOTIFY_HANDLE.
	Notify *Notify `idl:"name:phNotify" json:"notify"`
	// Return: The RpcUnRegisterAsyncNotification return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *UnregisterAsyncNotificationResponse) xxx_ToOp(ctx context.Context, op *xxx_UnregisterAsyncNotificationOperation) *xxx_UnregisterAsyncNotificationOperation {
	if op == nil {
		op = &xxx_UnregisterAsyncNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.Notify = o.Notify
	op.Return = o.Return
	return op
}

func (o *UnregisterAsyncNotificationResponse) xxx_FromOp(ctx context.Context, op *xxx_UnregisterAsyncNotificationOperation) {
	if o == nil {
		return
	}
	o.Notify = op.Notify
	o.Return = op.Return
}
func (o *UnregisterAsyncNotificationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *UnregisterAsyncNotificationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnregisterAsyncNotificationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
