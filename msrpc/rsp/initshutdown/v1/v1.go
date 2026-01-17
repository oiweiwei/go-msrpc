package initshutdown

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	rsp "github.com/oiweiwei/go-msrpc/msrpc/rsp"
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
	_ = rsp.GoPackage
)

var (
	// import guard
	GoPackage = "rsp"
)

var (
	// Syntax UUID
	InitShutdownSyntaxUUID = &uuid.UUID{TimeLow: 0x894de0c0, TimeMid: 0xd55, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0x22, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}
	// Syntax ID
	InitShutdownSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: InitShutdownSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// InitShutdown interface.
type InitShutdownClient interface {

	// The BaseInitiateShutdown method is used to initiate the shutdown of the remote computer.<8>
	//
	// Return Values: The method returns ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it returns a nonzero error code.<9>
	BaseInitiateShutdown(context.Context, *BaseInitiateShutdownRequest, ...dcerpc.CallOption) (*BaseInitiateShutdownResponse, error)

	// The BaseAbortShutdown method is used to terminate the shutdown of the remote computer
	// within the waiting period.<10>
	//
	// Return Values: The method returns ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it returns a nonzero error code.
	BaseAbortShutdown(context.Context, *BaseAbortShutdownRequest, ...dcerpc.CallOption) (*BaseAbortShutdownResponse, error)

	// The BaseInitiateShutdownEx method is used to initiate the shutdown of the remote
	// computer.<11>
	//
	// Return Values: The method returns ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it returns a nonzero error code.
	BaseInitiateShutdownEx(context.Context, *BaseInitiateShutdownExRequest, ...dcerpc.CallOption) (*BaseInitiateShutdownExResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultInitShutdownClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultInitShutdownClient) BaseInitiateShutdown(ctx context.Context, in *BaseInitiateShutdownRequest, opts ...dcerpc.CallOption) (*BaseInitiateShutdownResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseInitiateShutdownResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInitShutdownClient) BaseAbortShutdown(ctx context.Context, in *BaseAbortShutdownRequest, opts ...dcerpc.CallOption) (*BaseAbortShutdownResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseAbortShutdownResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInitShutdownClient) BaseInitiateShutdownEx(ctx context.Context, in *BaseInitiateShutdownExRequest, opts ...dcerpc.CallOption) (*BaseInitiateShutdownExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseInitiateShutdownExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), o.cc.Error(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInitShutdownClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultInitShutdownClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewInitShutdownClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (InitShutdownClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(InitShutdownSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultInitShutdownClient{cc: cc}, nil
}

// xxx_BaseInitiateShutdownOperation structure represents the BaseInitiateShutdown operation
type xxx_BaseInitiateShutdownOperation struct {
	ServerName          string             `idl:"name:ServerName;pointer:unique" json:"server_name"`
	Message             *rsp.UnicodeString `idl:"name:lpMessage;pointer:unique" json:"message"`
	Timeout             uint32             `idl:"name:dwTimeout" json:"timeout"`
	ForceAppsClosed     uint8              `idl:"name:bForceAppsClosed" json:"force_apps_closed"`
	RebootAfterShutdown uint8              `idl:"name:bRebootAfterShutdown" json:"reboot_after_shutdown"`
	Return              uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseInitiateShutdownOperation) OpNum() int { return 0 }

func (o *xxx_BaseInitiateShutdownOperation) OpName() string {
	return "/InitShutdown/v1/BaseInitiateShutdown"
}

func (o *xxx_BaseInitiateShutdownOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseInitiateShutdownOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME}*(1)[dim:0,string](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// lpMessage {in} (1:{pointer=unique, alias=PREG_UNICODE_STRING}*(1))(2:{alias=REG_UNICODE_STRING}(struct))
	{
		if o.Message != nil {
			_ptr_lpMessage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&rsp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_lpMessage); err != nil {
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
	// dwTimeout {in} (1:(uint32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	// bForceAppsClosed {in} (1:(uchar))
	{
		if err := w.WriteData(o.ForceAppsClosed); err != nil {
			return err
		}
	}
	// bRebootAfterShutdown {in} (1:(uchar))
	{
		if err := w.WriteData(o.RebootAfterShutdown); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseInitiateShutdownOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME}*(1)[dim:0,string](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpMessage {in} (1:{pointer=unique, alias=PREG_UNICODE_STRING}*(1))(2:{alias=REG_UNICODE_STRING}(struct))
	{
		_ptr_lpMessage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &rsp.UnicodeString{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpMessage := func(ptr interface{}) { o.Message = *ptr.(**rsp.UnicodeString) }
		if err := w.ReadPointer(&o.Message, _s_lpMessage, _ptr_lpMessage); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwTimeout {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	// bForceAppsClosed {in} (1:(uchar))
	{
		if err := w.ReadData(&o.ForceAppsClosed); err != nil {
			return err
		}
	}
	// bRebootAfterShutdown {in} (1:(uchar))
	{
		if err := w.ReadData(&o.RebootAfterShutdown); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseInitiateShutdownOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseInitiateShutdownOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseInitiateShutdownOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BaseInitiateShutdownRequest structure represents the BaseInitiateShutdown operation request
type BaseInitiateShutdownRequest struct {
	// ServerName: The custom RPC binding handle (PREGISTRY_SERVER_NAME (section 2.2.1)).
	ServerName string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	// lpMessage: Null-terminated Unicode string that contains the message to display during
	// the shutdown waiting period. If this parameter is NULL, no message MUST be displayed.
	Message *rsp.UnicodeString `idl:"name:lpMessage;pointer:unique" json:"message"`
	// dwTimeout: Number of seconds to wait before shutting down.
	Timeout uint32 `idl:"name:dwTimeout" json:"timeout"`
	// bForceAppsClosed: If TRUE, all applications SHOULD be terminated unconditionally.
	ForceAppsClosed uint8 `idl:"name:bForceAppsClosed" json:"force_apps_closed"`
	// bRebootAfterShutdown: If TRUE, the system SHOULD shut down and reboot. If FALSE,
	// the system SHOULD only shut down.
	RebootAfterShutdown uint8 `idl:"name:bRebootAfterShutdown" json:"reboot_after_shutdown"`
}

func (o *BaseInitiateShutdownRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseInitiateShutdownOperation) *xxx_BaseInitiateShutdownOperation {
	if op == nil {
		op = &xxx_BaseInitiateShutdownOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.Message = o.Message
	op.Timeout = o.Timeout
	op.ForceAppsClosed = o.ForceAppsClosed
	op.RebootAfterShutdown = o.RebootAfterShutdown
	return op
}

func (o *BaseInitiateShutdownRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseInitiateShutdownOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.Message = op.Message
	o.Timeout = op.Timeout
	o.ForceAppsClosed = op.ForceAppsClosed
	o.RebootAfterShutdown = op.RebootAfterShutdown
}
func (o *BaseInitiateShutdownRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseInitiateShutdownRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseInitiateShutdownOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseInitiateShutdownResponse structure represents the BaseInitiateShutdown operation response
type BaseInitiateShutdownResponse struct {
	// Return: The BaseInitiateShutdown return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseInitiateShutdownResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseInitiateShutdownOperation) *xxx_BaseInitiateShutdownOperation {
	if op == nil {
		op = &xxx_BaseInitiateShutdownOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *BaseInitiateShutdownResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseInitiateShutdownOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BaseInitiateShutdownResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseInitiateShutdownResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseInitiateShutdownOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseAbortShutdownOperation structure represents the BaseAbortShutdown operation
type xxx_BaseAbortShutdownOperation struct {
	ServerName string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseAbortShutdownOperation) OpNum() int { return 1 }

func (o *xxx_BaseAbortShutdownOperation) OpName() string { return "/InitShutdown/v1/BaseAbortShutdown" }

func (o *xxx_BaseAbortShutdownOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseAbortShutdownOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME}*(1)[dim:0,string](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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

func (o *xxx_BaseAbortShutdownOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME}*(1)[dim:0,string](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseAbortShutdownOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseAbortShutdownOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseAbortShutdownOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BaseAbortShutdownRequest structure represents the BaseAbortShutdown operation request
type BaseAbortShutdownRequest struct {
	// ServerName: The custom RPC binding handle (PREGISTRY_SERVER_NAME (section 2.2.1)).
	ServerName string `idl:"name:ServerName;pointer:unique" json:"server_name"`
}

func (o *BaseAbortShutdownRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseAbortShutdownOperation) *xxx_BaseAbortShutdownOperation {
	if op == nil {
		op = &xxx_BaseAbortShutdownOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	return op
}

func (o *BaseAbortShutdownRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseAbortShutdownOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
}
func (o *BaseAbortShutdownRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseAbortShutdownRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseAbortShutdownOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseAbortShutdownResponse structure represents the BaseAbortShutdown operation response
type BaseAbortShutdownResponse struct {
	// Return: The BaseAbortShutdown return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseAbortShutdownResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseAbortShutdownOperation) *xxx_BaseAbortShutdownOperation {
	if op == nil {
		op = &xxx_BaseAbortShutdownOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *BaseAbortShutdownResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseAbortShutdownOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BaseAbortShutdownResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseAbortShutdownResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseAbortShutdownOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseInitiateShutdownExOperation structure represents the BaseInitiateShutdownEx operation
type xxx_BaseInitiateShutdownExOperation struct {
	ServerName          string             `idl:"name:ServerName;pointer:unique" json:"server_name"`
	Message             *rsp.UnicodeString `idl:"name:lpMessage;pointer:unique" json:"message"`
	Timeout             uint32             `idl:"name:dwTimeout" json:"timeout"`
	ForceAppsClosed     uint8              `idl:"name:bForceAppsClosed" json:"force_apps_closed"`
	RebootAfterShutdown uint8              `idl:"name:bRebootAfterShutdown" json:"reboot_after_shutdown"`
	Reason              uint32             `idl:"name:dwReason" json:"reason"`
	Return              uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseInitiateShutdownExOperation) OpNum() int { return 2 }

func (o *xxx_BaseInitiateShutdownExOperation) OpName() string {
	return "/InitShutdown/v1/BaseInitiateShutdownEx"
}

func (o *xxx_BaseInitiateShutdownExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseInitiateShutdownExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME}*(1)[dim:0,string](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// lpMessage {in} (1:{pointer=unique, alias=PREG_UNICODE_STRING}*(1))(2:{alias=REG_UNICODE_STRING}(struct))
	{
		if o.Message != nil {
			_ptr_lpMessage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Message != nil {
					if err := o.Message.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&rsp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Message, _ptr_lpMessage); err != nil {
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
	// dwTimeout {in} (1:(uint32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	// bForceAppsClosed {in} (1:(uchar))
	{
		if err := w.WriteData(o.ForceAppsClosed); err != nil {
			return err
		}
	}
	// bRebootAfterShutdown {in} (1:(uchar))
	{
		if err := w.WriteData(o.RebootAfterShutdown); err != nil {
			return err
		}
	}
	// dwReason {in} (1:(uint32))
	{
		if err := w.WriteData(o.Reason); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseInitiateShutdownExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME}*(1)[dim:0,string](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpMessage {in} (1:{pointer=unique, alias=PREG_UNICODE_STRING}*(1))(2:{alias=REG_UNICODE_STRING}(struct))
	{
		_ptr_lpMessage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Message == nil {
				o.Message = &rsp.UnicodeString{}
			}
			if err := o.Message.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpMessage := func(ptr interface{}) { o.Message = *ptr.(**rsp.UnicodeString) }
		if err := w.ReadPointer(&o.Message, _s_lpMessage, _ptr_lpMessage); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwTimeout {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	// bForceAppsClosed {in} (1:(uchar))
	{
		if err := w.ReadData(&o.ForceAppsClosed); err != nil {
			return err
		}
	}
	// bRebootAfterShutdown {in} (1:(uchar))
	{
		if err := w.ReadData(&o.RebootAfterShutdown); err != nil {
			return err
		}
	}
	// dwReason {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Reason); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseInitiateShutdownExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseInitiateShutdownExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseInitiateShutdownExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BaseInitiateShutdownExRequest structure represents the BaseInitiateShutdownEx operation request
type BaseInitiateShutdownExRequest struct {
	// ServerName: The custom RPC binding handle (PREGISTRY_SERVER_NAME (section 2.2.1)).
	ServerName string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	// lpMessage: Null-terminated Unicode string that contains the message to display during
	// the shutdown waiting period. If this parameter is NULL, no message MUST be displayed.
	Message *rsp.UnicodeString `idl:"name:lpMessage;pointer:unique" json:"message"`
	// dwTimeout: Number of seconds to wait before shutting down.
	Timeout uint32 `idl:"name:dwTimeout" json:"timeout"`
	// bForceAppsClosed: If TRUE, all applications SHOULD be terminated unconditionally.
	ForceAppsClosed uint8 `idl:"name:bForceAppsClosed" json:"force_apps_closed"`
	// bRebootAfterShutdown: If TRUE, the system SHOULD shut down and reboot. If FALSE,
	// the system SHOULD only shut down.
	RebootAfterShutdown uint8 `idl:"name:bRebootAfterShutdown" json:"reboot_after_shutdown"`
	// dwReason: Reason for initiating the shutdown (section 2.3).
	Reason uint32 `idl:"name:dwReason" json:"reason"`
}

func (o *BaseInitiateShutdownExRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseInitiateShutdownExOperation) *xxx_BaseInitiateShutdownExOperation {
	if op == nil {
		op = &xxx_BaseInitiateShutdownExOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.Message = o.Message
	op.Timeout = o.Timeout
	op.ForceAppsClosed = o.ForceAppsClosed
	op.RebootAfterShutdown = o.RebootAfterShutdown
	op.Reason = o.Reason
	return op
}

func (o *BaseInitiateShutdownExRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseInitiateShutdownExOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.Message = op.Message
	o.Timeout = op.Timeout
	o.ForceAppsClosed = op.ForceAppsClosed
	o.RebootAfterShutdown = op.RebootAfterShutdown
	o.Reason = op.Reason
}
func (o *BaseInitiateShutdownExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseInitiateShutdownExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseInitiateShutdownExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseInitiateShutdownExResponse structure represents the BaseInitiateShutdownEx operation response
type BaseInitiateShutdownExResponse struct {
	// Return: The BaseInitiateShutdownEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseInitiateShutdownExResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseInitiateShutdownExOperation) *xxx_BaseInitiateShutdownExOperation {
	if op == nil {
		op = &xxx_BaseInitiateShutdownExOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *BaseInitiateShutdownExResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseInitiateShutdownExOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BaseInitiateShutdownExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseInitiateShutdownExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseInitiateShutdownExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
