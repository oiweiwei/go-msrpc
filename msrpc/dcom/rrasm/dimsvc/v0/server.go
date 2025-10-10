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
	MprAdminServerGetInfo(context.Context, *MprAdminServerGetInfoRequest) (*MprAdminServerGetInfoResponse, error)

	RASAdminConnectionEnum(context.Context, *RASAdminConnectionEnumRequest) (*RASAdminConnectionEnumResponse, error)

	RASAdminConnectionGetInfo(context.Context, *RASAdminConnectionGetInfoRequest) (*RASAdminConnectionGetInfoResponse, error)

	RASAdminConnectionClearStats(context.Context, *RASAdminConnectionClearStatsRequest) (*RASAdminConnectionClearStatsResponse, error)

	RASAdminPortEnum(context.Context, *RASAdminPortEnumRequest) (*RASAdminPortEnumResponse, error)

	RASAdminPortGetInfo(context.Context, *RASAdminPortGetInfoRequest) (*RASAdminPortGetInfoResponse, error)

	RASAdminPortClearStats(context.Context, *RASAdminPortClearStatsRequest) (*RASAdminPortClearStatsResponse, error)

	RASAdminPortReset(context.Context, *RASAdminPortResetRequest) (*RASAdminPortResetResponse, error)

	RASAdminPortDisconnect(context.Context, *RASAdminPortDisconnectRequest) (*RASAdminPortDisconnectResponse, error)

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

	RmibEntryCreate(context.Context, *RmibEntryCreateRequest) (*RmibEntryCreateResponse, error)

	RmibEntryDelete(context.Context, *RmibEntryDeleteRequest) (*RmibEntryDeleteResponse, error)

	RmibEntrySet(context.Context, *RmibEntrySetRequest) (*RmibEntrySetResponse, error)

	RmibEntryGet(context.Context, *RmibEntryGetRequest) (*RmibEntryGetResponse, error)

	RmibEntryGetFirst(context.Context, *RmibEntryGetFirstRequest) (*RmibEntryGetFirstResponse, error)

	RmibEntryGetNext(context.Context, *RmibEntryGetNextRequest) (*RmibEntryGetNextResponse, error)

	RmibGetTrapInfo(context.Context, *RmibGetTrapInfoRequest) (*RmibGetTrapInfoResponse, error)

	RmibSetTrapInfo(context.Context, *RmibSetTrapInfoRequest) (*RmibSetTrapInfoResponse, error)

	RASAdminConnectionNotification(context.Context, *RASAdminConnectionNotificationRequest) (*RASAdminConnectionNotificationResponse, error)

	RASAdminSendUserMessage(context.Context, *RASAdminSendUserMessageRequest) (*RASAdminSendUserMessageResponse, error)

	RouterDeviceEnum(context.Context, *RouterDeviceEnumRequest) (*RouterDeviceEnumResponse, error)

	RouterInterfaceTransportCreate(context.Context, *RouterInterfaceTransportCreateRequest) (*RouterInterfaceTransportCreateResponse, error)

	RouterInterfaceDeviceGetInfo(context.Context, *RouterInterfaceDeviceGetInfoRequest) (*RouterInterfaceDeviceGetInfoResponse, error)

	RouterInterfaceDeviceSetInfo(context.Context, *RouterInterfaceDeviceSetInfoRequest) (*RouterInterfaceDeviceSetInfoResponse, error)

	RouterInterfaceSetCredentialsEx(context.Context, *RouterInterfaceSetCredentialsExRequest) (*RouterInterfaceSetCredentialsExResponse, error)

	RouterInterfaceGetCredentialsEx(context.Context, *RouterInterfaceGetCredentialsExRequest) (*RouterInterfaceGetCredentialsExResponse, error)

	RASAdminConnectionRemoveQuarantine(context.Context, *RASAdminConnectionRemoveQuarantineRequest) (*RASAdminConnectionRemoveQuarantineResponse, error)

	MprAdminServerSetInfo(context.Context, *MprAdminServerSetInfoRequest) (*MprAdminServerSetInfoResponse, error)

	MprAdminServerGetInfoEx(context.Context, *MprAdminServerGetInfoExRequest) (*MprAdminServerGetInfoExResponse, error)

	RASAdminConnectionEnumEx(context.Context, *RASAdminConnectionEnumExRequest) (*RASAdminConnectionEnumExResponse, error)

	RASAdminConnectionGetInfoEx(context.Context, *RASAdminConnectionGetInfoExRequest) (*RASAdminConnectionGetInfoExResponse, error)

	MprAdminServerSetInfoEx(context.Context, *MprAdminServerSetInfoExRequest) (*MprAdminServerSetInfoExResponse, error)

	RASAdminUpdateConnection(context.Context, *RASAdminUpdateConnectionRequest) (*RASAdminUpdateConnectionResponse, error)

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
		op := &xxx_MprAdminServerGetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MprAdminServerGetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MprAdminServerGetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // RRasAdminConnectionEnum
		op := &xxx_RASAdminConnectionEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASAdminConnectionEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASAdminConnectionEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // RRasAdminConnectionGetInfo
		op := &xxx_RASAdminConnectionGetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASAdminConnectionGetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASAdminConnectionGetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // RRasAdminConnectionClearStats
		op := &xxx_RASAdminConnectionClearStatsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASAdminConnectionClearStatsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASAdminConnectionClearStats(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // RRasAdminPortEnum
		op := &xxx_RASAdminPortEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASAdminPortEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASAdminPortEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // RRasAdminPortGetInfo
		op := &xxx_RASAdminPortGetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASAdminPortGetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASAdminPortGetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // RRasAdminPortClearStats
		op := &xxx_RASAdminPortClearStatsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASAdminPortClearStatsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASAdminPortClearStats(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // RRasAdminPortReset
		op := &xxx_RASAdminPortResetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASAdminPortResetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASAdminPortReset(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // RRasAdminPortDisconnect
		op := &xxx_RASAdminPortDisconnectOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASAdminPortDisconnectRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASAdminPortDisconnect(ctx, req)
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
		op := &xxx_RmibEntryCreateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RmibEntryCreateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RmibEntryCreate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // RMIBEntryDelete
		op := &xxx_RmibEntryDeleteOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RmibEntryDeleteRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RmibEntryDelete(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // RMIBEntrySet
		op := &xxx_RmibEntrySetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RmibEntrySetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RmibEntrySet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // RMIBEntryGet
		op := &xxx_RmibEntryGetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RmibEntryGetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RmibEntryGet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // RMIBEntryGetFirst
		op := &xxx_RmibEntryGetFirstOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RmibEntryGetFirstRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RmibEntryGetFirst(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // RMIBEntryGetNext
		op := &xxx_RmibEntryGetNextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RmibEntryGetNextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RmibEntryGetNext(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // RMIBGetTrapInfo
		op := &xxx_RmibGetTrapInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RmibGetTrapInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RmibGetTrapInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // RMIBSetTrapInfo
		op := &xxx_RmibSetTrapInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RmibSetTrapInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RmibSetTrapInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // RRasAdminConnectionNotification
		op := &xxx_RASAdminConnectionNotificationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASAdminConnectionNotificationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASAdminConnectionNotification(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // RRasAdminSendUserMessage
		op := &xxx_RASAdminSendUserMessageOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASAdminSendUserMessageRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASAdminSendUserMessage(ctx, req)
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
		op := &xxx_RASAdminConnectionRemoveQuarantineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASAdminConnectionRemoveQuarantineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASAdminConnectionRemoveQuarantine(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 43: // RMprAdminServerSetInfo
		op := &xxx_MprAdminServerSetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MprAdminServerSetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MprAdminServerSetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // RMprAdminServerGetInfoEx
		op := &xxx_MprAdminServerGetInfoExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MprAdminServerGetInfoExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MprAdminServerGetInfoEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // RRasAdminConnectionEnumEx
		op := &xxx_RASAdminConnectionEnumExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASAdminConnectionEnumExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASAdminConnectionEnumEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // RRasAdminConnectionGetInfoEx
		op := &xxx_RASAdminConnectionGetInfoExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASAdminConnectionGetInfoExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASAdminConnectionGetInfoEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 47: // RMprAdminServerSetInfoEx
		op := &xxx_MprAdminServerSetInfoExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MprAdminServerSetInfoExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MprAdminServerSetInfoEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 48: // RRasAdminUpdateConnection
		op := &xxx_RASAdminUpdateConnectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RASAdminUpdateConnectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RASAdminUpdateConnection(ctx, req)
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

func (UnimplementedDimsvcServer) MprAdminServerGetInfo(context.Context, *MprAdminServerGetInfoRequest) (*MprAdminServerGetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RASAdminConnectionEnum(context.Context, *RASAdminConnectionEnumRequest) (*RASAdminConnectionEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RASAdminConnectionGetInfo(context.Context, *RASAdminConnectionGetInfoRequest) (*RASAdminConnectionGetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RASAdminConnectionClearStats(context.Context, *RASAdminConnectionClearStatsRequest) (*RASAdminConnectionClearStatsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RASAdminPortEnum(context.Context, *RASAdminPortEnumRequest) (*RASAdminPortEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RASAdminPortGetInfo(context.Context, *RASAdminPortGetInfoRequest) (*RASAdminPortGetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RASAdminPortClearStats(context.Context, *RASAdminPortClearStatsRequest) (*RASAdminPortClearStatsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RASAdminPortReset(context.Context, *RASAdminPortResetRequest) (*RASAdminPortResetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RASAdminPortDisconnect(context.Context, *RASAdminPortDisconnectRequest) (*RASAdminPortDisconnectResponse, error) {
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
func (UnimplementedDimsvcServer) RmibEntryCreate(context.Context, *RmibEntryCreateRequest) (*RmibEntryCreateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RmibEntryDelete(context.Context, *RmibEntryDeleteRequest) (*RmibEntryDeleteResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RmibEntrySet(context.Context, *RmibEntrySetRequest) (*RmibEntrySetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RmibEntryGet(context.Context, *RmibEntryGetRequest) (*RmibEntryGetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RmibEntryGetFirst(context.Context, *RmibEntryGetFirstRequest) (*RmibEntryGetFirstResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RmibEntryGetNext(context.Context, *RmibEntryGetNextRequest) (*RmibEntryGetNextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RmibGetTrapInfo(context.Context, *RmibGetTrapInfoRequest) (*RmibGetTrapInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RmibSetTrapInfo(context.Context, *RmibSetTrapInfoRequest) (*RmibSetTrapInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RASAdminConnectionNotification(context.Context, *RASAdminConnectionNotificationRequest) (*RASAdminConnectionNotificationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RASAdminSendUserMessage(context.Context, *RASAdminSendUserMessageRequest) (*RASAdminSendUserMessageResponse, error) {
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
func (UnimplementedDimsvcServer) RASAdminConnectionRemoveQuarantine(context.Context, *RASAdminConnectionRemoveQuarantineRequest) (*RASAdminConnectionRemoveQuarantineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) MprAdminServerSetInfo(context.Context, *MprAdminServerSetInfoRequest) (*MprAdminServerSetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) MprAdminServerGetInfoEx(context.Context, *MprAdminServerGetInfoExRequest) (*MprAdminServerGetInfoExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RASAdminConnectionEnumEx(context.Context, *RASAdminConnectionEnumExRequest) (*RASAdminConnectionEnumExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RASAdminConnectionGetInfoEx(context.Context, *RASAdminConnectionGetInfoExRequest) (*RASAdminConnectionGetInfoExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) MprAdminServerSetInfoEx(context.Context, *MprAdminServerSetInfoExRequest) (*MprAdminServerSetInfoExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDimsvcServer) RASAdminUpdateConnection(context.Context, *RASAdminUpdateConnectionRequest) (*RASAdminUpdateConnectionResponse, error) {
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
