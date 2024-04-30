package winsif

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

// winsif server interface.
type WinsifServer interface {

	// The R_WinsRecordAction method inserts, modifies, deletes, releases, or queries a
	// name record from the WINS database.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS indicates that the operation was completed successfully. Otherwise,
	// the TimeStamp member SHOULD contain one of the following Win32 error codes, as specified
	// in [MS-ERREF]:
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN               |                                                                                  |
	//	|            VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS          | The call was successful.                                                         |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000FA0 ERROR_WINS_INTERNAL    | An error occurred while processing the RPC call.                                 |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000FA5 ERROR_REC_NON_EXISTENT | The name does not exist in the database. This error is returned only if a        |
	//	|                                   | requested WINSINTF_E_QUERY operation is not successful.                          |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED    | The caller doesn't have sufficient permissions.                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	RecordAction(context.Context, *RecordActionRequest) (*RecordActionResponse, error)

	// The R_WinsStatus method retrieves configuration settings and statistics from a WINS
	// server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation completed successfully.
	// Otherwise, this return value is a Win32 error code, as specified in [MS-ERREF]. The
	// following Win32 error codes can be returned.
	//
	//	+--------------------------------+--------------------------------------------------+
	//	|             RETURN             |                                                  |
	//	|           VALUE/CODE           |                   DESCRIPTION                    |
	//	|                                |                                                  |
	//	+--------------------------------+--------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                         |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000FA0 ERROR_WINS_INTERNAL | An error occurred while processing the RPC call. |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The caller doesn't have sufficient permissions.  |
	//	+--------------------------------+--------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	Status(context.Context, *StatusRequest) (*StatusResponse, error)

	// The R_WinsTrigger method triggers a replication operation between a target WINS server
	// and another WINS server.
	//
	// Return Values: A 32 bit unsigned integer value that indicates the return status.
	// A return value of ERROR_SUCCESS (0x00000000) indicates that the operation completed
	// successfully. Any other return value is a Win32 error code, as specified in [MS-ERREF].
	// The following Win32 error codes can be returned:
	//
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN              |                                                                                  |
	//	|            VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                  |                                                                                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS         | The call was successful.                                                         |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000FA0 ERROR_WINS_INTERNAL   | An error occurred while processing the RPC call.                                 |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000FA6 ERROR_RPL_NOT_ALLOWED | The WINS server requested for the replication operation is requested is not      |
	//	|                                  | configured as a replication partner for the target WINS server.                  |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED   | The caller does not have sufficient permissions.                                 |
	//	+----------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	Trigger(context.Context, *TriggerRequest) (*TriggerResponse, error)

	// The R_WinsDoStaticInit method performs static initialization of a WINS database by
	// registering the names specified in a data file.
	//
	// Return Values: A 32 bit unsigned integer value that indicates the return status.
	// A return value of ERROR_SUCCESS (0x00000000) indicates that the operation completed
	// successfully. Any other return value is a Win32 error code as specified in [MS-ERREF].
	// The following Win32 error codes can be returned:
	//
	//	+-------------------------------------+----------------------------------------------------------------------+
	//	|               RETURN                |                                                                      |
	//	|             VALUE/CODE              |                             DESCRIPTION                              |
	//	|                                     |                                                                      |
	//	+-------------------------------------+----------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS            | The call was successful.                                             |
	//	+-------------------------------------+----------------------------------------------------------------------+
	//	| 0x00000FA2 ERROR_STATIC_INIT_FAILED | An error occurred during static initialization of the database file. |
	//	+-------------------------------------+----------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED      | The caller does not have sufficient permissions.                     |
	//	+-------------------------------------+----------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	DoStaticInit(context.Context, *DoStaticInitRequest) (*DoStaticInitResponse, error)

	// The R_WinsDoScavenging method queues a scavenging request on the target WINS server.
	//
	// Return Values: A 32 bit unsigned integer that indicates the return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation completed successfully.
	// A nonzero return value is a Win32 error code, as specified in [MS-ERREF]. The following
	// Win32 error codes can be returned:
	//
	//	+--------------------------------+-------------------------------------------------+
	//	|             RETURN             |                                                 |
	//	|           VALUE/CODE           |                   DESCRIPTION                   |
	//	|                                |                                                 |
	//	+--------------------------------+-------------------------------------------------+
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                        |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The caller doesn't have sufficient permissions. |
	//	+--------------------------------+-------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	DoScavenging(context.Context, *DoScavengingRequest) (*DoScavengingResponse, error)

	// The R_WinsGetDbRecs method returns the records whose version numbers are within a
	// specified range and that are owned by a specified WINS server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates the return status.
	// A return value of ERROR_SUCCESS (0x00000000) indicates that the operation finished
	// successfully. Any nonzero value is a Win32 error code, as specified in [MS-ERREF].
	// The following Win32 error codes can be returned:
	//
	//	+--------------------------------+--------------------------------------------------+
	//	|             RETURN             |                                                  |
	//	|           VALUE/CODE           |                   DESCRIPTION                    |
	//	|                                |                                                  |
	//	+--------------------------------+--------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                         |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000FA0 ERROR_WINS_INTERNAL | An error occurred while processing the RPC call. |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The caller does not have sufficient permissions. |
	//	+--------------------------------+--------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetDBRecords(context.Context, *GetDBRecordsRequest) (*GetDBRecordsResponse, error)

	// The R_WinsTerm method sends a termination signal to the WINS process on a target
	// WINS server.
	//
	// Return Values: A 32 bit unsigned integer that indicates the return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that operation completed successfully.
	// A nonzero return value is a Win32 error code, as specified in [MS-ERREF]. The following
	// Win32 error codes can be returned:
	//
	//	+--------------------------------+-------------------------------------------------+
	//	|             RETURN             |                                                 |
	//	|           VALUE/CODE           |                   DESCRIPTION                   |
	//	|                                |                                                 |
	//	+--------------------------------+-------------------------------------------------+
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                        |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The caller doesn't have sufficient permissions. |
	//	+--------------------------------+-------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	Terminate(context.Context, *TerminateRequest) (*TerminateResponse, error)

	// The R_WinsBackup method backs up the WINS database to a specified directory.
	//
	// Return Values: A 32-bit unsigned integer that indicates the return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation completed successfully.
	// Any nonzero return value is a Win32 error code, as specified in [MS-ERREF]. The following
	// Win32 error codes can be returned.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                                                         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000FA0 ERROR_WINS_INTERNAL | An error occurred while processing the RPC call.                                 |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000FA4 ERROR_FULL_BACKUP   | The backup failed. Check the directory to which you are backing up the database. |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The caller doesn't have sufficient permissions.                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	Backup(context.Context, *BackupRequest) (*BackupResponse, error)

	// The R_WinsDelDbRecs method deletes the records whose version numbers are within a
	// specified range and that are owned by a specified WINS server.
	//
	// Return Values: A 32-bit unsigned integer that indicates the return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation completed successfully.
	// Any nonzero return value is a Win32 error code, as specified in [MS-ERREF]. The following
	// Win32 error codes can be returned.
	//
	//	+--------------------------------+--------------------------------------------------+
	//	|             RETURN             |                                                  |
	//	|           VALUE/CODE           |                   DESCRIPTION                    |
	//	|                                |                                                  |
	//	+--------------------------------+--------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                         |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000FA0 ERROR_WINS_INTERNAL | An error occurred while processing the RPC call. |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The caller doesn't have sufficient permissions.  |
	//	+--------------------------------+--------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	DeleteDBRecords(context.Context, *DeleteDBRecordsRequest) (*DeleteDBRecordsResponse, error)

	// The R_WinsPullRange method pulls a range of records owned by a WINS server from another
	// WINS server, and replicates them within the target WINS server database.<8>
	//
	// Return Values: A 32 bit unsigned integer that indicates the return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that operation completed successfully.
	// A nonzero return value is a Win32 error code, as specified in [MS-ERREF]. The following
	// Win32 error codes can be returned:
	//
	//	+--------------------------------+--------------------------------------------------+
	//	|             RETURN             |                                                  |
	//	|           VALUE/CODE           |                   DESCRIPTION                    |
	//	|                                |                                                  |
	//	+--------------------------------+--------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                         |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000FA0 ERROR_WINS_INTERNAL | An error occurred while processing the RPC call. |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The caller doesn't have sufficient permissions.  |
	//	+--------------------------------+--------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	PullRange(context.Context, *PullRangeRequest) (*PullRangeResponse, error)

	// The R_WinsSetPriorityClass method sets the priority class for the WINS process running
	// on the target WINS server.
	//
	// Return Values: A 32 bit unsigned integer that indicates the return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation completed successfully.
	// A nonzero return value is a Win32 error code, as specified in [MS-ERREF]. The following
	// Win32 error codes can be returned:
	//
	//	+--------------------------------+--------------------------------------------------+
	//	|             RETURN             |                                                  |
	//	|           VALUE/CODE           |                   DESCRIPTION                    |
	//	|                                |                                                  |
	//	+--------------------------------+--------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                         |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000FA0 ERROR_WINS_INTERNAL | An error occurred while processing the RPC call. |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The caller doesn't have sufficient permissions.  |
	//	+--------------------------------+--------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetPriorityClass(context.Context, *SetPriorityClassRequest) (*SetPriorityClassResponse, error)

	// The R_WinsResetCounters method resets the pull replication counters for all partners
	// of the target WINS server.
	//
	// Return Values: A 32 bit unsigned integer that indicates the return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that operation completed successfully.
	// A nonzero return value is a Win32 error code, as specified in [MS-ERREF]. The following
	// Win32 error codes can be returned:
	//
	//	+--------------------------------+-------------------------------------------------+
	//	|             RETURN             |                                                 |
	//	|           VALUE/CODE           |                   DESCRIPTION                   |
	//	|                                |                                                 |
	//	+--------------------------------+-------------------------------------------------+
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                        |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The caller doesn't have sufficient permissions. |
	//	+--------------------------------+-------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	ResetCounters(context.Context, *ResetCountersRequest) (*ResetCountersResponse, error)

	// The R_WinsWorkerThdUpd method updates the number of threads that have been created
	// to serve NetBIOS requests.
	//
	// Return Values: A 32 bit unsigned integer that indicates the return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that operation completed successfully.
	// A nonzero return value is a Win32 error code, as specified in [MS-ERREF]. The following
	// Win32 error codes can be returned:
	//
	//	+--------------------------------+--------------------------------------------------+
	//	|             RETURN             |                                                  |
	//	|           VALUE/CODE           |                   DESCRIPTION                    |
	//	|                                |                                                  |
	//	+--------------------------------+--------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                         |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000FA0 ERROR_WINS_INTERNAL | An error occurred while processing the RPC call. |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The caller doesn't have sufficient permissions.  |
	//	+--------------------------------+--------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	WorkerThreadUpdate(context.Context, *WorkerThreadUpdateRequest) (*WorkerThreadUpdateResponse, error)

	// R_WinsGetNameAndAddr operation.
	GetNameAndAddr(context.Context, *GetNameAndAddrRequest) (*GetNameAndAddrResponse, error)

	// The R_WinsGetBrowserNames_Old method always returns an ERROR_WINS_INTERNAL error
	// code.
	//
	// Return Values: A 32-bit unsigned integer value that indicates the return status.
	// The method always returns the ERROR_WINS_INTERNAL error code.
	//
	//	+--------------------------------+--------------------------------------------------+
	//	|             RETURN             |                                                  |
	//	|           VALUE/CODE           |                   DESCRIPTION                    |
	//	|                                |                                                  |
	//	+--------------------------------+--------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000FA0 ERROR_WINS_INTERNAL | An error occurred while processing the RPC call. |
	//	+--------------------------------+--------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown.
	GetBrowserNamesOld(context.Context, *GetBrowserNamesOldRequest) (*GetBrowserNamesOldResponse, error)

	// The R_WinsDeleteWins method deletes all the records owned by a particular WINS server
	// from the target WINS server database.
	//
	// Return Values: A 32 bit unsigned integer that indicates the return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that operation completed successfully.
	// A nonzero return value is a Win32 error code, as specified in [MS-ERREF]. The following
	// Win32 error codes can be returned:
	//
	//	+--------------------------------+--------------------------------------------------+
	//	|             RETURN             |                                                  |
	//	|           VALUE/CODE           |                   DESCRIPTION                    |
	//	|                                |                                                  |
	//	+--------------------------------+--------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                         |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000FA0 ERROR_WINS_INTERNAL | An error occurred while processing the RPC call. |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The caller doesn't have sufficient permissions.  |
	//	+--------------------------------+--------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	DeleteWINS(context.Context, *DeleteWINSRequest) (*DeleteWINSResponse, error)

	// The R_WinsSetFlags method always returns ERROR_SUCCESS.
	//
	// Return Values: A 32-bit unsigned integer value that indicates the return status.
	// A return value of ERROR_SUCCESS (0x00000000) indicates that the operation completed
	// successfully.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown.
	SetFlags(context.Context, *SetFlagsRequest) (*SetFlagsResponse, error)

	// The R_WinsGetBrowserNames method retrieves browser name records from the target WINS
	// server database.
	//
	// Return Values: A 32-bit unsigned integer that indicates the return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that operation completed successfully.
	// A nonzero return value is a Win32 error code, as specified in [MS-ERREF]. The following
	// Win32 error codes can be returned:
	//
	//	+--------------------------------+--------------------------------------------------+
	//	|             RETURN             |                                                  |
	//	|           VALUE/CODE           |                   DESCRIPTION                    |
	//	|                                |                                                  |
	//	+--------------------------------+--------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                         |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000FA0 ERROR_WINS_INTERNAL | An error occurred while processing the RPC call. |
	//	+--------------------------------+--------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetBrowserNames(context.Context, *GetBrowserNamesRequest) (*GetBrowserNamesResponse, error)

	// The R_WinsGetDbRecsByName method retrieves records matching an owner address from
	// a target WINS server database starting at a specified cursor.
	//
	// Return Values: A 32 bit unsigned integer that indicates the return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that operation completed successfully.
	// A nonzero return value is a Win32 error code, as specified in [MS-ERREF]. The following
	// Win32 error codes can be returned:
	//
	//	+-----------------------------------+--------------------------------------------------+
	//	|              RETURN               |                                                  |
	//	|            VALUE/CODE             |                   DESCRIPTION                    |
	//	|                                   |                                                  |
	//	+-----------------------------------+--------------------------------------------------+
	//	+-----------------------------------+--------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS          | The call was successful.                         |
	//	+-----------------------------------+--------------------------------------------------+
	//	| 0x00000FA0 ERROR_WINS_INTERNAL    | An error occurred while processing the RPC call. |
	//	+-----------------------------------+--------------------------------------------------+
	//	| 0x00000FA5 ERROR_REC_NON_EXISTENT | No records were found matching the given data.   |
	//	+-----------------------------------+--------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED    | The caller doesn't have sufficient permissions.  |
	//	+-----------------------------------+--------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetDBRecordsByName(context.Context, *GetDBRecordsByNameRequest) (*GetDBRecordsByNameResponse, error)

	// The R_WinsStatusNew method retrieves configuration settings and statistics from a
	// WINS server.
	//
	// Return Values: A 32-bit unsigned integer that indicates the return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that operation completed successfully.
	// A nonzero return value is a Win32 error code, as specified in [MS-ERREF]. The following
	// Win32 error codes can be returned:
	//
	//	+--------------------------------+--------------------------------------------------+
	//	|             RETURN             |                                                  |
	//	|           VALUE/CODE           |                   DESCRIPTION                    |
	//	|                                |                                                  |
	//	+--------------------------------+--------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                         |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The caller does not have sufficient permissions. |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000FA0 ERROR_WINS_INTERNAL | An error occurred while processing the RPC call. |
	//	+--------------------------------+--------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	StatusNew(context.Context, *StatusNewRequest) (*StatusNewResponse, error)

	// The R_WinsStatusWHdl method retrieves various configuration settings and the statistics
	// of a WINS server.
	//
	// Return Values: A 32-bit unsigned integer that indicates the return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that operation completed successfully.
	// A nonzero return value is a Win32 error code, as specified in [MS-ERREF]. The following
	// Win32 error codes can be returned:
	//
	//	+--------------------------------+--------------------------------------------------+
	//	|             RETURN             |                                                  |
	//	|           VALUE/CODE           |                   DESCRIPTION                    |
	//	|                                |                                                  |
	//	+--------------------------------+--------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                         |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000FA0 ERROR_WINS_INTERNAL | An error occurred while processing the RPC call. |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The caller doesn't have sufficient permissions.  |
	//	+--------------------------------+--------------------------------------------------+
	//
	// The behavior of this method is the same as that of the R_WinsStatusNew method (section
	// 3.1.4.20).
	StatusW(context.Context, *StatusWRequest) (*StatusWResponse, error)

	// The R_WinsDoScavengingNew method requests a specific scavenging operation on the
	// target WINS server.
	//
	// Return Values: A 32-bit unsigned integer that indicates the return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation completed successfully.
	// A nonzero return value is a Win32 error code, as specified in [MS-ERREF]. The following
	// Win32 error codes can be returned:
	//
	//	+--------------------------------+-------------------------------------------------+
	//	|             RETURN             |                                                 |
	//	|           VALUE/CODE           |                   DESCRIPTION                   |
	//	|                                |                                                 |
	//	+--------------------------------+-------------------------------------------------+
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                        |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The caller doesn't have sufficient permissions. |
	//	+--------------------------------+-------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	DoScavengingNew(context.Context, *DoScavengingNewRequest) (*DoScavengingNewResponse, error)
}

