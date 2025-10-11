package iupdatesearcher

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	oaut "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut"
	idispatch "github.com/oiweiwei/go-msrpc/msrpc/dcom/oaut/idispatch/v0"
	uamg "github.com/oiweiwei/go-msrpc/msrpc/dcom/uamg"
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
	_ = dcom.GoPackage
	_ = idispatch.GoPackage
	_ = oaut.GoPackage
	_ = uamg.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/uamg"
)

var (
	// IUpdateSearcher interface identifier 8f45abf1-f9ae-4b95-a933-f0f66e5056ea
	UpdateSearcherIID = &dcom.IID{Data1: 0x8f45abf1, Data2: 0xf9ae, Data3: 0x4b95, Data4: []byte{0xa9, 0x33, 0xf0, 0xf6, 0x6e, 0x50, 0x56, 0xea}}
	// Syntax UUID
	UpdateSearcherSyntaxUUID = &uuid.UUID{TimeLow: 0x8f45abf1, TimeMid: 0xf9ae, TimeHiAndVersion: 0x4b95, ClockSeqHiAndReserved: 0xa9, ClockSeqLow: 0x33, Node: [6]uint8{0xf0, 0xf6, 0x6e, 0x50, 0x56, 0xea}}
	// Syntax ID
	UpdateSearcherSyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: UpdateSearcherSyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IUpdateSearcher interface.
type UpdateSearcherClient interface {

	// IDispatch retrieval method.
	Dispatch() idispatch.DispatchClient

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
	GetCanAutomaticallyUpgradeService(context.Context, *GetCanAutomaticallyUpgradeServiceRequest, ...dcerpc.CallOption) (*GetCanAutomaticallyUpgradeServiceResponse, error)

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
	SetCanAutomaticallyUpgradeService(context.Context, *SetCanAutomaticallyUpgradeServiceRequest, ...dcerpc.CallOption) (*SetCanAutomaticallyUpgradeServiceResponse, error)

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
	GetClientApplicationID(context.Context, *GetClientApplicationIDRequest, ...dcerpc.CallOption) (*GetClientApplicationIDResponse, error)

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
	SetClientApplicationID(context.Context, *SetClientApplicationIDRequest, ...dcerpc.CallOption) (*SetClientApplicationIDResponse, error)

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
	GetIncludePotentiallySupersededUpdates(context.Context, *GetIncludePotentiallySupersededUpdatesRequest, ...dcerpc.CallOption) (*GetIncludePotentiallySupersededUpdatesResponse, error)

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
	SetIncludePotentiallySupersededUpdates(context.Context, *SetIncludePotentiallySupersededUpdatesRequest, ...dcerpc.CallOption) (*SetIncludePotentiallySupersededUpdatesResponse, error)

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
	GetServerSelection(context.Context, *GetServerSelectionRequest, ...dcerpc.CallOption) (*GetServerSelectionResponse, error)

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
	SetServerSelection(context.Context, *SetServerSelectionRequest, ...dcerpc.CallOption) (*SetServerSelectionResponse, error)

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
	EscapeString(context.Context, *EscapeStringRequest, ...dcerpc.CallOption) (*EscapeStringResponse, error)

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
	QueryHistory(context.Context, *QueryHistoryRequest, ...dcerpc.CallOption) (*QueryHistoryResponse, error)

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
	Search(context.Context, *SearchRequest, ...dcerpc.CallOption) (*SearchResponse, error)

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
	GetOnline(context.Context, *GetOnlineRequest, ...dcerpc.CallOption) (*GetOnlineResponse, error)

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
	SetOnline(context.Context, *SetOnlineRequest, ...dcerpc.CallOption) (*SetOnlineResponse, error)

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
	GetTotalHistoryCount(context.Context, *GetTotalHistoryCountRequest, ...dcerpc.CallOption) (*GetTotalHistoryCountResponse, error)

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
	GetServiceID(context.Context, *GetServiceIDRequest, ...dcerpc.CallOption) (*GetServiceIDResponse, error)

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
	SetServiceID(context.Context, *SetServiceIDRequest, ...dcerpc.CallOption) (*SetServiceIDResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) UpdateSearcherClient
}

type xxx_DefaultUpdateSearcherClient struct {
	idispatch.DispatchClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultUpdateSearcherClient) Dispatch() idispatch.DispatchClient {
	return o.DispatchClient
}

func (o *xxx_DefaultUpdateSearcherClient) GetCanAutomaticallyUpgradeService(ctx context.Context, in *GetCanAutomaticallyUpgradeServiceRequest, opts ...dcerpc.CallOption) (*GetCanAutomaticallyUpgradeServiceResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetCanAutomaticallyUpgradeServiceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSearcherClient) SetCanAutomaticallyUpgradeService(ctx context.Context, in *SetCanAutomaticallyUpgradeServiceRequest, opts ...dcerpc.CallOption) (*SetCanAutomaticallyUpgradeServiceResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetCanAutomaticallyUpgradeServiceResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSearcherClient) GetClientApplicationID(ctx context.Context, in *GetClientApplicationIDRequest, opts ...dcerpc.CallOption) (*GetClientApplicationIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetClientApplicationIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSearcherClient) SetClientApplicationID(ctx context.Context, in *SetClientApplicationIDRequest, opts ...dcerpc.CallOption) (*SetClientApplicationIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetClientApplicationIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSearcherClient) GetIncludePotentiallySupersededUpdates(ctx context.Context, in *GetIncludePotentiallySupersededUpdatesRequest, opts ...dcerpc.CallOption) (*GetIncludePotentiallySupersededUpdatesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetIncludePotentiallySupersededUpdatesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSearcherClient) SetIncludePotentiallySupersededUpdates(ctx context.Context, in *SetIncludePotentiallySupersededUpdatesRequest, opts ...dcerpc.CallOption) (*SetIncludePotentiallySupersededUpdatesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetIncludePotentiallySupersededUpdatesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSearcherClient) GetServerSelection(ctx context.Context, in *GetServerSelectionRequest, opts ...dcerpc.CallOption) (*GetServerSelectionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetServerSelectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSearcherClient) SetServerSelection(ctx context.Context, in *SetServerSelectionRequest, opts ...dcerpc.CallOption) (*SetServerSelectionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetServerSelectionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSearcherClient) EscapeString(ctx context.Context, in *EscapeStringRequest, opts ...dcerpc.CallOption) (*EscapeStringResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EscapeStringResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSearcherClient) QueryHistory(ctx context.Context, in *QueryHistoryRequest, opts ...dcerpc.CallOption) (*QueryHistoryResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryHistoryResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSearcherClient) Search(ctx context.Context, in *SearchRequest, opts ...dcerpc.CallOption) (*SearchResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SearchResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSearcherClient) GetOnline(ctx context.Context, in *GetOnlineRequest, opts ...dcerpc.CallOption) (*GetOnlineResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetOnlineResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSearcherClient) SetOnline(ctx context.Context, in *SetOnlineRequest, opts ...dcerpc.CallOption) (*SetOnlineResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetOnlineResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSearcherClient) GetTotalHistoryCount(ctx context.Context, in *GetTotalHistoryCountRequest, opts ...dcerpc.CallOption) (*GetTotalHistoryCountResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetTotalHistoryCountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSearcherClient) GetServiceID(ctx context.Context, in *GetServiceIDRequest, opts ...dcerpc.CallOption) (*GetServiceIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetServiceIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSearcherClient) SetServiceID(ctx context.Context, in *SetServiceIDRequest, opts ...dcerpc.CallOption) (*SetServiceIDResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetServiceIDResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultUpdateSearcherClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultUpdateSearcherClient) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultUpdateSearcherClient) IPID(ctx context.Context, ipid *dcom.IPID) UpdateSearcherClient {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultUpdateSearcherClient{
		DispatchClient: o.DispatchClient.IPID(ctx, ipid),
		cc:             o.cc,
		ipid:           ipid,
	}
}

func NewUpdateSearcherClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (UpdateSearcherClient, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(UpdateSearcherSyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := idispatch.NewDispatchClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultUpdateSearcherClient{
		DispatchClient: base,
		cc:             cc,
		ipid:           ipid,
	}, nil
}

// xxx_GetCanAutomaticallyUpgradeServiceOperation structure represents the CanAutomaticallyUpgradeService operation
type xxx_GetCanAutomaticallyUpgradeServiceOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue int16          `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetCanAutomaticallyUpgradeServiceOperation) OpNum() int { return 7 }

func (o *xxx_GetCanAutomaticallyUpgradeServiceOperation) OpName() string {
	return "/IUpdateSearcher/v0/CanAutomaticallyUpgradeService"
}

func (o *xxx_GetCanAutomaticallyUpgradeServiceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCanAutomaticallyUpgradeServiceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCanAutomaticallyUpgradeServiceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCanAutomaticallyUpgradeServiceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCanAutomaticallyUpgradeServiceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.ReturnValue); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetCanAutomaticallyUpgradeServiceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.ReturnValue); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetCanAutomaticallyUpgradeServiceRequest structure represents the CanAutomaticallyUpgradeService operation request
type GetCanAutomaticallyUpgradeServiceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetCanAutomaticallyUpgradeServiceRequest) xxx_ToOp(ctx context.Context, op *xxx_GetCanAutomaticallyUpgradeServiceOperation) *xxx_GetCanAutomaticallyUpgradeServiceOperation {
	if op == nil {
		op = &xxx_GetCanAutomaticallyUpgradeServiceOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetCanAutomaticallyUpgradeServiceRequest) xxx_FromOp(ctx context.Context, op *xxx_GetCanAutomaticallyUpgradeServiceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetCanAutomaticallyUpgradeServiceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetCanAutomaticallyUpgradeServiceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCanAutomaticallyUpgradeServiceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetCanAutomaticallyUpgradeServiceResponse structure represents the CanAutomaticallyUpgradeService operation response
type GetCanAutomaticallyUpgradeServiceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: MUST be VARIANT_TRUE if the update agent can automatically upgrade itself
	// before performing a search initiated through this instance, or VARIANT_FALSE otherwise.
	ReturnValue int16 `idl:"name:retval" json:"return_value"`
	// Return: The CanAutomaticallyUpgradeService return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetCanAutomaticallyUpgradeServiceResponse) xxx_ToOp(ctx context.Context, op *xxx_GetCanAutomaticallyUpgradeServiceOperation) *xxx_GetCanAutomaticallyUpgradeServiceOperation {
	if op == nil {
		op = &xxx_GetCanAutomaticallyUpgradeServiceOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetCanAutomaticallyUpgradeServiceResponse) xxx_FromOp(ctx context.Context, op *xxx_GetCanAutomaticallyUpgradeServiceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetCanAutomaticallyUpgradeServiceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetCanAutomaticallyUpgradeServiceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetCanAutomaticallyUpgradeServiceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetCanAutomaticallyUpgradeServiceOperation structure represents the CanAutomaticallyUpgradeService operation
type xxx_SetCanAutomaticallyUpgradeServiceOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Value  int16          `idl:"name:value" json:"value"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetCanAutomaticallyUpgradeServiceOperation) OpNum() int { return 8 }

func (o *xxx_SetCanAutomaticallyUpgradeServiceOperation) OpName() string {
	return "/IUpdateSearcher/v0/CanAutomaticallyUpgradeService"
}

func (o *xxx_SetCanAutomaticallyUpgradeServiceOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCanAutomaticallyUpgradeServiceOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// value {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Value); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCanAutomaticallyUpgradeServiceOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// value {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Value); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCanAutomaticallyUpgradeServiceOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCanAutomaticallyUpgradeServiceOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetCanAutomaticallyUpgradeServiceOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetCanAutomaticallyUpgradeServiceRequest structure represents the CanAutomaticallyUpgradeService operation request
type SetCanAutomaticallyUpgradeServiceRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// value: Set to VARIANT_TRUE to allow the update agent to automatically upgrade itself,
	// if it is an older version, before performing a search through this instance; otherwise,
	// set to VARIANT_FALSE to disallow automatic upgrading of the update agent.
	Value int16 `idl:"name:value" json:"value"`
}

