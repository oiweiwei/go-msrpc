package winsif

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	raiw "github.com/oiweiwei/go-msrpc/msrpc/raiw"
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
	_ = raiw.GoPackage
)

var (
	// import guard
	GoPackage = "raiw"
)

var (
	// Syntax UUID
	WinsifSyntaxUUID = &uuid.UUID{TimeLow: 0x45f52c28, TimeMid: 0x7f9f, TimeHiAndVersion: 0x101a, ClockSeqHiAndReserved: 0xb5, ClockSeqLow: 0x2b, Node: [6]uint8{0x8, 0x0, 0x2b, 0x2e, 0xfa, 0xbe}}
	// Syntax ID
	WinsifSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: WinsifSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// winsif interface.
type WinsifClient interface {

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
	RecordAction(context.Context, *RecordActionRequest, ...dcerpc.CallOption) (*RecordActionResponse, error)

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
	Status(context.Context, *StatusRequest, ...dcerpc.CallOption) (*StatusResponse, error)

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
	Trigger(context.Context, *TriggerRequest, ...dcerpc.CallOption) (*TriggerResponse, error)

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
	DoStaticInit(context.Context, *DoStaticInitRequest, ...dcerpc.CallOption) (*DoStaticInitResponse, error)

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
	DoScavenging(context.Context, *DoScavengingRequest, ...dcerpc.CallOption) (*DoScavengingResponse, error)

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
	GetDBRecords(context.Context, *GetDBRecordsRequest, ...dcerpc.CallOption) (*GetDBRecordsResponse, error)

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
	Terminate(context.Context, *TerminateRequest, ...dcerpc.CallOption) (*TerminateResponse, error)

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
	Backup(context.Context, *BackupRequest, ...dcerpc.CallOption) (*BackupResponse, error)

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
	DeleteDBRecords(context.Context, *DeleteDBRecordsRequest, ...dcerpc.CallOption) (*DeleteDBRecordsResponse, error)

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
	PullRange(context.Context, *PullRangeRequest, ...dcerpc.CallOption) (*PullRangeResponse, error)

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
	SetPriorityClass(context.Context, *SetPriorityClassRequest, ...dcerpc.CallOption) (*SetPriorityClassResponse, error)

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
	ResetCounters(context.Context, *ResetCountersRequest, ...dcerpc.CallOption) (*ResetCountersResponse, error)

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
	WorkerThreadUpdate(context.Context, *WorkerThreadUpdateRequest, ...dcerpc.CallOption) (*WorkerThreadUpdateResponse, error)

	// R_WinsGetNameAndAddr operation.
	GetNameAndAddr(context.Context, *GetNameAndAddrRequest, ...dcerpc.CallOption) (*GetNameAndAddrResponse, error)

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
	GetBrowserNamesOld(context.Context, *GetBrowserNamesOldRequest, ...dcerpc.CallOption) (*GetBrowserNamesOldResponse, error)

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
	DeleteWINS(context.Context, *DeleteWINSRequest, ...dcerpc.CallOption) (*DeleteWINSResponse, error)

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
	SetFlags(context.Context, *SetFlagsRequest, ...dcerpc.CallOption) (*SetFlagsResponse, error)

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
	GetBrowserNames(context.Context, *GetBrowserNamesRequest, ...dcerpc.CallOption) (*GetBrowserNamesResponse, error)

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
	GetDBRecordsByName(context.Context, *GetDBRecordsByNameRequest, ...dcerpc.CallOption) (*GetDBRecordsByNameResponse, error)

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
	StatusNew(context.Context, *StatusNewRequest, ...dcerpc.CallOption) (*StatusNewResponse, error)

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
	StatusW(context.Context, *StatusWRequest, ...dcerpc.CallOption) (*StatusWResponse, error)

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
	DoScavengingNew(context.Context, *DoScavengingNewRequest, ...dcerpc.CallOption) (*DoScavengingNewResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error
}

type xxx_DefaultWinsifClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultWinsifClient) RecordAction(ctx context.Context, in *RecordActionRequest, opts ...dcerpc.CallOption) (*RecordActionResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RecordActionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) Status(ctx context.Context, in *StatusRequest, opts ...dcerpc.CallOption) (*StatusResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &StatusResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) Trigger(ctx context.Context, in *TriggerRequest, opts ...dcerpc.CallOption) (*TriggerResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &TriggerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) DoStaticInit(ctx context.Context, in *DoStaticInitRequest, opts ...dcerpc.CallOption) (*DoStaticInitResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DoStaticInitResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) DoScavenging(ctx context.Context, in *DoScavengingRequest, opts ...dcerpc.CallOption) (*DoScavengingResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DoScavengingResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) GetDBRecords(ctx context.Context, in *GetDBRecordsRequest, opts ...dcerpc.CallOption) (*GetDBRecordsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetDBRecordsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) Terminate(ctx context.Context, in *TerminateRequest, opts ...dcerpc.CallOption) (*TerminateResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &TerminateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) Backup(ctx context.Context, in *BackupRequest, opts ...dcerpc.CallOption) (*BackupResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BackupResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) DeleteDBRecords(ctx context.Context, in *DeleteDBRecordsRequest, opts ...dcerpc.CallOption) (*DeleteDBRecordsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteDBRecordsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) PullRange(ctx context.Context, in *PullRangeRequest, opts ...dcerpc.CallOption) (*PullRangeResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PullRangeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) SetPriorityClass(ctx context.Context, in *SetPriorityClassRequest, opts ...dcerpc.CallOption) (*SetPriorityClassResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetPriorityClassResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) ResetCounters(ctx context.Context, in *ResetCountersRequest, opts ...dcerpc.CallOption) (*ResetCountersResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ResetCountersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) WorkerThreadUpdate(ctx context.Context, in *WorkerThreadUpdateRequest, opts ...dcerpc.CallOption) (*WorkerThreadUpdateResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &WorkerThreadUpdateResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) GetNameAndAddr(ctx context.Context, in *GetNameAndAddrRequest, opts ...dcerpc.CallOption) (*GetNameAndAddrResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetNameAndAddrResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) GetBrowserNamesOld(ctx context.Context, in *GetBrowserNamesOldRequest, opts ...dcerpc.CallOption) (*GetBrowserNamesOldResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetBrowserNamesOldResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) DeleteWINS(ctx context.Context, in *DeleteWINSRequest, opts ...dcerpc.CallOption) (*DeleteWINSResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DeleteWINSResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) SetFlags(ctx context.Context, in *SetFlagsRequest, opts ...dcerpc.CallOption) (*SetFlagsResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetFlagsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) GetBrowserNames(ctx context.Context, in *GetBrowserNamesRequest, opts ...dcerpc.CallOption) (*GetBrowserNamesResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetBrowserNamesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) GetDBRecordsByName(ctx context.Context, in *GetDBRecordsByNameRequest, opts ...dcerpc.CallOption) (*GetDBRecordsByNameResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetDBRecordsByNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) StatusNew(ctx context.Context, in *StatusNewRequest, opts ...dcerpc.CallOption) (*StatusNewResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &StatusNewResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) StatusW(ctx context.Context, in *StatusWRequest, opts ...dcerpc.CallOption) (*StatusWResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &StatusWResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) DoScavengingNew(ctx context.Context, in *DoScavengingNewRequest, opts ...dcerpc.CallOption) (*DoScavengingNewResponse, error) {
	op := in.xxx_ToOp(ctx)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &DoScavengingNewResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinsifClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}
func NewWinsifClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (WinsifClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(WinsifSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultWinsifClient{cc: cc}, nil
}

// xxx_RecordActionOperation structure represents the R_WinsRecordAction operation
type xxx_RecordActionOperation struct {
	RecordAction *raiw.RecordAction `idl:"name:ppRecAction;pointer:ref" json:"record_action"`
	Return       uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_RecordActionOperation) OpNum() int { return 0 }

func (o *xxx_RecordActionOperation) OpName() string { return "/winsif/v1/R_WinsRecordAction" }

func (o *xxx_RecordActionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RecordActionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ppRecAction {in, out} (1:{pointer=ref}*(2))(2:{alias=PWINSINTF_RECORD_ACTION_T}*(1))(3:{alias=WINSINTF_RECORD_ACTION_T}(struct))
	{
		if o.RecordAction != nil {
			_ptr_ppRecAction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RecordAction != nil {
					if err := o.RecordAction.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&raiw.RecordAction{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RecordAction, _ptr_ppRecAction); err != nil {
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

func (o *xxx_RecordActionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ppRecAction {in, out} (1:{pointer=ref}*(2))(2:{alias=PWINSINTF_RECORD_ACTION_T,pointer=ref}*(1))(3:{alias=WINSINTF_RECORD_ACTION_T}(struct))
	{
		_ptr_ppRecAction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RecordAction == nil {
				o.RecordAction = &raiw.RecordAction{}
			}
			if err := o.RecordAction.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppRecAction := func(ptr interface{}) { o.RecordAction = *ptr.(**raiw.RecordAction) }
		if err := w.ReadPointer(&o.RecordAction, _s_ppRecAction, _ptr_ppRecAction); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RecordActionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RecordActionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppRecAction {in, out} (1:{pointer=ref}*(2))(2:{alias=PWINSINTF_RECORD_ACTION_T}*(1))(3:{alias=WINSINTF_RECORD_ACTION_T}(struct))
	{
		if o.RecordAction != nil {
			_ptr_ppRecAction := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RecordAction != nil {
					if err := o.RecordAction.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&raiw.RecordAction{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RecordAction, _ptr_ppRecAction); err != nil {
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
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RecordActionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppRecAction {in, out} (1:{pointer=ref}*(2))(2:{alias=PWINSINTF_RECORD_ACTION_T,pointer=ref}*(1))(3:{alias=WINSINTF_RECORD_ACTION_T}(struct))
	{
		_ptr_ppRecAction := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.RecordAction == nil {
				o.RecordAction = &raiw.RecordAction{}
			}
			if err := o.RecordAction.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppRecAction := func(ptr interface{}) { o.RecordAction = *ptr.(**raiw.RecordAction) }
		if err := w.ReadPointer(&o.RecordAction, _s_ppRecAction, _ptr_ppRecAction); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RecordActionRequest structure represents the R_WinsRecordAction operation request
type RecordActionRequest struct {
	// ppRecAction: A pointer to a WINSINTF_RECORD_ACTION_T structure (section 2.2.2.3)
	// that contains the details of the record and the action to be performed on it. The
	// interpretation of the member values in this structure depends on the type of action
	// specified by the WINSINTF_ACT_E enumeration (section 2.2.1.4) value in its Cmd_e
	// member, as follows.
	//
	// WINSINTF_E_INSERT:
	//
	// § Cmd_e is set to WINSINTF_E_INSERT.
	//
	// § pName points to a NULL-terminated string that contains the NetBIOS name and optionally
	// the NetBIOS scope name of the record.
	//
	// § NameLen contains the length of the string specified by pName.
	//
	// § TypOfRec_e is set to a value between 0x00000000 and 0x00000003 based on the record
	// type.
	//
	// § NoOfAdds is set to a positive value based on the number of IP address mappings
	// that the record has.
	//
	// § pAdd or Add is set with the mapping IP addresses, based on the TypOfRec_e member.
	//
	// § VersNo SHOULD be ignored by the server. The inserted record MUST be marked with
	// the current version number that is in use at the WINS server.
	//
	// § NodeTyp is set to a value between 0x00 and 0x03 based on the type of the node.
	//
	// § OwnerId SHOULD be ignored by the server. The record MUST be inserted into the
	// database with the OwnerId member set to the target WINS server address.
	//
	// § State_e SHOULD be ignored by the server. The record MUST be inserted into the
	// database with its state marked as ACTIVE.
	//
	// § fStatic is set to 0x00000001 if the record being inserted is a static record;
	// otherwise, it is set to 0x00000000.
	//
	// § TimeStamp SHOULD be ignored by the server. The inserted record SHOULD be time-stamped
	// with zero if the fStatic member is set to 0x00000001; otherwise, it SHOULD be time-stamped
	// with the current time on the server plus the refresh interval configured on the server.
	//
	// WINSINTF_E_DELETE:
	//
	// § Cmd_e is set to WINSINTF_E_DELETE.
	//
	// § pName points to a NULL-terminated string that contains the NetBIOS name and optionally
	// the NetBIOS scope name of the record to be deleted from the database.
	//
	// § NameLen contains the length of the string specified by pName.
	//
	// § State_e is set to 0x00000003.
	//
	// § All other members SHOULD be ignored by the server.
	//
	// § WINSINTF_E_RELEASE:
	//
	// § Cmd_e is set to WINSINTF_E_RELEASE.
	//
	// § pName points to a NULL-terminated string that contains the NetBIOS name and optionally
	// the NetBIOS scope name of the record.
	//
	// § NameLen contains the length of the string specified by pName.
	//
	// § TypOfRec_e is set to a value between zero and 0x00000003 based on the record type.
	//
	// § NoOfAdds MUST be set to 0x00000001.
	//
	// § pAdd or Add is set with the mapping IP address based on the TypOfRec_e member.
	//
	// § VersNo, NodeTyp, OwnerId, and fStatic SHOULD be ignored by the server.
	//
	// § State_e SHOULD be ignored by the server. The record MUST be inserted with state
	// marked as RELEASED.
	//
	// § TimeStamp SHOULD be ignored by the server. The released record SHOULD be time-stamped
	// with 0xFFFFFFFF if the fStatic member is set to 0x00000001; otherwise, it SHOULD
	// be time-stamped with the current time on the server plus the tombstone interval configured
	// on the server.
	//
	// WINSINTF_E_MODIFY:
	//
	// § Cmd_e is set to WINSINTF_E_MODIFY.
	//
	// § pName points to a NULL-terminated string that contains the NetBIOS name and optionally
	// the NetBIOS scope name of the record to be modified in the database.
	//
	// § NameLen contains the length of the string specified by pName.
	//
	// § TypOfRec_e contains the record type to be set for the record matching the pName
	// member in the WINS database.
	//
	// § NodeTyp contains the node type to be set for the record matching the pName member
	// in the WINS database.
	//
	// § State_e contains the state to be set for the record matching the pName member
	// in the WINS database.
	//
	// § fStatic contains the  value to be set for the record matching the pName member
	// in the WINS database.
	//
	// § All other members SHOULD be ignored by the server.
	//
	// WINSINTF_E_QUERY:
	//
	// § Cmd_e is set to WINSINTF_E_QUERY.
	//
	// § pName points to a NULL-terminated string that contains the NetBIOS name and optionally
	// the NetBIOS scope name of the record to be queried from the database.
	//
	// § NameLen contains the length of the string specified by pName.
	//
	// § All other members act as output, which are filled by the server if a matching
	// entry is found in the database.
	//
	// § TypOfRec_e contains the matching record type.
	//
	// § If the TyeOfRec_e member is set to 0x00000000 or 0x00000001, the NoOfAdds member
	// SHOULD contain 0x00000001 or the number of IP addresses that are mapped to the name
	// given in the pName member.
	//
	// § If the TypOfRec_e member is set to 0x00000002 or 0x00000003. The RPC method caller
	// SHOULD refer to this member for the set of IP addresses mapped to the name given
	// in the pName member.
	//
	// § If the TypOfRec_e member is set to 0x00000000 or 0x00000001. The RPC method caller
	// SHOULD refer to this member for the IP address mapped to the name given in the pName
	// member. If the TypOfRec_e member is set to 0x00000001, the IPAdd member of the Add
	// structure MUST contain 0xFFFFFFFF.
	//
	// § VersNo contains the version number of the matching record.
	//
	// § NodeTyp contains the node type of the matching record.
	//
	// § OwnerId contains the IP address of the owner of the matching record.
	//
	// § State_e contains the state of the matching record.
	//
	// § fStatic contains the value 0x00000001 if the record is entered into the database
	// by an administrator; otherwise, it contains 0x00000000.
	RecordAction *raiw.RecordAction `idl:"name:ppRecAction;pointer:ref" json:"record_action"`
}

func (o *RecordActionRequest) xxx_ToOp(ctx context.Context) *xxx_RecordActionOperation {
	if o == nil {
		return &xxx_RecordActionOperation{}
	}
	return &xxx_RecordActionOperation{
		RecordAction: o.RecordAction,
	}
}

func (o *RecordActionRequest) xxx_FromOp(ctx context.Context, op *xxx_RecordActionOperation) {
	if o == nil {
		return
	}
	o.RecordAction = op.RecordAction
}
func (o *RecordActionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *RecordActionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RecordActionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RecordActionResponse structure represents the R_WinsRecordAction operation response
type RecordActionResponse struct {
	// ppRecAction: A pointer to a WINSINTF_RECORD_ACTION_T structure (section 2.2.2.3)
	// that contains the details of the record and the action to be performed on it. The
	// interpretation of the member values in this structure depends on the type of action
	// specified by the WINSINTF_ACT_E enumeration (section 2.2.1.4) value in its Cmd_e
	// member, as follows.
	//
	// WINSINTF_E_INSERT:
	//
	// § Cmd_e is set to WINSINTF_E_INSERT.
	//
	// § pName points to a NULL-terminated string that contains the NetBIOS name and optionally
	// the NetBIOS scope name of the record.
	//
	// § NameLen contains the length of the string specified by pName.
	//
	// § TypOfRec_e is set to a value between 0x00000000 and 0x00000003 based on the record
	// type.
	//
	// § NoOfAdds is set to a positive value based on the number of IP address mappings
	// that the record has.
	//
	// § pAdd or Add is set with the mapping IP addresses, based on the TypOfRec_e member.
	//
	// § VersNo SHOULD be ignored by the server. The inserted record MUST be marked with
	// the current version number that is in use at the WINS server.
	//
	// § NodeTyp is set to a value between 0x00 and 0x03 based on the type of the node.
	//
	// § OwnerId SHOULD be ignored by the server. The record MUST be inserted into the
	// database with the OwnerId member set to the target WINS server address.
	//
	// § State_e SHOULD be ignored by the server. The record MUST be inserted into the
	// database with its state marked as ACTIVE.
	//
	// § fStatic is set to 0x00000001 if the record being inserted is a static record;
	// otherwise, it is set to 0x00000000.
	//
	// § TimeStamp SHOULD be ignored by the server. The inserted record SHOULD be time-stamped
	// with zero if the fStatic member is set to 0x00000001; otherwise, it SHOULD be time-stamped
	// with the current time on the server plus the refresh interval configured on the server.
	//
	// WINSINTF_E_DELETE:
	//
	// § Cmd_e is set to WINSINTF_E_DELETE.
	//
	// § pName points to a NULL-terminated string that contains the NetBIOS name and optionally
	// the NetBIOS scope name of the record to be deleted from the database.
	//
	// § NameLen contains the length of the string specified by pName.
	//
	// § State_e is set to 0x00000003.
	//
	// § All other members SHOULD be ignored by the server.
	//
	// § WINSINTF_E_RELEASE:
	//
	// § Cmd_e is set to WINSINTF_E_RELEASE.
	//
	// § pName points to a NULL-terminated string that contains the NetBIOS name and optionally
	// the NetBIOS scope name of the record.
	//
	// § NameLen contains the length of the string specified by pName.
	//
	// § TypOfRec_e is set to a value between zero and 0x00000003 based on the record type.
	//
	// § NoOfAdds MUST be set to 0x00000001.
	//
	// § pAdd or Add is set with the mapping IP address based on the TypOfRec_e member.
	//
	// § VersNo, NodeTyp, OwnerId, and fStatic SHOULD be ignored by the server.
	//
	// § State_e SHOULD be ignored by the server. The record MUST be inserted with state
	// marked as RELEASED.
	//
	// § TimeStamp SHOULD be ignored by the server. The released record SHOULD be time-stamped
	// with 0xFFFFFFFF if the fStatic member is set to 0x00000001; otherwise, it SHOULD
	// be time-stamped with the current time on the server plus the tombstone interval configured
	// on the server.
	//
	// WINSINTF_E_MODIFY:
	//
	// § Cmd_e is set to WINSINTF_E_MODIFY.
	//
	// § pName points to a NULL-terminated string that contains the NetBIOS name and optionally
	// the NetBIOS scope name of the record to be modified in the database.
	//
	// § NameLen contains the length of the string specified by pName.
	//
	// § TypOfRec_e contains the record type to be set for the record matching the pName
	// member in the WINS database.
	//
	// § NodeTyp contains the node type to be set for the record matching the pName member
	// in the WINS database.
	//
	// § State_e contains the state to be set for the record matching the pName member
	// in the WINS database.
	//
	// § fStatic contains the  value to be set for the record matching the pName member
	// in the WINS database.
	//
	// § All other members SHOULD be ignored by the server.
	//
	// WINSINTF_E_QUERY:
	//
	// § Cmd_e is set to WINSINTF_E_QUERY.
	//
	// § pName points to a NULL-terminated string that contains the NetBIOS name and optionally
	// the NetBIOS scope name of the record to be queried from the database.
	//
	// § NameLen contains the length of the string specified by pName.
	//
	// § All other members act as output, which are filled by the server if a matching
	// entry is found in the database.
	//
	// § TypOfRec_e contains the matching record type.
	//
	// § If the TyeOfRec_e member is set to 0x00000000 or 0x00000001, the NoOfAdds member
	// SHOULD contain 0x00000001 or the number of IP addresses that are mapped to the name
	// given in the pName member.
	//
	// § If the TypOfRec_e member is set to 0x00000002 or 0x00000003. The RPC method caller
	// SHOULD refer to this member for the set of IP addresses mapped to the name given
	// in the pName member.
	//
	// § If the TypOfRec_e member is set to 0x00000000 or 0x00000001. The RPC method caller
	// SHOULD refer to this member for the IP address mapped to the name given in the pName
	// member. If the TypOfRec_e member is set to 0x00000001, the IPAdd member of the Add
	// structure MUST contain 0xFFFFFFFF.
	//
	// § VersNo contains the version number of the matching record.
	//
	// § NodeTyp contains the node type of the matching record.
	//
	// § OwnerId contains the IP address of the owner of the matching record.
	//
	// § State_e contains the state of the matching record.
	//
	// § fStatic contains the value 0x00000001 if the record is entered into the database
	// by an administrator; otherwise, it contains 0x00000000.
	RecordAction *raiw.RecordAction `idl:"name:ppRecAction;pointer:ref" json:"record_action"`
	// Return: The R_WinsRecordAction return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RecordActionResponse) xxx_ToOp(ctx context.Context) *xxx_RecordActionOperation {
	if o == nil {
		return &xxx_RecordActionOperation{}
	}
	return &xxx_RecordActionOperation{
		RecordAction: o.RecordAction,
		Return:       o.Return,
	}
}

func (o *RecordActionResponse) xxx_FromOp(ctx context.Context, op *xxx_RecordActionOperation) {
	if o == nil {
		return
	}
	o.RecordAction = op.RecordAction
	o.Return = op.Return
}
func (o *RecordActionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *RecordActionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RecordActionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StatusOperation structure represents the R_WinsStatus operation
type xxx_StatusOperation struct {
	Cmd     raiw.Cmd      `idl:"name:Cmd_e" json:"cmd"`
	Results *raiw.Results `idl:"name:pResults;pointer:ref" json:"results"`
	Return  uint32        `idl:"name:Return" json:"return"`
}

func (o *xxx_StatusOperation) OpNum() int { return 1 }

func (o *xxx_StatusOperation) OpName() string { return "/winsif/v1/R_WinsStatus" }

func (o *xxx_StatusOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatusOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Cmd_e {in} (1:{alias=WINSINTF_CMD_E}(enum))
	{
		if err := w.WriteData(uint16(o.Cmd)); err != nil {
			return err
		}
	}
	// pResults {in, out} (1:{pointer=ref, alias=PWINSINTF_RESULTS_T}*(1))(2:{alias=WINSINTF_RESULTS_T}(struct))
	{
		if o.Results != nil {
			if err := o.Results.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.Results{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatusOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Cmd_e {in} (1:{alias=WINSINTF_CMD_E}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Cmd)); err != nil {
			return err
		}
	}
	// pResults {in, out} (1:{pointer=ref, alias=PWINSINTF_RESULTS_T}*(1))(2:{alias=WINSINTF_RESULTS_T}(struct))
	{
		if o.Results == nil {
			o.Results = &raiw.Results{}
		}
		if err := o.Results.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatusOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatusOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResults {in, out} (1:{pointer=ref, alias=PWINSINTF_RESULTS_T}*(1))(2:{alias=WINSINTF_RESULTS_T}(struct))
	{
		if o.Results != nil {
			if err := o.Results.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.Results{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatusOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResults {in, out} (1:{pointer=ref, alias=PWINSINTF_RESULTS_T}*(1))(2:{alias=WINSINTF_RESULTS_T}(struct))
	{
		if o.Results == nil {
			o.Results = &raiw.Results{}
		}
		if err := o.Results.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// StatusRequest structure represents the R_WinsStatus operation request
type StatusRequest struct {
	// Cmd_e: The command to be executed on the target WINS server from the WINSINTF_CMD_E
	// enumeration (section 2.2.1.5).
	Cmd raiw.Cmd `idl:"name:Cmd_e" json:"cmd"`
	// pResults: A pointer to a WINSINTF_RESULTS_T structure (section 2.2.2.7) that contains
	// configuration data and statistics for the target WINS server.
	Results *raiw.Results `idl:"name:pResults;pointer:ref" json:"results"`
}

func (o *StatusRequest) xxx_ToOp(ctx context.Context) *xxx_StatusOperation {
	if o == nil {
		return &xxx_StatusOperation{}
	}
	return &xxx_StatusOperation{
		Cmd:     o.Cmd,
		Results: o.Results,
	}
}

func (o *StatusRequest) xxx_FromOp(ctx context.Context, op *xxx_StatusOperation) {
	if o == nil {
		return
	}
	o.Cmd = op.Cmd
	o.Results = op.Results
}
func (o *StatusRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *StatusRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StatusOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StatusResponse structure represents the R_WinsStatus operation response
type StatusResponse struct {
	// pResults: A pointer to a WINSINTF_RESULTS_T structure (section 2.2.2.7) that contains
	// configuration data and statistics for the target WINS server.
	Results *raiw.Results `idl:"name:pResults;pointer:ref" json:"results"`
	// Return: The R_WinsStatus return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *StatusResponse) xxx_ToOp(ctx context.Context) *xxx_StatusOperation {
	if o == nil {
		return &xxx_StatusOperation{}
	}
	return &xxx_StatusOperation{
		Results: o.Results,
		Return:  o.Return,
	}
}

func (o *StatusResponse) xxx_FromOp(ctx context.Context, op *xxx_StatusOperation) {
	if o == nil {
		return
	}
	o.Results = op.Results
	o.Return = op.Return
}
func (o *StatusResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *StatusResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StatusOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_TriggerOperation structure represents the R_WinsTrigger operation
type xxx_TriggerOperation struct {
	WINSAddr     *raiw.Addr        `idl:"name:pWinsAddr;pointer:ref" json:"wins_addr"`
	TriggerTypeE raiw.TriggerTypeE `idl:"name:TrigType_e" json:"trigger_type_e"`
	Return       uint32            `idl:"name:Return" json:"return"`
}

func (o *xxx_TriggerOperation) OpNum() int { return 2 }

func (o *xxx_TriggerOperation) OpName() string { return "/winsif/v1/R_WinsTrigger" }

func (o *xxx_TriggerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TriggerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pWinsAddr {in} (1:{pointer=ref, alias=PWINSINTF_ADDR_T}*(1))(2:{alias=WINSINTF_ADDR_T}(struct))
	{
		if o.WINSAddr != nil {
			if err := o.WINSAddr.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.Addr{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// TrigType_e {in} (1:{alias=WINSINTF_TRIG_TYPE_E}(enum))
	{
		if err := w.WriteData(uint16(o.TriggerTypeE)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TriggerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pWinsAddr {in} (1:{pointer=ref, alias=PWINSINTF_ADDR_T}*(1))(2:{alias=WINSINTF_ADDR_T}(struct))
	{
		if o.WINSAddr == nil {
			o.WINSAddr = &raiw.Addr{}
		}
		if err := o.WINSAddr.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// TrigType_e {in} (1:{alias=WINSINTF_TRIG_TYPE_E}(enum))
	{
		if err := w.ReadData((*uint16)(&o.TriggerTypeE)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TriggerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TriggerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TriggerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// TriggerRequest structure represents the R_WinsTrigger operation request
type TriggerRequest struct {
	WINSAddr *raiw.Addr `idl:"name:pWinsAddr;pointer:ref" json:"wins_addr"`
	// TrigType_e: The type of replication operation requested.
	TriggerTypeE raiw.TriggerTypeE `idl:"name:TrigType_e" json:"trigger_type_e"`
}

func (o *TriggerRequest) xxx_ToOp(ctx context.Context) *xxx_TriggerOperation {
	if o == nil {
		return &xxx_TriggerOperation{}
	}
	return &xxx_TriggerOperation{
		WINSAddr:     o.WINSAddr,
		TriggerTypeE: o.TriggerTypeE,
	}
}

func (o *TriggerRequest) xxx_FromOp(ctx context.Context, op *xxx_TriggerOperation) {
	if o == nil {
		return
	}
	o.WINSAddr = op.WINSAddr
	o.TriggerTypeE = op.TriggerTypeE
}
func (o *TriggerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *TriggerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_TriggerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// TriggerResponse structure represents the R_WinsTrigger operation response
type TriggerResponse struct {
	// Return: The R_WinsTrigger return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *TriggerResponse) xxx_ToOp(ctx context.Context) *xxx_TriggerOperation {
	if o == nil {
		return &xxx_TriggerOperation{}
	}
	return &xxx_TriggerOperation{
		Return: o.Return,
	}
}

func (o *TriggerResponse) xxx_FromOp(ctx context.Context, op *xxx_TriggerOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *TriggerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *TriggerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_TriggerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DoStaticInitOperation structure represents the R_WinsDoStaticInit operation
type xxx_DoStaticInitOperation struct {
	DataFilePath string `idl:"name:pDataFilePath;string;pointer:unique" json:"data_file_path"`
	Delete       uint32 `idl:"name:fDel" json:"delete"`
	Return       uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_DoStaticInitOperation) OpNum() int { return 3 }

func (o *xxx_DoStaticInitOperation) OpName() string { return "/winsif/v1/R_WinsDoStaticInit" }

func (o *xxx_DoStaticInitOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoStaticInitOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pDataFilePath {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.DataFilePath != "" {
			_ptr_pDataFilePath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.DataFilePath); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DataFilePath, _ptr_pDataFilePath); err != nil {
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
	// fDel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Delete); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoStaticInitOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pDataFilePath {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pDataFilePath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.DataFilePath); err != nil {
				return err
			}
			return nil
		})
		_s_pDataFilePath := func(ptr interface{}) { o.DataFilePath = *ptr.(*string) }
		if err := w.ReadPointer(&o.DataFilePath, _s_pDataFilePath, _ptr_pDataFilePath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// fDel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Delete); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoStaticInitOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoStaticInitOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoStaticInitOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DoStaticInitRequest structure represents the R_WinsDoStaticInit operation request
type DoStaticInitRequest struct {
	// pDataFilePath: A pointer to a Unicode string containing the path to a text file on
	// the target WINS server. The file SHOULD contain entries that map NetBIOS names to
	// IPv4 addresses in string format using the following syntax:
	//
	// <IPv4 address 1> <one or more spaces> <NetBIOS name 1> <IPv4 address 2> <one or more
	// spaces> <NetBIOS name 2> ... <IPv4 address N> <one or more spaces> <NetBIOS name
	// N>
	//
	// An example of this syntax can be found in the Windows LMHOSTS file, as specified
	// in [MS-NBTE] section 2.2.3 LMHOSTS File Syntax. See [LMHOSTS] for more information.
	DataFilePath string `idl:"name:pDataFilePath;string;pointer:unique" json:"data_file_path"`
	// fDel: Value specifying whether or not to delete the file specified by pDataFilePath
	// from the target WINS server. A non-zero value deletes the file from the target WINS
	// server after the database initialization is complete.
	Delete uint32 `idl:"name:fDel" json:"delete"`
}

func (o *DoStaticInitRequest) xxx_ToOp(ctx context.Context) *xxx_DoStaticInitOperation {
	if o == nil {
		return &xxx_DoStaticInitOperation{}
	}
	return &xxx_DoStaticInitOperation{
		DataFilePath: o.DataFilePath,
		Delete:       o.Delete,
	}
}

func (o *DoStaticInitRequest) xxx_FromOp(ctx context.Context, op *xxx_DoStaticInitOperation) {
	if o == nil {
		return
	}
	o.DataFilePath = op.DataFilePath
	o.Delete = op.Delete
}
func (o *DoStaticInitRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DoStaticInitRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DoStaticInitOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DoStaticInitResponse structure represents the R_WinsDoStaticInit operation response
type DoStaticInitResponse struct {
	// Return: The R_WinsDoStaticInit return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *DoStaticInitResponse) xxx_ToOp(ctx context.Context) *xxx_DoStaticInitOperation {
	if o == nil {
		return &xxx_DoStaticInitOperation{}
	}
	return &xxx_DoStaticInitOperation{
		Return: o.Return,
	}
}

func (o *DoStaticInitResponse) xxx_FromOp(ctx context.Context, op *xxx_DoStaticInitOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DoStaticInitResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DoStaticInitResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DoStaticInitOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DoScavengingOperation structure represents the R_WinsDoScavenging operation
type xxx_DoScavengingOperation struct {
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_DoScavengingOperation) OpNum() int { return 4 }

func (o *xxx_DoScavengingOperation) OpName() string { return "/winsif/v1/R_WinsDoScavenging" }

func (o *xxx_DoScavengingOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoScavengingOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_DoScavengingOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_DoScavengingOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoScavengingOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoScavengingOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DoScavengingRequest structure represents the R_WinsDoScavenging operation request
type DoScavengingRequest struct {
}

func (o *DoScavengingRequest) xxx_ToOp(ctx context.Context) *xxx_DoScavengingOperation {
	if o == nil {
		return &xxx_DoScavengingOperation{}
	}
	return &xxx_DoScavengingOperation{}
}

func (o *DoScavengingRequest) xxx_FromOp(ctx context.Context, op *xxx_DoScavengingOperation) {
	if o == nil {
		return
	}
}
func (o *DoScavengingRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DoScavengingRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DoScavengingOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DoScavengingResponse structure represents the R_WinsDoScavenging operation response
type DoScavengingResponse struct {
	// Return: The R_WinsDoScavenging return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *DoScavengingResponse) xxx_ToOp(ctx context.Context) *xxx_DoScavengingOperation {
	if o == nil {
		return &xxx_DoScavengingOperation{}
	}
	return &xxx_DoScavengingOperation{
		Return: o.Return,
	}
}

func (o *DoScavengingResponse) xxx_FromOp(ctx context.Context, op *xxx_DoScavengingOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DoScavengingResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DoScavengingResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DoScavengingOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDBRecordsOperation structure represents the R_WinsGetDbRecs operation
type xxx_GetDBRecordsOperation struct {
	WINSAddr  *raiw.Addr    `idl:"name:pWinsAddr;pointer:ref" json:"wins_addr"`
	MinVersNo *raiw.VersNo  `idl:"name:MinVersNo" json:"min_vers_no"`
	MaxVersNo *raiw.VersNo  `idl:"name:MaxVersNo" json:"max_vers_no"`
	Records   *raiw.Records `idl:"name:pRecs" json:"records"`
	Return    uint32        `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDBRecordsOperation) OpNum() int { return 5 }

func (o *xxx_GetDBRecordsOperation) OpName() string { return "/winsif/v1/R_WinsGetDbRecs" }

func (o *xxx_GetDBRecordsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDBRecordsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pWinsAddr {in} (1:{pointer=ref, alias=PWINSINTF_ADDR_T}*(1))(2:{alias=WINSINTF_ADDR_T}(struct))
	{
		if o.WINSAddr != nil {
			if err := o.WINSAddr.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.Addr{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// MinVersNo {in} (1:{alias=WINSINTF_VERS_NO_T, names=LARGE_INTEGER}(struct))
	{
		if o.MinVersNo != nil {
			if err := o.MinVersNo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.VersNo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// MaxVersNo {in} (1:{alias=WINSINTF_VERS_NO_T, names=LARGE_INTEGER}(struct))
	{
		if o.MaxVersNo != nil {
			if err := o.MaxVersNo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.VersNo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetDBRecordsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pWinsAddr {in} (1:{pointer=ref, alias=PWINSINTF_ADDR_T}*(1))(2:{alias=WINSINTF_ADDR_T}(struct))
	{
		if o.WINSAddr == nil {
			o.WINSAddr = &raiw.Addr{}
		}
		if err := o.WINSAddr.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// MinVersNo {in} (1:{alias=WINSINTF_VERS_NO_T, names=LARGE_INTEGER}(struct))
	{
		if o.MinVersNo == nil {
			o.MinVersNo = &raiw.VersNo{}
		}
		if err := o.MinVersNo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// MaxVersNo {in} (1:{alias=WINSINTF_VERS_NO_T, names=LARGE_INTEGER}(struct))
	{
		if o.MaxVersNo == nil {
			o.MaxVersNo = &raiw.VersNo{}
		}
		if err := o.MaxVersNo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDBRecordsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDBRecordsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pRecs {out} (1:{alias=PWINSINTF_RECS_T}*(1))(2:{alias=WINSINTF_RECS_T}(struct))
	{
		if o.Records != nil {
			if err := o.Records.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.Records{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDBRecordsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pRecs {out} (1:{alias=PWINSINTF_RECS_T,pointer=ref}*(1))(2:{alias=WINSINTF_RECS_T}(struct))
	{
		if o.Records == nil {
			o.Records = &raiw.Records{}
		}
		if err := o.Records.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetDBRecordsRequest structure represents the R_WinsGetDbRecs operation request
type GetDBRecordsRequest struct {
	WINSAddr *raiw.Addr `idl:"name:pWinsAddr;pointer:ref" json:"wins_addr"`
	// MinVersNo: The lower bound on the version range of the records to be retrieved.
	MinVersNo *raiw.VersNo `idl:"name:MinVersNo" json:"min_vers_no"`
	// MaxVersNo: The upper bound on the version range of the records to be retrieved.
	MaxVersNo *raiw.VersNo `idl:"name:MaxVersNo" json:"max_vers_no"`
}

func (o *GetDBRecordsRequest) xxx_ToOp(ctx context.Context) *xxx_GetDBRecordsOperation {
	if o == nil {
		return &xxx_GetDBRecordsOperation{}
	}
	return &xxx_GetDBRecordsOperation{
		WINSAddr:  o.WINSAddr,
		MinVersNo: o.MinVersNo,
		MaxVersNo: o.MaxVersNo,
	}
}

func (o *GetDBRecordsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDBRecordsOperation) {
	if o == nil {
		return
	}
	o.WINSAddr = op.WINSAddr
	o.MinVersNo = op.MinVersNo
	o.MaxVersNo = op.MaxVersNo
}
func (o *GetDBRecordsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDBRecordsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDBRecordsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDBRecordsResponse structure represents the R_WinsGetDbRecs operation response
type GetDBRecordsResponse struct {
	// pRecs: Pointer to a structure of type WINSINTF_RECS_T, which contains the records
	// retrieved from the target WINS server.
	Records *raiw.Records `idl:"name:pRecs" json:"records"`
	// Return: The R_WinsGetDbRecs return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetDBRecordsResponse) xxx_ToOp(ctx context.Context) *xxx_GetDBRecordsOperation {
	if o == nil {
		return &xxx_GetDBRecordsOperation{}
	}
	return &xxx_GetDBRecordsOperation{
		Records: o.Records,
		Return:  o.Return,
	}
}

func (o *GetDBRecordsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDBRecordsOperation) {
	if o == nil {
		return
	}
	o.Records = op.Records
	o.Return = op.Return
}
func (o *GetDBRecordsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDBRecordsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDBRecordsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_TerminateOperation structure represents the R_WinsTerm operation
type xxx_TerminateOperation struct {
	AbruptTem int16  `idl:"name:fAbruptTem" json:"abrupt_tem"`
	Return    uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_TerminateOperation) OpNum() int { return 6 }

func (o *xxx_TerminateOperation) OpName() string { return "/winsif/v1/R_WinsTerm" }

func (o *xxx_TerminateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TerminateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// fAbruptTem {in} (1:(int16))
	{
		if err := w.WriteData(o.AbruptTem); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TerminateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// fAbruptTem {in} (1:(int16))
	{
		if err := w.ReadData(&o.AbruptTem); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TerminateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TerminateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_TerminateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// TerminateRequest structure represents the R_WinsTerm operation request
type TerminateRequest struct {
	// fAbruptTem: A value that indicates whether the WINS process terminates immediately.
	// If this value is nonzero, the service terminates immediately. Otherwise, the service
	// exits normally and frees all resources.
	AbruptTem int16 `idl:"name:fAbruptTem" json:"abrupt_tem"`
}

func (o *TerminateRequest) xxx_ToOp(ctx context.Context) *xxx_TerminateOperation {
	if o == nil {
		return &xxx_TerminateOperation{}
	}
	return &xxx_TerminateOperation{
		AbruptTem: o.AbruptTem,
	}
}

func (o *TerminateRequest) xxx_FromOp(ctx context.Context, op *xxx_TerminateOperation) {
	if o == nil {
		return
	}
	o.AbruptTem = op.AbruptTem
}
func (o *TerminateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *TerminateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_TerminateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// TerminateResponse structure represents the R_WinsTerm operation response
type TerminateResponse struct {
	// Return: The R_WinsTerm return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *TerminateResponse) xxx_ToOp(ctx context.Context) *xxx_TerminateOperation {
	if o == nil {
		return &xxx_TerminateOperation{}
	}
	return &xxx_TerminateOperation{
		Return: o.Return,
	}
}

func (o *TerminateResponse) xxx_FromOp(ctx context.Context, op *xxx_TerminateOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *TerminateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *TerminateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_TerminateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BackupOperation structure represents the R_WinsBackup operation
type xxx_BackupOperation struct {
	BackupPath  string `idl:"name:pBackupPath;string;pointer:ref" json:"backup_path"`
	Incremental int16  `idl:"name:fIncremental" json:"incremental"`
	Return      uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_BackupOperation) OpNum() int { return 7 }

func (o *xxx_BackupOperation) OpName() string { return "/winsif/v1/R_WinsBackup" }

func (o *xxx_BackupOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pBackupPath {in} (1:{string, pointer=ref, alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,string,null](uchar))
	{
		if err := ndr.WriteCharNString(ctx, w, o.BackupPath); err != nil {
			return err
		}
	}
	// fIncremental {in} (1:(int16))
	{
		if err := w.WriteData(o.Incremental); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pBackupPath {in} (1:{string, pointer=ref, alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,string,null](uchar))
	{
		if err := ndr.ReadCharNString(ctx, w, &o.BackupPath); err != nil {
			return err
		}
	}
	// fIncremental {in} (1:(int16))
	{
		if err := w.ReadData(&o.Incremental); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BackupOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BackupRequest structure represents the R_WinsBackup operation request
type BackupRequest struct {
	// pBackupPath: A pointer to a string that contains the name of the directory to which
	// to back up the database. This pointer MUST not be NULL.
	BackupPath string `idl:"name:pBackupPath;string;pointer:ref" json:"backup_path"`
	// fIncremental: A value that is ignored.
	Incremental int16 `idl:"name:fIncremental" json:"incremental"`
}

func (o *BackupRequest) xxx_ToOp(ctx context.Context) *xxx_BackupOperation {
	if o == nil {
		return &xxx_BackupOperation{}
	}
	return &xxx_BackupOperation{
		BackupPath:  o.BackupPath,
		Incremental: o.Incremental,
	}
}

func (o *BackupRequest) xxx_FromOp(ctx context.Context, op *xxx_BackupOperation) {
	if o == nil {
		return
	}
	o.BackupPath = op.BackupPath
	o.Incremental = op.Incremental
}
func (o *BackupRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *BackupRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BackupResponse structure represents the R_WinsBackup operation response
type BackupResponse struct {
	// Return: The R_WinsBackup return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BackupResponse) xxx_ToOp(ctx context.Context) *xxx_BackupOperation {
	if o == nil {
		return &xxx_BackupOperation{}
	}
	return &xxx_BackupOperation{
		Return: o.Return,
	}
}

func (o *BackupResponse) xxx_FromOp(ctx context.Context, op *xxx_BackupOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BackupResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *BackupResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BackupOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteDBRecordsOperation structure represents the R_WinsDelDbRecs operation
type xxx_DeleteDBRecordsOperation struct {
	WINSAddr  *raiw.Addr   `idl:"name:pWinsAddr;pointer:ref" json:"wins_addr"`
	MinVersNo *raiw.VersNo `idl:"name:MinVersNo" json:"min_vers_no"`
	MaxVersNo *raiw.VersNo `idl:"name:MaxVersNo" json:"max_vers_no"`
	Return    uint32       `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteDBRecordsOperation) OpNum() int { return 8 }

func (o *xxx_DeleteDBRecordsOperation) OpName() string { return "/winsif/v1/R_WinsDelDbRecs" }

func (o *xxx_DeleteDBRecordsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteDBRecordsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pWinsAddr {in} (1:{pointer=ref, alias=PWINSINTF_ADDR_T}*(1))(2:{alias=WINSINTF_ADDR_T}(struct))
	{
		if o.WINSAddr != nil {
			if err := o.WINSAddr.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.Addr{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// MinVersNo {in} (1:{alias=WINSINTF_VERS_NO_T, names=LARGE_INTEGER}(struct))
	{
		if o.MinVersNo != nil {
			if err := o.MinVersNo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.VersNo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// MaxVersNo {in} (1:{alias=WINSINTF_VERS_NO_T, names=LARGE_INTEGER}(struct))
	{
		if o.MaxVersNo != nil {
			if err := o.MaxVersNo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.VersNo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DeleteDBRecordsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pWinsAddr {in} (1:{pointer=ref, alias=PWINSINTF_ADDR_T}*(1))(2:{alias=WINSINTF_ADDR_T}(struct))
	{
		if o.WINSAddr == nil {
			o.WINSAddr = &raiw.Addr{}
		}
		if err := o.WINSAddr.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// MinVersNo {in} (1:{alias=WINSINTF_VERS_NO_T, names=LARGE_INTEGER}(struct))
	{
		if o.MinVersNo == nil {
			o.MinVersNo = &raiw.VersNo{}
		}
		if err := o.MinVersNo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// MaxVersNo {in} (1:{alias=WINSINTF_VERS_NO_T, names=LARGE_INTEGER}(struct))
	{
		if o.MaxVersNo == nil {
			o.MaxVersNo = &raiw.VersNo{}
		}
		if err := o.MaxVersNo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteDBRecordsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteDBRecordsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteDBRecordsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteDBRecordsRequest structure represents the R_WinsDelDbRecs operation request
type DeleteDBRecordsRequest struct {
	WINSAddr *raiw.Addr `idl:"name:pWinsAddr;pointer:ref" json:"wins_addr"`
	// MinVersNo: The lower bound on the version number of the records to be deleted.
	MinVersNo *raiw.VersNo `idl:"name:MinVersNo" json:"min_vers_no"`
	// MaxVersNo: The upper bound on the version number of the records to be deleted.
	MaxVersNo *raiw.VersNo `idl:"name:MaxVersNo" json:"max_vers_no"`
}

func (o *DeleteDBRecordsRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteDBRecordsOperation {
	if o == nil {
		return &xxx_DeleteDBRecordsOperation{}
	}
	return &xxx_DeleteDBRecordsOperation{
		WINSAddr:  o.WINSAddr,
		MinVersNo: o.MinVersNo,
		MaxVersNo: o.MaxVersNo,
	}
}

func (o *DeleteDBRecordsRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteDBRecordsOperation) {
	if o == nil {
		return
	}
	o.WINSAddr = op.WINSAddr
	o.MinVersNo = op.MinVersNo
	o.MaxVersNo = op.MaxVersNo
}
func (o *DeleteDBRecordsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteDBRecordsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteDBRecordsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteDBRecordsResponse structure represents the R_WinsDelDbRecs operation response
type DeleteDBRecordsResponse struct {
	// Return: The R_WinsDelDbRecs return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *DeleteDBRecordsResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteDBRecordsOperation {
	if o == nil {
		return &xxx_DeleteDBRecordsOperation{}
	}
	return &xxx_DeleteDBRecordsOperation{
		Return: o.Return,
	}
}

func (o *DeleteDBRecordsResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteDBRecordsOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DeleteDBRecordsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteDBRecordsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteDBRecordsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PullRangeOperation structure represents the R_WinsPullRange operation
type xxx_PullRangeOperation struct {
	WINSAddr  *raiw.Addr   `idl:"name:pWinsAddr;pointer:ref" json:"wins_addr"`
	OwnerAddr *raiw.Addr   `idl:"name:pOwnerAddr;pointer:ref" json:"owner_addr"`
	MinVersNo *raiw.VersNo `idl:"name:MinVersNo" json:"min_vers_no"`
	MaxVersNo *raiw.VersNo `idl:"name:MaxVersNo" json:"max_vers_no"`
	Return    uint32       `idl:"name:Return" json:"return"`
}

func (o *xxx_PullRangeOperation) OpNum() int { return 9 }

func (o *xxx_PullRangeOperation) OpName() string { return "/winsif/v1/R_WinsPullRange" }

func (o *xxx_PullRangeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PullRangeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pWinsAddr {in} (1:{pointer=ref, alias=PWINSINTF_ADDR_T}*(1))(2:{alias=WINSINTF_ADDR_T}(struct))
	{
		if o.WINSAddr != nil {
			if err := o.WINSAddr.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.Addr{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pOwnerAddr {in} (1:{pointer=ref, alias=PWINSINTF_ADDR_T}*(1))(2:{alias=WINSINTF_ADDR_T}(struct))
	{
		if o.OwnerAddr != nil {
			if err := o.OwnerAddr.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.Addr{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// MinVersNo {in} (1:{alias=WINSINTF_VERS_NO_T, names=LARGE_INTEGER}(struct))
	{
		if o.MinVersNo != nil {
			if err := o.MinVersNo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.VersNo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// MaxVersNo {in} (1:{alias=WINSINTF_VERS_NO_T, names=LARGE_INTEGER}(struct))
	{
		if o.MaxVersNo != nil {
			if err := o.MaxVersNo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.VersNo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_PullRangeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pWinsAddr {in} (1:{pointer=ref, alias=PWINSINTF_ADDR_T}*(1))(2:{alias=WINSINTF_ADDR_T}(struct))
	{
		if o.WINSAddr == nil {
			o.WINSAddr = &raiw.Addr{}
		}
		if err := o.WINSAddr.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pOwnerAddr {in} (1:{pointer=ref, alias=PWINSINTF_ADDR_T}*(1))(2:{alias=WINSINTF_ADDR_T}(struct))
	{
		if o.OwnerAddr == nil {
			o.OwnerAddr = &raiw.Addr{}
		}
		if err := o.OwnerAddr.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// MinVersNo {in} (1:{alias=WINSINTF_VERS_NO_T, names=LARGE_INTEGER}(struct))
	{
		if o.MinVersNo == nil {
			o.MinVersNo = &raiw.VersNo{}
		}
		if err := o.MinVersNo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// MaxVersNo {in} (1:{alias=WINSINTF_VERS_NO_T, names=LARGE_INTEGER}(struct))
	{
		if o.MaxVersNo == nil {
			o.MaxVersNo = &raiw.VersNo{}
		}
		if err := o.MaxVersNo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PullRangeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PullRangeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PullRangeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PullRangeRequest structure represents the R_WinsPullRange operation request
type PullRangeRequest struct {
	WINSAddr  *raiw.Addr `idl:"name:pWinsAddr;pointer:ref" json:"wins_addr"`
	OwnerAddr *raiw.Addr `idl:"name:pOwnerAddr;pointer:ref" json:"owner_addr"`
	// MinVersNo: The lower bound on the range of version numbers for the records to be
	// pulled.
	MinVersNo *raiw.VersNo `idl:"name:MinVersNo" json:"min_vers_no"`
	// MaxVersNo: The upper bound on the range of version numbers for the records to be
	// pulled.
	MaxVersNo *raiw.VersNo `idl:"name:MaxVersNo" json:"max_vers_no"`
}

func (o *PullRangeRequest) xxx_ToOp(ctx context.Context) *xxx_PullRangeOperation {
	if o == nil {
		return &xxx_PullRangeOperation{}
	}
	return &xxx_PullRangeOperation{
		WINSAddr:  o.WINSAddr,
		OwnerAddr: o.OwnerAddr,
		MinVersNo: o.MinVersNo,
		MaxVersNo: o.MaxVersNo,
	}
}

func (o *PullRangeRequest) xxx_FromOp(ctx context.Context, op *xxx_PullRangeOperation) {
	if o == nil {
		return
	}
	o.WINSAddr = op.WINSAddr
	o.OwnerAddr = op.OwnerAddr
	o.MinVersNo = op.MinVersNo
	o.MaxVersNo = op.MaxVersNo
}
func (o *PullRangeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PullRangeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PullRangeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PullRangeResponse structure represents the R_WinsPullRange operation response
type PullRangeResponse struct {
	// Return: The R_WinsPullRange return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *PullRangeResponse) xxx_ToOp(ctx context.Context) *xxx_PullRangeOperation {
	if o == nil {
		return &xxx_PullRangeOperation{}
	}
	return &xxx_PullRangeOperation{
		Return: o.Return,
	}
}

func (o *PullRangeResponse) xxx_FromOp(ctx context.Context, op *xxx_PullRangeOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *PullRangeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PullRangeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PullRangeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetPriorityClassOperation structure represents the R_WinsSetPriorityClass operation
type xxx_SetPriorityClassOperation struct {
	PriorityClass raiw.PriorityClass `idl:"name:PrCls_e" json:"priority_class"`
	Return        uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_SetPriorityClassOperation) OpNum() int { return 10 }

func (o *xxx_SetPriorityClassOperation) OpName() string { return "/winsif/v1/R_WinsSetPriorityClass" }

func (o *xxx_SetPriorityClassOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPriorityClassOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PrCls_e {in} (1:{alias=WINSINTF_PRIORITY_CLASS_E}(enum))
	{
		if err := w.WriteData(uint16(o.PriorityClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPriorityClassOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PrCls_e {in} (1:{alias=WINSINTF_PRIORITY_CLASS_E}(enum))
	{
		if err := w.ReadData((*uint16)(&o.PriorityClass)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPriorityClassOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPriorityClassOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPriorityClassOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetPriorityClassRequest structure represents the R_WinsSetPriorityClass operation request
type SetPriorityClassRequest struct {
	// PrCls_e: The priority class to be set.
	PriorityClass raiw.PriorityClass `idl:"name:PrCls_e" json:"priority_class"`
}

func (o *SetPriorityClassRequest) xxx_ToOp(ctx context.Context) *xxx_SetPriorityClassOperation {
	if o == nil {
		return &xxx_SetPriorityClassOperation{}
	}
	return &xxx_SetPriorityClassOperation{
		PriorityClass: o.PriorityClass,
	}
}

func (o *SetPriorityClassRequest) xxx_FromOp(ctx context.Context, op *xxx_SetPriorityClassOperation) {
	if o == nil {
		return
	}
	o.PriorityClass = op.PriorityClass
}
func (o *SetPriorityClassRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetPriorityClassRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPriorityClassOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetPriorityClassResponse structure represents the R_WinsSetPriorityClass operation response
type SetPriorityClassResponse struct {
	// Return: The R_WinsSetPriorityClass return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetPriorityClassResponse) xxx_ToOp(ctx context.Context) *xxx_SetPriorityClassOperation {
	if o == nil {
		return &xxx_SetPriorityClassOperation{}
	}
	return &xxx_SetPriorityClassOperation{
		Return: o.Return,
	}
}

func (o *SetPriorityClassResponse) xxx_FromOp(ctx context.Context, op *xxx_SetPriorityClassOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetPriorityClassResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetPriorityClassResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPriorityClassOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ResetCountersOperation structure represents the R_WinsResetCounters operation
type xxx_ResetCountersOperation struct {
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_ResetCountersOperation) OpNum() int { return 11 }

func (o *xxx_ResetCountersOperation) OpName() string { return "/winsif/v1/R_WinsResetCounters" }

func (o *xxx_ResetCountersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResetCountersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_ResetCountersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_ResetCountersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResetCountersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ResetCountersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ResetCountersRequest structure represents the R_WinsResetCounters operation request
type ResetCountersRequest struct {
}

func (o *ResetCountersRequest) xxx_ToOp(ctx context.Context) *xxx_ResetCountersOperation {
	if o == nil {
		return &xxx_ResetCountersOperation{}
	}
	return &xxx_ResetCountersOperation{}
}

func (o *ResetCountersRequest) xxx_FromOp(ctx context.Context, op *xxx_ResetCountersOperation) {
	if o == nil {
		return
	}
}
func (o *ResetCountersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ResetCountersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResetCountersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ResetCountersResponse structure represents the R_WinsResetCounters operation response
type ResetCountersResponse struct {
	// Return: The R_WinsResetCounters return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ResetCountersResponse) xxx_ToOp(ctx context.Context) *xxx_ResetCountersOperation {
	if o == nil {
		return &xxx_ResetCountersOperation{}
	}
	return &xxx_ResetCountersOperation{
		Return: o.Return,
	}
}

func (o *ResetCountersResponse) xxx_FromOp(ctx context.Context, op *xxx_ResetCountersOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ResetCountersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ResetCountersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ResetCountersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_WorkerThreadUpdateOperation structure represents the R_WinsWorkerThdUpd operation
type xxx_WorkerThreadUpdateOperation struct {
	NewNumberOfNbtThreads uint32 `idl:"name:NewNoOfNbtThds" json:"new_number_of_nbt_threads"`
	Return                uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_WorkerThreadUpdateOperation) OpNum() int { return 12 }

func (o *xxx_WorkerThreadUpdateOperation) OpName() string { return "/winsif/v1/R_WinsWorkerThdUpd" }

func (o *xxx_WorkerThreadUpdateOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WorkerThreadUpdateOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// NewNoOfNbtThds {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NewNumberOfNbtThreads); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WorkerThreadUpdateOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// NewNoOfNbtThds {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NewNumberOfNbtThreads); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WorkerThreadUpdateOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WorkerThreadUpdateOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_WorkerThreadUpdateOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// WorkerThreadUpdateRequest structure represents the R_WinsWorkerThdUpd operation request
type WorkerThreadUpdateRequest struct {
	// NewNoOfNbtThds: New value for the number of worker threads that have been created
	// for NetBIOS requests.
	NewNumberOfNbtThreads uint32 `idl:"name:NewNoOfNbtThds" json:"new_number_of_nbt_threads"`
}

func (o *WorkerThreadUpdateRequest) xxx_ToOp(ctx context.Context) *xxx_WorkerThreadUpdateOperation {
	if o == nil {
		return &xxx_WorkerThreadUpdateOperation{}
	}
	return &xxx_WorkerThreadUpdateOperation{
		NewNumberOfNbtThreads: o.NewNumberOfNbtThreads,
	}
}

func (o *WorkerThreadUpdateRequest) xxx_FromOp(ctx context.Context, op *xxx_WorkerThreadUpdateOperation) {
	if o == nil {
		return
	}
	o.NewNumberOfNbtThreads = op.NewNumberOfNbtThreads
}
func (o *WorkerThreadUpdateRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *WorkerThreadUpdateRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WorkerThreadUpdateOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// WorkerThreadUpdateResponse structure represents the R_WinsWorkerThdUpd operation response
type WorkerThreadUpdateResponse struct {
	// Return: The R_WinsWorkerThdUpd return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *WorkerThreadUpdateResponse) xxx_ToOp(ctx context.Context) *xxx_WorkerThreadUpdateOperation {
	if o == nil {
		return &xxx_WorkerThreadUpdateOperation{}
	}
	return &xxx_WorkerThreadUpdateOperation{
		Return: o.Return,
	}
}

func (o *WorkerThreadUpdateResponse) xxx_FromOp(ctx context.Context, op *xxx_WorkerThreadUpdateOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *WorkerThreadUpdateResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *WorkerThreadUpdateResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_WorkerThreadUpdateOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetNameAndAddrOperation structure represents the R_WinsGetNameAndAddr operation
type xxx_GetNameAndAddrOperation struct {
	WINSAddr *raiw.Addr `idl:"name:pWinsAddr;pointer:ref" json:"wins_addr"`
	UNCName  string     `idl:"name:pUncName;size_is:(80);string" json:"unc_name"`
	Return   uint32     `idl:"name:Return" json:"return"`
}

func (o *xxx_GetNameAndAddrOperation) OpNum() int { return 13 }

func (o *xxx_GetNameAndAddrOperation) OpName() string { return "/winsif/v1/R_WinsGetNameAndAddr" }

func (o *xxx_GetNameAndAddrOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNameAndAddrOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_GetNameAndAddrOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_GetNameAndAddrOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	// cannot evaluate expression 80
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNameAndAddrOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pWinsAddr {out} (1:{pointer=ref, alias=PWINSINTF_ADDR_T}*(1))(2:{alias=WINSINTF_ADDR_T}(struct))
	{
		if o.WINSAddr != nil {
			if err := o.WINSAddr.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.Addr{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// pUncName {out} (1:{string, alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=80,string,null](uchar))
	{
		dimSize1 := uint64(80)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := ndr.CharNLen(o.UNCName)
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
		_UNCName_buf := []byte(o.UNCName)
		if uint64(len(_UNCName_buf)) > sizeInfo[0]-1 {
			_UNCName_buf = _UNCName_buf[:sizeInfo[0]-1]
		}
		if o.UNCName != ndr.ZeroString {
			_UNCName_buf = append(_UNCName_buf, byte(0))
		}
		for i1 := range _UNCName_buf {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(_UNCName_buf[i1]); err != nil {
				return err
			}
		}
		for i1 := len(_UNCName_buf); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetNameAndAddrOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pWinsAddr {out} (1:{pointer=ref, alias=PWINSINTF_ADDR_T}*(1))(2:{alias=WINSINTF_ADDR_T}(struct))
	{
		if o.WINSAddr == nil {
			o.WINSAddr = &raiw.Addr{}
		}
		if err := o.WINSAddr.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// pUncName {out} (1:{string, alias=LPBYTE,pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=80,string,null](uchar))
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
		var _UNCName_buf []byte
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _UNCName_buf", sizeInfo[0])
		}
		_UNCName_buf = make([]byte, sizeInfo[0])
		for i1 := range _UNCName_buf {
			i1 := i1
			if err := w.ReadData(&_UNCName_buf[i1]); err != nil {
				return err
			}
		}
		o.UNCName = strings.TrimRight(string(_UNCName_buf), ndr.ZeroString)
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetNameAndAddrRequest structure represents the R_WinsGetNameAndAddr operation request
type GetNameAndAddrRequest struct {
}

func (o *GetNameAndAddrRequest) xxx_ToOp(ctx context.Context) *xxx_GetNameAndAddrOperation {
	if o == nil {
		return &xxx_GetNameAndAddrOperation{}
	}
	return &xxx_GetNameAndAddrOperation{}
}

func (o *GetNameAndAddrRequest) xxx_FromOp(ctx context.Context, op *xxx_GetNameAndAddrOperation) {
	if o == nil {
		return
	}
}
func (o *GetNameAndAddrRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetNameAndAddrRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNameAndAddrOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetNameAndAddrResponse structure represents the R_WinsGetNameAndAddr operation response
type GetNameAndAddrResponse struct {
	WINSAddr *raiw.Addr `idl:"name:pWinsAddr;pointer:ref" json:"wins_addr"`
	UNCName  string     `idl:"name:pUncName;size_is:(80);string" json:"unc_name"`
	// Return: The R_WinsGetNameAndAddr return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetNameAndAddrResponse) xxx_ToOp(ctx context.Context) *xxx_GetNameAndAddrOperation {
	if o == nil {
		return &xxx_GetNameAndAddrOperation{}
	}
	return &xxx_GetNameAndAddrOperation{
		WINSAddr: o.WINSAddr,
		UNCName:  o.UNCName,
		Return:   o.Return,
	}
}

func (o *GetNameAndAddrResponse) xxx_FromOp(ctx context.Context, op *xxx_GetNameAndAddrOperation) {
	if o == nil {
		return
	}
	o.WINSAddr = op.WINSAddr
	o.UNCName = op.UNCName
	o.Return = op.Return
}
func (o *GetNameAndAddrResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetNameAndAddrResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetNameAndAddrOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetBrowserNamesOldOperation structure represents the R_WinsGetBrowserNames_Old operation
type xxx_GetBrowserNamesOldOperation struct {
	Names  *raiw.BrowserNames `idl:"name:pNames" json:"names"`
	Return uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_GetBrowserNamesOldOperation) OpNum() int { return 14 }

func (o *xxx_GetBrowserNamesOldOperation) OpName() string {
	return "/winsif/v1/R_WinsGetBrowserNames_Old"
}

func (o *xxx_GetBrowserNamesOldOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBrowserNamesOldOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_GetBrowserNamesOldOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_GetBrowserNamesOldOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBrowserNamesOldOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pNames {out} (1:{alias=PWINSINTF_BROWSER_NAMES_T}*(1))(2:{alias=WINSINTF_BROWSER_NAMES_T}(struct))
	{
		if o.Names != nil {
			if err := o.Names.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.BrowserNames{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBrowserNamesOldOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pNames {out} (1:{alias=PWINSINTF_BROWSER_NAMES_T,pointer=ref}*(1))(2:{alias=WINSINTF_BROWSER_NAMES_T}(struct))
	{
		if o.Names == nil {
			o.Names = &raiw.BrowserNames{}
		}
		if err := o.Names.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetBrowserNamesOldRequest structure represents the R_WinsGetBrowserNames_Old operation request
type GetBrowserNamesOldRequest struct {
}

func (o *GetBrowserNamesOldRequest) xxx_ToOp(ctx context.Context) *xxx_GetBrowserNamesOldOperation {
	if o == nil {
		return &xxx_GetBrowserNamesOldOperation{}
	}
	return &xxx_GetBrowserNamesOldOperation{}
}

func (o *GetBrowserNamesOldRequest) xxx_FromOp(ctx context.Context, op *xxx_GetBrowserNamesOldOperation) {
	if o == nil {
		return
	}
}
func (o *GetBrowserNamesOldRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetBrowserNamesOldRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBrowserNamesOldOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetBrowserNamesOldResponse structure represents the R_WinsGetBrowserNames_Old operation response
type GetBrowserNamesOldResponse struct {
	// pNames: This field MUST be ignored.
	Names *raiw.BrowserNames `idl:"name:pNames" json:"names"`
	// Return: The R_WinsGetBrowserNames_Old return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetBrowserNamesOldResponse) xxx_ToOp(ctx context.Context) *xxx_GetBrowserNamesOldOperation {
	if o == nil {
		return &xxx_GetBrowserNamesOldOperation{}
	}
	return &xxx_GetBrowserNamesOldOperation{
		Names:  o.Names,
		Return: o.Return,
	}
}

func (o *GetBrowserNamesOldResponse) xxx_FromOp(ctx context.Context, op *xxx_GetBrowserNamesOldOperation) {
	if o == nil {
		return
	}
	o.Names = op.Names
	o.Return = op.Return
}
func (o *GetBrowserNamesOldResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetBrowserNamesOldResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBrowserNamesOldOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DeleteWINSOperation structure represents the R_WinsDeleteWins operation
type xxx_DeleteWINSOperation struct {
	WINSAddr *raiw.Addr `idl:"name:pWinsAddr;pointer:ref" json:"wins_addr"`
	Return   uint32     `idl:"name:Return" json:"return"`
}

func (o *xxx_DeleteWINSOperation) OpNum() int { return 15 }

func (o *xxx_DeleteWINSOperation) OpName() string { return "/winsif/v1/R_WinsDeleteWins" }

func (o *xxx_DeleteWINSOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteWINSOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pWinsAddr {in} (1:{pointer=ref, alias=PWINSINTF_ADDR_T}*(1))(2:{alias=WINSINTF_ADDR_T}(struct))
	{
		if o.WINSAddr != nil {
			if err := o.WINSAddr.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.Addr{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DeleteWINSOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pWinsAddr {in} (1:{pointer=ref, alias=PWINSINTF_ADDR_T}*(1))(2:{alias=WINSINTF_ADDR_T}(struct))
	{
		if o.WINSAddr == nil {
			o.WINSAddr = &raiw.Addr{}
		}
		if err := o.WINSAddr.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteWINSOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteWINSOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DeleteWINSOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DeleteWINSRequest structure represents the R_WinsDeleteWins operation request
type DeleteWINSRequest struct {
	WINSAddr *raiw.Addr `idl:"name:pWinsAddr;pointer:ref" json:"wins_addr"`
}

func (o *DeleteWINSRequest) xxx_ToOp(ctx context.Context) *xxx_DeleteWINSOperation {
	if o == nil {
		return &xxx_DeleteWINSOperation{}
	}
	return &xxx_DeleteWINSOperation{
		WINSAddr: o.WINSAddr,
	}
}

func (o *DeleteWINSRequest) xxx_FromOp(ctx context.Context, op *xxx_DeleteWINSOperation) {
	if o == nil {
		return
	}
	o.WINSAddr = op.WINSAddr
}
func (o *DeleteWINSRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DeleteWINSRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteWINSOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DeleteWINSResponse structure represents the R_WinsDeleteWins operation response
type DeleteWINSResponse struct {
	// Return: The R_WinsDeleteWins return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *DeleteWINSResponse) xxx_ToOp(ctx context.Context) *xxx_DeleteWINSOperation {
	if o == nil {
		return &xxx_DeleteWINSOperation{}
	}
	return &xxx_DeleteWINSOperation{
		Return: o.Return,
	}
}

func (o *DeleteWINSResponse) xxx_FromOp(ctx context.Context, op *xxx_DeleteWINSOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DeleteWINSResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DeleteWINSResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DeleteWINSOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetFlagsOperation structure represents the R_WinsSetFlags operation
type xxx_SetFlagsOperation struct {
	Flags  uint32 `idl:"name:fFlags" json:"flags"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SetFlagsOperation) OpNum() int { return 16 }

func (o *xxx_SetFlagsOperation) OpName() string { return "/winsif/v1/R_WinsSetFlags" }

func (o *xxx_SetFlagsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlagsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// fFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlagsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// fFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlagsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlagsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetFlagsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetFlagsRequest structure represents the R_WinsSetFlags operation request
type SetFlagsRequest struct {
	// fFlags: This field MUST be ignored.
	Flags uint32 `idl:"name:fFlags" json:"flags"`
}

func (o *SetFlagsRequest) xxx_ToOp(ctx context.Context) *xxx_SetFlagsOperation {
	if o == nil {
		return &xxx_SetFlagsOperation{}
	}
	return &xxx_SetFlagsOperation{
		Flags: o.Flags,
	}
}

func (o *SetFlagsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetFlagsOperation) {
	if o == nil {
		return
	}
	o.Flags = op.Flags
}
func (o *SetFlagsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *SetFlagsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFlagsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetFlagsResponse structure represents the R_WinsSetFlags operation response
type SetFlagsResponse struct {
	// Return: The R_WinsSetFlags return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetFlagsResponse) xxx_ToOp(ctx context.Context) *xxx_SetFlagsOperation {
	if o == nil {
		return &xxx_SetFlagsOperation{}
	}
	return &xxx_SetFlagsOperation{
		Return: o.Return,
	}
}

func (o *SetFlagsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetFlagsOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetFlagsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *SetFlagsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetFlagsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetBrowserNamesOperation structure represents the R_WinsGetBrowserNames operation
type xxx_GetBrowserNamesOperation struct {
	Server *raiw.BindData     `idl:"name:ServerHdl;pointer:ref" json:"server"`
	Names  *raiw.BrowserNames `idl:"name:pNames" json:"names"`
	Return uint32             `idl:"name:Return" json:"return"`
}

func (o *xxx_GetBrowserNamesOperation) OpNum() int { return 17 }

func (o *xxx_GetBrowserNamesOperation) OpName() string { return "/winsif/v1/R_WinsGetBrowserNames" }

func (o *xxx_GetBrowserNamesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBrowserNamesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerHdl {in} (1:{handle, pointer=ref, alias=WINSIF_HANDLE, names=PWINSINTF_BIND_DATA_T}*(1))(2:{alias=WINSINTF_BIND_DATA_T}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.BindData{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBrowserNamesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerHdl {in} (1:{handle, pointer=ref, alias=WINSIF_HANDLE, names=PWINSINTF_BIND_DATA_T}*(1))(2:{alias=WINSINTF_BIND_DATA_T}(struct))
	{
		if o.Server == nil {
			o.Server = &raiw.BindData{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBrowserNamesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBrowserNamesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pNames {out} (1:{alias=PWINSINTF_BROWSER_NAMES_T}*(1))(2:{alias=WINSINTF_BROWSER_NAMES_T}(struct))
	{
		if o.Names != nil {
			if err := o.Names.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.BrowserNames{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetBrowserNamesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pNames {out} (1:{alias=PWINSINTF_BROWSER_NAMES_T,pointer=ref}*(1))(2:{alias=WINSINTF_BROWSER_NAMES_T}(struct))
	{
		if o.Names == nil {
			o.Names = &raiw.BrowserNames{}
		}
		if err := o.Names.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetBrowserNamesRequest structure represents the R_WinsGetBrowserNames operation request
type GetBrowserNamesRequest struct {
	// ServerHdl: An RPC binding over IP address/HostName to the WINS server. RPC uses this
	// binding internally to determine which WINS server the call is directed to. See [MSDN-Handles]
	// for more information.
	Server *raiw.BindData `idl:"name:ServerHdl;pointer:ref" json:"server"`
}

func (o *GetBrowserNamesRequest) xxx_ToOp(ctx context.Context) *xxx_GetBrowserNamesOperation {
	if o == nil {
		return &xxx_GetBrowserNamesOperation{}
	}
	return &xxx_GetBrowserNamesOperation{
		Server: o.Server,
	}
}

func (o *GetBrowserNamesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetBrowserNamesOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
}
func (o *GetBrowserNamesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetBrowserNamesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBrowserNamesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetBrowserNamesResponse structure represents the R_WinsGetBrowserNames operation response
type GetBrowserNamesResponse struct {
	// pNames: A pointer to a structure of type WINSINTF_BROWSER_NAMES_T (section 2.2.2.10),
	// which contains the browser name records retrieved from the target WINS server.
	Names *raiw.BrowserNames `idl:"name:pNames" json:"names"`
	// Return: The R_WinsGetBrowserNames return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetBrowserNamesResponse) xxx_ToOp(ctx context.Context) *xxx_GetBrowserNamesOperation {
	if o == nil {
		return &xxx_GetBrowserNamesOperation{}
	}
	return &xxx_GetBrowserNamesOperation{
		Names:  o.Names,
		Return: o.Return,
	}
}

func (o *GetBrowserNamesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetBrowserNamesOperation) {
	if o == nil {
		return
	}
	o.Names = op.Names
	o.Return = op.Return
}
func (o *GetBrowserNamesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetBrowserNamesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetBrowserNamesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDBRecordsByNameOperation structure represents the R_WinsGetDbRecsByName operation
type xxx_GetDBRecordsByNameOperation struct {
	WINSAddr               *raiw.Addr    `idl:"name:pWinsAddr;pointer:unique" json:"wins_addr"`
	Location               uint32        `idl:"name:Location" json:"location"`
	Name                   []byte        `idl:"name:pName;size_is:((NameLen+1));pointer:unique" json:"name"`
	NameLength             uint32        `idl:"name:NameLen" json:"name_length"`
	NumberOfRecordsDesired uint32        `idl:"name:NoOfRecsDesired" json:"number_of_records_desired"`
	OnlyStatic             uint32        `idl:"name:fOnlyStatic" json:"only_static"`
	Records                *raiw.Records `idl:"name:pRecs" json:"records"`
	Return                 uint32        `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDBRecordsByNameOperation) OpNum() int { return 18 }

func (o *xxx_GetDBRecordsByNameOperation) OpName() string { return "/winsif/v1/R_WinsGetDbRecsByName" }

func (o *xxx_GetDBRecordsByNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Name != nil && o.NameLength == 0 {
		o.NameLength = uint32((len(o.Name) - 1))
		if len(o.Name) < 1 {
			o.NameLength = uint32(0)
		}
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDBRecordsByNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pWinsAddr {in} (1:{pointer=unique, alias=PWINSINTF_ADDR_T}*(1))(2:{alias=WINSINTF_ADDR_T}(struct))
	{
		if o.WINSAddr != nil {
			_ptr_pWinsAddr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.WINSAddr != nil {
					if err := o.WINSAddr.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&raiw.Addr{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.WINSAddr, _ptr_pWinsAddr); err != nil {
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
	// Location {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Location); err != nil {
			return err
		}
	}
	// pName {in} (1:{pointer=unique, alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=(NameLen+1)](uchar))
	{
		if o.Name != nil || (o.NameLength+1) > 0 {
			_ptr_pName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64((o.NameLength + 1))
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Name {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Name[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Name); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Name, _ptr_pName); err != nil {
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
	// NameLen {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NameLength); err != nil {
			return err
		}
	}
	// NoOfRecsDesired {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NumberOfRecordsDesired); err != nil {
			return err
		}
	}
	// fOnlyStatic {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.OnlyStatic); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDBRecordsByNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pWinsAddr {in} (1:{pointer=unique, alias=PWINSINTF_ADDR_T}*(1))(2:{alias=WINSINTF_ADDR_T}(struct))
	{
		_ptr_pWinsAddr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.WINSAddr == nil {
				o.WINSAddr = &raiw.Addr{}
			}
			if err := o.WINSAddr.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pWinsAddr := func(ptr interface{}) { o.WINSAddr = *ptr.(**raiw.Addr) }
		if err := w.ReadPointer(&o.WINSAddr, _s_pWinsAddr, _ptr_pWinsAddr); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Location {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Location); err != nil {
			return err
		}
	}
	// pName {in} (1:{pointer=unique, alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=(NameLen+1)](uchar))
	{
		_ptr_pName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Name", sizeInfo[0])
			}
			o.Name = make([]byte, sizeInfo[0])
			for i1 := range o.Name {
				i1 := i1
				if err := w.ReadData(&o.Name[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_pName := func(ptr interface{}) { o.Name = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Name, _s_pName, _ptr_pName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// NameLen {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NameLength); err != nil {
			return err
		}
	}
	// NoOfRecsDesired {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NumberOfRecordsDesired); err != nil {
			return err
		}
	}
	// fOnlyStatic {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.OnlyStatic); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDBRecordsByNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDBRecordsByNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pRecs {out} (1:{alias=PWINSINTF_RECS_T}*(1))(2:{alias=WINSINTF_RECS_T}(struct))
	{
		if o.Records != nil {
			if err := o.Records.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.Records{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDBRecordsByNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pRecs {out} (1:{alias=PWINSINTF_RECS_T,pointer=ref}*(1))(2:{alias=WINSINTF_RECS_T}(struct))
	{
		if o.Records == nil {
			o.Records = &raiw.Records{}
		}
		if err := o.Records.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetDBRecordsByNameRequest structure represents the R_WinsGetDbRecsByName operation request
type GetDBRecordsByNameRequest struct {
	WINSAddr *raiw.Addr `idl:"name:pWinsAddr;pointer:unique" json:"wins_addr"`
	// Location: A value specifying the direction in which the database is searched. If
	// the value is zero, the database is searched forward starting from the beginning.
	// If the value is 1, the database is searched backward starting from the last record
	// of the database.
	Location uint32 `idl:"name:Location" json:"location"`
	// pName: A pointer to a name that specifies the cursor from which the database retrieval
	// starts.
	Name []byte `idl:"name:pName;size_is:((NameLen+1));pointer:unique" json:"name"`
	// NameLen: The length of the name that pName points to, including terminating NULL
	// character.
	NameLength uint32 `idl:"name:NameLen" json:"name_length"`
	// NoOfRecsDesired: The number of records to be retrieved from the database.
	NumberOfRecordsDesired uint32 `idl:"name:NoOfRecsDesired" json:"number_of_records_desired"`
	// fOnlyStatic: Takes a value of 1, 2, or 4 to indicate whether static records, dynamic
	// records, or both are retrieved. A value of 1 retrieves only static records. A value
	// of 2 retrieves only dynamic records. A value of 4 retrieves both static records and
	// dynamic records.
	OnlyStatic uint32 `idl:"name:fOnlyStatic" json:"only_static"`
}

func (o *GetDBRecordsByNameRequest) xxx_ToOp(ctx context.Context) *xxx_GetDBRecordsByNameOperation {
	if o == nil {
		return &xxx_GetDBRecordsByNameOperation{}
	}
	return &xxx_GetDBRecordsByNameOperation{
		WINSAddr:               o.WINSAddr,
		Location:               o.Location,
		Name:                   o.Name,
		NameLength:             o.NameLength,
		NumberOfRecordsDesired: o.NumberOfRecordsDesired,
		OnlyStatic:             o.OnlyStatic,
	}
}

func (o *GetDBRecordsByNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDBRecordsByNameOperation) {
	if o == nil {
		return
	}
	o.WINSAddr = op.WINSAddr
	o.Location = op.Location
	o.Name = op.Name
	o.NameLength = op.NameLength
	o.NumberOfRecordsDesired = op.NumberOfRecordsDesired
	o.OnlyStatic = op.OnlyStatic
}
func (o *GetDBRecordsByNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetDBRecordsByNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDBRecordsByNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDBRecordsByNameResponse structure represents the R_WinsGetDbRecsByName operation response
type GetDBRecordsByNameResponse struct {
	// pRecs: A pointer to a structure containing the retrieved records.
	Records *raiw.Records `idl:"name:pRecs" json:"records"`
	// Return: The R_WinsGetDbRecsByName return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetDBRecordsByNameResponse) xxx_ToOp(ctx context.Context) *xxx_GetDBRecordsByNameOperation {
	if o == nil {
		return &xxx_GetDBRecordsByNameOperation{}
	}
	return &xxx_GetDBRecordsByNameOperation{
		Records: o.Records,
		Return:  o.Return,
	}
}

func (o *GetDBRecordsByNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDBRecordsByNameOperation) {
	if o == nil {
		return
	}
	o.Records = op.Records
	o.Return = op.Return
}
func (o *GetDBRecordsByNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetDBRecordsByNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDBRecordsByNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StatusNewOperation structure represents the R_WinsStatusNew operation
type xxx_StatusNewOperation struct {
	Cmd     raiw.Cmd         `idl:"name:Cmd_e" json:"cmd"`
	Results *raiw.ResultsNew `idl:"name:pResults" json:"results"`
	Return  uint32           `idl:"name:Return" json:"return"`
}

func (o *xxx_StatusNewOperation) OpNum() int { return 19 }

func (o *xxx_StatusNewOperation) OpName() string { return "/winsif/v1/R_WinsStatusNew" }

func (o *xxx_StatusNewOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatusNewOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Cmd_e {in} (1:{alias=WINSINTF_CMD_E}(enum))
	{
		if err := w.WriteData(uint16(o.Cmd)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatusNewOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Cmd_e {in} (1:{alias=WINSINTF_CMD_E}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Cmd)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatusNewOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatusNewOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResults {out} (1:{alias=PWINSINTF_RESULTS_NEW_T}*(1))(2:{alias=WINSINTF_RESULTS_NEW_T}(struct))
	{
		if o.Results != nil {
			if err := o.Results.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.ResultsNew{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatusNewOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResults {out} (1:{alias=PWINSINTF_RESULTS_NEW_T,pointer=ref}*(1))(2:{alias=WINSINTF_RESULTS_NEW_T}(struct))
	{
		if o.Results == nil {
			o.Results = &raiw.ResultsNew{}
		}
		if err := o.Results.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// StatusNewRequest structure represents the R_WinsStatusNew operation request
type StatusNewRequest struct {
	// Cmd_e: The command to be executed on the target WINS server, from the WINSINTF_CMD_E
	// enumeration (section 2.2.1.5).
	Cmd raiw.Cmd `idl:"name:Cmd_e" json:"cmd"`
}

func (o *StatusNewRequest) xxx_ToOp(ctx context.Context) *xxx_StatusNewOperation {
	if o == nil {
		return &xxx_StatusNewOperation{}
	}
	return &xxx_StatusNewOperation{
		Cmd: o.Cmd,
	}
}

func (o *StatusNewRequest) xxx_FromOp(ctx context.Context, op *xxx_StatusNewOperation) {
	if o == nil {
		return
	}
	o.Cmd = op.Cmd
}
func (o *StatusNewRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *StatusNewRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StatusNewOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StatusNewResponse structure represents the R_WinsStatusNew operation response
type StatusNewResponse struct {
	// pResults: A pointer to a WINSINTF_RESULTS_NEW_T structure (section 2.2.2.11), which
	// contains the results of the command execution.
	Results *raiw.ResultsNew `idl:"name:pResults" json:"results"`
	// Return: The R_WinsStatusNew return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *StatusNewResponse) xxx_ToOp(ctx context.Context) *xxx_StatusNewOperation {
	if o == nil {
		return &xxx_StatusNewOperation{}
	}
	return &xxx_StatusNewOperation{
		Results: o.Results,
		Return:  o.Return,
	}
}

func (o *StatusNewResponse) xxx_FromOp(ctx context.Context, op *xxx_StatusNewOperation) {
	if o == nil {
		return
	}
	o.Results = op.Results
	o.Return = op.Return
}
func (o *StatusNewResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *StatusNewResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StatusNewOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StatusWOperation structure represents the R_WinsStatusWHdl operation
type xxx_StatusWOperation struct {
	Server  *raiw.BindData   `idl:"name:ServerHdl;pointer:ref" json:"server"`
	Cmd     raiw.Cmd         `idl:"name:Cmd_e" json:"cmd"`
	Results *raiw.ResultsNew `idl:"name:pResults;pointer:ref" json:"results"`
	Return  uint32           `idl:"name:Return" json:"return"`
}

func (o *xxx_StatusWOperation) OpNum() int { return 20 }

func (o *xxx_StatusWOperation) OpName() string { return "/winsif/v1/R_WinsStatusWHdl" }

func (o *xxx_StatusWOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatusWOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerHdl {in} (1:{handle, pointer=ref, alias=WINSIF_HANDLE, names=PWINSINTF_BIND_DATA_T}*(1))(2:{alias=WINSINTF_BIND_DATA_T}(struct))
	{
		if o.Server != nil {
			if err := o.Server.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.BindData{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Cmd_e {in} (1:{alias=WINSINTF_CMD_E}(enum))
	{
		if err := w.WriteData(uint16(o.Cmd)); err != nil {
			return err
		}
	}
	// pResults {in, out} (1:{pointer=ref, alias=PWINSINTF_RESULTS_NEW_T}*(1))(2:{alias=WINSINTF_RESULTS_NEW_T}(struct))
	{
		if o.Results != nil {
			if err := o.Results.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.ResultsNew{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatusWOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerHdl {in} (1:{handle, pointer=ref, alias=WINSIF_HANDLE, names=PWINSINTF_BIND_DATA_T}*(1))(2:{alias=WINSINTF_BIND_DATA_T}(struct))
	{
		if o.Server == nil {
			o.Server = &raiw.BindData{}
		}
		if err := o.Server.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Cmd_e {in} (1:{alias=WINSINTF_CMD_E}(enum))
	{
		if err := w.ReadData((*uint16)(&o.Cmd)); err != nil {
			return err
		}
	}
	// pResults {in, out} (1:{pointer=ref, alias=PWINSINTF_RESULTS_NEW_T}*(1))(2:{alias=WINSINTF_RESULTS_NEW_T}(struct))
	{
		if o.Results == nil {
			o.Results = &raiw.ResultsNew{}
		}
		if err := o.Results.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatusWOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatusWOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pResults {in, out} (1:{pointer=ref, alias=PWINSINTF_RESULTS_NEW_T}*(1))(2:{alias=WINSINTF_RESULTS_NEW_T}(struct))
	{
		if o.Results != nil {
			if err := o.Results.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.ResultsNew{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StatusWOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pResults {in, out} (1:{pointer=ref, alias=PWINSINTF_RESULTS_NEW_T}*(1))(2:{alias=WINSINTF_RESULTS_NEW_T}(struct))
	{
		if o.Results == nil {
			o.Results = &raiw.ResultsNew{}
		}
		if err := o.Results.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// StatusWRequest structure represents the R_WinsStatusWHdl operation request
type StatusWRequest struct {
	// ServerHdl: An RPC binding over IP address/HostName to the WINS server. RPC uses this
	// binding internally to determine which WINS server the call is directed to. See [MSDN-Handles]
	// for more information.
	Server *raiw.BindData `idl:"name:ServerHdl;pointer:ref" json:"server"`
	// Cmd_e: The command to be executed on the target WINS server.
	Cmd raiw.Cmd `idl:"name:Cmd_e" json:"cmd"`
	// pResults: A pointer to a structure of type WINSINTF_RESULTS_NEW_T (section 2.2.2.11)
	// that contains the results of the command execution.
	Results *raiw.ResultsNew `idl:"name:pResults;pointer:ref" json:"results"`
}

func (o *StatusWRequest) xxx_ToOp(ctx context.Context) *xxx_StatusWOperation {
	if o == nil {
		return &xxx_StatusWOperation{}
	}
	return &xxx_StatusWOperation{
		Server:  o.Server,
		Cmd:     o.Cmd,
		Results: o.Results,
	}
}

func (o *StatusWRequest) xxx_FromOp(ctx context.Context, op *xxx_StatusWOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.Cmd = op.Cmd
	o.Results = op.Results
}
func (o *StatusWRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *StatusWRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StatusWOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StatusWResponse structure represents the R_WinsStatusWHdl operation response
type StatusWResponse struct {
	// pResults: A pointer to a structure of type WINSINTF_RESULTS_NEW_T (section 2.2.2.11)
	// that contains the results of the command execution.
	Results *raiw.ResultsNew `idl:"name:pResults;pointer:ref" json:"results"`
	// Return: The R_WinsStatusWHdl return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *StatusWResponse) xxx_ToOp(ctx context.Context) *xxx_StatusWOperation {
	if o == nil {
		return &xxx_StatusWOperation{}
	}
	return &xxx_StatusWOperation{
		Results: o.Results,
		Return:  o.Return,
	}
}

func (o *StatusWResponse) xxx_FromOp(ctx context.Context, op *xxx_StatusWOperation) {
	if o == nil {
		return
	}
	o.Results = op.Results
	o.Return = op.Return
}
func (o *StatusWResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *StatusWResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StatusWOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_DoScavengingNewOperation structure represents the R_WinsDoScavengingNew operation
type xxx_DoScavengingNewOperation struct {
	PeriodicScavengingRequest *raiw.ScavengingRequest `idl:"name:pScvReq;pointer:ref" json:"periodic_scavenging_request"`
	Return                    uint32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_DoScavengingNewOperation) OpNum() int { return 21 }

func (o *xxx_DoScavengingNewOperation) OpName() string { return "/winsif/v1/R_WinsDoScavengingNew" }

func (o *xxx_DoScavengingNewOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoScavengingNewOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pScvReq {in} (1:{pointer=ref, alias=PWINSINTF_SCV_REQ_T}*(1))(2:{alias=WINSINTF_SCV_REQ_T}(struct))
	{
		if o.PeriodicScavengingRequest != nil {
			if err := o.PeriodicScavengingRequest.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&raiw.ScavengingRequest{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_DoScavengingNewOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pScvReq {in} (1:{pointer=ref, alias=PWINSINTF_SCV_REQ_T}*(1))(2:{alias=WINSINTF_SCV_REQ_T}(struct))
	{
		if o.PeriodicScavengingRequest == nil {
			o.PeriodicScavengingRequest = &raiw.ScavengingRequest{}
		}
		if err := o.PeriodicScavengingRequest.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoScavengingNewOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoScavengingNewOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_DoScavengingNewOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// DoScavengingNewRequest structure represents the R_WinsDoScavengingNew operation request
type DoScavengingNewRequest struct {
	// pScvReq: A pointer to a WINSINTF_SCV_REQ_T structure (section 2.2.2.12) that defines
	// the type of scavenging operation.
	PeriodicScavengingRequest *raiw.ScavengingRequest `idl:"name:pScvReq;pointer:ref" json:"periodic_scavenging_request"`
}

func (o *DoScavengingNewRequest) xxx_ToOp(ctx context.Context) *xxx_DoScavengingNewOperation {
	if o == nil {
		return &xxx_DoScavengingNewOperation{}
	}
	return &xxx_DoScavengingNewOperation{
		PeriodicScavengingRequest: o.PeriodicScavengingRequest,
	}
}

func (o *DoScavengingNewRequest) xxx_FromOp(ctx context.Context, op *xxx_DoScavengingNewOperation) {
	if o == nil {
		return
	}
	o.PeriodicScavengingRequest = op.PeriodicScavengingRequest
}
func (o *DoScavengingNewRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *DoScavengingNewRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DoScavengingNewOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// DoScavengingNewResponse structure represents the R_WinsDoScavengingNew operation response
type DoScavengingNewResponse struct {
	// Return: The R_WinsDoScavengingNew return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *DoScavengingNewResponse) xxx_ToOp(ctx context.Context) *xxx_DoScavengingNewOperation {
	if o == nil {
		return &xxx_DoScavengingNewOperation{}
	}
	return &xxx_DoScavengingNewOperation{
		Return: o.Return,
	}
}

func (o *DoScavengingNewResponse) xxx_FromOp(ctx context.Context, op *xxx_DoScavengingNewOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *DoScavengingNewResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *DoScavengingNewResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_DoScavengingNewOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
