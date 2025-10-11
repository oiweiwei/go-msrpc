package iupdatesearcher

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
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
	_ = idispatch.GoPackage
)

// IUpdateSearcher server interface.
type UpdateSearcherServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The IUpdateSearcher::CanAutomaticallyUpgradeService (opnum 9) method sets whether
	// the update agent can automatically upgrade itself before performing a search initiated
	// through this instance.
	//
	// The IUpdateSearcher::CanAutomaticallyUpgradeService (opnum 8) method retrieves whether
	// the update agent can automatically upgrade itself before performing a search initiated
	// through this instance.
	//
	// Return Values: The method MUST return information an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// If value is neither VARIANT_TRUE nor VARIANT_FALSE, this method MUST return an error.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD set the value of the CanAutomaticallyUpgradeService ADM element
	// to the contents of the value parameter.
	GetCanAutomaticallyUpgradeService(context.Context, *GetCanAutomaticallyUpgradeServiceRequest) (*GetCanAutomaticallyUpgradeServiceResponse, error)

	// The IUpdateSearcher::CanAutomaticallyUpgradeService (opnum 9) method sets whether
	// the update agent can automatically upgrade itself before performing a search initiated
	// through this instance.
	//
	// The IUpdateSearcher::CanAutomaticallyUpgradeService (opnum 8) method retrieves whether
	// the update agent can automatically upgrade itself before performing a search initiated
	// through this instance.
	//
	// Return Values: The method MUST return information an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// If value is neither VARIANT_TRUE nor VARIANT_FALSE, this method MUST return an error.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD set the value of the CanAutomaticallyUpgradeService ADM element
	// to the contents of the value parameter.
	SetCanAutomaticallyUpgradeService(context.Context, *SetCanAutomaticallyUpgradeServiceRequest) (*SetCanAutomaticallyUpgradeServiceResponse, error)

	// The IUpdateSearcher::ClientApplicationID (opnum 10) method retrieves the string used
	// to identify the current client application.
	//
	// The IUpdateSession::ClientApplicationID (opnum 9) method sets the identifier of the
	// calling application.
	//
	// The IUpdateHistoryEntry::ClientApplicationID (opnum 16) method retrieves the ID of
	// the application that initiated the operation.
	//
	// The IUpdateServiceManager2::ClientApplicationID (opnum 16) method sets a string that
	// identifies the client application that is using this interface.
	//
	// The IUpdateSession::ClientApplicationID (opnum 8) method retrieves the identifier
	// of the calling application.
	//
	// The IUpdateSearcher::ClientApplicationID (opnum 11) method sets the string used to
	// identify the current client application.
	//
	// The IUpdateServiceManager2::ClientApplicationID (opnum 15) method retrieves a string
	// that identifies the client application that is using this interface.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the ClientApplicationID ADM element.
	GetClientApplicationID(context.Context, *GetClientApplicationIDRequest) (*GetClientApplicationIDResponse, error)

	// The IUpdateSearcher::ClientApplicationID (opnum 10) method retrieves the string used
	// to identify the current client application.
	//
	// The IUpdateSession::ClientApplicationID (opnum 9) method sets the identifier of the
	// calling application.
	//
	// The IUpdateHistoryEntry::ClientApplicationID (opnum 16) method retrieves the ID of
	// the application that initiated the operation.
	//
	// The IUpdateServiceManager2::ClientApplicationID (opnum 16) method sets a string that
	// identifies the client application that is using this interface.
	//
	// The IUpdateSession::ClientApplicationID (opnum 8) method retrieves the identifier
	// of the calling application.
	//
	// The IUpdateSearcher::ClientApplicationID (opnum 11) method sets the string used to
	// identify the current client application.
	//
	// The IUpdateServiceManager2::ClientApplicationID (opnum 15) method retrieves a string
	// that identifies the client application that is using this interface.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the ClientApplicationID ADM element.
	SetClientApplicationID(context.Context, *SetClientApplicationIDRequest) (*SetClientApplicationIDResponse, error)

	// The IUpdateSearcher::IncludePotentiallySupersededUpdates (opnum 12) method retrieves
	// whether the search results include updates that are superseded by other updates in
	// the search results.
	//
	// The IUpdateSearcher::IncludePotentiallySupersededUpdates (opnum 13) method sets whether
	// the search results include updates that are superseded by other updates in the search
	// results.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the IncludePotentiallySupersededUpdates ADM
	// element.
	GetIncludePotentiallySupersededUpdates(context.Context, *GetIncludePotentiallySupersededUpdatesRequest) (*GetIncludePotentiallySupersededUpdatesResponse, error)

	// The IUpdateSearcher::IncludePotentiallySupersededUpdates (opnum 12) method retrieves
	// whether the search results include updates that are superseded by other updates in
	// the search results.
	//
	// The IUpdateSearcher::IncludePotentiallySupersededUpdates (opnum 13) method sets whether
	// the search results include updates that are superseded by other updates in the search
	// results.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the IncludePotentiallySupersededUpdates ADM
	// element.
	SetIncludePotentiallySupersededUpdates(context.Context, *SetIncludePotentiallySupersededUpdatesRequest) (*SetIncludePotentiallySupersededUpdatesResponse, error)

	// The ServerSelection enumeration defines values that describe the type of server to
	// use for an update search operation.
	//
	// The IUpdateSearcher::ServerSelection (opnum 14) method retrieves the type of update
	// server used to search against.
	//
	// The IUpdateSearcher::ServerSelection (Opnum 15) method sets the type of update server
	// used to search against.
	//
	// The IUpdateHistoryEntry::ServerSelection (opnum 17) method describes the type of
	// update service that provided the update for which the operation was performed.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the ServerSelection ADM element.
	GetServerSelection(context.Context, *GetServerSelectionRequest) (*GetServerSelectionResponse, error)

	// The ServerSelection enumeration defines values that describe the type of server to
	// use for an update search operation.
	//
	// The IUpdateSearcher::ServerSelection (opnum 14) method retrieves the type of update
	// server used to search against.
	//
	// The IUpdateSearcher::ServerSelection (Opnum 15) method sets the type of update server
	// used to search against.
	//
	// The IUpdateHistoryEntry::ServerSelection (opnum 17) method describes the type of
	// update service that provided the update for which the operation was performed.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the ServerSelection ADM element.
	SetServerSelection(context.Context, *SetServerSelectionRequest) (*SetServerSelectionResponse, error)

	// Opnum16NotUsedOnWire operation.
	// Opnum16NotUsedOnWire

	// Opnum17NotUsedOnWire operation.
	// Opnum17NotUsedOnWire

	// The IUpdateSearcher::EscapeString (opnum 18) method escapes a string such that it
	// can be used as a literal value in a search criteria string.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method MUST return the escaped string.
	EscapeString(context.Context, *EscapeStringRequest) (*EscapeStringResponse, error)

	// The IUpdateSearcher::QueryHistory (opnum 19) method retrieves a collection of history
	// events.
	//
	// The IUpdateSession3::QueryHistory (Opnum 17) method retrieves relevant update history
	// entries.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	QueryHistory(context.Context, *QueryHistoryRequest) (*QueryHistoryResponse, error)

	// The IUpdateSearcher::Search (opnum 20) method performs an update search.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The server SHOULD trigger the update agent to perform an update search through an
	// implementation-dependent<41> interface.
	//
	// The update search MUST be restricted to updates meeting the criteria specified. If
	// the criteria parameter is NULL or the empty string, the server SHOULD trigger the
	// update agent to use default search criteria, as defined by the update agent.
	//
	// If the value of the CanAutomaticallyUpgradeService ADM element is VARIANT_TRUE, the
	// search operation MAY automatically upgrade the update agent.
	//
	// If the value of the IncludePotentiallySupersededUpdates ADM element is VARIANT_TRUE,
	// the search operation results SHOULD include updates superseded by other updates in
	// the search results. Otherwise, the search operation results SHOULD NOT include updates
	// superseded by other updates in the search results.
	//
	// If the value of the Online ADM element is VARIANT_TRUE, the search operation SHOULD
	// go online.
	//
	// The following table specifies the update server against which the search operation
	// SHOULD be performed.
	//
	//	+------------------------------------------+-------------------------------------------------------------------------+
	//	|     VALUE OF THE SERVERSELECTION ADM     |                         UPDATE SERVER TO SEARCH                         |
	//	|                 ELEMENT                  |                                 AGAINST                                 |
	//	+------------------------------------------+-------------------------------------------------------------------------+
	//	+------------------------------------------+-------------------------------------------------------------------------+
	//	| ssDefault                                | Implementation-defined by the update agent.                             |
	//	+------------------------------------------+-------------------------------------------------------------------------+
	//	| ssManagedServer                          | An update server managed by an administrator.                           |
	//	+------------------------------------------+-------------------------------------------------------------------------+
	//	| ssWindowsUpdate                          | The Windows Update update server                                        |
	//	+------------------------------------------+-------------------------------------------------------------------------+
	//	| ssOthers                                 | The update server identified by the value of the ServiceID ADM element. |
	//	+------------------------------------------+-------------------------------------------------------------------------+
	//
	// If the value of the IgnoreDownloadPriority ADM element is VARIANT_TRUE, the search
	// operation SHOULD ignore the download priority of updates when computing update supersedence.
	//
	// The following table specifies the set of updates for which the search operation SHOULD
	// be performed.
	//
	//	+--------------------------------------+----------------------------------------------------------------+
	//	|     VALUE OF THE SEARCHSCOPE ADM     |                       UPDATES TO SEARCH                        |
	//	|               ELEMENT                |                              FOR                               |
	//	+--------------------------------------+----------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------+
	//	| searchScopeDefault                   | Per-machine updates only.                                      |
	//	+--------------------------------------+----------------------------------------------------------------+
	//	| searchScopeMachineOnly               | Per-machine updates only.                                      |
	//	+--------------------------------------+----------------------------------------------------------------+
	//	| searchScopeCurrentUserOnly           | Per-user updates for the calling user only.                    |
	//	+--------------------------------------+----------------------------------------------------------------+
	//	| searchScopeMachineAndCurrentUser     | Per-machine updates and per-user updates for the calling user. |
	//	+--------------------------------------+----------------------------------------------------------------+
	//	| searchScopeMachineAndAllUsers        | Per-machine updates and per-user updates for all users.        |
	//	+--------------------------------------+----------------------------------------------------------------+
	//	| searchScopeAllUsers                  | Per-user updates for all users.                                |
	//	+--------------------------------------+----------------------------------------------------------------+
	Search(context.Context, *SearchRequest) (*SearchResponse, error)

	// The IUpdateSearcher::Online (opnum 22) method sets whether a search performed by
	// using this interface is an online search.
	//
	// The IUpdateSearcher::Online (opnum 21) method retrieves whether a search performed
	// by using this interface is an online search.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD set the value of the Online ADM element to the value parameter.
	GetOnline(context.Context, *GetOnlineRequest) (*GetOnlineResponse, error)

	// The IUpdateSearcher::Online (opnum 22) method sets whether a search performed by
	// using this interface is an online search.
	//
	// The IUpdateSearcher::Online (opnum 21) method retrieves whether a search performed
	// by using this interface is an online search.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD set the value of the Online ADM element to the value parameter.
	SetOnline(context.Context, *SetOnlineRequest) (*SetOnlineResponse, error)

	// The IUpdateSearcher::GetTotalHistoryCount (opnum 23) method retrieves the total number
	// of history events stored on the computer.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the HistoryCount ADM element.
	GetTotalHistoryCount(context.Context, *GetTotalHistoryCountRequest) (*GetTotalHistoryCountResponse, error)

	// The IUpdateService::ServiceID (opnum 16) method retrieves the unique identifier for
	// the update service.
	//
	// The IUpdateHistoryEntry::ServiceID (opnum 18) method retrieves the unique identifier
	// of the update service that provided the update for which the operation was performed.
	//
	// The IUpdateSearcher::ServiceID (opnum 24) method retrieves the unique identifier
	// of the update server used to search against.
	//
	// The IUpdateSearcher::ServiceID (opnum 25) method sets the unique identifier of the
	// update server used to search against.
	//
	// The IUpdateServiceRegistration::ServiceID (opnum 9) method retrieves the service
	// identifier.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the ServiceID ADM element.
	GetServiceID(context.Context, *GetServiceIDRequest) (*GetServiceIDResponse, error)

	// The IUpdateService::ServiceID (opnum 16) method retrieves the unique identifier for
	// the update service.
	//
	// The IUpdateHistoryEntry::ServiceID (opnum 18) method retrieves the unique identifier
	// of the update service that provided the update for which the operation was performed.
	//
	// The IUpdateSearcher::ServiceID (opnum 24) method retrieves the unique identifier
	// of the update server used to search against.
	//
	// The IUpdateSearcher::ServiceID (opnum 25) method sets the unique identifier of the
	// update server used to search against.
	//
	// The IUpdateServiceRegistration::ServiceID (opnum 9) method retrieves the service
	// identifier.
	//
	// Return Values: The method MUST return information in an HRESULT data structure. The
	// severity bit in the structure identifies the following conditions:
	//
	// * If the severity bit is set to 0, the method completed successfully.
	//
	// * If the severity bit is set to 1, the method failed and encountered a fatal error.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// This method SHOULD return the value of the ServiceID ADM element.
	SetServiceID(context.Context, *SetServiceIDRequest) (*SetServiceIDResponse, error)
}

