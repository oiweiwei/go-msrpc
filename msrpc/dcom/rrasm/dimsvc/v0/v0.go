package dimsvc

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	rrasm "github.com/oiweiwei/go-msrpc/msrpc/dcom/rrasm"
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
	_ = rrasm.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/rrasm"
)

var (
	// Syntax UUID
	DimsvcSyntaxUUID = &uuid.UUID{TimeLow: 0x8f09f000, TimeMid: 0xb7ed, TimeHiAndVersion: 0x11ce, ClockSeqHiAndReserved: 0xbb, ClockSeqLow: 0xd2, Node: [6]uint8{0x0, 0x0, 0x1a, 0x18, 0x1c, 0xad}}
	// Syntax ID
	DimsvcSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: DimsvcSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// dimsvc interface.
type DimsvcClient interface {
	ServerGetInfo(context.Context, *ServerGetInfoRequest, ...dcerpc.CallOption) (*ServerGetInfoResponse, error)

	ConnectionEnum(context.Context, *ConnectionEnumRequest, ...dcerpc.CallOption) (*ConnectionEnumResponse, error)

	ConnectionGetInfo(context.Context, *ConnectionGetInfoRequest, ...dcerpc.CallOption) (*ConnectionGetInfoResponse, error)

	ConnectionClearStats(context.Context, *ConnectionClearStatsRequest, ...dcerpc.CallOption) (*ConnectionClearStatsResponse, error)

	PortEnum(context.Context, *PortEnumRequest, ...dcerpc.CallOption) (*PortEnumResponse, error)

	PortGetInfo(context.Context, *PortGetInfoRequest, ...dcerpc.CallOption) (*PortGetInfoResponse, error)

	PortClearStats(context.Context, *PortClearStatsRequest, ...dcerpc.CallOption) (*PortClearStatsResponse, error)

	PortReset(context.Context, *PortResetRequest, ...dcerpc.CallOption) (*PortResetResponse, error)

	PortDisconnect(context.Context, *PortDisconnectRequest, ...dcerpc.CallOption) (*PortDisconnectResponse, error)

	RouterInterfaceTransportSetGlobalInfo(context.Context, *RouterInterfaceTransportSetGlobalInfoRequest, ...dcerpc.CallOption) (*RouterInterfaceTransportSetGlobalInfoResponse, error)

	RouterInterfaceTransportGetGlobalInfo(context.Context, *RouterInterfaceTransportGetGlobalInfoRequest, ...dcerpc.CallOption) (*RouterInterfaceTransportGetGlobalInfoResponse, error)

	RouterInterfaceGetHandle(context.Context, *RouterInterfaceGetHandleRequest, ...dcerpc.CallOption) (*RouterInterfaceGetHandleResponse, error)

	RouterInterfaceCreate(context.Context, *RouterInterfaceCreateRequest, ...dcerpc.CallOption) (*RouterInterfaceCreateResponse, error)

	RouterInterfaceGetInfo(context.Context, *RouterInterfaceGetInfoRequest, ...dcerpc.CallOption) (*RouterInterfaceGetInfoResponse, error)

	RouterInterfaceSetInfo(context.Context, *RouterInterfaceSetInfoRequest, ...dcerpc.CallOption) (*RouterInterfaceSetInfoResponse, error)

	RouterInterfaceDelete(context.Context, *RouterInterfaceDeleteRequest, ...dcerpc.CallOption) (*RouterInterfaceDeleteResponse, error)

	RouterInterfaceTransportRemove(context.Context, *RouterInterfaceTransportRemoveRequest, ...dcerpc.CallOption) (*RouterInterfaceTransportRemoveResponse, error)

	RouterInterfaceTransportAdd(context.Context, *RouterInterfaceTransportAddRequest, ...dcerpc.CallOption) (*RouterInterfaceTransportAddResponse, error)

	RouterInterfaceTransportGetInfo(context.Context, *RouterInterfaceTransportGetInfoRequest, ...dcerpc.CallOption) (*RouterInterfaceTransportGetInfoResponse, error)

	RouterInterfaceTransportSetInfo(context.Context, *RouterInterfaceTransportSetInfoRequest, ...dcerpc.CallOption) (*RouterInterfaceTransportSetInfoResponse, error)

	RouterInterfaceEnum(context.Context, *RouterInterfaceEnumRequest, ...dcerpc.CallOption) (*RouterInterfaceEnumResponse, error)

	RouterInterfaceConnect(context.Context, *RouterInterfaceConnectRequest, ...dcerpc.CallOption) (*RouterInterfaceConnectResponse, error)

	RouterInterfaceDisconnect(context.Context, *RouterInterfaceDisconnectRequest, ...dcerpc.CallOption) (*RouterInterfaceDisconnectResponse, error)

	RouterInterfaceUpdateRoutes(context.Context, *RouterInterfaceUpdateRoutesRequest, ...dcerpc.CallOption) (*RouterInterfaceUpdateRoutesResponse, error)

	RouterInterfaceQueryUpdateResult(context.Context, *RouterInterfaceQueryUpdateResultRequest, ...dcerpc.CallOption) (*RouterInterfaceQueryUpdateResultResponse, error)

	RouterInterfaceUpdatePhonebookInfo(context.Context, *RouterInterfaceUpdatePhonebookInfoRequest, ...dcerpc.CallOption) (*RouterInterfaceUpdatePhonebookInfoResponse, error)

	MIBEntryCreate(context.Context, *MIBEntryCreateRequest, ...dcerpc.CallOption) (*MIBEntryCreateResponse, error)

	MIBEntryDelete(context.Context, *MIBEntryDeleteRequest, ...dcerpc.CallOption) (*MIBEntryDeleteResponse, error)

	MIBEntrySet(context.Context, *MIBEntrySetRequest, ...dcerpc.CallOption) (*MIBEntrySetResponse, error)

	MIBEntryGet(context.Context, *MIBEntryGetRequest, ...dcerpc.CallOption) (*MIBEntryGetResponse, error)

	MIBEntryGetFirst(context.Context, *MIBEntryGetFirstRequest, ...dcerpc.CallOption) (*MIBEntryGetFirstResponse, error)

	MIBEntryGetNext(context.Context, *MIBEntryGetNextRequest, ...dcerpc.CallOption) (*MIBEntryGetNextResponse, error)

	MIBGetTrapInfo(context.Context, *MIBGetTrapInfoRequest, ...dcerpc.CallOption) (*MIBGetTrapInfoResponse, error)

	MIBSetTrapInfo(context.Context, *MIBSetTrapInfoRequest, ...dcerpc.CallOption) (*MIBSetTrapInfoResponse, error)

	ConnectionNotification(context.Context, *ConnectionNotificationRequest, ...dcerpc.CallOption) (*ConnectionNotificationResponse, error)

	SendUserMessage(context.Context, *SendUserMessageRequest, ...dcerpc.CallOption) (*SendUserMessageResponse, error)

	RouterDeviceEnum(context.Context, *RouterDeviceEnumRequest, ...dcerpc.CallOption) (*RouterDeviceEnumResponse, error)

	RouterInterfaceTransportCreate(context.Context, *RouterInterfaceTransportCreateRequest, ...dcerpc.CallOption) (*RouterInterfaceTransportCreateResponse, error)

	RouterInterfaceDeviceGetInfo(context.Context, *RouterInterfaceDeviceGetInfoRequest, ...dcerpc.CallOption) (*RouterInterfaceDeviceGetInfoResponse, error)

	RouterInterfaceDeviceSetInfo(context.Context, *RouterInterfaceDeviceSetInfoRequest, ...dcerpc.CallOption) (*RouterInterfaceDeviceSetInfoResponse, error)

	RouterInterfaceSetCredentialsEx(context.Context, *RouterInterfaceSetCredentialsExRequest, ...dcerpc.CallOption) (*RouterInterfaceSetCredentialsExResponse, error)

	RouterInterfaceGetCredentialsEx(context.Context, *RouterInterfaceGetCredentialsExRequest, ...dcerpc.CallOption) (*RouterInterfaceGetCredentialsExResponse, error)

	ConnectionRemoveQuarantine(context.Context, *ConnectionRemoveQuarantineRequest, ...dcerpc.CallOption) (*ConnectionRemoveQuarantineResponse, error)

	ServerSetInfo(context.Context, *ServerSetInfoRequest, ...dcerpc.CallOption) (*ServerSetInfoResponse, error)

	ServerGetInfoEx(context.Context, *ServerGetInfoExRequest, ...dcerpc.CallOption) (*ServerGetInfoExResponse, error)

	ConnectionEnumEx(context.Context, *ConnectionEnumExRequest, ...dcerpc.CallOption) (*ConnectionEnumExResponse, error)

	ConnectionGetInfoEx(context.Context, *ConnectionGetInfoExRequest, ...dcerpc.CallOption) (*ConnectionGetInfoExResponse, error)

	ServerSetInfoEx(context.Context, *ServerSetInfoExRequest, ...dcerpc.CallOption) (*ServerSetInfoExResponse, error)

	UpdateConnection(context.Context, *UpdateConnectionRequest, ...dcerpc.CallOption) (*UpdateConnectionResponse, error)

	RouterInterfaceSetCredentialsLocal(context.Context, *RouterInterfaceSetCredentialsLocalRequest, ...dcerpc.CallOption) (*RouterInterfaceSetCredentialsLocalResponse, error)

	RouterInterfaceGetCredentialsLocal(context.Context, *RouterInterfaceGetCredentialsLocalRequest, ...dcerpc.CallOption) (*RouterInterfaceGetCredentialsLocalResponse, error)

	RouterInterfaceGetCustomInfoEx(context.Context, *RouterInterfaceGetCustomInfoExRequest, ...dcerpc.CallOption) (*RouterInterfaceGetCustomInfoExResponse, error)

	RouterInterfaceSetCustomInfoEx(context.Context, *RouterInterfaceSetCustomInfoExRequest, ...dcerpc.CallOption) (*RouterInterfaceSetCustomInfoExResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultDimsvcClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultDimsvcClient) ServerGetInfo(ctx context.Context, in *ServerGetInfoRequest, opts ...dcerpc.CallOption) (*ServerGetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ServerGetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) ConnectionEnum(ctx context.Context, in *ConnectionEnumRequest, opts ...dcerpc.CallOption) (*ConnectionEnumResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConnectionEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) ConnectionGetInfo(ctx context.Context, in *ConnectionGetInfoRequest, opts ...dcerpc.CallOption) (*ConnectionGetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConnectionGetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) ConnectionClearStats(ctx context.Context, in *ConnectionClearStatsRequest, opts ...dcerpc.CallOption) (*ConnectionClearStatsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConnectionClearStatsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) PortEnum(ctx context.Context, in *PortEnumRequest, opts ...dcerpc.CallOption) (*PortEnumResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PortEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) PortGetInfo(ctx context.Context, in *PortGetInfoRequest, opts ...dcerpc.CallOption) (*PortGetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PortGetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) PortClearStats(ctx context.Context, in *PortClearStatsRequest, opts ...dcerpc.CallOption) (*PortClearStatsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PortClearStatsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) PortReset(ctx context.Context, in *PortResetRequest, opts ...dcerpc.CallOption) (*PortResetResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PortResetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) PortDisconnect(ctx context.Context, in *PortDisconnectRequest, opts ...dcerpc.CallOption) (*PortDisconnectResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PortDisconnectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceTransportSetGlobalInfo(ctx context.Context, in *RouterInterfaceTransportSetGlobalInfoRequest, opts ...dcerpc.CallOption) (*RouterInterfaceTransportSetGlobalInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceTransportSetGlobalInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceTransportGetGlobalInfo(ctx context.Context, in *RouterInterfaceTransportGetGlobalInfoRequest, opts ...dcerpc.CallOption) (*RouterInterfaceTransportGetGlobalInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceTransportGetGlobalInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceGetHandle(ctx context.Context, in *RouterInterfaceGetHandleRequest, opts ...dcerpc.CallOption) (*RouterInterfaceGetHandleResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceGetHandleResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceCreate(ctx context.Context, in *RouterInterfaceCreateRequest, opts ...dcerpc.CallOption) (*RouterInterfaceCreateResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceCreateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceGetInfo(ctx context.Context, in *RouterInterfaceGetInfoRequest, opts ...dcerpc.CallOption) (*RouterInterfaceGetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceGetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceSetInfo(ctx context.Context, in *RouterInterfaceSetInfoRequest, opts ...dcerpc.CallOption) (*RouterInterfaceSetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceSetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceDelete(ctx context.Context, in *RouterInterfaceDeleteRequest, opts ...dcerpc.CallOption) (*RouterInterfaceDeleteResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceDeleteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceTransportRemove(ctx context.Context, in *RouterInterfaceTransportRemoveRequest, opts ...dcerpc.CallOption) (*RouterInterfaceTransportRemoveResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceTransportRemoveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceTransportAdd(ctx context.Context, in *RouterInterfaceTransportAddRequest, opts ...dcerpc.CallOption) (*RouterInterfaceTransportAddResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceTransportAddResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceTransportGetInfo(ctx context.Context, in *RouterInterfaceTransportGetInfoRequest, opts ...dcerpc.CallOption) (*RouterInterfaceTransportGetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceTransportGetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceTransportSetInfo(ctx context.Context, in *RouterInterfaceTransportSetInfoRequest, opts ...dcerpc.CallOption) (*RouterInterfaceTransportSetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceTransportSetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceEnum(ctx context.Context, in *RouterInterfaceEnumRequest, opts ...dcerpc.CallOption) (*RouterInterfaceEnumResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceConnect(ctx context.Context, in *RouterInterfaceConnectRequest, opts ...dcerpc.CallOption) (*RouterInterfaceConnectResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceConnectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceDisconnect(ctx context.Context, in *RouterInterfaceDisconnectRequest, opts ...dcerpc.CallOption) (*RouterInterfaceDisconnectResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceDisconnectResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceUpdateRoutes(ctx context.Context, in *RouterInterfaceUpdateRoutesRequest, opts ...dcerpc.CallOption) (*RouterInterfaceUpdateRoutesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceUpdateRoutesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceQueryUpdateResult(ctx context.Context, in *RouterInterfaceQueryUpdateResultRequest, opts ...dcerpc.CallOption) (*RouterInterfaceQueryUpdateResultResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceQueryUpdateResultResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceUpdatePhonebookInfo(ctx context.Context, in *RouterInterfaceUpdatePhonebookInfoRequest, opts ...dcerpc.CallOption) (*RouterInterfaceUpdatePhonebookInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceUpdatePhonebookInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) MIBEntryCreate(ctx context.Context, in *MIBEntryCreateRequest, opts ...dcerpc.CallOption) (*MIBEntryCreateResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MIBEntryCreateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) MIBEntryDelete(ctx context.Context, in *MIBEntryDeleteRequest, opts ...dcerpc.CallOption) (*MIBEntryDeleteResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MIBEntryDeleteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) MIBEntrySet(ctx context.Context, in *MIBEntrySetRequest, opts ...dcerpc.CallOption) (*MIBEntrySetResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MIBEntrySetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) MIBEntryGet(ctx context.Context, in *MIBEntryGetRequest, opts ...dcerpc.CallOption) (*MIBEntryGetResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MIBEntryGetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) MIBEntryGetFirst(ctx context.Context, in *MIBEntryGetFirstRequest, opts ...dcerpc.CallOption) (*MIBEntryGetFirstResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MIBEntryGetFirstResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) MIBEntryGetNext(ctx context.Context, in *MIBEntryGetNextRequest, opts ...dcerpc.CallOption) (*MIBEntryGetNextResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MIBEntryGetNextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) MIBGetTrapInfo(ctx context.Context, in *MIBGetTrapInfoRequest, opts ...dcerpc.CallOption) (*MIBGetTrapInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MIBGetTrapInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) MIBSetTrapInfo(ctx context.Context, in *MIBSetTrapInfoRequest, opts ...dcerpc.CallOption) (*MIBSetTrapInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MIBSetTrapInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) ConnectionNotification(ctx context.Context, in *ConnectionNotificationRequest, opts ...dcerpc.CallOption) (*ConnectionNotificationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConnectionNotificationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) SendUserMessage(ctx context.Context, in *SendUserMessageRequest, opts ...dcerpc.CallOption) (*SendUserMessageResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SendUserMessageResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterDeviceEnum(ctx context.Context, in *RouterDeviceEnumRequest, opts ...dcerpc.CallOption) (*RouterDeviceEnumResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterDeviceEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceTransportCreate(ctx context.Context, in *RouterInterfaceTransportCreateRequest, opts ...dcerpc.CallOption) (*RouterInterfaceTransportCreateResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceTransportCreateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceDeviceGetInfo(ctx context.Context, in *RouterInterfaceDeviceGetInfoRequest, opts ...dcerpc.CallOption) (*RouterInterfaceDeviceGetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceDeviceGetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceDeviceSetInfo(ctx context.Context, in *RouterInterfaceDeviceSetInfoRequest, opts ...dcerpc.CallOption) (*RouterInterfaceDeviceSetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceDeviceSetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceSetCredentialsEx(ctx context.Context, in *RouterInterfaceSetCredentialsExRequest, opts ...dcerpc.CallOption) (*RouterInterfaceSetCredentialsExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceSetCredentialsExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceGetCredentialsEx(ctx context.Context, in *RouterInterfaceGetCredentialsExRequest, opts ...dcerpc.CallOption) (*RouterInterfaceGetCredentialsExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceGetCredentialsExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) ConnectionRemoveQuarantine(ctx context.Context, in *ConnectionRemoveQuarantineRequest, opts ...dcerpc.CallOption) (*ConnectionRemoveQuarantineResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConnectionRemoveQuarantineResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) ServerSetInfo(ctx context.Context, in *ServerSetInfoRequest, opts ...dcerpc.CallOption) (*ServerSetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ServerSetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) ServerGetInfoEx(ctx context.Context, in *ServerGetInfoExRequest, opts ...dcerpc.CallOption) (*ServerGetInfoExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ServerGetInfoExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) ConnectionEnumEx(ctx context.Context, in *ConnectionEnumExRequest, opts ...dcerpc.CallOption) (*ConnectionEnumExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConnectionEnumExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) ConnectionGetInfoEx(ctx context.Context, in *ConnectionGetInfoExRequest, opts ...dcerpc.CallOption) (*ConnectionGetInfoExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConnectionGetInfoExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) ServerSetInfoEx(ctx context.Context, in *ServerSetInfoExRequest, opts ...dcerpc.CallOption) (*ServerSetInfoExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ServerSetInfoExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) UpdateConnection(ctx context.Context, in *UpdateConnectionRequest, opts ...dcerpc.CallOption) (*UpdateConnectionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UpdateConnectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceSetCredentialsLocal(ctx context.Context, in *RouterInterfaceSetCredentialsLocalRequest, opts ...dcerpc.CallOption) (*RouterInterfaceSetCredentialsLocalResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceSetCredentialsLocalResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceGetCredentialsLocal(ctx context.Context, in *RouterInterfaceGetCredentialsLocalRequest, opts ...dcerpc.CallOption) (*RouterInterfaceGetCredentialsLocalResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceGetCredentialsLocalResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceGetCustomInfoEx(ctx context.Context, in *RouterInterfaceGetCustomInfoExRequest, opts ...dcerpc.CallOption) (*RouterInterfaceGetCustomInfoExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceGetCustomInfoExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RouterInterfaceSetCustomInfoEx(ctx context.Context, in *RouterInterfaceSetCustomInfoExRequest, opts ...dcerpc.CallOption) (*RouterInterfaceSetCustomInfoExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RouterInterfaceSetCustomInfoExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultDimsvcClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewDimsvcClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (DimsvcClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(DimsvcSyntaxV0_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultDimsvcClient{cc: cc}, nil
}

// xxx_ServerGetInfoOperation structure represents the RMprAdminServerGetInfo operation
type xxx_ServerGetInfoOperation struct {
	Level  uint32                      `idl:"name:dwLevel" json:"level"`
	Info   *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Return uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_ServerGetInfoOperation) OpNum() int { return 0 }

func (o *xxx_ServerGetInfoOperation) OpName() string { return "/dimsvc/v0/RMprAdminServerGetInfo" }

func (o *xxx_ServerGetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_ServerGetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
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

// ServerGetInfoRequest structure represents the RMprAdminServerGetInfo operation request
type ServerGetInfoRequest struct {
	Level uint32 `idl:"name:dwLevel" json:"level"`
}

func (o *ServerGetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_ServerGetInfoOperation) *xxx_ServerGetInfoOperation {
	if op == nil {
		op = &xxx_ServerGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	return op
}

func (o *ServerGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_ServerGetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
}
func (o *ServerGetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ServerGetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerGetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ServerGetInfoResponse structure represents the RMprAdminServerGetInfo operation response
type ServerGetInfoResponse struct {
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RMprAdminServerGetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ServerGetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_ServerGetInfoOperation) *xxx_ServerGetInfoOperation {
	if op == nil {
		op = &xxx_ServerGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *ServerGetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_ServerGetInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *ServerGetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ServerGetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerGetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConnectionEnumOperation structure represents the RRasAdminConnectionEnum operation
type xxx_ConnectionEnumOperation struct {
	Level                  uint32                      `idl:"name:dwLevel" json:"level"`
	Info                   *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	PreferredMaximumLength uint32                      `idl:"name:dwPreferedMaximumLength" json:"preferred_maximum_length"`
	EntriesRead            uint32                      `idl:"name:lpdwEntriesRead" json:"entries_read"`
	TotalEntries           uint32                      `idl:"name:lpdwTotalEntries" json:"total_entries"`
	Resume                 uint32                      `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	Return                 uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_ConnectionEnumOperation) OpNum() int { return 1 }

func (o *xxx_ConnectionEnumOperation) OpName() string { return "/dimsvc/v0/RRasAdminConnectionEnum" }

func (o *xxx_ConnectionEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwPreferedMaximumLength {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_default_null([cond]) attribute.
		_ptr_lpdwResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwPreferedMaximumLength {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpdwResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_lpdwResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_lpdwResumeHandle, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpdwEntriesRead {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.EntriesRead); err != nil {
			return err
		}
	}
	// lpdwTotalEntries {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TotalEntries); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_default_null([cond]) attribute.
		_ptr_lpdwResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_lpdwResumeHandle); err != nil {
			return err
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

func (o *xxx_ConnectionEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpdwEntriesRead {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.EntriesRead); err != nil {
			return err
		}
	}
	// lpdwTotalEntries {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TotalEntries); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpdwResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_lpdwResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_lpdwResumeHandle, _ptr_lpdwResumeHandle); err != nil {
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

// ConnectionEnumRequest structure represents the RRasAdminConnectionEnum operation request
type ConnectionEnumRequest struct {
	Level                  uint32                      `idl:"name:dwLevel" json:"level"`
	Info                   *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	PreferredMaximumLength uint32                      `idl:"name:dwPreferedMaximumLength" json:"preferred_maximum_length"`
	Resume                 uint32                      `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
}

func (o *ConnectionEnumRequest) xxx_ToOp(ctx context.Context, op *xxx_ConnectionEnumOperation) *xxx_ConnectionEnumOperation {
	if op == nil {
		op = &xxx_ConnectionEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	op.PreferredMaximumLength = o.PreferredMaximumLength
	op.Resume = o.Resume
	return op
}

func (o *ConnectionEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_ConnectionEnumOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.PreferredMaximumLength = op.PreferredMaximumLength
	o.Resume = op.Resume
}
func (o *ConnectionEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ConnectionEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConnectionEnumResponse structure represents the RRasAdminConnectionEnum operation response
type ConnectionEnumResponse struct {
	Info         *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	EntriesRead  uint32                      `idl:"name:lpdwEntriesRead" json:"entries_read"`
	TotalEntries uint32                      `idl:"name:lpdwTotalEntries" json:"total_entries"`
	Resume       uint32                      `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	// Return: The RRasAdminConnectionEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ConnectionEnumResponse) xxx_ToOp(ctx context.Context, op *xxx_ConnectionEnumOperation) *xxx_ConnectionEnumOperation {
	if op == nil {
		op = &xxx_ConnectionEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.EntriesRead = o.EntriesRead
	op.TotalEntries = o.TotalEntries
	op.Resume = o.Resume
	op.Return = o.Return
	return op
}

func (o *ConnectionEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_ConnectionEnumOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.EntriesRead = op.EntriesRead
	o.TotalEntries = op.TotalEntries
	o.Resume = op.Resume
	o.Return = op.Return
}
func (o *ConnectionEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ConnectionEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConnectionGetInfoOperation structure represents the RRasAdminConnectionGetInfo operation
type xxx_ConnectionGetInfoOperation struct {
	Level      uint32                      `idl:"name:dwLevel" json:"level"`
	Connection uint32                      `idl:"name:hDimConnection" json:"connection"`
	Info       *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Return     uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_ConnectionGetInfoOperation) OpNum() int { return 2 }

func (o *xxx_ConnectionGetInfoOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminConnectionGetInfo"
}

func (o *xxx_ConnectionGetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionGetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Connection); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionGetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Connection); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionGetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionGetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_ConnectionGetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
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

// ConnectionGetInfoRequest structure represents the RRasAdminConnectionGetInfo operation request
type ConnectionGetInfoRequest struct {
	Level      uint32 `idl:"name:dwLevel" json:"level"`
	Connection uint32 `idl:"name:hDimConnection" json:"connection"`
}

func (o *ConnectionGetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_ConnectionGetInfoOperation) *xxx_ConnectionGetInfoOperation {
	if op == nil {
		op = &xxx_ConnectionGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Connection = o.Connection
	return op
}

func (o *ConnectionGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_ConnectionGetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Connection = op.Connection
}
func (o *ConnectionGetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ConnectionGetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionGetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConnectionGetInfoResponse structure represents the RRasAdminConnectionGetInfo operation response
type ConnectionGetInfoResponse struct {
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RRasAdminConnectionGetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ConnectionGetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_ConnectionGetInfoOperation) *xxx_ConnectionGetInfoOperation {
	if op == nil {
		op = &xxx_ConnectionGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *ConnectionGetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_ConnectionGetInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *ConnectionGetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ConnectionGetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionGetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConnectionClearStatsOperation structure represents the RRasAdminConnectionClearStats operation
type xxx_ConnectionClearStatsOperation struct {
	Connection uint32 `idl:"name:hDimConnection" json:"connection"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_ConnectionClearStatsOperation) OpNum() int { return 3 }

func (o *xxx_ConnectionClearStatsOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminConnectionClearStats"
}

func (o *xxx_ConnectionClearStatsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionClearStatsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Connection); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionClearStatsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Connection); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionClearStatsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionClearStatsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionClearStatsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ConnectionClearStatsRequest structure represents the RRasAdminConnectionClearStats operation request
type ConnectionClearStatsRequest struct {
	Connection uint32 `idl:"name:hDimConnection" json:"connection"`
}

func (o *ConnectionClearStatsRequest) xxx_ToOp(ctx context.Context, op *xxx_ConnectionClearStatsOperation) *xxx_ConnectionClearStatsOperation {
	if op == nil {
		op = &xxx_ConnectionClearStatsOperation{}
	}
	if o == nil {
		return op
	}
	op.Connection = o.Connection
	return op
}

func (o *ConnectionClearStatsRequest) xxx_FromOp(ctx context.Context, op *xxx_ConnectionClearStatsOperation) {
	if o == nil {
		return
	}
	o.Connection = op.Connection
}
func (o *ConnectionClearStatsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ConnectionClearStatsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionClearStatsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConnectionClearStatsResponse structure represents the RRasAdminConnectionClearStats operation response
type ConnectionClearStatsResponse struct {
	// Return: The RRasAdminConnectionClearStats return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ConnectionClearStatsResponse) xxx_ToOp(ctx context.Context, op *xxx_ConnectionClearStatsOperation) *xxx_ConnectionClearStatsOperation {
	if op == nil {
		op = &xxx_ConnectionClearStatsOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *ConnectionClearStatsResponse) xxx_FromOp(ctx context.Context, op *xxx_ConnectionClearStatsOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ConnectionClearStatsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ConnectionClearStatsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionClearStatsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PortEnumOperation structure represents the RRasAdminPortEnum operation
type xxx_PortEnumOperation struct {
	Level                  uint32                      `idl:"name:dwLevel" json:"level"`
	Connection             uint32                      `idl:"name:hRasConnection" json:"connection"`
	Info                   *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	PreferredMaximumLength uint32                      `idl:"name:dwPreferedMaximumLength" json:"preferred_maximum_length"`
	EntriesRead            uint32                      `idl:"name:lpdwEntriesRead" json:"entries_read"`
	TotalEntries           uint32                      `idl:"name:lpdwTotalEntries" json:"total_entries"`
	Resume                 uint32                      `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	Return                 uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_PortEnumOperation) OpNum() int { return 4 }

func (o *xxx_PortEnumOperation) OpName() string { return "/dimsvc/v0/RRasAdminPortEnum" }

func (o *xxx_PortEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// hRasConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Connection); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwPreferedMaximumLength {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_default_null([cond]) attribute.
		_ptr_lpdwResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// hRasConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Connection); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwPreferedMaximumLength {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpdwResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_lpdwResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_lpdwResumeHandle, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpdwEntriesRead {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.EntriesRead); err != nil {
			return err
		}
	}
	// lpdwTotalEntries {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TotalEntries); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_default_null([cond]) attribute.
		_ptr_lpdwResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_lpdwResumeHandle); err != nil {
			return err
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

func (o *xxx_PortEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpdwEntriesRead {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.EntriesRead); err != nil {
			return err
		}
	}
	// lpdwTotalEntries {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TotalEntries); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpdwResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_lpdwResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_lpdwResumeHandle, _ptr_lpdwResumeHandle); err != nil {
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

// PortEnumRequest structure represents the RRasAdminPortEnum operation request
type PortEnumRequest struct {
	Level                  uint32                      `idl:"name:dwLevel" json:"level"`
	Connection             uint32                      `idl:"name:hRasConnection" json:"connection"`
	Info                   *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	PreferredMaximumLength uint32                      `idl:"name:dwPreferedMaximumLength" json:"preferred_maximum_length"`
	Resume                 uint32                      `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
}

func (o *PortEnumRequest) xxx_ToOp(ctx context.Context, op *xxx_PortEnumOperation) *xxx_PortEnumOperation {
	if op == nil {
		op = &xxx_PortEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Connection = o.Connection
	op.Info = o.Info
	op.PreferredMaximumLength = o.PreferredMaximumLength
	op.Resume = o.Resume
	return op
}

func (o *PortEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_PortEnumOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Connection = op.Connection
	o.Info = op.Info
	o.PreferredMaximumLength = op.PreferredMaximumLength
	o.Resume = op.Resume
}
func (o *PortEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PortEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PortEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PortEnumResponse structure represents the RRasAdminPortEnum operation response
type PortEnumResponse struct {
	Info         *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	EntriesRead  uint32                      `idl:"name:lpdwEntriesRead" json:"entries_read"`
	TotalEntries uint32                      `idl:"name:lpdwTotalEntries" json:"total_entries"`
	Resume       uint32                      `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	// Return: The RRasAdminPortEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *PortEnumResponse) xxx_ToOp(ctx context.Context, op *xxx_PortEnumOperation) *xxx_PortEnumOperation {
	if op == nil {
		op = &xxx_PortEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.EntriesRead = o.EntriesRead
	op.TotalEntries = o.TotalEntries
	op.Resume = o.Resume
	op.Return = o.Return
	return op
}

func (o *PortEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_PortEnumOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.EntriesRead = op.EntriesRead
	o.TotalEntries = op.TotalEntries
	o.Resume = op.Resume
	o.Return = op.Return
}
func (o *PortEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PortEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PortEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PortGetInfoOperation structure represents the RRasAdminPortGetInfo operation
type xxx_PortGetInfoOperation struct {
	Level  uint32                      `idl:"name:dwLevel" json:"level"`
	Port   uint32                      `idl:"name:hPort" json:"port"`
	Info   *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Return uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_PortGetInfoOperation) OpNum() int { return 5 }

func (o *xxx_PortGetInfoOperation) OpName() string { return "/dimsvc/v0/RRasAdminPortGetInfo" }

func (o *xxx_PortGetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortGetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Port); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortGetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Port); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortGetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortGetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_PortGetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
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

// PortGetInfoRequest structure represents the RRasAdminPortGetInfo operation request
type PortGetInfoRequest struct {
	Level uint32 `idl:"name:dwLevel" json:"level"`
	Port  uint32 `idl:"name:hPort" json:"port"`
}

func (o *PortGetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_PortGetInfoOperation) *xxx_PortGetInfoOperation {
	if op == nil {
		op = &xxx_PortGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Port = o.Port
	return op
}

func (o *PortGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_PortGetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Port = op.Port
}
func (o *PortGetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PortGetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PortGetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PortGetInfoResponse structure represents the RRasAdminPortGetInfo operation response
type PortGetInfoResponse struct {
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RRasAdminPortGetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *PortGetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_PortGetInfoOperation) *xxx_PortGetInfoOperation {
	if op == nil {
		op = &xxx_PortGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *PortGetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_PortGetInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *PortGetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PortGetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PortGetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PortClearStatsOperation structure represents the RRasAdminPortClearStats operation
type xxx_PortClearStatsOperation struct {
	Port   uint32 `idl:"name:hPort" json:"port"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_PortClearStatsOperation) OpNum() int { return 6 }

func (o *xxx_PortClearStatsOperation) OpName() string { return "/dimsvc/v0/RRasAdminPortClearStats" }

func (o *xxx_PortClearStatsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortClearStatsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Port); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortClearStatsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Port); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortClearStatsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortClearStatsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortClearStatsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PortClearStatsRequest structure represents the RRasAdminPortClearStats operation request
type PortClearStatsRequest struct {
	Port uint32 `idl:"name:hPort" json:"port"`
}

func (o *PortClearStatsRequest) xxx_ToOp(ctx context.Context, op *xxx_PortClearStatsOperation) *xxx_PortClearStatsOperation {
	if op == nil {
		op = &xxx_PortClearStatsOperation{}
	}
	if o == nil {
		return op
	}
	op.Port = o.Port
	return op
}

func (o *PortClearStatsRequest) xxx_FromOp(ctx context.Context, op *xxx_PortClearStatsOperation) {
	if o == nil {
		return
	}
	o.Port = op.Port
}
func (o *PortClearStatsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PortClearStatsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PortClearStatsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PortClearStatsResponse structure represents the RRasAdminPortClearStats operation response
type PortClearStatsResponse struct {
	// Return: The RRasAdminPortClearStats return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *PortClearStatsResponse) xxx_ToOp(ctx context.Context, op *xxx_PortClearStatsOperation) *xxx_PortClearStatsOperation {
	if op == nil {
		op = &xxx_PortClearStatsOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *PortClearStatsResponse) xxx_FromOp(ctx context.Context, op *xxx_PortClearStatsOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *PortClearStatsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PortClearStatsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PortClearStatsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PortResetOperation structure represents the RRasAdminPortReset operation
type xxx_PortResetOperation struct {
	Port   uint32 `idl:"name:hPort" json:"port"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_PortResetOperation) OpNum() int { return 7 }

func (o *xxx_PortResetOperation) OpName() string { return "/dimsvc/v0/RRasAdminPortReset" }

func (o *xxx_PortResetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortResetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Port); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortResetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Port); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortResetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortResetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortResetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PortResetRequest structure represents the RRasAdminPortReset operation request
type PortResetRequest struct {
	Port uint32 `idl:"name:hPort" json:"port"`
}

func (o *PortResetRequest) xxx_ToOp(ctx context.Context, op *xxx_PortResetOperation) *xxx_PortResetOperation {
	if op == nil {
		op = &xxx_PortResetOperation{}
	}
	if o == nil {
		return op
	}
	op.Port = o.Port
	return op
}

func (o *PortResetRequest) xxx_FromOp(ctx context.Context, op *xxx_PortResetOperation) {
	if o == nil {
		return
	}
	o.Port = op.Port
}
func (o *PortResetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PortResetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PortResetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PortResetResponse structure represents the RRasAdminPortReset operation response
type PortResetResponse struct {
	// Return: The RRasAdminPortReset return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *PortResetResponse) xxx_ToOp(ctx context.Context, op *xxx_PortResetOperation) *xxx_PortResetOperation {
	if op == nil {
		op = &xxx_PortResetOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *PortResetResponse) xxx_FromOp(ctx context.Context, op *xxx_PortResetOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *PortResetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PortResetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PortResetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PortDisconnectOperation structure represents the RRasAdminPortDisconnect operation
type xxx_PortDisconnectOperation struct {
	Port   uint32 `idl:"name:hPort" json:"port"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_PortDisconnectOperation) OpNum() int { return 8 }

func (o *xxx_PortDisconnectOperation) OpName() string { return "/dimsvc/v0/RRasAdminPortDisconnect" }

func (o *xxx_PortDisconnectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortDisconnectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Port); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortDisconnectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Port); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortDisconnectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortDisconnectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PortDisconnectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PortDisconnectRequest structure represents the RRasAdminPortDisconnect operation request
type PortDisconnectRequest struct {
	Port uint32 `idl:"name:hPort" json:"port"`
}

func (o *PortDisconnectRequest) xxx_ToOp(ctx context.Context, op *xxx_PortDisconnectOperation) *xxx_PortDisconnectOperation {
	if op == nil {
		op = &xxx_PortDisconnectOperation{}
	}
	if o == nil {
		return op
	}
	op.Port = o.Port
	return op
}

func (o *PortDisconnectRequest) xxx_FromOp(ctx context.Context, op *xxx_PortDisconnectOperation) {
	if o == nil {
		return
	}
	o.Port = op.Port
}
func (o *PortDisconnectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *PortDisconnectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PortDisconnectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PortDisconnectResponse structure represents the RRasAdminPortDisconnect operation response
type PortDisconnectResponse struct {
	// Return: The RRasAdminPortDisconnect return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *PortDisconnectResponse) xxx_ToOp(ctx context.Context, op *xxx_PortDisconnectOperation) *xxx_PortDisconnectOperation {
	if op == nil {
		op = &xxx_PortDisconnectOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *PortDisconnectResponse) xxx_FromOp(ctx context.Context, op *xxx_PortDisconnectOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *PortDisconnectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *PortDisconnectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PortDisconnectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceTransportSetGlobalInfoOperation structure represents the RRouterInterfaceTransportSetGlobalInfo operation
type xxx_RouterInterfaceTransportSetGlobalInfoOperation struct {
	TransportID uint32                    `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	Return      uint32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceTransportSetGlobalInfoOperation) OpNum() int { return 9 }

func (o *xxx_RouterInterfaceTransportSetGlobalInfoOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceTransportSetGlobalInfo"
}

func (o *xxx_RouterInterfaceTransportSetGlobalInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportSetGlobalInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TransportID); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INTERFACE_CONTAINER}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportSetGlobalInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TransportID); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INTERFACE_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InterfaceContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportSetGlobalInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportSetGlobalInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportSetGlobalInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceTransportSetGlobalInfoRequest structure represents the RRouterInterfaceTransportSetGlobalInfo operation request
type RouterInterfaceTransportSetGlobalInfoRequest struct {
	TransportID uint32                    `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *RouterInterfaceTransportSetGlobalInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportSetGlobalInfoOperation) *xxx_RouterInterfaceTransportSetGlobalInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportSetGlobalInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.TransportID = o.TransportID
	op.Info = o.Info
	return op
}

func (o *RouterInterfaceTransportSetGlobalInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportSetGlobalInfoOperation) {
	if o == nil {
		return
	}
	o.TransportID = op.TransportID
	o.Info = op.Info
}
func (o *RouterInterfaceTransportSetGlobalInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceTransportSetGlobalInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportSetGlobalInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceTransportSetGlobalInfoResponse structure represents the RRouterInterfaceTransportSetGlobalInfo operation response
type RouterInterfaceTransportSetGlobalInfoResponse struct {
	// Return: The RRouterInterfaceTransportSetGlobalInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceTransportSetGlobalInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportSetGlobalInfoOperation) *xxx_RouterInterfaceTransportSetGlobalInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportSetGlobalInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceTransportSetGlobalInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportSetGlobalInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceTransportSetGlobalInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceTransportSetGlobalInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportSetGlobalInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceTransportGetGlobalInfoOperation structure represents the RRouterInterfaceTransportGetGlobalInfo operation
type xxx_RouterInterfaceTransportGetGlobalInfoOperation struct {
	TransportID uint32                    `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	Return      uint32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceTransportGetGlobalInfoOperation) OpNum() int { return 10 }

func (o *xxx_RouterInterfaceTransportGetGlobalInfoOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceTransportGetGlobalInfo"
}

func (o *xxx_RouterInterfaceTransportGetGlobalInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportGetGlobalInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TransportID); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INTERFACE_CONTAINER}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportGetGlobalInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TransportID); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INTERFACE_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InterfaceContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportGetGlobalInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportGetGlobalInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INTERFACE_CONTAINER}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RouterInterfaceTransportGetGlobalInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_INTERFACE_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InterfaceContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
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

// RouterInterfaceTransportGetGlobalInfoRequest structure represents the RRouterInterfaceTransportGetGlobalInfo operation request
type RouterInterfaceTransportGetGlobalInfoRequest struct {
	TransportID uint32                    `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *RouterInterfaceTransportGetGlobalInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportGetGlobalInfoOperation) *xxx_RouterInterfaceTransportGetGlobalInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportGetGlobalInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.TransportID = o.TransportID
	op.Info = o.Info
	return op
}

func (o *RouterInterfaceTransportGetGlobalInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportGetGlobalInfoOperation) {
	if o == nil {
		return
	}
	o.TransportID = op.TransportID
	o.Info = op.Info
}
func (o *RouterInterfaceTransportGetGlobalInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceTransportGetGlobalInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportGetGlobalInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceTransportGetGlobalInfoResponse structure represents the RRouterInterfaceTransportGetGlobalInfo operation response
type RouterInterfaceTransportGetGlobalInfoResponse struct {
	Info *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RRouterInterfaceTransportGetGlobalInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceTransportGetGlobalInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportGetGlobalInfoOperation) *xxx_RouterInterfaceTransportGetGlobalInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportGetGlobalInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceTransportGetGlobalInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportGetGlobalInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *RouterInterfaceTransportGetGlobalInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceTransportGetGlobalInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportGetGlobalInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceGetHandleOperation structure represents the RRouterInterfaceGetHandle operation
type xxx_RouterInterfaceGetHandleOperation struct {
	InterfaceName           string `idl:"name:lpwsInterfaceName;string" json:"interface_name"`
	Interface               uint32 `idl:"name:phInterface" json:"interface"`
	IncludeClientInterfaces uint32 `idl:"name:fIncludeClientInterfaces" json:"include_client_interfaces"`
	Return                  uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceGetHandleOperation) OpNum() int { return 11 }

func (o *xxx_RouterInterfaceGetHandleOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceGetHandle"
}

func (o *xxx_RouterInterfaceGetHandleOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetHandleOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpwsInterfaceName {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.InterfaceName); err != nil {
			return err
		}
	}
	// phInterface {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// fIncludeClientInterfaces {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.IncludeClientInterfaces); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetHandleOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpwsInterfaceName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.InterfaceName); err != nil {
			return err
		}
	}
	// phInterface {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// fIncludeClientInterfaces {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.IncludeClientInterfaces); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetHandleOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetHandleOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phInterface {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
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

func (o *xxx_RouterInterfaceGetHandleOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phInterface {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
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

// RouterInterfaceGetHandleRequest structure represents the RRouterInterfaceGetHandle operation request
type RouterInterfaceGetHandleRequest struct {
	InterfaceName           string `idl:"name:lpwsInterfaceName;string" json:"interface_name"`
	Interface               uint32 `idl:"name:phInterface" json:"interface"`
	IncludeClientInterfaces uint32 `idl:"name:fIncludeClientInterfaces" json:"include_client_interfaces"`
}

func (o *RouterInterfaceGetHandleRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetHandleOperation) *xxx_RouterInterfaceGetHandleOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetHandleOperation{}
	}
	if o == nil {
		return op
	}
	op.InterfaceName = o.InterfaceName
	op.Interface = o.Interface
	op.IncludeClientInterfaces = o.IncludeClientInterfaces
	return op
}

func (o *RouterInterfaceGetHandleRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetHandleOperation) {
	if o == nil {
		return
	}
	o.InterfaceName = op.InterfaceName
	o.Interface = op.Interface
	o.IncludeClientInterfaces = op.IncludeClientInterfaces
}
func (o *RouterInterfaceGetHandleRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceGetHandleRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceGetHandleOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceGetHandleResponse structure represents the RRouterInterfaceGetHandle operation response
type RouterInterfaceGetHandleResponse struct {
	Interface uint32 `idl:"name:phInterface" json:"interface"`
	// Return: The RRouterInterfaceGetHandle return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceGetHandleResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetHandleOperation) *xxx_RouterInterfaceGetHandleOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetHandleOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceGetHandleResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetHandleOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.Return = op.Return
}
func (o *RouterInterfaceGetHandleResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceGetHandleResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceGetHandleOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceCreateOperation structure represents the RRouterInterfaceCreate operation
type xxx_RouterInterfaceCreateOperation struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Interface uint32                      `idl:"name:phInterface" json:"interface"`
	Return    uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceCreateOperation) OpNum() int { return 12 }

func (o *xxx_RouterInterfaceCreateOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceCreate"
}

func (o *xxx_RouterInterfaceCreateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceCreateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// phInterface {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceCreateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// phInterface {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceCreateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceCreateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phInterface {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
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

func (o *xxx_RouterInterfaceCreateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phInterface {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
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

// RouterInterfaceCreateRequest structure represents the RRouterInterfaceCreate operation request
type RouterInterfaceCreateRequest struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Interface uint32                      `idl:"name:phInterface" json:"interface"`
}

func (o *RouterInterfaceCreateRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceCreateOperation) *xxx_RouterInterfaceCreateOperation {
	if op == nil {
		op = &xxx_RouterInterfaceCreateOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	op.Interface = o.Interface
	return op
}

func (o *RouterInterfaceCreateRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceCreateOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.Interface = op.Interface
}
func (o *RouterInterfaceCreateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceCreateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceCreateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceCreateResponse structure represents the RRouterInterfaceCreate operation response
type RouterInterfaceCreateResponse struct {
	Interface uint32 `idl:"name:phInterface" json:"interface"`
	// Return: The RRouterInterfaceCreate return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceCreateResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceCreateOperation) *xxx_RouterInterfaceCreateOperation {
	if op == nil {
		op = &xxx_RouterInterfaceCreateOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceCreateResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceCreateOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.Return = op.Return
}
func (o *RouterInterfaceCreateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceCreateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceCreateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceGetInfoOperation structure represents the RRouterInterfaceGetInfo operation
type xxx_RouterInterfaceGetInfoOperation struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Interface uint32                      `idl:"name:hInterface" json:"interface"`
	Return    uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceGetInfoOperation) OpNum() int { return 13 }

func (o *xxx_RouterInterfaceGetInfoOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceGetInfo"
}

func (o *xxx_RouterInterfaceGetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RouterInterfaceGetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
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

// RouterInterfaceGetInfoRequest structure represents the RRouterInterfaceGetInfo operation request
type RouterInterfaceGetInfoRequest struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Interface uint32                      `idl:"name:hInterface" json:"interface"`
}

func (o *RouterInterfaceGetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetInfoOperation) *xxx_RouterInterfaceGetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	op.Interface = o.Interface
	return op
}

func (o *RouterInterfaceGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.Interface = op.Interface
}
func (o *RouterInterfaceGetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceGetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceGetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceGetInfoResponse structure represents the RRouterInterfaceGetInfo operation response
type RouterInterfaceGetInfoResponse struct {
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RRouterInterfaceGetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceGetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetInfoOperation) *xxx_RouterInterfaceGetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceGetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *RouterInterfaceGetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceGetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceGetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceSetInfoOperation structure represents the RRouterInterfaceSetInfo operation
type xxx_RouterInterfaceSetInfoOperation struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Interface uint32                      `idl:"name:hInterface" json:"interface"`
	Return    uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceSetInfoOperation) OpNum() int { return 14 }

func (o *xxx_RouterInterfaceSetInfoOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceSetInfo"
}

func (o *xxx_RouterInterfaceSetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceSetInfoRequest structure represents the RRouterInterfaceSetInfo operation request
type RouterInterfaceSetInfoRequest struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Interface uint32                      `idl:"name:hInterface" json:"interface"`
}

func (o *RouterInterfaceSetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceSetInfoOperation) *xxx_RouterInterfaceSetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceSetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	op.Interface = o.Interface
	return op
}

func (o *RouterInterfaceSetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceSetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.Interface = op.Interface
}
func (o *RouterInterfaceSetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceSetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceSetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceSetInfoResponse structure represents the RRouterInterfaceSetInfo operation response
type RouterInterfaceSetInfoResponse struct {
	// Return: The RRouterInterfaceSetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceSetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceSetInfoOperation) *xxx_RouterInterfaceSetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceSetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceSetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceSetInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceSetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceSetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceSetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceDeleteOperation structure represents the RRouterInterfaceDelete operation
type xxx_RouterInterfaceDeleteOperation struct {
	Interface uint32 `idl:"name:hInterface" json:"interface"`
	Return    uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceDeleteOperation) OpNum() int { return 15 }

func (o *xxx_RouterInterfaceDeleteOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceDelete"
}

func (o *xxx_RouterInterfaceDeleteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeleteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeleteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeleteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceDeleteRequest structure represents the RRouterInterfaceDelete operation request
type RouterInterfaceDeleteRequest struct {
	Interface uint32 `idl:"name:hInterface" json:"interface"`
}

func (o *RouterInterfaceDeleteRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceDeleteOperation) *xxx_RouterInterfaceDeleteOperation {
	if op == nil {
		op = &xxx_RouterInterfaceDeleteOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	return op
}

func (o *RouterInterfaceDeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceDeleteOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
}
func (o *RouterInterfaceDeleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceDeleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceDeleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceDeleteResponse structure represents the RRouterInterfaceDelete operation response
type RouterInterfaceDeleteResponse struct {
	// Return: The RRouterInterfaceDelete return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceDeleteResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceDeleteOperation) *xxx_RouterInterfaceDeleteOperation {
	if op == nil {
		op = &xxx_RouterInterfaceDeleteOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceDeleteResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceDeleteOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceDeleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceDeleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceDeleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceTransportRemoveOperation structure represents the RRouterInterfaceTransportRemove operation
type xxx_RouterInterfaceTransportRemoveOperation struct {
	Interface   uint32 `idl:"name:hInterface" json:"interface"`
	TransportID uint32 `idl:"name:dwTransportId" json:"transport_id"`
	Return      uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceTransportRemoveOperation) OpNum() int { return 16 }

func (o *xxx_RouterInterfaceTransportRemoveOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceTransportRemove"
}

func (o *xxx_RouterInterfaceTransportRemoveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportRemoveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TransportID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportRemoveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TransportID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportRemoveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportRemoveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportRemoveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceTransportRemoveRequest structure represents the RRouterInterfaceTransportRemove operation request
type RouterInterfaceTransportRemoveRequest struct {
	Interface   uint32 `idl:"name:hInterface" json:"interface"`
	TransportID uint32 `idl:"name:dwTransportId" json:"transport_id"`
}

func (o *RouterInterfaceTransportRemoveRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportRemoveOperation) *xxx_RouterInterfaceTransportRemoveOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportRemoveOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.TransportID = o.TransportID
	return op
}

func (o *RouterInterfaceTransportRemoveRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportRemoveOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.TransportID = op.TransportID
}
func (o *RouterInterfaceTransportRemoveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceTransportRemoveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportRemoveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceTransportRemoveResponse structure represents the RRouterInterfaceTransportRemove operation response
type RouterInterfaceTransportRemoveResponse struct {
	// Return: The RRouterInterfaceTransportRemove return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceTransportRemoveResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportRemoveOperation) *xxx_RouterInterfaceTransportRemoveOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportRemoveOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceTransportRemoveResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportRemoveOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceTransportRemoveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceTransportRemoveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportRemoveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceTransportAddOperation structure represents the RRouterInterfaceTransportAdd operation
type xxx_RouterInterfaceTransportAddOperation struct {
	Interface   uint32                    `idl:"name:hInterface" json:"interface"`
	TransportID uint32                    `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	Return      uint32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceTransportAddOperation) OpNum() int { return 17 }

func (o *xxx_RouterInterfaceTransportAddOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceTransportAdd"
}

func (o *xxx_RouterInterfaceTransportAddOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportAddOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TransportID); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INTERFACE_CONTAINER}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportAddOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TransportID); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INTERFACE_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InterfaceContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportAddOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportAddOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportAddOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceTransportAddRequest structure represents the RRouterInterfaceTransportAdd operation request
type RouterInterfaceTransportAddRequest struct {
	Interface   uint32                    `idl:"name:hInterface" json:"interface"`
	TransportID uint32                    `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *RouterInterfaceTransportAddRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportAddOperation) *xxx_RouterInterfaceTransportAddOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportAddOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.TransportID = o.TransportID
	op.Info = o.Info
	return op
}

func (o *RouterInterfaceTransportAddRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportAddOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.TransportID = op.TransportID
	o.Info = op.Info
}
func (o *RouterInterfaceTransportAddRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceTransportAddRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportAddOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceTransportAddResponse structure represents the RRouterInterfaceTransportAdd operation response
type RouterInterfaceTransportAddResponse struct {
	// Return: The RRouterInterfaceTransportAdd return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceTransportAddResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportAddOperation) *xxx_RouterInterfaceTransportAddOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportAddOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceTransportAddResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportAddOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceTransportAddResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceTransportAddResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportAddOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceTransportGetInfoOperation structure represents the RRouterInterfaceTransportGetInfo operation
type xxx_RouterInterfaceTransportGetInfoOperation struct {
	Interface   uint32                    `idl:"name:hInterface" json:"interface"`
	TransportID uint32                    `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	Return      uint32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceTransportGetInfoOperation) OpNum() int { return 18 }

func (o *xxx_RouterInterfaceTransportGetInfoOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceTransportGetInfo"
}

func (o *xxx_RouterInterfaceTransportGetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportGetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TransportID); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INTERFACE_CONTAINER}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportGetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TransportID); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INTERFACE_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InterfaceContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportGetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportGetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INTERFACE_CONTAINER}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RouterInterfaceTransportGetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_INTERFACE_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InterfaceContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
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

// RouterInterfaceTransportGetInfoRequest structure represents the RRouterInterfaceTransportGetInfo operation request
type RouterInterfaceTransportGetInfoRequest struct {
	Interface   uint32                    `idl:"name:hInterface" json:"interface"`
	TransportID uint32                    `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *RouterInterfaceTransportGetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportGetInfoOperation) *xxx_RouterInterfaceTransportGetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.TransportID = o.TransportID
	op.Info = o.Info
	return op
}

func (o *RouterInterfaceTransportGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportGetInfoOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.TransportID = op.TransportID
	o.Info = op.Info
}
func (o *RouterInterfaceTransportGetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceTransportGetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportGetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceTransportGetInfoResponse structure represents the RRouterInterfaceTransportGetInfo operation response
type RouterInterfaceTransportGetInfoResponse struct {
	Info *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RRouterInterfaceTransportGetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceTransportGetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportGetInfoOperation) *xxx_RouterInterfaceTransportGetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceTransportGetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportGetInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *RouterInterfaceTransportGetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceTransportGetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportGetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceTransportSetInfoOperation structure represents the RRouterInterfaceTransportSetInfo operation
type xxx_RouterInterfaceTransportSetInfoOperation struct {
	Interface   uint32                    `idl:"name:hInterface" json:"interface"`
	TransportID uint32                    `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	Return      uint32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceTransportSetInfoOperation) OpNum() int { return 19 }

func (o *xxx_RouterInterfaceTransportSetInfoOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceTransportSetInfo"
}

func (o *xxx_RouterInterfaceTransportSetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportSetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TransportID); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INTERFACE_CONTAINER}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportSetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TransportID); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INTERFACE_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InterfaceContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportSetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportSetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportSetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceTransportSetInfoRequest structure represents the RRouterInterfaceTransportSetInfo operation request
type RouterInterfaceTransportSetInfoRequest struct {
	Interface   uint32                    `idl:"name:hInterface" json:"interface"`
	TransportID uint32                    `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *RouterInterfaceTransportSetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportSetInfoOperation) *xxx_RouterInterfaceTransportSetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportSetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.TransportID = o.TransportID
	op.Info = o.Info
	return op
}

func (o *RouterInterfaceTransportSetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportSetInfoOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.TransportID = op.TransportID
	o.Info = op.Info
}
func (o *RouterInterfaceTransportSetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceTransportSetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportSetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceTransportSetInfoResponse structure represents the RRouterInterfaceTransportSetInfo operation response
type RouterInterfaceTransportSetInfoResponse struct {
	// Return: The RRouterInterfaceTransportSetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceTransportSetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportSetInfoOperation) *xxx_RouterInterfaceTransportSetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportSetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceTransportSetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportSetInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceTransportSetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceTransportSetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportSetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceEnumOperation structure represents the RRouterInterfaceEnum operation
type xxx_RouterInterfaceEnumOperation struct {
	Level                  uint32                      `idl:"name:dwLevel" json:"level"`
	Info                   *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	PreferredMaximumLength uint32                      `idl:"name:dwPreferedMaximumLength" json:"preferred_maximum_length"`
	EntriesRead            uint32                      `idl:"name:lpdwEntriesRead" json:"entries_read"`
	TotalEntries           uint32                      `idl:"name:lpdwTotalEntries" json:"total_entries"`
	Resume                 uint32                      `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	Return                 uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceEnumOperation) OpNum() int { return 20 }

func (o *xxx_RouterInterfaceEnumOperation) OpName() string { return "/dimsvc/v0/RRouterInterfaceEnum" }

func (o *xxx_RouterInterfaceEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwPreferedMaximumLength {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_default_null([cond]) attribute.
		_ptr_lpdwResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwPreferedMaximumLength {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PreferredMaximumLength); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpdwResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_lpdwResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_lpdwResumeHandle, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpdwEntriesRead {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.EntriesRead); err != nil {
			return err
		}
	}
	// lpdwTotalEntries {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TotalEntries); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_default_null([cond]) attribute.
		_ptr_lpdwResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_lpdwResumeHandle); err != nil {
			return err
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

func (o *xxx_RouterInterfaceEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpdwEntriesRead {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.EntriesRead); err != nil {
			return err
		}
	}
	// lpdwTotalEntries {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TotalEntries); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpdwResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_lpdwResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_lpdwResumeHandle, _ptr_lpdwResumeHandle); err != nil {
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

// RouterInterfaceEnumRequest structure represents the RRouterInterfaceEnum operation request
type RouterInterfaceEnumRequest struct {
	Level                  uint32                      `idl:"name:dwLevel" json:"level"`
	Info                   *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	PreferredMaximumLength uint32                      `idl:"name:dwPreferedMaximumLength" json:"preferred_maximum_length"`
	Resume                 uint32                      `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
}

func (o *RouterInterfaceEnumRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceEnumOperation) *xxx_RouterInterfaceEnumOperation {
	if op == nil {
		op = &xxx_RouterInterfaceEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	op.PreferredMaximumLength = o.PreferredMaximumLength
	op.Resume = o.Resume
	return op
}

func (o *RouterInterfaceEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceEnumOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.PreferredMaximumLength = op.PreferredMaximumLength
	o.Resume = op.Resume
}
func (o *RouterInterfaceEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceEnumResponse structure represents the RRouterInterfaceEnum operation response
type RouterInterfaceEnumResponse struct {
	Info         *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	EntriesRead  uint32                      `idl:"name:lpdwEntriesRead" json:"entries_read"`
	TotalEntries uint32                      `idl:"name:lpdwTotalEntries" json:"total_entries"`
	Resume       uint32                      `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	// Return: The RRouterInterfaceEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceEnumResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceEnumOperation) *xxx_RouterInterfaceEnumOperation {
	if op == nil {
		op = &xxx_RouterInterfaceEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.EntriesRead = o.EntriesRead
	op.TotalEntries = o.TotalEntries
	op.Resume = o.Resume
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceEnumOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.EntriesRead = op.EntriesRead
	o.TotalEntries = op.TotalEntries
	o.Resume = op.Resume
	o.Return = op.Return
}
func (o *RouterInterfaceEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceConnectOperation structure represents the RRouterInterfaceConnect operation
type xxx_RouterInterfaceConnectOperation struct {
	Interface        uint32 `idl:"name:hInterface" json:"interface"`
	Event            uint64 `idl:"name:hEvent" json:"event"`
	Blocking         uint32 `idl:"name:fBlocking" json:"blocking"`
	CallersProcessID uint32 `idl:"name:dwCallersProcessId" json:"callers_process_id"`
	Return           uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceConnectOperation) OpNum() int { return 21 }

func (o *xxx_RouterInterfaceConnectOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceConnect"
}

func (o *xxx_RouterInterfaceConnectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceConnectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// hEvent {in} (1:{public, alias=ULONG_PTR}(uint32_64))
	{
		if err := w.WriteData(ndr.Uint3264(o.Event)); err != nil {
			return err
		}
	}
	// fBlocking {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Blocking); err != nil {
			return err
		}
	}
	// dwCallersProcessId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.CallersProcessID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceConnectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// hEvent {in} (1:{public, alias=ULONG_PTR}(uint32_64))
	{
		if err := w.ReadData((*ndr.Uint3264)(&o.Event)); err != nil {
			return err
		}
	}
	// fBlocking {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Blocking); err != nil {
			return err
		}
	}
	// dwCallersProcessId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.CallersProcessID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceConnectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceConnectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceConnectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceConnectRequest structure represents the RRouterInterfaceConnect operation request
type RouterInterfaceConnectRequest struct {
	Interface        uint32 `idl:"name:hInterface" json:"interface"`
	Event            uint64 `idl:"name:hEvent" json:"event"`
	Blocking         uint32 `idl:"name:fBlocking" json:"blocking"`
	CallersProcessID uint32 `idl:"name:dwCallersProcessId" json:"callers_process_id"`
}

func (o *RouterInterfaceConnectRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceConnectOperation) *xxx_RouterInterfaceConnectOperation {
	if op == nil {
		op = &xxx_RouterInterfaceConnectOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.Event = o.Event
	op.Blocking = o.Blocking
	op.CallersProcessID = o.CallersProcessID
	return op
}

func (o *RouterInterfaceConnectRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceConnectOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.Event = op.Event
	o.Blocking = op.Blocking
	o.CallersProcessID = op.CallersProcessID
}
func (o *RouterInterfaceConnectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceConnectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceConnectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceConnectResponse structure represents the RRouterInterfaceConnect operation response
type RouterInterfaceConnectResponse struct {
	// Return: The RRouterInterfaceConnect return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceConnectResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceConnectOperation) *xxx_RouterInterfaceConnectOperation {
	if op == nil {
		op = &xxx_RouterInterfaceConnectOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceConnectResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceConnectOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceConnectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceConnectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceConnectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceDisconnectOperation structure represents the RRouterInterfaceDisconnect operation
type xxx_RouterInterfaceDisconnectOperation struct {
	Interface uint32 `idl:"name:hInterface" json:"interface"`
	Return    uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceDisconnectOperation) OpNum() int { return 22 }

func (o *xxx_RouterInterfaceDisconnectOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceDisconnect"
}

func (o *xxx_RouterInterfaceDisconnectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDisconnectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDisconnectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDisconnectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDisconnectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDisconnectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceDisconnectRequest structure represents the RRouterInterfaceDisconnect operation request
type RouterInterfaceDisconnectRequest struct {
	Interface uint32 `idl:"name:hInterface" json:"interface"`
}

func (o *RouterInterfaceDisconnectRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceDisconnectOperation) *xxx_RouterInterfaceDisconnectOperation {
	if op == nil {
		op = &xxx_RouterInterfaceDisconnectOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	return op
}

func (o *RouterInterfaceDisconnectRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceDisconnectOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
}
func (o *RouterInterfaceDisconnectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceDisconnectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceDisconnectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceDisconnectResponse structure represents the RRouterInterfaceDisconnect operation response
type RouterInterfaceDisconnectResponse struct {
	// Return: The RRouterInterfaceDisconnect return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceDisconnectResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceDisconnectOperation) *xxx_RouterInterfaceDisconnectOperation {
	if op == nil {
		op = &xxx_RouterInterfaceDisconnectOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceDisconnectResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceDisconnectOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceDisconnectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceDisconnectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceDisconnectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceUpdateRoutesOperation structure represents the RRouterInterfaceUpdateRoutes operation
type xxx_RouterInterfaceUpdateRoutesOperation struct {
	Interface       uint32 `idl:"name:hInterface" json:"interface"`
	TransportID     uint32 `idl:"name:dwTransportId" json:"transport_id"`
	Event           uint64 `idl:"name:hEvent" json:"event"`
	ClientProcessID uint32 `idl:"name:dwClientProcessId" json:"client_process_id"`
	Return          uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceUpdateRoutesOperation) OpNum() int { return 23 }

func (o *xxx_RouterInterfaceUpdateRoutesOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceUpdateRoutes"
}

func (o *xxx_RouterInterfaceUpdateRoutesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceUpdateRoutesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TransportID); err != nil {
			return err
		}
	}
	// hEvent {in} (1:{public, alias=ULONG_PTR}(uint32_64))
	{
		if err := w.WriteData(ndr.Uint3264(o.Event)); err != nil {
			return err
		}
	}
	// dwClientProcessId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientProcessID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceUpdateRoutesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TransportID); err != nil {
			return err
		}
	}
	// hEvent {in} (1:{public, alias=ULONG_PTR}(uint32_64))
	{
		if err := w.ReadData((*ndr.Uint3264)(&o.Event)); err != nil {
			return err
		}
	}
	// dwClientProcessId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientProcessID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceUpdateRoutesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceUpdateRoutesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceUpdateRoutesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceUpdateRoutesRequest structure represents the RRouterInterfaceUpdateRoutes operation request
type RouterInterfaceUpdateRoutesRequest struct {
	Interface       uint32 `idl:"name:hInterface" json:"interface"`
	TransportID     uint32 `idl:"name:dwTransportId" json:"transport_id"`
	Event           uint64 `idl:"name:hEvent" json:"event"`
	ClientProcessID uint32 `idl:"name:dwClientProcessId" json:"client_process_id"`
}

func (o *RouterInterfaceUpdateRoutesRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceUpdateRoutesOperation) *xxx_RouterInterfaceUpdateRoutesOperation {
	if op == nil {
		op = &xxx_RouterInterfaceUpdateRoutesOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.TransportID = o.TransportID
	op.Event = o.Event
	op.ClientProcessID = o.ClientProcessID
	return op
}

func (o *RouterInterfaceUpdateRoutesRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceUpdateRoutesOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.TransportID = op.TransportID
	o.Event = op.Event
	o.ClientProcessID = op.ClientProcessID
}
func (o *RouterInterfaceUpdateRoutesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceUpdateRoutesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceUpdateRoutesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceUpdateRoutesResponse structure represents the RRouterInterfaceUpdateRoutes operation response
type RouterInterfaceUpdateRoutesResponse struct {
	// Return: The RRouterInterfaceUpdateRoutes return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceUpdateRoutesResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceUpdateRoutesOperation) *xxx_RouterInterfaceUpdateRoutesOperation {
	if op == nil {
		op = &xxx_RouterInterfaceUpdateRoutesOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceUpdateRoutesResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceUpdateRoutesOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceUpdateRoutesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceUpdateRoutesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceUpdateRoutesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceQueryUpdateResultOperation structure represents the RRouterInterfaceQueryUpdateResult operation
type xxx_RouterInterfaceQueryUpdateResultOperation struct {
	Interface    uint32 `idl:"name:hInterface" json:"interface"`
	TransportID  uint32 `idl:"name:dwTransportId" json:"transport_id"`
	UpdateResult uint32 `idl:"name:pUpdateResult" json:"update_result"`
	Return       uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceQueryUpdateResultOperation) OpNum() int { return 24 }

func (o *xxx_RouterInterfaceQueryUpdateResultOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceQueryUpdateResult"
}

func (o *xxx_RouterInterfaceQueryUpdateResultOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceQueryUpdateResultOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TransportID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceQueryUpdateResultOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TransportID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceQueryUpdateResultOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceQueryUpdateResultOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pUpdateResult {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.UpdateResult); err != nil {
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

func (o *xxx_RouterInterfaceQueryUpdateResultOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pUpdateResult {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.UpdateResult); err != nil {
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

// RouterInterfaceQueryUpdateResultRequest structure represents the RRouterInterfaceQueryUpdateResult operation request
type RouterInterfaceQueryUpdateResultRequest struct {
	Interface   uint32 `idl:"name:hInterface" json:"interface"`
	TransportID uint32 `idl:"name:dwTransportId" json:"transport_id"`
}

func (o *RouterInterfaceQueryUpdateResultRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceQueryUpdateResultOperation) *xxx_RouterInterfaceQueryUpdateResultOperation {
	if op == nil {
		op = &xxx_RouterInterfaceQueryUpdateResultOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.TransportID = o.TransportID
	return op
}

func (o *RouterInterfaceQueryUpdateResultRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceQueryUpdateResultOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.TransportID = op.TransportID
}
func (o *RouterInterfaceQueryUpdateResultRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceQueryUpdateResultRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceQueryUpdateResultOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceQueryUpdateResultResponse structure represents the RRouterInterfaceQueryUpdateResult operation response
type RouterInterfaceQueryUpdateResultResponse struct {
	UpdateResult uint32 `idl:"name:pUpdateResult" json:"update_result"`
	// Return: The RRouterInterfaceQueryUpdateResult return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceQueryUpdateResultResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceQueryUpdateResultOperation) *xxx_RouterInterfaceQueryUpdateResultOperation {
	if op == nil {
		op = &xxx_RouterInterfaceQueryUpdateResultOperation{}
	}
	if o == nil {
		return op
	}
	op.UpdateResult = o.UpdateResult
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceQueryUpdateResultResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceQueryUpdateResultOperation) {
	if o == nil {
		return
	}
	o.UpdateResult = op.UpdateResult
	o.Return = op.Return
}
func (o *RouterInterfaceQueryUpdateResultResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceQueryUpdateResultResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceQueryUpdateResultOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceUpdatePhonebookInfoOperation structure represents the RRouterInterfaceUpdatePhonebookInfo operation
type xxx_RouterInterfaceUpdatePhonebookInfoOperation struct {
	Interface uint32 `idl:"name:hInterface" json:"interface"`
	Return    uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceUpdatePhonebookInfoOperation) OpNum() int { return 25 }

func (o *xxx_RouterInterfaceUpdatePhonebookInfoOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceUpdatePhonebookInfo"
}

func (o *xxx_RouterInterfaceUpdatePhonebookInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceUpdatePhonebookInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceUpdatePhonebookInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceUpdatePhonebookInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceUpdatePhonebookInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceUpdatePhonebookInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceUpdatePhonebookInfoRequest structure represents the RRouterInterfaceUpdatePhonebookInfo operation request
type RouterInterfaceUpdatePhonebookInfoRequest struct {
	Interface uint32 `idl:"name:hInterface" json:"interface"`
}

func (o *RouterInterfaceUpdatePhonebookInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceUpdatePhonebookInfoOperation) *xxx_RouterInterfaceUpdatePhonebookInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceUpdatePhonebookInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	return op
}

func (o *RouterInterfaceUpdatePhonebookInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceUpdatePhonebookInfoOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
}
func (o *RouterInterfaceUpdatePhonebookInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceUpdatePhonebookInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceUpdatePhonebookInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceUpdatePhonebookInfoResponse structure represents the RRouterInterfaceUpdatePhonebookInfo operation response
type RouterInterfaceUpdatePhonebookInfoResponse struct {
	// Return: The RRouterInterfaceUpdatePhonebookInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceUpdatePhonebookInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceUpdatePhonebookInfoOperation) *xxx_RouterInterfaceUpdatePhonebookInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceUpdatePhonebookInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceUpdatePhonebookInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceUpdatePhonebookInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceUpdatePhonebookInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceUpdatePhonebookInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceUpdatePhonebookInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MIBEntryCreateOperation structure represents the RMIBEntryCreate operation
type xxx_MIBEntryCreateOperation struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
	Return     uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_MIBEntryCreateOperation) OpNum() int { return 26 }

func (o *xxx_MIBEntryCreateOperation) OpName() string { return "/dimsvc/v0/RMIBEntryCreate" }

func (o *xxx_MIBEntryCreateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryCreateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryCreateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryCreateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryCreateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryCreateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MIBEntryCreateRequest structure represents the RMIBEntryCreate operation request
type MIBEntryCreateRequest struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
}

func (o *MIBEntryCreateRequest) xxx_ToOp(ctx context.Context, op *xxx_MIBEntryCreateOperation) *xxx_MIBEntryCreateOperation {
	if op == nil {
		op = &xxx_MIBEntryCreateOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.Info = o.Info
	return op
}

func (o *MIBEntryCreateRequest) xxx_FromOp(ctx context.Context, op *xxx_MIBEntryCreateOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.Info = op.Info
}
func (o *MIBEntryCreateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MIBEntryCreateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntryCreateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MIBEntryCreateResponse structure represents the RMIBEntryCreate operation response
type MIBEntryCreateResponse struct {
	// Return: The RMIBEntryCreate return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MIBEntryCreateResponse) xxx_ToOp(ctx context.Context, op *xxx_MIBEntryCreateOperation) *xxx_MIBEntryCreateOperation {
	if op == nil {
		op = &xxx_MIBEntryCreateOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *MIBEntryCreateResponse) xxx_FromOp(ctx context.Context, op *xxx_MIBEntryCreateOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *MIBEntryCreateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MIBEntryCreateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntryCreateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MIBEntryDeleteOperation structure represents the RMIBEntryDelete operation
type xxx_MIBEntryDeleteOperation struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
	Return     uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_MIBEntryDeleteOperation) OpNum() int { return 27 }

func (o *xxx_MIBEntryDeleteOperation) OpName() string { return "/dimsvc/v0/RMIBEntryDelete" }

func (o *xxx_MIBEntryDeleteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryDeleteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryDeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryDeleteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryDeleteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryDeleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MIBEntryDeleteRequest structure represents the RMIBEntryDelete operation request
type MIBEntryDeleteRequest struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
}

func (o *MIBEntryDeleteRequest) xxx_ToOp(ctx context.Context, op *xxx_MIBEntryDeleteOperation) *xxx_MIBEntryDeleteOperation {
	if op == nil {
		op = &xxx_MIBEntryDeleteOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.Info = o.Info
	return op
}

func (o *MIBEntryDeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_MIBEntryDeleteOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.Info = op.Info
}
func (o *MIBEntryDeleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MIBEntryDeleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntryDeleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MIBEntryDeleteResponse structure represents the RMIBEntryDelete operation response
type MIBEntryDeleteResponse struct {
	// Return: The RMIBEntryDelete return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MIBEntryDeleteResponse) xxx_ToOp(ctx context.Context, op *xxx_MIBEntryDeleteOperation) *xxx_MIBEntryDeleteOperation {
	if op == nil {
		op = &xxx_MIBEntryDeleteOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *MIBEntryDeleteResponse) xxx_FromOp(ctx context.Context, op *xxx_MIBEntryDeleteOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *MIBEntryDeleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MIBEntryDeleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntryDeleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MIBEntrySetOperation structure represents the RMIBEntrySet operation
type xxx_MIBEntrySetOperation struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
	Return     uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_MIBEntrySetOperation) OpNum() int { return 28 }

func (o *xxx_MIBEntrySetOperation) OpName() string { return "/dimsvc/v0/RMIBEntrySet" }

func (o *xxx_MIBEntrySetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntrySetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntrySetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntrySetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntrySetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntrySetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MIBEntrySetRequest structure represents the RMIBEntrySet operation request
type MIBEntrySetRequest struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
}

func (o *MIBEntrySetRequest) xxx_ToOp(ctx context.Context, op *xxx_MIBEntrySetOperation) *xxx_MIBEntrySetOperation {
	if op == nil {
		op = &xxx_MIBEntrySetOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.Info = o.Info
	return op
}

func (o *MIBEntrySetRequest) xxx_FromOp(ctx context.Context, op *xxx_MIBEntrySetOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.Info = op.Info
}
func (o *MIBEntrySetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MIBEntrySetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntrySetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MIBEntrySetResponse structure represents the RMIBEntrySet operation response
type MIBEntrySetResponse struct {
	// Return: The RMIBEntrySet return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MIBEntrySetResponse) xxx_ToOp(ctx context.Context, op *xxx_MIBEntrySetOperation) *xxx_MIBEntrySetOperation {
	if op == nil {
		op = &xxx_MIBEntrySetOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *MIBEntrySetResponse) xxx_FromOp(ctx context.Context, op *xxx_MIBEntrySetOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *MIBEntrySetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MIBEntrySetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntrySetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MIBEntryGetOperation structure represents the RMIBEntryGet operation
type xxx_MIBEntryGetOperation struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
	Return     uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_MIBEntryGetOperation) OpNum() int { return 29 }

func (o *xxx_MIBEntryGetOperation) OpName() string { return "/dimsvc/v0/RMIBEntryGet" }

func (o *xxx_MIBEntryGetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_MIBEntryGetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
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

// MIBEntryGetRequest structure represents the RMIBEntryGet operation request
type MIBEntryGetRequest struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
}

func (o *MIBEntryGetRequest) xxx_ToOp(ctx context.Context, op *xxx_MIBEntryGetOperation) *xxx_MIBEntryGetOperation {
	if op == nil {
		op = &xxx_MIBEntryGetOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.Info = o.Info
	return op
}

func (o *MIBEntryGetRequest) xxx_FromOp(ctx context.Context, op *xxx_MIBEntryGetOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.Info = op.Info
}
func (o *MIBEntryGetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MIBEntryGetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntryGetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MIBEntryGetResponse structure represents the RMIBEntryGet operation response
type MIBEntryGetResponse struct {
	Info *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
	// Return: The RMIBEntryGet return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MIBEntryGetResponse) xxx_ToOp(ctx context.Context, op *xxx_MIBEntryGetOperation) *xxx_MIBEntryGetOperation {
	if op == nil {
		op = &xxx_MIBEntryGetOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *MIBEntryGetResponse) xxx_FromOp(ctx context.Context, op *xxx_MIBEntryGetOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *MIBEntryGetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MIBEntryGetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntryGetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MIBEntryGetFirstOperation structure represents the RMIBEntryGetFirst operation
type xxx_MIBEntryGetFirstOperation struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
	Return     uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_MIBEntryGetFirstOperation) OpNum() int { return 30 }

func (o *xxx_MIBEntryGetFirstOperation) OpName() string { return "/dimsvc/v0/RMIBEntryGetFirst" }

func (o *xxx_MIBEntryGetFirstOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetFirstOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetFirstOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetFirstOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetFirstOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_MIBEntryGetFirstOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
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

// MIBEntryGetFirstRequest structure represents the RMIBEntryGetFirst operation request
type MIBEntryGetFirstRequest struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
}

func (o *MIBEntryGetFirstRequest) xxx_ToOp(ctx context.Context, op *xxx_MIBEntryGetFirstOperation) *xxx_MIBEntryGetFirstOperation {
	if op == nil {
		op = &xxx_MIBEntryGetFirstOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.Info = o.Info
	return op
}

func (o *MIBEntryGetFirstRequest) xxx_FromOp(ctx context.Context, op *xxx_MIBEntryGetFirstOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.Info = op.Info
}
func (o *MIBEntryGetFirstRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MIBEntryGetFirstRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntryGetFirstOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MIBEntryGetFirstResponse structure represents the RMIBEntryGetFirst operation response
type MIBEntryGetFirstResponse struct {
	Info *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
	// Return: The RMIBEntryGetFirst return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MIBEntryGetFirstResponse) xxx_ToOp(ctx context.Context, op *xxx_MIBEntryGetFirstOperation) *xxx_MIBEntryGetFirstOperation {
	if op == nil {
		op = &xxx_MIBEntryGetFirstOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *MIBEntryGetFirstResponse) xxx_FromOp(ctx context.Context, op *xxx_MIBEntryGetFirstOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *MIBEntryGetFirstResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MIBEntryGetFirstResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntryGetFirstOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MIBEntryGetNextOperation structure represents the RMIBEntryGetNext operation
type xxx_MIBEntryGetNextOperation struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
	Return     uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_MIBEntryGetNextOperation) OpNum() int { return 31 }

func (o *xxx_MIBEntryGetNextOperation) OpName() string { return "/dimsvc/v0/RMIBEntryGetNext" }

func (o *xxx_MIBEntryGetNextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetNextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetNextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetNextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBEntryGetNextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_MIBEntryGetNextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
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

// MIBEntryGetNextRequest structure represents the RMIBEntryGetNext operation request
type MIBEntryGetNextRequest struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
}

func (o *MIBEntryGetNextRequest) xxx_ToOp(ctx context.Context, op *xxx_MIBEntryGetNextOperation) *xxx_MIBEntryGetNextOperation {
	if op == nil {
		op = &xxx_MIBEntryGetNextOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.Info = o.Info
	return op
}

func (o *MIBEntryGetNextRequest) xxx_FromOp(ctx context.Context, op *xxx_MIBEntryGetNextOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.Info = op.Info
}
func (o *MIBEntryGetNextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MIBEntryGetNextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntryGetNextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MIBEntryGetNextResponse structure represents the RMIBEntryGetNext operation response
type MIBEntryGetNextResponse struct {
	Info *rrasm.MIBEntryContainer `idl:"name:pInfoStuct" json:"info"`
	// Return: The RMIBEntryGetNext return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MIBEntryGetNextResponse) xxx_ToOp(ctx context.Context, op *xxx_MIBEntryGetNextOperation) *xxx_MIBEntryGetNextOperation {
	if op == nil {
		op = &xxx_MIBEntryGetNextOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *MIBEntryGetNextResponse) xxx_FromOp(ctx context.Context, op *xxx_MIBEntryGetNextOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *MIBEntryGetNextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MIBEntryGetNextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBEntryGetNextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MIBGetTrapInfoOperation structure represents the RMIBGetTrapInfo operation
type xxx_MIBGetTrapInfoOperation struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStruct" json:"info"`
	Return     uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_MIBGetTrapInfoOperation) OpNum() int { return 32 }

func (o *xxx_MIBGetTrapInfoOperation) OpName() string { return "/dimsvc/v0/RMIBGetTrapInfo" }

func (o *xxx_MIBGetTrapInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBGetTrapInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBGetTrapInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingPID); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBGetTrapInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBGetTrapInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_MIBGetTrapInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
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

// MIBGetTrapInfoRequest structure represents the RMIBGetTrapInfo operation request
type MIBGetTrapInfoRequest struct {
	PID        uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.MIBEntryContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *MIBGetTrapInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_MIBGetTrapInfoOperation) *xxx_MIBGetTrapInfoOperation {
	if op == nil {
		op = &xxx_MIBGetTrapInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.Info = o.Info
	return op
}

func (o *MIBGetTrapInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_MIBGetTrapInfoOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.Info = op.Info
}
func (o *MIBGetTrapInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MIBGetTrapInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBGetTrapInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MIBGetTrapInfoResponse structure represents the RMIBGetTrapInfo operation response
type MIBGetTrapInfoResponse struct {
	Info *rrasm.MIBEntryContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RMIBGetTrapInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MIBGetTrapInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_MIBGetTrapInfoOperation) *xxx_MIBGetTrapInfoOperation {
	if op == nil {
		op = &xxx_MIBGetTrapInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *MIBGetTrapInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_MIBGetTrapInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *MIBGetTrapInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MIBGetTrapInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBGetTrapInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MIBSetTrapInfoOperation structure represents the RMIBSetTrapInfo operation
type xxx_MIBSetTrapInfoOperation struct {
	PID             uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID      uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Event           uint64                   `idl:"name:hEvent" json:"event"`
	ClientProcessID uint32                   `idl:"name:dwClientProcessId" json:"client_process_id"`
	Info            *rrasm.MIBEntryContainer `idl:"name:pInfoStruct" json:"info"`
	Return          uint32                   `idl:"name:Return" json:"return"`
}

func (o *xxx_MIBSetTrapInfoOperation) OpNum() int { return 33 }

func (o *xxx_MIBSetTrapInfoOperation) OpName() string { return "/dimsvc/v0/RMIBSetTrapInfo" }

func (o *xxx_MIBSetTrapInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBSetTrapInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingPID); err != nil {
			return err
		}
	}
	// hEvent {in} (1:{public, alias=ULONG_PTR}(uint32_64))
	{
		if err := w.WriteData(ndr.Uint3264(o.Event)); err != nil {
			return err
		}
	}
	// dwClientProcessId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientProcessID); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBSetTrapInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PID); err != nil {
			return err
		}
	}
	// dwRoutingPid {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingPID); err != nil {
			return err
		}
	}
	// hEvent {in} (1:{public, alias=ULONG_PTR}(uint32_64))
	{
		if err := w.ReadData((*ndr.Uint3264)(&o.Event)); err != nil {
			return err
		}
	}
	// dwClientProcessId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientProcessID); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBSetTrapInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MIBSetTrapInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.MIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_MIBSetTrapInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.MIBEntryContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
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

// MIBSetTrapInfoRequest structure represents the RMIBSetTrapInfo operation request
type MIBSetTrapInfoRequest struct {
	PID             uint32                   `idl:"name:dwPid" json:"pid"`
	RoutingPID      uint32                   `idl:"name:dwRoutingPid" json:"routing_pid"`
	Event           uint64                   `idl:"name:hEvent" json:"event"`
	ClientProcessID uint32                   `idl:"name:dwClientProcessId" json:"client_process_id"`
	Info            *rrasm.MIBEntryContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *MIBSetTrapInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_MIBSetTrapInfoOperation) *xxx_MIBSetTrapInfoOperation {
	if op == nil {
		op = &xxx_MIBSetTrapInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.Event = o.Event
	op.ClientProcessID = o.ClientProcessID
	op.Info = o.Info
	return op
}

func (o *MIBSetTrapInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_MIBSetTrapInfoOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.Event = op.Event
	o.ClientProcessID = op.ClientProcessID
	o.Info = op.Info
}
func (o *MIBSetTrapInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MIBSetTrapInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBSetTrapInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MIBSetTrapInfoResponse structure represents the RMIBSetTrapInfo operation response
type MIBSetTrapInfoResponse struct {
	Info *rrasm.MIBEntryContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RMIBSetTrapInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MIBSetTrapInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_MIBSetTrapInfoOperation) *xxx_MIBSetTrapInfoOperation {
	if op == nil {
		op = &xxx_MIBSetTrapInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *MIBSetTrapInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_MIBSetTrapInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *MIBSetTrapInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MIBSetTrapInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MIBSetTrapInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConnectionNotificationOperation structure represents the RRasAdminConnectionNotification operation
type xxx_ConnectionNotificationOperation struct {
	Register          uint32 `idl:"name:fRegister" json:"register"`
	ClientProcessID   uint32 `idl:"name:dwClientProcessId" json:"client_process_id"`
	EventNotification uint64 `idl:"name:hEventNotification" json:"event_notification"`
	Return            uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_ConnectionNotificationOperation) OpNum() int { return 34 }

func (o *xxx_ConnectionNotificationOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminConnectionNotification"
}

func (o *xxx_ConnectionNotificationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionNotificationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// fRegister {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Register); err != nil {
			return err
		}
	}
	// dwClientProcessId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ClientProcessID); err != nil {
			return err
		}
	}
	// hEventNotification {in} (1:{public, alias=ULONG_PTR}(uint32_64))
	{
		if err := w.WriteData(ndr.Uint3264(o.EventNotification)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionNotificationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// fRegister {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Register); err != nil {
			return err
		}
	}
	// dwClientProcessId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ClientProcessID); err != nil {
			return err
		}
	}
	// hEventNotification {in} (1:{public, alias=ULONG_PTR}(uint32_64))
	{
		if err := w.ReadData((*ndr.Uint3264)(&o.EventNotification)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionNotificationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionNotificationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionNotificationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ConnectionNotificationRequest structure represents the RRasAdminConnectionNotification operation request
type ConnectionNotificationRequest struct {
	Register          uint32 `idl:"name:fRegister" json:"register"`
	ClientProcessID   uint32 `idl:"name:dwClientProcessId" json:"client_process_id"`
	EventNotification uint64 `idl:"name:hEventNotification" json:"event_notification"`
}

func (o *ConnectionNotificationRequest) xxx_ToOp(ctx context.Context, op *xxx_ConnectionNotificationOperation) *xxx_ConnectionNotificationOperation {
	if op == nil {
		op = &xxx_ConnectionNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.Register = o.Register
	op.ClientProcessID = o.ClientProcessID
	op.EventNotification = o.EventNotification
	return op
}

func (o *ConnectionNotificationRequest) xxx_FromOp(ctx context.Context, op *xxx_ConnectionNotificationOperation) {
	if o == nil {
		return
	}
	o.Register = op.Register
	o.ClientProcessID = op.ClientProcessID
	o.EventNotification = op.EventNotification
}
func (o *ConnectionNotificationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ConnectionNotificationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionNotificationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConnectionNotificationResponse structure represents the RRasAdminConnectionNotification operation response
type ConnectionNotificationResponse struct {
	// Return: The RRasAdminConnectionNotification return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ConnectionNotificationResponse) xxx_ToOp(ctx context.Context, op *xxx_ConnectionNotificationOperation) *xxx_ConnectionNotificationOperation {
	if op == nil {
		op = &xxx_ConnectionNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *ConnectionNotificationResponse) xxx_FromOp(ctx context.Context, op *xxx_ConnectionNotificationOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ConnectionNotificationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ConnectionNotificationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionNotificationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SendUserMessageOperation structure represents the RRasAdminSendUserMessage operation
type xxx_SendUserMessageOperation struct {
	Connection uint32 `idl:"name:hDimConnection" json:"connection"`
	Message    string `idl:"name:lpwszMessage;string" json:"message"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SendUserMessageOperation) OpNum() int { return 35 }

func (o *xxx_SendUserMessageOperation) OpName() string { return "/dimsvc/v0/RRasAdminSendUserMessage" }

func (o *xxx_SendUserMessageOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendUserMessageOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Connection); err != nil {
			return err
		}
	}
	// lpwszMessage {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Message); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendUserMessageOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Connection); err != nil {
			return err
		}
	}
	// lpwszMessage {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Message); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendUserMessageOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendUserMessageOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendUserMessageOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SendUserMessageRequest structure represents the RRasAdminSendUserMessage operation request
type SendUserMessageRequest struct {
	Connection uint32 `idl:"name:hDimConnection" json:"connection"`
	Message    string `idl:"name:lpwszMessage;string" json:"message"`
}

func (o *SendUserMessageRequest) xxx_ToOp(ctx context.Context, op *xxx_SendUserMessageOperation) *xxx_SendUserMessageOperation {
	if op == nil {
		op = &xxx_SendUserMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.Connection = o.Connection
	op.Message = o.Message
	return op
}

func (o *SendUserMessageRequest) xxx_FromOp(ctx context.Context, op *xxx_SendUserMessageOperation) {
	if o == nil {
		return
	}
	o.Connection = op.Connection
	o.Message = op.Message
}
func (o *SendUserMessageRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SendUserMessageRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SendUserMessageOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SendUserMessageResponse structure represents the RRasAdminSendUserMessage operation response
type SendUserMessageResponse struct {
	// Return: The RRasAdminSendUserMessage return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SendUserMessageResponse) xxx_ToOp(ctx context.Context, op *xxx_SendUserMessageOperation) *xxx_SendUserMessageOperation {
	if op == nil {
		op = &xxx_SendUserMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SendUserMessageResponse) xxx_FromOp(ctx context.Context, op *xxx_SendUserMessageOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SendUserMessageResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SendUserMessageResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SendUserMessageOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterDeviceEnumOperation structure represents the RRouterDeviceEnum operation
type xxx_RouterDeviceEnumOperation struct {
	Level        uint32                      `idl:"name:dwLevel" json:"level"`
	Info         *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	TotalEntries uint32                      `idl:"name:lpdwTotalEntries" json:"total_entries"`
	Return       uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterDeviceEnumOperation) OpNum() int { return 36 }

func (o *xxx_RouterDeviceEnumOperation) OpName() string { return "/dimsvc/v0/RRouterDeviceEnum" }

func (o *xxx_RouterDeviceEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterDeviceEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpdwTotalEntries {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TotalEntries); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterDeviceEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpdwTotalEntries {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TotalEntries); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterDeviceEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterDeviceEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpdwTotalEntries {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TotalEntries); err != nil {
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

func (o *xxx_RouterDeviceEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpdwTotalEntries {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TotalEntries); err != nil {
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

// RouterDeviceEnumRequest structure represents the RRouterDeviceEnum operation request
type RouterDeviceEnumRequest struct {
	Level        uint32                      `idl:"name:dwLevel" json:"level"`
	Info         *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	TotalEntries uint32                      `idl:"name:lpdwTotalEntries" json:"total_entries"`
}

func (o *RouterDeviceEnumRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterDeviceEnumOperation) *xxx_RouterDeviceEnumOperation {
	if op == nil {
		op = &xxx_RouterDeviceEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	op.TotalEntries = o.TotalEntries
	return op
}

func (o *RouterDeviceEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterDeviceEnumOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.TotalEntries = op.TotalEntries
}
func (o *RouterDeviceEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterDeviceEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterDeviceEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterDeviceEnumResponse structure represents the RRouterDeviceEnum operation response
type RouterDeviceEnumResponse struct {
	Info         *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	TotalEntries uint32                      `idl:"name:lpdwTotalEntries" json:"total_entries"`
	// Return: The RRouterDeviceEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterDeviceEnumResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterDeviceEnumOperation) *xxx_RouterDeviceEnumOperation {
	if op == nil {
		op = &xxx_RouterDeviceEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.TotalEntries = o.TotalEntries
	op.Return = o.Return
	return op
}

func (o *RouterDeviceEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterDeviceEnumOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.TotalEntries = op.TotalEntries
	o.Return = op.Return
}
func (o *RouterDeviceEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterDeviceEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterDeviceEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceTransportCreateOperation structure represents the RRouterInterfaceTransportCreate operation
type xxx_RouterInterfaceTransportCreateOperation struct {
	TransportID   uint32                    `idl:"name:dwTransportId" json:"transport_id"`
	TransportName string                    `idl:"name:lpwsTransportName;string" json:"transport_name"`
	Info          *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	DLLPath       string                    `idl:"name:lpwsDLLPath;string" json:"dll_path"`
	Return        uint32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceTransportCreateOperation) OpNum() int { return 37 }

func (o *xxx_RouterInterfaceTransportCreateOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceTransportCreate"
}

func (o *xxx_RouterInterfaceTransportCreateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportCreateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TransportID); err != nil {
			return err
		}
	}
	// lpwsTransportName {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.TransportName); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INTERFACE_CONTAINER}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpwsDLLPath {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DLLPath); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportCreateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwTransportId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TransportID); err != nil {
			return err
		}
	}
	// lpwsTransportName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.TransportName); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INTERFACE_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InterfaceContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpwsDLLPath {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DLLPath); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportCreateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportCreateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceTransportCreateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceTransportCreateRequest structure represents the RRouterInterfaceTransportCreate operation request
type RouterInterfaceTransportCreateRequest struct {
	TransportID   uint32                    `idl:"name:dwTransportId" json:"transport_id"`
	TransportName string                    `idl:"name:lpwsTransportName;string" json:"transport_name"`
	Info          *rrasm.InterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	DLLPath       string                    `idl:"name:lpwsDLLPath;string" json:"dll_path"`
}

func (o *RouterInterfaceTransportCreateRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportCreateOperation) *xxx_RouterInterfaceTransportCreateOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportCreateOperation{}
	}
	if o == nil {
		return op
	}
	op.TransportID = o.TransportID
	op.TransportName = o.TransportName
	op.Info = o.Info
	op.DLLPath = o.DLLPath
	return op
}

func (o *RouterInterfaceTransportCreateRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportCreateOperation) {
	if o == nil {
		return
	}
	o.TransportID = op.TransportID
	o.TransportName = op.TransportName
	o.Info = op.Info
	o.DLLPath = op.DLLPath
}
func (o *RouterInterfaceTransportCreateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceTransportCreateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportCreateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceTransportCreateResponse structure represents the RRouterInterfaceTransportCreate operation response
type RouterInterfaceTransportCreateResponse struct {
	// Return: The RRouterInterfaceTransportCreate return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceTransportCreateResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportCreateOperation) *xxx_RouterInterfaceTransportCreateOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportCreateOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceTransportCreateResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportCreateOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceTransportCreateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceTransportCreateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceTransportCreateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceDeviceGetInfoOperation structure represents the RRouterInterfaceDeviceGetInfo operation
type xxx_RouterInterfaceDeviceGetInfoOperation struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Index     uint32                      `idl:"name:dwIndex" json:"index"`
	Interface uint32                      `idl:"name:hInterface" json:"interface"`
	Return    uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceDeviceGetInfoOperation) OpNum() int { return 38 }

func (o *xxx_RouterInterfaceDeviceGetInfoOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceDeviceGetInfo"
}

func (o *xxx_RouterInterfaceDeviceGetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeviceGetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeviceGetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeviceGetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeviceGetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RouterInterfaceDeviceGetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
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

// RouterInterfaceDeviceGetInfoRequest structure represents the RRouterInterfaceDeviceGetInfo operation request
type RouterInterfaceDeviceGetInfoRequest struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Index     uint32                      `idl:"name:dwIndex" json:"index"`
	Interface uint32                      `idl:"name:hInterface" json:"interface"`
}

func (o *RouterInterfaceDeviceGetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceDeviceGetInfoOperation) *xxx_RouterInterfaceDeviceGetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceDeviceGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	op.Index = o.Index
	op.Interface = o.Interface
	return op
}

func (o *RouterInterfaceDeviceGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceDeviceGetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.Index = op.Index
	o.Interface = op.Interface
}
func (o *RouterInterfaceDeviceGetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceDeviceGetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceDeviceGetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceDeviceGetInfoResponse structure represents the RRouterInterfaceDeviceGetInfo operation response
type RouterInterfaceDeviceGetInfoResponse struct {
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RRouterInterfaceDeviceGetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceDeviceGetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceDeviceGetInfoOperation) *xxx_RouterInterfaceDeviceGetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceDeviceGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceDeviceGetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceDeviceGetInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *RouterInterfaceDeviceGetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceDeviceGetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceDeviceGetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceDeviceSetInfoOperation structure represents the RRouterInterfaceDeviceSetInfo operation
type xxx_RouterInterfaceDeviceSetInfoOperation struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Index     uint32                      `idl:"name:dwIndex" json:"index"`
	Interface uint32                      `idl:"name:hInterface" json:"interface"`
	Return    uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceDeviceSetInfoOperation) OpNum() int { return 39 }

func (o *xxx_RouterInterfaceDeviceSetInfoOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceDeviceSetInfo"
}

func (o *xxx_RouterInterfaceDeviceSetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeviceSetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeviceSetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeviceSetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeviceSetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeviceSetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceDeviceSetInfoRequest structure represents the RRouterInterfaceDeviceSetInfo operation request
type RouterInterfaceDeviceSetInfoRequest struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Index     uint32                      `idl:"name:dwIndex" json:"index"`
	Interface uint32                      `idl:"name:hInterface" json:"interface"`
}

func (o *RouterInterfaceDeviceSetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceDeviceSetInfoOperation) *xxx_RouterInterfaceDeviceSetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceDeviceSetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	op.Index = o.Index
	op.Interface = o.Interface
	return op
}

func (o *RouterInterfaceDeviceSetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceDeviceSetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.Index = op.Index
	o.Interface = op.Interface
}
func (o *RouterInterfaceDeviceSetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceDeviceSetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceDeviceSetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceDeviceSetInfoResponse structure represents the RRouterInterfaceDeviceSetInfo operation response
type RouterInterfaceDeviceSetInfoResponse struct {
	// Return: The RRouterInterfaceDeviceSetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceDeviceSetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceDeviceSetInfoOperation) *xxx_RouterInterfaceDeviceSetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceDeviceSetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceDeviceSetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceDeviceSetInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceDeviceSetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceDeviceSetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceDeviceSetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceSetCredentialsExOperation structure represents the RRouterInterfaceSetCredentialsEx operation
type xxx_RouterInterfaceSetCredentialsExOperation struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Interface uint32                      `idl:"name:hInterface" json:"interface"`
	Return    uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceSetCredentialsExOperation) OpNum() int { return 40 }

func (o *xxx_RouterInterfaceSetCredentialsExOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceSetCredentialsEx"
}

func (o *xxx_RouterInterfaceSetCredentialsExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCredentialsExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCredentialsExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCredentialsExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCredentialsExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCredentialsExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceSetCredentialsExRequest structure represents the RRouterInterfaceSetCredentialsEx operation request
type RouterInterfaceSetCredentialsExRequest struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Interface uint32                      `idl:"name:hInterface" json:"interface"`
}

func (o *RouterInterfaceSetCredentialsExRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceSetCredentialsExOperation) *xxx_RouterInterfaceSetCredentialsExOperation {
	if op == nil {
		op = &xxx_RouterInterfaceSetCredentialsExOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	op.Interface = o.Interface
	return op
}

func (o *RouterInterfaceSetCredentialsExRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceSetCredentialsExOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.Interface = op.Interface
}
func (o *RouterInterfaceSetCredentialsExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceSetCredentialsExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceSetCredentialsExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceSetCredentialsExResponse structure represents the RRouterInterfaceSetCredentialsEx operation response
type RouterInterfaceSetCredentialsExResponse struct {
	// Return: The RRouterInterfaceSetCredentialsEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceSetCredentialsExResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceSetCredentialsExOperation) *xxx_RouterInterfaceSetCredentialsExOperation {
	if op == nil {
		op = &xxx_RouterInterfaceSetCredentialsExOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceSetCredentialsExResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceSetCredentialsExOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceSetCredentialsExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceSetCredentialsExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceSetCredentialsExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceGetCredentialsExOperation structure represents the RRouterInterfaceGetCredentialsEx operation
type xxx_RouterInterfaceGetCredentialsExOperation struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Interface uint32                      `idl:"name:hInterface" json:"interface"`
	Return    uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceGetCredentialsExOperation) OpNum() int { return 41 }

func (o *xxx_RouterInterfaceGetCredentialsExOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceGetCredentialsEx"
}

func (o *xxx_RouterInterfaceGetCredentialsExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCredentialsExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCredentialsExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCredentialsExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCredentialsExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RouterInterfaceGetCredentialsExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
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

// RouterInterfaceGetCredentialsExRequest structure represents the RRouterInterfaceGetCredentialsEx operation request
type RouterInterfaceGetCredentialsExRequest struct {
	Level     uint32                      `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Interface uint32                      `idl:"name:hInterface" json:"interface"`
}

func (o *RouterInterfaceGetCredentialsExRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetCredentialsExOperation) *xxx_RouterInterfaceGetCredentialsExOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetCredentialsExOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	op.Interface = o.Interface
	return op
}

func (o *RouterInterfaceGetCredentialsExRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetCredentialsExOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.Interface = op.Interface
}
func (o *RouterInterfaceGetCredentialsExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceGetCredentialsExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceGetCredentialsExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceGetCredentialsExResponse structure represents the RRouterInterfaceGetCredentialsEx operation response
type RouterInterfaceGetCredentialsExResponse struct {
	Info *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RRouterInterfaceGetCredentialsEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceGetCredentialsExResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetCredentialsExOperation) *xxx_RouterInterfaceGetCredentialsExOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetCredentialsExOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceGetCredentialsExResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetCredentialsExOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *RouterInterfaceGetCredentialsExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceGetCredentialsExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceGetCredentialsExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConnectionRemoveQuarantineOperation structure represents the RRasAdminConnectionRemoveQuarantine operation
type xxx_ConnectionRemoveQuarantineOperation struct {
	Connection  uint32 `idl:"name:hRasConnection" json:"connection"`
	IsIPAddress bool   `idl:"name:fIsIpAddress" json:"is_ip_address"`
	Return      uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_ConnectionRemoveQuarantineOperation) OpNum() int { return 42 }

func (o *xxx_ConnectionRemoveQuarantineOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminConnectionRemoveQuarantine"
}

func (o *xxx_ConnectionRemoveQuarantineOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionRemoveQuarantineOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRasConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Connection); err != nil {
			return err
		}
	}
	// fIsIpAddress {in} (1:{alias=BOOL}(int32))
	{
		if !o.IsIPAddress {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ConnectionRemoveQuarantineOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRasConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Connection); err != nil {
			return err
		}
	}
	// fIsIpAddress {in} (1:{alias=BOOL}(int32))
	{
		var _bIsIPAddress int32
		if err := w.ReadData(&_bIsIPAddress); err != nil {
			return err
		}
		o.IsIPAddress = _bIsIPAddress != 0
	}
	return nil
}

func (o *xxx_ConnectionRemoveQuarantineOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionRemoveQuarantineOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionRemoveQuarantineOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ConnectionRemoveQuarantineRequest structure represents the RRasAdminConnectionRemoveQuarantine operation request
type ConnectionRemoveQuarantineRequest struct {
	Connection  uint32 `idl:"name:hRasConnection" json:"connection"`
	IsIPAddress bool   `idl:"name:fIsIpAddress" json:"is_ip_address"`
}

func (o *ConnectionRemoveQuarantineRequest) xxx_ToOp(ctx context.Context, op *xxx_ConnectionRemoveQuarantineOperation) *xxx_ConnectionRemoveQuarantineOperation {
	if op == nil {
		op = &xxx_ConnectionRemoveQuarantineOperation{}
	}
	if o == nil {
		return op
	}
	op.Connection = o.Connection
	op.IsIPAddress = o.IsIPAddress
	return op
}

func (o *ConnectionRemoveQuarantineRequest) xxx_FromOp(ctx context.Context, op *xxx_ConnectionRemoveQuarantineOperation) {
	if o == nil {
		return
	}
	o.Connection = op.Connection
	o.IsIPAddress = op.IsIPAddress
}
func (o *ConnectionRemoveQuarantineRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ConnectionRemoveQuarantineRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionRemoveQuarantineOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConnectionRemoveQuarantineResponse structure represents the RRasAdminConnectionRemoveQuarantine operation response
type ConnectionRemoveQuarantineResponse struct {
	// Return: The RRasAdminConnectionRemoveQuarantine return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ConnectionRemoveQuarantineResponse) xxx_ToOp(ctx context.Context, op *xxx_ConnectionRemoveQuarantineOperation) *xxx_ConnectionRemoveQuarantineOperation {
	if op == nil {
		op = &xxx_ConnectionRemoveQuarantineOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *ConnectionRemoveQuarantineResponse) xxx_FromOp(ctx context.Context, op *xxx_ConnectionRemoveQuarantineOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ConnectionRemoveQuarantineResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ConnectionRemoveQuarantineResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionRemoveQuarantineOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ServerSetInfoOperation structure represents the RMprAdminServerSetInfo operation
type xxx_ServerSetInfoOperation struct {
	Level  uint32                      `idl:"name:dwLevel" json:"level"`
	Info   *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
	Return uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_ServerSetInfoOperation) OpNum() int { return 43 }

func (o *xxx_ServerSetInfoOperation) OpName() string { return "/dimsvc/v0/RMprAdminServerSetInfo" }

func (o *xxx_ServerSetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.InformationContainer{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ServerSetInfoRequest structure represents the RMprAdminServerSetInfo operation request
type ServerSetInfoRequest struct {
	Level uint32                      `idl:"name:dwLevel" json:"level"`
	Info  *rrasm.InformationContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *ServerSetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_ServerSetInfoOperation) *xxx_ServerSetInfoOperation {
	if op == nil {
		op = &xxx_ServerSetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	return op
}

func (o *ServerSetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_ServerSetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
}
func (o *ServerSetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ServerSetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerSetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ServerSetInfoResponse structure represents the RMprAdminServerSetInfo operation response
type ServerSetInfoResponse struct {
	// Return: The RMprAdminServerSetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ServerSetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_ServerSetInfoOperation) *xxx_ServerSetInfoOperation {
	if op == nil {
		op = &xxx_ServerSetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *ServerSetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_ServerSetInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ServerSetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ServerSetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerSetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ServerGetInfoExOperation structure represents the RMprAdminServerGetInfoEx operation
type xxx_ServerGetInfoExOperation struct {
	ServerConfig *rrasm.ServerExIDL `idl:"name:pServerConfig" json:"server_config"`
	Return       uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_ServerGetInfoExOperation) OpNum() int { return 44 }

func (o *xxx_ServerGetInfoExOperation) OpName() string { return "/dimsvc/v0/RMprAdminServerGetInfoEx" }

func (o *xxx_ServerGetInfoExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetInfoExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pServerConfig {in, out} (1:{alias=PMPR_SERVER_EX_IDL}*(1))(2:{alias=MPR_SERVER_EX_IDL}(union))
	{
		if o.ServerConfig != nil {
			if err := o.ServerConfig.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.ServerExIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetInfoExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pServerConfig {in, out} (1:{alias=PMPR_SERVER_EX_IDL,pointer=ref}*(1))(2:{alias=MPR_SERVER_EX_IDL}(union))
	{
		if o.ServerConfig == nil {
			o.ServerConfig = &rrasm.ServerExIDL{}
		}
		if err := o.ServerConfig.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetInfoExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerGetInfoExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pServerConfig {in, out} (1:{alias=PMPR_SERVER_EX_IDL}*(1))(2:{alias=MPR_SERVER_EX_IDL}(union))
	{
		if o.ServerConfig != nil {
			if err := o.ServerConfig.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.ServerExIDL{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_ServerGetInfoExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pServerConfig {in, out} (1:{alias=PMPR_SERVER_EX_IDL,pointer=ref}*(1))(2:{alias=MPR_SERVER_EX_IDL}(union))
	{
		if o.ServerConfig == nil {
			o.ServerConfig = &rrasm.ServerExIDL{}
		}
		if err := o.ServerConfig.UnmarshalNDR(ctx, w); err != nil {
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

// ServerGetInfoExRequest structure represents the RMprAdminServerGetInfoEx operation request
type ServerGetInfoExRequest struct {
	ServerConfig *rrasm.ServerExIDL `idl:"name:pServerConfig" json:"server_config"`
}

func (o *ServerGetInfoExRequest) xxx_ToOp(ctx context.Context, op *xxx_ServerGetInfoExOperation) *xxx_ServerGetInfoExOperation {
	if op == nil {
		op = &xxx_ServerGetInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerConfig = o.ServerConfig
	return op
}

func (o *ServerGetInfoExRequest) xxx_FromOp(ctx context.Context, op *xxx_ServerGetInfoExOperation) {
	if o == nil {
		return
	}
	o.ServerConfig = op.ServerConfig
}
func (o *ServerGetInfoExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ServerGetInfoExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerGetInfoExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ServerGetInfoExResponse structure represents the RMprAdminServerGetInfoEx operation response
type ServerGetInfoExResponse struct {
	ServerConfig *rrasm.ServerExIDL `idl:"name:pServerConfig" json:"server_config"`
	// Return: The RMprAdminServerGetInfoEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ServerGetInfoExResponse) xxx_ToOp(ctx context.Context, op *xxx_ServerGetInfoExOperation) *xxx_ServerGetInfoExOperation {
	if op == nil {
		op = &xxx_ServerGetInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerConfig = o.ServerConfig
	op.Return = o.Return
	return op
}

func (o *ServerGetInfoExResponse) xxx_FromOp(ctx context.Context, op *xxx_ServerGetInfoExOperation) {
	if o == nil {
		return
	}
	o.ServerConfig = op.ServerConfig
	o.Return = op.Return
}
func (o *ServerGetInfoExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ServerGetInfoExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerGetInfoExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConnectionEnumExOperation structure represents the RRasAdminConnectionEnumEx operation
type xxx_ConnectionEnumExOperation struct {
	ObjectHeader        *rrasm.ObjectHeaderIDL      `idl:"name:objectHeader" json:"object_header"`
	PreferredMaxLength  uint32                      `idl:"name:dwPreferedMaxLen" json:"preferred_max_length"`
	EntriesRead         uint32                      `idl:"name:lpdwEntriesRead" json:"entries_read"`
	LpdNumTotalElements uint32                      `idl:"name:lpdNumTotalElements" json:"lpd_num_total_elements"`
	RASConections       []*rrasm.RASConnectionExIDL `idl:"name:pRasConections;size_is:(, lpdwEntriesRead)" json:"ras_conections"`
	Resume              uint32                      `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	Return              uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_ConnectionEnumExOperation) OpNum() int { return 45 }

func (o *xxx_ConnectionEnumExOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminConnectionEnumEx"
}

func (o *xxx_ConnectionEnumExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionEnumExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// objectHeader {in} (1:{alias=PMPRAPI_OBJECT_HEADER_IDL}*(1))(2:{alias=MPRAPI_OBJECT_HEADER_IDL}(struct))
	{
		if o.ObjectHeader != nil {
			if err := o.ObjectHeader.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.ObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwPreferedMaxLen {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PreferredMaxLength); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_default_null([cond]) attribute.
		_ptr_lpdwResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionEnumExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// objectHeader {in} (1:{alias=PMPRAPI_OBJECT_HEADER_IDL,pointer=ref}*(1))(2:{alias=MPRAPI_OBJECT_HEADER_IDL}(struct))
	{
		if o.ObjectHeader == nil {
			o.ObjectHeader = &rrasm.ObjectHeaderIDL{}
		}
		if err := o.ObjectHeader.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwPreferedMaxLen {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PreferredMaxLength); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpdwResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_lpdwResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_lpdwResumeHandle, _ptr_lpdwResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionEnumExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.RASConections != nil && o.EntriesRead == 0 {
		o.EntriesRead = uint32(len(o.RASConections))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionEnumExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpdwEntriesRead {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.EntriesRead); err != nil {
			return err
		}
	}
	// lpdNumTotalElements {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.LpdNumTotalElements); err != nil {
			return err
		}
	}
	// pRasConections {out} (1:{pointer=ref}*(2))(2:{alias=PRAS_CONNECTION_EX_IDL}*(1))(3:{alias=RAS_CONNECTION_EX_IDL}[dim:0,size_is=lpdwEntriesRead](union))
	{
		if o.RASConections != nil || o.EntriesRead > 0 {
			_ptr_pRasConections := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.EntriesRead)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.RASConections {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.RASConections[i1] != nil {
						if err := o.RASConections[i1].MarshalNDR(ctx, w); err != nil {
							return err
						}
					} else {
						if err := (&rrasm.RASConnectionExIDL{}).MarshalNDR(ctx, w); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.RASConections); uint64(i1) < sizeInfo[0]; i1++ {
					if err := (&rrasm.RASConnectionExIDL{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RASConections, _ptr_pRasConections); err != nil {
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
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_default_null([cond]) attribute.
		_ptr_lpdwResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_lpdwResumeHandle); err != nil {
			return err
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

func (o *xxx_ConnectionEnumExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpdwEntriesRead {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.EntriesRead); err != nil {
			return err
		}
	}
	// lpdNumTotalElements {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.LpdNumTotalElements); err != nil {
			return err
		}
	}
	// pRasConections {out} (1:{pointer=ref}*(2))(2:{alias=PRAS_CONNECTION_EX_IDL,pointer=ref}*(1))(3:{alias=RAS_CONNECTION_EX_IDL}[dim:0,size_is=lpdwEntriesRead](union))
	{
		_ptr_pRasConections := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.RASConections", sizeInfo[0])
			}
			o.RASConections = make([]*rrasm.RASConnectionExIDL, sizeInfo[0])
			for i1 := range o.RASConections {
				i1 := i1
				if o.RASConections[i1] == nil {
					o.RASConections[i1] = &rrasm.RASConnectionExIDL{}
				}
				if err := o.RASConections[i1].UnmarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pRasConections := func(ptr interface{}) { o.RASConections = *ptr.(*[]*rrasm.RASConnectionExIDL) }
		if err := w.ReadPointer(&o.RASConections, _s_pRasConections, _ptr_pRasConections); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpdwResumeHandle {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpdwResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_lpdwResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_lpdwResumeHandle, _ptr_lpdwResumeHandle); err != nil {
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

// ConnectionEnumExRequest structure represents the RRasAdminConnectionEnumEx operation request
type ConnectionEnumExRequest struct {
	ObjectHeader       *rrasm.ObjectHeaderIDL `idl:"name:objectHeader" json:"object_header"`
	PreferredMaxLength uint32                 `idl:"name:dwPreferedMaxLen" json:"preferred_max_length"`
	Resume             uint32                 `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
}

func (o *ConnectionEnumExRequest) xxx_ToOp(ctx context.Context, op *xxx_ConnectionEnumExOperation) *xxx_ConnectionEnumExOperation {
	if op == nil {
		op = &xxx_ConnectionEnumExOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectHeader = o.ObjectHeader
	op.PreferredMaxLength = o.PreferredMaxLength
	op.Resume = o.Resume
	return op
}

func (o *ConnectionEnumExRequest) xxx_FromOp(ctx context.Context, op *xxx_ConnectionEnumExOperation) {
	if o == nil {
		return
	}
	o.ObjectHeader = op.ObjectHeader
	o.PreferredMaxLength = op.PreferredMaxLength
	o.Resume = op.Resume
}
func (o *ConnectionEnumExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ConnectionEnumExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionEnumExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConnectionEnumExResponse structure represents the RRasAdminConnectionEnumEx operation response
type ConnectionEnumExResponse struct {
	EntriesRead         uint32                      `idl:"name:lpdwEntriesRead" json:"entries_read"`
	LpdNumTotalElements uint32                      `idl:"name:lpdNumTotalElements" json:"lpd_num_total_elements"`
	RASConections       []*rrasm.RASConnectionExIDL `idl:"name:pRasConections;size_is:(, lpdwEntriesRead)" json:"ras_conections"`
	Resume              uint32                      `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	// Return: The RRasAdminConnectionEnumEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ConnectionEnumExResponse) xxx_ToOp(ctx context.Context, op *xxx_ConnectionEnumExOperation) *xxx_ConnectionEnumExOperation {
	if op == nil {
		op = &xxx_ConnectionEnumExOperation{}
	}
	if o == nil {
		return op
	}
	op.EntriesRead = o.EntriesRead
	op.LpdNumTotalElements = o.LpdNumTotalElements
	op.RASConections = o.RASConections
	op.Resume = o.Resume
	op.Return = o.Return
	return op
}

func (o *ConnectionEnumExResponse) xxx_FromOp(ctx context.Context, op *xxx_ConnectionEnumExOperation) {
	if o == nil {
		return
	}
	o.EntriesRead = op.EntriesRead
	o.LpdNumTotalElements = op.LpdNumTotalElements
	o.RASConections = op.RASConections
	o.Resume = op.Resume
	o.Return = op.Return
}
func (o *ConnectionEnumExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ConnectionEnumExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionEnumExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConnectionGetInfoExOperation structure represents the RRasAdminConnectionGetInfoEx operation
type xxx_ConnectionGetInfoExOperation struct {
	Connection    uint32                    `idl:"name:hDimConnection" json:"connection"`
	ObjectHeader  *rrasm.ObjectHeaderIDL    `idl:"name:objectHeader" json:"object_header"`
	RASConnection *rrasm.RASConnectionExIDL `idl:"name:pRasConnection" json:"ras_connection"`
	Return        uint32                    `idl:"name:Return" json:"return"`
}

func (o *xxx_ConnectionGetInfoExOperation) OpNum() int { return 46 }

func (o *xxx_ConnectionGetInfoExOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminConnectionGetInfoEx"
}

func (o *xxx_ConnectionGetInfoExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionGetInfoExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Connection); err != nil {
			return err
		}
	}
	// objectHeader {in} (1:{alias=PMPRAPI_OBJECT_HEADER_IDL}*(1))(2:{alias=MPRAPI_OBJECT_HEADER_IDL}(struct))
	{
		if o.ObjectHeader != nil {
			if err := o.ObjectHeader.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.ObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ConnectionGetInfoExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Connection); err != nil {
			return err
		}
	}
	// objectHeader {in} (1:{alias=PMPRAPI_OBJECT_HEADER_IDL,pointer=ref}*(1))(2:{alias=MPRAPI_OBJECT_HEADER_IDL}(struct))
	{
		if o.ObjectHeader == nil {
			o.ObjectHeader = &rrasm.ObjectHeaderIDL{}
		}
		if err := o.ObjectHeader.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionGetInfoExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionGetInfoExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pRasConnection {out} (1:{alias=PRAS_CONNECTION_EX_IDL}*(1))(2:{alias=RAS_CONNECTION_EX_IDL}(union))
	{
		if o.RASConnection != nil {
			if err := o.RASConnection.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.RASConnectionExIDL{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_ConnectionGetInfoExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pRasConnection {out} (1:{alias=PRAS_CONNECTION_EX_IDL,pointer=ref}*(1))(2:{alias=RAS_CONNECTION_EX_IDL}(union))
	{
		if o.RASConnection == nil {
			o.RASConnection = &rrasm.RASConnectionExIDL{}
		}
		if err := o.RASConnection.UnmarshalNDR(ctx, w); err != nil {
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

// ConnectionGetInfoExRequest structure represents the RRasAdminConnectionGetInfoEx operation request
type ConnectionGetInfoExRequest struct {
	Connection   uint32                 `idl:"name:hDimConnection" json:"connection"`
	ObjectHeader *rrasm.ObjectHeaderIDL `idl:"name:objectHeader" json:"object_header"`
}

func (o *ConnectionGetInfoExRequest) xxx_ToOp(ctx context.Context, op *xxx_ConnectionGetInfoExOperation) *xxx_ConnectionGetInfoExOperation {
	if op == nil {
		op = &xxx_ConnectionGetInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.Connection = o.Connection
	op.ObjectHeader = o.ObjectHeader
	return op
}

func (o *ConnectionGetInfoExRequest) xxx_FromOp(ctx context.Context, op *xxx_ConnectionGetInfoExOperation) {
	if o == nil {
		return
	}
	o.Connection = op.Connection
	o.ObjectHeader = op.ObjectHeader
}
func (o *ConnectionGetInfoExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ConnectionGetInfoExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionGetInfoExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConnectionGetInfoExResponse structure represents the RRasAdminConnectionGetInfoEx operation response
type ConnectionGetInfoExResponse struct {
	RASConnection *rrasm.RASConnectionExIDL `idl:"name:pRasConnection" json:"ras_connection"`
	// Return: The RRasAdminConnectionGetInfoEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ConnectionGetInfoExResponse) xxx_ToOp(ctx context.Context, op *xxx_ConnectionGetInfoExOperation) *xxx_ConnectionGetInfoExOperation {
	if op == nil {
		op = &xxx_ConnectionGetInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.RASConnection = o.RASConnection
	op.Return = o.Return
	return op
}

func (o *ConnectionGetInfoExResponse) xxx_FromOp(ctx context.Context, op *xxx_ConnectionGetInfoExOperation) {
	if o == nil {
		return
	}
	o.RASConnection = op.RASConnection
	o.Return = op.Return
}
func (o *ConnectionGetInfoExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ConnectionGetInfoExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionGetInfoExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ServerSetInfoExOperation structure represents the RMprAdminServerSetInfoEx operation
type xxx_ServerSetInfoExOperation struct {
	ServerConfig *rrasm.ServerSetConfigExIDL `idl:"name:pServerConfig" json:"server_config"`
	Return       uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_ServerSetInfoExOperation) OpNum() int { return 47 }

func (o *xxx_ServerSetInfoExOperation) OpName() string { return "/dimsvc/v0/RMprAdminServerSetInfoEx" }

func (o *xxx_ServerSetInfoExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetInfoExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pServerConfig {in} (1:{alias=PMPR_SERVER_SET_CONFIG_EX_IDL}*(1))(2:{alias=MPR_SERVER_SET_CONFIG_EX_IDL}(union))
	{
		if o.ServerConfig != nil {
			if err := o.ServerConfig.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.ServerSetConfigExIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetInfoExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pServerConfig {in} (1:{alias=PMPR_SERVER_SET_CONFIG_EX_IDL,pointer=ref}*(1))(2:{alias=MPR_SERVER_SET_CONFIG_EX_IDL}(union))
	{
		if o.ServerConfig == nil {
			o.ServerConfig = &rrasm.ServerSetConfigExIDL{}
		}
		if err := o.ServerConfig.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetInfoExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetInfoExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ServerSetInfoExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ServerSetInfoExRequest structure represents the RMprAdminServerSetInfoEx operation request
type ServerSetInfoExRequest struct {
	ServerConfig *rrasm.ServerSetConfigExIDL `idl:"name:pServerConfig" json:"server_config"`
}

func (o *ServerSetInfoExRequest) xxx_ToOp(ctx context.Context, op *xxx_ServerSetInfoExOperation) *xxx_ServerSetInfoExOperation {
	if op == nil {
		op = &xxx_ServerSetInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerConfig = o.ServerConfig
	return op
}

func (o *ServerSetInfoExRequest) xxx_FromOp(ctx context.Context, op *xxx_ServerSetInfoExOperation) {
	if o == nil {
		return
	}
	o.ServerConfig = op.ServerConfig
}
func (o *ServerSetInfoExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ServerSetInfoExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerSetInfoExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ServerSetInfoExResponse structure represents the RMprAdminServerSetInfoEx operation response
type ServerSetInfoExResponse struct {
	// Return: The RMprAdminServerSetInfoEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ServerSetInfoExResponse) xxx_ToOp(ctx context.Context, op *xxx_ServerSetInfoExOperation) *xxx_ServerSetInfoExOperation {
	if op == nil {
		op = &xxx_ServerSetInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *ServerSetInfoExResponse) xxx_FromOp(ctx context.Context, op *xxx_ServerSetInfoExOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ServerSetInfoExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ServerSetInfoExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ServerSetInfoExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UpdateConnectionOperation structure represents the RRasAdminUpdateConnection operation
type xxx_UpdateConnectionOperation struct {
	Connection   uint32                        `idl:"name:hDimConnection" json:"connection"`
	ServerConfig *rrasm.RASUpdateConnectionIDL `idl:"name:pServerConfig" json:"server_config"`
	Return       uint32                        `idl:"name:Return" json:"return"`
}

func (o *xxx_UpdateConnectionOperation) OpNum() int { return 48 }

func (o *xxx_UpdateConnectionOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminUpdateConnection"
}

func (o *xxx_UpdateConnectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateConnectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Connection); err != nil {
			return err
		}
	}
	// pServerConfig {in} (1:{alias=PRAS_UPDATE_CONNECTION_IDL}*(1))(2:{alias=RAS_UPDATE_CONNECTION_IDL}(union))
	{
		if o.ServerConfig != nil {
			if err := o.ServerConfig.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.RASUpdateConnectionIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_UpdateConnectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Connection); err != nil {
			return err
		}
	}
	// pServerConfig {in} (1:{alias=PRAS_UPDATE_CONNECTION_IDL,pointer=ref}*(1))(2:{alias=RAS_UPDATE_CONNECTION_IDL}(union))
	{
		if o.ServerConfig == nil {
			o.ServerConfig = &rrasm.RASUpdateConnectionIDL{}
		}
		if err := o.ServerConfig.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateConnectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateConnectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UpdateConnectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UpdateConnectionRequest structure represents the RRasAdminUpdateConnection operation request
type UpdateConnectionRequest struct {
	Connection   uint32                        `idl:"name:hDimConnection" json:"connection"`
	ServerConfig *rrasm.RASUpdateConnectionIDL `idl:"name:pServerConfig" json:"server_config"`
}

func (o *UpdateConnectionRequest) xxx_ToOp(ctx context.Context, op *xxx_UpdateConnectionOperation) *xxx_UpdateConnectionOperation {
	if op == nil {
		op = &xxx_UpdateConnectionOperation{}
	}
	if o == nil {
		return op
	}
	op.Connection = o.Connection
	op.ServerConfig = o.ServerConfig
	return op
}

func (o *UpdateConnectionRequest) xxx_FromOp(ctx context.Context, op *xxx_UpdateConnectionOperation) {
	if o == nil {
		return
	}
	o.Connection = op.Connection
	o.ServerConfig = op.ServerConfig
}
func (o *UpdateConnectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *UpdateConnectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UpdateConnectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UpdateConnectionResponse structure represents the RRasAdminUpdateConnection operation response
type UpdateConnectionResponse struct {
	// Return: The RRasAdminUpdateConnection return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *UpdateConnectionResponse) xxx_ToOp(ctx context.Context, op *xxx_UpdateConnectionOperation) *xxx_UpdateConnectionOperation {
	if op == nil {
		op = &xxx_UpdateConnectionOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *UpdateConnectionResponse) xxx_FromOp(ctx context.Context, op *xxx_UpdateConnectionOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *UpdateConnectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *UpdateConnectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UpdateConnectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceSetCredentialsLocalOperation structure represents the RRouterInterfaceSetCredentialsLocal operation
type xxx_RouterInterfaceSetCredentialsLocalOperation struct {
	InterfaceName string `idl:"name:lpwsInterfaceName;string" json:"interface_name"`
	UserName      string `idl:"name:lpwsUserName;string" json:"user_name"`
	DomainName    string `idl:"name:lpwsDomainName;string" json:"domain_name"`
	Password      string `idl:"name:lpwsPassword;string" json:"password"`
	Return        uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceSetCredentialsLocalOperation) OpNum() int { return 49 }

func (o *xxx_RouterInterfaceSetCredentialsLocalOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceSetCredentialsLocal"
}

func (o *xxx_RouterInterfaceSetCredentialsLocalOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCredentialsLocalOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpwsInterfaceName {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.InterfaceName); err != nil {
			return err
		}
	}
	// lpwsUserName {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.UserName); err != nil {
			return err
		}
	}
	// lpwsDomainName {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DomainName); err != nil {
			return err
		}
	}
	// lpwsPassword {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Password); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCredentialsLocalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpwsInterfaceName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.InterfaceName); err != nil {
			return err
		}
	}
	// lpwsUserName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.UserName); err != nil {
			return err
		}
	}
	// lpwsDomainName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DomainName); err != nil {
			return err
		}
	}
	// lpwsPassword {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Password); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCredentialsLocalOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCredentialsLocalOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCredentialsLocalOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RouterInterfaceSetCredentialsLocalRequest structure represents the RRouterInterfaceSetCredentialsLocal operation request
type RouterInterfaceSetCredentialsLocalRequest struct {
	InterfaceName string `idl:"name:lpwsInterfaceName;string" json:"interface_name"`
	UserName      string `idl:"name:lpwsUserName;string" json:"user_name"`
	DomainName    string `idl:"name:lpwsDomainName;string" json:"domain_name"`
	Password      string `idl:"name:lpwsPassword;string" json:"password"`
}

func (o *RouterInterfaceSetCredentialsLocalRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceSetCredentialsLocalOperation) *xxx_RouterInterfaceSetCredentialsLocalOperation {
	if op == nil {
		op = &xxx_RouterInterfaceSetCredentialsLocalOperation{}
	}
	if o == nil {
		return op
	}
	op.InterfaceName = o.InterfaceName
	op.UserName = o.UserName
	op.DomainName = o.DomainName
	op.Password = o.Password
	return op
}

func (o *RouterInterfaceSetCredentialsLocalRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceSetCredentialsLocalOperation) {
	if o == nil {
		return
	}
	o.InterfaceName = op.InterfaceName
	o.UserName = op.UserName
	o.DomainName = op.DomainName
	o.Password = op.Password
}
func (o *RouterInterfaceSetCredentialsLocalRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceSetCredentialsLocalRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceSetCredentialsLocalOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceSetCredentialsLocalResponse structure represents the RRouterInterfaceSetCredentialsLocal operation response
type RouterInterfaceSetCredentialsLocalResponse struct {
	// Return: The RRouterInterfaceSetCredentialsLocal return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceSetCredentialsLocalResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceSetCredentialsLocalOperation) *xxx_RouterInterfaceSetCredentialsLocalOperation {
	if op == nil {
		op = &xxx_RouterInterfaceSetCredentialsLocalOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceSetCredentialsLocalResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceSetCredentialsLocalOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RouterInterfaceSetCredentialsLocalResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceSetCredentialsLocalResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceSetCredentialsLocalOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceGetCredentialsLocalOperation structure represents the RRouterInterfaceGetCredentialsLocal operation
type xxx_RouterInterfaceGetCredentialsLocalOperation struct {
	InterfaceName string `idl:"name:lpwsInterfaceName;string" json:"interface_name"`
	UserName      string `idl:"name:lpwsUserName;string" json:"user_name"`
	DomainName    string `idl:"name:lpwsDomainName;string" json:"domain_name"`
	Password      string `idl:"name:lpwsPassword;string" json:"password"`
	Return        uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceGetCredentialsLocalOperation) OpNum() int { return 50 }

func (o *xxx_RouterInterfaceGetCredentialsLocalOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceGetCredentialsLocal"
}

func (o *xxx_RouterInterfaceGetCredentialsLocalOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCredentialsLocalOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpwsInterfaceName {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.InterfaceName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCredentialsLocalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpwsInterfaceName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.InterfaceName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCredentialsLocalOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCredentialsLocalOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpwsUserName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.UserName != "" {
			_ptr_lpwsUserName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.UserName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.UserName, _ptr_lpwsUserName); err != nil {
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
	// lpwsDomainName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.DomainName != "" {
			_ptr_lpwsDomainName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.DomainName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DomainName, _ptr_lpwsDomainName); err != nil {
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
	// lpwsPassword {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Password != "" {
			_ptr_lpwsPassword := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Password); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Password, _ptr_lpwsPassword); err != nil {
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
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCredentialsLocalOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpwsUserName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpwsUserName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.UserName); err != nil {
				return err
			}
			return nil
		})
		_s_lpwsUserName := func(ptr interface{}) { o.UserName = *ptr.(*string) }
		if err := w.ReadPointer(&o.UserName, _s_lpwsUserName, _ptr_lpwsUserName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpwsDomainName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpwsDomainName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.DomainName); err != nil {
				return err
			}
			return nil
		})
		_s_lpwsDomainName := func(ptr interface{}) { o.DomainName = *ptr.(*string) }
		if err := w.ReadPointer(&o.DomainName, _s_lpwsDomainName, _ptr_lpwsDomainName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpwsPassword {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpwsPassword := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Password); err != nil {
				return err
			}
			return nil
		})
		_s_lpwsPassword := func(ptr interface{}) { o.Password = *ptr.(*string) }
		if err := w.ReadPointer(&o.Password, _s_lpwsPassword, _ptr_lpwsPassword); err != nil {
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

// RouterInterfaceGetCredentialsLocalRequest structure represents the RRouterInterfaceGetCredentialsLocal operation request
type RouterInterfaceGetCredentialsLocalRequest struct {
	InterfaceName string `idl:"name:lpwsInterfaceName;string" json:"interface_name"`
}

func (o *RouterInterfaceGetCredentialsLocalRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetCredentialsLocalOperation) *xxx_RouterInterfaceGetCredentialsLocalOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetCredentialsLocalOperation{}
	}
	if o == nil {
		return op
	}
	op.InterfaceName = o.InterfaceName
	return op
}

func (o *RouterInterfaceGetCredentialsLocalRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetCredentialsLocalOperation) {
	if o == nil {
		return
	}
	o.InterfaceName = op.InterfaceName
}
func (o *RouterInterfaceGetCredentialsLocalRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceGetCredentialsLocalRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceGetCredentialsLocalOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceGetCredentialsLocalResponse structure represents the RRouterInterfaceGetCredentialsLocal operation response
type RouterInterfaceGetCredentialsLocalResponse struct {
	UserName   string `idl:"name:lpwsUserName;string" json:"user_name"`
	DomainName string `idl:"name:lpwsDomainName;string" json:"domain_name"`
	Password   string `idl:"name:lpwsPassword;string" json:"password"`
	// Return: The RRouterInterfaceGetCredentialsLocal return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceGetCredentialsLocalResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetCredentialsLocalOperation) *xxx_RouterInterfaceGetCredentialsLocalOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetCredentialsLocalOperation{}
	}
	if o == nil {
		return op
	}
	op.UserName = o.UserName
	op.DomainName = o.DomainName
	op.Password = o.Password
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceGetCredentialsLocalResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetCredentialsLocalOperation) {
	if o == nil {
		return
	}
	o.UserName = op.UserName
	o.DomainName = op.DomainName
	o.Password = op.Password
	o.Return = op.Return
}
func (o *RouterInterfaceGetCredentialsLocalResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceGetCredentialsLocalResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceGetCredentialsLocalOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceGetCustomInfoExOperation structure represents the RRouterInterfaceGetCustomInfoEx operation
type xxx_RouterInterfaceGetCustomInfoExOperation struct {
	Interface             uint32                          `idl:"name:hInterface" json:"interface"`
	InterfaceCustomConfig *rrasm.InterfaceCustominfoexIDL `idl:"name:pIfCustomConfig" json:"interface_custom_config"`
	Return                uint32                          `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceGetCustomInfoExOperation) OpNum() int { return 51 }

func (o *xxx_RouterInterfaceGetCustomInfoExOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceGetCustomInfoEx"
}

func (o *xxx_RouterInterfaceGetCustomInfoExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCustomInfoExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// pIfCustomConfig {in, out} (1:{alias=PMPR_IF_CUSTOMINFOEX_IDL}*(1))(2:{alias=MPR_IF_CUSTOMINFOEX_IDL}(union))
	{
		if o.InterfaceCustomConfig != nil {
			if err := o.InterfaceCustomConfig.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceCustominfoexIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCustomInfoExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// pIfCustomConfig {in, out} (1:{alias=PMPR_IF_CUSTOMINFOEX_IDL,pointer=ref}*(1))(2:{alias=MPR_IF_CUSTOMINFOEX_IDL}(union))
	{
		if o.InterfaceCustomConfig == nil {
			o.InterfaceCustomConfig = &rrasm.InterfaceCustominfoexIDL{}
		}
		if err := o.InterfaceCustomConfig.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCustomInfoExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCustomInfoExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pIfCustomConfig {in, out} (1:{alias=PMPR_IF_CUSTOMINFOEX_IDL}*(1))(2:{alias=MPR_IF_CUSTOMINFOEX_IDL}(union))
	{
		if o.InterfaceCustomConfig != nil {
			if err := o.InterfaceCustomConfig.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceCustominfoexIDL{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RouterInterfaceGetCustomInfoExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pIfCustomConfig {in, out} (1:{alias=PMPR_IF_CUSTOMINFOEX_IDL,pointer=ref}*(1))(2:{alias=MPR_IF_CUSTOMINFOEX_IDL}(union))
	{
		if o.InterfaceCustomConfig == nil {
			o.InterfaceCustomConfig = &rrasm.InterfaceCustominfoexIDL{}
		}
		if err := o.InterfaceCustomConfig.UnmarshalNDR(ctx, w); err != nil {
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

// RouterInterfaceGetCustomInfoExRequest structure represents the RRouterInterfaceGetCustomInfoEx operation request
type RouterInterfaceGetCustomInfoExRequest struct {
	Interface             uint32                          `idl:"name:hInterface" json:"interface"`
	InterfaceCustomConfig *rrasm.InterfaceCustominfoexIDL `idl:"name:pIfCustomConfig" json:"interface_custom_config"`
}

func (o *RouterInterfaceGetCustomInfoExRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetCustomInfoExOperation) *xxx_RouterInterfaceGetCustomInfoExOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetCustomInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.InterfaceCustomConfig = o.InterfaceCustomConfig
	return op
}

func (o *RouterInterfaceGetCustomInfoExRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetCustomInfoExOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.InterfaceCustomConfig = op.InterfaceCustomConfig
}
func (o *RouterInterfaceGetCustomInfoExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceGetCustomInfoExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceGetCustomInfoExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceGetCustomInfoExResponse structure represents the RRouterInterfaceGetCustomInfoEx operation response
type RouterInterfaceGetCustomInfoExResponse struct {
	InterfaceCustomConfig *rrasm.InterfaceCustominfoexIDL `idl:"name:pIfCustomConfig" json:"interface_custom_config"`
	// Return: The RRouterInterfaceGetCustomInfoEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceGetCustomInfoExResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetCustomInfoExOperation) *xxx_RouterInterfaceGetCustomInfoExOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetCustomInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.InterfaceCustomConfig = o.InterfaceCustomConfig
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceGetCustomInfoExResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetCustomInfoExOperation) {
	if o == nil {
		return
	}
	o.InterfaceCustomConfig = op.InterfaceCustomConfig
	o.Return = op.Return
}
func (o *RouterInterfaceGetCustomInfoExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceGetCustomInfoExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceGetCustomInfoExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceSetCustomInfoExOperation structure represents the RRouterInterfaceSetCustomInfoEx operation
type xxx_RouterInterfaceSetCustomInfoExOperation struct {
	Interface             uint32                          `idl:"name:hInterface" json:"interface"`
	InterfaceCustomConfig *rrasm.InterfaceCustominfoexIDL `idl:"name:pIfCustomConfig" json:"interface_custom_config"`
	Return                uint32                          `idl:"name:Return" json:"return"`
}

func (o *xxx_RouterInterfaceSetCustomInfoExOperation) OpNum() int { return 52 }

func (o *xxx_RouterInterfaceSetCustomInfoExOperation) OpName() string {
	return "/dimsvc/v0/RRouterInterfaceSetCustomInfoEx"
}

func (o *xxx_RouterInterfaceSetCustomInfoExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCustomInfoExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Interface); err != nil {
			return err
		}
	}
	// pIfCustomConfig {in, out} (1:{alias=PMPR_IF_CUSTOMINFOEX_IDL}*(1))(2:{alias=MPR_IF_CUSTOMINFOEX_IDL}(union))
	{
		if o.InterfaceCustomConfig != nil {
			if err := o.InterfaceCustomConfig.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceCustominfoexIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCustomInfoExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Interface); err != nil {
			return err
		}
	}
	// pIfCustomConfig {in, out} (1:{alias=PMPR_IF_CUSTOMINFOEX_IDL,pointer=ref}*(1))(2:{alias=MPR_IF_CUSTOMINFOEX_IDL}(union))
	{
		if o.InterfaceCustomConfig == nil {
			o.InterfaceCustomConfig = &rrasm.InterfaceCustominfoexIDL{}
		}
		if err := o.InterfaceCustomConfig.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCustomInfoExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCustomInfoExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pIfCustomConfig {in, out} (1:{alias=PMPR_IF_CUSTOMINFOEX_IDL}*(1))(2:{alias=MPR_IF_CUSTOMINFOEX_IDL}(union))
	{
		if o.InterfaceCustomConfig != nil {
			if err := o.InterfaceCustomConfig.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.InterfaceCustominfoexIDL{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RouterInterfaceSetCustomInfoExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pIfCustomConfig {in, out} (1:{alias=PMPR_IF_CUSTOMINFOEX_IDL,pointer=ref}*(1))(2:{alias=MPR_IF_CUSTOMINFOEX_IDL}(union))
	{
		if o.InterfaceCustomConfig == nil {
			o.InterfaceCustomConfig = &rrasm.InterfaceCustominfoexIDL{}
		}
		if err := o.InterfaceCustomConfig.UnmarshalNDR(ctx, w); err != nil {
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

// RouterInterfaceSetCustomInfoExRequest structure represents the RRouterInterfaceSetCustomInfoEx operation request
type RouterInterfaceSetCustomInfoExRequest struct {
	Interface             uint32                          `idl:"name:hInterface" json:"interface"`
	InterfaceCustomConfig *rrasm.InterfaceCustominfoexIDL `idl:"name:pIfCustomConfig" json:"interface_custom_config"`
}

func (o *RouterInterfaceSetCustomInfoExRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceSetCustomInfoExOperation) *xxx_RouterInterfaceSetCustomInfoExOperation {
	if op == nil {
		op = &xxx_RouterInterfaceSetCustomInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.Interface = o.Interface
	op.InterfaceCustomConfig = o.InterfaceCustomConfig
	return op
}

func (o *RouterInterfaceSetCustomInfoExRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceSetCustomInfoExOperation) {
	if o == nil {
		return
	}
	o.Interface = op.Interface
	o.InterfaceCustomConfig = op.InterfaceCustomConfig
}
func (o *RouterInterfaceSetCustomInfoExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RouterInterfaceSetCustomInfoExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceSetCustomInfoExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RouterInterfaceSetCustomInfoExResponse structure represents the RRouterInterfaceSetCustomInfoEx operation response
type RouterInterfaceSetCustomInfoExResponse struct {
	InterfaceCustomConfig *rrasm.InterfaceCustominfoexIDL `idl:"name:pIfCustomConfig" json:"interface_custom_config"`
	// Return: The RRouterInterfaceSetCustomInfoEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RouterInterfaceSetCustomInfoExResponse) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceSetCustomInfoExOperation) *xxx_RouterInterfaceSetCustomInfoExOperation {
	if op == nil {
		op = &xxx_RouterInterfaceSetCustomInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.InterfaceCustomConfig = o.InterfaceCustomConfig
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceSetCustomInfoExResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceSetCustomInfoExOperation) {
	if o == nil {
		return
	}
	o.InterfaceCustomConfig = op.InterfaceCustomConfig
	o.Return = op.Return
}
func (o *RouterInterfaceSetCustomInfoExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RouterInterfaceSetCustomInfoExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RouterInterfaceSetCustomInfoExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
