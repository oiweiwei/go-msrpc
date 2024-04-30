package inetinfo

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

// inetinfo server interface.
type InetinfoServer interface {

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
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)

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
	GetAdminInformation(context.Context, *GetAdminInformationRequest) (*GetAdminInformationResponse, error)

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
	GetSites(context.Context, *GetSitesRequest) (*GetSitesResponse, error)

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
	SetAdminInformation(context.Context, *SetAdminInformationRequest) (*SetAdminInformationResponse, error)

	// The R_InetInfoGetGlobalAdminInformation method is called by the client. In response,
	// the server retrieves configuration data shared by all Internet protocol servers.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2 or [MS-ERREF]
	// section 2.3.1.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol, as specified in [MS-RPCE].
	GetGlobalAdminInformation(context.Context, *GetGlobalAdminInformationRequest) (*GetGlobalAdminInformationResponse, error)

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
	SetGlobalAdminInformation(context.Context, *SetGlobalAdminInformationRequest) (*SetGlobalAdminInformationResponse, error)

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
	QueryStatistics(context.Context, *QueryStatisticsRequest) (*QueryStatisticsResponse, error)

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
	ClearStatistics(context.Context, *ClearStatisticsRequest) (*ClearStatisticsResponse, error)

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
	FlushMemoryCache(context.Context, *FlushMemoryCacheRequest) (*FlushMemoryCacheResponse, error)

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
	GetServerCapabilities(context.Context, *GetServerCapabilitiesRequest) (*GetServerCapabilitiesResponse, error)

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
	W3QueryStatistics2(context.Context, *W3QueryStatistics2Request) (*W3QueryStatistics2Response, error)

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
	W3ClearStatistics2(context.Context, *W3ClearStatistics2Request) (*W3ClearStatistics2Response, error)

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
	FTPQueryStatistics2(context.Context, *FTPQueryStatistics2Request) (*FTPQueryStatistics2Response, error)

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
	FTPClearStatistics2(context.Context, *FTPClearStatistics2Request) (*FTPClearStatistics2Response, error)

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
	IISEnumerateUsers(context.Context, *IISEnumerateUsersRequest) (*IISEnumerateUsersResponse, error)

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
	IISDisconnectUser(context.Context, *IISDisconnectUserRequest) (*IISDisconnectUserResponse, error)

	// Opnum16NotUsedOnWire operation.
	// Opnum16NotUsedOnWire

	// Opnum17NotUsedOnWire operation.
	// Opnum17NotUsedOnWire
}

func RegisterInetinfoServer(conn dcerpc.Conn, o InetinfoServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewInetinfoServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(InetinfoSyntaxV2_0))...)
}

func NewInetinfoServerHandle(o InetinfoServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return InetinfoServerHandle(ctx, o, opNum, r)
	}
}

func InetinfoServerHandle(ctx context.Context, o InetinfoServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // R_InetInfoGetVersion
		in := &GetVersionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetVersion(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // R_InetInfoGetAdminInformation
		in := &GetAdminInformationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAdminInformation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // R_InetInfoGetSites
		in := &GetSitesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSites(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // R_InetInfoSetAdminInformation
		in := &SetAdminInformationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetAdminInformation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // R_InetInfoGetGlobalAdminInformation
		in := &GetGlobalAdminInformationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetGlobalAdminInformation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // R_InetInfoSetGlobalAdminInformation
		in := &SetGlobalAdminInformationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetGlobalAdminInformation(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // R_InetInfoQueryStatistics
		in := &QueryStatisticsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryStatistics(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // R_InetInfoClearStatistics
		in := &ClearStatisticsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ClearStatistics(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // R_InetInfoFlushMemoryCache
		in := &FlushMemoryCacheRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FlushMemoryCache(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // R_InetInfoGetServerCapabilities
		in := &GetServerCapabilitiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetServerCapabilities(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // R_W3QueryStatistics2
		in := &W3QueryStatistics2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.W3QueryStatistics2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // R_W3ClearStatistics2
		in := &W3ClearStatistics2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.W3ClearStatistics2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // R_FtpQueryStatistics2
		in := &FTPQueryStatistics2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FTPQueryStatistics2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // R_FtpClearStatistics2
		in := &FTPClearStatistics2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FTPClearStatistics2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // R_IISEnumerateUsers
		in := &IISEnumerateUsersRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.IISEnumerateUsers(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // R_IISDisconnectUser
		in := &IISDisconnectUserRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.IISDisconnectUser(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // Opnum16NotUsedOnWire
		// Opnum16NotUsedOnWire
		return nil, nil
	case 17: // Opnum17NotUsedOnWire
		// Opnum17NotUsedOnWire
		return nil, nil
	}
	return nil, nil
}
