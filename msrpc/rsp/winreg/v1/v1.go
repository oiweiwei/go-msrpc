package winreg

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
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
	_ = (*errors.Error)(nil)
	_ = rsp.GoPackage
)

var (
	// import guard
	GoPackage = "rsp"
)

var (
	// Syntax UUID
	WinregSyntaxUUID = &uuid.UUID{TimeLow: 0x338cd001, TimeMid: 0x2244, TimeHiAndVersion: 0x31f1, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0xaa, Node: [6]uint8{0x90, 0x0, 0x38, 0x0, 0x10, 0x3}}
	// Syntax ID
	WinregSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: WinregSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// winreg interface.
type WinregClient interface {

	// Opnum0NotImplemented operation.
	// Opnum0NotImplemented

	// Opnum1NotImplemented operation.
	// Opnum1NotImplemented

	// Opnum2NotImplemented operation.
	// Opnum2NotImplemented

	// Opnum3NotImplemented operation.
	// Opnum3NotImplemented

	// Opnum4NotImplemented operation.
	// Opnum4NotImplemented

	// Opnum5NotImplemented operation.
	// Opnum5NotImplemented

	// Opnum6NotImplemented operation.
	// Opnum6NotImplemented

	// Opnum7NotImplemented operation.
	// Opnum7NotImplemented

	// Opnum8NotImplemented operation.
	// Opnum8NotImplemented

	// Opnum9NotImplemented operation.
	// Opnum9NotImplemented

	// Opnum10NotImplemented operation.
	// Opnum10NotImplemented

	// Opnum11NotImplemented operation.
	// Opnum11NotImplemented

	// Opnum12NotImplemented operation.
	// Opnum12NotImplemented

	// Opnum13NotImplemented operation.
	// Opnum13NotImplemented

	// Opnum14NotImplemented operation.
	// Opnum14NotImplemented

	// Opnum15NotImplemented operation.
	// Opnum15NotImplemented

	// Opnum16NotImplemented operation.
	// Opnum16NotImplemented

	// Opnum17NotImplemented operation.
	// Opnum17NotImplemented

	// Opnum18NotImplemented operation.
	// Opnum18NotImplemented

	// Opnum19NotImplemented operation.
	// Opnum19NotImplemented

	// Opnum20NotImplemented operation.
	// Opnum20NotImplemented

	// Opnum21NotImplemented operation.
	// Opnum21NotImplemented

	// Opnum22NotImplemented operation.
	// Opnum22NotImplemented

	// Opnum23NotImplemented operation.
	// Opnum23NotImplemented

	// The BaseInitiateSystemShutdown method is used to initiate the shutdown of the remote
	// computer.<4>
	//
	// Return Values: The method returns ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it returns a nonzero error code.
	BaseInitiateSystemShutdown(context.Context, *BaseInitiateSystemShutdownRequest, ...dcerpc.CallOption) (*BaseInitiateSystemShutdownResponse, error)

	// The BaseAbortSystemShutdown method is used to terminate the shutdown of the remote
	// computer within the waiting period.<5>
	//
	// Return Values: The method returns ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it returns a nonzero error code.
	BaseAbortSystemShutdown(context.Context, *BaseAbortSystemShutdownRequest, ...dcerpc.CallOption) (*BaseAbortSystemShutdownResponse, error)

	// Opnum26NotImplemented operation.
	// Opnum26NotImplemented

	// Opnum27NotImplemented operation.
	// Opnum27NotImplemented

	// Opnum28NotImplemented operation.
	// Opnum28NotImplemented

	// Opnum29NotImplemented operation.
	// Opnum29NotImplemented

	// The BaseInitiateSystemShutdownEx method is used to initiate the shutdown of the remote
	// computer with the reason for initiating the shutdown given as a parameter to the
	// call.<6>
	//
	// Return Values: The method returns ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it returns a nonzero error code.
	BaseInitiateSystemShutdownEx(context.Context, *BaseInitiateSystemShutdownExRequest, ...dcerpc.CallOption) (*BaseInitiateSystemShutdownExResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error
}

type xxx_DefaultWinregClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultWinregClient) BaseInitiateSystemShutdown(ctx context.Context, in *BaseInitiateSystemShutdownRequest, opts ...dcerpc.CallOption) (*BaseInitiateSystemShutdownResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseInitiateSystemShutdownResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseAbortSystemShutdown(ctx context.Context, in *BaseAbortSystemShutdownRequest, opts ...dcerpc.CallOption) (*BaseAbortSystemShutdownResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseAbortSystemShutdownResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseInitiateSystemShutdownEx(ctx context.Context, in *BaseInitiateSystemShutdownExRequest, opts ...dcerpc.CallOption) (*BaseInitiateSystemShutdownExResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseInitiateSystemShutdownExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}
func NewWinregClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (WinregClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(WinregSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultWinregClient{cc: cc}, nil
}

// xxx_BaseInitiateSystemShutdownOperation structure represents the BaseInitiateSystemShutdown operation
type xxx_BaseInitiateSystemShutdownOperation struct {
	ServerName          string             `idl:"name:ServerName;pointer:unique" json:"server_name"`
	Message             *rsp.UnicodeString `idl:"name:lpMessage;pointer:unique" json:"message"`
	Timeout             uint32             `idl:"name:dwTimeout" json:"timeout"`
	ForceAppsClosed     uint8              `idl:"name:bForceAppsClosed" json:"force_apps_closed"`
	RebootAfterShutdown uint8              `idl:"name:bRebootAfterShutdown" json:"reboot_after_shutdown"`
	Return              uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseInitiateSystemShutdownOperation) OpNum() int { return 24 }

func (o *xxx_BaseInitiateSystemShutdownOperation) OpName() string {
	return "/winreg/v1/BaseInitiateSystemShutdown"
}

func (o *xxx_BaseInitiateSystemShutdownOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseInitiateSystemShutdownOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BaseInitiateSystemShutdownOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_BaseInitiateSystemShutdownOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseInitiateSystemShutdownOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BaseInitiateSystemShutdownOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BaseInitiateSystemShutdownRequest structure represents the BaseInitiateSystemShutdown operation request
type BaseInitiateSystemShutdownRequest struct {
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

func (o *BaseInitiateSystemShutdownRequest) xxx_ToOp(ctx context.Context) *xxx_BaseInitiateSystemShutdownOperation {
	if o == nil {
		return &xxx_BaseInitiateSystemShutdownOperation{}
	}
	return &xxx_BaseInitiateSystemShutdownOperation{
		ServerName:          o.ServerName,
		Message:             o.Message,
		Timeout:             o.Timeout,
		ForceAppsClosed:     o.ForceAppsClosed,
		RebootAfterShutdown: o.RebootAfterShutdown,
	}
}

func (o *BaseInitiateSystemShutdownRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseInitiateSystemShutdownOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.Message = op.Message
	o.Timeout = op.Timeout
	o.ForceAppsClosed = op.ForceAppsClosed
	o.RebootAfterShutdown = op.RebootAfterShutdown
}
func (o *BaseInitiateSystemShutdownRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *BaseInitiateSystemShutdownRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseInitiateSystemShutdownOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseInitiateSystemShutdownResponse structure represents the BaseInitiateSystemShutdown operation response
type BaseInitiateSystemShutdownResponse struct {
	// Return: The BaseInitiateSystemShutdown return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseInitiateSystemShutdownResponse) xxx_ToOp(ctx context.Context) *xxx_BaseInitiateSystemShutdownOperation {
	if o == nil {
		return &xxx_BaseInitiateSystemShutdownOperation{}
	}
	return &xxx_BaseInitiateSystemShutdownOperation{
		Return: o.Return,
	}
}

func (o *BaseInitiateSystemShutdownResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseInitiateSystemShutdownOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BaseInitiateSystemShutdownResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *BaseInitiateSystemShutdownResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseInitiateSystemShutdownOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseAbortSystemShutdownOperation structure represents the BaseAbortSystemShutdown operation
type xxx_BaseAbortSystemShutdownOperation struct {
	ServerName string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseAbortSystemShutdownOperation) OpNum() int { return 25 }

func (o *xxx_BaseAbortSystemShutdownOperation) OpName() string {
	return "/winreg/v1/BaseAbortSystemShutdown"
}

func (o *xxx_BaseAbortSystemShutdownOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseAbortSystemShutdownOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BaseAbortSystemShutdownOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_BaseAbortSystemShutdownOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseAbortSystemShutdownOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BaseAbortSystemShutdownOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BaseAbortSystemShutdownRequest structure represents the BaseAbortSystemShutdown operation request
type BaseAbortSystemShutdownRequest struct {
	// ServerName: The custom RPC binding handle (PREGISTRY_SERVER_NAME (section 2.2.1)).
	ServerName string `idl:"name:ServerName;pointer:unique" json:"server_name"`
}

func (o *BaseAbortSystemShutdownRequest) xxx_ToOp(ctx context.Context) *xxx_BaseAbortSystemShutdownOperation {
	if o == nil {
		return &xxx_BaseAbortSystemShutdownOperation{}
	}
	return &xxx_BaseAbortSystemShutdownOperation{
		ServerName: o.ServerName,
	}
}

func (o *BaseAbortSystemShutdownRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseAbortSystemShutdownOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
}
func (o *BaseAbortSystemShutdownRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *BaseAbortSystemShutdownRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseAbortSystemShutdownOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseAbortSystemShutdownResponse structure represents the BaseAbortSystemShutdown operation response
type BaseAbortSystemShutdownResponse struct {
	// Return: The BaseAbortSystemShutdown return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseAbortSystemShutdownResponse) xxx_ToOp(ctx context.Context) *xxx_BaseAbortSystemShutdownOperation {
	if o == nil {
		return &xxx_BaseAbortSystemShutdownOperation{}
	}
	return &xxx_BaseAbortSystemShutdownOperation{
		Return: o.Return,
	}
}

func (o *BaseAbortSystemShutdownResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseAbortSystemShutdownOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BaseAbortSystemShutdownResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *BaseAbortSystemShutdownResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseAbortSystemShutdownOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseInitiateSystemShutdownExOperation structure represents the BaseInitiateSystemShutdownEx operation
type xxx_BaseInitiateSystemShutdownExOperation struct {
	ServerName          string             `idl:"name:ServerName;pointer:unique" json:"server_name"`
	Message             *rsp.UnicodeString `idl:"name:lpMessage;pointer:unique" json:"message"`
	Timeout             uint32             `idl:"name:dwTimeout" json:"timeout"`
	ForceAppsClosed     uint8              `idl:"name:bForceAppsClosed" json:"force_apps_closed"`
	RebootAfterShutdown uint8              `idl:"name:bRebootAfterShutdown" json:"reboot_after_shutdown"`
	Reason              uint32             `idl:"name:dwReason" json:"reason"`
	Return              uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseInitiateSystemShutdownExOperation) OpNum() int { return 30 }

func (o *xxx_BaseInitiateSystemShutdownExOperation) OpName() string {
	return "/winreg/v1/BaseInitiateSystemShutdownEx"
}

func (o *xxx_BaseInitiateSystemShutdownExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseInitiateSystemShutdownExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BaseInitiateSystemShutdownExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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

func (o *xxx_BaseInitiateSystemShutdownExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseInitiateSystemShutdownExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BaseInitiateSystemShutdownExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BaseInitiateSystemShutdownExRequest structure represents the BaseInitiateSystemShutdownEx operation request
type BaseInitiateSystemShutdownExRequest struct {
	// ServerName: The custom RPC binding handle (PREGISTRY_SERVER_NAME (section 2.2.1)).
	ServerName string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	// lpMessage: Null-terminated Unicode string that contains the message to display during
	// the shutdown waiting period. If this parameter is NULL, no message MUST be displayed.
	Message *rsp.UnicodeString `idl:"name:lpMessage;pointer:unique" json:"message"`
	// dwTimeout: Number of seconds to wait before shutting down.
	Timeout uint32 `idl:"name:dwTimeout" json:"timeout"`
	// bForceAppsClosed: If TRUE, all applications SHOULD be terminated unconditionally.
	ForceAppsClosed uint8 `idl:"name:bForceAppsClosed" json:"force_apps_closed"`
	// bRebootAfterShutdown: If TRUE, the system SHOULD shutdown and reboot. If FALSE, the
	// system SHOULD only shut down.
	RebootAfterShutdown uint8 `idl:"name:bRebootAfterShutdown" json:"reboot_after_shutdown"`
	// dwReason: Reason for initiating the shutdown (section 2.3). The dwReason SHOULD be
	// used for log entries for the shutdown event.
	Reason uint32 `idl:"name:dwReason" json:"reason"`
}

func (o *BaseInitiateSystemShutdownExRequest) xxx_ToOp(ctx context.Context) *xxx_BaseInitiateSystemShutdownExOperation {
	if o == nil {
		return &xxx_BaseInitiateSystemShutdownExOperation{}
	}
	return &xxx_BaseInitiateSystemShutdownExOperation{
		ServerName:          o.ServerName,
		Message:             o.Message,
		Timeout:             o.Timeout,
		ForceAppsClosed:     o.ForceAppsClosed,
		RebootAfterShutdown: o.RebootAfterShutdown,
		Reason:              o.Reason,
	}
}

func (o *BaseInitiateSystemShutdownExRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseInitiateSystemShutdownExOperation) {
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
func (o *BaseInitiateSystemShutdownExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *BaseInitiateSystemShutdownExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseInitiateSystemShutdownExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseInitiateSystemShutdownExResponse structure represents the BaseInitiateSystemShutdownEx operation response
type BaseInitiateSystemShutdownExResponse struct {
	// Return: The BaseInitiateSystemShutdownEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseInitiateSystemShutdownExResponse) xxx_ToOp(ctx context.Context) *xxx_BaseInitiateSystemShutdownExOperation {
	if o == nil {
		return &xxx_BaseInitiateSystemShutdownExOperation{}
	}
	return &xxx_BaseInitiateSystemShutdownExOperation{
		Return: o.Return,
	}
}

func (o *BaseInitiateSystemShutdownExResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseInitiateSystemShutdownExOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BaseInitiateSystemShutdownExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *BaseInitiateSystemShutdownExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseInitiateSystemShutdownExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
