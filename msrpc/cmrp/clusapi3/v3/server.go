package clusapi3

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
	case 109: // ApiSetResourceDependencyExpression
		in := &SetResourceDependencyExpressionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetResourceDependencyExpression(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 110: // ApiGetResourceDependencyExpression
		in := &GetResourceDependencyExpressionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetResourceDependencyExpression(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 111: // Opnum111NotUsedOnWire
		// Opnum111NotUsedOnWire
		return nil, nil
	case 112: // ApiGetResourceNetworkName
		in := &GetResourceNetworkNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetResourceNetworkName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 113: // ApiExecuteBatch
		in := &ExecuteBatchRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ExecuteBatch(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 114: // ApiCreateBatchPort
		in := &CreateBatchPortRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateBatchPort(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 115: // ApiGetBatchNotification
		in := &GetBatchNotificationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetBatchNotification(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 116: // ApiCloseBatchPort
		in := &CloseBatchPortRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseBatchPort(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 117: // ApiOpenClusterEx
		in := &OpenClusterExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenClusterEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 118: // ApiOpenNodeEx
		in := &OpenNodeExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenNodeEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 119: // ApiOpenGroupEx
		in := &OpenGroupExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenGroupEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 120: // ApiOpenResourceEx
		in := &OpenResourceExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenResourceEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 121: // ApiOpenNetworkEx
		in := &OpenNetworkExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenNetworkEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 122: // ApiOpenNetInterfaceEx
		in := &OpenNetInterfaceExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenNetInterfaceEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 123: // ApiChangeCsvState
		in := &ChangeCSVStateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ChangeCSVState(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 124: // ApiCreateNodeEnumEx
		in := &CreateNodeEnumExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateNodeEnumEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 125: // ApiCreateEnumEx
		in := &CreateEnumExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateEnumEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 126: // ApiPauseNodeEx
		in := &PauseNodeExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PauseNodeEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 127: // ApiPauseNodeWithDrainTarget
		in := &PauseNodeWithDrainTargetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PauseNodeWithDrainTarget(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 128: // ApiResumeNodeEx
		in := &ResumeNodeExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ResumeNodeEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 129: // ApiCreateGroupEx
		in := &CreateGroupExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateGroupEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 130: // ApiOnlineGroupEx
		in := &OnlineGroupExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OnlineGroupEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 131: // ApiOfflineGroupEx
		in := &OfflineGroupExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OfflineGroupEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 132: // ApiMoveGroupEx
		in := &MoveGroupExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.MoveGroupEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 133: // ApiMoveGroupToNodeEx
		in := &MoveGroupToNodeExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.MoveGroupToNodeEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 134: // ApiCancelClusterGroupOperation
		in := &CancelClusterGroupOperationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CancelClusterGroupOperation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 135: // ApiOnlineResourceEx
		in := &OnlineResourceExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OnlineResourceEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 136: // ApiOfflineResourceEx
		in := &OfflineResourceExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OfflineResourceEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 137: // ApiCreateNotifyV2
		in := &CreateNotifyV2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateNotifyV2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 138: // ApiAddNotifyV2
		in := &AddNotifyV2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddNotifyV2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 139: // ApiGetNotifyV2
		in := &GetNotifyV2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNotifyV2(ctx, in)
		return resp.xxx_ToOp(ctx), err
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
		in := &CreateGroupEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateGroupEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 144: // ApiCreateResourceEnum
		in := &CreateResourceEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateResourceEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 145: // ApiExecuteReadBatch
		in := &ExecuteReadBatchRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ExecuteReadBatch(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 146: // ApiRestartResource
		in := &RestartResourceRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RestartResource(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 147: // ApiGetNotifyAsync
		in := &GetNotifyAsyncRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNotifyAsync(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 148: // Opnum148NotUsedOnWire
		// Opnum148NotUsedOnWire
		return nil, nil
	case 149: // Opnum149otUsedOnWire
		in := &Opnum149otUsedOnWireRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Opnum149otUsedOnWire(ctx, in)
		return resp.xxx_ToOp(ctx), err
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
		in := &AddNotifyResourceTypeV2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddNotifyResourceTypeV2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 156: // Opnum156NotUsedOnWire
		// Opnum156NotUsedOnWire
		return nil, nil
	case 157: // ApiExecuteReadBatchEx
		in := &ExecuteReadBatchExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ExecuteReadBatchEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
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
		in := &CreateGroupSetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateGroupSet(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 164: // ApiOpenGroupSet
		in := &OpenGroupSetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.OpenGroupSet(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 165: // ApiCloseGroupSet
		in := &CloseGroupSetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CloseGroupSet(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 166: // ApiDeleteGroupSet
		in := &DeleteGroupSetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteGroupSet(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 167: // ApiAddGroupToGroupSet
		in := &AddGroupToGroupSetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddGroupToGroupSet(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 168: // ApiRemoveGroupFromGroupSet
		in := &RemoveGroupFromGroupSetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveGroupFromGroupSet(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 169: // ApiMoveGroupToGroupSet
		in := &MoveGroupToGroupSetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.MoveGroupToGroupSet(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 170: // Opnum170NotUsedOnWire
		// Opnum170NotUsedOnWire
		return nil, nil
	case 171: // ApiAddGroupSetDependency
		in := &AddGroupSetDependencyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddGroupSetDependency(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 172: // ApiAddGroupToGroupSetDependency
		in := &AddGroupToGroupSetDependencyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddGroupToGroupSetDependency(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 173: // ApiNodeGroupSetControl
		in := &NodeGroupSetControlRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.NodeGroupSetControl(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 174: // ApiGroupSetControl
		in := &GroupSetControlRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GroupSetControl(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 175: // ApiSetGroupDependencyExpression
		in := &SetGroupDependencyExpressionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetGroupDependencyExpression(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 176: // ApiRemoveClusterGroupDependency
		in := &RemoveClusterGroupDependencyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveClusterGroupDependency(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 177: // ApiSetGroupSetDependencyExpression
		in := &SetGroupSetDependencyExpressionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetGroupSetDependencyExpression(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 178: // ApiRemoveGroupSetDependency
		in := &RemoveGroupSetDependencyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveGroupSetDependency(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 179: // ApiRemoveClusterGroupToGroupSetDependency
		in := &RemoveClusterGroupToGroupSetDependencyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveClusterGroupToGroupSetDependency(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 180: // ApiCreateGroupSetEnum
		in := &CreateGroupSetEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateGroupSetEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 181: // ApiCreateNetInterfaceEnum
		in := &CreateNetInterfaceEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateNetInterfaceEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 182: // ApiChangeCsvStateEx
		in := &ChangeCSVStateExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ChangeCSVStateEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 183: // ApiAddGroupToGroupSetEx
		in := &AddGroupToGroupSetExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddGroupToGroupSetEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 184: // ApiChangeResourceGroupEx
		in := &ChangeResourceGroupExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ChangeResourceGroupEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
