package clusapi2

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

// clusapi2 server interface.
type Clusapi2Server interface {

	// ApiOpenCluster operation.
	OpenCluster(context.Context, *OpenClusterRequest) (*OpenClusterResponse, error)

	// ApiCloseCluster operation.
	CloseCluster(context.Context, *CloseClusterRequest) (*CloseClusterResponse, error)

	// ApiSetClusterName operation.
	SetClusterName(context.Context, *SetClusterNameRequest) (*SetClusterNameResponse, error)

	// ApiGetClusterName operation.
	GetClusterName(context.Context, *GetClusterNameRequest) (*GetClusterNameResponse, error)

	// ApiGetClusterVersion operation.
	GetClusterVersion(context.Context, *GetClusterVersionRequest) (*GetClusterVersionResponse, error)

	// ApiGetQuorumResource operation.
	GetQuorumResource(context.Context, *GetQuorumResourceRequest) (*GetQuorumResourceResponse, error)

	// ApiSetQuorumResource operation.
	SetQuorumResource(context.Context, *SetQuorumResourceRequest) (*SetQuorumResourceResponse, error)

	// ApiCreateEnum operation.
	CreateEnum(context.Context, *CreateEnumRequest) (*CreateEnumResponse, error)

	// ApiOpenResource operation.
	OpenResource(context.Context, *OpenResourceRequest) (*OpenResourceResponse, error)

	// ApiCreateResource operation.
	CreateResource(context.Context, *CreateResourceRequest) (*CreateResourceResponse, error)

	// ApiDeleteResource operation.
	DeleteResource(context.Context, *DeleteResourceRequest) (*DeleteResourceResponse, error)

	// ApiCloseResource operation.
	CloseResource(context.Context, *CloseResourceRequest) (*CloseResourceResponse, error)

	// ApiGetResourceState operation.
	GetResourceState(context.Context, *GetResourceStateRequest) (*GetResourceStateResponse, error)

	// ApiSetResourceName operation.
	SetResourceName(context.Context, *SetResourceNameRequest) (*SetResourceNameResponse, error)

	// ApiGetResourceId operation.
	GetResourceID(context.Context, *GetResourceIDRequest) (*GetResourceIDResponse, error)

	// ApiGetResourceType operation.
	GetResourceType(context.Context, *GetResourceTypeRequest) (*GetResourceTypeResponse, error)

	// ApiFailResource operation.
	FailResource(context.Context, *FailResourceRequest) (*FailResourceResponse, error)

	// ApiOnlineResource operation.
	OnlineResource(context.Context, *OnlineResourceRequest) (*OnlineResourceResponse, error)

	// ApiOfflineResource operation.
	OfflineResource(context.Context, *OfflineResourceRequest) (*OfflineResourceResponse, error)

	// ApiAddResourceDependency operation.
	AddResourceDependency(context.Context, *AddResourceDependencyRequest) (*AddResourceDependencyResponse, error)

	// ApiRemoveResourceDependency operation.
	RemoveResourceDependency(context.Context, *RemoveResourceDependencyRequest) (*RemoveResourceDependencyResponse, error)

	// ApiCanResourceBeDependent operation.
	CanResourceBeDependent(context.Context, *CanResourceBeDependentRequest) (*CanResourceBeDependentResponse, error)

	// ApiCreateResEnum operation.
	CreateRestrictionEnum(context.Context, *CreateRestrictionEnumRequest) (*CreateRestrictionEnumResponse, error)

	// ApiAddResourceNode operation.
	AddResourceNode(context.Context, *AddResourceNodeRequest) (*AddResourceNodeResponse, error)

	// ApiRemoveResourceNode operation.
	RemoveResourceNode(context.Context, *RemoveResourceNodeRequest) (*RemoveResourceNodeResponse, error)

	// ApiChangeResourceGroup operation.
	ChangeResourceGroup(context.Context, *ChangeResourceGroupRequest) (*ChangeResourceGroupResponse, error)

	// ApiCreateResourceType operation.
	CreateResourceType(context.Context, *CreateResourceTypeRequest) (*CreateResourceTypeResponse, error)

	// ApiDeleteResourceType operation.
	DeleteResourceType(context.Context, *DeleteResourceTypeRequest) (*DeleteResourceTypeResponse, error)

	// ApiGetRootKey operation.
	GetRootKey(context.Context, *GetRootKeyRequest) (*GetRootKeyResponse, error)

	// ApiCreateKey operation.
	CreateKey(context.Context, *CreateKeyRequest) (*CreateKeyResponse, error)

	// ApiOpenKey operation.
	OpenKey(context.Context, *OpenKeyRequest) (*OpenKeyResponse, error)

	// ApiEnumKey operation.
	EnumKey(context.Context, *EnumKeyRequest) (*EnumKeyResponse, error)

	// ApiSetValue operation.
	SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error)

	// ApiDeleteValue operation.
	DeleteValue(context.Context, *DeleteValueRequest) (*DeleteValueResponse, error)

	// ApiQueryValue operation.
	QueryValue(context.Context, *QueryValueRequest) (*QueryValueResponse, error)

	// ApiDeleteKey operation.
	DeleteKey(context.Context, *DeleteKeyRequest) (*DeleteKeyResponse, error)

	// ApiEnumValue operation.
	EnumValue(context.Context, *EnumValueRequest) (*EnumValueResponse, error)

	// ApiCloseKey operation.
	CloseKey(context.Context, *CloseKeyRequest) (*CloseKeyResponse, error)

	// ApiQueryInfoKey operation.
	QueryInfoKey(context.Context, *QueryInfoKeyRequest) (*QueryInfoKeyResponse, error)

	// ApiSetKeySecurity operation.
	SetKeySecurity(context.Context, *SetKeySecurityRequest) (*SetKeySecurityResponse, error)

	// ApiGetKeySecurity operation.
	GetKeySecurity(context.Context, *GetKeySecurityRequest) (*GetKeySecurityResponse, error)

	// ApiOpenGroup operation.
	OpenGroup(context.Context, *OpenGroupRequest) (*OpenGroupResponse, error)

	// ApiCreateGroup operation.
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error)

	// ApiDeleteGroup operation.
	DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupResponse, error)

	// ApiCloseGroup operation.
	CloseGroup(context.Context, *CloseGroupRequest) (*CloseGroupResponse, error)

	// ApiGetGroupState operation.
	GetGroupState(context.Context, *GetGroupStateRequest) (*GetGroupStateResponse, error)

	// ApiSetGroupName operation.
	SetGroupName(context.Context, *SetGroupNameRequest) (*SetGroupNameResponse, error)

	// ApiGetGroupId operation.
	GetGroupID(context.Context, *GetGroupIDRequest) (*GetGroupIDResponse, error)

	// ApiGetNodeId operation.
	GetNodeID(context.Context, *GetNodeIDRequest) (*GetNodeIDResponse, error)

	// ApiOnlineGroup operation.
	OnlineGroup(context.Context, *OnlineGroupRequest) (*OnlineGroupResponse, error)

	// ApiOfflineGroup operation.
	OfflineGroup(context.Context, *OfflineGroupRequest) (*OfflineGroupResponse, error)

	// ApiMoveGroup operation.
	MoveGroup(context.Context, *MoveGroupRequest) (*MoveGroupResponse, error)

	// ApiMoveGroupToNode operation.
	MoveGroupToNode(context.Context, *MoveGroupToNodeRequest) (*MoveGroupToNodeResponse, error)

	// ApiCreateGroupResourceEnum operation.
	CreateGroupResourceEnum(context.Context, *CreateGroupResourceEnumRequest) (*CreateGroupResourceEnumResponse, error)

	// ApiSetGroupNodeList operation.
	SetGroupNodeList(context.Context, *SetGroupNodeListRequest) (*SetGroupNodeListResponse, error)

	// ApiCreateNotify operation.
	CreateNotify(context.Context, *CreateNotifyRequest) (*CreateNotifyResponse, error)

	// ApiCloseNotify operation.
	CloseNotify(context.Context, *CloseNotifyRequest) (*CloseNotifyResponse, error)

	// ApiAddNotifyCluster operation.
	AddNotifyCluster(context.Context, *AddNotifyClusterRequest) (*AddNotifyClusterResponse, error)

	// ApiAddNotifyNode operation.
	AddNotifyNode(context.Context, *AddNotifyNodeRequest) (*AddNotifyNodeResponse, error)

	// ApiAddNotifyGroup operation.
	AddNotifyGroup(context.Context, *AddNotifyGroupRequest) (*AddNotifyGroupResponse, error)

	// ApiAddNotifyResource operation.
	AddNotifyResource(context.Context, *AddNotifyResourceRequest) (*AddNotifyResourceResponse, error)

	// ApiAddNotifyKey operation.
	AddNotifyKey(context.Context, *AddNotifyKeyRequest) (*AddNotifyKeyResponse, error)

	// ApiReAddNotifyNode operation.
	ReAddNotifyNode(context.Context, *ReAddNotifyNodeRequest) (*ReAddNotifyNodeResponse, error)

	// ApiReAddNotifyGroup operation.
	ReAddNotifyGroup(context.Context, *ReAddNotifyGroupRequest) (*ReAddNotifyGroupResponse, error)

	// ApiReAddNotifyResource operation.
	ReAddNotifyResource(context.Context, *ReAddNotifyResourceRequest) (*ReAddNotifyResourceResponse, error)

	// ApiGetNotify operation.
	GetNotify(context.Context, *GetNotifyRequest) (*GetNotifyResponse, error)

	// ApiOpenNode operation.
	OpenNode(context.Context, *OpenNodeRequest) (*OpenNodeResponse, error)

	// ApiCloseNode operation.
	CloseNode(context.Context, *CloseNodeRequest) (*CloseNodeResponse, error)

	// ApiGetNodeState operation.
	GetNodeState(context.Context, *GetNodeStateRequest) (*GetNodeStateResponse, error)

	// ApiPauseNode operation.
	PauseNode(context.Context, *PauseNodeRequest) (*PauseNodeResponse, error)

	// ApiResumeNode operation.
	ResumeNode(context.Context, *ResumeNodeRequest) (*ResumeNodeResponse, error)

	// ApiEvictNode operation.
	EvictNode(context.Context, *EvictNodeRequest) (*EvictNodeResponse, error)

	// ApiNodeResourceControl operation.
	NodeResourceControl(context.Context, *NodeResourceControlRequest) (*NodeResourceControlResponse, error)

	// ApiResourceControl operation.
	ResourceControl(context.Context, *ResourceControlRequest) (*ResourceControlResponse, error)

	// ApiNodeResourceTypeControl operation.
	NodeResourceTypeControl(context.Context, *NodeResourceTypeControlRequest) (*NodeResourceTypeControlResponse, error)

	// ApiResourceTypeControl operation.
	ResourceTypeControl(context.Context, *ResourceTypeControlRequest) (*ResourceTypeControlResponse, error)

	// ApiNodeGroupControl operation.
	NodeGroupControl(context.Context, *NodeGroupControlRequest) (*NodeGroupControlResponse, error)

	// ApiGroupControl operation.
	GroupControl(context.Context, *GroupControlRequest) (*GroupControlResponse, error)

	// ApiNodeNodeControl operation.
	NodeNodeControl(context.Context, *NodeNodeControlRequest) (*NodeNodeControlResponse, error)

	// ApiNodeControl operation.
	NodeControl(context.Context, *NodeControlRequest) (*NodeControlResponse, error)

	// Opnum80NotUsedOnWire operation.
	// Opnum80NotUsedOnWire

	// ApiOpenNetwork operation.
	OpenNetwork(context.Context, *OpenNetworkRequest) (*OpenNetworkResponse, error)

	// ApiCloseNetwork operation.
	CloseNetwork(context.Context, *CloseNetworkRequest) (*CloseNetworkResponse, error)

	// ApiGetNetworkState operation.
	GetNetworkState(context.Context, *GetNetworkStateRequest) (*GetNetworkStateResponse, error)

	// ApiSetNetworkName operation.
	SetNetworkName(context.Context, *SetNetworkNameRequest) (*SetNetworkNameResponse, error)

	// ApiCreateNetworkEnum operation.
	CreateNetworkEnum(context.Context, *CreateNetworkEnumRequest) (*CreateNetworkEnumResponse, error)

	// ApiGetNetworkId operation.
	GetNetworkID(context.Context, *GetNetworkIDRequest) (*GetNetworkIDResponse, error)

	// ApiSetNetworkPriorityOrder operation.
	SetNetworkPriorityOrder(context.Context, *SetNetworkPriorityOrderRequest) (*SetNetworkPriorityOrderResponse, error)

	// ApiNodeNetworkControl operation.
	NodeNetworkControl(context.Context, *NodeNetworkControlRequest) (*NodeNetworkControlResponse, error)

	// ApiNetworkControl operation.
	NetworkControl(context.Context, *NetworkControlRequest) (*NetworkControlResponse, error)

	// ApiAddNotifyNetwork operation.
	AddNotifyNetwork(context.Context, *AddNotifyNetworkRequest) (*AddNotifyNetworkResponse, error)

	// ApiReAddNotifyNetwork operation.
	ReAddNotifyNetwork(context.Context, *ReAddNotifyNetworkRequest) (*ReAddNotifyNetworkResponse, error)

	// ApiOpenNetInterface operation.
	OpenNetInterface(context.Context, *OpenNetInterfaceRequest) (*OpenNetInterfaceResponse, error)

	// ApiCloseNetInterface operation.
	CloseNetInterface(context.Context, *CloseNetInterfaceRequest) (*CloseNetInterfaceResponse, error)

	// ApiGetNetInterfaceState operation.
	GetNetInterfaceState(context.Context, *GetNetInterfaceStateRequest) (*GetNetInterfaceStateResponse, error)

	// ApiGetNetInterface operation.
	GetNetInterface(context.Context, *GetNetInterfaceRequest) (*GetNetInterfaceResponse, error)

	// ApiGetNetInterfaceId operation.
	GetNetInterfaceID(context.Context, *GetNetInterfaceIDRequest) (*GetNetInterfaceIDResponse, error)

	// ApiNodeNetInterfaceControl operation.
	NodeNetInterfaceControl(context.Context, *NodeNetInterfaceControlRequest) (*NodeNetInterfaceControlResponse, error)

	// ApiNetInterfaceControl operation.
	NetInterfaceControl(context.Context, *NetInterfaceControlRequest) (*NetInterfaceControlResponse, error)

	// ApiAddNotifyNetInterface operation.
	AddNotifyNetInterface(context.Context, *AddNotifyNetInterfaceRequest) (*AddNotifyNetInterfaceResponse, error)

	// ApiReAddNotifyNetInterface operation.
	ReAddNotifyNetInterface(context.Context, *ReAddNotifyNetInterfaceRequest) (*ReAddNotifyNetInterfaceResponse, error)

	// ApiCreateNodeEnum operation.
	CreateNodeEnum(context.Context, *CreateNodeEnumRequest) (*CreateNodeEnumResponse, error)

	// ApiGetClusterVersion2 operation.
	GetClusterVersion2(context.Context, *GetClusterVersion2Request) (*GetClusterVersion2Response, error)

	// ApiCreateResTypeEnum operation.
	CreateRestrictionTypeEnum(context.Context, *CreateRestrictionTypeEnumRequest) (*CreateRestrictionTypeEnumResponse, error)

	// ApiBackupClusterDatabase operation.
	BackupClusterDatabase(context.Context, *BackupClusterDatabaseRequest) (*BackupClusterDatabaseResponse, error)

	// ApiNodeClusterControl operation.
	NodeClusterControl(context.Context, *NodeClusterControlRequest) (*NodeClusterControlResponse, error)

	// ApiClusterControl operation.
	ClusterControl(context.Context, *ClusterControlRequest) (*ClusterControlResponse, error)

	// ApiUnblockGetNotifyCall operation.
	UnblockGetNotifyCall(context.Context, *UnblockGetNotifyCallRequest) (*UnblockGetNotifyCallResponse, error)

	// ApiSetServiceAccountPassword operation.
	SetServiceAccountPassword(context.Context, *SetServiceAccountPasswordRequest) (*SetServiceAccountPasswordResponse, error)
}

