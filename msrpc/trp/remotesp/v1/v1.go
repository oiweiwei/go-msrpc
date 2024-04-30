package remotesp

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
	GoPackage = "trp"
)

var (
	// Syntax UUID
	RemotespSyntaxUUID = &uuid.UUID{TimeLow: 0x2f5f6521, TimeMid: 0xca47, TimeHiAndVersion: 0x1068, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0x19, Node: [6]uint8{0x0, 0xdd, 0x1, 0x6, 0x62, 0xdb}}
	// Syntax ID
	RemotespSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: RemotespSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// remotesp interface.
type RemotespClient interface {

	// The RemoteSPAttach method is called by the server to establish a binding instance
	// in response to a client call to the server's ClientAttach method.
	//
	// Return Values: The method returns 0 on success; otherwise, it returns a nonzero error
	// code, as specified in [MS-ERREF]. On success, the Server List is updated with the
	// binding instance.
	//
	// Exceptions Thrown:
	//
	// The client raises an RPC_S_ACCESS_DENIED exception if it fails to obtain the RPC
	// call attributes. The client also raises an RPC_S_ACCESS_DENIED exception if it determines
	// from the call attributes that the server did not specify RPC_C_AUTHN_LEVEL_PKT_PRIVACY,
	// and the client configuration requires this authentication level.
	//
	// Except as noted above, no exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// The opnum field value for this method is 0.<10>
	Attach(context.Context, *AttachRequest, ...dcerpc.CallOption) (*AttachResponse, error)

	// The RemoteSPEventProc method is called by the server to "push" completion notifications
	// and unsolicited events to the client. The server MUST call this method of the remotesp
	// interface with the endpoint and protocol sequence as specified by the connection-oriented
	// client in the ClientAttach RPC packet.
	//
	// Return Values: This method has no return values.
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	//
	// The opnum field value for this method is 1.
	EventProc(context.Context, *EventProcRequest, ...dcerpc.CallOption) (*EventProcResponse, error)

	// The RemoteSPDetach method is called by the server in response to a Client call to
	// the server's ClientDetach method to free the binding instance and to release the
	// associated resources.
	//
	// Exceptions Thrown:
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	//
	// The opnum field value for this method is 2.
	Detach(context.Context, *DetachRequest, ...dcerpc.CallOption) (*DetachResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error
}

// Type2 structure represents PCONTEXT_HANDLE_TYPE2 RPC structure.
type Type2 dcetypes.ContextHandle

func (o *Type2) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Type2) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Type2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Type2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

type xxx_DefaultRemotespClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultRemotespClient) Attach(ctx context.Context, in *AttachRequest, opts ...dcerpc.CallOption) (*AttachResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AttachResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRemotespClient) EventProc(ctx context.Context, in *EventProcRequest, opts ...dcerpc.CallOption) (*EventProcResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EventProcResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultRemotespClient) Detach(ctx context.Context, in *DetachRequest, opts ...dcerpc.CallOption) (*DetachResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DetachResponse{}
	out.xxx_FromOp(ctx, op)
	return out, nil
}

func (o *xxx_DefaultRemotespClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}
func NewRemotespClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RemotespClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RemotespSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultRemotespClient{cc: cc}, nil
}

// xxx_AttachOperation structure represents the RemoteSPAttach operation
type xxx_AttachOperation struct {
	Context *Type2 `idl:"name:pphContext" json:"context"`
	Return  int32  `idl:"name:Return" json:"return"`
}

func (o *xxx_AttachOperation) OpNum() int { return 0 }

func (o *xxx_AttachOperation) OpName() string { return "/remotesp/v1/RemoteSPAttach" }

