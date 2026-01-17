package clusapi3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
)

// clusapi3 server interface.
type Clusapi3Server interface {

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

	// (Protocol Version 3) The ApiSetResourceDependencyExpression method instructs the
	// server to set the dependency relationship for the resource that is identified by
	// hResource to the complex dependency, as specified in section 3.1.1.1.2, using dependency
	// expression represented by lpszDependencyExpression. For successful completion of
	// the method, the server MUST add the dependency information to the nonvolatile cluster
	// state.
	//
	// Servers MUST maintain complex resource dependencies as nonvolatile configuration
	// data in their cluster state.
	//
	// Return Values: The method MUST return the following error codes for the specified
	// conditions.
	//
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                   RETURN                   |                                                                                  |
	//	|                 VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                            |                                                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                   | Success.                                                                         |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000006 ERROR_INVALID_HANDLE            | The data that is pointed to by the hResource parameter does not represent a      |
	//	|                                            | valid HRES_RPC context handle.                                                   |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER         | See the preceding text for details of when this error can occur.                 |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000423 ERROR_CIRCULAR_DEPENDENCY       | See the preceding text for details of when this error can occur.                 |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000138B ERROR_DEPENDENCY_ALREADY_EXISTS | See the preceding text for details of when this error can occur.                 |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000139B ERROR_RESOURCE_ONLINE           | See the preceding text for details of when this error can occur.                 |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000013CD ERROR_DEPENDENCY_NOT_ALLOWED    | Cannot depend on quorum resource.                                                |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//
	// For any other condition, this method MUST return a value that is not one of the values
	// listed in the preceding table. The client MUST behave in one consistent, identical
	// manner for all values that are not listed in the preceding table. The client SHOULD
	// treat errors specified in section 3.2.4.6 as recoverable errors and initiate the
	// reconnect procedure as specified in section 3.2.4.6.
	SetResourceDependencyExpression(context.Context, *SetResourceDependencyExpressionRequest) (*SetResourceDependencyExpressionResponse, error)

	// (Protocol Version 3) The ApiGetResourceDependencyExpression method instructs the
	// server to retrieve the complex dependency expression, as specified in section 3.1.1.1.2,
	// for the resource represented by hResource.
	//
	// The server SHOULD accept an ApiGetResourceDependencyExpression request if its protocol
	// server state is read-only, and the server MUST accept an ApiGetResourceDependencyExpression
	// request for processing if it is in the read/write state, as specified in section
	// 3.1.1.
	//
	// The server SHOULD accept an ApiGetResourceDependencyExpression request if the access
	// level associated with the hResource context handle is at least "Read" (section 3.1.4).
	//
	// Return Values: The method MUST return the following error codes for the specified
	// conditions.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | Success.                                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000006 ERROR_INVALID_HANDLE | The data that is pointed to by the hResource parameter does not represent a      |
	//	|                                 | valid HRES_RPC context handle.                                                   |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// For any other condition, this method MUST return a value that is not one of the values
	// listed in the preceding table. The client MUST behave in one consistent, identical
	// manner for all values that are not listed in the preceding table. The client SHOULD
	// treat errors as recoverable errors and initiate the reconnect procedure as specified
	// in section 3.2.4.6.
	GetResourceDependencyExpression(context.Context, *GetResourceDependencyExpressionRequest) (*GetResourceDependencyExpressionResponse, error)

	// Opnum111NotUsedOnWire operation.
	// Opnum111NotUsedOnWire

	// (Protocol Version 3) The ApiGetResourceNetworkName method retrieves the NetBIOS computer
	// name associated with the resource upon which the designated resource depends and
	// whose resource type, as specified in section 3.1.4.2.16, matches the Unicode string
	// "Network Name".
	//
	// Return Values: The method MUST return the following error codes for the specified
	// conditions.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS              | Success.                                                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000006 ERROR_INVALID_HANDLE       | The data that is pointed to by the hResource parameter does not represent a      |
	//	|                                       | valid HRES_RPC context handle.                                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000138A ERROR_DEPENDENCY_NOT_FOUND | A resource that has the resource type that matches the Unicode string "Network   |
	//	|                                       | Name" was not found in any dependency chains that start from the designated      |
	//	|                                       | resource.                                                                        |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//
	// For any other condition, this method MUST return a value that is not one of the values
	// listed in the preceding table. The client MUST behave in one consistent, identical
	// manner for all values that are not listed in the preceding table. The client SHOULD
	// treat errors specified in 3.2.4.6 as recoverable errors and initiate the reconnect
	// procedure as specified in section 3.2.4.6.
	GetResourceNetworkName(context.Context, *GetResourceNetworkNameRequest) (*GetResourceNetworkNameResponse, error)

	// (Protocol Version 3) The ApiExecuteBatch method instructs the server to perform a
	// transacted series of modifications, relative to the designated key, to the cluster
	// registry. The contents of the registry are changed only if all the modifications
	// succeed; the first failure causes all previous modifications to be backed out.
	//
	// The modifications are described in a CLUSTER_REG_BATCH_UPDATE structure, as described
	// in section 2.2.3.17. Each batch update command instructs the server as to how the
	// registry is modified.
	//
	// Return Values: The method MUST return the following error codes for the specified
	// conditions.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | Success.                                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000006 ERROR_INVALID_HANDLE    | The data that is pointed to by the hKey parameter does not represent a valid     |
	//	|                                    | HKEY_RPC context handle.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A CLUSREG_SET_VALUE or CLUSREG_DELETE_VALUE command has been encountered prior   |
	//	|                                    | to a CLUSREG_CREATE_KEY command.                                                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// For any other condition, this method MUST return a value that is not one of the values
	// listed in the preceding table. The client MUST behave in one consistent, identical
	// manner for all values that are not listed in the preceding table. The client SHOULD
	// treat errors specified in section 3.2.4.6 as recoverable errors and initiate the
	// reconnect procedure as specified in section 3.2.4.6.
	ExecuteBatch(context.Context, *ExecuteBatchRequest) (*ExecuteBatchResponse, error)

	// (Protocol Version 3) The ApiCreateBatchPort method establishes context on the server
	// about client interaction with a cluster registry batch update notification port by
	// means of the current RPC connection. ApiCreateBatchPort returns a context handle
	// so that the client can refer to the context that is created in subsequent method
	// invocations.
	//
	// The server SHOULD accept an ApiCreateBatchPort request if the client's access level
	// is at least "Read" (section 3.1.4).
	//
	// Return Values: The method MUST return the following error codes for the specified
	// conditions.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | Success.                                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000006 ERROR_INVALID_HANDLE | The data that is pointed to by the hKey parameter does not represent a valid     |
	//	|                                 | HKEY_RPC context handle.                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// For any other condition, this method MUST return a value that is not one of the values
	// listed in the preceding table. The client MUST behave in one consistent, identical
	// manner for all values that are not listed in the preceding table. The client SHOULD
	// treat errors as recoverable errors, and initiate the reconnect procedure as specified
	// in section 3.2.4.6.
	CreateBatchPort(context.Context, *CreateBatchPortRequest) (*CreateBatchPortResponse, error)

	// (Protocol Version 3) The ApiGetBatchNotification method instructs the server to retrieve
	// the first queued batch update event indication from the designated batch update notification_port.
	// The server MUST NOT complete this method until an indication has been queued to the
	// port or the port has been closed through a separate call to the ApiCloseBatchPort
	// method.
	//
	// Upon successful completion of this method, the server MUST write a CLUSTER_REG_BATCH_UPDATE
	// structure, as specified in section 2.2.3.17, indicating the modifications that have
	// been made to the cluster registry. This information essentially mirrors the CLUSTER_REG_BATCH_UPDATE
	// structure provided to the ApiExecuteBatch method with the modification that a CLUSREG_VALUE_DELETED
	// command precedes every CLUSREG_SET_VALUE and CLUSREG_DELETE_VALUE command in the
	// returned notification data if the value has existing data. The Name field identifies
	// the name of the value that was modified, and the Data field contains the value data
	// that existed prior to executing the aforementioned value command.
	//
	// For example, the client calls ApiExecuteBatch with the following series of commands
	// for a value named "NotifyTest" that does not already exist in the registry:
	//
	// * Delete Value
	//
	// * Set Value to "hello world"
	//
	// * Set Value to "hello universe"
	//
	// * Delete Value
	//
	// The series of change notifications that the server returns to the client through
	// this method are as follows:
	//
	// * Delete Value
	//
	// * Set Value to "hello world"
	//
	// * Value Deleted with the data set to "hello world"
	//
	// * Set Value to "hello universe"
	//
	// * Value Deleted with the data set to "hello universe"
	//
	// * Delete Value
	//
	// The following diagrams illustrate the preceding command and notification sequence.
	//
	// Return Values: The method MUST return the following error codes for the specified
	// conditions.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | Success.                                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000006 ERROR_INVALID_HANDLE | The data that is pointed to by the hBatchNotify parameter does not represent a   |
	//	|                                 | valid HBATCH_PORT_RPC context handle.                                            |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS  | The port referenced by the hBatchNotify parameter has been closed by a separate  |
	//	|                                 | call to the ApiCloseBatchPort method.                                            |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// For any other condition, this method MUST return a value that is not one of those
	// listed in the preceding table. The client MUST behave in one consistent, identical
	// manner for all values that are not listed in the preceding table.
	GetBatchNotification(context.Context, *GetBatchNotificationRequest) (*GetBatchNotificationResponse, error)

	// (Protocol Version 3) The ApiCloseBatchPort method instructs the server to free any
	// context information that is created in a previous ApiCreateBatchPort method.
	//
	// The server SHOULD accept an ApiCloseBatchPort request if its protocol server state
	// is read-only, and the server MUST accept the request for processing if it is in the
	// read/write state, as specified in section 3.1.1.
	//
	// The server SHOULD accept an ApiCloseBatchPort request if the client's access level
	// is at least "Read" (section 3.1.4).
	//
	// Return Values: The method MUST return the following error codes for the specified
	// conditions.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | Success.                                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000006 ERROR_INVALID_HANDLE | The data that is pointed to by the phBatchPort parameter does not represent a    |
	//	|                                 | valid HBATCH_PORT_RPC context handle.                                            |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// For any other condition, this method MUST return a value that is not one of the values
	// listed in the preceding table. The client MUST behave in one consistent, identical
	// manner for all values that are not listed in the preceding table.
	CloseBatchPort(context.Context, *CloseBatchPortRequest) (*CloseBatchPortResponse, error)

	// (Protocol Version 3) The ApiOpenClusterEx method SHOULD<96> establish context on
	// the server about client interaction with the cluster by means of the current RPC
	// connection. ApiOpenClusterEx returns a context handle so that the client can refer
	// to the context that is created in subsequent method calls.
	//
	// Return Values: The method MUST return a valid HCLUSTER_RPC (section 2.2.1.1) context
	// handle to indicate success; otherwise, it MUST return NULL.
	OpenClusterEx(context.Context, *OpenClusterExRequest) (*OpenClusterExResponse, error)

	// (Protocol Version 3) The ApiOpenNodeEx method SHOULD<97> establish context on the
	// server about the interaction of a client with the specified cluster node by using
	// the current RPC connection. ApiOpenNodeEx returns a context handle so that the client
	// can refer to the context that is created in subsequent method calls.
	//
	// Return Values: The method MUST return a valid HNODE_RPC (section 2.2.1.2) context
	// handle to indicate success; otherwise, it MUST return NULL.
	OpenNodeEx(context.Context, *OpenNodeExRequest) (*OpenNodeExResponse, error)

	// (Protocol Version 3) The ApiOpenGroupEx method SHOULD<98> establish context on the
	// server about the interaction of a client with a specified cluster group by means
	// of the current RPC connection. ApiOpenGroupEx returns a context handle so that the
	// client can refer to the group in subsequent method calls.
	//
	// Return Values: The method MUST return a valid HGROUP_RPC context handle (section
	// 2.2.1.3) to indicate success; otherwise, it MUST return NULL.
	OpenGroupEx(context.Context, *OpenGroupExRequest) (*OpenGroupExResponse, error)

	// (Protocol Version 3) The ApiOpenResourceEx method SHOULD<99> establish context on
	// the server about the interaction of a client with the specified cluster resource
	// by using the current RPC connection. ApiOpenResourceEx returns a context handle so
	// that the client can refer to the resource in subsequent method calls.
	//
	// Return Values: The method MUST return a valid HRES_RPC (section 2.2.1.4) context
	// handle to indicate success; otherwise, it MUST return NULL.
	OpenResourceEx(context.Context, *OpenResourceExRequest) (*OpenResourceExResponse, error)

	// (Protocol Version 3) The ApiOpenNetworkEx method SHOULD<100> establish context on
	// the server about the interaction of a client with the specified cluster network by
	// using the current RPC connection. ApiOpenNetworkEx returns a context handle so that
	// the client can refer to the context that is created in subsequent method calls.
	//
	// Return Values: The method MUST return a valid HNETWORK_RPC (section 2.2.1.7) context
	// handle to indicate success; otherwise, it MUST return NULL.
	OpenNetworkEx(context.Context, *OpenNetworkExRequest) (*OpenNetworkExResponse, error)

	// (Protocol Version 3) The ApiOpenNetInterfaceEx method SHOULD<101> establish context
	// on the server about the interaction of a client with the specified cluster network
	// interface by using the current RPC connection. ApiOpenNetInterfaceEx returns a context
	// handle so that the client can refer to the context that is created in subsequent
	// method calls.
	//
	// Return Values: The method MUST return a valid HNETINTERFACE_RPC (section 2.2.1.8)
	// context handle to indicate success; otherwise, it MUST return NULL.
	OpenNetInterfaceEx(context.Context, *OpenNetInterfaceExRequest) (*OpenNetInterfaceExResponse, error)

	// (Protocol Version 3) The ApiChangeCsvState method SHOULD<102> instruct the server
	// to change the accessibility of the disk associated with hResource.
	//
	// If dwState is 1, the server MUST set ResourceSharedVolumes to TRUE and convert all
	// volumes associated with hResource to cluster shared volumes. The server MUST set
	// the initial state of all cluster shared volumes associated with hResource such that
	// volume maintenance mode, redirected mode, and backup mode are all disabled.
	//
	// If dwState is 1, the server SHOULD also designate the group associated with hResource
	// as a special group, as specified in section 3.1.1.1.4.
	//
	// If dwState is 0, the server MUST set ResourceSharedVolumes to FALSE and stop making
	// the volumes associated with hResource accessible to all nodes as cluster shared volumes.
	//
	// If dwState is 0, the server SHOULD also remove the special group designation of the
	// group associated with hResource.
	//
	// The server SHOULD accept an ApiChangeCsvState request if its protocol server state
	// is read-only, and the server MUST accept the request for processing if it is in the
	// read/write state, as specified in section 3.1.1.
	//
	// Return Values: The method MUST return the following error codes for the specified
	// conditions.
	//
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                     RETURN                     |                                                                                  |
	//	|                   VALUE/CODE                   |                                   DESCRIPTION                                    |
	//	|                                                |                                                                                  |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                       | Success.                                                                         |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000003E5 ERROR_IO_PENDING                    | The operation is still in progress.                                              |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000046 ERROR_SHARING_PAUSED                | The current protocol server state of the server is not read/write.               |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000006D1 RPC_S_PROCNUM_OUT_OF_RANGE          | The server does not support this method.                                         |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000013B8 ERROR_CLUSTER_INVALID_REQUEST       | The operation is invalid for the cluster or for the specified resource . It is   |
	//	|                                                | invalid for the cluster if the dwState parameter is 1 and the requested state is |
	//	|                                                | not enabled for the cluster; for instance, the server EnableSharedVolumes state  |
	//	|                                                | is FALSE (indicating that the server does not support cluster shared volumes)    |
	//	|                                                | (see section 3.1.1.4). The operation is invalid for the specified resource if    |
	//	|                                                | any of the following conditions are met: The dwState parameter is 1, and the     |
	//	|                                                | specified resource is already deployed to an application/service. The dwState    |
	//	|                                                | parameter is 1, and the specified resource is in maintenance mode (see section   |
	//	|                                                | 3.1.1.1.1.2). The dwState parameter is 1, and the specified resource depends on  |
	//	|                                                | one or more additional resources. The dwState parameter is 0, and the specified  |
	//	|                                                | resource does not currently allow volumes to be shared to all nodes in a cluster |
	//	|                                                | (ResourceSharedVolumes is already FALSE).                                        |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000013D7 ERROR_CLUSTER_RESTYPE_NOT_SUPPORTED | The dwState parameter is 1 and the specified resource is not of the correct      |
	//	|                                                | type. Shared access can only be enabled for resources that are of the Physical   |
	//	|                                                | Disk Resource type.                                                              |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000138C ERROR_RESOURCE_NOT_ONLINE           | The dwState parameter is 1 and the specified resource is not online. The         |
	//	|                                                | resource MUST be online to enable shared access.                                 |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// For any other condition, this method returns a value that is not one of the values
	// listed in the preceding table. The client MUST behave in one consistent, identical
	// manner for all values that are not listed in the preceding table. The client treats
	// errors specified in 3.2.4.6 as recoverable errors and initiate the reconnect procedure
	// as specified in section 3.2.4.6.
	ChangeCSVState(context.Context, *ChangeCSVStateRequest) (*ChangeCSVStateResponse, error)

	// (Protocol Version 3) The ApiCreateNodeEnumEx method SHOULD<103> return two ENUM_LIST
	// structures of equal length containing the ID and Name attributes of the requested
	// objects of the specified enumeration type from the cluster state. Each Element in
	// the ReturnIdEnum parameter corresponds to the ID of the element at the same offset
	// in the ReturnNameEnum parameter.
	//
	// If multiple enumeration types are indicated, the resulting ENUM_LIST contains zero
	// or more entries of each type, and the type of each entry in the lists are indicated
	// by the ENUM_ENTRY data structure, as specified in section 2.2.3.4.
	//
	// The server SHOULD accept an ApiCreateNodeEnumEx request if its protocol server state
	// is read-only, and the server MUST accept the request for processing if it is in the
	// read/write state, as specified in section 3.1.1.
	//
	// Return Values: The method MUST return the following error codes for the specified
	// conditions.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS              | Success.                                                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000057 ERROR_INVALID_PARAMETER   | The dwType parameter is not one of the specified values, or the dwOptions        |
	//	|                                       | parameter is not 0x00000000.                                                     |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000006D1 RPC_S_PROCNUM_OUT_OF_RANGE | The server does not support this method.                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//
	// For any other condition, this method returns a value that is not one of the values
	// listed in the preceding table. The client MUST behave in one consistent, identical
	// manner for all values that are not listed in the preceding table. The client SHOULD
	// treat errors specified in 3.2.4.6 as recoverable errors and initiate the reconnect
	// procedure as specified in section 3.2.4.6.
	CreateNodeEnumEx(context.Context, *CreateNodeEnumExRequest) (*CreateNodeEnumExResponse, error)

	// (Protocol Version 3) The ApiCreateEnumEx method SHOULD<104> return two ENUM_LIST
	// structures of equal length containing the ID and Name attributes of the requested
	// objects of the specified enumeration type from the cluster state. Each Element in
	// the ReturnIdEnum parameter corresponds to the ID of the element at the same offset
	// in the ReturnNameEnum parameter.
	//
	// If multiple enumeration types are indicated, the resulting ENUM_LIST contains zero
	// or more entries of each type, and the type of each entry in the list is indicated
	// by the ENUM_ENTRY data structure, as specified in section 2.2.3.4.
	//
	// The server SHOULD accept an ApiCreateEnumEx request if its protocol server state
	// is read-only, as specified in section 3.1.1, and the dwType parameter is CLUSTER_ENUM_NODE.
	// The server MUST accept an ApiCreateEnumEx request if its protocol server state is
	// read/write.
	//
	// Return Values: The method MUST return the following error codes for the conditions
	// that are specified as follows.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS              | Success.                                                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY    | The server failed to allocate enough memory for the ReturnEnum parameter.        |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000057 ERROR_INVALID_PARAMETER   | The enumeration type that is specified by dwType is not valid or dwOptions is    |
	//	|                                       | not set to 0x00000000.                                                           |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000006D1 RPC_S_PROCNUM_OUT_OF_RANGE | The server does not support this method.                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//
	// For any other condition, the server returns a value that is not one of the values
	// listed in the preceding table. The client MUST behave in one consistent, identical
	// manner for all values that are not listed in the preceding table. However, the client
	// SHOULD treat errors specified in 3.2.4.6 as recoverable errors and initiate the reconnect
	// procedure as specified in section 3.2.4.6.
	CreateEnumEx(context.Context, *CreateEnumExRequest) (*CreateEnumExResponse, error)

	// (Protocol Version 3) The ApiPauseNodeEx method SHOULD<105> instruct the server to
	// suspend group ownership and failover activity on the designated node and, optionally,
	// to move groups on the designated node to different nodes in the cluster.
	//
	// Return Values: The method MUST return one of the following error codes.
	//
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                        RETURN                        |                                                                                  |
	//	|                      VALUE/CODE                      |                                   DESCRIPTION                                    |
	//	|                                                      |                                                                                  |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                             | The method completed successfully.                                               |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000006 ERROR_INVALID_HANDLE                      | The data that is designated by the hNode parameter does not represent a valid    |
	//	|                                                      | HNODE_RPC context handle.                                                        |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000013BA ERROR_CLUSTER_NODE_DOWN                   | There are no nodes in the cluster that are in the ClusterNodeUp state other than |
	//	|                                                      | the node designated by the hNode parameter.                                      |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000003E5 ERROR_IO_PENDING                          | The server is in the process of evacuating the specified node.                   |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000174A ERROR_CLUSTER_NODE_EVACUATION_IN_PROGRESS | The server is already evacuating the specified node due to a prior call to this  |
	//	|                                                      | method.                                                                          |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// For any other condition, the server returns a value that is not one of the values
	// listed in the preceding table. The client MUST behave identically for all return
	// values that are not listed in the preceding table; however, the client SHOULD treat
	// errors specified in section 3.2.4.6 as recoverable errors and initiate the reconnect
	// procedure as specified in that section.
	PauseNodeEx(context.Context, *PauseNodeExRequest) (*PauseNodeExResponse, error)

	// (Protocol Version 3) The ApiPauseNodeWithDrainTarget method SHOULD<106> instruct
	// the server to suspend group ownership and failover activity on the designated target
	// node and to move all groups from the designated node to a designated node in the
	// cluster.
	//
	// This server MUST handle this method in the same manner as ApiPauseNodeEx (section
	// 3.1.4.2.125) except that the server MUST attempt to move groups hosted by the node
	// specified by the hNode parameter to the node specified by the hNodeDrainTarget parameter.
	// The server SHOULD move the groups according to preferences, limitations, and other
	// policies as if ApiMoveGroupToNode (section 3.1.4.2.53) or ApiMoveGroupToNodeEx (section
	// 3.1.4.2.132) had been called for each of these groups individually.
	//
	// Return Values: This method MUST return one of the error codes returned by ApiPauseNodeEx
	// (section 3.1.4.2.125) or one of the following values:
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                                  |
	//	|                VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000138D ERROR_HOST_NODE_NOT_AVAILABLE | The node designated by the hNodeDrainTarget parameter is in an invalid state.    |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000072 ERROR_INVALID_TARGET_HANDLE   | The node designated by the hNodeDrainTarget parameter is an invalid destination  |
	//	|                                          | node. This method MUST return ERROR_INVALID_TARGET_HANDLE if the node designated |
	//	|                                          | by the hNodeDrainTarget parameter is the same as the node designated by the      |
	//	|                                          | hNode parameter.                                                                 |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//
	// For any other condition, the server returns a value that is not one of the values
	// listed in the preceding table. The client MUST behave identically for all return
	// values that are not listed in the preceding table; however, the client SHOULD treat
	// errors specified in section 3.2.4.6 as recoverable errors and initiate the reconnect
	// procedure as specified in that section.
	PauseNodeWithDrainTarget(context.Context, *PauseNodeWithDrainTargetRequest) (*PauseNodeWithDrainTargetResponse, error)

	// ApiResumeNodeEx operation.
	ResumeNodeEx(context.Context, *ResumeNodeExRequest) (*ResumeNodeExResponse, error)

	// (Protocol Version 3) The ApiCreateGroupEx method SHOULD<108> extend functionality
	// of the ApiCreateGroup (section 3.1.4.2.43) method, allowing the client to provide
	// additional information about the group being created.
	//
	// Return Values: This method MUST return a valid HGROUP_RPC context handle, as specified
	// in section 2.2.1.3, to indicate success. Otherwise, it MUST return NULL.
	CreateGroupEx(context.Context, *CreateGroupExRequest) (*CreateGroupExResponse, error)

	// (Protocol Version 3) The ApiOnlineGroupEx method SHOULD<109> instruct the server
	// to make all the resources in the designated group active or available on the node
	// that is hosting the group. The persistent state of the group is set to Online and
	// is updated in the nonvolatile cluster state.
	//
	// Return Values: This method MUST return the same error codes as specified for ApiOnlineGroup
	// (section 3.1.4.2.50), in addition to the following return value.
	//
	//	+------------------------------------+-----------------------------------------------------------------+
	//	|               RETURN               |                                                                 |
	//	|             VALUE/CODE             |                           DESCRIPTION                           |
	//	|                                    |                                                                 |
	//	+------------------------------------+-----------------------------------------------------------------+
	//	+------------------------------------+-----------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The dwOnlineFlags parameter is not one of the specified values. |
	//	+------------------------------------+-----------------------------------------------------------------+
	OnlineGroupEx(context.Context, *OnlineGroupExRequest) (*OnlineGroupExResponse, error)

	// (Protocol Version 3) The ApiOfflineGroupEx method SHOULD<113> instruct the server
	// to make all the resources in the designated group inactive or unavailable on the
	// node that is hosting the group.
	//
	// Return Values: This method MUST return the same error codes as specified for ApiOfflineGroup
	// (section 3.1.4.2.51).
	OfflineGroupEx(context.Context, *OfflineGroupExRequest) (*OfflineGroupExResponse, error)

	// (Protocol Version 3) The ApiMoveGroupEx method SHOULD<114> instruct the server to
	// move ownership of the specified group to another node in the cluster.
	//
	// Return Values: This method MUST return the same error codes as specified for ApiMoveGroup
	// (section 3.1.4.2.52), in addition to the following return value.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The client included both the CLUSAPI_GROUP_MOVE_IGNORE_RESOURCE_STATUS and       |
	//	|                                    | CLUSAPI_GROUP_MOVE_QUEUE_ENABLED flags in the dwMoveFlags parameter.             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	MoveGroupEx(context.Context, *MoveGroupExRequest) (*MoveGroupExResponse, error)

	// (Protocol Version 3) The ApiMoveGroupToNodeEx method SHOULD<116> instructs the server
	// to move ownership of a group to the specified node in the cluster.
	//
	// Return Values: This method MUST return the same error codes as specified for ApiMoveGroupToNode
	// (section 3.1.4.2.53) and ApiMoveGroupEx (section 3.1.4.2.131).
	MoveGroupToNodeEx(context.Context, *MoveGroupToNodeExRequest) (*MoveGroupToNodeExResponse, error)

	// (Protocol Version 3) The ApiCancelClusterGroupOperation SHOULD<117>allow a client
	// to cancel a pending group move operation.
	//
	// Return Values: This method MUST return one of the following values.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | The method completed successfully.                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The dwCancelFlags parameter is not set to 0.                                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000006 ERROR_INVALID_HANDLE    | The hGroup parameter does not represent a valid HGROUP_RPC context handle.       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000003E5 ERROR_IO_PENDING        | The server has accepted the request and will process it asynchronously.          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000139F ERROR_INVALID_STATE     | The specified group is not moving or the group move operation is no longer       |
	//	|                                    | cancellable.                                                                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// For any other condition, the server returns a value that is not one of the values
	// listed in the preceding table. The client MUST treat all values that are not listed
	// in the preceding table identically. However, the client SHOULD treat errors specified
	// in section 3.2.4.6 as recoverable errors and initiate the reconnect procedure as
	// specified in section 3.2.4.6.
	CancelClusterGroupOperation(context.Context, *CancelClusterGroupOperationRequest) (*CancelClusterGroupOperationResponse, error)

	// (Protocol Version 3) The ApiOnlineResourceEx method SHOULD<118> instruct the server
	// to make the specified resource active or available on the node that currently owns
	// it.
	//
	// Return Values: This method MUST return the same error codes as returned by the ApiOnlineResource
	// (section 3.1.4.2.18) method, in addition to the following return value.
	//
	//	+------------------------------------+-----------------------------------------------------------------+
	//	|               RETURN               |                                                                 |
	//	|             VALUE/CODE             |                           DESCRIPTION                           |
	//	|                                    |                                                                 |
	//	+------------------------------------+-----------------------------------------------------------------+
	//	+------------------------------------+-----------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The dwOnlineFlags parameter is not one of the specified values. |
	//	+------------------------------------+-----------------------------------------------------------------+
	OnlineResourceEx(context.Context, *OnlineResourceExRequest) (*OnlineResourceExResponse, error)

	// (Protocol Version 3) The ApiOfflineResourceEx SHOULD<121> instruct the server to
	// make the designated resource inactive or unavailable on the node that currently owns
	// it.
	//
	// Return Values: This method MUST return the same error codes returned by the ApiOfflineResource
	// (section 3.1.4.2.19) method, in addition to the following return value.
	//
	//	+------------------------------------+------------------------------------------------------------------+
	//	|               RETURN               |                                                                  |
	//	|             VALUE/CODE             |                           DESCRIPTION                            |
	//	|                                    |                                                                  |
	//	+------------------------------------+------------------------------------------------------------------+
	//	+------------------------------------+------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The dwOfflineFlags parameter is not one of the specified values. |
	//	+------------------------------------+------------------------------------------------------------------+
	OfflineResourceEx(context.Context, *OfflineResourceExRequest) (*OfflineResourceExResponse, error)

	// (Protocol Version 3) The ApiCreateNotifyV2 method uses the current RPC connection
	// to establish a context on the server about the interaction of a client with a cluster
	// version 2 notification port.
	//
	// Return Values: The method MUST return a valid HNOTIFY_RPC context handle, as specified
	// in section 2.2.1.6, to indicate success. Otherwise, it MUST return NULL.
	CreateNotifyV2(context.Context, *CreateNotifyV2Request) (*CreateNotifyV2Response, error)

	// (Protocol Version 3) The ApiAddNotifyV2 method SHOULD<123> register for notifications
	// for the object and type specified on an HNOTIFY_RPC context handle previously returned
	// by a call to ApiCreateNotifyV2 (section 3.1.4.2.136). Clients can use this method
	// to register for multiple notifications for a given object in a single call. However,
	// clients MUST call this method for every object for which it needs to receive notifications.
	//
	// Clients MUST only call this method with an HNOTIFY_RPC handle returned by the ApiCreateNotifyV2
	// (section 3.1.4.2.136) method. The server MUST use the root key of the cluster registry
	// as the key for which state and configuration changes are monitored if the object
	// type specified in the filter parameter is CLUSTER_OBJECT_TYPE_REGISTRY.
	//
	// If the filter includes the CLUSTER_CHANGE_REGISTRY_SUBTREE_V2 value, the server MUST
	// extend the scope of monitoring to include all subkeys, all subkey security descriptors,
	// and all values under the root key. The server MUST also report changes under the
	// specified key and all subkeys.
	//
	// The server SHOULD accept an ApiAddNotifyV2 request if its protocol server state is
	// read-only and MUST accept the request if its state is read/write, as specified in
	// section 3.1.1.
	//
	// The server SHOULD accept an ApiAddNotifyV2 request if the client's access level is
	// at least "Read" (section 3.1.4).
	//
	// Return Values: This method MUST return one of the following error codes.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | The method completed successfully.                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000006 ERROR_INVALID_HANDLE    | Either the hNotify parameter does not represent a valid HNOTIFY_RPC context      |
	//	|                                    | handle, or the hObject parameter does not represent a valid HGENERIC_RPC context |
	//	|                                    | handle, or both.                                                                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | Any of the following conditions: The filter parameter contains invalid flags.    |
	//	|                                    | The filter parameter specifies an invalid object type. The dwVersion parameter   |
	//	|                                    | contains an invalid value.                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// For any other condition, the server returns a value that is not one of the values
	// listed in the above table. The client MUST treat all values that are not listed in
	// the preceding table identically, except for recoverable errors specified in section
	// 3.2.4.6, for which the client SHOULD initiate the reconnect procedure.
	AddNotifyV2(context.Context, *AddNotifyV2Request) (*AddNotifyV2Response, error)

	// (Protocol Version 3) The ApiGetNotifyV2 method is an asynchronous RPC method that
	// SHOULD<124> instruct the server to return a set of queued events from a designated
	// version 2 notification port corresponding to a particular context handle. The server
	// MUST NOT complete this method until an indication has been queued to the port or
	// the port has been closed through a separate call to the ApiUnblockGetNotifyCall (section
	// 3.1.4.2.107) or ApiCloseNotify (section 3.1.4.2.57) method.
	//
	// Return Values: This method MUST return the following error codes for the following
	// conditions.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN               |                                                                                  |
	//	|            VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS          | The method completed successfully.                                               |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000006 ERROR_INVALID_HANDLE   | The data that is pointed to by the hNotify parameter does not represent a valid  |
	//	|                                   | HNOTIFY_RPC context handle.                                                      |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS    | The notification port represented by the hNotify parameter has been closed.      |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 ERROR_INVALID_FUNCTION | Either the ApiUnblockGetNotifyCall (section 3.1.4.2.107) method or the           |
	//	|                                   | ApiCloseNotify (section 3.1.4.2.57) method has been called in another thread.    |
	//	|                                   | The client SHOULD terminate the notification thread.                             |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//
	// For any other condition, the server returns a value that is not one of the values
	// listed in the preceding table. The client MUST treat all such values identically,
	// with the exception of errors specified in section 3.2.4.6, which are recoverable
	// and for which the client SHOULD initiate the reconnect procedure.
	GetNotifyV2(context.Context, *GetNotifyV2Request) (*GetNotifyV2Response, error)

	// Opnum140NotUsedOnWire operation.
	// Opnum140NotUsedOnWire

	// Opnum141NotUsedOnWire operation.
	// Opnum141NotUsedOnWire

	// Opnum142NotUsedOnWire operation.
	// Opnum142NotUsedOnWire

	// (Protocol Version 3) The ApiCreateGroupEnum method SHOULD<126> return an enumeration
	// of groups from the current set of groups in the cluster.
	//
	// Return Values: This method MUST return the following values for the following conditions.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | The method completed successfully.                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The client specified a property name in the pProperties or pRoProperties         |
	//	|                                    | parameters that is not the name of a group common or private property,           |
	//	|                                    | respectively.                                                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// For any other condition, the server returns a value that is not one of the values
	// listed in the preceding table. Clients MUST treat all such values identically, with
	// the exception of errors specified in section 3.2.4.6, which are recoverable errors
	// for which the client SHOULD initiate the reconnect procedure.
	CreateGroupEnum(context.Context, *CreateGroupEnumRequest) (*CreateGroupEnumResponse, error)

	// (Protocol Version 3) The ApiCreateResourceEnum method SHOULD<127> return an enumeration
	// of resources from the current set of resources in the cluster.
	//
	// Return Values: This method MUST return the following values for the following conditions.
	//
	//	+--------------------------+------------------------------------+
	//	|          RETURN          |                                    |
	//	|        VALUE/CODE        |            DESCRIPTION             |
	//	|                          |                                    |
	//	+--------------------------+------------------------------------+
	//	+--------------------------+------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The method completed successfully. |
	//	+--------------------------+------------------------------------+
	//
	// For any other condition, the server returns a value that is not one of the values
	// listed in the preceding table. Clients MUST treat all such values identically, with
	// the exception of errors specified in section 3.2.4.6, which are recoverable errors
	// for which the client SHOULD initiate the reconnect procedure.
	CreateResourceEnum(context.Context, *CreateResourceEnumRequest) (*CreateResourceEnumResponse, error)

	// ApiExecuteReadBatch operation.
	ExecuteReadBatch(context.Context, *ExecuteReadBatchRequest) (*ExecuteReadBatchResponse, error)

	// The ApiRestartResource method SHOULD<129> instruct the server to restart a resource
	// without affecting the resource's persistent state. The server SHOULD bring the resource
	// offline, then bring the resource back to its persistent state.
	//
	// Return Values: This method MUST return ERROR_SUCCESS (0x00000000) on success, or
	// a different value for any other condition. Clients MUST treat all values other than
	// ERROR_SUCCESS identically, except for recoverable errors as specified in section
	// 3.2.4.6, for which the client MUST initiate the reconnect procedure.
	RestartResource(context.Context, *RestartResourceRequest) (*RestartResourceResponse, error)

	// (Protocol Version 3) The ApiGetNotifyAsync method is an asynchronous RPC method that
	// SHOULD<130> be used instead of ApiGetNotify to instruct the server to return the
	// next set of queued events corresponding to a particular context handle. The server
	// MUST NOT complete this method until an indication has been queued to the port or
	// the port has been closed through a separate call to the ApiUnblockGetNotifyCall or
	// ApiCloseNotify method.
	//
	// Return Values: This method MUST return one of the following values.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN               |                                                                                  |
	//	|            VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS          | The method completed successfully.                                               |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000006 ERROR_INVALID_HANDLE   | The data that is pointed to by the hNotify parameter does not represent a valid  |
	//	|                                   | HNOTIFY_RPC context handle.                                                      |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS    | The notification port represented by the hNotify parameter has been closed.      |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000001 ERROR_INVALID_FUNCTION | Either the ApiUnblockGetNotifyCall (section 3.1.4.2.107) method or the           |
	//	|                                   | ApiCloseNotify (section 3.1.4.2.57) method has been called in another thread.    |
	//	|                                   | The client SHOULD terminate the notification thread.                             |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//
	// For any other condition, the server returns a value that is not one of the values
	// listed in the preceding table. The client MUST treat all such values identically,
	// with the exception of errors specified in section 3.2.4.6, which are recoverable
	// and for which the client SHOULD initiate the reconnect procedure.
	GetNotifyAsync(context.Context, *GetNotifyAsyncRequest) (*GetNotifyAsyncResponse, error)

	// Opnum148NotUsedOnWire operation.
	// Opnum148NotUsedOnWire

	// Opnum149otUsedOnWire operation.
	Opnum149otUsedOnWire(context.Context, *Opnum149otUsedOnWireRequest) (*Opnum149otUsedOnWireResponse, error)

	// Opnum150NotUsedOnWire operation.
	// Opnum150NotUsedOnWire

	// Opnum151NotUsedOnWire operation.
	// Opnum151NotUsedOnWire

	// Opnum152NotUsedOnWire operation.
	// Opnum152NotUsedOnWire

	// Opnum153NotUsedOnWire operation.
	// Opnum153NotUsedOnWire

	// Opnum154NotUsedOnWire operation.
	// Opnum154NotUsedOnWire

	// ApiAddNotifyResourceTypeV2 operation.
	AddNotifyResourceTypeV2(context.Context, *AddNotifyResourceTypeV2Request) (*AddNotifyResourceTypeV2Response, error)

	// Opnum156NotUsedOnWire operation.
	// Opnum156NotUsedOnWire

	// ApiExecuteReadBatchEx operation.
	ExecuteReadBatchEx(context.Context, *ExecuteReadBatchExRequest) (*ExecuteReadBatchExResponse, error)

	// Opnum158NotUsedOnWire operation.
	// Opnum158NotUsedOnWire

	// Opnum159NotUsedOnWire operation.
	// Opnum159NotUsedOnWire

	// Opnum160NotUsedOnWire operation.
	// Opnum160NotUsedOnWire

	// Opnum161NotUsedOnWire operation.
	// Opnum161NotUsedOnWire

	// Opnum162NotUsedOnWire operation.
	// Opnum162NotUsedOnWire

	// ApiCreateGroupSet operation.
	CreateGroupSet(context.Context, *CreateGroupSetRequest) (*CreateGroupSetResponse, error)

	// ApiOpenGroupSet operation.
	OpenGroupSet(context.Context, *OpenGroupSetRequest) (*OpenGroupSetResponse, error)

	// ApiCloseGroupSet operation.
	CloseGroupSet(context.Context, *CloseGroupSetRequest) (*CloseGroupSetResponse, error)

	// ApiDeleteGroupSet operation.
	DeleteGroupSet(context.Context, *DeleteGroupSetRequest) (*DeleteGroupSetResponse, error)

	// ApiAddGroupToGroupSet operation.
	AddGroupToGroupSet(context.Context, *AddGroupToGroupSetRequest) (*AddGroupToGroupSetResponse, error)

	// ApiRemoveGroupFromGroupSet operation.
	RemoveGroupFromGroupSet(context.Context, *RemoveGroupFromGroupSetRequest) (*RemoveGroupFromGroupSetResponse, error)

	// ApiMoveGroupToGroupSet operation.
	MoveGroupToGroupSet(context.Context, *MoveGroupToGroupSetRequest) (*MoveGroupToGroupSetResponse, error)

	// Opnum170NotUsedOnWire operation.
	// Opnum170NotUsedOnWire

	// ApiAddGroupSetDependency operation.
	AddGroupSetDependency(context.Context, *AddGroupSetDependencyRequest) (*AddGroupSetDependencyResponse, error)

	// ApiAddGroupToGroupSetDependency operation.
	AddGroupToGroupSetDependency(context.Context, *AddGroupToGroupSetDependencyRequest) (*AddGroupToGroupSetDependencyResponse, error)

	// ApiNodeGroupSetControl operation.
	NodeGroupSetControl(context.Context, *NodeGroupSetControlRequest) (*NodeGroupSetControlResponse, error)

	// ApiGroupSetControl operation.
	GroupSetControl(context.Context, *GroupSetControlRequest) (*GroupSetControlResponse, error)

	// ApiSetGroupDependencyExpression operation.
	SetGroupDependencyExpression(context.Context, *SetGroupDependencyExpressionRequest) (*SetGroupDependencyExpressionResponse, error)

	// ApiRemoveClusterGroupDependency operation.
	RemoveClusterGroupDependency(context.Context, *RemoveClusterGroupDependencyRequest) (*RemoveClusterGroupDependencyResponse, error)

	// ApiSetGroupSetDependencyExpression operation.
	SetGroupSetDependencyExpression(context.Context, *SetGroupSetDependencyExpressionRequest) (*SetGroupSetDependencyExpressionResponse, error)

	// ApiRemoveGroupSetDependency operation.
	RemoveGroupSetDependency(context.Context, *RemoveGroupSetDependencyRequest) (*RemoveGroupSetDependencyResponse, error)

	// ApiRemoveClusterGroupToGroupSetDependency operation.
	RemoveClusterGroupToGroupSetDependency(context.Context, *RemoveClusterGroupToGroupSetDependencyRequest) (*RemoveClusterGroupToGroupSetDependencyResponse, error)

	// ApiCreateGroupSetEnum operation.
	CreateGroupSetEnum(context.Context, *CreateGroupSetEnumRequest) (*CreateGroupSetEnumResponse, error)

	// ApiCreateNetInterfaceEnum operation.
	CreateNetInterfaceEnum(context.Context, *CreateNetInterfaceEnumRequest) (*CreateNetInterfaceEnumResponse, error)

	// ApiChangeCsvStateEx operation.
	ChangeCSVStateEx(context.Context, *ChangeCSVStateExRequest) (*ChangeCSVStateExResponse, error)

	// ApiAddGroupToGroupSetEx operation.
	AddGroupToGroupSetEx(context.Context, *AddGroupToGroupSetExRequest) (*AddGroupToGroupSetExResponse, error)

	// ApiChangeResourceGroupEx operation.
	ChangeResourceGroupEx(context.Context, *ChangeResourceGroupExRequest) (*ChangeResourceGroupExResponse, error)
}

