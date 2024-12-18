package iclusterstorage3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	dcom "github.com/oiweiwei/go-msrpc/msrpc/dcom"
	csvp "github.com/oiweiwei/go-msrpc/msrpc/dcom/csvp"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
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
	_ = dcom.GoPackage
	_ = iunknown.GoPackage
	_ = csvp.GoPackage
	_ = dtyp.GoPackage
)

var (
	// import guard
	GoPackage = "dcom/csvp"
)

var (
	// IClusterStorage3 interface identifier 11942d87-a1de-4e7f-83fb-a840d9c5928d
	ClusterStorage3IID = &dcom.IID{Data1: 0x11942d87, Data2: 0xa1de, Data3: 0x4e7f, Data4: []byte{0x83, 0xfb, 0xa8, 0x40, 0xd9, 0xc5, 0x92, 0x8d}}
	// Syntax UUID
	ClusterStorage3SyntaxUUID = &uuid.UUID{TimeLow: 0x11942d87, TimeMid: 0xa1de, TimeHiAndVersion: 0x4e7f, ClockSeqHiAndReserved: 0x83, ClockSeqLow: 0xfb, Node: [6]uint8{0xa8, 0x40, 0xd9, 0xc5, 0x92, 0x8d}}
	// Syntax ID
	ClusterStorage3SyntaxV0_0 = &dcerpc.SyntaxID{IfUUID: ClusterStorage3SyntaxUUID, IfVersionMajor: 0, IfVersionMinor: 0}
)

// IClusterStorage3 interface.
type ClusterStorage3Client interface {

	// IUnknown retrieval method.
	Unknown() iunknown.UnknownClient

	// The CprepDiskGetUniqueIds3 method returns device ID data about the ClusPrepDisk.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+--------------------------------+-------------------------------------------+
	//	|             RETURN             |                                           |
	//	|           VALUE/CODE           |                DESCRIPTION                |
	//	|                                |                                           |
	//	+--------------------------------+-------------------------------------------+
	//	+--------------------------------+-------------------------------------------+
	//	| 0x00000000 S_OK                | The call was successful.                  |
	//	+--------------------------------+-------------------------------------------+
	//	| 0x80070057 E_INVALIDARG        | One or more arguments are invalid.        |
	//	+--------------------------------+-------------------------------------------+
	//	| 0x80070032 ERROR_NOT_SUPPORTED | The disk does not support device ID data. |
	//	+--------------------------------+-------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 3.
	GetUniqueIDs3(context.Context, *GetUniqueIDs3Request, ...dcerpc.CallOption) (*GetUniqueIDs3Response, error)

	// The CprepCheckNetFtBindings3 method verifies that an implementation-specific mechanism
	// is available for use as a network file sharing protocol.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 4.
	CheckNetFTBindings3(context.Context, *CheckNetFTBindings3Request, ...dcerpc.CallOption) (*CheckNetFTBindings3Response, error)

	// ClusterFileShareTestSetupState transitions to ClusterFileShareTestSetup.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 5.
	CSVTestSetup3(context.Context, *CSVTestSetup3Request, ...dcerpc.CallOption) (*CSVTestSetup3Response, error)

	// The CprepIsNodeClustered3 method determines whether the server is a node within a
	// cluster.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+-------------------------+------------------------------------+
	//	|         RETURN          |                                    |
	//	|       VALUE/CODE        |            DESCRIPTION             |
	//	|                         |                                    |
	//	+-------------------------+------------------------------------+
	//	+-------------------------+------------------------------------+
	//	| 0x00000000 S_OK         | The call was successful.           |
	//	+-------------------------+------------------------------------+
	//	| 0x80070057 E_INVALIDARG | One or more arguments are invalid. |
	//	+-------------------------+------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 6.
	IsNodeClustered3(context.Context, *IsNodeClustered3Request, ...dcerpc.CallOption) (*IsNodeClustered3Response, error)

	// The CprepCreateNewSmbShares3 method retrieves the list of IP addresses, with \\ prepended.
	// This method can be used to access a share via an implementation-specific mechanism.
	//
	// # The output strings have the form
	//
	// * \\<IPv4 address>
	//
	// or
	//
	// * \\[<IPv6 address>].
	//
	// HRESULT CprepCreateNewSmbShares3(
	//
	// * LPWSTR** ppwszSharePaths,
	//
	// * );
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+-------------------------+------------------------------------+
	//	|         RETURN          |                                    |
	//	|       VALUE/CODE        |            DESCRIPTION             |
	//	|                         |                                    |
	//	+-------------------------+------------------------------------+
	//	+-------------------------+------------------------------------+
	//	| 0x00000000 S_OK         | The call was successful.           |
	//	+-------------------------+------------------------------------+
	//	| 0x80070057 E_INVALIDARG | One or more arguments are invalid. |
	//	+-------------------------+------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 7.
	CreateNewSMBShares3(context.Context, *CreateNewSMBShares3Request, ...dcerpc.CallOption) (*CreateNewSMBShares3Response, error)

	// The CprepConnectToNewSmbShares3 method attempts to connect to shares represented
	// by ClusPrepShares in the ClusprepShareList at the given list of IP address strings.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	|         RETURN          |                                                                                  |
	//	|       VALUE/CODE        |                                   DESCRIPTION                                    |
	//	|                         |                                                                                  |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 S_OK         | The call was successful. The connection used at least one of the file shares     |
	//	|                         | specified in ppwszSharePaths.                                                    |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//	| 0x80070057 E_INVALIDARG | One or more arguments are invalid.                                               |
	//	+-------------------------+----------------------------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 8.
	ConnectToNewSMBShares3(context.Context, *ConnectToNewSMBShares3Request, ...dcerpc.CallOption) (*ConnectToNewSMBShares3Response, error)

	// The CprepDiskGetProps3 method retrieves information about the configuration and status
	// of a given disk.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------+--------------------------+
	//	|             RETURN              |                          |
	//	|           VALUE/CODE            |       DESCRIPTION        |
	//	|                                 |                          |
	//	+---------------------------------+--------------------------+
	//	+---------------------------------+--------------------------+
	//	| 0x00000000 S_OK                 | The call was successful. |
	//	+---------------------------------+--------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND | The disk was not found.  |
	//	+---------------------------------+--------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 9.
	GetProperties3(context.Context, *GetProperties3Request, ...dcerpc.CallOption) (*GetProperties3Response, error)

	// The CprepDiskIsReadOnly3 method returns a Boolean value indicating whether the LUN
	// underlying the operating system disk is writable.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Zero or positive values indicate success,
	// with the lower 16 bits in positive nonzero values containing warnings or flags defined
	// in the method implementation. For more information about Win32 error codes and HRESULT
	// values, see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+-------------------------------------------+
	//	|                RETURN                 |                                           |
	//	|              VALUE/CODE               |                DESCRIPTION                |
	//	|                                       |                                           |
	//	+---------------------------------------+-------------------------------------------+
	//	+---------------------------------------+-------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                  |
	//	+---------------------------------------+-------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found.                   |
	//	+---------------------------------------+-------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online. |
	//	+---------------------------------------+-------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	//
	// The opnum field value for this method is 10.
	IsReadOnly3(context.Context, *IsReadOnly3Request, ...dcerpc.CallOption) (*IsReadOnly3Response, error)

	// The CprepDiskPRRegister3 method performs a SCSI PERSISTENT RESERVE OUT command (see
	// [SPC-3] section 6.12) with a REGISTER or REGISTER_IGNORE_EXISTING action.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Non-negative values indicate success, with
	// the lower 16 bits of the value containing warnings or flags defined in the method
	// implementation. For more information about Win32 error codes and HRESULT values,
	// see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+----------------------------------------------------------+
	//	|                RETURN                 |                                                          |
	//	|              VALUE/CODE               |                       DESCRIPTION                        |
	//	|                                       |                                                          |
	//	+---------------------------------------+----------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                 |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found.                                  |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER    | An internal error occurred.                              |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The value of ClusPrepDisk.AttachedState is not Attached. |
	//	+---------------------------------------+----------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol (see [MS-RPCE]).
	//
	// The opnum field value for this method is 11.
	PRRegister3(context.Context, *PRRegister3Request, ...dcerpc.CallOption) (*PRRegister3Response, error)

	// The CprepDiskFindKey3 method queries the SCSI Persistent Reserve registration table
	// for the disk and determines if the specified key is listed in the table.
	//
	// The user supplies a key and this method returns a BOOLEAN indicating whether the
	// key is found in the registration table.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Non-negative values indicate success, with
	// the lower 16 bits of the value containing warnings or flags defined in the method
	// implementation. For more information about Win32 error codes and HRESULT values,
	// see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+----------------------------------------------------------+
	//	|                RETURN                 |                                                          |
	//	|              VALUE/CODE               |                       DESCRIPTION                        |
	//	|                                       |                                                          |
	//	+---------------------------------------+----------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                 |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found.                                  |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER    | An internal error occurred.                              |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The value of ClusPrepDisk.AttachedState is not Attached. |
	//	+---------------------------------------+----------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol (see [MS-RPCE]).
	//
	// The opnum field value for this method is 12.
	FindKey3(context.Context, *FindKey3Request, ...dcerpc.CallOption) (*FindKey3Response, error)

	// The CprepDiskPRPreempt3 method performs a SCSI PERSISTENT RESERVE OUT command (see
	// [SPC-3] section 6.12) with a PREEMPT action.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Non-negative values indicate success, with
	// the lower 16 bits of the value containing warnings or flags defined in the method
	// implementation. For more information about Win32 error codes and HRESULT values,
	// see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+----------------------------------------------------------+
	//	|                RETURN                 |                                                          |
	//	|              VALUE/CODE               |                       DESCRIPTION                        |
	//	|                                       |                                                          |
	//	+---------------------------------------+----------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                 |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found.                                  |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER    | An internal error occurred.                              |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The value of ClusPrepDisk.AttachedState is not Attached. |
	//	+---------------------------------------+----------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol (see [MS-RPCE]).
	//
	// The opnum field value for this method is 13.
	PRPreempt3(context.Context, *PRPreempt3Request, ...dcerpc.CallOption) (*PRPreempt3Response, error)

	// The CprepDiskPRReserve3 method performs a SCSI PERSISTENT RESERVE OUT command (see
	// [SPC-3] section 6.12) with a RESERVE action.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Non-negative values indicate success, with
	// the lower 16 bits of the value containing warnings or flags defined in the method
	// implementation. For more information about Win32 error codes and HRESULT values,
	// see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+----------------------------------------------------------+
	//	|                RETURN                 |                                                          |
	//	|              VALUE/CODE               |                       DESCRIPTION                        |
	//	|                                       |                                                          |
	//	+---------------------------------------+----------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                 |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found.                                  |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER    | An internal error occurred.                              |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The value of ClusPrepDisk.AttachedState is not Attached. |
	//	+---------------------------------------+----------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol (see [MS-RPCE]).
	//
	// The opnum field value for this method is 14.
	PRReserve3(context.Context, *PRReserve3Request, ...dcerpc.CallOption) (*PRReserve3Response, error)

	// The CprepDiskIsPRPresent3 method queries the SCSI Persistent Reserve reservation
	// table for the disk and determines if the specified key is listed in the table.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Non-negative values indicate success, with
	// the lower 16 bits of the value containing warnings or flags defined in the method
	// implementation. For more information about Win32 error codes and HRESULT values,
	// see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+----------------------------------------------------------+
	//	|                RETURN                 |                                                          |
	//	|              VALUE/CODE               |                       DESCRIPTION                        |
	//	|                                       |                                                          |
	//	+---------------------------------------+----------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                 |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x80070490 ERROR_NOT_FOUND            | The key was not found in the reservation table.          |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found.                                  |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER    | An internal error occurred.                              |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The value of ClusPrepDisk.AttachedState is not Attached. |
	//	+---------------------------------------+----------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol (see [MS-RPCE]).
	//
	// The opnum field value for this method is 15.
	IsPRPresent3(context.Context, *IsPRPresent3Request, ...dcerpc.CallOption) (*IsPRPresent3Response, error)

	// The CprepDiskPRRelease3 method performs a SCSI PERSISTENT RESERVE OUT command (see
	// [SPC-3] section 6.12) with a RELEASE action.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Non-negative values indicate success, with
	// the lower 16 bits of the value containing warnings or flags defined in the method
	// implementation. For more information about Win32 error codes and HRESULT values,
	// see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+----------------------------------------------------------+
	//	|                RETURN                 |                                                          |
	//	|              VALUE/CODE               |                       DESCRIPTION                        |
	//	|                                       |                                                          |
	//	+---------------------------------------+----------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                 |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found.                                  |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER    | An internal error occurred.                              |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The value of ClusPrepDisk.AttachedState is not Attached. |
	//	+---------------------------------------+----------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol (see [MS-RPCE]).
	//
	// The opnum field value for this method is 16.
	PRRelease3(context.Context, *PRRelease3Request, ...dcerpc.CallOption) (*PRRelease3Response, error)

	// The CprepDiskPRClear3 method performs a SCSI PERSISTENT RESERVE OUT command (see
	// [SPC-3] section 6.12) with a CLEAR action.
	//
	// Return Values: A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it has failed. Non-negative values indicate success, with
	// the lower 16 bits of the value containing warnings or flags defined in the method
	// implementation. For more information about Win32 error codes and HRESULT values,
	// see [MS-ERREF] sections 2.2 and 2.1.
	//
	//	+---------------------------------------+----------------------------------------------------------+
	//	|                RETURN                 |                                                          |
	//	|              VALUE/CODE               |                       DESCRIPTION                        |
	//	|                                       |                                                          |
	//	+---------------------------------------+----------------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x00000000 S_OK                       | The call was successful.                                 |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x80070002 ERROR_FILE_NOT_FOUND       | The disk was not found.                                  |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER    | An internal error occurred.                              |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x80070548 ERROR_INVALID_SERVER_STATE | The server's Prepare State is not Online.                |
	//	+---------------------------------------+----------------------------------------------------------+
	//	| 0x8007139F ERROR_INVALID_STATE        | The value of ClusPrepDisk.AttachedState is not Attached. |
	//	+---------------------------------------+----------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol (see [MS-RPCE]).
	//
	// The opnum field value for this method is 17.
	PRClear3(context.Context, *PRClear3Request, ...dcerpc.CallOption) (*PRClear3Response, error)

	// AlterContext alters the client context.
	AlterContext(context.Context, ...dcerpc.Option) error

	// Conn returns the client connection (unsafe)
	Conn() dcerpc.Conn

	// IPID sets the object interface identifier.
	IPID(context.Context, *dcom.IPID) ClusterStorage3Client
}

