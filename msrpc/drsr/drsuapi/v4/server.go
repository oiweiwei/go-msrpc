package drsuapi

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

// drsuapi server interface.
type DrsuapiServer interface {

	// The IDL_DRSBind method creates a context handle that is necessary to call any other
	// method in this interface.
	//
	// Return Values: 0 if successful, otherwise a Windows error code.
	//
	// Exceptions Thrown: This method does not throw exceptions beyond those thrown by the
	// underlying RPC protocol.
	Bind(context.Context, *BindRequest) (*BindResponse, error)

	// The IDL_DRSUnbind method destroys a context handle previously created by the IDL_DRSBind
	// method.
	//
	// Return Values: 0 if successful, or a Windows error code if a failure occurs.
	//
	// Exceptions Thrown: This method might throw the following exception beyond those thrown
	// by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE.
	Unbind(context.Context, *UnbindRequest) (*UnbindResponse, error)

	// The IDL_DRSReplicaSync method triggers replication from another DC.
	//
	// Return Values: 0 if successful, otherwise a Windows error code.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	SyncReplica(context.Context, *SyncReplicaRequest) (*SyncReplicaResponse, error)

	// The IDL_DRSGetNCChanges method replicates updates from an NC replica on the server.
	//
	// Return Values: 0 if successful, otherwise a Windows error code.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	GetNCChanges(context.Context, *GetNCChangesRequest) (*GetNCChangesResponse, error)

	// The IDL_DRSUpdateRefs method adds or deletes a value from the repsTo of a specified
	// NC replica.
	//
	// Return Values: 0 if successful, otherwise a Windows error code.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	UpdateReferences(context.Context, *UpdateReferencesRequest) (*UpdateReferencesResponse, error)

	// The IDL_DRSReplicaAdd method adds a replication source reference for the specified
	// NC.
	//
	// Return Values: 0 if successful, otherwise a Windows error code.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	AddReplica(context.Context, *AddReplicaRequest) (*AddReplicaResponse, error)

	// The IDL_DRSReplicaDel method deletes a replication source reference for the specified
	// NC.
	//
	// Return Values: 0 if successful, otherwise a Windows error code.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	DeleteReplica(context.Context, *DeleteReplicaRequest) (*DeleteReplicaResponse, error)

	// The IDL_DRSReplicaModify method updates the value for repsFrom for the NC replica.
	//
	// Return Values: 0 if successful, or a Windows error code if a failure occurs.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	ModifyReplica(context.Context, *ModifyReplicaRequest) (*ModifyReplicaResponse, error)

	// The IDL_DRSVerifyNames method resolves a sequence of object identities.
	//
	// Return Values: 0 if successful, otherwise a Windows error code.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	VerifyNames(context.Context, *VerifyNamesRequest) (*VerifyNamesResponse, error)

	// The IDL_DRSGetMemberships method retrieves group membership for an object.
	//
	// Return Values: 0 if successful, otherwise a Windows error code.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	GetMemberships(context.Context, *GetMembershipsRequest) (*GetMembershipsResponse, error)

	// The IDL_DRSInterDomainMove method is a helper method used in a cross-NC move LDAP
	// operation.
	//
	// Return Values: 0 if successful, otherwise a Windows error code.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	InterdomainMove(context.Context, *InterdomainMoveRequest) (*InterdomainMoveResponse, error)

	// If the server is the PDC emulator FSMO role owner, the IDL_DRSGetNT4ChangeLog method
	// returns either a sequence of PDC change log entries or the NT4 replication state,
	// or both, as requested by the client.
	//
	// Return Values: 0 or ERROR_MORE_DATA if successful; another Windows error code if
	// a failure occurred.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	GetNT4ChangeLog(context.Context, *GetNT4ChangeLogRequest) (*GetNT4ChangeLogResponse, error)

	// The IDL_DRSCrackNames method looks up each of a set of objects in the directory and
	// returns it to the caller in the requested format.
	//
	// Return Values: 0 if successful, otherwise a Windows error code.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	CrackNames(context.Context, *CrackNamesRequest) (*CrackNamesResponse, error)

	// The IDL_DRSWriteSPN method updates the set of SPNs on an object.
	//
	// Return Values: 0 if successful, or a Windows error code if a failure occurs.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	WriteSPN(context.Context, *WriteSPNRequest) (*WriteSPNResponse, error)

	// The IDL_DRSRemoveDsServer method removes the representation (also known as metadata)
	// of a DC from the directory.
	//
	// Return Values: 0 if successful or a Windows error code if a failure occurs.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	RemoveDSServer(context.Context, *RemoveDSServerRequest) (*RemoveDSServerResponse, error)

	// The IDL_DRSRemoveDsDomain method removes the representation (also known as metadata)
	// of a domain from the directory.
	//
	// Return Values: 0 if successful, or a Windows error code if a failure occurs.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	RemoveDSDomain(context.Context, *RemoveDSDomainRequest) (*RemoveDSDomainResponse, error)

	// The IDL_DRSDomainControllerInfo method retrieves information about DCs in a given
	// domain.
	//
	// Return Values: 0 if successful, otherwise a Windows error code.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	DomainControllerInfo(context.Context, *DomainControllerInfoRequest) (*DomainControllerInfoResponse, error)

	// The IDL_DRSAddEntry method adds one or more objects.
	//
	// Return Values: 0 if successful, otherwise a Windows error code.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	AddEntry(context.Context, *AddEntryRequest) (*AddEntryResponse, error)

	// The IDL_DRSExecuteKCC method validates the replication interconnections of DCs and
	// updates them if necessary.
	//
	// Return Values: 0 if successful, otherwise a Windows error code.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and ERROR_INVALID_PARAMETER.
	ExecuteKCC(context.Context, *ExecuteKCCRequest) (*ExecuteKCCResponse, error)

	// The IDL_DRSGetReplInfo method retrieves the replication state of the server.
	//
	// Return Values: 0 if successful, otherwise a Windows error code.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	GetReplicationInfo(context.Context, *GetReplicationInfoRequest) (*GetReplicationInfoResponse, error)

	// The IDL_DRSAddSidHistory method adds one or more SIDs to the sIDHistory attribute
	// of a given object.
	//
	// Return Values: 0 or one of the following Windows error codes: ERROR_DS_MUST_RUN_ON_DST_DC
	// or ERROR_INVALID_PARAMETER.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and ERROR_INVALID_PARAMETER.
	AddSIDHistory(context.Context, *AddSIDHistoryRequest) (*AddSIDHistoryResponse, error)

	// The IDL_DRSGetMemberships2 method retrieves group memberships for a sequence of objects.
	//
	// Return Values: 0 if successful; otherwise, a Windows error code.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	GetMemberships2(context.Context, *GetMemberships2Request) (*GetMemberships2Response, error)

	// The IDL_DRSReplicaVerifyObjects method verifies the existence of objects in an NC
	// replica by comparing against a replica of the same NC on a reference DC, optionally
	// deleting any objects that do not exist on the reference DC.
	//
	// Return Values: 0 if successful, otherwise a Windows error code.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	VerifyObjectsReplica(context.Context, *VerifyObjectsReplicaRequest) (*VerifyObjectsReplicaResponse, error)

	// The IDL_DRSGetObjectExistence method helps the client check the consistency of object
	// existence between its replica of an NC and the server's replica of the same NC. Checking
	// the consistency of object existence means identifying objects that have replicated
	// to both replicas and that exist in one replica but not in the other. For the purposes
	// of this method, an object exists within a NC replica if it is either an object or
	// a tombstone.
	//
	// See IDL_DRSReplicaVerifyObjects for a use of this method.
	//
	// Return Values: 0 if successful, otherwise a Windows error code.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	GetObjectExistence(context.Context, *GetObjectExistenceRequest) (*GetObjectExistenceResponse, error)

	// The IDL_DRSQuerySitesByCost method determines the communication cost from a "from"
	// site to one or more "to" sites.
	//
	// Return Values: 0 if successful, or a Windows error code if a failure occurs.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	QuerySitesByCost(context.Context, *QuerySitesByCostRequest) (*QuerySitesByCostResponse, error)

	// The IDL_DRSInitDemotion method performs the first phase of the removal of a DC from
	// an AD LDS forest. This method is supported only by AD LDS.
	//
	// Return Values: 0 if successful, otherwise a Windows error code.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	InitDemotion(context.Context, *InitDemotionRequest) (*InitDemotionResponse, error)

	// The IDL_DRSReplicaDemotion method replicates off all changes to the specified NC
	// and moves any FSMOs held to another server.
	//
	// Return Values: 0 if successful, otherwise a Windows error code.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	DemotionReplica(context.Context, *DemotionReplicaRequest) (*DemotionReplicaResponse, error)

	// The IDL_DRSFinishDemotion method either performs one or more steps toward the complete
	// removal of a DC from an AD LDS forest, or it undoes the effects of the first phase
	// of removal (performed by IDL_DRSInitDemotion). This method is supported by AD LDS
	// only.
	//
	// Return Values: 0 if successful, otherwise a Windows error code.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, ERROR_DS_DIFFERENT_REPL_EPOCHS, and  ERROR_INVALID_PARAMETER.
	FinishDemotion(context.Context, *FinishDemotionRequest) (*FinishDemotionResponse, error)

	// The IDL_DRSAddCloneDC method is used to create a new DC object by copying attributes
	// from an existing DC object.
	//
	// Return Values: 0 if successful, otherwise a Windows error code.
	//
	// Exceptions Thrown: This method might throw the following exceptions beyond those
	// thrown by the underlying RPC protocol (as specified in [MS-RPCE]): ERROR_INVALID_HANDLE,
	// ERROR_DS_DRS_EXTENSIONS_CHANGED, and ERROR_INVALID_PARAMETER.
	AddCloneDC(context.Context, *AddCloneDCRequest) (*AddCloneDCResponse, error)

	// IDL_DRSWriteNgcKey operation.
	WriteNGCKey(context.Context, *WriteNGCKeyRequest) (*WriteNGCKeyResponse, error)

	// IDL_DRSReadNgcKey operation.
	ReadNGCKey(context.Context, *ReadNGCKeyRequest) (*ReadNGCKeyResponse, error)
}

