package winreg

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

// winreg server interface.
type WinregServer interface {

	// The OpenClassesRoot method is called by the client. In response, the server opens
	// the HKEY_CLASSES_ROOT  predefined key.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | Access is denied.                                                                |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                | server can no longer service registry requests because server shutdown has been  |
	//	|                                | initiated.                                                                       |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	OpenClassesRoot(context.Context, *OpenClassesRootRequest) (*OpenClassesRootResponse, error)

	// The OpenCurrentUser method is called by the client. In response, the server opens
	// a handle to the HKEY_CURRENT_USER key. The server MUST determine which subkey of
	// HKEY_USERS is the correct key to map to HKEY_CURRENT_USER, as explained in section
	// 3.1.1.8.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | Access is denied.                                                                |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                | server can no longer service registry requests because server shutdown has been  |
	//	|                                | initiated (SHUTDOWNINPROGRESS is set to TRUE).                                   |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	OpenCurrentUser(context.Context, *OpenCurrentUserRequest) (*OpenCurrentUserResponse, error)

	// The OpenLocalMachine method is called by the client. In response, the server opens
	// a handle to the HKEY_LOCAL_MACHINE predefined key.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | Access is denied.                                                                |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                | server can no longer service registry requests because server shutdown has been  |
	//	|                                | initiated.                                                                       |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	OpenLocalMachine(context.Context, *OpenLocalMachineRequest) (*OpenLocalMachineResponse, error)

	// The OpenPerformanceData method is called by the client. In response, the server opens
	// a handle to the HKEY_PERFORMANCE_DATA predefined key. The HKEY_PERFORMANCE_DATA  predefined
	// key is used to retrieve performance information from a registry server using only
	// the BaseRegQueryInfoKey, BaseRegQueryValue, BaseRegEnumValue, and BaseRegCloseKey
	// methods.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The server
	// SHOULD return without modification any other error code encountered in servicing
	// the client request.
	//
	// The most common error codes are listed in the following table.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | Access is denied.                                                                |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                | server can no longer service registry requests because server shutdown has been  |
	//	|                                | initiated.                                                                       |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	OpenPerformanceData(context.Context, *OpenPerformanceDataRequest) (*OpenPerformanceDataResponse, error)

	// The OpenUsers method is called by the client. In response, the server opens a handle
	// to the HKEY_USERS predefined key.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The server
	// SHOULD return without modification any error code encountered in servicing the client
	// request.
	//
	// The most common error codes are listed in the following table.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | Access is denied.                                                                |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                | server can no longer service registry requests because server shutdown has been  |
	//	|                                | initiated.                                                                       |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	OpenUsers(context.Context, *OpenUsersRequest) (*OpenUsersResponse, error)

	// The BaseRegCloseKey method is called by the client. In response, the server destroys
	// (closes) the handle to the specified registry key.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000006 ERROR_INVALID_HANDLE | The handle is invalid.                                                           |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000AA ERROR_BUSY           | The requested resource is in use.                                                |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT  | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                 | server can no longer service registry requests because server shutdown has been  |
	//	|                                 | initiated.                                                                       |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000015 ERROR_NOT_READY      | The service is not read. Calls can be repeated at a later time.                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000102 WAIT_TIMEOUT         | The wait operation timed out.                                                    |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	BaseRegCloseKey(context.Context, *BaseRegCloseKeyRequest) (*BaseRegCloseKeyResponse, error)

	// The BaseRegCreateKey method is called by the client. In response, the server creates
	// the specified registry key and returns a handle to the newly created key. If the
	// key already exists in the registry, a handle to the existing key is opened and returned.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The caller does not have KEY_CREATE_SUB_KEY access rights.                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied.                                                                |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter is incorrect.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT     | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                    | server can no longer service registry requests because server shutdown has been  |
	//	|                                    | initiated.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	BaseRegCreateKey(context.Context, *BaseRegCreateKeyRequest) (*BaseRegCreateKeyResponse, error)

	// The BaseRegDeleteKey method is called by the client. In response, the server deletes
	// the specified subkey.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter is incorrect.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Access is denied. For BaseRegDeleteKey, this error will be returned when the key |
	//	|                                    | indicated by the lpSubKey parameter has subkeys.                                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT     | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                    | server can no longer service registry requests because server shutdown has been  |
	//	|                                    | initiated.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	BaseRegDeleteKey(context.Context, *BaseRegDeleteKeyRequest) (*BaseRegDeleteKeyResponse, error)

	// The BaseRegDeleteValue method is called by the client. In response, the server removes
	// a named value from the specified registry key.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The caller does not have KEY_SET_VALUE access rights.                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter is incorrect.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT     | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                    | server can no longer service registry requests because server shutdown has been  |
	//	|                                    | initiated.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	BaseRegDeleteValue(context.Context, *BaseRegDeleteValueRequest) (*BaseRegDeleteValueResponse, error)

	// The BaseRegEnumKey method is called by the client in order to enumerate a subkey.
	// In response, the server returns a requested subkey.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The caller does not have KEY_ENUMERATE_SUB_KEYS access rights.                   |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000000E ERROR_OUTOFMEMORY       | Not enough storage is available to complete this operation.                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter is incorrect.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS     | No more data is available.                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT     | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                    | server can no longer service registry requests because server shutdown has been  |
	//	|                                    | initiated.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA         | The size of the buffer is not large enough to hold the requested data.           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	BaseRegEnumKey(context.Context, *BaseRegEnumKeyRequest) (*BaseRegEnumKeyResponse, error)

	// The BaseRegEnumValue method is called by the client. In response, the server enumerates
	// the value at the specified index for the specified registry key.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2 The most
	// common error codes are listed in the following table.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                |                                                                                  |
	//	|              VALUE/CODE              |                                   DESCRIPTION                                    |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED       | The caller does not have KEY_QUERY_VALUE access rights.                          |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000000E ERROR_OUTOFMEMORY         | Not enough storage is available to complete this operation.                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER   | A parameter is incorrect.                                                        |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000007A ERROR_INSUFFICIENT_BUFFER | The data area passed to a system call is too small.                              |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA           | More data is available.                                                          |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS       | No more data is available.                                                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT       | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                      | server can no longer service registry requests because server shutdown has been  |
	//	|                                      | initiated.                                                                       |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	BaseRegEnumValue(context.Context, *BaseRegEnumValueRequest) (*BaseRegEnumValueResponse, error)

	// The BaseRegFlushKey method is called by the client. In response, the server writes
	// all of the subkeys and values of the key indicated by the hKey parameter to the backing
	// store for registry data.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2.
	//
	//	+--------------------------------+---------------------------------------------------------+
	//	|             RETURN             |                                                         |
	//	|           VALUE/CODE           |                       DESCRIPTION                       |
	//	|                                |                                                         |
	//	+--------------------------------+---------------------------------------------------------+
	//	+--------------------------------+---------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | The caller does not have KEY_QUERY_VALUE access rights. |
	//	+--------------------------------+---------------------------------------------------------+
	BaseRegFlushKey(context.Context, *BaseRegFlushKeyRequest) (*BaseRegFlushKeyResponse, error)

	// The BaseRegGetKeySecurity method is called by the client. In response, the server
	// returns a copy of the security descriptor that protects the specified open registry
	// key.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x0000000E ERROR_OUTOFMEMORY       | Not enough storage is available to complete this operation.                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter is incorrect.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT     | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                    | server can no longer service registry requests because server shutdown has been  |
	//	|                                    | initiated.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	BaseRegGetKeySecurity(context.Context, *BaseRegGetKeySecurityRequest) (*BaseRegGetKeySecurityResponse, error)

	// The BaseRegLoadKey method is called by the client. In response, the server loads
	// key, subkey, and value data from a file and inserts the data into the registry hierarchy.
	//
	// The BaseRegLoadKey method is designed for use in backup and recovery scenarios where
	// the client first loads a registry hive from a file on disk using the BaseRegLoadKey
	// method. Then, after reading or writing key data from the loaded hive, the client
	// uses the BaseRegUnLoadKey method to unload the hive. For example, a backup application
	// loads another user hive (another user's HKEY_CURRENT_USER ) from a file on disk using
	// the BaseRegLoadKey method. After reading key and value data, it will unload the hive
	// using the BaseRegUnLoadKey method.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter is incorrect.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000003F9 ERROR_NOT_REGISTRY_FILE | The system attempted to load or restore a file into the registry, but the        |
	//	|                                    | specified file is not in a registry file format.                                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT     | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                    | server can no longer service registry requests because server shutdown has been  |
	//	|                                    | initiated.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	BaseRegLoadKey(context.Context, *BaseRegLoadKeyRequest) (*BaseRegLoadKeyResponse, error)

	// Opnum14NotImplemented operation.
	// Opnum14NotImplemented

	// The BaseRegOpenKey method is called by the client. In response, the server opens
	// a specified key for access and returns a handle to it.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2.
	BaseRegOpenKey(context.Context, *BaseRegOpenKeyRequest) (*BaseRegOpenKeyResponse, error)

	// The BaseRegQueryInfoKey method is called by the client. In response, the server returns
	// relevant information on the key that corresponds to the specified key handle.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The caller does not have KEY_QUERY_VALUE access rights.                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter is incorrect.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT     | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                    | server can no longer service registry requests because server shutdown has been  |
	//	|                                    | initiated.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA         | The size of the buffer is not large enough to hold the requested data.           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	BaseRegQueryInfoKey(context.Context, *BaseRegQueryInfoKeyRequest) (*BaseRegQueryInfoKeyResponse, error)

	// The BaseRegQueryValue method is called by the client. In response, the server returns
	// the data that is associated with the named value of a specified registry open key.
	// If a value name is not specified, the server returns the data that is associated
	// with the default value of the specified registry open key.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The caller does not have KEY_QUERY_VALUE access rights.                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter is incorrect.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The value specified by lpValueName was not found. If lpValueName was not         |
	//	|                                    | specified, the default value has not been defined.                               |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT     | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                    | server can no longer service registry requests because server shutdown has been  |
	//	|                                    | initiated.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA         | The data to be returned is larger than the buffer provided.                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	BaseRegQueryValue(context.Context, *BaseRegQueryValueRequest) (*BaseRegQueryValueResponse, error)

	// The BaseRegReplaceKey method is called by the client. In response, the server MUST
	// read the registry information from the specified file and replace the specified key
	// with the content of the file. When the system is started again, the key and subkeys
	// have the same values as those in the specified file.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter is incorrect.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT     | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                    | server can no longer service registry requests because server shutdown has been  |
	//	|                                    | initiated.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000011 ERROR_NOT_SAME_DEVICE   | The file indicated by lpOldFile is not on the same physical volume as the file   |
	//	|                                    | indicated by lpNewFile.                                                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	BaseRegReplaceKey(context.Context, *BaseRegReplaceKeyRequest) (*BaseRegReplaceKeyResponse, error)

	// The BaseRegRestoreKey method is called by the client. In response, the server reads
	// the registry information in a specified file and copies it over the specified key.
	// The registry information takes the form of a key and multiple levels of subkeys.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter is incorrect.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT     | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                    | server can no longer service registry requests because server shutdown has been  |
	//	|                                    | initiated.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	BaseRegRestoreKey(context.Context, *BaseRegRestoreKeyRequest) (*BaseRegRestoreKeyResponse, error)

	// The BaseRegSaveKey method is called by the client. In response, the server saves
	// the specified key, subkeys, and values to a new file.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in either [MS-ERREF] section 2.2 or
	// [MS-ERREF] section 2.3.1. The most common error codes are listed in the following
	// table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter is incorrect.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT     | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                    | server can no longer service registry requests because server shutdown has been  |
	//	|                                    | initiated.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	BaseRegSaveKey(context.Context, *BaseRegSaveKeyRequest) (*BaseRegSaveKeyResponse, error)

	// The BaseRegSetKeySecurity method is called by the client. In response, the server
	// sets the security descriptor that protects the specified open registry key.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter is incorrect.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT     | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                    | server can no longer service registry requests because server shutdown has been  |
	//	|                                    | initiated.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	BaseRegSetKeySecurity(context.Context, *BaseRegSetKeySecurityRequest) (*BaseRegSetKeySecurityResponse, error)

	// The BaseRegSetValue method is called by the client. In response, the server sets
	// the data for the specified value of a registry key.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The caller does not have KEY_SET_VALUE access rights, or the value being set to  |
	//	|                                    | a symbolic key is not the literal string "SymbolicLinkValue".                    |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter is incorrect.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT     | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                    | server can no longer service registry requests because server shutdown has been  |
	//	|                                    | initiated.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	BaseRegSetValue(context.Context, *BaseRegSetValueRequest) (*BaseRegSetValueResponse, error)

	// The BaseRegUnLoadKey method is called by the client. In response, the server removes
	// the specified discrete body of keys, subkeys, and values that is rooted at the top
	// of the registry hierarchy.
	//
	// The BaseRegUnLoadKey method is designed for use in backup and recovery scenarios
	// where the client first loads a registry hive from a file on disk using the BaseRegLoadKey
	// method. Then, after reading or writing key data from the loaded hive, the client
	// uses the BaseRegUnLoadKey method to unload the hive. For example, a backup application
	// can load another user hive (another user's HKEY_CURRENT_USER) from a file on disk
	// using the BaseRegLoadKey method. Then, after reading key and value data, it will
	// unload the hive using the BaseRegUnLoadKey method.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter is incorrect.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT     | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                    | server can no longer service registry requests because server shutdown has been  |
	//	|                                    | initiated.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The key specified by the handle hKey and the lpSubKey parameter does not exist   |
	//	|                                    | in the key namespace.                                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The key specified by the handle hKey and the lpSubKey parameter is not a         |
	//	|                                    | descendent key of the HKEY_LOCAL_MACHINE or HKEY_USERS root keys, or there are   |
	//	|                                    | open handles to the key or its descendants.                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	BaseRegUnloadKey(context.Context, *BaseRegUnloadKeyRequest) (*BaseRegUnloadKeyResponse, error)

	// Opnum24NotImplemented operation.
	// Opnum24NotImplemented

	// Opnum25NotImplemented operation.
	// Opnum25NotImplemented

	// The BaseRegGetVersion method is called by the client. In response, the server returns
	// the version of the remote registry server. The BaseRegGetVersion method is used by
	// the client and the server to determine if the remote registry server supports both
	// 32-bit and 64-bit key namespaces.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns the following nonzero error code.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000003E6 ERROR_NOACCESS      | Invalid access to memory location.                                               |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                | server can no longer service registry requests because server shutdown has been  |
	//	|                                | initiated.                                                                       |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	BaseRegGetVersion(context.Context, *BaseRegGetVersionRequest) (*BaseRegGetVersionResponse, error)

	// The OpenCurrentConfig method is called by the client. In response, the server attempts
	// to open a handle to the HKEY_CURRENT_CONFIG predefined key.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | Access is denied.                                                                |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                | server can no longer service registry requests because server shutdown has been  |
	//	|                                | initiated.                                                                       |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	OpenCurrentConfig(context.Context, *OpenCurrentConfigRequest) (*OpenCurrentConfigResponse, error)

	// Opnum28NotImplemented operation.
	// Opnum28NotImplemented

	// The BaseRegQueryMultipleValues method is called by the client. In response, the server
	// returns the type and data for a client-specified list of value names that are associated
	// with the specified registry key.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED        | The caller does not have KEY_QUERY_VALUE access rights.                          |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER    | A parameter is incorrect.                                                        |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000078 ERROR_CALL_NOT_IMPLEMENTED | This function is not supported on this system.                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT        | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                       | server can no longer service registry requests because server shutdown has been  |
	//	|                                       | initiated.                                                                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	BaseRegQueryMultipleValues(context.Context, *BaseRegQueryMultipleValuesRequest) (*BaseRegQueryMultipleValuesResponse, error)

	// Opnum30NotImplemented operation.
	// Opnum30NotImplemented

	// The BaseRegSaveKeyEx method is called by the client. In response, the server saves
	// the specified key, subkeys, and values to a new file. The BaseRegSaveKeyEx method
	// accepts flags that determine the format for the saved key or and values.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter is incorrect.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT     | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                    | server can no longer service registry requests because server shutdown has been  |
	//	|                                    | initiated.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | The server does not have access permissions to save the file at the specified    |
	//	|                                    | location.                                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS    | Cannot create a file when that file already exists.                              |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	BaseRegSaveKeyEx(context.Context, *BaseRegSaveKeyExRequest) (*BaseRegSaveKeyExResponse, error)

	// The OpenPerformanceText method is called by the client. In response, the server opens
	// a handle to the HKEY_PERFORMANCE_TEXT predefined key. The HKEY_PERFORMANCE_TEXT predefined
	// key is used to retrieve performance information from a registry server using only
	// the BaseRegQueryInfoKey, BaseRegQueryValue, BaseRegEnumValue, and BaseRegCloseKey
	// methods.
	//
	// Return Values: This method MUST always return a 0 (ERROR_SUCCESS), even in case of
	// errors.
	//
	//	+-------------------+------------------+
	//	|      RETURN       |                  |
	//	|    VALUE/CODE     |   DESCRIPTION    |
	//	|                   |                  |
	//	+-------------------+------------------+
	//	+-------------------+------------------+
	//	| 0 ERROR_SUCCESS   | Always returned. |
	//	+-------------------+------------------+
	OpenPerformanceText(context.Context, *OpenPerformanceTextRequest) (*OpenPerformanceTextResponse, error)

	// The OpenPerformanceNlsText method is called by the client. In response, the server
	// opens a handle to the HKEY_PERFORMANCE_NLSTEXT predefined key. The HKEY_PERFORMANCE_NLSTEXT
	//  predefined key is used to retrieve performance information from a registry server
	// using only the BaseRegQueryInfoKey, BaseRegQueryValue, BaseRegEnumValue, and BaseRegCloseKey
	// methods.
	//
	// Return Values: This method MUST always return a 0 (ERROR_SUCCESS), even in case of
	// errors.
	//
	//	+-------------------+------------------+
	//	|      RETURN       |                  |
	//	|    VALUE/CODE     |   DESCRIPTION    |
	//	|                   |                  |
	//	+-------------------+------------------+
	//	+-------------------+------------------+
	//	| 0 ERROR_SUCCESS   | Always returned. |
	//	+-------------------+------------------+
	OpenPerformanceNlsText(context.Context, *OpenPerformanceNlsTextRequest) (*OpenPerformanceNlsTextResponse, error)

	// The BaseRegQueryMultipleValues2 method is called by the client. In response, the
	// server returns the type and data for a client-specified list of value names that
	// are associated with the specified registry key.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                                  |
	//	|              VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED        | The caller does not have KEY_QUERY_VALUE access rights.                          |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER    | A parameter is incorrect.                                                        |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000078 ERROR_CALL_NOT_IMPLEMENTED | This function is not supported on this system.                                   |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA            | More data is available.                                                          |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT        | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                       | server can no longer service registry requests because server shutdown has been  |
	//	|                                       | initiated.                                                                       |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	BaseRegQueryMultipleValues2(context.Context, *BaseRegQueryMultipleValues2Request) (*BaseRegQueryMultipleValues2Response, error)

	// The BaseRegDeleteKeyEx method is called by the client. In response, the server deletes
	// the specified registry key.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2. The most
	// common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | A parameter is incorrect.                                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000013 ERROR_WRITE_PROTECT     | A read or write operation was attempted to a volume after it was dismounted. The |
	//	|                                    | server can no longer service registry requests because server shutdown has been  |
	//	|                                    | initiated.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	BaseRegDeleteKeyEx(context.Context, *BaseRegDeleteKeyExRequest) (*BaseRegDeleteKeyExResponse, error)
}

