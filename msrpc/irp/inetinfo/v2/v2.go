package inetinfo

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dtyp "github.com/oiweiwei/go-msrpc/msrpc/dtyp"
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
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "irp"
)

var (
	// Syntax UUID
	InetinfoSyntaxUUID = &uuid.UUID{TimeLow: 0x82ad4280, TimeMid: 0x36b, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x97, ClockSeqLow: 0x2c, Node: [6]uint8{0x0, 0xaa, 0x0, 0x68, 0x87, 0xb0}}
	// Syntax ID
	InetinfoSyntaxV2_0 = &dcerpc.SyntaxID{IfUUID: InetinfoSyntaxUUID, IfVersionMajor: 2, IfVersionMinor: 0}
)

// inetinfo interface.
type InetinfoClient interface {

	// The R_InetInfoGetVersion method is called by the client. In response, the server
	// returns its version information.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2 or [MS-ERREF]
	// section 2.3.1.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	//
	// The value returned in pdwVersion SHOULD<9> correspond to the version of the Internet
	// protocol servers managed by the Internet Information Services (IIS) Inetinfo Remote
	// Protocol server.
	GetVersion(context.Context, *GetVersionRequest, ...dcerpc.CallOption) (*GetVersionResponse, error)

	// The R_InetInfoGetAdminInformation method is called by the client. In response, the
	// server retrieves configuration data for the specified Internet protocol server.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2 or [MS-ERREF]
	// section 2.3.1. The most common error codes are listed in the following table.
	//
	//	+-------------------------------------+--------------------------------------------+
	//	|               RETURN                |                                            |
	//	|             VALUE/CODE              |                DESCRIPTION                 |
	//	|                                     |                                            |
	//	+-------------------------------------+--------------------------------------------+
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND     | The system cannot find the specified file. |
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000426 ERROR_SERVICE_NOT_ACTIVE | The service is not running.                |
	//	+-------------------------------------+--------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	GetAdminInformation(context.Context, *GetAdminInformationRequest, ...dcerpc.CallOption) (*GetAdminInformationResponse, error)

	// The R_InetInfoGetSites method is called by the client. In response, the server retrieves
	// a list of server instances for the specified Internet protocol server.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2 or [MS-ERREF]
	// section 2.3.1. The most common error codes are listed in the following table.
	//
	//	+-------------------------------------+-----------------------------+
	//	|               RETURN                |                             |
	//	|             VALUE/CODE              |         DESCRIPTION         |
	//	|                                     |                             |
	//	+-------------------------------------+-----------------------------+
	//	+-------------------------------------+-----------------------------+
	//	| 0x00000426 ERROR_SERVICE_NOT_ACTIVE | The service is not running. |
	//	+-------------------------------------+-----------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	GetSites(context.Context, *GetSitesRequest, ...dcerpc.CallOption) (*GetSitesResponse, error)

	// The R_InetInfoSetAdminInformation method is called by the client. In response, the
	// server sets configurable properties for the specified Internet protocol server.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2, or [MS-ERREF]
	// section 2.3.1. The most common error codes are listed in the following table.
	//
	//	+-------------------------------------+--------------------------------------------+
	//	|               RETURN                |                                            |
	//	|             VALUE/CODE              |                DESCRIPTION                 |
	//	|                                     |                                            |
	//	+-------------------------------------+--------------------------------------------+
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND     | The system cannot find the specified file. |
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED      | Access is denied.                          |
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000426 ERROR_SERVICE_NOT_ACTIVE | The service is not running.                |
	//	+-------------------------------------+--------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	SetAdminInformation(context.Context, *SetAdminInformationRequest, ...dcerpc.CallOption) (*SetAdminInformationResponse, error)

	// The R_InetInfoGetGlobalAdminInformation method is called by the client. In response,
	// the server retrieves configuration data shared by all Internet protocol servers.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2 or [MS-ERREF]
	// section 2.3.1.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	GetGlobalAdminInformation(context.Context, *GetGlobalAdminInformationRequest, ...dcerpc.CallOption) (*GetGlobalAdminInformationResponse, error)

	// The R_InetInfoSetGlobalAdminInformation assigns global settings for all Internet
	// protocol servers present on the host system.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2 or [MS-ERREF]
	// section 2.3.1. The most common error codes are listed in the following table.
	//
	//	+--------------------------------+-------------------------------+
	//	|             RETURN             |                               |
	//	|           VALUE/CODE           |          DESCRIPTION          |
	//	|                                |                               |
	//	+--------------------------------+-------------------------------+
	//	+--------------------------------+-------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | Access is denied.             |
	//	+--------------------------------+-------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED | The request is not supported. |
	//	+--------------------------------+-------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	SetGlobalAdminInformation(context.Context, *SetGlobalAdminInformationRequest, ...dcerpc.CallOption) (*SetGlobalAdminInformationResponse, error)

	// The R_InetInfoQueryStatistics method is called by the client. In response, the server
	// retrieves statistical data for the specified Internet protocol server.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2 or [MS-ERREF]
	// section 2.3.1. The following table describes possible error code values.
	//
	//	+--------------------------------+---------------------------------------+
	//	|             RETURN             |                                       |
	//	|           VALUE/CODE           |              DESCRIPTION              |
	//	|                                |                                       |
	//	+--------------------------------+---------------------------------------+
	//	+--------------------------------+---------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL | The system call level is not correct. |
	//	+--------------------------------+---------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	QueryStatistics(context.Context, *QueryStatisticsRequest, ...dcerpc.CallOption) (*QueryStatisticsResponse, error)

	// The R_InetInfoClearStatistics is called by the client. In response, the server resets
	// the statistical data maintained by the specified Internet protocol server.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2 or [MS-ERREF]
	// section 2.3.1. The most common error codes are listed in the following table.
	//
	//	+------------------------------------+-------------------------------+
	//	|               RETURN               |                               |
	//	|             VALUE/CODE             |          DESCRIPTION          |
	//	|                                    |                               |
	//	+------------------------------------+-------------------------------+
	//	+------------------------------------+-------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.             |
	//	+------------------------------------+-------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The request is not supported. |
	//	+------------------------------------+-------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The parameter is incorrect.   |
	//	+------------------------------------+-------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	ClearStatistics(context.Context, *ClearStatisticsRequest, ...dcerpc.CallOption) (*ClearStatisticsResponse, error)

	// The R_InetInfoFlushMemoryCache method is called by the client. In response, the server
	// flushes data from the internal caches of the specified Internet protocol server.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2 or  [MS-ERREF]
	// section 2.3.1. The most common error codes are listed in the following table.
	//
	//	+--------------------------------+-------------------+
	//	|             RETURN             |                   |
	//	|           VALUE/CODE           |    DESCRIPTION    |
	//	|                                |                   |
	//	+--------------------------------+-------------------+
	//	+--------------------------------+-------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | Access is denied. |
	//	+--------------------------------+-------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	FlushMemoryCache(context.Context, *FlushMemoryCacheRequest, ...dcerpc.CallOption) (*FlushMemoryCacheResponse, error)

	// The R_InetInfoGetServerCapabilities method is called by the client. In response,
	// the server returns information on the features of the Internet protocol servers and
	// the host operating system.
	//
	// Return Values:  The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2 or [MS-ERREF]
	// section 2.3.1.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	GetServerCapabilities(context.Context, *GetServerCapabilitiesRequest, ...dcerpc.CallOption) (*GetServerCapabilitiesResponse, error)

	// The R_W3QueryStatistics2 method is called by the client. In response, the server
	// returns statistical data from the HTTP server.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2 or [MS-ERREF]
	// section 2.3.1. The most common error codes are listed in the following table.
	//
	//	+-------------------------------------+--------------------------------------------+
	//	|               RETURN                |                                            |
	//	|             VALUE/CODE              |                DESCRIPTION                 |
	//	|                                     |                                            |
	//	+-------------------------------------+--------------------------------------------+
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND     | The system cannot find the specified file. |
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL      | The system call level is not correct.      |
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000426 ERROR_SERVICE_NOT_ACTIVE | The service is not running.                |
	//	+-------------------------------------+--------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	W3QueryStatistics2(context.Context, *W3QueryStatistics2Request, ...dcerpc.CallOption) (*W3QueryStatistics2Response, error)

	// The R_W3ClearStatistics2 method is called by the client. In response, the server
	// resets statistical data for the HTTP server.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2 or [MS-ERREF]
	// section 2.3.1. The most common error codes are listed in the following table.
	//
	//	+-------------------------------------+--------------------------------------------+
	//	|               RETURN                |                                            |
	//	|             VALUE/CODE              |                DESCRIPTION                 |
	//	|                                     |                                            |
	//	+-------------------------------------+--------------------------------------------+
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND     | The system cannot find the specified file. |
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED      | Access is denied.                          |
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000426 ERROR_SERVICE_NOT_ACTIVE | The service is not running.                |
	//	+-------------------------------------+--------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	W3ClearStatistics2(context.Context, *W3ClearStatistics2Request, ...dcerpc.CallOption) (*W3ClearStatistics2Response, error)

	// The R_FtpQueryStatistics2 method is called by the client. In response, the server
	// returns statistical data from the FTP server.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2 or [MS-ERREF]
	// section 2.3.1. The most common error codes are listed in the following table.
	//
	//	+-------------------------------------+--------------------------------------------+
	//	|               RETURN                |                                            |
	//	|             VALUE/CODE              |                DESCRIPTION                 |
	//	|                                     |                                            |
	//	+-------------------------------------+--------------------------------------------+
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND     | The system cannot find the specified file. |
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL      | The system call level is not correct.      |
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000426 ERROR_SERVICE_NOT_ACTIVE | The service is not running.                |
	//	+-------------------------------------+--------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	FTPQueryStatistics2(context.Context, *FTPQueryStatistics2Request, ...dcerpc.CallOption) (*FTPQueryStatistics2Response, error)

	// The R_FtpClearStatistics2 method is called by the client. In response, the server
	// resets statistical data for the FTP server.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2 or [MS-ERREF]
	// section 2.3.1. The most common error codes are listed in the following table.
	//
	//	+-------------------------------------+--------------------------------------------+
	//	|               RETURN                |                                            |
	//	|             VALUE/CODE              |                DESCRIPTION                 |
	//	|                                     |                                            |
	//	+-------------------------------------+--------------------------------------------+
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND     | The system cannot find the specified file. |
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED      | Access is denied.                          |
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000426 ERROR_SERVICE_NOT_ACTIVE | The service is not running.                |
	//	+-------------------------------------+--------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	FTPClearStatistics2(context.Context, *FTPClearStatistics2Request, ...dcerpc.CallOption) (*FTPClearStatistics2Response, error)

	// The R_IISEnumerateUsers method is called by the client. In response, the server returns
	// a list of clients connected to the specified Internet protocol server.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2 or[MS-ERREF]
	// section 2.3.1. The most common error codes are listed in the following table.
	//
	//	+-------------------------------------+--------------------------------------------+
	//	|               RETURN                |                                            |
	//	|             VALUE/CODE              |                DESCRIPTION                 |
	//	|                                     |                                            |
	//	+-------------------------------------+--------------------------------------------+
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND     | The system cannot find the specified file. |
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED      | Access is denied.                          |
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000426 ERROR_SERVICE_NOT_ACTIVE | The service is not running.                |
	//	+-------------------------------------+--------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	IISEnumerateUsers(context.Context, *IISEnumerateUsersRequest, ...dcerpc.CallOption) (*IISEnumerateUsersResponse, error)

	// The R_IISDisconnectUser method is called by the client. In response, the server disconnects
	// the specified user from the specified Internet protocol server.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2 or[MS-ERREF]
	// section 2.3.1. The most common error codes are listed in the following table.
	//
	//	+-------------------------------------+--------------------------------------------+
	//	|               RETURN                |                                            |
	//	|             VALUE/CODE              |                DESCRIPTION                 |
	//	|                                     |                                            |
	//	+-------------------------------------+--------------------------------------------+
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND     | The system cannot find the specified file. |
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED      | Access is denied.                          |
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x00000426 ERROR_SERVICE_NOT_ACTIVE | The service is not running.                |
	//	+-------------------------------------+--------------------------------------------+
	//	| 0x000008AD NERR_UserNotFound        | The user name could not be found.          |
	//	+-------------------------------------+--------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	IISDisconnectUser(context.Context, *IISDisconnectUserRequest, ...dcerpc.CallOption) (*IISDisconnectUserResponse, error)

	// Opnum16NotUsedOnWire operation.
	// Opnum16NotUsedOnWire

	// Opnum17NotUsedOnWire operation.
	// Opnum17NotUsedOnWire

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// CapFlags structure represents INET_INFO_CAP_FLAGS RPC structure.
//
// The INET_INFO_CAP_FLAGS structure contains information on features that are available
// for a protocol server.
//
// This type is declared as follows:
type CapFlags struct {
	// Flag:  A value that indicates the features supported by the current running instance
	// of the Internet protocol server implementation. The possible values for this member
	// result from a bitwise OR of zero or more of the flags defined in the following table.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                         |                                                                                  |
	//	|                  VALUE                  |                                     MEANING                                      |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_ODBC_LOGGING 0x00000001        | The Internet protocol server supports the Open Database Connectivity (ODBC) log  |
	//	|                                         | format feature.                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_FILE_LOGGING 0x00000002        | The Internet protocol server supports the file system logging feature.           |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_VIRTUAL_SERVER 0x00000004      | The Internet protocol server supports multiple instances of the protocol server  |
	//	|                                         | network endpoint (1) aliases.                                                    |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_BW_THROTTLING 0x00000008       | The Internet protocol server supports network bandwidth throttling.              |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_IP_ACCESS_CHECK 0x00000010     | The Internet protocol server supports blocking client connections using IP       |
	//	|                                         | restrictions.                                                                    |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_MAX_CONNECTIONS 0x00000020     | The Internet protocol server supports connection limiting.                       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_10_CONNECTION_LIMIT 0x00000040 | The Internet protocol server supports a limit of 10 concurrent connections.      |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_MULTIPLE_INSTANCE 0x00000080   | The Internet protocol server supports multiple instances.                        |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_SSL_SUPPORT 0x00000100         | The Internet protocol server supports the SSL protocol.                          |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_OPERATORS_LIST 0x00000200      | The Internet protocol server supports administrative operations by identities    |
	//	|                                         | other than Windows operating system administrators.                              |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_FP_INSTALLED 0x00000400        | Front Page Server Extensions are installed on the server.                        |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_CPU_AUDITING 0x00000800        | The Internet protocol server supports CPU limits.                                |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_SERVER_COMPRESSION 0x00001000  | The Internet protocol server supports compression of network data.               |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_DAV 0x00002000                 | The Internet protocol server supports the WebDAV protocol.                       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_DIGEST_SUPPORT 0x00004000      | The Internet protocol server supports the Digest Authentication Protocol.        |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_NT_CERTMAP_SUPPORT 0x00008000  | The Internet protocol server supports mapping client certificates to Windows     |
	//	|                                         | user accounts.                                                                   |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_POOLED_OOP 0x00010000          | The Internet protocol server supports running a pool of applications in a        |
	//	|                                         | separate process.                                                                |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	Flag uint32 `idl:"name:Flag" json:"flag"`
	// Mask:  A value that indicates the capabilities that can be enabled for the protocol
	// servers in the current implementation. The possible values for this member result
	// from a bitwise OR operation of zero or more of the following flags.
	//
	// Note  The Mask value contains all the capabilities that the current version of the
	// protocol server implementation can support. The Flag value indicates the features
	// that the current running server instance does support.The server sets the mask value
	// to a bitwise OR of all the flags it knows about. The server then sets the flags to
	// the bitwise OR of the features supported for the current platform, a subset of those
	// set in the mask field. A given version of the software reports the same mask values,
	// but might support different flags values depending on the operating system type.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                         |                                                                                  |
	//	|                  VALUE                  |                                     MEANING                                      |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_ODBC_LOGGING 0x00000001        | The Internet protocol server supports the Open Database Connectivity (ODBC) log  |
	//	|                                         | format feature.                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_FILE_LOGGING 0x00000002        | The Internet protocol server supports the file system logging feature.           |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_VIRTUAL_SERVER 0x00000004      | The Internet protocol server supports multiple instances of the protocol server  |
	//	|                                         | network endpoint (2) aliases.                                                    |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_BW_THROTTLING 0x00000008       | The Internet protocol server supports network bandwidth throttling.              |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_IP_ACCESS_CHECK 0x00000010     | The Internet protocol server supports blocking client connections by using IP    |
	//	|                                         | restrictions.                                                                    |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_MAX_CONNECTIONS 0x00000020     | The Internet protocol server supports connection limiting.                       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_10_CONNECTION_LIMIT 0x00000040 | The Internet protocol server supports a limit of 10 concurrent connections.      |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_MULTIPLE_INSTANCE 0x00000080   | The Internet protocol server supports multiple instances.                        |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_SSL_SUPPORT 0x00000100         | The Internet protocol server supports the SSL protocol.                          |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_OPERATORS_LIST 0x00000200      | The Internet protocol server supports administrative operations by identities    |
	//	|                                         | other than Windows operating system administrators.                              |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_FP_INSTALLED 0x00000400        | Front Page Server Extensions are installed on the server.                        |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_CPU_AUDITING 0x00000800        | The Internet protocol server supports CPU limits.                                |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_SERVER_COMPRESSION 0x00001000  | The Internet protocol server supports compression of network data.               |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_DAV 0x00002000                 | The Internet protocol server supports the WebDAV protocol.                       |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_DIGEST_SUPPORT 0x00004000      | The Internet protocol server supports the Digest Authentication Protocol.        |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_NT_CERTMAP_SUPPORT 0x00008000  | The Internet protocol server supports mapping client certificates to Windows     |
	//	|                                         | user accounts.                                                                   |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| IIS_CAP1_POOLED_OOP 0x00010000          | The Internet protocol server supports running a pool of applications in a        |
	//	|                                         | separate process.                                                                |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	Mask uint32 `idl:"name:Mask" json:"mask"`
}

func (o *CapFlags) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CapFlags) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Flag); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	return nil
}
func (o *CapFlags) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flag); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	return nil
}

// Capabilities structure represents INET_INFO_CAPABILITIES_STRUCT RPC structure.
//
// The INET_INFO_CAPABILITIES_STRUCT structure specifies the features supported by an
// Internet protocol server implementation.
//
// This type is declared as follows:
type Capabilities struct {
	// CapVersion:  The version number of this structure. MUST be 1.
	CapVersion uint32 `idl:"name:CapVersion" json:"cap_version"`
	// ProductType:  The value that indicates the Windows operating system product type
	// hosting the implementation. This field MUST be set to one of the following values.
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                                        |                                                                                  |
	//	|                 VALUE                  |                                     MEANING                                      |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| INET_INFO_PRODUCT_NTSERVER 0x00000001  | The operating system product type is a Windows server.                           |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| INET_INFO_PRODUCT_NTWKSTA 0x00000002   | The operating system product type is a Windows client or Windows NT Workstation  |
	//	|                                        | operating system.                                                                |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| INET_INFO_PRODUCT_UNKNOWN 0xFFFFFFFF   | The operating system product type is unknown.                                    |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| INET_INFO_PRODUCT_WINDOWS95 0x00000003 | The operating system product type is Windows 95 operating system.                |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	ProductType uint32 `idl:"name:ProductType" json:"product_type"`
	// MajorVersion:   The major version number of the Internet Information Services (IIS)
	// Inetinfo Remote Protocol server.
	MajorVersion uint32 `idl:"name:MajorVersion" json:"major_version"`
	// MinorVersion:  The minor version number of the Internet Information Services (IIS)
	// Inetinfo Remote Protocol server.
	MinorVersion uint32 `idl:"name:MinorVersion" json:"minor_version"`
	// BuildNumber:  The build number of the version of the Windows operating system running
	// on the Internet Information Services (IIS) Inetinfo Remote Protocol server.
	BuildNumber uint32 `idl:"name:BuildNumber" json:"build_number"`
	// NumCapFlags:  The number of INET_INFO_CAP_FLAGS structures pointed to by the CapFlags
	// member. MUST be 1.
	CapFlagsLength uint32 `idl:"name:NumCapFlags" json:"cap_flags_length"`
	// CapFlags:  The pointer to an array of INET_INFO_CAP_FLAGS structures that defines
	// the server's capabilities.
	CapFlags []*CapFlags `idl:"name:CapFlags;size_is:(NumCapFlags)" json:"cap_flags"`
}

