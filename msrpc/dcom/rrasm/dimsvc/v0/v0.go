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
	MprAdminServerGetInfo(context.Context, *MprAdminServerGetInfoRequest, ...dcerpc.CallOption) (*MprAdminServerGetInfoResponse, error)

	RASAdminConnectionEnum(context.Context, *RASAdminConnectionEnumRequest, ...dcerpc.CallOption) (*RASAdminConnectionEnumResponse, error)

	RASAdminConnectionGetInfo(context.Context, *RASAdminConnectionGetInfoRequest, ...dcerpc.CallOption) (*RASAdminConnectionGetInfoResponse, error)

	RASAdminConnectionClearStats(context.Context, *RASAdminConnectionClearStatsRequest, ...dcerpc.CallOption) (*RASAdminConnectionClearStatsResponse, error)

	RASAdminPortEnum(context.Context, *RASAdminPortEnumRequest, ...dcerpc.CallOption) (*RASAdminPortEnumResponse, error)

	RASAdminPortGetInfo(context.Context, *RASAdminPortGetInfoRequest, ...dcerpc.CallOption) (*RASAdminPortGetInfoResponse, error)

	RASAdminPortClearStats(context.Context, *RASAdminPortClearStatsRequest, ...dcerpc.CallOption) (*RASAdminPortClearStatsResponse, error)

	RASAdminPortReset(context.Context, *RASAdminPortResetRequest, ...dcerpc.CallOption) (*RASAdminPortResetResponse, error)

	RASAdminPortDisconnect(context.Context, *RASAdminPortDisconnectRequest, ...dcerpc.CallOption) (*RASAdminPortDisconnectResponse, error)

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

	RmibEntryCreate(context.Context, *RmibEntryCreateRequest, ...dcerpc.CallOption) (*RmibEntryCreateResponse, error)

	RmibEntryDelete(context.Context, *RmibEntryDeleteRequest, ...dcerpc.CallOption) (*RmibEntryDeleteResponse, error)

	RmibEntrySet(context.Context, *RmibEntrySetRequest, ...dcerpc.CallOption) (*RmibEntrySetResponse, error)

	RmibEntryGet(context.Context, *RmibEntryGetRequest, ...dcerpc.CallOption) (*RmibEntryGetResponse, error)

	RmibEntryGetFirst(context.Context, *RmibEntryGetFirstRequest, ...dcerpc.CallOption) (*RmibEntryGetFirstResponse, error)

	RmibEntryGetNext(context.Context, *RmibEntryGetNextRequest, ...dcerpc.CallOption) (*RmibEntryGetNextResponse, error)

	RmibGetTrapInfo(context.Context, *RmibGetTrapInfoRequest, ...dcerpc.CallOption) (*RmibGetTrapInfoResponse, error)

	RmibSetTrapInfo(context.Context, *RmibSetTrapInfoRequest, ...dcerpc.CallOption) (*RmibSetTrapInfoResponse, error)

	RASAdminConnectionNotification(context.Context, *RASAdminConnectionNotificationRequest, ...dcerpc.CallOption) (*RASAdminConnectionNotificationResponse, error)

	RASAdminSendUserMessage(context.Context, *RASAdminSendUserMessageRequest, ...dcerpc.CallOption) (*RASAdminSendUserMessageResponse, error)

	RouterDeviceEnum(context.Context, *RouterDeviceEnumRequest, ...dcerpc.CallOption) (*RouterDeviceEnumResponse, error)

	RouterInterfaceTransportCreate(context.Context, *RouterInterfaceTransportCreateRequest, ...dcerpc.CallOption) (*RouterInterfaceTransportCreateResponse, error)

	RouterInterfaceDeviceGetInfo(context.Context, *RouterInterfaceDeviceGetInfoRequest, ...dcerpc.CallOption) (*RouterInterfaceDeviceGetInfoResponse, error)

	RouterInterfaceDeviceSetInfo(context.Context, *RouterInterfaceDeviceSetInfoRequest, ...dcerpc.CallOption) (*RouterInterfaceDeviceSetInfoResponse, error)

	RouterInterfaceSetCredentialsEx(context.Context, *RouterInterfaceSetCredentialsExRequest, ...dcerpc.CallOption) (*RouterInterfaceSetCredentialsExResponse, error)

	RouterInterfaceGetCredentialsEx(context.Context, *RouterInterfaceGetCredentialsExRequest, ...dcerpc.CallOption) (*RouterInterfaceGetCredentialsExResponse, error)

	RASAdminConnectionRemoveQuarantine(context.Context, *RASAdminConnectionRemoveQuarantineRequest, ...dcerpc.CallOption) (*RASAdminConnectionRemoveQuarantineResponse, error)

	MprAdminServerSetInfo(context.Context, *MprAdminServerSetInfoRequest, ...dcerpc.CallOption) (*MprAdminServerSetInfoResponse, error)

	MprAdminServerGetInfoEx(context.Context, *MprAdminServerGetInfoExRequest, ...dcerpc.CallOption) (*MprAdminServerGetInfoExResponse, error)

	RASAdminConnectionEnumEx(context.Context, *RASAdminConnectionEnumExRequest, ...dcerpc.CallOption) (*RASAdminConnectionEnumExResponse, error)

	RASAdminConnectionGetInfoEx(context.Context, *RASAdminConnectionGetInfoExRequest, ...dcerpc.CallOption) (*RASAdminConnectionGetInfoExResponse, error)

	MprAdminServerSetInfoEx(context.Context, *MprAdminServerSetInfoExRequest, ...dcerpc.CallOption) (*MprAdminServerSetInfoExResponse, error)

	RASAdminUpdateConnection(context.Context, *RASAdminUpdateConnectionRequest, ...dcerpc.CallOption) (*RASAdminUpdateConnectionResponse, error)

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

