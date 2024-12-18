package asyncemsmdb

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	oxcrpc "github.com/oiweiwei/go-msrpc/msrpc/oxcrpc"
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
	_ = oxcrpc.GoPackage
)

var (
	// import guard
	GoPackage = "oxcrpc"
)

var (
	// Syntax UUID
	AsyncemsmdbSyntaxUUID = &uuid.UUID{TimeLow: 0x5261574a, TimeMid: 0x4572, TimeHiAndVersion: 0x206e, ClockSeqHiAndReserved: 0xb2, ClockSeqLow: 0x68, Node: [6]uint8{0x6b, 0x19, 0x92, 0x13, 0xb4, 0xe4}}
	// Syntax ID
	AsyncemsmdbSyntaxV0_1 = &dcerpc.SyntaxID{IfUUID: AsyncemsmdbSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 1}
)

// asyncemsmdb interface.
type AsyncemsmdbClient interface {

	// The EcDoAsyncWaitEx method is an asynchronous call that the server does not complete
	// until events are pending on the Session Context, up to a 5-minute duration of no
	// client activity. If no events are available within 5 minutes of the time that the
	// client last accessed the server<36> through a call to the EcDoRpcExt2 method, as
	// specified in section 3.1.4.2, the server returns the call and does not set the NotificationPending
	// flag in the pulFlagsOut parameter. If an event is pending, the server completes the
	// call immediately and returns the NotificationPending flag in the pulFlagsOut parameter.
	// This call requires an active asynchronous context handle to be returned from the
	// EcDoAsyncConnectEx method on the EMSMDB interface, as specified in section 3.1.4.1.
	// The asynchronous context handle is associated with the Session Context.
	//
	// This method is part of notification handling. For more information about notifications,
	// see [MS-OXCNOTIF].
	//
	// Return Values: If the method succeeds, the return value is 0. If the method fails,
	// the return value is an implementation-specific error code or one of the protocol-defined
	// error codes listed in the following table.
	//
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	|   ERROR CODE    |            |                                                                                  |
	//	|      NAME       |   VALUE    |                                     MEANING                                      |
	//	|                 |            |                                                                                  |
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	| Rejected        | 0x000007EE | An EcDoAsyncWaitEx method call is already outstanding on this asynchronous       |
	//	|                 |            | context handle.<37>                                                              |
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//	| Exiting         | 0x000003ED | The server is shutting down.                                                     |
	//	+-----------------+------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol as specified in [MS-RPCE].
	DoAsyncWaitEx(context.Context, *DoAsyncWaitExRequest, ...dcerpc.CallOption) (*DoAsyncWaitExResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultAsyncemsmdbClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultAsyncemsmdbClient) DoAsyncWaitEx(ctx context.Context, in *DoAsyncWaitExRequest, opts ...dcerpc.CallOption) (*DoAsyncWaitExResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DoAsyncWaitExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultAsyncemsmdbClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultAsyncemsmdbClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewAsyncemsmdbClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (AsyncemsmdbClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(AsyncemsmdbSyntaxV0_1))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultAsyncemsmdbClient{cc: cc}, nil
}

// xxx_DoAsyncWaitExOperation structure represents the EcDoAsyncWaitEx operation
type xxx_DoAsyncWaitExOperation struct {
	AsyncSession *oxcrpc.AsyncSession `idl:"name:acxh" json:"async_session"`
	FlagsIn      uint32               `idl:"name:ulFlagsIn" json:"flags_in"`
	FlagsOut     uint32               `idl:"name:pulFlagsOut" json:"flags_out"`
	Return       int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_DoAsyncWaitExOperation) OpNum() int { return 0 }

func (o *xxx_DoAsyncWaitExOperation) OpName() string { return "/asyncemsmdb/v0.1/EcDoAsyncWaitEx" }

func (o *xxx_DoAsyncWaitExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoAsyncWaitExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// acxh {in} (1:{context_handle, alias=ACXH, names=ndr_context_handle}(struct))
	{
		if o.AsyncSession != nil {
			if err := o.AsyncSession.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&oxcrpc.AsyncSession{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// ulFlagsIn {in} (1:(uint32))
	{
		if err := w.WriteData(o.FlagsIn); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoAsyncWaitExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// acxh {in} (1:{context_handle, alias=ACXH, names=ndr_context_handle}(struct))
	{
		if o.AsyncSession == nil {
			o.AsyncSession = &oxcrpc.AsyncSession{}
		}
		if err := o.AsyncSession.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// ulFlagsIn {in} (1:(uint32))
	{
		if err := w.ReadData(&o.FlagsIn); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoAsyncWaitExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoAsyncWaitExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pulFlagsOut {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.FlagsOut); err != nil {
			return err
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

func (o *xxx_DoAsyncWaitExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pulFlagsOut {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.FlagsOut); err != nil {
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

// DoAsyncWaitExRequest structure represents the EcDoAsyncWaitEx operation request
type DoAsyncWaitExRequest struct {
	// acxh: An asynchronous context handle. On input, the client MUST pass a valid asynchronous
	// context handle that was created by calling the EcDoAsyncConnectEx method on the EMSMDB
	// interface. The server uses the asynchronous context handle to identify the Session
	// Context to use for this call. If the asynchronous context handle is not valid, the
	// server successfully completes the call, setting the NotificationPending flag in the
	// pulFlagsOut parameter.
	AsyncSession *oxcrpc.AsyncSession `idl:"name:acxh" json:"async_session"`
	// ulFlagsIn: Unused. Reserved for future use. Client MUST pass a value of 0x00000000.
	FlagsIn uint32 `idl:"name:ulFlagsIn" json:"flags_in"`
}

func (o *DoAsyncWaitExRequest) xxx_ToOp(ctx context.Context) *xxx_DoAsyncWaitExOperation {
	if o == nil {
		return &xxx_DoAsyncWaitExOperation{}
	}
	return &xxx_DoAsyncWaitExOperation{
		AsyncSession: o.AsyncSession,
		FlagsIn:      o.FlagsIn,
	}
}

func (o *DoAsyncWaitExRequest) xxx_FromOp(ctx context.Context, op *xxx_DoAsyncWaitExOperation) {
	if o == nil {
		return
	}
	o.AsyncSession = op.AsyncSession
	o.FlagsIn = op.FlagsIn
}
func (o *DoAsyncWaitExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DoAsyncWaitExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DoAsyncWaitExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DoAsyncWaitExResponse structure represents the EcDoAsyncWaitEx operation response
type DoAsyncWaitExResponse struct {
	// pulFlagsOut: The output flags for the client. Flag values are specified in the following
	// table.
	//
	//	+---------------------+------------+----------------------------------------------------------------------------------+
	//	|        FLAG         |            |                                                                                  |
	//	|        NAME         |   VALUE    |                                   DESCRIPTION                                    |
	//	|                     |            |                                                                                  |
	//	+---------------------+------------+----------------------------------------------------------------------------------+
	//	+---------------------+------------+----------------------------------------------------------------------------------+
	//	| NotificationPending | 0x00000001 | Signals that events are pending for the client on the Session Context on the     |
	//	|                     |            | server. The client MUST call the EcDoRpcExt2 method (with additional data in the |
	//	|                     |            | ROP request buffer if there is additional data to send to the server, or with an |
	//	|                     |            | empty ROP request buffer if there is no additional data to send to the server).  |
	//	|                     |            | The server will return the event details in the ROP response buffer.             |
	//	+---------------------+------------+----------------------------------------------------------------------------------+
	FlagsOut uint32 `idl:"name:pulFlagsOut" json:"flags_out"`
	// Return: The EcDoAsyncWaitEx return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *DoAsyncWaitExResponse) xxx_ToOp(ctx context.Context) *xxx_DoAsyncWaitExOperation {
	if o == nil {
		return &xxx_DoAsyncWaitExOperation{}
	}
	return &xxx_DoAsyncWaitExOperation{
		FlagsOut: o.FlagsOut,
		Return:   o.Return,
	}
}

func (o *DoAsyncWaitExResponse) xxx_FromOp(ctx context.Context, op *xxx_DoAsyncWaitExOperation) {
	if o == nil {
		return
	}
	o.FlagsOut = op.FlagsOut
	o.Return = op.Return
}
func (o *DoAsyncWaitExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DoAsyncWaitExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DoAsyncWaitExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
