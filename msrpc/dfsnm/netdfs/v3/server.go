package netdfs

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

// netdfs server interface.
type NetdfsServer interface {

	// The NetrDfsManagerGetVersion method returns the version number of the DFS server
	// in use on the server.
	//
	// The NetrDfsManagerGetVersion method has the following MIDL syntax.
	//
	// This method has no parameters.
	//
	// Return Values: This method MUST return one of the following values.
	//
	//	+--------------+----------------------------------------------------------------------------------+
	//	|    RETURN    |                                                                                  |
	//	|    VALUE     |                                   DESCRIPTION                                    |
	//	|              |                                                                                  |
	//	+--------------+----------------------------------------------------------------------------------+
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x00000001   | The server MUST support stand-alone DFS namespaces and opnums from 0 through 5,  |
	//	|              | inclusive. The server MAY support domain-based DFS namespaces and other opnums.  |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x00000002   | In addition to the preceding, the server MUST support domainv1-based DFS         |
	//	|              | namespaces and opnums 10 through 22, inclusive. The server MAY support hosting   |
	//	|              | more than one DFS namespace on the same server.                                  |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x00000004   | In addition to the preceding, the server MUST support hosting more than one DFS  |
	//	|              | namespace on the same server and Level parameter value 200 of the NetrDfsEnumEx  |
	//	|              | method. It SHOULD support opnum 6.                                               |
	//	+--------------+----------------------------------------------------------------------------------+
	//	| 0x00000006   | In addition to the preceding, the server MUST support domainv2-based DFS         |
	//	|              | namespace and opnums 23 through 25, inclusive.                                   |
	//	+--------------+----------------------------------------------------------------------------------+
	//
	// The clients MAY use the version information to determine the RPC methods that the
	// DFS server supports.<37><38><39><40><41>
	ManagerGetVersion(context.Context, *ManagerGetVersionRequest) (*ManagerGetVersionResponse, error)

	// The NetrDfsAdd method creates a new DFS link or adds a new target to an existing
	// link of a DFS namespace.
	//
	// The NetrDfsAdd (Opnum 1) method has the following MIDL syntax.
	//
	// Return Values: The method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF] section
	// 2.2. The most common error codes are listed in the following table.
	//
	//	+------------------------------------+-------------------------------------------------------+
	//	|               RETURN               |                                                       |
	//	|             VALUE/CODE             |                      DESCRIPTION                      |
	//	|                                    |                                                       |
	//	+------------------------------------+-------------------------------------------------------+
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | Successful completion.                                |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Permission to perform the operation was denied.       |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x00000050 ERROR_FILE_EXISTS       | The specified DFS link target already exists.         |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | An incorrect parameter was specified.                 |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x00000490 ERROR_NOT_FOUND         | The specified DFS root namespace does not exist.      |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The method does not support a domain-based namespace. |
	//	+------------------------------------+-------------------------------------------------------+
	//	| 0x00000906 NERR_NetNameNotFound    | The DFS link target does not exist.                   |
	//	+------------------------------------+-------------------------------------------------------+
	//
	// The NetrDfsAdd method SHOULD<42> support a domain-based DFS namespace. If it does
	// not support a domain-based DFS namespace it MUST return ERROR_NOT_SUPPORTED.
	Add(context.Context, *AddRequest) (*AddResponse, error)

	// The NetrDfsRemove method removes a link or a link target from a DFS namespace. A
	// link can be removed regardless of the number of targets associated with it.
	//
	// The NetrDfsRemove method has the following MIDL syntax.
	//
	// Return Values: This method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF] section
	// 2.2. The most common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | Successful completion.                                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified DFS link target was not found as a target of the specified DFS     |
	//	|                                    | link.                                                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Permission to perform the operation was denied.                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | An incorrect parameter was specified.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The method does not support a domain-based namespace.                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000490 ERROR_NOT_FOUND         | The specified DFS namespace or DFS link does not exist.                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// The NetrDfsRemove method SHOULD<49> support a domain-based DFS namespace. If it does
	// not support a domain-based DFS namespace it MUST return ERROR_NOT_SUPPORTED.
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)

	// The NetrDfsSetInfo method sets or modifies information relevant to a specific DFS
	// root, DFS root target, DFS link, or DFS link target.
	//
	// The NetrDfsSetInfo method uses the following MIDL syntax.
	//
	// Return Values: This method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF],
	// section 2.2. The most common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | Successful completion.                                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified DFS link target was not found as a target of the specified DFS     |
	//	|                                    | link.                                                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Permission to perform the operation was denied.                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified operation is not supported.                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | An incorrect parameter was specified.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000490 ERROR_NOT_FOUND         | The specified DFS root namespace or DFS link, or DFS link or root target, does   |
	//	|                                    | not exist.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// The NetrDfsSetInfo method SHOULD<56> support a domain-based DFS namespace. If it
	// does not support a domain-based DFS namespace it MUST return ERROR_NOT_SUPPORTED.
	SetInfo(context.Context, *SetInfoRequest) (*SetInfoResponse, error)

	// The NetrDfsGetInfo method returns information about a DFS root or a DFS link of the
	// specified DFS namespace.
	//
	// The NetrDfsGetInfo method has the following MIDL syntax.
	//
	// Return Values: The method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF] section
	// 2.2. The most common error codes are listed in the following table.
	//
	//	+------------------------------------+--------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                |
	//	|             VALUE/CODE             |                                  DESCRIPTION                                   |
	//	|                                    |                                                                                |
	//	+------------------------------------+--------------------------------------------------------------------------------+
	//	+------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | Successful completion.                                                         |
	//	+------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified operation is not supported.                                      |
	//	+------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | An incorrect parameter was specified.                                          |
	//	+------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00000490 ERROR_NOT_FOUND         | The specified DFS root or DFS link, or DFS link or root target does not exist. |
	//	+------------------------------------+--------------------------------------------------------------------------------+
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error)

	// The NetrDfsEnum method enumerates the DFS root hosted on a server or the DFS links
	// of the namespace hosted by a server. Depending on the information level, the targets
	// of the root and links are also displayed.
	//
	// The NetrDfsEnum method uses the following MIDL syntax.
	//
	// Return Values: The method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF] section
	// 2.2. The most common error codes are listed in the following table.
	//
	//	+---------------------------------------+--------------------------------------------------+
	//	|                RETURN                 |                                                  |
	//	|              VALUE/CODE               |                   DESCRIPTION                    |
	//	|                                       |                                                  |
	//	+---------------------------------------+--------------------------------------------------+
	//	+---------------------------------------+--------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS              | Successful completion.                           |
	//	+---------------------------------------+--------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER    | An incorrect parameter was specified.            |
	//	+---------------------------------------+--------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS        | There is no data to return.                      |
	//	+---------------------------------------+--------------------------------------------------+
	//	| 0x00000490 ERROR_NOT_FOUND            | The specified DFS root namespace does not exist. |
	//	+---------------------------------------+--------------------------------------------------+
	//	| 0x000010DF ERROR_DEVICE_NOT_AVAILABLE | The server hosts more than one root.             |
	//	+---------------------------------------+--------------------------------------------------+
	//
	// A server MAY<66> implement this method.
	Enum(context.Context, *EnumRequest) (*EnumResponse, error)

	// The NetrDfsMove (Opnum 6) method renames or moves a DFS link. This method has the
	// following MIDL syntax.
	//
	// Return Values: The method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF] section
	// 2.2. The most common error codes are listed in the following table.
	//
	//	+------------------------------------+--------------------------------------------------------------------------+
	//	|               RETURN               |                                                                          |
	//	|             VALUE/CODE             |                               DESCRIPTION                                |
	//	|                                    |                                                                          |
	//	+------------------------------------+--------------------------------------------------------------------------+
	//	+------------------------------------+--------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | Successful completion.                                                   |
	//	+------------------------------------+--------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Permission to perform the operation was denied.                          |
	//	+------------------------------------+--------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified operation is not supported.                                |
	//	+------------------------------------+--------------------------------------------------------------------------+
	//	| 0x00000050 ERROR_FILE_EXISTS       | The destination path specifies an existing link.                         |
	//	+------------------------------------+--------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | An incorrect parameter was specified.                                    |
	//	+------------------------------------+--------------------------------------------------------------------------+
	//	| 0x00000490 ERROR_NOT_FOUND         | A specified DFS root namespace does not exist, or no links were matched. |
	//	+------------------------------------+--------------------------------------------------------------------------+
	Move(context.Context, *MoveRequest) (*MoveResponse, error)

	// Opnum7NotUsedOnWire operation.
	// Opnum7NotUsedOnWire

	// Opnum8NotUsedOnWire operation.
	// Opnum8NotUsedOnWire

	// Opnum9NotUsedOnWire operation.
	// Opnum9NotUsedOnWire

	// The NetrDfsAddFtRoot (Opnum 10) method creates a new domainv1-based DFS namespace
	// or adds a root target to an existing namespace.
	//
	// The NetrDfsAddFtRoot method uses the following MIDL syntax.
	//
	// Return Values: The method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF] section
	// 2.2. The most common error codes are listed in the following table.
	//
	//	+------------------------------------+--------------------------------------------------------------------------+
	//	|               RETURN               |                                                                          |
	//	|             VALUE/CODE             |                               DESCRIPTION                                |
	//	|                                    |                                                                          |
	//	+------------------------------------+--------------------------------------------------------------------------+
	//	+------------------------------------+--------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | Successful completion.                                                   |
	//	+------------------------------------+--------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Permission to perform the operation was denied.                          |
	//	+------------------------------------+--------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | An incorrect parameter was specified.                                    |
	//	+------------------------------------+--------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS    | A namespace of the specified name already exists on the server.          |
	//	+------------------------------------+--------------------------------------------------------------------------+
	//	| 0x00000906 NERR_NetNameNotFound    | The share that the RootShare parameter specifies does not already exist. |
	//	+------------------------------------+--------------------------------------------------------------------------+
	//
	// The share that the RootShare parameter specifies MUST already exist on the server.
	//
	// If the DcName parameter is not NULL, the server assumes that this is the PDC for
	// the domain in which the DFS namespace is to be created.
	//
	// If the domain-based DFS namespace already exists, and the ServerName and RootShare
	// parameters are a root target, the server MUST fail with ERROR_ALREADY_EXISTS.
	//
	// If the share that the RootShare parameter specifies does not already exist, the RPC
	// method MUST fail with NERR_NetNameNotFound (0x00000906).
	AddFTRoot(context.Context, *AddFTRootRequest) (*AddFTRootResponse, error)

	// The NetrDfsRemoveFtRoot (Opnum 11) method removes the specified root target from
	// a domainv1-based DFS namespace.<112> If the target is the last one associated with
	// the DFS namespace, then this method also deletes the DFS namespace. The DFS namespace
	// can be removed without first removing all of the links in it.
	//
	// If a client tries to use this method on a domainv2-based DFS namespace target, then
	// the server MUST fail with the return value of ERROR_NOT_SUPPORTED.
	//
	// The NetrDfsRemoveFtRoot method uses the following MIDL syntax.
	//
	// Return Values: The method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF] section
	// 2.2.The most common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | Successful completion.                                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified DFS root target was not found as a target of the specified DFS     |
	//	|                                    | namespace.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Permission to perform the operation was denied.                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | An incorrect parameter was specified.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000490 ERROR_NOT_FOUND         | The specified DFS rootnamespace does not exist.                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	RemoveFTRoot(context.Context, *RemoveFTRootRequest) (*RemoveFTRootResponse, error)

	// The NetrDfsAddStdRoot (Opnum 12) method creates a new stand-alone DFS namespace.<118><119>
	//
	// The NetrDfsAddStdRoot method uses the following MIDL syntax.
	//
	// Return Values: The method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF] section
	// 2.2. The most common error codes are listed in the following table.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | Successful completion.                                                           |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED  | Permission to perform the operation was denied.                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000050 ERROR_FILE_EXISTS    | The DFS namespace that the ServerName and RootShare parameters specify already   |
	//	|                                 | exists.<120>                                                                     |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS | The DFS namespace that the ServerName and RootShare parameters specify already   |
	//	|                                 | exists.<121>                                                                     |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000906 NERR_NetNameNotFound | The share that the RootShare parameter specifies does not already exist.         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	AddStdRoot(context.Context, *AddStdRootRequest) (*AddStdRootResponse, error)

	// The NetrDfsRemoveStdRoot (Opnum 13) method deletes the specified stand-alone DFS
	// namespace.<122> The DFS namespace can be removed without first removing all of the
	// links in it.
	//
	// The NetrDfsRemoveStdRoot method uses the following MIDL syntax.
	//
	// Return Values: The method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF] section
	// 2.2. The most common error codes are listed in the following table.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | Successful completion.                                                           |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | Permission to perform the operation was denied.                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000490 ERROR_NOT_FOUND     | The DFS namespace that the ServerName and RootShare parameters specify does not  |
	//	|                                | already exist.                                                                   |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	RemoveStdRoot(context.Context, *RemoveStdRootRequest) (*RemoveStdRootResponse, error)

	// The NetrDfsManagerInitialize method instructs the DFS server to discard its current
	// state and reinitialize itself from its stored configuration settings. The server
	// SHOULD<35> choose to implement this method.
	//
	// The NetrDfsManagerInitialize method has the following Microsoft Interface Definition
	// Language (MIDL) syntax.
	//
	// Return Values: The method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF] section
	// 2.2. The most common error codes are listed in the following table.
	//
	//	+--------------------------------+--------------------------------------------------+
	//	|             RETURN             |                                                  |
	//	|           VALUE/CODE           |                   DESCRIPTION                    |
	//	|                                |                                                  |
	//	+--------------------------------+--------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | Successful completion.                           |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | Permission to perform the operation was denied.  |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED | Server does not support the requested operation. |
	//	+--------------------------------+--------------------------------------------------+
	//
	// If this method is implemented, the DFS server SHOULD<36> discard its current state
	// and reinitialize itself from its stored configuration settings.
	ManagerInitialize(context.Context, *ManagerInitializeRequest) (*ManagerInitializeResponse, error)

	// The NetrDfsAddStdRootForced (Opnum 15) method creates a new stand-alone DFS namespace
	// without checking for the availability and accessibility of the specified share.<123><124><125>
	//
	// The NetrDfsAddStdRootForced method uses the following MIDL syntax.
	//
	// Return Values: The method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF] section
	// 2.2. The most common error codes are listed in the following table.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | Successful completion.                                                           |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED  | Permission to perform the operation was denied.                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS | The DFS namespace that the ServerName and RootShare parameters specify already   |
	//	|                                 | exists.                                                                          |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// The support for this method is optional. If supported, then the server MUST support
	// the ability to create a DFS namespace even when the share that the RootShare parameter
	// specifies is not available or accessible.
	AddStdRootForced(context.Context, *AddStdRootForcedRequest) (*AddStdRootForcedResponse, error)

	// The NetrDfsGetDcAddress (Opnum 16) method returns the DC host name that is used by
	// the DFS server to which the RPC method is issued.<126> The client MUST use this DC
	// to create a domain-based DFS namespace, add a root target to a domain-based DFS namespace,
	// remove a root target from a domain-based DFS namespace, or remove a domain-based
	// DFS namespace.
	//
	// The NetrDfsGetDcAddress method uses the following MIDL syntax.
	//
	// Return Values: The method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF] section
	// 2.2. The most common error codes are listed in the following table.
	//
	//	+--------------------------+------------------------+
	//	|          RETURN          |                        |
	//	|        VALUE/CODE        |      DESCRIPTION       |
	//	|                          |                        |
	//	+--------------------------+------------------------+
	//	+--------------------------+------------------------+
	//	| 0x00000000 ERROR_SUCCESS | Successful completion. |
	//	+--------------------------+------------------------+
	//
	// A server MAY<131> implement this method if it supports domain-based DFS namespaces.
	//
	// In the DcName parameter, the server SHOULD return the host name of the DC it is using
	// to access DFS metadata for any domain-based DFS namespace that the server is hosting.
	// If the server is not currently using a DC, it MUST determine a DC and return its
	// name.
	//
	// The server SHOULD also return a time-out value, in seconds, that is equal to the
	// length of time that the server will be using the DC in the Timeout parameter, assuming
	// that another RPC method does not change it.
	//
	// The server uses the IsRoot parameter to specify whether it supports the ability to
	// host more than one DFS namespace, and to indicate whether it is currently hosting
	// a DFS namespace. If the server supports the ability to host more than one DFS namespace,
	// it MUST return a value of FALSE in the IsRoot parameter, regardless of whether it
	// is actually hosting a DFS namespace. If the server does not support the ability to
	// host more than one DFS namespace, and if it currently hosts a DFS namespace, it SHOULD
	// return a value of TRUE in the IsRoot parameter; otherwise, it SHOULD return FALSE.
	GetDCAddress(context.Context, *GetDCAddressRequest) (*GetDCAddressResponse, error)

	// The NetrDfsSetDcAddress (Opnum 17) method instructs the server receiving the RPC
	// method to use the specified DC for DFS metadata accesses for domain-based DFS namespaces.<132>
	//
	// The NetrDfsSetDcAddress method uses the following MIDL syntax.
	//
	// Return Values: The method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF] section
	// 2.2. The most common error codes are listed in the following table.
	//
	//	+--------------------------------+-------------------------------------------------+
	//	|             RETURN             |                                                 |
	//	|           VALUE/CODE           |                   DESCRIPTION                   |
	//	|                                |                                                 |
	//	+--------------------------------+-------------------------------------------------+
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | Successful completion.                          |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED | Permission to perform the operation was denied. |
	//	+--------------------------------+-------------------------------------------------+
	//
	// Servers MAY choose not to implement this method or implement it as a method with
	// no effect that returns ERROR_SUCCESS.<134>
	//
	// Otherwise, the server MUST update the DC it uses for accessing DFS metadata for the
	// domain-based DFS namespace it is hosting with the specified DC. If no time-out is
	// specified in the Timeout parameter (NET_DFS_SETDC_TIMEOUT is not set in the Flags
	// parameter), the server MUST use its default time-out. The DC the server uses at the
	// end of this time-out is implementation-defined.
	//
	// When NET_DFS_SETDC_INIT_PKT is set in the Flags parameter, the server SHOULD initiate
	// a background synchronization of the domain-based DFS namespace it is hosting with
	// either the DC specified by this method or the default DC the server is using. This
	// MUST be treated as functionally equivalent to receiving a NetrDfsSetInfo (Opnum 3)
	// method with the Level parameter value 101 and the State field of DFS_INFO_101 set
	// to DFS_VOLUME_STATE_RESYNCHRONIZE.<135>
	SetDCAddress(context.Context, *SetDCAddressRequest) (*SetDCAddressResponse, error)

	// The NetrDfsFlushFtTable method instructs the DFS server on a DC to purge the specified
	// domainv1-based DFS entry from any DFS root referral cache it might have.
	//
	// Note  This method MUST fail on non-DC servers, as specified in this section.
	//
	// The NetrDfsFlushFtTable method uses the following MIDL syntax.
	//
	// Return Values: This method MUST return 0 on success and a nonzero error code on failure.
	// The values transmitted in this field are implementation-specific. For protocol purposes,
	// all nonzero values MUST be treated as equivalent failures.
	//
	// Note  This method MUST return ERROR_NOT_SUPPORTED on non-DC servers.
	//
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN             |                                                                                  |
	//	|           VALUE/CODE           |                                   DESCRIPTION                                    |
	//	|                                |                                                                                  |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | Successful completion.                                                           |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED | Operation not supported. This MUST be returned if the server does not implement  |
	//	|                                | the method.                                                                      |
	//	+--------------------------------+----------------------------------------------------------------------------------+
	//
	// The server MAY choose not to implement this method.<149> If it does, ERROR_NOT_SUPPORTED
	// MUST be returned.
	FlushFTTable(context.Context, *FlushFTTableRequest) (*FlushFTTableResponse, error)

	// The NetrDfsAdd2 (Opnum 19) method creates a new DFS link or adds a new target to
	// an existing link of a DFS namespace.
	//
	// The NetrDfsAdd2 method has the following MIDL syntax.
	//
	// Return Values: The method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF] section
	// 2.2. The most common error codes are listed in the following table.
	//
	//	+------------------------------------+-------------------------------------------------+
	//	|               RETURN               |                                                 |
	//	|             VALUE/CODE             |                   DESCRIPTION                   |
	//	|                                    |                                                 |
	//	+------------------------------------+-------------------------------------------------+
	//	+------------------------------------+-------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | The operation completed successfully.           |
	//	+------------------------------------+-------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Permission to perform the operation was denied. |
	//	+------------------------------------+-------------------------------------------------+
	//	| 0x00000050 ERROR_FILE_EXISTS       | The specified DFS link target already exists.   |
	//	+------------------------------------+-------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | An incorrect parameter was specified.           |
	//	+------------------------------------+-------------------------------------------------+
	//	| 0x00000490 ERROR_NOT_FOUND         | The specified DFS namespace does not exist.     |
	//	+------------------------------------+-------------------------------------------------+
	//	| 0x00000906 NERR_NetNameNotFound    | The DFS link target does not exist.             |
	//	+------------------------------------+-------------------------------------------------+
	//
	// A server MAY<81> implement this method.
	//
	// If the NetrDfsAdd (Opnum 1) method on a server does not support a domain-based namespace,
	// the server SHOULD support a domain-based namespace in the NetrDfsAdd2 (Opnum 19)
	// method. <82><83>
	Add2(context.Context, *Add2Request) (*Add2Response, error)

	// The NetrDfsRemove2 (Opnum 20) method removes the specified link or link target.
	//
	// The NetrDfsRemove2 method uses the following MIDL syntax.
	//
	// Return Values: The method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF] section
	// 2.2. The most common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | Successful completion.                                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified DFS link target was not found as a target of the specified DFS     |
	//	|                                    | link.                                                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Permission to perform the operation was denied.                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | An incorrect parameter was specified.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000490 ERROR_NOT_FOUND         | The specified DFS namespace or DFS link does not exist.                          |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// The server MAY<89> implement this method.
	//
	// If the NetrDfsRemove (Opnum 2) method on a server does not support a domain-based
	// namespace, the server SHOULD support a domain-based namespace in the NetrDfsRemove2
	// (Opnum 20) method.<90><91>
	Remove2(context.Context, *Remove2Request) (*Remove2Response, error)

	// The NetrDfsEnumEx (Opnum 21) method enumerates the DFS roots hosted on a server,
	// or DFS links of a namespace hosted by the server.<92><93> Depending on the information
	// level, the targets associated with the roots and links are also displayed.
	//
	// The NetrDfsEnumEx method uses the following MIDL syntax.
	//
	// Return Values: The method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF],
	// section 2.2. The most common error codes are listed in the following table.
	//
	//	+------------------------------------+--------------------------------------------------+
	//	|               RETURN               |                                                  |
	//	|             VALUE/CODE             |                   DESCRIPTION                    |
	//	|                                    |                                                  |
	//	+------------------------------------+--------------------------------------------------+
	//	+------------------------------------+--------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | Successful completion.                           |
	//	+------------------------------------+--------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | An incorrect parameter was specified.            |
	//	+------------------------------------+--------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS     | There is no data to return.                      |
	//	+------------------------------------+--------------------------------------------------+
	//	| 0x00000490 ERROR_NOT_FOUND         | The specified DFS root namespace does not exist. |
	//	+------------------------------------+--------------------------------------------------+
	EnumEx(context.Context, *EnumExRequest) (*EnumExResponse, error)

	// The NetrDfsSetInfo2 (Opnum 22) method sets or modifies the information associated
	// with a DFS root, a DFS root target, a DFS link, or a DFS link target.
	//
	// The NetrDfsSetInfo2 method has the following MIDL syntax.
	//
	// Return Values: The method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF] section
	// 2.2. The most common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | Successful completion.                                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified DFS link target was not found as a target of the specified DFS     |
	//	|                                    | link.                                                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Permission to perform the operation was denied.                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000032 ERROR_NOT_SUPPORTED     | The specified operation is not supported.                                        |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | An incorrect parameter was specified.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000490 ERROR_NOT_FOUND         | The specified DFS root, DFS link, or DFS link or root target does not exist.     |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//
	// The server MAY<105> implement this method.
	//
	// If the NetrDfsSetInfo (Opnum 3) method on a server does not support a domain-based
	// namespace, the server SHOULD support a domain-based namespace in the NetrDfsSetInfo2
	// (Opnum 22) method.<106><107>
	SetInfo2(context.Context, *SetInfo2Request) (*SetInfo2Response, error)

	// The NetrDfsAddRootTarget method is used to create a stand-alone DFS namespace, a
	// domainv1-based DFS namespace, or a domainv2-based DFS namespace.<73>
	//
	// The NetrDfsAddRootTarget method uses the following MIDL syntax.
	//
	// Return Values: The method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF] section
	// 2.2. The most common error codes are listed in the following table.
	//
	//	+---------------------------------+---------------------------------------------------------------------------+
	//	|             RETURN              |                                                                           |
	//	|           VALUE/CODE            |                                DESCRIPTION                                |
	//	|                                 |                                                                           |
	//	+---------------------------------+---------------------------------------------------------------------------+
	//	+---------------------------------+---------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | Successful completion.                                                    |
	//	+---------------------------------+---------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED  | Permission to perform the operation was denied.                           |
	//	+---------------------------------+---------------------------------------------------------------------------+
	//	| 0x000000B7 ERROR_ALREADY_EXISTS | The specified namespace already exists on this server.                    |
	//	+---------------------------------+---------------------------------------------------------------------------+
	//	| 0x00000906 NERR_NetNameNotFound | The share that the TargetPath parameter specifies does not already exist. |
	//	+---------------------------------+---------------------------------------------------------------------------+
	//
	// The following table summarizes the various actions that the NetrDfsAddRootTarget
	// method takes based on the parameter values.
	//
	//	+----------------------+-----------------------+------------------------+----------------------------------------------------------------------------------+
	//	|       PDFSPATH       |      PTARGETPATH      |      MAJORVERSION      |                                                                                  |
	//	|      PARAMETER       |       PARAMETER       |       PARAMETER        |                                   EXPLANATION                                    |
	//	|                      |                       |                        |                                                                                  |
	//	+----------------------+-----------------------+------------------------+----------------------------------------------------------------------------------+
	//	+----------------------+-----------------------+------------------------+----------------------------------------------------------------------------------+
	//	| \\<domain>\<dfsname> | \\<server>\<share>    |                      1 | Creates a new domainv1-based DFS namespace or adds a new root target to an       |
	//	|                      |                       |                        | existing domainv1-based DFS namespace. If a DFS namespace already exists, the    |
	//	|                      |                       |                        | version specified MUST match the DFS namespace; otherwise, the call fails.       |
	//	+----------------------+-----------------------+------------------------+----------------------------------------------------------------------------------+
	//	| \\<domain>\<dfsname> | \\<server>\<share>    |                      2 | Creates a new domainv2-based DFS namespace or adds a new root target to an       |
	//	|                      |                       |                        | existing domainv2-based DFS namespace. If a DFS namespace already exists, the    |
	//	|                      |                       |                        | version specified MUST match the DFS namespace; otherwise, the call fails.       |
	//	+----------------------+-----------------------+------------------------+----------------------------------------------------------------------------------+
	//	| \\<domain>\<dfsname> | \\<server>\<share>    |                      0 | Adds a new root target to an existing domain-based DFS namespace or a            |
	//	|                      |                       |                        | domainv2-based DFS namespace. If a DFS namespace does not already exist, the     |
	//	|                      |                       |                        | call fails.                                                                      |
	//	+----------------------+-----------------------+------------------------+----------------------------------------------------------------------------------+
	//	| \\<server>\<share>   | NULL                  |                      1 | Creates a new stand-alone DFS namespace.                                         |
	//	+----------------------+-----------------------+------------------------+----------------------------------------------------------------------------------+
	//
	// The following scheme is used to create a new domainv2-based DFS namespace:
	//
	// * NetrDfsGetSupportedNamespaceVersion ( 7c34ae60-9d4d-46b5-b2e1-e8ca09ac4d10 ) is
	// called to determine an appropriate version number to pass to the NetrDfsAddRootTarget()
	// method.
	//
	// * The client-side method creates a DFS metadata, format-independent LDAP ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_45643bfb-b4c4-432c-a10f-b98790063f8d
	// ) entry called the DFS namespace anchor. It contains only the DFS metadata major
	// version number.
	//
	// * Updates the access control list (ACL) ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_9f92aa05-dd0a-45f2-88d6-89f1fb654395
	// ) on the object ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_8bb43a65-7a8c-4585-a7ed-23044772f8ca
	// ) of the DFS namespace to permit read/write access by the DFS root target server.
	//
	// * The client-side method then issues an RPC ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_8a7f6700-8311-45bc-af10-82e10accd331
	// ) call to the DFS root target server.
	//
	// * The DFS server ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_3de8e640-501a-4b25-80a7-0ba769f3c0b9
	// ) creates a new DFS namespace LDAP entry with the DFS namespace anchor LDAP entry
	// as its parent.
	//
	// * All DFS links ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_0611e93d-f0e7-42ee-a591-d77ebcbb6619
	// ) are created with the DFS namespace LDAP entry as the parent. For more information,
	// see section 2.3.2 ( 6317bce5-9588-421d-a007-f61e459b4083 ).
	//
	// This results in two LDAP entries in domainv2 corresponding to the single LDAP entry
	// in domainv1.
	//
	// If the domain-based DFS namespace already exists, and the ServerName and RootShare
	// parameters are a root target, the server MUST fail with ERROR_ALREADY_EXISTS.
	//
	// If the share that the pTargetPath parameter specifies does not already exist, the
	// RPC method MUST fail with NERR_NetNameNotFound (0x00000906).
	AddRootTarget(context.Context, *AddRootTargetRequest) (*AddRootTargetResponse, error)

	// The NetrDfsRemoveRootTarget (Opnum 24) method is the unified DFS namespace deletion
	// method. It deletes stand-alone DFS namespaces, domainv1-based DFS namespaces, or
	// domainv2-based DFS namespaces based on parameters specified.<74>
	//
	// The NetrDfsRemoveRootTarget (Opnum 24) method has the following MIDL syntax.
	//
	// Return Values: The method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF] section
	// 2.2. The most common error codes are listed in the following table.
	//
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	|               RETURN               |                                                                                  |
	//	|             VALUE/CODE             |                                   DESCRIPTION                                    |
	//	|                                    |                                                                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS           | Successful completion.                                                           |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000002 ERROR_FILE_NOT_FOUND    | The specified DFS root target was not found as a target of the specified DFS     |
	//	|                                    | namespace.                                                                       |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000005 ERROR_ACCESS_DENIED     | Permission to perform the operation was denied.                                  |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER | An incorrect parameter was specified.                                            |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000490 ERROR_NOT_FOUND         | The specified DFS root namespace does not exist.                                 |
	//	+------------------------------------+----------------------------------------------------------------------------------+
	RemoveRootTarget(context.Context, *RemoveRootTargetRequest) (*RemoveRootTargetResponse, error)

	// The NetrDfsGetSupportedNamespaceVersion (Opnum 25) method is used to determine the
	// supported DFS metadata version number.<78>
	//
	// The NetrDfsGetSupportedNamespaceVersion (Opnum 25) method has the following MIDL
	// syntax.
	//
	// Return Values: The method MUST return 0 on success and a nonzero error code on failure.
	// The method can return any specific error code value, as specified in [MS-ERREF] section
	// 2.2. The most common error codes are listed in the following table.
	//
	//	+--------------------------+------------------------+
	//	|          RETURN          |                        |
	//	|        VALUE/CODE        |      DESCRIPTION       |
	//	|                          |                        |
	//	+--------------------------+------------------------+
	//	+--------------------------+------------------------+
	//	| 0x00000000 ERROR_SUCCESS | Successful completion. |
	//	+--------------------------+------------------------+
	//
	// The standalone namespace version supported by a server can be unaffected by the domain
	// metadata schema. If this is the case, the server MUST return a standalone DFS major
	// and minor version of zero for the DFS_NAMESPACE_VERSION_ORIGIN_DOMAIN query. In this
	// case, the standalone DFS capability field has no meaning and MUST also be zero.
	//
	// The version number of the DFS metadata that can be used for a new DFS namespace depends
	// on the following:
	//
	// * For domain-based DFS namespaces ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_c37de1c8-4bd3-406f-ad8c-50c877666f91
	// ) , the version supported by the DFS metadata schema in use in the server's domain
	// ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_b0276eb2-4e65-4cf1-a718-e0920a614aca ).
	//
	// * The version supported by the server that is to host the DFS root target ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_ac90b498-3ba4-48d6-bcd6-5495f1654671
	// ).
	//
	// Thus, the version that can be used for creating a new DFS namespace is the minimum
	// version that the domain and the server support.
	//
	// This method is useful in determining an appropriate version number to pass to the
	// NetrDfsAddRootTarget method.
	GetSupportedNamespaceVersion(context.Context, *GetSupportedNamespaceVersionRequest) (*GetSupportedNamespaceVersionResponse, error)
}

