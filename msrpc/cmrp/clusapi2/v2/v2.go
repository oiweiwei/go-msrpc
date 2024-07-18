package clusapi2

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcetypes "github.com/oiweiwei/go-msrpc/msrpc/dcetypes"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
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
	_ = dcetypes.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "cmrp"
)

var (
	// Syntax UUID
	Clusapi2SyntaxUUID = &uuid.UUID{TimeLow: 0xb97db8b2, TimeMid: 0x4c63, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0xbf, ClockSeqLow: 0xf6, Node: [6]uint8{0x8, 0x0, 0x2b, 0xe2, 0x3f, 0x2f}}
	// Syntax ID
	Clusapi2SyntaxV2_0 = &dcerpc.SyntaxID{IfUUID: Clusapi2SyntaxUUID, IfVersionMajor: 2, IfVersionMinor: 0}
)

// clusapi2 interface.
type Clusapi2Client interface {

	// ApiOpenCluster operation.
	OpenCluster(context.Context, *OpenClusterRequest, ...dcerpc.CallOption) (*OpenClusterResponse, error)

	// ApiCloseCluster operation.
	CloseCluster(context.Context, *CloseClusterRequest, ...dcerpc.CallOption) (*CloseClusterResponse, error)

	// ApiSetClusterName operation.
	SetClusterName(context.Context, *SetClusterNameRequest, ...dcerpc.CallOption) (*SetClusterNameResponse, error)

	// ApiGetClusterName operation.
	GetClusterName(context.Context, *GetClusterNameRequest, ...dcerpc.CallOption) (*GetClusterNameResponse, error)

	// ApiGetClusterVersion operation.
	GetClusterVersion(context.Context, *GetClusterVersionRequest, ...dcerpc.CallOption) (*GetClusterVersionResponse, error)

	// ApiGetQuorumResource operation.
	GetQuorumResource(context.Context, *GetQuorumResourceRequest, ...dcerpc.CallOption) (*GetQuorumResourceResponse, error)

	// ApiSetQuorumResource operation.
	SetQuorumResource(context.Context, *SetQuorumResourceRequest, ...dcerpc.CallOption) (*SetQuorumResourceResponse, error)

	// ApiCreateEnum operation.
	CreateEnum(context.Context, *CreateEnumRequest, ...dcerpc.CallOption) (*CreateEnumResponse, error)

	// ApiOpenResource operation.
	OpenResource(context.Context, *OpenResourceRequest, ...dcerpc.CallOption) (*OpenResourceResponse, error)

	// ApiCreateResource operation.
	CreateResource(context.Context, *CreateResourceRequest, ...dcerpc.CallOption) (*CreateResourceResponse, error)

	// ApiDeleteResource operation.
	DeleteResource(context.Context, *DeleteResourceRequest, ...dcerpc.CallOption) (*DeleteResourceResponse, error)

	// ApiCloseResource operation.
	CloseResource(context.Context, *CloseResourceRequest, ...dcerpc.CallOption) (*CloseResourceResponse, error)

	// ApiGetResourceState operation.
	GetResourceState(context.Context, *GetResourceStateRequest, ...dcerpc.CallOption) (*GetResourceStateResponse, error)

	// ApiSetResourceName operation.
	SetResourceName(context.Context, *SetResourceNameRequest, ...dcerpc.CallOption) (*SetResourceNameResponse, error)

	// ApiGetResourceId operation.
	GetResourceID(context.Context, *GetResourceIDRequest, ...dcerpc.CallOption) (*GetResourceIDResponse, error)

	// ApiGetResourceType operation.
	GetResourceType(context.Context, *GetResourceTypeRequest, ...dcerpc.CallOption) (*GetResourceTypeResponse, error)

	// ApiFailResource operation.
	FailResource(context.Context, *FailResourceRequest, ...dcerpc.CallOption) (*FailResourceResponse, error)

	// ApiOnlineResource operation.
	OnlineResource(context.Context, *OnlineResourceRequest, ...dcerpc.CallOption) (*OnlineResourceResponse, error)

	// ApiOfflineResource operation.
	OfflineResource(context.Context, *OfflineResourceRequest, ...dcerpc.CallOption) (*OfflineResourceResponse, error)

	// ApiAddResourceDependency operation.
	AddResourceDependency(context.Context, *AddResourceDependencyRequest, ...dcerpc.CallOption) (*AddResourceDependencyResponse, error)

	// ApiRemoveResourceDependency operation.
	RemoveResourceDependency(context.Context, *RemoveResourceDependencyRequest, ...dcerpc.CallOption) (*RemoveResourceDependencyResponse, error)

	// ApiCanResourceBeDependent operation.
	CanResourceBeDependent(context.Context, *CanResourceBeDependentRequest, ...dcerpc.CallOption) (*CanResourceBeDependentResponse, error)

	// ApiCreateResEnum operation.
	CreateRestrictionEnum(context.Context, *CreateRestrictionEnumRequest, ...dcerpc.CallOption) (*CreateRestrictionEnumResponse, error)

	// ApiAddResourceNode operation.
	AddResourceNode(context.Context, *AddResourceNodeRequest, ...dcerpc.CallOption) (*AddResourceNodeResponse, error)

	// ApiRemoveResourceNode operation.
	RemoveResourceNode(context.Context, *RemoveResourceNodeRequest, ...dcerpc.CallOption) (*RemoveResourceNodeResponse, error)

	// ApiChangeResourceGroup operation.
	ChangeResourceGroup(context.Context, *ChangeResourceGroupRequest, ...dcerpc.CallOption) (*ChangeResourceGroupResponse, error)

	// ApiCreateResourceType operation.
	CreateResourceType(context.Context, *CreateResourceTypeRequest, ...dcerpc.CallOption) (*CreateResourceTypeResponse, error)

	// ApiDeleteResourceType operation.
	DeleteResourceType(context.Context, *DeleteResourceTypeRequest, ...dcerpc.CallOption) (*DeleteResourceTypeResponse, error)

	// ApiGetRootKey operation.
	GetRootKey(context.Context, *GetRootKeyRequest, ...dcerpc.CallOption) (*GetRootKeyResponse, error)

	// ApiCreateKey operation.
	CreateKey(context.Context, *CreateKeyRequest, ...dcerpc.CallOption) (*CreateKeyResponse, error)

	// ApiOpenKey operation.
	OpenKey(context.Context, *OpenKeyRequest, ...dcerpc.CallOption) (*OpenKeyResponse, error)

	// ApiEnumKey operation.
	EnumKey(context.Context, *EnumKeyRequest, ...dcerpc.CallOption) (*EnumKeyResponse, error)

	// ApiSetValue operation.
	SetValue(context.Context, *SetValueRequest, ...dcerpc.CallOption) (*SetValueResponse, error)

	// ApiDeleteValue operation.
	DeleteValue(context.Context, *DeleteValueRequest, ...dcerpc.CallOption) (*DeleteValueResponse, error)

	// ApiQueryValue operation.
	QueryValue(context.Context, *QueryValueRequest, ...dcerpc.CallOption) (*QueryValueResponse, error)

	// ApiDeleteKey operation.
	DeleteKey(context.Context, *DeleteKeyRequest, ...dcerpc.CallOption) (*DeleteKeyResponse, error)

	// ApiEnumValue operation.
	EnumValue(context.Context, *EnumValueRequest, ...dcerpc.CallOption) (*EnumValueResponse, error)

	// ApiCloseKey operation.
	CloseKey(context.Context, *CloseKeyRequest, ...dcerpc.CallOption) (*CloseKeyResponse, error)

	// ApiQueryInfoKey operation.
	QueryInfoKey(context.Context, *QueryInfoKeyRequest, ...dcerpc.CallOption) (*QueryInfoKeyResponse, error)

	// ApiSetKeySecurity operation.
	SetKeySecurity(context.Context, *SetKeySecurityRequest, ...dcerpc.CallOption) (*SetKeySecurityResponse, error)

	// ApiGetKeySecurity operation.
	GetKeySecurity(context.Context, *GetKeySecurityRequest, ...dcerpc.CallOption) (*GetKeySecurityResponse, error)

	// ApiOpenGroup operation.
	OpenGroup(context.Context, *OpenGroupRequest, ...dcerpc.CallOption) (*OpenGroupResponse, error)

	// ApiCreateGroup operation.
	CreateGroup(context.Context, *CreateGroupRequest, ...dcerpc.CallOption) (*CreateGroupResponse, error)

	// ApiDeleteGroup operation.
	DeleteGroup(context.Context, *DeleteGroupRequest, ...dcerpc.CallOption) (*DeleteGroupResponse, error)

	// ApiCloseGroup operation.
	CloseGroup(context.Context, *CloseGroupRequest, ...dcerpc.CallOption) (*CloseGroupResponse, error)

	// ApiGetGroupState operation.
	GetGroupState(context.Context, *GetGroupStateRequest, ...dcerpc.CallOption) (*GetGroupStateResponse, error)

	// ApiSetGroupName operation.
	SetGroupName(context.Context, *SetGroupNameRequest, ...dcerpc.CallOption) (*SetGroupNameResponse, error)

	// ApiGetGroupId operation.
	GetGroupID(context.Context, *GetGroupIDRequest, ...dcerpc.CallOption) (*GetGroupIDResponse, error)

	// ApiGetNodeId operation.
	GetNodeID(context.Context, *GetNodeIDRequest, ...dcerpc.CallOption) (*GetNodeIDResponse, error)

	// ApiOnlineGroup operation.
	OnlineGroup(context.Context, *OnlineGroupRequest, ...dcerpc.CallOption) (*OnlineGroupResponse, error)

	// ApiOfflineGroup operation.
	OfflineGroup(context.Context, *OfflineGroupRequest, ...dcerpc.CallOption) (*OfflineGroupResponse, error)

	// ApiMoveGroup operation.
	MoveGroup(context.Context, *MoveGroupRequest, ...dcerpc.CallOption) (*MoveGroupResponse, error)

	// ApiMoveGroupToNode operation.
	MoveGroupToNode(context.Context, *MoveGroupToNodeRequest, ...dcerpc.CallOption) (*MoveGroupToNodeResponse, error)

	// ApiCreateGroupResourceEnum operation.
	CreateGroupResourceEnum(context.Context, *CreateGroupResourceEnumRequest, ...dcerpc.CallOption) (*CreateGroupResourceEnumResponse, error)

	// ApiSetGroupNodeList operation.
	SetGroupNodeList(context.Context, *SetGroupNodeListRequest, ...dcerpc.CallOption) (*SetGroupNodeListResponse, error)

	// ApiCreateNotify operation.
	CreateNotify(context.Context, *CreateNotifyRequest, ...dcerpc.CallOption) (*CreateNotifyResponse, error)

	// ApiCloseNotify operation.
	CloseNotify(context.Context, *CloseNotifyRequest, ...dcerpc.CallOption) (*CloseNotifyResponse, error)

	// ApiAddNotifyCluster operation.
	AddNotifyCluster(context.Context, *AddNotifyClusterRequest, ...dcerpc.CallOption) (*AddNotifyClusterResponse, error)

	// ApiAddNotifyNode operation.
	AddNotifyNode(context.Context, *AddNotifyNodeRequest, ...dcerpc.CallOption) (*AddNotifyNodeResponse, error)

	// ApiAddNotifyGroup operation.
	AddNotifyGroup(context.Context, *AddNotifyGroupRequest, ...dcerpc.CallOption) (*AddNotifyGroupResponse, error)

	// ApiAddNotifyResource operation.
	AddNotifyResource(context.Context, *AddNotifyResourceRequest, ...dcerpc.CallOption) (*AddNotifyResourceResponse, error)

	// ApiAddNotifyKey operation.
	AddNotifyKey(context.Context, *AddNotifyKeyRequest, ...dcerpc.CallOption) (*AddNotifyKeyResponse, error)

	// ApiReAddNotifyNode operation.
	ReAddNotifyNode(context.Context, *ReAddNotifyNodeRequest, ...dcerpc.CallOption) (*ReAddNotifyNodeResponse, error)

	// ApiReAddNotifyGroup operation.
	ReAddNotifyGroup(context.Context, *ReAddNotifyGroupRequest, ...dcerpc.CallOption) (*ReAddNotifyGroupResponse, error)

	// ApiReAddNotifyResource operation.
	ReAddNotifyResource(context.Context, *ReAddNotifyResourceRequest, ...dcerpc.CallOption) (*ReAddNotifyResourceResponse, error)

	// ApiGetNotify operation.
	GetNotify(context.Context, *GetNotifyRequest, ...dcerpc.CallOption) (*GetNotifyResponse, error)

	// ApiOpenNode operation.
	OpenNode(context.Context, *OpenNodeRequest, ...dcerpc.CallOption) (*OpenNodeResponse, error)

	// ApiCloseNode operation.
	CloseNode(context.Context, *CloseNodeRequest, ...dcerpc.CallOption) (*CloseNodeResponse, error)

	// ApiGetNodeState operation.
	GetNodeState(context.Context, *GetNodeStateRequest, ...dcerpc.CallOption) (*GetNodeStateResponse, error)

	// ApiPauseNode operation.
	PauseNode(context.Context, *PauseNodeRequest, ...dcerpc.CallOption) (*PauseNodeResponse, error)

	// ApiResumeNode operation.
	ResumeNode(context.Context, *ResumeNodeRequest, ...dcerpc.CallOption) (*ResumeNodeResponse, error)

	// ApiEvictNode operation.
	EvictNode(context.Context, *EvictNodeRequest, ...dcerpc.CallOption) (*EvictNodeResponse, error)

	// ApiNodeResourceControl operation.
	NodeResourceControl(context.Context, *NodeResourceControlRequest, ...dcerpc.CallOption) (*NodeResourceControlResponse, error)

	// ApiResourceControl operation.
	ResourceControl(context.Context, *ResourceControlRequest, ...dcerpc.CallOption) (*ResourceControlResponse, error)

	// ApiNodeResourceTypeControl operation.
	NodeResourceTypeControl(context.Context, *NodeResourceTypeControlRequest, ...dcerpc.CallOption) (*NodeResourceTypeControlResponse, error)

	// ApiResourceTypeControl operation.
	ResourceTypeControl(context.Context, *ResourceTypeControlRequest, ...dcerpc.CallOption) (*ResourceTypeControlResponse, error)

	// ApiNodeGroupControl operation.
	NodeGroupControl(context.Context, *NodeGroupControlRequest, ...dcerpc.CallOption) (*NodeGroupControlResponse, error)

	// ApiGroupControl operation.
	GroupControl(context.Context, *GroupControlRequest, ...dcerpc.CallOption) (*GroupControlResponse, error)

	// ApiNodeNodeControl operation.
	NodeNodeControl(context.Context, *NodeNodeControlRequest, ...dcerpc.CallOption) (*NodeNodeControlResponse, error)

	// ApiNodeControl operation.
	NodeControl(context.Context, *NodeControlRequest, ...dcerpc.CallOption) (*NodeControlResponse, error)

	// Opnum80NotUsedOnWire operation.
	// Opnum80NotUsedOnWire

	// ApiOpenNetwork operation.
	OpenNetwork(context.Context, *OpenNetworkRequest, ...dcerpc.CallOption) (*OpenNetworkResponse, error)

	// ApiCloseNetwork operation.
	CloseNetwork(context.Context, *CloseNetworkRequest, ...dcerpc.CallOption) (*CloseNetworkResponse, error)

	// ApiGetNetworkState operation.
	GetNetworkState(context.Context, *GetNetworkStateRequest, ...dcerpc.CallOption) (*GetNetworkStateResponse, error)

	// ApiSetNetworkName operation.
	SetNetworkName(context.Context, *SetNetworkNameRequest, ...dcerpc.CallOption) (*SetNetworkNameResponse, error)

	// ApiCreateNetworkEnum operation.
	CreateNetworkEnum(context.Context, *CreateNetworkEnumRequest, ...dcerpc.CallOption) (*CreateNetworkEnumResponse, error)

	// ApiGetNetworkId operation.
	GetNetworkID(context.Context, *GetNetworkIDRequest, ...dcerpc.CallOption) (*GetNetworkIDResponse, error)

	// ApiSetNetworkPriorityOrder operation.
	SetNetworkPriorityOrder(context.Context, *SetNetworkPriorityOrderRequest, ...dcerpc.CallOption) (*SetNetworkPriorityOrderResponse, error)

	// ApiNodeNetworkControl operation.
	NodeNetworkControl(context.Context, *NodeNetworkControlRequest, ...dcerpc.CallOption) (*NodeNetworkControlResponse, error)

	// ApiNetworkControl operation.
	NetworkControl(context.Context, *NetworkControlRequest, ...dcerpc.CallOption) (*NetworkControlResponse, error)

	// ApiAddNotifyNetwork operation.
	AddNotifyNetwork(context.Context, *AddNotifyNetworkRequest, ...dcerpc.CallOption) (*AddNotifyNetworkResponse, error)

	// ApiReAddNotifyNetwork operation.
	ReAddNotifyNetwork(context.Context, *ReAddNotifyNetworkRequest, ...dcerpc.CallOption) (*ReAddNotifyNetworkResponse, error)

	// ApiOpenNetInterface operation.
	OpenNetInterface(context.Context, *OpenNetInterfaceRequest, ...dcerpc.CallOption) (*OpenNetInterfaceResponse, error)

	// ApiCloseNetInterface operation.
	CloseNetInterface(context.Context, *CloseNetInterfaceRequest, ...dcerpc.CallOption) (*CloseNetInterfaceResponse, error)

	// ApiGetNetInterfaceState operation.
	GetNetInterfaceState(context.Context, *GetNetInterfaceStateRequest, ...dcerpc.CallOption) (*GetNetInterfaceStateResponse, error)

	// ApiGetNetInterface operation.
	GetNetInterface(context.Context, *GetNetInterfaceRequest, ...dcerpc.CallOption) (*GetNetInterfaceResponse, error)

	// ApiGetNetInterfaceId operation.
	GetNetInterfaceID(context.Context, *GetNetInterfaceIDRequest, ...dcerpc.CallOption) (*GetNetInterfaceIDResponse, error)

	// ApiNodeNetInterfaceControl operation.
	NodeNetInterfaceControl(context.Context, *NodeNetInterfaceControlRequest, ...dcerpc.CallOption) (*NodeNetInterfaceControlResponse, error)

	// ApiNetInterfaceControl operation.
	NetInterfaceControl(context.Context, *NetInterfaceControlRequest, ...dcerpc.CallOption) (*NetInterfaceControlResponse, error)

	// ApiAddNotifyNetInterface operation.
	AddNotifyNetInterface(context.Context, *AddNotifyNetInterfaceRequest, ...dcerpc.CallOption) (*AddNotifyNetInterfaceResponse, error)

	// ApiReAddNotifyNetInterface operation.
	ReAddNotifyNetInterface(context.Context, *ReAddNotifyNetInterfaceRequest, ...dcerpc.CallOption) (*ReAddNotifyNetInterfaceResponse, error)

	// ApiCreateNodeEnum operation.
	CreateNodeEnum(context.Context, *CreateNodeEnumRequest, ...dcerpc.CallOption) (*CreateNodeEnumResponse, error)

	// ApiGetClusterVersion2 operation.
	GetClusterVersion2(context.Context, *GetClusterVersion2Request, ...dcerpc.CallOption) (*GetClusterVersion2Response, error)

	// ApiCreateResTypeEnum operation.
	CreateRestrictionTypeEnum(context.Context, *CreateRestrictionTypeEnumRequest, ...dcerpc.CallOption) (*CreateRestrictionTypeEnumResponse, error)

	// ApiBackupClusterDatabase operation.
	BackupClusterDatabase(context.Context, *BackupClusterDatabaseRequest, ...dcerpc.CallOption) (*BackupClusterDatabaseResponse, error)

	// ApiNodeClusterControl operation.
	NodeClusterControl(context.Context, *NodeClusterControlRequest, ...dcerpc.CallOption) (*NodeClusterControlResponse, error)

	// ApiClusterControl operation.
	ClusterControl(context.Context, *ClusterControlRequest, ...dcerpc.CallOption) (*ClusterControlResponse, error)

	// ApiUnblockGetNotifyCall operation.
	UnblockGetNotifyCall(context.Context, *UnblockGetNotifyCallRequest, ...dcerpc.CallOption) (*UnblockGetNotifyCallResponse, error)

	// ApiSetServiceAccountPassword operation.
	SetServiceAccountPassword(context.Context, *SetServiceAccountPasswordRequest, ...dcerpc.CallOption) (*SetServiceAccountPasswordResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error
}

// Cluster structure represents HCLUSTER_RPC RPC structure.
type Cluster dcetypes.ContextHandle

func (o *Cluster) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Cluster) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Cluster) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Cluster) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Node structure represents HNODE_RPC RPC structure.
type Node dcetypes.ContextHandle

func (o *Node) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Node) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Node) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Node) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Group structure represents HGROUP_RPC RPC structure.
type Group dcetypes.ContextHandle

func (o *Group) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Group) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Group) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Group) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Resource structure represents HRES_RPC RPC structure.
type Resource dcetypes.ContextHandle

func (o *Resource) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Resource) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Resource) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Resource) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Key structure represents HKEY_RPC RPC structure.
type Key dcetypes.ContextHandle

func (o *Key) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Key) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Key) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Key) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Notify structure represents HNOTIFY_RPC RPC structure.
type Notify dcetypes.ContextHandle

func (o *Notify) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Notify) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Notify) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Notify) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Network structure represents HNETWORK_RPC RPC structure.
type Network dcetypes.ContextHandle

func (o *Network) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Network) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Network) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Network) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// NetInterface structure represents HNETINTERFACE_RPC RPC structure.
type NetInterface dcetypes.ContextHandle

func (o *NetInterface) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *NetInterface) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *NetInterface) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *NetInterface) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// SecurityDescriptor structure represents RPC_SECURITY_DESCRIPTOR RPC structure.
//
// The RPC_SECURITY_DESCRIPTOR structure is a container for passing a security descriptor
// that can be marshaled and unmarshaled by RPC. In this protocol it is part of the
// RPC_SECURITY_ATTRIBUTES structure (section 2.2.3.2).
type SecurityDescriptor struct {
	// lpSecurityDescriptor: A variable-length buffer that contains a security descriptor
	// in self-relative form.
	SecurityDescriptor []byte `idl:"name:lpSecurityDescriptor;size_is:(cbInSecurityDescriptor);length_is:(cbOutSecurityDescriptor)" json:"security_descriptor"`
	// cbInSecurityDescriptor: The size, in bytes, of memory that is allocated for the security
	// descriptor. If no security descriptor is specified, this field MUST be zero.
	InSecurityDescriptorLength uint32 `idl:"name:cbInSecurityDescriptor" json:"in_security_descriptor_length"`
	// cbOutSecurityDescriptor: The number of bytes of the lpSecurityDescriptor to be transmitted.
	// If no security descriptor is specified, this field MUST be zero.
	OutSecurityDescriptorLength uint32 `idl:"name:cbOutSecurityDescriptor" json:"out_security_descriptor_length"`
}

func (o *SecurityDescriptor) xxx_PreparePayload(ctx context.Context) error {
	if o.SecurityDescriptor != nil && o.InSecurityDescriptorLength == 0 {
		o.InSecurityDescriptorLength = uint32(len(o.SecurityDescriptor))
	}
	if o.SecurityDescriptor != nil && o.OutSecurityDescriptorLength == 0 {
		o.OutSecurityDescriptorLength = uint32(len(o.SecurityDescriptor))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityDescriptor) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.SecurityDescriptor != nil || o.InSecurityDescriptorLength > 0 {
		_ptr_lpSecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.InSecurityDescriptorLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := uint64(o.OutSecurityDescriptorLength)
			if dimLength1 > sizeInfo[0] {
				dimLength1 = sizeInfo[0]
			} else {
				sizeInfo[0] = dimLength1
			}
			if err := w.WriteSize(0); err != nil {
				return err
			}
			if err := w.WriteSize(dimLength1); err != nil {
				return err
			}
			for i1 := range o.SecurityDescriptor {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.SecurityDescriptor[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.SecurityDescriptor); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SecurityDescriptor, _ptr_lpSecurityDescriptor); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.InSecurityDescriptorLength); err != nil {
		return err
	}
	if err := w.WriteData(o.OutSecurityDescriptorLength); err != nil {
		return err
	}
	return nil
}
func (o *SecurityDescriptor) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_lpSecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.SecurityDescriptor", sizeInfo[0])
		}
		o.SecurityDescriptor = make([]byte, sizeInfo[0])
		for i1 := range o.SecurityDescriptor {
			i1 := i1
			if err := w.ReadData(&o.SecurityDescriptor[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_lpSecurityDescriptor := func(ptr interface{}) { o.SecurityDescriptor = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.SecurityDescriptor, _s_lpSecurityDescriptor, _ptr_lpSecurityDescriptor); err != nil {
		return err
	}
	if err := w.ReadData(&o.InSecurityDescriptorLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutSecurityDescriptorLength); err != nil {
		return err
	}
	return nil
}

// SecurityAttributes structure represents RPC_SECURITY_ATTRIBUTES RPC structure.
//
// The RPC_SECURITY_ATTRIBUTES structure represents security attributes that can be
// marshaled and unmarshaled by RPC.
//
// The RPC_SECURITY_ATTRIBUTES is used by a client to indicate the security attributes
// that are assigned when creating a new cluster registry key, as specified in section
// 3.1.4.2.30.
type SecurityAttributes struct {
	// nLength: The length of the structure, in bytes.
	Length uint32 `idl:"name:nLength" json:"length"`
	// RpcSecurityDescriptor: A self-relative security descriptor that can be marshaled
	// and unmarshaled by RPC, as specified in section 2.2.3.1.
	SecurityDescriptor *SecurityDescriptor `idl:"name:RpcSecurityDescriptor" json:"security_descriptor"`
	// bInheritHandle: Any nonzero value if a new spawned process inherits the handle; however,
	// because cluster registry keys are not inheritable, this field MUST be set to zero
	// for use in ApiCreateKey (section 3.1.4.2.30).
	InheritHandle int32 `idl:"name:bInheritHandle" json:"inherit_handle"`
}

func (o *SecurityAttributes) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityAttributes) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if o.SecurityDescriptor != nil {
		if err := o.SecurityDescriptor.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SecurityDescriptor{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.InheritHandle); err != nil {
		return err
	}
	return nil
}
func (o *SecurityAttributes) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if o.SecurityDescriptor == nil {
		o.SecurityDescriptor = &SecurityDescriptor{}
	}
	if err := o.SecurityDescriptor.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.InheritHandle); err != nil {
		return err
	}
	return nil
}

// ClusterOperationalVersionInfo structure represents CLUSTER_OPERATIONAL_VERSION_INFO RPC structure.
//
// The CLUSTER_OPERATIONAL_VERSION_INFO structure contains information about the versions
// of cluster software with which all nodes in the cluster are compatible.
type ClusterOperationalVersionInfo struct {
	// dwSize: The size, in bytes, of this structure. MUST be set to 20 bytes.
	Size uint32 `idl:"name:dwSize" json:"size"`
	// dwClusterHighestVersion: The highest version of cluster software with which all nodes
	// in the cluster are compatible. The upper 2 bytes of dwClusterHighestVersion identify
	// the cluster software internal major version number, and the lower 2 bytes identify
	// the cluster software internal minor version number. Note that the internal version
	// and build numbers are not necessarily identical to the operating system version and
	// build numbers. The dwClusterHighestVersion member SHOULD<23> be set to one of the
	// following values.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x000100e0 | One or more nodes support an internal version that is no higher than 0x0001 with |
	//	|            | internal build number 0x00e0.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x000200e0 | One or more nodes support an internal version that is no higher than 0x0002 with |
	//	|            | internal build number 0x00e0.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00030893 | One or more nodes support an internal version that is no higher than 0x0003 with |
	//	|            | internal build number 0x0893.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00040ECE | One or more nodes support an internal version that is no higher than 0x0004 with |
	//	|            | internal build number 0x0ECE.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00051771 | One or more nodes support an internal version that is no higher than 0x0005 with |
	//	|            | internal build number 0x1771.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00061DB0 | One or more nodes support an internal version that is no higher than 0x0006 with |
	//	|            | internal build number 0x1DB0.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00061DB1 | One or more nodes support an internal version that is no higher than 0x0006 with |
	//	|            | internal build number 0x1DB1.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x000723F0 | One or more nodes support an internal version that is no higher than 0x0007 with |
	//	|            | internal build number 0x23F0.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00082580 | One or more nodes support an internal version that is no higher than 0x0008 with |
	//	|            | internal build number 0x2580.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00090003 | One or more nodes support an internal version that is no higher than 0x0009 with |
	//	|            | internal build number 0x0003.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x000a0001 | One or more nodes support an internal version that is no higher than 0x000a with |
	//	|            | internal build number 0x0001.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x000a0002 | One or more nodes support an internal version that is no higher than 0x000a with |
	//	|            | internal build number 0x0002.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x000a0003 | One or more nodes support an internal version that is no higher than 0x000a with |
	//	|            | internal build number 0x0003.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x000b0003 | One or more nodes support an internal version that is no higher than 0x000b with |
	//	|            | internal build number 0x0003.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x000c0004 | One or more nodes support an internal version that is no higher than 0x000c with |
	//	|            | internal build number 0x0004.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	ClusterHighestVersion uint32 `idl:"name:dwClusterHighestVersion" json:"cluster_highest_version"`
	// dwClusterLowestVersion: The lowest version of cluster software with which all nodes
	// in the cluster are compatible. The upper 2 bytes of dwClusterLowestVersion identify
	// the cluster software internal major version number, and the lower 2 bytes of dwClusterLowestVersion
	// identify the cluster software internal minor version number. The dwClusterLowestVersion
	// member SHOULD<24> be set to one of the following values.
	//
	//	+------------+----------------------------------------------------------------------------------+
	//	|            |                                                                                  |
	//	|   VALUE    |                                     MEANING                                      |
	//	|            |                                                                                  |
	//	+------------+----------------------------------------------------------------------------------+
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x000100e0 | One or more nodes support an internal version that is no lower than 0x0001 with  |
	//	|            | internal build number 0x00e0.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x000200e0 | One or more nodes support an internal version that is no lower than 0x0002 with  |
	//	|            | internal build number 0x00e0.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00030893 | One or more nodes support an internal version that is no lower than 0x0003 with  |
	//	|            | internal build number 0x0893.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00040ECE | One or more nodes support an internal version that is no lower than 0x0004 with  |
	//	|            | internal build number 0x0ECE.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00051771 | One or more nodes support an internal version that is no lower than 0x0005 with  |
	//	|            | internal build number 0x1771.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00061DB0 | One or more nodes support an internal version that is no lower than 0x0006 with  |
	//	|            | internal build number 0x1DB0.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00061DB1 | One or more nodes support an internal version that is no lower than 0x0003 with  |
	//	|            | internal build number 0x1DB1.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x000723F0 | One or more nodes support an internal version that is no lower than 0x0007 with  |
	//	|            | internal build number 0x23F0.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00082580 | One or more nodes support an internal version that is no lower than 0x0008 with  |
	//	|            | internal build number 0x2580.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x00090003 | One or more nodes support an internal version that is no lower than 0x0009 with  |
	//	|            | internal build number 0x0003.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x000a0001 | One or more nodes support an internal version that is no lower than 0x000a with  |
	//	|            | internal build number 0x0001.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x000a0002 | One or more nodes support an internal version that is no lower than 0x000a with  |
	//	|            | internal build number 0x0002.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x000a0003 | One or more nodes support an internal version that is no higher than 0x000a with |
	//	|            | internal build number 0x0003.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x000b0003 | One or more nodes support an internal version that is no lower than 0x000b with  |
	//	|            | internal build number 0x0003.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	//	| 0x000c0004 | One or more nodes support an internal version that is no lower than 0x000c with  |
	//	|            | internal build number 0x0004.                                                    |
	//	+------------+----------------------------------------------------------------------------------+
	ClusterLowestVersion uint32 `idl:"name:dwClusterLowestVersion" json:"cluster_lowest_version"`
	// dwFlags:  The flags that identify the characteristics of the cluster operational
	// version. The dwFlags member MUST be set to one of the following values.
	//
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                            |                                                                                  |
	//	|                   VALUE                    |                                     MEANING                                      |
	//	|                                            |                                                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000                                 | All nodes in the cluster are running the same version of the cluster software.   |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLUSTER_VERSION_FLAG_MIXED_MODE 0x00000001 | The cluster is configured with nodes that are running different versions of the  |
	//	|                                            | cluster software.                                                                |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:dwFlags" json:"flags"`
	// dwReserved: Reserved. MUST be set to 0 and ignored by the client.
	_ uint32 `idl:"name:dwReserved"`
}

func (o *ClusterOperationalVersionInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClusterOperationalVersionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Size); err != nil {
		return err
	}
	if err := w.WriteData(o.ClusterHighestVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.ClusterLowestVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	// reserved dwReserved
	if err := w.WriteData(uint32(0)); err != nil {
		return err
	}
	return nil
}
func (o *ClusterOperationalVersionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Size); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClusterHighestVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.ClusterLowestVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	// reserved dwReserved
	var _dwReserved uint32
	if err := w.ReadData(&_dwReserved); err != nil {
		return err
	}
	return nil
}

// EnumEntry structure represents ENUM_ENTRY RPC structure.
//
// The ENUM_ENTRY structure contains information about a single element in an ENUM_LIST
// (section 2.2.3.5). An ENUM_ENTRY contains the name of a cluster object.
type EnumEntry struct {
	// Type: Specifies the type of the object that is represented by the list element. This
	// member MUST be set to one of the following values.
	//
	// If the ENUM_LIST is returned by a call to ApiCreateEnum (section 3.1.4.2.8) or ApiCreateEnumEx
	// (section 3.1.4.2.124), Type is set to one of the following values.
	//
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                                |                                                                                  |
	//	|                     VALUE                      |                                     MEANING                                      |
	//	|                                                |                                                                                  |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLUSTER_ENUM_NODE 0x00000001                   | If returned by ApiCreateEnum, this ENUM_ENTRY list element contains the name     |
	//	|                                                | of a cluster node. If returned by ApiCreateEnumEx, this ENUM_ENTRY list element  |
	//	|                                                | contains the name or ID, as specified in ApiCreateEnumEx, of a cluster node.     |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLUSTER_ENUM_RESTYPE 0x00000002                | If returned by ApiCreateEnum, this ENUM_ENTRY list element contains the name of  |
	//	|                                                | a cluster resource type. If returned by ApiCreateEnumEx, this ENUM_ENTRY list    |
	//	|                                                | element contains the name or ID, as specified in ApiCreateEnumEx, of a cluster   |
	//	|                                                | resource type.                                                                   |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLUSTER_ENUM_RESOURCE 0x00000004               | If returned by ApiCreateEnum, this ENUM_ENTRY list element contains the name of  |
	//	|                                                | a cluster resource. If returned by ApiCreateEnumEx, this ENUM_ENTRY list element |
	//	|                                                | contains the name or ID, as specified in ApiCreateEnumEx, of a cluster resource. |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLUSTER_ENUM_GROUP 0x00000008                  | If returned by ApiCreateEnum, this ENUM_ENTRY list element contains the name of  |
	//	|                                                | a cluster group. If returned by ApiCreateEnumEx, this ENUM_ENTRY list element    |
	//	|                                                | contains the name or ID, as specified in ApiCreateEnumEx, of a cluster group.    |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLUSTER_ENUM_NETWORK 0x00000010                | If returned by ApiCreateEnum, this ENUM_ENTRY list element contains the name of  |
	//	|                                                | a cluster network. If returned by ApiCreateEnumEx, this ENUM_ENTRY list element  |
	//	|                                                | contains the name or ID, as specified in ApiCreateEnumEx, of a cluster network.  |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLUSTER_ENUM_NETINTERFACE 0x00000020           | If returned by ApiCreateEnum, this ENUM_ENTRY list element contains the name     |
	//	|                                                | of a cluster network interface. If returned by ApiCreateEnumEx, this ENUM_ENTRY  |
	//	|                                                | list element contains the name or ID, as specified in ApiCreateEnumEx, of a      |
	//	|                                                | cluster network interface.                                                       |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLUSTER_ENUM_SHARED_VOLUME_RESOURCE 0x40000000 | If returned by ApiCreateEnum, this ENUM_ENTRY list element contains the name of  |
	//	|                                                | a cluster shared volume. If returned by ApiCreateEnumEx, this ENUM_ENTRY list    |
	//	|                                                | element contains the name or ID, as specified in ApiCreateEnumEx, of a cluster   |
	//	|                                                | shared volume.                                                                   |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLUSTER_ENUM_INTERNAL_NETWORK 0x80000000       | If returned by ApiCreateEnum, this ENUM_ENTRY list element contains the name     |
	//	|                                                | of a cluster network used only for internal cluster communications. If returned  |
	//	|                                                | by ApiCreateEnumEx, this ENUM_ENTRY list element contains the name or ID,        |
	//	|                                                | as specified in ApiCreateEnumEx, of a cluster network used only for internal     |
	//	|                                                | cluster communications.                                                          |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// If the ENUM_LIST is returned by a call to ApiCreateResEnum (section 3.1.4.2.23),
	// Type is set to one of the following values.
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                           |                                                                                  |
	//	|                   VALUE                   |                                     MEANING                                      |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLUSTER_RESOURCE_ENUM_DEPENDS 0x00000001  | This ENUM_ENTRY list element contains the name of a cluster resource that is     |
	//	|                                           | depended upon, as described in Resources (section 3.1.1.1.1) by the resource     |
	//	|                                           | designated in the call to ApiCreateResEnum.                                      |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLUSTER_RESOURCE_ENUM_PROVIDES 0x00000002 | This ENUM_ENTRY list element contains the name of a cluster resource that        |
	//	|                                           | depends upon, as described in Resources by the resource designated in the call   |
	//	|                                           | to ApiCreateResEnum.                                                             |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLUSTER_RESOURCE_ENUM_NODES 0x00000004    | This ENUM_ENTRY list element contains the name of a cluster node that can host   |
	//	|                                           | the resource designated in the call to ApiCreateResEnum.                         |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//
	// If the ENUM_LIST is returned by a call to ApiCreateGroupResourceEnum (section 3.1.4.2.54),
	// Type is set to one of the following values.
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                        |                                                                                  |
	//	|                 VALUE                  |                                     MEANING                                      |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| CLUSTER_GROUP_ENUM_CONTAINS 0x00000001 | This ENUM_ENTRY list element contains the name of a cluster resource that is     |
	//	|                                        | contained in the group designated in the call to ApiCreateGroupResourceEnum.     |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| CLUSTER_GROUP_ENUM_NODES 0x00000002    | This ENUM_ENTRY list element contains the name of a cluster node that can host   |
	//	|                                        | the group designated in the call to ApiCreateGroupResourceEnum.                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//
	// If the ENUM_LIST is returned by a call to ApiCreateNetworkEnum (section 3.1.4.2.85),
	// Type is set to one of the following values.
	//
	//	+-----------------------------------------------+--------------------------------------------------------------------------------+
	//	|                                               |                                                                                |
	//	|                     VALUE                     |                                    MEANING                                     |
	//	|                                               |                                                                                |
	//	+-----------------------------------------------+--------------------------------------------------------------------------------+
	//	+-----------------------------------------------+--------------------------------------------------------------------------------+
	//	| CLUSTER_NETWORK_ENUM_NETINTERFACES 0x00000001 | This ENUM_ENTRY list element contains the name of a cluster network interface. |
	//	+-----------------------------------------------+--------------------------------------------------------------------------------+
	//
	// If the ENUM_LIST is returned by a call to ApiCreateNodeEnum (section 3.1.4.2.101)
	// or ApiCreateNodeEnumEx (section 3.1.4.2.123), Type is set to one of the following
	// values.
	//
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                            |                                                                                  |
	//	|                   VALUE                    |                                     MEANING                                      |
	//	|                                            |                                                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLUSTER_NODE_ENUM_NETINTERFACES 0x00000001 | This ENUM_ENTRY list element contains the name of a cluster network interface.   |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| CLUSTER_NODE_ENUM_GROUPS 0x00000002        | If returned by ApiCreateNodeEnum, this ENUM_ENTRY list element contains the      |
	//	|                                            | name of a cluster group. If returned by ApiCreateNodeEnumEx, this ENUM_ENTRY     |
	//	|                                            | list element contains the name or ID, as specified in ApiCreateNodeEnumEx, of a  |
	//	|                                            | cluster group.                                                                   |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//
	// If the ENUM_LIST is returned by a call to ApiCreateResTypeEnum (section 3.1.4.2.103),
	// Type is set to one of the following values.
	//
	//	+-------------------------------------------------+-----------------------------------------------------------------------+
	//	|                                                 |                                                                       |
	//	|                      VALUE                      |                                MEANING                                |
	//	|                                                 |                                                                       |
	//	+-------------------------------------------------+-----------------------------------------------------------------------+
	//	+-------------------------------------------------+-----------------------------------------------------------------------+
	//	| CLUSTER_RESOURCE_TYPE_ENUM_NODES 0x00000001     | This ENUM_ENTRY list element contains the name of a cluster node.     |
	//	+-------------------------------------------------+-----------------------------------------------------------------------+
	//	| CLUSTER_RESOURCE_TYPE_ENUM_RESOURCES 0x00000002 | This ENUM_ENTRY list element contains the name of a cluster resource. |
	//	+-------------------------------------------------+-----------------------------------------------------------------------+
	//
	// If the ENUM_LIST is returned by a call to ApiCreateNetInterfaceEnum, as specified
	// in section 3.1.4.2.163, Type is set to one of the following values.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                                      |                                                                                  |
	//	|                VALUE                 |                                     MEANING                                      |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| CLUSTER_ENUM_NETINTERFACE 0x00000020 | If returned by ApiCreateNetInterfaceEnum, this ENUM_ENTRY list element contains  |
	//	|                                      | the list of cluster network interfaces.                                          |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	Type uint32 `idl:"name:Type" json:"type"`
	// Name: If the ENUM_LIST is returned by a call to ApiCreateNodeEnumEx, it contains
	// the return data specified in section 3.1.4.2.123.
	//
	// If the ENUM_LIST is returned by a call to ApiCreateEnumEx, it contains the return
	// data specified in section 3.1.4.2.124.
	Name string `idl:"name:Name;string" json:"name"`
}

func (o *EnumEntry) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *EnumEntry) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	if o.Name != "" {
		_ptr_Name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_Name); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *EnumEntry) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	_ptr_Name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_Name := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_Name, _ptr_Name); err != nil {
		return err
	}
	return nil
}

// EnumList structure represents ENUM_LIST RPC structure.
//
// The ENUM_LIST structure is a container for a list of ENUM_ENTRY (section 2.2.3.4)
// structures. An ENUM_LIST encapsulates the results of a query that is performed on
// the cluster state. The semantics of the collection of named cluster objects depends
// on the query that is performed. For example, the ApiCreateEnum (section 3.1.4.2.8)
// method can be used to query a list of nodes in the cluster or a list of resources
// in the cluster. The result of either query is returned to the client as an ENUM_LIST.
type EnumList struct {
	// EntryCount: An unsigned 32-bit integer. The number of elements in the field Entry.
	EntryCount uint32 `idl:"name:EntryCount" json:"entry_count"`
	// Entry: An array of ENUM_ENTRY structures that represents the contents of the list.
	Entry []*EnumEntry `idl:"name:Entry;size_is:(EntryCount)" json:"entry"`
}

func (o *EnumList) xxx_PreparePayload(ctx context.Context) error {
	if o.Entry != nil && o.EntryCount == 0 {
		o.EntryCount = uint32(len(o.Entry))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *EnumList) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.EntryCount)
	return []uint64{
		dimSize1,
	}
}
func (o *EnumList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntryCount); err != nil {
		return err
	}
	for i1 := range o.Entry {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.Entry[i1] != nil {
			if err := o.Entry[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&EnumEntry{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Entry); uint64(i1) < sizeInfo[0]; i1++ {
		if err := (&EnumEntry{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *EnumList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntryCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.EntryCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.EntryCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.Entry", sizeInfo[0])
	}
	o.Entry = make([]*EnumEntry, sizeInfo[0])
	for i1 := range o.Entry {
		i1 := i1
		if o.Entry[i1] == nil {
			o.Entry[i1] = &EnumEntry{}
		}
		if err := o.Entry[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// ClusterSetPasswordStatus structure represents IDL_CLUSTER_SET_PASSWORD_STATUS RPC structure.
//
// The IDL_CLUSTER_SET_PASSWORD_STATUS structure contains information about the results
// of a service account password change, as specified in section 3.1.4.2.108, on a particular
// node in the cluster.
type ClusterSetPasswordStatus struct {
	// NodeId: A 32-bit integer containing the ID of a configured node in the cluster.
	NodeID uint32 `idl:"name:NodeId" json:"node_id"`
	// SetAttempted: A Boolean where TRUE indicates that the password change was attempted
	// on this node; any other value indicates that no attempt was made.
	SetAttempted bool `idl:"name:SetAttempted" json:"set_attempted"`
	// ReturnStatus: The Win32 error code associated with the attempt. This field MUST be
	// ignored if the SetAttempted field is FALSE.
	ReturnStatus uint32 `idl:"name:ReturnStatus" json:"return_status"`
}

func (o *ClusterSetPasswordStatus) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ClusterSetPasswordStatus) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.NodeID); err != nil {
		return err
	}
	if err := w.WriteData(o.SetAttempted); err != nil {
		return err
	}
	if err := w.WriteData(o.ReturnStatus); err != nil {
		return err
	}
	return nil
}
func (o *ClusterSetPasswordStatus) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.NodeID); err != nil {
		return err
	}
	if err := w.ReadData(&o.SetAttempted); err != nil {
		return err
	}
	if err := w.ReadData(&o.ReturnStatus); err != nil {
		return err
	}
	return nil
}

// ClusterSetPasswordFlags type represents IDL_CLUSTER_SET_PASSWORD_FLAGS RPC enumeration.
//
// The IDL_CLUSTER_SET_PASSWORD_FLAGS enumeration defines the possible values for the
// dwFlags parameter of the method ApiSetServiceAccountPassword (section 3.1.4.1.108).
// The valid constant values are as follows. All other values are reserved.
type ClusterSetPasswordFlags uint16

var (
	// IDL_CLUSTER_SET_PASSWORD_IGNORE_DOWN_NODES:  Indicates to the server to proceed
	// with the password change operation even if there are nodes configured in the cluster
	// that are currently in the ClusterNodeUp state, as specified in section 3.1.4.1.69.
	ClusterSetPasswordFlagsIgnoreDownNodes ClusterSetPasswordFlags = 1
)

func (o ClusterSetPasswordFlags) String() string {
	switch o {
	case ClusterSetPasswordFlagsIgnoreDownNodes:
		return "ClusterSetPasswordFlagsIgnoreDownNodes"
	}
	return "Invalid"
}

type xxx_DefaultClusapi2Client struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultClusapi2Client) OpenCluster(ctx context.Context, in *OpenClusterRequest, opts ...dcerpc.CallOption) (*OpenClusterResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenClusterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != nil {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) CloseCluster(ctx context.Context, in *CloseClusterRequest, opts ...dcerpc.CallOption) (*CloseClusterResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseClusterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) SetClusterName(ctx context.Context, in *SetClusterNameRequest, opts ...dcerpc.CallOption) (*SetClusterNameResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetClusterNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) GetClusterName(ctx context.Context, in *GetClusterNameRequest, opts ...dcerpc.CallOption) (*GetClusterNameResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetClusterNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) GetClusterVersion(ctx context.Context, in *GetClusterVersionRequest, opts ...dcerpc.CallOption) (*GetClusterVersionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetClusterVersionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) GetQuorumResource(ctx context.Context, in *GetQuorumResourceRequest, opts ...dcerpc.CallOption) (*GetQuorumResourceResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetQuorumResourceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) SetQuorumResource(ctx context.Context, in *SetQuorumResourceRequest, opts ...dcerpc.CallOption) (*SetQuorumResourceResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetQuorumResourceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) CreateEnum(ctx context.Context, in *CreateEnumRequest, opts ...dcerpc.CallOption) (*CreateEnumResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) OpenResource(ctx context.Context, in *OpenResourceRequest, opts ...dcerpc.CallOption) (*OpenResourceResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenResourceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != nil {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) CreateResource(ctx context.Context, in *CreateResourceRequest, opts ...dcerpc.CallOption) (*CreateResourceResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateResourceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != nil {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...dcerpc.CallOption) (*DeleteResourceResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteResourceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) CloseResource(ctx context.Context, in *CloseResourceRequest, opts ...dcerpc.CallOption) (*CloseResourceResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseResourceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) GetResourceState(ctx context.Context, in *GetResourceStateRequest, opts ...dcerpc.CallOption) (*GetResourceStateResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetResourceStateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) SetResourceName(ctx context.Context, in *SetResourceNameRequest, opts ...dcerpc.CallOption) (*SetResourceNameResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetResourceNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) GetResourceID(ctx context.Context, in *GetResourceIDRequest, opts ...dcerpc.CallOption) (*GetResourceIDResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetResourceIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) GetResourceType(ctx context.Context, in *GetResourceTypeRequest, opts ...dcerpc.CallOption) (*GetResourceTypeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetResourceTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) FailResource(ctx context.Context, in *FailResourceRequest, opts ...dcerpc.CallOption) (*FailResourceResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FailResourceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) OnlineResource(ctx context.Context, in *OnlineResourceRequest, opts ...dcerpc.CallOption) (*OnlineResourceResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OnlineResourceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) OfflineResource(ctx context.Context, in *OfflineResourceRequest, opts ...dcerpc.CallOption) (*OfflineResourceResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OfflineResourceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) AddResourceDependency(ctx context.Context, in *AddResourceDependencyRequest, opts ...dcerpc.CallOption) (*AddResourceDependencyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddResourceDependencyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) RemoveResourceDependency(ctx context.Context, in *RemoveResourceDependencyRequest, opts ...dcerpc.CallOption) (*RemoveResourceDependencyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveResourceDependencyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) CanResourceBeDependent(ctx context.Context, in *CanResourceBeDependentRequest, opts ...dcerpc.CallOption) (*CanResourceBeDependentResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CanResourceBeDependentResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) CreateRestrictionEnum(ctx context.Context, in *CreateRestrictionEnumRequest, opts ...dcerpc.CallOption) (*CreateRestrictionEnumResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateRestrictionEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) AddResourceNode(ctx context.Context, in *AddResourceNodeRequest, opts ...dcerpc.CallOption) (*AddResourceNodeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddResourceNodeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) RemoveResourceNode(ctx context.Context, in *RemoveResourceNodeRequest, opts ...dcerpc.CallOption) (*RemoveResourceNodeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveResourceNodeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) ChangeResourceGroup(ctx context.Context, in *ChangeResourceGroupRequest, opts ...dcerpc.CallOption) (*ChangeResourceGroupResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ChangeResourceGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) CreateResourceType(ctx context.Context, in *CreateResourceTypeRequest, opts ...dcerpc.CallOption) (*CreateResourceTypeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateResourceTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) DeleteResourceType(ctx context.Context, in *DeleteResourceTypeRequest, opts ...dcerpc.CallOption) (*DeleteResourceTypeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteResourceTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) GetRootKey(ctx context.Context, in *GetRootKeyRequest, opts ...dcerpc.CallOption) (*GetRootKeyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetRootKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != nil {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) CreateKey(ctx context.Context, in *CreateKeyRequest, opts ...dcerpc.CallOption) (*CreateKeyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != nil {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) OpenKey(ctx context.Context, in *OpenKeyRequest, opts ...dcerpc.CallOption) (*OpenKeyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != nil {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) EnumKey(ctx context.Context, in *EnumKeyRequest, opts ...dcerpc.CallOption) (*EnumKeyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) SetValue(ctx context.Context, in *SetValueRequest, opts ...dcerpc.CallOption) (*SetValueResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetValueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) DeleteValue(ctx context.Context, in *DeleteValueRequest, opts ...dcerpc.CallOption) (*DeleteValueResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteValueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) QueryValue(ctx context.Context, in *QueryValueRequest, opts ...dcerpc.CallOption) (*QueryValueResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryValueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) DeleteKey(ctx context.Context, in *DeleteKeyRequest, opts ...dcerpc.CallOption) (*DeleteKeyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) EnumValue(ctx context.Context, in *EnumValueRequest, opts ...dcerpc.CallOption) (*EnumValueResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumValueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) CloseKey(ctx context.Context, in *CloseKeyRequest, opts ...dcerpc.CallOption) (*CloseKeyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) QueryInfoKey(ctx context.Context, in *QueryInfoKeyRequest, opts ...dcerpc.CallOption) (*QueryInfoKeyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryInfoKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) SetKeySecurity(ctx context.Context, in *SetKeySecurityRequest, opts ...dcerpc.CallOption) (*SetKeySecurityResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetKeySecurityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) GetKeySecurity(ctx context.Context, in *GetKeySecurityRequest, opts ...dcerpc.CallOption) (*GetKeySecurityResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetKeySecurityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) OpenGroup(ctx context.Context, in *OpenGroupRequest, opts ...dcerpc.CallOption) (*OpenGroupResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != nil {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...dcerpc.CallOption) (*CreateGroupResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != nil {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...dcerpc.CallOption) (*DeleteGroupResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) CloseGroup(ctx context.Context, in *CloseGroupRequest, opts ...dcerpc.CallOption) (*CloseGroupResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) GetGroupState(ctx context.Context, in *GetGroupStateRequest, opts ...dcerpc.CallOption) (*GetGroupStateResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetGroupStateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) SetGroupName(ctx context.Context, in *SetGroupNameRequest, opts ...dcerpc.CallOption) (*SetGroupNameResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetGroupNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) GetGroupID(ctx context.Context, in *GetGroupIDRequest, opts ...dcerpc.CallOption) (*GetGroupIDResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetGroupIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) GetNodeID(ctx context.Context, in *GetNodeIDRequest, opts ...dcerpc.CallOption) (*GetNodeIDResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNodeIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) OnlineGroup(ctx context.Context, in *OnlineGroupRequest, opts ...dcerpc.CallOption) (*OnlineGroupResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OnlineGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) OfflineGroup(ctx context.Context, in *OfflineGroupRequest, opts ...dcerpc.CallOption) (*OfflineGroupResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OfflineGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) MoveGroup(ctx context.Context, in *MoveGroupRequest, opts ...dcerpc.CallOption) (*MoveGroupResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MoveGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) MoveGroupToNode(ctx context.Context, in *MoveGroupToNodeRequest, opts ...dcerpc.CallOption) (*MoveGroupToNodeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MoveGroupToNodeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) CreateGroupResourceEnum(ctx context.Context, in *CreateGroupResourceEnumRequest, opts ...dcerpc.CallOption) (*CreateGroupResourceEnumResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateGroupResourceEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) SetGroupNodeList(ctx context.Context, in *SetGroupNodeListRequest, opts ...dcerpc.CallOption) (*SetGroupNodeListResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetGroupNodeListResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) CreateNotify(ctx context.Context, in *CreateNotifyRequest, opts ...dcerpc.CallOption) (*CreateNotifyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateNotifyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != nil {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) CloseNotify(ctx context.Context, in *CloseNotifyRequest, opts ...dcerpc.CallOption) (*CloseNotifyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseNotifyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) AddNotifyCluster(ctx context.Context, in *AddNotifyClusterRequest, opts ...dcerpc.CallOption) (*AddNotifyClusterResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddNotifyClusterResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) AddNotifyNode(ctx context.Context, in *AddNotifyNodeRequest, opts ...dcerpc.CallOption) (*AddNotifyNodeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddNotifyNodeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) AddNotifyGroup(ctx context.Context, in *AddNotifyGroupRequest, opts ...dcerpc.CallOption) (*AddNotifyGroupResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddNotifyGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) AddNotifyResource(ctx context.Context, in *AddNotifyResourceRequest, opts ...dcerpc.CallOption) (*AddNotifyResourceResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddNotifyResourceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) AddNotifyKey(ctx context.Context, in *AddNotifyKeyRequest, opts ...dcerpc.CallOption) (*AddNotifyKeyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddNotifyKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) ReAddNotifyNode(ctx context.Context, in *ReAddNotifyNodeRequest, opts ...dcerpc.CallOption) (*ReAddNotifyNodeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ReAddNotifyNodeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) ReAddNotifyGroup(ctx context.Context, in *ReAddNotifyGroupRequest, opts ...dcerpc.CallOption) (*ReAddNotifyGroupResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ReAddNotifyGroupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) ReAddNotifyResource(ctx context.Context, in *ReAddNotifyResourceRequest, opts ...dcerpc.CallOption) (*ReAddNotifyResourceResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ReAddNotifyResourceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) GetNotify(ctx context.Context, in *GetNotifyRequest, opts ...dcerpc.CallOption) (*GetNotifyResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNotifyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) OpenNode(ctx context.Context, in *OpenNodeRequest, opts ...dcerpc.CallOption) (*OpenNodeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenNodeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != nil {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) CloseNode(ctx context.Context, in *CloseNodeRequest, opts ...dcerpc.CallOption) (*CloseNodeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseNodeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) GetNodeState(ctx context.Context, in *GetNodeStateRequest, opts ...dcerpc.CallOption) (*GetNodeStateResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNodeStateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) PauseNode(ctx context.Context, in *PauseNodeRequest, opts ...dcerpc.CallOption) (*PauseNodeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PauseNodeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) ResumeNode(ctx context.Context, in *ResumeNodeRequest, opts ...dcerpc.CallOption) (*ResumeNodeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ResumeNodeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) EvictNode(ctx context.Context, in *EvictNodeRequest, opts ...dcerpc.CallOption) (*EvictNodeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EvictNodeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) NodeResourceControl(ctx context.Context, in *NodeResourceControlRequest, opts ...dcerpc.CallOption) (*NodeResourceControlResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &NodeResourceControlResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) ResourceControl(ctx context.Context, in *ResourceControlRequest, opts ...dcerpc.CallOption) (*ResourceControlResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ResourceControlResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) NodeResourceTypeControl(ctx context.Context, in *NodeResourceTypeControlRequest, opts ...dcerpc.CallOption) (*NodeResourceTypeControlResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &NodeResourceTypeControlResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) ResourceTypeControl(ctx context.Context, in *ResourceTypeControlRequest, opts ...dcerpc.CallOption) (*ResourceTypeControlResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ResourceTypeControlResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) NodeGroupControl(ctx context.Context, in *NodeGroupControlRequest, opts ...dcerpc.CallOption) (*NodeGroupControlResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &NodeGroupControlResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) GroupControl(ctx context.Context, in *GroupControlRequest, opts ...dcerpc.CallOption) (*GroupControlResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GroupControlResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) NodeNodeControl(ctx context.Context, in *NodeNodeControlRequest, opts ...dcerpc.CallOption) (*NodeNodeControlResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &NodeNodeControlResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) NodeControl(ctx context.Context, in *NodeControlRequest, opts ...dcerpc.CallOption) (*NodeControlResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &NodeControlResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) OpenNetwork(ctx context.Context, in *OpenNetworkRequest, opts ...dcerpc.CallOption) (*OpenNetworkResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenNetworkResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != nil {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) CloseNetwork(ctx context.Context, in *CloseNetworkRequest, opts ...dcerpc.CallOption) (*CloseNetworkResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseNetworkResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) GetNetworkState(ctx context.Context, in *GetNetworkStateRequest, opts ...dcerpc.CallOption) (*GetNetworkStateResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNetworkStateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) SetNetworkName(ctx context.Context, in *SetNetworkNameRequest, opts ...dcerpc.CallOption) (*SetNetworkNameResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetNetworkNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) CreateNetworkEnum(ctx context.Context, in *CreateNetworkEnumRequest, opts ...dcerpc.CallOption) (*CreateNetworkEnumResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateNetworkEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) GetNetworkID(ctx context.Context, in *GetNetworkIDRequest, opts ...dcerpc.CallOption) (*GetNetworkIDResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNetworkIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) SetNetworkPriorityOrder(ctx context.Context, in *SetNetworkPriorityOrderRequest, opts ...dcerpc.CallOption) (*SetNetworkPriorityOrderResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetNetworkPriorityOrderResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) NodeNetworkControl(ctx context.Context, in *NodeNetworkControlRequest, opts ...dcerpc.CallOption) (*NodeNetworkControlResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &NodeNetworkControlResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) NetworkControl(ctx context.Context, in *NetworkControlRequest, opts ...dcerpc.CallOption) (*NetworkControlResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &NetworkControlResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) AddNotifyNetwork(ctx context.Context, in *AddNotifyNetworkRequest, opts ...dcerpc.CallOption) (*AddNotifyNetworkResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddNotifyNetworkResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) ReAddNotifyNetwork(ctx context.Context, in *ReAddNotifyNetworkRequest, opts ...dcerpc.CallOption) (*ReAddNotifyNetworkResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ReAddNotifyNetworkResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) OpenNetInterface(ctx context.Context, in *OpenNetInterfaceRequest, opts ...dcerpc.CallOption) (*OpenNetInterfaceResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenNetInterfaceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != nil {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) CloseNetInterface(ctx context.Context, in *CloseNetInterfaceRequest, opts ...dcerpc.CallOption) (*CloseNetInterfaceResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CloseNetInterfaceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) GetNetInterfaceState(ctx context.Context, in *GetNetInterfaceStateRequest, opts ...dcerpc.CallOption) (*GetNetInterfaceStateResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNetInterfaceStateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) GetNetInterface(ctx context.Context, in *GetNetInterfaceRequest, opts ...dcerpc.CallOption) (*GetNetInterfaceResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNetInterfaceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) GetNetInterfaceID(ctx context.Context, in *GetNetInterfaceIDRequest, opts ...dcerpc.CallOption) (*GetNetInterfaceIDResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNetInterfaceIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) NodeNetInterfaceControl(ctx context.Context, in *NodeNetInterfaceControlRequest, opts ...dcerpc.CallOption) (*NodeNetInterfaceControlResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &NodeNetInterfaceControlResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) NetInterfaceControl(ctx context.Context, in *NetInterfaceControlRequest, opts ...dcerpc.CallOption) (*NetInterfaceControlResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &NetInterfaceControlResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) AddNotifyNetInterface(ctx context.Context, in *AddNotifyNetInterfaceRequest, opts ...dcerpc.CallOption) (*AddNotifyNetInterfaceResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddNotifyNetInterfaceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) ReAddNotifyNetInterface(ctx context.Context, in *ReAddNotifyNetInterfaceRequest, opts ...dcerpc.CallOption) (*ReAddNotifyNetInterfaceResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ReAddNotifyNetInterfaceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) CreateNodeEnum(ctx context.Context, in *CreateNodeEnumRequest, opts ...dcerpc.CallOption) (*CreateNodeEnumResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateNodeEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) GetClusterVersion2(ctx context.Context, in *GetClusterVersion2Request, opts ...dcerpc.CallOption) (*GetClusterVersion2Response, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetClusterVersion2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) CreateRestrictionTypeEnum(ctx context.Context, in *CreateRestrictionTypeEnumRequest, opts ...dcerpc.CallOption) (*CreateRestrictionTypeEnumResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateRestrictionTypeEnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) BackupClusterDatabase(ctx context.Context, in *BackupClusterDatabaseRequest, opts ...dcerpc.CallOption) (*BackupClusterDatabaseResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BackupClusterDatabaseResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) NodeClusterControl(ctx context.Context, in *NodeClusterControlRequest, opts ...dcerpc.CallOption) (*NodeClusterControlResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &NodeClusterControlResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) ClusterControl(ctx context.Context, in *ClusterControlRequest, opts ...dcerpc.CallOption) (*ClusterControlResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ClusterControlResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) UnblockGetNotifyCall(ctx context.Context, in *UnblockGetNotifyCallRequest, opts ...dcerpc.CallOption) (*UnblockGetNotifyCallResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &UnblockGetNotifyCallResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) SetServiceAccountPassword(ctx context.Context, in *SetServiceAccountPasswordRequest, opts ...dcerpc.CallOption) (*SetServiceAccountPasswordResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetServiceAccountPasswordResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusapi2Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}
func NewClusapi2Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (Clusapi2Client, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(Clusapi2SyntaxV2_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultClusapi2Client{cc: cc}, nil
}

// xxx_OpenClusterOperation structure represents the ApiOpenCluster operation
type xxx_OpenClusterOperation struct {
	Status uint32   `idl:"name:Status" json:"status"`
	Return *Cluster `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenClusterOperation) OpNum() int { return 0 }

func (o *xxx_OpenClusterOperation) OpName() string { return "/clusapi2/v2/ApiOpenCluster" }

func (o *xxx_OpenClusterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenClusterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_OpenClusterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_OpenClusterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenClusterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HCLUSTER_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return != nil {
			if err := o.Return.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Cluster{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_OpenClusterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HCLUSTER_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return == nil {
			o.Return = &Cluster{}
		}
		if err := o.Return.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// OpenClusterRequest structure represents the ApiOpenCluster operation request
type OpenClusterRequest struct {
}

func (o *OpenClusterRequest) xxx_ToOp(ctx context.Context) *xxx_OpenClusterOperation {
	if o == nil {
		return &xxx_OpenClusterOperation{}
	}
	return &xxx_OpenClusterOperation{}
}

func (o *OpenClusterRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenClusterOperation) {
	if o == nil {
		return
	}
}
func (o *OpenClusterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OpenClusterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenClusterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenClusterResponse structure represents the ApiOpenCluster operation response
type OpenClusterResponse struct {
	Status uint32 `idl:"name:Status" json:"status"`
	// Return: The ApiOpenCluster return value.
	Return *Cluster `idl:"name:Return" json:"return"`
}

func (o *OpenClusterResponse) xxx_ToOp(ctx context.Context) *xxx_OpenClusterOperation {
	if o == nil {
		return &xxx_OpenClusterOperation{}
	}
	return &xxx_OpenClusterOperation{
		Status: o.Status,
		Return: o.Return,
	}
}

func (o *OpenClusterResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenClusterOperation) {
	if o == nil {
		return
	}
	o.Status = op.Status
	o.Return = op.Return
}
func (o *OpenClusterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OpenClusterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenClusterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseClusterOperation structure represents the ApiCloseCluster operation
type xxx_CloseClusterOperation struct {
	Cluster *Cluster `idl:"name:Cluster" json:"cluster"`
	Return  uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseClusterOperation) OpNum() int { return 1 }

func (o *xxx_CloseClusterOperation) OpName() string { return "/clusapi2/v2/ApiCloseCluster" }

func (o *xxx_CloseClusterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseClusterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Cluster {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HCLUSTER_RPC, names=ndr_context_handle}(struct))
	{
		if o.Cluster != nil {
			if err := o.Cluster.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Cluster{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseClusterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Cluster {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HCLUSTER_RPC, names=ndr_context_handle}(struct))
	{
		if o.Cluster == nil {
			o.Cluster = &Cluster{}
		}
		if err := o.Cluster.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseClusterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseClusterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Cluster {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HCLUSTER_RPC, names=ndr_context_handle}(struct))
	{
		if o.Cluster != nil {
			if err := o.Cluster.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Cluster{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseClusterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Cluster {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HCLUSTER_RPC, names=ndr_context_handle}(struct))
	{
		if o.Cluster == nil {
			o.Cluster = &Cluster{}
		}
		if err := o.Cluster.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloseClusterRequest structure represents the ApiCloseCluster operation request
type CloseClusterRequest struct {
	Cluster *Cluster `idl:"name:Cluster" json:"cluster"`
}

func (o *CloseClusterRequest) xxx_ToOp(ctx context.Context) *xxx_CloseClusterOperation {
	if o == nil {
		return &xxx_CloseClusterOperation{}
	}
	return &xxx_CloseClusterOperation{
		Cluster: o.Cluster,
	}
}

func (o *CloseClusterRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseClusterOperation) {
	if o == nil {
		return
	}
	o.Cluster = op.Cluster
}
func (o *CloseClusterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CloseClusterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseClusterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseClusterResponse structure represents the ApiCloseCluster operation response
type CloseClusterResponse struct {
	Cluster *Cluster `idl:"name:Cluster" json:"cluster"`
	// Return: The ApiCloseCluster return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CloseClusterResponse) xxx_ToOp(ctx context.Context) *xxx_CloseClusterOperation {
	if o == nil {
		return &xxx_CloseClusterOperation{}
	}
	return &xxx_CloseClusterOperation{
		Cluster: o.Cluster,
		Return:  o.Return,
	}
}

func (o *CloseClusterResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseClusterOperation) {
	if o == nil {
		return
	}
	o.Cluster = op.Cluster
	o.Return = op.Return
}
func (o *CloseClusterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CloseClusterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseClusterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetClusterNameOperation structure represents the ApiSetClusterName operation
type xxx_SetClusterNameOperation struct {
	NewClusterName string `idl:"name:NewClusterName;string" json:"new_cluster_name"`
	Return         uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SetClusterNameOperation) OpNum() int { return 2 }

func (o *xxx_SetClusterNameOperation) OpName() string { return "/clusapi2/v2/ApiSetClusterName" }

func (o *xxx_SetClusterNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClusterNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// NewClusterName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.NewClusterName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClusterNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// NewClusterName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.NewClusterName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClusterNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClusterNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClusterNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetClusterNameRequest structure represents the ApiSetClusterName operation request
type SetClusterNameRequest struct {
	NewClusterName string `idl:"name:NewClusterName;string" json:"new_cluster_name"`
}

func (o *SetClusterNameRequest) xxx_ToOp(ctx context.Context) *xxx_SetClusterNameOperation {
	if o == nil {
		return &xxx_SetClusterNameOperation{}
	}
	return &xxx_SetClusterNameOperation{
		NewClusterName: o.NewClusterName,
	}
}

func (o *SetClusterNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetClusterNameOperation) {
	if o == nil {
		return
	}
	o.NewClusterName = op.NewClusterName
}
func (o *SetClusterNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetClusterNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClusterNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetClusterNameResponse structure represents the ApiSetClusterName operation response
type SetClusterNameResponse struct {
	// Return: The ApiSetClusterName return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetClusterNameResponse) xxx_ToOp(ctx context.Context) *xxx_SetClusterNameOperation {
	if o == nil {
		return &xxx_SetClusterNameOperation{}
	}
	return &xxx_SetClusterNameOperation{
		Return: o.Return,
	}
}

func (o *SetClusterNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetClusterNameOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetClusterNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetClusterNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClusterNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetClusterNameOperation structure represents the ApiGetClusterName operation
type xxx_GetClusterNameOperation struct {
	ClusterName string `idl:"name:ClusterName;string" json:"cluster_name"`
	NodeName    string `idl:"name:NodeName;string" json:"node_name"`
	Return      uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClusterNameOperation) OpNum() int { return 3 }

func (o *xxx_GetClusterNameOperation) OpName() string { return "/clusapi2/v2/ApiGetClusterName" }

func (o *xxx_GetClusterNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClusterNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_GetClusterNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_GetClusterNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClusterNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ClusterName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ClusterName != "" {
			_ptr_ClusterName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ClusterName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ClusterName, _ptr_ClusterName); err != nil {
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
	// NodeName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.NodeName != "" {
			_ptr_NodeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.NodeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.NodeName, _ptr_NodeName); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClusterNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ClusterName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ClusterName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ClusterName); err != nil {
				return err
			}
			return nil
		})
		_s_ClusterName := func(ptr interface{}) { o.ClusterName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ClusterName, _s_ClusterName, _ptr_ClusterName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// NodeName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_NodeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.NodeName); err != nil {
				return err
			}
			return nil
		})
		_s_NodeName := func(ptr interface{}) { o.NodeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.NodeName, _s_NodeName, _ptr_NodeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetClusterNameRequest structure represents the ApiGetClusterName operation request
type GetClusterNameRequest struct {
}

func (o *GetClusterNameRequest) xxx_ToOp(ctx context.Context) *xxx_GetClusterNameOperation {
	if o == nil {
		return &xxx_GetClusterNameOperation{}
	}
	return &xxx_GetClusterNameOperation{}
}

func (o *GetClusterNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClusterNameOperation) {
	if o == nil {
		return
	}
}
func (o *GetClusterNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetClusterNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClusterNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClusterNameResponse structure represents the ApiGetClusterName operation response
type GetClusterNameResponse struct {
	ClusterName string `idl:"name:ClusterName;string" json:"cluster_name"`
	NodeName    string `idl:"name:NodeName;string" json:"node_name"`
	// Return: The ApiGetClusterName return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetClusterNameResponse) xxx_ToOp(ctx context.Context) *xxx_GetClusterNameOperation {
	if o == nil {
		return &xxx_GetClusterNameOperation{}
	}
	return &xxx_GetClusterNameOperation{
		ClusterName: o.ClusterName,
		NodeName:    o.NodeName,
		Return:      o.Return,
	}
}

func (o *GetClusterNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClusterNameOperation) {
	if o == nil {
		return
	}
	o.ClusterName = op.ClusterName
	o.NodeName = op.NodeName
	o.Return = op.Return
}
func (o *GetClusterNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetClusterNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClusterNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetClusterVersionOperation structure represents the ApiGetClusterVersion operation
type xxx_GetClusterVersionOperation struct {
	MajorVersion uint16 `idl:"name:lpwMajorVersion" json:"major_version"`
	MinorVersion uint16 `idl:"name:lpwMinorVersion" json:"minor_version"`
	BuildNumber  uint16 `idl:"name:lpwBuildNumber" json:"build_number"`
	VendorID     string `idl:"name:lpszVendorId;string" json:"vendor_id"`
	CSDVersion   string `idl:"name:lpszCSDVersion;string" json:"csd_version"`
	Return       uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClusterVersionOperation) OpNum() int { return 4 }

func (o *xxx_GetClusterVersionOperation) OpName() string { return "/clusapi2/v2/ApiGetClusterVersion" }

func (o *xxx_GetClusterVersionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClusterVersionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_GetClusterVersionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_GetClusterVersionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClusterVersionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpwMajorVersion {out} (1:{pointer=ref}*(1))(2:{alias=WORD}(uint16))
	{
		if err := w.WriteData(o.MajorVersion); err != nil {
			return err
		}
	}
	// lpwMinorVersion {out} (1:{pointer=ref}*(1))(2:{alias=WORD}(uint16))
	{
		if err := w.WriteData(o.MinorVersion); err != nil {
			return err
		}
	}
	// lpwBuildNumber {out} (1:{pointer=ref}*(1))(2:{alias=WORD}(uint16))
	{
		if err := w.WriteData(o.BuildNumber); err != nil {
			return err
		}
	}
	// lpszVendorId {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.VendorID != "" {
			_ptr_lpszVendorId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.VendorID); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.VendorID, _ptr_lpszVendorId); err != nil {
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
	// lpszCSDVersion {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.CSDVersion != "" {
			_ptr_lpszCSDVersion := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.CSDVersion); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.CSDVersion, _ptr_lpszCSDVersion); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClusterVersionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpwMajorVersion {out} (1:{pointer=ref}*(1))(2:{alias=WORD}(uint16))
	{
		if err := w.ReadData(&o.MajorVersion); err != nil {
			return err
		}
	}
	// lpwMinorVersion {out} (1:{pointer=ref}*(1))(2:{alias=WORD}(uint16))
	{
		if err := w.ReadData(&o.MinorVersion); err != nil {
			return err
		}
	}
	// lpwBuildNumber {out} (1:{pointer=ref}*(1))(2:{alias=WORD}(uint16))
	{
		if err := w.ReadData(&o.BuildNumber); err != nil {
			return err
		}
	}
	// lpszVendorId {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpszVendorId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.VendorID); err != nil {
				return err
			}
			return nil
		})
		_s_lpszVendorId := func(ptr interface{}) { o.VendorID = *ptr.(*string) }
		if err := w.ReadPointer(&o.VendorID, _s_lpszVendorId, _ptr_lpszVendorId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpszCSDVersion {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpszCSDVersion := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.CSDVersion); err != nil {
				return err
			}
			return nil
		})
		_s_lpszCSDVersion := func(ptr interface{}) { o.CSDVersion = *ptr.(*string) }
		if err := w.ReadPointer(&o.CSDVersion, _s_lpszCSDVersion, _ptr_lpszCSDVersion); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetClusterVersionRequest structure represents the ApiGetClusterVersion operation request
type GetClusterVersionRequest struct {
}

func (o *GetClusterVersionRequest) xxx_ToOp(ctx context.Context) *xxx_GetClusterVersionOperation {
	if o == nil {
		return &xxx_GetClusterVersionOperation{}
	}
	return &xxx_GetClusterVersionOperation{}
}

func (o *GetClusterVersionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClusterVersionOperation) {
	if o == nil {
		return
	}
}
func (o *GetClusterVersionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetClusterVersionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClusterVersionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClusterVersionResponse structure represents the ApiGetClusterVersion operation response
type GetClusterVersionResponse struct {
	MajorVersion uint16 `idl:"name:lpwMajorVersion" json:"major_version"`
	MinorVersion uint16 `idl:"name:lpwMinorVersion" json:"minor_version"`
	BuildNumber  uint16 `idl:"name:lpwBuildNumber" json:"build_number"`
	VendorID     string `idl:"name:lpszVendorId;string" json:"vendor_id"`
	CSDVersion   string `idl:"name:lpszCSDVersion;string" json:"csd_version"`
	// Return: The ApiGetClusterVersion return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetClusterVersionResponse) xxx_ToOp(ctx context.Context) *xxx_GetClusterVersionOperation {
	if o == nil {
		return &xxx_GetClusterVersionOperation{}
	}
	return &xxx_GetClusterVersionOperation{
		MajorVersion: o.MajorVersion,
		MinorVersion: o.MinorVersion,
		BuildNumber:  o.BuildNumber,
		VendorID:     o.VendorID,
		CSDVersion:   o.CSDVersion,
		Return:       o.Return,
	}
}

func (o *GetClusterVersionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClusterVersionOperation) {
	if o == nil {
		return
	}
	o.MajorVersion = op.MajorVersion
	o.MinorVersion = op.MinorVersion
	o.BuildNumber = op.BuildNumber
	o.VendorID = op.VendorID
	o.CSDVersion = op.CSDVersion
	o.Return = op.Return
}
func (o *GetClusterVersionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetClusterVersionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClusterVersionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetQuorumResourceOperation structure represents the ApiGetQuorumResource operation
type xxx_GetQuorumResourceOperation struct {
	ResourceName     string `idl:"name:lpszResourceName;string" json:"resource_name"`
	DeviceName       string `idl:"name:lpszDeviceName;string" json:"device_name"`
	MaxQuorumLogSize uint32 `idl:"name:pdwMaxQuorumLogSize" json:"max_quorum_log_size"`
	Return           uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetQuorumResourceOperation) OpNum() int { return 5 }

func (o *xxx_GetQuorumResourceOperation) OpName() string { return "/clusapi2/v2/ApiGetQuorumResource" }

func (o *xxx_GetQuorumResourceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQuorumResourceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_GetQuorumResourceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_GetQuorumResourceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQuorumResourceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpszResourceName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ResourceName != "" {
			_ptr_lpszResourceName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ResourceName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ResourceName, _ptr_lpszResourceName); err != nil {
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
	// lpszDeviceName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.DeviceName != "" {
			_ptr_lpszDeviceName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.DeviceName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DeviceName, _ptr_lpszDeviceName); err != nil {
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
	// pdwMaxQuorumLogSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MaxQuorumLogSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQuorumResourceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpszResourceName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpszResourceName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ResourceName); err != nil {
				return err
			}
			return nil
		})
		_s_lpszResourceName := func(ptr interface{}) { o.ResourceName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ResourceName, _s_lpszResourceName, _ptr_lpszResourceName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpszDeviceName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpszDeviceName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.DeviceName); err != nil {
				return err
			}
			return nil
		})
		_s_lpszDeviceName := func(ptr interface{}) { o.DeviceName = *ptr.(*string) }
		if err := w.ReadPointer(&o.DeviceName, _s_lpszDeviceName, _ptr_lpszDeviceName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwMaxQuorumLogSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MaxQuorumLogSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetQuorumResourceRequest structure represents the ApiGetQuorumResource operation request
type GetQuorumResourceRequest struct {
}

func (o *GetQuorumResourceRequest) xxx_ToOp(ctx context.Context) *xxx_GetQuorumResourceOperation {
	if o == nil {
		return &xxx_GetQuorumResourceOperation{}
	}
	return &xxx_GetQuorumResourceOperation{}
}

func (o *GetQuorumResourceRequest) xxx_FromOp(ctx context.Context, op *xxx_GetQuorumResourceOperation) {
	if o == nil {
		return
	}
}
func (o *GetQuorumResourceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetQuorumResourceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQuorumResourceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetQuorumResourceResponse structure represents the ApiGetQuorumResource operation response
type GetQuorumResourceResponse struct {
	ResourceName     string `idl:"name:lpszResourceName;string" json:"resource_name"`
	DeviceName       string `idl:"name:lpszDeviceName;string" json:"device_name"`
	MaxQuorumLogSize uint32 `idl:"name:pdwMaxQuorumLogSize" json:"max_quorum_log_size"`
	// Return: The ApiGetQuorumResource return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetQuorumResourceResponse) xxx_ToOp(ctx context.Context) *xxx_GetQuorumResourceOperation {
	if o == nil {
		return &xxx_GetQuorumResourceOperation{}
	}
	return &xxx_GetQuorumResourceOperation{
		ResourceName:     o.ResourceName,
		DeviceName:       o.DeviceName,
		MaxQuorumLogSize: o.MaxQuorumLogSize,
		Return:           o.Return,
	}
}

func (o *GetQuorumResourceResponse) xxx_FromOp(ctx context.Context, op *xxx_GetQuorumResourceOperation) {
	if o == nil {
		return
	}
	o.ResourceName = op.ResourceName
	o.DeviceName = op.DeviceName
	o.MaxQuorumLogSize = op.MaxQuorumLogSize
	o.Return = op.Return
}
func (o *GetQuorumResourceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetQuorumResourceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQuorumResourceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetQuorumResourceOperation structure represents the ApiSetQuorumResource operation
type xxx_SetQuorumResourceOperation struct {
	Resource         *Resource `idl:"name:hResource" json:"resource"`
	DeviceName       string    `idl:"name:lpszDeviceName;string" json:"device_name"`
	MaxQuorumLogSize uint32    `idl:"name:dwMaxQuorumLogSize" json:"max_quorum_log_size"`
	Return           uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_SetQuorumResourceOperation) OpNum() int { return 6 }

func (o *xxx_SetQuorumResourceOperation) OpName() string { return "/clusapi2/v2/ApiSetQuorumResource" }

func (o *xxx_SetQuorumResourceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetQuorumResourceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpszDeviceName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DeviceName); err != nil {
			return err
		}
	}
	// dwMaxQuorumLogSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MaxQuorumLogSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetQuorumResourceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpszDeviceName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DeviceName); err != nil {
			return err
		}
	}
	// dwMaxQuorumLogSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MaxQuorumLogSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetQuorumResourceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetQuorumResourceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetQuorumResourceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetQuorumResourceRequest structure represents the ApiSetQuorumResource operation request
type SetQuorumResourceRequest struct {
	Resource         *Resource `idl:"name:hResource" json:"resource"`
	DeviceName       string    `idl:"name:lpszDeviceName;string" json:"device_name"`
	MaxQuorumLogSize uint32    `idl:"name:dwMaxQuorumLogSize" json:"max_quorum_log_size"`
}

func (o *SetQuorumResourceRequest) xxx_ToOp(ctx context.Context) *xxx_SetQuorumResourceOperation {
	if o == nil {
		return &xxx_SetQuorumResourceOperation{}
	}
	return &xxx_SetQuorumResourceOperation{
		Resource:         o.Resource,
		DeviceName:       o.DeviceName,
		MaxQuorumLogSize: o.MaxQuorumLogSize,
	}
}

func (o *SetQuorumResourceRequest) xxx_FromOp(ctx context.Context, op *xxx_SetQuorumResourceOperation) {
	if o == nil {
		return
	}
	o.Resource = op.Resource
	o.DeviceName = op.DeviceName
	o.MaxQuorumLogSize = op.MaxQuorumLogSize
}
func (o *SetQuorumResourceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetQuorumResourceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetQuorumResourceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetQuorumResourceResponse structure represents the ApiSetQuorumResource operation response
type SetQuorumResourceResponse struct {
	// Return: The ApiSetQuorumResource return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetQuorumResourceResponse) xxx_ToOp(ctx context.Context) *xxx_SetQuorumResourceOperation {
	if o == nil {
		return &xxx_SetQuorumResourceOperation{}
	}
	return &xxx_SetQuorumResourceOperation{
		Return: o.Return,
	}
}

func (o *SetQuorumResourceResponse) xxx_FromOp(ctx context.Context, op *xxx_SetQuorumResourceOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetQuorumResourceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetQuorumResourceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetQuorumResourceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateEnumOperation structure represents the ApiCreateEnum operation
type xxx_CreateEnumOperation struct {
	Type       uint32    `idl:"name:dwType" json:"type"`
	ReturnEnum *EnumList `idl:"name:ReturnEnum" json:"return_enum"`
	Return     uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateEnumOperation) OpNum() int { return 7 }

func (o *xxx_CreateEnumOperation) OpName() string { return "/clusapi2/v2/ApiCreateEnum" }

func (o *xxx_CreateEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ReturnEnum {out} (1:{pointer=ref}*(2))(2:{alias=PENUM_LIST}*(1))(3:{alias=ENUM_LIST}(struct))
	{
		if o.ReturnEnum != nil {
			_ptr_ReturnEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnEnum != nil {
					if err := o.ReturnEnum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EnumList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnEnum, _ptr_ReturnEnum); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ReturnEnum {out} (1:{pointer=ref}*(2))(2:{alias=PENUM_LIST,pointer=ref}*(1))(3:{alias=ENUM_LIST}(struct))
	{
		_ptr_ReturnEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnEnum == nil {
				o.ReturnEnum = &EnumList{}
			}
			if err := o.ReturnEnum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReturnEnum := func(ptr interface{}) { o.ReturnEnum = *ptr.(**EnumList) }
		if err := w.ReadPointer(&o.ReturnEnum, _s_ReturnEnum, _ptr_ReturnEnum); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateEnumRequest structure represents the ApiCreateEnum operation request
type CreateEnumRequest struct {
	Type uint32 `idl:"name:dwType" json:"type"`
}

func (o *CreateEnumRequest) xxx_ToOp(ctx context.Context) *xxx_CreateEnumOperation {
	if o == nil {
		return &xxx_CreateEnumOperation{}
	}
	return &xxx_CreateEnumOperation{
		Type: o.Type,
	}
}

func (o *CreateEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateEnumOperation) {
	if o == nil {
		return
	}
	o.Type = op.Type
}
func (o *CreateEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateEnumResponse structure represents the ApiCreateEnum operation response
type CreateEnumResponse struct {
	ReturnEnum *EnumList `idl:"name:ReturnEnum" json:"return_enum"`
	// Return: The ApiCreateEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CreateEnumResponse) xxx_ToOp(ctx context.Context) *xxx_CreateEnumOperation {
	if o == nil {
		return &xxx_CreateEnumOperation{}
	}
	return &xxx_CreateEnumOperation{
		ReturnEnum: o.ReturnEnum,
		Return:     o.Return,
	}
}

func (o *CreateEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateEnumOperation) {
	if o == nil {
		return
	}
	o.ReturnEnum = op.ReturnEnum
	o.Return = op.Return
}
func (o *CreateEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenResourceOperation structure represents the ApiOpenResource operation
type xxx_OpenResourceOperation struct {
	ResourceName string    `idl:"name:lpszResourceName;string" json:"resource_name"`
	Status       uint32    `idl:"name:Status" json:"status"`
	Return       *Resource `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenResourceOperation) OpNum() int { return 8 }

func (o *xxx_OpenResourceOperation) OpName() string { return "/clusapi2/v2/ApiOpenResource" }

func (o *xxx_OpenResourceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenResourceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpszResourceName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ResourceName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenResourceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpszResourceName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ResourceName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenResourceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenResourceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return != nil {
			if err := o.Return.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_OpenResourceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return == nil {
			o.Return = &Resource{}
		}
		if err := o.Return.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// OpenResourceRequest structure represents the ApiOpenResource operation request
type OpenResourceRequest struct {
	ResourceName string `idl:"name:lpszResourceName;string" json:"resource_name"`
}

func (o *OpenResourceRequest) xxx_ToOp(ctx context.Context) *xxx_OpenResourceOperation {
	if o == nil {
		return &xxx_OpenResourceOperation{}
	}
	return &xxx_OpenResourceOperation{
		ResourceName: o.ResourceName,
	}
}

func (o *OpenResourceRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenResourceOperation) {
	if o == nil {
		return
	}
	o.ResourceName = op.ResourceName
}
func (o *OpenResourceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OpenResourceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenResourceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenResourceResponse structure represents the ApiOpenResource operation response
type OpenResourceResponse struct {
	Status uint32 `idl:"name:Status" json:"status"`
	// Return: The ApiOpenResource return value.
	Return *Resource `idl:"name:Return" json:"return"`
}

func (o *OpenResourceResponse) xxx_ToOp(ctx context.Context) *xxx_OpenResourceOperation {
	if o == nil {
		return &xxx_OpenResourceOperation{}
	}
	return &xxx_OpenResourceOperation{
		Status: o.Status,
		Return: o.Return,
	}
}

func (o *OpenResourceResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenResourceOperation) {
	if o == nil {
		return
	}
	o.Status = op.Status
	o.Return = op.Return
}
func (o *OpenResourceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OpenResourceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenResourceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateResourceOperation structure represents the ApiCreateResource operation
type xxx_CreateResourceOperation struct {
	Group        *Group    `idl:"name:hGroup" json:"group"`
	ResourceName string    `idl:"name:lpszResourceName;string" json:"resource_name"`
	ResourceType string    `idl:"name:lpszResourceType;string" json:"resource_type"`
	Flags        uint32    `idl:"name:dwFlags" json:"flags"`
	Status       uint32    `idl:"name:Status" json:"status"`
	Return       *Resource `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateResourceOperation) OpNum() int { return 9 }

func (o *xxx_CreateResourceOperation) OpName() string { return "/clusapi2/v2/ApiCreateResource" }

func (o *xxx_CreateResourceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateResourceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group != nil {
			if err := o.Group.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Group{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpszResourceName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ResourceName); err != nil {
			return err
		}
	}
	// lpszResourceType {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ResourceType); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateResourceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group == nil {
			o.Group = &Group{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpszResourceName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ResourceName); err != nil {
			return err
		}
	}
	// lpszResourceType {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ResourceType); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateResourceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateResourceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return != nil {
			if err := o.Return.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CreateResourceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return == nil {
			o.Return = &Resource{}
		}
		if err := o.Return.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// CreateResourceRequest structure represents the ApiCreateResource operation request
type CreateResourceRequest struct {
	Group        *Group `idl:"name:hGroup" json:"group"`
	ResourceName string `idl:"name:lpszResourceName;string" json:"resource_name"`
	ResourceType string `idl:"name:lpszResourceType;string" json:"resource_type"`
	Flags        uint32 `idl:"name:dwFlags" json:"flags"`
}

func (o *CreateResourceRequest) xxx_ToOp(ctx context.Context) *xxx_CreateResourceOperation {
	if o == nil {
		return &xxx_CreateResourceOperation{}
	}
	return &xxx_CreateResourceOperation{
		Group:        o.Group,
		ResourceName: o.ResourceName,
		ResourceType: o.ResourceType,
		Flags:        o.Flags,
	}
}

func (o *CreateResourceRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateResourceOperation) {
	if o == nil {
		return
	}
	o.Group = op.Group
	o.ResourceName = op.ResourceName
	o.ResourceType = op.ResourceType
	o.Flags = op.Flags
}
func (o *CreateResourceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateResourceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateResourceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateResourceResponse structure represents the ApiCreateResource operation response
type CreateResourceResponse struct {
	Status uint32 `idl:"name:Status" json:"status"`
	// Return: The ApiCreateResource return value.
	Return *Resource `idl:"name:Return" json:"return"`
}

func (o *CreateResourceResponse) xxx_ToOp(ctx context.Context) *xxx_CreateResourceOperation {
	if o == nil {
		return &xxx_CreateResourceOperation{}
	}
	return &xxx_CreateResourceOperation{
		Status: o.Status,
		Return: o.Return,
	}
}

func (o *CreateResourceResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateResourceOperation) {
	if o == nil {
		return
	}
	o.Status = op.Status
	o.Return = op.Return
}
func (o *CreateResourceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateResourceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateResourceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteResourceOperation structure represents the ApiDeleteResource operation
type xxx_DeleteResourceOperation struct {
	Resource *Resource `idl:"name:hResource" json:"resource"`
	Return   uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteResourceOperation) OpNum() int { return 10 }

func (o *xxx_DeleteResourceOperation) OpName() string { return "/clusapi2/v2/ApiDeleteResource" }

func (o *xxx_DeleteResourceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteResourceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DeleteResourceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteResourceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteResourceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteResourceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteResourceRequest structure represents the ApiDeleteResource operation request
type DeleteResourceRequest struct {
	Resource *Resource `idl:"name:hResource" json:"resource"`
}

func (o *DeleteResourceRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteResourceOperation {
	if o == nil {
		return &xxx_DeleteResourceOperation{}
	}
	return &xxx_DeleteResourceOperation{
		Resource: o.Resource,
	}
}

func (o *DeleteResourceRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteResourceOperation) {
	if o == nil {
		return
	}
	o.Resource = op.Resource
}
func (o *DeleteResourceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteResourceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteResourceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteResourceResponse structure represents the ApiDeleteResource operation response
type DeleteResourceResponse struct {
	// Return: The ApiDeleteResource return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *DeleteResourceResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteResourceOperation {
	if o == nil {
		return &xxx_DeleteResourceOperation{}
	}
	return &xxx_DeleteResourceOperation{
		Return: o.Return,
	}
}

func (o *DeleteResourceResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteResourceOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DeleteResourceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteResourceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteResourceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseResourceOperation structure represents the ApiCloseResource operation
type xxx_CloseResourceOperation struct {
	Resource *Resource `idl:"name:Resource" json:"resource"`
	Return   uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseResourceOperation) OpNum() int { return 11 }

func (o *xxx_CloseResourceOperation) OpName() string { return "/clusapi2/v2/ApiCloseResource" }

func (o *xxx_CloseResourceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseResourceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Resource {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseResourceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Resource {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseResourceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseResourceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Resource {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseResourceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Resource {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloseResourceRequest structure represents the ApiCloseResource operation request
type CloseResourceRequest struct {
	Resource *Resource `idl:"name:Resource" json:"resource"`
}

func (o *CloseResourceRequest) xxx_ToOp(ctx context.Context) *xxx_CloseResourceOperation {
	if o == nil {
		return &xxx_CloseResourceOperation{}
	}
	return &xxx_CloseResourceOperation{
		Resource: o.Resource,
	}
}

func (o *CloseResourceRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseResourceOperation) {
	if o == nil {
		return
	}
	o.Resource = op.Resource
}
func (o *CloseResourceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CloseResourceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseResourceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseResourceResponse structure represents the ApiCloseResource operation response
type CloseResourceResponse struct {
	Resource *Resource `idl:"name:Resource" json:"resource"`
	// Return: The ApiCloseResource return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CloseResourceResponse) xxx_ToOp(ctx context.Context) *xxx_CloseResourceOperation {
	if o == nil {
		return &xxx_CloseResourceOperation{}
	}
	return &xxx_CloseResourceOperation{
		Resource: o.Resource,
		Return:   o.Return,
	}
}

func (o *CloseResourceResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseResourceOperation) {
	if o == nil {
		return
	}
	o.Resource = op.Resource
	o.Return = op.Return
}
func (o *CloseResourceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CloseResourceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseResourceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetResourceStateOperation structure represents the ApiGetResourceState operation
type xxx_GetResourceStateOperation struct {
	Resource  *Resource `idl:"name:hResource" json:"resource"`
	State     uint32    `idl:"name:State" json:"state"`
	NodeName  string    `idl:"name:NodeName;string" json:"node_name"`
	GroupName string    `idl:"name:GroupName;string" json:"group_name"`
	Return    uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_GetResourceStateOperation) OpNum() int { return 12 }

func (o *xxx_GetResourceStateOperation) OpName() string { return "/clusapi2/v2/ApiGetResourceState" }

func (o *xxx_GetResourceStateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResourceStateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetResourceStateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResourceStateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResourceStateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// State {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.State); err != nil {
			return err
		}
	}
	// NodeName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.NodeName != "" {
			_ptr_NodeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.NodeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.NodeName, _ptr_NodeName); err != nil {
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
	// GroupName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.GroupName != "" {
			_ptr_GroupName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.GroupName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.GroupName, _ptr_GroupName); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResourceStateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// State {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.State); err != nil {
			return err
		}
	}
	// NodeName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_NodeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.NodeName); err != nil {
				return err
			}
			return nil
		})
		_s_NodeName := func(ptr interface{}) { o.NodeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.NodeName, _s_NodeName, _ptr_NodeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// GroupName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_GroupName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.GroupName); err != nil {
				return err
			}
			return nil
		})
		_s_GroupName := func(ptr interface{}) { o.GroupName = *ptr.(*string) }
		if err := w.ReadPointer(&o.GroupName, _s_GroupName, _ptr_GroupName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetResourceStateRequest structure represents the ApiGetResourceState operation request
type GetResourceStateRequest struct {
	Resource *Resource `idl:"name:hResource" json:"resource"`
}

func (o *GetResourceStateRequest) xxx_ToOp(ctx context.Context) *xxx_GetResourceStateOperation {
	if o == nil {
		return &xxx_GetResourceStateOperation{}
	}
	return &xxx_GetResourceStateOperation{
		Resource: o.Resource,
	}
}

func (o *GetResourceStateRequest) xxx_FromOp(ctx context.Context, op *xxx_GetResourceStateOperation) {
	if o == nil {
		return
	}
	o.Resource = op.Resource
}
func (o *GetResourceStateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetResourceStateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResourceStateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetResourceStateResponse structure represents the ApiGetResourceState operation response
type GetResourceStateResponse struct {
	State     uint32 `idl:"name:State" json:"state"`
	NodeName  string `idl:"name:NodeName;string" json:"node_name"`
	GroupName string `idl:"name:GroupName;string" json:"group_name"`
	// Return: The ApiGetResourceState return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetResourceStateResponse) xxx_ToOp(ctx context.Context) *xxx_GetResourceStateOperation {
	if o == nil {
		return &xxx_GetResourceStateOperation{}
	}
	return &xxx_GetResourceStateOperation{
		State:     o.State,
		NodeName:  o.NodeName,
		GroupName: o.GroupName,
		Return:    o.Return,
	}
}

func (o *GetResourceStateResponse) xxx_FromOp(ctx context.Context, op *xxx_GetResourceStateOperation) {
	if o == nil {
		return
	}
	o.State = op.State
	o.NodeName = op.NodeName
	o.GroupName = op.GroupName
	o.Return = op.Return
}
func (o *GetResourceStateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetResourceStateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResourceStateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetResourceNameOperation structure represents the ApiSetResourceName operation
type xxx_SetResourceNameOperation struct {
	Resource     *Resource `idl:"name:hResource" json:"resource"`
	ResourceName string    `idl:"name:lpszResourceName;string" json:"resource_name"`
	Return       uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_SetResourceNameOperation) OpNum() int { return 13 }

func (o *xxx_SetResourceNameOperation) OpName() string { return "/clusapi2/v2/ApiSetResourceName" }

func (o *xxx_SetResourceNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResourceNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpszResourceName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ResourceName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResourceNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpszResourceName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ResourceName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResourceNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResourceNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetResourceNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetResourceNameRequest structure represents the ApiSetResourceName operation request
type SetResourceNameRequest struct {
	Resource     *Resource `idl:"name:hResource" json:"resource"`
	ResourceName string    `idl:"name:lpszResourceName;string" json:"resource_name"`
}

func (o *SetResourceNameRequest) xxx_ToOp(ctx context.Context) *xxx_SetResourceNameOperation {
	if o == nil {
		return &xxx_SetResourceNameOperation{}
	}
	return &xxx_SetResourceNameOperation{
		Resource:     o.Resource,
		ResourceName: o.ResourceName,
	}
}

func (o *SetResourceNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetResourceNameOperation) {
	if o == nil {
		return
	}
	o.Resource = op.Resource
	o.ResourceName = op.ResourceName
}
func (o *SetResourceNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetResourceNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetResourceNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetResourceNameResponse structure represents the ApiSetResourceName operation response
type SetResourceNameResponse struct {
	// Return: The ApiSetResourceName return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetResourceNameResponse) xxx_ToOp(ctx context.Context) *xxx_SetResourceNameOperation {
	if o == nil {
		return &xxx_SetResourceNameOperation{}
	}
	return &xxx_SetResourceNameOperation{
		Return: o.Return,
	}
}

func (o *SetResourceNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetResourceNameOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetResourceNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetResourceNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetResourceNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetResourceIDOperation structure represents the ApiGetResourceId operation
type xxx_GetResourceIDOperation struct {
	Resource *Resource `idl:"name:hResource" json:"resource"`
	GUID     string    `idl:"name:pGuid;string" json:"guid"`
	Return   uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_GetResourceIDOperation) OpNum() int { return 14 }

func (o *xxx_GetResourceIDOperation) OpName() string { return "/clusapi2/v2/ApiGetResourceId" }

func (o *xxx_GetResourceIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResourceIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetResourceIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResourceIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResourceIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pGuid {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.GUID != "" {
			_ptr_pGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.GUID); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.GUID, _ptr_pGuid); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResourceIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pGuid {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.GUID); err != nil {
				return err
			}
			return nil
		})
		_s_pGuid := func(ptr interface{}) { o.GUID = *ptr.(*string) }
		if err := w.ReadPointer(&o.GUID, _s_pGuid, _ptr_pGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetResourceIDRequest structure represents the ApiGetResourceId operation request
type GetResourceIDRequest struct {
	Resource *Resource `idl:"name:hResource" json:"resource"`
}

func (o *GetResourceIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetResourceIDOperation {
	if o == nil {
		return &xxx_GetResourceIDOperation{}
	}
	return &xxx_GetResourceIDOperation{
		Resource: o.Resource,
	}
}

func (o *GetResourceIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetResourceIDOperation) {
	if o == nil {
		return
	}
	o.Resource = op.Resource
}
func (o *GetResourceIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetResourceIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResourceIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetResourceIDResponse structure represents the ApiGetResourceId operation response
type GetResourceIDResponse struct {
	GUID string `idl:"name:pGuid;string" json:"guid"`
	// Return: The ApiGetResourceId return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetResourceIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetResourceIDOperation {
	if o == nil {
		return &xxx_GetResourceIDOperation{}
	}
	return &xxx_GetResourceIDOperation{
		GUID:   o.GUID,
		Return: o.Return,
	}
}

func (o *GetResourceIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetResourceIDOperation) {
	if o == nil {
		return
	}
	o.GUID = op.GUID
	o.Return = op.Return
}
func (o *GetResourceIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetResourceIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResourceIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetResourceTypeOperation structure represents the ApiGetResourceType operation
type xxx_GetResourceTypeOperation struct {
	Resource     *Resource `idl:"name:hResource" json:"resource"`
	ResourceType string    `idl:"name:lpszResourceType;string" json:"resource_type"`
	Return       uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_GetResourceTypeOperation) OpNum() int { return 15 }

func (o *xxx_GetResourceTypeOperation) OpName() string { return "/clusapi2/v2/ApiGetResourceType" }

func (o *xxx_GetResourceTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResourceTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetResourceTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResourceTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResourceTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpszResourceType {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ResourceType != "" {
			_ptr_lpszResourceType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ResourceType); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ResourceType, _ptr_lpszResourceType); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetResourceTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpszResourceType {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpszResourceType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ResourceType); err != nil {
				return err
			}
			return nil
		})
		_s_lpszResourceType := func(ptr interface{}) { o.ResourceType = *ptr.(*string) }
		if err := w.ReadPointer(&o.ResourceType, _s_lpszResourceType, _ptr_lpszResourceType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetResourceTypeRequest structure represents the ApiGetResourceType operation request
type GetResourceTypeRequest struct {
	Resource *Resource `idl:"name:hResource" json:"resource"`
}

func (o *GetResourceTypeRequest) xxx_ToOp(ctx context.Context) *xxx_GetResourceTypeOperation {
	if o == nil {
		return &xxx_GetResourceTypeOperation{}
	}
	return &xxx_GetResourceTypeOperation{
		Resource: o.Resource,
	}
}

func (o *GetResourceTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetResourceTypeOperation) {
	if o == nil {
		return
	}
	o.Resource = op.Resource
}
func (o *GetResourceTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetResourceTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResourceTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetResourceTypeResponse structure represents the ApiGetResourceType operation response
type GetResourceTypeResponse struct {
	ResourceType string `idl:"name:lpszResourceType;string" json:"resource_type"`
	// Return: The ApiGetResourceType return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetResourceTypeResponse) xxx_ToOp(ctx context.Context) *xxx_GetResourceTypeOperation {
	if o == nil {
		return &xxx_GetResourceTypeOperation{}
	}
	return &xxx_GetResourceTypeOperation{
		ResourceType: o.ResourceType,
		Return:       o.Return,
	}
}

func (o *GetResourceTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetResourceTypeOperation) {
	if o == nil {
		return
	}
	o.ResourceType = op.ResourceType
	o.Return = op.Return
}
func (o *GetResourceTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetResourceTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetResourceTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FailResourceOperation structure represents the ApiFailResource operation
type xxx_FailResourceOperation struct {
	Resource *Resource `idl:"name:hResource" json:"resource"`
	Return   uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_FailResourceOperation) OpNum() int { return 16 }

func (o *xxx_FailResourceOperation) OpName() string { return "/clusapi2/v2/ApiFailResource" }

func (o *xxx_FailResourceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FailResourceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_FailResourceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FailResourceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FailResourceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FailResourceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FailResourceRequest structure represents the ApiFailResource operation request
type FailResourceRequest struct {
	Resource *Resource `idl:"name:hResource" json:"resource"`
}

func (o *FailResourceRequest) xxx_ToOp(ctx context.Context) *xxx_FailResourceOperation {
	if o == nil {
		return &xxx_FailResourceOperation{}
	}
	return &xxx_FailResourceOperation{
		Resource: o.Resource,
	}
}

func (o *FailResourceRequest) xxx_FromOp(ctx context.Context, op *xxx_FailResourceOperation) {
	if o == nil {
		return
	}
	o.Resource = op.Resource
}
func (o *FailResourceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *FailResourceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FailResourceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FailResourceResponse structure represents the ApiFailResource operation response
type FailResourceResponse struct {
	// Return: The ApiFailResource return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *FailResourceResponse) xxx_ToOp(ctx context.Context) *xxx_FailResourceOperation {
	if o == nil {
		return &xxx_FailResourceOperation{}
	}
	return &xxx_FailResourceOperation{
		Return: o.Return,
	}
}

func (o *FailResourceResponse) xxx_FromOp(ctx context.Context, op *xxx_FailResourceOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *FailResourceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *FailResourceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FailResourceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OnlineResourceOperation structure represents the ApiOnlineResource operation
type xxx_OnlineResourceOperation struct {
	Resource *Resource `idl:"name:hResource" json:"resource"`
	Return   uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_OnlineResourceOperation) OpNum() int { return 17 }

func (o *xxx_OnlineResourceOperation) OpName() string { return "/clusapi2/v2/ApiOnlineResource" }

func (o *xxx_OnlineResourceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnlineResourceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_OnlineResourceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnlineResourceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnlineResourceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnlineResourceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OnlineResourceRequest structure represents the ApiOnlineResource operation request
type OnlineResourceRequest struct {
	Resource *Resource `idl:"name:hResource" json:"resource"`
}

func (o *OnlineResourceRequest) xxx_ToOp(ctx context.Context) *xxx_OnlineResourceOperation {
	if o == nil {
		return &xxx_OnlineResourceOperation{}
	}
	return &xxx_OnlineResourceOperation{
		Resource: o.Resource,
	}
}

func (o *OnlineResourceRequest) xxx_FromOp(ctx context.Context, op *xxx_OnlineResourceOperation) {
	if o == nil {
		return
	}
	o.Resource = op.Resource
}
func (o *OnlineResourceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OnlineResourceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OnlineResourceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OnlineResourceResponse structure represents the ApiOnlineResource operation response
type OnlineResourceResponse struct {
	// Return: The ApiOnlineResource return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *OnlineResourceResponse) xxx_ToOp(ctx context.Context) *xxx_OnlineResourceOperation {
	if o == nil {
		return &xxx_OnlineResourceOperation{}
	}
	return &xxx_OnlineResourceOperation{
		Return: o.Return,
	}
}

func (o *OnlineResourceResponse) xxx_FromOp(ctx context.Context, op *xxx_OnlineResourceOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *OnlineResourceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OnlineResourceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OnlineResourceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OfflineResourceOperation structure represents the ApiOfflineResource operation
type xxx_OfflineResourceOperation struct {
	Resource *Resource `idl:"name:hResource" json:"resource"`
	Return   uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_OfflineResourceOperation) OpNum() int { return 18 }

func (o *xxx_OfflineResourceOperation) OpName() string { return "/clusapi2/v2/ApiOfflineResource" }

func (o *xxx_OfflineResourceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OfflineResourceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_OfflineResourceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OfflineResourceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OfflineResourceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OfflineResourceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OfflineResourceRequest structure represents the ApiOfflineResource operation request
type OfflineResourceRequest struct {
	Resource *Resource `idl:"name:hResource" json:"resource"`
}

func (o *OfflineResourceRequest) xxx_ToOp(ctx context.Context) *xxx_OfflineResourceOperation {
	if o == nil {
		return &xxx_OfflineResourceOperation{}
	}
	return &xxx_OfflineResourceOperation{
		Resource: o.Resource,
	}
}

func (o *OfflineResourceRequest) xxx_FromOp(ctx context.Context, op *xxx_OfflineResourceOperation) {
	if o == nil {
		return
	}
	o.Resource = op.Resource
}
func (o *OfflineResourceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OfflineResourceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OfflineResourceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OfflineResourceResponse structure represents the ApiOfflineResource operation response
type OfflineResourceResponse struct {
	// Return: The ApiOfflineResource return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *OfflineResourceResponse) xxx_ToOp(ctx context.Context) *xxx_OfflineResourceOperation {
	if o == nil {
		return &xxx_OfflineResourceOperation{}
	}
	return &xxx_OfflineResourceOperation{
		Return: o.Return,
	}
}

func (o *OfflineResourceResponse) xxx_FromOp(ctx context.Context, op *xxx_OfflineResourceOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *OfflineResourceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OfflineResourceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OfflineResourceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddResourceDependencyOperation structure represents the ApiAddResourceDependency operation
type xxx_AddResourceDependencyOperation struct {
	Resource   *Resource `idl:"name:hResource" json:"resource"`
	DepdendsOn *Resource `idl:"name:hDependsOn" json:"depdends_on"`
	Return     uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_AddResourceDependencyOperation) OpNum() int { return 19 }

func (o *xxx_AddResourceDependencyOperation) OpName() string {
	return "/clusapi2/v2/ApiAddResourceDependency"
}

func (o *xxx_AddResourceDependencyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddResourceDependencyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hDependsOn {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.DepdendsOn != nil {
			if err := o.DepdendsOn.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_AddResourceDependencyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hDependsOn {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.DepdendsOn == nil {
			o.DepdendsOn = &Resource{}
		}
		if err := o.DepdendsOn.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddResourceDependencyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddResourceDependencyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddResourceDependencyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddResourceDependencyRequest structure represents the ApiAddResourceDependency operation request
type AddResourceDependencyRequest struct {
	Resource   *Resource `idl:"name:hResource" json:"resource"`
	DepdendsOn *Resource `idl:"name:hDependsOn" json:"depdends_on"`
}

func (o *AddResourceDependencyRequest) xxx_ToOp(ctx context.Context) *xxx_AddResourceDependencyOperation {
	if o == nil {
		return &xxx_AddResourceDependencyOperation{}
	}
	return &xxx_AddResourceDependencyOperation{
		Resource:   o.Resource,
		DepdendsOn: o.DepdendsOn,
	}
}

func (o *AddResourceDependencyRequest) xxx_FromOp(ctx context.Context, op *xxx_AddResourceDependencyOperation) {
	if o == nil {
		return
	}
	o.Resource = op.Resource
	o.DepdendsOn = op.DepdendsOn
}
func (o *AddResourceDependencyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddResourceDependencyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddResourceDependencyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddResourceDependencyResponse structure represents the ApiAddResourceDependency operation response
type AddResourceDependencyResponse struct {
	// Return: The ApiAddResourceDependency return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AddResourceDependencyResponse) xxx_ToOp(ctx context.Context) *xxx_AddResourceDependencyOperation {
	if o == nil {
		return &xxx_AddResourceDependencyOperation{}
	}
	return &xxx_AddResourceDependencyOperation{
		Return: o.Return,
	}
}

func (o *AddResourceDependencyResponse) xxx_FromOp(ctx context.Context, op *xxx_AddResourceDependencyOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AddResourceDependencyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddResourceDependencyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddResourceDependencyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveResourceDependencyOperation structure represents the ApiRemoveResourceDependency operation
type xxx_RemoveResourceDependencyOperation struct {
	Resource   *Resource `idl:"name:hResource" json:"resource"`
	DepdendsOn *Resource `idl:"name:hDependsOn" json:"depdends_on"`
	Return     uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveResourceDependencyOperation) OpNum() int { return 20 }

func (o *xxx_RemoveResourceDependencyOperation) OpName() string {
	return "/clusapi2/v2/ApiRemoveResourceDependency"
}

func (o *xxx_RemoveResourceDependencyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveResourceDependencyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hDependsOn {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.DepdendsOn != nil {
			if err := o.DepdendsOn.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_RemoveResourceDependencyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hDependsOn {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.DepdendsOn == nil {
			o.DepdendsOn = &Resource{}
		}
		if err := o.DepdendsOn.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveResourceDependencyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveResourceDependencyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveResourceDependencyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveResourceDependencyRequest structure represents the ApiRemoveResourceDependency operation request
type RemoveResourceDependencyRequest struct {
	Resource   *Resource `idl:"name:hResource" json:"resource"`
	DepdendsOn *Resource `idl:"name:hDependsOn" json:"depdends_on"`
}

func (o *RemoveResourceDependencyRequest) xxx_ToOp(ctx context.Context) *xxx_RemoveResourceDependencyOperation {
	if o == nil {
		return &xxx_RemoveResourceDependencyOperation{}
	}
	return &xxx_RemoveResourceDependencyOperation{
		Resource:   o.Resource,
		DepdendsOn: o.DepdendsOn,
	}
}

func (o *RemoveResourceDependencyRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveResourceDependencyOperation) {
	if o == nil {
		return
	}
	o.Resource = op.Resource
	o.DepdendsOn = op.DepdendsOn
}
func (o *RemoveResourceDependencyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoveResourceDependencyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveResourceDependencyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveResourceDependencyResponse structure represents the ApiRemoveResourceDependency operation response
type RemoveResourceDependencyResponse struct {
	// Return: The ApiRemoveResourceDependency return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RemoveResourceDependencyResponse) xxx_ToOp(ctx context.Context) *xxx_RemoveResourceDependencyOperation {
	if o == nil {
		return &xxx_RemoveResourceDependencyOperation{}
	}
	return &xxx_RemoveResourceDependencyOperation{
		Return: o.Return,
	}
}

func (o *RemoveResourceDependencyResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveResourceDependencyOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RemoveResourceDependencyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoveResourceDependencyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveResourceDependencyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CanResourceBeDependentOperation structure represents the ApiCanResourceBeDependent operation
type xxx_CanResourceBeDependentOperation struct {
	Resource          *Resource `idl:"name:hResource" json:"resource"`
	ResourceDependent *Resource `idl:"name:hResourceDependent" json:"resource_dependent"`
	Return            uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_CanResourceBeDependentOperation) OpNum() int { return 21 }

func (o *xxx_CanResourceBeDependentOperation) OpName() string {
	return "/clusapi2/v2/ApiCanResourceBeDependent"
}

func (o *xxx_CanResourceBeDependentOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CanResourceBeDependentOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hResourceDependent {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.ResourceDependent != nil {
			if err := o.ResourceDependent.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CanResourceBeDependentOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hResourceDependent {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.ResourceDependent == nil {
			o.ResourceDependent = &Resource{}
		}
		if err := o.ResourceDependent.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CanResourceBeDependentOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CanResourceBeDependentOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CanResourceBeDependentOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CanResourceBeDependentRequest structure represents the ApiCanResourceBeDependent operation request
type CanResourceBeDependentRequest struct {
	Resource          *Resource `idl:"name:hResource" json:"resource"`
	ResourceDependent *Resource `idl:"name:hResourceDependent" json:"resource_dependent"`
}

func (o *CanResourceBeDependentRequest) xxx_ToOp(ctx context.Context) *xxx_CanResourceBeDependentOperation {
	if o == nil {
		return &xxx_CanResourceBeDependentOperation{}
	}
	return &xxx_CanResourceBeDependentOperation{
		Resource:          o.Resource,
		ResourceDependent: o.ResourceDependent,
	}
}

func (o *CanResourceBeDependentRequest) xxx_FromOp(ctx context.Context, op *xxx_CanResourceBeDependentOperation) {
	if o == nil {
		return
	}
	o.Resource = op.Resource
	o.ResourceDependent = op.ResourceDependent
}
func (o *CanResourceBeDependentRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CanResourceBeDependentRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CanResourceBeDependentOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CanResourceBeDependentResponse structure represents the ApiCanResourceBeDependent operation response
type CanResourceBeDependentResponse struct {
	// Return: The ApiCanResourceBeDependent return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CanResourceBeDependentResponse) xxx_ToOp(ctx context.Context) *xxx_CanResourceBeDependentOperation {
	if o == nil {
		return &xxx_CanResourceBeDependentOperation{}
	}
	return &xxx_CanResourceBeDependentOperation{
		Return: o.Return,
	}
}

func (o *CanResourceBeDependentResponse) xxx_FromOp(ctx context.Context, op *xxx_CanResourceBeDependentOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *CanResourceBeDependentResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CanResourceBeDependentResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CanResourceBeDependentOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateRestrictionEnumOperation structure represents the ApiCreateResEnum operation
type xxx_CreateRestrictionEnumOperation struct {
	Resource   *Resource `idl:"name:hResource" json:"resource"`
	Type       uint32    `idl:"name:dwType" json:"type"`
	ReturnEnum *EnumList `idl:"name:ReturnEnum" json:"return_enum"`
	Return     uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateRestrictionEnumOperation) OpNum() int { return 22 }

func (o *xxx_CreateRestrictionEnumOperation) OpName() string { return "/clusapi2/v2/ApiCreateResEnum" }

func (o *xxx_CreateRestrictionEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRestrictionEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRestrictionEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRestrictionEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRestrictionEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ReturnEnum {out} (1:{pointer=ref}*(2))(2:{alias=PENUM_LIST}*(1))(3:{alias=ENUM_LIST}(struct))
	{
		if o.ReturnEnum != nil {
			_ptr_ReturnEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnEnum != nil {
					if err := o.ReturnEnum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EnumList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnEnum, _ptr_ReturnEnum); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRestrictionEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ReturnEnum {out} (1:{pointer=ref}*(2))(2:{alias=PENUM_LIST,pointer=ref}*(1))(3:{alias=ENUM_LIST}(struct))
	{
		_ptr_ReturnEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnEnum == nil {
				o.ReturnEnum = &EnumList{}
			}
			if err := o.ReturnEnum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReturnEnum := func(ptr interface{}) { o.ReturnEnum = *ptr.(**EnumList) }
		if err := w.ReadPointer(&o.ReturnEnum, _s_ReturnEnum, _ptr_ReturnEnum); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateRestrictionEnumRequest structure represents the ApiCreateResEnum operation request
type CreateRestrictionEnumRequest struct {
	Resource *Resource `idl:"name:hResource" json:"resource"`
	Type     uint32    `idl:"name:dwType" json:"type"`
}

func (o *CreateRestrictionEnumRequest) xxx_ToOp(ctx context.Context) *xxx_CreateRestrictionEnumOperation {
	if o == nil {
		return &xxx_CreateRestrictionEnumOperation{}
	}
	return &xxx_CreateRestrictionEnumOperation{
		Resource: o.Resource,
		Type:     o.Type,
	}
}

func (o *CreateRestrictionEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateRestrictionEnumOperation) {
	if o == nil {
		return
	}
	o.Resource = op.Resource
	o.Type = op.Type
}
func (o *CreateRestrictionEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateRestrictionEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateRestrictionEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateRestrictionEnumResponse structure represents the ApiCreateResEnum operation response
type CreateRestrictionEnumResponse struct {
	ReturnEnum *EnumList `idl:"name:ReturnEnum" json:"return_enum"`
	// Return: The ApiCreateResEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CreateRestrictionEnumResponse) xxx_ToOp(ctx context.Context) *xxx_CreateRestrictionEnumOperation {
	if o == nil {
		return &xxx_CreateRestrictionEnumOperation{}
	}
	return &xxx_CreateRestrictionEnumOperation{
		ReturnEnum: o.ReturnEnum,
		Return:     o.Return,
	}
}

func (o *CreateRestrictionEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateRestrictionEnumOperation) {
	if o == nil {
		return
	}
	o.ReturnEnum = op.ReturnEnum
	o.Return = op.Return
}
func (o *CreateRestrictionEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateRestrictionEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateRestrictionEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddResourceNodeOperation structure represents the ApiAddResourceNode operation
type xxx_AddResourceNodeOperation struct {
	Resource *Resource `idl:"name:hResource" json:"resource"`
	Node     *Node     `idl:"name:hNode" json:"node"`
	Return   uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_AddResourceNodeOperation) OpNum() int { return 23 }

func (o *xxx_AddResourceNodeOperation) OpName() string { return "/clusapi2/v2/ApiAddResourceNode" }

func (o *xxx_AddResourceNodeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddResourceNodeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node != nil {
			if err := o.Node.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_AddResourceNodeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node == nil {
			o.Node = &Node{}
		}
		if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddResourceNodeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddResourceNodeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddResourceNodeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddResourceNodeRequest structure represents the ApiAddResourceNode operation request
type AddResourceNodeRequest struct {
	Resource *Resource `idl:"name:hResource" json:"resource"`
	Node     *Node     `idl:"name:hNode" json:"node"`
}

func (o *AddResourceNodeRequest) xxx_ToOp(ctx context.Context) *xxx_AddResourceNodeOperation {
	if o == nil {
		return &xxx_AddResourceNodeOperation{}
	}
	return &xxx_AddResourceNodeOperation{
		Resource: o.Resource,
		Node:     o.Node,
	}
}

func (o *AddResourceNodeRequest) xxx_FromOp(ctx context.Context, op *xxx_AddResourceNodeOperation) {
	if o == nil {
		return
	}
	o.Resource = op.Resource
	o.Node = op.Node
}
func (o *AddResourceNodeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddResourceNodeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddResourceNodeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddResourceNodeResponse structure represents the ApiAddResourceNode operation response
type AddResourceNodeResponse struct {
	// Return: The ApiAddResourceNode return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AddResourceNodeResponse) xxx_ToOp(ctx context.Context) *xxx_AddResourceNodeOperation {
	if o == nil {
		return &xxx_AddResourceNodeOperation{}
	}
	return &xxx_AddResourceNodeOperation{
		Return: o.Return,
	}
}

func (o *AddResourceNodeResponse) xxx_FromOp(ctx context.Context, op *xxx_AddResourceNodeOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AddResourceNodeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddResourceNodeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddResourceNodeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveResourceNodeOperation structure represents the ApiRemoveResourceNode operation
type xxx_RemoveResourceNodeOperation struct {
	Resource *Resource `idl:"name:hResource" json:"resource"`
	Node     *Node     `idl:"name:hNode" json:"node"`
	Return   uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveResourceNodeOperation) OpNum() int { return 24 }

func (o *xxx_RemoveResourceNodeOperation) OpName() string {
	return "/clusapi2/v2/ApiRemoveResourceNode"
}

func (o *xxx_RemoveResourceNodeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveResourceNodeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node != nil {
			if err := o.Node.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_RemoveResourceNodeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node == nil {
			o.Node = &Node{}
		}
		if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveResourceNodeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveResourceNodeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveResourceNodeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveResourceNodeRequest structure represents the ApiRemoveResourceNode operation request
type RemoveResourceNodeRequest struct {
	Resource *Resource `idl:"name:hResource" json:"resource"`
	Node     *Node     `idl:"name:hNode" json:"node"`
}

func (o *RemoveResourceNodeRequest) xxx_ToOp(ctx context.Context) *xxx_RemoveResourceNodeOperation {
	if o == nil {
		return &xxx_RemoveResourceNodeOperation{}
	}
	return &xxx_RemoveResourceNodeOperation{
		Resource: o.Resource,
		Node:     o.Node,
	}
}

func (o *RemoveResourceNodeRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveResourceNodeOperation) {
	if o == nil {
		return
	}
	o.Resource = op.Resource
	o.Node = op.Node
}
func (o *RemoveResourceNodeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RemoveResourceNodeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveResourceNodeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveResourceNodeResponse structure represents the ApiRemoveResourceNode operation response
type RemoveResourceNodeResponse struct {
	// Return: The ApiRemoveResourceNode return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RemoveResourceNodeResponse) xxx_ToOp(ctx context.Context) *xxx_RemoveResourceNodeOperation {
	if o == nil {
		return &xxx_RemoveResourceNodeOperation{}
	}
	return &xxx_RemoveResourceNodeOperation{
		Return: o.Return,
	}
}

func (o *RemoveResourceNodeResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveResourceNodeOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RemoveResourceNodeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RemoveResourceNodeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveResourceNodeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ChangeResourceGroupOperation structure represents the ApiChangeResourceGroup operation
type xxx_ChangeResourceGroupOperation struct {
	Resource *Resource `idl:"name:hResource" json:"resource"`
	Group    *Group    `idl:"name:hGroup" json:"group"`
	Return   uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_ChangeResourceGroupOperation) OpNum() int { return 25 }

func (o *xxx_ChangeResourceGroupOperation) OpName() string {
	return "/clusapi2/v2/ApiChangeResourceGroup"
}

func (o *xxx_ChangeResourceGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangeResourceGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group != nil {
			if err := o.Group.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Group{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ChangeResourceGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group == nil {
			o.Group = &Group{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangeResourceGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangeResourceGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ChangeResourceGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ChangeResourceGroupRequest structure represents the ApiChangeResourceGroup operation request
type ChangeResourceGroupRequest struct {
	Resource *Resource `idl:"name:hResource" json:"resource"`
	Group    *Group    `idl:"name:hGroup" json:"group"`
}

func (o *ChangeResourceGroupRequest) xxx_ToOp(ctx context.Context) *xxx_ChangeResourceGroupOperation {
	if o == nil {
		return &xxx_ChangeResourceGroupOperation{}
	}
	return &xxx_ChangeResourceGroupOperation{
		Resource: o.Resource,
		Group:    o.Group,
	}
}

func (o *ChangeResourceGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_ChangeResourceGroupOperation) {
	if o == nil {
		return
	}
	o.Resource = op.Resource
	o.Group = op.Group
}
func (o *ChangeResourceGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ChangeResourceGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ChangeResourceGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ChangeResourceGroupResponse structure represents the ApiChangeResourceGroup operation response
type ChangeResourceGroupResponse struct {
	// Return: The ApiChangeResourceGroup return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ChangeResourceGroupResponse) xxx_ToOp(ctx context.Context) *xxx_ChangeResourceGroupOperation {
	if o == nil {
		return &xxx_ChangeResourceGroupOperation{}
	}
	return &xxx_ChangeResourceGroupOperation{
		Return: o.Return,
	}
}

func (o *ChangeResourceGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_ChangeResourceGroupOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ChangeResourceGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ChangeResourceGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ChangeResourceGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateResourceTypeOperation structure represents the ApiCreateResourceType operation
type xxx_CreateResourceTypeOperation struct {
	TypeName    string `idl:"name:lpszTypeName;string" json:"type_name"`
	DisplayName string `idl:"name:lpszDisplayName;string" json:"display_name"`
	DLLName     string `idl:"name:lpszDllName;string" json:"dll_name"`
	LooksAlive  uint32 `idl:"name:dwLooksAlive" json:"looks_alive"`
	IsAlive     uint32 `idl:"name:dwIsAlive" json:"is_alive"`
	Return      uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateResourceTypeOperation) OpNum() int { return 26 }

func (o *xxx_CreateResourceTypeOperation) OpName() string {
	return "/clusapi2/v2/ApiCreateResourceType"
}

func (o *xxx_CreateResourceTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateResourceTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpszTypeName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.TypeName); err != nil {
			return err
		}
	}
	// lpszDisplayName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DisplayName); err != nil {
			return err
		}
	}
	// lpszDllName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DLLName); err != nil {
			return err
		}
	}
	// dwLooksAlive {in} (1:(uint32))
	{
		if err := w.WriteData(o.LooksAlive); err != nil {
			return err
		}
	}
	// dwIsAlive {in} (1:(uint32))
	{
		if err := w.WriteData(o.IsAlive); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateResourceTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpszTypeName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.TypeName); err != nil {
			return err
		}
	}
	// lpszDisplayName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DisplayName); err != nil {
			return err
		}
	}
	// lpszDllName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DLLName); err != nil {
			return err
		}
	}
	// dwLooksAlive {in} (1:(uint32))
	{
		if err := w.ReadData(&o.LooksAlive); err != nil {
			return err
		}
	}
	// dwIsAlive {in} (1:(uint32))
	{
		if err := w.ReadData(&o.IsAlive); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateResourceTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateResourceTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateResourceTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateResourceTypeRequest structure represents the ApiCreateResourceType operation request
type CreateResourceTypeRequest struct {
	TypeName    string `idl:"name:lpszTypeName;string" json:"type_name"`
	DisplayName string `idl:"name:lpszDisplayName;string" json:"display_name"`
	DLLName     string `idl:"name:lpszDllName;string" json:"dll_name"`
	LooksAlive  uint32 `idl:"name:dwLooksAlive" json:"looks_alive"`
	IsAlive     uint32 `idl:"name:dwIsAlive" json:"is_alive"`
}

func (o *CreateResourceTypeRequest) xxx_ToOp(ctx context.Context) *xxx_CreateResourceTypeOperation {
	if o == nil {
		return &xxx_CreateResourceTypeOperation{}
	}
	return &xxx_CreateResourceTypeOperation{
		TypeName:    o.TypeName,
		DisplayName: o.DisplayName,
		DLLName:     o.DLLName,
		LooksAlive:  o.LooksAlive,
		IsAlive:     o.IsAlive,
	}
}

func (o *CreateResourceTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateResourceTypeOperation) {
	if o == nil {
		return
	}
	o.TypeName = op.TypeName
	o.DisplayName = op.DisplayName
	o.DLLName = op.DLLName
	o.LooksAlive = op.LooksAlive
	o.IsAlive = op.IsAlive
}
func (o *CreateResourceTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateResourceTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateResourceTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateResourceTypeResponse structure represents the ApiCreateResourceType operation response
type CreateResourceTypeResponse struct {
	// Return: The ApiCreateResourceType return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CreateResourceTypeResponse) xxx_ToOp(ctx context.Context) *xxx_CreateResourceTypeOperation {
	if o == nil {
		return &xxx_CreateResourceTypeOperation{}
	}
	return &xxx_CreateResourceTypeOperation{
		Return: o.Return,
	}
}

func (o *CreateResourceTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateResourceTypeOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *CreateResourceTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateResourceTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateResourceTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteResourceTypeOperation structure represents the ApiDeleteResourceType operation
type xxx_DeleteResourceTypeOperation struct {
	TypeName string `idl:"name:lpszTypeName;string" json:"type_name"`
	Return   uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteResourceTypeOperation) OpNum() int { return 27 }

func (o *xxx_DeleteResourceTypeOperation) OpName() string {
	return "/clusapi2/v2/ApiDeleteResourceType"
}

func (o *xxx_DeleteResourceTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteResourceTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpszTypeName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.TypeName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteResourceTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpszTypeName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.TypeName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteResourceTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteResourceTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteResourceTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteResourceTypeRequest structure represents the ApiDeleteResourceType operation request
type DeleteResourceTypeRequest struct {
	TypeName string `idl:"name:lpszTypeName;string" json:"type_name"`
}

func (o *DeleteResourceTypeRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteResourceTypeOperation {
	if o == nil {
		return &xxx_DeleteResourceTypeOperation{}
	}
	return &xxx_DeleteResourceTypeOperation{
		TypeName: o.TypeName,
	}
}

func (o *DeleteResourceTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteResourceTypeOperation) {
	if o == nil {
		return
	}
	o.TypeName = op.TypeName
}
func (o *DeleteResourceTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteResourceTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteResourceTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteResourceTypeResponse structure represents the ApiDeleteResourceType operation response
type DeleteResourceTypeResponse struct {
	// Return: The ApiDeleteResourceType return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *DeleteResourceTypeResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteResourceTypeOperation {
	if o == nil {
		return &xxx_DeleteResourceTypeOperation{}
	}
	return &xxx_DeleteResourceTypeOperation{
		Return: o.Return,
	}
}

func (o *DeleteResourceTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteResourceTypeOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DeleteResourceTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteResourceTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteResourceTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRootKeyOperation structure represents the ApiGetRootKey operation
type xxx_GetRootKeyOperation struct {
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
	Status        uint32 `idl:"name:Status" json:"status"`
	Return        *Key   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRootKeyOperation) OpNum() int { return 28 }

func (o *xxx_GetRootKeyOperation) OpName() string { return "/clusapi2/v2/ApiGetRootKey" }

func (o *xxx_GetRootKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRootKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// samDesired {in} (1:(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRootKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// samDesired {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRootKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRootKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return != nil {
			if err := o.Return.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetRootKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return == nil {
			o.Return = &Key{}
		}
		if err := o.Return.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// GetRootKeyRequest structure represents the ApiGetRootKey operation request
type GetRootKeyRequest struct {
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
}

func (o *GetRootKeyRequest) xxx_ToOp(ctx context.Context) *xxx_GetRootKeyOperation {
	if o == nil {
		return &xxx_GetRootKeyOperation{}
	}
	return &xxx_GetRootKeyOperation{
		DesiredAccess: o.DesiredAccess,
	}
}

func (o *GetRootKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRootKeyOperation) {
	if o == nil {
		return
	}
	o.DesiredAccess = op.DesiredAccess
}
func (o *GetRootKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetRootKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRootKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRootKeyResponse structure represents the ApiGetRootKey operation response
type GetRootKeyResponse struct {
	Status uint32 `idl:"name:Status" json:"status"`
	// Return: The ApiGetRootKey return value.
	Return *Key `idl:"name:Return" json:"return"`
}

func (o *GetRootKeyResponse) xxx_ToOp(ctx context.Context) *xxx_GetRootKeyOperation {
	if o == nil {
		return &xxx_GetRootKeyOperation{}
	}
	return &xxx_GetRootKeyOperation{
		Status: o.Status,
		Return: o.Return,
	}
}

func (o *GetRootKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRootKeyOperation) {
	if o == nil {
		return
	}
	o.Status = op.Status
	o.Return = op.Return
}
func (o *GetRootKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetRootKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRootKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateKeyOperation structure represents the ApiCreateKey operation
type xxx_CreateKeyOperation struct {
	Key                *Key                `idl:"name:hKey" json:"key"`
	SubKey             string              `idl:"name:lpSubKey;string" json:"sub_key"`
	Options            uint32              `idl:"name:dwOptions" json:"options"`
	DesiredAccess      uint32              `idl:"name:samDesired" json:"desired_access"`
	SecurityAttributes *SecurityAttributes `idl:"name:lpSecurityAttributes;pointer:unique" json:"security_attributes"`
	Disposition        uint32              `idl:"name:lpdwDisposition" json:"disposition"`
	Status             uint32              `idl:"name:Status" json:"status"`
	Return             *Key                `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateKeyOperation) OpNum() int { return 29 }

func (o *xxx_CreateKeyOperation) OpName() string { return "/clusapi2/v2/ApiCreateKey" }

func (o *xxx_CreateKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpSubKey {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.SubKey); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	// samDesired {in} (1:(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	// lpSecurityAttributes {in} (1:{pointer=unique, alias=PRPC_SECURITY_ATTRIBUTES}*(1))(2:{alias=RPC_SECURITY_ATTRIBUTES}(struct))
	{
		if o.SecurityAttributes != nil {
			_ptr_lpSecurityAttributes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SecurityAttributes != nil {
					if err := o.SecurityAttributes.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&SecurityAttributes{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SecurityAttributes, _ptr_lpSecurityAttributes); err != nil {
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

func (o *xxx_CreateKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpSubKey {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.SubKey); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	// samDesired {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	// lpSecurityAttributes {in} (1:{pointer=unique, alias=PRPC_SECURITY_ATTRIBUTES}*(1))(2:{alias=RPC_SECURITY_ATTRIBUTES}(struct))
	{
		_ptr_lpSecurityAttributes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SecurityAttributes == nil {
				o.SecurityAttributes = &SecurityAttributes{}
			}
			if err := o.SecurityAttributes.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpSecurityAttributes := func(ptr interface{}) { o.SecurityAttributes = *ptr.(**SecurityAttributes) }
		if err := w.ReadPointer(&o.SecurityAttributes, _s_lpSecurityAttributes, _ptr_lpSecurityAttributes); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpdwDisposition {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Disposition); err != nil {
			return err
		}
	}
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return != nil {
			if err := o.Return.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CreateKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpdwDisposition {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Disposition); err != nil {
			return err
		}
	}
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return == nil {
			o.Return = &Key{}
		}
		if err := o.Return.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// CreateKeyRequest structure represents the ApiCreateKey operation request
type CreateKeyRequest struct {
	Key                *Key                `idl:"name:hKey" json:"key"`
	SubKey             string              `idl:"name:lpSubKey;string" json:"sub_key"`
	Options            uint32              `idl:"name:dwOptions" json:"options"`
	DesiredAccess      uint32              `idl:"name:samDesired" json:"desired_access"`
	SecurityAttributes *SecurityAttributes `idl:"name:lpSecurityAttributes;pointer:unique" json:"security_attributes"`
}

func (o *CreateKeyRequest) xxx_ToOp(ctx context.Context) *xxx_CreateKeyOperation {
	if o == nil {
		return &xxx_CreateKeyOperation{}
	}
	return &xxx_CreateKeyOperation{
		Key:                o.Key,
		SubKey:             o.SubKey,
		Options:            o.Options,
		DesiredAccess:      o.DesiredAccess,
		SecurityAttributes: o.SecurityAttributes,
	}
}

func (o *CreateKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateKeyOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.SubKey = op.SubKey
	o.Options = op.Options
	o.DesiredAccess = op.DesiredAccess
	o.SecurityAttributes = op.SecurityAttributes
}
func (o *CreateKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateKeyResponse structure represents the ApiCreateKey operation response
type CreateKeyResponse struct {
	Disposition uint32 `idl:"name:lpdwDisposition" json:"disposition"`
	Status      uint32 `idl:"name:Status" json:"status"`
	// Return: The ApiCreateKey return value.
	Return *Key `idl:"name:Return" json:"return"`
}

func (o *CreateKeyResponse) xxx_ToOp(ctx context.Context) *xxx_CreateKeyOperation {
	if o == nil {
		return &xxx_CreateKeyOperation{}
	}
	return &xxx_CreateKeyOperation{
		Disposition: o.Disposition,
		Status:      o.Status,
		Return:      o.Return,
	}
}

func (o *CreateKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateKeyOperation) {
	if o == nil {
		return
	}
	o.Disposition = op.Disposition
	o.Status = op.Status
	o.Return = op.Return
}
func (o *CreateKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenKeyOperation structure represents the ApiOpenKey operation
type xxx_OpenKeyOperation struct {
	Key           *Key   `idl:"name:hKey" json:"key"`
	SubKey        string `idl:"name:lpSubKey;string" json:"sub_key"`
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
	Status        uint32 `idl:"name:Status" json:"status"`
	Return        *Key   `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenKeyOperation) OpNum() int { return 30 }

func (o *xxx_OpenKeyOperation) OpName() string { return "/clusapi2/v2/ApiOpenKey" }

func (o *xxx_OpenKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpSubKey {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.SubKey); err != nil {
			return err
		}
	}
	// samDesired {in} (1:(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpSubKey {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.SubKey); err != nil {
			return err
		}
	}
	// samDesired {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return != nil {
			if err := o.Return.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_OpenKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return == nil {
			o.Return = &Key{}
		}
		if err := o.Return.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// OpenKeyRequest structure represents the ApiOpenKey operation request
type OpenKeyRequest struct {
	Key           *Key   `idl:"name:hKey" json:"key"`
	SubKey        string `idl:"name:lpSubKey;string" json:"sub_key"`
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
}

func (o *OpenKeyRequest) xxx_ToOp(ctx context.Context) *xxx_OpenKeyOperation {
	if o == nil {
		return &xxx_OpenKeyOperation{}
	}
	return &xxx_OpenKeyOperation{
		Key:           o.Key,
		SubKey:        o.SubKey,
		DesiredAccess: o.DesiredAccess,
	}
}

func (o *OpenKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenKeyOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.SubKey = op.SubKey
	o.DesiredAccess = op.DesiredAccess
}
func (o *OpenKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OpenKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenKeyResponse structure represents the ApiOpenKey operation response
type OpenKeyResponse struct {
	Status uint32 `idl:"name:Status" json:"status"`
	// Return: The ApiOpenKey return value.
	Return *Key `idl:"name:Return" json:"return"`
}

func (o *OpenKeyResponse) xxx_ToOp(ctx context.Context) *xxx_OpenKeyOperation {
	if o == nil {
		return &xxx_OpenKeyOperation{}
	}
	return &xxx_OpenKeyOperation{
		Status: o.Status,
		Return: o.Return,
	}
}

func (o *OpenKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenKeyOperation) {
	if o == nil {
		return
	}
	o.Status = op.Status
	o.Return = op.Return
}
func (o *OpenKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OpenKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumKeyOperation structure represents the ApiEnumKey operation
type xxx_EnumKeyOperation struct {
	Key           *Key           `idl:"name:hKey" json:"key"`
	Index         uint32         `idl:"name:dwIndex" json:"index"`
	KeyName       string         `idl:"name:KeyName;string" json:"key_name"`
	LastWriteTime *dtyp.Filetime `idl:"name:lpftLastWriteTime" json:"last_write_time"`
	Return        uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumKeyOperation) OpNum() int { return 31 }

func (o *xxx_EnumKeyOperation) OpName() string { return "/clusapi2/v2/ApiEnumKey" }

func (o *xxx_EnumKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwIndex {in} (1:(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwIndex {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// KeyName {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		if o.KeyName != "" {
			_ptr_KeyName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.KeyName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.KeyName, _ptr_KeyName); err != nil {
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
	// lpftLastWriteTime {out} (1:{pointer=ref}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.LastWriteTime != nil {
			if err := o.LastWriteTime.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// KeyName {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		_ptr_KeyName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.KeyName); err != nil {
				return err
			}
			return nil
		})
		_s_KeyName := func(ptr interface{}) { o.KeyName = *ptr.(*string) }
		if err := w.ReadPointer(&o.KeyName, _s_KeyName, _ptr_KeyName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpftLastWriteTime {out} (1:{pointer=ref}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.LastWriteTime == nil {
			o.LastWriteTime = &dtyp.Filetime{}
		}
		if err := o.LastWriteTime.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumKeyRequest structure represents the ApiEnumKey operation request
type EnumKeyRequest struct {
	Key   *Key   `idl:"name:hKey" json:"key"`
	Index uint32 `idl:"name:dwIndex" json:"index"`
}

func (o *EnumKeyRequest) xxx_ToOp(ctx context.Context) *xxx_EnumKeyOperation {
	if o == nil {
		return &xxx_EnumKeyOperation{}
	}
	return &xxx_EnumKeyOperation{
		Key:   o.Key,
		Index: o.Index,
	}
}

func (o *EnumKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumKeyOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.Index = op.Index
}
func (o *EnumKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumKeyResponse structure represents the ApiEnumKey operation response
type EnumKeyResponse struct {
	KeyName       string         `idl:"name:KeyName;string" json:"key_name"`
	LastWriteTime *dtyp.Filetime `idl:"name:lpftLastWriteTime" json:"last_write_time"`
	// Return: The ApiEnumKey return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EnumKeyResponse) xxx_ToOp(ctx context.Context) *xxx_EnumKeyOperation {
	if o == nil {
		return &xxx_EnumKeyOperation{}
	}
	return &xxx_EnumKeyOperation{
		KeyName:       o.KeyName,
		LastWriteTime: o.LastWriteTime,
		Return:        o.Return,
	}
}

func (o *EnumKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumKeyOperation) {
	if o == nil {
		return
	}
	o.KeyName = op.KeyName
	o.LastWriteTime = op.LastWriteTime
	o.Return = op.Return
}
func (o *EnumKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetValueOperation structure represents the ApiSetValue operation
type xxx_SetValueOperation struct {
	Key        *Key   `idl:"name:hKey" json:"key"`
	ValueName  string `idl:"name:lpValueName;string" json:"value_name"`
	Type       uint32 `idl:"name:dwType" json:"type"`
	Data       []byte `idl:"name:lpData;size_is:(cbData)" json:"data"`
	DataLength uint32 `idl:"name:cbData" json:"data_length"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SetValueOperation) OpNum() int { return 32 }

func (o *xxx_SetValueOperation) OpName() string { return "/clusapi2/v2/ApiSetValue" }

func (o *xxx_SetValueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Data != nil && o.DataLength == 0 {
		o.DataLength = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetValueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpValueName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ValueName); err != nil {
			return err
		}
	}
	// dwType {in} (1:(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// lpData {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbData](uchar))
	{
		dimSize1 := uint64(o.DataLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Data {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Data[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// cbData {in} (1:(uint32))
	{
		if err := w.WriteData(o.DataLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetValueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpValueName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ValueName); err != nil {
			return err
		}
	}
	// dwType {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// lpData {in} (1:{pointer=ref}*(1)[dim:0,size_is=cbData](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]byte, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
		}
	}
	// cbData {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DataLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetValueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetValueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetValueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetValueRequest structure represents the ApiSetValue operation request
type SetValueRequest struct {
	Key        *Key   `idl:"name:hKey" json:"key"`
	ValueName  string `idl:"name:lpValueName;string" json:"value_name"`
	Type       uint32 `idl:"name:dwType" json:"type"`
	Data       []byte `idl:"name:lpData;size_is:(cbData)" json:"data"`
	DataLength uint32 `idl:"name:cbData" json:"data_length"`
}

func (o *SetValueRequest) xxx_ToOp(ctx context.Context) *xxx_SetValueOperation {
	if o == nil {
		return &xxx_SetValueOperation{}
	}
	return &xxx_SetValueOperation{
		Key:        o.Key,
		ValueName:  o.ValueName,
		Type:       o.Type,
		Data:       o.Data,
		DataLength: o.DataLength,
	}
}

func (o *SetValueRequest) xxx_FromOp(ctx context.Context, op *xxx_SetValueOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.ValueName = op.ValueName
	o.Type = op.Type
	o.Data = op.Data
	o.DataLength = op.DataLength
}
func (o *SetValueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetValueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetValueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetValueResponse structure represents the ApiSetValue operation response
type SetValueResponse struct {
	// Return: The ApiSetValue return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetValueResponse) xxx_ToOp(ctx context.Context) *xxx_SetValueOperation {
	if o == nil {
		return &xxx_SetValueOperation{}
	}
	return &xxx_SetValueOperation{
		Return: o.Return,
	}
}

func (o *SetValueResponse) xxx_FromOp(ctx context.Context, op *xxx_SetValueOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetValueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetValueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetValueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteValueOperation structure represents the ApiDeleteValue operation
type xxx_DeleteValueOperation struct {
	Key       *Key   `idl:"name:hKey" json:"key"`
	ValueName string `idl:"name:lpValueName;string" json:"value_name"`
	Return    uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteValueOperation) OpNum() int { return 33 }

func (o *xxx_DeleteValueOperation) OpName() string { return "/clusapi2/v2/ApiDeleteValue" }

func (o *xxx_DeleteValueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteValueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpValueName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ValueName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteValueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpValueName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ValueName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteValueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteValueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteValueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteValueRequest structure represents the ApiDeleteValue operation request
type DeleteValueRequest struct {
	Key       *Key   `idl:"name:hKey" json:"key"`
	ValueName string `idl:"name:lpValueName;string" json:"value_name"`
}

func (o *DeleteValueRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteValueOperation {
	if o == nil {
		return &xxx_DeleteValueOperation{}
	}
	return &xxx_DeleteValueOperation{
		Key:       o.Key,
		ValueName: o.ValueName,
	}
}

func (o *DeleteValueRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteValueOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.ValueName = op.ValueName
}
func (o *DeleteValueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteValueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteValueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteValueResponse structure represents the ApiDeleteValue operation response
type DeleteValueResponse struct {
	// Return: The ApiDeleteValue return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *DeleteValueResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteValueOperation {
	if o == nil {
		return &xxx_DeleteValueOperation{}
	}
	return &xxx_DeleteValueOperation{
		Return: o.Return,
	}
}

func (o *DeleteValueResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteValueOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DeleteValueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteValueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteValueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryValueOperation structure represents the ApiQueryValue operation
type xxx_QueryValueOperation struct {
	Key            *Key   `idl:"name:hKey" json:"key"`
	ValueName      string `idl:"name:lpValueName;string" json:"value_name"`
	ValueType      uint32 `idl:"name:lpValueType" json:"value_type"`
	Data           []byte `idl:"name:lpData;size_is:(cbData)" json:"data"`
	DataLength     uint32 `idl:"name:cbData" json:"data_length"`
	RequiredLength uint32 `idl:"name:lpcbRequired" json:"required_length"`
	Return         uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryValueOperation) OpNum() int { return 34 }

func (o *xxx_QueryValueOperation) OpName() string { return "/clusapi2/v2/ApiQueryValue" }

func (o *xxx_QueryValueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryValueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpValueName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ValueName); err != nil {
			return err
		}
	}
	// cbData {in} (1:(uint32))
	{
		if err := w.WriteData(o.DataLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryValueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpValueName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ValueName); err != nil {
			return err
		}
	}
	// cbData {in} (1:(uint32))
	{
		if err := w.ReadData(&o.DataLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryValueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryValueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpValueType {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.ValueType); err != nil {
			return err
		}
	}
	// lpData {out} (1:{pointer=ref}*(1)[dim:0,size_is=cbData](uchar))
	{
		dimSize1 := uint64(o.DataLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Data {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Data[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryValueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpValueType {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.ValueType); err != nil {
			return err
		}
	}
	// lpData {out} (1:{pointer=ref}*(1)[dim:0,size_is=cbData](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]byte, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryValueRequest structure represents the ApiQueryValue operation request
type QueryValueRequest struct {
	Key        *Key   `idl:"name:hKey" json:"key"`
	ValueName  string `idl:"name:lpValueName;string" json:"value_name"`
	DataLength uint32 `idl:"name:cbData" json:"data_length"`
}

func (o *QueryValueRequest) xxx_ToOp(ctx context.Context) *xxx_QueryValueOperation {
	if o == nil {
		return &xxx_QueryValueOperation{}
	}
	return &xxx_QueryValueOperation{
		Key:        o.Key,
		ValueName:  o.ValueName,
		DataLength: o.DataLength,
	}
}

func (o *QueryValueRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryValueOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.ValueName = op.ValueName
	o.DataLength = op.DataLength
}
func (o *QueryValueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryValueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryValueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryValueResponse structure represents the ApiQueryValue operation response
type QueryValueResponse struct {
	ValueType      uint32 `idl:"name:lpValueType" json:"value_type"`
	Data           []byte `idl:"name:lpData;size_is:(cbData)" json:"data"`
	RequiredLength uint32 `idl:"name:lpcbRequired" json:"required_length"`
	// Return: The ApiQueryValue return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *QueryValueResponse) xxx_ToOp(ctx context.Context) *xxx_QueryValueOperation {
	if o == nil {
		return &xxx_QueryValueOperation{}
	}
	return &xxx_QueryValueOperation{
		ValueType:      o.ValueType,
		Data:           o.Data,
		RequiredLength: o.RequiredLength,
		Return:         o.Return,
	}
}

func (o *QueryValueResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryValueOperation) {
	if o == nil {
		return
	}
	o.ValueType = op.ValueType
	o.Data = op.Data
	o.RequiredLength = op.RequiredLength
	o.Return = op.Return
}
func (o *QueryValueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryValueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryValueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteKeyOperation structure represents the ApiDeleteKey operation
type xxx_DeleteKeyOperation struct {
	Key    *Key   `idl:"name:hKey" json:"key"`
	SubKey string `idl:"name:lpSubKey;string" json:"sub_key"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteKeyOperation) OpNum() int { return 35 }

func (o *xxx_DeleteKeyOperation) OpName() string { return "/clusapi2/v2/ApiDeleteKey" }

func (o *xxx_DeleteKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpSubKey {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.SubKey); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpSubKey {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.SubKey); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteKeyRequest structure represents the ApiDeleteKey operation request
type DeleteKeyRequest struct {
	Key    *Key   `idl:"name:hKey" json:"key"`
	SubKey string `idl:"name:lpSubKey;string" json:"sub_key"`
}

func (o *DeleteKeyRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteKeyOperation {
	if o == nil {
		return &xxx_DeleteKeyOperation{}
	}
	return &xxx_DeleteKeyOperation{
		Key:    o.Key,
		SubKey: o.SubKey,
	}
}

func (o *DeleteKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteKeyOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.SubKey = op.SubKey
}
func (o *DeleteKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteKeyResponse structure represents the ApiDeleteKey operation response
type DeleteKeyResponse struct {
	// Return: The ApiDeleteKey return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *DeleteKeyResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteKeyOperation {
	if o == nil {
		return &xxx_DeleteKeyOperation{}
	}
	return &xxx_DeleteKeyOperation{
		Return: o.Return,
	}
}

func (o *DeleteKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteKeyOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DeleteKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumValueOperation structure represents the ApiEnumValue operation
type xxx_EnumValueOperation struct {
	Key        *Key   `idl:"name:hKey" json:"key"`
	Index      uint32 `idl:"name:dwIndex" json:"index"`
	ValueName  string `idl:"name:lpValueName;string" json:"value_name"`
	Type       uint32 `idl:"name:lpType" json:"type"`
	Data       []byte `idl:"name:lpData;size_is:(lpcbData)" json:"data"`
	DataLength uint32 `idl:"name:lpcbData" json:"data_length"`
	TotalSize  uint32 `idl:"name:TotalSize" json:"total_size"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumValueOperation) OpNum() int { return 36 }

func (o *xxx_EnumValueOperation) OpName() string { return "/clusapi2/v2/ApiEnumValue" }

func (o *xxx_EnumValueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumValueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwIndex {in} (1:(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	// lpcbData {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.DataLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumValueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwIndex {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	// lpcbData {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.DataLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumValueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Data != nil && o.DataLength == 0 {
		o.DataLength = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumValueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpValueName {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		if o.ValueName != "" {
			_ptr_lpValueName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ValueName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ValueName, _ptr_lpValueName); err != nil {
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
	// lpType {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// lpData {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpcbData](uchar))
	{
		dimSize1 := uint64(o.DataLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Data {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Data[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpcbData {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.DataLength); err != nil {
			return err
		}
	}
	// TotalSize {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.TotalSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumValueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpValueName {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpValueName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ValueName); err != nil {
				return err
			}
			return nil
		})
		_s_lpValueName := func(ptr interface{}) { o.ValueName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ValueName, _s_lpValueName, _ptr_lpValueName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpType {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// lpData {out} (1:{pointer=ref}*(1)[dim:0,size_is=lpcbData](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]byte, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
		}
	}
	// lpcbData {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.DataLength); err != nil {
			return err
		}
	}
	// TotalSize {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.TotalSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumValueRequest structure represents the ApiEnumValue operation request
type EnumValueRequest struct {
	Key        *Key   `idl:"name:hKey" json:"key"`
	Index      uint32 `idl:"name:dwIndex" json:"index"`
	DataLength uint32 `idl:"name:lpcbData" json:"data_length"`
}

func (o *EnumValueRequest) xxx_ToOp(ctx context.Context) *xxx_EnumValueOperation {
	if o == nil {
		return &xxx_EnumValueOperation{}
	}
	return &xxx_EnumValueOperation{
		Key:        o.Key,
		Index:      o.Index,
		DataLength: o.DataLength,
	}
}

func (o *EnumValueRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumValueOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.Index = op.Index
	o.DataLength = op.DataLength
}
func (o *EnumValueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EnumValueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumValueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumValueResponse structure represents the ApiEnumValue operation response
type EnumValueResponse struct {
	ValueName  string `idl:"name:lpValueName;string" json:"value_name"`
	Type       uint32 `idl:"name:lpType" json:"type"`
	Data       []byte `idl:"name:lpData;size_is:(lpcbData)" json:"data"`
	DataLength uint32 `idl:"name:lpcbData" json:"data_length"`
	TotalSize  uint32 `idl:"name:TotalSize" json:"total_size"`
	// Return: The ApiEnumValue return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EnumValueResponse) xxx_ToOp(ctx context.Context) *xxx_EnumValueOperation {
	if o == nil {
		return &xxx_EnumValueOperation{}
	}
	return &xxx_EnumValueOperation{
		ValueName:  o.ValueName,
		Type:       o.Type,
		Data:       o.Data,
		DataLength: o.DataLength,
		TotalSize:  o.TotalSize,
		Return:     o.Return,
	}
}

func (o *EnumValueResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumValueOperation) {
	if o == nil {
		return
	}
	o.ValueName = op.ValueName
	o.Type = op.Type
	o.Data = op.Data
	o.DataLength = op.DataLength
	o.TotalSize = op.TotalSize
	o.Return = op.Return
}
func (o *EnumValueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EnumValueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumValueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseKeyOperation structure represents the ApiCloseKey operation
type xxx_CloseKeyOperation struct {
	Key    *Key   `idl:"name:pKey" json:"key"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseKeyOperation) OpNum() int { return 37 }

func (o *xxx_CloseKeyOperation) OpName() string { return "/clusapi2/v2/ApiCloseKey" }

func (o *xxx_CloseKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pKey {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pKey {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pKey {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pKey {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloseKeyRequest structure represents the ApiCloseKey operation request
type CloseKeyRequest struct {
	Key *Key `idl:"name:pKey" json:"key"`
}

func (o *CloseKeyRequest) xxx_ToOp(ctx context.Context) *xxx_CloseKeyOperation {
	if o == nil {
		return &xxx_CloseKeyOperation{}
	}
	return &xxx_CloseKeyOperation{
		Key: o.Key,
	}
}

func (o *CloseKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseKeyOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
}
func (o *CloseKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CloseKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseKeyResponse structure represents the ApiCloseKey operation response
type CloseKeyResponse struct {
	Key *Key `idl:"name:pKey" json:"key"`
	// Return: The ApiCloseKey return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CloseKeyResponse) xxx_ToOp(ctx context.Context) *xxx_CloseKeyOperation {
	if o == nil {
		return &xxx_CloseKeyOperation{}
	}
	return &xxx_CloseKeyOperation{
		Key:    o.Key,
		Return: o.Return,
	}
}

func (o *CloseKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseKeyOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.Return = op.Return
}
func (o *CloseKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CloseKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryInfoKeyOperation structure represents the ApiQueryInfoKey operation
type xxx_QueryInfoKeyOperation struct {
	Key                      *Key           `idl:"name:hKey" json:"key"`
	SubKeysCount             uint32         `idl:"name:lpcSubKeys" json:"sub_keys_count"`
	MaxSubKeyLength          uint32         `idl:"name:lpcbMaxSubKeyLen" json:"max_sub_key_length"`
	ValuesCount              uint32         `idl:"name:lpcValues" json:"values_count"`
	MaxValueNameLength       uint32         `idl:"name:lpcbMaxValueNameLen" json:"max_value_name_length"`
	MaxValueLength           uint32         `idl:"name:lpcbMaxValueLen" json:"max_value_length"`
	SecurityDescriptorLength uint32         `idl:"name:lpcbSecurityDescriptor" json:"security_descriptor_length"`
	LastWriteTime            *dtyp.Filetime `idl:"name:lpftLastWriteTime" json:"last_write_time"`
	Return                   uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryInfoKeyOperation) OpNum() int { return 38 }

func (o *xxx_QueryInfoKeyOperation) OpName() string { return "/clusapi2/v2/ApiQueryInfoKey" }

func (o *xxx_QueryInfoKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInfoKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_QueryInfoKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInfoKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInfoKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpcSubKeys {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SubKeysCount); err != nil {
			return err
		}
	}
	// lpcbMaxSubKeyLen {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MaxSubKeyLength); err != nil {
			return err
		}
	}
	// lpcValues {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ValuesCount); err != nil {
			return err
		}
	}
	// lpcbMaxValueNameLen {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MaxValueNameLength); err != nil {
			return err
		}
	}
	// lpcbMaxValueLen {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MaxValueLength); err != nil {
			return err
		}
	}
	// lpcbSecurityDescriptor {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SecurityDescriptorLength); err != nil {
			return err
		}
	}
	// lpftLastWriteTime {out} (1:{alias=PFILETIME}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.LastWriteTime != nil {
			if err := o.LastWriteTime.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryInfoKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpcSubKeys {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SubKeysCount); err != nil {
			return err
		}
	}
	// lpcbMaxSubKeyLen {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MaxSubKeyLength); err != nil {
			return err
		}
	}
	// lpcValues {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ValuesCount); err != nil {
			return err
		}
	}
	// lpcbMaxValueNameLen {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MaxValueNameLength); err != nil {
			return err
		}
	}
	// lpcbMaxValueLen {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MaxValueLength); err != nil {
			return err
		}
	}
	// lpcbSecurityDescriptor {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SecurityDescriptorLength); err != nil {
			return err
		}
	}
	// lpftLastWriteTime {out} (1:{alias=PFILETIME,pointer=ref}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.LastWriteTime == nil {
			o.LastWriteTime = &dtyp.Filetime{}
		}
		if err := o.LastWriteTime.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryInfoKeyRequest structure represents the ApiQueryInfoKey operation request
type QueryInfoKeyRequest struct {
	Key *Key `idl:"name:hKey" json:"key"`
}

func (o *QueryInfoKeyRequest) xxx_ToOp(ctx context.Context) *xxx_QueryInfoKeyOperation {
	if o == nil {
		return &xxx_QueryInfoKeyOperation{}
	}
	return &xxx_QueryInfoKeyOperation{
		Key: o.Key,
	}
}

func (o *QueryInfoKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryInfoKeyOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
}
func (o *QueryInfoKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *QueryInfoKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInfoKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryInfoKeyResponse structure represents the ApiQueryInfoKey operation response
type QueryInfoKeyResponse struct {
	SubKeysCount             uint32         `idl:"name:lpcSubKeys" json:"sub_keys_count"`
	MaxSubKeyLength          uint32         `idl:"name:lpcbMaxSubKeyLen" json:"max_sub_key_length"`
	ValuesCount              uint32         `idl:"name:lpcValues" json:"values_count"`
	MaxValueNameLength       uint32         `idl:"name:lpcbMaxValueNameLen" json:"max_value_name_length"`
	MaxValueLength           uint32         `idl:"name:lpcbMaxValueLen" json:"max_value_length"`
	SecurityDescriptorLength uint32         `idl:"name:lpcbSecurityDescriptor" json:"security_descriptor_length"`
	LastWriteTime            *dtyp.Filetime `idl:"name:lpftLastWriteTime" json:"last_write_time"`
	// Return: The ApiQueryInfoKey return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *QueryInfoKeyResponse) xxx_ToOp(ctx context.Context) *xxx_QueryInfoKeyOperation {
	if o == nil {
		return &xxx_QueryInfoKeyOperation{}
	}
	return &xxx_QueryInfoKeyOperation{
		SubKeysCount:             o.SubKeysCount,
		MaxSubKeyLength:          o.MaxSubKeyLength,
		ValuesCount:              o.ValuesCount,
		MaxValueNameLength:       o.MaxValueNameLength,
		MaxValueLength:           o.MaxValueLength,
		SecurityDescriptorLength: o.SecurityDescriptorLength,
		LastWriteTime:            o.LastWriteTime,
		Return:                   o.Return,
	}
}

func (o *QueryInfoKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryInfoKeyOperation) {
	if o == nil {
		return
	}
	o.SubKeysCount = op.SubKeysCount
	o.MaxSubKeyLength = op.MaxSubKeyLength
	o.ValuesCount = op.ValuesCount
	o.MaxValueNameLength = op.MaxValueNameLength
	o.MaxValueLength = op.MaxValueLength
	o.SecurityDescriptorLength = op.SecurityDescriptorLength
	o.LastWriteTime = op.LastWriteTime
	o.Return = op.Return
}
func (o *QueryInfoKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *QueryInfoKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryInfoKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetKeySecurityOperation structure represents the ApiSetKeySecurity operation
type xxx_SetKeySecurityOperation struct {
	Key                 *Key                `idl:"name:hKey" json:"key"`
	SecurityInformation uint32              `idl:"name:SecurityInformation" json:"security_information"`
	SecurityDescriptor  *SecurityDescriptor `idl:"name:pRpcSecurityDescriptor" json:"security_descriptor"`
	Return              uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_SetKeySecurityOperation) OpNum() int { return 39 }

func (o *xxx_SetKeySecurityOperation) OpName() string { return "/clusapi2/v2/ApiSetKeySecurity" }

func (o *xxx_SetKeySecurityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetKeySecurityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// SecurityInformation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SecurityInformation); err != nil {
			return err
		}
	}
	// pRpcSecurityDescriptor {in} (1:{alias=PRPC_SECURITY_DESCRIPTOR}*(1))(2:{alias=RPC_SECURITY_DESCRIPTOR}(struct))
	{
		if o.SecurityDescriptor != nil {
			if err := o.SecurityDescriptor.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SecurityDescriptor{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetKeySecurityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// SecurityInformation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SecurityInformation); err != nil {
			return err
		}
	}
	// pRpcSecurityDescriptor {in} (1:{alias=PRPC_SECURITY_DESCRIPTOR,pointer=ref}*(1))(2:{alias=RPC_SECURITY_DESCRIPTOR}(struct))
	{
		if o.SecurityDescriptor == nil {
			o.SecurityDescriptor = &SecurityDescriptor{}
		}
		if err := o.SecurityDescriptor.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetKeySecurityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetKeySecurityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetKeySecurityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetKeySecurityRequest structure represents the ApiSetKeySecurity operation request
type SetKeySecurityRequest struct {
	Key                 *Key                `idl:"name:hKey" json:"key"`
	SecurityInformation uint32              `idl:"name:SecurityInformation" json:"security_information"`
	SecurityDescriptor  *SecurityDescriptor `idl:"name:pRpcSecurityDescriptor" json:"security_descriptor"`
}

func (o *SetKeySecurityRequest) xxx_ToOp(ctx context.Context) *xxx_SetKeySecurityOperation {
	if o == nil {
		return &xxx_SetKeySecurityOperation{}
	}
	return &xxx_SetKeySecurityOperation{
		Key:                 o.Key,
		SecurityInformation: o.SecurityInformation,
		SecurityDescriptor:  o.SecurityDescriptor,
	}
}

func (o *SetKeySecurityRequest) xxx_FromOp(ctx context.Context, op *xxx_SetKeySecurityOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.SecurityInformation = op.SecurityInformation
	o.SecurityDescriptor = op.SecurityDescriptor
}
func (o *SetKeySecurityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetKeySecurityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetKeySecurityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetKeySecurityResponse structure represents the ApiSetKeySecurity operation response
type SetKeySecurityResponse struct {
	// Return: The ApiSetKeySecurity return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetKeySecurityResponse) xxx_ToOp(ctx context.Context) *xxx_SetKeySecurityOperation {
	if o == nil {
		return &xxx_SetKeySecurityOperation{}
	}
	return &xxx_SetKeySecurityOperation{
		Return: o.Return,
	}
}

func (o *SetKeySecurityResponse) xxx_FromOp(ctx context.Context, op *xxx_SetKeySecurityOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetKeySecurityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetKeySecurityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetKeySecurityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetKeySecurityOperation structure represents the ApiGetKeySecurity operation
type xxx_GetKeySecurityOperation struct {
	Key                 *Key                `idl:"name:hKey" json:"key"`
	SecurityInformation uint32              `idl:"name:SecurityInformation" json:"security_information"`
	SecurityDescriptor  *SecurityDescriptor `idl:"name:pRpcSecurityDescriptor" json:"security_descriptor"`
	Return              uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_GetKeySecurityOperation) OpNum() int { return 40 }

func (o *xxx_GetKeySecurityOperation) OpName() string { return "/clusapi2/v2/ApiGetKeySecurity" }

func (o *xxx_GetKeySecurityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetKeySecurityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// SecurityInformation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SecurityInformation); err != nil {
			return err
		}
	}
	// pRpcSecurityDescriptor {in, out} (1:{alias=PRPC_SECURITY_DESCRIPTOR}*(1))(2:{alias=RPC_SECURITY_DESCRIPTOR}(struct))
	{
		if o.SecurityDescriptor != nil {
			if err := o.SecurityDescriptor.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SecurityDescriptor{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetKeySecurityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// SecurityInformation {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SecurityInformation); err != nil {
			return err
		}
	}
	// pRpcSecurityDescriptor {in, out} (1:{alias=PRPC_SECURITY_DESCRIPTOR,pointer=ref}*(1))(2:{alias=RPC_SECURITY_DESCRIPTOR}(struct))
	{
		if o.SecurityDescriptor == nil {
			o.SecurityDescriptor = &SecurityDescriptor{}
		}
		if err := o.SecurityDescriptor.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetKeySecurityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetKeySecurityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pRpcSecurityDescriptor {in, out} (1:{alias=PRPC_SECURITY_DESCRIPTOR}*(1))(2:{alias=RPC_SECURITY_DESCRIPTOR}(struct))
	{
		if o.SecurityDescriptor != nil {
			if err := o.SecurityDescriptor.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SecurityDescriptor{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetKeySecurityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pRpcSecurityDescriptor {in, out} (1:{alias=PRPC_SECURITY_DESCRIPTOR,pointer=ref}*(1))(2:{alias=RPC_SECURITY_DESCRIPTOR}(struct))
	{
		if o.SecurityDescriptor == nil {
			o.SecurityDescriptor = &SecurityDescriptor{}
		}
		if err := o.SecurityDescriptor.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetKeySecurityRequest structure represents the ApiGetKeySecurity operation request
type GetKeySecurityRequest struct {
	Key                 *Key                `idl:"name:hKey" json:"key"`
	SecurityInformation uint32              `idl:"name:SecurityInformation" json:"security_information"`
	SecurityDescriptor  *SecurityDescriptor `idl:"name:pRpcSecurityDescriptor" json:"security_descriptor"`
}

func (o *GetKeySecurityRequest) xxx_ToOp(ctx context.Context) *xxx_GetKeySecurityOperation {
	if o == nil {
		return &xxx_GetKeySecurityOperation{}
	}
	return &xxx_GetKeySecurityOperation{
		Key:                 o.Key,
		SecurityInformation: o.SecurityInformation,
		SecurityDescriptor:  o.SecurityDescriptor,
	}
}

func (o *GetKeySecurityRequest) xxx_FromOp(ctx context.Context, op *xxx_GetKeySecurityOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.SecurityInformation = op.SecurityInformation
	o.SecurityDescriptor = op.SecurityDescriptor
}
func (o *GetKeySecurityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetKeySecurityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetKeySecurityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetKeySecurityResponse structure represents the ApiGetKeySecurity operation response
type GetKeySecurityResponse struct {
	SecurityDescriptor *SecurityDescriptor `idl:"name:pRpcSecurityDescriptor" json:"security_descriptor"`
	// Return: The ApiGetKeySecurity return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetKeySecurityResponse) xxx_ToOp(ctx context.Context) *xxx_GetKeySecurityOperation {
	if o == nil {
		return &xxx_GetKeySecurityOperation{}
	}
	return &xxx_GetKeySecurityOperation{
		SecurityDescriptor: o.SecurityDescriptor,
		Return:             o.Return,
	}
}

func (o *GetKeySecurityResponse) xxx_FromOp(ctx context.Context, op *xxx_GetKeySecurityOperation) {
	if o == nil {
		return
	}
	o.SecurityDescriptor = op.SecurityDescriptor
	o.Return = op.Return
}
func (o *GetKeySecurityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetKeySecurityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetKeySecurityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenGroupOperation structure represents the ApiOpenGroup operation
type xxx_OpenGroupOperation struct {
	GroupName string `idl:"name:lpszGroupName;string" json:"group_name"`
	Status    uint32 `idl:"name:Status" json:"status"`
	Return    *Group `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenGroupOperation) OpNum() int { return 41 }

func (o *xxx_OpenGroupOperation) OpName() string { return "/clusapi2/v2/ApiOpenGroup" }

func (o *xxx_OpenGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpszGroupName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.GroupName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpszGroupName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.GroupName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return != nil {
			if err := o.Return.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Group{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_OpenGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return == nil {
			o.Return = &Group{}
		}
		if err := o.Return.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// OpenGroupRequest structure represents the ApiOpenGroup operation request
type OpenGroupRequest struct {
	GroupName string `idl:"name:lpszGroupName;string" json:"group_name"`
}

func (o *OpenGroupRequest) xxx_ToOp(ctx context.Context) *xxx_OpenGroupOperation {
	if o == nil {
		return &xxx_OpenGroupOperation{}
	}
	return &xxx_OpenGroupOperation{
		GroupName: o.GroupName,
	}
}

func (o *OpenGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenGroupOperation) {
	if o == nil {
		return
	}
	o.GroupName = op.GroupName
}
func (o *OpenGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OpenGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenGroupResponse structure represents the ApiOpenGroup operation response
type OpenGroupResponse struct {
	Status uint32 `idl:"name:Status" json:"status"`
	// Return: The ApiOpenGroup return value.
	Return *Group `idl:"name:Return" json:"return"`
}

func (o *OpenGroupResponse) xxx_ToOp(ctx context.Context) *xxx_OpenGroupOperation {
	if o == nil {
		return &xxx_OpenGroupOperation{}
	}
	return &xxx_OpenGroupOperation{
		Status: o.Status,
		Return: o.Return,
	}
}

func (o *OpenGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenGroupOperation) {
	if o == nil {
		return
	}
	o.Status = op.Status
	o.Return = op.Return
}
func (o *OpenGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OpenGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateGroupOperation structure represents the ApiCreateGroup operation
type xxx_CreateGroupOperation struct {
	GroupName string `idl:"name:lpszGroupName;string" json:"group_name"`
	Status    uint32 `idl:"name:Status" json:"status"`
	Return    *Group `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateGroupOperation) OpNum() int { return 42 }

func (o *xxx_CreateGroupOperation) OpName() string { return "/clusapi2/v2/ApiCreateGroup" }

func (o *xxx_CreateGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpszGroupName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.GroupName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpszGroupName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.GroupName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return != nil {
			if err := o.Return.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Group{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CreateGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return == nil {
			o.Return = &Group{}
		}
		if err := o.Return.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// CreateGroupRequest structure represents the ApiCreateGroup operation request
type CreateGroupRequest struct {
	GroupName string `idl:"name:lpszGroupName;string" json:"group_name"`
}

func (o *CreateGroupRequest) xxx_ToOp(ctx context.Context) *xxx_CreateGroupOperation {
	if o == nil {
		return &xxx_CreateGroupOperation{}
	}
	return &xxx_CreateGroupOperation{
		GroupName: o.GroupName,
	}
}

func (o *CreateGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateGroupOperation) {
	if o == nil {
		return
	}
	o.GroupName = op.GroupName
}
func (o *CreateGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateGroupResponse structure represents the ApiCreateGroup operation response
type CreateGroupResponse struct {
	Status uint32 `idl:"name:Status" json:"status"`
	// Return: The ApiCreateGroup return value.
	Return *Group `idl:"name:Return" json:"return"`
}

func (o *CreateGroupResponse) xxx_ToOp(ctx context.Context) *xxx_CreateGroupOperation {
	if o == nil {
		return &xxx_CreateGroupOperation{}
	}
	return &xxx_CreateGroupOperation{
		Status: o.Status,
		Return: o.Return,
	}
}

func (o *CreateGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateGroupOperation) {
	if o == nil {
		return
	}
	o.Status = op.Status
	o.Return = op.Return
}
func (o *CreateGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteGroupOperation structure represents the ApiDeleteGroup operation
type xxx_DeleteGroupOperation struct {
	Group  *Group `idl:"name:Group" json:"group"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteGroupOperation) OpNum() int { return 43 }

func (o *xxx_DeleteGroupOperation) OpName() string { return "/clusapi2/v2/ApiDeleteGroup" }

func (o *xxx_DeleteGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Group {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group != nil {
			if err := o.Group.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Group{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DeleteGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Group {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group == nil {
			o.Group = &Group{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteGroupRequest structure represents the ApiDeleteGroup operation request
type DeleteGroupRequest struct {
	Group *Group `idl:"name:Group" json:"group"`
}

func (o *DeleteGroupRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteGroupOperation {
	if o == nil {
		return &xxx_DeleteGroupOperation{}
	}
	return &xxx_DeleteGroupOperation{
		Group: o.Group,
	}
}

func (o *DeleteGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteGroupOperation) {
	if o == nil {
		return
	}
	o.Group = op.Group
}
func (o *DeleteGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteGroupResponse structure represents the ApiDeleteGroup operation response
type DeleteGroupResponse struct {
	// Return: The ApiDeleteGroup return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *DeleteGroupResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteGroupOperation {
	if o == nil {
		return &xxx_DeleteGroupOperation{}
	}
	return &xxx_DeleteGroupOperation{
		Return: o.Return,
	}
}

func (o *DeleteGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteGroupOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DeleteGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseGroupOperation structure represents the ApiCloseGroup operation
type xxx_CloseGroupOperation struct {
	Group  *Group `idl:"name:Group" json:"group"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseGroupOperation) OpNum() int { return 44 }

func (o *xxx_CloseGroupOperation) OpName() string { return "/clusapi2/v2/ApiCloseGroup" }

func (o *xxx_CloseGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Group {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group != nil {
			if err := o.Group.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Group{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Group {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group == nil {
			o.Group = &Group{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Group {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group != nil {
			if err := o.Group.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Group{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Group {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group == nil {
			o.Group = &Group{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloseGroupRequest structure represents the ApiCloseGroup operation request
type CloseGroupRequest struct {
	Group *Group `idl:"name:Group" json:"group"`
}

func (o *CloseGroupRequest) xxx_ToOp(ctx context.Context) *xxx_CloseGroupOperation {
	if o == nil {
		return &xxx_CloseGroupOperation{}
	}
	return &xxx_CloseGroupOperation{
		Group: o.Group,
	}
}

func (o *CloseGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseGroupOperation) {
	if o == nil {
		return
	}
	o.Group = op.Group
}
func (o *CloseGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CloseGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseGroupResponse structure represents the ApiCloseGroup operation response
type CloseGroupResponse struct {
	Group *Group `idl:"name:Group" json:"group"`
	// Return: The ApiCloseGroup return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CloseGroupResponse) xxx_ToOp(ctx context.Context) *xxx_CloseGroupOperation {
	if o == nil {
		return &xxx_CloseGroupOperation{}
	}
	return &xxx_CloseGroupOperation{
		Group:  o.Group,
		Return: o.Return,
	}
}

func (o *CloseGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseGroupOperation) {
	if o == nil {
		return
	}
	o.Group = op.Group
	o.Return = op.Return
}
func (o *CloseGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CloseGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetGroupStateOperation structure represents the ApiGetGroupState operation
type xxx_GetGroupStateOperation struct {
	Group    *Group `idl:"name:hGroup" json:"group"`
	State    uint32 `idl:"name:State" json:"state"`
	NodeName string `idl:"name:NodeName;string" json:"node_name"`
	Return   uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetGroupStateOperation) OpNum() int { return 45 }

func (o *xxx_GetGroupStateOperation) OpName() string { return "/clusapi2/v2/ApiGetGroupState" }

func (o *xxx_GetGroupStateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGroupStateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group != nil {
			if err := o.Group.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Group{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetGroupStateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group == nil {
			o.Group = &Group{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGroupStateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGroupStateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// State {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.State); err != nil {
			return err
		}
	}
	// NodeName {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		if o.NodeName != "" {
			_ptr_NodeName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.NodeName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.NodeName, _ptr_NodeName); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGroupStateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// State {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.State); err != nil {
			return err
		}
	}
	// NodeName {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		_ptr_NodeName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.NodeName); err != nil {
				return err
			}
			return nil
		})
		_s_NodeName := func(ptr interface{}) { o.NodeName = *ptr.(*string) }
		if err := w.ReadPointer(&o.NodeName, _s_NodeName, _ptr_NodeName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetGroupStateRequest structure represents the ApiGetGroupState operation request
type GetGroupStateRequest struct {
	Group *Group `idl:"name:hGroup" json:"group"`
}

func (o *GetGroupStateRequest) xxx_ToOp(ctx context.Context) *xxx_GetGroupStateOperation {
	if o == nil {
		return &xxx_GetGroupStateOperation{}
	}
	return &xxx_GetGroupStateOperation{
		Group: o.Group,
	}
}

func (o *GetGroupStateRequest) xxx_FromOp(ctx context.Context, op *xxx_GetGroupStateOperation) {
	if o == nil {
		return
	}
	o.Group = op.Group
}
func (o *GetGroupStateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetGroupStateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetGroupStateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetGroupStateResponse structure represents the ApiGetGroupState operation response
type GetGroupStateResponse struct {
	State    uint32 `idl:"name:State" json:"state"`
	NodeName string `idl:"name:NodeName;string" json:"node_name"`
	// Return: The ApiGetGroupState return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetGroupStateResponse) xxx_ToOp(ctx context.Context) *xxx_GetGroupStateOperation {
	if o == nil {
		return &xxx_GetGroupStateOperation{}
	}
	return &xxx_GetGroupStateOperation{
		State:    o.State,
		NodeName: o.NodeName,
		Return:   o.Return,
	}
}

func (o *GetGroupStateResponse) xxx_FromOp(ctx context.Context, op *xxx_GetGroupStateOperation) {
	if o == nil {
		return
	}
	o.State = op.State
	o.NodeName = op.NodeName
	o.Return = op.Return
}
func (o *GetGroupStateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetGroupStateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetGroupStateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetGroupNameOperation structure represents the ApiSetGroupName operation
type xxx_SetGroupNameOperation struct {
	Group     *Group `idl:"name:hGroup" json:"group"`
	GroupName string `idl:"name:lpszGroupName;string" json:"group_name"`
	Return    uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SetGroupNameOperation) OpNum() int { return 46 }

func (o *xxx_SetGroupNameOperation) OpName() string { return "/clusapi2/v2/ApiSetGroupName" }

func (o *xxx_SetGroupNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGroupNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group != nil {
			if err := o.Group.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Group{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpszGroupName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.GroupName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGroupNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group == nil {
			o.Group = &Group{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpszGroupName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.GroupName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGroupNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGroupNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGroupNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetGroupNameRequest structure represents the ApiSetGroupName operation request
type SetGroupNameRequest struct {
	Group     *Group `idl:"name:hGroup" json:"group"`
	GroupName string `idl:"name:lpszGroupName;string" json:"group_name"`
}

func (o *SetGroupNameRequest) xxx_ToOp(ctx context.Context) *xxx_SetGroupNameOperation {
	if o == nil {
		return &xxx_SetGroupNameOperation{}
	}
	return &xxx_SetGroupNameOperation{
		Group:     o.Group,
		GroupName: o.GroupName,
	}
}

func (o *SetGroupNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetGroupNameOperation) {
	if o == nil {
		return
	}
	o.Group = op.Group
	o.GroupName = op.GroupName
}
func (o *SetGroupNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetGroupNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetGroupNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetGroupNameResponse structure represents the ApiSetGroupName operation response
type SetGroupNameResponse struct {
	// Return: The ApiSetGroupName return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetGroupNameResponse) xxx_ToOp(ctx context.Context) *xxx_SetGroupNameOperation {
	if o == nil {
		return &xxx_SetGroupNameOperation{}
	}
	return &xxx_SetGroupNameOperation{
		Return: o.Return,
	}
}

func (o *SetGroupNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetGroupNameOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetGroupNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetGroupNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetGroupNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetGroupIDOperation structure represents the ApiGetGroupId operation
type xxx_GetGroupIDOperation struct {
	Group  *Group `idl:"name:hGroup" json:"group"`
	GUID   string `idl:"name:pGuid;string" json:"guid"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetGroupIDOperation) OpNum() int { return 47 }

func (o *xxx_GetGroupIDOperation) OpName() string { return "/clusapi2/v2/ApiGetGroupId" }

func (o *xxx_GetGroupIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGroupIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group != nil {
			if err := o.Group.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Group{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetGroupIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group == nil {
			o.Group = &Group{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGroupIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGroupIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pGuid {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.GUID != "" {
			_ptr_pGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.GUID); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.GUID, _ptr_pGuid); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGroupIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pGuid {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.GUID); err != nil {
				return err
			}
			return nil
		})
		_s_pGuid := func(ptr interface{}) { o.GUID = *ptr.(*string) }
		if err := w.ReadPointer(&o.GUID, _s_pGuid, _ptr_pGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetGroupIDRequest structure represents the ApiGetGroupId operation request
type GetGroupIDRequest struct {
	Group *Group `idl:"name:hGroup" json:"group"`
}

func (o *GetGroupIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetGroupIDOperation {
	if o == nil {
		return &xxx_GetGroupIDOperation{}
	}
	return &xxx_GetGroupIDOperation{
		Group: o.Group,
	}
}

func (o *GetGroupIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetGroupIDOperation) {
	if o == nil {
		return
	}
	o.Group = op.Group
}
func (o *GetGroupIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetGroupIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetGroupIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetGroupIDResponse structure represents the ApiGetGroupId operation response
type GetGroupIDResponse struct {
	GUID string `idl:"name:pGuid;string" json:"guid"`
	// Return: The ApiGetGroupId return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetGroupIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetGroupIDOperation {
	if o == nil {
		return &xxx_GetGroupIDOperation{}
	}
	return &xxx_GetGroupIDOperation{
		GUID:   o.GUID,
		Return: o.Return,
	}
}

func (o *GetGroupIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetGroupIDOperation) {
	if o == nil {
		return
	}
	o.GUID = op.GUID
	o.Return = op.Return
}
func (o *GetGroupIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetGroupIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetGroupIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNodeIDOperation structure represents the ApiGetNodeId operation
type xxx_GetNodeIDOperation struct {
	Node   *Node  `idl:"name:hNode" json:"node"`
	GUID   string `idl:"name:pGuid;string" json:"guid"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNodeIDOperation) OpNum() int { return 48 }

func (o *xxx_GetNodeIDOperation) OpName() string { return "/clusapi2/v2/ApiGetNodeId" }

func (o *xxx_GetNodeIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNodeIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node != nil {
			if err := o.Node.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetNodeIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node == nil {
			o.Node = &Node{}
		}
		if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNodeIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNodeIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pGuid {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.GUID != "" {
			_ptr_pGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.GUID); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.GUID, _ptr_pGuid); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNodeIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pGuid {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.GUID); err != nil {
				return err
			}
			return nil
		})
		_s_pGuid := func(ptr interface{}) { o.GUID = *ptr.(*string) }
		if err := w.ReadPointer(&o.GUID, _s_pGuid, _ptr_pGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNodeIDRequest structure represents the ApiGetNodeId operation request
type GetNodeIDRequest struct {
	Node *Node `idl:"name:hNode" json:"node"`
}

func (o *GetNodeIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetNodeIDOperation {
	if o == nil {
		return &xxx_GetNodeIDOperation{}
	}
	return &xxx_GetNodeIDOperation{
		Node: o.Node,
	}
}

func (o *GetNodeIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNodeIDOperation) {
	if o == nil {
		return
	}
	o.Node = op.Node
}
func (o *GetNodeIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetNodeIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNodeIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNodeIDResponse structure represents the ApiGetNodeId operation response
type GetNodeIDResponse struct {
	GUID string `idl:"name:pGuid;string" json:"guid"`
	// Return: The ApiGetNodeId return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetNodeIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetNodeIDOperation {
	if o == nil {
		return &xxx_GetNodeIDOperation{}
	}
	return &xxx_GetNodeIDOperation{
		GUID:   o.GUID,
		Return: o.Return,
	}
}

func (o *GetNodeIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNodeIDOperation) {
	if o == nil {
		return
	}
	o.GUID = op.GUID
	o.Return = op.Return
}
func (o *GetNodeIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetNodeIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNodeIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OnlineGroupOperation structure represents the ApiOnlineGroup operation
type xxx_OnlineGroupOperation struct {
	Group  *Group `idl:"name:hGroup" json:"group"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_OnlineGroupOperation) OpNum() int { return 49 }

func (o *xxx_OnlineGroupOperation) OpName() string { return "/clusapi2/v2/ApiOnlineGroup" }

func (o *xxx_OnlineGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnlineGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group != nil {
			if err := o.Group.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Group{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_OnlineGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group == nil {
			o.Group = &Group{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnlineGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnlineGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OnlineGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OnlineGroupRequest structure represents the ApiOnlineGroup operation request
type OnlineGroupRequest struct {
	Group *Group `idl:"name:hGroup" json:"group"`
}

func (o *OnlineGroupRequest) xxx_ToOp(ctx context.Context) *xxx_OnlineGroupOperation {
	if o == nil {
		return &xxx_OnlineGroupOperation{}
	}
	return &xxx_OnlineGroupOperation{
		Group: o.Group,
	}
}

func (o *OnlineGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_OnlineGroupOperation) {
	if o == nil {
		return
	}
	o.Group = op.Group
}
func (o *OnlineGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OnlineGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OnlineGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OnlineGroupResponse structure represents the ApiOnlineGroup operation response
type OnlineGroupResponse struct {
	// Return: The ApiOnlineGroup return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *OnlineGroupResponse) xxx_ToOp(ctx context.Context) *xxx_OnlineGroupOperation {
	if o == nil {
		return &xxx_OnlineGroupOperation{}
	}
	return &xxx_OnlineGroupOperation{
		Return: o.Return,
	}
}

func (o *OnlineGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_OnlineGroupOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *OnlineGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OnlineGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OnlineGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OfflineGroupOperation structure represents the ApiOfflineGroup operation
type xxx_OfflineGroupOperation struct {
	Group  *Group `idl:"name:hGroup" json:"group"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_OfflineGroupOperation) OpNum() int { return 50 }

func (o *xxx_OfflineGroupOperation) OpName() string { return "/clusapi2/v2/ApiOfflineGroup" }

func (o *xxx_OfflineGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OfflineGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group != nil {
			if err := o.Group.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Group{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_OfflineGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group == nil {
			o.Group = &Group{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OfflineGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OfflineGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OfflineGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OfflineGroupRequest structure represents the ApiOfflineGroup operation request
type OfflineGroupRequest struct {
	Group *Group `idl:"name:hGroup" json:"group"`
}

func (o *OfflineGroupRequest) xxx_ToOp(ctx context.Context) *xxx_OfflineGroupOperation {
	if o == nil {
		return &xxx_OfflineGroupOperation{}
	}
	return &xxx_OfflineGroupOperation{
		Group: o.Group,
	}
}

func (o *OfflineGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_OfflineGroupOperation) {
	if o == nil {
		return
	}
	o.Group = op.Group
}
func (o *OfflineGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OfflineGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OfflineGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OfflineGroupResponse structure represents the ApiOfflineGroup operation response
type OfflineGroupResponse struct {
	// Return: The ApiOfflineGroup return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *OfflineGroupResponse) xxx_ToOp(ctx context.Context) *xxx_OfflineGroupOperation {
	if o == nil {
		return &xxx_OfflineGroupOperation{}
	}
	return &xxx_OfflineGroupOperation{
		Return: o.Return,
	}
}

func (o *OfflineGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_OfflineGroupOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *OfflineGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OfflineGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OfflineGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MoveGroupOperation structure represents the ApiMoveGroup operation
type xxx_MoveGroupOperation struct {
	Group  *Group `idl:"name:hGroup" json:"group"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_MoveGroupOperation) OpNum() int { return 51 }

func (o *xxx_MoveGroupOperation) OpName() string { return "/clusapi2/v2/ApiMoveGroup" }

func (o *xxx_MoveGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group != nil {
			if err := o.Group.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Group{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_MoveGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group == nil {
			o.Group = &Group{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MoveGroupRequest structure represents the ApiMoveGroup operation request
type MoveGroupRequest struct {
	Group *Group `idl:"name:hGroup" json:"group"`
}

func (o *MoveGroupRequest) xxx_ToOp(ctx context.Context) *xxx_MoveGroupOperation {
	if o == nil {
		return &xxx_MoveGroupOperation{}
	}
	return &xxx_MoveGroupOperation{
		Group: o.Group,
	}
}

func (o *MoveGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_MoveGroupOperation) {
	if o == nil {
		return
	}
	o.Group = op.Group
}
func (o *MoveGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *MoveGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MoveGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MoveGroupResponse structure represents the ApiMoveGroup operation response
type MoveGroupResponse struct {
	// Return: The ApiMoveGroup return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MoveGroupResponse) xxx_ToOp(ctx context.Context) *xxx_MoveGroupOperation {
	if o == nil {
		return &xxx_MoveGroupOperation{}
	}
	return &xxx_MoveGroupOperation{
		Return: o.Return,
	}
}

func (o *MoveGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_MoveGroupOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *MoveGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *MoveGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MoveGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MoveGroupToNodeOperation structure represents the ApiMoveGroupToNode operation
type xxx_MoveGroupToNodeOperation struct {
	Group  *Group `idl:"name:hGroup" json:"group"`
	Node   *Node  `idl:"name:hNode" json:"node"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_MoveGroupToNodeOperation) OpNum() int { return 52 }

func (o *xxx_MoveGroupToNodeOperation) OpName() string { return "/clusapi2/v2/ApiMoveGroupToNode" }

func (o *xxx_MoveGroupToNodeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveGroupToNodeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group != nil {
			if err := o.Group.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Group{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node != nil {
			if err := o.Node.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_MoveGroupToNodeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group == nil {
			o.Group = &Group{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node == nil {
			o.Node = &Node{}
		}
		if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveGroupToNodeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveGroupToNodeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveGroupToNodeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MoveGroupToNodeRequest structure represents the ApiMoveGroupToNode operation request
type MoveGroupToNodeRequest struct {
	Group *Group `idl:"name:hGroup" json:"group"`
	Node  *Node  `idl:"name:hNode" json:"node"`
}

func (o *MoveGroupToNodeRequest) xxx_ToOp(ctx context.Context) *xxx_MoveGroupToNodeOperation {
	if o == nil {
		return &xxx_MoveGroupToNodeOperation{}
	}
	return &xxx_MoveGroupToNodeOperation{
		Group: o.Group,
		Node:  o.Node,
	}
}

func (o *MoveGroupToNodeRequest) xxx_FromOp(ctx context.Context, op *xxx_MoveGroupToNodeOperation) {
	if o == nil {
		return
	}
	o.Group = op.Group
	o.Node = op.Node
}
func (o *MoveGroupToNodeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *MoveGroupToNodeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MoveGroupToNodeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MoveGroupToNodeResponse structure represents the ApiMoveGroupToNode operation response
type MoveGroupToNodeResponse struct {
	// Return: The ApiMoveGroupToNode return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MoveGroupToNodeResponse) xxx_ToOp(ctx context.Context) *xxx_MoveGroupToNodeOperation {
	if o == nil {
		return &xxx_MoveGroupToNodeOperation{}
	}
	return &xxx_MoveGroupToNodeOperation{
		Return: o.Return,
	}
}

func (o *MoveGroupToNodeResponse) xxx_FromOp(ctx context.Context, op *xxx_MoveGroupToNodeOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *MoveGroupToNodeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *MoveGroupToNodeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MoveGroupToNodeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateGroupResourceEnumOperation structure represents the ApiCreateGroupResourceEnum operation
type xxx_CreateGroupResourceEnumOperation struct {
	Group      *Group    `idl:"name:hGroup" json:"group"`
	Type       uint32    `idl:"name:dwType" json:"type"`
	ReturnEnum *EnumList `idl:"name:ReturnEnum" json:"return_enum"`
	Return     uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateGroupResourceEnumOperation) OpNum() int { return 53 }

func (o *xxx_CreateGroupResourceEnumOperation) OpName() string {
	return "/clusapi2/v2/ApiCreateGroupResourceEnum"
}

func (o *xxx_CreateGroupResourceEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateGroupResourceEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group != nil {
			if err := o.Group.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Group{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwType {in} (1:(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateGroupResourceEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group == nil {
			o.Group = &Group{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwType {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateGroupResourceEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateGroupResourceEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ReturnEnum {out} (1:{pointer=ref}*(2))(2:{alias=PENUM_LIST}*(1))(3:{alias=ENUM_LIST}(struct))
	{
		if o.ReturnEnum != nil {
			_ptr_ReturnEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnEnum != nil {
					if err := o.ReturnEnum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EnumList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnEnum, _ptr_ReturnEnum); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateGroupResourceEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ReturnEnum {out} (1:{pointer=ref}*(2))(2:{alias=PENUM_LIST,pointer=ref}*(1))(3:{alias=ENUM_LIST}(struct))
	{
		_ptr_ReturnEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnEnum == nil {
				o.ReturnEnum = &EnumList{}
			}
			if err := o.ReturnEnum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReturnEnum := func(ptr interface{}) { o.ReturnEnum = *ptr.(**EnumList) }
		if err := w.ReadPointer(&o.ReturnEnum, _s_ReturnEnum, _ptr_ReturnEnum); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateGroupResourceEnumRequest structure represents the ApiCreateGroupResourceEnum operation request
type CreateGroupResourceEnumRequest struct {
	Group *Group `idl:"name:hGroup" json:"group"`
	Type  uint32 `idl:"name:dwType" json:"type"`
}

func (o *CreateGroupResourceEnumRequest) xxx_ToOp(ctx context.Context) *xxx_CreateGroupResourceEnumOperation {
	if o == nil {
		return &xxx_CreateGroupResourceEnumOperation{}
	}
	return &xxx_CreateGroupResourceEnumOperation{
		Group: o.Group,
		Type:  o.Type,
	}
}

func (o *CreateGroupResourceEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateGroupResourceEnumOperation) {
	if o == nil {
		return
	}
	o.Group = op.Group
	o.Type = op.Type
}
func (o *CreateGroupResourceEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateGroupResourceEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateGroupResourceEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateGroupResourceEnumResponse structure represents the ApiCreateGroupResourceEnum operation response
type CreateGroupResourceEnumResponse struct {
	ReturnEnum *EnumList `idl:"name:ReturnEnum" json:"return_enum"`
	// Return: The ApiCreateGroupResourceEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CreateGroupResourceEnumResponse) xxx_ToOp(ctx context.Context) *xxx_CreateGroupResourceEnumOperation {
	if o == nil {
		return &xxx_CreateGroupResourceEnumOperation{}
	}
	return &xxx_CreateGroupResourceEnumOperation{
		ReturnEnum: o.ReturnEnum,
		Return:     o.Return,
	}
}

func (o *CreateGroupResourceEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateGroupResourceEnumOperation) {
	if o == nil {
		return
	}
	o.ReturnEnum = op.ReturnEnum
	o.Return = op.Return
}
func (o *CreateGroupResourceEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateGroupResourceEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateGroupResourceEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetGroupNodeListOperation structure represents the ApiSetGroupNodeList operation
type xxx_SetGroupNodeListOperation struct {
	Group      *Group `idl:"name:hGroup" json:"group"`
	NodeList   []byte `idl:"name:lpNodeList;size_is:(cbListSize);pointer:unique" json:"node_list"`
	ListLength uint32 `idl:"name:cbListSize" json:"list_length"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SetGroupNodeListOperation) OpNum() int { return 54 }

func (o *xxx_SetGroupNodeListOperation) OpName() string { return "/clusapi2/v2/ApiSetGroupNodeList" }

func (o *xxx_SetGroupNodeListOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.NodeList != nil && o.ListLength == 0 {
		o.ListLength = uint32(len(o.NodeList))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGroupNodeListOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group != nil {
			if err := o.Group.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Group{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpNodeList {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=cbListSize](uchar))
	{
		if o.NodeList != nil || o.ListLength > 0 {
			_ptr_lpNodeList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.ListLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.NodeList {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.NodeList[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.NodeList); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.NodeList, _ptr_lpNodeList); err != nil {
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
	// cbListSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ListLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGroupNodeListOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group == nil {
			o.Group = &Group{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpNodeList {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=cbListSize](uchar))
	{
		_ptr_lpNodeList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.NodeList", sizeInfo[0])
			}
			o.NodeList = make([]byte, sizeInfo[0])
			for i1 := range o.NodeList {
				i1 := i1
				if err := w.ReadData(&o.NodeList[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_lpNodeList := func(ptr interface{}) { o.NodeList = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.NodeList, _s_lpNodeList, _ptr_lpNodeList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// cbListSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ListLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGroupNodeListOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGroupNodeListOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGroupNodeListOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetGroupNodeListRequest structure represents the ApiSetGroupNodeList operation request
type SetGroupNodeListRequest struct {
	Group      *Group `idl:"name:hGroup" json:"group"`
	NodeList   []byte `idl:"name:lpNodeList;size_is:(cbListSize);pointer:unique" json:"node_list"`
	ListLength uint32 `idl:"name:cbListSize" json:"list_length"`
}

func (o *SetGroupNodeListRequest) xxx_ToOp(ctx context.Context) *xxx_SetGroupNodeListOperation {
	if o == nil {
		return &xxx_SetGroupNodeListOperation{}
	}
	return &xxx_SetGroupNodeListOperation{
		Group:      o.Group,
		NodeList:   o.NodeList,
		ListLength: o.ListLength,
	}
}

func (o *SetGroupNodeListRequest) xxx_FromOp(ctx context.Context, op *xxx_SetGroupNodeListOperation) {
	if o == nil {
		return
	}
	o.Group = op.Group
	o.NodeList = op.NodeList
	o.ListLength = op.ListLength
}
func (o *SetGroupNodeListRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetGroupNodeListRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetGroupNodeListOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetGroupNodeListResponse structure represents the ApiSetGroupNodeList operation response
type SetGroupNodeListResponse struct {
	// Return: The ApiSetGroupNodeList return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetGroupNodeListResponse) xxx_ToOp(ctx context.Context) *xxx_SetGroupNodeListOperation {
	if o == nil {
		return &xxx_SetGroupNodeListOperation{}
	}
	return &xxx_SetGroupNodeListOperation{
		Return: o.Return,
	}
}

func (o *SetGroupNodeListResponse) xxx_FromOp(ctx context.Context, op *xxx_SetGroupNodeListOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetGroupNodeListResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetGroupNodeListResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetGroupNodeListOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateNotifyOperation structure represents the ApiCreateNotify operation
type xxx_CreateNotifyOperation struct {
	Status uint32  `idl:"name:Status" json:"status"`
	Return *Notify `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateNotifyOperation) OpNum() int { return 55 }

func (o *xxx_CreateNotifyOperation) OpName() string { return "/clusapi2/v2/ApiCreateNotify" }

func (o *xxx_CreateNotifyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNotifyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_CreateNotifyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_CreateNotifyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNotifyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return != nil {
			if err := o.Return.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CreateNotifyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return == nil {
			o.Return = &Notify{}
		}
		if err := o.Return.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// CreateNotifyRequest structure represents the ApiCreateNotify operation request
type CreateNotifyRequest struct {
}

func (o *CreateNotifyRequest) xxx_ToOp(ctx context.Context) *xxx_CreateNotifyOperation {
	if o == nil {
		return &xxx_CreateNotifyOperation{}
	}
	return &xxx_CreateNotifyOperation{}
}

func (o *CreateNotifyRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateNotifyOperation) {
	if o == nil {
		return
	}
}
func (o *CreateNotifyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateNotifyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateNotifyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateNotifyResponse structure represents the ApiCreateNotify operation response
type CreateNotifyResponse struct {
	Status uint32 `idl:"name:Status" json:"status"`
	// Return: The ApiCreateNotify return value.
	Return *Notify `idl:"name:Return" json:"return"`
}

func (o *CreateNotifyResponse) xxx_ToOp(ctx context.Context) *xxx_CreateNotifyOperation {
	if o == nil {
		return &xxx_CreateNotifyOperation{}
	}
	return &xxx_CreateNotifyOperation{
		Status: o.Status,
		Return: o.Return,
	}
}

func (o *CreateNotifyResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateNotifyOperation) {
	if o == nil {
		return
	}
	o.Status = op.Status
	o.Return = op.Return
}
func (o *CreateNotifyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateNotifyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateNotifyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseNotifyOperation structure represents the ApiCloseNotify operation
type xxx_CloseNotifyOperation struct {
	Notify *Notify `idl:"name:hNotify" json:"notify"`
	Return uint32  `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseNotifyOperation) OpNum() int { return 56 }

func (o *xxx_CloseNotifyOperation) OpName() string { return "/clusapi2/v2/ApiCloseNotify" }

func (o *xxx_CloseNotifyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseNotifyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNotify {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify != nil {
			if err := o.Notify.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseNotifyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNotify {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify == nil {
			o.Notify = &Notify{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseNotifyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseNotifyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// hNotify {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify != nil {
			if err := o.Notify.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseNotifyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// hNotify {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify == nil {
			o.Notify = &Notify{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloseNotifyRequest structure represents the ApiCloseNotify operation request
type CloseNotifyRequest struct {
	Notify *Notify `idl:"name:hNotify" json:"notify"`
}

func (o *CloseNotifyRequest) xxx_ToOp(ctx context.Context) *xxx_CloseNotifyOperation {
	if o == nil {
		return &xxx_CloseNotifyOperation{}
	}
	return &xxx_CloseNotifyOperation{
		Notify: o.Notify,
	}
}

func (o *CloseNotifyRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseNotifyOperation) {
	if o == nil {
		return
	}
	o.Notify = op.Notify
}
func (o *CloseNotifyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CloseNotifyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseNotifyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseNotifyResponse structure represents the ApiCloseNotify operation response
type CloseNotifyResponse struct {
	Notify *Notify `idl:"name:hNotify" json:"notify"`
	// Return: The ApiCloseNotify return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CloseNotifyResponse) xxx_ToOp(ctx context.Context) *xxx_CloseNotifyOperation {
	if o == nil {
		return &xxx_CloseNotifyOperation{}
	}
	return &xxx_CloseNotifyOperation{
		Notify: o.Notify,
		Return: o.Return,
	}
}

func (o *CloseNotifyResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseNotifyOperation) {
	if o == nil {
		return
	}
	o.Notify = op.Notify
	o.Return = op.Return
}
func (o *CloseNotifyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CloseNotifyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseNotifyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddNotifyClusterOperation structure represents the ApiAddNotifyCluster operation
type xxx_AddNotifyClusterOperation struct {
	Notify    *Notify  `idl:"name:hNotify" json:"notify"`
	Cluster   *Cluster `idl:"name:hCluster" json:"cluster"`
	Filter    uint32   `idl:"name:dwFilter" json:"filter"`
	NotifyKey uint32   `idl:"name:dwNotifyKey" json:"notify_key"`
	Return    uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_AddNotifyClusterOperation) OpNum() int { return 57 }

func (o *xxx_AddNotifyClusterOperation) OpName() string { return "/clusapi2/v2/ApiAddNotifyCluster" }

func (o *xxx_AddNotifyClusterOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyClusterOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify != nil {
			if err := o.Notify.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hCluster {in} (1:{context_handle, alias=HCLUSTER_RPC, names=ndr_context_handle}(struct))
	{
		if o.Cluster != nil {
			if err := o.Cluster.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Cluster{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NotifyKey); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyClusterOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify == nil {
			o.Notify = &Notify{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hCluster {in} (1:{context_handle, alias=HCLUSTER_RPC, names=ndr_context_handle}(struct))
	{
		if o.Cluster == nil {
			o.Cluster = &Cluster{}
		}
		if err := o.Cluster.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NotifyKey); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyClusterOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyClusterOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyClusterOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddNotifyClusterRequest structure represents the ApiAddNotifyCluster operation request
type AddNotifyClusterRequest struct {
	Notify    *Notify  `idl:"name:hNotify" json:"notify"`
	Cluster   *Cluster `idl:"name:hCluster" json:"cluster"`
	Filter    uint32   `idl:"name:dwFilter" json:"filter"`
	NotifyKey uint32   `idl:"name:dwNotifyKey" json:"notify_key"`
}

func (o *AddNotifyClusterRequest) xxx_ToOp(ctx context.Context) *xxx_AddNotifyClusterOperation {
	if o == nil {
		return &xxx_AddNotifyClusterOperation{}
	}
	return &xxx_AddNotifyClusterOperation{
		Notify:    o.Notify,
		Cluster:   o.Cluster,
		Filter:    o.Filter,
		NotifyKey: o.NotifyKey,
	}
}

func (o *AddNotifyClusterRequest) xxx_FromOp(ctx context.Context, op *xxx_AddNotifyClusterOperation) {
	if o == nil {
		return
	}
	o.Notify = op.Notify
	o.Cluster = op.Cluster
	o.Filter = op.Filter
	o.NotifyKey = op.NotifyKey
}
func (o *AddNotifyClusterRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddNotifyClusterRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddNotifyClusterOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddNotifyClusterResponse structure represents the ApiAddNotifyCluster operation response
type AddNotifyClusterResponse struct {
	// Return: The ApiAddNotifyCluster return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AddNotifyClusterResponse) xxx_ToOp(ctx context.Context) *xxx_AddNotifyClusterOperation {
	if o == nil {
		return &xxx_AddNotifyClusterOperation{}
	}
	return &xxx_AddNotifyClusterOperation{
		Return: o.Return,
	}
}

func (o *AddNotifyClusterResponse) xxx_FromOp(ctx context.Context, op *xxx_AddNotifyClusterOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AddNotifyClusterResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddNotifyClusterResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddNotifyClusterOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddNotifyNodeOperation structure represents the ApiAddNotifyNode operation
type xxx_AddNotifyNodeOperation struct {
	Notify        *Notify `idl:"name:hNotify" json:"notify"`
	Node          *Node   `idl:"name:hNode" json:"node"`
	Filter        uint32  `idl:"name:dwFilter" json:"filter"`
	NotifyKey     uint32  `idl:"name:dwNotifyKey" json:"notify_key"`
	StateSequence uint32  `idl:"name:dwStateSequence" json:"state_sequence"`
	Return        uint32  `idl:"name:Return" json:"return"`
}

func (o *xxx_AddNotifyNodeOperation) OpNum() int { return 58 }

func (o *xxx_AddNotifyNodeOperation) OpName() string { return "/clusapi2/v2/ApiAddNotifyNode" }

func (o *xxx_AddNotifyNodeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyNodeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify != nil {
			if err := o.Notify.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node != nil {
			if err := o.Node.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NotifyKey); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyNodeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify == nil {
			o.Notify = &Notify{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node == nil {
			o.Node = &Node{}
		}
		if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NotifyKey); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyNodeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyNodeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// dwStateSequence {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.StateSequence); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyNodeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// dwStateSequence {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.StateSequence); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddNotifyNodeRequest structure represents the ApiAddNotifyNode operation request
type AddNotifyNodeRequest struct {
	Notify    *Notify `idl:"name:hNotify" json:"notify"`
	Node      *Node   `idl:"name:hNode" json:"node"`
	Filter    uint32  `idl:"name:dwFilter" json:"filter"`
	NotifyKey uint32  `idl:"name:dwNotifyKey" json:"notify_key"`
}

func (o *AddNotifyNodeRequest) xxx_ToOp(ctx context.Context) *xxx_AddNotifyNodeOperation {
	if o == nil {
		return &xxx_AddNotifyNodeOperation{}
	}
	return &xxx_AddNotifyNodeOperation{
		Notify:    o.Notify,
		Node:      o.Node,
		Filter:    o.Filter,
		NotifyKey: o.NotifyKey,
	}
}

func (o *AddNotifyNodeRequest) xxx_FromOp(ctx context.Context, op *xxx_AddNotifyNodeOperation) {
	if o == nil {
		return
	}
	o.Notify = op.Notify
	o.Node = op.Node
	o.Filter = op.Filter
	o.NotifyKey = op.NotifyKey
}
func (o *AddNotifyNodeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddNotifyNodeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddNotifyNodeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddNotifyNodeResponse structure represents the ApiAddNotifyNode operation response
type AddNotifyNodeResponse struct {
	StateSequence uint32 `idl:"name:dwStateSequence" json:"state_sequence"`
	// Return: The ApiAddNotifyNode return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AddNotifyNodeResponse) xxx_ToOp(ctx context.Context) *xxx_AddNotifyNodeOperation {
	if o == nil {
		return &xxx_AddNotifyNodeOperation{}
	}
	return &xxx_AddNotifyNodeOperation{
		StateSequence: o.StateSequence,
		Return:        o.Return,
	}
}

func (o *AddNotifyNodeResponse) xxx_FromOp(ctx context.Context, op *xxx_AddNotifyNodeOperation) {
	if o == nil {
		return
	}
	o.StateSequence = op.StateSequence
	o.Return = op.Return
}
func (o *AddNotifyNodeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddNotifyNodeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddNotifyNodeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddNotifyGroupOperation structure represents the ApiAddNotifyGroup operation
type xxx_AddNotifyGroupOperation struct {
	Notify        *Notify `idl:"name:hNotify" json:"notify"`
	Group         *Group  `idl:"name:hGroup" json:"group"`
	Filter        uint32  `idl:"name:dwFilter" json:"filter"`
	NotifyKey     uint32  `idl:"name:dwNotifyKey" json:"notify_key"`
	StateSequence uint32  `idl:"name:dwStateSequence" json:"state_sequence"`
	Return        uint32  `idl:"name:Return" json:"return"`
}

func (o *xxx_AddNotifyGroupOperation) OpNum() int { return 59 }

func (o *xxx_AddNotifyGroupOperation) OpName() string { return "/clusapi2/v2/ApiAddNotifyGroup" }

func (o *xxx_AddNotifyGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify != nil {
			if err := o.Notify.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group != nil {
			if err := o.Group.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Group{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NotifyKey); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify == nil {
			o.Notify = &Notify{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group == nil {
			o.Group = &Group{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NotifyKey); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// dwStateSequence {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.StateSequence); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// dwStateSequence {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.StateSequence); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddNotifyGroupRequest structure represents the ApiAddNotifyGroup operation request
type AddNotifyGroupRequest struct {
	Notify    *Notify `idl:"name:hNotify" json:"notify"`
	Group     *Group  `idl:"name:hGroup" json:"group"`
	Filter    uint32  `idl:"name:dwFilter" json:"filter"`
	NotifyKey uint32  `idl:"name:dwNotifyKey" json:"notify_key"`
}

func (o *AddNotifyGroupRequest) xxx_ToOp(ctx context.Context) *xxx_AddNotifyGroupOperation {
	if o == nil {
		return &xxx_AddNotifyGroupOperation{}
	}
	return &xxx_AddNotifyGroupOperation{
		Notify:    o.Notify,
		Group:     o.Group,
		Filter:    o.Filter,
		NotifyKey: o.NotifyKey,
	}
}

func (o *AddNotifyGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_AddNotifyGroupOperation) {
	if o == nil {
		return
	}
	o.Notify = op.Notify
	o.Group = op.Group
	o.Filter = op.Filter
	o.NotifyKey = op.NotifyKey
}
func (o *AddNotifyGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddNotifyGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddNotifyGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddNotifyGroupResponse structure represents the ApiAddNotifyGroup operation response
type AddNotifyGroupResponse struct {
	StateSequence uint32 `idl:"name:dwStateSequence" json:"state_sequence"`
	// Return: The ApiAddNotifyGroup return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AddNotifyGroupResponse) xxx_ToOp(ctx context.Context) *xxx_AddNotifyGroupOperation {
	if o == nil {
		return &xxx_AddNotifyGroupOperation{}
	}
	return &xxx_AddNotifyGroupOperation{
		StateSequence: o.StateSequence,
		Return:        o.Return,
	}
}

func (o *AddNotifyGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_AddNotifyGroupOperation) {
	if o == nil {
		return
	}
	o.StateSequence = op.StateSequence
	o.Return = op.Return
}
func (o *AddNotifyGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddNotifyGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddNotifyGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddNotifyResourceOperation structure represents the ApiAddNotifyResource operation
type xxx_AddNotifyResourceOperation struct {
	Notify        *Notify   `idl:"name:hNotify" json:"notify"`
	Resource      *Resource `idl:"name:hResource" json:"resource"`
	Filter        uint32    `idl:"name:dwFilter" json:"filter"`
	NotifyKey     uint32    `idl:"name:dwNotifyKey" json:"notify_key"`
	StateSequence uint32    `idl:"name:dwStateSequence" json:"state_sequence"`
	Return        uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_AddNotifyResourceOperation) OpNum() int { return 60 }

func (o *xxx_AddNotifyResourceOperation) OpName() string { return "/clusapi2/v2/ApiAddNotifyResource" }

func (o *xxx_AddNotifyResourceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyResourceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify != nil {
			if err := o.Notify.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NotifyKey); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyResourceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify == nil {
			o.Notify = &Notify{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NotifyKey); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyResourceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyResourceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// dwStateSequence {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.StateSequence); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyResourceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// dwStateSequence {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.StateSequence); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddNotifyResourceRequest structure represents the ApiAddNotifyResource operation request
type AddNotifyResourceRequest struct {
	Notify    *Notify   `idl:"name:hNotify" json:"notify"`
	Resource  *Resource `idl:"name:hResource" json:"resource"`
	Filter    uint32    `idl:"name:dwFilter" json:"filter"`
	NotifyKey uint32    `idl:"name:dwNotifyKey" json:"notify_key"`
}

func (o *AddNotifyResourceRequest) xxx_ToOp(ctx context.Context) *xxx_AddNotifyResourceOperation {
	if o == nil {
		return &xxx_AddNotifyResourceOperation{}
	}
	return &xxx_AddNotifyResourceOperation{
		Notify:    o.Notify,
		Resource:  o.Resource,
		Filter:    o.Filter,
		NotifyKey: o.NotifyKey,
	}
}

func (o *AddNotifyResourceRequest) xxx_FromOp(ctx context.Context, op *xxx_AddNotifyResourceOperation) {
	if o == nil {
		return
	}
	o.Notify = op.Notify
	o.Resource = op.Resource
	o.Filter = op.Filter
	o.NotifyKey = op.NotifyKey
}
func (o *AddNotifyResourceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddNotifyResourceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddNotifyResourceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddNotifyResourceResponse structure represents the ApiAddNotifyResource operation response
type AddNotifyResourceResponse struct {
	StateSequence uint32 `idl:"name:dwStateSequence" json:"state_sequence"`
	// Return: The ApiAddNotifyResource return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AddNotifyResourceResponse) xxx_ToOp(ctx context.Context) *xxx_AddNotifyResourceOperation {
	if o == nil {
		return &xxx_AddNotifyResourceOperation{}
	}
	return &xxx_AddNotifyResourceOperation{
		StateSequence: o.StateSequence,
		Return:        o.Return,
	}
}

func (o *AddNotifyResourceResponse) xxx_FromOp(ctx context.Context, op *xxx_AddNotifyResourceOperation) {
	if o == nil {
		return
	}
	o.StateSequence = op.StateSequence
	o.Return = op.Return
}
func (o *AddNotifyResourceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddNotifyResourceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddNotifyResourceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddNotifyKeyOperation structure represents the ApiAddNotifyKey operation
type xxx_AddNotifyKeyOperation struct {
	Notify       *Notify `idl:"name:hNotify" json:"notify"`
	Key          *Key    `idl:"name:hKey" json:"key"`
	NotifyKey    uint32  `idl:"name:dwNotifyKey" json:"notify_key"`
	Filter       uint32  `idl:"name:dwFilter" json:"filter"`
	WatchSubTree bool    `idl:"name:WatchSubTree" json:"watch_sub_tree"`
	Return       uint32  `idl:"name:Return" json:"return"`
}

func (o *xxx_AddNotifyKeyOperation) OpNum() int { return 61 }

func (o *xxx_AddNotifyKeyOperation) OpName() string { return "/clusapi2/v2/ApiAddNotifyKey" }

func (o *xxx_AddNotifyKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify != nil {
			if err := o.Notify.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NotifyKey); err != nil {
			return err
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Filter); err != nil {
			return err
		}
	}
	// WatchSubTree {in} (1:{alias=BOOL}(int32))
	{
		if !o.WatchSubTree {
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

func (o *xxx_AddNotifyKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify == nil {
			o.Notify = &Notify{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hKey {in} (1:{context_handle, alias=HKEY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NotifyKey); err != nil {
			return err
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Filter); err != nil {
			return err
		}
	}
	// WatchSubTree {in} (1:{alias=BOOL}(int32))
	{
		var _bWatchSubTree int32
		if err := w.ReadData(&_bWatchSubTree); err != nil {
			return err
		}
		o.WatchSubTree = _bWatchSubTree != 0
	}
	return nil
}

func (o *xxx_AddNotifyKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddNotifyKeyRequest structure represents the ApiAddNotifyKey operation request
type AddNotifyKeyRequest struct {
	Notify       *Notify `idl:"name:hNotify" json:"notify"`
	Key          *Key    `idl:"name:hKey" json:"key"`
	NotifyKey    uint32  `idl:"name:dwNotifyKey" json:"notify_key"`
	Filter       uint32  `idl:"name:dwFilter" json:"filter"`
	WatchSubTree bool    `idl:"name:WatchSubTree" json:"watch_sub_tree"`
}

func (o *AddNotifyKeyRequest) xxx_ToOp(ctx context.Context) *xxx_AddNotifyKeyOperation {
	if o == nil {
		return &xxx_AddNotifyKeyOperation{}
	}
	return &xxx_AddNotifyKeyOperation{
		Notify:       o.Notify,
		Key:          o.Key,
		NotifyKey:    o.NotifyKey,
		Filter:       o.Filter,
		WatchSubTree: o.WatchSubTree,
	}
}

func (o *AddNotifyKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_AddNotifyKeyOperation) {
	if o == nil {
		return
	}
	o.Notify = op.Notify
	o.Key = op.Key
	o.NotifyKey = op.NotifyKey
	o.Filter = op.Filter
	o.WatchSubTree = op.WatchSubTree
}
func (o *AddNotifyKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddNotifyKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddNotifyKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddNotifyKeyResponse structure represents the ApiAddNotifyKey operation response
type AddNotifyKeyResponse struct {
	// Return: The ApiAddNotifyKey return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AddNotifyKeyResponse) xxx_ToOp(ctx context.Context) *xxx_AddNotifyKeyOperation {
	if o == nil {
		return &xxx_AddNotifyKeyOperation{}
	}
	return &xxx_AddNotifyKeyOperation{
		Return: o.Return,
	}
}

func (o *AddNotifyKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_AddNotifyKeyOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AddNotifyKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddNotifyKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddNotifyKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReAddNotifyNodeOperation structure represents the ApiReAddNotifyNode operation
type xxx_ReAddNotifyNodeOperation struct {
	Notify        *Notify `idl:"name:hNotify" json:"notify"`
	Node          *Node   `idl:"name:hNode" json:"node"`
	Filter        uint32  `idl:"name:dwFilter" json:"filter"`
	NotifyKey     uint32  `idl:"name:dwNotifyKey" json:"notify_key"`
	StateSequence uint32  `idl:"name:StateSequence" json:"state_sequence"`
	Return        uint32  `idl:"name:Return" json:"return"`
}

func (o *xxx_ReAddNotifyNodeOperation) OpNum() int { return 62 }

func (o *xxx_ReAddNotifyNodeOperation) OpName() string { return "/clusapi2/v2/ApiReAddNotifyNode" }

func (o *xxx_ReAddNotifyNodeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyNodeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify != nil {
			if err := o.Notify.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node != nil {
			if err := o.Node.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NotifyKey); err != nil {
			return err
		}
	}
	// StateSequence {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.StateSequence); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyNodeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify == nil {
			o.Notify = &Notify{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node == nil {
			o.Node = &Node{}
		}
		if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NotifyKey); err != nil {
			return err
		}
	}
	// StateSequence {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.StateSequence); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyNodeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyNodeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyNodeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ReAddNotifyNodeRequest structure represents the ApiReAddNotifyNode operation request
type ReAddNotifyNodeRequest struct {
	Notify        *Notify `idl:"name:hNotify" json:"notify"`
	Node          *Node   `idl:"name:hNode" json:"node"`
	Filter        uint32  `idl:"name:dwFilter" json:"filter"`
	NotifyKey     uint32  `idl:"name:dwNotifyKey" json:"notify_key"`
	StateSequence uint32  `idl:"name:StateSequence" json:"state_sequence"`
}

func (o *ReAddNotifyNodeRequest) xxx_ToOp(ctx context.Context) *xxx_ReAddNotifyNodeOperation {
	if o == nil {
		return &xxx_ReAddNotifyNodeOperation{}
	}
	return &xxx_ReAddNotifyNodeOperation{
		Notify:        o.Notify,
		Node:          o.Node,
		Filter:        o.Filter,
		NotifyKey:     o.NotifyKey,
		StateSequence: o.StateSequence,
	}
}

func (o *ReAddNotifyNodeRequest) xxx_FromOp(ctx context.Context, op *xxx_ReAddNotifyNodeOperation) {
	if o == nil {
		return
	}
	o.Notify = op.Notify
	o.Node = op.Node
	o.Filter = op.Filter
	o.NotifyKey = op.NotifyKey
	o.StateSequence = op.StateSequence
}
func (o *ReAddNotifyNodeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ReAddNotifyNodeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReAddNotifyNodeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReAddNotifyNodeResponse structure represents the ApiReAddNotifyNode operation response
type ReAddNotifyNodeResponse struct {
	// Return: The ApiReAddNotifyNode return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ReAddNotifyNodeResponse) xxx_ToOp(ctx context.Context) *xxx_ReAddNotifyNodeOperation {
	if o == nil {
		return &xxx_ReAddNotifyNodeOperation{}
	}
	return &xxx_ReAddNotifyNodeOperation{
		Return: o.Return,
	}
}

func (o *ReAddNotifyNodeResponse) xxx_FromOp(ctx context.Context, op *xxx_ReAddNotifyNodeOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ReAddNotifyNodeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ReAddNotifyNodeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReAddNotifyNodeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReAddNotifyGroupOperation structure represents the ApiReAddNotifyGroup operation
type xxx_ReAddNotifyGroupOperation struct {
	Notify        *Notify `idl:"name:hNotify" json:"notify"`
	Group         *Group  `idl:"name:hGroup" json:"group"`
	Filter        uint32  `idl:"name:dwFilter" json:"filter"`
	NotifyKey     uint32  `idl:"name:dwNotifyKey" json:"notify_key"`
	StateSequence uint32  `idl:"name:StateSequence" json:"state_sequence"`
	Return        uint32  `idl:"name:Return" json:"return"`
}

func (o *xxx_ReAddNotifyGroupOperation) OpNum() int { return 63 }

func (o *xxx_ReAddNotifyGroupOperation) OpName() string { return "/clusapi2/v2/ApiReAddNotifyGroup" }

func (o *xxx_ReAddNotifyGroupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyGroupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify != nil {
			if err := o.Notify.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group != nil {
			if err := o.Group.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Group{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NotifyKey); err != nil {
			return err
		}
	}
	// StateSequence {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.StateSequence); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyGroupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify == nil {
			o.Notify = &Notify{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group == nil {
			o.Group = &Group{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NotifyKey); err != nil {
			return err
		}
	}
	// StateSequence {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.StateSequence); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyGroupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyGroupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyGroupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ReAddNotifyGroupRequest structure represents the ApiReAddNotifyGroup operation request
type ReAddNotifyGroupRequest struct {
	Notify        *Notify `idl:"name:hNotify" json:"notify"`
	Group         *Group  `idl:"name:hGroup" json:"group"`
	Filter        uint32  `idl:"name:dwFilter" json:"filter"`
	NotifyKey     uint32  `idl:"name:dwNotifyKey" json:"notify_key"`
	StateSequence uint32  `idl:"name:StateSequence" json:"state_sequence"`
}

func (o *ReAddNotifyGroupRequest) xxx_ToOp(ctx context.Context) *xxx_ReAddNotifyGroupOperation {
	if o == nil {
		return &xxx_ReAddNotifyGroupOperation{}
	}
	return &xxx_ReAddNotifyGroupOperation{
		Notify:        o.Notify,
		Group:         o.Group,
		Filter:        o.Filter,
		NotifyKey:     o.NotifyKey,
		StateSequence: o.StateSequence,
	}
}

func (o *ReAddNotifyGroupRequest) xxx_FromOp(ctx context.Context, op *xxx_ReAddNotifyGroupOperation) {
	if o == nil {
		return
	}
	o.Notify = op.Notify
	o.Group = op.Group
	o.Filter = op.Filter
	o.NotifyKey = op.NotifyKey
	o.StateSequence = op.StateSequence
}
func (o *ReAddNotifyGroupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ReAddNotifyGroupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReAddNotifyGroupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReAddNotifyGroupResponse structure represents the ApiReAddNotifyGroup operation response
type ReAddNotifyGroupResponse struct {
	// Return: The ApiReAddNotifyGroup return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ReAddNotifyGroupResponse) xxx_ToOp(ctx context.Context) *xxx_ReAddNotifyGroupOperation {
	if o == nil {
		return &xxx_ReAddNotifyGroupOperation{}
	}
	return &xxx_ReAddNotifyGroupOperation{
		Return: o.Return,
	}
}

func (o *ReAddNotifyGroupResponse) xxx_FromOp(ctx context.Context, op *xxx_ReAddNotifyGroupOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ReAddNotifyGroupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ReAddNotifyGroupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReAddNotifyGroupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReAddNotifyResourceOperation structure represents the ApiReAddNotifyResource operation
type xxx_ReAddNotifyResourceOperation struct {
	Notify        *Notify   `idl:"name:hNotify" json:"notify"`
	Resource      *Resource `idl:"name:hResource" json:"resource"`
	Filter        uint32    `idl:"name:dwFilter" json:"filter"`
	NotifyKey     uint32    `idl:"name:dwNotifyKey" json:"notify_key"`
	StateSequence uint32    `idl:"name:StateSequence" json:"state_sequence"`
	Return        uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_ReAddNotifyResourceOperation) OpNum() int { return 64 }

func (o *xxx_ReAddNotifyResourceOperation) OpName() string {
	return "/clusapi2/v2/ApiReAddNotifyResource"
}

func (o *xxx_ReAddNotifyResourceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyResourceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify != nil {
			if err := o.Notify.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NotifyKey); err != nil {
			return err
		}
	}
	// StateSequence {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.StateSequence); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyResourceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify == nil {
			o.Notify = &Notify{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NotifyKey); err != nil {
			return err
		}
	}
	// StateSequence {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.StateSequence); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyResourceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyResourceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyResourceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ReAddNotifyResourceRequest structure represents the ApiReAddNotifyResource operation request
type ReAddNotifyResourceRequest struct {
	Notify        *Notify   `idl:"name:hNotify" json:"notify"`
	Resource      *Resource `idl:"name:hResource" json:"resource"`
	Filter        uint32    `idl:"name:dwFilter" json:"filter"`
	NotifyKey     uint32    `idl:"name:dwNotifyKey" json:"notify_key"`
	StateSequence uint32    `idl:"name:StateSequence" json:"state_sequence"`
}

func (o *ReAddNotifyResourceRequest) xxx_ToOp(ctx context.Context) *xxx_ReAddNotifyResourceOperation {
	if o == nil {
		return &xxx_ReAddNotifyResourceOperation{}
	}
	return &xxx_ReAddNotifyResourceOperation{
		Notify:        o.Notify,
		Resource:      o.Resource,
		Filter:        o.Filter,
		NotifyKey:     o.NotifyKey,
		StateSequence: o.StateSequence,
	}
}

func (o *ReAddNotifyResourceRequest) xxx_FromOp(ctx context.Context, op *xxx_ReAddNotifyResourceOperation) {
	if o == nil {
		return
	}
	o.Notify = op.Notify
	o.Resource = op.Resource
	o.Filter = op.Filter
	o.NotifyKey = op.NotifyKey
	o.StateSequence = op.StateSequence
}
func (o *ReAddNotifyResourceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ReAddNotifyResourceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReAddNotifyResourceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReAddNotifyResourceResponse structure represents the ApiReAddNotifyResource operation response
type ReAddNotifyResourceResponse struct {
	// Return: The ApiReAddNotifyResource return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ReAddNotifyResourceResponse) xxx_ToOp(ctx context.Context) *xxx_ReAddNotifyResourceOperation {
	if o == nil {
		return &xxx_ReAddNotifyResourceOperation{}
	}
	return &xxx_ReAddNotifyResourceOperation{
		Return: o.Return,
	}
}

func (o *ReAddNotifyResourceResponse) xxx_FromOp(ctx context.Context, op *xxx_ReAddNotifyResourceOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ReAddNotifyResourceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ReAddNotifyResourceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReAddNotifyResourceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNotifyOperation structure represents the ApiGetNotify operation
type xxx_GetNotifyOperation struct {
	Notify        *Notify `idl:"name:hNotify" json:"notify"`
	Timeout       uint32  `idl:"name:Timeout" json:"timeout"`
	NotifyKey     uint32  `idl:"name:dwNotifyKey" json:"notify_key"`
	Filter        uint32  `idl:"name:dwFilter" json:"filter"`
	StateSequence uint32  `idl:"name:dwStateSequence" json:"state_sequence"`
	Name          string  `idl:"name:Name;string" json:"name"`
	Return        uint32  `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNotifyOperation) OpNum() int { return 65 }

func (o *xxx_GetNotifyOperation) OpName() string { return "/clusapi2/v2/ApiGetNotify" }

func (o *xxx_GetNotifyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNotifyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify != nil {
			if err := o.Notify.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Timeout {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNotifyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify == nil {
			o.Notify = &Notify{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Timeout {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNotifyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNotifyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// dwNotifyKey {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NotifyKey); err != nil {
			return err
		}
	}
	// dwFilter {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Filter); err != nil {
			return err
		}
	}
	// dwStateSequence {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.StateSequence); err != nil {
			return err
		}
	}
	// Name {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Name != "" {
			_ptr_Name := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_Name); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNotifyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// dwNotifyKey {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NotifyKey); err != nil {
			return err
		}
	}
	// dwFilter {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Filter); err != nil {
			return err
		}
	}
	// dwStateSequence {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.StateSequence); err != nil {
			return err
		}
	}
	// Name {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_Name := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
				return err
			}
			return nil
		})
		_s_Name := func(ptr interface{}) { o.Name = *ptr.(*string) }
		if err := w.ReadPointer(&o.Name, _s_Name, _ptr_Name); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNotifyRequest structure represents the ApiGetNotify operation request
type GetNotifyRequest struct {
	Notify  *Notify `idl:"name:hNotify" json:"notify"`
	Timeout uint32  `idl:"name:Timeout" json:"timeout"`
}

func (o *GetNotifyRequest) xxx_ToOp(ctx context.Context) *xxx_GetNotifyOperation {
	if o == nil {
		return &xxx_GetNotifyOperation{}
	}
	return &xxx_GetNotifyOperation{
		Notify:  o.Notify,
		Timeout: o.Timeout,
	}
}

func (o *GetNotifyRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNotifyOperation) {
	if o == nil {
		return
	}
	o.Notify = op.Notify
	o.Timeout = op.Timeout
}
func (o *GetNotifyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetNotifyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNotifyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNotifyResponse structure represents the ApiGetNotify operation response
type GetNotifyResponse struct {
	NotifyKey     uint32 `idl:"name:dwNotifyKey" json:"notify_key"`
	Filter        uint32 `idl:"name:dwFilter" json:"filter"`
	StateSequence uint32 `idl:"name:dwStateSequence" json:"state_sequence"`
	Name          string `idl:"name:Name;string" json:"name"`
	// Return: The ApiGetNotify return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetNotifyResponse) xxx_ToOp(ctx context.Context) *xxx_GetNotifyOperation {
	if o == nil {
		return &xxx_GetNotifyOperation{}
	}
	return &xxx_GetNotifyOperation{
		NotifyKey:     o.NotifyKey,
		Filter:        o.Filter,
		StateSequence: o.StateSequence,
		Name:          o.Name,
		Return:        o.Return,
	}
}

func (o *GetNotifyResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNotifyOperation) {
	if o == nil {
		return
	}
	o.NotifyKey = op.NotifyKey
	o.Filter = op.Filter
	o.StateSequence = op.StateSequence
	o.Name = op.Name
	o.Return = op.Return
}
func (o *GetNotifyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetNotifyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNotifyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenNodeOperation structure represents the ApiOpenNode operation
type xxx_OpenNodeOperation struct {
	NodeName string `idl:"name:lpszNodeName;string" json:"node_name"`
	Status   uint32 `idl:"name:Status" json:"status"`
	Return   *Node  `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenNodeOperation) OpNum() int { return 66 }

func (o *xxx_OpenNodeOperation) OpName() string { return "/clusapi2/v2/ApiOpenNode" }

func (o *xxx_OpenNodeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNodeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpszNodeName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.NodeName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNodeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpszNodeName {in} (1:{string, pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.NodeName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNodeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNodeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return != nil {
			if err := o.Return.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_OpenNodeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return == nil {
			o.Return = &Node{}
		}
		if err := o.Return.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// OpenNodeRequest structure represents the ApiOpenNode operation request
type OpenNodeRequest struct {
	NodeName string `idl:"name:lpszNodeName;string" json:"node_name"`
}

func (o *OpenNodeRequest) xxx_ToOp(ctx context.Context) *xxx_OpenNodeOperation {
	if o == nil {
		return &xxx_OpenNodeOperation{}
	}
	return &xxx_OpenNodeOperation{
		NodeName: o.NodeName,
	}
}

func (o *OpenNodeRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenNodeOperation) {
	if o == nil {
		return
	}
	o.NodeName = op.NodeName
}
func (o *OpenNodeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OpenNodeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenNodeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenNodeResponse structure represents the ApiOpenNode operation response
type OpenNodeResponse struct {
	Status uint32 `idl:"name:Status" json:"status"`
	// Return: The ApiOpenNode return value.
	Return *Node `idl:"name:Return" json:"return"`
}

func (o *OpenNodeResponse) xxx_ToOp(ctx context.Context) *xxx_OpenNodeOperation {
	if o == nil {
		return &xxx_OpenNodeOperation{}
	}
	return &xxx_OpenNodeOperation{
		Status: o.Status,
		Return: o.Return,
	}
}

func (o *OpenNodeResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenNodeOperation) {
	if o == nil {
		return
	}
	o.Status = op.Status
	o.Return = op.Return
}
func (o *OpenNodeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OpenNodeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenNodeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseNodeOperation structure represents the ApiCloseNode operation
type xxx_CloseNodeOperation struct {
	Node   *Node  `idl:"name:Node" json:"node"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseNodeOperation) OpNum() int { return 67 }

func (o *xxx_CloseNodeOperation) OpName() string { return "/clusapi2/v2/ApiCloseNode" }

func (o *xxx_CloseNodeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseNodeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Node {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node != nil {
			if err := o.Node.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseNodeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Node {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node == nil {
			o.Node = &Node{}
		}
		if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseNodeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseNodeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Node {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node != nil {
			if err := o.Node.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseNodeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Node {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node == nil {
			o.Node = &Node{}
		}
		if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloseNodeRequest structure represents the ApiCloseNode operation request
type CloseNodeRequest struct {
	Node *Node `idl:"name:Node" json:"node"`
}

func (o *CloseNodeRequest) xxx_ToOp(ctx context.Context) *xxx_CloseNodeOperation {
	if o == nil {
		return &xxx_CloseNodeOperation{}
	}
	return &xxx_CloseNodeOperation{
		Node: o.Node,
	}
}

func (o *CloseNodeRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseNodeOperation) {
	if o == nil {
		return
	}
	o.Node = op.Node
}
func (o *CloseNodeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CloseNodeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseNodeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseNodeResponse structure represents the ApiCloseNode operation response
type CloseNodeResponse struct {
	Node *Node `idl:"name:Node" json:"node"`
	// Return: The ApiCloseNode return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CloseNodeResponse) xxx_ToOp(ctx context.Context) *xxx_CloseNodeOperation {
	if o == nil {
		return &xxx_CloseNodeOperation{}
	}
	return &xxx_CloseNodeOperation{
		Node:   o.Node,
		Return: o.Return,
	}
}

func (o *CloseNodeResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseNodeOperation) {
	if o == nil {
		return
	}
	o.Node = op.Node
	o.Return = op.Return
}
func (o *CloseNodeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CloseNodeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseNodeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNodeStateOperation structure represents the ApiGetNodeState operation
type xxx_GetNodeStateOperation struct {
	Node   *Node  `idl:"name:hNode" json:"node"`
	State  uint32 `idl:"name:State" json:"state"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNodeStateOperation) OpNum() int { return 68 }

func (o *xxx_GetNodeStateOperation) OpName() string { return "/clusapi2/v2/ApiGetNodeState" }

func (o *xxx_GetNodeStateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNodeStateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node != nil {
			if err := o.Node.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetNodeStateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node == nil {
			o.Node = &Node{}
		}
		if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNodeStateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNodeStateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// State {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.State); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNodeStateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// State {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.State); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNodeStateRequest structure represents the ApiGetNodeState operation request
type GetNodeStateRequest struct {
	Node *Node `idl:"name:hNode" json:"node"`
}

func (o *GetNodeStateRequest) xxx_ToOp(ctx context.Context) *xxx_GetNodeStateOperation {
	if o == nil {
		return &xxx_GetNodeStateOperation{}
	}
	return &xxx_GetNodeStateOperation{
		Node: o.Node,
	}
}

func (o *GetNodeStateRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNodeStateOperation) {
	if o == nil {
		return
	}
	o.Node = op.Node
}
func (o *GetNodeStateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetNodeStateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNodeStateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNodeStateResponse structure represents the ApiGetNodeState operation response
type GetNodeStateResponse struct {
	State uint32 `idl:"name:State" json:"state"`
	// Return: The ApiGetNodeState return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetNodeStateResponse) xxx_ToOp(ctx context.Context) *xxx_GetNodeStateOperation {
	if o == nil {
		return &xxx_GetNodeStateOperation{}
	}
	return &xxx_GetNodeStateOperation{
		State:  o.State,
		Return: o.Return,
	}
}

func (o *GetNodeStateResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNodeStateOperation) {
	if o == nil {
		return
	}
	o.State = op.State
	o.Return = op.Return
}
func (o *GetNodeStateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetNodeStateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNodeStateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PauseNodeOperation structure represents the ApiPauseNode operation
type xxx_PauseNodeOperation struct {
	Node   *Node  `idl:"name:hNode" json:"node"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_PauseNodeOperation) OpNum() int { return 69 }

func (o *xxx_PauseNodeOperation) OpName() string { return "/clusapi2/v2/ApiPauseNode" }

func (o *xxx_PauseNodeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PauseNodeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node != nil {
			if err := o.Node.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_PauseNodeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node == nil {
			o.Node = &Node{}
		}
		if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PauseNodeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PauseNodeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PauseNodeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PauseNodeRequest structure represents the ApiPauseNode operation request
type PauseNodeRequest struct {
	Node *Node `idl:"name:hNode" json:"node"`
}

func (o *PauseNodeRequest) xxx_ToOp(ctx context.Context) *xxx_PauseNodeOperation {
	if o == nil {
		return &xxx_PauseNodeOperation{}
	}
	return &xxx_PauseNodeOperation{
		Node: o.Node,
	}
}

func (o *PauseNodeRequest) xxx_FromOp(ctx context.Context, op *xxx_PauseNodeOperation) {
	if o == nil {
		return
	}
	o.Node = op.Node
}
func (o *PauseNodeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PauseNodeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PauseNodeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PauseNodeResponse structure represents the ApiPauseNode operation response
type PauseNodeResponse struct {
	// Return: The ApiPauseNode return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *PauseNodeResponse) xxx_ToOp(ctx context.Context) *xxx_PauseNodeOperation {
	if o == nil {
		return &xxx_PauseNodeOperation{}
	}
	return &xxx_PauseNodeOperation{
		Return: o.Return,
	}
}

func (o *PauseNodeResponse) xxx_FromOp(ctx context.Context, op *xxx_PauseNodeOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *PauseNodeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PauseNodeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PauseNodeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ResumeNodeOperation structure represents the ApiResumeNode operation
type xxx_ResumeNodeOperation struct {
	Node   *Node  `idl:"name:hNode" json:"node"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_ResumeNodeOperation) OpNum() int { return 70 }

func (o *xxx_ResumeNodeOperation) OpName() string { return "/clusapi2/v2/ApiResumeNode" }

func (o *xxx_ResumeNodeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResumeNodeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node != nil {
			if err := o.Node.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ResumeNodeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node == nil {
			o.Node = &Node{}
		}
		if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResumeNodeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResumeNodeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResumeNodeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ResumeNodeRequest structure represents the ApiResumeNode operation request
type ResumeNodeRequest struct {
	Node *Node `idl:"name:hNode" json:"node"`
}

func (o *ResumeNodeRequest) xxx_ToOp(ctx context.Context) *xxx_ResumeNodeOperation {
	if o == nil {
		return &xxx_ResumeNodeOperation{}
	}
	return &xxx_ResumeNodeOperation{
		Node: o.Node,
	}
}

func (o *ResumeNodeRequest) xxx_FromOp(ctx context.Context, op *xxx_ResumeNodeOperation) {
	if o == nil {
		return
	}
	o.Node = op.Node
}
func (o *ResumeNodeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ResumeNodeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResumeNodeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ResumeNodeResponse structure represents the ApiResumeNode operation response
type ResumeNodeResponse struct {
	// Return: The ApiResumeNode return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ResumeNodeResponse) xxx_ToOp(ctx context.Context) *xxx_ResumeNodeOperation {
	if o == nil {
		return &xxx_ResumeNodeOperation{}
	}
	return &xxx_ResumeNodeOperation{
		Return: o.Return,
	}
}

func (o *ResumeNodeResponse) xxx_FromOp(ctx context.Context, op *xxx_ResumeNodeOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ResumeNodeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ResumeNodeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResumeNodeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EvictNodeOperation structure represents the ApiEvictNode operation
type xxx_EvictNodeOperation struct {
	Node   *Node  `idl:"name:hNode" json:"node"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_EvictNodeOperation) OpNum() int { return 71 }

func (o *xxx_EvictNodeOperation) OpName() string { return "/clusapi2/v2/ApiEvictNode" }

func (o *xxx_EvictNodeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EvictNodeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node != nil {
			if err := o.Node.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_EvictNodeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node == nil {
			o.Node = &Node{}
		}
		if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EvictNodeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EvictNodeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EvictNodeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EvictNodeRequest structure represents the ApiEvictNode operation request
type EvictNodeRequest struct {
	Node *Node `idl:"name:hNode" json:"node"`
}

func (o *EvictNodeRequest) xxx_ToOp(ctx context.Context) *xxx_EvictNodeOperation {
	if o == nil {
		return &xxx_EvictNodeOperation{}
	}
	return &xxx_EvictNodeOperation{
		Node: o.Node,
	}
}

func (o *EvictNodeRequest) xxx_FromOp(ctx context.Context, op *xxx_EvictNodeOperation) {
	if o == nil {
		return
	}
	o.Node = op.Node
}
func (o *EvictNodeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *EvictNodeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EvictNodeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EvictNodeResponse structure represents the ApiEvictNode operation response
type EvictNodeResponse struct {
	// Return: The ApiEvictNode return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EvictNodeResponse) xxx_ToOp(ctx context.Context) *xxx_EvictNodeOperation {
	if o == nil {
		return &xxx_EvictNodeOperation{}
	}
	return &xxx_EvictNodeOperation{
		Return: o.Return,
	}
}

func (o *EvictNodeResponse) xxx_FromOp(ctx context.Context, op *xxx_EvictNodeOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *EvictNodeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *EvictNodeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EvictNodeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_NodeResourceControlOperation structure represents the ApiNodeResourceControl operation
type xxx_NodeResourceControlOperation struct {
	Resource       *Resource `idl:"name:hResource" json:"resource"`
	Node           *Node     `idl:"name:hNode" json:"node"`
	ControlCode    uint32    `idl:"name:dwControlCode" json:"control_code"`
	InBuffer       []byte    `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize   uint32    `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBuffer      []byte    `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	OutBufferSize  uint32    `idl:"name:nOutBufferSize" json:"out_buffer_size"`
	BytesReturned  uint32    `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32    `idl:"name:lpcbRequired" json:"required_length"`
	Return         uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_NodeResourceControlOperation) OpNum() int { return 72 }

func (o *xxx_NodeResourceControlOperation) OpName() string {
	return "/clusapi2/v2/ApiNodeResourceControl"
}

func (o *xxx_NodeResourceControlOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InBuffer != nil && o.InBufferSize == 0 {
		o.InBufferSize = uint32(len(o.InBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeResourceControlOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node != nil {
			if err := o.Node.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		if o.InBuffer != nil || o.InBufferSize > 0 {
			_ptr_lpInBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.InBufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.InBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.InBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.InBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InBuffer, _ptr_lpInBuffer); err != nil {
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
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeResourceControlOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node == nil {
			o.Node = &Node{}
		}
		if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		_ptr_lpInBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.InBuffer", sizeInfo[0])
			}
			o.InBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.InBuffer {
				i1 := i1
				if err := w.ReadData(&o.InBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_lpInBuffer := func(ptr interface{}) { o.InBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.InBuffer, _s_lpInBuffer, _ptr_lpInBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeResourceControlOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.OutBuffer != nil && o.BytesReturned == 0 {
		o.BytesReturned = uint32(len(o.OutBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeResourceControlOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		dimSize1 := uint64(o.OutBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.BytesReturned)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.OutBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.OutBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.OutBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeResourceControlOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OutBuffer", sizeInfo[0])
		}
		o.OutBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.OutBuffer {
			i1 := i1
			if err := w.ReadData(&o.OutBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// NodeResourceControlRequest structure represents the ApiNodeResourceControl operation request
type NodeResourceControlRequest struct {
	Resource      *Resource `idl:"name:hResource" json:"resource"`
	Node          *Node     `idl:"name:hNode" json:"node"`
	ControlCode   uint32    `idl:"name:dwControlCode" json:"control_code"`
	InBuffer      []byte    `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize  uint32    `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBufferSize uint32    `idl:"name:nOutBufferSize" json:"out_buffer_size"`
}

func (o *NodeResourceControlRequest) xxx_ToOp(ctx context.Context) *xxx_NodeResourceControlOperation {
	if o == nil {
		return &xxx_NodeResourceControlOperation{}
	}
	return &xxx_NodeResourceControlOperation{
		Resource:      o.Resource,
		Node:          o.Node,
		ControlCode:   o.ControlCode,
		InBuffer:      o.InBuffer,
		InBufferSize:  o.InBufferSize,
		OutBufferSize: o.OutBufferSize,
	}
}

func (o *NodeResourceControlRequest) xxx_FromOp(ctx context.Context, op *xxx_NodeResourceControlOperation) {
	if o == nil {
		return
	}
	o.Resource = op.Resource
	o.Node = op.Node
	o.ControlCode = op.ControlCode
	o.InBuffer = op.InBuffer
	o.InBufferSize = op.InBufferSize
	o.OutBufferSize = op.OutBufferSize
}
func (o *NodeResourceControlRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *NodeResourceControlRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NodeResourceControlOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NodeResourceControlResponse structure represents the ApiNodeResourceControl operation response
type NodeResourceControlResponse struct {
	OutBuffer      []byte `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	BytesReturned  uint32 `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32 `idl:"name:lpcbRequired" json:"required_length"`
	// Return: The ApiNodeResourceControl return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *NodeResourceControlResponse) xxx_ToOp(ctx context.Context) *xxx_NodeResourceControlOperation {
	if o == nil {
		return &xxx_NodeResourceControlOperation{}
	}
	return &xxx_NodeResourceControlOperation{
		OutBuffer:      o.OutBuffer,
		BytesReturned:  o.BytesReturned,
		RequiredLength: o.RequiredLength,
		Return:         o.Return,
	}
}

func (o *NodeResourceControlResponse) xxx_FromOp(ctx context.Context, op *xxx_NodeResourceControlOperation) {
	if o == nil {
		return
	}
	o.OutBuffer = op.OutBuffer
	o.BytesReturned = op.BytesReturned
	o.RequiredLength = op.RequiredLength
	o.Return = op.Return
}
func (o *NodeResourceControlResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *NodeResourceControlResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NodeResourceControlOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ResourceControlOperation structure represents the ApiResourceControl operation
type xxx_ResourceControlOperation struct {
	Resource       *Resource `idl:"name:hResource" json:"resource"`
	ControlCode    uint32    `idl:"name:dwControlCode" json:"control_code"`
	InBuffer       []byte    `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize   uint32    `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBuffer      []byte    `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	OutBufferSize  uint32    `idl:"name:nOutBufferSize" json:"out_buffer_size"`
	BytesReturned  uint32    `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32    `idl:"name:lpcbRequired" json:"required_length"`
	Return         uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_ResourceControlOperation) OpNum() int { return 73 }

func (o *xxx_ResourceControlOperation) OpName() string { return "/clusapi2/v2/ApiResourceControl" }

func (o *xxx_ResourceControlOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InBuffer != nil && o.InBufferSize == 0 {
		o.InBufferSize = uint32(len(o.InBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResourceControlOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource != nil {
			if err := o.Resource.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Resource{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwControlCode {in} (1:(uint32))
	{
		if err := w.WriteData(o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1)[dim:0,size_is=nInBufferSize](uchar))
	{
		if o.InBuffer != nil || o.InBufferSize > 0 {
			_ptr_lpInBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.InBufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.InBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.InBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.InBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InBuffer, _ptr_lpInBuffer); err != nil {
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
	// nInBufferSize {in} (1:(uint32))
	{
		if err := w.WriteData(o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:(uint32))
	{
		if err := w.WriteData(o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResourceControlOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hResource {in} (1:{context_handle, alias=HRES_RPC, names=ndr_context_handle}(struct))
	{
		if o.Resource == nil {
			o.Resource = &Resource{}
		}
		if err := o.Resource.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwControlCode {in} (1:(uint32))
	{
		if err := w.ReadData(&o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1)[dim:0,size_is=nInBufferSize](uchar))
	{
		_ptr_lpInBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.InBuffer", sizeInfo[0])
			}
			o.InBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.InBuffer {
				i1 := i1
				if err := w.ReadData(&o.InBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_lpInBuffer := func(ptr interface{}) { o.InBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.InBuffer, _s_lpInBuffer, _ptr_lpInBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// nInBufferSize {in} (1:(uint32))
	{
		if err := w.ReadData(&o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:(uint32))
	{
		if err := w.ReadData(&o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResourceControlOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.OutBuffer != nil && o.BytesReturned == 0 {
		o.BytesReturned = uint32(len(o.OutBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResourceControlOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpOutBuffer {out} (1:{pointer=ref}*(1)[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		dimSize1 := uint64(o.OutBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.BytesReturned)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.OutBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.OutBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.OutBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResourceControlOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpOutBuffer {out} (1:{pointer=ref}*(1)[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OutBuffer", sizeInfo[0])
		}
		o.OutBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.OutBuffer {
			i1 := i1
			if err := w.ReadData(&o.OutBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ResourceControlRequest structure represents the ApiResourceControl operation request
type ResourceControlRequest struct {
	Resource      *Resource `idl:"name:hResource" json:"resource"`
	ControlCode   uint32    `idl:"name:dwControlCode" json:"control_code"`
	InBuffer      []byte    `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize  uint32    `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBufferSize uint32    `idl:"name:nOutBufferSize" json:"out_buffer_size"`
}

func (o *ResourceControlRequest) xxx_ToOp(ctx context.Context) *xxx_ResourceControlOperation {
	if o == nil {
		return &xxx_ResourceControlOperation{}
	}
	return &xxx_ResourceControlOperation{
		Resource:      o.Resource,
		ControlCode:   o.ControlCode,
		InBuffer:      o.InBuffer,
		InBufferSize:  o.InBufferSize,
		OutBufferSize: o.OutBufferSize,
	}
}

func (o *ResourceControlRequest) xxx_FromOp(ctx context.Context, op *xxx_ResourceControlOperation) {
	if o == nil {
		return
	}
	o.Resource = op.Resource
	o.ControlCode = op.ControlCode
	o.InBuffer = op.InBuffer
	o.InBufferSize = op.InBufferSize
	o.OutBufferSize = op.OutBufferSize
}
func (o *ResourceControlRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ResourceControlRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResourceControlOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ResourceControlResponse structure represents the ApiResourceControl operation response
type ResourceControlResponse struct {
	OutBuffer      []byte `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	BytesReturned  uint32 `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32 `idl:"name:lpcbRequired" json:"required_length"`
	// Return: The ApiResourceControl return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ResourceControlResponse) xxx_ToOp(ctx context.Context) *xxx_ResourceControlOperation {
	if o == nil {
		return &xxx_ResourceControlOperation{}
	}
	return &xxx_ResourceControlOperation{
		OutBuffer:      o.OutBuffer,
		BytesReturned:  o.BytesReturned,
		RequiredLength: o.RequiredLength,
		Return:         o.Return,
	}
}

func (o *ResourceControlResponse) xxx_FromOp(ctx context.Context, op *xxx_ResourceControlOperation) {
	if o == nil {
		return
	}
	o.OutBuffer = op.OutBuffer
	o.BytesReturned = op.BytesReturned
	o.RequiredLength = op.RequiredLength
	o.Return = op.Return
}
func (o *ResourceControlResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ResourceControlResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResourceControlOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_NodeResourceTypeControlOperation structure represents the ApiNodeResourceTypeControl operation
type xxx_NodeResourceTypeControlOperation struct {
	Cluster          *Cluster `idl:"name:hCluster" json:"cluster"`
	ResourceTypeName string   `idl:"name:lpszResourceTypeName;string" json:"resource_type_name"`
	Node             *Node    `idl:"name:hNode" json:"node"`
	ControlCode      uint32   `idl:"name:dwControlCode" json:"control_code"`
	InBuffer         []byte   `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize     uint32   `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBuffer        []byte   `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	OutBufferSize    uint32   `idl:"name:nOutBufferSize" json:"out_buffer_size"`
	BytesReturned    uint32   `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength   uint32   `idl:"name:lpcbRequired" json:"required_length"`
	Return           uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_NodeResourceTypeControlOperation) OpNum() int { return 74 }

func (o *xxx_NodeResourceTypeControlOperation) OpName() string {
	return "/clusapi2/v2/ApiNodeResourceTypeControl"
}

func (o *xxx_NodeResourceTypeControlOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InBuffer != nil && o.InBufferSize == 0 {
		o.InBufferSize = uint32(len(o.InBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeResourceTypeControlOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hCluster {in} (1:{context_handle, alias=HCLUSTER_RPC, names=ndr_context_handle}(struct))
	{
		if o.Cluster != nil {
			if err := o.Cluster.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Cluster{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpszResourceTypeName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ResourceTypeName); err != nil {
			return err
		}
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node != nil {
			if err := o.Node.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		if o.InBuffer != nil || o.InBufferSize > 0 {
			_ptr_lpInBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.InBufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.InBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.InBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.InBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InBuffer, _ptr_lpInBuffer); err != nil {
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
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeResourceTypeControlOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hCluster {in} (1:{context_handle, alias=HCLUSTER_RPC, names=ndr_context_handle}(struct))
	{
		if o.Cluster == nil {
			o.Cluster = &Cluster{}
		}
		if err := o.Cluster.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpszResourceTypeName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ResourceTypeName); err != nil {
			return err
		}
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node == nil {
			o.Node = &Node{}
		}
		if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		_ptr_lpInBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.InBuffer", sizeInfo[0])
			}
			o.InBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.InBuffer {
				i1 := i1
				if err := w.ReadData(&o.InBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_lpInBuffer := func(ptr interface{}) { o.InBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.InBuffer, _s_lpInBuffer, _ptr_lpInBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeResourceTypeControlOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.OutBuffer != nil && o.BytesReturned == 0 {
		o.BytesReturned = uint32(len(o.OutBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeResourceTypeControlOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		dimSize1 := uint64(o.OutBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.BytesReturned)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.OutBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.OutBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.OutBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeResourceTypeControlOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OutBuffer", sizeInfo[0])
		}
		o.OutBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.OutBuffer {
			i1 := i1
			if err := w.ReadData(&o.OutBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// NodeResourceTypeControlRequest structure represents the ApiNodeResourceTypeControl operation request
type NodeResourceTypeControlRequest struct {
	Cluster          *Cluster `idl:"name:hCluster" json:"cluster"`
	ResourceTypeName string   `idl:"name:lpszResourceTypeName;string" json:"resource_type_name"`
	Node             *Node    `idl:"name:hNode" json:"node"`
	ControlCode      uint32   `idl:"name:dwControlCode" json:"control_code"`
	InBuffer         []byte   `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize     uint32   `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBufferSize    uint32   `idl:"name:nOutBufferSize" json:"out_buffer_size"`
}

func (o *NodeResourceTypeControlRequest) xxx_ToOp(ctx context.Context) *xxx_NodeResourceTypeControlOperation {
	if o == nil {
		return &xxx_NodeResourceTypeControlOperation{}
	}
	return &xxx_NodeResourceTypeControlOperation{
		Cluster:          o.Cluster,
		ResourceTypeName: o.ResourceTypeName,
		Node:             o.Node,
		ControlCode:      o.ControlCode,
		InBuffer:         o.InBuffer,
		InBufferSize:     o.InBufferSize,
		OutBufferSize:    o.OutBufferSize,
	}
}

func (o *NodeResourceTypeControlRequest) xxx_FromOp(ctx context.Context, op *xxx_NodeResourceTypeControlOperation) {
	if o == nil {
		return
	}
	o.Cluster = op.Cluster
	o.ResourceTypeName = op.ResourceTypeName
	o.Node = op.Node
	o.ControlCode = op.ControlCode
	o.InBuffer = op.InBuffer
	o.InBufferSize = op.InBufferSize
	o.OutBufferSize = op.OutBufferSize
}
func (o *NodeResourceTypeControlRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *NodeResourceTypeControlRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NodeResourceTypeControlOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NodeResourceTypeControlResponse structure represents the ApiNodeResourceTypeControl operation response
type NodeResourceTypeControlResponse struct {
	OutBuffer      []byte `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	BytesReturned  uint32 `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32 `idl:"name:lpcbRequired" json:"required_length"`
	// Return: The ApiNodeResourceTypeControl return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *NodeResourceTypeControlResponse) xxx_ToOp(ctx context.Context) *xxx_NodeResourceTypeControlOperation {
	if o == nil {
		return &xxx_NodeResourceTypeControlOperation{}
	}
	return &xxx_NodeResourceTypeControlOperation{
		OutBuffer:      o.OutBuffer,
		BytesReturned:  o.BytesReturned,
		RequiredLength: o.RequiredLength,
		Return:         o.Return,
	}
}

func (o *NodeResourceTypeControlResponse) xxx_FromOp(ctx context.Context, op *xxx_NodeResourceTypeControlOperation) {
	if o == nil {
		return
	}
	o.OutBuffer = op.OutBuffer
	o.BytesReturned = op.BytesReturned
	o.RequiredLength = op.RequiredLength
	o.Return = op.Return
}
func (o *NodeResourceTypeControlResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *NodeResourceTypeControlResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NodeResourceTypeControlOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ResourceTypeControlOperation structure represents the ApiResourceTypeControl operation
type xxx_ResourceTypeControlOperation struct {
	Cluster          *Cluster `idl:"name:hCluster" json:"cluster"`
	ResourceTypeName string   `idl:"name:lpszResourceTypeName;string" json:"resource_type_name"`
	ControlCode      uint32   `idl:"name:dwControlCode" json:"control_code"`
	InBuffer         []byte   `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize     uint32   `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBuffer        []byte   `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	OutBufferSize    uint32   `idl:"name:nOutBufferSize" json:"out_buffer_size"`
	BytesReturned    uint32   `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength   uint32   `idl:"name:lpcbRequired" json:"required_length"`
	Return           uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_ResourceTypeControlOperation) OpNum() int { return 75 }

func (o *xxx_ResourceTypeControlOperation) OpName() string {
	return "/clusapi2/v2/ApiResourceTypeControl"
}

func (o *xxx_ResourceTypeControlOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InBuffer != nil && o.InBufferSize == 0 {
		o.InBufferSize = uint32(len(o.InBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResourceTypeControlOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hCluster {in} (1:{context_handle, alias=HCLUSTER_RPC, names=ndr_context_handle}(struct))
	{
		if o.Cluster != nil {
			if err := o.Cluster.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Cluster{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpszResourceTypeName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ResourceTypeName); err != nil {
			return err
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		if o.InBuffer != nil || o.InBufferSize > 0 {
			_ptr_lpInBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.InBufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.InBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.InBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.InBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InBuffer, _ptr_lpInBuffer); err != nil {
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
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResourceTypeControlOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hCluster {in} (1:{context_handle, alias=HCLUSTER_RPC, names=ndr_context_handle}(struct))
	{
		if o.Cluster == nil {
			o.Cluster = &Cluster{}
		}
		if err := o.Cluster.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpszResourceTypeName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ResourceTypeName); err != nil {
			return err
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		_ptr_lpInBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.InBuffer", sizeInfo[0])
			}
			o.InBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.InBuffer {
				i1 := i1
				if err := w.ReadData(&o.InBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_lpInBuffer := func(ptr interface{}) { o.InBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.InBuffer, _s_lpInBuffer, _ptr_lpInBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResourceTypeControlOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.OutBuffer != nil && o.BytesReturned == 0 {
		o.BytesReturned = uint32(len(o.OutBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResourceTypeControlOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		dimSize1 := uint64(o.OutBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.BytesReturned)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.OutBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.OutBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.OutBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResourceTypeControlOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OutBuffer", sizeInfo[0])
		}
		o.OutBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.OutBuffer {
			i1 := i1
			if err := w.ReadData(&o.OutBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ResourceTypeControlRequest structure represents the ApiResourceTypeControl operation request
type ResourceTypeControlRequest struct {
	Cluster          *Cluster `idl:"name:hCluster" json:"cluster"`
	ResourceTypeName string   `idl:"name:lpszResourceTypeName;string" json:"resource_type_name"`
	ControlCode      uint32   `idl:"name:dwControlCode" json:"control_code"`
	InBuffer         []byte   `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize     uint32   `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBufferSize    uint32   `idl:"name:nOutBufferSize" json:"out_buffer_size"`
}

func (o *ResourceTypeControlRequest) xxx_ToOp(ctx context.Context) *xxx_ResourceTypeControlOperation {
	if o == nil {
		return &xxx_ResourceTypeControlOperation{}
	}
	return &xxx_ResourceTypeControlOperation{
		Cluster:          o.Cluster,
		ResourceTypeName: o.ResourceTypeName,
		ControlCode:      o.ControlCode,
		InBuffer:         o.InBuffer,
		InBufferSize:     o.InBufferSize,
		OutBufferSize:    o.OutBufferSize,
	}
}

func (o *ResourceTypeControlRequest) xxx_FromOp(ctx context.Context, op *xxx_ResourceTypeControlOperation) {
	if o == nil {
		return
	}
	o.Cluster = op.Cluster
	o.ResourceTypeName = op.ResourceTypeName
	o.ControlCode = op.ControlCode
	o.InBuffer = op.InBuffer
	o.InBufferSize = op.InBufferSize
	o.OutBufferSize = op.OutBufferSize
}
func (o *ResourceTypeControlRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ResourceTypeControlRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResourceTypeControlOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ResourceTypeControlResponse structure represents the ApiResourceTypeControl operation response
type ResourceTypeControlResponse struct {
	OutBuffer      []byte `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	BytesReturned  uint32 `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32 `idl:"name:lpcbRequired" json:"required_length"`
	// Return: The ApiResourceTypeControl return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ResourceTypeControlResponse) xxx_ToOp(ctx context.Context) *xxx_ResourceTypeControlOperation {
	if o == nil {
		return &xxx_ResourceTypeControlOperation{}
	}
	return &xxx_ResourceTypeControlOperation{
		OutBuffer:      o.OutBuffer,
		BytesReturned:  o.BytesReturned,
		RequiredLength: o.RequiredLength,
		Return:         o.Return,
	}
}

func (o *ResourceTypeControlResponse) xxx_FromOp(ctx context.Context, op *xxx_ResourceTypeControlOperation) {
	if o == nil {
		return
	}
	o.OutBuffer = op.OutBuffer
	o.BytesReturned = op.BytesReturned
	o.RequiredLength = op.RequiredLength
	o.Return = op.Return
}
func (o *ResourceTypeControlResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ResourceTypeControlResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResourceTypeControlOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_NodeGroupControlOperation structure represents the ApiNodeGroupControl operation
type xxx_NodeGroupControlOperation struct {
	Group          *Group `idl:"name:hGroup" json:"group"`
	Node           *Node  `idl:"name:hNode" json:"node"`
	ControlCode    uint32 `idl:"name:dwControlCode" json:"control_code"`
	InBuffer       []byte `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize   uint32 `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBuffer      []byte `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	OutBufferSize  uint32 `idl:"name:nOutBufferSize" json:"out_buffer_size"`
	BytesReturned  uint32 `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32 `idl:"name:lpcbRequired" json:"required_length"`
	Return         uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_NodeGroupControlOperation) OpNum() int { return 76 }

func (o *xxx_NodeGroupControlOperation) OpName() string { return "/clusapi2/v2/ApiNodeGroupControl" }

func (o *xxx_NodeGroupControlOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InBuffer != nil && o.InBufferSize == 0 {
		o.InBufferSize = uint32(len(o.InBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeGroupControlOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group != nil {
			if err := o.Group.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Group{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node != nil {
			if err := o.Node.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		if o.InBuffer != nil || o.InBufferSize > 0 {
			_ptr_lpInBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.InBufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.InBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.InBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.InBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InBuffer, _ptr_lpInBuffer); err != nil {
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
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeGroupControlOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group == nil {
			o.Group = &Group{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node == nil {
			o.Node = &Node{}
		}
		if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		_ptr_lpInBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.InBuffer", sizeInfo[0])
			}
			o.InBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.InBuffer {
				i1 := i1
				if err := w.ReadData(&o.InBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_lpInBuffer := func(ptr interface{}) { o.InBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.InBuffer, _s_lpInBuffer, _ptr_lpInBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeGroupControlOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.OutBuffer != nil && o.BytesReturned == 0 {
		o.BytesReturned = uint32(len(o.OutBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeGroupControlOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		dimSize1 := uint64(o.OutBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.BytesReturned)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.OutBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.OutBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.OutBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeGroupControlOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OutBuffer", sizeInfo[0])
		}
		o.OutBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.OutBuffer {
			i1 := i1
			if err := w.ReadData(&o.OutBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// NodeGroupControlRequest structure represents the ApiNodeGroupControl operation request
type NodeGroupControlRequest struct {
	Group         *Group `idl:"name:hGroup" json:"group"`
	Node          *Node  `idl:"name:hNode" json:"node"`
	ControlCode   uint32 `idl:"name:dwControlCode" json:"control_code"`
	InBuffer      []byte `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize  uint32 `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBufferSize uint32 `idl:"name:nOutBufferSize" json:"out_buffer_size"`
}

func (o *NodeGroupControlRequest) xxx_ToOp(ctx context.Context) *xxx_NodeGroupControlOperation {
	if o == nil {
		return &xxx_NodeGroupControlOperation{}
	}
	return &xxx_NodeGroupControlOperation{
		Group:         o.Group,
		Node:          o.Node,
		ControlCode:   o.ControlCode,
		InBuffer:      o.InBuffer,
		InBufferSize:  o.InBufferSize,
		OutBufferSize: o.OutBufferSize,
	}
}

func (o *NodeGroupControlRequest) xxx_FromOp(ctx context.Context, op *xxx_NodeGroupControlOperation) {
	if o == nil {
		return
	}
	o.Group = op.Group
	o.Node = op.Node
	o.ControlCode = op.ControlCode
	o.InBuffer = op.InBuffer
	o.InBufferSize = op.InBufferSize
	o.OutBufferSize = op.OutBufferSize
}
func (o *NodeGroupControlRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *NodeGroupControlRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NodeGroupControlOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NodeGroupControlResponse structure represents the ApiNodeGroupControl operation response
type NodeGroupControlResponse struct {
	OutBuffer      []byte `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	BytesReturned  uint32 `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32 `idl:"name:lpcbRequired" json:"required_length"`
	// Return: The ApiNodeGroupControl return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *NodeGroupControlResponse) xxx_ToOp(ctx context.Context) *xxx_NodeGroupControlOperation {
	if o == nil {
		return &xxx_NodeGroupControlOperation{}
	}
	return &xxx_NodeGroupControlOperation{
		OutBuffer:      o.OutBuffer,
		BytesReturned:  o.BytesReturned,
		RequiredLength: o.RequiredLength,
		Return:         o.Return,
	}
}

func (o *NodeGroupControlResponse) xxx_FromOp(ctx context.Context, op *xxx_NodeGroupControlOperation) {
	if o == nil {
		return
	}
	o.OutBuffer = op.OutBuffer
	o.BytesReturned = op.BytesReturned
	o.RequiredLength = op.RequiredLength
	o.Return = op.Return
}
func (o *NodeGroupControlResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *NodeGroupControlResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NodeGroupControlOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GroupControlOperation structure represents the ApiGroupControl operation
type xxx_GroupControlOperation struct {
	Group          *Group `idl:"name:hGroup" json:"group"`
	ControlCode    uint32 `idl:"name:dwControlCode" json:"control_code"`
	InBuffer       []byte `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize   uint32 `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBuffer      []byte `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	OutBufferSize  uint32 `idl:"name:nOutBufferSize" json:"out_buffer_size"`
	BytesReturned  uint32 `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32 `idl:"name:lpcbRequired" json:"required_length"`
	Return         uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GroupControlOperation) OpNum() int { return 77 }

func (o *xxx_GroupControlOperation) OpName() string { return "/clusapi2/v2/ApiGroupControl" }

func (o *xxx_GroupControlOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InBuffer != nil && o.InBufferSize == 0 {
		o.InBufferSize = uint32(len(o.InBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GroupControlOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group != nil {
			if err := o.Group.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Group{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		if o.InBuffer != nil || o.InBufferSize > 0 {
			_ptr_lpInBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.InBufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.InBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.InBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.InBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InBuffer, _ptr_lpInBuffer); err != nil {
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
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GroupControlOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hGroup {in} (1:{context_handle, alias=HGROUP_RPC, names=ndr_context_handle}(struct))
	{
		if o.Group == nil {
			o.Group = &Group{}
		}
		if err := o.Group.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		_ptr_lpInBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.InBuffer", sizeInfo[0])
			}
			o.InBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.InBuffer {
				i1 := i1
				if err := w.ReadData(&o.InBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_lpInBuffer := func(ptr interface{}) { o.InBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.InBuffer, _s_lpInBuffer, _ptr_lpInBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GroupControlOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.OutBuffer != nil && o.BytesReturned == 0 {
		o.BytesReturned = uint32(len(o.OutBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GroupControlOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		dimSize1 := uint64(o.OutBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.BytesReturned)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.OutBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.OutBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.OutBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GroupControlOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OutBuffer", sizeInfo[0])
		}
		o.OutBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.OutBuffer {
			i1 := i1
			if err := w.ReadData(&o.OutBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GroupControlRequest structure represents the ApiGroupControl operation request
type GroupControlRequest struct {
	Group         *Group `idl:"name:hGroup" json:"group"`
	ControlCode   uint32 `idl:"name:dwControlCode" json:"control_code"`
	InBuffer      []byte `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize  uint32 `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBufferSize uint32 `idl:"name:nOutBufferSize" json:"out_buffer_size"`
}

func (o *GroupControlRequest) xxx_ToOp(ctx context.Context) *xxx_GroupControlOperation {
	if o == nil {
		return &xxx_GroupControlOperation{}
	}
	return &xxx_GroupControlOperation{
		Group:         o.Group,
		ControlCode:   o.ControlCode,
		InBuffer:      o.InBuffer,
		InBufferSize:  o.InBufferSize,
		OutBufferSize: o.OutBufferSize,
	}
}

func (o *GroupControlRequest) xxx_FromOp(ctx context.Context, op *xxx_GroupControlOperation) {
	if o == nil {
		return
	}
	o.Group = op.Group
	o.ControlCode = op.ControlCode
	o.InBuffer = op.InBuffer
	o.InBufferSize = op.InBufferSize
	o.OutBufferSize = op.OutBufferSize
}
func (o *GroupControlRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GroupControlRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GroupControlOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GroupControlResponse structure represents the ApiGroupControl operation response
type GroupControlResponse struct {
	OutBuffer      []byte `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	BytesReturned  uint32 `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32 `idl:"name:lpcbRequired" json:"required_length"`
	// Return: The ApiGroupControl return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GroupControlResponse) xxx_ToOp(ctx context.Context) *xxx_GroupControlOperation {
	if o == nil {
		return &xxx_GroupControlOperation{}
	}
	return &xxx_GroupControlOperation{
		OutBuffer:      o.OutBuffer,
		BytesReturned:  o.BytesReturned,
		RequiredLength: o.RequiredLength,
		Return:         o.Return,
	}
}

func (o *GroupControlResponse) xxx_FromOp(ctx context.Context, op *xxx_GroupControlOperation) {
	if o == nil {
		return
	}
	o.OutBuffer = op.OutBuffer
	o.BytesReturned = op.BytesReturned
	o.RequiredLength = op.RequiredLength
	o.Return = op.Return
}
func (o *GroupControlResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GroupControlResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GroupControlOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_NodeNodeControlOperation structure represents the ApiNodeNodeControl operation
type xxx_NodeNodeControlOperation struct {
	Node           *Node  `idl:"name:hNode" json:"node"`
	HostNode       *Node  `idl:"name:hHostNode" json:"host_node"`
	ControlCode    uint32 `idl:"name:dwControlCode" json:"control_code"`
	InBuffer       []byte `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize   uint32 `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBuffer      []byte `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	OutBufferSize  uint32 `idl:"name:nOutBufferSize" json:"out_buffer_size"`
	BytesReturned  uint32 `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32 `idl:"name:lpcbRequired" json:"required_length"`
	Return         uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_NodeNodeControlOperation) OpNum() int { return 78 }

func (o *xxx_NodeNodeControlOperation) OpName() string { return "/clusapi2/v2/ApiNodeNodeControl" }

func (o *xxx_NodeNodeControlOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InBuffer != nil && o.InBufferSize == 0 {
		o.InBufferSize = uint32(len(o.InBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeNodeControlOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node != nil {
			if err := o.Node.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hHostNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.HostNode != nil {
			if err := o.HostNode.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		if o.InBuffer != nil || o.InBufferSize > 0 {
			_ptr_lpInBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.InBufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.InBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.InBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.InBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InBuffer, _ptr_lpInBuffer); err != nil {
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
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeNodeControlOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node == nil {
			o.Node = &Node{}
		}
		if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hHostNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.HostNode == nil {
			o.HostNode = &Node{}
		}
		if err := o.HostNode.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		_ptr_lpInBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.InBuffer", sizeInfo[0])
			}
			o.InBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.InBuffer {
				i1 := i1
				if err := w.ReadData(&o.InBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_lpInBuffer := func(ptr interface{}) { o.InBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.InBuffer, _s_lpInBuffer, _ptr_lpInBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeNodeControlOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.OutBuffer != nil && o.BytesReturned == 0 {
		o.BytesReturned = uint32(len(o.OutBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeNodeControlOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		dimSize1 := uint64(o.OutBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.BytesReturned)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.OutBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.OutBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.OutBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeNodeControlOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OutBuffer", sizeInfo[0])
		}
		o.OutBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.OutBuffer {
			i1 := i1
			if err := w.ReadData(&o.OutBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// NodeNodeControlRequest structure represents the ApiNodeNodeControl operation request
type NodeNodeControlRequest struct {
	Node          *Node  `idl:"name:hNode" json:"node"`
	HostNode      *Node  `idl:"name:hHostNode" json:"host_node"`
	ControlCode   uint32 `idl:"name:dwControlCode" json:"control_code"`
	InBuffer      []byte `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize  uint32 `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBufferSize uint32 `idl:"name:nOutBufferSize" json:"out_buffer_size"`
}

func (o *NodeNodeControlRequest) xxx_ToOp(ctx context.Context) *xxx_NodeNodeControlOperation {
	if o == nil {
		return &xxx_NodeNodeControlOperation{}
	}
	return &xxx_NodeNodeControlOperation{
		Node:          o.Node,
		HostNode:      o.HostNode,
		ControlCode:   o.ControlCode,
		InBuffer:      o.InBuffer,
		InBufferSize:  o.InBufferSize,
		OutBufferSize: o.OutBufferSize,
	}
}

func (o *NodeNodeControlRequest) xxx_FromOp(ctx context.Context, op *xxx_NodeNodeControlOperation) {
	if o == nil {
		return
	}
	o.Node = op.Node
	o.HostNode = op.HostNode
	o.ControlCode = op.ControlCode
	o.InBuffer = op.InBuffer
	o.InBufferSize = op.InBufferSize
	o.OutBufferSize = op.OutBufferSize
}
func (o *NodeNodeControlRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *NodeNodeControlRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NodeNodeControlOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NodeNodeControlResponse structure represents the ApiNodeNodeControl operation response
type NodeNodeControlResponse struct {
	OutBuffer      []byte `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	BytesReturned  uint32 `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32 `idl:"name:lpcbRequired" json:"required_length"`
	// Return: The ApiNodeNodeControl return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *NodeNodeControlResponse) xxx_ToOp(ctx context.Context) *xxx_NodeNodeControlOperation {
	if o == nil {
		return &xxx_NodeNodeControlOperation{}
	}
	return &xxx_NodeNodeControlOperation{
		OutBuffer:      o.OutBuffer,
		BytesReturned:  o.BytesReturned,
		RequiredLength: o.RequiredLength,
		Return:         o.Return,
	}
}

func (o *NodeNodeControlResponse) xxx_FromOp(ctx context.Context, op *xxx_NodeNodeControlOperation) {
	if o == nil {
		return
	}
	o.OutBuffer = op.OutBuffer
	o.BytesReturned = op.BytesReturned
	o.RequiredLength = op.RequiredLength
	o.Return = op.Return
}
func (o *NodeNodeControlResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *NodeNodeControlResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NodeNodeControlOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_NodeControlOperation structure represents the ApiNodeControl operation
type xxx_NodeControlOperation struct {
	Node           *Node  `idl:"name:hNode" json:"node"`
	ControlCode    uint32 `idl:"name:dwControlCode" json:"control_code"`
	InBuffer       []byte `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize   uint32 `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBuffer      []byte `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	OutBufferSize  uint32 `idl:"name:nOutBufferSize" json:"out_buffer_size"`
	BytesReturned  uint32 `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32 `idl:"name:lpcbRequired" json:"required_length"`
	Return         uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_NodeControlOperation) OpNum() int { return 79 }

func (o *xxx_NodeControlOperation) OpName() string { return "/clusapi2/v2/ApiNodeControl" }

func (o *xxx_NodeControlOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InBuffer != nil && o.InBufferSize == 0 {
		o.InBufferSize = uint32(len(o.InBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeControlOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node != nil {
			if err := o.Node.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		if o.InBuffer != nil || o.InBufferSize > 0 {
			_ptr_lpInBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.InBufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.InBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.InBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.InBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InBuffer, _ptr_lpInBuffer); err != nil {
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
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeControlOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node == nil {
			o.Node = &Node{}
		}
		if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		_ptr_lpInBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.InBuffer", sizeInfo[0])
			}
			o.InBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.InBuffer {
				i1 := i1
				if err := w.ReadData(&o.InBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_lpInBuffer := func(ptr interface{}) { o.InBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.InBuffer, _s_lpInBuffer, _ptr_lpInBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeControlOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.OutBuffer != nil && o.BytesReturned == 0 {
		o.BytesReturned = uint32(len(o.OutBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeControlOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		dimSize1 := uint64(o.OutBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.BytesReturned)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.OutBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.OutBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.OutBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeControlOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OutBuffer", sizeInfo[0])
		}
		o.OutBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.OutBuffer {
			i1 := i1
			if err := w.ReadData(&o.OutBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// NodeControlRequest structure represents the ApiNodeControl operation request
type NodeControlRequest struct {
	Node          *Node  `idl:"name:hNode" json:"node"`
	ControlCode   uint32 `idl:"name:dwControlCode" json:"control_code"`
	InBuffer      []byte `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize  uint32 `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBufferSize uint32 `idl:"name:nOutBufferSize" json:"out_buffer_size"`
}

func (o *NodeControlRequest) xxx_ToOp(ctx context.Context) *xxx_NodeControlOperation {
	if o == nil {
		return &xxx_NodeControlOperation{}
	}
	return &xxx_NodeControlOperation{
		Node:          o.Node,
		ControlCode:   o.ControlCode,
		InBuffer:      o.InBuffer,
		InBufferSize:  o.InBufferSize,
		OutBufferSize: o.OutBufferSize,
	}
}

func (o *NodeControlRequest) xxx_FromOp(ctx context.Context, op *xxx_NodeControlOperation) {
	if o == nil {
		return
	}
	o.Node = op.Node
	o.ControlCode = op.ControlCode
	o.InBuffer = op.InBuffer
	o.InBufferSize = op.InBufferSize
	o.OutBufferSize = op.OutBufferSize
}
func (o *NodeControlRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *NodeControlRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NodeControlOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NodeControlResponse structure represents the ApiNodeControl operation response
type NodeControlResponse struct {
	OutBuffer      []byte `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	BytesReturned  uint32 `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32 `idl:"name:lpcbRequired" json:"required_length"`
	// Return: The ApiNodeControl return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *NodeControlResponse) xxx_ToOp(ctx context.Context) *xxx_NodeControlOperation {
	if o == nil {
		return &xxx_NodeControlOperation{}
	}
	return &xxx_NodeControlOperation{
		OutBuffer:      o.OutBuffer,
		BytesReturned:  o.BytesReturned,
		RequiredLength: o.RequiredLength,
		Return:         o.Return,
	}
}

func (o *NodeControlResponse) xxx_FromOp(ctx context.Context, op *xxx_NodeControlOperation) {
	if o == nil {
		return
	}
	o.OutBuffer = op.OutBuffer
	o.BytesReturned = op.BytesReturned
	o.RequiredLength = op.RequiredLength
	o.Return = op.Return
}
func (o *NodeControlResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *NodeControlResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NodeControlOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenNetworkOperation structure represents the ApiOpenNetwork operation
type xxx_OpenNetworkOperation struct {
	NetworkName string   `idl:"name:lpszNetworkName;string" json:"network_name"`
	Status      uint32   `idl:"name:Status" json:"status"`
	Return      *Network `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenNetworkOperation) OpNum() int { return 81 }

func (o *xxx_OpenNetworkOperation) OpName() string { return "/clusapi2/v2/ApiOpenNetwork" }

func (o *xxx_OpenNetworkOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNetworkOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpszNetworkName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.NetworkName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNetworkOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpszNetworkName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.NetworkName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNetworkOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNetworkOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return != nil {
			if err := o.Return.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Network{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_OpenNetworkOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return == nil {
			o.Return = &Network{}
		}
		if err := o.Return.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// OpenNetworkRequest structure represents the ApiOpenNetwork operation request
type OpenNetworkRequest struct {
	NetworkName string `idl:"name:lpszNetworkName;string" json:"network_name"`
}

func (o *OpenNetworkRequest) xxx_ToOp(ctx context.Context) *xxx_OpenNetworkOperation {
	if o == nil {
		return &xxx_OpenNetworkOperation{}
	}
	return &xxx_OpenNetworkOperation{
		NetworkName: o.NetworkName,
	}
}

func (o *OpenNetworkRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenNetworkOperation) {
	if o == nil {
		return
	}
	o.NetworkName = op.NetworkName
}
func (o *OpenNetworkRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OpenNetworkRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenNetworkOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenNetworkResponse structure represents the ApiOpenNetwork operation response
type OpenNetworkResponse struct {
	Status uint32 `idl:"name:Status" json:"status"`
	// Return: The ApiOpenNetwork return value.
	Return *Network `idl:"name:Return" json:"return"`
}

func (o *OpenNetworkResponse) xxx_ToOp(ctx context.Context) *xxx_OpenNetworkOperation {
	if o == nil {
		return &xxx_OpenNetworkOperation{}
	}
	return &xxx_OpenNetworkOperation{
		Status: o.Status,
		Return: o.Return,
	}
}

func (o *OpenNetworkResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenNetworkOperation) {
	if o == nil {
		return
	}
	o.Status = op.Status
	o.Return = op.Return
}
func (o *OpenNetworkResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OpenNetworkResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenNetworkOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseNetworkOperation structure represents the ApiCloseNetwork operation
type xxx_CloseNetworkOperation struct {
	Network *Network `idl:"name:Network" json:"network"`
	Return  uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseNetworkOperation) OpNum() int { return 82 }

func (o *xxx_CloseNetworkOperation) OpName() string { return "/clusapi2/v2/ApiCloseNetwork" }

func (o *xxx_CloseNetworkOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseNetworkOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Network {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Network != nil {
			if err := o.Network.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Network{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseNetworkOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Network {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Network == nil {
			o.Network = &Network{}
		}
		if err := o.Network.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseNetworkOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseNetworkOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Network {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Network != nil {
			if err := o.Network.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Network{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseNetworkOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Network {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Network == nil {
			o.Network = &Network{}
		}
		if err := o.Network.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloseNetworkRequest structure represents the ApiCloseNetwork operation request
type CloseNetworkRequest struct {
	Network *Network `idl:"name:Network" json:"network"`
}

func (o *CloseNetworkRequest) xxx_ToOp(ctx context.Context) *xxx_CloseNetworkOperation {
	if o == nil {
		return &xxx_CloseNetworkOperation{}
	}
	return &xxx_CloseNetworkOperation{
		Network: o.Network,
	}
}

func (o *CloseNetworkRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseNetworkOperation) {
	if o == nil {
		return
	}
	o.Network = op.Network
}
func (o *CloseNetworkRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CloseNetworkRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseNetworkOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseNetworkResponse structure represents the ApiCloseNetwork operation response
type CloseNetworkResponse struct {
	Network *Network `idl:"name:Network" json:"network"`
	// Return: The ApiCloseNetwork return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CloseNetworkResponse) xxx_ToOp(ctx context.Context) *xxx_CloseNetworkOperation {
	if o == nil {
		return &xxx_CloseNetworkOperation{}
	}
	return &xxx_CloseNetworkOperation{
		Network: o.Network,
		Return:  o.Return,
	}
}

func (o *CloseNetworkResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseNetworkOperation) {
	if o == nil {
		return
	}
	o.Network = op.Network
	o.Return = op.Return
}
func (o *CloseNetworkResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CloseNetworkResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseNetworkOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNetworkStateOperation structure represents the ApiGetNetworkState operation
type xxx_GetNetworkStateOperation struct {
	Network *Network `idl:"name:hNetwork" json:"network"`
	State   uint32   `idl:"name:State" json:"state"`
	Return  uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNetworkStateOperation) OpNum() int { return 83 }

func (o *xxx_GetNetworkStateOperation) OpName() string { return "/clusapi2/v2/ApiGetNetworkState" }

func (o *xxx_GetNetworkStateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetworkStateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNetwork {in} (1:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Network != nil {
			if err := o.Network.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Network{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetNetworkStateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNetwork {in} (1:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Network == nil {
			o.Network = &Network{}
		}
		if err := o.Network.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetworkStateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetworkStateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// State {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.State); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetworkStateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// State {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.State); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNetworkStateRequest structure represents the ApiGetNetworkState operation request
type GetNetworkStateRequest struct {
	Network *Network `idl:"name:hNetwork" json:"network"`
}

func (o *GetNetworkStateRequest) xxx_ToOp(ctx context.Context) *xxx_GetNetworkStateOperation {
	if o == nil {
		return &xxx_GetNetworkStateOperation{}
	}
	return &xxx_GetNetworkStateOperation{
		Network: o.Network,
	}
}

func (o *GetNetworkStateRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNetworkStateOperation) {
	if o == nil {
		return
	}
	o.Network = op.Network
}
func (o *GetNetworkStateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetNetworkStateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNetworkStateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNetworkStateResponse structure represents the ApiGetNetworkState operation response
type GetNetworkStateResponse struct {
	State uint32 `idl:"name:State" json:"state"`
	// Return: The ApiGetNetworkState return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetNetworkStateResponse) xxx_ToOp(ctx context.Context) *xxx_GetNetworkStateOperation {
	if o == nil {
		return &xxx_GetNetworkStateOperation{}
	}
	return &xxx_GetNetworkStateOperation{
		State:  o.State,
		Return: o.Return,
	}
}

func (o *GetNetworkStateResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNetworkStateOperation) {
	if o == nil {
		return
	}
	o.State = op.State
	o.Return = op.Return
}
func (o *GetNetworkStateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetNetworkStateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNetworkStateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetNetworkNameOperation structure represents the ApiSetNetworkName operation
type xxx_SetNetworkNameOperation struct {
	Network     *Network `idl:"name:hNetwork" json:"network"`
	NetworkName string   `idl:"name:lpszNetworkName;string" json:"network_name"`
	Return      uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNetworkNameOperation) OpNum() int { return 84 }

func (o *xxx_SetNetworkNameOperation) OpName() string { return "/clusapi2/v2/ApiSetNetworkName" }

func (o *xxx_SetNetworkNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNetworkNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNetwork {in} (1:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Network != nil {
			if err := o.Network.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Network{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpszNetworkName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.NetworkName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNetworkNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNetwork {in} (1:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Network == nil {
			o.Network = &Network{}
		}
		if err := o.Network.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpszNetworkName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.NetworkName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNetworkNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNetworkNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNetworkNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetNetworkNameRequest structure represents the ApiSetNetworkName operation request
type SetNetworkNameRequest struct {
	Network     *Network `idl:"name:hNetwork" json:"network"`
	NetworkName string   `idl:"name:lpszNetworkName;string" json:"network_name"`
}

func (o *SetNetworkNameRequest) xxx_ToOp(ctx context.Context) *xxx_SetNetworkNameOperation {
	if o == nil {
		return &xxx_SetNetworkNameOperation{}
	}
	return &xxx_SetNetworkNameOperation{
		Network:     o.Network,
		NetworkName: o.NetworkName,
	}
}

func (o *SetNetworkNameRequest) xxx_FromOp(ctx context.Context, op *xxx_SetNetworkNameOperation) {
	if o == nil {
		return
	}
	o.Network = op.Network
	o.NetworkName = op.NetworkName
}
func (o *SetNetworkNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetNetworkNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNetworkNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNetworkNameResponse structure represents the ApiSetNetworkName operation response
type SetNetworkNameResponse struct {
	// Return: The ApiSetNetworkName return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetNetworkNameResponse) xxx_ToOp(ctx context.Context) *xxx_SetNetworkNameOperation {
	if o == nil {
		return &xxx_SetNetworkNameOperation{}
	}
	return &xxx_SetNetworkNameOperation{
		Return: o.Return,
	}
}

func (o *SetNetworkNameResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNetworkNameOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetNetworkNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetNetworkNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNetworkNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateNetworkEnumOperation structure represents the ApiCreateNetworkEnum operation
type xxx_CreateNetworkEnumOperation struct {
	Network    *Network  `idl:"name:hNetwork" json:"network"`
	Type       uint32    `idl:"name:dwType" json:"type"`
	ReturnEnum *EnumList `idl:"name:ReturnEnum" json:"return_enum"`
	Return     uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateNetworkEnumOperation) OpNum() int { return 85 }

func (o *xxx_CreateNetworkEnumOperation) OpName() string { return "/clusapi2/v2/ApiCreateNetworkEnum" }

func (o *xxx_CreateNetworkEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNetworkEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNetwork {in} (1:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Network != nil {
			if err := o.Network.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Network{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNetworkEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNetwork {in} (1:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Network == nil {
			o.Network = &Network{}
		}
		if err := o.Network.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNetworkEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNetworkEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ReturnEnum {out} (1:{pointer=ref}*(2))(2:{alias=PENUM_LIST}*(1))(3:{alias=ENUM_LIST}(struct))
	{
		if o.ReturnEnum != nil {
			_ptr_ReturnEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnEnum != nil {
					if err := o.ReturnEnum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EnumList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnEnum, _ptr_ReturnEnum); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNetworkEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ReturnEnum {out} (1:{pointer=ref}*(2))(2:{alias=PENUM_LIST,pointer=ref}*(1))(3:{alias=ENUM_LIST}(struct))
	{
		_ptr_ReturnEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnEnum == nil {
				o.ReturnEnum = &EnumList{}
			}
			if err := o.ReturnEnum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReturnEnum := func(ptr interface{}) { o.ReturnEnum = *ptr.(**EnumList) }
		if err := w.ReadPointer(&o.ReturnEnum, _s_ReturnEnum, _ptr_ReturnEnum); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateNetworkEnumRequest structure represents the ApiCreateNetworkEnum operation request
type CreateNetworkEnumRequest struct {
	Network *Network `idl:"name:hNetwork" json:"network"`
	Type    uint32   `idl:"name:dwType" json:"type"`
}

func (o *CreateNetworkEnumRequest) xxx_ToOp(ctx context.Context) *xxx_CreateNetworkEnumOperation {
	if o == nil {
		return &xxx_CreateNetworkEnumOperation{}
	}
	return &xxx_CreateNetworkEnumOperation{
		Network: o.Network,
		Type:    o.Type,
	}
}

func (o *CreateNetworkEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateNetworkEnumOperation) {
	if o == nil {
		return
	}
	o.Network = op.Network
	o.Type = op.Type
}
func (o *CreateNetworkEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateNetworkEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateNetworkEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateNetworkEnumResponse structure represents the ApiCreateNetworkEnum operation response
type CreateNetworkEnumResponse struct {
	ReturnEnum *EnumList `idl:"name:ReturnEnum" json:"return_enum"`
	// Return: The ApiCreateNetworkEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CreateNetworkEnumResponse) xxx_ToOp(ctx context.Context) *xxx_CreateNetworkEnumOperation {
	if o == nil {
		return &xxx_CreateNetworkEnumOperation{}
	}
	return &xxx_CreateNetworkEnumOperation{
		ReturnEnum: o.ReturnEnum,
		Return:     o.Return,
	}
}

func (o *CreateNetworkEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateNetworkEnumOperation) {
	if o == nil {
		return
	}
	o.ReturnEnum = op.ReturnEnum
	o.Return = op.Return
}
func (o *CreateNetworkEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateNetworkEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateNetworkEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNetworkIDOperation structure represents the ApiGetNetworkId operation
type xxx_GetNetworkIDOperation struct {
	Network *Network `idl:"name:hNetwork" json:"network"`
	GUID    string   `idl:"name:pGuid;string" json:"guid"`
	Return  uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNetworkIDOperation) OpNum() int { return 86 }

func (o *xxx_GetNetworkIDOperation) OpName() string { return "/clusapi2/v2/ApiGetNetworkId" }

func (o *xxx_GetNetworkIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetworkIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNetwork {in} (1:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Network != nil {
			if err := o.Network.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Network{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetNetworkIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNetwork {in} (1:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Network == nil {
			o.Network = &Network{}
		}
		if err := o.Network.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetworkIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetworkIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pGuid {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.GUID != "" {
			_ptr_pGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.GUID); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.GUID, _ptr_pGuid); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetworkIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pGuid {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.GUID); err != nil {
				return err
			}
			return nil
		})
		_s_pGuid := func(ptr interface{}) { o.GUID = *ptr.(*string) }
		if err := w.ReadPointer(&o.GUID, _s_pGuid, _ptr_pGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNetworkIDRequest structure represents the ApiGetNetworkId operation request
type GetNetworkIDRequest struct {
	Network *Network `idl:"name:hNetwork" json:"network"`
}

func (o *GetNetworkIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetNetworkIDOperation {
	if o == nil {
		return &xxx_GetNetworkIDOperation{}
	}
	return &xxx_GetNetworkIDOperation{
		Network: o.Network,
	}
}

func (o *GetNetworkIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNetworkIDOperation) {
	if o == nil {
		return
	}
	o.Network = op.Network
}
func (o *GetNetworkIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetNetworkIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNetworkIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNetworkIDResponse structure represents the ApiGetNetworkId operation response
type GetNetworkIDResponse struct {
	GUID string `idl:"name:pGuid;string" json:"guid"`
	// Return: The ApiGetNetworkId return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetNetworkIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetNetworkIDOperation {
	if o == nil {
		return &xxx_GetNetworkIDOperation{}
	}
	return &xxx_GetNetworkIDOperation{
		GUID:   o.GUID,
		Return: o.Return,
	}
}

func (o *GetNetworkIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNetworkIDOperation) {
	if o == nil {
		return
	}
	o.GUID = op.GUID
	o.Return = op.Return
}
func (o *GetNetworkIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetNetworkIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNetworkIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetNetworkPriorityOrderOperation structure represents the ApiSetNetworkPriorityOrder operation
type xxx_SetNetworkPriorityOrderOperation struct {
	NetworkCount  uint32   `idl:"name:NetworkCount" json:"network_count"`
	NetworkIDList []string `idl:"name:NetworkIdList;size_is:(NetworkCount);string" json:"network_id_list"`
	Return        uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_SetNetworkPriorityOrderOperation) OpNum() int { return 87 }

func (o *xxx_SetNetworkPriorityOrderOperation) OpName() string {
	return "/clusapi2/v2/ApiSetNetworkPriorityOrder"
}

func (o *xxx_SetNetworkPriorityOrderOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.NetworkIDList != nil && o.NetworkCount == 0 {
		o.NetworkCount = uint32(len(o.NetworkIDList))
	}
	if o.NetworkCount > uint32(1000) {
		return fmt.Errorf("NetworkCount is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNetworkPriorityOrderOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// NetworkCount {in} (1:{range=(0,1000), alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NetworkCount); err != nil {
			return err
		}
	}
	// NetworkIdList {in} (1:{string}[dim:0,size_is=NetworkCount])(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		dimSize1 := uint64(o.NetworkCount)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.NetworkIDList {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.NetworkIDList[i1] != "" {
				_ptr_NetworkIdList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if err := ndr.WriteUTF16NString(ctx, w, o.NetworkIDList[i1]); err != nil {
						return err
					}
					return nil
				})
				if err := w.WritePointer(&o.NetworkIDList[i1], _ptr_NetworkIdList); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.NetworkIDList); uint64(i1) < sizeInfo[0]; i1++ {
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

func (o *xxx_SetNetworkPriorityOrderOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// NetworkCount {in} (1:{range=(0,1000), alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NetworkCount); err != nil {
			return err
		}
	}
	// NetworkIdList {in} (1:{string}[dim:0,size_is=NetworkCount])(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.NetworkIDList", sizeInfo[0])
		}
		o.NetworkIDList = make([]string, sizeInfo[0])
		for i1 := range o.NetworkIDList {
			i1 := i1
			_ptr_NetworkIdList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadUTF16NString(ctx, w, &o.NetworkIDList[i1]); err != nil {
					return err
				}
				return nil
			})
			_s_NetworkIdList := func(ptr interface{}) { o.NetworkIDList[i1] = *ptr.(*string) }
			if err := w.ReadPointer(&o.NetworkIDList[i1], _s_NetworkIdList, _ptr_NetworkIdList); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNetworkPriorityOrderOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNetworkPriorityOrderOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetNetworkPriorityOrderOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetNetworkPriorityOrderRequest structure represents the ApiSetNetworkPriorityOrder operation request
type SetNetworkPriorityOrderRequest struct {
	NetworkCount  uint32   `idl:"name:NetworkCount" json:"network_count"`
	NetworkIDList []string `idl:"name:NetworkIdList;size_is:(NetworkCount);string" json:"network_id_list"`
}

func (o *SetNetworkPriorityOrderRequest) xxx_ToOp(ctx context.Context) *xxx_SetNetworkPriorityOrderOperation {
	if o == nil {
		return &xxx_SetNetworkPriorityOrderOperation{}
	}
	return &xxx_SetNetworkPriorityOrderOperation{
		NetworkCount:  o.NetworkCount,
		NetworkIDList: o.NetworkIDList,
	}
}

func (o *SetNetworkPriorityOrderRequest) xxx_FromOp(ctx context.Context, op *xxx_SetNetworkPriorityOrderOperation) {
	if o == nil {
		return
	}
	o.NetworkCount = op.NetworkCount
	o.NetworkIDList = op.NetworkIDList
}
func (o *SetNetworkPriorityOrderRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetNetworkPriorityOrderRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNetworkPriorityOrderOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetNetworkPriorityOrderResponse structure represents the ApiSetNetworkPriorityOrder operation response
type SetNetworkPriorityOrderResponse struct {
	// Return: The ApiSetNetworkPriorityOrder return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetNetworkPriorityOrderResponse) xxx_ToOp(ctx context.Context) *xxx_SetNetworkPriorityOrderOperation {
	if o == nil {
		return &xxx_SetNetworkPriorityOrderOperation{}
	}
	return &xxx_SetNetworkPriorityOrderOperation{
		Return: o.Return,
	}
}

func (o *SetNetworkPriorityOrderResponse) xxx_FromOp(ctx context.Context, op *xxx_SetNetworkPriorityOrderOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetNetworkPriorityOrderResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetNetworkPriorityOrderResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetNetworkPriorityOrderOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_NodeNetworkControlOperation structure represents the ApiNodeNetworkControl operation
type xxx_NodeNetworkControlOperation struct {
	Network        *Network `idl:"name:hNetwork" json:"network"`
	Node           *Node    `idl:"name:hNode" json:"node"`
	ControlCode    uint32   `idl:"name:dwControlCode" json:"control_code"`
	InBuffer       []byte   `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize   uint32   `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBuffer      []byte   `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	OutBufferSize  uint32   `idl:"name:nOutBufferSize" json:"out_buffer_size"`
	BytesReturned  uint32   `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32   `idl:"name:lpcbRequired" json:"required_length"`
	Return         uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_NodeNetworkControlOperation) OpNum() int { return 88 }

func (o *xxx_NodeNetworkControlOperation) OpName() string {
	return "/clusapi2/v2/ApiNodeNetworkControl"
}

func (o *xxx_NodeNetworkControlOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InBuffer != nil && o.InBufferSize == 0 {
		o.InBufferSize = uint32(len(o.InBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeNetworkControlOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNetwork {in} (1:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Network != nil {
			if err := o.Network.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Network{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node != nil {
			if err := o.Node.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		if o.InBuffer != nil || o.InBufferSize > 0 {
			_ptr_lpInBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.InBufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.InBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.InBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.InBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InBuffer, _ptr_lpInBuffer); err != nil {
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
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeNetworkControlOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNetwork {in} (1:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Network == nil {
			o.Network = &Network{}
		}
		if err := o.Network.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node == nil {
			o.Node = &Node{}
		}
		if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		_ptr_lpInBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.InBuffer", sizeInfo[0])
			}
			o.InBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.InBuffer {
				i1 := i1
				if err := w.ReadData(&o.InBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_lpInBuffer := func(ptr interface{}) { o.InBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.InBuffer, _s_lpInBuffer, _ptr_lpInBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeNetworkControlOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.OutBuffer != nil && o.BytesReturned == 0 {
		o.BytesReturned = uint32(len(o.OutBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeNetworkControlOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		dimSize1 := uint64(o.OutBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.BytesReturned)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.OutBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.OutBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.OutBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeNetworkControlOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OutBuffer", sizeInfo[0])
		}
		o.OutBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.OutBuffer {
			i1 := i1
			if err := w.ReadData(&o.OutBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// NodeNetworkControlRequest structure represents the ApiNodeNetworkControl operation request
type NodeNetworkControlRequest struct {
	Network       *Network `idl:"name:hNetwork" json:"network"`
	Node          *Node    `idl:"name:hNode" json:"node"`
	ControlCode   uint32   `idl:"name:dwControlCode" json:"control_code"`
	InBuffer      []byte   `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize  uint32   `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBufferSize uint32   `idl:"name:nOutBufferSize" json:"out_buffer_size"`
}

func (o *NodeNetworkControlRequest) xxx_ToOp(ctx context.Context) *xxx_NodeNetworkControlOperation {
	if o == nil {
		return &xxx_NodeNetworkControlOperation{}
	}
	return &xxx_NodeNetworkControlOperation{
		Network:       o.Network,
		Node:          o.Node,
		ControlCode:   o.ControlCode,
		InBuffer:      o.InBuffer,
		InBufferSize:  o.InBufferSize,
		OutBufferSize: o.OutBufferSize,
	}
}

func (o *NodeNetworkControlRequest) xxx_FromOp(ctx context.Context, op *xxx_NodeNetworkControlOperation) {
	if o == nil {
		return
	}
	o.Network = op.Network
	o.Node = op.Node
	o.ControlCode = op.ControlCode
	o.InBuffer = op.InBuffer
	o.InBufferSize = op.InBufferSize
	o.OutBufferSize = op.OutBufferSize
}
func (o *NodeNetworkControlRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *NodeNetworkControlRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NodeNetworkControlOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NodeNetworkControlResponse structure represents the ApiNodeNetworkControl operation response
type NodeNetworkControlResponse struct {
	OutBuffer      []byte `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	BytesReturned  uint32 `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32 `idl:"name:lpcbRequired" json:"required_length"`
	// Return: The ApiNodeNetworkControl return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *NodeNetworkControlResponse) xxx_ToOp(ctx context.Context) *xxx_NodeNetworkControlOperation {
	if o == nil {
		return &xxx_NodeNetworkControlOperation{}
	}
	return &xxx_NodeNetworkControlOperation{
		OutBuffer:      o.OutBuffer,
		BytesReturned:  o.BytesReturned,
		RequiredLength: o.RequiredLength,
		Return:         o.Return,
	}
}

func (o *NodeNetworkControlResponse) xxx_FromOp(ctx context.Context, op *xxx_NodeNetworkControlOperation) {
	if o == nil {
		return
	}
	o.OutBuffer = op.OutBuffer
	o.BytesReturned = op.BytesReturned
	o.RequiredLength = op.RequiredLength
	o.Return = op.Return
}
func (o *NodeNetworkControlResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *NodeNetworkControlResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NodeNetworkControlOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_NetworkControlOperation structure represents the ApiNetworkControl operation
type xxx_NetworkControlOperation struct {
	Network        *Network `idl:"name:hNetwork" json:"network"`
	ControlCode    uint32   `idl:"name:dwControlCode" json:"control_code"`
	InBuffer       []byte   `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize   uint32   `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBuffer      []byte   `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	OutBufferSize  uint32   `idl:"name:nOutBufferSize" json:"out_buffer_size"`
	BytesReturned  uint32   `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32   `idl:"name:lpcbRequired" json:"required_length"`
	Return         uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_NetworkControlOperation) OpNum() int { return 89 }

func (o *xxx_NetworkControlOperation) OpName() string { return "/clusapi2/v2/ApiNetworkControl" }

func (o *xxx_NetworkControlOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InBuffer != nil && o.InBufferSize == 0 {
		o.InBufferSize = uint32(len(o.InBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NetworkControlOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNetwork {in} (1:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Network != nil {
			if err := o.Network.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Network{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		if o.InBuffer != nil || o.InBufferSize > 0 {
			_ptr_lpInBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.InBufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.InBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.InBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.InBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InBuffer, _ptr_lpInBuffer); err != nil {
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
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NetworkControlOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNetwork {in} (1:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Network == nil {
			o.Network = &Network{}
		}
		if err := o.Network.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		_ptr_lpInBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.InBuffer", sizeInfo[0])
			}
			o.InBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.InBuffer {
				i1 := i1
				if err := w.ReadData(&o.InBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_lpInBuffer := func(ptr interface{}) { o.InBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.InBuffer, _s_lpInBuffer, _ptr_lpInBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NetworkControlOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.OutBuffer != nil && o.BytesReturned == 0 {
		o.BytesReturned = uint32(len(o.OutBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NetworkControlOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		dimSize1 := uint64(o.OutBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.BytesReturned)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.OutBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.OutBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.OutBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NetworkControlOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OutBuffer", sizeInfo[0])
		}
		o.OutBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.OutBuffer {
			i1 := i1
			if err := w.ReadData(&o.OutBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// NetworkControlRequest structure represents the ApiNetworkControl operation request
type NetworkControlRequest struct {
	Network       *Network `idl:"name:hNetwork" json:"network"`
	ControlCode   uint32   `idl:"name:dwControlCode" json:"control_code"`
	InBuffer      []byte   `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize  uint32   `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBufferSize uint32   `idl:"name:nOutBufferSize" json:"out_buffer_size"`
}

func (o *NetworkControlRequest) xxx_ToOp(ctx context.Context) *xxx_NetworkControlOperation {
	if o == nil {
		return &xxx_NetworkControlOperation{}
	}
	return &xxx_NetworkControlOperation{
		Network:       o.Network,
		ControlCode:   o.ControlCode,
		InBuffer:      o.InBuffer,
		InBufferSize:  o.InBufferSize,
		OutBufferSize: o.OutBufferSize,
	}
}

func (o *NetworkControlRequest) xxx_FromOp(ctx context.Context, op *xxx_NetworkControlOperation) {
	if o == nil {
		return
	}
	o.Network = op.Network
	o.ControlCode = op.ControlCode
	o.InBuffer = op.InBuffer
	o.InBufferSize = op.InBufferSize
	o.OutBufferSize = op.OutBufferSize
}
func (o *NetworkControlRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *NetworkControlRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NetworkControlOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NetworkControlResponse structure represents the ApiNetworkControl operation response
type NetworkControlResponse struct {
	OutBuffer      []byte `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	BytesReturned  uint32 `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32 `idl:"name:lpcbRequired" json:"required_length"`
	// Return: The ApiNetworkControl return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *NetworkControlResponse) xxx_ToOp(ctx context.Context) *xxx_NetworkControlOperation {
	if o == nil {
		return &xxx_NetworkControlOperation{}
	}
	return &xxx_NetworkControlOperation{
		OutBuffer:      o.OutBuffer,
		BytesReturned:  o.BytesReturned,
		RequiredLength: o.RequiredLength,
		Return:         o.Return,
	}
}

func (o *NetworkControlResponse) xxx_FromOp(ctx context.Context, op *xxx_NetworkControlOperation) {
	if o == nil {
		return
	}
	o.OutBuffer = op.OutBuffer
	o.BytesReturned = op.BytesReturned
	o.RequiredLength = op.RequiredLength
	o.Return = op.Return
}
func (o *NetworkControlResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *NetworkControlResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NetworkControlOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddNotifyNetworkOperation structure represents the ApiAddNotifyNetwork operation
type xxx_AddNotifyNetworkOperation struct {
	Notify        *Notify  `idl:"name:hNotify" json:"notify"`
	Network       *Network `idl:"name:hNetwork" json:"network"`
	Filter        uint32   `idl:"name:dwFilter" json:"filter"`
	NotifyKey     uint32   `idl:"name:dwNotifyKey" json:"notify_key"`
	StateSequence uint32   `idl:"name:dwStateSequence" json:"state_sequence"`
	Return        uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_AddNotifyNetworkOperation) OpNum() int { return 90 }

func (o *xxx_AddNotifyNetworkOperation) OpName() string { return "/clusapi2/v2/ApiAddNotifyNetwork" }

func (o *xxx_AddNotifyNetworkOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyNetworkOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify != nil {
			if err := o.Notify.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hNetwork {in} (1:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Network != nil {
			if err := o.Network.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Network{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NotifyKey); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyNetworkOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify == nil {
			o.Notify = &Notify{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hNetwork {in} (1:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Network == nil {
			o.Network = &Network{}
		}
		if err := o.Network.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NotifyKey); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyNetworkOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyNetworkOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// dwStateSequence {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.StateSequence); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyNetworkOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// dwStateSequence {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.StateSequence); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddNotifyNetworkRequest structure represents the ApiAddNotifyNetwork operation request
type AddNotifyNetworkRequest struct {
	Notify    *Notify  `idl:"name:hNotify" json:"notify"`
	Network   *Network `idl:"name:hNetwork" json:"network"`
	Filter    uint32   `idl:"name:dwFilter" json:"filter"`
	NotifyKey uint32   `idl:"name:dwNotifyKey" json:"notify_key"`
}

func (o *AddNotifyNetworkRequest) xxx_ToOp(ctx context.Context) *xxx_AddNotifyNetworkOperation {
	if o == nil {
		return &xxx_AddNotifyNetworkOperation{}
	}
	return &xxx_AddNotifyNetworkOperation{
		Notify:    o.Notify,
		Network:   o.Network,
		Filter:    o.Filter,
		NotifyKey: o.NotifyKey,
	}
}

func (o *AddNotifyNetworkRequest) xxx_FromOp(ctx context.Context, op *xxx_AddNotifyNetworkOperation) {
	if o == nil {
		return
	}
	o.Notify = op.Notify
	o.Network = op.Network
	o.Filter = op.Filter
	o.NotifyKey = op.NotifyKey
}
func (o *AddNotifyNetworkRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddNotifyNetworkRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddNotifyNetworkOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddNotifyNetworkResponse structure represents the ApiAddNotifyNetwork operation response
type AddNotifyNetworkResponse struct {
	StateSequence uint32 `idl:"name:dwStateSequence" json:"state_sequence"`
	// Return: The ApiAddNotifyNetwork return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AddNotifyNetworkResponse) xxx_ToOp(ctx context.Context) *xxx_AddNotifyNetworkOperation {
	if o == nil {
		return &xxx_AddNotifyNetworkOperation{}
	}
	return &xxx_AddNotifyNetworkOperation{
		StateSequence: o.StateSequence,
		Return:        o.Return,
	}
}

func (o *AddNotifyNetworkResponse) xxx_FromOp(ctx context.Context, op *xxx_AddNotifyNetworkOperation) {
	if o == nil {
		return
	}
	o.StateSequence = op.StateSequence
	o.Return = op.Return
}
func (o *AddNotifyNetworkResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddNotifyNetworkResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddNotifyNetworkOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReAddNotifyNetworkOperation structure represents the ApiReAddNotifyNetwork operation
type xxx_ReAddNotifyNetworkOperation struct {
	Notify        *Notify  `idl:"name:hNotify" json:"notify"`
	Network       *Network `idl:"name:hNetwork" json:"network"`
	Filter        uint32   `idl:"name:dwFilter" json:"filter"`
	NotifyKey     uint32   `idl:"name:dwNotifyKey" json:"notify_key"`
	StateSequence uint32   `idl:"name:StateSequence" json:"state_sequence"`
	Return        uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_ReAddNotifyNetworkOperation) OpNum() int { return 91 }

func (o *xxx_ReAddNotifyNetworkOperation) OpName() string {
	return "/clusapi2/v2/ApiReAddNotifyNetwork"
}

func (o *xxx_ReAddNotifyNetworkOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyNetworkOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify != nil {
			if err := o.Notify.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hNetwork {in} (1:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Network != nil {
			if err := o.Network.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Network{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NotifyKey); err != nil {
			return err
		}
	}
	// StateSequence {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.StateSequence); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyNetworkOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify == nil {
			o.Notify = &Notify{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hNetwork {in} (1:{context_handle, alias=HNETWORK_RPC, names=ndr_context_handle}(struct))
	{
		if o.Network == nil {
			o.Network = &Network{}
		}
		if err := o.Network.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NotifyKey); err != nil {
			return err
		}
	}
	// StateSequence {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.StateSequence); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyNetworkOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyNetworkOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyNetworkOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ReAddNotifyNetworkRequest structure represents the ApiReAddNotifyNetwork operation request
type ReAddNotifyNetworkRequest struct {
	Notify        *Notify  `idl:"name:hNotify" json:"notify"`
	Network       *Network `idl:"name:hNetwork" json:"network"`
	Filter        uint32   `idl:"name:dwFilter" json:"filter"`
	NotifyKey     uint32   `idl:"name:dwNotifyKey" json:"notify_key"`
	StateSequence uint32   `idl:"name:StateSequence" json:"state_sequence"`
}

func (o *ReAddNotifyNetworkRequest) xxx_ToOp(ctx context.Context) *xxx_ReAddNotifyNetworkOperation {
	if o == nil {
		return &xxx_ReAddNotifyNetworkOperation{}
	}
	return &xxx_ReAddNotifyNetworkOperation{
		Notify:        o.Notify,
		Network:       o.Network,
		Filter:        o.Filter,
		NotifyKey:     o.NotifyKey,
		StateSequence: o.StateSequence,
	}
}

func (o *ReAddNotifyNetworkRequest) xxx_FromOp(ctx context.Context, op *xxx_ReAddNotifyNetworkOperation) {
	if o == nil {
		return
	}
	o.Notify = op.Notify
	o.Network = op.Network
	o.Filter = op.Filter
	o.NotifyKey = op.NotifyKey
	o.StateSequence = op.StateSequence
}
func (o *ReAddNotifyNetworkRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ReAddNotifyNetworkRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReAddNotifyNetworkOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReAddNotifyNetworkResponse structure represents the ApiReAddNotifyNetwork operation response
type ReAddNotifyNetworkResponse struct {
	// Return: The ApiReAddNotifyNetwork return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ReAddNotifyNetworkResponse) xxx_ToOp(ctx context.Context) *xxx_ReAddNotifyNetworkOperation {
	if o == nil {
		return &xxx_ReAddNotifyNetworkOperation{}
	}
	return &xxx_ReAddNotifyNetworkOperation{
		Return: o.Return,
	}
}

func (o *ReAddNotifyNetworkResponse) xxx_FromOp(ctx context.Context, op *xxx_ReAddNotifyNetworkOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ReAddNotifyNetworkResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ReAddNotifyNetworkResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReAddNotifyNetworkOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenNetInterfaceOperation structure represents the ApiOpenNetInterface operation
type xxx_OpenNetInterfaceOperation struct {
	NetInterfaceName string        `idl:"name:lpszNetInterfaceName;string" json:"net_interface_name"`
	Status           uint32        `idl:"name:Status" json:"status"`
	Return           *NetInterface `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenNetInterfaceOperation) OpNum() int { return 92 }

func (o *xxx_OpenNetInterfaceOperation) OpName() string { return "/clusapi2/v2/ApiOpenNetInterface" }

func (o *xxx_OpenNetInterfaceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNetInterfaceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpszNetInterfaceName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.NetInterfaceName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNetInterfaceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpszNetInterfaceName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.NetInterfaceName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNetInterfaceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenNetInterfaceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.WriteData(o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HNETINTERFACE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return != nil {
			if err := o.Return.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NetInterface{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_OpenNetInterfaceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Status {out} (1:{pointer=ref}*(1)(error_status_t))
	{
		if err := w.ReadData(&o.Status); err != nil {
			return err
		}
	}
	// Return {out} (1:{context_handle, alias=HNETINTERFACE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Return == nil {
			o.Return = &NetInterface{}
		}
		if err := o.Return.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// OpenNetInterfaceRequest structure represents the ApiOpenNetInterface operation request
type OpenNetInterfaceRequest struct {
	NetInterfaceName string `idl:"name:lpszNetInterfaceName;string" json:"net_interface_name"`
}

func (o *OpenNetInterfaceRequest) xxx_ToOp(ctx context.Context) *xxx_OpenNetInterfaceOperation {
	if o == nil {
		return &xxx_OpenNetInterfaceOperation{}
	}
	return &xxx_OpenNetInterfaceOperation{
		NetInterfaceName: o.NetInterfaceName,
	}
}

func (o *OpenNetInterfaceRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenNetInterfaceOperation) {
	if o == nil {
		return
	}
	o.NetInterfaceName = op.NetInterfaceName
}
func (o *OpenNetInterfaceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *OpenNetInterfaceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenNetInterfaceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenNetInterfaceResponse structure represents the ApiOpenNetInterface operation response
type OpenNetInterfaceResponse struct {
	Status uint32 `idl:"name:Status" json:"status"`
	// Return: The ApiOpenNetInterface return value.
	Return *NetInterface `idl:"name:Return" json:"return"`
}

func (o *OpenNetInterfaceResponse) xxx_ToOp(ctx context.Context) *xxx_OpenNetInterfaceOperation {
	if o == nil {
		return &xxx_OpenNetInterfaceOperation{}
	}
	return &xxx_OpenNetInterfaceOperation{
		Status: o.Status,
		Return: o.Return,
	}
}

func (o *OpenNetInterfaceResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenNetInterfaceOperation) {
	if o == nil {
		return
	}
	o.Status = op.Status
	o.Return = op.Return
}
func (o *OpenNetInterfaceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *OpenNetInterfaceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenNetInterfaceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CloseNetInterfaceOperation structure represents the ApiCloseNetInterface operation
type xxx_CloseNetInterfaceOperation struct {
	NetInterface *NetInterface `idl:"name:NetInterface" json:"net_interface"`
	Return       uint32        `idl:"name:Return" json:"return"`
}

func (o *xxx_CloseNetInterfaceOperation) OpNum() int { return 93 }

func (o *xxx_CloseNetInterfaceOperation) OpName() string { return "/clusapi2/v2/ApiCloseNetInterface" }

func (o *xxx_CloseNetInterfaceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseNetInterfaceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// NetInterface {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HNETINTERFACE_RPC, names=ndr_context_handle}(struct))
	{
		if o.NetInterface != nil {
			if err := o.NetInterface.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NetInterface{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_CloseNetInterfaceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// NetInterface {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HNETINTERFACE_RPC, names=ndr_context_handle}(struct))
	{
		if o.NetInterface == nil {
			o.NetInterface = &NetInterface{}
		}
		if err := o.NetInterface.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseNetInterfaceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseNetInterfaceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// NetInterface {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HNETINTERFACE_RPC, names=ndr_context_handle}(struct))
	{
		if o.NetInterface != nil {
			if err := o.NetInterface.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NetInterface{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CloseNetInterfaceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// NetInterface {in, out} (1:{pointer=ref}*(1))(2:{context_handle, alias=HNETINTERFACE_RPC, names=ndr_context_handle}(struct))
	{
		if o.NetInterface == nil {
			o.NetInterface = &NetInterface{}
		}
		if err := o.NetInterface.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CloseNetInterfaceRequest structure represents the ApiCloseNetInterface operation request
type CloseNetInterfaceRequest struct {
	NetInterface *NetInterface `idl:"name:NetInterface" json:"net_interface"`
}

func (o *CloseNetInterfaceRequest) xxx_ToOp(ctx context.Context) *xxx_CloseNetInterfaceOperation {
	if o == nil {
		return &xxx_CloseNetInterfaceOperation{}
	}
	return &xxx_CloseNetInterfaceOperation{
		NetInterface: o.NetInterface,
	}
}

func (o *CloseNetInterfaceRequest) xxx_FromOp(ctx context.Context, op *xxx_CloseNetInterfaceOperation) {
	if o == nil {
		return
	}
	o.NetInterface = op.NetInterface
}
func (o *CloseNetInterfaceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CloseNetInterfaceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseNetInterfaceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CloseNetInterfaceResponse structure represents the ApiCloseNetInterface operation response
type CloseNetInterfaceResponse struct {
	NetInterface *NetInterface `idl:"name:NetInterface" json:"net_interface"`
	// Return: The ApiCloseNetInterface return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CloseNetInterfaceResponse) xxx_ToOp(ctx context.Context) *xxx_CloseNetInterfaceOperation {
	if o == nil {
		return &xxx_CloseNetInterfaceOperation{}
	}
	return &xxx_CloseNetInterfaceOperation{
		NetInterface: o.NetInterface,
		Return:       o.Return,
	}
}

func (o *CloseNetInterfaceResponse) xxx_FromOp(ctx context.Context, op *xxx_CloseNetInterfaceOperation) {
	if o == nil {
		return
	}
	o.NetInterface = op.NetInterface
	o.Return = op.Return
}
func (o *CloseNetInterfaceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CloseNetInterfaceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CloseNetInterfaceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNetInterfaceStateOperation structure represents the ApiGetNetInterfaceState operation
type xxx_GetNetInterfaceStateOperation struct {
	NetInterface *NetInterface `idl:"name:hNetInterface" json:"net_interface"`
	State        uint32        `idl:"name:State" json:"state"`
	Return       uint32        `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNetInterfaceStateOperation) OpNum() int { return 94 }

func (o *xxx_GetNetInterfaceStateOperation) OpName() string {
	return "/clusapi2/v2/ApiGetNetInterfaceState"
}

func (o *xxx_GetNetInterfaceStateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetInterfaceStateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNetInterface {in} (1:{context_handle, alias=HNETINTERFACE_RPC, names=ndr_context_handle}(struct))
	{
		if o.NetInterface != nil {
			if err := o.NetInterface.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NetInterface{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetNetInterfaceStateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNetInterface {in} (1:{context_handle, alias=HNETINTERFACE_RPC, names=ndr_context_handle}(struct))
	{
		if o.NetInterface == nil {
			o.NetInterface = &NetInterface{}
		}
		if err := o.NetInterface.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetInterfaceStateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetInterfaceStateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// State {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.State); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetInterfaceStateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// State {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.State); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNetInterfaceStateRequest structure represents the ApiGetNetInterfaceState operation request
type GetNetInterfaceStateRequest struct {
	NetInterface *NetInterface `idl:"name:hNetInterface" json:"net_interface"`
}

func (o *GetNetInterfaceStateRequest) xxx_ToOp(ctx context.Context) *xxx_GetNetInterfaceStateOperation {
	if o == nil {
		return &xxx_GetNetInterfaceStateOperation{}
	}
	return &xxx_GetNetInterfaceStateOperation{
		NetInterface: o.NetInterface,
	}
}

func (o *GetNetInterfaceStateRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNetInterfaceStateOperation) {
	if o == nil {
		return
	}
	o.NetInterface = op.NetInterface
}
func (o *GetNetInterfaceStateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetNetInterfaceStateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNetInterfaceStateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNetInterfaceStateResponse structure represents the ApiGetNetInterfaceState operation response
type GetNetInterfaceStateResponse struct {
	State uint32 `idl:"name:State" json:"state"`
	// Return: The ApiGetNetInterfaceState return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetNetInterfaceStateResponse) xxx_ToOp(ctx context.Context) *xxx_GetNetInterfaceStateOperation {
	if o == nil {
		return &xxx_GetNetInterfaceStateOperation{}
	}
	return &xxx_GetNetInterfaceStateOperation{
		State:  o.State,
		Return: o.Return,
	}
}

func (o *GetNetInterfaceStateResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNetInterfaceStateOperation) {
	if o == nil {
		return
	}
	o.State = op.State
	o.Return = op.Return
}
func (o *GetNetInterfaceStateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetNetInterfaceStateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNetInterfaceStateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNetInterfaceOperation structure represents the ApiGetNetInterface operation
type xxx_GetNetInterfaceOperation struct {
	NodeName      string `idl:"name:lpszNodeName;string" json:"node_name"`
	NetworkName   string `idl:"name:lpszNetworkName;string" json:"network_name"`
	InterfaceName string `idl:"name:lppszInterfaceName;string" json:"interface_name"`
	Return        uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNetInterfaceOperation) OpNum() int { return 95 }

func (o *xxx_GetNetInterfaceOperation) OpName() string { return "/clusapi2/v2/ApiGetNetInterface" }

func (o *xxx_GetNetInterfaceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetInterfaceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpszNodeName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.NodeName); err != nil {
			return err
		}
	}
	// lpszNetworkName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.NetworkName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetInterfaceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpszNodeName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.NodeName); err != nil {
			return err
		}
	}
	// lpszNetworkName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.NetworkName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetInterfaceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetInterfaceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lppszInterfaceName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.InterfaceName != "" {
			_ptr_lppszInterfaceName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.InterfaceName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.InterfaceName, _ptr_lppszInterfaceName); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetInterfaceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lppszInterfaceName {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_lppszInterfaceName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.InterfaceName); err != nil {
				return err
			}
			return nil
		})
		_s_lppszInterfaceName := func(ptr interface{}) { o.InterfaceName = *ptr.(*string) }
		if err := w.ReadPointer(&o.InterfaceName, _s_lppszInterfaceName, _ptr_lppszInterfaceName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNetInterfaceRequest structure represents the ApiGetNetInterface operation request
type GetNetInterfaceRequest struct {
	NodeName    string `idl:"name:lpszNodeName;string" json:"node_name"`
	NetworkName string `idl:"name:lpszNetworkName;string" json:"network_name"`
}

func (o *GetNetInterfaceRequest) xxx_ToOp(ctx context.Context) *xxx_GetNetInterfaceOperation {
	if o == nil {
		return &xxx_GetNetInterfaceOperation{}
	}
	return &xxx_GetNetInterfaceOperation{
		NodeName:    o.NodeName,
		NetworkName: o.NetworkName,
	}
}

func (o *GetNetInterfaceRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNetInterfaceOperation) {
	if o == nil {
		return
	}
	o.NodeName = op.NodeName
	o.NetworkName = op.NetworkName
}
func (o *GetNetInterfaceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetNetInterfaceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNetInterfaceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNetInterfaceResponse structure represents the ApiGetNetInterface operation response
type GetNetInterfaceResponse struct {
	InterfaceName string `idl:"name:lppszInterfaceName;string" json:"interface_name"`
	// Return: The ApiGetNetInterface return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetNetInterfaceResponse) xxx_ToOp(ctx context.Context) *xxx_GetNetInterfaceOperation {
	if o == nil {
		return &xxx_GetNetInterfaceOperation{}
	}
	return &xxx_GetNetInterfaceOperation{
		InterfaceName: o.InterfaceName,
		Return:        o.Return,
	}
}

func (o *GetNetInterfaceResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNetInterfaceOperation) {
	if o == nil {
		return
	}
	o.InterfaceName = op.InterfaceName
	o.Return = op.Return
}
func (o *GetNetInterfaceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetNetInterfaceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNetInterfaceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNetInterfaceIDOperation structure represents the ApiGetNetInterfaceId operation
type xxx_GetNetInterfaceIDOperation struct {
	NetInterface *NetInterface `idl:"name:hNetInterface" json:"net_interface"`
	GUID         string        `idl:"name:pGuid;string" json:"guid"`
	Return       uint32        `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNetInterfaceIDOperation) OpNum() int { return 96 }

func (o *xxx_GetNetInterfaceIDOperation) OpName() string { return "/clusapi2/v2/ApiGetNetInterfaceId" }

func (o *xxx_GetNetInterfaceIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetInterfaceIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNetInterface {in} (1:{context_handle, alias=HNETINTERFACE_RPC, names=ndr_context_handle}(struct))
	{
		if o.NetInterface != nil {
			if err := o.NetInterface.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NetInterface{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetNetInterfaceIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNetInterface {in} (1:{context_handle, alias=HNETINTERFACE_RPC, names=ndr_context_handle}(struct))
	{
		if o.NetInterface == nil {
			o.NetInterface = &NetInterface{}
		}
		if err := o.NetInterface.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetInterfaceIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetInterfaceIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pGuid {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.GUID != "" {
			_ptr_pGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.GUID); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.GUID, _ptr_pGuid); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNetInterfaceIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pGuid {out} (1:{string, pointer=ref}*(2))(2:{alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.GUID); err != nil {
				return err
			}
			return nil
		})
		_s_pGuid := func(ptr interface{}) { o.GUID = *ptr.(*string) }
		if err := w.ReadPointer(&o.GUID, _s_pGuid, _ptr_pGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNetInterfaceIDRequest structure represents the ApiGetNetInterfaceId operation request
type GetNetInterfaceIDRequest struct {
	NetInterface *NetInterface `idl:"name:hNetInterface" json:"net_interface"`
}

func (o *GetNetInterfaceIDRequest) xxx_ToOp(ctx context.Context) *xxx_GetNetInterfaceIDOperation {
	if o == nil {
		return &xxx_GetNetInterfaceIDOperation{}
	}
	return &xxx_GetNetInterfaceIDOperation{
		NetInterface: o.NetInterface,
	}
}

func (o *GetNetInterfaceIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNetInterfaceIDOperation) {
	if o == nil {
		return
	}
	o.NetInterface = op.NetInterface
}
func (o *GetNetInterfaceIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetNetInterfaceIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNetInterfaceIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNetInterfaceIDResponse structure represents the ApiGetNetInterfaceId operation response
type GetNetInterfaceIDResponse struct {
	GUID string `idl:"name:pGuid;string" json:"guid"`
	// Return: The ApiGetNetInterfaceId return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetNetInterfaceIDResponse) xxx_ToOp(ctx context.Context) *xxx_GetNetInterfaceIDOperation {
	if o == nil {
		return &xxx_GetNetInterfaceIDOperation{}
	}
	return &xxx_GetNetInterfaceIDOperation{
		GUID:   o.GUID,
		Return: o.Return,
	}
}

func (o *GetNetInterfaceIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNetInterfaceIDOperation) {
	if o == nil {
		return
	}
	o.GUID = op.GUID
	o.Return = op.Return
}
func (o *GetNetInterfaceIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetNetInterfaceIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNetInterfaceIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_NodeNetInterfaceControlOperation structure represents the ApiNodeNetInterfaceControl operation
type xxx_NodeNetInterfaceControlOperation struct {
	NetInterface   *NetInterface `idl:"name:hNetInterface" json:"net_interface"`
	Node           *Node         `idl:"name:hNode" json:"node"`
	ControlCode    uint32        `idl:"name:dwControlCode" json:"control_code"`
	InBuffer       []byte        `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize   uint32        `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBuffer      []byte        `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	OutBufferSize  uint32        `idl:"name:nOutBufferSize" json:"out_buffer_size"`
	BytesReturned  uint32        `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32        `idl:"name:lpcbRequired" json:"required_length"`
	Return         uint32        `idl:"name:Return" json:"return"`
}

func (o *xxx_NodeNetInterfaceControlOperation) OpNum() int { return 97 }

func (o *xxx_NodeNetInterfaceControlOperation) OpName() string {
	return "/clusapi2/v2/ApiNodeNetInterfaceControl"
}

func (o *xxx_NodeNetInterfaceControlOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InBuffer != nil && o.InBufferSize == 0 {
		o.InBufferSize = uint32(len(o.InBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeNetInterfaceControlOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNetInterface {in} (1:{context_handle, alias=HNETINTERFACE_RPC, names=ndr_context_handle}(struct))
	{
		if o.NetInterface != nil {
			if err := o.NetInterface.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NetInterface{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node != nil {
			if err := o.Node.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		if o.InBuffer != nil || o.InBufferSize > 0 {
			_ptr_lpInBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.InBufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.InBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.InBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.InBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InBuffer, _ptr_lpInBuffer); err != nil {
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
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeNetInterfaceControlOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNetInterface {in} (1:{context_handle, alias=HNETINTERFACE_RPC, names=ndr_context_handle}(struct))
	{
		if o.NetInterface == nil {
			o.NetInterface = &NetInterface{}
		}
		if err := o.NetInterface.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node == nil {
			o.Node = &Node{}
		}
		if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		_ptr_lpInBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.InBuffer", sizeInfo[0])
			}
			o.InBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.InBuffer {
				i1 := i1
				if err := w.ReadData(&o.InBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_lpInBuffer := func(ptr interface{}) { o.InBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.InBuffer, _s_lpInBuffer, _ptr_lpInBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeNetInterfaceControlOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.OutBuffer != nil && o.BytesReturned == 0 {
		o.BytesReturned = uint32(len(o.OutBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeNetInterfaceControlOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		dimSize1 := uint64(o.OutBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.BytesReturned)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.OutBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.OutBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.OutBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeNetInterfaceControlOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OutBuffer", sizeInfo[0])
		}
		o.OutBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.OutBuffer {
			i1 := i1
			if err := w.ReadData(&o.OutBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// NodeNetInterfaceControlRequest structure represents the ApiNodeNetInterfaceControl operation request
type NodeNetInterfaceControlRequest struct {
	NetInterface  *NetInterface `idl:"name:hNetInterface" json:"net_interface"`
	Node          *Node         `idl:"name:hNode" json:"node"`
	ControlCode   uint32        `idl:"name:dwControlCode" json:"control_code"`
	InBuffer      []byte        `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize  uint32        `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBufferSize uint32        `idl:"name:nOutBufferSize" json:"out_buffer_size"`
}

func (o *NodeNetInterfaceControlRequest) xxx_ToOp(ctx context.Context) *xxx_NodeNetInterfaceControlOperation {
	if o == nil {
		return &xxx_NodeNetInterfaceControlOperation{}
	}
	return &xxx_NodeNetInterfaceControlOperation{
		NetInterface:  o.NetInterface,
		Node:          o.Node,
		ControlCode:   o.ControlCode,
		InBuffer:      o.InBuffer,
		InBufferSize:  o.InBufferSize,
		OutBufferSize: o.OutBufferSize,
	}
}

func (o *NodeNetInterfaceControlRequest) xxx_FromOp(ctx context.Context, op *xxx_NodeNetInterfaceControlOperation) {
	if o == nil {
		return
	}
	o.NetInterface = op.NetInterface
	o.Node = op.Node
	o.ControlCode = op.ControlCode
	o.InBuffer = op.InBuffer
	o.InBufferSize = op.InBufferSize
	o.OutBufferSize = op.OutBufferSize
}
func (o *NodeNetInterfaceControlRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *NodeNetInterfaceControlRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NodeNetInterfaceControlOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NodeNetInterfaceControlResponse structure represents the ApiNodeNetInterfaceControl operation response
type NodeNetInterfaceControlResponse struct {
	OutBuffer      []byte `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	BytesReturned  uint32 `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32 `idl:"name:lpcbRequired" json:"required_length"`
	// Return: The ApiNodeNetInterfaceControl return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *NodeNetInterfaceControlResponse) xxx_ToOp(ctx context.Context) *xxx_NodeNetInterfaceControlOperation {
	if o == nil {
		return &xxx_NodeNetInterfaceControlOperation{}
	}
	return &xxx_NodeNetInterfaceControlOperation{
		OutBuffer:      o.OutBuffer,
		BytesReturned:  o.BytesReturned,
		RequiredLength: o.RequiredLength,
		Return:         o.Return,
	}
}

func (o *NodeNetInterfaceControlResponse) xxx_FromOp(ctx context.Context, op *xxx_NodeNetInterfaceControlOperation) {
	if o == nil {
		return
	}
	o.OutBuffer = op.OutBuffer
	o.BytesReturned = op.BytesReturned
	o.RequiredLength = op.RequiredLength
	o.Return = op.Return
}
func (o *NodeNetInterfaceControlResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *NodeNetInterfaceControlResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NodeNetInterfaceControlOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_NetInterfaceControlOperation structure represents the ApiNetInterfaceControl operation
type xxx_NetInterfaceControlOperation struct {
	NetInterface   *NetInterface `idl:"name:hNetInterface" json:"net_interface"`
	ControlCode    uint32        `idl:"name:dwControlCode" json:"control_code"`
	InBuffer       []byte        `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize   uint32        `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBuffer      []byte        `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	OutBufferSize  uint32        `idl:"name:nOutBufferSize" json:"out_buffer_size"`
	BytesReturned  uint32        `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32        `idl:"name:lpcbRequired" json:"required_length"`
	Return         uint32        `idl:"name:Return" json:"return"`
}

func (o *xxx_NetInterfaceControlOperation) OpNum() int { return 98 }

func (o *xxx_NetInterfaceControlOperation) OpName() string {
	return "/clusapi2/v2/ApiNetInterfaceControl"
}

func (o *xxx_NetInterfaceControlOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InBuffer != nil && o.InBufferSize == 0 {
		o.InBufferSize = uint32(len(o.InBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NetInterfaceControlOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNetInterface {in} (1:{context_handle, alias=HNETINTERFACE_RPC, names=ndr_context_handle}(struct))
	{
		if o.NetInterface != nil {
			if err := o.NetInterface.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NetInterface{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		if o.InBuffer != nil || o.InBufferSize > 0 {
			_ptr_lpInBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.InBufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.InBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.InBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.InBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InBuffer, _ptr_lpInBuffer); err != nil {
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
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NetInterfaceControlOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNetInterface {in} (1:{context_handle, alias=HNETINTERFACE_RPC, names=ndr_context_handle}(struct))
	{
		if o.NetInterface == nil {
			o.NetInterface = &NetInterface{}
		}
		if err := o.NetInterface.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		_ptr_lpInBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.InBuffer", sizeInfo[0])
			}
			o.InBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.InBuffer {
				i1 := i1
				if err := w.ReadData(&o.InBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_lpInBuffer := func(ptr interface{}) { o.InBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.InBuffer, _s_lpInBuffer, _ptr_lpInBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NetInterfaceControlOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.OutBuffer != nil && o.BytesReturned == 0 {
		o.BytesReturned = uint32(len(o.OutBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NetInterfaceControlOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		dimSize1 := uint64(o.OutBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.BytesReturned)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.OutBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.OutBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.OutBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NetInterfaceControlOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OutBuffer", sizeInfo[0])
		}
		o.OutBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.OutBuffer {
			i1 := i1
			if err := w.ReadData(&o.OutBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// NetInterfaceControlRequest structure represents the ApiNetInterfaceControl operation request
type NetInterfaceControlRequest struct {
	NetInterface  *NetInterface `idl:"name:hNetInterface" json:"net_interface"`
	ControlCode   uint32        `idl:"name:dwControlCode" json:"control_code"`
	InBuffer      []byte        `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize  uint32        `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBufferSize uint32        `idl:"name:nOutBufferSize" json:"out_buffer_size"`
}

func (o *NetInterfaceControlRequest) xxx_ToOp(ctx context.Context) *xxx_NetInterfaceControlOperation {
	if o == nil {
		return &xxx_NetInterfaceControlOperation{}
	}
	return &xxx_NetInterfaceControlOperation{
		NetInterface:  o.NetInterface,
		ControlCode:   o.ControlCode,
		InBuffer:      o.InBuffer,
		InBufferSize:  o.InBufferSize,
		OutBufferSize: o.OutBufferSize,
	}
}

func (o *NetInterfaceControlRequest) xxx_FromOp(ctx context.Context, op *xxx_NetInterfaceControlOperation) {
	if o == nil {
		return
	}
	o.NetInterface = op.NetInterface
	o.ControlCode = op.ControlCode
	o.InBuffer = op.InBuffer
	o.InBufferSize = op.InBufferSize
	o.OutBufferSize = op.OutBufferSize
}
func (o *NetInterfaceControlRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *NetInterfaceControlRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NetInterfaceControlOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NetInterfaceControlResponse structure represents the ApiNetInterfaceControl operation response
type NetInterfaceControlResponse struct {
	OutBuffer      []byte `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	BytesReturned  uint32 `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32 `idl:"name:lpcbRequired" json:"required_length"`
	// Return: The ApiNetInterfaceControl return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *NetInterfaceControlResponse) xxx_ToOp(ctx context.Context) *xxx_NetInterfaceControlOperation {
	if o == nil {
		return &xxx_NetInterfaceControlOperation{}
	}
	return &xxx_NetInterfaceControlOperation{
		OutBuffer:      o.OutBuffer,
		BytesReturned:  o.BytesReturned,
		RequiredLength: o.RequiredLength,
		Return:         o.Return,
	}
}

func (o *NetInterfaceControlResponse) xxx_FromOp(ctx context.Context, op *xxx_NetInterfaceControlOperation) {
	if o == nil {
		return
	}
	o.OutBuffer = op.OutBuffer
	o.BytesReturned = op.BytesReturned
	o.RequiredLength = op.RequiredLength
	o.Return = op.Return
}
func (o *NetInterfaceControlResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *NetInterfaceControlResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NetInterfaceControlOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddNotifyNetInterfaceOperation structure represents the ApiAddNotifyNetInterface operation
type xxx_AddNotifyNetInterfaceOperation struct {
	Notify        *Notify       `idl:"name:hNotify" json:"notify"`
	NetInterface  *NetInterface `idl:"name:hNetInterface" json:"net_interface"`
	Filter        uint32        `idl:"name:dwFilter" json:"filter"`
	NotifyKey     uint32        `idl:"name:dwNotifyKey" json:"notify_key"`
	StateSequence uint32        `idl:"name:dwStateSequence" json:"state_sequence"`
	Return        uint32        `idl:"name:Return" json:"return"`
}

func (o *xxx_AddNotifyNetInterfaceOperation) OpNum() int { return 99 }

func (o *xxx_AddNotifyNetInterfaceOperation) OpName() string {
	return "/clusapi2/v2/ApiAddNotifyNetInterface"
}

func (o *xxx_AddNotifyNetInterfaceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyNetInterfaceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify != nil {
			if err := o.Notify.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hNetInterface {in} (1:{context_handle, alias=HNETINTERFACE_RPC, names=ndr_context_handle}(struct))
	{
		if o.NetInterface != nil {
			if err := o.NetInterface.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NetInterface{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NotifyKey); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyNetInterfaceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify == nil {
			o.Notify = &Notify{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hNetInterface {in} (1:{context_handle, alias=HNETINTERFACE_RPC, names=ndr_context_handle}(struct))
	{
		if o.NetInterface == nil {
			o.NetInterface = &NetInterface{}
		}
		if err := o.NetInterface.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NotifyKey); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyNetInterfaceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyNetInterfaceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// dwStateSequence {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.StateSequence); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddNotifyNetInterfaceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// dwStateSequence {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.StateSequence); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddNotifyNetInterfaceRequest structure represents the ApiAddNotifyNetInterface operation request
type AddNotifyNetInterfaceRequest struct {
	Notify       *Notify       `idl:"name:hNotify" json:"notify"`
	NetInterface *NetInterface `idl:"name:hNetInterface" json:"net_interface"`
	Filter       uint32        `idl:"name:dwFilter" json:"filter"`
	NotifyKey    uint32        `idl:"name:dwNotifyKey" json:"notify_key"`
}

func (o *AddNotifyNetInterfaceRequest) xxx_ToOp(ctx context.Context) *xxx_AddNotifyNetInterfaceOperation {
	if o == nil {
		return &xxx_AddNotifyNetInterfaceOperation{}
	}
	return &xxx_AddNotifyNetInterfaceOperation{
		Notify:       o.Notify,
		NetInterface: o.NetInterface,
		Filter:       o.Filter,
		NotifyKey:    o.NotifyKey,
	}
}

func (o *AddNotifyNetInterfaceRequest) xxx_FromOp(ctx context.Context, op *xxx_AddNotifyNetInterfaceOperation) {
	if o == nil {
		return
	}
	o.Notify = op.Notify
	o.NetInterface = op.NetInterface
	o.Filter = op.Filter
	o.NotifyKey = op.NotifyKey
}
func (o *AddNotifyNetInterfaceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *AddNotifyNetInterfaceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddNotifyNetInterfaceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddNotifyNetInterfaceResponse structure represents the ApiAddNotifyNetInterface operation response
type AddNotifyNetInterfaceResponse struct {
	StateSequence uint32 `idl:"name:dwStateSequence" json:"state_sequence"`
	// Return: The ApiAddNotifyNetInterface return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AddNotifyNetInterfaceResponse) xxx_ToOp(ctx context.Context) *xxx_AddNotifyNetInterfaceOperation {
	if o == nil {
		return &xxx_AddNotifyNetInterfaceOperation{}
	}
	return &xxx_AddNotifyNetInterfaceOperation{
		StateSequence: o.StateSequence,
		Return:        o.Return,
	}
}

func (o *AddNotifyNetInterfaceResponse) xxx_FromOp(ctx context.Context, op *xxx_AddNotifyNetInterfaceOperation) {
	if o == nil {
		return
	}
	o.StateSequence = op.StateSequence
	o.Return = op.Return
}
func (o *AddNotifyNetInterfaceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *AddNotifyNetInterfaceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddNotifyNetInterfaceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ReAddNotifyNetInterfaceOperation structure represents the ApiReAddNotifyNetInterface operation
type xxx_ReAddNotifyNetInterfaceOperation struct {
	Notify        *Notify       `idl:"name:hNotify" json:"notify"`
	NetInterface  *NetInterface `idl:"name:hNetInterface" json:"net_interface"`
	Filter        uint32        `idl:"name:dwFilter" json:"filter"`
	NotifyKey     uint32        `idl:"name:dwNotifyKey" json:"notify_key"`
	StateSequence uint32        `idl:"name:StateSequence" json:"state_sequence"`
	Return        uint32        `idl:"name:Return" json:"return"`
}

func (o *xxx_ReAddNotifyNetInterfaceOperation) OpNum() int { return 100 }

func (o *xxx_ReAddNotifyNetInterfaceOperation) OpName() string {
	return "/clusapi2/v2/ApiReAddNotifyNetInterface"
}

func (o *xxx_ReAddNotifyNetInterfaceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyNetInterfaceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify != nil {
			if err := o.Notify.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hNetInterface {in} (1:{context_handle, alias=HNETINTERFACE_RPC, names=ndr_context_handle}(struct))
	{
		if o.NetInterface != nil {
			if err := o.NetInterface.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&NetInterface{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NotifyKey); err != nil {
			return err
		}
	}
	// StateSequence {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.StateSequence); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyNetInterfaceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify == nil {
			o.Notify = &Notify{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hNetInterface {in} (1:{context_handle, alias=HNETINTERFACE_RPC, names=ndr_context_handle}(struct))
	{
		if o.NetInterface == nil {
			o.NetInterface = &NetInterface{}
		}
		if err := o.NetInterface.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwFilter {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Filter); err != nil {
			return err
		}
	}
	// dwNotifyKey {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NotifyKey); err != nil {
			return err
		}
	}
	// StateSequence {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.StateSequence); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyNetInterfaceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyNetInterfaceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ReAddNotifyNetInterfaceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ReAddNotifyNetInterfaceRequest structure represents the ApiReAddNotifyNetInterface operation request
type ReAddNotifyNetInterfaceRequest struct {
	Notify        *Notify       `idl:"name:hNotify" json:"notify"`
	NetInterface  *NetInterface `idl:"name:hNetInterface" json:"net_interface"`
	Filter        uint32        `idl:"name:dwFilter" json:"filter"`
	NotifyKey     uint32        `idl:"name:dwNotifyKey" json:"notify_key"`
	StateSequence uint32        `idl:"name:StateSequence" json:"state_sequence"`
}

func (o *ReAddNotifyNetInterfaceRequest) xxx_ToOp(ctx context.Context) *xxx_ReAddNotifyNetInterfaceOperation {
	if o == nil {
		return &xxx_ReAddNotifyNetInterfaceOperation{}
	}
	return &xxx_ReAddNotifyNetInterfaceOperation{
		Notify:        o.Notify,
		NetInterface:  o.NetInterface,
		Filter:        o.Filter,
		NotifyKey:     o.NotifyKey,
		StateSequence: o.StateSequence,
	}
}

func (o *ReAddNotifyNetInterfaceRequest) xxx_FromOp(ctx context.Context, op *xxx_ReAddNotifyNetInterfaceOperation) {
	if o == nil {
		return
	}
	o.Notify = op.Notify
	o.NetInterface = op.NetInterface
	o.Filter = op.Filter
	o.NotifyKey = op.NotifyKey
	o.StateSequence = op.StateSequence
}
func (o *ReAddNotifyNetInterfaceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ReAddNotifyNetInterfaceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReAddNotifyNetInterfaceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ReAddNotifyNetInterfaceResponse structure represents the ApiReAddNotifyNetInterface operation response
type ReAddNotifyNetInterfaceResponse struct {
	// Return: The ApiReAddNotifyNetInterface return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ReAddNotifyNetInterfaceResponse) xxx_ToOp(ctx context.Context) *xxx_ReAddNotifyNetInterfaceOperation {
	if o == nil {
		return &xxx_ReAddNotifyNetInterfaceOperation{}
	}
	return &xxx_ReAddNotifyNetInterfaceOperation{
		Return: o.Return,
	}
}

func (o *ReAddNotifyNetInterfaceResponse) xxx_FromOp(ctx context.Context, op *xxx_ReAddNotifyNetInterfaceOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ReAddNotifyNetInterfaceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ReAddNotifyNetInterfaceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ReAddNotifyNetInterfaceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateNodeEnumOperation structure represents the ApiCreateNodeEnum operation
type xxx_CreateNodeEnumOperation struct {
	Node       *Node     `idl:"name:hNode" json:"node"`
	Type       uint32    `idl:"name:dwType" json:"type"`
	ReturnEnum *EnumList `idl:"name:ReturnEnum" json:"return_enum"`
	Return     uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateNodeEnumOperation) OpNum() int { return 101 }

func (o *xxx_CreateNodeEnumOperation) OpName() string { return "/clusapi2/v2/ApiCreateNodeEnum" }

func (o *xxx_CreateNodeEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNodeEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node != nil {
			if err := o.Node.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNodeEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.Node == nil {
			o.Node = &Node{}
		}
		if err := o.Node.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNodeEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNodeEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ReturnEnum {out} (1:{pointer=ref}*(2))(2:{alias=PENUM_LIST}*(1))(3:{alias=ENUM_LIST}(struct))
	{
		if o.ReturnEnum != nil {
			_ptr_ReturnEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnEnum != nil {
					if err := o.ReturnEnum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EnumList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnEnum, _ptr_ReturnEnum); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNodeEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ReturnEnum {out} (1:{pointer=ref}*(2))(2:{alias=PENUM_LIST,pointer=ref}*(1))(3:{alias=ENUM_LIST}(struct))
	{
		_ptr_ReturnEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnEnum == nil {
				o.ReturnEnum = &EnumList{}
			}
			if err := o.ReturnEnum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReturnEnum := func(ptr interface{}) { o.ReturnEnum = *ptr.(**EnumList) }
		if err := w.ReadPointer(&o.ReturnEnum, _s_ReturnEnum, _ptr_ReturnEnum); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateNodeEnumRequest structure represents the ApiCreateNodeEnum operation request
type CreateNodeEnumRequest struct {
	Node *Node  `idl:"name:hNode" json:"node"`
	Type uint32 `idl:"name:dwType" json:"type"`
}

func (o *CreateNodeEnumRequest) xxx_ToOp(ctx context.Context) *xxx_CreateNodeEnumOperation {
	if o == nil {
		return &xxx_CreateNodeEnumOperation{}
	}
	return &xxx_CreateNodeEnumOperation{
		Node: o.Node,
		Type: o.Type,
	}
}

func (o *CreateNodeEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateNodeEnumOperation) {
	if o == nil {
		return
	}
	o.Node = op.Node
	o.Type = op.Type
}
func (o *CreateNodeEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateNodeEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateNodeEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateNodeEnumResponse structure represents the ApiCreateNodeEnum operation response
type CreateNodeEnumResponse struct {
	ReturnEnum *EnumList `idl:"name:ReturnEnum" json:"return_enum"`
	// Return: The ApiCreateNodeEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CreateNodeEnumResponse) xxx_ToOp(ctx context.Context) *xxx_CreateNodeEnumOperation {
	if o == nil {
		return &xxx_CreateNodeEnumOperation{}
	}
	return &xxx_CreateNodeEnumOperation{
		ReturnEnum: o.ReturnEnum,
		Return:     o.Return,
	}
}

func (o *CreateNodeEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateNodeEnumOperation) {
	if o == nil {
		return
	}
	o.ReturnEnum = op.ReturnEnum
	o.Return = op.Return
}
func (o *CreateNodeEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateNodeEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateNodeEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetClusterVersion2Operation structure represents the ApiGetClusterVersion2 operation
type xxx_GetClusterVersion2Operation struct {
	MajorVersion            uint16                         `idl:"name:lpwMajorVersion" json:"major_version"`
	MinorVersion            uint16                         `idl:"name:lpwMinorVersion" json:"minor_version"`
	BuildNumber             uint16                         `idl:"name:lpwBuildNumber" json:"build_number"`
	VendorID                string                         `idl:"name:lpszVendorId;string" json:"vendor_id"`
	CSDVersion              string                         `idl:"name:lpszCSDVersion;string" json:"csd_version"`
	ClusterOperationVerInfo *ClusterOperationalVersionInfo `idl:"name:ppClusterOpVerInfo" json:"cluster_operation_ver_info"`
	Return                  uint32                         `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClusterVersion2Operation) OpNum() int { return 102 }

func (o *xxx_GetClusterVersion2Operation) OpName() string {
	return "/clusapi2/v2/ApiGetClusterVersion2"
}

func (o *xxx_GetClusterVersion2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClusterVersion2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_GetClusterVersion2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_GetClusterVersion2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClusterVersion2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpwMajorVersion {out} (1:{pointer=ref}*(1))(2:{alias=WORD}(uint16))
	{
		if err := w.WriteData(o.MajorVersion); err != nil {
			return err
		}
	}
	// lpwMinorVersion {out} (1:{pointer=ref}*(1))(2:{alias=WORD}(uint16))
	{
		if err := w.WriteData(o.MinorVersion); err != nil {
			return err
		}
	}
	// lpwBuildNumber {out} (1:{pointer=ref}*(1))(2:{alias=WORD}(uint16))
	{
		if err := w.WriteData(o.BuildNumber); err != nil {
			return err
		}
	}
	// lpszVendorId {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		if o.VendorID != "" {
			_ptr_lpszVendorId := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.VendorID); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.VendorID, _ptr_lpszVendorId); err != nil {
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
	// lpszCSDVersion {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		if o.CSDVersion != "" {
			_ptr_lpszCSDVersion := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.CSDVersion); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.CSDVersion, _ptr_lpszCSDVersion); err != nil {
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
	// ppClusterOpVerInfo {out} (1:{pointer=ref}*(2))(2:{alias=PCLUSTER_OPERATIONAL_VERSION_INFO}*(1))(3:{alias=CLUSTER_OPERATIONAL_VERSION_INFO}(struct))
	{
		if o.ClusterOperationVerInfo != nil {
			_ptr_ppClusterOpVerInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ClusterOperationVerInfo != nil {
					if err := o.ClusterOperationVerInfo.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ClusterOperationalVersionInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClusterOperationVerInfo, _ptr_ppClusterOpVerInfo); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClusterVersion2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpwMajorVersion {out} (1:{pointer=ref}*(1))(2:{alias=WORD}(uint16))
	{
		if err := w.ReadData(&o.MajorVersion); err != nil {
			return err
		}
	}
	// lpwMinorVersion {out} (1:{pointer=ref}*(1))(2:{alias=WORD}(uint16))
	{
		if err := w.ReadData(&o.MinorVersion); err != nil {
			return err
		}
	}
	// lpwBuildNumber {out} (1:{pointer=ref}*(1))(2:{alias=WORD}(uint16))
	{
		if err := w.ReadData(&o.BuildNumber); err != nil {
			return err
		}
	}
	// lpszVendorId {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpszVendorId := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.VendorID); err != nil {
				return err
			}
			return nil
		})
		_s_lpszVendorId := func(ptr interface{}) { o.VendorID = *ptr.(*string) }
		if err := w.ReadPointer(&o.VendorID, _s_lpszVendorId, _ptr_lpszVendorId); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpszCSDVersion {out} (1:{string, pointer=ref}*(2)*(1)[dim:0,string,null](wchar))
	{
		_ptr_lpszCSDVersion := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.CSDVersion); err != nil {
				return err
			}
			return nil
		})
		_s_lpszCSDVersion := func(ptr interface{}) { o.CSDVersion = *ptr.(*string) }
		if err := w.ReadPointer(&o.CSDVersion, _s_lpszCSDVersion, _ptr_lpszCSDVersion); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppClusterOpVerInfo {out} (1:{pointer=ref}*(2))(2:{alias=PCLUSTER_OPERATIONAL_VERSION_INFO,pointer=ref}*(1))(3:{alias=CLUSTER_OPERATIONAL_VERSION_INFO}(struct))
	{
		_ptr_ppClusterOpVerInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ClusterOperationVerInfo == nil {
				o.ClusterOperationVerInfo = &ClusterOperationalVersionInfo{}
			}
			if err := o.ClusterOperationVerInfo.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppClusterOpVerInfo := func(ptr interface{}) { o.ClusterOperationVerInfo = *ptr.(**ClusterOperationalVersionInfo) }
		if err := w.ReadPointer(&o.ClusterOperationVerInfo, _s_ppClusterOpVerInfo, _ptr_ppClusterOpVerInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetClusterVersion2Request structure represents the ApiGetClusterVersion2 operation request
type GetClusterVersion2Request struct {
}

func (o *GetClusterVersion2Request) xxx_ToOp(ctx context.Context) *xxx_GetClusterVersion2Operation {
	if o == nil {
		return &xxx_GetClusterVersion2Operation{}
	}
	return &xxx_GetClusterVersion2Operation{}
}

func (o *GetClusterVersion2Request) xxx_FromOp(ctx context.Context, op *xxx_GetClusterVersion2Operation) {
	if o == nil {
		return
	}
}
func (o *GetClusterVersion2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetClusterVersion2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClusterVersion2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClusterVersion2Response structure represents the ApiGetClusterVersion2 operation response
type GetClusterVersion2Response struct {
	MajorVersion            uint16                         `idl:"name:lpwMajorVersion" json:"major_version"`
	MinorVersion            uint16                         `idl:"name:lpwMinorVersion" json:"minor_version"`
	BuildNumber             uint16                         `idl:"name:lpwBuildNumber" json:"build_number"`
	VendorID                string                         `idl:"name:lpszVendorId;string" json:"vendor_id"`
	CSDVersion              string                         `idl:"name:lpszCSDVersion;string" json:"csd_version"`
	ClusterOperationVerInfo *ClusterOperationalVersionInfo `idl:"name:ppClusterOpVerInfo" json:"cluster_operation_ver_info"`
	// Return: The ApiGetClusterVersion2 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetClusterVersion2Response) xxx_ToOp(ctx context.Context) *xxx_GetClusterVersion2Operation {
	if o == nil {
		return &xxx_GetClusterVersion2Operation{}
	}
	return &xxx_GetClusterVersion2Operation{
		MajorVersion:            o.MajorVersion,
		MinorVersion:            o.MinorVersion,
		BuildNumber:             o.BuildNumber,
		VendorID:                o.VendorID,
		CSDVersion:              o.CSDVersion,
		ClusterOperationVerInfo: o.ClusterOperationVerInfo,
		Return:                  o.Return,
	}
}

func (o *GetClusterVersion2Response) xxx_FromOp(ctx context.Context, op *xxx_GetClusterVersion2Operation) {
	if o == nil {
		return
	}
	o.MajorVersion = op.MajorVersion
	o.MinorVersion = op.MinorVersion
	o.BuildNumber = op.BuildNumber
	o.VendorID = op.VendorID
	o.CSDVersion = op.CSDVersion
	o.ClusterOperationVerInfo = op.ClusterOperationVerInfo
	o.Return = op.Return
}
func (o *GetClusterVersion2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetClusterVersion2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClusterVersion2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateRestrictionTypeEnumOperation structure represents the ApiCreateResTypeEnum operation
type xxx_CreateRestrictionTypeEnumOperation struct {
	TypeName   string    `idl:"name:lpszTypeName;string" json:"type_name"`
	Type       uint32    `idl:"name:dwType" json:"type"`
	ReturnEnum *EnumList `idl:"name:ReturnEnum" json:"return_enum"`
	Return     uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateRestrictionTypeEnumOperation) OpNum() int { return 103 }

func (o *xxx_CreateRestrictionTypeEnumOperation) OpName() string {
	return "/clusapi2/v2/ApiCreateResTypeEnum"
}

func (o *xxx_CreateRestrictionTypeEnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRestrictionTypeEnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpszTypeName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.TypeName); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRestrictionTypeEnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpszTypeName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.TypeName); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRestrictionTypeEnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRestrictionTypeEnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ReturnEnum {out} (1:{pointer=ref}*(2))(2:{alias=PENUM_LIST}*(1))(3:{alias=ENUM_LIST}(struct))
	{
		if o.ReturnEnum != nil {
			_ptr_ReturnEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnEnum != nil {
					if err := o.ReturnEnum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&EnumList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnEnum, _ptr_ReturnEnum); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateRestrictionTypeEnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ReturnEnum {out} (1:{pointer=ref}*(2))(2:{alias=PENUM_LIST,pointer=ref}*(1))(3:{alias=ENUM_LIST}(struct))
	{
		_ptr_ReturnEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnEnum == nil {
				o.ReturnEnum = &EnumList{}
			}
			if err := o.ReturnEnum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ReturnEnum := func(ptr interface{}) { o.ReturnEnum = *ptr.(**EnumList) }
		if err := w.ReadPointer(&o.ReturnEnum, _s_ReturnEnum, _ptr_ReturnEnum); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateRestrictionTypeEnumRequest structure represents the ApiCreateResTypeEnum operation request
type CreateRestrictionTypeEnumRequest struct {
	TypeName string `idl:"name:lpszTypeName;string" json:"type_name"`
	Type     uint32 `idl:"name:dwType" json:"type"`
}

func (o *CreateRestrictionTypeEnumRequest) xxx_ToOp(ctx context.Context) *xxx_CreateRestrictionTypeEnumOperation {
	if o == nil {
		return &xxx_CreateRestrictionTypeEnumOperation{}
	}
	return &xxx_CreateRestrictionTypeEnumOperation{
		TypeName: o.TypeName,
		Type:     o.Type,
	}
}

func (o *CreateRestrictionTypeEnumRequest) xxx_FromOp(ctx context.Context, op *xxx_CreateRestrictionTypeEnumOperation) {
	if o == nil {
		return
	}
	o.TypeName = op.TypeName
	o.Type = op.Type
}
func (o *CreateRestrictionTypeEnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateRestrictionTypeEnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateRestrictionTypeEnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateRestrictionTypeEnumResponse structure represents the ApiCreateResTypeEnum operation response
type CreateRestrictionTypeEnumResponse struct {
	ReturnEnum *EnumList `idl:"name:ReturnEnum" json:"return_enum"`
	// Return: The ApiCreateResTypeEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *CreateRestrictionTypeEnumResponse) xxx_ToOp(ctx context.Context) *xxx_CreateRestrictionTypeEnumOperation {
	if o == nil {
		return &xxx_CreateRestrictionTypeEnumOperation{}
	}
	return &xxx_CreateRestrictionTypeEnumOperation{
		ReturnEnum: o.ReturnEnum,
		Return:     o.Return,
	}
}

func (o *CreateRestrictionTypeEnumResponse) xxx_FromOp(ctx context.Context, op *xxx_CreateRestrictionTypeEnumOperation) {
	if o == nil {
		return
	}
	o.ReturnEnum = op.ReturnEnum
	o.Return = op.Return
}
func (o *CreateRestrictionTypeEnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateRestrictionTypeEnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateRestrictionTypeEnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BackupClusterDatabaseOperation structure represents the ApiBackupClusterDatabase operation
type xxx_BackupClusterDatabaseOperation struct {
	PathName string `idl:"name:lpszPathName;string" json:"path_name"`
	Return   uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_BackupClusterDatabaseOperation) OpNum() int { return 104 }

func (o *xxx_BackupClusterDatabaseOperation) OpName() string {
	return "/clusapi2/v2/ApiBackupClusterDatabase"
}

func (o *xxx_BackupClusterDatabaseOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupClusterDatabaseOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpszPathName {in} (1:{string, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.PathName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupClusterDatabaseOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpszPathName {in} (1:{string, alias=LPCWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.PathName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupClusterDatabaseOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupClusterDatabaseOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupClusterDatabaseOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BackupClusterDatabaseRequest structure represents the ApiBackupClusterDatabase operation request
type BackupClusterDatabaseRequest struct {
	PathName string `idl:"name:lpszPathName;string" json:"path_name"`
}

func (o *BackupClusterDatabaseRequest) xxx_ToOp(ctx context.Context) *xxx_BackupClusterDatabaseOperation {
	if o == nil {
		return &xxx_BackupClusterDatabaseOperation{}
	}
	return &xxx_BackupClusterDatabaseOperation{
		PathName: o.PathName,
	}
}

func (o *BackupClusterDatabaseRequest) xxx_FromOp(ctx context.Context, op *xxx_BackupClusterDatabaseOperation) {
	if o == nil {
		return
	}
	o.PathName = op.PathName
}
func (o *BackupClusterDatabaseRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *BackupClusterDatabaseRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupClusterDatabaseOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BackupClusterDatabaseResponse structure represents the ApiBackupClusterDatabase operation response
type BackupClusterDatabaseResponse struct {
	// Return: The ApiBackupClusterDatabase return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BackupClusterDatabaseResponse) xxx_ToOp(ctx context.Context) *xxx_BackupClusterDatabaseOperation {
	if o == nil {
		return &xxx_BackupClusterDatabaseOperation{}
	}
	return &xxx_BackupClusterDatabaseOperation{
		Return: o.Return,
	}
}

func (o *BackupClusterDatabaseResponse) xxx_FromOp(ctx context.Context, op *xxx_BackupClusterDatabaseOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BackupClusterDatabaseResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *BackupClusterDatabaseResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupClusterDatabaseOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_NodeClusterControlOperation structure represents the ApiNodeClusterControl operation
type xxx_NodeClusterControlOperation struct {
	Cluster        *Cluster `idl:"name:hCluster" json:"cluster"`
	HostNode       *Node    `idl:"name:hHostNode" json:"host_node"`
	ControlCode    uint32   `idl:"name:dwControlCode" json:"control_code"`
	InBuffer       []byte   `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize   uint32   `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBuffer      []byte   `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	OutBufferSize  uint32   `idl:"name:nOutBufferSize" json:"out_buffer_size"`
	BytesReturned  uint32   `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32   `idl:"name:lpcbRequired" json:"required_length"`
	Return         uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_NodeClusterControlOperation) OpNum() int { return 105 }

func (o *xxx_NodeClusterControlOperation) OpName() string {
	return "/clusapi2/v2/ApiNodeClusterControl"
}

func (o *xxx_NodeClusterControlOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InBuffer != nil && o.InBufferSize == 0 {
		o.InBufferSize = uint32(len(o.InBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeClusterControlOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hCluster {in} (1:{context_handle, alias=HCLUSTER_RPC, names=ndr_context_handle}(struct))
	{
		if o.Cluster != nil {
			if err := o.Cluster.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Cluster{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// hHostNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.HostNode != nil {
			if err := o.HostNode.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Node{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		if o.InBuffer != nil || o.InBufferSize > 0 {
			_ptr_lpInBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.InBufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.InBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.InBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.InBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InBuffer, _ptr_lpInBuffer); err != nil {
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
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeClusterControlOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hCluster {in} (1:{context_handle, alias=HCLUSTER_RPC, names=ndr_context_handle}(struct))
	{
		if o.Cluster == nil {
			o.Cluster = &Cluster{}
		}
		if err := o.Cluster.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// hHostNode {in} (1:{context_handle, alias=HNODE_RPC, names=ndr_context_handle}(struct))
	{
		if o.HostNode == nil {
			o.HostNode = &Node{}
		}
		if err := o.HostNode.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		_ptr_lpInBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.InBuffer", sizeInfo[0])
			}
			o.InBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.InBuffer {
				i1 := i1
				if err := w.ReadData(&o.InBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_lpInBuffer := func(ptr interface{}) { o.InBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.InBuffer, _s_lpInBuffer, _ptr_lpInBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeClusterControlOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.OutBuffer != nil && o.BytesReturned == 0 {
		o.BytesReturned = uint32(len(o.OutBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeClusterControlOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		dimSize1 := uint64(o.OutBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.BytesReturned)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.OutBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.OutBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.OutBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_NodeClusterControlOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OutBuffer", sizeInfo[0])
		}
		o.OutBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.OutBuffer {
			i1 := i1
			if err := w.ReadData(&o.OutBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// NodeClusterControlRequest structure represents the ApiNodeClusterControl operation request
type NodeClusterControlRequest struct {
	Cluster       *Cluster `idl:"name:hCluster" json:"cluster"`
	HostNode      *Node    `idl:"name:hHostNode" json:"host_node"`
	ControlCode   uint32   `idl:"name:dwControlCode" json:"control_code"`
	InBuffer      []byte   `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize  uint32   `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBufferSize uint32   `idl:"name:nOutBufferSize" json:"out_buffer_size"`
}

func (o *NodeClusterControlRequest) xxx_ToOp(ctx context.Context) *xxx_NodeClusterControlOperation {
	if o == nil {
		return &xxx_NodeClusterControlOperation{}
	}
	return &xxx_NodeClusterControlOperation{
		Cluster:       o.Cluster,
		HostNode:      o.HostNode,
		ControlCode:   o.ControlCode,
		InBuffer:      o.InBuffer,
		InBufferSize:  o.InBufferSize,
		OutBufferSize: o.OutBufferSize,
	}
}

func (o *NodeClusterControlRequest) xxx_FromOp(ctx context.Context, op *xxx_NodeClusterControlOperation) {
	if o == nil {
		return
	}
	o.Cluster = op.Cluster
	o.HostNode = op.HostNode
	o.ControlCode = op.ControlCode
	o.InBuffer = op.InBuffer
	o.InBufferSize = op.InBufferSize
	o.OutBufferSize = op.OutBufferSize
}
func (o *NodeClusterControlRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *NodeClusterControlRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NodeClusterControlOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// NodeClusterControlResponse structure represents the ApiNodeClusterControl operation response
type NodeClusterControlResponse struct {
	OutBuffer      []byte `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	BytesReturned  uint32 `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32 `idl:"name:lpcbRequired" json:"required_length"`
	// Return: The ApiNodeClusterControl return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *NodeClusterControlResponse) xxx_ToOp(ctx context.Context) *xxx_NodeClusterControlOperation {
	if o == nil {
		return &xxx_NodeClusterControlOperation{}
	}
	return &xxx_NodeClusterControlOperation{
		OutBuffer:      o.OutBuffer,
		BytesReturned:  o.BytesReturned,
		RequiredLength: o.RequiredLength,
		Return:         o.Return,
	}
}

func (o *NodeClusterControlResponse) xxx_FromOp(ctx context.Context, op *xxx_NodeClusterControlOperation) {
	if o == nil {
		return
	}
	o.OutBuffer = op.OutBuffer
	o.BytesReturned = op.BytesReturned
	o.RequiredLength = op.RequiredLength
	o.Return = op.Return
}
func (o *NodeClusterControlResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *NodeClusterControlResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_NodeClusterControlOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ClusterControlOperation structure represents the ApiClusterControl operation
type xxx_ClusterControlOperation struct {
	Cluster        *Cluster `idl:"name:hCluster" json:"cluster"`
	ControlCode    uint32   `idl:"name:dwControlCode" json:"control_code"`
	InBuffer       []byte   `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize   uint32   `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBuffer      []byte   `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	OutBufferSize  uint32   `idl:"name:nOutBufferSize" json:"out_buffer_size"`
	BytesReturned  uint32   `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32   `idl:"name:lpcbRequired" json:"required_length"`
	Return         uint32   `idl:"name:Return" json:"return"`
}

func (o *xxx_ClusterControlOperation) OpNum() int { return 106 }

func (o *xxx_ClusterControlOperation) OpName() string { return "/clusapi2/v2/ApiClusterControl" }

func (o *xxx_ClusterControlOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.InBuffer != nil && o.InBufferSize == 0 {
		o.InBufferSize = uint32(len(o.InBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClusterControlOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hCluster {in} (1:{context_handle, alias=HCLUSTER_RPC, names=ndr_context_handle}(struct))
	{
		if o.Cluster != nil {
			if err := o.Cluster.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Cluster{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		if o.InBuffer != nil || o.InBufferSize > 0 {
			_ptr_lpInBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.InBufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.InBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.InBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.InBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.InBuffer, _ptr_lpInBuffer); err != nil {
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
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClusterControlOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hCluster {in} (1:{context_handle, alias=HCLUSTER_RPC, names=ndr_context_handle}(struct))
	{
		if o.Cluster == nil {
			o.Cluster = &Cluster{}
		}
		if err := o.Cluster.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwControlCode {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ControlCode); err != nil {
			return err
		}
	}
	// lpInBuffer {in} (1:{pointer=unique}*(1))(2:{alias=UCHAR}[dim:0,size_is=nInBufferSize](uchar))
	{
		_ptr_lpInBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.InBuffer", sizeInfo[0])
			}
			o.InBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.InBuffer {
				i1 := i1
				if err := w.ReadData(&o.InBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_lpInBuffer := func(ptr interface{}) { o.InBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.InBuffer, _s_lpInBuffer, _ptr_lpInBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// nInBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InBufferSize); err != nil {
			return err
		}
	}
	// nOutBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OutBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClusterControlOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.OutBuffer != nil && o.BytesReturned == 0 {
		o.BytesReturned = uint32(len(o.OutBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClusterControlOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		dimSize1 := uint64(o.OutBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.BytesReturned)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.OutBuffer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.OutBuffer[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.OutBuffer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClusterControlOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpOutBuffer {out} (1:{pointer=ref}*(1))(2:{alias=UCHAR}[dim:0,size_is=nOutBufferSize,length_is=lpBytesReturned](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.OutBuffer", sizeInfo[0])
		}
		o.OutBuffer = make([]byte, sizeInfo[0])
		for i1 := range o.OutBuffer {
			i1 := i1
			if err := w.ReadData(&o.OutBuffer[i1]); err != nil {
				return err
			}
		}
	}
	// lpBytesReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BytesReturned); err != nil {
			return err
		}
	}
	// lpcbRequired {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequiredLength); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ClusterControlRequest structure represents the ApiClusterControl operation request
type ClusterControlRequest struct {
	Cluster       *Cluster `idl:"name:hCluster" json:"cluster"`
	ControlCode   uint32   `idl:"name:dwControlCode" json:"control_code"`
	InBuffer      []byte   `idl:"name:lpInBuffer;size_is:(nInBufferSize);pointer:unique" json:"in_buffer"`
	InBufferSize  uint32   `idl:"name:nInBufferSize" json:"in_buffer_size"`
	OutBufferSize uint32   `idl:"name:nOutBufferSize" json:"out_buffer_size"`
}

func (o *ClusterControlRequest) xxx_ToOp(ctx context.Context) *xxx_ClusterControlOperation {
	if o == nil {
		return &xxx_ClusterControlOperation{}
	}
	return &xxx_ClusterControlOperation{
		Cluster:       o.Cluster,
		ControlCode:   o.ControlCode,
		InBuffer:      o.InBuffer,
		InBufferSize:  o.InBufferSize,
		OutBufferSize: o.OutBufferSize,
	}
}

func (o *ClusterControlRequest) xxx_FromOp(ctx context.Context, op *xxx_ClusterControlOperation) {
	if o == nil {
		return
	}
	o.Cluster = op.Cluster
	o.ControlCode = op.ControlCode
	o.InBuffer = op.InBuffer
	o.InBufferSize = op.InBufferSize
	o.OutBufferSize = op.OutBufferSize
}
func (o *ClusterControlRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ClusterControlRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClusterControlOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ClusterControlResponse structure represents the ApiClusterControl operation response
type ClusterControlResponse struct {
	OutBuffer      []byte `idl:"name:lpOutBuffer;size_is:(nOutBufferSize);length_is:(lpBytesReturned)" json:"out_buffer"`
	BytesReturned  uint32 `idl:"name:lpBytesReturned" json:"bytes_returned"`
	RequiredLength uint32 `idl:"name:lpcbRequired" json:"required_length"`
	// Return: The ApiClusterControl return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ClusterControlResponse) xxx_ToOp(ctx context.Context) *xxx_ClusterControlOperation {
	if o == nil {
		return &xxx_ClusterControlOperation{}
	}
	return &xxx_ClusterControlOperation{
		OutBuffer:      o.OutBuffer,
		BytesReturned:  o.BytesReturned,
		RequiredLength: o.RequiredLength,
		Return:         o.Return,
	}
}

func (o *ClusterControlResponse) xxx_FromOp(ctx context.Context, op *xxx_ClusterControlOperation) {
	if o == nil {
		return
	}
	o.OutBuffer = op.OutBuffer
	o.BytesReturned = op.BytesReturned
	o.RequiredLength = op.RequiredLength
	o.Return = op.Return
}
func (o *ClusterControlResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ClusterControlResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClusterControlOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_UnblockGetNotifyCallOperation structure represents the ApiUnblockGetNotifyCall operation
type xxx_UnblockGetNotifyCallOperation struct {
	Notify *Notify `idl:"name:hNotify" json:"notify"`
	Return uint32  `idl:"name:Return" json:"return"`
}

func (o *xxx_UnblockGetNotifyCallOperation) OpNum() int { return 107 }

func (o *xxx_UnblockGetNotifyCallOperation) OpName() string {
	return "/clusapi2/v2/ApiUnblockGetNotifyCall"
}

func (o *xxx_UnblockGetNotifyCallOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnblockGetNotifyCallOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify != nil {
			if err := o.Notify.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Notify{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_UnblockGetNotifyCallOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hNotify {in} (1:{context_handle, alias=HNOTIFY_RPC, names=ndr_context_handle}(struct))
	{
		if o.Notify == nil {
			o.Notify = &Notify{}
		}
		if err := o.Notify.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnblockGetNotifyCallOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnblockGetNotifyCallOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_UnblockGetNotifyCallOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// UnblockGetNotifyCallRequest structure represents the ApiUnblockGetNotifyCall operation request
type UnblockGetNotifyCallRequest struct {
	Notify *Notify `idl:"name:hNotify" json:"notify"`
}

func (o *UnblockGetNotifyCallRequest) xxx_ToOp(ctx context.Context) *xxx_UnblockGetNotifyCallOperation {
	if o == nil {
		return &xxx_UnblockGetNotifyCallOperation{}
	}
	return &xxx_UnblockGetNotifyCallOperation{
		Notify: o.Notify,
	}
}

func (o *UnblockGetNotifyCallRequest) xxx_FromOp(ctx context.Context, op *xxx_UnblockGetNotifyCallOperation) {
	if o == nil {
		return
	}
	o.Notify = op.Notify
}
func (o *UnblockGetNotifyCallRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *UnblockGetNotifyCallRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnblockGetNotifyCallOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// UnblockGetNotifyCallResponse structure represents the ApiUnblockGetNotifyCall operation response
type UnblockGetNotifyCallResponse struct {
	// Return: The ApiUnblockGetNotifyCall return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *UnblockGetNotifyCallResponse) xxx_ToOp(ctx context.Context) *xxx_UnblockGetNotifyCallOperation {
	if o == nil {
		return &xxx_UnblockGetNotifyCallOperation{}
	}
	return &xxx_UnblockGetNotifyCallOperation{
		Return: o.Return,
	}
}

func (o *UnblockGetNotifyCallResponse) xxx_FromOp(ctx context.Context, op *xxx_UnblockGetNotifyCallOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *UnblockGetNotifyCallResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *UnblockGetNotifyCallResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_UnblockGetNotifyCallOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetServiceAccountPasswordOperation structure represents the ApiSetServiceAccountPassword operation
type xxx_SetServiceAccountPasswordOperation struct {
	NewPassword               string                      `idl:"name:lpszNewPassword;string" json:"new_password"`
	Flags                     ClusterSetPasswordFlags     `idl:"name:dwFlags" json:"flags"`
	ReturnStatusBufferPointer []*ClusterSetPasswordStatus `idl:"name:ReturnStatusBufferPtr;size_is:(ReturnStatusBufferSize);length_is:(SizeReturned)" json:"return_status_buffer_pointer"`
	ReturnStatusBufferSize    uint32                      `idl:"name:ReturnStatusBufferSize" json:"return_status_buffer_size"`
	SizeReturned              uint32                      `idl:"name:SizeReturned" json:"size_returned"`
	ExpectedBufferSize        uint32                      `idl:"name:ExpectedBufferSize" json:"expected_buffer_size"`
	Return                    uint32                      `idl:"name:Return" json:"return"`
}

func (o *xxx_SetServiceAccountPasswordOperation) OpNum() int { return 108 }

func (o *xxx_SetServiceAccountPasswordOperation) OpName() string {
	return "/clusapi2/v2/ApiSetServiceAccountPassword"
}

func (o *xxx_SetServiceAccountPasswordOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetServiceAccountPasswordOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// lpszNewPassword {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.NewPassword); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=IDL_CLUSTER_SET_PASSWORD_FLAGS}(enum))
	{
		if err := w.WriteData(uint16(o.Flags)); err != nil {
			return err
		}
	}
	// ReturnStatusBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ReturnStatusBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetServiceAccountPasswordOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// lpszNewPassword {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.NewPassword); err != nil {
			return err
		}
	}
	// dwFlags {in} (1:{alias=IDL_CLUSTER_SET_PASSWORD_FLAGS}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Flags)); err != nil {
			return err
		}
	}
	// ReturnStatusBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ReturnStatusBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetServiceAccountPasswordOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.ReturnStatusBufferPointer != nil && o.SizeReturned == 0 {
		o.SizeReturned = uint32(len(o.ReturnStatusBufferPointer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetServiceAccountPasswordOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ReturnStatusBufferPtr {out} (1:[dim:0,size_is=ReturnStatusBufferSize,length_is=SizeReturned])(2:{alias=IDL_CLUSTER_SET_PASSWORD_STATUS}(struct))
	{
		dimSize1 := uint64(o.ReturnStatusBufferSize)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.SizeReturned)
		if dimLength1 > sizeInfo[0] {
			dimLength1 = sizeInfo[0]
		} else {
			sizeInfo[0] = dimLength1
		}
		if err := w.WriteSize(0); err != nil {
			return err
		}
		if err := w.WriteSize(dimLength1); err != nil {
			return err
		}
		for i1 := range o.ReturnStatusBufferPointer {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.ReturnStatusBufferPointer[i1] != nil {
				if err := o.ReturnStatusBufferPointer[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ClusterSetPasswordStatus{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.ReturnStatusBufferPointer); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&ClusterSetPasswordStatus{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// SizeReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SizeReturned); err != nil {
			return err
		}
	}
	// ExpectedBufferSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ExpectedBufferSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetServiceAccountPasswordOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ReturnStatusBufferPtr {out} (1:[dim:0,size_is=ReturnStatusBufferSize,length_is=SizeReturned])(2:{alias=IDL_CLUSTER_SET_PASSWORD_STATUS}(struct))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ReturnStatusBufferPointer", sizeInfo[0])
		}
		o.ReturnStatusBufferPointer = make([]*ClusterSetPasswordStatus, sizeInfo[0])
		for i1 := range o.ReturnStatusBufferPointer {
			i1 := i1
			if o.ReturnStatusBufferPointer[i1] == nil {
				o.ReturnStatusBufferPointer[i1] = &ClusterSetPasswordStatus{}
			}
			if err := o.ReturnStatusBufferPointer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// SizeReturned {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SizeReturned); err != nil {
			return err
		}
	}
	// ExpectedBufferSize {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ExpectedBufferSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetServiceAccountPasswordRequest structure represents the ApiSetServiceAccountPassword operation request
type SetServiceAccountPasswordRequest struct {
	NewPassword            string                  `idl:"name:lpszNewPassword;string" json:"new_password"`
	Flags                  ClusterSetPasswordFlags `idl:"name:dwFlags" json:"flags"`
	ReturnStatusBufferSize uint32                  `idl:"name:ReturnStatusBufferSize" json:"return_status_buffer_size"`
}

func (o *SetServiceAccountPasswordRequest) xxx_ToOp(ctx context.Context) *xxx_SetServiceAccountPasswordOperation {
	if o == nil {
		return &xxx_SetServiceAccountPasswordOperation{}
	}
	return &xxx_SetServiceAccountPasswordOperation{
		NewPassword:            o.NewPassword,
		Flags:                  o.Flags,
		ReturnStatusBufferSize: o.ReturnStatusBufferSize,
	}
}

func (o *SetServiceAccountPasswordRequest) xxx_FromOp(ctx context.Context, op *xxx_SetServiceAccountPasswordOperation) {
	if o == nil {
		return
	}
	o.NewPassword = op.NewPassword
	o.Flags = op.Flags
	o.ReturnStatusBufferSize = op.ReturnStatusBufferSize
}
func (o *SetServiceAccountPasswordRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetServiceAccountPasswordRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetServiceAccountPasswordOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetServiceAccountPasswordResponse structure represents the ApiSetServiceAccountPassword operation response
type SetServiceAccountPasswordResponse struct {
	ReturnStatusBufferPointer []*ClusterSetPasswordStatus `idl:"name:ReturnStatusBufferPtr;size_is:(ReturnStatusBufferSize);length_is:(SizeReturned)" json:"return_status_buffer_pointer"`
	SizeReturned              uint32                      `idl:"name:SizeReturned" json:"size_returned"`
	ExpectedBufferSize        uint32                      `idl:"name:ExpectedBufferSize" json:"expected_buffer_size"`
	// Return: The ApiSetServiceAccountPassword return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetServiceAccountPasswordResponse) xxx_ToOp(ctx context.Context) *xxx_SetServiceAccountPasswordOperation {
	if o == nil {
		return &xxx_SetServiceAccountPasswordOperation{}
	}
	return &xxx_SetServiceAccountPasswordOperation{
		ReturnStatusBufferPointer: o.ReturnStatusBufferPointer,
		SizeReturned:              o.SizeReturned,
		ExpectedBufferSize:        o.ExpectedBufferSize,
		Return:                    o.Return,
	}
}

func (o *SetServiceAccountPasswordResponse) xxx_FromOp(ctx context.Context, op *xxx_SetServiceAccountPasswordOperation) {
	if o == nil {
		return
	}
	o.ReturnStatusBufferPointer = op.ReturnStatusBufferPointer
	o.SizeReturned = op.SizeReturned
	o.ExpectedBufferSize = op.ExpectedBufferSize
	o.Return = op.Return
}
func (o *SetServiceAccountPasswordResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetServiceAccountPasswordResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetServiceAccountPasswordOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