type xxx_DefaultClusterStorage3Client struct {
	iunknown.UnknownClient
	cc   dcerpc.Conn
	ipid *dcom.IPID
}

func (o *xxx_DefaultClusterStorage3Client) Unknown() iunknown.UnknownClient {
	return o.UnknownClient
}

func (o *xxx_DefaultClusterStorage3Client) GetUniqueIDs3(ctx context.Context, in *GetUniqueIDs3Request, opts ...dcerpc.CallOption) (*GetUniqueIDs3Response, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetUniqueIDs3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage3Client) CheckNetFTBindings3(ctx context.Context, in *CheckNetFTBindings3Request, opts ...dcerpc.CallOption) (*CheckNetFTBindings3Response, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CheckNetFTBindings3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage3Client) CSVTestSetup3(ctx context.Context, in *CSVTestSetup3Request, opts ...dcerpc.CallOption) (*CSVTestSetup3Response, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CSVTestSetup3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage3Client) IsNodeClustered3(ctx context.Context, in *IsNodeClustered3Request, opts ...dcerpc.CallOption) (*IsNodeClustered3Response, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &IsNodeClustered3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage3Client) CreateNewSMBShares3(ctx context.Context, in *CreateNewSMBShares3Request, opts ...dcerpc.CallOption) (*CreateNewSMBShares3Response, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &CreateNewSMBShares3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage3Client) ConnectToNewSMBShares3(ctx context.Context, in *ConnectToNewSMBShares3Request, opts ...dcerpc.CallOption) (*ConnectToNewSMBShares3Response, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &ConnectToNewSMBShares3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage3Client) GetProperties3(ctx context.Context, in *GetProperties3Request, opts ...dcerpc.CallOption) (*GetProperties3Response, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &GetProperties3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage3Client) IsReadOnly3(ctx context.Context, in *IsReadOnly3Request, opts ...dcerpc.CallOption) (*IsReadOnly3Response, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &IsReadOnly3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage3Client) PRRegister3(ctx context.Context, in *PRRegister3Request, opts ...dcerpc.CallOption) (*PRRegister3Response, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PRRegister3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage3Client) FindKey3(ctx context.Context, in *FindKey3Request, opts ...dcerpc.CallOption) (*FindKey3Response, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &FindKey3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage3Client) PRPreempt3(ctx context.Context, in *PRPreempt3Request, opts ...dcerpc.CallOption) (*PRPreempt3Response, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PRPreempt3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage3Client) PRReserve3(ctx context.Context, in *PRReserve3Request, opts ...dcerpc.CallOption) (*PRReserve3Response, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PRReserve3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage3Client) IsPRPresent3(ctx context.Context, in *IsPRPresent3Request, opts ...dcerpc.CallOption) (*IsPRPresent3Response, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &IsPRPresent3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage3Client) PRRelease3(ctx context.Context, in *PRRelease3Request, opts ...dcerpc.CallOption) (*PRRelease3Response, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PRRelease3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage3Client) PRClear3(ctx context.Context, in *PRClear3Request, opts ...dcerpc.CallOption) (*PRClear3Response, error) {
	op := in.xxx_ToOp(ctx)
	if _, ok := dcom.HasIPID(opts); !ok {
		if o.ipid != nil {
			opts = append(opts, dcom.WithIPID(o.ipid))
		} else {
			return nil, fmt.Errorf("%s: ipid is missing", op.OpName())
		}
	}
	if err := o.cc.Invoke(ctx, op, opts...); err != nil {
		return nil, err
	}
	out := &PRClear3Response{}
	out.xxx_FromOp(ctx, op)
	if op.Return != int32(0) {
		return out, fmt.Errorf("%s: %w", op.OpName(), errors.New(ctx, op.Return))
	}
	return out, nil
}

