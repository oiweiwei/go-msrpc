package winreg

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcetypes "github.com/oiweiwei/go-msrpc/msrpc/dcetypes"
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
	_ = dcetypes.GoPackage
)

var (
	// import guard
	GoPackage = "rrp"
)

var (
	// Syntax UUID
	WinregSyntaxUUID = &uuid.UUID{TimeLow: 0x338cd001, TimeMid: 0x2244, TimeHiAndVersion: 0x31f1, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0xaa, Node: [6]uint8{0x90, 0x0, 0x38, 0x0, 0x10, 0x3}}
	// Syntax ID
	WinregSyntaxV1_0 = &dcerpc.SyntaxID{IfUUID: WinregSyntaxUUID, IfVersionMajor: 1, IfVersionMinor: 0}
)

// winreg interface.
type WinregClient interface {

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
	OpenClassesRoot(context.Context, *OpenClassesRootRequest, ...dcerpc.CallOption) (*OpenClassesRootResponse, error)

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
	OpenCurrentUser(context.Context, *OpenCurrentUserRequest, ...dcerpc.CallOption) (*OpenCurrentUserResponse, error)

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
	OpenLocalMachine(context.Context, *OpenLocalMachineRequest, ...dcerpc.CallOption) (*OpenLocalMachineResponse, error)

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
	OpenPerformanceData(context.Context, *OpenPerformanceDataRequest, ...dcerpc.CallOption) (*OpenPerformanceDataResponse, error)

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
	OpenUsers(context.Context, *OpenUsersRequest, ...dcerpc.CallOption) (*OpenUsersResponse, error)

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
	BaseRegCloseKey(context.Context, *BaseRegCloseKeyRequest, ...dcerpc.CallOption) (*BaseRegCloseKeyResponse, error)

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
	BaseRegCreateKey(context.Context, *BaseRegCreateKeyRequest, ...dcerpc.CallOption) (*BaseRegCreateKeyResponse, error)

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
	BaseRegDeleteKey(context.Context, *BaseRegDeleteKeyRequest, ...dcerpc.CallOption) (*BaseRegDeleteKeyResponse, error)

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
	BaseRegDeleteValue(context.Context, *BaseRegDeleteValueRequest, ...dcerpc.CallOption) (*BaseRegDeleteValueResponse, error)

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
	BaseRegEnumKey(context.Context, *BaseRegEnumKeyRequest, ...dcerpc.CallOption) (*BaseRegEnumKeyResponse, error)

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
	BaseRegEnumValue(context.Context, *BaseRegEnumValueRequest, ...dcerpc.CallOption) (*BaseRegEnumValueResponse, error)

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
	BaseRegFlushKey(context.Context, *BaseRegFlushKeyRequest, ...dcerpc.CallOption) (*BaseRegFlushKeyResponse, error)

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
	BaseRegGetKeySecurity(context.Context, *BaseRegGetKeySecurityRequest, ...dcerpc.CallOption) (*BaseRegGetKeySecurityResponse, error)

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
	BaseRegLoadKey(context.Context, *BaseRegLoadKeyRequest, ...dcerpc.CallOption) (*BaseRegLoadKeyResponse, error)

	// Opnum14NotImplemented operation.
	// Opnum14NotImplemented

	// The BaseRegOpenKey method is called by the client. In response, the server opens
	// a specified key for access and returns a handle to it.
	//
	// Return Values: The method returns 0 (ERROR_SUCCESS) to indicate success; otherwise,
	// it returns a nonzero error code, as specified in [MS-ERREF] section 2.2.
	BaseRegOpenKey(context.Context, *BaseRegOpenKeyRequest, ...dcerpc.CallOption) (*BaseRegOpenKeyResponse, error)

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
	BaseRegQueryInfoKey(context.Context, *BaseRegQueryInfoKeyRequest, ...dcerpc.CallOption) (*BaseRegQueryInfoKeyResponse, error)

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
	BaseRegQueryValue(context.Context, *BaseRegQueryValueRequest, ...dcerpc.CallOption) (*BaseRegQueryValueResponse, error)

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
	BaseRegReplaceKey(context.Context, *BaseRegReplaceKeyRequest, ...dcerpc.CallOption) (*BaseRegReplaceKeyResponse, error)

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
	BaseRegRestoreKey(context.Context, *BaseRegRestoreKeyRequest, ...dcerpc.CallOption) (*BaseRegRestoreKeyResponse, error)

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
	BaseRegSaveKey(context.Context, *BaseRegSaveKeyRequest, ...dcerpc.CallOption) (*BaseRegSaveKeyResponse, error)

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
	BaseRegSetKeySecurity(context.Context, *BaseRegSetKeySecurityRequest, ...dcerpc.CallOption) (*BaseRegSetKeySecurityResponse, error)

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
	BaseRegSetValue(context.Context, *BaseRegSetValueRequest, ...dcerpc.CallOption) (*BaseRegSetValueResponse, error)

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
	BaseRegUnloadKey(context.Context, *BaseRegUnloadKeyRequest, ...dcerpc.CallOption) (*BaseRegUnloadKeyResponse, error)

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
	BaseRegGetVersion(context.Context, *BaseRegGetVersionRequest, ...dcerpc.CallOption) (*BaseRegGetVersionResponse, error)

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
	OpenCurrentConfig(context.Context, *OpenCurrentConfigRequest, ...dcerpc.CallOption) (*OpenCurrentConfigResponse, error)

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
	BaseRegQueryMultipleValues(context.Context, *BaseRegQueryMultipleValuesRequest, ...dcerpc.CallOption) (*BaseRegQueryMultipleValuesResponse, error)

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
	BaseRegSaveKeyEx(context.Context, *BaseRegSaveKeyExRequest, ...dcerpc.CallOption) (*BaseRegSaveKeyExResponse, error)

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
	OpenPerformanceText(context.Context, *OpenPerformanceTextRequest, ...dcerpc.CallOption) (*OpenPerformanceTextResponse, error)

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
	OpenPerformanceNlsText(context.Context, *OpenPerformanceNlsTextRequest, ...dcerpc.CallOption) (*OpenPerformanceNlsTextResponse, error)

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
	BaseRegQueryMultipleValues2(context.Context, *BaseRegQueryMultipleValues2Request, ...dcerpc.CallOption) (*BaseRegQueryMultipleValues2Response, error)

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
	BaseRegDeleteKeyEx(context.Context, *BaseRegDeleteKeyExRequest, ...dcerpc.CallOption) (*BaseRegDeleteKeyExResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// OwnerSecurityInformation represents the OWNER_SECURITY_INFORMATION RPC constant
const OwnerSecurityInformation = 0x00000001

// GroupSecurityInformation represents the GROUP_SECURITY_INFORMATION RPC constant
const GroupSecurityInformation = 0x00000002

// DACLSecurityInformation represents the DACL_SECURITY_INFORMATION RPC constant
const DACLSecurityInformation = 0x00000004

// SACLSecurityInformation represents the SACL_SECURITY_INFORMATION RPC constant
const SACLSecurityInformation = 0x00000008

// RegBinary represents the REG_BINARY RPC constant
const RegBinary = 0x00000003

// RegDword represents the REG_DWORD RPC constant
const RegDword = 0x00000004

// RegDwordLittleEndian represents the REG_DWORD_LITTLE_ENDIAN RPC constant
const RegDwordLittleEndian = 0x00000004

// RegDwordBigEndian represents the REG_DWORD_BIG_ENDIAN RPC constant
const RegDwordBigEndian = 0x00000005

// RegExpandString represents the REG_EXPAND_SZ RPC constant
const RegExpandString = 0x00000002

// RegLink represents the REG_LINK RPC constant
const RegLink = 0x00000006

// RegMultistring represents the REG_MULTI_SZ RPC constant
const RegMultistring = 0x00000007

// RegNone represents the REG_NONE RPC constant
const RegNone = 0x00000000

// RegQword represents the REG_QWORD RPC constant
const RegQword = 0x0000000B

// RegQwordLittleEndian represents the REG_QWORD_LITTLE_ENDIAN RPC constant
const RegQwordLittleEndian = 0x0000000B

// RegString represents the REG_SZ RPC constant
const RegString = 0x00000001

// KeyQueryValue represents the KEY_QUERY_VALUE RPC constant
const KeyQueryValue = 0x00000001

// KeySetValue represents the KEY_SET_VALUE RPC constant
const KeySetValue = 0x00000002

// KeyCreateSubKey represents the KEY_CREATE_SUB_KEY RPC constant
const KeyCreateSubKey = 0x00000004

// KeyEnumerateSubKeys represents the KEY_ENUMERATE_SUB_KEYS RPC constant
const KeyEnumerateSubKeys = 0x00000008

// KeyCreateLink represents the KEY_CREATE_LINK RPC constant
const KeyCreateLink = 0x00000020

// KeyWOW6464Key represents the KEY_WOW64_64KEY RPC constant
const KeyWOW6464Key = 0x00000100

// KeyWOW6432Key represents the KEY_WOW64_32KEY RPC constant
const KeyWOW6432Key = 0x00000200

// UnicodeString structure represents RRP_UNICODE_STRING RPC structure.
type UnicodeString dtyp.UnicodeString

func (o *UnicodeString) UnicodeString() *dtyp.UnicodeString { return (*dtyp.UnicodeString)(o) }

func (o *UnicodeString) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != "" && o.MaximumLength == 0 {
		o.MaximumLength = uint16((len(o.Buffer) * 2))
	}
	if o.Buffer != "" && o.Length == 0 {
		o.Length = uint16((len(o.Buffer) * 2))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UnicodeString) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.MaximumLength); err != nil {
		return err
	}
	if o.Buffer != "" || (o.MaximumLength/2) > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64((o.MaximumLength / 2))
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := uint64((o.Length / 2))
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
			_Buffer_buf := utf16.Encode([]rune(o.Buffer))
			if uint64(len(_Buffer_buf)) > sizeInfo[0] {
				_Buffer_buf = _Buffer_buf[:sizeInfo[0]]
			}
			for i1 := range _Buffer_buf {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(_Buffer_buf[i1]); err != nil {
					return err
				}
			}
			for i1 := len(_Buffer_buf); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
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
func (o *UnicodeString) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaximumLength); err != nil {
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
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		var _Buffer_buf []uint16
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array _Buffer_buf", sizeInfo[0])
		}
		_Buffer_buf = make([]uint16, sizeInfo[0])
		for i1 := range _Buffer_buf {
			i1 := i1
			if err := w.ReadData(&_Buffer_buf[i1]); err != nil {
				return err
			}
		}
		o.Buffer = strings.TrimRight(string(utf16.Decode(_Buffer_buf)), ndr.ZeroString)
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*string) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// Key structure represents RPC_HKEY RPC structure.
type Key dcetypes.ContextHandle

func (o *Key) ContextHandle() *dcetypes.ContextHandle { return (*dcetypes.ContextHandle)(o) }

func (o *Key) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Key) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Attributes); err != nil {
		return err
	}
	if o.UUID != nil {
		if err := o.UUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Key) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Attributes); err != nil {
		return err
	}
	if o.UUID == nil {
		o.UUID = &dtyp.GUID{}
	}
	if err := o.UUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// ValueEntry structure represents RVALENT RPC structure.
//
// The RVALENT structure is used to store the values and data that are associated with
// a key, as specified in section 3.1.5.26. The format of the RVALENT structure is as
// follows.
type ValueEntry struct {
	// ve_valuename:  A pointer to an RRP_UNICODE_STRING structure that MUST contain the
	// name of the specified value to be retrieved.
	ValueName *dtyp.UnicodeString `idl:"name:ve_valuename" json:"value_name"`
	// ve_valuelen:   The length in bytes of the ve_valueptr buffer.
	ValueLength uint32 `idl:"name:ve_valuelen" json:"value_length"`
	// ve_valueptr:  A pointer to the data that is associated with a specified value.
	ValuePointer uint32 `idl:"name:ve_valueptr" json:"value_pointer"`
	// ve_type:  The type of the data that is associated with a specified value. For additional
	// specification of the possible values, see section 3.1.1.5.
	//
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	|                            |                                                                                  |
	//	|           VALUE            |                                     MEANING                                      |
	//	|                            |                                                                                  |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| REG_BINARY 3               | Binary data in any form.                                                         |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| REG_DWORD 4                | A 32-bit number.                                                                 |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| REG_DWORD_LITTLE_ENDIAN 4  | A 32-bit number in little-endian format.                                         |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| REG_DWORD_BIG_ENDIAN 5     | A 32-bit number in big-endian format.                                            |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| REG_EXPAND_SZ 2            | A null-terminated string that contains unexpanded references to environment      |
	//	|                            | variables (for example, "%PATH%"). It will be a Unicode or system code page      |
	//	|                            | string, depending on the functions used to manipulate the string.                |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| REG_LINK 6                 | A symbolic link.                                                                 |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| REG_MULTI_SZ 7             | A sequence of null-terminated strings, terminated by an empty string (\0).       |
	//	|                            | For example: String1\0String2\0String3\0LastString\0\0 The first \0 terminates   |
	//	|                            | the first string, the second to the last \0 terminates the last string, and      |
	//	|                            | the final \0 terminates the sequence. Note that the final terminator MUST be     |
	//	|                            | factored into the length of the string.                                          |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| REG_NONE 0                 | No defined value type.                                                           |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| REG_QWORD 11               | A 64-bit number.                                                                 |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| REG_QWORD_LITTLE_ENDIAN 11 | A 64-bit number in little-endian format.                                         |
	//	+----------------------------+----------------------------------------------------------------------------------+
	//	| REG_SZ 1                   | A null-terminated string. This string is either a Unicode or an system code page |
	//	|                            | string, depending on the functions used to manipulate the string.                |
	//	+----------------------------+----------------------------------------------------------------------------------+
	Type uint32 `idl:"name:ve_type" json:"type"`
}

func (o *ValueEntry) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *ValueEntry) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.ValueName != nil {
		_ptr_ve_valuename := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.ValueName != nil {
				if err := o.ValueName.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.ValueName, _ptr_ve_valuename); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.ValueLength); err != nil {
		return err
	}
	// XXX pointer to primitive type, default behavior is to write non-null pointer.
	// if this behavior is not desired, use goext_null_if(cond) attribute.
	_ptr_ve_valueptr := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
		if err := w.WriteData(o.ValuePointer); err != nil {
			return err
		}
		return nil
	})
	if err := w.WritePointer(&o.ValuePointer, _ptr_ve_valueptr); err != nil {
		return err
	}
	if err := w.WriteData(o.Type); err != nil {
		return err
	}
	return nil
}
func (o *ValueEntry) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_ve_valuename := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.ValueName == nil {
			o.ValueName = &dtyp.UnicodeString{}
		}
		if err := o.ValueName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_ve_valuename := func(ptr interface{}) { o.ValueName = *ptr.(**dtyp.UnicodeString) }
	if err := w.ReadPointer(&o.ValueName, _s_ve_valuename, _ptr_ve_valuename); err != nil {
		return err
	}
	if err := w.ReadData(&o.ValueLength); err != nil {
		return err
	}
	_ptr_ve_valueptr := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := w.ReadData(&o.ValuePointer); err != nil {
			return err
		}
		return nil
	})
	_s_ve_valueptr := func(ptr interface{}) { o.ValuePointer = *ptr.(*uint32) }
	if err := w.ReadPointer(&o.ValuePointer, _s_ve_valueptr, _ptr_ve_valueptr); err != nil {
		return err
	}
	if err := w.ReadData(&o.Type); err != nil {
		return err
	}
	return nil
}

// SecurityDescriptor structure represents RPC_SECURITY_DESCRIPTOR RPC structure.
//
// The RPC_SECURITY_DESCRIPTOR structure represents the RPC security descriptors.
type SecurityDescriptor struct {
	// lpSecurityDescriptor:  A buffer that contains a SECURITY_DESCRIPTOR, as specified
	// in [MS-DTYP] section 2.4.6.
	SecurityDescriptor []byte `idl:"name:lpSecurityDescriptor;size_is:(cbInSecurityDescriptor);length_is:(cbOutSecurityDescriptor)" json:"security_descriptor"`
	// cbInSecurityDescriptor:  The size in bytes of the security descriptor.
	InSecurityDescriptorLength uint32 `idl:"name:cbInSecurityDescriptor" json:"in_security_descriptor_length"`
	// cbOutSecurityDescriptor:  The size in bytes of the security descriptor.
	OutSecurityDescriptorLength uint32 `idl:"name:cbOutSecurityDescriptor" json:"out_security_descriptor_length"`
}