func (o *Capabilities) xxx_PreparePayload(ctx context.Context) error {
	if o.CapFlags != nil && o.CapFlagsLength == 0 {
		o.CapFlagsLength = uint32(len(o.CapFlags))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Capabilities) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.CapVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.ProductType); err != nil {
		return err
	}
	if err := w.WriteData(o.MajorVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.MinorVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.BuildNumber); err != nil {
		return err
	}
	if err := w.WriteData(o.CapFlagsLength); err != nil {
		return err
	}
	if o.CapFlags != nil || o.CapFlagsLength > 0 {
		_ptr_CapFlags := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.CapFlagsLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.CapFlags {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.CapFlags[i1] != nil {
					if err := o.CapFlags[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&CapFlags{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.CapFlags); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&CapFlags{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.CapFlags, _ptr_CapFlags); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Capabilities) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.CapVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.ProductType); err != nil {
		return err
	}
	if err := w.ReadData(&o.MajorVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.MinorVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.BuildNumber); err != nil {
		return err
	}
	if err := w.ReadData(&o.CapFlagsLength); err != nil {
		return err
	}
	_ptr_CapFlags := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.CapFlagsLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.CapFlagsLength)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.CapFlags", sizeInfo[0])
		}
		o.CapFlags = make([]*CapFlags, sizeInfo[0])
		for i1 := range o.CapFlags {
			i1 := i1
			if o.CapFlags[i1] == nil {
				o.CapFlags[i1] = &CapFlags{}
			}
			if err := o.CapFlags[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_CapFlags := func(ptr interface{}) { o.CapFlags = *ptr.(*[]*CapFlags) }
	if err := w.ReadPointer(&o.CapFlags, _s_CapFlags, _ptr_CapFlags); err != nil {
		return err
	}
	return nil
}

// InetLogConfiguration structure represents INET_LOG_CONFIGURATION RPC structure.
//
// The INET_LOG_CONFIGURATION structure contains configuration information for Internet
// protocol server logging.
//
// This type is declared as follows:
type InetLogConfiguration struct {
	// inetLogType:  A 32-bit integer that specifies the type of log to be written. This
	// field MUST be set to one of the following values.
	//
	//	+------------------------------+-------------------------------------------------------------------+
	//	|                              |                                                                   |
	//	|            VALUE             |                              MEANING                              |
	//	|                              |                                                                   |
	//	+------------------------------+-------------------------------------------------------------------+
	//	+------------------------------+-------------------------------------------------------------------+
	//	| INET_LOG_DISABLED 0x00000000 | Logging is disabled.                                              |
	//	+------------------------------+-------------------------------------------------------------------+
	//	| INET_LOG_TO_FILE 0x00000001  | The log is written to a file.                                     |
	//	+------------------------------+-------------------------------------------------------------------+
	//	| INET_LOG_TO_SQL 0x00000002   | The log is written to a Structured Query Language (SQL) database. |
	//	+------------------------------+-------------------------------------------------------------------+
	//	| INET_LOG_INVALID 0xFFFFFFFF  | The log is not valid.                                             |
	//	+------------------------------+-------------------------------------------------------------------+
	InetLogType uint32 `idl:"name:inetLogType" json:"inet_log_type"`
	// ilPeriod:  Specifies the periodicity of Internet protocol server logging. This field
	// MUST be set to one of the following values.
	//
	//	+------------------------------------+------------------------------+
	//	|                                    |                              |
	//	|               VALUE                |           MEANING            |
	//	|                                    |                              |
	//	+------------------------------------+------------------------------+
	//	+------------------------------------+------------------------------+
	//	| INET_LOG_PERIOD_NONE 0x00000000    | There is no log period.      |
	//	+------------------------------------+------------------------------+
	//	| INET_LOG_PERIOD_DAILY 0x00000001   | The log period is daily.     |
	//	+------------------------------------+------------------------------+
	//	| INET_LOG_PERIOD_WEEKLY 0x00000002  | The log period is weekly.    |
	//	+------------------------------------+------------------------------+
	//	| INET_LOG_PERIOD_MONTHLY 0x00000003 | The log period is monthly.   |
	//	+------------------------------------+------------------------------+
	//	| INET_LOG_PERIOD_HOURLY 0x00000004  | The log period is hourly.    |
	//	+------------------------------------+------------------------------+
	//	| INET_LOG_PERIOD_YEARLY 0x00000005  | The log period is yearly.    |
	//	+------------------------------------+------------------------------+
	//	| INET_LOG_PERIOD_INVALID 0xFFFFFFFF | The log period is not valid. |
	//	+------------------------------------+------------------------------+
	Period uint32 `idl:"name:ilPeriod" json:"period"`
	// rgchLogFileDirectory:  A null-terminated string that specifies the destination of
	// the Internet protocol server log.
	LogFileDirectory []uint16 `idl:"name:rgchLogFileDirectory" json:"log_file_directory"`
	// cbSizeForTruncation:  Specifies the maximum size in bytes for each log file.
	LengthForTruncation uint32 `idl:"name:cbSizeForTruncation" json:"length_for_truncation"`
	// rgchDataSource:  A null-terminated string that specifies the Open Database Connectivity
	// (ODBC) data source name to which the Internet protocol server log is to be written.
	DataSource []uint16 `idl:"name:rgchDataSource" json:"data_source"`
	// rgchTableName:  A null-terminated string that specifies the name of the table on
	// rgchDataSource to which the Internet protocol server log is to be written.
	TableName []uint16 `idl:"name:rgchTableName" json:"table_name"`
	// rgchUserName:  A null-terminated string that specifies the name of the user for the
	// ODBC connection.
	UserName []uint16 `idl:"name:rgchUserName" json:"user_name"`
	// rgchPassword:  A null-terminated string that specifies the password associated with
	// the rgchUserName user name.
	Password []uint16 `idl:"name:rgchPassword" json:"password"`
}

func (o *InetLogConfiguration) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *InetLogConfiguration) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.InetLogType); err != nil {
		return err
	}
	if err := w.WriteData(o.Period); err != nil {
		return err
	}
	for i1 := range o.LogFileDirectory {
		i1 := i1
		if uint64(i1) >= 260 {
			break
		}
		if err := w.WriteData(o.LogFileDirectory[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.LogFileDirectory); uint64(i1) < 260; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LengthForTruncation); err != nil {
		return err
	}
	for i1 := range o.DataSource {
		i1 := i1
		if uint64(i1) >= 260 {
			break
		}
		if err := w.WriteData(o.DataSource[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.DataSource); uint64(i1) < 260; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.TableName {
		i1 := i1
		if uint64(i1) >= 30 {
			break
		}
		if err := w.WriteData(o.TableName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.TableName); uint64(i1) < 30; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.UserName {
		i1 := i1
		if uint64(i1) >= 257 {
			break
		}
		if err := w.WriteData(o.UserName[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.UserName); uint64(i1) < 257; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	for i1 := range o.Password {
		i1 := i1
		if uint64(i1) >= 257 {
			break
		}
		if err := w.WriteData(o.Password[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Password); uint64(i1) < 257; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *InetLogConfiguration) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.InetLogType); err != nil {
		return err
	}
	if err := w.ReadData(&o.Period); err != nil {
		return err
	}
	o.LogFileDirectory = make([]uint16, 260)
	for i1 := range o.LogFileDirectory {
		i1 := i1
		if err := w.ReadData(&o.LogFileDirectory[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.LengthForTruncation); err != nil {
		return err
	}
	o.DataSource = make([]uint16, 260)
	for i1 := range o.DataSource {
		i1 := i1
		if err := w.ReadData(&o.DataSource[i1]); err != nil {
			return err
		}
	}
	o.TableName = make([]uint16, 30)
	for i1 := range o.TableName {
		i1 := i1
		if err := w.ReadData(&o.TableName[i1]); err != nil {
			return err
		}
	}
	o.UserName = make([]uint16, 257)
	for i1 := range o.UserName {
		i1 := i1
		if err := w.ReadData(&o.UserName[i1]); err != nil {
			return err
		}
	}
	o.Password = make([]uint16, 257)
	for i1 := range o.Password {
		i1 := i1
		if err := w.ReadData(&o.Password[i1]); err != nil {
			return err
		}
	}
	return nil
}

// IPsecEntry structure represents INET_INFO_IP_SEC_ENTRY RPC structure.
//
// The INET_INFO_IP_SEC_ENTRY structure contains Internet Protocol Security (IPv4) entries.
//
// This type is declared as follows:
type IPsecEntry struct {
	// dwMask:  The subnet mask. Data is stored in network byte order.
	Mask uint32 `idl:"name:dwMask" json:"mask"`
	// dwNetwork:  The IP address. Data is stored in network byte order.
	Network uint32 `idl:"name:dwNetwork" json:"network"`
}

func (o *IPsecEntry) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPsecEntry) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	if err := w.WriteData(o.Network); err != nil {
		return err
	}
	return nil
}
func (o *IPsecEntry) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	if err := w.ReadData(&o.Network); err != nil {
		return err
	}
	return nil
}

// IPsecList structure represents INET_INFO_IP_SEC_LIST RPC structure.
//
// The INET_INFO_IP_SEC_LIST structure contains a list of INET_INFO_IP_SEC_ENTRY entries.
//
// This type is declared as follows:
type IPsecList struct {
	// cEntries:  The number of entries contained in the list.
	EntriesCount uint32 `idl:"name:cEntries" json:"entries_count"`
	// aIPSecEntry:  An array of INET_INFO_IP_SEC_ENTRY entries.
	IPsecEntry []*IPsecEntry `idl:"name:aIPSecEntry;size_is:(cEntries)" json:"ipsec_entry"`
}

func (o *IPsecList) xxx_PreparePayload(ctx context.Context) error {
	if o.IPsecEntry != nil && o.EntriesCount == 0 {
		o.EntriesCount = uint32(len(o.IPsecEntry))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *IPsecList) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.EntriesCount)
	return []uint64{
		dimSize1,
	}
}
func (o *IPsecList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesCount); err != nil {
		return err
	}
	for i1 := range o.IPsecEntry {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.IPsecEntry[i1] != nil {
			if err := o.IPsecEntry[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&IPsecEntry{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.IPsecEntry); uint64(i1) < sizeInfo[0]; i1++ {
		if err := (&IPsecEntry{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *IPsecList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.EntriesCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.EntriesCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.IPsecEntry", sizeInfo[0])
	}
	o.IPsecEntry = make([]*IPsecEntry, sizeInfo[0])
	for i1 := range o.IPsecEntry {
		i1 := i1
		if o.IPsecEntry[i1] == nil {
			o.IPsecEntry[i1] = &IPsecEntry{}
		}
		if err := o.IPsecEntry[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// VirtualRootEntry structure represents INET_INFO_VIRTUAL_ROOT_ENTRY RPC structure.
//
// The INET_INFO_VIRTUAL_ROOT_ENTRY structure contains data describing a virtual root
// for the Internet protocol server instance.
//
// This type is declared as follows:
//
// pszRoot:  The virtual root name.
//
// pszAddress:  The optional IP address.
//
// pszDirectory:  The physical directory.
//
// dwMask:  The Access Mask for this virtual root.
//
// pszAccountName:  The account to connect as.
//
// AccountPassword:  Null-terminated WCHAR array containing the password for pszAccountName.<5>
//
// dwError:  The error code stored for the virtual root.
type VirtualRootEntry struct {
	Root            string   `idl:"name:pszRoot;string" json:"root"`
	Address         string   `idl:"name:pszAddress;string" json:"address"`
	Directory       string   `idl:"name:pszDirectory;string" json:"directory"`
	Mask            uint32   `idl:"name:dwMask" json:"mask"`
	AccountName     string   `idl:"name:pszAccountName;string" json:"account_name"`
	AccountPassword []uint16 `idl:"name:AccountPassword" json:"account_password"`
	Error           uint32   `idl:"name:dwError" json:"error"`
}

func (o *VirtualRootEntry) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *VirtualRootEntry) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Root != "" {
		_ptr_pszRoot := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Root); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Root, _ptr_pszRoot); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Address != "" {
		_ptr_pszAddress := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Address); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Address, _ptr_pszAddress); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Directory != "" {
		_ptr_pszDirectory := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Directory); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Directory, _ptr_pszDirectory); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Mask); err != nil {
		return err
	}
	if o.AccountName != "" {
		_ptr_pszAccountName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.AccountName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.AccountName, _ptr_pszAccountName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	for i1 := range o.AccountPassword {
		i1 := i1
		if uint64(i1) >= 257 {
			break
		}
		if err := w.WriteData(o.AccountPassword[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.AccountPassword); uint64(i1) < 257; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Error); err != nil {
		return err
	}
	return nil
}
func (o *VirtualRootEntry) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pszRoot := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Root); err != nil {
			return err
		}
		return nil
	})
	_s_pszRoot := func(ptr interface{}) { o.Root = *ptr.(*string) }
	if err := w.ReadPointer(&o.Root, _s_pszRoot, _ptr_pszRoot); err != nil {
		return err
	}
	_ptr_pszAddress := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Address); err != nil {
			return err
		}
		return nil
	})
	_s_pszAddress := func(ptr interface{}) { o.Address = *ptr.(*string) }
	if err := w.ReadPointer(&o.Address, _s_pszAddress, _ptr_pszAddress); err != nil {
		return err
	}
	_ptr_pszDirectory := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Directory); err != nil {
			return err
		}
		return nil
	})
	_s_pszDirectory := func(ptr interface{}) { o.Directory = *ptr.(*string) }
	if err := w.ReadPointer(&o.Directory, _s_pszDirectory, _ptr_pszDirectory); err != nil {
		return err
	}
	if err := w.ReadData(&o.Mask); err != nil {
		return err
	}
	_ptr_pszAccountName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.AccountName); err != nil {
			return err
		}
		return nil
	})
	_s_pszAccountName := func(ptr interface{}) { o.AccountName = *ptr.(*string) }
	if err := w.ReadPointer(&o.AccountName, _s_pszAccountName, _ptr_pszAccountName); err != nil {
		return err
	}
	o.AccountPassword = make([]uint16, 257)
	for i1 := range o.AccountPassword {
		i1 := i1
		if err := w.ReadData(&o.AccountPassword[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Error); err != nil {
		return err
	}
	return nil
}

// VirtualRootList structure represents INET_INFO_VIRTUAL_ROOT_LIST RPC structure.
//
// The INET_INFO_VIRTUAL_ROOT_LIST structure contains a list of INET_INFO_VIRTUAL_ROOT_ENTRY
// virtual root entries.
//
// This type is declared as follows:
type VirtualRootList struct {
	// cEntries:  The number of entries contained in the list.
	EntriesCount uint32 `idl:"name:cEntries" json:"entries_count"`
	// aVirtRootEntry:  An array of INET_INFO_VIRTUAL_ROOT_ENTRY entries.
	VirtRootEntry []*VirtualRootEntry `idl:"name:aVirtRootEntry;size_is:(cEntries)" json:"virt_root_entry"`
}

func (o *VirtualRootList) xxx_PreparePayload(ctx context.Context) error {
	if o.VirtRootEntry != nil && o.EntriesCount == 0 {
		o.EntriesCount = uint32(len(o.VirtRootEntry))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *VirtualRootList) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.EntriesCount)
	return []uint64{
		dimSize1,
	}
}
func (o *VirtualRootList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesCount); err != nil {
		return err
	}
	for i1 := range o.VirtRootEntry {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.VirtRootEntry[i1] != nil {
			if err := o.VirtRootEntry[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&VirtualRootEntry{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.VirtRootEntry); uint64(i1) < sizeInfo[0]; i1++ {
		if err := (&VirtualRootEntry{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *VirtualRootList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.EntriesCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.EntriesCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.VirtRootEntry", sizeInfo[0])
	}
	o.VirtRootEntry = make([]*VirtualRootEntry, sizeInfo[0])
	for i1 := range o.VirtRootEntry {
		i1 := i1
		if o.VirtRootEntry[i1] == nil {
			o.VirtRootEntry[i1] = &VirtualRootEntry{}
		}
		if err := o.VirtRootEntry[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// ConfigInfo structure represents INET_INFO_CONFIG_INFO RPC structure.
//
// The INET_INFO_CONFIG_INFO structure stores configuration values for an Internet protocol
// server.<3>
//
// This type is declared as follows:
type ConfigInfo struct {
	// FieldControl:  A 32-bit unsigned integer that specifies a bit field. This field specifies
	// the values of the INET_INFO_CONFIG_INFO structure that contain valid data. An implementation
	// MUST set the flag corresponding to the structure field when returning or updating
	// configuration data. This field MUST be set to a valid combination of the following
	// values.
	//
	//	+---------------------------------+------------+
	//	|                                 |            |
	//	|              NAME               |   VALUE    |
	//	|                                 |            |
	//	+---------------------------------+------------+
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_CONNECTION_TIMEOUT | 0x00000001 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_MAX_CONNECTIONS    | 0x00000002 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_LOG_CONFIG         | 0x00000004 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_ADMIN_NAME         | 0x00000008 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_SERVER_COMMENT     | 0x00000010 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_ADMIN_EMAIL        | 0x00000020 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_HOST_NAME          | 0x00000040 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_SERVER_SIZE        | 0x00000080 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_DEF_LOGON_DOMAIN   | 0x00008000 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_AUTHENTICATION     | 0x00010000 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_ALLOW_ANONYMOUS    | 0x00020000 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_LOG_ANONYMOUS      | 0x00040000 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_LOG_NONANONYMOUS   | 0x00080000 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_ANON_USER_NAME     | 0x00100000 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_ANON_PASSWORD      | 0x00200000 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_PORT_NUMBER        | 0x00400000 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_SITE_SECURITY      | 0x00800000 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_VIRTUAL_ROOTS      | 0x01000000 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_SECURE_PORT_NUMBER | 0x02000000 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_SERVER_NAME        | 0x04000000 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_AUTO_START         | 0x08000000 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_ADDRESS_TYPE       | 0x10000000 |
	//	+---------------------------------+------------+
	//	| FC_INET_INFO_IP_ADDRESS         | 0x20000000 |
	//	+---------------------------------+------------+
	FieldControl uint32 `idl:"name:FieldControl" json:"field_control"`
	// dwConnectionTimeout:  The time limit to maintain an inactive connection specified
	// as the number of seconds from the last request.
	ConnectionTimeout uint32 `idl:"name:dwConnectionTimeout" json:"connection_timeout"`
	// dwMaxConnections:  The maximum number of allowed connections to the Internet protocol
	// server.
	MaxConnections uint32 `idl:"name:dwMaxConnections" json:"max_connections"`
	// lpszAdminName:  A pointer to a null-terminated Unicode string that contains the name
	// of the Internet protocol server administrator.
	AdminName string `idl:"name:lpszAdminName;string" json:"admin_name"`
	// lpszAdminEmail:  A pointer to a null-terminated Unicode string that contains the
	// email address of the Internet protocol server administrator.
	AdminEmail string `idl:"name:lpszAdminEmail;string" json:"admin_email"`
	// lpszServerComment:  A pointer to a null-terminated Unicode string that contains a
	// comment that describes the Internet protocol server instance.
	ServerComment string `idl:"name:lpszServerComment;string" json:"server_comment"`
	// lpLogConfig:  A pointer to an INET_LOG_CONFIGURATION structure that specifies the
	// configuration of the Internet protocol server log.
	LogConfig *InetLogConfiguration `idl:"name:lpLogConfig" json:"log_config"`
	// LangId:   A WORD containing the language identifier, a standard international numerical
	// identifier for the language in the country or region where the server is in use,
	// as specified in [MS-LCID] section 2.1.
	LangID uint16 `idl:"name:LangId" json:"lang_id"`
	// LocalId:  A locale identifier that is a DWORD value that consists of a language identifier,
	// such as one specified for the LangID member, combined with a sort identifier indicating
	// location. For example, the LangID member might indicate French, where the LocalID
	// indicates Canada. The LocalID member is given as specified in [MS-LCID] section 2.1.
	LocalID uint32 `idl:"name:LocalId" json:"local_id"`
	// ProductId:  An array of 64 bytes that MAY contain a string value, which identifies
	// the server implementation.
	ProductID []byte `idl:"name:ProductId" json:"product_id"`
	// fLogAnonymous:  A value that is set to TRUE if data transfers from anonymous users
	// are to be logged.
	LogAnonymous bool `idl:"name:fLogAnonymous" json:"log_anonymous"`
	// fLogNonAnonymous:  A value that is set to TRUE if data transfers from non-anonymous
	// users are to be logged.
	LogNonAnonymous bool `idl:"name:fLogNonAnonymous" json:"log_non_anonymous"`
	// lpszAnonUserName:  A pointer to a null-terminated Unicode string that contains the
	// name requested and accepted from an anonymous user.
	AnonUserName string `idl:"name:lpszAnonUserName;string" json:"anon_user_name"`
	// szAnonPassword:  An array of 257 Unicode characters containing a null-terminated
	// Unicode string that contains a password requested and accepted from an anonymous
	// user.<4>
	AnonPassword []uint16 `idl:"name:szAnonPassword" json:"anon_password"`
	// dwAuthentication:  A value that indicates what authentication methods can be used.
	Authentication uint32 `idl:"name:dwAuthentication" json:"authentication"`
	// sPort:  A 16-bit unsigned integer that specifies the network port on which the Internet
	// protocol server instance is running.
	Port int16 `idl:"name:sPort" json:"port"`
	// DenyIPList:  A pointer to an INET_INFO_IP_SEC_LIST structure that contains a list
	// of IP addresses that will be denied connections to the Internet protocol server.
	DenyIPList *IPsecList `idl:"name:DenyIPList" json:"deny_ip_list"`
	// GrantIPList:  A pointer to an INET_INFO_IP_SEC_LIST structure that contains a list
	// of IP addresses that will be granted connections to the Internet protocol server.
	GrantIPList *IPsecList `idl:"name:GrantIPList" json:"grant_ip_list"`
	// VirtualRoots:  A pointer to an INET_INFO_VIRTUAL_ROOT_LIST structure that contains
	// a list of virtual root directories for the Internet protocol server instance.
	VirtualRoots *VirtualRootList `idl:"name:VirtualRoots" json:"virtual_roots"`
}

func (o *ConfigInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ConfigInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.FieldControl); err != nil {
		return err
	}
	if err := w.WriteData(o.ConnectionTimeout); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxConnections); err != nil {
		return err
	}
	if o.AdminName != "" {
		_ptr_lpszAdminName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.AdminName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.AdminName, _ptr_lpszAdminName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.AdminEmail != "" {
		_ptr_lpszAdminEmail := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.AdminEmail); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.AdminEmail, _ptr_lpszAdminEmail); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.ServerComment != "" {
		_ptr_lpszServerComment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ServerComment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ServerComment, _ptr_lpszServerComment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.LogConfig != nil {
		_ptr_lpLogConfig := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.LogConfig != nil {
				if err := o.LogConfig.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&InetLogConfiguration{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.LogConfig, _ptr_lpLogConfig); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.LangID); err != nil {
		return err
	}
	if err := w.WriteData(o.LocalID); err != nil {
		return err
	}
	for i1 := range o.ProductID {
		i1 := i1
		if uint64(i1) >= 64 {
			break
		}
		if err := w.WriteData(o.ProductID[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.ProductID); uint64(i1) < 64; i1++ {
		if err := w.WriteData(uint8(0)); err != nil {
			return err
		}
	}
	if !o.LogAnonymous {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if !o.LogNonAnonymous {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if o.AnonUserName != "" {
		_ptr_lpszAnonUserName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.AnonUserName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.AnonUserName, _ptr_lpszAnonUserName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	for i1 := range o.AnonPassword {
		i1 := i1
		if uint64(i1) >= 257 {
			break
		}
		if err := w.WriteData(o.AnonPassword[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.AnonPassword); uint64(i1) < 257; i1++ {
		if err := w.WriteData(uint16(0)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Authentication); err != nil {
		return err
	}
	if err := w.WriteData(o.Port); err != nil {
		return err
	}
	if o.DenyIPList != nil {
		_ptr_DenyIPList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.DenyIPList != nil {
				if err := o.DenyIPList.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPsecList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.DenyIPList, _ptr_DenyIPList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.GrantIPList != nil {
		_ptr_GrantIPList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.GrantIPList != nil {
				if err := o.GrantIPList.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IPsecList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.GrantIPList, _ptr_GrantIPList); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.VirtualRoots != nil {
		_ptr_VirtualRoots := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.VirtualRoots != nil {
				if err := o.VirtualRoots.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&VirtualRootList{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.VirtualRoots, _ptr_VirtualRoots); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *ConfigInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.FieldControl); err != nil {
		return err
	}
	if err := w.ReadData(&o.ConnectionTimeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxConnections); err != nil {
		return err
	}
	_ptr_lpszAdminName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.AdminName); err != nil {
			return err
		}
		return nil
	})
	_s_lpszAdminName := func(ptr interface{}) { o.AdminName = *ptr.(*string) }
	if err := w.ReadPointer(&o.AdminName, _s_lpszAdminName, _ptr_lpszAdminName); err != nil {
		return err
	}
	_ptr_lpszAdminEmail := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.AdminEmail); err != nil {
			return err
		}
		return nil
	})
	_s_lpszAdminEmail := func(ptr interface{}) { o.AdminEmail = *ptr.(*string) }
	if err := w.ReadPointer(&o.AdminEmail, _s_lpszAdminEmail, _ptr_lpszAdminEmail); err != nil {
		return err
	}
	_ptr_lpszServerComment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ServerComment); err != nil {
			return err
		}
		return nil
	})
	_s_lpszServerComment := func(ptr interface{}) { o.ServerComment = *ptr.(*string) }
	if err := w.ReadPointer(&o.ServerComment, _s_lpszServerComment, _ptr_lpszServerComment); err != nil {
		return err
	}
	_ptr_lpLogConfig := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.LogConfig == nil {
			o.LogConfig = &InetLogConfiguration{}
		}
		if err := o.LogConfig.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_lpLogConfig := func(ptr interface{}) { o.LogConfig = *ptr.(**InetLogConfiguration) }
	if err := w.ReadPointer(&o.LogConfig, _s_lpLogConfig, _ptr_lpLogConfig); err != nil {
		return err
	}
	if err := w.ReadData(&o.LangID); err != nil {
		return err
	}
	if err := w.ReadData(&o.LocalID); err != nil {
		return err
	}
	o.ProductID = make([]byte, 64)
	for i1 := range o.ProductID {
		i1 := i1
		if err := w.ReadData(&o.ProductID[i1]); err != nil {
			return err
		}
	}
	var _bLogAnonymous int32
	if err := w.ReadData(&_bLogAnonymous); err != nil {
		return err
	}
	o.LogAnonymous = _bLogAnonymous != 0
	var _bLogNonAnonymous int32
	if err := w.ReadData(&_bLogNonAnonymous); err != nil {
		return err
	}
	o.LogNonAnonymous = _bLogNonAnonymous != 0
	_ptr_lpszAnonUserName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.AnonUserName); err != nil {
			return err
		}
		return nil
	})
	_s_lpszAnonUserName := func(ptr interface{}) { o.AnonUserName = *ptr.(*string) }
	if err := w.ReadPointer(&o.AnonUserName, _s_lpszAnonUserName, _ptr_lpszAnonUserName); err != nil {
		return err
	}
	o.AnonPassword = make([]uint16, 257)
	for i1 := range o.AnonPassword {
		i1 := i1
		if err := w.ReadData(&o.AnonPassword[i1]); err != nil {
			return err
		}
	}
	if err := w.ReadData(&o.Authentication); err != nil {
		return err
	}
	if err := w.ReadData(&o.Port); err != nil {
		return err
	}
	_ptr_DenyIPList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.DenyIPList == nil {
			o.DenyIPList = &IPsecList{}
		}
		if err := o.DenyIPList.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DenyIPList := func(ptr interface{}) { o.DenyIPList = *ptr.(**IPsecList) }
	if err := w.ReadPointer(&o.DenyIPList, _s_DenyIPList, _ptr_DenyIPList); err != nil {
		return err
	}
	_ptr_GrantIPList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.GrantIPList == nil {
			o.GrantIPList = &IPsecList{}
		}
		if err := o.GrantIPList.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_GrantIPList := func(ptr interface{}) { o.GrantIPList = *ptr.(**IPsecList) }
	if err := w.ReadPointer(&o.GrantIPList, _s_GrantIPList, _ptr_GrantIPList); err != nil {
		return err
	}
	_ptr_VirtualRoots := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.VirtualRoots == nil {
			o.VirtualRoots = &VirtualRootList{}
		}
		if err := o.VirtualRoots.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_VirtualRoots := func(ptr interface{}) { o.VirtualRoots = *ptr.(**VirtualRootList) }
	if err := w.ReadPointer(&o.VirtualRoots, _s_VirtualRoots, _ptr_VirtualRoots); err != nil {
		return err
	}
	return nil
}

// SiteEntry structure represents INET_INFO_SITE_ENTRY RPC structure.
//
// The INET_INFO_SITE_ENTRY structure contains information describing an Internet protocol
// server instance.
//
// This type is declared as follows:
//
// pszComment:  The server instance comment.
//
// dwInstance:  The server instance identifier.
type SiteEntry struct {
	Comment  string `idl:"name:pszComment;string" json:"comment"`
	Instance uint32 `idl:"name:dwInstance" json:"instance"`
}

func (o *SiteEntry) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SiteEntry) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Comment != "" {
		_ptr_pszComment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Comment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Comment, _ptr_pszComment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.Instance); err != nil {
		return err
	}
	return nil
}
func (o *SiteEntry) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_pszComment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Comment); err != nil {
			return err
		}
		return nil
	})
	_s_pszComment := func(ptr interface{}) { o.Comment = *ptr.(*string) }
	if err := w.ReadPointer(&o.Comment, _s_pszComment, _ptr_pszComment); err != nil {
		return err
	}
	if err := w.ReadData(&o.Instance); err != nil {
		return err
	}
	return nil
}

// SiteList structure represents INET_INFO_SITE_LIST RPC structure.
//
// The INET_INFO_SITE_LIST structure contains a list of INET_INFO_SITE_ENTRY site entries.
//
// This type is declared as follows:
type SiteList struct {
	// cEntries:  The number of entries contained in the list.
	EntriesCount uint32 `idl:"name:cEntries" json:"entries_count"`
	// aSiteEntry:   An array of INET_INFO_SITE_ENTRY site entries.
	SiteEntry []*SiteEntry `idl:"name:aSiteEntry;size_is:(cEntries)" json:"site_entry"`
}

func (o *SiteList) xxx_PreparePayload(ctx context.Context) error {
	if o.SiteEntry != nil && o.EntriesCount == 0 {
		o.EntriesCount = uint32(len(o.SiteEntry))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *SiteList) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.EntriesCount)
	return []uint64{
		dimSize1,
	}
}
func (o *SiteList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for sz1 := range sizeInfo {
			if err := w.WriteSize(sizeInfo[sz1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesCount); err != nil {
		return err
	}
	for i1 := range o.SiteEntry {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.SiteEntry[i1] != nil {
			if err := o.SiteEntry[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SiteEntry{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.SiteEntry); uint64(i1) < sizeInfo[0]; i1++ {
		if err := (&SiteEntry{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *SiteList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	sizeInfo, ok := ctx.Value(ndr.SizeInfo).([]uint64)
	if !ok {
		sizeInfo = o.NDRSizeInfo()
		for i1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[i1]); err != nil {
				return err
			}
		}
		ctx = context.WithValue(ctx, ndr.SizeInfo, sizeInfo)
	}
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesCount); err != nil {
		return err
	}
	// XXX: for opaque unmarshaling
	if o.EntriesCount > 0 && sizeInfo[0] == 0 {
		sizeInfo[0] = uint64(o.EntriesCount)
	}
	if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
		return fmt.Errorf("buffer overflow for size %d of array o.SiteEntry", sizeInfo[0])
	}
	o.SiteEntry = make([]*SiteEntry, sizeInfo[0])
	for i1 := range o.SiteEntry {
		i1 := i1
		if o.SiteEntry[i1] == nil {
			o.SiteEntry[i1] = &SiteEntry{}
		}
		if err := o.SiteEntry[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// GlobalConfigInfo structure represents INET_INFO_GLOBAL_CONFIG_INFO RPC structure.
//
// The INET_INFO_GLOBAL_CONFIG_INFO structure contains configuration data global to
// all the Internet protocol services managed by this protocol.
//
// This type is declared as follows:
type GlobalConfigInfo struct {
	// FieldControl:  A bit-field that specifies the values of the INET_INFO_GLOBAL_CONFIG_INFO
	// structure that have been initialized. An implementation MUST set the flag corresponding
	// to the structure field when returning or updating configuration data. This field
	// MUST be set to a valid combination of the following values.
	//
	//	+------------+---------------------------------+
	//	|            |                                 |
	//	|   VALUE    |             MEANING             |
	//	|            |                                 |
	//	+------------+---------------------------------+
	//	+------------+---------------------------------+
	//	| 0x00000001 | FC_GINET_INFO_BANDWIDTH_LEVEL   |
	//	+------------+---------------------------------+
	//	| 0x00000002 | FC_GINET_INFO_MEMORY_CACHE_SIZE |
	//	+------------+---------------------------------+
	FieldControl uint32 `idl:"name:FieldControl" json:"field_control"`
	// BandwidthLevel:  The bytes per second to send over the network for the Internet protocol
	// servers.
	BandwidthLevel uint32 `idl:"name:BandwidthLevel" json:"bandwidth_level"`
	// cbMemoryCacheSize:  The size of the in-memory file cache for the Internet protocol
	// servers.
	MemoryCacheLength uint32 `idl:"name:cbMemoryCacheSize" json:"memory_cache_length"`
}

func (o *GlobalConfigInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *GlobalConfigInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.FieldControl); err != nil {
		return err
	}
	if err := w.WriteData(o.BandwidthLevel); err != nil {
		return err
	}
	if err := w.WriteData(o.MemoryCacheLength); err != nil {
		return err
	}
	return nil
}
func (o *GlobalConfigInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.FieldControl); err != nil {
		return err
	}
	if err := w.ReadData(&o.BandwidthLevel); err != nil {
		return err
	}
	if err := w.ReadData(&o.MemoryCacheLength); err != nil {
		return err
	}
	return nil
}

// CacheStatistics structure represents INETA_CACHE_STATISTICS RPC structure.
//
// The INETA_CACHE_STATISTICS structure contains statistics for the Internet protocol
// server's caches.<6>
//
// This type is declared as follows:
type CacheStatistics struct {
	// FilesCached:  The current number of files whose content is in the Internet protocol
	// server file cache.
	FilesCached uint32 `idl:"name:FilesCached" json:"files_cached"`
	// TotalFilesCached:  The total number of files whose content has been cached since
	// Internet protocol server startup.
	TotalFilesCached uint32 `idl:"name:TotalFilesCached" json:"total_files_cached"`
	// FileHits:  The number of successful lookups in the Internet protocol server's file
	// cache.
	FileHits uint32 `idl:"name:FileHits" json:"file_hits"`
	// FileMisses:  The number of unsuccessful lookups in the Internet protocol server's
	// file cache.
	FileMisses uint32 `idl:"name:FileMisses" json:"file_misses"`
	// FileFlushes:  The number of file cache flushes since Internet protocol server startup.
	FileFlushes uint32 `idl:"name:FileFlushes" json:"file_flushes"`
	// CurrentFileCacheSize:  The current number of bytes used for the Internet protocol
	// server's file cache.
	CurrentFileCacheSize uint64 `idl:"name:CurrentFileCacheSize" json:"current_file_cache_size"`
	// MaximumFileCacheSize:  The maximum number of bytes used for the Internet protocol
	// server's file cache.
	MaximumFileCacheSize uint64 `idl:"name:MaximumFileCacheSize" json:"maximum_file_cache_size"`
	// FlushedEntries:  The number of files that are marked for removal from the Internet
	// protocol server cache after the current transfers are complete.
	FlushedEntries uint32 `idl:"name:FlushedEntries" json:"flushed_entries"`
	// TotalFlushed:  The number of files removed from the cache since Internet protocol
	// server startup.
	TotalFlushed uint32 `idl:"name:TotalFlushed" json:"total_flushed"`
	// URICached:  The number of URI information blocks currently cached by the Internet
	// protocol server.
	URICached uint32 `idl:"name:URICached" json:"uri_cached"`
	// TotalURICached:  The total number of URI information blocks ever added to the cache
	// for the Internet protocol server.
	TotalURICached uint32 `idl:"name:TotalURICached" json:"total_uri_cached"`
	// URIHits:   The number of successful lookups in the Internet protocol server's URI
	// cache.
	URIHits uint32 `idl:"name:URIHits" json:"uri_hits"`
	// URIMisses:  The number of unsuccessful lookups in the Internet protocol server's
	// URI cache.
	URIMisses uint32 `idl:"name:URIMisses" json:"uri_misses"`
	// URIFlushes:  The number of URI cache flushes since Internet protocol server startup.
	URIFlushes uint32 `idl:"name:URIFlushes" json:"uri_flushes"`
	// TotalURIFlushed:  The total number of URI information blocks that have been removed
	// from the cache since Internet protocol server startup.
	TotalURIFlushed uint32 `idl:"name:TotalURIFlushed" json:"total_uri_flushed"`
	// BlobCached:  The number of BLOB information blocks currently cached by the Internet
	// protocol server.
	BlobCached uint32 `idl:"name:BlobCached" json:"blob_cached"`
	// TotalBlobCached:  The total number of BLOB information blocks ever added to the Internet
	// protocol server's cache.
	TotalBlobCached uint32 `idl:"name:TotalBlobCached" json:"total_blob_cached"`
	// BlobHits:  The number of successful lookups in the Internet protocol server's BLOB
	// cache.
	BlobHits uint32 `idl:"name:BlobHits" json:"blob_hits"`
	// BlobMisses:  The number of unsuccessful lookups in the Internet protocol server's
	// BLOB cache.
	BlobMisses uint32 `idl:"name:BlobMisses" json:"blob_misses"`
	// BlobFlushes:  The number of BLOB cache flushes since Internet protocol server startup.
	BlobFlushes uint32 `idl:"name:BlobFlushes" json:"blob_flushes"`
	// TotalBlobFlushed:  The total number of BLOB information blocks that have been removed
	// from the cache since Internet protocol server startup.
	TotalBlobFlushed uint32 `idl:"name:TotalBlobFlushed" json:"total_blob_flushed"`
}

func (o *CacheStatistics) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *CacheStatistics) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.FilesCached); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalFilesCached); err != nil {
		return err
	}
	if err := w.WriteData(o.FileHits); err != nil {
		return err
	}
	if err := w.WriteData(o.FileMisses); err != nil {
		return err
	}
	if err := w.WriteData(o.FileFlushes); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentFileCacheSize); err != nil {
		return err
	}
	if err := w.WriteData(o.MaximumFileCacheSize); err != nil {
		return err
	}
	if err := w.WriteData(o.FlushedEntries); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalFlushed); err != nil {
		return err
	}
	if err := w.WriteData(o.URICached); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalURICached); err != nil {
		return err
	}
	if err := w.WriteData(o.URIHits); err != nil {
		return err
	}
	if err := w.WriteData(o.URIMisses); err != nil {
		return err
	}
	if err := w.WriteData(o.URIFlushes); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalURIFlushed); err != nil {
		return err
	}
	if err := w.WriteData(o.BlobCached); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalBlobCached); err != nil {
		return err
	}
	if err := w.WriteData(o.BlobHits); err != nil {
		return err
	}
	if err := w.WriteData(o.BlobMisses); err != nil {
		return err
	}
	if err := w.WriteData(o.BlobFlushes); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalBlobFlushed); err != nil {
		return err
	}
	return nil
}
func (o *CacheStatistics) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.FilesCached); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalFilesCached); err != nil {
		return err
	}
	if err := w.ReadData(&o.FileHits); err != nil {
		return err
	}
	if err := w.ReadData(&o.FileMisses); err != nil {
		return err
	}
	if err := w.ReadData(&o.FileFlushes); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentFileCacheSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaximumFileCacheSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.FlushedEntries); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalFlushed); err != nil {
		return err
	}
	if err := w.ReadData(&o.URICached); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalURICached); err != nil {
		return err
	}
	if err := w.ReadData(&o.URIHits); err != nil {
		return err
	}
	if err := w.ReadData(&o.URIMisses); err != nil {
		return err
	}
	if err := w.ReadData(&o.URIFlushes); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalURIFlushed); err != nil {
		return err
	}
	if err := w.ReadData(&o.BlobCached); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalBlobCached); err != nil {
		return err
	}
	if err := w.ReadData(&o.BlobHits); err != nil {
		return err
	}
	if err := w.ReadData(&o.BlobMisses); err != nil {
		return err
	}
	if err := w.ReadData(&o.BlobFlushes); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalBlobFlushed); err != nil {
		return err
	}
	return nil
}