func RegisterClusapi3Server(conn dcerpc.Conn, o Clusapi3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewClusapi3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Clusapi3SyntaxV3_0))...)
}

func NewClusapi3ServerHandle(o Clusapi3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Clusapi3ServerHandle(ctx, o, opNum, r)
	}
}

func Clusapi3ServerHandle(ctx context.Context, o Clusapi3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	case 109: // ApiSetResourceDependencyExpression
		op := &xxx_SetResourceDependencyExpressionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetResourceDependencyExpressionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetResourceDependencyExpression(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 110: // ApiGetResourceDependencyExpression
		op := &xxx_GetResourceDependencyExpressionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetResourceDependencyExpressionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetResourceDependencyExpression(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 111: // Opnum111NotUsedOnWire
		// Opnum111NotUsedOnWire
		return nil, nil
	case 112: // ApiGetResourceNetworkName
		op := &xxx_GetResourceNetworkNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetResourceNetworkNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetResourceNetworkName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 113: // ApiExecuteBatch
		op := &xxx_ExecuteBatchOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExecuteBatchRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExecuteBatch(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 114: // ApiCreateBatchPort
		op := &xxx_CreateBatchPortOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateBatchPortRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateBatchPort(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 115: // ApiGetBatchNotification
		op := &xxx_GetBatchNotificationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetBatchNotificationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetBatchNotification(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 116: // ApiCloseBatchPort
		op := &xxx_CloseBatchPortOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseBatchPortRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseBatchPort(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 117: // ApiOpenClusterEx
		op := &xxx_OpenClusterExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenClusterExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenClusterEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 118: // ApiOpenNodeEx
		op := &xxx_OpenNodeExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenNodeExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenNodeEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 119: // ApiOpenGroupEx
		op := &xxx_OpenGroupExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenGroupExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenGroupEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 120: // ApiOpenResourceEx
		op := &xxx_OpenResourceExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenResourceExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenResourceEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 121: // ApiOpenNetworkEx
		op := &xxx_OpenNetworkExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenNetworkExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenNetworkEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 122: // ApiOpenNetInterfaceEx
		op := &xxx_OpenNetInterfaceExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenNetInterfaceExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenNetInterfaceEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 123: // ApiChangeCsvState
		op := &xxx_ChangeCSVStateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ChangeCSVStateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ChangeCSVState(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 124: // ApiCreateNodeEnumEx
		op := &xxx_CreateNodeEnumExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateNodeEnumExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateNodeEnumEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 125: // ApiCreateEnumEx
		op := &xxx_CreateEnumExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateEnumExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateEnumEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 126: // ApiPauseNodeEx
		op := &xxx_PauseNodeExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PauseNodeExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PauseNodeEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 127: // ApiPauseNodeWithDrainTarget
		op := &xxx_PauseNodeWithDrainTargetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PauseNodeWithDrainTargetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PauseNodeWithDrainTarget(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 128: // ApiResumeNodeEx
		op := &xxx_ResumeNodeExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResumeNodeExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ResumeNodeEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 129: // ApiCreateGroupEx
		op := &xxx_CreateGroupExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateGroupExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateGroupEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 130: // ApiOnlineGroupEx
		op := &xxx_OnlineGroupExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OnlineGroupExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OnlineGroupEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 131: // ApiOfflineGroupEx
		op := &xxx_OfflineGroupExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OfflineGroupExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OfflineGroupEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 132: // ApiMoveGroupEx
		op := &xxx_MoveGroupExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MoveGroupExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MoveGroupEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 133: // ApiMoveGroupToNodeEx
		op := &xxx_MoveGroupToNodeExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MoveGroupToNodeExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MoveGroupToNodeEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 134: // ApiCancelClusterGroupOperation
		op := &xxx_CancelClusterGroupOperationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CancelClusterGroupOperationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CancelClusterGroupOperation(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 135: // ApiOnlineResourceEx
		op := &xxx_OnlineResourceExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OnlineResourceExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OnlineResourceEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 136: // ApiOfflineResourceEx
		op := &xxx_OfflineResourceExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OfflineResourceExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OfflineResourceEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 137: // ApiCreateNotifyV2
		op := &xxx_CreateNotifyV2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateNotifyV2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateNotifyV2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 138: // ApiAddNotifyV2
		op := &xxx_AddNotifyV2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddNotifyV2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddNotifyV2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 139: // ApiGetNotifyV2
		op := &xxx_GetNotifyV2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNotifyV2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNotifyV2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 140: // Opnum140NotUsedOnWire
		// Opnum140NotUsedOnWire
		return nil, nil
	case 141: // Opnum141NotUsedOnWire
		// Opnum141NotUsedOnWire
		return nil, nil
	case 142: // Opnum142NotUsedOnWire
		// Opnum142NotUsedOnWire
		return nil, nil
	case 143: // ApiCreateGroupEnum
		op := &xxx_CreateGroupEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateGroupEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateGroupEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 144: // ApiCreateResourceEnum
		op := &xxx_CreateResourceEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateResourceEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateResourceEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 145: // ApiExecuteReadBatch
		op := &xxx_ExecuteReadBatchOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExecuteReadBatchRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExecuteReadBatch(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 146: // ApiRestartResource
		op := &xxx_RestartResourceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RestartResourceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RestartResource(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 147: // ApiGetNotifyAsync
		op := &xxx_GetNotifyAsyncOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNotifyAsyncRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNotifyAsync(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 148: // Opnum148NotUsedOnWire
		// Opnum148NotUsedOnWire
		return nil, nil
	case 149: // Opnum149otUsedOnWire
		op := &xxx_Opnum149otUsedOnWireOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Opnum149otUsedOnWireRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Opnum149otUsedOnWire(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 150: // Opnum150NotUsedOnWire
		// Opnum150NotUsedOnWire
		return nil, nil
	case 151: // Opnum151NotUsedOnWire
		// Opnum151NotUsedOnWire
		return nil, nil
	case 152: // Opnum152NotUsedOnWire
		// Opnum152NotUsedOnWire
		return nil, nil
	case 153: // Opnum153NotUsedOnWire
		// Opnum153NotUsedOnWire
		return nil, nil
	case 154: // Opnum154NotUsedOnWire
		// Opnum154NotUsedOnWire
		return nil, nil
	case 155: // ApiAddNotifyResourceTypeV2
		op := &xxx_AddNotifyResourceTypeV2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddNotifyResourceTypeV2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddNotifyResourceTypeV2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 156: // Opnum156NotUsedOnWire
		// Opnum156NotUsedOnWire
		return nil, nil
	case 157: // ApiExecuteReadBatchEx
		op := &xxx_ExecuteReadBatchExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExecuteReadBatchExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExecuteReadBatchEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 158: // Opnum158NotUsedOnWire
		// Opnum158NotUsedOnWire
		return nil, nil
	case 159: // Opnum159NotUsedOnWire
		// Opnum159NotUsedOnWire
		return nil, nil
	case 160: // Opnum160NotUsedOnWire
		// Opnum160NotUsedOnWire
		return nil, nil
	case 161: // Opnum161NotUsedOnWire
		// Opnum161NotUsedOnWire
		return nil, nil
	case 162: // Opnum162NotUsedOnWire
		// Opnum162NotUsedOnWire
		return nil, nil
	case 163: // ApiCreateGroupSet
		op := &xxx_CreateGroupSetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateGroupSetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateGroupSet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 164: // ApiOpenGroupSet
		op := &xxx_OpenGroupSetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenGroupSetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenGroupSet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 165: // ApiCloseGroupSet
		op := &xxx_CloseGroupSetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CloseGroupSetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CloseGroupSet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 166: // ApiDeleteGroupSet
		op := &xxx_DeleteGroupSetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteGroupSetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteGroupSet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 167: // ApiAddGroupToGroupSet
		op := &xxx_AddGroupToGroupSetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddGroupToGroupSetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddGroupToGroupSet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 168: // ApiRemoveGroupFromGroupSet
		op := &xxx_RemoveGroupFromGroupSetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveGroupFromGroupSetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveGroupFromGroupSet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 169: // ApiMoveGroupToGroupSet
		op := &xxx_MoveGroupToGroupSetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MoveGroupToGroupSetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.MoveGroupToGroupSet(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 170: // Opnum170NotUsedOnWire
		// Opnum170NotUsedOnWire
		return nil, nil
	case 171: // ApiAddGroupSetDependency
		op := &xxx_AddGroupSetDependencyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddGroupSetDependencyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddGroupSetDependency(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 172: // ApiAddGroupToGroupSetDependency
		op := &xxx_AddGroupToGroupSetDependencyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddGroupToGroupSetDependencyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddGroupToGroupSetDependency(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 173: // ApiNodeGroupSetControl
		op := &xxx_NodeGroupSetControlOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &NodeGroupSetControlRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.NodeGroupSetControl(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 174: // ApiGroupSetControl
		op := &xxx_GroupSetControlOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GroupSetControlRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GroupSetControl(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 175: // ApiSetGroupDependencyExpression
		op := &xxx_SetGroupDependencyExpressionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetGroupDependencyExpressionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetGroupDependencyExpression(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 176: // ApiRemoveClusterGroupDependency
		op := &xxx_RemoveClusterGroupDependencyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveClusterGroupDependencyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveClusterGroupDependency(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 177: // ApiSetGroupSetDependencyExpression
		op := &xxx_SetGroupSetDependencyExpressionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetGroupSetDependencyExpressionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetGroupSetDependencyExpression(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 178: // ApiRemoveGroupSetDependency
		op := &xxx_RemoveGroupSetDependencyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveGroupSetDependencyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveGroupSetDependency(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 179: // ApiRemoveClusterGroupToGroupSetDependency
		op := &xxx_RemoveClusterGroupToGroupSetDependencyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveClusterGroupToGroupSetDependencyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveClusterGroupToGroupSetDependency(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 180: // ApiCreateGroupSetEnum
		op := &xxx_CreateGroupSetEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateGroupSetEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateGroupSetEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 181: // ApiCreateNetInterfaceEnum
		op := &xxx_CreateNetInterfaceEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateNetInterfaceEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateNetInterfaceEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 182: // ApiChangeCsvStateEx
		op := &xxx_ChangeCSVStateExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ChangeCSVStateExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ChangeCSVStateEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 183: // ApiAddGroupToGroupSetEx
		op := &xxx_AddGroupToGroupSetExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddGroupToGroupSetExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddGroupToGroupSetEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 184: // ApiChangeResourceGroupEx
		op := &xxx_ChangeResourceGroupExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ChangeResourceGroupExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ChangeResourceGroupEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented clusapi3
type UnimplementedClusapi3Server struct {
}

func (UnimplementedClusapi3Server) OpenCluster(context.Context, *OpenClusterRequest) (*OpenClusterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CloseCluster(context.Context, *CloseClusterRequest) (*CloseClusterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) SetClusterName(context.Context, *SetClusterNameRequest) (*SetClusterNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetClusterName(context.Context, *GetClusterNameRequest) (*GetClusterNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetClusterVersion(context.Context, *GetClusterVersionRequest) (*GetClusterVersionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetQuorumResource(context.Context, *GetQuorumResourceRequest) (*GetQuorumResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) SetQuorumResource(context.Context, *SetQuorumResourceRequest) (*SetQuorumResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateEnum(context.Context, *CreateEnumRequest) (*CreateEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OpenResource(context.Context, *OpenResourceRequest) (*OpenResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateResource(context.Context, *CreateResourceRequest) (*CreateResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) DeleteResource(context.Context, *DeleteResourceRequest) (*DeleteResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CloseResource(context.Context, *CloseResourceRequest) (*CloseResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetResourceState(context.Context, *GetResourceStateRequest) (*GetResourceStateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) SetResourceName(context.Context, *SetResourceNameRequest) (*SetResourceNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetResourceID(context.Context, *GetResourceIDRequest) (*GetResourceIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetResourceType(context.Context, *GetResourceTypeRequest) (*GetResourceTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) FailResource(context.Context, *FailResourceRequest) (*FailResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OnlineResource(context.Context, *OnlineResourceRequest) (*OnlineResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OfflineResource(context.Context, *OfflineResourceRequest) (*OfflineResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) AddResourceDependency(context.Context, *AddResourceDependencyRequest) (*AddResourceDependencyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) RemoveResourceDependency(context.Context, *RemoveResourceDependencyRequest) (*RemoveResourceDependencyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CanResourceBeDependent(context.Context, *CanResourceBeDependentRequest) (*CanResourceBeDependentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateRestrictionEnum(context.Context, *CreateRestrictionEnumRequest) (*CreateRestrictionEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) AddResourceNode(context.Context, *AddResourceNodeRequest) (*AddResourceNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) RemoveResourceNode(context.Context, *RemoveResourceNodeRequest) (*RemoveResourceNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) ChangeResourceGroup(context.Context, *ChangeResourceGroupRequest) (*ChangeResourceGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateResourceType(context.Context, *CreateResourceTypeRequest) (*CreateResourceTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) DeleteResourceType(context.Context, *DeleteResourceTypeRequest) (*DeleteResourceTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetRootKey(context.Context, *GetRootKeyRequest) (*GetRootKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateKey(context.Context, *CreateKeyRequest) (*CreateKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OpenKey(context.Context, *OpenKeyRequest) (*OpenKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) EnumKey(context.Context, *EnumKeyRequest) (*EnumKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) SetValue(context.Context, *SetValueRequest) (*SetValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) DeleteValue(context.Context, *DeleteValueRequest) (*DeleteValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) QueryValue(context.Context, *QueryValueRequest) (*QueryValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) DeleteKey(context.Context, *DeleteKeyRequest) (*DeleteKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) EnumValue(context.Context, *EnumValueRequest) (*EnumValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CloseKey(context.Context, *CloseKeyRequest) (*CloseKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) QueryInfoKey(context.Context, *QueryInfoKeyRequest) (*QueryInfoKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) SetKeySecurity(context.Context, *SetKeySecurityRequest) (*SetKeySecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetKeySecurity(context.Context, *GetKeySecurityRequest) (*GetKeySecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OpenGroup(context.Context, *OpenGroupRequest) (*OpenGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) DeleteGroup(context.Context, *DeleteGroupRequest) (*DeleteGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CloseGroup(context.Context, *CloseGroupRequest) (*CloseGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetGroupState(context.Context, *GetGroupStateRequest) (*GetGroupStateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) SetGroupName(context.Context, *SetGroupNameRequest) (*SetGroupNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetGroupID(context.Context, *GetGroupIDRequest) (*GetGroupIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetNodeID(context.Context, *GetNodeIDRequest) (*GetNodeIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OnlineGroup(context.Context, *OnlineGroupRequest) (*OnlineGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OfflineGroup(context.Context, *OfflineGroupRequest) (*OfflineGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) MoveGroup(context.Context, *MoveGroupRequest) (*MoveGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) MoveGroupToNode(context.Context, *MoveGroupToNodeRequest) (*MoveGroupToNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateGroupResourceEnum(context.Context, *CreateGroupResourceEnumRequest) (*CreateGroupResourceEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) SetGroupNodeList(context.Context, *SetGroupNodeListRequest) (*SetGroupNodeListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateNotify(context.Context, *CreateNotifyRequest) (*CreateNotifyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CloseNotify(context.Context, *CloseNotifyRequest) (*CloseNotifyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) AddNotifyCluster(context.Context, *AddNotifyClusterRequest) (*AddNotifyClusterResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) AddNotifyNode(context.Context, *AddNotifyNodeRequest) (*AddNotifyNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) AddNotifyGroup(context.Context, *AddNotifyGroupRequest) (*AddNotifyGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) AddNotifyResource(context.Context, *AddNotifyResourceRequest) (*AddNotifyResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) AddNotifyKey(context.Context, *AddNotifyKeyRequest) (*AddNotifyKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) ReAddNotifyNode(context.Context, *ReAddNotifyNodeRequest) (*ReAddNotifyNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) ReAddNotifyGroup(context.Context, *ReAddNotifyGroupRequest) (*ReAddNotifyGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) ReAddNotifyResource(context.Context, *ReAddNotifyResourceRequest) (*ReAddNotifyResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetNotify(context.Context, *GetNotifyRequest) (*GetNotifyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OpenNode(context.Context, *OpenNodeRequest) (*OpenNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CloseNode(context.Context, *CloseNodeRequest) (*CloseNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetNodeState(context.Context, *GetNodeStateRequest) (*GetNodeStateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) PauseNode(context.Context, *PauseNodeRequest) (*PauseNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) ResumeNode(context.Context, *ResumeNodeRequest) (*ResumeNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) EvictNode(context.Context, *EvictNodeRequest) (*EvictNodeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) NodeResourceControl(context.Context, *NodeResourceControlRequest) (*NodeResourceControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) ResourceControl(context.Context, *ResourceControlRequest) (*ResourceControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) NodeResourceTypeControl(context.Context, *NodeResourceTypeControlRequest) (*NodeResourceTypeControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) ResourceTypeControl(context.Context, *ResourceTypeControlRequest) (*ResourceTypeControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) NodeGroupControl(context.Context, *NodeGroupControlRequest) (*NodeGroupControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GroupControl(context.Context, *GroupControlRequest) (*GroupControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) NodeNodeControl(context.Context, *NodeNodeControlRequest) (*NodeNodeControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) NodeControl(context.Context, *NodeControlRequest) (*NodeControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OpenNetwork(context.Context, *OpenNetworkRequest) (*OpenNetworkResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CloseNetwork(context.Context, *CloseNetworkRequest) (*CloseNetworkResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetNetworkState(context.Context, *GetNetworkStateRequest) (*GetNetworkStateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) SetNetworkName(context.Context, *SetNetworkNameRequest) (*SetNetworkNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateNetworkEnum(context.Context, *CreateNetworkEnumRequest) (*CreateNetworkEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetNetworkID(context.Context, *GetNetworkIDRequest) (*GetNetworkIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) SetNetworkPriorityOrder(context.Context, *SetNetworkPriorityOrderRequest) (*SetNetworkPriorityOrderResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) NodeNetworkControl(context.Context, *NodeNetworkControlRequest) (*NodeNetworkControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) NetworkControl(context.Context, *NetworkControlRequest) (*NetworkControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) AddNotifyNetwork(context.Context, *AddNotifyNetworkRequest) (*AddNotifyNetworkResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) ReAddNotifyNetwork(context.Context, *ReAddNotifyNetworkRequest) (*ReAddNotifyNetworkResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OpenNetInterface(context.Context, *OpenNetInterfaceRequest) (*OpenNetInterfaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CloseNetInterface(context.Context, *CloseNetInterfaceRequest) (*CloseNetInterfaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetNetInterfaceState(context.Context, *GetNetInterfaceStateRequest) (*GetNetInterfaceStateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetNetInterface(context.Context, *GetNetInterfaceRequest) (*GetNetInterfaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetNetInterfaceID(context.Context, *GetNetInterfaceIDRequest) (*GetNetInterfaceIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) NodeNetInterfaceControl(context.Context, *NodeNetInterfaceControlRequest) (*NodeNetInterfaceControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) NetInterfaceControl(context.Context, *NetInterfaceControlRequest) (*NetInterfaceControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) AddNotifyNetInterface(context.Context, *AddNotifyNetInterfaceRequest) (*AddNotifyNetInterfaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) ReAddNotifyNetInterface(context.Context, *ReAddNotifyNetInterfaceRequest) (*ReAddNotifyNetInterfaceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateNodeEnum(context.Context, *CreateNodeEnumRequest) (*CreateNodeEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetClusterVersion2(context.Context, *GetClusterVersion2Request) (*GetClusterVersion2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateRestrictionTypeEnum(context.Context, *CreateRestrictionTypeEnumRequest) (*CreateRestrictionTypeEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) BackupClusterDatabase(context.Context, *BackupClusterDatabaseRequest) (*BackupClusterDatabaseResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) NodeClusterControl(context.Context, *NodeClusterControlRequest) (*NodeClusterControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) ClusterControl(context.Context, *ClusterControlRequest) (*ClusterControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) UnblockGetNotifyCall(context.Context, *UnblockGetNotifyCallRequest) (*UnblockGetNotifyCallResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) SetServiceAccountPassword(context.Context, *SetServiceAccountPasswordRequest) (*SetServiceAccountPasswordResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) SetResourceDependencyExpression(context.Context, *SetResourceDependencyExpressionRequest) (*SetResourceDependencyExpressionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetResourceDependencyExpression(context.Context, *GetResourceDependencyExpressionRequest) (*GetResourceDependencyExpressionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetResourceNetworkName(context.Context, *GetResourceNetworkNameRequest) (*GetResourceNetworkNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) ExecuteBatch(context.Context, *ExecuteBatchRequest) (*ExecuteBatchResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateBatchPort(context.Context, *CreateBatchPortRequest) (*CreateBatchPortResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetBatchNotification(context.Context, *GetBatchNotificationRequest) (*GetBatchNotificationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CloseBatchPort(context.Context, *CloseBatchPortRequest) (*CloseBatchPortResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OpenClusterEx(context.Context, *OpenClusterExRequest) (*OpenClusterExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OpenNodeEx(context.Context, *OpenNodeExRequest) (*OpenNodeExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OpenGroupEx(context.Context, *OpenGroupExRequest) (*OpenGroupExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OpenResourceEx(context.Context, *OpenResourceExRequest) (*OpenResourceExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OpenNetworkEx(context.Context, *OpenNetworkExRequest) (*OpenNetworkExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OpenNetInterfaceEx(context.Context, *OpenNetInterfaceExRequest) (*OpenNetInterfaceExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) ChangeCSVState(context.Context, *ChangeCSVStateRequest) (*ChangeCSVStateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateNodeEnumEx(context.Context, *CreateNodeEnumExRequest) (*CreateNodeEnumExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateEnumEx(context.Context, *CreateEnumExRequest) (*CreateEnumExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) PauseNodeEx(context.Context, *PauseNodeExRequest) (*PauseNodeExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) PauseNodeWithDrainTarget(context.Context, *PauseNodeWithDrainTargetRequest) (*PauseNodeWithDrainTargetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) ResumeNodeEx(context.Context, *ResumeNodeExRequest) (*ResumeNodeExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateGroupEx(context.Context, *CreateGroupExRequest) (*CreateGroupExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OnlineGroupEx(context.Context, *OnlineGroupExRequest) (*OnlineGroupExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OfflineGroupEx(context.Context, *OfflineGroupExRequest) (*OfflineGroupExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) MoveGroupEx(context.Context, *MoveGroupExRequest) (*MoveGroupExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) MoveGroupToNodeEx(context.Context, *MoveGroupToNodeExRequest) (*MoveGroupToNodeExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CancelClusterGroupOperation(context.Context, *CancelClusterGroupOperationRequest) (*CancelClusterGroupOperationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OnlineResourceEx(context.Context, *OnlineResourceExRequest) (*OnlineResourceExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OfflineResourceEx(context.Context, *OfflineResourceExRequest) (*OfflineResourceExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateNotifyV2(context.Context, *CreateNotifyV2Request) (*CreateNotifyV2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) AddNotifyV2(context.Context, *AddNotifyV2Request) (*AddNotifyV2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetNotifyV2(context.Context, *GetNotifyV2Request) (*GetNotifyV2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateGroupEnum(context.Context, *CreateGroupEnumRequest) (*CreateGroupEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateResourceEnum(context.Context, *CreateResourceEnumRequest) (*CreateResourceEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) ExecuteReadBatch(context.Context, *ExecuteReadBatchRequest) (*ExecuteReadBatchResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) RestartResource(context.Context, *RestartResourceRequest) (*RestartResourceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GetNotifyAsync(context.Context, *GetNotifyAsyncRequest) (*GetNotifyAsyncResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) Opnum149otUsedOnWire(context.Context, *Opnum149otUsedOnWireRequest) (*Opnum149otUsedOnWireResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) AddNotifyResourceTypeV2(context.Context, *AddNotifyResourceTypeV2Request) (*AddNotifyResourceTypeV2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) ExecuteReadBatchEx(context.Context, *ExecuteReadBatchExRequest) (*ExecuteReadBatchExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateGroupSet(context.Context, *CreateGroupSetRequest) (*CreateGroupSetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) OpenGroupSet(context.Context, *OpenGroupSetRequest) (*OpenGroupSetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CloseGroupSet(context.Context, *CloseGroupSetRequest) (*CloseGroupSetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) DeleteGroupSet(context.Context, *DeleteGroupSetRequest) (*DeleteGroupSetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) AddGroupToGroupSet(context.Context, *AddGroupToGroupSetRequest) (*AddGroupToGroupSetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) RemoveGroupFromGroupSet(context.Context, *RemoveGroupFromGroupSetRequest) (*RemoveGroupFromGroupSetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) MoveGroupToGroupSet(context.Context, *MoveGroupToGroupSetRequest) (*MoveGroupToGroupSetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) AddGroupSetDependency(context.Context, *AddGroupSetDependencyRequest) (*AddGroupSetDependencyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) AddGroupToGroupSetDependency(context.Context, *AddGroupToGroupSetDependencyRequest) (*AddGroupToGroupSetDependencyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) NodeGroupSetControl(context.Context, *NodeGroupSetControlRequest) (*NodeGroupSetControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) GroupSetControl(context.Context, *GroupSetControlRequest) (*GroupSetControlResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) SetGroupDependencyExpression(context.Context, *SetGroupDependencyExpressionRequest) (*SetGroupDependencyExpressionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) RemoveClusterGroupDependency(context.Context, *RemoveClusterGroupDependencyRequest) (*RemoveClusterGroupDependencyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) SetGroupSetDependencyExpression(context.Context, *SetGroupSetDependencyExpressionRequest) (*SetGroupSetDependencyExpressionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) RemoveGroupSetDependency(context.Context, *RemoveGroupSetDependencyRequest) (*RemoveGroupSetDependencyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) RemoveClusterGroupToGroupSetDependency(context.Context, *RemoveClusterGroupToGroupSetDependencyRequest) (*RemoveClusterGroupToGroupSetDependencyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateGroupSetEnum(context.Context, *CreateGroupSetEnumRequest) (*CreateGroupSetEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) CreateNetInterfaceEnum(context.Context, *CreateNetInterfaceEnumRequest) (*CreateNetInterfaceEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) ChangeCSVStateEx(context.Context, *ChangeCSVStateExRequest) (*ChangeCSVStateExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) AddGroupToGroupSetEx(context.Context, *AddGroupToGroupSetExRequest) (*AddGroupToGroupSetExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedClusapi3Server) ChangeResourceGroupEx(context.Context, *ChangeResourceGroupExRequest) (*ChangeResourceGroupExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ Clusapi3Server = (*UnimplementedClusapi3Server)(nil)
