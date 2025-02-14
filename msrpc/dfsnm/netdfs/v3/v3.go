package netdfs

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
	GoPackage = "dfsnm"
)

var (
	// Syntax UUID
	NetdfsSyntaxUUID = &uuid.UUID{TimeLow: 0x4fc742e0, TimeMid: 0x4a10, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x82, ClockSeqLow: 0x73, Node: [6]uint8{0x0, 0xaa, 0x0, 0x4a, 0xe6, 0x73}}
	// Syntax ID
	NetdfsSyntaxV3_0 = &dcerpc.SyntaxID{IfUUID: NetdfsSyntaxUUID, IfVersionMajor: 3, IfVersionMinor: 0}
)

// netdfs interface.
type NetdfsClient interface {

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
	ManagerGetVersion(context.Context, *ManagerGetVersionRequest, ...dcerpc.CallOption) (*ManagerGetVersionResponse, error)

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
	Add(context.Context, *AddRequest, ...dcerpc.CallOption) (*AddResponse, error)

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
	Remove(context.Context, *RemoveRequest, ...dcerpc.CallOption) (*RemoveResponse, error)

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
	SetInfo(context.Context, *SetInfoRequest, ...dcerpc.CallOption) (*SetInfoResponse, error)

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
	GetInfo(context.Context, *GetInfoRequest, ...dcerpc.CallOption) (*GetInfoResponse, error)

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
	Enum(context.Context, *EnumRequest, ...dcerpc.CallOption) (*EnumResponse, error)

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
	Move(context.Context, *MoveRequest, ...dcerpc.CallOption) (*MoveResponse, error)

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
	AddFTRoot(context.Context, *AddFTRootRequest, ...dcerpc.CallOption) (*AddFTRootResponse, error)

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
	RemoveFTRoot(context.Context, *RemoveFTRootRequest, ...dcerpc.CallOption) (*RemoveFTRootResponse, error)

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
	AddStdRoot(context.Context, *AddStdRootRequest, ...dcerpc.CallOption) (*AddStdRootResponse, error)

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
	RemoveStdRoot(context.Context, *RemoveStdRootRequest, ...dcerpc.CallOption) (*RemoveStdRootResponse, error)

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
	ManagerInitialize(context.Context, *ManagerInitializeRequest, ...dcerpc.CallOption) (*ManagerInitializeResponse, error)

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
	AddStdRootForced(context.Context, *AddStdRootForcedRequest, ...dcerpc.CallOption) (*AddStdRootForcedResponse, error)

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
	GetDCAddress(context.Context, *GetDCAddressRequest, ...dcerpc.CallOption) (*GetDCAddressResponse, error)

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
	SetDCAddress(context.Context, *SetDCAddressRequest, ...dcerpc.CallOption) (*SetDCAddressResponse, error)

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
	FlushFTTable(context.Context, *FlushFTTableRequest, ...dcerpc.CallOption) (*FlushFTTableResponse, error)

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
	Add2(context.Context, *Add2Request, ...dcerpc.CallOption) (*Add2Response, error)

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
	Remove2(context.Context, *Remove2Request, ...dcerpc.CallOption) (*Remove2Response, error)

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
	EnumEx(context.Context, *EnumExRequest, ...dcerpc.CallOption) (*EnumExResponse, error)

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
	SetInfo2(context.Context, *SetInfo2Request, ...dcerpc.CallOption) (*SetInfo2Response, error)

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
	AddRootTarget(context.Context, *AddRootTargetRequest, ...dcerpc.CallOption) (*AddRootTargetResponse, error)

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
	RemoveRootTarget(context.Context, *RemoveRootTargetRequest, ...dcerpc.CallOption) (*RemoveRootTargetResponse, error)

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
	GetSupportedNamespaceVersion(context.Context, *GetSupportedNamespaceVersionRequest, ...dcerpc.CallOption) (*GetSupportedNamespaceVersionResponse, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn
}

// TargetPriorityClass type represents DFS_TARGET_PRIORITY_CLASS RPC enumeration.
//
// The DFS_TARGET_PRIORITY_CLASS enumeration relates to the NetrDfsSetInfo and NetrDfsSetInfo2
// methods when used to set the priority of DFS targets in referrals from a server.
// For more information on prioritization, see section 2.2.2.7. The enumeration defines
// five possible DFS target priority class settings.
type TargetPriorityClass uint32

var (
	// DfsInvalidPriorityClass:  This is not a valid priority class.
	TargetPriorityClassInvalidPriorityClass TargetPriorityClass = 1
	// DfsSiteCostNormalPriorityClass:  The default or "normal" site cost priority class
	// for a DFS target.
	TargetPriorityClassSiteCostNormalPriorityClass TargetPriorityClass = 0
	// DfsGlobalHighPriorityClass:  The highest priority class for a DFS target. Targets
	// assigned to this class receive global preference.
	TargetPriorityClassGlobalHighPriorityClass TargetPriorityClass = 1
	// DfsSiteCostHighPriorityClass:  The highest site cost priority class for a DFS target.
	// Targets assigned to this class receive the highest preference among targets of the
	// same site cost for a given DFS client.
	TargetPriorityClassSiteCostHighPriorityClass TargetPriorityClass = 2
	// DfsSiteCostLowPriorityClass:  The lowest site cost priority class for a DFS target.
	// Targets assigned to this class receive the least preference among targets of the
	// same site cost for a given DFS client.
	TargetPriorityClassSiteCostLowPriorityClass TargetPriorityClass = 3
	// DfsGlobalLowPriorityClass:  The lowest priority class level for a DFS target. Targets
	// assigned to this class receive the least preference globally.
	//
	// The underlying data type of this enumeration is long integer.
	//
	// The order of priority classes, from highest to lowest, is as follows:
	//
	// * DfsGlobalHighPriorityClass
	//
	// * DfsSiteCostHighPriorityClass
	//
	// * DfsSiteCostNormalPriorityClass
	//
	// * DfsSiteCostLowPriorityClass
	//
	// * DfsGlobalLowPriorityClass
	//
	// Server targets are initially grouped into global high-priority, normal-priority,
	// and global low-priority classes. The normal-priority class is then subdivided, based
	// on site cost, into site cost high-priority, site cost normal-priority, and site-cost
	// low-priority classes.
	//
	// For example, all server targets with a site cost value of 0 are grouped into site
	// cost high-priority, normal-priority, and low-priority classes. Then, all server targets
	// with higher site costs are likewise separated into site cost high-priority, normal-priority,
	// and low-priority classes. Thus, a server target with a site cost value of 0 and a
	// site cost low-priority class is still ranked higher than a server target with a site
	// cost value of 1 and a site cost high-priority class.
	//
	// Be aware that the value for a "normal-priority class" is set to 0 even though it
	// is lower in priority than DfsGlobalHighPriorityClass and DfsSiteCostHighPriorityClass.
	// This is the default priority class setting. For added granularity, priority rank
	// can be used to discriminate within a priority class.
	TargetPriorityClassGlobalLowPriorityClass TargetPriorityClass = 4
)

func (o TargetPriorityClass) String() string {
	switch o {
	case TargetPriorityClassInvalidPriorityClass:
		return "TargetPriorityClassInvalidPriorityClass"
	case TargetPriorityClassSiteCostNormalPriorityClass:
		return "TargetPriorityClassSiteCostNormalPriorityClass"
	case TargetPriorityClassGlobalHighPriorityClass:
		return "TargetPriorityClassGlobalHighPriorityClass"
	case TargetPriorityClassSiteCostHighPriorityClass:
		return "TargetPriorityClassSiteCostHighPriorityClass"
	case TargetPriorityClassSiteCostLowPriorityClass:
		return "TargetPriorityClassSiteCostLowPriorityClass"
	case TargetPriorityClassGlobalLowPriorityClass:
		return "TargetPriorityClassGlobalLowPriorityClass"
	}
	return "Invalid"
}

// TargetPriority structure represents DFS_TARGET_PRIORITY RPC structure.
//
// The DFS_TARGET_PRIORITY structure relates to the NetrDfsSetInfo and NetrDfsSetInfo2
// methods when used to set the priority of a DFS target in referrals from a server.
// It also relates to the DFS_STORAGE_INFO_1 structure that the NetrDfsEnum, NetrDfsEnumEx,
// and NetrDfsGetInfo methods return. The structure defines the priority of a DFS target.
// The DFS targets can be prioritized independently of site cost. The DFS target priority
// is manually assigned to link targets and root targets and allows for load balancing
// of clients.
//
// The DFS_TARGET_PRIORITY structure has the following format.
type TargetPriority struct {
	// TargetPriorityClass:  The DFS_TARGET_PRIORITY_CLASS enumeration value that specifies
	// the priority class of the target. For more information, see section 2.2.2.8.
	TargetPriorityClass TargetPriorityClass `idl:"name:TargetPriorityClass" json:"target_priority_class"`
	// TargetPriorityRank:  The priority rank of the target, ranging in value from 0x0000
	// to 0x001F, where 0x0000 is the highest rank. Priority ranks apply only within a priority
	// class, not across priority classes.
	TargetPriorityRank uint16 `idl:"name:TargetPriorityRank" json:"target_priority_rank"`
	// Reserved:  MUST be set to 0 by the sender and ignored by the receiver.
	_ uint16 `idl:"name:Reserved"`
}

func (o *TargetPriority) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *TargetPriority) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteEnum(uint32(o.TargetPriorityClass)); err != nil {
		return err
	}
	if err := w.WriteData(o.TargetPriorityRank); err != nil {
		return err
	}
	// reserved Reserved
	if err := w.WriteData(uint16(0)); err != nil {
		return err
	}
	return nil
}
func (o *TargetPriority) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadEnum((*uint32)(&o.TargetPriorityClass)); err != nil {
		return err
	}
	if err := w.ReadData(&o.TargetPriorityRank); err != nil {
		return err
	}
	// reserved Reserved
	var _Reserved uint16
	if err := w.ReadData(&_Reserved); err != nil {
		return err
	}
	return nil
}

// StorageInfo structure represents DFS_STORAGE_INFO RPC structure.
//
// The DFS_STORAGE_INFO structure relates to the NetrDfsEnum, NetrDfsEnumEx, and NetrDfsGetInfo
// methods when used to enumerate DFS links and DFS targets in a namespace or to get
// information about a DFS link. The structure contains information about the target
// of a DFS root or DFS link.
//
// The DFS_STORAGE_INFO structure has the following format.
type StorageInfo struct {
	// State:  Refers to the State field of DFS_INFO_106. For more information, see section
	// 2.2.4.6.
	State uint32 `idl:"name:State" json:"state"`
	// ServerName:  The pointer to a null-terminated Unicode string containing the DFS target
	// host name.
	ServerName string `idl:"name:ServerName;string" json:"server_name"`
	// ShareName:  The pointer to a null-terminated Unicode string containing the DFS target
	// share name.
	//
	// DFS_INFO_3 and DFS_INFO_4 structures contain one or more DFS_STORAGE_INFO structures,
	// one for each DFS target.
	ShareName string `idl:"name:ShareName;string" json:"share_name"`
}

func (o *StorageInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *StorageInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if o.ServerName != "" {
		_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
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
	if o.ShareName != "" {
		_ptr_ShareName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ShareName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ShareName, _ptr_ShareName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *StorageInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
			return err
		}
		return nil
	})
	_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
		return err
	}
	_ptr_ShareName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ShareName); err != nil {
			return err
		}
		return nil
	})
	_s_ShareName := func(ptr interface{}) { o.ShareName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ShareName, _s_ShareName, _ptr_ShareName); err != nil {
		return err
	}
	return nil
}

// StorageInfo1 structure represents DFS_STORAGE_INFO_1 RPC structure.
//
// The DFS_STORAGE_INFO_1 structure relates to the NetrDfsEnum, NetrDfsEnumEx, and NetrDfsGetInfo
// methods when used to enumerate DFS links and targets in a namespace or to get information
// about a DFS link. The structure contains data about a DFS target, including the host
// name and share name, as well as the target state and priority. For more information
// on prioritization, see section 2.2.2.7.
//
// The DFS_STORAGE_INFO_1 structure has the following format.
type StorageInfo1 struct {
	// State:  Refers to the State field of DFS_INFO_106. For more information, see section
	// 2.2.4.6.
	State uint32 `idl:"name:State" json:"state"`
	// ServerName:  A pointer to a null-terminated Unicode string containing the DFS target
	// host name.
	ServerName string `idl:"name:ServerName;string" json:"server_name"`
	// ShareName:  A pointer to a null-terminated Unicode string containing the DFS target
	// share name.
	ShareName string `idl:"name:ShareName;string" json:"share_name"`
	// TargetPriority:  A DFS_TARGET_PRIORITY structure containing the priority class and
	// priority rank.
	TargetPriority *TargetPriority `idl:"name:TargetPriority" json:"target_priority"`
}

func (o *StorageInfo1) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *StorageInfo1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if o.ServerName != "" {
		_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
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
	if o.ShareName != "" {
		_ptr_ShareName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ShareName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ShareName, _ptr_ShareName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.TargetPriority != nil {
		if err := o.TargetPriority.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TargetPriority{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *StorageInfo1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
			return err
		}
		return nil
	})
	_s_ServerName := func(ptr interface{}) { o.ServerName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ServerName, _s_ServerName, _ptr_ServerName); err != nil {
		return err
	}
	_ptr_ShareName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ShareName); err != nil {
			return err
		}
		return nil
	})
	_s_ShareName := func(ptr interface{}) { o.ShareName = *ptr.(*string) }
	if err := w.ReadPointer(&o.ShareName, _s_ShareName, _ptr_ShareName); err != nil {
		return err
	}
	if o.TargetPriority == nil {
		o.TargetPriority = &TargetPriority{}
	}
	if err := o.TargetPriority.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// RootListEntry structure represents DFSM_ROOT_LIST_ENTRY RPC structure.
//
// The DFSM_ROOT_LIST_ENTRY structure relates to the NetrDfsAdd2, NetrDfsAddFtRoot,
// and NetrDfsSetInfo2 methods when used to add a DFS link or a DFS root target, or
// to modify the configuration of a domain-based DFS namespace. The structure contains
// information about a DFS root target.
//
// The DFSM_ROOT_LIST_ENTRY structure has the following format.
type RootListEntry struct {
	// ServerShare:  Specifies a DFS root target.
	ServerShare string `idl:"name:ServerShare;string;pointer:unique" json:"server_share"`
}

func (o *RootListEntry) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *RootListEntry) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.ServerShare != "" {
		_ptr_ServerShare := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.ServerShare); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.ServerShare, _ptr_ServerShare); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *RootListEntry) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_ServerShare := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.ServerShare); err != nil {
			return err
		}
		return nil
	})
	_s_ServerShare := func(ptr interface{}) { o.ServerShare = *ptr.(*string) }
	if err := w.ReadPointer(&o.ServerShare, _s_ServerShare, _ptr_ServerShare); err != nil {
		return err
	}
	return nil
}

// RootList structure represents DFSM_ROOT_LIST RPC structure.
//
// The DFSM_ROOT_LIST structure relates to the NetrDfsAdd2, NetrDfsAddFtRoot, and NetrDfsSetInfo2
// methods when used to add a DFS link or a DFS root target, or to modify the configuration
// of a domain-based DFS namespace. The structure contains an array of DFSM_ROOT_LIST_ENTRY
// structures, each of which contains information about a DFS root target.
//
// The DFSM_ROOT_LIST structure has the following format.
type RootList struct {
	// cEntries:  The number of DFS targets. The value of this member indicates the size
	// of the array in the Entry member.
	EntriesCount uint32 `idl:"name:cEntries" json:"entries_count"`
	// Entry:  An array of DFSM_ROOT_LIST_ENTRY structures. Each structure provides information
	// about one DFS target. For more information, see section 2.2.2.10.
	Entry []*RootListEntry `idl:"name:Entry;size_is:(cEntries)" json:"entry"`
}

func (o *RootList) xxx_PreparePayload(ctx context.Context) error {
	if o.Entry != nil && o.EntriesCount == 0 {
		o.EntriesCount = uint32(len(o.Entry))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *RootList) NDRSizeInfo() []uint64 {
	dimSize1 := uint64(o.EntriesCount)
	return []uint64{
		dimSize1,
	}
}
func (o *RootList) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
	for i1 := range o.Entry {
		i1 := i1
		if uint64(i1) >= sizeInfo[0] {
			break
		}
		if o.Entry[i1] != nil {
			if err := o.Entry[i1].MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&RootListEntry{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	for i1 := len(o.Entry); uint64(i1) < sizeInfo[0]; i1++ {
		if err := (&RootListEntry{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *RootList) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		return fmt.Errorf("buffer overflow for size %d of array o.Entry", sizeInfo[0])
	}
	o.Entry = make([]*RootListEntry, sizeInfo[0])
	for i1 := range o.Entry {
		i1 := i1
		if o.Entry[i1] == nil {
			o.Entry[i1] = &RootListEntry{}
		}
		if err := o.Entry[i1].UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

// NamespaceVersionOrigin type represents DFS_NAMESPACE_VERSION_ORIGIN RPC enumeration.
//
// The DFS_NAMESPACE_VERSION_ORIGIN is an enumeration that relates to the NetrDfsGetSupportedNamespaceVersion
// method when used to determine the supported DFS metadata version number.
//
// The DFS_NAMESPACE_VERSION_ORIGIN enumeration has the following format.
type NamespaceVersionOrigin uint16

var (
	// DFS_NAMESPACE_VERSION_ORIGIN_COMBINED:  This value is not used in communication.
	NamespaceVersionOriginCombined NamespaceVersionOrigin = 0
	// DFS_NAMESPACE_VERSION_ORIGIN_SERVER:  The maximum version that a server can support.
	NamespaceVersionOriginServer NamespaceVersionOrigin = 1
	// DFS_NAMESPACE_VERSION_ORIGIN_DOMAIN:  The maximum version that the domain can support.
	NamespaceVersionOriginDomain NamespaceVersionOrigin = 2
)

func (o NamespaceVersionOrigin) String() string {
	switch o {
	case NamespaceVersionOriginCombined:
		return "NamespaceVersionOriginCombined"
	case NamespaceVersionOriginServer:
		return "NamespaceVersionOriginServer"
	case NamespaceVersionOriginDomain:
		return "NamespaceVersionOriginDomain"
	}
	return "Invalid"
}

// SupportedNamespaceVersionInfo structure represents DFS_SUPPORTED_NAMESPACE_VERSION_INFO RPC structure.
//
// The DFS_SUPPORTED_NAMESPACE_VERSION_INFO structure relates to the NetrDfsGetSupportedNamespaceVersion
// method when used to determine the domain-based or standalone-based DFS major and
// minor version information.
//
// The DFS_SUPPORTED_NAMESPACE_VERSION_INFO structure has the following format.
type SupportedNamespaceVersionInfo struct {
	// DomainDfsMajorVersion:  A value containing the major version number of the DFS metadata
	// format supported by a domain-based DFS namespace.
	DomainDFSMajorVersion uint32 `idl:"name:DomainDfsMajorVersion" json:"domain_dfs_major_version"`
	// DomainDfsMinorVersion:  A value containing the minor version number of the DFS metadata
	// format supported by a domain-based DFS namespace.
	DomainDFSMinorVersion uint32 `idl:"name:DomainDfsMinorVersion" json:"domain_dfs_minor_version"`
	// DomainDfsCapabilities:  A value containing the capability information of a domain-based
	// DFS namespace.
	DomainDFSCapabilities uint64 `idl:"name:DomainDfsCapabilities" json:"domain_dfs_capabilities"`
	// StandaloneDfsMajorVersion:  A value containing the major version number of a stand-alone
	// DFS namespace.
	StandaloneDFSMajorVersion uint32 `idl:"name:StandaloneDfsMajorVersion" json:"standalone_dfs_major_version"`
	// StandaloneDfsMinorVersion:  A value containing the minor version number of a stand-alone
	// DFS namespace.
	StandaloneDFSMinorVersion uint32 `idl:"name:StandaloneDfsMinorVersion" json:"standalone_dfs_minor_version"`
	// StandaloneDfsCapabilities:  A value containing the capability information of a stand-alone
	// DFS namespace.
	//
	// DomainDfsCapabilities and StandaloneDfsCapabilities are bit fields with the following
	// defined value.
	//
	//	+--------------------------------------------------+-------------------------------------------------------------------------------+
	//	|                                                  |                                                                               |
	//	|                      VALUE                       |                                    MEANING                                    |
	//	|                                                  |                                                                               |
	//	+--------------------------------------------------+-------------------------------------------------------------------------------+
	//	+--------------------------------------------------+-------------------------------------------------------------------------------+
	//	| DFS_NAMESPACE_CAPABILITY_ABDE 0x0000000000000001 | This specifies support for Access Based Directory Enumeration (ABDE) mode.<4> |
	//	+--------------------------------------------------+-------------------------------------------------------------------------------+
	//
	// When this structure is used for communication, all undefined bit fields MUST be set
	// to zero. A client SHOULD ignore all bit fields it does not understand.
	StandaloneDFSCapabilities uint64 `idl:"name:StandaloneDfsCapabilities" json:"standalone_dfs_capabilities"`
}

func (o *SupportedNamespaceVersionInfo) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *SupportedNamespaceVersionInfo) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.DomainDFSMajorVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.DomainDFSMinorVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.DomainDFSCapabilities); err != nil {
		return err
	}
	if err := w.WriteData(o.StandaloneDFSMajorVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.StandaloneDFSMinorVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.StandaloneDFSCapabilities); err != nil {
		return err
	}
	return nil
}
func (o *SupportedNamespaceVersionInfo) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.DomainDFSMajorVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.DomainDFSMinorVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.DomainDFSCapabilities); err != nil {
		return err
	}
	if err := w.ReadData(&o.StandaloneDFSMajorVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.StandaloneDFSMinorVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.StandaloneDFSCapabilities); err != nil {
		return err
	}
	return nil
}

// Info1 structure represents DFS_INFO_1 RPC structure.
//
// The DFS_INFO_1 structure contains the name of a DFS root or DFS link.
//
// The DFS_INFO_1 structure has the following format.
type Info1 struct {
	// EntryPath:  The pointer to a DFS root or a DFS link path.
	EntryPath string `idl:"name:EntryPath;string" json:"entry_path"`
}

func (o *Info1) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.EntryPath != "" {
		_ptr_EntryPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EntryPath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EntryPath, _ptr_EntryPath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_EntryPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EntryPath); err != nil {
			return err
		}
		return nil
	})
	_s_EntryPath := func(ptr interface{}) { o.EntryPath = *ptr.(*string) }
	if err := w.ReadPointer(&o.EntryPath, _s_EntryPath, _ptr_EntryPath); err != nil {
		return err
	}
	return nil
}

// Info2 structure represents DFS_INFO_2 RPC structure.
//
// The DFS_INFO_2 structure contains information for a DFS root or DFS link.
//
// The DFS_INFO_2 structure has the following format.
type Info2 struct {
	// EntryPath:  A pointer to a DFS root or a DFS link path.
	EntryPath string `idl:"name:EntryPath;string" json:"entry_path"`
	// Comment:  A pointer to a null-terminated Unicode string containing a comment that
	// is used for informational purposes and is associated with the DFS root or DFS link.
	// This string has no protocol-specified restrictions on length or content. The comment
	// is meant for human consumption and does not affect server functionality.
	Comment string `idl:"name:Comment;string" json:"comment"`
	// State:   This field has the state of the DFS root or DFS link. For a DFS root, this
	// field also specifies whether the DFS namespace is stand-alone or domain-based.
	//
	// The DFS_VOLUME_STATES bitmask (0x0000000F) MUST be used to extract the following
	// DFS root or DFS link state from this field. For more information about some of these
	// states, see section 2.2.2.13.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|                                     |                                                                                  |
	//	|                VALUE                |                                     MEANING                                      |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_VOLUME_STATE_OK 0x00000001      | The specified DFS root or DFS link is in the normal state.                       |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| RESERVED 0x00000002                 | This value is reserved and MUST NOT be used.                                     |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_VOLUME_STATE_OFFLINE 0x00000003 | The DFS link is offline, and none of the DFS targets will be included in the     |
	//	|                                     | referral response. This flag is valid only for a DFS link and cannot be set on a |
	//	|                                     | DFS root. This state is persisted to the DFS metadata.                           |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_VOLUME_STATE_ONLINE 0x00000004  | The DFS link is online and available for referral request. This flag is valid    |
	//	|                                     | only for a DFS link and cannot be set on a DFS root. This state is persisted to  |
	//	|                                     | the DFS metadata.                                                                |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//
	// The DFS_VOLUME_FLAVORS bitmask (0x00000300) MUST be used to extract the following
	// DFS namespace flavor from this field.
	//
	//	+-----------------------------------------+-------------------------------------------------+
	//	|                                         |                                                 |
	//	|                  VALUE                  |                     MEANING                     |
	//	|                                         |                                                 |
	//	+-----------------------------------------+-------------------------------------------------+
	//	+-----------------------------------------+-------------------------------------------------+
	//	| DFS_VOLUME_FLAVOR_STANDALONE 0x00000100 | Stand-alone DFS namespace.                      |
	//	+-----------------------------------------+-------------------------------------------------+
	//	| DFS_VOLUME_FLAVOR_AD_BLOB 0x00000200    | domainv1-based or domainv2-based DFS namespace. |
	//	+-----------------------------------------+-------------------------------------------------+
	State uint32 `idl:"name:State" json:"state"`
	// NumberOfStorages:  Number of DFS targets for the root or link.
	NumberOfStorages uint32 `idl:"name:NumberOfStorages" json:"number_of_storages"`
}

func (o *Info2) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.EntryPath != "" {
		_ptr_EntryPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EntryPath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EntryPath, _ptr_EntryPath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Comment != "" {
		_ptr_Comment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Comment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Comment, _ptr_Comment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfStorages); err != nil {
		return err
	}
	return nil
}
func (o *Info2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_EntryPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EntryPath); err != nil {
			return err
		}
		return nil
	})
	_s_EntryPath := func(ptr interface{}) { o.EntryPath = *ptr.(*string) }
	if err := w.ReadPointer(&o.EntryPath, _s_EntryPath, _ptr_EntryPath); err != nil {
		return err
	}
	_ptr_Comment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Comment); err != nil {
			return err
		}
		return nil
	})
	_s_Comment := func(ptr interface{}) { o.Comment = *ptr.(*string) }
	if err := w.ReadPointer(&o.Comment, _s_Comment, _ptr_Comment); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfStorages); err != nil {
		return err
	}
	return nil
}

// Info3 structure represents DFS_INFO_3 RPC structure.
//
// The DFS_INFO_3 structure contains information for a DFS root or a DFS link.
//
// The DFS_INFO_3 structure has the following format.
type Info3 struct {
	// EntryPath:  Pointer to a DFS root or DFS link path.
	EntryPath string `idl:"name:EntryPath;string" json:"entry_path"`
	// Comment:  A pointer to a null-terminated Unicode string containing a comment associated
	// with the DFS root or DFS link that is for informational purposes. This string has
	// no protocol-specified restrictions on length or content. The comment is meant for
	// human consumption and does not affect server functionality.
	Comment string `idl:"name:Comment;string" json:"comment"`
	// State:  Refers to the State field of DFS_INFO_2. For more information, see section
	// 2.2.3.2.
	State uint32 `idl:"name:State" json:"state"`
	// NumberOfStorages:  The number of DFS targets for this root or link.
	NumberOfStorages uint32 `idl:"name:NumberOfStorages" json:"number_of_storages"`
	// Storage:  A pointer to an array of DFS_STORAGE_INFO structures containing information
	// about each target. (For more information, see section 2.2.2.5). The NumberOfStorages
	// member specifies the number of structures within this storage array.
	Storage []*StorageInfo `idl:"name:Storage;size_is:(NumberOfStorages)" json:"storage"`
}

func (o *Info3) xxx_PreparePayload(ctx context.Context) error {
	if o.Storage != nil && o.NumberOfStorages == 0 {
		o.NumberOfStorages = uint32(len(o.Storage))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.EntryPath != "" {
		_ptr_EntryPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EntryPath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EntryPath, _ptr_EntryPath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Comment != "" {
		_ptr_Comment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Comment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Comment, _ptr_Comment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfStorages); err != nil {
		return err
	}
	if o.Storage != nil || o.NumberOfStorages > 0 {
		_ptr_Storage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.NumberOfStorages)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Storage {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Storage[i1] != nil {
					if err := o.Storage[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&StorageInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Storage); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&StorageInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Storage, _ptr_Storage); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_EntryPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EntryPath); err != nil {
			return err
		}
		return nil
	})
	_s_EntryPath := func(ptr interface{}) { o.EntryPath = *ptr.(*string) }
	if err := w.ReadPointer(&o.EntryPath, _s_EntryPath, _ptr_EntryPath); err != nil {
		return err
	}
	_ptr_Comment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Comment); err != nil {
			return err
		}
		return nil
	})
	_s_Comment := func(ptr interface{}) { o.Comment = *ptr.(*string) }
	if err := w.ReadPointer(&o.Comment, _s_Comment, _ptr_Comment); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfStorages); err != nil {
		return err
	}
	_ptr_Storage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.NumberOfStorages > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.NumberOfStorages)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Storage", sizeInfo[0])
		}
		o.Storage = make([]*StorageInfo, sizeInfo[0])
		for i1 := range o.Storage {
			i1 := i1
			if o.Storage[i1] == nil {
				o.Storage[i1] = &StorageInfo{}
			}
			if err := o.Storage[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Storage := func(ptr interface{}) { o.Storage = *ptr.(*[]*StorageInfo) }
	if err := w.ReadPointer(&o.Storage, _s_Storage, _ptr_Storage); err != nil {
		return err
	}
	return nil
}

// Info4 structure represents DFS_INFO_4 RPC structure.
//
// The DFS_INFO_4 structure contains information for a DFS root or a DFS link.
//
// The DFS_INFO_4 structure has the following format.
type Info4 struct {
	// EntryPath:  A pointer to a DFS root or a DFS link path.
	EntryPath string `idl:"name:EntryPath;string" json:"entry_path"`
	// Comment:  A pointer to a null-terminated Unicode string containing a comment associated
	// with the DFS root or DFS link that is for informational purposes. This string has
	// no protocol-specified restrictions on length or content. The comment is meant for
	// human consumption and does not affect server functionality.
	Comment string `idl:"name:Comment;string" json:"comment"`
	// State:  Refers to the State field of DFS_INFO_2. For more information, see section
	// 2.2.3.2.
	State uint32 `idl:"name:State" json:"state"`
	// Timeout:  The time-out, in seconds, associated with the root or link and used in
	// a DFS referral response to a DFS client.
	Timeout uint32 `idl:"name:Timeout" json:"timeout"`
	// Guid:  The GUID of this root or link.
	GUID *dtyp.GUID `idl:"name:Guid" json:"guid"`
	// NumberOfStorages:  The number of DFS targets for this root or link. There are no
	// protocol-specified restrictions on the number of targets for a root or link.
	NumberOfStorages uint32 `idl:"name:NumberOfStorages" json:"number_of_storages"`
	// Storage:  A pointer to an array of DFS_STORAGE_INFO structures containing information
	// about each target. (For more information, see section 2.2.2.5). The NumberOfStorages
	// member specifies the number of structures within this storage array.
	Storage []*StorageInfo `idl:"name:Storage;size_is:(NumberOfStorages)" json:"storage"`
}

func (o *Info4) xxx_PreparePayload(ctx context.Context) error {
	if o.Storage != nil && o.NumberOfStorages == 0 {
		o.NumberOfStorages = uint32(len(o.Storage))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info4) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.EntryPath != "" {
		_ptr_EntryPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EntryPath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EntryPath, _ptr_EntryPath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Comment != "" {
		_ptr_Comment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Comment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Comment, _ptr_Comment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.Timeout); err != nil {
		return err
	}
	if o.GUID != nil {
		if err := o.GUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.NumberOfStorages); err != nil {
		return err
	}
	if o.Storage != nil || o.NumberOfStorages > 0 {
		_ptr_Storage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.NumberOfStorages)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Storage {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Storage[i1] != nil {
					if err := o.Storage[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&StorageInfo{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Storage); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&StorageInfo{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Storage, _ptr_Storage); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info4) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_EntryPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EntryPath); err != nil {
			return err
		}
		return nil
	})
	_s_EntryPath := func(ptr interface{}) { o.EntryPath = *ptr.(*string) }
	if err := w.ReadPointer(&o.EntryPath, _s_EntryPath, _ptr_EntryPath); err != nil {
		return err
	}
	_ptr_Comment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Comment); err != nil {
			return err
		}
		return nil
	})
	_s_Comment := func(ptr interface{}) { o.Comment = *ptr.(*string) }
	if err := w.ReadPointer(&o.Comment, _s_Comment, _ptr_Comment); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.Timeout); err != nil {
		return err
	}
	if o.GUID == nil {
		o.GUID = &dtyp.GUID{}
	}
	if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfStorages); err != nil {
		return err
	}
	_ptr_Storage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.NumberOfStorages > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.NumberOfStorages)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Storage", sizeInfo[0])
		}
		o.Storage = make([]*StorageInfo, sizeInfo[0])
		for i1 := range o.Storage {
			i1 := i1
			if o.Storage[i1] == nil {
				o.Storage[i1] = &StorageInfo{}
			}
			if err := o.Storage[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Storage := func(ptr interface{}) { o.Storage = *ptr.(*[]*StorageInfo) }
	if err := w.ReadPointer(&o.Storage, _s_Storage, _ptr_Storage); err != nil {
		return err
	}
	return nil
}

// Info5 structure represents DFS_INFO_5 RPC structure.
//
// The DFS_INFO_5 structure contains information for a DFS root or a DFS link.
//
// The DFS_INFO_5 structure has the following format.
type Info5 struct {
	// EntryPath:   A pointer to a DFS root or a DFS link path.
	EntryPath string `idl:"name:EntryPath;string" json:"entry_path"`
	// Comment:   A pointer to a null-terminated Unicode string containing a comment associated
	// with the DFS root or DFS link that is for informational purposes. This string has
	// no protocol-specified restrictions on length or content. The comment is meant for
	// human consumption and does not affect server functionality.
	Comment string `idl:"name:Comment;string" json:"comment"`
	// State:  Refers to the State field of DFS_INFO_2. For more information, see section
	// 2.2.3.2.
	State uint32 `idl:"name:State" json:"state"`
	// Timeout:   The time-out, in seconds, associated with the root or link and used in
	// a DFS referral response to a DFS client.
	Timeout uint32 `idl:"name:Timeout" json:"timeout"`
	// Guid:  The GUID of this root or link.
	GUID *dtyp.GUID `idl:"name:Guid" json:"guid"`
	// PropertyFlags:   A bit field in which each bit is responsible for a specific property
	// applicable to the entire DFS namespace, the DFS root, or an individual DFS link,
	// depending on the actual property. Any combination of bits is allowed, unless indicated
	// otherwise. The following are valid bit definitions for this field.
	//
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                               |                                                                                  |
	//	|                     VALUE                     |                                     MEANING                                      |
	//	|                                               |                                                                                  |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_PROPERTY_FLAG_INSITE_REFERRALS 0x00000001 | When set, indicates that DFS in-site referral mode is enabled.                   |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_PROPERTY_FLAG_ROOT_SCALABILITY 0x00000002 | When set, indicates DFS root scalability mode is enabled. This flag is valid     |
	//	|                                               | only for the DFS root of a domain-based DFS namespace.                           |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_PROPERTY_FLAG_SITE_COSTING 0x00000004     | When set, indicates DFS referral site costing is enabled. This flag is valid     |
	//	|                                               | only for a DFS root.                                                             |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_PROPERTY_FLAG_TARGET_FAILBACK 0x00000008  | When set, indicates DFS client target failback is enabled.                       |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_PROPERTY_FLAG_CLUSTER_ENABLED 0x00000010  | When set, indicates clustered DFS namespace is enabled.                          |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_PROPERTY_FLAG_ABDE 0x00000020             | When set, enables Access Based Directory Enumeration (ABDE) mode on a            |
	//	|                                               | domainv2-based DFS namespace or a stand-alone DFS namespace.<5>This flag is not  |
	//	|                                               | supported on domainv1-based namespaces.                                          |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	PropertyFlags uint32 `idl:"name:PropertyFlags" json:"property_flags"`
	// MetadataSize:  The size, in bytes, of the DFS metadata of the DFS namespace. For
	// a DFS link, this MUST be 0.
	MetadataSize uint32 `idl:"name:MetadataSize" json:"metadata_size"`
	// NumberOfStorages:  The number of DFS targets for this root or link.
	NumberOfStorages uint32 `idl:"name:NumberOfStorages" json:"number_of_storages"`
}

func (o *Info5) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info5) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.EntryPath != "" {
		_ptr_EntryPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EntryPath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EntryPath, _ptr_EntryPath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Comment != "" {
		_ptr_Comment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Comment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Comment, _ptr_Comment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.Timeout); err != nil {
		return err
	}
	if o.GUID != nil {
		if err := o.GUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PropertyFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.MetadataSize); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfStorages); err != nil {
		return err
	}
	return nil
}
func (o *Info5) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_EntryPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EntryPath); err != nil {
			return err
		}
		return nil
	})
	_s_EntryPath := func(ptr interface{}) { o.EntryPath = *ptr.(*string) }
	if err := w.ReadPointer(&o.EntryPath, _s_EntryPath, _ptr_EntryPath); err != nil {
		return err
	}
	_ptr_Comment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Comment); err != nil {
			return err
		}
		return nil
	})
	_s_Comment := func(ptr interface{}) { o.Comment = *ptr.(*string) }
	if err := w.ReadPointer(&o.Comment, _s_Comment, _ptr_Comment); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.Timeout); err != nil {
		return err
	}
	if o.GUID == nil {
		o.GUID = &dtyp.GUID{}
	}
	if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.MetadataSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfStorages); err != nil {
		return err
	}
	return nil
}