// ATQStatistics structure represents INETA_ATQ_STATISTICS RPC structure.
//
// The INETA_ATQ_STATISTICS structure contains network I/O statistics and client connection
// information for an Internet protocol server.
//
// This type is declared as follows:
type ATQStatistics struct {
	// TotalBlockedRequests:  The total number of blocked requests.
	TotalBlockedRequests uint32 `idl:"name:TotalBlockedRequests" json:"total_blocked_requests"`
	// TotalRejectedRequests:  The total number of rejected requests.
	TotalRejectedRequests uint32 `idl:"name:TotalRejectedRequests" json:"total_rejected_requests"`
	// TotalAllowedRequests:  The total number of allowed requests.
	TotalAllowedRequests uint32 `idl:"name:TotalAllowedRequests" json:"total_allowed_requests"`
	// CurrentBlockedRequests:  The total number of currently blocked requests.
	CurrentBlockedRequests uint32 `idl:"name:CurrentBlockedRequests" json:"current_blocked_requests"`
	// MeasuredBandwidth:  The measured network bandwidth in bytes per second.
	MeasuredBandwidth uint32 `idl:"name:MeasuredBandwidth" json:"measured_bandwidth"`
}

func (o *ATQStatistics) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ATQStatistics) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalBlockedRequests); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalRejectedRequests); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalAllowedRequests); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentBlockedRequests); err != nil {
		return err
	}
	if err := w.WriteData(o.MeasuredBandwidth); err != nil {
		return err
	}
	return nil
}
func (o *ATQStatistics) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalBlockedRequests); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalRejectedRequests); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalAllowedRequests); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentBlockedRequests); err != nil {
		return err
	}
	if err := w.ReadData(&o.MeasuredBandwidth); err != nil {
		return err
	}
	return nil
}

