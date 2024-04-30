package srvsvc

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

// srvsvc server interface.
type SrvsvcServer interface {

	// Opnum0NotUsedOnWire operation.
	// Opnum0NotUsedOnWire

	// Opnum1NotUsedOnWire operation.
	// Opnum1NotUsedOnWire

	// Opnum2NotUsedOnWire operation.
	// Opnum2NotUsedOnWire

	// Opnum3NotUsedOnWire operation.
	// Opnum3NotUsedOnWire

	// Opnum4NotUsedOnWire operation.
	// Opnum4NotUsedOnWire

	// Opnum5NotUsedOnWire operation.
	// Opnum5NotUsedOnWire

	// Opnum6NotUsedOnWire operation.
	// Opnum6NotUsedOnWire

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// The NetrConnectionEnum method lists all the treeconnects made to a shared resource
	// on the server or all treeconnects established from a particular computer.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.
	ConnectionEnum(context.Context, *ConnectionEnumRequest) (*ConnectionEnumResponse, error)

	// The NetrFileEnum method MUST return information about some or all open files on a
	// server, depending on the parameters specified, or return an error code.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The client request succeeded.                                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | The system call level is not correct.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA         | The client request succeeded. More entries are available. Not all entries could  |
	//	|                                    | be returned in the buffer size that is specified by PreferedMaximumLength.       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000084B NERR_BufTooSmall        | The client request succeeded. More entries are available. The buffer size that   |
	//	|                                    | is specified by PreferedMaximumLength was too small to fit even a single entry.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	FileEnum(context.Context, *FileEnumRequest) (*FileEnumResponse, error)

	// The NetrFileGetInfo method MUST retrieve information about a particular open server
	// resource or return an error code.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------+
	//	|               RETURN               |                                                          |
	//	|             VALUE/CODE             |                       DESCRIPTION                        |
	//	|                                    |                                                          |
	//	+------------------------------------+----------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The client request succeeded.                            |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                        |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The system cannot find the file specified.               |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | The system call level is not correct.                    |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command. |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x0000084B NERR_BufTooSmall        | The supplied buffer is too small.                        |
	//	+------------------------------------+----------------------------------------------------------+
	FileGetInfo(context.Context, *FileGetInfoRequest) (*FileGetInfoResponse, error)

	// The server receives the NetrFileClose method in an RPC_REQUEST packet. In response,
	// the server MUST force an open resource instance (for example, file, device, or named
	// pipe) on the server to close. This message can be used when an error prevents closure
	// by any other means.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+--------------------------------+-----------------------------------------------------------------+
	//	|             RETURN             |                                                                 |
	//	|           VALUE/CODE           |                           DESCRIPTION                           |
	//	|                                |                                                                 |
	//	+--------------------------------+-----------------------------------------------------------------+
	//	+--------------------------------+-----------------------------------------------------------------+
	//	| 0x00000000 NERR_Success        | The client request succeeded.                                   |
	//	+--------------------------------+-----------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | Access is denied.                                               |
	//	+--------------------------------+-----------------------------------------------------------------+
	//	| 0x0000090A NERR_FileIdNotFound | There is no open file with the specified identification number. |
	//	+--------------------------------+-----------------------------------------------------------------+
	//
	// This message can be used when an error prevents closure by any other means.
	//
	// The FileId parameter specifies the file identifier of the Open in FileList to close.
	// The value of the FileId parameter MUST correspond to a FileId that is returned in
	// a previous NetrFileEnum message response by the server. The server MUST look up Open
	// in the FileList where FileId matches Open.GlobalFileId. If no match is found, the
	// server MUST return NERR_FileIdNotFound. If a match is found, the server MUST close
	// the Open by invoking an underlying server event as specified in [MS-CIFS] section
	// 3.3.4.13 or [MS-SMB2] section 3.3.4.17, providing FileId as the input parameter.
	//
	// If either CIFS or SMB2 servers return STATUS_SUCCESS, the server MUST return NERR_Success.
	// Otherwise, the server MUST fail the call with a NERR_FileIdNotFound error code.
	//
	// The server SHOULD<49> enforce security measures to verify that the caller has the
	// required permissions to execute this routine. If the caller does not have the required
	// credentials, the server SHOULD<50> fail the call.
	FileClose(context.Context, *FileCloseRequest) (*FileCloseResponse, error)

	// The NetrSessionEnum method MUST return information about sessions that are established
	// on a server or return an error code.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The client request succeeded.                                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | The system call level is not correct.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The parameter is incorrect.                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA         | The client request succeeded. More entries are available. Not all entries could  |
	//	|                                    | be returned in the buffer size that is specified by PreferedMaximumLength.       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000908 NERR_ClientNameNotFound | A session does not exist with the computer name.                                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000092F NERR_InvalidComputer    | The computer name is not valid.                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000008AD NERR_UserNotFound       | The user name could not be found.                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	SessionEnum(context.Context, *SessionEnumRequest) (*SessionEnumResponse, error)

	// The NetrSessionDel method MUST end one or more network sessions between a server
	// and a client.
	//
	// Return Values: This method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. This method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------+
	//	|               RETURN               |                                                          |
	//	|             VALUE/CODE             |                       DESCRIPTION                        |
	//	|                                    |                                                          |
	//	+------------------------------------+----------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The client request succeeded.                            |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                        |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The parameter is incorrect.                              |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command. |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000908 NERR_ClientNameNotFound | A session does not exist with the computer name.         |
	//	+------------------------------------+----------------------------------------------------------+
	SessionDelete(context.Context, *SessionDeleteRequest) (*SessionDeleteResponse, error)

	// The NetrShareAdd method shares a server resource.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The client request succeeded.                                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | The system call level is not correct.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000007B ERROR_INVALID_NAME      | The file name, directory name, or volume label syntax is incorrect.              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The client request failed because the specified parameter is invalid. For        |
	//	|                                    | details, see the description that follows for the ParmErr parameter.             |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000846 NERR_DuplicateShare     | The share name is already in use on this server.                                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000844 NERR_UnknownDevDir      | The device or directory does not exist.                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	ShareAdd(context.Context, *ShareAddRequest) (*ShareAddResponse, error)

	// The NetrShareEnum method retrieves information about each shared resource on a server.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success        | The client request succeeded.                                                    |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA     | The client request succeeded. More entries are available. Not all entries could  |
	//	|                                | be returned in the buffer size that is specified by PreferedMaximumLength.       |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL | The system call level is not correct.                                            |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// If ServerName does not match any Transport.ServerName in TransportList with the SVTI2_SCOPED_NAME
	// bit set in Transport.Flags, the server MUST reset ServerName as "*".
	ShareEnum(context.Context, *ShareEnumRequest) (*ShareEnumResponse, error)

	// The NetrShareGetInfo method retrieves information about a particular shared resource
	// on the server from the ShareList.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	|               RETURN               |                                                                       |
	//	|             VALUE/CODE             |                              DESCRIPTION                              |
	//	|                                    |                                                                       |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The client request succeeded.                                         |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                                     |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | The system call level is not correct.                                 |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The client request failed because the specified parameter is invalid. |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command.              |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x0000084B NERR_BufTooSmall        | The supplied buffer is too small.                                     |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000906 NERR_NetNameNotFound    | The share name does not exist.                                        |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//
	// If ServerName does not match any Transport.ServerName in TransportList with the SVTI2_SCOPED_NAME
	// bit set in Transport.Flags, the server MUST reset ServerName as "*".
	ShareGetInfo(context.Context, *ShareGetInfoRequest) (*ShareGetInfoResponse, error)

	// The NetrShareSetInfo method sets the parameters of a shared resource in a ShareList.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                                  |
	//	|               VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success                 | The client request succeeded.                                                    |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED          | Access is denied.                                                                |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER      | The client request failed because the specified parameter is invalid. For        |
	//	|                                         | details, see the description that follows for the ParmErr parameter.             |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY      | Not enough storage is available to process this command.                         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000906 NERR_NetNameNotFound         | The share name does not exist.                                                   |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED          | The server does not support branch cache. <62>                                   |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000424 ERROR_SERVICE_DOES_NOT_EXIST | The branch cache component does not exist as an installed service. <63>          |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL          | The system call level is not correct.                                            |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	ShareSetInfo(context.Context, *ShareSetInfoRequest) (*ShareSetInfoResponse, error)

	// The NetrShareDel method deletes a share name from the ShareList, which disconnects
	// all connections to the shared resource. If the share is sticky, all information about
	// the share is also deleted from permanent storage.<67>
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	|               RETURN               |                                                                       |
	//	|             VALUE/CODE             |                              DESCRIPTION                              |
	//	|                                    |                                                                       |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The client request succeeded.                                         |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                                     |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The client request failed because the specified parameter is invalid. |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command.              |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000906 NERR_NetNameNotFound    | The share name does not exist.                                        |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//
	// If ServerName does not match any Transport.ServerName in TransportList with the SVTI2_SCOPED_NAME
	// bit set in Transport.Flags, the server MUST reset ServerName as "*".
	ShareDelete(context.Context, *ShareDeleteRequest) (*ShareDeleteResponse, error)

	// The NetrShareDelSticky method marks the share as nonpersistent by clearing the IsPersistent
	// member of a Share in the ShareList.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.
	//
	// The primary use of this method is to delete a sticky share whose root directory has
	// been deleted (thus preventing actual re-creation of the share) but whose entry still
	// exists in permanent storage.<72> This method can also be used to remove the persistence
	// of a share without deleting the current incarnation of the share.
	//
	// If ServerName does not match any Transport.ServerName in TransportList with the SVTI2_SCOPED_NAME
	// bit set in Transport.Flags, the server MUST reset ServerName as "*".
	ShareDeleteSticky(context.Context, *ShareDeleteStickyRequest) (*ShareDeleteStickyResponse, error)

	// The NetrShareCheck method checks whether a server is sharing a device.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------+
	//	|               RETURN               |                                                          |
	//	|             VALUE/CODE             |                       DESCRIPTION                        |
	//	|                                    |                                                          |
	//	+------------------------------------+----------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The client request succeeded.                            |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command. |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000907 NERR_DeviceNotShared    | The device is not shared.                                |
	//	+------------------------------------+----------------------------------------------------------+
	ShareCheck(context.Context, *ShareCheckRequest) (*ShareCheckResponse, error)

	// The NetrServerGetInfo method retrieves current configuration information for CIFS
	// and SMB Version 1.0 servers.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	|               RETURN               |                                                                       |
	//	|             VALUE/CODE             |                              DESCRIPTION                              |
	//	|                                    |                                                                       |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The client request succeeded.                                         |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                                     |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | The system call level is not correct.                                 |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The client request failed because the specified parameter is invalid. |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command.              |
	//	+------------------------------------+-----------------------------------------------------------------------+
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error)

	// The NetrServerSetInfo method sets server operating parameters for CIFS and SMB Version
	// 1.0 file servers; it can set them individually or collectively. The information is
	// stored in a way that allows it to remain in effect after the system is reinitialized.<81>
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The client request succeeded.                                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | The system call level is not correct.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The client request failed because the specified parameter is invalid. For        |
	//	|                                    | details see the description that follows for the ParmErr parameter.              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	SetInfo(context.Context, *SetInfoRequest) (*SetInfoResponse, error)

	// The NetrServerDiskEnum method retrieves a list of disk drives on a server. The method
	// returns an array of three-character strings (a drive letter, a colon, and a terminating
	// null character).
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+------------------------------------+--------------------------------------------------------------------+
	//	|               RETURN               |                                                                    |
	//	|             VALUE/CODE             |                            DESCRIPTION                             |
	//	|                                    |                                                                    |
	//	+------------------------------------+--------------------------------------------------------------------+
	//	+------------------------------------+--------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The client request succeeded.                                      |
	//	+------------------------------------+--------------------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | The system call level is not correct.                              |
	//	+------------------------------------+--------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command.           |
	//	+------------------------------------+--------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The caller does not have the permissions to perform the operation. |
	//	+------------------------------------+--------------------------------------------------------------------+
	DiskEnum(context.Context, *DiskEnumRequest) (*DiskEnumResponse, error)

	// The NetrServerStatisticsGet method retrieves the operating statistics for a service.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.
	StatisticsGet(context.Context, *StatisticsGetRequest) (*StatisticsGetResponse, error)

	// The NetrServerTransportAdd method binds the server to the transport protocol.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------+
	//	|               RETURN               |                                                          |
	//	|             VALUE/CODE             |                       DESCRIPTION                        |
	//	|                                    |                                                          |
	//	+------------------------------------+----------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The client request succeeded.                            |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                        |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000034 ERROR_DUP_NAME          | A duplicate name exists on the network.                  |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | The system call level is not correct.                    |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The parameter is incorrect.                              |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command. |
	//	+------------------------------------+----------------------------------------------------------+
	//
	// The NetrServerTransportAdd message MUST be processed in the same way as the NetrServerTransportAddEx
	// message, except that it MUST allow only level 0 (that is, SERVER_TRANSPORT_INFO_0).
	// The NetrServerTransportAddEx message is specified in section 3.1.4.23.
	//
	// The server MAY<91> enforce security measures to verify that the caller has the required
	// permissions to execute this call. If the server enforces these security measures
	// and the caller does not have the required credentials, the server SHOULD<92> fail
	// the call.
	TransportAdd(context.Context, *TransportAddRequest) (*TransportAddResponse, error)

	// The NetrServerTransportEnum method enumerates the information about transport protocols
	// that the server manages in TransportList.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The client request succeeded.                                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | The system call level is not correct.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA         | The client request succeeded. More entries are available. Not all entries could  |
	//	|                                    | be returned in the buffer size that is specified by PreferedMaximumLength.       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000084B NERR_BufTooSmall        | The client request succeeded. More entries are available. The buffer size that   |
	//	|                                    | is specified by PreferedMaximumLength was too small to fit even a single entry.  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	TransportEnum(context.Context, *TransportEnumRequest) (*TransportEnumResponse, error)

	// The NetrServerTransportDel method unbinds (or disconnects) the transport protocol
	// from the server. If this method succeeds, the server can no longer communicate with
	// clients by using the specified transport protocol (such as TCP or XNS).
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------+
	//	|               RETURN               |                                                          |
	//	|             VALUE/CODE             |                       DESCRIPTION                        |
	//	|                                    |                                                          |
	//	+------------------------------------+----------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The client request succeeded.                            |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                        |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The parameter is incorrect.                              |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command. |
	//	+------------------------------------+----------------------------------------------------------+
	//
	// The NetrServerTransportDel message MUST be processed in the same way as the NetrServerTransportDelEx
	// message, except that it MUST allow only level 0 (that is, SERVER_TRANSPORT_INFO_0).
	// The processing for this message is specified in section 3.1.4.26.
	//
	// The server MAY<97> enforce security measures to verify that the caller has the required
	// permissions to execute this call. If the server enforces these security measures
	// and the caller does not have the required credentials, the server SHOULD<98> fail
	// the call.
	TransportDelete(context.Context, *TransportDeleteRequest) (*TransportDeleteResponse, error)

	// The NetrRemoteTOD method returns the time of day information on a server.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.
	RemoteToD(context.Context, *RemoteToDRequest) (*RemoteToDResponse, error)

	// Opnum29NotUsedOnWire operation.
	// Opnum29NotUsedOnWire

	// The NetprPathType method checks a path name to determine its type.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.
	PathType(context.Context, *PathTypeRequest) (*PathTypeResponse, error)

	// The NetprPathCanonicalize method converts a path name to the canonical format.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.
	//
	// If the Flags parameter is not equal to zero, the server SHOULD fail the call with
	// an implementation-specific error code.<110>
	PathCanonicalize(context.Context, *PathCanonicalizeRequest) (*PathCanonicalizeResponse, error)

	// The NetprPathCompare method performs comparison of two paths.
	//
	// Return Values: Upon successful processing, the server MUST return 0 if both paths
	// are the same, –1 if the first is less than the second, and 1 otherwise. If the
	// method fails, it can return any specific error code value as specified in [MS-ERREF]
	// section 2.2.
	PathCompare(context.Context, *PathCompareRequest) (*PathCompareResponse, error)

	// The NetprNameValidate method performs checks to ensure that the specified name is
	// a valid name for the specified type.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.
	//
	// If the Flags parameter is not equal to zero, the server SHOULD fail the call with
	// an implementation-specific error code.<119>
	NameValidate(context.Context, *NameValidateRequest) (*NameValidateResponse, error)

	// The NetprNameCanonicalize method converts a name to the canonical format for the
	// specified type.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.
	NameCanonicalize(context.Context, *NameCanonicalizeRequest) (*NameCanonicalizeResponse, error)

	// The NetprNameCompare method does comparison of two names of a specific name type.
	//
	// Return Values: MUST return 0 if both paths are the same. Other values indicate that
	// either the paths are different or an error occurred when the client request was processed.
	NameCompare(context.Context, *NameCompareRequest) (*NameCompareResponse, error)

	// The NetrShareEnumSticky method retrieves information about each sticky shared resource
	// whose IsPersistent setting is set in a ShareList.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|           RETURN            |                                                                                  |
	//	|         VALUE/CODE          |                                   DESCRIPTION                                    |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success     | The client request succeeded.                                                    |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA  | The client request succeeded. More entries are available. Not all entries could  |
	//	|                             | be returned in the buffer size that is specified by PreferedMaximumLength.       |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000084B NERR_BufTooSmall | The client request succeeded. More entries are available. The buffer size that   |
	//	|                             | is specified by PreferedMaximumLength was too small to fit even a single entry.  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	ShareEnumSticky(context.Context, *ShareEnumStickyRequest) (*ShareEnumStickyResponse, error)

	// The NetrShareDelStart method performs the initial phase of a two-phase share delete.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.
	ShareDeleteStart(context.Context, *ShareDeleteStartRequest) (*ShareDeleteStartResponse, error)

	// The NetrShareDelCommit method performs the final phase of a two-phase share delete.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success.
	// Otherwise, the method returns a nonzero error code unless the share being deleted
	// is IPC$. If the share being deleted is IPC$, the return value is not meaningful.
	// The method can take any specific error code value, as specified in [MS-ERREF] section
	// 2.2.
	//
	// The NetrShareDelCommit message is the continuation of the NetrShareDelStart message
	// and MUST cause the share to be actually deleted, which disconnects all connections
	// to the share, or MUST return an error code.
	//
	// This method can be used to delete the IPC$ share as well as other shares. When the
	// share is not IPC$, only a return value of 0 indicates success.
	//
	// This two-phase deletion MUST be used to delete IPC$, which is the share that is used
	// for named pipes. Deleting IPC$ results in the closing of the pipe on which the RPC
	// is being executed. Thus, the client never receives the response to the RPC. The two-phase
	// delete offers a positive response in phase 1 and then an expected error in phase
	// 2.
	//
	// ContextHandle MUST reference the share to be deleted in the NetrShareDelStart method.
	// If a share is not found, the server MUST fail the call with an ERROR_INVALID_PARAMETER
	// error code.
	//
	// If a share is found, but the IsMarkedForDeletion member of the Share is not set,
	// the server MUST fail the call with an ERROR_INVALID_PARAMETER error code.
	//
	// Otherwise, the server MUST delete the share by invoking the underlying server event,
	// as specified in [MS-CIFS] section 3.3.4.11and [MS-SMB2] section 3.3.4.15, providing
	// tuple <ServerName, NetName> as input parameters.
	//
	// The server does not enforce any security measures when processing this call.
	ShareDeleteCommit(context.Context, *ShareDeleteCommitRequest) (*ShareDeleteCommitResponse, error)

	// The NetrpGetFileSecurity method returns to the caller a copy of the security descriptor
	// that protects a file or directory.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.
	GetFileSecurity(context.Context, *GetFileSecurityRequest) (*GetFileSecurityResponse, error)

	// The NetrpSetFileSecurity method sets the security of a file or directory.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.
	SetFileSecurity(context.Context, *SetFileSecurityRequest) (*SetFileSecurityResponse, error)

	// The NetrServerTransportAddEx method binds the specified server to the transport protocol.
	// This extended method allows the caller to specify information levels 1, 2, and 3
	// beyond what the NetrServerTransportAdd method allows.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------+
	//	|               RETURN               |                                                          |
	//	|             VALUE/CODE             |                       DESCRIPTION                        |
	//	|                                    |                                                          |
	//	+------------------------------------+----------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The client request succeeded.                            |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                        |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000034 ERROR_DUP_NAME          | A duplicate name exists on the network.                  |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | The system call level is not correct.                    |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The parameter is incorrect.                              |
	//	+------------------------------------+----------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command. |
	//	+------------------------------------+----------------------------------------------------------+
	//
	// The server SHOULD<93> enforce security measures to verify that the caller has the
	// required permissions to execute this call. If the caller does not have the required
	// credentials, the server SHOULD<94> fail the call.
	//
	// The Level parameter determines the type of structure that the client has used to
	// specify information about the new transport. The value MUST be 0, 1, 2, or 3. If
	// the Level parameter is not equal to one of the valid values, the server MUST fail
	// the call with an ERROR_INVALID_LEVEL error code.
	//
	// If the Level parameter is 0, the Buffer parameter points to a SERVER_TRANSPORT_INFO_0
	// structure.
	//
	// If the Level parameter is 1, the Buffer parameter points to a SERVER_TRANSPORT_INFO_1
	// structure.
	//
	// If the Level parameter is 2, the Buffer parameter points to a SERVER_TRANSPORT_INFO_2
	// structure.
	//
	// If the Level parameter is 3, the Buffer parameter points to a SERVER_TRANSPORT_INFO_3
	// structure.
	TransportAddEx(context.Context, *TransportAddExRequest) (*TransportAddExResponse, error)

	// Opnum42NotUsedOnWire operation.
	// Opnum42NotUsedOnWire

	// The NetrDfsGetVersion method checks whether the server is a DFS server and if so,
	// returns the DFS version. An implementation MAY<127> choose to support this method.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)

	// The NetrDfsCreateLocalPartition method marks a share as being a DFS share. In addition,
	// if the RelationInfo parameter is non-NULL, it creates DFS links, as specified in
	// [MS-DFSC], for each of the entries in the RelationInfo parameter. An implementation
	// MAY<132> choose to support this method.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.
	CreateLocalPartition(context.Context, *CreateLocalPartitionRequest) (*CreateLocalPartitionResponse, error)

	// The NetrDfsDeleteLocalPartition method deletes a DFS share (Prefix) on the server.
	// An implementation MAY<138> choose to support this method.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.
	DeleteLocalPartition(context.Context, *DeleteLocalPartitionRequest) (*DeleteLocalPartitionResponse, error)

	// The NetrDfsSetLocalVolumeState method sets a local DFS share online or offline. An
	// implementation MAY<142> choose to support this method.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.
	SetLocalVolumeState(context.Context, *SetLocalVolumeStateRequest) (*SetLocalVolumeStateResponse, error)

	// Opnum47NotUsedOnWire operation.
	// Opnum47NotUsedOnWire

	// The NetrDfsCreateExitPoint method creates a DFS link on the server. An implementation
	// MAY<146> choose to support this method.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.
	CreateExitPoint(context.Context, *CreateExitPointRequest) (*CreateExitPointResponse, error)

	// The NetrDfsDeleteExitPoint method deletes a DFS link on the server. An implementation
	// MAY<155> choose to support this method.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.
	DeleteExitPoint(context.Context, *DeleteExitPointRequest) (*DeleteExitPointResponse, error)

	// The NetrDfsModifyPrefix method changes the path that corresponds to a DFS link on
	// the server. An implementation MAY<151> choose to support this method.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.
	ModifyPrefix(context.Context, *ModifyPrefixRequest) (*ModifyPrefixResponse, error)

	// The NetrDfsFixLocalVolume method provides knowledge of a new DFS share on the server.
	// An implementation MAY<159> choose to support this method.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.
	FixLocalVolume(context.Context, *FixLocalVolumeRequest) (*FixLocalVolumeResponse, error)

	// The NetrDfsManagerReportSiteInfo method obtains a list of names that SHOULD<165>
	// correspond to the Active Directory sites covered by the specified server. An implementation
	// MAY<166> choose to support this method.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.
	//
	// The ppSiteInfo parameter is a pointer to a LPDFS_SITELIST_INFO member, which in turn
	// points to the location of a DFS_SITELIST_INFO structure in which the information
	// is returned. That structure has a cSites member that the server SHOULD set to the
	// number of sites returned. The information about the sites themselves MUST be returned
	// in the Site member, which is an array of DFS_SITENAME_INFO structures. The sites
	// the server returns are implementation-specific.<167>
	//
	// The server MAY<168> enforce security measures to verify that the caller has the required
	// permissions to execute this call. If the server enforces these security measures
	// and the caller does not have the required credentials, the server SHOULD<169> fail
	// the call.
	ManagerReportSiteInfo(context.Context, *ManagerReportSiteInfoRequest) (*ManagerReportSiteInfoResponse, error)

	// The server receives the NetrServerTransportDelEx method in an RPC_REQUEST packet.
	// In response, the server unbinds (or disconnects) the transport protocol from the
	// server. If this method succeeds, the server can no longer communicate with clients
	// by using the specified transport protocol (such as TCP or XNS). This extended method
	// allows level 1 beyond what the NetrServerTransportDel method allows.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2.<99>
	//
	// The Level parameter determines the type of structure the client has used to specify
	// information about the new transport. Valid values are 0 and 1. If the Level parameter
	// is not equal to one of the valid values, the server MUST fail the call with an ERROR_INVALID_LEVEL
	// error code.
	//
	// If the Level parameter is 0, the Buffer parameter points to a SERVER_TRANSPORT_INFO_0
	// structure. If the Level parameter is 1, the Buffer parameter points to a SERVER_TRANSPORT_INFO_1
	// structure.
	TransportDeleteEx(context.Context, *TransportDeleteExRequest) (*TransportDeleteExResponse, error)

	// The NetrServerAliasAdd method attaches an alias name to an existing server name and
	// inserts Alias objects into AliasList, through which the shared resource can be accessed
	// either with server name or alias name. An alias is used to identify which resources
	// are visible to an SMB client based on the server name presented in each tree connect
	// request.
	//
	// );
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	|               RETURN               |                                                                       |
	//	|             VALUE/CODE             |                              DESCRIPTION                              |
	//	|                                    |                                                                       |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The client request succeeded.                                         |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                                     |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The client request failed because the specified parameter is invalid. |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command.              |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000846 NERR_DuplicateShare     | The alias already exists.                                             |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | The system call level is not correct.                                 |
	//	+------------------------------------+-----------------------------------------------------------------------+
	AliasAdd(context.Context, *AliasAddRequest) (*AliasAddResponse, error)

	// The NetrServerAliasEnum method retrieves alias information for a server based on
	// specified alias name or server name.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The client request succeeded.                                                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The client request failed because the specified parameter is invalid.            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command.                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000084B NERR_BufTooSmall        | The allocated buffer is too small to hold single entry.                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | The system call level is not correct.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA         | The client request succeeded. More entries are available. Not all entries could  |
	//	|                                    | be returned in the buffer size that is specified by PreferedMaximumLength.       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	AliasEnum(context.Context, *AliasEnumRequest) (*AliasEnumResponse, error)

	// The NetrServerAliasDel method deletes an alias name from a server alias list based
	// on specified alias name.
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	|               RETURN               |                                                                       |
	//	|             VALUE/CODE             |                              DESCRIPTION                              |
	//	|                                    |                                                                       |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The client request succeeded.                                         |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                                     |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The client request failed because the specified parameter is invalid. |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command.              |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000906 NERR_NetNameNotFound    | The alias does not exist.                                             |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | The system call level is not correct.                                 |
	//	+------------------------------------+-----------------------------------------------------------------------+
	AliasDelete(context.Context, *AliasDeleteRequest) (*AliasDeleteResponse, error)

	// The NetrShareDelEx method deletes a share from the ShareList, which disconnects all
	// connections to the shared resource. If the share is sticky, all information about
	// the share is also deleted from permanent storage.<176>
	//
	// Return Values: The method returns 0x00000000 (NERR_Success) to indicate success;
	// otherwise, it returns a nonzero error code. The method can take any specific error
	// code value, as specified in [MS-ERREF] section 2.2. The most common error codes are
	// listed in the following table.
	//
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	|               RETURN               |                                                                       |
	//	|             VALUE/CODE             |                              DESCRIPTION                              |
	//	|                                    |                                                                       |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000000 NERR_Success            | The client request succeeded.                                         |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                                     |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | The client request failed because the specified parameter is invalid. |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000008 ERROR_NOT_ENOUGH_MEMORY | Not enough storage is available to process this command.              |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x00000906 NERR_NetNameNotFound    | The share name does not exist.                                        |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//	| 0x0000007C ERROR_INVALID_LEVEL     | The system call level is not correct.                                 |
	//	+------------------------------------+-----------------------------------------------------------------------+
	//
	// The ShareInfo.shi503_netname parameter specifies the name of the share to delete
	// from the ShareList. This MUST be a nonempty null-terminated UTF-16 string; otherwise,
	// the server MUST fail the call with an ERROR_INVALID_PARAMETER error code.
	ShareDeleteEx(context.Context, *ShareDeleteExRequest) (*ShareDeleteExResponse, error)
}

