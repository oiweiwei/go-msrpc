package iwrmaccounting

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

// IWRMAccounting server interface.
type AccountingServer interface {

	// IDispatch base class.
	idispatch.DispatchServer

	// The CreateAccountingDb method creates the database for accounting data.<40>
	//
	// Return Values: This method returns 0x00000000 for success, or a negative HRESULT
	// value (shown in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------------+------------------------------------------+
	//	|                RETURN                |                                          |
	//	|              VALUE/CODE              |               DESCRIPTION                |
	//	|                                      |                                          |
	//	+--------------------------------------+------------------------------------------+
	//	+--------------------------------------+------------------------------------------+
	//	| 0x00000000 S_OK                      | Operation is successful.                 |
	//	+--------------------------------------+------------------------------------------+
	//	| 0xC1FF01F7 WRM_ERR_ACCOUNTING_FAILED | WSRM encountered an error in accounting. |
	//	+--------------------------------------+------------------------------------------+
	//
	// Additional IWRMAccounting interface methods are specified in section 3.2.4.3.
	CreateAccountingDB(context.Context, *CreateAccountingDBRequest) (*CreateAccountingDBResponse, error)

	// The GetAccountingMetadata method retrieves accounting metadata, which includes column
	// names, types, and other attributes of the accounting tables.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (int the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Operation successful. |
	//	+-------------------+-----------------------+
	//
	// Additional IWRMAccounting interface methods are specified in section 3.2.4.3.
	GetAccountingMetadata(context.Context, *GetAccountingMetadataRequest) (*GetAccountingMetadataResponse, error)

	// The ExecuteAccountingQuery method executes an accounting query.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                         RETURN                         |                                                                                  |
	//	|                       VALUE/CODE                       |                                   DESCRIPTION                                    |
	//	|                                                        |                                                                                  |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                                        | Operation successful.                                                            |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                                | One or more arguments are invalid.                                               |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF01FA WRM_ERR_WYUKON_NOT_CONNECTABLE              | Cannot establish a connection to the accounting database.<41>                    |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF01FE WRM_ERR_JET_INVALID_COLUMN_NAME             | The query has been canceled. One or more of the column names specified in the    |
	//	|                                                        | accounting query are invalid.                                                    |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0200 WRM_ERR_JET_PIVOTABLE_COLUMN_NOT_GROUPED_BY | One or more SQL SELECT columns cannot be selected because of the current SQL     |
	//	|                                                        | GROUP BY settings. These columns MUST be grouped to be selected.                 |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0201 WRM_ERR_JET_INVALID_GROUP_BY_COL            | One or more columns specified for SQL GROUP BY is either invalid or cannot be    |
	//	|                                                        | grouped on.                                                                      |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0203 WRM_ERR_JET_SERVER_TOO_BUSY                 | The server can service only one accounting request at a time.<42>                |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0207 WRM_ERR_JET_SERVICE_BEING_SHUT_DOWN         | The query has been aborted since the management service is being shut down.      |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The error WRM_ERR_JET_PIVOTABLE_COLUMN_NOT_GROUPED_BY is returned in cases where
	// a column with the IsVisible attribute set to FALSE is included in the SQL SELECT
	// column while there are some columns in the group column collection. The following
	// sample AccountingQueryCondition XML (section 2.2.5.5) SHOULD return this error:
	//
	// Additional IWRMAccounting interface methods are specified in section 3.2.4.3.
	ExecuteAccountingQuery(context.Context, *ExecuteAccountingQueryRequest) (*ExecuteAccountingQueryResponse, error)

	// The GetRawAccountingData method returns raw accounting data from the accounting database
	// (section 3.2.1.2).
	//
	// Return Values: This method returns 0x00000000 for success, or a negative HRESULT
	// value (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                                  |
	//	|                VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                           | Operation successful.                                                            |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                   | One or more arguments are invalid.                                               |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF01F7 WRM_ERR_ACCOUNTING_FAILED      | WSRM encountered an error in accounting.                                         |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF01FA WRM_ERR_WYUKON_NOT_CONNECTABLE | Cannot establish a connection to the accounting database due to an error other   |
	//	|                                           | than one of the errors of the WRM_ERR_WYUKON_CORRUPTED return value.             |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF01FB WRM_ERR_WYUKON_CORRUPTED       | Cannot establish a connection to the accounting database; either the database is |
	//	|                                           | in single-user mode and already connected, or it is in an invalid or corrupted   |
	//	|                                           | state.                                                                           |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF01FC WRM_ERR_WYUKON_NOT_INSTALLED   | The accounting database is not installed on the specified server.                |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0203 WRM_ERR_JET_SERVER_TOO_BUSY    | The server can service only one accounting request at a time.<44>                |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMAccounting interface methods are specified in section 3.2.4.3.
	GetRawAccountingData(context.Context, *GetRawAccountingDataRequest) (*GetRawAccountingDataResponse, error)

	// The GetNextAccountingDataBatch method gets the next batch of data in a previously
	// initiated query to the accounting database.<45>
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                                  |
	//	|                VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                           | Operation successful.                                                            |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x81FF0217 WRM_WRN_WSRM_INCOMPLETE_FETCH  | Data was fetched incompletely; the connection to the database might have         |
	//	|                                           | terminated.                                                                      |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF01FA WRM_ERR_WYUKON_NOT_CONNECTABLE | Cannot establish a connection to the accounting database.                        |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0203 WRM_ERR_JET_SERVER_TOO_BUSY    | The server can service only one accounting request at a time.<46>                |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The method GetNextAccountingDataBatch returns data from the accounting database if
	// all the data was not retrieved by previous calls to this method or either of the
	// methods ExecuteAccountingQuery (section 3.2.4.3.3) or GetRawAccountingData (section
	// 3.2.4.3.4). The availability of additional database data is indicated by the value
	// returned in the pbIsThereMoreData parameter of each of these methods.
	//
	// If ExecuteAccountingQuery or GetRawAccountingData had returned indicating no more
	// accounting data to be retrieved, and still GetNextAccountingDataBatch is called,
	// pbstrResult is returned as NULL.
	//
	// Additional IWRMAccounting interface methods are specified in section 3.2.4.3.
	GetNextAccountingDataBatch(context.Context, *GetNextAccountingDataBatchRequest) (*GetNextAccountingDataBatchResponse, error)

	// The DeleteAccountingData method deletes accounting data within a specified time period
	// from the accounting database (section 3.2.1.2). If there is no accounting data present
	// between the specified dates, the functions returns SUCCESS while no accounting data
	// is deleted.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                                  |
	//	|                VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                           | Operation successful.                                                            |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                   | One or more arguments are invalid.<48>                                           |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF01FA WRM_ERR_WYUKON_NOT_CONNECTABLE | Cannot establish a connection to the accounting database.                        |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF01FB WRM_ERR_WYUKON_CORRUPTED       | Cannot establish a connection to the accounting database; either the database    |
	//	|                                           | is in single user mode and already connected or it is in an invalid or corrupted |
	//	|                                           | state.                                                                           |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0203 WRM_ERR_JET_SERVER_TOO_BUSY    | The server can service only one accounting request at a time.                    |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Additional IWRMAccounting interface methods are specified in section 3.2.4.3.
	DeleteAccountingData(context.Context, *DeleteAccountingDataRequest) (*DeleteAccountingDataResponse, error)

	// The DefragmentDB method is not implemented. It MUST return a success code.
	//
	// This method has no parameters.
	//
	// Return Values: This method returns 0x00000000 for success.
	//
	//	+-------------------+-----------------------+
	//	|      RETURN       |                       |
	//	|    VALUE/CODE     |      DESCRIPTION      |
	//	|                   |                       |
	//	+-------------------+-----------------------+
	//	+-------------------+-----------------------+
	//	| 0x00000000 S_OK   | Operation successful. |
	//	+-------------------+-----------------------+
	//
	// Additional IWRMAccounting interface methods are specified in section 3.2.4.3.
	DefragmentDB(context.Context, *DefragmentDBRequest) (*DefragmentDBResponse, error)

	// The CancelAccountingQuery method cancels a previously-initiated query to the accounting
	// database.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+-------------------------------------------+---------------------------------------------------------------+
	//	|                  RETURN                   |                                                               |
	//	|                VALUE/CODE                 |                          DESCRIPTION                          |
	//	|                                           |                                                               |
	//	+-------------------------------------------+---------------------------------------------------------------+
	//	+-------------------------------------------+---------------------------------------------------------------+
	//	| 0x00000000 S_OK                           | Operation successful.                                         |
	//	+-------------------------------------------+---------------------------------------------------------------+
	//	| 0xC1FF01FA WRM_ERR_WYUKON_NOT_CONNECTABLE | Cannot establish a connection to the accounting database.<49> |
	//	+-------------------------------------------+---------------------------------------------------------------+
	//
	// The method CancelAccountingQuery cancels a query to the accounting database after
	// previous calls to either of the methods ExecuteAccountingQuery (section 3.2.4.3.3)
	// or GetRawAccountingData (section 3.2.4.3.4) and before a call to the method GetNextAccountingDataBatch
	// (section 3.2.4.3.5). The availability of additional database data is indicated by
	// the value returned in the pbIsThereMoreData parameter of each method.
	//
	// Additional IWRMAccounting interface methods are specified in section 3.2.4.3.
	CancelAccountingQuery(context.Context, *CancelAccountingQueryRequest) (*CancelAccountingQueryResponse, error)

	// The RegisterAccountingClient method registers an accounting client for remote accounting
	// on an accounting server. A default accounting database SHOULD<50> be defined.
	//
	// Note  This method is expected to be called remotely by the WSRM management service
	// that is acting as an accounting client. A client SHOULD NOT call this method.
	//
	// Return Values: This method returns 0x00000000 for success, or a negative HRESULT
	// value (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+----------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                          RETURN                          |                                                                                  |
	//	|                        VALUE/CODE                        |                                   DESCRIPTION                                    |
	//	|                                                          |                                                                                  |
	//	+----------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                                          | Operation successful.                                                            |
	//	+----------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                                  | One or more arguments are invalid.                                               |
	//	+----------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0209 WRM_ERR_ACC_DISABLED_FOR_REMOTE_CLIENT        | WSRM encountered an error in accounting.                                         |
	//	+----------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF020C WRM_ERR_REMOTE_SERVICE_NOT_SETUP_FOR_REMOTING | Connection to the remote server could not be established. The server is not set  |
	//	|                                                          | up for remote accounting. This error is returned when an accounting client is    |
	//	|                                                          | trying to register itself on an accounting server machine which itself is an     |
	//	|                                                          | accounting client.<51>                                                           |
	//	+----------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0212 WRM_ERR_INVALID_OPERATION                     | The operation is invalid. This error is returned when a WSRM management service  |
	//	|                                                          | acting as an accounting client tries to register itself twice, as when this      |
	//	|                                                          | method is called with the same DCOM interface instance.                          |
	//	+----------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0216 WRM_ERR_DBSERVER_CANNOT_BE_REMOTE             | Data cannot be logged on the remote system. This error is returned when an       |
	//	|                                                          | accounting client calls this method on the accounting server but accounting is   |
	//	|                                                          | not yet initialized.<52>                                                         |
	//	+----------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// It is possible for multiple WSRM servers to act as accounting clients and have a
	// common accounting database server.
	//
	// Note  Remote accounting is not supported in a workgroup environment.
	//
	// Additional IWRMAccounting interface methods are specified in section 3.2.4.3.
	RegisterAccountingClient(context.Context, *RegisterAccountingClientRequest) (*RegisterAccountingClientResponse, error)

	// The DumpAccountingData method dumps accounting data from a remote server acting as
	// an accounting client to the server currently acting as its accounting database server.
	// The time interval of dumping data SHOULD be set by the client by using the SetConfig
	// method call of the management service running on the accounting client.
	//
	// Note  This method is expected to be called remotely by the WSRM management service
	// that is acting as an accounting client. A client SHOULD NOT call this method.
	//
	// Return Values: This method returns 0x00000000 for success or a negative HRESULT value
	// (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                      RETURN                       |                                                                                  |
	//	|                    VALUE/CODE                     |                                   DESCRIPTION                                    |
	//	|                                                   |                                                                                  |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                                   | Operation successful.                                                            |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                           | One or more arguments are invalid.                                               |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0209 WRM_ERR_ACC_DISABLED_FOR_REMOTE_CLIENT | Accounting is disabled on the remote accounting server.                          |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF020A WRM_ERR_ACC_CLIENT_CONN_TERMINATED     | Connection to remote server failed. The remote connection has been terminated    |
	//	|                                                   | by the accounting server because the accounting functionality of the accounting  |
	//	|                                                   | client was disabled.                                                             |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF020E WRM_ERR_ACC_NO_CONNECTION              | The operation failed. There is no existing connection to the remote database     |
	//	|                                                   | server.                                                                          |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// It is possible for multiple WSRM servers to have a common accounting database server.
	//
	// Note  Remote accounting is not supported in a workgroup environment.
	//
	// Additional IWRMAccounting interface methods are specified in section 3.2.4.3.
	DumpAccountingData(context.Context, *DumpAccountingDataRequest) (*DumpAccountingDataResponse, error)

	// The GetAccountingClients method gets the names of all servers that are acting as
	// accounting clients on the current server.
	//
	// Return Values: This method returns 0x00000000 for success, or a negative HRESULT
	// value (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                   RETURN                   |                                                                                  |
	//	|                 VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                            |                                                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                            | Operation successful.                                                            |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0210 WRM_ERR_CLIENTSTATUS_ACC_OFF    | Accounting functionality is currently turned off on the accounting server. There |
	//	|                                            | are no clients logging data to this machine.                                     |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0211 WRM_ERR_CLIENTSTATUS_ACC_REMOTE | Accounting is currently being done on a remote machine. There are no clients     |
	//	|                                            | logging data to this machine.                                                    |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The GetAccountingClients method can be used to find the names of WSRM servers, acting
	// as accounting clients, for which the current server is acting as an accounting server.
	//
	// Note  Remote accounting is not supported in a workgroup environment.
	//
	// Additional IWRMAccounting interface methods are specified in section 3.2.4.3.
	GetAccountingClients(context.Context, *GetAccountingClientsRequest) (*GetAccountingClientsResponse, error)

	// The SetAccountingClientStatus method sets the status of accounting functionality
	// of servers that are in the remote accounting client role. Status of accounting functionality
	// is controlled by the Enabled attribute in bstrClientIds XML. If the Enabled attribute
	// in the bstrClientIds XML is specified as "true", the accounting functionality for
	// the WSRM server, whose name is specified as the value of the respective AccountingClient
	// node, is enabled; otherwise, if it is specified as "false", the accounting functionality
	// is disabled. The DumpAccountingData method does not store accounting data in the
	// database if the status of the accounting functionality of the respective server is
	// disabled using the SetAccountingClientStatus method.<53>
	//
	// Return Values: This method returns 0x00000000 for success, or a negative HRESULT
	// value (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                     RETURN                     |                                                                                  |
	//	|                   VALUE/CODE                   |                                   DESCRIPTION                                    |
	//	|                                                |                                                                                  |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                                | Operation successful.                                                            |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG                        | One or more arguments are invalid.                                               |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF020F WRM_ERR_SETCLIENTSTATUS_XML_INVALID | The operation failed. XML data provided is invalid.                              |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0210 WRM_ERR_CLIENTSTATUS_ACC_OFF        | Accounting is currently turned off. There are no clients logging data to this    |
	//	|                                                | machine.                                                                         |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0211 WRM_ERR_CLIENTSTATUS_ACC_REMOTE     | Accounting is currently being done on a remote machine. There are no clients     |
	//	|                                                | logging data to this machine.                                                    |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Note  Remote accounting is not supported in a workgroup environment.
	//
	// Additional IWRMAccounting interface methods are specified in section 3.2.4.3.
	SetAccountingClientStatus(context.Context, *SetAccountingClientStatusRequest) (*SetAccountingClientStatusResponse, error)

	// The CheckAccountingConnection method determines the status of the connection to the
	// accounting server.
	//
	// Note  This method is expected to be called remotely by the WSRM management service.
	// In a remote accounting scenario, this method is called by the accounting client's
	// WSRM service to check the connection status. A WSRM client SHOULD NOT call this method.
	//
	// This method has no parameters.
	//
	// Return Values: This method returns 0x00000000 for success, or a negative HRESULT
	// value (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                      RETURN                       |                                                                                  |
	//	|                    VALUE/CODE                     |                                   DESCRIPTION                                    |
	//	|                                                   |                                                                                  |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                                   | Operation successful.                                                            |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0209 WRM_ERR_ACC_DISABLED_FOR_REMOTE_CLIENT | Accounting is disabled on the remote accounting server.                          |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF020A WRM_ERR_ACC_CLIENT_CONN_TERMINATED     | Connection to remote accounting server failed. The remote connection has been    |
	//	|                                                   | terminated by the accounting server because the accounting functionality of the  |
	//	|                                                   | accounting client was disabled.<54>                                              |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF020E WRM_ERR_ACC_NO_CONNECTION              | The operation failed. There is no existing connection to the remote database     |
	//	|                                                   | server.                                                                          |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// It is possible for multiple WSRM servers to have a common accounting database server.
	// <55>
	//
	// Note  Remote accounting is not supported in a workgroup environment.
	//
	// Additional IWRMAccounting interface methods are specified in section 3.2.4.3.
	CheckAccountingConnection(context.Context, *CheckAccountingConnectionRequest) (*CheckAccountingConnectionResponse, error)

	// The SetClientPermissions method adds or removes a specified client for remote accounting
	// on a server and makes the required changes for DCOM permissions; that is, it adds
	// the machine account of the accounting client to the Distributed COM Users group (see
	// [MS-SAMR] section 3.1.4.2) of the accounting server.
	//
	// Return Values: This method returns 0x00000000 for success, or a negative HRESULT
	// value (in the following table or in [MS-ERREF] section 2.1.1) if an error occurs.
	//
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                      RETURN                      |                                                                                  |
	//	|                    VALUE/CODE                    |                                   DESCRIPTION                                    |
	//	|                                                  |                                                                                  |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK                                  | Operation successful.                                                            |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0213 WRM_ERR_INVALID_DBSERVER_NAME         | Could not establish a connection to the database server. Either the server name  |
	//	|                                                  | is invalid or the machine executing this method does not have permission to      |
	//	|                                                  | perform an account lookup.                                                       |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0xC1FF0214 WRM_ERR_PERMISSION_GRANT_UNSUCCESSFUL | Could not grant permissions to the current machine to log data remotely. This    |
	//	|                                                  | error is returned when a server machine denies access to a requesting machine.   |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// It is possible for multiple WSRM servers to have a common accounting server.
	//
	// Note  Remote accounting is not supported in a workgroup environment.
	//
	// Additional IWRMAccounting interface methods are specified in section 3.2.4.3.
	SetClientPermissions(context.Context, *SetClientPermissionsRequest) (*SetClientPermissionsResponse, error)
}

