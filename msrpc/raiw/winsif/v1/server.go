package winsif

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
		op := &xxx_RecordActionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RecordActionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RecordAction(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // R_WinsStatus
		op := &xxx_StatusOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StatusRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Status(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // R_WinsTrigger
		op := &xxx_TriggerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &TriggerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Trigger(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // R_WinsDoStaticInit
		op := &xxx_DoStaticInitOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DoStaticInitRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DoStaticInit(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // R_WinsDoScavenging
		op := &xxx_DoScavengingOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DoScavengingRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DoScavenging(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // R_WinsGetDbRecs
		op := &xxx_GetDBRecordsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDBRecordsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDBRecords(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // R_WinsTerm
		op := &xxx_TerminateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &TerminateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Terminate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // R_WinsBackup
		op := &xxx_BackupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BackupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Backup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // R_WinsDelDbRecs
		op := &xxx_DeleteDBRecordsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteDBRecordsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteDBRecords(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // R_WinsPullRange
		op := &xxx_PullRangeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &PullRangeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.PullRange(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // R_WinsSetPriorityClass
		op := &xxx_SetPriorityClassOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPriorityClassRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPriorityClass(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // R_WinsResetCounters
		op := &xxx_ResetCountersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ResetCountersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ResetCounters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // R_WinsWorkerThdUpd
		op := &xxx_WorkerThreadUpdateOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WorkerThreadUpdateRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WorkerThreadUpdate(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // R_WinsGetNameAndAddr
		op := &xxx_GetNameAndAddrOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetNameAndAddrRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetNameAndAddr(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // R_WinsGetBrowserNames_Old
		op := &xxx_GetBrowserNamesOldOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetBrowserNamesOldRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetBrowserNamesOld(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // R_WinsDeleteWins
		op := &xxx_DeleteWINSOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteWINSRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteWINS(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // R_WinsSetFlags
		op := &xxx_SetFlagsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetFlagsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetFlags(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // R_WinsGetBrowserNames
		op := &xxx_GetBrowserNamesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetBrowserNamesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetBrowserNames(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // R_WinsGetDbRecsByName
		op := &xxx_GetDBRecordsByNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDBRecordsByNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDBRecordsByName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // R_WinsStatusNew
		op := &xxx_StatusNewOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StatusNewRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StatusNew(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // R_WinsStatusWHdl
		op := &xxx_StatusWOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StatusWRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StatusW(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // R_WinsDoScavengingNew
		op := &xxx_DoScavengingNewOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DoScavengingNewRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DoScavengingNew(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented winsif
type UnimplementedWinsifServer struct {
}

func (UnimplementedWinsifServer) RecordAction(context.Context, *RecordActionRequest) (*RecordActionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) Trigger(context.Context, *TriggerRequest) (*TriggerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) DoStaticInit(context.Context, *DoStaticInitRequest) (*DoStaticInitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) DoScavenging(context.Context, *DoScavengingRequest) (*DoScavengingResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) GetDBRecords(context.Context, *GetDBRecordsRequest) (*GetDBRecordsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) Terminate(context.Context, *TerminateRequest) (*TerminateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) Backup(context.Context, *BackupRequest) (*BackupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) DeleteDBRecords(context.Context, *DeleteDBRecordsRequest) (*DeleteDBRecordsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) PullRange(context.Context, *PullRangeRequest) (*PullRangeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) SetPriorityClass(context.Context, *SetPriorityClassRequest) (*SetPriorityClassResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) ResetCounters(context.Context, *ResetCountersRequest) (*ResetCountersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) WorkerThreadUpdate(context.Context, *WorkerThreadUpdateRequest) (*WorkerThreadUpdateResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) GetNameAndAddr(context.Context, *GetNameAndAddrRequest) (*GetNameAndAddrResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) GetBrowserNamesOld(context.Context, *GetBrowserNamesOldRequest) (*GetBrowserNamesOldResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) DeleteWINS(context.Context, *DeleteWINSRequest) (*DeleteWINSResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) SetFlags(context.Context, *SetFlagsRequest) (*SetFlagsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) GetBrowserNames(context.Context, *GetBrowserNamesRequest) (*GetBrowserNamesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) GetDBRecordsByName(context.Context, *GetDBRecordsByNameRequest) (*GetDBRecordsByNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) StatusNew(context.Context, *StatusNewRequest) (*StatusNewResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) StatusW(context.Context, *StatusWRequest) (*StatusWResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinsifServer) DoScavengingNew(context.Context, *DoScavengingNewRequest) (*DoScavengingNewResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ WinsifServer = (*UnimplementedWinsifServer)(nil)