func RegisterWinsifServer(conn dcerpc.Conn, o WinsifServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWinsifServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WinsifSyntaxV1_0))...)
}

func NewWinsifServerHandle(o WinsifServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WinsifServerHandle(ctx, o, opNum, r)
	}
}

func WinsifServerHandle(ctx context.Context, o WinsifServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // R_WinsRecordAction
		in := &RecordActionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RecordAction(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // R_WinsStatus
		in := &StatusRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Status(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // R_WinsTrigger
		in := &TriggerRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Trigger(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // R_WinsDoStaticInit
		in := &DoStaticInitRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DoStaticInit(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // R_WinsDoScavenging
		in := &DoScavengingRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DoScavenging(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // R_WinsGetDbRecs
		in := &GetDBRecordsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDBRecords(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // R_WinsTerm
		in := &TerminateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Terminate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // R_WinsBackup
		in := &BackupRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Backup(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // R_WinsDelDbRecs
		in := &DeleteDBRecordsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteDBRecords(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // R_WinsPullRange
		in := &PullRangeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PullRange(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // R_WinsSetPriorityClass
		in := &SetPriorityClassRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPriorityClass(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // R_WinsResetCounters
		in := &ResetCountersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ResetCounters(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // R_WinsWorkerThdUpd
		in := &WorkerThreadUpdateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.WorkerThreadUpdate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // R_WinsGetNameAndAddr
		in := &GetNameAndAddrRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNameAndAddr(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // R_WinsGetBrowserNames_Old
		in := &GetBrowserNamesOldRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetBrowserNamesOld(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // R_WinsDeleteWins
		in := &DeleteWINSRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteWINS(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // R_WinsSetFlags
		in := &SetFlagsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFlags(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // R_WinsGetBrowserNames
		in := &GetBrowserNamesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetBrowserNames(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // R_WinsGetDbRecsByName
		in := &GetDBRecordsByNameRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDBRecordsByName(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // R_WinsStatusNew
		in := &StatusNewRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.StatusNew(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // R_WinsStatusWHdl
		in := &StatusWRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.StatusW(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // R_WinsDoScavengingNew
		in := &DoScavengingNewRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DoScavengingNew(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
