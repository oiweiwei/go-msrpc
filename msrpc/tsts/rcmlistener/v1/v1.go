package rcmlistener

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
	_ = dcetypes.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "tsts"
)

var (
	// Syntax UUID
	RcmListenerSyntaxUUID = &uuid.UUID{TimeLow: 0x497d95a6, TimeMid: 0x2d27, TimeHiAndVersion: 0x4bf5, ClockSeqHiAndReserved: 0x9b, ClockSeqLow: 0xbd, Node: [6]uint8{0xa6, 0x4, 0x69, 0x57, 0x13, 0x3c}}
	// Syntax ID
	RcmListenerSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: RcmListenerSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// RCMListener interface.
type RcmListenerClient interface {

	// The RpcOpenListener method returns a handle to the specified Terminal Services listener
	// running on a terminal server. No special permissions are required to call this method.
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
	OpenListener(context.Context, *OpenListenerRequest, ...dcerpc.CallOption) (*OpenListenerResponse, error)

	// The RpcCloseListener method closes the handle for a Terminal Services listener running
	// on a terminal server. This MUST be called after RpcOpenListener. The call to this
	// method MUST be serialized if there are multiple threads running otherwise the behavior
	// of this function is unknown. No special permissions are required to call this method.
	//
	// Return Values:  The method MUST return S_OK (0x00000000) on success; otherwise,
	// it MUST return an implementation-specific negative value.<159>
	//
	//	+-------------------+------------------------+
	//	|      RETURN       |                        |
	//	|    VALUE/CODE     |      DESCRIPTION       |
	//	|                   |                        |
	//	+-------------------+------------------------+
	//	+-------------------+------------------------+
	//	| 0x00000000 S_OK   | Successful completion. |
	//	+-------------------+------------------------+
	CloseListener(context.Context, *CloseListenerRequest, ...dcerpc.CallOption) (*CloseListenerResponse, error)

	// The RpcStopListener method stops the specified Terminal Services listener running
	// on a terminal server. This MUST be called after RpcOpenListener. The caller MUST
	// have WINSTATION_RESET permission to the listener. The method checks whether the caller
	// has WINSTATION_RESET permission (section 3.1.1) by setting it as the Access Request
	// mask, and fails if the caller does not have the permission.
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
	StopListener(context.Context, *StopListenerRequest, ...dcerpc.CallOption) (*StopListenerResponse, error)

	// The RpcStartListener method starts the specified Terminal Services listener on a
	// terminal server. This MUST be called after RpcOpenListener. The caller MUST have
	// WINSTATION_RESET and WINSTATION_QUERY permissions to the listener. The method checks
	// whether the caller has WINSTATION_RESET and WINSTATION_QUERY permission (section
	// 3.1.1) by setting it as the Access Request mask, and fails if the caller does not
	// have the permissions.
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
	StartListener(context.Context, *StartListenerRequest, ...dcerpc.CallOption) (*StartListenerResponse, error)

	// The RpcIsListening method checks whether the specified Terminal Services listener
	// is running on a terminal server. This MUST be called after RpcOpenListener. The caller
	// MUST have WINSTATION_QUERY permission to the listener. The method checks whether
	// the caller has WINSTATION_QUERY permission (section 3.1.1) by setting it as the Access
	// Request mask, and fails if the caller does not have the permission.
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
	IsListening(context.Context, *IsListeningRequest, ...dcerpc.CallOption) (*IsListeningResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// Listener structure represents HLISTENER RPC structure.
type Listener dcetypes.ContextHandle

func (o *Listener) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Listener) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *Listener) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
func (o *Listener) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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

type xxx_DefaultRcmListenerClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultRcmListenerClient) OpenListener(ctx context.Context, in *OpenListenerRequest, opts ...dcerpc.CallOption) (*OpenListenerResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenListenerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRcmListenerClient) CloseListener(ctx context.Context, in *CloseListenerRequest, opts ...dcerpc.CallOption) (*CloseListenerResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseListenerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRcmListenerClient) StopListener(ctx context.Context, in *StopListenerRequest, opts ...dcerpc.CallOption) (*StopListenerResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &StopListenerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRcmListenerClient) StartListener(ctx context.Context, in *StartListenerRequest, opts ...dcerpc.CallOption) (*StartListenerResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &StartListenerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRcmListenerClient) IsListening(ctx context.Context, in *IsListeningRequest, opts ...dcerpc.CallOption) (*IsListeningResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &IsListeningResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultRcmListenerClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultRcmListenerClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewRcmListenerClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (RcmListenerClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(RcmListenerSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultRcmListenerClient{cc: cc}, nil
}

// xxx_OpenListenerOperation structure represents the RpcOpenListener operation
type xxx_OpenListenerOperation struct {
	ListenerName string    `idl:"name:szListenerName;string" json:"listener_name"`
	Listener     *Listener `idl:"name:phListener" json:"listener"`
	Return       int32     `idl:"name:Return" json:"return"`
}

// OpNum returns the operation number of RpcOpenListener operation.
func (o *xxx_OpenListenerOperation) OpNum() int { return 0 }

// OpName returns the operation name of RpcOpenListener operation.
func (o *xxx_OpenListenerOperation) OpName() string { return "/RCMListener/v1/RpcOpenListener" }

func (o *xxx_OpenListenerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenListenerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// szListenerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ListenerName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenListenerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// szListenerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ListenerName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenListenerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenListenerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phListener {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HLISTENER, names=ndr_context_handle}(struct))
	{
		if o.Listener != nil {
			if err := o.Listener.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Listener{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_OpenListenerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phListener {out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HLISTENER, names=ndr_context_handle}(struct))
	{
		if o.Listener == nil {
			o.Listener = &Listener{}
		}
		if err := o.Listener.UnmarshalNDR(ctx, w); err != nil {
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

// OpenListenerRequest structure represents the RpcOpenListener operation request
type OpenListenerRequest struct {
	// szListenerName:  The name of the listener for which to retrieve a handle.
	ListenerName string `idl:"name:szListenerName;string" json:"listener_name"`
}

func (o *OpenListenerRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenListenerOperation) *xxx_OpenListenerOperation {
	if op == nil {
		op = &xxx_OpenListenerOperation{}
	}
	if o == nil {
		return op
	}
	op.ListenerName = o.ListenerName
	return op
}

func (o *OpenListenerRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenListenerOperation) {
	if o == nil {
		return
	}
	o.ListenerName = op.ListenerName
}
func (o *OpenListenerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenListenerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenListenerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MakeOpenListenerRequest build a response structure from the given request structure.
func (o *OpenListenerRequest) MakeResponse() *OpenListenerResponse {
	return &OpenListenerResponse{}
}

// OpNum returns the operation number of RpcOpenListener operation.
func (o *OpenListenerRequest) OpNum() int { return 0 }

// OpName returns the operation name of RpcOpenListener operation.
func (o *OpenListenerRequest) OpName() string { return "/RCMListener/v1/RpcOpenListener" }

// OpenListenerResponse structure represents the RpcOpenListener operation response
type OpenListenerResponse struct {
	// phListener:  Pointer to a handle to the listener. The handle is of type HLISTENER.
	Listener *Listener `idl:"name:phListener" json:"listener"`
	// Return: The RpcOpenListener return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *OpenListenerResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenListenerOperation) *xxx_OpenListenerOperation {
	if op == nil {
		op = &xxx_OpenListenerOperation{}
	}
	if o == nil {
		return op
	}
	op.Listener = o.Listener
	op.Return = o.Return
	return op
}

func (o *OpenListenerResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenListenerOperation) {
	if o == nil {
		return
	}
	o.Listener = op.Listener
	o.Return = op.Return
}
func (o *OpenListenerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenListenerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenListenerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseListenerOperation structure represents the RpcCloseListener operation
type xxx_CloseListenerOperation struct {
	Listener *Listener `idl:"name:phListener" json:"listener"`
	Return   int32     `idl:"name:Return" json:"return"`
}

// OpNum returns the operation number of RpcCloseListener operation.
func (o *xxx_CloseListenerOperation) OpNum() int { return 1 }

// OpName returns the operation name of RpcCloseListener operation.
func (o *xxx_CloseListenerOperation) OpName() string { return "/RCMListener/v1/RpcCloseListener" }

func (o *xxx_CloseListenerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseListenerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// phListener {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HLISTENER, names=ndr_context_handle}(struct))
	{
		if o.Listener != nil {
			if err := o.Listener.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Listener{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseListenerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// phListener {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HLISTENER, names=ndr_context_handle}(struct))
	{
		if o.Listener == nil {
			o.Listener = &Listener{}
		}
		if err := o.Listener.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseListenerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseListenerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phListener {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HLISTENER, names=ndr_context_handle}(struct))
	{
		if o.Listener != nil {
			if err := o.Listener.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Listener{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_CloseListenerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phListener {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HLISTENER, names=ndr_context_handle}(struct))
	{
		if o.Listener == nil {
			o.Listener = &Listener{}
		}
		if err := o.Listener.UnmarshalNDR(ctx, w); err != nil {
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

// CloseListenerRequest structure represents the RpcCloseListener operation request
type CloseListenerRequest struct {
	// phListener: Pointer to a handle to the listener as returned by RpcOpenListener. The
	// handle is of type HLISTENER. The handle is set to NULL when the call returns.
	Listener *Listener `idl:"name:phListener" json:"listener"`
}

func (o *CloseListenerRequest) xxx_ToOp(ctx context.Context, op *xxx_CloseListenerOperation) *xxx_CloseListenerOperation {
	if op == nil {
		op = &xxx_CloseListenerOperation{}
	}
	if o == nil {
		return op
	}
	op.Listener = o.Listener
	return op
}

func (o *CloseListenerRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseListenerOperation) {
	if o == nil {
		return
	}
	o.Listener = op.Listener
}
func (o *CloseListenerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *CloseListenerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseListenerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MakeCloseListenerRequest build a response structure from the given request structure.
func (o *CloseListenerRequest) MakeResponse() *CloseListenerResponse {
	return &CloseListenerResponse{}
}

// OpNum returns the operation number of RpcCloseListener operation.
func (o *CloseListenerRequest) OpNum() int { return 1 }

// OpName returns the operation name of RpcCloseListener operation.
func (o *CloseListenerRequest) OpName() string { return "/RCMListener/v1/RpcCloseListener" }

// CloseListenerResponse structure represents the RpcCloseListener operation response
type CloseListenerResponse struct {
	// phListener: Pointer to a handle to the listener as returned by RpcOpenListener. The
	// handle is of type HLISTENER. The handle is set to NULL when the call returns.
	Listener *Listener `idl:"name:phListener" json:"listener"`
	// Return: The RpcCloseListener return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CloseListenerResponse) xxx_ToOp(ctx context.Context, op *xxx_CloseListenerOperation) *xxx_CloseListenerOperation {
	if op == nil {
		op = &xxx_CloseListenerOperation{}
	}
	if o == nil {
		return op
	}
	op.Listener = o.Listener
	op.Return = o.Return
	return op
}

func (o *CloseListenerResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseListenerOperation) {
	if o == nil {
		return
	}
	o.Listener = op.Listener
	o.Return = op.Return
}
func (o *CloseListenerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *CloseListenerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseListenerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StopListenerOperation structure represents the RpcStopListener operation
type xxx_StopListenerOperation struct {
	Listener *Listener `idl:"name:hListener" json:"listener"`
	Return   int32     `idl:"name:Return" json:"return"`
}

// OpNum returns the operation number of RpcStopListener operation.
func (o *xxx_StopListenerOperation) OpNum() int { return 2 }

// OpName returns the operation name of RpcStopListener operation.
func (o *xxx_StopListenerOperation) OpName() string { return "/RCMListener/v1/RpcStopListener" }

func (o *xxx_StopListenerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopListenerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hListener {in} (1:{context_handle, alias=HLISTENER, names=ndr_context_handle}(struct))
	{
		if o.Listener != nil {
			if err := o.Listener.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Listener{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_StopListenerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hListener {in} (1:{context_handle, alias=HLISTENER, names=ndr_context_handle}(struct))
	{
		if o.Listener == nil {
			o.Listener = &Listener{}
		}
		if err := o.Listener.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopListenerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StopListenerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_StopListenerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// StopListenerRequest structure represents the RpcStopListener operation request
type StopListenerRequest struct {
	// hListener: The handle to the listener. This is of type HLISTENER.
	Listener *Listener `idl:"name:hListener" json:"listener"`
}

func (o *StopListenerRequest) xxx_ToOp(ctx context.Context, op *xxx_StopListenerOperation) *xxx_StopListenerOperation {
	if op == nil {
		op = &xxx_StopListenerOperation{}
	}
	if o == nil {
		return op
	}
	op.Listener = o.Listener
	return op
}

func (o *StopListenerRequest) xxx_FromOp(ctx context.Context, op *xxx_StopListenerOperation) {
	if o == nil {
		return
	}
	o.Listener = op.Listener
}
func (o *StopListenerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *StopListenerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StopListenerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MakeStopListenerRequest build a response structure from the given request structure.
func (o *StopListenerRequest) MakeResponse() *StopListenerResponse {
	return &StopListenerResponse{}
}

// OpNum returns the operation number of RpcStopListener operation.
func (o *StopListenerRequest) OpNum() int { return 2 }

// OpName returns the operation name of RpcStopListener operation.
func (o *StopListenerRequest) OpName() string { return "/RCMListener/v1/RpcStopListener" }

// StopListenerResponse structure represents the RpcStopListener operation response
type StopListenerResponse struct {
	// Return: The RpcStopListener return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StopListenerResponse) xxx_ToOp(ctx context.Context, op *xxx_StopListenerOperation) *xxx_StopListenerOperation {
	if op == nil {
		op = &xxx_StopListenerOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *StopListenerResponse) xxx_FromOp(ctx context.Context, op *xxx_StopListenerOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *StopListenerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *StopListenerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StopListenerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StartListenerOperation structure represents the RpcStartListener operation
type xxx_StartListenerOperation struct {
	Listener *Listener `idl:"name:hListener" json:"listener"`
	Return   int32     `idl:"name:Return" json:"return"`
}

// OpNum returns the operation number of RpcStartListener operation.
func (o *xxx_StartListenerOperation) OpNum() int { return 3 }

// OpName returns the operation name of RpcStartListener operation.
func (o *xxx_StartListenerOperation) OpName() string { return "/RCMListener/v1/RpcStartListener" }

func (o *xxx_StartListenerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartListenerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hListener {in} (1:{context_handle, alias=HLISTENER, names=ndr_context_handle}(struct))
	{
		if o.Listener != nil {
			if err := o.Listener.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Listener{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_StartListenerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hListener {in} (1:{context_handle, alias=HLISTENER, names=ndr_context_handle}(struct))
	{
		if o.Listener == nil {
			o.Listener = &Listener{}
		}
		if err := o.Listener.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartListenerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartListenerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_StartListenerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// StartListenerRequest structure represents the RpcStartListener operation request
type StartListenerRequest struct {
	// hListener:  The handle to the listener. This is of type HLISTENER.
	Listener *Listener `idl:"name:hListener" json:"listener"`
}

func (o *StartListenerRequest) xxx_ToOp(ctx context.Context, op *xxx_StartListenerOperation) *xxx_StartListenerOperation {
	if op == nil {
		op = &xxx_StartListenerOperation{}
	}
	if o == nil {
		return op
	}
	op.Listener = o.Listener
	return op
}

func (o *StartListenerRequest) xxx_FromOp(ctx context.Context, op *xxx_StartListenerOperation) {
	if o == nil {
		return
	}
	o.Listener = op.Listener
}
func (o *StartListenerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *StartListenerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StartListenerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MakeStartListenerRequest build a response structure from the given request structure.
func (o *StartListenerRequest) MakeResponse() *StartListenerResponse {
	return &StartListenerResponse{}
}

// OpNum returns the operation number of RpcStartListener operation.
func (o *StartListenerRequest) OpNum() int { return 3 }

// OpName returns the operation name of RpcStartListener operation.
func (o *StartListenerRequest) OpName() string { return "/RCMListener/v1/RpcStartListener" }

// StartListenerResponse structure represents the RpcStartListener operation response
type StartListenerResponse struct {
	// Return: The RpcStartListener return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *StartListenerResponse) xxx_ToOp(ctx context.Context, op *xxx_StartListenerOperation) *xxx_StartListenerOperation {
	if op == nil {
		op = &xxx_StartListenerOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *StartListenerResponse) xxx_FromOp(ctx context.Context, op *xxx_StartListenerOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *StartListenerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *StartListenerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StartListenerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsListeningOperation structure represents the RpcIsListening operation
type xxx_IsListeningOperation struct {
	Listener    *Listener `idl:"name:hListener" json:"listener"`
	IsListening bool      `idl:"name:pbIsListening" json:"is_listening"`
	Return      int32     `idl:"name:Return" json:"return"`
}

// OpNum returns the operation number of RpcIsListening operation.
func (o *xxx_IsListeningOperation) OpNum() int { return 4 }

// OpName returns the operation name of RpcIsListening operation.
func (o *xxx_IsListeningOperation) OpName() string { return "/RCMListener/v1/RpcIsListening" }

func (o *xxx_IsListeningOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsListeningOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hListener {in} (1:{context_handle, alias=HLISTENER, names=ndr_context_handle}(struct))
	{
		if o.Listener != nil {
			if err := o.Listener.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Listener{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_IsListeningOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hListener {in} (1:{context_handle, alias=HLISTENER, names=ndr_context_handle}(struct))
	{
		if o.Listener == nil {
			o.Listener = &Listener{}
		}
		if err := o.Listener.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsListeningOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsListeningOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pbIsListening {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		if !o.IsListening {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
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

func (o *xxx_IsListeningOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pbIsListening {out} (1:{pointer=ref}*(1))(2:{alias=BOOL}(int32))
	{
		var _bIsListening int32
		if err := w.ReadData(&_bIsListening); err != nil {
			return err
		}
		o.IsListening = _bIsListening != 0
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// IsListeningRequest structure represents the RpcIsListening operation request
type IsListeningRequest struct {
	// hListener: The handle to the listener. This is of type HLISTENER.
	Listener *Listener `idl:"name:hListener" json:"listener"`
}

func (o *IsListeningRequest) xxx_ToOp(ctx context.Context, op *xxx_IsListeningOperation) *xxx_IsListeningOperation {
	if op == nil {
		op = &xxx_IsListeningOperation{}
	}
	if o == nil {
		return op
	}
	op.Listener = o.Listener
	return op
}

func (o *IsListeningRequest) xxx_FromOp(ctx context.Context, op *xxx_IsListeningOperation) {
	if o == nil {
		return
	}
	o.Listener = op.Listener
}
func (o *IsListeningRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *IsListeningRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsListeningOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MakeIsListeningRequest build a response structure from the given request structure.
func (o *IsListeningRequest) MakeResponse() *IsListeningResponse {
	return &IsListeningResponse{}
}

// OpNum returns the operation number of RpcIsListening operation.
func (o *IsListeningRequest) OpNum() int { return 4 }

// OpName returns the operation name of RpcIsListening operation.
func (o *IsListeningRequest) OpName() string { return "/RCMListener/v1/RpcIsListening" }

// IsListeningResponse structure represents the RpcIsListening operation response
type IsListeningResponse struct {
	// pbIsListening: Set to TRUE if the listener is listening for a connection, FALSE otherwise.
	IsListening bool `idl:"name:pbIsListening" json:"is_listening"`
	// Return: The RpcIsListening return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsListeningResponse) xxx_ToOp(ctx context.Context, op *xxx_IsListeningOperation) *xxx_IsListeningOperation {
	if op == nil {
		op = &xxx_IsListeningOperation{}
	}
	if o == nil {
		return op
	}
	op.IsListening = o.IsListening
	op.Return = o.Return
	return op
}

func (o *IsListeningResponse) xxx_FromOp(ctx context.Context, op *xxx_IsListeningOperation) {
	if o == nil {
		return
	}
	o.IsListening = op.IsListening
	o.Return = op.Return
}
func (o *IsListeningResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *IsListeningResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsListeningOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