func (o *xxx_DefaultClusterStorage3Client) AlterContext(ctx context.Context, opts ...dcerpc.Option) error {
	return o.cc.AlterContext(ctx, opts...)
}

func (o *xxx_DefaultClusterStorage3Client) Conn() dcerpc.Conn {
	return o.cc
}

func (o *xxx_DefaultClusterStorage3Client) IPID(ctx context.Context, ipid *dcom.IPID) ClusterStorage3Client {
	if ipid == nil {
		ipid = &dcom.IPID{}
	}
	return &xxx_DefaultClusterStorage3Client{
		UnknownClient: o.UnknownClient.IPID(ctx, ipid),
		cc:            o.cc,
		ipid:          ipid,
	}
}

func NewClusterStorage3Client(ctx context.Context, cc dcerpc.Conn, opts ...dcerpc.Option) (ClusterStorage3Client, error) {
	var err error
	if !dcom.IsSuperclass(opts) {
		cc, err = cc.Bind(ctx, append(opts, dcerpc.WithAbstractSyntax(ClusterStorage3SyntaxV0_0))...)
		if err != nil {
			return nil, err
		}
	}
	base, err := iunknown.NewUnknownClient(ctx, cc, append(opts, dcom.Superclass(cc))...)
	if err != nil {
		return nil, err
	}
	ipid, ok := dcom.HasIPID(opts)
	if ok {
		base = base.IPID(ctx, ipid)
	}
	return &xxx_DefaultClusterStorage3Client{
		UnknownClient: base,
		cc:            cc,
		ipid:          ipid,
	}, nil
}

// xxx_GetUniqueIDs3Operation structure represents the CprepDiskGetUniqueIds3 operation
type xxx_GetUniqueIDs3Operation struct {
	This                   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That                   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID                 *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	DeviceIDHeader         []byte              `idl:"name:ppbDeviceIdHeader;size_is:(, pcbDihSize)" json:"device_id_header"`
	DeviceIDHeaderLength   uint32              `idl:"name:pcbDihSize" json:"device_id_header_length"`
	DeviceDescriptor       []byte              `idl:"name:ppDeviceDescriptor;size_is:(, pcbDdSize)" json:"device_descriptor"`
	DeviceDescriptorLength uint32              `idl:"name:pcbDdSize" json:"device_descriptor_length"`
	Return                 int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_GetUniqueIDs3Operation) OpNum() int { return 3 }

func (o *xxx_GetUniqueIDs3Operation) OpName() string {
	return "/IClusterStorage3/v0/CprepDiskGetUniqueIds3"
}