func RegisterUpdateSearcherServer(conn dcerpc.Conn, o UpdateSearcherServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewUpdateSearcherServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(UpdateSearcherSyntaxV0_0))...)
}

func NewUpdateSearcherServerHandle(o UpdateSearcherServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return UpdateSearcherServerHandle(ctx, o, opNum, r)
	}
}

func UpdateSearcherServerHandle(ctx context.Context, o UpdateSearcherServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // CanAutomaticallyUpgradeService
		op := &xxx_GetCanAutomaticallyUpgradeServiceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCanAutomaticallyUpgradeServiceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCanAutomaticallyUpgradeService(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // CanAutomaticallyUpgradeService
		op := &xxx_SetCanAutomaticallyUpgradeServiceOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetCanAutomaticallyUpgradeServiceRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetCanAutomaticallyUpgradeService(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ClientApplicationID
		op := &xxx_GetClientApplicationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetClientApplicationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetClientApplicationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // ClientApplicationID
		op := &xxx_SetClientApplicationIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetClientApplicationIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetClientApplicationID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // IncludePotentiallySupersededUpdates
		op := &xxx_GetIncludePotentiallySupersededUpdatesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetIncludePotentiallySupersededUpdatesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetIncludePotentiallySupersededUpdates(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // IncludePotentiallySupersededUpdates
		op := &xxx_SetIncludePotentiallySupersededUpdatesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetIncludePotentiallySupersededUpdatesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetIncludePotentiallySupersededUpdates(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // ServerSelection
		op := &xxx_GetServerSelectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServerSelectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServerSelection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // ServerSelection
		op := &xxx_SetServerSelectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetServerSelectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetServerSelection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // Opnum16NotUsedOnWire
		// Opnum16NotUsedOnWire
		return nil, nil
	case 16: // Opnum17NotUsedOnWire
		// Opnum17NotUsedOnWire
		return nil, nil
	case 17: // EscapeString
		op := &xxx_EscapeStringOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EscapeStringRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EscapeString(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // QueryHistory
		op := &xxx_QueryHistoryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &QueryHistoryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.QueryHistory(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // Search
		op := &xxx_SearchOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SearchRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Search(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // Online
		op := &xxx_GetOnlineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOnlineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOnline(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // Online
		op := &xxx_SetOnlineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetOnlineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetOnline(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // GetTotalHistoryCount
		op := &xxx_GetTotalHistoryCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTotalHistoryCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTotalHistoryCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // ServiceID
		op := &xxx_GetServiceIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServiceIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServiceID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // ServiceID
		op := &xxx_SetServiceIDOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetServiceIDRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetServiceID(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IUpdateSearcher
type UnimplementedUpdateSearcherServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedUpdateSearcherServer) GetCanAutomaticallyUpgradeService(context.Context, *GetCanAutomaticallyUpgradeServiceRequest) (*GetCanAutomaticallyUpgradeServiceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) SetCanAutomaticallyUpgradeService(context.Context, *SetCanAutomaticallyUpgradeServiceRequest) (*SetCanAutomaticallyUpgradeServiceResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) GetClientApplicationID(context.Context, *GetClientApplicationIDRequest) (*GetClientApplicationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) SetClientApplicationID(context.Context, *SetClientApplicationIDRequest) (*SetClientApplicationIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) GetIncludePotentiallySupersededUpdates(context.Context, *GetIncludePotentiallySupersededUpdatesRequest) (*GetIncludePotentiallySupersededUpdatesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) SetIncludePotentiallySupersededUpdates(context.Context, *SetIncludePotentiallySupersededUpdatesRequest) (*SetIncludePotentiallySupersededUpdatesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) GetServerSelection(context.Context, *GetServerSelectionRequest) (*GetServerSelectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) SetServerSelection(context.Context, *SetServerSelectionRequest) (*SetServerSelectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) EscapeString(context.Context, *EscapeStringRequest) (*EscapeStringResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) QueryHistory(context.Context, *QueryHistoryRequest) (*QueryHistoryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) GetOnline(context.Context, *GetOnlineRequest) (*GetOnlineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) SetOnline(context.Context, *SetOnlineRequest) (*SetOnlineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) GetTotalHistoryCount(context.Context, *GetTotalHistoryCountRequest) (*GetTotalHistoryCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) GetServiceID(context.Context, *GetServiceIDRequest) (*GetServiceIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedUpdateSearcherServer) SetServiceID(context.Context, *SetServiceIDRequest) (*SetServiceIDResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ UpdateSearcherServer = (*UnimplementedUpdateSearcherServer)(nil)
