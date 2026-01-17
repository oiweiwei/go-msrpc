package fax

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

// fax server interface.
type FaxServer interface {

	// The fax client application calls the FAX_GetServicePrinters (Opnum 0) method to obtain
	// a list of printers that are visible to the fax server.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the                |
	//	|                                    | FAX_ACCESS_QUERY_CONFIG access rights required for this operation.               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | Not enough storage is available to process this command. The fax server failed   |
	//	|                                    | to allocate sufficient memory to hold the array of FAX_PRINTER_INFOW structures  |
	//	|                                    | to be returned to the client.                                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. The lpBuffer parameter is set to a NULL pointer      |
	//	|                                    | value.<136>                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F    | The fax server failed to custom marshal the array of FAX_PRINTER_INFOW           |
	//	|                                    | structures to be returned to the client.                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetServicePrinters(context.Context, *GetServicePrintersRequest) (*GetServicePrintersResponse, error)

	// The FAX_ConnectionRefCount (Opnum 1) method is called by the client.<73>
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The Connect parameter is set to a value of 0x00000001          |
	//	|                                    | (Connect), and the client's fax user account does not have the                   |
	//	|                                    | ALL_FAX_USER_ACCESS_RIGHTS access rights required for the connect operation.     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned if any of the following  |
	//	|                                    | conditions are met: § The Connect parameter is set to a value of 0x00000000      |
	//	|                                    | (Disconnect) or to a value of 0x00000002 (Release), and the Handle parameter is  |
	//	|                                    | set to a NULL value.<75> § The Connect parameter is set to a value of 0x00000001 |
	//	|                                    | (Connect), and the CanShare parameter is set to a NULL pointer value.<76> §      |
	//	|                                    | The Connect parameter is set to a value other than 0x00000000 (Disconnect),      |
	//	|                                    | 0x00000001 (Connect), or 0x00000002 (Release).                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Fax clients call this method to connect or disconnect from the fax server.
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	ConnectionReferenceCount(context.Context, *ConnectionReferenceCountRequest) (*ConnectionReferenceCountResponse, error)

	// The FAX_OpenPort (Opnum 2) method is called by the client. In response, the server
	// opens a fax port for subsequent use in other fax methods, and it returns a fax port
	// handle for use by the fax client application.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied: the client's fax user account does not have either the         |
	//	|                                    | FAX_ACCESS_QUERY_CONFIG or the FAX_ACCESS_MANAGE_CONFIG access permissions.      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_HANDLE 0x00000006    | The call was made with the Flags argument containing the PORT_OPEN_MODIFY flag   |
	//	|                                    | and the port is already opened to be modified by another call.                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BAD_UNIT 0x00000014          | The system cannot find the port for the receiving device by using the line       |
	//	|                                    | identifier specified by the DeviceId argument.                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The FaxPortHandle argument is NULL.<140>                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	OpenPort(context.Context, *OpenPortRequest) (*OpenPortResponse, error)

	// The FAX_ClosePort (Opnum 3) method is called by the client to close an open fax port.
	// The client passes FaxPortHandle, which it received from a call to FAX_OpenPort (section
	// 3.1.4.1.65).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_HANDLE 0x00000006    | FaxPortHandle is not a valid open port handle returned by FAX_OpenPort. <71>     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The RPC_FAX_PORT_HANDLE fax data type referenced by the FaxPortHandle parameter  |
	//	|                                    | is set to a NULL pointer value.<72>                                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	ClosePort(context.Context, *ClosePortRequest) (*ClosePortResponse, error)

	// The FAX_EnumJobs (Opnum 4) method is called by the client to enumerate all the fax
	// jobs on the specified fax server.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the                |
	//	|                                    | FAX_ACCESS_QUERY_JOBS access right.                                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The server cannot allocate sufficient memory to hold the array of FAX_JOB_ENTRY  |
	//	|                                    | (section 2.2.5) structures to be returned to the client.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. The Buffer parameter is set to a NULL pointer        |
	//	|                                    | value.<87>                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F    | The server failed to custom marshal the array of FAX_JOB_ENTRY to be returned to |
	//	|                                    | the client.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EnumJobs(context.Context, *EnumJobsRequest) (*EnumJobsResponse, error)

	// The FAX_GetJob (Opnum 5) method is called by the client to retrieve information regarding
	// a specific job. The job is specified by the JobId parameter. The value for the JobId
	// parameter can be obtained using one of the following methods: FAX_EnumJobs (section
	// 3.1.4.1.21), FAX_EnumJobsEx (section 3.1.4.1.22), or FAX_EnumJobsEx2 (section 3.1.4.1.23).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | The client's fax user account does not have access to query the job specified    |
	//	|                                    | by the JobId parameter. This error can happen in any of the following cases: §   |
	//	|                                    | The job specified by the JobId parameter is an outgoing job; the client (the fax |
	//	|                                    | user account currently logged in on the client) is not the owner of this job and |
	//	|                                    | does not have the FAX_ACCESS_QUERY_OUT_JOBS access rights. § The job specified   |
	//	|                                    | by the JobId parameter is an incoming job; incoming faxes are not public and     |
	//	|                                    | the client's fax user account does not have the FAX_ACCESS_MANAGE_RECEIVE_FOLDER |
	//	|                                    | rights.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The server cannot allocate sufficient memory to hold the FAX_JOB_ENTRY (section  |
	//	|                                    | 2.2.5) structure to be returned to the client.                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned if any of the following  |
	//	|                                    | conditions are met: § The Buffer parameter is set to a NULL pointer value.<110>  |
	//	|                                    | § The fax job specified by the JobId parameter cannot be found (does not         |
	//	|                                    | exist) in the fax server queue. § The fax job specified by JobId cannot be       |
	//	|                                    | queried: the job type is JT_BROADCAST, JS_DELETING, or JS_COMPLETED. For more    |
	//	|                                    | information about job types, see the description of the dwJobType member of the  |
	//	|                                    | FAX_JOB_STATUS (section 2.2.36) structure.                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F    | The server failed to custom marshal the FAX_JOB_ENTRY to be returned to the      |
	//	|                                    | client.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetJob(context.Context, *GetJobRequest) (*GetJobResponse, error)

	// The FAX_SetJob (Opnum 6) method is called by the client. The value for the JobId
	// parameter can be obtained using one of the following methods: FAX_EnumJobs (section
	// 3.1.4.1.21), FAX_EnumJobsEx (section 3.1.4.1.22), or FAX_EnumJobsEx2 (section 3.1.4.1.23).
	//
	// On success, the server MUST pause, resume, cancel, or restart the specified fax job
	// and MUST set the job status (section 3.1.1) to reflect the new job state.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the access rights  |
	//	|                                    | required to perform this operation: § For an outgoing fax job, the client's fax  |
	//	|                                    | user account is not the owner of the fax job and the client's fax user account   |
	//	|                                    | does not have the FAX_ACCESS_MANAGE_OUT_JOBS rights. § For an incoming fax job,  |
	//	|                                    | incoming faxes are not public, and the client's fax user account does not have   |
	//	|                                    | the FAX_ACCESS_MANAGE_RECEIVE_FOLDER rights.                                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                    | following conditions: § The fax job indicated by the JobId argument cannot be    |
	//	|                                    | found by the fax server. § The specified Command argument value is JC_DELETE and |
	//	|                                    | the fax job specified by the JobId argument indicates a fax job that is already  |
	//	|                                    | in a state of being deleted or was already deleted. § The specified Command      |
	//	|                                    | argument value is JC_UNKNOWN (0). § The specified Command argument value is not  |
	//	|                                    | JC_DELETE, JC_PAUSE, or JC_RESUME/JC_RESTART. § The specified Command argument   |
	//	|                                    | value is JC_DELETE and the type of the fax job specified by the JobId parameter  |
	//	|                                    | is JT_BROADCAST See the definition of the dwJobType member of the FAX_JOB_STATUS |
	//	|                                    | structure (section 2.2.36).                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_OPERATION 0x000010DD | The specified Command argument value is JC_DELETE, and the specified JobId       |
	//	|                                    | represents a fax job with a current job status other than JS_PENDING or          |
	//	|                                    | JS_RETRYING.                                                                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	SetJob(context.Context, *SetJobRequest) (*SetJobResponse, error)

	// The FAX_GetPageData (Opnum 7) method is called by the client to retrieve data in
	// the first page of an outgoing fax job. The information that is returned in the buffer
	// is an in-memory copy of the first page of the TIFF file. The value for the JobId
	// parameter can be obtained using one of the following methods: FAX_EnumJobs (section
	// 3.1.4.1.21), FAX_EnumJobsEx (section 3.1.4.1.22), or FAX_EnumJobsEx2 (section 3.1.4.1.23).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The caller does not have the access rights required for this   |
	//	|                                    | operation: ALL_FAX_USER_ACCESS_RIGHTS (section 2.2.83).                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | Not enough storage is available to process this command.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | The job identified by the JobId parameter is not an outgoing fax job or is not   |
	//	|                                    | a valid fax job for which the fax server can extract the first page of the TIFF  |
	//	|                                    | file.                                                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                    | following conditions: § One or more of the following parameters are null         |
	//	|                                    | pointers: Buffer, ImageWidth, ImageHeight.<120> § The fax server cannot find the |
	//	|                                    | fax job indicated by JobId.                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetPageData(context.Context, *GetPageDataRequest) (*GetPageDataResponse, error)

	// The FAX_GetDeviceStatus (Opnum 8) method is called by the client to retrieve information
	// about a specified fax device (port).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the                |
	//	|                                    | FAX_ACCESS_QUERY_CONFIG access rights required for this operation.               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The server cannot allocate sufficient memory to hold the FAX_DEVICE_STATUS to be |
	//	|                                    | returned to the client.                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | The FaxPortHandle parameter is not set to a valid port handle obtained using     |
	//	|                                    | FAX_OpenPort.<106>                                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                    | following conditions: § The StatusBuffer parameter is set to a NULL pointer      |
	//	|                                    | value.<107> § The FaxPortHandle parameter is set to a NULL value.<108>           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F    | The server failed to custom marshal the FAX_DEVICE_STATUS to be returned to the  |
	//	|                                    | client.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetDeviceStatus(context.Context, *GetDeviceStatusRequest) (*GetDeviceStatusResponse, error)

	// The FAX_Abort (Opnum 9) method is called by the client to abort the specified fax
	// job on the server. The value for the JobId parameter can be obtained using one of
	// the following methods: FAX_EnumJobs (section 3.1.4.1.21), FAX_EnumJobsEx (section
	// 3.1.4.1.22), or FAX_EnumJobsEx2 (section 3.1.4.1.23).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return either one of the following error codes, one of the fax-specific errors
	// that are defined in section 2.2.52, or one of the other standard errors defined in
	// [MS-ERREF] section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. This error code is returned when any of the following          |
	//	|                                    | conditions happen: § The client's fax user account does not have                 |
	//	|                                    | FAX_ACCESS_MANAGE_RECEIVE_FOLDER permission, and the specified JobId             |
	//	|                                    | represents an incoming fax job. § The client's fax user account does not have    |
	//	|                                    | FAX_ACCESS_MANAGE_OUT_JOBS permission, and the specified JobId represents an     |
	//	|                                    | outgoing fax job of a different user.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                    | following conditions: § The fax job identified by the specified JobId is not     |
	//	|                                    | found. § The specified job has already been canceled or is in the process of     |
	//	|                                    | being canceled. § The type of the fax job specified by the JobId parameter is    |
	//	|                                    | JT_BROADCAST (the description of the dwJobType member of the FAX_JOB_STATUS      |
	//	|                                    | structure specified in section 2.2.36).                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_OPERATION 0x000010DD | The operation is invalid. This error code is returned under any of the following |
	//	|                                    | conditions: § The specified JobId represents an incoming fax job (the dwJobType  |
	//	|                                    | member of the FAX_JOB_STATUS describing the job is set to 0x0002), which is      |
	//	|                                    | being routed (the dwQueueStatus member of the FAX_JOB_STATUS describing the job  |
	//	|                                    | is set to JS_ROUTING) and cannot be aborted at this stage. § The specified JobId |
	//	|                                    | represents a fax job in progress (the dwJobType member of the FAX_JOB_STATUS     |
	//	|                                    | describing the job is set to 0x0003), which the fax server failed to route       |
	//	|                                    | (the dwQueueStatus member of the FAX_JOB_STATUS describing the job is set to     |
	//	|                                    | JS_IN_PROGRESS) and cannot be aborted at this stage.                             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	Abort(context.Context, *AbortRequest) (*AbortResponse, error)

	// The FAX_EnumPorts (Opnum 10) method is called by the client to obtain port state
	// information.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the access rights  |
	//	|                                    | (FAX_ACCESS_QUERY_CONFIG) required for this operation.                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | Not enough storage is available to process this command. The fax server cannot   |
	//	|                                    | allocate sufficient memory to hold the array of _FAX_PORT_INFO to be returned to |
	//	|                                    | the client.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned if any of the following  |
	//	|                                    | conditions are met:<95> § The PortBuffer parameter is set to a NULL pointer      |
	//	|                                    | value. § The PortsReturned parameter is set to a NULL pointer value.             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F    | The fax server failed to custom marshal the array of _FAX_PORT_INFO to be        |
	//	|                                    | returned to the client.                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EnumPorts(context.Context, *EnumPortsRequest) (*EnumPortsResponse, error)

	// The FAX_GetPort (Opnum 11) method is called by the client to retrieve port status
	// information for a requested port at the server.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the                |
	//	|                                    | FAX_ACCESS_QUERY_CONFIG access rights required for this operation.               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The server cannot allocate sufficient memory to hold the FAX_PORT_INFO (section  |
	//	|                                    | 2.2.7) data structure to be returned to the client.                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | The FaxPortHandle parameter is not set to a valid port handle obtained using     |
	//	|                                    | FAX_OpenPort.<124>                                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned if any of the following  |
	//	|                                    | conditions are met: § The PortBuffer parameter is set to a NULL pointer          |
	//	|                                    | value.<125> § The FaxPortHandle parameter is set to NULL.<126>                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F    | The server failed to custom marshal the FAX_PORT_INFO to be returned to the      |
	//	|                                    | client.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetPort(context.Context, *GetPortRequest) (*GetPortResponse, error)

	// A fax client application uses the FAX_SetPort (Opnum 12) method to set fax device
	// information. The function sets extension configuration properties that are stored
	// at the device level, such as enabling or disabling sending and receiving, and the
	// auto or manual answering of calls.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                                  |
	//	|                  VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005               | Access is denied. The client's fax user account does not have the                |
	//	|                                              | FAX_ACCESS_MANAGE_CONFIG fax access right.                                       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D                | The handle specified by the FaxPortHandle argument is not a valid fax port       |
	//	|                                              | handle obtained by a call to FAX_OpenPort.<166>                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057           | The parameter is incorrect. This error code is returned if any of the following  |
	//	|                                              | conditions are met: § The pointer specified by the PortInfo argument is NULL.    |
	//	|                                              | § The handle specified by the FaxPortHandle argument is NULL. § The requested    |
	//	|                                              | value of the Priority member of the PortInfo parameter is 0, or is greater than  |
	//	|                                              | the total number of installed fax devices. § The size of the PortInfo argument   |
	//	|                                              | specified in the SizeOfStruct field is not the correct size, in bytes, for the   |
	//	|                                              | FAX_PORT_INFO. § When FAX_ERR_DEVICE_NUM_LIMIT_EXCEEDED is to be returned, but   |
	//	|                                              | the fax client does not support this error code (the fax client API version      |
	//	|                                              | described in section 3.1.4.1.10 is FAX_API_VERSION_0).                           |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_DEVICE_NUM_LIMIT_EXCEEDED 0x00001B62 | The fax server cannot complete the operation because all of the following        |
	//	|                                              | conditions are true: If the fax server has an implementation-dependent maximum   |
	//	|                                              | number of supported devices set<167> and the number of fax devices currently     |
	//	|                                              | connected to the server is equal to or exceeds this maximum number of supported  |
	//	|                                              | devices. The device is not send, receive, or manual-receive enabled.             |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetPort(context.Context, *SetPortRequest) (*SetPortResponse, error)

	// The FAX_EnumRoutingMethods (Opnum 13) method is called by the client to enumerate
	// all the routing methods for a specified port that are registered with the fax server
	// in the fax server's list of routing methods. The client calls Fax_OpenPort (section
	// 3.1.4.1.65) to get the value for FaxPortHandle. The function returns detailed information
	// about each of the enumerated routing methods.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_FUNCTION 0x00000001  | The fax server failed to enumerate any routing methods for the fax port          |
	//	|                                    | specified through the FaxPortHandle parameter.                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the access rights  |
	//	|                                    | (FAX_ACCESS_QUERY_CONFIG) required for this operation.                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | Not enough storage is available to process this command. The fax server cannot   |
	//	|                                    | allocate sufficient memory to hold the array of FAX_ROUTING_METHOD to be         |
	//	|                                    | returned to the client.                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | This error SHOULD be returned if the FaxPortHandle argument is not a valid       |
	//	|                                    | handle obtained using FAX_OpenPort.<99>                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The RoutingInfoBuffer parameter is set to a NULL pointer value. <100>            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F    | The fax server failed to custom marshal the array of FAX_ROUTING_METHOD          |
	//	|                                    | structures to be returned to the client.                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EnumRoutingMethods(context.Context, *EnumRoutingMethodsRequest) (*EnumRoutingMethodsResponse, error)

	// The FAX_EnableRoutingMethod (Opnum 14) method is called by the client for a specified
	// fax device (port).
	//
	// The user is expected to set the proper configuration settings on the client before
	// enabling any routing method. For example, in order to enable email, the user SHOULD
	// specify the proper SMTP details, and the client SHOULD then call the FAX_SetReceiptsConfiguration
	// (section 3.1.4.1.91) method, with the bIsToUseForMSRouteThroughEmailMethod value
	// in the pReceipts parameter set to true. Also, the user can ensure that the proper
	// routing method destinations, such as EmailID, Printer, and Folder values, have been
	// specified. The client can use the FAX_SetExtensionData (section 3.1.4.1.79) method
	// to set the EmailID, Printer, and Folder once the user has entered the proper values.
	//
	// On success, the server MUST enable or disable a fax routing method for a specific
	// fax device. The server MUST validate that the client's fax user account has access
	// to enable or disable routing methods. The RoutingGUID parameter MUST be for a valid
	// routing method.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the                |
	//	|                                    | FAX_ACCESS_MANAGE_CONFIG access right.                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | The data is invalid. The GUID specified by the RoutingGuid parameter is not a    |
	//	|                                    | routing method GUID.                                                             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This is returned when RoutingGuid is set to NULL.    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EnableRoutingMethod(context.Context, *EnableRoutingMethodRequest) (*EnableRoutingMethodResponse, error)

	// The FAX_GetRoutingInfo (Opnum 15) method is called by the client to retrieve information
	// about a specified routing method that is identified by the passed-in GUID.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the                |
	//	|                                    | FAX_ACCESS_QUERY_CONFIG access rights required for this operation.               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The server cannot allocate sufficient memory to hold the fax routing information |
	//	|                                    | data to be returned to the client.                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | This error code is returned if any of the following conditions are met: § The    |
	//	|                                    | port handle specified by the FaxPortHandle parameter is not a valid fax port     |
	//	|                                    | handle obtained with FAX_OpenPort (section 3.1.4.1.65).<131> § The server cannot |
	//	|                                    | find the routing method identified by the GUID specified by the RoutingGuid      |
	//	|                                    | parameter.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned if any of the following  |
	//	|                                    | conditions are met: § The RoutingGuid parameter is set to a NULL pointer value.  |
	//	|                                    | § The RoutingInfoBuffer parameter is set to a NULL pointer value.<132> § The     |
	//	|                                    | FaxPortHandle parameter is set to a NULL value.<133>                             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetRoutingInfo(context.Context, *GetRoutingInfoRequest) (*GetRoutingInfoResponse, error)

	// The FAX_SetRoutingInfo (Opnum 16) method is called by the client to set routing information
	// for a fax routing method.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | The client's fax user account does not have the FAX_ACCESS_MANAGE_CONFIG         |
	//	|                                    | permission.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | This error code is returned if any of the following conditions are met: § The    |
	//	|                                    | FaxPortHandle parameter is not set to a valid fax port handle obtained with      |
	//	|                                    | FAX_OpenPort. § The RoutingGuid parameter is not set to a GUID representing a    |
	//	|                                    | valid routing method.                                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | This error code is returned if any of the following conditions are met: § The    |
	//	|                                    | RoutingGuid parameter is set to a NULL pointer value. § The RoutingInfoBuffer    |
	//	|                                    | parameter is set to a NULL pointer value. § The RoutingInfoBufferSize parameter  |
	//	|                                    | is set to a value of 0. § The FaxPortHandle parameter is set to a NULL           |
	//	|                                    | value.<172>                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	SetRoutingInfo(context.Context, *SetRoutingInfoRequest) (*SetRoutingInfoResponse, error)

	// The FAX_EnumGlobalRoutingInfo (Opnum 17) method is called by the client to enumerate
	// global routing information.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_FUNCTION 0x00000001  | The server failed to enumerate the routing methods.                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the                |
	//	|                                    | FAX_ACCESS_QUERY_CONFIG access rights required for this operation.               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The server cannot allocate sufficient memory to hold the array of                |
	//	|                                    | _FAX_GLOBAL_ROUTING_INFOW structures to be returned to the client.               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The RoutingInfoBuffer parameter is set to a NULL pointer value.<86>              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F    | The server failed to custom marshal the array of _FAX_GLOBAL_ROUTING_INFOW       |
	//	|                                    | structures to be returned to the client.                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EnumGlobalRoutingInfo(context.Context, *EnumGlobalRoutingInfoRequest) (*EnumGlobalRoutingInfoResponse, error)

	// The fax client application calls the FAX_SetGlobalRoutingInfo (Opnum 18) method to
	// set global routing properties such as the routing method priority.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the required       |
	//	|                                    | FAX_ACCESS_MANAGE_CONFIG authorization for this operation.                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | The fax server cannot find the routing method specified by the Guid structure    |
	//	|                                    | field of the RoutingInfo parameter.                                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned if any of the following  |
	//	|                                    | conditions are met: § The RoutingInfo parameter is set to a NULL pointer value.  |
	//	|                                    | § The SizeOfStruct structure field of RoutingInfo is not set to the correct      |
	//	|                                    | size, in bytes, for the FAX_GLOBAL_ROUTING_INFOW structure.                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetGlobalRoutingInfo(context.Context, *SetGlobalRoutingInfoRequest) (*SetGlobalRoutingInfoResponse, error)

	// The FAX_GetConfiguration (Opnum 19) method is called by the client to query the general
	// configuration of the fax server that is described by the _FAX_CONFIGURATIONW (section
	// 2.2.29) structure.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the                |
	//	|                                    | FAX_ACCESS_QUERY_CONFIG access rights required for this operation.               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The server cannot allocate sufficient memory to hold the FAX_CONFIGURATION       |
	//	|                                    | (section 2.2.28) data structure to be returned to the client.                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The buffer parameter is set to a NULL pointer value.<104>                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F    | The server failed to custom marshal the FAX_CONFIGURATION to be returned to the  |
	//	|                                    | client.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetConfiguration(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error)

	// The fax client application calls the FAX_SetConfiguration (Opnum 20) method to change
	// the general configuration of the fax server. The FAX_CONFIGURATIONW (section 2.2.28)
	// structure describes the general configuration of the fax server.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the required       |
	//	|                                    | access rights, in this case FAX_ACCESS_MANAGE_CONFIG.                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                    | following conditions: § The pointer specified with the FaxConfig argument is     |
	//	|                                    | NULL.<157> § The dwSizeOfStruct member of the FAX_CONFIGURATIONW specified by    |
	//	|                                    | the FaxConfig parameter is different from the correct size, in bytes, for the    |
	//	|                                    | FAX_CONFIGURATIONW, described in section 2.2.28. § The ArchiveOutgoingFaxes      |
	//	|                                    | member of the FAX_CONFIGURATIONW specified by FaxConfig is set to TRUE, and      |
	//	|                                    | the ArchiveDirectory member of the same data structure is set to a NULL pointer  |
	//	|                                    | value.<158>                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7  | The registry is corrupted. The structure of one of the files that contain        |
	//	|                                    | registry data is corrupted, or the system's memory image of the file is          |
	//	|                                    | corrupted, or the file could not be recovered because the alternate copy or log  |
	//	|                                    | was absent or corrupted.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetConfiguration(context.Context, *SetConfigurationRequest) (*SetConfigurationResponse, error)

	// The FAX_GetLoggingCategories (Opnum 21) method is called by the client. In response,
	// the server MUST return the current logging categories for the fax server to which
	// the client has connected. A logging category determines the errors or other events
	// that the fax server records in the application event log.
	//
	// The client SHOULD free the returned buffer.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005       | Access is denied. The client's fax user account does not have the                |
	//	|                                      | FAX_ACCESS_QUERY_CONFIG access rights required for this operation.               |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008   | The server cannot allocate sufficient memory to hold the array of                |
	//	|                                      | FAX_LOG_CATEGORY to be returned to the client.                                   |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057   | The parameter is incorrect. This error code is returned if any of the            |
	//	|                                      | following conditions are met: § The BufferSize parameter is set to a NULL        |
	//	|                                      | pointer value.<114> § The NumberCategories parameter is set to a NULL pointer    |
	//	|                                      | value.<115>                                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ARITHMETIC_OVERFLOW 0x00000216 | This error code is returned if any of the following conditions are met: § The    |
	//	|                                      | total number of logging categories multiplied by the size, in bytes, of the      |
	//	|                                      | FAX_LOG_CATEGORY Fixed_Size block results in a number that exceeds the maximum   |
	//	|                                      | value for a DWORD (0xFFFFFFFF). § The total number of logging categories         |
	//	|                                      | multiplied by the size, in bytes, of the FAX_LOG_CATEGORY Fixed_Size block       |
	//	|                                      | plus the sum of all lengths, in bytes, including NULL terminators, of all Name   |
	//	|                                      | strings from the Variable_Data blocks results in a number that exceeds the       |
	//	|                                      | maximum value for a DWORD (0xFFFFFFFF).                                          |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F      | The server failed to custom marshal the array of FAX_LOG_CATEGORY to be returned |
	//	|                                      | to the client.                                                                   |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetLoggingCategories(context.Context, *GetLoggingCategoriesRequest) (*GetLoggingCategoriesResponse, error)

	// The FAX_SetLoggingCategories (Opnum 22) method is called by the client. On success,
	// the server MUST modify the current logging categories for the fax server to which
	// the client has connected. A logging category determines the errors or other events
	// that the fax server records in the application event log.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the                |
	//	|                                    | FAX_ACCESS_MANAGE_CONFIG access rights required for this call.                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                    | following conditions: § The value specified for the Buffer parameter is NULL. §  |
	//	|                                    | The value specified for the BufferSize parameter is 0. § The fax server cannot   |
	//	|                                    | parse the FAX_LOG_CATEGORY pointed at by Buffer, possibly because the buffer     |
	//	|                                    | data is corrupted. § One of the pointer fields of one of the FAX_LOG_CATEGORY    |
	//	|                                    | pointed at by Buffer point to memory locations outside of the memory block       |
	//	|                                    | specified by the Buffer and BufferSize parameters.                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7  | The registry is corrupted. The fax server cannot write to register the modified  |
	//	|                                    | logging categories.                                                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The FAX_LOG_CATEGORY array MUST be serialized. For more information, see [MSDN-FAX_LOG_CATEGORY].
	// The variable data fields, such as strings, SHOULD be filled with the offset to the
	// string from the beginning of the buffer and not the actual address.
	SetLoggingCategories(context.Context, *SetLoggingCategoriesRequest) (*SetLoggingCategoriesResponse, error)

	// The FAX_GetSecurity (Opnum 23) method is called by the client to retrieve information
	// about the fax security descriptor from the fax server.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                                  |
	//	|               VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005          | Access is denied. The client's fax user account does not have a permission level |
	//	|                                         | of at least READ_CONTROL.                                                        |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008      | Not enough storage is available to process this command.                         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057      | The parameter is incorrect. This is returned when pSecurityDescriptor is NULL.   |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_SECURITY_DESCR 0x0000053A | The security descriptor structure is invalid.                                    |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetSecurity(context.Context, *GetSecurityRequest) (*GetSecurityResponse, error)

	// The FAX_SetSecurity (Opnum 24) method is called by the client. On success, the server
	// MUST set the fax server's security descriptor.
	//
	// Protocol version FAX_API_VERSION_3 (0x00030000) fax servers SHOULD fail this call
	// by returning ERROR_NOT_SUPPORTED (0x00000032). The fax client SHOULD NOT call this
	// method if the protocol version reported by the server is FAX_API_VERSION_3 (0x00030000).
	// For more information, see FAX_ConnectFaxServer (section 3.1.4.1.10). The fax client
	// SHOULD call FAX_SetSecurityEx2 (section 3.1.4.1.95) instead.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the access         |
	//	|                                    | rights required for this operation. This error code is returned under any        |
	//	|                                    | of the following conditions, listed by required access right: § WRITE_OWNER,     |
	//	|                                    | when the fax server is a FAX_API_VERSION_1 server and the SecurityInformation    |
	//	|                                    | parameter contains the OWNER_SECURITY_INFORMATION value. § WRITE_DAC, when the   |
	//	|                                    | fax server is a FAX_API_VERSION_1 server and SecurityInformation contains the    |
	//	|                                    | GROUP_SECURITY_INFORMATION or DACL_SECURITY_INFORMATION values. § READ_CONTROL,  |
	//	|                                    | when the fax server is a FAX_API_VERSION_2 server and SecurityInformation        |
	//	|                                    | contains the GROUP_SECURITY_INFORMATION, DACL_SECURITY_INFORMATION,              |
	//	|                                    | or OWNER_SECURITY_INFORMATION values. § ACCESS_SYSTEM_SECURITY, when             |
	//	|                                    | SecurityInformation contains the SACL_SECURITY_INFORMATION value.                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | The data contained in the buffer specified by the pSecurityDescriptor parameter  |
	//	|                                    | is not a valid SECURITY_DESCRIPTOR.                                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_SUPPORTED 0x00000032     | The fax server does not support this operation. This error SHOULD be returned by |
	//	|                                    | FAX_API_VERSION_3 servers.                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                    | following conditions: § pSecurityInformation is set to a NULL pointer value. §   |
	//	|                                    | The dwBufferSize parameter is set to a value of 0.                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7  | The registry is corrupted. The structure of one of the files containing registry |
	//	|                                    | data is corrupted, or the system's memory image of the file is corrupted, or     |
	//	|                                    | the file could not be recovered because the alternate copy or log was absent or  |
	//	|                                    | corrupted.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	SetSecurity(context.Context, *SetSecurityRequest) (*SetSecurityResponse, error)

	// The FAX_AccessCheck (Opnum 25) method is called when the client needs to check whether
	// the client's fax user account has certain access permissions on the server.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The user does not have a valid fax user account on the         |
	//	|                                    | server.<65>                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. The pfAccess argument is NULL,<66> or the access     |
	//	|                                    | mask specified by the AccessMask argument contains invalid fax-specific access   |
	//	|                                    | rights.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	AccessCheck(context.Context, *AccessCheckRequest) (*AccessCheckResponse, error)

	// The FAX_CheckServerProtSeq (Opnum 26) method is called by the client to validate
	// whether a specified protocol sequence is supported by the server.<70> In response,
	// the server MUST validate the specified protocol sequence.
	//
	// Protocol version FAX_API_VERSION_2 (0x00020000) and FAX_API_VERSION_3 (0x00030000)
	// fax servers SHOULD fail this call by returning ERROR_NOT_SUPPORTED (0x00000032).
	// The fax client SHOULD NOT call this method if the protocol version reported by the
	// server is FAX_API_VERSION_2 (0x00020000) or FAX_API_VERSION_3 (0x00030000). For more
	// information, see FAX_ConnectFaxServer (section 3.1.4.1.10).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+-----------------------------------------+-------------------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                               |
	//	|               VALUE/CODE                |                                  DESCRIPTION                                  |
	//	|                                         |                                                                               |
	//	+-----------------------------------------+-------------------------------------------------------------------------------+
	//	+-----------------------------------------+-------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057      | The parameter is incorrect. The lpdwProtSeq argument specified is NULL.       |
	//	+-----------------------------------------+-------------------------------------------------------------------------------+
	//	| RPC_S_PROTSEQ_NOT_SUPPORTED 0x0000006A7 | The protocol sequence specified by the lpdwProtSeq argument is not supported. |
	//	+-----------------------------------------+-------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	CheckServerProtocolSeq(context.Context, *CheckServerProtocolSeqRequest) (*CheckServerProtocolSeqResponse, error)

	// The FAX_SendDocumentEx (Opnum 27) method is called by the client.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors (section
	// 2.2.52) or one of the standard errors ([MS-ERREF] section 2.2).
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                                  |
	//	|                  VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005               | Access is denied. This error is returned when any of the following conditions    |
	//	|                                              | are met: § The limit on the number of recipients for a single fax broadcast      |
	//	|                                              | was reached and FAX_ERR_RECIPIENTS_LIMIT couldn't be returned because this       |
	//	|                                              | error code is unsupported by the fax client API version (FAX_API_VERSION_0       |
	//	|                                              | or FAX_API_VERSION_1, described in section 3.1.4.1.10). § The client's fax       |
	//	|                                              | user account does not have the required access rights to submit the fax:         |
	//	|                                              | FAX_ACCESS_SUBMIT (for FAX_PRIORITY_TYPE_LOW), FAX_ACCESS_SUBMIT_NORMAL          |
	//	|                                              | (for FAX_PRIORITY_TYPE_NORMAL), or FAX_ACCESS_SUBMIT_HIGH (for                   |
	//	|                                              | FAX_PRIORITY_TYPE_HIGH), where the FAX_PRIORITY_TYPE value comes from the        |
	//	|                                              | Priority field of the specified lpJobParams structure.                           |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008           | Not enough storage is available to process this command.                         |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D                | The file specified by lpcwstrFileName argument is empty (has a size of 0 bytes). |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_WRITE_PROTECT 0x00000013               | The outgoing fax queue is blocked: The queue state is FAX_OUTBOX_BLOCKED as      |
	//	|                                              | described in the FAX_SetQueue (section 3.1.4.1.90) method.                       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_SUPPORTED 0x00000032               | The fax server SHOULD return this error code when the request described by the   |
	//	|                                              | lpJobParams argument is not supported by the fax server.<152>                    |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057           | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                              | following conditions: § One or more of the following arguments are NULL          |
	//	|                                              | or 0: dwNumRecipients, lpcSenderProfile, lpcRecipientList, lpJobParams,          |
	//	|                                              | lpdwlMessageId, lpdwlRecipientMessageIds,<153>lpcCoverPageInfo. § The file       |
	//	|                                              | name indicated by the lpcstwrFileName argument does not indicate a file          |
	//	|                                              | of the expected TIFF format. § The lpwstrCoverPageFileName field of the          |
	//	|                                              | lpcCoverPageInfo structure is not in the expected COV format. § The Priority     |
	//	|                                              | field of the lpJobParams structure is not one of the following values:           |
	//	|                                              | FAX_PRIORITY_TYPE_LOW, FAX_PRIORITY_TYPE_NORMAL, FAX_PRIORITY_TYPE_HIGH. § The   |
	//	|                                              | requested receipt delivery types are invalid (the dwReceiptDeliveryType field    |
	//	|                                              | of the lpJobParams structure), not DRT_EMAIL, DRT_MSGBOX, and/or DRT_NONE. §     |
	//	|                                              | The fax server tried to return FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU but the client  |
	//	|                                              | fax API version. FAX_API_VERSION_0 (described in section 3.1.4.1.10) does not    |
	//	|                                              | support this error code.                                                         |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_UNSUPPORTED_TYPE 0x0000065E            | Data of this type is not supported. The requested receipt delivery type          |
	//	|                                              | specified by the dwReceiptDeliveryType field of the lpJobParams structure is not |
	//	|                                              | supported by the fax server.                                                     |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU 0x00001B63 | The fax client module API version (described in section 3.1.4.1.10) is           |
	//	|                                              | FAX_API_VERSION_1 or above, and the fax server is running on a version of the    |
	//	|                                              | operating system that does not support the requested operation.<154>             |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_RECIPIENTS_LIMIT 0x00001B65          | The limit on the number of recipients for a single fax broadcast was reached.    |
	//	|                                              | ERROR_ACCESS_DENIED is returned instead of this error code when the client       |
	//	|                                              | does not support it (client-supported fax API version is FAX_API_VERSION_0 or    |
	//	|                                              | FAX_API_VERSION_1, described in section 3.1.4.1.10).                             |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	SendDocumentEx(context.Context, *SendDocumentExRequest) (*SendDocumentExResponse, error)

	// The FAX_EnumJobsEx (Opnum 28) method is called by the client to enumerate a specified
	// set of jobs on the server's queue. The type of jobs to enumerate is described by
	// the dwJobTypes argument.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have any of the         |
	//	|                                    | access rights defined in ALL_FAX_USER_ACCESS_RIGHTS (section 2.2.83) when        |
	//	|                                    | enumerating jobs of type JT_SEND. The client's fax user account does not have    |
	//	|                                    | the FAX_ACCESS_MANAGE_RECEIVE_FOLDER access right when enumerating jobs of type  |
	//	|                                    | JT_RECEIVE or JT_ROUTING.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | Not enough storage is available to process this command.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | At least one of the following arguments has been specified as NULL: Buffer,      |
	//	|                                    | BufferSize, or lpdwJobs.<88>                                                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EnumJobsEx(context.Context, *EnumJobsExRequest) (*EnumJobsExResponse, error)

	// The FAX_GetJobEx (Opnum 29) is called by the client to retrieve information about
	// a specified job at the server. The job is identified by the job message ID. The job
	// message ID can be obtained using one of the following methods: FAX_EnumJobs (section
	// 3.1.4.1.21), FAX_EnumJobsEx (section 3.1.4.1.22), or FAX_EnumJobsEx2 (section 3.1.4.1.23).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005       | Access is denied. The client's fax user account does not have the access rights  |
	//	|                                      | required for this operation: ALL_FAX_USER_ACCESS_RIGHTS (section 2.2.83).        |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008   | The server cannot allocate memory for the data to be returned to the caller.     |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057   | This error code is returned under any of the following conditions:<111> § The    |
	//	|                                      | pointer submitted for the Buffer argument is NULL. § The pointer submitted for   |
	//	|                                      | the BufferSize argument is NULL.                                                 |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_MESSAGE_NOT_FOUND 0x00001B61 | The fax server cannot find the fax job indicated by the dwlMessageId argument.   |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetJobEx(context.Context, *GetJobExRequest) (*GetJobExResponse, error)

	// The FAX_GetCountryList (Opnum 30) method is called by the client to retrieve the
	// list of country/region information that is defined on the server. TAPI maintains
	// this list, which contains information like the country/region name or country/region
	// ID.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have any of the         |
	//	|                                    | permissions covered by the ALL_FAX_USER_ACCESS_RIGHTS enumeration (section       |
	//	|                                    | 2.2.83).                                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | Not enough storage is available to process this command. The fax server cannot   |
	//	|                                    | allocate sufficient memory to hold the FAX_TAPI_LINECOUNTRY_LISTW to be returned |
	//	|                                    | to the client.                                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error is returned if the Buffer parameter is    |
	//	|                                    | set to a NULL pointer value.<105>                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F    | The fax server failed to custom marshal the FAX_TAPI_LINECOUNTRY_LISTW to be     |
	//	|                                    | returned to the client.                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetCountryList(context.Context, *GetCountryListRequest) (*GetCountryListResponse, error)

	// The FAX_GetPersonalProfileInfo (Opnum 31) method is called by the client to retrieve
	// information about the personal profile (section 3.1.1) of the sender or the receiver
	// of a fax from the specified fax message that is present in the described message
	// folder. The Folder value MUST be one of the enumerations defined by FAX_ENUM_MESSAGE_FOLDER
	// (section 2.2.2) except FAX_MESSAGE_FOLDER_INBOX. The ProfType value MUST be one of
	// the enumerations that are defined by FAX_ENUM_PERSONAL_PROF_TYPES (section 2.2.4).
	// The dwlMessageId parameter specifies a particular message and can be obtained using
	// the following methods: FAX_EnumJobs (section 3.1.4.1.21), FAX_EnumJobsEx (section
	// 3.1.4.1.22), FAX_EnumMessages (section 3.1.4.1.24), or FAX_EnumMessagesEx (section
	// 3.1.4.1.25).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005       | Access is denied. The client's fax user account does not have any of the         |
	//	|                                      | permissions covered by ALL_FAX_USER_ACCESS_RIGHTS.                               |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BAD_FORMAT 0x0000000B          | The message identified by dwlMessageId is an archived sent message and there was |
	//	|                                      | an error in reading the message file.                                            |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057   | The parameter is incorrect. This error code is returned when any of              |
	//	|                                      | the following conditions happen: § The pointer specified by the Buffer           |
	//	|                                      | parameter is NULL.<122> § The value of the Folder parameter specified is not     |
	//	|                                      | FAX_MESSAGE_FOLDER_QUEUE or FAX_MESSAGE_FOLDER_SENTITEMS.                        |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_SRV_OUTOFMEMORY 0x00001B59   | The fax server failed to allocate memory needed for internal execution of this   |
	//	|                                      | operation.                                                                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_MESSAGE_NOT_FOUND 0x00001B61 | This error code is returned when any of the following conditions happen: § The   |
	//	|                                      | message identified by the dwlMessageId parameter is not found. § The dwFolder    |
	//	|                                      | parameter is specified as FAX_MESSAGE_FOLDER_QUEUE, but the message identified   |
	//	|                                      | by dwlMessageId is not an outgoing queued message. § The message identified      |
	//	|                                      | by dwlMessageId is an outgoing queued message being deleted. § The message       |
	//	|                                      | identified by dwlMessageId is an outgoing queued message of a different user,    |
	//	|                                      | and this user does not have FAX_ACCESS_QUERY_OUT_JOBS permission. § The message  |
	//	|                                      | identified by dwlMessageId is an archived sent message of a different user, and  |
	//	|                                      | this user does not have FAX_ACCESS_QUERY_ARCHIVES permission.                    |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetPersonalProfileInfo(context.Context, *GetPersonalProfileInfoRequest) (*GetPersonalProfileInfoResponse, error)

	// The FAX_GetQueueStates (Opnum 32) method is called by the client to retrieve the
	// state of the fax queues at the server.
	//
	// The pdwQueueStates parameter MUST NOT be NULL. On success, the server MUST return
	// the state information about the fax service.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have any of the         |
	//	|                                    | permissions covered by ALL_FAX_USER_ACCESS_RIGHTS (section 2.2.83).              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. The pdwQueueStates parameter is set to a NULL        |
	//	|                                    | pointer value. <128>                                                             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetQueueStates(context.Context, *GetQueueStatesRequest) (*GetQueueStatesResponse, error)

	// The fax client application calls the FAX_SetQueue (Opnum 33) method to change the
	// state of the server queue. In response, the server MUST validate whether the client's
	// fax user account has authorization to set the Queue State of the server. On success,
	// the server MUST set its Queue State as specified by the client.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the required       |
	//	|                                    | FAX_ACCESS_MANAGE_CONFIG authorization for this operation.                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_OUTOFMEMORY 0x0000000E       | The fax server cannot allocate sufficient memory for a FAX_EVENT_EX_1 (section   |
	//	|                                    | 2.2.68) structure that describes a FAX_EVENT_QUEUE_TYPE_QUEUE_STATE event to be  |
	//	|                                    | signaled to the client. For more details, see FAX_ClientEventQueueEx (section    |
	//	|                                    | 3.2.4.3).                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. The dwQueueStates parameter is set to a              |
	//	|                                    | combination of values that does not contain any of the supported values:         |
	//	|                                    | FAX_INCOMING_BLOCKED, FAX_OUTBOX_BLOCKED, or FAX_OUTBOX_PAUSED.                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7  | The registry is corrupted. The structure of one of the files containing registry |
	//	|                                    | data is corrupted, or the system's memory image of the file is corrupted, or     |
	//	|                                    | the file could not be recovered because the alternate copy or log was absent or  |
	//	|                                    | corrupted.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| RPC_E_SYS_CALL_FAILED 0x80010100   | dwQueueStates includes the FAX_OUTBOX_PAUSED value and the fax server cannot     |
	//	|                                    | pause the server queue, or dwQueueStates does not include the FAX_OUTBOX_PAUSED  |
	//	|                                    | value and the fax server cannot resume the server queue.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetQueue(context.Context, *SetQueueRequest) (*SetQueueResponse, error)

	// The FAX_GetReceiptsConfiguration (Opnum 34) method is called by the client. On success,
	// the server MUST return the receipt configuration information of the fax server.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors ([MS-ERREF] section
	// 2.2).
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the                |
	//	|                                    | FAX_ACCESS_QUERY_CONFIG access rights required for this operation.               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | Not enough storage is available to process this command. The fax server cannot   |
	//	|                                    | allocate sufficient memory to hold the _FAX_RECEIPTS_CONFIGW to be returned to   |
	//	|                                    | the client.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. The Buffer parameter is set to a NULL pointer        |
	//	|                                    | value.<129>                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F    | The fax server failed to custom marshal the _FAX_RECEIPTS_CONFIGW to be returned |
	//	|                                    | to the client.                                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetReceiptsConfiguration(context.Context, *GetReceiptsConfigurationRequest) (*GetReceiptsConfigurationResponse, error)

	// The FAX_SetReceiptsConfiguration (Opnum 35) method is called by the client. On success,
	// the server MUST set the receipt configuration information that is used by the fax
	// server to send delivery receipts for fax transmissions.<169>
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                                  |
	//	|                  VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005               | Access is denied. The client's fax user account does not have the required       |
	//	|                                              | FAX_ACCESS_MANAGE_CONFIG authorization for this operation.                       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_SUPPORTED 0x00000032               | The dwAllowedReceipts member of the pReceipts parameter contains the DRT_MSGBOX  |
	//	|                                              | flag value (described in FAX_ENUM_DELIVERY_REPORT_TYPES), and the fax server     |
	//	|                                              | does not support this type of fax receipts.<170>                                 |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057           | The parameter is incorrect. This error code is returned if any of the            |
	//	|                                              | following conditions are met: § The dwSizeOfStruct member of pReceipts           |
	//	|                                              | is not equal to the correct size, in bytes, for the FAX_RECEIPTS_CONFIGW.        |
	//	|                                              | § The dwAllowedReceipts member of pReceipts contains one or more invalid         |
	//	|                                              | flag values (not defined for the FAX_ENUM_DELIVERY_REPORT_TYPES (section         |
	//	|                                              | 2.2.76) enumeration). § The following member values are set in the submitted     |
	//	|                                              | pReceipts: the DRT_EMAIL flag is set within the dwAllowedReceipts member,        |
	//	|                                              | the bIsToUseForMSRouteThroughEmailMethod member is set to FALSE, and the         |
	//	|                                              | lpwstrSMTPPasswordmember is set to a non-NULL pointer value. § In the submitted  |
	//	|                                              | pReceipts, the DRT_EMAIL flag is set within the dwAllowedReceipts member         |
	//	|                                              | or the bIsToUseForMSRouteThroughEmailMethod member is set to TRUE and the        |
	//	|                                              | SMTPAuthOption member is set to a value lower than FAX_SMTP_AUTH_ANONYMOUS       |
	//	|                                              | or greater than FAX_SMTP_AUTH_NTLM. § The fax server tried to return             |
	//	|                                              | FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU but the client fax API version                 |
	//	|                                              | (FAX_API_VERSION_0, described in section 3.1.4.1.10) does not support this error |
	//	|                                              | code.                                                                            |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7            | The registry is corrupted. The structure of one of the files containing registry |
	//	|                                              | data is corrupted, or the system's memory image of the file is corrupted, or     |
	//	|                                              | the file could not be recovered because the alternate copy or log was absent or  |
	//	|                                              | corrupted.                                                                       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU 0x00001B63 | The fax client module API version (described in section 3.1.4.1.10) is           |
	//	|                                              | FAX_API_VERSION_1 or above, and the fax server is running on a version of the    |
	//	|                                              | operating system that does not support the requested operation.<171>             |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetReceiptsConfiguration(context.Context, *SetReceiptsConfigurationRequest) (*SetReceiptsConfigurationResponse, error)

	// The FAX_GetReceiptsOptions (Opnum 36) method is called by the client to retrieve
	// the supported receipt options on the server.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return the following error code, one of the fax-specific errors that are
	// defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | Access is denied. The client's fax user account does not have any of the         |
	//	|                                | permissions covered by ALL_FAX_USER_ACCESS_RIGHTS (section 2.2.83).              |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetReceiptsOptions(context.Context, *GetReceiptsOptionsRequest) (*GetReceiptsOptionsResponse, error)

	// The fax client application calls the FAX_GetVersion (Opnum 37) method to obtain the
	// version of the fax server it is connected to.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return the following error codes, one of the fax-specific errors that are
	// defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | Access is denied. The client's fax user account does not have the                |
	//	|                                | ALL_FAX_USER_ACCESS_RIGHTS (section 2.2.83) access rights required for this      |
	//	|                                | operation.                                                                       |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)

	// The FAX_GetOutboxConfiguration (Opnum 38) method is called by the client to retrieve
	// the outbox configuration at the server.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the                |
	//	|                                    | FAX_ACCESS_QUERY_CONFIG access rights required for this operation.               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | Not enough storage is available to process this command. The fax server cannot   |
	//	|                                    | allocate sufficient memory to hold the FAX_OUTBOX_CONFIG to be returned to the   |
	//	|                                    | client.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. The Buffer parameter is set to a NULL pointer        |
	//	|                                    | value.<119>                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetOutboxConfiguration(context.Context, *GetOutboxConfigurationRequest) (*GetOutboxConfigurationResponse, error)

	// The fax client application calls the FAX_SetOutboxConfiguration (Opnum 39) method
	// to set the current Outbox configuration such as the Discount Time.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the required       |
	//	|                                    | FAX_ACCESS_MANAGE_CONFIG authorization for this operation.                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned if any of the following  |
	//	|                                    | conditions occur: § The dwSizeOfStruct member of the structure pointed at by     |
	//	|                                    | the pOutboxCfg parameter is not the correct size for the FAX_OUTBOX_CONFIG.      |
	//	|                                    | § The dtDiscountStart or dtDiscountEnd members of the structure pointed to by    |
	//	|                                    | pOutboxCfg contain one or more invalid Hour (value greater than 24) or Minute    |
	//	|                                    | (value greater than 60) fields.                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7  | The registry is corrupted. The structure of one of the files containing registry |
	//	|                                    | data is corrupted, or the system's memory image of the file is corrupted, or     |
	//	|                                    | the file could not be recovered because the alternate copy or log was absent or  |
	//	|                                    | corrupted.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetOutboxConfiguration(context.Context, *SetOutboxConfigurationRequest) (*SetOutboxConfigurationResponse, error)

	// The FAX_GetPersonalCoverPagesOption (Opnum 40) method is called by the client to
	// retrieve information about the supported personal cover-page options.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return the following error code, one of the fax-specific errors that are
	// defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | Access is denied. The client's fax user account does not have any of the         |
	//	|                                | permissions covered by ALL_FAX_USER_ACCESS_RIGHTS (section 2.2.83).<121>         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetPersonalCoverPagesOption(context.Context, *GetPersonalCoverPagesOptionRequest) (*GetPersonalCoverPagesOptionResponse, error)

	// The FAX_GetArchiveConfiguration (Opnum 41) method is called by the client to retrieve
	// the current archive configuration on the fax server. In response, the server returns
	// archive configuration information about the fax server.
	//
	// Protocol version FAX_API_VERSION_3 (0x00030000) fax servers SHOULD fail this call
	// by returning ERROR_NOT_SUPPORTED (0x00000032). The fax client SHOULD NOT call this
	// method if the protocol version reported by the server is FAX_API_VERSION_3 (0x00030000).
	// For more information, see FAX_ConnectFaxServer (section 3.1.4.1.10). The fax client
	// SHOULD call FAX_GetGeneralConfiguration (section 3.1.4.1.40) instead.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the access rights  |
	//	|                                    | (FAX_ACCESS_QUERY_CONFIG) required for this operation.                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned if any of the            |
	//	|                                    | following conditions are met: § The Buffer parameter is set to a NULL            |
	//	|                                    | pointer value. <102> § The value specified for the Folder parameter is not       |
	//	|                                    | FAX_MESSAGE_FOLDER_SENTITEMS or FAX_MESSAGE_FOLDER_INBOX.                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | Not enough storage is available to process this command. The fax server cannot   |
	//	|                                    | allocate sufficient memory to hold the FAX_ARCHIVE_CONFIGW to be returned to the |
	//	|                                    | client.                                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_SUPPORTED 0x00000032     | The fax server does not implement this method. Protocol version                  |
	//	|                                    | FAX_API_VERSION_3 (0x00030000) fax servers SHOULD fail this call by returning    |
	//	|                                    | this error code.                                                                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetArchiveConfiguration(context.Context, *GetArchiveConfigurationRequest) (*GetArchiveConfigurationResponse, error)

	// The fax client application calls the FAX_SetArchiveConfiguration (Opnum 42) method
	// to set the archive configuration for a specific fax folder on the fax server.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors (section
	// 2.2.52), or one of the other standard errors ([MS-ERREF] section 2.2).
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005        | Access is denied. The client's fax user account does not have the required       |
	//	|                                       | FAX_ACCESS_MANAGE_CONFIG authorization for this operation.                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_SUPPORTED 0x00000032        | The fax server does not support this operation. This error code SHOULD be        |
	//	|                                       | returned by the FAX_API_VERSION_3 servers.                                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057    | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                       | following conditions: § The dwSizeOfStruct member of the FAX_ARCHIVE_CONFIGW     |
	//	|                                       | specified by the pArchiveCfg parameter is set to an incorrect value. § The       |
	//	|                                       | Folder parameter is set to a value other than FAX_MESSAGE_FOLDER_SENTITEMS and   |
	//	|                                       | FAX_MESSAGE_FOLDER_INBOX. § The bUseArchive member of the FAX_ARCHIVE_CONFIGW    |
	//	|                                       | specified by pArchiveCfg is set to TRUE, and any of the three following          |
	//	|                                       | conditions are also met: § The value of the dwSizeQuotaHighWatermark member      |
	//	|                                       | of the same structure is smaller than the value of the dwSizeQuotaLowWatermark   |
	//	|                                       | member. § The lpcstrFolder member of the same structure is set to NULL. § The    |
	//	|                                       | lpcstrFolder member of the same structure is set to point to an empty character  |
	//	|                                       | string.                                                                          |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BUFFER_OVERFLOW 0x0000006F      | The file name is too long. The bUseArchive member of the FAX_ARCHIVE_CONFIGW     |
	//	|                                       | specified by pArchiveCfg is set to a value of TRUE, and the length of the file   |
	//	|                                       | name specified by the lpcstrFolder of the same structure is set to a character   |
	//	|                                       | string longer than 180 characters, excluding the length of the null terminator.  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7     | The fax server cannot store the new archive configuration to the registry. The   |
	//	|                                       | registry could be corrupted.                                                     |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_DIRECTORY_IN_USE 0x00001B5F   | The bUseArchive member of the FAX_ARCHIVE_CONFIGW specified by pArchiveCfg is    |
	//	|                                       | set to a value of TRUE, and the file specified by the lpcstrFolder member of     |
	//	|                                       | the same structure is invalid: the directory of the file is the same as the fax  |
	//	|                                       | queue directory.                                                                 |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_FILE_ACCESS_DENIED 0x00001B60 | The bUseArchive member of the FAX_ARCHIVE_CONFIGW specified by pArchiveCfg       |
	//	|                                       | is set to a value of TRUE, and the file specified by the lpcstrFolder member     |
	//	|                                       | of the same structure is invalid: the fax server encountered an access denied    |
	//	|                                       | (ERROR_ACCESS_DENIED) or a sharing violation (ERROR_SHARING_VIOLATION) error     |
	//	|                                       | when attempting to access the file.                                              |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetArchiveConfiguration(context.Context, *SetArchiveConfigurationRequest) (*SetArchiveConfigurationResponse, error)

	// The FAX_GetActivityLoggingConfiguration (Opnum 43) method is called by the client
	// to retrieve the current activity logging configuration.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the access rights  |
	//	|                                    | (FAX_ACCESS_QUERY_CONFIG) required for this operation.                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | Not enough storage is available to process this command. The fax server cannot   |
	//	|                                    | allocate sufficient memory to hold the _FAX_ACTIVITY_LOGGING_CONFIGW to be       |
	//	|                                    | returned to the client.                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The Buffer parameter is set to a NULL pointer value. <101>                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F    | The fax server failed to custom marshal the _FAX_ACTIVITY_LOGGING_CONFIGW to be  |
	//	|                                    | returned to the client.                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetActivityLoggingConfiguration(context.Context, *GetActivityLoggingConfigurationRequest) (*GetActivityLoggingConfigurationResponse, error)

	// The fax client application calls the FAX_SetActivityLoggingConfiguration (Opnum 44)
	// method to set options for activity logging. This includes setting whether entries
	// for incoming and outgoing faxes SHOULD be logged and the location of the log file.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005        | Access denied. This error code is returned under any of the following            |
	//	|                                       | conditions: § The client's fax user account does not have the required           |
	//	|                                       | FAX_ACCESS_MANAGE_CONFIG authorization to manage the server configuration. §     |
	//	|                                       | The directory specified by the lpwstrDBPath member of the pActivLogCfg structure |
	//	|                                       | is not a valid fax folder (the fax server does not have rights to create files,  |
	//	|                                       | write to files, enumerate files and/or delete files), the server needs to return |
	//	|                                       | FAX_ERR_FILE_ACCESS_DENIED, and the client does not support this error code.     |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057    | This error code is returned under any of the following conditions: § The         |
	//	|                                       | dwSizeOfStruct member of the FAX_ACTIVITY_LOGGING_CONFIGW specified by the       |
	//	|                                       | pActivLogCfg parameter is set to an incorrect value. The correct size, in bytes, |
	//	|                                       | described in the FAX_ACTIVITY_LOGGING_CONFIGW structure. § The lpwstrDBPath      |
	//	|                                       | field of the structure specified by the pActivLogCfg argument contains an empty  |
	//	|                                       | string or is set to NULL. § The lpwstrDBPath member of the structure specified   |
	//	|                                       | by the pActivLogCfg argument does not indicate a complete path name.             |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BUFFER_OVERFLOW 0x0000006F      | The file name is too long. The lpwstrDBPath field of the structure specified by  |
	//	|                                       | the pActivityLog argument contains a path name longer than 248 characters, not   |
	//	|                                       | counting the terminating null character.                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7     | The fax server cannot store the new activity logging configuration in the        |
	//	|                                       | registry. This error can happen if the registry is corrupted.                    |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_FILE_ACCESS_DENIED 0x00001B60 | This error code is returned under any of the following conditions: § The         |
	//	|                                       | folder specified by the lpwstrDBPath member of the pActivityLog structure        |
	//	|                                       | is not a valid fax folder where the server has rights to create, enumerate,      |
	//	|                                       | write to, and delete files. § The fax server cannot create a new (if different   |
	//	|                                       | from the respective existing file) file specified by the lpwstrDBPath            |
	//	|                                       | member of pActivLogCfg, because the server encountered an access denied          |
	//	|                                       | (ERROR_ACCESS_DENIED) or sharing violation error (ERROR_SHARING_VIOLATION) when  |
	//	|                                       | attempting to create the specified file.                                         |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetActivityLoggingConfiguration(context.Context, *SetActivityLoggingConfigurationRequest) (*SetActivityLoggingConfigurationResponse, error)

	// The FAX_EnumerateProviders (Opnum 45) method is called by the client to enumerate
	// all the FSPs installed on the server.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the                |
	//	|                                    | FAX_ACCESS_QUERY_CONFIG access rights required for this operation.               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error is returned if the BufferSize and/or the  |
	//	|                                    | lpdwNumProviders parameters are set to NULL pointer values.                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F    | The fax server failed to custom marshal the array of FAX_DEVICE_PROVIDER_INFO    |
	//	|                                    | structures to be returned to the client.                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_SRV_OUTOFMEMORY 0x00001B59 | The fax server failed to allocate sufficient memory for the return buffer to     |
	//	|                                    | hold the FAX_DEVICE_PROVIDER_INFO to be returned to the client.                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EnumerateProviders(context.Context, *EnumerateProvidersRequest) (*EnumerateProvidersResponse, error)

	// The FAX_GetPortEx (Opnum 46) method is called by the client to retrieve port status
	// information for a requested port at the server. The device ID that is passed in SHOULD
	// be obtained from FAX_EnumPorts (section 3.1.4.1.28). This method is an extended version
	// of FAX_GetPort (section 3.1.4.1.51).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the                |
	//	|                                    | FAX_ACCESS_QUERY_CONFIG access rights required for this call.                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | Not enough storage is available to process this command. The fax server cannot   |
	//	|                                    | allocate sufficient memory to hold the FAX_PORT_INFO_EXW (section 2.2.45)        |
	//	|                                    | structure to be returned to the client.                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BAD_UNIT 0x00000014          | The system cannot find the port for the receiving device by using the line       |
	//	|                                    | identifier specified by the dwDeviceId argument.                                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. The Buffer parameter is set to a NULL pointer        |
	//	|                                    | value.<127>                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F    | The fax server failed to custom marshal the FAX_PORT_INFO_EXW to be returned to  |
	//	|                                    | the client.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetPortEx(context.Context, *GetPortExRequest) (*GetPortExResponse, error)

	// A fax client application uses the FAX_SetPortEx (Opnum 47) method to set fax device
	// information. The function sets extension configuration properties that are stored
	// at the device level, such as enable or disable sending and receiving, and the auto
	// or manual answering of calls. The value for the dwDeviceId parameter can be obtained
	// using the FAX_EnumPorts (section 3.1.4.1.28) method or the FAX_EnumPortsEx (section
	// 3.1.4.1.29) method.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                                  |
	//	|                  VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005               | Access is denied. The caller does not have the required FAX_ACCESS_MANAGE_CONFIG |
	//	|                                              | authorization for this operation.                                                |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BAD_UNIT 0x00000014                    | The fax server cannot find the device specified by the dwDeviceId parameter.     |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057           | This error code is returned under any of the following conditions: § The         |
	//	|                                              | value of dwDeviceId is zero. § The size of the structure pointed at by           |
	//	|                                              | the pPortInfo parameter, or the value of the dwSizeOfStruct field of this        |
	//	|                                              | structure, do not match the correct size for the FAX_PORT_INFO_EXW. § The        |
	//	|                                              | ReceiveMode field of the structure pointed at by pPortInfo does not contain a    |
	//	|                                              | valid value of the FAX_ENUM_DEVICE_RECEIVE_MODE (section 2.2.55) enumeration.    |
	//	|                                              | § The device specified by dwDeviceId is a virtual device (the Flags field        |
	//	|                                              | of the FAX_PORT_INFO (section 2.2.7) structure is set to FPF_VIRTUAL)            |
	//	|                                              | and the ReceiveMode field of the structure pointed at by pPortInfo is            |
	//	|                                              | FAX_DEVICE_RECEIVE_MODE_MANUAL.                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BUFFER_OVERFLOW 0x0000006F             | The length of the lpwstrDescription character string field of the data structure |
	//	|                                              | pointed at by pPortInfo exceeds 253 characters, excluding the length of the NULL |
	//	|                                              | string terminator.                                                               |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7            | The fax server cannot store the updated device information to Registry. The      |
	//	|                                              | Registry might be corrupt.                                                       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_DEVICE_NUM_LIMIT_EXCEEDED 0x00001B62 | The fax server cannot complete the operation because the number of active        |
	//	|                                              | fax devices that are allowed for this version of the operating system was        |
	//	|                                              | exceeded.<168>                                                                   |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetPortEx(context.Context, *SetPortExRequest) (*SetPortExResponse, error)

	// The FAX_EnumPortsEx (Opnum 48) method is called by the client to enumerate detailed
	// port state information for each device that is connected to the fax server.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the access rights  |
	//	|                                    | (FAX_ACCESS_QUERY_CONFIG) required for this operation.                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | Not enough storage is available to process this command. The fax server cannot   |
	//	|                                    | allocate sufficient memory to hold the array of _FAX_PORT_INFO_EXW to be         |
	//	|                                    | returned to the client.                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The Buffer parameter is set to a NULL pointer value. <96>                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F    | The fax server failed to custom marshal the array of _FAX_PORT_INFO_EXW to be    |
	//	|                                    | returned to the client.                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EnumPortsEx(context.Context, *EnumPortsExRequest) (*EnumPortsExResponse, error)

	// The FAX_GetExtensionData (Opnum 49) method is called by the client to retrieve the
	// private configuration data for a fax routing extension or a routing method. Such
	// private configuration data is written with a FAX_SetExtensionData (section 3.1.4.1.79)
	// call. The lpcwstrNameGUID parameter MUST be for a valid routing extension or routing
	// method for which the client requests the private data. The value for the dwDeviceId
	// parameter can be obtained from either the FAX_EnumPorts (section 3.1.4.1.28) or the
	// FAX_EnumPortsEx (section 3.1.4.1.29) methods.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_FILE_NOT_FOUND 0x00000002    | The fax server cannot find the requested data.                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the access rights  |
	//	|                                    | (FAX_ACCESS_QUERY_CONFIG) required for this operation.                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned if any of the following  |
	//	|                                    | conditions are met: <109> § The lpcwstrNameGUID parameter is set to a NULL       |
	//	|                                    | pointer value. § The ppData parameter is set to a NULL pointer value. § The      |
	//	|                                    | lpdwDataSize parameter is set to a NULL pointer value. § The lpcwstrNameGUID     |
	//	|                                    | parameter holds an invalid curly-braced GUID string.                             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7  | The registry is corrupted. The structure of one of the files that contain        |
	//	|                                    | registry data is corrupted, or the system's memory image of the file is          |
	//	|                                    | corrupted, or the file could not be recovered because the alternate copy or log  |
	//	|                                    | was absent or corrupted.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetExtensionData(context.Context, *GetExtensionDataRequest) (*GetExtensionDataResponse, error)

	// The fax client application calls the FAX_SetExtensionData (Opnum 50) method in order
	// to write the private data for a routing extension or routing method for one or all
	// fax devices installed on the fax server. The corresponding method that the fax client
	// calls to read this private data is FAX_GetExtensionData (section 3.1.4.1.39). The
	// value for the dwDeviceId parameter can be obtained using either the FAX_EnumPorts
	// (section 3.1.4.1.28) or FAX_EnumPortsEx (section 3.1.4.1.29) method. The lpcwstrNameGUID
	// parameter MUST be for a valid routing extension or routing method for which the client
	// requests the private data to be written.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the required       |
	//	|                                    | FAX_ACCESS_MANAGE_CONFIG authorization for this operation.                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned if any of the following  |
	//	|                                    | conditions are met:<161> § The lpcwstrComputerName parameter is set to a NULL    |
	//	|                                    | pointer value. § The lpcwstrNameGUID parameter is set to a NULL pointer value.   |
	//	|                                    | § The pData parameter is set to a NULL pointer value. § The dwDataSize parameter |
	//	|                                    | is set to zero. § The lpcwstrNameGUID parameter holds an invalid curly-braced    |
	//	|                                    | GUID string.                                                                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7  | The registry is corrupted. The structure of one of the files that contain        |
	//	|                                    | registry data is corrupted, or the system's memory image of the file is          |
	//	|                                    | corrupted, or the file could not be recovered because the alternate copy or log  |
	//	|                                    | was absent or corrupted.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetExtensionData(context.Context, *SetExtensionDataRequest) (*SetExtensionDataResponse, error)

	// The FAX_AddOutboundGroup (Opnum 51) method is called by the client to add a new outbound
	// routing group.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                                  |
	//	|                  VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005               | Access is denied. The client's fax user account does not have the                |
	//	|                                              | FAX_ACCESS_MANAGE_CONFIG access rights.                                          |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_DUP_NAME 0x00000034                    | The group name specified by the lpwstrGroupName parameter is "< All devices>".   |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057           | The fax server tried to return FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU, but the client |
	//	|                                              | fax API version FAX_API_VERSION_0 (described in section 3.1.4.1.10) does not     |
	//	|                                              | support this error code.                                                         |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BUFFER_OVERFLOW 0x0000006F             | The length of the character string specified by the lpwstrGroupName parameter,   |
	//	|                                              | excluding the length of the terminating null terminator, is equal to or greater  |
	//	|                                              | than 128 characters.                                                             |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7            | The fax server cannot store the new outbound routing group configuration in the  |
	//	|                                              | registry. The registry could be corrupted.                                       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU 0x00001B63 | The fax client module API version (described in section 3.1.4.1.10) is           |
	//	|                                              | FAX_API_VERSION_1 or above, and the fax server is running on a version of the    |
	//	|                                              | operating system that does not support the requested operation.<68>              |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those by the underlying RPC protocol
	// [MS-RPCE].
	AddOutboundGroup(context.Context, *AddOutboundGroupRequest) (*AddOutboundGroupResponse, error)

	// The fax client application calls the FAX_SetOutboundGroup (Opnum 52) method to set
	// a new device list to an existing group. The name of the group to remove is specified
	// using the lpwstrGroupName parameter. The value for lpwstrGroupName can be obtained
	// using FAX_EnumOutboundGroups (section 3.1.4.1.26).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                                  |
	//	|                  VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005               | The client's fax user account does not have the required                         |
	//	|                                              | FAX_ACCESS_MANAGE_CONFIG authorization for this operation.                       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057           | The parameter is incorrect. This error code is returned when any of the          |
	//	|                                              | following conditions occur: § The dwSizeOfStruct member of the pGroup parameter  |
	//	|                                              | is not equal to the correct size of the RPC_FAX_OUTBOUND_ROUTING_GROUPW. §       |
	//	|                                              | The lpwstrGroupName member of pGroup is set to a NULL pointer value. § The       |
	//	|                                              | lpdwDevices member of pGroup is set to NULL and the dwNumDevices member of       |
	//	|                                              | the same pGroup is set to a value greater than zero. § The fax server tried      |
	//	|                                              | to return FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU but the client fax API version       |
	//	|                                              | (FAX_API_VERSION_0, described in section 3.1.4.1.10) does not support this error |
	//	|                                              | code.                                                                            |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BUFFER_OVERFLOW 0x0000006F             | The length (excluding the terminating null character) of the character           |
	//	|                                              | string pointed at by the lpwstrGroupName member of pGroup is longer than 128     |
	//	|                                              | characters.                                                                      |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7            | The registry is corrupted. The structure of one of the files containing registry |
	//	|                                              | data is corrupted, or the system's memory image of the file is corrupted, or     |
	//	|                                              | the file could not be recovered because the alternate copy or log was absent or  |
	//	|                                              | corrupted.                                                                       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU 0x00001B63 | The fax client module API version (described in section 3.1.4.1.10) is           |
	//	|                                              | FAX_API_VERSION_1 or above, and the fax server is running on a version of the    |
	//	|                                              | operating system that does not support the requested operation.<163>             |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetOutboundGroup(context.Context, *SetOutboundGroupRequest) (*SetOutboundGroupResponse, error)

	// The fax client application calls the FAX_RemoveOutboundGroup (Opnum 53) method to
	// remove an existing outbound routing group from the fax server. The name of the group
	// to remove is specified using the lpwstrGroupName parameter. The value for lpwstrGroupName
	// can be obtained using FAX_EnumOutboundGroups (section 3.1.4.1.26).<149>
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                                  |
	//	|                  VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005               | Access is denied. The client's fax user account does not have the                |
	//	|                                              | FAX_ACCESS_MANAGE_CONFIG access rights.                                          |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057           | The fax server tried to return FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU but the client  |
	//	|                                              | fax API version (FAX_API_VERSION_0, described in the section 3.1.4.1.10) does    |
	//	|                                              | not support this error code.                                                     |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BUFFER_OVERFLOW 0x0000006F             | The group name specified by the lpwstrGroupName argument (excluding the length   |
	//	|                                              | of the terminating null character) is longer than 128 characters.                |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7            | The fax server cannot access the local machine's fax routing group information   |
	//	|                                              | in the registry. The registry could be corrupt.                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_OPERATION 0x000010DD           | The lpwstrGroupName parameter specifies the special routing group "<All          |
	//	|                                              | Devices>".                                                                       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_GROUP_NOT_FOUND 0x000001B5A          | The group specified by the lpwstrGroupName argument cannot be found.             |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_GROUP_IN_USE 0x00001B5C              | The fax server cannot remove the outbound routing group identified by            |
	//	|                                              | lpwstrGroupName. The outbound routing group is in use by one or more outbound    |
	//	|                                              | routing rules.                                                                   |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU 0x00001B63 | The fax client module API version (described in the section 3.1.4.1.10) is       |
	//	|                                              | FAX_API_VERSION_1 or above, and the fax server is running on a version of the    |
	//	|                                              | operating system that does not support the requested operation.<150>             |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	RemoveOutboundGroup(context.Context, *RemoveOutboundGroupRequest) (*RemoveOutboundGroupResponse, error)

	// The FAX_EnumOutboundGroups (Opnum 54) method is called by the client.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                                  |
	//	|                  VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005               | Access is denied. The client's fax user account does not have the access rights  |
	//	|                                              | (FAX_ACCESS_QUERY_CONFIG) required for this operation.                           |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057           | This error code is returned under any of the following conditions: § The         |
	//	|                                              | pointer referenced by the ppData argument is NULL.<92> § The fax server tried    |
	//	|                                              | to return FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU, but the client fax API version      |
	//	|                                              | (FAX_API_VERSION_0, described in section 3.1.4.1.10) does not support this error |
	//	|                                              | code.                                                                            |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU 0x00001B63 | The fax client module's API version (described in section 3.1.4.1.10) is         |
	//	|                                              | FAX_API_VERSION_1 or above, and the fax server is running. on a version of the   |
	//	|                                              | operating system that does not support the requested operation.<93>              |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EnumOutboundGroups(context.Context, *EnumOutboundGroupsRequest) (*EnumOutboundGroupsResponse, error)

	// The FAX_SetDeviceOrderInGroup (Opnum 55) method is called by the client. The value
	// for the dwDeviceId parameter can be obtained using the FAX_EnumPorts (section 3.1.4.1.28)
	// method or the FAX_EnumPortsEx (section 3.1.4.1.29) method. The name of the group
	// to remove is specified using the lpwstrGroupName parameter. The value for lpwstrGroupName
	// can be obtained using FAX_EnumOutboundGroups (section 3.1.4.1.26).
	//
	// The order is the 1-based location of the device in the group. The value of 1 indicates
	// the device is ordered first in the group. The order of devices in the group determines
	// the order in which they are used to send outgoing faxes when the group is selected
	// by an outbound routing rule.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                                  |
	//	|                  VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005               | Access is denied. The client's fax user account does not have the required       |
	//	|                                              | FAX_ACCESS_MANAGE_CONFIG authorization for this operation.                       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BUFFER_OVERFLOW 0x0000006F             | The group name is too long. The length of the character string specified by the  |
	//	|                                              | lpwstrGroupName parameter, excluding the length of the null terminator, exceeds  |
	//	|                                              | 128 characters.                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057           | The parameter is incorrect. This error code is returned if any of the following  |
	//	|                                              | conditions are met: § The dwDeviceId or the dwNewOrder parameter is set to a     |
	//	|                                              | value of 0. § The fax server tried to return FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU,  |
	//	|                                              | but the client fax API version (FAX_API_VERSION_0, described in section          |
	//	|                                              | 3.1.4.1.10) does not support this error code.                                    |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7            | The registry is corrupted. The structure of one of the files containing registry |
	//	|                                              | data is corrupted, or the system's memory image of the file is corrupted, or     |
	//	|                                              | the file could not be recovered because the alternate copy or log was absent or  |
	//	|                                              | corrupted.                                                                       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_BAD_GROUP_CONFIGURATION 0x00001B5B   | The device specified by dwDeviceId does not exist in the group identified by     |
	//	|                                              | the lpwstrGroupName, or the value of dwNewOrder is greater than the number of    |
	//	|                                              | devices in the group.                                                            |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU 0x00001B63 | The fax client module API version (as described in section 3.1.4.1.10) is        |
	//	|                                              | FAX_API_VERSION_1 or above, and the fax server is running on a version of the    |
	//	|                                              | operating system that does not support the requested operation.<160>             |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetDeviceOrderInGroup(context.Context, *SetDeviceOrderInGroupRequest) (*SetDeviceOrderInGroupResponse, error)

	// The FAX_AddOutboundRule (Opnum 56) method is called by the client to add a new outbound
	// rule for the specified outbound group to the fax server’s rules map. The value
	// for the dwDeviceId parameter can be obtained using the FAX_EnumPorts (section 3.1.4.1.28)
	// method or the FAX_EnumPortsEx (section 3.1.4.1.29) method.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                                  |
	//	|                  VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005               | Access is denied. The client's fax user account does not have authorization      |
	//	|                                              | required for this call (FAX_ACCESS_MANAGE_CONFIG).                               |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BAD_UNIT 0x00000014                    | The system cannot find the device specified by the dwDeviceId argument.          |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_GEN_FAILURE 0x0000001F                 | The fax server encountered an exception while processing the character string    |
	//	|                                              | specified by the lpwstrGroupName argument.                                       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057           | The parameter is incorrect. This error code is returned when any of              |
	//	|                                              | the following conditions are met: § The country code specified with the          |
	//	|                                              | dwCountryCode argument is ROUTING_RULE_COUNTRY_CODE_ANY (not a valid             |
	//	|                                              | rule dialing location). § The lpwstrGroupName argument value is NULL.            |
	//	|                                              | § The value of the dwDeviceId argument is 0. § The fax server needs to           |
	//	|                                              | return FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU, but the client fax API version         |
	//	|                                              | (FAX_API_VERSION_0, described in section 3.1.4.1.10) does not support this error |
	//	|                                              | code.                                                                            |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BUFFER_OVERFLOW 0x0000006F             | The destination group of the rule specified by the lpwstrGroupName argument      |
	//	|                                              | is longer than the maximum supported value of 128 characters (excluding the      |
	//	|                                              | terminating null character).                                                     |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7            | The registry is corrupted. The structure of one of the files containing registry |
	//	|                                              | data is corrupted, or the system's memory image of the file is corrupted, or     |
	//	|                                              | the file could not be recovered because the alternate copy or log was absent or  |
	//	|                                              | corrupted.                                                                       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_BAD_GROUP_CONFIGURATION 0x00001B5B   | The fax server encountered an outbound routing group with a bad configuration,   |
	//	|                                              | or the group device list is empty; the status for the new rule object created    |
	//	|                                              | by the server based on the specified dialing location and device ID is           |
	//	|                                              | FAX_RULE_STATUS_ALL_GROUP_DEV_NOT_VALID or FAX_RULE_STATUS_EMPTY_GROUP.          |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU 0x00001B63 | The fax client module API version (described in section 3.1.4.1.10) is           |
	//	|                                              | FAX_API_VERSION_1 or above, and the fax server is running on a version of the    |
	//	|                                              | operating system that does not support the requested operation.<69>              |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	AddOutboundRule(context.Context, *AddOutboundRuleRequest) (*AddOutboundRuleResponse, error)

	// The FAX_RemoveOutboundRule (Opnum 57) method removes an existing outbound routing
	// rule from the rules map. The default outbound rule cannot be removed.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                                  |
	//	|                  VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005               | Access is denied. The client's fax user account does not have the required       |
	//	|                                              | FAX_ACCESS_MANAGE_CONFIG access rights to perform this operation.                |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057           | The parameter is incorrect. The country code specified by the dwCountryCode      |
	//	|                                              | argument is ROUTING_RULE_COUNTRY_CODE_ANY. Also returned if the fax server       |
	//	|                                              | tried to return FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU but the client fax API version |
	//	|                                              | (FAX_API_VERSION_0, described in section 3.1.4.1.10) does not support this error |
	//	|                                              | code.                                                                            |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7            | The fax server encountered a registry error when attempting to remove the        |
	//	|                                              | specified outbound rule registration. The registry could be corrupt.             |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_RULE_NOT_FOUND 0x00001B5D            | The fax server failed to locate an outbound routing rule by country/region code  |
	//	|                                              | and area code.                                                                   |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU 0x00001B63 | The fax server is running on a version of the operating system that does not     |
	//	|                                              | support the requested operation.<151>                                            |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	RemoveOutboundRule(context.Context, *RemoveOutboundRuleRequest) (*RemoveOutboundRuleResponse, error)

	// A fax client application uses the FAX_SetOutboundRule (Opnum 58) method to set the
	// information about an individual fax outbound routing rule.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                                  |
	//	|                  VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005               | Access is denied. The client's fax user account does not have the required       |
	//	|                                              | FAX_ACCESS_MANAGE_CONFIG authorization for this operation.                       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057           | The parameter is incorrect. This error code is returned if any of the following  |
	//	|                                              | conditions occur: § The lpwstrGroupName member of the Destination member of      |
	//	|                                              | the pRule parameter is set to a NULL pointer value. § The dwDeviceId member      |
	//	|                                              | of the Destination member of pRule is set to zero. § The fax server tried        |
	//	|                                              | to return FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU but the client fax API version       |
	//	|                                              | (FAX_API_VERSION_0 (described in section 3.1.4.1.10) does not support this error |
	//	|                                              | code.                                                                            |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BUFFER_OVERFLOW 0x0000006F             | The length (excluding the terminating null character) of the character string    |
	//	|                                              | pointed at by the lpwstrGroupName member of the Destination member of pRule is   |
	//	|                                              | greater than 128 characters.                                                     |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7            | The registry is corrupted. The structure of one of the files containing registry |
	//	|                                              | data is corrupted, or the system's memory image of the file is corrupted, or     |
	//	|                                              | the file could not be recovered because the alternate copy or log was absent or  |
	//	|                                              | corrupted.                                                                       |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_BAD_GROUP_CONFIGURATION 0x00001B5B   | This error code is returned if any of the following conditions occur: § The      |
	//	|                                              | fax server encountered an outbound routing group with a bad configuration,       |
	//	|                                              | or the group device list is empty. § The status for the outbound routing rule    |
	//	|                                              | object indicated by the specified dialing location (the dwCountryCode and        |
	//	|                                              | dwAreaCode members of pRule) and group name (the lpwstrGroupName member of       |
	//	|                                              | the Destination member of pRule) is FAX_GROUP_STATUS_ALL_ DEV_NOT_VALID or       |
	//	|                                              | FAX_RULE_STATUS_EMPTY_GROUP.                                                     |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU 0x00001B63 | The fax client module API version (described in section 3.1.4.1.10) is           |
	//	|                                              | FAX_API_VERSION_1 or above, and the fax server is running on a version of the    |
	//	|                                              | operating system that does not support the requested operation.<164>             |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetOutboundRule(context.Context, *SetOutboundRuleRequest) (*SetOutboundRuleResponse, error)

	// The FAX_EnumOutboundRules (Opnum 59) method is called by the client to enumerate
	// all the outbound routing rules that are present on the specified fax server.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                                  |
	//	|                  VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005               | Access is denied. The client's fax user account does not have the access rights  |
	//	|                                              | (FAX_ACCESS_QUERY_CONFIG) required for this operation.                           |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057           | The fax server tried to return FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU but the client  |
	//	|                                              | fax API version (FAX_API_VERSION_0, described in section 3.1.4.1.10) does not    |
	//	|                                              | support this error code.                                                         |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_NOT_SUPPORTED_ON_THIS_SKU 0x00001B63 | The fax client module's API version (described in section 3.1.4.1.10) is         |
	//	|                                              | FAX_API_VERSION_1 or above, and the fax server is running on a version of the    |
	//	|                                              | operating system that does not support the requested operation.<94>              |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EnumOutboundRules(context.Context, *EnumOutboundRulesRequest) (*EnumOutboundRulesResponse, error)

	// The fax client application calls the FAX_RegisterServiceProviderEx (Opnum 60) method
	// to register a fax service provider (FSP) with the Fax Service. Registration takes
	// place after the Fax Service restarts.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the authorization  |
	//	|                                    | for FAX_ACCESS_MANAGE_CONFIG required for this operation.                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned if any of the following  |
	//	|                                    | conditions are met: § The lpcwstrGuid parameter contains an invalid GUID. §      |
	//	|                                    | The dwFSPIVersion parameter is set to a value other than 0x00010000. § The       |
	//	|                                    | dwCapabilities parameter is set to a value other than 0. § The file path         |
	//	|                                    | specified by the lpcwstrImageName parameter does not exist, or the fax server    |
	//	|                                    | does not have access to the file.                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BUFFER_OVERFLOW 0x0000006F   | The length of the lpcwstrFriendlyName, lpcwstrImageName or lpcwstrTspName        |
	//	|                                    | character strings exceeds MAX_FAX_STRING_LEN characters, excluding the length of |
	//	|                                    | the NULL string terminator.                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ALREADY_EXISTS 0x000000B7    | An FSP is already registered with the same GUID (specified by the lpcwstrGUID    |
	//	|                                    | parameter) or with the same telephony service provider (specified by the         |
	//	|                                    | lpcwstrTspName parameter).                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7  | The registry is corrupted. The structure of one of the files that contains       |
	//	|                                    | registry data is corrupted, or the system's memory image of the file is          |
	//	|                                    | corrupted, or the file could not be recovered because the alternate copy or log  |
	//	|                                    | was absent or corrupted.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	RegisterServiceProviderEx(context.Context, *RegisterServiceProviderExRequest) (*RegisterServiceProviderExResponse, error)

	// The FAX_UnregisterServiceProviderEx (Opnum 61) method is called when the client needs
	// to unregister a fax service provider (FSP). In response, the server MUST validate
	// that the client's fax user account has write access. On success, the server MUST
	// remove the service provider for the fax server.<179>
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | Access is denied. The caller does not have the required FAX_ACCESS_MANAGE_CONFIG |
	//	|                                | authorization for this operation.                                                |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	UnregisterServiceProviderEx(context.Context, *UnregisterServiceProviderExRequest) (*UnregisterServiceProviderExResponse, error)

	// The FAX_UnregisterRoutingExtension (Opnum 62) method unregisters an existing inbound
	// routing extension.<178>
	//
	// There are no corresponding routing extension registration functions exposed to the
	// FAX client. Registration is a non-RPC process that is done locally on the fax server
	// using any implementation-specific method.
	//
	// On success, the server MUST unregister the specified routing extension.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return the following error code, one of the fax-specific errors that are
	// defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | Access is denied. The client's fax user account does not have sufficient rights  |
	//	|                                | to perform the operation (FAX_ACCESS_MANAGE_CONFIG) or the user account does not |
	//	|                                | exist.                                                                           |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	UnregisterRoutingExtension(context.Context, *UnregisterRoutingExtensionRequest) (*UnregisterRoutingExtensionResponse, error)

	// The FAX_StartMessagesEnum (Opnum 63) method is called by the client. On success,
	// the server MUST start enumerating messages in one of the archives.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The caller does not have the required                          |
	//	|                                    | ALL_FAX_USERS_ACCESS_RIGHTS access right to execute this call.                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned under any of             |
	//	|                                    | the following conditions: § The value specified for the level argument           |
	//	|                                    | is not equal to 1. § The value specified for the Folder argument is not          |
	//	|                                    | equal to FAX_MESSAGE_FOLDER_INBOX or FAX_MESSAGE_FOLDER_SENTITEMS. § The         |
	//	|                                    | lpcwstrAccountName parameter is not NULL and passes validation checks but does   |
	//	|                                    | not correspond to an existing account name. § The account name specified by the  |
	//	|                                    | lpcwstrAccountName argument is a valid account name but it refers to a different |
	//	|                                    | user than the caller.                                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NO_MORE_ITEMS 0x00000103     | No data is available. There are no messages to be enumerated.                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	StartMessagesEnum(context.Context, *StartMessagesEnumRequest) (*StartMessagesEnumResponse, error)

	// The FAX_EndMessagesEnum (Opnum 64) method is called by the client.
	//
	// On success, the server MUST halt the enumerating of messages in the specified archives.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_HANDLE 0x00000006    | This error code SHOULD be returned if the handle pointed to by the specified     |
	//	|                                    | lpHandle parameter is not a valid handle returned by FAX_StartMessagesEnum.<79>  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | This error SHOULD be returned if the handle pointed to by the specified lpHandle |
	//	|                                    | is NULL.<80>                                                                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EndMessagesEnum(context.Context, *EndMessagesEnumRequest) (*EndMessagesEnumResponse, error)

	// The FAX_EnumMessages (Opnum 65) method is called by the client.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The fax server failed to allocate memory for the return buffer.                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | This error code is returned if any of the following conditions are met: § One    |
	//	|                                    | or more of the pointer values specified by the following arguments are NULL:     |
	//	|                                    | lppBuffer, lpdwBufferSize, and lpdwNumMessagesRetrieved. § The hEnum parameter   |
	//	|                                    | is NULL.<89> § The maximum number of messages, specified by the dwNumMessages    |
	//	|                                    | argument, is set to zero.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BUFFER_OVERFLOW 0x0000006F   | The fax server encountered an integer overflow condition while processing        |
	//	|                                    | the request for the maximum number of messages specified by the dwNumMessages    |
	//	|                                    | argument.                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NO_MORE_ITEMS 0x00000103     | No more data is available. The method reached the end of the lppBuffer message   |
	//	|                                    | buffer and there are no more messages to be enumerated.                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_SRV_OUTOFMEMORY 0x00001B59 | The fax server failed to allocate memory needed for internal execution of the    |
	//	|                                    | command.                                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// The client expects that this method is incremental and uses an internal context cursor
	// to point to the next set of messages to retrieve for each call. The cursor is set
	// to point to the beginning of the messages in the archive after a successful call
	// to FAX_StartMessagesEnum. Each successful call to FAX_EnumMessages advances the cursor
	// by the number of messages retrieved. After the cursor reaches the end of the enumeration,
	// the method fails with the 0x00000103 (ERROR_NO_MORE_ITEMS) error code. The FAX_EndMessagesEnum
	// (section 3.1.4.1.16) method SHOULD then be called.
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EnumMessages(context.Context, *EnumMessagesRequest) (*EnumMessagesResponse, error)

	// The FAX_GetMessage (Opnum 66) method is called by the client. The archive can be
	// one of the enumerations that are defined by FAX_ENUM_MESSAGE_FOLDER (section 2.2.2)
	// except FAX_MESSAGE_FOLDER_QUEUE. The dwlMessageId parameter specifies a particular
	// message and can be obtained using the FAX_EnumMessages (section 3.1.4.1.24) method
	// or the FAX_EnumMessagesEx (section 3.1.4.1.25) method.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005       | Access is denied. The client's fax user account does not have any of the         |
	//	|                                      | permissions covered by ALL_FAX_USER_ACCESS_RIGHTS (section 2.2.83).              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008   | Not enough storage is available to process this command. The fax server failed   |
	//	|                                      | to allocate sufficient memory to hold the FAX_MESSAGEW to be returned to the     |
	//	|                                      | client.                                                                          |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057   | The parameter is incorrect. This error code is returned when any of the          |
	//	|                                      | following conditions are met: § The lppBuffer or lpdwBufferSize parameters       |
	//	|                                      | are set to NULL pointer values.<116> § The dwlMessageId parameter is set         |
	//	|                                      | to a value of 0, and the value of the specified Folder parameter is not          |
	//	|                                      | FAX_MESSAGE_FOLDER_INBOX or FAX_MESSAGE_FOLDER_SENTITEMS.                        |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F      | The fax server failed to custom marshal the FAX_MESSAGEW to be returned to the   |
	//	|                                      | client.                                                                          |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_MESSAGE_NOT_FOUND 0x00001B61 | The fax server cannot find the job or message by its identifier.  This           |
	//	|                                      | error code is returned when any of the following conditions are met:             |
	//	|                                      | § The message identified by dwlMessageId is not found. § The message             |
	//	|                                      | identified by dwlMessageId is an unassigned incoming fax. The incoming           |
	//	|                                      | faxes are not public (accessible to all users), and the user does not            |
	//	|                                      | have FAX_ACCESS_MANAGE_RECEIVE_FOLDER permission. § The message identified       |
	//	|                                      | by dwlMessageId is for a different user, and this user does not have             |
	//	|                                      | FAX_ACCESS_QUERY_ARCHIVES permission.                                            |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error)

	// The fax client application calls the FAX_RemoveMessage (Opnum 67) method to remove
	// a message from a specific Fax Archive Folder. The dwlMessageId parameter specifies
	// a particular message and can be obtained using the FAX_EnumMessages (section 3.1.4.1.24)
	// method or the FAX_EnumMessagesEx (section 3.1.4.1.25) method.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section  2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005        | Access is denied. The caller does not have the required access rights            |
	//	|                                       | (ALL_FAX_USER_ACCESS_RIGHTS (section 2.2.83)) for this operation.                |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057    | This error code is returned under any of the following conditions: § The value   |
	//	|                                       | of the dwlMessageId parameter is 0. § The value of the Folder parameter is not   |
	//	|                                       | FAX_MESSAGE_FOLDER_INBOX or FAX_MESSAGE_FOLDER_SENT_ITEMS.                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_FILE_ACCESS_DENIED 0x00001B60 | The fax server failed to remove the fax message. When trying to delete the fax   |
	//	|                                       | archived file (the file that represents the fax message to be removed), the fax  |
	//	|                                       | server internally encountered an access denied or sharing violation error.       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_MESSAGE_NOT_FOUND 0x00001B61  | The fax server cannot find the message to be deleted (indicated by               |
	//	|                                       | dwlMessageId). When trying to delete the fax archived file (the file that        |
	//	|                                       | represents the fax message to be removed), the fax server internally encountered |
	//	|                                       | a file not found error.                                                          |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	RemoveMessage(context.Context, *RemoveMessageRequest) (*RemoveMessageResponse, error)

	// The client calls the FAX_StartCopyToServer (Opnum 68) method to start a copy of a
	// file to the server queue directory (section 3.1.1) for which the client's fax user
	// account has access to submit faxes. The server MUST generate a unique file name and
	// create a file with that name in the server queue directory. Then the server MUST
	// create a copy handle and return it to the client to indicate success.
	//
	// The copy handle returned by the fax server with the lpHandle output argument is valid
	// until the fax client calls FAX_EndCopy (section 3.1.4.1.15), until the fax server
	// is shut down or restarted, or until an implementation-specific condition occurs that
	// invalidates the copy handle on the fax server.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have any of the access  |
	//	|                                    | rights required for this operation: FAX_ACCESS_SUBMIT, FAX_ACCESS_SUBMIT_NORMAL, |
	//	|                                    | or FAX_ACCESS_SUBMIT_HIGH.                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                    | following conditions: § The lpcwstrFileExt parameter is set to a NULL pointer    |
	//	|                                    | value.<173> § The file extension that is specified by lpcwstrFileExt is not      |
	//	|                                    | ".cov" or ".tif".                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BUFFER_OVERFLOW 0x0000006F   | The file name is too long. The buffer specified by the lpwstrServerFileName      |
	//	|                                    | parameter does not have sufficient space to accommodate the server file name.    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	StartCopyToServer(context.Context, *StartCopyToServerRequest) (*StartCopyToServerResponse, error)

	// The FAX_StartCopyMessageFromServer (Opnum 69) method is called by the fax client
	// to start a copy operation of a fax message from the Fax Archive Folder (section 3.1.1)
	// or of a fax job from the server queue directory (section 3.1.1).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|          RETURN/ERROR CODE           |                                                                                  |
	//	|                VALUES                |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057   | The parameter is incorrect. The dwlMessageId specified is 0 and/or the           |
	//	|                                      | specified Folder enumeration value is not FAX_MESSAGE_FOLDER_QUEUE,              |
	//	|                                      | FAX_MESSAGE_FOLDER_INBOX, or FAX_MESSAGE_FOLDER_SENTITEMS.                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_MESSAGE_NOT_FOUND 0x00001B61 | This error code is returned if any of the following conditions are met: §        |
	//	|                                      | The fax server cannot find the fax queue entry referenced by the specified       |
	//	|                                      | dwMessageId (invalid job identifier) or the user does not have sufficient        |
	//	|                                      | rights to access the fax queue. § FAX_MESSAGE_FOLDER_QUEUE was specified         |
	//	|                                      | for the Folder parameter, the specified message is not owned by the client's     |
	//	|                                      | fax user account, and the client's fax user account does not have either         |
	//	|                                      | the FAX_ACCESS_QUERY_OUT_JOBS permission or the FAX_ACCESS_MANAGE_OUT_JOBS       |
	//	|                                      | permission. § FAX_MESSAGE_FOLDER_INBOX was specified for Folder, the client's    |
	//	|                                      | fax user account does not have the FAX_ACCESS_MANAGE_RECEIVE_FOLDER permission,  |
	//	|                                      | and the bIncomingMessagesArePublic option is not specified in the server         |
	//	|                                      | configuration. § FAX_MESSAGE_FOLDER_SENTITEMS was specified for Folder,          |
	//	|                                      | the client's fax user account does not have the FAX_ACCESS_QUERY_ARCHIVES        |
	//	|                                      | permission, and the specified message is not owned by the client's fax user      |
	//	|                                      | account.                                                                         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	StartCopyMessageFromServer(context.Context, *StartCopyMessageFromServerRequest) (*StartCopyMessageFromServerResponse, error)

	// The FAX_WriteFile (Opnum 70) method is called by the client.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.<181>
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_HANDLE 0x00000006    | The handle value specified by the hCopy argument is not a valid copy handle      |
	//	|                                    | returned by FAX_StartCopyToServer.                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_GEN_FAILURE 0x0000001F       | A device attached to the system is not functioning. The call was unable to write |
	//	|                                    | the full amount of the data that was requested to be written.                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The following conditions can lead to this value being returned: § The parameter  |
	//	|                                    | is incorrect. § The dwDataSize parameter is 0. § The handle value specified by   |
	//	|                                    | the hCopy argument is NULL or the buffer size specified by the dwDataSize is     |
	//	|                                    | zero.                                                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	WriteFile(context.Context, *WriteFileRequest) (*WriteFileResponse, error)

	// The fax client application calls the FAX_ReadFile (Opnum 71) method to copy a file
	// from the server (in chunks).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_HANDLE 0x00000006    | The handle specified by the hCopy argument is not a valid copy handle            |
	//	|                                    | returned by FAX_StartCopyMessageFromServer for which FAX_EndCopy has not been    |
	//	|                                    | called.<142>                                                                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                    | following conditions: § The handle specified by the hCopy argument is NULL.<143> |
	//	|                                    | § The value specified for the dwMaxDataSize argument is zero. § The value        |
	//	|                                    | specified for the *lpdwDataSize at input (when the call is made) is different    |
	//	|                                    | than the value specified for the dwMaxDataSize argument.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	ReadFile(context.Context, *ReadFileRequest) (*ReadFileResponse, error)

	// The FAX_EndCopy (Opnum 72) method is called by the client to end a copy operation
	// process from or to the server, and to close the respective copy handle.
	//
	// On success, the server MUST terminate the specified copy operation previously begun
	// with FAX_StartCopyToServer (section 3.1.4.1.97) or FAX_StartCopyMessageFromServer
	// (section 3.1.4.1.96).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_HANDLE 0x00000006 | This error code SHOULD be returned if the handle pointed to by the specified     |
	//	|                                 | lphCopy parameter is not a valid handle returned by FAX_StartCopyToServer or     |
	//	|                                 | FAX_StartCopyMessageFromServer.<78>                                              |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EndCopy(context.Context, *EndCopyRequest) (*EndCopyResponse, error)

	// The FAX_StartServerNotification (Opnum 73) method is called by the client to get
	// notification about legacy events. On success, the server MUST start to notify the
	// fax client about the occurring fax events.
	//
	// Protocol version FAX_API_VERSION_2 (0x00020000) and FAX_API_VERSION_3 (0x00030000)
	// fax servers SHOULD fail this call by returning ERROR_NOT_SUPPORTED (0x00000032).
	// The fax client SHOULD NOT call this method if the protocol version reported by the
	// server is FAX_API_VERSION_2 (0x00020000) or FAX_API_VERSION_3 (0x00030000). For more
	// information, see FAX_ConnectFaxServer (section 3.1.4.1.10).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | The client's fax user account does not have sufficient rights for this call,     |
	//	|                                | which is ALL_FAX_USER_ACCESS_RIGHTS, or the user account does not exist.         |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BAD_FORMAT 0x0000000B    | The length, including the terminating null character, of the string specified    |
	//	|                                | by the lpcwstrMachineName argument is greater than 256 characters. The length,   |
	//	|                                | including the terminating null character, of the string specified by the         |
	//	|                                | lpcwstrEndPoint argument is greater than 11 characters.                          |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_SUPPORTED 0x00000032 | The request is not supported.                                                    |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	//
	// A fax client calls FAX_StartServerNotification to inform the server that it needs
	// to receive the notifications of legacy fax events. The fax server SHOULD call FAX_OpenConnection
	// on the client by using the supplied endpoint, protocol sequence information, and
	// context handle information. The server then sends the notification of legacy events
	// to the client by using FAX_ClientEventQueue (section 3.2.4.2). When the client no
	// longer needs to receive notifications, it calls FAX_EndServerNotification, and the
	// server SHOULD call FAX_CloseConnection (section 3.2.4.4) to close the connection.
	//
	// Note  This method only supports TCP/IP as the transport protocol.
	StartServerNotification(context.Context, *StartServerNotificationRequest) (*StartServerNotificationResponse, error)

	// The FAX_StartServerNotificationEx (Opnum 74) method is called by the client to get
	// notification about extended or legacy events. On success, the server MUST start to
	// notify the fax client about the occurring fax events.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section  2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. This error code is returned under any of the following         |
	//	|                                    | conditions: § The value specified for the dwEventTypes argument contains         |
	//	|                                    | the FAX_EVENT_TYPE_NEW_CALL and/or FAX_EVENT_TYPE_IN_QUEUE flags and the         |
	//	|                                    | caller cannot access unsigned faxes: incoming faxes are not public and the       |
	//	|                                    | caller does not have the FAX_ACCESS_MANAGE_RECEIVE_FOLDER rights. § The value    |
	//	|                                    | specified for the dwEventTypes argument contains the FAX_EVENT_TYPE_CONFIG,      |
	//	|                                    | FAX_EVENT_TYPE_DEVICE_STATUS and/or the FAX_EVENT_TYPE_ACTIVITY flags and the    |
	//	|                                    | caller does not have the FAX_ACCESS_QUERY_CONFIG rights.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BAD_FORMAT 0x0000000B        | This error code is returned under any of the following conditions: § The length  |
	//	|                                    | of the fax client machine name specified by the lpcwstrMachineName argument,     |
	//	|                                    | excluding the length of the terminating null character, is longer than 256       |
	//	|                                    | characters. § The length of the endpoint string specified by the lpcwstrEndPoint |
	//	|                                    | argument, excluding the length of the terminating null character, is longer than |
	//	|                                    | or equal to 11 characters.                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_OUTOFMEMORY 0x0000000E       | Not enough storage is available to complete this operation.                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | Invalid parameter. This error code is returned under any of the following        |
	//	|                                    | conditions: § Any of these arguments specify a NULL pointer value:               |
	//	|                                    | lcpwstrEndPoint, lpcwstrMachineName, lpHandle.<175> § The value specified for    |
	//	|                                    | the dwEventTypes argument is not a bitwise OR combination of the following       |
	//	|                                    | FAX_ENUM_EVENT_TYPE values: § FAX_EVENT_TYPE_IN_QUEUE § FAX_EVENT_TYPE_OUT_QUEUE |
	//	|                                    | § FAX_EVENT_TYPE_CONFIG § FAX_EVENT_TYPE_ACTIVITY § FAX_EVENT_TYPE_QUEUE_STATE   |
	//	|                                    | § FAX_EVENT_TYPE_IN_ARCHIVE § FAX_EVENT_TYPE_OUT_ARCHIVE §                       |
	//	|                                    | FAX_EVENT_TYPE_FXSSVC_ENDED § FAX_EVENT_TYPE_DEVICE_STATUS §                     |
	//	|                                    | FAX_EVENT_TYPE_NEW_CALL                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	//
	// A fax client calls FAX_StartServerNotificationEx to inform the server that it needs
	// to receive the notifications of extended or legacy fax events. The fax server SHOULD
	// call FAX_OpenConnection on the client by using the supplied endpoint, protocol sequence
	// information, and context handle information. The server then sends notification of
	// events to the client by using either FAX_ClientEventQueueEx (section 3.2.4.3) or
	// FAX_ClientEventQueue (section 3.2.4.2) as specified by the bEventEx parameter. When
	// the client no longer needs to receive notifications, it calls FAX_EndServerNotification;
	// the server SHOULD call FAX_CloseConnection (section 3.2.4.4) to close the connection.
	StartServerNotificationEx(context.Context, *StartServerNotificationExRequest) (*StartServerNotificationExResponse, error)

	// The FAX_EndServerNotification (Opnum 75) method is called by the client to stop the
	// notifications from the server, which were initiated by a call to FAX_StartServerNotification
	// (section 3.1.4.1.100), FAX_StartServerNotificationEx (section 3.1.4.1.101), or FAX_StartServerNotificationEx2
	// (section 3.1.4.1.102).
	//
	// On success, the server MUST stop notifying the client of events.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | This error SHOULD be returned if the lpHandle parameter is not a valid           |
	//	|                                    | handle obtained using the FAX_StartServerNotification method, the                |
	//	|                                    | FAX_StartServerNotificationEx method, or the FAX_StartServerNotificationEx2      |
	//	|                                    | method.<81>                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | lpHandle is specified as NULL.<82>                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// To stop notifications, the client SHOULD call FAX_EndServerNotification; the server
	// SHOULD call FAX_CloseConnection section 3.2.4.4) to close the connection.
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EndServerNotification(context.Context, *EndServerNotificationRequest) (*EndServerNotificationResponse, error)

	// The fax client application calls the FAX_GetServerActivity (Opnum 76) method to retrieve
	// the status of the fax queue activity and event log reports.
	//
	// The client MUST allocate memory for the pServerActivity argument. It MUST also set
	// the dwSizeOfStruct field to the correct size, in bytes, of the FAX_SERVER_ACTIVITY
	// structure as described in section 2.2.19.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the                |
	//	|                                    | FAX_ACCESS_QUERY_CONFIG access rights required for this operation.               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. The dwSizeOfStruct member of the FAX_SERVER_ACTIVITY |
	//	|                                    | specified by the pServerActivity parameter on input is set by the client to an   |
	//	|                                    | incorrect value. For more details about the correct size to be filled in this    |
	//	|                                    | member, see the FAX_SERVER_ACTIVITY.                                             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetServerActivity(context.Context, *GetServerActivityRequest) (*GetServerActivityResponse, error)

	// The FAX_SetConfigWizardUsed (Opnum 77) method is called by the client. The server
	// MUST validate that the client's fax user account has access to manage configuration
	// information on the server. On success, the server MUST set a value in the registry
	// indicating whether or not the configuration wizard was used. <159>
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	|              RETURN               |                                                                                  |
	//	|            VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                   |                                                                                  |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005    | Access is denied. The caller does not have the required rights                   |
	//	|                                   | (FAX_ACCESS_MANAGE_CONFIG) to perform this operation.                            |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7 | The registry is corrupted. The structure of one of the files containing registry |
	//	|                                   | data is corrupted, or the system's memory image of the file is corrupted.        |
	//	+-----------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetConfigWizardUsed(context.Context, *SetConfigWizardUsedRequest) (*SetConfigWizardUsedResponse, error)

	// The FAX_EnumRoutingExtensions (Opnum 78) function is called by the client to enumerate
	// all the routing extensions that are registered with the specified fax server. The
	// function returns detailed information about each of the routing extensions.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the access rights  |
	//	|                                    | (FAX_ACCESS_QUERY_CONFIG) required for this operation.                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | Not enough storage is available to process this command. The fax server cannot   |
	//	|                                    | allocate sufficient memory to hold the array of FAX_ROUTING_EXTENSION_INFO       |
	//	|                                    | structures to be returned to the client.                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The Buffer parameter is set to a NULL pointer value. <97>                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F    | The fax server failed to custom marshal the array of FAX_ROUTING_EXTENSION_INFO  |
	//	|                                    | structures to be returned to the client.                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EnumRoutingExtensions(context.Context, *EnumRoutingExtensionsRequest) (*EnumRoutingExtensionsResponse, error)

	// The FAX_ConnectFaxServer (Opnum 80) method is called by the client to create a connection
	// to the fax server.
	//
	// Protocol version FAX_API_VERSION_0 (0x00000000) fax servers SHOULD NOT implement
	// this method.
	//
	// If the underlying RPC layer fails this call by returning RPC_S_PROCNUM_OUT_OF_RANGE
	// (0x000006D1), the fax client SHOULD consider the server protocol (and API version)
	// to be FAX_API_VERSION_0 and MAY retry this request by switching to the FaxObs Server
	// Interface (section 3.1.4.2) and calling the FaxObs_ConnectionRefCount (section 3.1.4.2.2)
	// method.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return the following error code, one of the fax-specific errors that are
	// defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | Access is denied. The bAutoCreateAccountOnConnect field of the                   |
	//	|                                | FAX_GENERAL_CONFIG structure is set to FALSE and the calling user's              |
	//	|                                | authenticated user identity does not have a fax user account associated on       |
	//	|                                | the fax server, or the does not have any of the access rights defined in         |
	//	|                                | ALL_FAX_USER_ACCESS_RIGHTS (section 2.2.83).                                     |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	ConnectFaxServer(context.Context, *ConnectFaxServerRequest) (*ConnectFaxServerResponse, error)

	// The FAX_GetSecurityEx (Opnum 81) method is called by the clients to retrieve information
	// about the fax security descriptor from the fax server.<134>
	//
	// Protocol version FAX_API_VERSION_3 (0x00030000) fax servers SHOULD fail this call
	// by returning ERROR_NOT_SUPPORTED (0x00000032). The fax client SHOULD NOT call this
	// method if the protocol version reported by the server is FAX_API_VERSION_3 (0x00030000).
	// For more information, see FAX_ConnectFaxServer (section 3.1.4.1.10). The fax client
	// SHOULD call the FAX_GetSecurityEx2 (section 3.1.4.1.60) method instead.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                                  |
	//	|               VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005          | Access is denied. This error code is returned if any of the following            |
	//	|                                         | conditions occur: § The client's fax user account does not have READ_CONTROL     |
	//	|                                         | access but the requesting SecurityInformation parameter contains one             |
	//	|                                         | of these flags: GROUP_SECURITY_INFORMATION, DACL_SECURITY_INFORMATION,           |
	//	|                                         | or OWNER_SECURITY_INFORMATION. § The client's fax user account does not          |
	//	|                                         | have ACCESS_SYSTEM_SECURITY but the SecurityInformation contains the flag        |
	//	|                                         | SACL_SECURITY_INFORMATION.                                                       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008      | Not enough storage is available to process this command.                         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057      | The parameter is incorrect. This is returned when the pSecurityDescriptor        |
	//	|                                         | parameter is NULL.                                                               |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_SECURITY_DESCR 0x0000053A | The security descriptor structure is invalid.                                    |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetSecurityEx(context.Context, *GetSecurityExRequest) (*GetSecurityExResponse, error)

	// A fax client application calls the FAX_RefreshArchive (Opnum 82) method to notify
	// the server that the archive folder has been changed and SHOULD be refreshed.<145>
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The required access level SHOULD be FAX_ACCESS_MANAGE_CONFIG.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. The folder parameter SHOULD either be                |
	//	|                                    | FAX_MESSAGE_FOLDER_INBOX or FAX_MESSAGE_FOLDER_SENTITEMS.<147>                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	RefreshArchive(context.Context, *RefreshArchiveRequest) (*RefreshArchiveResponse, error)

	// The FAX_SetRecipientsLimit (Opnum 83) method is called by the client. A fax client
	// application calls FAX_SetRecipientsLimit to set the recipient limit of a single broadcast
	// job. On success, the server MUST set the recipient limit of a single broadcast job.
	//
	// Protocol version FAX_API_VERSION_0 (0x00000000) and FAX_API_VERSION_1 (0x00010000)
	// fax servers SHOULD NOT implement this call. Protocol version FAX_API_VERSION_2 (0x00020000)
	// and FAX_API_VERSION_3 (0x00030000) fax servers SHOULD fail this call by returning
	// ERROR_NOT_SUPPORTED (0x00000032). The fax client MUST NOT call this method if the
	// protocol version reported by the server is FAX_API_VERSION_0 (0x00000000) or FAX_API_VERSION_1
	// (0x00010000). For more information,  see FAX_ConnectFaxServer (section 3.1.4.1.10).
	//
	// Return Values: This method SHOULD return 0x00000032 (ERROR_NOT_SUPPORTED).
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	SetRecipientsLimit(context.Context, *SetRecipientsLimitRequest) (*SetRecipientsLimitResponse, error)

	// The FAX_GetRecipientsLimit (Opnum 84) method is called by the client to retrieve
	// information about the recipient limit of a single broadcast job.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return the following error code, one of the fax-specific errors that are
	// defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | Access is denied. The client's fax user account does not have any of the         |
	//	|                                | following access rights: FAX_ACCESS_SUBMIT, FAX_ACCESS_SUBMIT_NORMAL, or         |
	//	|                                | FAX_ACCESS_SUBMIT_HIGH.                                                          |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetRecipientsLimit(context.Context, *GetRecipientsLimitRequest) (*GetRecipientsLimitResponse, error)

	// The FAX_GetServerSKU (Opnum 85) method is called by the client. In response, the
	// server returns the SKU of the fax server operating system.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return the following error code, one of the fax-specific errors that are
	// defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005 | Access is denied. The client's fax user account does not have any of the         |
	//	|                                | permissions covered by ALL_FAX_USER_ACCESS_RIGHTS (section 2.2.83).              |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetServerSKU(context.Context, *GetServerSKURequest) (*GetServerSKUResponse, error)

	// The FAX_CheckValidFaxFolder (Opnum 86) method is called by the client to check whether
	// the specified path is accessible to the fax server.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN                |                                                                                  |
	//	|             VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_FILE_NOT_FOUND 0x00000002     | The path specified by the lpcwstrPath argument has a valid structure (the folder |
	//	|                                     | path is valid), but the file does not exist.                                     |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_PATH_NOT_FOUND 0x00000003     | The path specified by the lpcwstrPath argument has a valid structure, but the    |
	//	|                                     | folder path does not exist.                                                      |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005      | Access is denied. The client's fax user account does not have sufficient rights  |
	//	|                                     | for this operation (ALL_FAX_USER_ACCESS_RIGHTS).                                 |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057  | The lpcwstrPath argument is NULL, or the path specified by the lpcwstrPath       |
	//	|                                     | argument is incomplete.                                                          |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BUFFER_OVERFLOW 0x0000006F    | The length of the path (including length of the terminating null character)      |
	//	|                                     | specified by the lpcwstrPath argument exceeds 180 characters.                    |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_DIRECTORY_IN_USE 0x00001B5F | The path specified by the lpcwstrPath argument points to a folder currently in   |
	//	|                                     | use by the fax server, such as the server queue directory or the Fax Archive     |
	//	|                                     | Folder (section 3.1.1).                                                          |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	CheckValidFaxFolder(context.Context, *CheckValidFaxFolderRequest) (*CheckValidFaxFolderResponse, error)

	// The FAX_GetJobEx2 (Opnum 87) method is called by the client to retrieve information
	// about a specified job. The job is identified by the job message ID. The job message
	// ID can be obtained using one of the following methods: FAX_EnumJobs (section 3.1.4.1.21),
	// FAX_EnumJobsEx (section 3.1.4.1.22), or FAX_EnumJobsEx2 (section 3.1.4.1.23).
	//
	// Protocol version FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000),
	// and FAX_API_VERSION_2 (0x00020000) fax servers SHOULD NOT implement this call. The
	// fax client MUST NOT call this method if the protocol version reported by the server
	// is FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000), or FAX_API_VERSION_2
	// (0x00020000). For more information, see FAX_ConnectFaxServer (section 3.1.4.1.10).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005       | Access is denied. This error code is returned under any of the following         |
	//	|                                      | conditions: § The client's fax user account does not have any of the permissions |
	//	|                                      | covered by ALL_FAX_USER_ACCESS_RIGHTS (section 2.2.83). § For an outgoing fax    |
	//	|                                      | job, the caller is not the owner of the fax job, and the caller does not have    |
	//	|                                      | the FAX_ACCESS_MANAGE_OUT_JOBS rights. § For an incoming fax job, the caller is  |
	//	|                                      | not the receiver of the call, incoming faxes are not public, and the client's    |
	//	|                                      | fax user account does not have the FAX_ACCESS_MANAGE_RECEIVE_FOLDER rights.      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008   | Not enough storage is available to process this command. The fax server failed   |
	//	|                                      | to allocate sufficient memory to hold the FAX_JOB_ENTRY_EX_1 to be returned to   |
	//	|                                      | the client.                                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057   | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                      | following conditions: § The Buffer and/or BufferSize parameters are set to NULL  |
	//	|                                      | pointer values.<112> § The level parameter is set to a value other than 1.       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F      | The fax server failed to custom marshal the FAX_JOB_ENTRY_EX_1 to be returned to |
	//	|                                      | the client.                                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_MESSAGE_NOT_FOUND 0x00001B61 | This error code is returned under any of the following conditions: § The fax     |
	//	|                                      | server cannot find the fax job identified by the value of the dwlMessageID       |
	//	|                                      | parameter. § The user is not the owner of the fax job identified by the value of |
	//	|                                      | dwlMessageID.                                                                    |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetJobEx2(context.Context, *GetJobEx2Request) (*GetJobEx2Response, error)

	// The FAX_EnumJobsEx2 (Opnum 88) method is called by the client to enumerate a specified
	// set of jobs on the server's queue for a specific fax account. The type of jobs to
	// enumerate is described by the dwJobTypes argument.
	//
	// Protocol version FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000),
	// and FAX_API_VERSION_2 (0x00020000) fax servers SHOULD NOT implement this call. The
	// fax client MUST NOT call this method if the protocol version reported by the server
	// is FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000), or FAX_API_VERSION_2
	// (0x00020000). For more information, see FAX_ConnectFaxServer (section 3.1.4.1.10).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. This error can be returned when any of the following           |
	//	|                                    | conditions are true: § The client's fax user account does not have any of the    |
	//	|                                    | access rights defined in ALL_FAX_USER_ACCESS_RIGHTS (section 2.2.83) that are    |
	//	|                                    | required in order to enumerate jobs of type JT_SEND on its own account. That     |
	//	|                                    | is, the fAllAccounts parameter is FALSE. § The client's fax user account does    |
	//	|                                    | not have the FAX_ACCESS_QUERY_OUT_JOBS access right that is required in order to |
	//	|                                    | enumerate jobs of type JT_SEND on all accounts. That is, fAllAccounts is TRUE. § |
	//	|                                    | The client's fax user account does not have the FAX_ACCESS_MANAGE_RECEIVE_FOLDER |
	//	|                                    | access right that is required in order to enumerate jobs of type JT_RECEIVE or   |
	//	|                                    | JT_ROUTING.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | Not enough storage is available to process this command.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This value is returned when any of the following     |
	//	|                                    | conditions are true: § Either the lpwdJobs or the Buffer parameter is NULL.      |
	//	|                                    | § The BufferSize parameter is 0. § The level parameter is not set to 1. § The    |
	//	|                                    | lpcwstrAccountName parameter contains an improperly formatted account name or    |
	//	|                                    | points to a nonexistent or other user account.                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// The account name that lpcwstrAccountName indicates MUST be in one of the following
	// formats. Any other format is invalid.
	//
	//	+----------------------------+-----------------------------------------------------------------+
	//	|                            |                                                                 |
	//	|           FORMAT           |                           DESCRIPTION                           |
	//	|                            |                                                                 |
	//	+----------------------------+-----------------------------------------------------------------+
	//	+----------------------------+-----------------------------------------------------------------+
	//	| <machine_name>\<user_name> | For a local user with machine_name as the local machine's name. |
	//	+----------------------------+-----------------------------------------------------------------+
	//	| <domain_name>\<user_name>  | For a nonlocal user.                                            |
	//	+----------------------------+-----------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumJobsEx2(context.Context, *EnumJobsEx2Request) (*EnumJobsEx2Response, error)

	// The FAX_GetMessageEx (Opnum 89) method is called by the client to retrieve a particular
	// message from one of the specified fax message archives. The dwlMessageId parameter
	// specifies a particular message and can be obtained using the FAX_EnumMessages (section
	// 3.1.4.1.24) method or the FAX_EnumMessagesEx (section 3.1.4.1.25) method. The folder
	// value MUST be one of the enumerations that are defined by FAX_ENUM_MESSAGE_FOLDER
	// (section 2.2.2), except FAX_MESSAGE_FOLDER_QUEUE. This is an extended version of
	// FAX_GetMessage (section 3.1.4.1.45), because it takes an additional level parameter
	// supporting the extended structure FAX_MESSAGE_1 (section 2.2.37).
	//
	// Protocol version FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1  (0x00010000),
	// and  FAX_API_VERSION_2 (0x00020000) fax servers SHOULD NOT implement this call.
	// The fax client MUST NOT call this method if the protocol version reported by the
	// server is FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000), or FAX_API_VERSION_2
	// (0x00020000). For more information, see FAX_ConnectFaxServer (section 3.1.4.1.10).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005       | Access is denied. The client's fax user account does not have any of the         |
	//	|                                      | permissions covered by ALL_FAX_USER_ACCESS_RIGHTS (section 2.2.83).              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057   | The parameter is incorrect. This error code results under any of the following   |
	//	|                                      | conditions: § The value of the specified level parameter is not 1. § The pointer |
	//	|                                      | specified by the lppBuffer parameter is NULL.<117> § The pointer specified       |
	//	|                                      | by the lpdwBufferSize parameter is NULL.<118> § The value of the dwlMessageId    |
	//	|                                      | parameter is specified as 0. § The value of the specified Folder parameter is    |
	//	|                                      | not FAX_MESSAGE_FOLDER_INBOX or FAX_MESSAGE_FOLDER_SENTITEMS.                    |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_SRV_OUTOFMEMORY 0x00001B59   | The fax server failed to allocate memory needed for internal execution of this   |
	//	|                                      | operation.                                                                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_MESSAGE_NOT_FOUND 0x00001B61 | This error code is returned under any of the following conditions: § The         |
	//	|                                      | message identified by the dwlMessageId parameter is not found. § The message     |
	//	|                                      | identified by dwlMessageId is an unassigned incoming fax. The incoming           |
	//	|                                      | faxes are not public (accessible to all users), and the user does not            |
	//	|                                      | have FAX_ACCESS_MANAGE_RECEIVE_FOLDER permission. § The message identified       |
	//	|                                      | by dwlMessageId is for a different user, and this user does not have             |
	//	|                                      | FAX_ACCESS_QUERY_ARCHIVES permission.                                            |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetMessageEx(context.Context, *GetMessageExRequest) (*GetMessageExResponse, error)

	// The FAX_StartMessagesEnumEx (Opnum 90) method is called by the client. On success,
	// the server MUST start enumerating messages in the specified archive folder.
	//
	// Protocol version FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000),
	// and FAX_API_VERSION_2 (0x00020000) fax servers SHOULD NOT implement this call. The
	// fax client MUST NOT call this method if the protocol version reported by the server
	// is FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000), or FAX_API_VERSION_2
	// (0x00020000). For more information, see FAX_ConnectFaxServer (section 3.1.4.1.10).
	//
	// A fax client application calls the FAX_StartMessagesEnumEx to start enumerating messages
	// from the archives. Each enumerated message has more information than those that are
	// returned by the FAX_StartMessagesEnum (section 3.1.4.1.98) method, namely whether
	// or not the message has a cover page, the type of receipts selected, the email address
	// for receipts, and the flags from FAX_ENUM_MSG_FLAGS (section 2.2.53) enumeration.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section  2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. This error code is returned under any of the following         |
	//	|                                    | conditions: § The caller does not have the required basic access rights to       |
	//	|                                    | execute this call (ALL_FAX_USERS_ACCESS_RIGHTS). § The client's fax user account |
	//	|                                    | does not have access to query messages for all accounts. The value specified for |
	//	|                                    | the fAllAccounts parameter is not equal to zero, and the caller does not have    |
	//	|                                    | the FAX_ACCESS_QUERY_ARCHIVES rights.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                    | following conditions: § The value specified for the level argument is not        |
	//	|                                    | equal to 1. § The value specified for the Folder argument is not equal to        |
	//	|                                    | FAX_MESSAGE_FOLDER_INBOX or FAX_MESSAGE_FOLDER_SENTITEMS. § The account name     |
	//	|                                    | specified for the lpcwstrAccountName argument appears valid (pointer is not      |
	//	|                                    | NULL), but the account name is not a valid fax account name. § The account name  |
	//	|                                    | specified by the lpcwstrAccountName argument is a valid account name, but it     |
	//	|                                    | refers to a different user than the caller.                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NO_MORE_ITEMS 0x00000103     | No data is available. There are no messages to be enumerated.                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	//
	// A fax client application calls the FAX_StartMessagesEnumEx function to start enumerating
	// messages in one of the archives. The enumerated messages have more information than
	// those that are returned by FAX_StartMessagesEnum, namely, whether it has a cover
	// page, the type of receipts selected, the email address for receipts, and the flags
	// from FAX_ENUM_MSG_FLAGS.
	//
	// The account name that lpcwstrAccountName indicates MUST be in one of the following
	// formats. Any other format is invalid.
	//
	//	+----------------------------+---------------------------------------------------------------------+
	//	|                            |                                                                     |
	//	|           FORMAT           |                             DESCRIPTION                             |
	//	|                            |                                                                     |
	//	+----------------------------+---------------------------------------------------------------------+
	//	+----------------------------+---------------------------------------------------------------------+
	//	| <machine_name>\<user_name> | For a local user that has machine_name as the local machine's name. |
	//	+----------------------------+---------------------------------------------------------------------+
	//	| <domain_name>\<user_name>  | For a nonlocal user.                                                |
	//	+----------------------------+---------------------------------------------------------------------+
	StartMessagesEnumEx(context.Context, *StartMessagesEnumExRequest) (*StartMessagesEnumExResponse, error)

	// The FAX_EnumMessagesEx (Opnum 91) method is called by the client. This message differs
	// from the FAX_EnumMessages (section 3.1.4.1.24) in that this function takes a level
	// parameter, which differentiates the type of message information structure that the
	// function returns.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The fax server failed to allocate memory for the return buffer.                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned if any of                |
	//	|                                    | the following conditions are met: § One or more of the pointer values            |
	//	|                                    | specified by the following arguments are NULL: lppBuffer, lpdwBufferSize,        |
	//	|                                    | lpdwNumMessagesRetrieved, and lpdwLevel. § hEnum is NULL<90> or is an            |
	//	|                                    | invalid handle that is not returned by a call to FAX_StartMessagesEnum or        |
	//	|                                    | FAX_StartMessagesEnumEx <91>. § dwNumMessages is zero.                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BUFFER_OVERFLOW 0x0000006F   | The fax server encountered an integer overflow condition while processing        |
	//	|                                    | the request for the maximum number of messages specified by the dwNumMessages    |
	//	|                                    | argument.                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NO_MORE_ITEMS 0x00000103     | No more data is available. The method reached the end of the lppBuffer message   |
	//	|                                    | buffer and there are no more messages to be enumerated.                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_SRV_OUTOFMEMORY 0x00001B59 | The fax server failed to allocate memory needed for internal execution of the    |
	//	|                                    | command.                                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// The client implementation assumes that this method is incremental and uses an internal
	// context cursor to point to the next set of messages to retrieve for each call. The
	// cursor is set to point to the beginning of the messages in the archive after a successful
	// call to FAX_StartMessagesEnum or FAX_StartMessagesEnumEx. Each successful call to
	// FAX_EnumMessagesEx advances the cursor by the number of messages retrieved. After
	// the cursor reaches the end of the enumeration, the method fails with the 0x00000103
	// (ERROR_NO_MORE_ITEMS) error code. The FAX_EndMessagesEnum (section 3.1.4.1.16) method
	// can then be called to halt the enumeration of messages.
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EnumMessagesEx(context.Context, *EnumMessagesExRequest) (*EnumMessagesExResponse, error)

	// The FAX_StartServerNotificationEx2 (Opnum 92) method is called by the client to get
	// notification about extended events. On success, the server MUST start to notify the
	// fax client about the occurring fax events.
	//
	// Protocol version FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000),
	// and FAX_API_VERSION_2 (0x00020000) fax servers SHOULD NOT implement this call. The
	// fax client MUST NOT call this method if the protocol version reported by the server
	// is FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000), or FAX_API_VERSION_2
	// (0x00020000). For more information, see FAX_ConnectFaxServer (section 3.1.4.1.10).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. This error is returned when any of the following conditions    |
	//	|                                    | occur: § The dwEventTypes parameter is set to a value containing the             |
	//	|                                    | FAX_EVENT_TYPE_NEW_CALL or FAX_EVENT_TYPE_IN_QUEUE flags, the incoming faxes are |
	//	|                                    | not public (accessible to all users), and the client's fax user account does not |
	//	|                                    | have the FAX_ACCESS_MANAGE_RECEIVE_FOLDER permission. § dwEventTypes is set to   |
	//	|                                    | a value containing the FAX_EVENT_TYPE_CONFIG, FAX_EVENT_TYPE_DEVICE_STATUS, or   |
	//	|                                    | FAX_EVENT_TYPE_ACTIVITY flags and the client's fax user account does not have    |
	//	|                                    | the FAX_ACCESS_QUERY_CONFIG permission.                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_OUTOFMEMORY 0x0000000E       | The fax server failed to allocate the memory required for the internal server's  |
	//	|                                    | execution of this operation request.                                             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_GEN_FAILURE 0x0000001F       | The server threw internally an exception during the execution of this operation, |
	//	|                                    | and the server handled this exception.                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                    | following conditions: § dwEventTypes is set to a value containing the            |
	//	|                                    | FAX_EVENT_TYPE_LEGACY or FAX_EVENT_TYPE_LOCAL_ONLY flags, or to another          |
	//	|                                    | value that is not a combination made exclusively from the flags valid            |
	//	|                                    | for this operation: FAX_EVENT_TYPE_IN_QUEUE, FAX_EVENT_TYPE_OUT_QUEUE,           |
	//	|                                    | FAX_EVENT_TYPE_CONFIG, FAX_EVENT_TYPE_ACTIVITY, FAX_EVENT_TYPE_QUEUE_STATE,      |
	//	|                                    | FAX_EVENT_TYPE_IN_ARCHIVE, FAX_EVENT_TYPE_OUT_ARCHIVE,                           |
	//	|                                    | FAX_EVENT_TYPE_FXSSVC_ENDED, FAX_EVENT_TYPE_DEVICE_STATUS, or                    |
	//	|                                    | FAX_EVENT_TYPE_NEW_CALL. § The level parameter is not set to 1. § One or more    |
	//	|                                    | of the following parameters are set to NULL pointer values: lpcwstrEndpoint,     |
	//	|                                    | lpcwstrMachineName, or lpHandle.<177> § The lpcwstrAccountName parameter is set  |
	//	|                                    | to a non-null character string pointer value which does not specify a valid fax  |
	//	|                                    | account name. § lpcwstrAccountName is set to a non-null character string pointer |
	//	|                                    | value which specifies a valid fax account name for a different user than the     |
	//	|                                    | user who is currently logged on the client.                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	//
	// The account name is the one on which to listen for events and a level that specifies
	// the type of the structure that describes each event. The name lpcwstrAccountName
	// is accessed only for account-based events.
	//
	// The account name that lpcwstrAccountName indicates MUST be in one of the following
	// formats. Any other format is invalid.
	//
	//	+----------------------------+--------------------------------------------------------------------------+
	//	|                            |                                                                          |
	//	|           FORMAT           |                               DESCRIPTION                                |
	//	|                            |                                                                          |
	//	+----------------------------+--------------------------------------------------------------------------+
	//	+----------------------------+--------------------------------------------------------------------------+
	//	| <machine_name>\<user_name> | For a local user that has machine_name as the name of the local machine. |
	//	+----------------------------+--------------------------------------------------------------------------+
	//	| <domain_name>\<user_name>  | For a nonlocal user.                                                     |
	//	+----------------------------+--------------------------------------------------------------------------+
	//
	// A fax client calls FAX_StartServerNotificationEx2 (section 3.1.4.1.102) to inform
	// the server that it needs to receive notifications of extended fax events. The fax
	// server SHOULD call FAX_OpenConnection on the client by using the supplied endpoint,
	// protocol sequence information, and context handle information. The server then sends
	// notification of events to the client by using FAX_ClientEventQueueEx (section 3.2.4.3).
	// When the client no longer needs to receive notifications, it calls FAX_EndServerNotification
	// (section 3.1.4.1.17), and the server SHOULD call FAX_CloseConnection (section 3.2.4.4)
	// to close the connection.
	StartServerNotificationEx2(context.Context, *StartServerNotificationEx2Request) (*StartServerNotificationEx2Response, error)

	// The FAX_CreateAccount (Opnum 93) method is called by the client to request a new
	// fax user account to be created based on an existing valid operating system user account.
	//
	// Protocol version FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000),
	// and FAX_API_VERSION_2 (0x00020000) fax servers SHOULD NOT implement this call. The
	// fax client MUST NOT call this method if the protocol version reported by the server
	// is FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000), or FAX_API_VERSION_2
	// (0x00020000). For more information, see FAX_ConnectFaxServer (section 3.1.4.1.10).
	//
	// Return Values: This method MUST return 0 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors that are defined
	// in [MS-ERREF] section 2.2.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005        | Access is denied. The client's fax user account does not have the                |
	//	|                                       | FAX_ACCESS_MANAGE_CONFIG access right.                                           |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057    | The parameter is incorrect. This value is returned when any of the following     |
	//	|                                       | conditions are true: § The Buffer parameter is NULL. § The BufferSize parameter  |
	//	|                                       | is 0. § The level parameter is greater than zero. § The account name, as pointed |
	//	|                                       | to by the account information contained in Buffer, is NULL or is specified using |
	//	|                                       | an invalid format.                                                               |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ALREADY_EXISTS 0x000000B7       | The fax account already exists.                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_HANDLE_STATE 0x00000649 | The handle is in an invalid state.                                               |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//
	// The account name contained in the lpcwstrAccountName member of the FAX_ACCOUNT_INFO_0,
	// as pointed to by the Buffer parameter, MUST be in one of the following formats. Any
	// other format is invalid.
	//
	//	+----------------------------+-----------------------------------------------------------------+
	//	|                            |                                                                 |
	//	|           FORMAT           |                           DESCRIPTION                           |
	//	|                            |                                                                 |
	//	+----------------------------+-----------------------------------------------------------------+
	//	+----------------------------+-----------------------------------------------------------------+
	//	| <machine_name>\<user_name> | For a local user with machine_name as the local machine's name. |
	//	+----------------------------+-----------------------------------------------------------------+
	//	| <domain_name>\<user_name>  | For a nonlocal user.                                            |
	//	+----------------------------+-----------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)

	// The FAX_DeleteAccount (Opnum 94) method is called by the client to delete a fax user
	// account previously created with FAX_CreateAccount (section 3.1.4.1.12).
	//
	// Protocol version FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000),
	// and FAX_API_VERSION_2 (0x00020000) fax servers SHOULD NOT implement this call. The
	// fax client MUST NOT call this method if the protocol version reported by the server
	// is FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000), or FAX_API_VERSION_2
	// (0x00020000). For more information, see FAX_ConnectFaxServer (section 3.1.4.1.10).
	//
	// On success, the server MUST delete the specified fax account. The server MUST validate
	// that the client's fax user account has access to delete fax user accounts on the
	// fax server. The client can delete any fax user accounts or the current fax user account.
	// Any subsequent operations on the deleted fax user accounts MUST be failed by the
	// server with the error code ERROR_ACCESS_DENIED. The fax server SHOULD allow deleting
	// a fax user account even if the underlying operating system's user account has been
	// deleted after this fax user account was created.
	//
	// Return Values: This method MUST return 0 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, a fax-specific error defined in
	// section 2.2.52, or one of the standard errors that are defined in [MS-ERREF] section
	// 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|            RETURN ERROR            |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the                |
	//	|                                    | FAX_ACCESS_MANAGE_CONFIG access right.                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. The account name pointed to by the                   |
	//	|                                    | lpcwstrAccountName parameter is NULL or improperly formatted.                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// The account name that lpcwstrAccountName indicates MUST be in one of the following
	// formats. Any other format is invalid.
	//
	//	+----------------------------+-----------------------------------------------------------------+
	//	|                            |                                                                 |
	//	|           FORMAT           |                           DESCRIPTION                           |
	//	|                            |                                                                 |
	//	+----------------------------+-----------------------------------------------------------------+
	//	+----------------------------+-----------------------------------------------------------------+
	//	| <machine_name>\<user_name> | For a local user with machine_name as the local machine's name. |
	//	+----------------------------+-----------------------------------------------------------------+
	//	| <domain_name>\<user_name>  | For a remote (not local) user.                                  |
	//	+----------------------------+-----------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error)

	// The FAX_EnumAccounts (Opnum 95) method is called by the client to enumerate all the
	// fax accounts on the server.
	//
	// Protocol version FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1  (0x00010000),
	// and FAX_API_VERSION_2 (0x00020000) fax servers SHOULD NOT implement this call. The
	// fax client MUST NOT call this method if the protocol version reported by the server
	// is FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000), or FAX_API_VERSION_2
	// (0x00020000). For more information, see FAX_ConnectFaxServer (section 3.1.4.1.10).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | The client's fax user account does not have the access right                     |
	//	|                                    | FAX_ACCESS_QUERY_CONFIG to perform this operation.                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The fax server failed to allocate the amount of memory needed to process this    |
	//	|                                    | request.                                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned if any of the following  |
	//	|                                    | conditions are met: § The value specified for the level parameter is not         |
	//	|                                    | equal to zero. § The pointer specified by the Buffer parameter is NULL.<83>      |
	//	|                                    | § The value pointed to by the BufferSize parameter is zero.<84> § The pointer    |
	//	|                                    | specified by the lpdwAccounts parameter is NULL.<85>                             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// The account name that lpcwstrAccountName indicates MUST be in one of the following
	// formats. Any other format is invalid.
	//
	//	+----------------------------+-----------------------------------------------------------------+
	//	|                            |                                                                 |
	//	|           FORMAT           |                           DESCRIPTION                           |
	//	|                            |                                                                 |
	//	+----------------------------+-----------------------------------------------------------------+
	//	+----------------------------+-----------------------------------------------------------------+
	//	| <machine_name>\<user_name> | For a local user with machine_name as the local machine's name. |
	//	+----------------------------+-----------------------------------------------------------------+
	//	| <domain_name>\<user_name>  | For a nonlocal user.                                            |
	//	+----------------------------+-----------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EnumAccounts(context.Context, *EnumAccountsRequest) (*EnumAccountsResponse, error)

	// The FAX_GetAccountInfo (Opnum 96) method is called by the client to retrieve information
	// about a specified fax user account. The fax user account for which information is
	// retrieved is specified by the lpcwstrAccountName parameter, which can be obtained
	// using the FAX_EnumAccounts (section 3.1.4.1.18) method.
	//
	// Protocol version FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000),
	// and FAX_API_VERSION_2 (0x00020000) fax servers SHOULD NOT implement this call. The
	// fax client MUST NOT call this method if the protocol version reported by the server
	// is FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000), or FAX_API_VERSION_2
	// (0x00020000). For more information, see FAX_ConnectFaxServer (section 3.1.4.1.10).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_FILE_NOT_FOUND 0x00000002    | The account name specified by the lpcwstrAccountName parameter appears valid but |
	//	|                                    | does not exist.                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The user account specified by the lpwstrAccountName            |
	//	|                                    | argument is not the caller, and it does not have the fax access rights           |
	//	|                                    | FAX_ACCESS_QUERY_CONFIG.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. The value of the level argument is greater than      |
	//	|                                    | zero. The account name specified by lpcwstrAccountName is not a valid fax        |
	//	|                                    | account name.                                                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// The account name that lpcwstrAccountName indicates MUST be in one of the following
	// formats. Any other format is invalid.
	//
	//	+----------------------------+--------------------------------------------------------------------------+
	//	|                            |                                                                          |
	//	|           FORMAT           |                               DESCRIPTION                                |
	//	|                            |                                                                          |
	//	+----------------------------+--------------------------------------------------------------------------+
	//	+----------------------------+--------------------------------------------------------------------------+
	//	| <machine_name>\<user_name> | For a local user that has machine_name as the name of the local machine. |
	//	+----------------------------+--------------------------------------------------------------------------+
	//	| <domain_name>\<user_name>  | For a nonlocal user.                                                     |
	//	+----------------------------+--------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetAccountInfo(context.Context, *GetAccountInfoRequest) (*GetAccountInfoResponse, error)

	// The FAX_GetGeneralConfiguration (Opnum 97) method is called by the client to request
	// information about the general configuration at the server.
	//
	// Protocol version FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000),
	// and FAX_API_VERSION_2 (0x00020000) fax servers SHOULD NOT implement this call. The
	// fax client MUST NOT call this method if the protocol version reported by the server
	// is FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000), or FAX_API_VERSION_2
	// (0x00020000). For more information, see FAX_ConnectFaxServer (section 3.1.4.1.10).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the                |
	//	|                                    | FAX_ACCESS_QUERY_CONFIG access rights required for this operation.               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | Not enough storage is available to process this command. The fax server failed   |
	//	|                                    | to allocate sufficient memory to hold the FAX_GENERAL_CONFIG to be returned to   |
	//	|                                    | the client.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. The level parameter is set to a value other than 0.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F    | The fax server failed to custom marshal the FAX_GENERAL_CONFIG to be returned to |
	//	|                                    | the client.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetGeneralConfiguration(context.Context, *GetGeneralConfigurationRequest) (*GetGeneralConfigurationResponse, error)

	// The fax client application calls the FAX_SetGeneralConfiguration (Opnum 98) method
	// to set the configuration options provided for the fax service. The FAX_GENERAL_CONFIG
	// (section 2.2.31) structure MUST be serialized. The variable data fields, such as
	// strings, MUST be filled with the offset to the string from the beginning of the buffer
	// and not the actual address. In response, the server MUST validate that the client's
	// fax user account has access to manage configuration on the server. On success, the
	// server MUST set the requested configuration options.
	//
	// Protocol version FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1  (0x00010000),
	// and FAX_API_VERSION_2 (0x00020000) fax servers SHOULD NOT implement this call. The
	// fax client MUST NOT call this method if the protocol version reported by the server
	// is FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000), or FAX_API_VERSION_2
	// (0x00020000). For more information, see FAX_ConnectFaxServer (section 3.1.4.1.10).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the required       |
	//	|                                    | FAX_ACCESS_MANAGE_CONFIG authorization for this operation.                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | In order to process the data for custom marshaling, the server needs to make     |
	//	|                                    | a copy of the FAX_GENERAL_CONFIG provided by the client; but the server cannot   |
	//	|                                    | allocate sufficient memory to hold the copy of the FAX_GENERAL_CONFIG.           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | The lpcwstrArchiveLocationOffset member of the Fixed_Portion of the              |
	//	|                                    | FAX_GENERAL_CONFIG referenced by the Buffer parameter is set to an invalid       |
	//	|                                    | offset value.                                                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned if any of the following  |
	//	|                                    | conditions are met: § The level parameter is set to a value greater than 0. §    |
	//	|                                    | In the FAX_GENERAL_CONFIG referenced by Buffer, the dtDiscountStart.Hour member  |
	//	|                                    | is set to a value greater than or equal to 24, and the dtDiscountStart.Minute    |
	//	|                                    | member is set to a value greater than or equal to 60. § In the                   |
	//	|                                    | FAX_GENERAL_CONFIG referenced by Buffer, the dtDiscountEnd.Hour member is set    |
	//	|                                    | to a value greater than or equal to 24, and the dtDiscountEnd.Minute member      |
	//	|                                    | is set to a value greater than or equal to 60. § In the FAX_GENERAL_CONFIG       |
	//	|                                    | referenced by Buffer, the bUseArchive member is set to TRUE and the              |
	//	|                                    | lpcwstrArchiveLocationOffset member is set to 0. § In the FAX_GENERAL_CONFIG     |
	//	|                                    | referenced by Buffer, the dwSizeQuotaHighWaterMark member is set to a value      |
	//	|                                    | greater than the value of the dwSizeQuotaLowWaterMark member. § In the           |
	//	|                                    | FAX_GENERAL_CONFIG referenced by Buffer, the dwQueueState member contains one    |
	//	|                                    | or more of the following flag values: FAX_INCOMING_BLOCKED, FAX_OUTBOX_BLOCKED,  |
	//	|                                    | or FAX_OUTBOX_PAUSED. § The value of the BufferSize parameter is less than       |
	//	|                                    | the size required to hold the custom marshaled FAX_GENERAL_CONFIG structure      |
	//	|                                    | referenced by Buffer. This size is specified by the dwSizeOfStruct member of the |
	//	|                                    | FAX_GENERAL_CONFIG referenced by Buffer.                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INTERNAL_ERROR 0x0000054F    | The server failed to parse the custom marshaled FAX_GENERAL_CONFIG.              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetGeneralConfiguration(context.Context, *SetGeneralConfigurationRequest) (*SetGeneralConfigurationResponse, error)

	// The FAX_GetSecurityEx2 (Opnum 99) method is called by the client to retrieve information
	// about the fax security descriptor from the fax server.
	//
	// Protocol version FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000),
	// and FAX_API_VERSION_2 (0x00020000) fax servers SHOULD NOT implement this call. The
	// fax client MUST NOT call this method if the protocol version reported by the server
	// is FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000), or FAX_API_VERSION_2
	// (0x00020000). For more information, see FAX_ConnectFaxServer (section 3.1.4.1.10).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                                  |
	//	|               VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005          | Access is denied. This error is returned when there is a mismatch between        |
	//	|                                         | the access level requested (Read control, access to set/get SACL security        |
	//	|                                         | information or both) through the bit pattern in SecurityInformation and the      |
	//	|                                         | current authorized level.                                                        |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008      | Not enough storage is available to process this command.                         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057      | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                         | following conditions: § The pSecurityDescriptor parameter is NULL. § The value   |
	//	|                                         | of the SecurityInformation parameter does not conform to the definition of valid |
	//	|                                         | bit patterns for this parameter.                                                 |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_SECURITY_DESCR 0x0000053A | The security descriptor structure is invalid.                                    |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetSecurityEx2(context.Context, *GetSecurityEx2Request) (*GetSecurityEx2Response, error)

	// The FAX_SetSecurityEx2 (Opnum 100) method is called by the client. On success, the
	// server MUST set the fax server's security descriptor.
	//
	// Protocol version FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1  (0x00010000),
	// and FAX_API_VERSION_2 (0x00020000) fax servers SHOULD NOT implement this call. The
	// fax client MUST NOT call this method if the protocol version reported by the server
	// is FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000), or FAX_API_VERSION_2
	// (0x00020000). For more information, see FAX_ConnectFaxServer (section 3.1.4.1.10).
	// The fax client SHOULD call FAX_SetSecurity (section 3.1.4.1.94) instead.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the access         |
	//	|                                    | rights required for this operation: § WRITE_OWNER, when the SecurityInformation  |
	//	|                                    | parameter contains the OWNER_SECURITY_INFORMATION value. § WRITE_DAC,            |
	//	|                                    | when SecurityInformation contains the GROUP_SECURITY_INFORMATION                 |
	//	|                                    | or DACL_SECURITY_INFORMATION values. § ACCESS_SYSTEM_SECURITY, when              |
	//	|                                    | SecurityInformation contains the SACL_SECURITY_INFORMATION value.                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | The data is invalid. The data contained in the buffer specified by the           |
	//	|                                    | pSecurityDescriptor parameter is not a valid SECURITY_DESCRIPTOR.                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_OUTOFMEMORY 0x0000000E       | The fax server cannot allocate sufficient memory for a properly constructed      |
	//	|                                    | FAX_EVENT_EX_1 (section 2.2.68) structure describing a FAX_EVENT_TYPE_CONFIG     |
	//	|                                    | event to be signaled to the client. A properly constructed structure has the     |
	//	|                                    | ConfigType member of the FAX_EVENT_EX_1 set to FAX_CONFIG_TYPE_SECURITY. For     |
	//	|                                    | more details, see FAX_ClientEventQueueEx (section 3.2.4.3).                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned under any                |
	//	|                                    | of the following conditions: § pSecurityDescriptor is set to a NULL              |
	//	|                                    | pointer value. § The dwBufferSize parameter is set to 0x00000000. §              |
	//	|                                    | SecurityInformation is set to a value that does not contain any of the           |
	//	|                                    | following flags: OWNER_SECURITY_INFORMATION, GROUP_SECURITY_INFORMATION,         |
	//	|                                    | DACL_SECURITY_INFORMATION, or SACL_SECURITY_INFORMATION.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_REGISTRY_CORRUPT 0x000003F7  | The registry is corrupted. The structure of one of the files that contains       |
	//	|                                    | registry data is corrupted, or the system's memory image of the file is          |
	//	|                                    | corrupted, or the file could not be recovered because the alternate copy or log  |
	//	|                                    | was absent or corrupted.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	SetSecurityEx2(context.Context, *SetSecurityEx2Request) (*SetSecurityEx2Response, error)

	// The FAX_AccessCheckEx2 (Opnum 101) method is called by the client when the client
	// needs to check whether the client's fax user account has certain access permissions
	// on the server.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The caller does not have the required permissions for this     |
	//	|                                    | request (the caller does not have a valid fax user account).                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                    | following conditions: § The pointer specified in the pfAccess parameter is NULL. |
	//	|                                    | § The fax access rights specified in the lpdwRights parameter contain invalid    |
	//	|                                    | access values.                                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	AccessCheckEx2(context.Context, *AccessCheckEx2Request) (*AccessCheckEx2Response, error)

	// The fax client application calls the FAX_ReAssignMessage (Opnum 102) method to reassign
	// the specified fax message to a set of users.<144>
	//
	// Protocol version FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000),
	// and FAX_API_VERSION_2 (0x00020000) fax servers SHOULD NOT implement this call. The
	// fax client MUST NOT call this method if the protocol version reported by the server
	// is FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000), or FAX_API_VERSION_2
	// (0x00020000). For more information, see FAX_ConnectFaxServer (section 3.1.4.1.10).
	//
	// The dwlMessageId parameter specifies a particular message and can be obtained using
	// the FAX_EnumMessages (section 3.1.4.1.24) method or the FAX_EnumMessagesEx (section
	// 3.1.4.1.25) methods.
	//
	// The client MUST specify the recipients for a reassigned message in a semicolon (;)
	// separated format. In response, the server MUST validate whether the bIncomingFaxesArePublic
	// option in the FAX_GENERAL_CONFIG (section 2.2.31) data structure is not set in the
	// server configuration. The server MUST also validate whether the message that is specified
	// by the dwlMessageId argument refers to a valid message on the server. The server
	// MUST validate that there are recipient numbers for each of the recipients that are
	// listed in pReAssignInfo structure. On success, the server MUST reassign the specified
	// fax message.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_FILE_NOT_FOUND 0x00000002    | This error code is returned under any of the following conditions: § The         |
	//	|                                    | dwlMessageId argument does not specify a valid message. § One or more recipients |
	//	|                                    | specified in the lpcwstrRecipients field of the pReAssignInfo argument do not    |
	//	|                                    | have a corresponding fax user account.                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | The caller does not have the FAX_ACCESS_MANAGE_RECEIVE_FOLDER access rights.     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | This error code is returned under any of the following conditions: § The value   |
	//	|                                    | specified for the dwlMessageId parameter is zero. § The lpcwstrRecipients member |
	//	|                                    | of the data structure specified by the pReAssignInfo parameter is set to NULL or |
	//	|                                    | to an empty string.                                                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BUFFER_OVERFLOW 0x0000006F   | The number of recipients specified in the lpcwstrRecipients member of the        |
	//	|                                    | data structure pointed at by pReAssignInfo is greater than FAX_MAX_RECIPIENTS    |
	//	|                                    | (10000).                                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_OPERATION 0x000010DD | This error code is returned under any of the following conditions: § Incoming    |
	//	|                                    | faxes are public (section 2.2.31). Reassignment is not supported when incoming   |
	//	|                                    | faxes are public. § The server does not support reassignment, or the server      |
	//	|                                    | is configured with a policy that is currently set to disable fax message         |
	//	|                                    | reassignment.                                                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	ReassignMessage(context.Context, *ReassignMessageRequest) (*ReassignMessageResponse, error)

	// The fax client application calls the FAX_SetMessage (Opnum 103) method to set the
	// specific message properties for the message identified by its ID.<162> The dwlMessageId
	// parameter specifies a particular message and can be obtained using the FAX_EnumMessages
	// (section 3.1.4.1.24) method or the FAX_EnumMessagesEx (section 3.1.4.1.25) method.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005       | Access is denied. The client's fax user account does not have the                |
	//	|                                      | ALL_FAX_USER_ACCESS_RIGHTS access rights required for this operation.            |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057   | The parameter is incorrect. This error code is returned if any of the            |
	//	|                                      | following conditions are met: § The structure pointed to by the lpMessageProps   |
	//	|                                      | argument contains invalid data. § The Folder argument has an invalid value (a    |
	//	|                                      | value other than FAX_MESSAGE_FOLDER_INBOX or FAX_MESSAGE_FOLDER_SENTITEMS).      |
	//	|                                      | § The dwlMessageId parameter is zero. § In the structure pointed at              |
	//	|                                      | by the lpMessageProps argument, the dwValidityMask field contains the            |
	//	|                                      | FAX_MSG_PROP_FIELD_MSG_FLAGS, and the dwMsgFlags field does not contain the      |
	//	|                                      | FAX_MSG_ALL_FLAGS flag.                                                          |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ERR_MESSAGE_NOT_FOUND 0x00001B61 | The fax message specified by the dwlMessageId argument cannot be found by the    |
	//	|                                      | fax server in the folder specified by the Folder argument.                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetMessage(context.Context, *SetMessageRequest) (*SetMessageResponse, error)

	// The FAX_GetConfigOption (Opnum 104) is called by the client to retrieve a configuration
	// setting at the server using an RPC_REQUEST packet.
	//
	// Protocol version FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000),
	// and FAX_API_VERSION_2 (0x00020000) fax servers SHOULD NOT implement this call. The
	// fax client MUST NOT call this method if the protocol version reported by the server
	// is FAX_API_VERSION_0 (0x00000000), FAX_API_VERSION_1 (0x00010000), or FAX_API_VERSION_2
	// (0x00020000). For more information, see FAX_ConnectFaxServer (section 3.1.4.1.10).
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have any of the         |
	//	|                                    | permissions covered by ALL_FAX_USER_ACCESS_RIGHTS.                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | This error is returned when any of the following conditions are met: § The       |
	//	|                                    | lpdwValue parameter is set to a NULL pointer value.<103> § The configuration     |
	//	|                                    | option specified by the option parameter is not one of the following             |
	//	|                                    | values: FAX_CONFIG_OPTION_ALLOW_PERSONAL_CP, FAX_CONFIG_OPTION_QUEUE_STATE,      |
	//	|                                    | FAX_CONFIG_OPTION_ALLOW_RECEIPTS or FAX_CONFIG_OPTION_INCOMING_FAXES_PUBLIC.     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Use of this method does not require FAX_ACCESS_QUERY_CONFIG access rights. A calling
	// user with any ACE on the server can use this method.
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetConfigOption(context.Context, *GetConfigOptionRequest) (*GetConfigOptionResponse, error)
}

