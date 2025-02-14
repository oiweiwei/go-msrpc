package w32time

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	w32t "github.com/oiweiwei/go-msrpc/msrpc/w32t"
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
	_ = w32t.GoPackage
)

var (
	// import guard
	GoPackage = "w32t"
)

var (
	// Syntax UUID
	W32TimeSyntaxUUID = &uuid.UUID{TimeLow: 0x8fb6d884, TimeMid: 0x2388, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0x35, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xda, 0x27, 0x95}}
	// Syntax ID
	W32TimeSyntaxV4_1 = &dcerpc.SyntaxID{IfUUID: W32TimeSyntaxUUID, IfVersionMajor: 4, IfVersionMinor: 1}
)

// W32Time interface.
type W32TimeClient interface {

	// The W32TimeSync method is invoked to request that the time service immediately initiate
	// an attempt to synchronize its time. The MIDL syntax of this method is specified as
	// follows.
	//
	// Return Values: If the TimeSyncFlag_ReturnResult flag is specified, the return value
	// MUST be one of the following specific TimeSync_ReturnResult codes. Otherwise, this
	// method MUST return zero on success or an implementation-specific nonzero error code
	// on failure.<36>
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ResyncResult_Success      | Synchronization between the time provider and the caller was successful. For     |
	//	|                                      | asynchronous calls, this result does not guarantee that the server has acquired  |
	//	|                                      | a new time sample. It merely states that the synchronization attempt has been    |
	//	|                                      | successfully initiated.                                                          |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 ResyncResult_NoData       | The time service could not obtain a new time sample from the time provider.      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ResyncResult_StaleData    | The time service received data that was time stamped earlier than the last good  |
	//	|                                      | sample.                                                                          |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000003 ResyncResult_ChangeTooBig | The time service received data in which the time difference from the local clock |
	//	|                                      | was too large to trust.                                                          |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000004 ResyncResult_Shutdown     | The time service was shutting down.                                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	Sync(context.Context, *SyncRequest, ...dcerpc.CallOption) (*SyncResponse, error)

	// The W32TimeGetNetlogonServiceBits method returns information about the functionality
	// that the time service provides. The MIDL syntax of this method is specified as follows.
	//
	// Return Values: This method returns an unsigned 32-bit field that contains information
	// about the functionality that the time service provides. Multiple bits can be set
	// in the return value. Any bits not defined as follows MUST be set to zero by servers
	// and ignored by clients.
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN              |                                                                                  |
	//	|            VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000040 DS_TIMESERV_FLAG      | The time service provides a time source with which clients can synchronize using |
	//	|                                  | NTP, as specified in [RFC1305].                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000200 DS_GOOD_TIMESERV_FLAG | The time service provides a reliable time source with which clients can          |
	//	|                                  | synchronize using NTP, as specified in [RFC1305].                                |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	GetNetlogonServiceBits(context.Context, *GetNetlogonServiceBitsRequest, ...dcerpc.CallOption) (*GetNetlogonServiceBitsResponse, error)

	// The W32TimeQueryProviderStatus method returns operational information for a specified
	// time provider (either an NTP or a hardware time provider) within the time service's
	// list of time providers.<39> The MIDL syntax of this method is specified as follows.
	//
	// Return Values: This method MUST return zero on success; on failure, it returns a
	// nonzero error code. The values transmitted in this field are implementation specific.
	// All nonzero values MUST be treated as equivalent for protocol purposes.<41>
	QueryProviderStatus(context.Context, *QueryProviderStatusRequest, ...dcerpc.CallOption) (*QueryProviderStatusResponse, error)

	// The W32TimeQuerySource method returns the current time source of the time service.
	// The MIDL syntax of this method is specified as follows.
	//
	// Return Values: This method MUST return 0 on success; on failure, it returns a nonzero
	// error code. The values transmitted in this field are implementation specific. All
	// nonzero values MUST be treated as equivalent for the purposes of this protocol.<42>
	QuerySource(context.Context, *QuerySourceRequest, ...dcerpc.CallOption) (*QuerySourceResponse, error)

	// The W32TimeQueryProviderConfiguration method returns configuration data for a specific
	// time provider within the time service's list of time providers.<44> The MIDL syntax
	// of this method is specified as follows.
	//
	// Return Values: This method MUST return zero on success; on failure, it returns a
	// nonzero error code. The values transmitted in this field are implementation specific.
	// All nonzero values MUST be treated as equivalent for the purposes of this protocol.<46>
	QueryProviderConfiguration(context.Context, *QueryProviderConfigurationRequest, ...dcerpc.CallOption) (*QueryProviderConfigurationResponse, error)

	// The W32TimeQueryConfiguration method returns the configuration data of the time service.<47>
	// The MIDL syntax of this method is specified as follows.
	//
	// Return Values: This method MUST return 0 on success; on failure, it returns a nonzero
	// error code. The values transmitted in this field are implementation specific. All
	// nonzero values MUST be treated as equivalent for the purposes of this protocol.<48>
	QueryConfiguration(context.Context, *QueryConfigurationRequest, ...dcerpc.CallOption) (*QueryConfigurationResponse, error)

	// The W32TimeQueryStatus method returns the service status data of the time service.<49>
	// The MIDL syntax of this method is specified as follows.
	//
	// Return Values: This method MUST return 0 on success; on failure, it returns a nonzero
	// error code. The values transmitted in this field are implementation specific. All
	// nonzero values MUST be treated as equivalent for the purposes of this protocol.<50>
	QueryStatus(context.Context, *QueryStatusRequest, ...dcerpc.CallOption) (*QueryStatusResponse, error)

	// The W32TimeLog method is invoked to request that the time service update its logging
	// configuration.<51> The logging of the time service is implementation specific.<52>
	//
	// The MIDL syntax of this method is specified as follows.
	//
	// Return Values: This method MUST return 0 on success; on failure, it returns a nonzero
	// error code. The values transmitted in this field are implementation specific. All
	// nonzero values MUST be treated as equivalent for protocol purposes.<53>
	Log(context.Context, *LogRequest, ...dcerpc.CallOption) (*LogResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultW32TimeClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultW32TimeClient) Sync(ctx context.Context, in *SyncRequest, opts ...dcerpc.CallOption) (*SyncResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SyncResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultW32TimeClient) GetNetlogonServiceBits(ctx context.Context, in *GetNetlogonServiceBitsRequest, opts ...dcerpc.CallOption) (*GetNetlogonServiceBitsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNetlogonServiceBitsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultW32TimeClient) QueryProviderStatus(ctx context.Context, in *QueryProviderStatusRequest, opts ...dcerpc.CallOption) (*QueryProviderStatusResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryProviderStatusResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultW32TimeClient) QuerySource(ctx context.Context, in *QuerySourceRequest, opts ...dcerpc.CallOption) (*QuerySourceResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QuerySourceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultW32TimeClient) QueryProviderConfiguration(ctx context.Context, in *QueryProviderConfigurationRequest, opts ...dcerpc.CallOption) (*QueryProviderConfigurationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryProviderConfigurationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultW32TimeClient) QueryConfiguration(ctx context.Context, in *QueryConfigurationRequest, opts ...dcerpc.CallOption) (*QueryConfigurationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryConfigurationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultW32TimeClient) QueryStatus(ctx context.Context, in *QueryStatusRequest, opts ...dcerpc.CallOption) (*QueryStatusResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryStatusResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultW32TimeClient) Log(ctx context.Context, in *LogRequest, opts ...dcerpc.CallOption) (*LogResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &LogResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultW32TimeClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultW32TimeClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewW32TimeClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (W32TimeClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(W32TimeSyntaxV4_1))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultW32TimeClient{cc: cc}, nil
}

// xxx_SyncOperation structure represents the W32TimeSync operation
type xxx_SyncOperation struct {
	Wait   uint32 `idl:"name:uWait" json:"wait"`
	Flags  uint32 `idl:"name:ulFlags" json:"flags"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SyncOperation) OpNum() int { return 0 }

func (o *xxx_SyncOperation) OpName() string { return "/W32Time/v4.1/W32TimeSync" }

func (o *xxx_SyncOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SyncOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// uWait {in} (1:(uint32))
	{
		if err := w.WriteData(o.Wait); err != nil {
			return err
		}
	}
	// ulFlags {in} (1:(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SyncOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// uWait {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Wait); err != nil {
			return err
		}
	}
	// ulFlags {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SyncOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SyncOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SyncOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SyncRequest structure represents the W32TimeSync operation request
type SyncRequest struct {
	// uWait: Blocking status of the call. The value MUST be one of the following.
	//
	//	+--------------------+----------------------------------------------------------------------------------+
	//	|                    |                                                                                  |
	//	|       VALUE        |                                     MEANING                                      |
	//	|                    |                                                                                  |
	//	+--------------------+----------------------------------------------------------------------------------+
	//	+--------------------+----------------------------------------------------------------------------------+
	//	|                  0 | Method MUST return RPC_S_OK without waiting for the outcome of time              |
	//	|                    | synchronization. In this case, the final outcome of the attempt is not available |
	//	|                    | to the caller.                                                                   |
	//	+--------------------+----------------------------------------------------------------------------------+
	//	| Non-zero 0 < value | Method MUST NOT return until time synchronization is complete. The server must   |
	//	|                    | block the response to the caller until the attempt at time synchronization is    |
	//	|                    | complete, regardless of the status. An implementation can choose to respond to   |
	//	|                    | the caller after an implementation-specific timeout occurs.<34>                  |
	//	+--------------------+----------------------------------------------------------------------------------+
	Wait uint32 `idl:"name:uWait" json:"wait"`
	// ulFlags: Time synchronization behaviors.
	//
	// The following values SHOULD be mutually exclusive. When multiple values are set,
	// the value whose bit is least significant SHOULD take precedence.<35>
	//
	// Note  The TimeSyncFlag_SoftResync value MUST NOT be used in conjunction with any
	// other value in the following table.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                         |                                                                                  |
	//	|                  VALUE                  |                                     MEANING                                      |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| TimeSyncFlag_SoftResync 0x00000000      | The time service MUST synchronize itself with the currently available time       |
	//	|                                         | samples. It MUST NOT poll the network or hardware time providers for new time    |
	//	|                                         | data.                                                                            |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| TimeSyncFlag_HardResync 0x00000001      | The time service MUST discard its old time samples and MUST acquire new samples  |
	//	|                                         | from the network or hardware time providers.                                     |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| TimeSyncFlag_Rediscover 0x00000004      | Identical to the TimeSyncFlag_HardResync flag, except that the time service MUST |
	//	|                                         | attempt to discover new network time sources prior to discarding and reacquiring |
	//	|                                         | new time samples.                                                                |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| TimeSyncFlag_UpdateAndResync 0x00000008 | Identical to the TimeSyncFlag_Rediscover flag, except that prior to attempting   |
	//	|                                         | to discover new time sources, the time service MUST update its configuration.    |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| TimeSyncFlag_ForceResync 0x00000010     | Identical to the TimeSyncFlag_HardResync flag, except that it causes the         |
	//	|                                         | processing of the next time sample to ignore any phase correction boundaries     |
	//	|                                         | imposed by W32Time.                                                              |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//
	// The following value can be joined in a bitwise OR with the preceding values. If uWait
	// is set to zero, the following value MUST be ignored.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                                      |                                                                                  |
	//	|                VALUE                 |                                     MEANING                                      |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| TimeSyncFlag_ReturnResult 0x00000002 | Used only for synchronous calls. If set, the method MUST return one of the       |
	//	|                                      | following return values.                                                         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:ulFlags" json:"flags"`
}

func (o *SyncRequest) xxx_ToOp(ctx context.Context, op *xxx_SyncOperation) *xxx_SyncOperation {
	if op == nil {
		op = &xxx_SyncOperation{}
	}
	if o == nil {
		return op
	}
	o.Wait = op.Wait
	o.Flags = op.Flags
	return op
}

func (o *SyncRequest) xxx_FromOp(ctx context.Context, op *xxx_SyncOperation) {
	if o == nil {
		return
	}
	o.Wait = op.Wait
	o.Flags = op.Flags
}
func (o *SyncRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SyncRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SyncOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SyncResponse structure represents the W32TimeSync operation response
type SyncResponse struct {
	// Return: The W32TimeSync return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SyncResponse) xxx_ToOp(ctx context.Context, op *xxx_SyncOperation) *xxx_SyncOperation {
	if op == nil {
		op = &xxx_SyncOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *SyncResponse) xxx_FromOp(ctx context.Context, op *xxx_SyncOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SyncResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SyncResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SyncOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNetlogonServiceBitsOperation structure represents the W32TimeGetNetlogonServiceBits operation
type xxx_GetNetlogonServiceBitsOperation struct {
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNetlogonServiceBitsOperation) OpNum() int { return 1 }

func (o *xxx_GetNetlogonServiceBitsOperation) OpName() string {
	return "/W32Time/v4.1/W32TimeGetNetlogonServiceBits"
}

func (o *xxx_GetNetlogonServiceBitsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetlogonServiceBitsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_GetNetlogonServiceBitsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_GetNetlogonServiceBitsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetlogonServiceBitsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_GetNetlogonServiceBitsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNetlogonServiceBitsRequest structure represents the W32TimeGetNetlogonServiceBits operation request
type GetNetlogonServiceBitsRequest struct {
}

func (o *GetNetlogonServiceBitsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetNetlogonServiceBitsOperation) *xxx_GetNetlogonServiceBitsOperation {
	if op == nil {
		op = &xxx_GetNetlogonServiceBitsOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *GetNetlogonServiceBitsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNetlogonServiceBitsOperation) {
	if o == nil {
		return
	}
}
func (o *GetNetlogonServiceBitsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetNetlogonServiceBitsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNetlogonServiceBitsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNetlogonServiceBitsResponse structure represents the W32TimeGetNetlogonServiceBits operation response
type GetNetlogonServiceBitsResponse struct {
	// Return: The W32TimeGetNetlogonServiceBits return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetNetlogonServiceBitsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetNetlogonServiceBitsOperation) *xxx_GetNetlogonServiceBitsOperation {
	if op == nil {
		op = &xxx_GetNetlogonServiceBitsOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *GetNetlogonServiceBitsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNetlogonServiceBitsOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *GetNetlogonServiceBitsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetNetlogonServiceBitsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNetlogonServiceBitsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryProviderStatusOperation structure represents the W32TimeQueryProviderStatus operation
type xxx_QueryProviderStatusOperation struct {
	Flags        uint32             `idl:"name:ulFlags" json:"flags"`
	Provider     string             `idl:"name:pwszProvider;string" json:"provider"`
	ProviderInfo *w32t.ProviderInfo `idl:"name:pProviderInfo;pointer:ref" json:"provider_info"`
	Return       uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryProviderStatusOperation) OpNum() int { return 2 }

func (o *xxx_QueryProviderStatusOperation) OpName() string {
	return "/W32Time/v4.1/W32TimeQueryProviderStatus"
}

func (o *xxx_QueryProviderStatusOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryProviderStatusOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ulFlags {in} (1:(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pwszProvider {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Provider); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryProviderStatusOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ulFlags {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pwszProvider {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Provider); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryProviderStatusOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryProviderStatusOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pProviderInfo {out} (1:{pointer=ref}*(2))(2:{alias=PW32TIME_PROVIDER_INFO}*(1))(3:{alias=W32TIME_PROVIDER_INFO}(struct))
	{
		if o.ProviderInfo != nil {
			_ptr_pProviderInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ProviderInfo != nil {
					if err := o.ProviderInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&w32t.ProviderInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ProviderInfo, _ptr_pProviderInfo); err != nil {
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
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryProviderStatusOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pProviderInfo {out} (1:{pointer=ref}*(2))(2:{alias=PW32TIME_PROVIDER_INFO,pointer=ref}*(1))(3:{alias=W32TIME_PROVIDER_INFO}(struct))
	{
		_ptr_pProviderInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ProviderInfo == nil {
				o.ProviderInfo = &w32t.ProviderInfo{}
			}
			if err := o.ProviderInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pProviderInfo := func(ptr interface{}) { o.ProviderInfo = *ptr.(**w32t.ProviderInfo) }
		if err := w.ReadPointer(&o.ProviderInfo, _s_pProviderInfo, _ptr_pProviderInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryProviderStatusRequest structure represents the W32TimeQueryProviderStatus operation request
type QueryProviderStatusRequest struct {
	// ulFlags: Reserved. This parameter MUST be set to zero and MUST be ignored on receipt.
	Flags uint32 `idl:"name:ulFlags" json:"flags"`
	// pwszProvider: Name of the time provider to query. This name is implementation specific.<40>
	Provider string `idl:"name:pwszProvider;string" json:"provider"`
}

func (o *QueryProviderStatusRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryProviderStatusOperation) *xxx_QueryProviderStatusOperation {
	if op == nil {
		op = &xxx_QueryProviderStatusOperation{}
	}
	if o == nil {
		return op
	}
	o.Flags = op.Flags
	o.Provider = op.Provider
	return op
}

func (o *QueryProviderStatusRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryProviderStatusOperation) {
	if o == nil {
		return
	}
	o.Flags = op.Flags
	o.Provider = op.Provider
}
func (o *QueryProviderStatusRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryProviderStatusRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryProviderStatusOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryProviderStatusResponse structure represents the W32TimeQueryProviderStatus operation response
type QueryProviderStatusResponse struct {
	// pProviderInfo: A pointer that receives a pointer to a W32TIME_PROVIDER_INFO structure
	// containing operational information for the time provider.
	ProviderInfo *w32t.ProviderInfo `idl:"name:pProviderInfo;pointer:ref" json:"provider_info"`
	// Return: The W32TimeQueryProviderStatus return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *QueryProviderStatusResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryProviderStatusOperation) *xxx_QueryProviderStatusOperation {
	if op == nil {
		op = &xxx_QueryProviderStatusOperation{}
	}
	if o == nil {
		return op
	}
	o.ProviderInfo = op.ProviderInfo
	o.Return = op.Return
	return op
}

func (o *QueryProviderStatusResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryProviderStatusOperation) {
	if o == nil {
		return
	}
	o.ProviderInfo = op.ProviderInfo
	o.Return = op.Return
}
func (o *QueryProviderStatusResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryProviderStatusResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryProviderStatusOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QuerySourceOperation structure represents the W32TimeQuerySource operation
type xxx_QuerySourceOperation struct {
	Source string `idl:"name:pwszSource;string" json:"source"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_QuerySourceOperation) OpNum() int { return 3 }

func (o *xxx_QuerySourceOperation) OpName() string { return "/W32Time/v4.1/W32TimeQuerySource" }

func (o *xxx_QuerySourceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySourceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_QuerySourceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_QuerySourceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySourceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pwszSource {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		if o.Source != "" {
			_ptr_pwszSource := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Source); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Source, _ptr_pwszSource); err != nil {
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
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QuerySourceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pwszSource {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		_ptr_pwszSource := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Source); err != nil {
				return err
			}
			return nil
		})
		_s_pwszSource := func(ptr interface{}) { o.Source = *ptr.(*string) }
		if err := w.ReadPointer(&o.Source, _s_pwszSource, _ptr_pwszSource); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QuerySourceRequest structure represents the W32TimeQuerySource operation request
type QuerySourceRequest struct {
}

func (o *QuerySourceRequest) xxx_ToOp(ctx context.Context, op *xxx_QuerySourceOperation) *xxx_QuerySourceOperation {
	if op == nil {
		op = &xxx_QuerySourceOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *QuerySourceRequest) xxx_FromOp(ctx context.Context, op *xxx_QuerySourceOperation) {
	if o == nil {
		return
	}
}
func (o *QuerySourceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QuerySourceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QuerySourceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QuerySourceResponse structure represents the W32TimeQuerySource operation response
type QuerySourceResponse struct {
	// pwszSource: A pointer to a null-terminated string that is the name of the time source
	// that the time service is synchronizing with. If the time service is not synchronizing
	// with any time source, the string MUST be set to a null-terminated empty string. This
	// string SHOULD be either the FQDN or the IP address of the time source in the form
	// of a string, for example, "ntp1.nist.gov" or "10.0.0.1".
	Source string `idl:"name:pwszSource;string" json:"source"`
	// Return: The W32TimeQuerySource return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *QuerySourceResponse) xxx_ToOp(ctx context.Context, op *xxx_QuerySourceOperation) *xxx_QuerySourceOperation {
	if op == nil {
		op = &xxx_QuerySourceOperation{}
	}
	if o == nil {
		return op
	}
	o.Source = op.Source
	o.Return = op.Return
	return op
}

func (o *QuerySourceResponse) xxx_FromOp(ctx context.Context, op *xxx_QuerySourceOperation) {
	if o == nil {
		return
	}
	o.Source = op.Source
	o.Return = op.Return
}
func (o *QuerySourceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QuerySourceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QuerySourceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryProviderConfigurationOperation structure represents the W32TimeQueryProviderConfiguration operation
type xxx_QueryProviderConfigurationOperation struct {
	Flags                     uint32                      `idl:"name:ulFlags" json:"flags"`
	Provider                  string                      `idl:"name:pwszProvider;string" json:"provider"`
	ConfigurationProviderInfo *w32t.ConfigurationProvider `idl:"name:pConfigurationProviderInfo;pointer:ref" json:"configuration_provider_info"`
	Return                    uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryProviderConfigurationOperation) OpNum() int { return 4 }

func (o *xxx_QueryProviderConfigurationOperation) OpName() string {
	return "/W32Time/v4.1/W32TimeQueryProviderConfiguration"
}

func (o *xxx_QueryProviderConfigurationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryProviderConfigurationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ulFlags {in} (1:(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// pwszProvider {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Provider); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryProviderConfigurationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ulFlags {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// pwszProvider {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Provider); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryProviderConfigurationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryProviderConfigurationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pConfigurationProviderInfo {out} (1:{pointer=ref}*(2))(2:{alias=PW32TIME_CONFIGURATION_PROVIDER}*(1))(3:{alias=W32TIME_CONFIGURATION_PROVIDER}(struct))
	{
		if o.ConfigurationProviderInfo != nil {
			_ptr_pConfigurationProviderInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConfigurationProviderInfo != nil {
					if err := o.ConfigurationProviderInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&w32t.ConfigurationProvider{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConfigurationProviderInfo, _ptr_pConfigurationProviderInfo); err != nil {
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
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryProviderConfigurationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pConfigurationProviderInfo {out} (1:{pointer=ref}*(2))(2:{alias=PW32TIME_CONFIGURATION_PROVIDER,pointer=ref}*(1))(3:{alias=W32TIME_CONFIGURATION_PROVIDER}(struct))
	{
		_ptr_pConfigurationProviderInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConfigurationProviderInfo == nil {
				o.ConfigurationProviderInfo = &w32t.ConfigurationProvider{}
			}
			if err := o.ConfigurationProviderInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pConfigurationProviderInfo := func(ptr interface{}) { o.ConfigurationProviderInfo = *ptr.(**w32t.ConfigurationProvider) }
		if err := w.ReadPointer(&o.ConfigurationProviderInfo, _s_pConfigurationProviderInfo, _ptr_pConfigurationProviderInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryProviderConfigurationRequest structure represents the W32TimeQueryProviderConfiguration operation request
type QueryProviderConfigurationRequest struct {
	// ulFlags: Reserved. This parameter MUST be set to zero and MUST be ignored on receipt.
	Flags uint32 `idl:"name:ulFlags" json:"flags"`
	// pwszProvider: A null-terminated string that is the name of the time provider to query.
	// This name is implementation specific.<45>
	Provider string `idl:"name:pwszProvider;string" json:"provider"`
}

func (o *QueryProviderConfigurationRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryProviderConfigurationOperation) *xxx_QueryProviderConfigurationOperation {
	if op == nil {
		op = &xxx_QueryProviderConfigurationOperation{}
	}
	if o == nil {
		return op
	}
	o.Flags = op.Flags
	o.Provider = op.Provider
	return op
}

func (o *QueryProviderConfigurationRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryProviderConfigurationOperation) {
	if o == nil {
		return
	}
	o.Flags = op.Flags
	o.Provider = op.Provider
}
func (o *QueryProviderConfigurationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryProviderConfigurationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryProviderConfigurationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryProviderConfigurationResponse structure represents the W32TimeQueryProviderConfiguration operation response
type QueryProviderConfigurationResponse struct {
	// pConfigurationProviderInfo: A pointer that receives a pointer to a W32TIME_CONFIGURATION_PROVIDER
	// structure containing configuration data for the time provider.
	ConfigurationProviderInfo *w32t.ConfigurationProvider `idl:"name:pConfigurationProviderInfo;pointer:ref" json:"configuration_provider_info"`
	// Return: The W32TimeQueryProviderConfiguration return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *QueryProviderConfigurationResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryProviderConfigurationOperation) *xxx_QueryProviderConfigurationOperation {
	if op == nil {
		op = &xxx_QueryProviderConfigurationOperation{}
	}
	if o == nil {
		return op
	}
	o.ConfigurationProviderInfo = op.ConfigurationProviderInfo
	o.Return = op.Return
	return op
}

func (o *QueryProviderConfigurationResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryProviderConfigurationOperation) {
	if o == nil {
		return
	}
	o.ConfigurationProviderInfo = op.ConfigurationProviderInfo
	o.Return = op.Return
}
func (o *QueryProviderConfigurationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryProviderConfigurationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryProviderConfigurationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryConfigurationOperation structure represents the W32TimeQueryConfiguration operation
type xxx_QueryConfigurationOperation struct {
	ConfigurationInfo *w32t.ConfigurationInfo `idl:"name:pConfigurationInfo;pointer:ref" json:"configuration_info"`
	Return            uint32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryConfigurationOperation) OpNum() int { return 5 }

func (o *xxx_QueryConfigurationOperation) OpName() string {
	return "/W32Time/v4.1/W32TimeQueryConfiguration"
}

func (o *xxx_QueryConfigurationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryConfigurationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_QueryConfigurationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_QueryConfigurationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryConfigurationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pConfigurationInfo {out} (1:{pointer=ref}*(2))(2:{alias=PW32TIME_CONFIGURATION_INFO}*(1))(3:{alias=W32TIME_CONFIGURATION_INFO}(struct))
	{
		if o.ConfigurationInfo != nil {
			_ptr_pConfigurationInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ConfigurationInfo != nil {
					if err := o.ConfigurationInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&w32t.ConfigurationInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ConfigurationInfo, _ptr_pConfigurationInfo); err != nil {
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
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryConfigurationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pConfigurationInfo {out} (1:{pointer=ref}*(2))(2:{alias=PW32TIME_CONFIGURATION_INFO,pointer=ref}*(1))(3:{alias=W32TIME_CONFIGURATION_INFO}(struct))
	{
		_ptr_pConfigurationInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ConfigurationInfo == nil {
				o.ConfigurationInfo = &w32t.ConfigurationInfo{}
			}
			if err := o.ConfigurationInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pConfigurationInfo := func(ptr interface{}) { o.ConfigurationInfo = *ptr.(**w32t.ConfigurationInfo) }
		if err := w.ReadPointer(&o.ConfigurationInfo, _s_pConfigurationInfo, _ptr_pConfigurationInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryConfigurationRequest structure represents the W32TimeQueryConfiguration operation request
type QueryConfigurationRequest struct {
}

func (o *QueryConfigurationRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryConfigurationOperation) *xxx_QueryConfigurationOperation {
	if op == nil {
		op = &xxx_QueryConfigurationOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *QueryConfigurationRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryConfigurationOperation) {
	if o == nil {
		return
	}
}
func (o *QueryConfigurationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryConfigurationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryConfigurationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryConfigurationResponse structure represents the W32TimeQueryConfiguration operation response
type QueryConfigurationResponse struct {
	// pConfigurationInfo: A pointer that receives a pointer to a W32TIME_CONFIGURATION_INFO
	// structure containing configuration data for the time service.
	ConfigurationInfo *w32t.ConfigurationInfo `idl:"name:pConfigurationInfo;pointer:ref" json:"configuration_info"`
	// Return: The W32TimeQueryConfiguration return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *QueryConfigurationResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryConfigurationOperation) *xxx_QueryConfigurationOperation {
	if op == nil {
		op = &xxx_QueryConfigurationOperation{}
	}
	if o == nil {
		return op
	}
	o.ConfigurationInfo = op.ConfigurationInfo
	o.Return = op.Return
	return op
}

func (o *QueryConfigurationResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryConfigurationOperation) {
	if o == nil {
		return
	}
	o.ConfigurationInfo = op.ConfigurationInfo
	o.Return = op.Return
}
func (o *QueryConfigurationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryConfigurationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryConfigurationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryStatusOperation structure represents the W32TimeQueryStatus operation
type xxx_QueryStatusOperation struct {
	StatusInfo *w32t.StatusInfo `idl:"name:pStatusInfo;pointer:ref" json:"status_info"`
	Return     uint32           `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryStatusOperation) OpNum() int { return 6 }

func (o *xxx_QueryStatusOperation) OpName() string { return "/W32Time/v4.1/W32TimeQueryStatus" }

func (o *xxx_QueryStatusOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryStatusOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_QueryStatusOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_QueryStatusOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryStatusOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pStatusInfo {out} (1:{pointer=ref}*(2))(2:{alias=PW32TIME_STATUS_INFO}*(1))(3:{alias=W32TIME_STATUS_INFO}(struct))
	{
		if o.StatusInfo != nil {
			_ptr_pStatusInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.StatusInfo != nil {
					if err := o.StatusInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&w32t.StatusInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.StatusInfo, _ptr_pStatusInfo); err != nil {
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
	// Return {out} (1:(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryStatusOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pStatusInfo {out} (1:{pointer=ref}*(2))(2:{alias=PW32TIME_STATUS_INFO,pointer=ref}*(1))(3:{alias=W32TIME_STATUS_INFO}(struct))
	{
		_ptr_pStatusInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.StatusInfo == nil {
				o.StatusInfo = &w32t.StatusInfo{}
			}
			if err := o.StatusInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pStatusInfo := func(ptr interface{}) { o.StatusInfo = *ptr.(**w32t.StatusInfo) }
		if err := w.ReadPointer(&o.StatusInfo, _s_pStatusInfo, _ptr_pStatusInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryStatusRequest structure represents the W32TimeQueryStatus operation request
type QueryStatusRequest struct {
}

func (o *QueryStatusRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryStatusOperation) *xxx_QueryStatusOperation {
	if op == nil {
		op = &xxx_QueryStatusOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *QueryStatusRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryStatusOperation) {
	if o == nil {
		return
	}
}
func (o *QueryStatusRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryStatusRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryStatusOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryStatusResponse structure represents the W32TimeQueryStatus operation response
type QueryStatusResponse struct {
	// pStatusInfo: A pointer that receives a pointer to a W32TIME_STATUS_INFO structure
	// containing status data for the time service.
	StatusInfo *w32t.StatusInfo `idl:"name:pStatusInfo;pointer:ref" json:"status_info"`
	// Return: The W32TimeQueryStatus return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *QueryStatusResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryStatusOperation) *xxx_QueryStatusOperation {
	if op == nil {
		op = &xxx_QueryStatusOperation{}
	}
	if o == nil {
		return op
	}
	o.StatusInfo = op.StatusInfo
	o.Return = op.Return
	return op
}

func (o *QueryStatusResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryStatusOperation) {
	if o == nil {
		return
	}
	o.StatusInfo = op.StatusInfo
	o.Return = op.Return
}
func (o *QueryStatusResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryStatusResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryStatusOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_LogOperation structure represents the W32TimeLog operation
type xxx_LogOperation struct {
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_LogOperation) OpNum() int { return 7 }

func (o *xxx_LogOperation) OpName() string { return "/W32Time/v4.1/W32TimeLog" }

func (o *xxx_LogOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LogOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_LogOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_LogOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_LogOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_LogOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// LogRequest structure represents the W32TimeLog operation request
type LogRequest struct {
}

func (o *LogRequest) xxx_ToOp(ctx context.Context, op *xxx_LogOperation) *xxx_LogOperation {
	if op == nil {
		op = &xxx_LogOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *LogRequest) xxx_FromOp(ctx context.Context, op *xxx_LogOperation) {
	if o == nil {
		return
	}
}
func (o *LogRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *LogRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LogOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// LogResponse structure represents the W32TimeLog operation response
type LogResponse struct {
	// Return: The W32TimeLog return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *LogResponse) xxx_ToOp(ctx context.Context, op *xxx_LogOperation) *xxx_LogOperation {
	if op == nil {
		op = &xxx_LogOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *LogResponse) xxx_FromOp(ctx context.Context, op *xxx_LogOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *LogResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *LogResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_LogOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