func RegisterAccountingServer(conn dcerpc.Conn, o AccountingServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAccountingServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AccountingSyntaxV0_0))...)
}

func NewAccountingServerHandle(o AccountingServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AccountingServerHandle(ctx, o, opNum, r)
	}
}

func AccountingServerHandle(ctx context.Context, o AccountingServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 7 {
		// IDispatch base method.
		return idispatch.DispatchServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 7: // CreateAccountingDb
		op := &xxx_CreateAccountingDBOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateAccountingDBRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateAccountingDB(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // GetAccountingMetadata
		op := &xxx_GetAccountingMetadataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAccountingMetadataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAccountingMetadata(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // ExecuteAccountingQuery
		op := &xxx_ExecuteAccountingQueryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ExecuteAccountingQueryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ExecuteAccountingQuery(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // GetRawAccountingData
		op := &xxx_GetRawAccountingDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRawAccountingDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRawAccountingData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // GetNextAccountingDataBatch
		op := &xxx_GetNextAccountingDataBatchOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNextAccountingDataBatchRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNextAccountingDataBatch(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // DeleteAccountingData
		op := &xxx_DeleteAccountingDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteAccountingDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteAccountingData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // DefragmentDB
		op := &xxx_DefragmentDBOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DefragmentDBRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DefragmentDB(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // CancelAccountingQuery
		op := &xxx_CancelAccountingQueryOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CancelAccountingQueryRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CancelAccountingQuery(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // RegisterAccountingClient
		op := &xxx_RegisterAccountingClientOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RegisterAccountingClientRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RegisterAccountingClient(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // DumpAccountingData
		op := &xxx_DumpAccountingDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DumpAccountingDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DumpAccountingData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // GetAccountingClients
		op := &xxx_GetAccountingClientsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAccountingClientsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAccountingClients(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // SetAccountingClientStatus
		op := &xxx_SetAccountingClientStatusOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetAccountingClientStatusRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetAccountingClientStatus(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // CheckAccountingConnection
		op := &xxx_CheckAccountingConnectionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CheckAccountingConnectionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CheckAccountingConnection(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // SetClientPermissions
		op := &xxx_SetClientPermissionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetClientPermissionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetClientPermissions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented IWRMAccounting
type UnimplementedAccountingServer struct {
	idispatch.UnimplementedDispatchServer
}

func (UnimplementedAccountingServer) CreateAccountingDB(context.Context, *CreateAccountingDBRequest) (*CreateAccountingDBResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) GetAccountingMetadata(context.Context, *GetAccountingMetadataRequest) (*GetAccountingMetadataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) ExecuteAccountingQuery(context.Context, *ExecuteAccountingQueryRequest) (*ExecuteAccountingQueryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) GetRawAccountingData(context.Context, *GetRawAccountingDataRequest) (*GetRawAccountingDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) GetNextAccountingDataBatch(context.Context, *GetNextAccountingDataBatchRequest) (*GetNextAccountingDataBatchResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) DeleteAccountingData(context.Context, *DeleteAccountingDataRequest) (*DeleteAccountingDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) DefragmentDB(context.Context, *DefragmentDBRequest) (*DefragmentDBResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) CancelAccountingQuery(context.Context, *CancelAccountingQueryRequest) (*CancelAccountingQueryResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) RegisterAccountingClient(context.Context, *RegisterAccountingClientRequest) (*RegisterAccountingClientResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) DumpAccountingData(context.Context, *DumpAccountingDataRequest) (*DumpAccountingDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) GetAccountingClients(context.Context, *GetAccountingClientsRequest) (*GetAccountingClientsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) SetAccountingClientStatus(context.Context, *SetAccountingClientStatusRequest) (*SetAccountingClientStatusResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) CheckAccountingConnection(context.Context, *CheckAccountingConnectionRequest) (*CheckAccountingConnectionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedAccountingServer) SetClientPermissions(context.Context, *SetClientPermissionsRequest) (*SetClientPermissionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ AccountingServer = (*UnimplementedAccountingServer)(nil)