func RegisterSrvsvcServer(conn dcerpc.Conn, o SrvsvcServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewSrvsvcServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(SrvsvcSyntaxV3_0))...)
}

func NewSrvsvcServerHandle(o SrvsvcServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return SrvsvcServerHandle(ctx, o, opNum, r)
	}
}

func SrvsvcServerHandle(ctx context.Context, o SrvsvcServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // Opnum0NotUsedOnWire
		// Opnum0NotUsedOnWire
		return nil, nil
	case 1: // Opnum1NotUsedOnWire
		// Opnum1NotUsedOnWire
		return nil, nil
	case 2: // Opnum2NotUsedOnWire
		// Opnum2NotUsedOnWire
		return nil, nil
	case 3: // Opnum3NotUsedOnWire
		// Opnum3NotUsedOnWire
		return nil, nil
	case 4: // Opnum4NotUsedOnWire
		// Opnum4NotUsedOnWire
		return nil, nil
	case 5: // Opnum5NotUsedOnWire
		// Opnum5NotUsedOnWire
		return nil, nil
	case 6: // Opnum6NotUsedOnWire
		// Opnum6NotUsedOnWire
		return nil, nil
	case 7: // Opnum7NotUsedOnWire
		// Opnum7NotUsedOnWire
		return nil, nil
	case 8: // NetrConnectionEnum
		in := &ConnectionEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ConnectionEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // NetrFileEnum
		in := &FileEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FileEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // NetrFileGetInfo
		in := &FileGetInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FileGetInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // NetrFileClose
		in := &FileCloseRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FileClose(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // NetrSessionEnum
		in := &SessionEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SessionEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // NetrSessionDel
		in := &SessionDeleteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SessionDelete(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // NetrShareAdd
		in := &ShareAddRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ShareAdd(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // NetrShareEnum
		in := &ShareEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ShareEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // NetrShareGetInfo
		in := &ShareGetInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ShareGetInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // NetrShareSetInfo
		in := &ShareSetInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ShareSetInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // NetrShareDel
		in := &ShareDeleteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ShareDelete(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // NetrShareDelSticky
		in := &ShareDeleteStickyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ShareDeleteSticky(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // NetrShareCheck
		in := &ShareCheckRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ShareCheck(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // NetrServerGetInfo
		in := &GetInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // NetrServerSetInfo
		in := &SetInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // NetrServerDiskEnum
		in := &DiskEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DiskEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // NetrServerStatisticsGet
		in := &StatisticsGetRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.StatisticsGet(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // NetrServerTransportAdd
		in := &TransportAddRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.TransportAdd(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // NetrServerTransportEnum
		in := &TransportEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.TransportEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // NetrServerTransportDel
		in := &TransportDeleteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.TransportDelete(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // NetrRemoteTOD
		in := &RemoteToDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoteToD(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // Opnum29NotUsedOnWire
		// Opnum29NotUsedOnWire
		return nil, nil
	case 30: // NetprPathType
		in := &PathTypeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PathType(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // NetprPathCanonicalize
		in := &PathCanonicalizeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PathCanonicalize(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // NetprPathCompare
		in := &PathCompareRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.PathCompare(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // NetprNameValidate
		in := &NameValidateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.NameValidate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // NetprNameCanonicalize
		in := &NameCanonicalizeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.NameCanonicalize(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // NetprNameCompare
		in := &NameCompareRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.NameCompare(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // NetrShareEnumSticky
		in := &ShareEnumStickyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ShareEnumSticky(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // NetrShareDelStart
		in := &ShareDeleteStartRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ShareDeleteStart(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 38: // NetrShareDelCommit
		in := &ShareDeleteCommitRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ShareDeleteCommit(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 39: // NetrpGetFileSecurity
		in := &GetFileSecurityRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFileSecurity(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 40: // NetrpSetFileSecurity
		in := &SetFileSecurityRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFileSecurity(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 41: // NetrServerTransportAddEx
		in := &TransportAddExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.TransportAddEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 42: // Opnum42NotUsedOnWire
		// Opnum42NotUsedOnWire
		return nil, nil
	case 43: // NetrDfsGetVersion
		in := &GetVersionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetVersion(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 44: // NetrDfsCreateLocalPartition
		in := &CreateLocalPartitionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateLocalPartition(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 45: // NetrDfsDeleteLocalPartition
		in := &DeleteLocalPartitionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteLocalPartition(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 46: // NetrDfsSetLocalVolumeState
		in := &SetLocalVolumeStateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetLocalVolumeState(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 47: // Opnum47NotUsedOnWire
		// Opnum47NotUsedOnWire
		return nil, nil
	case 48: // NetrDfsCreateExitPoint
		in := &CreateExitPointRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateExitPoint(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 49: // NetrDfsDeleteExitPoint
		in := &DeleteExitPointRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteExitPoint(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 50: // NetrDfsModifyPrefix
		in := &ModifyPrefixRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ModifyPrefix(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 51: // NetrDfsFixLocalVolume
		in := &FixLocalVolumeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FixLocalVolume(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 52: // NetrDfsManagerReportSiteInfo
		in := &ManagerReportSiteInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ManagerReportSiteInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 53: // NetrServerTransportDelEx
		in := &TransportDeleteExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.TransportDeleteEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 54: // NetrServerAliasAdd
		in := &AliasAddRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AliasAdd(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 55: // NetrServerAliasEnum
		in := &AliasEnumRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AliasEnum(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 56: // NetrServerAliasDel
		in := &AliasDeleteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AliasDelete(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 57: // NetrShareDelEx
		in := &ShareDeleteExRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ShareDeleteEx(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
