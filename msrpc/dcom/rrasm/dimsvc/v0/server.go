package dimsvc

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// dimsvc server interface.
type DimsvcServer interface {
	ServerGetInfo(context.Context, *ServerGetInfoRequest) (*ServerGetInfoResponse, error)

	ConnectionEnum(context.Context, *ConnectionEnumRequest) (*ConnectionEnumResponse, error)

	ConnectionGetInfo(context.Context, *ConnectionGetInfoRequest) (*ConnectionGetInfoResponse, error)

	ConnectionClearStats(context.Context, *ConnectionClearStatsRequest) (*ConnectionClearStatsResponse, error)

	PortEnum(context.Context, *PortEnumRequest) (*PortEnumResponse, error)

	PortGetInfo(context.Context, *PortGetInfoRequest) (*PortGetInfoResponse, error)

	PortClearStats(context.Context, *PortClearStatsRequest) (*PortClearStatsResponse, error)

	PortReset(context.Context, *PortResetRequest) (*PortResetResponse, error)

	PortDisconnect(context.Context, *PortDisconnectRequest) (*PortDisconnectResponse, error)

	RouterInterfaceTransportSetGlobalInfo(context.Context, *RouterInterfaceTransportSetGlobalInfoRequest) (*RouterInterfaceTransportSetGlobalInfoResponse, error)

	RouterInterfaceTransportGetGlobalInfo(context.Context, *RouterInterfaceTransportGetGlobalInfoRequest) (*RouterInterfaceTransportGetGlobalInfoResponse, error)

	RouterInterfaceGetHandle(context.Context, *RouterInterfaceGetHandleRequest) (*RouterInterfaceGetHandleResponse, error)

	RouterInterfaceCreate(context.Context, *RouterInterfaceCreateRequest) (*RouterInterfaceCreateResponse, error)

	RouterInterfaceGetInfo(context.Context, *RouterInterfaceGetInfoRequest) (*RouterInterfaceGetInfoResponse, error)

	RouterInterfaceSetInfo(context.Context, *RouterInterfaceSetInfoRequest) (*RouterInterfaceSetInfoResponse, error)

	RouterInterfaceDelete(context.Context, *RouterInterfaceDeleteRequest) (*RouterInterfaceDeleteResponse, error)

	RouterInterfaceTransportRemove(context.Context, *RouterInterfaceTransportRemoveRequest) (*RouterInterfaceTransportRemoveResponse, error)

	RouterInterfaceTransportAdd(context.Context, *RouterInterfaceTransportAddRequest) (*RouterInterfaceTransportAddResponse, error)

	RouterInterfaceTransportGetInfo(context.Context, *RouterInterfaceTransportGetInfoRequest) (*RouterInterfaceTransportGetInfoResponse, error)

	RouterInterfaceTransportSetInfo(context.Context, *RouterInterfaceTransportSetInfoRequest) (*RouterInterfaceTransportSetInfoResponse, error)

	RouterInterfaceEnum(context.Context, *RouterInterfaceEnumRequest) (*RouterInterfaceEnumResponse, error)

	RouterInterfaceConnect(context.Context, *RouterInterfaceConnectRequest) (*RouterInterfaceConnectResponse, error)

	RouterInterfaceDisconnect(context.Context, *RouterInterfaceDisconnectRequest) (*RouterInterfaceDisconnectResponse, error)

	RouterInterfaceUpdateRoutes(context.Context, *RouterInterfaceUpdateRoutesRequest) (*RouterInterfaceUpdateRoutesResponse, error)

	RouterInterfaceQueryUpdateResult(context.Context, *RouterInterfaceQueryUpdateResultRequest) (*RouterInterfaceQueryUpdateResultResponse, error)

	RouterInterfaceUpdatePhonebookInfo(context.Context, *RouterInterfaceUpdatePhonebookInfoRequest) (*RouterInterfaceUpdatePhonebookInfoResponse, error)

	MIBEntryCreate(context.Context, *MIBEntryCreateRequest) (*MIBEntryCreateResponse, error)

	MIBEntryDelete(context.Context, *MIBEntryDeleteRequest) (*MIBEntryDeleteResponse, error)

	MIBEntrySet(context.Context, *MIBEntrySetRequest) (*MIBEntrySetResponse, error)

	MIBEntryGet(context.Context, *MIBEntryGetRequest) (*MIBEntryGetResponse, error)

	MIBEntryGetFirst(context.Context, *MIBEntryGetFirstRequest) (*MIBEntryGetFirstResponse, error)

	MIBEntryGetNext(context.Context, *MIBEntryGetNextRequest) (*MIBEntryGetNextResponse, error)

	MIBGetTrapInfo(context.Context, *MIBGetTrapInfoRequest) (*MIBGetTrapInfoResponse, error)

	MIBSetTrapInfo(context.Context, *MIBSetTrapInfoRequest) (*MIBSetTrapInfoResponse, error)

	ConnectionNotification(context.Context, *ConnectionNotificationRequest) (*ConnectionNotificationResponse, error)

	SendUserMessage(context.Context, *SendUserMessageRequest) (*SendUserMessageResponse, error)

	RouterDeviceEnum(context.Context, *RouterDeviceEnumRequest) (*RouterDeviceEnumResponse, error)

	RouterInterfaceTransportCreate(context.Context, *RouterInterfaceTransportCreateRequest) (*RouterInterfaceTransportCreateResponse, error)

	RouterInterfaceDeviceGetInfo(context.Context, *RouterInterfaceDeviceGetInfoRequest) (*RouterInterfaceDeviceGetInfoResponse, error)

	RouterInterfaceDeviceSetInfo(context.Context, *RouterInterfaceDeviceSetInfoRequest) (*RouterInterfaceDeviceSetInfoResponse, error)

	RouterInterfaceSetCredentialsEx(context.Context, *RouterInterfaceSetCredentialsExRequest) (*RouterInterfaceSetCredentialsExResponse, error)

	RouterInterfaceGetCredentialsEx(context.Context, *RouterInterfaceGetCredentialsExRequest) (*RouterInterfaceGetCredentialsExResponse, error)

	ConnectionRemoveQuarantine(context.Context, *ConnectionRemoveQuarantineRequest) (*ConnectionRemoveQuarantineResponse, error)

	ServerSetInfo(context.Context, *ServerSetInfoRequest) (*ServerSetInfoResponse, error)

	ServerGetInfoEx(context.Context, *ServerGetInfoExRequest) (*ServerGetInfoExResponse, error)

	ConnectionEnumEx(context.Context, *ConnectionEnumExRequest) (*ConnectionEnumExResponse, error)

	ConnectionGetInfoEx(context.Context, *ConnectionGetInfoExRequest) (*ConnectionGetInfoExResponse, error)

	ServerSetInfoEx(context.Context, *ServerSetInfoExRequest) (*ServerSetInfoExResponse, error)

	UpdateConnection(context.Context, *UpdateConnectionRequest) (*UpdateConnectionResponse, error)

	RouterInterfaceSetCredentialsLocal(context.Context, *RouterInterfaceSetCredentialsLocalRequest) (*RouterInterfaceSetCredentialsLocalResponse, error)

	RouterInterfaceGetCredentialsLocal(context.Context, *RouterInterfaceGetCredentialsLocalRequest) (*RouterInterfaceGetCredentialsLocalResponse, error)

	RouterInterfaceGetCustomInfoEx(context.Context, *RouterInterfaceGetCustomInfoExRequest) (*RouterInterfaceGetCustomInfoExResponse, error)

	RouterInterfaceSetCustomInfoEx(context.Context, *RouterInterfaceSetCustomInfoExRequest) (*RouterInterfaceSetCustomInfoExResponse, error)
}

