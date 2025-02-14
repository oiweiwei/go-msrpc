package faxobs

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	fax "github.com/oiweiwei/go-msrpc/msrpc/fax"
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
	_ = fax.GoPackage
)

var (
	// import guard
	GoPackage = "fax"
)

var (
	// Syntax UUID
	FaxobsSyntaxUUID = &uuid.UUID{TimeLow: 0xea0a3165, TimeMid: 0x4834, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0xa6, ClockSeqLow: 0xf8, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xa3, 0x46, 0xcc}}
	// Syntax ID
	FaxobsSyntaxV4_0 = &dcerpc.SyntaxID{IfUUID: FaxobsSyntaxUUID, IfVersionMajor: 4, IfVersionMinor: 0}
)

// faxobs interface.
type FaxobsClient interface {

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
	ConnectionReferenceCount(context.Context, *ConnectionReferenceCountRequest, ...dcerpc.CallOption) (*ConnectionReferenceCountResponse, error)

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
	GetVersion(context.Context, *GetVersionRequest, ...dcerpc.CallOption) (*GetVersionResponse, error)

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
	GetInstallType(context.Context, *GetInstallTypeRequest, ...dcerpc.CallOption) (*GetInstallTypeResponse, error)

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
	OpenPort(context.Context, *OpenPortRequest, ...dcerpc.CallOption) (*OpenPortResponse, error)

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
	ClosePort(context.Context, *ClosePortRequest, ...dcerpc.CallOption) (*ClosePortResponse, error)

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
	SendDocument(context.Context, *SendDocumentRequest, ...dcerpc.CallOption) (*SendDocumentResponse, error)

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
	GetQueueFileName(context.Context, *GetQueueFileNameRequest, ...dcerpc.CallOption) (*GetQueueFileNameResponse, error)

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
	EnumJobs(context.Context, *EnumJobsRequest, ...dcerpc.CallOption) (*EnumJobsResponse, error)

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
	GetJob(context.Context, *GetJobRequest, ...dcerpc.CallOption) (*GetJobResponse, error)

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
	SetJob(context.Context, *SetJobRequest, ...dcerpc.CallOption) (*SetJobResponse, error)

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
	GetPageData(context.Context, *GetPageDataRequest, ...dcerpc.CallOption) (*GetPageDataResponse, error)

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
	GetDeviceStatus(context.Context, *GetDeviceStatusRequest, ...dcerpc.CallOption) (*GetDeviceStatusResponse, error)

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
	Abort(context.Context, *AbortRequest, ...dcerpc.CallOption) (*AbortResponse, error)

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
	EnumPorts(context.Context, *EnumPortsRequest, ...dcerpc.CallOption) (*EnumPortsResponse, error)

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
	GetPort(context.Context, *GetPortRequest, ...dcerpc.CallOption) (*GetPortResponse, error)

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
	SetPort(context.Context, *SetPortRequest, ...dcerpc.CallOption) (*SetPortResponse, error)

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
	EnumRoutingMethods(context.Context, *EnumRoutingMethodsRequest, ...dcerpc.CallOption) (*EnumRoutingMethodsResponse, error)

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
	EnableRoutingMethod(context.Context, *EnableRoutingMethodRequest, ...dcerpc.CallOption) (*EnableRoutingMethodResponse, error)

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
	GetRoutingInfo(context.Context, *GetRoutingInfoRequest, ...dcerpc.CallOption) (*GetRoutingInfoResponse, error)

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
	SetRoutingInfo(context.Context, *SetRoutingInfoRequest, ...dcerpc.CallOption) (*SetRoutingInfoResponse, error)

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
	EnumGlobalRoutingInfo(context.Context, *EnumGlobalRoutingInfoRequest, ...dcerpc.CallOption) (*EnumGlobalRoutingInfoResponse, error)

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
	SetGlobalRoutingInfo(context.Context, *SetGlobalRoutingInfoRequest, ...dcerpc.CallOption) (*SetGlobalRoutingInfoResponse, error)

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
	GetConfiguration(context.Context, *GetConfigurationRequest, ...dcerpc.CallOption) (*GetConfigurationResponse, error)

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
	SetConfiguration(context.Context, *SetConfigurationRequest, ...dcerpc.CallOption) (*SetConfigurationResponse, error)

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
	GetLoggingCategories(context.Context, *GetLoggingCategoriesRequest, ...dcerpc.CallOption) (*GetLoggingCategoriesResponse, error)

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
	SetLoggingCategories(context.Context, *SetLoggingCategoriesRequest, ...dcerpc.CallOption) (*SetLoggingCategoriesResponse, error)

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
	GetTAPILocations(context.Context, *GetTAPILocationsRequest, ...dcerpc.CallOption) (*GetTAPILocationsResponse, error)

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
	SetTAPILocations(context.Context, *SetTAPILocationsRequest, ...dcerpc.CallOption) (*SetTAPILocationsResponse, error)

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
	GetMAPIProfiles(context.Context, *GetMAPIProfilesRequest, ...dcerpc.CallOption) (*GetMAPIProfilesResponse, error)

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
	StartClientServer(context.Context, *StartClientServerRequest, ...dcerpc.CallOption) (*StartClientServerResponse, error)

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
	GetSecurityDescriptor(context.Context, *GetSecurityDescriptorRequest, ...dcerpc.CallOption) (*GetSecurityDescriptorResponse, error)

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
	SetSecurityDescriptor(context.Context, *SetSecurityDescriptorRequest, ...dcerpc.CallOption) (*SetSecurityDescriptorResponse, error)

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
	GetSecurityDescriptorCount(context.Context, *GetSecurityDescriptorCountRequest, ...dcerpc.CallOption) (*GetSecurityDescriptorCountResponse, error)

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
	AccessCheck(context.Context, *AccessCheckRequest, ...dcerpc.CallOption) (*AccessCheckResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

type xxx_DefaultFaxobsClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultFaxobsClient) ConnectionReferenceCount(ctx context.Context, in *ConnectionReferenceCountRequest, opts ...dcerpc.CallOption) (*ConnectionReferenceCountResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConnectionReferenceCountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...dcerpc.CallOption) (*GetVersionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetVersionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) GetInstallType(ctx context.Context, in *GetInstallTypeRequest, opts ...dcerpc.CallOption) (*GetInstallTypeResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetInstallTypeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) OpenPort(ctx context.Context, in *OpenPortRequest, opts ...dcerpc.CallOption) (*OpenPortResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenPortResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) ClosePort(ctx context.Context, in *ClosePortRequest, opts ...dcerpc.CallOption) (*ClosePortResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ClosePortResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) SendDocument(ctx context.Context, in *SendDocumentRequest, opts ...dcerpc.CallOption) (*SendDocumentResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SendDocumentResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) GetQueueFileName(ctx context.Context, in *GetQueueFileNameRequest, opts ...dcerpc.CallOption) (*GetQueueFileNameResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetQueueFileNameResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) EnumJobs(ctx context.Context, in *EnumJobsRequest, opts ...dcerpc.CallOption) (*EnumJobsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumJobsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) GetJob(ctx context.Context, in *GetJobRequest, opts ...dcerpc.CallOption) (*GetJobResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetJobResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) SetJob(ctx context.Context, in *SetJobRequest, opts ...dcerpc.CallOption) (*SetJobResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetJobResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) GetPageData(ctx context.Context, in *GetPageDataRequest, opts ...dcerpc.CallOption) (*GetPageDataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetPageDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) GetDeviceStatus(ctx context.Context, in *GetDeviceStatusRequest, opts ...dcerpc.CallOption) (*GetDeviceStatusResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetDeviceStatusResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) Abort(ctx context.Context, in *AbortRequest, opts ...dcerpc.CallOption) (*AbortResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AbortResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) EnumPorts(ctx context.Context, in *EnumPortsRequest, opts ...dcerpc.CallOption) (*EnumPortsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumPortsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) GetPort(ctx context.Context, in *GetPortRequest, opts ...dcerpc.CallOption) (*GetPortResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetPortResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) SetPort(ctx context.Context, in *SetPortRequest, opts ...dcerpc.CallOption) (*SetPortResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetPortResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) EnumRoutingMethods(ctx context.Context, in *EnumRoutingMethodsRequest, opts ...dcerpc.CallOption) (*EnumRoutingMethodsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumRoutingMethodsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) EnableRoutingMethod(ctx context.Context, in *EnableRoutingMethodRequest, opts ...dcerpc.CallOption) (*EnableRoutingMethodResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnableRoutingMethodResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) GetRoutingInfo(ctx context.Context, in *GetRoutingInfoRequest, opts ...dcerpc.CallOption) (*GetRoutingInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetRoutingInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) SetRoutingInfo(ctx context.Context, in *SetRoutingInfoRequest, opts ...dcerpc.CallOption) (*SetRoutingInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetRoutingInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) EnumGlobalRoutingInfo(ctx context.Context, in *EnumGlobalRoutingInfoRequest, opts ...dcerpc.CallOption) (*EnumGlobalRoutingInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumGlobalRoutingInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) SetGlobalRoutingInfo(ctx context.Context, in *SetGlobalRoutingInfoRequest, opts ...dcerpc.CallOption) (*SetGlobalRoutingInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetGlobalRoutingInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...dcerpc.CallOption) (*GetConfigurationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetConfigurationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) SetConfiguration(ctx context.Context, in *SetConfigurationRequest, opts ...dcerpc.CallOption) (*SetConfigurationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetConfigurationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) GetLoggingCategories(ctx context.Context, in *GetLoggingCategoriesRequest, opts ...dcerpc.CallOption) (*GetLoggingCategoriesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetLoggingCategoriesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) SetLoggingCategories(ctx context.Context, in *SetLoggingCategoriesRequest, opts ...dcerpc.CallOption) (*SetLoggingCategoriesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetLoggingCategoriesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) GetTAPILocations(ctx context.Context, in *GetTAPILocationsRequest, opts ...dcerpc.CallOption) (*GetTAPILocationsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetTAPILocationsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) SetTAPILocations(ctx context.Context, in *SetTAPILocationsRequest, opts ...dcerpc.CallOption) (*SetTAPILocationsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetTAPILocationsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) GetMAPIProfiles(ctx context.Context, in *GetMAPIProfilesRequest, opts ...dcerpc.CallOption) (*GetMAPIProfilesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetMAPIProfilesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) StartClientServer(ctx context.Context, in *StartClientServerRequest, opts ...dcerpc.CallOption) (*StartClientServerResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &StartClientServerResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) GetSecurityDescriptor(ctx context.Context, in *GetSecurityDescriptorRequest, opts ...dcerpc.CallOption) (*GetSecurityDescriptorResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSecurityDescriptorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) SetSecurityDescriptor(ctx context.Context, in *SetSecurityDescriptorRequest, opts ...dcerpc.CallOption) (*SetSecurityDescriptorResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetSecurityDescriptorResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) GetSecurityDescriptorCount(ctx context.Context, in *GetSecurityDescriptorCountRequest, opts ...dcerpc.CallOption) (*GetSecurityDescriptorCountResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSecurityDescriptorCountResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) AccessCheck(ctx context.Context, in *AccessCheckRequest, opts ...dcerpc.CallOption) (*AccessCheckResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AccessCheckResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultFaxobsClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultFaxobsClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewFaxobsClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (FaxobsClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(FaxobsSyntaxV4_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultFaxobsClient{cc: cc}, nil
}

// xxx_ConnectionReferenceCountOperation structure represents the FaxObs_ConnectionRefCount operation
type xxx_ConnectionReferenceCountOperation struct {
	Handle   *fax.Service `idl:"name:Handle" json:"handle"`
	Connect  uint32       `idl:"name:Connect" json:"connect"`
	CanShare uint32       `idl:"name:CanShare" json:"can_share"`
	Return   uint32       `idl:"name:Return" json:"return"`
}

func (o *xxx_ConnectionReferenceCountOperation) OpNum() int { return 0 }

func (o *xxx_ConnectionReferenceCountOperation) OpName() string {
	return "/faxobs/v4/FaxObs_ConnectionRefCount"
}

func (o *xxx_ConnectionReferenceCountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionReferenceCountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Handle {in, out} (1:{alias=PRPC_FAX_SVC_HANDLE}*(1))(2:{context_handle, alias=RPC_FAX_SVC_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.Service{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Connect {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Connect); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionReferenceCountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Handle {in, out} (1:{alias=PRPC_FAX_SVC_HANDLE,pointer=ref}*(1))(2:{context_handle, alias=RPC_FAX_SVC_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &fax.Service{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Connect {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Connect); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionReferenceCountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionReferenceCountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Handle {in, out} (1:{alias=PRPC_FAX_SVC_HANDLE}*(1))(2:{context_handle, alias=RPC_FAX_SVC_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle != nil {
			if err := o.Handle.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.Service{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// CanShare {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.CanShare); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectionReferenceCountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Handle {in, out} (1:{alias=PRPC_FAX_SVC_HANDLE,pointer=ref}*(1))(2:{context_handle, alias=RPC_FAX_SVC_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.Handle == nil {
			o.Handle = &fax.Service{}
		}
		if err := o.Handle.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// CanShare {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.CanShare); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ConnectionReferenceCountRequest structure represents the FaxObs_ConnectionRefCount operation request
type ConnectionReferenceCountRequest struct {
	// Handle: The connection handle that references a connection between the client and
	// the server. If Connect is set to 0x00000001 (Connect), a new handle is returned in
	// this parameter. Otherwise, this parameter MUST be set to a handle returned from a
	// previous call to this method.
	Handle *fax.Service `idl:"name:Handle" json:"handle"`
	// Connect: A DWORD ([MS-DTYP] section 2.2.9) value that specifies connection information.
	//
	//	+-----------------------+----------------------------------------------------------------------------------+
	//	|                       |                                                                                  |
	//	|      VALUE/CODE       |                                     MEANING                                      |
	//	|                       |                                                                                  |
	//	+-----------------------+----------------------------------------------------------------------------------+
	//	+-----------------------+----------------------------------------------------------------------------------+
	//	| Disconnect 0x00000000 | Close the Fax Server connection. The handle specified in Handle MUST have been   |
	//	|                       | returned by a previous call to FaxObs_ConnectionRefCount with a Connect value of |
	//	|                       | 0x00000001 (Connect). After this call, the handle in Handle will be invalid and  |
	//	|                       | MUST NOT be used in any subsequent calls.                                        |
	//	+-----------------------+----------------------------------------------------------------------------------+
	//	| Connect 0x00000001    | Connect to the Fax Server.                                                       |
	//	+-----------------------+----------------------------------------------------------------------------------+
	Connect uint32 `idl:"name:Connect" json:"connect"`
}

func (o *ConnectionReferenceCountRequest) xxx_ToOp(ctx context.Context, op *xxx_ConnectionReferenceCountOperation) *xxx_ConnectionReferenceCountOperation {
	if op == nil {
		op = &xxx_ConnectionReferenceCountOperation{}
	}
	if o == nil {
		return op
	}
	o.Handle = op.Handle
	o.Connect = op.Connect
	return op
}

func (o *ConnectionReferenceCountRequest) xxx_FromOp(ctx context.Context, op *xxx_ConnectionReferenceCountOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.Connect = op.Connect
}
func (o *ConnectionReferenceCountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ConnectionReferenceCountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionReferenceCountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConnectionReferenceCountResponse structure represents the FaxObs_ConnectionRefCount operation response
type ConnectionReferenceCountResponse struct {
	// Handle: The connection handle that references a connection between the client and
	// the server. If Connect is set to 0x00000001 (Connect), a new handle is returned in
	// this parameter. Otherwise, this parameter MUST be set to a handle returned from a
	// previous call to this method.
	Handle *fax.Service `idl:"name:Handle" json:"handle"`
	// CanShare: The fax server MUST return a nonzero value in the DWORD referenced by this
	// parameter if the fax print queues can be shared as described in section 3.1.1, and
	// a zero value otherwise.<183>
	CanShare uint32 `idl:"name:CanShare" json:"can_share"`
	// Return: The FaxObs_ConnectionRefCount return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ConnectionReferenceCountResponse) xxx_ToOp(ctx context.Context, op *xxx_ConnectionReferenceCountOperation) *xxx_ConnectionReferenceCountOperation {
	if op == nil {
		op = &xxx_ConnectionReferenceCountOperation{}
	}
	if o == nil {
		return op
	}
	o.Handle = op.Handle
	o.CanShare = op.CanShare
	o.Return = op.Return
	return op
}

func (o *ConnectionReferenceCountResponse) xxx_FromOp(ctx context.Context, op *xxx_ConnectionReferenceCountOperation) {
	if o == nil {
		return
	}
	o.Handle = op.Handle
	o.CanShare = op.CanShare
	o.Return = op.Return
}
func (o *ConnectionReferenceCountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ConnectionReferenceCountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectionReferenceCountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetVersionOperation structure represents the FaxObs_GetVersion operation
type xxx_GetVersionOperation struct {
	Version uint32 `idl:"name:Version" json:"version"`
	Return  uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetVersionOperation) OpNum() int { return 1 }

func (o *xxx_GetVersionOperation) OpName() string { return "/faxobs/v4/FaxObs_GetVersion" }

func (o *xxx_GetVersionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_GetVersionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_GetVersionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Version {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Version); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Version {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Version); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetVersionRequest structure represents the FaxObs_GetVersion operation request
type GetVersionRequest struct {
}

func (o *GetVersionRequest) xxx_ToOp(ctx context.Context, op *xxx_GetVersionOperation) *xxx_GetVersionOperation {
	if op == nil {
		op = &xxx_GetVersionOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *GetVersionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetVersionOperation) {
	if o == nil {
		return
	}
}
func (o *GetVersionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetVersionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVersionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetVersionResponse structure represents the FaxObs_GetVersion operation response
type GetVersionResponse struct {
	// Version: A pointer to a DWORD ([MS-DTYP] section 2.2.9) value where on return from
	// this call, the server MUST write its version number. The server MUST write to the
	// low-order WORD ([MS-DTYP] section 2.2.61) of this DWORD value the major version number
	// and to the high-order WORD the minor version number. The returned DWORD value format
	// is as follows:
	//
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 1 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 2 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 3 | 1 |
	//	|   |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |   |   |   |   |   |   |   |   | 0 |   |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	//	| Minor version number                                          | Major version number                                          |
	//	+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
	Version uint32 `idl:"name:Version" json:"version"`
	// Return: The FaxObs_GetVersion return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetVersionResponse) xxx_ToOp(ctx context.Context, op *xxx_GetVersionOperation) *xxx_GetVersionOperation {
	if op == nil {
		op = &xxx_GetVersionOperation{}
	}
	if o == nil {
		return op
	}
	o.Version = op.Version
	o.Return = op.Return
	return op
}

func (o *GetVersionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetVersionOperation) {
	if o == nil {
		return
	}
	o.Version = op.Version
	o.Return = op.Return
}
func (o *GetVersionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetVersionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetVersionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetInstallTypeOperation structure represents the FaxObs_GetInstallType operation
type xxx_GetInstallTypeOperation struct {
	InstallType        uint32 `idl:"name:InstallType" json:"install_type"`
	InstalledPlatforms uint32 `idl:"name:InstalledPlatforms" json:"installed_platforms"`
	ProductType        uint32 `idl:"name:ProductType" json:"product_type"`
	Return             uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetInstallTypeOperation) OpNum() int { return 2 }

func (o *xxx_GetInstallTypeOperation) OpName() string { return "/faxobs/v4/FaxObs_GetInstallType" }

func (o *xxx_GetInstallTypeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInstallTypeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_GetInstallTypeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_GetInstallTypeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInstallTypeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// InstallType {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InstallType); err != nil {
			return err
		}
	}
	// InstalledPlatforms {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.InstalledPlatforms); err != nil {
			return err
		}
	}
	// ProductType {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ProductType); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInstallTypeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// InstallType {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InstallType); err != nil {
			return err
		}
	}
	// InstalledPlatforms {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.InstalledPlatforms); err != nil {
			return err
		}
	}
	// ProductType {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ProductType); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetInstallTypeRequest structure represents the FaxObs_GetInstallType operation request
type GetInstallTypeRequest struct {
}

func (o *GetInstallTypeRequest) xxx_ToOp(ctx context.Context, op *xxx_GetInstallTypeOperation) *xxx_GetInstallTypeOperation {
	if op == nil {
		op = &xxx_GetInstallTypeOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *GetInstallTypeRequest) xxx_FromOp(ctx context.Context, op *xxx_GetInstallTypeOperation) {
	if o == nil {
		return
	}
}
func (o *GetInstallTypeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetInstallTypeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInstallTypeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetInstallTypeResponse structure represents the FaxObs_GetInstallType operation response
type GetInstallTypeResponse struct {
	// InstallType: A pointer to a DWORD ([MS-DTYP] section 2.2.9) value where, upon return
	// from this call, the fax server MUST write the install type of the fax server. This
	// value MUST be 0x00000002 for FAX_INSTALL_SERVER. The values 0x00000001, 0x00000004,
	// and 0x00000008 are reserved for local use.
	InstallType uint32 `idl:"name:InstallType" json:"install_type"`
	// InstalledPlatforms: A pointer to a DWORD value where on return from this call, the
	// fax server MUST write the installed platform (microprocessor type) of the fax server.
	// This value MUST be one of the following:
	//
	//	+-----------------------------------------+-----------------------------------------------------+
	//	|                                         |                                                     |
	//	|               VALUE/CODE                |                       MEANING                       |
	//	|                                         |                                                     |
	//	+-----------------------------------------+-----------------------------------------------------+
	//	+-----------------------------------------+-----------------------------------------------------+
	//	| FAX_INSTALLED_PLATFORM_X86 0x00000001   | The fax server is running on an Intel x86 platform. |
	//	+-----------------------------------------+-----------------------------------------------------+
	//	| FAX_INSTALLED_PLATFORM_MIPS 0x00000002  | The fax server is running on a MIPS platform.       |
	//	+-----------------------------------------+-----------------------------------------------------+
	//	| FAX_INSTALLED_PLATFORM_ALPHA 0x00000004 | The fax server is running on a DEC Alpha platform.  |
	//	+-----------------------------------------+-----------------------------------------------------+
	//	| FAX_INSTALLED_PLATFORM_PPC 0x00000008   | The fax server is running on a PowerPC platform.    |
	//	+-----------------------------------------+-----------------------------------------------------+
	InstalledPlatforms uint32 `idl:"name:InstalledPlatforms" json:"installed_platforms"`
	// ProductType: A pointer to a DWORD value where on return from this call, the fax server
	// MUST write the installed product type of the fax server. This value MUST be one of
	// the following:
	//
	//	+--------------------------------+--------------------------------------------------------+
	//	|                                |                                                        |
	//	|           VALUE/CODE           |                        MEANING                         |
	//	|                                |                                                        |
	//	+--------------------------------+--------------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------------+
	//	| PRODUCT_TYPE_WINNT 0x00000001  | The fax server is a workstation-type operating system. |
	//	+--------------------------------+--------------------------------------------------------+
	//	| PRODUCT_TYPE_SERVER 0x00000002 | The fax server is a server-type operating system.      |
	//	+--------------------------------+--------------------------------------------------------+
	ProductType uint32 `idl:"name:ProductType" json:"product_type"`
	// Return: The FaxObs_GetInstallType return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetInstallTypeResponse) xxx_ToOp(ctx context.Context, op *xxx_GetInstallTypeOperation) *xxx_GetInstallTypeOperation {
	if op == nil {
		op = &xxx_GetInstallTypeOperation{}
	}
	if o == nil {
		return op
	}
	o.InstallType = op.InstallType
	o.InstalledPlatforms = op.InstalledPlatforms
	o.ProductType = op.ProductType
	o.Return = op.Return
	return op
}

func (o *GetInstallTypeResponse) xxx_FromOp(ctx context.Context, op *xxx_GetInstallTypeOperation) {
	if o == nil {
		return
	}
	o.InstallType = op.InstallType
	o.InstalledPlatforms = op.InstalledPlatforms
	o.ProductType = op.ProductType
	o.Return = op.Return
}
func (o *GetInstallTypeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetInstallTypeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInstallTypeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenPortOperation structure represents the FaxObs_OpenPort operation
type xxx_OpenPortOperation struct {
	DeviceID uint32    `idl:"name:DeviceId" json:"device_id"`
	Flags    uint32    `idl:"name:Flags" json:"flags"`
	FaxPort  *fax.Port `idl:"name:FaxPortHandle" json:"fax_port"`
	Return   uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenPortOperation) OpNum() int { return 3 }

func (o *xxx_OpenPortOperation) OpName() string { return "/faxobs/v4/FaxObs_OpenPort" }

func (o *xxx_OpenPortOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPortOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DeviceId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.DeviceID); err != nil {
			return err
		}
	}
	// Flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPortOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DeviceId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DeviceID); err != nil {
			return err
		}
	}
	// Flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPortOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPortOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// FaxPortHandle {out} (1:{alias=PRPC_FAX_PORT_HANDLE}*(1))(2:{context_handle, alias=RPC_FAX_PORT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort != nil {
			if err := o.FaxPort.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.Port{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPortOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// FaxPortHandle {out} (1:{alias=PRPC_FAX_PORT_HANDLE,pointer=ref}*(1))(2:{context_handle, alias=RPC_FAX_PORT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort == nil {
			o.FaxPort = &fax.Port{}
		}
		if err := o.FaxPort.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// OpenPortRequest structure represents the FaxObs_OpenPort operation request
type OpenPortRequest struct {
	// DeviceId: A DWORD ([MS-DTYP] section 2.2.9) variable that contains the line identifier
	// for the receiving device (port). The client SHOULD call the FaxObs_EnumPorts (section
	// 3.1.4.2.15) method to retrieve a valid value for this parameter.
	DeviceID uint32 `idl:"name:DeviceId" json:"device_id"`
	// Flags: A DWORD variable that contains a set of bit flags defining the access mode
	// for the port.<185>
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|                             |                                                                                  |
	//	|         VALUE/CODE          |                                     MEANING                                      |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000                  | No port access mode flags are specified.                                         |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PORT_OPEN_QUERY 0x00000001  | The port access mode that is required to obtain a fax port handle. This access   |
	//	|                             | level is also required to call the FaxObs_GetPort (section 3.1.4.2.16) method to |
	//	|                             | query fax port information.<186>                                                 |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| PORT_OPEN_MODIFY 0x00000002 | The port access mode that allows changes to the configuration of a fax           |
	//	|                             | port. The fax server can use this port access mode to allow execution of the     |
	//	|                             | FaxObs_SetPort (section 3.1.4.2.17) method. This access mode also includes the   |
	//	|                             | allowance that is associated with the PORT_OPEN_QUERY access mode.<187>          |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
}

func (o *OpenPortRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenPortOperation) *xxx_OpenPortOperation {
	if op == nil {
		op = &xxx_OpenPortOperation{}
	}
	if o == nil {
		return op
	}
	o.DeviceID = op.DeviceID
	o.Flags = op.Flags
	return op
}

func (o *OpenPortRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenPortOperation) {
	if o == nil {
		return
	}
	o.DeviceID = op.DeviceID
	o.Flags = op.Flags
}
func (o *OpenPortRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenPortRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenPortOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenPortResponse structure represents the FaxObs_OpenPort operation response
type OpenPortResponse struct {
	// FaxPortHandle: A pointer to a variable that receives a fax port handle (from the
	// Fax Data Types (section 2.2.74) enumeration) that is required on subsequent calls
	// by other fax client methods.
	FaxPort *fax.Port `idl:"name:FaxPortHandle" json:"fax_port"`
	// Return: The FaxObs_OpenPort return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *OpenPortResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenPortOperation) *xxx_OpenPortOperation {
	if op == nil {
		op = &xxx_OpenPortOperation{}
	}
	if o == nil {
		return op
	}
	o.FaxPort = op.FaxPort
	o.Return = op.Return
	return op
}

func (o *OpenPortResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenPortOperation) {
	if o == nil {
		return
	}
	o.FaxPort = op.FaxPort
	o.Return = op.Return
}
func (o *OpenPortResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenPortResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenPortOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ClosePortOperation structure represents the FaxObs_ClosePort operation
type xxx_ClosePortOperation struct {
	FaxPort *fax.Port `idl:"name:FaxPortHandle" json:"fax_port"`
	Return  uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_ClosePortOperation) OpNum() int { return 4 }

func (o *xxx_ClosePortOperation) OpName() string { return "/faxobs/v4/FaxObs_ClosePort" }

func (o *xxx_ClosePortOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClosePortOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FaxPortHandle {in, out} (1:{alias=PRPC_FAX_PORT_HANDLE}*(1))(2:{context_handle, alias=RPC_FAX_PORT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort != nil {
			if err := o.FaxPort.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.Port{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_ClosePortOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FaxPortHandle {in, out} (1:{alias=PRPC_FAX_PORT_HANDLE,pointer=ref}*(1))(2:{context_handle, alias=RPC_FAX_PORT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort == nil {
			o.FaxPort = &fax.Port{}
		}
		if err := o.FaxPort.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClosePortOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClosePortOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// FaxPortHandle {in, out} (1:{alias=PRPC_FAX_PORT_HANDLE}*(1))(2:{context_handle, alias=RPC_FAX_PORT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort != nil {
			if err := o.FaxPort.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.Port{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClosePortOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// FaxPortHandle {in, out} (1:{alias=PRPC_FAX_PORT_HANDLE,pointer=ref}*(1))(2:{context_handle, alias=RPC_FAX_PORT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort == nil {
			o.FaxPort = &fax.Port{}
		}
		if err := o.FaxPort.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ClosePortRequest structure represents the FaxObs_ClosePort operation request
type ClosePortRequest struct {
	// FaxPortHandle: A pointer to a fax port handle obtained with a FaxObs_OpenPort call.
	FaxPort *fax.Port `idl:"name:FaxPortHandle" json:"fax_port"`
}

func (o *ClosePortRequest) xxx_ToOp(ctx context.Context, op *xxx_ClosePortOperation) *xxx_ClosePortOperation {
	if op == nil {
		op = &xxx_ClosePortOperation{}
	}
	if o == nil {
		return op
	}
	o.FaxPort = op.FaxPort
	return op
}

func (o *ClosePortRequest) xxx_FromOp(ctx context.Context, op *xxx_ClosePortOperation) {
	if o == nil {
		return
	}
	o.FaxPort = op.FaxPort
}
func (o *ClosePortRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ClosePortRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClosePortOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ClosePortResponse structure represents the FaxObs_ClosePort operation response
type ClosePortResponse struct {
	// FaxPortHandle: A pointer to a fax port handle obtained with a FaxObs_OpenPort call.
	FaxPort *fax.Port `idl:"name:FaxPortHandle" json:"fax_port"`
	// Return: The FaxObs_ClosePort return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ClosePortResponse) xxx_ToOp(ctx context.Context, op *xxx_ClosePortOperation) *xxx_ClosePortOperation {
	if op == nil {
		op = &xxx_ClosePortOperation{}
	}
	if o == nil {
		return op
	}
	o.FaxPort = op.FaxPort
	o.Return = op.Return
	return op
}

func (o *ClosePortResponse) xxx_FromOp(ctx context.Context, op *xxx_ClosePortOperation) {
	if o == nil {
		return
	}
	o.FaxPort = op.FaxPort
	o.Return = op.Return
}
func (o *ClosePortResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ClosePortResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClosePortOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SendDocumentOperation structure represents the FaxObs_SendDocument operation
type xxx_SendDocumentOperation struct {
	FileName  string         `idl:"name:FileName;string;pointer:unique" json:"file_name"`
	JobParams *fax.JobParamW `idl:"name:JobParams" json:"job_params"`
	FaxJobID  uint32         `idl:"name:FaxJobId" json:"fax_job_id"`
	Return    uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_SendDocumentOperation) OpNum() int { return 5 }

func (o *xxx_SendDocumentOperation) OpName() string { return "/faxobs/v4/FaxObs_SendDocument" }

func (o *xxx_SendDocumentOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendDocumentOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FileName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.FileName != "" {
			_ptr_FileName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.FileName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.FileName, _ptr_FileName); err != nil {
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
	// JobParams {in} (1:{pointer=ref}*(1))(2:{alias=FAX_JOB_PARAMW}(struct))
	{
		if o.JobParams != nil {
			if err := o.JobParams.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.JobParamW{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendDocumentOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FileName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_FileName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.FileName); err != nil {
				return err
			}
			return nil
		})
		_s_FileName := func(ptr interface{}) { o.FileName = *ptr.(*string) }
		if err := w.ReadPointer(&o.FileName, _s_FileName, _ptr_FileName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// JobParams {in} (1:{pointer=ref}*(1))(2:{alias=FAX_JOB_PARAMW}(struct))
	{
		if o.JobParams == nil {
			o.JobParams = &fax.JobParamW{}
		}
		if err := o.JobParams.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendDocumentOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendDocumentOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// FaxJobId {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.FaxJobID); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SendDocumentOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// FaxJobId {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.FaxJobID); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SendDocumentRequest structure represents the FaxObs_SendDocument operation request
type SendDocumentRequest struct {
	// FileName: A null-terminated character string that contains the name of the file,
	// without path information, of the fax document in TIFF format. The server checks the
	// server queue directory for this file. Before making this call, the client can create
	// a file on the server by calling FaxObs_GetQueueFileName (section 3.1.4.2.8) and then
	// provide content for the file by using a protocol outside of this specification, such
	// as [MS-SMB].
	FileName string `idl:"name:FileName;string;pointer:unique" json:"file_name"`
	// JobParams: A pointer to a FAX_JOB_PARAMW (section 2.2.13) structure that contains
	// the information necessary for the server to send the fax transmission, including
	// information describing the personal profiles (section 3.1.1) for the sender and the
	// recipient of the fax.
	JobParams *fax.JobParamW `idl:"name:JobParams" json:"job_params"`
}

func (o *SendDocumentRequest) xxx_ToOp(ctx context.Context, op *xxx_SendDocumentOperation) *xxx_SendDocumentOperation {
	if op == nil {
		op = &xxx_SendDocumentOperation{}
	}
	if o == nil {
		return op
	}
	o.FileName = op.FileName
	o.JobParams = op.JobParams
	return op
}

func (o *SendDocumentRequest) xxx_FromOp(ctx context.Context, op *xxx_SendDocumentOperation) {
	if o == nil {
		return
	}
	o.FileName = op.FileName
	o.JobParams = op.JobParams
}
func (o *SendDocumentRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SendDocumentRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SendDocumentOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SendDocumentResponse structure represents the FaxObs_SendDocument operation response
type SendDocumentResponse struct {
	// FaxJobId: A pointer to a DWORD ([MS-DTYP] section 2.2.9) that returns the job ID.
	FaxJobID uint32 `idl:"name:FaxJobId" json:"fax_job_id"`
	// Return: The FaxObs_SendDocument return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SendDocumentResponse) xxx_ToOp(ctx context.Context, op *xxx_SendDocumentOperation) *xxx_SendDocumentOperation {
	if op == nil {
		op = &xxx_SendDocumentOperation{}
	}
	if o == nil {
		return op
	}
	o.FaxJobID = op.FaxJobID
	o.Return = op.Return
	return op
}

func (o *SendDocumentResponse) xxx_FromOp(ctx context.Context, op *xxx_SendDocumentOperation) {
	if o == nil {
		return
	}
	o.FaxJobID = op.FaxJobID
	o.Return = op.Return
}
func (o *SendDocumentResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SendDocumentResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SendDocumentOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetQueueFileNameOperation structure represents the FaxObs_GetQueueFileName operation
type xxx_GetQueueFileNameOperation struct {
	FileName     string `idl:"name:FileName;size_is:(FileNameSize);pointer:unique" json:"file_name"`
	FileNameSize uint32 `idl:"name:FileNameSize" json:"file_name_size"`
	Return       uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetQueueFileNameOperation) OpNum() int { return 6 }

func (o *xxx_GetQueueFileNameOperation) OpName() string { return "/faxobs/v4/FaxObs_GetQueueFileName" }

func (o *xxx_GetQueueFileNameOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.FileName != "" && o.FileNameSize == 0 {
		o.FileNameSize = uint32(len(o.FileName))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQueueFileNameOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FileName {in, out} (1:{pointer=unique, alias=LPWSTR}*(1)[dim:0,size_is=FileNameSize,string](wchar))
	{
		if o.FileName != "" || o.FileNameSize > 0 {
			_ptr_FileName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.FileNameSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				_FileName_buf := utf16.Encode([]rune(o.FileName))
				if uint64(len(_FileName_buf)) > sizeInfo[0] {
					_FileName_buf = _FileName_buf[:sizeInfo[0]]
				}
				for i1 := range _FileName_buf {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(_FileName_buf[i1]); err != nil {
						return err
					}
				}
				for i1 := len(_FileName_buf); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint16(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileName, _ptr_FileName); err != nil {
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
	// FileNameSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.FileNameSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQueueFileNameOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FileName {in, out} (1:{pointer=unique, alias=LPWSTR}*(1)[dim:0,size_is=FileNameSize,string](wchar))
	{
		_ptr_FileName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			var _FileName_buf []uint16
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _FileName_buf", sizeInfo[0])
			}
			_FileName_buf = make([]uint16, sizeInfo[0])
			for i1 := range _FileName_buf {
				i1 := i1
				if err := w.ReadData(&_FileName_buf[i1]); err != nil {
					return err
				}
			}
			o.FileName = strings.TrimRight(string(utf16.Decode(_FileName_buf)), ndr.ZeroString)
			return nil
		})
		_s_FileName := func(ptr interface{}) { o.FileName = *ptr.(*string) }
		if err := w.ReadPointer(&o.FileName, _s_FileName, _ptr_FileName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// FileNameSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.FileNameSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQueueFileNameOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQueueFileNameOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// FileName {in, out} (1:{pointer=unique, alias=LPWSTR}*(1)[dim:0,size_is=FileNameSize,string](wchar))
	{
		if o.FileName != "" || o.FileNameSize > 0 {
			_ptr_FileName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.FileNameSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				_FileName_buf := utf16.Encode([]rune(o.FileName))
				if uint64(len(_FileName_buf)) > sizeInfo[0] {
					_FileName_buf = _FileName_buf[:sizeInfo[0]]
				}
				for i1 := range _FileName_buf {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(_FileName_buf[i1]); err != nil {
						return err
					}
				}
				for i1 := len(_FileName_buf); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint16(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FileName, _ptr_FileName); err != nil {
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
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetQueueFileNameOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// FileName {in, out} (1:{pointer=unique, alias=LPWSTR}*(1)[dim:0,size_is=FileNameSize,string](wchar))
	{
		_ptr_FileName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			var _FileName_buf []uint16
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array _FileName_buf", sizeInfo[0])
			}
			_FileName_buf = make([]uint16, sizeInfo[0])
			for i1 := range _FileName_buf {
				i1 := i1
				if err := w.ReadData(&_FileName_buf[i1]); err != nil {
					return err
				}
			}
			o.FileName = strings.TrimRight(string(utf16.Decode(_FileName_buf)), ndr.ZeroString)
			return nil
		})
		_s_FileName := func(ptr interface{}) { o.FileName = *ptr.(*string) }
		if err := w.ReadPointer(&o.FileName, _s_FileName, _ptr_FileName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetQueueFileNameRequest structure represents the FaxObs_GetQueueFileName operation request
type GetQueueFileNameRequest struct {
	// FileName: A buffer that MUST be allocated by the client to hold FileNameSize characters.
	// On successful return from this call the server MUST write to this buffer a null-terminated
	// character string containing the path name, including file name and extension, for
	// a new unique file name within the fax server queue directory.<192>
	FileName string `idl:"name:FileName;size_is:(FileNameSize);pointer:unique" json:"file_name"`
	// FileNameSize: A DWORD ([MS-DTYP] section 2.2.9) value that specifies the size, in
	// characters, of the FileName buffer.<193>
	FileNameSize uint32 `idl:"name:FileNameSize" json:"file_name_size"`
}

func (o *GetQueueFileNameRequest) xxx_ToOp(ctx context.Context, op *xxx_GetQueueFileNameOperation) *xxx_GetQueueFileNameOperation {
	if op == nil {
		op = &xxx_GetQueueFileNameOperation{}
	}
	if o == nil {
		return op
	}
	o.FileName = op.FileName
	o.FileNameSize = op.FileNameSize
	return op
}

func (o *GetQueueFileNameRequest) xxx_FromOp(ctx context.Context, op *xxx_GetQueueFileNameOperation) {
	if o == nil {
		return
	}
	o.FileName = op.FileName
	o.FileNameSize = op.FileNameSize
}
func (o *GetQueueFileNameRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetQueueFileNameRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQueueFileNameOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetQueueFileNameResponse structure represents the FaxObs_GetQueueFileName operation response
type GetQueueFileNameResponse struct {
	// FileName: A buffer that MUST be allocated by the client to hold FileNameSize characters.
	// On successful return from this call the server MUST write to this buffer a null-terminated
	// character string containing the path name, including file name and extension, for
	// a new unique file name within the fax server queue directory.<192>
	FileName string `idl:"name:FileName;size_is:(FileNameSize);pointer:unique" json:"file_name"`
	// Return: The FaxObs_GetQueueFileName return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetQueueFileNameResponse) xxx_ToOp(ctx context.Context, op *xxx_GetQueueFileNameOperation) *xxx_GetQueueFileNameOperation {
	if op == nil {
		op = &xxx_GetQueueFileNameOperation{}
	}
	if o == nil {
		return op
	}
	o.FileName = op.FileName
	o.Return = op.Return
	return op
}

func (o *GetQueueFileNameResponse) xxx_FromOp(ctx context.Context, op *xxx_GetQueueFileNameOperation) {
	if o == nil {
		return
	}
	o.FileName = op.FileName
	o.Return = op.Return
}
func (o *GetQueueFileNameResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetQueueFileNameResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetQueueFileNameOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumJobsOperation structure represents the FaxObs_EnumJobs operation
type xxx_EnumJobsOperation struct {
	Buffer       []byte `idl:"name:Buffer;size_is:(, BufferSize);pointer:unique" json:"buffer"`
	BufferSize   uint32 `idl:"name:BufferSize" json:"buffer_size"`
	JobsReturned uint32 `idl:"name:JobsReturned" json:"jobs_returned"`
	Return       uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumJobsOperation) OpNum() int { return 7 }

func (o *xxx_EnumJobsOperation) OpName() string { return "/faxobs/v4/FaxObs_EnumJobs" }

func (o *xxx_EnumJobsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Buffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumJobsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Buffer != nil || o.BufferSize > 0 {
					_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.Buffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.Buffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumJobsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
				}
				o.Buffer = make([]byte, sizeInfo[0])
				for i1 := range o.Buffer {
					i1 := i1
					if err := w.ReadData(&o.Buffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumJobsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Buffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumJobsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Buffer != nil || o.BufferSize > 0 {
					_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.Buffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.Buffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	// JobsReturned {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.JobsReturned); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumJobsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
				}
				o.Buffer = make([]byte, sizeInfo[0])
				for i1 := range o.Buffer {
					i1 := i1
					if err := w.ReadData(&o.Buffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	// JobsReturned {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.JobsReturned); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumJobsRequest structure represents the FaxObs_EnumJobs operation request
type EnumJobsRequest struct {
	// Buffer: A pointer to the address of the returned buffer containing an array of _FAX_JOB_ENTRY
	// (section 2.2.6) structures.
	Buffer []byte `idl:"name:Buffer;size_is:(, BufferSize);pointer:unique" json:"buffer"`
	// BufferSize: A variable to return the size, in bytes, of the job information buffer.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
}

func (o *EnumJobsRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumJobsOperation) *xxx_EnumJobsOperation {
	if op == nil {
		op = &xxx_EnumJobsOperation{}
	}
	if o == nil {
		return op
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	return op
}

func (o *EnumJobsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumJobsOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
}
func (o *EnumJobsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumJobsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumJobsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumJobsResponse structure represents the FaxObs_EnumJobs operation response
type EnumJobsResponse struct {
	// Buffer: A pointer to the address of the returned buffer containing an array of _FAX_JOB_ENTRY
	// (section 2.2.6) structures.
	Buffer []byte `idl:"name:Buffer;size_is:(, BufferSize);pointer:unique" json:"buffer"`
	// BufferSize: A variable to return the size, in bytes, of the job information buffer.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
	// JobsReturned: A pointer to a DWORD ([MS-DTYP] section 2.2.9) variable to receive
	// the number of _FAX_JOB_ENTRY structures that the method returns in Buffer.
	JobsReturned uint32 `idl:"name:JobsReturned" json:"jobs_returned"`
	// Return: The FaxObs_EnumJobs return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EnumJobsResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumJobsOperation) *xxx_EnumJobsOperation {
	if op == nil {
		op = &xxx_EnumJobsOperation{}
	}
	if o == nil {
		return op
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	o.JobsReturned = op.JobsReturned
	o.Return = op.Return
	return op
}

func (o *EnumJobsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumJobsOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	o.JobsReturned = op.JobsReturned
	o.Return = op.Return
}
func (o *EnumJobsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumJobsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumJobsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetJobOperation structure represents the FaxObs_GetJob operation
type xxx_GetJobOperation struct {
	JobID      uint32 `idl:"name:JobId" json:"job_id"`
	Buffer     []byte `idl:"name:Buffer;size_is:(, BufferSize);pointer:unique" json:"buffer"`
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetJobOperation) OpNum() int { return 8 }

func (o *xxx_GetJobOperation) OpName() string { return "/faxobs/v4/FaxObs_GetJob" }

func (o *xxx_GetJobOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Buffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJobOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// JobId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.JobID); err != nil {
			return err
		}
	}
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Buffer != nil || o.BufferSize > 0 {
					_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.Buffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.Buffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJobOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// JobId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.JobID); err != nil {
			return err
		}
	}
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
				}
				o.Buffer = make([]byte, sizeInfo[0])
				for i1 := range o.Buffer {
					i1 := i1
					if err := w.ReadData(&o.Buffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJobOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Buffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJobOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Buffer != nil || o.BufferSize > 0 {
					_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.Buffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.Buffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetJobOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
				}
				o.Buffer = make([]byte, sizeInfo[0])
				for i1 := range o.Buffer {
					i1 := i1
					if err := w.ReadData(&o.Buffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetJobRequest structure represents the FaxObs_GetJob operation request
type GetJobRequest struct {
	// JobId: A number that uniquely identifies a queued or active fax job on the server.
	JobID uint32 `idl:"name:JobId" json:"job_id"`
	// Buffer: A pointer to the address of the returned buffer containing a _FAX_JOB_ENTRY
	// (section 2.2.6) structure.
	Buffer []byte `idl:"name:Buffer;size_is:(, BufferSize);pointer:unique" json:"buffer"`
	// BufferSize: A variable to return the size, in bytes, of the job information buffer.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
}

func (o *GetJobRequest) xxx_ToOp(ctx context.Context, op *xxx_GetJobOperation) *xxx_GetJobOperation {
	if op == nil {
		op = &xxx_GetJobOperation{}
	}
	if o == nil {
		return op
	}
	o.JobID = op.JobID
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	return op
}

func (o *GetJobRequest) xxx_FromOp(ctx context.Context, op *xxx_GetJobOperation) {
	if o == nil {
		return
	}
	o.JobID = op.JobID
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
}
func (o *GetJobRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetJobRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetJobOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetJobResponse structure represents the FaxObs_GetJob operation response
type GetJobResponse struct {
	// Buffer: A pointer to the address of the returned buffer containing a _FAX_JOB_ENTRY
	// (section 2.2.6) structure.
	Buffer []byte `idl:"name:Buffer;size_is:(, BufferSize);pointer:unique" json:"buffer"`
	// BufferSize: A variable to return the size, in bytes, of the job information buffer.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
	// Return: The FaxObs_GetJob return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetJobResponse) xxx_ToOp(ctx context.Context, op *xxx_GetJobOperation) *xxx_GetJobOperation {
	if op == nil {
		op = &xxx_GetJobOperation{}
	}
	if o == nil {
		return op
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	o.Return = op.Return
	return op
}

func (o *GetJobResponse) xxx_FromOp(ctx context.Context, op *xxx_GetJobOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	o.Return = op.Return
}
func (o *GetJobResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetJobResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetJobOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetJobOperation structure represents the FaxObs_SetJob operation
type xxx_SetJobOperation struct {
	JobID    uint32        `idl:"name:JobId" json:"job_id"`
	Command  uint32        `idl:"name:Command" json:"command"`
	JobEntry *fax.JobEntry `idl:"name:JobEntry" json:"job_entry"`
	Return   uint32        `idl:"name:Return" json:"return"`
}

func (o *xxx_SetJobOperation) OpNum() int { return 9 }

func (o *xxx_SetJobOperation) OpName() string { return "/faxobs/v4/FaxObs_SetJob" }

func (o *xxx_SetJobOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetJobOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// JobId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.JobID); err != nil {
			return err
		}
	}
	// Command {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Command); err != nil {
			return err
		}
	}
	// JobEntry {in} (1:{pointer=ref}*(1))(2:{alias=FAX_JOB_ENTRY}(struct))
	{
		if o.JobEntry != nil {
			if err := o.JobEntry.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.JobEntry{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetJobOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// JobId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.JobID); err != nil {
			return err
		}
	}
	// Command {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Command); err != nil {
			return err
		}
	}
	// JobEntry {in} (1:{pointer=ref}*(1))(2:{alias=FAX_JOB_ENTRY}(struct))
	{
		if o.JobEntry == nil {
			o.JobEntry = &fax.JobEntry{}
		}
		if err := o.JobEntry.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetJobOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetJobOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetJobOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetJobRequest structure represents the FaxObs_SetJob operation request
type SetJobRequest struct {
	// JobId: A DWORD ([MS-DTYP] section 2.2.9) containing a value that uniquely identifies
	// the fax job to modify.
	JobID uint32 `idl:"name:JobId" json:"job_id"`
	// Command: A DWORD containing a job command that the fax server is requested to perform.
	//
	//	+----------------------+----------------------------------------------------------------------------------+
	//	|                      |                                                                                  |
	//	|      VALUE/CODE      |                                     MEANING                                      |
	//	|                      |                                                                                  |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| JC_DELETE 0x00000001 | The fax server MUST cancel the specified fax job. This job can be in an active   |
	//	|                      | or queued state. This is equivalent to calling the FaxObs_Abort (section         |
	//	|                      | 3.1.4.2.14) method.                                                              |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| JC_PAUSE 0x00000002  | The fax server MUST pause the specified fax job if the job's status is           |
	//	|                      | JS_PENDING or JS_RETRYING.                                                       |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| JC_RESUME 0x00000003 | The fax server MUST resume the specified fax job if it is in a paused state and  |
	//	|                      | return the job status to the value it had when the job was paused: JS_PENDING or |
	//	|                      | JS_RETRYING.                                                                     |
	//	+----------------------+----------------------------------------------------------------------------------+
	Command uint32 `idl:"name:Command" json:"command"`
	// JobEntry: A pointer to a FAX_JOB_ENTRY (section 2.2.5) structure. The fax server
	// MUST validate this pointer to be not NULL and fail to return ERROR_INVALID_PARAMETER
	// otherwise. Except for this validation requirement, the fax server SHOULD ignore this
	// parameter. The fax client SHOULD submit the pointer to a valid FAX_JOB_ENTRY. This
	// structure contains data obtained from FaxObs_GetJob (section 3.1.4.2.10) or FaxObs_EnumJobs
	// (section 3.1.4.2.9). This data represents the job identified by the JobId parameter.<196>
	JobEntry *fax.JobEntry `idl:"name:JobEntry" json:"job_entry"`
}

func (o *SetJobRequest) xxx_ToOp(ctx context.Context, op *xxx_SetJobOperation) *xxx_SetJobOperation {
	if op == nil {
		op = &xxx_SetJobOperation{}
	}
	if o == nil {
		return op
	}
	o.JobID = op.JobID
	o.Command = op.Command
	o.JobEntry = op.JobEntry
	return op
}

func (o *SetJobRequest) xxx_FromOp(ctx context.Context, op *xxx_SetJobOperation) {
	if o == nil {
		return
	}
	o.JobID = op.JobID
	o.Command = op.Command
	o.JobEntry = op.JobEntry
}
func (o *SetJobRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetJobRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetJobOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetJobResponse structure represents the FaxObs_SetJob operation response
type SetJobResponse struct {
	// Return: The FaxObs_SetJob return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetJobResponse) xxx_ToOp(ctx context.Context, op *xxx_SetJobOperation) *xxx_SetJobOperation {
	if op == nil {
		op = &xxx_SetJobOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *SetJobResponse) xxx_FromOp(ctx context.Context, op *xxx_SetJobOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetJobResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetJobResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetJobOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPageDataOperation structure represents the FaxObs_GetPageData operation
type xxx_GetPageDataOperation struct {
	JobID       uint32 `idl:"name:JobId" json:"job_id"`
	Buffer      []byte `idl:"name:Buffer;size_is:(, BufferSize);pointer:unique" json:"buffer"`
	BufferSize  uint32 `idl:"name:BufferSize" json:"buffer_size"`
	ImageWidth  uint32 `idl:"name:ImageWidth" json:"image_width"`
	ImageHeight uint32 `idl:"name:ImageHeight" json:"image_height"`
	Return      uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPageDataOperation) OpNum() int { return 10 }

func (o *xxx_GetPageDataOperation) OpName() string { return "/faxobs/v4/FaxObs_GetPageData" }

func (o *xxx_GetPageDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Buffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPageDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// JobId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.JobID); err != nil {
			return err
		}
	}
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Buffer != nil || o.BufferSize > 0 {
					_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.Buffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.Buffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	// ImageWidth {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ImageWidth); err != nil {
			return err
		}
	}
	// ImageHeight {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ImageHeight); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPageDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// JobId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.JobID); err != nil {
			return err
		}
	}
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
				}
				o.Buffer = make([]byte, sizeInfo[0])
				for i1 := range o.Buffer {
					i1 := i1
					if err := w.ReadData(&o.Buffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	// ImageWidth {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ImageWidth); err != nil {
			return err
		}
	}
	// ImageHeight {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ImageHeight); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPageDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Buffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPageDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Buffer != nil || o.BufferSize > 0 {
					_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.Buffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.Buffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	// ImageWidth {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ImageWidth); err != nil {
			return err
		}
	}
	// ImageHeight {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ImageHeight); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPageDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
				}
				o.Buffer = make([]byte, sizeInfo[0])
				for i1 := range o.Buffer {
					i1 := i1
					if err := w.ReadData(&o.Buffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	// ImageWidth {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ImageWidth); err != nil {
			return err
		}
	}
	// ImageHeight {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ImageHeight); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetPageDataRequest structure represents the FaxObs_GetPageData operation request
type GetPageDataRequest struct {
	// JobId: A DWORD ([MS-DTYP] section 2.2.9) containing the unique number identifying
	// the fax job that is associated with the page of data.
	JobID uint32 `idl:"name:JobId" json:"job_id"`
	// Buffer: A pointer to the address of the returned buffer containing the first page
	// of data in the fax document.
	Buffer []byte `idl:"name:Buffer;size_is:(, BufferSize);pointer:unique" json:"buffer"`
	// BufferSize: A pointer to a DWORD variable to receive the size, in bytes, of the buffer
	// pointed to by the Buffer parameter.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
	// ImageWidth: A pointer to a DWORD variable to receive the width, in pixels, of the
	// fax image.
	ImageWidth uint32 `idl:"name:ImageWidth" json:"image_width"`
	// ImageHeight: A pointer to a DWORD variable to receive the height, in pixels, of the
	// fax image.
	ImageHeight uint32 `idl:"name:ImageHeight" json:"image_height"`
}

func (o *GetPageDataRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPageDataOperation) *xxx_GetPageDataOperation {
	if op == nil {
		op = &xxx_GetPageDataOperation{}
	}
	if o == nil {
		return op
	}
	o.JobID = op.JobID
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	o.ImageWidth = op.ImageWidth
	o.ImageHeight = op.ImageHeight
	return op
}

func (o *GetPageDataRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPageDataOperation) {
	if o == nil {
		return
	}
	o.JobID = op.JobID
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	o.ImageWidth = op.ImageWidth
	o.ImageHeight = op.ImageHeight
}
func (o *GetPageDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPageDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPageDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPageDataResponse structure represents the FaxObs_GetPageData operation response
type GetPageDataResponse struct {
	// Buffer: A pointer to the address of the returned buffer containing the first page
	// of data in the fax document.
	Buffer []byte `idl:"name:Buffer;size_is:(, BufferSize);pointer:unique" json:"buffer"`
	// BufferSize: A pointer to a DWORD variable to receive the size, in bytes, of the buffer
	// pointed to by the Buffer parameter.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
	// ImageWidth: A pointer to a DWORD variable to receive the width, in pixels, of the
	// fax image.
	ImageWidth uint32 `idl:"name:ImageWidth" json:"image_width"`
	// ImageHeight: A pointer to a DWORD variable to receive the height, in pixels, of the
	// fax image.
	ImageHeight uint32 `idl:"name:ImageHeight" json:"image_height"`
	// Return: The FaxObs_GetPageData return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetPageDataResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPageDataOperation) *xxx_GetPageDataOperation {
	if op == nil {
		op = &xxx_GetPageDataOperation{}
	}
	if o == nil {
		return op
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	o.ImageWidth = op.ImageWidth
	o.ImageHeight = op.ImageHeight
	o.Return = op.Return
	return op
}

func (o *GetPageDataResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPageDataOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	o.ImageWidth = op.ImageWidth
	o.ImageHeight = op.ImageHeight
	o.Return = op.Return
}
func (o *GetPageDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPageDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPageDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDeviceStatusOperation structure represents the FaxObs_GetDeviceStatus operation
type xxx_GetDeviceStatusOperation struct {
	FaxPort      *fax.Port `idl:"name:FaxPortHandle" json:"fax_port"`
	StatusBuffer []byte    `idl:"name:StatusBuffer;size_is:(, BufferSize);pointer:unique" json:"status_buffer"`
	BufferSize   uint32    `idl:"name:BufferSize" json:"buffer_size"`
	Return       uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDeviceStatusOperation) OpNum() int { return 11 }

func (o *xxx_GetDeviceStatusOperation) OpName() string { return "/faxobs/v4/FaxObs_GetDeviceStatus" }

func (o *xxx_GetDeviceStatusOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.StatusBuffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.StatusBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDeviceStatusOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FaxPortHandle {in} (1:{context_handle, alias=RPC_FAX_PORT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort != nil {
			if err := o.FaxPort.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.Port{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// StatusBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.StatusBuffer != nil {
			_ptr_StatusBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.StatusBuffer != nil || o.BufferSize > 0 {
					_ptr_StatusBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.StatusBuffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.StatusBuffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.StatusBuffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.StatusBuffer, _ptr_StatusBuffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.StatusBuffer, _ptr_StatusBuffer); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDeviceStatusOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FaxPortHandle {in} (1:{context_handle, alias=RPC_FAX_PORT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort == nil {
			o.FaxPort = &fax.Port{}
		}
		if err := o.FaxPort.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// StatusBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_StatusBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_StatusBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.StatusBuffer", sizeInfo[0])
				}
				o.StatusBuffer = make([]byte, sizeInfo[0])
				for i1 := range o.StatusBuffer {
					i1 := i1
					if err := w.ReadData(&o.StatusBuffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_StatusBuffer := func(ptr interface{}) { o.StatusBuffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.StatusBuffer, _s_StatusBuffer, _ptr_StatusBuffer); err != nil {
				return err
			}
			return nil
		})
		_s_StatusBuffer := func(ptr interface{}) { o.StatusBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.StatusBuffer, _s_StatusBuffer, _ptr_StatusBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDeviceStatusOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.StatusBuffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.StatusBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDeviceStatusOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// StatusBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.StatusBuffer != nil {
			_ptr_StatusBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.StatusBuffer != nil || o.BufferSize > 0 {
					_ptr_StatusBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.StatusBuffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.StatusBuffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.StatusBuffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.StatusBuffer, _ptr_StatusBuffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.StatusBuffer, _ptr_StatusBuffer); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDeviceStatusOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// StatusBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_StatusBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_StatusBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.StatusBuffer", sizeInfo[0])
				}
				o.StatusBuffer = make([]byte, sizeInfo[0])
				for i1 := range o.StatusBuffer {
					i1 := i1
					if err := w.ReadData(&o.StatusBuffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_StatusBuffer := func(ptr interface{}) { o.StatusBuffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.StatusBuffer, _s_StatusBuffer, _ptr_StatusBuffer); err != nil {
				return err
			}
			return nil
		})
		_s_StatusBuffer := func(ptr interface{}) { o.StatusBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.StatusBuffer, _s_StatusBuffer, _ptr_StatusBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetDeviceStatusRequest structure represents the FaxObs_GetDeviceStatus operation request
type GetDeviceStatusRequest struct {
	// FaxPortHandle: An RPC context handle that references a specified fax port. This context
	// handle MUST be obtained using the FaxObs_OpenPort (section 3.1.4.2.5) method.
	FaxPort *fax.Port `idl:"name:FaxPortHandle" json:"fax_port"`
	// StatusBuffer: A pointer to the address of the returned buffer containing a FAX_DEVICE_STATUS
	// (section 2.2.10) structure. The structure describes the status of one fax device.
	StatusBuffer []byte `idl:"name:StatusBuffer;size_is:(, BufferSize);pointer:unique" json:"status_buffer"`
	// BufferSize: A variable to return the size, in bytes, of the status buffer.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
}

func (o *GetDeviceStatusRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDeviceStatusOperation) *xxx_GetDeviceStatusOperation {
	if op == nil {
		op = &xxx_GetDeviceStatusOperation{}
	}
	if o == nil {
		return op
	}
	o.FaxPort = op.FaxPort
	o.StatusBuffer = op.StatusBuffer
	o.BufferSize = op.BufferSize
	return op
}

func (o *GetDeviceStatusRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDeviceStatusOperation) {
	if o == nil {
		return
	}
	o.FaxPort = op.FaxPort
	o.StatusBuffer = op.StatusBuffer
	o.BufferSize = op.BufferSize
}
func (o *GetDeviceStatusRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDeviceStatusRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDeviceStatusOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDeviceStatusResponse structure represents the FaxObs_GetDeviceStatus operation response
type GetDeviceStatusResponse struct {
	// StatusBuffer: A pointer to the address of the returned buffer containing a FAX_DEVICE_STATUS
	// (section 2.2.10) structure. The structure describes the status of one fax device.
	StatusBuffer []byte `idl:"name:StatusBuffer;size_is:(, BufferSize);pointer:unique" json:"status_buffer"`
	// BufferSize: A variable to return the size, in bytes, of the status buffer.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
	// Return: The FaxObs_GetDeviceStatus return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetDeviceStatusResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDeviceStatusOperation) *xxx_GetDeviceStatusOperation {
	if op == nil {
		op = &xxx_GetDeviceStatusOperation{}
	}
	if o == nil {
		return op
	}
	o.StatusBuffer = op.StatusBuffer
	o.BufferSize = op.BufferSize
	o.Return = op.Return
	return op
}

func (o *GetDeviceStatusResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDeviceStatusOperation) {
	if o == nil {
		return
	}
	o.StatusBuffer = op.StatusBuffer
	o.BufferSize = op.BufferSize
	o.Return = op.Return
}
func (o *GetDeviceStatusResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDeviceStatusResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDeviceStatusOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AbortOperation structure represents the FaxObs_Abort operation
type xxx_AbortOperation struct {
	JobID  uint32 `idl:"name:JobId" json:"job_id"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_AbortOperation) OpNum() int { return 12 }

func (o *xxx_AbortOperation) OpName() string { return "/faxobs/v4/FaxObs_Abort" }

func (o *xxx_AbortOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// JobId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.JobID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// JobId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.JobID); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AbortOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AbortRequest structure represents the FaxObs_Abort operation request
type AbortRequest struct {
	// JobId: A DWORD ([MS-DTYP] section 2.2.9) containing a unique number that identifies
	// the fax job to terminate.
	JobID uint32 `idl:"name:JobId" json:"job_id"`
}

func (o *AbortRequest) xxx_ToOp(ctx context.Context, op *xxx_AbortOperation) *xxx_AbortOperation {
	if op == nil {
		op = &xxx_AbortOperation{}
	}
	if o == nil {
		return op
	}
	o.JobID = op.JobID
	return op
}

func (o *AbortRequest) xxx_FromOp(ctx context.Context, op *xxx_AbortOperation) {
	if o == nil {
		return
	}
	o.JobID = op.JobID
}
func (o *AbortRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AbortRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AbortOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AbortResponse structure represents the FaxObs_Abort operation response
type AbortResponse struct {
	// Return: The FaxObs_Abort return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AbortResponse) xxx_ToOp(ctx context.Context, op *xxx_AbortOperation) *xxx_AbortOperation {
	if op == nil {
		op = &xxx_AbortOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *AbortResponse) xxx_FromOp(ctx context.Context, op *xxx_AbortOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AbortResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AbortResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AbortOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumPortsOperation structure represents the FaxObs_EnumPorts operation
type xxx_EnumPortsOperation struct {
	PortBuffer    []byte `idl:"name:PortBuffer;size_is:(, BufferSize);pointer:unique" json:"port_buffer"`
	BufferSize    uint32 `idl:"name:BufferSize" json:"buffer_size"`
	PortsReturned uint32 `idl:"name:PortsReturned" json:"ports_returned"`
	Return        uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumPortsOperation) OpNum() int { return 13 }

func (o *xxx_EnumPortsOperation) OpName() string { return "/faxobs/v4/FaxObs_EnumPorts" }

func (o *xxx_EnumPortsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.PortBuffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.PortBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumPortsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// PortBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.PortBuffer != nil {
			_ptr_PortBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PortBuffer != nil || o.BufferSize > 0 {
					_ptr_PortBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.PortBuffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.PortBuffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.PortBuffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.PortBuffer, _ptr_PortBuffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PortBuffer, _ptr_PortBuffer); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumPortsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// PortBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_PortBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_PortBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.PortBuffer", sizeInfo[0])
				}
				o.PortBuffer = make([]byte, sizeInfo[0])
				for i1 := range o.PortBuffer {
					i1 := i1
					if err := w.ReadData(&o.PortBuffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_PortBuffer := func(ptr interface{}) { o.PortBuffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.PortBuffer, _s_PortBuffer, _ptr_PortBuffer); err != nil {
				return err
			}
			return nil
		})
		_s_PortBuffer := func(ptr interface{}) { o.PortBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.PortBuffer, _s_PortBuffer, _ptr_PortBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumPortsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.PortBuffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.PortBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumPortsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// PortBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.PortBuffer != nil {
			_ptr_PortBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PortBuffer != nil || o.BufferSize > 0 {
					_ptr_PortBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.PortBuffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.PortBuffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.PortBuffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.PortBuffer, _ptr_PortBuffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PortBuffer, _ptr_PortBuffer); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	// PortsReturned {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PortsReturned); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumPortsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// PortBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_PortBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_PortBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.PortBuffer", sizeInfo[0])
				}
				o.PortBuffer = make([]byte, sizeInfo[0])
				for i1 := range o.PortBuffer {
					i1 := i1
					if err := w.ReadData(&o.PortBuffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_PortBuffer := func(ptr interface{}) { o.PortBuffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.PortBuffer, _s_PortBuffer, _ptr_PortBuffer); err != nil {
				return err
			}
			return nil
		})
		_s_PortBuffer := func(ptr interface{}) { o.PortBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.PortBuffer, _s_PortBuffer, _ptr_PortBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	// PortsReturned {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PortsReturned); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumPortsRequest structure represents the FaxObs_EnumPorts operation request
type EnumPortsRequest struct {
	// PortBuffer: A pointer to the address of the returned buffer containing an array of
	// _FAX_PORT_INFO (section 2.2.8) structures. Each structure describes one fax port
	// (device).
	PortBuffer []byte `idl:"name:PortBuffer;size_is:(, BufferSize);pointer:unique" json:"port_buffer"`
	// BufferSize: A variable to return the size, in bytes, of the PortBuffer.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
}

func (o *EnumPortsRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumPortsOperation) *xxx_EnumPortsOperation {
	if op == nil {
		op = &xxx_EnumPortsOperation{}
	}
	if o == nil {
		return op
	}
	o.PortBuffer = op.PortBuffer
	o.BufferSize = op.BufferSize
	return op
}

func (o *EnumPortsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumPortsOperation) {
	if o == nil {
		return
	}
	o.PortBuffer = op.PortBuffer
	o.BufferSize = op.BufferSize
}
func (o *EnumPortsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumPortsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumPortsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumPortsResponse structure represents the FaxObs_EnumPorts operation response
type EnumPortsResponse struct {
	// PortBuffer: A pointer to the address of the returned buffer containing an array of
	// _FAX_PORT_INFO (section 2.2.8) structures. Each structure describes one fax port
	// (device).
	PortBuffer []byte `idl:"name:PortBuffer;size_is:(, BufferSize);pointer:unique" json:"port_buffer"`
	// BufferSize: A variable to return the size, in bytes, of the PortBuffer.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
	// PortsReturned: A pointer to a DWORD ([MS-DTYP] section 2.2.9) variable to receive
	// the number of _FAX_PORT_INFO that the method returns in the PortBuffer parameter.
	PortsReturned uint32 `idl:"name:PortsReturned" json:"ports_returned"`
	// Return: The FaxObs_EnumPorts return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EnumPortsResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumPortsOperation) *xxx_EnumPortsOperation {
	if op == nil {
		op = &xxx_EnumPortsOperation{}
	}
	if o == nil {
		return op
	}
	o.PortBuffer = op.PortBuffer
	o.BufferSize = op.BufferSize
	o.PortsReturned = op.PortsReturned
	o.Return = op.Return
	return op
}

func (o *EnumPortsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumPortsOperation) {
	if o == nil {
		return
	}
	o.PortBuffer = op.PortBuffer
	o.BufferSize = op.BufferSize
	o.PortsReturned = op.PortsReturned
	o.Return = op.Return
}
func (o *EnumPortsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumPortsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumPortsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetPortOperation structure represents the FaxObs_GetPort operation
type xxx_GetPortOperation struct {
	FaxPort    *fax.Port `idl:"name:FaxPortHandle" json:"fax_port"`
	PortBuffer []byte    `idl:"name:PortBuffer;size_is:(, BufferSize);pointer:unique" json:"port_buffer"`
	BufferSize uint32    `idl:"name:BufferSize" json:"buffer_size"`
	Return     uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_GetPortOperation) OpNum() int { return 14 }

func (o *xxx_GetPortOperation) OpName() string { return "/faxobs/v4/FaxObs_GetPort" }

func (o *xxx_GetPortOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.PortBuffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.PortBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPortOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FaxPortHandle {in} (1:{context_handle, alias=RPC_FAX_PORT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort != nil {
			if err := o.FaxPort.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.Port{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// PortBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.PortBuffer != nil {
			_ptr_PortBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PortBuffer != nil || o.BufferSize > 0 {
					_ptr_PortBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.PortBuffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.PortBuffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.PortBuffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.PortBuffer, _ptr_PortBuffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PortBuffer, _ptr_PortBuffer); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPortOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FaxPortHandle {in} (1:{context_handle, alias=RPC_FAX_PORT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort == nil {
			o.FaxPort = &fax.Port{}
		}
		if err := o.FaxPort.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// PortBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_PortBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_PortBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.PortBuffer", sizeInfo[0])
				}
				o.PortBuffer = make([]byte, sizeInfo[0])
				for i1 := range o.PortBuffer {
					i1 := i1
					if err := w.ReadData(&o.PortBuffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_PortBuffer := func(ptr interface{}) { o.PortBuffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.PortBuffer, _s_PortBuffer, _ptr_PortBuffer); err != nil {
				return err
			}
			return nil
		})
		_s_PortBuffer := func(ptr interface{}) { o.PortBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.PortBuffer, _s_PortBuffer, _ptr_PortBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPortOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.PortBuffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.PortBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPortOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// PortBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.PortBuffer != nil {
			_ptr_PortBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.PortBuffer != nil || o.BufferSize > 0 {
					_ptr_PortBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.PortBuffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.PortBuffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.PortBuffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.PortBuffer, _ptr_PortBuffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.PortBuffer, _ptr_PortBuffer); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetPortOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// PortBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_PortBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_PortBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.PortBuffer", sizeInfo[0])
				}
				o.PortBuffer = make([]byte, sizeInfo[0])
				for i1 := range o.PortBuffer {
					i1 := i1
					if err := w.ReadData(&o.PortBuffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_PortBuffer := func(ptr interface{}) { o.PortBuffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.PortBuffer, _s_PortBuffer, _ptr_PortBuffer); err != nil {
				return err
			}
			return nil
		})
		_s_PortBuffer := func(ptr interface{}) { o.PortBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.PortBuffer, _s_PortBuffer, _ptr_PortBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetPortRequest structure represents the FaxObs_GetPort operation request
type GetPortRequest struct {
	// FaxPortHandle: An RPC context handle that references a fax port.
	FaxPort *fax.Port `idl:"name:FaxPortHandle" json:"fax_port"`
	// PortBuffer: A pointer to the address of the returned buffer containing a _FAX_PORT_INFO
	// (section 2.2.8) structure. The structure describes one fax port (device).
	PortBuffer []byte `idl:"name:PortBuffer;size_is:(, BufferSize);pointer:unique" json:"port_buffer"`
	// BufferSize: A variable to return the size, in bytes, of the port buffer.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
}

func (o *GetPortRequest) xxx_ToOp(ctx context.Context, op *xxx_GetPortOperation) *xxx_GetPortOperation {
	if op == nil {
		op = &xxx_GetPortOperation{}
	}
	if o == nil {
		return op
	}
	o.FaxPort = op.FaxPort
	o.PortBuffer = op.PortBuffer
	o.BufferSize = op.BufferSize
	return op
}

func (o *GetPortRequest) xxx_FromOp(ctx context.Context, op *xxx_GetPortOperation) {
	if o == nil {
		return
	}
	o.FaxPort = op.FaxPort
	o.PortBuffer = op.PortBuffer
	o.BufferSize = op.BufferSize
}
func (o *GetPortRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetPortRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPortOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetPortResponse structure represents the FaxObs_GetPort operation response
type GetPortResponse struct {
	// PortBuffer: A pointer to the address of the returned buffer containing a _FAX_PORT_INFO
	// (section 2.2.8) structure. The structure describes one fax port (device).
	PortBuffer []byte `idl:"name:PortBuffer;size_is:(, BufferSize);pointer:unique" json:"port_buffer"`
	// BufferSize: A variable to return the size, in bytes, of the port buffer.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
	// Return: The FaxObs_GetPort return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetPortResponse) xxx_ToOp(ctx context.Context, op *xxx_GetPortOperation) *xxx_GetPortOperation {
	if op == nil {
		op = &xxx_GetPortOperation{}
	}
	if o == nil {
		return op
	}
	o.PortBuffer = op.PortBuffer
	o.BufferSize = op.BufferSize
	o.Return = op.Return
	return op
}

func (o *GetPortResponse) xxx_FromOp(ctx context.Context, op *xxx_GetPortOperation) {
	if o == nil {
		return
	}
	o.PortBuffer = op.PortBuffer
	o.BufferSize = op.BufferSize
	o.Return = op.Return
}
func (o *GetPortResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetPortResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetPortOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetPortOperation structure represents the FaxObs_SetPort operation
type xxx_SetPortOperation struct {
	FaxPort  *fax.Port     `idl:"name:FaxPortHandle" json:"fax_port"`
	PortInfo *fax.PortInfo `idl:"name:PortInfo" json:"port_info"`
	Return   uint32        `idl:"name:Return" json:"return"`
}

func (o *xxx_SetPortOperation) OpNum() int { return 15 }

func (o *xxx_SetPortOperation) OpName() string { return "/faxobs/v4/FaxObs_SetPort" }

func (o *xxx_SetPortOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPortOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FaxPortHandle {in} (1:{context_handle, alias=RPC_FAX_PORT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort != nil {
			if err := o.FaxPort.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.Port{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// PortInfo {in} (1:{pointer=ref}*(1))(2:{alias=FAX_PORT_INFO}(struct))
	{
		if o.PortInfo != nil {
			if err := o.PortInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.PortInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPortOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FaxPortHandle {in} (1:{context_handle, alias=RPC_FAX_PORT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort == nil {
			o.FaxPort = &fax.Port{}
		}
		if err := o.FaxPort.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// PortInfo {in} (1:{pointer=ref}*(1))(2:{alias=FAX_PORT_INFO}(struct))
	{
		if o.PortInfo == nil {
			o.PortInfo = &fax.PortInfo{}
		}
		if err := o.PortInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPortOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPortOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetPortOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetPortRequest structure represents the FaxObs_SetPort operation request
type SetPortRequest struct {
	// FaxPortHandle: An RPC context handle that references a fax port.
	FaxPort *fax.Port `idl:"name:FaxPortHandle" json:"fax_port"`
	// PortInfo: A pointer to a FAX_PORT_INFO (section 2.2.7) structure. The structure contains
	// data to modify the configuration of the specified fax port. The server MUST ignore
	// the State field of this structure.
	PortInfo *fax.PortInfo `idl:"name:PortInfo" json:"port_info"`
}

func (o *SetPortRequest) xxx_ToOp(ctx context.Context, op *xxx_SetPortOperation) *xxx_SetPortOperation {
	if op == nil {
		op = &xxx_SetPortOperation{}
	}
	if o == nil {
		return op
	}
	o.FaxPort = op.FaxPort
	o.PortInfo = op.PortInfo
	return op
}

func (o *SetPortRequest) xxx_FromOp(ctx context.Context, op *xxx_SetPortOperation) {
	if o == nil {
		return
	}
	o.FaxPort = op.FaxPort
	o.PortInfo = op.PortInfo
}
func (o *SetPortRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetPortRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPortOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetPortResponse structure represents the FaxObs_SetPort operation response
type SetPortResponse struct {
	// Return: The FaxObs_SetPort return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetPortResponse) xxx_ToOp(ctx context.Context, op *xxx_SetPortOperation) *xxx_SetPortOperation {
	if op == nil {
		op = &xxx_SetPortOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *SetPortResponse) xxx_FromOp(ctx context.Context, op *xxx_SetPortOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetPortResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetPortResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetPortOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumRoutingMethodsOperation structure represents the FaxObs_EnumRoutingMethods operation
type xxx_EnumRoutingMethodsOperation struct {
	FaxPort               *fax.Port `idl:"name:FaxPortHandle" json:"fax_port"`
	RoutingInfoBuffer     []byte    `idl:"name:RoutingInfoBuffer;size_is:(, RoutingInfoBufferSize);pointer:unique" json:"routing_info_buffer"`
	RoutingInfoBufferSize uint32    `idl:"name:RoutingInfoBufferSize" json:"routing_info_buffer_size"`
	PortsReturned         uint32    `idl:"name:PortsReturned" json:"ports_returned"`
	Return                uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumRoutingMethodsOperation) OpNum() int { return 16 }

func (o *xxx_EnumRoutingMethodsOperation) OpName() string {
	return "/faxobs/v4/FaxObs_EnumRoutingMethods"
}

func (o *xxx_EnumRoutingMethodsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.RoutingInfoBuffer != nil && o.RoutingInfoBufferSize == 0 {
		o.RoutingInfoBufferSize = uint32(len(o.RoutingInfoBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRoutingMethodsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FaxPortHandle {in} (1:{context_handle, alias=RPC_FAX_PORT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort != nil {
			if err := o.FaxPort.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.Port{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// RoutingInfoBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=RoutingInfoBufferSize](uchar))
	{
		if o.RoutingInfoBuffer != nil {
			_ptr_RoutingInfoBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RoutingInfoBuffer != nil || o.RoutingInfoBufferSize > 0 {
					_ptr_RoutingInfoBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.RoutingInfoBufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.RoutingInfoBuffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.RoutingInfoBuffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.RoutingInfoBuffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
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
	// RoutingInfoBufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingInfoBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRoutingMethodsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FaxPortHandle {in} (1:{context_handle, alias=RPC_FAX_PORT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort == nil {
			o.FaxPort = &fax.Port{}
		}
		if err := o.FaxPort.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// RoutingInfoBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=RoutingInfoBufferSize](uchar))
	{
		_ptr_RoutingInfoBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_RoutingInfoBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.RoutingInfoBuffer", sizeInfo[0])
				}
				o.RoutingInfoBuffer = make([]byte, sizeInfo[0])
				for i1 := range o.RoutingInfoBuffer {
					i1 := i1
					if err := w.ReadData(&o.RoutingInfoBuffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_RoutingInfoBuffer := func(ptr interface{}) { o.RoutingInfoBuffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.RoutingInfoBuffer, _s_RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
				return err
			}
			return nil
		})
		_s_RoutingInfoBuffer := func(ptr interface{}) { o.RoutingInfoBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.RoutingInfoBuffer, _s_RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RoutingInfoBufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingInfoBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRoutingMethodsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.RoutingInfoBuffer != nil && o.RoutingInfoBufferSize == 0 {
		o.RoutingInfoBufferSize = uint32(len(o.RoutingInfoBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRoutingMethodsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// RoutingInfoBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=RoutingInfoBufferSize](uchar))
	{
		if o.RoutingInfoBuffer != nil {
			_ptr_RoutingInfoBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RoutingInfoBuffer != nil || o.RoutingInfoBufferSize > 0 {
					_ptr_RoutingInfoBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.RoutingInfoBufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.RoutingInfoBuffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.RoutingInfoBuffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.RoutingInfoBuffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
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
	// RoutingInfoBufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingInfoBufferSize); err != nil {
			return err
		}
	}
	// PortsReturned {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PortsReturned); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumRoutingMethodsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// RoutingInfoBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=RoutingInfoBufferSize](uchar))
	{
		_ptr_RoutingInfoBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_RoutingInfoBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.RoutingInfoBuffer", sizeInfo[0])
				}
				o.RoutingInfoBuffer = make([]byte, sizeInfo[0])
				for i1 := range o.RoutingInfoBuffer {
					i1 := i1
					if err := w.ReadData(&o.RoutingInfoBuffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_RoutingInfoBuffer := func(ptr interface{}) { o.RoutingInfoBuffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.RoutingInfoBuffer, _s_RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
				return err
			}
			return nil
		})
		_s_RoutingInfoBuffer := func(ptr interface{}) { o.RoutingInfoBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.RoutingInfoBuffer, _s_RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RoutingInfoBufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingInfoBufferSize); err != nil {
			return err
		}
	}
	// PortsReturned {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PortsReturned); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumRoutingMethodsRequest structure represents the FaxObs_EnumRoutingMethods operation request
type EnumRoutingMethodsRequest struct {
	// FaxPortHandle: An RPC context handle that references a fax port.
	FaxPort *fax.Port `idl:"name:FaxPortHandle" json:"fax_port"`
	// RoutingInfoBuffer: A pointer to the address of the returned buffer containing an
	// array of FAX_ROUTING_METHOD (section 2.2.9) structures. Each structure contains information
	// about one fax routing method.
	RoutingInfoBuffer []byte `idl:"name:RoutingInfoBuffer;size_is:(, RoutingInfoBufferSize);pointer:unique" json:"routing_info_buffer"`
	// RoutingInfoBufferSize: A variable to return the size, in bytes, of the routing method
	// buffer.
	RoutingInfoBufferSize uint32 `idl:"name:RoutingInfoBufferSize" json:"routing_info_buffer_size"`
}

func (o *EnumRoutingMethodsRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumRoutingMethodsOperation) *xxx_EnumRoutingMethodsOperation {
	if op == nil {
		op = &xxx_EnumRoutingMethodsOperation{}
	}
	if o == nil {
		return op
	}
	o.FaxPort = op.FaxPort
	o.RoutingInfoBuffer = op.RoutingInfoBuffer
	o.RoutingInfoBufferSize = op.RoutingInfoBufferSize
	return op
}

func (o *EnumRoutingMethodsRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumRoutingMethodsOperation) {
	if o == nil {
		return
	}
	o.FaxPort = op.FaxPort
	o.RoutingInfoBuffer = op.RoutingInfoBuffer
	o.RoutingInfoBufferSize = op.RoutingInfoBufferSize
}
func (o *EnumRoutingMethodsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumRoutingMethodsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumRoutingMethodsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumRoutingMethodsResponse structure represents the FaxObs_EnumRoutingMethods operation response
type EnumRoutingMethodsResponse struct {
	// RoutingInfoBuffer: A pointer to the address of the returned buffer containing an
	// array of FAX_ROUTING_METHOD (section 2.2.9) structures. Each structure contains information
	// about one fax routing method.
	RoutingInfoBuffer []byte `idl:"name:RoutingInfoBuffer;size_is:(, RoutingInfoBufferSize);pointer:unique" json:"routing_info_buffer"`
	// RoutingInfoBufferSize: A variable to return the size, in bytes, of the routing method
	// buffer.
	RoutingInfoBufferSize uint32 `idl:"name:RoutingInfoBufferSize" json:"routing_info_buffer_size"`
	// PortsReturned: A pointer to a DWORD ([MS-DTYP] section 2.2.9) variable to receive
	// the number of FAX_ROUTING_METHOD that are returned by the RoutingInfoBuffer parameter.
	PortsReturned uint32 `idl:"name:PortsReturned" json:"ports_returned"`
	// Return: The FaxObs_EnumRoutingMethods return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EnumRoutingMethodsResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumRoutingMethodsOperation) *xxx_EnumRoutingMethodsOperation {
	if op == nil {
		op = &xxx_EnumRoutingMethodsOperation{}
	}
	if o == nil {
		return op
	}
	o.RoutingInfoBuffer = op.RoutingInfoBuffer
	o.RoutingInfoBufferSize = op.RoutingInfoBufferSize
	o.PortsReturned = op.PortsReturned
	o.Return = op.Return
	return op
}

func (o *EnumRoutingMethodsResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumRoutingMethodsOperation) {
	if o == nil {
		return
	}
	o.RoutingInfoBuffer = op.RoutingInfoBuffer
	o.RoutingInfoBufferSize = op.RoutingInfoBufferSize
	o.PortsReturned = op.PortsReturned
	o.Return = op.Return
}
func (o *EnumRoutingMethodsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumRoutingMethodsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumRoutingMethodsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnableRoutingMethodOperation structure represents the FaxObs_EnableRoutingMethod operation
type xxx_EnableRoutingMethodOperation struct {
	FaxPort     *fax.Port `idl:"name:FaxPortHandle" json:"fax_port"`
	RoutingGUID string    `idl:"name:RoutingGuid;string;pointer:unique" json:"routing_guid"`
	Enabled     bool      `idl:"name:Enabled" json:"enabled"`
	Return      uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_EnableRoutingMethodOperation) OpNum() int { return 17 }

func (o *xxx_EnableRoutingMethodOperation) OpName() string {
	return "/faxobs/v4/FaxObs_EnableRoutingMethod"
}

func (o *xxx_EnableRoutingMethodOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableRoutingMethodOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FaxPortHandle {in} (1:{context_handle, alias=RPC_FAX_PORT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort != nil {
			if err := o.FaxPort.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.Port{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// RoutingGuid {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.RoutingGUID != "" {
			_ptr_RoutingGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.RoutingGUID); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.RoutingGUID, _ptr_RoutingGuid); err != nil {
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
	// Enabled {in} (1:{alias=BOOL}(int32))
	{
		if !o.Enabled {
			if err := w.WriteData(int32(0)); err != nil {
				return err
			}
		} else {
			if err := w.WriteData(int32(1)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_EnableRoutingMethodOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FaxPortHandle {in} (1:{context_handle, alias=RPC_FAX_PORT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort == nil {
			o.FaxPort = &fax.Port{}
		}
		if err := o.FaxPort.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// RoutingGuid {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_RoutingGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.RoutingGUID); err != nil {
				return err
			}
			return nil
		})
		_s_RoutingGuid := func(ptr interface{}) { o.RoutingGUID = *ptr.(*string) }
		if err := w.ReadPointer(&o.RoutingGUID, _s_RoutingGuid, _ptr_RoutingGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Enabled {in} (1:{alias=BOOL}(int32))
	{
		var _bEnabled int32
		if err := w.ReadData(&_bEnabled); err != nil {
			return err
		}
		o.Enabled = _bEnabled != 0
	}
	return nil
}

func (o *xxx_EnableRoutingMethodOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableRoutingMethodOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnableRoutingMethodOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnableRoutingMethodRequest structure represents the FaxObs_EnableRoutingMethod operation request
type EnableRoutingMethodRequest struct {
	// FaxPortHandle: An RPC context handle that references a fax port.
	FaxPort *fax.Port `idl:"name:FaxPortHandle" json:"fax_port"`
	// RoutingGuid: A curly braced GUID string that uniquely identifies the fax routing
	// method on which to act. For more information about routing methods, see [MSDN-FRM].
	// The routing methods and their curly braced GUID string values that can be used for
	// this parameter are discoverable by calling FaxObs_EnumRoutingMethods (section 3.1.4.2.18).
	// Included in this list are the default routing methods described in section 2.2.87.
	RoutingGUID string `idl:"name:RoutingGuid;string;pointer:unique" json:"routing_guid"`
	// Enabled: A Boolean variable that indicates whether the client request is to enable
	// (when set to TRUE) or disable (when set to FALSE) the fax routing method specified
	// by RoutingGuid.
	Enabled bool `idl:"name:Enabled" json:"enabled"`
}

func (o *EnableRoutingMethodRequest) xxx_ToOp(ctx context.Context, op *xxx_EnableRoutingMethodOperation) *xxx_EnableRoutingMethodOperation {
	if op == nil {
		op = &xxx_EnableRoutingMethodOperation{}
	}
	if o == nil {
		return op
	}
	o.FaxPort = op.FaxPort
	o.RoutingGUID = op.RoutingGUID
	o.Enabled = op.Enabled
	return op
}

func (o *EnableRoutingMethodRequest) xxx_FromOp(ctx context.Context, op *xxx_EnableRoutingMethodOperation) {
	if o == nil {
		return
	}
	o.FaxPort = op.FaxPort
	o.RoutingGUID = op.RoutingGUID
	o.Enabled = op.Enabled
}
func (o *EnableRoutingMethodRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnableRoutingMethodRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnableRoutingMethodOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnableRoutingMethodResponse structure represents the FaxObs_EnableRoutingMethod operation response
type EnableRoutingMethodResponse struct {
	// Return: The FaxObs_EnableRoutingMethod return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EnableRoutingMethodResponse) xxx_ToOp(ctx context.Context, op *xxx_EnableRoutingMethodOperation) *xxx_EnableRoutingMethodOperation {
	if op == nil {
		op = &xxx_EnableRoutingMethodOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *EnableRoutingMethodResponse) xxx_FromOp(ctx context.Context, op *xxx_EnableRoutingMethodOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *EnableRoutingMethodResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnableRoutingMethodResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnableRoutingMethodOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetRoutingInfoOperation structure represents the FaxObs_GetRoutingInfo operation
type xxx_GetRoutingInfoOperation struct {
	FaxPort               *fax.Port `idl:"name:FaxPortHandle" json:"fax_port"`
	RoutingGUID           string    `idl:"name:RoutingGuid;string;pointer:unique" json:"routing_guid"`
	RoutingInfoBuffer     []byte    `idl:"name:RoutingInfoBuffer;size_is:(, RoutingInfoBufferSize);pointer:unique" json:"routing_info_buffer"`
	RoutingInfoBufferSize uint32    `idl:"name:RoutingInfoBufferSize" json:"routing_info_buffer_size"`
	Return                uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_GetRoutingInfoOperation) OpNum() int { return 18 }

func (o *xxx_GetRoutingInfoOperation) OpName() string { return "/faxobs/v4/FaxObs_GetRoutingInfo" }

func (o *xxx_GetRoutingInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.RoutingInfoBuffer != nil && o.RoutingInfoBufferSize == 0 {
		o.RoutingInfoBufferSize = uint32(len(o.RoutingInfoBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRoutingInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FaxPortHandle {in} (1:{context_handle, alias=RPC_FAX_PORT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort != nil {
			if err := o.FaxPort.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.Port{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// RoutingGuid {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.RoutingGUID != "" {
			_ptr_RoutingGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.RoutingGUID); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.RoutingGUID, _ptr_RoutingGuid); err != nil {
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
	// RoutingInfoBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=RoutingInfoBufferSize](uchar))
	{
		if o.RoutingInfoBuffer != nil {
			_ptr_RoutingInfoBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RoutingInfoBuffer != nil || o.RoutingInfoBufferSize > 0 {
					_ptr_RoutingInfoBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.RoutingInfoBufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.RoutingInfoBuffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.RoutingInfoBuffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.RoutingInfoBuffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
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
	// RoutingInfoBufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingInfoBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRoutingInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FaxPortHandle {in} (1:{context_handle, alias=RPC_FAX_PORT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort == nil {
			o.FaxPort = &fax.Port{}
		}
		if err := o.FaxPort.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// RoutingGuid {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_RoutingGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.RoutingGUID); err != nil {
				return err
			}
			return nil
		})
		_s_RoutingGuid := func(ptr interface{}) { o.RoutingGUID = *ptr.(*string) }
		if err := w.ReadPointer(&o.RoutingGUID, _s_RoutingGuid, _ptr_RoutingGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RoutingInfoBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=RoutingInfoBufferSize](uchar))
	{
		_ptr_RoutingInfoBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_RoutingInfoBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.RoutingInfoBuffer", sizeInfo[0])
				}
				o.RoutingInfoBuffer = make([]byte, sizeInfo[0])
				for i1 := range o.RoutingInfoBuffer {
					i1 := i1
					if err := w.ReadData(&o.RoutingInfoBuffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_RoutingInfoBuffer := func(ptr interface{}) { o.RoutingInfoBuffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.RoutingInfoBuffer, _s_RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
				return err
			}
			return nil
		})
		_s_RoutingInfoBuffer := func(ptr interface{}) { o.RoutingInfoBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.RoutingInfoBuffer, _s_RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RoutingInfoBufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingInfoBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRoutingInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.RoutingInfoBuffer != nil && o.RoutingInfoBufferSize == 0 {
		o.RoutingInfoBufferSize = uint32(len(o.RoutingInfoBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRoutingInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// RoutingInfoBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=RoutingInfoBufferSize](uchar))
	{
		if o.RoutingInfoBuffer != nil {
			_ptr_RoutingInfoBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RoutingInfoBuffer != nil || o.RoutingInfoBufferSize > 0 {
					_ptr_RoutingInfoBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.RoutingInfoBufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.RoutingInfoBuffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.RoutingInfoBuffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.RoutingInfoBuffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
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
	// RoutingInfoBufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingInfoBufferSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetRoutingInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// RoutingInfoBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=RoutingInfoBufferSize](uchar))
	{
		_ptr_RoutingInfoBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_RoutingInfoBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.RoutingInfoBuffer", sizeInfo[0])
				}
				o.RoutingInfoBuffer = make([]byte, sizeInfo[0])
				for i1 := range o.RoutingInfoBuffer {
					i1 := i1
					if err := w.ReadData(&o.RoutingInfoBuffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_RoutingInfoBuffer := func(ptr interface{}) { o.RoutingInfoBuffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.RoutingInfoBuffer, _s_RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
				return err
			}
			return nil
		})
		_s_RoutingInfoBuffer := func(ptr interface{}) { o.RoutingInfoBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.RoutingInfoBuffer, _s_RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RoutingInfoBufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingInfoBufferSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetRoutingInfoRequest structure represents the FaxObs_GetRoutingInfo operation request
type GetRoutingInfoRequest struct {
	// FaxPortHandle: An RPC context handle that references a specified fax port.
	FaxPort *fax.Port `idl:"name:FaxPortHandle" json:"fax_port"`
	// RoutingGuid: A curly braced GUID string that uniquely identifies the fax routing
	// method for which to obtain the routing information. Fax routing methods are defined
	// by a fax-routing extension and each method is identified by a GUID. For more information
	// about routing methods, see [MSDN-FRM]. The routing methods and their curly braced
	// GUID string values that can be used for this parameter are discoverable by calling
	// FaxObs_EnumRoutingMethods (section 3.1.4.2.18). Included in this list are the default
	// routing methods described in section 2.2.87.
	RoutingGUID string `idl:"name:RoutingGuid;string;pointer:unique" json:"routing_guid"`
	// RoutingInfoBuffer: A pointer to the address of the returned buffer containing the
	// fax-routing information. The buffer format and contents depend on the routing method
	// that is identified by RoutingGuid.
	RoutingInfoBuffer []byte `idl:"name:RoutingInfoBuffer;size_is:(, RoutingInfoBufferSize);pointer:unique" json:"routing_info_buffer"`
	// RoutingInfoBufferSize: A pointer to a DWORD ([MS-DTYP] section 2.2.9) variable that
	// receives the size, in bytes, of the RoutingInfoBuffer buffer.
	RoutingInfoBufferSize uint32 `idl:"name:RoutingInfoBufferSize" json:"routing_info_buffer_size"`
}

func (o *GetRoutingInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetRoutingInfoOperation) *xxx_GetRoutingInfoOperation {
	if op == nil {
		op = &xxx_GetRoutingInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.FaxPort = op.FaxPort
	o.RoutingGUID = op.RoutingGUID
	o.RoutingInfoBuffer = op.RoutingInfoBuffer
	o.RoutingInfoBufferSize = op.RoutingInfoBufferSize
	return op
}

func (o *GetRoutingInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetRoutingInfoOperation) {
	if o == nil {
		return
	}
	o.FaxPort = op.FaxPort
	o.RoutingGUID = op.RoutingGUID
	o.RoutingInfoBuffer = op.RoutingInfoBuffer
	o.RoutingInfoBufferSize = op.RoutingInfoBufferSize
}
func (o *GetRoutingInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetRoutingInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRoutingInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetRoutingInfoResponse structure represents the FaxObs_GetRoutingInfo operation response
type GetRoutingInfoResponse struct {
	// RoutingInfoBuffer: A pointer to the address of the returned buffer containing the
	// fax-routing information. The buffer format and contents depend on the routing method
	// that is identified by RoutingGuid.
	RoutingInfoBuffer []byte `idl:"name:RoutingInfoBuffer;size_is:(, RoutingInfoBufferSize);pointer:unique" json:"routing_info_buffer"`
	// RoutingInfoBufferSize: A pointer to a DWORD ([MS-DTYP] section 2.2.9) variable that
	// receives the size, in bytes, of the RoutingInfoBuffer buffer.
	RoutingInfoBufferSize uint32 `idl:"name:RoutingInfoBufferSize" json:"routing_info_buffer_size"`
	// Return: The FaxObs_GetRoutingInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetRoutingInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetRoutingInfoOperation) *xxx_GetRoutingInfoOperation {
	if op == nil {
		op = &xxx_GetRoutingInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.RoutingInfoBuffer = op.RoutingInfoBuffer
	o.RoutingInfoBufferSize = op.RoutingInfoBufferSize
	o.Return = op.Return
	return op
}

func (o *GetRoutingInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetRoutingInfoOperation) {
	if o == nil {
		return
	}
	o.RoutingInfoBuffer = op.RoutingInfoBuffer
	o.RoutingInfoBufferSize = op.RoutingInfoBufferSize
	o.Return = op.Return
}
func (o *GetRoutingInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetRoutingInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetRoutingInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetRoutingInfoOperation structure represents the FaxObs_SetRoutingInfo operation
type xxx_SetRoutingInfoOperation struct {
	FaxPort               *fax.Port `idl:"name:FaxPortHandle" json:"fax_port"`
	RoutingGUID           string    `idl:"name:RoutingGuid;string;pointer:unique" json:"routing_guid"`
	RoutingInfoBuffer     []byte    `idl:"name:RoutingInfoBuffer;size_is:(RoutingInfoBufferSize);pointer:unique" json:"routing_info_buffer"`
	RoutingInfoBufferSize uint32    `idl:"name:RoutingInfoBufferSize" json:"routing_info_buffer_size"`
	Return                uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_SetRoutingInfoOperation) OpNum() int { return 19 }

func (o *xxx_SetRoutingInfoOperation) OpName() string { return "/faxobs/v4/FaxObs_SetRoutingInfo" }

func (o *xxx_SetRoutingInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.RoutingInfoBuffer != nil && o.RoutingInfoBufferSize == 0 {
		o.RoutingInfoBufferSize = uint32(len(o.RoutingInfoBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRoutingInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FaxPortHandle {in} (1:{context_handle, alias=RPC_FAX_PORT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort != nil {
			if err := o.FaxPort.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.Port{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// RoutingGuid {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.RoutingGUID != "" {
			_ptr_RoutingGuid := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.RoutingGUID); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.RoutingGUID, _ptr_RoutingGuid); err != nil {
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
	// RoutingInfoBuffer {in} (1:{pointer=unique}*(1))(2:{alias=BYTE}[dim:0,size_is=RoutingInfoBufferSize](uchar))
	{
		if o.RoutingInfoBuffer != nil || o.RoutingInfoBufferSize > 0 {
			_ptr_RoutingInfoBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.RoutingInfoBufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.RoutingInfoBuffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.RoutingInfoBuffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.RoutingInfoBuffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
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
	// RoutingInfoBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingInfoBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRoutingInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FaxPortHandle {in} (1:{context_handle, alias=RPC_FAX_PORT_HANDLE, names=ndr_context_handle}(struct))
	{
		if o.FaxPort == nil {
			o.FaxPort = &fax.Port{}
		}
		if err := o.FaxPort.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// RoutingGuid {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_RoutingGuid := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.RoutingGUID); err != nil {
				return err
			}
			return nil
		})
		_s_RoutingGuid := func(ptr interface{}) { o.RoutingGUID = *ptr.(*string) }
		if err := w.ReadPointer(&o.RoutingGUID, _s_RoutingGuid, _ptr_RoutingGuid); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RoutingInfoBuffer {in} (1:{pointer=unique}*(1))(2:{alias=BYTE}[dim:0,size_is=RoutingInfoBufferSize](uchar))
	{
		_ptr_RoutingInfoBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.RoutingInfoBuffer", sizeInfo[0])
			}
			o.RoutingInfoBuffer = make([]byte, sizeInfo[0])
			for i1 := range o.RoutingInfoBuffer {
				i1 := i1
				if err := w.ReadData(&o.RoutingInfoBuffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_RoutingInfoBuffer := func(ptr interface{}) { o.RoutingInfoBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.RoutingInfoBuffer, _s_RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RoutingInfoBufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingInfoBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRoutingInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRoutingInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetRoutingInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetRoutingInfoRequest structure represents the FaxObs_SetRoutingInfo operation request
type SetRoutingInfoRequest struct {
	// FaxPortHandle: An RPC context handle that references a specified fax port.
	FaxPort *fax.Port `idl:"name:FaxPortHandle" json:"fax_port"`
	// RoutingGuid: A curly braced GUID string that uniquely identifies the fax routing
	// method to set the routing information for. Fax routing methods are defined by a fax
	// routing extension, and the method is identified by a GUID. For more information about
	// routing methods, see [MSDN-FRM]. The routing methods and their curly braced GUID
	// string values, which can be used for this parameter, are discoverable by calling
	// FaxObs_EnumRoutingMethods (section 3.1.4.2.18). Included in this list are the default
	// routing methods described in section 2.2.87.
	RoutingGUID string `idl:"name:RoutingGuid;string;pointer:unique" json:"routing_guid"`
	// RoutingInfoBuffer: A pointer to a buffer that contains the fax routing information
	// to be set. The format and contents of this buffer depend on the routing method identified
	// by the RoutingGuid parameter.
	RoutingInfoBuffer []byte `idl:"name:RoutingInfoBuffer;size_is:(RoutingInfoBufferSize);pointer:unique" json:"routing_info_buffer"`
	// RoutingInfoBufferSize: The size, in bytes, of the RoutingInfoBuffer buffer.
	RoutingInfoBufferSize uint32 `idl:"name:RoutingInfoBufferSize" json:"routing_info_buffer_size"`
}

func (o *SetRoutingInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_SetRoutingInfoOperation) *xxx_SetRoutingInfoOperation {
	if op == nil {
		op = &xxx_SetRoutingInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.FaxPort = op.FaxPort
	o.RoutingGUID = op.RoutingGUID
	o.RoutingInfoBuffer = op.RoutingInfoBuffer
	o.RoutingInfoBufferSize = op.RoutingInfoBufferSize
	return op
}

func (o *SetRoutingInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_SetRoutingInfoOperation) {
	if o == nil {
		return
	}
	o.FaxPort = op.FaxPort
	o.RoutingGUID = op.RoutingGUID
	o.RoutingInfoBuffer = op.RoutingInfoBuffer
	o.RoutingInfoBufferSize = op.RoutingInfoBufferSize
}
func (o *SetRoutingInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetRoutingInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetRoutingInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetRoutingInfoResponse structure represents the FaxObs_SetRoutingInfo operation response
type SetRoutingInfoResponse struct {
	// Return: The FaxObs_SetRoutingInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetRoutingInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_SetRoutingInfoOperation) *xxx_SetRoutingInfoOperation {
	if op == nil {
		op = &xxx_SetRoutingInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *SetRoutingInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_SetRoutingInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetRoutingInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetRoutingInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetRoutingInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumGlobalRoutingInfoOperation structure represents the FaxObs_EnumGlobalRoutingInfo operation
type xxx_EnumGlobalRoutingInfoOperation struct {
	RoutingInfoBuffer     []byte `idl:"name:RoutingInfoBuffer;size_is:(, RoutingInfoBufferSize);pointer:unique" json:"routing_info_buffer"`
	RoutingInfoBufferSize uint32 `idl:"name:RoutingInfoBufferSize" json:"routing_info_buffer_size"`
	MethodsReturned       uint32 `idl:"name:MethodsReturned" json:"methods_returned"`
	Return                uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumGlobalRoutingInfoOperation) OpNum() int { return 20 }

func (o *xxx_EnumGlobalRoutingInfoOperation) OpName() string {
	return "/faxobs/v4/FaxObs_EnumGlobalRoutingInfo"
}

func (o *xxx_EnumGlobalRoutingInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.RoutingInfoBuffer != nil && o.RoutingInfoBufferSize == 0 {
		o.RoutingInfoBufferSize = uint32(len(o.RoutingInfoBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumGlobalRoutingInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// RoutingInfoBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=RoutingInfoBufferSize](uchar))
	{
		if o.RoutingInfoBuffer != nil {
			_ptr_RoutingInfoBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RoutingInfoBuffer != nil || o.RoutingInfoBufferSize > 0 {
					_ptr_RoutingInfoBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.RoutingInfoBufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.RoutingInfoBuffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.RoutingInfoBuffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.RoutingInfoBuffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
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
	// RoutingInfoBufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingInfoBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumGlobalRoutingInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// RoutingInfoBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=RoutingInfoBufferSize](uchar))
	{
		_ptr_RoutingInfoBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_RoutingInfoBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.RoutingInfoBuffer", sizeInfo[0])
				}
				o.RoutingInfoBuffer = make([]byte, sizeInfo[0])
				for i1 := range o.RoutingInfoBuffer {
					i1 := i1
					if err := w.ReadData(&o.RoutingInfoBuffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_RoutingInfoBuffer := func(ptr interface{}) { o.RoutingInfoBuffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.RoutingInfoBuffer, _s_RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
				return err
			}
			return nil
		})
		_s_RoutingInfoBuffer := func(ptr interface{}) { o.RoutingInfoBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.RoutingInfoBuffer, _s_RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RoutingInfoBufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingInfoBufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumGlobalRoutingInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.RoutingInfoBuffer != nil && o.RoutingInfoBufferSize == 0 {
		o.RoutingInfoBufferSize = uint32(len(o.RoutingInfoBuffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumGlobalRoutingInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// RoutingInfoBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=RoutingInfoBufferSize](uchar))
	{
		if o.RoutingInfoBuffer != nil {
			_ptr_RoutingInfoBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RoutingInfoBuffer != nil || o.RoutingInfoBufferSize > 0 {
					_ptr_RoutingInfoBuffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.RoutingInfoBufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.RoutingInfoBuffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.RoutingInfoBuffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.RoutingInfoBuffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
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
	// RoutingInfoBufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RoutingInfoBufferSize); err != nil {
			return err
		}
	}
	// MethodsReturned {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MethodsReturned); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumGlobalRoutingInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// RoutingInfoBuffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=RoutingInfoBufferSize](uchar))
	{
		_ptr_RoutingInfoBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_RoutingInfoBuffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.RoutingInfoBuffer", sizeInfo[0])
				}
				o.RoutingInfoBuffer = make([]byte, sizeInfo[0])
				for i1 := range o.RoutingInfoBuffer {
					i1 := i1
					if err := w.ReadData(&o.RoutingInfoBuffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_RoutingInfoBuffer := func(ptr interface{}) { o.RoutingInfoBuffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.RoutingInfoBuffer, _s_RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
				return err
			}
			return nil
		})
		_s_RoutingInfoBuffer := func(ptr interface{}) { o.RoutingInfoBuffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.RoutingInfoBuffer, _s_RoutingInfoBuffer, _ptr_RoutingInfoBuffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// RoutingInfoBufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RoutingInfoBufferSize); err != nil {
			return err
		}
	}
	// MethodsReturned {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MethodsReturned); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumGlobalRoutingInfoRequest structure represents the FaxObs_EnumGlobalRoutingInfo operation request
type EnumGlobalRoutingInfoRequest struct {
	// RoutingInfoBuffer: A pointer to the address of the returned buffer containing an
	// array of _FAX_GLOBAL_ROUTING_INFOW (section 2.2.33) structures. Each structure contains
	// information about one fax routing method, as it pertains to the entire fax server.
	RoutingInfoBuffer []byte `idl:"name:RoutingInfoBuffer;size_is:(, RoutingInfoBufferSize);pointer:unique" json:"routing_info_buffer"`
	// RoutingInfoBufferSize: A variable to return the size, in bytes, of the routing information
	// buffer.
	RoutingInfoBufferSize uint32 `idl:"name:RoutingInfoBufferSize" json:"routing_info_buffer_size"`
}

func (o *EnumGlobalRoutingInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumGlobalRoutingInfoOperation) *xxx_EnumGlobalRoutingInfoOperation {
	if op == nil {
		op = &xxx_EnumGlobalRoutingInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.RoutingInfoBuffer = op.RoutingInfoBuffer
	o.RoutingInfoBufferSize = op.RoutingInfoBufferSize
	return op
}

func (o *EnumGlobalRoutingInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumGlobalRoutingInfoOperation) {
	if o == nil {
		return
	}
	o.RoutingInfoBuffer = op.RoutingInfoBuffer
	o.RoutingInfoBufferSize = op.RoutingInfoBufferSize
}
func (o *EnumGlobalRoutingInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumGlobalRoutingInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumGlobalRoutingInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumGlobalRoutingInfoResponse structure represents the FaxObs_EnumGlobalRoutingInfo operation response
type EnumGlobalRoutingInfoResponse struct {
	// RoutingInfoBuffer: A pointer to the address of the returned buffer containing an
	// array of _FAX_GLOBAL_ROUTING_INFOW (section 2.2.33) structures. Each structure contains
	// information about one fax routing method, as it pertains to the entire fax server.
	RoutingInfoBuffer []byte `idl:"name:RoutingInfoBuffer;size_is:(, RoutingInfoBufferSize);pointer:unique" json:"routing_info_buffer"`
	// RoutingInfoBufferSize: A variable to return the size, in bytes, of the routing information
	// buffer.
	RoutingInfoBufferSize uint32 `idl:"name:RoutingInfoBufferSize" json:"routing_info_buffer_size"`
	// MethodsReturned: A pointer to a DWORD ([MS-DTYP] section 2.2.9) variable to receive
	// the number of _FAX_GLOBAL_ROUTING_INFOW structures that the method returns in the
	// RoutingInfoBuffer parameter. This number SHOULD equal the total number of fax routing
	// methods installed on the fax server.
	MethodsReturned uint32 `idl:"name:MethodsReturned" json:"methods_returned"`
	// Return: The FaxObs_EnumGlobalRoutingInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EnumGlobalRoutingInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumGlobalRoutingInfoOperation) *xxx_EnumGlobalRoutingInfoOperation {
	if op == nil {
		op = &xxx_EnumGlobalRoutingInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.RoutingInfoBuffer = op.RoutingInfoBuffer
	o.RoutingInfoBufferSize = op.RoutingInfoBufferSize
	o.MethodsReturned = op.MethodsReturned
	o.Return = op.Return
	return op
}

func (o *EnumGlobalRoutingInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumGlobalRoutingInfoOperation) {
	if o == nil {
		return
	}
	o.RoutingInfoBuffer = op.RoutingInfoBuffer
	o.RoutingInfoBufferSize = op.RoutingInfoBufferSize
	o.MethodsReturned = op.MethodsReturned
	o.Return = op.Return
}
func (o *EnumGlobalRoutingInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumGlobalRoutingInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumGlobalRoutingInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetGlobalRoutingInfoOperation structure represents the FaxObs_SetGlobalRoutingInfo operation
type xxx_SetGlobalRoutingInfoOperation struct {
	RoutingInfo *fax.GlobalRoutingInfoW `idl:"name:RoutingInfo" json:"routing_info"`
	Return      uint32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_SetGlobalRoutingInfoOperation) OpNum() int { return 21 }

func (o *xxx_SetGlobalRoutingInfoOperation) OpName() string {
	return "/faxobs/v4/FaxObs_SetGlobalRoutingInfo"
}

func (o *xxx_SetGlobalRoutingInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGlobalRoutingInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// RoutingInfo {in} (1:{pointer=ref}*(1))(2:{alias=FAX_GLOBAL_ROUTING_INFOW}(struct))
	{
		if o.RoutingInfo != nil {
			if err := o.RoutingInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.GlobalRoutingInfoW{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGlobalRoutingInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// RoutingInfo {in} (1:{pointer=ref}*(1))(2:{alias=FAX_GLOBAL_ROUTING_INFOW}(struct))
	{
		if o.RoutingInfo == nil {
			o.RoutingInfo = &fax.GlobalRoutingInfoW{}
		}
		if err := o.RoutingInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGlobalRoutingInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGlobalRoutingInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGlobalRoutingInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetGlobalRoutingInfoRequest structure represents the FaxObs_SetGlobalRoutingInfo operation request
type SetGlobalRoutingInfoRequest struct {
	// RoutingInfo: A pointer to a buffer that contains a FAX_GLOBAL_ROUTING_INFOW (section
	// 2.2.32) structure.
	RoutingInfo *fax.GlobalRoutingInfoW `idl:"name:RoutingInfo" json:"routing_info"`
}

func (o *SetGlobalRoutingInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_SetGlobalRoutingInfoOperation) *xxx_SetGlobalRoutingInfoOperation {
	if op == nil {
		op = &xxx_SetGlobalRoutingInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.RoutingInfo = op.RoutingInfo
	return op
}

func (o *SetGlobalRoutingInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_SetGlobalRoutingInfoOperation) {
	if o == nil {
		return
	}
	o.RoutingInfo = op.RoutingInfo
}
func (o *SetGlobalRoutingInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetGlobalRoutingInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetGlobalRoutingInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetGlobalRoutingInfoResponse structure represents the FaxObs_SetGlobalRoutingInfo operation response
type SetGlobalRoutingInfoResponse struct {
	// Return: The FaxObs_SetGlobalRoutingInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetGlobalRoutingInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_SetGlobalRoutingInfoOperation) *xxx_SetGlobalRoutingInfoOperation {
	if op == nil {
		op = &xxx_SetGlobalRoutingInfoOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *SetGlobalRoutingInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_SetGlobalRoutingInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetGlobalRoutingInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetGlobalRoutingInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetGlobalRoutingInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetConfigurationOperation structure represents the FaxObs_GetConfiguration operation
type xxx_GetConfigurationOperation struct {
	Buffer     []byte `idl:"name:Buffer;size_is:(, BufferSize);pointer:unique" json:"buffer"`
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetConfigurationOperation) OpNum() int { return 22 }

func (o *xxx_GetConfigurationOperation) OpName() string { return "/faxobs/v4/FaxObs_GetConfiguration" }

func (o *xxx_GetConfigurationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Buffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigurationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Buffer != nil || o.BufferSize > 0 {
					_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.Buffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.Buffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigurationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
				}
				o.Buffer = make([]byte, sizeInfo[0])
				for i1 := range o.Buffer {
					i1 := i1
					if err := w.ReadData(&o.Buffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigurationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Buffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigurationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Buffer != nil || o.BufferSize > 0 {
					_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.Buffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.Buffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetConfigurationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
				}
				o.Buffer = make([]byte, sizeInfo[0])
				for i1 := range o.Buffer {
					i1 := i1
					if err := w.ReadData(&o.Buffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetConfigurationRequest structure represents the FaxObs_GetConfiguration operation request
type GetConfigurationRequest struct {
	// Buffer: A pointer to the address of the returned buffer containing a _FAX_CONFIGURATIONW
	// (section 2.2.29) structure. The structure contains the current configuration settings
	// for the fax server.
	Buffer []byte `idl:"name:Buffer;size_is:(, BufferSize);pointer:unique" json:"buffer"`
	// BufferSize: A variable to return the size, in bytes, of the buffer.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
}

func (o *GetConfigurationRequest) xxx_ToOp(ctx context.Context, op *xxx_GetConfigurationOperation) *xxx_GetConfigurationOperation {
	if op == nil {
		op = &xxx_GetConfigurationOperation{}
	}
	if o == nil {
		return op
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	return op
}

func (o *GetConfigurationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetConfigurationOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
}
func (o *GetConfigurationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetConfigurationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConfigurationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetConfigurationResponse structure represents the FaxObs_GetConfiguration operation response
type GetConfigurationResponse struct {
	// Buffer: A pointer to the address of the returned buffer containing a _FAX_CONFIGURATIONW
	// (section 2.2.29) structure. The structure contains the current configuration settings
	// for the fax server.
	Buffer []byte `idl:"name:Buffer;size_is:(, BufferSize);pointer:unique" json:"buffer"`
	// BufferSize: A variable to return the size, in bytes, of the buffer.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
	// Return: The FaxObs_GetConfiguration return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetConfigurationResponse) xxx_ToOp(ctx context.Context, op *xxx_GetConfigurationOperation) *xxx_GetConfigurationOperation {
	if op == nil {
		op = &xxx_GetConfigurationOperation{}
	}
	if o == nil {
		return op
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	o.Return = op.Return
	return op
}

func (o *GetConfigurationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetConfigurationOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	o.Return = op.Return
}
func (o *GetConfigurationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetConfigurationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetConfigurationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetConfigurationOperation structure represents the FaxObs_SetConfiguration operation
type xxx_SetConfigurationOperation struct {
	FaxConfig *fax.ConfigW `idl:"name:FaxConfig" json:"fax_config"`
	Return    uint32       `idl:"name:Return" json:"return"`
}

func (o *xxx_SetConfigurationOperation) OpNum() int { return 23 }

func (o *xxx_SetConfigurationOperation) OpName() string { return "/faxobs/v4/FaxObs_SetConfiguration" }

func (o *xxx_SetConfigurationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetConfigurationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FaxConfig {in} (1:{pointer=ref}*(1))(2:{alias=FAX_CONFIGURATIONW}(struct))
	{
		if o.FaxConfig != nil {
			if err := o.FaxConfig.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&fax.ConfigW{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetConfigurationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FaxConfig {in} (1:{pointer=ref}*(1))(2:{alias=FAX_CONFIGURATIONW}(struct))
	{
		if o.FaxConfig == nil {
			o.FaxConfig = &fax.ConfigW{}
		}
		if err := o.FaxConfig.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetConfigurationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetConfigurationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetConfigurationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetConfigurationRequest structure represents the FaxObs_SetConfiguration operation request
type SetConfigurationRequest struct {
	// FaxConfig: A pointer to a FAX_CONFIGURATIONW (section 2.2.28) structure. If the Branding
	// structure member is TRUE, the fax server SHOULD generate a brand that contains transmission-related
	// information, such as the transmitting subscriber identifier, date, time, and page
	// count. If the UseDeviceTsid structure member is TRUE, the server SHOULD use the device's
	// transmitting subscriber identifier. If the ServerCp structure member is TRUE, the
	// client SHOULD use a common cover page stored on the fax server; if this member is
	// FALSE, the client SHOULD use a personal cover page template. If the PauseServerQueue
	// structure member is TRUE, the server SHOULD pause the outgoing fax queue. If the
	// ArchiveOutgoingFaxes structure member is TRUE, the server SHOULD archive transmissions
	// in the directory specified by the ArchiveDirectory member. The fax server SHOULD
	// ignore the ArchiveDirectory structure member if the ArchiveOutgoingFaxes member is
	// FALSE. The fax server SHOULD retain the discount time period submitted by the client
	// with the StartCheapTime and StopCheapTime structure members. <209>
	FaxConfig *fax.ConfigW `idl:"name:FaxConfig" json:"fax_config"`
}

func (o *SetConfigurationRequest) xxx_ToOp(ctx context.Context, op *xxx_SetConfigurationOperation) *xxx_SetConfigurationOperation {
	if op == nil {
		op = &xxx_SetConfigurationOperation{}
	}
	if o == nil {
		return op
	}
	o.FaxConfig = op.FaxConfig
	return op
}

func (o *SetConfigurationRequest) xxx_FromOp(ctx context.Context, op *xxx_SetConfigurationOperation) {
	if o == nil {
		return
	}
	o.FaxConfig = op.FaxConfig
}
func (o *SetConfigurationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetConfigurationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetConfigurationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetConfigurationResponse structure represents the FaxObs_SetConfiguration operation response
type SetConfigurationResponse struct {
	// Return: The FaxObs_SetConfiguration return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetConfigurationResponse) xxx_ToOp(ctx context.Context, op *xxx_SetConfigurationOperation) *xxx_SetConfigurationOperation {
	if op == nil {
		op = &xxx_SetConfigurationOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *SetConfigurationResponse) xxx_FromOp(ctx context.Context, op *xxx_SetConfigurationOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetConfigurationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetConfigurationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetConfigurationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetLoggingCategoriesOperation structure represents the FaxObs_GetLoggingCategories operation
type xxx_GetLoggingCategoriesOperation struct {
	Buffer           []byte `idl:"name:Buffer;size_is:(, BufferSize);pointer:unique" json:"buffer"`
	BufferSize       uint32 `idl:"name:BufferSize" json:"buffer_size"`
	NumberCategories uint32 `idl:"name:NumberCategories" json:"number_categories"`
	Return           uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetLoggingCategoriesOperation) OpNum() int { return 24 }

func (o *xxx_GetLoggingCategoriesOperation) OpName() string {
	return "/faxobs/v4/FaxObs_GetLoggingCategories"
}

func (o *xxx_GetLoggingCategoriesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Buffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLoggingCategoriesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Buffer != nil || o.BufferSize > 0 {
					_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.Buffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.Buffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	// NumberCategories {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NumberCategories); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLoggingCategoriesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
				}
				o.Buffer = make([]byte, sizeInfo[0])
				for i1 := range o.Buffer {
					i1 := i1
					if err := w.ReadData(&o.Buffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	// NumberCategories {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NumberCategories); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLoggingCategoriesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Buffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLoggingCategoriesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Buffer != nil || o.BufferSize > 0 {
					_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.Buffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.Buffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	// NumberCategories {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NumberCategories); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetLoggingCategoriesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
				}
				o.Buffer = make([]byte, sizeInfo[0])
				for i1 := range o.Buffer {
					i1 := i1
					if err := w.ReadData(&o.Buffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	// NumberCategories {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NumberCategories); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetLoggingCategoriesRequest structure represents the FaxObs_GetLoggingCategories operation request
type GetLoggingCategoriesRequest struct {
	// Buffer: A pointer to the address of the returned buffer containing an array of FAX_LOG_CATEGORY
	// (section 2.2.11) structures. The number of structures included in the array is set
	// by NumberCategories. Each structure describes one current logging category. The Name
	// strings are appended after the FAX_LOG_CATEGORY entries.
	Buffer []byte `idl:"name:Buffer;size_is:(, BufferSize);pointer:unique" json:"buffer"`
	// BufferSize: A variable to return the size, in bytes, of the buffer.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
	// NumberCategories: A pointer to a DWORD ([MS-DTYP] section 2.2.9) variable to receive
	// the number of FAX_LOG_CATEGORY structures that the method returns in the Buffer parameter.
	NumberCategories uint32 `idl:"name:NumberCategories" json:"number_categories"`
}

func (o *GetLoggingCategoriesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetLoggingCategoriesOperation) *xxx_GetLoggingCategoriesOperation {
	if op == nil {
		op = &xxx_GetLoggingCategoriesOperation{}
	}
	if o == nil {
		return op
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	o.NumberCategories = op.NumberCategories
	return op
}

func (o *GetLoggingCategoriesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetLoggingCategoriesOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	o.NumberCategories = op.NumberCategories
}
func (o *GetLoggingCategoriesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetLoggingCategoriesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLoggingCategoriesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetLoggingCategoriesResponse structure represents the FaxObs_GetLoggingCategories operation response
type GetLoggingCategoriesResponse struct {
	// Buffer: A pointer to the address of the returned buffer containing an array of FAX_LOG_CATEGORY
	// (section 2.2.11) structures. The number of structures included in the array is set
	// by NumberCategories. Each structure describes one current logging category. The Name
	// strings are appended after the FAX_LOG_CATEGORY entries.
	Buffer []byte `idl:"name:Buffer;size_is:(, BufferSize);pointer:unique" json:"buffer"`
	// BufferSize: A variable to return the size, in bytes, of the buffer.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
	// NumberCategories: A pointer to a DWORD ([MS-DTYP] section 2.2.9) variable to receive
	// the number of FAX_LOG_CATEGORY structures that the method returns in the Buffer parameter.
	NumberCategories uint32 `idl:"name:NumberCategories" json:"number_categories"`
	// Return: The FaxObs_GetLoggingCategories return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetLoggingCategoriesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetLoggingCategoriesOperation) *xxx_GetLoggingCategoriesOperation {
	if op == nil {
		op = &xxx_GetLoggingCategoriesOperation{}
	}
	if o == nil {
		return op
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	o.NumberCategories = op.NumberCategories
	o.Return = op.Return
	return op
}

func (o *GetLoggingCategoriesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetLoggingCategoriesOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	o.NumberCategories = op.NumberCategories
	o.Return = op.Return
}
func (o *GetLoggingCategoriesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetLoggingCategoriesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetLoggingCategoriesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetLoggingCategoriesOperation structure represents the FaxObs_SetLoggingCategories operation
type xxx_SetLoggingCategoriesOperation struct {
	Buffer           []byte `idl:"name:Buffer;size_is:(BufferSize);pointer:unique" json:"buffer"`
	BufferSize       uint32 `idl:"name:BufferSize" json:"buffer_size"`
	NumberCategories uint32 `idl:"name:NumberCategories" json:"number_categories"`
	Return           uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SetLoggingCategoriesOperation) OpNum() int { return 25 }

func (o *xxx_SetLoggingCategoriesOperation) OpName() string {
	return "/faxobs/v4/FaxObs_SetLoggingCategories"
}

func (o *xxx_SetLoggingCategoriesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Buffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLoggingCategoriesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Buffer {in} (1:{pointer=unique, alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.Buffer != nil || o.BufferSize > 0 {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.BufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Buffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Buffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
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
	// BufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	// NumberCategories {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NumberCategories); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLoggingCategoriesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Buffer {in} (1:{pointer=unique, alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
			}
			o.Buffer = make([]byte, sizeInfo[0])
			for i1 := range o.Buffer {
				i1 := i1
				if err := w.ReadData(&o.Buffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	// NumberCategories {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NumberCategories); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLoggingCategoriesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLoggingCategoriesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetLoggingCategoriesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetLoggingCategoriesRequest structure represents the FaxObs_SetLoggingCategories operation request
type SetLoggingCategoriesRequest struct {
	// Buffer: A pointer to an array of FAX_LOG_CATEGORY (section 2.2.11) structures. Each
	// structure contains the data to modify one logging category. The data includes a friendly
	// name of the logging category, a numeric identifier for the category, and the current
	// severity-level threshold for the category. For more information, see [MSDN-FSCAR].
	Buffer []byte `idl:"name:Buffer;size_is:(BufferSize);pointer:unique" json:"buffer"`
	// BufferSize: A DWORD ([MS-DTYP] section 2.2.9) variable that contains the size, in
	// bytes, of the data buffer.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
	// NumberCategories: A DWORD variable that contains the number of FAX_LOG_CATEGORY structures
	// that the method passes in the Buffer parameter.
	NumberCategories uint32 `idl:"name:NumberCategories" json:"number_categories"`
}

func (o *SetLoggingCategoriesRequest) xxx_ToOp(ctx context.Context, op *xxx_SetLoggingCategoriesOperation) *xxx_SetLoggingCategoriesOperation {
	if op == nil {
		op = &xxx_SetLoggingCategoriesOperation{}
	}
	if o == nil {
		return op
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	o.NumberCategories = op.NumberCategories
	return op
}

func (o *SetLoggingCategoriesRequest) xxx_FromOp(ctx context.Context, op *xxx_SetLoggingCategoriesOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	o.NumberCategories = op.NumberCategories
}
func (o *SetLoggingCategoriesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetLoggingCategoriesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLoggingCategoriesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetLoggingCategoriesResponse structure represents the FaxObs_SetLoggingCategories operation response
type SetLoggingCategoriesResponse struct {
	// Return: The FaxObs_SetLoggingCategories return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetLoggingCategoriesResponse) xxx_ToOp(ctx context.Context, op *xxx_SetLoggingCategoriesOperation) *xxx_SetLoggingCategoriesOperation {
	if op == nil {
		op = &xxx_SetLoggingCategoriesOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *SetLoggingCategoriesResponse) xxx_FromOp(ctx context.Context, op *xxx_SetLoggingCategoriesOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetLoggingCategoriesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetLoggingCategoriesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetLoggingCategoriesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetTAPILocationsOperation structure represents the FaxObs_GetTapiLocations operation
type xxx_GetTAPILocationsOperation struct {
	Buffer     []byte `idl:"name:Buffer;size_is:(, BufferSize);pointer:unique" json:"buffer"`
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetTAPILocationsOperation) OpNum() int { return 26 }

func (o *xxx_GetTAPILocationsOperation) OpName() string { return "/faxobs/v4/FaxObs_GetTapiLocations" }

func (o *xxx_GetTAPILocationsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Buffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTAPILocationsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Buffer != nil || o.BufferSize > 0 {
					_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.Buffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.Buffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTAPILocationsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
				}
				o.Buffer = make([]byte, sizeInfo[0])
				for i1 := range o.Buffer {
					i1 := i1
					if err := w.ReadData(&o.Buffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTAPILocationsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Buffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTAPILocationsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.Buffer != nil {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Buffer != nil || o.BufferSize > 0 {
					_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.Buffer {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.Buffer[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetTAPILocationsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Buffer {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
				}
				o.Buffer = make([]byte, sizeInfo[0])
				for i1 := range o.Buffer {
					i1 := i1
					if err := w.ReadData(&o.Buffer[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
				return err
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetTAPILocationsRequest structure represents the FaxObs_GetTapiLocations operation request
type GetTAPILocationsRequest struct {
	// Buffer: A pointer to the address of the returned buffer containing a FAX_TAPI_LOCATION_INFO
	// (section 2.2.89) structure that contains a list of FAX_TAPI_LOCATIONS, each FAX_TAPI_LOCATIONS
	// describes one location. Each structure includes information such as a friendly name,
	// country code, and area code. The current location can be identified in this list
	// of FAX_TAPI_LOCATIONS by searching for a structure with the location identifier value
	// (contained by the PermanentLocationID member) described by the CurrentLocationID
	// member of the FAX_TAPI_LOCATION_INFO.
	Buffer []byte `idl:"name:Buffer;size_is:(, BufferSize);pointer:unique" json:"buffer"`
	// BufferSize: Pointer to a DWORD ([MS-DTYP] section 2.2.9) variable that receives the
	// size, in bytes, of the data returned in the buffer.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
}

func (o *GetTAPILocationsRequest) xxx_ToOp(ctx context.Context, op *xxx_GetTAPILocationsOperation) *xxx_GetTAPILocationsOperation {
	if op == nil {
		op = &xxx_GetTAPILocationsOperation{}
	}
	if o == nil {
		return op
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	return op
}

func (o *GetTAPILocationsRequest) xxx_FromOp(ctx context.Context, op *xxx_GetTAPILocationsOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
}
func (o *GetTAPILocationsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetTAPILocationsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTAPILocationsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetTAPILocationsResponse structure represents the FaxObs_GetTapiLocations operation response
type GetTAPILocationsResponse struct {
	// Buffer: A pointer to the address of the returned buffer containing a FAX_TAPI_LOCATION_INFO
	// (section 2.2.89) structure that contains a list of FAX_TAPI_LOCATIONS, each FAX_TAPI_LOCATIONS
	// describes one location. Each structure includes information such as a friendly name,
	// country code, and area code. The current location can be identified in this list
	// of FAX_TAPI_LOCATIONS by searching for a structure with the location identifier value
	// (contained by the PermanentLocationID member) described by the CurrentLocationID
	// member of the FAX_TAPI_LOCATION_INFO.
	Buffer []byte `idl:"name:Buffer;size_is:(, BufferSize);pointer:unique" json:"buffer"`
	// BufferSize: Pointer to a DWORD ([MS-DTYP] section 2.2.9) variable that receives the
	// size, in bytes, of the data returned in the buffer.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
	// Return: The FaxObs_GetTapiLocations return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetTAPILocationsResponse) xxx_ToOp(ctx context.Context, op *xxx_GetTAPILocationsOperation) *xxx_GetTAPILocationsOperation {
	if op == nil {
		op = &xxx_GetTAPILocationsOperation{}
	}
	if o == nil {
		return op
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	o.Return = op.Return
	return op
}

func (o *GetTAPILocationsResponse) xxx_FromOp(ctx context.Context, op *xxx_GetTAPILocationsOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	o.Return = op.Return
}
func (o *GetTAPILocationsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetTAPILocationsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetTAPILocationsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetTAPILocationsOperation structure represents the FaxObs_SetTapiLocations operation
type xxx_SetTAPILocationsOperation struct {
	Buffer     []byte `idl:"name:Buffer;size_is:(BufferSize);pointer:unique" json:"buffer"`
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SetTAPILocationsOperation) OpNum() int { return 27 }

func (o *xxx_SetTAPILocationsOperation) OpName() string { return "/faxobs/v4/FaxObs_SetTapiLocations" }

func (o *xxx_SetTAPILocationsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Buffer != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTAPILocationsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Buffer {in} (1:{pointer=unique, alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.Buffer != nil || o.BufferSize > 0 {
			_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.BufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.Buffer {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Buffer[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
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
	// BufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTAPILocationsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Buffer {in} (1:{pointer=unique, alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
			}
			o.Buffer = make([]byte, sizeInfo[0])
			for i1 := range o.Buffer {
				i1 := i1
				if err := w.ReadData(&o.Buffer[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTAPILocationsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTAPILocationsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetTAPILocationsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetTAPILocationsRequest structure represents the FaxObs_SetTapiLocations operation request
type SetTAPILocationsRequest struct {
	// Buffer: A pointer to a FAX_TAPI_LOCATION_INFO (section 2.2.89) structure containing
	// a list of FAX_TAPI_LOCATIONS, each structure describing one location. The current
	// location is identified in this list of by the FAX_TAPI_LOCATIONS with the location
	// identifier value (contained by the PermanentLocationID member) described by the CurrentLocationID
	// member of the FAX_TAPI_LOCATION_INFO.
	Buffer []byte `idl:"name:Buffer;size_is:(BufferSize);pointer:unique" json:"buffer"`
	// BufferSize: Pointer to a DWORD ([MS-DTYP] section 2.2.9) variable that contains the
	// size, in bytes, of the data contained in the buffer.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
}

func (o *SetTAPILocationsRequest) xxx_ToOp(ctx context.Context, op *xxx_SetTAPILocationsOperation) *xxx_SetTAPILocationsOperation {
	if op == nil {
		op = &xxx_SetTAPILocationsOperation{}
	}
	if o == nil {
		return op
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
	return op
}

func (o *SetTAPILocationsRequest) xxx_FromOp(ctx context.Context, op *xxx_SetTAPILocationsOperation) {
	if o == nil {
		return
	}
	o.Buffer = op.Buffer
	o.BufferSize = op.BufferSize
}
func (o *SetTAPILocationsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetTAPILocationsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTAPILocationsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetTAPILocationsResponse structure represents the FaxObs_SetTapiLocations operation response
type SetTAPILocationsResponse struct {
	// Return: The FaxObs_SetTapiLocations return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetTAPILocationsResponse) xxx_ToOp(ctx context.Context, op *xxx_SetTAPILocationsOperation) *xxx_SetTAPILocationsOperation {
	if op == nil {
		op = &xxx_SetTAPILocationsOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *SetTAPILocationsResponse) xxx_FromOp(ctx context.Context, op *xxx_SetTAPILocationsOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetTAPILocationsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetTAPILocationsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetTAPILocationsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetMAPIProfilesOperation structure represents the FaxObs_GetMapiProfiles operation
type xxx_GetMAPIProfilesOperation struct {
	MAPIProfiles []byte `idl:"name:MapiProfiles;size_is:(, BufferSize);pointer:unique" json:"mapi_profiles"`
	BufferSize   uint32 `idl:"name:BufferSize" json:"buffer_size"`
	Return       uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetMAPIProfilesOperation) OpNum() int { return 28 }

func (o *xxx_GetMAPIProfilesOperation) OpName() string { return "/faxobs/v4/FaxObs_GetMapiProfiles" }

func (o *xxx_GetMAPIProfilesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.MAPIProfiles != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.MAPIProfiles))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMAPIProfilesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// MapiProfiles {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.MAPIProfiles != nil {
			_ptr_MapiProfiles := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MAPIProfiles != nil || o.BufferSize > 0 {
					_ptr_MapiProfiles := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.MAPIProfiles {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.MAPIProfiles[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.MAPIProfiles); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.MAPIProfiles, _ptr_MapiProfiles); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MAPIProfiles, _ptr_MapiProfiles); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMAPIProfilesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// MapiProfiles {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_MapiProfiles := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_MapiProfiles := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.MAPIProfiles", sizeInfo[0])
				}
				o.MAPIProfiles = make([]byte, sizeInfo[0])
				for i1 := range o.MAPIProfiles {
					i1 := i1
					if err := w.ReadData(&o.MAPIProfiles[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_MapiProfiles := func(ptr interface{}) { o.MAPIProfiles = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.MAPIProfiles, _s_MapiProfiles, _ptr_MapiProfiles); err != nil {
				return err
			}
			return nil
		})
		_s_MapiProfiles := func(ptr interface{}) { o.MAPIProfiles = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.MAPIProfiles, _s_MapiProfiles, _ptr_MapiProfiles); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMAPIProfilesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.MAPIProfiles != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.MAPIProfiles))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMAPIProfilesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// MapiProfiles {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.MAPIProfiles != nil {
			_ptr_MapiProfiles := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.MAPIProfiles != nil || o.BufferSize > 0 {
					_ptr_MapiProfiles := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.MAPIProfiles {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.MAPIProfiles[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.MAPIProfiles); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.MAPIProfiles, _ptr_MapiProfiles); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.MAPIProfiles, _ptr_MapiProfiles); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetMAPIProfilesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// MapiProfiles {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_MapiProfiles := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_MapiProfiles := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.MAPIProfiles", sizeInfo[0])
				}
				o.MAPIProfiles = make([]byte, sizeInfo[0])
				for i1 := range o.MAPIProfiles {
					i1 := i1
					if err := w.ReadData(&o.MAPIProfiles[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_MapiProfiles := func(ptr interface{}) { o.MAPIProfiles = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.MAPIProfiles, _s_MapiProfiles, _ptr_MapiProfiles); err != nil {
				return err
			}
			return nil
		})
		_s_MapiProfiles := func(ptr interface{}) { o.MAPIProfiles = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.MAPIProfiles, _s_MapiProfiles, _ptr_MapiProfiles); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetMAPIProfilesRequest structure represents the FaxObs_GetMapiProfiles operation request
type GetMAPIProfilesRequest struct {
	// MapiProfiles: A pointer to the address of the returned buffer. This buffer contains
	// a sequence of null-terminated character strings; each of these strings contains the
	// name of a MAPI profile. The sequence is terminated by an empty null-terminated character
	// string.
	MAPIProfiles []byte `idl:"name:MapiProfiles;size_is:(, BufferSize);pointer:unique" json:"mapi_profiles"`
	// BufferSize: Pointer to a DWORD ([MS-DTYP] section 2.2.9) variable that receives the
	// size, in bytes, of the data returned in the MapiProfiles character strings sequence.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
}

func (o *GetMAPIProfilesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetMAPIProfilesOperation) *xxx_GetMAPIProfilesOperation {
	if op == nil {
		op = &xxx_GetMAPIProfilesOperation{}
	}
	if o == nil {
		return op
	}
	o.MAPIProfiles = op.MAPIProfiles
	o.BufferSize = op.BufferSize
	return op
}

func (o *GetMAPIProfilesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetMAPIProfilesOperation) {
	if o == nil {
		return
	}
	o.MAPIProfiles = op.MAPIProfiles
	o.BufferSize = op.BufferSize
}
func (o *GetMAPIProfilesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetMAPIProfilesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMAPIProfilesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetMAPIProfilesResponse structure represents the FaxObs_GetMapiProfiles operation response
type GetMAPIProfilesResponse struct {
	// MapiProfiles: A pointer to the address of the returned buffer. This buffer contains
	// a sequence of null-terminated character strings; each of these strings contains the
	// name of a MAPI profile. The sequence is terminated by an empty null-terminated character
	// string.
	MAPIProfiles []byte `idl:"name:MapiProfiles;size_is:(, BufferSize);pointer:unique" json:"mapi_profiles"`
	// BufferSize: Pointer to a DWORD ([MS-DTYP] section 2.2.9) variable that receives the
	// size, in bytes, of the data returned in the MapiProfiles character strings sequence.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
	// Return: The FaxObs_GetMapiProfiles return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetMAPIProfilesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetMAPIProfilesOperation) *xxx_GetMAPIProfilesOperation {
	if op == nil {
		op = &xxx_GetMAPIProfilesOperation{}
	}
	if o == nil {
		return op
	}
	o.MAPIProfiles = op.MAPIProfiles
	o.BufferSize = op.BufferSize
	o.Return = op.Return
	return op
}

func (o *GetMAPIProfilesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetMAPIProfilesOperation) {
	if o == nil {
		return
	}
	o.MAPIProfiles = op.MAPIProfiles
	o.BufferSize = op.BufferSize
	o.Return = op.Return
}
func (o *GetMAPIProfilesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetMAPIProfilesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetMAPIProfilesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_StartClientServerOperation structure represents the FaxObs_StartClientServer operation
type xxx_StartClientServerOperation struct {
	MachineName string `idl:"name:MachineName;string;pointer:unique" json:"machine_name"`
	ClientName  string `idl:"name:ClientName;string;pointer:unique" json:"client_name"`
	Context     uint64 `idl:"name:Context" json:"context"`
	Return      uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_StartClientServerOperation) OpNum() int { return 29 }

func (o *xxx_StartClientServerOperation) OpName() string {
	return "/faxobs/v4/FaxObs_StartClientServer"
}

func (o *xxx_StartClientServerOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartClientServerOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// MachineName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.MachineName != "" {
			_ptr_MachineName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.MachineName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.MachineName, _ptr_MachineName); err != nil {
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
	// ClientName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.ClientName != "" {
			_ptr_ClientName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ClientName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ClientName, _ptr_ClientName); err != nil {
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
	// Context {in} (1:{alias=ULONG64}(uint64))
	{
		if err := w.WriteData(o.Context); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartClientServerOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// MachineName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_MachineName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.MachineName); err != nil {
				return err
			}
			return nil
		})
		_s_MachineName := func(ptr interface{}) { o.MachineName = *ptr.(*string) }
		if err := w.ReadPointer(&o.MachineName, _s_MachineName, _ptr_MachineName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ClientName {in} (1:{string, pointer=unique, alias=LPCWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_ClientName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ClientName); err != nil {
				return err
			}
			return nil
		})
		_s_ClientName := func(ptr interface{}) { o.ClientName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ClientName, _s_ClientName, _ptr_ClientName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Context {in} (1:{alias=ULONG64}(uint64))
	{
		if err := w.ReadData(&o.Context); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartClientServerOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartClientServerOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_StartClientServerOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// StartClientServerRequest structure represents the FaxObs_StartClientServer operation request
type StartClientServerRequest struct {
	// MachineName: A pointer to a null-terminated character string containing the name
	// of the fax client machine. The machine name MUST be NULL for a local machine and
	// a fully qualified domain name (FQDN) for a remote machine.
	MachineName string `idl:"name:MachineName;string;pointer:unique" json:"machine_name"`
	// ClientName: A pointer to a null-terminated character string containing the friendly
	// name of the fax client application. This name MUST be unique for each fax client
	// application running on the same fax client machine.
	ClientName string `idl:"name:ClientName;string;pointer:unique" json:"client_name"`
	// Context: A ULONG64 ([MS-DTYP] section 2.2.54) value that can be passed to FAX_OpenConnection
	// as a notification context. This context is equivalent to the subscription context
	// used in the Fax Server Interface methods FAX_StartServerNotification (section 3.1.4.1.100),
	// FAX_StartServerNotificationEx (section 3.1.4.1.101), and FAX_StartServerNotificationEx2
	// (section 3.1.4.1.102), with the difference that the FaxObs Server Interface (section
	// 3.1.4.2) does not have a method similar to FAX_EndServerNotification (section 3.1.4.1.17)
	// that the client calls to close this context.
	Context uint64 `idl:"name:Context" json:"context"`
}

func (o *StartClientServerRequest) xxx_ToOp(ctx context.Context, op *xxx_StartClientServerOperation) *xxx_StartClientServerOperation {
	if op == nil {
		op = &xxx_StartClientServerOperation{}
	}
	if o == nil {
		return op
	}
	o.MachineName = op.MachineName
	o.ClientName = op.ClientName
	o.Context = op.Context
	return op
}

func (o *StartClientServerRequest) xxx_FromOp(ctx context.Context, op *xxx_StartClientServerOperation) {
	if o == nil {
		return
	}
	o.MachineName = op.MachineName
	o.ClientName = op.ClientName
	o.Context = op.Context
}
func (o *StartClientServerRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *StartClientServerRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StartClientServerOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// StartClientServerResponse structure represents the FaxObs_StartClientServer operation response
type StartClientServerResponse struct {
	// Return: The FaxObs_StartClientServer return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *StartClientServerResponse) xxx_ToOp(ctx context.Context, op *xxx_StartClientServerOperation) *xxx_StartClientServerOperation {
	if op == nil {
		op = &xxx_StartClientServerOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *StartClientServerResponse) xxx_FromOp(ctx context.Context, op *xxx_StartClientServerOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *StartClientServerResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *StartClientServerResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_StartClientServerOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSecurityDescriptorOperation structure represents the FaxObs_GetSecurityDescriptor operation
type xxx_GetSecurityDescriptorOperation struct {
	ID                    uint32 `idl:"name:Id" json:"id"`
	FaxSecurityDescriptor []byte `idl:"name:FaxSecurityDescriptor;size_is:(, BufferSize);pointer:unique" json:"fax_security_descriptor"`
	BufferSize            uint32 `idl:"name:BufferSize" json:"buffer_size"`
	Return                uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSecurityDescriptorOperation) OpNum() int { return 31 }

func (o *xxx_GetSecurityDescriptorOperation) OpName() string {
	return "/faxobs/v4/FaxObs_GetSecurityDescriptor"
}

func (o *xxx_GetSecurityDescriptorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.FaxSecurityDescriptor != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.FaxSecurityDescriptor))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSecurityDescriptorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Id {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ID); err != nil {
			return err
		}
	}
	// FaxSecurityDescriptor {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.FaxSecurityDescriptor != nil {
			_ptr_FaxSecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FaxSecurityDescriptor != nil || o.BufferSize > 0 {
					_ptr_FaxSecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.FaxSecurityDescriptor {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.FaxSecurityDescriptor[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.FaxSecurityDescriptor); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.FaxSecurityDescriptor, _ptr_FaxSecurityDescriptor); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FaxSecurityDescriptor, _ptr_FaxSecurityDescriptor); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSecurityDescriptorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Id {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ID); err != nil {
			return err
		}
	}
	// FaxSecurityDescriptor {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_FaxSecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_FaxSecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.FaxSecurityDescriptor", sizeInfo[0])
				}
				o.FaxSecurityDescriptor = make([]byte, sizeInfo[0])
				for i1 := range o.FaxSecurityDescriptor {
					i1 := i1
					if err := w.ReadData(&o.FaxSecurityDescriptor[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_FaxSecurityDescriptor := func(ptr interface{}) { o.FaxSecurityDescriptor = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.FaxSecurityDescriptor, _s_FaxSecurityDescriptor, _ptr_FaxSecurityDescriptor); err != nil {
				return err
			}
			return nil
		})
		_s_FaxSecurityDescriptor := func(ptr interface{}) { o.FaxSecurityDescriptor = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.FaxSecurityDescriptor, _s_FaxSecurityDescriptor, _ptr_FaxSecurityDescriptor); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSecurityDescriptorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.FaxSecurityDescriptor != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.FaxSecurityDescriptor))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSecurityDescriptorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// FaxSecurityDescriptor {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.FaxSecurityDescriptor != nil {
			_ptr_FaxSecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.FaxSecurityDescriptor != nil || o.BufferSize > 0 {
					_ptr_FaxSecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						dimSize1 := uint64(o.BufferSize)
						if err := w.WriteSize(dimSize1); err != nil {
							return err
						}
						sizeInfo := []uint64{
							dimSize1,
						}
						for i1 := range o.FaxSecurityDescriptor {
							i1 := i1
							if uint64(i1) >= sizeInfo[0] {
								break
							}
							if err := w.WriteData(o.FaxSecurityDescriptor[i1]); err != nil {
								return err
							}
						}
						for i1 := len(o.FaxSecurityDescriptor); uint64(i1) < sizeInfo[0]; i1++ {
							if err := w.WriteData(uint8(0)); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.FaxSecurityDescriptor, _ptr_FaxSecurityDescriptor); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FaxSecurityDescriptor, _ptr_FaxSecurityDescriptor); err != nil {
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
	// BufferSize {in, out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSecurityDescriptorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// FaxSecurityDescriptor {in, out} (1:{pointer=unique}*(2))(2:{alias=LPBYTE,pointer=ref}*(1))(3:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_FaxSecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_FaxSecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				sizeInfo := []uint64{
					0,
				}
				for sz1 := range sizeInfo {
					if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
						return err
					}
				}
				if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
					return fmt.Errorf("buffer overflow for size %d of array o.FaxSecurityDescriptor", sizeInfo[0])
				}
				o.FaxSecurityDescriptor = make([]byte, sizeInfo[0])
				for i1 := range o.FaxSecurityDescriptor {
					i1 := i1
					if err := w.ReadData(&o.FaxSecurityDescriptor[i1]); err != nil {
						return err
					}
				}
				return nil
			})
			_s_FaxSecurityDescriptor := func(ptr interface{}) { o.FaxSecurityDescriptor = *ptr.(*[]byte) }
			if err := w.ReadPointer(&o.FaxSecurityDescriptor, _s_FaxSecurityDescriptor, _ptr_FaxSecurityDescriptor); err != nil {
				return err
			}
			return nil
		})
		_s_FaxSecurityDescriptor := func(ptr interface{}) { o.FaxSecurityDescriptor = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.FaxSecurityDescriptor, _s_FaxSecurityDescriptor, _ptr_FaxSecurityDescriptor); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in, out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetSecurityDescriptorRequest structure represents the FaxObs_GetSecurityDescriptor operation request
type GetSecurityDescriptorRequest struct {
	// Id: A DWORD ([MS-DTYP] section 2.2.9) containing the identifier of the security descriptor
	// to request. The client MUST set this parameter to 0.
	ID uint32 `idl:"name:Id" json:"id"`
	// FaxSecurityDescriptor: A pointer to the address of the returned buffer containing
	// a FAX_SECURITY_DESCRIPTOR.
	FaxSecurityDescriptor []byte `idl:"name:FaxSecurityDescriptor;size_is:(, BufferSize);pointer:unique" json:"fax_security_descriptor"`
	// BufferSize: Pointer to a DWORD variable containing the number of bytes returned in
	// the FaxSecurityDescriptor buffer.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
}

func (o *GetSecurityDescriptorRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSecurityDescriptorOperation) *xxx_GetSecurityDescriptorOperation {
	if op == nil {
		op = &xxx_GetSecurityDescriptorOperation{}
	}
	if o == nil {
		return op
	}
	o.ID = op.ID
	o.FaxSecurityDescriptor = op.FaxSecurityDescriptor
	o.BufferSize = op.BufferSize
	return op
}

func (o *GetSecurityDescriptorRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSecurityDescriptorOperation) {
	if o == nil {
		return
	}
	o.ID = op.ID
	o.FaxSecurityDescriptor = op.FaxSecurityDescriptor
	o.BufferSize = op.BufferSize
}
func (o *GetSecurityDescriptorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSecurityDescriptorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSecurityDescriptorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSecurityDescriptorResponse structure represents the FaxObs_GetSecurityDescriptor operation response
type GetSecurityDescriptorResponse struct {
	// FaxSecurityDescriptor: A pointer to the address of the returned buffer containing
	// a FAX_SECURITY_DESCRIPTOR.
	FaxSecurityDescriptor []byte `idl:"name:FaxSecurityDescriptor;size_is:(, BufferSize);pointer:unique" json:"fax_security_descriptor"`
	// BufferSize: Pointer to a DWORD variable containing the number of bytes returned in
	// the FaxSecurityDescriptor buffer.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
	// Return: The FaxObs_GetSecurityDescriptor return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetSecurityDescriptorResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSecurityDescriptorOperation) *xxx_GetSecurityDescriptorOperation {
	if op == nil {
		op = &xxx_GetSecurityDescriptorOperation{}
	}
	if o == nil {
		return op
	}
	o.FaxSecurityDescriptor = op.FaxSecurityDescriptor
	o.BufferSize = op.BufferSize
	o.Return = op.Return
	return op
}

func (o *GetSecurityDescriptorResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSecurityDescriptorOperation) {
	if o == nil {
		return
	}
	o.FaxSecurityDescriptor = op.FaxSecurityDescriptor
	o.BufferSize = op.BufferSize
	o.Return = op.Return
}
func (o *GetSecurityDescriptorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSecurityDescriptorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSecurityDescriptorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetSecurityDescriptorOperation structure represents the FaxObs_SetSecurityDescriptor operation
type xxx_SetSecurityDescriptorOperation struct {
	FaxSecurityDescriptor []byte `idl:"name:FaxSecurityDescriptor;size_is:(BufferSize);pointer:unique" json:"fax_security_descriptor"`
	BufferSize            uint32 `idl:"name:BufferSize" json:"buffer_size"`
	Return                uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SetSecurityDescriptorOperation) OpNum() int { return 32 }

func (o *xxx_SetSecurityDescriptorOperation) OpName() string {
	return "/faxobs/v4/FaxObs_SetSecurityDescriptor"
}

func (o *xxx_SetSecurityDescriptorOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.FaxSecurityDescriptor != nil && o.BufferSize == 0 {
		o.BufferSize = uint32(len(o.FaxSecurityDescriptor))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecurityDescriptorOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// FaxSecurityDescriptor {in} (1:{pointer=unique, alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		if o.FaxSecurityDescriptor != nil || o.BufferSize > 0 {
			_ptr_FaxSecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.BufferSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.FaxSecurityDescriptor {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.FaxSecurityDescriptor[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.FaxSecurityDescriptor); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.FaxSecurityDescriptor, _ptr_FaxSecurityDescriptor); err != nil {
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
	// BufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecurityDescriptorOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// FaxSecurityDescriptor {in} (1:{pointer=unique, alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=BufferSize](uchar))
	{
		_ptr_FaxSecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.FaxSecurityDescriptor", sizeInfo[0])
			}
			o.FaxSecurityDescriptor = make([]byte, sizeInfo[0])
			for i1 := range o.FaxSecurityDescriptor {
				i1 := i1
				if err := w.ReadData(&o.FaxSecurityDescriptor[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_FaxSecurityDescriptor := func(ptr interface{}) { o.FaxSecurityDescriptor = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.FaxSecurityDescriptor, _s_FaxSecurityDescriptor, _ptr_FaxSecurityDescriptor); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// BufferSize {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.BufferSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecurityDescriptorOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecurityDescriptorOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetSecurityDescriptorOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetSecurityDescriptorRequest structure represents the FaxObs_SetSecurityDescriptor operation request
type SetSecurityDescriptorRequest struct {
	// FaxSecurityDescriptor: A pointer to a buffer containing a FAX_SECURITY_DESCRIPTOR
	// to be set.
	FaxSecurityDescriptor []byte `idl:"name:FaxSecurityDescriptor;size_is:(BufferSize);pointer:unique" json:"fax_security_descriptor"`
	// BufferSize: A DWORD ([MS-DTYP] section 2.2.9) containing the size, in bytes, of the
	// data pointed at by the FaxSecurityDescriptor parameter.
	BufferSize uint32 `idl:"name:BufferSize" json:"buffer_size"`
}

func (o *SetSecurityDescriptorRequest) xxx_ToOp(ctx context.Context, op *xxx_SetSecurityDescriptorOperation) *xxx_SetSecurityDescriptorOperation {
	if op == nil {
		op = &xxx_SetSecurityDescriptorOperation{}
	}
	if o == nil {
		return op
	}
	o.FaxSecurityDescriptor = op.FaxSecurityDescriptor
	o.BufferSize = op.BufferSize
	return op
}

func (o *SetSecurityDescriptorRequest) xxx_FromOp(ctx context.Context, op *xxx_SetSecurityDescriptorOperation) {
	if o == nil {
		return
	}
	o.FaxSecurityDescriptor = op.FaxSecurityDescriptor
	o.BufferSize = op.BufferSize
}
func (o *SetSecurityDescriptorRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetSecurityDescriptorRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSecurityDescriptorOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetSecurityDescriptorResponse structure represents the FaxObs_SetSecurityDescriptor operation response
type SetSecurityDescriptorResponse struct {
	// Return: The FaxObs_SetSecurityDescriptor return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetSecurityDescriptorResponse) xxx_ToOp(ctx context.Context, op *xxx_SetSecurityDescriptorOperation) *xxx_SetSecurityDescriptorOperation {
	if op == nil {
		op = &xxx_SetSecurityDescriptorOperation{}
	}
	if o == nil {
		return op
	}
	o.Return = op.Return
	return op
}

func (o *SetSecurityDescriptorResponse) xxx_FromOp(ctx context.Context, op *xxx_SetSecurityDescriptorOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetSecurityDescriptorResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetSecurityDescriptorResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetSecurityDescriptorOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSecurityDescriptorCountOperation structure represents the FaxObs_GetSecurityDescriptorCount operation
type xxx_GetSecurityDescriptorCountOperation struct {
	Count  uint32 `idl:"name:Count" json:"count"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSecurityDescriptorCountOperation) OpNum() int { return 33 }

func (o *xxx_GetSecurityDescriptorCountOperation) OpName() string {
	return "/faxobs/v4/FaxObs_GetSecurityDescriptorCount"
}

func (o *xxx_GetSecurityDescriptorCountOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSecurityDescriptorCountOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_GetSecurityDescriptorCountOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_GetSecurityDescriptorCountOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSecurityDescriptorCountOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Count {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Count); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSecurityDescriptorCountOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Count {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Count); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetSecurityDescriptorCountRequest structure represents the FaxObs_GetSecurityDescriptorCount operation request
type GetSecurityDescriptorCountRequest struct {
}

func (o *GetSecurityDescriptorCountRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSecurityDescriptorCountOperation) *xxx_GetSecurityDescriptorCountOperation {
	if op == nil {
		op = &xxx_GetSecurityDescriptorCountOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *GetSecurityDescriptorCountRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSecurityDescriptorCountOperation) {
	if o == nil {
		return
	}
}
func (o *GetSecurityDescriptorCountRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSecurityDescriptorCountRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSecurityDescriptorCountOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSecurityDescriptorCountResponse structure represents the FaxObs_GetSecurityDescriptorCount operation response
type GetSecurityDescriptorCountResponse struct {
	// Count: A pointer to a DWORD ([MS-DTYP] section 2.2.9) value containing on return
	// the number of security descriptors. On a successful return, the server MUST set this
	// parameter to a value of 1.
	Count uint32 `idl:"name:Count" json:"count"`
	// Return: The FaxObs_GetSecurityDescriptorCount return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetSecurityDescriptorCountResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSecurityDescriptorCountOperation) *xxx_GetSecurityDescriptorCountOperation {
	if op == nil {
		op = &xxx_GetSecurityDescriptorCountOperation{}
	}
	if o == nil {
		return op
	}
	o.Count = op.Count
	o.Return = op.Return
	return op
}

func (o *GetSecurityDescriptorCountResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSecurityDescriptorCountOperation) {
	if o == nil {
		return
	}
	o.Count = op.Count
	o.Return = op.Return
}
func (o *GetSecurityDescriptorCountResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSecurityDescriptorCountResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSecurityDescriptorCountOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AccessCheckOperation structure represents the FaxObs_AccessCheck operation
type xxx_AccessCheckOperation struct {
	AccessMask uint32 `idl:"name:AccessMask" json:"access_mask"`
	Access     uint32 `idl:"name:fAccess" json:"access"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_AccessCheckOperation) OpNum() int { return 34 }

func (o *xxx_AccessCheckOperation) OpName() string { return "/faxobs/v4/FaxObs_AccessCheck" }

func (o *xxx_AccessCheckOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AccessCheckOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// AccessMask {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.AccessMask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AccessCheckOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// AccessMask {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.AccessMask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AccessCheckOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AccessCheckOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// fAccess {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Access); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AccessCheckOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// fAccess {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Access); err != nil {
			return err
		}
	}
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AccessCheckRequest structure represents the FaxObs_AccessCheck operation request
type AccessCheckRequest struct {
	// AccessMask: A DWORD ([MS-DTYP] section 2.2.9) containing a set of bit flags that
	// define the fax access permissions to check for the fax client user account. This
	// parameter MUST be a bitwise OR combination of generic FaxObs access rights and specific
	// FaxObs access rights that are described in the following tables.
	//
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	|    GENERIC FAXOBS ACCESS     |                                                                                  |
	//	|            RIGHTS            |                                   DESCRIPTION                                    |
	//	|                              |                                                                                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_WRITE 0x00020001         | Includes the read-only rights that are granted by the following specific         |
	//	|                              | FaxObs access rights in combination with the standard access rights              |
	//	|                              | STANDARD_RIGHTS_WRITE: § FAX_JOB_SUBMIT                                          |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_READ 0x00020016          | Includes the read-only rights that are granted by the following specific         |
	//	|                              | FaxObs access rights in combination with the standard access rights              |
	//	|                              | STANDARD_RIGHTS_READ: § FAX_JOB_QUERY § FAX_CONFIG_QUERY § FAX_PORT_QUERY        |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_ALL_ACCESS 0x001F007F    | Includes the read-only rights that are granted by the following specific         |
	//	|                              | FaxObs access rights in combination with the standard access rights              |
	//	|                              | STANDARD_RIGHTS_ALL: § FAX_JOB_SUBMIT § FAX_JOB_QUERY § FAX_CONFIG_QUERY §       |
	//	|                              | FAX_CONFIG_SET § FAX_PORT_QUERY § FAX_PORT_SET § FAX_JOB_MANAGE                  |
	//	+------------------------------+----------------------------------------------------------------------------------+
	//
	// The generic FaxObs access rights are bitwise OR combination of specific FaxObs access
	// rights and standard access rights. For more information about the standard access
	// rights, see [MSDN-SAR].
	//
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	|    SPECIFIC FAXOBS ACCESS     |                                                                                  |
	//	|            RIGHTS             |                                   DESCRIPTION                                    |
	//	|                               |                                                                                  |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_JOB_SUBMIT 0x00000001     | The user can submit documents to be faxed. Example method: FaxObs_SendDocument   |
	//	|                               | (section 3.1.4.2.7).                                                             |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_JOB_QUERY 0x00000002      | The user can query information about submitted fax jobs. Example method:         |
	//	|                               | FaxObs_EnumJobs (section 3.1.4.2.9).                                             |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_CONFIG_QUERY 0x00000004   | The user can query the fax server configuration. Example method:                 |
	//	|                               | FaxObs_GetConfiguration (section 3.1.4.2.24).                                    |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_CONFIG_SET 0x00000008     | The user can change the fax server configuration. Example method:                |
	//	|                               | FaxObs_SetConfiguration (section 3.1.4.2.25).                                    |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_PORT_QUERY 0x00000010     | The user can query information about the fax ports (devices) installed on the    |
	//	|                               | fax server. Example method: FaxObs_EnumPorts (section 3.1.4.2.15).               |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_PORT_SET 0x00000020       | The user can change the configuration of the fax ports (devices) installed on    |
	//	|                               | the fax server. Example method: FaxObs_SetPort (section 3.1.4.2.17).             |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| FAX_JOB_MANAGE 0x00000040     | The user can pause, resume, and cancel submitted fax jobs. Example method:       |
	//	|                               | FaxObs_SetJob (section 3.1.4.2.11).                                              |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	AccessMask uint32 `idl:"name:AccessMask" json:"access_mask"`
}

func (o *AccessCheckRequest) xxx_ToOp(ctx context.Context, op *xxx_AccessCheckOperation) *xxx_AccessCheckOperation {
	if op == nil {
		op = &xxx_AccessCheckOperation{}
	}
	if o == nil {
		return op
	}
	o.AccessMask = op.AccessMask
	return op
}

func (o *AccessCheckRequest) xxx_FromOp(ctx context.Context, op *xxx_AccessCheckOperation) {
	if o == nil {
		return
	}
	o.AccessMask = op.AccessMask
}
func (o *AccessCheckRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AccessCheckRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AccessCheckOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AccessCheckResponse structure represents the FaxObs_AccessCheck operation response
type AccessCheckResponse struct {
	// fAccess: A pointer to a BOOL variable to receive on successful return the access
	// check return value. A TRUE value indicates that access is allowed. A FALSE value
	// indicates that access is denied.
	Access uint32 `idl:"name:fAccess" json:"access"`
	// Return: The FaxObs_AccessCheck return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AccessCheckResponse) xxx_ToOp(ctx context.Context, op *xxx_AccessCheckOperation) *xxx_AccessCheckOperation {
	if op == nil {
		op = &xxx_AccessCheckOperation{}
	}
	if o == nil {
		return op
	}
	o.Access = op.Access
	o.Return = op.Return
	return op
}

func (o *AccessCheckResponse) xxx_FromOp(ctx context.Context, op *xxx_AccessCheckOperation) {
	if o == nil {
		return
	}
	o.Access = op.Access
	o.Return = op.Return
}
func (o *AccessCheckResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AccessCheckResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AccessCheckOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