// Statistics0 structure represents INET_INFO_STATISTICS_0 RPC structure.
//
// The INET_INFO_STATISTICS_0 structure contains statistics for an Internet protocol
// server.
//
// This type is declared as follows:
type Statistics0 struct {
	// CacheCtrs:  The structure of type INETA_CACHE_STATISTICS that contains statistics
	// on the Internet protocol server cache.
	CacheCounters *CacheStatistics `idl:"name:CacheCtrs" json:"cache_counters"`
	// AtqCtrs:  The structure of type INETA_ATQ_STATISTICS that contains statistics on
	// the Internet protocol server network I/O.
	ATQCounters *ATQStatistics `idl:"name:AtqCtrs" json:"atq_counters"`
	// nAuxCounters:  The number of initialized elements in rgCounters. MUST be 0.
	AuxCounters uint32 `idl:"name:nAuxCounters" json:"aux_counters"`
	// rgCounters:  An array of 20 elements. This field is unused and MUST be ignored by
	// clients.
	Counters []uint32 `idl:"name:rgCounters" json:"counters"`
}

func (o *Statistics0) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Statistics0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.CacheCounters != nil {
		if err := o.CacheCounters.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&CacheStatistics{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.ATQCounters != nil {
		if err := o.ATQCounters.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&ATQStatistics{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.AuxCounters); err != nil {
		return err
	}
	for i1 := range o.Counters {
		i1 := i1
		if uint64(i1) >= 20 {
			break
		}
		if err := w.WriteData(o.Counters[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Counters); uint64(i1) < 20; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *Statistics0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.CacheCounters == nil {
		o.CacheCounters = &CacheStatistics{}
	}
	if err := o.CacheCounters.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.ATQCounters == nil {
		o.ATQCounters = &ATQStatistics{}
	}
	if err := o.ATQCounters.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuxCounters); err != nil {
		return err
	}
	o.Counters = make([]uint32, 20)
	for i1 := range o.Counters {
		i1 := i1
		if err := w.ReadData(&o.Counters[i1]); err != nil {
			return err
		}
	}
	return nil
}

// StatisticsInfo structure represents INET_INFO_STATISTICS_INFO RPC union.
type StatisticsInfo struct {
	// Types that are assignable to Value
	//
	// *StatisticsInfo_InetStats0
	Value is_StatisticsInfo `json:"value"`
}

func (o *StatisticsInfo) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *StatisticsInfo_InetStats0:
		if value != nil {
			return value.InetStats0
		}
	}
	return nil
}

type is_StatisticsInfo interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_StatisticsInfo()
}

func (o *StatisticsInfo) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *StatisticsInfo_InetStats0:
		return uint32(0)
	}
	return uint32(0)
}

func (o *StatisticsInfo) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		_o, _ := o.Value.(*StatisticsInfo_InetStats0)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&StatisticsInfo_InetStats0{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *StatisticsInfo) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		o.Value = &StatisticsInfo_InetStats0{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// StatisticsInfo_InetStats0 structure represents INET_INFO_STATISTICS_INFO RPC union arm.
//
// It has following labels: 0
type StatisticsInfo_InetStats0 struct {
	InetStats0 *Statistics0 `idl:"name:InetStats0" json:"inet_stats0"`
}

func (*StatisticsInfo_InetStats0) is_StatisticsInfo() {}

func (o *StatisticsInfo_InetStats0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.InetStats0 != nil {
		_ptr_InetStats0 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.InetStats0 != nil {
				if err := o.InetStats0.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Statistics0{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.InetStats0, _ptr_InetStats0); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *StatisticsInfo_InetStats0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_InetStats0 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.InetStats0 == nil {
			o.InetStats0 = &Statistics0{}
		}
		if err := o.InetStats0.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_InetStats0 := func(ptr interface{}) { o.InetStats0 = *ptr.(**Statistics0) }
	if err := w.ReadPointer(&o.InetStats0, _s_InetStats0, _ptr_InetStats0); err != nil {
		return err
	}
	return nil
}

// W3Statistics1 structure represents W3_STATISTICS_1 RPC structure.
//
// The W3_STATISTICS_1 structure contains statistics on the usage of the HTTP server.
//
// This type is declared as follows:
type W3Statistics1 struct {
	// TotalBytesSent:  The total number of bytes sent.
	TotalBytesSent *dtyp.LargeInteger `idl:"name:TotalBytesSent" json:"total_bytes_sent"`
	// TotalBytesReceived:  The total number of bytes received.
	TotalBytesReceived *dtyp.LargeInteger `idl:"name:TotalBytesReceived" json:"total_bytes_received"`
	// TotalFilesSent:  The total number of files sent by the HTTP server.
	TotalFilesSent uint32 `idl:"name:TotalFilesSent" json:"total_files_sent"`
	// TotalFilesReceived:  The total number of files received by the HTTP server.
	TotalFilesReceived uint32 `idl:"name:TotalFilesReceived" json:"total_files_received"`
	// CurrentAnonymousUsers:  The current number of anonymous users connected to the HTTP
	// server.
	CurrentAnonymousUsers uint32 `idl:"name:CurrentAnonymousUsers" json:"current_anonymous_users"`
	// CurrentNonAnonymousUsers:  The current number of non-anonymous users connected to
	// the HTTP server.
	CurrentNonAnonymousUsers uint32 `idl:"name:CurrentNonAnonymousUsers" json:"current_non_anonymous_users"`
	// TotalAnonymousUsers:  The total number of anonymous users that have ever connected
	// to the HTTP server.
	TotalAnonymousUsers uint32 `idl:"name:TotalAnonymousUsers" json:"total_anonymous_users"`
	// TotalNonAnonymousUsers:  The total number of non-anonymous users that have ever connected
	// to the HTTP server.
	TotalNonAnonymousUsers uint32 `idl:"name:TotalNonAnonymousUsers" json:"total_non_anonymous_users"`
	// MaxAnonymousUsers:  The maximum number of anonymous users who simultaneously connected
	// to the HTTP server.
	MaxAnonymousUsers uint32 `idl:"name:MaxAnonymousUsers" json:"max_anonymous_users"`
	// MaxNonAnonymousUsers:  The maximum number of non-anonymous users who simultaneously
	// connected to the HTTP server.
	MaxNonAnonymousUsers uint32 `idl:"name:MaxNonAnonymousUsers" json:"max_non_anonymous_users"`
	// CurrentConnections:  The current number of connections to the HTTP server.
	CurrentConnections uint32 `idl:"name:CurrentConnections" json:"current_connections"`
	// MaxConnections:  The maximum number of connections to the HTTP server.
	MaxConnections uint32 `idl:"name:MaxConnections" json:"max_connections"`
	// ConnectionAttempts:  The number of connection attempts that have been made to the
	// HTTP server.
	ConnectionAttempts uint32 `idl:"name:ConnectionAttempts" json:"connection_attempts"`
	// LogonAttempts:  The number of logon attempts that have been made to the HTTP server.
	LogonAttempts uint32 `idl:"name:LogonAttempts" json:"logon_attempts"`
	// TotalOptions:  The total number of HTTP requests made with the OPTIONS method.
	TotalOptions uint32 `idl:"name:TotalOptions" json:"total_options"`
	// TotalGets:  The total number of HTTP requests made using the GET method.
	TotalGets uint32 `idl:"name:TotalGets" json:"total_gets"`
	// TotalPosts:  The total number of HTTP requests made using the POST method.
	TotalPosts uint32 `idl:"name:TotalPosts" json:"total_posts"`
	// TotalHeads:  The total number of HTTP requests made using the HEAD method.
	TotalHeads uint32 `idl:"name:TotalHeads" json:"total_heads"`
	// TotalPuts:  The total number of HTTP requests made using the PUT method.
	TotalPuts uint32 `idl:"name:TotalPuts" json:"total_puts"`
	// TotalDeletes:  The total number of HTTP requests made using the DELETE method.
	TotalDeletes uint32 `idl:"name:TotalDeletes" json:"total_deletes"`
	// TotalTraces:   The total number of HTTP requests made using the TRACE method.
	TotalTraces uint32 `idl:"name:TotalTraces" json:"total_traces"`
	// TotalMove:   The total number of WebDAV requests made using the MOVE method. For
	// more information on WebDAV requests, see [RFC2518].
	TotalMove uint32 `idl:"name:TotalMove" json:"total_move"`
	// TotalCopy:  The total number of WebDAV requests made using the COPY method.
	TotalCopy uint32 `idl:"name:TotalCopy" json:"total_copy"`
	// TotalMkcol:  The total number of WebDAV requests made using the MKCOL method.
	TotalMkcol uint32 `idl:"name:TotalMkcol" json:"total_mkcol"`
	// TotalPropfind:  The total number of WebDAV requests made using the PROPFIND method.
	TotalPropfind uint32 `idl:"name:TotalPropfind" json:"total_propfind"`
	// TotalProppatch:  The total number of WebDAV requests made using the PROPPATCH method.
	TotalProppatch uint32 `idl:"name:TotalProppatch" json:"total_proppatch"`
	// TotalSearch:  The total number of requests made using the SEARCH method.
	TotalSearch uint32 `idl:"name:TotalSearch" json:"total_search"`
	// TotalLock:  The total number of WebDAV requests made using the LOCK method.
	TotalLock uint32 `idl:"name:TotalLock" json:"total_lock"`
	// TotalUnlock:  The total number of WebDAV requests made using the UNLOCK method.
	TotalUnlock uint32 `idl:"name:TotalUnlock" json:"total_unlock"`
	// TotalOthers:  The total number of HTTP requests made to methods not already listed.
	TotalOthers uint32 `idl:"name:TotalOthers" json:"total_others"`
	// TotalCGIRequests:  The total number of Common Gateway Interface (CGI) requests ever
	// made to the HTTP server.
	TotalCGIRequests uint32 `idl:"name:TotalCGIRequests" json:"total_cgi_requests"`
	// TotalBGIRequests:  The total number of Binary Gateway Interface (BGI) requests ever
	// made to the HTTP server.
	TotalBGIRequests uint32 `idl:"name:TotalBGIRequests" json:"total_bgi_requests"`
	// TotalNotFoundErrors:  The total number of requests that could not be satisfied by
	// the server because the requested document could not be found. These requests are
	// generally reported as an HTTP 404 error code to the client.
	TotalNotFoundErrors uint32 `idl:"name:TotalNotFoundErrors" json:"total_not_found_errors"`
	// TotalLockedErrors:  The total number of locked errors.
	TotalLockedErrors uint32 `idl:"name:TotalLockedErrors" json:"total_locked_errors"`
	// CurrentCalAuth:  The current number of Client Access Licenses (CALs) that are authorized
	// (for more information, see [MSFT-CAL]).
	CurrentCALAuth uint32 `idl:"name:CurrentCalAuth" json:"current_cal_auth"`
	// MaxCalAuth:  The maximum number of CALs that are authorized.
	MaxCALAuth uint32 `idl:"name:MaxCalAuth" json:"max_cal_auth"`
	// TotalFailedCalAuth:  The total number of failed CAL authorizations.
	TotalFailedCALAuth uint32 `idl:"name:TotalFailedCalAuth" json:"total_failed_cal_auth"`
	// CurrentCalSsl:  The current number of CALs for a Secure Sockets Layer (SSL) connection.
	CurrentCALSSL uint32 `idl:"name:CurrentCalSsl" json:"current_cal_ssl"`
	// MaxCalSsl:  The maximum number of CALs for an SSL connection.
	MaxCALSSL uint32 `idl:"name:MaxCalSsl" json:"max_cal_ssl"`
	// TotalFailedCalSsl:  The total number of failed CAL SSL connections.
	TotalFailedCALSSL uint32 `idl:"name:TotalFailedCalSsl" json:"total_failed_cal_ssl"`
	// CurrentCGIRequests:  The current number of CGI requests. For more information on
	// CGI, see [RFC3875].
	CurrentCGIRequests uint32 `idl:"name:CurrentCGIRequests" json:"current_cgi_requests"`
	// CurrentBGIRequests:  The current number of BGI requests.
	CurrentBGIRequests uint32 `idl:"name:CurrentBGIRequests" json:"current_bgi_requests"`
	// MaxCGIRequests:  The maximum number of CGI requests allowed.
	MaxCGIRequests uint32 `idl:"name:MaxCGIRequests" json:"max_cgi_requests"`
	// MaxBGIRequests:  The maximum number of BGI requests allowed.
	MaxBGIRequests uint32 `idl:"name:MaxBGIRequests" json:"max_bgi_requests"`
	// CurrentBlockedRequests:  The current number of blocked requests.
	CurrentBlockedRequests uint32 `idl:"name:CurrentBlockedRequests" json:"current_blocked_requests"`
	// TotalBlockedRequests:  The total number of blocked requests.
	TotalBlockedRequests uint32 `idl:"name:TotalBlockedRequests" json:"total_blocked_requests"`
	// TotalAllowedRequests:  The total number of allowed requests to the HTTP server.
	TotalAllowedRequests uint32 `idl:"name:TotalAllowedRequests" json:"total_allowed_requests"`
	// TotalRejectedRequests:  The total number of rejected requests.
	TotalRejectedRequests uint32 `idl:"name:TotalRejectedRequests" json:"total_rejected_requests"`
	// MeasuredBw:  The measured network bandwidth for the HTTP server.
	MeasuredBandwidth uint32 `idl:"name:MeasuredBw" json:"measured_bandwidth"`
	// ServiceUptime:  The HTTP server uptime.
	ServiceUptime uint32 `idl:"name:ServiceUptime" json:"service_uptime"`
	// TimeOfLastClear:  The time of the last clear.
	TimeOfLastClear uint32 `idl:"name:TimeOfLastClear" json:"time_of_last_clear"`
	// nAuxCounters:  The number of initialized elements in rgCounters. MUST be 0.
	AuxCounters uint32 `idl:"name:nAuxCounters" json:"aux_counters"`
	// rgCounters:  An array of 20 elements. This field is unused and MUST be ignored by
	// clients.
	Counters []uint32 `idl:"name:rgCounters" json:"counters"`
}

func (o *W3Statistics1) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *W3Statistics1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.TotalBytesSent != nil {
		if err := o.TotalBytesSent.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.TotalBytesReceived != nil {
		if err := o.TotalBytesReceived.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TotalFilesSent); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalFilesReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentAnonymousUsers); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentNonAnonymousUsers); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalAnonymousUsers); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalNonAnonymousUsers); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxAnonymousUsers); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxNonAnonymousUsers); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentConnections); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxConnections); err != nil {
		return err
	}
	if err := w.WriteData(o.ConnectionAttempts); err != nil {
		return err
	}
	if err := w.WriteData(o.LogonAttempts); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalOptions); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalGets); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalPosts); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalHeads); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalPuts); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalDeletes); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalTraces); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalMove); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalCopy); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalMkcol); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalPropfind); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalProppatch); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalSearch); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalLock); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalUnlock); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalOthers); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalCGIRequests); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalBGIRequests); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalNotFoundErrors); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalLockedErrors); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentCALAuth); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxCALAuth); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalFailedCALAuth); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentCALSSL); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxCALSSL); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalFailedCALSSL); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentCGIRequests); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentBGIRequests); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxCGIRequests); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxBGIRequests); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentBlockedRequests); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalBlockedRequests); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalAllowedRequests); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalRejectedRequests); err != nil {
		return err
	}
	if err := w.WriteData(o.MeasuredBandwidth); err != nil {
		return err
	}
	if err := w.WriteData(o.ServiceUptime); err != nil {
		return err
	}
	if err := w.WriteData(o.TimeOfLastClear); err != nil {
		return err
	}
	if err := w.WriteData(o.AuxCounters); err != nil {
		return err
	}
	for i1 := range o.Counters {
		i1 := i1
		if uint64(i1) >= 20 {
			break
		}
		if err := w.WriteData(o.Counters[i1]); err != nil {
			return err
		}
	}
	for i1 := len(o.Counters); uint64(i1) < 20; i1++ {
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}
func (o *W3Statistics1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.TotalBytesSent == nil {
		o.TotalBytesSent = &dtyp.LargeInteger{}
	}
	if err := o.TotalBytesSent.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.TotalBytesReceived == nil {
		o.TotalBytesReceived = &dtyp.LargeInteger{}
	}
	if err := o.TotalBytesReceived.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalFilesSent); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalFilesReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentAnonymousUsers); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentNonAnonymousUsers); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalAnonymousUsers); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalNonAnonymousUsers); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxAnonymousUsers); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxNonAnonymousUsers); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentConnections); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxConnections); err != nil {
		return err
	}
	if err := w.ReadData(&o.ConnectionAttempts); err != nil {
		return err
	}
	if err := w.ReadData(&o.LogonAttempts); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalOptions); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalGets); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalPosts); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalHeads); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalPuts); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalDeletes); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalTraces); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalMove); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalCopy); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalMkcol); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalPropfind); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalProppatch); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalSearch); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalLock); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalUnlock); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalOthers); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalCGIRequests); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalBGIRequests); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalNotFoundErrors); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalLockedErrors); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentCALAuth); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxCALAuth); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalFailedCALAuth); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentCALSSL); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxCALSSL); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalFailedCALSSL); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentCGIRequests); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentBGIRequests); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxCGIRequests); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxBGIRequests); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentBlockedRequests); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalBlockedRequests); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalAllowedRequests); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalRejectedRequests); err != nil {
		return err
	}
	if err := w.ReadData(&o.MeasuredBandwidth); err != nil {
		return err
	}
	if err := w.ReadData(&o.ServiceUptime); err != nil {
		return err
	}
	if err := w.ReadData(&o.TimeOfLastClear); err != nil {
		return err
	}
	if err := w.ReadData(&o.AuxCounters); err != nil {
		return err
	}
	o.Counters = make([]uint32, 20)
	for i1 := range o.Counters {
		i1 := i1
		if err := w.ReadData(&o.Counters[i1]); err != nil {
			return err
		}
	}
	return nil
}

// W3Statistics structure represents W3_STATISTICS_STRUCT RPC union.
type W3Statistics struct {
	// Types that are assignable to Value
	//
	// *W3Statistics_StatInfo1
	Value is_W3Statistics `json:"value"`
}

func (o *W3Statistics) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *W3Statistics_StatInfo1:
		if value != nil {
			return value.StatInfo1
		}
	}
	return nil
}

type is_W3Statistics interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_W3Statistics()
}

func (o *W3Statistics) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *W3Statistics_StatInfo1:
		return uint32(0)
	}
	return uint32(0)
}

func (o *W3Statistics) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		_o, _ := o.Value.(*W3Statistics_StatInfo1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&W3Statistics_StatInfo1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *W3Statistics) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		o.Value = &W3Statistics_StatInfo1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// W3Statistics_StatInfo1 structure represents W3_STATISTICS_STRUCT RPC union arm.
//
// It has following labels: 0
type W3Statistics_StatInfo1 struct {
	StatInfo1 *W3Statistics1 `idl:"name:StatInfo1" json:"stat_info1"`
}

func (*W3Statistics_StatInfo1) is_W3Statistics() {}