func RegisterDimsvcServer(conn dcerpc.Conn, o DimsvcServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDimsvcServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DimsvcSyntaxV0_0))...)
}

func NewDimsvcServerHandle(o DimsvcServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DimsvcServerHandle(ctx, o, opNum, r)
	}
}

func DimsvcServerHandle(ctx context.Context, o DimsvcServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // RMprAdminServerGetInfo
		op := &xxx_ServerGetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ServerGetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ServerGetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // RRasAdminConnectionEnum
		op := &xxx_ConnectionEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectionEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConnectionEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // RRasAdminConnectionGetInfo
		op := &xxx_ConnectionGetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectionGetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConnectionGetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // RRasAdminConnectionClearStats
		op := &xxx_ConnectionClearStatsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectionClearStatsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConnectionClearStats(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // RRasAdminPortEnum
		op := &xxx_PortEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PortEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PortEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // RRasAdminPortGetInfo
		op := &xxx_PortGetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PortGetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PortGetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // RRasAdminPortClearStats
		op := &xxx_PortClearStatsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PortClearStatsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PortClearStats(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // RRasAdminPortReset
		op := &xxx_PortResetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PortResetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PortReset(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // RRasAdminPortDisconnect
		op := &xxx_PortDisconnectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PortDisconnectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PortDisconnect(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // RRouterInterfaceTransportSetGlobalInfo
		op := &xxx_RouterInterfaceTransportSetGlobalInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceTransportSetGlobalInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceTransportSetGlobalInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // RRouterInterfaceTransportGetGlobalInfo
		op := &xxx_RouterInterfaceTransportGetGlobalInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceTransportGetGlobalInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceTransportGetGlobalInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // RRouterInterfaceGetHandle
		op := &xxx_RouterInterfaceGetHandleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceGetHandleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceGetHandle(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // RRouterInterfaceCreate
		op := &xxx_RouterInterfaceCreateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceCreateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceCreate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // RRouterInterfaceGetInfo
		op := &xxx_RouterInterfaceGetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceGetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceGetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // RRouterInterfaceSetInfo
		op := &xxx_RouterInterfaceSetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceSetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceSetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // RRouterInterfaceDelete
		op := &xxx_RouterInterfaceDeleteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceDeleteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceDelete(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // RRouterInterfaceTransportRemove
		op := &xxx_RouterInterfaceTransportRemoveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceTransportRemoveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceTransportRemove(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // RRouterInterfaceTransportAdd
		op := &xxx_RouterInterfaceTransportAddOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceTransportAddRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceTransportAdd(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // RRouterInterfaceTransportGetInfo
		op := &xxx_RouterInterfaceTransportGetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceTransportGetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceTransportGetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // RRouterInterfaceTransportSetInfo
		op := &xxx_RouterInterfaceTransportSetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceTransportSetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceTransportSetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // RRouterInterfaceEnum
		op := &xxx_RouterInterfaceEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // RRouterInterfaceConnect
		op := &xxx_RouterInterfaceConnectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceConnectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceConnect(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // RRouterInterfaceDisconnect
		op := &xxx_RouterInterfaceDisconnectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceDisconnectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceDisconnect(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // RRouterInterfaceUpdateRoutes
		op := &xxx_RouterInterfaceUpdateRoutesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceUpdateRoutesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceUpdateRoutes(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // RRouterInterfaceQueryUpdateResult
		op := &xxx_RouterInterfaceQueryUpdateResultOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceQueryUpdateResultRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceQueryUpdateResult(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // RRouterInterfaceUpdatePhonebookInfo
		op := &xxx_RouterInterfaceUpdatePhonebookInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceUpdatePhonebookInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceUpdatePhonebookInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // RMIBEntryCreate
		op := &xxx_MIBEntryCreateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MIBEntryCreateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MIBEntryCreate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // RMIBEntryDelete
		op := &xxx_MIBEntryDeleteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MIBEntryDeleteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MIBEntryDelete(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // RMIBEntrySet
		op := &xxx_MIBEntrySetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MIBEntrySetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MIBEntrySet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // RMIBEntryGet
		op := &xxx_MIBEntryGetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MIBEntryGetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MIBEntryGet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // RMIBEntryGetFirst
		op := &xxx_MIBEntryGetFirstOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MIBEntryGetFirstRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MIBEntryGetFirst(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // RMIBEntryGetNext
		op := &xxx_MIBEntryGetNextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MIBEntryGetNextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MIBEntryGetNext(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // RMIBGetTrapInfo
		op := &xxx_MIBGetTrapInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MIBGetTrapInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MIBGetTrapInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // RMIBSetTrapInfo
		op := &xxx_MIBSetTrapInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MIBSetTrapInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MIBSetTrapInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // RRasAdminConnectionNotification
		op := &xxx_ConnectionNotificationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectionNotificationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConnectionNotification(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // RRasAdminSendUserMessage
		op := &xxx_SendUserMessageOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SendUserMessageRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SendUserMessage(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // RRouterDeviceEnum
		op := &xxx_RouterDeviceEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterDeviceEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterDeviceEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // RRouterInterfaceTransportCreate
		op := &xxx_RouterInterfaceTransportCreateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceTransportCreateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceTransportCreate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // RRouterInterfaceDeviceGetInfo
		op := &xxx_RouterInterfaceDeviceGetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceDeviceGetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceDeviceGetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // RRouterInterfaceDeviceSetInfo
		op := &xxx_RouterInterfaceDeviceSetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceDeviceSetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceDeviceSetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // RRouterInterfaceSetCredentialsEx
		op := &xxx_RouterInterfaceSetCredentialsExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceSetCredentialsExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceSetCredentialsEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // RRouterInterfaceGetCredentialsEx
		op := &xxx_RouterInterfaceGetCredentialsExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceGetCredentialsExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceGetCredentialsEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 42: // RRasAdminConnectionRemoveQuarantine
		op := &xxx_ConnectionRemoveQuarantineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectionRemoveQuarantineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConnectionRemoveQuarantine(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 43: // RMprAdminServerSetInfo
		op := &xxx_ServerSetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ServerSetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ServerSetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // RMprAdminServerGetInfoEx
		op := &xxx_ServerGetInfoExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ServerGetInfoExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ServerGetInfoEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // RRasAdminConnectionEnumEx
		op := &xxx_ConnectionEnumExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectionEnumExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConnectionEnumEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // RRasAdminConnectionGetInfoEx
		op := &xxx_ConnectionGetInfoExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectionGetInfoExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConnectionGetInfoEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 47: // RMprAdminServerSetInfoEx
		op := &xxx_ServerSetInfoExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ServerSetInfoExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ServerSetInfoEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 48: // RRasAdminUpdateConnection
		op := &xxx_UpdateConnectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UpdateConnectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UpdateConnection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 49: // RRouterInterfaceSetCredentialsLocal
		op := &xxx_RouterInterfaceSetCredentialsLocalOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceSetCredentialsLocalRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceSetCredentialsLocal(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 50: // RRouterInterfaceGetCredentialsLocal
		op := &xxx_RouterInterfaceGetCredentialsLocalOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceGetCredentialsLocalRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceGetCredentialsLocal(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 51: // RRouterInterfaceGetCustomInfoEx
		op := &xxx_RouterInterfaceGetCustomInfoExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceGetCustomInfoExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceGetCustomInfoEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 52: // RRouterInterfaceSetCustomInfoEx
		op := &xxx_RouterInterfaceSetCustomInfoExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RouterInterfaceSetCustomInfoExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RouterInterfaceSetCustomInfoEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented dimsvc
type UnimplementedDimsvcServer struct {
}

func (UnimplementedDimsvcServer) ServerGetInfo(context.Context, *ServerGetInfoRequest) (*ServerGetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) ConnectionEnum(context.Context, *ConnectionEnumRequest) (*ConnectionEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) ConnectionGetInfo(context.Context, *ConnectionGetInfoRequest) (*ConnectionGetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) ConnectionClearStats(context.Context, *ConnectionClearStatsRequest) (*ConnectionClearStatsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) PortEnum(context.Context, *PortEnumRequest) (*PortEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) PortGetInfo(context.Context, *PortGetInfoRequest) (*PortGetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) PortClearStats(context.Context, *PortClearStatsRequest) (*PortClearStatsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) PortReset(context.Context, *PortResetRequest) (*PortResetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) PortDisconnect(context.Context, *PortDisconnectRequest) (*PortDisconnectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceTransportSetGlobalInfo(context.Context, *RouterInterfaceTransportSetGlobalInfoRequest) (*RouterInterfaceTransportSetGlobalInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceTransportGetGlobalInfo(context.Context, *RouterInterfaceTransportGetGlobalInfoRequest) (*RouterInterfaceTransportGetGlobalInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceGetHandle(context.Context, *RouterInterfaceGetHandleRequest) (*RouterInterfaceGetHandleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceCreate(context.Context, *RouterInterfaceCreateRequest) (*RouterInterfaceCreateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceGetInfo(context.Context, *RouterInterfaceGetInfoRequest) (*RouterInterfaceGetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceSetInfo(context.Context, *RouterInterfaceSetInfoRequest) (*RouterInterfaceSetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceDelete(context.Context, *RouterInterfaceDeleteRequest) (*RouterInterfaceDeleteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceTransportRemove(context.Context, *RouterInterfaceTransportRemoveRequest) (*RouterInterfaceTransportRemoveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceTransportAdd(context.Context, *RouterInterfaceTransportAddRequest) (*RouterInterfaceTransportAddResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceTransportGetInfo(context.Context, *RouterInterfaceTransportGetInfoRequest) (*RouterInterfaceTransportGetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceTransportSetInfo(context.Context, *RouterInterfaceTransportSetInfoRequest) (*RouterInterfaceTransportSetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceEnum(context.Context, *RouterInterfaceEnumRequest) (*RouterInterfaceEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceConnect(context.Context, *RouterInterfaceConnectRequest) (*RouterInterfaceConnectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceDisconnect(context.Context, *RouterInterfaceDisconnectRequest) (*RouterInterfaceDisconnectResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceUpdateRoutes(context.Context, *RouterInterfaceUpdateRoutesRequest) (*RouterInterfaceUpdateRoutesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceQueryUpdateResult(context.Context, *RouterInterfaceQueryUpdateResultRequest) (*RouterInterfaceQueryUpdateResultResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceUpdatePhonebookInfo(context.Context, *RouterInterfaceUpdatePhonebookInfoRequest) (*RouterInterfaceUpdatePhonebookInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) MIBEntryCreate(context.Context, *MIBEntryCreateRequest) (*MIBEntryCreateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) MIBEntryDelete(context.Context, *MIBEntryDeleteRequest) (*MIBEntryDeleteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) MIBEntrySet(context.Context, *MIBEntrySetRequest) (*MIBEntrySetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) MIBEntryGet(context.Context, *MIBEntryGetRequest) (*MIBEntryGetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) MIBEntryGetFirst(context.Context, *MIBEntryGetFirstRequest) (*MIBEntryGetFirstResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) MIBEntryGetNext(context.Context, *MIBEntryGetNextRequest) (*MIBEntryGetNextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) MIBGetTrapInfo(context.Context, *MIBGetTrapInfoRequest) (*MIBGetTrapInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) MIBSetTrapInfo(context.Context, *MIBSetTrapInfoRequest) (*MIBSetTrapInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) ConnectionNotification(context.Context, *ConnectionNotificationRequest) (*ConnectionNotificationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) SendUserMessage(context.Context, *SendUserMessageRequest) (*SendUserMessageResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterDeviceEnum(context.Context, *RouterDeviceEnumRequest) (*RouterDeviceEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceTransportCreate(context.Context, *RouterInterfaceTransportCreateRequest) (*RouterInterfaceTransportCreateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceDeviceGetInfo(context.Context, *RouterInterfaceDeviceGetInfoRequest) (*RouterInterfaceDeviceGetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceDeviceSetInfo(context.Context, *RouterInterfaceDeviceSetInfoRequest) (*RouterInterfaceDeviceSetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceSetCredentialsEx(context.Context, *RouterInterfaceSetCredentialsExRequest) (*RouterInterfaceSetCredentialsExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceGetCredentialsEx(context.Context, *RouterInterfaceGetCredentialsExRequest) (*RouterInterfaceGetCredentialsExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) ConnectionRemoveQuarantine(context.Context, *ConnectionRemoveQuarantineRequest) (*ConnectionRemoveQuarantineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) ServerSetInfo(context.Context, *ServerSetInfoRequest) (*ServerSetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) ServerGetInfoEx(context.Context, *ServerGetInfoExRequest) (*ServerGetInfoExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) ConnectionEnumEx(context.Context, *ConnectionEnumExRequest) (*ConnectionEnumExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) ConnectionGetInfoEx(context.Context, *ConnectionGetInfoExRequest) (*ConnectionGetInfoExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) ServerSetInfoEx(context.Context, *ServerSetInfoExRequest) (*ServerSetInfoExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) UpdateConnection(context.Context, *UpdateConnectionRequest) (*UpdateConnectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceSetCredentialsLocal(context.Context, *RouterInterfaceSetCredentialsLocalRequest) (*RouterInterfaceSetCredentialsLocalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceGetCredentialsLocal(context.Context, *RouterInterfaceGetCredentialsLocalRequest) (*RouterInterfaceGetCredentialsLocalResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceGetCustomInfoEx(context.Context, *RouterInterfaceGetCustomInfoExRequest) (*RouterInterfaceGetCustomInfoExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RouterInterfaceSetCustomInfoEx(context.Context, *RouterInterfaceSetCustomInfoExRequest) (*RouterInterfaceSetCustomInfoExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ DimsvcServer = (*UnimplementedDimsvcServer)(nil)