func RegisterNetdfsServer(conn dcerpc.Conn, o NetdfsServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewNetdfsServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(NetdfsSyntaxV3_0))...)
}

func NewNetdfsServerHandle(o NetdfsServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return NetdfsServerHandle(ctx, o, opNum, r)
	}
}

func NetdfsServerHandle(ctx context.Context, o NetdfsServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // NetrDfsManagerGetVersion
		op := &xxx_ManagerGetVersionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ManagerGetVersionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ManagerGetVersion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 1: // NetrDfsAdd
		op := &xxx_AddOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Add(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 2: // NetrDfsRemove
		op := &xxx_RemoveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Remove(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 3: // NetrDfsSetInfo
		op := &xxx_SetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 4: // NetrDfsGetInfo
		op := &xxx_GetInfoOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetInfoRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetInfo(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 5: // NetrDfsEnum
		op := &xxx_EnumOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Enum(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 6: // NetrDfsMove
		op := &xxx_MoveOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &MoveRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Move(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 7: // Opnum7NotUsedOnWire
		// Opnum7NotUsedOnWire
		return nil, nil
	case 8: // Opnum8NotUsedOnWire
		// Opnum8NotUsedOnWire
		return nil, nil
	case 9: // Opnum9NotUsedOnWire
		// Opnum9NotUsedOnWire
		return nil, nil
	case 10: // NetrDfsAddFtRoot
		op := &xxx_AddFTRootOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddFTRootRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddFTRoot(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 11: // NetrDfsRemoveFtRoot
		op := &xxx_RemoveFTRootOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveFTRootRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveFTRoot(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 12: // NetrDfsAddStdRoot
		op := &xxx_AddStdRootOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddStdRootRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddStdRoot(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 13: // NetrDfsRemoveStdRoot
		op := &xxx_RemoveStdRootOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveStdRootRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveStdRoot(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 14: // NetrDfsManagerInitialize
		op := &xxx_ManagerInitializeOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &ManagerInitializeRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.ManagerInitialize(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 15: // NetrDfsAddStdRootForced
		op := &xxx_AddStdRootForcedOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddStdRootForcedRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddStdRootForced(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 16: // NetrDfsGetDcAddress
		op := &xxx_GetDCAddressOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetDCAddressRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetDCAddress(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 17: // NetrDfsSetDcAddress
		op := &xxx_SetDCAddressOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetDCAddressRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetDCAddress(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 18: // NetrDfsFlushFtTable
		op := &xxx_FlushFTTableOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &FlushFTTableRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.FlushFTTable(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 19: // NetrDfsAdd2
		op := &xxx_Add2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Add2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Add2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 20: // NetrDfsRemove2
		op := &xxx_Remove2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &Remove2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.Remove2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 21: // NetrDfsEnumEx
		op := &xxx_EnumExOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &EnumExRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.EnumEx(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 22: // NetrDfsSetInfo2
		op := &xxx_SetInfo2Operation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &SetInfo2Request{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.SetInfo2(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 23: // NetrDfsAddRootTarget
		op := &xxx_AddRootTargetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &AddRootTargetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.AddRootTarget(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 24: // NetrDfsRemoveRootTarget
		op := &xxx_RemoveRootTargetOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &RemoveRootTargetRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.RemoveRootTarget(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	case 25: // NetrDfsGetSupportedNamespaceVersion
		op := &xxx_GetSupportedNamespaceVersionOperation{}
		if err := op.UnmarshalNDRRequest(ctx, r); err != nil {
			return nil, err
		}
		req := &GetSupportedNamespaceVersionRequest{}
		req.xxx_FromOp(ctx, op)
		resp, err := o.GetSupportedNamespaceVersion(ctx, req)
		return resp.xxx_ToOp(ctx, op), err
	}
	return nil, nil
}

// Unimplemented netdfs
type UnimplementedNetdfsServer struct {
}

func (UnimplementedNetdfsServer) ManagerGetVersion(context.Context, *ManagerGetVersionRequest) (*ManagerGetVersionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) Remove(context.Context, *RemoveRequest) (*RemoveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) SetInfo(context.Context, *SetInfoRequest) (*SetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) GetInfo(context.Context, *GetInfoRequest) (*GetInfoResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) Enum(context.Context, *EnumRequest) (*EnumResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) Move(context.Context, *MoveRequest) (*MoveResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) AddFTRoot(context.Context, *AddFTRootRequest) (*AddFTRootResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) RemoveFTRoot(context.Context, *RemoveFTRootRequest) (*RemoveFTRootResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) AddStdRoot(context.Context, *AddStdRootRequest) (*AddStdRootResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) RemoveStdRoot(context.Context, *RemoveStdRootRequest) (*RemoveStdRootResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) ManagerInitialize(context.Context, *ManagerInitializeRequest) (*ManagerInitializeResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) AddStdRootForced(context.Context, *AddStdRootForcedRequest) (*AddStdRootForcedResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) GetDCAddress(context.Context, *GetDCAddressRequest) (*GetDCAddressResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) SetDCAddress(context.Context, *SetDCAddressRequest) (*SetDCAddressResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) FlushFTTable(context.Context, *FlushFTTableRequest) (*FlushFTTableResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) Add2(context.Context, *Add2Request) (*Add2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) Remove2(context.Context, *Remove2Request) (*Remove2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) EnumEx(context.Context, *EnumExRequest) (*EnumExResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) SetInfo2(context.Context, *SetInfo2Request) (*SetInfo2Response, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) AddRootTarget(context.Context, *AddRootTargetRequest) (*AddRootTargetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) RemoveRootTarget(context.Context, *RemoveRootTargetRequest) (*RemoveRootTargetResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}
func (UnimplementedNetdfsServer) GetSupportedNamespaceVersion(context.Context, *GetSupportedNamespaceVersionRequest) (*GetSupportedNamespaceVersionResponse, error) {
	return nil, dcerpc.ErrNotImplemented
}

var _ NetdfsServer = (*UnimplementedNetdfsServer)(nil)