func RegisterClusapi2Server(conn dcerpc.Conn, o Clusapi2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewClusapi2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Clusapi2SyntaxV2_0))...)
}

func NewClusapi2ServerHandle(o Clusapi2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Clusapi2ServerHandle(ctx, o, opNum, r)
	}
}

func Clusapi2ServerHandle(ctx context.Context, o Clusapi2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // ApiOpenCluster
		op := &xxx_OpenClusterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenClusterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenCluster(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // ApiCloseCluster
		op := &xxx_CloseClusterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseClusterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseCluster(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // ApiSetClusterName
		op := &xxx_SetClusterNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetClusterNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetClusterName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // ApiGetClusterName
		op := &xxx_GetClusterNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClusterNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClusterName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // ApiGetClusterVersion
		op := &xxx_GetClusterVersionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClusterVersionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClusterVersion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // ApiGetQuorumResource
		op := &xxx_GetQuorumResourceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetQuorumResourceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetQuorumResource(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // ApiSetQuorumResource
		op := &xxx_SetQuorumResourceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetQuorumResourceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetQuorumResource(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // ApiCreateEnum
		op := &xxx_CreateEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // ApiOpenResource
		op := &xxx_OpenResourceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenResourceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenResource(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ApiCreateResource
		op := &xxx_CreateResourceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateResourceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateResource(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // ApiDeleteResource
		op := &xxx_DeleteResourceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteResourceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteResource(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // ApiCloseResource
		op := &xxx_CloseResourceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseResourceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseResource(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // ApiGetResourceState
		op := &xxx_GetResourceStateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetResourceStateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetResourceState(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // ApiSetResourceName
		op := &xxx_SetResourceNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetResourceNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetResourceName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // ApiGetResourceId
		op := &xxx_GetResourceIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetResourceIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetResourceID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // ApiGetResourceType
		op := &xxx_GetResourceTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetResourceTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetResourceType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // ApiFailResource
		op := &xxx_FailResourceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FailResourceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FailResource(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // ApiOnlineResource
		op := &xxx_OnlineResourceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OnlineResourceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OnlineResource(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // ApiOfflineResource
		op := &xxx_OfflineResourceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OfflineResourceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OfflineResource(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // ApiAddResourceDependency
		op := &xxx_AddResourceDependencyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddResourceDependencyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddResourceDependency(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // ApiRemoveResourceDependency
		op := &xxx_RemoveResourceDependencyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveResourceDependencyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveResourceDependency(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // ApiCanResourceBeDependent
		op := &xxx_CanResourceBeDependentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CanResourceBeDependentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CanResourceBeDependent(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // ApiCreateResEnum
		op := &xxx_CreateRestrictionEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateRestrictionEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateRestrictionEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // ApiAddResourceNode
		op := &xxx_AddResourceNodeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddResourceNodeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddResourceNode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // ApiRemoveResourceNode
		op := &xxx_RemoveResourceNodeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveResourceNodeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveResourceNode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // ApiChangeResourceGroup
		op := &xxx_ChangeResourceGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ChangeResourceGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ChangeResourceGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // ApiCreateResourceType
		op := &xxx_CreateResourceTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateResourceTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateResourceType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // ApiDeleteResourceType
		op := &xxx_DeleteResourceTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteResourceTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteResourceType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // ApiGetRootKey
		op := &xxx_GetRootKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRootKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRootKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // ApiCreateKey
		op := &xxx_CreateKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // ApiOpenKey
		op := &xxx_OpenKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // ApiEnumKey
		op := &xxx_EnumKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // ApiSetValue
		op := &xxx_SetValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // ApiDeleteValue
		op := &xxx_DeleteValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // ApiQueryValue
		op := &xxx_QueryValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // ApiDeleteKey
		op := &xxx_DeleteKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // ApiEnumValue
		op := &xxx_EnumValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // ApiCloseKey
		op := &xxx_CloseKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // ApiQueryInfoKey
		op := &xxx_QueryInfoKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryInfoKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryInfoKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // ApiSetKeySecurity
		op := &xxx_SetKeySecurityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetKeySecurityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetKeySecurity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // ApiGetKeySecurity
		op := &xxx_GetKeySecurityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetKeySecurityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetKeySecurity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // ApiOpenGroup
		op := &xxx_OpenGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 42: // ApiCreateGroup
		op := &xxx_CreateGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 43: // ApiDeleteGroup
		op := &xxx_DeleteGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // ApiCloseGroup
		op := &xxx_CloseGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // ApiGetGroupState
		op := &xxx_GetGroupStateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetGroupStateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetGroupState(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // ApiSetGroupName
		op := &xxx_SetGroupNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetGroupNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetGroupName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 47: // ApiGetGroupId
		op := &xxx_GetGroupIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetGroupIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetGroupID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 48: // ApiGetNodeId
		op := &xxx_GetNodeIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNodeIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNodeID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 49: // ApiOnlineGroup
		op := &xxx_OnlineGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OnlineGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OnlineGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 50: // ApiOfflineGroup
		op := &xxx_OfflineGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OfflineGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OfflineGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 51: // ApiMoveGroup
		op := &xxx_MoveGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MoveGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MoveGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 52: // ApiMoveGroupToNode
		op := &xxx_MoveGroupToNodeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MoveGroupToNodeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MoveGroupToNode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 53: // ApiCreateGroupResourceEnum
		op := &xxx_CreateGroupResourceEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateGroupResourceEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateGroupResourceEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 54: // ApiSetGroupNodeList
		op := &xxx_SetGroupNodeListOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetGroupNodeListRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetGroupNodeList(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 55: // ApiCreateNotify
		op := &xxx_CreateNotifyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateNotifyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateNotify(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 56: // ApiCloseNotify
		op := &xxx_CloseNotifyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseNotifyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseNotify(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 57: // ApiAddNotifyCluster
		op := &xxx_AddNotifyClusterOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddNotifyClusterRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddNotifyCluster(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 58: // ApiAddNotifyNode
		op := &xxx_AddNotifyNodeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddNotifyNodeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddNotifyNode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 59: // ApiAddNotifyGroup
		op := &xxx_AddNotifyGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddNotifyGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddNotifyGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 60: // ApiAddNotifyResource
		op := &xxx_AddNotifyResourceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddNotifyResourceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddNotifyResource(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 61: // ApiAddNotifyKey
		op := &xxx_AddNotifyKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddNotifyKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddNotifyKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 62: // ApiReAddNotifyNode
		op := &xxx_ReAddNotifyNodeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReAddNotifyNodeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReAddNotifyNode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 63: // ApiReAddNotifyGroup
		op := &xxx_ReAddNotifyGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReAddNotifyGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReAddNotifyGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 64: // ApiReAddNotifyResource
		op := &xxx_ReAddNotifyResourceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReAddNotifyResourceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReAddNotifyResource(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 65: // ApiGetNotify
		op := &xxx_GetNotifyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNotifyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNotify(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 66: // ApiOpenNode
		op := &xxx_OpenNodeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenNodeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenNode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 67: // ApiCloseNode
		op := &xxx_CloseNodeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseNodeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseNode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 68: // ApiGetNodeState
		op := &xxx_GetNodeStateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNodeStateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNodeState(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 69: // ApiPauseNode
		op := &xxx_PauseNodeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PauseNodeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PauseNode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 70: // ApiResumeNode
		op := &xxx_ResumeNodeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResumeNodeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ResumeNode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 71: // ApiEvictNode
		op := &xxx_EvictNodeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EvictNodeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EvictNode(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 72: // ApiNodeResourceControl
		op := &xxx_NodeResourceControlOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NodeResourceControlRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NodeResourceControl(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 73: // ApiResourceControl
		op := &xxx_ResourceControlOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResourceControlRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ResourceControl(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 74: // ApiNodeResourceTypeControl
		op := &xxx_NodeResourceTypeControlOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NodeResourceTypeControlRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NodeResourceTypeControl(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 75: // ApiResourceTypeControl
		op := &xxx_ResourceTypeControlOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResourceTypeControlRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ResourceTypeControl(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 76: // ApiNodeGroupControl
		op := &xxx_NodeGroupControlOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NodeGroupControlRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NodeGroupControl(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 77: // ApiGroupControl
		op := &xxx_GroupControlOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GroupControlRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GroupControl(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 78: // ApiNodeNodeControl
		op := &xxx_NodeNodeControlOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NodeNodeControlRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NodeNodeControl(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 79: // ApiNodeControl
		op := &xxx_NodeControlOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NodeControlRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NodeControl(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 80: // Opnum80NotUsedOnWire
		// Opnum80NotUsedOnWire
		return nil, nil
	case 81: // ApiOpenNetwork
		op := &xxx_OpenNetworkOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenNetworkRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenNetwork(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 82: // ApiCloseNetwork
		op := &xxx_CloseNetworkOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseNetworkRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseNetwork(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 83: // ApiGetNetworkState
		op := &xxx_GetNetworkStateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNetworkStateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNetworkState(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 84: // ApiSetNetworkName
		op := &xxx_SetNetworkNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNetworkNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetNetworkName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 85: // ApiCreateNetworkEnum
		op := &xxx_CreateNetworkEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateNetworkEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateNetworkEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 86: // ApiGetNetworkId
		op := &xxx_GetNetworkIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNetworkIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNetworkID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 87: // ApiSetNetworkPriorityOrder
		op := &xxx_SetNetworkPriorityOrderOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetNetworkPriorityOrderRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetNetworkPriorityOrder(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 88: // ApiNodeNetworkControl
		op := &xxx_NodeNetworkControlOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NodeNetworkControlRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NodeNetworkControl(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 89: // ApiNetworkControl
		op := &xxx_NetworkControlOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NetworkControlRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NetworkControl(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 90: // ApiAddNotifyNetwork
		op := &xxx_AddNotifyNetworkOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddNotifyNetworkRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddNotifyNetwork(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 91: // ApiReAddNotifyNetwork
		op := &xxx_ReAddNotifyNetworkOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReAddNotifyNetworkRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReAddNotifyNetwork(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 92: // ApiOpenNetInterface
		op := &xxx_OpenNetInterfaceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenNetInterfaceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenNetInterface(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 93: // ApiCloseNetInterface
		op := &xxx_CloseNetInterfaceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseNetInterfaceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseNetInterface(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 94: // ApiGetNetInterfaceState
		op := &xxx_GetNetInterfaceStateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNetInterfaceStateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNetInterfaceState(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 95: // ApiGetNetInterface
		op := &xxx_GetNetInterfaceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNetInterfaceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNetInterface(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 96: // ApiGetNetInterfaceId
		op := &xxx_GetNetInterfaceIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNetInterfaceIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNetInterfaceID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 97: // ApiNodeNetInterfaceControl
		op := &xxx_NodeNetInterfaceControlOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NodeNetInterfaceControlRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NodeNetInterfaceControl(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 98: // ApiNetInterfaceControl
		op := &xxx_NetInterfaceControlOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NetInterfaceControlRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NetInterfaceControl(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 99: // ApiAddNotifyNetInterface
		op := &xxx_AddNotifyNetInterfaceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddNotifyNetInterfaceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddNotifyNetInterface(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 100: // ApiReAddNotifyNetInterface
		op := &xxx_ReAddNotifyNetInterfaceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReAddNotifyNetInterfaceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReAddNotifyNetInterface(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 101: // ApiCreateNodeEnum
		op := &xxx_CreateNodeEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateNodeEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateNodeEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 102: // ApiGetClusterVersion2
		op := &xxx_GetClusterVersion2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClusterVersion2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClusterVersion2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 103: // ApiCreateResTypeEnum
		op := &xxx_CreateRestrictionTypeEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateRestrictionTypeEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateRestrictionTypeEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 104: // ApiBackupClusterDatabase
		op := &xxx_BackupClusterDatabaseOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BackupClusterDatabaseRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BackupClusterDatabase(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 105: // ApiNodeClusterControl
		op := &xxx_NodeClusterControlOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NodeClusterControlRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NodeClusterControl(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 106: // ApiClusterControl
		op := &xxx_ClusterControlOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ClusterControlRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ClusterControl(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 107: // ApiUnblockGetNotifyCall
		op := &xxx_UnblockGetNotifyCallOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UnblockGetNotifyCallRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UnblockGetNotifyCall(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 108: // ApiSetServiceAccountPassword
		op := &xxx_SetServiceAccountPasswordOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetServiceAccountPasswordRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetServiceAccountPassword(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented clusapi2
type UnimplementedClusapi2Server struct {
}

func (UnimplementedClusapi2Server) OpenCluster(context.Context, *OpenClusterRequest) (*OpenClusterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) CloseCluster(context.Context, *CloseClusterRequest) (*CloseClusterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) SetClusterName(context.Context, *SetClusterNameRequest) (*SetClusterNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) GetClusterName(context.Context, *GetClusterNameRequest) (*GetClusterNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) GetClusterVersion(context.Context, *GetClusterVersionRequest) (*GetClusterVersionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) GetQuorumResource(context.Context, *GetQuorumResourceRequest) (*GetQuorumResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) SetQuorumResource(context.Context, *SetQuorumResourceRequest) (*SetQuorumResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) CreateEnum(context.Context, *CreateEnumRequest) (*CreateEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) OpenResource(context.Context, *OpenResourceRequest) (*OpenResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) CreateResource(context.Context, *CreateResourceRequest) (*CreateResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) DeleteResource(context.Context, *DeleteResourceRequest) (*DeleteResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) CloseResource(context.Context, *CloseResourceRequest) (*CloseResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) GetResourceState(context.Context, *GetResourceStateRequest) (*GetResourceStateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) SetResourceName(context.Context, *SetResourceNameRequest) (*SetResourceNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) GetResourceID(context.Context, *GetResourceIDRequest) (*GetResourceIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) GetResourceType(context.Context, *GetResourceTypeRequest) (*GetResourceTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) FailResource(context.Context, *FailResourceRequest) (*FailResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) OnlineResource(context.Context, *OnlineResourceRequest) (*OnlineResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) OfflineResource(context.Context, *OfflineResourceRequest) (*OfflineResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) AddResourceDependency(context.Context, *AddResourceDependencyRequest) (*AddResourceDependencyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) RemoveResourceDependency(context.Context, *RemoveResourceDependencyRequest) (*RemoveResourceDependencyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) CanResourceBeDependent(context.Context, *CanResourceBeDependentRequest) (*CanResourceBeDependentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) CreateRestrictionEnum(context.Context, *CreateRestrictionEnumRequest) (*CreateRestrictionEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) AddResourceNode(context.Context, *AddResourceNodeRequest) (*AddResourceNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) RemoveResourceNode(context.Context, *RemoveResourceNodeRequest) (*RemoveResourceNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) ChangeResourceGroup(context.Context, *ChangeResourceGroupRequest) (*ChangeResourceGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) CreateResourceType(context.Context, *CreateResourceTypeRequest) (*CreateResourceTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) DeleteResourceType(context.Context, *DeleteResourceTypeRequest) (*DeleteResourceTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) GetRootKey(context.Context, *GetRootKeyRequest) (*GetRootKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) CreateKey(context.Context, *CreateKeyRequest) (*CreateKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) OpenKey(context.Context, *OpenKeyRequest) (*OpenKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) EnumKey(context.Context, *EnumKeyRequest) (*EnumKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) DeleteValue(context.Context, *DeleteValueRequest) (*DeleteValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) QueryValue(context.Context, *QueryValueRequest) (*QueryValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) DeleteKey(context.Context, *DeleteKeyRequest) (*DeleteKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) EnumValue(context.Context, *EnumValueRequest) (*EnumValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) CloseKey(context.Context, *CloseKeyRequest) (*CloseKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) QueryInfoKey(context.Context, *QueryInfoKeyRequest) (*QueryInfoKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) SetKeySecurity(context.Context, *SetKeySecurityRequest) (*SetKeySecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) GetKeySecurity(context.Context, *GetKeySecurityRequest) (*GetKeySecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) OpenGroup(context.Context, *OpenGroupRequest) (*OpenGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) CloseGroup(context.Context, *CloseGroupRequest) (*CloseGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) GetGroupState(context.Context, *GetGroupStateRequest) (*GetGroupStateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) SetGroupName(context.Context, *SetGroupNameRequest) (*SetGroupNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) GetGroupID(context.Context, *GetGroupIDRequest) (*GetGroupIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) GetNodeID(context.Context, *GetNodeIDRequest) (*GetNodeIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) OnlineGroup(context.Context, *OnlineGroupRequest) (*OnlineGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) OfflineGroup(context.Context, *OfflineGroupRequest) (*OfflineGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) MoveGroup(context.Context, *MoveGroupRequest) (*MoveGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) MoveGroupToNode(context.Context, *MoveGroupToNodeRequest) (*MoveGroupToNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) CreateGroupResourceEnum(context.Context, *CreateGroupResourceEnumRequest) (*CreateGroupResourceEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) SetGroupNodeList(context.Context, *SetGroupNodeListRequest) (*SetGroupNodeListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) CreateNotify(context.Context, *CreateNotifyRequest) (*CreateNotifyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) CloseNotify(context.Context, *CloseNotifyRequest) (*CloseNotifyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) AddNotifyCluster(context.Context, *AddNotifyClusterRequest) (*AddNotifyClusterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) AddNotifyNode(context.Context, *AddNotifyNodeRequest) (*AddNotifyNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) AddNotifyGroup(context.Context, *AddNotifyGroupRequest) (*AddNotifyGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) AddNotifyResource(context.Context, *AddNotifyResourceRequest) (*AddNotifyResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) AddNotifyKey(context.Context, *AddNotifyKeyRequest) (*AddNotifyKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) ReAddNotifyNode(context.Context, *ReAddNotifyNodeRequest) (*ReAddNotifyNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) ReAddNotifyGroup(context.Context, *ReAddNotifyGroupRequest) (*ReAddNotifyGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) ReAddNotifyResource(context.Context, *ReAddNotifyResourceRequest) (*ReAddNotifyResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) GetNotify(context.Context, *GetNotifyRequest) (*GetNotifyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) OpenNode(context.Context, *OpenNodeRequest) (*OpenNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) CloseNode(context.Context, *CloseNodeRequest) (*CloseNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) GetNodeState(context.Context, *GetNodeStateRequest) (*GetNodeStateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) PauseNode(context.Context, *PauseNodeRequest) (*PauseNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) ResumeNode(context.Context, *ResumeNodeRequest) (*ResumeNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) EvictNode(context.Context, *EvictNodeRequest) (*EvictNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) NodeResourceControl(context.Context, *NodeResourceControlRequest) (*NodeResourceControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) ResourceControl(context.Context, *ResourceControlRequest) (*ResourceControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) NodeResourceTypeControl(context.Context, *NodeResourceTypeControlRequest) (*NodeResourceTypeControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) ResourceTypeControl(context.Context, *ResourceTypeControlRequest) (*ResourceTypeControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) NodeGroupControl(context.Context, *NodeGroupControlRequest) (*NodeGroupControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) GroupControl(context.Context, *GroupControlRequest) (*GroupControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) NodeNodeControl(context.Context, *NodeNodeControlRequest) (*NodeNodeControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) NodeControl(context.Context, *NodeControlRequest) (*NodeControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) OpenNetwork(context.Context, *OpenNetworkRequest) (*OpenNetworkResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) CloseNetwork(context.Context, *CloseNetworkRequest) (*CloseNetworkResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) GetNetworkState(context.Context, *GetNetworkStateRequest) (*GetNetworkStateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) SetNetworkName(context.Context, *SetNetworkNameRequest) (*SetNetworkNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) CreateNetworkEnum(context.Context, *CreateNetworkEnumRequest) (*CreateNetworkEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) GetNetworkID(context.Context, *GetNetworkIDRequest) (*GetNetworkIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) SetNetworkPriorityOrder(context.Context, *SetNetworkPriorityOrderRequest) (*SetNetworkPriorityOrderResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) NodeNetworkControl(context.Context, *NodeNetworkControlRequest) (*NodeNetworkControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) NetworkControl(context.Context, *NetworkControlRequest) (*NetworkControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) AddNotifyNetwork(context.Context, *AddNotifyNetworkRequest) (*AddNotifyNetworkResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) ReAddNotifyNetwork(context.Context, *ReAddNotifyNetworkRequest) (*ReAddNotifyNetworkResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) OpenNetInterface(context.Context, *OpenNetInterfaceRequest) (*OpenNetInterfaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) CloseNetInterface(context.Context, *CloseNetInterfaceRequest) (*CloseNetInterfaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) GetNetInterfaceState(context.Context, *GetNetInterfaceStateRequest) (*GetNetInterfaceStateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) GetNetInterface(context.Context, *GetNetInterfaceRequest) (*GetNetInterfaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) GetNetInterfaceID(context.Context, *GetNetInterfaceIDRequest) (*GetNetInterfaceIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) NodeNetInterfaceControl(context.Context, *NodeNetInterfaceControlRequest) (*NodeNetInterfaceControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) NetInterfaceControl(context.Context, *NetInterfaceControlRequest) (*NetInterfaceControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) AddNotifyNetInterface(context.Context, *AddNotifyNetInterfaceRequest) (*AddNotifyNetInterfaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) ReAddNotifyNetInterface(context.Context, *ReAddNotifyNetInterfaceRequest) (*ReAddNotifyNetInterfaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) CreateNodeEnum(context.Context, *CreateNodeEnumRequest) (*CreateNodeEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) GetClusterVersion2(context.Context, *GetClusterVersion2Request) (*GetClusterVersion2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) CreateRestrictionTypeEnum(context.Context, *CreateRestrictionTypeEnumRequest) (*CreateRestrictionTypeEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) BackupClusterDatabase(context.Context, *BackupClusterDatabaseRequest) (*BackupClusterDatabaseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) NodeClusterControl(context.Context, *NodeClusterControlRequest) (*NodeClusterControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) ClusterControl(context.Context, *ClusterControlRequest) (*ClusterControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) UnblockGetNotifyCall(context.Context, *UnblockGetNotifyCallRequest) (*UnblockGetNotifyCallResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi2Server) SetServiceAccountPassword(context.Context, *SetServiceAccountPasswordRequest) (*SetServiceAccountPasswordResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Clusapi2Server = (*UnimplementedClusapi2Server)(nil)
