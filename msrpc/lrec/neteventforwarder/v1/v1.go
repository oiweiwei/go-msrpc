package neteventforwarder

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcetypes "github.com/oiweiwei/go-msrpc/msrpc/dcetypes"
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
	_ = dcetypes.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "lrec"
)

var (
	// Syntax UUID
	NetEventForwarderSyntaxUUID = &uuid.UUID{TimeLow: 0x22e5386d, TimeMid: 0x8b12, TimeHiAndVersion: 0x4bf0, ClockSeqHiAndReserved: 0xb0, ClockSeqLow: 0xec, Node: [6]uint8{0x6a, 0x1e, 0xa4, 0x19, 0xe3, 0x66}}
	// Syntax ID
	NetEventForwarderSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: NetEventForwarderSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// NetEventForwarder interface.
type NetEventForwarderClient interface {

	// RpcNetEventOpenSession operation.
	OpenSession(context.Context, *OpenSessionRequest, ...dcerpc.CallOption) (*OpenSessionResponse, error)

	// RpcNetEventReceiveData operation.
	ReceiveData(context.Context, *ReceiveDataRequest, ...dcerpc.CallOption) (*ReceiveDataResponse, error)

	// RpcNetEventCloseSession operation.
	CloseSession(context.Context, *CloseSessionRequest, ...dcerpc.CallOption) (*CloseSessionResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// Session structure represents PSESSION_HANDLE RPC structure.
type Session dcetypes.ContextHandle

func (o *Session) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Session) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Session) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Session) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

// EventBuffer structure represents EVENT_BUFFER RPC structure.
//
// The EVENT_BUFFER structure defines a data structure for transferring a generic payload.
// The LREC protocol uses this structure to pass event records in the RpcNetEventReceiveData
// (section 3.1.4.2.2) method.
type EventBuffer struct {
	// BufferLength: This property specifies the length, in bytes, of the data stored in
	// the Buffer field.
	BufferLength uint32 `idl:"name:BufferLength" json:"buffer_length"`
	// Buffer: This property specifies a collection of one or more NET_EVENT_DATA_HEADER
	// (section 2.3.2.2) structures each followed by an event payload.
	Buffer []byte `idl:"name:Buffer;size_is:(BufferLength)" json:"buffer"`
}