func (o *W3Statistics_StatInfo1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.StatInfo1 != nil {
		_ptr_StatInfo1 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.StatInfo1 != nil {
				if err := o.StatInfo1.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&W3Statistics1{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.StatInfo1, _ptr_StatInfo1); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *W3Statistics_StatInfo1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_StatInfo1 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.StatInfo1 == nil {
			o.StatInfo1 = &W3Statistics1{}
		}
		if err := o.StatInfo1.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_StatInfo1 := func(ptr interface{}) { o.StatInfo1 = *ptr.(**W3Statistics1) }
	if err := w.ReadPointer(&o.StatInfo1, _s_StatInfo1, _ptr_StatInfo1); err != nil {
		return err
	}
	return nil
}

// FTPStatistics0 structure represents FTP_STATISTICS_0 RPC structure.
//
// The FTP_STATISTICS_0 structure contains statistics on the usage of the FTP server.
//
// This type is declared as follows:
type FTPStatistics0 struct {
	// TotalBytesSent:  The total number of bytes sent.
	TotalBytesSent *dtyp.LargeInteger `idl:"name:TotalBytesSent" json:"total_bytes_sent"`
	// TotalBytesReceived:  The total number of bytes received.
	TotalBytesReceived *dtyp.LargeInteger `idl:"name:TotalBytesReceived" json:"total_bytes_received"`
	// TotalFilesSent:  The total number of files sent by the FTP server.
	TotalFilesSent uint32 `idl:"name:TotalFilesSent" json:"total_files_sent"`
	// TotalFilesReceived:  The total number of files received by the FTP server.
	TotalFilesReceived uint32 `idl:"name:TotalFilesReceived" json:"total_files_received"`
	// CurrentAnonymousUsers:  The current number of anonymous users connected to the FTP
	// server.
	CurrentAnonymousUsers uint32 `idl:"name:CurrentAnonymousUsers" json:"current_anonymous_users"`
	// CurrentNonAnonymousUsers:  The current number of non-anonymous users connected to
	// the FTP server.
	CurrentNonAnonymousUsers uint32 `idl:"name:CurrentNonAnonymousUsers" json:"current_non_anonymous_users"`
	// TotalAnonymousUsers:  The total number of anonymous users that have ever connected
	// to the FTP server.
	TotalAnonymousUsers uint32 `idl:"name:TotalAnonymousUsers" json:"total_anonymous_users"`
	// TotalNonAnonymousUsers:  The total number of non-anonymous users that have ever connected
	// to the FTP server.
	TotalNonAnonymousUsers uint32 `idl:"name:TotalNonAnonymousUsers" json:"total_non_anonymous_users"`
	// MaxAnonymousUsers:  The maximum number of anonymous users allowed to simultaneously
	// connect to the FTP server.
	MaxAnonymousUsers uint32 `idl:"name:MaxAnonymousUsers" json:"max_anonymous_users"`
	// MaxNonAnonymousUsers:  The maximum number of non-anonymous users allowed to simultaneously
	// connect to the FTP server.
	MaxNonAnonymousUsers uint32 `idl:"name:MaxNonAnonymousUsers" json:"max_non_anonymous_users"`
	// CurrentConnections:  The current number of connections to the FTP server.
	CurrentConnections uint32 `idl:"name:CurrentConnections" json:"current_connections"`
	// MaxConnections:  The maximum number of connections to the FTP server.
	MaxConnections uint32 `idl:"name:MaxConnections" json:"max_connections"`
	// ConnectionAttempts:  The number of connection attempts that have been made to the
	// FTP server.
	ConnectionAttempts uint32 `idl:"name:ConnectionAttempts" json:"connection_attempts"`
	// LogonAttempts:  The number of logon attempts that have been made to the FTP server.
	LogonAttempts uint32 `idl:"name:LogonAttempts" json:"logon_attempts"`
	// ServiceUptime:  The time that the FTP server has been operational.
	ServiceUptime uint32 `idl:"name:ServiceUptime" json:"service_uptime"`
	// TotalAllowedRequests:  The total number of requests allowed to the FTP server.
	TotalAllowedRequests uint32 `idl:"name:TotalAllowedRequests" json:"total_allowed_requests"`
	// TotalRejectedRequests:  The total number of rejected requests.
	TotalRejectedRequests uint32 `idl:"name:TotalRejectedRequests" json:"total_rejected_requests"`
	// TotalBlockedRequests:  The total number of blocked requests.
	TotalBlockedRequests uint32 `idl:"name:TotalBlockedRequests" json:"total_blocked_requests"`
	// CurrentBlockedRequests:  The current number of blocked requests.
	CurrentBlockedRequests uint32 `idl:"name:CurrentBlockedRequests" json:"current_blocked_requests"`
	// MeasuredBandwidth:  The measured network bandwidth for the FTP server.
	MeasuredBandwidth uint32 `idl:"name:MeasuredBandwidth" json:"measured_bandwidth"`
	// TimeOfLastClear:  The time of the last clear.
	TimeOfLastClear uint32 `idl:"name:TimeOfLastClear" json:"time_of_last_clear"`
}

func (o *FTPStatistics0) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *FTPStatistics0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if o.TotalBytesSent != nil {
		if err := o.TotalBytesSent.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if o.TotalBytesReceived != nil {
		if err := o.TotalBytesReceived.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.LargeInteger{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.TotalFilesSent); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalFilesReceived); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentAnonymousUsers); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentNonAnonymousUsers); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalAnonymousUsers); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalNonAnonymousUsers); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxAnonymousUsers); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxNonAnonymousUsers); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentConnections); err != nil {
		return err
	}
	if err := w.WriteData(o.MaxConnections); err != nil {
		return err
	}
	if err := w.WriteData(o.ConnectionAttempts); err != nil {
		return err
	}
	if err := w.WriteData(o.LogonAttempts); err != nil {
		return err
	}
	if err := w.WriteData(o.ServiceUptime); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalAllowedRequests); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalRejectedRequests); err != nil {
		return err
	}
	if err := w.WriteData(o.TotalBlockedRequests); err != nil {
		return err
	}
	if err := w.WriteData(o.CurrentBlockedRequests); err != nil {
		return err
	}
	if err := w.WriteData(o.MeasuredBandwidth); err != nil {
		return err
	}
	if err := w.WriteData(o.TimeOfLastClear); err != nil {
		return err
	}
	return nil
}
func (o *FTPStatistics0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if o.TotalBytesSent == nil {
		o.TotalBytesSent = &dtyp.LargeInteger{}
	}
	if err := o.TotalBytesSent.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if o.TotalBytesReceived == nil {
		o.TotalBytesReceived = &dtyp.LargeInteger{}
	}
	if err := o.TotalBytesReceived.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalFilesSent); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalFilesReceived); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentAnonymousUsers); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentNonAnonymousUsers); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalAnonymousUsers); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalNonAnonymousUsers); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxAnonymousUsers); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxNonAnonymousUsers); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentConnections); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaxConnections); err != nil {
		return err
	}
	if err := w.ReadData(&o.ConnectionAttempts); err != nil {
		return err
	}
	if err := w.ReadData(&o.LogonAttempts); err != nil {
		return err
	}
	if err := w.ReadData(&o.ServiceUptime); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalAllowedRequests); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalRejectedRequests); err != nil {
		return err
	}
	if err := w.ReadData(&o.TotalBlockedRequests); err != nil {
		return err
	}
	if err := w.ReadData(&o.CurrentBlockedRequests); err != nil {
		return err
	}
	if err := w.ReadData(&o.MeasuredBandwidth); err != nil {
		return err
	}
	if err := w.ReadData(&o.TimeOfLastClear); err != nil {
		return err
	}
	return nil
}

// FTPStatistics structure represents FTP_STATISTICS_STRUCT RPC union.
type FTPStatistics struct {
	// Types that are assignable to Value
	//
	// *FTPStatistics_StatInfo0
	Value is_FTPStatistics `json:"value"`
}

func (o *FTPStatistics) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *FTPStatistics_StatInfo0:
		if value != nil {
			return value.StatInfo0
		}
	}
	return nil
}

type is_FTPStatistics interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_FTPStatistics()
}

func (o *FTPStatistics) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *FTPStatistics_StatInfo0:
		return uint32(0)
	}
	return uint32(0)
}

func (o *FTPStatistics) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		_o, _ := o.Value.(*FTPStatistics_StatInfo0)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&FTPStatistics_StatInfo0{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *FTPStatistics) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(0):
		o.Value = &FTPStatistics_StatInfo0{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// FTPStatistics_StatInfo0 structure represents FTP_STATISTICS_STRUCT RPC union arm.
//
// It has following labels: 0
type FTPStatistics_StatInfo0 struct {
	StatInfo0 *FTPStatistics0 `idl:"name:StatInfo0" json:"stat_info0"`
}

func (*FTPStatistics_StatInfo0) is_FTPStatistics() {}

func (o *FTPStatistics_StatInfo0) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.StatInfo0 != nil {
		_ptr_StatInfo0 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.StatInfo0 != nil {
				if err := o.StatInfo0.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&FTPStatistics0{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.StatInfo0, _ptr_StatInfo0); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *FTPStatistics_StatInfo0) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_StatInfo0 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.StatInfo0 == nil {
			o.StatInfo0 = &FTPStatistics0{}
		}
		if err := o.StatInfo0.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_StatInfo0 := func(ptr interface{}) { o.StatInfo0 = *ptr.(**FTPStatistics0) }
	if err := w.ReadPointer(&o.StatInfo0, _s_StatInfo0, _ptr_StatInfo0); err != nil {
		return err
	}
	return nil
}

// IISUserInfo1 structure represents IIS_USER_INFO_1 RPC structure.
//
// The IIS_USER_INFO_1 structure describes a user connected to an Internet protocol
// server.
//
// This type is declared as follows:
type IISUserInfo1 struct {
	// idUser:  A unique identifier for the user.
	UserID uint32 `idl:"name:idUser" json:"user_id"`
	// pszUser:  A name for the user, not necessarily unique.
	User string `idl:"name:pszUser;string" json:"user"`
	// fAnonymous:  Indicates whether or not the user connected anonymously. This field
	// MUST be one of the following values.
	//
	//	+---------+-----------------------------------------+
	//	|         |                                         |
	//	|  VALUE  |                 MEANING                 |
	//	|         |                                         |
	//	+---------+-----------------------------------------+
	//	+---------+-----------------------------------------+
	//	| TRUE 1  | The user is logged on as Anonymous.     |
	//	+---------+-----------------------------------------+
	//	| FALSE 0 | The user is not logged on as Anonymous. |
	//	+---------+-----------------------------------------+
	Anonymous bool `idl:"name:fAnonymous" json:"anonymous"`
	// inetHost:  The host IPv4 address. Data is stored in network byte order.
	InetHost uint32 `idl:"name:inetHost" json:"inet_host"`
	// tConnect:  The user connection time measured in elapsed seconds.
	Connect uint32 `idl:"name:tConnect" json:"connect"`
}

func (o *IISUserInfo1) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IISUserInfo1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.UserID); err != nil {
		return err
	}
	if o.User != "" {
		_ptr_pszUser := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.User); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.User, _ptr_pszUser); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if !o.Anonymous {
		if err := w.WriteData(int32(0)); err != nil {
			return err
		}
	} else {
		if err := w.WriteData(int32(1)); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.InetHost); err != nil {
		return err
	}
	if err := w.WriteData(o.Connect); err != nil {
		return err
	}
	return nil
}
func (o *IISUserInfo1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.UserID); err != nil {
		return err
	}
	_ptr_pszUser := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.User); err != nil {
			return err
		}
		return nil
	})
	_s_pszUser := func(ptr interface{}) { o.User = *ptr.(*string) }
	if err := w.ReadPointer(&o.User, _s_pszUser, _ptr_pszUser); err != nil {
		return err
	}
	var _bAnonymous int32
	if err := w.ReadData(&_bAnonymous); err != nil {
		return err
	}
	o.Anonymous = _bAnonymous != 0
	if err := w.ReadData(&o.InetHost); err != nil {
		return err
	}
	if err := w.ReadData(&o.Connect); err != nil {
		return err
	}
	return nil
}