// Info6 structure represents DFS_INFO_6 RPC structure.
//
// The DFS_INFO_6 structure contains information for a DFS root or a DFS link.
//
// The DFS_INFO_6 structure has the following format.
type Info6 struct {
	// EntryPath:   A pointer to a DFS root or a DFS link path.
	EntryPath string `idl:"name:EntryPath;string" json:"entry_path"`
	// Comment:   A pointer to a null-terminated Unicode string containing a comment associated
	// with the DFS root or DFS link that is for informational purposes. This string has
	// no protocol-specified restrictions on length or content. The comment is meant for
	// human consumption and does not affect server functionality.
	Comment string `idl:"name:Comment;string" json:"comment"`
	// State:  Refers to the State field of DFS_INFO_2. For more information, see section
	// 2.2.3.2.
	State uint32 `idl:"name:State" json:"state"`
	// Timeout:   The time-out, in seconds, associated with the root or link and used in
	// a DFS referral response to a DFS client.
	Timeout uint32 `idl:"name:Timeout" json:"timeout"`
	// Guid:   The GUID of this root or link.
	GUID *dtyp.GUID `idl:"name:Guid" json:"guid"`
	// PropertyFlags:   Refers to the PropertyFlags field of DFS_INFO_5. For more information,
	// see section 2.2.3.5.
	PropertyFlags uint32 `idl:"name:PropertyFlags" json:"property_flags"`
	// MetadataSize:  The size of the DFS metadata of the DFS namespace. This MUST be 0
	// for a DFS link.
	MetadataSize uint32 `idl:"name:MetadataSize" json:"metadata_size"`
	// NumberOfStorages:   The number of DFS targets for this root or link. The protocol
	// imposes no restrictions on the number of roots or links.
	NumberOfStorages uint32 `idl:"name:NumberOfStorages" json:"number_of_storages"`
	// Storage:  A pointer to an array of DFS_STORAGE_INFO_1 structures containing information
	// about each target. The NumberOfStorages member specifies the number of structures
	// within this storage array.
	Storage []*StorageInfo1 `idl:"name:Storage;size_is:(NumberOfStorages)" json:"storage"`
}

func (o *Info6) xxx_PreparePayload(ctx context.Context) error {
	if o.Storage != nil && o.NumberOfStorages == 0 {
		o.NumberOfStorages = uint32(len(o.Storage))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.EntryPath != "" {
		_ptr_EntryPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EntryPath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EntryPath, _ptr_EntryPath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Comment != "" {
		_ptr_Comment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Comment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Comment, _ptr_Comment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.Timeout); err != nil {
		return err
	}
	if o.GUID != nil {
		if err := o.GUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PropertyFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.MetadataSize); err != nil {
		return err
	}
	if err := w.WriteData(o.NumberOfStorages); err != nil {
		return err
	}
	if o.Storage != nil || o.NumberOfStorages > 0 {
		_ptr_Storage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.NumberOfStorages)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Storage {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Storage[i1] != nil {
					if err := o.Storage[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&StorageInfo1{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Storage); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&StorageInfo1{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Storage, _ptr_Storage); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_EntryPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EntryPath); err != nil {
			return err
		}
		return nil
	})
	_s_EntryPath := func(ptr interface{}) { o.EntryPath = *ptr.(*string) }
	if err := w.ReadPointer(&o.EntryPath, _s_EntryPath, _ptr_EntryPath); err != nil {
		return err
	}
	_ptr_Comment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Comment); err != nil {
			return err
		}
		return nil
	})
	_s_Comment := func(ptr interface{}) { o.Comment = *ptr.(*string) }
	if err := w.ReadPointer(&o.Comment, _s_Comment, _ptr_Comment); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.Timeout); err != nil {
		return err
	}
	if o.GUID == nil {
		o.GUID = &dtyp.GUID{}
	}
	if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.MetadataSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfStorages); err != nil {
		return err
	}
	_ptr_Storage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.NumberOfStorages > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.NumberOfStorages)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Storage", sizeInfo[0])
		}
		o.Storage = make([]*StorageInfo1, sizeInfo[0])
		for i1 := range o.Storage {
			i1 := i1
			if o.Storage[i1] == nil {
				o.Storage[i1] = &StorageInfo1{}
			}
			if err := o.Storage[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Storage := func(ptr interface{}) { o.Storage = *ptr.(*[]*StorageInfo1) }
	if err := w.ReadPointer(&o.Storage, _s_Storage, _ptr_Storage); err != nil {
		return err
	}
	return nil
}

// Info7 structure represents DFS_INFO_7 RPC structure.
//
// The DFS_INFO_7 structure contains information about a DFS root.
//
// The DFS_INFO_7 structure has the following format.
type Info7 struct {
	// GenerationGuid:  This GUID is modified each time DFS metadata is updated.
	//
	// This data type is used to detect when the metadata of a DFS namespace has changed.
	// It MUST be supported for domain-based DFS namespaces. It MAY be supported for stand-alone
	// DFS namespaces; a null GUID (all 128-bits are 0) MUST be returned if this is not
	// supported.<6>
	GenerationGUID *dtyp.GUID `idl:"name:GenerationGuid" json:"generation_guid"`
}

func (o *Info7) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info7) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.GenerationGUID != nil {
		if err := o.GenerationGUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info7) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.GenerationGUID == nil {
		o.GenerationGUID = &dtyp.GUID{}
	}
	if err := o.GenerationGUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Info8 structure represents DFS_INFO_8 RPC structure.
//
// The DFS_INFO_8 structure contains information for a DFS root or a DFS link.
//
// The DFS_INFO_8 structure has the following format.
type Info8 struct {
	// EntryPath:  A pointer to a DFS root or a DFS link path.
	EntryPath string `idl:"name:EntryPath;string" json:"entry_path"`
	// Comment:  A pointer to a null-terminated Unicode string containing a comment associated
	// with the DFS root or DFS link that is for informational purposes. This string has
	// no protocol-specified restrictions on length or content. The comment is meant for
	// human consumption and does not affect server functionality.
	Comment string `idl:"name:Comment;string" json:"comment"`
	// State:  Refers to the State field of DFS_INFO_2. For more information, see section
	// 2.2.3.2.
	State uint32 `idl:"name:State" json:"state"`
	// Timeout:  The time-out, in seconds, associated with the root or link and used in
	// a DFS referral response to a DFS client.
	Timeout uint32 `idl:"name:Timeout" json:"timeout"`
	// Guid:  The GUID of this root or link.
	GUID *dtyp.GUID `idl:"name:Guid" json:"guid"`
	// PropertyFlags:  Refers to the PropertyFlags field of DFS_INFO_5. For more information,
	// see section 2.2.3.5.
	PropertyFlags uint32 `idl:"name:PropertyFlags" json:"property_flags"`
	// MetadataSize:  The size, in bytes, of the DFS metadata of the DFS namespace. For
	// a DFS link, this MUST be 0.
	MetadataSize uint32 `idl:"name:MetadataSize" json:"metadata_size"`
	// SecurityDescriptorLength:  The length, in bytes, of the buffer that the pSecurityDescriptor
	// field points to.
	SecurityDescriptorLength uint32 `idl:"name:SecurityDescriptorLength" json:"security_descriptor_length"`
	// pSecurityDescriptor:  A self-relative security descriptor to be associated with a
	// DFS link. For more information on security descriptors, see [MS-DTYP] section 2.4.6.
	SecurityDescriptor []byte `idl:"name:pSecurityDescriptor;size_is:(SecurityDescriptorLength)" json:"security_descriptor"`
	// NumberOfStorages:  The number of DFS targets for this root or link. The protocol
	// imposes no restrictions on the number of roots or links.
	NumberOfStorages uint32 `idl:"name:NumberOfStorages" json:"number_of_storages"`
}

func (o *Info8) xxx_PreparePayload(ctx context.Context) error {
	if o.SecurityDescriptor != nil && o.SecurityDescriptorLength == 0 {
		o.SecurityDescriptorLength = uint32(len(o.SecurityDescriptor))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info8) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.EntryPath != "" {
		_ptr_EntryPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EntryPath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EntryPath, _ptr_EntryPath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Comment != "" {
		_ptr_Comment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Comment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Comment, _ptr_Comment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.Timeout); err != nil {
		return err
	}
	if o.GUID != nil {
		if err := o.GUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PropertyFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.MetadataSize); err != nil {
		return err
	}
	if err := w.WriteData(o.SecurityDescriptorLength); err != nil {
		return err
	}
	if o.SecurityDescriptor != nil || o.SecurityDescriptorLength > 0 {
		_ptr_pSecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.SecurityDescriptorLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
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
		if err := w.WritePointer(&o.SecurityDescriptor, _ptr_pSecurityDescriptor); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.NumberOfStorages); err != nil {
		return err
	}
	return nil
}
func (o *Info8) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_EntryPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EntryPath); err != nil {
			return err
		}
		return nil
	})
	_s_EntryPath := func(ptr interface{}) { o.EntryPath = *ptr.(*string) }
	if err := w.ReadPointer(&o.EntryPath, _s_EntryPath, _ptr_EntryPath); err != nil {
		return err
	}
	_ptr_Comment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Comment); err != nil {
			return err
		}
		return nil
	})
	_s_Comment := func(ptr interface{}) { o.Comment = *ptr.(*string) }
	if err := w.ReadPointer(&o.Comment, _s_Comment, _ptr_Comment); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.Timeout); err != nil {
		return err
	}
	if o.GUID == nil {
		o.GUID = &dtyp.GUID{}
	}
	if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.MetadataSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecurityDescriptorLength); err != nil {
		return err
	}
	_ptr_pSecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.SecurityDescriptorLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.SecurityDescriptorLength)
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
	_s_pSecurityDescriptor := func(ptr interface{}) { o.SecurityDescriptor = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.SecurityDescriptor, _s_pSecurityDescriptor, _ptr_pSecurityDescriptor); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfStorages); err != nil {
		return err
	}
	return nil
}

// Info9 structure represents DFS_INFO_9 RPC structure.
//
// The DFS_INFO_9 structure contains information for a DFS root or a DFS link.
//
// The DFS_INFO_9 structure has the following format.
type Info9 struct {
	// EntryPath:  A pointer to a DFS root or a DFS link path.
	EntryPath string `idl:"name:EntryPath;string" json:"entry_path"`
	// Comment:  Pointer to a null-terminated Unicode string containing a comment associated
	// with the DFS root or DFS link that is for informational purposes. There are no protocol-specified
	// restrictions on the length or content of this string. The comment is meant for human
	// readability and has no effect on server functionality.
	Comment string `idl:"name:Comment;string" json:"comment"`
	// State:  Refers to the State field of DFS_INFO_2. For more information, see section
	// 2.2.3.2.
	State uint32 `idl:"name:State" json:"state"`
	// Timeout:  The time-out, in seconds, associated with the root or link and used in
	// a DFS referral response to a DFS client.
	Timeout uint32 `idl:"name:Timeout" json:"timeout"`
	// Guid:  The GUID of this root or link.
	GUID *dtyp.GUID `idl:"name:Guid" json:"guid"`
	// PropertyFlags:  Refers to the PropertyFlags field of DFS_INFO_5. For more information,
	// see section 2.2.3.5.
	PropertyFlags uint32 `idl:"name:PropertyFlags" json:"property_flags"`
	// MetadataSize:  The size, in bytes, of the DFS metadata of the DFS namespace. For
	// a DFS link, this MUST be 0.
	MetadataSize uint32 `idl:"name:MetadataSize" json:"metadata_size"`
	// SecurityDescriptorLength:  The length, in bytes, of the buffer that the pSecurityDescriptor
	// field points to.
	SecurityDescriptorLength uint32 `idl:"name:SecurityDescriptorLength" json:"security_descriptor_length"`
	// pSecurityDescriptor:  A self-relative security descriptor to be associated with a
	// DFS link. For more information on security descriptors, see [MS-DTYP] section 2.4.6.
	SecurityDescriptor []byte `idl:"name:pSecurityDescriptor;size_is:(SecurityDescriptorLength)" json:"security_descriptor"`
	// NumberOfStorages:  The number of DFS targets for this root or link. The protocol
	// imposes no restrictions on the number of roots or links.
	NumberOfStorages uint32 `idl:"name:NumberOfStorages" json:"number_of_storages"`
	// Storage:  A pointer to an array of DFS_STORAGE_INFO_1 structures containing information
	// about each target. The NumberOfStorages member specifies the number of structures
	// within this storage array.
	//
	// For information on target priority rank and class information, see section 2.2.2.6.
	Storage []*StorageInfo1 `idl:"name:Storage;size_is:(NumberOfStorages)" json:"storage"`
}

func (o *Info9) xxx_PreparePayload(ctx context.Context) error {
	if o.SecurityDescriptor != nil && o.SecurityDescriptorLength == 0 {
		o.SecurityDescriptorLength = uint32(len(o.SecurityDescriptor))
	}
	if o.Storage != nil && o.NumberOfStorages == 0 {
		o.NumberOfStorages = uint32(len(o.Storage))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info9) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.EntryPath != "" {
		_ptr_EntryPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.EntryPath); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.EntryPath, _ptr_EntryPath); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if o.Comment != "" {
		_ptr_Comment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Comment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Comment, _ptr_Comment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.Timeout); err != nil {
		return err
	}
	if o.GUID != nil {
		if err := o.GUID.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.PropertyFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.MetadataSize); err != nil {
		return err
	}
	if err := w.WriteData(o.SecurityDescriptorLength); err != nil {
		return err
	}
	if o.SecurityDescriptor != nil || o.SecurityDescriptorLength > 0 {
		_ptr_pSecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.SecurityDescriptorLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
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
		if err := w.WritePointer(&o.SecurityDescriptor, _ptr_pSecurityDescriptor); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.NumberOfStorages); err != nil {
		return err
	}
	if o.Storage != nil || o.NumberOfStorages > 0 {
		_ptr_Storage := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.NumberOfStorages)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.Storage {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.Storage[i1] != nil {
					if err := o.Storage[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&StorageInfo1{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Storage); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&StorageInfo1{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Storage, _ptr_Storage); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info9) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_EntryPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.EntryPath); err != nil {
			return err
		}
		return nil
	})
	_s_EntryPath := func(ptr interface{}) { o.EntryPath = *ptr.(*string) }
	if err := w.ReadPointer(&o.EntryPath, _s_EntryPath, _ptr_EntryPath); err != nil {
		return err
	}
	_ptr_Comment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Comment); err != nil {
			return err
		}
		return nil
	})
	_s_Comment := func(ptr interface{}) { o.Comment = *ptr.(*string) }
	if err := w.ReadPointer(&o.Comment, _s_Comment, _ptr_Comment); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.Timeout); err != nil {
		return err
	}
	if o.GUID == nil {
		o.GUID = &dtyp.GUID{}
	}
	if err := o.GUID.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.MetadataSize); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecurityDescriptorLength); err != nil {
		return err
	}
	_ptr_pSecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.SecurityDescriptorLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.SecurityDescriptorLength)
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
	_s_pSecurityDescriptor := func(ptr interface{}) { o.SecurityDescriptor = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.SecurityDescriptor, _s_pSecurityDescriptor, _ptr_pSecurityDescriptor); err != nil {
		return err
	}
	if err := w.ReadData(&o.NumberOfStorages); err != nil {
		return err
	}
	_ptr_Storage := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.NumberOfStorages > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.NumberOfStorages)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Storage", sizeInfo[0])
		}
		o.Storage = make([]*StorageInfo1, sizeInfo[0])
		for i1 := range o.Storage {
			i1 := i1
			if o.Storage[i1] == nil {
				o.Storage[i1] = &StorageInfo1{}
			}
			if err := o.Storage[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Storage := func(ptr interface{}) { o.Storage = *ptr.(*[]*StorageInfo1) }
	if err := w.ReadPointer(&o.Storage, _s_Storage, _ptr_Storage); err != nil {
		return err
	}
	return nil
}

// Info50 structure represents DFS_INFO_50 RPC structure.
//
// The DFS_INFO_50 structure is used to get the DFS metadata version and the capability
// information of an existing DFS namespace.
//
// The DFS_INFO_50 structure has the following format.
type Info50 struct {
	// NamespaceMajorVersion:  A value containing the major version number used to determine
	// the DFS metadata format supported in a domain-based DFS namespace or a stand-alone
	// DFS namespace.<7>
	NamespaceMajorVersion uint32 `idl:"name:NamespaceMajorVersion" json:"namespace_major_version"`
	// NamespaceMinorVersion:  A value containing the minor version number used to determine
	// the DFS metadata format supported in a domain-based DFS namespace or stand-alone
	// DFS namespace.<8>
	NamespaceMinorVersion uint32 `idl:"name:NamespaceMinorVersion" json:"namespace_minor_version"`
	// NamespaceCapabilities:  A value containing the capability information of a DFS namespace.
	//
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                                   |                                                                                  |
	//	|                       VALUE                       |                                     MEANING                                      |
	//	|                                                   |                                                                                  |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS__NAMESPACE_CAPABILITY_ABDE 0x0000000000000001 | The specified DFS root supports using Access Based Directory Enumeration (ABDE)  |
	//	|                                                   | mode.<9>                                                                         |
	//	+---------------------------------------------------+----------------------------------------------------------------------------------+
	NamespaceCapabilities uint64 `idl:"name:NamespaceCapabilities" json:"namespace_capabilities"`
}

func (o *Info50) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info50) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(8); err != nil {
		return err
	}
	if err := w.WriteData(o.NamespaceMajorVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.NamespaceMinorVersion); err != nil {
		return err
	}
	if err := w.WriteData(o.NamespaceCapabilities); err != nil {
		return err
	}
	return nil
}
func (o *Info50) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(8); err != nil {
		return err
	}
	if err := w.ReadData(&o.NamespaceMajorVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.NamespaceMinorVersion); err != nil {
		return err
	}
	if err := w.ReadData(&o.NamespaceCapabilities); err != nil {
		return err
	}
	return nil
}

// Info100 structure represents DFS_INFO_100 RPC structure.
//
// The DFS_INFO_100 structure relates to the NetrDfsGetInfo, NetrDfsSetInfo, and NetrDfsSetInfo2
// methods when used to retrieve or set comment text about a DFS root or a DFS link.
// The structure contains a comment associated with a DFS root or a DFS link.
//
// The DFS_INFO_100 structure has the following format.
type Info100 struct {
	// Comment:   A pointer to a null-terminated Unicode string containing a comment associated
	// with the DFS root or DFS link that is for informational purposes. This string has
	// no protocol-specified restrictions on length or content. The comment is meant for
	// human readability and does not affect server functionality.
	Comment string `idl:"name:Comment;string" json:"comment"`
}

func (o *Info100) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info100) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.Comment != "" {
		_ptr_Comment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Comment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Comment, _ptr_Comment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info100) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_Comment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Comment); err != nil {
			return err
		}
		return nil
	})
	_s_Comment := func(ptr interface{}) { o.Comment = *ptr.(*string) }
	if err := w.ReadPointer(&o.Comment, _s_Comment, _ptr_Comment); err != nil {
		return err
	}
	return nil
}

// Info101 structure represents DFS_INFO_101 RPC structure.
//
// The DFS_INFO_101 structure describes the storage state on a root, link, root target,
// or link target.
//
// The DFS_INFO_101 structure has the following format.
type Info101 struct {
	// State:   The state of the root, link, root target, or link target.
	//
	// The following table lists the valid states that can be set for a root or a link.
	// Some of these states are used to perform a server operation and are not persisted
	// to the DFS metadata, as specified below. For more information about some of these
	// states, see section 2.2.2.13.
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                           |                                                                                  |
	//	|                   VALUE                   |                                     MEANING                                      |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_VOLUME_STATE_OK 0x00000001            | The specified DFS root or DFS link is in the normal state.                       |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_VOLUME_STATE_OFFLINE 0x00000003       | The specified DFS link is offline or unavailable. This flag is valid only for a  |
	//	|                                           | DFS link. This state is persisted to the DFS metadata.                           |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_VOLUME_STATE_ONLINE 0x00000004        | The specified DFS link is available. This flag is valid only for a DFS link.     |
	//	|                                           | This state is persisted to the DFS metadata.                                     |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_VOLUME_STATE_RESYNCHRONIZE 0x00000010 | Forces a resynchronization on the DFS root. This flag is valid only for a        |
	//	|                                           | DFS root. This operation is an incremental synchronization that picks up         |
	//	|                                           | only changed objects in the metadata. This state is used to perform a server     |
	//	|                                           | operation. It is not persisted to the DFS metadata.                              |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_VOLUME_STATE_STANDBY 0x00000020       | Sets a root volume to standby mode. This flag is valid only for a clustered DFS  |
	//	|                                           | root. This state is used to perform a server operation. It is not persisted to   |
	//	|                                           | the DFS metadata.                                                                |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_VOLUME_STATE_FORCE_SYNC 0x00000040    | Forces a full resynchronization operation on the DFS root target of a specified  |
	//	|                                           | domainv2-based DFS namespace or stand-alone DFS namespace to identify DFS links  |
	//	|                                           | that have been added or deleted. This is not supported on a domainv1-based DFS   |
	//	|                                           | namespace. DFS links MUST NOT be specified. This state is used to perform a      |
	//	|                                           | server operation. It is not persisted to the DFS metadata.                       |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//
	// DFS_VOLUME_STATES (0x0000000F) is not relevant here, because it is a mask used when
	// reading the volume state, not for setting it.
	//
	// The following table lists the valid states that can be set for a root target or a
	// link target.
	//
	//	+--------------------------------------+-------------------------------------------------+
	//	|                                      |                                                 |
	//	|                VALUE                 |                     MEANING                     |
	//	|                                      |                                                 |
	//	+--------------------------------------+-------------------------------------------------+
	//	+--------------------------------------+-------------------------------------------------+
	//	| DFS_STORAGE_STATE_OFFLINE 0x00000001 | This target is offline and unavailable for use. |
	//	+--------------------------------------+-------------------------------------------------+
	//	| DFS_STORAGE_STATE_ONLINE 0x00000002  | This target is online and available for use.    |
	//	+--------------------------------------+-------------------------------------------------+
	State uint32 `idl:"name:State" json:"state"`
}

func (o *Info101) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info101) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	return nil
}
func (o *Info101) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	return nil
}

// Info102 structure represents DFS_INFO_102 RPC structure.
//
// The DFS_INFO_102 structure contains a time-out value for a DFS root or a DFS link.
//
// The DFS_INFO_102 structure has the following format.
type Info102 struct {
	// Timeout:   The time-out, in seconds, associated with the root or link and used in
	// a DFS referral response to a DFS client.
	Timeout uint32 `idl:"name:Timeout" json:"timeout"`
}

func (o *Info102) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info102) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.Timeout); err != nil {
		return err
	}
	return nil
}
func (o *Info102) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.Timeout); err != nil {
		return err
	}
	return nil
}

// Info103 structure represents DFS_INFO_103 RPC structure.
//
// The DFS_INFO_103 structure contains properties that set specific behaviors for a
// DFS root or a DFS link.
//
// The DFS_INFO_103 structure has the following format.
type Info103 struct {
	// PropertyFlagMask:  Indicates which bits in the PropertyFlags field are considered
	// for modification of DFS namespace root or link properties.
	//
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                               |                                                                                  |
	//	|                     VALUE                     |                                     MEANING                                      |
	//	|                                               |                                                                                  |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_PROPERTY_FLAG_INSITE_REFERRALS 0x00000001 | Valid for domain and stand-alone DFS roots and links.                            |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_PROPERTY_FLAG_ROOT_SCALABILITY 0x00000002 | This flag is valid only for the DFS root of a domain-based DFS namespace.        |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_PROPERTY_FLAG_SITE_COSTING 0x00000004     | This flag is valid only for a DFS root.                                          |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_PROPERTY_FLAG_TARGET_FAILBACK 0x00000008  | Valid for domain and stand-alone DFS roots and links.                            |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_PROPERTY_FLAG_ABDE 0x00000020             | Valid only for a domainv2-based DFS namespace or stand-alone DFS namespace       |
	//	|                                               | root.<10> This flag is not supported on domainv1-based namespaces.               |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	PropertyFlagMask uint32 `idl:"name:PropertyFlagMask" json:"property_flag_mask"`
	// PropertyFlags:  A bit field in which each bit is responsible for a specific property
	// applicable to the whole DFS namespace, the DFS root, or an individual DFS link, depending
	// on the actual property. Any combination of bits is allowed, unless indicated otherwise.
	// The server considers the bits in this field only when the corresponding bit in the
	// PropertyFlagMask field is set. The following table lists the valid bits for this
	// field and describes the actions taken when each bit is set or not set in the request.
	//
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                               |                                                                                  |
	//	|                     VALUE                     |                                     MEANING                                      |
	//	|                                               |                                                                                  |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_PROPERTY_FLAG_INSITE_REFERRALS 0x00000001 | When set, enables DFS in-site referral mode. When not set, disables DFS in-site  |
	//	|                                               | referral mode.                                                                   |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_PROPERTY_FLAG_ROOT_SCALABILITY 0x00000002 | When set, enables DFS root scalability mode. When not set, disables DFS root     |
	//	|                                               | scalability mode.                                                                |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_PROPERTY_FLAG_SITE_COSTING 0x00000004     | When set, enables DFS referral site costing. When not set, disables DFS referral |
	//	|                                               | site costing.                                                                    |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_PROPERTY_FLAG_TARGET_FAILBACK 0x00000008  | When set, enables DFS client target failback. When not set, disables DFS client  |
	//	|                                               | target failback.                                                                 |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_PROPERTY_FLAG_ABDE 0x00000020             | When set, enables ABDE mode on a domainv2-based DFS namespace or stand-alone DFS |
	//	|                                               | namespace. When not set, disables ABDE mode on a domainv2-based DFS namespace or |
	//	|                                               | stand-alone DFS namespace.                                                       |
	//	+-----------------------------------------------+----------------------------------------------------------------------------------+
	PropertyFlags uint32 `idl:"name:PropertyFlags" json:"property_flags"`
}

func (o *Info103) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info103) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.PropertyFlagMask); err != nil {
		return err
	}
	if err := w.WriteData(o.PropertyFlags); err != nil {
		return err
	}
	return nil
}
func (o *Info103) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyFlagMask); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyFlags); err != nil {
		return err
	}
	return nil
}

// Info104 structure represents DFS_INFO_104 RPC structure.
//
// The DFS_INFO_104 structure contains the priority of a DFS root target or a DFS link
// target.
//
// The DFS_INFO_104 structure has the following format.
type Info104 struct {
	// TargetPriority:  A DFS_TARGET_PRIORITY structure that indicates the priority rank
	// and priority class of a target. For more information on prioritization, see section
	// 2.2.2.7.
	TargetPriority *TargetPriority `idl:"name:TargetPriority" json:"target_priority"`
}

