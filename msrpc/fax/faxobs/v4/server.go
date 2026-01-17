package faxobs

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

// faxobs server interface.
type FaxobsServer interface {

	// The FaxObs_ConnectionRefCount (Opnum 0) method is called by the client to connect
	// to or disconnect from the server.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the fax-specific errors that are defined in section 2.2.52
	// or one of the other standard errors defined in [MS-ERREF] section 2.2. There are
	// no predefined specific error codes to be returned by this method.
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	ConnectionReferenceCount(context.Context, *ConnectionReferenceCountRequest) (*ConnectionReferenceCountResponse, error)

	// The client calls the FaxObs_GetVersion (Opnum 1) method to obtain the version number
	// of the server.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return the following error code, one of the fax-specific errors that are
	// defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+------------------------------------------------------------+
	//	|               RETURN               |                                                            |
	//	|             VALUE/CODE             |                        DESCRIPTION                         |
	//	|                                    |                                                            |
	//	+------------------------------------+------------------------------------------------------------+
	//	+------------------------------------+------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The Version parameter is set to a NULL pointer value.<184> |
	//	+------------------------------------+------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)

	// The client calls the FaxObs_GetInstallType (Opnum 2) method to obtain information
	// about the server installation.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return the following error code, one of the fax-specific errors that are
	// defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+-----------------------------------+------------------------------------------------------------------------------+
	//	|              RETURN               |                                                                              |
	//	|            VALUE/CODE             |                                 DESCRIPTION                                  |
	//	|                                   |                                                                              |
	//	+-----------------------------------+------------------------------------------------------------------------------+
	//	+-----------------------------------+------------------------------------------------------------------------------+
	//	| ERROR_INVALID_FUNCTION 0x00000001 | The fax server is unable to retrieve the requested installation information. |
	//	+-----------------------------------+------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetInstallType(context.Context, *GetInstallTypeRequest) (*GetInstallTypeResponse, error)

	// The client calls the FaxObs_OpenPort (Opnum 3) method to open a fax port and obtain
	// a fax port handle.
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
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The caller does not have the FAX_PORT_QUERY access rights      |
	//	|                                    | required for this operation.                                                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_HANDLE 0x00000006    | The call was made with the Flags parameter containing the PORT_OPEN_MODIFY flag, |
	//	|                                    | and the port is already opened to be modified by another call.                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_BAD_UNIT 0x00000014          | The system cannot find the port for the receiving device by using the line       |
	//	|                                    | identifier specified by the DeviceId parameter.                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The FaxPortHandle parameter is set to a NULL pointer value.<188>                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	OpenPort(context.Context, *OpenPortRequest) (*OpenPortResponse, error)

	// The client calls the FaxObs_ClosePort (Opnum 4) method to close a fax port and release
	// the fax port handle obtained with a FaxObs_OpenPort (section 3.1.4.2.5) call.
	//
	// On success, the server MUST close the specified port and release the handle.
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
	//	| ERROR_ACCESS_DENIED 0x00000005 | Access is denied. The client's fax user account does not have the FAX_PORT_QUERY |
	//	|                                | access rights required for this operation.                                       |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	ClosePort(context.Context, *ClosePortRequest) (*ClosePortResponse, error)

	// The client calls the FaxObs_SendDocument (Opnum 5) method to send a document.
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
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the FAX_JOB_SUBMIT |
	//	|                                    | access rights required for this operation.                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                    | following conditions: § The JobParams parameter is set to a NULL pointer         |
	//	|                                    | value. § The FileName parameter is set to a NULL pointer value. § The FaxJobId   |
	//	|                                    | parameter is set to a NULL pointer value.<190> § The length of the character     |
	//	|                                    | string specified by the FileName parameter (excluding the length of the          |
	//	|                                    | terminating null character) plus the length of the fax queue directory path name |
	//	|                                    | (excluding the length of the terminating null character) exceeds 253 characters. |
	//	|                                    | This error can occur if the fax client is not using a file path name obtained    |
	//	|                                    | from FaxObs_GetQueueFileName. § Either of the following conditions are true: §   |
	//	|                                    | The first value of the Reserved field of the structure referenced by JobParams   |
	//	|                                    | is set to 0xFFFFFFFE (32-bit) or 0x00000000FFFFFFFE (64-bit). The second value   |
	//	|                                    | of this same Reserved field is set to 0x00000002 (32-bit) or 0x0000000000000002  |
	//	|                                    | (64-bit). The RecipientNumber field of the same structure is set to NULL. § The  |
	//	|                                    | first value of the Reserved field of the structure referenced by JobParams is    |
	//	|                                    | not set to 0xFFFFFFFE (32-bit) or 0x00000000FFFFFFFE (64-bit). The CallHandle    |
	//	|                                    | field of the same structure is not set to 0x00000000. The RecipientNumber field  |
	//	|                                    | of the JobParams structure is NULL.                                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	SendDocument(context.Context, *SendDocumentRequest) (*SendDocumentResponse, error)

	// The client calls FaxObs_GetQueueFileName (Opnum 6) to obtain from the server the
	// name of a new file located in the fax server queue directory. The client can copy
	// to this file the fax data to be transmitted and submit the file name to FaxObs_SendDocument
	// (section 3.1.4.2.7).
	//
	// The client SHOULD set the FileNameSize parameter to a value of 255 characters. If
	// the client sets FileNameSize to a value exceeding 255 characters, the server SHOULD
	// NOT write more than 255 characters, including the terminating null character, to
	// the FileName output parameter.
	//
	// The client's fax user account SHOULD have write file access under the fax server
	// queue directory.<191>
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the fax-specific errors that are defined in section 2.2.52
	// or one of the other standard errors defined in [MS-ERREF] section 2.2. There are
	// no predefined specific error codes to be returned by this method.
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetQueueFileName(context.Context, *GetQueueFileNameRequest) (*GetQueueFileNameResponse, error)

	// The FaxObs_EnumJobs (Opnum 7) method is called by the client to enumerate the fax
	// jobs on the server.
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
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the FAX_ JOB_QUERY |
	//	|                                    | access rights.                                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The fax server cannot allocate memory for the data to be returned to the client. |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned when any of the          |
	//	|                                    | following conditions are met: § The Buffer parameter is set to a NULL pointer    |
	//	|                                    | value. § The BufferSize parameter is set to a NULL pointer value. <194> § The    |
	//	|                                    | JobsReturned parameter is set to a NULL pointer value. <195>                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EnumJobs(context.Context, *EnumJobsRequest) (*EnumJobsResponse, error)

	// The FaxObs_GetJob (Opnum 8) method is called by the client to retrieve information
	// regarding a specific job. The job is specified by the JobId parameter. The value
	// for JobId can be obtained by calling the FaxObs_EnumJobs (section 3.1.4.2.9) or FaxObs_SendDocument
	// (section 3.1.4.2.7) method.
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
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the FAX_JOB_QUERY  |
	//	|                                    | access rights.                                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The fax server cannot allocate memory for the data to be returned to the client. |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	GetJob(context.Context, *GetJobRequest) (*GetJobResponse, error)

	// The FaxObs_SetJob (Opnum 9) method is called by the client to pause, resume, or cancel
	// a fax job. The value for the JobId parameter can be obtained by calling the FaxObs_EnumJobs
	// (section 3.1.4.2.9) or FaxObs_SendDocument (section 3.1.4.2.7) method.
	//
	// On success, the server MUST pause, resume, or cancel the specified fax job and MUST
	// set the job status (section 3.1.1) to reflect the new job state.
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
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. This error code is returned under any of the following         |
	//	|                                    | conditions: § The client's fax user account does not have the FAX_ JOB_MANAGE    |
	//	|                                    | access rights. § The client's fax user account is not the owner of the fax job   |
	//	|                                    | identified by the JobId parameter.                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                    | following conditions: § The fax job indicated by JobId cannot be found by the    |
	//	|                                    | fax server. § The specified Command parameter value is not JC_DELETE, JC_PAUSE,  |
	//	|                                    | or JC_RESUME. § The specified Command value is JC_DELETE, the specified JobId    |
	//	|                                    | represents the job for an outgoing broadcast message, and aborting outgoing      |
	//	|                                    | broadcast messages is not supported. For more information, see FaxObs_Abort. §   |
	//	|                                    | The JobId is set to a NULL pointer value.                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	SetJob(context.Context, *SetJobRequest) (*SetJobResponse, error)

	// The FaxObs_GetPageData (Opnum 10) method is called by the client to retrieve the
	// data from the first page of an outgoing fax job. The information that is returned
	// in the buffer is an in-memory copy of the first page of the TIFF file for the job.
	// The value for the JobId parameter can be obtained by calling the FaxObs_EnumJobs
	// (section 3.1.4.2.9) method.
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
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the FAX_JOB_QUERY  |
	//	|                                    | access rights.                                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The fax server cannot allocate memory for the data to be returned to the client. |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | The job identified by the JobId parameter is not an outgoing fax job.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                    | following conditions: § One or more of the following parameters are set to NULL  |
	//	|                                    | pointer values: Buffer, BufferSize, ImageWidth, and ImageHeight.<197> § The fax  |
	//	|                                    | server cannot find the fax job indicated by the JobId parameter.                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetPageData(context.Context, *GetPageDataRequest) (*GetPageDataResponse, error)

	// The FaxObs_GetDeviceStatus (Opnum 11) method is called by the client to retrieve
	// information about a specified fax device (port).
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
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the FAX_PORT_QUERY |
	//	|                                    | access rights.                                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The fax server cannot allocate memory for the data to be returned to the client. |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | This error SHOULD be returned if the FaxPortHandle parameter is not a valid      |
	//	|                                    | handle obtained using FaxObs_OpenPort.<198>                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetDeviceStatus(context.Context, *GetDeviceStatusRequest) (*GetDeviceStatusResponse, error)

	// The FaxObs_Abort (Opnum 12) method is called by the client to abort the specified
	// fax job on the server. The value for the JobId parameter can be obtained by calling
	// the FaxObs_EnumJobs (section 3.1.4.2.9) or FaxObs_SendDocument (section 3.1.4.2.7)
	// method.
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
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. This error code is returned under any of the following         |
	//	|                                    | conditions: § The client's fax user account does not have the FAX_ JOB_MANAGE    |
	//	|                                    | access rights. § The client's fax user account is not the owner of the fax job   |
	//	|                                    | identified by the JobId parameter.                                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned when any of the          |
	//	|                                    | following conditions are met: § The fax job identified by the specified JobId    |
	//	|                                    | cannot be found by the fax server. § The specified job has already been canceled |
	//	|                                    | or is in the process of being canceled. § The specified JobId represents the job |
	//	|                                    | for an outgoing broadcast message; aborting outgoing broadcast messages is not   |
	//	|                                    | supported.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	Abort(context.Context, *AbortRequest) (*AbortResponse, error)

	// The FaxObs_EnumPorts (Opnum 13) method is called by the client to enumerate the fax
	// ports (devices) on the server and retrieve information describing these ports (devices).
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
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the FAX_           |
	//	|                                    | PORT_QUERY access rights.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The fax server cannot allocate sufficient memory to hold the array of            |
	//	|                                    | _FAX_PORT_INFO structures to be returned to the client.                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The PortsReturned parameter is set to a NULL pointer value.<199>                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EnumPorts(context.Context, *EnumPortsRequest) (*EnumPortsResponse, error)

	// The FaxObs_GetPort (Opnum 14) method is called by the client to retrieve status information
	// from the server about the specified fax port (device).
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
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the FAX_PORT_QUERY |
	//	|                                    | access rights.                                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The fax server cannot allocate sufficient memory to hold the _FAX_PORT_INFO to   |
	//	|                                    | be returned to the client.                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | This error SHOULD be returned if the FaxPortHandle argument is not a valid       |
	//	|                                    | handle obtained using FaxObs_OpenPort.<200>                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetPort(context.Context, *GetPortRequest) (*GetPortResponse, error)

	// The client calls the FaxObs_SetPort (Opnum 15) method to change the configuration
	// of a fax port (device). The function sets extension configuration properties that
	// are stored at the device level, such as enabling or disabling sending and receiving,
	// and the automatic or manual answering of calls.
	//
	// The client MUST set the FaxPortHandle parameter to a valid fax port handle value
	// returned by the FaxObs_OpenPort (section 3.1.4.2.5) method. The server MUST validate
	// that the client's fax user account has the access to change the port configuration.
	// On success, the server MUST modify the properties of the port as specified by the
	// client.
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
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the FAX_PORT_SET   |
	//	|                                    | access rights.                                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | The handle specified by the FaxPortHandle argument is not a valid fax port       |
	//	|                                    | handle obtained by a call to FaxObs_OpenPort.<201>                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The size of FAX_PORT_INFO, specified in the SizeOfStruct field, is incorrect.    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_DEVICE_IN_USE 0x00000964     | The specified fax port is currently sending or receiving a fax transmission.     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	SetPort(context.Context, *SetPortRequest) (*SetPortResponse, error)

	// The FaxObs_EnumRoutingMethods (Opnum 16) method is called by the client to enumerate
	// all of the routing methods that are registered with the server for a specified port
	// (device). The function returns detailed information about each enumerated routing
	// method.
	//
	// The client MUST set the FaxPortHandle parameter to a valid fax port handle value
	// returned by the FaxObs_OpenPort (section 3.1.4.2.5) method.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the following error codes, one of the fax-specific errors that
	// are defined in section 2.2.52 or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_FUNCTION 0x00000001  | There are no routing methods configured on the fax server for the fax port       |
	//	|                                    | specified through the FaxPortHandle parameter.                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the FAX_PORT_QUERY |
	//	|                                    | access rights.                                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The fax server cannot allocate sufficient memory to hold the array of            |
	//	|                                    | FAX_ROUTING_METHOD structures to be returned to the client.                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | This error SHOULD be returned if the FaxPortHandle argument is not a valid       |
	//	|                                    | handle obtained using FaxObs_OpenPort.<202>                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned if any of the            |
	//	|                                    | following parameters is set to a NULL pointer value: RoutingInfoBuffer,          |
	//	|                                    | RoutingInfoBufferSize, PortsReturned.<203>                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EnumRoutingMethods(context.Context, *EnumRoutingMethodsRequest) (*EnumRoutingMethodsResponse, error)

	// The FaxObs_EnableRoutingMethod (Opnum 17) method is called by the client to enable
	// or disable a routing method for a fax port (device).
	//
	// The client MUST set the FaxPortHandle parameter to a valid fax port handle value
	// returned by the FaxObs_OpenPort (section 3.1.4.2.5) method. In response, the server
	// MUST validate that the client's fax user account has access to enable or disable
	// routing methods. The client MUST set the RoutingGUID parameter to point to a valid
	// routing method.
	//
	// On success, the server MUST enable or disable a fax routing method for the specified
	// fax port (device).
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
	//	| ERROR_ACCESS_DENIED 0x00000005 | Access is denied. The client's fax user account does not have the FAX_PORT_SET   |
	//	|                                | access rights.                                                                   |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D  | The data is invalid. This error code is returned when any of the following       |
	//	|                                | conditions are met: § The FaxPortHandle argument is not a valid handle obtained  |
	//	|                                | using FaxObs_OpenPort.<204> § The GUID specified by the RoutingGuid parameter is |
	//	|                                | not a routing method GUID.                                                       |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EnableRoutingMethod(context.Context, *EnableRoutingMethodRequest) (*EnableRoutingMethodResponse, error)

	// The FaxObs_GetRoutingInfo (Opnum 18) method is called by the client to retrieve information
	// about a routing method.
	//
	// The client MUST set the FaxPortHandle parameter to a valid fax port handle value
	// returned by the FaxObs_OpenPort (section 3.1.4.2.5) method.
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
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the FAX_PORT_QUERY |
	//	|                                    | access rights.                                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The fax server cannot allocate sufficient memory for the data to be returned to  |
	//	|                                    | the client.                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | The data is invalid. This error code is returned when any of the following       |
	//	|                                    | conditions are met: § The FaxPortHandle argument is not a valid handle obtained  |
	//	|                                    | using FaxObs_OpenPort.<205> § The GUID specified by the RoutingGuid parameter is |
	//	|                                    | not a routing method GUID.                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. One or more of the following parameters are set to   |
	//	|                                    | NULL pointer values: RoutingGuid, RoutingInfoBuffer, RoutingInfoBufferSize.<206> |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetRoutingInfo(context.Context, *GetRoutingInfoRequest) (*GetRoutingInfoResponse, error)

	// The FaxObs_SetRoutingInfo (Opnum 19) method is called by the client to set routing
	// information for a fax routing method.
	//
	// The client MUST set the FaxPortHandle parameter to a valid fax port handle value
	// returned by the FaxObs_OpenPort (section 3.1.4.2.5) method.
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
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the FAX_PORT_SET   |
	//	|                                    | access rights.                                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | The data is invalid. This error code is returned when any of the following       |
	//	|                                    | conditions are met: § The FaxPortHandle argument is not a valid handle obtained  |
	//	|                                    | using FaxObs_OpenPort.<207> § The GUID specified by the RoutingGuid parameter is |
	//	|                                    | not a routing method GUID. § The routing extension specified by the RoutingGuid  |
	//	|                                    | or the routing extension implementing the routing method specified by the        |
	//	|                                    | RoutingGuid denies the request to set the data requested by the fax client.      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. One or more of the following parameters are set to   |
	//	|                                    | NULL pointer values: RoutingGuid, RoutingInfoBuffer, RoutingInfoBufferSize.      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	SetRoutingInfo(context.Context, *SetRoutingInfoRequest) (*SetRoutingInfoResponse, error)

	// The FaxObs_EnumGlobalRoutingInfo (Opnum 20) method is called by the client to enumerate
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
	//	| ERROR_INVALID_FUNCTION 0x00000001  | There are no routing methods currently configured on the fax server.             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the                |
	//	|                                    | FAX_CONFIG_QUERY access rights.                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The fax server cannot allocate sufficient memory to hold the array of            |
	//	|                                    | _FAX_GLOBAL_ROUTING_INFOW structures to be returned to the client.               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. One or more of the following parameters are          |
	//	|                                    | set to NULL pointer values: RoutingInfoBuffer, RoutingInfoBufferSize, and        |
	//	|                                    | MethodsReturned. <208>                                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	EnumGlobalRoutingInfo(context.Context, *EnumGlobalRoutingInfoRequest) (*EnumGlobalRoutingInfoResponse, error)

	// The fax client calls the FaxObs_SetGlobalRoutingInfo (Opnum 21) method to set global
	// routing properties, such as the routing method priority.
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
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the FAX_CONFIG_SET |
	//	|                                    | access rights.                                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | The fax server cannot find the routing method specified by the Guid structure    |
	//	|                                    | field of the RoutingInfo parameter.                                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned if any of the following  |
	//	|                                    | conditions are met: § The RoutingInfo is set to a NULL pointer value. § The      |
	//	|                                    | SizeOfStruct structure field of the RoutingInfo is not set to the correct size   |
	//	|                                    | in bytes for the FAX_GLOBAL_ROUTING_INFO.                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	SetGlobalRoutingInfo(context.Context, *SetGlobalRoutingInfoRequest) (*SetGlobalRoutingInfoResponse, error)

	// The FaxObs_GetConfiguration (Opnum 22) method is called by the client to retrieve
	// information about the configuration of the fax server.
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
	//	|                                    | FAX_CONFIG_QUERY access rights.                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The fax server cannot allocate sufficient memory to hold the _FAX_CONFIGURATIONW |
	//	|                                    | to be returned to the client.                                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned if the Buffer or         |
	//	|                                    | BufferSize parameters are set to NULL pointer values.                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetConfiguration(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error)

	// The client calls the FaxObs_SetConfiguration (Opnum 23) method to change the fax
	// server configuration.
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
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the FAX_CONFIG_SET |
	//	|                                    | access rights.                                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | The fax server failed to add (apply) the new MAPI profile specified by the       |
	//	|                                    | Reserved member of the FaxConfig structure.                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                    | following conditions: § The pointer specified with the FaxConfig argument is     |
	//	|                                    | NULL. § The dwSizeOfStruct member of the FAX_CONFIGURATIONW specified by the     |
	//	|                                    | FaxConfig parameter is set to an incorrect value. § The ArchiveOutgoingFaxes     |
	//	|                                    | member of the FaxConfig structure is set to TRUE and the ArchiveDirectory member |
	//	|                                    | of this same structure is set to a NULL pointer value.                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	SetConfiguration(context.Context, *SetConfigurationRequest) (*SetConfigurationResponse, error)

	// The FaxObs_GetLoggingCategories (Opnum 24) method is called by the client to retrieve
	// the current logging categories from the server. A logging category determines the
	// errors or other events that the fax server records in the application event log.
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
	//	|                                    | FAX_CONFIG_QUERY access rights.                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The fax server cannot allocate sufficient memory to hold the array of            |
	//	|                                    | FAX_LOG_CATEGORY structures to be returned to the client.                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error is returned if one of more of the         |
	//	|                                    | following parameters are set to NULL pointer values: Buffer, BufferSize,         |
	//	|                                    | NumberCategories. <210>                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetLoggingCategories(context.Context, *GetLoggingCategoriesRequest) (*GetLoggingCategoriesResponse, error)

	// The FaxObs_SetLoggingCategories (Opnum 25) method is called by the client to set
	// the current logging categories on the server. A logging category determines the errors
	// or other events that the fax server records in the application event log.
	//
	// On success, the server MUST modify its current logging categories.
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
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the FAX_CONFIG_SET |
	//	|                                    | access rights.                                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                    | following conditions: § The value specified for the Buffer parameter is NULL. §  |
	//	|                                    | The value specified for the BufferSize parameter is 0.                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	SetLoggingCategories(context.Context, *SetLoggingCategoriesRequest) (*SetLoggingCategoriesResponse, error)

	// The FaxObs_GetTapiLocations (Opnum 26) method is called by the client to retrieve
	// the current and other available TAPI locations configured for the server. The TAPI
	// locations can be set by the client with the FaxObs_SetTapiLocations (section 3.1.4.2.29)
	// method.
	//
	// A TAPI location is described by a FAX_TAPI_LOCATIONS (section 2.2.88) data structure,
	// which includes information such as a friendly name, country code, and area code for
	// the respective location. For more information about TAPI see [MSDN-TAPI2.2].
	//
	// On success, the server MUST allocate memory for and return the data describing its
	// current and other available TAPI locations.
	//
	// The client SHOULD free the buffer.
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
	//	|                                    | FAX_CONFIG_QUERY access rights.                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The fax server cannot allocate sufficient memory to hold the data to be returned |
	//	|                                    | to the client.                                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error is returned if any of the following       |
	//	|                                    | parameters are set to a NULL pointer value: Buffer or BufferSize.<211>           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetTAPILocations(context.Context, *GetTAPILocationsRequest) (*GetTAPILocationsResponse, error)

	// The FaxObs_SetTapiLocations (Opnum 27) method is called by the client to set the
	// current and other available TAPI locations for the server. The TAPI locations can
	// be retrieved by the client with the FaxObs_GetTapiLocations (section 3.1.4.2.28)
	// method.
	//
	// A TAPI location is described by a FAX_TAPI_LOCATIONS (section 2.2.88) data structure,
	// which includes information such as a friendly name, country code, and area code for
	// the respective location. For more information about TAPI, see [MSDN-TAPI2.2].
	//
	// On success, the server MUST apply the new locations configuration that was requested
	// by the client.
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
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the FAX_CONFIG_SET |
	//	|                                    | access rights.                                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error is returned if the Buffer parameter is    |
	//	|                                    | set to a NULL pointer value.                                                     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	SetTAPILocations(context.Context, *SetTAPILocationsRequest) (*SetTAPILocationsResponse, error)

	// The FaxObs_GetMapiProfiles (Opnum 28) method is called by the client to retrieve
	// the names of the current MAPI profiles set for the fax server. For more information
	// about MAPI profiles, see [MSDN-MAPIPRF].
	//
	// On success, the server MUST allocate memory for, and return the list of, the current
	// MAPI profile names. They MUST be formatted as a sequence of null-terminated character
	// strings, with the sequence terminated by a single empty, null-terminated character
	// string.
	//
	// The client SHOULD free the buffer.
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
	//	|                                    | FAX_CONFIG_QUERY access rights.                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error is returned if the MapiProfiles or the    |
	//	|                                    | BufferSize parameters are set to NULL pointer values. <212>                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetMAPIProfiles(context.Context, *GetMAPIProfilesRequest) (*GetMAPIProfilesResponse, error)

	// The FaxObs_StartClientServer (Opnum 29) method is called by the client to register
	// itself to receive notifications of fax events from the server.
	//
	// On success, the server MUST start notifying the client about the occurring fax events.
	//
	// If the server receives a FaxObs_StartClientServer request for a client machine name
	// and client name that are already registered to receive fax event notifications with
	// a previously executed FaxObs_StartClientServer call, the fax server MUST consider
	// the new request a success and keep the existing fax client registration.
	//
	// To notify the client about a fax event, the server MUST first open a connection with
	// the client by calling the FAX_OpenConnection (section 3.2.4.5) method. Then the fax
	// server MUST notify the client by calling the FAX_ClientEventQueue (section 3.2.4.2)
	// method. Finally, the server SHOULD close the connection with the client by calling
	// the FAX_CloseConnection (section 3.2.4.4) method. <213>
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return the following error code, one of the fax-specific errors that are
	// defined in section 2.2.52, or one of the other standard errors defined in [MS-ERREF]
	// section 2.2.
	//
	//	+------------------------------------+--------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                |
	//	|             VALUE/CODE             |                                  DESCRIPTION                                   |
	//	|                                    |                                                                                |
	//	+------------------------------------+--------------------------------------------------------------------------------+
	//	+------------------------------------+--------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The fax server cannot allocate memory necessary for the fax client connection. |
	//	+------------------------------------+--------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	StartClientServer(context.Context, *StartClientServerRequest) (*StartClientServerResponse, error)

	// Opnum30NotUsedOnWire operation.
	// Opnum30NotUsedOnWire

	// The client calls FaxObs_GetSecurityDescriptor (Opnum 31) method to retrieve the fax
	// security descriptor of the server. The client can set the fax security descriptor
	// of the server with the FaxObs_SetSecurityDescriptor (section 3.1.4.2.33) method.
	//
	// On success, the server MUST allocate memory for the return data buffer and return
	// a FAX_SECURITY_DESCRIPTOR (section 2.2.90) structure.
	//
	// The client SHOULD free the buffer.
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
	//	|                                    | FAX_CONFIG_QUERY access rights.                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_NOT_ENOUGH_MEMORY 0x00000008 | The server cannot allocate sufficient memory to hold the FAX_SECURITY_DESCRIPTOR |
	//	|                                    | to be returned to the client.                                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_CATEGORY 0x00000075  | The server SHOULD return this error code if the value of the Id parameter is     |
	//	|                                    | greater than 0.<214>                                                             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetSecurityDescriptor(context.Context, *GetSecurityDescriptorRequest) (*GetSecurityDescriptorResponse, error)

	// The client calls FaxObs_SetSecurityDescriptor (Opnum 32) method to set the fax security
	// descriptor of the server. The client can retrieve the security descriptor of the
	// server with the FaxObs_GetSecurityDescriptor (section 3.1.4.2.32) method.
	//
	// On success, the server MUST apply the security descriptor described in the submitted
	// FAX_SECURITY_DESCRIPTOR (section 2.2.90) structure.
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
	//	| ERROR_ACCESS_DENIED 0x00000005     | Access is denied. The client's fax user account does not have the FAX_CONFIG_SET |
	//	|                                    | access rights.                                                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_DATA 0x0000000D      | The security descriptor described by the FAX_SECURITY_DESCRIPTOR pointed at by   |
	//	|                                    | the FaxSecurityDecriptor parameter is invalid.                                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned under any of the         |
	//	|                                    | following conditions: § The value of the BufferSize parameter is less than the   |
	//	|                                    | size of the Fixed_Portion block of the FAX_SECURITY_DESCRIPTOR. § An offset      |
	//	|                                    | into the Fixed_Portion block of the FAX_SECURITY_DESCRIPTOR pointed at by the    |
	//	|                                    | FaxSecurityDecriptor points to outside of the buffer.                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| ERROR_INVALID_CATEGORY 0x00000075  | The fax server SHOULD return this error if the Id member of the                  |
	//	|                                    | FAX_SECURITY_DESCRIPTOR specified by the FaxSecurityDescriptor argument is set   |
	//	|                                    | to a value greater than 0.<215>                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	SetSecurityDescriptor(context.Context, *SetSecurityDescriptorRequest) (*SetSecurityDescriptorResponse, error)

	// The client calls the FaxObs_GetSecurityDescriptorCount (Opnum 33) method to retrieve
	// the total number of fax security descriptors from the server.
	//
	// On success, the server MUST return the total number of security descriptors. This
	// number MUST be 1.
	//
	// Return Values: This method MUST return 0x00000000 (ERROR_SUCCESS) for success; otherwise,
	// it MUST return one of the fax-specific errors that are defined in section 2.2.52
	// or one of the other standard errors defined in [MS-ERREF] section 2.2. There are
	// no predefined error codes to be returned by this method.
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	GetSecurityDescriptorCount(context.Context, *GetSecurityDescriptorCountRequest) (*GetSecurityDescriptorCountResponse, error)

	// The FaxObs_AccessCheck (Opnum 34) method is called by the client to check whether
	// the currently logged-on client user account has access permissions to execute specific
	// fax operations on the fax server.
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
	//	| ERROR_INVALID_PARAMETER 0x00000057 | The parameter is incorrect. This error code is returned if any of the following  |
	//	|                                    | conditions are met: § The hBinding parameter is set to a NULL value. <216> § The |
	//	|                                    | fAccess parameter is set to a NULL pointer value. <217>                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown except those that are thrown by the underlying
	// RPC protocol, [MS-RPCE].
	AccessCheck(context.Context, *AccessCheckRequest) (*AccessCheckResponse, error)
}

