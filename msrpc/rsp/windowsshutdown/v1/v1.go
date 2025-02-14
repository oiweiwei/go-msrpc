package windowsshutdown

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
	WindowsShutdownSyntaxUUID = &uuid.UUID{TimeLow: 0xd95afe70, TimeMid: 0xa6d5, TimeHiAndVersion: 0x4259, ClockSeqHiAndReserved: 0x82, ClockSeqLow: 0x2e, Node: [6]uint8{0x2c, 0x84, 0xda, 0x1d, 0xdb, 0xd}}
	// Syntax ID
	WindowsShutdownSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: WindowsShutdownSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// WindowsShutdown interface.
type WindowsShutdownClient interface {

	// The WsdrInitiateShutdown method is used to initiate the shutdown of the remote computer.<14>
	//
	// Return Values: The method returns ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it returns a nonzero error code.
	InitiateShutdown(context.Context, *InitiateShutdownRequest, ...dcerpc.CallOption) (*InitiateShutdownResponse, error)

	// The WsdrAbortShutdown method is used to terminate the shutdown of the remote computer
	// within the waiting period.<15>
	//
	// Return Values: The method returns ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it returns a nonzero error code.
	AbortShutdown(context.Context, *AbortShutdownRequest, ...dcerpc.CallOption) (*AbortShutdownResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultWindowsShutdownClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultWindowsShutdownClient) InitiateShutdown(ctx context.Context, in *InitiateShutdownRequest, opts ...dcerpc.CallOption) (*InitiateShutdownResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &InitiateShutdownResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWindowsShutdownClient) AbortShutdown(ctx context.Context, in *AbortShutdownRequest, opts ...dcerpc.CallOption) (*AbortShutdownResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AbortShutdownResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWindowsShutdownClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultWindowsShutdownClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewWindowsShutdownClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (WindowsShutdownClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(WindowsShutdownSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultWindowsShutdownClient{cc: cc}, nil
}

// xxx_InitiateShutdownOperation structure represents the WsdrInitiateShutdown operation
type xxx_InitiateShutdownOperation struct {
	Message      *rsp.UnicodeString `idl:"name:lpMessage;pointer:unique" json:"message"`
	GracePeriod  uint32             `idl:"name:dwGracePeriod" json:"grace_period"`
	ShudownFlags uint32             `idl:"name:dwShudownFlags" json:"shudown_flags"`
	Reason       uint32             `idl:"name:dwReason" json:"reason"`
	ClientHint   *rsp.UnicodeString `idl:"name:lpClientHint;pointer:unique" json:"client_hint"`
	Return       uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_InitiateShutdownOperation) OpNum() int { return 0 }

func (o *xxx_InitiateShutdownOperation) OpName() string {
	return "/WindowsShutdown/v1/WsdrInitiateShutdown"
}

func (o *xxx_InitiateShutdownOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitiateShutdownOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
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
	// dwGracePeriod {in} (1:(uint32))
	{
		if err := w.WriteData(o.GracePeriod); err != nil {
			return err
		}
	}
	// dwShudownFlags {in} (1:(uint32))
	{
		if err := w.WriteData(o.ShudownFlags); err != nil {
			return err
		}
	}
	// dwReason {in} (1:(uint32))
	{
		if err := w.WriteData(o.Reason); err != nil {
			return err
		}
	}
	// lpClientHint {in} (1:{pointer=unique, alias=PREG_UNICODE_STRING}*(1))(2:{alias=REG_UNICODE_STRING}(struct))
	{
		if o.ClientHint != nil {
			_ptr_lpClientHint := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ClientHint != nil {
					if err := o.ClientHint.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&rsp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientHint, _ptr_lpClientHint); err != nil {
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

func (o *xxx_InitiateShutdownOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
	// dwGracePeriod {in} (1:(uint32))
	{
		if err := w.ReadData(&o.GracePeriod); err != nil {
			return err
		}
	}
	// dwShudownFlags {in} (1:(uint32))
	{
		if err := w.ReadData(&o.ShudownFlags); err != nil {
			return err
		}
	}
	// dwReason {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Reason); err != nil {
			return err
		}
	}
	// lpClientHint {in} (1:{pointer=unique, alias=PREG_UNICODE_STRING}*(1))(2:{alias=REG_UNICODE_STRING}(struct))
	{
		_ptr_lpClientHint := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ClientHint == nil {
				o.ClientHint = &rsp.UnicodeString{}
			}
			if err := o.ClientHint.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpClientHint := func(ptr interface{}) { o.ClientHint = *ptr.(**rsp.UnicodeString) }
		if err := w.ReadPointer(&o.ClientHint, _s_lpClientHint, _ptr_lpClientHint); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitiateShutdownOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_InitiateShutdownOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_InitiateShutdownOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// InitiateShutdownRequest structure represents the WsdrInitiateShutdown operation request
type InitiateShutdownRequest struct {
	// lpMessage: Null-terminated Unicode string that contains the message to display during
	// the shutdown waiting period. If this parameter is NULL, no message MUST be displayed.
	Message *rsp.UnicodeString `idl:"name:lpMessage;pointer:unique" json:"message"`
	// dwGracePeriod: Number of seconds to wait before shutting down.
	GracePeriod uint32 `idl:"name:dwGracePeriod" json:"grace_period"`
	// dwShudownFlags: A set of bit flags in little-endian format used as a mask to indicate
	// shutdown options. The value is constructed from zero or more bit flags from the following
	// table, with the exception that flag "B" cannot be combined with "C" or "D".
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | G | F | E | D | C | B | 0 | A |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//
	// The bits are defined as follows.
	//
	//	+--------------+----------------------------------------------------------------------------------+
	//	|              |                                                                                  |
	//	|    VALUE     |                                     MEANING                                      |
	//	|              |                                                                                  |
	//	+--------------+----------------------------------------------------------------------------------+
	//	+--------------+----------------------------------------------------------------------------------+
	//	| A 0x00000001 | All applications SHOULD be terminated unconditionally. An alternate for this     |
	//	|              | field is SHUTDOWN_FORCE_OTHERS.                                                  |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| B 0x00000004 | Restart computer. Cannot be used with "C" or "D". An alternate name for this     |
	//	|              | field is SHUTDOWN_RESTART.                                                       |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| C 0x00000008 | The shutdown SHOULD turn off the computer. Cannot be used with "B" or "D". An    |
	//	|              | alternate name for this field is SHUTDOWN_POWEROFF.                              |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| D 0x00000010 | The shutdown SHOULD leave the computer powered but SHOULD NOT cause a            |
	//	|              | reboot. Cannot be used with "B" or "C". An alternate name for this field is      |
	//	|              | SHUTDOWN_NOREBOOT.                                                               |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| E 0x00000020 | If a shutdown is currently in progress, setting this bit on a subsequent         |
	//	|              | shutdown request SHOULD cause the ongoing request's waiting period to be ignored |
	//	|              | and SHOULD cause an immediate shutdown. An alternate name for this field is      |
	//	|              | SHUTDOWN_GRACE_OVERRIDE.                                                         |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| F 0x00000040 | The shutdown SHOULD install pending software updates before proceeding. An       |
	//	|              | alternate name for this field is SHUTDOWN_INSTALL_UPDATES.                       |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| G 0x00000080 | The shutdown SHOULD restart the computer and then restart any applications       |
	//	|              | that have registered for restart. An alternate name for this field is            |
	//	|              | SHUTDOWN_RESTARTAPPS.                                                            |
	//	+--------------+----------------------------------------------------------------------------------+
	ShudownFlags uint32 `idl:"name:dwShudownFlags" json:"shudown_flags"`
	// dwReason: Reason for initiating the shutdown (section 2.3). The dwReason SHOULD be
	// used for log entries for the shutdown event.
	Reason uint32 `idl:"name:dwReason" json:"reason"`
	// lpClientHint: Used only for diagnostic purposes (logging the image file name of the
	// process initiating a shutdown).
	ClientHint *rsp.UnicodeString `idl:"name:lpClientHint;pointer:unique" json:"client_hint"`
}

func (o *InitiateShutdownRequest) xxx_ToOp(ctx context.Context, op *xxx_InitiateShutdownOperation) *xxx_InitiateShutdownOperation {
	if op == nil {
		op = &xxx_InitiateShutdownOperation{}
	}
	if o == nil {
		return op
	}
	op.Message = o.Message
	op.GracePeriod = o.GracePeriod
	op.ShudownFlags = o.ShudownFlags
	op.Reason = o.Reason
	op.ClientHint = o.ClientHint
	return op
}

func (o *InitiateShutdownRequest) xxx_FromOp(ctx context.Context, op *xxx_InitiateShutdownOperation) {
	if o == nil {
		return
	}
	o.Message = op.Message
	o.GracePeriod = op.GracePeriod
	o.ShudownFlags = op.ShudownFlags
	o.Reason = op.Reason
	o.ClientHint = op.ClientHint
}
func (o *InitiateShutdownRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *InitiateShutdownRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitiateShutdownOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// InitiateShutdownResponse structure represents the WsdrInitiateShutdown operation response
type InitiateShutdownResponse struct {
	// Return: The WsdrInitiateShutdown return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *InitiateShutdownResponse) xxx_ToOp(ctx context.Context, op *xxx_InitiateShutdownOperation) *xxx_InitiateShutdownOperation {
	if op == nil {
		op = &xxx_InitiateShutdownOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *InitiateShutdownResponse) xxx_FromOp(ctx context.Context, op *xxx_InitiateShutdownOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *InitiateShutdownResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *InitiateShutdownResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_InitiateShutdownOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AbortShutdownOperation structure represents the WsdrAbortShutdown operation
type xxx_AbortShutdownOperation struct {
	ClientHint *rsp.UnicodeString `idl:"name:lpClientHint;pointer:unique" json:"client_hint"`
	Return     uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_AbortShutdownOperation) OpNum() int { return 1 }

func (o *xxx_AbortShutdownOperation) OpName() string { return "/WindowsShutdown/v1/WsdrAbortShutdown" }

func (o *xxx_AbortShutdownOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortShutdownOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpClientHint {in} (1:{pointer=unique, alias=PREG_UNICODE_STRING}*(1))(2:{alias=REG_UNICODE_STRING}(struct))
	{
		if o.ClientHint != nil {
			_ptr_lpClientHint := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ClientHint != nil {
					if err := o.ClientHint.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&rsp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientHint, _ptr_lpClientHint); err != nil {
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

func (o *xxx_AbortShutdownOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpClientHint {in} (1:{pointer=unique, alias=PREG_UNICODE_STRING}*(1))(2:{alias=REG_UNICODE_STRING}(struct))
	{
		_ptr_lpClientHint := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ClientHint == nil {
				o.ClientHint = &rsp.UnicodeString{}
			}
			if err := o.ClientHint.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpClientHint := func(ptr interface{}) { o.ClientHint = *ptr.(**rsp.UnicodeString) }
		if err := w.ReadPointer(&o.ClientHint, _s_lpClientHint, _ptr_lpClientHint); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortShutdownOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortShutdownOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_AbortShutdownOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AbortShutdownRequest structure represents the WsdrAbortShutdown operation request
type AbortShutdownRequest struct {
	// lpClientHint: Used only for diagnostic purposes (logging the image file name of the
	// process canceling a shutdown).
	ClientHint *rsp.UnicodeString `idl:"name:lpClientHint;pointer:unique" json:"client_hint"`
}

func (o *AbortShutdownRequest) xxx_ToOp(ctx context.Context, op *xxx_AbortShutdownOperation) *xxx_AbortShutdownOperation {
	if op == nil {
		op = &xxx_AbortShutdownOperation{}
	}
	if o == nil {
		return op
	}
	op.ClientHint = o.ClientHint
	return op
}

func (o *AbortShutdownRequest) xxx_FromOp(ctx context.Context, op *xxx_AbortShutdownOperation) {
	if o == nil {
		return
	}
	o.ClientHint = op.ClientHint
}
func (o *AbortShutdownRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AbortShutdownRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AbortShutdownOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AbortShutdownResponse structure represents the WsdrAbortShutdown operation response
type AbortShutdownResponse struct {
	// Return: The WsdrAbortShutdown return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AbortShutdownResponse) xxx_ToOp(ctx context.Context, op *xxx_AbortShutdownOperation) *xxx_AbortShutdownOperation {
	if op == nil {
		op = &xxx_AbortShutdownOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *AbortShutdownResponse) xxx_FromOp(ctx context.Context, op *xxx_AbortShutdownOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AbortShutdownResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AbortShutdownResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AbortShutdownOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
