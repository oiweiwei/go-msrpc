package faxclient

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	fax "github.com/oiweiwei/go-msrpc/msrpc/fax"
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
	_ = fax.GoPackage
)

var (
	// import guard
	GoPackage = "fax"
)

var (
	// Syntax UUID
	FaxclientSyntaxUUID = &uuid.UUID{TimeLow: 0x6099fc12, TimeMid: 0x3eff, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0xd0, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xd9, 0x1a, 0x4e}}
	// Syntax ID
	FaxclientSyntaxV3_0 = &dcerpc.SyntaxID{IfUUID: FaxclientSyntaxUUID, IfVersionMajor: 3, IfVersionMinor: 0}
)

// faxclient interface.
type FaxclientClient interface {

	// The FAX_OpenConnection (Opnum 0) method returns the context handle that is supplied
	// by the FAX_StartServerNotification family of calls. This is done to provide a security
	// layer, by verifying that the notifications are coming from an expected source.
	//
	// Exceptions Thrown: No exceptions are thrown except those thrown by the underlying
	// RPC protocol, [MS-RPCE].
	//
	// The returned PRPC_FAX_HANDLE is the Context parameter cast to a HANDLE.
	//
	// The FAX_OpenConnection method returns the context handle supplied by the FAX_StartServerNotification
	// family of calls. This is done to provide a security layer, by verifying that the
	// notifications are coming from an expected source.
	OpenConnection(context.Context, *OpenConnectionRequest, ...dcerpc.CallOption) (*OpenConnectionResponse, error)

	// The FAX_ClientEventQueue (Opnum 1) method is called by the fax server (acting as
	// an RPC client for this call) when it needs to deliver a legacy fax event to the fax
	// client (acting as an RPC server for this call). The fax client registers for notifications
	// of legacy events with the fax server by calling FAX_StartServerNotification (section
	// 3.1.4.1.100) or FAX_StartServerNotificationEx (section 3.1.4.1.101). In this call,
	// the fax client MUST pass a fax client notification context, which the fax server
	// MUST pass back to the fax client when it sends an event. This is done to provide
	// a security layer, by verifying that the notifications are coming from an expected
	// source.
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	ClientEventQueue(context.Context, *ClientEventQueueRequest, ...dcerpc.CallOption) (*ClientEventQueueResponse, error)

	// The FAX_CloseConnection (Opnum 2) method is called by the fax server (acting as an
	// RPC client for this call) when it needs to release the connection to the fax client
	// (acting as an RPC server for this call). When the fax client calls FAX_EndServerNotification
	// (section 3.1.4.1.17), the fax server MUST release the RPC connection to the fax client
	// through this call.
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	CloseConnection(context.Context, *CloseConnectionRequest, ...dcerpc.CallOption) (*CloseConnectionResponse, error)

	// The FAX_ClientEventQueueEx (Opnum 3) method is called by the fax server (acting as
	// an RPC client for this call) when it needs to deliver an extended fax event to the
	// fax client (acting as an RPC server for this call). The fax client registers for
	// notifications with the fax server by calling either FAX_StartServerNotificationEx
	// (section 3.1.4.1.101) or FAX_StartServerNotificationEx2 (section 3.1.4.1.102). In
	// this call, the fax client MUST pass a fax client notification context, which the
	// fax server MUST pass back to the fax client when it sends an event. This is done
	// to provide a security layer, by verifying that the notifications are coming from
	// an expected source.
	//
	// Data in FAX_ClientEventQueueEx is serialized. Pointers to variable-size data (such
	// as strings) are replaced with offsets from the beginning of the buffer.
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	//
	// Data in FAX_ClientEventQueueEx is serialized. Pointers to variable size data (such
	// as strings) are replaced with offsets from the beginning of the buffer.
	ClientEventQueueEx(context.Context, *ClientEventQueueExRequest, ...dcerpc.CallOption) (*ClientEventQueueExResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error
}

type xxx_DefaultFaxclientClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultFaxclientClient) OpenConnection(ctx context.Context, in *OpenConnectionRequest, opts ...dcerpc.CallOption) (*OpenConnectionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenConnectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxclientClient) ClientEventQueue(ctx context.Context, in *ClientEventQueueRequest, opts ...dcerpc.CallOption) (*ClientEventQueueResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ClientEventQueueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxclientClient) CloseConnection(ctx context.Context, in *CloseConnectionRequest, opts ...dcerpc.CallOption) (*CloseConnectionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseConnectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxclientClient) ClientEventQueueEx(ctx context.Context, in *ClientEventQueueExRequest, opts ...dcerpc.CallOption) (*ClientEventQueueExResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ClientEventQueueExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxclientClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}
func NewFaxclientClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (FaxclientClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(FaxclientSyntaxV3_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultFaxclientClient{cc: cc}, nil
}

// xxx_OpenConnectionOperation structure represents the FAX_OpenConnection operation
type xxx_OpenConnectionOperation struct {
	Context uint64   `idl:"name:Context" json:"context"`
	Fax     *fax.Fax `idl:"name:FaxHandle" json:"fax"`
	Return  uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenConnectionOperation) OpNum() int { return 0 }

func (o *xxx_OpenConnectionOperation) OpName() string { return "/faxclient/v3/FAX_OpenConnection" }

func (o *xxx_OpenConnectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenConnectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Context {in} (1:(uint64))
	{
		if err := w.WriteData(o.Context); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenConnectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Context {in} (1:(uint64))
	{
		if err := w.ReadData(&o.Context); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenConnectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenConnectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// FaxHandle {out} (1:{pointer=ref, alias=PRPC_FAX_HANDLE}*(1))(2:{context_handle, alias=RPC_FAX_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Fax != nil {
			if err := o.Fax.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.Fax{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenConnectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// FaxHandle {out} (1:{pointer=ref, alias=PRPC_FAX_HANDLE}*(1))(2:{context_handle, alias=RPC_FAX_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Fax == nil {
			o.Fax = &fax.Fax{}
		}
		if err := o.Fax.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenConnectionRequest structure represents the FAX_OpenConnection operation request
type OpenConnectionRequest struct {
	// Context: A ULONG64 ([MS-DTYP] section 2.2.51) containing a context information handle.
	// This handle SHOULD match the one supplied to the server when using the FAX_StartServerNotification
	// family of calls. For more information, see the following topics:
	//
	// §  FAX_StartServerNotification
	//
	// * *FAX_StartServerNotificationEx*
	//
	// * *FAX_StartServerNotificationEx2*
	Context uint64 `idl:"name:Context" json:"context"`
}

func (o *OpenConnectionRequest) xxx_ToOp(ctx context.Context) *xxx_OpenConnectionOperation {
	if o == nil {
		return &xxx_OpenConnectionOperation{}
	}
	return &xxx_OpenConnectionOperation{
		Context: o.Context,
	}
}

func (o *OpenConnectionRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenConnectionOperation) {
	if o == nil {
		return
	}
	o.Context = op.Context
}
func (o *OpenConnectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OpenConnectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenConnectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenConnectionResponse structure represents the FAX_OpenConnection operation response
type OpenConnectionResponse struct {
	// FaxHandle: A pointer to an RPC_FAX_HANDLE indicating a context handle to open. This
	// value is used in other fax client calls.
	//
	// This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise, it MUST
	// return one of the following error codes, one of the fax-specific errors that are
	// defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | The method requires at least packet-level privacy. The server                    |
	//	|                                    | checks the authentication level of the client. If it is less than                |
	//	|                                    | RPC_C_AUTHN_LEVEL_PKT_PRIVACY, refuse access. Or there are other access-related  |
	//	|                                    | problems.                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | An invalid AssyncInfo structure is pointed to by the Context parameter or there  |
	//	|                                    | are parameter-related problems.                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	Fax *fax.Fax `idl:"name:FaxHandle" json:"fax"`
	// Return: The FAX_OpenConnection return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *OpenConnectionResponse) xxx_ToOp(ctx context.Context) *xxx_OpenConnectionOperation {
	if o == nil {
		return &xxx_OpenConnectionOperation{}
	}
	return &xxx_OpenConnectionOperation{
		Fax:    o.Fax,
		Return: o.Return,
	}
}

func (o *OpenConnectionResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenConnectionOperation) {
	if o == nil {
		return
	}
	o.Fax = op.Fax
	o.Return = op.Return
}
func (o *OpenConnectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OpenConnectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenConnectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ClientEventQueueOperation structure represents the FAX_ClientEventQueue operation
type xxx_ClientEventQueueOperation struct {
	FaxPort  *fax.Fax             `idl:"name:FaxPortHandle" json:"fax_port"`
	FaxEvent *fax.CompletionEvent `idl:"name:FaxEvent" json:"fax_event"`
	Return   uint32               `idl:"name:Return" json:"return"`
}

func (o *xxx_ClientEventQueueOperation) OpNum() int { return 1 }

func (o *xxx_ClientEventQueueOperation) OpName() string { return "/faxclient/v3/FAX_ClientEventQueue" }

func (o *xxx_ClientEventQueueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientEventQueueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FaxPortHandle {in} (1:{context_handle, alias=RPC_FAX_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort != nil {
			if err := o.FaxPort.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.Fax{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// FaxEvent {in} (1:{alias=FAX_EVENT}(struct))
	{
		if o.FaxEvent != nil {
			if err := o.FaxEvent.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.CompletionEvent{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ClientEventQueueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FaxPortHandle {in} (1:{context_handle, alias=RPC_FAX_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort == nil {
			o.FaxPort = &fax.Fax{}
		}
		if err := o.FaxPort.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// FaxEvent {in} (1:{alias=FAX_EVENT}(struct))
	{
		if o.FaxEvent == nil {
			o.FaxEvent = &fax.CompletionEvent{}
		}
		if err := o.FaxEvent.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientEventQueueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientEventQueueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientEventQueueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ClientEventQueueRequest structure represents the FAX_ClientEventQueue operation request
type ClientEventQueueRequest struct {
	// FaxPortHandle: A fax data type that indicates a context handle for this call.
	FaxPort *fax.Fax `idl:"name:FaxPortHandle" json:"fax_port"`
	// FaxEvent: A FAX_EVENT (section 2.2.66) structure that contains the contents of an
	// I/O completion packet. The fax server sends the completion packet to notify a fax
	// client application about an asynchronous fax server event. Since the client is to
	// be notified of each event separately, in this case ORing of events is not allowed.
	//
	// This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise, it MUST
	// return one of the following error codes, one of the fax-specific errors that are
	// defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------+
	//	|               RETURN               |                                                          |
	//	|             VALUE/CODE             |                       DESCRIPTION                        |
	//	|                                    |                                                          |
	//	+------------------------------------+----------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | Not enough storage is available to process this command. |
	//	+------------------------------------+----------------------------------------------------------+
	FaxEvent *fax.CompletionEvent `idl:"name:FaxEvent" json:"fax_event"`
}

func (o *ClientEventQueueRequest) xxx_ToOp(ctx context.Context) *xxx_ClientEventQueueOperation {
	if o == nil {
		return &xxx_ClientEventQueueOperation{}
	}
	return &xxx_ClientEventQueueOperation{
		FaxPort:  o.FaxPort,
		FaxEvent: o.FaxEvent,
	}
}

func (o *ClientEventQueueRequest) xxx_FromOp(ctx context.Context, op *xxx_ClientEventQueueOperation) {
	if o == nil {
		return
	}
	o.FaxPort = op.FaxPort
	o.FaxEvent = op.FaxEvent
}
func (o *ClientEventQueueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ClientEventQueueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClientEventQueueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ClientEventQueueResponse structure represents the FAX_ClientEventQueue operation response
type ClientEventQueueResponse struct {
	// Return: The FAX_ClientEventQueue return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ClientEventQueueResponse) xxx_ToOp(ctx context.Context) *xxx_ClientEventQueueOperation {
	if o == nil {
		return &xxx_ClientEventQueueOperation{}
	}
	return &xxx_ClientEventQueueOperation{
		Return: o.Return,
	}
}

func (o *ClientEventQueueResponse) xxx_FromOp(ctx context.Context, op *xxx_ClientEventQueueOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ClientEventQueueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ClientEventQueueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClientEventQueueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseConnectionOperation structure represents the FAX_CloseConnection operation
type xxx_CloseConnectionOperation struct {
	Fax    *fax.Fax `idl:"name:FaxHandle" json:"fax"`
	Return uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseConnectionOperation) OpNum() int { return 2 }

func (o *xxx_CloseConnectionOperation) OpName() string { return "/faxclient/v3/FAX_CloseConnection" }

func (o *xxx_CloseConnectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseConnectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FaxHandle {in, out} (1:{pointer=ref, alias=PRPC_FAX_HANDLE}*(1))(2:{context_handle, alias=RPC_FAX_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Fax != nil {
			if err := o.Fax.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.Fax{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseConnectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FaxHandle {in, out} (1:{pointer=ref, alias=PRPC_FAX_HANDLE}*(1))(2:{context_handle, alias=RPC_FAX_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Fax == nil {
			o.Fax = &fax.Fax{}
		}
		if err := o.Fax.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseConnectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseConnectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// FaxHandle {in, out} (1:{pointer=ref, alias=PRPC_FAX_HANDLE}*(1))(2:{context_handle, alias=RPC_FAX_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Fax != nil {
			if err := o.Fax.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.Fax{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseConnectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// FaxHandle {in, out} (1:{pointer=ref, alias=PRPC_FAX_HANDLE}*(1))(2:{context_handle, alias=RPC_FAX_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Fax == nil {
			o.Fax = &fax.Fax{}
		}
		if err := o.Fax.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloseConnectionRequest structure represents the FAX_CloseConnection operation request
type CloseConnectionRequest struct {
	// FaxHandle: A pointer to an RPC_FAX_HANDLE that indicates a context handle to close.
	// For more information about RPC_FAX_HANDLE, see FAX Data Types (section 2.2.74).
	Fax *fax.Fax `idl:"name:FaxHandle" json:"fax"`
}

func (o *CloseConnectionRequest) xxx_ToOp(ctx context.Context) *xxx_CloseConnectionOperation {
	if o == nil {
		return &xxx_CloseConnectionOperation{}
	}
	return &xxx_CloseConnectionOperation{
		Fax: o.Fax,
	}
}

func (o *CloseConnectionRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseConnectionOperation) {
	if o == nil {
		return
	}
	o.Fax = op.Fax
}
func (o *CloseConnectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CloseConnectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseConnectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseConnectionResponse structure represents the FAX_CloseConnection operation response
type CloseConnectionResponse struct {
	// FaxHandle: A pointer to an RPC_FAX_HANDLE that indicates a context handle to close.
	// For more information about RPC_FAX_HANDLE, see FAX Data Types (section 2.2.74).
	Fax *fax.Fax `idl:"name:FaxHandle" json:"fax"`
	// Return: The FAX_CloseConnection return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CloseConnectionResponse) xxx_ToOp(ctx context.Context) *xxx_CloseConnectionOperation {
	if o == nil {
		return &xxx_CloseConnectionOperation{}
	}
	return &xxx_CloseConnectionOperation{
		Fax:    o.Fax,
		Return: o.Return,
	}
}

func (o *CloseConnectionResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseConnectionOperation) {
	if o == nil {
		return
	}
	o.Fax = op.Fax
	o.Return = op.Return
}
func (o *CloseConnectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CloseConnectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseConnectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ClientEventQueueExOperation structure represents the FAX_ClientEventQueueEx operation
type xxx_ClientEventQueueExOperation struct {
	ClientContext *fax.Fax `idl:"name:hClientContext;pointer:ref" json:"client_context"`
	Data          []byte   `idl:"name:lpbData;size_is:(dwDataSize);pointer:ref" json:"data"`
	DataSize      uint32   `idl:"name:dwDataSize" json:"data_size"`
	Return        uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_ClientEventQueueExOperation) OpNum() int { return 3 }

func (o *xxx_ClientEventQueueExOperation) OpName() string {
	return "/faxclient/v3/FAX_ClientEventQueueEx"
}

func (o *xxx_ClientEventQueueExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Data != nil && o.DataSize == 0 {
		o.DataSize = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientEventQueueExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hClientContext {in} (1:{context_handle, pointer=ref, alias=RPC_FAX_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ClientContext != nil {
			if err := o.ClientContext.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.Fax{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpbData {in} (1:{pointer=ref, alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=dwDataSize](uchar))
	{
		dimSize1 := uint64(o.DataSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Data {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Data[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// dwDataSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.DataSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientEventQueueExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hClientContext {in} (1:{context_handle, pointer=ref, alias=RPC_FAX_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.ClientContext == nil {
			o.ClientContext = &fax.Fax{}
		}
		if err := o.ClientContext.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpbData {in} (1:{pointer=ref, alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=dwDataSize](uchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]byte, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
		}
	}
	// dwDataSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DataSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientEventQueueExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientEventQueueExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClientEventQueueExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ClientEventQueueExRequest structure represents the FAX_ClientEventQueueEx operation request
type ClientEventQueueExRequest struct {
	// hClientContext: A fax data type indicating a context handle for this call.
	ClientContext *fax.Fax `idl:"name:hClientContext;pointer:ref" json:"client_context"`
	// lpbData: A pointer to a FAX_EVENT_EX (section 2.2.67) or FAX_EVENT_EX_1 (section
	// 2.2.68) structure. The data is serialized. Pointers to variable size data (such as
	// strings) are replaced with offsets from the beginning of the buffer. Since the client
	// is to be notified of each event separately, in this case ORing of events is not allowed.
	//
	// If the client requested extended events by calling FAX_StartServerNotificationEx,
	// the client MUST use a FAX_EVENT_EX. If the client called FAX_StartServerNotificationEx2
	// to receive these events, the client MUST use a FAX_EVENT_EX_1.
	Data []byte `idl:"name:lpbData;size_is:(dwDataSize);pointer:ref" json:"data"`
	// dwDataSize: A DWORD ([MS-DTYP] section 2.2.9) containing the size of the buffer pointed
	// to by the lpbData parameter.
	//
	// This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise, it MUST
	// return one of the following error codes, one of the fax-specific errors that are
	// defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D   | The hClientContext handle is not a valid subscription context handle returned by |
	//	|                                 | FAX_StartServerNotificationEx or FAX_StartServerNotificationEx2.<218>            |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_OUTOFMEMORY 0x0000000E    | The fax client needs to make a copy of the data provided by the fax server in    |
	//	|                                 | the lpbData buffer, and the fax client failed to allocate dwDataSize bytes to    |
	//	|                                 | hold this copy.                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F | The fax client failed to recognize the custom marshaled FAX_EVENT_EX or          |
	//	|                                 | FAX_EVENT_EX_1 provided by the fax server in the lpbData buffer.                 |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	DataSize uint32 `idl:"name:dwDataSize" json:"data_size"`
}

func (o *ClientEventQueueExRequest) xxx_ToOp(ctx context.Context) *xxx_ClientEventQueueExOperation {
	if o == nil {
		return &xxx_ClientEventQueueExOperation{}
	}
	return &xxx_ClientEventQueueExOperation{
		ClientContext: o.ClientContext,
		Data:          o.Data,
		DataSize:      o.DataSize,
	}
}

func (o *ClientEventQueueExRequest) xxx_FromOp(ctx context.Context, op *xxx_ClientEventQueueExOperation) {
	if o == nil {
		return
	}
	o.ClientContext = op.ClientContext
	o.Data = op.Data
	o.DataSize = op.DataSize
}
func (o *ClientEventQueueExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ClientEventQueueExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClientEventQueueExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ClientEventQueueExResponse structure represents the FAX_ClientEventQueueEx operation response
type ClientEventQueueExResponse struct {
	// Return: The FAX_ClientEventQueueEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ClientEventQueueExResponse) xxx_ToOp(ctx context.Context) *xxx_ClientEventQueueExOperation {
	if o == nil {
		return &xxx_ClientEventQueueExOperation{}
	}
	return &xxx_ClientEventQueueExOperation{
		Return: o.Return,
	}
}

func (o *ClientEventQueueExResponse) xxx_FromOp(ctx context.Context, op *xxx_ClientEventQueueExOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ClientEventQueueExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ClientEventQueueExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClientEventQueueExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