func RegisterWinregServer(conn dcerpc.Conn, o WinregServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWinregServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WinregSyntaxV1_0))...)
}

func NewWinregServerHandle(o WinregServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WinregServerHandle(ctx, o, opNum, r)
	}
}

func WinregServerHandle(ctx context.Context, o WinregServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // OpenClassesRoot
		op := &xxx_OpenClassesRootOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenClassesRootRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenClassesRoot(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // OpenCurrentUser
		op := &xxx_OpenCurrentUserOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenCurrentUserRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenCurrentUser(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // OpenLocalMachine
		op := &xxx_OpenLocalMachineOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenLocalMachineRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenLocalMachine(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // OpenPerformanceData
		op := &xxx_OpenPerformanceDataOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenPerformanceDataRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenPerformanceData(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // OpenUsers
		op := &xxx_OpenUsersOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenUsersRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenUsers(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // BaseRegCloseKey
		op := &xxx_BaseRegCloseKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegCloseKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegCloseKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // BaseRegCreateKey
		op := &xxx_BaseRegCreateKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegCreateKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegCreateKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // BaseRegDeleteKey
		op := &xxx_BaseRegDeleteKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegDeleteKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegDeleteKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 8: // BaseRegDeleteValue
		op := &xxx_BaseRegDeleteValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegDeleteValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegDeleteValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 9: // BaseRegEnumKey
		op := &xxx_BaseRegEnumKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegEnumKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegEnumKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 10: // BaseRegEnumValue
		op := &xxx_BaseRegEnumValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegEnumValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegEnumValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // BaseRegFlushKey
		op := &xxx_BaseRegFlushKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegFlushKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegFlushKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // BaseRegGetKeySecurity
		op := &xxx_BaseRegGetKeySecurityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegGetKeySecurityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegGetKeySecurity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // BaseRegLoadKey
		op := &xxx_BaseRegLoadKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegLoadKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegLoadKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // Opnum14NotImplemented
		// Opnum14NotImplemented
		return nil, nil
	case 15: // BaseRegOpenKey
		op := &xxx_BaseRegOpenKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegOpenKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegOpenKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // BaseRegQueryInfoKey
		op := &xxx_BaseRegQueryInfoKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegQueryInfoKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegQueryInfoKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // BaseRegQueryValue
		op := &xxx_BaseRegQueryValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegQueryValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegQueryValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // BaseRegReplaceKey
		op := &xxx_BaseRegReplaceKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegReplaceKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegReplaceKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // BaseRegRestoreKey
		op := &xxx_BaseRegRestoreKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegRestoreKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegRestoreKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // BaseRegSaveKey
		op := &xxx_BaseRegSaveKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegSaveKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegSaveKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // BaseRegSetKeySecurity
		op := &xxx_BaseRegSetKeySecurityOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegSetKeySecurityRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegSetKeySecurity(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // BaseRegSetValue
		op := &xxx_BaseRegSetValueOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegSetValueRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegSetValue(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // BaseRegUnLoadKey
		op := &xxx_BaseRegUnloadKeyOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegUnloadKeyRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegUnloadKey(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // Opnum24NotImplemented
		// Opnum24NotImplemented
		return nil, nil
	case 25: // Opnum25NotImplemented
		// Opnum25NotImplemented
		return nil, nil
	case 26: // BaseRegGetVersion
		op := &xxx_BaseRegGetVersionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegGetVersionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegGetVersion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 27: // OpenCurrentConfig
		op := &xxx_OpenCurrentConfigOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenCurrentConfigRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenCurrentConfig(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 28: // Opnum28NotImplemented
		// Opnum28NotImplemented
		return nil, nil
	case 29: // BaseRegQueryMultipleValues
		op := &xxx_BaseRegQueryMultipleValuesOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegQueryMultipleValuesRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegQueryMultipleValues(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 30: // Opnum30NotImplemented
		// Opnum30NotImplemented
		return nil, nil
	case 31: // BaseRegSaveKeyEx
		op := &xxx_BaseRegSaveKeyExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegSaveKeyExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegSaveKeyEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 32: // OpenPerformanceText
		op := &xxx_OpenPerformanceTextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenPerformanceTextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenPerformanceText(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 33: // OpenPerformanceNlsText
		op := &xxx_OpenPerformanceNlsTextOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &OpenPerformanceNlsTextRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.OpenPerformanceNlsText(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 34: // BaseRegQueryMultipleValues2
		op := &xxx_BaseRegQueryMultipleValues2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegQueryMultipleValues2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegQueryMultipleValues2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 35: // BaseRegDeleteKeyEx
		op := &xxx_BaseRegDeleteKeyExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &BaseRegDeleteKeyExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.BaseRegDeleteKeyEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented winreg
type UnimplementedWinregServer struct {
}

func (UnimplementedWinregServer) OpenClassesRoot(context.Context, *OpenClassesRootRequest) (*OpenClassesRootResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) OpenCurrentUser(context.Context, *OpenCurrentUserRequest) (*OpenCurrentUserResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) OpenLocalMachine(context.Context, *OpenLocalMachineRequest) (*OpenLocalMachineResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) OpenPerformanceData(context.Context, *OpenPerformanceDataRequest) (*OpenPerformanceDataResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) OpenUsers(context.Context, *OpenUsersRequest) (*OpenUsersResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegCloseKey(context.Context, *BaseRegCloseKeyRequest) (*BaseRegCloseKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegCreateKey(context.Context, *BaseRegCreateKeyRequest) (*BaseRegCreateKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegDeleteKey(context.Context, *BaseRegDeleteKeyRequest) (*BaseRegDeleteKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegDeleteValue(context.Context, *BaseRegDeleteValueRequest) (*BaseRegDeleteValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegEnumKey(context.Context, *BaseRegEnumKeyRequest) (*BaseRegEnumKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegEnumValue(context.Context, *BaseRegEnumValueRequest) (*BaseRegEnumValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegFlushKey(context.Context, *BaseRegFlushKeyRequest) (*BaseRegFlushKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegGetKeySecurity(context.Context, *BaseRegGetKeySecurityRequest) (*BaseRegGetKeySecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegLoadKey(context.Context, *BaseRegLoadKeyRequest) (*BaseRegLoadKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegOpenKey(context.Context, *BaseRegOpenKeyRequest) (*BaseRegOpenKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegQueryInfoKey(context.Context, *BaseRegQueryInfoKeyRequest) (*BaseRegQueryInfoKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegQueryValue(context.Context, *BaseRegQueryValueRequest) (*BaseRegQueryValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegReplaceKey(context.Context, *BaseRegReplaceKeyRequest) (*BaseRegReplaceKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegRestoreKey(context.Context, *BaseRegRestoreKeyRequest) (*BaseRegRestoreKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegSaveKey(context.Context, *BaseRegSaveKeyRequest) (*BaseRegSaveKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegSetKeySecurity(context.Context, *BaseRegSetKeySecurityRequest) (*BaseRegSetKeySecurityResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegSetValue(context.Context, *BaseRegSetValueRequest) (*BaseRegSetValueResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegUnloadKey(context.Context, *BaseRegUnloadKeyRequest) (*BaseRegUnloadKeyResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegGetVersion(context.Context, *BaseRegGetVersionRequest) (*BaseRegGetVersionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) OpenCurrentConfig(context.Context, *OpenCurrentConfigRequest) (*OpenCurrentConfigResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegQueryMultipleValues(context.Context, *BaseRegQueryMultipleValuesRequest) (*BaseRegQueryMultipleValuesResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegSaveKeyEx(context.Context, *BaseRegSaveKeyExRequest) (*BaseRegSaveKeyExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) OpenPerformanceText(context.Context, *OpenPerformanceTextRequest) (*OpenPerformanceTextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) OpenPerformanceNlsText(context.Context, *OpenPerformanceNlsTextRequest) (*OpenPerformanceNlsTextResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegQueryMultipleValues2(context.Context, *BaseRegQueryMultipleValues2Request) (*BaseRegQueryMultipleValues2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedWinregServer) BaseRegDeleteKeyEx(context.Context, *BaseRegDeleteKeyExRequest) (*BaseRegDeleteKeyExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ WinregServer = (*UnimplementedWinregServer)(nil)