func (o *Info104) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info104) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if o.TargetPriority != nil {
		if err := o.TargetPriority.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TargetPriority{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info104) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if o.TargetPriority == nil {
		o.TargetPriority = &TargetPriority{}
	}
	if err := o.TargetPriority.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Info105 structure represents DFS_INFO_105 RPC structure.
//
// The DFS_INFO_105 structure contains information about a DFS root or DFS link, including
// comment, state, time-out, and DFS behaviors specified by property flags.
//
// The DFS_INFO_105 structure has the following format.
type Info105 struct {
	// Comment:   A pointer to a null-terminated Unicode string containing a comment associated
	// with the DFS root or DFS link that is for informational purposes. This string has
	// no protocol-specified restrictions on length or content. The comment is meant for
	// human consumption and does not affect server functionality.
	Comment string `idl:"name:Comment;string" json:"comment"`
	// State:  The following table lists the valid states that can be set for links. All
	// other values are reserved and MUST NOT be used. For more information about some of
	// these states, see section 2.2.2.13.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|                                     |                                                                                  |
	//	|                VALUE                |                                     MEANING                                      |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000                          | Indicates that the existing state MUST NOT be changed.                           |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_VOLUME_STATE_OFFLINE 0x00000003 | The specified DFS link is offline or unavailable. This flag is valid only for a  |
	//	|                                     | DFS link.                                                                        |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_VOLUME_STATE_ONLINE 0x00000004  | The specified DFS link is available. This flag is valid only for a DFS link.     |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	State uint32 `idl:"name:State" json:"state"`
	// Timeout:   The time-out, in seconds, associated with the root or link and used in
	// a DFS referral response to a DFS client.
	Timeout uint32 `idl:"name:Timeout" json:"timeout"`
	// PropertyFlagMask:   Indicates which bits in the PropertyFlags field are valid.
	PropertyFlagMask uint32 `idl:"name:PropertyFlagMask" json:"property_flag_mask"`
	// PropertyFlags:  Refers to the PropertyFlags field of DFS_INFO_103, as specified in
	// section 2.2.4.3.
	PropertyFlags uint32 `idl:"name:PropertyFlags" json:"property_flags"`
}

func (o *Info105) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info105) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Comment != "" {
		_ptr_Comment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Comment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Comment, _ptr_Comment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.Timeout); err != nil {
		return err
	}
	if err := w.WriteData(o.PropertyFlagMask); err != nil {
		return err
	}
	if err := w.WriteData(o.PropertyFlags); err != nil {
		return err
	}
	return nil
}
func (o *Info105) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_Comment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Comment); err != nil {
			return err
		}
		return nil
	})
	_s_Comment := func(ptr interface{}) { o.Comment = *ptr.(*string) }
	if err := w.ReadPointer(&o.Comment, _s_Comment, _ptr_Comment); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.Timeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyFlagMask); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyFlags); err != nil {
		return err
	}
	return nil
}

// Info106 structure represents DFS_INFO_106 RPC structure.
//
// The DFS_INFO_106 structure contains the storage state and priority of a DFS root
// target or a DFS link target. For more information on prioritization, see section
// 2.2.2.7.
//
// The DFS_INFO_106 structure has the following format.
type Info106 struct {
	// State:  The state of the target. Contains one of the following valid state values.
	//
	//	+--------------------------------------+-------------------------------------------------+
	//	|                                      |                                                 |
	//	|                VALUE                 |                     MEANING                     |
	//	|                                      |                                                 |
	//	+--------------------------------------+-------------------------------------------------+
	//	+--------------------------------------+-------------------------------------------------+
	//	| DFS_STORAGE_STATE_OFFLINE 0x00000001 | This target is offline and unavailable for use. |
	//	+--------------------------------------+-------------------------------------------------+
	//	| DFS_STORAGE_STATE_ONLINE 0x00000002  | This target is online and available for use.    |
	//	+--------------------------------------+-------------------------------------------------+
	State uint32 `idl:"name:State" json:"state"`
	// TargetPriority:  A DFS_TARGET_PRIORITY structure that indicates the priority class
	// and rank of the DFS target.
	TargetPriority *TargetPriority `idl:"name:TargetPriority" json:"target_priority"`
}

func (o *Info106) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info106) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(4); err != nil {
		return err
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if o.TargetPriority != nil {
		if err := o.TargetPriority.MarshalNDR(ctx, w); err != nil {
			return err
		}
	} else {
		if err := (&TargetPriority{}).MarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info106) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(4); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if o.TargetPriority == nil {
		o.TargetPriority = &TargetPriority{}
	}
	if err := o.TargetPriority.UnmarshalNDR(ctx, w); err != nil {
		return err
	}
	return nil
}

// Info107 structure represents DFS_INFO_107 RPC structure.
//
// The DFS_INFO_107 structure contains information about a DFS root or DFS link, including
// comment, state, time-out, security descriptor, and DFS behaviors specified by property
// flags.
//
// The DFS_INFO_107 structure has the following format.
type Info107 struct {
	// Comment:  A pointer to a null-terminated Unicode string containing a comment associated
	// with the DFS root or DFS link that is for informational purposes. This string has
	// no protocol-specified restrictions on length or content. The comment is meant for
	// human readability and does not affect server functionality.
	Comment string `idl:"name:Comment;string" json:"comment"`
	// State:  The states that can be set for links. The following table lists such states.
	// All other values are reserved and MUST NOT be used. For more information about some
	// of these states, see section 2.2.2.13.
	//
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	|                                     |                                                                                  |
	//	|                VALUE                |                                     MEANING                                      |
	//	|                                     |                                                                                  |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000                          | Indicates that the existing state MUST NOT be changed.                           |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_VOLUME_STATE_OFFLINE 0x00000003 | The specified DFS link is offline or unavailable. This flag is valid only for a  |
	//	|                                     | DFS link.                                                                        |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_VOLUME_STATE_ONLINE 0x00000004  | The specified DFS link is available. This flag is valid only for a DFS link.     |
	//	+-------------------------------------+----------------------------------------------------------------------------------+
	State uint32 `idl:"name:State" json:"state"`
	// Timeout:  The time-out, in seconds, associated with the root or link and used in
	// a DFS referral response to a DFS client.
	Timeout uint32 `idl:"name:Timeout" json:"timeout"`
	// PropertyFlagMask:  Indicates which bits in the PropertyFlags field are valid.
	PropertyFlagMask uint32 `idl:"name:PropertyFlagMask" json:"property_flag_mask"`
	// PropertyFlags:  Refers to the PropertyFlags field of DFS_INFO_103, as specified in
	// section 2.2.4.3.
	PropertyFlags uint32 `idl:"name:PropertyFlags" json:"property_flags"`
	// SecurityDescriptorLength:  The length, in bytes, of the buffer that the pSecurityDescriptor
	// field points to.
	SecurityDescriptorLength uint32 `idl:"name:SecurityDescriptorLength" json:"security_descriptor_length"`
	// pSecurityDescriptor:  A self-relative security descriptor to be associated with a
	// DFS link. For more information on security descriptors, see [MS-DTYP] section 2.4.6.
	SecurityDescriptor []byte `idl:"name:pSecurityDescriptor;size_is:(SecurityDescriptorLength)" json:"security_descriptor"`
}

func (o *Info107) xxx_PreparePayload(ctx context.Context) error {
	if o.SecurityDescriptor != nil && o.SecurityDescriptorLength == 0 {
		o.SecurityDescriptorLength = uint32(len(o.SecurityDescriptor))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info107) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if o.Comment != "" {
		_ptr_Comment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Comment); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Comment, _ptr_Comment); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	if err := w.WriteData(o.State); err != nil {
		return err
	}
	if err := w.WriteData(o.Timeout); err != nil {
		return err
	}
	if err := w.WriteData(o.PropertyFlagMask); err != nil {
		return err
	}
	if err := w.WriteData(o.PropertyFlags); err != nil {
		return err
	}
	if err := w.WriteData(o.SecurityDescriptorLength); err != nil {
		return err
	}
	if o.SecurityDescriptor != nil || o.SecurityDescriptorLength > 0 {
		_ptr_pSecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.SecurityDescriptorLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
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
		if err := w.WritePointer(&o.SecurityDescriptor, _ptr_pSecurityDescriptor); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info107) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	_ptr_Comment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Comment); err != nil {
			return err
		}
		return nil
	})
	_s_Comment := func(ptr interface{}) { o.Comment = *ptr.(*string) }
	if err := w.ReadPointer(&o.Comment, _s_Comment, _ptr_Comment); err != nil {
		return err
	}
	if err := w.ReadData(&o.State); err != nil {
		return err
	}
	if err := w.ReadData(&o.Timeout); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyFlagMask); err != nil {
		return err
	}
	if err := w.ReadData(&o.PropertyFlags); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecurityDescriptorLength); err != nil {
		return err
	}
	_ptr_pSecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.SecurityDescriptorLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.SecurityDescriptorLength)
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
	_s_pSecurityDescriptor := func(ptr interface{}) { o.SecurityDescriptor = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.SecurityDescriptor, _s_pSecurityDescriptor, _ptr_pSecurityDescriptor); err != nil {
		return err
	}
	return nil
}

// Info150 structure represents DFS_INFO_150 RPC structure.
//
// The DFS_INFO_150 structure relates to the NetrDfsGetInfo, NetrDfsSetInfo, and NetrDfsSetInfo2
// methods when used to retrieve or set security descriptors associated with a DFS link.
// The structure contains the self-relative security descriptor associated with a DFS
// link.
//
// The DFS_INFO_150 structure has the following format.
type Info150 struct {
	// SecurityDescriptorLength:  The length, in bytes, of the buffer that the pSecurityDescriptor
	// field points to.
	SecurityDescriptorLength uint32 `idl:"name:SecurityDescriptorLength" json:"security_descriptor_length"`
	// pSecurityDescriptor:  A self-relative security descriptor associated with DFS. For
	// more information on security descriptors, see [MS-DTYP] section 2.4.6.
	SecurityDescriptor []byte `idl:"name:pSecurityDescriptor;size_is:(SecurityDescriptorLength)" json:"security_descriptor"`
}

func (o *Info150) xxx_PreparePayload(ctx context.Context) error {
	if o.SecurityDescriptor != nil && o.SecurityDescriptorLength == 0 {
		o.SecurityDescriptorLength = uint32(len(o.SecurityDescriptor))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info150) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.SecurityDescriptorLength); err != nil {
		return err
	}
	if o.SecurityDescriptor != nil || o.SecurityDescriptorLength > 0 {
		_ptr_pSecurityDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.SecurityDescriptorLength)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
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
		if err := w.WritePointer(&o.SecurityDescriptor, _ptr_pSecurityDescriptor); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info150) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.SecurityDescriptorLength); err != nil {
		return err
	}
	_ptr_pSecurityDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.SecurityDescriptorLength > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.SecurityDescriptorLength)
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
	_s_pSecurityDescriptor := func(ptr interface{}) { o.SecurityDescriptor = *ptr.(*[]byte) }
	if err := w.ReadPointer(&o.SecurityDescriptor, _s_pSecurityDescriptor, _ptr_pSecurityDescriptor); err != nil {
		return err
	}
	return nil
}

// Info200 structure represents DFS_INFO_200 RPC structure.
//
// The DFS_INFO_200 structure relates to the NetrDfsEnumEx method when used to enumerate
// all of the domain-based DFS namespace in a domain. The structure contains the name
// of a domain-based DFS namespace. The DFS_INFO_200 structure has the following format.
type Info200 struct {
	// FtDfsName:  A pointer to a DFS root path.
	FTDFSName string `idl:"name:FtDfsName;string" json:"ft_dfs_name"`
}

func (o *Info200) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info200) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(6); err != nil {
		return err
	}
	if o.FTDFSName != "" {
		_ptr_FtDfsName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.FTDFSName); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.FTDFSName, _ptr_FtDfsName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info200) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(6); err != nil {
		return err
	}
	_ptr_FtDfsName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.FTDFSName); err != nil {
			return err
		}
		return nil
	})
	_s_FtDfsName := func(ptr interface{}) { o.FTDFSName = *ptr.(*string) }
	if err := w.ReadPointer(&o.FTDFSName, _s_FtDfsName, _ptr_FtDfsName); err != nil {
		return err
	}
	return nil
}

// Info300 structure represents DFS_INFO_300 RPC structure.
//
// The DFS_INFO_300 structure relates to the NetrDfsEnum and NetrDfsEnumEx methods when
// used to enumerate DFS roots hosted on a server. The structure contains the name and
// type (domain-based or stand-alone) of a DFS namespace. The DFS_INFO_300 structure
// has the following format.
type Info300 struct {
	// Flags:  This value specifies the type of the DFS namespace. This MUST have one of
	// the following two permitted values.
	//
	//	+-----------------------------------------+-----------------------------+
	//	|                                         |                             |
	//	|                  VALUE                  |           MEANING           |
	//	|                                         |                             |
	//	+-----------------------------------------+-----------------------------+
	//	+-----------------------------------------+-----------------------------+
	//	| DFS_VOLUME_FLAVOR_STANDALONE 0x00000100 | Stand-alone DFS namespace.  |
	//	+-----------------------------------------+-----------------------------+
	//	| DFS_VOLUME_FLAVOR_AD_BLOB 0x00000200    | Domain-based DFS namespace. |
	//	+-----------------------------------------+-----------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// DfsName:  A pointer to a DFS root path.
	Name string `idl:"name:DfsName;string" json:"name"`
}

func (o *Info300) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info300) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Flags); err != nil {
		return err
	}
	if o.Name != "" {
		_ptr_DfsName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Name, _ptr_DfsName); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info300) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Flags); err != nil {
		return err
	}
	_ptr_DfsName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
			return err
		}
		return nil
	})
	_s_DfsName := func(ptr interface{}) { o.Name = *ptr.(*string) }
	if err := w.ReadPointer(&o.Name, _s_DfsName, _ptr_DfsName); err != nil {
		return err
	}
	return nil
}

// Info structure represents DFS_INFO_STRUCT RPC union.
//
// The DFS_INFO_STRUCT union relates to the NetrDfsGetInfo, NetrDfsSetInfo, and NetrDfsSetInfo2
// methods when used to retrieve or set the configuration of the DFS server. The usage
// model of this union is for the client to specify a Level parameter to determine which
// case of the DFS_INFO_STRUCT to use.
//
// The DFS_INFO_STRUCT union has the following format.
type Info struct {
	// Types that are assignable to Value
	//
	// *Info_1
	// *Info_2
	// *Info_3
	// *Info_4
	// *Info_5
	// *Info_6
	// *Info_7
	// *Info_8
	// *Info_9
	// *Info_50
	// *Info_100
	// *Info_101
	// *Info_102
	// *Info_103
	// *Info_104
	// *Info_105
	// *Info_106
	// *Info_107
	// *Info_150
	Value is_Info `json:"value"`
}

func (o *Info) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *Info_1:
		if value != nil {
			return value.Info1
		}
	case *Info_2:
		if value != nil {
			return value.Info2
		}
	case *Info_3:
		if value != nil {
			return value.Info3
		}
	case *Info_4:
		if value != nil {
			return value.Info4
		}
	case *Info_5:
		if value != nil {
			return value.Info5
		}
	case *Info_6:
		if value != nil {
			return value.Info6
		}
	case *Info_7:
		if value != nil {
			return value.Info7
		}
	case *Info_8:
		if value != nil {
			return value.Info8
		}
	case *Info_9:
		if value != nil {
			return value.Info9
		}
	case *Info_50:
		if value != nil {
			return value.Info50
		}
	case *Info_100:
		if value != nil {
			return value.Info100
		}
	case *Info_101:
		if value != nil {
			return value.Info101
		}
	case *Info_102:
		if value != nil {
			return value.Info102
		}
	case *Info_103:
		if value != nil {
			return value.Info103
		}
	case *Info_104:
		if value != nil {
			return value.Info104
		}
	case *Info_105:
		if value != nil {
			return value.Info105
		}
	case *Info_106:
		if value != nil {
			return value.Info106
		}
	case *Info_107:
		if value != nil {
			return value.Info107
		}
	case *Info_150:
		if value != nil {
			return value.Info150
		}
	}
	return nil
}

type is_Info interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_Info()
}

func (o *Info) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *Info_1:
		return uint32(1)
	case *Info_2:
		return uint32(2)
	case *Info_3:
		return uint32(3)
	case *Info_4:
		return uint32(4)
	case *Info_5:
		return uint32(5)
	case *Info_6:
		return uint32(6)
	case *Info_7:
		return uint32(7)
	case *Info_8:
		return uint32(8)
	case *Info_9:
		return uint32(9)
	case *Info_50:
		return uint32(50)
	case *Info_100:
		return uint32(100)
	case *Info_101:
		return uint32(101)
	case *Info_102:
		return uint32(102)
	case *Info_103:
		return uint32(103)
	case *Info_104:
		return uint32(104)
	case *Info_105:
		return uint32(105)
	case *Info_106:
		return uint32(106)
	case *Info_107:
		return uint32(107)
	case *Info_150:
		return uint32(150)
	}
	return uint32(0)
}

