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
		in := &OpenClusterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenCluster(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // ApiCloseCluster
		in := &CloseClusterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseCluster(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // ApiSetClusterName
		in := &SetClusterNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetClusterName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // ApiGetClusterName
		in := &GetClusterNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClusterName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // ApiGetClusterVersion
		in := &GetClusterVersionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClusterVersion(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // ApiGetQuorumResource
		in := &GetQuorumResourceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetQuorumResource(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // ApiSetQuorumResource
		in := &SetQuorumResourceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetQuorumResource(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // ApiCreateEnum
		in := &CreateEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // ApiOpenResource
		in := &OpenResourceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenResource(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // ApiCreateResource
		in := &CreateResourceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateResource(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // ApiDeleteResource
		in := &DeleteResourceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteResource(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // ApiCloseResource
		in := &CloseResourceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseResource(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // ApiGetResourceState
		in := &GetResourceStateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetResourceState(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // ApiSetResourceName
		in := &SetResourceNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetResourceName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // ApiGetResourceId
		in := &GetResourceIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetResourceID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // ApiGetResourceType
		in := &GetResourceTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetResourceType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // ApiFailResource
		in := &FailResourceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FailResource(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // ApiOnlineResource
		in := &OnlineResourceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OnlineResource(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // ApiOfflineResource
		in := &OfflineResourceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OfflineResource(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // ApiAddResourceDependency
		in := &AddResourceDependencyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddResourceDependency(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // ApiRemoveResourceDependency
		in := &RemoveResourceDependencyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveResourceDependency(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // ApiCanResourceBeDependent
		in := &CanResourceBeDependentRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CanResourceBeDependent(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // ApiCreateResEnum
		in := &CreateRestrictionEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateRestrictionEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // ApiAddResourceNode
		in := &AddResourceNodeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddResourceNode(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // ApiRemoveResourceNode
		in := &RemoveResourceNodeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveResourceNode(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // ApiChangeResourceGroup
		in := &ChangeResourceGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ChangeResourceGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // ApiCreateResourceType
		in := &CreateResourceTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateResourceType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // ApiDeleteResourceType
		in := &DeleteResourceTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteResourceType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // ApiGetRootKey
		in := &GetRootKeyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetRootKey(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // ApiCreateKey
		in := &CreateKeyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateKey(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // ApiOpenKey
		in := &OpenKeyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenKey(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // ApiEnumKey
		in := &EnumKeyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumKey(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // ApiSetValue
		in := &SetValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // ApiDeleteValue
		in := &DeleteValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // ApiQueryValue
		in := &QueryValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // ApiDeleteKey
		in := &DeleteKeyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteKey(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // ApiEnumValue
		in := &EnumValueRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumValue(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // ApiCloseKey
		in := &CloseKeyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseKey(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 38: // ApiQueryInfoKey
		in := &QueryInfoKeyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryInfoKey(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 39: // ApiSetKeySecurity
		in := &SetKeySecurityRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetKeySecurity(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 40: // ApiGetKeySecurity
		in := &GetKeySecurityRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetKeySecurity(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 41: // ApiOpenGroup
		in := &OpenGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 42: // ApiCreateGroup
		in := &CreateGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 43: // ApiDeleteGroup
		in := &DeleteGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 44: // ApiCloseGroup
		in := &CloseGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 45: // ApiGetGroupState
		in := &GetGroupStateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetGroupState(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 46: // ApiSetGroupName
		in := &SetGroupNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetGroupName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 47: // ApiGetGroupId
		in := &GetGroupIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetGroupID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 48: // ApiGetNodeId
		in := &GetNodeIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNodeID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 49: // ApiOnlineGroup
		in := &OnlineGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OnlineGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 50: // ApiOfflineGroup
		in := &OfflineGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OfflineGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 51: // ApiMoveGroup
		in := &MoveGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.MoveGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 52: // ApiMoveGroupToNode
		in := &MoveGroupToNodeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.MoveGroupToNode(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 53: // ApiCreateGroupResourceEnum
		in := &CreateGroupResourceEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateGroupResourceEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 54: // ApiSetGroupNodeList
		in := &SetGroupNodeListRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetGroupNodeList(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 55: // ApiCreateNotify
		in := &CreateNotifyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateNotify(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 56: // ApiCloseNotify
		in := &CloseNotifyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseNotify(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 57: // ApiAddNotifyCluster
		in := &AddNotifyClusterRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddNotifyCluster(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 58: // ApiAddNotifyNode
		in := &AddNotifyNodeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddNotifyNode(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 59: // ApiAddNotifyGroup
		in := &AddNotifyGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddNotifyGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 60: // ApiAddNotifyResource
		in := &AddNotifyResourceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddNotifyResource(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 61: // ApiAddNotifyKey
		in := &AddNotifyKeyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddNotifyKey(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 62: // ApiReAddNotifyNode
		in := &ReAddNotifyNodeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ReAddNotifyNode(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 63: // ApiReAddNotifyGroup
		in := &ReAddNotifyGroupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ReAddNotifyGroup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 64: // ApiReAddNotifyResource
		in := &ReAddNotifyResourceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ReAddNotifyResource(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 65: // ApiGetNotify
		in := &GetNotifyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNotify(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 66: // ApiOpenNode
		in := &OpenNodeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenNode(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 67: // ApiCloseNode
		in := &CloseNodeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseNode(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 68: // ApiGetNodeState
		in := &GetNodeStateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNodeState(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 69: // ApiPauseNode
		in := &PauseNodeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PauseNode(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 70: // ApiResumeNode
		in := &ResumeNodeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ResumeNode(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 71: // ApiEvictNode
		in := &EvictNodeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EvictNode(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 72: // ApiNodeResourceControl
		in := &NodeResourceControlRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.NodeResourceControl(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 73: // ApiResourceControl
		in := &ResourceControlRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ResourceControl(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 74: // ApiNodeResourceTypeControl
		in := &NodeResourceTypeControlRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.NodeResourceTypeControl(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 75: // ApiResourceTypeControl
		in := &ResourceTypeControlRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ResourceTypeControl(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 76: // ApiNodeGroupControl
		in := &NodeGroupControlRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.NodeGroupControl(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 77: // ApiGroupControl
		in := &GroupControlRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GroupControl(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 78: // ApiNodeNodeControl
		in := &NodeNodeControlRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.NodeNodeControl(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 79: // ApiNodeControl
		in := &NodeControlRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.NodeControl(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 80: // Opnum80NotUsedOnWire
		// Opnum80NotUsedOnWire
		return nil, nil
	case 81: // ApiOpenNetwork
		in := &OpenNetworkRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenNetwork(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 82: // ApiCloseNetwork
		in := &CloseNetworkRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseNetwork(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 83: // ApiGetNetworkState
		in := &GetNetworkStateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNetworkState(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 84: // ApiSetNetworkName
		in := &SetNetworkNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetNetworkName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 85: // ApiCreateNetworkEnum
		in := &CreateNetworkEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateNetworkEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 86: // ApiGetNetworkId
		in := &GetNetworkIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNetworkID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 87: // ApiSetNetworkPriorityOrder
		in := &SetNetworkPriorityOrderRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetNetworkPriorityOrder(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 88: // ApiNodeNetworkControl
		in := &NodeNetworkControlRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.NodeNetworkControl(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 89: // ApiNetworkControl
		in := &NetworkControlRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.NetworkControl(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 90: // ApiAddNotifyNetwork
		in := &AddNotifyNetworkRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddNotifyNetwork(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 91: // ApiReAddNotifyNetwork
		in := &ReAddNotifyNetworkRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ReAddNotifyNetwork(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 92: // ApiOpenNetInterface
		in := &OpenNetInterfaceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenNetInterface(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 93: // ApiCloseNetInterface
		in := &CloseNetInterfaceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseNetInterface(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 94: // ApiGetNetInterfaceState
		in := &GetNetInterfaceStateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNetInterfaceState(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 95: // ApiGetNetInterface
		in := &GetNetInterfaceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNetInterface(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 96: // ApiGetNetInterfaceId
		in := &GetNetInterfaceIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNetInterfaceID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 97: // ApiNodeNetInterfaceControl
		in := &NodeNetInterfaceControlRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.NodeNetInterfaceControl(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 98: // ApiNetInterfaceControl
		in := &NetInterfaceControlRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.NetInterfaceControl(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 99: // ApiAddNotifyNetInterface
		in := &AddNotifyNetInterfaceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddNotifyNetInterface(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 100: // ApiReAddNotifyNetInterface
		in := &ReAddNotifyNetInterfaceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ReAddNotifyNetInterface(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 101: // ApiCreateNodeEnum
		in := &CreateNodeEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateNodeEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 102: // ApiGetClusterVersion2
		in := &GetClusterVersion2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClusterVersion2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 103: // ApiCreateResTypeEnum
		in := &CreateRestrictionTypeEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateRestrictionTypeEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 104: // ApiBackupClusterDatabase
		in := &BackupClusterDatabaseRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.BackupClusterDatabase(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 105: // ApiNodeClusterControl
		in := &NodeClusterControlRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.NodeClusterControl(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 106: // ApiClusterControl
		in := &ClusterControlRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ClusterControl(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 107: // ApiUnblockGetNotifyCall
		in := &UnblockGetNotifyCallRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.UnblockGetNotifyCall(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 108: // ApiSetServiceAccountPassword
		in := &SetServiceAccountPasswordRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetServiceAccountPassword(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