func (o *xxx_DefaultDimsvcClient) MprAdminServerGetInfo(ctx context.Context, in *MprAdminServerGetInfoRequest, opts ...dcerpc.CallOption) (*MprAdminServerGetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MprAdminServerGetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RASAdminConnectionEnum(ctx context.Context, in *RASAdminConnectionEnumRequest, opts ...dcerpc.CallOption) (*RASAdminConnectionEnumResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASAdminConnectionEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RASAdminConnectionGetInfo(ctx context.Context, in *RASAdminConnectionGetInfoRequest, opts ...dcerpc.CallOption) (*RASAdminConnectionGetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASAdminConnectionGetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RASAdminConnectionClearStats(ctx context.Context, in *RASAdminConnectionClearStatsRequest, opts ...dcerpc.CallOption) (*RASAdminConnectionClearStatsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASAdminConnectionClearStatsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RASAdminPortEnum(ctx context.Context, in *RASAdminPortEnumRequest, opts ...dcerpc.CallOption) (*RASAdminPortEnumResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASAdminPortEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RASAdminPortGetInfo(ctx context.Context, in *RASAdminPortGetInfoRequest, opts ...dcerpc.CallOption) (*RASAdminPortGetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASAdminPortGetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RASAdminPortClearStats(ctx context.Context, in *RASAdminPortClearStatsRequest, opts ...dcerpc.CallOption) (*RASAdminPortClearStatsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASAdminPortClearStatsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RASAdminPortReset(ctx context.Context, in *RASAdminPortResetRequest, opts ...dcerpc.CallOption) (*RASAdminPortResetResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASAdminPortResetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RASAdminPortDisconnect(ctx context.Context, in *RASAdminPortDisconnectRequest, opts ...dcerpc.CallOption) (*RASAdminPortDisconnectResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASAdminPortDisconnectResponse{}
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

func (o *xxx_DefaultDimsvcClient) RmibEntryCreate(ctx context.Context, in *RmibEntryCreateRequest, opts ...dcerpc.CallOption) (*RmibEntryCreateResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RmibEntryCreateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RmibEntryDelete(ctx context.Context, in *RmibEntryDeleteRequest, opts ...dcerpc.CallOption) (*RmibEntryDeleteResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RmibEntryDeleteResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RmibEntrySet(ctx context.Context, in *RmibEntrySetRequest, opts ...dcerpc.CallOption) (*RmibEntrySetResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RmibEntrySetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RmibEntryGet(ctx context.Context, in *RmibEntryGetRequest, opts ...dcerpc.CallOption) (*RmibEntryGetResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RmibEntryGetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RmibEntryGetFirst(ctx context.Context, in *RmibEntryGetFirstRequest, opts ...dcerpc.CallOption) (*RmibEntryGetFirstResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RmibEntryGetFirstResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RmibEntryGetNext(ctx context.Context, in *RmibEntryGetNextRequest, opts ...dcerpc.CallOption) (*RmibEntryGetNextResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RmibEntryGetNextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RmibGetTrapInfo(ctx context.Context, in *RmibGetTrapInfoRequest, opts ...dcerpc.CallOption) (*RmibGetTrapInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RmibGetTrapInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RmibSetTrapInfo(ctx context.Context, in *RmibSetTrapInfoRequest, opts ...dcerpc.CallOption) (*RmibSetTrapInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RmibSetTrapInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RASAdminConnectionNotification(ctx context.Context, in *RASAdminConnectionNotificationRequest, opts ...dcerpc.CallOption) (*RASAdminConnectionNotificationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASAdminConnectionNotificationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RASAdminSendUserMessage(ctx context.Context, in *RASAdminSendUserMessageRequest, opts ...dcerpc.CallOption) (*RASAdminSendUserMessageResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASAdminSendUserMessageResponse{}
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

func (o *xxx_DefaultDimsvcClient) RASAdminConnectionRemoveQuarantine(ctx context.Context, in *RASAdminConnectionRemoveQuarantineRequest, opts ...dcerpc.CallOption) (*RASAdminConnectionRemoveQuarantineResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASAdminConnectionRemoveQuarantineResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) MprAdminServerSetInfo(ctx context.Context, in *MprAdminServerSetInfoRequest, opts ...dcerpc.CallOption) (*MprAdminServerSetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MprAdminServerSetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) MprAdminServerGetInfoEx(ctx context.Context, in *MprAdminServerGetInfoExRequest, opts ...dcerpc.CallOption) (*MprAdminServerGetInfoExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MprAdminServerGetInfoExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RASAdminConnectionEnumEx(ctx context.Context, in *RASAdminConnectionEnumExRequest, opts ...dcerpc.CallOption) (*RASAdminConnectionEnumExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASAdminConnectionEnumExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RASAdminConnectionGetInfoEx(ctx context.Context, in *RASAdminConnectionGetInfoExRequest, opts ...dcerpc.CallOption) (*RASAdminConnectionGetInfoExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASAdminConnectionGetInfoExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) MprAdminServerSetInfoEx(ctx context.Context, in *MprAdminServerSetInfoExRequest, opts ...dcerpc.CallOption) (*MprAdminServerSetInfoExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MprAdminServerSetInfoExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultDimsvcClient) RASAdminUpdateConnection(ctx context.Context, in *RASAdminUpdateConnectionRequest, opts ...dcerpc.CallOption) (*RASAdminUpdateConnectionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RASAdminUpdateConnectionResponse{}
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

// xxx_MprAdminServerGetInfoOperation structure represents the RMprAdminServerGetInfo operation
type xxx_MprAdminServerGetInfoOperation struct {
	Level  uint32                         `idl:"name:dwLevel" json:"level"`
	Info   *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	Return uint32                         `idl:"name:Return" json:"return"`
}

func (o *xxx_MprAdminServerGetInfoOperation) OpNum() int { return 0 }

func (o *xxx_MprAdminServerGetInfoOperation) OpName() string {
	return "/dimsvc/v0/RMprAdminServerGetInfo"
}

func (o *xxx_MprAdminServerGetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MprAdminServerGetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_MprAdminServerGetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MprAdminServerGetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MprAdminServerGetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_MprAdminServerGetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.DimInformationContainer{}
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

// MprAdminServerGetInfoRequest structure represents the RMprAdminServerGetInfo operation request
type MprAdminServerGetInfoRequest struct {
	Level uint32 `idl:"name:dwLevel" json:"level"`
}

func (o *MprAdminServerGetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_MprAdminServerGetInfoOperation) *xxx_MprAdminServerGetInfoOperation {
	if op == nil {
		op = &xxx_MprAdminServerGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	return op
}

func (o *MprAdminServerGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_MprAdminServerGetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
}
func (o *MprAdminServerGetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MprAdminServerGetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MprAdminServerGetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MprAdminServerGetInfoResponse structure represents the RMprAdminServerGetInfo operation response
type MprAdminServerGetInfoResponse struct {
	Info *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RMprAdminServerGetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MprAdminServerGetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_MprAdminServerGetInfoOperation) *xxx_MprAdminServerGetInfoOperation {
	if op == nil {
		op = &xxx_MprAdminServerGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *MprAdminServerGetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_MprAdminServerGetInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *MprAdminServerGetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MprAdminServerGetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MprAdminServerGetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RASAdminConnectionEnumOperation structure represents the RRasAdminConnectionEnum operation
type xxx_RASAdminConnectionEnumOperation struct {
	Level                  uint32                         `idl:"name:dwLevel" json:"level"`
	Info                   *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	PreferredMaximumLength uint32                         `idl:"name:dwPreferedMaximumLength" json:"preferred_maximum_length"`
	EntriesRead            uint32                         `idl:"name:lpdwEntriesRead" json:"entries_read"`
	TotalEntries           uint32                         `idl:"name:lpdwTotalEntries" json:"total_entries"`
	Resume                 uint32                         `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	Return                 uint32                         `idl:"name:Return" json:"return"`
}

func (o *xxx_RASAdminConnectionEnumOperation) OpNum() int { return 1 }

func (o *xxx_RASAdminConnectionEnumOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminConnectionEnum"
}

func (o *xxx_RASAdminConnectionEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RASAdminConnectionEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.DimInformationContainer{}
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

func (o *xxx_RASAdminConnectionEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RASAdminConnectionEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.DimInformationContainer{}
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

// RASAdminConnectionEnumRequest structure represents the RRasAdminConnectionEnum operation request
type RASAdminConnectionEnumRequest struct {
	Level                  uint32                         `idl:"name:dwLevel" json:"level"`
	Info                   *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	PreferredMaximumLength uint32                         `idl:"name:dwPreferedMaximumLength" json:"preferred_maximum_length"`
	Resume                 uint32                         `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
}

func (o *RASAdminConnectionEnumRequest) xxx_ToOp(ctx context.Context, op *xxx_RASAdminConnectionEnumOperation) *xxx_RASAdminConnectionEnumOperation {
	if op == nil {
		op = &xxx_RASAdminConnectionEnumOperation{}
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

func (o *RASAdminConnectionEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_RASAdminConnectionEnumOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.PreferredMaximumLength = op.PreferredMaximumLength
	o.Resume = op.Resume
}
func (o *RASAdminConnectionEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASAdminConnectionEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminConnectionEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASAdminConnectionEnumResponse structure represents the RRasAdminConnectionEnum operation response
type RASAdminConnectionEnumResponse struct {
	Info         *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	EntriesRead  uint32                         `idl:"name:lpdwEntriesRead" json:"entries_read"`
	TotalEntries uint32                         `idl:"name:lpdwTotalEntries" json:"total_entries"`
	Resume       uint32                         `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	// Return: The RRasAdminConnectionEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASAdminConnectionEnumResponse) xxx_ToOp(ctx context.Context, op *xxx_RASAdminConnectionEnumOperation) *xxx_RASAdminConnectionEnumOperation {
	if op == nil {
		op = &xxx_RASAdminConnectionEnumOperation{}
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

func (o *RASAdminConnectionEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_RASAdminConnectionEnumOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.EntriesRead = op.EntriesRead
	o.TotalEntries = op.TotalEntries
	o.Resume = op.Resume
	o.Return = op.Return
}
func (o *RASAdminConnectionEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASAdminConnectionEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminConnectionEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RASAdminConnectionGetInfoOperation structure represents the RRasAdminConnectionGetInfo operation
type xxx_RASAdminConnectionGetInfoOperation struct {
	Level          uint32                         `idl:"name:dwLevel" json:"level"`
	HDimConnection uint32                         `idl:"name:hDimConnection" json:"h_dim_connection"`
	Info           *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	Return         uint32                         `idl:"name:Return" json:"return"`
}

func (o *xxx_RASAdminConnectionGetInfoOperation) OpNum() int { return 2 }

func (o *xxx_RASAdminConnectionGetInfoOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminConnectionGetInfo"
}

func (o *xxx_RASAdminConnectionGetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionGetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
		if err := w.WriteData(o.HDimConnection); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionGetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.HDimConnection); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionGetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionGetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RASAdminConnectionGetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.DimInformationContainer{}
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

// RASAdminConnectionGetInfoRequest structure represents the RRasAdminConnectionGetInfo operation request
type RASAdminConnectionGetInfoRequest struct {
	Level          uint32 `idl:"name:dwLevel" json:"level"`
	HDimConnection uint32 `idl:"name:hDimConnection" json:"h_dim_connection"`
}

func (o *RASAdminConnectionGetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RASAdminConnectionGetInfoOperation) *xxx_RASAdminConnectionGetInfoOperation {
	if op == nil {
		op = &xxx_RASAdminConnectionGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.HDimConnection = o.HDimConnection
	return op
}

func (o *RASAdminConnectionGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RASAdminConnectionGetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.HDimConnection = op.HDimConnection
}
func (o *RASAdminConnectionGetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASAdminConnectionGetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminConnectionGetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASAdminConnectionGetInfoResponse structure represents the RRasAdminConnectionGetInfo operation response
type RASAdminConnectionGetInfoResponse struct {
	Info *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RRasAdminConnectionGetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASAdminConnectionGetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RASAdminConnectionGetInfoOperation) *xxx_RASAdminConnectionGetInfoOperation {
	if op == nil {
		op = &xxx_RASAdminConnectionGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *RASAdminConnectionGetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RASAdminConnectionGetInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *RASAdminConnectionGetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASAdminConnectionGetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminConnectionGetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RASAdminConnectionClearStatsOperation structure represents the RRasAdminConnectionClearStats operation
type xxx_RASAdminConnectionClearStatsOperation struct {
	HDimConnection uint32 `idl:"name:hDimConnection" json:"h_dim_connection"`
	Return         uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RASAdminConnectionClearStatsOperation) OpNum() int { return 3 }

func (o *xxx_RASAdminConnectionClearStatsOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminConnectionClearStats"
}

func (o *xxx_RASAdminConnectionClearStatsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionClearStatsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.HDimConnection); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionClearStatsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.HDimConnection); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionClearStatsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionClearStatsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RASAdminConnectionClearStatsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RASAdminConnectionClearStatsRequest structure represents the RRasAdminConnectionClearStats operation request
type RASAdminConnectionClearStatsRequest struct {
	HDimConnection uint32 `idl:"name:hDimConnection" json:"h_dim_connection"`
}

func (o *RASAdminConnectionClearStatsRequest) xxx_ToOp(ctx context.Context, op *xxx_RASAdminConnectionClearStatsOperation) *xxx_RASAdminConnectionClearStatsOperation {
	if op == nil {
		op = &xxx_RASAdminConnectionClearStatsOperation{}
	}
	if o == nil {
		return op
	}
	op.HDimConnection = o.HDimConnection
	return op
}

func (o *RASAdminConnectionClearStatsRequest) xxx_FromOp(ctx context.Context, op *xxx_RASAdminConnectionClearStatsOperation) {
	if o == nil {
		return
	}
	o.HDimConnection = op.HDimConnection
}
func (o *RASAdminConnectionClearStatsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASAdminConnectionClearStatsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminConnectionClearStatsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASAdminConnectionClearStatsResponse structure represents the RRasAdminConnectionClearStats operation response
type RASAdminConnectionClearStatsResponse struct {
	// Return: The RRasAdminConnectionClearStats return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASAdminConnectionClearStatsResponse) xxx_ToOp(ctx context.Context, op *xxx_RASAdminConnectionClearStatsOperation) *xxx_RASAdminConnectionClearStatsOperation {
	if op == nil {
		op = &xxx_RASAdminConnectionClearStatsOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RASAdminConnectionClearStatsResponse) xxx_FromOp(ctx context.Context, op *xxx_RASAdminConnectionClearStatsOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RASAdminConnectionClearStatsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASAdminConnectionClearStatsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminConnectionClearStatsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RASAdminPortEnumOperation structure represents the RRasAdminPortEnum operation
type xxx_RASAdminPortEnumOperation struct {
	Level                  uint32                         `idl:"name:dwLevel" json:"level"`
	HRASConnection         uint32                         `idl:"name:hRasConnection" json:"h_ras_connection"`
	Info                   *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	PreferredMaximumLength uint32                         `idl:"name:dwPreferedMaximumLength" json:"preferred_maximum_length"`
	EntriesRead            uint32                         `idl:"name:lpdwEntriesRead" json:"entries_read"`
	TotalEntries           uint32                         `idl:"name:lpdwTotalEntries" json:"total_entries"`
	Resume                 uint32                         `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	Return                 uint32                         `idl:"name:Return" json:"return"`
}

func (o *xxx_RASAdminPortEnumOperation) OpNum() int { return 4 }

func (o *xxx_RASAdminPortEnumOperation) OpName() string { return "/dimsvc/v0/RRasAdminPortEnum" }

func (o *xxx_RASAdminPortEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminPortEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
		if err := w.WriteData(o.HRASConnection); err != nil {
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RASAdminPortEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// hRasConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.HRASConnection); err != nil {
			return err
		}
	}
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.DimInformationContainer{}
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

func (o *xxx_RASAdminPortEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminPortEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RASAdminPortEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.DimInformationContainer{}
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

// RASAdminPortEnumRequest structure represents the RRasAdminPortEnum operation request
type RASAdminPortEnumRequest struct {
	Level                  uint32                         `idl:"name:dwLevel" json:"level"`
	HRASConnection         uint32                         `idl:"name:hRasConnection" json:"h_ras_connection"`
	Info                   *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	PreferredMaximumLength uint32                         `idl:"name:dwPreferedMaximumLength" json:"preferred_maximum_length"`
	Resume                 uint32                         `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
}

func (o *RASAdminPortEnumRequest) xxx_ToOp(ctx context.Context, op *xxx_RASAdminPortEnumOperation) *xxx_RASAdminPortEnumOperation {
	if op == nil {
		op = &xxx_RASAdminPortEnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.HRASConnection = o.HRASConnection
	op.Info = o.Info
	op.PreferredMaximumLength = o.PreferredMaximumLength
	op.Resume = o.Resume
	return op
}

func (o *RASAdminPortEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_RASAdminPortEnumOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.HRASConnection = op.HRASConnection
	o.Info = op.Info
	o.PreferredMaximumLength = op.PreferredMaximumLength
	o.Resume = op.Resume
}
func (o *RASAdminPortEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASAdminPortEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminPortEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASAdminPortEnumResponse structure represents the RRasAdminPortEnum operation response
type RASAdminPortEnumResponse struct {
	Info         *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	EntriesRead  uint32                         `idl:"name:lpdwEntriesRead" json:"entries_read"`
	TotalEntries uint32                         `idl:"name:lpdwTotalEntries" json:"total_entries"`
	Resume       uint32                         `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	// Return: The RRasAdminPortEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASAdminPortEnumResponse) xxx_ToOp(ctx context.Context, op *xxx_RASAdminPortEnumOperation) *xxx_RASAdminPortEnumOperation {
	if op == nil {
		op = &xxx_RASAdminPortEnumOperation{}
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

func (o *RASAdminPortEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_RASAdminPortEnumOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.EntriesRead = op.EntriesRead
	o.TotalEntries = op.TotalEntries
	o.Resume = op.Resume
	o.Return = op.Return
}
func (o *RASAdminPortEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASAdminPortEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminPortEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RASAdminPortGetInfoOperation structure represents the RRasAdminPortGetInfo operation
type xxx_RASAdminPortGetInfoOperation struct {
	Level  uint32                         `idl:"name:dwLevel" json:"level"`
	HPort  uint32                         `idl:"name:hPort" json:"h_port"`
	Info   *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	Return uint32                         `idl:"name:Return" json:"return"`
}

func (o *xxx_RASAdminPortGetInfoOperation) OpNum() int { return 5 }

func (o *xxx_RASAdminPortGetInfoOperation) OpName() string { return "/dimsvc/v0/RRasAdminPortGetInfo" }

func (o *xxx_RASAdminPortGetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminPortGetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
		if err := w.WriteData(o.HPort); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminPortGetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.HPort); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminPortGetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminPortGetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RASAdminPortGetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {out} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.DimInformationContainer{}
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

// RASAdminPortGetInfoRequest structure represents the RRasAdminPortGetInfo operation request
type RASAdminPortGetInfoRequest struct {
	Level uint32 `idl:"name:dwLevel" json:"level"`
	HPort uint32 `idl:"name:hPort" json:"h_port"`
}

func (o *RASAdminPortGetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RASAdminPortGetInfoOperation) *xxx_RASAdminPortGetInfoOperation {
	if op == nil {
		op = &xxx_RASAdminPortGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.HPort = o.HPort
	return op
}

func (o *RASAdminPortGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RASAdminPortGetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.HPort = op.HPort
}
func (o *RASAdminPortGetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASAdminPortGetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminPortGetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASAdminPortGetInfoResponse structure represents the RRasAdminPortGetInfo operation response
type RASAdminPortGetInfoResponse struct {
	Info *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RRasAdminPortGetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASAdminPortGetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RASAdminPortGetInfoOperation) *xxx_RASAdminPortGetInfoOperation {
	if op == nil {
		op = &xxx_RASAdminPortGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *RASAdminPortGetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RASAdminPortGetInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *RASAdminPortGetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASAdminPortGetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminPortGetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RASAdminPortClearStatsOperation structure represents the RRasAdminPortClearStats operation
type xxx_RASAdminPortClearStatsOperation struct {
	HPort  uint32 `idl:"name:hPort" json:"h_port"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RASAdminPortClearStatsOperation) OpNum() int { return 6 }

func (o *xxx_RASAdminPortClearStatsOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminPortClearStats"
}

func (o *xxx_RASAdminPortClearStatsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminPortClearStatsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.HPort); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminPortClearStatsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.HPort); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminPortClearStatsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminPortClearStatsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RASAdminPortClearStatsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RASAdminPortClearStatsRequest structure represents the RRasAdminPortClearStats operation request
type RASAdminPortClearStatsRequest struct {
	HPort uint32 `idl:"name:hPort" json:"h_port"`
}

func (o *RASAdminPortClearStatsRequest) xxx_ToOp(ctx context.Context, op *xxx_RASAdminPortClearStatsOperation) *xxx_RASAdminPortClearStatsOperation {
	if op == nil {
		op = &xxx_RASAdminPortClearStatsOperation{}
	}
	if o == nil {
		return op
	}
	op.HPort = o.HPort
	return op
}

func (o *RASAdminPortClearStatsRequest) xxx_FromOp(ctx context.Context, op *xxx_RASAdminPortClearStatsOperation) {
	if o == nil {
		return
	}
	o.HPort = op.HPort
}
func (o *RASAdminPortClearStatsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASAdminPortClearStatsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminPortClearStatsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASAdminPortClearStatsResponse structure represents the RRasAdminPortClearStats operation response
type RASAdminPortClearStatsResponse struct {
	// Return: The RRasAdminPortClearStats return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASAdminPortClearStatsResponse) xxx_ToOp(ctx context.Context, op *xxx_RASAdminPortClearStatsOperation) *xxx_RASAdminPortClearStatsOperation {
	if op == nil {
		op = &xxx_RASAdminPortClearStatsOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RASAdminPortClearStatsResponse) xxx_FromOp(ctx context.Context, op *xxx_RASAdminPortClearStatsOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RASAdminPortClearStatsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASAdminPortClearStatsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminPortClearStatsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RASAdminPortResetOperation structure represents the RRasAdminPortReset operation
type xxx_RASAdminPortResetOperation struct {
	HPort  uint32 `idl:"name:hPort" json:"h_port"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RASAdminPortResetOperation) OpNum() int { return 7 }

func (o *xxx_RASAdminPortResetOperation) OpName() string { return "/dimsvc/v0/RRasAdminPortReset" }

func (o *xxx_RASAdminPortResetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminPortResetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.HPort); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminPortResetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.HPort); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminPortResetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminPortResetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RASAdminPortResetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RASAdminPortResetRequest structure represents the RRasAdminPortReset operation request
type RASAdminPortResetRequest struct {
	HPort uint32 `idl:"name:hPort" json:"h_port"`
}

func (o *RASAdminPortResetRequest) xxx_ToOp(ctx context.Context, op *xxx_RASAdminPortResetOperation) *xxx_RASAdminPortResetOperation {
	if op == nil {
		op = &xxx_RASAdminPortResetOperation{}
	}
	if o == nil {
		return op
	}
	op.HPort = o.HPort
	return op
}

func (o *RASAdminPortResetRequest) xxx_FromOp(ctx context.Context, op *xxx_RASAdminPortResetOperation) {
	if o == nil {
		return
	}
	o.HPort = op.HPort
}
func (o *RASAdminPortResetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASAdminPortResetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminPortResetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASAdminPortResetResponse structure represents the RRasAdminPortReset operation response
type RASAdminPortResetResponse struct {
	// Return: The RRasAdminPortReset return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASAdminPortResetResponse) xxx_ToOp(ctx context.Context, op *xxx_RASAdminPortResetOperation) *xxx_RASAdminPortResetOperation {
	if op == nil {
		op = &xxx_RASAdminPortResetOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RASAdminPortResetResponse) xxx_FromOp(ctx context.Context, op *xxx_RASAdminPortResetOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RASAdminPortResetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASAdminPortResetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminPortResetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RASAdminPortDisconnectOperation structure represents the RRasAdminPortDisconnect operation
type xxx_RASAdminPortDisconnectOperation struct {
	HPort  uint32 `idl:"name:hPort" json:"h_port"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RASAdminPortDisconnectOperation) OpNum() int { return 8 }

func (o *xxx_RASAdminPortDisconnectOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminPortDisconnect"
}

func (o *xxx_RASAdminPortDisconnectOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminPortDisconnectOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.HPort); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminPortDisconnectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hPort {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.HPort); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminPortDisconnectOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminPortDisconnectOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RASAdminPortDisconnectOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RASAdminPortDisconnectRequest structure represents the RRasAdminPortDisconnect operation request
type RASAdminPortDisconnectRequest struct {
	HPort uint32 `idl:"name:hPort" json:"h_port"`
}

func (o *RASAdminPortDisconnectRequest) xxx_ToOp(ctx context.Context, op *xxx_RASAdminPortDisconnectOperation) *xxx_RASAdminPortDisconnectOperation {
	if op == nil {
		op = &xxx_RASAdminPortDisconnectOperation{}
	}
	if o == nil {
		return op
	}
	op.HPort = o.HPort
	return op
}

func (o *RASAdminPortDisconnectRequest) xxx_FromOp(ctx context.Context, op *xxx_RASAdminPortDisconnectOperation) {
	if o == nil {
		return
	}
	o.HPort = op.HPort
}
func (o *RASAdminPortDisconnectRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASAdminPortDisconnectRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminPortDisconnectOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASAdminPortDisconnectResponse structure represents the RRasAdminPortDisconnect operation response
type RASAdminPortDisconnectResponse struct {
	// Return: The RRasAdminPortDisconnect return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASAdminPortDisconnectResponse) xxx_ToOp(ctx context.Context, op *xxx_RASAdminPortDisconnectOperation) *xxx_RASAdminPortDisconnectOperation {
	if op == nil {
		op = &xxx_RASAdminPortDisconnectOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RASAdminPortDisconnectResponse) xxx_FromOp(ctx context.Context, op *xxx_RASAdminPortDisconnectOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RASAdminPortDisconnectResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASAdminPortDisconnectResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminPortDisconnectOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceTransportSetGlobalInfoOperation structure represents the RRouterInterfaceTransportSetGlobalInfo operation
type xxx_RouterInterfaceTransportSetGlobalInfoOperation struct {
	TransportID uint32                       `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.DimInterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	Return      uint32                       `idl:"name:Return" json:"return"`
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
			if err := (&rrasm.DimInterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
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
			o.Info = &rrasm.DimInterfaceContainer{}
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
	TransportID uint32                       `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.DimInterfaceContainer `idl:"name:pInfoStruct" json:"info"`
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
	TransportID uint32                       `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.DimInterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	Return      uint32                       `idl:"name:Return" json:"return"`
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
			if err := (&rrasm.DimInterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
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
			o.Info = &rrasm.DimInterfaceContainer{}
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
			if err := (&rrasm.DimInterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
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
			o.Info = &rrasm.DimInterfaceContainer{}
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
	TransportID uint32                       `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.DimInterfaceContainer `idl:"name:pInfoStruct" json:"info"`
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
	Info *rrasm.DimInterfaceContainer `idl:"name:pInfoStruct" json:"info"`
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
	LpwsInterfaceName       string `idl:"name:lpwsInterfaceName;string" json:"lpws_interface_name"`
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
		if err := ndr.WriteUTF16NString(ctx, w, o.LpwsInterfaceName); err != nil {
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
		if err := ndr.ReadUTF16NString(ctx, w, &o.LpwsInterfaceName); err != nil {
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
	LpwsInterfaceName       string `idl:"name:lpwsInterfaceName;string" json:"lpws_interface_name"`
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
	op.LpwsInterfaceName = o.LpwsInterfaceName
	op.Interface = o.Interface
	op.IncludeClientInterfaces = o.IncludeClientInterfaces
	return op
}

func (o *RouterInterfaceGetHandleRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetHandleOperation) {
	if o == nil {
		return
	}
	o.LpwsInterfaceName = op.LpwsInterfaceName
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
	Level     uint32                         `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	Interface uint32                         `idl:"name:phInterface" json:"interface"`
	Return    uint32                         `idl:"name:Return" json:"return"`
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
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
			o.Info = &rrasm.DimInformationContainer{}
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
	Level     uint32                         `idl:"name:dwLevel" json:"level"`
	Info      *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	Interface uint32                         `idl:"name:phInterface" json:"interface"`
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
	Level      uint32                         `idl:"name:dwLevel" json:"level"`
	Info       *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	HInterface uint32                         `idl:"name:hInterface" json:"h_interface"`
	Return     uint32                         `idl:"name:Return" json:"return"`
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.HInterface); err != nil {
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
			o.Info = &rrasm.DimInformationContainer{}
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
		if err := w.ReadData(&o.HInterface); err != nil {
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
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
			o.Info = &rrasm.DimInformationContainer{}
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
	Level      uint32                         `idl:"name:dwLevel" json:"level"`
	Info       *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	HInterface uint32                         `idl:"name:hInterface" json:"h_interface"`
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
	op.HInterface = o.HInterface
	return op
}

func (o *RouterInterfaceGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.HInterface = op.HInterface
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
	Info *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
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
	Level      uint32                         `idl:"name:dwLevel" json:"level"`
	Info       *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	HInterface uint32                         `idl:"name:hInterface" json:"h_interface"`
	Return     uint32                         `idl:"name:Return" json:"return"`
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.HInterface); err != nil {
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
			o.Info = &rrasm.DimInformationContainer{}
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
		if err := w.ReadData(&o.HInterface); err != nil {
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
	Level      uint32                         `idl:"name:dwLevel" json:"level"`
	Info       *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	HInterface uint32                         `idl:"name:hInterface" json:"h_interface"`
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
	op.HInterface = o.HInterface
	return op
}

func (o *RouterInterfaceSetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceSetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.HInterface = op.HInterface
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
	HInterface uint32 `idl:"name:hInterface" json:"h_interface"`
	Return     uint32 `idl:"name:Return" json:"return"`
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
		if err := w.WriteData(o.HInterface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.HInterface); err != nil {
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
	HInterface uint32 `idl:"name:hInterface" json:"h_interface"`
}

func (o *RouterInterfaceDeleteRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceDeleteOperation) *xxx_RouterInterfaceDeleteOperation {
	if op == nil {
		op = &xxx_RouterInterfaceDeleteOperation{}
	}
	if o == nil {
		return op
	}
	op.HInterface = o.HInterface
	return op
}

func (o *RouterInterfaceDeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceDeleteOperation) {
	if o == nil {
		return
	}
	o.HInterface = op.HInterface
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
	HInterface  uint32 `idl:"name:hInterface" json:"h_interface"`
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
		if err := w.WriteData(o.HInterface); err != nil {
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
		if err := w.ReadData(&o.HInterface); err != nil {
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
	HInterface  uint32 `idl:"name:hInterface" json:"h_interface"`
	TransportID uint32 `idl:"name:dwTransportId" json:"transport_id"`
}

func (o *RouterInterfaceTransportRemoveRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportRemoveOperation) *xxx_RouterInterfaceTransportRemoveOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportRemoveOperation{}
	}
	if o == nil {
		return op
	}
	op.HInterface = o.HInterface
	op.TransportID = o.TransportID
	return op
}

func (o *RouterInterfaceTransportRemoveRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportRemoveOperation) {
	if o == nil {
		return
	}
	o.HInterface = op.HInterface
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
	HInterface  uint32                       `idl:"name:hInterface" json:"h_interface"`
	TransportID uint32                       `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.DimInterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	Return      uint32                       `idl:"name:Return" json:"return"`
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
		if err := w.WriteData(o.HInterface); err != nil {
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
			if err := (&rrasm.DimInterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
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
		if err := w.ReadData(&o.HInterface); err != nil {
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
			o.Info = &rrasm.DimInterfaceContainer{}
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
	HInterface  uint32                       `idl:"name:hInterface" json:"h_interface"`
	TransportID uint32                       `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.DimInterfaceContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *RouterInterfaceTransportAddRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportAddOperation) *xxx_RouterInterfaceTransportAddOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportAddOperation{}
	}
	if o == nil {
		return op
	}
	op.HInterface = o.HInterface
	op.TransportID = o.TransportID
	op.Info = o.Info
	return op
}

func (o *RouterInterfaceTransportAddRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportAddOperation) {
	if o == nil {
		return
	}
	o.HInterface = op.HInterface
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
	HInterface  uint32                       `idl:"name:hInterface" json:"h_interface"`
	TransportID uint32                       `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.DimInterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	Return      uint32                       `idl:"name:Return" json:"return"`
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
		if err := w.WriteData(o.HInterface); err != nil {
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
			if err := (&rrasm.DimInterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
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
		if err := w.ReadData(&o.HInterface); err != nil {
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
			o.Info = &rrasm.DimInterfaceContainer{}
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
			if err := (&rrasm.DimInterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
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
			o.Info = &rrasm.DimInterfaceContainer{}
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
	HInterface  uint32                       `idl:"name:hInterface" json:"h_interface"`
	TransportID uint32                       `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.DimInterfaceContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *RouterInterfaceTransportGetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportGetInfoOperation) *xxx_RouterInterfaceTransportGetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportGetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.HInterface = o.HInterface
	op.TransportID = o.TransportID
	op.Info = o.Info
	return op
}

func (o *RouterInterfaceTransportGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportGetInfoOperation) {
	if o == nil {
		return
	}
	o.HInterface = op.HInterface
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
	Info *rrasm.DimInterfaceContainer `idl:"name:pInfoStruct" json:"info"`
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
	HInterface  uint32                       `idl:"name:hInterface" json:"h_interface"`
	TransportID uint32                       `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.DimInterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	Return      uint32                       `idl:"name:Return" json:"return"`
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
		if err := w.WriteData(o.HInterface); err != nil {
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
			if err := (&rrasm.DimInterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
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
		if err := w.ReadData(&o.HInterface); err != nil {
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
			o.Info = &rrasm.DimInterfaceContainer{}
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
	HInterface  uint32                       `idl:"name:hInterface" json:"h_interface"`
	TransportID uint32                       `idl:"name:dwTransportId" json:"transport_id"`
	Info        *rrasm.DimInterfaceContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *RouterInterfaceTransportSetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportSetInfoOperation) *xxx_RouterInterfaceTransportSetInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportSetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.HInterface = o.HInterface
	op.TransportID = o.TransportID
	op.Info = o.Info
	return op
}

func (o *RouterInterfaceTransportSetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportSetInfoOperation) {
	if o == nil {
		return
	}
	o.HInterface = op.HInterface
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
	Level                  uint32                         `idl:"name:dwLevel" json:"level"`
	Info                   *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	PreferredMaximumLength uint32                         `idl:"name:dwPreferedMaximumLength" json:"preferred_maximum_length"`
	EntriesRead            uint32                         `idl:"name:lpdwEntriesRead" json:"entries_read"`
	TotalEntries           uint32                         `idl:"name:lpdwTotalEntries" json:"total_entries"`
	Resume                 uint32                         `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	Return                 uint32                         `idl:"name:Return" json:"return"`
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
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
			o.Info = &rrasm.DimInformationContainer{}
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
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
			o.Info = &rrasm.DimInformationContainer{}
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
	Level                  uint32                         `idl:"name:dwLevel" json:"level"`
	Info                   *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	PreferredMaximumLength uint32                         `idl:"name:dwPreferedMaximumLength" json:"preferred_maximum_length"`
	Resume                 uint32                         `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
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
	Info         *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	EntriesRead  uint32                         `idl:"name:lpdwEntriesRead" json:"entries_read"`
	TotalEntries uint32                         `idl:"name:lpdwTotalEntries" json:"total_entries"`
	Resume       uint32                         `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
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
	HInterface       uint32 `idl:"name:hInterface" json:"h_interface"`
	HEvent           uint64 `idl:"name:hEvent" json:"h_event"`
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
		if err := w.WriteData(o.HInterface); err != nil {
			return err
		}
	}
	// hEvent {in} (1:{public, alias=ULONG_PTR}(uint32_64))
	{
		if err := w.WriteData(ndr.Uint3264(o.HEvent)); err != nil {
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
		if err := w.ReadData(&o.HInterface); err != nil {
			return err
		}
	}
	// hEvent {in} (1:{public, alias=ULONG_PTR}(uint32_64))
	{
		if err := w.ReadData((*ndr.Uint3264)(&o.HEvent)); err != nil {
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
	HInterface       uint32 `idl:"name:hInterface" json:"h_interface"`
	HEvent           uint64 `idl:"name:hEvent" json:"h_event"`
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
	op.HInterface = o.HInterface
	op.HEvent = o.HEvent
	op.Blocking = o.Blocking
	op.CallersProcessID = o.CallersProcessID
	return op
}

func (o *RouterInterfaceConnectRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceConnectOperation) {
	if o == nil {
		return
	}
	o.HInterface = op.HInterface
	o.HEvent = op.HEvent
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
	HInterface uint32 `idl:"name:hInterface" json:"h_interface"`
	Return     uint32 `idl:"name:Return" json:"return"`
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
		if err := w.WriteData(o.HInterface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceDisconnectOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.HInterface); err != nil {
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
	HInterface uint32 `idl:"name:hInterface" json:"h_interface"`
}

func (o *RouterInterfaceDisconnectRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceDisconnectOperation) *xxx_RouterInterfaceDisconnectOperation {
	if op == nil {
		op = &xxx_RouterInterfaceDisconnectOperation{}
	}
	if o == nil {
		return op
	}
	op.HInterface = o.HInterface
	return op
}

func (o *RouterInterfaceDisconnectRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceDisconnectOperation) {
	if o == nil {
		return
	}
	o.HInterface = op.HInterface
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
	HInterface      uint32 `idl:"name:hInterface" json:"h_interface"`
	TransportID     uint32 `idl:"name:dwTransportId" json:"transport_id"`
	HEvent          uint64 `idl:"name:hEvent" json:"h_event"`
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
		if err := w.WriteData(o.HInterface); err != nil {
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
		if err := w.WriteData(ndr.Uint3264(o.HEvent)); err != nil {
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
		if err := w.ReadData(&o.HInterface); err != nil {
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
		if err := w.ReadData((*ndr.Uint3264)(&o.HEvent)); err != nil {
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
	HInterface      uint32 `idl:"name:hInterface" json:"h_interface"`
	TransportID     uint32 `idl:"name:dwTransportId" json:"transport_id"`
	HEvent          uint64 `idl:"name:hEvent" json:"h_event"`
	ClientProcessID uint32 `idl:"name:dwClientProcessId" json:"client_process_id"`
}

func (o *RouterInterfaceUpdateRoutesRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceUpdateRoutesOperation) *xxx_RouterInterfaceUpdateRoutesOperation {
	if op == nil {
		op = &xxx_RouterInterfaceUpdateRoutesOperation{}
	}
	if o == nil {
		return op
	}
	op.HInterface = o.HInterface
	op.TransportID = o.TransportID
	op.HEvent = o.HEvent
	op.ClientProcessID = o.ClientProcessID
	return op
}

func (o *RouterInterfaceUpdateRoutesRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceUpdateRoutesOperation) {
	if o == nil {
		return
	}
	o.HInterface = op.HInterface
	o.TransportID = op.TransportID
	o.HEvent = op.HEvent
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
	HInterface   uint32 `idl:"name:hInterface" json:"h_interface"`
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
		if err := w.WriteData(o.HInterface); err != nil {
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
		if err := w.ReadData(&o.HInterface); err != nil {
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
	HInterface  uint32 `idl:"name:hInterface" json:"h_interface"`
	TransportID uint32 `idl:"name:dwTransportId" json:"transport_id"`
}

func (o *RouterInterfaceQueryUpdateResultRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceQueryUpdateResultOperation) *xxx_RouterInterfaceQueryUpdateResultOperation {
	if op == nil {
		op = &xxx_RouterInterfaceQueryUpdateResultOperation{}
	}
	if o == nil {
		return op
	}
	op.HInterface = o.HInterface
	op.TransportID = o.TransportID
	return op
}

func (o *RouterInterfaceQueryUpdateResultRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceQueryUpdateResultOperation) {
	if o == nil {
		return
	}
	o.HInterface = op.HInterface
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
	HInterface uint32 `idl:"name:hInterface" json:"h_interface"`
	Return     uint32 `idl:"name:Return" json:"return"`
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
		if err := w.WriteData(o.HInterface); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceUpdatePhonebookInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.HInterface); err != nil {
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
	HInterface uint32 `idl:"name:hInterface" json:"h_interface"`
}

func (o *RouterInterfaceUpdatePhonebookInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceUpdatePhonebookInfoOperation) *xxx_RouterInterfaceUpdatePhonebookInfoOperation {
	if op == nil {
		op = &xxx_RouterInterfaceUpdatePhonebookInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.HInterface = o.HInterface
	return op
}

func (o *RouterInterfaceUpdatePhonebookInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceUpdatePhonebookInfoOperation) {
	if o == nil {
		return
	}
	o.HInterface = op.HInterface
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

// xxx_RmibEntryCreateOperation structure represents the RMIBEntryCreate operation
type xxx_RmibEntryCreateOperation struct {
	PID        uint32                      `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                      `idl:"name:dwRoutingPid" json:"routing_pid"`
	InfoStuct  *rrasm.DimMIBEntryContainer `idl:"name:pInfoStuct" json:"info_stuct"`
	Return     uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RmibEntryCreateOperation) OpNum() int { return 26 }

func (o *xxx_RmibEntryCreateOperation) OpName() string { return "/dimsvc/v0/RMIBEntryCreate" }

func (o *xxx_RmibEntryCreateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntryCreateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
		if o.InfoStuct != nil {
			if err := o.InfoStuct.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.DimMIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntryCreateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
		if o.InfoStuct == nil {
			o.InfoStuct = &rrasm.DimMIBEntryContainer{}
		}
		if err := o.InfoStuct.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntryCreateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntryCreateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RmibEntryCreateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RmibEntryCreateRequest structure represents the RMIBEntryCreate operation request
type RmibEntryCreateRequest struct {
	PID        uint32                      `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                      `idl:"name:dwRoutingPid" json:"routing_pid"`
	InfoStuct  *rrasm.DimMIBEntryContainer `idl:"name:pInfoStuct" json:"info_stuct"`
}

func (o *RmibEntryCreateRequest) xxx_ToOp(ctx context.Context, op *xxx_RmibEntryCreateOperation) *xxx_RmibEntryCreateOperation {
	if op == nil {
		op = &xxx_RmibEntryCreateOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.InfoStuct = o.InfoStuct
	return op
}

func (o *RmibEntryCreateRequest) xxx_FromOp(ctx context.Context, op *xxx_RmibEntryCreateOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.InfoStuct = op.InfoStuct
}
func (o *RmibEntryCreateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RmibEntryCreateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RmibEntryCreateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RmibEntryCreateResponse structure represents the RMIBEntryCreate operation response
type RmibEntryCreateResponse struct {
	// Return: The RMIBEntryCreate return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RmibEntryCreateResponse) xxx_ToOp(ctx context.Context, op *xxx_RmibEntryCreateOperation) *xxx_RmibEntryCreateOperation {
	if op == nil {
		op = &xxx_RmibEntryCreateOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RmibEntryCreateResponse) xxx_FromOp(ctx context.Context, op *xxx_RmibEntryCreateOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RmibEntryCreateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RmibEntryCreateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RmibEntryCreateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RmibEntryDeleteOperation structure represents the RMIBEntryDelete operation
type xxx_RmibEntryDeleteOperation struct {
	PID        uint32                      `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                      `idl:"name:dwRoutingPid" json:"routing_pid"`
	InfoStuct  *rrasm.DimMIBEntryContainer `idl:"name:pInfoStuct" json:"info_stuct"`
	Return     uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RmibEntryDeleteOperation) OpNum() int { return 27 }

func (o *xxx_RmibEntryDeleteOperation) OpName() string { return "/dimsvc/v0/RMIBEntryDelete" }

func (o *xxx_RmibEntryDeleteOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntryDeleteOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
		if o.InfoStuct != nil {
			if err := o.InfoStuct.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.DimMIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntryDeleteOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
		if o.InfoStuct == nil {
			o.InfoStuct = &rrasm.DimMIBEntryContainer{}
		}
		if err := o.InfoStuct.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntryDeleteOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntryDeleteOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RmibEntryDeleteOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RmibEntryDeleteRequest structure represents the RMIBEntryDelete operation request
type RmibEntryDeleteRequest struct {
	PID        uint32                      `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                      `idl:"name:dwRoutingPid" json:"routing_pid"`
	InfoStuct  *rrasm.DimMIBEntryContainer `idl:"name:pInfoStuct" json:"info_stuct"`
}

func (o *RmibEntryDeleteRequest) xxx_ToOp(ctx context.Context, op *xxx_RmibEntryDeleteOperation) *xxx_RmibEntryDeleteOperation {
	if op == nil {
		op = &xxx_RmibEntryDeleteOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.InfoStuct = o.InfoStuct
	return op
}

func (o *RmibEntryDeleteRequest) xxx_FromOp(ctx context.Context, op *xxx_RmibEntryDeleteOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.InfoStuct = op.InfoStuct
}
func (o *RmibEntryDeleteRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RmibEntryDeleteRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RmibEntryDeleteOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RmibEntryDeleteResponse structure represents the RMIBEntryDelete operation response
type RmibEntryDeleteResponse struct {
	// Return: The RMIBEntryDelete return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RmibEntryDeleteResponse) xxx_ToOp(ctx context.Context, op *xxx_RmibEntryDeleteOperation) *xxx_RmibEntryDeleteOperation {
	if op == nil {
		op = &xxx_RmibEntryDeleteOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RmibEntryDeleteResponse) xxx_FromOp(ctx context.Context, op *xxx_RmibEntryDeleteOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RmibEntryDeleteResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RmibEntryDeleteResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RmibEntryDeleteOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RmibEntrySetOperation structure represents the RMIBEntrySet operation
type xxx_RmibEntrySetOperation struct {
	PID        uint32                      `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                      `idl:"name:dwRoutingPid" json:"routing_pid"`
	InfoStuct  *rrasm.DimMIBEntryContainer `idl:"name:pInfoStuct" json:"info_stuct"`
	Return     uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RmibEntrySetOperation) OpNum() int { return 28 }

func (o *xxx_RmibEntrySetOperation) OpName() string { return "/dimsvc/v0/RMIBEntrySet" }

func (o *xxx_RmibEntrySetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntrySetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
		if o.InfoStuct != nil {
			if err := o.InfoStuct.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.DimMIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntrySetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
		if o.InfoStuct == nil {
			o.InfoStuct = &rrasm.DimMIBEntryContainer{}
		}
		if err := o.InfoStuct.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntrySetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntrySetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RmibEntrySetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RmibEntrySetRequest structure represents the RMIBEntrySet operation request
type RmibEntrySetRequest struct {
	PID        uint32                      `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                      `idl:"name:dwRoutingPid" json:"routing_pid"`
	InfoStuct  *rrasm.DimMIBEntryContainer `idl:"name:pInfoStuct" json:"info_stuct"`
}

func (o *RmibEntrySetRequest) xxx_ToOp(ctx context.Context, op *xxx_RmibEntrySetOperation) *xxx_RmibEntrySetOperation {
	if op == nil {
		op = &xxx_RmibEntrySetOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.InfoStuct = o.InfoStuct
	return op
}

func (o *RmibEntrySetRequest) xxx_FromOp(ctx context.Context, op *xxx_RmibEntrySetOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.InfoStuct = op.InfoStuct
}
func (o *RmibEntrySetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RmibEntrySetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RmibEntrySetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RmibEntrySetResponse structure represents the RMIBEntrySet operation response
type RmibEntrySetResponse struct {
	// Return: The RMIBEntrySet return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RmibEntrySetResponse) xxx_ToOp(ctx context.Context, op *xxx_RmibEntrySetOperation) *xxx_RmibEntrySetOperation {
	if op == nil {
		op = &xxx_RmibEntrySetOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RmibEntrySetResponse) xxx_FromOp(ctx context.Context, op *xxx_RmibEntrySetOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RmibEntrySetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RmibEntrySetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RmibEntrySetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RmibEntryGetOperation structure represents the RMIBEntryGet operation
type xxx_RmibEntryGetOperation struct {
	PID        uint32                      `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                      `idl:"name:dwRoutingPid" json:"routing_pid"`
	InfoStuct  *rrasm.DimMIBEntryContainer `idl:"name:pInfoStuct" json:"info_stuct"`
	Return     uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RmibEntryGetOperation) OpNum() int { return 29 }

func (o *xxx_RmibEntryGetOperation) OpName() string { return "/dimsvc/v0/RMIBEntryGet" }

func (o *xxx_RmibEntryGetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntryGetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
		if o.InfoStuct != nil {
			if err := o.InfoStuct.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.DimMIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntryGetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
		if o.InfoStuct == nil {
			o.InfoStuct = &rrasm.DimMIBEntryContainer{}
		}
		if err := o.InfoStuct.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntryGetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntryGetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.InfoStuct != nil {
			if err := o.InfoStuct.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.DimMIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RmibEntryGetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.InfoStuct == nil {
			o.InfoStuct = &rrasm.DimMIBEntryContainer{}
		}
		if err := o.InfoStuct.UnmarshalNDR(ctx, w); err != nil {
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

// RmibEntryGetRequest structure represents the RMIBEntryGet operation request
type RmibEntryGetRequest struct {
	PID        uint32                      `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                      `idl:"name:dwRoutingPid" json:"routing_pid"`
	InfoStuct  *rrasm.DimMIBEntryContainer `idl:"name:pInfoStuct" json:"info_stuct"`
}

func (o *RmibEntryGetRequest) xxx_ToOp(ctx context.Context, op *xxx_RmibEntryGetOperation) *xxx_RmibEntryGetOperation {
	if op == nil {
		op = &xxx_RmibEntryGetOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.InfoStuct = o.InfoStuct
	return op
}

func (o *RmibEntryGetRequest) xxx_FromOp(ctx context.Context, op *xxx_RmibEntryGetOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.InfoStuct = op.InfoStuct
}
func (o *RmibEntryGetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RmibEntryGetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RmibEntryGetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RmibEntryGetResponse structure represents the RMIBEntryGet operation response
type RmibEntryGetResponse struct {
	InfoStuct *rrasm.DimMIBEntryContainer `idl:"name:pInfoStuct" json:"info_stuct"`
	// Return: The RMIBEntryGet return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RmibEntryGetResponse) xxx_ToOp(ctx context.Context, op *xxx_RmibEntryGetOperation) *xxx_RmibEntryGetOperation {
	if op == nil {
		op = &xxx_RmibEntryGetOperation{}
	}
	if o == nil {
		return op
	}
	op.InfoStuct = o.InfoStuct
	op.Return = o.Return
	return op
}

func (o *RmibEntryGetResponse) xxx_FromOp(ctx context.Context, op *xxx_RmibEntryGetOperation) {
	if o == nil {
		return
	}
	o.InfoStuct = op.InfoStuct
	o.Return = op.Return
}
func (o *RmibEntryGetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RmibEntryGetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RmibEntryGetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RmibEntryGetFirstOperation structure represents the RMIBEntryGetFirst operation
type xxx_RmibEntryGetFirstOperation struct {
	PID        uint32                      `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                      `idl:"name:dwRoutingPid" json:"routing_pid"`
	InfoStuct  *rrasm.DimMIBEntryContainer `idl:"name:pInfoStuct" json:"info_stuct"`
	Return     uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RmibEntryGetFirstOperation) OpNum() int { return 30 }

func (o *xxx_RmibEntryGetFirstOperation) OpName() string { return "/dimsvc/v0/RMIBEntryGetFirst" }

func (o *xxx_RmibEntryGetFirstOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntryGetFirstOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
		if o.InfoStuct != nil {
			if err := o.InfoStuct.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.DimMIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntryGetFirstOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
		if o.InfoStuct == nil {
			o.InfoStuct = &rrasm.DimMIBEntryContainer{}
		}
		if err := o.InfoStuct.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntryGetFirstOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntryGetFirstOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.InfoStuct != nil {
			if err := o.InfoStuct.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.DimMIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RmibEntryGetFirstOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.InfoStuct == nil {
			o.InfoStuct = &rrasm.DimMIBEntryContainer{}
		}
		if err := o.InfoStuct.UnmarshalNDR(ctx, w); err != nil {
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

// RmibEntryGetFirstRequest structure represents the RMIBEntryGetFirst operation request
type RmibEntryGetFirstRequest struct {
	PID        uint32                      `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                      `idl:"name:dwRoutingPid" json:"routing_pid"`
	InfoStuct  *rrasm.DimMIBEntryContainer `idl:"name:pInfoStuct" json:"info_stuct"`
}

func (o *RmibEntryGetFirstRequest) xxx_ToOp(ctx context.Context, op *xxx_RmibEntryGetFirstOperation) *xxx_RmibEntryGetFirstOperation {
	if op == nil {
		op = &xxx_RmibEntryGetFirstOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.InfoStuct = o.InfoStuct
	return op
}

func (o *RmibEntryGetFirstRequest) xxx_FromOp(ctx context.Context, op *xxx_RmibEntryGetFirstOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.InfoStuct = op.InfoStuct
}
func (o *RmibEntryGetFirstRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RmibEntryGetFirstRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RmibEntryGetFirstOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RmibEntryGetFirstResponse structure represents the RMIBEntryGetFirst operation response
type RmibEntryGetFirstResponse struct {
	InfoStuct *rrasm.DimMIBEntryContainer `idl:"name:pInfoStuct" json:"info_stuct"`
	// Return: The RMIBEntryGetFirst return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RmibEntryGetFirstResponse) xxx_ToOp(ctx context.Context, op *xxx_RmibEntryGetFirstOperation) *xxx_RmibEntryGetFirstOperation {
	if op == nil {
		op = &xxx_RmibEntryGetFirstOperation{}
	}
	if o == nil {
		return op
	}
	op.InfoStuct = o.InfoStuct
	op.Return = o.Return
	return op
}

func (o *RmibEntryGetFirstResponse) xxx_FromOp(ctx context.Context, op *xxx_RmibEntryGetFirstOperation) {
	if o == nil {
		return
	}
	o.InfoStuct = op.InfoStuct
	o.Return = op.Return
}
func (o *RmibEntryGetFirstResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RmibEntryGetFirstResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RmibEntryGetFirstOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RmibEntryGetNextOperation structure represents the RMIBEntryGetNext operation
type xxx_RmibEntryGetNextOperation struct {
	PID        uint32                      `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                      `idl:"name:dwRoutingPid" json:"routing_pid"`
	InfoStuct  *rrasm.DimMIBEntryContainer `idl:"name:pInfoStuct" json:"info_stuct"`
	Return     uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RmibEntryGetNextOperation) OpNum() int { return 31 }

func (o *xxx_RmibEntryGetNextOperation) OpName() string { return "/dimsvc/v0/RMIBEntryGetNext" }

func (o *xxx_RmibEntryGetNextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntryGetNextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
		if o.InfoStuct != nil {
			if err := o.InfoStuct.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.DimMIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntryGetNextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
		if o.InfoStuct == nil {
			o.InfoStuct = &rrasm.DimMIBEntryContainer{}
		}
		if err := o.InfoStuct.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntryGetNextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibEntryGetNextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.InfoStuct != nil {
			if err := o.InfoStuct.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&rrasm.DimMIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RmibEntryGetNextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStuct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.InfoStuct == nil {
			o.InfoStuct = &rrasm.DimMIBEntryContainer{}
		}
		if err := o.InfoStuct.UnmarshalNDR(ctx, w); err != nil {
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

// RmibEntryGetNextRequest structure represents the RMIBEntryGetNext operation request
type RmibEntryGetNextRequest struct {
	PID        uint32                      `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                      `idl:"name:dwRoutingPid" json:"routing_pid"`
	InfoStuct  *rrasm.DimMIBEntryContainer `idl:"name:pInfoStuct" json:"info_stuct"`
}

func (o *RmibEntryGetNextRequest) xxx_ToOp(ctx context.Context, op *xxx_RmibEntryGetNextOperation) *xxx_RmibEntryGetNextOperation {
	if op == nil {
		op = &xxx_RmibEntryGetNextOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.InfoStuct = o.InfoStuct
	return op
}

func (o *RmibEntryGetNextRequest) xxx_FromOp(ctx context.Context, op *xxx_RmibEntryGetNextOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.InfoStuct = op.InfoStuct
}
func (o *RmibEntryGetNextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RmibEntryGetNextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RmibEntryGetNextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RmibEntryGetNextResponse structure represents the RMIBEntryGetNext operation response
type RmibEntryGetNextResponse struct {
	InfoStuct *rrasm.DimMIBEntryContainer `idl:"name:pInfoStuct" json:"info_stuct"`
	// Return: The RMIBEntryGetNext return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RmibEntryGetNextResponse) xxx_ToOp(ctx context.Context, op *xxx_RmibEntryGetNextOperation) *xxx_RmibEntryGetNextOperation {
	if op == nil {
		op = &xxx_RmibEntryGetNextOperation{}
	}
	if o == nil {
		return op
	}
	op.InfoStuct = o.InfoStuct
	op.Return = o.Return
	return op
}

func (o *RmibEntryGetNextResponse) xxx_FromOp(ctx context.Context, op *xxx_RmibEntryGetNextOperation) {
	if o == nil {
		return
	}
	o.InfoStuct = op.InfoStuct
	o.Return = op.Return
}
func (o *RmibEntryGetNextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RmibEntryGetNextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RmibEntryGetNextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RmibGetTrapInfoOperation structure represents the RMIBGetTrapInfo operation
type xxx_RmibGetTrapInfoOperation struct {
	PID        uint32                      `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                      `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.DimMIBEntryContainer `idl:"name:pInfoStruct" json:"info"`
	Return     uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RmibGetTrapInfoOperation) OpNum() int { return 32 }

func (o *xxx_RmibGetTrapInfoOperation) OpName() string { return "/dimsvc/v0/RMIBGetTrapInfo" }

func (o *xxx_RmibGetTrapInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibGetTrapInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
			if err := (&rrasm.DimMIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibGetTrapInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
			o.Info = &rrasm.DimMIBEntryContainer{}
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

func (o *xxx_RmibGetTrapInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibGetTrapInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
			if err := (&rrasm.DimMIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RmibGetTrapInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.DimMIBEntryContainer{}
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

// RmibGetTrapInfoRequest structure represents the RMIBGetTrapInfo operation request
type RmibGetTrapInfoRequest struct {
	PID        uint32                      `idl:"name:dwPid" json:"pid"`
	RoutingPID uint32                      `idl:"name:dwRoutingPid" json:"routing_pid"`
	Info       *rrasm.DimMIBEntryContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *RmibGetTrapInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RmibGetTrapInfoOperation) *xxx_RmibGetTrapInfoOperation {
	if op == nil {
		op = &xxx_RmibGetTrapInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.Info = o.Info
	return op
}

func (o *RmibGetTrapInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RmibGetTrapInfoOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.Info = op.Info
}
func (o *RmibGetTrapInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RmibGetTrapInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RmibGetTrapInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RmibGetTrapInfoResponse structure represents the RMIBGetTrapInfo operation response
type RmibGetTrapInfoResponse struct {
	Info *rrasm.DimMIBEntryContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RMIBGetTrapInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RmibGetTrapInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RmibGetTrapInfoOperation) *xxx_RmibGetTrapInfoOperation {
	if op == nil {
		op = &xxx_RmibGetTrapInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *RmibGetTrapInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RmibGetTrapInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *RmibGetTrapInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RmibGetTrapInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RmibGetTrapInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RmibSetTrapInfoOperation structure represents the RMIBSetTrapInfo operation
type xxx_RmibSetTrapInfoOperation struct {
	PID             uint32                      `idl:"name:dwPid" json:"pid"`
	RoutingPID      uint32                      `idl:"name:dwRoutingPid" json:"routing_pid"`
	HEvent          uint64                      `idl:"name:hEvent" json:"h_event"`
	ClientProcessID uint32                      `idl:"name:dwClientProcessId" json:"client_process_id"`
	Info            *rrasm.DimMIBEntryContainer `idl:"name:pInfoStruct" json:"info"`
	Return          uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_RmibSetTrapInfoOperation) OpNum() int { return 33 }

func (o *xxx_RmibSetTrapInfoOperation) OpName() string { return "/dimsvc/v0/RMIBSetTrapInfo" }

func (o *xxx_RmibSetTrapInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibSetTrapInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
		if err := w.WriteData(ndr.Uint3264(o.HEvent)); err != nil {
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
			if err := (&rrasm.DimMIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibSetTrapInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
		if err := w.ReadData((*ndr.Uint3264)(&o.HEvent)); err != nil {
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
			o.Info = &rrasm.DimMIBEntryContainer{}
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

func (o *xxx_RmibSetTrapInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RmibSetTrapInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
			if err := (&rrasm.DimMIBEntryContainer{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RmibSetTrapInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pInfoStruct {in, out} (1:{alias=PDIM_MIB_ENTRY_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_MIB_ENTRY_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.DimMIBEntryContainer{}
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

// RmibSetTrapInfoRequest structure represents the RMIBSetTrapInfo operation request
type RmibSetTrapInfoRequest struct {
	PID             uint32                      `idl:"name:dwPid" json:"pid"`
	RoutingPID      uint32                      `idl:"name:dwRoutingPid" json:"routing_pid"`
	HEvent          uint64                      `idl:"name:hEvent" json:"h_event"`
	ClientProcessID uint32                      `idl:"name:dwClientProcessId" json:"client_process_id"`
	Info            *rrasm.DimMIBEntryContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *RmibSetTrapInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_RmibSetTrapInfoOperation) *xxx_RmibSetTrapInfoOperation {
	if op == nil {
		op = &xxx_RmibSetTrapInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.PID = o.PID
	op.RoutingPID = o.RoutingPID
	op.HEvent = o.HEvent
	op.ClientProcessID = o.ClientProcessID
	op.Info = o.Info
	return op
}

func (o *RmibSetTrapInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RmibSetTrapInfoOperation) {
	if o == nil {
		return
	}
	o.PID = op.PID
	o.RoutingPID = op.RoutingPID
	o.HEvent = op.HEvent
	o.ClientProcessID = op.ClientProcessID
	o.Info = op.Info
}
func (o *RmibSetTrapInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RmibSetTrapInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RmibSetTrapInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RmibSetTrapInfoResponse structure represents the RMIBSetTrapInfo operation response
type RmibSetTrapInfoResponse struct {
	Info *rrasm.DimMIBEntryContainer `idl:"name:pInfoStruct" json:"info"`
	// Return: The RMIBSetTrapInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RmibSetTrapInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_RmibSetTrapInfoOperation) *xxx_RmibSetTrapInfoOperation {
	if op == nil {
		op = &xxx_RmibSetTrapInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *RmibSetTrapInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_RmibSetTrapInfoOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *RmibSetTrapInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RmibSetTrapInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RmibSetTrapInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RASAdminConnectionNotificationOperation structure represents the RRasAdminConnectionNotification operation
type xxx_RASAdminConnectionNotificationOperation struct {
	Register           uint32 `idl:"name:fRegister" json:"register"`
	ClientProcessID    uint32 `idl:"name:dwClientProcessId" json:"client_process_id"`
	HEventNotification uint64 `idl:"name:hEventNotification" json:"h_event_notification"`
	Return             uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RASAdminConnectionNotificationOperation) OpNum() int { return 34 }

func (o *xxx_RASAdminConnectionNotificationOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminConnectionNotification"
}

func (o *xxx_RASAdminConnectionNotificationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionNotificationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
		if err := w.WriteData(ndr.Uint3264(o.HEventNotification)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionNotificationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
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
		if err := w.ReadData((*ndr.Uint3264)(&o.HEventNotification)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionNotificationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionNotificationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RASAdminConnectionNotificationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RASAdminConnectionNotificationRequest structure represents the RRasAdminConnectionNotification operation request
type RASAdminConnectionNotificationRequest struct {
	Register           uint32 `idl:"name:fRegister" json:"register"`
	ClientProcessID    uint32 `idl:"name:dwClientProcessId" json:"client_process_id"`
	HEventNotification uint64 `idl:"name:hEventNotification" json:"h_event_notification"`
}

func (o *RASAdminConnectionNotificationRequest) xxx_ToOp(ctx context.Context, op *xxx_RASAdminConnectionNotificationOperation) *xxx_RASAdminConnectionNotificationOperation {
	if op == nil {
		op = &xxx_RASAdminConnectionNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.Register = o.Register
	op.ClientProcessID = o.ClientProcessID
	op.HEventNotification = o.HEventNotification
	return op
}

func (o *RASAdminConnectionNotificationRequest) xxx_FromOp(ctx context.Context, op *xxx_RASAdminConnectionNotificationOperation) {
	if o == nil {
		return
	}
	o.Register = op.Register
	o.ClientProcessID = op.ClientProcessID
	o.HEventNotification = op.HEventNotification
}
func (o *RASAdminConnectionNotificationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASAdminConnectionNotificationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminConnectionNotificationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASAdminConnectionNotificationResponse structure represents the RRasAdminConnectionNotification operation response
type RASAdminConnectionNotificationResponse struct {
	// Return: The RRasAdminConnectionNotification return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASAdminConnectionNotificationResponse) xxx_ToOp(ctx context.Context, op *xxx_RASAdminConnectionNotificationOperation) *xxx_RASAdminConnectionNotificationOperation {
	if op == nil {
		op = &xxx_RASAdminConnectionNotificationOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RASAdminConnectionNotificationResponse) xxx_FromOp(ctx context.Context, op *xxx_RASAdminConnectionNotificationOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RASAdminConnectionNotificationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASAdminConnectionNotificationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminConnectionNotificationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RASAdminSendUserMessageOperation structure represents the RRasAdminSendUserMessage operation
type xxx_RASAdminSendUserMessageOperation struct {
	HDimConnection uint32 `idl:"name:hDimConnection" json:"h_dim_connection"`
	Message        string `idl:"name:lpwszMessage;string" json:"message"`
	Return         uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RASAdminSendUserMessageOperation) OpNum() int { return 35 }

func (o *xxx_RASAdminSendUserMessageOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminSendUserMessage"
}

func (o *xxx_RASAdminSendUserMessageOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminSendUserMessageOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.HDimConnection); err != nil {
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

func (o *xxx_RASAdminSendUserMessageOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.HDimConnection); err != nil {
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

func (o *xxx_RASAdminSendUserMessageOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminSendUserMessageOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RASAdminSendUserMessageOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RASAdminSendUserMessageRequest structure represents the RRasAdminSendUserMessage operation request
type RASAdminSendUserMessageRequest struct {
	HDimConnection uint32 `idl:"name:hDimConnection" json:"h_dim_connection"`
	Message        string `idl:"name:lpwszMessage;string" json:"message"`
}

func (o *RASAdminSendUserMessageRequest) xxx_ToOp(ctx context.Context, op *xxx_RASAdminSendUserMessageOperation) *xxx_RASAdminSendUserMessageOperation {
	if op == nil {
		op = &xxx_RASAdminSendUserMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.HDimConnection = o.HDimConnection
	op.Message = o.Message
	return op
}

func (o *RASAdminSendUserMessageRequest) xxx_FromOp(ctx context.Context, op *xxx_RASAdminSendUserMessageOperation) {
	if o == nil {
		return
	}
	o.HDimConnection = op.HDimConnection
	o.Message = op.Message
}
func (o *RASAdminSendUserMessageRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASAdminSendUserMessageRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminSendUserMessageOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASAdminSendUserMessageResponse structure represents the RRasAdminSendUserMessage operation response
type RASAdminSendUserMessageResponse struct {
	// Return: The RRasAdminSendUserMessage return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASAdminSendUserMessageResponse) xxx_ToOp(ctx context.Context, op *xxx_RASAdminSendUserMessageOperation) *xxx_RASAdminSendUserMessageOperation {
	if op == nil {
		op = &xxx_RASAdminSendUserMessageOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RASAdminSendUserMessageResponse) xxx_FromOp(ctx context.Context, op *xxx_RASAdminSendUserMessageOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RASAdminSendUserMessageResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASAdminSendUserMessageResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminSendUserMessageOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterDeviceEnumOperation structure represents the RRouterDeviceEnum operation
type xxx_RouterDeviceEnumOperation struct {
	Level        uint32                         `idl:"name:dwLevel" json:"level"`
	Info         *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	TotalEntries uint32                         `idl:"name:lpdwTotalEntries" json:"total_entries"`
	Return       uint32                         `idl:"name:Return" json:"return"`
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
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
			o.Info = &rrasm.DimInformationContainer{}
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
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
			o.Info = &rrasm.DimInformationContainer{}
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
	Level        uint32                         `idl:"name:dwLevel" json:"level"`
	Info         *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	TotalEntries uint32                         `idl:"name:lpdwTotalEntries" json:"total_entries"`
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
	Info         *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	TotalEntries uint32                         `idl:"name:lpdwTotalEntries" json:"total_entries"`
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
	TransportID       uint32                       `idl:"name:dwTransportId" json:"transport_id"`
	LpwsTransportName string                       `idl:"name:lpwsTransportName;string" json:"lpws_transport_name"`
	Info              *rrasm.DimInterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	LpwsDLLPath       string                       `idl:"name:lpwsDLLPath;string" json:"lpws_dll_path"`
	Return            uint32                       `idl:"name:Return" json:"return"`
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
		if err := ndr.WriteUTF16NString(ctx, w, o.LpwsTransportName); err != nil {
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
			if err := (&rrasm.DimInterfaceContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpwsDLLPath {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.LpwsDLLPath); err != nil {
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
		if err := ndr.ReadUTF16NString(ctx, w, &o.LpwsTransportName); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INTERFACE_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INTERFACE_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.DimInterfaceContainer{}
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
		if err := ndr.ReadUTF16NString(ctx, w, &o.LpwsDLLPath); err != nil {
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
	TransportID       uint32                       `idl:"name:dwTransportId" json:"transport_id"`
	LpwsTransportName string                       `idl:"name:lpwsTransportName;string" json:"lpws_transport_name"`
	Info              *rrasm.DimInterfaceContainer `idl:"name:pInfoStruct" json:"info"`
	LpwsDLLPath       string                       `idl:"name:lpwsDLLPath;string" json:"lpws_dll_path"`
}

func (o *RouterInterfaceTransportCreateRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceTransportCreateOperation) *xxx_RouterInterfaceTransportCreateOperation {
	if op == nil {
		op = &xxx_RouterInterfaceTransportCreateOperation{}
	}
	if o == nil {
		return op
	}
	op.TransportID = o.TransportID
	op.LpwsTransportName = o.LpwsTransportName
	op.Info = o.Info
	op.LpwsDLLPath = o.LpwsDLLPath
	return op
}

func (o *RouterInterfaceTransportCreateRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceTransportCreateOperation) {
	if o == nil {
		return
	}
	o.TransportID = op.TransportID
	o.LpwsTransportName = op.LpwsTransportName
	o.Info = op.Info
	o.LpwsDLLPath = op.LpwsDLLPath
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
	Level      uint32                         `idl:"name:dwLevel" json:"level"`
	Info       *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	Index      uint32                         `idl:"name:dwIndex" json:"index"`
	HInterface uint32                         `idl:"name:hInterface" json:"h_interface"`
	Return     uint32                         `idl:"name:Return" json:"return"`
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
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
		if err := w.WriteData(o.HInterface); err != nil {
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
			o.Info = &rrasm.DimInformationContainer{}
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
		if err := w.ReadData(&o.HInterface); err != nil {
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
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
			o.Info = &rrasm.DimInformationContainer{}
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
	Level      uint32                         `idl:"name:dwLevel" json:"level"`
	Info       *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	Index      uint32                         `idl:"name:dwIndex" json:"index"`
	HInterface uint32                         `idl:"name:hInterface" json:"h_interface"`
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
	op.HInterface = o.HInterface
	return op
}

func (o *RouterInterfaceDeviceGetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceDeviceGetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.Index = op.Index
	o.HInterface = op.HInterface
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
	Info *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
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
	Level      uint32                         `idl:"name:dwLevel" json:"level"`
	Info       *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	Index      uint32                         `idl:"name:dwIndex" json:"index"`
	HInterface uint32                         `idl:"name:hInterface" json:"h_interface"`
	Return     uint32                         `idl:"name:Return" json:"return"`
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
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
		if err := w.WriteData(o.HInterface); err != nil {
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
			o.Info = &rrasm.DimInformationContainer{}
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
		if err := w.ReadData(&o.HInterface); err != nil {
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
	Level      uint32                         `idl:"name:dwLevel" json:"level"`
	Info       *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	Index      uint32                         `idl:"name:dwIndex" json:"index"`
	HInterface uint32                         `idl:"name:hInterface" json:"h_interface"`
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
	op.HInterface = o.HInterface
	return op
}

func (o *RouterInterfaceDeviceSetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceDeviceSetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.Index = op.Index
	o.HInterface = op.HInterface
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
	Level      uint32                         `idl:"name:dwLevel" json:"level"`
	Info       *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	HInterface uint32                         `idl:"name:hInterface" json:"h_interface"`
	Return     uint32                         `idl:"name:Return" json:"return"`
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.HInterface); err != nil {
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
			o.Info = &rrasm.DimInformationContainer{}
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
		if err := w.ReadData(&o.HInterface); err != nil {
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
	Level      uint32                         `idl:"name:dwLevel" json:"level"`
	Info       *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	HInterface uint32                         `idl:"name:hInterface" json:"h_interface"`
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
	op.HInterface = o.HInterface
	return op
}

func (o *RouterInterfaceSetCredentialsExRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceSetCredentialsExOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.HInterface = op.HInterface
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
	Level      uint32                         `idl:"name:dwLevel" json:"level"`
	Info       *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	HInterface uint32                         `idl:"name:hInterface" json:"h_interface"`
	Return     uint32                         `idl:"name:Return" json:"return"`
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// hInterface {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.HInterface); err != nil {
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
			o.Info = &rrasm.DimInformationContainer{}
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
		if err := w.ReadData(&o.HInterface); err != nil {
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
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
			o.Info = &rrasm.DimInformationContainer{}
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
	Level      uint32                         `idl:"name:dwLevel" json:"level"`
	Info       *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	HInterface uint32                         `idl:"name:hInterface" json:"h_interface"`
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
	op.HInterface = o.HInterface
	return op
}

func (o *RouterInterfaceGetCredentialsExRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetCredentialsExOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
	o.HInterface = op.HInterface
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
	Info *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
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

// xxx_RASAdminConnectionRemoveQuarantineOperation structure represents the RRasAdminConnectionRemoveQuarantine operation
type xxx_RASAdminConnectionRemoveQuarantineOperation struct {
	HRASConnection uint32 `idl:"name:hRasConnection" json:"h_ras_connection"`
	IsIPAddress    bool   `idl:"name:fIsIpAddress" json:"is_ip_address"`
	Return         uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RASAdminConnectionRemoveQuarantineOperation) OpNum() int { return 42 }

func (o *xxx_RASAdminConnectionRemoveQuarantineOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminConnectionRemoveQuarantine"
}

func (o *xxx_RASAdminConnectionRemoveQuarantineOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionRemoveQuarantineOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hRasConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.HRASConnection); err != nil {
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

func (o *xxx_RASAdminConnectionRemoveQuarantineOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hRasConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.HRASConnection); err != nil {
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

func (o *xxx_RASAdminConnectionRemoveQuarantineOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionRemoveQuarantineOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RASAdminConnectionRemoveQuarantineOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RASAdminConnectionRemoveQuarantineRequest structure represents the RRasAdminConnectionRemoveQuarantine operation request
type RASAdminConnectionRemoveQuarantineRequest struct {
	HRASConnection uint32 `idl:"name:hRasConnection" json:"h_ras_connection"`
	IsIPAddress    bool   `idl:"name:fIsIpAddress" json:"is_ip_address"`
}

func (o *RASAdminConnectionRemoveQuarantineRequest) xxx_ToOp(ctx context.Context, op *xxx_RASAdminConnectionRemoveQuarantineOperation) *xxx_RASAdminConnectionRemoveQuarantineOperation {
	if op == nil {
		op = &xxx_RASAdminConnectionRemoveQuarantineOperation{}
	}
	if o == nil {
		return op
	}
	op.HRASConnection = o.HRASConnection
	op.IsIPAddress = o.IsIPAddress
	return op
}

func (o *RASAdminConnectionRemoveQuarantineRequest) xxx_FromOp(ctx context.Context, op *xxx_RASAdminConnectionRemoveQuarantineOperation) {
	if o == nil {
		return
	}
	o.HRASConnection = op.HRASConnection
	o.IsIPAddress = op.IsIPAddress
}
func (o *RASAdminConnectionRemoveQuarantineRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASAdminConnectionRemoveQuarantineRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminConnectionRemoveQuarantineOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASAdminConnectionRemoveQuarantineResponse structure represents the RRasAdminConnectionRemoveQuarantine operation response
type RASAdminConnectionRemoveQuarantineResponse struct {
	// Return: The RRasAdminConnectionRemoveQuarantine return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASAdminConnectionRemoveQuarantineResponse) xxx_ToOp(ctx context.Context, op *xxx_RASAdminConnectionRemoveQuarantineOperation) *xxx_RASAdminConnectionRemoveQuarantineOperation {
	if op == nil {
		op = &xxx_RASAdminConnectionRemoveQuarantineOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RASAdminConnectionRemoveQuarantineResponse) xxx_FromOp(ctx context.Context, op *xxx_RASAdminConnectionRemoveQuarantineOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RASAdminConnectionRemoveQuarantineResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASAdminConnectionRemoveQuarantineResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminConnectionRemoveQuarantineOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MprAdminServerSetInfoOperation structure represents the RMprAdminServerSetInfo operation
type xxx_MprAdminServerSetInfoOperation struct {
	Level  uint32                         `idl:"name:dwLevel" json:"level"`
	Info   *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
	Return uint32                         `idl:"name:Return" json:"return"`
}

func (o *xxx_MprAdminServerSetInfoOperation) OpNum() int { return 43 }

func (o *xxx_MprAdminServerSetInfoOperation) OpName() string {
	return "/dimsvc/v0/RMprAdminServerSetInfo"
}

func (o *xxx_MprAdminServerSetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MprAdminServerSetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
			if err := (&rrasm.DimInformationContainer{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MprAdminServerSetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// pInfoStruct {in} (1:{alias=PDIM_INFORMATION_CONTAINER,pointer=ref}*(1))(2:{alias=DIM_INFORMATION_CONTAINER}(struct))
	{
		if o.Info == nil {
			o.Info = &rrasm.DimInformationContainer{}
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

func (o *xxx_MprAdminServerSetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MprAdminServerSetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_MprAdminServerSetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MprAdminServerSetInfoRequest structure represents the RMprAdminServerSetInfo operation request
type MprAdminServerSetInfoRequest struct {
	Level uint32                         `idl:"name:dwLevel" json:"level"`
	Info  *rrasm.DimInformationContainer `idl:"name:pInfoStruct" json:"info"`
}

func (o *MprAdminServerSetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_MprAdminServerSetInfoOperation) *xxx_MprAdminServerSetInfoOperation {
	if op == nil {
		op = &xxx_MprAdminServerSetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.Info = o.Info
	return op
}

func (o *MprAdminServerSetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_MprAdminServerSetInfoOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.Info = op.Info
}
func (o *MprAdminServerSetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MprAdminServerSetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MprAdminServerSetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MprAdminServerSetInfoResponse structure represents the RMprAdminServerSetInfo operation response
type MprAdminServerSetInfoResponse struct {
	// Return: The RMprAdminServerSetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MprAdminServerSetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_MprAdminServerSetInfoOperation) *xxx_MprAdminServerSetInfoOperation {
	if op == nil {
		op = &xxx_MprAdminServerSetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *MprAdminServerSetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_MprAdminServerSetInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *MprAdminServerSetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MprAdminServerSetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MprAdminServerSetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MprAdminServerGetInfoExOperation structure represents the RMprAdminServerGetInfoEx operation
type xxx_MprAdminServerGetInfoExOperation struct {
	ServerConfig *rrasm.MprServerExIDL `idl:"name:pServerConfig" json:"server_config"`
	Return       uint32                `idl:"name:Return" json:"return"`
}

func (o *xxx_MprAdminServerGetInfoExOperation) OpNum() int { return 44 }

func (o *xxx_MprAdminServerGetInfoExOperation) OpName() string {
	return "/dimsvc/v0/RMprAdminServerGetInfoEx"
}

func (o *xxx_MprAdminServerGetInfoExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MprAdminServerGetInfoExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
			if err := (&rrasm.MprServerExIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MprAdminServerGetInfoExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pServerConfig {in, out} (1:{alias=PMPR_SERVER_EX_IDL,pointer=ref}*(1))(2:{alias=MPR_SERVER_EX_IDL}(union))
	{
		if o.ServerConfig == nil {
			o.ServerConfig = &rrasm.MprServerExIDL{}
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

func (o *xxx_MprAdminServerGetInfoExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MprAdminServerGetInfoExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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
			if err := (&rrasm.MprServerExIDL{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_MprAdminServerGetInfoExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pServerConfig {in, out} (1:{alias=PMPR_SERVER_EX_IDL,pointer=ref}*(1))(2:{alias=MPR_SERVER_EX_IDL}(union))
	{
		if o.ServerConfig == nil {
			o.ServerConfig = &rrasm.MprServerExIDL{}
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

// MprAdminServerGetInfoExRequest structure represents the RMprAdminServerGetInfoEx operation request
type MprAdminServerGetInfoExRequest struct {
	ServerConfig *rrasm.MprServerExIDL `idl:"name:pServerConfig" json:"server_config"`
}

func (o *MprAdminServerGetInfoExRequest) xxx_ToOp(ctx context.Context, op *xxx_MprAdminServerGetInfoExOperation) *xxx_MprAdminServerGetInfoExOperation {
	if op == nil {
		op = &xxx_MprAdminServerGetInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerConfig = o.ServerConfig
	return op
}

func (o *MprAdminServerGetInfoExRequest) xxx_FromOp(ctx context.Context, op *xxx_MprAdminServerGetInfoExOperation) {
	if o == nil {
		return
	}
	o.ServerConfig = op.ServerConfig
}
func (o *MprAdminServerGetInfoExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MprAdminServerGetInfoExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MprAdminServerGetInfoExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MprAdminServerGetInfoExResponse structure represents the RMprAdminServerGetInfoEx operation response
type MprAdminServerGetInfoExResponse struct {
	ServerConfig *rrasm.MprServerExIDL `idl:"name:pServerConfig" json:"server_config"`
	// Return: The RMprAdminServerGetInfoEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MprAdminServerGetInfoExResponse) xxx_ToOp(ctx context.Context, op *xxx_MprAdminServerGetInfoExOperation) *xxx_MprAdminServerGetInfoExOperation {
	if op == nil {
		op = &xxx_MprAdminServerGetInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerConfig = o.ServerConfig
	op.Return = o.Return
	return op
}

func (o *MprAdminServerGetInfoExResponse) xxx_FromOp(ctx context.Context, op *xxx_MprAdminServerGetInfoExOperation) {
	if o == nil {
		return
	}
	o.ServerConfig = op.ServerConfig
	o.Return = op.Return
}
func (o *MprAdminServerGetInfoExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MprAdminServerGetInfoExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MprAdminServerGetInfoExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RASAdminConnectionEnumExOperation structure represents the RRasAdminConnectionEnumEx operation
type xxx_RASAdminConnectionEnumExOperation struct {
	ObjectHeader        *rrasm.MprapiObjectHeaderIDL `idl:"name:objectHeader" json:"object_header"`
	PreferredMaxLength  uint32                       `idl:"name:dwPreferedMaxLen" json:"preferred_max_length"`
	EntriesRead         uint32                       `idl:"name:lpdwEntriesRead" json:"entries_read"`
	LpdNumTotalElements uint32                       `idl:"name:lpdNumTotalElements" json:"lpd_num_total_elements"`
	RASConections       []*rrasm.RASConnectionExIDL  `idl:"name:pRasConections;size_is:(, lpdwEntriesRead)" json:"ras_conections"`
	Resume              uint32                       `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	Return              uint32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_RASAdminConnectionEnumExOperation) OpNum() int { return 45 }

func (o *xxx_RASAdminConnectionEnumExOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminConnectionEnumEx"
}

func (o *xxx_RASAdminConnectionEnumExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionEnumExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
			if err := (&rrasm.MprapiObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_RASAdminConnectionEnumExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// objectHeader {in} (1:{alias=PMPRAPI_OBJECT_HEADER_IDL,pointer=ref}*(1))(2:{alias=MPRAPI_OBJECT_HEADER_IDL}(struct))
	{
		if o.ObjectHeader == nil {
			o.ObjectHeader = &rrasm.MprapiObjectHeaderIDL{}
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

func (o *xxx_RASAdminConnectionEnumExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
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

func (o *xxx_RASAdminConnectionEnumExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RASAdminConnectionEnumExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RASAdminConnectionEnumExRequest structure represents the RRasAdminConnectionEnumEx operation request
type RASAdminConnectionEnumExRequest struct {
	ObjectHeader       *rrasm.MprapiObjectHeaderIDL `idl:"name:objectHeader" json:"object_header"`
	PreferredMaxLength uint32                       `idl:"name:dwPreferedMaxLen" json:"preferred_max_length"`
	Resume             uint32                       `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
}

func (o *RASAdminConnectionEnumExRequest) xxx_ToOp(ctx context.Context, op *xxx_RASAdminConnectionEnumExOperation) *xxx_RASAdminConnectionEnumExOperation {
	if op == nil {
		op = &xxx_RASAdminConnectionEnumExOperation{}
	}
	if o == nil {
		return op
	}
	op.ObjectHeader = o.ObjectHeader
	op.PreferredMaxLength = o.PreferredMaxLength
	op.Resume = o.Resume
	return op
}

func (o *RASAdminConnectionEnumExRequest) xxx_FromOp(ctx context.Context, op *xxx_RASAdminConnectionEnumExOperation) {
	if o == nil {
		return
	}
	o.ObjectHeader = op.ObjectHeader
	o.PreferredMaxLength = op.PreferredMaxLength
	o.Resume = op.Resume
}
func (o *RASAdminConnectionEnumExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASAdminConnectionEnumExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminConnectionEnumExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASAdminConnectionEnumExResponse structure represents the RRasAdminConnectionEnumEx operation response
type RASAdminConnectionEnumExResponse struct {
	EntriesRead         uint32                      `idl:"name:lpdwEntriesRead" json:"entries_read"`
	LpdNumTotalElements uint32                      `idl:"name:lpdNumTotalElements" json:"lpd_num_total_elements"`
	RASConections       []*rrasm.RASConnectionExIDL `idl:"name:pRasConections;size_is:(, lpdwEntriesRead)" json:"ras_conections"`
	Resume              uint32                      `idl:"name:lpdwResumeHandle;pointer:unique" json:"resume"`
	// Return: The RRasAdminConnectionEnumEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASAdminConnectionEnumExResponse) xxx_ToOp(ctx context.Context, op *xxx_RASAdminConnectionEnumExOperation) *xxx_RASAdminConnectionEnumExOperation {
	if op == nil {
		op = &xxx_RASAdminConnectionEnumExOperation{}
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

func (o *RASAdminConnectionEnumExResponse) xxx_FromOp(ctx context.Context, op *xxx_RASAdminConnectionEnumExOperation) {
	if o == nil {
		return
	}
	o.EntriesRead = op.EntriesRead
	o.LpdNumTotalElements = op.LpdNumTotalElements
	o.RASConections = op.RASConections
	o.Resume = op.Resume
	o.Return = op.Return
}
func (o *RASAdminConnectionEnumExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASAdminConnectionEnumExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminConnectionEnumExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RASAdminConnectionGetInfoExOperation structure represents the RRasAdminConnectionGetInfoEx operation
type xxx_RASAdminConnectionGetInfoExOperation struct {
	HDimConnection uint32                       `idl:"name:hDimConnection" json:"h_dim_connection"`
	ObjectHeader   *rrasm.MprapiObjectHeaderIDL `idl:"name:objectHeader" json:"object_header"`
	RASConnection  *rrasm.RASConnectionExIDL    `idl:"name:pRasConnection" json:"ras_connection"`
	Return         uint32                       `idl:"name:Return" json:"return"`
}

func (o *xxx_RASAdminConnectionGetInfoExOperation) OpNum() int { return 46 }

func (o *xxx_RASAdminConnectionGetInfoExOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminConnectionGetInfoEx"
}

func (o *xxx_RASAdminConnectionGetInfoExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionGetInfoExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.HDimConnection); err != nil {
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
			if err := (&rrasm.MprapiObjectHeaderIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionGetInfoExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.HDimConnection); err != nil {
			return err
		}
	}
	// objectHeader {in} (1:{alias=PMPRAPI_OBJECT_HEADER_IDL,pointer=ref}*(1))(2:{alias=MPRAPI_OBJECT_HEADER_IDL}(struct))
	{
		if o.ObjectHeader == nil {
			o.ObjectHeader = &rrasm.MprapiObjectHeaderIDL{}
		}
		if err := o.ObjectHeader.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionGetInfoExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminConnectionGetInfoExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RASAdminConnectionGetInfoExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
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

// RASAdminConnectionGetInfoExRequest structure represents the RRasAdminConnectionGetInfoEx operation request
type RASAdminConnectionGetInfoExRequest struct {
	HDimConnection uint32                       `idl:"name:hDimConnection" json:"h_dim_connection"`
	ObjectHeader   *rrasm.MprapiObjectHeaderIDL `idl:"name:objectHeader" json:"object_header"`
}

func (o *RASAdminConnectionGetInfoExRequest) xxx_ToOp(ctx context.Context, op *xxx_RASAdminConnectionGetInfoExOperation) *xxx_RASAdminConnectionGetInfoExOperation {
	if op == nil {
		op = &xxx_RASAdminConnectionGetInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.HDimConnection = o.HDimConnection
	op.ObjectHeader = o.ObjectHeader
	return op
}

func (o *RASAdminConnectionGetInfoExRequest) xxx_FromOp(ctx context.Context, op *xxx_RASAdminConnectionGetInfoExOperation) {
	if o == nil {
		return
	}
	o.HDimConnection = op.HDimConnection
	o.ObjectHeader = op.ObjectHeader
}
func (o *RASAdminConnectionGetInfoExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASAdminConnectionGetInfoExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminConnectionGetInfoExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASAdminConnectionGetInfoExResponse structure represents the RRasAdminConnectionGetInfoEx operation response
type RASAdminConnectionGetInfoExResponse struct {
	RASConnection *rrasm.RASConnectionExIDL `idl:"name:pRasConnection" json:"ras_connection"`
	// Return: The RRasAdminConnectionGetInfoEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASAdminConnectionGetInfoExResponse) xxx_ToOp(ctx context.Context, op *xxx_RASAdminConnectionGetInfoExOperation) *xxx_RASAdminConnectionGetInfoExOperation {
	if op == nil {
		op = &xxx_RASAdminConnectionGetInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.RASConnection = o.RASConnection
	op.Return = o.Return
	return op
}

func (o *RASAdminConnectionGetInfoExResponse) xxx_FromOp(ctx context.Context, op *xxx_RASAdminConnectionGetInfoExOperation) {
	if o == nil {
		return
	}
	o.RASConnection = op.RASConnection
	o.Return = op.Return
}
func (o *RASAdminConnectionGetInfoExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASAdminConnectionGetInfoExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminConnectionGetInfoExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MprAdminServerSetInfoExOperation structure represents the RMprAdminServerSetInfoEx operation
type xxx_MprAdminServerSetInfoExOperation struct {
	ServerConfig *rrasm.MprServerSetConfigExIDL `idl:"name:pServerConfig" json:"server_config"`
	Return       uint32                         `idl:"name:Return" json:"return"`
}

func (o *xxx_MprAdminServerSetInfoExOperation) OpNum() int { return 47 }

func (o *xxx_MprAdminServerSetInfoExOperation) OpName() string {
	return "/dimsvc/v0/RMprAdminServerSetInfoEx"
}

func (o *xxx_MprAdminServerSetInfoExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MprAdminServerSetInfoExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
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
			if err := (&rrasm.MprServerSetConfigExIDL{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MprAdminServerSetInfoExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pServerConfig {in} (1:{alias=PMPR_SERVER_SET_CONFIG_EX_IDL,pointer=ref}*(1))(2:{alias=MPR_SERVER_SET_CONFIG_EX_IDL}(union))
	{
		if o.ServerConfig == nil {
			o.ServerConfig = &rrasm.MprServerSetConfigExIDL{}
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

func (o *xxx_MprAdminServerSetInfoExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MprAdminServerSetInfoExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_MprAdminServerSetInfoExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MprAdminServerSetInfoExRequest structure represents the RMprAdminServerSetInfoEx operation request
type MprAdminServerSetInfoExRequest struct {
	ServerConfig *rrasm.MprServerSetConfigExIDL `idl:"name:pServerConfig" json:"server_config"`
}

func (o *MprAdminServerSetInfoExRequest) xxx_ToOp(ctx context.Context, op *xxx_MprAdminServerSetInfoExOperation) *xxx_MprAdminServerSetInfoExOperation {
	if op == nil {
		op = &xxx_MprAdminServerSetInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerConfig = o.ServerConfig
	return op
}

func (o *MprAdminServerSetInfoExRequest) xxx_FromOp(ctx context.Context, op *xxx_MprAdminServerSetInfoExOperation) {
	if o == nil {
		return
	}
	o.ServerConfig = op.ServerConfig
}
func (o *MprAdminServerSetInfoExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MprAdminServerSetInfoExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MprAdminServerSetInfoExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MprAdminServerSetInfoExResponse structure represents the RMprAdminServerSetInfoEx operation response
type MprAdminServerSetInfoExResponse struct {
	// Return: The RMprAdminServerSetInfoEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MprAdminServerSetInfoExResponse) xxx_ToOp(ctx context.Context, op *xxx_MprAdminServerSetInfoExOperation) *xxx_MprAdminServerSetInfoExOperation {
	if op == nil {
		op = &xxx_MprAdminServerSetInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *MprAdminServerSetInfoExResponse) xxx_FromOp(ctx context.Context, op *xxx_MprAdminServerSetInfoExOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *MprAdminServerSetInfoExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MprAdminServerSetInfoExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MprAdminServerSetInfoExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RASAdminUpdateConnectionOperation structure represents the RRasAdminUpdateConnection operation
type xxx_RASAdminUpdateConnectionOperation struct {
	HDimConnection uint32                        `idl:"name:hDimConnection" json:"h_dim_connection"`
	ServerConfig   *rrasm.RASUpdateConnectionIDL `idl:"name:pServerConfig" json:"server_config"`
	Return         uint32                        `idl:"name:Return" json:"return"`
}

func (o *xxx_RASAdminUpdateConnectionOperation) OpNum() int { return 48 }

func (o *xxx_RASAdminUpdateConnectionOperation) OpName() string {
	return "/dimsvc/v0/RRasAdminUpdateConnection"
}

func (o *xxx_RASAdminUpdateConnectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminUpdateConnectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.HDimConnection); err != nil {
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

func (o *xxx_RASAdminUpdateConnectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hDimConnection {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.HDimConnection); err != nil {
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

func (o *xxx_RASAdminUpdateConnectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RASAdminUpdateConnectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_RASAdminUpdateConnectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RASAdminUpdateConnectionRequest structure represents the RRasAdminUpdateConnection operation request
type RASAdminUpdateConnectionRequest struct {
	HDimConnection uint32                        `idl:"name:hDimConnection" json:"h_dim_connection"`
	ServerConfig   *rrasm.RASUpdateConnectionIDL `idl:"name:pServerConfig" json:"server_config"`
}

func (o *RASAdminUpdateConnectionRequest) xxx_ToOp(ctx context.Context, op *xxx_RASAdminUpdateConnectionOperation) *xxx_RASAdminUpdateConnectionOperation {
	if op == nil {
		op = &xxx_RASAdminUpdateConnectionOperation{}
	}
	if o == nil {
		return op
	}
	op.HDimConnection = o.HDimConnection
	op.ServerConfig = o.ServerConfig
	return op
}

func (o *RASAdminUpdateConnectionRequest) xxx_FromOp(ctx context.Context, op *xxx_RASAdminUpdateConnectionOperation) {
	if o == nil {
		return
	}
	o.HDimConnection = op.HDimConnection
	o.ServerConfig = op.ServerConfig
}
func (o *RASAdminUpdateConnectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RASAdminUpdateConnectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminUpdateConnectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RASAdminUpdateConnectionResponse structure represents the RRasAdminUpdateConnection operation response
type RASAdminUpdateConnectionResponse struct {
	// Return: The RRasAdminUpdateConnection return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RASAdminUpdateConnectionResponse) xxx_ToOp(ctx context.Context, op *xxx_RASAdminUpdateConnectionOperation) *xxx_RASAdminUpdateConnectionOperation {
	if op == nil {
		op = &xxx_RASAdminUpdateConnectionOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RASAdminUpdateConnectionResponse) xxx_FromOp(ctx context.Context, op *xxx_RASAdminUpdateConnectionOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RASAdminUpdateConnectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RASAdminUpdateConnectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RASAdminUpdateConnectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RouterInterfaceSetCredentialsLocalOperation structure represents the RRouterInterfaceSetCredentialsLocal operation
type xxx_RouterInterfaceSetCredentialsLocalOperation struct {
	LpwsInterfaceName string `idl:"name:lpwsInterfaceName;string" json:"lpws_interface_name"`
	LpwsUserName      string `idl:"name:lpwsUserName;string" json:"lpws_user_name"`
	LpwsDomainName    string `idl:"name:lpwsDomainName;string" json:"lpws_domain_name"`
	LpwsPassword      string `idl:"name:lpwsPassword;string" json:"lpws_password"`
	Return            uint32 `idl:"name:Return" json:"return"`
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
		if err := ndr.WriteUTF16NString(ctx, w, o.LpwsInterfaceName); err != nil {
			return err
		}
	}
	// lpwsUserName {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.LpwsUserName); err != nil {
			return err
		}
	}
	// lpwsDomainName {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.LpwsDomainName); err != nil {
			return err
		}
	}
	// lpwsPassword {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.LpwsPassword); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceSetCredentialsLocalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpwsInterfaceName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.LpwsInterfaceName); err != nil {
			return err
		}
	}
	// lpwsUserName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.LpwsUserName); err != nil {
			return err
		}
	}
	// lpwsDomainName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.LpwsDomainName); err != nil {
			return err
		}
	}
	// lpwsPassword {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.LpwsPassword); err != nil {
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
	LpwsInterfaceName string `idl:"name:lpwsInterfaceName;string" json:"lpws_interface_name"`
	LpwsUserName      string `idl:"name:lpwsUserName;string" json:"lpws_user_name"`
	LpwsDomainName    string `idl:"name:lpwsDomainName;string" json:"lpws_domain_name"`
	LpwsPassword      string `idl:"name:lpwsPassword;string" json:"lpws_password"`
}

func (o *RouterInterfaceSetCredentialsLocalRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceSetCredentialsLocalOperation) *xxx_RouterInterfaceSetCredentialsLocalOperation {
	if op == nil {
		op = &xxx_RouterInterfaceSetCredentialsLocalOperation{}
	}
	if o == nil {
		return op
	}
	op.LpwsInterfaceName = o.LpwsInterfaceName
	op.LpwsUserName = o.LpwsUserName
	op.LpwsDomainName = o.LpwsDomainName
	op.LpwsPassword = o.LpwsPassword
	return op
}

func (o *RouterInterfaceSetCredentialsLocalRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceSetCredentialsLocalOperation) {
	if o == nil {
		return
	}
	o.LpwsInterfaceName = op.LpwsInterfaceName
	o.LpwsUserName = op.LpwsUserName
	o.LpwsDomainName = op.LpwsDomainName
	o.LpwsPassword = op.LpwsPassword
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
	LpwsInterfaceName string `idl:"name:lpwsInterfaceName;string" json:"lpws_interface_name"`
	LpwsUserName      string `idl:"name:lpwsUserName;string" json:"lpws_user_name"`
	LpwsDomainName    string `idl:"name:lpwsDomainName;string" json:"lpws_domain_name"`
	LpwsPassword      string `idl:"name:lpwsPassword;string" json:"lpws_password"`
	Return            uint32 `idl:"name:Return" json:"return"`
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
		if err := ndr.WriteUTF16NString(ctx, w, o.LpwsInterfaceName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RouterInterfaceGetCredentialsLocalOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpwsInterfaceName {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.LpwsInterfaceName); err != nil {
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
		if o.LpwsUserName != "" {
			_ptr_lpwsUserName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.LpwsUserName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.LpwsUserName, _ptr_lpwsUserName); err != nil {
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
		if o.LpwsDomainName != "" {
			_ptr_lpwsDomainName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.LpwsDomainName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.LpwsDomainName, _ptr_lpwsDomainName); err != nil {
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
		if o.LpwsPassword != "" {
			_ptr_lpwsPassword := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.LpwsPassword); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.LpwsPassword, _ptr_lpwsPassword); err != nil {
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
			if err := ndr.ReadUTF16NString(ctx, w, &o.LpwsUserName); err != nil {
				return err
			}
			return nil
		})
		_s_lpwsUserName := func(ptr interface{}) { o.LpwsUserName = *ptr.(*string) }
		if err := w.ReadPointer(&o.LpwsUserName, _s_lpwsUserName, _ptr_lpwsUserName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpwsDomainName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpwsDomainName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.LpwsDomainName); err != nil {
				return err
			}
			return nil
		})
		_s_lpwsDomainName := func(ptr interface{}) { o.LpwsDomainName = *ptr.(*string) }
		if err := w.ReadPointer(&o.LpwsDomainName, _s_lpwsDomainName, _ptr_lpwsDomainName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpwsPassword {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpwsPassword := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.LpwsPassword); err != nil {
				return err
			}
			return nil
		})
		_s_lpwsPassword := func(ptr interface{}) { o.LpwsPassword = *ptr.(*string) }
		if err := w.ReadPointer(&o.LpwsPassword, _s_lpwsPassword, _ptr_lpwsPassword); err != nil {
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
	LpwsInterfaceName string `idl:"name:lpwsInterfaceName;string" json:"lpws_interface_name"`
}

func (o *RouterInterfaceGetCredentialsLocalRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetCredentialsLocalOperation) *xxx_RouterInterfaceGetCredentialsLocalOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetCredentialsLocalOperation{}
	}
	if o == nil {
		return op
	}
	op.LpwsInterfaceName = o.LpwsInterfaceName
	return op
}

func (o *RouterInterfaceGetCredentialsLocalRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetCredentialsLocalOperation) {
	if o == nil {
		return
	}
	o.LpwsInterfaceName = op.LpwsInterfaceName
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
	LpwsUserName   string `idl:"name:lpwsUserName;string" json:"lpws_user_name"`
	LpwsDomainName string `idl:"name:lpwsDomainName;string" json:"lpws_domain_name"`
	LpwsPassword   string `idl:"name:lpwsPassword;string" json:"lpws_password"`
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
	op.LpwsUserName = o.LpwsUserName
	op.LpwsDomainName = o.LpwsDomainName
	op.LpwsPassword = o.LpwsPassword
	op.Return = o.Return
	return op
}

func (o *RouterInterfaceGetCredentialsLocalResponse) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetCredentialsLocalOperation) {
	if o == nil {
		return
	}
	o.LpwsUserName = op.LpwsUserName
	o.LpwsDomainName = op.LpwsDomainName
	o.LpwsPassword = op.LpwsPassword
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
	HInterface            uint32                             `idl:"name:hInterface" json:"h_interface"`
	InterfaceCustomConfig *rrasm.MprInterfaceCustominfoexIDL `idl:"name:pIfCustomConfig" json:"interface_custom_config"`
	Return                uint32                             `idl:"name:Return" json:"return"`
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
		if err := w.WriteData(o.HInterface); err != nil {
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
			if err := (&rrasm.MprInterfaceCustominfoexIDL{}).MarshalNDR(ctx, w); err != nil {
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
		if err := w.ReadData(&o.HInterface); err != nil {
			return err
		}
	}
	// pIfCustomConfig {in, out} (1:{alias=PMPR_IF_CUSTOMINFOEX_IDL,pointer=ref}*(1))(2:{alias=MPR_IF_CUSTOMINFOEX_IDL}(union))
	{
		if o.InterfaceCustomConfig == nil {
			o.InterfaceCustomConfig = &rrasm.MprInterfaceCustominfoexIDL{}
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
			if err := (&rrasm.MprInterfaceCustominfoexIDL{}).MarshalNDR(ctx, w); err != nil {
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
			o.InterfaceCustomConfig = &rrasm.MprInterfaceCustominfoexIDL{}
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
	HInterface            uint32                             `idl:"name:hInterface" json:"h_interface"`
	InterfaceCustomConfig *rrasm.MprInterfaceCustominfoexIDL `idl:"name:pIfCustomConfig" json:"interface_custom_config"`
}

func (o *RouterInterfaceGetCustomInfoExRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceGetCustomInfoExOperation) *xxx_RouterInterfaceGetCustomInfoExOperation {
	if op == nil {
		op = &xxx_RouterInterfaceGetCustomInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.HInterface = o.HInterface
	op.InterfaceCustomConfig = o.InterfaceCustomConfig
	return op
}

func (o *RouterInterfaceGetCustomInfoExRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceGetCustomInfoExOperation) {
	if o == nil {
		return
	}
	o.HInterface = op.HInterface
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
	InterfaceCustomConfig *rrasm.MprInterfaceCustominfoexIDL `idl:"name:pIfCustomConfig" json:"interface_custom_config"`
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
	HInterface            uint32                             `idl:"name:hInterface" json:"h_interface"`
	InterfaceCustomConfig *rrasm.MprInterfaceCustominfoexIDL `idl:"name:pIfCustomConfig" json:"interface_custom_config"`
	Return                uint32                             `idl:"name:Return" json:"return"`
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
		if err := w.WriteData(o.HInterface); err != nil {
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
			if err := (&rrasm.MprInterfaceCustominfoexIDL{}).MarshalNDR(ctx, w); err != nil {
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
		if err := w.ReadData(&o.HInterface); err != nil {
			return err
		}
	}
	// pIfCustomConfig {in, out} (1:{alias=PMPR_IF_CUSTOMINFOEX_IDL,pointer=ref}*(1))(2:{alias=MPR_IF_CUSTOMINFOEX_IDL}(union))
	{
		if o.InterfaceCustomConfig == nil {
			o.InterfaceCustomConfig = &rrasm.MprInterfaceCustominfoexIDL{}
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
			if err := (&rrasm.MprInterfaceCustominfoexIDL{}).MarshalNDR(ctx, w); err != nil {
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
			o.InterfaceCustomConfig = &rrasm.MprInterfaceCustominfoexIDL{}
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
	HInterface            uint32                             `idl:"name:hInterface" json:"h_interface"`
	InterfaceCustomConfig *rrasm.MprInterfaceCustominfoexIDL `idl:"name:pIfCustomConfig" json:"interface_custom_config"`
}

func (o *RouterInterfaceSetCustomInfoExRequest) xxx_ToOp(ctx context.Context, op *xxx_RouterInterfaceSetCustomInfoExOperation) *xxx_RouterInterfaceSetCustomInfoExOperation {
	if op == nil {
		op = &xxx_RouterInterfaceSetCustomInfoExOperation{}
	}
	if o == nil {
		return op
	}
	op.HInterface = o.HInterface
	op.InterfaceCustomConfig = o.InterfaceCustomConfig
	return op
}

func (o *RouterInterfaceSetCustomInfoExRequest) xxx_FromOp(ctx context.Context, op *xxx_RouterInterfaceSetCustomInfoExOperation) {
	if o == nil {
		return
	}
	o.HInterface = op.HInterface
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
	InterfaceCustomConfig *rrasm.MprInterfaceCustominfoexIDL `idl:"name:pIfCustomConfig" json:"interface_custom_config"`
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