func (o *SecurityDescriptor) xxx_PreparePayload(ctx context.Context) error {
	if o.SecurityDescriptor != nil && o.InSecurityDescriptorLength == 0 {
		o.InSecurityDescriptorLength = uint32(len(o.SecurityDescriptor))
	}
	if o.SecurityDescriptor != nil && o.OutSecurityDescriptorLength == 0 {
		o.OutSecurityDescriptorLength = uint32(len(o.SecurityDescriptor))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityDescriptor) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.SecurityDescriptor != nil || o.InSecurityDescriptorLength > 0 {
		_ptr_lpSecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.InSecurityDescriptorLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := uint64(o.OutSecurityDescriptorLength)
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
			for i1 := range o.SecurityDescriptor {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.SecurityDescriptor[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.SecurityDescriptor); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint8(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SecurityDescriptor, _ptr_lpSecurityDescriptor); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.InSecurityDescriptorLength); err != nil {
		return err
	}
	if err := w.WriteData(o.OutSecurityDescriptorLength); err != nil {
		return err
	}
	return nil
}
func (o *SecurityDescriptor) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_lpSecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.SecurityDescriptor", sizeInfo[0])
		}
		o.SecurityDescriptor = make([]byte, sizeInfo[0])
		for i1 := range o.SecurityDescriptor {
			i1 := i1
			if err := w.ReadData(&o.SecurityDescriptor[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_lpSecurityDescriptor := func(ptr interface{}) { o.SecurityDescriptor = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.SecurityDescriptor, _s_lpSecurityDescriptor, _ptr_lpSecurityDescriptor); err != nil {
		return err
	}
	if err := w.ReadData(&o.InSecurityDescriptorLength); err != nil {
		return err
	}
	if err := w.ReadData(&o.OutSecurityDescriptorLength); err != nil {
		return err
	}
	return nil
}

// SecurityAttributes structure represents RPC_SECURITY_ATTRIBUTES RPC structure.
//
// The RPC_SECURITY_ATTRIBUTES structure represents security attributes that can be
// set through the Remote Procedure Call Protocol Extensions, as specified in [MS-CMRP]
// section 2.2.3.2.
type SecurityAttributes struct {
	// nLength:  The length in bytes of the security descriptor.
	Length uint32 `idl:"name:nLength" json:"length"`
	// RpcSecurityDescriptor:  The security descriptor that MUST be as specified in RPC_SECURITY_DESCRIPTOR.
	SecurityDescriptor *SecurityDescriptor `idl:"name:RpcSecurityDescriptor" json:"security_descriptor"`
	// bInheritHandle:  TRUE if the new process inherits the handle; otherwise, FALSE.
	InheritHandle bool `idl:"name:bInheritHandle" json:"inherit_handle"`
}

func (o *SecurityAttributes) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SecurityAttributes) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if o.SecurityDescriptor != nil {
		if err := o.SecurityDescriptor.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&SecurityDescriptor{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.InheritHandle); err != nil {
		return err
	}
	return nil
}
func (o *SecurityAttributes) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if o.SecurityDescriptor == nil {
		o.SecurityDescriptor = &SecurityDescriptor{}
	}
	if err := o.SecurityDescriptor.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.InheritHandle); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultWinregClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultWinregClient) OpenClassesRoot(ctx context.Context, in *OpenClassesRootRequest, opts ...dcerpc.CallOption) (*OpenClassesRootResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenClassesRootResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) OpenCurrentUser(ctx context.Context, in *OpenCurrentUserRequest, opts ...dcerpc.CallOption) (*OpenCurrentUserResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenCurrentUserResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) OpenLocalMachine(ctx context.Context, in *OpenLocalMachineRequest, opts ...dcerpc.CallOption) (*OpenLocalMachineResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenLocalMachineResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) OpenPerformanceData(ctx context.Context, in *OpenPerformanceDataRequest, opts ...dcerpc.CallOption) (*OpenPerformanceDataResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenPerformanceDataResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) OpenUsers(ctx context.Context, in *OpenUsersRequest, opts ...dcerpc.CallOption) (*OpenUsersResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenUsersResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegCloseKey(ctx context.Context, in *BaseRegCloseKeyRequest, opts ...dcerpc.CallOption) (*BaseRegCloseKeyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegCloseKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegCreateKey(ctx context.Context, in *BaseRegCreateKeyRequest, opts ...dcerpc.CallOption) (*BaseRegCreateKeyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegCreateKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegDeleteKey(ctx context.Context, in *BaseRegDeleteKeyRequest, opts ...dcerpc.CallOption) (*BaseRegDeleteKeyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegDeleteKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegDeleteValue(ctx context.Context, in *BaseRegDeleteValueRequest, opts ...dcerpc.CallOption) (*BaseRegDeleteValueResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegDeleteValueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegEnumKey(ctx context.Context, in *BaseRegEnumKeyRequest, opts ...dcerpc.CallOption) (*BaseRegEnumKeyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegEnumKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegEnumValue(ctx context.Context, in *BaseRegEnumValueRequest, opts ...dcerpc.CallOption) (*BaseRegEnumValueResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegEnumValueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegFlushKey(ctx context.Context, in *BaseRegFlushKeyRequest, opts ...dcerpc.CallOption) (*BaseRegFlushKeyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegFlushKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegGetKeySecurity(ctx context.Context, in *BaseRegGetKeySecurityRequest, opts ...dcerpc.CallOption) (*BaseRegGetKeySecurityResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegGetKeySecurityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegLoadKey(ctx context.Context, in *BaseRegLoadKeyRequest, opts ...dcerpc.CallOption) (*BaseRegLoadKeyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegLoadKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegOpenKey(ctx context.Context, in *BaseRegOpenKeyRequest, opts ...dcerpc.CallOption) (*BaseRegOpenKeyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegOpenKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegQueryInfoKey(ctx context.Context, in *BaseRegQueryInfoKeyRequest, opts ...dcerpc.CallOption) (*BaseRegQueryInfoKeyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegQueryInfoKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegQueryValue(ctx context.Context, in *BaseRegQueryValueRequest, opts ...dcerpc.CallOption) (*BaseRegQueryValueResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegQueryValueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegReplaceKey(ctx context.Context, in *BaseRegReplaceKeyRequest, opts ...dcerpc.CallOption) (*BaseRegReplaceKeyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegReplaceKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegRestoreKey(ctx context.Context, in *BaseRegRestoreKeyRequest, opts ...dcerpc.CallOption) (*BaseRegRestoreKeyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegRestoreKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegSaveKey(ctx context.Context, in *BaseRegSaveKeyRequest, opts ...dcerpc.CallOption) (*BaseRegSaveKeyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegSaveKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegSetKeySecurity(ctx context.Context, in *BaseRegSetKeySecurityRequest, opts ...dcerpc.CallOption) (*BaseRegSetKeySecurityResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegSetKeySecurityResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegSetValue(ctx context.Context, in *BaseRegSetValueRequest, opts ...dcerpc.CallOption) (*BaseRegSetValueResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegSetValueResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegUnloadKey(ctx context.Context, in *BaseRegUnloadKeyRequest, opts ...dcerpc.CallOption) (*BaseRegUnloadKeyResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegUnloadKeyResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegGetVersion(ctx context.Context, in *BaseRegGetVersionRequest, opts ...dcerpc.CallOption) (*BaseRegGetVersionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegGetVersionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) OpenCurrentConfig(ctx context.Context, in *OpenCurrentConfigRequest, opts ...dcerpc.CallOption) (*OpenCurrentConfigResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenCurrentConfigResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegQueryMultipleValues(ctx context.Context, in *BaseRegQueryMultipleValuesRequest, opts ...dcerpc.CallOption) (*BaseRegQueryMultipleValuesResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegQueryMultipleValuesResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegSaveKeyEx(ctx context.Context, in *BaseRegSaveKeyExRequest, opts ...dcerpc.CallOption) (*BaseRegSaveKeyExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegSaveKeyExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) OpenPerformanceText(ctx context.Context, in *OpenPerformanceTextRequest, opts ...dcerpc.CallOption) (*OpenPerformanceTextResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenPerformanceTextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) OpenPerformanceNlsText(ctx context.Context, in *OpenPerformanceNlsTextRequest, opts ...dcerpc.CallOption) (*OpenPerformanceNlsTextResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &OpenPerformanceNlsTextResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegQueryMultipleValues2(ctx context.Context, in *BaseRegQueryMultipleValues2Request, opts ...dcerpc.CallOption) (*BaseRegQueryMultipleValues2Response, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegQueryMultipleValues2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) BaseRegDeleteKeyEx(ctx context.Context, in *BaseRegDeleteKeyExRequest, opts ...dcerpc.CallOption) (*BaseRegDeleteKeyExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &BaseRegDeleteKeyExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultWinregClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultWinregClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewWinregClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (WinregClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(WinregSyntaxV1_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultWinregClient{cc: cc}, nil
}

// xxx_OpenClassesRootOperation structure represents the OpenClassesRoot operation
type xxx_OpenClassesRootOperation struct {
	ServerName    string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
	Key           *Key   `idl:"name:phKey" json:"key"`
	Return        uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenClassesRootOperation) OpNum() int { return 0 }

func (o *xxx_OpenClassesRootOperation) OpName() string { return "/winreg/v1/OpenClassesRoot" }

func (o *xxx_OpenClassesRootOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenClassesRootOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME, names=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// samDesired {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenClassesRootOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME, names=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// samDesired {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenClassesRootOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenClassesRootOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phKey {out} (1:{alias=PRPC_HKEY}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_OpenClassesRootOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phKey {out} (1:{alias=PRPC_HKEY,pointer=ref}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
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

// OpenClassesRootRequest structure represents the OpenClassesRoot operation request
type OpenClassesRootRequest struct {
	// ServerName: The server name. The ServerName SHOULD be sent as NULL, and MUST be ignored
	// when it is received because binding to the server is already complete at this stage.
	ServerName string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	// samDesired: A bit field that describes the requested security access for the key.
	// It MUST be constructed from one or more of the values specified in section 2.2.3.
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
}

func (o *OpenClassesRootRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenClassesRootOperation) *xxx_OpenClassesRootOperation {
	if op == nil {
		op = &xxx_OpenClassesRootOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.DesiredAccess = o.DesiredAccess
	return op
}

func (o *OpenClassesRootRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenClassesRootOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.DesiredAccess = op.DesiredAccess
}
func (o *OpenClassesRootRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenClassesRootRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenClassesRootOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenClassesRootResponse structure represents the OpenClassesRoot operation response
type OpenClassesRootResponse struct {
	// phKey: A pointer to an RPC context handle for the root key, HKEY_CLASSES_ROOT, as
	// specified in section 3.1.1.7. The handle is found in the handle table (HANDLETABLE).
	Key *Key `idl:"name:phKey" json:"key"`
	// Return: The OpenClassesRoot return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *OpenClassesRootResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenClassesRootOperation) *xxx_OpenClassesRootOperation {
	if op == nil {
		op = &xxx_OpenClassesRootOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.Return = o.Return
	return op
}

func (o *OpenClassesRootResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenClassesRootOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.Return = op.Return
}
func (o *OpenClassesRootResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenClassesRootResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenClassesRootOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenCurrentUserOperation structure represents the OpenCurrentUser operation
type xxx_OpenCurrentUserOperation struct {
	ServerName    string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
	Key           *Key   `idl:"name:phKey" json:"key"`
	Return        uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenCurrentUserOperation) OpNum() int { return 1 }

func (o *xxx_OpenCurrentUserOperation) OpName() string { return "/winreg/v1/OpenCurrentUser" }

func (o *xxx_OpenCurrentUserOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenCurrentUserOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME, names=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// samDesired {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenCurrentUserOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME, names=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// samDesired {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenCurrentUserOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenCurrentUserOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phKey {out} (1:{alias=PRPC_HKEY}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_OpenCurrentUserOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phKey {out} (1:{alias=PRPC_HKEY,pointer=ref}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
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

// OpenCurrentUserRequest structure represents the OpenCurrentUser operation request
type OpenCurrentUserRequest struct {
	// ServerName: SHOULD be sent as NULL, and MUST be ignored on receipt because the binding
	// to the server is already complete at this stage.
	ServerName string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	// samDesired: A bit field that describes the wanted security access for the key. It
	// MUST be constructed from one or more of the values that are specified in section
	// 2.2.3.
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
}

func (o *OpenCurrentUserRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenCurrentUserOperation) *xxx_OpenCurrentUserOperation {
	if op == nil {
		op = &xxx_OpenCurrentUserOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.DesiredAccess = o.DesiredAccess
	return op
}

func (o *OpenCurrentUserRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenCurrentUserOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.DesiredAccess = op.DesiredAccess
}
func (o *OpenCurrentUserRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenCurrentUserRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenCurrentUserOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenCurrentUserResponse structure represents the OpenCurrentUser operation response
type OpenCurrentUserResponse struct {
	// phKey: A pointer to an RPC context handle for the root key, HKEY_CURRENT_USER, as
	// specified in section 3.1.1.7. The handle is found in the handle table (HANDLETABLE).
	Key *Key `idl:"name:phKey" json:"key"`
	// Return: The OpenCurrentUser return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *OpenCurrentUserResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenCurrentUserOperation) *xxx_OpenCurrentUserOperation {
	if op == nil {
		op = &xxx_OpenCurrentUserOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.Return = o.Return
	return op
}

func (o *OpenCurrentUserResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenCurrentUserOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.Return = op.Return
}
func (o *OpenCurrentUserResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenCurrentUserResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenCurrentUserOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenLocalMachineOperation structure represents the OpenLocalMachine operation
type xxx_OpenLocalMachineOperation struct {
	ServerName    string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
	Key           *Key   `idl:"name:phKey" json:"key"`
	Return        uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenLocalMachineOperation) OpNum() int { return 2 }

func (o *xxx_OpenLocalMachineOperation) OpName() string { return "/winreg/v1/OpenLocalMachine" }

func (o *xxx_OpenLocalMachineOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenLocalMachineOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME, names=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// samDesired {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenLocalMachineOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME, names=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// samDesired {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenLocalMachineOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenLocalMachineOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phKey {out} (1:{alias=PRPC_HKEY}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_OpenLocalMachineOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phKey {out} (1:{alias=PRPC_HKEY,pointer=ref}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
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

// OpenLocalMachineRequest structure represents the OpenLocalMachine operation request
type OpenLocalMachineRequest struct {
	// ServerName: SHOULD be sent as NULL and MUST be ignored on receipt because the binding
	// to the server is already complete at this stage.
	ServerName string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	// samDesired: A bit field that describes the wanted security access for the key. It
	// MUST be constructed from one or more of the values that are specified in section
	// 2.2.3.
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
}

func (o *OpenLocalMachineRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenLocalMachineOperation) *xxx_OpenLocalMachineOperation {
	if op == nil {
		op = &xxx_OpenLocalMachineOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.DesiredAccess = o.DesiredAccess
	return op
}

func (o *OpenLocalMachineRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenLocalMachineOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.DesiredAccess = op.DesiredAccess
}
func (o *OpenLocalMachineRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenLocalMachineRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenLocalMachineOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenLocalMachineResponse structure represents the OpenLocalMachine operation response
type OpenLocalMachineResponse struct {
	// phKey: A pointer to an RPC context handle for the root key, HKEY_LOCAL_MACHINE, as
	// specified in section 3.1.1.7. The handle is found in the handle table (HANDLETABLE).
	Key *Key `idl:"name:phKey" json:"key"`
	// Return: The OpenLocalMachine return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *OpenLocalMachineResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenLocalMachineOperation) *xxx_OpenLocalMachineOperation {
	if op == nil {
		op = &xxx_OpenLocalMachineOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.Return = o.Return
	return op
}

func (o *OpenLocalMachineResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenLocalMachineOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.Return = op.Return
}
func (o *OpenLocalMachineResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenLocalMachineResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenLocalMachineOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenPerformanceDataOperation structure represents the OpenPerformanceData operation
type xxx_OpenPerformanceDataOperation struct {
	ServerName    string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
	Key           *Key   `idl:"name:phKey" json:"key"`
	Return        uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenPerformanceDataOperation) OpNum() int { return 3 }

func (o *xxx_OpenPerformanceDataOperation) OpName() string { return "/winreg/v1/OpenPerformanceData" }

func (o *xxx_OpenPerformanceDataOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPerformanceDataOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME, names=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// samDesired {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPerformanceDataOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME, names=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// samDesired {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPerformanceDataOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPerformanceDataOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phKey {out} (1:{alias=PRPC_HKEY}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_OpenPerformanceDataOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phKey {out} (1:{alias=PRPC_HKEY,pointer=ref}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
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

// OpenPerformanceDataRequest structure represents the OpenPerformanceData operation request
type OpenPerformanceDataRequest struct {
	// ServerName: SHOULD be sent as NULL and MUST be ignored on receipt because the binding
	// to the server is already complete at this stage.
	ServerName string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	// samDesired: SHOULD be sent as 0 and MUST be ignored on receipt.
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
}

func (o *OpenPerformanceDataRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenPerformanceDataOperation) *xxx_OpenPerformanceDataOperation {
	if op == nil {
		op = &xxx_OpenPerformanceDataOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.DesiredAccess = o.DesiredAccess
	return op
}

func (o *OpenPerformanceDataRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenPerformanceDataOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.DesiredAccess = op.DesiredAccess
}
func (o *OpenPerformanceDataRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenPerformanceDataRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenPerformanceDataOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenPerformanceDataResponse structure represents the OpenPerformanceData operation response
type OpenPerformanceDataResponse struct {
	// phKey: A pointer to an RPC context handle for the root key, HKEY_PERFORMANCE_DATA,
	// as specified in section 3.1.1.7. The handle is found in the handle table (HANDLETABLE).
	Key *Key `idl:"name:phKey" json:"key"`
	// Return: The OpenPerformanceData return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *OpenPerformanceDataResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenPerformanceDataOperation) *xxx_OpenPerformanceDataOperation {
	if op == nil {
		op = &xxx_OpenPerformanceDataOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.Return = o.Return
	return op
}

func (o *OpenPerformanceDataResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenPerformanceDataOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.Return = op.Return
}
func (o *OpenPerformanceDataResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenPerformanceDataResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenPerformanceDataOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenUsersOperation structure represents the OpenUsers operation
type xxx_OpenUsersOperation struct {
	ServerName    string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
	Key           *Key   `idl:"name:phKey" json:"key"`
	Return        uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenUsersOperation) OpNum() int { return 4 }

func (o *xxx_OpenUsersOperation) OpName() string { return "/winreg/v1/OpenUsers" }

func (o *xxx_OpenUsersOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenUsersOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME, names=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// samDesired {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenUsersOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME, names=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// samDesired {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenUsersOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenUsersOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phKey {out} (1:{alias=PRPC_HKEY}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_OpenUsersOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phKey {out} (1:{alias=PRPC_HKEY,pointer=ref}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
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

// OpenUsersRequest structure represents the OpenUsers operation request
type OpenUsersRequest struct {
	// ServerName: SHOULD be sent as NULL and MUST be ignored on receipt because the binding
	// to the server is already complete at this stage.
	ServerName string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	// samDesired: The bit field that describes the wanted security access for the key.
	// It MUST be constructed from one or more of the values that are specified in section
	// 2.2.3.
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
}

func (o *OpenUsersRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenUsersOperation) *xxx_OpenUsersOperation {
	if op == nil {
		op = &xxx_OpenUsersOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.DesiredAccess = o.DesiredAccess
	return op
}

func (o *OpenUsersRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenUsersOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.DesiredAccess = op.DesiredAccess
}
func (o *OpenUsersRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenUsersRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenUsersOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenUsersResponse structure represents the OpenUsers operation response
type OpenUsersResponse struct {
	// phKey: A pointer to an RPC context handle for the root key, HKEY_USERS, as specified
	// in section 3.1.1.7. The handle is found in the handle table (HANDLETABLE).
	Key *Key `idl:"name:phKey" json:"key"`
	// Return: The OpenUsers return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *OpenUsersResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenUsersOperation) *xxx_OpenUsersOperation {
	if op == nil {
		op = &xxx_OpenUsersOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.Return = o.Return
	return op
}

func (o *OpenUsersResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenUsersOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.Return = op.Return
}
func (o *OpenUsersResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenUsersResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenUsersOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegCloseKeyOperation structure represents the BaseRegCloseKey operation
type xxx_BaseRegCloseKeyOperation struct {
	Key    *Key   `idl:"name:hKey" json:"key"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegCloseKeyOperation) OpNum() int { return 5 }

func (o *xxx_BaseRegCloseKeyOperation) OpName() string { return "/winreg/v1/BaseRegCloseKey" }

func (o *xxx_BaseRegCloseKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegCloseKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in, out} (1:{alias=PRPC_HKEY}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_BaseRegCloseKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in, out} (1:{alias=PRPC_HKEY,pointer=ref}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegCloseKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegCloseKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// hKey {in, out} (1:{alias=PRPC_HKEY}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_BaseRegCloseKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// hKey {in, out} (1:{alias=PRPC_HKEY,pointer=ref}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
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

// BaseRegCloseKeyRequest structure represents the BaseRegCloseKey operation request
type BaseRegCloseKeyRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegCreateKey, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
}

func (o *BaseRegCloseKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegCloseKeyOperation) *xxx_BaseRegCloseKeyOperation {
	if op == nil {
		op = &xxx_BaseRegCloseKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	return op
}

func (o *BaseRegCloseKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegCloseKeyOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
}
func (o *BaseRegCloseKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegCloseKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegCloseKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegCloseKeyResponse structure represents the BaseRegCloseKey operation response
type BaseRegCloseKeyResponse struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegCreateKey, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
	// Return: The BaseRegCloseKey return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegCloseKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegCloseKeyOperation) *xxx_BaseRegCloseKeyOperation {
	if op == nil {
		op = &xxx_BaseRegCloseKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.Return = o.Return
	return op
}

func (o *BaseRegCloseKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegCloseKeyOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.Return = op.Return
}
func (o *BaseRegCloseKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegCloseKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegCloseKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegCreateKeyOperation structure represents the BaseRegCreateKey operation
type xxx_BaseRegCreateKeyOperation struct {
	Key                *Key                `idl:"name:hKey" json:"key"`
	SubKey             *UnicodeString      `idl:"name:lpSubKey" json:"sub_key"`
	Class              *UnicodeString      `idl:"name:lpClass" json:"class"`
	Options            uint32              `idl:"name:dwOptions" json:"options"`
	DesiredAccess      uint32              `idl:"name:samDesired" json:"desired_access"`
	SecurityAttributes *SecurityAttributes `idl:"name:lpSecurityAttributes;pointer:unique" json:"security_attributes"`
	ResultKey          *Key                `idl:"name:phkResult" json:"result_key"`
	Disposition        uint32              `idl:"name:lpdwDisposition;pointer:unique" json:"disposition"`
	Return             uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegCreateKeyOperation) OpNum() int { return 6 }

func (o *xxx_BaseRegCreateKeyOperation) OpName() string { return "/winreg/v1/BaseRegCreateKey" }

func (o *xxx_BaseRegCreateKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegCreateKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpSubKey {in} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.SubKey != nil {
			if err := o.SubKey.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpClass {in} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.Class != nil {
			if err := o.Class.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	// samDesired {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	// lpSecurityAttributes {in} (1:{pointer=unique, alias=PRPC_SECURITY_ATTRIBUTES}*(1))(2:{alias=RPC_SECURITY_ATTRIBUTES}(struct))
	{
		if o.SecurityAttributes != nil {
			_ptr_lpSecurityAttributes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SecurityAttributes != nil {
					if err := o.SecurityAttributes.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&SecurityAttributes{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SecurityAttributes, _ptr_lpSecurityAttributes); err != nil {
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
	// lpdwDisposition {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_null_if(cond) attribute.
		_ptr_lpdwDisposition := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Disposition); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Disposition, _ptr_lpdwDisposition); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegCreateKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpSubKey {in} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.SubKey == nil {
			o.SubKey = &UnicodeString{}
		}
		if err := o.SubKey.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpClass {in} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.Class == nil {
			o.Class = &UnicodeString{}
		}
		if err := o.Class.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	// samDesired {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	// lpSecurityAttributes {in} (1:{pointer=unique, alias=PRPC_SECURITY_ATTRIBUTES}*(1))(2:{alias=RPC_SECURITY_ATTRIBUTES}(struct))
	{
		_ptr_lpSecurityAttributes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SecurityAttributes == nil {
				o.SecurityAttributes = &SecurityAttributes{}
			}
			if err := o.SecurityAttributes.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpSecurityAttributes := func(ptr interface{}) { o.SecurityAttributes = *ptr.(**SecurityAttributes) }
		if err := w.ReadPointer(&o.SecurityAttributes, _s_lpSecurityAttributes, _ptr_lpSecurityAttributes); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpdwDisposition {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpdwDisposition := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Disposition); err != nil {
				return err
			}
			return nil
		})
		_s_lpdwDisposition := func(ptr interface{}) { o.Disposition = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Disposition, _s_lpdwDisposition, _ptr_lpdwDisposition); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegCreateKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegCreateKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phkResult {out} (1:{alias=PRPC_HKEY}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.ResultKey != nil {
			if err := o.ResultKey.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpdwDisposition {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_null_if(cond) attribute.
		_ptr_lpdwDisposition := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Disposition); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Disposition, _ptr_lpdwDisposition); err != nil {
			return err
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

func (o *xxx_BaseRegCreateKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phkResult {out} (1:{alias=PRPC_HKEY,pointer=ref}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.ResultKey == nil {
			o.ResultKey = &Key{}
		}
		if err := o.ResultKey.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpdwDisposition {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpdwDisposition := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Disposition); err != nil {
				return err
			}
			return nil
		})
		_s_lpdwDisposition := func(ptr interface{}) { o.Disposition = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Disposition, _s_lpdwDisposition, _ptr_lpdwDisposition); err != nil {
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

// BaseRegCreateKeyRequest structure represents the BaseRegCreateKey operation request
type BaseRegCreateKeyRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: BaseRegCreateKey, OpenClassesRoot,
	// OpenCurrentUser, OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
	// lpSubKey: A pointer to an RRP_UNICODE_STRING structure that specifies the name of
	// the key (as specified in section 3.1.1.1) that this method opens or creates. The
	// name of the key specified is relative to the key specified by the hkey parameter.
	SubKey *UnicodeString `idl:"name:lpSubKey" json:"sub_key"`
	// lpClass: A pointer to an RRP_UNICODE_STRING structure that specifies the class of
	// the key (as specified in section 3.1.1.6).<16>
	Class *UnicodeString `idl:"name:lpClass" json:"class"`
	// dwOptions: Registry key options. MUST be the bitwise OR of one of the key types defined
	// in section 3.1.1.2, and any or none of the following options. The BaseRegCreateKey
	// method fails with ERROR_INVALID_PARAMETER if an unlisted value is specified.
	//
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	|                                       |                                                                                  |
	//	|                *VALUE*                |                                    *MEANING*                                     |
	//	|                                       |                                                                                  |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| REG_OPTION_BACKUP_RESTORE 0x00000004  | Indicates that the caller wishes to assert its backup and/or restore privileges. |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| REG_OPTION_OPEN_LINK 0x00000008       | Indicates that the caller wishes to open the targeted symlink source rather than |
	//	|                                       | the symlink target.                                                              |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	//	| REG_OPTION_DONT_VIRTUALIZE 0x00000010 | Indicates that the caller wishes to disable limited user access virtualization   |
	//	|                                       | for this operation.                                                              |
	//	+---------------------------------------+----------------------------------------------------------------------------------+
	Options uint32 `idl:"name:dwOptions" json:"options"`
	// samDesired: A bit field that describes the wanted security access for the handle
	// to the key that is being created or opened. It MUST be constructed from one or more
	// of the values that are specified in section 2.2.3.
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
	// lpSecurityAttributes: A pointer to an RPC_SECURITY_ATTRIBUTES structure for the new
	// subkey provided a new subkey is created.
	SecurityAttributes *SecurityAttributes `idl:"name:lpSecurityAttributes;pointer:unique" json:"security_attributes"`
	// lpdwDisposition: The disposition of the returned key indicated by phkResult. The
	// value of this parameter set by the client is ignored by the server. Its value MUST
	// be one of the following.
	//
	//	+------------------------------------+---------------------------------------------------------------+
	//	|                                    |                                                               |
	//	|               VALUE                |                            MEANING                            |
	//	|                                    |                                                               |
	//	+------------------------------------+---------------------------------------------------------------+
	//	+------------------------------------+---------------------------------------------------------------+
	//	| REG_CREATED_NEW_KEY 0x00000001     | The key did not exist and was created.                        |
	//	+------------------------------------+---------------------------------------------------------------+
	//	| REG_OPENED_EXISTING_KEY 0x00000002 | The key already existed and was opened without being changed. |
	//	+------------------------------------+---------------------------------------------------------------+
	Disposition uint32 `idl:"name:lpdwDisposition;pointer:unique" json:"disposition"`
}

func (o *BaseRegCreateKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegCreateKeyOperation) *xxx_BaseRegCreateKeyOperation {
	if op == nil {
		op = &xxx_BaseRegCreateKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.SubKey = o.SubKey
	op.Class = o.Class
	op.Options = o.Options
	op.DesiredAccess = o.DesiredAccess
	op.SecurityAttributes = o.SecurityAttributes
	op.Disposition = o.Disposition
	return op
}

func (o *BaseRegCreateKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegCreateKeyOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.SubKey = op.SubKey
	o.Class = op.Class
	o.Options = op.Options
	o.DesiredAccess = op.DesiredAccess
	o.SecurityAttributes = op.SecurityAttributes
	o.Disposition = op.Disposition
}
func (o *BaseRegCreateKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegCreateKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegCreateKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegCreateKeyResponse structure represents the BaseRegCreateKey operation response
type BaseRegCreateKeyResponse struct {
	// phkResult: A pointer to a variable that receives a handle to the opened or created
	// key.
	ResultKey *Key `idl:"name:phkResult" json:"result_key"`
	// lpdwDisposition: The disposition of the returned key indicated by phkResult. The
	// value of this parameter set by the client is ignored by the server. Its value MUST
	// be one of the following.
	//
	//	+------------------------------------+---------------------------------------------------------------+
	//	|                                    |                                                               |
	//	|               VALUE                |                            MEANING                            |
	//	|                                    |                                                               |
	//	+------------------------------------+---------------------------------------------------------------+
	//	+------------------------------------+---------------------------------------------------------------+
	//	| REG_CREATED_NEW_KEY 0x00000001     | The key did not exist and was created.                        |
	//	+------------------------------------+---------------------------------------------------------------+
	//	| REG_OPENED_EXISTING_KEY 0x00000002 | The key already existed and was opened without being changed. |
	//	+------------------------------------+---------------------------------------------------------------+
	Disposition uint32 `idl:"name:lpdwDisposition;pointer:unique" json:"disposition"`
	// Return: The BaseRegCreateKey return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegCreateKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegCreateKeyOperation) *xxx_BaseRegCreateKeyOperation {
	if op == nil {
		op = &xxx_BaseRegCreateKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.ResultKey = o.ResultKey
	op.Disposition = o.Disposition
	op.Return = o.Return
	return op
}

func (o *BaseRegCreateKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegCreateKeyOperation) {
	if o == nil {
		return
	}
	o.ResultKey = op.ResultKey
	o.Disposition = op.Disposition
	o.Return = op.Return
}
func (o *BaseRegCreateKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegCreateKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegCreateKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegDeleteKeyOperation structure represents the BaseRegDeleteKey operation
type xxx_BaseRegDeleteKeyOperation struct {
	Key    *Key           `idl:"name:hKey" json:"key"`
	SubKey *UnicodeString `idl:"name:lpSubKey" json:"sub_key"`
	Return uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegDeleteKeyOperation) OpNum() int { return 7 }

func (o *xxx_BaseRegDeleteKeyOperation) OpName() string { return "/winreg/v1/BaseRegDeleteKey" }

func (o *xxx_BaseRegDeleteKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegDeleteKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpSubKey {in} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.SubKey != nil {
			if err := o.SubKey.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegDeleteKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpSubKey {in} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.SubKey == nil {
			o.SubKey = &UnicodeString{}
		}
		if err := o.SubKey.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegDeleteKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegDeleteKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BaseRegDeleteKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BaseRegDeleteKeyRequest structure represents the BaseRegDeleteKey operation request
type BaseRegDeleteKeyRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegCreateKey, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
	// lpSubKey: A pointer to an RRP_UNICODE_STRING structure that MUST contain the name
	// of the key (as specified in section 3.1.1) to delete.
	SubKey *UnicodeString `idl:"name:lpSubKey" json:"sub_key"`
}

func (o *BaseRegDeleteKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegDeleteKeyOperation) *xxx_BaseRegDeleteKeyOperation {
	if op == nil {
		op = &xxx_BaseRegDeleteKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.SubKey = o.SubKey
	return op
}

func (o *BaseRegDeleteKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegDeleteKeyOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.SubKey = op.SubKey
}
func (o *BaseRegDeleteKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegDeleteKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegDeleteKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegDeleteKeyResponse structure represents the BaseRegDeleteKey operation response
type BaseRegDeleteKeyResponse struct {
	// Return: The BaseRegDeleteKey return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegDeleteKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegDeleteKeyOperation) *xxx_BaseRegDeleteKeyOperation {
	if op == nil {
		op = &xxx_BaseRegDeleteKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *BaseRegDeleteKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegDeleteKeyOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BaseRegDeleteKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegDeleteKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegDeleteKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegDeleteValueOperation structure represents the BaseRegDeleteValue operation
type xxx_BaseRegDeleteValueOperation struct {
	Key       *Key           `idl:"name:hKey" json:"key"`
	ValueName *UnicodeString `idl:"name:lpValueName" json:"value_name"`
	Return    uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegDeleteValueOperation) OpNum() int { return 8 }

func (o *xxx_BaseRegDeleteValueOperation) OpName() string { return "/winreg/v1/BaseRegDeleteValue" }

func (o *xxx_BaseRegDeleteValueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegDeleteValueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpValueName {in} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.ValueName != nil {
			if err := o.ValueName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegDeleteValueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpValueName {in} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.ValueName == nil {
			o.ValueName = &UnicodeString{}
		}
		if err := o.ValueName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegDeleteValueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegDeleteValueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BaseRegDeleteValueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BaseRegDeleteValueRequest structure represents the BaseRegDeleteValue operation request
type BaseRegDeleteValueRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegCreateKey, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
	// lpValueName: A pointer to an RRP_UNICODE_STRING structure that MUST contain the name
	// of the value (as specified in section 3.1.1) to remove. If the client sets the lpValueName
	// parameter to NULL, the server SHOULD<18> fail this method and return ERROR_INVALID_PARAMETER.
	ValueName *UnicodeString `idl:"name:lpValueName" json:"value_name"`
}

func (o *BaseRegDeleteValueRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegDeleteValueOperation) *xxx_BaseRegDeleteValueOperation {
	if op == nil {
		op = &xxx_BaseRegDeleteValueOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.ValueName = o.ValueName
	return op
}

func (o *BaseRegDeleteValueRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegDeleteValueOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.ValueName = op.ValueName
}
func (o *BaseRegDeleteValueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegDeleteValueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegDeleteValueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegDeleteValueResponse structure represents the BaseRegDeleteValue operation response
type BaseRegDeleteValueResponse struct {
	// Return: The BaseRegDeleteValue return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegDeleteValueResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegDeleteValueOperation) *xxx_BaseRegDeleteValueOperation {
	if op == nil {
		op = &xxx_BaseRegDeleteValueOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *BaseRegDeleteValueResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegDeleteValueOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BaseRegDeleteValueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegDeleteValueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegDeleteValueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegEnumKeyOperation structure represents the BaseRegEnumKey operation
type xxx_BaseRegEnumKeyOperation struct {
	Key           *Key                `idl:"name:hKey" json:"key"`
	Index         uint32              `idl:"name:dwIndex" json:"index"`
	NameIn        *UnicodeString      `idl:"name:lpNameIn" json:"name_in"`
	NameOut       *UnicodeString      `idl:"name:lpNameOut" json:"name_out"`
	ClassIn       *UnicodeString      `idl:"name:lpClassIn;pointer:unique" json:"class_in"`
	ClassOut      *dtyp.UnicodeString `idl:"name:lplpClassOut" json:"class_out"`
	LastWriteTime *dtyp.Filetime      `idl:"name:lpftLastWriteTime;pointer:unique" json:"last_write_time"`
	Return        uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegEnumKeyOperation) OpNum() int { return 9 }

func (o *xxx_BaseRegEnumKeyOperation) OpName() string { return "/winreg/v1/BaseRegEnumKey" }

func (o *xxx_BaseRegEnumKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegEnumKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	// lpNameIn {in} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.NameIn != nil {
			if err := o.NameIn.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpClassIn {in} (1:{pointer=unique, alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.ClassIn != nil {
			_ptr_lpClassIn := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ClassIn != nil {
					if err := o.ClassIn.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClassIn, _ptr_lpClassIn); err != nil {
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
	// lpftLastWriteTime {in, out} (1:{pointer=unique, alias=PFILETIME}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.LastWriteTime != nil {
			_ptr_lpftLastWriteTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.LastWriteTime != nil {
					if err := o.LastWriteTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.LastWriteTime, _ptr_lpftLastWriteTime); err != nil {
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

func (o *xxx_BaseRegEnumKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	// lpNameIn {in} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.NameIn == nil {
			o.NameIn = &UnicodeString{}
		}
		if err := o.NameIn.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpClassIn {in} (1:{pointer=unique, alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		_ptr_lpClassIn := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ClassIn == nil {
				o.ClassIn = &UnicodeString{}
			}
			if err := o.ClassIn.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpClassIn := func(ptr interface{}) { o.ClassIn = *ptr.(**UnicodeString) }
		if err := w.ReadPointer(&o.ClassIn, _s_lpClassIn, _ptr_lpClassIn); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpftLastWriteTime {in, out} (1:{pointer=unique, alias=PFILETIME}*(1))(2:{alias=FILETIME}(struct))
	{
		_ptr_lpftLastWriteTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.LastWriteTime == nil {
				o.LastWriteTime = &dtyp.Filetime{}
			}
			if err := o.LastWriteTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpftLastWriteTime := func(ptr interface{}) { o.LastWriteTime = *ptr.(**dtyp.Filetime) }
		if err := w.ReadPointer(&o.LastWriteTime, _s_lpftLastWriteTime, _ptr_lpftLastWriteTime); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegEnumKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegEnumKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpNameOut {out} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.NameOut != nil {
			if err := o.NameOut.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lplpClassOut {out} (1:{pointer=ref}*(2))(2:{alias=PRPC_UNICODE_STRING}*(1))(3:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.ClassOut != nil {
			_ptr_lplpClassOut := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.ClassOut != nil {
					if err := o.ClassOut.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.ClassOut, _ptr_lplpClassOut); err != nil {
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
	// lpftLastWriteTime {in, out} (1:{pointer=unique, alias=PFILETIME}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.LastWriteTime != nil {
			_ptr_lpftLastWriteTime := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.LastWriteTime != nil {
					if err := o.LastWriteTime.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.LastWriteTime, _ptr_lpftLastWriteTime); err != nil {
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

func (o *xxx_BaseRegEnumKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpNameOut {out} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.NameOut == nil {
			o.NameOut = &UnicodeString{}
		}
		if err := o.NameOut.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lplpClassOut {out} (1:{pointer=ref}*(2))(2:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(3:{alias=RPC_UNICODE_STRING}(struct))
	{
		_ptr_lplpClassOut := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.ClassOut == nil {
				o.ClassOut = &dtyp.UnicodeString{}
			}
			if err := o.ClassOut.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lplpClassOut := func(ptr interface{}) { o.ClassOut = *ptr.(**dtyp.UnicodeString) }
		if err := w.ReadPointer(&o.ClassOut, _s_lplpClassOut, _ptr_lplpClassOut); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpftLastWriteTime {in, out} (1:{pointer=unique, alias=PFILETIME}*(1))(2:{alias=FILETIME}(struct))
	{
		_ptr_lpftLastWriteTime := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.LastWriteTime == nil {
				o.LastWriteTime = &dtyp.Filetime{}
			}
			if err := o.LastWriteTime.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_lpftLastWriteTime := func(ptr interface{}) { o.LastWriteTime = *ptr.(**dtyp.Filetime) }
		if err := w.ReadPointer(&o.LastWriteTime, _s_lpftLastWriteTime, _ptr_lpftLastWriteTime); err != nil {
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

// BaseRegEnumKeyRequest structure represents the BaseRegEnumKey operation request
type BaseRegEnumKeyRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegCreateKey, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
	// dwIndex: The index of the subkey to retrieve, as specified in section 3.1.1.1.
	Index uint32 `idl:"name:dwIndex" json:"index"`
	// lpNameIn: A pointer to an RRP_UNICODE_STRING structure (section 2.2.4) that contains
	// the key name to be retrieved, as specified in section 3.1.1.1. This is used by the
	// server to determine the maximum length for the output name parameter and to allocate
	// space accordingly. The content is ignored, and only the maximum length is significant.
	// The Length field MUST be set to 0.
	NameIn *UnicodeString `idl:"name:lpNameIn" json:"name_in"`
	// lpClassIn: A pointer to an RRP_UNICODE_STRING structure (section 2.2.4) that contains
	// the class to be retrieved, as specified in section 3.1.1.6. This is used by the server
	// to determine the maximum length for the output class parameter and to allocate space
	// accordingly. The content is ignored.
	ClassIn *UnicodeString `idl:"name:lpClassIn;pointer:unique" json:"class_in"`
	// lpftLastWriteTime: MUST be the time when the value was last written (set or created).
	LastWriteTime *dtyp.Filetime `idl:"name:lpftLastWriteTime;pointer:unique" json:"last_write_time"`
}

func (o *BaseRegEnumKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegEnumKeyOperation) *xxx_BaseRegEnumKeyOperation {
	if op == nil {
		op = &xxx_BaseRegEnumKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.Index = o.Index
	op.NameIn = o.NameIn
	op.ClassIn = o.ClassIn
	op.LastWriteTime = o.LastWriteTime
	return op
}

func (o *BaseRegEnumKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegEnumKeyOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.Index = op.Index
	o.NameIn = op.NameIn
	o.ClassIn = op.ClassIn
	o.LastWriteTime = op.LastWriteTime
}
func (o *BaseRegEnumKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegEnumKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegEnumKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegEnumKeyResponse structure represents the BaseRegEnumKey operation response
type BaseRegEnumKeyResponse struct {
	// lpNameOut: A pointer to an RRP_UNICODE_STRING structure that receives the name of
	// the retrieved key, as specified in section 3.1.1.1. All fields MUST be set to 0.
	NameOut *UnicodeString `idl:"name:lpNameOut" json:"name_out"`
	// lplpClassOut: A pointer to a PRPC_UNICODE_STRING structure ([MS-DTYP] section 2.3.10),
	// that receives the class of the retrieved key, as specified in section 3.1.1.6. This
	// parameter is optional.
	ClassOut *dtyp.UnicodeString `idl:"name:lplpClassOut" json:"class_out"`
	// lpftLastWriteTime: MUST be the time when the value was last written (set or created).
	LastWriteTime *dtyp.Filetime `idl:"name:lpftLastWriteTime;pointer:unique" json:"last_write_time"`
	// Return: The BaseRegEnumKey return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegEnumKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegEnumKeyOperation) *xxx_BaseRegEnumKeyOperation {
	if op == nil {
		op = &xxx_BaseRegEnumKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.NameOut = o.NameOut
	op.ClassOut = o.ClassOut
	op.LastWriteTime = o.LastWriteTime
	op.Return = o.Return
	return op
}

func (o *BaseRegEnumKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegEnumKeyOperation) {
	if o == nil {
		return
	}
	o.NameOut = op.NameOut
	o.ClassOut = op.ClassOut
	o.LastWriteTime = op.LastWriteTime
	o.Return = op.Return
}
func (o *BaseRegEnumKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegEnumKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegEnumKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegEnumValueOperation structure represents the BaseRegEnumValue operation
type xxx_BaseRegEnumValueOperation struct {
	Key          *Key                `idl:"name:hKey" json:"key"`
	Index        uint32              `idl:"name:dwIndex" json:"index"`
	ValueNameIn  *UnicodeString      `idl:"name:lpValueNameIn" json:"value_name_in"`
	ValueNameOut *dtyp.UnicodeString `idl:"name:lpValueNameOut" json:"value_name_out"`
	Type         uint32              `idl:"name:lpType;pointer:unique" json:"type"`
	Data         []byte              `idl:"name:lpData;size_is:(lpcbData);length_is:(lpcbLen);pointer:unique" json:"data"`
	DataLength   uint32              `idl:"name:lpcbData;pointer:unique" json:"data_length"`
	Length       uint32              `idl:"name:lpcbLen;pointer:unique" json:"length"`
	Return       uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegEnumValueOperation) OpNum() int { return 10 }

func (o *xxx_BaseRegEnumValueOperation) OpName() string { return "/winreg/v1/BaseRegEnumValue" }

func (o *xxx_BaseRegEnumValueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Data != nil && o.DataLength == 0 {
		o.DataLength = uint32(len(o.Data))
	}
	if o.Data != nil && o.Length == 0 {
		o.Length = uint32(len(o.Data))
	}
	if len(o.Data) > int(67108864) {
		return fmt.Errorf("Data is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegEnumValueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// dwIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Index); err != nil {
			return err
		}
	}
	// lpValueNameIn {in} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.ValueNameIn != nil {
			if err := o.ValueNameIn.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpType {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_null_if(cond) attribute.
		_ptr_lpType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Type); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Type, _ptr_lpType); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpData {in, out} (1:{pointer=unique, range=(0,67108864), alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=lpcbData,length_is=lpcbLen](uchar))
	{
		if o.Data != nil || o.DataLength > 0 {
			_ptr_lpData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.DataLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				dimLength1 := uint64(o.Length)
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
				for i1 := range o.Data {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Data[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Data, _ptr_lpData); err != nil {
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
	// lpcbData {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_null_if(cond) attribute.
		_ptr_lpcbData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.DataLength); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DataLength, _ptr_lpcbData); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpcbLen {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_null_if(cond) attribute.
		_ptr_lpcbLen := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Length); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Length, _ptr_lpcbLen); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegEnumValueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// dwIndex {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Index); err != nil {
			return err
		}
	}
	// lpValueNameIn {in} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.ValueNameIn == nil {
			o.ValueNameIn = &UnicodeString{}
		}
		if err := o.ValueNameIn.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpType {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Type); err != nil {
				return err
			}
			return nil
		})
		_s_lpType := func(ptr interface{}) { o.Type = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Type, _s_lpType, _ptr_lpType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpData {in, out} (1:{pointer=unique, range=(0,67108864), alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=lpcbData,length_is=lpcbLen](uchar))
	{
		_ptr_lpData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
			}
			o.Data = make([]byte, sizeInfo[0])
			for i1 := range o.Data {
				i1 := i1
				if err := w.ReadData(&o.Data[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_lpData := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Data, _s_lpData, _ptr_lpData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpcbData {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpcbData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.DataLength); err != nil {
				return err
			}
			return nil
		})
		_s_lpcbData := func(ptr interface{}) { o.DataLength = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.DataLength, _s_lpcbData, _ptr_lpcbData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpcbLen {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpcbLen := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Length); err != nil {
				return err
			}
			return nil
		})
		_s_lpcbLen := func(ptr interface{}) { o.Length = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Length, _s_lpcbLen, _ptr_lpcbLen); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegEnumValueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Data != nil && o.DataLength == 0 {
		o.DataLength = uint32(len(o.Data))
	}
	if o.Data != nil && o.Length == 0 {
		o.Length = uint32(len(o.Data))
	}
	if len(o.Data) > int(67108864) {
		return fmt.Errorf("Data is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegEnumValueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpValueNameOut {out} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.ValueNameOut != nil {
			if err := o.ValueNameOut.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpType {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_null_if(cond) attribute.
		_ptr_lpType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Type); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Type, _ptr_lpType); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpData {in, out} (1:{pointer=unique, range=(0,67108864), alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=lpcbData,length_is=lpcbLen](uchar))
	{
		if o.Data != nil || o.DataLength > 0 {
			_ptr_lpData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.DataLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				dimLength1 := uint64(o.Length)
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
				for i1 := range o.Data {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Data[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Data, _ptr_lpData); err != nil {
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
	// lpcbData {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_null_if(cond) attribute.
		_ptr_lpcbData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.DataLength); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DataLength, _ptr_lpcbData); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpcbLen {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_null_if(cond) attribute.
		_ptr_lpcbLen := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Length); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Length, _ptr_lpcbLen); err != nil {
			return err
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

func (o *xxx_BaseRegEnumValueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpValueNameOut {out} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.ValueNameOut == nil {
			o.ValueNameOut = &dtyp.UnicodeString{}
		}
		if err := o.ValueNameOut.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpType {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Type); err != nil {
				return err
			}
			return nil
		})
		_s_lpType := func(ptr interface{}) { o.Type = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Type, _s_lpType, _ptr_lpType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpData {in, out} (1:{pointer=unique, range=(0,67108864), alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=lpcbData,length_is=lpcbLen](uchar))
	{
		_ptr_lpData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
			}
			o.Data = make([]byte, sizeInfo[0])
			for i1 := range o.Data {
				i1 := i1
				if err := w.ReadData(&o.Data[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_lpData := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Data, _s_lpData, _ptr_lpData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpcbData {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpcbData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.DataLength); err != nil {
				return err
			}
			return nil
		})
		_s_lpcbData := func(ptr interface{}) { o.DataLength = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.DataLength, _s_lpcbData, _ptr_lpcbData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpcbLen {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpcbLen := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Length); err != nil {
				return err
			}
			return nil
		})
		_s_lpcbLen := func(ptr interface{}) { o.Length = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Length, _s_lpcbLen, _ptr_lpcbLen); err != nil {
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

// BaseRegEnumValueRequest structure represents the BaseRegEnumValue operation request
type BaseRegEnumValueRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegCreateKey, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
	// dwIndex: MUST be the index of the value to be retrieved, as specified in section
	// 3.1.1.5.
	Index uint32 `idl:"name:dwIndex" json:"index"`
	// lpValueNameIn: A pointer to an RRP_UNICODE_STRING structure that contains the value
	// name to be retrieved, as specified in section 3.1.1.5. This is used by the server
	// to determine the maximum length for the output name parameter and to allocate space
	// accordingly. The content is ignored, and only the maximum length is significant.
	ValueNameIn *UnicodeString `idl:"name:lpValueNameIn" json:"value_name_in"`
	// lpType: An optional pointer to a buffer that receives the REG_VALUE_TYPE of the value
	// (as specified in section 3.1.1.5), or it MUST be NULL.
	Type uint32 `idl:"name:lpType;pointer:unique" json:"type"`
	// lpData: An optional pointer to a buffer that receives the data of the value entry.
	Data []byte `idl:"name:lpData;size_is:(lpcbData);length_is:(lpcbLen);pointer:unique" json:"data"`
	// lpcbData: A pointer to a variable that MUST contain the size of the buffer that is
	// pointed to by lpData. MUST NOT be NULL if lpData is present.
	DataLength uint32 `idl:"name:lpcbData;pointer:unique" json:"data_length"`
	// lpcbLen: MUST specify the number of bytes to transmit to the client.
	Length uint32 `idl:"name:lpcbLen;pointer:unique" json:"length"`
}

func (o *BaseRegEnumValueRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegEnumValueOperation) *xxx_BaseRegEnumValueOperation {
	if op == nil {
		op = &xxx_BaseRegEnumValueOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.Index = o.Index
	op.ValueNameIn = o.ValueNameIn
	op.Type = o.Type
	op.Data = o.Data
	op.DataLength = o.DataLength
	op.Length = o.Length
	return op
}

func (o *BaseRegEnumValueRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegEnumValueOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.Index = op.Index
	o.ValueNameIn = op.ValueNameIn
	o.Type = op.Type
	o.Data = op.Data
	o.DataLength = op.DataLength
	o.Length = op.Length
}
func (o *BaseRegEnumValueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegEnumValueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegEnumValueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegEnumValueResponse structure represents the BaseRegEnumValue operation response
type BaseRegEnumValueResponse struct {
	// lpValueNameOut: A pointer to an RPC_UNICODE_STRING structure that receives the retrieved
	// value name, as specified in section 3.1.1.5.
	ValueNameOut *dtyp.UnicodeString `idl:"name:lpValueNameOut" json:"value_name_out"`
	// lpType: An optional pointer to a buffer that receives the REG_VALUE_TYPE of the value
	// (as specified in section 3.1.1.5), or it MUST be NULL.
	Type uint32 `idl:"name:lpType;pointer:unique" json:"type"`
	// lpData: An optional pointer to a buffer that receives the data of the value entry.
	Data []byte `idl:"name:lpData;size_is:(lpcbData);length_is:(lpcbLen);pointer:unique" json:"data"`
	// lpcbData: A pointer to a variable that MUST contain the size of the buffer that is
	// pointed to by lpData. MUST NOT be NULL if lpData is present.
	DataLength uint32 `idl:"name:lpcbData;pointer:unique" json:"data_length"`
	// lpcbLen: MUST specify the number of bytes to transmit to the client.
	Length uint32 `idl:"name:lpcbLen;pointer:unique" json:"length"`
	// Return: The BaseRegEnumValue return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegEnumValueResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegEnumValueOperation) *xxx_BaseRegEnumValueOperation {
	if op == nil {
		op = &xxx_BaseRegEnumValueOperation{}
	}
	if o == nil {
		return op
	}
	op.ValueNameOut = o.ValueNameOut
	op.Type = o.Type
	op.Data = o.Data
	op.DataLength = o.DataLength
	op.Length = o.Length
	op.Return = o.Return
	return op
}

func (o *BaseRegEnumValueResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegEnumValueOperation) {
	if o == nil {
		return
	}
	o.ValueNameOut = op.ValueNameOut
	o.Type = op.Type
	o.Data = op.Data
	o.DataLength = op.DataLength
	o.Length = op.Length
	o.Return = op.Return
}
func (o *BaseRegEnumValueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegEnumValueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegEnumValueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegFlushKeyOperation structure represents the BaseRegFlushKey operation
type xxx_BaseRegFlushKeyOperation struct {
	Key    *Key   `idl:"name:hKey" json:"key"`
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegFlushKeyOperation) OpNum() int { return 11 }

func (o *xxx_BaseRegFlushKeyOperation) OpName() string { return "/winreg/v1/BaseRegFlushKey" }

func (o *xxx_BaseRegFlushKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegFlushKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_BaseRegFlushKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegFlushKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegFlushKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BaseRegFlushKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BaseRegFlushKeyRequest structure represents the BaseRegFlushKey operation request
type BaseRegFlushKeyRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegCreateKey, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
}

func (o *BaseRegFlushKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegFlushKeyOperation) *xxx_BaseRegFlushKeyOperation {
	if op == nil {
		op = &xxx_BaseRegFlushKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	return op
}

func (o *BaseRegFlushKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegFlushKeyOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
}
func (o *BaseRegFlushKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegFlushKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegFlushKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegFlushKeyResponse structure represents the BaseRegFlushKey operation response
type BaseRegFlushKeyResponse struct {
	// Return: The BaseRegFlushKey return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegFlushKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegFlushKeyOperation) *xxx_BaseRegFlushKeyOperation {
	if op == nil {
		op = &xxx_BaseRegFlushKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *BaseRegFlushKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegFlushKeyOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BaseRegFlushKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegFlushKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegFlushKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegGetKeySecurityOperation structure represents the BaseRegGetKeySecurity operation
type xxx_BaseRegGetKeySecurityOperation struct {
	Key                   *Key                `idl:"name:hKey" json:"key"`
	SecurityInformation   uint32              `idl:"name:SecurityInformation" json:"security_information"`
	SecurityDescriptorIn  *SecurityDescriptor `idl:"name:pRpcSecurityDescriptorIn" json:"security_descriptor_in"`
	SecurityDescriptorOut *SecurityDescriptor `idl:"name:pRpcSecurityDescriptorOut" json:"security_descriptor_out"`
	Return                uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegGetKeySecurityOperation) OpNum() int { return 12 }

func (o *xxx_BaseRegGetKeySecurityOperation) OpName() string {
	return "/winreg/v1/BaseRegGetKeySecurity"
}

func (o *xxx_BaseRegGetKeySecurityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegGetKeySecurityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// SecurityInformation {in} (1:{alias=SECURITY_INFORMATION, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SecurityInformation); err != nil {
			return err
		}
	}
	// pRpcSecurityDescriptorIn {in} (1:{alias=PRPC_SECURITY_DESCRIPTOR}*(1))(2:{alias=RPC_SECURITY_DESCRIPTOR}(struct))
	{
		if o.SecurityDescriptorIn != nil {
			if err := o.SecurityDescriptorIn.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SecurityDescriptor{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegGetKeySecurityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// SecurityInformation {in} (1:{alias=SECURITY_INFORMATION, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SecurityInformation); err != nil {
			return err
		}
	}
	// pRpcSecurityDescriptorIn {in} (1:{alias=PRPC_SECURITY_DESCRIPTOR,pointer=ref}*(1))(2:{alias=RPC_SECURITY_DESCRIPTOR}(struct))
	{
		if o.SecurityDescriptorIn == nil {
			o.SecurityDescriptorIn = &SecurityDescriptor{}
		}
		if err := o.SecurityDescriptorIn.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegGetKeySecurityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegGetKeySecurityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pRpcSecurityDescriptorOut {out} (1:{alias=PRPC_SECURITY_DESCRIPTOR}*(1))(2:{alias=RPC_SECURITY_DESCRIPTOR}(struct))
	{
		if o.SecurityDescriptorOut != nil {
			if err := o.SecurityDescriptorOut.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SecurityDescriptor{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_BaseRegGetKeySecurityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pRpcSecurityDescriptorOut {out} (1:{alias=PRPC_SECURITY_DESCRIPTOR,pointer=ref}*(1))(2:{alias=RPC_SECURITY_DESCRIPTOR}(struct))
	{
		if o.SecurityDescriptorOut == nil {
			o.SecurityDescriptorOut = &SecurityDescriptor{}
		}
		if err := o.SecurityDescriptorOut.UnmarshalNDR(ctx, w); err != nil {
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

// BaseRegGetKeySecurityRequest structure represents the BaseRegGetKeySecurity operation request
type BaseRegGetKeySecurityRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegCreateKey, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
	// SecurityInformation: The information that is needed to determine the type of security
	// that is returned in pRpcSecurityDescriptorOut. See SECURITY_INFORMATION (includes
	// a list of possible values).
	SecurityInformation uint32 `idl:"name:SecurityInformation" json:"security_information"`
	// pRpcSecurityDescriptorIn: A pointer to a buffer containing a security descriptor.
	// The client MUST provide a pointer to an RPC_SECURITY_DESCRIPTOR with arbitrary contents.
	// The server uses the size of this security descriptor to validate the client has the
	// correct amount of memory allocated for the RPC_SECURITY_DESCRIPTOR pointed to by
	// the pRpcSecurityDescriptorOut parameter
	SecurityDescriptorIn *SecurityDescriptor `idl:"name:pRpcSecurityDescriptorIn" json:"security_descriptor_in"`
}

func (o *BaseRegGetKeySecurityRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegGetKeySecurityOperation) *xxx_BaseRegGetKeySecurityOperation {
	if op == nil {
		op = &xxx_BaseRegGetKeySecurityOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.SecurityInformation = o.SecurityInformation
	op.SecurityDescriptorIn = o.SecurityDescriptorIn
	return op
}

func (o *BaseRegGetKeySecurityRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegGetKeySecurityOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.SecurityInformation = op.SecurityInformation
	o.SecurityDescriptorIn = op.SecurityDescriptorIn
}
func (o *BaseRegGetKeySecurityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegGetKeySecurityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegGetKeySecurityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegGetKeySecurityResponse structure represents the BaseRegGetKeySecurity operation response
type BaseRegGetKeySecurityResponse struct {
	// pRpcSecurityDescriptorOut: A pointer to a buffer to which the requested security
	// descriptor MUST be written.
	SecurityDescriptorOut *SecurityDescriptor `idl:"name:pRpcSecurityDescriptorOut" json:"security_descriptor_out"`
	// Return: The BaseRegGetKeySecurity return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegGetKeySecurityResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegGetKeySecurityOperation) *xxx_BaseRegGetKeySecurityOperation {
	if op == nil {
		op = &xxx_BaseRegGetKeySecurityOperation{}
	}
	if o == nil {
		return op
	}
	op.SecurityDescriptorOut = o.SecurityDescriptorOut
	op.Return = o.Return
	return op
}

func (o *BaseRegGetKeySecurityResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegGetKeySecurityOperation) {
	if o == nil {
		return
	}
	o.SecurityDescriptorOut = op.SecurityDescriptorOut
	o.Return = op.Return
}
func (o *BaseRegGetKeySecurityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegGetKeySecurityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegGetKeySecurityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegLoadKeyOperation structure represents the BaseRegLoadKey operation
type xxx_BaseRegLoadKeyOperation struct {
	Key    *Key           `idl:"name:hKey" json:"key"`
	SubKey *UnicodeString `idl:"name:lpSubKey" json:"sub_key"`
	File   *UnicodeString `idl:"name:lpFile" json:"file"`
	Return uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegLoadKeyOperation) OpNum() int { return 13 }

func (o *xxx_BaseRegLoadKeyOperation) OpName() string { return "/winreg/v1/BaseRegLoadKey" }

func (o *xxx_BaseRegLoadKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegLoadKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpSubKey {in} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.SubKey != nil {
			if err := o.SubKey.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpFile {in} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.File != nil {
			if err := o.File.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegLoadKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpSubKey {in} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.SubKey == nil {
			o.SubKey = &UnicodeString{}
		}
		if err := o.SubKey.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpFile {in} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.File == nil {
			o.File = &UnicodeString{}
		}
		if err := o.File.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegLoadKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegLoadKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BaseRegLoadKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BaseRegLoadKeyRequest structure represents the BaseRegLoadKey operation request
type BaseRegLoadKeyRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenUsers and OpenLocalMachine.
	//
	// Note  The other open methods in this protocol cannot be used to obtain the hKey parameter
	// because the server checks that the key specified by lpSubKey is a descendent of the
	// HKEY_LOCAL_MACHINE or HKEY_USERS root keys.
	Key *Key `idl:"name:hKey" json:"key"`
	// lpSubKey: A pointer to an RRP_UNICODE_STRING structure that specifies the name of
	// the key (as specified in section 3.1.1.1) that MUST be created under hKey.
	SubKey *UnicodeString `idl:"name:lpSubKey" json:"sub_key"`
	// lpFile: A pointer to a null-terminated RRP_UNICODE_STRING structure that contains
	// the name of an existing registry file in which the specified key and subkeys are
	// to be saved. The format of the file name is implementation-specific. It is assumed
	// that this file was created with the BaseRegSaveKey method. If it does not exist,
	// the server creates a file with the specified name.
	File *UnicodeString `idl:"name:lpFile" json:"file"`
}

func (o *BaseRegLoadKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegLoadKeyOperation) *xxx_BaseRegLoadKeyOperation {
	if op == nil {
		op = &xxx_BaseRegLoadKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.SubKey = o.SubKey
	op.File = o.File
	return op
}

func (o *BaseRegLoadKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegLoadKeyOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.SubKey = op.SubKey
	o.File = op.File
}
func (o *BaseRegLoadKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegLoadKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegLoadKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegLoadKeyResponse structure represents the BaseRegLoadKey operation response
type BaseRegLoadKeyResponse struct {
	// Return: The BaseRegLoadKey return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegLoadKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegLoadKeyOperation) *xxx_BaseRegLoadKeyOperation {
	if op == nil {
		op = &xxx_BaseRegLoadKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *BaseRegLoadKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegLoadKeyOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BaseRegLoadKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegLoadKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegLoadKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegOpenKeyOperation structure represents the BaseRegOpenKey operation
type xxx_BaseRegOpenKeyOperation struct {
	Key           *Key           `idl:"name:hKey" json:"key"`
	SubKey        *UnicodeString `idl:"name:lpSubKey" json:"sub_key"`
	Options       uint32         `idl:"name:dwOptions" json:"options"`
	DesiredAccess uint32         `idl:"name:samDesired" json:"desired_access"`
	ResultKey     *Key           `idl:"name:phkResult" json:"result_key"`
	Return        uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegOpenKeyOperation) OpNum() int { return 15 }

func (o *xxx_BaseRegOpenKeyOperation) OpName() string { return "/winreg/v1/BaseRegOpenKey" }

func (o *xxx_BaseRegOpenKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegOpenKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpSubKey {in} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.SubKey != nil {
			if err := o.SubKey.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Options); err != nil {
			return err
		}
	}
	// samDesired {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegOpenKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpSubKey {in} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.SubKey == nil {
			o.SubKey = &UnicodeString{}
		}
		if err := o.SubKey.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwOptions {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Options); err != nil {
			return err
		}
	}
	// samDesired {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegOpenKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegOpenKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phkResult {out} (1:{alias=PRPC_HKEY}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.ResultKey != nil {
			if err := o.ResultKey.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_BaseRegOpenKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phkResult {out} (1:{alias=PRPC_HKEY,pointer=ref}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.ResultKey == nil {
			o.ResultKey = &Key{}
		}
		if err := o.ResultKey.UnmarshalNDR(ctx, w); err != nil {
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

// BaseRegOpenKeyRequest structure represents the BaseRegOpenKey operation request
type BaseRegOpenKeyRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegOpenKey, OpenCurrentConfig,
	// OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
	// lpSubKey: A pointer to an RRP_UNICODE_STRING structure that MUST contain the name
	// of a key to open. This parameter is always relative to the key that is specified
	// by the hKey parameter and is a pointer to a null-terminated string that contains
	// the name of the subkey to open, as specified in section 3.1.1. This key MUST be an
	// existing subkey of the key that is identified by the hKey parameter.
	SubKey *UnicodeString `idl:"name:lpSubKey" json:"sub_key"`
	// dwOptions: Registry key options. MUST be the bitwise OR of any or none of the following
	// values.
	//
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	|                                      |                                                                                  |
	//	|               *VALUE*                |                                    *MEANING*                                     |
	//	|                                      |                                                                                  |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| REG_OPTION_BACKUP_RESTORE 0x00000004 | Indicates that the caller requests to assert its backup and/or restore           |
	//	|                                      | privileges.                                                                      |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	//	| REG_OPTION_OPEN_LINK 0x00000008      | Indicates that the caller requests to open the targeted symlink source rather    |
	//	|                                      | than the symlink target.                                                         |
	//	+--------------------------------------+----------------------------------------------------------------------------------+
	Options uint32 `idl:"name:dwOptions" json:"options"`
	// samDesired: A bit field that describes the requested security access for the handle
	// to the key that is being opened. It MUST be constructed from one or more of the values
	// that are specified in section 2.2.3.
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
}

func (o *BaseRegOpenKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegOpenKeyOperation) *xxx_BaseRegOpenKeyOperation {
	if op == nil {
		op = &xxx_BaseRegOpenKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.SubKey = o.SubKey
	op.Options = o.Options
	op.DesiredAccess = o.DesiredAccess
	return op
}

func (o *BaseRegOpenKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegOpenKeyOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.SubKey = op.SubKey
	o.Options = op.Options
	o.DesiredAccess = op.DesiredAccess
}
func (o *BaseRegOpenKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegOpenKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegOpenKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegOpenKeyResponse structure represents the BaseRegOpenKey operation response
type BaseRegOpenKeyResponse struct {
	// phkResult: A pointer to the handle of the open key. The server MUST return a NULL
	// for phkResult in case of failure.
	ResultKey *Key `idl:"name:phkResult" json:"result_key"`
	// Return: The BaseRegOpenKey return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegOpenKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegOpenKeyOperation) *xxx_BaseRegOpenKeyOperation {
	if op == nil {
		op = &xxx_BaseRegOpenKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.ResultKey = o.ResultKey
	op.Return = o.Return
	return op
}

func (o *BaseRegOpenKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegOpenKeyOperation) {
	if o == nil {
		return
	}
	o.ResultKey = op.ResultKey
	o.Return = op.Return
}
func (o *BaseRegOpenKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegOpenKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegOpenKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegQueryInfoKeyOperation structure represents the BaseRegQueryInfoKey operation
type xxx_BaseRegQueryInfoKeyOperation struct {
	Key                      *Key                `idl:"name:hKey" json:"key"`
	ClassIn                  *UnicodeString      `idl:"name:lpClassIn" json:"class_in"`
	ClassOut                 *dtyp.UnicodeString `idl:"name:lpClassOut" json:"class_out"`
	SubKeysCount             uint32              `idl:"name:lpcSubKeys" json:"sub_keys_count"`
	MaxSubKeyLength          uint32              `idl:"name:lpcbMaxSubKeyLen" json:"max_sub_key_length"`
	MaxClassLength           uint32              `idl:"name:lpcbMaxClassLen" json:"max_class_length"`
	ValuesCount              uint32              `idl:"name:lpcValues" json:"values_count"`
	MaxValueNameLength       uint32              `idl:"name:lpcbMaxValueNameLen" json:"max_value_name_length"`
	MaxValueLength           uint32              `idl:"name:lpcbMaxValueLen" json:"max_value_length"`
	SecurityDescriptorLength uint32              `idl:"name:lpcbSecurityDescriptor" json:"security_descriptor_length"`
	LastWriteTime            *dtyp.Filetime      `idl:"name:lpftLastWriteTime" json:"last_write_time"`
	Return                   uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegQueryInfoKeyOperation) OpNum() int { return 16 }

func (o *xxx_BaseRegQueryInfoKeyOperation) OpName() string { return "/winreg/v1/BaseRegQueryInfoKey" }

func (o *xxx_BaseRegQueryInfoKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegQueryInfoKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpClassIn {in} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.ClassIn != nil {
			if err := o.ClassIn.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegQueryInfoKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpClassIn {in} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.ClassIn == nil {
			o.ClassIn = &UnicodeString{}
		}
		if err := o.ClassIn.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegQueryInfoKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegQueryInfoKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpClassOut {out} (1:{alias=PRPC_UNICODE_STRING}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.ClassOut != nil {
			if err := o.ClassOut.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpcSubKeys {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SubKeysCount); err != nil {
			return err
		}
	}
	// lpcbMaxSubKeyLen {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MaxSubKeyLength); err != nil {
			return err
		}
	}
	// lpcbMaxClassLen {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MaxClassLength); err != nil {
			return err
		}
	}
	// lpcValues {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ValuesCount); err != nil {
			return err
		}
	}
	// lpcbMaxValueNameLen {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MaxValueNameLength); err != nil {
			return err
		}
	}
	// lpcbMaxValueLen {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.MaxValueLength); err != nil {
			return err
		}
	}
	// lpcbSecurityDescriptor {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.SecurityDescriptorLength); err != nil {
			return err
		}
	}
	// lpftLastWriteTime {out} (1:{alias=PFILETIME}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.LastWriteTime != nil {
			if err := o.LastWriteTime.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.Filetime{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_BaseRegQueryInfoKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpClassOut {out} (1:{alias=PRPC_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RPC_UNICODE_STRING}(struct))
	{
		if o.ClassOut == nil {
			o.ClassOut = &dtyp.UnicodeString{}
		}
		if err := o.ClassOut.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpcSubKeys {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SubKeysCount); err != nil {
			return err
		}
	}
	// lpcbMaxSubKeyLen {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MaxSubKeyLength); err != nil {
			return err
		}
	}
	// lpcbMaxClassLen {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MaxClassLength); err != nil {
			return err
		}
	}
	// lpcValues {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ValuesCount); err != nil {
			return err
		}
	}
	// lpcbMaxValueNameLen {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MaxValueNameLength); err != nil {
			return err
		}
	}
	// lpcbMaxValueLen {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.MaxValueLength); err != nil {
			return err
		}
	}
	// lpcbSecurityDescriptor {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SecurityDescriptorLength); err != nil {
			return err
		}
	}
	// lpftLastWriteTime {out} (1:{alias=PFILETIME,pointer=ref}*(1))(2:{alias=FILETIME}(struct))
	{
		if o.LastWriteTime == nil {
			o.LastWriteTime = &dtyp.Filetime{}
		}
		if err := o.LastWriteTime.UnmarshalNDR(ctx, w); err != nil {
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

// BaseRegQueryInfoKeyRequest structure represents the BaseRegQueryInfoKey operation request
type BaseRegQueryInfoKeyRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegCreateKey, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
	// lpClassIn: A pointer to an RRP_UNICODE_STRING structure that contains the class of
	// the key to be retrieved, as specified in section 3.1.1.6. This string is optional;
	// it is saved but is not used by the registry.
	ClassIn *UnicodeString `idl:"name:lpClassIn" json:"class_in"`
}

func (o *BaseRegQueryInfoKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegQueryInfoKeyOperation) *xxx_BaseRegQueryInfoKeyOperation {
	if op == nil {
		op = &xxx_BaseRegQueryInfoKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.ClassIn = o.ClassIn
	return op
}

func (o *BaseRegQueryInfoKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegQueryInfoKeyOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.ClassIn = op.ClassIn
}
func (o *BaseRegQueryInfoKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegQueryInfoKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegQueryInfoKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegQueryInfoKeyResponse structure represents the BaseRegQueryInfoKey operation response
type BaseRegQueryInfoKeyResponse struct {
	// lpClassOut: A pointer to an RPC_UNICODE_STRING structure that receives the class
	// of this key, as specified in section 3.1.1.6.
	ClassOut *dtyp.UnicodeString `idl:"name:lpClassOut" json:"class_out"`
	// lpcSubKeys: A pointer to a DWORD that MUST receive the count of the subkeys of the
	// specified key.
	SubKeysCount uint32 `idl:"name:lpcSubKeys" json:"sub_keys_count"`
	// lpcbMaxSubKeyLen: A pointer to a DWORD that receives the size of the key's subkey
	// with the longest name, or a greater size, as the number of TCHAR elements.
	//
	// TCHAR elements are defined as follows.
	MaxSubKeyLength uint32 `idl:"name:lpcbMaxSubKeyLen" json:"max_sub_key_length"`
	// lpcbMaxClassLen: A pointer to a DWORD that receives the size of the longest string
	// that specifies a subkey class, or a greater size, in Unicode characters.
	MaxClassLength uint32 `idl:"name:lpcbMaxClassLen" json:"max_class_length"`
	// lpcValues: A pointer to a DWORD that receives the number of values that are associated
	// with the key.
	ValuesCount uint32 `idl:"name:lpcValues" json:"values_count"`
	// lpcbMaxValueNameLen: A pointer to a DWORD that receives the size of the key's longest
	// value name, or a greater size, as the number of TCHAR elements.
	MaxValueNameLength uint32 `idl:"name:lpcbMaxValueNameLen" json:"max_value_name_length"`
	// lpcbMaxValueLen: A pointer to a DWORD that receives the size in bytes of the longest
	// data component, or a greater size, in the key's values.
	MaxValueLength uint32 `idl:"name:lpcbMaxValueLen" json:"max_value_length"`
	// lpcbSecurityDescriptor: A pointer to a DWORD that receives the size in bytes of the
	// key's SECURITY_DESCRIPTOR.
	SecurityDescriptorLength uint32 `idl:"name:lpcbSecurityDescriptor" json:"security_descriptor_length"`
	// lpftLastWriteTime: A pointer to a FILETIME structure that receives the last write
	// time.
	LastWriteTime *dtyp.Filetime `idl:"name:lpftLastWriteTime" json:"last_write_time"`
	// Return: The BaseRegQueryInfoKey return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegQueryInfoKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegQueryInfoKeyOperation) *xxx_BaseRegQueryInfoKeyOperation {
	if op == nil {
		op = &xxx_BaseRegQueryInfoKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.ClassOut = o.ClassOut
	op.SubKeysCount = o.SubKeysCount
	op.MaxSubKeyLength = o.MaxSubKeyLength
	op.MaxClassLength = o.MaxClassLength
	op.ValuesCount = o.ValuesCount
	op.MaxValueNameLength = o.MaxValueNameLength
	op.MaxValueLength = o.MaxValueLength
	op.SecurityDescriptorLength = o.SecurityDescriptorLength
	op.LastWriteTime = o.LastWriteTime
	op.Return = o.Return
	return op
}

func (o *BaseRegQueryInfoKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegQueryInfoKeyOperation) {
	if o == nil {
		return
	}
	o.ClassOut = op.ClassOut
	o.SubKeysCount = op.SubKeysCount
	o.MaxSubKeyLength = op.MaxSubKeyLength
	o.MaxClassLength = op.MaxClassLength
	o.ValuesCount = op.ValuesCount
	o.MaxValueNameLength = op.MaxValueNameLength
	o.MaxValueLength = op.MaxValueLength
	o.SecurityDescriptorLength = op.SecurityDescriptorLength
	o.LastWriteTime = op.LastWriteTime
	o.Return = op.Return
}
func (o *BaseRegQueryInfoKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegQueryInfoKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegQueryInfoKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegQueryValueOperation structure represents the BaseRegQueryValue operation
type xxx_BaseRegQueryValueOperation struct {
	Key        *Key           `idl:"name:hKey" json:"key"`
	ValueName  *UnicodeString `idl:"name:lpValueName" json:"value_name"`
	Type       uint32         `idl:"name:lpType;pointer:unique" json:"type"`
	Data       []byte         `idl:"name:lpData;size_is:(lpcbData);length_is:(lpcbLen);pointer:unique" json:"data"`
	DataLength uint32         `idl:"name:lpcbData;pointer:unique" json:"data_length"`
	Length     uint32         `idl:"name:lpcbLen;pointer:unique" json:"length"`
	Return     uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegQueryValueOperation) OpNum() int { return 17 }

func (o *xxx_BaseRegQueryValueOperation) OpName() string { return "/winreg/v1/BaseRegQueryValue" }

func (o *xxx_BaseRegQueryValueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Data != nil && o.DataLength == 0 {
		o.DataLength = uint32(len(o.Data))
	}
	if o.Data != nil && o.Length == 0 {
		o.Length = uint32(len(o.Data))
	}
	if len(o.Data) > int(67108864) {
		return fmt.Errorf("Data is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegQueryValueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpValueName {in} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.ValueName != nil {
			if err := o.ValueName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpType {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_null_if(cond) attribute.
		_ptr_lpType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Type); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Type, _ptr_lpType); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpData {in, out} (1:{pointer=unique, range=(0,67108864), alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=lpcbData,length_is=lpcbLen](uchar))
	{
		if o.Data != nil || o.DataLength > 0 {
			_ptr_lpData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.DataLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				dimLength1 := uint64(o.Length)
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
				for i1 := range o.Data {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Data[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Data, _ptr_lpData); err != nil {
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
	// lpcbData {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_null_if(cond) attribute.
		_ptr_lpcbData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.DataLength); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DataLength, _ptr_lpcbData); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpcbLen {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_null_if(cond) attribute.
		_ptr_lpcbLen := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Length); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Length, _ptr_lpcbLen); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegQueryValueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpValueName {in} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.ValueName == nil {
			o.ValueName = &UnicodeString{}
		}
		if err := o.ValueName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpType {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Type); err != nil {
				return err
			}
			return nil
		})
		_s_lpType := func(ptr interface{}) { o.Type = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Type, _s_lpType, _ptr_lpType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpData {in, out} (1:{pointer=unique, range=(0,67108864), alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=lpcbData,length_is=lpcbLen](uchar))
	{
		_ptr_lpData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
			}
			o.Data = make([]byte, sizeInfo[0])
			for i1 := range o.Data {
				i1 := i1
				if err := w.ReadData(&o.Data[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_lpData := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Data, _s_lpData, _ptr_lpData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpcbData {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpcbData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.DataLength); err != nil {
				return err
			}
			return nil
		})
		_s_lpcbData := func(ptr interface{}) { o.DataLength = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.DataLength, _s_lpcbData, _ptr_lpcbData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpcbLen {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpcbLen := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Length); err != nil {
				return err
			}
			return nil
		})
		_s_lpcbLen := func(ptr interface{}) { o.Length = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Length, _s_lpcbLen, _ptr_lpcbLen); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegQueryValueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Data != nil && o.DataLength == 0 {
		o.DataLength = uint32(len(o.Data))
	}
	if o.Data != nil && o.Length == 0 {
		o.Length = uint32(len(o.Data))
	}
	if len(o.Data) > int(67108864) {
		return fmt.Errorf("Data is out of range")
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegQueryValueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpType {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_null_if(cond) attribute.
		_ptr_lpType := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Type); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Type, _ptr_lpType); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpData {in, out} (1:{pointer=unique, range=(0,67108864), alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=lpcbData,length_is=lpcbLen](uchar))
	{
		if o.Data != nil || o.DataLength > 0 {
			_ptr_lpData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.DataLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				dimLength1 := uint64(o.Length)
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
				for i1 := range o.Data {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.Data[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Data, _ptr_lpData); err != nil {
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
	// lpcbData {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_null_if(cond) attribute.
		_ptr_lpcbData := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.DataLength); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.DataLength, _ptr_lpcbData); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpcbLen {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		// XXX pointer to primitive type, default behavior is to write non-null pointer.
		// if this behavior is not desired, use goext_null_if(cond) attribute.
		_ptr_lpcbLen := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Length); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Length, _ptr_lpcbLen); err != nil {
			return err
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

func (o *xxx_BaseRegQueryValueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpType {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpType := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Type); err != nil {
				return err
			}
			return nil
		})
		_s_lpType := func(ptr interface{}) { o.Type = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Type, _s_lpType, _ptr_lpType); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpData {in, out} (1:{pointer=unique, range=(0,67108864), alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=lpcbData,length_is=lpcbLen](uchar))
	{
		_ptr_lpData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
			}
			o.Data = make([]byte, sizeInfo[0])
			for i1 := range o.Data {
				i1 := i1
				if err := w.ReadData(&o.Data[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_lpData := func(ptr interface{}) { o.Data = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Data, _s_lpData, _ptr_lpData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpcbData {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpcbData := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.DataLength); err != nil {
				return err
			}
			return nil
		})
		_s_lpcbData := func(ptr interface{}) { o.DataLength = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.DataLength, _s_lpcbData, _ptr_lpcbData); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpcbLen {in, out} (1:{pointer=unique, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_lpcbLen := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Length); err != nil {
				return err
			}
			return nil
		})
		_s_lpcbLen := func(ptr interface{}) { o.Length = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Length, _s_lpcbLen, _ptr_lpcbLen); err != nil {
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

// BaseRegQueryValueRequest structure represents the BaseRegQueryValue operation request
type BaseRegQueryValueRequest struct {
	// hKey: On input, a handle to a key that MUST have been opened previously by using
	// one of the open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegCreateKey, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
	// lpValueName: On input, the client sets lpValueName to a pointer to an RRP_UNICODE_STRING
	// structure that MUST contain the name of the value, as specified in section 3.1.1.
	// If the client sets lpValueName to NULL, the server MUST fail this method and return
	// ERROR_INVALID_PARAMETER.
	ValueName *UnicodeString `idl:"name:lpValueName" json:"value_name"`
	// lpType: On input, the client sets lpType to a pointer to a variable to receive the
	// type code of a value entry. On output, the server MUST set this parameter to NULL
	// if the value specified by the lpValueName parameter is not found. If the client sets
	// lpType to NULL, the server MUST fail this method and return ERROR_INVALID_PARAMETER.
	Type uint32 `idl:"name:lpType;pointer:unique" json:"type"`
	// lpData: On input, the client sets lpData to a pointer to a buffer to receive the
	// data of the value entry.
	Data []byte `idl:"name:lpData;size_is:(lpcbData);length_is:(lpcbLen);pointer:unique" json:"data"`
	// lpcbData: A pointer to a variable that, on input, contains the size in bytes of the
	// buffer that is pointed to by the lpData parameter. On output, the variable receives
	// the number of bytes that are returned in lpData. This length variable MUST be set
	// to 0 by the server if the client provides NULL for the lpData parameter.
	DataLength uint32 `idl:"name:lpcbData;pointer:unique" json:"data_length"`
	// lpcbLen: A pointer to a variable that contains the number of bytes to transmit to
	// the client. On input, the client MUST allocate the memory for this parameter and
	// the pointer value of this parameter MUST not be NULL. On output, the server MUST
	// set this parameter to the size (in bytes) of the buffer pointed to by the lpData
	// parameter. If the client sets lpcbLen to NULL, the server MUST fail this method and
	// return ERROR_INVALID_PARAMETER.
	Length uint32 `idl:"name:lpcbLen;pointer:unique" json:"length"`
}

func (o *BaseRegQueryValueRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegQueryValueOperation) *xxx_BaseRegQueryValueOperation {
	if op == nil {
		op = &xxx_BaseRegQueryValueOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.ValueName = o.ValueName
	op.Type = o.Type
	op.Data = o.Data
	op.DataLength = o.DataLength
	op.Length = o.Length
	return op
}

func (o *BaseRegQueryValueRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegQueryValueOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.ValueName = op.ValueName
	o.Type = op.Type
	o.Data = op.Data
	o.DataLength = op.DataLength
	o.Length = op.Length
}
func (o *BaseRegQueryValueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegQueryValueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegQueryValueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegQueryValueResponse structure represents the BaseRegQueryValue operation response
type BaseRegQueryValueResponse struct {
	// lpType: On input, the client sets lpType to a pointer to a variable to receive the
	// type code of a value entry. On output, the server MUST set this parameter to NULL
	// if the value specified by the lpValueName parameter is not found. If the client sets
	// lpType to NULL, the server MUST fail this method and return ERROR_INVALID_PARAMETER.
	Type uint32 `idl:"name:lpType;pointer:unique" json:"type"`
	// lpData: On input, the client sets lpData to a pointer to a buffer to receive the
	// data of the value entry.
	Data []byte `idl:"name:lpData;size_is:(lpcbData);length_is:(lpcbLen);pointer:unique" json:"data"`
	// lpcbData: A pointer to a variable that, on input, contains the size in bytes of the
	// buffer that is pointed to by the lpData parameter. On output, the variable receives
	// the number of bytes that are returned in lpData. This length variable MUST be set
	// to 0 by the server if the client provides NULL for the lpData parameter.
	DataLength uint32 `idl:"name:lpcbData;pointer:unique" json:"data_length"`
	// lpcbLen: A pointer to a variable that contains the number of bytes to transmit to
	// the client. On input, the client MUST allocate the memory for this parameter and
	// the pointer value of this parameter MUST not be NULL. On output, the server MUST
	// set this parameter to the size (in bytes) of the buffer pointed to by the lpData
	// parameter. If the client sets lpcbLen to NULL, the server MUST fail this method and
	// return ERROR_INVALID_PARAMETER.
	Length uint32 `idl:"name:lpcbLen;pointer:unique" json:"length"`
	// Return: The BaseRegQueryValue return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegQueryValueResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegQueryValueOperation) *xxx_BaseRegQueryValueOperation {
	if op == nil {
		op = &xxx_BaseRegQueryValueOperation{}
	}
	if o == nil {
		return op
	}
	op.Type = o.Type
	op.Data = o.Data
	op.DataLength = o.DataLength
	op.Length = o.Length
	op.Return = o.Return
	return op
}

func (o *BaseRegQueryValueResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegQueryValueOperation) {
	if o == nil {
		return
	}
	o.Type = op.Type
	o.Data = op.Data
	o.DataLength = op.DataLength
	o.Length = op.Length
	o.Return = op.Return
}
func (o *BaseRegQueryValueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegQueryValueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegQueryValueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegReplaceKeyOperation structure represents the BaseRegReplaceKey operation
type xxx_BaseRegReplaceKeyOperation struct {
	Key     *Key           `idl:"name:hKey" json:"key"`
	SubKey  *UnicodeString `idl:"name:lpSubKey" json:"sub_key"`
	NewFile *UnicodeString `idl:"name:lpNewFile" json:"new_file"`
	OldFile *UnicodeString `idl:"name:lpOldFile" json:"old_file"`
	Return  uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegReplaceKeyOperation) OpNum() int { return 18 }

func (o *xxx_BaseRegReplaceKeyOperation) OpName() string { return "/winreg/v1/BaseRegReplaceKey" }

func (o *xxx_BaseRegReplaceKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegReplaceKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpSubKey {in} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.SubKey != nil {
			if err := o.SubKey.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpNewFile {in} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.NewFile != nil {
			if err := o.NewFile.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpOldFile {in} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.OldFile != nil {
			if err := o.OldFile.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegReplaceKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpSubKey {in} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.SubKey == nil {
			o.SubKey = &UnicodeString{}
		}
		if err := o.SubKey.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpNewFile {in} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.NewFile == nil {
			o.NewFile = &UnicodeString{}
		}
		if err := o.NewFile.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpOldFile {in} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.OldFile == nil {
			o.OldFile = &UnicodeString{}
		}
		if err := o.OldFile.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegReplaceKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegReplaceKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BaseRegReplaceKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BaseRegReplaceKeyRequest structure represents the BaseRegReplaceKey operation request
type BaseRegReplaceKeyRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegCreateKey, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
	// lpSubKey: A pointer to an RRP_UNICODE_STRING structure that MUST either contain the
	// name of the key whose subkeys and values are replaced by this method (as specified
	// in section 3.1.1), or be NULL.
	SubKey *UnicodeString `idl:"name:lpSubKey" json:"sub_key"`
	// lpNewFile: A pointer to an RRP_UNICODE_STRING structure that MUST contain a registry
	// file name with the registration information, as specified in section 3.1.1. The format
	// of the file name is implementation-specific, but both lpNewFile and lpOldFile SHOULD<22>
	// be in the same format.
	NewFile *UnicodeString `idl:"name:lpNewFile" json:"new_file"`
	// lpOldFile: A pointer to an RRP_UNICODE_STRING structure that MUST contain the registry
	// file name that receives a backup copy of the replaced registry information. The format
	// of the file name is implementation-specific, but is in the same format as lpNewFile.
	OldFile *UnicodeString `idl:"name:lpOldFile" json:"old_file"`
}

func (o *BaseRegReplaceKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegReplaceKeyOperation) *xxx_BaseRegReplaceKeyOperation {
	if op == nil {
		op = &xxx_BaseRegReplaceKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.SubKey = o.SubKey
	op.NewFile = o.NewFile
	op.OldFile = o.OldFile
	return op
}

func (o *BaseRegReplaceKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegReplaceKeyOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.SubKey = op.SubKey
	o.NewFile = op.NewFile
	o.OldFile = op.OldFile
}
func (o *BaseRegReplaceKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegReplaceKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegReplaceKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegReplaceKeyResponse structure represents the BaseRegReplaceKey operation response
type BaseRegReplaceKeyResponse struct {
	// Return: The BaseRegReplaceKey return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegReplaceKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegReplaceKeyOperation) *xxx_BaseRegReplaceKeyOperation {
	if op == nil {
		op = &xxx_BaseRegReplaceKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *BaseRegReplaceKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegReplaceKeyOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BaseRegReplaceKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegReplaceKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegReplaceKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegRestoreKeyOperation structure represents the BaseRegRestoreKey operation
type xxx_BaseRegRestoreKeyOperation struct {
	Key    *Key           `idl:"name:hKey" json:"key"`
	File   *UnicodeString `idl:"name:lpFile" json:"file"`
	Flags  uint32         `idl:"name:Flags" json:"flags"`
	Return uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegRestoreKeyOperation) OpNum() int { return 19 }

func (o *xxx_BaseRegRestoreKeyOperation) OpName() string { return "/winreg/v1/BaseRegRestoreKey" }

func (o *xxx_BaseRegRestoreKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegRestoreKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpFile {in} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.File != nil {
			if err := o.File.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
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

func (o *xxx_BaseRegRestoreKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpFile {in} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.File == nil {
			o.File = &UnicodeString{}
		}
		if err := o.File.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
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

func (o *xxx_BaseRegRestoreKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegRestoreKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BaseRegRestoreKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BaseRegRestoreKeyRequest structure represents the BaseRegRestoreKey operation request
type BaseRegRestoreKeyRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegCreateKey, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
	// lpFile: A pointer to an RRP_UNICODE_STRING structure that contains an existing registry
	// file name. The format of the file name SHOULD<24> be implementation-specific.
	File *UnicodeString `idl:"name:lpFile" json:"file"`
	// Flags: An optional flag argument that is the bitwise OR of any of the following options.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|                                    |                                                                                  |
	//	|               VALUE                |                                     MEANING                                      |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| REG_WHOLE_HIVE_VOLATILE 0x00000001 | If set, registry keys created in the Data Store from the file indicated by       |
	//	|                                    | lpFile MUST be VOLATILE.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| REG_REFRESH_HIVE 0x00000002        | If set, the location of the subtree that the hKey parameter points to is         |
	//	|                                    | restored to its state immediately following the last flush. The subtree MUST     |
	//	|                                    | NOT be lazy flushed (by calling BaseRegRestoreKey with REG_NO_LAZY_FLUSH         |
	//	|                                    | specified as the value of this parameter); the caller MUST be a member of the    |
	//	|                                    | Administrators Group; and the handle the hKey parameter refers to MUST point to  |
	//	|                                    | the root of the subtree.                                                         |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| REG_NO_LAZY_FLUSH 0x00000004       | If set, the key or subtree that is specified by the hKey parameter does not      |
	//	|                                    | automatically flush at regular intervals of time. The server MUST set the        |
	//	|                                    | property KEYNOPERIODICFLUSH equal to TRUE for the key specified by the hKey      |
	//	|                                    | parameter and all subkeys (see section 3.1.1.3).                                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| REG_FORCE_RESTORE 0x00000008       | If set, the restore operation is executed even if open handles exist at (or      |
	//	|                                    | beneath) the location in the registry hierarchy to which the hKey parameter      |
	//	|                                    | points.<25>                                                                      |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
}

func (o *BaseRegRestoreKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegRestoreKeyOperation) *xxx_BaseRegRestoreKeyOperation {
	if op == nil {
		op = &xxx_BaseRegRestoreKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.File = o.File
	op.Flags = o.Flags
	return op
}

func (o *BaseRegRestoreKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegRestoreKeyOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.File = op.File
	o.Flags = op.Flags
}
func (o *BaseRegRestoreKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegRestoreKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegRestoreKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegRestoreKeyResponse structure represents the BaseRegRestoreKey operation response
type BaseRegRestoreKeyResponse struct {
	// Return: The BaseRegRestoreKey return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegRestoreKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegRestoreKeyOperation) *xxx_BaseRegRestoreKeyOperation {
	if op == nil {
		op = &xxx_BaseRegRestoreKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *BaseRegRestoreKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegRestoreKeyOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BaseRegRestoreKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegRestoreKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegRestoreKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegSaveKeyOperation structure represents the BaseRegSaveKey operation
type xxx_BaseRegSaveKeyOperation struct {
	Key                *Key                `idl:"name:hKey" json:"key"`
	File               *UnicodeString      `idl:"name:lpFile" json:"file"`
	SecurityAttributes *SecurityAttributes `idl:"name:pSecurityAttributes;pointer:unique" json:"security_attributes"`
	Return             uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegSaveKeyOperation) OpNum() int { return 20 }

func (o *xxx_BaseRegSaveKeyOperation) OpName() string { return "/winreg/v1/BaseRegSaveKey" }

func (o *xxx_BaseRegSaveKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegSaveKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpFile {in} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.File != nil {
			if err := o.File.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pSecurityAttributes {in} (1:{pointer=unique, alias=PRPC_SECURITY_ATTRIBUTES}*(1))(2:{alias=RPC_SECURITY_ATTRIBUTES}(struct))
	{
		if o.SecurityAttributes != nil {
			_ptr_pSecurityAttributes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SecurityAttributes != nil {
					if err := o.SecurityAttributes.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&SecurityAttributes{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SecurityAttributes, _ptr_pSecurityAttributes); err != nil {
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

func (o *xxx_BaseRegSaveKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpFile {in} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.File == nil {
			o.File = &UnicodeString{}
		}
		if err := o.File.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pSecurityAttributes {in} (1:{pointer=unique, alias=PRPC_SECURITY_ATTRIBUTES}*(1))(2:{alias=RPC_SECURITY_ATTRIBUTES}(struct))
	{
		_ptr_pSecurityAttributes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SecurityAttributes == nil {
				o.SecurityAttributes = &SecurityAttributes{}
			}
			if err := o.SecurityAttributes.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pSecurityAttributes := func(ptr interface{}) { o.SecurityAttributes = *ptr.(**SecurityAttributes) }
		if err := w.ReadPointer(&o.SecurityAttributes, _s_pSecurityAttributes, _ptr_pSecurityAttributes); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegSaveKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegSaveKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BaseRegSaveKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BaseRegSaveKeyRequest structure represents the BaseRegSaveKey operation request
type BaseRegSaveKeyRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegCreateKey, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
	// lpFile: A pointer to an RRP_UNICODE_STRING structure that contains the name of the
	// registry file in which the specified key and subkeys are to be saved. The format
	// of the file name SHOULD<26> be implementation-specific.
	File *UnicodeString `idl:"name:lpFile" json:"file"`
	// pSecurityAttributes: A pointer to an RPC_SECURITY_ATTRIBUTES structure.
	SecurityAttributes *SecurityAttributes `idl:"name:pSecurityAttributes;pointer:unique" json:"security_attributes"`
}

func (o *BaseRegSaveKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegSaveKeyOperation) *xxx_BaseRegSaveKeyOperation {
	if op == nil {
		op = &xxx_BaseRegSaveKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.File = o.File
	op.SecurityAttributes = o.SecurityAttributes
	return op
}

func (o *BaseRegSaveKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegSaveKeyOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.File = op.File
	o.SecurityAttributes = op.SecurityAttributes
}
func (o *BaseRegSaveKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegSaveKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegSaveKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegSaveKeyResponse structure represents the BaseRegSaveKey operation response
type BaseRegSaveKeyResponse struct {
	// Return: The BaseRegSaveKey return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegSaveKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegSaveKeyOperation) *xxx_BaseRegSaveKeyOperation {
	if op == nil {
		op = &xxx_BaseRegSaveKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *BaseRegSaveKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegSaveKeyOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BaseRegSaveKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegSaveKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegSaveKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegSetKeySecurityOperation structure represents the BaseRegSetKeySecurity operation
type xxx_BaseRegSetKeySecurityOperation struct {
	Key                 *Key                `idl:"name:hKey" json:"key"`
	SecurityInformation uint32              `idl:"name:SecurityInformation" json:"security_information"`
	SecurityDescriptor  *SecurityDescriptor `idl:"name:pRpcSecurityDescriptor" json:"security_descriptor"`
	Return              uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegSetKeySecurityOperation) OpNum() int { return 21 }

func (o *xxx_BaseRegSetKeySecurityOperation) OpName() string {
	return "/winreg/v1/BaseRegSetKeySecurity"
}

func (o *xxx_BaseRegSetKeySecurityOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegSetKeySecurityOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// SecurityInformation {in} (1:{alias=SECURITY_INFORMATION, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.SecurityInformation); err != nil {
			return err
		}
	}
	// pRpcSecurityDescriptor {in} (1:{alias=PRPC_SECURITY_DESCRIPTOR}*(1))(2:{alias=RPC_SECURITY_DESCRIPTOR}(struct))
	{
		if o.SecurityDescriptor != nil {
			if err := o.SecurityDescriptor.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SecurityDescriptor{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegSetKeySecurityOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// SecurityInformation {in} (1:{alias=SECURITY_INFORMATION, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.SecurityInformation); err != nil {
			return err
		}
	}
	// pRpcSecurityDescriptor {in} (1:{alias=PRPC_SECURITY_DESCRIPTOR,pointer=ref}*(1))(2:{alias=RPC_SECURITY_DESCRIPTOR}(struct))
	{
		if o.SecurityDescriptor == nil {
			o.SecurityDescriptor = &SecurityDescriptor{}
		}
		if err := o.SecurityDescriptor.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegSetKeySecurityOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegSetKeySecurityOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BaseRegSetKeySecurityOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BaseRegSetKeySecurityRequest structure represents the BaseRegSetKeySecurity operation request
type BaseRegSetKeySecurityRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegCreateKey, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
	// SecurityInformation: The SECURITY_INFORMATION that specifies the content of the pRpcSecurityDescriptor
	// parameter.
	SecurityInformation uint32 `idl:"name:SecurityInformation" json:"security_information"`
	// pRpcSecurityDescriptor: A pointer to the RPC_SECURITY_DESCRIPTOR to set for the supplied
	// key.
	SecurityDescriptor *SecurityDescriptor `idl:"name:pRpcSecurityDescriptor" json:"security_descriptor"`
}

func (o *BaseRegSetKeySecurityRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegSetKeySecurityOperation) *xxx_BaseRegSetKeySecurityOperation {
	if op == nil {
		op = &xxx_BaseRegSetKeySecurityOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.SecurityInformation = o.SecurityInformation
	op.SecurityDescriptor = o.SecurityDescriptor
	return op
}

func (o *BaseRegSetKeySecurityRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegSetKeySecurityOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.SecurityInformation = op.SecurityInformation
	o.SecurityDescriptor = op.SecurityDescriptor
}
func (o *BaseRegSetKeySecurityRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegSetKeySecurityRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegSetKeySecurityOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegSetKeySecurityResponse structure represents the BaseRegSetKeySecurity operation response
type BaseRegSetKeySecurityResponse struct {
	// Return: The BaseRegSetKeySecurity return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegSetKeySecurityResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegSetKeySecurityOperation) *xxx_BaseRegSetKeySecurityOperation {
	if op == nil {
		op = &xxx_BaseRegSetKeySecurityOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *BaseRegSetKeySecurityResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegSetKeySecurityOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BaseRegSetKeySecurityResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegSetKeySecurityResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegSetKeySecurityOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegSetValueOperation structure represents the BaseRegSetValue operation
type xxx_BaseRegSetValueOperation struct {
	Key        *Key           `idl:"name:hKey" json:"key"`
	ValueName  *UnicodeString `idl:"name:lpValueName" json:"value_name"`
	Type       uint32         `idl:"name:dwType" json:"type"`
	Data       []byte         `idl:"name:lpData;size_is:(cbData)" json:"data"`
	DataLength uint32         `idl:"name:cbData" json:"data_length"`
	Return     uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegSetValueOperation) OpNum() int { return 22 }

func (o *xxx_BaseRegSetValueOperation) OpName() string { return "/winreg/v1/BaseRegSetValue" }

func (o *xxx_BaseRegSetValueOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.Data != nil && o.DataLength == 0 {
		o.DataLength = uint32(len(o.Data))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegSetValueOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpValueName {in} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.ValueName != nil {
			if err := o.ValueName.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Type); err != nil {
			return err
		}
	}
	// lpData {in} (1:{alias=LPBYTE}*(1))(2:{alias=BYTE}[dim:0,size_is=cbData](uchar))
	{
		dimSize1 := uint64(o.DataLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.Data {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if err := w.WriteData(o.Data[i1]); err != nil {
				return err
			}
		}
		for i1 := len(o.Data); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WriteData(uint8(0)); err != nil {
				return err
			}
		}
	}
	// cbData {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.DataLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegSetValueOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpValueName {in} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.ValueName == nil {
			o.ValueName = &UnicodeString{}
		}
		if err := o.ValueName.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwType {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Type); err != nil {
			return err
		}
	}
	// lpData {in} (1:{alias=LPBYTE,pointer=ref}*(1))(2:{alias=BYTE}[dim:0,size_is=cbData](uchar))
	{
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Data", sizeInfo[0])
		}
		o.Data = make([]byte, sizeInfo[0])
		for i1 := range o.Data {
			i1 := i1
			if err := w.ReadData(&o.Data[i1]); err != nil {
				return err
			}
		}
	}
	// cbData {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.DataLength); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegSetValueOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegSetValueOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BaseRegSetValueOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BaseRegSetValueRequest structure represents the BaseRegSetValue operation request
type BaseRegSetValueRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegCreateKey, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
	// lpValueName: MUST be a pointer to an RRP_UNICODE_STRING structure that contains the
	// name of the value (as specified in section 3.1.1.5) to set.
	ValueName *UnicodeString `idl:"name:lpValueName" json:"value_name"`
	// dwType: The type of data to be stored. This MUST be one of the REG_VALUE_TYPE values
	// specified in section 3.1.1.5.
	Type uint32 `idl:"name:dwType" json:"type"`
	// lpData: A pointer to a buffer that contains the data to set for the value entry.
	Data []byte `idl:"name:lpData;size_is:(cbData)" json:"data"`
	// cbData: The length in bytes of the information to be stored.
	DataLength uint32 `idl:"name:cbData" json:"data_length"`
}

func (o *BaseRegSetValueRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegSetValueOperation) *xxx_BaseRegSetValueOperation {
	if op == nil {
		op = &xxx_BaseRegSetValueOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.ValueName = o.ValueName
	op.Type = o.Type
	op.Data = o.Data
	op.DataLength = o.DataLength
	return op
}

func (o *BaseRegSetValueRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegSetValueOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.ValueName = op.ValueName
	o.Type = op.Type
	o.Data = op.Data
	o.DataLength = op.DataLength
}
func (o *BaseRegSetValueRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegSetValueRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegSetValueOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegSetValueResponse structure represents the BaseRegSetValue operation response
type BaseRegSetValueResponse struct {
	// Return: The BaseRegSetValue return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegSetValueResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegSetValueOperation) *xxx_BaseRegSetValueOperation {
	if op == nil {
		op = &xxx_BaseRegSetValueOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *BaseRegSetValueResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegSetValueOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BaseRegSetValueResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegSetValueResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegSetValueOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegUnloadKeyOperation structure represents the BaseRegUnLoadKey operation
type xxx_BaseRegUnloadKeyOperation struct {
	Key    *Key           `idl:"name:hKey" json:"key"`
	SubKey *UnicodeString `idl:"name:lpSubKey" json:"sub_key"`
	Return uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegUnloadKeyOperation) OpNum() int { return 23 }

func (o *xxx_BaseRegUnloadKeyOperation) OpName() string { return "/winreg/v1/BaseRegUnLoadKey" }

func (o *xxx_BaseRegUnloadKeyOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegUnloadKeyOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpSubKey {in} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.SubKey != nil {
			if err := o.SubKey.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegUnloadKeyOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpSubKey {in} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.SubKey == nil {
			o.SubKey = &UnicodeString{}
		}
		if err := o.SubKey.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegUnloadKeyOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegUnloadKeyOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BaseRegUnloadKeyOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BaseRegUnloadKeyRequest structure represents the BaseRegUnLoadKey operation request
type BaseRegUnloadKeyRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenCurrentUser and OpenLocalMachine.
	Key *Key `idl:"name:hKey" json:"key"`
	// lpSubKey: An optional pointer to an RRP_UNICODE_STRING structure that MUST contain
	// the relative name, as specified in section 3.1.1.1.2. The lpSubKey parameter points
	// to the name of the key that is to be unloaded.
	SubKey *UnicodeString `idl:"name:lpSubKey" json:"sub_key"`
}

func (o *BaseRegUnloadKeyRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegUnloadKeyOperation) *xxx_BaseRegUnloadKeyOperation {
	if op == nil {
		op = &xxx_BaseRegUnloadKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.SubKey = o.SubKey
	return op
}

func (o *BaseRegUnloadKeyRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegUnloadKeyOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.SubKey = op.SubKey
}
func (o *BaseRegUnloadKeyRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegUnloadKeyRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegUnloadKeyOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegUnloadKeyResponse structure represents the BaseRegUnLoadKey operation response
type BaseRegUnloadKeyResponse struct {
	// Return: The BaseRegUnLoadKey return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegUnloadKeyResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegUnloadKeyOperation) *xxx_BaseRegUnloadKeyOperation {
	if op == nil {
		op = &xxx_BaseRegUnloadKeyOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *BaseRegUnloadKeyResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegUnloadKeyOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BaseRegUnloadKeyResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegUnloadKeyResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegUnloadKeyOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegGetVersionOperation structure represents the BaseRegGetVersion operation
type xxx_BaseRegGetVersionOperation struct {
	Key     *Key   `idl:"name:hKey" json:"key"`
	Version uint32 `idl:"name:lpdwVersion" json:"version"`
	Return  uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegGetVersionOperation) OpNum() int { return 26 }

func (o *xxx_BaseRegGetVersionOperation) OpName() string { return "/winreg/v1/BaseRegGetVersion" }

func (o *xxx_BaseRegGetVersionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegGetVersionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_BaseRegGetVersionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegGetVersionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegGetVersionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// lpdwVersion {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
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

func (o *xxx_BaseRegGetVersionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// lpdwVersion {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
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

// BaseRegGetVersionRequest structure represents the BaseRegGetVersion operation request
type BaseRegGetVersionRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegCreateKey, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
}

func (o *BaseRegGetVersionRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegGetVersionOperation) *xxx_BaseRegGetVersionOperation {
	if op == nil {
		op = &xxx_BaseRegGetVersionOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	return op
}

func (o *BaseRegGetVersionRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegGetVersionOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
}
func (o *BaseRegGetVersionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegGetVersionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegGetVersionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegGetVersionResponse structure represents the BaseRegGetVersion operation response
type BaseRegGetVersionResponse struct {
	// lpdwVersion: A buffer in which the registry version MUST be returned. The registry
	// version is implementation-specific.<27>
	Version uint32 `idl:"name:lpdwVersion" json:"version"`
	// Return: The BaseRegGetVersion return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegGetVersionResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegGetVersionOperation) *xxx_BaseRegGetVersionOperation {
	if op == nil {
		op = &xxx_BaseRegGetVersionOperation{}
	}
	if o == nil {
		return op
	}
	op.Version = o.Version
	op.Return = o.Return
	return op
}

func (o *BaseRegGetVersionResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegGetVersionOperation) {
	if o == nil {
		return
	}
	o.Version = op.Version
	o.Return = op.Return
}
func (o *BaseRegGetVersionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegGetVersionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegGetVersionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenCurrentConfigOperation structure represents the OpenCurrentConfig operation
type xxx_OpenCurrentConfigOperation struct {
	ServerName    string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
	Key           *Key   `idl:"name:phKey" json:"key"`
	Return        uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenCurrentConfigOperation) OpNum() int { return 27 }

func (o *xxx_OpenCurrentConfigOperation) OpName() string { return "/winreg/v1/OpenCurrentConfig" }

func (o *xxx_OpenCurrentConfigOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenCurrentConfigOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME, names=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// samDesired {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenCurrentConfigOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME, names=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// samDesired {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenCurrentConfigOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenCurrentConfigOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phKey {out} (1:{alias=PRPC_HKEY}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_OpenCurrentConfigOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phKey {out} (1:{alias=PRPC_HKEY,pointer=ref}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
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

// OpenCurrentConfigRequest structure represents the OpenCurrentConfig operation request
type OpenCurrentConfigRequest struct {
	// ServerName: This SHOULD be sent as NULL and MUST be ignored on receipt because the
	// binding to the server is already complete at this stage.
	ServerName string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	// samDesired: A bit field that describes the wanted security access for the key. It
	// MUST be constructed from one or more of the values that are specified in section
	// 2.2.3.
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
}

func (o *OpenCurrentConfigRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenCurrentConfigOperation) *xxx_OpenCurrentConfigOperation {
	if op == nil {
		op = &xxx_OpenCurrentConfigOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.DesiredAccess = o.DesiredAccess
	return op
}

func (o *OpenCurrentConfigRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenCurrentConfigOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.DesiredAccess = op.DesiredAccess
}
func (o *OpenCurrentConfigRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenCurrentConfigRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenCurrentConfigOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenCurrentConfigResponse structure represents the OpenCurrentConfig operation response
type OpenCurrentConfigResponse struct {
	// phKey: A handle to the root key, HKEY_CURRENT_CONFIG, as specified in section 3.1.1.7.
	Key *Key `idl:"name:phKey" json:"key"`
	// Return: The OpenCurrentConfig return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *OpenCurrentConfigResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenCurrentConfigOperation) *xxx_OpenCurrentConfigOperation {
	if op == nil {
		op = &xxx_OpenCurrentConfigOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.Return = o.Return
	return op
}

func (o *OpenCurrentConfigResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenCurrentConfigOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.Return = op.Return
}
func (o *OpenCurrentConfigResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenCurrentConfigResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenCurrentConfigOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegQueryMultipleValuesOperation structure represents the BaseRegQueryMultipleValues operation
type xxx_BaseRegQueryMultipleValuesOperation struct {
	Key          *Key          `idl:"name:hKey" json:"key"`
	ValueListIn  []*ValueEntry `idl:"name:val_listIn;size_is:(num_vals);length_is:(num_vals)" json:"value_list_in"`
	ValueListOut []*ValueEntry `idl:"name:val_listOut;size_is:(num_vals);length_is:(num_vals)" json:"value_list_out"`
	ValsLength   uint32        `idl:"name:num_vals" json:"vals_length"`
	Buffer       []byte        `idl:"name:lpvalueBuf;size_is:(ldwTotsize);length_is:(ldwTotsize);pointer:unique" json:"buffer"`
	TotalSize    uint32        `idl:"name:ldwTotsize;pointer:ref" json:"total_size"`
	Return       uint32        `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegQueryMultipleValuesOperation) OpNum() int { return 29 }

func (o *xxx_BaseRegQueryMultipleValuesOperation) OpName() string {
	return "/winreg/v1/BaseRegQueryMultipleValues"
}

func (o *xxx_BaseRegQueryMultipleValuesOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.ValueListIn != nil && o.ValsLength == 0 {
		o.ValsLength = uint32(len(o.ValueListIn))
	}
	if o.ValueListIn != nil && o.ValsLength == 0 {
		o.ValsLength = uint32(len(o.ValueListIn))
	}
	if o.Buffer != nil && o.TotalSize == 0 {
		o.TotalSize = uint32(len(o.Buffer))
	}
	if o.Buffer != nil && o.TotalSize == 0 {
		o.TotalSize = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegQueryMultipleValuesOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// val_listIn {in} (1:{alias=PRVALENT}*(1))(2:{alias=RVALENT}[dim:0,size_is=num_vals,length_is=num_vals](struct))
	{
		dimSize1 := uint64(o.ValsLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.ValsLength)
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
		for i1 := range o.ValueListIn {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.ValueListIn[i1] != nil {
				if err := o.ValueListIn[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ValueEntry{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.ValueListIn); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&ValueEntry{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// num_vals {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ValsLength); err != nil {
			return err
		}
	}
	// lpvalueBuf {in, out} (1:{pointer=unique}*(1)[dim:0,size_is=ldwTotsize,length_is=ldwTotsize](char))
	{
		if o.Buffer != nil || o.TotalSize > 0 {
			_ptr_lpvalueBuf := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.TotalSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				dimLength1 := uint64(o.TotalSize)
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
			if err := w.WritePointer(&o.Buffer, _ptr_lpvalueBuf); err != nil {
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
	// ldwTotsize {in, out} (1:{pointer=ref, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TotalSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegQueryMultipleValuesOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// val_listIn {in} (1:{alias=PRVALENT,pointer=ref}*(1))(2:{alias=RVALENT}[dim:0,size_is=num_vals,length_is=num_vals](struct))
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
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ValueListIn", sizeInfo[0])
		}
		o.ValueListIn = make([]*ValueEntry, sizeInfo[0])
		for i1 := range o.ValueListIn {
			i1 := i1
			if o.ValueListIn[i1] == nil {
				o.ValueListIn[i1] = &ValueEntry{}
			}
			if err := o.ValueListIn[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// num_vals {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ValsLength); err != nil {
			return err
		}
	}
	// lpvalueBuf {in, out} (1:{pointer=unique}*(1)[dim:0,size_is=ldwTotsize,length_is=ldwTotsize](char))
	{
		_ptr_lpvalueBuf := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		_s_lpvalueBuf := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_lpvalueBuf, _ptr_lpvalueBuf); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ldwTotsize {in, out} (1:{pointer=ref, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TotalSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegQueryMultipleValuesOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.Buffer != nil && o.TotalSize == 0 {
		o.TotalSize = uint32(len(o.Buffer))
	}
	if o.Buffer != nil && o.TotalSize == 0 {
		o.TotalSize = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegQueryMultipleValuesOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// val_listOut {out} (1:{alias=PRVALENT}*(1))(2:{alias=RVALENT}[dim:0,size_is=num_vals,length_is=num_vals](struct))
	{
		dimSize1 := uint64(o.ValsLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.ValsLength)
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
		for i1 := range o.ValueListOut {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.ValueListOut[i1] != nil {
				if err := o.ValueListOut[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ValueEntry{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.ValueListOut); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&ValueEntry{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpvalueBuf {in, out} (1:{pointer=unique}*(1)[dim:0,size_is=ldwTotsize,length_is=ldwTotsize](char))
	{
		if o.Buffer != nil || o.TotalSize > 0 {
			_ptr_lpvalueBuf := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.TotalSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				dimLength1 := uint64(o.TotalSize)
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
			if err := w.WritePointer(&o.Buffer, _ptr_lpvalueBuf); err != nil {
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
	// ldwTotsize {in, out} (1:{pointer=ref, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TotalSize); err != nil {
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

func (o *xxx_BaseRegQueryMultipleValuesOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// val_listOut {out} (1:{alias=PRVALENT,pointer=ref}*(1))(2:{alias=RVALENT}[dim:0,size_is=num_vals,length_is=num_vals](struct))
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
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ValueListOut", sizeInfo[0])
		}
		o.ValueListOut = make([]*ValueEntry, sizeInfo[0])
		for i1 := range o.ValueListOut {
			i1 := i1
			if o.ValueListOut[i1] == nil {
				o.ValueListOut[i1] = &ValueEntry{}
			}
			if err := o.ValueListOut[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpvalueBuf {in, out} (1:{pointer=unique}*(1)[dim:0,size_is=ldwTotsize,length_is=ldwTotsize](char))
	{
		_ptr_lpvalueBuf := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		_s_lpvalueBuf := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_lpvalueBuf, _ptr_lpvalueBuf); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ldwTotsize {in, out} (1:{pointer=ref, alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TotalSize); err != nil {
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

// BaseRegQueryMultipleValuesRequest structure represents the BaseRegQueryMultipleValues operation request
type BaseRegQueryMultipleValuesRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegCreateKey, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
	// val_listIn: A pointer to an array of RVALENT structures, one for each value to be
	// queried. The array holds the list of value names for which the type and data MUST
	// be returned.
	ValueListIn []*ValueEntry `idl:"name:val_listIn;size_is:(num_vals);length_is:(num_vals)" json:"value_list_in"`
	// num_vals: The size in bytes of the val_list array.
	ValsLength uint32 `idl:"name:num_vals" json:"vals_length"`
	// lpvalueBuf: Returns the data for each value that is specified by the val_listOut
	// parameter.
	Buffer []byte `idl:"name:lpvalueBuf;size_is:(ldwTotsize);length_is:(ldwTotsize);pointer:unique" json:"buffer"`
	// ldwTotsize: The value that indicates the length in bytes of the lpvalueBuf parameter.
	TotalSize uint32 `idl:"name:ldwTotsize;pointer:ref" json:"total_size"`
}

func (o *BaseRegQueryMultipleValuesRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegQueryMultipleValuesOperation) *xxx_BaseRegQueryMultipleValuesOperation {
	if op == nil {
		op = &xxx_BaseRegQueryMultipleValuesOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.ValueListIn = o.ValueListIn
	op.ValsLength = o.ValsLength
	op.Buffer = o.Buffer
	op.TotalSize = o.TotalSize
	return op
}

func (o *BaseRegQueryMultipleValuesRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegQueryMultipleValuesOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.ValueListIn = op.ValueListIn
	o.ValsLength = op.ValsLength
	o.Buffer = op.Buffer
	o.TotalSize = op.TotalSize
}
func (o *BaseRegQueryMultipleValuesRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegQueryMultipleValuesRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegQueryMultipleValuesOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegQueryMultipleValuesResponse structure represents the BaseRegQueryMultipleValues operation response
type BaseRegQueryMultipleValuesResponse struct {
	// XXX: num_vals is an implicit input depedency for output parameters
	ValsLength uint32 `idl:"name:num_vals" json:"vals_length"`

	// val_listOut: A pointer to an array of RVALENT structures, one for each value to be
	// queried.
	ValueListOut []*ValueEntry `idl:"name:val_listOut;size_is:(num_vals);length_is:(num_vals)" json:"value_list_out"`
	// lpvalueBuf: Returns the data for each value that is specified by the val_listOut
	// parameter.
	Buffer []byte `idl:"name:lpvalueBuf;size_is:(ldwTotsize);length_is:(ldwTotsize);pointer:unique" json:"buffer"`
	// ldwTotsize: The value that indicates the length in bytes of the lpvalueBuf parameter.
	TotalSize uint32 `idl:"name:ldwTotsize;pointer:ref" json:"total_size"`
	// Return: The BaseRegQueryMultipleValues return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegQueryMultipleValuesResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegQueryMultipleValuesOperation) *xxx_BaseRegQueryMultipleValuesOperation {
	if op == nil {
		op = &xxx_BaseRegQueryMultipleValuesOperation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.ValsLength == uint32(0) {
		op.ValsLength = o.ValsLength
	}

	op.ValueListOut = o.ValueListOut
	op.Buffer = o.Buffer
	op.TotalSize = o.TotalSize
	op.Return = o.Return
	return op
}

func (o *BaseRegQueryMultipleValuesResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegQueryMultipleValuesOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.ValsLength = op.ValsLength

	o.ValueListOut = op.ValueListOut
	o.Buffer = op.Buffer
	o.TotalSize = op.TotalSize
	o.Return = op.Return
}
func (o *BaseRegQueryMultipleValuesResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegQueryMultipleValuesResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegQueryMultipleValuesOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegSaveKeyExOperation structure represents the BaseRegSaveKeyEx operation
type xxx_BaseRegSaveKeyExOperation struct {
	Key                *Key                `idl:"name:hKey" json:"key"`
	File               *UnicodeString      `idl:"name:lpFile" json:"file"`
	SecurityAttributes *SecurityAttributes `idl:"name:pSecurityAttributes;pointer:unique" json:"security_attributes"`
	Flags              uint32              `idl:"name:Flags" json:"flags"`
	Return             uint32              `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegSaveKeyExOperation) OpNum() int { return 31 }

func (o *xxx_BaseRegSaveKeyExOperation) OpName() string { return "/winreg/v1/BaseRegSaveKeyEx" }

func (o *xxx_BaseRegSaveKeyExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegSaveKeyExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpFile {in} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.File != nil {
			if err := o.File.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pSecurityAttributes {in} (1:{pointer=unique, alias=PRPC_SECURITY_ATTRIBUTES}*(1))(2:{alias=RPC_SECURITY_ATTRIBUTES}(struct))
	{
		if o.SecurityAttributes != nil {
			_ptr_pSecurityAttributes := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.SecurityAttributes != nil {
					if err := o.SecurityAttributes.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&SecurityAttributes{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SecurityAttributes, _ptr_pSecurityAttributes); err != nil {
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
	// Flags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegSaveKeyExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpFile {in} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.File == nil {
			o.File = &UnicodeString{}
		}
		if err := o.File.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pSecurityAttributes {in} (1:{pointer=unique, alias=PRPC_SECURITY_ATTRIBUTES}*(1))(2:{alias=RPC_SECURITY_ATTRIBUTES}(struct))
	{
		_ptr_pSecurityAttributes := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.SecurityAttributes == nil {
				o.SecurityAttributes = &SecurityAttributes{}
			}
			if err := o.SecurityAttributes.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_pSecurityAttributes := func(ptr interface{}) { o.SecurityAttributes = *ptr.(**SecurityAttributes) }
		if err := w.ReadPointer(&o.SecurityAttributes, _s_pSecurityAttributes, _ptr_pSecurityAttributes); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
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

func (o *xxx_BaseRegSaveKeyExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegSaveKeyExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BaseRegSaveKeyExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BaseRegSaveKeyExRequest structure represents the BaseRegSaveKeyEx operation request
type BaseRegSaveKeyExRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegCreateKey, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
	// lpFile: A pointer to an RRP_UNICODE_STRING structure that contains the name of the
	// file in which the specified key and subkeys are saved. The format of the file name
	// is implementation-specific.
	File *UnicodeString `idl:"name:lpFile" json:"file"`
	// pSecurityAttributes: A pointer to an RPC_SECURITY_ATTRIBUTES structure that specifies
	// a security descriptor for the new file. If the pSecurityAttributes parameter is NULL,
	// the file receives a default security descriptor.
	SecurityAttributes *SecurityAttributes `idl:"name:pSecurityAttributes;pointer:unique" json:"security_attributes"`
	// Flags: Specifies the format for the saved key. This MUST be one of the following
	// values.
	//
	//	+-------+-----------------------------------------------------+
	//	|       |                                                     |
	//	| VALUE |                       MEANING                       |
	//	|       |                                                     |
	//	+-------+-----------------------------------------------------+
	//	+-------+-----------------------------------------------------+
	//	|     1 | The key or subtree is saved in the original format. |
	//	+-------+-----------------------------------------------------+
	//	|     2 | The key or subtree is saved in the latest format.   |
	//	+-------+-----------------------------------------------------+
	//	|     4 | The key or subtree is saved without compression.    |
	//	+-------+-----------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
}

func (o *BaseRegSaveKeyExRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegSaveKeyExOperation) *xxx_BaseRegSaveKeyExOperation {
	if op == nil {
		op = &xxx_BaseRegSaveKeyExOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.File = o.File
	op.SecurityAttributes = o.SecurityAttributes
	op.Flags = o.Flags
	return op
}

func (o *BaseRegSaveKeyExRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegSaveKeyExOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.File = op.File
	o.SecurityAttributes = op.SecurityAttributes
	o.Flags = op.Flags
}
func (o *BaseRegSaveKeyExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegSaveKeyExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegSaveKeyExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegSaveKeyExResponse structure represents the BaseRegSaveKeyEx operation response
type BaseRegSaveKeyExResponse struct {
	// Return: The BaseRegSaveKeyEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegSaveKeyExResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegSaveKeyExOperation) *xxx_BaseRegSaveKeyExOperation {
	if op == nil {
		op = &xxx_BaseRegSaveKeyExOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *BaseRegSaveKeyExResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegSaveKeyExOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BaseRegSaveKeyExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegSaveKeyExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegSaveKeyExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenPerformanceTextOperation structure represents the OpenPerformanceText operation
type xxx_OpenPerformanceTextOperation struct {
	ServerName    string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
	Key           *Key   `idl:"name:phKey" json:"key"`
	Return        uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenPerformanceTextOperation) OpNum() int { return 32 }

func (o *xxx_OpenPerformanceTextOperation) OpName() string { return "/winreg/v1/OpenPerformanceText" }

func (o *xxx_OpenPerformanceTextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPerformanceTextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME, names=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// samDesired {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPerformanceTextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME, names=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// samDesired {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPerformanceTextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPerformanceTextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phKey {out} (1:{alias=PRPC_HKEY}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_OpenPerformanceTextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phKey {out} (1:{alias=PRPC_HKEY,pointer=ref}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
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

// OpenPerformanceTextRequest structure represents the OpenPerformanceText operation request
type OpenPerformanceTextRequest struct {
	// ServerName: SHOULD be sent as NULL and MUST be ignored on receipt because the binding
	// to the server is already complete at this stage.
	ServerName string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	// samDesired: SHOULD be sent as 0 and MUST be ignored on receipt.
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
}

func (o *OpenPerformanceTextRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenPerformanceTextOperation) *xxx_OpenPerformanceTextOperation {
	if op == nil {
		op = &xxx_OpenPerformanceTextOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.DesiredAccess = o.DesiredAccess
	return op
}

func (o *OpenPerformanceTextRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenPerformanceTextOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.DesiredAccess = op.DesiredAccess
}
func (o *OpenPerformanceTextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenPerformanceTextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenPerformanceTextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenPerformanceTextResponse structure represents the OpenPerformanceText operation response
type OpenPerformanceTextResponse struct {
	// phKey: A pointer to a variable that receives a handle to the root key HKEY_PERFORMANCE_TEXT.
	Key *Key `idl:"name:phKey" json:"key"`
	// Return: The OpenPerformanceText return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *OpenPerformanceTextResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenPerformanceTextOperation) *xxx_OpenPerformanceTextOperation {
	if op == nil {
		op = &xxx_OpenPerformanceTextOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.Return = o.Return
	return op
}

func (o *OpenPerformanceTextResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenPerformanceTextOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.Return = op.Return
}
func (o *OpenPerformanceTextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenPerformanceTextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenPerformanceTextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_OpenPerformanceNlsTextOperation structure represents the OpenPerformanceNlsText operation
type xxx_OpenPerformanceNlsTextOperation struct {
	ServerName    string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
	Key           *Key   `idl:"name:phKey" json:"key"`
	Return        uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_OpenPerformanceNlsTextOperation) OpNum() int { return 33 }

func (o *xxx_OpenPerformanceNlsTextOperation) OpName() string {
	return "/winreg/v1/OpenPerformanceNlsText"
}

func (o *xxx_OpenPerformanceNlsTextOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPerformanceNlsTextOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME, names=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16String(ctx, w, o.ServerName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ServerName, _ptr_ServerName); err != nil {
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
	// samDesired {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.WriteData(o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPerformanceNlsTextOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{handle, pointer=unique, alias=PREGISTRY_SERVER_NAME, names=PWCHAR}*(1))(2:{alias=WCHAR}[dim:0,string](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16String(ctx, w, &o.ServerName); err != nil {
				return err
			}
			return nil
		})
		_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// samDesired {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.ReadData(&o.DesiredAccess); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPerformanceNlsTextOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_OpenPerformanceNlsTextOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// phKey {out} (1:{alias=PRPC_HKEY}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
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

func (o *xxx_OpenPerformanceNlsTextOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// phKey {out} (1:{alias=PRPC_HKEY,pointer=ref}*(1))(2:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
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

// OpenPerformanceNlsTextRequest structure represents the OpenPerformanceNlsText operation request
type OpenPerformanceNlsTextRequest struct {
	// ServerName: This SHOULD be sent as NULL and MUST be ignored on receipt because the
	// binding to the server is already complete at this stage.
	ServerName string `idl:"name:ServerName;pointer:unique" json:"server_name"`
	// samDesired: This SHOULD be sent as 0 and MUST be ignored on receipt.
	DesiredAccess uint32 `idl:"name:samDesired" json:"desired_access"`
}

func (o *OpenPerformanceNlsTextRequest) xxx_ToOp(ctx context.Context, op *xxx_OpenPerformanceNlsTextOperation) *xxx_OpenPerformanceNlsTextOperation {
	if op == nil {
		op = &xxx_OpenPerformanceNlsTextOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.DesiredAccess = o.DesiredAccess
	return op
}

func (o *OpenPerformanceNlsTextRequest) xxx_FromOp(ctx context.Context, op *xxx_OpenPerformanceNlsTextOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.DesiredAccess = op.DesiredAccess
}
func (o *OpenPerformanceNlsTextRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *OpenPerformanceNlsTextRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenPerformanceNlsTextOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// OpenPerformanceNlsTextResponse structure represents the OpenPerformanceNlsText operation response
type OpenPerformanceNlsTextResponse struct {
	// phKey: A pointer to a variable that receives a handle to the root key HKEY_PERFORMANCE_NLSTEXT,
	// as specified in section 3.1.1.7.
	Key *Key `idl:"name:phKey" json:"key"`
	// Return: The OpenPerformanceNlsText return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *OpenPerformanceNlsTextResponse) xxx_ToOp(ctx context.Context, op *xxx_OpenPerformanceNlsTextOperation) *xxx_OpenPerformanceNlsTextOperation {
	if op == nil {
		op = &xxx_OpenPerformanceNlsTextOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.Return = o.Return
	return op
}

func (o *OpenPerformanceNlsTextResponse) xxx_FromOp(ctx context.Context, op *xxx_OpenPerformanceNlsTextOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.Return = op.Return
}
func (o *OpenPerformanceNlsTextResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *OpenPerformanceNlsTextResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_OpenPerformanceNlsTextOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegQueryMultipleValues2Operation structure represents the BaseRegQueryMultipleValues2 operation
type xxx_BaseRegQueryMultipleValues2Operation struct {
	Key          *Key          `idl:"name:hKey" json:"key"`
	ValueListIn  []*ValueEntry `idl:"name:val_listIn;size_is:(num_vals);length_is:(num_vals)" json:"value_list_in"`
	ValueListOut []*ValueEntry `idl:"name:val_listOut;size_is:(num_vals);length_is:(num_vals)" json:"value_list_out"`
	ValsLength   uint32        `idl:"name:num_vals" json:"vals_length"`
	Buffer       []byte        `idl:"name:lpvalueBuf;size_is:(ldwTotsize);length_is:(ldwTotsize);pointer:unique" json:"buffer"`
	TotalSize    uint32        `idl:"name:ldwTotsize" json:"total_size"`
	RequiredSize uint32        `idl:"name:ldwRequiredSize" json:"required_size"`
	Return       uint32        `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegQueryMultipleValues2Operation) OpNum() int { return 34 }

func (o *xxx_BaseRegQueryMultipleValues2Operation) OpName() string {
	return "/winreg/v1/BaseRegQueryMultipleValues2"
}

func (o *xxx_BaseRegQueryMultipleValues2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.ValueListIn != nil && o.ValsLength == 0 {
		o.ValsLength = uint32(len(o.ValueListIn))
	}
	if o.ValueListIn != nil && o.ValsLength == 0 {
		o.ValsLength = uint32(len(o.ValueListIn))
	}
	if o.Buffer != nil && o.TotalSize == 0 {
		o.TotalSize = uint32(len(o.Buffer))
	}
	if o.Buffer != nil && o.TotalSize == 0 {
		o.TotalSize = uint32(len(o.Buffer))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegQueryMultipleValues2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// val_listIn {in} (1:{alias=PRVALENT}*(1))(2:{alias=RVALENT}[dim:0,size_is=num_vals,length_is=num_vals](struct))
	{
		dimSize1 := uint64(o.ValsLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.ValsLength)
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
		for i1 := range o.ValueListIn {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.ValueListIn[i1] != nil {
				if err := o.ValueListIn[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ValueEntry{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.ValueListIn); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&ValueEntry{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// num_vals {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.ValsLength); err != nil {
			return err
		}
	}
	// lpvalueBuf {in, out} (1:{pointer=unique}*(1)[dim:0,size_is=ldwTotsize,length_is=ldwTotsize](char))
	{
		if o.Buffer != nil || o.TotalSize > 0 {
			_ptr_lpvalueBuf := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.TotalSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				dimLength1 := uint64(o.TotalSize)
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
			if err := w.WritePointer(&o.Buffer, _ptr_lpvalueBuf); err != nil {
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
	// ldwTotsize {in} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.TotalSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegQueryMultipleValues2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// val_listIn {in} (1:{alias=PRVALENT,pointer=ref}*(1))(2:{alias=RVALENT}[dim:0,size_is=num_vals,length_is=num_vals](struct))
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
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ValueListIn", sizeInfo[0])
		}
		o.ValueListIn = make([]*ValueEntry, sizeInfo[0])
		for i1 := range o.ValueListIn {
			i1 := i1
			if o.ValueListIn[i1] == nil {
				o.ValueListIn[i1] = &ValueEntry{}
			}
			if err := o.ValueListIn[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// num_vals {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.ValsLength); err != nil {
			return err
		}
	}
	// lpvalueBuf {in, out} (1:{pointer=unique}*(1)[dim:0,size_is=ldwTotsize,length_is=ldwTotsize](char))
	{
		_ptr_lpvalueBuf := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		_s_lpvalueBuf := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_lpvalueBuf, _ptr_lpvalueBuf); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ldwTotsize {in} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.TotalSize); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegQueryMultipleValues2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegQueryMultipleValues2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// val_listOut {out} (1:{alias=PRVALENT}*(1))(2:{alias=RVALENT}[dim:0,size_is=num_vals,length_is=num_vals](struct))
	{
		dimSize1 := uint64(o.ValsLength)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		dimLength1 := uint64(o.ValsLength)
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
		for i1 := range o.ValueListOut {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.ValueListOut[i1] != nil {
				if err := o.ValueListOut[i1].MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&ValueEntry{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.ValueListOut); uint64(i1) < sizeInfo[0]; i1++ {
			if err := (&ValueEntry{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// lpvalueBuf {in, out} (1:{pointer=unique}*(1)[dim:0,size_is=ldwTotsize,length_is=ldwTotsize](char))
	{
		if o.Buffer != nil || o.TotalSize > 0 {
			_ptr_lpvalueBuf := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.TotalSize)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				dimLength1 := uint64(o.TotalSize)
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
			if err := w.WritePointer(&o.Buffer, _ptr_lpvalueBuf); err != nil {
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
	// ldwRequiredSize {out} (1:{alias=LPDWORD}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.RequiredSize); err != nil {
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

func (o *xxx_BaseRegQueryMultipleValues2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// val_listOut {out} (1:{alias=PRVALENT,pointer=ref}*(1))(2:{alias=RVALENT}[dim:0,size_is=num_vals,length_is=num_vals](struct))
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
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.ValueListOut", sizeInfo[0])
		}
		o.ValueListOut = make([]*ValueEntry, sizeInfo[0])
		for i1 := range o.ValueListOut {
			i1 := i1
			if o.ValueListOut[i1] == nil {
				o.ValueListOut[i1] = &ValueEntry{}
			}
			if err := o.ValueListOut[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// lpvalueBuf {in, out} (1:{pointer=unique}*(1)[dim:0,size_is=ldwTotsize,length_is=ldwTotsize](char))
	{
		_ptr_lpvalueBuf := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
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
		_s_lpvalueBuf := func(ptr interface{}) { o.Buffer = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.Buffer, _s_lpvalueBuf, _ptr_lpvalueBuf); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ldwRequiredSize {out} (1:{alias=LPDWORD,pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.RequiredSize); err != nil {
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

// BaseRegQueryMultipleValues2Request structure represents the BaseRegQueryMultipleValues2 operation request
type BaseRegQueryMultipleValues2Request struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegCreateKey, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText. The server SHOULD
	// NOT process requests on predefined keys.
	Key *Key `idl:"name:hKey" json:"key"`
	// val_listIn: A pointer to an array of RVALENT structures, one for each value to query.
	// The array holds the list of value names for which the type and data MUST be returned.
	ValueListIn []*ValueEntry `idl:"name:val_listIn;size_is:(num_vals);length_is:(num_vals)" json:"value_list_in"`
	// num_vals: The size as the number of RVALENT structures of the val_list array.
	ValsLength uint32 `idl:"name:num_vals" json:"vals_length"`
	// lpvalueBuf: The data for each value that is specified by the val_listOut parameter.
	Buffer []byte `idl:"name:lpvalueBuf;size_is:(ldwTotsize);length_is:(ldwTotsize);pointer:unique" json:"buffer"`
	// ldwTotsize: A value that indicates the size in bytes of lpvalueBuf.
	TotalSize uint32 `idl:"name:ldwTotsize" json:"total_size"`
}

func (o *BaseRegQueryMultipleValues2Request) xxx_ToOp(ctx context.Context, op *xxx_BaseRegQueryMultipleValues2Operation) *xxx_BaseRegQueryMultipleValues2Operation {
	if op == nil {
		op = &xxx_BaseRegQueryMultipleValues2Operation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.ValueListIn = o.ValueListIn
	op.ValsLength = o.ValsLength
	op.Buffer = o.Buffer
	op.TotalSize = o.TotalSize
	return op
}

func (o *BaseRegQueryMultipleValues2Request) xxx_FromOp(ctx context.Context, op *xxx_BaseRegQueryMultipleValues2Operation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.ValueListIn = op.ValueListIn
	o.ValsLength = op.ValsLength
	o.Buffer = op.Buffer
	o.TotalSize = op.TotalSize
}
func (o *BaseRegQueryMultipleValues2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegQueryMultipleValues2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegQueryMultipleValues2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegQueryMultipleValues2Response structure represents the BaseRegQueryMultipleValues2 operation response
type BaseRegQueryMultipleValues2Response struct {
	// XXX: num_vals is an implicit input depedency for output parameters
	ValsLength uint32 `idl:"name:num_vals" json:"vals_length"`
	// XXX: ldwTotsize is an implicit input depedency for output parameters
	TotalSize uint32 `idl:"name:ldwTotsize" json:"total_size"`

	// val_listOut: A pointer to an array of RVALENT structures, one for each value to be
	// queried. This parameter is a placeholder to return the type, size, and data offset
	// for each requested value.
	ValueListOut []*ValueEntry `idl:"name:val_listOut;size_is:(num_vals);length_is:(num_vals)" json:"value_list_out"`
	// lpvalueBuf: The data for each value that is specified by the val_listOut parameter.
	Buffer []byte `idl:"name:lpvalueBuf;size_is:(ldwTotsize);length_is:(ldwTotsize);pointer:unique" json:"buffer"`
	// ldwRequiredSize: If lpvalueBuf is not large enough to contain all the data, this
	// parameter MUST return the size in bytes that is needed for lpvalueBuf to contain
	// all the required data.
	RequiredSize uint32 `idl:"name:ldwRequiredSize" json:"required_size"`
	// Return: The BaseRegQueryMultipleValues2 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegQueryMultipleValues2Response) xxx_ToOp(ctx context.Context, op *xxx_BaseRegQueryMultipleValues2Operation) *xxx_BaseRegQueryMultipleValues2Operation {
	if op == nil {
		op = &xxx_BaseRegQueryMultipleValues2Operation{}
	}
	if o == nil {
		return op
	}
	// XXX: implicit input dependencies for output parameters
	if op.ValsLength == uint32(0) {
		op.ValsLength = o.ValsLength
	}
	if op.TotalSize == uint32(0) {
		op.TotalSize = o.TotalSize
	}

	op.ValueListOut = o.ValueListOut
	op.Buffer = o.Buffer
	op.RequiredSize = o.RequiredSize
	op.Return = o.Return
	return op
}

func (o *BaseRegQueryMultipleValues2Response) xxx_FromOp(ctx context.Context, op *xxx_BaseRegQueryMultipleValues2Operation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.ValsLength = op.ValsLength
	o.TotalSize = op.TotalSize

	o.ValueListOut = op.ValueListOut
	o.Buffer = op.Buffer
	o.RequiredSize = op.RequiredSize
	o.Return = op.Return
}
func (o *BaseRegQueryMultipleValues2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegQueryMultipleValues2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegQueryMultipleValues2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_BaseRegDeleteKeyExOperation structure represents the BaseRegDeleteKeyEx operation
type xxx_BaseRegDeleteKeyExOperation struct {
	Key        *Key           `idl:"name:hKey" json:"key"`
	SubKey     *UnicodeString `idl:"name:lpSubKey" json:"sub_key"`
	AccessMask uint32         `idl:"name:AccessMask" json:"access_mask"`
	_          uint32         `idl:"name:Reserved"`
	Return     uint32         `idl:"name:Return" json:"return"`
}

func (o *xxx_BaseRegDeleteKeyExOperation) OpNum() int { return 35 }

func (o *xxx_BaseRegDeleteKeyExOperation) OpName() string { return "/winreg/v1/BaseRegDeleteKeyEx" }

func (o *xxx_BaseRegDeleteKeyExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegDeleteKeyExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key != nil {
			if err := o.Key.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Key{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// lpSubKey {in} (1:{alias=PRRP_UNICODE_STRING}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.SubKey != nil {
			if err := o.SubKey.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&UnicodeString{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// AccessMask {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.WriteData(o.AccessMask); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		if err := w.WriteData(uint32(0)); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegDeleteKeyExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// hKey {in} (1:{context_handle, alias=RPC_HKEY, names=ndr_context_handle}(struct))
	{
		if o.Key == nil {
			o.Key = &Key{}
		}
		if err := o.Key.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// lpSubKey {in} (1:{alias=PRRP_UNICODE_STRING,pointer=ref}*(1))(2:{alias=RRP_UNICODE_STRING, names=RPC_UNICODE_STRING}(struct))
	{
		if o.SubKey == nil {
			o.SubKey = &UnicodeString{}
		}
		if err := o.SubKey.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// AccessMask {in} (1:{alias=REGSAM, names=ULONG}(uint32))
	{
		if err := w.ReadData(&o.AccessMask); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{alias=DWORD}(uint32))
	{
		// reserved Reserved
		var _Reserved uint32
		if err := w.ReadData(&_Reserved); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegDeleteKeyExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_BaseRegDeleteKeyExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_BaseRegDeleteKeyExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:(error_status_t))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// BaseRegDeleteKeyExRequest structure represents the BaseRegDeleteKeyEx operation request
type BaseRegDeleteKeyExRequest struct {
	// hKey: A handle to a key that MUST have been opened previously by using one of the
	// open methods that are specified in section 3.1.5: OpenClassesRoot, OpenCurrentUser,
	// OpenLocalMachine, OpenPerformanceData, OpenUsers, BaseRegCreateKey, BaseRegOpenKey,
	// OpenCurrentConfig, OpenPerformanceText, OpenPerformanceNlsText.
	Key *Key `idl:"name:hKey" json:"key"`
	// lpSubKey: A pointer to an RRP_UNICODE_STRING structure that MUST specify the name
	// of the key (as specified in section 3.1.1) to delete.
	SubKey *UnicodeString `idl:"name:lpSubKey" json:"sub_key"`
	// AccessMask: A bit field that describes the wanted security access for the key.
	//
	//	+----------------------------+--------------------------------------------------------+
	//	|                            |                                                        |
	//	|           VALUE            |                        MEANING                         |
	//	|                            |                                                        |
	//	+----------------------------+--------------------------------------------------------+
	//	+----------------------------+--------------------------------------------------------+
	//	| KEY_WOW64_64KEY 0x00000100 | Explicitly delete the key in the 64-bit key namespace. |
	//	+----------------------------+--------------------------------------------------------+
	//	| KEY_WOW64_32KEY 0x00000200 | Explicitly delete the key in the 32-bit key namespace. |
	//	+----------------------------+--------------------------------------------------------+
	AccessMask uint32 `idl:"name:AccessMask" json:"access_mask"`
}

func (o *BaseRegDeleteKeyExRequest) xxx_ToOp(ctx context.Context, op *xxx_BaseRegDeleteKeyExOperation) *xxx_BaseRegDeleteKeyExOperation {
	if op == nil {
		op = &xxx_BaseRegDeleteKeyExOperation{}
	}
	if o == nil {
		return op
	}
	op.Key = o.Key
	op.SubKey = o.SubKey
	op.AccessMask = o.AccessMask
	return op
}

func (o *BaseRegDeleteKeyExRequest) xxx_FromOp(ctx context.Context, op *xxx_BaseRegDeleteKeyExOperation) {
	if o == nil {
		return
	}
	o.Key = op.Key
	o.SubKey = op.SubKey
	o.AccessMask = op.AccessMask
}
func (o *BaseRegDeleteKeyExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *BaseRegDeleteKeyExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegDeleteKeyExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// BaseRegDeleteKeyExResponse structure represents the BaseRegDeleteKeyEx operation response
type BaseRegDeleteKeyExResponse struct {
	// Return: The BaseRegDeleteKeyEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *BaseRegDeleteKeyExResponse) xxx_ToOp(ctx context.Context, op *xxx_BaseRegDeleteKeyExOperation) *xxx_BaseRegDeleteKeyExOperation {
	if op == nil {
		op = &xxx_BaseRegDeleteKeyExOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *BaseRegDeleteKeyExResponse) xxx_FromOp(ctx context.Context, op *xxx_BaseRegDeleteKeyExOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *BaseRegDeleteKeyExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *BaseRegDeleteKeyExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_BaseRegDeleteKeyExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