func (o *SetCanAutomaticallyUpgradeServiceRequest) xxx_ToOp(ctx context.Context, op *xxx_SetCanAutomaticallyUpgradeServiceOperation) *xxx_SetCanAutomaticallyUpgradeServiceOperation {
	if op == nil {
		op = &xxx_SetCanAutomaticallyUpgradeServiceOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Value = o.Value
	return op
}

func (o *SetCanAutomaticallyUpgradeServiceRequest) xxx_FromOp(ctx context.Context, op *xxx_SetCanAutomaticallyUpgradeServiceOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Value = op.Value
}
func (o *SetCanAutomaticallyUpgradeServiceRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetCanAutomaticallyUpgradeServiceRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCanAutomaticallyUpgradeServiceOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetCanAutomaticallyUpgradeServiceResponse structure represents the CanAutomaticallyUpgradeService operation response
type SetCanAutomaticallyUpgradeServiceResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CanAutomaticallyUpgradeService return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetCanAutomaticallyUpgradeServiceResponse) xxx_ToOp(ctx context.Context, op *xxx_SetCanAutomaticallyUpgradeServiceOperation) *xxx_SetCanAutomaticallyUpgradeServiceOperation {
	if op == nil {
		op = &xxx_SetCanAutomaticallyUpgradeServiceOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetCanAutomaticallyUpgradeServiceResponse) xxx_FromOp(ctx context.Context, op *xxx_SetCanAutomaticallyUpgradeServiceOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetCanAutomaticallyUpgradeServiceResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetCanAutomaticallyUpgradeServiceResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetCanAutomaticallyUpgradeServiceOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetClientApplicationIDOperation structure represents the ClientApplicationID operation
type xxx_GetClientApplicationIDOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue *oaut.String   `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetClientApplicationIDOperation) OpNum() int { return 9 }

func (o *xxx_GetClientApplicationIDOperation) OpName() string {
	return "/IUpdateSearcher/v0/ClientApplicationID"
}

func (o *xxx_GetClientApplicationIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientApplicationIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientApplicationIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientApplicationIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientApplicationIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// retval {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ReturnValue != nil {
			_ptr_retval := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnValue != nil {
					if err := o.ReturnValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnValue, _ptr_retval); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetClientApplicationIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// retval {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_retval := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnValue == nil {
				o.ReturnValue = &oaut.String{}
			}
			if err := o.ReturnValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_retval := func(ptr interface{}) { o.ReturnValue = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ReturnValue, _s_retval, _ptr_retval); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetClientApplicationIDRequest structure represents the ClientApplicationID operation request
type GetClientApplicationIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetClientApplicationIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetClientApplicationIDOperation) *xxx_GetClientApplicationIDOperation {
	if op == nil {
		op = &xxx_GetClientApplicationIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetClientApplicationIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetClientApplicationIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetClientApplicationIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetClientApplicationIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClientApplicationIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetClientApplicationIDResponse structure represents the ClientApplicationID operation response
type GetClientApplicationIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: Returns a string identifying the client application using the interface.
	//
	// retval: A string identifying the client application that initiated the operation.
	//
	// retval: The identifier of the calling application previously set using the IUpdateSession::ClientApplicationID
	// (opnum 9) (section 3.16.4.2) method. If no identifier was previously set, this MUST
	// be NULL or the empty string.
	//
	// retval: A string identifying the client application that is using this interface.
	ReturnValue *oaut.String `idl:"name:retval" json:"return_value"`
	// Return: The ClientApplicationID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetClientApplicationIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetClientApplicationIDOperation) *xxx_GetClientApplicationIDOperation {
	if op == nil {
		op = &xxx_GetClientApplicationIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetClientApplicationIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetClientApplicationIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetClientApplicationIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetClientApplicationIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetClientApplicationIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetClientApplicationIDOperation structure represents the ClientApplicationID operation
type xxx_SetClientApplicationIDOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Value  *oaut.String   `idl:"name:value" json:"value"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetClientApplicationIDOperation) OpNum() int { return 10 }

func (o *xxx_SetClientApplicationIDOperation) OpName() string {
	return "/IUpdateSearcher/v0/ClientApplicationID"
}

func (o *xxx_SetClientApplicationIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientApplicationIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// value {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Value != nil {
			_ptr_value := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Value != nil {
					if err := o.Value.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Value, _ptr_value); err != nil {
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

func (o *xxx_SetClientApplicationIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// value {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_value := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Value == nil {
				o.Value = &oaut.String{}
			}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_value := func(ptr interface{}) { o.Value = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Value, _s_value, _ptr_value); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientApplicationIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientApplicationIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetClientApplicationIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetClientApplicationIDRequest structure represents the ClientApplicationID operation request
type SetClientApplicationIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// value: The identifier of the calling application.
	//
	// value: A string that identifies the client application that is using this interface.
	//
	// value: A string used to identify the client application using the interface.
	Value *oaut.String `idl:"name:value" json:"value"`
}

func (o *SetClientApplicationIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetClientApplicationIDOperation) *xxx_SetClientApplicationIDOperation {
	if op == nil {
		op = &xxx_SetClientApplicationIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Value = o.Value
	return op
}

func (o *SetClientApplicationIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetClientApplicationIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Value = op.Value
}
func (o *SetClientApplicationIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetClientApplicationIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClientApplicationIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetClientApplicationIDResponse structure represents the ClientApplicationID operation response
type SetClientApplicationIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ClientApplicationID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetClientApplicationIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetClientApplicationIDOperation) *xxx_SetClientApplicationIDOperation {
	if op == nil {
		op = &xxx_SetClientApplicationIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetClientApplicationIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetClientApplicationIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetClientApplicationIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetClientApplicationIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetClientApplicationIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetIncludePotentiallySupersededUpdatesOperation structure represents the IncludePotentiallySupersededUpdates operation
type xxx_GetIncludePotentiallySupersededUpdatesOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue int16          `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetIncludePotentiallySupersededUpdatesOperation) OpNum() int { return 11 }

func (o *xxx_GetIncludePotentiallySupersededUpdatesOperation) OpName() string {
	return "/IUpdateSearcher/v0/IncludePotentiallySupersededUpdates"
}

func (o *xxx_GetIncludePotentiallySupersededUpdatesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIncludePotentiallySupersededUpdatesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIncludePotentiallySupersededUpdatesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIncludePotentiallySupersededUpdatesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIncludePotentiallySupersededUpdatesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.ReturnValue); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetIncludePotentiallySupersededUpdatesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.ReturnValue); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetIncludePotentiallySupersededUpdatesRequest structure represents the IncludePotentiallySupersededUpdates operation request
type GetIncludePotentiallySupersededUpdatesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetIncludePotentiallySupersededUpdatesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetIncludePotentiallySupersededUpdatesOperation) *xxx_GetIncludePotentiallySupersededUpdatesOperation {
	if op == nil {
		op = &xxx_GetIncludePotentiallySupersededUpdatesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetIncludePotentiallySupersededUpdatesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetIncludePotentiallySupersededUpdatesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetIncludePotentiallySupersededUpdatesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetIncludePotentiallySupersededUpdatesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIncludePotentiallySupersededUpdatesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetIncludePotentiallySupersededUpdatesResponse structure represents the IncludePotentiallySupersededUpdates operation response
type GetIncludePotentiallySupersededUpdatesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: Indicates whether the results of searches performed by using this interface
	// include updates that are superseded by other updates.
	ReturnValue int16 `idl:"name:retval" json:"return_value"`
	// Return: The IncludePotentiallySupersededUpdates return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetIncludePotentiallySupersededUpdatesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetIncludePotentiallySupersededUpdatesOperation) *xxx_GetIncludePotentiallySupersededUpdatesOperation {
	if op == nil {
		op = &xxx_GetIncludePotentiallySupersededUpdatesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetIncludePotentiallySupersededUpdatesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetIncludePotentiallySupersededUpdatesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetIncludePotentiallySupersededUpdatesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetIncludePotentiallySupersededUpdatesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetIncludePotentiallySupersededUpdatesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetIncludePotentiallySupersededUpdatesOperation structure represents the IncludePotentiallySupersededUpdates operation
type xxx_SetIncludePotentiallySupersededUpdatesOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Value  int16          `idl:"name:value" json:"value"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetIncludePotentiallySupersededUpdatesOperation) OpNum() int { return 12 }

func (o *xxx_SetIncludePotentiallySupersededUpdatesOperation) OpName() string {
	return "/IUpdateSearcher/v0/IncludePotentiallySupersededUpdates"
}

func (o *xxx_SetIncludePotentiallySupersededUpdatesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetIncludePotentiallySupersededUpdatesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// value {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Value); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetIncludePotentiallySupersededUpdatesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// value {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Value); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetIncludePotentiallySupersededUpdatesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetIncludePotentiallySupersededUpdatesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetIncludePotentiallySupersededUpdatesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetIncludePotentiallySupersededUpdatesRequest structure represents the IncludePotentiallySupersededUpdates operation request
type SetIncludePotentiallySupersededUpdatesRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// value: Indicates whether the results of searches performed by using this interface
	// include updates that are superseded by other updates.
	Value int16 `idl:"name:value" json:"value"`
}

func (o *SetIncludePotentiallySupersededUpdatesRequest) xxx_ToOp(ctx context.Context, op *xxx_SetIncludePotentiallySupersededUpdatesOperation) *xxx_SetIncludePotentiallySupersededUpdatesOperation {
	if op == nil {
		op = &xxx_SetIncludePotentiallySupersededUpdatesOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Value = o.Value
	return op
}

func (o *SetIncludePotentiallySupersededUpdatesRequest) xxx_FromOp(ctx context.Context, op *xxx_SetIncludePotentiallySupersededUpdatesOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Value = op.Value
}
func (o *SetIncludePotentiallySupersededUpdatesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetIncludePotentiallySupersededUpdatesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetIncludePotentiallySupersededUpdatesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetIncludePotentiallySupersededUpdatesResponse structure represents the IncludePotentiallySupersededUpdates operation response
type SetIncludePotentiallySupersededUpdatesResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The IncludePotentiallySupersededUpdates return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetIncludePotentiallySupersededUpdatesResponse) xxx_ToOp(ctx context.Context, op *xxx_SetIncludePotentiallySupersededUpdatesOperation) *xxx_SetIncludePotentiallySupersededUpdatesOperation {
	if op == nil {
		op = &xxx_SetIncludePotentiallySupersededUpdatesOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetIncludePotentiallySupersededUpdatesResponse) xxx_FromOp(ctx context.Context, op *xxx_SetIncludePotentiallySupersededUpdatesOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetIncludePotentiallySupersededUpdatesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetIncludePotentiallySupersededUpdatesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetIncludePotentiallySupersededUpdatesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetServerSelectionOperation structure represents the ServerSelection operation
type xxx_GetServerSelectionOperation struct {
	This        *dcom.ORPCThis       `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat       `idl:"name:That" json:"that"`
	ReturnValue uamg.ServerSelection `idl:"name:retval" json:"return_value"`
	Return      int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_GetServerSelectionOperation) OpNum() int { return 13 }

func (o *xxx_GetServerSelectionOperation) OpName() string {
	return "/IUpdateSearcher/v0/ServerSelection"
}

func (o *xxx_GetServerSelectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerSelectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerSelectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerSelectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerSelectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{v1_enum, alias=ServerSelection}(enum))
	{
		if err := w.WriteEnum(uint32(o.ReturnValue)); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerSelectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{v1_enum, alias=ServerSelection}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.ReturnValue)); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetServerSelectionRequest structure represents the ServerSelection operation request
type GetServerSelectionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetServerSelectionRequest) xxx_ToOp(ctx context.Context, op *xxx_GetServerSelectionOperation) *xxx_GetServerSelectionOperation {
	if op == nil {
		op = &xxx_GetServerSelectionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetServerSelectionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetServerSelectionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetServerSelectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetServerSelectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServerSelectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetServerSelectionResponse structure represents the ServerSelection operation response
type GetServerSelectionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: A value from the ServerSelection (section 2.2.11) enumeration that indicates
	// the type of update server against which to perform search operations initiated by
	// using this interface.
	//
	// retval: A value from the ServerSelection (section 2.2.11) enumeration that indicates
	// the type of update service that provided the update for which the operation was performed.
	ReturnValue uamg.ServerSelection `idl:"name:retval" json:"return_value"`
	// Return: The ServerSelection return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetServerSelectionResponse) xxx_ToOp(ctx context.Context, op *xxx_GetServerSelectionOperation) *xxx_GetServerSelectionOperation {
	if op == nil {
		op = &xxx_GetServerSelectionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetServerSelectionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetServerSelectionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetServerSelectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetServerSelectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServerSelectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetServerSelectionOperation structure represents the ServerSelection operation
type xxx_SetServerSelectionOperation struct {
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// value: A value from the ServerSelection (section 2.2.11) enumeration that indicates
	// the type of update server against which to perform search operations initiated by
	// using this interface.
	Value  uamg.ServerSelection `idl:"name:value" json:"value"`
	Return int32                `idl:"name:Return" json:"return"`
}

func (o *xxx_SetServerSelectionOperation) OpNum() int { return 14 }

func (o *xxx_SetServerSelectionOperation) OpName() string {
	return "/IUpdateSearcher/v0/ServerSelection"
}

func (o *xxx_SetServerSelectionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetServerSelectionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// value {in} (1:{v1_enum, alias=ServerSelection}(enum))
	{
		if err := w.WriteEnum(uint32(o.Value)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetServerSelectionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// value {in} (1:{v1_enum, alias=ServerSelection}(enum))
	{
		if err := w.ReadEnum((*uint32)(&o.Value)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetServerSelectionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetServerSelectionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetServerSelectionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetServerSelectionRequest structure represents the ServerSelection operation request
type SetServerSelectionRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// value: A value from the ServerSelection (section 2.2.11) enumeration that indicates
	// the type of update server against which to perform search operations initiated by
	// using this interface.
	//
	// value: A value from the ServerSelection (section 2.2.11) enumeration that indicates
	// the type of update server against which to perform search operations initiated by
	// using this interface.
	Value uamg.ServerSelection `idl:"name:value" json:"value"`
}

func (o *SetServerSelectionRequest) xxx_ToOp(ctx context.Context, op *xxx_SetServerSelectionOperation) *xxx_SetServerSelectionOperation {
	if op == nil {
		op = &xxx_SetServerSelectionOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Value = o.Value
	return op
}

func (o *SetServerSelectionRequest) xxx_FromOp(ctx context.Context, op *xxx_SetServerSelectionOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Value = op.Value
}
func (o *SetServerSelectionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetServerSelectionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetServerSelectionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetServerSelectionResponse structure represents the ServerSelection operation response
type SetServerSelectionResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ServerSelection return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetServerSelectionResponse) xxx_ToOp(ctx context.Context, op *xxx_SetServerSelectionOperation) *xxx_SetServerSelectionOperation {
	if op == nil {
		op = &xxx_SetServerSelectionOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetServerSelectionResponse) xxx_FromOp(ctx context.Context, op *xxx_SetServerSelectionOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetServerSelectionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetServerSelectionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetServerSelectionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EscapeStringOperation structure represents the EscapeString operation
type xxx_EscapeStringOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	Unescaped   *oaut.String   `idl:"name:unescaped" json:"unescaped"`
	ReturnValue *oaut.String   `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_EscapeStringOperation) OpNum() int { return 17 }

func (o *xxx_EscapeStringOperation) OpName() string { return "/IUpdateSearcher/v0/EscapeString" }

func (o *xxx_EscapeStringOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EscapeStringOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// unescaped {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Unescaped != nil {
			_ptr_unescaped := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Unescaped != nil {
					if err := o.Unescaped.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Unescaped, _ptr_unescaped); err != nil {
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

func (o *xxx_EscapeStringOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// unescaped {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_unescaped := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Unescaped == nil {
				o.Unescaped = &oaut.String{}
			}
			if err := o.Unescaped.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_unescaped := func(ptr interface{}) { o.Unescaped = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Unescaped, _s_unescaped, _ptr_unescaped); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EscapeStringOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EscapeStringOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// retval {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ReturnValue != nil {
			_ptr_retval := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnValue != nil {
					if err := o.ReturnValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnValue, _ptr_retval); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EscapeStringOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// retval {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_retval := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnValue == nil {
				o.ReturnValue = &oaut.String{}
			}
			if err := o.ReturnValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_retval := func(ptr interface{}) { o.ReturnValue = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ReturnValue, _s_retval, _ptr_retval); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EscapeStringRequest structure represents the EscapeString operation request
type EscapeStringRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// unescaped: The string to escape.
	Unescaped *oaut.String `idl:"name:unescaped" json:"unescaped"`
}

func (o *EscapeStringRequest) xxx_ToOp(ctx context.Context, op *xxx_EscapeStringOperation) *xxx_EscapeStringOperation {
	if op == nil {
		op = &xxx_EscapeStringOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Unescaped = o.Unescaped
	return op
}

func (o *EscapeStringRequest) xxx_FromOp(ctx context.Context, op *xxx_EscapeStringOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Unescaped = op.Unescaped
}
func (o *EscapeStringRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EscapeStringRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EscapeStringOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EscapeStringResponse structure represents the EscapeString operation response
type EscapeStringResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: Resulting escaped string.
	ReturnValue *oaut.String `idl:"name:retval" json:"return_value"`
	// Return: The EscapeString return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *EscapeStringResponse) xxx_ToOp(ctx context.Context, op *xxx_EscapeStringOperation) *xxx_EscapeStringOperation {
	if op == nil {
		op = &xxx_EscapeStringOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *EscapeStringResponse) xxx_FromOp(ctx context.Context, op *xxx_EscapeStringOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *EscapeStringResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EscapeStringResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EscapeStringOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryHistoryOperation structure represents the QueryHistory operation
type xxx_QueryHistoryOperation struct {
	This        *dcom.ORPCThis                     `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat                     `idl:"name:That" json:"that"`
	StartIndex  int32                              `idl:"name:startIndex" json:"start_index"`
	Count       int32                              `idl:"name:count" json:"count"`
	ReturnValue *uamg.UpdateHistoryEntryCollection `idl:"name:retval" json:"return_value"`
	Return      int32                              `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryHistoryOperation) OpNum() int { return 18 }

func (o *xxx_QueryHistoryOperation) OpName() string { return "/IUpdateSearcher/v0/QueryHistory" }

func (o *xxx_QueryHistoryOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryHistoryOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// startIndex {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.StartIndex); err != nil {
			return err
		}
	}
	// count {in} (1:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryHistoryOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// startIndex {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.StartIndex); err != nil {
			return err
		}
	}
	// count {in} (1:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryHistoryOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryHistoryOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// retval {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUpdateHistoryEntryCollection}(interface))
	{
		if o.ReturnValue != nil {
			_ptr_retval := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnValue != nil {
					if err := o.ReturnValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&uamg.UpdateHistoryEntryCollection{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnValue, _ptr_retval); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryHistoryOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// retval {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=IUpdateHistoryEntryCollection}(interface))
	{
		_ptr_retval := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnValue == nil {
				o.ReturnValue = &uamg.UpdateHistoryEntryCollection{}
			}
			if err := o.ReturnValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_retval := func(ptr interface{}) { o.ReturnValue = *ptr.(**uamg.UpdateHistoryEntryCollection) }
		if err := w.ReadPointer(&o.ReturnValue, _s_retval, _ptr_retval); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// QueryHistoryRequest structure represents the QueryHistory operation request
type QueryHistoryRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// startIndex: The zero-based index of the first history entry to retrieve.
	StartIndex int32 `idl:"name:startIndex" json:"start_index"`
	// count: The number of entries to retrieve.
	Count int32 `idl:"name:count" json:"count"`
}

func (o *QueryHistoryRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryHistoryOperation) *xxx_QueryHistoryOperation {
	if op == nil {
		op = &xxx_QueryHistoryOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.StartIndex = o.StartIndex
	op.Count = o.Count
	return op
}

func (o *QueryHistoryRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryHistoryOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.StartIndex = op.StartIndex
	o.Count = op.Count
}
func (o *QueryHistoryRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryHistoryRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryHistoryOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryHistoryResponse structure represents the QueryHistory operation response
type QueryHistoryResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: An IUpdateHistoryEntryCollection containing the history entries requested.
	// If fewer entries are available than requested in the count parameter, only the entries
	// available are retrieved.
	ReturnValue *uamg.UpdateHistoryEntryCollection `idl:"name:retval" json:"return_value"`
	// Return: The QueryHistory return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *QueryHistoryResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryHistoryOperation) *xxx_QueryHistoryOperation {
	if op == nil {
		op = &xxx_QueryHistoryOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *QueryHistoryResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryHistoryOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *QueryHistoryResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryHistoryResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryHistoryOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SearchOperation structure represents the Search operation
type xxx_SearchOperation struct {
	This        *dcom.ORPCThis     `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat     `idl:"name:That" json:"that"`
	Criteria    *oaut.String       `idl:"name:criteria" json:"criteria"`
	ReturnValue *uamg.SearchResult `idl:"name:retval" json:"return_value"`
	Return      int32              `idl:"name:Return" json:"return"`
}

func (o *xxx_SearchOperation) OpNum() int { return 19 }

func (o *xxx_SearchOperation) OpName() string { return "/IUpdateSearcher/v0/Search" }

func (o *xxx_SearchOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SearchOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// criteria {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Criteria != nil {
			_ptr_criteria := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Criteria != nil {
					if err := o.Criteria.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Criteria, _ptr_criteria); err != nil {
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

func (o *xxx_SearchOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// criteria {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_criteria := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Criteria == nil {
				o.Criteria = &oaut.String{}
			}
			if err := o.Criteria.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_criteria := func(ptr interface{}) { o.Criteria = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Criteria, _s_criteria, _ptr_criteria); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SearchOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SearchOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// retval {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=ISearchResult}(interface))
	{
		if o.ReturnValue != nil {
			_ptr_retval := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnValue != nil {
					if err := o.ReturnValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&uamg.SearchResult{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnValue, _ptr_retval); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SearchOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// retval {out, retval} (1:{pointer=ref}*(2)*(1))(2:{alias=ISearchResult}(interface))
	{
		_ptr_retval := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnValue == nil {
				o.ReturnValue = &uamg.SearchResult{}
			}
			if err := o.ReturnValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_retval := func(ptr interface{}) { o.ReturnValue = *ptr.(**uamg.SearchResult) }
		if err := w.ReadPointer(&o.ReturnValue, _s_retval, _ptr_retval); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SearchRequest structure represents the Search operation request
type SearchRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// criteria: A string specifying criteria that updates are required to meet in order
	// to be returned in the search results.
	//
	// The server MUST support criteria specified by the following ABNF grammar.
	//
	// criteria = andgroup *[1,*WSP "or" 1,*WSP andgroup] /  andgroup = andgroup *[1,*WSP
	// "and" 1,*WSP andgroup] / "(" *WSP andgroup *WSP ")" / criterion  criterion = propertyname
	// *WSP operator *WSP value / "(" *WSP criterion *WSP ")"  propertyname = "Type" / "DeploymentAction"
	// / "IsAssigned" / "AutoSelectOnWebSites" / "BrowseOnly" / "UpdateID" / "RevisionNumber"
	// / "CategoryIDs" / "IsInstalled" / "IsHidden" /  "IsPresent" / "RebootRequired"
	// operator = "=" / "!=" / "contains"  value = string / integer  string = "'" *CHAR
	// "'"  integer = ["-"] 1*DIGIT
	//
	// The CHAR and DIGIT production rules are specified in [RFC5234] section B.1.
	//
	// Integer values MUST be specified in base 10. String values MUST be escaped (using
	// the IUpdateSearcher::EscapeString (opnum 18) method) before being enclosed in single
	// quotation marks, as specified by the preceding grammar. String comparisons MUST be
	// case-insensitive.
	//
	// The server MUST support the following update properties with the given operators
	// and value types.
	//
	//	+----------------------+---------------------+------------------------+----------------------------------------------------------------------------------+
	//	|   UPDATE PROPERTY    |      OPERATORS      |         VALUE          |                                                                                  |
	//	|         NAME         |      SUPPORTED      |          TYPE          |                                   DESCRIPTION                                    |
	//	|                      |                     |                        |                                                                                  |
	//	+----------------------+---------------------+------------------------+----------------------------------------------------------------------------------+
	//	+----------------------+---------------------+------------------------+----------------------------------------------------------------------------------+
	//	| Type                 | =, !=               | String                 | Compares the update's type to that given.                                        |
	//	+----------------------+---------------------+------------------------+----------------------------------------------------------------------------------+
	//	| DeploymentAction     | =                   | String                 | Finds updates with the given deployment action. If this criterion is not         |
	//	|                      |                     |                        | specified, the search operation MUST return only updates with a DeploymentAction |
	//	|                      |                     |                        | of "Installation".                                                               |
	//	+----------------------+---------------------+------------------------+----------------------------------------------------------------------------------+
	//	| IsAssigned           | =                   | Integer MUST be 0 or 1 | The value is treated as a Boolean. The value determines whether the search       |
	//	|                      |                     |                        | operation finds updates that are intended for processing by an automatic update  |
	//	|                      |                     |                        | agent.                                                                           |
	//	+----------------------+---------------------+------------------------+----------------------------------------------------------------------------------+
	//	| AutoSelectOnWebsites | =                   | Integer MUST be 0 or 1 | The value is treated as a Boolean. The value determines whether the search       |
	//	|                      |                     |                        | operation finds updates that are intended to be automatically selected from web  |
	//	|                      |                     |                        | user interfaces.                                                                 |
	//	+----------------------+---------------------+------------------------+----------------------------------------------------------------------------------+
	//	| BrowseOnly           | =                   | Integer MUST be 0 or 1 | The value is treated as a Boolean. The value determines whether the search       |
	//	|                      |                     |                        | operation finds updates that are not intended for processing by an automatic     |
	//	|                      |                     |                        | update agent.                                                                    |
	//	+----------------------+---------------------+------------------------+----------------------------------------------------------------------------------+
	//	| UpdateID             | =, !=               | String                 | Compares the update's ID to that given.                                          |
	//	+----------------------+---------------------+------------------------+----------------------------------------------------------------------------------+
	//	| RevisionNumber       | =                   | Integer                | Finds updates with the given revision number.                                    |
	//	+----------------------+---------------------+------------------------+----------------------------------------------------------------------------------+
	//	| CategoryIDs          | Contains            | String                 | Finds updates belonging to the given category.                                   |
	//	+----------------------+---------------------+------------------------+----------------------------------------------------------------------------------+
	//	| IsInstalled          | =                   | Integer MUST be 0 or 1 | The value is treated as a Boolean. The value determines whether the search       |
	//	|                      |                     |                        | operation finds updates that are installed.                                      |
	//	+----------------------+---------------------+------------------------+----------------------------------------------------------------------------------+
	//	| IsHidden             | =                   | Integer MUST be 0 or 1 | The value is treated as a Boolean. The value determines whether the search       |
	//	|                      |                     |                        | operation finds updates that have been hidden by a user.                         |
	//	+----------------------+---------------------+------------------------+----------------------------------------------------------------------------------+
	//	| IsPresent            | =                   | Integer MUST be 0 or 1 | The value is treated as a Boolean. The value determines whether the search       |
	//	|                      |                     |                        | operation finds updates that are present, that is, updates that have been        |
	//	|                      |                     |                        | installed for one or more products.                                              |
	//	+----------------------+---------------------+------------------------+----------------------------------------------------------------------------------+
	//	| RebootRequired       | =                   | Integer MUST be 0 or 1 | The value is treated as a Boolean. The value determines whether the search       |
	//	|                      |                     |                        | operation finds updates that currently require a reboot to complete installation |
	//	|                      |                     |                        | or uninstallation.                                                               |
	//	+----------------------+---------------------+------------------------+----------------------------------------------------------------------------------+
	Criteria *oaut.String `idl:"name:criteria" json:"criteria"`
}

func (o *SearchRequest) xxx_ToOp(ctx context.Context, op *xxx_SearchOperation) *xxx_SearchOperation {
	if op == nil {
		op = &xxx_SearchOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Criteria = o.Criteria
	return op
}

func (o *SearchRequest) xxx_FromOp(ctx context.Context, op *xxx_SearchOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Criteria = op.Criteria
}
func (o *SearchRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SearchRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SearchOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SearchResponse structure represents the Search operation response
type SearchResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: An ISearchResult instance with the results of the search.
	ReturnValue *uamg.SearchResult `idl:"name:retval" json:"return_value"`
	// Return: The Search return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SearchResponse) xxx_ToOp(ctx context.Context, op *xxx_SearchOperation) *xxx_SearchOperation {
	if op == nil {
		op = &xxx_SearchOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *SearchResponse) xxx_FromOp(ctx context.Context, op *xxx_SearchOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *SearchResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SearchResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SearchOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetOnlineOperation structure represents the Online operation
type xxx_GetOnlineOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue int16          `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetOnlineOperation) OpNum() int { return 20 }

func (o *xxx_GetOnlineOperation) OpName() string { return "/IUpdateSearcher/v0/Online" }

func (o *xxx_GetOnlineOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOnlineOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOnlineOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOnlineOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOnlineOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.ReturnValue); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetOnlineOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.ReturnValue); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetOnlineRequest structure represents the Online operation request
type GetOnlineRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetOnlineRequest) xxx_ToOp(ctx context.Context, op *xxx_GetOnlineOperation) *xxx_GetOnlineOperation {
	if op == nil {
		op = &xxx_GetOnlineOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetOnlineRequest) xxx_FromOp(ctx context.Context, op *xxx_GetOnlineOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetOnlineRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetOnlineRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOnlineOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetOnlineResponse structure represents the Online operation response
type GetOnlineResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: MUST be set either to VARIANT_TRUE if the search operation can contact the
	// server or to VARIANT_FALSE if it uses local data only.
	ReturnValue int16 `idl:"name:retval" json:"return_value"`
	// Return: The Online return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetOnlineResponse) xxx_ToOp(ctx context.Context, op *xxx_GetOnlineOperation) *xxx_GetOnlineOperation {
	if op == nil {
		op = &xxx_GetOnlineOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetOnlineResponse) xxx_FromOp(ctx context.Context, op *xxx_GetOnlineOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetOnlineResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetOnlineResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetOnlineOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetOnlineOperation structure represents the Online operation
type xxx_SetOnlineOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Value  int16          `idl:"name:value" json:"value"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetOnlineOperation) OpNum() int { return 21 }

func (o *xxx_SetOnlineOperation) OpName() string { return "/IUpdateSearcher/v0/Online" }

func (o *xxx_SetOnlineOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOnlineOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// value {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.WriteData(o.Value); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOnlineOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// value {in} (1:{alias=VARIANT_BOOL}(int16))
	{
		if err := w.ReadData(&o.Value); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOnlineOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOnlineOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetOnlineOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetOnlineRequest structure represents the Online operation request
type SetOnlineRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// value: MUST be set either to VARIANT_TRUE if the search operation can contact the
	// server or to VARIANT_FALSE if it uses local data only.
	Value int16 `idl:"name:value" json:"value"`
}

func (o *SetOnlineRequest) xxx_ToOp(ctx context.Context, op *xxx_SetOnlineOperation) *xxx_SetOnlineOperation {
	if op == nil {
		op = &xxx_SetOnlineOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Value = o.Value
	return op
}

func (o *SetOnlineRequest) xxx_FromOp(ctx context.Context, op *xxx_SetOnlineOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Value = op.Value
}
func (o *SetOnlineRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetOnlineRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOnlineOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetOnlineResponse structure represents the Online operation response
type SetOnlineResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The Online return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetOnlineResponse) xxx_ToOp(ctx context.Context, op *xxx_SetOnlineOperation) *xxx_SetOnlineOperation {
	if op == nil {
		op = &xxx_SetOnlineOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetOnlineResponse) xxx_FromOp(ctx context.Context, op *xxx_SetOnlineOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetOnlineResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetOnlineResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetOnlineOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTotalHistoryCountOperation structure represents the GetTotalHistoryCount operation
type xxx_GetTotalHistoryCountOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue int32          `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTotalHistoryCountOperation) OpNum() int { return 22 }

func (o *xxx_GetTotalHistoryCountOperation) OpName() string {
	return "/IUpdateSearcher/v0/GetTotalHistoryCount"
}

func (o *xxx_GetTotalHistoryCountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTotalHistoryCountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTotalHistoryCountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTotalHistoryCountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTotalHistoryCountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.WriteData(o.ReturnValue); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTotalHistoryCountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// retval {out, retval} (1:{pointer=ref}*(1))(2:{alias=LONG}(int32))
	{
		if err := w.ReadData(&o.ReturnValue); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetTotalHistoryCountRequest structure represents the GetTotalHistoryCount operation request
type GetTotalHistoryCountRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetTotalHistoryCountRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTotalHistoryCountOperation) *xxx_GetTotalHistoryCountOperation {
	if op == nil {
		op = &xxx_GetTotalHistoryCountOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetTotalHistoryCountRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTotalHistoryCountOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetTotalHistoryCountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetTotalHistoryCountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTotalHistoryCountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTotalHistoryCountResponse structure represents the GetTotalHistoryCount operation response
type GetTotalHistoryCountResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: The total number of history entries stored on the computer.
	ReturnValue int32 `idl:"name:retval" json:"return_value"`
	// Return: The GetTotalHistoryCount return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetTotalHistoryCountResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTotalHistoryCountOperation) *xxx_GetTotalHistoryCountOperation {
	if op == nil {
		op = &xxx_GetTotalHistoryCountOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetTotalHistoryCountResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTotalHistoryCountOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetTotalHistoryCountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetTotalHistoryCountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTotalHistoryCountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetServiceIDOperation structure represents the ServiceID operation
type xxx_GetServiceIDOperation struct {
	This        *dcom.ORPCThis `idl:"name:This" json:"this"`
	That        *dcom.ORPCThat `idl:"name:That" json:"that"`
	ReturnValue *oaut.String   `idl:"name:retval" json:"return_value"`
	Return      int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_GetServiceIDOperation) OpNum() int { return 23 }

func (o *xxx_GetServiceIDOperation) OpName() string { return "/IUpdateSearcher/v0/ServiceID" }

func (o *xxx_GetServiceIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServiceIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServiceIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServiceIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServiceIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// retval {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.ReturnValue != nil {
			_ptr_retval := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ReturnValue != nil {
					if err := o.ReturnValue.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ReturnValue, _ptr_retval); err != nil {
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
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServiceIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// retval {out, retval} (1:{pointer=ref}*(2))(2:{pointer=unique, alias=BSTR}*(1))(3:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_retval := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ReturnValue == nil {
				o.ReturnValue = &oaut.String{}
			}
			if err := o.ReturnValue.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_retval := func(ptr interface{}) { o.ReturnValue = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.ReturnValue, _s_retval, _ptr_retval); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetServiceIDRequest structure represents the ServiceID operation request
type GetServiceIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *GetServiceIDRequest) xxx_ToOp(ctx context.Context, op *xxx_GetServiceIDOperation) *xxx_GetServiceIDOperation {
	if op == nil {
		op = &xxx_GetServiceIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	return op
}

func (o *GetServiceIDRequest) xxx_FromOp(ctx context.Context, op *xxx_GetServiceIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *GetServiceIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetServiceIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServiceIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetServiceIDResponse structure represents the ServiceID operation response
type GetServiceIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// retval: A string uniquely identifying the service.
	//
	// retval: A string containing the unique identifier of the update service that provided
	// the update for which the operation was performed.
	//
	// retval: A string containing the unique identifier of the update server to use for
	// search.
	//
	// retval: A string identifying the service.
	ReturnValue *oaut.String `idl:"name:retval" json:"return_value"`
	// Return: The ServiceID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetServiceIDResponse) xxx_ToOp(ctx context.Context, op *xxx_GetServiceIDOperation) *xxx_GetServiceIDOperation {
	if op == nil {
		op = &xxx_GetServiceIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.ReturnValue = o.ReturnValue
	op.Return = o.Return
	return op
}

func (o *GetServiceIDResponse) xxx_FromOp(ctx context.Context, op *xxx_GetServiceIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReturnValue = op.ReturnValue
	o.Return = op.Return
}
func (o *GetServiceIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetServiceIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServiceIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetServiceIDOperation structure represents the ServiceID operation
type xxx_SetServiceIDOperation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Value  *oaut.String   `idl:"name:value" json:"value"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_SetServiceIDOperation) OpNum() int { return 24 }

func (o *xxx_SetServiceIDOperation) OpName() string { return "/IUpdateSearcher/v0/ServiceID" }

func (o *xxx_SetServiceIDOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetServiceIDOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// value {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		if o.Value != nil {
			_ptr_value := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Value != nil {
					if err := o.Value.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&oaut.String{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Value, _ptr_value); err != nil {
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

func (o *xxx_SetServiceIDOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// value {in} (1:{pointer=unique, alias=BSTR}*(1))(2:{pointer=unique, alias=_BSTR, names=FLAGGED_WORD_BLOB}(struct))
	{
		_ptr_value := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Value == nil {
				o.Value = &oaut.String{}
			}
			if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_value := func(ptr interface{}) { o.Value = *ptr.(**oaut.String) }
		if err := w.ReadPointer(&o.Value, _s_value, _ptr_value); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetServiceIDOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetServiceIDOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetServiceIDOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetServiceIDRequest structure represents the ServiceID operation request
type SetServiceIDRequest struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// value: A string containing the unique identifier of the update server to use for
	// the search.
	Value *oaut.String `idl:"name:value" json:"value"`
}

func (o *SetServiceIDRequest) xxx_ToOp(ctx context.Context, op *xxx_SetServiceIDOperation) *xxx_SetServiceIDOperation {
	if op == nil {
		op = &xxx_SetServiceIDOperation{}
	}
	if o == nil {
		return op
	}
	op.This = o.This
	op.Value = o.Value
	return op
}

func (o *SetServiceIDRequest) xxx_FromOp(ctx context.Context, op *xxx_SetServiceIDOperation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.Value = op.Value
}
func (o *SetServiceIDRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetServiceIDRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetServiceIDOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetServiceIDResponse structure represents the ServiceID operation response
type SetServiceIDResponse struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The ServiceID return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *SetServiceIDResponse) xxx_ToOp(ctx context.Context, op *xxx_SetServiceIDOperation) *xxx_SetServiceIDOperation {
	if op == nil {
		op = &xxx_SetServiceIDOperation{}
	}
	if o == nil {
		return op
	}
	op.That = o.That
	op.Return = o.Return
	return op
}

func (o *SetServiceIDResponse) xxx_FromOp(ctx context.Context, op *xxx_SetServiceIDOperation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *SetServiceIDResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetServiceIDResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetServiceIDOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