func (o *xxx_AttachOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AttachOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_AttachOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_AttachOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AttachOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pphContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_TYPE2, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Type2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AttachOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pphContext {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_TYPE2, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Type2{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AttachRequest structure represents the RemoteSPAttach operation request
type AttachRequest struct {
}

func (o *AttachRequest) xxx_ToOp(ctx context.Context) *xxx_AttachOperation {
	if o == nil {
		return &xxx_AttachOperation{}
	}
	return &xxx_AttachOperation{}
}

func (o *AttachRequest) xxx_FromOp(ctx context.Context, op *xxx_AttachOperation) {
	if o == nil {
		return
	}
}
func (o *AttachRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AttachRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AttachOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AttachResponse structure represents the RemoteSPAttach operation response
type AttachResponse struct {
	// pphContext: Client handle of type PCONTEXT_HANDLE_TYPE2.
	Context *Type2 `idl:"name:pphContext" json:"context"`
	// Return: The RemoteSPAttach return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *AttachResponse) xxx_ToOp(ctx context.Context) *xxx_AttachOperation {
	if o == nil {
		return &xxx_AttachOperation{}
	}
	return &xxx_AttachOperation{
		Context: o.Context,
		Return:  o.Return,
	}
}

func (o *AttachResponse) xxx_FromOp(ctx context.Context, op *xxx_AttachOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.Return = op.Return
}
func (o *AttachResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AttachResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AttachOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EventProcOperation structure represents the RemoteSPEventProc operation
type xxx_EventProcOperation struct {
	Context *Type2 `idl:"name:phContext" json:"context"`
	Buffer  []byte `idl:"name:pBuffer;size_is:(lSize);length_is:(lSize)" json:"buffer"`
	Size    int32  `idl:"name:lSize" json:"size"`
}

func (o *xxx_EventProcOperation) OpNum() int { return 1 }

func (o *xxx_EventProcOperation) OpName() string { return "/remotesp/v1/RemoteSPEventProc" }

func (o *xxx_EventProcOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Buffer != nil && o.Size == 0 {
		o.Size = int32(len(o.Buffer))
	}
	if o.Buffer != nil && o.Size == 0 {
		o.Size = int32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EventProcOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phContext {in} (1:{context_handle, alias=PCONTEXT_HANDLE_TYPE2, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Type2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pBuffer {in} (1:[dim:0,size_is=lSize,length_is=lSize](uchar))
	{
		dimSize1 := uint64(o.Size)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.Size)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
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
	}
	// lSize {in} (1:(int32))
	{
		if err := w.WriteData(o.Size); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EventProcOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phContext {in} (1:{context_handle, alias=PCONTEXT_HANDLE_TYPE2, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Type2{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pBuffer {in} (1:[dim:0,size_is=lSize,length_is=lSize](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
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
	}
	// lSize {in} (1:(int32))
	{
		if err := w.ReadData(&o.Size); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EventProcOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EventProcOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_EventProcOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	return nil
}

// EventProcRequest structure represents the RemoteSPEventProc operation request
type EventProcRequest struct {
	// phContext: Client handle of type PCONTEXT_HANDLE_TYPE2.
	Context *Type2 `idl:"name:phContext" json:"context"`
	// pBuffer: Packet MUST contain a list of ASYNCEVENTMSG structures, each of which MUST
	// be ASYNCEVENTMSG.TotalSize bytes in size.
	Buffer []byte `idl:"name:pBuffer;size_is:(lSize);length_is:(lSize)" json:"buffer"`
	// lSize: Size of the pBuffer.
	Size int32 `idl:"name:lSize" json:"size"`
}

func (o *EventProcRequest) xxx_ToOp(ctx context.Context) *xxx_EventProcOperation {
	if o == nil {
		return &xxx_EventProcOperation{}
	}
	return &xxx_EventProcOperation{
		Context: o.Context,
		Buffer:  o.Buffer,
		Size:    o.Size,
	}
}

func (o *EventProcRequest) xxx_FromOp(ctx context.Context, op *xxx_EventProcOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
	o.Buffer = op.Buffer
	o.Size = op.Size
}
func (o *EventProcRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EventProcRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EventProcOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EventProcResponse structure represents the RemoteSPEventProc operation response
type EventProcResponse struct {
}

func (o *EventProcResponse) xxx_ToOp(ctx context.Context) *xxx_EventProcOperation {
	if o == nil {
		return &xxx_EventProcOperation{}
	}
	return &xxx_EventProcOperation{}
}

func (o *EventProcResponse) xxx_FromOp(ctx context.Context, op *xxx_EventProcOperation) {
	if o == nil {
		return
	}
}
func (o *EventProcResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EventProcResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EventProcOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DetachOperation structure represents the RemoteSPDetach operation
type xxx_DetachOperation struct {
	Context *Type2 `idl:"name:pphContext" json:"context"`
}

func (o *xxx_DetachOperation) OpNum() int { return 2 }

func (o *xxx_DetachOperation) OpName() string { return "/remotesp/v1/RemoteSPDetach" }

func (o *xxx_DetachOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DetachOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_TYPE2, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Type2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DetachOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_TYPE2, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Type2{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DetachOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DetachOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_TYPE2, names=ndr_context_handle}(struct))
	{
		if o.Context != nil {
			if err := o.Context.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Type2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DetachOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pphContext {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=PCONTEXT_HANDLE_TYPE2, names=ndr_context_handle}(struct))
	{
		if o.Context == nil {
			o.Context = &Type2{}
		}
		if err := o.Context.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// DetachRequest structure represents the RemoteSPDetach operation request
type DetachRequest struct {
	// pphContext: Pointer to a PCONTEXT_HANDLE_TYPE2 handle to the binding instance being
	// terminated.
	//
	// This method has no return values.
	//
	// On success, the binding instance is removed from the Server List.
	Context *Type2 `idl:"name:pphContext" json:"context"`
}

func (o *DetachRequest) xxx_ToOp(ctx context.Context) *xxx_DetachOperation {
	if o == nil {
		return &xxx_DetachOperation{}
	}
	return &xxx_DetachOperation{
		Context: o.Context,
	}
}

func (o *DetachRequest) xxx_FromOp(ctx context.Context, op *xxx_DetachOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
}
func (o *DetachRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DetachRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DetachOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DetachResponse structure represents the RemoteSPDetach operation response
type DetachResponse struct {
	// pphContext: Pointer to a PCONTEXT_HANDLE_TYPE2 handle to the binding instance being
	// terminated.
	//
	// This method has no return values.
	//
	// On success, the binding instance is removed from the Server List.
	Context *Type2 `idl:"name:pphContext" json:"context"`
}

func (o *DetachResponse) xxx_ToOp(ctx context.Context) *xxx_DetachOperation {
	if o == nil {
		return &xxx_DetachOperation{}
	}
	return &xxx_DetachOperation{
		Context: o.Context,
	}
}

func (o *DetachResponse) xxx_FromOp(ctx context.Context, op *xxx_DetachOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
}
func (o *DetachResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DetachResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DetachOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