func (o *EventBuffer) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.BufferLength == 0 {
		o.BufferLength = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *EventBuffer) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.BufferLength); err != nil {
		return err
	}
	if o.Buffer != nil || o.BufferLength > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.BufferLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Buffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Buffer[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *EventBuffer) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.BufferLength); err != nil {
		return err
	}
	_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.BufferLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.BufferLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]byte, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if err := w.ReadData(&o.Buffer[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultNetEventForwarderClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultNetEventForwarderClient) OpenSession(ctx context.Context, in *OpenSessionRequest, opts ...dcerpc.CallOption) (*OpenSessionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenSessionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetEventForwarderClient) ReceiveData(ctx context.Context, in *ReceiveDataRequest, opts ...dcerpc.CallOption) (*ReceiveDataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ReceiveDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetEventForwarderClient) CloseSession(ctx context.Context, in *CloseSessionRequest, opts ...dcerpc.CallOption) (*CloseSessionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseSessionResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultNetEventForwarderClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultNetEventForwarderClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewNetEventForwarderClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (NetEventForwarderClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(NetEventForwarderSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultNetEventForwarderClient{cc: cc}, nil
}

// xxx_OpenSessionOperation structure represents the RpcNetEventOpenSession operation
type xxx_OpenSessionOperation struct {
	LoggerName    string   `idl:"name:LoggerName;string" json:"logger_name"`
	SessionHandle *Session `idl:"name:SessionHandle" json:"session_handle"`
	Return        uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenSessionOperation) OpNum() int { return 0 }

func (o *xxx_OpenSessionOperation) OpName() string {
	return "/NetEventForwarder/v1/RpcNetEventOpenSession"
}

func (o *xxx_OpenSessionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenSessionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// LoggerName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.LoggerName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenSessionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// LoggerName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.LoggerName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenSessionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenSessionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// SessionHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PSESSION_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.SessionHandle != nil {
			if err := o.SessionHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Session{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenSessionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// SessionHandle {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PSESSION_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.SessionHandle == nil {
			o.SessionHandle = &Session{}
		}
		if err := o.SessionHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenSessionRequest structure represents the RpcNetEventOpenSession operation request
type OpenSessionRequest struct {
	LoggerName string `idl:"name:LoggerName;string" json:"logger_name"`
}

func (o *OpenSessionRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenSessionOperation) *xxx_OpenSessionOperation {
	if op == nil {
		op = &xxx_OpenSessionOperation{}
	}
	if o == nil {
		return op
	}
	op.LoggerName = o.LoggerName
	return op
}

func (o *OpenSessionRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenSessionOperation) {
	if o == nil {
		return
	}
	o.LoggerName = op.LoggerName
}
func (o *OpenSessionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenSessionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenSessionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenSessionResponse structure represents the RpcNetEventOpenSession operation response
type OpenSessionResponse struct {
	SessionHandle *Session `idl:"name:SessionHandle" json:"session_handle"`
	// Return: The RpcNetEventOpenSession return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *OpenSessionResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenSessionOperation) *xxx_OpenSessionOperation {
	if op == nil {
		op = &xxx_OpenSessionOperation{}
	}
	if o == nil {
		return op
	}
	op.SessionHandle = o.SessionHandle
	op.Return = o.Return
	return op
}

func (o *OpenSessionResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenSessionOperation) {
	if o == nil {
		return
	}
	o.SessionHandle = op.SessionHandle
	o.Return = op.Return
}
func (o *OpenSessionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenSessionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenSessionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReceiveDataOperation structure represents the RpcNetEventReceiveData operation
type xxx_ReceiveDataOperation struct {
	SessionHandle *Session     `idl:"name:SessionHandle" json:"session_handle"`
	EventBuffer   *EventBuffer `idl:"name:EventBuffer" json:"event_buffer"`
	Return        uint32       `idl:"name:Return" json:"return"`
}

func (o *xxx_ReceiveDataOperation) OpNum() int { return 1 }

func (o *xxx_ReceiveDataOperation) OpName() string {
	return "/NetEventForwarder/v1/RpcNetEventReceiveData"
}

func (o *xxx_ReceiveDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SessionHandle {in} (1:{context_handle, alias=PSESSION_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.SessionHandle != nil {
			if err := o.SessionHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Session{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ReceiveDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SessionHandle {in} (1:{context_handle, alias=PSESSION_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.SessionHandle == nil {
			o.SessionHandle = &Session{}
		}
		if err := o.SessionHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// EventBuffer {out} (1:{pointer=ref}*(1))(2:{alias=EVENT_BUFFER}(struct))
	{
		if o.EventBuffer != nil {
			if err := o.EventBuffer.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&EventBuffer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReceiveDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// EventBuffer {out} (1:{pointer=ref}*(1))(2:{alias=EVENT_BUFFER}(struct))
	{
		if o.EventBuffer == nil {
			o.EventBuffer = &EventBuffer{}
		}
		if err := o.EventBuffer.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ReceiveDataRequest structure represents the RpcNetEventReceiveData operation request
type ReceiveDataRequest struct {
	SessionHandle *Session `idl:"name:SessionHandle" json:"session_handle"`
}

func (o *ReceiveDataRequest) xxx_ToOp(ctx context.Context, op *xxx_ReceiveDataOperation) *xxx_ReceiveDataOperation {
	if op == nil {
		op = &xxx_ReceiveDataOperation{}
	}
	if o == nil {
		return op
	}
	op.SessionHandle = o.SessionHandle
	return op
}

func (o *ReceiveDataRequest) xxx_FromOp(ctx context.Context, op *xxx_ReceiveDataOperation) {
	if o == nil {
		return
	}
	o.SessionHandle = op.SessionHandle
}
func (o *ReceiveDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ReceiveDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReceiveDataResponse structure represents the RpcNetEventReceiveData operation response
type ReceiveDataResponse struct {
	EventBuffer *EventBuffer `idl:"name:EventBuffer" json:"event_buffer"`
	// Return: The RpcNetEventReceiveData return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ReceiveDataResponse) xxx_ToOp(ctx context.Context, op *xxx_ReceiveDataOperation) *xxx_ReceiveDataOperation {
	if op == nil {
		op = &xxx_ReceiveDataOperation{}
	}
	if o == nil {
		return op
	}
	op.EventBuffer = o.EventBuffer
	op.Return = o.Return
	return op
}

func (o *ReceiveDataResponse) xxx_FromOp(ctx context.Context, op *xxx_ReceiveDataOperation) {
	if o == nil {
		return
	}
	o.EventBuffer = op.EventBuffer
	o.Return = op.Return
}
func (o *ReceiveDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ReceiveDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReceiveDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseSessionOperation structure represents the RpcNetEventCloseSession operation
type xxx_CloseSessionOperation struct {
	SessionHandle *Session `idl:"name:SessionHandle" json:"session_handle"`
}

func (o *xxx_CloseSessionOperation) OpNum() int { return 2 }

func (o *xxx_CloseSessionOperation) OpName() string {
	return "/NetEventForwarder/v1/RpcNetEventCloseSession"
}

func (o *xxx_CloseSessionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseSessionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// SessionHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PSESSION_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.SessionHandle != nil {
			if err := o.SessionHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Session{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseSessionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// SessionHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PSESSION_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.SessionHandle == nil {
			o.SessionHandle = &Session{}
		}
		if err := o.SessionHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseSessionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseSessionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// SessionHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PSESSION_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.SessionHandle != nil {
			if err := o.SessionHandle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Session{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseSessionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// SessionHandle {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PSESSION_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.SessionHandle == nil {
			o.SessionHandle = &Session{}
		}
		if err := o.SessionHandle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// CloseSessionRequest structure represents the RpcNetEventCloseSession operation request
type CloseSessionRequest struct {
	SessionHandle *Session `idl:"name:SessionHandle" json:"session_handle"`
}

func (o *CloseSessionRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseSessionOperation) *xxx_CloseSessionOperation {
	if op == nil {
		op = &xxx_CloseSessionOperation{}
	}
	if o == nil {
		return op
	}
	op.SessionHandle = o.SessionHandle
	return op
}

func (o *CloseSessionRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseSessionOperation) {
	if o == nil {
		return
	}
	o.SessionHandle = op.SessionHandle
}
func (o *CloseSessionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloseSessionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseSessionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseSessionResponse structure represents the RpcNetEventCloseSession operation response
type CloseSessionResponse struct {
	SessionHandle *Session `idl:"name:SessionHandle" json:"session_handle"`
}

func (o *CloseSessionResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseSessionOperation) *xxx_CloseSessionOperation {
	if op == nil {
		op = &xxx_CloseSessionOperation{}
	}
	if o == nil {
		return op
	}
	op.SessionHandle = o.SessionHandle
	return op
}

func (o *CloseSessionResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseSessionOperation) {
	if o == nil {
		return
	}
	o.SessionHandle = op.SessionHandle
}
func (o *CloseSessionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloseSessionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseSessionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