func (o *Info) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
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
		_o, _ := o.Value.(*Info_1)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info_1{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*Info_2)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info_2{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(3):
		_o, _ := o.Value.(*Info_3)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info_3{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(4):
		_o, _ := o.Value.(*Info_4)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info_4{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(5):
		_o, _ := o.Value.(*Info_5)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info_5{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(6):
		_o, _ := o.Value.(*Info_6)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info_6{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(7):
		_o, _ := o.Value.(*Info_7)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info_7{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(8):
		_o, _ := o.Value.(*Info_8)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info_8{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(9):
		_o, _ := o.Value.(*Info_9)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info_9{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(50):
		_o, _ := o.Value.(*Info_50)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info_50{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(100):
		_o, _ := o.Value.(*Info_100)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info_100{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(101):
		_o, _ := o.Value.(*Info_101)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info_101{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(102):
		_o, _ := o.Value.(*Info_102)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info_102{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(103):
		_o, _ := o.Value.(*Info_103)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info_103{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(104):
		_o, _ := o.Value.(*Info_104)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info_104{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(105):
		_o, _ := o.Value.(*Info_105)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info_105{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(106):
		_o, _ := o.Value.(*Info_106)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info_106{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(107):
		_o, _ := o.Value.(*Info_107)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info_107{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(150):
		_o, _ := o.Value.(*Info_150)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&Info_150{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
	}
	return nil
}

func (o *Info) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
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
		o.Value = &Info_1{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &Info_2{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(3):
		o.Value = &Info_3{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(4):
		o.Value = &Info_4{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(5):
		o.Value = &Info_5{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(6):
		o.Value = &Info_6{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(7):
		o.Value = &Info_7{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(8):
		o.Value = &Info_8{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(9):
		o.Value = &Info_9{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(50):
		o.Value = &Info_50{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(100):
		o.Value = &Info_100{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(101):
		o.Value = &Info_101{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(102):
		o.Value = &Info_102{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(103):
		o.Value = &Info_103{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(104):
		o.Value = &Info_104{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(105):
		o.Value = &Info_105{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(106):
		o.Value = &Info_106{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(107):
		o.Value = &Info_107{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(150):
		o.Value = &Info_150{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
	}
	return nil
}

// Info_1 structure represents DFS_INFO_STRUCT RPC union arm.
//
// It has following labels: 1
type Info_1 struct {
	// DfsInfo1:  The DFS_INFO_1 structure contains the name of a DFS root or DFS link.
	// For more information on the specifications, see section 2.2.3.1.
	Info1 *Info1 `idl:"name:DfsInfo1" json:"info1"`
}

func (*Info_1) is_Info() {}

func (o *Info_1) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info1 != nil {
		_ptr_DfsInfo1 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info1 != nil {
				if err := o.Info1.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info1{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info1, _ptr_DfsInfo1); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info_1) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo1 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info1 == nil {
			o.Info1 = &Info1{}
		}
		if err := o.Info1.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo1 := func(ptr interface{}) { o.Info1 = *ptr.(**Info1) }
	if err := w.ReadPointer(&o.Info1, _s_DfsInfo1, _ptr_DfsInfo1); err != nil {
		return err
	}
	return nil
}

// Info_2 structure represents DFS_INFO_STRUCT RPC union arm.
//
// It has following labels: 2
type Info_2 struct {
	// DfsInfo2:  The DFS_INFO_2 structure contains information for a DFS root or DFS link.
	// For more information on specifications, see section 2.2.3.2.
	Info2 *Info2 `idl:"name:DfsInfo2" json:"info2"`
}

func (*Info_2) is_Info() {}

func (o *Info_2) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info2 != nil {
		_ptr_DfsInfo2 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info2 != nil {
				if err := o.Info2.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info2{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info2, _ptr_DfsInfo2); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info_2) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo2 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info2 == nil {
			o.Info2 = &Info2{}
		}
		if err := o.Info2.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo2 := func(ptr interface{}) { o.Info2 = *ptr.(**Info2) }
	if err := w.ReadPointer(&o.Info2, _s_DfsInfo2, _ptr_DfsInfo2); err != nil {
		return err
	}
	return nil
}

// Info_3 structure represents DFS_INFO_STRUCT RPC union arm.
//
// It has following labels: 3
type Info_3 struct {
	// DfsInfo3:  The DFS_INFO_3 structure contains information for a DFS root or DFS link.
	// For more information on specifications, see section 2.2.3.3.
	Info3 *Info3 `idl:"name:DfsInfo3" json:"info3"`
}

func (*Info_3) is_Info() {}

func (o *Info_3) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info3 != nil {
		_ptr_DfsInfo3 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info3 != nil {
				if err := o.Info3.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info3{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info3, _ptr_DfsInfo3); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info_3) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo3 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info3 == nil {
			o.Info3 = &Info3{}
		}
		if err := o.Info3.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo3 := func(ptr interface{}) { o.Info3 = *ptr.(**Info3) }
	if err := w.ReadPointer(&o.Info3, _s_DfsInfo3, _ptr_DfsInfo3); err != nil {
		return err
	}
	return nil
}

// Info_4 structure represents DFS_INFO_STRUCT RPC union arm.
//
// It has following labels: 4
type Info_4 struct {
	// DfsInfo4:  The DFS_INFO_4 structure contains information for a DFS root or DFS link.
	// For more information on specifications, see section 2.2.3.4.
	Info4 *Info4 `idl:"name:DfsInfo4" json:"info4"`
}

func (*Info_4) is_Info() {}

func (o *Info_4) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info4 != nil {
		_ptr_DfsInfo4 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info4 != nil {
				if err := o.Info4.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info4{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info4, _ptr_DfsInfo4); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info_4) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo4 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info4 == nil {
			o.Info4 = &Info4{}
		}
		if err := o.Info4.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo4 := func(ptr interface{}) { o.Info4 = *ptr.(**Info4) }
	if err := w.ReadPointer(&o.Info4, _s_DfsInfo4, _ptr_DfsInfo4); err != nil {
		return err
	}
	return nil
}

// Info_5 structure represents DFS_INFO_STRUCT RPC union arm.
//
// It has following labels: 5
type Info_5 struct {
	// DfsInfo5:  The DFS_INFO_5 structure contains information for a DFS root or DFS link.
	// For more information on specifications, see section 2.2.3.5.
	Info5 *Info5 `idl:"name:DfsInfo5" json:"info5"`
}

func (*Info_5) is_Info() {}

func (o *Info_5) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info5 != nil {
		_ptr_DfsInfo5 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info5 != nil {
				if err := o.Info5.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info5{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info5, _ptr_DfsInfo5); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info_5) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo5 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info5 == nil {
			o.Info5 = &Info5{}
		}
		if err := o.Info5.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo5 := func(ptr interface{}) { o.Info5 = *ptr.(**Info5) }
	if err := w.ReadPointer(&o.Info5, _s_DfsInfo5, _ptr_DfsInfo5); err != nil {
		return err
	}
	return nil
}

// Info_6 structure represents DFS_INFO_STRUCT RPC union arm.
//
// It has following labels: 6
type Info_6 struct {
	// DfsInfo6:  The DFS_INFO_6 structure contains information for a DFS root or DFS link.
	// For more information on specifications, see section 2.2.3.6.
	Info6 *Info6 `idl:"name:DfsInfo6" json:"info6"`
}

func (*Info_6) is_Info() {}

func (o *Info_6) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info6 != nil {
		_ptr_DfsInfo6 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info6 != nil {
				if err := o.Info6.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info6{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info6, _ptr_DfsInfo6); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info_6) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo6 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info6 == nil {
			o.Info6 = &Info6{}
		}
		if err := o.Info6.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo6 := func(ptr interface{}) { o.Info6 = *ptr.(**Info6) }
	if err := w.ReadPointer(&o.Info6, _s_DfsInfo6, _ptr_DfsInfo6); err != nil {
		return err
	}
	return nil
}

// Info_7 structure represents DFS_INFO_STRUCT RPC union arm.
//
// It has following labels: 7
type Info_7 struct {
	// DfsInfo7:  The DFS_INFO_7 structure contains information about a DFS root or DFS
	// link. For more information on specifications, see section 2.2.3.7.
	Info7 *Info7 `idl:"name:DfsInfo7" json:"info7"`
}

func (*Info_7) is_Info() {}

func (o *Info_7) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info7 != nil {
		_ptr_DfsInfo7 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info7 != nil {
				if err := o.Info7.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info7{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info7, _ptr_DfsInfo7); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info_7) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo7 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info7 == nil {
			o.Info7 = &Info7{}
		}
		if err := o.Info7.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo7 := func(ptr interface{}) { o.Info7 = *ptr.(**Info7) }
	if err := w.ReadPointer(&o.Info7, _s_DfsInfo7, _ptr_DfsInfo7); err != nil {
		return err
	}
	return nil
}

// Info_8 structure represents DFS_INFO_STRUCT RPC union arm.
//
// It has following labels: 8
type Info_8 struct {
	// DfsInfo8:  The DFS_INFO_8 structure contains information about a DFS root or DFS
	// link. For more information on specifications, see section 2.2.3.8.
	Info8 *Info8 `idl:"name:DfsInfo8" json:"info8"`
}

func (*Info_8) is_Info() {}

func (o *Info_8) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info8 != nil {
		_ptr_DfsInfo8 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info8 != nil {
				if err := o.Info8.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info8{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info8, _ptr_DfsInfo8); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info_8) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo8 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info8 == nil {
			o.Info8 = &Info8{}
		}
		if err := o.Info8.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo8 := func(ptr interface{}) { o.Info8 = *ptr.(**Info8) }
	if err := w.ReadPointer(&o.Info8, _s_DfsInfo8, _ptr_DfsInfo8); err != nil {
		return err
	}
	return nil
}

// Info_9 structure represents DFS_INFO_STRUCT RPC union arm.
//
// It has following labels: 9
type Info_9 struct {
	// DfsInfo9:  The DFS_INFO_9 structure contains information about a DFS root or DFS
	// link. For more information on specifications, see section 2.2.3.9.
	Info9 *Info9 `idl:"name:DfsInfo9" json:"info9"`
}

func (*Info_9) is_Info() {}

func (o *Info_9) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info9 != nil {
		_ptr_DfsInfo9 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info9 != nil {
				if err := o.Info9.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info9{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info9, _ptr_DfsInfo9); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info_9) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo9 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info9 == nil {
			o.Info9 = &Info9{}
		}
		if err := o.Info9.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo9 := func(ptr interface{}) { o.Info9 = *ptr.(**Info9) }
	if err := w.ReadPointer(&o.Info9, _s_DfsInfo9, _ptr_DfsInfo9); err != nil {
		return err
	}
	return nil
}

// Info_50 structure represents DFS_INFO_STRUCT RPC union arm.
//
// It has following labels: 50
type Info_50 struct {
	// DfsInfo50:  The DFS_INFO_50 structure contains information about a DFS root or DFS
	// link. For more information on specifications, see section 2.2.3.10.
	Info50 *Info50 `idl:"name:DfsInfo50" json:"info50"`
}

func (*Info_50) is_Info() {}

func (o *Info_50) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info50 != nil {
		_ptr_DfsInfo50 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info50 != nil {
				if err := o.Info50.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info50{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info50, _ptr_DfsInfo50); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info_50) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo50 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info50 == nil {
			o.Info50 = &Info50{}
		}
		if err := o.Info50.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo50 := func(ptr interface{}) { o.Info50 = *ptr.(**Info50) }
	if err := w.ReadPointer(&o.Info50, _s_DfsInfo50, _ptr_DfsInfo50); err != nil {
		return err
	}
	return nil
}

// Info_100 structure represents DFS_INFO_STRUCT RPC union arm.
//
// It has following labels: 100
type Info_100 struct {
	// DfsInfo100:  The DFS_INFO_100 structure contains a comment associated with a DFS
	// root or DFS link. For more information on specifications, see section 2.2.5.1.
	Info100 *Info100 `idl:"name:DfsInfo100" json:"info100"`
}

func (*Info_100) is_Info() {}

func (o *Info_100) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info100 != nil {
		_ptr_DfsInfo100 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info100 != nil {
				if err := o.Info100.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info100{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info100, _ptr_DfsInfo100); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info_100) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo100 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info100 == nil {
			o.Info100 = &Info100{}
		}
		if err := o.Info100.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo100 := func(ptr interface{}) { o.Info100 = *ptr.(**Info100) }
	if err := w.ReadPointer(&o.Info100, _s_DfsInfo100, _ptr_DfsInfo100); err != nil {
		return err
	}
	return nil
}

// Info_101 structure represents DFS_INFO_STRUCT RPC union arm.
//
// It has following labels: 101
type Info_101 struct {
	// DfsInfo101:  The DFS_INFO_101 structure describes the storage state on a DFS root,
	// DFS link, DFS root target, or DFS link target. For more information on specifications,
	// see section 2.2.4.1.
	Info101 *Info101 `idl:"name:DfsInfo101" json:"info101"`
}

func (*Info_101) is_Info() {}

func (o *Info_101) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info101 != nil {
		_ptr_DfsInfo101 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info101 != nil {
				if err := o.Info101.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info101{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info101, _ptr_DfsInfo101); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info_101) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo101 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info101 == nil {
			o.Info101 = &Info101{}
		}
		if err := o.Info101.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo101 := func(ptr interface{}) { o.Info101 = *ptr.(**Info101) }
	if err := w.ReadPointer(&o.Info101, _s_DfsInfo101, _ptr_DfsInfo101); err != nil {
		return err
	}
	return nil
}

// Info_102 structure represents DFS_INFO_STRUCT RPC union arm.
//
// It has following labels: 102
type Info_102 struct {
	// DfsInfo102:  The DFS_INFO_102 structure contains a time-out value for a DFS root
	// or DFS link. For more information on specifications, see section 2.2.4.2.
	Info102 *Info102 `idl:"name:DfsInfo102" json:"info102"`
}

func (*Info_102) is_Info() {}

func (o *Info_102) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info102 != nil {
		_ptr_DfsInfo102 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info102 != nil {
				if err := o.Info102.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info102{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info102, _ptr_DfsInfo102); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info_102) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo102 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info102 == nil {
			o.Info102 = &Info102{}
		}
		if err := o.Info102.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo102 := func(ptr interface{}) { o.Info102 = *ptr.(**Info102) }
	if err := w.ReadPointer(&o.Info102, _s_DfsInfo102, _ptr_DfsInfo102); err != nil {
		return err
	}
	return nil
}

// Info_103 structure represents DFS_INFO_STRUCT RPC union arm.
//
// It has following labels: 103
type Info_103 struct {
	// DfsInfo103:  The DFS_INFO_103 structure contains properties that set specific behaviors
	// for a DFS root or DFS link. For more information on specifications, see section 2.2.4.3.
	Info103 *Info103 `idl:"name:DfsInfo103" json:"info103"`
}

func (*Info_103) is_Info() {}

func (o *Info_103) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info103 != nil {
		_ptr_DfsInfo103 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info103 != nil {
				if err := o.Info103.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info103{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info103, _ptr_DfsInfo103); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info_103) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo103 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info103 == nil {
			o.Info103 = &Info103{}
		}
		if err := o.Info103.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo103 := func(ptr interface{}) { o.Info103 = *ptr.(**Info103) }
	if err := w.ReadPointer(&o.Info103, _s_DfsInfo103, _ptr_DfsInfo103); err != nil {
		return err
	}
	return nil
}

// Info_104 structure represents DFS_INFO_STRUCT RPC union arm.
//
// It has following labels: 104
type Info_104 struct {
	// DfsInfo104:  The DFS_INFO_104 structure contains the priority of a DFS root target
	// or DFS link target. For more information on specifications, see section 2.2.4.4.
	Info104 *Info104 `idl:"name:DfsInfo104" json:"info104"`
}

func (*Info_104) is_Info() {}

func (o *Info_104) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info104 != nil {
		_ptr_DfsInfo104 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info104 != nil {
				if err := o.Info104.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info104{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info104, _ptr_DfsInfo104); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info_104) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo104 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info104 == nil {
			o.Info104 = &Info104{}
		}
		if err := o.Info104.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo104 := func(ptr interface{}) { o.Info104 = *ptr.(**Info104) }
	if err := w.ReadPointer(&o.Info104, _s_DfsInfo104, _ptr_DfsInfo104); err != nil {
		return err
	}
	return nil
}

// Info_105 structure represents DFS_INFO_STRUCT RPC union arm.
//
// It has following labels: 105
type Info_105 struct {
	// DfsInfo105:  The DFS_INFO_105 structure contains information about a DFS root or
	// DFS link, including comment, state, time-out, and DFS behaviors that property flags
	// specify. For more information on specifications, see section 2.2.4.5.
	Info105 *Info105 `idl:"name:DfsInfo105" json:"info105"`
}

func (*Info_105) is_Info() {}

func (o *Info_105) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info105 != nil {
		_ptr_DfsInfo105 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info105 != nil {
				if err := o.Info105.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info105{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info105, _ptr_DfsInfo105); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info_105) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo105 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info105 == nil {
			o.Info105 = &Info105{}
		}
		if err := o.Info105.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo105 := func(ptr interface{}) { o.Info105 = *ptr.(**Info105) }
	if err := w.ReadPointer(&o.Info105, _s_DfsInfo105, _ptr_DfsInfo105); err != nil {
		return err
	}
	return nil
}

// Info_106 structure represents DFS_INFO_STRUCT RPC union arm.
//
// It has following labels: 106
type Info_106 struct {
	// DfsInfo106:  The DFS_INFO_106 structure contains the storage state and priority for
	// a DFS root target or DFS link target. For more information on specifications, see
	// section 2.2.4.6.
	Info106 *Info106 `idl:"name:DfsInfo106" json:"info106"`
}

func (*Info_106) is_Info() {}

func (o *Info_106) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info106 != nil {
		_ptr_DfsInfo106 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info106 != nil {
				if err := o.Info106.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info106{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info106, _ptr_DfsInfo106); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info_106) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo106 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info106 == nil {
			o.Info106 = &Info106{}
		}
		if err := o.Info106.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo106 := func(ptr interface{}) { o.Info106 = *ptr.(**Info106) }
	if err := w.ReadPointer(&o.Info106, _s_DfsInfo106, _ptr_DfsInfo106); err != nil {
		return err
	}
	return nil
}

// Info_107 structure represents DFS_INFO_STRUCT RPC union arm.
//
// It has following labels: 107
type Info_107 struct {
	// DfsInfo107:  The DFS_INFO_107 structure contains the storage state and priority for
	// a DFS root target or DFS link target. For more information on specifications, see
	// section 2.2.4.7.
	Info107 *Info107 `idl:"name:DfsInfo107" json:"info107"`
}

func (*Info_107) is_Info() {}

func (o *Info_107) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info107 != nil {
		_ptr_DfsInfo107 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info107 != nil {
				if err := o.Info107.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info107{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info107, _ptr_DfsInfo107); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info_107) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo107 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info107 == nil {
			o.Info107 = &Info107{}
		}
		if err := o.Info107.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo107 := func(ptr interface{}) { o.Info107 = *ptr.(**Info107) }
	if err := w.ReadPointer(&o.Info107, _s_DfsInfo107, _ptr_DfsInfo107); err != nil {
		return err
	}
	return nil
}

// Info_150 structure represents DFS_INFO_STRUCT RPC union arm.
//
// It has following labels: 150
type Info_150 struct {
	// DfsInfo150:  The DFS_INFO_150 structure contains the self-relative security descriptor
	// associated with the DFS link. For more information on specifications, see section
	// 2.2.5.2.
	Info150 *Info150 `idl:"name:DfsInfo150" json:"info150"`
}

func (*Info_150) is_Info() {}

func (o *Info_150) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info150 != nil {
		_ptr_DfsInfo150 := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info150 != nil {
				if err := o.Info150.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info150{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info150, _ptr_DfsInfo150); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *Info_150) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo150 := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info150 == nil {
			o.Info150 = &Info150{}
		}
		if err := o.Info150.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo150 := func(ptr interface{}) { o.Info150 = *ptr.(**Info150) }
	if err := w.ReadPointer(&o.Info150, _s_DfsInfo150, _ptr_DfsInfo150); err != nil {
		return err
	}
	return nil
}

// Info1Container structure represents DFS_INFO_1_CONTAINER RPC structure.
//
// The DFS_INFO_1_CONTAINER structure contains an array of DFS_INFO_1 structures. The
// DFS_INFO_1_CONTAINER structure has the following format.
type Info1Container struct {
	// EntriesRead:  The number of elements in the array.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  The array of DFS_INFO_1 structures.
	Buffer []*Info1 `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *Info1Container) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Info1Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
					if err := (&Info1{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Info1{}).MarshalNDR(ctx, w); err != nil {
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
func (o *Info1Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Buffer = make([]*Info1, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &Info1{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*Info1) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// Info2Container structure represents DFS_INFO_2_CONTAINER RPC structure.
//
// The DFS_INFO_2_CONTAINER structure contains an array of DFS_INFO_2 structures. The
// DFS_INFO_2_CONTAINER structure has the following format.
type Info2Container struct {
	// EntriesRead:  The number of elements in the array.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  The array of DFS_INFO_2 structures.
	Buffer []*Info2 `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *Info2Container) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Info2Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
					if err := (&Info2{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Info2{}).MarshalNDR(ctx, w); err != nil {
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
func (o *Info2Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Buffer = make([]*Info2, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &Info2{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*Info2) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// Info3Container structure represents DFS_INFO_3_CONTAINER RPC structure.
//
// The DFS_INFO_3_CONTAINER structure contains an array of DFS_INFO_3 structures. The
// DFS_INFO_3_CONTAINER structure has the following format.
type Info3Container struct {
	// EntriesRead:  The number of elements in the array.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  The array of DFS_INFO_3 structures.
	Buffer []*Info3 `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *Info3Container) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Info3Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
					if err := (&Info3{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Info3{}).MarshalNDR(ctx, w); err != nil {
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
func (o *Info3Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Buffer = make([]*Info3, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &Info3{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*Info3) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// Info4Container structure represents DFS_INFO_4_CONTAINER RPC structure.
//
// The DFS_INFO_4_CONTAINER structure contains an array of DFS_INFO_4 structures. The
// DFS_INFO_4_CONTAINER structure has the following format.
type Info4Container struct {
	// EntriesRead:  The number of elements in the array.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  The array of DFS_INFO_4 structures.
	Buffer []*Info4 `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *Info4Container) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Info4Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
					if err := (&Info4{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Info4{}).MarshalNDR(ctx, w); err != nil {
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
func (o *Info4Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Buffer = make([]*Info4, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &Info4{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*Info4) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// Info5Container structure represents DFS_INFO_5_CONTAINER RPC structure.
//
// The DFS_INFO_5_CONTAINER structure contains an array of DFS_INFO_5 structures. The
// DFS_INFO_5_CONTAINER structure has the following format.
type Info5Container struct {
	// EntriesRead:  The number of elements in the array.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  The array of DFS_INFO_5 structures.
	Buffer []*Info5 `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *Info5Container) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Info5Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
					if err := (&Info5{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Info5{}).MarshalNDR(ctx, w); err != nil {
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
func (o *Info5Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Buffer = make([]*Info5, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &Info5{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*Info5) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// Info6Container structure represents DFS_INFO_6_CONTAINER RPC structure.
//
// The DFS_INFO_6_CONTAINER structure contains an array of DFS_INFO_6 structures. The
// DFS_INFO_6_CONTAINER structure has the following format.
type Info6Container struct {
	// EntriesRead:  The number of elements in the array.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  The array of DFS_INFO_6 structures.
	Buffer []*Info6 `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *Info6Container) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Info6Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
					if err := (&Info6{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Info6{}).MarshalNDR(ctx, w); err != nil {
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
func (o *Info6Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Buffer = make([]*Info6, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &Info6{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*Info6) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// Info8Container structure represents DFS_INFO_8_CONTAINER RPC structure.
//
// The DFS_INFO_8_CONTAINER structure contains an array of DFS_INFO_8 structures. The
// DFS_INFO_8_CONTAINER structure has the following format.
type Info8Container struct {
	// EntriesRead:  The number of DFS_INFO_8 elements in the array.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  The array of DFS_INFO_8 structures.
	Buffer []*Info8 `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *Info8Container) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Info8Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
					if err := (&Info8{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Info8{}).MarshalNDR(ctx, w); err != nil {
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
func (o *Info8Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Buffer = make([]*Info8, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &Info8{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*Info8) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// Info9Container structure represents DFS_INFO_9_CONTAINER RPC structure.
//
// The DFS_INFO_9_CONTAINER structure contains an array of DFS_INFO_9 structures. The
// DFS_INFO_9_CONTAINER structure has the following format.
type Info9Container struct {
	// EntriesRead:  The number of DFS_INFO_9 elements in the array.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  The array of DFS_INFO_9 structures.
	Buffer []*Info9 `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *Info9Container) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Info9Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
					if err := (&Info9{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Info9{}).MarshalNDR(ctx, w); err != nil {
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
func (o *Info9Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Buffer = make([]*Info9, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &Info9{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*Info9) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// Info200Container structure represents DFS_INFO_200_CONTAINER RPC structure.
//
// The DFS_INFO_200_CONTAINER structure contains an array of DFS_INFO_200 structures.
// The DFS_INFO_200_CONTAINER structure has the following format.
type Info200Container struct {
	// EntriesRead:  The number of elements in the array.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  The array of DFS_INFO_200 structures.
	Buffer []*Info200 `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *Info200Container) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Info200Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
					if err := (&Info200{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Info200{}).MarshalNDR(ctx, w); err != nil {
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
func (o *Info200Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Buffer = make([]*Info200, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &Info200{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*Info200) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// Info300Container structure represents DFS_INFO_300_CONTAINER RPC structure.
//
// The DFS_INFO_300_CONTAINER structure contains an array of DFS_INFO_300 structures.
// The DFS_INFO_300_CONTAINER structure has the following format.
type Info300Container struct {
	// EntriesRead:  The number of elements in the array.
	EntriesRead uint32 `idl:"name:EntriesRead" json:"entries_read"`
	// Buffer:  The array of DFS_INFO_300 structures.
	Buffer []*Info300 `idl:"name:Buffer;size_is:(EntriesRead)" json:"buffer"`
}

func (o *Info300Container) xxx_PreparePayload(ctx context.Context) error {
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
func (o *Info300Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
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
					if err := (&Info300{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&Info300{}).MarshalNDR(ctx, w); err != nil {
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
func (o *Info300Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
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
		o.Buffer = make([]*Info300, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if o.Buffer[i1] == nil {
				o.Buffer[i1] = &Info300{}
			}
			if err := o.Buffer[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]*Info300) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}

// InfoEnum structure represents DFS_INFO_ENUM_STRUCT RPC structure.
//
// The DFS_INFO_ENUM_STRUCT union relates to the NetrDfsEnum and NetrDfsEnumEx methods
// when used to enumerate the configuration of the DFS server.
//
// The DFS_INFO_ENUM_STRUCT union structure has the following format.
type InfoEnum struct {
	// Level:  Specifies the case of the DfsInfoContainer union.
	Level uint32 `idl:"name:Level" json:"level"`
	// DfsInfoContainer:  Union of the possible enumeration containers.
	InfoContainer *InfoEnum_InfoContainer `idl:"name:DfsInfoContainer;switch_is:Level" json:"info_container"`
}

func (o *InfoEnum) xxx_PreparePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *InfoEnum) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Level); err != nil {
		return err
	}
	_swInfoContainer := uint32(o.Level)
	if o.InfoContainer != nil {
		if err := o.InfoContainer.MarshalUnionNDR(ctx, w, _swInfoContainer); err != nil {
			return err
		}
	} else {
		if err := (&InfoEnum_InfoContainer{}).MarshalUnionNDR(ctx, w, _swInfoContainer); err != nil {
			return err
		}
	}
	return nil
}
func (o *InfoEnum) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Level); err != nil {
		return err
	}
	if o.InfoContainer == nil {
		o.InfoContainer = &InfoEnum_InfoContainer{}
	}
	_swInfoContainer := uint32(o.Level)
	if err := o.InfoContainer.UnmarshalUnionNDR(ctx, w, _swInfoContainer); err != nil {
		return err
	}
	return nil
}

// InfoEnum_InfoContainer structure represents DFS_INFO_ENUM_STRUCT union anonymous member.
//
// The DFS_INFO_ENUM_STRUCT union relates to the NetrDfsEnum and NetrDfsEnumEx methods
// when used to enumerate the configuration of the DFS server.
//
// The DFS_INFO_ENUM_STRUCT union structure has the following format.
type InfoEnum_InfoContainer struct {
	// Types that are assignable to Value
	//
	// *InfoContainer_Info1Container
	// *InfoContainer_Info2Container
	// *InfoContainer_Info3Container
	// *InfoContainer_Info4Container
	// *InfoContainer_Info5Container
	// *InfoContainer_Info6Container
	// *InfoContainer_Info8Container
	// *InfoContainer_Info9Container
	// *InfoContainer_Info200Container
	// *InfoContainer_Info300Container
	Value is_InfoEnum_InfoContainer `json:"value"`
}

func (o *InfoEnum_InfoContainer) GetValue() any {
	if o == nil {
		return nil
	}
	switch value := (interface{})(o.Value).(type) {
	case *InfoContainer_Info1Container:
		if value != nil {
			return value.Info1Container
		}
	case *InfoContainer_Info2Container:
		if value != nil {
			return value.Info2Container
		}
	case *InfoContainer_Info3Container:
		if value != nil {
			return value.Info3Container
		}
	case *InfoContainer_Info4Container:
		if value != nil {
			return value.Info4Container
		}
	case *InfoContainer_Info5Container:
		if value != nil {
			return value.Info5Container
		}
	case *InfoContainer_Info6Container:
		if value != nil {
			return value.Info6Container
		}
	case *InfoContainer_Info8Container:
		if value != nil {
			return value.Info8Container
		}
	case *InfoContainer_Info9Container:
		if value != nil {
			return value.Info9Container
		}
	case *InfoContainer_Info200Container:
		if value != nil {
			return value.Info200Container
		}
	case *InfoContainer_Info300Container:
		if value != nil {
			return value.Info300Container
		}
	}
	return nil
}

type is_InfoEnum_InfoContainer interface {
	ndr.Marshaler
	ndr.Unmarshaler
	is_InfoEnum_InfoContainer()
}

func (o *InfoEnum_InfoContainer) NDRSwitchValue(sw uint32) uint32 {
	if o == nil {
		return uint32(0)
	}
	switch (interface{})(o.Value).(type) {
	case *InfoContainer_Info1Container:
		return uint32(1)
	case *InfoContainer_Info2Container:
		return uint32(2)
	case *InfoContainer_Info3Container:
		return uint32(3)
	case *InfoContainer_Info4Container:
		return uint32(4)
	case *InfoContainer_Info5Container:
		return uint32(5)
	case *InfoContainer_Info6Container:
		return uint32(6)
	case *InfoContainer_Info8Container:
		return uint32(8)
	case *InfoContainer_Info9Container:
		return uint32(9)
	case *InfoContainer_Info200Container:
		return uint32(200)
	case *InfoContainer_Info300Container:
		return uint32(300)
	}
	return uint32(0)
}

func (o *InfoEnum_InfoContainer) MarshalUnionNDR(ctx context.Context, w ndr.Writer, sw uint32) error {
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
		_o, _ := o.Value.(*InfoContainer_Info1Container)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&InfoContainer_Info1Container{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(2):
		_o, _ := o.Value.(*InfoContainer_Info2Container)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&InfoContainer_Info2Container{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(3):
		_o, _ := o.Value.(*InfoContainer_Info3Container)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&InfoContainer_Info3Container{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(4):
		_o, _ := o.Value.(*InfoContainer_Info4Container)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&InfoContainer_Info4Container{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(5):
		_o, _ := o.Value.(*InfoContainer_Info5Container)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&InfoContainer_Info5Container{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(6):
		_o, _ := o.Value.(*InfoContainer_Info6Container)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&InfoContainer_Info6Container{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(8):
		_o, _ := o.Value.(*InfoContainer_Info8Container)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&InfoContainer_Info8Container{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(9):
		_o, _ := o.Value.(*InfoContainer_Info9Container)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&InfoContainer_Info9Container{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(200):
		_o, _ := o.Value.(*InfoContainer_Info200Container)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&InfoContainer_Info200Container{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	case uint32(300):
		_o, _ := o.Value.(*InfoContainer_Info300Container)
		if _o != nil {
			if err := _o.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&InfoContainer_Info300Container{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

func (o *InfoEnum_InfoContainer) UnmarshalUnionNDR(ctx context.Context, w ndr.Reader, sw uint32) error {
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
		o.Value = &InfoContainer_Info1Container{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(2):
		o.Value = &InfoContainer_Info2Container{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(3):
		o.Value = &InfoContainer_Info3Container{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(4):
		o.Value = &InfoContainer_Info4Container{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(5):
		o.Value = &InfoContainer_Info5Container{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(6):
		o.Value = &InfoContainer_Info6Container{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(8):
		o.Value = &InfoContainer_Info8Container{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(9):
		o.Value = &InfoContainer_Info9Container{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(200):
		o.Value = &InfoContainer_Info200Container{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	case uint32(300):
		o.Value = &InfoContainer_Info300Container{}
		if err := o.Value.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported switch case value %v", sw)
	}
	return nil
}

// InfoContainer_Info1Container structure represents InfoEnum_InfoContainer RPC union arm.
//
// It has following labels: 1
type InfoContainer_Info1Container struct {
	// DfsInfo1Container:  The DFS_INFO_1_CONTAINER structure contains an array of the names
	// of DFS roots or DFS links. For more information, see section 2.2.6.1.
	Info1Container *Info1Container `idl:"name:DfsInfo1Container" json:"info1_container"`
}

func (*InfoContainer_Info1Container) is_InfoEnum_InfoContainer() {}

func (o *InfoContainer_Info1Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info1Container != nil {
		_ptr_DfsInfo1Container := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info1Container != nil {
				if err := o.Info1Container.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info1Container{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info1Container, _ptr_DfsInfo1Container); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *InfoContainer_Info1Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo1Container := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info1Container == nil {
			o.Info1Container = &Info1Container{}
		}
		if err := o.Info1Container.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo1Container := func(ptr interface{}) { o.Info1Container = *ptr.(**Info1Container) }
	if err := w.ReadPointer(&o.Info1Container, _s_DfsInfo1Container, _ptr_DfsInfo1Container); err != nil {
		return err
	}
	return nil
}

// InfoContainer_Info2Container structure represents InfoEnum_InfoContainer RPC union arm.
//
// It has following labels: 2
type InfoContainer_Info2Container struct {
	// DfsInfo2Container:  The DFS_INFO_2_CONTAINER structure contains an array of information
	// for DFS roots or DFS links. For more information, see section 2.2.6.2.
	Info2Container *Info2Container `idl:"name:DfsInfo2Container" json:"info2_container"`
}

func (*InfoContainer_Info2Container) is_InfoEnum_InfoContainer() {}

func (o *InfoContainer_Info2Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info2Container != nil {
		_ptr_DfsInfo2Container := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info2Container != nil {
				if err := o.Info2Container.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info2Container{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info2Container, _ptr_DfsInfo2Container); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *InfoContainer_Info2Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo2Container := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info2Container == nil {
			o.Info2Container = &Info2Container{}
		}
		if err := o.Info2Container.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo2Container := func(ptr interface{}) { o.Info2Container = *ptr.(**Info2Container) }
	if err := w.ReadPointer(&o.Info2Container, _s_DfsInfo2Container, _ptr_DfsInfo2Container); err != nil {
		return err
	}
	return nil
}

// InfoContainer_Info3Container structure represents InfoEnum_InfoContainer RPC union arm.
//
// It has following labels: 3
type InfoContainer_Info3Container struct {
	// DfsInfo3Container:  The DFS_INFO_3_CONTAINER structure contains an array of information
	// for DFS roots or DFS links. For more information, see section 2.2.6.3.
	Info3Container *Info3Container `idl:"name:DfsInfo3Container" json:"info3_container"`
}

func (*InfoContainer_Info3Container) is_InfoEnum_InfoContainer() {}

func (o *InfoContainer_Info3Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info3Container != nil {
		_ptr_DfsInfo3Container := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info3Container != nil {
				if err := o.Info3Container.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info3Container{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info3Container, _ptr_DfsInfo3Container); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *InfoContainer_Info3Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo3Container := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info3Container == nil {
			o.Info3Container = &Info3Container{}
		}
		if err := o.Info3Container.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo3Container := func(ptr interface{}) { o.Info3Container = *ptr.(**Info3Container) }
	if err := w.ReadPointer(&o.Info3Container, _s_DfsInfo3Container, _ptr_DfsInfo3Container); err != nil {
		return err
	}
	return nil
}

// InfoContainer_Info4Container structure represents InfoEnum_InfoContainer RPC union arm.
//
// It has following labels: 4
type InfoContainer_Info4Container struct {
	// DfsInfo4Container:  The DFS_INFO_4_CONTAINER structure contains an array of information
	// for DFS roots or DFS links. For more information, see section 2.2.6.4.
	Info4Container *Info4Container `idl:"name:DfsInfo4Container" json:"info4_container"`
}

func (*InfoContainer_Info4Container) is_InfoEnum_InfoContainer() {}

func (o *InfoContainer_Info4Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info4Container != nil {
		_ptr_DfsInfo4Container := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info4Container != nil {
				if err := o.Info4Container.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info4Container{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info4Container, _ptr_DfsInfo4Container); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *InfoContainer_Info4Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo4Container := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info4Container == nil {
			o.Info4Container = &Info4Container{}
		}
		if err := o.Info4Container.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo4Container := func(ptr interface{}) { o.Info4Container = *ptr.(**Info4Container) }
	if err := w.ReadPointer(&o.Info4Container, _s_DfsInfo4Container, _ptr_DfsInfo4Container); err != nil {
		return err
	}
	return nil
}

// InfoContainer_Info5Container structure represents InfoEnum_InfoContainer RPC union arm.
//
// It has following labels: 5
type InfoContainer_Info5Container struct {
	// DfsInfo5Container:  The DFS_INFO_5_CONTAINER structure contains an array of information
	// for DFS roots or DFS links. For more information, see section 2.2.6.5.
	Info5Container *Info5Container `idl:"name:DfsInfo5Container" json:"info5_container"`
}

func (*InfoContainer_Info5Container) is_InfoEnum_InfoContainer() {}

func (o *InfoContainer_Info5Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info5Container != nil {
		_ptr_DfsInfo5Container := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info5Container != nil {
				if err := o.Info5Container.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info5Container{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info5Container, _ptr_DfsInfo5Container); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *InfoContainer_Info5Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo5Container := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info5Container == nil {
			o.Info5Container = &Info5Container{}
		}
		if err := o.Info5Container.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo5Container := func(ptr interface{}) { o.Info5Container = *ptr.(**Info5Container) }
	if err := w.ReadPointer(&o.Info5Container, _s_DfsInfo5Container, _ptr_DfsInfo5Container); err != nil {
		return err
	}
	return nil
}

// InfoContainer_Info6Container structure represents InfoEnum_InfoContainer RPC union arm.
//
// It has following labels: 6
type InfoContainer_Info6Container struct {
	// DfsInfo6Container:  The DFS_INFO_6_CONTAINER structure contains an array of information
	// for DFS roots or DFS links. For more information, see section 2.2.6.6.
	Info6Container *Info6Container `idl:"name:DfsInfo6Container" json:"info6_container"`
}

func (*InfoContainer_Info6Container) is_InfoEnum_InfoContainer() {}

func (o *InfoContainer_Info6Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info6Container != nil {
		_ptr_DfsInfo6Container := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info6Container != nil {
				if err := o.Info6Container.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info6Container{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info6Container, _ptr_DfsInfo6Container); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *InfoContainer_Info6Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo6Container := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info6Container == nil {
			o.Info6Container = &Info6Container{}
		}
		if err := o.Info6Container.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo6Container := func(ptr interface{}) { o.Info6Container = *ptr.(**Info6Container) }
	if err := w.ReadPointer(&o.Info6Container, _s_DfsInfo6Container, _ptr_DfsInfo6Container); err != nil {
		return err
	}
	return nil
}

// InfoContainer_Info8Container structure represents InfoEnum_InfoContainer RPC union arm.
//
// It has following labels: 8
type InfoContainer_Info8Container struct {
	// DfsInfo8Container:  The DFS_INFO_8_CONTAINER structure contains an array of information
	// for DFS roots or DFS links. For more information, see section 2.2.6.7.
	Info8Container *Info8Container `idl:"name:DfsInfo8Container" json:"info8_container"`
}

func (*InfoContainer_Info8Container) is_InfoEnum_InfoContainer() {}

func (o *InfoContainer_Info8Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info8Container != nil {
		_ptr_DfsInfo8Container := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info8Container != nil {
				if err := o.Info8Container.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info8Container{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info8Container, _ptr_DfsInfo8Container); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *InfoContainer_Info8Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo8Container := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info8Container == nil {
			o.Info8Container = &Info8Container{}
		}
		if err := o.Info8Container.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo8Container := func(ptr interface{}) { o.Info8Container = *ptr.(**Info8Container) }
	if err := w.ReadPointer(&o.Info8Container, _s_DfsInfo8Container, _ptr_DfsInfo8Container); err != nil {
		return err
	}
	return nil
}

// InfoContainer_Info9Container structure represents InfoEnum_InfoContainer RPC union arm.
//
// It has following labels: 9
type InfoContainer_Info9Container struct {
	// DfsInfo9Container:  The DFS_INFO_9_CONTAINER structure contains an array of information
	// for DFS roots or DFS links. For more information, see section 2.2.6.8.
	Info9Container *Info9Container `idl:"name:DfsInfo9Container" json:"info9_container"`
}

func (*InfoContainer_Info9Container) is_InfoEnum_InfoContainer() {}

func (o *InfoContainer_Info9Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info9Container != nil {
		_ptr_DfsInfo9Container := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info9Container != nil {
				if err := o.Info9Container.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info9Container{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info9Container, _ptr_DfsInfo9Container); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *InfoContainer_Info9Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo9Container := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info9Container == nil {
			o.Info9Container = &Info9Container{}
		}
		if err := o.Info9Container.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo9Container := func(ptr interface{}) { o.Info9Container = *ptr.(**Info9Container) }
	if err := w.ReadPointer(&o.Info9Container, _s_DfsInfo9Container, _ptr_DfsInfo9Container); err != nil {
		return err
	}
	return nil
}

// InfoContainer_Info200Container structure represents InfoEnum_InfoContainer RPC union arm.
//
// It has following labels: 200
type InfoContainer_Info200Container struct {
	// DfsInfo200Container:  The DFS_INFO_200_CONTAINER structure contains an array of the
	// names of domain-based DFS namespaces in a domain-based DFS. For more information,
	// see section 2.2.6.9.
	Info200Container *Info200Container `idl:"name:DfsInfo200Container" json:"info200_container"`
}

func (*InfoContainer_Info200Container) is_InfoEnum_InfoContainer() {}

func (o *InfoContainer_Info200Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info200Container != nil {
		_ptr_DfsInfo200Container := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info200Container != nil {
				if err := o.Info200Container.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info200Container{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info200Container, _ptr_DfsInfo200Container); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *InfoContainer_Info200Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo200Container := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info200Container == nil {
			o.Info200Container = &Info200Container{}
		}
		if err := o.Info200Container.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo200Container := func(ptr interface{}) { o.Info200Container = *ptr.(**Info200Container) }
	if err := w.ReadPointer(&o.Info200Container, _s_DfsInfo200Container, _ptr_DfsInfo200Container); err != nil {
		return err
	}
	return nil
}

// InfoContainer_Info300Container structure represents InfoEnum_InfoContainer RPC union arm.
//
// It has following labels: 300
type InfoContainer_Info300Container struct {
	// DfsInfo300Container:  The DFS_INFO_300_CONTAINER structure contains an array of the
	// DFS roots hosted on a server. For more information, see section 2.2.6.10.
	Info300Container *Info300Container `idl:"name:DfsInfo300Container" json:"info300_container"`
}

func (*InfoContainer_Info300Container) is_InfoEnum_InfoContainer() {}

func (o *InfoContainer_Info300Container) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if o.Info300Container != nil {
		_ptr_DfsInfo300Container := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if o.Info300Container != nil {
				if err := o.Info300Container.MarshalNDR(ctx, w); err != nil {
					return err
				}
			} else {
				if err := (&Info300Container{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Info300Container, _ptr_DfsInfo300Container); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *InfoContainer_Info300Container) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	_ptr_DfsInfo300Container := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		if o.Info300Container == nil {
			o.Info300Container = &Info300Container{}
		}
		if err := o.Info300Container.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		return nil
	})
	_s_DfsInfo300Container := func(ptr interface{}) { o.Info300Container = *ptr.(**Info300Container) }
	if err := w.ReadPointer(&o.Info300Container, _s_DfsInfo300Container, _ptr_DfsInfo300Container); err != nil {
		return err
	}
	return nil
}

type xxx_DefaultNetdfsClient struct {
	cc dcerpc.Conn
}

func (o *xxx_DefaultNetdfsClient) ManagerGetVersion(ctx context.Context, in *ManagerGetVersionRequest, opts ...dcerpc.CallOption) (*ManagerGetVersionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ManagerGetVersionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) Add(ctx context.Context, in *AddRequest, opts ...dcerpc.CallOption) (*AddResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) Remove(ctx context.Context, in *RemoveRequest, opts ...dcerpc.CallOption) (*RemoveResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) SetInfo(ctx context.Context, in *SetInfoRequest, opts ...dcerpc.CallOption) (*SetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...dcerpc.CallOption) (*GetInfoResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetInfoResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) Enum(ctx context.Context, in *EnumRequest, opts ...dcerpc.CallOption) (*EnumResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) Move(ctx context.Context, in *MoveRequest, opts ...dcerpc.CallOption) (*MoveResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &MoveResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) AddFTRoot(ctx context.Context, in *AddFTRootRequest, opts ...dcerpc.CallOption) (*AddFTRootResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddFTRootResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) RemoveFTRoot(ctx context.Context, in *RemoveFTRootRequest, opts ...dcerpc.CallOption) (*RemoveFTRootResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveFTRootResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) AddStdRoot(ctx context.Context, in *AddStdRootRequest, opts ...dcerpc.CallOption) (*AddStdRootResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddStdRootResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) RemoveStdRoot(ctx context.Context, in *RemoveStdRootRequest, opts ...dcerpc.CallOption) (*RemoveStdRootResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveStdRootResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) ManagerInitialize(ctx context.Context, in *ManagerInitializeRequest, opts ...dcerpc.CallOption) (*ManagerInitializeResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ManagerInitializeResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) AddStdRootForced(ctx context.Context, in *AddStdRootForcedRequest, opts ...dcerpc.CallOption) (*AddStdRootForcedResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddStdRootForcedResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) GetDCAddress(ctx context.Context, in *GetDCAddressRequest, opts ...dcerpc.CallOption) (*GetDCAddressResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetDCAddressResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) SetDCAddress(ctx context.Context, in *SetDCAddressRequest, opts ...dcerpc.CallOption) (*SetDCAddressResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetDCAddressResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) FlushFTTable(ctx context.Context, in *FlushFTTableRequest, opts ...dcerpc.CallOption) (*FlushFTTableResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FlushFTTableResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) Add2(ctx context.Context, in *Add2Request, opts ...dcerpc.CallOption) (*Add2Response, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &Add2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) Remove2(ctx context.Context, in *Remove2Request, opts ...dcerpc.CallOption) (*Remove2Response, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &Remove2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) EnumEx(ctx context.Context, in *EnumExRequest, opts ...dcerpc.CallOption) (*EnumExResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &EnumExResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) SetInfo2(ctx context.Context, in *SetInfo2Request, opts ...dcerpc.CallOption) (*SetInfo2Response, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &SetInfo2Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) AddRootTarget(ctx context.Context, in *AddRootTargetRequest, opts ...dcerpc.CallOption) (*AddRootTargetResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &AddRootTargetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) RemoveRootTarget(ctx context.Context, in *RemoveRootTargetRequest, opts ...dcerpc.CallOption) (*RemoveRootTargetResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &RemoveRootTargetResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) GetSupportedNamespaceVersion(ctx context.Context, in *GetSupportedNamespaceVersionRequest, opts ...dcerpc.CallOption) (*GetSupportedNamespaceVersionResponse, error) {
	op := in.xxx_ToOp(ctx, nil)
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetSupportedNamespaceVersionResponse{}
	out.xxx_FromOp(ctx, op)
	if op.Return != uint32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultNetdfsClient) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultNetdfsClient) Conn() dcerpc.Conn {
	return o.cc
}

func NewNetdfsClient(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (NetdfsClient, error) {
	cc, err := cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(NetdfsSyntaxV3_0))...)
	if err != nil {
		return nil, err
	}
	return &xxx_DefaultNetdfsClient{cc: cc}, nil
}

// xxx_ManagerGetVersionOperation structure represents the NetrDfsManagerGetVersion operation
type xxx_ManagerGetVersionOperation struct {
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_ManagerGetVersionOperation) OpNum() int { return 0 }

func (o *xxx_ManagerGetVersionOperation) OpName() string {
	return "/netdfs/v3/NetrDfsManagerGetVersion"
}

func (o *xxx_ManagerGetVersionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ManagerGetVersionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	return nil
}

func (o *xxx_ManagerGetVersionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	return nil
}

func (o *xxx_ManagerGetVersionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ManagerGetVersionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
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

func (o *xxx_ManagerGetVersionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ManagerGetVersionRequest structure represents the NetrDfsManagerGetVersion operation request
type ManagerGetVersionRequest struct {
}

func (o *ManagerGetVersionRequest) xxx_ToOp(ctx context.Context, op *xxx_ManagerGetVersionOperation) *xxx_ManagerGetVersionOperation {
	if op == nil {
		op = &xxx_ManagerGetVersionOperation{}
	}
	if o == nil {
		return op
	}
	return op
}

func (o *ManagerGetVersionRequest) xxx_FromOp(ctx context.Context, op *xxx_ManagerGetVersionOperation) {
	if o == nil {
		return
	}
}
func (o *ManagerGetVersionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ManagerGetVersionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ManagerGetVersionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ManagerGetVersionResponse structure represents the NetrDfsManagerGetVersion operation response
type ManagerGetVersionResponse struct {
	// Return: The NetrDfsManagerGetVersion return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ManagerGetVersionResponse) xxx_ToOp(ctx context.Context, op *xxx_ManagerGetVersionOperation) *xxx_ManagerGetVersionOperation {
	if op == nil {
		op = &xxx_ManagerGetVersionOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *ManagerGetVersionResponse) xxx_FromOp(ctx context.Context, op *xxx_ManagerGetVersionOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ManagerGetVersionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ManagerGetVersionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ManagerGetVersionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddOperation structure represents the NetrDfsAdd operation
type xxx_AddOperation struct {
	EntryPath  string `idl:"name:DfsEntryPath;string" json:"entry_path"`
	ServerName string `idl:"name:ServerName;string" json:"server_name"`
	ShareName  string `idl:"name:ShareName;string;pointer:unique" json:"share_name"`
	Comment    string `idl:"name:Comment;string;pointer:unique" json:"comment"`
	Flags      uint32 `idl:"name:Flags" json:"flags"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_AddOperation) OpNum() int { return 1 }

func (o *xxx_AddOperation) OpName() string { return "/netdfs/v3/NetrDfsAdd" }

func (o *xxx_AddOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DfsEntryPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.EntryPath); err != nil {
			return err
		}
	}
	// ServerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
			return err
		}
	}
	// ShareName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.ShareName != "" {
			_ptr_ShareName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ShareName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ShareName, _ptr_ShareName); err != nil {
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
	// Comment {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.Comment != "" {
			_ptr_Comment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Comment); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Comment, _ptr_Comment); err != nil {
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

func (o *xxx_AddOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DfsEntryPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.EntryPath); err != nil {
			return err
		}
	}
	// ServerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
			return err
		}
	}
	// ShareName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ShareName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ShareName); err != nil {
				return err
			}
			return nil
		})
		_s_ShareName := func(ptr interface{}) { o.ShareName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ShareName, _s_ShareName, _ptr_ShareName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Comment {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_Comment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Comment); err != nil {
				return err
			}
			return nil
		})
		_s_Comment := func(ptr interface{}) { o.Comment = *ptr.(*string) }
		if err := w.ReadPointer(&o.Comment, _s_Comment, _ptr_Comment); err != nil {
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

func (o *xxx_AddOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddRequest structure represents the NetrDfsAdd operation request
type AddRequest struct {
	// DfsEntryPath: The pointer to a DFS link path that contains the name of an existing
	// link when additional link targets are being added or the name of a new link is being
	// created.
	EntryPath string `idl:"name:DfsEntryPath;string" json:"entry_path"`
	// ServerName: The pointer to a null-terminated Unicode string that specifies the DFS
	// link target host name.
	ServerName string `idl:"name:ServerName;string" json:"server_name"`
	// ShareName: The pointer to a null-terminated Unicode DFS link target share name string.
	// This can also be a share name with a path relative to the share, for example, "share1\mydir1\mydir2".
	// When specified this way, each pathname component MUST be a directory.
	ShareName string `idl:"name:ShareName;string;pointer:unique" json:"share_name"`
	// Comment: The pointer to a null-terminated Unicode string that contains a comment
	// associated with this root or link. This string has no protocol-specified restrictions
	// on length or content. The comment is meant for human consumption and does not affect
	// server functionality. The comment MUST be ignored when adding a target to an existing
	// link.
	Comment string `idl:"name:Comment;string;pointer:unique" json:"comment"`
	// Flags: A value indicating the operation to perform. The following table lists such
	// flags.
	//
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	|                               |                                                                                  |
	//	|             VALUE             |                                     MEANING                                      |
	//	|                               |                                                                                  |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000                    | Create a new link or adds a new target to an existing link.                      |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_ADD_VOLUME 0x00000001     | Create a new link in the DFS namespace if one does not already exist or fails if |
	//	|                               | a link already exists.                                                           |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_RESTORE_VOLUME 0x00000002 | Add a target without verifying its existence.                                    |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
}

func (o *AddRequest) xxx_ToOp(ctx context.Context, op *xxx_AddOperation) *xxx_AddOperation {
	if op == nil {
		op = &xxx_AddOperation{}
	}
	if o == nil {
		return op
	}
	op.EntryPath = o.EntryPath
	op.ServerName = o.ServerName
	op.ShareName = o.ShareName
	op.Comment = o.Comment
	op.Flags = o.Flags
	return op
}

func (o *AddRequest) xxx_FromOp(ctx context.Context, op *xxx_AddOperation) {
	if o == nil {
		return
	}
	o.EntryPath = op.EntryPath
	o.ServerName = op.ServerName
	o.ShareName = op.ShareName
	o.Comment = op.Comment
	o.Flags = op.Flags
}
func (o *AddRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AddRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddResponse structure represents the NetrDfsAdd operation response
type AddResponse struct {
	// Return: The NetrDfsAdd return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AddResponse) xxx_ToOp(ctx context.Context, op *xxx_AddOperation) *xxx_AddOperation {
	if op == nil {
		op = &xxx_AddOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *AddResponse) xxx_FromOp(ctx context.Context, op *xxx_AddOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AddResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AddResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveOperation structure represents the NetrDfsRemove operation
type xxx_RemoveOperation struct {
	EntryPath  string `idl:"name:DfsEntryPath;string" json:"entry_path"`
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	ShareName  string `idl:"name:ShareName;string;pointer:unique" json:"share_name"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveOperation) OpNum() int { return 2 }

func (o *xxx_RemoveOperation) OpName() string { return "/netdfs/v3/NetrDfsRemove" }

func (o *xxx_RemoveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DfsEntryPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.EntryPath); err != nil {
			return err
		}
	}
	// ServerName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
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
	// ShareName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.ShareName != "" {
			_ptr_ShareName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ShareName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ShareName, _ptr_ShareName); err != nil {
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

func (o *xxx_RemoveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DfsEntryPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.EntryPath); err != nil {
			return err
		}
	}
	// ServerName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
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
	// ShareName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ShareName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ShareName); err != nil {
				return err
			}
			return nil
		})
		_s_ShareName := func(ptr interface{}) { o.ShareName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ShareName, _s_ShareName, _ptr_ShareName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveRequest structure represents the NetrDfsRemove operation request
type RemoveRequest struct {
	// DfsEntryPath: The pointer to the DFS link path that contains the name of an existing
	// link.
	EntryPath string `idl:"name:DfsEntryPath;string" json:"entry_path"`
	// ServerName: The pointer to a null-terminated Unicode DFS link target host name string.
	// Clients MUST set ServerName to a NULL pointer in requests to remove the link and
	// all its link targets.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// ShareName: The pointer to a null-terminated Unicode DFS link target share name string.
	// This can also be a share name with a path relative to the share, for example, "share1\mydir1\mydir2".
	// Clients MUST set ShareName to a NULL pointer in requests to remove the link and all
	// its link targets.
	ShareName string `idl:"name:ShareName;string;pointer:unique" json:"share_name"`
}

func (o *RemoveRequest) xxx_ToOp(ctx context.Context, op *xxx_RemoveOperation) *xxx_RemoveOperation {
	if op == nil {
		op = &xxx_RemoveOperation{}
	}
	if o == nil {
		return op
	}
	op.EntryPath = o.EntryPath
	op.ServerName = o.ServerName
	op.ShareName = o.ShareName
	return op
}

func (o *RemoveRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveOperation) {
	if o == nil {
		return
	}
	o.EntryPath = op.EntryPath
	o.ServerName = op.ServerName
	o.ShareName = op.ShareName
}
func (o *RemoveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemoveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveResponse structure represents the NetrDfsRemove operation response
type RemoveResponse struct {
	// Return: The NetrDfsRemove return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RemoveResponse) xxx_ToOp(ctx context.Context, op *xxx_RemoveOperation) *xxx_RemoveOperation {
	if op == nil {
		op = &xxx_RemoveOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RemoveResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RemoveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemoveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetInfoOperation structure represents the NetrDfsSetInfo operation
type xxx_SetInfoOperation struct {
	EntryPath  string `idl:"name:DfsEntryPath;string" json:"entry_path"`
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	ShareName  string `idl:"name:ShareName;string;pointer:unique" json:"share_name"`
	Level      uint32 `idl:"name:Level" json:"level"`
	Info       *Info  `idl:"name:DfsInfo;switch_is:Level" json:"info"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SetInfoOperation) OpNum() int { return 3 }

func (o *xxx_SetInfoOperation) OpName() string { return "/netdfs/v3/NetrDfsSetInfo" }

func (o *xxx_SetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DfsEntryPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.EntryPath); err != nil {
			return err
		}
	}
	// ServerName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
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
	// ShareName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.ShareName != "" {
			_ptr_ShareName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ShareName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ShareName, _ptr_ShareName); err != nil {
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
	// DfsInfo {in} (1:{pointer=ref}*(1))(2:{switch_type={}(uint32), alias=DFS_INFO_STRUCT}(union))
	{
		_swInfo := uint32(o.Level)
		if o.Info != nil {
			if err := o.Info.MarshalUnionNDR(ctx, w, _swInfo); err != nil {
				return err
			}
		} else {
			if err := (&Info{}).MarshalUnionNDR(ctx, w, _swInfo); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DfsEntryPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.EntryPath); err != nil {
			return err
		}
	}
	// ServerName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
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
	// ShareName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ShareName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ShareName); err != nil {
				return err
			}
			return nil
		})
		_s_ShareName := func(ptr interface{}) { o.ShareName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ShareName, _s_ShareName, _ptr_ShareName); err != nil {
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
	// DfsInfo {in} (1:{pointer=ref}*(1))(2:{switch_type={}(uint32), alias=DFS_INFO_STRUCT}(union))
	{
		if o.Info == nil {
			o.Info = &Info{}
		}
		_swInfo := uint32(o.Level)
		if err := o.Info.UnmarshalUnionNDR(ctx, w, _swInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetInfoRequest structure represents the NetrDfsSetInfo operation request
type SetInfoRequest struct {
	// DfsEntryPath: The pointer to a DFS root or a DFS link path.
	EntryPath string `idl:"name:DfsEntryPath;string" json:"entry_path"`
	// ServerName: The pointer to a null-terminated Unicode DFS root target or DFS link
	// target host name string. Clients MUST set this to a NULL pointer when the DFS root
	// or DFS link is used and not the DFS root target or DFS link target.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// ShareName: The pointer to a null-terminated Unicode string DFS root target or DFS
	// link target host name. Clients MUST set this to a NULL pointer when the DFS root
	// or DFS link is used and not the DFS root target or DFS link target.
	ShareName string `idl:"name:ShareName;string;pointer:unique" json:"share_name"`
	// Level: Specifies the information level of the data and, in turn, determines the action
	// the method performs.
	//
	//	+----------------------+----------------------------------------------------------------------------------+
	//	|                      |                                                                                  |
	//	|        VALUE         |                                     MEANING                                      |
	//	|                      |                                                                                  |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_100 0x00000064 | Sets the comment associated with the root or link specified in the DfsInfo       |
	//	|                      | parameter. The ServerName and ShareName parameters MUST be NULL.                 |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_101 0x00000065 | Sets the state associated with the root, link, root target, or link target       |
	//	|                      | specified in DfsInfo.<52>                                                        |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_102 0x00000066 | Sets the time-out value associated with the root or link specified in DfsInfo.   |
	//	|                      | The ServerName and ShareName parameters MUST be ignored.                         |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_103 0x00000067 | Sets the property flags for the root or link specified in DfsInfo. The           |
	//	|                      | ServerName and ShareName parameters MUST be NULL.                                |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_104 0x00000068 | Sets the target priority rank and class for the root target or link target       |
	//	|                      | specified in DfsInfo.                                                            |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_105 0x00000069 | Sets the comment, state, time-out information, and property flags for the        |
	//	|                      | namespace root or link specified in DfsInfo. Does not apply to a root target or  |
	//	|                      | link target. The ServerName and ShareName parameters MUST be NULL.               |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_106 0x0000006A | Sets the target state and priority for the DFS root target or DFS link target    |
	//	|                      | specified in DfsInfo.<53> This does not apply to the DFS namespace root or link. |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_107 0x0000006B | Sets the comment, state, time-out, security descriptor information, and property |
	//	|                      | flags for the namespace root or link specified in DfsInfo. Does not apply to     |
	//	|                      | a root target or link target. The ServerName and ShareName parameters MUST be    |
	//	|                      | NULL. The security descriptor MUST NOT have owner, group, or SACLs in it. The    |
	//	|                      | security descriptor MUST be a NULL, zero length value if used on a namespace     |
	//	|                      | root. In this case, note that it is equivalent to using Level_105.               |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_150 0x00000096 | Sets the security descriptor associated with a DFS link. Only stand-alone DFS    |
	//	|                      | namespaces and domainv2-based DFS namespaces are supported. The ServerName and   |
	//	|                      | ShareName parameters MUST both be NULL. The security descriptor MUST NOT have    |
	//	|                      | owner, group, or SACLs in it.                                                    |
	//	+----------------------+----------------------------------------------------------------------------------+
	Level uint32 `idl:"name:Level" json:"level"`
	// DfsInfo: The pointer to a DFS_INFO_STRUCT union that contains the specified data.
	// The value of the Level parameter selects the case of the union.
	Info *Info `idl:"name:DfsInfo;switch_is:Level" json:"info"`
}

func (o *SetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_SetInfoOperation) *xxx_SetInfoOperation {
	if op == nil {
		op = &xxx_SetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.EntryPath = o.EntryPath
	op.ServerName = o.ServerName
	op.ShareName = o.ShareName
	op.Level = o.Level
	op.Info = o.Info
	return op
}

func (o *SetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_SetInfoOperation) {
	if o == nil {
		return
	}
	o.EntryPath = op.EntryPath
	o.ServerName = op.ServerName
	o.ShareName = op.ShareName
	o.Level = op.Level
	o.Info = op.Info
}
func (o *SetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetInfoResponse structure represents the NetrDfsSetInfo operation response
type SetInfoResponse struct {
	// Return: The NetrDfsSetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_SetInfoOperation) *xxx_SetInfoOperation {
	if op == nil {
		op = &xxx_SetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_SetInfoOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetInfoOperation structure represents the NetrDfsGetInfo operation
type xxx_GetInfoOperation struct {
	EntryPath  string `idl:"name:DfsEntryPath;string" json:"entry_path"`
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	ShareName  string `idl:"name:ShareName;string;pointer:unique" json:"share_name"`
	Level      uint32 `idl:"name:Level" json:"level"`
	Info       *Info  `idl:"name:DfsInfo;switch_is:Level" json:"info"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetInfoOperation) OpNum() int { return 4 }

func (o *xxx_GetInfoOperation) OpName() string { return "/netdfs/v3/NetrDfsGetInfo" }

func (o *xxx_GetInfoOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInfoOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DfsEntryPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.EntryPath); err != nil {
			return err
		}
	}
	// ServerName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
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
	// ShareName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.ShareName != "" {
			_ptr_ShareName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ShareName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ShareName, _ptr_ShareName); err != nil {
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
	return nil
}

func (o *xxx_GetInfoOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DfsEntryPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.EntryPath); err != nil {
			return err
		}
	}
	// ServerName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
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
	// ShareName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ShareName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ShareName); err != nil {
				return err
			}
			return nil
		})
		_s_ShareName := func(ptr interface{}) { o.ShareName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ShareName, _s_ShareName, _ptr_ShareName); err != nil {
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
	return nil
}

func (o *xxx_GetInfoOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInfoOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// DfsInfo {out} (1:{pointer=ref}*(1))(2:{switch_type={}(uint32), alias=DFS_INFO_STRUCT}(union))
	{
		_swInfo := uint32(o.Level)
		if o.Info != nil {
			if err := o.Info.MarshalUnionNDR(ctx, w, _swInfo); err != nil {
				return err
			}
		} else {
			if err := (&Info{}).MarshalUnionNDR(ctx, w, _swInfo); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetInfoOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// DfsInfo {out} (1:{pointer=ref}*(1))(2:{switch_type={}(uint32), alias=DFS_INFO_STRUCT}(union))
	{
		if o.Info == nil {
			o.Info = &Info{}
		}
		_swInfo := uint32(o.Level)
		if err := o.Info.UnmarshalUnionNDR(ctx, w, _swInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetInfoRequest structure represents the NetrDfsGetInfo operation request
type GetInfoRequest struct {
	// DfsEntryPath: The pointer to a DFS root or a DFS link path.
	EntryPath string `idl:"name:DfsEntryPath;string" json:"entry_path"`
	// ServerName: This parameter MUST be a NULL pointer for Level_50 and MUST be ignored
	// for other levels.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// ShareName: This parameter MUST be a NULL pointer for Level_50 and MUST be ignored
	// for other levels.
	ShareName string `idl:"name:ShareName;string;pointer:unique" json:"share_name"`
	// Level: This parameter specifies the information level of the data and, in turn, determines
	// the action the method performs.
	//
	//	+----------------------+----------------------------------------------------------------------------------+
	//	|                      |                                                                                  |
	//	|        VALUE         |                                     MEANING                                      |
	//	|                      |                                                                                  |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_1 0x00000001   | Returns the name of the DFS root or the DFS link.                                |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_2 0x00000002   | Returns the name, comment, state, and number of targets for the DFS root or the  |
	//	|                      | DFS link.                                                                        |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_3 0x00000003   | Returns the name, comment, state, number of targets, and target information for  |
	//	|                      | the DFS root or the DFS link.                                                    |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_4 0x00000004   | Returns the name, comment, state, time-out, GUID, number of targets, and target  |
	//	|                      | information for the DFS root or the DFS link.                                    |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_5 0x00000005   | Returns the name, comment, state, time-out, GUID, property flags, metadata size, |
	//	|                      | and number of targets for the DFS root or the DFS link.                          |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_6 0x00000006   | Returns the name, comment, state, GUID, time-out, property flags, metadata size, |
	//	|                      | number of targets, and target information for the DFS root or the DFS link.      |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_7 0x00000007   | Returns the version number GUID of the DFS metadata. This value only applies to  |
	//	|                      | the DFS root.                                                                    |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_8 0x00000008   | Returns the name, comment, state, time-out, GUID, property flags, metadata size, |
	//	|                      | number of targets, and security descriptor associated with the DFS root or the   |
	//	|                      | DFS link. Only stand-alone DFS namespaces and domainv2-based DFS namespaces are  |
	//	|                      | supported.                                                                       |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_9 0x00000009   | Returns the name, comment, state, GUID, time-out, property flags, metadata size, |
	//	|                      | number of targets, list of targets, and security descriptor for the DFS root or  |
	//	|                      | the DFS link. Only stand-alone DFS namespaces and domainv2-based DFS namespaces  |
	//	|                      | are supported.                                                                   |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_50 0x00000032  | Returns the DFS metadata version and capability information of an existing DFS   |
	//	|                      | namespace. This level is valid only for the DFS namespace root, not for DFS      |
	//	|                      | links. The ServerName and ShareName parameters MUST both be NULL.<58>            |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_100 0x00000064 | Returns the comment associated with the root or DFS link specified in the        |
	//	|                      | DfsEntryPath parameter.                                                          |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_150 0x00000096 | Returns the security descriptor associated with a DFS link. Only stand-alone DFS |
	//	|                      | namespaces and domainv2-based DFS namespaces are supported.                      |
	//	+----------------------+----------------------------------------------------------------------------------+
	Level uint32 `idl:"name:Level" json:"level"`
}

func (o *GetInfoRequest) xxx_ToOp(ctx context.Context, op *xxx_GetInfoOperation) *xxx_GetInfoOperation {
	if op == nil {
		op = &xxx_GetInfoOperation{}
	}
	if o == nil {
		return op
	}
	op.EntryPath = o.EntryPath
	op.ServerName = o.ServerName
	op.ShareName = o.ShareName
	op.Level = o.Level
	return op
}

func (o *GetInfoRequest) xxx_FromOp(ctx context.Context, op *xxx_GetInfoOperation) {
	if o == nil {
		return
	}
	o.EntryPath = op.EntryPath
	o.ServerName = op.ServerName
	o.ShareName = op.ShareName
	o.Level = op.Level
}
func (o *GetInfoRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetInfoRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInfoOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetInfoResponse structure represents the NetrDfsGetInfo operation response
type GetInfoResponse struct {
	// XXX: Level is an implicit input depedency for output parameters
	Level uint32 `idl:"name:Level" json:"level"`
	// DfsInfo: The pointer to a DFS_INFO_STRUCT union to receive the returned information.
	// The case of the union is selected by the value of the Level parameter.
	Info *Info `idl:"name:DfsInfo;switch_is:Level" json:"info"`
	// Return: The NetrDfsGetInfo return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetInfoResponse) xxx_ToOp(ctx context.Context, op *xxx_GetInfoOperation) *xxx_GetInfoOperation {
	if op == nil {
		op = &xxx_GetInfoOperation{}
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

func (o *GetInfoResponse) xxx_FromOp(ctx context.Context, op *xxx_GetInfoOperation) {
	if o == nil {
		return
	}
	// XXX: implicit input dependencies for output parameters
	o.Level = op.Level

	o.Info = op.Info
	o.Return = op.Return
}
func (o *GetInfoResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetInfoResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetInfoOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumOperation structure represents the NetrDfsEnum operation
type xxx_EnumOperation struct {
	Level         uint32    `idl:"name:Level" json:"level"`
	PrefMaxLength uint32    `idl:"name:PrefMaxLen" json:"pref_max_length"`
	Enum          *InfoEnum `idl:"name:DfsEnum;pointer:unique" json:"enum"`
	Resume        uint32    `idl:"name:ResumeHandle;pointer:unique" json:"resume"`
	Return        uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumOperation) OpNum() int { return 5 }

func (o *xxx_EnumOperation) OpName() string { return "/netdfs/v3/NetrDfsEnum" }

func (o *xxx_EnumOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Level {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// PrefMaxLen {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PrefMaxLength); err != nil {
			return err
		}
	}
	// DfsEnum {in, out} (1:{pointer=unique}*(1))(2:{alias=DFS_INFO_ENUM_STRUCT}(struct))
	{
		if o.Enum != nil {
			_ptr_DfsEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&InfoEnum{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_DfsEnum); err != nil {
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
	// ResumeHandle {in, out} (1:{pointer=unique}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_ResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Level {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// PrefMaxLen {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PrefMaxLength); err != nil {
			return err
		}
	}
	// DfsEnum {in, out} (1:{pointer=unique}*(1))(2:{alias=DFS_INFO_ENUM_STRUCT}(struct))
	{
		_ptr_DfsEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &InfoEnum{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_DfsEnum := func(ptr interface{}) { o.Enum = *ptr.(**InfoEnum) }
		if err := w.ReadPointer(&o.Enum, _s_DfsEnum, _ptr_DfsEnum); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=unique}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_ResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_ResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_ResumeHandle, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// DfsEnum {in, out} (1:{pointer=unique}*(1))(2:{alias=DFS_INFO_ENUM_STRUCT}(struct))
	{
		if o.Enum != nil {
			_ptr_DfsEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&InfoEnum{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_DfsEnum); err != nil {
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
	// ResumeHandle {in, out} (1:{pointer=unique}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_ResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// DfsEnum {in, out} (1:{pointer=unique}*(1))(2:{alias=DFS_INFO_ENUM_STRUCT}(struct))
	{
		_ptr_DfsEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &InfoEnum{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_DfsEnum := func(ptr interface{}) { o.Enum = *ptr.(**InfoEnum) }
		if err := w.ReadPointer(&o.Enum, _s_DfsEnum, _ptr_DfsEnum); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=unique}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_ResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_ResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_ResumeHandle, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumRequest structure represents the NetrDfsEnum operation request
type EnumRequest struct {
	// Level: This parameter specifies the information level of the data and, in turn, determines
	// the action that the method performs. On successful return, the server MUST return
	// an array of the corresponding structures in the buffer pointed to by DfsEnum.
	//
	//	+--------------------+----------------------------------------------------------------------------------+
	//	|                    |                                                                                  |
	//	|       VALUE        |                                     MEANING                                      |
	//	|                    |                                                                                  |
	//	+--------------------+----------------------------------------------------------------------------------+
	//	+--------------------+----------------------------------------------------------------------------------+
	//	| Level_1 0x00000001 | Gets the name of the DFS root and all links beneath it. In this case, on         |
	//	|                    | successful return DfsEnum MUST point to an array of DFS_INFO_1 structures.       |
	//	+--------------------+----------------------------------------------------------------------------------+
	//	| Level_2 0x00000002 | Gets the name, comment, state, and number of targets for the DFS root and all    |
	//	|                    | links under the root. In this case, on successful return DfsEnum MUST point to   |
	//	|                    | an array of DFS_INFO_2 structures.                                               |
	//	+--------------------+----------------------------------------------------------------------------------+
	//	| Level_3 0x00000003 | Gets the name, comment, state, number of targets, and target information for     |
	//	|                    | the DFS root and all links under the root. In this case, on successful return    |
	//	|                    | DfsEnum MUST point to an array of DFS_INFO_3 structures.                         |
	//	+--------------------+----------------------------------------------------------------------------------+
	//	| Level_4 0x00000004 | Gets the name, comment, state, time-out, GUID, number of targets, and target     |
	//	|                    | information for the DFS root and all links under the root. In this case, on      |
	//	|                    | successful return DfsEnum MUST point to an array of DFS_INFO_4 structures.       |
	//	+--------------------+----------------------------------------------------------------------------------+
	//	| Level_5 0x00000005 | Gets the name, comment, state, time-out, GUID, property flags, metadata size,    |
	//	|                    | and number of targets for a DFS root and all links under the root. In this case, |
	//	|                    | on successful return DfsEnum MUST point to an array of DFS_INFO_5 structures.    |
	//	+--------------------+----------------------------------------------------------------------------------+
	//	| Level_6 0x00000006 | Gets the name, comment, state, time-out, GUID, property flags, metadata size,    |
	//	|                    | number of targets, and target information for a DFS root or DFS links. In        |
	//	|                    | this case, on successful return DfsEnum MUST point to an array of DFS_INFO_6     |
	//	|                    | structures.                                                                      |
	//	+--------------------+----------------------------------------------------------------------------------+
	//	| Level_8 0x00000008 | Gets the name, comment, state, time-out, GUID, property flags, metadata size,    |
	//	|                    | and number of targets for a DFS root and all DFS links under the root. Also      |
	//	|                    | returns the security descriptor associated with each of the DFS links. In        |
	//	|                    | this case, on successful return DfsEnum MUST point to an array of DFS_INFO_8     |
	//	|                    | structures.                                                                      |
	//	+--------------------+----------------------------------------------------------------------------------+
	//	| Level_9 0x00000009 | Gets the name, comment, state, time-out, GUID, property flags, metadata size,    |
	//	|                    | and number of targets, and target information for a DFS root and all DFS links   |
	//	|                    | under the root. Also returns the security descriptor associated with each of the |
	//	|                    | DFS links. In this case, on successful return DfsEnum MUST point to an array of  |
	//	|                    | DFS_INFO_9 structures.                                                           |
	//	+--------------------+----------------------------------------------------------------------------------+
	Level uint32 `idl:"name:Level" json:"level"`
	// PrefMaxLen: This parameter specifies restrictions on the number of elements returned.
	// A value of 0xFFFFFFFF means there are no restrictions, in which case all entries
	// MUST be returned.<64>
	PrefMaxLength uint32 `idl:"name:PrefMaxLen" json:"pref_max_length"`
	// DfsEnum: A pointer to a DFS_INFO_ENUM_STRUCT union to receive the returned information.
	// The client SHOULD set the Level member to the same value as the method's Level parameter,
	// and MUST set the DfsInfoContainer union member to a pointer to the corresponding
	// container structure as specified in section 2.2.6. The client MUST initialize the
	// container structure's EntriesRead member to zero and Buffer member to a NULL pointer.
	// The value of the Level member determines the case of the union.
	Enum *InfoEnum `idl:"name:DfsEnum;pointer:unique" json:"enum"`
	// ResumeHandle: This parameter is used to continue an enumeration when more data is
	// available than can be returned in a single invocation of this method.
	//
	// * If this parameter is not a NULL pointer, and the method returns ERROR_SUCCESS,
	// this parameter receives an implementation-specific nonzero value that can be passed
	// in subsequent calls to this method to continue the enumeration.
	//
	// * If this parameter is a NULL pointer or points to a 0 value, it indicates that this
	// is an initial enumeration request.
	//
	// * If this parameter is not a NULL pointer and points to a nonzero value returned
	// in ResumeHandle by an earlier invocation of this method, the server will attempt
	// to continue a previous enumeration, but MAY produce incomplete or inconsistent results
	// due to the possibility of concurrent updates to the DFS namespace ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_6a3f0be9-b9b4-49df-9d1c-a3b89e4e9890
	// ). <65> ( 3a466440-1ef6-4439-b4e2-be9eaddeb511#Appendix_A_65 )
	Resume uint32 `idl:"name:ResumeHandle;pointer:unique" json:"resume"`
}

func (o *EnumRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumOperation) *xxx_EnumOperation {
	if op == nil {
		op = &xxx_EnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Level = o.Level
	op.PrefMaxLength = o.PrefMaxLength
	op.Enum = o.Enum
	op.Resume = o.Resume
	return op
}

func (o *EnumRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumOperation) {
	if o == nil {
		return
	}
	o.Level = op.Level
	o.PrefMaxLength = op.PrefMaxLength
	o.Enum = op.Enum
	o.Resume = op.Resume
}
func (o *EnumRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumResponse structure represents the NetrDfsEnum operation response
type EnumResponse struct {
	// DfsEnum: A pointer to a DFS_INFO_ENUM_STRUCT union to receive the returned information.
	// The client SHOULD set the Level member to the same value as the method's Level parameter,
	// and MUST set the DfsInfoContainer union member to a pointer to the corresponding
	// container structure as specified in section 2.2.6. The client MUST initialize the
	// container structure's EntriesRead member to zero and Buffer member to a NULL pointer.
	// The value of the Level member determines the case of the union.
	Enum *InfoEnum `idl:"name:DfsEnum;pointer:unique" json:"enum"`
	// ResumeHandle: This parameter is used to continue an enumeration when more data is
	// available than can be returned in a single invocation of this method.
	//
	// * If this parameter is not a NULL pointer, and the method returns ERROR_SUCCESS,
	// this parameter receives an implementation-specific nonzero value that can be passed
	// in subsequent calls to this method to continue the enumeration.
	//
	// * If this parameter is a NULL pointer or points to a 0 value, it indicates that this
	// is an initial enumeration request.
	//
	// * If this parameter is not a NULL pointer and points to a nonzero value returned
	// in ResumeHandle by an earlier invocation of this method, the server will attempt
	// to continue a previous enumeration, but MAY produce incomplete or inconsistent results
	// due to the possibility of concurrent updates to the DFS namespace ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_6a3f0be9-b9b4-49df-9d1c-a3b89e4e9890
	// ). <65> ( 3a466440-1ef6-4439-b4e2-be9eaddeb511#Appendix_A_65 )
	Resume uint32 `idl:"name:ResumeHandle;pointer:unique" json:"resume"`
	// Return: The NetrDfsEnum return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EnumResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumOperation) *xxx_EnumOperation {
	if op == nil {
		op = &xxx_EnumOperation{}
	}
	if o == nil {
		return op
	}
	op.Enum = o.Enum
	op.Resume = o.Resume
	op.Return = o.Return
	return op
}

func (o *EnumResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumOperation) {
	if o == nil {
		return
	}
	o.Enum = op.Enum
	o.Resume = op.Resume
	o.Return = op.Return
}
func (o *EnumResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_MoveOperation structure represents the NetrDfsMove operation
type xxx_MoveOperation struct {
	EntryPath       string `idl:"name:DfsEntryPath;string" json:"entry_path"`
	NewDFSEntryPath string `idl:"name:NewDfsEntryPath;string" json:"new_dfs_entry_path"`
	Flags           uint32 `idl:"name:Flags" json:"flags"`
	Return          uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_MoveOperation) OpNum() int { return 6 }

func (o *xxx_MoveOperation) OpName() string { return "/netdfs/v3/NetrDfsMove" }

func (o *xxx_MoveOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DfsEntryPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.EntryPath); err != nil {
			return err
		}
	}
	// NewDfsEntryPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.NewDFSEntryPath); err != nil {
			return err
		}
	}
	// Flags {in} (1:(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DfsEntryPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.EntryPath); err != nil {
			return err
		}
	}
	// NewDfsEntryPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.NewDFSEntryPath); err != nil {
			return err
		}
	}
	// Flags {in} (1:(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_MoveOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// MoveRequest structure represents the NetrDfsMove operation request
type MoveRequest struct {
	// DfsEntryPath: The pointer to a DFS path, this parameter specifies the source path
	// for the move operation. This MUST be a DFS link or the path prefix of any DFS link
	// in the DFS namespace.
	EntryPath string `idl:"name:DfsEntryPath;string" json:"entry_path"`
	// NewDfsEntryPath: The pointer to a DFS path, this parameter specifies the destination
	// DFS path for the move operation. This MUST be a path or a DFS link in the same DFS
	// namespace.
	NewDFSEntryPath string `idl:"name:NewDfsEntryPath;string" json:"new_dfs_entry_path"`
	// Flags: A bit field specifying additional actions to take.
	//
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                            |                                                                                  |
	//	|                   VALUE                    |                                     MEANING                                      |
	//	|                                            |                                                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_MOVE_FLAG_REPLACE_IF_EXISTS 0x00000001 | If the destination path is an existing link, replace it as part of the move      |
	//	|                                            | operation.                                                                       |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
}

func (o *MoveRequest) xxx_ToOp(ctx context.Context, op *xxx_MoveOperation) *xxx_MoveOperation {
	if op == nil {
		op = &xxx_MoveOperation{}
	}
	if o == nil {
		return op
	}
	op.EntryPath = o.EntryPath
	op.NewDFSEntryPath = o.NewDFSEntryPath
	op.Flags = o.Flags
	return op
}

func (o *MoveRequest) xxx_FromOp(ctx context.Context, op *xxx_MoveOperation) {
	if o == nil {
		return
	}
	o.EntryPath = op.EntryPath
	o.NewDFSEntryPath = op.NewDFSEntryPath
	o.Flags = op.Flags
}
func (o *MoveRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *MoveRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MoveOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// MoveResponse structure represents the NetrDfsMove operation response
type MoveResponse struct {
	// Return: The NetrDfsMove return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *MoveResponse) xxx_ToOp(ctx context.Context, op *xxx_MoveOperation) *xxx_MoveOperation {
	if op == nil {
		op = &xxx_MoveOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *MoveResponse) xxx_FromOp(ctx context.Context, op *xxx_MoveOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *MoveResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *MoveResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_MoveOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddFTRootOperation structure represents the NetrDfsAddFtRoot operation
type xxx_AddFTRootOperation struct {
	ServerName string    `idl:"name:ServerName;string" json:"server_name"`
	DCName     string    `idl:"name:DcName;string" json:"dc_name"`
	RootShare  string    `idl:"name:RootShare;string" json:"root_share"`
	FTDFSName  string    `idl:"name:FtDfsName;string" json:"ft_dfs_name"`
	Comment    string    `idl:"name:Comment;string" json:"comment"`
	ConfigDN   string    `idl:"name:ConfigDN;string" json:"config_dn"`
	NewFTDFS   bool      `idl:"name:NewFtDfs" json:"new_ft_dfs"`
	Flags      uint32    `idl:"name:ApiFlags" json:"flags"`
	RootList   *RootList `idl:"name:ppRootList;pointer:unique" json:"root_list"`
	Return     uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_AddFTRootOperation) OpNum() int { return 10 }

func (o *xxx_AddFTRootOperation) OpName() string { return "/netdfs/v3/NetrDfsAddFtRoot" }

func (o *xxx_AddFTRootOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddFTRootOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
			return err
		}
	}
	// DcName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DCName); err != nil {
			return err
		}
	}
	// RootShare {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.RootShare); err != nil {
			return err
		}
	}
	// FtDfsName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FTDFSName); err != nil {
			return err
		}
	}
	// Comment {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Comment); err != nil {
			return err
		}
	}
	// ConfigDN {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ConfigDN); err != nil {
			return err
		}
	}
	// NewFtDfs {in} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.NewFTDFS); err != nil {
			return err
		}
	}
	// ApiFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// ppRootList {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=DFSM_ROOT_LIST}(struct))
	{
		if o.RootList != nil {
			_ptr_ppRootList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RootList != nil {
					_ptr_ppRootList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.RootList != nil {
							if err := o.RootList.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&RootList{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.RootList, _ptr_ppRootList); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RootList, _ptr_ppRootList); err != nil {
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

func (o *xxx_AddFTRootOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
			return err
		}
	}
	// DcName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DCName); err != nil {
			return err
		}
	}
	// RootShare {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.RootShare); err != nil {
			return err
		}
	}
	// FtDfsName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FTDFSName); err != nil {
			return err
		}
	}
	// Comment {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Comment); err != nil {
			return err
		}
	}
	// ConfigDN {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ConfigDN); err != nil {
			return err
		}
	}
	// NewFtDfs {in} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.NewFTDFS); err != nil {
			return err
		}
	}
	// ApiFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// ppRootList {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=DFSM_ROOT_LIST}(struct))
	{
		_ptr_ppRootList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppRootList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.RootList == nil {
					o.RootList = &RootList{}
				}
				if err := o.RootList.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppRootList := func(ptr interface{}) { o.RootList = *ptr.(**RootList) }
			if err := w.ReadPointer(&o.RootList, _s_ppRootList, _ptr_ppRootList); err != nil {
				return err
			}
			return nil
		})
		_s_ppRootList := func(ptr interface{}) { o.RootList = *ptr.(**RootList) }
		if err := w.ReadPointer(&o.RootList, _s_ppRootList, _ptr_ppRootList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddFTRootOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddFTRootOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppRootList {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=DFSM_ROOT_LIST}(struct))
	{
		if o.RootList != nil {
			_ptr_ppRootList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RootList != nil {
					_ptr_ppRootList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.RootList != nil {
							if err := o.RootList.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&RootList{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.RootList, _ptr_ppRootList); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RootList, _ptr_ppRootList); err != nil {
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
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddFTRootOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppRootList {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=DFSM_ROOT_LIST}(struct))
	{
		_ptr_ppRootList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppRootList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.RootList == nil {
					o.RootList = &RootList{}
				}
				if err := o.RootList.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppRootList := func(ptr interface{}) { o.RootList = *ptr.(**RootList) }
			if err := w.ReadPointer(&o.RootList, _s_ppRootList, _ptr_ppRootList); err != nil {
				return err
			}
			return nil
		})
		_s_ppRootList := func(ptr interface{}) { o.RootList = *ptr.(**RootList) }
		if err := w.ReadPointer(&o.RootList, _s_ppRootList, _ptr_ppRootList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddFTRootRequest structure represents the NetrDfsAddFtRoot operation request
type AddFTRootRequest struct {
	// ServerName: The pointer to a null-terminated Unicode string. This MUST be used as
	// the host name of the new DFS root target in the metadata.<108>
	ServerName string `idl:"name:ServerName;string" json:"server_name"`
	// DcName: The pointer to a null-terminated Unicode string. For a domainv1-based DFS
	// namespace, this string contains the host name of the DC that the new DFS root target
	// is to use to get or update DFS metadata for the DFS namespace. This parameter MAY
	// be a NULL pointer, otherwise, it MUST be the PDC for the domain of the DFS namespace.
	DCName string `idl:"name:DcName;string" json:"dc_name"`
	// RootShare: The pointer to a null-terminated Unicode string. This is the new DFS root
	// target share name. This can be different from the FtDfsName parameter. The share
	// MUST already exist.
	RootShare string `idl:"name:RootShare;string" json:"root_share"`
	// FtDfsName: The pointer to a null-terminated Unicode string. This is the name of the
	// new or existing domain-based DFS namespace.
	FTDFSName string `idl:"name:FtDfsName;string" json:"ft_dfs_name"`
	// Comment: The pointer to a null-terminated Unicode string that contains a comment
	// associated with the DFS namespace. Used for informational purposes, this string has
	// no protocol-specified restrictions on length or content. The comment is meant for
	// human consumption and does not affect server functionality. This parameter MAY be
	// NULL.
	Comment string `idl:"name:Comment;string" json:"comment"`
	// ConfigDN: The pointer to a null-terminated Unicode string. This string MUST be the
	// path of the DFS namespace object entry in the DFS Configuration Container (see section
	// 2.3.3).<109>
	ConfigDN string `idl:"name:ConfigDN;string" json:"config_dn"`
	// NewFtDfs: A Boolean value that, if TRUE, indicates a request to create a new root.
	// If FALSE, then this value indicates a request to add a new root target to an existing
	// root.
	NewFTDFS bool `idl:"name:NewFtDfs" json:"new_ft_dfs"`
	// ApiFlags: This parameter MUST be 0.
	Flags uint32 `idl:"name:ApiFlags" json:"flags"`
	// ppRootList: On success, returns a list of DFS root targets in the domain-based DFS
	// namespace that the client will be responsible for notifying of the change in the
	// DFS namespace. See section 3.2.4.3.1. The list MAY be empty if the server has performed
	// the notification.<110>
	RootList *RootList `idl:"name:ppRootList;pointer:unique" json:"root_list"`
}

func (o *AddFTRootRequest) xxx_ToOp(ctx context.Context, op *xxx_AddFTRootOperation) *xxx_AddFTRootOperation {
	if op == nil {
		op = &xxx_AddFTRootOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.DCName = o.DCName
	op.RootShare = o.RootShare
	op.FTDFSName = o.FTDFSName
	op.Comment = o.Comment
	op.ConfigDN = o.ConfigDN
	op.NewFTDFS = o.NewFTDFS
	op.Flags = o.Flags
	op.RootList = o.RootList
	return op
}

func (o *AddFTRootRequest) xxx_FromOp(ctx context.Context, op *xxx_AddFTRootOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.DCName = op.DCName
	o.RootShare = op.RootShare
	o.FTDFSName = op.FTDFSName
	o.Comment = op.Comment
	o.ConfigDN = op.ConfigDN
	o.NewFTDFS = op.NewFTDFS
	o.Flags = op.Flags
	o.RootList = op.RootList
}
func (o *AddFTRootRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AddFTRootRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddFTRootOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddFTRootResponse structure represents the NetrDfsAddFtRoot operation response
type AddFTRootResponse struct {
	// ppRootList: On success, returns a list of DFS root targets in the domain-based DFS
	// namespace that the client will be responsible for notifying of the change in the
	// DFS namespace. See section 3.2.4.3.1. The list MAY be empty if the server has performed
	// the notification.<110>
	RootList *RootList `idl:"name:ppRootList;pointer:unique" json:"root_list"`
	// Return: The NetrDfsAddFtRoot return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AddFTRootResponse) xxx_ToOp(ctx context.Context, op *xxx_AddFTRootOperation) *xxx_AddFTRootOperation {
	if op == nil {
		op = &xxx_AddFTRootOperation{}
	}
	if o == nil {
		return op
	}
	op.RootList = o.RootList
	op.Return = o.Return
	return op
}

func (o *AddFTRootResponse) xxx_FromOp(ctx context.Context, op *xxx_AddFTRootOperation) {
	if o == nil {
		return
	}
	o.RootList = op.RootList
	o.Return = op.Return
}
func (o *AddFTRootResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AddFTRootResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddFTRootOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveFTRootOperation structure represents the NetrDfsRemoveFtRoot operation
type xxx_RemoveFTRootOperation struct {
	ServerName string    `idl:"name:ServerName;string" json:"server_name"`
	DCName     string    `idl:"name:DcName;string" json:"dc_name"`
	RootShare  string    `idl:"name:RootShare;string" json:"root_share"`
	FTDFSName  string    `idl:"name:FtDfsName;string" json:"ft_dfs_name"`
	Flags      uint32    `idl:"name:ApiFlags" json:"flags"`
	RootList   *RootList `idl:"name:ppRootList;pointer:unique" json:"root_list"`
	Return     uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveFTRootOperation) OpNum() int { return 11 }

func (o *xxx_RemoveFTRootOperation) OpName() string { return "/netdfs/v3/NetrDfsRemoveFtRoot" }

func (o *xxx_RemoveFTRootOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveFTRootOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
			return err
		}
	}
	// DcName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DCName); err != nil {
			return err
		}
	}
	// RootShare {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.RootShare); err != nil {
			return err
		}
	}
	// FtDfsName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FTDFSName); err != nil {
			return err
		}
	}
	// ApiFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	// ppRootList {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=DFSM_ROOT_LIST}(struct))
	{
		if o.RootList != nil {
			_ptr_ppRootList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RootList != nil {
					_ptr_ppRootList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.RootList != nil {
							if err := o.RootList.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&RootList{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.RootList, _ptr_ppRootList); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RootList, _ptr_ppRootList); err != nil {
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

func (o *xxx_RemoveFTRootOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
			return err
		}
	}
	// DcName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DCName); err != nil {
			return err
		}
	}
	// RootShare {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.RootShare); err != nil {
			return err
		}
	}
	// FtDfsName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FTDFSName); err != nil {
			return err
		}
	}
	// ApiFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	// ppRootList {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=DFSM_ROOT_LIST}(struct))
	{
		_ptr_ppRootList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppRootList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.RootList == nil {
					o.RootList = &RootList{}
				}
				if err := o.RootList.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppRootList := func(ptr interface{}) { o.RootList = *ptr.(**RootList) }
			if err := w.ReadPointer(&o.RootList, _s_ppRootList, _ptr_ppRootList); err != nil {
				return err
			}
			return nil
		})
		_s_ppRootList := func(ptr interface{}) { o.RootList = *ptr.(**RootList) }
		if err := w.ReadPointer(&o.RootList, _s_ppRootList, _ptr_ppRootList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveFTRootOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveFTRootOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppRootList {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=DFSM_ROOT_LIST}(struct))
	{
		if o.RootList != nil {
			_ptr_ppRootList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RootList != nil {
					_ptr_ppRootList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.RootList != nil {
							if err := o.RootList.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&RootList{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.RootList, _ptr_ppRootList); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RootList, _ptr_ppRootList); err != nil {
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
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveFTRootOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppRootList {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=DFSM_ROOT_LIST}(struct))
	{
		_ptr_ppRootList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppRootList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.RootList == nil {
					o.RootList = &RootList{}
				}
				if err := o.RootList.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppRootList := func(ptr interface{}) { o.RootList = *ptr.(**RootList) }
			if err := w.ReadPointer(&o.RootList, _s_ppRootList, _ptr_ppRootList); err != nil {
				return err
			}
			return nil
		})
		_s_ppRootList := func(ptr interface{}) { o.RootList = *ptr.(**RootList) }
		if err := w.ReadPointer(&o.RootList, _s_ppRootList, _ptr_ppRootList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveFTRootRequest structure represents the NetrDfsRemoveFtRoot operation request
type RemoveFTRootRequest struct {
	// ServerName: The pointer to a null-terminated Unicode string. This is the host name
	// DFS root target to be removed.
	ServerName string `idl:"name:ServerName;string" json:"server_name"`
	// DcName: The pointer to a null-terminated Unicode string. For a domainv1-based DFS
	// namespace, this string contains the host name of the DC to be used by the DFS root
	// targets being removed for getting or updating DFS metadata for the DFS namespace.
	// This parameter MAY be a NULL pointer; otherwise, it MUST be the PDC for the domain
	// of the DFS namespace.
	DCName string `idl:"name:DcName;string" json:"dc_name"`
	// RootShare: The pointer to a null-terminated Unicode DFS root target share name string.
	// The share is not removed automatically when the method is successful; it MUST be
	// removed explicitly as needed.
	RootShare string `idl:"name:RootShare;string" json:"root_share"`
	// FtDfsName: The pointer to a null-terminated Unicode string that contains the DFS
	// namespace in which the operation is to be performed. It MAY be different from the
	// RootShare.
	FTDFSName string `idl:"name:FtDfsName;string" json:"ft_dfs_name"`
	// ApiFlags: The only supported bit in the ApiFlags parameter is DFS_FORCE_REMOVE.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|                             |                                                                                  |
	//	|            VALUE            |                                     MEANING                                      |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| DFS_FORCE_REMOVE 0x80000000 | Removes the named DFS root target from the namespace's directory service         |
	//	|                             | metadata only.                                                                   |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:ApiFlags" json:"flags"`
	// ppRootList: On success, returns a list of DFS root targets in the domain-based DFS
	// namespace which the client will be responsible for notifying about the change in
	// the DFS namespace. See section 3.2.4.3.2. The list MAY be empty if the server has
	// performed the notification.<114>
	RootList *RootList `idl:"name:ppRootList;pointer:unique" json:"root_list"`
}

func (o *RemoveFTRootRequest) xxx_ToOp(ctx context.Context, op *xxx_RemoveFTRootOperation) *xxx_RemoveFTRootOperation {
	if op == nil {
		op = &xxx_RemoveFTRootOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.DCName = o.DCName
	op.RootShare = o.RootShare
	op.FTDFSName = o.FTDFSName
	op.Flags = o.Flags
	op.RootList = o.RootList
	return op
}

func (o *RemoveFTRootRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveFTRootOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.DCName = op.DCName
	o.RootShare = op.RootShare
	o.FTDFSName = op.FTDFSName
	o.Flags = op.Flags
	o.RootList = op.RootList
}
func (o *RemoveFTRootRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemoveFTRootRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveFTRootOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveFTRootResponse structure represents the NetrDfsRemoveFtRoot operation response
type RemoveFTRootResponse struct {
	// ppRootList: On success, returns a list of DFS root targets in the domain-based DFS
	// namespace which the client will be responsible for notifying about the change in
	// the DFS namespace. See section 3.2.4.3.2. The list MAY be empty if the server has
	// performed the notification.<114>
	RootList *RootList `idl:"name:ppRootList;pointer:unique" json:"root_list"`
	// Return: The NetrDfsRemoveFtRoot return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RemoveFTRootResponse) xxx_ToOp(ctx context.Context, op *xxx_RemoveFTRootOperation) *xxx_RemoveFTRootOperation {
	if op == nil {
		op = &xxx_RemoveFTRootOperation{}
	}
	if o == nil {
		return op
	}
	op.RootList = o.RootList
	op.Return = o.Return
	return op
}

func (o *RemoveFTRootResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveFTRootOperation) {
	if o == nil {
		return
	}
	o.RootList = op.RootList
	o.Return = op.Return
}
func (o *RemoveFTRootResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemoveFTRootResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveFTRootOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddStdRootOperation structure represents the NetrDfsAddStdRoot operation
type xxx_AddStdRootOperation struct {
	ServerName string `idl:"name:ServerName;string" json:"server_name"`
	RootShare  string `idl:"name:RootShare;string" json:"root_share"`
	Comment    string `idl:"name:Comment;string" json:"comment"`
	Flags      uint32 `idl:"name:ApiFlags" json:"flags"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_AddStdRootOperation) OpNum() int { return 12 }

func (o *xxx_AddStdRootOperation) OpName() string { return "/netdfs/v3/NetrDfsAddStdRoot" }

func (o *xxx_AddStdRootOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddStdRootOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
			return err
		}
	}
	// RootShare {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.RootShare); err != nil {
			return err
		}
	}
	// Comment {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Comment); err != nil {
			return err
		}
	}
	// ApiFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddStdRootOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
			return err
		}
	}
	// RootShare {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.RootShare); err != nil {
			return err
		}
	}
	// Comment {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Comment); err != nil {
			return err
		}
	}
	// ApiFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddStdRootOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddStdRootOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddStdRootOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddStdRootRequest structure represents the NetrDfsAddStdRoot operation request
type AddStdRootRequest struct {
	// ServerName: The pointer to a null-terminated Unicode string. This is the host name
	// of the new DFS root target.
	ServerName string `idl:"name:ServerName;string" json:"server_name"`
	// RootShare: The pointer to a null-terminated Unicode string. This is the new DFS root
	// target share name as well as the DFS namespace name. The share MUST already exist.
	RootShare string `idl:"name:RootShare;string" json:"root_share"`
	// Comment: The pointer to a null-terminated Unicode string that contains a comment
	// associated with the DFS namespace. Used for informational purposes, this string has
	// no protocol-specified restrictions on length or content. The comment is meant for
	// human consumption and does not affect server functionality. This parameter MAY be
	// a NULL pointer.
	Comment string `idl:"name:Comment;string" json:"comment"`
	// ApiFlags: This parameter is reserved for future use and is ignored by the server.
	Flags uint32 `idl:"name:ApiFlags" json:"flags"`
}

func (o *AddStdRootRequest) xxx_ToOp(ctx context.Context, op *xxx_AddStdRootOperation) *xxx_AddStdRootOperation {
	if op == nil {
		op = &xxx_AddStdRootOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.RootShare = o.RootShare
	op.Comment = o.Comment
	op.Flags = o.Flags
	return op
}

func (o *AddStdRootRequest) xxx_FromOp(ctx context.Context, op *xxx_AddStdRootOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.RootShare = op.RootShare
	o.Comment = op.Comment
	o.Flags = op.Flags
}
func (o *AddStdRootRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AddStdRootRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddStdRootOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddStdRootResponse structure represents the NetrDfsAddStdRoot operation response
type AddStdRootResponse struct {
	// Return: The NetrDfsAddStdRoot return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AddStdRootResponse) xxx_ToOp(ctx context.Context, op *xxx_AddStdRootOperation) *xxx_AddStdRootOperation {
	if op == nil {
		op = &xxx_AddStdRootOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *AddStdRootResponse) xxx_FromOp(ctx context.Context, op *xxx_AddStdRootOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AddStdRootResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AddStdRootResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddStdRootOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveStdRootOperation structure represents the NetrDfsRemoveStdRoot operation
type xxx_RemoveStdRootOperation struct {
	ServerName string `idl:"name:ServerName;string" json:"server_name"`
	RootShare  string `idl:"name:RootShare;string" json:"root_share"`
	Flags      uint32 `idl:"name:ApiFlags" json:"flags"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveStdRootOperation) OpNum() int { return 13 }

func (o *xxx_RemoveStdRootOperation) OpName() string { return "/netdfs/v3/NetrDfsRemoveStdRoot" }

func (o *xxx_RemoveStdRootOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveStdRootOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
			return err
		}
	}
	// RootShare {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.RootShare); err != nil {
			return err
		}
	}
	// ApiFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveStdRootOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
			return err
		}
	}
	// RootShare {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.RootShare); err != nil {
			return err
		}
	}
	// ApiFlags {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveStdRootOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveStdRootOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveStdRootOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveStdRootRequest structure represents the NetrDfsRemoveStdRoot operation request
type RemoveStdRootRequest struct {
	// ServerName: The pointer to a null-terminated Unicode string. This is the host name
	// of the DFS root target to be removed.
	ServerName string `idl:"name:ServerName;string" json:"server_name"`
	// RootShare: The pointer to a null-terminated Unicode DFS root target share name string.
	// This is also the DFS namespace name. The share is not removed automatically when
	// the method is successful; it MUST be removed explicitly, as needed.
	RootShare string `idl:"name:RootShare;string" json:"root_share"`
	// ApiFlags: This parameter is reserved for future use and is ignored by the server.
	Flags uint32 `idl:"name:ApiFlags" json:"flags"`
}

func (o *RemoveStdRootRequest) xxx_ToOp(ctx context.Context, op *xxx_RemoveStdRootOperation) *xxx_RemoveStdRootOperation {
	if op == nil {
		op = &xxx_RemoveStdRootOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.RootShare = o.RootShare
	op.Flags = o.Flags
	return op
}

func (o *RemoveStdRootRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveStdRootOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.RootShare = op.RootShare
	o.Flags = op.Flags
}
func (o *RemoveStdRootRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemoveStdRootRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveStdRootOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveStdRootResponse structure represents the NetrDfsRemoveStdRoot operation response
type RemoveStdRootResponse struct {
	// Return: The NetrDfsRemoveStdRoot return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RemoveStdRootResponse) xxx_ToOp(ctx context.Context, op *xxx_RemoveStdRootOperation) *xxx_RemoveStdRootOperation {
	if op == nil {
		op = &xxx_RemoveStdRootOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RemoveStdRootResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveStdRootOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RemoveStdRootResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemoveStdRootResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveStdRootOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ManagerInitializeOperation structure represents the NetrDfsManagerInitialize operation
type xxx_ManagerInitializeOperation struct {
	ServerName string `idl:"name:ServerName;string" json:"server_name"`
	Flags      uint32 `idl:"name:Flags" json:"flags"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_ManagerInitializeOperation) OpNum() int { return 14 }

func (o *xxx_ManagerInitializeOperation) OpName() string {
	return "/netdfs/v3/NetrDfsManagerInitialize"
}

func (o *xxx_ManagerInitializeOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ManagerInitializeOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
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

func (o *xxx_ManagerInitializeOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
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

func (o *xxx_ManagerInitializeOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ManagerInitializeOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ManagerInitializeOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ManagerInitializeRequest structure represents the NetrDfsManagerInitialize operation request
type ManagerInitializeRequest struct {
	// ServerName: The pointer to a null-terminated Unicode host name string of the DFS
	// root target server or DC where the DFS service is to be reinitialized.
	ServerName string `idl:"name:ServerName;string" json:"server_name"`
	// Flags: This parameter MUST be zero.
	Flags uint32 `idl:"name:Flags" json:"flags"`
}

func (o *ManagerInitializeRequest) xxx_ToOp(ctx context.Context, op *xxx_ManagerInitializeOperation) *xxx_ManagerInitializeOperation {
	if op == nil {
		op = &xxx_ManagerInitializeOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.Flags = o.Flags
	return op
}

func (o *ManagerInitializeRequest) xxx_FromOp(ctx context.Context, op *xxx_ManagerInitializeOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.Flags = op.Flags
}
func (o *ManagerInitializeRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *ManagerInitializeRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ManagerInitializeOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ManagerInitializeResponse structure represents the NetrDfsManagerInitialize operation response
type ManagerInitializeResponse struct {
	// Return: The NetrDfsManagerInitialize return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *ManagerInitializeResponse) xxx_ToOp(ctx context.Context, op *xxx_ManagerInitializeOperation) *xxx_ManagerInitializeOperation {
	if op == nil {
		op = &xxx_ManagerInitializeOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *ManagerInitializeResponse) xxx_FromOp(ctx context.Context, op *xxx_ManagerInitializeOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *ManagerInitializeResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *ManagerInitializeResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ManagerInitializeOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddStdRootForcedOperation structure represents the NetrDfsAddStdRootForced operation
type xxx_AddStdRootForcedOperation struct {
	ServerName string `idl:"name:ServerName;string" json:"server_name"`
	RootShare  string `idl:"name:RootShare;string" json:"root_share"`
	Comment    string `idl:"name:Comment;string" json:"comment"`
	Share      string `idl:"name:Share;string" json:"share"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_AddStdRootForcedOperation) OpNum() int { return 15 }

func (o *xxx_AddStdRootForcedOperation) OpName() string { return "/netdfs/v3/NetrDfsAddStdRootForced" }

func (o *xxx_AddStdRootForcedOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddStdRootForcedOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
			return err
		}
	}
	// RootShare {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.RootShare); err != nil {
			return err
		}
	}
	// Comment {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Comment); err != nil {
			return err
		}
	}
	// Share {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.Share); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddStdRootForcedOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
			return err
		}
	}
	// RootShare {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.RootShare); err != nil {
			return err
		}
	}
	// Comment {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Comment); err != nil {
			return err
		}
	}
	// Share {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.Share); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddStdRootForcedOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddStdRootForcedOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddStdRootForcedOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddStdRootForcedRequest structure represents the NetrDfsAddStdRootForced operation request
type AddStdRootForcedRequest struct {
	// ServerName: The pointer to a null-terminated Unicode string. This is the host name
	// of the new DFS root target.
	ServerName string `idl:"name:ServerName;string" json:"server_name"`
	// RootShare: The pointer to a null-terminated Unicode DFS root target share name string.
	// This is also the DFS namespace name. This method does not create the share; it MUST
	// be created separately.
	RootShare string `idl:"name:RootShare;string" json:"root_share"`
	// Comment: The pointer to a null-terminated Unicode string that contains a comment
	// associated with the DFS namespace. Used for informational purposes, this string has
	// no protocol-specified restrictions on length or content. The comment is meant for
	// human consumption and does not affect server functionality. This parameter MAY be
	// a NULL pointer.
	Comment string `idl:"name:Comment;string" json:"comment"`
	// Share:  The pointer to a null-terminated Unicode string that contains the local
	// file system path corresponding to the share on the server receiving the RPC method,
	// in the following form:
	//
	// <X>:\<path>
	Share string `idl:"name:Share;string" json:"share"`
}

func (o *AddStdRootForcedRequest) xxx_ToOp(ctx context.Context, op *xxx_AddStdRootForcedOperation) *xxx_AddStdRootForcedOperation {
	if op == nil {
		op = &xxx_AddStdRootForcedOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.RootShare = o.RootShare
	op.Comment = o.Comment
	op.Share = o.Share
	return op
}

func (o *AddStdRootForcedRequest) xxx_FromOp(ctx context.Context, op *xxx_AddStdRootForcedOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.RootShare = op.RootShare
	o.Comment = op.Comment
	o.Share = op.Share
}
func (o *AddStdRootForcedRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AddStdRootForcedRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddStdRootForcedOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddStdRootForcedResponse structure represents the NetrDfsAddStdRootForced operation response
type AddStdRootForcedResponse struct {
	// Return: The NetrDfsAddStdRootForced return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AddStdRootForcedResponse) xxx_ToOp(ctx context.Context, op *xxx_AddStdRootForcedOperation) *xxx_AddStdRootForcedOperation {
	if op == nil {
		op = &xxx_AddStdRootForcedOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *AddStdRootForcedResponse) xxx_FromOp(ctx context.Context, op *xxx_AddStdRootForcedOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AddStdRootForcedResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AddStdRootForcedResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddStdRootForcedOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetDCAddressOperation structure represents the NetrDfsGetDcAddress operation
type xxx_GetDCAddressOperation struct {
	ServerName string `idl:"name:ServerName;string" json:"server_name"`
	DCName     string `idl:"name:DcName;string" json:"dc_name"`
	IsRoot     bool   `idl:"name:IsRoot" json:"is_root"`
	Timeout    uint32 `idl:"name:Timeout" json:"timeout"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_GetDCAddressOperation) OpNum() int { return 16 }

func (o *xxx_GetDCAddressOperation) OpName() string { return "/netdfs/v3/NetrDfsGetDcAddress" }

func (o *xxx_GetDCAddressOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDCAddressOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
			return err
		}
	}
	// DcName {in, out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.DCName != "" {
			_ptr_DcName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.DCName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DCName, _ptr_DcName); err != nil {
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
	// IsRoot {in, out} (1:{pointer=ref}*(1))(2:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.IsRoot); err != nil {
			return err
		}
	}
	// Timeout {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDCAddressOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
			return err
		}
	}
	// DcName {in, out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_DcName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.DCName); err != nil {
				return err
			}
			return nil
		})
		_s_DcName := func(ptr interface{}) { o.DCName = *ptr.(*string) }
		if err := w.ReadPointer(&o.DCName, _s_DcName, _ptr_DcName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// IsRoot {in, out} (1:{pointer=ref}*(1))(2:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.IsRoot); err != nil {
			return err
		}
	}
	// Timeout {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDCAddressOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDCAddressOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// DcName {in, out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.DCName != "" {
			_ptr_DcName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.DCName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DCName, _ptr_DcName); err != nil {
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
	// IsRoot {in, out} (1:{pointer=ref}*(1))(2:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.IsRoot); err != nil {
			return err
		}
	}
	// Timeout {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetDCAddressOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// DcName {in, out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_DcName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.DCName); err != nil {
				return err
			}
			return nil
		})
		_s_DcName := func(ptr interface{}) { o.DCName = *ptr.(*string) }
		if err := w.ReadPointer(&o.DCName, _s_DcName, _ptr_DcName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// IsRoot {in, out} (1:{pointer=ref}*(1))(2:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.IsRoot); err != nil {
			return err
		}
	}
	// Timeout {in, out} (1:{pointer=ref}*(1)(uint32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetDCAddressRequest structure represents the NetrDfsGetDcAddress operation request
type GetDCAddressRequest struct {
	// ServerName: A pointer to a null-terminated Unicode string. This is the host name
	// of the server to which the RPC method is issued.<127>
	ServerName string `idl:"name:ServerName;string" json:"server_name"`
	// DcName: A null-terminated Unicode string that contains the DC host name when the
	// NetrDfsGetDcAddress method is successful.<128>
	DCName string `idl:"name:DcName;string" json:"dc_name"`
	// IsRoot: A pointer to a Boolean value, set to TRUE on return if the server hosts any
	// DFS root target, and FALSE otherwise.<129>
	IsRoot bool `idl:"name:IsRoot" json:"is_root"`
	// Timeout: A pointer to an unsigned 32-bit integer value indicating the count of seconds
	// for which the server will use the DC that is returned to access DFS metadata.<130>
	Timeout uint32 `idl:"name:Timeout" json:"timeout"`
}

func (o *GetDCAddressRequest) xxx_ToOp(ctx context.Context, op *xxx_GetDCAddressOperation) *xxx_GetDCAddressOperation {
	if op == nil {
		op = &xxx_GetDCAddressOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.DCName = o.DCName
	op.IsRoot = o.IsRoot
	op.Timeout = o.Timeout
	return op
}

func (o *GetDCAddressRequest) xxx_FromOp(ctx context.Context, op *xxx_GetDCAddressOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.DCName = op.DCName
	o.IsRoot = op.IsRoot
	o.Timeout = op.Timeout
}
func (o *GetDCAddressRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetDCAddressRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDCAddressOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetDCAddressResponse structure represents the NetrDfsGetDcAddress operation response
type GetDCAddressResponse struct {
	// DcName: A null-terminated Unicode string that contains the DC host name when the
	// NetrDfsGetDcAddress method is successful.<128>
	DCName string `idl:"name:DcName;string" json:"dc_name"`
	// IsRoot: A pointer to a Boolean value, set to TRUE on return if the server hosts any
	// DFS root target, and FALSE otherwise.<129>
	IsRoot bool `idl:"name:IsRoot" json:"is_root"`
	// Timeout: A pointer to an unsigned 32-bit integer value indicating the count of seconds
	// for which the server will use the DC that is returned to access DFS metadata.<130>
	Timeout uint32 `idl:"name:Timeout" json:"timeout"`
	// Return: The NetrDfsGetDcAddress return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetDCAddressResponse) xxx_ToOp(ctx context.Context, op *xxx_GetDCAddressOperation) *xxx_GetDCAddressOperation {
	if op == nil {
		op = &xxx_GetDCAddressOperation{}
	}
	if o == nil {
		return op
	}
	op.DCName = o.DCName
	op.IsRoot = o.IsRoot
	op.Timeout = o.Timeout
	op.Return = o.Return
	return op
}

func (o *GetDCAddressResponse) xxx_FromOp(ctx context.Context, op *xxx_GetDCAddressOperation) {
	if o == nil {
		return
	}
	o.DCName = op.DCName
	o.IsRoot = op.IsRoot
	o.Timeout = op.Timeout
	o.Return = op.Return
}
func (o *GetDCAddressResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetDCAddressResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetDCAddressOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetDCAddressOperation structure represents the NetrDfsSetDcAddress operation
type xxx_SetDCAddressOperation struct {
	ServerName string `idl:"name:ServerName;string" json:"server_name"`
	DCName     string `idl:"name:DcName;string" json:"dc_name"`
	Timeout    uint32 `idl:"name:Timeout" json:"timeout"`
	Flags      uint32 `idl:"name:Flags" json:"flags"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_SetDCAddressOperation) OpNum() int { return 17 }

func (o *xxx_SetDCAddressOperation) OpName() string { return "/netdfs/v3/NetrDfsSetDcAddress" }

func (o *xxx_SetDCAddressOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDCAddressOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// ServerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
			return err
		}
	}
	// DcName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DCName); err != nil {
			return err
		}
	}
	// Timeout {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Timeout); err != nil {
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

func (o *xxx_SetDCAddressOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// ServerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
			return err
		}
	}
	// DcName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DCName); err != nil {
			return err
		}
	}
	// Timeout {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Timeout); err != nil {
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

func (o *xxx_SetDCAddressOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDCAddressOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetDCAddressOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetDCAddressRequest structure represents the NetrDfsSetDcAddress operation request
type SetDCAddressRequest struct {
	// ServerName: The pointer to a null-terminated Unicode string. This is the host name
	// of the server to which the RPC method is issued.
	ServerName string `idl:"name:ServerName;string" json:"server_name"`
	// DcName: The pointer to a null-terminated Unicode DC host name string.
	DCName string `idl:"name:DcName;string" json:"dc_name"`
	// Timeout: The time period, in seconds, that the server uses the specified DC when
	// storing and retrieving domain-based DFS metadata. This is valid only when the NET_DFS_SETDC_TIMEOUT
	// bit of the Flags parameter is set.
	Timeout uint32 `idl:"name:Timeout" json:"timeout"`
	// Flags: The bit field specifying additional operations to perform. Possible values
	// are as follows.
	//
	//	+-----------------------------------+-------------------------------------------------------------------------------+
	//	|                                   |                                                                               |
	//	|               VALUE               |                                    MEANING                                    |
	//	|                                   |                                                                               |
	//	+-----------------------------------+-------------------------------------------------------------------------------+
	//	+-----------------------------------+-------------------------------------------------------------------------------+
	//	| NET_DFS_SETDC_FLAGS 0x00000000    | Indicates that no additional operation is requested.                          |
	//	+-----------------------------------+-------------------------------------------------------------------------------+
	//	| NET_DFS_SETDC_TIMEOUT 0x00000001  | Sets the time-out value based on the Timeout parameter.                       |
	//	+-----------------------------------+-------------------------------------------------------------------------------+
	//	| NET_DFS_SETDC_INIT_PKT 0x00000002 | Instructs the called server to reload its DFS metadata from the specified DC. |
	//	+-----------------------------------+-------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
}

func (o *SetDCAddressRequest) xxx_ToOp(ctx context.Context, op *xxx_SetDCAddressOperation) *xxx_SetDCAddressOperation {
	if op == nil {
		op = &xxx_SetDCAddressOperation{}
	}
	if o == nil {
		return op
	}
	op.ServerName = o.ServerName
	op.DCName = o.DCName
	op.Timeout = o.Timeout
	op.Flags = o.Flags
	return op
}

func (o *SetDCAddressRequest) xxx_FromOp(ctx context.Context, op *xxx_SetDCAddressOperation) {
	if o == nil {
		return
	}
	o.ServerName = op.ServerName
	o.DCName = op.DCName
	o.Timeout = op.Timeout
	o.Flags = op.Flags
}
func (o *SetDCAddressRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetDCAddressRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDCAddressOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetDCAddressResponse structure represents the NetrDfsSetDcAddress operation response
type SetDCAddressResponse struct {
	// Return: The NetrDfsSetDcAddress return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetDCAddressResponse) xxx_ToOp(ctx context.Context, op *xxx_SetDCAddressOperation) *xxx_SetDCAddressOperation {
	if op == nil {
		op = &xxx_SetDCAddressOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *SetDCAddressResponse) xxx_FromOp(ctx context.Context, op *xxx_SetDCAddressOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *SetDCAddressResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetDCAddressResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetDCAddressOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FlushFTTableOperation structure represents the NetrDfsFlushFtTable operation
type xxx_FlushFTTableOperation struct {
	DCName    string `idl:"name:DcName;string" json:"dc_name"`
	FTDFSName string `idl:"name:wszFtDfsName;string" json:"ft_dfs_name"`
	Return    uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_FlushFTTableOperation) OpNum() int { return 18 }

func (o *xxx_FlushFTTableOperation) OpName() string { return "/netdfs/v3/NetrDfsFlushFtTable" }

func (o *xxx_FlushFTTableOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FlushFTTableOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DcName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DCName); err != nil {
			return err
		}
	}
	// wszFtDfsName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.FTDFSName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FlushFTTableOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DcName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DCName); err != nil {
			return err
		}
	}
	// wszFtDfsName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.FTDFSName); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FlushFTTableOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FlushFTTableOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FlushFTTableOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FlushFTTableRequest structure represents the NetrDfsFlushFtTable operation request
type FlushFTTableRequest struct {
	// DcName: The pointer to a null-terminated Unicode string that contains the host name
	// of the DC to which the RPC method is issued.
	DCName string `idl:"name:DcName;string" json:"dc_name"`
	// wszFtDfsName: The pointer to a null-terminated Unicode string that contains the name
	// of the domain-based DFS namespace.
	FTDFSName string `idl:"name:wszFtDfsName;string" json:"ft_dfs_name"`
}

func (o *FlushFTTableRequest) xxx_ToOp(ctx context.Context, op *xxx_FlushFTTableOperation) *xxx_FlushFTTableOperation {
	if op == nil {
		op = &xxx_FlushFTTableOperation{}
	}
	if o == nil {
		return op
	}
	op.DCName = o.DCName
	op.FTDFSName = o.FTDFSName
	return op
}

func (o *FlushFTTableRequest) xxx_FromOp(ctx context.Context, op *xxx_FlushFTTableOperation) {
	if o == nil {
		return
	}
	o.DCName = op.DCName
	o.FTDFSName = op.FTDFSName
}
func (o *FlushFTTableRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *FlushFTTableRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FlushFTTableOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FlushFTTableResponse structure represents the NetrDfsFlushFtTable operation response
type FlushFTTableResponse struct {
	// Return: The NetrDfsFlushFtTable return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *FlushFTTableResponse) xxx_ToOp(ctx context.Context, op *xxx_FlushFTTableOperation) *xxx_FlushFTTableOperation {
	if op == nil {
		op = &xxx_FlushFTTableOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *FlushFTTableResponse) xxx_FromOp(ctx context.Context, op *xxx_FlushFTTableOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *FlushFTTableResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *FlushFTTableResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FlushFTTableOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_Add2Operation structure represents the NetrDfsAdd2 operation
type xxx_Add2Operation struct {
	EntryPath  string    `idl:"name:DfsEntryPath;string" json:"entry_path"`
	DCName     string    `idl:"name:DcName;string" json:"dc_name"`
	ServerName string    `idl:"name:ServerName;string" json:"server_name"`
	ShareName  string    `idl:"name:ShareName;string;pointer:unique" json:"share_name"`
	Comment    string    `idl:"name:Comment;string;pointer:unique" json:"comment"`
	Flags      uint32    `idl:"name:Flags" json:"flags"`
	RootList   *RootList `idl:"name:ppRootList;pointer:unique" json:"root_list"`
	Return     uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_Add2Operation) OpNum() int { return 19 }

func (o *xxx_Add2Operation) OpName() string { return "/netdfs/v3/NetrDfsAdd2" }

func (o *xxx_Add2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Add2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DfsEntryPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.EntryPath); err != nil {
			return err
		}
	}
	// DcName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DCName); err != nil {
			return err
		}
	}
	// ServerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
			return err
		}
	}
	// ShareName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.ShareName != "" {
			_ptr_ShareName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ShareName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ShareName, _ptr_ShareName); err != nil {
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
	// Comment {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.Comment != "" {
			_ptr_Comment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Comment); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Comment, _ptr_Comment); err != nil {
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
	// ppRootList {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=DFSM_ROOT_LIST}(struct))
	{
		if o.RootList != nil {
			_ptr_ppRootList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RootList != nil {
					_ptr_ppRootList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.RootList != nil {
							if err := o.RootList.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&RootList{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.RootList, _ptr_ppRootList); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RootList, _ptr_ppRootList); err != nil {
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

func (o *xxx_Add2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DfsEntryPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.EntryPath); err != nil {
			return err
		}
	}
	// DcName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DCName); err != nil {
			return err
		}
	}
	// ServerName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
			return err
		}
	}
	// ShareName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ShareName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ShareName); err != nil {
				return err
			}
			return nil
		})
		_s_ShareName := func(ptr interface{}) { o.ShareName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ShareName, _s_ShareName, _ptr_ShareName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Comment {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_Comment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Comment); err != nil {
				return err
			}
			return nil
		})
		_s_Comment := func(ptr interface{}) { o.Comment = *ptr.(*string) }
		if err := w.ReadPointer(&o.Comment, _s_Comment, _ptr_Comment); err != nil {
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
	// ppRootList {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=DFSM_ROOT_LIST}(struct))
	{
		_ptr_ppRootList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppRootList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.RootList == nil {
					o.RootList = &RootList{}
				}
				if err := o.RootList.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppRootList := func(ptr interface{}) { o.RootList = *ptr.(**RootList) }
			if err := w.ReadPointer(&o.RootList, _s_ppRootList, _ptr_ppRootList); err != nil {
				return err
			}
			return nil
		})
		_s_ppRootList := func(ptr interface{}) { o.RootList = *ptr.(**RootList) }
		if err := w.ReadPointer(&o.RootList, _s_ppRootList, _ptr_ppRootList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Add2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Add2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppRootList {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=DFSM_ROOT_LIST}(struct))
	{
		if o.RootList != nil {
			_ptr_ppRootList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RootList != nil {
					_ptr_ppRootList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.RootList != nil {
							if err := o.RootList.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&RootList{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.RootList, _ptr_ppRootList); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RootList, _ptr_ppRootList); err != nil {
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
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Add2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppRootList {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=DFSM_ROOT_LIST}(struct))
	{
		_ptr_ppRootList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppRootList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.RootList == nil {
					o.RootList = &RootList{}
				}
				if err := o.RootList.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppRootList := func(ptr interface{}) { o.RootList = *ptr.(**RootList) }
			if err := w.ReadPointer(&o.RootList, _s_ppRootList, _ptr_ppRootList); err != nil {
				return err
			}
			return nil
		})
		_s_ppRootList := func(ptr interface{}) { o.RootList = *ptr.(**RootList) }
		if err := w.ReadPointer(&o.RootList, _s_ppRootList, _ptr_ppRootList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// Add2Request structure represents the NetrDfsAdd2 operation request
type Add2Request struct {
	// DfsEntryPath: A pointer to a DFS link path that contains the name of an existing
	// link when additional link targets are added or the name of a new link is created.
	EntryPath string `idl:"name:DfsEntryPath;string" json:"entry_path"`
	// DcName: A pointer to a null-terminated Unicode string. For a domain-based DFS namespace,
	// this is the host name of the DC that the DFS root target uses to get or update DFS
	// metadata for the DFS namespace. This parameter MAY be a NULL pointer; otherwise,
	// it MUST be the host name of the PDC for the domain of the DFS namespace.<79>
	DCName string `idl:"name:DcName;string" json:"dc_name"`
	// ServerName: A pointer to a null-terminated Unicode string that specifies the DFS
	// link target host name.
	ServerName string `idl:"name:ServerName;string" json:"server_name"`
	// ShareName: A pointer to a null-terminated Unicode DFS link target share name string.
	// This can also be a share name with a path relative to the share (for example, share1\mydir1\mydir2).
	// When specified in this manner, each pathname component MUST be a directory.
	ShareName string `idl:"name:ShareName;string;pointer:unique" json:"share_name"`
	// Comment: A pointer to a null-terminated, human-readable Unicode string description
	// associated with this root or link. This string is not subject to protocol-specified
	// restrictions on length or content and does not affect server functionality. The description
	// MUST be ignored when adding a target to an existing link.
	Comment string `idl:"name:Comment;string;pointer:unique" json:"comment"`
	// Flags: The flag that indicates the operation to perform. The following table lists
	// the possible values.
	//
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	|                               |                                                                                  |
	//	|             VALUE             |                                     MEANING                                      |
	//	|                               |                                                                                  |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000                    | Create a new link or add a new target to an existing link.                       |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_ADD_VOLUME 0x00000001     | Create a new link in the DFS namespace if one does not already exist or fail if  |
	//	|                               | it exists.                                                                       |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_RESTORE_VOLUME 0x00000002 | Add a target without verifying its existence.                                    |
	//	+-------------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
	// ppRootList: On success, returns a list of DFS root targets in the domain-based DFS
	// namespace that the client will be responsible for notifying of the change in the
	// DFS namespace. See section 3.2.4.2.1. This list MAY be empty if the server performs
	// the notification.<80>
	RootList *RootList `idl:"name:ppRootList;pointer:unique" json:"root_list"`
}

func (o *Add2Request) xxx_ToOp(ctx context.Context, op *xxx_Add2Operation) *xxx_Add2Operation {
	if op == nil {
		op = &xxx_Add2Operation{}
	}
	if o == nil {
		return op
	}
	op.EntryPath = o.EntryPath
	op.DCName = o.DCName
	op.ServerName = o.ServerName
	op.ShareName = o.ShareName
	op.Comment = o.Comment
	op.Flags = o.Flags
	op.RootList = o.RootList
	return op
}

func (o *Add2Request) xxx_FromOp(ctx context.Context, op *xxx_Add2Operation) {
	if o == nil {
		return
	}
	o.EntryPath = op.EntryPath
	o.DCName = op.DCName
	o.ServerName = op.ServerName
	o.ShareName = op.ShareName
	o.Comment = op.Comment
	o.Flags = op.Flags
	o.RootList = op.RootList
}
func (o *Add2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *Add2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Add2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// Add2Response structure represents the NetrDfsAdd2 operation response
type Add2Response struct {
	// ppRootList: On success, returns a list of DFS root targets in the domain-based DFS
	// namespace that the client will be responsible for notifying of the change in the
	// DFS namespace. See section 3.2.4.2.1. This list MAY be empty if the server performs
	// the notification.<80>
	RootList *RootList `idl:"name:ppRootList;pointer:unique" json:"root_list"`
	// Return: The NetrDfsAdd2 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *Add2Response) xxx_ToOp(ctx context.Context, op *xxx_Add2Operation) *xxx_Add2Operation {
	if op == nil {
		op = &xxx_Add2Operation{}
	}
	if o == nil {
		return op
	}
	op.RootList = o.RootList
	op.Return = o.Return
	return op
}

func (o *Add2Response) xxx_FromOp(ctx context.Context, op *xxx_Add2Operation) {
	if o == nil {
		return
	}
	o.RootList = op.RootList
	o.Return = op.Return
}
func (o *Add2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *Add2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Add2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_Remove2Operation structure represents the NetrDfsRemove2 operation
type xxx_Remove2Operation struct {
	EntryPath  string    `idl:"name:DfsEntryPath;string" json:"entry_path"`
	DCName     string    `idl:"name:DcName;string" json:"dc_name"`
	ServerName string    `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	ShareName  string    `idl:"name:ShareName;string;pointer:unique" json:"share_name"`
	RootList   *RootList `idl:"name:ppRootList;pointer:unique" json:"root_list"`
	Return     uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_Remove2Operation) OpNum() int { return 20 }

func (o *xxx_Remove2Operation) OpName() string { return "/netdfs/v3/NetrDfsRemove2" }

func (o *xxx_Remove2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Remove2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DfsEntryPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.EntryPath); err != nil {
			return err
		}
	}
	// DcName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DCName); err != nil {
			return err
		}
	}
	// ServerName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
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
	// ShareName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.ShareName != "" {
			_ptr_ShareName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ShareName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ShareName, _ptr_ShareName); err != nil {
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
	// ppRootList {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=DFSM_ROOT_LIST}(struct))
	{
		if o.RootList != nil {
			_ptr_ppRootList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RootList != nil {
					_ptr_ppRootList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.RootList != nil {
							if err := o.RootList.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&RootList{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.RootList, _ptr_ppRootList); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RootList, _ptr_ppRootList); err != nil {
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

func (o *xxx_Remove2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DfsEntryPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.EntryPath); err != nil {
			return err
		}
	}
	// DcName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DCName); err != nil {
			return err
		}
	}
	// ServerName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
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
	// ShareName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ShareName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ShareName); err != nil {
				return err
			}
			return nil
		})
		_s_ShareName := func(ptr interface{}) { o.ShareName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ShareName, _s_ShareName, _ptr_ShareName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppRootList {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=DFSM_ROOT_LIST}(struct))
	{
		_ptr_ppRootList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppRootList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.RootList == nil {
					o.RootList = &RootList{}
				}
				if err := o.RootList.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppRootList := func(ptr interface{}) { o.RootList = *ptr.(**RootList) }
			if err := w.ReadPointer(&o.RootList, _s_ppRootList, _ptr_ppRootList); err != nil {
				return err
			}
			return nil
		})
		_s_ppRootList := func(ptr interface{}) { o.RootList = *ptr.(**RootList) }
		if err := w.ReadPointer(&o.RootList, _s_ppRootList, _ptr_ppRootList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Remove2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Remove2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppRootList {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=DFSM_ROOT_LIST}(struct))
	{
		if o.RootList != nil {
			_ptr_ppRootList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RootList != nil {
					_ptr_ppRootList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.RootList != nil {
							if err := o.RootList.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&RootList{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.RootList, _ptr_ppRootList); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RootList, _ptr_ppRootList); err != nil {
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
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_Remove2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppRootList {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=DFSM_ROOT_LIST}(struct))
	{
		_ptr_ppRootList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppRootList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.RootList == nil {
					o.RootList = &RootList{}
				}
				if err := o.RootList.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppRootList := func(ptr interface{}) { o.RootList = *ptr.(**RootList) }
			if err := w.ReadPointer(&o.RootList, _s_ppRootList, _ptr_ppRootList); err != nil {
				return err
			}
			return nil
		})
		_s_ppRootList := func(ptr interface{}) { o.RootList = *ptr.(**RootList) }
		if err := w.ReadPointer(&o.RootList, _s_ppRootList, _ptr_ppRootList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// Remove2Request structure represents the NetrDfsRemove2 operation request
type Remove2Request struct {
	// DfsEntryPath: The pointer to a DFS link path that contains the name of the DFS link
	// to remove.
	EntryPath string `idl:"name:DfsEntryPath;string" json:"entry_path"`
	// DcName: The pointer to a null-terminated Unicode string. For a domain-based DFS namespace,
	// this string contains the host name of the DC to be used by the DFS root target that
	// is removing the DFS link. This parameter MAY be a NULL pointer; otherwise, it MUST
	// be the PDC for the domain of the DFS namespace.<87>
	DCName string `idl:"name:DcName;string" json:"dc_name"`
	// ServerName: The pointer to a null-terminated Unicode DFS link target host name string.
	// This MUST be a NULL pointer when the link and all of the link targets are to be removed.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// ShareName: The pointer to a null-terminated Unicode DFS link target share name string.
	// This MUST be a NULL pointer when the link and all of the link targets are to be removed.
	ShareName string `idl:"name:ShareName;string;pointer:unique" json:"share_name"`
	// ppRootList: On success, returns a list of DFS root targets in the domain-based DFS
	// namespace that the client will be responsible for notifying of the change in the
	// DFS namespace. See section 3.2.4.2.2. This list MAY be empty if the server has performed
	// the notification.<88>
	RootList *RootList `idl:"name:ppRootList;pointer:unique" json:"root_list"`
}

func (o *Remove2Request) xxx_ToOp(ctx context.Context, op *xxx_Remove2Operation) *xxx_Remove2Operation {
	if op == nil {
		op = &xxx_Remove2Operation{}
	}
	if o == nil {
		return op
	}
	op.EntryPath = o.EntryPath
	op.DCName = o.DCName
	op.ServerName = o.ServerName
	op.ShareName = o.ShareName
	op.RootList = o.RootList
	return op
}

func (o *Remove2Request) xxx_FromOp(ctx context.Context, op *xxx_Remove2Operation) {
	if o == nil {
		return
	}
	o.EntryPath = op.EntryPath
	o.DCName = op.DCName
	o.ServerName = op.ServerName
	o.ShareName = op.ShareName
	o.RootList = op.RootList
}
func (o *Remove2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *Remove2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Remove2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// Remove2Response structure represents the NetrDfsRemove2 operation response
type Remove2Response struct {
	// ppRootList: On success, returns a list of DFS root targets in the domain-based DFS
	// namespace that the client will be responsible for notifying of the change in the
	// DFS namespace. See section 3.2.4.2.2. This list MAY be empty if the server has performed
	// the notification.<88>
	RootList *RootList `idl:"name:ppRootList;pointer:unique" json:"root_list"`
	// Return: The NetrDfsRemove2 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *Remove2Response) xxx_ToOp(ctx context.Context, op *xxx_Remove2Operation) *xxx_Remove2Operation {
	if op == nil {
		op = &xxx_Remove2Operation{}
	}
	if o == nil {
		return op
	}
	op.RootList = o.RootList
	op.Return = o.Return
	return op
}

func (o *Remove2Response) xxx_FromOp(ctx context.Context, op *xxx_Remove2Operation) {
	if o == nil {
		return
	}
	o.RootList = op.RootList
	o.Return = op.Return
}
func (o *Remove2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *Remove2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_Remove2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_EnumExOperation structure represents the NetrDfsEnumEx operation
type xxx_EnumExOperation struct {
	EntryPath     string    `idl:"name:DfsEntryPath;string" json:"entry_path"`
	Level         uint32    `idl:"name:Level" json:"level"`
	PrefMaxLength uint32    `idl:"name:PrefMaxLen" json:"pref_max_length"`
	Enum          *InfoEnum `idl:"name:DfsEnum;pointer:unique" json:"enum"`
	Resume        uint32    `idl:"name:ResumeHandle;pointer:unique" json:"resume"`
	Return        uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_EnumExOperation) OpNum() int { return 21 }

func (o *xxx_EnumExOperation) OpName() string { return "/netdfs/v3/NetrDfsEnumEx" }

func (o *xxx_EnumExOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumExOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DfsEntryPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.EntryPath); err != nil {
			return err
		}
	}
	// Level {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.Level); err != nil {
			return err
		}
	}
	// PrefMaxLen {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.PrefMaxLength); err != nil {
			return err
		}
	}
	// DfsEnum {in, out} (1:{pointer=unique}*(1))(2:{alias=DFS_INFO_ENUM_STRUCT}(struct))
	{
		if o.Enum != nil {
			_ptr_DfsEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&InfoEnum{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_DfsEnum); err != nil {
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
	// ResumeHandle {in, out} (1:{pointer=unique}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_ResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumExOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DfsEntryPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.EntryPath); err != nil {
			return err
		}
	}
	// Level {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Level); err != nil {
			return err
		}
	}
	// PrefMaxLen {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.PrefMaxLength); err != nil {
			return err
		}
	}
	// DfsEnum {in, out} (1:{pointer=unique}*(1))(2:{alias=DFS_INFO_ENUM_STRUCT}(struct))
	{
		_ptr_DfsEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &InfoEnum{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_DfsEnum := func(ptr interface{}) { o.Enum = *ptr.(**InfoEnum) }
		if err := w.ReadPointer(&o.Enum, _s_DfsEnum, _ptr_DfsEnum); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=unique}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_ResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_ResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_ResumeHandle, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumExOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumExOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// DfsEnum {in, out} (1:{pointer=unique}*(1))(2:{alias=DFS_INFO_ENUM_STRUCT}(struct))
	{
		if o.Enum != nil {
			_ptr_DfsEnum := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.Enum != nil {
					if err := o.Enum.MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&InfoEnum{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.Enum, _ptr_DfsEnum); err != nil {
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
	// ResumeHandle {in, out} (1:{pointer=unique}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_ResumeHandle := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			if err := w.WriteData(o.Resume); err != nil {
				return err
			}
			return nil
		})
		if err := w.WritePointer(&o.Resume, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_EnumExOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// DfsEnum {in, out} (1:{pointer=unique}*(1))(2:{alias=DFS_INFO_ENUM_STRUCT}(struct))
	{
		_ptr_DfsEnum := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if o.Enum == nil {
				o.Enum = &InfoEnum{}
			}
			if err := o.Enum.UnmarshalNDR(ctx, w); err != nil {
				return err
			}
			return nil
		})
		_s_DfsEnum := func(ptr interface{}) { o.Enum = *ptr.(**InfoEnum) }
		if err := w.ReadPointer(&o.Enum, _s_DfsEnum, _ptr_DfsEnum); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ResumeHandle {in, out} (1:{pointer=unique}*(1))(2:{alias=DWORD}(uint32))
	{
		_ptr_ResumeHandle := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := w.ReadData(&o.Resume); err != nil {
				return err
			}
			return nil
		})
		_s_ResumeHandle := func(ptr interface{}) { o.Resume = *ptr.(*uint32) }
		if err := w.ReadPointer(&o.Resume, _s_ResumeHandle, _ptr_ResumeHandle); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// EnumExRequest structure represents the NetrDfsEnumEx operation request
type EnumExRequest struct {
	// DfsEntryPath: The pointer to a domain name, a host name, or a DFS path, depending
	// on the Level parameter.
	//
	// * A domain name MUST be a null-terminated Unicode ( a9bc4403-a862-48b9-b99b-1b44a887d177#gt_c305d0ab-8b94-461a-bd76-13b40cb8c4d8
	// ) string in the following forms:
	//
	// <DomainName> or \<DomainName> or \\<DomainName>
	//
	// where <DomainName> is the domain name to use for the enumeration.
	//
	// * A host name MUST be a null-terminated Unicode string in the following formats:
	//
	// <ServerName> or \<ServerName> or \\<ServerName>
	//
	// where <ServerName> is a host name.
	//
	// * A DFS root or a DFS link path.
	//
	// When DfsEntryPath points to a DFS link path, the remaining path after the DFS namespace
	// name MUST be ignored.
	EntryPath string `idl:"name:DfsEntryPath;string" json:"entry_path"`
	// Level: This parameter specifies the information level of the data and in turn determines
	// the action the method performs. On successful return, the server MUST return an array
	// of the corresponding structures in the buffer pointed to by DfsEnum.
	//
	//	+----------------------+----------------------------------------------------------------------------------+
	//	|                      |                                                                                  |
	//	|        VALUE         |                                     MEANING                                      |
	//	|                      |                                                                                  |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_1 0x00000001   | Gets the name of the DFS root and all links beneath it. In this case, on         |
	//	|                      | successful return DfsEnum MUST point to an array of DFS_INFO_1 structures.       |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_2 0x00000002   | Gets the name, comment, state, and number of targets for the DFS root and all    |
	//	|                      | links under the root. In this case, on successful return DfsEnum MUST point to   |
	//	|                      | an array of DFS_INFO_2 structures.                                               |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_3 0x00000003   | Gets the name, comment, state, number of targets, and information about          |
	//	|                      | each target for the DFS root and all links under the root. In this case, on      |
	//	|                      | successful return DfsEnum MUST point to an array of DFS_INFO_3 structures.       |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_4 0x00000004   | Gets the name, comment, state, time-out, GUID, number of targets, and            |
	//	|                      | information about each target for the DFS root and all links under the root.     |
	//	|                      | In this case, on successful return DfsEnum MUST point to an array of DFS_INFO_4  |
	//	|                      | structures.                                                                      |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_5 0x00000005   | Gets the name, comment, state, time-out, GUID, property flags, metadata size,    |
	//	|                      | and number of targets for a DFS root and all links under the root. In this case, |
	//	|                      | on successful return DfsEnum MUST point to an array of DFS_INFO_5 structures.    |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_6 0x00000006   | Gets the name, comment, state, time-out, GUID, property flags, metadata size,    |
	//	|                      | number of targets, and target information for a root or link. In this case, on   |
	//	|                      | successful return DfsEnum MUST point to an array of DFS_INFO_6 structures.       |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_8 0x00000008   | Gets the name, comment, state, time-out, GUID, property flags, metadata size,    |
	//	|                      | and number of targets for a DFS root and all links under the root. Also returns  |
	//	|                      | the security descriptor associated with each of the DFS links. In this case, on  |
	//	|                      | successful return DfsEnum MUST point to an array of DFS_INFO_8 structures.       |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_9 0x00000009   | Gets the name, comment, state, time-out, GUID, property flags, metadata size,    |
	//	|                      | and number of targets and target information for a DFS root and all links under  |
	//	|                      | the root. Also returns the security descriptor associated with each of the DFS   |
	//	|                      | links. In this case, on successful return DfsEnum MUST point to an array of      |
	//	|                      | DFS_INFO_9 structures.                                                           |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_200 0x000000C8 | Enumerates all of the domain-based DFS namespace in the specified domain. In     |
	//	|                      | this case, on successful return DfsEnum MUST point to an array of DFS_INFO_200   |
	//	|                      | structures.                                                                      |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_300 0x0000012C | Enumerates the stand-alone and domain-based DFS roots that the server hosts. In  |
	//	|                      | this case, on successful return DfsEnum MUST point to an array of DFS_INFO_300   |
	//	|                      | structures.                                                                      |
	//	+----------------------+----------------------------------------------------------------------------------+
	Level uint32 `idl:"name:Level" json:"level"`
	// PrefMaxLen: This parameter specifies restrictions on the number of elements returned.
	// A value of 0xFFFFFFFF means there are no restrictions, in which case all entries
	// MUST be returned.<96>
	PrefMaxLength uint32 `idl:"name:PrefMaxLen" json:"pref_max_length"`
	// DfsEnum: A pointer to a DFS_INFO_ENUM_STRUCT union to receive the returned information.
	// The client SHOULD set the Level member to the same value as the method's Level parameter,
	// and MUST set the DfsInfoContainer union member to a pointer to the corresponding
	// container structure as specified in section 2.2.6. The client MUST initialize the
	// container structure's EntriesRead member to zero and the Buffer member to a NULL
	// pointer. The value of the Level member determines the case of the union.
	Enum *InfoEnum `idl:"name:DfsEnum;pointer:unique" json:"enum"`
	// ResumeHandle: This parameter is used to continue an enumeration when more data is
	// available than can be returned in a single invocation of this method.
	//
	// * If this parameter is not a NULL pointer, and the method returns ERROR_SUCCESS,
	// this parameter receives an implementation-specific nonzero value that can be passed
	// in subsequent calls to this method to continue the enumeration.
	//
	// * If this parameter is a NULL pointer, or it points to a zero value, it indicates
	// that this is an initial enumeration.
	//
	// * If this parameter is not a NULL pointer, and it points to a nonzero value returned
	// in ResumeHandle by an earlier invocation of this method, the server will attempt
	// to continue a previous enumeration. <97> ( 3a466440-1ef6-4439-b4e2-be9eaddeb511#Appendix_A_97
	// )
	Resume uint32 `idl:"name:ResumeHandle;pointer:unique" json:"resume"`
}

func (o *EnumExRequest) xxx_ToOp(ctx context.Context, op *xxx_EnumExOperation) *xxx_EnumExOperation {
	if op == nil {
		op = &xxx_EnumExOperation{}
	}
	if o == nil {
		return op
	}
	op.EntryPath = o.EntryPath
	op.Level = o.Level
	op.PrefMaxLength = o.PrefMaxLength
	op.Enum = o.Enum
	op.Resume = o.Resume
	return op
}

func (o *EnumExRequest) xxx_FromOp(ctx context.Context, op *xxx_EnumExOperation) {
	if o == nil {
		return
	}
	o.EntryPath = op.EntryPath
	o.Level = op.Level
	o.PrefMaxLength = op.PrefMaxLength
	o.Enum = op.Enum
	o.Resume = op.Resume
}
func (o *EnumExRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *EnumExRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumExOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// EnumExResponse structure represents the NetrDfsEnumEx operation response
type EnumExResponse struct {
	// DfsEnum: A pointer to a DFS_INFO_ENUM_STRUCT union to receive the returned information.
	// The client SHOULD set the Level member to the same value as the method's Level parameter,
	// and MUST set the DfsInfoContainer union member to a pointer to the corresponding
	// container structure as specified in section 2.2.6. The client MUST initialize the
	// container structure's EntriesRead member to zero and the Buffer member to a NULL
	// pointer. The value of the Level member determines the case of the union.
	Enum *InfoEnum `idl:"name:DfsEnum;pointer:unique" json:"enum"`
	// ResumeHandle: This parameter is used to continue an enumeration when more data is
	// available than can be returned in a single invocation of this method.
	//
	// * If this parameter is not a NULL pointer, and the method returns ERROR_SUCCESS,
	// this parameter receives an implementation-specific nonzero value that can be passed
	// in subsequent calls to this method to continue the enumeration.
	//
	// * If this parameter is a NULL pointer, or it points to a zero value, it indicates
	// that this is an initial enumeration.
	//
	// * If this parameter is not a NULL pointer, and it points to a nonzero value returned
	// in ResumeHandle by an earlier invocation of this method, the server will attempt
	// to continue a previous enumeration. <97> ( 3a466440-1ef6-4439-b4e2-be9eaddeb511#Appendix_A_97
	// )
	Resume uint32 `idl:"name:ResumeHandle;pointer:unique" json:"resume"`
	// Return: The NetrDfsEnumEx return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *EnumExResponse) xxx_ToOp(ctx context.Context, op *xxx_EnumExOperation) *xxx_EnumExOperation {
	if op == nil {
		op = &xxx_EnumExOperation{}
	}
	if o == nil {
		return op
	}
	op.Enum = o.Enum
	op.Resume = o.Resume
	op.Return = o.Return
	return op
}

func (o *EnumExResponse) xxx_FromOp(ctx context.Context, op *xxx_EnumExOperation) {
	if o == nil {
		return
	}
	o.Enum = op.Enum
	o.Resume = op.Resume
	o.Return = op.Return
}
func (o *EnumExResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *EnumExResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_EnumExOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_SetInfo2Operation structure represents the NetrDfsSetInfo2 operation
type xxx_SetInfo2Operation struct {
	EntryPath  string    `idl:"name:DfsEntryPath;string" json:"entry_path"`
	DCName     string    `idl:"name:DcName;string" json:"dc_name"`
	ServerName string    `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	ShareName  string    `idl:"name:ShareName;string;pointer:unique" json:"share_name"`
	Level      uint32    `idl:"name:Level" json:"level"`
	DFSInfo    *Info     `idl:"name:pDfsInfo;switch_is:Level" json:"dfs_info"`
	RootList   *RootList `idl:"name:ppRootList;pointer:unique" json:"root_list"`
	Return     uint32    `idl:"name:Return" json:"return"`
}

func (o *xxx_SetInfo2Operation) OpNum() int { return 22 }

func (o *xxx_SetInfo2Operation) OpName() string { return "/netdfs/v3/NetrDfsSetInfo2" }

func (o *xxx_SetInfo2Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInfo2Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// DfsEntryPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.EntryPath); err != nil {
			return err
		}
	}
	// DcName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.WriteUTF16NString(ctx, w, o.DCName); err != nil {
			return err
		}
	}
	// ServerName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.ServerName != "" {
			_ptr_ServerName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ServerName); err != nil {
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
	// ShareName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.ShareName != "" {
			_ptr_ShareName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.ShareName); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.ShareName, _ptr_ShareName); err != nil {
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
	// pDfsInfo {in} (1:{pointer=ref}*(1))(2:{switch_type={}(uint32), alias=DFS_INFO_STRUCT}(union))
	{
		_swDFSInfo := uint32(o.Level)
		if o.DFSInfo != nil {
			if err := o.DFSInfo.MarshalUnionNDR(ctx, w, _swDFSInfo); err != nil {
				return err
			}
		} else {
			if err := (&Info{}).MarshalUnionNDR(ctx, w, _swDFSInfo); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppRootList {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=DFSM_ROOT_LIST}(struct))
	{
		if o.RootList != nil {
			_ptr_ppRootList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RootList != nil {
					_ptr_ppRootList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.RootList != nil {
							if err := o.RootList.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&RootList{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.RootList, _ptr_ppRootList); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RootList, _ptr_ppRootList); err != nil {
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

func (o *xxx_SetInfo2Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// DfsEntryPath {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.EntryPath); err != nil {
			return err
		}
	}
	// DcName {in} (1:{string, pointer=ref}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if err := ndr.ReadUTF16NString(ctx, w, &o.DCName); err != nil {
			return err
		}
	}
	// ServerName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ServerName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ServerName); err != nil {
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
	// ShareName {in} (1:{string, pointer=unique}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_ShareName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.ShareName); err != nil {
				return err
			}
			return nil
		})
		_s_ShareName := func(ptr interface{}) { o.ShareName = *ptr.(*string) }
		if err := w.ReadPointer(&o.ShareName, _s_ShareName, _ptr_ShareName); err != nil {
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
	// pDfsInfo {in} (1:{pointer=ref}*(1))(2:{switch_type={}(uint32), alias=DFS_INFO_STRUCT}(union))
	{
		if o.DFSInfo == nil {
			o.DFSInfo = &Info{}
		}
		_swDFSInfo := uint32(o.Level)
		if err := o.DFSInfo.UnmarshalUnionNDR(ctx, w, _swDFSInfo); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppRootList {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=DFSM_ROOT_LIST}(struct))
	{
		_ptr_ppRootList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppRootList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.RootList == nil {
					o.RootList = &RootList{}
				}
				if err := o.RootList.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppRootList := func(ptr interface{}) { o.RootList = *ptr.(**RootList) }
			if err := w.ReadPointer(&o.RootList, _s_ppRootList, _ptr_ppRootList); err != nil {
				return err
			}
			return nil
		})
		_s_ppRootList := func(ptr interface{}) { o.RootList = *ptr.(**RootList) }
		if err := w.ReadPointer(&o.RootList, _s_ppRootList, _ptr_ppRootList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInfo2Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInfo2Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// ppRootList {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=DFSM_ROOT_LIST}(struct))
	{
		if o.RootList != nil {
			_ptr_ppRootList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if o.RootList != nil {
					_ptr_ppRootList := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
						if o.RootList != nil {
							if err := o.RootList.MarshalNDR(ctx, w); err != nil {
								return err
							}
						} else {
							if err := (&RootList{}).MarshalNDR(ctx, w); err != nil {
								return err
							}
						}
						return nil
					})
					if err := w.WritePointer(&o.RootList, _ptr_ppRootList); err != nil {
						return err
					}
				} else {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.RootList, _ptr_ppRootList); err != nil {
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
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_SetInfo2Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// ppRootList {in, out} (1:{pointer=unique}*(2)*(1))(2:{alias=DFSM_ROOT_LIST}(struct))
	{
		_ptr_ppRootList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			_ptr_ppRootList := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if o.RootList == nil {
					o.RootList = &RootList{}
				}
				if err := o.RootList.UnmarshalNDR(ctx, w); err != nil {
					return err
				}
				return nil
			})
			_s_ppRootList := func(ptr interface{}) { o.RootList = *ptr.(**RootList) }
			if err := w.ReadPointer(&o.RootList, _s_ppRootList, _ptr_ppRootList); err != nil {
				return err
			}
			return nil
		})
		_s_ppRootList := func(ptr interface{}) { o.RootList = *ptr.(**RootList) }
		if err := w.ReadPointer(&o.RootList, _s_ppRootList, _ptr_ppRootList); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// SetInfo2Request structure represents the NetrDfsSetInfo2 operation request
type SetInfo2Request struct {
	// DfsEntryPath: The pointer to a DFS root path or a DFS link path that contains the
	// name of a DFS root or DFS link name.
	EntryPath string `idl:"name:DfsEntryPath;string" json:"entry_path"`
	// DcName: The pointer to a null-terminated Unicode string. It MUST be ignored for a
	// stand-alone DFS namespace. For a domain-based DFS namespace, this string contains
	// the host name of the DC that the DFS root target uses to get or update DFS metadata
	// for the DFS namespace. This parameter MAY be a NULL pointer; otherwise, it MUST be
	// the PDC for the domain of the DFS namespace.<99>
	DCName string `idl:"name:DcName;string" json:"dc_name"`
	// ServerName: The pointer to a null-terminated Unicode DFS root target or a DFS link
	// target host name string. This parameter MUST be a NULL pointer if the operation is
	// intended for a DFS root or a DFS link and not for targets.
	ServerName string `idl:"name:ServerName;string;pointer:unique" json:"server_name"`
	// ShareName: The pointer to a null-terminated Unicode DFS root target or a DFS link
	// target share name string. This parameter MUST be a NULL pointer if the operation
	// is intended for a DFS root or a DFS link and not for targets.
	ShareName string `idl:"name:ShareName;string;pointer:unique" json:"share_name"`
	// Level: This parameter specifies the information level of the data and, in turn, determines
	// the action the method performs.
	//
	//	+----------------------+----------------------------------------------------------------------------------+
	//	|                      |                                                                                  |
	//	|        VALUE         |                                     MEANING                                      |
	//	|                      |                                                                                  |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_100 0x00000064 | Sets the comment associated with the root or link that specified in DfsInfo.     |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_101 0x00000065 | Sets the storage state associated with the root, link, root target, or link      |
	//	|                      | target specified in DfsInfo.<100>                                                |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_102 0x00000066 | Sets the time-out value associated with the root or link specified in DfsInfo.   |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_103 0x00000067 | Sets the property flags for the root or link specified in DfsInfo.               |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_104 0x00000068 | Sets the target priority rank and class for the root target or link target       |
	//	|                      | specified in DfsInfo.                                                            |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_105 0x00000069 | Sets the comment, state, time-out information, and property flags for the root   |
	//	|                      | or link specified in DfsInfo. This does not apply to a root target or link       |
	//	|                      | target.                                                                          |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_106 0x0000006A | Sets the target state and priority for the root target or link target specified  |
	//	|                      | in DfsInfo. This does not apply to the DFS namespace root or link.<101>          |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_107 0x0000006B | Sets the comment, state, time-out, security descriptor information, and property |
	//	|                      | flags for the root or link specified in DfsInfo. Does not apply to a root        |
	//	|                      | target or link target. The ServerName and ShareName parameters MUST be NULL. The |
	//	|                      | security descriptor MUST NOT have an owner, group, or SACLs in it. The security  |
	//	|                      | descriptor MUST be a NULL, zero length value if used on a namespace root. In     |
	//	|                      | this case, note that it is equivalent to using Level_105.                        |
	//	+----------------------+----------------------------------------------------------------------------------+
	//	| Level_150 0x00000096 | Sets the security descriptor associated with a link. Only stand-alone DFS        |
	//	|                      | namespaces and domainv2-based DFS namespaces are supported. The ServerName and   |
	//	|                      | ShareName parameters MUST both be NULL. The security descriptor MUST NOT have an |
	//	|                      | owner, group, or SACLs in it.                                                    |
	//	+----------------------+----------------------------------------------------------------------------------+
	Level uint32 `idl:"name:Level" json:"level"`
	// pDfsInfo: The pointer to a DFS_INFO_STRUCT union that contains the specified data.
	// The Level parameter value determines the case of the union.
	DFSInfo *Info `idl:"name:pDfsInfo;switch_is:Level" json:"dfs_info"`
	// ppRootList: On success, returns a list of DFS root targets in the domain-based DFS
	// namespace which the client will be responsible for notifying about the change in
	// the DFS namespace. See section 3.2.4.2.3. This list MAY be empty if the server has
	// performed the notification.<104>
	RootList *RootList `idl:"name:ppRootList;pointer:unique" json:"root_list"`
}

func (o *SetInfo2Request) xxx_ToOp(ctx context.Context, op *xxx_SetInfo2Operation) *xxx_SetInfo2Operation {
	if op == nil {
		op = &xxx_SetInfo2Operation{}
	}
	if o == nil {
		return op
	}
	op.EntryPath = o.EntryPath
	op.DCName = o.DCName
	op.ServerName = o.ServerName
	op.ShareName = o.ShareName
	op.Level = o.Level
	op.DFSInfo = o.DFSInfo
	op.RootList = o.RootList
	return op
}

func (o *SetInfo2Request) xxx_FromOp(ctx context.Context, op *xxx_SetInfo2Operation) {
	if o == nil {
		return
	}
	o.EntryPath = op.EntryPath
	o.DCName = op.DCName
	o.ServerName = op.ServerName
	o.ShareName = op.ShareName
	o.Level = op.Level
	o.DFSInfo = op.DFSInfo
	o.RootList = op.RootList
}
func (o *SetInfo2Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *SetInfo2Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInfo2Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// SetInfo2Response structure represents the NetrDfsSetInfo2 operation response
type SetInfo2Response struct {
	// ppRootList: On success, returns a list of DFS root targets in the domain-based DFS
	// namespace which the client will be responsible for notifying about the change in
	// the DFS namespace. See section 3.2.4.2.3. This list MAY be empty if the server has
	// performed the notification.<104>
	RootList *RootList `idl:"name:ppRootList;pointer:unique" json:"root_list"`
	// Return: The NetrDfsSetInfo2 return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *SetInfo2Response) xxx_ToOp(ctx context.Context, op *xxx_SetInfo2Operation) *xxx_SetInfo2Operation {
	if op == nil {
		op = &xxx_SetInfo2Operation{}
	}
	if o == nil {
		return op
	}
	op.RootList = o.RootList
	op.Return = o.Return
	return op
}

func (o *SetInfo2Response) xxx_FromOp(ctx context.Context, op *xxx_SetInfo2Operation) {
	if o == nil {
		return
	}
	o.RootList = op.RootList
	o.Return = op.Return
}
func (o *SetInfo2Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *SetInfo2Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_SetInfo2Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_AddRootTargetOperation structure represents the NetrDfsAddRootTarget operation
type xxx_AddRootTargetOperation struct {
	DFSPath      string `idl:"name:pDfsPath;string;pointer:unique" json:"dfs_path"`
	TargetPath   string `idl:"name:pTargetPath;string;pointer:unique" json:"target_path"`
	MajorVersion uint32 `idl:"name:MajorVersion" json:"major_version"`
	Comment      string `idl:"name:pComment;string;pointer:unique" json:"comment"`
	NewNamespace bool   `idl:"name:NewNamespace" json:"new_namespace"`
	Flags        uint32 `idl:"name:Flags" json:"flags"`
	Return       uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_AddRootTargetOperation) OpNum() int { return 23 }

func (o *xxx_AddRootTargetOperation) OpName() string { return "/netdfs/v3/NetrDfsAddRootTarget" }

func (o *xxx_AddRootTargetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddRootTargetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pDfsPath {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.DFSPath != "" {
			_ptr_pDfsPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.DFSPath); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DFSPath, _ptr_pDfsPath); err != nil {
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
	// pTargetPath {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.TargetPath != "" {
			_ptr_pTargetPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.TargetPath); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.TargetPath, _ptr_pTargetPath); err != nil {
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
	// MajorVersion {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.MajorVersion); err != nil {
			return err
		}
	}
	// pComment {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.Comment != "" {
			_ptr_pComment := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Comment); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.Comment, _ptr_pComment); err != nil {
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
	// NewNamespace {in} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.NewNamespace); err != nil {
			return err
		}
	}
	// Flags {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddRootTargetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pDfsPath {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pDfsPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.DFSPath); err != nil {
				return err
			}
			return nil
		})
		_s_pDfsPath := func(ptr interface{}) { o.DFSPath = *ptr.(*string) }
		if err := w.ReadPointer(&o.DFSPath, _s_pDfsPath, _ptr_pDfsPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pTargetPath {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pTargetPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.TargetPath); err != nil {
				return err
			}
			return nil
		})
		_s_pTargetPath := func(ptr interface{}) { o.TargetPath = *ptr.(*string) }
		if err := w.ReadPointer(&o.TargetPath, _s_pTargetPath, _ptr_pTargetPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// MajorVersion {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.MajorVersion); err != nil {
			return err
		}
	}
	// pComment {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pComment := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Comment); err != nil {
				return err
			}
			return nil
		})
		_s_pComment := func(ptr interface{}) { o.Comment = *ptr.(*string) }
		if err := w.ReadPointer(&o.Comment, _s_pComment, _ptr_pComment); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// NewNamespace {in} (1:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.NewNamespace); err != nil {
			return err
		}
	}
	// Flags {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddRootTargetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddRootTargetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_AddRootTargetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// AddRootTargetRequest structure represents the NetrDfsAddRootTarget operation request
type AddRootTargetRequest struct {
	// pDfsPath: The pointer to a null-terminated Unicode string. This MUST be \\<domain>\<dfsname>
	// for domain-based DFS or \\<server>\<share> for stand-alone DFS.
	DFSPath string `idl:"name:pDfsPath;string;pointer:unique" json:"dfs_path"`
	// pTargetPath: The pointer to a null-terminated Unicode string. This MUST be \\<server>\<share>[\<path>]
	// for domain-based DFS or NULL for stand-alone DFS. The latter restriction is required
	// to ensure that a typographic error in the domain name, while attempting to create
	// a domain-based DFS, does not result in a stand-alone DFS namespace being created
	// on the DFS root target server, if the first pathname component of the pDfsPath parameter
	// is used to detect whether a domain-based DFS namespace or stand-alone DFS namespace
	// is being created. When pTargetPath is not NULL, the <server> MUST be used as the
	// host name of the new DFS root target in the metadata.
	TargetPath string `idl:"name:pTargetPath;string;pointer:unique" json:"target_path"`
	// MajorVersion: The DFS metadata version to use to create the DFS namespace. When adding
	// a DFS root target to an existing DFS namespace, MajorVersion MUST be either 0 or
	// the major version number of the existing DFS namespace. Otherwise, the call MUST
	// fail.
	MajorVersion uint32 `idl:"name:MajorVersion" json:"major_version"`
	// pComment: The pointer to a null-terminated Unicode string that contains a comment
	// associated with this root or link. This string has no protocol-specified restrictions
	// on length or content. The comment is meant for human consumption and does not affect
	// server functionality. The comment MUST be ignored when adding a target to an existing
	// link.
	Comment string `idl:"name:pComment;string;pointer:unique" json:"comment"`
	// NewNamespace: A Boolean value that, if TRUE, indicates a request to create a new
	// root. If FALSE, this value indicates a request to add a new root target to an existing
	// root.
	NewNamespace bool `idl:"name:NewNamespace" json:"new_namespace"`
	// Flags: This parameter MUST be zero for a domain-based DFS namespace and MUST be ignored
	// for a stand-alone DFS namespace.
	Flags uint32 `idl:"name:Flags" json:"flags"`
}

func (o *AddRootTargetRequest) xxx_ToOp(ctx context.Context, op *xxx_AddRootTargetOperation) *xxx_AddRootTargetOperation {
	if op == nil {
		op = &xxx_AddRootTargetOperation{}
	}
	if o == nil {
		return op
	}
	op.DFSPath = o.DFSPath
	op.TargetPath = o.TargetPath
	op.MajorVersion = o.MajorVersion
	op.Comment = o.Comment
	op.NewNamespace = o.NewNamespace
	op.Flags = o.Flags
	return op
}

func (o *AddRootTargetRequest) xxx_FromOp(ctx context.Context, op *xxx_AddRootTargetOperation) {
	if o == nil {
		return
	}
	o.DFSPath = op.DFSPath
	o.TargetPath = op.TargetPath
	o.MajorVersion = op.MajorVersion
	o.Comment = op.Comment
	o.NewNamespace = op.NewNamespace
	o.Flags = op.Flags
}
func (o *AddRootTargetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *AddRootTargetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddRootTargetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// AddRootTargetResponse structure represents the NetrDfsAddRootTarget operation response
type AddRootTargetResponse struct {
	// Return: The NetrDfsAddRootTarget return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *AddRootTargetResponse) xxx_ToOp(ctx context.Context, op *xxx_AddRootTargetOperation) *xxx_AddRootTargetOperation {
	if op == nil {
		op = &xxx_AddRootTargetOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *AddRootTargetResponse) xxx_FromOp(ctx context.Context, op *xxx_AddRootTargetOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *AddRootTargetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *AddRootTargetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_AddRootTargetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_RemoveRootTargetOperation structure represents the NetrDfsRemoveRootTarget operation
type xxx_RemoveRootTargetOperation struct {
	DFSPath    string `idl:"name:pDfsPath;string;pointer:unique" json:"dfs_path"`
	TargetPath string `idl:"name:pTargetPath;string;pointer:unique" json:"target_path"`
	Flags      uint32 `idl:"name:Flags" json:"flags"`
	Return     uint32 `idl:"name:Return" json:"return"`
}

func (o *xxx_RemoveRootTargetOperation) OpNum() int { return 24 }

func (o *xxx_RemoveRootTargetOperation) OpName() string { return "/netdfs/v3/NetrDfsRemoveRootTarget" }

func (o *xxx_RemoveRootTargetOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveRootTargetOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// pDfsPath {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.DFSPath != "" {
			_ptr_pDfsPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.DFSPath); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.DFSPath, _ptr_pDfsPath); err != nil {
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
	// pTargetPath {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		if o.TargetPath != "" {
			_ptr_pTargetPath := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.TargetPath); err != nil {
					return err
				}
				return nil
			})
			if err := w.WritePointer(&o.TargetPath, _ptr_pTargetPath); err != nil {
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
	// Flags {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveRootTargetOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// pDfsPath {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pDfsPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.DFSPath); err != nil {
				return err
			}
			return nil
		})
		_s_pDfsPath := func(ptr interface{}) { o.DFSPath = *ptr.(*string) }
		if err := w.ReadPointer(&o.DFSPath, _s_pDfsPath, _ptr_pDfsPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pTargetPath {in} (1:{string, pointer=unique, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		_ptr_pTargetPath := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.TargetPath); err != nil {
				return err
			}
			return nil
		})
		_s_pTargetPath := func(ptr interface{}) { o.TargetPath = *ptr.(*string) }
		if err := w.ReadPointer(&o.TargetPath, _s_pTargetPath, _ptr_pTargetPath); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Flags {in} (1:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.Flags); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveRootTargetOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveRootTargetOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_RemoveRootTargetOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// RemoveRootTargetRequest structure represents the NetrDfsRemoveRootTarget operation request
type RemoveRootTargetRequest struct {
	// pDfsPath: The pointer to a null-terminated Unicode string. This MUST be \\<domain>\<dfsname>
	// for domain-based DFS or \\<server>\<share> for stand-alone DFS.
	DFSPath string `idl:"name:pDfsPath;string;pointer:unique" json:"dfs_path"`
	// pTargetPath: The pointer to a null-terminated Unicode string. This MUST be \\<server>\<share>[\<path>]
	// for domain-based DFS or NULL for stand-alone DFS.
	TargetPath string `idl:"name:pTargetPath;string;pointer:unique" json:"target_path"`
	// Flags: A bit field specifying the type of removal operation. For a standalone namespace,
	// this bit-field parameter MUST be zero. For a domain-based DFS namespace, it can be
	// zero or the following value. Zero indicates a normal removal operation.
	//
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	|                             |                                                                                  |
	//	|            VALUE            |                                     MEANING                                      |
	//	|                             |                                                                                  |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------+----------------------------------------------------------------------------------+
	//	| DFS_FORCE_REMOVE 0x80000000 | Specifying this flag for a domain-based DFS namespace removes the root target    |
	//	|                             | even if it is not accessible.                                                    |
	//	+-----------------------------+----------------------------------------------------------------------------------+
	Flags uint32 `idl:"name:Flags" json:"flags"`
}

func (o *RemoveRootTargetRequest) xxx_ToOp(ctx context.Context, op *xxx_RemoveRootTargetOperation) *xxx_RemoveRootTargetOperation {
	if op == nil {
		op = &xxx_RemoveRootTargetOperation{}
	}
	if o == nil {
		return op
	}
	op.DFSPath = o.DFSPath
	op.TargetPath = o.TargetPath
	op.Flags = o.Flags
	return op
}

func (o *RemoveRootTargetRequest) xxx_FromOp(ctx context.Context, op *xxx_RemoveRootTargetOperation) {
	if o == nil {
		return
	}
	o.DFSPath = op.DFSPath
	o.TargetPath = op.TargetPath
	o.Flags = op.Flags
}
func (o *RemoveRootTargetRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *RemoveRootTargetRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveRootTargetOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// RemoveRootTargetResponse structure represents the NetrDfsRemoveRootTarget operation response
type RemoveRootTargetResponse struct {
	// Return: The NetrDfsRemoveRootTarget return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *RemoveRootTargetResponse) xxx_ToOp(ctx context.Context, op *xxx_RemoveRootTargetOperation) *xxx_RemoveRootTargetOperation {
	if op == nil {
		op = &xxx_RemoveRootTargetOperation{}
	}
	if o == nil {
		return op
	}
	op.Return = o.Return
	return op
}

func (o *RemoveRootTargetResponse) xxx_FromOp(ctx context.Context, op *xxx_RemoveRootTargetOperation) {
	if o == nil {
		return
	}
	o.Return = op.Return
}
func (o *RemoveRootTargetResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *RemoveRootTargetResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_RemoveRootTargetOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetSupportedNamespaceVersionOperation structure represents the NetrDfsGetSupportedNamespaceVersion operation
type xxx_GetSupportedNamespaceVersionOperation struct {
	Origin      NamespaceVersionOrigin         `idl:"name:Origin" json:"origin"`
	Name        string                         `idl:"name:pName;string;pointer:unique" json:"name"`
	VersionInfo *SupportedNamespaceVersionInfo `idl:"name:pVersionInfo" json:"version_info"`
	Return      uint32                         `idl:"name:Return" json:"return"`
}

func (o *xxx_GetSupportedNamespaceVersionOperation) OpNum() int { return 25 }

func (o *xxx_GetSupportedNamespaceVersionOperation) OpName() string {
	return "/netdfs/v3/NetrDfsGetSupportedNamespaceVersion"
}

func (o *xxx_GetSupportedNamespaceVersionOperation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSupportedNamespaceVersionOperation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// Origin {in} (1:{alias=DFS_NAMESPACE_VERSION_ORIGIN}(enum))
	{
		if err := w.WriteEnum(uint16(o.Origin)); err != nil {
			return err
		}
	}
	// pName {in} (1:{string, pointer=unique, alias=NETDFS_SERVER_OR_DOMAIN_HANDLE}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		if o.Name != "" {
			_ptr_pName := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				if err := ndr.WriteUTF16NString(ctx, w, o.Name); err != nil {
					return err
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
	return nil
}

func (o *xxx_GetSupportedNamespaceVersionOperation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// Origin {in} (1:{alias=DFS_NAMESPACE_VERSION_ORIGIN}(enum))
	{
		if err := w.ReadEnum((*uint16)(&o.Origin)); err != nil {
			return err
		}
	}
	// pName {in} (1:{string, pointer=unique, alias=NETDFS_SERVER_OR_DOMAIN_HANDLE}*(1))(2:{alias=WCHAR}[dim:0,string,null](wchar))
	{
		_ptr_pName := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			if err := ndr.ReadUTF16NString(ctx, w, &o.Name); err != nil {
				return err
			}
			return nil
		})
		_s_pName := func(ptr interface{}) { o.Name = *ptr.(*string) }
		if err := w.ReadPointer(&o.Name, _s_pName, _ptr_pName); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSupportedNamespaceVersionOperation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSupportedNamespaceVersionOperation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// pVersionInfo {out} (1:{alias=PDFS_SUPPORTED_NAMESPACE_VERSION_INFO}*(1))(2:{alias=DFS_SUPPORTED_NAMESPACE_VERSION_INFO}(struct))
	{
		if o.VersionInfo != nil {
			if err := o.VersionInfo.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&SupportedNamespaceVersionInfo{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetSupportedNamespaceVersionOperation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// pVersionInfo {out} (1:{alias=PDFS_SUPPORTED_NAMESPACE_VERSION_INFO,pointer=ref}*(1))(2:{alias=DFS_SUPPORTED_NAMESPACE_VERSION_INFO}(struct))
	{
		if o.VersionInfo == nil {
			o.VersionInfo = &SupportedNamespaceVersionInfo{}
		}
		if err := o.VersionInfo.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=NET_API_STATUS, names=DWORD}(uint32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetSupportedNamespaceVersionRequest structure represents the NetrDfsGetSupportedNamespaceVersion operation request
type GetSupportedNamespaceVersionRequest struct {
	// Origin: This parameter specifies the version information requested.
	//
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	|                                            |                                                                                  |
	//	|                   VALUE                    |                                     MEANING                                      |
	//	|                                            |                                                                                  |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_NAMESPACE_VERSION_ORIGIN_SERVER 0x0001 | This specifies that the returned information MUST reflect the metadata versions  |
	//	|                                            | supported by the server. Versions supported by the server can be higher (or      |
	//	|                                            | lower) than those supported by the domain.                                       |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	//	| DFS_NAMESPACE_VERSION_ORIGIN_DOMAIN 0x0002 | This specifies that the returned information MUST reflect the metadata versions  |
	//	|                                            | supported by the domain schema of the domain to which the server is joined.      |
	//	|                                            | Versions supported by the domain schema can be higher (or lower) than those      |
	//	|                                            | supported by the server.                                                         |
	//	+--------------------------------------------+----------------------------------------------------------------------------------+
	Origin NamespaceVersionOrigin `idl:"name:Origin" json:"origin"`
	// pName: The pointer to a null-terminated Unicode string. The server MUST ignore the
	// pName parameter.
	Name string `idl:"name:pName;string;pointer:unique" json:"name"`
}

func (o *GetSupportedNamespaceVersionRequest) xxx_ToOp(ctx context.Context, op *xxx_GetSupportedNamespaceVersionOperation) *xxx_GetSupportedNamespaceVersionOperation {
	if op == nil {
		op = &xxx_GetSupportedNamespaceVersionOperation{}
	}
	if o == nil {
		return op
	}
	op.Origin = o.Origin
	op.Name = o.Name
	return op
}

func (o *GetSupportedNamespaceVersionRequest) xxx_FromOp(ctx context.Context, op *xxx_GetSupportedNamespaceVersionOperation) {
	if o == nil {
		return
	}
	o.Origin = op.Origin
	o.Name = op.Name
}
func (o *GetSupportedNamespaceVersionRequest) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRRequest(ctx, w)
}
func (o *GetSupportedNamespaceVersionRequest) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSupportedNamespaceVersionOperation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetSupportedNamespaceVersionResponse structure represents the NetrDfsGetSupportedNamespaceVersion operation response
type GetSupportedNamespaceVersionResponse struct {
	// pVersionInfo: The pointer to a DFS_SUPPORTED_NAMESPACE_VERSION_INFO structure to
	// receive the DFS metadata version number determined.
	VersionInfo *SupportedNamespaceVersionInfo `idl:"name:pVersionInfo" json:"version_info"`
	// Return: The NetrDfsGetSupportedNamespaceVersion return value.
	Return uint32 `idl:"name:Return" json:"return"`
}

func (o *GetSupportedNamespaceVersionResponse) xxx_ToOp(ctx context.Context, op *xxx_GetSupportedNamespaceVersionOperation) *xxx_GetSupportedNamespaceVersionOperation {
	if op == nil {
		op = &xxx_GetSupportedNamespaceVersionOperation{}
	}
	if o == nil {
		return op
	}
	op.VersionInfo = o.VersionInfo
	op.Return = o.Return
	return op
}

func (o *GetSupportedNamespaceVersionResponse) xxx_FromOp(ctx context.Context, op *xxx_GetSupportedNamespaceVersionOperation) {
	if o == nil {
		return
	}
	o.VersionInfo = op.VersionInfo
	o.Return = op.Return
}
func (o *GetSupportedNamespaceVersionResponse) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx, nil).MarshalNDRResponse(ctx, w)
}
func (o *GetSupportedNamespaceVersionResponse) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetSupportedNamespaceVersionOperation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