// IISUserInfo1Container structure represents IIS_USER_INFO_1_CONTAINER RPC structure.
//
// The IIS_USER_INFO_1_CONTAINER structure contains a list of IIS_USER_INFO_1 structures
// describing users who are actively connected to the Internet protocol server.
//
// This type is declared as follows:
type IISUserInfo1Container struct {
	// EntriesRead:  The total number of IIS_USER_INFO objects in Buffer.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  The pointer to an array of IIS_USER_INFO_1 structures.
	Buffer []*IISUserInfo1 `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *IISUserInfo1Container) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.EntriesRead == 0 {
		o.EntriesRead = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IISUserInfo1Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.EntriesRead); err != nil {
		return err
	}
	if o.Buffer != nil || o.EntriesRead > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.EntriesRead)
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
				if o.Buffer[i1] != nil {
					if err := o.Buffer[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&IISUserInfo1{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&IISUserInfo1{}).MarshalNDR(ctx, w); err != nil {
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
}
func (o *IISUserInfo1Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.EntriesRead); err != nil {
		return err
	}
	_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.EntriesRead > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.EntriesRead)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]*IISUserInfo1, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &IISUserInfo1{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*IISUserInfo1) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// IISUserEnum structure represents IIS_USER_ENUM_STRUCT RPC structure.
//
// The IIS_USER_ENUM_STRUCT structure contains a pointer to an IIS_USER_INFO_1_CONTAINER.
//
// This type is declared as follows:
type IISUserEnum struct {
	// Level:  The value that indicates the level of detail in the information provided.
	// This member MUST be set to 1.
	Level uint32 `idl:"name:Level" json:"level"`
	// ConfigInfo:  The name of the contained union.
	ConfigInfo *IISUserEnum_ConfigInfo `idl:"name:ConfigInfo;switch_is:Level" json:"config_info"`
}

func (o *IISUserEnum) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *IISUserEnum) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Level); err != nil {
		return err
	}
	_swConfigInfo := uint32(o.Level)
	if o.ConfigInfo != nil {
		if err := o.ConfigInfo.MarshalUnionNDR(ctx, w, _swConfigInfo); err != nil {
			return err
		}
	} else {
		if err := (&IISUserEnum_ConfigInfo{}).MarshalUnionNDR(ctx, w, _swConfigInfo); err != nil {
			return err
		}
	}
	return nil
}
func (o *IISUserEnum) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Level); err != nil {
		return err
	}
	if o.ConfigInfo == nil {
		o.ConfigInfo = &IISUserEnum_ConfigInfo{}
	}
	_swConfigInfo := uint32(o.Level)
	if err := o.ConfigInfo.UnmarshalUnionNDR(ctx, w, _swConfigInfo); err != nil {
		return err
	}
	return nil
}

// IISUserEnum_ConfigInfo structure represents IIS_USER_ENUM_STRUCT union anonymous member.
//
// The IIS_USER_ENUM_STRUCT structure contains a pointer to an IIS_USER_INFO_1_CONTAINER.
//
// This type is declared as follows:
type IISUserEnum_ConfigInfo struct {
	// Types that are assignable to Value
	//
	// *IISUserEnum_ConfigInfo_Level1
	Value is_IISUserEnum_ConfigInfo `json:"value"`
}

func (o *IISUserEnum_ConfigInfo) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *IISUserEnum_ConfigInfo_Level1:
		if value != nil {
			return value.Level1
		}
	}
	return nil
}

type is_IISUserEnum_ConfigInfo interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_IISUserEnum_ConfigInfo()
}

func (o *IISUserEnum_ConfigInfo) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *IISUserEnum_ConfigInfo_Level1:
		return uint32(1)
	}
	return uint32(0)
}

func (o *IISUserEnum_ConfigInfo) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
	if err := w.WriteUnionAlign(9); err != nil {
		return err
	}
	if err := w.WriteSwitch(uint32(sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		_o, _ := o.Value.(*IISUserEnum_ConfigInfo_Level1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&IISUserEnum_ConfigInfo_Level1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *IISUserEnum_ConfigInfo) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
	if err := w.ReadUnionAlign(9); err != nil {
		return err
	}
	if err := w.ReadSwitch((*uint32)(&sw)); err != nil {
		return err
	}
	// ms_union
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	switch sw {
	case uint32(1):
		o.Value = &IISUserEnum_ConfigInfo_Level1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// IISUserEnum_ConfigInfo_Level1 structure represents IISUserEnum_ConfigInfo RPC union arm.
//
// It has following labels: 1
type IISUserEnum_ConfigInfo_Level1 struct {
	// Level1:  The pointer to an IIS_USER_INFO_1_CONTAINER structure that contains the
	// user information collection.
	Level1 *IISUserInfo1Container `idl:"name:Level1" json:"level1"`
}

func (*IISUserEnum_ConfigInfo_Level1) is_IISUserEnum_ConfigInfo() {}

func (o *IISUserEnum_ConfigInfo_Level1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Level1 != nil {
		_ptr_Level1 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Level1 != nil {
				if err := o.Level1.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&IISUserInfo1Container{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Level1, _ptr_Level1); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *IISUserEnum_ConfigInfo_Level1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_Level1 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Level1 == nil {
			o.Level1 = &IISUserInfo1Container{}
		}
		if err := o.Level1.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_Level1 := func(ptr interface{}) { o.Level1 = *ptr.(**IISUserInfo1Container) }
	if err := w.ReadPointer(&o.Level1, _s_Level1, _ptr_Level1); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultInetinfoClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultInetinfoClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...dcerpc.CallOption) (*GetVersionResponse, error) {
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

func (o *xxx_DefaultInetinfoClient) GetAdminInformation(ctx context.Context, in *GetAdminInformationRequest, opts ...dcerpc.CallOption) (*GetAdminInformationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetAdminInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInetinfoClient) GetSites(ctx context.Context, in *GetSitesRequest, opts ...dcerpc.CallOption) (*GetSitesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSitesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInetinfoClient) SetAdminInformation(ctx context.Context, in *SetAdminInformationRequest, opts ...dcerpc.CallOption) (*SetAdminInformationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetAdminInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInetinfoClient) GetGlobalAdminInformation(ctx context.Context, in *GetGlobalAdminInformationRequest, opts ...dcerpc.CallOption) (*GetGlobalAdminInformationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetGlobalAdminInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInetinfoClient) SetGlobalAdminInformation(ctx context.Context, in *SetGlobalAdminInformationRequest, opts ...dcerpc.CallOption) (*SetGlobalAdminInformationResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetGlobalAdminInformationResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInetinfoClient) QueryStatistics(ctx context.Context, in *QueryStatisticsRequest, opts ...dcerpc.CallOption) (*QueryStatisticsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &QueryStatisticsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInetinfoClient) ClearStatistics(ctx context.Context, in *ClearStatisticsRequest, opts ...dcerpc.CallOption) (*ClearStatisticsResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ClearStatisticsResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInetinfoClient) FlushMemoryCache(ctx context.Context, in *FlushMemoryCacheRequest, opts ...dcerpc.CallOption) (*FlushMemoryCacheResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FlushMemoryCacheResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInetinfoClient) GetServerCapabilities(ctx context.Context, in *GetServerCapabilitiesRequest, opts ...dcerpc.CallOption) (*GetServerCapabilitiesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetServerCapabilitiesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInetinfoClient) W3QueryStatistics2(ctx context.Context, in *W3QueryStatistics2Request, opts ...dcerpc.CallOption) (*W3QueryStatistics2Response, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &W3QueryStatistics2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInetinfoClient) W3ClearStatistics2(ctx context.Context, in *W3ClearStatistics2Request, opts ...dcerpc.CallOption) (*W3ClearStatistics2Response, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &W3ClearStatistics2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInetinfoClient) FTPQueryStatistics2(ctx context.Context, in *FTPQueryStatistics2Request, opts ...dcerpc.CallOption) (*FTPQueryStatistics2Response, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FTPQueryStatistics2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInetinfoClient) FTPClearStatistics2(ctx context.Context, in *FTPClearStatistics2Request, opts ...dcerpc.CallOption) (*FTPClearStatistics2Response, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FTPClearStatistics2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInetinfoClient) IISEnumerateUsers(ctx context.Context, in *IISEnumerateUsersRequest, opts ...dcerpc.CallOption) (*IISEnumerateUsersResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &IISEnumerateUsersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInetinfoClient) IISDisconnectUser(ctx context.Context, in *IISDisconnectUserRequest, opts ...dcerpc.CallOption) (*IISDisconnectUserResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &IISDisconnectUserResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultInetinfoClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultInetinfoClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewInetinfoClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (InetinfoClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(InetinfoSyntaxV2_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultInetinfoClient{cc: cc}, nil
}

// xxx_GetVersionOperation structure represents the R_InetInfoGetVersion operation
type xxx_GetVersionOperation struct {
	Server  string `idl:"name:pszServer;string;pointer:unique" json:"server"`
	_       uint32 `idl:"name:dwReserved"`
	Version uint32 `idl:"name:pdwVersion" json:"version"`
	Return  uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetVersionOperation) OpNum() int { return 0 }

func (o *xxx_GetVersionOperation) OpName() string { return "/inetinfo/v2/R_InetInfoGetVersion" }

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
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Server != "" {
			_ptr_pszServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Server); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_pszServer); err != nil {
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
	// dwReserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved dwReserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetVersionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Server); err != nil {
				return err
			}
			return nil
		})
		_s_pszServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
		if err := w.ReadPointer(&o.Server, _s_pszServer, _ptr_pszServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwReserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved dwReserved
		var _dwReserved uint32
		if err := w.ReadData(&_dwReserved); err != nil {
			return err
		}
	}
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
	// pdwVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Version); err != nil {
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

func (o *xxx_GetVersionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pdwVersion {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Version); err != nil {
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

// GetVersionRequest structure represents the R_InetInfoGetVersion operation request
type GetVersionRequest struct {
	// pszServer: A custom binding handle for the target system, as specified in section
	// 2.1.1. The value MUST NOT be used by the server implementation.
	Server string `idl:"name:pszServer;string;pointer:unique" json:"server"`
}

func (o *GetVersionRequest) xxx_ToOp(ctx context.Context, op *xxx_GetVersionOperation) *xxx_GetVersionOperation {
	if op == nil {
		op = &xxx_GetVersionOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	return op
}

func (o *GetVersionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetVersionOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
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

// GetVersionResponse structure represents the R_InetInfoGetVersion operation response
type GetVersionResponse struct {
	// pdwVersion: A pointer to a variable. On successful return, it MUST contain a major
	// and minor version number for the server implementation. The major version is stored
	// in the low WORD, and the minor version is stored in the high WORD.
	Version uint32 `idl:"name:pdwVersion" json:"version"`
	// Return: The R_InetInfoGetVersion return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetVersionResponse) xxx_ToOp(ctx context.Context, op *xxx_GetVersionOperation) *xxx_GetVersionOperation {
	if op == nil {
		op = &xxx_GetVersionOperation{}
	}
	if o == nil {
		return op
	}
	op.Version = o.Version
	op.Return = o.Return
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

// xxx_GetAdminInformationOperation structure represents the R_InetInfoGetAdminInformation operation
type xxx_GetAdminInformationOperation struct {
	Server     string      `idl:"name:pszServer;string;pointer:unique" json:"server"`
	ServerMask uint32      `idl:"name:dwServerMask" json:"server_mask"`
	Config     *ConfigInfo `idl:"name:ppConfig" json:"config"`
	Return     uint32      `idl:"name:Return" json:"return"`
}

func (o *xxx_GetAdminInformationOperation) OpNum() int { return 1 }

func (o *xxx_GetAdminInformationOperation) OpName() string {
	return "/inetinfo/v2/R_InetInfoGetAdminInformation"
}

func (o *xxx_GetAdminInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAdminInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Server != "" {
			_ptr_pszServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Server); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_pszServer); err != nil {
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
	// dwServerMask {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ServerMask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAdminInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Server); err != nil {
				return err
			}
			return nil
		})
		_s_pszServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
		if err := w.ReadPointer(&o.Server, _s_pszServer, _ptr_pszServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwServerMask {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ServerMask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAdminInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetAdminInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppConfig {out} (1:{pointer=ref}*(2))(2:{alias=LPINET_INFO_CONFIG_INFO}*(1))(3:{alias=INET_INFO_CONFIG_INFO}(struct))
	{
		if o.Config != nil {
			_ptr_ppConfig := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Config != nil {
					if err := o.Config.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&ConfigInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Config, _ptr_ppConfig); err != nil {
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

func (o *xxx_GetAdminInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppConfig {out} (1:{pointer=ref}*(2))(2:{alias=LPINET_INFO_CONFIG_INFO,pointer=ref}*(1))(3:{alias=INET_INFO_CONFIG_INFO}(struct))
	{
		_ptr_ppConfig := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Config == nil {
				o.Config = &ConfigInfo{}
			}
			if err := o.Config.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppConfig := func(ptr interface{}) { o.Config = *ptr.(**ConfigInfo) }
		if err := w.ReadPointer(&o.Config, _s_ppConfig, _ptr_ppConfig); err != nil {
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

// GetAdminInformationRequest structure represents the R_InetInfoGetAdminInformation operation request
type GetAdminInformationRequest struct {
	// pszServer: The custom binding handle for the target system, as specified in section
	// 2.1.1. The value MUST NOT be used by the server implementation.
	Server string `idl:"name:pszServer;string;pointer:unique" json:"server"`
	// dwServerMask: The identifier for the target Internet protocol server, as specified
	// in section 2.2.2.
	ServerMask uint32 `idl:"name:dwServerMask" json:"server_mask"`
}

func (o *GetAdminInformationRequest) xxx_ToOp(ctx context.Context, op *xxx_GetAdminInformationOperation) *xxx_GetAdminInformationOperation {
	if op == nil {
		op = &xxx_GetAdminInformationOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.ServerMask = o.ServerMask
	return op
}

func (o *GetAdminInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetAdminInformationOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.ServerMask = op.ServerMask
}
func (o *GetAdminInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetAdminInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAdminInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetAdminInformationResponse structure represents the R_InetInfoGetAdminInformation operation response
type GetAdminInformationResponse struct {
	// ppConfig: The pointer to a pointer to an INET_INFO_CONFIG_INFO structure that contains
	// configuration data for the specified Internet protocol server.
	Config *ConfigInfo `idl:"name:ppConfig" json:"config"`
	// Return: The R_InetInfoGetAdminInformation return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetAdminInformationResponse) xxx_ToOp(ctx context.Context, op *xxx_GetAdminInformationOperation) *xxx_GetAdminInformationOperation {
	if op == nil {
		op = &xxx_GetAdminInformationOperation{}
	}
	if o == nil {
		return op
	}
	op.Config = o.Config
	op.Return = o.Return
	return op
}

func (o *GetAdminInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetAdminInformationOperation) {
	if o == nil {
		return
	}
	o.Config = op.Config
	o.Return = op.Return
}
func (o *GetAdminInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetAdminInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetAdminInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSitesOperation structure represents the R_InetInfoGetSites operation
type xxx_GetSitesOperation struct {
	Server     string    `idl:"name:pszServer;string;pointer:unique" json:"server"`
	ServerMask uint32    `idl:"name:dwServerMask" json:"server_mask"`
	Sites      *SiteList `idl:"name:ppSites" json:"sites"`
	Return     uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSitesOperation) OpNum() int { return 2 }

func (o *xxx_GetSitesOperation) OpName() string { return "/inetinfo/v2/R_InetInfoGetSites" }

func (o *xxx_GetSitesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSitesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Server != "" {
			_ptr_pszServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Server); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_pszServer); err != nil {
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
	// dwServerMask {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ServerMask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSitesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Server); err != nil {
				return err
			}
			return nil
		})
		_s_pszServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
		if err := w.ReadPointer(&o.Server, _s_pszServer, _ptr_pszServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwServerMask {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ServerMask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSitesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSitesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppSites {out} (1:{pointer=ref}*(2))(2:{alias=LPINET_INFO_SITE_LIST}*(1))(3:{alias=INET_INFO_SITE_LIST}(struct))
	{
		if o.Sites != nil {
			_ptr_ppSites := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Sites != nil {
					if err := o.Sites.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&SiteList{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Sites, _ptr_ppSites); err != nil {
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

func (o *xxx_GetSitesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppSites {out} (1:{pointer=ref}*(2))(2:{alias=LPINET_INFO_SITE_LIST,pointer=ref}*(1))(3:{alias=INET_INFO_SITE_LIST}(struct))
	{
		_ptr_ppSites := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Sites == nil {
				o.Sites = &SiteList{}
			}
			if err := o.Sites.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppSites := func(ptr interface{}) { o.Sites = *ptr.(**SiteList) }
		if err := w.ReadPointer(&o.Sites, _s_ppSites, _ptr_ppSites); err != nil {
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

// GetSitesRequest structure represents the R_InetInfoGetSites operation request
type GetSitesRequest struct {
	// pszServer: The custom binding handle for the target system, as specified in section
	// 2.1.1. The value MUST NOT be used by the server implementation.
	Server string `idl:"name:pszServer;string;pointer:unique" json:"server"`
	// dwServerMask: The identifier for the target Internet protocol server, as specified
	// in section 2.2.2.
	ServerMask uint32 `idl:"name:dwServerMask" json:"server_mask"`
}

func (o *GetSitesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSitesOperation) *xxx_GetSitesOperation {
	if op == nil {
		op = &xxx_GetSitesOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.ServerMask = o.ServerMask
	return op
}

func (o *GetSitesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSitesOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.ServerMask = op.ServerMask
}
func (o *GetSitesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSitesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSitesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSitesResponse structure represents the R_InetInfoGetSites operation response
type GetSitesResponse struct {
	// ppSites: The pointer to a pointer to INET_INFO_SITE_LIST that specifies the list
	// of defined server instances for the Internet protocol server specified by dwServerMask.
	Sites *SiteList `idl:"name:ppSites" json:"sites"`
	// Return: The R_InetInfoGetSites return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetSitesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSitesOperation) *xxx_GetSitesOperation {
	if op == nil {
		op = &xxx_GetSitesOperation{}
	}
	if o == nil {
		return op
	}
	op.Sites = o.Sites
	op.Return = o.Return
	return op
}

func (o *GetSitesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSitesOperation) {
	if o == nil {
		return
	}
	o.Sites = op.Sites
	o.Return = op.Return
}
func (o *GetSitesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSitesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSitesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetAdminInformationOperation structure represents the R_InetInfoSetAdminInformation operation
type xxx_SetAdminInformationOperation struct {
	Server     string      `idl:"name:pszServer;string;pointer:unique" json:"server"`
	ServerMask uint32      `idl:"name:dwServerMask" json:"server_mask"`
	Config     *ConfigInfo `idl:"name:pConfig;pointer:ref" json:"config"`
	Return     uint32      `idl:"name:Return" json:"return"`
}

func (o *xxx_SetAdminInformationOperation) OpNum() int { return 3 }

func (o *xxx_SetAdminInformationOperation) OpName() string {
	return "/inetinfo/v2/R_InetInfoSetAdminInformation"
}

func (o *xxx_SetAdminInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAdminInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Server != "" {
			_ptr_pszServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Server); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_pszServer); err != nil {
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
	// dwServerMask {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ServerMask); err != nil {
			return err
		}
	}
	// pConfig {in} (1:{pointer=ref}*(1))(2:{alias=INET_INFO_CONFIG_INFO}(struct))
	{
		if o.Config != nil {
			if err := o.Config.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&ConfigInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAdminInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Server); err != nil {
				return err
			}
			return nil
		})
		_s_pszServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
		if err := w.ReadPointer(&o.Server, _s_pszServer, _ptr_pszServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwServerMask {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ServerMask); err != nil {
			return err
		}
	}
	// pConfig {in} (1:{pointer=ref}*(1))(2:{alias=INET_INFO_CONFIG_INFO}(struct))
	{
		if o.Config == nil {
			o.Config = &ConfigInfo{}
		}
		if err := o.Config.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAdminInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetAdminInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetAdminInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetAdminInformationRequest structure represents the R_InetInfoSetAdminInformation operation request
type SetAdminInformationRequest struct {
	// pszServer: The custom binding handle for the target system, as specified in section
	// 2.1.1. The value MUST NOT be used by the server implementation.
	Server string `idl:"name:pszServer;string;pointer:unique" json:"server"`
	// dwServerMask: The identifier for the target Internet protocol server, as specified
	// in section 2.2.2.
	ServerMask uint32 `idl:"name:dwServerMask" json:"server_mask"`
	// pConfig: The pointer to an INET_INFO_CONFIG_INFO structure containing the property
	// configuration to set. The client MUST set the appropriate flag in the FieldControl
	// member for any data field in pConfig that is to be set by the server.
	Config *ConfigInfo `idl:"name:pConfig;pointer:ref" json:"config"`
}

func (o *SetAdminInformationRequest) xxx_ToOp(ctx context.Context, op *xxx_SetAdminInformationOperation) *xxx_SetAdminInformationOperation {
	if op == nil {
		op = &xxx_SetAdminInformationOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.ServerMask = o.ServerMask
	op.Config = o.Config
	return op
}

func (o *SetAdminInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_SetAdminInformationOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.ServerMask = op.ServerMask
	o.Config = op.Config
}
func (o *SetAdminInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetAdminInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAdminInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetAdminInformationResponse structure represents the R_InetInfoSetAdminInformation operation response
type SetAdminInformationResponse struct {
	// Return: The R_InetInfoSetAdminInformation return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetAdminInformationResponse) xxx_ToOp(ctx context.Context, op *xxx_SetAdminInformationOperation) *xxx_SetAdminInformationOperation {
	if op == nil {
		op = &xxx_SetAdminInformationOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SetAdminInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_SetAdminInformationOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetAdminInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetAdminInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetAdminInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetGlobalAdminInformationOperation structure represents the R_InetInfoGetGlobalAdminInformation operation
type xxx_GetGlobalAdminInformationOperation struct {
	Server     string            `idl:"name:pszServer;string;pointer:unique" json:"server"`
	ServerMask uint32            `idl:"name:dwServerMask" json:"server_mask"`
	Config     *GlobalConfigInfo `idl:"name:ppConfig" json:"config"`
	Return     uint32            `idl:"name:Return" json:"return"`
}

func (o *xxx_GetGlobalAdminInformationOperation) OpNum() int { return 4 }

func (o *xxx_GetGlobalAdminInformationOperation) OpName() string {
	return "/inetinfo/v2/R_InetInfoGetGlobalAdminInformation"
}

func (o *xxx_GetGlobalAdminInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGlobalAdminInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Server != "" {
			_ptr_pszServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Server); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_pszServer); err != nil {
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
	// dwServerMask {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ServerMask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGlobalAdminInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Server); err != nil {
				return err
			}
			return nil
		})
		_s_pszServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
		if err := w.ReadPointer(&o.Server, _s_pszServer, _ptr_pszServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwServerMask {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ServerMask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGlobalAdminInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetGlobalAdminInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppConfig {out} (1:{pointer=ref}*(2))(2:{alias=LPINET_INFO_GLOBAL_CONFIG_INFO}*(1))(3:{alias=INET_INFO_GLOBAL_CONFIG_INFO}(struct))
	{
		if o.Config != nil {
			_ptr_ppConfig := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Config != nil {
					if err := o.Config.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&GlobalConfigInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Config, _ptr_ppConfig); err != nil {
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

func (o *xxx_GetGlobalAdminInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppConfig {out} (1:{pointer=ref}*(2))(2:{alias=LPINET_INFO_GLOBAL_CONFIG_INFO,pointer=ref}*(1))(3:{alias=INET_INFO_GLOBAL_CONFIG_INFO}(struct))
	{
		_ptr_ppConfig := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Config == nil {
				o.Config = &GlobalConfigInfo{}
			}
			if err := o.Config.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppConfig := func(ptr interface{}) { o.Config = *ptr.(**GlobalConfigInfo) }
		if err := w.ReadPointer(&o.Config, _s_ppConfig, _ptr_ppConfig); err != nil {
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

// GetGlobalAdminInformationRequest structure represents the R_InetInfoGetGlobalAdminInformation operation request
type GetGlobalAdminInformationRequest struct {
	// pszServer: The custom binding handle for the target system, as specified in section
	// 2.1.1. The value MUST NOT be used by the server implementation.
	Server string `idl:"name:pszServer;string;pointer:unique" json:"server"`
	// dwServerMask: This value MUST be ignored by the server.
	ServerMask uint32 `idl:"name:dwServerMask" json:"server_mask"`
}

func (o *GetGlobalAdminInformationRequest) xxx_ToOp(ctx context.Context, op *xxx_GetGlobalAdminInformationOperation) *xxx_GetGlobalAdminInformationOperation {
	if op == nil {
		op = &xxx_GetGlobalAdminInformationOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.ServerMask = o.ServerMask
	return op
}

func (o *GetGlobalAdminInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_GetGlobalAdminInformationOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.ServerMask = op.ServerMask
}
func (o *GetGlobalAdminInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetGlobalAdminInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetGlobalAdminInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetGlobalAdminInformationResponse structure represents the R_InetInfoGetGlobalAdminInformation operation response
type GetGlobalAdminInformationResponse struct {
	// ppConfig: The pointer to a pointer to an INET_INFO_GLOBAL_CONFIG_INFO structure.
	Config *GlobalConfigInfo `idl:"name:ppConfig" json:"config"`
	// Return: The R_InetInfoGetGlobalAdminInformation return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetGlobalAdminInformationResponse) xxx_ToOp(ctx context.Context, op *xxx_GetGlobalAdminInformationOperation) *xxx_GetGlobalAdminInformationOperation {
	if op == nil {
		op = &xxx_GetGlobalAdminInformationOperation{}
	}
	if o == nil {
		return op
	}
	op.Config = o.Config
	op.Return = o.Return
	return op
}

func (o *GetGlobalAdminInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_GetGlobalAdminInformationOperation) {
	if o == nil {
		return
	}
	o.Config = op.Config
	o.Return = op.Return
}
func (o *GetGlobalAdminInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetGlobalAdminInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetGlobalAdminInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetGlobalAdminInformationOperation structure represents the R_InetInfoSetGlobalAdminInformation operation
type xxx_SetGlobalAdminInformationOperation struct {
	Server     string            `idl:"name:pszServer;string;pointer:unique" json:"server"`
	ServerMask uint32            `idl:"name:dwServerMask" json:"server_mask"`
	Config     *GlobalConfigInfo `idl:"name:pConfig;pointer:ref" json:"config"`
	Return     uint32            `idl:"name:Return" json:"return"`
}

func (o *xxx_SetGlobalAdminInformationOperation) OpNum() int { return 5 }

func (o *xxx_SetGlobalAdminInformationOperation) OpName() string {
	return "/inetinfo/v2/R_InetInfoSetGlobalAdminInformation"
}

func (o *xxx_SetGlobalAdminInformationOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGlobalAdminInformationOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Server != "" {
			_ptr_pszServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Server); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_pszServer); err != nil {
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
	// dwServerMask {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ServerMask); err != nil {
			return err
		}
	}
	// pConfig {in} (1:{pointer=ref}*(1))(2:{alias=INET_INFO_GLOBAL_CONFIG_INFO}(struct))
	{
		if o.Config != nil {
			if err := o.Config.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&GlobalConfigInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_SetGlobalAdminInformationOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Server); err != nil {
				return err
			}
			return nil
		})
		_s_pszServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
		if err := w.ReadPointer(&o.Server, _s_pszServer, _ptr_pszServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwServerMask {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ServerMask); err != nil {
			return err
		}
	}
	// pConfig {in} (1:{pointer=ref}*(1))(2:{alias=INET_INFO_GLOBAL_CONFIG_INFO}(struct))
	{
		if o.Config == nil {
			o.Config = &GlobalConfigInfo{}
		}
		if err := o.Config.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGlobalAdminInformationOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetGlobalAdminInformationOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_SetGlobalAdminInformationOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetGlobalAdminInformationRequest structure represents the R_InetInfoSetGlobalAdminInformation operation request
type SetGlobalAdminInformationRequest struct {
	// pszServer: The custom binding handle for the target system, as specified in section
	// 2.1.1. The value MUST NOT be used by the server implementation.
	Server string `idl:"name:pszServer;string;pointer:unique" json:"server"`
	// dwServerMask: This value MUST be ignored by the server.
	ServerMask uint32 `idl:"name:dwServerMask" json:"server_mask"`
	// pConfig: The pointer to an INET_INFO_GLOBAL_CONFIG_INFO structure that contains global
	// administrative information.
	Config *GlobalConfigInfo `idl:"name:pConfig;pointer:ref" json:"config"`
}

func (o *SetGlobalAdminInformationRequest) xxx_ToOp(ctx context.Context, op *xxx_SetGlobalAdminInformationOperation) *xxx_SetGlobalAdminInformationOperation {
	if op == nil {
		op = &xxx_SetGlobalAdminInformationOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.ServerMask = o.ServerMask
	op.Config = o.Config
	return op
}

func (o *SetGlobalAdminInformationRequest) xxx_FromOp(ctx context.Context, op *xxx_SetGlobalAdminInformationOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.ServerMask = op.ServerMask
	o.Config = op.Config
}
func (o *SetGlobalAdminInformationRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetGlobalAdminInformationRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetGlobalAdminInformationOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetGlobalAdminInformationResponse structure represents the R_InetInfoSetGlobalAdminInformation operation response
type SetGlobalAdminInformationResponse struct {
	// Return: The R_InetInfoSetGlobalAdminInformation return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetGlobalAdminInformationResponse) xxx_ToOp(ctx context.Context, op *xxx_SetGlobalAdminInformationOperation) *xxx_SetGlobalAdminInformationOperation {
	if op == nil {
		op = &xxx_SetGlobalAdminInformationOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SetGlobalAdminInformationResponse) xxx_FromOp(ctx context.Context, op *xxx_SetGlobalAdminInformationOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetGlobalAdminInformationResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetGlobalAdminInformationResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetGlobalAdminInformationOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_QueryStatisticsOperation structure represents the R_InetInfoQueryStatistics operation
type xxx_QueryStatisticsOperation struct {
	Server     string          `idl:"name:pszServer;string;pointer:unique" json:"server"`
	Level      uint32          `idl:"name:Level" json:"level"`
	ServerMask uint32          `idl:"name:dwServerMask" json:"server_mask"`
	StatsInfo  *StatisticsInfo `idl:"name:StatsInfo;switch_is:Level" json:"stats_info"`
	Return     uint32          `idl:"name:Return" json:"return"`
}

func (o *xxx_QueryStatisticsOperation) OpNum() int { return 6 }

func (o *xxx_QueryStatisticsOperation) OpName() string {
	return "/inetinfo/v2/R_InetInfoQueryStatistics"
}

func (o *xxx_QueryStatisticsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryStatisticsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Server != "" {
			_ptr_pszServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Server); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_pszServer); err != nil {
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
	// Level {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// dwServerMask {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ServerMask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryStatisticsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Server); err != nil {
				return err
			}
			return nil
		})
		_s_pszServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
		if err := w.ReadPointer(&o.Server, _s_pszServer, _ptr_pszServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Level {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// dwServerMask {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ServerMask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryStatisticsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_QueryStatisticsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// StatsInfo {out} (1:{switch_type={}(uint32), alias=LPINET_INFO_STATISTICS_INFO}*(1))(2:{switch_type={}(uint32), alias=INET_INFO_STATISTICS_INFO}(union))
	{
		_swStatsInfo := uint32(o.Level)
		if o.StatsInfo != nil {
			if err := o.StatsInfo.MarshalUnionNDR(ctx, w, _swStatsInfo); err != nil {
				return err
			}
		} else {
			if err := (&StatisticsInfo{}).MarshalUnionNDR(ctx, w, _swStatsInfo); err != nil {
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

func (o *xxx_QueryStatisticsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// StatsInfo {out} (1:{switch_type={}(uint32), alias=LPINET_INFO_STATISTICS_INFO,pointer=ref}*(1))(2:{switch_type={}(uint32), alias=INET_INFO_STATISTICS_INFO}(union))
	{
		if o.StatsInfo == nil {
			o.StatsInfo = &StatisticsInfo{}
		}
		_swStatsInfo := uint32(o.Level)
		if err := o.StatsInfo.UnmarshalUnionNDR(ctx, w, _swStatsInfo); err != nil {
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

// QueryStatisticsRequest structure represents the R_InetInfoQueryStatistics operation request
type QueryStatisticsRequest struct {
	// pszServer: The custom binding handle for the target system, as specified in section
	// 2.1.1. The value MUST NOT be used by the server implementation.
	Server string `idl:"name:pszServer;string;pointer:unique" json:"server"`
	// Level: The level of detail to be retrieved. This member MUST be set to 0. If another
	// value is sent by the client, the server MUST return ERROR_INVALID_LEVEL (0x0000007C).
	Level uint32 `idl:"name:Level" json:"level"`
	// dwServerMask:  The identifier for the target Internet protocol server, as specified
	// in section 2.2.2. A value of 0 indicates that aggregate statistical data is to be
	// returned for all protocol servers.
	ServerMask uint32 `idl:"name:dwServerMask" json:"server_mask"`
}

func (o *QueryStatisticsRequest) xxx_ToOp(ctx context.Context, op *xxx_QueryStatisticsOperation) *xxx_QueryStatisticsOperation {
	if op == nil {
		op = &xxx_QueryStatisticsOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.Level = o.Level
	op.ServerMask = o.ServerMask
	return op
}

func (o *QueryStatisticsRequest) xxx_FromOp(ctx context.Context, op *xxx_QueryStatisticsOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.Level = op.Level
	o.ServerMask = op.ServerMask
}
func (o *QueryStatisticsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *QueryStatisticsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryStatisticsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// QueryStatisticsResponse structure represents the R_InetInfoQueryStatistics operation response
type QueryStatisticsResponse struct {
	// XXX: Level is an implicit input depedency for output parameters
	Level uint32 `idl:"name:Level" json:"level"`
	// StatsInfo: The pointer to an INET_INFO_STATISTICS_INFO union that contains the data
	// to be returned.
	StatsInfo *StatisticsInfo `idl:"name:StatsInfo;switch_is:Level" json:"stats_info"`
	// Return: The R_InetInfoQueryStatistics return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *QueryStatisticsResponse) xxx_ToOp(ctx context.Context, op *xxx_QueryStatisticsOperation) *xxx_QueryStatisticsOperation {
	if op == nil {
		op = &xxx_QueryStatisticsOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.Level == uint32(0) {
		op.Level = o.Level
	}

	op.StatsInfo = o.StatsInfo
	op.Return = o.Return
	return op
}

func (o *QueryStatisticsResponse) xxx_FromOp(ctx context.Context, op *xxx_QueryStatisticsOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.Level = op.Level

	o.StatsInfo = op.StatsInfo
	o.Return = op.Return
}
func (o *QueryStatisticsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *QueryStatisticsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_QueryStatisticsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ClearStatisticsOperation structure represents the R_InetInfoClearStatistics operation
type xxx_ClearStatisticsOperation struct {
	Server     string `idl:"name:pszServer;string;pointer:unique" json:"server"`
	ServerMask uint32 `idl:"name:dwServerMask" json:"server_mask"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_ClearStatisticsOperation) OpNum() int { return 7 }