func RegisterFaxobsServer(conn dcerpc.Conn, o FaxobsServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewFaxobsServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(FaxobsSyntaxV4_0))...)
}

func NewFaxobsServerHandle(o FaxobsServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return FaxobsServerHandle(ctx, o, opNum, r)
	}
}

func FaxobsServerHandle(ctx context.Context, o FaxobsServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // FaxObs_ConnectionRefCount
		op := &xxx_ConnectionReferenceCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ConnectionReferenceCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ConnectionReferenceCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // FaxObs_GetVersion
		op := &xxx_GetVersionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetVersionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetVersion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // FaxObs_GetInstallType
		op := &xxx_GetInstallTypeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetInstallTypeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetInstallType(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // FaxObs_OpenPort
		op := &xxx_OpenPortOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenPortRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenPort(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // FaxObs_ClosePort
		op := &xxx_ClosePortOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ClosePortRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ClosePort(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // FaxObs_SendDocument
		op := &xxx_SendDocumentOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SendDocumentRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SendDocument(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // FaxObs_GetQueueFileName
		op := &xxx_GetQueueFileNameOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetQueueFileNameRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetQueueFileName(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // FaxObs_EnumJobs
		op := &xxx_EnumJobsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumJobsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumJobs(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // FaxObs_GetJob
		op := &xxx_GetJobOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetJobRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetJob(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // FaxObs_SetJob
		op := &xxx_SetJobOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetJobRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetJob(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // FaxObs_GetPageData
		op := &xxx_GetPageDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPageDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPageData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // FaxObs_GetDeviceStatus
		op := &xxx_GetDeviceStatusOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDeviceStatusRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDeviceStatus(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // FaxObs_Abort
		op := &xxx_AbortOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AbortRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Abort(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // FaxObs_EnumPorts
		op := &xxx_EnumPortsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumPortsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumPorts(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // FaxObs_GetPort
		op := &xxx_GetPortOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetPortRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetPort(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // FaxObs_SetPort
		op := &xxx_SetPortOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetPortRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetPort(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // FaxObs_EnumRoutingMethods
		op := &xxx_EnumRoutingMethodsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumRoutingMethodsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumRoutingMethods(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // FaxObs_EnableRoutingMethod
		op := &xxx_EnableRoutingMethodOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnableRoutingMethodRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnableRoutingMethod(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // FaxObs_GetRoutingInfo
		op := &xxx_GetRoutingInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetRoutingInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetRoutingInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // FaxObs_SetRoutingInfo
		op := &xxx_SetRoutingInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetRoutingInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetRoutingInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // FaxObs_EnumGlobalRoutingInfo
		op := &xxx_EnumGlobalRoutingInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumGlobalRoutingInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumGlobalRoutingInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // FaxObs_SetGlobalRoutingInfo
		op := &xxx_SetGlobalRoutingInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetGlobalRoutingInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetGlobalRoutingInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // FaxObs_GetConfiguration
		op := &xxx_GetConfigurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetConfigurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetConfiguration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // FaxObs_SetConfiguration
		op := &xxx_SetConfigurationOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetConfigurationRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetConfiguration(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // FaxObs_GetLoggingCategories
		op := &xxx_GetLoggingCategoriesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetLoggingCategoriesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetLoggingCategories(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // FaxObs_SetLoggingCategories
		op := &xxx_SetLoggingCategoriesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetLoggingCategoriesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetLoggingCategories(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 26: // FaxObs_GetTapiLocations
		op := &xxx_GetTAPILocationsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetTAPILocationsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetTAPILocations(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // FaxObs_SetTapiLocations
		op := &xxx_SetTAPILocationsOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetTAPILocationsRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetTAPILocations(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // FaxObs_GetMapiProfiles
		op := &xxx_GetMAPIProfilesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetMAPIProfilesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetMAPIProfiles(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 29: // FaxObs_StartClientServer
		op := &xxx_StartClientServerOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &StartClientServerRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.StartClientServer(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // Opnum30NotUsedOnWire
		// Opnum30NotUsedOnWire
		return nil, nil
	case 31: // FaxObs_GetSecurityDescriptor
		op := &xxx_GetSecurityDescriptorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSecurityDescriptorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSecurityDescriptor(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // FaxObs_SetSecurityDescriptor
		op := &xxx_SetSecurityDescriptorOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetSecurityDescriptorRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetSecurityDescriptor(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // FaxObs_GetSecurityDescriptorCount
		op := &xxx_GetSecurityDescriptorCountOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSecurityDescriptorCountRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSecurityDescriptorCount(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // FaxObs_AccessCheck
		op := &xxx_AccessCheckOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AccessCheckRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AccessCheck(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented faxobs
type UnimplementedFaxobsServer struct {
}

func (UnimplementedFaxobsServer) ConnectionReferenceCount(context.Context, *ConnectionReferenceCountRequest) (*ConnectionReferenceCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) GetInstallType(context.Context, *GetInstallTypeRequest) (*GetInstallTypeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) OpenPort(context.Context, *OpenPortRequest) (*OpenPortResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) ClosePort(context.Context, *ClosePortRequest) (*ClosePortResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) SendDocument(context.Context, *SendDocumentRequest) (*SendDocumentResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) GetQueueFileName(context.Context, *GetQueueFileNameRequest) (*GetQueueFileNameResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) EnumJobs(context.Context, *EnumJobsRequest) (*EnumJobsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) GetJob(context.Context, *GetJobRequest) (*GetJobResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) SetJob(context.Context, *SetJobRequest) (*SetJobResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) GetPageData(context.Context, *GetPageDataRequest) (*GetPageDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) GetDeviceStatus(context.Context, *GetDeviceStatusRequest) (*GetDeviceStatusResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) Abort(context.Context, *AbortRequest) (*AbortResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) EnumPorts(context.Context, *EnumPortsRequest) (*EnumPortsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) GetPort(context.Context, *GetPortRequest) (*GetPortResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) SetPort(context.Context, *SetPortRequest) (*SetPortResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) EnumRoutingMethods(context.Context, *EnumRoutingMethodsRequest) (*EnumRoutingMethodsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) EnableRoutingMethod(context.Context, *EnableRoutingMethodRequest) (*EnableRoutingMethodResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) GetRoutingInfo(context.Context, *GetRoutingInfoRequest) (*GetRoutingInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) SetRoutingInfo(context.Context, *SetRoutingInfoRequest) (*SetRoutingInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) EnumGlobalRoutingInfo(context.Context, *EnumGlobalRoutingInfoRequest) (*EnumGlobalRoutingInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) SetGlobalRoutingInfo(context.Context, *SetGlobalRoutingInfoRequest) (*SetGlobalRoutingInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) GetConfiguration(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) SetConfiguration(context.Context, *SetConfigurationRequest) (*SetConfigurationResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) GetLoggingCategories(context.Context, *GetLoggingCategoriesRequest) (*GetLoggingCategoriesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) SetLoggingCategories(context.Context, *SetLoggingCategoriesRequest) (*SetLoggingCategoriesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) GetTAPILocations(context.Context, *GetTAPILocationsRequest) (*GetTAPILocationsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) SetTAPILocations(context.Context, *SetTAPILocationsRequest) (*SetTAPILocationsResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) GetMAPIProfiles(context.Context, *GetMAPIProfilesRequest) (*GetMAPIProfilesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) StartClientServer(context.Context, *StartClientServerRequest) (*StartClientServerResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) GetSecurityDescriptor(context.Context, *GetSecurityDescriptorRequest) (*GetSecurityDescriptorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) SetSecurityDescriptor(context.Context, *SetSecurityDescriptorRequest) (*SetSecurityDescriptorResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) GetSecurityDescriptorCount(context.Context, *GetSecurityDescriptorCountRequest) (*GetSecurityDescriptorCountResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedFaxobsServer) AccessCheck(context.Context, *AccessCheckRequest) (*AccessCheckResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ FaxobsServer = (*UnimplementedFaxobsServer)(nil)