func (o *xxx_GetUniqueIDs3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUniqueIDs3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DiskId {in} (1:{alias=CPREP_DISKID}(struct))
	{
		if o.DiskID != nil {
			if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.ClusterDiskID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetUniqueIDs3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DiskId {in} (1:{alias=CPREP_DISKID}(struct))
	{
		if o.DiskID == nil {
			o.DiskID = &csvp.ClusterDiskID{}
		}
		if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUniqueIDs3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.DeviceIDHeader != nil && o.DeviceIDHeaderLength == 0 {
		o.DeviceIDHeaderLength = uint32(len(o.DeviceIDHeader))
	}
	if o.DeviceDescriptor != nil && o.DeviceDescriptorLength == 0 {
		o.DeviceDescriptorLength = uint32(len(o.DeviceDescriptor))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUniqueIDs3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppbDeviceIdHeader {out} (1:{pointer=ref}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=pcbDihSize](uint8))
	{
		if o.DeviceIDHeader != nil || o.DeviceIDHeaderLength > 0 {
			_ptr_ppbDeviceIdHeader := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.DeviceIDHeaderLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.DeviceIDHeader {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.DeviceIDHeader[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.DeviceIDHeader); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DeviceIDHeader, _ptr_ppbDeviceIdHeader); err != nil {
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
	// pcbDihSize {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.DeviceIDHeaderLength); err != nil {
			return err
		}
	}
	// ppDeviceDescriptor {out} (1:{pointer=ref}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=pcbDdSize](uint8))
	{
		if o.DeviceDescriptor != nil || o.DeviceDescriptorLength > 0 {
			_ptr_ppDeviceDescriptor := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.DeviceDescriptorLength)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.DeviceDescriptor {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if err := w.WriteData(o.DeviceDescriptor[i1]); err != nil {
						return err
					}
				}
				for i1 := len(o.DeviceDescriptor); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WriteData(uint8(0)); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.DeviceDescriptor, _ptr_ppDeviceDescriptor); err != nil {
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
	// pcbDdSize {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.WriteData(o.DeviceDescriptorLength); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetUniqueIDs3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppbDeviceIdHeader {out} (1:{pointer=ref}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=pcbDihSize](uint8))
	{
		_ptr_ppbDeviceIdHeader := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.DeviceIDHeader", sizeInfo[0])
			}
			o.DeviceIDHeader = make([]byte, sizeInfo[0])
			for i1 := range o.DeviceIDHeader {
				i1 := i1
				if err := w.ReadData(&o.DeviceIDHeader[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppbDeviceIdHeader := func(ptr interface{}) { o.DeviceIDHeader = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.DeviceIDHeader, _s_ppbDeviceIdHeader, _ptr_ppbDeviceIdHeader); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcbDihSize {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.DeviceIDHeaderLength); err != nil {
			return err
		}
	}
	// ppDeviceDescriptor {out} (1:{pointer=ref}*(2)*(1))(2:{alias=BYTE}[dim:0,size_is=pcbDdSize](uint8))
	{
		_ptr_ppDeviceDescriptor := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.DeviceDescriptor", sizeInfo[0])
			}
			o.DeviceDescriptor = make([]byte, sizeInfo[0])
			for i1 := range o.DeviceDescriptor {
				i1 := i1
				if err := w.ReadData(&o.DeviceDescriptor[i1]); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppDeviceDescriptor := func(ptr interface{}) { o.DeviceDescriptor = *ptr.(*[]byte) }
		if err := w.ReadPointer(&o.DeviceDescriptor, _s_ppDeviceDescriptor, _ptr_ppDeviceDescriptor); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pcbDdSize {out} (1:{pointer=ref}*(1))(2:{alias=ULONG}(uint32))
	{
		if err := w.ReadData(&o.DeviceDescriptorLength); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetUniqueIDs3Request structure represents the CprepDiskGetUniqueIds3 operation request
type GetUniqueIDs3Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier representing the ClusPrepDisk for which to retrieve the device
	// ID data.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *GetUniqueIDs3Request) xxx_ToOp(ctx context.Context) *xxx_GetUniqueIDs3Operation {
	if o == nil {
		return &xxx_GetUniqueIDs3Operation{}
	}
	return &xxx_GetUniqueIDs3Operation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *GetUniqueIDs3Request) xxx_FromOp(ctx context.Context, op *xxx_GetUniqueIDs3Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *GetUniqueIDs3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetUniqueIDs3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUniqueIDs3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetUniqueIDs3Response structure represents the CprepDiskGetUniqueIds3 operation response
type GetUniqueIDs3Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppbDeviceIdHeader: On successful return, the address of a pointer to a *pcbDihSize-sized
	// block of BYTEs. The server allocates and initializes the returned buffer. Callers
	// MUST free this memory when they are finished with it. On unsuccessful return, the
	// client MUST ignore this value.
	DeviceIDHeader []byte `idl:"name:ppbDeviceIdHeader;size_is:(, pcbDihSize)" json:"device_id_header"`
	// pcbDihSize: On successful return, the number of BYTEs returned in ppbDeviceIdHeader.
	// On unsuccessful return, the client MUST ignore this value.
	DeviceIDHeaderLength uint32 `idl:"name:pcbDihSize" json:"device_id_header_length"`
	// ppDeviceDescriptor: On successful return, the address of a pointer to a *pcbDdSize-sized
	// block of BYTEs. The server allocates and initializes the returned buffer. Callers
	// MUST free this memory when they are finished with it. On unsuccessful return, the
	// client MUST ignore this value.
	DeviceDescriptor []byte `idl:"name:ppDeviceDescriptor;size_is:(, pcbDdSize)" json:"device_descriptor"`
	// pcbDdSize: On successful return, the number of BYTEs returned in ppDeviceDescriptor.
	// On unsuccessful return, the client MUST ignore this value.
	DeviceDescriptorLength uint32 `idl:"name:pcbDdSize" json:"device_descriptor_length"`
	// Return: The CprepDiskGetUniqueIds3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetUniqueIDs3Response) xxx_ToOp(ctx context.Context) *xxx_GetUniqueIDs3Operation {
	if o == nil {
		return &xxx_GetUniqueIDs3Operation{}
	}
	return &xxx_GetUniqueIDs3Operation{
		That:                   o.That,
		DeviceIDHeader:         o.DeviceIDHeader,
		DeviceIDHeaderLength:   o.DeviceIDHeaderLength,
		DeviceDescriptor:       o.DeviceDescriptor,
		DeviceDescriptorLength: o.DeviceDescriptorLength,
		Return:                 o.Return,
	}
}

func (o *GetUniqueIDs3Response) xxx_FromOp(ctx context.Context, op *xxx_GetUniqueIDs3Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DeviceIDHeader = op.DeviceIDHeader
	o.DeviceIDHeaderLength = op.DeviceIDHeaderLength
	o.DeviceDescriptor = op.DeviceDescriptor
	o.DeviceDescriptorLength = op.DeviceDescriptorLength
	o.Return = op.Return
}
func (o *GetUniqueIDs3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetUniqueIDs3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetUniqueIDs3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CheckNetFTBindings3Operation structure represents the CprepCheckNetFtBindings3 operation
type xxx_CheckNetFTBindings3Operation struct {
	This   *dcom.ORPCThis `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat `idl:"name:That" json:"that"`
	Return int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CheckNetFTBindings3Operation) OpNum() int { return 4 }

func (o *xxx_CheckNetFTBindings3Operation) OpName() string {
	return "/IClusterStorage3/v0/CprepCheckNetFtBindings3"
}

func (o *xxx_CheckNetFTBindings3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CheckNetFTBindings3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CheckNetFTBindings3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CheckNetFTBindings3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CheckNetFTBindings3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CheckNetFTBindings3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CheckNetFTBindings3Request structure represents the CprepCheckNetFtBindings3 operation request
type CheckNetFTBindings3Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CheckNetFTBindings3Request) xxx_ToOp(ctx context.Context) *xxx_CheckNetFTBindings3Operation {
	if o == nil {
		return &xxx_CheckNetFTBindings3Operation{}
	}
	return &xxx_CheckNetFTBindings3Operation{
		This: o.This,
	}
}

func (o *CheckNetFTBindings3Request) xxx_FromOp(ctx context.Context, op *xxx_CheckNetFTBindings3Operation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CheckNetFTBindings3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CheckNetFTBindings3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CheckNetFTBindings3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CheckNetFTBindings3Response structure represents the CprepCheckNetFtBindings3 operation response
type CheckNetFTBindings3Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepCheckNetFtBindings3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CheckNetFTBindings3Response) xxx_ToOp(ctx context.Context) *xxx_CheckNetFTBindings3Operation {
	if o == nil {
		return &xxx_CheckNetFTBindings3Operation{}
	}
	return &xxx_CheckNetFTBindings3Operation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *CheckNetFTBindings3Response) xxx_FromOp(ctx context.Context, op *xxx_CheckNetFTBindings3Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CheckNetFTBindings3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CheckNetFTBindings3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CheckNetFTBindings3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CSVTestSetup3Operation structure represents the CprepCsvTestSetup3 operation
type xxx_CSVTestSetup3Operation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	TestShareGUID *dtyp.GUID     `idl:"name:TestShareGuid" json:"test_share_guid"`
	_             string         `idl:"name:Reserved;string"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CSVTestSetup3Operation) OpNum() int { return 5 }

func (o *xxx_CSVTestSetup3Operation) OpName() string {
	return "/IClusterStorage3/v0/CprepCsvTestSetup3"
}

func (o *xxx_CSVTestSetup3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CSVTestSetup3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// TestShareGuid {in} (1:{alias=GUID}(struct))
	{
		if o.TestShareGUID != nil {
			if err := o.TestShareGUID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dtyp.GUID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Reserved {in} (1:{string, alias=LPWSTR}*(1)[dim:0,string,null](wchar))
	{
		// reserved Reserved
		if err := ndr.WriteUTF16NString(ctx, w, ""); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CSVTestSetup3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// TestShareGuid {in} (1:{alias=GUID}(struct))
	{
		if o.TestShareGUID == nil {
			o.TestShareGUID = &dtyp.GUID{}
		}
		if err := o.TestShareGUID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Reserved {in} (1:{string, alias=LPWSTR,pointer=ref}*(1)[dim:0,string,null](wchar))
	{
		// reserved Reserved
		var _Reserved string
		if err := ndr.ReadUTF16NString(ctx, w, &_Reserved); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CSVTestSetup3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CSVTestSetup3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CSVTestSetup3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CSVTestSetup3Request structure represents the CprepCsvTestSetup3 operation request
type CSVTestSetup3Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// TestShareGuid: The client generates a GUID and passes it to all nodes. The GUID can
	// be used to form a unique share name and create the ClusPrepShareList.
	TestShareGUID *dtyp.GUID `idl:"name:TestShareGuid" json:"test_share_guid"`
}

func (o *CSVTestSetup3Request) xxx_ToOp(ctx context.Context) *xxx_CSVTestSetup3Operation {
	if o == nil {
		return &xxx_CSVTestSetup3Operation{}
	}
	return &xxx_CSVTestSetup3Operation{
		This:          o.This,
		TestShareGUID: o.TestShareGUID,
	}
}

func (o *CSVTestSetup3Request) xxx_FromOp(ctx context.Context, op *xxx_CSVTestSetup3Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.TestShareGUID = op.TestShareGUID
}
func (o *CSVTestSetup3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CSVTestSetup3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CSVTestSetup3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CSVTestSetup3Response structure represents the CprepCsvTestSetup3 operation response
type CSVTestSetup3Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepCsvTestSetup3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CSVTestSetup3Response) xxx_ToOp(ctx context.Context) *xxx_CSVTestSetup3Operation {
	if o == nil {
		return &xxx_CSVTestSetup3Operation{}
	}
	return &xxx_CSVTestSetup3Operation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *CSVTestSetup3Response) xxx_FromOp(ctx context.Context, op *xxx_CSVTestSetup3Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *CSVTestSetup3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CSVTestSetup3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CSVTestSetup3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsNodeClustered3Operation structure represents the CprepIsNodeClustered3 operation
type xxx_IsNodeClustered3Operation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	IsClusterNode bool           `idl:"name:pbIsClusterNode" json:"is_cluster_node"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_IsNodeClustered3Operation) OpNum() int { return 6 }

func (o *xxx_IsNodeClustered3Operation) OpName() string {
	return "/IClusterStorage3/v0/CprepIsNodeClustered3"
}

func (o *xxx_IsNodeClustered3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsNodeClustered3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsNodeClustered3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsNodeClustered3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsNodeClustered3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pbIsClusterNode {out} (1:{pointer=ref}*(1))(2:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.IsClusterNode); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsNodeClustered3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbIsClusterNode {out} (1:{pointer=ref}*(1))(2:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.IsClusterNode); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// IsNodeClustered3Request structure represents the CprepIsNodeClustered3 operation request
type IsNodeClustered3Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *IsNodeClustered3Request) xxx_ToOp(ctx context.Context) *xxx_IsNodeClustered3Operation {
	if o == nil {
		return &xxx_IsNodeClustered3Operation{}
	}
	return &xxx_IsNodeClustered3Operation{
		This: o.This,
	}
}

func (o *IsNodeClustered3Request) xxx_FromOp(ctx context.Context, op *xxx_IsNodeClustered3Operation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *IsNodeClustered3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *IsNodeClustered3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsNodeClustered3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsNodeClustered3Response structure represents the CprepIsNodeClustered3 operation response
type IsNodeClustered3Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbIsClusterNode: The address of a pointer to a BOOLEAN value. Returns TRUE if the
	// server is a node within a cluster.
	IsClusterNode bool `idl:"name:pbIsClusterNode" json:"is_cluster_node"`
	// Return: The CprepIsNodeClustered3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsNodeClustered3Response) xxx_ToOp(ctx context.Context) *xxx_IsNodeClustered3Operation {
	if o == nil {
		return &xxx_IsNodeClustered3Operation{}
	}
	return &xxx_IsNodeClustered3Operation{
		That:          o.That,
		IsClusterNode: o.IsClusterNode,
		Return:        o.Return,
	}
}

func (o *IsNodeClustered3Response) xxx_FromOp(ctx context.Context, op *xxx_IsNodeClustered3Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.IsClusterNode = op.IsClusterNode
	o.Return = op.Return
}
func (o *IsNodeClustered3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *IsNodeClustered3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsNodeClustered3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_CreateNewSMBShares3Operation structure represents the CprepCreateNewSmbShares3 operation
type xxx_CreateNewSMBShares3Operation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	SharePaths    []string       `idl:"name:ppwszSharePaths;size_is:(, pdwNumberOfPaths);string" json:"share_paths"`
	NumberOfPaths uint32         `idl:"name:pdwNumberOfPaths" json:"number_of_paths"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_CreateNewSMBShares3Operation) OpNum() int { return 7 }

func (o *xxx_CreateNewSMBShares3Operation) OpName() string {
	return "/IClusterStorage3/v0/CprepCreateNewSmbShares3"
}

func (o *xxx_CreateNewSMBShares3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNewSMBShares3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNewSMBShares3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNewSMBShares3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if o.SharePaths != nil && o.NumberOfPaths == 0 {
		o.NumberOfPaths = uint32(len(o.SharePaths))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNewSMBShares3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppwszSharePaths {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pdwNumberOfPaths]*(1)[dim:0,string,null](wchar))
	{
		if o.SharePaths != nil || o.NumberOfPaths > 0 {
			_ptr_ppwszSharePaths := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
				dimSize1 := uint64(o.NumberOfPaths)
				if err := w.WriteSize(dimSize1); err != nil {
					return err
				}
				sizeInfo := []uint64{
					dimSize1,
				}
				for i1 := range o.SharePaths {
					i1 := i1
					if uint64(i1) >= sizeInfo[0] {
						break
					}
					if o.SharePaths[i1] != "" {
						_ptr_ppwszSharePaths := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
							if err := ndr.WriteUTF16NString(ctx, w, o.SharePaths[i1]); err != nil {
								return err
							}
							return nil
						})
						if err := w.WritePointer(&o.SharePaths[i1], _ptr_ppwszSharePaths); err != nil {
							return err
						}
					} else {
						if err := w.WritePointer(nil); err != nil {
							return err
						}
					}
				}
				for i1 := len(o.SharePaths); uint64(i1) < sizeInfo[0]; i1++ {
					if err := w.WritePointer(nil); err != nil {
						return err
					}
				}
				return nil
			})
			if err := w.WritePointer(&o.SharePaths, _ptr_ppwszSharePaths); err != nil {
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
	// pdwNumberOfPaths {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NumberOfPaths); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_CreateNewSMBShares3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppwszSharePaths {out} (1:{string, pointer=ref}*(2)*(1))(2:{alias=LPWSTR}[dim:0,size_is=pdwNumberOfPaths]*(1)[dim:0,string,null](wchar))
	{
		_ptr_ppwszSharePaths := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
			sizeInfo := []uint64{
				0,
			}
			for sz1 := range sizeInfo {
				if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
					return err
				}
			}
			if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
				return fmt.Errorf("buffer overflow for size %d of array o.SharePaths", sizeInfo[0])
			}
			o.SharePaths = make([]string, sizeInfo[0])
			for i1 := range o.SharePaths {
				i1 := i1
				_ptr_ppwszSharePaths := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
					if err := ndr.ReadUTF16NString(ctx, w, &o.SharePaths[i1]); err != nil {
						return err
					}
					return nil
				})
				_s_ppwszSharePaths := func(ptr interface{}) { o.SharePaths[i1] = *ptr.(*string) }
				if err := w.ReadPointer(&o.SharePaths[i1], _s_ppwszSharePaths, _ptr_ppwszSharePaths); err != nil {
					return err
				}
			}
			return nil
		})
		_s_ppwszSharePaths := func(ptr interface{}) { o.SharePaths = *ptr.(*[]string) }
		if err := w.ReadPointer(&o.SharePaths, _s_ppwszSharePaths, _ptr_ppwszSharePaths); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pdwNumberOfPaths {out} (1:{pointer=ref}*(1))(2:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NumberOfPaths); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// CreateNewSMBShares3Request structure represents the CprepCreateNewSmbShares3 operation request
type CreateNewSMBShares3Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
}

func (o *CreateNewSMBShares3Request) xxx_ToOp(ctx context.Context) *xxx_CreateNewSMBShares3Operation {
	if o == nil {
		return &xxx_CreateNewSMBShares3Operation{}
	}
	return &xxx_CreateNewSMBShares3Operation{
		This: o.This,
	}
}

func (o *CreateNewSMBShares3Request) xxx_FromOp(ctx context.Context, op *xxx_CreateNewSMBShares3Operation) {
	if o == nil {
		return
	}
	o.This = op.This
}
func (o *CreateNewSMBShares3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *CreateNewSMBShares3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateNewSMBShares3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// CreateNewSMBShares3Response structure represents the CprepCreateNewSmbShares3 operation response
type CreateNewSMBShares3Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// ppwszSharePaths: On successful return, specifies the address of a pointer to a *pdwNumberOfPaths-sized
	// block of LPWSTRs. The server allocates and initializes the returned buffer. Callers
	// MUST free this memory when they are finished with it.
	SharePaths []string `idl:"name:ppwszSharePaths;size_is:(, pdwNumberOfPaths);string" json:"share_paths"`
	// pdwNumberOfPaths: The number of file share path names returned in ppwszSharePaths.
	NumberOfPaths uint32 `idl:"name:pdwNumberOfPaths" json:"number_of_paths"`
	// Return: The CprepCreateNewSmbShares3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *CreateNewSMBShares3Response) xxx_ToOp(ctx context.Context) *xxx_CreateNewSMBShares3Operation {
	if o == nil {
		return &xxx_CreateNewSMBShares3Operation{}
	}
	return &xxx_CreateNewSMBShares3Operation{
		That:          o.That,
		SharePaths:    o.SharePaths,
		NumberOfPaths: o.NumberOfPaths,
		Return:        o.Return,
	}
}

func (o *CreateNewSMBShares3Response) xxx_FromOp(ctx context.Context, op *xxx_CreateNewSMBShares3Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.SharePaths = op.SharePaths
	o.NumberOfPaths = op.NumberOfPaths
	o.Return = op.Return
}
func (o *CreateNewSMBShares3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *CreateNewSMBShares3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_CreateNewSMBShares3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_ConnectToNewSMBShares3Operation structure represents the CprepConnectToNewSmbShares3 operation
type xxx_ConnectToNewSMBShares3Operation struct {
	This          *dcom.ORPCThis `idl:"name:This" json:"this"`
	That          *dcom.ORPCThat `idl:"name:That" json:"that"`
	SharePaths    []string       `idl:"name:ppwszSharePaths;size_is:(dwNumberOfPaths, );string" json:"share_paths"`
	NumberOfPaths uint32         `idl:"name:dwNumberOfPaths" json:"number_of_paths"`
	Return        int32          `idl:"name:Return" json:"return"`
}

func (o *xxx_ConnectToNewSMBShares3Operation) OpNum() int { return 8 }

func (o *xxx_ConnectToNewSMBShares3Operation) OpName() string {
	return "/IClusterStorage3/v0/CprepConnectToNewSmbShares3"
}

func (o *xxx_ConnectToNewSMBShares3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if o.SharePaths != nil && o.NumberOfPaths == 0 {
		o.NumberOfPaths = uint32(len(o.SharePaths))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectToNewSMBShares3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// ppwszSharePaths {in} (1:{string, pointer=ref}*(1))(2:{alias=LPWSTR}[dim:0,size_is=dwNumberOfPaths]*(1)[dim:0,string,null](wchar))
	{
		dimSize1 := uint64(o.NumberOfPaths)
		if err := w.WriteSize(dimSize1); err != nil {
			return err
		}
		sizeInfo := []uint64{
			dimSize1,
		}
		for i1 := range o.SharePaths {
			i1 := i1
			if uint64(i1) >= sizeInfo[0] {
				break
			}
			if o.SharePaths[i1] != "" {
				_ptr_ppwszSharePaths := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
					if err := ndr.WriteUTF16NString(ctx, w, o.SharePaths[i1]); err != nil {
						return err
					}
					return nil
				})
				if err := w.WritePointer(&o.SharePaths[i1], _ptr_ppwszSharePaths); err != nil {
					return err
				}
			} else {
				if err := w.WritePointer(nil); err != nil {
					return err
				}
			}
		}
		for i1 := len(o.SharePaths); uint64(i1) < sizeInfo[0]; i1++ {
			if err := w.WritePointer(nil); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// dwNumberOfPaths {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.WriteData(o.NumberOfPaths); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectToNewSMBShares3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// ppwszSharePaths {in} (1:{string, pointer=ref}*(1))(2:{alias=LPWSTR}[dim:0,size_is=dwNumberOfPaths]*(1)[dim:0,string,null](wchar))
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
			return fmt.Errorf("buffer overflow for size %d of array o.SharePaths", sizeInfo[0])
		}
		o.SharePaths = make([]string, sizeInfo[0])
		for i1 := range o.SharePaths {
			i1 := i1
			_ptr_ppwszSharePaths := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
				if err := ndr.ReadUTF16NString(ctx, w, &o.SharePaths[i1]); err != nil {
					return err
				}
				return nil
			})
			_s_ppwszSharePaths := func(ptr interface{}) { o.SharePaths[i1] = *ptr.(*string) }
			if err := w.ReadPointer(&o.SharePaths[i1], _s_ppwszSharePaths, _ptr_ppwszSharePaths); err != nil {
				return err
			}
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// dwNumberOfPaths {in} (1:{alias=DWORD}(uint32))
	{
		if err := w.ReadData(&o.NumberOfPaths); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectToNewSMBShares3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectToNewSMBShares3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_ConnectToNewSMBShares3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// ConnectToNewSMBShares3Request structure represents the CprepConnectToNewSmbShares3 operation request
type ConnectToNewSMBShares3Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// ppwszSharePaths: Specifies the address of a dwNumberOfPaths-sized block of LPWSTRs.
	SharePaths    []string `idl:"name:ppwszSharePaths;size_is:(dwNumberOfPaths, );string" json:"share_paths"`
	NumberOfPaths uint32   `idl:"name:dwNumberOfPaths" json:"number_of_paths"`
}

func (o *ConnectToNewSMBShares3Request) xxx_ToOp(ctx context.Context) *xxx_ConnectToNewSMBShares3Operation {
	if o == nil {
		return &xxx_ConnectToNewSMBShares3Operation{}
	}
	return &xxx_ConnectToNewSMBShares3Operation{
		This:          o.This,
		SharePaths:    o.SharePaths,
		NumberOfPaths: o.NumberOfPaths,
	}
}

func (o *ConnectToNewSMBShares3Request) xxx_FromOp(ctx context.Context, op *xxx_ConnectToNewSMBShares3Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.SharePaths = op.SharePaths
	o.NumberOfPaths = op.NumberOfPaths
}
func (o *ConnectToNewSMBShares3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *ConnectToNewSMBShares3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectToNewSMBShares3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// ConnectToNewSMBShares3Response structure represents the CprepConnectToNewSmbShares3 operation response
type ConnectToNewSMBShares3Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepConnectToNewSmbShares3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *ConnectToNewSMBShares3Response) xxx_ToOp(ctx context.Context) *xxx_ConnectToNewSMBShares3Operation {
	if o == nil {
		return &xxx_ConnectToNewSMBShares3Operation{}
	}
	return &xxx_ConnectToNewSMBShares3Operation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *ConnectToNewSMBShares3Response) xxx_FromOp(ctx context.Context, op *xxx_ConnectToNewSMBShares3Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *ConnectToNewSMBShares3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *ConnectToNewSMBShares3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_ConnectToNewSMBShares3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_GetProperties3Operation structure represents the CprepDiskGetProps3 operation
type xxx_GetProperties3Operation struct {
	This           *dcom.ORPCThis         `idl:"name:This" json:"this"`
	That           *dcom.ORPCThat         `idl:"name:That" json:"that"`
	DiskID         *csvp.ClusterDiskID    `idl:"name:DiskId" json:"disk_id"`
	DiskProperties *csvp.DiskPropertiesEx `idl:"name:pDiskProps" json:"disk_properties"`
	Return         int32                  `idl:"name:Return" json:"return"`
}

func (o *xxx_GetProperties3Operation) OpNum() int { return 9 }

func (o *xxx_GetProperties3Operation) OpName() string {
	return "/IClusterStorage3/v0/CprepDiskGetProps3"
}

func (o *xxx_GetProperties3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProperties3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DiskId {in} (1:{alias=CPREP_DISKID}(struct))
	{
		if o.DiskID != nil {
			if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.ClusterDiskID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_GetProperties3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DiskId {in} (1:{alias=CPREP_DISKID}(struct))
	{
		if o.DiskID == nil {
			o.DiskID = &csvp.ClusterDiskID{}
		}
		if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProperties3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProperties3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pDiskProps {out} (1:{pointer=ref}*(1))(2:{alias=DISK_PROPS_EX}(struct))
	{
		if o.DiskProperties != nil {
			if err := o.DiskProperties.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.DiskPropertiesEx{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_GetProperties3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pDiskProps {out} (1:{pointer=ref}*(1))(2:{alias=DISK_PROPS_EX}(struct))
	{
		if o.DiskProperties == nil {
			o.DiskProperties = &csvp.DiskPropertiesEx{}
		}
		if err := o.DiskProperties.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// GetProperties3Request structure represents the CprepDiskGetProps3 operation request
type GetProperties3Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk for which to get the disk properties.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *GetProperties3Request) xxx_ToOp(ctx context.Context) *xxx_GetProperties3Operation {
	if o == nil {
		return &xxx_GetProperties3Operation{}
	}
	return &xxx_GetProperties3Operation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *GetProperties3Request) xxx_FromOp(ctx context.Context, op *xxx_GetProperties3Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *GetProperties3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *GetProperties3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetProperties3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// GetProperties3Response structure represents the CprepDiskGetProps3 operation response
type GetProperties3Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pDiskProps: The properties of the selected ClusPrepDisk.
	DiskProperties *csvp.DiskPropertiesEx `idl:"name:pDiskProps" json:"disk_properties"`
	// Return: The CprepDiskGetProps3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *GetProperties3Response) xxx_ToOp(ctx context.Context) *xxx_GetProperties3Operation {
	if o == nil {
		return &xxx_GetProperties3Operation{}
	}
	return &xxx_GetProperties3Operation{
		That:           o.That,
		DiskProperties: o.DiskProperties,
		Return:         o.Return,
	}
}

func (o *GetProperties3Response) xxx_FromOp(ctx context.Context, op *xxx_GetProperties3Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.DiskProperties = op.DiskProperties
	o.Return = op.Return
}
func (o *GetProperties3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *GetProperties3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_GetProperties3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsReadOnly3Operation structure represents the CprepDiskIsReadOnly3 operation
type xxx_IsReadOnly3Operation struct {
	This     *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID   *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	ReadOnly bool                `idl:"name:pbReadOnly" json:"read_only"`
	Return   int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_IsReadOnly3Operation) OpNum() int { return 10 }

func (o *xxx_IsReadOnly3Operation) OpName() string {
	return "/IClusterStorage3/v0/CprepDiskIsReadOnly3"
}

func (o *xxx_IsReadOnly3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsReadOnly3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DiskId {in} (1:{alias=CPREP_DISKID}(struct))
	{
		if o.DiskID != nil {
			if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.ClusterDiskID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *xxx_IsReadOnly3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DiskId {in} (1:{alias=CPREP_DISKID}(struct))
	{
		if o.DiskID == nil {
			o.DiskID = &csvp.ClusterDiskID{}
		}
		if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsReadOnly3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsReadOnly3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pbReadOnly {out} (1:{pointer=ref}*(1))(2:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.ReadOnly); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsReadOnly3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbReadOnly {out} (1:{pointer=ref}*(1))(2:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.ReadOnly); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// IsReadOnly3Request structure represents the CprepDiskIsReadOnly3 operation request
type IsReadOnly3Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk for which to return the disk writable
	// status.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
}

func (o *IsReadOnly3Request) xxx_ToOp(ctx context.Context) *xxx_IsReadOnly3Operation {
	if o == nil {
		return &xxx_IsReadOnly3Operation{}
	}
	return &xxx_IsReadOnly3Operation{
		This:   o.This,
		DiskID: o.DiskID,
	}
}

func (o *IsReadOnly3Request) xxx_FromOp(ctx context.Context, op *xxx_IsReadOnly3Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
}
func (o *IsReadOnly3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *IsReadOnly3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsReadOnly3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsReadOnly3Response structure represents the CprepDiskIsReadOnly3 operation response
type IsReadOnly3Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbReadOnly: Returns a nonzero value if the LUN underlying the operating system disk
	// identified by ClusPrepDisk is not writable.
	ReadOnly bool `idl:"name:pbReadOnly" json:"read_only"`
	// Return: The CprepDiskIsReadOnly3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsReadOnly3Response) xxx_ToOp(ctx context.Context) *xxx_IsReadOnly3Operation {
	if o == nil {
		return &xxx_IsReadOnly3Operation{}
	}
	return &xxx_IsReadOnly3Operation{
		That:     o.That,
		ReadOnly: o.ReadOnly,
		Return:   o.Return,
	}
}

func (o *IsReadOnly3Response) xxx_FromOp(ctx context.Context, op *xxx_IsReadOnly3Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.ReadOnly = op.ReadOnly
	o.Return = op.Return
}
func (o *IsReadOnly3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *IsReadOnly3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsReadOnly3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PRRegister3Operation structure represents the CprepDiskPRRegister3 operation
type xxx_PRRegister3Operation struct {
	This     *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID   *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	OldPRKey uint64              `idl:"name:OldPrKey" json:"old_pr_key"`
	NewPRKey uint64              `idl:"name:NewPrKey" json:"new_pr_key"`
	Return   int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_PRRegister3Operation) OpNum() int { return 11 }

func (o *xxx_PRRegister3Operation) OpName() string {
	return "/IClusterStorage3/v0/CprepDiskPRRegister3"
}

func (o *xxx_PRRegister3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRRegister3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DiskId {in} (1:{alias=CPREP_DISKID}(struct))
	{
		if o.DiskID != nil {
			if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.ClusterDiskID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// OldPrKey {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.OldPRKey); err != nil {
			return err
		}
	}
	// NewPrKey {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.NewPRKey); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRRegister3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DiskId {in} (1:{alias=CPREP_DISKID}(struct))
	{
		if o.DiskID == nil {
			o.DiskID = &csvp.ClusterDiskID{}
		}
		if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// OldPrKey {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.OldPRKey); err != nil {
			return err
		}
	}
	// NewPrKey {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.NewPRKey); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRRegister3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRRegister3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRRegister3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PRRegister3Request structure represents the CprepDiskPRRegister3 operation request
type PRRegister3Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	// OldPrKey: The key used in the Reservation Key field of the SCSI PERSISTENT RESERVE
	// OUT command. If the value of OldPrKey is zero, the REGISTER_IGNORE_EXISTING action
	// is used. Otherwise, the REGISTER action is used.
	OldPRKey uint64 `idl:"name:OldPrKey" json:"old_pr_key"`
	// NewPrKey: The key used in the Service Action Reservation Key field of the SCSI PERSISTENT
	// RESERVE OUT command.
	NewPRKey uint64 `idl:"name:NewPrKey" json:"new_pr_key"`
}

func (o *PRRegister3Request) xxx_ToOp(ctx context.Context) *xxx_PRRegister3Operation {
	if o == nil {
		return &xxx_PRRegister3Operation{}
	}
	return &xxx_PRRegister3Operation{
		This:     o.This,
		DiskID:   o.DiskID,
		OldPRKey: o.OldPRKey,
		NewPRKey: o.NewPRKey,
	}
}

func (o *PRRegister3Request) xxx_FromOp(ctx context.Context, op *xxx_PRRegister3Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.OldPRKey = op.OldPRKey
	o.NewPRKey = op.NewPRKey
}
func (o *PRRegister3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PRRegister3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRRegister3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PRRegister3Response structure represents the CprepDiskPRRegister3 operation response
type PRRegister3Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskPRRegister3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PRRegister3Response) xxx_ToOp(ctx context.Context) *xxx_PRRegister3Operation {
	if o == nil {
		return &xxx_PRRegister3Operation{}
	}
	return &xxx_PRRegister3Operation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PRRegister3Response) xxx_FromOp(ctx context.Context, op *xxx_PRRegister3Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PRRegister3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PRRegister3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRRegister3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_FindKey3Operation structure represents the CprepDiskFindKey3 operation
type xxx_FindKey3Operation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Key    uint64              `idl:"name:Key" json:"key"`
	Found  bool                `idl:"name:pbFound" json:"found"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_FindKey3Operation) OpNum() int { return 12 }

func (o *xxx_FindKey3Operation) OpName() string { return "/IClusterStorage3/v0/CprepDiskFindKey3" }

func (o *xxx_FindKey3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FindKey3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DiskId {in} (1:{alias=CPREP_DISKID}(struct))
	{
		if o.DiskID != nil {
			if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.ClusterDiskID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Key {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.Key); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FindKey3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DiskId {in} (1:{alias=CPREP_DISKID}(struct))
	{
		if o.DiskID == nil {
			o.DiskID = &csvp.ClusterDiskID{}
		}
		if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Key {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.Key); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FindKey3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FindKey3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// pbFound {out} (1:{pointer=ref}*(1))(2:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.WriteData(o.Found); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_FindKey3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// pbFound {out} (1:{pointer=ref}*(1))(2:{alias=BOOLEAN, names=BYTE}(uchar))
	{
		if err := w.ReadData(&o.Found); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// FindKey3Request structure represents the CprepDiskFindKey3 operation request
type FindKey3Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	// Key: The registration key to search for in the SCSI Persistent Reserve registration
	// table for the LUN underlying the operating system disk identified by the DiskId parameter.
	// A value of zero for this parameter indicates that the caller is querying to discover
	// whether any keys are registered.
	Key uint64 `idl:"name:Key" json:"key"`
}

func (o *FindKey3Request) xxx_ToOp(ctx context.Context) *xxx_FindKey3Operation {
	if o == nil {
		return &xxx_FindKey3Operation{}
	}
	return &xxx_FindKey3Operation{
		This:   o.This,
		DiskID: o.DiskID,
		Key:    o.Key,
	}
}

func (o *FindKey3Request) xxx_FromOp(ctx context.Context, op *xxx_FindKey3Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.Key = op.Key
}
func (o *FindKey3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *FindKey3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FindKey3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// FindKey3Response structure represents the CprepDiskFindKey3 operation response
type FindKey3Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// pbFound: Returns a nonzero value if the registration key is found in the SCSI Persistent
	// Reserve registration table for the LUN underlying the operating system disk identified
	// by the DiskId parameter.
	Found bool `idl:"name:pbFound" json:"found"`
	// Return: The CprepDiskFindKey3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *FindKey3Response) xxx_ToOp(ctx context.Context) *xxx_FindKey3Operation {
	if o == nil {
		return &xxx_FindKey3Operation{}
	}
	return &xxx_FindKey3Operation{
		That:   o.That,
		Found:  o.Found,
		Return: o.Return,
	}
}

func (o *FindKey3Response) xxx_FromOp(ctx context.Context, op *xxx_FindKey3Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Found = op.Found
	o.Return = op.Return
}
func (o *FindKey3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *FindKey3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_FindKey3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PRPreempt3Operation structure represents the CprepDiskPRPreempt3 operation
type xxx_PRPreempt3Operation struct {
	This     *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That     *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID   *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	OwnerKey uint64              `idl:"name:OwnerKey" json:"owner_key"`
	NewKey   uint64              `idl:"name:NewKey" json:"new_key"`
	Return   int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_PRPreempt3Operation) OpNum() int { return 13 }

func (o *xxx_PRPreempt3Operation) OpName() string { return "/IClusterStorage3/v0/CprepDiskPRPreempt3" }

func (o *xxx_PRPreempt3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRPreempt3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DiskId {in} (1:{alias=CPREP_DISKID}(struct))
	{
		if o.DiskID != nil {
			if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.ClusterDiskID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// OwnerKey {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.OwnerKey); err != nil {
			return err
		}
	}
	// NewKey {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.NewKey); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRPreempt3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DiskId {in} (1:{alias=CPREP_DISKID}(struct))
	{
		if o.DiskID == nil {
			o.DiskID = &csvp.ClusterDiskID{}
		}
		if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// OwnerKey {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.OwnerKey); err != nil {
			return err
		}
	}
	// NewKey {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.NewKey); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRPreempt3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRPreempt3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRPreempt3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PRPreempt3Request structure represents the CprepDiskPRPreempt3 operation request
type PRPreempt3Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	// OwnerKey: The key used in the Service Action Reservation Key field of the SCSI PERSISTENT
	// RESERVE OUT command.
	OwnerKey uint64 `idl:"name:OwnerKey" json:"owner_key"`
	// NewKey: The key used in the Reservation Key field of the SCSI PERSISTENT RESERVE
	// OUT command.
	NewKey uint64 `idl:"name:NewKey" json:"new_key"`
}

func (o *PRPreempt3Request) xxx_ToOp(ctx context.Context) *xxx_PRPreempt3Operation {
	if o == nil {
		return &xxx_PRPreempt3Operation{}
	}
	return &xxx_PRPreempt3Operation{
		This:     o.This,
		DiskID:   o.DiskID,
		OwnerKey: o.OwnerKey,
		NewKey:   o.NewKey,
	}
}

func (o *PRPreempt3Request) xxx_FromOp(ctx context.Context, op *xxx_PRPreempt3Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.OwnerKey = op.OwnerKey
	o.NewKey = op.NewKey
}
func (o *PRPreempt3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PRPreempt3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRPreempt3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PRPreempt3Response structure represents the CprepDiskPRPreempt3 operation response
type PRPreempt3Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskPRPreempt3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PRPreempt3Response) xxx_ToOp(ctx context.Context) *xxx_PRPreempt3Operation {
	if o == nil {
		return &xxx_PRPreempt3Operation{}
	}
	return &xxx_PRPreempt3Operation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PRPreempt3Response) xxx_FromOp(ctx context.Context, op *xxx_PRPreempt3Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PRPreempt3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PRPreempt3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRPreempt3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PRReserve3Operation structure represents the CprepDiskPRReserve3 operation
type xxx_PRReserve3Operation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Key    uint64              `idl:"name:Key" json:"key"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_PRReserve3Operation) OpNum() int { return 14 }

func (o *xxx_PRReserve3Operation) OpName() string { return "/IClusterStorage3/v0/CprepDiskPRReserve3" }

func (o *xxx_PRReserve3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRReserve3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DiskId {in} (1:{alias=CPREP_DISKID}(struct))
	{
		if o.DiskID != nil {
			if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.ClusterDiskID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Key {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.Key); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRReserve3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DiskId {in} (1:{alias=CPREP_DISKID}(struct))
	{
		if o.DiskID == nil {
			o.DiskID = &csvp.ClusterDiskID{}
		}
		if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Key {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.Key); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRReserve3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRReserve3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRReserve3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PRReserve3Request structure represents the CprepDiskPRReserve3 operation request
type PRReserve3Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	// Key: The key used in the Reservation Key field of the SCSI PERSISTENT RESERVE OUT
	// command.
	Key uint64 `idl:"name:Key" json:"key"`
}

func (o *PRReserve3Request) xxx_ToOp(ctx context.Context) *xxx_PRReserve3Operation {
	if o == nil {
		return &xxx_PRReserve3Operation{}
	}
	return &xxx_PRReserve3Operation{
		This:   o.This,
		DiskID: o.DiskID,
		Key:    o.Key,
	}
}

func (o *PRReserve3Request) xxx_FromOp(ctx context.Context, op *xxx_PRReserve3Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.Key = op.Key
}
func (o *PRReserve3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PRReserve3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRReserve3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PRReserve3Response structure represents the CprepDiskPRReserve3 operation response
type PRReserve3Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskPRReserve3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PRReserve3Response) xxx_ToOp(ctx context.Context) *xxx_PRReserve3Operation {
	if o == nil {
		return &xxx_PRReserve3Operation{}
	}
	return &xxx_PRReserve3Operation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PRReserve3Response) xxx_FromOp(ctx context.Context, op *xxx_PRReserve3Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PRReserve3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PRReserve3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRReserve3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_IsPRPresent3Operation structure represents the CprepDiskIsPRPresent3 operation
type xxx_IsPRPresent3Operation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Key    uint64              `idl:"name:Key" json:"key"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_IsPRPresent3Operation) OpNum() int { return 15 }

func (o *xxx_IsPRPresent3Operation) OpName() string {
	return "/IClusterStorage3/v0/CprepDiskIsPRPresent3"
}

func (o *xxx_IsPRPresent3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPRPresent3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DiskId {in} (1:{alias=CPREP_DISKID}(struct))
	{
		if o.DiskID != nil {
			if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.ClusterDiskID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Key {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.Key); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPRPresent3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DiskId {in} (1:{alias=CPREP_DISKID}(struct))
	{
		if o.DiskID == nil {
			o.DiskID = &csvp.ClusterDiskID{}
		}
		if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Key {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.Key); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPRPresent3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPRPresent3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_IsPRPresent3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// IsPRPresent3Request structure represents the CprepDiskIsPRPresent3 operation request
type IsPRPresent3Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	// Key: The reservation key to search for in the SCSI Persistent Reserve reservation
	// table for the LUN underlying the operating system disk identified by the DiskId parameter.
	Key uint64 `idl:"name:Key" json:"key"`
}

func (o *IsPRPresent3Request) xxx_ToOp(ctx context.Context) *xxx_IsPRPresent3Operation {
	if o == nil {
		return &xxx_IsPRPresent3Operation{}
	}
	return &xxx_IsPRPresent3Operation{
		This:   o.This,
		DiskID: o.DiskID,
		Key:    o.Key,
	}
}

func (o *IsPRPresent3Request) xxx_FromOp(ctx context.Context, op *xxx_IsPRPresent3Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.Key = op.Key
}
func (o *IsPRPresent3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *IsPRPresent3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsPRPresent3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// IsPRPresent3Response structure represents the CprepDiskIsPRPresent3 operation response
type IsPRPresent3Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskIsPRPresent3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *IsPRPresent3Response) xxx_ToOp(ctx context.Context) *xxx_IsPRPresent3Operation {
	if o == nil {
		return &xxx_IsPRPresent3Operation{}
	}
	return &xxx_IsPRPresent3Operation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *IsPRPresent3Response) xxx_FromOp(ctx context.Context, op *xxx_IsPRPresent3Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *IsPRPresent3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *IsPRPresent3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_IsPRPresent3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PRRelease3Operation structure represents the CprepDiskPRRelease3 operation
type xxx_PRRelease3Operation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Key    uint64              `idl:"name:Key" json:"key"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_PRRelease3Operation) OpNum() int { return 16 }

func (o *xxx_PRRelease3Operation) OpName() string { return "/IClusterStorage3/v0/CprepDiskPRRelease3" }

func (o *xxx_PRRelease3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRRelease3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DiskId {in} (1:{alias=CPREP_DISKID}(struct))
	{
		if o.DiskID != nil {
			if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.ClusterDiskID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Key {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.Key); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRRelease3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DiskId {in} (1:{alias=CPREP_DISKID}(struct))
	{
		if o.DiskID == nil {
			o.DiskID = &csvp.ClusterDiskID{}
		}
		if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Key {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.Key); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRRelease3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRRelease3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRRelease3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PRRelease3Request structure represents the CprepDiskPRRelease3 operation request
type PRRelease3Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	// Key: The key used in the Reservation Key field of the SCSI PERSISTENT RESERVE OUT
	// command.
	Key uint64 `idl:"name:Key" json:"key"`
}

func (o *PRRelease3Request) xxx_ToOp(ctx context.Context) *xxx_PRRelease3Operation {
	if o == nil {
		return &xxx_PRRelease3Operation{}
	}
	return &xxx_PRRelease3Operation{
		This:   o.This,
		DiskID: o.DiskID,
		Key:    o.Key,
	}
}

func (o *PRRelease3Request) xxx_FromOp(ctx context.Context, op *xxx_PRRelease3Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.Key = op.Key
}
func (o *PRRelease3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PRRelease3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRRelease3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PRRelease3Response structure represents the CprepDiskPRRelease3 operation response
type PRRelease3Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskPRRelease3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PRRelease3Response) xxx_ToOp(ctx context.Context) *xxx_PRRelease3Operation {
	if o == nil {
		return &xxx_PRRelease3Operation{}
	}
	return &xxx_PRRelease3Operation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PRRelease3Response) xxx_FromOp(ctx context.Context, op *xxx_PRRelease3Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PRRelease3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PRRelease3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRRelease3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// xxx_PRClear3Operation structure represents the CprepDiskPRClear3 operation
type xxx_PRClear3Operation struct {
	This   *dcom.ORPCThis      `idl:"name:This" json:"this"`
	That   *dcom.ORPCThat      `idl:"name:That" json:"that"`
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	Key    uint64              `idl:"name:Key" json:"key"`
	Return int32               `idl:"name:Return" json:"return"`
}

func (o *xxx_PRClear3Operation) OpNum() int { return 17 }

func (o *xxx_PRClear3Operation) OpName() string { return "/IClusterStorage3/v0/CprepDiskPRClear3" }

func (o *xxx_PRClear3Operation) xxx_PrepareRequestPayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareRequestPayload(context.Context) error }); ok {
		if err := hook.AfterPrepareRequestPayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRClear3Operation) MarshalNDRRequest(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareRequestPayload(ctx); err != nil {
		return err
	}
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This != nil {
			if err := o.This.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThis{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// DiskId {in} (1:{alias=CPREP_DISKID}(struct))
	{
		if o.DiskID != nil {
			if err := o.DiskID.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&csvp.ClusterDiskID{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
	}
	// Key {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.WriteData(o.Key); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRClear3Operation) UnmarshalNDRRequest(ctx context.Context, w ndr.Reader) error {
	// This {in} (1:{alias=ORPCTHIS}(struct))
	{
		if o.This == nil {
			o.This = &dcom.ORPCThis{}
		}
		if err := o.This.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// DiskId {in} (1:{alias=CPREP_DISKID}(struct))
	{
		if o.DiskID == nil {
			o.DiskID = &csvp.ClusterDiskID{}
		}
		if err := o.DiskID.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
	}
	// Key {in} (1:{alias=ULONGLONG}(uint64))
	{
		if err := w.ReadData(&o.Key); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRClear3Operation) xxx_PrepareResponsePayload(ctx context.Context) error {
	if hook, ok := (interface{})(o).(interface{ AfterPrepareResponsePayload(context.Context) error }); ok {
		if err := hook.AfterPrepareResponsePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRClear3Operation) MarshalNDRResponse(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PrepareResponsePayload(ctx); err != nil {
		return err
	}
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That != nil {
			if err := o.That.MarshalNDR(ctx, w); err != nil {
				return err
			}
		} else {
			if err := (&dcom.ORPCThat{}).MarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		if err := w.WriteDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.WriteData(o.Return); err != nil {
			return err
		}
	}
	return nil
}

func (o *xxx_PRClear3Operation) UnmarshalNDRResponse(ctx context.Context, w ndr.Reader) error {
	// That {out} (1:{alias=ORPCTHAT}(struct))
	{
		if o.That == nil {
			o.That = &dcom.ORPCThat{}
		}
		if err := o.That.UnmarshalNDR(ctx, w); err != nil {
			return err
		}
		if err := w.ReadDeferred(); err != nil {
			return err
		}
	}
	// Return {out} (1:{alias=HRESULT, names=LONG}(int32))
	{
		if err := w.ReadData(&o.Return); err != nil {
			return err
		}
	}
	return nil
}

// PRClear3Request structure represents the CprepDiskPRClear3 operation request
type PRClear3Request struct {
	// This: ORPCTHIS structure that is used to send ORPC extension data to the server.
	This *dcom.ORPCThis `idl:"name:This" json:"this"`
	// DiskId: The identifier of the ClusPrepDisk representing the disk.
	DiskID *csvp.ClusterDiskID `idl:"name:DiskId" json:"disk_id"`
	// Key: The key used in the Reservation Key field of the SCSI PERSISTENT RESERVE OUT
	// command.
	Key uint64 `idl:"name:Key" json:"key"`
}

func (o *PRClear3Request) xxx_ToOp(ctx context.Context) *xxx_PRClear3Operation {
	if o == nil {
		return &xxx_PRClear3Operation{}
	}
	return &xxx_PRClear3Operation{
		This:   o.This,
		DiskID: o.DiskID,
		Key:    o.Key,
	}
}

func (o *PRClear3Request) xxx_FromOp(ctx context.Context, op *xxx_PRClear3Operation) {
	if o == nil {
		return
	}
	o.This = op.This
	o.DiskID = op.DiskID
	o.Key = op.Key
}
func (o *PRClear3Request) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRRequest(ctx, w)
}
func (o *PRClear3Request) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRClear3Operation{}
	if err := _o.UnmarshalNDRRequest(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}

// PRClear3Response structure represents the CprepDiskPRClear3 operation response
type PRClear3Response struct {
	// That: ORPCTHAT structure that is used to return ORPC extension data to the client.
	That *dcom.ORPCThat `idl:"name:That" json:"that"`
	// Return: The CprepDiskPRClear3 return value.
	Return int32 `idl:"name:Return" json:"return"`
}

func (o *PRClear3Response) xxx_ToOp(ctx context.Context) *xxx_PRClear3Operation {
	if o == nil {
		return &xxx_PRClear3Operation{}
	}
	return &xxx_PRClear3Operation{
		That:   o.That,
		Return: o.Return,
	}
}

func (o *PRClear3Response) xxx_FromOp(ctx context.Context, op *xxx_PRClear3Operation) {
	if o == nil {
		return
	}
	o.That = op.That
	o.Return = op.Return
}
func (o *PRClear3Response) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	return o.xxx_ToOp(ctx).MarshalNDRResponse(ctx, w)
}
func (o *PRClear3Response) UnmarshalNDR(ctx context.Context, r ndr.Reader) error {
	_o := &xxx_PRClear3Operation{}
	if err := _o.UnmarshalNDRResponse(ctx, r); err != nil {
		return err
	}
	o.xxx_FromOp(ctx, _o)
	return nil
}