func (o *xxx_ClearStatisticsOperation) OpName() string {
	return "/inetinfo/v2/R_InetInfoClearStatistics"
}

func (o *xxx_ClearStatisticsOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearStatisticsOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Server != "" {
			_ptr_pszServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Server); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_pszServer); err != nil {
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
	// dwServerMask {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ServerMask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearStatisticsOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Server); err != nil {
				return err
			}
			return nil
		})
		_s_pszServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
		if err := w.ReadPointer(&o.Server, _s_pszServer, _ptr_pszServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwServerMask {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ServerMask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearStatisticsOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ClearStatisticsOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ClearStatisticsOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ClearStatisticsRequest structure represents the R_InetInfoClearStatistics operation request
type ClearStatisticsRequest struct {
	// pszServer: The custom binding handle for the target system, as specified in section
	// 2.1.1. The value MUST NOT be used by the server implementation.
	Server string `idl:"name:pszServer;string;pointer:unique" json:"server"`
	// dwServerMask: The identifier for the target Internet protocol server, as specified
	// in section 2.2.2.
	ServerMask uint32 `idl:"name:dwServerMask" json:"server_mask"`
}

func (o *ClearStatisticsRequest) xxx_ToOp(ctx context.Context, op *xxx_ClearStatisticsOperation) *xxx_ClearStatisticsOperation {
	if op == nil {
		op = &xxx_ClearStatisticsOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.ServerMask = o.ServerMask
	return op
}

func (o *ClearStatisticsRequest) xxx_FromOp(ctx context.Context, op *xxx_ClearStatisticsOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.ServerMask = op.ServerMask
}
func (o *ClearStatisticsRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ClearStatisticsRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearStatisticsOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ClearStatisticsResponse structure represents the R_InetInfoClearStatistics operation response
type ClearStatisticsResponse struct {
	// Return: The R_InetInfoClearStatistics return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ClearStatisticsResponse) xxx_ToOp(ctx context.Context, op *xxx_ClearStatisticsOperation) *xxx_ClearStatisticsOperation {
	if op == nil {
		op = &xxx_ClearStatisticsOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *ClearStatisticsResponse) xxx_FromOp(ctx context.Context, op *xxx_ClearStatisticsOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ClearStatisticsResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ClearStatisticsResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ClearStatisticsOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FlushMemoryCacheOperation structure represents the R_InetInfoFlushMemoryCache operation
type xxx_FlushMemoryCacheOperation struct {
	Server     string `idl:"name:pszServer;string;pointer:unique" json:"server"`
	ServerMask uint32 `idl:"name:dwServerMask" json:"server_mask"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_FlushMemoryCacheOperation) OpNum() int { return 8 }

func (o *xxx_FlushMemoryCacheOperation) OpName() string {
	return "/inetinfo/v2/R_InetInfoFlushMemoryCache"
}

func (o *xxx_FlushMemoryCacheOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FlushMemoryCacheOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Server != "" {
			_ptr_pszServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Server); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_pszServer); err != nil {
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
	// dwServerMask {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ServerMask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FlushMemoryCacheOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Server); err != nil {
				return err
			}
			return nil
		})
		_s_pszServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
		if err := w.ReadPointer(&o.Server, _s_pszServer, _ptr_pszServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwServerMask {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ServerMask); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FlushMemoryCacheOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FlushMemoryCacheOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_FlushMemoryCacheOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FlushMemoryCacheRequest structure represents the R_InetInfoFlushMemoryCache operation request
type FlushMemoryCacheRequest struct {
	// pszServer: The custom binding handle for the target system, as specified in section
	// 2.1.1. The value MUST NOT be used by the server implementation.
	Server string `idl:"name:pszServer;string;pointer:unique" json:"server"`
	// dwServerMask: The identifier for the target Internet protocol server, as specified
	// in section 2.2.2.
	ServerMask uint32 `idl:"name:dwServerMask" json:"server_mask"`
}

func (o *FlushMemoryCacheRequest) xxx_ToOp(ctx context.Context, op *xxx_FlushMemoryCacheOperation) *xxx_FlushMemoryCacheOperation {
	if op == nil {
		op = &xxx_FlushMemoryCacheOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.ServerMask = o.ServerMask
	return op
}

func (o *FlushMemoryCacheRequest) xxx_FromOp(ctx context.Context, op *xxx_FlushMemoryCacheOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.ServerMask = op.ServerMask
}
func (o *FlushMemoryCacheRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FlushMemoryCacheRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FlushMemoryCacheOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FlushMemoryCacheResponse structure represents the R_InetInfoFlushMemoryCache operation response
type FlushMemoryCacheResponse struct {
	// Return: The R_InetInfoFlushMemoryCache return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *FlushMemoryCacheResponse) xxx_ToOp(ctx context.Context, op *xxx_FlushMemoryCacheOperation) *xxx_FlushMemoryCacheOperation {
	if op == nil {
		op = &xxx_FlushMemoryCacheOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *FlushMemoryCacheResponse) xxx_FromOp(ctx context.Context, op *xxx_FlushMemoryCacheOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *FlushMemoryCacheResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FlushMemoryCacheResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FlushMemoryCacheOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetServerCapabilitiesOperation structure represents the R_InetInfoGetServerCapabilities operation
type xxx_GetServerCapabilitiesOperation struct {
	Server string        `idl:"name:pszServer;string;pointer:unique" json:"server"`
	_      uint32        `idl:"name:dwReserved"`
	Cap    *Capabilities `idl:"name:ppCap" json:"cap"`
	Return uint32        `idl:"name:Return" json:"return"`
}

func (o *xxx_GetServerCapabilitiesOperation) OpNum() int { return 9 }

func (o *xxx_GetServerCapabilitiesOperation) OpName() string {
	return "/inetinfo/v2/R_InetInfoGetServerCapabilities"
}

func (o *xxx_GetServerCapabilitiesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerCapabilitiesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Server != "" {
			_ptr_pszServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Server); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_pszServer); err != nil {
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
	// dwReserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved dwReserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerCapabilitiesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Server); err != nil {
				return err
			}
			return nil
		})
		_s_pszServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
		if err := w.ReadPointer(&o.Server, _s_pszServer, _ptr_pszServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwReserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved dwReserved
		var _dwReserved uint32
		if err := w.ReadData(&_dwReserved); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerCapabilitiesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetServerCapabilitiesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppCap {out} (1:{pointer=ref}*(2))(2:{alias=LPINET_INFO_CAPABILITIES_STRUCT}*(1))(3:{alias=INET_INFO_CAPABILITIES_STRUCT}(struct))
	{
		if o.Cap != nil {
			_ptr_ppCap := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Cap != nil {
					if err := o.Cap.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&Capabilities{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Cap, _ptr_ppCap); err != nil {
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

func (o *xxx_GetServerCapabilitiesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppCap {out} (1:{pointer=ref}*(2))(2:{alias=LPINET_INFO_CAPABILITIES_STRUCT,pointer=ref}*(1))(3:{alias=INET_INFO_CAPABILITIES_STRUCT}(struct))
	{
		_ptr_ppCap := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Cap == nil {
				o.Cap = &Capabilities{}
			}
			if err := o.Cap.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_ppCap := func(ptr interface{}) { o.Cap = *ptr.(**Capabilities) }
		if err := w.ReadPointer(&o.Cap, _s_ppCap, _ptr_ppCap); err != nil {
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

// GetServerCapabilitiesRequest structure represents the R_InetInfoGetServerCapabilities operation request
type GetServerCapabilitiesRequest struct {
	// pszServer: The custom binding handle for the target system, as specified in section
	// 2.1.1. The value MUST NOT be used by the server implementation.
	Server string `idl:"name:pszServer;string;pointer:unique" json:"server"`
}

func (o *GetServerCapabilitiesRequest) xxx_ToOp(ctx context.Context, op *xxx_GetServerCapabilitiesOperation) *xxx_GetServerCapabilitiesOperation {
	if op == nil {
		op = &xxx_GetServerCapabilitiesOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	return op
}

func (o *GetServerCapabilitiesRequest) xxx_FromOp(ctx context.Context, op *xxx_GetServerCapabilitiesOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
}
func (o *GetServerCapabilitiesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetServerCapabilitiesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServerCapabilitiesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetServerCapabilitiesResponse structure represents the R_InetInfoGetServerCapabilities operation response
type GetServerCapabilitiesResponse struct {
	// ppCap: The pointer to an INET_INFO_CAPABILITIES_STRUCT structure that indicates the
	// capabilities of the Internet protocol servers on the host system.
	Cap *Capabilities `idl:"name:ppCap" json:"cap"`
	// Return: The R_InetInfoGetServerCapabilities return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetServerCapabilitiesResponse) xxx_ToOp(ctx context.Context, op *xxx_GetServerCapabilitiesOperation) *xxx_GetServerCapabilitiesOperation {
	if op == nil {
		op = &xxx_GetServerCapabilitiesOperation{}
	}
	if o == nil {
		return op
	}
	op.Cap = o.Cap
	op.Return = o.Return
	return op
}

func (o *GetServerCapabilitiesResponse) xxx_FromOp(ctx context.Context, op *xxx_GetServerCapabilitiesOperation) {
	if o == nil {
		return
	}
	o.Cap = op.Cap
	o.Return = op.Return
}
func (o *GetServerCapabilitiesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetServerCapabilitiesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetServerCapabilitiesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_W3QueryStatistics2Operation structure represents the R_W3QueryStatistics2 operation
type xxx_W3QueryStatistics2Operation struct {
	Server   string        `idl:"name:pszServer;string;pointer:unique" json:"server"`
	Level    uint32        `idl:"name:dwLevel" json:"level"`
	Instance uint32        `idl:"name:dwInstance" json:"instance"`
	_        uint32        `idl:"name:dwReserved"`
	Info     *W3Statistics `idl:"name:InfoStruct;switch_is:dwLevel" json:"info"`
	Return   uint32        `idl:"name:Return" json:"return"`
}

func (o *xxx_W3QueryStatistics2Operation) OpNum() int { return 10 }

func (o *xxx_W3QueryStatistics2Operation) OpName() string { return "/inetinfo/v2/R_W3QueryStatistics2" }

func (o *xxx_W3QueryStatistics2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_W3QueryStatistics2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Server != "" {
			_ptr_pszServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Server); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_pszServer); err != nil {
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
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// dwInstance {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Instance); err != nil {
			return err
		}
	}
	// dwReserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved dwReserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_W3QueryStatistics2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Server); err != nil {
				return err
			}
			return nil
		})
		_s_pszServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
		if err := w.ReadPointer(&o.Server, _s_pszServer, _ptr_pszServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// dwInstance {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Instance); err != nil {
			return err
		}
	}
	// dwReserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved dwReserved
		var _dwReserved uint32
		if err := w.ReadData(&_dwReserved); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_W3QueryStatistics2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_W3QueryStatistics2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// InfoStruct {out} (1:{switch_type={}(uint32), alias=LPW3_STATISTICS_STRUCT}*(1))(2:{switch_type={}(uint32), alias=W3_STATISTICS_STRUCT}(union))
	{
		_swInfo := uint32(o.Level)
		if o.Info != nil {
			if err := o.Info.MarshalUnionNDR(ctx, w, _swInfo); err != nil {
				return err
			}
		} else {
			if err := (&W3Statistics{}).MarshalUnionNDR(ctx, w, _swInfo); err != nil {
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

func (o *xxx_W3QueryStatistics2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// InfoStruct {out} (1:{switch_type={}(uint32), alias=LPW3_STATISTICS_STRUCT,pointer=ref}*(1))(2:{switch_type={}(uint32), alias=W3_STATISTICS_STRUCT}(union))
	{
		if o.Info == nil {
			o.Info = &W3Statistics{}
		}
		_swInfo := uint32(o.Level)
		if err := o.Info.UnmarshalUnionNDR(ctx, w, _swInfo); err != nil {
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

// W3QueryStatistics2Request structure represents the R_W3QueryStatistics2 operation request
type W3QueryStatistics2Request struct {
	// pszServer: The custom binding handle for the target system, as specified in section
	// 2.1.1. The value MUST NOT be used by the server implementation.
	Server string `idl:"name:pszServer;string;pointer:unique" json:"server"`
	// dwLevel: The level of detail to be retrieved. This parameter MUST be set to 0 by
	// the client. Other values MUST generate a return code of ERROR_INVALID_LEVEL (0x0000007C).
	Level uint32 `idl:"name:dwLevel" json:"level"`
	// dwInstance: The ID of the protocol server instance whose statistical data is being
	// requested. The following values have special meanings.
	//
	//	+-------------+--------------------------------------------------------------------+
	//	|             |                                                                    |
	//	|    VALUE    |                              MEANING                               |
	//	|             |                                                                    |
	//	+-------------+--------------------------------------------------------------------+
	//	+-------------+--------------------------------------------------------------------+
	//	| 0x000000000 | Return global (not per server instance) statistics data.           |
	//	+-------------+--------------------------------------------------------------------+
	//	| 0xf0000003  | Return statistics aggregated across all protocol server instances. |
	//	+-------------+--------------------------------------------------------------------+
	Instance uint32 `idl:"name:dwInstance" json:"instance"`
}

func (o *W3QueryStatistics2Request) xxx_ToOp(ctx context.Context, op *xxx_W3QueryStatistics2Operation) *xxx_W3QueryStatistics2Operation {
	if op == nil {
		op = &xxx_W3QueryStatistics2Operation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.Level = o.Level
	op.Instance = o.Instance
	return op
}

func (o *W3QueryStatistics2Request) xxx_FromOp(ctx context.Context, op *xxx_W3QueryStatistics2Operation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.Level = op.Level
	o.Instance = op.Instance
}
func (o *W3QueryStatistics2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *W3QueryStatistics2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_W3QueryStatistics2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// W3QueryStatistics2Response structure represents the R_W3QueryStatistics2 operation response
type W3QueryStatistics2Response struct {
	// XXX: dwLevel is an implicit input depedency for output parameters
	Level uint32 `idl:"name:dwLevel" json:"level"`
	// InfoStruct: The pointer to a W3_STATISTICS_STRUCT union to contain the retrieved
	// statistics data.
	Info *W3Statistics `idl:"name:InfoStruct;switch_is:dwLevel" json:"info"`
	// Return: The R_W3QueryStatistics2 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *W3QueryStatistics2Response) xxx_ToOp(ctx context.Context, op *xxx_W3QueryStatistics2Operation) *xxx_W3QueryStatistics2Operation {
	if op == nil {
		op = &xxx_W3QueryStatistics2Operation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.Level == uint32(0) {
		op.Level = o.Level
	}

	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *W3QueryStatistics2Response) xxx_FromOp(ctx context.Context, op *xxx_W3QueryStatistics2Operation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.Level = op.Level

	o.Info = op.Info
	o.Return = op.Return
}
func (o *W3QueryStatistics2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *W3QueryStatistics2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_W3QueryStatistics2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_W3ClearStatistics2Operation structure represents the R_W3ClearStatistics2 operation
type xxx_W3ClearStatistics2Operation struct {
	Server   string `idl:"name:pszServer;string;pointer:unique" json:"server"`
	Instance uint32 `idl:"name:dwInstance" json:"instance"`
	Return   uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_W3ClearStatistics2Operation) OpNum() int { return 11 }

func (o *xxx_W3ClearStatistics2Operation) OpName() string { return "/inetinfo/v2/R_W3ClearStatistics2" }

func (o *xxx_W3ClearStatistics2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_W3ClearStatistics2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Server != "" {
			_ptr_pszServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Server); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_pszServer); err != nil {
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
	// dwInstance {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Instance); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_W3ClearStatistics2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Server); err != nil {
				return err
			}
			return nil
		})
		_s_pszServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
		if err := w.ReadPointer(&o.Server, _s_pszServer, _ptr_pszServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwInstance {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Instance); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_W3ClearStatistics2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_W3ClearStatistics2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_W3ClearStatistics2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// W3ClearStatistics2Request structure represents the R_W3ClearStatistics2 operation request
type W3ClearStatistics2Request struct {
	// pszServer: The custom binding handle for the target system, as specified in section
	// 2.1.1. Value MUST NOT be used by the server implementation.
	Server string `idl:"name:pszServer;string;pointer:unique" json:"server"`
	// dwInstance: The ID of the protocol server instance whose statistical data is being
	// cleared.
	Instance uint32 `idl:"name:dwInstance" json:"instance"`
}

func (o *W3ClearStatistics2Request) xxx_ToOp(ctx context.Context, op *xxx_W3ClearStatistics2Operation) *xxx_W3ClearStatistics2Operation {
	if op == nil {
		op = &xxx_W3ClearStatistics2Operation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.Instance = o.Instance
	return op
}

func (o *W3ClearStatistics2Request) xxx_FromOp(ctx context.Context, op *xxx_W3ClearStatistics2Operation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.Instance = op.Instance
}
func (o *W3ClearStatistics2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *W3ClearStatistics2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_W3ClearStatistics2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// W3ClearStatistics2Response structure represents the R_W3ClearStatistics2 operation response
type W3ClearStatistics2Response struct {
	// Return: The R_W3ClearStatistics2 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *W3ClearStatistics2Response) xxx_ToOp(ctx context.Context, op *xxx_W3ClearStatistics2Operation) *xxx_W3ClearStatistics2Operation {
	if op == nil {
		op = &xxx_W3ClearStatistics2Operation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *W3ClearStatistics2Response) xxx_FromOp(ctx context.Context, op *xxx_W3ClearStatistics2Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *W3ClearStatistics2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *W3ClearStatistics2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_W3ClearStatistics2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FTPQueryStatistics2Operation structure represents the R_FtpQueryStatistics2 operation
type xxx_FTPQueryStatistics2Operation struct {
	Server   string         `idl:"name:pszServer;string;pointer:unique" json:"server"`
	Level    uint32         `idl:"name:dwLevel" json:"level"`
	Instance uint32         `idl:"name:dwInstance" json:"instance"`
	_        uint32         `idl:"name:dwReserved"`
	Info     *FTPStatistics `idl:"name:InfoStruct;switch_is:dwLevel" json:"info"`
	Return   uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_FTPQueryStatistics2Operation) OpNum() int { return 12 }

func (o *xxx_FTPQueryStatistics2Operation) OpName() string {
	return "/inetinfo/v2/R_FtpQueryStatistics2"
}

func (o *xxx_FTPQueryStatistics2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTPQueryStatistics2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Server != "" {
			_ptr_pszServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Server); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_pszServer); err != nil {
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
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// dwInstance {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Instance); err != nil {
			return err
		}
	}
	// dwReserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved dwReserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTPQueryStatistics2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Server); err != nil {
				return err
			}
			return nil
		})
		_s_pszServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
		if err := w.ReadPointer(&o.Server, _s_pszServer, _ptr_pszServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwLevel {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// dwInstance {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Instance); err != nil {
			return err
		}
	}
	// dwReserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved dwReserved
		var _dwReserved uint32
		if err := w.ReadData(&_dwReserved); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTPQueryStatistics2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTPQueryStatistics2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// InfoStruct {out} (1:{switch_type={}(uint32), alias=LPFTP_STATISTICS_STRUCT}*(1))(2:{switch_type={}(uint32), alias=FTP_STATISTICS_STRUCT}(union))
	{
		_swInfo := uint32(o.Level)
		if o.Info != nil {
			if err := o.Info.MarshalUnionNDR(ctx, w, _swInfo); err != nil {
				return err
			}
		} else {
			if err := (&FTPStatistics{}).MarshalUnionNDR(ctx, w, _swInfo); err != nil {
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

func (o *xxx_FTPQueryStatistics2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// InfoStruct {out} (1:{switch_type={}(uint32), alias=LPFTP_STATISTICS_STRUCT,pointer=ref}*(1))(2:{switch_type={}(uint32), alias=FTP_STATISTICS_STRUCT}(union))
	{
		if o.Info == nil {
			o.Info = &FTPStatistics{}
		}
		_swInfo := uint32(o.Level)
		if err := o.Info.UnmarshalUnionNDR(ctx, w, _swInfo); err != nil {
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

// FTPQueryStatistics2Request structure represents the R_FtpQueryStatistics2 operation request
type FTPQueryStatistics2Request struct {
	// pszServer: The custom binding handle for the target system, as specified in section
	// 2.1.1. The value MUST NOT be used by the server implementation.
	Server string `idl:"name:pszServer;string;pointer:unique" json:"server"`
	// dwLevel: The level of detail to be retrieved. This parameter MUST be set to 0 by
	// the client. Other values MUST generate a return code of ERROR_INVALID_LEVEL (0x0000007C).
	Level uint32 `idl:"name:dwLevel" json:"level"`
	// dwInstance: The ID of the protocol server instance whose statistical data is being
	// requested. The following values have special meanings.
	//
	//	+------------+--------------------------------------------------------------------+
	//	|            |                                                                    |
	//	|   VALUE    |                              MEANING                               |
	//	|            |                                                                    |
	//	+------------+--------------------------------------------------------------------+
	//	+------------+--------------------------------------------------------------------+
	//	| 0x00000000 | Return global (not per server instance) statistics data.           |
	//	+------------+--------------------------------------------------------------------+
	//	| 0xF0000003 | Return statistics aggregated across all protocol server instances. |
	//	+------------+--------------------------------------------------------------------+
	Instance uint32 `idl:"name:dwInstance" json:"instance"`
}

func (o *FTPQueryStatistics2Request) xxx_ToOp(ctx context.Context, op *xxx_FTPQueryStatistics2Operation) *xxx_FTPQueryStatistics2Operation {
	if op == nil {
		op = &xxx_FTPQueryStatistics2Operation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.Level = o.Level
	op.Instance = o.Instance
	return op
}

func (o *FTPQueryStatistics2Request) xxx_FromOp(ctx context.Context, op *xxx_FTPQueryStatistics2Operation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.Level = op.Level
	o.Instance = op.Instance
}
func (o *FTPQueryStatistics2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FTPQueryStatistics2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FTPQueryStatistics2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FTPQueryStatistics2Response structure represents the R_FtpQueryStatistics2 operation response
type FTPQueryStatistics2Response struct {
	// XXX: dwLevel is an implicit input depedency for output parameters
	Level uint32 `idl:"name:dwLevel" json:"level"`
	// InfoStruct: The pointer to an FTP_STATISTICS_STRUCT union to contain the retrieved
	// statistics data.
	Info *FTPStatistics `idl:"name:InfoStruct;switch_is:dwLevel" json:"info"`
	// Return: The R_FtpQueryStatistics2 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *FTPQueryStatistics2Response) xxx_ToOp(ctx context.Context, op *xxx_FTPQueryStatistics2Operation) *xxx_FTPQueryStatistics2Operation {
	if op == nil {
		op = &xxx_FTPQueryStatistics2Operation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.Level == uint32(0) {
		op.Level = o.Level
	}

	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *FTPQueryStatistics2Response) xxx_FromOp(ctx context.Context, op *xxx_FTPQueryStatistics2Operation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.Level = op.Level

	o.Info = op.Info
	o.Return = op.Return
}
func (o *FTPQueryStatistics2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FTPQueryStatistics2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FTPQueryStatistics2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FTPClearStatistics2Operation structure represents the R_FtpClearStatistics2 operation
type xxx_FTPClearStatistics2Operation struct {
	Server   string `idl:"name:pszServer;string;pointer:unique" json:"server"`
	Instance uint32 `idl:"name:dwInstance" json:"instance"`
	Return   uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_FTPClearStatistics2Operation) OpNum() int { return 13 }

func (o *xxx_FTPClearStatistics2Operation) OpName() string {
	return "/inetinfo/v2/R_FtpClearStatistics2"
}

func (o *xxx_FTPClearStatistics2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTPClearStatistics2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Server != "" {
			_ptr_pszServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Server); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_pszServer); err != nil {
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
	// dwInstance {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Instance); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTPClearStatistics2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Server); err != nil {
				return err
			}
			return nil
		})
		_s_pszServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
		if err := w.ReadPointer(&o.Server, _s_pszServer, _ptr_pszServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwInstance {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Instance); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTPClearStatistics2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FTPClearStatistics2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_FTPClearStatistics2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FTPClearStatistics2Request structure represents the R_FtpClearStatistics2 operation request
type FTPClearStatistics2Request struct {
	// pszServer: The custom binding handle for the target system, as specified in section
	// 2.1.1. The value MUST NOT be used by the server implementation.
	Server string `idl:"name:pszServer;string;pointer:unique" json:"server"`
	// dwInstance: The ID of the protocol server instance whose statistical data is being
	// cleared.
	Instance uint32 `idl:"name:dwInstance" json:"instance"`
}

func (o *FTPClearStatistics2Request) xxx_ToOp(ctx context.Context, op *xxx_FTPClearStatistics2Operation) *xxx_FTPClearStatistics2Operation {
	if op == nil {
		op = &xxx_FTPClearStatistics2Operation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.Instance = o.Instance
	return op
}

func (o *FTPClearStatistics2Request) xxx_FromOp(ctx context.Context, op *xxx_FTPClearStatistics2Operation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.Instance = op.Instance
}
func (o *FTPClearStatistics2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FTPClearStatistics2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FTPClearStatistics2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FTPClearStatistics2Response structure represents the R_FtpClearStatistics2 operation response
type FTPClearStatistics2Response struct {
	// Return: The R_FtpClearStatistics2 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *FTPClearStatistics2Response) xxx_ToOp(ctx context.Context, op *xxx_FTPClearStatistics2Operation) *xxx_FTPClearStatistics2Operation {
	if op == nil {
		op = &xxx_FTPClearStatistics2Operation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *FTPClearStatistics2Response) xxx_FromOp(ctx context.Context, op *xxx_FTPClearStatistics2Operation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *FTPClearStatistics2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FTPClearStatistics2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FTPClearStatistics2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IISEnumerateUsersOperation structure represents the R_IISEnumerateUsers operation
type xxx_IISEnumerateUsersOperation struct {
	Server    string       `idl:"name:pszServer;string;pointer:unique" json:"server"`
	ServiceID uint32       `idl:"name:dwServiceId" json:"service_id"`
	Instance  uint32       `idl:"name:dwInstance" json:"instance"`
	Info      *IISUserEnum `idl:"name:InfoStruct" json:"info"`
	Return    uint32       `idl:"name:Return" json:"return"`
}

func (o *xxx_IISEnumerateUsersOperation) OpNum() int { return 14 }

func (o *xxx_IISEnumerateUsersOperation) OpName() string { return "/inetinfo/v2/R_IISEnumerateUsers" }

func (o *xxx_IISEnumerateUsersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IISEnumerateUsersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Server != "" {
			_ptr_pszServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Server); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_pszServer); err != nil {
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
	// dwServiceId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ServiceID); err != nil {
			return err
		}
	}
	// dwInstance {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Instance); err != nil {
			return err
		}
	}
	// InfoStruct {in, out} (1:{alias=LPIIS_USER_ENUM_STRUCT}*(1))(2:{alias=IIS_USER_ENUM_STRUCT}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&IISUserEnum{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IISEnumerateUsersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Server); err != nil {
				return err
			}
			return nil
		})
		_s_pszServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
		if err := w.ReadPointer(&o.Server, _s_pszServer, _ptr_pszServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwServiceId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ServiceID); err != nil {
			return err
		}
	}
	// dwInstance {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Instance); err != nil {
			return err
		}
	}
	// InfoStruct {in, out} (1:{alias=LPIIS_USER_ENUM_STRUCT,pointer=ref}*(1))(2:{alias=IIS_USER_ENUM_STRUCT}(struct))
	{
		if o.Info == nil {
			o.Info = &IISUserEnum{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IISEnumerateUsersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IISEnumerateUsersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// InfoStruct {in, out} (1:{alias=LPIIS_USER_ENUM_STRUCT}*(1))(2:{alias=IIS_USER_ENUM_STRUCT}(struct))
	{
		if o.Info != nil {
			if err := o.Info.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&IISUserEnum{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_IISEnumerateUsersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// InfoStruct {in, out} (1:{alias=LPIIS_USER_ENUM_STRUCT,pointer=ref}*(1))(2:{alias=IIS_USER_ENUM_STRUCT}(struct))
	{
		if o.Info == nil {
			o.Info = &IISUserEnum{}
		}
		if err := o.Info.UnmarshalNDR(ctx, w); err != nil {
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

// IISEnumerateUsersRequest structure represents the R_IISEnumerateUsers operation request
type IISEnumerateUsersRequest struct {
	// pszServer: The custom binding handle for the target system, as specified in section
	// 2.1.1. The value MUST NOT be used by the server implementation.
	Server string `idl:"name:pszServer;string;pointer:unique" json:"server"`
	// dwServiceId: The identifier for the specified Internet protocol server, as specified
	// in section 2.2.2.
	ServiceID uint32 `idl:"name:dwServiceId" json:"service_id"`
	// dwInstance: The ID of the Internet protocol server instance whose users are being
	// enumerated.
	Instance uint32 `idl:"name:dwInstance" json:"instance"`
	// InfoStruct: The pointer to an IIS_USER_ENUM_STRUCT that contains the list of active
	// users for this server.
	Info *IISUserEnum `idl:"name:InfoStruct" json:"info"`
}

func (o *IISEnumerateUsersRequest) xxx_ToOp(ctx context.Context, op *xxx_IISEnumerateUsersOperation) *xxx_IISEnumerateUsersOperation {
	if op == nil {
		op = &xxx_IISEnumerateUsersOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.ServiceID = o.ServiceID
	op.Instance = o.Instance
	op.Info = o.Info
	return op
}

func (o *IISEnumerateUsersRequest) xxx_FromOp(ctx context.Context, op *xxx_IISEnumerateUsersOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.ServiceID = op.ServiceID
	o.Instance = op.Instance
	o.Info = op.Info
}
func (o *IISEnumerateUsersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *IISEnumerateUsersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IISEnumerateUsersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IISEnumerateUsersResponse structure represents the R_IISEnumerateUsers operation response
type IISEnumerateUsersResponse struct {
	// InfoStruct: The pointer to an IIS_USER_ENUM_STRUCT that contains the list of active
	// users for this server.
	Info *IISUserEnum `idl:"name:InfoStruct" json:"info"`
	// Return: The R_IISEnumerateUsers return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *IISEnumerateUsersResponse) xxx_ToOp(ctx context.Context, op *xxx_IISEnumerateUsersOperation) *xxx_IISEnumerateUsersOperation {
	if op == nil {
		op = &xxx_IISEnumerateUsersOperation{}
	}
	if o == nil {
		return op
	}
	op.Info = o.Info
	op.Return = o.Return
	return op
}

func (o *IISEnumerateUsersResponse) xxx_FromOp(ctx context.Context, op *xxx_IISEnumerateUsersOperation) {
	if o == nil {
		return
	}
	o.Info = op.Info
	o.Return = op.Return
}
func (o *IISEnumerateUsersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *IISEnumerateUsersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IISEnumerateUsersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IISDisconnectUserOperation structure represents the R_IISDisconnectUser operation
type xxx_IISDisconnectUserOperation struct {
	Server    string `idl:"name:pszServer;string;pointer:unique" json:"server"`
	ServiceID uint32 `idl:"name:dwServiceId" json:"service_id"`
	Instance  uint32 `idl:"name:dwInstance" json:"instance"`
	IDUser    uint32 `idl:"name:dwIdUser" json:"id_user"`
	Return    uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_IISDisconnectUserOperation) OpNum() int { return 15 }

func (o *xxx_IISDisconnectUserOperation) OpName() string { return "/inetinfo/v2/R_IISDisconnectUser" }

func (o *xxx_IISDisconnectUserOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IISDisconnectUserOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Server != "" {
			_ptr_pszServer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Server); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Server, _ptr_pszServer); err != nil {
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
	// dwServiceId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ServiceID); err != nil {
			return err
		}
	}
	// dwInstance {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Instance); err != nil {
			return err
		}
	}
	// dwIdUser {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.IDUser); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IISDisconnectUserOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pszServer {in} (1:{handle, string, pointer=unique, alias=INET_INFO_IMPERSONATE_HANDLE, names=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pszServer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Server); err != nil {
				return err
			}
			return nil
		})
		_s_pszServer := func(ptr interface{}) { o.Server = *ptr.(*string) }
		if err := w.ReadPointer(&o.Server, _s_pszServer, _ptr_pszServer); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwServiceId {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ServiceID); err != nil {
			return err
		}
	}
	// dwInstance {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Instance); err != nil {
			return err
		}
	}
	// dwIdUser {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.IDUser); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IISDisconnectUserOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IISDisconnectUserOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_IISDisconnectUserOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// IISDisconnectUserRequest structure represents the R_IISDisconnectUser operation request
type IISDisconnectUserRequest struct {
	// pszServer: The custom binding handle for the target system, as specified in section
	// 2.1.1. The value MUST NOT be used by the server implementation.
	Server string `idl:"name:pszServer;string;pointer:unique" json:"server"`
	// dwServiceId: The identifier for the specified Internet protocol server, as specified
	// in section 2.2.2.
	ServiceID uint32 `idl:"name:dwServiceId" json:"service_id"`
	// dwInstance: The ID of the Internet protocol server instance whose user is being disconnected.
	Instance uint32 `idl:"name:dwInstance" json:"instance"`
	// dwIdUser: The identifier of the user to disconnect, as found in the idUser field
	// of the IIS_USER_INFO_1 structure returned by the R_IISEnumerateUsers method. A value
	// of 0 for this parameter indicates that the server implementation MUST attempt to
	// disconnect all users from this Internet protocol server instance.
	IDUser uint32 `idl:"name:dwIdUser" json:"id_user"`
}

func (o *IISDisconnectUserRequest) xxx_ToOp(ctx context.Context, op *xxx_IISDisconnectUserOperation) *xxx_IISDisconnectUserOperation {
	if op == nil {
		op = &xxx_IISDisconnectUserOperation{}
	}
	if o == nil {
		return op
	}
	op.Server = o.Server
	op.ServiceID = o.ServiceID
	op.Instance = o.Instance
	op.IDUser = o.IDUser
	return op
}

func (o *IISDisconnectUserRequest) xxx_FromOp(ctx context.Context, op *xxx_IISDisconnectUserOperation) {
	if o == nil {
		return
	}
	o.Server = op.Server
	o.ServiceID = op.ServiceID
	o.Instance = op.Instance
	o.IDUser = op.IDUser
}
func (o *IISDisconnectUserRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *IISDisconnectUserRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IISDisconnectUserOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IISDisconnectUserResponse structure represents the R_IISDisconnectUser operation response
type IISDisconnectUserResponse struct {
	// Return: The R_IISDisconnectUser return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *IISDisconnectUserResponse) xxx_ToOp(ctx context.Context, op *xxx_IISDisconnectUserOperation) *xxx_IISDisconnectUserOperation {
	if op == nil {
		op = &xxx_IISDisconnectUserOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *IISDisconnectUserResponse) xxx_FromOp(ctx context.Context, op *xxx_IISDisconnectUserOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *IISDisconnectUserResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *IISDisconnectUserResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IISDisconnectUserOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