func RegisterFaxServer(conn dcerpc.Conn, o FaxServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewFaxServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(FaxSyntaxV4_0))...)
}

func NewFaxServerHandle(o FaxServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return FaxServerHandle(ctx, o, opNum, r)
	}
}

func FaxServerHandle(ctx context.Context, o FaxServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // FAX_GetServicePrinters
		op := &xxx_GetServicePrintersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServicePrintersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServicePrinters(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // FAX_ConnectionRefCount
		op := &xxx_ConnectionReferenceCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectionReferenceCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConnectionReferenceCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // FAX_OpenPort
		op := &xxx_OpenPortOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenPortRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenPort(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // FAX_ClosePort
		op := &xxx_ClosePortOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ClosePortRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ClosePort(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // FAX_EnumJobs
		op := &xxx_EnumJobsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumJobsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumJobs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // FAX_GetJob
		op := &xxx_GetJobOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetJobRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetJob(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // FAX_SetJob
		op := &xxx_SetJobOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetJobRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetJob(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // FAX_GetPageData
		op := &xxx_GetPageDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPageDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPageData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // FAX_GetDeviceStatus
		op := &xxx_GetDeviceStatusOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDeviceStatusRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDeviceStatus(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // FAX_Abort
		op := &xxx_AbortOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AbortRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Abort(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // FAX_EnumPorts
		op := &xxx_EnumPortsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumPortsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumPorts(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // FAX_GetPort
		op := &xxx_GetPortOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPortRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPort(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // FAX_SetPort
		op := &xxx_SetPortOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPortRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPort(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // FAX_EnumRoutingMethods
		op := &xxx_EnumRoutingMethodsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumRoutingMethodsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumRoutingMethods(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // FAX_EnableRoutingMethod
		op := &xxx_EnableRoutingMethodOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnableRoutingMethodRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnableRoutingMethod(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // FAX_GetRoutingInfo
		op := &xxx_GetRoutingInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRoutingInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRoutingInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // FAX_SetRoutingInfo
		op := &xxx_SetRoutingInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetRoutingInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetRoutingInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // FAX_EnumGlobalRoutingInfo
		op := &xxx_EnumGlobalRoutingInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumGlobalRoutingInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumGlobalRoutingInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // FAX_SetGlobalRoutingInfo
		op := &xxx_SetGlobalRoutingInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetGlobalRoutingInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetGlobalRoutingInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // FAX_GetConfiguration
		op := &xxx_GetConfigurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetConfigurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetConfiguration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // FAX_SetConfiguration
		op := &xxx_SetConfigurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetConfigurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetConfiguration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // FAX_GetLoggingCategories
		op := &xxx_GetLoggingCategoriesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLoggingCategoriesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLoggingCategories(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // FAX_SetLoggingCategories
		op := &xxx_SetLoggingCategoriesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetLoggingCategoriesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetLoggingCategories(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // FAX_GetSecurity
		op := &xxx_GetSecurityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSecurityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSecurity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // FAX_SetSecurity
		op := &xxx_SetSecurityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSecurityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSecurity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // FAX_AccessCheck
		op := &xxx_AccessCheckOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AccessCheckRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AccessCheck(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // FAX_CheckServerProtSeq
		op := &xxx_CheckServerProtocolSeqOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CheckServerProtocolSeqRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CheckServerProtocolSeq(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // FAX_SendDocumentEx
		op := &xxx_SendDocumentExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SendDocumentExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SendDocumentEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // FAX_EnumJobsEx
		op := &xxx_EnumJobsExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumJobsExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumJobsEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // FAX_GetJobEx
		op := &xxx_GetJobExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetJobExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetJobEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // FAX_GetCountryList
		op := &xxx_GetCountryListOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetCountryListRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetCountryList(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 31: // FAX_GetPersonalProfileInfo
		op := &xxx_GetPersonalProfileInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPersonalProfileInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPersonalProfileInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // FAX_GetQueueStates
		op := &xxx_GetQueueStatesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetQueueStatesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetQueueStates(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // FAX_SetQueue
		op := &xxx_SetQueueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetQueueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetQueue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // FAX_GetReceiptsConfiguration
		op := &xxx_GetReceiptsConfigurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetReceiptsConfigurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetReceiptsConfiguration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // FAX_SetReceiptsConfiguration
		op := &xxx_SetReceiptsConfigurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetReceiptsConfigurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetReceiptsConfiguration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 36: // FAX_GetReceiptsOptions
		op := &xxx_GetReceiptsOptionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetReceiptsOptionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetReceiptsOptions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 37: // FAX_GetVersion
		op := &xxx_GetVersionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetVersionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetVersion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 38: // FAX_GetOutboxConfiguration
		op := &xxx_GetOutboxConfigurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetOutboxConfigurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetOutboxConfiguration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 39: // FAX_SetOutboxConfiguration
		op := &xxx_SetOutboxConfigurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetOutboxConfigurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetOutboxConfiguration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 40: // FAX_GetPersonalCoverPagesOption
		op := &xxx_GetPersonalCoverPagesOptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPersonalCoverPagesOptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPersonalCoverPagesOption(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 41: // FAX_GetArchiveConfiguration
		op := &xxx_GetArchiveConfigurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetArchiveConfigurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetArchiveConfiguration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 42: // FAX_SetArchiveConfiguration
		op := &xxx_SetArchiveConfigurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetArchiveConfigurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetArchiveConfiguration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 43: // FAX_GetActivityLoggingConfiguration
		op := &xxx_GetActivityLoggingConfigurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetActivityLoggingConfigurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetActivityLoggingConfiguration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 44: // FAX_SetActivityLoggingConfiguration
		op := &xxx_SetActivityLoggingConfigurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetActivityLoggingConfigurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetActivityLoggingConfiguration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 45: // FAX_EnumerateProviders
		op := &xxx_EnumerateProvidersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumerateProvidersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumerateProviders(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 46: // FAX_GetPortEx
		op := &xxx_GetPortExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPortExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPortEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 47: // FAX_SetPortEx
		op := &xxx_SetPortExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPortExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPortEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 48: // FAX_EnumPortsEx
		op := &xxx_EnumPortsExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumPortsExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumPortsEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 49: // FAX_GetExtensionData
		op := &xxx_GetExtensionDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetExtensionDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetExtensionData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 50: // FAX_SetExtensionData
		op := &xxx_SetExtensionDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetExtensionDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetExtensionData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 51: // FAX_AddOutboundGroup
		op := &xxx_AddOutboundGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddOutboundGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddOutboundGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 52: // FAX_SetOutboundGroup
		op := &xxx_SetOutboundGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetOutboundGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetOutboundGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 53: // FAX_RemoveOutboundGroup
		op := &xxx_RemoveOutboundGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveOutboundGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveOutboundGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 54: // FAX_EnumOutboundGroups
		op := &xxx_EnumOutboundGroupsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumOutboundGroupsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumOutboundGroups(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 55: // FAX_SetDeviceOrderInGroup
		op := &xxx_SetDeviceOrderInGroupOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDeviceOrderInGroupRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDeviceOrderInGroup(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 56: // FAX_AddOutboundRule
		op := &xxx_AddOutboundRuleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddOutboundRuleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddOutboundRule(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 57: // FAX_RemoveOutboundRule
		op := &xxx_RemoveOutboundRuleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveOutboundRuleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveOutboundRule(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 58: // FAX_SetOutboundRule
		op := &xxx_SetOutboundRuleOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetOutboundRuleRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetOutboundRule(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 59: // FAX_EnumOutboundRules
		op := &xxx_EnumOutboundRulesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumOutboundRulesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumOutboundRules(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 60: // FAX_RegisterServiceProviderEx
		op := &xxx_RegisterServiceProviderExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RegisterServiceProviderExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RegisterServiceProviderEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 61: // FAX_UnregisterServiceProviderEx
		op := &xxx_UnregisterServiceProviderExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UnregisterServiceProviderExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UnregisterServiceProviderEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 62: // FAX_UnregisterRoutingExtension
		op := &xxx_UnregisterRoutingExtensionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &UnregisterRoutingExtensionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.UnregisterRoutingExtension(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 63: // FAX_StartMessagesEnum
		op := &xxx_StartMessagesEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StartMessagesEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StartMessagesEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 64: // FAX_EndMessagesEnum
		op := &xxx_EndMessagesEnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EndMessagesEnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EndMessagesEnum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 65: // FAX_EnumMessages
		op := &xxx_EnumMessagesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumMessagesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumMessages(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 66: // FAX_GetMessage
		op := &xxx_GetMessageOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMessageRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMessage(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 67: // FAX_RemoveMessage
		op := &xxx_RemoveMessageOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveMessageRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveMessage(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 68: // FAX_StartCopyToServer
		op := &xxx_StartCopyToServerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StartCopyToServerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StartCopyToServer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 69: // FAX_StartCopyMessageFromServer
		op := &xxx_StartCopyMessageFromServerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StartCopyMessageFromServerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StartCopyMessageFromServer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 70: // FAX_WriteFile
		op := &xxx_WriteFileOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &WriteFileRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.WriteFile(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 71: // FAX_ReadFile
		op := &xxx_ReadFileOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReadFileRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReadFile(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 72: // FAX_EndCopy
		op := &xxx_EndCopyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EndCopyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EndCopy(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 73: // FAX_StartServerNotification
		op := &xxx_StartServerNotificationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StartServerNotificationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StartServerNotification(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 74: // FAX_StartServerNotificationEx
		op := &xxx_StartServerNotificationExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StartServerNotificationExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StartServerNotificationEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 75: // FAX_EndServerNotification
		op := &xxx_EndServerNotificationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EndServerNotificationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EndServerNotification(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 76: // FAX_GetServerActivity
		op := &xxx_GetServerActivityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServerActivityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServerActivity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 77: // FAX_SetConfigWizardUsed
		op := &xxx_SetConfigWizardUsedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetConfigWizardUsedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetConfigWizardUsed(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 78: // FAX_EnumRoutingExtensions
		op := &xxx_EnumRoutingExtensionsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumRoutingExtensionsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumRoutingExtensions(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 79: // FAX_ConnectFaxServer
		op := &xxx_ConnectFaxServerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectFaxServerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConnectFaxServer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 80: // FAX_GetSecurityEx
		op := &xxx_GetSecurityExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSecurityExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSecurityEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 81: // FAX_RefreshArchive
		op := &xxx_RefreshArchiveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RefreshArchiveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RefreshArchive(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 82: // FAX_SetRecipientsLimit
		op := &xxx_SetRecipientsLimitOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetRecipientsLimitRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetRecipientsLimit(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 83: // FAX_GetRecipientsLimit
		op := &xxx_GetRecipientsLimitOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRecipientsLimitRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRecipientsLimit(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 84: // FAX_GetServerSKU
		op := &xxx_GetServerSKUOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetServerSKURequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetServerSKU(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 85: // FAX_CheckValidFaxFolder
		op := &xxx_CheckValidFaxFolderOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CheckValidFaxFolderRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CheckValidFaxFolder(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 86: // FAX_GetJobEx2
		op := &xxx_GetJobEx2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetJobEx2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetJobEx2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 87: // FAX_EnumJobsEx2
		op := &xxx_EnumJobsEx2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumJobsEx2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumJobsEx2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 88: // FAX_GetMessageEx
		op := &xxx_GetMessageExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMessageExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMessageEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 89: // FAX_StartMessagesEnumEx
		op := &xxx_StartMessagesEnumExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StartMessagesEnumExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StartMessagesEnumEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 90: // FAX_EnumMessagesEx
		op := &xxx_EnumMessagesExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumMessagesExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumMessagesEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 91: // FAX_StartServerNotificationEx2
		op := &xxx_StartServerNotificationEx2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StartServerNotificationEx2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StartServerNotificationEx2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 92: // FAX_CreateAccount
		op := &xxx_CreateAccountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &CreateAccountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.CreateAccount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 93: // FAX_DeleteAccount
		op := &xxx_DeleteAccountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &DeleteAccountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.DeleteAccount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 94: // FAX_EnumAccounts
		op := &xxx_EnumAccountsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumAccountsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumAccounts(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 95: // FAX_GetAccountInfo
		op := &xxx_GetAccountInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetAccountInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetAccountInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 96: // FAX_GetGeneralConfiguration
		op := &xxx_GetGeneralConfigurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetGeneralConfigurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetGeneralConfiguration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 97: // FAX_SetGeneralConfiguration
		op := &xxx_SetGeneralConfigurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetGeneralConfigurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetGeneralConfiguration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 98: // FAX_GetSecurityEx2
		op := &xxx_GetSecurityEx2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSecurityEx2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSecurityEx2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 99: // FAX_SetSecurityEx2
		op := &xxx_SetSecurityEx2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSecurityEx2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSecurityEx2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 100: // FAX_AccessCheckEx2
		op := &xxx_AccessCheckEx2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AccessCheckEx2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AccessCheckEx2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 101: // FAX_ReAssignMessage
		op := &xxx_ReassignMessageOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ReassignMessageRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ReassignMessage(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 102: // FAX_SetMessage
		op := &xxx_SetMessageOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetMessageRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetMessage(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 103: // FAX_GetConfigOption
		op := &xxx_GetConfigOptionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetConfigOptionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetConfigOption(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented fax
type UnimplementedFaxServer struct {
}

func (UnimplementedFaxServer) GetServicePrinters(context.Context, *GetServicePrintersRequest) (*GetServicePrintersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) ConnectionReferenceCount(context.Context, *ConnectionReferenceCountRequest) (*ConnectionReferenceCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) OpenPort(context.Context, *OpenPortRequest) (*OpenPortResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) ClosePort(context.Context, *ClosePortRequest) (*ClosePortResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) EnumJobs(context.Context, *EnumJobsRequest) (*EnumJobsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetJob(context.Context, *GetJobRequest) (*GetJobResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetJob(context.Context, *SetJobRequest) (*SetJobResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetPageData(context.Context, *GetPageDataRequest) (*GetPageDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetDeviceStatus(context.Context, *GetDeviceStatusRequest) (*GetDeviceStatusResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) Abort(context.Context, *AbortRequest) (*AbortResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) EnumPorts(context.Context, *EnumPortsRequest) (*EnumPortsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetPort(context.Context, *GetPortRequest) (*GetPortResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetPort(context.Context, *SetPortRequest) (*SetPortResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) EnumRoutingMethods(context.Context, *EnumRoutingMethodsRequest) (*EnumRoutingMethodsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) EnableRoutingMethod(context.Context, *EnableRoutingMethodRequest) (*EnableRoutingMethodResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetRoutingInfo(context.Context, *GetRoutingInfoRequest) (*GetRoutingInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetRoutingInfo(context.Context, *SetRoutingInfoRequest) (*SetRoutingInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) EnumGlobalRoutingInfo(context.Context, *EnumGlobalRoutingInfoRequest) (*EnumGlobalRoutingInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetGlobalRoutingInfo(context.Context, *SetGlobalRoutingInfoRequest) (*SetGlobalRoutingInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetConfiguration(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetConfiguration(context.Context, *SetConfigurationRequest) (*SetConfigurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetLoggingCategories(context.Context, *GetLoggingCategoriesRequest) (*GetLoggingCategoriesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetLoggingCategories(context.Context, *SetLoggingCategoriesRequest) (*SetLoggingCategoriesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetSecurity(context.Context, *GetSecurityRequest) (*GetSecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetSecurity(context.Context, *SetSecurityRequest) (*SetSecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) AccessCheck(context.Context, *AccessCheckRequest) (*AccessCheckResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) CheckServerProtocolSeq(context.Context, *CheckServerProtocolSeqRequest) (*CheckServerProtocolSeqResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SendDocumentEx(context.Context, *SendDocumentExRequest) (*SendDocumentExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) EnumJobsEx(context.Context, *EnumJobsExRequest) (*EnumJobsExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetJobEx(context.Context, *GetJobExRequest) (*GetJobExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetCountryList(context.Context, *GetCountryListRequest) (*GetCountryListResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetPersonalProfileInfo(context.Context, *GetPersonalProfileInfoRequest) (*GetPersonalProfileInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetQueueStates(context.Context, *GetQueueStatesRequest) (*GetQueueStatesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetQueue(context.Context, *SetQueueRequest) (*SetQueueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetReceiptsConfiguration(context.Context, *GetReceiptsConfigurationRequest) (*GetReceiptsConfigurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetReceiptsConfiguration(context.Context, *SetReceiptsConfigurationRequest) (*SetReceiptsConfigurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetReceiptsOptions(context.Context, *GetReceiptsOptionsRequest) (*GetReceiptsOptionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetOutboxConfiguration(context.Context, *GetOutboxConfigurationRequest) (*GetOutboxConfigurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetOutboxConfiguration(context.Context, *SetOutboxConfigurationRequest) (*SetOutboxConfigurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetPersonalCoverPagesOption(context.Context, *GetPersonalCoverPagesOptionRequest) (*GetPersonalCoverPagesOptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetArchiveConfiguration(context.Context, *GetArchiveConfigurationRequest) (*GetArchiveConfigurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetArchiveConfiguration(context.Context, *SetArchiveConfigurationRequest) (*SetArchiveConfigurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetActivityLoggingConfiguration(context.Context, *GetActivityLoggingConfigurationRequest) (*GetActivityLoggingConfigurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetActivityLoggingConfiguration(context.Context, *SetActivityLoggingConfigurationRequest) (*SetActivityLoggingConfigurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) EnumerateProviders(context.Context, *EnumerateProvidersRequest) (*EnumerateProvidersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetPortEx(context.Context, *GetPortExRequest) (*GetPortExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetPortEx(context.Context, *SetPortExRequest) (*SetPortExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) EnumPortsEx(context.Context, *EnumPortsExRequest) (*EnumPortsExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetExtensionData(context.Context, *GetExtensionDataRequest) (*GetExtensionDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetExtensionData(context.Context, *SetExtensionDataRequest) (*SetExtensionDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) AddOutboundGroup(context.Context, *AddOutboundGroupRequest) (*AddOutboundGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetOutboundGroup(context.Context, *SetOutboundGroupRequest) (*SetOutboundGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) RemoveOutboundGroup(context.Context, *RemoveOutboundGroupRequest) (*RemoveOutboundGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) EnumOutboundGroups(context.Context, *EnumOutboundGroupsRequest) (*EnumOutboundGroupsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetDeviceOrderInGroup(context.Context, *SetDeviceOrderInGroupRequest) (*SetDeviceOrderInGroupResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) AddOutboundRule(context.Context, *AddOutboundRuleRequest) (*AddOutboundRuleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) RemoveOutboundRule(context.Context, *RemoveOutboundRuleRequest) (*RemoveOutboundRuleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetOutboundRule(context.Context, *SetOutboundRuleRequest) (*SetOutboundRuleResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) EnumOutboundRules(context.Context, *EnumOutboundRulesRequest) (*EnumOutboundRulesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) RegisterServiceProviderEx(context.Context, *RegisterServiceProviderExRequest) (*RegisterServiceProviderExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) UnregisterServiceProviderEx(context.Context, *UnregisterServiceProviderExRequest) (*UnregisterServiceProviderExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) UnregisterRoutingExtension(context.Context, *UnregisterRoutingExtensionRequest) (*UnregisterRoutingExtensionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) StartMessagesEnum(context.Context, *StartMessagesEnumRequest) (*StartMessagesEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) EndMessagesEnum(context.Context, *EndMessagesEnumRequest) (*EndMessagesEnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) EnumMessages(context.Context, *EnumMessagesRequest) (*EnumMessagesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) RemoveMessage(context.Context, *RemoveMessageRequest) (*RemoveMessageResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) StartCopyToServer(context.Context, *StartCopyToServerRequest) (*StartCopyToServerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) StartCopyMessageFromServer(context.Context, *StartCopyMessageFromServerRequest) (*StartCopyMessageFromServerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) WriteFile(context.Context, *WriteFileRequest) (*WriteFileResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) ReadFile(context.Context, *ReadFileRequest) (*ReadFileResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) EndCopy(context.Context, *EndCopyRequest) (*EndCopyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) StartServerNotification(context.Context, *StartServerNotificationRequest) (*StartServerNotificationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) StartServerNotificationEx(context.Context, *StartServerNotificationExRequest) (*StartServerNotificationExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) EndServerNotification(context.Context, *EndServerNotificationRequest) (*EndServerNotificationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetServerActivity(context.Context, *GetServerActivityRequest) (*GetServerActivityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetConfigWizardUsed(context.Context, *SetConfigWizardUsedRequest) (*SetConfigWizardUsedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) EnumRoutingExtensions(context.Context, *EnumRoutingExtensionsRequest) (*EnumRoutingExtensionsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) ConnectFaxServer(context.Context, *ConnectFaxServerRequest) (*ConnectFaxServerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetSecurityEx(context.Context, *GetSecurityExRequest) (*GetSecurityExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) RefreshArchive(context.Context, *RefreshArchiveRequest) (*RefreshArchiveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetRecipientsLimit(context.Context, *SetRecipientsLimitRequest) (*SetRecipientsLimitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetRecipientsLimit(context.Context, *GetRecipientsLimitRequest) (*GetRecipientsLimitResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetServerSKU(context.Context, *GetServerSKURequest) (*GetServerSKUResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) CheckValidFaxFolder(context.Context, *CheckValidFaxFolderRequest) (*CheckValidFaxFolderResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetJobEx2(context.Context, *GetJobEx2Request) (*GetJobEx2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) EnumJobsEx2(context.Context, *EnumJobsEx2Request) (*EnumJobsEx2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetMessageEx(context.Context, *GetMessageExRequest) (*GetMessageExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) StartMessagesEnumEx(context.Context, *StartMessagesEnumExRequest) (*StartMessagesEnumExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) EnumMessagesEx(context.Context, *EnumMessagesExRequest) (*EnumMessagesExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) StartServerNotificationEx2(context.Context, *StartServerNotificationEx2Request) (*StartServerNotificationEx2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) EnumAccounts(context.Context, *EnumAccountsRequest) (*EnumAccountsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetAccountInfo(context.Context, *GetAccountInfoRequest) (*GetAccountInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetGeneralConfiguration(context.Context, *GetGeneralConfigurationRequest) (*GetGeneralConfigurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetGeneralConfiguration(context.Context, *SetGeneralConfigurationRequest) (*SetGeneralConfigurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetSecurityEx2(context.Context, *GetSecurityEx2Request) (*GetSecurityEx2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetSecurityEx2(context.Context, *SetSecurityEx2Request) (*SetSecurityEx2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) AccessCheckEx2(context.Context, *AccessCheckEx2Request) (*AccessCheckEx2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) ReassignMessage(context.Context, *ReassignMessageRequest) (*ReassignMessageResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) SetMessage(context.Context, *SetMessageRequest) (*SetMessageResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxServer) GetConfigOption(context.Context, *GetConfigOptionRequest) (*GetConfigOptionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ FaxServer = (*UnimplementedFaxServer)(nil)