func RegisterDrsuapiServer(conn dcerpc.Conn, o DrsuapiServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDrsuapiServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DrsuapiSyntaxV4_0))...)
}

func NewDrsuapiServerHandle(o DrsuapiServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DrsuapiServerHandle(ctx, o, opNum, r)
	}
}

func DrsuapiServerHandle(ctx context.Context, o DrsuapiServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // IDL_DRSBind
		op := &xxx_BindOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BindRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Bind(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // IDL_DRSUnbind
		op := &xxx_UnbindOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UnbindRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Unbind(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // IDL_DRSReplicaSync
		op := &xxx_SyncReplicaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SyncReplicaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SyncReplica(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // IDL_DRSGetNCChanges
		op := &xxx_GetNCChangesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNCChangesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNCChanges(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // IDL_DRSUpdateRefs
		op := &xxx_UpdateReferencesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UpdateReferencesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UpdateReferences(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // IDL_DRSReplicaAdd
		op := &xxx_AddReplicaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddReplicaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddReplica(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // IDL_DRSReplicaDel
		op := &xxx_DeleteReplicaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteReplicaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteReplica(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // IDL_DRSReplicaModify
		op := &xxx_ModifyReplicaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ModifyReplicaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ModifyReplica(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // IDL_DRSVerifyNames
		op := &xxx_VerifyNamesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &VerifyNamesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.VerifyNames(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // IDL_DRSGetMemberships
		op := &xxx_GetMembershipsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMembershipsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMemberships(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // IDL_DRSInterDomainMove
		op := &xxx_InterdomainMoveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InterdomainMoveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InterdomainMove(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // IDL_DRSGetNT4ChangeLog
		op := &xxx_GetNT4ChangeLogOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNT4ChangeLogRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNT4ChangeLog(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // IDL_DRSCrackNames
		op := &xxx_CrackNamesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CrackNamesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CrackNames(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // IDL_DRSWriteSPN
		op := &xxx_WriteSPNOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WriteSPNRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WriteSPN(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // IDL_DRSRemoveDsServer
		op := &xxx_RemoveDSServerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveDSServerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveDSServer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // IDL_DRSRemoveDsDomain
		op := &xxx_RemoveDSDomainOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveDSDomainRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveDSDomain(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // IDL_DRSDomainControllerInfo
		op := &xxx_DomainControllerInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DomainControllerInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DomainControllerInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // IDL_DRSAddEntry
		op := &xxx_AddEntryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddEntryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddEntry(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // IDL_DRSExecuteKCC
		op := &xxx_ExecuteKCCOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExecuteKCCRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExecuteKCC(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // IDL_DRSGetReplInfo
		op := &xxx_GetReplicationInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetReplicationInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetReplicationInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // IDL_DRSAddSidHistory
		op := &xxx_AddSIDHistoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddSIDHistoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddSIDHistory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // IDL_DRSGetMemberships2
		op := &xxx_GetMemberships2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMemberships2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMemberships2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // IDL_DRSReplicaVerifyObjects
		op := &xxx_VerifyObjectsReplicaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &VerifyObjectsReplicaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.VerifyObjectsReplica(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // IDL_DRSGetObjectExistence
		op := &xxx_GetObjectExistenceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetObjectExistenceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetObjectExistence(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // IDL_DRSQuerySitesByCost
		op := &xxx_QuerySitesByCostOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QuerySitesByCostRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QuerySitesByCost(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // IDL_DRSInitDemotion
		op := &xxx_InitDemotionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &InitDemotionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.InitDemotion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // IDL_DRSReplicaDemotion
		op := &xxx_DemotionReplicaOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DemotionReplicaRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DemotionReplica(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // IDL_DRSFinishDemotion
		op := &xxx_FinishDemotionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FinishDemotionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FinishDemotion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // IDL_DRSAddCloneDC
		op := &xxx_AddCloneDCOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddCloneDCRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddCloneDC(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // IDL_DRSWriteNgcKey
		op := &xxx_WriteNGCKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WriteNGCKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WriteNGCKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // IDL_DRSReadNgcKey
		op := &xxx_ReadNGCKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReadNGCKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReadNGCKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented drsuapi
type UnimplementedDrsuapiServer struct {
}

func (UnimplementedDrsuapiServer) Bind(context.Context, *BindRequest) (*BindResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) Unbind(context.Context, *UnbindRequest) (*UnbindResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) SyncReplica(context.Context, *SyncReplicaRequest) (*SyncReplicaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) GetNCChanges(context.Context, *GetNCChangesRequest) (*GetNCChangesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) UpdateReferences(context.Context, *UpdateReferencesRequest) (*UpdateReferencesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) AddReplica(context.Context, *AddReplicaRequest) (*AddReplicaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) DeleteReplica(context.Context, *DeleteReplicaRequest) (*DeleteReplicaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) ModifyReplica(context.Context, *ModifyReplicaRequest) (*ModifyReplicaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) VerifyNames(context.Context, *VerifyNamesRequest) (*VerifyNamesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) GetMemberships(context.Context, *GetMembershipsRequest) (*GetMembershipsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) InterdomainMove(context.Context, *InterdomainMoveRequest) (*InterdomainMoveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) GetNT4ChangeLog(context.Context, *GetNT4ChangeLogRequest) (*GetNT4ChangeLogResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) CrackNames(context.Context, *CrackNamesRequest) (*CrackNamesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) WriteSPN(context.Context, *WriteSPNRequest) (*WriteSPNResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) RemoveDSServer(context.Context, *RemoveDSServerRequest) (*RemoveDSServerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) RemoveDSDomain(context.Context, *RemoveDSDomainRequest) (*RemoveDSDomainResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) DomainControllerInfo(context.Context, *DomainControllerInfoRequest) (*DomainControllerInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) AddEntry(context.Context, *AddEntryRequest) (*AddEntryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) ExecuteKCC(context.Context, *ExecuteKCCRequest) (*ExecuteKCCResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) GetReplicationInfo(context.Context, *GetReplicationInfoRequest) (*GetReplicationInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) AddSIDHistory(context.Context, *AddSIDHistoryRequest) (*AddSIDHistoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) GetMemberships2(context.Context, *GetMemberships2Request) (*GetMemberships2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) VerifyObjectsReplica(context.Context, *VerifyObjectsReplicaRequest) (*VerifyObjectsReplicaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) GetObjectExistence(context.Context, *GetObjectExistenceRequest) (*GetObjectExistenceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) QuerySitesByCost(context.Context, *QuerySitesByCostRequest) (*QuerySitesByCostResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) InitDemotion(context.Context, *InitDemotionRequest) (*InitDemotionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) DemotionReplica(context.Context, *DemotionReplicaRequest) (*DemotionReplicaResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) FinishDemotion(context.Context, *FinishDemotionRequest) (*FinishDemotionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) AddCloneDC(context.Context, *AddCloneDCRequest) (*AddCloneDCResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) WriteNGCKey(context.Context, *WriteNGCKeyRequest) (*WriteNGCKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedDrsuapiServer) ReadNGCKey(context.Context, *ReadNGCKeyRequest) (*ReadNGCKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ DrsuapiServer = (*UnimplementedDrsuapiServer)(nil)
